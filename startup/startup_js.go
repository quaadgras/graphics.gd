package startup

import (
	"fmt"
	"syscall/js"
	"unsafe"

	gd "graphics.gd/internal"
	"graphics.gd/internal/callframe"
	"graphics.gd/internal/pointers"
)

func init() {
	GDExtension := js.Global().Get("GDExtension")
	if GDExtension.IsUndefined() {
		GDExtension = js.Global().Get("Object").New()
		js.Global().Set("GDExtension", GDExtension)
	}
	GDExtension.Set("res://library.gdextension", js.FuncOf(func(_ js.Value, args []js.Value) any {
		get_proc_func := args[0]
		linkJS(&gd.Global, func(s string) js.Value {
			return get_proc_func.Invoke(s)
		})
		initialization := args[2]
		initialization.Set("minimum_initialization_level", uint32(gd.GDExtensionInitializationLevelScene))
		initialization.Set("initialize", js.FuncOf(func(_ js.Value, args []js.Value) any {
			level := gd.GDExtensionInitializationLevel(args[0].Int())
			gd.Global.Init(level)
			if level == 2 {
				for _, fn := range gd.StartupFunctions {
					fn()
				}
				close(intialized)
				for _, fn := range gd.PostStartupFunctions {
					fn()
				}
			}
			return nil
		}))
		initialization.Set("deinitialize", js.FuncOf(func(_ js.Value, args []js.Value) any {
			for _, cleanup := range gd.Cleanups() {
				cleanup()
			}
			pointers.Cycle()
			pointers.Cycle()
			return nil
		}))
		return 0
	}))
}

var read_result_buffer js.Value
var write_params_buffer js.Value

func writeCallFrameArguments(buffer int, args callframe.Args) {
	for i := range args.Len() {
		writeCallFrameArgument(buffer, i, args.Index(i))
	}
}

func writeCallFrameArgument(buffer int, argn int, arg callframe.Addr) {
	val := arg.Pointer()
	write_params_buffer.Invoke(buffer, argn,
		uint32(val[0]), uint32(val[1]), uint32(val[2]), uint32(val[3]),
		uint32(val[4]), uint32(val[5]), uint32(val[6]), uint32(val[7]),
		uint32(val[8]), uint32(val[9]), uint32(val[10]), uint32(val[11]),
		uint32(val[12]), uint32(val[13]), uint32(val[14]), uint32(val[15]),
	)
}

func readCallFrameResult(buffer int, addr callframe.Addr) {
	result := addr.Pointer()
	for i := range result {
		result[i] = uintptr(read_result_buffer.Invoke(buffer, 0, i).Int())
	}
}

func linkJS(API *gd.API, dlsym func(string) js.Value) {
	read_result_buffer = dlsym("read_result_buffer")
	write_params_buffer = dlsym("write_params_buffer")

	string_name_new := dlsym("string_name_new")
	API.StringNames.New = func(s string) gd.StringName {
		return pointers.New[gd.StringName]([1]uintptr{uintptr(string_name_new.Invoke(s).Float())})
	}
	get_class_tag := dlsym("classdb_get_class_tag")
	API.ClassDB.GetClassTag = func(name gd.StringName) gd.ClassTag {
		return gd.ClassTag(get_class_tag.Invoke(pointers.Get(name)[0]).Float())
	}
	callable_custom_create2 := dlsym("callable_custom_create2")
	API.Callables.Create = func(fn func(...gd.Variant) (gd.Variant, error)) gd.Callable {
		var info = js.Global().Get("Object").New()
		info.Set("call_func", js.FuncOf(func(_ js.Value, js_args []js.Value) any {
			var argc = js_args[0].Int()
			var args = make([]gd.Variant, argc)
			for i := 0; i < argc; i++ {
				var raw gd.VariantPointers
				for j := 0; j < len(raw); j++ {
					raw[j] = uintptr(read_result_buffer.Invoke(0, i, j).Int())
				}
				args[i] = pointers.New[gd.Variant](raw)
			}
			result, err := fn(args...)
			if err != nil {
				return 1
			}
			raw := pointers.Get(result)
			for i := 0; i < len(raw); i++ {
				write_params_buffer.Invoke(0, 0, i)
			}
			return 0
		}))
		callable_custom_create2.Invoke(info)
		var raw gd.CallablePointers
		for i := 0; i < len(raw); i++ {
			raw[i] = uintptr(read_result_buffer.Invoke(0, 0, i).Int())
		}
		return pointers.New[gd.Callable](raw)
	}
	variant_get_ptr_constructor := dlsym("variant_get_ptr_constructor")
	call_variant_get_ptr_constructor := dlsym("call_variant_get_ptr_constructor")
	API.Variants.GetPointerConstructor = func(vtype gd.VariantType, index int32) func(base callframe.Addr, args callframe.Args) {
		fn := variant_get_ptr_constructor.Invoke(uint32(vtype), index)
		return func(base callframe.Addr, args callframe.Args) {
			writeCallFrameArguments(0, args)
			call_variant_get_ptr_constructor.Invoke(fn)
			readCallFrameResult(0, base)
		}
	}
	variant_get_ptr_operator_evaluator := dlsym("variant_get_ptr_operator_evaluator")
	call_variant_get_ptr_operator_evaluator := dlsym("call_variant_get_ptr_operator_evaluator")
	API.Variants.PointerOperatorEvaluator = func(op gd.Operator, a, b gd.VariantType) func(a, b, ret callframe.Addr) {
		fn := variant_get_ptr_operator_evaluator.Invoke(uint32(op), uint32(a), uint32(b))
		return func(a callframe.Addr, b callframe.Addr, ret callframe.Addr) {
			writeCallFrameArgument(0, 0, a)
			writeCallFrameArgument(1, 0, b)
			call_variant_get_ptr_operator_evaluator.Invoke(fn)
			readCallFrameResult(0, ret)
		}
	}
	variant_get_ptr_destructor := dlsym("variant_get_ptr_destructor")
	call_variant_get_ptr_destructor := dlsym("call_variant_get_ptr_destructor")
	API.Variants.GetPointerDestructor = func(vtype gd.VariantType) func(base callframe.Addr) {
		fn := variant_get_ptr_destructor.Invoke(uint32(vtype))
		return func(base callframe.Addr) {
			writeCallFrameArgument(0, 0, base)
			call_variant_get_ptr_destructor.Invoke(fn)
		}
	}
	get_variant_from_type_constructor := dlsym("get_variant_from_type_constructor")
	call_get_variant_from_type_constructor := dlsym("call_get_variant_from_type_constructor")
	API.Variants.FromTypeConstructor = func(vt gd.VariantType) func(ret callframe.Ptr[gd.VariantPointers], arg callframe.Addr) {
		fn := get_variant_from_type_constructor.Invoke(uint32(vt))
		return func(ret callframe.Ptr[gd.VariantPointers], arg callframe.Addr) {
			writeCallFrameArgument(0, 0, arg)
			call_get_variant_from_type_constructor.Invoke(fn)
			readCallFrameResult(0, ret.Addr())
		}
	}
	get_variant_to_type_constructor := dlsym("get_variant_to_type_constructor")
	call_get_variant_to_type_constructor := dlsym("call_get_variant_to_type_constructor")
	API.Variants.ToTypeConstructor = func(vt gd.VariantType) func(ret callframe.Addr, arg callframe.Ptr[gd.VariantPointers]) {
		fn := get_variant_to_type_constructor.Invoke(uint32(vt))
		return func(ret callframe.Addr, arg callframe.Ptr[gd.VariantPointers]) {
			writeCallFrameArgument(0, 0, arg.Addr())
			call_get_variant_to_type_constructor.Invoke(fn)
			readCallFrameResult(0, ret)
		}
	}
	variant_get_ptr_utility_function := dlsym("variant_get_ptr_utility_function")
	call_variant_get_ptr_utility_function := dlsym("call_variant_get_ptr_utility_function")
	API.Variants.GetPointerUtilityFunction = func(name gd.StringName, hash gd.Int) func(ret callframe.Addr, args callframe.Args, c int32) {
		fn := variant_get_ptr_utility_function.Invoke(pointers.Get(name)[0], uint32(hash))
		return func(ret callframe.Addr, args callframe.Args, c int32) {
			writeCallFrameArguments(0, args)
			call_variant_get_ptr_utility_function.Invoke(fn, c)
			readCallFrameResult(0, ret)
		}
	}
	variant_get_ptr_builtin_method := dlsym("variant_get_ptr_builtin_method")
	call_variant_get_ptr_builtin_method := dlsym("call_variant_get_ptr_builtin_method")
	API.Variants.GetPointerBuiltinMethod = func(vt gd.VariantType, sn gd.StringName, i gd.Int) func(base callframe.Addr, args callframe.Args, ret callframe.Addr, c int32) {
		fn := variant_get_ptr_builtin_method.Invoke(uint32(vt), pointers.Get(sn)[0], uint32(i))
		return func(base callframe.Addr, args callframe.Args, ret callframe.Addr, c int32) {
			writeCallFrameArguments(0, args)
			call_variant_get_ptr_builtin_method.Invoke(fn, c)
			readCallFrameResult(0, ret)
		}
	}
	classdb_get_method_bind := dlsym("classdb_get_method_bind")
	fmt.Println(classdb_get_method_bind)
	API.ClassDB.GetMethodBind = func(class, method gd.StringName, hash gd.Int) gd.MethodBind {
		return gd.MethodBind(classdb_get_method_bind.Invoke(pointers.Get(class)[0], pointers.Get(method)[0], *(*float64)(unsafe.Pointer(&hash))).Int())
	}
}
