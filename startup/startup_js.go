package startup

import (
	"math"
	"syscall/js"
	"unsafe"

	gd "graphics.gd/internal"
	"graphics.gd/internal/callframe"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Packed"
)

func init() {
	GDExtension := js.Global().Get("GDExtension")
	if GDExtension.IsUndefined() {
		GDExtension = js.Global().Get("Object").New()
		js.Global().Set("GDExtension", GDExtension)
	}
	GDExtension.Set("res://library.gdextension", js.FuncOf(func(_ js.Value, args []js.Value) any {
		get_proc_func := args[0]
		dlsym = func(s string) js.Value {
			return get_proc_func.Invoke(s)
		}
		linkJS(&gd.Global)
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
var write_params_buffer16 js.Value

func writeCallFrameArguments(buffer int, args callframe.Args) {
	for i := range args.Len() {
		writeCallFrameArgument(buffer, i, args.Index(i))
	}
}

func writeCallFrameArgument(buffer int, argn int, arg callframe.Addr) {
	val := arg.Pointer()
	write_params_buffer16.Invoke(buffer, argn,
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
		result[i] = uint32(read_result_buffer.Invoke(buffer, 0, i).Int())
	}
}

var dlsym func(string) js.Value

func linkJS(API *gd.API) {
	read_result_buffer = dlsym("read_result_buffer")
	write_params_buffer = dlsym("write_params_buffer")
	write_params_buffer16 = dlsym("write_params_buffer16")

	string_name_new := dlsym("string_name_new")
	API.StringNames.New = func(s string) gd.StringName {
		return pointers.New[gd.StringName]([1]gd.EnginePointer{gd.EnginePointer(string_name_new.Invoke(s).Int())})
	}
	string_new := dlsym("string_new")
	API.Strings.New = func(s string) gd.String {
		return pointers.New[gd.String]([1]gd.EnginePointer{gd.EnginePointer(string_new.Invoke(s).Int())})
	}
	string_get := dlsym("string_get")
	API.Strings.Get = func(s gd.String) string {
		return string_get.Invoke(pointers.Get(s)[0], s.Len()).String()
	}
	string_operator_plus_eq_string := dlsym("string_operator_plus_eq_string")
	API.Strings.Append = func(s gd.String, other gd.String) {
		result := uint32(string_operator_plus_eq_string.Invoke(pointers.Get(s)[0], pointers.Get(other)[0]).Int())
		pointers.Set(s, [1]gd.EnginePointer{gd.EnginePointer(result)})
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
		return pointers.Let[gd.String]([1]gd.EnginePointer{gd.EnginePointer(get_library_path.Invoke(uint32(token)).Int())})
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
				var raw [6]uint32
				for j := 0; j < len(raw); j++ {
					raw[j] = uint32(read_result_buffer.Invoke(0, i, j).Int())
				}
				args[i] = pointers.Let[gd.Variant](*(*gd.VariantPointers)(unsafe.Pointer(&raw)))
			}
			result, err := fn(args...)
			if err != nil {
				return 1
			}
			raw := pointers.Get(result)
			buf := *(*[6]uint32)(unsafe.Pointer(&raw))
			for i := 0; i < len(buf); i++ {
				write_params_buffer.Invoke(0, 0, i, buf[i])
			}
			return 0
		}))
		callable_custom_create2.Invoke(info)
		var raw [4]uint32
		for i := 0; i < len(raw); i++ {
			raw[i] = uint32(read_result_buffer.Invoke(0, 0, i).Int())
		}
		return pointers.New[gd.Callable](*(*[2]uint64)(unsafe.Pointer(&raw)))
	}
	variant_get_type := dlsym("variant_get_type")
	API.Variants.GetType = func(v gd.Variant) gd.VariantType {
		var raw = pointers.Get(v)
		var vnt = *(*[6]uint32)(unsafe.Pointer(&raw))
		return gd.VariantType(variant_get_type.Invoke(
			vnt[0], vnt[1], vnt[2], vnt[3], vnt[4], vnt[5],
		).Int())
	}
	variant_destroy := dlsym("variant_destroy")
	API.Variants.Destroy = func(v gd.Variant) {
		var raw = pointers.Get(v)
		var vnt = *(*[6]uint32)(unsafe.Pointer(&raw))
		variant_destroy.Invoke(
			vnt[0], vnt[1], vnt[2], vnt[3], vnt[4], vnt[5],
		)
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
	call_variant_from_type_constructor := dlsym("call_variant_from_type_constructor")
	API.Variants.FromTypeConstructor = func(vt gd.VariantType) func(ret callframe.Ptr[gd.VariantPointers], arg callframe.Addr) {
		fn := get_variant_from_type_constructor.Invoke(uint32(vt))
		return func(ret callframe.Ptr[gd.VariantPointers], arg callframe.Addr) {
			writeCallFrameArgument(0, 0, arg)
			call_variant_from_type_constructor.Invoke(fn)
			readCallFrameResult(0, ret.Addr())
		}
	}
	get_variant_to_type_constructor := dlsym("get_variant_to_type_constructor")
	call_variant_to_type_constructor := dlsym("call_variant_to_type_constructor")
	API.Variants.ToTypeConstructor = func(vt gd.VariantType) func(ret callframe.Addr, arg callframe.Ptr[gd.VariantPointers]) {
		fn := get_variant_to_type_constructor.Invoke(uint32(vt))
		return func(ret callframe.Addr, arg callframe.Ptr[gd.VariantPointers]) {
			writeCallFrameArgument(0, 0, arg.Addr())
			call_variant_to_type_constructor.Invoke(fn)
			readCallFrameResult(0, ret)
		}
	}
	variant_get_ptr_utility_function := dlsym("variant_get_ptr_utility_function")
	call_variant_get_ptr_utility_function := dlsym("call_variant_get_ptr_utility_function")
	API.Variants.GetPointerUtilityFunction = func(name gd.StringName, hash gd.Int) func(ret callframe.Addr, args callframe.Args, c int32) {
		fn := variant_get_ptr_utility_function.Invoke(pointers.Get(name)[0], *(*float64)(unsafe.Pointer(&hash)))
		return func(ret callframe.Addr, args callframe.Args, c int32) {
			writeCallFrameArguments(0, args)
			call_variant_get_ptr_utility_function.Invoke(fn, c)
			readCallFrameResult(0, ret)
		}
	}
	variant_get_type_name := dlsym("variant_get_type_name")
	API.Variants.GetTypeName = func(vt gd.VariantType) gd.String {
		return pointers.Let[gd.String]([1]gd.EnginePointer{gd.EnginePointer(variant_get_type_name.Invoke(uint32(vt)).Int())})
	}
	variant_get_ptr_builtin_method := dlsym("variant_get_ptr_builtin_method")
	call_variant_get_ptr_builtin_method := dlsym("call_variant_get_ptr_builtin_method")
	API.Variants.GetPointerBuiltinMethod = func(vt gd.VariantType, sn gd.StringName, hash gd.Int) func(base callframe.Addr, args callframe.Args, ret callframe.Addr, c int32) {
		fn := variant_get_ptr_builtin_method.Invoke(uint32(vt), pointers.Get(sn)[0], *(*float64)(unsafe.Pointer(&hash)))
		return func(base callframe.Addr, args callframe.Args, ret callframe.Addr, c int32) {
			writeCallFrameArgument(1, 0, base)
			writeCallFrameArguments(0, args)
			call_variant_get_ptr_builtin_method.Invoke(fn, c)
			readCallFrameResult(0, ret)
			base.Pointer()[0] = uint32(read_result_buffer.Invoke(1, 0, 0).Int())
			base.Pointer()[1] = uint32(read_result_buffer.Invoke(1, 0, 1).Int())
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
		info.Set("create_instance", js.FuncOf(func(_ js.Value, args []js.Value) any {
			return pointers.Get(info_go.CreateInstance()[0])[0]
		}))
		info.Set("get_virtual_call_data", js.FuncOf(func(_ js.Value, args []js.Value) any {
			p_name := args[0].Int()
			var name = pointers.Let[gd.StringName]([1]gd.EnginePointer{gd.EnginePointer(p_name)})
			virtual := info_go.GetVirtual(name)
			if virtual == nil {
				return 0
			}
			return uintptr(cgoNewHandle(virtual))
		}))
		classdb_register_extension_class3.Invoke(uint32(library), pointers.Get(name)[0], pointers.Get(extends)[0], info)
	}
	classdb_construct_object := dlsym("classdb_construct_object")
	API.ClassDB.ConstructObject = func(class gd.StringName) [1]gd.Object {
		return [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(classdb_construct_object.Invoke(pointers.Get(class)[0]).Int())})}
	}
	object_method_bind_call := dlsym("object_bind_method_call")
	API.Object.MethodBindCall = func(method gd.MethodBind, obj [1]gd.Object, arg ...gd.Variant) (gd.Variant, error) {
		if obj == [1]gd.Object{} {
			panic("nil gd.Object dereference")
		}
		var self uint32
		if obj != ([1]gd.Object{}) {
			self = uint32(pointers.Get(obj[0])[0])
		}
		if self == 0 {
			panic("nil gd.Object dereference")
		}
		for i, v := range arg {
			raw := pointers.Get(v)
			buf := *(*[6]uint32)(unsafe.Pointer(&raw))
			for j := 0; j < len(buf); j++ {
				write_params_buffer.Invoke(0, i, j, buf[j])
			}
		}
		object_method_bind_call.Invoke(uint32(method), self, uint32(len(arg)))
		var raw [6]uint32
		for i := 0; i < len(raw); i++ {
			raw[i] = uint32(read_result_buffer.Invoke(0, 0, i).Int())
		}
		return pointers.New[gd.Variant](*(*[3]uint64)(unsafe.Pointer(&raw))), nil
	}
	object_method_bind_ptrcall := dlsym("object_method_bind_ptrcall")
	API.Object.MethodBindPointerCall = func(method gd.MethodBind, obj [1]gd.Object, arg callframe.Args, ret callframe.Addr) {
		var self uint32
		if obj != ([1]gd.Object{}) {
			self = uint32(pointers.Get(obj[0])[0])
		}
		writeCallFrameArguments(0, arg)
		object_method_bind_ptrcall.Invoke(uint32(method), self)
		readCallFrameResult(0, ret)
	}
	API.Object.MethodBindPointerCallStatic = func(method gd.MethodBind, arg callframe.Args, ret callframe.Addr) {
		writeCallFrameArguments(0, arg)
		object_method_bind_ptrcall.Invoke(uint32(method), 0)
		readCallFrameResult(0, ret)
	}
	global_get_singleton := dlsym("global_get_singleton")
	API.Object.GetSingleton = func(name gd.StringName) [1]gd.Object {
		return [1]gd.Object{pointers.Raw[gd.Object]([3]uint64{uint64(global_get_singleton.Invoke(pointers.Get(name)[0]).Int())})}
	}
	object_get_instance_from_id := dlsym("object_get_instance_from_id")
	API.Object.GetInstanceFromID = func(id gd.ObjectID) [1]gd.Object {
		fid := math.Float64frombits(uint64(id))
		return [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(object_get_instance_from_id.Invoke(fid).Int())})}
	}
	object_set_instance := dlsym("object_set_instance")
	API.Object.SetInstance = func(o [1]gd.Object, sn gd.StringName, oi gd.ObjectInterface) {
		wrapper := js.Global().Get("Object").New()
		methods := []js.Func{}
		cleanup := func(fn js.Func) js.Func {
			methods = append(methods, fn)
			return fn
		}
		handle := cgoNewHandle(oi)
		wrapper.Set("ref", uint32(handle))
		wrapper.Set("set", cleanup(js.FuncOf(func(_ js.Value, args []js.Value) any {
			var field = pointers.Let[gd.StringName]([1]gd.EnginePointer{gd.EnginePointer(args[0].Int())})
			var variant [6]uint32
			for i := 0; i < len(variant); i++ {
				variant[i] = uint32(args[i+1].Int())
			}
			return oi.Set(field, pointers.Let[gd.Variant](*(*[3]uint64)(unsafe.Pointer(&variant))))
		})))
		wrapper.Set("get", cleanup(js.FuncOf(func(_ js.Value, args []js.Value) any {
			var field = pointers.Let[gd.StringName]([1]gd.EnginePointer{gd.EnginePointer(args[0].Int())})
			variant, ok := oi.Get(field)
			if !ok {
				return false
			}
			var raw = pointers.Get(variant)
			buf := *(*[6]uint32)(unsafe.Pointer(&raw))
			for i := 0; i < len(buf); i++ {
				write_params_buffer.Invoke(0, 0, i, buf[i])
			}
			return true
		})))
		wrapper.Set("get_property_list", cleanup(js.FuncOf(func(_ js.Value, args []js.Value) any {
			var list = oi.GetPropertyList()
			var arr = js.Global().Get("Array").New(len(list))
			for i, item_go := range list {
				var item = js.Global().Get("Object").New()
				item.Set("name", pointers.Get(item_go.Name))
				item.Set("type", uint32(item_go.Type))
				item.Set("hint", uint32(item_go.Hint))
				item.Set("hint_string", pointers.Get(item_go.HintString))
				item.Set("usage", uint32(item_go.Usage))
				item.Set("class_name", pointers.Get(item_go.ClassName))
				arr.SetIndex(i, item)
			}
			return len(list)
		})))
		wrapper.Set("property_can_revert", cleanup(js.FuncOf(func(_ js.Value, args []js.Value) any {
			var field = pointers.Let[gd.StringName]([1]gd.EnginePointer{gd.EnginePointer(args[0].Int())})
			return oi.PropertyCanRevert(field)
		})))
		wrapper.Set("property_get_revert", cleanup(js.FuncOf(func(_ js.Value, args []js.Value) any {
			var field = pointers.Let[gd.StringName]([1]gd.EnginePointer{gd.EnginePointer(args[0].Int())})
			variant, ok := oi.PropertyGetRevert(field)
			if !ok {
				return false
			}
			var raw = pointers.Get(variant)
			buf := *(*[6]uint32)(unsafe.Pointer(&raw))
			for i := 0; i < len(buf); i++ {
				write_params_buffer.Invoke(0, 0, i, buf[i])
			}
			return true
		})))
		/*wrapper.Set("validate_property", cleanup(js.FuncOf(func(_ js.Value, args []js.Value) any {
		var object = args[0]
		var property gd.PropertyInfo
		property.Name = pointers.New[gd.StringName]([1]uintptr{uintptr(object.Get("name").Int())})
		property.Type = gd.VariantType(object.Get("type").Int())
		property.Hint = int64(object.Get("hint").Int())
		property.HintString = pointers.New[gd.String]([1]uintptr{uintptr(object.Get("hint_string").Int())})
		property.Usage = int64(object.Get("usage").Int())
		property.ClassName = pointers.New[gd.StringName]([1]uintptr{uintptr(object.Get("class_name").Int())})
		return oi.ValidateProperty(property)
		})))*/
		wrapper.Set("notification", cleanup(js.FuncOf(func(_ js.Value, args []js.Value) any {
			var what = int32(args[0].Int())
			var reversed = args[1].Bool()
			oi.Notification(what, reversed)
			return nil
		})))
		wrapper.Set("to_string", cleanup(js.FuncOf(func(_ js.Value, args []js.Value) any {
			str, ok := oi.ToString()
			if !ok {
				return false
			}
			write_params_buffer.Invoke(0, 0, 0, pointers.Get(str)[0])
			return true
		})))
		wrapper.Set("reference", cleanup(js.FuncOf(func(_ js.Value, args []js.Value) any {
			oi.Reference()
			return nil
		})))
		wrapper.Set("unreference", cleanup(js.FuncOf(func(_ js.Value, args []js.Value) any {
			oi.Unreference()
			return nil
		})))
		wrapper.Set("free", cleanup(js.FuncOf(func(_ js.Value, args []js.Value) any {
			oi.Free()
			for _, fn := range methods {
				fn.Release()
			}
			handle.Delete()
			return nil
		})))
		wrapper.Set("call_virtual", cleanup(js.FuncOf(func(_ js.Value, args []js.Value) any {
			name := pointers.Let[gd.StringName]([1]gd.EnginePointer{gd.EnginePointer(args[0].Int())})
			wrap := cgoHandle(args[1].Int()).Value()
			oi.CallVirtual(name, wrap, gd.Address(args[2].Int()), gd.Address(args[3].Int()))
			return nil
		})))
		object_set_instance.Invoke(pointers.Get(o[0])[0], pointers.Get(sn)[0], wrapper)
	}
	object_set_instance_binding := dlsym("object_set_instance_binding")
	API.Object.SetInstanceBinding = func(o [1]gd.Object, et gd.ExtensionToken, a any, ibt gd.InstanceBindingType) {
		object_set_instance_binding.Invoke(pointers.Get(o[0])[0], uint32(et), 0, 0)
	}
	object_get_instance_id := dlsym("object_get_instance_id")
	API.Object.GetInstanceID = func(o [1]gd.Object) gd.ObjectID {
		return gd.ObjectID(math.Float64bits(object_get_instance_id.Invoke(pointers.Get(o[0])[0]).Float()))
	}
	object_cast_to := dlsym("object_cast_to")
	API.Object.CastTo = func(o [1]gd.Object, sn gd.ClassTag) [1]gd.Object {
		casted := object_cast_to.Invoke(pointers.Get(o[0])[0], uint32(sn)).Int()
		if casted == 0 {
			return [1]gd.Object{}
		}
		return o
	}
	object_destroy := dlsym("object_destroy")
	API.Object.Destroy = func(o [1]gd.Object) {
		object_destroy.Invoke(pointers.Get(o[0])[0])
	}
	variant_new_nil := dlsym("variant_new_nil")
	API.Variants.NewNil = func() gd.Variant {
		variant_new_nil.Invoke()
		var variant [6]uint32
		for i := 0; i < len(variant); i++ {
			variant[i] = uint32(read_result_buffer.Invoke(0, 0, i).Int())
		}
		return pointers.New[gd.Variant](*(*[3]uint64)(unsafe.Pointer(&variant)))
	}
	API.PackedByteArray = makePackedFunctions[gd.PackedByteArray, byte]("byte_array")
	API.PackedInt32Array = makePackedFunctions[gd.PackedInt32Array, int32]("int32_array")
	API.PackedByteArray = makePackedFunctions[gd.PackedByteArray, byte]("byte_array")
	API.PackedColorArray = makePackedFunctions[gd.PackedColorArray, gd.Color]("color_array")
	API.PackedFloat32Array = makePackedFunctions[gd.PackedFloat32Array, float32]("float32_array")
	API.PackedFloat64Array = makePackedFunctions[gd.PackedFloat64Array, float64]("float64_array")
	API.PackedInt32Array = makePackedFunctions[gd.PackedInt32Array, int32]("int32_array")
	API.PackedInt64Array = makePackedFunctions[gd.PackedInt64Array, int64]("int64_array")
	API.PackedVector2Array = makePackedFunctions[gd.PackedVector2Array, gd.Vector2]("vector2_array")
	API.PackedVector3Array = makePackedFunctions[gd.PackedVector3Array, gd.Vector3]("vector3_array")
	API.PackedVector4Array = makePackedFunctions[gd.PackedVector4Array, gd.Vector4]("vector4_array")
	packed_string_array_operator_index := dlsym("packed_string_array_operator_index")
	API.PackedStringArray.Index = func(pia gd.PackedStringArray, index gd.Int) gd.String {
		arr := pointers.Get(pia)
		raw := *(*[2]uint32)(unsafe.Pointer(&arr))
		s := packed_string_array_operator_index.Invoke(int(raw[0]), int(raw[1]), uint32(index)).Int()
		return pointers.Let[gd.String]([1]gd.EnginePointer{gd.EnginePointer(s)})
	}
	packed_string_array_operator_index_set := dlsym("packed_string_array_operator_index_set")
	API.PackedStringArray.SetIndex = func(pia gd.PackedStringArray, index gd.Int, v gd.String) {
		arr := pointers.Get(pia)
		raw := *(*[2]uint32)(unsafe.Pointer(&arr))
		packed_string_array_operator_index_set.Invoke(int(raw[0]), int(raw[1]), uint32(index), pointers.Get(v)[0])
	}
	API.PackedStringArray.CopyAsSlice = func(pia gd.PackedStringArray) []gd.String {
		var slice = make([]gd.String, pia.Size())
		for i := 0; i < int(pia.Size()); i++ {
			slice[i] = API.PackedStringArray.Index(pia, gd.Int(i))
		}
		return slice
	}
	API.PackedStringArray.CopyFromSlice = func(pia gd.PackedStringArray, slice []gd.String) {
		for i, v := range slice {
			API.PackedStringArray.SetIndex(pia, gd.Int(i), v)
		}
	}
	classdb_register_extension_class_property := dlsym("classdb_register_extension_class_property")
	API.ClassDB.RegisterClassProperty = func(library gd.ExtensionToken, class gd.StringName, info gd.PropertyInfo, getter, setter gd.StringName) {
		converted := js.Global().Get("Object").New()
		converted.Set("name", pointers.Get(info.Name)[0])
		converted.Set("type", uint32(info.Type))
		converted.Set("hint", uint32(info.Hint))
		converted.Set("hint_string", pointers.Get(info.HintString)[0])
		converted.Set("usage", uint32(info.Usage))
		converted.Set("class_name", pointers.Get(info.ClassName)[0])
		classdb_register_extension_class_property.Invoke(uint32(library), pointers.Get(class)[0], converted, pointers.Get(getter)[0], pointers.Get(setter)[0])
	}
	classdb_register_extension_class_signal := dlsym("classdb_register_extension_class_signal")
	API.ClassDB.RegisterClassSignal = func(library gd.ExtensionToken, class, signal gd.StringName, args []gd.PropertyInfo) {
		converted := js.Global().Get("Array").New()
		for _, arg := range args {
			argument := js.Global().Get("Object").New()
			argument.Set("name", pointers.Get(arg.Name)[0])
			argument.Set("type", uint32(arg.Type))
			argument.Set("hint", uint32(arg.Hint))
			argument.Set("hint_string", pointers.Get(arg.HintString)[0])
			argument.Set("usage", uint32(arg.Usage))
			argument.Set("class_name", pointers.Get(arg.ClassName)[0])
			converted.Call("push", argument)
		}
		classdb_register_extension_class_signal.Invoke(uint32(library), pointers.Get(class)[0], pointers.Get(signal)[0], converted)
	}
	classdb_register_extension_class_method := dlsym("classdb_register_extension_class_method")
	API.ClassDB.RegisterClassMethod = func(library gd.ExtensionToken, class gd.StringName, info gd.Method) {
		converted := js.Global().Get("Object").New()
		converted.Set("name", pointers.Get(info.Name)[0])
		converted.Set("method_flags", uint32(info.MethodFlags))
		converted.Set("call", js.FuncOf(func(_ js.Value, args []js.Value) any {
			instance := cgoHandle(args[0].Int()).Value()
			arg_count := args[1].Int()
			var arguments = make([]gd.Variant, arg_count)
			for i := range arg_count {
				var raw [6]uint32
				for j := range raw {
					raw[j] = uint32(read_result_buffer.Invoke(0, i, j).Int())
				}
				arguments[i] = pointers.Let[gd.Variant](*(*[3]uint64)(unsafe.Pointer(&raw)))
			}
			result, err := info.Call(instance, arguments...)
			if err != nil {
				return 1
			}
			raw := pointers.Get(result)
			vnt := *(*[6]uint32)(unsafe.Pointer(&raw))
			for i := range vnt {
				write_params_buffer.Invoke(0, 0, i, vnt[i])
			}
			return 0
		}))
		converted.Set("ptrcall", js.FuncOf(func(_ js.Value, args []js.Value) any {
			info.PointerCall(cgoHandle(args[0].Int()).Value(), gd.Address(args[1].Int()), gd.Address(args[2].Int()))
			return nil
		}))
		if info.ReturnValueInfo != nil {
			returnValueInfo := js.Global().Get("Object").New()
			returnValueInfo.Set("name", pointers.Get(info.ReturnValueInfo.Name)[0])
			returnValueInfo.Set("type", uint32(info.ReturnValueInfo.Type))
			returnValueInfo.Set("hint", uint32(info.ReturnValueInfo.Hint))
			returnValueInfo.Set("hint_string", pointers.Get(info.ReturnValueInfo.HintString)[0])
			returnValueInfo.Set("usage", uint32(info.ReturnValueInfo.Usage))
			returnValueInfo.Set("class_name", pointers.Get(info.ReturnValueInfo.ClassName)[0])
			converted.Set("return_value_info", returnValueInfo)
			converted.Set("return_value_metadata", uint32(info.ReturnValueMetadata))
		}
		var arguments = js.Global().Get("Array").New()
		var argument_metadatas = js.Global().Get("Array").New()
		for i, arg := range info.Arguments {
			argument := js.Global().Get("Object").New()
			argument.Set("name", pointers.Get(arg.Name)[0])
			argument.Set("type", uint32(arg.Type))
			argument.Set("hint", uint32(arg.Hint))
			argument.Set("hint_string", pointers.Get(arg.HintString)[0])
			argument.Set("usage", uint32(arg.Usage))
			argument.Set("class_name", pointers.Get(arg.ClassName)[0])
			arguments.Call("push", argument)
			argument_metadatas.Call("push", uint32(info.ArgumentsMetadata[i]))
		}
		converted.Set("arguments_info", arguments)
		converted.Set("arguments_metadata", argument_metadatas)
		converted.Set("default_argument_count", uint32(len(info.DefaultArguments)))
		for i, arg := range info.DefaultArguments {
			u64 := pointers.Get(arg)
			u32 := *(*[6]uint32)(unsafe.Pointer(&u64))
			for j, v := range u32 {
				write_params_buffer.Invoke(0, i, j, v)
			}
		}
		classdb_register_extension_class_method.Invoke(uint32(library), pointers.Get(class)[0], converted)
	}
	var scratch [16 * 16]uint32
	memory_index := dlsym("mem_index")
	API.Memory.Index = func(frame gd.Address, index int, size uintptr) unsafe.Pointer {
		if index < 0 {
			return unsafe.Pointer(&scratch)
		}
		memory_index.Invoke(uint32(frame), uint32(index), uint32(size))
		for i := uintptr(0); i < size/4; i++ {
			scratch[i] = uint32(read_result_buffer.Invoke(0, 0, i).Int())
		}
		return unsafe.Pointer(&scratch)
	}
	memory_write := dlsym("mem_write")
	API.Memory.Write = func(frame gd.Address, ptr unsafe.Pointer, size uintptr) {
		for i := uintptr(0); i < (size/4)+1; i++ {
			write_params_buffer.Invoke(0, 0, i, scratch[i])
		}
		memory_write.Invoke(uint32(frame), uint32(size))
	}

	dictionary_operator_index_set := dlsym("dictionary_operator_index_set")
	API.Dictionary.SetIndex = func(dict gd.Dictionary, k, v gd.Variant) {
		raw_key := pointers.Get(k)
		key := *(*[6]uint32)(unsafe.Pointer(&raw_key))
		raw_val := pointers.Get(v)
		val := *(*[6]uint32)(unsafe.Pointer(&raw_val))
		dictionary_operator_index_set.Invoke(pointers.Get(dict)[0], key[0], key[1], key[2], key[3], key[4], key[5], val[0], val[1], val[2], val[3], val[4], val[5])
	}
	dictionary_operator_index := dlsym("dictionary_operator_index")
	API.Dictionary.Index = func(dict gd.Dictionary, k gd.Variant) gd.Variant {
		raw_key := pointers.Get(k)
		key := *(*[6]uint32)(unsafe.Pointer(&raw_key))
		dictionary_operator_index.Invoke(pointers.Get(dict)[0], key[0], key[1], key[2], key[3], key[4], key[5])
		var buf [6]uint32
		for i := range buf {
			buf[i] = uint32(read_result_buffer.Invoke(0, 0, i).Int())
		}
		return pointers.Raw[gd.Variant](*(*[3]uint64)(unsafe.Pointer(&buf))).Copy()
	}
	array_set_typed := dlsym("array_set_typed")
	API.Array.SetTyped = func(self gd.Array, t gd.VariantType, className gd.StringName, script gd.Object) {
		array_set_typed.Invoke(pointers.Get(self)[0], uint32(t), pointers.Get(className)[0], pointers.Get(script)[0])
	}
	array_operator_index_set := dlsym("array_operator_index_set")
	API.Array.SetIndex = func(self gd.Array, index gd.Int, value gd.Variant) {
		raw := pointers.Get(value)
		val := *(*[6]uint32)(unsafe.Pointer(&raw))
		array_operator_index_set.Invoke(pointers.Get(self)[0], uint32(index), val[0], val[1], val[2], val[3], val[4], val[5])
	}
	array_operator_index := dlsym("array_operator_index")
	API.Array.Index = func(self gd.Array, index gd.Int) gd.Variant {
		array_operator_index.Invoke(pointers.Get(self)[0], uint32(index))
		var buf [6]uint32
		for i := range buf {
			buf[i] = uint32(read_result_buffer.Invoke(0, 0, i).Int())
		}
		return pointers.Raw[gd.Variant](*(*[3]uint64)(unsafe.Pointer(&buf))).Copy()
	}
}

func makePackedFunctions[T gd.Packed[T, V], V Packed.Type](prefix string) gd.PackedFunctionsFor[T, V] {
	var API gd.PackedFunctionsFor[T, V]
	packed_int32_array_operator_index := dlsym("packed_" + prefix + "_operator_index")
	API.Index = func(pia T, index gd.Int) V {
		arr := pointers.Get[T, gd.PackedPointers](pia)
		raw := *(*[2]uint32)(unsafe.Pointer(&arr))
		packed_int32_array_operator_index.Invoke(int(raw[0]), int(raw[1]), uint32(index))
		var buf [4]uint32
		for i := range buf {
			buf[i] = uint32(read_result_buffer.Invoke(0, 0, uint32(i)).Int())
		}
		return *(*V)(unsafe.Pointer(&buf))
	}
	packed_int32_array_operator_index_set := dlsym("packed_" + prefix + "_operator_index_set")
	API.SetIndex = func(pia T, index gd.Int, v V) {
		arr := pointers.Get[T, gd.PackedPointers](pia)
		raw := *(*[2]uint32)(unsafe.Pointer(&arr))
		var buf [4]uint32
		*(*V)(unsafe.Pointer(&buf)) = v
		for i, v := range buf {
			write_params_buffer.Invoke(0, 0, i, v)
		}
		packed_int32_array_operator_index_set.Invoke(int(raw[0]), int(raw[1]), uint32(index))
	}
	API.CopyAsSlice = func(pia T) []V {
		var slice = make([]V, pia.Len())
		for i := 0; i < pia.Len(); i++ {
			slice[i] = API.Index(pia, gd.Int(i))
		}
		return slice
	}
	API.CopyFromSlice = func(pia T, slice []V) {
		for i, v := range slice {
			API.SetIndex(pia, gd.Int(i), v)
		}
	}
	return API
}
