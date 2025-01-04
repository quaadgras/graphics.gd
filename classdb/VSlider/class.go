// Package VSlider provides methods for working with VSlider object instances.
package VSlider

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Slider"
import "graphics.gd/classdb/Range"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A vertical slider, used to adjust a value by moving a grabber along a vertical axis. It is a [Range]-based control and goes from bottom (min) to top (max). Note that this direction is the opposite of [VScrollBar]'s.
*/
type Instance [1]gdclass.VSlider
type Any interface {
	gd.IsClass
	AsVSlider() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VSlider

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VSlider"))
	return Instance{*(*gdclass.VSlider)(unsafe.Pointer(&object))}
}

func (self class) AsVSlider() Advanced          { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVSlider() Instance       { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsSlider() Slider.Advanced    { return *((*Slider.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSlider() Slider.Instance { return *((*Slider.Instance)(unsafe.Pointer(&self))) }
func (self class) AsRange() Range.Advanced      { return *((*Range.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRange() Range.Instance   { return *((*Range.Instance)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.Advanced  { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsSlider(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsSlider(), name)
	}
}
func init() {
	gdclass.Register("VSlider", func(ptr gd.Object) any { return [1]gdclass.VSlider{*(*gdclass.VSlider)(unsafe.Pointer(&ptr))} })
}
