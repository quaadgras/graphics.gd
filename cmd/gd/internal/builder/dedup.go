package builder

import (
	"crypto/sha256"
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"graphics.gd/cmd/gd/internal/gdpaths"
	"graphics.gd/cmd/gd/internal/tooling"
)

// Shared generic instantiations ("dedup").
//
// A Go package that mentions an instantiated generic type compiles that
// instantiation itself (for a type, its whole method set) even when a
// package it imports has already compiled it, and the linker then throws
// the copies away (golang.org/issue/56718). Every generated classdb package
// mentions the variant containers (Array.Contains[T], Packed.Array[T],
// Dictionary.Map[K, V]...), so without this most of a cold build is spent
// compiling the same instantiations over and over.
//
// So gd builds its own copy of the toolchain's compiler, from the
// toolchain's own sources with a small patch (bundled/dedup). Each compile
// records in its archive the instantiations it provides, along with what
// export data records about an imported function (ABI, escape analysis
// tags, inlining cost); a compile then treats an instantiation that one of
// its direct imports provides exactly like an imported function: its body
// is not read, analysed or compiled, except where a call site inlines it.
// The compiler is used through a -toolexec wrapper added to GOFLAGS, so
// every go command gd runs (and the rebuilds of the hot-reload host, which
// inherits the environment) use it, and the go command keeps its build
// cache entries apart from those of the stock compiler.
//
// The copy is built once per toolchain, under gdpaths.Lib/dedup. Set
// GD_NO_DEDUP=1 to build with the stock compiler.

//go:embed bundled/dedup/extinst_gc.go.overlay
var dedup_gc []byte

//go:embed bundled/dedup/extinst_noder.go.overlay
var dedup_noder []byte

//go:embed bundled/dedup/toolexec.go.overlay
var dedup_toolexec []byte

// dedupEdits are the hooks the patch adds to the compiler's own files. Each
// old text must occur exactly once in the toolchain's copy of the file, or
// gd builds with the stock compiler.
var dedupEdits = []struct{ file, old, new string }{
	{
		"cmd/compile/internal/base/debug.go",
		"\tSlice                 int    `help:\"print information about slice compilation\"`\n",
		"\tSkipExtInst           int    `help:\"rely on imported packages for their generic instantiations (1 skip codegen, 3 treat as imported; 2 and 4 log)\" concurrent:\"ok\"`\n" +
			"\tSlice                 int    `help:\"print information about slice compilation\"`\n",
	},
	{
		"cmd/compile/internal/gc/compile.go",
		"\tif fn.IsClosure() {\n\t\treturn // we'll get this as part of its enclosing function\n\t}\n",
		"\tif fn.IsClosure() {\n\t\treturn // we'll get this as part of its enclosing function\n\t}\n" +
			"\tif skipExtInstFunc(fn) {\n\t\treturn // an imported package provides it (-d=skipextinst)\n\t}\n",
	},
	{
		"cmd/compile/internal/gc/obj.go",
		"\t\tfinishArchiveEntry(bout, start, \"_go_.o\")\n\t}\n",
		"\t\tfinishArchiveEntry(bout, start, \"_go_.o\")\n\t}\n" +
			"\tdumpExtInstMember(bout, mode) // -d=skipextinst\n",
	},
	{
		"cmd/compile/internal/noder/reader.go",
		"\t} else {\n\t\tr.addBody(name.Func, method)\n\t}\n",
		"\t} else if !r.hasTypeParams() || !r.importExtInst(name, method) {\n\t\tr.addBody(name.Func, method)\n\t}\n",
	},
	{
		"cmd/compile/internal/noder/reader.go",
		"\t\t!reflectdata.NeedEmit(tbase) {\n\t\treturn\n\t}\n",
		"\t\t!reflectdata.NeedEmit(tbase) ||\n\t\textInstWrapper(sym) {\n\t\treturn\n\t}\n",
	},
	{
		"cmd/compile/internal/noder/reader.go",
		"\tif !needed {\n\t\treturn\n\t}\n\n\taddTailCall(pos, fn, recv, method)\n",
		"\tif !needed || extInstWrapper(sym) {\n\t\treturn\n\t}\n\n\taddTailCall(pos, fn, recv, method)\n",
	},
}

// dedupNewFiles are the files the patch adds to the compiler.
var dedupNewFiles = map[string][]byte{
	"cmd/compile/internal/gc/extinst.go":    dedup_gc,
	"cmd/compile/internal/noder/extinst.go": dedup_noder,
}

// EnableDedup has the go commands that gd runs from now on compile with its
// copy of the compiler that shares generic instantiations between packages,
// building that copy first if need be. When it cannot, it says why and the
// builds use the stock compiler.
func EnableDedup() {
	if os.Getenv("GD_NO_DEDUP") != "" || strings.Contains(os.Getenv("GOFLAGS"), "-toolexec") {
		return // opted out, or the user runs their own -toolexec
	}
	if gomod, err := tooling.Go.Output("env", "GOMOD"); err != nil || gomod == "" || gomod == os.DevNull {
		return // no Go module here, so nothing for gd to compile
	}
	toolexec, err := dedupToolexec()
	if err != nil {
		fmt.Fprintf(os.Stderr, "gd: compiling with the stock Go compiler: %v\n", err)
		return
	}
	if strings.ContainsAny(toolexec, " \t") {
		// GOFLAGS is split on spaces.
		fmt.Fprintf(os.Stderr, "gd: compiling with the stock Go compiler: its path has a space in it: %v\n", toolexec)
		return
	}
	os.Setenv("GOFLAGS", strings.TrimSpace(os.Getenv("GOFLAGS")+" -toolexec="+toolexec))
}

// dedupToolexec returns the path of the -toolexec wrapper for the deduplicating
// compiler, building both for the toolchain in use if they are not built yet.
func dedupToolexec() (string, error) {
	env, err := tooling.Go.Output("env", "-json", "GOVERSION", "GOROOT", "GOHOSTOS", "GOHOSTARCH", "GOMODCACHE")
	if err != nil {
		return "", err
	}
	var goenv struct{ GOVERSION, GOROOT, GOHOSTOS, GOHOSTARCH, GOMODCACHE string }
	if err := json.Unmarshal([]byte(env), &goenv); err != nil {
		return "", err
	}
	// The patch hooks into compiler internals, which change between Go
	// releases.
	if !strings.HasPrefix(goenv.GOVERSION, "go1.27.") && goenv.GOVERSION != "go1.27" {
		return "", fmt.Errorf("the generic instantiation sharing patch tracks go1.27 and this toolchain is %v", goenv.GOVERSION)
	}
	// An auto-downloaded toolchain lives beneath GOMODCACHE, where go build
	// refuses -overlay replacements.
	if goenv.GOMODCACHE != "" && strings.HasPrefix(goenv.GOROOT, goenv.GOMODCACHE+string(filepath.Separator)) {
		return "", errors.New("the selected toolchain was auto-downloaded into the module cache, where overlays are not permitted: install a toolchain matching go.mod to share generic instantiations")
	}
	// Patch the toolchain's own sources.
	patched := make(map[string][]byte)
	for _, edit := range dedupEdits {
		src, ok := patched[edit.file]
		if !ok {
			if src, err = os.ReadFile(filepath.Join(goenv.GOROOT, "src", filepath.FromSlash(edit.file))); err != nil {
				return "", err
			}
		}
		if n := strings.Count(string(src), edit.old); n != 1 {
			return "", fmt.Errorf("the generic instantiation sharing patch does not apply to %v's %v", goenv.GOVERSION, edit.file)
		}
		patched[edit.file] = []byte(strings.Replace(string(src), edit.old, edit.new, 1))
	}
	for file, src := range dedupNewFiles {
		patched[file] = src
	}
	hash := sha256.New()
	fmt.Fprintf(hash, "%s %s %s/%s\n", goenv.GOVERSION, goenv.GOROOT, goenv.GOHOSTOS, goenv.GOHOSTARCH)
	for _, file := range slices.Sorted(maps.Keys(patched)) {
		fmt.Fprintf(hash, "%s %d\n", file, len(patched[file]))
		hash.Write(patched[file])
	}
	hash.Write(dedup_toolexec)
	dir := filepath.Join(gdpaths.Lib, "dedup", goenv.GOVERSION+"-"+hex.EncodeToString(hash.Sum(nil))[:12])
	exe := "" // the host's, not the target's GOEXE
	if goenv.GOHOSTOS == "windows" {
		exe = ".exe"
	}
	toolexec := filepath.Join(dir, "toolexec"+exe)
	if _, err := os.Stat(toolexec); err == nil {
		return toolexec, nil
	}

	fmt.Fprintf(os.Stderr, "gd: building a copy of the %v compiler that shares generic instantiations between packages (once per toolchain)\n", goenv.GOVERSION)
	if err := os.MkdirAll(filepath.Dir(dir), 0755); err != nil {
		return "", err
	}
	tmp, err := os.MkdirTemp(filepath.Dir(dir), "build-")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(tmp)
	replace := make(map[string]string)
	for file, src := range patched {
		path := filepath.Join(tmp, "src", filepath.FromSlash(file))
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return "", err
		}
		if err := os.WriteFile(path, src, 0644); err != nil {
			return "", err
		}
		replace[filepath.Join(goenv.GOROOT, "src", filepath.FromSlash(file))] = path
	}
	overlay, err := json.Marshal(struct{ Replace map[string]string }{replace})
	if err != nil {
		return "", err
	}
	if err := os.WriteFile(filepath.Join(tmp, "overlay.json"), overlay, 0644); err != nil {
		return "", err
	}
	out := filepath.Join(tmp, "out")
	// The compiler must be built by the toolchain whose sources were
	// patched, which the current directory's go.mod may have selected, so
	// it is built from here; the wrapper is plain Go that any toolchain
	// builds, so it is built as its own module.
	if err := dedupGo("", goenv.GOHOSTOS, goenv.GOHOSTARCH, "build", "-buildvcs=false", "-overlay="+filepath.Join(tmp, "overlay.json"), "-o", filepath.Join(out, "compile"+exe), "cmd/compile"); err != nil {
		return "", fmt.Errorf("building the compiler: %w", err)
	}
	shim := filepath.Join(tmp, "toolexec")
	if err := os.MkdirAll(shim, 0755); err != nil {
		return "", err
	}
	if err := os.WriteFile(filepath.Join(shim, "go.mod"), []byte("module toolexec\n\ngo 1.22\n"), 0644); err != nil {
		return "", err
	}
	if err := os.WriteFile(filepath.Join(shim, "main.go"), dedup_toolexec, 0644); err != nil {
		return "", err
	}
	if err := dedupGo(shim, goenv.GOHOSTOS, goenv.GOHOSTARCH, "build", "-buildvcs=false", "-o", filepath.Join(out, "toolexec"+exe), "."); err != nil {
		return "", fmt.Errorf("building the toolexec wrapper: %w", err)
	}
	if err := os.Rename(out, dir); err != nil {
		if _, statErr := os.Stat(toolexec); statErr != nil {
			return "", err
		}
		// another gd process got there first
	}
	return toolexec, nil
}

// dedupGo runs a go command in dir (the current directory if empty) that
// builds for the host, whatever target gd is building for.
func dedupGo(dir, goos, goarch string, args ...string) error {
	path, err := tooling.Go.Lookup()
	if err != nil {
		return err
	}
	cmd := exec.Command(path, args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GOOS="+goos, "GOARCH="+goarch, "GOFLAGS=", "CGO_ENABLED=0")
	if dir != "" {
		cmd.Env = append(cmd.Env, "GOTOOLCHAIN=local")
	}
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
