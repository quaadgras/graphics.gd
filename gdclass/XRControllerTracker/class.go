package XRControllerTracker

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/XRPositionalTracker"
import "grow.graphics/gd/gdclass/XRTracker"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
An instance of this object represents a controller that is tracked.
As controllers are turned on and the [XRInterface] detects them, instances of this object are automatically added to this list of active tracking objects accessible through the [XRServer].
The [XRController3D] consumes objects of this type and should be used in your project.

*/
type Go [1]classdb.XRControllerTracker
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.XRControllerTracker
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("XRControllerTracker"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsXRControllerTracker() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRControllerTracker() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsXRPositionalTracker() XRPositionalTracker.GD { return *((*XRPositionalTracker.GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRPositionalTracker() XRPositionalTracker.Go { return *((*XRPositionalTracker.Go)(unsafe.Pointer(&self))) }
func (self class) AsXRTracker() XRTracker.GD { return *((*XRTracker.GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRTracker() XRTracker.Go { return *((*XRTracker.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsXRPositionalTracker(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsXRPositionalTracker(), name)
	}
}
func init() {classdb.Register("XRControllerTracker", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
