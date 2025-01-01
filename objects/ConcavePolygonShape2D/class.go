package ConcavePolygonShape2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Shape2D"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Vector2"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A 2D polyline shape, intended for use in physics. Used internally in [CollisionPolygon2D] when it's in [constant CollisionPolygon2D.BUILD_SEGMENTS] mode.
Being just a collection of interconnected line segments, [ConcavePolygonShape2D] is the most freely configurable single 2D shape. It can be used to form polygons of any nature, or even shapes that don't enclose an area. However, [ConcavePolygonShape2D] is [i]hollow[/i] even if the interconnected line segments do enclose an area, which often makes it unsuitable for physics or detection.
[b]Note:[/b] When used for collision, [ConcavePolygonShape2D] is intended to work with static [CollisionShape2D] nodes like [StaticBody2D] and will likely not behave well for [CharacterBody2D]s or [RigidBody2D]s in a mode other than Static.
[b]Warning:[/b] Physics bodies that are small have a chance to clip through this shape when moving fast. This happens because on one frame, the physics body may be on the "outside" of the shape, and on the next frame it may be "inside" it. [ConcavePolygonShape2D] is hollow, so it won't detect a collision.
[b]Performance:[/b] Due to its complexity, [ConcavePolygonShape2D] is the slowest 2D collision shape to check collisions against. Its use should generally be limited to level geometry. If the polyline is closed, [CollisionPolygon2D]'s [constant CollisionPolygon2D.BUILD_SOLIDS] mode can be used, which decomposes the polygon into convex ones; see [ConvexPolygonShape2D]'s documentation for instructions.
*/
type Instance [1]classdb.ConcavePolygonShape2D
type Any interface {
	gd.IsClass
	AsConcavePolygonShape2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ConcavePolygonShape2D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ConcavePolygonShape2D"))
	return Instance{classdb.ConcavePolygonShape2D(object)}
}

func (self Instance) Segments() []Vector2.XY {
	return []Vector2.XY(class(self).GetSegments().AsSlice())
}

func (self Instance) SetSegments(value []Vector2.XY) {
	class(self).SetSegments(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&value))))
}

//go:nosplit
func (self class) SetSegments(segments gd.PackedVector2Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(segments))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConcavePolygonShape2D.Bind_set_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSegments() gd.PackedVector2Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConcavePolygonShape2D.Bind_get_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsConcavePolygonShape2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsConcavePolygonShape2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsShape2D() Shape2D.Advanced          { return *((*Shape2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsShape2D() Shape2D.Instance {
	return *((*Shape2D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsShape2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsShape2D(), name)
	}
}
func init() {
	classdb.Register("ConcavePolygonShape2D", func(ptr gd.Object) any { return [1]classdb.ConcavePolygonShape2D{classdb.ConcavePolygonShape2D(ptr)} })
}
