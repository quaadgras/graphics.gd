package PhysicsBody2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/CollisionObject2D"
import "grow.graphics/gd/object/Node2D"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[PhysicsBody2D] is an abstract base class for 2D game objects affected by physics. All 2D physics bodies inherit from it.

*/
type Simple [1]classdb.PhysicsBody2D
func (self Simple) MoveAndCollide(motion gd.Vector2, test_only bool, safe_margin float64, recovery_as_collision bool) [1]classdb.KinematicCollision2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.KinematicCollision2D(Expert(self).MoveAndCollide(gc, motion, test_only, gd.Float(safe_margin), recovery_as_collision))
}
func (self Simple) TestMove(from gd.Transform2D, motion gd.Vector2, collision [1]classdb.KinematicCollision2D, safe_margin float64, recovery_as_collision bool) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).TestMove(from, motion, collision, gd.Float(safe_margin), recovery_as_collision))
}
func (self Simple) GetGravity() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetGravity())
}
func (self Simple) GetCollisionExceptions() gd.ArrayOf[[1]classdb.PhysicsBody2D] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.PhysicsBody2D](Expert(self).GetCollisionExceptions(gc))
}
func (self Simple) AddCollisionExceptionWith(body [1]classdb.Node) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddCollisionExceptionWith(body)
}
func (self Simple) RemoveCollisionExceptionWith(body [1]classdb.Node) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveCollisionExceptionWith(body)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PhysicsBody2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Moves the body along the vector [param motion]. In order to be frame rate independent in [method Node._physics_process] or [method Node._process], [param motion] should be computed using [code]delta[/code].
Returns a [KinematicCollision2D], which contains information about the collision when stopped, or when touching another body along the motion.
If [param test_only] is [code]true[/code], the body does not move but the would-be collision information is given.
[param safe_margin] is the extra margin used for collision recovery (see [member CharacterBody2D.safe_margin] for more details).
If [param recovery_as_collision] is [code]true[/code], any depenetration from the recovery phase is also reported as a collision; this is used e.g. by [CharacterBody2D] for improving floor detection during floor snapping.
*/
//go:nosplit
func (self class) MoveAndCollide(ctx gd.Lifetime, motion gd.Vector2, test_only bool, safe_margin gd.Float, recovery_as_collision bool) object.KinematicCollision2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, motion)
	callframe.Arg(frame, test_only)
	callframe.Arg(frame, safe_margin)
	callframe.Arg(frame, recovery_as_collision)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsBody2D.Bind_move_and_collide, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.KinematicCollision2D
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
func (self class) TestMove(from gd.Transform2D, motion gd.Vector2, collision object.KinematicCollision2D, safe_margin gd.Float, recovery_as_collision bool) bool {
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
func (self class) GetCollisionExceptions(ctx gd.Lifetime) gd.ArrayOf[object.PhysicsBody2D] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsBody2D.Bind_get_collision_exceptions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.PhysicsBody2D](ret)
}
/*
Adds a body to the list of bodies that this body can't collide with.
*/
//go:nosplit
func (self class) AddCollisionExceptionWith(body object.Node)  {
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
func (self class) RemoveCollisionExceptionWith(body object.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(body[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsBody2D.Bind_remove_collision_exception_with, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsPhysicsBody2D() Expert { return self[0].AsPhysicsBody2D() }


//go:nosplit
func (self Simple) AsPhysicsBody2D() Simple { return self[0].AsPhysicsBody2D() }


//go:nosplit
func (self class) AsCollisionObject2D() CollisionObject2D.Expert { return self[0].AsCollisionObject2D() }


//go:nosplit
func (self Simple) AsCollisionObject2D() CollisionObject2D.Simple { return self[0].AsCollisionObject2D() }


//go:nosplit
func (self class) AsNode2D() Node2D.Expert { return self[0].AsNode2D() }


//go:nosplit
func (self Simple) AsNode2D() Node2D.Simple { return self[0].AsNode2D() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("PhysicsBody2D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
