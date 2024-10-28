package EditorProperty

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Container"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

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
type Go [1]classdb.EditorProperty

/*
When this virtual function is called, you must update your editor.
*/
func (Go) _update_property(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
	}
}

/*
Called when the read-only status of the property is changed. It may be used to change custom controls into a read-only or modifiable state.
*/
func (Go) _set_read_only(impl func(ptr unsafe.Pointer, read_only bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var read_only = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, read_only)
	}
}

/*
Gets the edited property. If your editor is for a single property (added via [method EditorInspectorPlugin._parse_property]), then this will return the property.
*/
func (self Go) GetEditedProperty() string {
	return string(class(self).GetEditedProperty().String())
}

/*
Gets the edited object.
*/
func (self Go) GetEditedObject() gd.Object {
	return gd.Object(class(self).GetEditedObject())
}

/*
Forces refresh of the property display.
*/
func (self Go) UpdateProperty() {
	class(self).UpdateProperty()
}

/*
If any of the controls added can gain keyboard focus, add it here. This ensures that focus will be restored if the inspector is refreshed.
*/
func (self Go) AddFocusable(control gdclass.Control) {
	class(self).AddFocusable(control)
}

/*
Puts the [param editor] control below the property label. The control must be previously added using [method Node.add_child].
*/
func (self Go) SetBottomEditor(editor gdclass.Control) {
	class(self).SetBottomEditor(editor)
}

/*
If one or several properties have changed, this must be called. [param field] is used in case your editor can modify fields separately (as an example, Vector3.x). The [param changing] argument avoids the editor requesting this property to be refreshed (leave as [code]false[/code] if unsure).
*/
func (self Go) EmitChanged(property string, value gd.Variant) {
	class(self).EmitChanged(gd.NewStringName(property), value, gd.NewStringName(""), false)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorProperty
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorProperty"))
	return Go{classdb.EditorProperty(object)}
}

func (self Go) Label() string {
		return string(class(self).GetLabel().String())
}

func (self Go) SetLabel(value string) {
	class(self).SetLabel(gd.NewString(value))
}

func (self Go) ReadOnly() bool {
		return bool(class(self).IsReadOnly())
}

func (self Go) Checkable() bool {
		return bool(class(self).IsCheckable())
}

func (self Go) SetCheckable(value bool) {
	class(self).SetCheckable(value)
}

func (self Go) Checked() bool {
		return bool(class(self).IsChecked())
}

func (self Go) SetChecked(value bool) {
	class(self).SetChecked(value)
}

func (self Go) DrawWarning() bool {
		return bool(class(self).IsDrawWarning())
}

func (self Go) SetDrawWarning(value bool) {
	class(self).SetDrawWarning(value)
}

func (self Go) Keying() bool {
		return bool(class(self).IsKeying())
}

func (self Go) SetKeying(value bool) {
	class(self).SetKeying(value)
}

func (self Go) Deletable() bool {
		return bool(class(self).IsDeletable())
}

func (self Go) SetDeletable(value bool) {
	class(self).SetDeletable(value)
}

/*
When this virtual function is called, you must update your editor.
*/
func (class) _update_property(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
	}
}

/*
Called when the read-only status of the property is changed. It may be used to change custom controls into a read-only or modifiable state.
*/
func (class) _set_read_only(impl func(ptr unsafe.Pointer, read_only bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var read_only = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, read_only)
	}
}

//go:nosplit
func (self class) SetLabel(text gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_set_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLabel() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_get_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetReadOnly(read_only bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, read_only)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_set_read_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsReadOnly() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_is_read_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCheckable(checkable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, checkable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_set_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCheckable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_is_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetChecked(checked bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, checked)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_set_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsChecked() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_is_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDrawWarning(draw_warning bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, draw_warning)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_set_draw_warning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDrawWarning() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_is_draw_warning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetKeying(keying bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, keying)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_set_keying, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsKeying() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_is_keying, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDeletable(deletable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, deletable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_set_deletable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDeletable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_is_deletable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets the edited property. If your editor is for a single property (added via [method EditorInspectorPlugin._parse_property]), then this will return the property.
*/
//go:nosplit
func (self class) GetEditedProperty() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_get_edited_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
/*
Gets the edited object.
*/
//go:nosplit
func (self class) GetEditedObject() gd.Object {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_get_edited_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gd.PointerWithOwnershipTransferredToGo(r_ret.Get())
	frame.Free()
	return ret
}
/*
Forces refresh of the property display.
*/
//go:nosplit
func (self class) UpdateProperty()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_update_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If any of the controls added can gain keyboard focus, add it here. This ensures that focus will be restored if the inspector is refreshed.
*/
//go:nosplit
func (self class) AddFocusable(control gdclass.Control)  {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(gd.Object(control[0])))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_add_focusable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Puts the [param editor] control below the property label. The control must be previously added using [method Node.add_child].
*/
//go:nosplit
func (self class) SetBottomEditor(editor gdclass.Control)  {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(gd.Object(editor[0])))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_set_bottom_editor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If one or several properties have changed, this must be called. [param field] is used in case your editor can modify fields separately (as an example, Vector3.x). The [param changing] argument avoids the editor requesting this property to be refreshed (leave as [code]false[/code] if unsure).
*/
//go:nosplit
func (self class) EmitChanged(property gd.StringName, value gd.Variant, field gd.StringName, changing bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(property))
	callframe.Arg(frame, discreet.Get(value))
	callframe.Arg(frame, discreet.Get(field))
	callframe.Arg(frame, changing)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_emit_changed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnPropertyChanged(cb func(property string, value gd.Variant, field string, changing bool)) {
	self[0].AsObject().Connect(gd.NewStringName("property_changed"), gd.NewCallable(cb), 0)
}


func (self Go) OnMultiplePropertiesChanged(cb func(properties []string, value gd.Array)) {
	self[0].AsObject().Connect(gd.NewStringName("multiple_properties_changed"), gd.NewCallable(cb), 0)
}


func (self Go) OnPropertyKeyed(cb func(property string)) {
	self[0].AsObject().Connect(gd.NewStringName("property_keyed"), gd.NewCallable(cb), 0)
}


func (self Go) OnPropertyDeleted(cb func(property string)) {
	self[0].AsObject().Connect(gd.NewStringName("property_deleted"), gd.NewCallable(cb), 0)
}


func (self Go) OnPropertyKeyedWithValue(cb func(property string, value gd.Variant)) {
	self[0].AsObject().Connect(gd.NewStringName("property_keyed_with_value"), gd.NewCallable(cb), 0)
}


func (self Go) OnPropertyChecked(cb func(property string, checked bool)) {
	self[0].AsObject().Connect(gd.NewStringName("property_checked"), gd.NewCallable(cb), 0)
}


func (self Go) OnPropertyPinned(cb func(property string, pinned bool)) {
	self[0].AsObject().Connect(gd.NewStringName("property_pinned"), gd.NewCallable(cb), 0)
}


func (self Go) OnPropertyCanRevertChanged(cb func(property string, can_revert bool)) {
	self[0].AsObject().Connect(gd.NewStringName("property_can_revert_changed"), gd.NewCallable(cb), 0)
}


func (self Go) OnResourceSelected(cb func(path string, resource gdclass.Resource)) {
	self[0].AsObject().Connect(gd.NewStringName("resource_selected"), gd.NewCallable(cb), 0)
}


func (self Go) OnObjectIdSelected(cb func(property string, id int)) {
	self[0].AsObject().Connect(gd.NewStringName("object_id_selected"), gd.NewCallable(cb), 0)
}


func (self Go) OnSelected(cb func(path string, focusable_idx int)) {
	self[0].AsObject().Connect(gd.NewStringName("selected"), gd.NewCallable(cb), 0)
}


func (self class) AsEditorProperty() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorProperty() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsContainer() Container.GD { return *((*Container.GD)(unsafe.Pointer(&self))) }
func (self Go) AsContainer() Container.Go { return *((*Container.Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_update_property": return reflect.ValueOf(self._update_property);
	case "_set_read_only": return reflect.ValueOf(self._set_read_only);
	default: return gd.VirtualByName(self.AsContainer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_update_property": return reflect.ValueOf(self._update_property);
	case "_set_read_only": return reflect.ValueOf(self._set_read_only);
	default: return gd.VirtualByName(self.AsContainer(), name)
	}
}
func init() {classdb.Register("EditorProperty", func(ptr gd.Object) any { return classdb.EditorProperty(ptr) })}
