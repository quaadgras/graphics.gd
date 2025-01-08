// Package SpinBox provides methods for working with SpinBox object instances.
package SpinBox

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Range"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[SpinBox] is a numerical input text field. It allows entering integers and floating-point numbers.
[b]Example:[/b]
[codeblocks]
[gdscript]
var spin_box = SpinBox.new()
add_child(spin_box)
var line_edit = spin_box.get_line_edit()
line_edit.context_menu_enabled = false
spin_box.horizontal_alignment = LineEdit.HORIZONTAL_ALIGNMENT_RIGHT
[/gdscript]
[csharp]
var spinBox = new SpinBox();
AddChild(spinBox);
var lineEdit = spinBox.GetLineEdit();
lineEdit.ContextMenuEnabled = false;
spinBox.AlignHorizontal = LineEdit.HorizontalAlignEnum.Right;
[/csharp]
[/codeblocks]
The above code will create a [SpinBox], disable context menu on it and set the text alignment to right.
See [Range] class for more options over the [SpinBox].
[b]Note:[/b] With the [SpinBox]'s context menu disabled, you can right-click the bottom half of the spinbox to set the value to its minimum, while right-clicking the top half sets the value to its maximum.
[b]Note:[/b] [SpinBox] relies on an underlying [LineEdit] node. To theme a [SpinBox]'s background, add theme items for [LineEdit] and customize them.
[b]Note:[/b] If you want to implement drag and drop for the underlying [LineEdit], you can use [method Control.set_drag_forwarding] on the node returned by [method get_line_edit].
*/
type Instance [1]gdclass.SpinBox

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSpinBox() Instance
}

/*
Applies the current value of this [SpinBox].
*/
func (self Instance) Apply() {
	class(self).Apply()
}

/*
Returns the [LineEdit] instance from this [SpinBox]. You can use it to access properties and methods of [LineEdit].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Instance) GetLineEdit() [1]gdclass.LineEdit {
	return [1]gdclass.LineEdit(class(self).GetLineEdit())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SpinBox

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SpinBox"))
	casted := Instance{*(*gdclass.SpinBox)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Alignment() HorizontalAlignment {
	return HorizontalAlignment(class(self).GetHorizontalAlignment())
}

func (self Instance) SetAlignment(value HorizontalAlignment) {
	class(self).SetHorizontalAlignment(value)
}

func (self Instance) Editable() bool {
	return bool(class(self).IsEditable())
}

func (self Instance) SetEditable(value bool) {
	class(self).SetEditable(value)
}

func (self Instance) UpdateOnTextChanged() bool {
	return bool(class(self).GetUpdateOnTextChanged())
}

func (self Instance) SetUpdateOnTextChanged(value bool) {
	class(self).SetUpdateOnTextChanged(value)
}

func (self Instance) Prefix() string {
	return string(class(self).GetPrefix().String())
}

func (self Instance) SetPrefix(value string) {
	class(self).SetPrefix(gd.NewString(value))
}

func (self Instance) Suffix() string {
	return string(class(self).GetSuffix().String())
}

func (self Instance) SetSuffix(value string) {
	class(self).SetSuffix(gd.NewString(value))
}

func (self Instance) CustomArrowStep() Float.X {
	return Float.X(Float.X(class(self).GetCustomArrowStep()))
}

func (self Instance) SetCustomArrowStep(value Float.X) {
	class(self).SetCustomArrowStep(gd.Float(value))
}

func (self Instance) SelectAllOnFocus() bool {
	return bool(class(self).IsSelectAllOnFocus())
}

func (self Instance) SetSelectAllOnFocus(value bool) {
	class(self).SetSelectAllOnFocus(value)
}

//go:nosplit
func (self class) SetHorizontalAlignment(alignment HorizontalAlignment) {
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpinBox.Bind_set_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHorizontalAlignment() HorizontalAlignment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[HorizontalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpinBox.Bind_get_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSuffix(suffix gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(suffix))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpinBox.Bind_set_suffix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSuffix() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpinBox.Bind_get_suffix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPrefix(prefix gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(prefix))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpinBox.Bind_set_prefix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPrefix() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpinBox.Bind_get_prefix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEditable(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpinBox.Bind_set_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetCustomArrowStep(arrow_step gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, arrow_step)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpinBox.Bind_set_custom_arrow_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCustomArrowStep() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpinBox.Bind_get_custom_arrow_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) IsEditable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpinBox.Bind_is_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUpdateOnTextChanged(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpinBox.Bind_set_update_on_text_changed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetUpdateOnTextChanged() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpinBox.Bind_get_update_on_text_changed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSelectAllOnFocus(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpinBox.Bind_set_select_all_on_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSelectAllOnFocus() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpinBox.Bind_is_select_all_on_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Applies the current value of this [SpinBox].
*/
//go:nosplit
func (self class) Apply() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpinBox.Bind_apply, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [LineEdit] instance from this [SpinBox]. You can use it to access properties and methods of [LineEdit].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetLineEdit() [1]gdclass.LineEdit {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpinBox.Bind_get_line_edit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.LineEdit{gd.PointerLifetimeBoundTo[gdclass.LineEdit](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsSpinBox() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSpinBox() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRange() Range.Advanced     { return *((*Range.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRange() Range.Instance  { return *((*Range.Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("SpinBox", func(ptr gd.Object) any { return [1]gdclass.SpinBox{*(*gdclass.SpinBox)(unsafe.Pointer(&ptr))} })
}

type HorizontalAlignment int

const (
	/*Horizontal left alignment, usually for text-derived classes.*/
	HorizontalAlignmentLeft HorizontalAlignment = 0
	/*Horizontal center alignment, usually for text-derived classes.*/
	HorizontalAlignmentCenter HorizontalAlignment = 1
	/*Horizontal right alignment, usually for text-derived classes.*/
	HorizontalAlignmentRight HorizontalAlignment = 2
	/*Expand row to fit width, usually for text-derived classes.*/
	HorizontalAlignmentFill HorizontalAlignment = 3
)
