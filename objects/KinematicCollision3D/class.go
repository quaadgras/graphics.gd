package KinematicCollision3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Float"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Holds collision data from the movement of a [PhysicsBody3D], usually from [method PhysicsBody3D.move_and_collide]. When a [PhysicsBody3D] is moved, it stops if it detects a collision with another body. If a collision is detected, a [KinematicCollision3D] object is returned.
The collision data includes the colliding object, the remaining motion, and the collision position. This data can be used to determine a custom response to the collision.
*/
type Instance [1]classdb.KinematicCollision3D
type Any interface {
	gd.IsClass
	AsKinematicCollision3D() Instance
}

/*
Returns the moving object's travel before collision.
*/
func (self Instance) GetTravel() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetTravel())
}

/*
Returns the moving object's remaining movement vector.
*/
func (self Instance) GetRemainder() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetRemainder())
}

/*
Returns the colliding body's length of overlap along the collision normal.
*/
func (self Instance) GetDepth() Float.X {
	return Float.X(Float.X(class(self).GetDepth()))
}

/*
Returns the number of detected collisions.
*/
func (self Instance) GetCollisionCount() int {
	return int(int(class(self).GetCollisionCount()))
}

/*
Returns the point of collision in global coordinates given a collision index (the deepest collision by default).
*/
func (self Instance) GetPosition() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetPosition(gd.Int(0)))
}

/*
Returns the colliding body's shape's normal at the point of collision given a collision index (the deepest collision by default).
*/
func (self Instance) GetNormal() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetNormal(gd.Int(0)))
}

/*
Returns the collision angle according to [param up_direction], which is [constant Vector3.UP] by default. This value is always positive.
*/
func (self Instance) GetAngle() Float.X {
	return Float.X(Float.X(class(self).GetAngle(gd.Int(0), gd.Vector3(gd.Vector3{0, 1, 0}))))
}

/*
Returns the moving object's colliding shape given a collision index (the deepest collision by default).
*/
func (self Instance) GetLocalShape() gd.Object {
	return gd.Object(class(self).GetLocalShape(gd.Int(0)))
}

/*
Returns the colliding body's attached [Object] given a collision index (the deepest collision by default).
*/
func (self Instance) GetCollider() gd.Object {
	return gd.Object(class(self).GetCollider(gd.Int(0)))
}

/*
Returns the unique instance ID of the colliding body's attached [Object] given a collision index (the deepest collision by default). See [method Object.get_instance_id].
*/
func (self Instance) GetColliderId() int {
	return int(int(class(self).GetColliderId(gd.Int(0))))
}

/*
Returns the colliding body's [RID] used by the [PhysicsServer3D] given a collision index (the deepest collision by default).
*/
func (self Instance) GetColliderRid() Resource.ID {
	return Resource.ID(class(self).GetColliderRid(gd.Int(0)))
}

/*
Returns the colliding body's shape given a collision index (the deepest collision by default).
*/
func (self Instance) GetColliderShape() gd.Object {
	return gd.Object(class(self).GetColliderShape(gd.Int(0)))
}

/*
Returns the colliding body's shape index given a collision index (the deepest collision by default). See [CollisionObject3D].
*/
func (self Instance) GetColliderShapeIndex() int {
	return int(int(class(self).GetColliderShapeIndex(gd.Int(0))))
}

/*
Returns the colliding body's velocity given a collision index (the deepest collision by default).
*/
func (self Instance) GetColliderVelocity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetColliderVelocity(gd.Int(0)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.KinematicCollision3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("KinematicCollision3D"))
	return Instance{*(*classdb.KinematicCollision3D)(unsafe.Pointer(&object))}
}

/*
Returns the moving object's travel before collision.
*/
//go:nosplit
func (self class) GetTravel() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision3D.Bind_get_travel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the moving object's remaining movement vector.
*/
//go:nosplit
func (self class) GetRemainder() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision3D.Bind_get_remainder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the colliding body's length of overlap along the collision normal.
*/
//go:nosplit
func (self class) GetDepth() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision3D.Bind_get_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of detected collisions.
*/
//go:nosplit
func (self class) GetCollisionCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision3D.Bind_get_collision_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the point of collision in global coordinates given a collision index (the deepest collision by default).
*/
//go:nosplit
func (self class) GetPosition(collision_index gd.Int) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision3D.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the colliding body's shape's normal at the point of collision given a collision index (the deepest collision by default).
*/
//go:nosplit
func (self class) GetNormal(collision_index gd.Int) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision3D.Bind_get_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the collision angle according to [param up_direction], which is [constant Vector3.UP] by default. This value is always positive.
*/
//go:nosplit
func (self class) GetAngle(collision_index gd.Int, up_direction gd.Vector3) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	callframe.Arg(frame, up_direction)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision3D.Bind_get_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the moving object's colliding shape given a collision index (the deepest collision by default).
*/
//go:nosplit
func (self class) GetLocalShape(collision_index gd.Int) gd.Object {
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision3D.Bind_get_local_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gd.PointerWithOwnershipTransferredToGo[gd.Object](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the colliding body's attached [Object] given a collision index (the deepest collision by default).
*/
//go:nosplit
func (self class) GetCollider(collision_index gd.Int) gd.Object {
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision3D.Bind_get_collider, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gd.PointerWithOwnershipTransferredToGo[gd.Object](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the unique instance ID of the colliding body's attached [Object] given a collision index (the deepest collision by default). See [method Object.get_instance_id].
*/
//go:nosplit
func (self class) GetColliderId(collision_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision3D.Bind_get_collider_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the colliding body's [RID] used by the [PhysicsServer3D] given a collision index (the deepest collision by default).
*/
//go:nosplit
func (self class) GetColliderRid(collision_index gd.Int) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision3D.Bind_get_collider_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the colliding body's shape given a collision index (the deepest collision by default).
*/
//go:nosplit
func (self class) GetColliderShape(collision_index gd.Int) gd.Object {
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision3D.Bind_get_collider_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gd.PointerWithOwnershipTransferredToGo[gd.Object](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the colliding body's shape index given a collision index (the deepest collision by default). See [CollisionObject3D].
*/
//go:nosplit
func (self class) GetColliderShapeIndex(collision_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision3D.Bind_get_collider_shape_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the colliding body's velocity given a collision index (the deepest collision by default).
*/
//go:nosplit
func (self class) GetColliderVelocity(collision_index gd.Int) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision3D.Bind_get_collider_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsKinematicCollision3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsKinematicCollision3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted         { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted      { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

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
	classdb.Register("KinematicCollision3D", func(ptr gd.Object) any {
		return [1]classdb.KinematicCollision3D{*(*classdb.KinematicCollision3D)(unsafe.Pointer(&ptr))}
	})
}
