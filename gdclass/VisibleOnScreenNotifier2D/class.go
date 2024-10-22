package VisibleOnScreenNotifier2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
var _ mmm.Lifetime

/*
[VisibleOnScreenNotifier2D] represents a rectangular region of 2D space. When any part of this region becomes visible on screen or in a viewport, it will emit a [signal screen_entered] signal, and likewise it will emit a [signal screen_exited] signal when no part of it remains visible.
If you want a node to be enabled automatically when this region is visible on screen, use [VisibleOnScreenEnabler2D].
[b]Note:[/b] [VisibleOnScreenNotifier2D] uses the render culling code to determine whether it's visible on screen, so it won't function unless [member CanvasItem.visible] is set to [code]true[/code].

*/
type Go [1]classdb.VisibleOnScreenNotifier2D

/*
If [code]true[/code], the bounding rectangle is on the screen.
[b]Note:[/b] It takes one frame for the [VisibleOnScreenNotifier2D]'s visibility to be determined once added to the scene tree, so this method will always return [code]false[/code] right after it is instantiated, before the draw pass.
*/
func (self Go) IsOnScreen() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsOnScreen())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisibleOnScreenNotifier2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VisibleOnScreenNotifier2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Rect() gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Rect2(class(self).GetRect())
}

func (self Go) SetRect(value gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRect(value)
}

//go:nosplit
func (self class) SetRect(rect gd.Rect2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisibleOnScreenNotifier2D.Bind_set_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRect() gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisibleOnScreenNotifier2D.Bind_get_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], the bounding rectangle is on the screen.
[b]Note:[/b] It takes one frame for the [VisibleOnScreenNotifier2D]'s visibility to be determined once added to the scene tree, so this method will always return [code]false[/code] right after it is instantiated, before the draw pass.
*/
//go:nosplit
func (self class) IsOnScreen() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisibleOnScreenNotifier2D.Bind_is_on_screen, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnScreenEntered(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("screen_entered"), gc.Callable(cb), 0)
}


func (self Go) OnScreenExited(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("screen_exited"), gc.Callable(cb), 0)
}


func (self class) AsVisibleOnScreenNotifier2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisibleOnScreenNotifier2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("VisibleOnScreenNotifier2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
