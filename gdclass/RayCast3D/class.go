package RayCast3D

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
A raycast represents a ray from its origin to its [member target_position] that finds the closest [CollisionObject3D] along its path, if it intersects any.
[RayCast3D] can ignore some objects by adding them to an exception list, by making its detection reporting ignore [Area3D]s ([member collide_with_areas]) or [PhysicsBody3D]s ([member collide_with_bodies]), or by configuring physics layers.
[RayCast3D] calculates intersection every physics frame, and it holds the result until the next physics frame. For an immediate raycast, or if you want to configure a [RayCast3D] multiple times within the same physics frame, use [method force_raycast_update].
To sweep over a region of 3D space, you can approximate the region with multiple [RayCast3D]s or use [ShapeCast3D].

*/
type Go [1]classdb.RayCast3D

/*
Returns whether any object is intersecting with the ray's vector (considering the vector length).
*/
func (self Go) IsColliding() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsColliding())
}

/*
Updates the collision information for the ray immediately, without waiting for the next [code]_physics_process[/code] call. Use this method, for example, when the ray or its parent has changed state.
[b]Note:[/b] [member enabled] does not need to be [code]true[/code] for this to work.
*/
func (self Go) ForceRaycastUpdate() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ForceRaycastUpdate()
}

/*
Returns the first object that the ray intersects, or [code]null[/code] if no object is intersecting the ray (i.e. [method is_colliding] returns [code]false[/code]).
*/
func (self Go) GetCollider() gd.Object {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Object(class(self).GetCollider(gc))
}

/*
Returns the [RID] of the first object that the ray intersects, or an empty [RID] if no object is intersecting the ray (i.e. [method is_colliding] returns [code]false[/code]).
*/
func (self Go) GetColliderRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetColliderRid())
}

/*
Returns the shape ID of the first object that the ray intersects, or [code]0[/code] if no object is intersecting the ray (i.e. [method is_colliding] returns [code]false[/code]).
To get the intersected shape node, for a [CollisionObject3D] target, use:
[codeblocks]
[gdscript]
var target = get_collider() # A CollisionObject3D.
var shape_id = get_collider_shape() # The shape index in the collider.
var owner_id = target.shape_find_owner(shape_id) # The owner ID in the collider.
var shape = target.shape_owner_get_owner(owner_id)
[/gdscript]
[csharp]
var target = (CollisionObject3D)GetCollider(); // A CollisionObject3D.
var shapeId = GetColliderShape(); // The shape index in the collider.
var ownerId = target.ShapeFindOwner(shapeId); // The owner ID in the collider.
var shape = target.ShapeOwnerGetOwner(ownerId);
[/csharp]
[/codeblocks]
*/
func (self Go) GetColliderShape() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetColliderShape()))
}

/*
Returns the collision point at which the ray intersects the closest object, in the global coordinate system. If [member hit_from_inside] is [code]true[/code] and the ray starts inside of a collision shape, this function will return the origin point of the ray.
[b]Note:[/b] Check that [method is_colliding] returns [code]true[/code] before calling this method to ensure the returned point is valid and up-to-date.
*/
func (self Go) GetCollisionPoint() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetCollisionPoint())
}

/*
Returns the normal of the intersecting object's shape at the collision point, or [code]Vector3(0, 0, 0)[/code] if the ray starts inside the shape and [member hit_from_inside] is [code]true[/code].
[b]Note:[/b] Check that [method is_colliding] returns [code]true[/code] before calling this method to ensure the returned normal is valid and up-to-date.
*/
func (self Go) GetCollisionNormal() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetCollisionNormal())
}

/*
Returns the collision object's face index at the collision point, or [code]-1[/code] if the shape intersecting the ray is not a [ConcavePolygonShape3D].
*/
func (self Go) GetCollisionFaceIndex() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCollisionFaceIndex()))
}

/*
Adds a collision exception so the ray does not report collisions with the specified [RID].
*/
func (self Go) AddExceptionRid(rid gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddExceptionRid(rid)
}

/*
Adds a collision exception so the ray does not report collisions with the specified [CollisionObject3D] node.
*/
func (self Go) AddException(node gdclass.CollisionObject3D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddException(node)
}

/*
Removes a collision exception so the ray does report collisions with the specified [RID].
*/
func (self Go) RemoveExceptionRid(rid gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveExceptionRid(rid)
}

/*
Removes a collision exception so the ray does report collisions with the specified [CollisionObject3D] node.
*/
func (self Go) RemoveException(node gdclass.CollisionObject3D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveException(node)
}

/*
Removes all collision exceptions for this ray.
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
type class [1]classdb.RayCast3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RayCast3D"))
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

func (self Go) CollisionMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCollisionMask()))
}

func (self Go) SetCollisionMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionMask(gd.Int(value))
}

func (self Go) HitFromInside() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsHitFromInsideEnabled())
}

func (self Go) SetHitFromInside(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHitFromInside(value)
}

func (self Go) HitBackFaces() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsHitBackFacesEnabled())
}

func (self Go) SetHitBackFaces(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHitBackFaces(value)
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

func (self Go) DebugShapeThickness() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetDebugShapeThickness()))
}

func (self Go) SetDebugShapeThickness(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDebugShapeThickness(gd.Int(value))
}

//go:nosplit
func (self class) SetEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_is_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTargetPosition(local_point gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, local_point)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_set_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTargetPosition() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_get_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether any object is intersecting with the ray's vector (considering the vector length).
*/
//go:nosplit
func (self class) IsColliding() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_is_colliding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Updates the collision information for the ray immediately, without waiting for the next [code]_physics_process[/code] call. Use this method, for example, when the ray or its parent has changed state.
[b]Note:[/b] [member enabled] does not need to be [code]true[/code] for this to work.
*/
//go:nosplit
func (self class) ForceRaycastUpdate()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_force_raycast_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the first object that the ray intersects, or [code]null[/code] if no object is intersecting the ray (i.e. [method is_colliding] returns [code]false[/code]).
*/
//go:nosplit
func (self class) GetCollider(ctx gd.Lifetime) gd.Object {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_get_collider, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gd.Object
	ret.SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [RID] of the first object that the ray intersects, or an empty [RID] if no object is intersecting the ray (i.e. [method is_colliding] returns [code]false[/code]).
*/
//go:nosplit
func (self class) GetColliderRid() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_get_collider_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the shape ID of the first object that the ray intersects, or [code]0[/code] if no object is intersecting the ray (i.e. [method is_colliding] returns [code]false[/code]).
To get the intersected shape node, for a [CollisionObject3D] target, use:
[codeblocks]
[gdscript]
var target = get_collider() # A CollisionObject3D.
var shape_id = get_collider_shape() # The shape index in the collider.
var owner_id = target.shape_find_owner(shape_id) # The owner ID in the collider.
var shape = target.shape_owner_get_owner(owner_id)
[/gdscript]
[csharp]
var target = (CollisionObject3D)GetCollider(); // A CollisionObject3D.
var shapeId = GetColliderShape(); // The shape index in the collider.
var ownerId = target.ShapeFindOwner(shapeId); // The owner ID in the collider.
var shape = target.ShapeOwnerGetOwner(ownerId);
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetColliderShape() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_get_collider_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the collision point at which the ray intersects the closest object, in the global coordinate system. If [member hit_from_inside] is [code]true[/code] and the ray starts inside of a collision shape, this function will return the origin point of the ray.
[b]Note:[/b] Check that [method is_colliding] returns [code]true[/code] before calling this method to ensure the returned point is valid and up-to-date.
*/
//go:nosplit
func (self class) GetCollisionPoint() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_get_collision_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the normal of the intersecting object's shape at the collision point, or [code]Vector3(0, 0, 0)[/code] if the ray starts inside the shape and [member hit_from_inside] is [code]true[/code].
[b]Note:[/b] Check that [method is_colliding] returns [code]true[/code] before calling this method to ensure the returned normal is valid and up-to-date.
*/
//go:nosplit
func (self class) GetCollisionNormal() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_get_collision_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the collision object's face index at the collision point, or [code]-1[/code] if the shape intersecting the ray is not a [ConcavePolygonShape3D].
*/
//go:nosplit
func (self class) GetCollisionFaceIndex() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_get_collision_face_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a collision exception so the ray does not report collisions with the specified [RID].
*/
//go:nosplit
func (self class) AddExceptionRid(rid gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_add_exception_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a collision exception so the ray does not report collisions with the specified [CollisionObject3D] node.
*/
//go:nosplit
func (self class) AddException(node gdclass.CollisionObject3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_add_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a collision exception so the ray does report collisions with the specified [RID].
*/
//go:nosplit
func (self class) RemoveExceptionRid(rid gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_remove_exception_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a collision exception so the ray does report collisions with the specified [CollisionObject3D] node.
*/
//go:nosplit
func (self class) RemoveException(node gdclass.CollisionObject3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_remove_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all collision exceptions for this ray.
*/
//go:nosplit
func (self class) ClearExceptions()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_clear_exceptions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCollisionMask(mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_set_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_get_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_set_exclude_parent_body, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetExcludeParentBody() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_get_exclude_parent_body, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_set_collide_with_areas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCollideWithAreasEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_is_collide_with_areas_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_set_collide_with_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCollideWithBodiesEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_is_collide_with_bodies_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHitFromInside(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_set_hit_from_inside, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsHitFromInsideEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_is_hit_from_inside_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHitBackFaces(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_set_hit_back_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsHitBackFacesEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_is_hit_back_faces_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_set_debug_shape_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDebugShapeCustomColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_get_debug_shape_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDebugShapeThickness(debug_shape_thickness gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, debug_shape_thickness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_set_debug_shape_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDebugShapeThickness() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RayCast3D.Bind_get_debug_shape_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRayCast3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRayCast3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("RayCast3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
