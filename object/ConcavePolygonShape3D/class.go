package ConcavePolygonShape3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Shape3D"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
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
type Simple [1]classdb.ConcavePolygonShape3D
func (self Simple) SetFaces(faces gd.PackedVector3Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFaces(faces)
}
func (self Simple) GetFaces() gd.PackedVector3Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector3Array(Expert(self).GetFaces(gc))
}
func (self Simple) SetBackfaceCollisionEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBackfaceCollisionEnabled(enabled)
}
func (self Simple) IsBackfaceCollisionEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsBackfaceCollisionEnabled())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ConcavePolygonShape3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsConcavePolygonShape3D() Expert { return self[0].AsConcavePolygonShape3D() }


//go:nosplit
func (self Simple) AsConcavePolygonShape3D() Simple { return self[0].AsConcavePolygonShape3D() }


//go:nosplit
func (self class) AsShape3D() Shape3D.Expert { return self[0].AsShape3D() }


//go:nosplit
func (self Simple) AsShape3D() Shape3D.Simple { return self[0].AsShape3D() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


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
func init() {classdb.Register("ConcavePolygonShape3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
