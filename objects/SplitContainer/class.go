package SplitContainer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Container"
import "grow.graphics/gd/objects/Control"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
A container that accepts only two child controls, then arranges them horizontally or vertically and creates a divisor between them. The divisor can be dragged around to change the size relation between the child controls.
*/
type Instance [1]classdb.SplitContainer

/*
Clamps the [member split_offset] value to not go outside the currently possible minimal and maximum values.
*/
func (self Instance) ClampSplitOffset() {
	class(self).ClampSplitOffset()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.SplitContainer

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SplitContainer"))
	return Instance{classdb.SplitContainer(object)}
}

func (self Instance) SplitOffset() int {
	return int(int(class(self).GetSplitOffset()))
}

func (self Instance) SetSplitOffset(value int) {
	class(self).SetSplitOffset(gd.Int(value))
}

func (self Instance) Collapsed() bool {
	return bool(class(self).IsCollapsed())
}

func (self Instance) SetCollapsed(value bool) {
	class(self).SetCollapsed(value)
}

func (self Instance) DraggerVisibility() classdb.SplitContainerDraggerVisibility {
	return classdb.SplitContainerDraggerVisibility(class(self).GetDraggerVisibility())
}

func (self Instance) SetDraggerVisibility(value classdb.SplitContainerDraggerVisibility) {
	class(self).SetDraggerVisibility(value)
}

func (self Instance) Vertical() bool {
	return bool(class(self).IsVertical())
}

func (self Instance) SetVertical(value bool) {
	class(self).SetVertical(value)
}

//go:nosplit
func (self class) SetSplitOffset(offset gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SplitContainer.Bind_set_split_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSplitOffset() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SplitContainer.Bind_get_split_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clamps the [member split_offset] value to not go outside the currently possible minimal and maximum values.
*/
//go:nosplit
func (self class) ClampSplitOffset() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SplitContainer.Bind_clamp_split_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetCollapsed(collapsed bool) {
	var frame = callframe.New()
	callframe.Arg(frame, collapsed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SplitContainer.Bind_set_collapsed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsCollapsed() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SplitContainer.Bind_is_collapsed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDraggerVisibility(mode classdb.SplitContainerDraggerVisibility) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SplitContainer.Bind_set_dragger_visibility, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDraggerVisibility() classdb.SplitContainerDraggerVisibility {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.SplitContainerDraggerVisibility](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SplitContainer.Bind_get_dragger_visibility, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVertical(vertical bool) {
	var frame = callframe.New()
	callframe.Arg(frame, vertical)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SplitContainer.Bind_set_vertical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsVertical() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SplitContainer.Bind_is_vertical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnDragged(cb func(offset int)) {
	self[0].AsObject().Connect(gd.NewStringName("dragged"), gd.NewCallable(cb), 0)
}

func (self class) AsSplitContainer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSplitContainer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsContainer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsContainer(), name)
	}
}
func init() {
	classdb.Register("SplitContainer", func(ptr gd.Object) any { return classdb.SplitContainer(ptr) })
}

type DraggerVisibility = classdb.SplitContainerDraggerVisibility

const (
	/*The split dragger is visible when the cursor hovers it.*/
	DraggerVisible DraggerVisibility = 0
	/*The split dragger is never visible.*/
	DraggerHidden DraggerVisibility = 1
	/*The split dragger is never visible and its space collapsed.*/
	DraggerHiddenCollapsed DraggerVisibility = 2
)
