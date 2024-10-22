package FlowContainer

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
A container that arranges its child controls horizontally or vertically and wraps them around at the borders. This is similar to how text in a book wraps around when no more words can fit on a line.

*/
type Go [1]classdb.FlowContainer

/*
Returns the current line count.
*/
func (self Go) GetLineCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetLineCount()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.FlowContainer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("FlowContainer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Alignment() classdb.FlowContainerAlignmentMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.FlowContainerAlignmentMode(class(self).GetAlignment())
}

func (self Go) SetAlignment(value classdb.FlowContainerAlignmentMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAlignment(value)
}

func (self Go) LastWrapAlignment() classdb.FlowContainerLastWrapAlignmentMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.FlowContainerLastWrapAlignmentMode(class(self).GetLastWrapAlignment())
}

func (self Go) SetLastWrapAlignment(value classdb.FlowContainerLastWrapAlignmentMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLastWrapAlignment(value)
}

func (self Go) Vertical() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsVertical())
}

func (self Go) SetVertical(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVertical(value)
}

func (self Go) ReverseFill() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsReverseFill())
}

func (self Go) SetReverseFill(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetReverseFill(value)
}

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
func (self class) AsFlowContainer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsFlowContainer() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("FlowContainer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
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
