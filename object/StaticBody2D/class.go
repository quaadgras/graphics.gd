package StaticBody2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/PhysicsBody2D"
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
A static 2D physics body. It can't be moved by external forces or contacts, but can be moved manually by other means such as code, [AnimationMixer]s (with [member AnimationMixer.callback_mode_process] set to [constant AnimationMixer.ANIMATION_CALLBACK_MODE_PROCESS_PHYSICS]), and [RemoteTransform2D].
When [StaticBody2D] is moved, it is teleported to its new position without affecting other physics bodies in its path. If this is not desired, use [AnimatableBody2D] instead.
[StaticBody2D] is useful for completely static objects like floors and walls, as well as moving surfaces like conveyor belts and circular revolving platforms (by using [member constant_linear_velocity] and [member constant_angular_velocity]).

*/
type Simple [1]classdb.StaticBody2D
func (self Simple) SetConstantLinearVelocity(vel gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetConstantLinearVelocity(vel)
}
func (self Simple) SetConstantAngularVelocity(vel float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetConstantAngularVelocity(gd.Float(vel))
}
func (self Simple) GetConstantLinearVelocity() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetConstantLinearVelocity())
}
func (self Simple) GetConstantAngularVelocity() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetConstantAngularVelocity()))
}
func (self Simple) SetPhysicsMaterialOverride(physics_material_override [1]classdb.PhysicsMaterial) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPhysicsMaterialOverride(physics_material_override)
}
func (self Simple) GetPhysicsMaterialOverride() [1]classdb.PhysicsMaterial {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PhysicsMaterial(Expert(self).GetPhysicsMaterialOverride(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.StaticBody2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetConstantLinearVelocity(vel gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, vel)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StaticBody2D.Bind_set_constant_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetConstantAngularVelocity(vel gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, vel)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StaticBody2D.Bind_set_constant_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetConstantLinearVelocity() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StaticBody2D.Bind_get_constant_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetConstantAngularVelocity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StaticBody2D.Bind_get_constant_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPhysicsMaterialOverride(physics_material_override object.PhysicsMaterial)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(physics_material_override[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StaticBody2D.Bind_set_physics_material_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPhysicsMaterialOverride(ctx gd.Lifetime) object.PhysicsMaterial {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StaticBody2D.Bind_get_physics_material_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.PhysicsMaterial
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsStaticBody2D() Expert { return self[0].AsStaticBody2D() }


//go:nosplit
func (self Simple) AsStaticBody2D() Simple { return self[0].AsStaticBody2D() }


//go:nosplit
func (self class) AsPhysicsBody2D() PhysicsBody2D.Expert { return self[0].AsPhysicsBody2D() }


//go:nosplit
func (self Simple) AsPhysicsBody2D() PhysicsBody2D.Simple { return self[0].AsPhysicsBody2D() }


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
func init() {classdb.Register("StaticBody2D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
