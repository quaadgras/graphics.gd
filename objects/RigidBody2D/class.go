package RigidBody2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/PhysicsBody2D"
import "grow.graphics/gd/objects/CollisionObject2D"
import "grow.graphics/gd/objects/Node2D"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
[RigidBody2D] implements full 2D physics. It cannot be controlled directly, instead, you must apply forces to it (gravity, impulses, etc.), and the physics simulation will calculate the resulting movement, rotation, react to collisions, and affect other physics bodies in its path.
The body's behavior can be adjusted via [member lock_rotation], [member freeze], and [member freeze_mode]. By changing various properties of the object, such as [member mass], you can control how the physics simulation acts on it.
A rigid body will always maintain its shape and size, even when forces are applied to it. It is useful for objects that can be interacted with in an environment, such as a tree that can be knocked over or a stack of crates that can be pushed around.
If you need to override the default physics behavior, you can write a custom force integration function. See [member custom_integrator].
[b]Note:[/b] Changing the 2D transform or [member linear_velocity] of a [RigidBody2D] very often may lead to some unpredictable behaviors. If you need to directly affect the body, prefer [method _integrate_forces] as it allows you to directly access the physics state.

	// RigidBody2D methods that can be overridden by a [Class] that extends it.
	type RigidBody2D interface {
		//Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
		IntegrateForces(state objects.PhysicsDirectBodyState2D)
	}
*/
type Instance [1]classdb.RigidBody2D

/*
Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
*/
func (Instance) _integrate_forces(impl func(ptr unsafe.Pointer, state objects.PhysicsDirectBodyState2D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = objects.PhysicsDirectBodyState2D{pointers.New[classdb.PhysicsDirectBodyState2D]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, state)
	}
}

/*
Returns the number of contacts this body has with other bodies. By default, this returns 0 unless bodies are configured to monitor contacts (see [member contact_monitor]).
[b]Note:[/b] To retrieve the colliding bodies, use [method get_colliding_bodies].
*/
func (self Instance) GetContactCount() int {
	return int(int(class(self).GetContactCount()))
}

/*
Sets the body's velocity on the given axis. The velocity in the given vector axis will be set as the given vector length. This is useful for jumping behavior.
*/
func (self Instance) SetAxisVelocity(axis_velocity gd.Vector2) {
	class(self).SetAxisVelocity(axis_velocity)
}

/*
Applies a directional impulse without affecting rotation.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
This is equivalent to using [method apply_impulse] at the body's center of mass.
*/
func (self Instance) ApplyCentralImpulse() {
	class(self).ApplyCentralImpulse(gd.Vector2{0, 0})
}

/*
Applies a positioned impulse to the body.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[param position] is the offset from the body origin in global coordinates.
*/
func (self Instance) ApplyImpulse(impulse gd.Vector2) {
	class(self).ApplyImpulse(impulse, gd.Vector2{0, 0})
}

/*
Applies a rotational impulse to the body without affecting the position.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[b]Note:[/b] [member inertia] is required for this to work. To have [member inertia], an active [CollisionShape2D] must be a child of the node, or you can manually set [member inertia].
*/
func (self Instance) ApplyTorqueImpulse(torque float64) {
	class(self).ApplyTorqueImpulse(gd.Float(torque))
}

/*
Applies a directional force without affecting rotation. A force is time dependent and meant to be applied every physics update.
This is equivalent to using [method apply_force] at the body's center of mass.
*/
func (self Instance) ApplyCentralForce(force gd.Vector2) {
	class(self).ApplyCentralForce(force)
}

/*
Applies a positioned force to the body. A force is time dependent and meant to be applied every physics update.
[param position] is the offset from the body origin in global coordinates.
*/
func (self Instance) ApplyForce(force gd.Vector2) {
	class(self).ApplyForce(force, gd.Vector2{0, 0})
}

/*
Applies a rotational force without affecting position. A force is time dependent and meant to be applied every physics update.
[b]Note:[/b] [member inertia] is required for this to work. To have [member inertia], an active [CollisionShape2D] must be a child of the node, or you can manually set [member inertia].
*/
func (self Instance) ApplyTorque(torque float64) {
	class(self).ApplyTorque(gd.Float(torque))
}

/*
Adds a constant directional force without affecting rotation that keeps being applied over time until cleared with [code]constant_force = Vector2(0, 0)[/code].
This is equivalent to using [method add_constant_force] at the body's center of mass.
*/
func (self Instance) AddConstantCentralForce(force gd.Vector2) {
	class(self).AddConstantCentralForce(force)
}

/*
Adds a constant positioned force to the body that keeps being applied over time until cleared with [code]constant_force = Vector2(0, 0)[/code].
[param position] is the offset from the body origin in global coordinates.
*/
func (self Instance) AddConstantForce(force gd.Vector2) {
	class(self).AddConstantForce(force, gd.Vector2{0, 0})
}

/*
Adds a constant rotational force without affecting position that keeps being applied over time until cleared with [code]constant_torque = 0[/code].
*/
func (self Instance) AddConstantTorque(torque float64) {
	class(self).AddConstantTorque(gd.Float(torque))
}

/*
Returns a list of the bodies colliding with this one. Requires [member contact_monitor] to be set to [code]true[/code] and [member max_contacts_reported] to be set high enough to detect all the collisions.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of collisions is updated once per frame and before the physics step. Consider using signals instead.
*/
func (self Instance) GetCollidingBodies() gd.Array {
	return gd.Array(class(self).GetCollidingBodies())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RigidBody2D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RigidBody2D"))
	return Instance{classdb.RigidBody2D(object)}
}

func (self Instance) Mass() float64 {
	return float64(float64(class(self).GetMass()))
}

func (self Instance) SetMass(value float64) {
	class(self).SetMass(gd.Float(value))
}

func (self Instance) PhysicsMaterialOverride() objects.PhysicsMaterial {
	return objects.PhysicsMaterial(class(self).GetPhysicsMaterialOverride())
}

func (self Instance) SetPhysicsMaterialOverride(value objects.PhysicsMaterial) {
	class(self).SetPhysicsMaterialOverride(value)
}

func (self Instance) GravityScale() float64 {
	return float64(float64(class(self).GetGravityScale()))
}

func (self Instance) SetGravityScale(value float64) {
	class(self).SetGravityScale(gd.Float(value))
}

func (self Instance) CenterOfMassMode() classdb.RigidBody2DCenterOfMassMode {
	return classdb.RigidBody2DCenterOfMassMode(class(self).GetCenterOfMassMode())
}

func (self Instance) SetCenterOfMassMode(value classdb.RigidBody2DCenterOfMassMode) {
	class(self).SetCenterOfMassMode(value)
}

func (self Instance) CenterOfMass() gd.Vector2 {
	return gd.Vector2(class(self).GetCenterOfMass())
}

func (self Instance) SetCenterOfMass(value gd.Vector2) {
	class(self).SetCenterOfMass(value)
}

func (self Instance) Inertia() float64 {
	return float64(float64(class(self).GetInertia()))
}

func (self Instance) SetInertia(value float64) {
	class(self).SetInertia(gd.Float(value))
}

func (self Instance) Sleeping() bool {
	return bool(class(self).IsSleeping())
}

func (self Instance) SetSleeping(value bool) {
	class(self).SetSleeping(value)
}

func (self Instance) CanSleep() bool {
	return bool(class(self).IsAbleToSleep())
}

func (self Instance) SetCanSleep(value bool) {
	class(self).SetCanSleep(value)
}

func (self Instance) LockRotation() bool {
	return bool(class(self).IsLockRotationEnabled())
}

func (self Instance) SetLockRotation(value bool) {
	class(self).SetLockRotationEnabled(value)
}

func (self Instance) Freeze() bool {
	return bool(class(self).IsFreezeEnabled())
}

func (self Instance) SetFreeze(value bool) {
	class(self).SetFreezeEnabled(value)
}

func (self Instance) FreezeMode() classdb.RigidBody2DFreezeMode {
	return classdb.RigidBody2DFreezeMode(class(self).GetFreezeMode())
}

func (self Instance) SetFreezeMode(value classdb.RigidBody2DFreezeMode) {
	class(self).SetFreezeMode(value)
}

func (self Instance) CustomIntegrator() bool {
	return bool(class(self).IsUsingCustomIntegrator())
}

func (self Instance) SetCustomIntegrator(value bool) {
	class(self).SetUseCustomIntegrator(value)
}

func (self Instance) ContinuousCd() classdb.RigidBody2DCCDMode {
	return classdb.RigidBody2DCCDMode(class(self).GetContinuousCollisionDetectionMode())
}

func (self Instance) SetContinuousCd(value classdb.RigidBody2DCCDMode) {
	class(self).SetContinuousCollisionDetectionMode(value)
}

func (self Instance) ContactMonitor() bool {
	return bool(class(self).IsContactMonitorEnabled())
}

func (self Instance) SetContactMonitor(value bool) {
	class(self).SetContactMonitor(value)
}

func (self Instance) MaxContactsReported() int {
	return int(int(class(self).GetMaxContactsReported()))
}

func (self Instance) SetMaxContactsReported(value int) {
	class(self).SetMaxContactsReported(gd.Int(value))
}

func (self Instance) LinearVelocity() gd.Vector2 {
	return gd.Vector2(class(self).GetLinearVelocity())
}

func (self Instance) SetLinearVelocity(value gd.Vector2) {
	class(self).SetLinearVelocity(value)
}

func (self Instance) LinearDampMode() classdb.RigidBody2DDampMode {
	return classdb.RigidBody2DDampMode(class(self).GetLinearDampMode())
}

func (self Instance) SetLinearDampMode(value classdb.RigidBody2DDampMode) {
	class(self).SetLinearDampMode(value)
}

func (self Instance) LinearDamp() float64 {
	return float64(float64(class(self).GetLinearDamp()))
}

func (self Instance) SetLinearDamp(value float64) {
	class(self).SetLinearDamp(gd.Float(value))
}

func (self Instance) AngularVelocity() float64 {
	return float64(float64(class(self).GetAngularVelocity()))
}

func (self Instance) SetAngularVelocity(value float64) {
	class(self).SetAngularVelocity(gd.Float(value))
}

func (self Instance) AngularDampMode() classdb.RigidBody2DDampMode {
	return classdb.RigidBody2DDampMode(class(self).GetAngularDampMode())
}

func (self Instance) SetAngularDampMode(value classdb.RigidBody2DDampMode) {
	class(self).SetAngularDampMode(value)
}

func (self Instance) AngularDamp() float64 {
	return float64(float64(class(self).GetAngularDamp()))
}

func (self Instance) SetAngularDamp(value float64) {
	class(self).SetAngularDamp(gd.Float(value))
}

func (self Instance) ConstantForce() gd.Vector2 {
	return gd.Vector2(class(self).GetConstantForce())
}

func (self Instance) SetConstantForce(value gd.Vector2) {
	class(self).SetConstantForce(value)
}

func (self Instance) ConstantTorque() float64 {
	return float64(float64(class(self).GetConstantTorque()))
}

func (self Instance) SetConstantTorque(value float64) {
	class(self).SetConstantTorque(gd.Float(value))
}

/*
Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
*/
func (class) _integrate_forces(impl func(ptr unsafe.Pointer, state objects.PhysicsDirectBodyState2D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = objects.PhysicsDirectBodyState2D{pointers.New[classdb.PhysicsDirectBodyState2D]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, state)
	}
}

//go:nosplit
func (self class) SetMass(mass gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, mass)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMass() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_get_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetInertia() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_get_inertia, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInertia(inertia gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, inertia)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_inertia, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetCenterOfMassMode(mode classdb.RigidBody2DCenterOfMassMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_center_of_mass_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCenterOfMassMode() classdb.RigidBody2DCenterOfMassMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RigidBody2DCenterOfMassMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_get_center_of_mass_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCenterOfMass(center_of_mass gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, center_of_mass)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_center_of_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCenterOfMass() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_get_center_of_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPhysicsMaterialOverride(physics_material_override objects.PhysicsMaterial) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(physics_material_override[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_physics_material_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPhysicsMaterialOverride() objects.PhysicsMaterial {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_get_physics_material_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.PhysicsMaterial{classdb.PhysicsMaterial(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravityScale(gravity_scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, gravity_scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_gravity_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravityScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_get_gravity_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearDampMode(linear_damp_mode classdb.RigidBody2DDampMode) {
	var frame = callframe.New()
	callframe.Arg(frame, linear_damp_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_linear_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearDampMode() classdb.RigidBody2DDampMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RigidBody2DDampMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_get_linear_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularDampMode(angular_damp_mode classdb.RigidBody2DDampMode) {
	var frame = callframe.New()
	callframe.Arg(frame, angular_damp_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_angular_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularDampMode() classdb.RigidBody2DDampMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RigidBody2DDampMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_get_angular_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearDamp(linear_damp gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, linear_damp)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_linear_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearDamp() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_get_linear_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularDamp(angular_damp gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, angular_damp)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_angular_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularDamp() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_get_angular_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearVelocity(linear_velocity gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, linear_velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearVelocity() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_get_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularVelocity(angular_velocity gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, angular_velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularVelocity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_get_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxContactsReported(amount gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_max_contacts_reported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxContactsReported() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_get_max_contacts_reported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_get_contact_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseCustomIntegrator(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_use_custom_integrator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingCustomIntegrator() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_is_using_custom_integrator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetContactMonitor(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_contact_monitor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsContactMonitorEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_is_contact_monitor_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetContinuousCollisionDetectionMode(mode classdb.RigidBody2DCCDMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_continuous_collision_detection_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetContinuousCollisionDetectionMode() classdb.RigidBody2DCCDMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RigidBody2DCCDMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_get_continuous_collision_detection_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the body's velocity on the given axis. The velocity in the given vector axis will be set as the given vector length. This is useful for jumping behavior.
*/
//go:nosplit
func (self class) SetAxisVelocity(axis_velocity gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, axis_velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_axis_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Applies a directional impulse without affecting rotation.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
This is equivalent to using [method apply_impulse] at the body's center of mass.
*/
//go:nosplit
func (self class) ApplyCentralImpulse(impulse gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, impulse)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_apply_central_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Applies a positioned impulse to the body.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) ApplyImpulse(impulse gd.Vector2, position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, impulse)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_apply_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Applies a rotational impulse to the body without affecting the position.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[b]Note:[/b] [member inertia] is required for this to work. To have [member inertia], an active [CollisionShape2D] must be a child of the node, or you can manually set [member inertia].
*/
//go:nosplit
func (self class) ApplyTorqueImpulse(torque gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_apply_torque_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Applies a directional force without affecting rotation. A force is time dependent and meant to be applied every physics update.
This is equivalent to using [method apply_force] at the body's center of mass.
*/
//go:nosplit
func (self class) ApplyCentralForce(force gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_apply_central_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Applies a positioned force to the body. A force is time dependent and meant to be applied every physics update.
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) ApplyForce(force gd.Vector2, position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, force)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_apply_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Applies a rotational force without affecting position. A force is time dependent and meant to be applied every physics update.
[b]Note:[/b] [member inertia] is required for this to work. To have [member inertia], an active [CollisionShape2D] must be a child of the node, or you can manually set [member inertia].
*/
//go:nosplit
func (self class) ApplyTorque(torque gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_apply_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a constant directional force without affecting rotation that keeps being applied over time until cleared with [code]constant_force = Vector2(0, 0)[/code].
This is equivalent to using [method add_constant_force] at the body's center of mass.
*/
//go:nosplit
func (self class) AddConstantCentralForce(force gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_add_constant_central_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a constant positioned force to the body that keeps being applied over time until cleared with [code]constant_force = Vector2(0, 0)[/code].
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) AddConstantForce(force gd.Vector2, position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, force)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_add_constant_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a constant rotational force without affecting position that keeps being applied over time until cleared with [code]constant_torque = 0[/code].
*/
//go:nosplit
func (self class) AddConstantTorque(torque gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_add_constant_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetConstantForce(force gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_constant_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetConstantForce() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_get_constant_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetConstantTorque(torque gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_constant_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetConstantTorque() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_get_constant_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSleeping(sleeping bool) {
	var frame = callframe.New()
	callframe.Arg(frame, sleeping)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_sleeping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSleeping() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_is_sleeping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCanSleep(able_to_sleep bool) {
	var frame = callframe.New()
	callframe.Arg(frame, able_to_sleep)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_can_sleep, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsAbleToSleep() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_is_able_to_sleep, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLockRotationEnabled(lock_rotation bool) {
	var frame = callframe.New()
	callframe.Arg(frame, lock_rotation)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_lock_rotation_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsLockRotationEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_is_lock_rotation_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFreezeEnabled(freeze_mode bool) {
	var frame = callframe.New()
	callframe.Arg(frame, freeze_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_freeze_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsFreezeEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_is_freeze_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFreezeMode(freeze_mode classdb.RigidBody2DFreezeMode) {
	var frame = callframe.New()
	callframe.Arg(frame, freeze_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_set_freeze_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFreezeMode() classdb.RigidBody2DFreezeMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RigidBody2DFreezeMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_get_freeze_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a list of the bodies colliding with this one. Requires [member contact_monitor] to be set to [code]true[/code] and [member max_contacts_reported] to be set high enough to detect all the collisions.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of collisions is updated once per frame and before the physics step. Consider using signals instead.
*/
//go:nosplit
func (self class) GetCollidingBodies() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody2D.Bind_get_colliding_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self Instance) OnBodyShapeEntered(cb func(body_rid gd.RID, body objects.Node, body_shape_index int, local_shape_index int)) {
	self[0].AsObject().Connect(gd.NewStringName("body_shape_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBodyShapeExited(cb func(body_rid gd.RID, body objects.Node, body_shape_index int, local_shape_index int)) {
	self[0].AsObject().Connect(gd.NewStringName("body_shape_exited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBodyEntered(cb func(body objects.Node)) {
	self[0].AsObject().Connect(gd.NewStringName("body_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBodyExited(cb func(body objects.Node)) {
	self[0].AsObject().Connect(gd.NewStringName("body_exited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSleepingStateChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("sleeping_state_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsRigidBody2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRigidBody2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPhysicsBody2D() PhysicsBody2D.Advanced {
	return *((*PhysicsBody2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsBody2D() PhysicsBody2D.Instance {
	return *((*PhysicsBody2D.Instance)(unsafe.Pointer(&self)))
}
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
	case "_integrate_forces":
		return reflect.ValueOf(self._integrate_forces)
	default:
		return gd.VirtualByName(self.AsPhysicsBody2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_integrate_forces":
		return reflect.ValueOf(self._integrate_forces)
	default:
		return gd.VirtualByName(self.AsPhysicsBody2D(), name)
	}
}
func init() {
	classdb.Register("RigidBody2D", func(ptr gd.Object) any { return classdb.RigidBody2D(ptr) })
}

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
