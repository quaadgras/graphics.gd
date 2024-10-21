package ProgressBar

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Range"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A control used for visual representation of a percentage. Shows fill percentage from right to left.

*/
type Simple [1]classdb.ProgressBar
func (self Simple) SetFillMode(mode int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFillMode(gd.Int(mode))
}
func (self Simple) GetFillMode() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetFillMode()))
}
func (self Simple) SetShowPercentage(visible bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShowPercentage(visible)
}
func (self Simple) IsPercentageShown() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsPercentageShown())
}
func (self Simple) SetIndeterminate(indeterminate bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIndeterminate(indeterminate)
}
func (self Simple) IsIndeterminate() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsIndeterminate())
}
func (self Simple) SetEditorPreviewIndeterminate(preview_indeterminate bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEditorPreviewIndeterminate(preview_indeterminate)
}
func (self Simple) IsEditorPreviewIndeterminateEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsEditorPreviewIndeterminateEnabled())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ProgressBar
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsProgressBar() Expert { return self[0].AsProgressBar() }


//go:nosplit
func (self Simple) AsProgressBar() Simple { return self[0].AsProgressBar() }


//go:nosplit
func (self class) AsRange() Range.Expert { return self[0].AsRange() }


//go:nosplit
func (self Simple) AsRange() Range.Simple { return self[0].AsRange() }


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
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("ProgressBar", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
