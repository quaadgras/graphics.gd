package EditorInspectorPlugin

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
[EditorInspectorPlugin] allows adding custom property editors to [EditorInspector].
When an object is edited, the [method _can_handle] function is called and must return [code]true[/code] if the object type is supported.
If supported, the function [method _parse_begin] will be called, allowing to place custom controls at the beginning of the class.
Subsequently, the [method _parse_category] and [method _parse_property] are called for every category and property. They offer the ability to add custom controls to the inspector too.
Finally, [method _parse_end] will be called.
On each of these calls, the "add" functions can be called.
To use [EditorInspectorPlugin], register it using the [method EditorPlugin.add_inspector_plugin] method first.
	// EditorInspectorPlugin methods that can be overridden by a [Class] that extends it.
	type EditorInspectorPlugin interface {
		//Returns [code]true[/code] if this object can be handled by this plugin.
		CanHandle(obj gd.Object) bool
		//Called to allow adding controls at the beginning of the property list for [param object].
		ParseBegin(obj gd.Object) 
		//Called to allow adding controls at the beginning of a category in the property list for [param object].
		ParseCategory(obj gd.Object, category string) 
		//Called to allow adding controls at the beginning of a group or a sub-group in the property list for [param object].
		ParseGroup(obj gd.Object, group string) 
		//Called to allow adding property-specific editors to the property list for [param object]. The added editor control must extend [EditorProperty]. Returning [code]true[/code] removes the built-in editor for this property, otherwise allows to insert a custom editor before the built-in one.
		ParseProperty(obj gd.Object, atype gd.VariantType, name string, hint_type gd.PropertyHint, hint_string string, usage_flags gd.PropertyUsageFlags, wide bool) bool
		//Called to allow adding controls at the end of the property list for [param object].
		ParseEnd(obj gd.Object) 
	}

*/
type Go [1]classdb.EditorInspectorPlugin

/*
Returns [code]true[/code] if this object can be handled by this plugin.
*/
func (Go) _can_handle(impl func(ptr unsafe.Pointer, obj gd.Object) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = discreet.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})
		defer discreet.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called to allow adding controls at the beginning of the property list for [param object].
*/
func (Go) _parse_begin(impl func(ptr unsafe.Pointer, obj gd.Object) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = discreet.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})
		defer discreet.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, obj)
	}
}

/*
Called to allow adding controls at the beginning of a category in the property list for [param object].
*/
func (Go) _parse_category(impl func(ptr unsafe.Pointer, obj gd.Object, category string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = discreet.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})
		defer discreet.End(obj)
		var category = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,1))
		defer discreet.End(category)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, obj, category.String())
	}
}

/*
Called to allow adding controls at the beginning of a group or a sub-group in the property list for [param object].
*/
func (Go) _parse_group(impl func(ptr unsafe.Pointer, obj gd.Object, group string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = discreet.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})
		defer discreet.End(obj)
		var group = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,1))
		defer discreet.End(group)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, obj, group.String())
	}
}

/*
Called to allow adding property-specific editors to the property list for [param object]. The added editor control must extend [EditorProperty]. Returning [code]true[/code] removes the built-in editor for this property, otherwise allows to insert a custom editor before the built-in one.
*/
func (Go) _parse_property(impl func(ptr unsafe.Pointer, obj gd.Object, atype gd.VariantType, name string, hint_type gd.PropertyHint, hint_string string, usage_flags gd.PropertyUsageFlags, wide bool) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = discreet.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})
		defer discreet.End(obj)
		var atype = gd.UnsafeGet[gd.VariantType](p_args,1)
		var name = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,2))
		defer discreet.End(name)
		var hint_type = gd.UnsafeGet[gd.PropertyHint](p_args,3)
		var hint_string = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,4))
		defer discreet.End(hint_string)
		var usage_flags = gd.UnsafeGet[gd.PropertyUsageFlags](p_args,5)
		var wide = gd.UnsafeGet[bool](p_args,6)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, atype, name.String(), hint_type, hint_string.String(), usage_flags, wide)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called to allow adding controls at the end of the property list for [param object].
*/
func (Go) _parse_end(impl func(ptr unsafe.Pointer, obj gd.Object) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = discreet.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})
		defer discreet.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, obj)
	}
}

/*
Adds a custom control, which is not necessarily a property editor.
*/
func (self Go) AddCustomControl(control gdclass.Control) {
	class(self).AddCustomControl(control)
}

/*
Adds a property editor for an individual property. The [param editor] control must extend [EditorProperty].
There can be multiple property editors for a property. If [param add_to_end] is [code]true[/code], this newly added editor will be displayed after all the other editors of the property whose [param add_to_end] is [code]false[/code]. For example, the editor uses this parameter to add an "Edit Region" button for [member Sprite2D.region_rect] below the regular [Rect2] editor.
[param label] can be used to choose a custom label for the property editor in the inspector. If left empty, the label is computed from the name of the property instead.
*/
func (self Go) AddPropertyEditor(property string, editor gdclass.Control) {
	class(self).AddPropertyEditor(gd.NewString(property), editor, false, gd.NewString(""))
}

/*
Adds an editor that allows modifying multiple properties. The [param editor] control must extend [EditorProperty].
*/
func (self Go) AddPropertyEditorForMultipleProperties(label string, properties []string, editor gdclass.Control) {
	class(self).AddPropertyEditorForMultipleProperties(gd.NewString(label), gd.NewPackedStringSlice(properties), editor)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorInspectorPlugin
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorInspectorPlugin"))
	return Go{classdb.EditorInspectorPlugin(object)}
}

/*
Returns [code]true[/code] if this object can be handled by this plugin.
*/
func (class) _can_handle(impl func(ptr unsafe.Pointer, obj gd.Object) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = discreet.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})
		defer discreet.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called to allow adding controls at the beginning of the property list for [param object].
*/
func (class) _parse_begin(impl func(ptr unsafe.Pointer, obj gd.Object) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = discreet.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})
		defer discreet.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, obj)
	}
}

/*
Called to allow adding controls at the beginning of a category in the property list for [param object].
*/
func (class) _parse_category(impl func(ptr unsafe.Pointer, obj gd.Object, category gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = discreet.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})
		defer discreet.End(obj)
		var category = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, obj, category)
	}
}

/*
Called to allow adding controls at the beginning of a group or a sub-group in the property list for [param object].
*/
func (class) _parse_group(impl func(ptr unsafe.Pointer, obj gd.Object, group gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = discreet.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})
		defer discreet.End(obj)
		var group = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, obj, group)
	}
}

/*
Called to allow adding property-specific editors to the property list for [param object]. The added editor control must extend [EditorProperty]. Returning [code]true[/code] removes the built-in editor for this property, otherwise allows to insert a custom editor before the built-in one.
*/
func (class) _parse_property(impl func(ptr unsafe.Pointer, obj gd.Object, atype gd.VariantType, name gd.String, hint_type gd.PropertyHint, hint_string gd.String, usage_flags gd.PropertyUsageFlags, wide bool) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = discreet.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})
		defer discreet.End(obj)
		var atype = gd.UnsafeGet[gd.VariantType](p_args,1)
		var name = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,2))
		var hint_type = gd.UnsafeGet[gd.PropertyHint](p_args,3)
		var hint_string = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,4))
		var usage_flags = gd.UnsafeGet[gd.PropertyUsageFlags](p_args,5)
		var wide = gd.UnsafeGet[bool](p_args,6)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, atype, name, hint_type, hint_string, usage_flags, wide)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called to allow adding controls at the end of the property list for [param object].
*/
func (class) _parse_end(impl func(ptr unsafe.Pointer, obj gd.Object) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = discreet.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})
		defer discreet.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, obj)
	}
}

/*
Adds a custom control, which is not necessarily a property editor.
*/
//go:nosplit
func (self class) AddCustomControl(control gdclass.Control)  {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(gd.Object(control[0])))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInspectorPlugin.Bind_add_custom_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a property editor for an individual property. The [param editor] control must extend [EditorProperty].
There can be multiple property editors for a property. If [param add_to_end] is [code]true[/code], this newly added editor will be displayed after all the other editors of the property whose [param add_to_end] is [code]false[/code]. For example, the editor uses this parameter to add an "Edit Region" button for [member Sprite2D.region_rect] below the regular [Rect2] editor.
[param label] can be used to choose a custom label for the property editor in the inspector. If left empty, the label is computed from the name of the property instead.
*/
//go:nosplit
func (self class) AddPropertyEditor(property gd.String, editor gdclass.Control, add_to_end bool, label gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(property))
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(gd.Object(editor[0])))
	callframe.Arg(frame, add_to_end)
	callframe.Arg(frame, discreet.Get(label))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInspectorPlugin.Bind_add_property_editor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds an editor that allows modifying multiple properties. The [param editor] control must extend [EditorProperty].
*/
//go:nosplit
func (self class) AddPropertyEditorForMultipleProperties(label gd.String, properties gd.PackedStringArray, editor gdclass.Control)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(label))
	callframe.Arg(frame, discreet.Get(properties))
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(gd.Object(editor[0])))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInspectorPlugin.Bind_add_property_editor_for_multiple_properties, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsEditorInspectorPlugin() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorInspectorPlugin() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_can_handle": return reflect.ValueOf(self._can_handle);
	case "_parse_begin": return reflect.ValueOf(self._parse_begin);
	case "_parse_category": return reflect.ValueOf(self._parse_category);
	case "_parse_group": return reflect.ValueOf(self._parse_group);
	case "_parse_property": return reflect.ValueOf(self._parse_property);
	case "_parse_end": return reflect.ValueOf(self._parse_end);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_can_handle": return reflect.ValueOf(self._can_handle);
	case "_parse_begin": return reflect.ValueOf(self._parse_begin);
	case "_parse_category": return reflect.ValueOf(self._parse_category);
	case "_parse_group": return reflect.ValueOf(self._parse_group);
	case "_parse_property": return reflect.ValueOf(self._parse_property);
	case "_parse_end": return reflect.ValueOf(self._parse_end);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("EditorInspectorPlugin", func(ptr gd.Object) any { return classdb.EditorInspectorPlugin(ptr) })}
