package StaticBody2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/PhysicsBody2D"
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
A static 2D physics body. It can't be moved by external forces or contacts, but can be moved manually by other means such as code, [AnimationMixer]s (with [member AnimationMixer.callback_mode_process] set to [constant AnimationMixer.ANIMATION_CALLBACK_MODE_PROCESS_PHYSICS]), and [RemoteTransform2D].
When [StaticBody2D] is moved, it is teleported to its new position without affecting other physics bodies in its path. If this is not desired, use [AnimatableBody2D] instead.
[StaticBody2D] is useful for completely static objects like floors and walls, as well as moving surfaces like conveyor belts and circular revolving platforms (by using [member constant_linear_velocity] and [member constant_angular_velocity]).

*/
type Go [1]classdb.StaticBody2D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.StaticBody2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("StaticBody2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) PhysicsMaterialOverride() gdclass.PhysicsMaterial {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.PhysicsMaterial(class(self).GetPhysicsMaterialOverride(gc))
}

func (self Go) SetPhysicsMaterialOverride(value gdclass.PhysicsMaterial) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPhysicsMaterialOverride(value)
}

func (self Go) ConstantLinearVelocity() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetConstantLinearVelocity())
}

func (self Go) SetConstantLinearVelocity(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetConstantLinearVelocity(value)
}

func (self Go) ConstantAngularVelocity() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetConstantAngularVelocity()))
}

func (self Go) SetConstantAngularVelocity(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetConstantAngularVelocity(gd.Float(value))
}

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
func (self class) SetPhysicsMaterialOverride(physics_material_override gdclass.PhysicsMaterial)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(physics_material_override[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StaticBody2D.Bind_set_physics_material_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPhysicsMaterialOverride(ctx gd.Lifetime) gdclass.PhysicsMaterial {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StaticBody2D.Bind_get_physics_material_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.PhysicsMaterial
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsStaticBody2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsStaticBody2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsPhysicsBody2D() PhysicsBody2D.GD { return *((*PhysicsBody2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsPhysicsBody2D() PhysicsBody2D.Go { return *((*PhysicsBody2D.Go)(unsafe.Pointer(&self))) }
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
	default: return gd.VirtualByName(self.AsPhysicsBody2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsPhysicsBody2D(), name)
	}
}
func init() {classdb.Register("StaticBody2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
