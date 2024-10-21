package AspectRatioContainer

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
A container type that arranges its child controls in a way that preserves their proportions automatically when the container is resized. Useful when a container has a dynamic size and the child nodes must adjust their sizes accordingly without losing their aspect ratios.

*/
type Simple [1]classdb.AspectRatioContainer
func (self Simple) SetRatio(ratio float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRatio(gd.Float(ratio))
}
func (self Simple) GetRatio() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetRatio()))
}
func (self Simple) SetStretchMode(stretch_mode classdb.AspectRatioContainerStretchMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStretchMode(stretch_mode)
}
func (self Simple) GetStretchMode() classdb.AspectRatioContainerStretchMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AspectRatioContainerStretchMode(Expert(self).GetStretchMode())
}
func (self Simple) SetAlignmentHorizontal(alignment_horizontal classdb.AspectRatioContainerAlignmentMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAlignmentHorizontal(alignment_horizontal)
}
func (self Simple) GetAlignmentHorizontal() classdb.AspectRatioContainerAlignmentMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AspectRatioContainerAlignmentMode(Expert(self).GetAlignmentHorizontal())
}
func (self Simple) SetAlignmentVertical(alignment_vertical classdb.AspectRatioContainerAlignmentMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAlignmentVertical(alignment_vertical)
}
func (self Simple) GetAlignmentVertical() classdb.AspectRatioContainerAlignmentMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AspectRatioContainerAlignmentMode(Expert(self).GetAlignmentVertical())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AspectRatioContainer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetRatio(ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AspectRatioContainer.Bind_set_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRatio() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AspectRatioContainer.Bind_get_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStretchMode(stretch_mode classdb.AspectRatioContainerStretchMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stretch_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AspectRatioContainer.Bind_set_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStretchMode() classdb.AspectRatioContainerStretchMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AspectRatioContainerStretchMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AspectRatioContainer.Bind_get_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlignmentHorizontal(alignment_horizontal classdb.AspectRatioContainerAlignmentMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, alignment_horizontal)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AspectRatioContainer.Bind_set_alignment_horizontal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlignmentHorizontal() classdb.AspectRatioContainerAlignmentMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AspectRatioContainerAlignmentMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AspectRatioContainer.Bind_get_alignment_horizontal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlignmentVertical(alignment_vertical classdb.AspectRatioContainerAlignmentMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, alignment_vertical)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AspectRatioContainer.Bind_set_alignment_vertical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlignmentVertical() classdb.AspectRatioContainerAlignmentMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AspectRatioContainerAlignmentMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AspectRatioContainer.Bind_get_alignment_vertical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsAspectRatioContainer() Expert { return self[0].AsAspectRatioContainer() }


//go:nosplit
func (self Simple) AsAspectRatioContainer() Simple { return self[0].AsAspectRatioContainer() }


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
func init() {classdb.Register("AspectRatioContainer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type StretchMode = classdb.AspectRatioContainerStretchMode

const (
/*The height of child controls is automatically adjusted based on the width of the container.*/
	StretchWidthControlsHeight StretchMode = 0
/*The width of child controls is automatically adjusted based on the height of the container.*/
	StretchHeightControlsWidth StretchMode = 1
/*The bounding rectangle of child controls is automatically adjusted to fit inside the container while keeping the aspect ratio.*/
	StretchFit StretchMode = 2
/*The width and height of child controls is automatically adjusted to make their bounding rectangle cover the entire area of the container while keeping the aspect ratio.
When the bounding rectangle of child controls exceed the container's size and [member Control.clip_contents] is enabled, this allows to show only the container's area restricted by its own bounding rectangle.*/
	StretchCover StretchMode = 3
)
type AlignmentMode = classdb.AspectRatioContainerAlignmentMode

const (
/*Aligns child controls with the beginning (left or top) of the container.*/
	AlignmentBegin AlignmentMode = 0
/*Aligns child controls with the center of the container.*/
	AlignmentCenter AlignmentMode = 1
/*Aligns child controls with the end (right or bottom) of the container.*/
	AlignmentEnd AlignmentMode = 2
)
