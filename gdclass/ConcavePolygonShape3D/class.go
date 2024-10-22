package ConcavePolygonShape3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Shape3D"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A 3D trimesh shape, intended for use in physics. Usually used to provide a shape for a [CollisionShape3D].
Being just a collection of interconnected triangles, [ConcavePolygonShape3D] is the most freely configurable single 3D shape. It can be used to form polyhedra of any nature, or even shapes that don't enclose a volume. However, [ConcavePolygonShape3D] is [i]hollow[/i] even if the interconnected triangles do enclose a volume, which often makes it unsuitable for physics or detection.
[b]Note:[/b] When used for collision, [ConcavePolygonShape3D] is intended to work with static [CollisionShape3D] nodes like [StaticBody3D] and will likely not behave well for [CharacterBody3D]s or [RigidBody3D]s in a mode other than Static.
[b]Warning:[/b] Physics bodies that are small have a chance to clip through this shape when moving fast. This happens because on one frame, the physics body may be on the "outside" of the shape, and on the next frame it may be "inside" it. [ConcavePolygonShape3D] is hollow, so it won't detect a collision.
[b]Performance:[/b] Due to its complexity, [ConcavePolygonShape3D] is the slowest 3D collision shape to check collisions against. Its use should generally be limited to level geometry. For convex geometry, [ConvexPolygonShape3D] should be used. For dynamic physics bodies that need concave collision, several [ConvexPolygonShape3D]s can be used to represent its collision by using convex decomposition; see [ConvexPolygonShape3D]'s documentation for instructions.

*/
type Go [1]classdb.ConcavePolygonShape3D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ConcavePolygonShape3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ConcavePolygonShape3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Data() []gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return []gd.Vector3(class(self).GetFaces(gc).AsSlice())
}

func (self Go) SetData(value []gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFaces(gc.PackedVector3Slice(value))
}

func (self Go) BackfaceCollision() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsBackfaceCollisionEnabled())
}

func (self Go) SetBackfaceCollision(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBackfaceCollisionEnabled(value)
}

/*
Sets the faces of the trimesh shape from an array of vertices. The [param faces] array should be composed of triples such that each triple of vertices defines a triangle.
*/
//go:nosplit
func (self class) SetFaces(faces gd.PackedVector3Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(faces))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConcavePolygonShape3D.Bind_set_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the faces of the trimesh shape as an array of vertices. The array (of length divisible by three) is naturally divided into triples; each triple of vertices defines a triangle.
*/
//go:nosplit
func (self class) GetFaces(ctx gd.Lifetime) gd.PackedVector3Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConcavePolygonShape3D.Bind_get_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector3Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBackfaceCollisionEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConcavePolygonShape3D.Bind_set_backface_collision_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsBackfaceCollisionEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConcavePolygonShape3D.Bind_is_backface_collision_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsConcavePolygonShape3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsConcavePolygonShape3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsShape3D() Shape3D.GD { return *((*Shape3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsShape3D() Shape3D.Go { return *((*Shape3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsShape3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsShape3D(), name)
	}
}
func init() {classdb.Register("ConcavePolygonShape3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
