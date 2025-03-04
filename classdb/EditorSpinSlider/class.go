// Package EditorSpinSlider provides methods for working with EditorSpinSlider object instances.
package EditorSpinSlider

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Range"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
This [Control] node is used in the editor's Inspector dock to allow editing of numeric values. Can be used with [EditorInspectorPlugin] to recreate the same behavior.
If the [member Range.step] value is [code]1[/code], the [EditorSpinSlider] will display up/down arrows, similar to [SpinBox]. If the [member Range.step] value is not [code]1[/code], a slider will be displayed instead.
*/
type Instance [1]gdclass.EditorSpinSlider

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorSpinSlider() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorSpinSlider

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorSpinSlider"))
	casted := Instance{*(*gdclass.EditorSpinSlider)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Label() string {
	return string(class(self).GetLabel().String())
}

func (self Instance) SetLabel(value string) {
	class(self).SetLabel(String.New(value))
}

func (self Instance) Suffix() string {
	return string(class(self).GetSuffix().String())
}

func (self Instance) SetSuffix(value string) {
	class(self).SetSuffix(String.New(value))
}

func (self Instance) ReadOnly() bool {
	return bool(class(self).IsReadOnly())
}

func (self Instance) SetReadOnly(value bool) {
	class(self).SetReadOnly(value)
}

func (self Instance) Flat() bool {
	return bool(class(self).IsFlat())
}

func (self Instance) SetFlat(value bool) {
	class(self).SetFlat(value)
}

func (self Instance) HideSlider() bool {
	return bool(class(self).IsHidingSlider())
}

func (self Instance) SetHideSlider(value bool) {
	class(self).SetHideSlider(value)
}

func (self Instance) EditingInteger() bool {
	return bool(class(self).IsEditingInteger())
}

func (self Instance) SetEditingInteger(value bool) {
	class(self).SetEditingInteger(value)
}

//go:nosplit
func (self class) SetLabel(label String.Readable) { //gd:EditorSpinSlider.set_label
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(label)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSpinSlider.Bind_set_label, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLabel() String.Readable { //gd:EditorSpinSlider.get_label
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSpinSlider.Bind_get_label, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSuffix(suffix String.Readable) { //gd:EditorSpinSlider.set_suffix
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(suffix)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSpinSlider.Bind_set_suffix, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSuffix() String.Readable { //gd:EditorSpinSlider.get_suffix
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSpinSlider.Bind_get_suffix, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetReadOnly(read_only bool) { //gd:EditorSpinSlider.set_read_only
	var frame = callframe.New()
	callframe.Arg(frame, read_only)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSpinSlider.Bind_set_read_only, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsReadOnly() bool { //gd:EditorSpinSlider.is_read_only
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSpinSlider.Bind_is_read_only, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlat(flat bool) { //gd:EditorSpinSlider.set_flat
	var frame = callframe.New()
	callframe.Arg(frame, flat)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSpinSlider.Bind_set_flat, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlat() bool { //gd:EditorSpinSlider.is_flat
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSpinSlider.Bind_is_flat, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHideSlider(hide_slider bool) { //gd:EditorSpinSlider.set_hide_slider
	var frame = callframe.New()
	callframe.Arg(frame, hide_slider)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSpinSlider.Bind_set_hide_slider, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsHidingSlider() bool { //gd:EditorSpinSlider.is_hiding_slider
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSpinSlider.Bind_is_hiding_slider, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEditingInteger(editing_integer bool) { //gd:EditorSpinSlider.set_editing_integer
	var frame = callframe.New()
	callframe.Arg(frame, editing_integer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSpinSlider.Bind_set_editing_integer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsEditingInteger() bool { //gd:EditorSpinSlider.is_editing_integer
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSpinSlider.Bind_is_editing_integer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnGrabbed(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("grabbed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnUngrabbed(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("ungrabbed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnUpdownPressed(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("updown_pressed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnValueFocusEntered(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("value_focus_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnValueFocusExited(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("value_focus_exited"), gd.NewCallable(cb), 0)
}

func (self class) AsEditorSpinSlider() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorSpinSlider() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRange() Range.Advanced         { return *((*Range.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRange() Range.Instance      { return *((*Range.Instance)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.Advanced     { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
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
	default:
		return gd.VirtualByName(Range.Advanced(self.AsRange()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Range.Instance(self.AsRange()), name)
	}
}
func init() {
	gdclass.Register("EditorSpinSlider", func(ptr gd.Object) any {
		return [1]gdclass.EditorSpinSlider{*(*gdclass.EditorSpinSlider)(unsafe.Pointer(&ptr))}
	})
}
