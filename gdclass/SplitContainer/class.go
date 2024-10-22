package SplitContainer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Container"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A container that accepts only two child controls, then arranges them horizontally or vertically and creates a divisor between them. The divisor can be dragged around to change the size relation between the child controls.

*/
type Go [1]classdb.SplitContainer

/*
Clamps the [member split_offset] value to not go outside the currently possible minimal and maximum values.
*/
func (self Go) ClampSplitOffset() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClampSplitOffset()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SplitContainer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SplitContainer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) SplitOffset() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSplitOffset()))
}

func (self Go) SetSplitOffset(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSplitOffset(gd.Int(value))
}

func (self Go) Collapsed() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsCollapsed())
}

func (self Go) SetCollapsed(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollapsed(value)
}

func (self Go) DraggerVisibility() classdb.SplitContainerDraggerVisibility {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.SplitContainerDraggerVisibility(class(self).GetDraggerVisibility())
}

func (self Go) SetDraggerVisibility(value classdb.SplitContainerDraggerVisibility) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDraggerVisibility(value)
}

func (self Go) Vertical() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsVertical())
}

func (self Go) SetVertical(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVertical(value)
}

//go:nosplit
func (self class) SetSplitOffset(offset gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SplitContainer.Bind_set_split_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSplitOffset() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SplitContainer.Bind_get_split_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Clamps the [member split_offset] value to not go outside the currently possible minimal and maximum values.
*/
//go:nosplit
func (self class) ClampSplitOffset()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SplitContainer.Bind_clamp_split_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCollapsed(collapsed bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collapsed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SplitContainer.Bind_set_collapsed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCollapsed() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SplitContainer.Bind_is_collapsed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDraggerVisibility(mode classdb.SplitContainerDraggerVisibility)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SplitContainer.Bind_set_dragger_visibility, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDraggerVisibility() classdb.SplitContainerDraggerVisibility {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.SplitContainerDraggerVisibility](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SplitContainer.Bind_get_dragger_visibility, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVertical(vertical bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, vertical)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SplitContainer.Bind_set_vertical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsVertical() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SplitContainer.Bind_is_vertical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnDragged(cb func(offset int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("dragged"), gc.Callable(cb), 0)
}


func (self class) AsSplitContainer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSplitContainer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsContainer() Container.GD { return *((*Container.GD)(unsafe.Pointer(&self))) }
func (self Go) AsContainer() Container.Go { return *((*Container.Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsContainer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsContainer(), name)
	}
}
func init() {classdb.Register("SplitContainer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type DraggerVisibility = classdb.SplitContainerDraggerVisibility

const (
/*The split dragger is visible when the cursor hovers it.*/
	DraggerVisible DraggerVisibility = 0
/*The split dragger is never visible.*/
	DraggerHidden DraggerVisibility = 1
/*The split dragger is never visible and its space collapsed.*/
	DraggerHiddenCollapsed DraggerVisibility = 2
)
