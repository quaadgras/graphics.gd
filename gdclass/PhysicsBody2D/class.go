package PhysicsBody2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/CollisionObject2D"
import "grow.graphics/gd/gdclass/Node2D"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[PhysicsBody2D] is an abstract base class for 2D game objects affected by physics. All 2D physics bodies inherit from it.

*/
type Go [1]classdb.PhysicsBody2D

/*
Moves the body along the vector [param motion]. In order to be frame rate independent in [method Node._physics_process] or [method Node._process], [param motion] should be computed using [code]delta[/code].
Returns a [KinematicCollision2D], which contains information about the collision when stopped, or when touching another body along the motion.
If [param test_only] is [code]true[/code], the body does not move but the would-be collision information is given.
[param safe_margin] is the extra margin used for collision recovery (see [member CharacterBody2D.safe_margin] for more details).
If [param recovery_as_collision] is [code]true[/code], any depenetration from the recovery phase is also reported as a collision; this is used e.g. by [CharacterBody2D] for improving floor detection during floor snapping.
*/
func (self Go) MoveAndCollide(motion gd.Vector2) gdclass.KinematicCollision2D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.KinematicCollision2D(class(self).MoveAndCollide(gc, motion, false, gd.Float(0.08), false))
}

/*
Checks for collisions without moving the body. In order to be frame rate independent in [method Node._physics_process] or [method Node._process], [param motion] should be computed using [code]delta[/code].
Virtually sets the node's position, scale and rotation to that of the given [Transform2D], then tries to move the body along the vector [param motion]. Returns [code]true[/code] if a collision would stop the body from moving along the whole path.
[param collision] is an optional object of type [KinematicCollision2D], which contains additional information about the collision when stopped, or when touching another body along the motion.
[param safe_margin] is the extra margin used for collision recovery (see [member CharacterBody2D.safe_margin] for more details).
If [param recovery_as_collision] is [code]true[/code], any depenetration from the recovery phase is also reported as a collision; this is useful for checking whether the body would [i]touch[/i] any other bodies.
*/
func (self Go) TestMove(from gd.Transform2D, motion gd.Vector2) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).TestMove(from, motion, ([1]gdclass.KinematicCollision2D{}[0]), gd.Float(0.08), false))
}

/*
Returns the gravity vector computed from all sources that can affect the body, including all gravity overrides from [Area2D] nodes and the global world gravity.
*/
func (self Go) GetGravity() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetGravity())
}

/*
Returns an array of nodes that were added as collision exceptions for this body.
*/
func (self Go) GetCollisionExceptions() gd.ArrayOf[gdclass.PhysicsBody2D] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gdclass.PhysicsBody2D](class(self).GetCollisionExceptions(gc))
}

/*
Adds a body to the list of bodies that this body can't collide with.
*/
func (self Go) AddCollisionExceptionWith(body gdclass.Node) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddCollisionExceptionWith(body)
}

/*
Removes a body from the list of bodies that this body can't collide with.
*/
func (self Go) RemoveCollisionExceptionWith(body gdclass.Node) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveCollisionExceptionWith(body)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PhysicsBody2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PhysicsBody2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Moves the body along the vector [param motion]. In order to be frame rate independent in [method Node._physics_process] or [method Node._process], [param motion] should be computed using [code]delta[/code].
Returns a [KinematicCollision2D], which contains information about the collision when stopped, or when touching another body along the motion.
If [param test_only] is [code]true[/code], the body does not move but the would-be collision information is given.
[param safe_margin] is the extra margin used for collision recovery (see [member CharacterBody2D.safe_margin] for more details).
If [param recovery_as_collision] is [code]true[/code], any depenetration from the recovery phase is also reported as a collision; this is used e.g. by [CharacterBody2D] for improving floor detection during floor snapping.
*/
//go:nosplit
func (self class) MoveAndCollide(ctx gd.Lifetime, motion gd.Vector2, test_only bool, safe_margin gd.Float, recovery_as_collision bool) gdclass.KinematicCollision2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, motion)
	callframe.Arg(frame, test_only)
	callframe.Arg(frame, safe_margin)
	callframe.Arg(frame, recovery_as_collision)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsBody2D.Bind_move_and_collide, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.KinematicCollision2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
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
func (self class) TestMove(from gd.Transform2D, motion gd.Vector2, collision gdclass.KinematicCollision2D, safe_margin gd.Float, recovery_as_collision bool) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, motion)
	callframe.Arg(frame, mmm.Get(collision[0].AsPointer())[0])
	callframe.Arg(frame, safe_margin)
	callframe.Arg(frame, recovery_as_collision)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsBody2D.Bind_test_move, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the gravity vector computed from all sources that can affect the body, including all gravity overrides from [Area2D] nodes and the global world gravity.
*/
//go:nosplit
func (self class) GetGravity() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsBody2D.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array of nodes that were added as collision exceptions for this body.
*/
//go:nosplit
func (self class) GetCollisionExceptions(ctx gd.Lifetime) gd.ArrayOf[gdclass.PhysicsBody2D] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsBody2D.Bind_get_collision_exceptions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.PhysicsBody2D](ret)
}
/*
Adds a body to the list of bodies that this body can't collide with.
*/
//go:nosplit
func (self class) AddCollisionExceptionWith(body gdclass.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(body[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsBody2D.Bind_add_collision_exception_with, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a body from the list of bodies that this body can't collide with.
*/
//go:nosplit
func (self class) RemoveCollisionExceptionWith(body gdclass.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(body[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsBody2D.Bind_remove_collision_exception_with, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsPhysicsBody2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPhysicsBody2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsCollisionObject2D() CollisionObject2D.GD { return *((*CollisionObject2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCollisionObject2D() CollisionObject2D.Go { return *((*CollisionObject2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.GD { return *((*Node2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode2D() Node2D.Go { return *((*Node2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsCollisionObject2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsCollisionObject2D(), name)
	}
}
func init() {classdb.Register("PhysicsBody2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
