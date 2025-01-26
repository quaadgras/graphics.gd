// Package ConcavePolygonShape3D provides methods for working with ConcavePolygonShape3D object instances.
package ConcavePolygonShape3D

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
A 3D trimesh shape, intended for use in physics. Usually used to provide a shape for a [CollisionShape3D].
Being just a collection of interconnected triangles, [ConcavePolygonShape3D] is the most freely configurable single 3D shape. It can be used to form polyhedra of any nature, or even shapes that don't enclose a volume. However, [ConcavePolygonShape3D] is [i]hollow[/i] even if the interconnected triangles do enclose a volume, which often makes it unsuitable for physics or detection.
[b]Note:[/b] When used for collision, [ConcavePolygonShape3D] is intended to work with static [CollisionShape3D] nodes like [StaticBody3D] and will likely not behave well for [CharacterBody3D]s or [RigidBody3D]s in a mode other than Static.
[b]Warning:[/b] Physics bodies that are small have a chance to clip through this shape when moving fast. This happens because on one frame, the physics body may be on the "outside" of the shape, and on the next frame it may be "inside" it. [ConcavePolygonShape3D] is hollow, so it won't detect a collision.
[b]Performance:[/b] Due to its complexity, [ConcavePolygonShape3D] is the slowest 3D collision shape to check collisions against. Its use should generally be limited to level geometry. For convex geometry, [ConvexPolygonShape3D] should be used. For dynamic physics bodies that need concave collision, several [ConvexPolygonShape3D]s can be used to represent its collision by using convex decomposition; see [ConvexPolygonShape3D]'s documentation for instructions.
*/
type Instance [1]gdclass.ConcavePolygonShape3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsConcavePolygonShape3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ConcavePolygonShape3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ConcavePolygonShape3D"))
	casted := Instance{*(*gdclass.ConcavePolygonShape3D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Data() []Vector3.XYZ {
	return []Vector3.XYZ(class(self).GetFaces().AsSlice())
}

func (self Instance) SetData(value []Vector3.XYZ) {
	class(self).SetFaces(gd.NewPackedVector3Slice(*(*[]gd.Vector3)(unsafe.Pointer(&value))))
}

func (self Instance) BackfaceCollision() bool {
	return bool(class(self).IsBackfaceCollisionEnabled())
}

func (self Instance) SetBackfaceCollision(value bool) {
	class(self).SetBackfaceCollisionEnabled(value)
}

/*
Sets the faces of the trimesh shape from an array of vertices. The [param faces] array should be composed of triples such that each triple of vertices defines a triangle.
*/
//go:nosplit
func (self class) SetFaces(faces gd.PackedVector3Array) { //gd:ConcavePolygonShape3D.set_faces
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(faces))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConcavePolygonShape3D.Bind_set_faces, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the faces of the trimesh shape as an array of vertices. The array (of length divisible by three) is naturally divided into triples; each triple of vertices defines a triangle.
*/
//go:nosplit
func (self class) GetFaces() gd.PackedVector3Array { //gd:ConcavePolygonShape3D.get_faces
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConcavePolygonShape3D.Bind_get_faces, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedVector3Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBackfaceCollisionEnabled(enabled bool) { //gd:ConcavePolygonShape3D.set_backface_collision_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConcavePolygonShape3D.Bind_set_backface_collision_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsBackfaceCollisionEnabled() bool { //gd:ConcavePolygonShape3D.is_backface_collision_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConcavePolygonShape3D.Bind_is_backface_collision_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsConcavePolygonShape3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsConcavePolygonShape3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsShape3D() Shape3D.Advanced          { return *((*Shape3D.Advanced)(unsafe.Pointer(&self))) }
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
	gdclass.Register("ConcavePolygonShape3D", func(ptr gd.Object) any {
		return [1]gdclass.ConcavePolygonShape3D{*(*gdclass.ConcavePolygonShape3D)(unsafe.Pointer(&ptr))}
	})
}
