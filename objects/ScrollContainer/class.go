package ScrollContainer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Container"
import "grow.graphics/gd/objects/Control"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"
import "grow.graphics/gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A container used to provide a child control with scrollbars when needed. Scrollbars will automatically be drawn at the right (for vertical) or bottom (for horizontal) and will enable dragging to move the viewable Control (and its children) within the ScrollContainer. Scrollbars will also automatically resize the grabber based on the [member Control.custom_minimum_size] of the Control relative to the ScrollContainer.
*/
type Instance [1]classdb.ScrollContainer

/*
Returns the horizontal scrollbar [HScrollBar] of this [ScrollContainer].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to disable or hide a scrollbar, you can use [member horizontal_scroll_mode].
*/
func (self Instance) GetHScrollBar() objects.HScrollBar {
	return objects.HScrollBar(class(self).GetHScrollBar())
}

/*
Returns the vertical scrollbar [VScrollBar] of this [ScrollContainer].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to disable or hide a scrollbar, you can use [member vertical_scroll_mode].
*/
func (self Instance) GetVScrollBar() objects.VScrollBar {
	return objects.VScrollBar(class(self).GetVScrollBar())
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
func (self Instance) EnsureControlVisible(control objects.Control) {
	class(self).EnsureControlVisible(control)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ScrollContainer

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ScrollContainer"))
	return Instance{classdb.ScrollContainer(object)}
}

func (self Instance) FollowFocus() bool {
	return bool(class(self).IsFollowingFocus())
}

func (self Instance) SetFollowFocus(value bool) {
	class(self).SetFollowFocus(value)
}

func (self Instance) ScrollHorizontal() int {
	return int(int(class(self).GetHScroll()))
}

func (self Instance) SetScrollHorizontal(value int) {
	class(self).SetHScroll(gd.Int(value))
}

func (self Instance) ScrollVertical() int {
	return int(int(class(self).GetVScroll()))
}

func (self Instance) SetScrollVertical(value int) {
	class(self).SetVScroll(gd.Int(value))
}

func (self Instance) ScrollHorizontalCustomStep() Float.X {
	return Float.X(Float.X(class(self).GetHorizontalCustomStep()))
}

func (self Instance) SetScrollHorizontalCustomStep(value Float.X) {
	class(self).SetHorizontalCustomStep(gd.Float(value))
}

func (self Instance) ScrollVerticalCustomStep() Float.X {
	return Float.X(Float.X(class(self).GetVerticalCustomStep()))
}

func (self Instance) SetScrollVerticalCustomStep(value Float.X) {
	class(self).SetVerticalCustomStep(gd.Float(value))
}

func (self Instance) HorizontalScrollMode() classdb.ScrollContainerScrollMode {
	return classdb.ScrollContainerScrollMode(class(self).GetHorizontalScrollMode())
}

func (self Instance) SetHorizontalScrollMode(value classdb.ScrollContainerScrollMode) {
	class(self).SetHorizontalScrollMode(value)
}

func (self Instance) VerticalScrollMode() classdb.ScrollContainerScrollMode {
	return classdb.ScrollContainerScrollMode(class(self).GetVerticalScrollMode())
}

func (self Instance) SetVerticalScrollMode(value classdb.ScrollContainerScrollMode) {
	class(self).SetVerticalScrollMode(value)
}

func (self Instance) ScrollDeadzone() int {
	return int(int(class(self).GetDeadzone()))
}

func (self Instance) SetScrollDeadzone(value int) {
	class(self).SetDeadzone(gd.Int(value))
}

//go:nosplit
func (self class) SetHScroll(value gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_set_h_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHScroll() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_get_h_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVScroll(value gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_set_v_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVScroll() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_get_v_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHorizontalCustomStep(value gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_set_horizontal_custom_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHorizontalCustomStep() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_get_horizontal_custom_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVerticalCustomStep(value gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_set_vertical_custom_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVerticalCustomStep() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_get_vertical_custom_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHorizontalScrollMode(enable classdb.ScrollContainerScrollMode) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_set_horizontal_scroll_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHorizontalScrollMode() classdb.ScrollContainerScrollMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ScrollContainerScrollMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_get_horizontal_scroll_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVerticalScrollMode(enable classdb.ScrollContainerScrollMode) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_set_vertical_scroll_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVerticalScrollMode() classdb.ScrollContainerScrollMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ScrollContainerScrollMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_get_vertical_scroll_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDeadzone(deadzone gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, deadzone)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_set_deadzone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDeadzone() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_get_deadzone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFollowFocus(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_set_follow_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsFollowingFocus() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_is_following_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the horizontal scrollbar [HScrollBar] of this [ScrollContainer].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to disable or hide a scrollbar, you can use [member horizontal_scroll_mode].
*/
//go:nosplit
func (self class) GetHScrollBar() objects.HScrollBar {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_get_h_scroll_bar, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.HScrollBar{classdb.HScrollBar(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the vertical scrollbar [VScrollBar] of this [ScrollContainer].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to disable or hide a scrollbar, you can use [member vertical_scroll_mode].
*/
//go:nosplit
func (self class) GetVScrollBar() objects.VScrollBar {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_get_v_scroll_bar, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.VScrollBar{classdb.VScrollBar(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
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
func (self class) EnsureControlVisible(control objects.Control) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(control[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_ensure_control_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Instance) OnScrollStarted(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("scroll_started"), gd.NewCallable(cb), 0)
}

func (self Instance) OnScrollEnded(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("scroll_ended"), gd.NewCallable(cb), 0)
}

func (self class) AsScrollContainer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsScrollContainer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	default:
		return gd.VirtualByName(self.AsContainer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsContainer(), name)
	}
}
func init() {
	classdb.Register("ScrollContainer", func(ptr gd.Object) any { return classdb.ScrollContainer(ptr) })
}

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
