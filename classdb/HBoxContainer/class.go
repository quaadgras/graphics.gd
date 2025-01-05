// Package HBoxContainer provides methods for working with HBoxContainer object instances.
package HBoxContainer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/BoxContainer"
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

/*
A variant of [BoxContainer] that can only arrange its child controls horizontally. Child controls are rearranged automatically when their minimum size changes.
*/
type Instance [1]gdclass.HBoxContainer
type Any interface {
	gd.IsClass
	AsHBoxContainer() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.HBoxContainer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("HBoxContainer"))
	return Instance{*(*gdclass.HBoxContainer)(unsafe.Pointer(&object))}
}

func (self class) AsHBoxContainer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsHBoxContainer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsBoxContainer() BoxContainer.Advanced {
	return *((*BoxContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsBoxContainer() BoxContainer.Instance {
	return *((*BoxContainer.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(BoxContainer.Advanced(self.AsBoxContainer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(BoxContainer.Instance(self.AsBoxContainer()), name)
	}
}
func init() {
	gdclass.Register("HBoxContainer", func(ptr gd.Object) any {
		return [1]gdclass.HBoxContainer{*(*gdclass.HBoxContainer)(unsafe.Pointer(&ptr))}
	})
}
