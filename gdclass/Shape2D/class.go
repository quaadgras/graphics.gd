package Shape2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Abstract base class for all 2D shapes, intended for use in physics.
[b]Performance:[/b] Primitive shapes, especially [CircleShape2D], are fast to check collisions against. [ConvexPolygonShape2D] is slower, and [ConcavePolygonShape2D] is the slowest.

*/
type Go [1]classdb.Shape2D

/*
Returns [code]true[/code] if this shape is colliding with another.
This method needs the transformation matrix for this shape ([param local_xform]), the shape to check collisions with ([param with_shape]), and the transformation matrix of that shape ([param shape_xform]).
*/
func (self Go) Collide(local_xform gd.Transform2D, with_shape gdclass.Shape2D, shape_xform gd.Transform2D) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).Collide(local_xform, with_shape, shape_xform))
}

/*
Returns whether this shape would collide with another, if a given movement was applied.
This method needs the transformation matrix for this shape ([param local_xform]), the movement to test on this shape ([param local_motion]), the shape to check collisions with ([param with_shape]), the transformation matrix of that shape ([param shape_xform]), and the movement to test onto the other object ([param shape_motion]).
*/
func (self Go) CollideWithMotion(local_xform gd.Transform2D, local_motion gd.Vector2, with_shape gdclass.Shape2D, shape_xform gd.Transform2D, shape_motion gd.Vector2) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).CollideWithMotion(local_xform, local_motion, with_shape, shape_xform, shape_motion))
}

/*
Returns a list of contact point pairs where this shape touches another.
If there are no collisions, the returned list is empty. Otherwise, the returned list contains contact points arranged in pairs, with entries alternating between points on the boundary of this shape and points on the boundary of [param with_shape].
A collision pair A, B can be used to calculate the collision normal with [code](B - A).normalized()[/code], and the collision depth with [code](B - A).length()[/code]. This information is typically used to separate shapes, particularly in collision solvers.
This method needs the transformation matrix for this shape ([param local_xform]), the shape to check collisions with ([param with_shape]), and the transformation matrix of that shape ([param shape_xform]).
*/
func (self Go) CollideAndGetContacts(local_xform gd.Transform2D, with_shape gdclass.Shape2D, shape_xform gd.Transform2D) []gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return []gd.Vector2(class(self).CollideAndGetContacts(gc, local_xform, with_shape, shape_xform).AsSlice())
}

/*
Returns a list of contact point pairs where this shape would touch another, if a given movement was applied.
If there would be no collisions, the returned list is empty. Otherwise, the returned list contains contact points arranged in pairs, with entries alternating between points on the boundary of this shape and points on the boundary of [param with_shape].
A collision pair A, B can be used to calculate the collision normal with [code](B - A).normalized()[/code], and the collision depth with [code](B - A).length()[/code]. This information is typically used to separate shapes, particularly in collision solvers.
This method needs the transformation matrix for this shape ([param local_xform]), the movement to test on this shape ([param local_motion]), the shape to check collisions with ([param with_shape]), the transformation matrix of that shape ([param shape_xform]), and the movement to test onto the other object ([param shape_motion]).
*/
func (self Go) CollideWithMotionAndGetContacts(local_xform gd.Transform2D, local_motion gd.Vector2, with_shape gdclass.Shape2D, shape_xform gd.Transform2D, shape_motion gd.Vector2) []gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return []gd.Vector2(class(self).CollideWithMotionAndGetContacts(gc, local_xform, local_motion, with_shape, shape_xform, shape_motion).AsSlice())
}

/*
Draws a solid shape onto a [CanvasItem] with the [RenderingServer] API filled with the specified [param color]. The exact drawing method is specific for each shape and cannot be configured.
*/
func (self Go) Draw(canvas_item gd.RID, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Draw(canvas_item, color)
}

/*
Returns a [Rect2] representing the shapes boundary.
*/
func (self Go) GetRect() gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(class(self).GetRect())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Shape2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Shape2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) CustomSolverBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetCustomSolverBias()))
}

func (self Go) SetCustomSolverBias(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCustomSolverBias(gd.Float(value))
}

//go:nosplit
func (self class) SetCustomSolverBias(bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shape2D.Bind_set_custom_solver_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCustomSolverBias() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shape2D.Bind_get_custom_solver_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this shape is colliding with another.
This method needs the transformation matrix for this shape ([param local_xform]), the shape to check collisions with ([param with_shape]), and the transformation matrix of that shape ([param shape_xform]).
*/
//go:nosplit
func (self class) Collide(local_xform gd.Transform2D, with_shape gdclass.Shape2D, shape_xform gd.Transform2D) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, local_xform)
	callframe.Arg(frame, mmm.Get(with_shape[0].AsPointer())[0])
	callframe.Arg(frame, shape_xform)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shape2D.Bind_collide, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether this shape would collide with another, if a given movement was applied.
This method needs the transformation matrix for this shape ([param local_xform]), the movement to test on this shape ([param local_motion]), the shape to check collisions with ([param with_shape]), the transformation matrix of that shape ([param shape_xform]), and the movement to test onto the other object ([param shape_motion]).
*/
//go:nosplit
func (self class) CollideWithMotion(local_xform gd.Transform2D, local_motion gd.Vector2, with_shape gdclass.Shape2D, shape_xform gd.Transform2D, shape_motion gd.Vector2) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, local_xform)
	callframe.Arg(frame, local_motion)
	callframe.Arg(frame, mmm.Get(with_shape[0].AsPointer())[0])
	callframe.Arg(frame, shape_xform)
	callframe.Arg(frame, shape_motion)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shape2D.Bind_collide_with_motion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a list of contact point pairs where this shape touches another.
If there are no collisions, the returned list is empty. Otherwise, the returned list contains contact points arranged in pairs, with entries alternating between points on the boundary of this shape and points on the boundary of [param with_shape].
A collision pair A, B can be used to calculate the collision normal with [code](B - A).normalized()[/code], and the collision depth with [code](B - A).length()[/code]. This information is typically used to separate shapes, particularly in collision solvers.
This method needs the transformation matrix for this shape ([param local_xform]), the shape to check collisions with ([param with_shape]), and the transformation matrix of that shape ([param shape_xform]).
*/
//go:nosplit
func (self class) CollideAndGetContacts(ctx gd.Lifetime, local_xform gd.Transform2D, with_shape gdclass.Shape2D, shape_xform gd.Transform2D) gd.PackedVector2Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, local_xform)
	callframe.Arg(frame, mmm.Get(with_shape[0].AsPointer())[0])
	callframe.Arg(frame, shape_xform)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shape2D.Bind_collide_and_get_contacts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector2Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a list of contact point pairs where this shape would touch another, if a given movement was applied.
If there would be no collisions, the returned list is empty. Otherwise, the returned list contains contact points arranged in pairs, with entries alternating between points on the boundary of this shape and points on the boundary of [param with_shape].
A collision pair A, B can be used to calculate the collision normal with [code](B - A).normalized()[/code], and the collision depth with [code](B - A).length()[/code]. This information is typically used to separate shapes, particularly in collision solvers.
This method needs the transformation matrix for this shape ([param local_xform]), the movement to test on this shape ([param local_motion]), the shape to check collisions with ([param with_shape]), the transformation matrix of that shape ([param shape_xform]), and the movement to test onto the other object ([param shape_motion]).
*/
//go:nosplit
func (self class) CollideWithMotionAndGetContacts(ctx gd.Lifetime, local_xform gd.Transform2D, local_motion gd.Vector2, with_shape gdclass.Shape2D, shape_xform gd.Transform2D, shape_motion gd.Vector2) gd.PackedVector2Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, local_xform)
	callframe.Arg(frame, local_motion)
	callframe.Arg(frame, mmm.Get(with_shape[0].AsPointer())[0])
	callframe.Arg(frame, shape_xform)
	callframe.Arg(frame, shape_motion)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shape2D.Bind_collide_with_motion_and_get_contacts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector2Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Draws a solid shape onto a [CanvasItem] with the [RenderingServer] API filled with the specified [param color]. The exact drawing method is specific for each shape and cannot be configured.
*/
//go:nosplit
func (self class) Draw(canvas_item gd.RID, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shape2D.Bind_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [Rect2] representing the shapes boundary.
*/
//go:nosplit
func (self class) GetRect() gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shape2D.Bind_get_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsShape2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsShape2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("Shape2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
