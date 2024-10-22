package ProgressBar

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
A control used for visual representation of a percentage. Shows fill percentage from right to left.

*/
type Go [1]classdb.ProgressBar
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ProgressBar
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ProgressBar"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) FillMode() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetFillMode()))
}

func (self Go) SetFillMode(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFillMode(gd.Int(value))
}

func (self Go) ShowPercentage() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsPercentageShown())
}

func (self Go) SetShowPercentage(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShowPercentage(value)
}

func (self Go) Indeterminate() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsIndeterminate())
}

func (self Go) SetIndeterminate(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIndeterminate(value)
}

func (self Go) EditorPreviewIndeterminate() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsEditorPreviewIndeterminateEnabled())
}

func (self Go) SetEditorPreviewIndeterminate(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEditorPreviewIndeterminate(value)
}

//go:nosplit
func (self class) SetFillMode(mode gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProgressBar.Bind_set_fill_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFillMode() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProgressBar.Bind_get_fill_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowPercentage(visible bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProgressBar.Bind_set_show_percentage, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsPercentageShown() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProgressBar.Bind_is_percentage_shown, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIndeterminate(indeterminate bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, indeterminate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProgressBar.Bind_set_indeterminate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsIndeterminate() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProgressBar.Bind_is_indeterminate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEditorPreviewIndeterminate(preview_indeterminate bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, preview_indeterminate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProgressBar.Bind_set_editor_preview_indeterminate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEditorPreviewIndeterminateEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProgressBar.Bind_is_editor_preview_indeterminate_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsProgressBar() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsProgressBar() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("ProgressBar", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type FillMode = classdb.ProgressBarFillMode

const (
/*The progress bar fills from begin to end horizontally, according to the language direction. If [method Control.is_layout_rtl] returns [code]false[/code], it fills from left to right, and if it returns [code]true[/code], it fills from right to left.*/
	FillBeginToEnd FillMode = 0
/*The progress bar fills from end to begin horizontally, according to the language direction. If [method Control.is_layout_rtl] returns [code]false[/code], it fills from right to left, and if it returns [code]true[/code], it fills from left to right.*/
	FillEndToBegin FillMode = 1
/*The progress fills from top to bottom.*/
	FillTopToBottom FillMode = 2
/*The progress fills from bottom to top.*/
	FillBottomToTop FillMode = 3
)
