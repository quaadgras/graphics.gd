package EditorProperty

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Container"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A custom control for editing properties that can be added to the [EditorInspector]. It is added via [EditorInspectorPlugin].
	// EditorProperty methods that can be overridden by a [Class] that extends it.
	type EditorProperty interface {
		//When this virtual function is called, you must update your editor.
		UpdateProperty() 
		//Called when the read-only status of the property is changed. It may be used to change custom controls into a read-only or modifiable state.
		SetReadOnly(read_only bool) 
	}

*/
type Simple [1]classdb.EditorProperty
func (Simple) _update_property(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (Simple) _set_read_only(impl func(ptr unsafe.Pointer, read_only bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var read_only = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, read_only)
		gc.End()
	}
}
func (self Simple) SetLabel(text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLabel(gc.String(text))
}
func (self Simple) GetLabel() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetLabel(gc).String())
}
func (self Simple) SetReadOnly(read_only bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetReadOnly(read_only)
}
func (self Simple) IsReadOnly() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsReadOnly())
}
func (self Simple) SetCheckable(checkable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCheckable(checkable)
}
func (self Simple) IsCheckable() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsCheckable())
}
func (self Simple) SetChecked(checked bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetChecked(checked)
}
func (self Simple) IsChecked() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsChecked())
}
func (self Simple) SetDrawWarning(draw_warning bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDrawWarning(draw_warning)
}
func (self Simple) IsDrawWarning() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDrawWarning())
}
func (self Simple) SetKeying(keying bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetKeying(keying)
}
func (self Simple) IsKeying() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsKeying())
}
func (self Simple) SetDeletable(deletable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDeletable(deletable)
}
func (self Simple) IsDeletable() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDeletable())
}
func (self Simple) GetEditedProperty() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetEditedProperty(gc).String())
}
func (self Simple) GetEditedObject() gd.Object {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Object(Expert(self).GetEditedObject(gc))
}
func (self Simple) UpdateProperty() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UpdateProperty()
}
func (self Simple) AddFocusable(control [1]classdb.Control) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddFocusable(control)
}
func (self Simple) SetBottomEditor(editor [1]classdb.Control) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBottomEditor(editor)
}
func (self Simple) EmitChanged(property string, value gd.Variant, field string, changing bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EmitChanged(gc.StringName(property), value, gc.StringName(field), changing)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorProperty
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
When this virtual function is called, you must update your editor.
*/
func (class) _update_property(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called when the read-only status of the property is changed. It may be used to change custom controls into a read-only or modifiable state.
*/
func (class) _set_read_only(impl func(ptr unsafe.Pointer, read_only bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var read_only = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, read_only)
		ctx.End()
	}
}

//go:nosplit
func (self class) SetLabel(text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_set_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLabel(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_get_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetReadOnly(read_only bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, read_only)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_set_read_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsReadOnly() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_is_read_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCheckable(checkable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, checkable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_set_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCheckable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_is_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetChecked(checked bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, checked)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_set_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsChecked() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_is_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDrawWarning(draw_warning bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, draw_warning)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_set_draw_warning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDrawWarning() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_is_draw_warning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetKeying(keying bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, keying)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_set_keying, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsKeying() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_is_keying, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDeletable(deletable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, deletable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_set_deletable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDeletable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_is_deletable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets the edited property. If your editor is for a single property (added via [method EditorInspectorPlugin._parse_property]), then this will return the property.
*/
//go:nosplit
func (self class) GetEditedProperty(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_get_edited_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Gets the edited object.
*/
//go:nosplit
func (self class) GetEditedObject(ctx gd.Lifetime) gd.Object {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_get_edited_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gd.Object
	ret.SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Forces refresh of the property display.
*/
//go:nosplit
func (self class) UpdateProperty()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_update_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If any of the controls added can gain keyboard focus, add it here. This ensures that focus will be restored if the inspector is refreshed.
*/
//go:nosplit
func (self class) AddFocusable(control object.Control)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(control[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_add_focusable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Puts the [param editor] control below the property label. The control must be previously added using [method Node.add_child].
*/
//go:nosplit
func (self class) SetBottomEditor(editor object.Control)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(editor[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_set_bottom_editor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If one or several properties have changed, this must be called. [param field] is used in case your editor can modify fields separately (as an example, Vector3.x). The [param changing] argument avoids the editor requesting this property to be refreshed (leave as [code]false[/code] if unsure).
*/
//go:nosplit
func (self class) EmitChanged(property gd.StringName, value gd.Variant, field gd.StringName, changing bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(property))
	callframe.Arg(frame, mmm.Get(value))
	callframe.Arg(frame, mmm.Get(field))
	callframe.Arg(frame, changing)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorProperty.Bind_emit_changed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsEditorProperty() Expert { return self[0].AsEditorProperty() }


//go:nosplit
func (self Simple) AsEditorProperty() Simple { return self[0].AsEditorProperty() }


//go:nosplit
func (self class) AsContainer() Container.Expert { return self[0].AsContainer() }


//go:nosplit
func (self Simple) AsContainer() Container.Simple { return self[0].AsContainer() }


//go:nosplit
func (self class) AsControl() Control.Expert { return self[0].AsControl() }


//go:nosplit
func (self Simple) AsControl() Control.Simple { return self[0].AsControl() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_update_property": return reflect.ValueOf(self._update_property);
	case "_set_read_only": return reflect.ValueOf(self._set_read_only);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_update_property": return reflect.ValueOf(self._update_property);
	case "_set_read_only": return reflect.ValueOf(self._set_read_only);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("EditorProperty", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
