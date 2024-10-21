package FlowContainer

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
A container that arranges its child controls horizontally or vertically and wraps them around at the borders. This is similar to how text in a book wraps around when no more words can fit on a line.

*/
type Simple [1]classdb.FlowContainer
func (self Simple) GetLineCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetLineCount()))
}
func (self Simple) SetAlignment(alignment classdb.FlowContainerAlignmentMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAlignment(alignment)
}
func (self Simple) GetAlignment() classdb.FlowContainerAlignmentMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.FlowContainerAlignmentMode(Expert(self).GetAlignment())
}
func (self Simple) SetLastWrapAlignment(last_wrap_alignment classdb.FlowContainerLastWrapAlignmentMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLastWrapAlignment(last_wrap_alignment)
}
func (self Simple) GetLastWrapAlignment() classdb.FlowContainerLastWrapAlignmentMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.FlowContainerLastWrapAlignmentMode(Expert(self).GetLastWrapAlignment())
}
func (self Simple) SetVertical(vertical bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVertical(vertical)
}
func (self Simple) IsVertical() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsVertical())
}
func (self Simple) SetReverseFill(reverse_fill bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetReverseFill(reverse_fill)
}
func (self Simple) IsReverseFill() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsReverseFill())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.FlowContainer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the current line count.
*/
//go:nosplit
func (self class) GetLineCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FlowContainer.Bind_get_line_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlignment(alignment classdb.FlowContainerAlignmentMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FlowContainer.Bind_set_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlignment() classdb.FlowContainerAlignmentMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.FlowContainerAlignmentMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FlowContainer.Bind_get_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLastWrapAlignment(last_wrap_alignment classdb.FlowContainerLastWrapAlignmentMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, last_wrap_alignment)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FlowContainer.Bind_set_last_wrap_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLastWrapAlignment() classdb.FlowContainerLastWrapAlignmentMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.FlowContainerLastWrapAlignmentMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FlowContainer.Bind_get_last_wrap_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FlowContainer.Bind_set_vertical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsVertical() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FlowContainer.Bind_is_vertical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetReverseFill(reverse_fill bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, reverse_fill)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FlowContainer.Bind_set_reverse_fill, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsReverseFill() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FlowContainer.Bind_is_reverse_fill, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsFlowContainer() Expert { return self[0].AsFlowContainer() }


//go:nosplit
func (self Simple) AsFlowContainer() Simple { return self[0].AsFlowContainer() }


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
func init() {classdb.Register("FlowContainer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type AlignmentMode = classdb.FlowContainerAlignmentMode

const (
/*The child controls will be arranged at the beginning of the container, i.e. top if orientation is vertical, left if orientation is horizontal (right for RTL layout).*/
	AlignmentBegin AlignmentMode = 0
/*The child controls will be centered in the container.*/
	AlignmentCenter AlignmentMode = 1
/*The child controls will be arranged at the end of the container, i.e. bottom if orientation is vertical, right if orientation is horizontal (left for RTL layout).*/
	AlignmentEnd AlignmentMode = 2
)
type LastWrapAlignmentMode = classdb.FlowContainerLastWrapAlignmentMode

const (
/*The last partially filled row or column will wrap aligned to the previous row or column in accordance with [member alignment].*/
	LastWrapAlignmentInherit LastWrapAlignmentMode = 0
/*The last partially filled row or column will wrap aligned to the beginning of the previous row or column.*/
	LastWrapAlignmentBegin LastWrapAlignmentMode = 1
/*The last partially filled row or column will wrap aligned to the center of the previous row or column.*/
	LastWrapAlignmentCenter LastWrapAlignmentMode = 2
/*The last partially filled row or column will wrap aligned to the end of the previous row or column.*/
	LastWrapAlignmentEnd LastWrapAlignmentMode = 3
)
