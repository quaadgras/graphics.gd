package Slider

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Range"
import "grow.graphics/gd/objects/Control"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Abstract base class for sliders, used to adjust a value by moving a grabber along a horizontal or vertical axis. Sliders are [Range]-based controls.
*/
type Instance [1]classdb.Slider

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Slider

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Slider"))
	return Instance{classdb.Slider(object)}
}

func (self Instance) Editable() bool {
	return bool(class(self).IsEditable())
}

func (self Instance) SetEditable(value bool) {
	class(self).SetEditable(value)
}

func (self Instance) Scrollable() bool {
	return bool(class(self).IsScrollable())
}

func (self Instance) SetScrollable(value bool) {
	class(self).SetScrollable(value)
}

func (self Instance) TickCount() int {
	return int(int(class(self).GetTicks()))
}

func (self Instance) SetTickCount(value int) {
	class(self).SetTicks(gd.Int(value))
}

func (self Instance) TicksOnBorders() bool {
	return bool(class(self).GetTicksOnBorders())
}

func (self Instance) SetTicksOnBorders(value bool) {
	class(self).SetTicksOnBorders(value)
}

//go:nosplit
func (self class) SetTicks(count gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Slider.Bind_set_ticks, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTicks() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Slider.Bind_get_ticks, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTicksOnBorders() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Slider.Bind_get_ticks_on_borders, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTicksOnBorders(ticks_on_border bool) {
	var frame = callframe.New()
	callframe.Arg(frame, ticks_on_border)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Slider.Bind_set_ticks_on_borders, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetEditable(editable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, editable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Slider.Bind_set_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsEditable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Slider.Bind_is_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScrollable(scrollable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, scrollable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Slider.Bind_set_scrollable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsScrollable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Slider.Bind_is_scrollable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnDragStarted(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("drag_started"), gd.NewCallable(cb), 0)
}

func (self Instance) OnDragEnded(cb func(value_changed bool)) {
	self[0].AsObject().Connect(gd.NewStringName("drag_ended"), gd.NewCallable(cb), 0)
}

func (self class) AsSlider() Advanced          { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSlider() Instance       { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRange() Range.Advanced     { return *((*Range.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRange() Range.Instance  { return *((*Range.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsRange(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRange(), name)
	}
}
func init() { classdb.Register("Slider", func(ptr gd.Object) any { return classdb.Slider(ptr) }) }
