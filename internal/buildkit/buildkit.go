// Package buildkit runs the Go toolchain with zero process spawns:
// cmd/compile and cmd/link cross-compiled to wasip1 and executed as
// wazero modules inside this process. That is the shape iOS forces —
// no exec — and it also buys perfect re-entrancy and crash isolation,
// because every invocation gets a fresh module instance with fresh
// globals. Cross-targeting is native: the guest's GOOS/GOARCH env picks
// the output platform, so a wasm-hosted toolchain on an iPhone emits
// ios/arm64 objects like any other cross-compile.
//
// On iOS the runtime must be wazero's interpreter (no JIT allowed);
// elsewhere the compiling runtime is used automatically.
package buildkit

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"runtime"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"github.com/tetratelabs/wazero/sys"
)

// Runner executes wasip1 toolchain modules. It caches compiled modules
// by content hash, so repeated invocations of the same tool skip
// recompilation of the wasm itself.
type Runner struct {
	runtime  wazero.Runtime
	compiled map[[sha256.Size]byte]wazero.CompiledModule
}

func NewRunner(ctx context.Context) *Runner {
	config := wazero.NewRuntimeConfig()
	if runtime.GOOS == "ios" {
		config = wazero.NewRuntimeConfigInterpreter()
	}
	r := &Runner{
		runtime:  wazero.NewRuntimeWithConfig(ctx, config),
		compiled: map[[sha256.Size]byte]wazero.CompiledModule{},
	}
	wasi_snapshot_preview1.MustInstantiate(ctx, r.runtime)
	return r
}

func (r *Runner) Close(ctx context.Context) { r.runtime.Close(ctx) }

// Invocation describes one tool run.
type Invocation struct {
	Tool   []byte            // the wasip1 module (e.g. compile.wasm)
	Args   []string          // argv, including argv[0]
	Env    map[string]string // e.g. GOOS/GOARCH/TMPDIR for the guest
	Root   string            // host directory mounted as the guest's /
	Stdout io.Writer
	Stderr io.Writer
}

// Run executes the invocation and returns an error carrying the tool's
// diagnostics if it exited non-zero.
func (r *Runner) Run(ctx context.Context, inv Invocation) error {
	compiled, err := r.compile(ctx, inv.Tool)
	if err != nil {
		return err
	}
	root := inv.Root
	if root == "" {
		root = "/"
	}
	config := wazero.NewModuleConfig().
		WithArgs(inv.Args...).
		WithFSConfig(wazero.NewFSConfig().WithDirMount(root, "/")).
		WithStdout(inv.Stdout).
		WithStderr(inv.Stderr).
		WithName("") // anonymous: concurrent instances of one tool may coexist
	for key, value := range inv.Env {
		config = config.WithEnv(key, value)
	}
	module, err := r.runtime.InstantiateModule(ctx, compiled, config)
	if module != nil {
		module.Close(ctx)
	}
	var exit *sys.ExitError
	if errors.As(err, &exit) {
		if exit.ExitCode() == 0 {
			return nil
		}
		return fmt.Errorf("%s exited with code %d", inv.Args[0], exit.ExitCode())
	}
	return err
}

func (r *Runner) compile(ctx context.Context, tool []byte) (wazero.CompiledModule, error) {
	key := sha256.Sum256(tool)
	if compiled, ok := r.compiled[key]; ok {
		return compiled, nil
	}
	compiled, err := r.runtime.CompileModule(ctx, tool)
	if err != nil {
		return nil, err
	}
	r.compiled[key] = compiled
	return compiled, nil
}
