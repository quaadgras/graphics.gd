package ScriptLanguageExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/ScriptLanguage"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

type Go [1]classdb.ScriptLanguageExtension
func (Go) _get_name(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Go) _init(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (Go) _get_type(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Go) _get_extension(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Go) _finish(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (Go) _get_reserved_words(impl func(ptr unsafe.Pointer) []string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.PackedStringSlice(ret)))
		gc.End()
	}
}
func (Go) _is_control_flow_keyword(impl func(ptr unsafe.Pointer, keyword string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var keyword = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, keyword.String())
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _get_comment_delimiters(impl func(ptr unsafe.Pointer) []string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.PackedStringSlice(ret)))
		gc.End()
	}
}
func (Go) _get_doc_comment_delimiters(impl func(ptr unsafe.Pointer) []string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.PackedStringSlice(ret)))
		gc.End()
	}
}
func (Go) _get_string_delimiters(impl func(ptr unsafe.Pointer) []string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.PackedStringSlice(ret)))
		gc.End()
	}
}
func (Go) _make_template(impl func(ptr unsafe.Pointer, template string, class_name string, base_class_name string) gdclass.Script, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var template = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var class_name = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		var base_class_name = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, template.String(), class_name.String(), base_class_name.String())
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		gc.End()
	}
}
func (Go) _get_built_in_templates(impl func(ptr unsafe.Pointer, obj string) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var obj = mmm.Let[gd.StringName](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj.String())
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		gc.End()
	}
}
func (Go) _is_using_templates(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _validate(impl func(ptr unsafe.Pointer, script string, path string, validate_functions bool, validate_errors bool, validate_warnings bool, validate_safe_lines bool) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var script = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var path = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		var validate_functions = gd.UnsafeGet[bool](p_args,2)
		var validate_errors = gd.UnsafeGet[bool](p_args,3)
		var validate_warnings = gd.UnsafeGet[bool](p_args,4)
		var validate_safe_lines = gd.UnsafeGet[bool](p_args,5)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, script.String(), path.String(), validate_functions, validate_errors, validate_warnings, validate_safe_lines)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Go) _validate_path(impl func(ptr unsafe.Pointer, path string) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var path = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String())
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Go) _create_script(impl func(ptr unsafe.Pointer) gd.Object, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.AsPointer()))
		gc.End()
	}
}
func (Go) _has_named_classes(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _supports_builtin_mode(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _supports_documentation(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _can_inherit_from_file(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Returns the line where the function is defined in the code, or [code]-1[/code] if the function is not present.
*/
func (Go) _find_function(impl func(ptr unsafe.Pointer, function string, code string) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var function = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var code = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, function.String(), code.String())
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Go) _make_function(impl func(ptr unsafe.Pointer, class_name string, function_name string, function_args []string) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var class_name = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var function_name = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		var function_args = mmm.Let[gd.PackedStringArray](gc.Lifetime, gc.API, gd.UnsafeGet[[2]uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, class_name.String(), function_name.String(), function_args.Strings(gc))
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Go) _can_make_function(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _open_in_external_editor(impl func(ptr unsafe.Pointer, script gdclass.Script, line int, column int) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var script gdclass.Script
		script[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var line = gd.UnsafeGet[gd.Int](p_args,1)
		var column = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, script, int(line), int(column))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _overrides_external_editor(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _preferred_file_name_casing(impl func(ptr unsafe.Pointer) classdb.ScriptLanguageScriptNameCasing, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _complete_code(impl func(ptr unsafe.Pointer, code string, path string, owner gd.Object) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var code = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var path = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		var owner gd.Object
		owner.SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,2)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, code.String(), path.String(), owner)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Go) _lookup_code(impl func(ptr unsafe.Pointer, code string, symbol string, path string, owner gd.Object) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var code = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var symbol = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		var path = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,2))
		var owner gd.Object
		owner.SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,3)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, code.String(), symbol.String(), path.String(), owner)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Go) _auto_indent_code(impl func(ptr unsafe.Pointer, code string, from_line int, to_line int) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var code = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var from_line = gd.UnsafeGet[gd.Int](p_args,1)
		var to_line = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, code.String(), int(from_line), int(to_line))
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Go) _add_global_constant(impl func(ptr unsafe.Pointer, name string, value gd.Variant) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var name = mmm.Let[gd.StringName](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var value = mmm.Let[gd.Variant](gc.Lifetime, gc.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, name.String(), value)
		gc.End()
	}
}
func (Go) _add_named_global_constant(impl func(ptr unsafe.Pointer, name string, value gd.Variant) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var name = mmm.Let[gd.StringName](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var value = mmm.Let[gd.Variant](gc.Lifetime, gc.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, name.String(), value)
		gc.End()
	}
}
func (Go) _remove_named_global_constant(impl func(ptr unsafe.Pointer, name string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var name = mmm.Let[gd.StringName](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, name.String())
		gc.End()
	}
}
func (Go) _thread_enter(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (Go) _thread_exit(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (Go) _debug_get_error(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Go) _debug_get_stack_level_count(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Go) _debug_get_stack_level_line(impl func(ptr unsafe.Pointer, level int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var level = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(level))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Go) _debug_get_stack_level_function(impl func(ptr unsafe.Pointer, level int) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var level = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(level))
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}

/*
Returns the source associated with a given debug stack position.
*/
func (Go) _debug_get_stack_level_source(impl func(ptr unsafe.Pointer, level int) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var level = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(level))
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Go) _debug_get_stack_level_locals(impl func(ptr unsafe.Pointer, level int, max_subitems int, max_depth int) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var level = gd.UnsafeGet[gd.Int](p_args,0)
		var max_subitems = gd.UnsafeGet[gd.Int](p_args,1)
		var max_depth = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(level), int(max_subitems), int(max_depth))
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Go) _debug_get_stack_level_members(impl func(ptr unsafe.Pointer, level int, max_subitems int, max_depth int) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var level = gd.UnsafeGet[gd.Int](p_args,0)
		var max_subitems = gd.UnsafeGet[gd.Int](p_args,1)
		var max_depth = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(level), int(max_subitems), int(max_depth))
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Go) _debug_get_stack_level_instance(impl func(ptr unsafe.Pointer, level int) unsafe.Pointer, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var level = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(level))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _debug_get_globals(impl func(ptr unsafe.Pointer, max_subitems int, max_depth int) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var max_subitems = gd.UnsafeGet[gd.Int](p_args,0)
		var max_depth = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(max_subitems), int(max_depth))
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Go) _debug_parse_stack_level_expression(impl func(ptr unsafe.Pointer, level int, expression string, max_subitems int, max_depth int) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var level = gd.UnsafeGet[gd.Int](p_args,0)
		var expression = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		var max_subitems = gd.UnsafeGet[gd.Int](p_args,2)
		var max_depth = gd.UnsafeGet[gd.Int](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(level), expression.String(), int(max_subitems), int(max_depth))
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Go) _debug_get_current_stack_info(impl func(ptr unsafe.Pointer) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		gc.End()
	}
}
func (Go) _reload_all_scripts(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (Go) _reload_tool_script(impl func(ptr unsafe.Pointer, script gdclass.Script, soft_reload bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var script gdclass.Script
		script[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var soft_reload = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, script, soft_reload)
		gc.End()
	}
}
func (Go) _get_recognized_extensions(impl func(ptr unsafe.Pointer) []string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.PackedStringSlice(ret)))
		gc.End()
	}
}
func (Go) _get_public_functions(impl func(ptr unsafe.Pointer) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		gc.End()
	}
}
func (Go) _get_public_constants(impl func(ptr unsafe.Pointer) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Go) _get_public_annotations(impl func(ptr unsafe.Pointer) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		gc.End()
	}
}
func (Go) _profiling_start(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (Go) _profiling_stop(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (Go) _profiling_set_save_native_calls(impl func(ptr unsafe.Pointer, enable bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var enable = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, enable)
		gc.End()
	}
}
func (Go) _profiling_get_accumulated_data(impl func(ptr unsafe.Pointer, info_array *classdb.ScriptLanguageExtensionProfilingInfo, info_max int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var info_array = gd.UnsafeGet[*classdb.ScriptLanguageExtensionProfilingInfo](p_args,0)
		var info_max = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, info_array, int(info_max))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Go) _profiling_get_frame_data(impl func(ptr unsafe.Pointer, info_array *classdb.ScriptLanguageExtensionProfilingInfo, info_max int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var info_array = gd.UnsafeGet[*classdb.ScriptLanguageExtensionProfilingInfo](p_args,0)
		var info_max = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, info_array, int(info_max))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Go) _frame(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (Go) _handles_global_class_type(impl func(ptr unsafe.Pointer, atype string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var atype = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, atype.String())
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _get_global_class_name(impl func(ptr unsafe.Pointer, path string) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var path = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String())
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ScriptLanguageExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ScriptLanguageExtension"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (class) _get_name(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _init(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

func (class) _get_type(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _get_extension(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _finish(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

func (class) _get_reserved_words(impl func(ptr unsafe.Pointer) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _is_control_flow_keyword(impl func(ptr unsafe.Pointer, keyword gd.String) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var keyword = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, keyword)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_comment_delimiters(impl func(ptr unsafe.Pointer) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _get_doc_comment_delimiters(impl func(ptr unsafe.Pointer) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _get_string_delimiters(impl func(ptr unsafe.Pointer) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _make_template(impl func(ptr unsafe.Pointer, template gd.String, class_name gd.String, base_class_name gd.String) gdclass.Script, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var template = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var class_name = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var base_class_name = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, template, class_name, base_class_name)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}

func (class) _get_built_in_templates(impl func(ptr unsafe.Pointer, obj gd.StringName) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var obj = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

func (class) _is_using_templates(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _validate(impl func(ptr unsafe.Pointer, script gd.String, path gd.String, validate_functions bool, validate_errors bool, validate_warnings bool, validate_safe_lines bool) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var script = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var validate_functions = gd.UnsafeGet[bool](p_args,2)
		var validate_errors = gd.UnsafeGet[bool](p_args,3)
		var validate_warnings = gd.UnsafeGet[bool](p_args,4)
		var validate_safe_lines = gd.UnsafeGet[bool](p_args,5)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, script, path, validate_functions, validate_errors, validate_warnings, validate_safe_lines)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _validate_path(impl func(ptr unsafe.Pointer, path gd.String) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _create_script(impl func(ptr unsafe.Pointer) gd.Object, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.AsPointer()))
		ctx.End()
	}
}

func (class) _has_named_classes(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _supports_builtin_mode(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _supports_documentation(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _can_inherit_from_file(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Returns the line where the function is defined in the code, or [code]-1[/code] if the function is not present.
*/
func (class) _find_function(impl func(ptr unsafe.Pointer, function gd.String, code gd.String) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var function = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var code = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, function, code)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _make_function(impl func(ptr unsafe.Pointer, class_name gd.String, function_name gd.String, function_args gd.PackedStringArray) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var class_name = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var function_name = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var function_args = mmm.Let[gd.PackedStringArray](ctx.Lifetime, ctx.API, gd.UnsafeGet[[2]uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, class_name, function_name, function_args)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _can_make_function(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _open_in_external_editor(impl func(ptr unsafe.Pointer, script gdclass.Script, line gd.Int, column gd.Int) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var script gdclass.Script
		script[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var line = gd.UnsafeGet[gd.Int](p_args,1)
		var column = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, script, line, column)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _overrides_external_editor(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _preferred_file_name_casing(impl func(ptr unsafe.Pointer) classdb.ScriptLanguageScriptNameCasing, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _complete_code(impl func(ptr unsafe.Pointer, code gd.String, path gd.String, owner gd.Object) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var code = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var owner gd.Object
		owner.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,2)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, code, path, owner)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _lookup_code(impl func(ptr unsafe.Pointer, code gd.String, symbol gd.String, path gd.String, owner gd.Object) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var code = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var symbol = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,2))
		var owner gd.Object
		owner.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,3)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, code, symbol, path, owner)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _auto_indent_code(impl func(ptr unsafe.Pointer, code gd.String, from_line gd.Int, to_line gd.Int) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var code = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var from_line = gd.UnsafeGet[gd.Int](p_args,1)
		var to_line = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, code, from_line, to_line)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _add_global_constant(impl func(ptr unsafe.Pointer, name gd.StringName, value gd.Variant) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var name = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var value = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, name, value)
		ctx.End()
	}
}

func (class) _add_named_global_constant(impl func(ptr unsafe.Pointer, name gd.StringName, value gd.Variant) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var name = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var value = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, name, value)
		ctx.End()
	}
}

func (class) _remove_named_global_constant(impl func(ptr unsafe.Pointer, name gd.StringName) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var name = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, name)
		ctx.End()
	}
}

func (class) _thread_enter(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

func (class) _thread_exit(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

func (class) _debug_get_error(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _debug_get_stack_level_count(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _debug_get_stack_level_line(impl func(ptr unsafe.Pointer, level gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var level = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, level)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _debug_get_stack_level_function(impl func(ptr unsafe.Pointer, level gd.Int) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var level = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, level)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Returns the source associated with a given debug stack position.
*/
func (class) _debug_get_stack_level_source(impl func(ptr unsafe.Pointer, level gd.Int) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var level = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, level)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _debug_get_stack_level_locals(impl func(ptr unsafe.Pointer, level gd.Int, max_subitems gd.Int, max_depth gd.Int) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var level = gd.UnsafeGet[gd.Int](p_args,0)
		var max_subitems = gd.UnsafeGet[gd.Int](p_args,1)
		var max_depth = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, level, max_subitems, max_depth)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _debug_get_stack_level_members(impl func(ptr unsafe.Pointer, level gd.Int, max_subitems gd.Int, max_depth gd.Int) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var level = gd.UnsafeGet[gd.Int](p_args,0)
		var max_subitems = gd.UnsafeGet[gd.Int](p_args,1)
		var max_depth = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, level, max_subitems, max_depth)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _debug_get_stack_level_instance(impl func(ptr unsafe.Pointer, level gd.Int) unsafe.Pointer, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var level = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, level)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _debug_get_globals(impl func(ptr unsafe.Pointer, max_subitems gd.Int, max_depth gd.Int) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var max_subitems = gd.UnsafeGet[gd.Int](p_args,0)
		var max_depth = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, max_subitems, max_depth)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _debug_parse_stack_level_expression(impl func(ptr unsafe.Pointer, level gd.Int, expression gd.String, max_subitems gd.Int, max_depth gd.Int) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var level = gd.UnsafeGet[gd.Int](p_args,0)
		var expression = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var max_subitems = gd.UnsafeGet[gd.Int](p_args,2)
		var max_depth = gd.UnsafeGet[gd.Int](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, level, expression, max_subitems, max_depth)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _debug_get_current_stack_info(impl func(ptr unsafe.Pointer) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

func (class) _reload_all_scripts(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

func (class) _reload_tool_script(impl func(ptr unsafe.Pointer, script gdclass.Script, soft_reload bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var script gdclass.Script
		script[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var soft_reload = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, script, soft_reload)
		ctx.End()
	}
}

func (class) _get_recognized_extensions(impl func(ptr unsafe.Pointer) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _get_public_functions(impl func(ptr unsafe.Pointer) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

func (class) _get_public_constants(impl func(ptr unsafe.Pointer) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _get_public_annotations(impl func(ptr unsafe.Pointer) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

func (class) _profiling_start(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

func (class) _profiling_stop(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

func (class) _profiling_set_save_native_calls(impl func(ptr unsafe.Pointer, enable bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var enable = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, enable)
		ctx.End()
	}
}

func (class) _profiling_get_accumulated_data(impl func(ptr unsafe.Pointer, info_array *classdb.ScriptLanguageExtensionProfilingInfo, info_max gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var info_array = gd.UnsafeGet[*classdb.ScriptLanguageExtensionProfilingInfo](p_args,0)
		var info_max = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, info_array, info_max)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _profiling_get_frame_data(impl func(ptr unsafe.Pointer, info_array *classdb.ScriptLanguageExtensionProfilingInfo, info_max gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var info_array = gd.UnsafeGet[*classdb.ScriptLanguageExtensionProfilingInfo](p_args,0)
		var info_max = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, info_array, info_max)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _frame(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

func (class) _handles_global_class_type(impl func(ptr unsafe.Pointer, atype gd.String) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var atype = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, atype)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_global_class_name(impl func(ptr unsafe.Pointer, path gd.String) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (self class) AsScriptLanguageExtension() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsScriptLanguageExtension() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsScriptLanguage() ScriptLanguage.GD { return *((*ScriptLanguage.GD)(unsafe.Pointer(&self))) }
func (self Go) AsScriptLanguage() ScriptLanguage.Go { return *((*ScriptLanguage.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_name": return reflect.ValueOf(self._get_name);
	case "_init": return reflect.ValueOf(self._init);
	case "_get_type": return reflect.ValueOf(self._get_type);
	case "_get_extension": return reflect.ValueOf(self._get_extension);
	case "_finish": return reflect.ValueOf(self._finish);
	case "_get_reserved_words": return reflect.ValueOf(self._get_reserved_words);
	case "_is_control_flow_keyword": return reflect.ValueOf(self._is_control_flow_keyword);
	case "_get_comment_delimiters": return reflect.ValueOf(self._get_comment_delimiters);
	case "_get_doc_comment_delimiters": return reflect.ValueOf(self._get_doc_comment_delimiters);
	case "_get_string_delimiters": return reflect.ValueOf(self._get_string_delimiters);
	case "_make_template": return reflect.ValueOf(self._make_template);
	case "_get_built_in_templates": return reflect.ValueOf(self._get_built_in_templates);
	case "_is_using_templates": return reflect.ValueOf(self._is_using_templates);
	case "_validate": return reflect.ValueOf(self._validate);
	case "_validate_path": return reflect.ValueOf(self._validate_path);
	case "_create_script": return reflect.ValueOf(self._create_script);
	case "_has_named_classes": return reflect.ValueOf(self._has_named_classes);
	case "_supports_builtin_mode": return reflect.ValueOf(self._supports_builtin_mode);
	case "_supports_documentation": return reflect.ValueOf(self._supports_documentation);
	case "_can_inherit_from_file": return reflect.ValueOf(self._can_inherit_from_file);
	case "_find_function": return reflect.ValueOf(self._find_function);
	case "_make_function": return reflect.ValueOf(self._make_function);
	case "_can_make_function": return reflect.ValueOf(self._can_make_function);
	case "_open_in_external_editor": return reflect.ValueOf(self._open_in_external_editor);
	case "_overrides_external_editor": return reflect.ValueOf(self._overrides_external_editor);
	case "_preferred_file_name_casing": return reflect.ValueOf(self._preferred_file_name_casing);
	case "_complete_code": return reflect.ValueOf(self._complete_code);
	case "_lookup_code": return reflect.ValueOf(self._lookup_code);
	case "_auto_indent_code": return reflect.ValueOf(self._auto_indent_code);
	case "_add_global_constant": return reflect.ValueOf(self._add_global_constant);
	case "_add_named_global_constant": return reflect.ValueOf(self._add_named_global_constant);
	case "_remove_named_global_constant": return reflect.ValueOf(self._remove_named_global_constant);
	case "_thread_enter": return reflect.ValueOf(self._thread_enter);
	case "_thread_exit": return reflect.ValueOf(self._thread_exit);
	case "_debug_get_error": return reflect.ValueOf(self._debug_get_error);
	case "_debug_get_stack_level_count": return reflect.ValueOf(self._debug_get_stack_level_count);
	case "_debug_get_stack_level_line": return reflect.ValueOf(self._debug_get_stack_level_line);
	case "_debug_get_stack_level_function": return reflect.ValueOf(self._debug_get_stack_level_function);
	case "_debug_get_stack_level_source": return reflect.ValueOf(self._debug_get_stack_level_source);
	case "_debug_get_stack_level_locals": return reflect.ValueOf(self._debug_get_stack_level_locals);
	case "_debug_get_stack_level_members": return reflect.ValueOf(self._debug_get_stack_level_members);
	case "_debug_get_stack_level_instance": return reflect.ValueOf(self._debug_get_stack_level_instance);
	case "_debug_get_globals": return reflect.ValueOf(self._debug_get_globals);
	case "_debug_parse_stack_level_expression": return reflect.ValueOf(self._debug_parse_stack_level_expression);
	case "_debug_get_current_stack_info": return reflect.ValueOf(self._debug_get_current_stack_info);
	case "_reload_all_scripts": return reflect.ValueOf(self._reload_all_scripts);
	case "_reload_tool_script": return reflect.ValueOf(self._reload_tool_script);
	case "_get_recognized_extensions": return reflect.ValueOf(self._get_recognized_extensions);
	case "_get_public_functions": return reflect.ValueOf(self._get_public_functions);
	case "_get_public_constants": return reflect.ValueOf(self._get_public_constants);
	case "_get_public_annotations": return reflect.ValueOf(self._get_public_annotations);
	case "_profiling_start": return reflect.ValueOf(self._profiling_start);
	case "_profiling_stop": return reflect.ValueOf(self._profiling_stop);
	case "_profiling_set_save_native_calls": return reflect.ValueOf(self._profiling_set_save_native_calls);
	case "_profiling_get_accumulated_data": return reflect.ValueOf(self._profiling_get_accumulated_data);
	case "_profiling_get_frame_data": return reflect.ValueOf(self._profiling_get_frame_data);
	case "_frame": return reflect.ValueOf(self._frame);
	case "_handles_global_class_type": return reflect.ValueOf(self._handles_global_class_type);
	case "_get_global_class_name": return reflect.ValueOf(self._get_global_class_name);
	default: return gd.VirtualByName(self.AsScriptLanguage(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_get_name": return reflect.ValueOf(self._get_name);
	case "_init": return reflect.ValueOf(self._init);
	case "_get_type": return reflect.ValueOf(self._get_type);
	case "_get_extension": return reflect.ValueOf(self._get_extension);
	case "_finish": return reflect.ValueOf(self._finish);
	case "_get_reserved_words": return reflect.ValueOf(self._get_reserved_words);
	case "_is_control_flow_keyword": return reflect.ValueOf(self._is_control_flow_keyword);
	case "_get_comment_delimiters": return reflect.ValueOf(self._get_comment_delimiters);
	case "_get_doc_comment_delimiters": return reflect.ValueOf(self._get_doc_comment_delimiters);
	case "_get_string_delimiters": return reflect.ValueOf(self._get_string_delimiters);
	case "_make_template": return reflect.ValueOf(self._make_template);
	case "_get_built_in_templates": return reflect.ValueOf(self._get_built_in_templates);
	case "_is_using_templates": return reflect.ValueOf(self._is_using_templates);
	case "_validate": return reflect.ValueOf(self._validate);
	case "_validate_path": return reflect.ValueOf(self._validate_path);
	case "_create_script": return reflect.ValueOf(self._create_script);
	case "_has_named_classes": return reflect.ValueOf(self._has_named_classes);
	case "_supports_builtin_mode": return reflect.ValueOf(self._supports_builtin_mode);
	case "_supports_documentation": return reflect.ValueOf(self._supports_documentation);
	case "_can_inherit_from_file": return reflect.ValueOf(self._can_inherit_from_file);
	case "_find_function": return reflect.ValueOf(self._find_function);
	case "_make_function": return reflect.ValueOf(self._make_function);
	case "_can_make_function": return reflect.ValueOf(self._can_make_function);
	case "_open_in_external_editor": return reflect.ValueOf(self._open_in_external_editor);
	case "_overrides_external_editor": return reflect.ValueOf(self._overrides_external_editor);
	case "_preferred_file_name_casing": return reflect.ValueOf(self._preferred_file_name_casing);
	case "_complete_code": return reflect.ValueOf(self._complete_code);
	case "_lookup_code": return reflect.ValueOf(self._lookup_code);
	case "_auto_indent_code": return reflect.ValueOf(self._auto_indent_code);
	case "_add_global_constant": return reflect.ValueOf(self._add_global_constant);
	case "_add_named_global_constant": return reflect.ValueOf(self._add_named_global_constant);
	case "_remove_named_global_constant": return reflect.ValueOf(self._remove_named_global_constant);
	case "_thread_enter": return reflect.ValueOf(self._thread_enter);
	case "_thread_exit": return reflect.ValueOf(self._thread_exit);
	case "_debug_get_error": return reflect.ValueOf(self._debug_get_error);
	case "_debug_get_stack_level_count": return reflect.ValueOf(self._debug_get_stack_level_count);
	case "_debug_get_stack_level_line": return reflect.ValueOf(self._debug_get_stack_level_line);
	case "_debug_get_stack_level_function": return reflect.ValueOf(self._debug_get_stack_level_function);
	case "_debug_get_stack_level_source": return reflect.ValueOf(self._debug_get_stack_level_source);
	case "_debug_get_stack_level_locals": return reflect.ValueOf(self._debug_get_stack_level_locals);
	case "_debug_get_stack_level_members": return reflect.ValueOf(self._debug_get_stack_level_members);
	case "_debug_get_stack_level_instance": return reflect.ValueOf(self._debug_get_stack_level_instance);
	case "_debug_get_globals": return reflect.ValueOf(self._debug_get_globals);
	case "_debug_parse_stack_level_expression": return reflect.ValueOf(self._debug_parse_stack_level_expression);
	case "_debug_get_current_stack_info": return reflect.ValueOf(self._debug_get_current_stack_info);
	case "_reload_all_scripts": return reflect.ValueOf(self._reload_all_scripts);
	case "_reload_tool_script": return reflect.ValueOf(self._reload_tool_script);
	case "_get_recognized_extensions": return reflect.ValueOf(self._get_recognized_extensions);
	case "_get_public_functions": return reflect.ValueOf(self._get_public_functions);
	case "_get_public_constants": return reflect.ValueOf(self._get_public_constants);
	case "_get_public_annotations": return reflect.ValueOf(self._get_public_annotations);
	case "_profiling_start": return reflect.ValueOf(self._profiling_start);
	case "_profiling_stop": return reflect.ValueOf(self._profiling_stop);
	case "_profiling_set_save_native_calls": return reflect.ValueOf(self._profiling_set_save_native_calls);
	case "_profiling_get_accumulated_data": return reflect.ValueOf(self._profiling_get_accumulated_data);
	case "_profiling_get_frame_data": return reflect.ValueOf(self._profiling_get_frame_data);
	case "_frame": return reflect.ValueOf(self._frame);
	case "_handles_global_class_type": return reflect.ValueOf(self._handles_global_class_type);
	case "_get_global_class_name": return reflect.ValueOf(self._get_global_class_name);
	default: return gd.VirtualByName(self.AsScriptLanguage(), name)
	}
}
func init() {classdb.Register("ScriptLanguageExtension", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type LookupResultType = classdb.ScriptLanguageExtensionLookupResultType

const (
	LookupResultScriptLocation LookupResultType = 0
	LookupResultClass LookupResultType = 1
	LookupResultClassConstant LookupResultType = 2
	LookupResultClassProperty LookupResultType = 3
	LookupResultClassMethod LookupResultType = 4
	LookupResultClassSignal LookupResultType = 5
	LookupResultClassEnum LookupResultType = 6
	LookupResultClassTbdGlobalscope LookupResultType = 7
	LookupResultClassAnnotation LookupResultType = 8
	LookupResultMax LookupResultType = 9
)
type CodeCompletionLocation = classdb.ScriptLanguageExtensionCodeCompletionLocation

const (
/*The option is local to the location of the code completion query - e.g. a local variable. Subsequent value of location represent options from the outer class, the exact value represent how far they are (in terms of inner classes).*/
	LocationLocal CodeCompletionLocation = 0
/*The option is from the containing class or a parent class, relative to the location of the code completion query. Perform a bitwise OR with the class depth (e.g. [code]0[/code] for the local class, [code]1[/code] for the parent, [code]2[/code] for the grandparent, etc.) to store the depth of an option in the class or a parent class.*/
	LocationParentMask CodeCompletionLocation = 256
/*The option is from user code which is not local and not in a derived class (e.g. Autoload Singletons).*/
	LocationOtherUserCode CodeCompletionLocation = 512
/*The option is from other engine code, not covered by the other enum constants - e.g. built-in classes.*/
	LocationOther CodeCompletionLocation = 1024
)
type CodeCompletionKind = classdb.ScriptLanguageExtensionCodeCompletionKind

const (
	CodeCompletionKindClass CodeCompletionKind = 0
	CodeCompletionKindFunction CodeCompletionKind = 1
	CodeCompletionKindSignal CodeCompletionKind = 2
	CodeCompletionKindVariable CodeCompletionKind = 3
	CodeCompletionKindMember CodeCompletionKind = 4
	CodeCompletionKindEnum CodeCompletionKind = 5
	CodeCompletionKindConstant CodeCompletionKind = 6
	CodeCompletionKindNodePath CodeCompletionKind = 7
	CodeCompletionKindFilePath CodeCompletionKind = 8
	CodeCompletionKindPlainText CodeCompletionKind = 9
	CodeCompletionKindMax CodeCompletionKind = 10
)
