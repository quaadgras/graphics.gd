package ConcavePolygonShape2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Shape2D"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A 2D polyline shape, intended for use in physics. Used internally in [CollisionPolygon2D] when it's in [constant CollisionPolygon2D.BUILD_SEGMENTS] mode.
Being just a collection of interconnected line segments, [ConcavePolygonShape2D] is the most freely configurable single 2D shape. It can be used to form polygons of any nature, or even shapes that don't enclose an area. However, [ConcavePolygonShape2D] is [i]hollow[/i] even if the interconnected line segments do enclose an area, which often makes it unsuitable for physics or detection.
[b]Note:[/b] When used for collision, [ConcavePolygonShape2D] is intended to work with static [CollisionShape2D] nodes like [StaticBody2D] and will likely not behave well for [CharacterBody2D]s or [RigidBody2D]s in a mode other than Static.
[b]Warning:[/b] Physics bodies that are small have a chance to clip through this shape when moving fast. This happens because on one frame, the physics body may be on the "outside" of the shape, and on the next frame it may be "inside" it. [ConcavePolygonShape2D] is hollow, so it won't detect a collision.
[b]Performance:[/b] Due to its complexity, [ConcavePolygonShape2D] is the slowest 2D collision shape to check collisions against. Its use should generally be limited to level geometry. If the polyline is closed, [CollisionPolygon2D]'s [constant CollisionPolygon2D.BUILD_SOLIDS] mode can be used, which decomposes the polygon into convex ones; see [ConvexPolygonShape2D]'s documentation for instructions.

*/
type Go [1]classdb.ConcavePolygonShape2D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ConcavePolygonShape2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ConcavePolygonShape2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Segments() []gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return []gd.Vector2(class(self).GetSegments(gc).AsSlice())
}

func (self Go) SetSegments(value []gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSegments(gc.PackedVector2Slice(value))
}

//go:nosplit
func (self class) SetSegments(segments gd.PackedVector2Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(segments))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConcavePolygonShape2D.Bind_set_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSegments(ctx gd.Lifetime) gd.PackedVector2Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConcavePolygonShape2D.Bind_get_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector2Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsConcavePolygonShape2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsConcavePolygonShape2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsShape2D() Shape2D.GD { return *((*Shape2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsShape2D() Shape2D.Go { return *((*Shape2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsShape2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsShape2D(), name)
	}
}
func init() {classdb.Register("ConcavePolygonShape2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
