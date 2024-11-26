package WorldBoundaryShape3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Shape3D"
import "grow.graphics/gd/objects/Resource"
import "grow.graphics/gd/variant/Plane"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A 3D world boundary shape, intended for use in physics. [WorldBoundaryShape3D] works like an infinite plane that forces all physics bodies to stay above it. The [member plane]'s normal determines which direction is considered as "above" and in the editor, the line over the plane represents this direction. It can for example be used for endless flat floors.
*/
type Instance [1]classdb.WorldBoundaryShape3D

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.WorldBoundaryShape3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WorldBoundaryShape3D"))
	return Instance{classdb.WorldBoundaryShape3D(object)}
}

func (self Instance) Plane() Plane.NormalD {
	return Plane.NormalD(class(self).GetPlane())
}

func (self Instance) SetPlane(value Plane.NormalD) {
	class(self).SetPlane(gd.Plane(value))
}

//go:nosplit
func (self class) SetPlane(plane gd.Plane) {
	var frame = callframe.New()
	callframe.Arg(frame, plane)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldBoundaryShape3D.Bind_set_plane, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPlane() gd.Plane {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Plane](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldBoundaryShape3D.Bind_get_plane, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsWorldBoundaryShape3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWorldBoundaryShape3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsShape3D() Shape3D.Advanced         { return *((*Shape3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsShape3D() Shape3D.Instance {
	return *((*Shape3D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsShape3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsShape3D(), name)
	}
}
func init() {
	classdb.Register("WorldBoundaryShape3D", func(ptr gd.Object) any { return classdb.WorldBoundaryShape3D(ptr) })
}
