//go:build reloads

package startup

import (
	"context"
	"os"
	"os/exec"
	"strings"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	gd "graphics.gd/internal"
)

func init() {
	ctx := context.Background()
	Wazero := wazero.NewRuntime(ctx)
	wasi_snapshot_preview1.Instantiate(ctx, Wazero)

	wd, err := os.Getwd()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}

	cmd := exec.Command("go", "build", "-o", "./graphics/library.wasm")
	cmd.Env = append(os.Environ(), "GOOS=wasip1", "GOARCH=wasm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = strings.TrimSuffix(wd, "/graphics")
	if err := cmd.Run(); err != nil {
		os.Stderr.WriteString(err.Error())
	}

	wasm, err := os.ReadFile("./library.wasm")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}

	var gdextension wasmRuntime
	Wazero.NewHostModuleBuilder("gdextension").
		NewFunctionBuilder().
		WithGoModuleFunction(api.GoModuleFunc(gdextension.get_godot_version_major), nil, []api.ValueType{api.ValueTypeI32}).
		Export("get_godot_version_major").
		NewFunctionBuilder().
		WithGoModuleFunction(api.GoModuleFunc(gdextension.get_godot_version_minor), nil, []api.ValueType{api.ValueTypeI32}).
		Export("get_godot_version_minor").
		NewFunctionBuilder().
		WithGoModuleFunction(api.GoModuleFunc(gdextension.get_godot_version_patch), nil, []api.ValueType{api.ValueTypeI32}).
		Export("get_godot_version_patch").
		Instantiate(ctx)

	gd.StartupFunctions = append(gd.StartupFunctions, func() {
		program, err := Wazero.InstantiateWithConfig(ctx, wasm,
			wazero.NewModuleConfig().WithStderr(os.Stderr).WithStdout(os.Stdout),
		)
		if err != nil {
			os.Stderr.WriteString(err.Error())
			return
		}
		program.ExportedFunction("go")
	})
}

type wasmRuntime struct{}

func (wasm wasmRuntime) get_godot_version_major(ctx context.Context, m api.Module, stack []uint64) {
	stack[0] = api.EncodeU32(gd.Global.GetGodotVersion().Major)
}

func (wasm wasmRuntime) get_godot_version_minor(ctx context.Context, m api.Module, stack []uint64) {
	stack[0] = api.EncodeU32(gd.Global.GetGodotVersion().Minor)
}

func (wasm wasmRuntime) get_godot_version_patch(ctx context.Context, m api.Module, stack []uint64) {
	stack[0] = api.EncodeU32(gd.Global.GetGodotVersion().Patch)
}
