// Package PhysicsBody3D provides methods for working with PhysicsBody3D object instances.
package PhysicsBody3D

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
import "graphics.gd/classdb/CollisionObject3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Transform3D"

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

/*
[PhysicsBody3D] is an abstract base class for 3D game objects affected by physics. All 3D physics bodies inherit from it.
[b]Warning:[/b] With a non-uniform scale, this node will likely not behave as expected. It is advised to keep its scale the same on all axes and adjust its collision shape(s) instead.
*/
type Instance [1]gdclass.PhysicsBody3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPhysicsBody3D() Instance
}

/*
Moves the body along the vector [param motion]. In order to be frame rate independent in [method Node._physics_process] or [method Node._process], [param motion] should be computed using [code]delta[/code].
The body will stop if it collides. Returns a [KinematicCollision3D], which contains information about the collision when stopped, or when touching another body along the motion.
If [param test_only] is [code]true[/code], the body does not move but the would-be collision information is given.
[param safe_margin] is the extra margin used for collision recovery (see [member CharacterBody3D.safe_margin] for more details).
If [param recovery_as_collision] is [code]true[/code], any depenetration from the recovery phase is also reported as a collision; this is used e.g. by [CharacterBody3D] for improving floor detection during floor snapping.
[param max_collisions] allows to retrieve more than one collision result.
*/
func (self Instance) MoveAndCollide(motion Vector3.XYZ) [1]gdclass.KinematicCollision3D { //gd:PhysicsBody3D.move_and_collide
	return [1]gdclass.KinematicCollision3D(class(self).MoveAndCollide(gd.Vector3(motion), false, gd.Float(0.001), false, gd.Int(1)))
}

/*
Checks for collisions without moving the body. In order to be frame rate independent in [method Node._physics_process] or [method Node._process], [param motion] should be computed using [code]delta[/code].
Virtually sets the node's position, scale and rotation to that of the given [Transform3D], then tries to move the body along the vector [param motion]. Returns [code]true[/code] if a collision would stop the body from moving along the whole path.
[param collision] is an optional object of type [KinematicCollision3D], which contains additional information about the collision when stopped, or when touching another body along the motion.
[param safe_margin] is the extra margin used for collision recovery (see [member CharacterBody3D.safe_margin] for more details).
If [param recovery_as_collision] is [code]true[/code], any depenetration from the recovery phase is also reported as a collision; this is useful for checking whether the body would [i]touch[/i] any other bodies.
[param max_collisions] allows to retrieve more than one collision result.
*/
func (self Instance) TestMove(from Transform3D.BasisOrigin, motion Vector3.XYZ) bool { //gd:PhysicsBody3D.test_move
	return bool(class(self).TestMove(gd.Transform3D(from), gd.Vector3(motion), [1][1]gdclass.KinematicCollision3D{}[0], gd.Float(0.001), false, gd.Int(1)))
}

/*
Returns the gravity vector computed from all sources that can affect the body, including all gravity overrides from [Area3D] nodes and the global world gravity.
*/
func (self Instance) GetGravity() Vector3.XYZ { //gd:PhysicsBody3D.get_gravity
	return Vector3.XYZ(class(self).GetGravity())
}

/*
Returns an array of nodes that were added as collision exceptions for this body.
*/
func (self Instance) GetCollisionExceptions() [][1]gdclass.PhysicsBody3D { //gd:PhysicsBody3D.get_collision_exceptions
	return [][1]gdclass.PhysicsBody3D(gd.ArrayAs[[][1]gdclass.PhysicsBody3D](gd.InternalArray(class(self).GetCollisionExceptions())))
}

/*
Adds a body to the list of bodies that this body can't collide with.
*/
func (self Instance) AddCollisionExceptionWith(body [1]gdclass.Node) { //gd:PhysicsBody3D.add_collision_exception_with
	class(self).AddCollisionExceptionWith(body)
}

/*
Removes a body from the list of bodies that this body can't collide with.
*/
func (self Instance) RemoveCollisionExceptionWith(body [1]gdclass.Node) { //gd:PhysicsBody3D.remove_collision_exception_with
	class(self).RemoveCollisionExceptionWith(body)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicsBody3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsBody3D"))
	casted := Instance{*(*gdclass.PhysicsBody3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) AxisLockLinearX() bool {
	return bool(class(self).GetAxisLock(1))
}

func (self Instance) SetAxisLockLinearX(value bool) {
	class(self).SetAxisLock(1, value)
}

func (self Instance) AxisLockLinearY() bool {
	return bool(class(self).GetAxisLock(2))
}

func (self Instance) SetAxisLockLinearY(value bool) {
	class(self).SetAxisLock(2, value)
}

func (self Instance) AxisLockLinearZ() bool {
	return bool(class(self).GetAxisLock(4))
}

func (self Instance) SetAxisLockLinearZ(value bool) {
	class(self).SetAxisLock(4, value)
}

func (self Instance) AxisLockAngularX() bool {
	return bool(class(self).GetAxisLock(8))
}

func (self Instance) SetAxisLockAngularX(value bool) {
	class(self).SetAxisLock(8, value)
}

func (self Instance) AxisLockAngularY() bool {
	return bool(class(self).GetAxisLock(16))
}

func (self Instance) SetAxisLockAngularY(value bool) {
	class(self).SetAxisLock(16, value)
}

func (self Instance) AxisLockAngularZ() bool {
	return bool(class(self).GetAxisLock(32))
}

func (self Instance) SetAxisLockAngularZ(value bool) {
	class(self).SetAxisLock(32, value)
}

/*
Moves the body along the vector [param motion]. In order to be frame rate independent in [method Node._physics_process] or [method Node._process], [param motion] should be computed using [code]delta[/code].
The body will stop if it collides. Returns a [KinematicCollision3D], which contains information about the collision when stopped, or when touching another body along the motion.
If [param test_only] is [code]true[/code], the body does not move but the would-be collision information is given.
[param safe_margin] is the extra margin used for collision recovery (see [member CharacterBody3D.safe_margin] for more details).
If [param recovery_as_collision] is [code]true[/code], any depenetration from the recovery phase is also reported as a collision; this is used e.g. by [CharacterBody3D] for improving floor detection during floor snapping.
[param max_collisions] allows to retrieve more than one collision result.
*/
//go:nosplit
func (self class) MoveAndCollide(motion gd.Vector3, test_only bool, safe_margin gd.Float, recovery_as_collision bool, max_collisions gd.Int) [1]gdclass.KinematicCollision3D { //gd:PhysicsBody3D.move_and_collide
	var frame = callframe.New()
	callframe.Arg(frame, motion)
	callframe.Arg(frame, test_only)
	callframe.Arg(frame, safe_margin)
	callframe.Arg(frame, recovery_as_collision)
	callframe.Arg(frame, max_collisions)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsBody3D.Bind_move_and_collide, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.KinematicCollision3D{gd.PointerWithOwnershipTransferredToGo[gdclass.KinematicCollision3D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Checks for collisions without moving the body. In order to be frame rate independent in [method Node._physics_process] or [method Node._process], [param motion] should be computed using [code]delta[/code].
Virtually sets the node's position, scale and rotation to that of the given [Transform3D], then tries to move the body along the vector [param motion]. Returns [code]true[/code] if a collision would stop the body from moving along the whole path.
[param collision] is an optional object of type [KinematicCollision3D], which contains additional information about the collision when stopped, or when touching another body along the motion.
[param safe_margin] is the extra margin used for collision recovery (see [member CharacterBody3D.safe_margin] for more details).
If [param recovery_as_collision] is [code]true[/code], any depenetration from the recovery phase is also reported as a collision; this is useful for checking whether the body would [i]touch[/i] any other bodies.
[param max_collisions] allows to retrieve more than one collision result.
*/
//go:nosplit
func (self class) TestMove(from gd.Transform3D, motion gd.Vector3, collision [1]gdclass.KinematicCollision3D, safe_margin gd.Float, recovery_as_collision bool, max_collisions gd.Int) bool { //gd:PhysicsBody3D.test_move
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, motion)
	callframe.Arg(frame, pointers.Get(collision[0])[0])
	callframe.Arg(frame, safe_margin)
	callframe.Arg(frame, recovery_as_collision)
	callframe.Arg(frame, max_collisions)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsBody3D.Bind_test_move, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the gravity vector computed from all sources that can affect the body, including all gravity overrides from [Area3D] nodes and the global world gravity.
*/
//go:nosplit
func (self class) GetGravity() gd.Vector3 { //gd:PhysicsBody3D.get_gravity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsBody3D.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Locks or unlocks the specified linear or rotational [param axis] depending on the value of [param lock].
*/
//go:nosplit
func (self class) SetAxisLock(axis gdclass.PhysicsServer3DBodyAxis, lock bool) { //gd:PhysicsBody3D.set_axis_lock
	var frame = callframe.New()
	callframe.Arg(frame, axis)
	callframe.Arg(frame, lock)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsBody3D.Bind_set_axis_lock, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the specified linear or rotational [param axis] is locked.
*/
//go:nosplit
func (self class) GetAxisLock(axis gdclass.PhysicsServer3DBodyAxis) bool { //gd:PhysicsBody3D.get_axis_lock
	var frame = callframe.New()
	callframe.Arg(frame, axis)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsBody3D.Bind_get_axis_lock, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array of nodes that were added as collision exceptions for this body.
*/
//go:nosplit
func (self class) GetCollisionExceptions() Array.Contains[[1]gdclass.PhysicsBody3D] { //gd:PhysicsBody3D.get_collision_exceptions
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsBody3D.Bind_get_collision_exceptions, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.PhysicsBody3D]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Adds a body to the list of bodies that this body can't collide with.
*/
//go:nosplit
func (self class) AddCollisionExceptionWith(body [1]gdclass.Node) { //gd:PhysicsBody3D.add_collision_exception_with
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(body[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsBody3D.Bind_add_collision_exception_with, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a body from the list of bodies that this body can't collide with.
*/
//go:nosplit
func (self class) RemoveCollisionExceptionWith(body [1]gdclass.Node) { //gd:PhysicsBody3D.remove_collision_exception_with
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(body[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsBody3D.Bind_remove_collision_exception_with, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsPhysicsBody3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPhysicsBody3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsCollisionObject3D() CollisionObject3D.Advanced {
	return *((*CollisionObject3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCollisionObject3D() CollisionObject3D.Instance {
	return *((*CollisionObject3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(CollisionObject3D.Advanced(self.AsCollisionObject3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(CollisionObject3D.Instance(self.AsCollisionObject3D()), name)
	}
}
func init() {
	gdclass.Register("PhysicsBody3D", func(ptr gd.Object) any {
		return [1]gdclass.PhysicsBody3D{*(*gdclass.PhysicsBody3D)(unsafe.Pointer(&ptr))}
	})
}
