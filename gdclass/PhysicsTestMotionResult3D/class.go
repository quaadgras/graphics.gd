package PhysicsTestMotionResult3D

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
Describes the motion and collision result from [method PhysicsServer3D.body_test_motion].

*/
type Go [1]classdb.PhysicsTestMotionResult3D

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
Returns the maximum fraction of the motion that can occur without a collision, between [code]0[/code] and [code]1[/code].
*/
func (self Go) GetCollisionSafeFraction() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetCollisionSafeFraction()))
}

/*
Returns the minimum fraction of the motion needed to collide, if a collision occurred, between [code]0[/code] and [code]1[/code].
*/
func (self Go) GetCollisionUnsafeFraction() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetCollisionUnsafeFraction()))
}

/*
Returns the number of detected collisions.
*/
func (self Go) GetCollisionCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCollisionCount()))
}

/*
Returns the point of collision in global coordinates given a collision index (the deepest collision by default), if a collision occurred.
*/
func (self Go) GetCollisionPoint() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetCollisionPoint(gd.Int(0)))
}

/*
Returns the colliding body's shape's normal at the point of collision given a collision index (the deepest collision by default), if a collision occurred.
*/
func (self Go) GetCollisionNormal() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetCollisionNormal(gd.Int(0)))
}

/*
Returns the colliding body's velocity given a collision index (the deepest collision by default), if a collision occurred.
*/
func (self Go) GetColliderVelocity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetColliderVelocity(gd.Int(0)))
}

/*
Returns the unique instance ID of the colliding body's attached [Object] given a collision index (the deepest collision by default), if a collision occurred. See [method Object.get_instance_id].
*/
func (self Go) GetColliderId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetColliderId(gd.Int(0))))
}

/*
Returns the colliding body's [RID] used by the [PhysicsServer3D] given a collision index (the deepest collision by default), if a collision occurred.
*/
func (self Go) GetColliderRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetColliderRid(gd.Int(0)))
}

/*
Returns the colliding body's attached [Object] given a collision index (the deepest collision by default), if a collision occurred.
*/
func (self Go) GetCollider() gd.Object {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Object(class(self).GetCollider(gc, gd.Int(0)))
}

/*
Returns the colliding body's shape index given a collision index (the deepest collision by default), if a collision occurred. See [CollisionObject3D].
*/
func (self Go) GetColliderShape() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetColliderShape(gd.Int(0))))
}

/*
Returns the moving object's colliding shape given a collision index (the deepest collision by default), if a collision occurred.
*/
func (self Go) GetCollisionLocalShape() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCollisionLocalShape(gd.Int(0))))
}

/*
Returns the length of overlap along the collision normal given a collision index (the deepest collision by default), if a collision occurred.
*/
func (self Go) GetCollisionDepth() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetCollisionDepth(gd.Int(0))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PhysicsTestMotionResult3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PhysicsTestMotionResult3D"))
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionResult3D.Bind_get_travel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionResult3D.Bind_get_remainder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the maximum fraction of the motion that can occur without a collision, between [code]0[/code] and [code]1[/code].
*/
//go:nosplit
func (self class) GetCollisionSafeFraction() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionResult3D.Bind_get_collision_safe_fraction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the minimum fraction of the motion needed to collide, if a collision occurred, between [code]0[/code] and [code]1[/code].
*/
//go:nosplit
func (self class) GetCollisionUnsafeFraction() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionResult3D.Bind_get_collision_unsafe_fraction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionResult3D.Bind_get_collision_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the point of collision in global coordinates given a collision index (the deepest collision by default), if a collision occurred.
*/
//go:nosplit
func (self class) GetCollisionPoint(collision_index gd.Int) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionResult3D.Bind_get_collision_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the colliding body's shape's normal at the point of collision given a collision index (the deepest collision by default), if a collision occurred.
*/
//go:nosplit
func (self class) GetCollisionNormal(collision_index gd.Int) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionResult3D.Bind_get_collision_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the colliding body's velocity given a collision index (the deepest collision by default), if a collision occurred.
*/
//go:nosplit
func (self class) GetColliderVelocity(collision_index gd.Int) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionResult3D.Bind_get_collider_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the unique instance ID of the colliding body's attached [Object] given a collision index (the deepest collision by default), if a collision occurred. See [method Object.get_instance_id].
*/
//go:nosplit
func (self class) GetColliderId(collision_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionResult3D.Bind_get_collider_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the colliding body's [RID] used by the [PhysicsServer3D] given a collision index (the deepest collision by default), if a collision occurred.
*/
//go:nosplit
func (self class) GetColliderRid(collision_index gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionResult3D.Bind_get_collider_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the colliding body's attached [Object] given a collision index (the deepest collision by default), if a collision occurred.
*/
//go:nosplit
func (self class) GetCollider(ctx gd.Lifetime, collision_index gd.Int) gd.Object {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionResult3D.Bind_get_collider, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gd.Object
	ret.SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the colliding body's shape index given a collision index (the deepest collision by default), if a collision occurred. See [CollisionObject3D].
*/
//go:nosplit
func (self class) GetColliderShape(collision_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionResult3D.Bind_get_collider_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the moving object's colliding shape given a collision index (the deepest collision by default), if a collision occurred.
*/
//go:nosplit
func (self class) GetCollisionLocalShape(collision_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionResult3D.Bind_get_collision_local_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the length of overlap along the collision normal given a collision index (the deepest collision by default), if a collision occurred.
*/
//go:nosplit
func (self class) GetCollisionDepth(collision_index gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsTestMotionResult3D.Bind_get_collision_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPhysicsTestMotionResult3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPhysicsTestMotionResult3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("PhysicsTestMotionResult3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
