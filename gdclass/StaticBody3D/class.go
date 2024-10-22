package StaticBody3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/PhysicsBody3D"
import "grow.graphics/gd/gdclass/CollisionObject3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A static 3D physics body. It can't be moved by external forces or contacts, but can be moved manually by other means such as code, [AnimationMixer]s (with [member AnimationMixer.callback_mode_process] set to [constant AnimationMixer.ANIMATION_CALLBACK_MODE_PROCESS_PHYSICS]), and [RemoteTransform3D].
When [StaticBody3D] is moved, it is teleported to its new position without affecting other physics bodies in its path. If this is not desired, use [AnimatableBody3D] instead.
[StaticBody3D] is useful for completely static objects like floors and walls, as well as moving surfaces like conveyor belts and circular revolving platforms (by using [member constant_linear_velocity] and [member constant_angular_velocity]).

*/
type Go [1]classdb.StaticBody3D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.StaticBody3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("StaticBody3D"))
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

func (self Go) ConstantLinearVelocity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetConstantLinearVelocity())
}

func (self Go) SetConstantLinearVelocity(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetConstantLinearVelocity(value)
}

func (self Go) ConstantAngularVelocity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetConstantAngularVelocity())
}

func (self Go) SetConstantAngularVelocity(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetConstantAngularVelocity(value)
}

//go:nosplit
func (self class) SetConstantLinearVelocity(vel gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, vel)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StaticBody3D.Bind_set_constant_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetConstantAngularVelocity(vel gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, vel)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StaticBody3D.Bind_set_constant_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetConstantLinearVelocity() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StaticBody3D.Bind_get_constant_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetConstantAngularVelocity() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StaticBody3D.Bind_get_constant_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StaticBody3D.Bind_set_physics_material_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPhysicsMaterialOverride(ctx gd.Lifetime) gdclass.PhysicsMaterial {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StaticBody3D.Bind_get_physics_material_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.PhysicsMaterial
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsStaticBody3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsStaticBody3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsPhysicsBody3D() PhysicsBody3D.GD { return *((*PhysicsBody3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsPhysicsBody3D() PhysicsBody3D.Go { return *((*PhysicsBody3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCollisionObject3D() CollisionObject3D.GD { return *((*CollisionObject3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCollisionObject3D() CollisionObject3D.Go { return *((*CollisionObject3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsPhysicsBody3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsPhysicsBody3D(), name)
	}
}
func init() {classdb.Register("StaticBody3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
