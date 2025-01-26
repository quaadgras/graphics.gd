// Package ConvexPolygonShape3D provides methods for working with ConvexPolygonShape3D object instances.
package ConvexPolygonShape3D

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
import "graphics.gd/variant/Dictionary"
import "graphics.gd/classdb/Shape3D"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector3"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any

/*
A 3D convex polyhedron shape, intended for use in physics. Usually used to provide a shape for a [CollisionShape3D].
[ConvexPolygonShape3D] is [i]solid[/i], which means it detects collisions from objects that are fully inside it, unlike [ConcavePolygonShape3D] which is hollow. This makes it more suitable for both detection and physics.
[b]Convex decomposition:[/b] A concave polyhedron can be split up into several convex polyhedra. This allows dynamic physics bodies to have complex concave collisions (at a performance cost) and can be achieved by using several [ConvexPolygonShape3D] nodes. To generate a convex decomposition from a mesh, select the [MeshInstance3D] node, go to the [b]Mesh[/b] menu that appears above the viewport, and choose [b]Create Multiple Convex Collision Siblings[/b]. Alternatively, [method MeshInstance3D.create_multiple_convex_collisions] can be called in a script to perform this decomposition at run-time.
[b]Performance:[/b] [ConvexPolygonShape3D] is faster to check collisions against compared to [ConcavePolygonShape3D], but it is slower than primitive collision shapes such as [SphereShape3D] and [BoxShape3D]. Its use should generally be limited to medium-sized objects that cannot have their collision accurately represented by primitive shapes.
*/
type Instance [1]gdclass.ConvexPolygonShape3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsConvexPolygonShape3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ConvexPolygonShape3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ConvexPolygonShape3D"))
	casted := Instance{*(*gdclass.ConvexPolygonShape3D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Points() []Vector3.XYZ {
	return []Vector3.XYZ(class(self).GetPoints().AsSlice())
}

func (self Instance) SetPoints(value []Vector3.XYZ) {
	class(self).SetPoints(gd.NewPackedVector3Slice(*(*[]gd.Vector3)(unsafe.Pointer(&value))))
}

//go:nosplit
func (self class) SetPoints(points gd.PackedVector3Array) { //gd:ConvexPolygonShape3D.set_points
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(points))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConvexPolygonShape3D.Bind_set_points, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPoints() gd.PackedVector3Array { //gd:ConvexPolygonShape3D.get_points
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConvexPolygonShape3D.Bind_get_points, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedVector3Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsConvexPolygonShape3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsConvexPolygonShape3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("ConvexPolygonShape3D", func(ptr gd.Object) any {
		return [1]gdclass.ConvexPolygonShape3D{*(*gdclass.ConvexPolygonShape3D)(unsafe.Pointer(&ptr))}
	})
}
