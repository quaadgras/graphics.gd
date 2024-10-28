package Parallax2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node2D"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
A [Parallax2D] is used to create a parallax effect. It can move at a different speed relative to the camera movement using [member scroll_scale]. This creates an illusion of depth in a 2D game. If manual scrolling is desired, the [Camera2D] position can be ignored with [member ignore_camera_scroll].
[b]Note:[/b] Any changes to this node's position made after it enters the scene tree will be overridden if [member ignore_camera_scroll] is [code]false[/code] or [member screen_offset] is modified.

*/
type Go [1]classdb.Parallax2D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Parallax2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Parallax2D"))
	return Go{classdb.Parallax2D(object)}
}

func (self Go) ScrollScale() gd.Vector2 {
		return gd.Vector2(class(self).GetScrollScale())
}

func (self Go) SetScrollScale(value gd.Vector2) {
	class(self).SetScrollScale(value)
}

func (self Go) ScrollOffset() gd.Vector2 {
		return gd.Vector2(class(self).GetScrollOffset())
}

func (self Go) SetScrollOffset(value gd.Vector2) {
	class(self).SetScrollOffset(value)
}

func (self Go) RepeatSize() gd.Vector2 {
		return gd.Vector2(class(self).GetRepeatSize())
}

func (self Go) SetRepeatSize(value gd.Vector2) {
	class(self).SetRepeatSize(value)
}

func (self Go) Autoscroll() gd.Vector2 {
		return gd.Vector2(class(self).GetAutoscroll())
}

func (self Go) SetAutoscroll(value gd.Vector2) {
	class(self).SetAutoscroll(value)
}

func (self Go) RepeatTimes() int {
		return int(int(class(self).GetRepeatTimes()))
}

func (self Go) SetRepeatTimes(value int) {
	class(self).SetRepeatTimes(gd.Int(value))
}

func (self Go) LimitBegin() gd.Vector2 {
		return gd.Vector2(class(self).GetLimitBegin())
}

func (self Go) SetLimitBegin(value gd.Vector2) {
	class(self).SetLimitBegin(value)
}

func (self Go) LimitEnd() gd.Vector2 {
		return gd.Vector2(class(self).GetLimitEnd())
}

func (self Go) SetLimitEnd(value gd.Vector2) {
	class(self).SetLimitEnd(value)
}

func (self Go) FollowViewport() bool {
		return bool(class(self).GetFollowViewport())
}

func (self Go) SetFollowViewport(value bool) {
	class(self).SetFollowViewport(value)
}

func (self Go) IgnoreCameraScroll() bool {
		return bool(class(self).IsIgnoreCameraScroll())
}

func (self Go) SetIgnoreCameraScroll(value bool) {
	class(self).SetIgnoreCameraScroll(value)
}

func (self Go) ScreenOffset() gd.Vector2 {
		return gd.Vector2(class(self).GetScreenOffset())
}

func (self Go) SetScreenOffset(value gd.Vector2) {
	class(self).SetScreenOffset(value)
}

//go:nosplit
func (self class) SetScrollScale(scale gd.Vector2)  {
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
func (self class) SetRepeatSize(repeat_size gd.Vector2)  {
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
func (self class) SetRepeatTimes(repeat_times gd.Int)  {
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
func (self class) SetAutoscroll(autoscroll gd.Vector2)  {
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
func (self class) SetScrollOffset(offset gd.Vector2)  {
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
func (self class) SetScreenOffset(offset gd.Vector2)  {
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
func (self class) SetLimitBegin(offset gd.Vector2)  {
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
func (self class) SetLimitEnd(offset gd.Vector2)  {
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
func (self class) SetFollowViewport(follow bool)  {
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
func (self class) SetIgnoreCameraScroll(ignore bool)  {
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
func (self class) AsParallax2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsParallax2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.GD { return *((*Node2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode2D() Node2D.Go { return *((*Node2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {classdb.Register("Parallax2D", func(ptr gd.Object) any { return classdb.Parallax2D(ptr) })}
