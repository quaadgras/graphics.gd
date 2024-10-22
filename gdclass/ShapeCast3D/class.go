package ShapeCast3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Shape casting allows to detect collision objects by sweeping its [member shape] along the cast direction determined by [member target_position]. This is similar to [RayCast3D], but it allows for sweeping a region of space, rather than just a straight line. [ShapeCast3D] can detect multiple collision objects. It is useful for things like wide laser beams or snapping a simple shape to a floor.
Immediate collision overlaps can be done with the [member target_position] set to [code]Vector3(0, 0, 0)[/code] and by calling [method force_shapecast_update] within the same physics frame. This helps to overcome some limitations of [Area3D] when used as an instantaneous detection area, as collision information isn't immediately available to it.
[b]Note:[/b] Shape casting is more computationally expensive than ray casting.

*/
type Go [1]classdb.ShapeCast3D

/*
This method does nothing.
*/
func (self Go) ResourceChanged(resource gdclass.Resource) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ResourceChanged(resource)
}

/*
Returns whether any object is intersecting with the shape's vector (considering the vector length).
*/
func (self Go) IsColliding() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsColliding())
}

/*
The number of collisions detected at the point of impact. Use this to iterate over multiple collisions as provided by [method get_collider], [method get_collider_shape], [method get_collision_point], and [method get_collision_normal] methods.
*/
func (self Go) GetCollisionCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCollisionCount()))
}

/*
Updates the collision information for the shape immediately, without waiting for the next [code]_physics_process[/code] call. Use this method, for example, when the shape or its parent has changed state.
[b]Note:[/b] [code]enabled == true[/code] is not required for this to work.
*/
func (self Go) ForceShapecastUpdate() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ForceShapecastUpdate()
}

/*
Returns the collided [Object] of one of the multiple collisions at [param index], or [code]null[/code] if no object is intersecting the shape (i.e. [method is_colliding] returns [code]false[/code]).
*/
func (self Go) GetCollider(index int) gd.Object {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Object(class(self).GetCollider(gc, gd.Int(index)))
}

/*
Returns the [RID] of the collided object of one of the multiple collisions at [param index].
*/
func (self Go) GetColliderRid(index int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetColliderRid(gd.Int(index)))
}

/*
Returns the shape ID of the colliding shape of one of the multiple collisions at [param index], or [code]0[/code] if no object is intersecting the shape (i.e. [method is_colliding] returns [code]false[/code]).
*/
func (self Go) GetColliderShape(index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetColliderShape(gd.Int(index))))
}

/*
Returns the collision point of one of the multiple collisions at [param index] where the shape intersects the colliding object.
[b]Note:[/b] this point is in the [b]global[/b] coordinate system.
*/
func (self Go) GetCollisionPoint(index int) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetCollisionPoint(gd.Int(index)))
}

/*
Returns the normal of one of the multiple collisions at [param index] of the intersecting object.
*/
func (self Go) GetCollisionNormal(index int) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetCollisionNormal(gd.Int(index)))
}

/*
The fraction from the [ShapeCast3D]'s origin to its [member target_position] (between 0 and 1) of how far the shape can move without triggering a collision.
*/
func (self Go) GetClosestCollisionSafeFraction() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetClosestCollisionSafeFraction()))
}

/*
The fraction from the [ShapeCast3D]'s origin to its [member target_position] (between 0 and 1) of how far the shape must move to trigger a collision.
In ideal conditions this would be the same as [method get_closest_collision_safe_fraction], however shape casting is calculated in discrete steps, so the precise point of collision can occur between two calculated positions.
*/
func (self Go) GetClosestCollisionUnsafeFraction() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetClosestCollisionUnsafeFraction()))
}

/*
Adds a collision exception so the shape does not report collisions with the specified [RID].
*/
func (self Go) AddExceptionRid(rid gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddExceptionRid(rid)
}

/*
Adds a collision exception so the shape does not report collisions with the specified [CollisionObject3D] node.
*/
func (self Go) AddException(node gdclass.CollisionObject3D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddException(node)
}

/*
Removes a collision exception so the shape does report collisions with the specified [RID].
*/
func (self Go) RemoveExceptionRid(rid gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveExceptionRid(rid)
}

/*
Removes a collision exception so the shape does report collisions with the specified [CollisionObject3D] node.
*/
func (self Go) RemoveException(node gdclass.CollisionObject3D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveException(node)
}

/*
Removes all collision exceptions for this [ShapeCast3D].
*/
func (self Go) ClearExceptions() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearExceptions()
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
func (self Go) SetCollisionMaskValue(layer_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Go) GetCollisionMaskValue(layer_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetCollisionMaskValue(gd.Int(layer_number)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ShapeCast3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ShapeCast3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Enabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsEnabled())
}

func (self Go) SetEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnabled(value)
}

func (self Go) Shape() gdclass.Shape3D {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Shape3D(class(self).GetShape(gc))
}

func (self Go) SetShape(value gdclass.Shape3D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShape(value)
}

func (self Go) ExcludeParent() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetExcludeParentBody())
}

func (self Go) SetExcludeParent(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetExcludeParentBody(value)
}

func (self Go) TargetPosition() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetTargetPosition())
}

func (self Go) SetTargetPosition(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTargetPosition(value)
}

func (self Go) Margin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMargin()))
}

func (self Go) SetMargin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMargin(gd.Float(value))
}

func (self Go) MaxResults() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMaxResults()))
}

func (self Go) SetMaxResults(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxResults(gd.Int(value))
}

func (self Go) CollisionMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCollisionMask()))
}

func (self Go) SetCollisionMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionMask(gd.Int(value))
}

func (self Go) CollideWithAreas() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsCollideWithAreasEnabled())
}

func (self Go) SetCollideWithAreas(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollideWithAreas(value)
}

func (self Go) CollideWithBodies() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsCollideWithBodiesEnabled())
}

func (self Go) SetCollideWithBodies(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollideWithBodies(value)
}

func (self Go) DebugShapeCustomColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetDebugShapeCustomColor())
}

func (self Go) SetDebugShapeCustomColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDebugShapeCustomColor(value)
}

/*
This method does nothing.
*/
//go:nosplit
func (self class) ResourceChanged(resource gdclass.Resource)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(resource[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_resource_changed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_is_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShape(shape gdclass.Shape3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(shape[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_set_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShape(ctx gd.Lifetime) gdclass.Shape3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_get_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Shape3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTargetPosition(local_point gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, local_point)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_set_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTargetPosition() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_get_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMargin(margin gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_set_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMargin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_get_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxResults(max_results gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_results)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_set_max_results, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxResults() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_get_max_results, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether any object is intersecting with the shape's vector (considering the vector length).
*/
//go:nosplit
func (self class) IsColliding() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_is_colliding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
The number of collisions detected at the point of impact. Use this to iterate over multiple collisions as provided by [method get_collider], [method get_collider_shape], [method get_collision_point], and [method get_collision_normal] methods.
*/
//go:nosplit
func (self class) GetCollisionCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_get_collision_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Updates the collision information for the shape immediately, without waiting for the next [code]_physics_process[/code] call. Use this method, for example, when the shape or its parent has changed state.
[b]Note:[/b] [code]enabled == true[/code] is not required for this to work.
*/
//go:nosplit
func (self class) ForceShapecastUpdate()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_force_shapecast_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the collided [Object] of one of the multiple collisions at [param index], or [code]null[/code] if no object is intersecting the shape (i.e. [method is_colliding] returns [code]false[/code]).
*/
//go:nosplit
func (self class) GetCollider(ctx gd.Lifetime, index gd.Int) gd.Object {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_get_collider, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gd.Object
	ret.SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [RID] of the collided object of one of the multiple collisions at [param index].
*/
//go:nosplit
func (self class) GetColliderRid(index gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_get_collider_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the shape ID of the colliding shape of one of the multiple collisions at [param index], or [code]0[/code] if no object is intersecting the shape (i.e. [method is_colliding] returns [code]false[/code]).
*/
//go:nosplit
func (self class) GetColliderShape(index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_get_collider_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the collision point of one of the multiple collisions at [param index] where the shape intersects the colliding object.
[b]Note:[/b] this point is in the [b]global[/b] coordinate system.
*/
//go:nosplit
func (self class) GetCollisionPoint(index gd.Int) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_get_collision_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the normal of one of the multiple collisions at [param index] of the intersecting object.
*/
//go:nosplit
func (self class) GetCollisionNormal(index gd.Int) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_get_collision_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
The fraction from the [ShapeCast3D]'s origin to its [member target_position] (between 0 and 1) of how far the shape can move without triggering a collision.
*/
//go:nosplit
func (self class) GetClosestCollisionSafeFraction() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_get_closest_collision_safe_fraction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
The fraction from the [ShapeCast3D]'s origin to its [member target_position] (between 0 and 1) of how far the shape must move to trigger a collision.
In ideal conditions this would be the same as [method get_closest_collision_safe_fraction], however shape casting is calculated in discrete steps, so the precise point of collision can occur between two calculated positions.
*/
//go:nosplit
func (self class) GetClosestCollisionUnsafeFraction() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_get_closest_collision_unsafe_fraction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a collision exception so the shape does not report collisions with the specified [RID].
*/
//go:nosplit
func (self class) AddExceptionRid(rid gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_add_exception_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a collision exception so the shape does not report collisions with the specified [CollisionObject3D] node.
*/
//go:nosplit
func (self class) AddException(node gdclass.CollisionObject3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_add_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a collision exception so the shape does report collisions with the specified [RID].
*/
//go:nosplit
func (self class) RemoveExceptionRid(rid gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_remove_exception_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a collision exception so the shape does report collisions with the specified [CollisionObject3D] node.
*/
//go:nosplit
func (self class) RemoveException(node gdclass.CollisionObject3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_remove_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all collision exceptions for this [ShapeCast3D].
*/
//go:nosplit
func (self class) ClearExceptions()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_clear_exceptions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCollisionMask(mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionMaskValue(layer_number gd.Int, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_set_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionMaskValue(layer_number gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_get_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetExcludeParentBody(mask bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_set_exclude_parent_body, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetExcludeParentBody() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_get_exclude_parent_body, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollideWithAreas(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_set_collide_with_areas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCollideWithAreasEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_is_collide_with_areas_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollideWithBodies(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_set_collide_with_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCollideWithBodiesEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_is_collide_with_bodies_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDebugShapeCustomColor(debug_shape_custom_color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, debug_shape_custom_color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_set_debug_shape_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDebugShapeCustomColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShapeCast3D.Bind_get_debug_shape_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsShapeCast3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsShapeCast3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {classdb.Register("ShapeCast3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
