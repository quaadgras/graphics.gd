package ShapeCast2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Node2D"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Shape casting allows to detect collision objects by sweeping its [member shape] along the cast direction determined by [member target_position]. This is similar to [RayCast2D], but it allows for sweeping a region of space, rather than just a straight line. [ShapeCast2D] can detect multiple collision objects. It is useful for things like wide laser beams or snapping a simple shape to a floor.
Immediate collision overlaps can be done with the [member target_position] set to [code]Vector2(0, 0)[/code] and by calling [method force_shapecast_update] within the same physics frame. This helps to overcome some limitations of [Area2D] when used as an instantaneous detection area, as collision information isn't immediately available to it.
[b]Note:[/b] Shape casting is more computationally expensive than ray casting.
*/
type Instance [1]classdb.ShapeCast2D
type Any interface {
	gd.IsClass
	AsShapeCast2D() Instance
}

/*
Returns whether any object is intersecting with the shape's vector (considering the vector length).
*/
func (self Instance) IsColliding() bool {
	return bool(class(self).IsColliding())
}

/*
The number of collisions detected at the point of impact. Use this to iterate over multiple collisions as provided by [method get_collider], [method get_collider_shape], [method get_collision_point], and [method get_collision_normal] methods.
*/
func (self Instance) GetCollisionCount() int {
	return int(int(class(self).GetCollisionCount()))
}

/*
Updates the collision information for the shape immediately, without waiting for the next [code]_physics_process[/code] call. Use this method, for example, when the shape or its parent has changed state.
[b]Note:[/b] [code]enabled == true[/code] is not required for this to work.
*/
func (self Instance) ForceShapecastUpdate() {
	class(self).ForceShapecastUpdate()
}

/*
Returns the collided [Object] of one of the multiple collisions at [param index], or [code]null[/code] if no object is intersecting the shape (i.e. [method is_colliding] returns [code]false[/code]).
*/
func (self Instance) GetCollider(index int) gd.Object {
	return gd.Object(class(self).GetCollider(gd.Int(index)))
}

/*
Returns the [RID] of the collided object of one of the multiple collisions at [param index].
*/
func (self Instance) GetColliderRid(index int) Resource.ID {
	return Resource.ID(class(self).GetColliderRid(gd.Int(index)))
}

/*
Returns the shape ID of the colliding shape of one of the multiple collisions at [param index], or [code]0[/code] if no object is intersecting the shape (i.e. [method is_colliding] returns [code]false[/code]).
*/
func (self Instance) GetColliderShape(index int) int {
	return int(int(class(self).GetColliderShape(gd.Int(index))))
}

/*
Returns the collision point of one of the multiple collisions at [param index] where the shape intersects the colliding object.
[b]Note:[/b] this point is in the [b]global[/b] coordinate system.
*/
func (self Instance) GetCollisionPoint(index int) Vector2.XY {
	return Vector2.XY(class(self).GetCollisionPoint(gd.Int(index)))
}

/*
Returns the normal of one of the multiple collisions at [param index] of the intersecting object.
*/
func (self Instance) GetCollisionNormal(index int) Vector2.XY {
	return Vector2.XY(class(self).GetCollisionNormal(gd.Int(index)))
}

/*
The fraction from the [ShapeCast2D]'s origin to its [member target_position] (between 0 and 1) of how far the shape can move without triggering a collision.
*/
func (self Instance) GetClosestCollisionSafeFraction() Float.X {
	return Float.X(Float.X(class(self).GetClosestCollisionSafeFraction()))
}

/*
The fraction from the [ShapeCast2D]'s origin to its [member target_position] (between 0 and 1) of how far the shape must move to trigger a collision.
In ideal conditions this would be the same as [method get_closest_collision_safe_fraction], however shape casting is calculated in discrete steps, so the precise point of collision can occur between two calculated positions.
*/
func (self Instance) GetClosestCollisionUnsafeFraction() Float.X {
	return Float.X(Float.X(class(self).GetClosestCollisionUnsafeFraction()))
}

/*
Adds a collision exception so the shape does not report collisions with the specified [RID].
*/
func (self Instance) AddExceptionRid(rid Resource.ID) {
	class(self).AddExceptionRid(rid)
}

/*
Adds a collision exception so the shape does not report collisions with the specified [CollisionObject2D] node.
*/
func (self Instance) AddException(node objects.CollisionObject2D) {
	class(self).AddException(node)
}

/*
Removes a collision exception so the shape does report collisions with the specified [RID].
*/
func (self Instance) RemoveExceptionRid(rid Resource.ID) {
	class(self).RemoveExceptionRid(rid)
}

/*
Removes a collision exception so the shape does report collisions with the specified [CollisionObject2D] node.
*/
func (self Instance) RemoveException(node objects.CollisionObject2D) {
	class(self).RemoveException(node)
}

/*
Removes all collision exceptions for this shape.
*/
func (self Instance) ClearExceptions() {
	class(self).ClearExceptions()
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetCollisionMaskValue(layer_number int, value bool) {
	class(self).SetCollisionMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetCollisionMaskValue(layer_number int) bool {
	return bool(class(self).GetCollisionMaskValue(gd.Int(layer_number)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ShapeCast2D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ShapeCast2D"))
	return Instance{classdb.ShapeCast2D(object)}
}

func (self Instance) Enabled() bool {
	return bool(class(self).IsEnabled())
}

func (self Instance) SetEnabled(value bool) {
	class(self).SetEnabled(value)
}

func (self Instance) Shape() objects.Shape2D {
	return objects.Shape2D(class(self).GetShape())
}

func (self Instance) SetShape(value objects.Shape2D) {
	class(self).SetShape(value)
}

func (self Instance) ExcludeParent() bool {
	return bool(class(self).GetExcludeParentBody())
}

func (self Instance) SetExcludeParent(value bool) {
	class(self).SetExcludeParentBody(value)
}

func (self Instance) TargetPosition() Vector2.XY {
	return Vector2.XY(class(self).GetTargetPosition())
}

func (self Instance) SetTargetPosition(value Vector2.XY) {
	class(self).SetTargetPosition(gd.Vector2(value))
}

func (self Instance) Margin() Float.X {
	return Float.X(Float.X(class(self).GetMargin()))
}

func (self Instance) SetMargin(value Float.X) {
	class(self).SetMargin(gd.Float(value))
}

func (self Instance) MaxResults() int {
	return int(int(class(self).GetMaxResults()))
}

func (self Instance) SetMaxResults(value int) {
	class(self).SetMaxResults(gd.Int(value))
}

func (self Instance) CollisionMask() int {
	return int(int(class(self).GetCollisionMask()))
}

func (self Instance) SetCollisionMask(value int) {
	class(self).SetCollisionMask(gd.Int(value))
}

func (self Instance) CollideWithAreas() bool {
	return bool(class(self).IsCollideWithAreasEnabled())
}

func (self Instance) SetCollideWithAreas(value bool) {
	class(self).SetCollideWithAreas(value)
}

func (self Instance) CollideWithBodies() bool {
	return bool(class(self).IsCollideWithBodiesEnabled())
}

func (self Instance) SetCollideWithBodies(value bool) {
	class(self).SetCollideWithBodies(value)
}

//go:nosplit
func (self class) SetEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_is_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShape(shape objects.Shape2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shape[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_set_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetShape() objects.Shape2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_get_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Shape2D{classdb.Shape2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTargetPosition(local_point gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, local_point)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_set_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTargetPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_get_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMargin(margin gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_set_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMargin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_get_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxResults(max_results gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, max_results)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_set_max_results, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxResults() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_get_max_results, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether any object is intersecting with the shape's vector (considering the vector length).
*/
//go:nosplit
func (self class) IsColliding() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_is_colliding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
The number of collisions detected at the point of impact. Use this to iterate over multiple collisions as provided by [method get_collider], [method get_collider_shape], [method get_collision_point], and [method get_collision_normal] methods.
*/
//go:nosplit
func (self class) GetCollisionCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_get_collision_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Updates the collision information for the shape immediately, without waiting for the next [code]_physics_process[/code] call. Use this method, for example, when the shape or its parent has changed state.
[b]Note:[/b] [code]enabled == true[/code] is not required for this to work.
*/
//go:nosplit
func (self class) ForceShapecastUpdate() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_force_shapecast_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the collided [Object] of one of the multiple collisions at [param index], or [code]null[/code] if no object is intersecting the shape (i.e. [method is_colliding] returns [code]false[/code]).
*/
//go:nosplit
func (self class) GetCollider(index gd.Int) gd.Object {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_get_collider, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gd.PointerWithOwnershipTransferredToGo(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the [RID] of the collided object of one of the multiple collisions at [param index].
*/
//go:nosplit
func (self class) GetColliderRid(index gd.Int) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_get_collider_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the shape ID of the colliding shape of one of the multiple collisions at [param index], or [code]0[/code] if no object is intersecting the shape (i.e. [method is_colliding] returns [code]false[/code]).
*/
//go:nosplit
func (self class) GetColliderShape(index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_get_collider_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the collision point of one of the multiple collisions at [param index] where the shape intersects the colliding object.
[b]Note:[/b] this point is in the [b]global[/b] coordinate system.
*/
//go:nosplit
func (self class) GetCollisionPoint(index gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_get_collision_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the normal of one of the multiple collisions at [param index] of the intersecting object.
*/
//go:nosplit
func (self class) GetCollisionNormal(index gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_get_collision_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
The fraction from the [ShapeCast2D]'s origin to its [member target_position] (between 0 and 1) of how far the shape can move without triggering a collision.
*/
//go:nosplit
func (self class) GetClosestCollisionSafeFraction() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_get_closest_collision_safe_fraction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
The fraction from the [ShapeCast2D]'s origin to its [member target_position] (between 0 and 1) of how far the shape must move to trigger a collision.
In ideal conditions this would be the same as [method get_closest_collision_safe_fraction], however shape casting is calculated in discrete steps, so the precise point of collision can occur between two calculated positions.
*/
//go:nosplit
func (self class) GetClosestCollisionUnsafeFraction() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_get_closest_collision_unsafe_fraction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a collision exception so the shape does not report collisions with the specified [RID].
*/
//go:nosplit
func (self class) AddExceptionRid(rid gd.RID) {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_add_exception_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a collision exception so the shape does not report collisions with the specified [CollisionObject2D] node.
*/
//go:nosplit
func (self class) AddException(node objects.CollisionObject2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_add_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes a collision exception so the shape does report collisions with the specified [RID].
*/
//go:nosplit
func (self class) RemoveExceptionRid(rid gd.RID) {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_remove_exception_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes a collision exception so the shape does report collisions with the specified [CollisionObject2D] node.
*/
//go:nosplit
func (self class) RemoveException(node objects.CollisionObject2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_remove_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes all collision exceptions for this shape.
*/
//go:nosplit
func (self class) ClearExceptions() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_clear_exceptions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetCollisionMask(mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionMaskValue(layer_number gd.Int, value bool) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_set_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionMaskValue(layer_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_get_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExcludeParentBody(mask bool) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_set_exclude_parent_body, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetExcludeParentBody() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_get_exclude_parent_body, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollideWithAreas(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_set_collide_with_areas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsCollideWithAreasEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_is_collide_with_areas_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollideWithBodies(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_set_collide_with_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsCollideWithBodiesEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShapeCast2D.Bind_is_collide_with_bodies_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsShapeCast2D() Advanced      { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsShapeCast2D() Instance   { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced    { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {
	classdb.Register("ShapeCast2D", func(ptr gd.Object) any { return [1]classdb.ShapeCast2D{classdb.ShapeCast2D(ptr)} })
}
