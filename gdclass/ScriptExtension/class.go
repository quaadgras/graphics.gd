package ScriptExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Script"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

type Instance [1]classdb.ScriptExtension

func (Instance) _editor_can_reload_from_file(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _placeholder_erased(impl func(ptr unsafe.Pointer, placeholder unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var placeholder = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, placeholder)
	}
}
func (Instance) _can_instantiate(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_base_script(impl func(ptr unsafe.Pointer) gdclass.Script) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_global_name(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewStringName(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _inherits_script(impl func(ptr unsafe.Pointer, script gdclass.Script) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var script = gdclass.Script{pointers.New[classdb.Script]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(script[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, script)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_instance_base_type(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewStringName(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _instance_create(impl func(ptr unsafe.Pointer, for_object gd.Object) unsafe.Pointer) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var for_object = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(for_object)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, for_object)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _placeholder_instance_create(impl func(ptr unsafe.Pointer, for_object gd.Object) unsafe.Pointer) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var for_object = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(for_object)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, for_object)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _instance_has(impl func(ptr unsafe.Pointer, obj gd.Object) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _has_source_code(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_source_code(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _set_source_code(impl func(ptr unsafe.Pointer, code string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var code = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(code)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, code.String())
	}
}
func (Instance) _reload(impl func(ptr unsafe.Pointer, keep_state bool) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var keep_state = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, keep_state)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_documentation(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_class_icon_path(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _has_method(impl func(ptr unsafe.Pointer, method string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var method = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(method)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, method.String())
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _has_static_method(impl func(ptr unsafe.Pointer, method string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var method = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(method)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, method.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Return the expected argument count for the given [param method], or [code]null[/code] if it can't be determined (which will then fall back to the default behavior).
*/
func (Instance) _get_script_method_argument_count(impl func(ptr unsafe.Pointer, method string) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var method = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(method)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, method.String())
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_method_info(impl func(ptr unsafe.Pointer, method string) gd.Dictionary) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var method = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(method)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, method.String())
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _is_tool(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _is_valid(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns [code]true[/code] if the script is an abstract script. An abstract script does not have a constructor and cannot be instantiated.
*/
func (Instance) _is_abstract(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_language(impl func(ptr unsafe.Pointer) gdclass.ScriptLanguage) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _has_script_signal(impl func(ptr unsafe.Pointer, signal string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var signal = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(signal)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, signal.String())
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_script_signal_list(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _has_property_default_value(impl func(ptr unsafe.Pointer, property string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var property = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(property)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, property.String())
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_property_default_value(impl func(ptr unsafe.Pointer, property string) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var property = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(property)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, property.String())
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _update_exports(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _get_script_method_list(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_script_property_list(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_member_line(impl func(ptr unsafe.Pointer, member string) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var member = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(member)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, member.String())
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _get_constants(impl func(ptr unsafe.Pointer) gd.Dictionary) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_members(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _is_placeholder_fallback_enabled(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_rpc_config(impl func(ptr unsafe.Pointer) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ScriptExtension

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ScriptExtension"))
	return Instance{classdb.ScriptExtension(object)}
}

func (class) _editor_can_reload_from_file(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _placeholder_erased(impl func(ptr unsafe.Pointer, placeholder unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var placeholder = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, placeholder)
	}
}

func (class) _can_instantiate(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_base_script(impl func(ptr unsafe.Pointer) gdclass.Script) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_global_name(impl func(ptr unsafe.Pointer) gd.StringName) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _inherits_script(impl func(ptr unsafe.Pointer, script gdclass.Script) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var script = gdclass.Script{pointers.New[classdb.Script]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(script[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, script)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_instance_base_type(impl func(ptr unsafe.Pointer) gd.StringName) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _instance_create(impl func(ptr unsafe.Pointer, for_object gd.Object) unsafe.Pointer) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var for_object = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(for_object)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, for_object)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _placeholder_instance_create(impl func(ptr unsafe.Pointer, for_object gd.Object) unsafe.Pointer) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var for_object = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(for_object)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, for_object)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _instance_has(impl func(ptr unsafe.Pointer, obj gd.Object) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _has_source_code(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_source_code(impl func(ptr unsafe.Pointer) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _set_source_code(impl func(ptr unsafe.Pointer, code gd.String)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var code = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, code)
	}
}

func (class) _reload(impl func(ptr unsafe.Pointer, keep_state bool) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var keep_state = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, keep_state)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_documentation(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_class_icon_path(impl func(ptr unsafe.Pointer) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _has_method(impl func(ptr unsafe.Pointer, method gd.StringName) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var method = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, method)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _has_static_method(impl func(ptr unsafe.Pointer, method gd.StringName) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var method = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, method)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Return the expected argument count for the given [param method], or [code]null[/code] if it can't be determined (which will then fall back to the default behavior).
*/
func (class) _get_script_method_argument_count(impl func(ptr unsafe.Pointer, method gd.StringName) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var method = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, method)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_method_info(impl func(ptr unsafe.Pointer, method gd.StringName) gd.Dictionary) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var method = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, method)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _is_tool(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _is_valid(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns [code]true[/code] if the script is an abstract script. An abstract script does not have a constructor and cannot be instantiated.
*/
func (class) _is_abstract(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_language(impl func(ptr unsafe.Pointer) gdclass.ScriptLanguage) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _has_script_signal(impl func(ptr unsafe.Pointer, signal gd.StringName) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var signal = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, signal)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_script_signal_list(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _has_property_default_value(impl func(ptr unsafe.Pointer, property gd.StringName) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var property = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, property)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_property_default_value(impl func(ptr unsafe.Pointer, property gd.StringName) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var property = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, property)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _update_exports(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _get_script_method_list(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_script_property_list(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_member_line(impl func(ptr unsafe.Pointer, member gd.StringName) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var member = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, member)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_constants(impl func(ptr unsafe.Pointer) gd.Dictionary) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_members(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _is_placeholder_fallback_enabled(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_rpc_config(impl func(ptr unsafe.Pointer) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (self class) AsScriptExtension() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsScriptExtension() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsScript() Script.Advanced      { return *((*Script.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsScript() Script.Instance   { return *((*Script.Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_editor_can_reload_from_file":
		return reflect.ValueOf(self._editor_can_reload_from_file)
	case "_placeholder_erased":
		return reflect.ValueOf(self._placeholder_erased)
	case "_can_instantiate":
		return reflect.ValueOf(self._can_instantiate)
	case "_get_base_script":
		return reflect.ValueOf(self._get_base_script)
	case "_get_global_name":
		return reflect.ValueOf(self._get_global_name)
	case "_inherits_script":
		return reflect.ValueOf(self._inherits_script)
	case "_get_instance_base_type":
		return reflect.ValueOf(self._get_instance_base_type)
	case "_instance_create":
		return reflect.ValueOf(self._instance_create)
	case "_placeholder_instance_create":
		return reflect.ValueOf(self._placeholder_instance_create)
	case "_instance_has":
		return reflect.ValueOf(self._instance_has)
	case "_has_source_code":
		return reflect.ValueOf(self._has_source_code)
	case "_get_source_code":
		return reflect.ValueOf(self._get_source_code)
	case "_set_source_code":
		return reflect.ValueOf(self._set_source_code)
	case "_reload":
		return reflect.ValueOf(self._reload)
	case "_get_documentation":
		return reflect.ValueOf(self._get_documentation)
	case "_get_class_icon_path":
		return reflect.ValueOf(self._get_class_icon_path)
	case "_has_method":
		return reflect.ValueOf(self._has_method)
	case "_has_static_method":
		return reflect.ValueOf(self._has_static_method)
	case "_get_script_method_argument_count":
		return reflect.ValueOf(self._get_script_method_argument_count)
	case "_get_method_info":
		return reflect.ValueOf(self._get_method_info)
	case "_is_tool":
		return reflect.ValueOf(self._is_tool)
	case "_is_valid":
		return reflect.ValueOf(self._is_valid)
	case "_is_abstract":
		return reflect.ValueOf(self._is_abstract)
	case "_get_language":
		return reflect.ValueOf(self._get_language)
	case "_has_script_signal":
		return reflect.ValueOf(self._has_script_signal)
	case "_get_script_signal_list":
		return reflect.ValueOf(self._get_script_signal_list)
	case "_has_property_default_value":
		return reflect.ValueOf(self._has_property_default_value)
	case "_get_property_default_value":
		return reflect.ValueOf(self._get_property_default_value)
	case "_update_exports":
		return reflect.ValueOf(self._update_exports)
	case "_get_script_method_list":
		return reflect.ValueOf(self._get_script_method_list)
	case "_get_script_property_list":
		return reflect.ValueOf(self._get_script_property_list)
	case "_get_member_line":
		return reflect.ValueOf(self._get_member_line)
	case "_get_constants":
		return reflect.ValueOf(self._get_constants)
	case "_get_members":
		return reflect.ValueOf(self._get_members)
	case "_is_placeholder_fallback_enabled":
		return reflect.ValueOf(self._is_placeholder_fallback_enabled)
	case "_get_rpc_config":
		return reflect.ValueOf(self._get_rpc_config)
	default:
		return gd.VirtualByName(self.AsScript(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_editor_can_reload_from_file":
		return reflect.ValueOf(self._editor_can_reload_from_file)
	case "_placeholder_erased":
		return reflect.ValueOf(self._placeholder_erased)
	case "_can_instantiate":
		return reflect.ValueOf(self._can_instantiate)
	case "_get_base_script":
		return reflect.ValueOf(self._get_base_script)
	case "_get_global_name":
		return reflect.ValueOf(self._get_global_name)
	case "_inherits_script":
		return reflect.ValueOf(self._inherits_script)
	case "_get_instance_base_type":
		return reflect.ValueOf(self._get_instance_base_type)
	case "_instance_create":
		return reflect.ValueOf(self._instance_create)
	case "_placeholder_instance_create":
		return reflect.ValueOf(self._placeholder_instance_create)
	case "_instance_has":
		return reflect.ValueOf(self._instance_has)
	case "_has_source_code":
		return reflect.ValueOf(self._has_source_code)
	case "_get_source_code":
		return reflect.ValueOf(self._get_source_code)
	case "_set_source_code":
		return reflect.ValueOf(self._set_source_code)
	case "_reload":
		return reflect.ValueOf(self._reload)
	case "_get_documentation":
		return reflect.ValueOf(self._get_documentation)
	case "_get_class_icon_path":
		return reflect.ValueOf(self._get_class_icon_path)
	case "_has_method":
		return reflect.ValueOf(self._has_method)
	case "_has_static_method":
		return reflect.ValueOf(self._has_static_method)
	case "_get_script_method_argument_count":
		return reflect.ValueOf(self._get_script_method_argument_count)
	case "_get_method_info":
		return reflect.ValueOf(self._get_method_info)
	case "_is_tool":
		return reflect.ValueOf(self._is_tool)
	case "_is_valid":
		return reflect.ValueOf(self._is_valid)
	case "_is_abstract":
		return reflect.ValueOf(self._is_abstract)
	case "_get_language":
		return reflect.ValueOf(self._get_language)
	case "_has_script_signal":
		return reflect.ValueOf(self._has_script_signal)
	case "_get_script_signal_list":
		return reflect.ValueOf(self._get_script_signal_list)
	case "_has_property_default_value":
		return reflect.ValueOf(self._has_property_default_value)
	case "_get_property_default_value":
		return reflect.ValueOf(self._get_property_default_value)
	case "_update_exports":
		return reflect.ValueOf(self._update_exports)
	case "_get_script_method_list":
		return reflect.ValueOf(self._get_script_method_list)
	case "_get_script_property_list":
		return reflect.ValueOf(self._get_script_property_list)
	case "_get_member_line":
		return reflect.ValueOf(self._get_member_line)
	case "_get_constants":
		return reflect.ValueOf(self._get_constants)
	case "_get_members":
		return reflect.ValueOf(self._get_members)
	case "_is_placeholder_fallback_enabled":
		return reflect.ValueOf(self._is_placeholder_fallback_enabled)
	case "_get_rpc_config":
		return reflect.ValueOf(self._get_rpc_config)
	default:
		return gd.VirtualByName(self.AsScript(), name)
	}
}
func init() {
	classdb.Register("ScriptExtension", func(ptr gd.Object) any { return classdb.ScriptExtension(ptr) })
}
