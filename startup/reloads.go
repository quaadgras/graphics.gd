//go:build reloads

package startup

import (
	"context"
	"os"
	"os/exec"
	"strings"
	"unsafe"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	gd "graphics.gd/internal"
	"graphics.gd/internal/callframe"
	"graphics.gd/internal/pointers"
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

	type (
		FUN = api.GoModuleFunc
		ARG = []api.ValueType
		RET = []api.ValueType
	)
	const (
		I32 = api.ValueTypeI32
		I64 = api.ValueTypeI64
	)

	var fn = newWasmRuntime()
	Wazero.NewHostModuleBuilder("gdextension").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.read_result_buffer), ARG{I32, I32, I32}, RET{I32}).Export("read_result_buffer").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.write_params_buffer), ARG{I32, I32, I64}, RET{}).Export("write_params_buffer").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.write_params_buffer8), ARG{I32, I32, I64, I64, I64, I64, I64, I64, I64, I64}, RET{}).Export("write_params_buffer8").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.get_godot_version), ARG{}, RET{I32, I32, I32, I32}).Export("get_godot_version").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.mem_alloc), ARG{I64}, RET{I64}).Export("mem_alloc").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.mem_realloc), ARG{I64, I64}, RET{I64}).Export("mem_realloc").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.mem_free), ARG{I64}, RET{}).Export("mem_free").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.print_error), ARG{I32, I32, I32, I32, I32, I32, I32, I32}, RET{}).Export("print_error").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.print_error_with_message), ARG{I32, I32, I32, I32, I32, I32, I32, I32, I32, I32}, RET{}).Export("print_error_with_message").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.print_warning), ARG{I32, I32, I32, I32, I32, I32, I32, I32}, RET{}).Export("print_warning").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.print_warning_with_message), ARG{I32, I32, I32, I32, I32, I32, I32, I32, I32}, RET{}).Export("print_warning_with_message").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.print_script_error), ARG{I32, I32, I32, I32, I32, I32, I32, I32}, RET{}).Export("print_script_error").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.print_script_error_with_message), ARG{I32, I32, I32, I32, I32, I32, I32, I32, I32}, RET{}).Export("print_script_error_with_message").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.get_native_struct_size), ARG{I64}, RET{I64}).Export("get_native_struct_size").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_new_copy), ARG{I64, I64, I64}, RET{I64}).Export("variant_new_copy").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_destroy), ARG{I64, I64, I64}, RET{}).Export("variant_destroy").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_call), ARG{I64, I64, I64, I64, I32}, RET{I32}).Export("variant_call").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_call_static), ARG{I32, I64, I32}, RET{I32}).Export("variant_call_static").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_evaluate), ARG{I32, I64, I64, I64, I64, I64, I64}, RET{I32}).Export("variant_evaluate").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_set), ARG{I64, I64, I64, I64, I64, I64, I64, I64, I64}, RET{I32}).Export("variant_set").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_set_named), ARG{I64, I64, I64, I32, I64, I64, I64}, RET{I32}).Export("variant_set_named").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_set_keyed), ARG{I64, I64, I64, I64, I64, I64, I64, I64, I64}, RET{I32}).Export("variant_set_keyed").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_set_indexed), ARG{I64, I64, I64, I64, I64, I64, I64}, RET{I32}).Export("variant_set_indexed").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_get), ARG{I64, I64, I64, I64, I64, I64}, RET{I32}).Export("variant_get").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_get_named), ARG{I64, I64, I64, I64}, RET{I32}).Export("variant_get_named").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_get_keyed), ARG{I64, I64, I64, I64, I64}, RET{I32}).Export("variant_get_keyed").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_get_indexed), ARG{I64, I64, I64}, RET{I32}).Export("variant_get_indexed").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_iter_init), ARG{I64, I64, I64}, RET{I32}).Export("variant_iter_init").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_iter_next), ARG{I64, I64, I64, I64, I64, I64}, RET{I32}).Export("variant_iter_next").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_iter_get), ARG{I64, I64, I64, I64, I64, I64}, RET{I32}).Export("variant_iter_get").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_hash), ARG{I64, I64, I64}, RET{I64}).Export("variant_hash").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_recursive_hash), ARG{I64, I64, I64}, RET{I64}).Export("variant_recursive_hash").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_hash_compare), ARG{I64, I64, I64}, RET{I32}).Export("variant_hash_compare").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_booleanize), ARG{I64, I64, I64}, RET{I32}).Export("variant_booleanize").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_duplicate), ARG{I64, I64, I64, I32}, RET{I64, I64, I64}).Export("variant_duplicate").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_stringify), ARG{I64, I64, I64}, RET{I32}).Export("variant_stringify").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_get_type), ARG{I64, I64, I64}, RET{I32}).Export("variant_get_type").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_has_method), ARG{I64, I64, I64, I64}, RET{I32}).Export("variant_has_method").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_has_member), ARG{I32, I64}, RET{I32}).Export("variant_has_member").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_has_key), ARG{I64, I64, I64, I64, I64, I64}, RET{I32}).Export("variant_has_key").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_get_object_instance_id), ARG{I64, I64, I64}, RET{I64}).Export("variant_get_object_instance_id").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_get_type_name), ARG{I32}, RET{I64}).Export("variant_get_type_name").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_can_convert), ARG{I32, I32}, RET{I32}).Export("variant_can_convert").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.variant_can_convert_strict), ARG{I32, I32}, RET{I32}).Export("variant_can_convert_static").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.get_variant_from_type_constructor), ARG{I32}, RET{I32}).Export("get_variant_from_type_constructor").
		NewFunctionBuilder().WithGoModuleFunction(FUN(fn.call_variant_to_type_constructor), ARG{I32}, RET{}).Export("call_variant_to_type_constructor").
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

type wasmRuntime struct {
	engine gd.API

	result *callframe.Frame
	frames [3]*callframe.Frame
	error  gd.CallError

	malloc api.Function
	stack  [1]uint64
	free   api.Function

	variant_from_type_constructors [gd.TypeMax]func(ret callframe.Ptr[gd.VariantPointers], arg callframe.Addr)
}

func newWasmRuntime() *wasmRuntime {
	var wasm wasmRuntime
	for i := range wasm.frames {
		wasm.result = callframe.New()
		wasm.frames[i] = callframe.New()
	}
	return &wasm
}

func (wasm *wasmRuntime) u32(v uint64) uint32 { return api.DecodeU32(v) }
func (wasm *wasmRuntime) i64(v uint64) gd.Int { return *(*gd.Int)(unsafe.Pointer(&v)) }
func (wasm *wasmRuntime) i32(v uint64) int32  { return api.DecodeI32(v) }
func (wasm *wasmRuntime) bool(v uint64) bool  { return v != 0 }
func (wasm *wasmRuntime) str(m api.Module, stack []uint64) string {
	buf, _ := m.Memory().Read(wasm.u32(stack[0]), wasm.u32(stack[1]))
	return string(buf)
}
func (wasm *wasmRuntime) variant(stack []uint64) gd.Variant {
	return pointers.Raw[gd.Variant]([3]uint64{stack[0], stack[1], stack[2]})
}
func (wasm *wasmRuntime) name(v uint64) gd.StringName {
	return pointers.Raw[gd.StringName]([1]gd.EnginePointer{gd.EnginePointer(v)})
}
func (wasm *wasmRuntime) returns(v []uint64) {
	for i := range v {
		callframe.Set64(wasm.result, 0, i, v[i])
	}
}

func (wasm *wasmRuntime) read_result_buffer(ctx context.Context, m api.Module, stack []uint64) {
	buffer, index, word := wasm.u32(stack[0]), wasm.u32(stack[1]), wasm.u32(stack[2])
	stack[0] = callframe.Get64(wasm.frames[buffer], index, word)
}
func (wasm *wasmRuntime) write_params_buffer(ctx context.Context, m api.Module, stack []uint64) {
	buffer, index := wasm.u32(stack[0]), wasm.u32(stack[1])
	callframe.Set64(wasm.frames[buffer], index, 0, stack[2])
}
func (wasm *wasmRuntime) write_params_buffer8(ctx context.Context, m api.Module, stack []uint64) {
	buffer, index := wasm.u32(stack[0]), wasm.u32(stack[1])
	for i := range 8 {
		callframe.Set64(wasm.frames[buffer], index, i, stack[2+i])
	}
}
func (wasm *wasmRuntime) get_godot_version(ctx context.Context, m api.Module, stack []uint64) {
	version := wasm.engine.GetGodotVersion()
	stack[0] = api.EncodeU32(version.Major)
	stack[1] = api.EncodeU32(version.Minor)
	stack[2] = api.EncodeU32(version.Patch)
	stack[3] = uint64(len(version.Value))
	wasm.malloc.CallWithStack(ctx, stack[3:4])
	m.Memory().Write(wasm.u32(stack[3]), []byte(version.Value))
}
func (wasm *wasmRuntime) mem_alloc(ctx context.Context, m api.Module, stack []uint64) {
	stack[0] = uint64(wasm.engine.Memory.Allocate(uintptr(stack[0])))
}
func (wasm *wasmRuntime) mem_realloc(ctx context.Context, m api.Module, stack []uint64) {
	stack[0] = uint64(wasm.engine.Memory.Reallocate(gd.Address(stack[0]), uintptr(stack[1])))
}
func (wasm *wasmRuntime) mem_free(ctx context.Context, m api.Module, stack []uint64) {
	wasm.engine.Memory.Free(gd.Address(stack[0]))
}
func (wasm *wasmRuntime) print_error(ctx context.Context, m api.Module, stack []uint64) {
	wasm.engine.PrintError(wasm.str(m, stack[0:2]), wasm.str(m, stack[2:4]), wasm.str(m, stack[4:6]), wasm.i32(stack[6]), wasm.bool(stack[7]))
}
func (wasm *wasmRuntime) print_error_with_message(ctx context.Context, m api.Module, stack []uint64) {
	wasm.engine.PrintErrorMessage(wasm.str(m, stack[0:2]), wasm.str(m, stack[2:4]), wasm.str(m, stack[4:6]), wasm.str(m, stack[6:8]), wasm.i32(stack[8]), wasm.bool(stack[9]))
}
func (wasm *wasmRuntime) print_warning(ctx context.Context, m api.Module, stack []uint64) {
	wasm.engine.PrintWarning(wasm.str(m, stack[0:2]), wasm.str(m, stack[2:4]), wasm.str(m, stack[4:6]), wasm.i32(stack[6]), wasm.bool(stack[7]))
}
func (wasm *wasmRuntime) print_warning_with_message(ctx context.Context, m api.Module, stack []uint64) {
	wasm.engine.PrintWarningMessage(wasm.str(m, stack[0:2]), wasm.str(m, stack[2:4]), wasm.str(m, stack[4:6]), wasm.str(m, stack[6:8]), wasm.i32(stack[8]), wasm.bool(stack[9]))
}
func (wasm *wasmRuntime) print_script_error(ctx context.Context, m api.Module, stack []uint64) {
	wasm.engine.PrintScriptError(wasm.str(m, stack[0:2]), wasm.str(m, stack[2:4]), wasm.str(m, stack[4:6]), wasm.i32(stack[6]), wasm.bool(stack[7]))
}
func (wasm *wasmRuntime) print_script_error_with_message(ctx context.Context, m api.Module, stack []uint64) {
	wasm.engine.PrintScriptErrorMessage(wasm.str(m, stack[0:2]), wasm.str(m, stack[2:4]), wasm.str(m, stack[4:6]), wasm.str(m, stack[6:8]), wasm.i32(stack[8]), wasm.bool(stack[9]))
}
func (wasm *wasmRuntime) get_native_struct_size(ctx context.Context, m api.Module, stack []uint64) {
	stack[0] = uint64(wasm.engine.GetNativeStructSize(wasm.name(stack[0])))
}
func (wasm *wasmRuntime) variant_new_copy(ctx context.Context, m api.Module, stack []uint64) {
	copy := wasm.engine.Variants.NewCopy(wasm.variant(stack[0:3]))
	raw, _ := pointers.End(copy)
	wasm.returns(raw[:])
}
func (wasm *wasmRuntime) variant_destroy(ctx context.Context, m api.Module, stack []uint64) {
	wasm.engine.Variants.Destroy(wasm.variant(stack[0:3]))
}
func (wasm *wasmRuntime) variant_call(ctx context.Context, m api.Module, stack []uint64) {
	var args []gd.Variant
	for i := range stack[4] {
		args = append(args, wasm.variant([]uint64{
			callframe.Get64(wasm.frames[0], i, 0),
			callframe.Get64(wasm.frames[0], i, 1),
			callframe.Get64(wasm.frames[0], i, 2),
		}))
	}
	result, err := wasm.engine.Variants.Call(wasm.variant(stack[0:3]), wasm.name(stack[4]), args...)
	if err != nil {
		wasm.error = *err.(*gd.CallError)
		stack[0] = uint64(wasm.error.ErrorType)
		return
	}
	stack[0] = 0
	raw, _ := pointers.End(result)
	wasm.returns(raw[:])
}
func (wasm *wasmRuntime) variant_call_static(ctx context.Context, m api.Module, stack []uint64) {
	vtype, method := gd.VariantType(wasm.u32(stack[0])), wasm.name(stack[1])
	var args []gd.Variant
	for i := range stack[4] {
		args = append(args, wasm.variant([]uint64{
			callframe.Get64(wasm.frames[0], i, 0),
			callframe.Get64(wasm.frames[0], i, 1),
			callframe.Get64(wasm.frames[0], i, 2),
		}))
	}
	result, err := wasm.engine.Variants.CallStatic(vtype, method, args...)
	if err != nil {
		wasm.error = *err.(*gd.CallError)
		stack[0] = uint64(wasm.error.ErrorType)
		return
	}
	stack[0] = 0
	raw, _ := pointers.End(result)
	wasm.returns(raw[:])
}
func (wasm *wasmRuntime) variant_evaluate(ctx context.Context, m api.Module, stack []uint64) {
	op := gd.Operator(wasm.u32(stack[0]))
	result, ok := wasm.engine.Variants.Evaluate(op, wasm.variant(stack[1:4]), wasm.variant(stack[4:7]))
	if !ok {
		stack[0] = 0
		return
	}
	raw, _ := pointers.End(result)
	wasm.returns(raw[:])
	stack[0] = 1 // Indicate success
}
func (wasm *wasmRuntime) variant_set(ctx context.Context, m api.Module, stack []uint64) {
	if wasm.engine.Variants.Set(wasm.variant(stack[0:3]), wasm.variant(stack[3:6]), wasm.variant(stack[6:9])) {
		stack[0] = 1
	}
	stack[0] = 0
}
func (wasm *wasmRuntime) variant_set_named(ctx context.Context, m api.Module, stack []uint64) {
	if wasm.engine.Variants.SetNamed(wasm.variant(stack[0:3]), wasm.name(stack[3]), wasm.variant(stack[4:7])) {
		stack[0] = 1
	}
	stack[0] = 0
}
func (wasm *wasmRuntime) variant_set_keyed(ctx context.Context, m api.Module, stack []uint64) {
	if wasm.engine.Variants.SetKeyed(wasm.variant(stack[0:3]), wasm.variant(stack[3:6]), wasm.variant(stack[6:9])) {
		stack[0] = 1
	}
	stack[0] = 0
}
func (wasm *wasmRuntime) variant_set_indexed(ctx context.Context, m api.Module, stack []uint64) {
	ok, oob := wasm.engine.Variants.SetIndexed(wasm.variant(stack[0:3]), wasm.i64(stack[1]), wasm.variant(stack[2:5]))
	stack[0] = 0
	if ok {
		stack[0] = 1
	}
	if oob {
		stack[0] = 2 // Out of bounds
	}
}
func (wasm *wasmRuntime) variant_get(ctx context.Context, m api.Module, stack []uint64) {
	result, ok := wasm.engine.Variants.Get(wasm.variant(stack[0:3]), wasm.variant(stack[3:6]))
	if !ok {
		stack[0] = 0
		return
	}
	raw, _ := pointers.End(result)
	wasm.returns(raw[:])
	stack[0] = 1 // Indicate success
}
func (wasm *wasmRuntime) variant_get_named(ctx context.Context, m api.Module, stack []uint64) {
	result, ok := wasm.engine.Variants.GetNamed(wasm.variant(stack[0:3]), wasm.name(stack[3]))
	if !ok {
		stack[0] = 0
		return
	}
	raw, _ := pointers.End(result)
	wasm.returns(raw[:])
	stack[0] = 1 // Indicate success
}
func (wasm *wasmRuntime) variant_get_keyed(ctx context.Context, m api.Module, stack []uint64) {
	result, ok := wasm.engine.Variants.GetKeyed(wasm.variant(stack[0:3]), wasm.variant(stack[3:6]))
	if !ok {
		stack[0] = 0
		return
	}
	raw, _ := pointers.End(result)
	wasm.returns(raw[:])
	stack[0] = 1 // Indicate success
}
func (wasm *wasmRuntime) variant_get_indexed(ctx context.Context, m api.Module, stack []uint64) {
	result, ok, oob := wasm.engine.Variants.GetIndexed(wasm.variant(stack[0:3]), wasm.i64(stack[3]))
	if !ok {
		stack[0] = 0
		return
	}
	if oob {
		stack[0] = 2 // Out of bounds
		return
	}
	raw, _ := pointers.End(result)
	wasm.returns(raw[:])
	stack[0] = 1 // Indicate success
}
func (wasm *wasmRuntime) variant_iter_init(ctx context.Context, m api.Module, stack []uint64) {
	result, ok := wasm.engine.Variants.IteratorInitialize(wasm.variant(stack[0:3]))
	if !ok {
		stack[0] = 0
		return
	}
	raw, _ := pointers.End(result)
	wasm.returns(raw[:])
	stack[0] = 1 // Indicate success
}
func (wasm *wasmRuntime) variant_iter_next(ctx context.Context, m api.Module, stack []uint64) {
	ok := wasm.engine.Variants.IteratorNext(wasm.variant(stack[0:3]), wasm.variant(stack[3:6]))
	if !ok {
		stack[0] = 0
		return
	}
	stack[0] = 1
}
func (wasm *wasmRuntime) variant_iter_get(ctx context.Context, m api.Module, stack []uint64) {
	result, ok := wasm.engine.Variants.IteratorGet(wasm.variant(stack[0:3]), wasm.variant(stack[3:6]))
	if !ok {
		stack[0] = 0
		return
	}
	raw, _ := pointers.End(result)
	wasm.returns(raw[:])
	stack[0] = 1 // Indicate success
}
func (wasm *wasmRuntime) variant_hash(ctx context.Context, m api.Module, stack []uint64) {
	hash := wasm.engine.Variants.Hash(wasm.variant(stack[0:3]))
	stack[0] = *(*uint64)(unsafe.Pointer(&hash))
}
func (wasm *wasmRuntime) variant_recursive_hash(ctx context.Context, m api.Module, stack []uint64) {
	hash := wasm.engine.Variants.RecursiveHash(wasm.variant(stack[0:3]), wasm.i64(stack[3]))
	stack[0] = *(*uint64)(unsafe.Pointer(&hash))
}
func (wasm *wasmRuntime) variant_hash_compare(ctx context.Context, m api.Module, stack []uint64) {
	if wasm.engine.Variants.HashCompare(wasm.variant(stack[0:3]), wasm.variant(stack[3:6])) {
		stack[0] = 1 // Indicate success
	} else {
		stack[0] = 0 // Indicate failure
	}
}
func (wasm *wasmRuntime) variant_booleanize(ctx context.Context, m api.Module, stack []uint64) {
	if wasm.engine.Variants.Booleanize(wasm.variant(stack[0:3])) {
		stack[0] = 1 // Indicate true
	} else {
		stack[0] = 0 // Indicate false
	}
}
func (wasm *wasmRuntime) variant_duplicate(ctx context.Context, m api.Module, stack []uint64) {
	copy := wasm.engine.Variants.Duplicate(wasm.variant(stack[0:3]), wasm.bool(stack[3]))
	raw, _ := pointers.End(copy)
	wasm.returns(raw[:])
}
func (wasm *wasmRuntime) variant_stringify(ctx context.Context, m api.Module, stack []uint64) {
	result := wasm.engine.Variants.Stringify(wasm.variant(stack[0:3]))
	raw, _ := pointers.End(result)
	stack[0] = uint64(raw[0])
}
func (wasm *wasmRuntime) variant_get_type(ctx context.Context, m api.Module, stack []uint64) {
	stack[0] = uint64(wasm.engine.Variants.GetType(wasm.variant(stack[0:3])))
}
func (wasm *wasmRuntime) variant_has_method(ctx context.Context, m api.Module, stack []uint64) {
	if wasm.engine.Variants.HasMethod(wasm.variant(stack[0:3]), wasm.name(stack[3])) {
		stack[0] = 1 // Indicate true
	} else {
		stack[0] = 0 // Indicate false
	}
}
func (wasm *wasmRuntime) variant_has_member(ctx context.Context, m api.Module, stack []uint64) {
	if wasm.engine.Variants.HasMember(gd.VariantType(wasm.u32(stack[0])), wasm.name(stack[1])) {
		stack[0] = 1 // Indicate true
	} else {
		stack[0] = 0 // Indicate false
	}
}
func (wasm *wasmRuntime) variant_has_key(ctx context.Context, m api.Module, stack []uint64) {
	hasKey, valid := wasm.engine.Variants.HasKey(wasm.variant(stack[0:3]), wasm.variant(stack[3:6]))
	if !valid {
		err := -1
		stack[0] = *(*uint64)(unsafe.Pointer(&err))
		return
	}
	if hasKey {
		stack[0] = 1 // Indicate true
	} else {
		stack[0] = 0 // Indicate false
	}
}
func (wasm *wasmRuntime) variant_get_object_instance_id(ctx context.Context, m api.Module, stack []uint64) {
	id := wasm.engine.Variants.GetObjectInstanceID(wasm.variant(stack[0:3]))
	stack[0] = *(*uint64)(unsafe.Pointer(&id))
}
func (wasm *wasmRuntime) variant_get_type_name(ctx context.Context, m api.Module, stack []uint64) {
	name := wasm.engine.Variants.GetTypeName(gd.VariantType(wasm.u32(stack[0])))
	raw, _ := pointers.End(name)
	stack[0] = uint64(raw[0])
}
func (wasm *wasmRuntime) variant_can_convert(ctx context.Context, m api.Module, stack []uint64) {
	if wasm.engine.Variants.CanConvert(gd.VariantType(wasm.u32(stack[0])), gd.VariantType(wasm.u32(stack[1]))) {
		stack[0] = 1 // Indicate true
	} else {
		stack[0] = 0 // Indicate false
	}
}
func (wasm *wasmRuntime) variant_can_convert_strict(ctx context.Context, m api.Module, stack []uint64) {
	if wasm.engine.Variants.CanConvertStrict(gd.VariantType(wasm.u32(stack[0])), gd.VariantType(wasm.u32(stack[1]))) {
		stack[0] = 1 // Indicate true
	} else {
		stack[0] = 0 // Indicate false
	}
}
func (wasm *wasmRuntime) get_variant_from_type_constructor(ctx context.Context, m api.Module, stack []uint64) {
	vtype := gd.VariantType(wasm.u32(stack[0]))
	wasm.variant_from_type_constructors[vtype] = wasm.engine.Variants.FromTypeConstructor(vtype)
}
func (wasm *wasmRuntime) call_variant_to_type_constructor(ctx context.Context, m api.Module, stack []uint64) {
	wasm.variant_from_type_constructors[gd.VariantType(wasm.u32(stack[0]))](callframe.GetPtr[gd.VariantPointers](wasm.result, 0), wasm.frames[0].Array(0).Index(0))
}
