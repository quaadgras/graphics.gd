package ScrollContainer

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
A container used to provide a child control with scrollbars when needed. Scrollbars will automatically be drawn at the right (for vertical) or bottom (for horizontal) and will enable dragging to move the viewable Control (and its children) within the ScrollContainer. Scrollbars will also automatically resize the grabber based on the [member Control.custom_minimum_size] of the Control relative to the ScrollContainer.

*/
type Simple [1]classdb.ScrollContainer
func (self Simple) SetHScroll(value int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHScroll(gd.Int(value))
}
func (self Simple) GetHScroll() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetHScroll()))
}
func (self Simple) SetVScroll(value int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVScroll(gd.Int(value))
}
func (self Simple) GetVScroll() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetVScroll()))
}
func (self Simple) SetHorizontalCustomStep(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHorizontalCustomStep(gd.Float(value))
}
func (self Simple) GetHorizontalCustomStep() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetHorizontalCustomStep()))
}
func (self Simple) SetVerticalCustomStep(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVerticalCustomStep(gd.Float(value))
}
func (self Simple) GetVerticalCustomStep() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVerticalCustomStep()))
}
func (self Simple) SetHorizontalScrollMode(enable classdb.ScrollContainerScrollMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHorizontalScrollMode(enable)
}
func (self Simple) GetHorizontalScrollMode() classdb.ScrollContainerScrollMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ScrollContainerScrollMode(Expert(self).GetHorizontalScrollMode())
}
func (self Simple) SetVerticalScrollMode(enable classdb.ScrollContainerScrollMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVerticalScrollMode(enable)
}
func (self Simple) GetVerticalScrollMode() classdb.ScrollContainerScrollMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ScrollContainerScrollMode(Expert(self).GetVerticalScrollMode())
}
func (self Simple) SetDeadzone(deadzone int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDeadzone(gd.Int(deadzone))
}
func (self Simple) GetDeadzone() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetDeadzone()))
}
func (self Simple) SetFollowFocus(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFollowFocus(enabled)
}
func (self Simple) IsFollowingFocus() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsFollowingFocus())
}
func (self Simple) GetHScrollBar() [1]classdb.HScrollBar {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.HScrollBar(Expert(self).GetHScrollBar(gc))
}
func (self Simple) GetVScrollBar() [1]classdb.VScrollBar {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.VScrollBar(Expert(self).GetVScrollBar(gc))
}
func (self Simple) EnsureControlVisible(control [1]classdb.Control) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnsureControlVisible(control)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ScrollContainer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetHScroll(value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScrollContainer.Bind_set_h_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHScroll() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScrollContainer.Bind_get_h_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVScroll(value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScrollContainer.Bind_set_v_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVScroll() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScrollContainer.Bind_get_v_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHorizontalCustomStep(value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScrollContainer.Bind_set_horizontal_custom_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHorizontalCustomStep() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScrollContainer.Bind_get_horizontal_custom_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVerticalCustomStep(value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScrollContainer.Bind_set_vertical_custom_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVerticalCustomStep() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScrollContainer.Bind_get_vertical_custom_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHorizontalScrollMode(enable classdb.ScrollContainerScrollMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScrollContainer.Bind_set_horizontal_scroll_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHorizontalScrollMode() classdb.ScrollContainerScrollMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ScrollContainerScrollMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScrollContainer.Bind_get_horizontal_scroll_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVerticalScrollMode(enable classdb.ScrollContainerScrollMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScrollContainer.Bind_set_vertical_scroll_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVerticalScrollMode() classdb.ScrollContainerScrollMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ScrollContainerScrollMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScrollContainer.Bind_get_vertical_scroll_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDeadzone(deadzone gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, deadzone)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScrollContainer.Bind_set_deadzone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDeadzone() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScrollContainer.Bind_get_deadzone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFollowFocus(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScrollContainer.Bind_set_follow_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFollowingFocus() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScrollContainer.Bind_is_following_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the horizontal scrollbar [HScrollBar] of this [ScrollContainer].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to disable or hide a scrollbar, you can use [member horizontal_scroll_mode].
*/
//go:nosplit
func (self class) GetHScrollBar(ctx gd.Lifetime) object.HScrollBar {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScrollContainer.Bind_get_h_scroll_bar, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.HScrollBar
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the vertical scrollbar [VScrollBar] of this [ScrollContainer].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to disable or hide a scrollbar, you can use [member vertical_scroll_mode].
*/
//go:nosplit
func (self class) GetVScrollBar(ctx gd.Lifetime) object.VScrollBar {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScrollContainer.Bind_get_v_scroll_bar, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.VScrollBar
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Ensures the given [param control] is visible (must be a direct or indirect child of the ScrollContainer). Used by [member follow_focus].
[b]Note:[/b] This will not work on a node that was just added during the same frame. If you want to scroll to a newly added child, you must wait until the next frame using [signal SceneTree.process_frame]:
[codeblock]
add_child(child_node)
await get_tree().process_frame
ensure_control_visible(child_node)
[/codeblock]
*/
//go:nosplit
func (self class) EnsureControlVisible(control object.Control)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(control[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScrollContainer.Bind_ensure_control_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsScrollContainer() Expert { return self[0].AsScrollContainer() }


//go:nosplit
func (self Simple) AsScrollContainer() Simple { return self[0].AsScrollContainer() }


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
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("ScrollContainer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type ScrollMode = classdb.ScrollContainerScrollMode

const (
/*Scrolling disabled, scrollbar will be invisible.*/
	ScrollModeDisabled ScrollMode = 0
/*Scrolling enabled, scrollbar will be visible only if necessary, i.e. container's content is bigger than the container.*/
	ScrollModeAuto ScrollMode = 1
/*Scrolling enabled, scrollbar will be always visible.*/
	ScrollModeShowAlways ScrollMode = 2
/*Scrolling enabled, scrollbar will be hidden.*/
	ScrollModeShowNever ScrollMode = 3
)
