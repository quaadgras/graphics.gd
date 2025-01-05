// Package XRControllerTracker provides methods for working with XRControllerTracker object instances.
package XRControllerTracker

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/XRPositionalTracker"
import "graphics.gd/classdb/XRTracker"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
An instance of this object represents a controller that is tracked.
As controllers are turned on and the [XRInterface] detects them, instances of this object are automatically added to this list of active tracking objects accessible through the [XRServer].
The [XRController3D] consumes objects of this type and should be used in your project.
*/
type Instance [1]gdclass.XRControllerTracker
type Any interface {
	gd.IsClass
	AsXRControllerTracker() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.XRControllerTracker

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRControllerTracker"))
	return Instance{*(*gdclass.XRControllerTracker)(unsafe.Pointer(&object))}
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(XRPositionalTracker.Advanced(self.AsXRPositionalTracker()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(XRPositionalTracker.Instance(self.AsXRPositionalTracker()), name)
	}
}
func init() {
	gdclass.Register("XRControllerTracker", func(ptr gd.Object) any {
		return [1]gdclass.XRControllerTracker{*(*gdclass.XRControllerTracker)(unsafe.Pointer(&ptr))}
	})
}
