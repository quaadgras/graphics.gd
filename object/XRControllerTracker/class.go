package XRControllerTracker

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/XRPositionalTracker"
import "grow.graphics/gd/object/XRTracker"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
An instance of this object represents a controller that is tracked.
As controllers are turned on and the [XRInterface] detects them, instances of this object are automatically added to this list of active tracking objects accessible through the [XRServer].
The [XRController3D] consumes objects of this type and should be used in your project.

*/
type Simple [1]classdb.XRControllerTracker
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.XRControllerTracker
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsXRControllerTracker() Expert { return self[0].AsXRControllerTracker() }


//go:nosplit
func (self Simple) AsXRControllerTracker() Simple { return self[0].AsXRControllerTracker() }


//go:nosplit
func (self class) AsXRPositionalTracker() XRPositionalTracker.Expert { return self[0].AsXRPositionalTracker() }


//go:nosplit
func (self Simple) AsXRPositionalTracker() XRPositionalTracker.Simple { return self[0].AsXRPositionalTracker() }


//go:nosplit
func (self class) AsXRTracker() XRTracker.Expert { return self[0].AsXRTracker() }


//go:nosplit
func (self Simple) AsXRTracker() XRTracker.Simple { return self[0].AsXRTracker() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("XRControllerTracker", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
