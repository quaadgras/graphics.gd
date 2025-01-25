// Package WorldBoundaryShape3D provides methods for working with WorldBoundaryShape3D object instances.
package WorldBoundaryShape3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/classdb/Shape3D"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Plane"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
A 3D world boundary shape, intended for use in physics. [WorldBoundaryShape3D] works like an infinite plane that forces all physics bodies to stay above it. The [member plane]'s normal determines which direction is considered as "above" and in the editor, the line over the plane represents this direction. It can for example be used for endless flat floors.
*/
type Instance [1]gdclass.WorldBoundaryShape3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsWorldBoundaryShape3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.WorldBoundaryShape3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WorldBoundaryShape3D"))
	casted := Instance{*(*gdclass.WorldBoundaryShape3D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
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
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldBoundaryShape3D.Bind_set_plane, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPlane() gd.Plane {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Plane](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldBoundaryShape3D.Bind_get_plane, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Shape3D.Advanced(self.AsShape3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Shape3D.Instance(self.AsShape3D()), name)
	}
}
func init() {
	gdclass.Register("WorldBoundaryShape3D", func(ptr gd.Object) any {
		return [1]gdclass.WorldBoundaryShape3D{*(*gdclass.WorldBoundaryShape3D)(unsafe.Pointer(&ptr))}
	})
}
