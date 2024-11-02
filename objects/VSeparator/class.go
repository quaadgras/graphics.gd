package VSeparator

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Separator"
import "grow.graphics/gd/objects/Control"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A vertical separator used for separating other controls that are arranged [b]horizontally[/b]. [VSeparator] is purely visual and normally drawn as a [StyleBoxLine].
*/
type Instance [1]classdb.VSeparator

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VSeparator

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VSeparator"))
	return Instance{classdb.VSeparator(object)}
}

func (self class) AsVSeparator() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVSeparator() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("VSeparator", func(ptr gd.Object) any { return classdb.VSeparator(ptr) })
}
