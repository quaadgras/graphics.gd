package HFlowContainer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/FlowContainer"
import "graphics.gd/objects/Container"
import "graphics.gd/objects/Control"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A variant of [FlowContainer] that can only arrange its child controls horizontally, wrapping them around at the borders. This is similar to how text in a book wraps around when no more words can fit on a line.
*/
type Instance [1]classdb.HFlowContainer
type Any interface {
	gd.IsClass
	AsHFlowContainer() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.HFlowContainer

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("HFlowContainer"))
	return Instance{*(*classdb.HFlowContainer)(unsafe.Pointer(&object))}
}

func (self class) AsHFlowContainer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsHFlowContainer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsFlowContainer() FlowContainer.Advanced {
	return *((*FlowContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsFlowContainer() FlowContainer.Instance {
	return *((*FlowContainer.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsFlowContainer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsFlowContainer(), name)
	}
}
func init() {
	classdb.Register("HFlowContainer", func(ptr gd.Object) any {
		return [1]classdb.HFlowContainer{*(*classdb.HFlowContainer)(unsafe.Pointer(&ptr))}
	})
}
