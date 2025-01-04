// Package HSeparator provides methods for working with HSeparator object instances.
package HSeparator

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Separator"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A horizontal separator used for separating other controls that are arranged [b]vertically[/b]. [HSeparator] is purely visual and normally drawn as a [StyleBoxLine].
*/
type Instance [1]gdclass.HSeparator
type Any interface {
	gd.IsClass
	AsHSeparator() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.HSeparator

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("HSeparator"))
	return Instance{*(*gdclass.HSeparator)(unsafe.Pointer(&object))}
}

func (self class) AsHSeparator() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsHSeparator() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsSeparator() Separator.Advanced {
	return *((*Separator.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSeparator() Separator.Instance {
	return *((*Separator.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsSeparator(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsSeparator(), name)
	}
}
func init() {
	gdclass.Register("HSeparator", func(ptr gd.Object) any { return [1]gdclass.HSeparator{*(*gdclass.HSeparator)(unsafe.Pointer(&ptr))} })
}
