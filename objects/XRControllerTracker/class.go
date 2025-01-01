package XRControllerTracker

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/XRPositionalTracker"
import "graphics.gd/objects/XRTracker"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
An instance of this object represents a controller that is tracked.
As controllers are turned on and the [XRInterface] detects them, instances of this object are automatically added to this list of active tracking objects accessible through the [XRServer].
The [XRController3D] consumes objects of this type and should be used in your project.
*/
type Instance [1]classdb.XRControllerTracker
type Any interface {
	gd.IsClass
	AsXRControllerTracker() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.XRControllerTracker

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRControllerTracker"))
	return Instance{classdb.XRControllerTracker(object)}
}

func (self class) AsXRControllerTracker() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsXRControllerTracker() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsXRPositionalTracker() XRPositionalTracker.Advanced {
	return *((*XRPositionalTracker.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsXRPositionalTracker() XRPositionalTracker.Instance {
	return *((*XRPositionalTracker.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsXRTracker() XRTracker.Advanced {
	return *((*XRTracker.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsXRTracker() XRTracker.Instance {
	return *((*XRTracker.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsXRPositionalTracker(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsXRPositionalTracker(), name)
	}
}
func init() {
	classdb.Register("XRControllerTracker", func(ptr gd.Object) any { return [1]classdb.XRControllerTracker{classdb.XRControllerTracker(ptr)} })
}
