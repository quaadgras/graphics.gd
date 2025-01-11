package startup

import (
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
			close(shutdown)
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
	if addr.Uintptr() == 0 {
		return
	}
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
		return pointers.New[gd.StringName]([1]uintptr{uintptr(string_name_new.Invoke(s).Int())})
	}
	string_new := dlsym("string_new")
	API.Strings.New = func(s string) gd.String {
		return pointers.New[gd.String]([1]uintptr{uintptr(string_new.Invoke(s).Int())})
	}
	string_get := dlsym("string_get")
	API.Strings.Get = func(s gd.String) string {
		return string_get.Invoke(pointers.Get(s)[0]).String()
	}
	get_godot_version := dlsym("get_godot_version")
	API.GetGodotVersion = func() gd.Version {
		version := js.Global().Get("Object").New()
		get_godot_version.Invoke(version)
		return gd.Version{
			Major: uint32(version.Get("major").Int()),
			Minor: uint32(version.Get("minor").Int()),
			Patch: uint32(version.Get("patch").Int()),
			Value: version.Get("string").String(),
		}
	}
	get_native_struct_size := dlsym("get_native_struct_size")
	API.GetNativeStructSize = func(name gd.StringName) uintptr {
		return uintptr(get_native_struct_size.Invoke(pointers.Get(name)[0]).Int())
	}
	get_library_path := dlsym("get_library_path")
	API.GetLibraryPath = func(token gd.ExtensionToken) gd.String {
		return pointers.New[gd.String]([1]uintptr{uintptr(get_library_path.Invoke(uint32(token)).Int())})
	}
	get_class_tag := dlsym("classdb_get_class_tag")
	API.ClassDB.GetClassTag = func(name gd.StringName) gd.ClassTag {
		return gd.ClassTag(get_class_tag.Invoke(pointers.Get(name)[0]).Int())
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
	variant_get_type_name := dlsym("variant_get_type_name")
	API.Variants.GetTypeName = func(vt gd.VariantType) gd.String {
		return pointers.New[gd.String]([1]uintptr{uintptr(variant_get_type_name.Invoke(uint32(vt)).Int())})
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
	API.ClassDB.GetMethodBind = func(class, method gd.StringName, hash gd.Int) gd.MethodBind {
		return gd.MethodBind(classdb_get_method_bind.Invoke(pointers.Get(class)[0], pointers.Get(method)[0], *(*float64)(unsafe.Pointer(&hash))).Int())
	}
	classdb_register_extension_class3 := dlsym("classdb_register_extension_class3")
	API.ClassDB.RegisterClass = func(library gd.ExtensionToken, name, extends gd.StringName, info_go gd.ClassInterface) {
		info := js.Global().Get("Object").New()
		info.Set("is_virtual", info_go.IsVirtual())
		info.Set("is_abstract", info_go.IsAbstract())
		info.Set("is_exposed", info_go.IsExposed())
		info.Set("is_runtime", false)
		info.Set("create_instance_func", js.FuncOf(func(_ js.Value, args []js.Value) any {
			return pointers.Get(info_go.CreateInstance()[0])[0]
		}))
		info.Set("get_virtual_call_data_func", js.FuncOf(func(_ js.Value, args []js.Value) any {
			p_class := args[0].Int()
			p_name := args[1].Int()
			var name = pointers.Let[gd.StringName]([1]uintptr{uintptr(p_name)})
			virtual := cgoHandle(p_class).Value().(gd.ClassInterface).GetVirtual(name)
			if virtual == nil {
				return 0
			}
			return uintptr(cgoNewHandle(virtual))
		}))
		info.Set("call_virtual_with_data_func", js.FuncOf(func(_ js.Value, args []js.Value) any {
			/*p_instance := args[0].Int()
			p_name := args[1].Int()
			p_data := args[2].Int()
			var name = pointers.Let[gd.StringName]([1]uintptr{uintptr(p_name)})*/
			panic("not supported")
			//cgoHandle(p_instance).Value().(gd.ObjectInterface).CallVirtual(name, cgoHandle(p_data).Value(), gd.UnsafeArgs(p_args), gd.UnsafeBack(p_ret))
		}))
		classdb_register_extension_class3.Invoke(uint32(library), pointers.Get(name)[0], pointers.Get(extends)[0], info)
	}
	API.ClassDB.RegisterClassMethod = func(library gd.ExtensionToken, class gd.StringName, info gd.Method) {}
}
