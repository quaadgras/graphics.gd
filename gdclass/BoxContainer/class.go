package BoxContainer

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
A container that arranges its child controls horizontally or vertically, rearranging them automatically when their minimum size changes.

*/
type Go [1]classdb.BoxContainer

/*
Adds a [Control] node to the box as a spacer. If [param begin] is [code]true[/code], it will insert the [Control] node in front of all other children.
*/
func (self Go) AddSpacer(begin bool) gdclass.Control {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Control(class(self).AddSpacer(gc, begin))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.BoxContainer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("BoxContainer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Alignment() classdb.BoxContainerAlignmentMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.BoxContainerAlignmentMode(class(self).GetAlignment())
}

func (self Go) SetAlignment(value classdb.BoxContainerAlignmentMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAlignment(value)
}

func (self Go) Vertical() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsVertical())
}

func (self Go) SetVertical(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVertical(value)
}

/*
Adds a [Control] node to the box as a spacer. If [param begin] is [code]true[/code], it will insert the [Control] node in front of all other children.
*/
//go:nosplit
func (self class) AddSpacer(ctx gd.Lifetime, begin bool) gdclass.Control {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, begin)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoxContainer.Bind_add_spacer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Control
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlignment(alignment classdb.BoxContainerAlignmentMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoxContainer.Bind_set_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlignment() classdb.BoxContainerAlignmentMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BoxContainerAlignmentMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoxContainer.Bind_get_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoxContainer.Bind_set_vertical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsVertical() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoxContainer.Bind_is_vertical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsBoxContainer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsBoxContainer() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("BoxContainer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type AlignmentMode = classdb.BoxContainerAlignmentMode

const (
/*The child controls will be arranged at the beginning of the container, i.e. top if orientation is vertical, left if orientation is horizontal (right for RTL layout).*/
	AlignmentBegin AlignmentMode = 0
/*The child controls will be centered in the container.*/
	AlignmentCenter AlignmentMode = 1
/*The child controls will be arranged at the end of the container, i.e. bottom if orientation is vertical, right if orientation is horizontal (left for RTL layout).*/
	AlignmentEnd AlignmentMode = 2
)
