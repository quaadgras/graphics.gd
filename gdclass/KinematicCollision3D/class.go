package KinematicCollision3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Holds collision data from the movement of a [PhysicsBody3D], usually from [method PhysicsBody3D.move_and_collide]. When a [PhysicsBody3D] is moved, it stops if it detects a collision with another body. If a collision is detected, a [KinematicCollision3D] object is returned.
The collision data includes the colliding object, the remaining motion, and the collision position. This data can be used to determine a custom response to the collision.

*/
type Go [1]classdb.KinematicCollision3D

/*
Returns the moving object's travel before collision.
*/
func (self Go) GetTravel() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetTravel())
}

/*
Returns the moving object's remaining movement vector.
*/
func (self Go) GetRemainder() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetRemainder())
}

/*
Returns the colliding body's length of overlap along the collision normal.
*/
func (self Go) GetDepth() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetDepth()))
}

/*
Returns the number of detected collisions.
*/
func (self Go) GetCollisionCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCollisionCount()))
}

/*
Returns the point of collision in global coordinates given a collision index (the deepest collision by default).
*/
func (self Go) GetPosition() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetPosition(gd.Int(0)))
}

/*
Returns the colliding body's shape's normal at the point of collision given a collision index (the deepest collision by default).
*/
func (self Go) GetNormal() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetNormal(gd.Int(0)))
}

/*
Returns the collision angle according to [param up_direction], which is [constant Vector3.UP] by default. This value is always positive.
*/
func (self Go) GetAngle() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetAngle(gd.Int(0), gd.Vector3{0, 1, 0})))
}

/*
Returns the moving object's colliding shape given a collision index (the deepest collision by default).
*/
func (self Go) GetLocalShape() gd.Object {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Object(class(self).GetLocalShape(gc, gd.Int(0)))
}

/*
Returns the colliding body's attached [Object] given a collision index (the deepest collision by default).
*/
func (self Go) GetCollider() gd.Object {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Object(class(self).GetCollider(gc, gd.Int(0)))
}

/*
Returns the unique instance ID of the colliding body's attached [Object] given a collision index (the deepest collision by default). See [method Object.get_instance_id].
*/
func (self Go) GetColliderId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetColliderId(gd.Int(0))))
}

/*
Returns the colliding body's [RID] used by the [PhysicsServer3D] given a collision index (the deepest collision by default).
*/
func (self Go) GetColliderRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetColliderRid(gd.Int(0)))
}

/*
Returns the colliding body's shape given a collision index (the deepest collision by default).
*/
func (self Go) GetColliderShape() gd.Object {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Object(class(self).GetColliderShape(gc, gd.Int(0)))
}

/*
Returns the colliding body's shape index given a collision index (the deepest collision by default). See [CollisionObject3D].
*/
func (self Go) GetColliderShapeIndex() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetColliderShapeIndex(gd.Int(0))))
}

/*
Returns the colliding body's velocity given a collision index (the deepest collision by default).
*/
func (self Go) GetColliderVelocity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetColliderVelocity(gd.Int(0)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.KinematicCollision3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("KinematicCollision3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns the moving object's travel before collision.
*/
//go:nosplit
func (self class) GetTravel() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision3D.Bind_get_travel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the moving object's remaining movement vector.
*/
//go:nosplit
func (self class) GetRemainder() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision3D.Bind_get_remainder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the colliding body's length of overlap along the collision normal.
*/
//go:nosplit
func (self class) GetDepth() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision3D.Bind_get_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of detected collisions.
*/
//go:nosplit
func (self class) GetCollisionCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision3D.Bind_get_collision_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the point of collision in global coordinates given a collision index (the deepest collision by default).
*/
//go:nosplit
func (self class) GetPosition(collision_index gd.Int) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision3D.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the colliding body's shape's normal at the point of collision given a collision index (the deepest collision by default).
*/
//go:nosplit
func (self class) GetNormal(collision_index gd.Int) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision3D.Bind_get_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the collision angle according to [param up_direction], which is [constant Vector3.UP] by default. This value is always positive.
*/
//go:nosplit
func (self class) GetAngle(collision_index gd.Int, up_direction gd.Vector3) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	callframe.Arg(frame, up_direction)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision3D.Bind_get_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the moving object's colliding shape given a collision index (the deepest collision by default).
*/
//go:nosplit
func (self class) GetLocalShape(ctx gd.Lifetime, collision_index gd.Int) gd.Object {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision3D.Bind_get_local_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gd.Object
	ret.SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the colliding body's attached [Object] given a collision index (the deepest collision by default).
*/
//go:nosplit
func (self class) GetCollider(ctx gd.Lifetime, collision_index gd.Int) gd.Object {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision3D.Bind_get_collider, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gd.Object
	ret.SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the unique instance ID of the colliding body's attached [Object] given a collision index (the deepest collision by default). See [method Object.get_instance_id].
*/
//go:nosplit
func (self class) GetColliderId(collision_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision3D.Bind_get_collider_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the colliding body's [RID] used by the [PhysicsServer3D] given a collision index (the deepest collision by default).
*/
//go:nosplit
func (self class) GetColliderRid(collision_index gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision3D.Bind_get_collider_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the colliding body's shape given a collision index (the deepest collision by default).
*/
//go:nosplit
func (self class) GetColliderShape(ctx gd.Lifetime, collision_index gd.Int) gd.Object {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision3D.Bind_get_collider_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gd.Object
	ret.SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the colliding body's shape index given a collision index (the deepest collision by default). See [CollisionObject3D].
*/
//go:nosplit
func (self class) GetColliderShapeIndex(collision_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision3D.Bind_get_collider_shape_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the colliding body's velocity given a collision index (the deepest collision by default).
*/
//go:nosplit
func (self class) GetColliderVelocity(collision_index gd.Int) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision3D.Bind_get_collider_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsKinematicCollision3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsKinematicCollision3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("KinematicCollision3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
