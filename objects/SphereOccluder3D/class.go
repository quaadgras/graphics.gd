package SphereOccluder3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Occluder3D"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[SphereOccluder3D] stores a sphere shape that can be used by the engine's occlusion culling system.
See [OccluderInstance3D]'s documentation for instructions on setting up occlusion culling.
*/
type Instance [1]classdb.SphereOccluder3D
type Any interface {
	gd.IsClass
	AsSphereOccluder3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.SphereOccluder3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SphereOccluder3D"))
	return Instance{*(*classdb.SphereOccluder3D)(unsafe.Pointer(&object))}
}

func (self Instance) Radius() Float.X {
	return Float.X(Float.X(class(self).GetRadius()))
}

func (self Instance) SetRadius(value Float.X) {
	class(self).SetRadius(gd.Float(value))
}

//go:nosplit
func (self class) SetRadius(radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SphereOccluder3D.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SphereOccluder3D.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSphereOccluder3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSphereOccluder3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsOccluder3D() Occluder3D.Advanced {
	return *((*Occluder3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsOccluder3D() Occluder3D.Instance {
	return *((*Occluder3D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsOccluder3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsOccluder3D(), name)
	}
}
func init() {
	classdb.Register("SphereOccluder3D", func(ptr gd.Object) any {
		return [1]classdb.SphereOccluder3D{*(*classdb.SphereOccluder3D)(unsafe.Pointer(&ptr))}
	})
}
