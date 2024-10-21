package RigidBody2D

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
[RigidBody2D] implements full 2D physics. It cannot be controlled directly, instead, you must apply forces to it (gravity, impulses, etc.), and the physics simulation will calculate the resulting movement, rotation, react to collisions, and affect other physics bodies in its path.
The body's behavior can be adjusted via [member lock_rotation], [member freeze], and [member freeze_mode]. By changing various properties of the object, such as [member mass], you can control how the physics simulation acts on it.
A rigid body will always maintain its shape and size, even when forces are applied to it. It is useful for objects that can be interacted with in an environment, such as a tree that can be knocked over or a stack of crates that can be pushed around.
If you need to override the default physics behavior, you can write a custom force integration function. See [member custom_integrator].
[b]Note:[/b] Changing the 2D transform or [member linear_velocity] of a [RigidBody2D] very often may lead to some unpredictable behaviors. If you need to directly affect the body, prefer [method _integrate_forces] as it allows you to directly access the physics state.
	// RigidBody2D methods that can be overridden by a [Class] that extends it.
	type RigidBody2D interface {
		//Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
		IntegrateForces(state object.PhysicsDirectBodyState2D) 
	}

*/
type Simple [1]classdb.RigidBody2D
func (Simple) _integrate_forces(impl func(ptr unsafe.Pointer, state [1]classdb.PhysicsDirectBodyState2D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var state [1]classdb.PhysicsDirectBodyState2D
		state[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, state)
		gc.End()
	}
}
func (self Simple) SetMass(mass float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMass(gd.Float(mass))
}
func (self Simple) GetMass() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMass()))
}
func (self Simple) GetInertia() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetInertia()))
}
func (self Simple) SetInertia(inertia float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInertia(gd.Float(inertia))
}
func (self Simple) SetCenterOfMassMode(mode classdb.RigidBody2DCenterOfMassMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCenterOfMassMode(mode)
}
func (self Simple) GetCenterOfMassMode() classdb.RigidBody2DCenterOfMassMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RigidBody2DCenterOfMassMode(Expert(self).GetCenterOfMassMode())
}
func (self Simple) SetCenterOfMass(center_of_mass gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCenterOfMass(center_of_mass)
}
func (self Simple) GetCenterOfMass() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetCenterOfMass())
}
func (self Simple) SetPhysicsMaterialOverride(physics_material_override [1]classdb.PhysicsMaterial) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPhysicsMaterialOverride(physics_material_override)
}
func (self Simple) GetPhysicsMaterialOverride() [1]classdb.PhysicsMaterial {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PhysicsMaterial(Expert(self).GetPhysicsMaterialOverride(gc))
}
func (self Simple) SetGravityScale(gravity_scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGravityScale(gd.Float(gravity_scale))
}
func (self Simple) GetGravityScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetGravityScale()))
}
func (self Simple) SetLinearDampMode(linear_damp_mode classdb.RigidBody2DDampMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLinearDampMode(linear_damp_mode)
}
func (self Simple) GetLinearDampMode() classdb.RigidBody2DDampMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RigidBody2DDampMode(Expert(self).GetLinearDampMode())
}
func (self Simple) SetAngularDampMode(angular_damp_mode classdb.RigidBody2DDampMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAngularDampMode(angular_damp_mode)
}
func (self Simple) GetAngularDampMode() classdb.RigidBody2DDampMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RigidBody2DDampMode(Expert(self).GetAngularDampMode())
}
func (self Simple) SetLinearDamp(linear_damp float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLinearDamp(gd.Float(linear_damp))
}
func (self Simple) GetLinearDamp() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetLinearDamp()))
}
func (self Simple) SetAngularDamp(angular_damp float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAngularDamp(gd.Float(angular_damp))
}
func (self Simple) GetAngularDamp() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAngularDamp()))
}
func (self Simple) SetLinearVelocity(linear_velocity gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLinearVelocity(linear_velocity)
}
func (self Simple) GetLinearVelocity() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetLinearVelocity())
}
func (self Simple) SetAngularVelocity(angular_velocity float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAngularVelocity(gd.Float(angular_velocity))
}
func (self Simple) GetAngularVelocity() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAngularVelocity()))
}
func (self Simple) SetMaxContactsReported(amount int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaxContactsReported(gd.Int(amount))
}
func (self Simple) GetMaxContactsReported() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMaxContactsReported()))
}
func (self Simple) GetContactCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetContactCount()))
}
func (self Simple) SetUseCustomIntegrator(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseCustomIntegrator(enable)
}
func (self Simple) IsUsingCustomIntegrator() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUsingCustomIntegrator())
}
func (self Simple) SetContactMonitor(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetContactMonitor(enabled)
}
func (self Simple) IsContactMonitorEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsContactMonitorEnabled())
}
func (self Simple) SetContinuousCollisionDetectionMode(mode classdb.RigidBody2DCCDMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetContinuousCollisionDetectionMode(mode)
}
func (self Simple) GetContinuousCollisionDetectionMode() classdb.RigidBody2DCCDMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RigidBody2DCCDMode(Expert(self).GetContinuousCollisionDetectionMode())
}
func (self Simple) SetAxisVelocity(axis_velocity gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAxisVelocity(axis_velocity)
}
func (self Simple) ApplyCentralImpulse(impulse gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ApplyCentralImpulse(impulse)
}
func (self Simple) ApplyImpulse(impulse gd.Vector2, position gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ApplyImpulse(impulse, position)
}
func (self Simple) ApplyTorqueImpulse(torque float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ApplyTorqueImpulse(gd.Float(torque))
}
func (self Simple) ApplyCentralForce(force gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ApplyCentralForce(force)
}
func (self Simple) ApplyForce(force gd.Vector2, position gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ApplyForce(force, position)
}
func (self Simple) ApplyTorque(torque float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ApplyTorque(gd.Float(torque))
}
func (self Simple) AddConstantCentralForce(force gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddConstantCentralForce(force)
}
func (self Simple) AddConstantForce(force gd.Vector2, position gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddConstantForce(force, position)
}
func (self Simple) AddConstantTorque(torque float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddConstantTorque(gd.Float(torque))
}
func (self Simple) SetConstantForce(force gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetConstantForce(force)
}
func (self Simple) GetConstantForce() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetConstantForce())
}
func (self Simple) SetConstantTorque(torque float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetConstantTorque(gd.Float(torque))
}
func (self Simple) GetConstantTorque() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetConstantTorque()))
}
func (self Simple) SetSleeping(sleeping bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSleeping(sleeping)
}
func (self Simple) IsSleeping() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSleeping())
}
func (self Simple) SetCanSleep(able_to_sleep bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCanSleep(able_to_sleep)
}
func (self Simple) IsAbleToSleep() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAbleToSleep())
}
func (self Simple) SetLockRotationEnabled(lock_rotation bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLockRotationEnabled(lock_rotation)
}
func (self Simple) IsLockRotationEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsLockRotationEnabled())
}
func (self Simple) SetFreezeEnabled(freeze_mode bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFreezeEnabled(freeze_mode)
}
func (self Simple) IsFreezeEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsFreezeEnabled())
}
func (self Simple) SetFreezeMode(freeze_mode classdb.RigidBody2DFreezeMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFreezeMode(freeze_mode)
}
func (self Simple) GetFreezeMode() classdb.RigidBody2DFreezeMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RigidBody2DFreezeMode(Expert(self).GetFreezeMode())
}
func (self Simple) GetCollidingBodies() gd.ArrayOf[[1]classdb.Node2D] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.Node2D](Expert(self).GetCollidingBodies(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RigidBody2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
*/
func (class) _integrate_forces(impl func(ptr unsafe.Pointer, state object.PhysicsDirectBodyState2D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var state object.PhysicsDirectBodyState2D
		state[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, state)
		ctx.End()
	}
}

//go:nosplit
func (self class) SetMass(mass gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mass)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMass() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_get_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetInertia() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_get_inertia, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInertia(inertia gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, inertia)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_inertia, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCenterOfMassMode(mode classdb.RigidBody2DCenterOfMassMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_center_of_mass_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCenterOfMassMode() classdb.RigidBody2DCenterOfMassMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RigidBody2DCenterOfMassMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_get_center_of_mass_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCenterOfMass(center_of_mass gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, center_of_mass)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_center_of_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCenterOfMass() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_get_center_of_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_physics_material_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPhysicsMaterialOverride(ctx gd.Lifetime) object.PhysicsMaterial {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_get_physics_material_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.PhysicsMaterial
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGravityScale(gravity_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gravity_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_gravity_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGravityScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_get_gravity_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLinearDampMode(linear_damp_mode classdb.RigidBody2DDampMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, linear_damp_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_linear_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLinearDampMode() classdb.RigidBody2DDampMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RigidBody2DDampMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_get_linear_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAngularDampMode(angular_damp_mode classdb.RigidBody2DDampMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angular_damp_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_angular_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAngularDampMode() classdb.RigidBody2DDampMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RigidBody2DDampMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_get_angular_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLinearDamp(linear_damp gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, linear_damp)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_linear_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLinearDamp() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_get_linear_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAngularDamp(angular_damp gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angular_damp)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_angular_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAngularDamp() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_get_angular_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLinearVelocity(linear_velocity gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, linear_velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLinearVelocity() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_get_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAngularVelocity(angular_velocity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angular_velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAngularVelocity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_get_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxContactsReported(amount gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_max_contacts_reported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxContactsReported() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_get_max_contacts_reported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of contacts this body has with other bodies. By default, this returns 0 unless bodies are configured to monitor contacts (see [member contact_monitor]).
[b]Note:[/b] To retrieve the colliding bodies, use [method get_colliding_bodies].
*/
//go:nosplit
func (self class) GetContactCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_get_contact_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseCustomIntegrator(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_use_custom_integrator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingCustomIntegrator() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_is_using_custom_integrator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetContactMonitor(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_contact_monitor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsContactMonitorEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_is_contact_monitor_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetContinuousCollisionDetectionMode(mode classdb.RigidBody2DCCDMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_continuous_collision_detection_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetContinuousCollisionDetectionMode() classdb.RigidBody2DCCDMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RigidBody2DCCDMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_get_continuous_collision_detection_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the body's velocity on the given axis. The velocity in the given vector axis will be set as the given vector length. This is useful for jumping behavior.
*/
//go:nosplit
func (self class) SetAxisVelocity(axis_velocity gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, axis_velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_axis_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a directional impulse without affecting rotation.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
This is equivalent to using [method apply_impulse] at the body's center of mass.
*/
//go:nosplit
func (self class) ApplyCentralImpulse(impulse gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, impulse)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_apply_central_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a positioned impulse to the body.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) ApplyImpulse(impulse gd.Vector2, position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, impulse)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_apply_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a rotational impulse to the body without affecting the position.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[b]Note:[/b] [member inertia] is required for this to work. To have [member inertia], an active [CollisionShape2D] must be a child of the node, or you can manually set [member inertia].
*/
//go:nosplit
func (self class) ApplyTorqueImpulse(torque gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_apply_torque_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a directional force without affecting rotation. A force is time dependent and meant to be applied every physics update.
This is equivalent to using [method apply_force] at the body's center of mass.
*/
//go:nosplit
func (self class) ApplyCentralForce(force gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_apply_central_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a positioned force to the body. A force is time dependent and meant to be applied every physics update.
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) ApplyForce(force gd.Vector2, position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, force)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_apply_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a rotational force without affecting position. A force is time dependent and meant to be applied every physics update.
[b]Note:[/b] [member inertia] is required for this to work. To have [member inertia], an active [CollisionShape2D] must be a child of the node, or you can manually set [member inertia].
*/
//go:nosplit
func (self class) ApplyTorque(torque gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_apply_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a constant directional force without affecting rotation that keeps being applied over time until cleared with [code]constant_force = Vector2(0, 0)[/code].
This is equivalent to using [method add_constant_force] at the body's center of mass.
*/
//go:nosplit
func (self class) AddConstantCentralForce(force gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_add_constant_central_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a constant positioned force to the body that keeps being applied over time until cleared with [code]constant_force = Vector2(0, 0)[/code].
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) AddConstantForce(force gd.Vector2, position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, force)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_add_constant_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a constant rotational force without affecting position that keeps being applied over time until cleared with [code]constant_torque = 0[/code].
*/
//go:nosplit
func (self class) AddConstantTorque(torque gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_add_constant_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetConstantForce(force gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_constant_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetConstantForce() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_get_constant_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetConstantTorque(torque gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_constant_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetConstantTorque() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_get_constant_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSleeping(sleeping bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sleeping)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_sleeping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSleeping() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_is_sleeping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCanSleep(able_to_sleep bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, able_to_sleep)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_can_sleep, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAbleToSleep() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_is_able_to_sleep, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLockRotationEnabled(lock_rotation bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, lock_rotation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_lock_rotation_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsLockRotationEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_is_lock_rotation_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFreezeEnabled(freeze_mode bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, freeze_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_freeze_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFreezeEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_is_freeze_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFreezeMode(freeze_mode classdb.RigidBody2DFreezeMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, freeze_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_set_freeze_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFreezeMode() classdb.RigidBody2DFreezeMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RigidBody2DFreezeMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_get_freeze_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a list of the bodies colliding with this one. Requires [member contact_monitor] to be set to [code]true[/code] and [member max_contacts_reported] to be set high enough to detect all the collisions.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of collisions is updated once per frame and before the physics step. Consider using signals instead.
*/
//go:nosplit
func (self class) GetCollidingBodies(ctx gd.Lifetime) gd.ArrayOf[object.Node2D] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RigidBody2D.Bind_get_colliding_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.Node2D](ret)
}

//go:nosplit
func (self class) AsRigidBody2D() Expert { return self[0].AsRigidBody2D() }


//go:nosplit
func (self Simple) AsRigidBody2D() Simple { return self[0].AsRigidBody2D() }


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
	case "_integrate_forces": return reflect.ValueOf(self._integrate_forces);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_integrate_forces": return reflect.ValueOf(self._integrate_forces);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("RigidBody2D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type FreezeMode = classdb.RigidBody2DFreezeMode

const (
/*Static body freeze mode (default). The body is not affected by gravity and forces. It can be only moved by user code and doesn't collide with other bodies along its path.*/
	FreezeModeStatic FreezeMode = 0
/*Kinematic body freeze mode. Similar to [constant FREEZE_MODE_STATIC], but collides with other bodies along its path when moved. Useful for a frozen body that needs to be animated.*/
	FreezeModeKinematic FreezeMode = 1
)
type CenterOfMassMode = classdb.RigidBody2DCenterOfMassMode

const (
/*In this mode, the body's center of mass is calculated automatically based on its shapes. This assumes that the shapes' origins are also their center of mass.*/
	CenterOfMassModeAuto CenterOfMassMode = 0
/*In this mode, the body's center of mass is set through [member center_of_mass]. Defaults to the body's origin position.*/
	CenterOfMassModeCustom CenterOfMassMode = 1
)
type DampMode = classdb.RigidBody2DDampMode

const (
/*In this mode, the body's damping value is added to any value set in areas or the default value.*/
	DampModeCombine DampMode = 0
/*In this mode, the body's damping value replaces any value set in areas or the default value.*/
	DampModeReplace DampMode = 1
)
type CCDMode = classdb.RigidBody2DCCDMode

const (
/*Continuous collision detection disabled. This is the fastest way to detect body collisions, but can miss small, fast-moving objects.*/
	CcdModeDisabled CCDMode = 0
/*Continuous collision detection enabled using raycasting. This is faster than shapecasting but less precise.*/
	CcdModeCastRay CCDMode = 1
/*Continuous collision detection enabled using shapecasting. This is the slowest CCD method and the most precise.*/
	CcdModeCastShape CCDMode = 2
)
