package Parallax2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Node2D"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Vector2"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A [Parallax2D] is used to create a parallax effect. It can move at a different speed relative to the camera movement using [member scroll_scale]. This creates an illusion of depth in a 2D game. If manual scrolling is desired, the [Camera2D] position can be ignored with [member ignore_camera_scroll].
[b]Note:[/b] Any changes to this node's position made after it enters the scene tree will be overridden if [member ignore_camera_scroll] is [code]false[/code] or [member screen_offset] is modified.
*/
type Instance [1]classdb.Parallax2D
type Any interface {
	gd.IsClass
	AsParallax2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Parallax2D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Parallax2D"))
	return Instance{*(*classdb.Parallax2D)(unsafe.Pointer(&object))}
}

func (self Instance) ScrollScale() Vector2.XY {
	return Vector2.XY(class(self).GetScrollScale())
}

func (self Instance) SetScrollScale(value Vector2.XY) {
	class(self).SetScrollScale(gd.Vector2(value))
}

func (self Instance) ScrollOffset() Vector2.XY {
	return Vector2.XY(class(self).GetScrollOffset())
}

func (self Instance) SetScrollOffset(value Vector2.XY) {
	class(self).SetScrollOffset(gd.Vector2(value))
}

func (self Instance) RepeatSize() Vector2.XY {
	return Vector2.XY(class(self).GetRepeatSize())
}

func (self Instance) SetRepeatSize(value Vector2.XY) {
	class(self).SetRepeatSize(gd.Vector2(value))
}

func (self Instance) Autoscroll() Vector2.XY {
	return Vector2.XY(class(self).GetAutoscroll())
}

func (self Instance) SetAutoscroll(value Vector2.XY) {
	class(self).SetAutoscroll(gd.Vector2(value))
}

func (self Instance) RepeatTimes() int {
	return int(int(class(self).GetRepeatTimes()))
}

func (self Instance) SetRepeatTimes(value int) {
	class(self).SetRepeatTimes(gd.Int(value))
}

func (self Instance) LimitBegin() Vector2.XY {
	return Vector2.XY(class(self).GetLimitBegin())
}

func (self Instance) SetLimitBegin(value Vector2.XY) {
	class(self).SetLimitBegin(gd.Vector2(value))
}

func (self Instance) LimitEnd() Vector2.XY {
	return Vector2.XY(class(self).GetLimitEnd())
}

func (self Instance) SetLimitEnd(value Vector2.XY) {
	class(self).SetLimitEnd(gd.Vector2(value))
}

func (self Instance) FollowViewport() bool {
	return bool(class(self).GetFollowViewport())
}

func (self Instance) SetFollowViewport(value bool) {
	class(self).SetFollowViewport(value)
}

func (self Instance) IgnoreCameraScroll() bool {
	return bool(class(self).IsIgnoreCameraScroll())
}

func (self Instance) SetIgnoreCameraScroll(value bool) {
	class(self).SetIgnoreCameraScroll(value)
}

func (self Instance) ScreenOffset() Vector2.XY {
	return Vector2.XY(class(self).GetScreenOffset())
}

func (self Instance) SetScreenOffset(value Vector2.XY) {
	class(self).SetScreenOffset(gd.Vector2(value))
}

//go:nosplit
func (self class) SetScrollScale(scale gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_set_scroll_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScrollScale() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_get_scroll_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRepeatSize(repeat_size gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, repeat_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_set_repeat_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRepeatSize() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_get_repeat_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRepeatTimes(repeat_times gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, repeat_times)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_set_repeat_times, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRepeatTimes() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_get_repeat_times, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoscroll(autoscroll gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, autoscroll)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_set_autoscroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoscroll() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_get_autoscroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScrollOffset(offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_set_scroll_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScrollOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_get_scroll_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScreenOffset(offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_set_screen_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScreenOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_get_screen_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLimitBegin(offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_set_limit_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLimitBegin() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_get_limit_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLimitEnd(offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_set_limit_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLimitEnd() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_get_limit_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFollowViewport(follow bool) {
	var frame = callframe.New()
	callframe.Arg(frame, follow)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_set_follow_viewport, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFollowViewport() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_get_follow_viewport, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIgnoreCameraScroll(ignore bool) {
	var frame = callframe.New()
	callframe.Arg(frame, ignore)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_set_ignore_camera_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsIgnoreCameraScroll() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Parallax2D.Bind_is_ignore_camera_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsParallax2D() Advanced       { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsParallax2D() Instance    { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced    { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {
	classdb.Register("Parallax2D", func(ptr gd.Object) any { return [1]classdb.Parallax2D{*(*classdb.Parallax2D)(unsafe.Pointer(&ptr))} })
}
