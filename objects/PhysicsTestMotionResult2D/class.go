package PhysicsTestMotionResult2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Describes the motion and collision result from [method PhysicsServer2D.body_test_motion].
*/
type Instance [1]classdb.PhysicsTestMotionResult2D

/*
Returns the moving object's travel before collision.
*/
func (self Instance) GetTravel() gd.Vector2 {
	return gd.Vector2(class(self).GetTravel())
}

/*
Returns the moving object's remaining movement vector.
*/
func (self Instance) GetRemainder() gd.Vector2 {
	return gd.Vector2(class(self).GetRemainder())
}

/*
Returns the point of collision in global coordinates, if a collision occurred.
*/
func (self Instance) GetCollisionPoint() gd.Vector2 {
	return gd.Vector2(class(self).GetCollisionPoint())
}

/*
Returns the colliding body's shape's normal at the point of collision, if a collision occurred.
*/
func (self Instance) GetCollisionNormal() gd.Vector2 {
	return gd.Vector2(class(self).GetCollisionNormal())
}

/*
Returns the colliding body's velocity, if a collision occurred.
*/
func (self Instance) GetColliderVelocity() gd.Vector2 {
	return gd.Vector2(class(self).GetColliderVelocity())
}

/*
Returns the unique instance ID of the colliding body's attached [Object], if a collision occurred. See [method Object.get_instance_id].
*/
func (self Instance) GetColliderId() int {
	return int(int(class(self).GetColliderId()))
}

/*
Returns the colliding body's [RID] used by the [PhysicsServer2D], if a collision occurred.
*/
func (self Instance) GetColliderRid() gd.RID {
	return gd.RID(class(self).GetColliderRid())
}

/*
Returns the colliding body's attached [Object], if a collision occurred.
*/
func (self Instance) GetCollider() gd.Object {
	return gd.Object(class(self).GetCollider())
}

/*
Returns the colliding body's shape index, if a collision occurred. See [CollisionObject2D].
*/
func (self Instance) GetColliderShape() int {
	return int(int(class(self).GetColliderShape()))
}

/*
Returns the moving object's colliding shape, if a collision occurred.
*/
func (self Instance) GetCollisionLocalShape() int {
	return int(int(class(self).GetCollisionLocalShape()))
}

/*
Returns the length of overlap along the collision normal, if a collision occurred.
*/
func (self Instance) GetCollisionDepth() float64 {
	return float64(float64(class(self).GetCollisionDepth()))
}

/*
Returns the maximum fraction of the motion that can occur without a collision, between [code]0[/code] and [code]1[/code].
*/
func (self Instance) GetCollisionSafeFraction() float64 {
	return float64(float64(class(self).GetCollisionSafeFraction()))
}

/*
Returns the minimum fraction of the motion needed to collide, if a collision occurred, between [code]0[/code] and [code]1[/code].
*/
func (self Instance) GetCollisionUnsafeFraction() float64 {
	return float64(float64(class(self).GetCollisionUnsafeFraction()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.PhysicsTestMotionResult2D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsTestMotionResult2D"))
	return Instance{classdb.PhysicsTestMotionResult2D(object)}
}

/*
Returns the moving object's travel before collision.
*/
//go:nosplit
func (self class) GetTravel() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult2D.Bind_get_travel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the moving object's remaining movement vector.
*/
//go:nosplit
func (self class) GetRemainder() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult2D.Bind_get_remainder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the point of collision in global coordinates, if a collision occurred.
*/
//go:nosplit
func (self class) GetCollisionPoint() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult2D.Bind_get_collision_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the colliding body's shape's normal at the point of collision, if a collision occurred.
*/
//go:nosplit
func (self class) GetCollisionNormal() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult2D.Bind_get_collision_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the colliding body's velocity, if a collision occurred.
*/
//go:nosplit
func (self class) GetColliderVelocity() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult2D.Bind_get_collider_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the unique instance ID of the colliding body's attached [Object], if a collision occurred. See [method Object.get_instance_id].
*/
//go:nosplit
func (self class) GetColliderId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult2D.Bind_get_collider_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the colliding body's [RID] used by the [PhysicsServer2D], if a collision occurred.
*/
//go:nosplit
func (self class) GetColliderRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult2D.Bind_get_collider_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the colliding body's attached [Object], if a collision occurred.
*/
//go:nosplit
func (self class) GetCollider() gd.Object {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult2D.Bind_get_collider, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gd.PointerWithOwnershipTransferredToGo(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the colliding body's shape index, if a collision occurred. See [CollisionObject2D].
*/
//go:nosplit
func (self class) GetColliderShape() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult2D.Bind_get_collider_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the moving object's colliding shape, if a collision occurred.
*/
//go:nosplit
func (self class) GetCollisionLocalShape() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult2D.Bind_get_collision_local_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the length of overlap along the collision normal, if a collision occurred.
*/
//go:nosplit
func (self class) GetCollisionDepth() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult2D.Bind_get_collision_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the maximum fraction of the motion that can occur without a collision, between [code]0[/code] and [code]1[/code].
*/
//go:nosplit
func (self class) GetCollisionSafeFraction() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult2D.Bind_get_collision_safe_fraction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the minimum fraction of the motion needed to collide, if a collision occurred, between [code]0[/code] and [code]1[/code].
*/
//go:nosplit
func (self class) GetCollisionUnsafeFraction() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult2D.Bind_get_collision_unsafe_fraction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPhysicsTestMotionResult2D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsTestMotionResult2D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("PhysicsTestMotionResult2D", func(ptr gd.Object) any { return classdb.PhysicsTestMotionResult2D(ptr) })
}
