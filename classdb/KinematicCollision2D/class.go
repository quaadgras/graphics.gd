// Package KinematicCollision2D provides methods for working with KinematicCollision2D object instances.
package KinematicCollision2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable

/*
Holds collision data from the movement of a [PhysicsBody2D], usually from [method PhysicsBody2D.move_and_collide]. When a [PhysicsBody2D] is moved, it stops if it detects a collision with another body. If a collision is detected, a [KinematicCollision2D] object is returned.
The collision data includes the colliding object, the remaining motion, and the collision position. This data can be used to determine a custom response to the collision.
*/
type Instance [1]gdclass.KinematicCollision2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsKinematicCollision2D() Instance
}

/*
Returns the point of collision in global coordinates.
*/
func (self Instance) GetPosition() Vector2.XY { //gd:KinematicCollision2D.get_position
	return Vector2.XY(class(self).GetPosition())
}

/*
Returns the colliding body's shape's normal at the point of collision.
*/
func (self Instance) GetNormal() Vector2.XY { //gd:KinematicCollision2D.get_normal
	return Vector2.XY(class(self).GetNormal())
}

/*
Returns the moving object's travel before collision.
*/
func (self Instance) GetTravel() Vector2.XY { //gd:KinematicCollision2D.get_travel
	return Vector2.XY(class(self).GetTravel())
}

/*
Returns the moving object's remaining movement vector.
*/
func (self Instance) GetRemainder() Vector2.XY { //gd:KinematicCollision2D.get_remainder
	return Vector2.XY(class(self).GetRemainder())
}

/*
Returns the collision angle according to [param up_direction], which is [constant Vector2.UP] by default. This value is always positive.
*/
func (self Instance) GetAngle() Float.X { //gd:KinematicCollision2D.get_angle
	return Float.X(Float.X(class(self).GetAngle(gd.Vector2(gd.Vector2{0, -1}))))
}

/*
Returns the colliding body's length of overlap along the collision normal.
*/
func (self Instance) GetDepth() Float.X { //gd:KinematicCollision2D.get_depth
	return Float.X(Float.X(class(self).GetDepth()))
}

/*
Returns the moving object's colliding shape.
*/
func (self Instance) GetLocalShape() Object.Instance { //gd:KinematicCollision2D.get_local_shape
	return Object.Instance(class(self).GetLocalShape())
}

/*
Returns the colliding body's attached [Object].
*/
func (self Instance) GetCollider() Object.Instance { //gd:KinematicCollision2D.get_collider
	return Object.Instance(class(self).GetCollider())
}

/*
Returns the unique instance ID of the colliding body's attached [Object]. See [method Object.get_instance_id].
*/
func (self Instance) GetColliderId() int { //gd:KinematicCollision2D.get_collider_id
	return int(int(class(self).GetColliderId()))
}

/*
Returns the colliding body's [RID] used by the [PhysicsServer2D].
*/
func (self Instance) GetColliderRid() RID.Body2D { //gd:KinematicCollision2D.get_collider_rid
	return RID.Body2D(class(self).GetColliderRid())
}

/*
Returns the colliding body's shape.
*/
func (self Instance) GetColliderShape() Object.Instance { //gd:KinematicCollision2D.get_collider_shape
	return Object.Instance(class(self).GetColliderShape())
}

/*
Returns the colliding body's shape index. See [CollisionObject2D].
*/
func (self Instance) GetColliderShapeIndex() int { //gd:KinematicCollision2D.get_collider_shape_index
	return int(int(class(self).GetColliderShapeIndex()))
}

/*
Returns the colliding body's velocity.
*/
func (self Instance) GetColliderVelocity() Vector2.XY { //gd:KinematicCollision2D.get_collider_velocity
	return Vector2.XY(class(self).GetColliderVelocity())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.KinematicCollision2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("KinematicCollision2D"))
	casted := Instance{*(*gdclass.KinematicCollision2D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Returns the point of collision in global coordinates.
*/
//go:nosplit
func (self class) GetPosition() gd.Vector2 { //gd:KinematicCollision2D.get_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision2D.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the colliding body's shape's normal at the point of collision.
*/
//go:nosplit
func (self class) GetNormal() gd.Vector2 { //gd:KinematicCollision2D.get_normal
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision2D.Bind_get_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the moving object's travel before collision.
*/
//go:nosplit
func (self class) GetTravel() gd.Vector2 { //gd:KinematicCollision2D.get_travel
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision2D.Bind_get_travel, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the moving object's remaining movement vector.
*/
//go:nosplit
func (self class) GetRemainder() gd.Vector2 { //gd:KinematicCollision2D.get_remainder
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision2D.Bind_get_remainder, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the collision angle according to [param up_direction], which is [constant Vector2.UP] by default. This value is always positive.
*/
//go:nosplit
func (self class) GetAngle(up_direction gd.Vector2) gd.Float { //gd:KinematicCollision2D.get_angle
	var frame = callframe.New()
	callframe.Arg(frame, up_direction)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision2D.Bind_get_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the colliding body's length of overlap along the collision normal.
*/
//go:nosplit
func (self class) GetDepth() gd.Float { //gd:KinematicCollision2D.get_depth
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision2D.Bind_get_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the moving object's colliding shape.
*/
//go:nosplit
func (self class) GetLocalShape() [1]gd.Object { //gd:KinematicCollision2D.get_local_shape
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision2D.Bind_get_local_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(r_ret.Get())})}
	frame.Free()
	return ret
}

/*
Returns the colliding body's attached [Object].
*/
//go:nosplit
func (self class) GetCollider() [1]gd.Object { //gd:KinematicCollision2D.get_collider
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision2D.Bind_get_collider, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(r_ret.Get())})}
	frame.Free()
	return ret
}

/*
Returns the unique instance ID of the colliding body's attached [Object]. See [method Object.get_instance_id].
*/
//go:nosplit
func (self class) GetColliderId() gd.Int { //gd:KinematicCollision2D.get_collider_id
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision2D.Bind_get_collider_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the colliding body's [RID] used by the [PhysicsServer2D].
*/
//go:nosplit
func (self class) GetColliderRid() gd.RID { //gd:KinematicCollision2D.get_collider_rid
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision2D.Bind_get_collider_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the colliding body's shape.
*/
//go:nosplit
func (self class) GetColliderShape() [1]gd.Object { //gd:KinematicCollision2D.get_collider_shape
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision2D.Bind_get_collider_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(r_ret.Get())})}
	frame.Free()
	return ret
}

/*
Returns the colliding body's shape index. See [CollisionObject2D].
*/
//go:nosplit
func (self class) GetColliderShapeIndex() gd.Int { //gd:KinematicCollision2D.get_collider_shape_index
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision2D.Bind_get_collider_shape_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the colliding body's velocity.
*/
//go:nosplit
func (self class) GetColliderVelocity() gd.Vector2 { //gd:KinematicCollision2D.get_collider_velocity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.KinematicCollision2D.Bind_get_collider_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsKinematicCollision2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsKinematicCollision2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("KinematicCollision2D", func(ptr gd.Object) any {
		return [1]gdclass.KinematicCollision2D{*(*gdclass.KinematicCollision2D)(unsafe.Pointer(&ptr))}
	})
}
