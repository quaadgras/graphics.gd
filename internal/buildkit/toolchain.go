package buildkit

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Module names as staged on disk. These are large build artifacts (tens
// of MB each), not committed to the repo: they are cross-compiled by
// internal/buildkit/patches (the Go toolchain) and .../lld-wasm (the
// linker), then placed in a modules directory the app ships or caches.
const (
	moduleCompile = "compile.wasm"  // cmd/compile, patched for ios
	moduleLink    = "link.wasm"     // cmd/link, patched for ios
	moduleLD64    = "ld64.lld.wasm" // ld64.lld Mach-O linker
)

// Toolchain is the named build operations of the on-device toolchain,
// each backed by a wasip1 module run through a Runner. It is the typed
// surface the rest of the harness calls; Runner stays the generic
// module executor beneath it.
type Toolchain struct {
	runner  *Runner
	compile []byte
	link    []byte
	ld64    []byte
}

// ModulesDir returns where the toolchain modules are staged: the
// GD_HARNESS_WASMTC_DIR override, else a per-user cache directory.
func ModulesDir() string {
	if dir := os.Getenv("GD_HARNESS_WASMTC_DIR"); dir != "" {
		return dir
	}
	cache, err := os.UserCacheDir()
	if err != nil {
		cache = os.TempDir()
	}
	return filepath.Join(cache, "gd-harness", "wasmtc")
}

// Load reads the toolchain modules from dir (see ModulesDir) into a
// ready Toolchain. Missing modules are reported by name so the caller
// can point the user at the build recipe. ld64.lld.wasm is optional:
// only Mach-O linking needs it.
func Load(ctx context.Context, dir string) (*Toolchain, error) {
	read := func(name string, required bool) ([]byte, error) {
		data, err := os.ReadFile(filepath.Join(dir, name))
		if os.IsNotExist(err) && !required {
			return nil, nil
		}
		if err != nil {
			return nil, fmt.Errorf("toolchain module %s: %w (build it per internal/buildkit/patches; set GD_HARNESS_WASMTC_DIR)", name, err)
		}
		return data, nil
	}
	compile, err := read(moduleCompile, true)
	if err != nil {
		return nil, err
	}
	link, err := read(moduleLink, true)
	if err != nil {
		return nil, err
	}
	ld64, err := read(moduleLD64, false)
	if err != nil {
		return nil, err
	}
	return &Toolchain{
		runner:  NewRunner(ctx),
		compile: compile,
		link:    link,
		ld64:    ld64,
	}, nil
}

func (t *Toolchain) Close(ctx context.Context) { t.runner.Close(ctx) }

// Target names an output platform for the guest toolchain.
type Target struct {
	GOOS, GOARCH string
}

// IOSARM64 is the target this toolchain exists to serve.
var IOSARM64 = Target{GOOS: "ios", GOARCH: "arm64"}

func (t Target) env(root, tmp string) map[string]string {
	if tmp == "" {
		tmp = "/tmp"
	}
	return map[string]string{"GOOS": t.GOOS, "GOARCH": t.GOARCH, "TMPDIR": tmp}
}

// Compile runs cmd/compile to turn Go sources into a package archive.
// importcfg lists the package's dependencies' export data; out is the
// .a to write. Paths are as the guest sees them under root (default /).
func (t *Toolchain) Compile(ctx context.Context, target Target, root, importcfg, out string, sources []string, log io.Writer) error {
	args := append([]string{
		"compile", "-p", "main", "-complete",
		"-importcfg", importcfg, "-pack", "-o", out,
	}, sources...)
	return t.runner.Run(ctx, Invocation{
		Tool: t.compile, Args: args, Env: target.env(root, ""),
		Root: root, Stdout: log, Stderr: log,
	})
}

// LinkGo runs cmd/link's internal linker to produce a pure-Go
// executable (no C archives). For ios/arm64 this needs the toolchain
// built from internal/buildkit/patches (which lifts Go's
// external-linking refusal); the result is a signed, installable
// Mach-O. Not for programs that link C — use LinkMachO.
func (t *Toolchain) LinkGo(ctx context.Context, target Target, root, importcfg, out string, archive string, log io.Writer) error {
	return t.runner.Run(ctx, Invocation{
		Tool: t.link, Args: []string{
			"link", "-importcfg", importcfg,
			"-buildmode=pie", "-linkmode=internal",
			"-o", out, archive,
		},
		Env: target.env(root, ""), Root: root, Stdout: log, Stderr: log,
	})
}

// LinkMachO runs ld64.lld to link a Mach-O binary from arbitrary
// objects, archives, and frameworks — the full-game link against
// libgodot's C++ archive that internal linking cannot do. args is the
// complete ld64 argument list after argv[0]; every path must be
// absolute because the guest's working directory is the mount root.
func (t *Toolchain) LinkMachO(ctx context.Context, root string, args []string, log io.Writer) error {
	if t.ld64 == nil {
		return fmt.Errorf("Mach-O linking needs %s (build it per internal/buildkit/patches/lld-wasm; set GD_HARNESS_WASMTC_DIR)", moduleLD64)
	}
	if root == "" {
		root = "/"
	}
	return t.runner.Run(ctx, Invocation{
		Tool: t.ld64, Args: append([]string{"ld64.lld"}, args...),
		Root: root, Stdout: log, Stderr: log,
	})
}

// CanLinkMachO reports whether the ld64.lld module is loaded.
func (t *Toolchain) CanLinkMachO() bool { return t.ld64 != nil }
