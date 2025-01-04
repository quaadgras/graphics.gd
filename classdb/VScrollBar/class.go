package VScrollBar

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/ScrollBar"
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
A vertical scrollbar, typically used to navigate through content that extends beyond the visible height of a control. It is a [Range]-based control and goes from top (min) to bottom (max). Note that this direction is the opposite of [VSlider]'s.
*/
type Instance [1]gdclass.VScrollBar
type Any interface {
	gd.IsClass
	AsVScrollBar() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VScrollBar

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VScrollBar"))
	return Instance{*(*gdclass.VScrollBar)(unsafe.Pointer(&object))}
}

func (self class) AsVScrollBar() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVScrollBar() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsScrollBar() ScrollBar.Advanced {
	return *((*ScrollBar.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsScrollBar() ScrollBar.Instance {
	return *((*ScrollBar.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(self.AsScrollBar(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsScrollBar(), name)
	}
}
func init() {
	gdclass.Register("VScrollBar", func(ptr gd.Object) any { return [1]gdclass.VScrollBar{*(*gdclass.VScrollBar)(unsafe.Pointer(&ptr))} })
}
