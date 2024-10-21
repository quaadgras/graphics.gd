package SplitContainer

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
A container that accepts only two child controls, then arranges them horizontally or vertically and creates a divisor between them. The divisor can be dragged around to change the size relation between the child controls.

*/
type Simple [1]classdb.SplitContainer
func (self Simple) SetSplitOffset(offset int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSplitOffset(gd.Int(offset))
}
func (self Simple) GetSplitOffset() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSplitOffset()))
}
func (self Simple) ClampSplitOffset() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClampSplitOffset()
}
func (self Simple) SetCollapsed(collapsed bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCollapsed(collapsed)
}
func (self Simple) IsCollapsed() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsCollapsed())
}
func (self Simple) SetDraggerVisibility(mode classdb.SplitContainerDraggerVisibility) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDraggerVisibility(mode)
}
func (self Simple) GetDraggerVisibility() classdb.SplitContainerDraggerVisibility {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.SplitContainerDraggerVisibility(Expert(self).GetDraggerVisibility())
}
func (self Simple) SetVertical(vertical bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVertical(vertical)
}
func (self Simple) IsVertical() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsVertical())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.SplitContainer
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsSplitContainer() Expert { return self[0].AsSplitContainer() }


//go:nosplit
func (self Simple) AsSplitContainer() Simple { return self[0].AsSplitContainer() }


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
func init() {classdb.Register("SplitContainer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type DraggerVisibility = classdb.SplitContainerDraggerVisibility

const (
/*The split dragger is visible when the cursor hovers it.*/
	DraggerVisible DraggerVisibility = 0
/*The split dragger is never visible.*/
	DraggerHidden DraggerVisibility = 1
/*The split dragger is never visible and its space collapsed.*/
	DraggerHiddenCollapsed DraggerVisibility = 2
)
