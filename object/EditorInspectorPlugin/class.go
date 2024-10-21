package EditorInspectorPlugin

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
		ParseCategory(obj gd.Object, category gd.String) 
		//Called to allow adding controls at the beginning of a group or a sub-group in the property list for [param object].
		ParseGroup(obj gd.Object, group gd.String) 
		//Called to allow adding property-specific editors to the property list for [param object]. The added editor control must extend [EditorProperty]. Returning [code]true[/code] removes the built-in editor for this property, otherwise allows to insert a custom editor before the built-in one.
		ParseProperty(obj gd.Object, atype gd.VariantType, name gd.String, hint_type gd.PropertyHint, hint_string gd.String, usage_flags gd.PropertyUsageFlags, wide bool) bool
		//Called to allow adding controls at the end of the property list for [param object].
		ParseEnd(obj gd.Object) 
	}

*/
type Simple [1]classdb.EditorInspectorPlugin
func (Simple) _can_handle(impl func(ptr unsafe.Pointer, obj gd.Object) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var obj gd.Object
		obj.SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _parse_begin(impl func(ptr unsafe.Pointer, obj gd.Object) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var obj gd.Object
		obj.SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, obj)
		gc.End()
	}
}
func (Simple) _parse_category(impl func(ptr unsafe.Pointer, obj gd.Object, category string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var obj gd.Object
		obj.SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var category = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, obj, category.String())
		gc.End()
	}
}
func (Simple) _parse_group(impl func(ptr unsafe.Pointer, obj gd.Object, group string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var obj gd.Object
		obj.SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var group = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, obj, group.String())
		gc.End()
	}
}
func (Simple) _parse_property(impl func(ptr unsafe.Pointer, obj gd.Object, atype gd.VariantType, name string, hint_type gd.PropertyHint, hint_string string, usage_flags gd.PropertyUsageFlags, wide bool) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var obj gd.Object
		obj.SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var atype = gd.UnsafeGet[gd.VariantType](p_args,1)
		var name = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,2))
		var hint_type = gd.UnsafeGet[gd.PropertyHint](p_args,3)
		var hint_string = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,4))
		var usage_flags = gd.UnsafeGet[gd.PropertyUsageFlags](p_args,5)
		var wide = gd.UnsafeGet[bool](p_args,6)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, atype, name.String(), hint_type, hint_string.String(), usage_flags, wide)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _parse_end(impl func(ptr unsafe.Pointer, obj gd.Object) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var obj gd.Object
		obj.SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, obj)
		gc.End()
	}
}
func (self Simple) AddCustomControl(control [1]classdb.Control) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddCustomControl(control)
}
func (self Simple) AddPropertyEditor(property string, editor [1]classdb.Control, add_to_end bool, label string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddPropertyEditor(gc.String(property), editor, add_to_end, gc.String(label))
}
func (self Simple) AddPropertyEditorForMultipleProperties(label string, properties gd.PackedStringArray, editor [1]classdb.Control) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddPropertyEditorForMultipleProperties(gc.String(label), properties, editor)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorInspectorPlugin
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns [code]true[/code] if this object can be handled by this plugin.
*/
func (class) _can_handle(impl func(ptr unsafe.Pointer, obj gd.Object) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var obj gd.Object
		obj.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Called to allow adding controls at the beginning of the property list for [param object].
*/
func (class) _parse_begin(impl func(ptr unsafe.Pointer, obj gd.Object) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var obj gd.Object
		obj.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, obj)
		ctx.End()
	}
}

/*
Called to allow adding controls at the beginning of a category in the property list for [param object].
*/
func (class) _parse_category(impl func(ptr unsafe.Pointer, obj gd.Object, category gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var obj gd.Object
		obj.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var category = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, obj, category)
		ctx.End()
	}
}

/*
Called to allow adding controls at the beginning of a group or a sub-group in the property list for [param object].
*/
func (class) _parse_group(impl func(ptr unsafe.Pointer, obj gd.Object, group gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var obj gd.Object
		obj.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var group = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, obj, group)
		ctx.End()
	}
}

/*
Called to allow adding property-specific editors to the property list for [param object]. The added editor control must extend [EditorProperty]. Returning [code]true[/code] removes the built-in editor for this property, otherwise allows to insert a custom editor before the built-in one.
*/
func (class) _parse_property(impl func(ptr unsafe.Pointer, obj gd.Object, atype gd.VariantType, name gd.String, hint_type gd.PropertyHint, hint_string gd.String, usage_flags gd.PropertyUsageFlags, wide bool) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var obj gd.Object
		obj.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var atype = gd.UnsafeGet[gd.VariantType](p_args,1)
		var name = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,2))
		var hint_type = gd.UnsafeGet[gd.PropertyHint](p_args,3)
		var hint_string = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,4))
		var usage_flags = gd.UnsafeGet[gd.PropertyUsageFlags](p_args,5)
		var wide = gd.UnsafeGet[bool](p_args,6)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, atype, name, hint_type, hint_string, usage_flags, wide)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Called to allow adding controls at the end of the property list for [param object].
*/
func (class) _parse_end(impl func(ptr unsafe.Pointer, obj gd.Object) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var obj gd.Object
		obj.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, obj)
		ctx.End()
	}
}

/*
Adds a custom control, which is not necessarily a property editor.
*/
//go:nosplit
func (self class) AddCustomControl(control object.Control)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(control[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInspectorPlugin.Bind_add_custom_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a property editor for an individual property. The [param editor] control must extend [EditorProperty].
There can be multiple property editors for a property. If [param add_to_end] is [code]true[/code], this newly added editor will be displayed after all the other editors of the property whose [param add_to_end] is [code]false[/code]. For example, the editor uses this parameter to add an "Edit Region" button for [member Sprite2D.region_rect] below the regular [Rect2] editor.
[param label] can be used to choose a custom label for the property editor in the inspector. If left empty, the label is computed from the name of the property instead.
*/
//go:nosplit
func (self class) AddPropertyEditor(property gd.String, editor object.Control, add_to_end bool, label gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(property))
	callframe.Arg(frame, mmm.End(editor[0].AsPointer())[0])
	callframe.Arg(frame, add_to_end)
	callframe.Arg(frame, mmm.Get(label))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInspectorPlugin.Bind_add_property_editor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds an editor that allows modifying multiple properties. The [param editor] control must extend [EditorProperty].
*/
//go:nosplit
func (self class) AddPropertyEditorForMultipleProperties(label gd.String, properties gd.PackedStringArray, editor object.Control)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, mmm.Get(properties))
	callframe.Arg(frame, mmm.End(editor[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorInspectorPlugin.Bind_add_property_editor_for_multiple_properties, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsEditorInspectorPlugin() Expert { return self[0].AsEditorInspectorPlugin() }


//go:nosplit
func (self Simple) AsEditorInspectorPlugin() Simple { return self[0].AsEditorInspectorPlugin() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_can_handle": return reflect.ValueOf(self._can_handle);
	case "_parse_begin": return reflect.ValueOf(self._parse_begin);
	case "_parse_category": return reflect.ValueOf(self._parse_category);
	case "_parse_group": return reflect.ValueOf(self._parse_group);
	case "_parse_property": return reflect.ValueOf(self._parse_property);
	case "_parse_end": return reflect.ValueOf(self._parse_end);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_can_handle": return reflect.ValueOf(self._can_handle);
	case "_parse_begin": return reflect.ValueOf(self._parse_begin);
	case "_parse_category": return reflect.ValueOf(self._parse_category);
	case "_parse_group": return reflect.ValueOf(self._parse_group);
	case "_parse_property": return reflect.ValueOf(self._parse_property);
	case "_parse_end": return reflect.ValueOf(self._parse_end);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("EditorInspectorPlugin", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
