package StyleBoxEmpty

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/StyleBox"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
An empty [StyleBox] that can be used to display nothing instead of the default style (e.g. it can "disable" [code]focus[/code] styles).
*/
type Instance [1]classdb.StyleBoxEmpty
type Any interface {
	gd.IsClass
	AsStyleBoxEmpty() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.StyleBoxEmpty

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StyleBoxEmpty"))
	return Instance{classdb.StyleBoxEmpty(object)}
}

func (self class) AsStyleBoxEmpty() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStyleBoxEmpty() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsStyleBox() StyleBox.Advanced {
	return *((*StyleBox.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsStyleBox() StyleBox.Instance {
	return *((*StyleBox.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsStyleBox(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsStyleBox(), name)
	}
}
func init() {
	classdb.Register("StyleBoxEmpty", func(ptr gd.Object) any { return [1]classdb.StyleBoxEmpty{classdb.StyleBoxEmpty(ptr)} })
}
