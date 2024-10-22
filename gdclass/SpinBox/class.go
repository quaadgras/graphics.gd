package SpinBox

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Range"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Go [1]classdb.SpinBox

/*
Applies the current value of this [SpinBox].
*/
func (self Go) Apply() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Apply()
}

/*
Returns the [LineEdit] instance from this [SpinBox]. You can use it to access properties and methods of [LineEdit].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Go) GetLineEdit() gdclass.LineEdit {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.LineEdit(class(self).GetLineEdit(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SpinBox
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SpinBox"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Alignment() gd.HorizontalAlignment {
	gc := gd.GarbageCollector(); _ = gc
		return gd.HorizontalAlignment(class(self).GetHorizontalAlignment())
}

func (self Go) SetAlignment(value gd.HorizontalAlignment) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHorizontalAlignment(value)
}

func (self Go) Editable() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsEditable())
}

func (self Go) SetEditable(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEditable(value)
}

func (self Go) UpdateOnTextChanged() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetUpdateOnTextChanged())
}

func (self Go) SetUpdateOnTextChanged(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUpdateOnTextChanged(value)
}

func (self Go) Prefix() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetPrefix(gc).String())
}

func (self Go) SetPrefix(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPrefix(gc.String(value))
}

func (self Go) Suffix() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetSuffix(gc).String())
}

func (self Go) SetSuffix(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSuffix(gc.String(value))
}

func (self Go) CustomArrowStep() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetCustomArrowStep()))
}

func (self Go) SetCustomArrowStep(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCustomArrowStep(gd.Float(value))
}

func (self Go) SelectAllOnFocus() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSelectAllOnFocus())
}

func (self Go) SetSelectAllOnFocus(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSelectAllOnFocus(value)
}

//go:nosplit
func (self class) SetHorizontalAlignment(alignment gd.HorizontalAlignment)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpinBox.Bind_set_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHorizontalAlignment() gd.HorizontalAlignment {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.HorizontalAlignment](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpinBox.Bind_get_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSuffix(suffix gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(suffix))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpinBox.Bind_set_suffix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSuffix(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpinBox.Bind_get_suffix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPrefix(prefix gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(prefix))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpinBox.Bind_set_prefix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPrefix(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpinBox.Bind_get_prefix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEditable(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpinBox.Bind_set_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCustomArrowStep(arrow_step gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, arrow_step)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpinBox.Bind_set_custom_arrow_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCustomArrowStep() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpinBox.Bind_get_custom_arrow_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) IsEditable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpinBox.Bind_is_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUpdateOnTextChanged(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpinBox.Bind_set_update_on_text_changed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUpdateOnTextChanged() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpinBox.Bind_get_update_on_text_changed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSelectAllOnFocus(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpinBox.Bind_set_select_all_on_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSelectAllOnFocus() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpinBox.Bind_is_select_all_on_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Applies the current value of this [SpinBox].
*/
//go:nosplit
func (self class) Apply()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpinBox.Bind_apply, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [LineEdit] instance from this [SpinBox]. You can use it to access properties and methods of [LineEdit].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetLineEdit(ctx gd.Lifetime) gdclass.LineEdit {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SpinBox.Bind_get_line_edit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.LineEdit
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsSpinBox() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSpinBox() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRange() Range.GD { return *((*Range.GD)(unsafe.Pointer(&self))) }
func (self Go) AsRange() Range.Go { return *((*Range.Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRange(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRange(), name)
	}
}
func init() {classdb.Register("SpinBox", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
