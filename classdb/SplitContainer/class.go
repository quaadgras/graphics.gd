// Package SplitContainer provides methods for working with SplitContainer object instances.
package SplitContainer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/classdb/Container"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any

/*
A container that accepts only two child controls, then arranges them horizontally or vertically and creates a divisor between them. The divisor can be dragged around to change the size relation between the child controls.
*/
type Instance [1]gdclass.SplitContainer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSplitContainer() Instance
}

/*
Clamps the [member split_offset] value to not go outside the currently possible minimal and maximum values.
*/
func (self Instance) ClampSplitOffset() { //gd:SplitContainer.clamp_split_offset
	class(self).ClampSplitOffset()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SplitContainer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SplitContainer"))
	casted := Instance{*(*gdclass.SplitContainer)(unsafe.Pointer(&object))}
	return casted
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

func (self Instance) DraggerVisibility() gdclass.SplitContainerDraggerVisibility {
	return gdclass.SplitContainerDraggerVisibility(class(self).GetDraggerVisibility())
}

func (self Instance) SetDraggerVisibility(value gdclass.SplitContainerDraggerVisibility) {
	class(self).SetDraggerVisibility(value)
}

func (self Instance) Vertical() bool {
	return bool(class(self).IsVertical())
}

func (self Instance) SetVertical(value bool) {
	class(self).SetVertical(value)
}

//go:nosplit
func (self class) SetSplitOffset(offset gd.Int) { //gd:SplitContainer.set_split_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SplitContainer.Bind_set_split_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSplitOffset() gd.Int { //gd:SplitContainer.get_split_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SplitContainer.Bind_get_split_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clamps the [member split_offset] value to not go outside the currently possible minimal and maximum values.
*/
//go:nosplit
func (self class) ClampSplitOffset() { //gd:SplitContainer.clamp_split_offset
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SplitContainer.Bind_clamp_split_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetCollapsed(collapsed bool) { //gd:SplitContainer.set_collapsed
	var frame = callframe.New()
	callframe.Arg(frame, collapsed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SplitContainer.Bind_set_collapsed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCollapsed() bool { //gd:SplitContainer.is_collapsed
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SplitContainer.Bind_is_collapsed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDraggerVisibility(mode gdclass.SplitContainerDraggerVisibility) { //gd:SplitContainer.set_dragger_visibility
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SplitContainer.Bind_set_dragger_visibility, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDraggerVisibility() gdclass.SplitContainerDraggerVisibility { //gd:SplitContainer.get_dragger_visibility
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.SplitContainerDraggerVisibility](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SplitContainer.Bind_get_dragger_visibility, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVertical(vertical bool) { //gd:SplitContainer.set_vertical
	var frame = callframe.New()
	callframe.Arg(frame, vertical)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SplitContainer.Bind_set_vertical, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsVertical() bool { //gd:SplitContainer.is_vertical
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SplitContainer.Bind_is_vertical, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnDragged(cb func(offset int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("dragged"), gd.NewCallable(cb), 0)
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
		return gd.VirtualByName(Container.Advanced(self.AsContainer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Container.Instance(self.AsContainer()), name)
	}
}
func init() {
	gdclass.Register("SplitContainer", func(ptr gd.Object) any {
		return [1]gdclass.SplitContainer{*(*gdclass.SplitContainer)(unsafe.Pointer(&ptr))}
	})
}

type DraggerVisibility = gdclass.SplitContainerDraggerVisibility //gd:SplitContainer.DraggerVisibility

const (
	/*The split dragger is visible when the cursor hovers it.*/
	DraggerVisible DraggerVisibility = 0
	/*The split dragger is never visible.*/
	DraggerHidden DraggerVisibility = 1
	/*The split dragger is never visible and its space collapsed.*/
	DraggerHiddenCollapsed DraggerVisibility = 2
)
