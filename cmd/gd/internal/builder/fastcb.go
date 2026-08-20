package builder

import (
	"encoding/json"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"graphics.gd/cmd/gd/internal/gdpaths"
	"graphics.gd/cmd/gd/internal/tooling"
)

// The resident-callback runtime patch ("fastcb"): a replacement
// runtime/cgocall.go, applied as a `go build` overlay, that lets graphics.gd
// keep the engine thread's goroutine _Grunning with its P wired across
// engine->Go callbacks. graphics.gd detects the patch at runtime (see
// graphics.gd/internal/fastcb) and engages it automatically, so applying the
// overlay is all a build needs — no build tags and no application wiring.
//
//go:embed bundled/fastcb/cgocall.go.overlay
var fastcb_cgocall []byte

// The fused resident outbound crossing (see fastcb_callc_amd64.s.overlay):
// three companion files added to package runtime alongside the cgocall.go
// replacement. The .s and its amd64 declaration file are arch-gated by
// filename; the stub keeps fastcbCallCFastPC defined elsewhere. go build
// overlays support adding files that do not exist in the original tree.
//
//go:embed bundled/fastcb/fastcb_callc_amd64.s.overlay
var fastcb_callc_asm []byte

//go:embed bundled/fastcb/fastcb_callc_amd64.go.overlay
var fastcb_callc_decl []byte

//go:embed bundled/fastcb/fastcb_callc_arm64.s.overlay
var fastcb_callc_asm_arm64 []byte

//go:embed bundled/fastcb/fastcb_callc_arm64.go.overlay
var fastcb_callc_decl_arm64 []byte

//go:embed bundled/fastcb/fastcb_callc_windows_amd64.s.overlay
var fastcb_callc_asm_windows []byte

//go:embed bundled/fastcb/fastcb_callc_stub.go.overlay
var fastcb_callc_stub []byte

// fastcbAllowed lists the targets where the fastcb patch is applied. On
// Windows the patch preserves the stock callback path's osPreemptExtEnter/
// Exit pairing and m.winsyscall save/restore across the resident fast path.
var fastcbAllowed = map[string]bool{
	"linux":   true,
	"macos":   true,
	"ios":     true,
	"android": true,
	"musl":    true,
	"windows": true,
}

// fastcbCgocall resolves the runtime/cgocall.go replacement to overlay for a
// target, or "" when the patch must not be applied. The bundled copy is the
// default; environment variables adjust it:
//
//	GD_NO_FASTCB        disable the patch entirely
//	GD_OVERLAY_CGOCALL  path to an alternative replacement cgocall.go
//
// The bundled copy tracks the go1.27 runtime, so any other toolchain builds
// stock rather than overlaying a mismatched file.
func fastcbCgocall(target string) string {
	if os.Getenv("GD_NO_FASTCB") != "" {
		return ""
	}
	if !fastcbAllowed[target] {
		if os.Getenv("GD_OVERLAY_CGOCALL") != "" {
			fmt.Fprintf(os.Stderr, "gd: ignoring GD_OVERLAY_CGOCALL for %v: the resident-callback runtime patch is not applied on this target\n", target)
		}
		return ""
	}
	version, err := tooling.Go.Output("env", "GOVERSION")
	if err != nil || !strings.HasPrefix(strings.TrimSpace(version), "go1.27") {
		fmt.Fprintf(os.Stderr, "gd: building without the resident-callback runtime patch: it tracks go1.27 and this toolchain is %q\n", strings.TrimSpace(version))
		return ""
	}
	if p := os.Getenv("GD_OVERLAY_CGOCALL"); p != "" {
		return p
	}
	bundled := filepath.Join(gdpaths.Lib, "fastcb", "cgocall.go")
	if err := os.MkdirAll(filepath.Dir(bundled), 0755); err == nil {
		err = os.WriteFile(bundled, fastcb_cgocall, 0644)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "gd: building without the resident-callback runtime patch: cannot write %v: %v\n", bundled, err)
		return ""
	}
	return bundled
}

// fastcbFlags returns the extra `go build` flags for a target: the merged
// -tags list (baseTags plus the optional GD_EXTRA_TAGS) and, on targets where
// the fastcb runtime patch applies, the -overlay flag mapping
// runtime/cgocall.go to the patched copy.
func fastcbFlags(target, baseTags string) []string {
	var flags []string
	if tags := mergeTags(target, baseTags); tags != "" {
		flags = append(flags, "-tags="+tags)
	}
	cgocall := fastcbCgocall(target)
	if cgocall == "" {
		return flags
	}
	GOROOT, err := tooling.Go.Output("env", "GOROOT")
	if err != nil {
		fmt.Fprintf(os.Stderr, "gd: building without the resident-callback runtime patch: cannot resolve GOROOT: %v\n", err)
		return flags
	}
	// An auto-downloaded toolchain (go.mod requiring a newer point release
	// than the installed go) lives beneath GOMODCACHE, where `go build`
	// refuses -overlay replacements outright — skip the patch rather than
	// fail the whole build.
	GOMODCACHE, err := tooling.Go.Output("env", "GOMODCACHE")
	if err == nil {
		if cache := strings.TrimSpace(GOMODCACHE); cache != "" && strings.HasPrefix(strings.TrimSpace(GOROOT), cache+string(filepath.Separator)) {
			fmt.Fprintf(os.Stderr, "gd: building without the resident-callback runtime patch: the selected toolchain was auto-downloaded into the module cache, where overlays are not permitted — install a toolchain matching go.mod to enable the patch\n")
			return flags
		}
	}
	goroot := strings.TrimSpace(GOROOT)
	replace := map[string]string{
		filepath.Join(goroot, "src", "runtime", "cgocall.go"): cgocall,
	}
	// The fused-crossing companions are pure additions to package runtime;
	// if writing them fails the build proceeds with the cgocall.go patch
	// alone (fastcbCallCFastPC is then undefined only if cgocall.go is the
	// bundled copy, so treat a partial write as all-or-nothing).
	if companions, err := fastcbCallCFiles(); err == nil {
		for name, path := range companions {
			replace[filepath.Join(goroot, "src", "runtime", name)] = path
		}
	} else {
		fmt.Fprintf(os.Stderr, "gd: building without the resident-callback runtime patch: cannot write the fused-crossing files: %v\n", err)
		return flags
	}
	// Windows needs two more runtime replacements for extra-M retention: see
	// fastcbWindowsRetention. Without them every C->Go callback pays a
	// needm/dropm cycle, which both defeats resident-callback mode and races
	// the garbage collector's scan of the extra M's goroutine (dropm's
	// casgstatus spin) — deadlocking or killing the engine mid-frame.
	if target == "windows" {
		if libinit, cgoWindows, proc, ok := fastcbWindowsRetention(goroot); ok {
			replace[filepath.Join(goroot, "src", "runtime", "cgo", "gcc_libinit_windows.c")] = libinit
			replace[filepath.Join(goroot, "src", "runtime", "cgo", "windows.go")] = cgoWindows
			replace[filepath.Join(goroot, "src", "runtime", "proc.go")] = proc
		} else {
			fmt.Fprintf(os.Stderr, "gd: building without the resident-callback runtime patch: the windows extra-M retention patch does not apply to this toolchain\n")
			return flags
		}
	}
	overlay, err := writeOverlay("fastcb.json", replace)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gd: building without the resident-callback runtime patch: cannot write overlay: %v\n", err)
		return flags
	}
	return append(flags, "-overlay="+overlay)
}

// fastcbCallCFiles writes the fused-crossing companion files under
// gdpaths.Lib/fastcb and returns their runtime-relative names mapped to the
// written paths.
func fastcbCallCFiles() (map[string]string, error) {
	dir := filepath.Join(gdpaths.Lib, "fastcb")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	files := map[string][]byte{
		"fastcb_callc_amd64.s":         fastcb_callc_asm,
		"fastcb_callc_amd64.go":        fastcb_callc_decl,
		"fastcb_callc_arm64.s":         fastcb_callc_asm_arm64,
		"fastcb_callc_arm64.go":        fastcb_callc_decl_arm64,
		"fastcb_callc_windows_amd64.s": fastcb_callc_asm_windows,
		"fastcb_callc_stub.go":         fastcb_callc_stub,
	}
	out := make(map[string]string, len(files))
	for name, blob := range files {
		path := filepath.Join(dir, name)
		if err := os.WriteFile(path, blob, 0644); err != nil {
			return nil, err
		}
		out[name] = path
	}
	return out, nil
}

// mergeTags returns base plus the GD_EXTRA_TAGS additions (comma-separated),
// warning and ignoring them on targets where the fastcb patch does not apply.
func mergeTags(target, base string) string {
	extra := os.Getenv("GD_EXTRA_TAGS")
	if extra == "" {
		return base
	}
	if !fastcbAllowed[target] {
		fmt.Fprintf(os.Stderr, "gd: ignoring GD_EXTRA_TAGS for %v\n", target)
		return base
	}
	if base == "" {
		return extra
	}
	return base + "," + extra
}

// writeOverlay marshals a runtime-overlay replacement map to a JSON file named
// name under gdpaths.Lib and returns its path. encoding/json (rather than
// hand-assembled JSON) is required: on a windows host the paths contain
// backslashes, which hand-assembled JSON turns into invalid escapes.
func writeOverlay(name string, replace map[string]string) (string, error) {
	blob, err := json.Marshal(struct{ Replace map[string]string }{replace})
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(gdpaths.Lib, 0755); err != nil {
		return "", err
	}
	path := filepath.Join(gdpaths.Lib, name)
	if err := os.WriteFile(path, blob, 0644); err != nil {
		return "", err
	}
	return path, nil
}

// The windows halves of the resident-callback patch: extra-M retention.
//
// Upstream Go retains a C thread's extra M between callbacks only on pthread
// platforms (a pthread key destructor performs the deferred dropm at thread
// exit); on Windows the retention flag is never set, so every callback runs
// needm/dropm. gcc_libinit_windows.c.overlay implements the identical
// mechanism with fiber-local storage and publishes the retention flag, and a
// generated copy of proc.go drops cgoBindM's "bindm in unexpected GOOS"
// windows guard (the only reason bindm is 'unexpected' there is that the flag
// was assumed impossible to set). Since go1.27 runtime/cgo declares the
// windows _cgo_bindm as a nil pointer in its own windows.go, so that file is
// replaced too, importing the C x_cgo_bindm the way callbacks_unix.go does.
//
//go:embed bundled/fastcb/gcc_libinit_windows.c.overlay
var fastcb_libinit_windows []byte

//go:embed bundled/fastcb/windows.go.overlay
var fastcb_cgo_windows []byte

// fastcbWindowsRetention writes the libinit replacement and generates the
// patched proc.go from the active toolchain's own source, so it tracks point
// releases; if the guard's text ever changes shape the strict single-match
// requirement fails and the caller builds without the patch instead.
func fastcbWindowsRetention(goroot string) (libinit, cgoWindows, proc string, ok bool) {
	src, err := os.ReadFile(filepath.Join(goroot, "src", "runtime", "proc.go"))
	if err != nil {
		return "", "", "", false
	}
	const guard = `if GOOS == "windows" || GOOS == "plan9" {`
	if strings.Count(string(src), guard) != 1 {
		return "", "", "", false
	}
	patched := strings.Replace(string(src), guard, `if GOOS == "plan9" {`, 1)
	dir := filepath.Join(gdpaths.Lib, "fastcb")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", "", "", false
	}
	libinit = filepath.Join(dir, "gcc_libinit_windows.c")
	cgoWindows = filepath.Join(dir, "cgo_windows.go")
	proc = filepath.Join(dir, "proc.go")
	if os.WriteFile(libinit, fastcb_libinit_windows, 0644) != nil ||
		os.WriteFile(cgoWindows, fastcb_cgo_windows, 0644) != nil ||
		os.WriteFile(proc, []byte(patched), 0644) != nil {
		return "", "", "", false
	}
	return libinit, cgoWindows, proc, true
}
