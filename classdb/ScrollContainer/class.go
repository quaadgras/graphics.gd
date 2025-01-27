// Package ScrollContainer provides methods for working with ScrollContainer object instances.
package ScrollContainer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/classdb/Container"
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
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable

/*
A container used to provide a child control with scrollbars when needed. Scrollbars will automatically be drawn at the right (for vertical) or bottom (for horizontal) and will enable dragging to move the viewable Control (and its children) within the ScrollContainer. Scrollbars will also automatically resize the grabber based on the [member Control.custom_minimum_size] of the Control relative to the ScrollContainer.
*/
type Instance [1]gdclass.ScrollContainer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsScrollContainer() Instance
}

/*
Returns the horizontal scrollbar [HScrollBar] of this [ScrollContainer].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to disable or hide a scrollbar, you can use [member horizontal_scroll_mode].
*/
func (self Instance) GetHScrollBar() [1]gdclass.HScrollBar { //gd:ScrollContainer.get_h_scroll_bar
	return [1]gdclass.HScrollBar(class(self).GetHScrollBar())
}

/*
Returns the vertical scrollbar [VScrollBar] of this [ScrollContainer].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to disable or hide a scrollbar, you can use [member vertical_scroll_mode].
*/
func (self Instance) GetVScrollBar() [1]gdclass.VScrollBar { //gd:ScrollContainer.get_v_scroll_bar
	return [1]gdclass.VScrollBar(class(self).GetVScrollBar())
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
func (self Instance) EnsureControlVisible(control [1]gdclass.Control) { //gd:ScrollContainer.ensure_control_visible
	class(self).EnsureControlVisible(control)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ScrollContainer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ScrollContainer"))
	casted := Instance{*(*gdclass.ScrollContainer)(unsafe.Pointer(&object))}
	return casted
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

func (self Instance) HorizontalScrollMode() gdclass.ScrollContainerScrollMode {
	return gdclass.ScrollContainerScrollMode(class(self).GetHorizontalScrollMode())
}

func (self Instance) SetHorizontalScrollMode(value gdclass.ScrollContainerScrollMode) {
	class(self).SetHorizontalScrollMode(value)
}

func (self Instance) VerticalScrollMode() gdclass.ScrollContainerScrollMode {
	return gdclass.ScrollContainerScrollMode(class(self).GetVerticalScrollMode())
}

func (self Instance) SetVerticalScrollMode(value gdclass.ScrollContainerScrollMode) {
	class(self).SetVerticalScrollMode(value)
}

func (self Instance) ScrollDeadzone() int {
	return int(int(class(self).GetDeadzone()))
}

func (self Instance) SetScrollDeadzone(value int) {
	class(self).SetDeadzone(gd.Int(value))
}

//go:nosplit
func (self class) SetHScroll(value gd.Int) { //gd:ScrollContainer.set_h_scroll
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_set_h_scroll, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHScroll() gd.Int { //gd:ScrollContainer.get_h_scroll
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_get_h_scroll, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVScroll(value gd.Int) { //gd:ScrollContainer.set_v_scroll
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_set_v_scroll, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVScroll() gd.Int { //gd:ScrollContainer.get_v_scroll
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_get_v_scroll, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHorizontalCustomStep(value gd.Float) { //gd:ScrollContainer.set_horizontal_custom_step
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_set_horizontal_custom_step, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHorizontalCustomStep() gd.Float { //gd:ScrollContainer.get_horizontal_custom_step
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_get_horizontal_custom_step, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVerticalCustomStep(value gd.Float) { //gd:ScrollContainer.set_vertical_custom_step
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_set_vertical_custom_step, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVerticalCustomStep() gd.Float { //gd:ScrollContainer.get_vertical_custom_step
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_get_vertical_custom_step, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHorizontalScrollMode(enable gdclass.ScrollContainerScrollMode) { //gd:ScrollContainer.set_horizontal_scroll_mode
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_set_horizontal_scroll_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHorizontalScrollMode() gdclass.ScrollContainerScrollMode { //gd:ScrollContainer.get_horizontal_scroll_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.ScrollContainerScrollMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_get_horizontal_scroll_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVerticalScrollMode(enable gdclass.ScrollContainerScrollMode) { //gd:ScrollContainer.set_vertical_scroll_mode
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_set_vertical_scroll_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVerticalScrollMode() gdclass.ScrollContainerScrollMode { //gd:ScrollContainer.get_vertical_scroll_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.ScrollContainerScrollMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_get_vertical_scroll_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDeadzone(deadzone gd.Int) { //gd:ScrollContainer.set_deadzone
	var frame = callframe.New()
	callframe.Arg(frame, deadzone)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_set_deadzone, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDeadzone() gd.Int { //gd:ScrollContainer.get_deadzone
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_get_deadzone, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFollowFocus(enabled bool) { //gd:ScrollContainer.set_follow_focus
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_set_follow_focus, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFollowingFocus() bool { //gd:ScrollContainer.is_following_focus
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_is_following_focus, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the horizontal scrollbar [HScrollBar] of this [ScrollContainer].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to disable or hide a scrollbar, you can use [member horizontal_scroll_mode].
*/
//go:nosplit
func (self class) GetHScrollBar() [1]gdclass.HScrollBar { //gd:ScrollContainer.get_h_scroll_bar
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_get_h_scroll_bar, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.HScrollBar{gd.PointerLifetimeBoundTo[gdclass.HScrollBar](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the vertical scrollbar [VScrollBar] of this [ScrollContainer].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to disable or hide a scrollbar, you can use [member vertical_scroll_mode].
*/
//go:nosplit
func (self class) GetVScrollBar() [1]gdclass.VScrollBar { //gd:ScrollContainer.get_v_scroll_bar
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_get_v_scroll_bar, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.VScrollBar{gd.PointerLifetimeBoundTo[gdclass.VScrollBar](self.AsObject(), r_ret.Get())}
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
func (self class) EnsureControlVisible(control [1]gdclass.Control) { //gd:ScrollContainer.ensure_control_visible
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(control[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScrollContainer.Bind_ensure_control_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnScrollStarted(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("scroll_started"), gd.NewCallable(cb), 0)
}

func (self Instance) OnScrollEnded(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("scroll_ended"), gd.NewCallable(cb), 0)
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
		return gd.VirtualByName(Container.Advanced(self.AsContainer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Container.Instance(self.AsContainer()), name)
	}
}
func init() {
	gdclass.Register("ScrollContainer", func(ptr gd.Object) any {
		return [1]gdclass.ScrollContainer{*(*gdclass.ScrollContainer)(unsafe.Pointer(&ptr))}
	})
}

type ScrollMode = gdclass.ScrollContainerScrollMode //gd:ScrollContainer.ScrollMode

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
