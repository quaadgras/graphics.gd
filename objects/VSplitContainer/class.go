package VSplitContainer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/SplitContainer"
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
A container that accepts only two child controls, then arranges them vertically and creates a divisor between them. The divisor can be dragged around to change the size relation between the child controls.
*/
type Instance [1]classdb.VSplitContainer
type Any interface {
	gd.IsClass
	AsVSplitContainer() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VSplitContainer

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VSplitContainer"))
	return Instance{*(*classdb.VSplitContainer)(unsafe.Pointer(&object))}
}

func (self class) AsVSplitContainer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVSplitContainer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsSplitContainer() SplitContainer.Advanced {
	return *((*SplitContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSplitContainer() SplitContainer.Instance {
	return *((*SplitContainer.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsSplitContainer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsSplitContainer(), name)
	}
}
func init() {
	classdb.Register("VSplitContainer", func(ptr gd.Object) any {
		return [1]classdb.VSplitContainer{*(*classdb.VSplitContainer)(unsafe.Pointer(&ptr))}
	})
}
