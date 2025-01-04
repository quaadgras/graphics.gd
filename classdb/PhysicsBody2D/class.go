// Package PhysicsBody2D provides methods for working with PhysicsBody2D object instances.
package PhysicsBody2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/CollisionObject2D"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Transform2D"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[PhysicsBody2D] is an abstract base class for 2D game objects affected by physics. All 2D physics bodies inherit from it.
*/
type Instance [1]gdclass.PhysicsBody2D
type Any interface {
	gd.IsClass
	AsPhysicsBody2D() Instance
}

/*
Moves the body along the vector [param motion]. In order to be frame rate independent in [method Node._physics_process] or [method Node._process], [param motion] should be computed using [code]delta[/code].
Returns a [KinematicCollision2D], which contains information about the collision when stopped, or when touching another body along the motion.
If [param test_only] is [code]true[/code], the body does not move but the would-be collision information is given.
[param safe_margin] is the extra margin used for collision recovery (see [member CharacterBody2D.safe_margin] for more details).
If [param recovery_as_collision] is [code]true[/code], any depenetration from the recovery phase is also reported as a collision; this is used e.g. by [CharacterBody2D] for improving floor detection during floor snapping.
*/
func (self Instance) MoveAndCollide(motion Vector2.XY) [1]gdclass.KinematicCollision2D {
	return [1]gdclass.KinematicCollision2D(class(self).MoveAndCollide(gd.Vector2(motion), false, gd.Float(0.08), false))
}

/*
Checks for collisions without moving the body. In order to be frame rate independent in [method Node._physics_process] or [method Node._process], [param motion] should be computed using [code]delta[/code].
Virtually sets the node's position, scale and rotation to that of the given [Transform2D], then tries to move the body along the vector [param motion]. Returns [code]true[/code] if a collision would stop the body from moving along the whole path.
[param collision] is an optional object of type [KinematicCollision2D], which contains additional information about the collision when stopped, or when touching another body along the motion.
[param safe_margin] is the extra margin used for collision recovery (see [member CharacterBody2D.safe_margin] for more details).
If [param recovery_as_collision] is [code]true[/code], any depenetration from the recovery phase is also reported as a collision; this is useful for checking whether the body would [i]touch[/i] any other bodies.
*/
func (self Instance) TestMove(from Transform2D.OriginXY, motion Vector2.XY) bool {
	return bool(class(self).TestMove(gd.Transform2D(from), gd.Vector2(motion), [1][1]gdclass.KinematicCollision2D{}[0], gd.Float(0.08), false))
}

/*
Returns the gravity vector computed from all sources that can affect the body, including all gravity overrides from [Area2D] nodes and the global world gravity.
*/
func (self Instance) GetGravity() Vector2.XY {
	return Vector2.XY(class(self).GetGravity())
}

/*
Returns an array of nodes that were added as collision exceptions for this body.
*/
func (self Instance) GetCollisionExceptions() gd.Array {
	return gd.Array(class(self).GetCollisionExceptions())
}

/*
Adds a body to the list of bodies that this body can't collide with.
*/
func (self Instance) AddCollisionExceptionWith(body [1]gdclass.Node) {
	class(self).AddCollisionExceptionWith(body)
}

/*
Removes a body from the list of bodies that this body can't collide with.
*/
func (self Instance) RemoveCollisionExceptionWith(body [1]gdclass.Node) {
	class(self).RemoveCollisionExceptionWith(body)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicsBody2D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsBody2D"))
	return Instance{*(*gdclass.PhysicsBody2D)(unsafe.Pointer(&object))}
}

/*
Moves the body along the vector [param motion]. In order to be frame rate independent in [method Node._physics_process] or [method Node._process], [param motion] should be computed using [code]delta[/code].
Returns a [KinematicCollision2D], which contains information about the collision when stopped, or when touching another body along the motion.
If [param test_only] is [code]true[/code], the body does not move but the would-be collision information is given.
[param safe_margin] is the extra margin used for collision recovery (see [member CharacterBody2D.safe_margin] for more details).
If [param recovery_as_collision] is [code]true[/code], any depenetration from the recovery phase is also reported as a collision; this is used e.g. by [CharacterBody2D] for improving floor detection during floor snapping.
*/
//go:nosplit
func (self class) MoveAndCollide(motion gd.Vector2, test_only bool, safe_margin gd.Float, recovery_as_collision bool) [1]gdclass.KinematicCollision2D {
	var frame = callframe.New()
	callframe.Arg(frame, motion)
	callframe.Arg(frame, test_only)
	callframe.Arg(frame, safe_margin)
	callframe.Arg(frame, recovery_as_collision)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsBody2D.Bind_move_and_collide, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.KinematicCollision2D{gd.PointerWithOwnershipTransferredToGo[gdclass.KinematicCollision2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Checks for collisions without moving the body. In order to be frame rate independent in [method Node._physics_process] or [method Node._process], [param motion] should be computed using [code]delta[/code].
Virtually sets the node's position, scale and rotation to that of the given [Transform2D], then tries to move the body along the vector [param motion]. Returns [code]true[/code] if a collision would stop the body from moving along the whole path.
[param collision] is an optional object of type [KinematicCollision2D], which contains additional information about the collision when stopped, or when touching another body along the motion.
[param safe_margin] is the extra margin used for collision recovery (see [member CharacterBody2D.safe_margin] for more details).
If [param recovery_as_collision] is [code]true[/code], any depenetration from the recovery phase is also reported as a collision; this is useful for checking whether the body would [i]touch[/i] any other bodies.
*/
//go:nosplit
func (self class) TestMove(from gd.Transform2D, motion gd.Vector2, collision [1]gdclass.KinematicCollision2D, safe_margin gd.Float, recovery_as_collision bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, motion)
	callframe.Arg(frame, pointers.Get(collision[0])[0])
	callframe.Arg(frame, safe_margin)
	callframe.Arg(frame, recovery_as_collision)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsBody2D.Bind_test_move, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the gravity vector computed from all sources that can affect the body, including all gravity overrides from [Area2D] nodes and the global world gravity.
*/
//go:nosplit
func (self class) GetGravity() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsBody2D.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array of nodes that were added as collision exceptions for this body.
*/
//go:nosplit
func (self class) GetCollisionExceptions() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsBody2D.Bind_get_collision_exceptions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Adds a body to the list of bodies that this body can't collide with.
*/
//go:nosplit
func (self class) AddCollisionExceptionWith(body [1]gdclass.Node) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(body[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsBody2D.Bind_add_collision_exception_with, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes a body from the list of bodies that this body can't collide with.
*/
//go:nosplit
func (self class) RemoveCollisionExceptionWith(body [1]gdclass.Node) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(body[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsBody2D.Bind_remove_collision_exception_with, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsPhysicsBody2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPhysicsBody2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsCollisionObject2D() CollisionObject2D.Advanced {
	return *((*CollisionObject2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCollisionObject2D() CollisionObject2D.Instance {
	return *((*CollisionObject2D.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(self.AsCollisionObject2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsCollisionObject2D(), name)
	}
}
func init() {
	gdclass.Register("PhysicsBody2D", func(ptr gd.Object) any {
		return [1]gdclass.PhysicsBody2D{*(*gdclass.PhysicsBody2D)(unsafe.Pointer(&ptr))}
	})
}
