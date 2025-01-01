package EditorProperty

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Container"
import "graphics.gd/objects/Control"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Array"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

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
type Instance [1]classdb.EditorProperty
type Any interface {
	gd.IsClass
	AsEditorProperty() Instance
}

/*
When this virtual function is called, you must update your editor.
*/
func (Instance) _update_property(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the read-only status of the property is changed. It may be used to change custom controls into a read-only or modifiable state.
*/
func (Instance) _set_read_only(impl func(ptr unsafe.Pointer, read_only bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var read_only = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, read_only)
	}
}

/*
Gets the edited property. If your editor is for a single property (added via [method EditorInspectorPlugin._parse_property]), then this will return the property.
*/
func (self Instance) GetEditedProperty() string {
	return string(class(self).GetEditedProperty().String())
}

/*
Gets the edited object.
*/
func (self Instance) GetEditedObject() gd.Object {
	return gd.Object(class(self).GetEditedObject())
}

/*
Forces refresh of the property display.
*/
func (self Instance) UpdateProperty() {
	class(self).UpdateProperty()
}

/*
If any of the controls added can gain keyboard focus, add it here. This ensures that focus will be restored if the inspector is refreshed.
*/
func (self Instance) AddFocusable(control objects.Control) {
	class(self).AddFocusable(control)
}

/*
Puts the [param editor] control below the property label. The control must be previously added using [method Node.add_child].
*/
func (self Instance) SetBottomEditor(editor objects.Control) {
	class(self).SetBottomEditor(editor)
}

/*
If one or several properties have changed, this must be called. [param field] is used in case your editor can modify fields separately (as an example, Vector3.x). The [param changing] argument avoids the editor requesting this property to be refreshed (leave as [code]false[/code] if unsure).
*/
func (self Instance) EmitChanged(property string, value any) {
	class(self).EmitChanged(gd.NewStringName(property), gd.NewVariant(value), gd.NewStringName(""), false)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EditorProperty

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorProperty"))
	return Instance{classdb.EditorProperty(object)}
}

func (self Instance) Label() string {
	return string(class(self).GetLabel().String())
}

func (self Instance) SetLabel(value string) {
	class(self).SetLabel(gd.NewString(value))
}

func (self Instance) ReadOnly() bool {
	return bool(class(self).IsReadOnly())
}

func (self Instance) Checkable() bool {
	return bool(class(self).IsCheckable())
}

func (self Instance) SetCheckable(value bool) {
	class(self).SetCheckable(value)
}

func (self Instance) Checked() bool {
	return bool(class(self).IsChecked())
}

func (self Instance) SetChecked(value bool) {
	class(self).SetChecked(value)
}

func (self Instance) DrawWarning() bool {
	return bool(class(self).IsDrawWarning())
}

func (self Instance) SetDrawWarning(value bool) {
	class(self).SetDrawWarning(value)
}

func (self Instance) Keying() bool {
	return bool(class(self).IsKeying())
}

func (self Instance) SetKeying(value bool) {
	class(self).SetKeying(value)
}

func (self Instance) Deletable() bool {
	return bool(class(self).IsDeletable())
}

func (self Instance) SetDeletable(value bool) {
	class(self).SetDeletable(value)
}

/*
When this virtual function is called, you must update your editor.
*/
func (class) _update_property(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the read-only status of the property is changed. It may be used to change custom controls into a read-only or modifiable state.
*/
func (class) _set_read_only(impl func(ptr unsafe.Pointer, read_only bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var read_only = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, read_only)
	}
}

//go:nosplit
func (self class) SetLabel(text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_set_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLabel() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_get_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetReadOnly(read_only bool) {
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
func (self class) SetCheckable(checkable bool) {
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
func (self class) SetChecked(checked bool) {
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
func (self class) SetDrawWarning(draw_warning bool) {
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
func (self class) SetKeying(keying bool) {
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
func (self class) SetDeletable(deletable bool) {
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
	var ret = pointers.New[gd.StringName](r_ret.Get())
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
func (self class) UpdateProperty() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_update_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
If any of the controls added can gain keyboard focus, add it here. This ensures that focus will be restored if the inspector is refreshed.
*/
//go:nosplit
func (self class) AddFocusable(control objects.Control) {
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
func (self class) SetBottomEditor(editor objects.Control) {
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
func (self class) EmitChanged(property gd.StringName, value gd.Variant, field gd.StringName, changing bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(property))
	callframe.Arg(frame, pointers.Get(value))
	callframe.Arg(frame, pointers.Get(field))
	callframe.Arg(frame, changing)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorProperty.Bind_emit_changed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Instance) OnPropertyChanged(cb func(property string, value any, field string, changing bool)) {
	self[0].AsObject().Connect(gd.NewStringName("property_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnMultiplePropertiesChanged(cb func(properties []string, value Array.Any)) {
	self[0].AsObject().Connect(gd.NewStringName("multiple_properties_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPropertyKeyed(cb func(property string)) {
	self[0].AsObject().Connect(gd.NewStringName("property_keyed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPropertyDeleted(cb func(property string)) {
	self[0].AsObject().Connect(gd.NewStringName("property_deleted"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPropertyKeyedWithValue(cb func(property string, value any)) {
	self[0].AsObject().Connect(gd.NewStringName("property_keyed_with_value"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPropertyChecked(cb func(property string, checked bool)) {
	self[0].AsObject().Connect(gd.NewStringName("property_checked"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPropertyPinned(cb func(property string, pinned bool)) {
	self[0].AsObject().Connect(gd.NewStringName("property_pinned"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPropertyCanRevertChanged(cb func(property string, can_revert bool)) {
	self[0].AsObject().Connect(gd.NewStringName("property_can_revert_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnResourceSelected(cb func(path string, resource objects.Resource)) {
	self[0].AsObject().Connect(gd.NewStringName("resource_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnObjectIdSelected(cb func(property string, id int)) {
	self[0].AsObject().Connect(gd.NewStringName("object_id_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSelected(cb func(path string, focusable_idx int)) {
	self[0].AsObject().Connect(gd.NewStringName("selected"), gd.NewCallable(cb), 0)
}

func (self class) AsEditorProperty() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorProperty() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsContainer() Container.Advanced {
	return *((*Container.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsContainer() Container.Instance {
	return *((*Container.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsControl() Control.Advanced { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_update_property":
		return reflect.ValueOf(self._update_property)
	case "_set_read_only":
		return reflect.ValueOf(self._set_read_only)
	default:
		return gd.VirtualByName(self.AsContainer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_update_property":
		return reflect.ValueOf(self._update_property)
	case "_set_read_only":
		return reflect.ValueOf(self._set_read_only)
	default:
		return gd.VirtualByName(self.AsContainer(), name)
	}
}
func init() {
	classdb.Register("EditorProperty", func(ptr gd.Object) any { return [1]classdb.EditorProperty{classdb.EditorProperty(ptr)} })
}
