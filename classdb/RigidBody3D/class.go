// Package RigidBody3D provides methods for working with RigidBody3D object instances.
package RigidBody3D

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
import "graphics.gd/classdb/PhysicsBody3D"
import "graphics.gd/classdb/CollisionObject3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Basis"

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
[RigidBody3D] implements full 3D physics. It cannot be controlled directly, instead, you must apply forces to it (gravity, impulses, etc.), and the physics simulation will calculate the resulting movement, rotation, react to collisions, and affect other physics bodies in its path.
The body's behavior can be adjusted via [member lock_rotation], [member freeze], and [member freeze_mode]. By changing various properties of the object, such as [member mass], you can control how the physics simulation acts on it.
A rigid body will always maintain its shape and size, even when forces are applied to it. It is useful for objects that can be interacted with in an environment, such as a tree that can be knocked over or a stack of crates that can be pushed around.
If you need to override the default physics behavior, you can write a custom force integration function. See [member custom_integrator].
[b]Note:[/b] Changing the 3D transform or [member linear_velocity] of a [RigidBody3D] very often may lead to some unpredictable behaviors. If you need to directly affect the body, prefer [method _integrate_forces] as it allows you to directly access the physics state.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=RigidBody3D)
*/
type Instance [1]gdclass.RigidBody3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRigidBody3D() Instance
}
type Interface interface {
	//Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
	IntegrateForces(state [1]gdclass.PhysicsDirectBodyState3D)
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) IntegrateForces(state [1]gdclass.PhysicsDirectBodyState3D) { return }

/*
Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
*/
func (Instance) _integrate_forces(impl func(ptr unsafe.Pointer, state [1]gdclass.PhysicsDirectBodyState3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var state = [1]gdclass.PhysicsDirectBodyState3D{pointers.New[gdclass.PhysicsDirectBodyState3D]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(state[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, state)
	}
}

/*
Returns the inverse inertia tensor basis. This is used to calculate the angular acceleration resulting from a torque applied to the [RigidBody3D].
*/
func (self Instance) GetInverseInertiaTensor() Basis.XYZ { //gd:RigidBody3D.get_inverse_inertia_tensor
	return Basis.XYZ(class(self).GetInverseInertiaTensor())
}

/*
Returns the number of contacts this body has with other bodies. By default, this returns 0 unless bodies are configured to monitor contacts (see [member contact_monitor]).
[b]Note:[/b] To retrieve the colliding bodies, use [method get_colliding_bodies].
*/
func (self Instance) GetContactCount() int { //gd:RigidBody3D.get_contact_count
	return int(int(class(self).GetContactCount()))
}

/*
Sets an axis velocity. The velocity in the given vector axis will be set as the given vector length. This is useful for jumping behavior.
*/
func (self Instance) SetAxisVelocity(axis_velocity Vector3.XYZ) { //gd:RigidBody3D.set_axis_velocity
	class(self).SetAxisVelocity(gd.Vector3(axis_velocity))
}

/*
Applies a directional impulse without affecting rotation.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
This is equivalent to using [method apply_impulse] at the body's center of mass.
*/
func (self Instance) ApplyCentralImpulse(impulse Vector3.XYZ) { //gd:RigidBody3D.apply_central_impulse
	class(self).ApplyCentralImpulse(gd.Vector3(impulse))
}

/*
Applies a positioned impulse to the body.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[param position] is the offset from the body origin in global coordinates.
*/
func (self Instance) ApplyImpulse(impulse Vector3.XYZ) { //gd:RigidBody3D.apply_impulse
	class(self).ApplyImpulse(gd.Vector3(impulse), gd.Vector3(gd.Vector3{0, 0, 0}))
}

/*
Applies a rotational impulse to the body without affecting the position.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[b]Note:[/b] [member inertia] is required for this to work. To have [member inertia], an active [CollisionShape3D] must be a child of the node, or you can manually set [member inertia].
*/
func (self Instance) ApplyTorqueImpulse(impulse Vector3.XYZ) { //gd:RigidBody3D.apply_torque_impulse
	class(self).ApplyTorqueImpulse(gd.Vector3(impulse))
}

/*
Applies a directional force without affecting rotation. A force is time dependent and meant to be applied every physics update.
This is equivalent to using [method apply_force] at the body's center of mass.
*/
func (self Instance) ApplyCentralForce(force Vector3.XYZ) { //gd:RigidBody3D.apply_central_force
	class(self).ApplyCentralForce(gd.Vector3(force))
}

/*
Applies a positioned force to the body. A force is time dependent and meant to be applied every physics update.
[param position] is the offset from the body origin in global coordinates.
*/
func (self Instance) ApplyForce(force Vector3.XYZ) { //gd:RigidBody3D.apply_force
	class(self).ApplyForce(gd.Vector3(force), gd.Vector3(gd.Vector3{0, 0, 0}))
}

/*
Applies a rotational force without affecting position. A force is time dependent and meant to be applied every physics update.
[b]Note:[/b] [member inertia] is required for this to work. To have [member inertia], an active [CollisionShape3D] must be a child of the node, or you can manually set [member inertia].
*/
func (self Instance) ApplyTorque(torque Vector3.XYZ) { //gd:RigidBody3D.apply_torque
	class(self).ApplyTorque(gd.Vector3(torque))
}

/*
Adds a constant directional force without affecting rotation that keeps being applied over time until cleared with [code]constant_force = Vector3(0, 0, 0)[/code].
This is equivalent to using [method add_constant_force] at the body's center of mass.
*/
func (self Instance) AddConstantCentralForce(force Vector3.XYZ) { //gd:RigidBody3D.add_constant_central_force
	class(self).AddConstantCentralForce(gd.Vector3(force))
}

/*
Adds a constant positioned force to the body that keeps being applied over time until cleared with [code]constant_force = Vector3(0, 0, 0)[/code].
[param position] is the offset from the body origin in global coordinates.
*/
func (self Instance) AddConstantForce(force Vector3.XYZ) { //gd:RigidBody3D.add_constant_force
	class(self).AddConstantForce(gd.Vector3(force), gd.Vector3(gd.Vector3{0, 0, 0}))
}

/*
Adds a constant rotational force without affecting position that keeps being applied over time until cleared with [code]constant_torque = Vector3(0, 0, 0)[/code].
*/
func (self Instance) AddConstantTorque(torque Vector3.XYZ) { //gd:RigidBody3D.add_constant_torque
	class(self).AddConstantTorque(gd.Vector3(torque))
}

/*
Returns a list of the bodies colliding with this one. Requires [member contact_monitor] to be set to [code]true[/code] and [member max_contacts_reported] to be set high enough to detect all the collisions.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of collisions is updated once per frame and before the physics step. Consider using signals instead.
*/
func (self Instance) GetCollidingBodies() [][1]gdclass.Node3D { //gd:RigidBody3D.get_colliding_bodies
	return [][1]gdclass.Node3D(gd.ArrayAs[[][1]gdclass.Node3D](gd.InternalArray(class(self).GetCollidingBodies())))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RigidBody3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RigidBody3D"))
	casted := Instance{*(*gdclass.RigidBody3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Mass() Float.X {
	return Float.X(Float.X(class(self).GetMass()))
}

func (self Instance) SetMass(value Float.X) {
	class(self).SetMass(gd.Float(value))
}

func (self Instance) PhysicsMaterialOverride() [1]gdclass.PhysicsMaterial {
	return [1]gdclass.PhysicsMaterial(class(self).GetPhysicsMaterialOverride())
}

func (self Instance) SetPhysicsMaterialOverride(value [1]gdclass.PhysicsMaterial) {
	class(self).SetPhysicsMaterialOverride(value)
}

func (self Instance) GravityScale() Float.X {
	return Float.X(Float.X(class(self).GetGravityScale()))
}

func (self Instance) SetGravityScale(value Float.X) {
	class(self).SetGravityScale(gd.Float(value))
}

func (self Instance) CenterOfMassMode() gdclass.RigidBody3DCenterOfMassMode {
	return gdclass.RigidBody3DCenterOfMassMode(class(self).GetCenterOfMassMode())
}

func (self Instance) SetCenterOfMassMode(value gdclass.RigidBody3DCenterOfMassMode) {
	class(self).SetCenterOfMassMode(value)
}

func (self Instance) CenterOfMass() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetCenterOfMass())
}

func (self Instance) SetCenterOfMass(value Vector3.XYZ) {
	class(self).SetCenterOfMass(gd.Vector3(value))
}

func (self Instance) Inertia() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetInertia())
}

func (self Instance) SetInertia(value Vector3.XYZ) {
	class(self).SetInertia(gd.Vector3(value))
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

func (self Instance) FreezeMode() gdclass.RigidBody3DFreezeMode {
	return gdclass.RigidBody3DFreezeMode(class(self).GetFreezeMode())
}

func (self Instance) SetFreezeMode(value gdclass.RigidBody3DFreezeMode) {
	class(self).SetFreezeMode(value)
}

func (self Instance) CustomIntegrator() bool {
	return bool(class(self).IsUsingCustomIntegrator())
}

func (self Instance) SetCustomIntegrator(value bool) {
	class(self).SetUseCustomIntegrator(value)
}

func (self Instance) ContinuousCd() bool {
	return bool(class(self).IsUsingContinuousCollisionDetection())
}

func (self Instance) SetContinuousCd(value bool) {
	class(self).SetUseContinuousCollisionDetection(value)
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

func (self Instance) LinearVelocity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetLinearVelocity())
}

func (self Instance) SetLinearVelocity(value Vector3.XYZ) {
	class(self).SetLinearVelocity(gd.Vector3(value))
}

func (self Instance) LinearDampMode() gdclass.RigidBody3DDampMode {
	return gdclass.RigidBody3DDampMode(class(self).GetLinearDampMode())
}

func (self Instance) SetLinearDampMode(value gdclass.RigidBody3DDampMode) {
	class(self).SetLinearDampMode(value)
}

func (self Instance) LinearDamp() Float.X {
	return Float.X(Float.X(class(self).GetLinearDamp()))
}

func (self Instance) SetLinearDamp(value Float.X) {
	class(self).SetLinearDamp(gd.Float(value))
}

func (self Instance) AngularVelocity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetAngularVelocity())
}

func (self Instance) SetAngularVelocity(value Vector3.XYZ) {
	class(self).SetAngularVelocity(gd.Vector3(value))
}

func (self Instance) AngularDampMode() gdclass.RigidBody3DDampMode {
	return gdclass.RigidBody3DDampMode(class(self).GetAngularDampMode())
}

func (self Instance) SetAngularDampMode(value gdclass.RigidBody3DDampMode) {
	class(self).SetAngularDampMode(value)
}

func (self Instance) AngularDamp() Float.X {
	return Float.X(Float.X(class(self).GetAngularDamp()))
}

func (self Instance) SetAngularDamp(value Float.X) {
	class(self).SetAngularDamp(gd.Float(value))
}

func (self Instance) ConstantForce() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetConstantForce())
}

func (self Instance) SetConstantForce(value Vector3.XYZ) {
	class(self).SetConstantForce(gd.Vector3(value))
}

func (self Instance) ConstantTorque() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetConstantTorque())
}

func (self Instance) SetConstantTorque(value Vector3.XYZ) {
	class(self).SetConstantTorque(gd.Vector3(value))
}

/*
Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
*/
func (class) _integrate_forces(impl func(ptr unsafe.Pointer, state [1]gdclass.PhysicsDirectBodyState3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var state = [1]gdclass.PhysicsDirectBodyState3D{pointers.New[gdclass.PhysicsDirectBodyState3D]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(state[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, state)
	}
}

//go:nosplit
func (self class) SetMass(mass gd.Float) { //gd:RigidBody3D.set_mass
	var frame = callframe.New()
	callframe.Arg(frame, mass)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_mass, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMass() gd.Float { //gd:RigidBody3D.get_mass
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_mass, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInertia(inertia gd.Vector3) { //gd:RigidBody3D.set_inertia
	var frame = callframe.New()
	callframe.Arg(frame, inertia)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_inertia, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInertia() gd.Vector3 { //gd:RigidBody3D.get_inertia
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_inertia, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCenterOfMassMode(mode gdclass.RigidBody3DCenterOfMassMode) { //gd:RigidBody3D.set_center_of_mass_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_center_of_mass_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCenterOfMassMode() gdclass.RigidBody3DCenterOfMassMode { //gd:RigidBody3D.get_center_of_mass_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RigidBody3DCenterOfMassMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_center_of_mass_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCenterOfMass(center_of_mass gd.Vector3) { //gd:RigidBody3D.set_center_of_mass
	var frame = callframe.New()
	callframe.Arg(frame, center_of_mass)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_center_of_mass, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCenterOfMass() gd.Vector3 { //gd:RigidBody3D.get_center_of_mass
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_center_of_mass, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPhysicsMaterialOverride(physics_material_override [1]gdclass.PhysicsMaterial) { //gd:RigidBody3D.set_physics_material_override
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(physics_material_override[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_physics_material_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPhysicsMaterialOverride() [1]gdclass.PhysicsMaterial { //gd:RigidBody3D.get_physics_material_override
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_physics_material_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PhysicsMaterial{gd.PointerWithOwnershipTransferredToGo[gdclass.PhysicsMaterial](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearVelocity(linear_velocity gd.Vector3) { //gd:RigidBody3D.set_linear_velocity
	var frame = callframe.New()
	callframe.Arg(frame, linear_velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearVelocity() gd.Vector3 { //gd:RigidBody3D.get_linear_velocity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularVelocity(angular_velocity gd.Vector3) { //gd:RigidBody3D.set_angular_velocity
	var frame = callframe.New()
	callframe.Arg(frame, angular_velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularVelocity() gd.Vector3 { //gd:RigidBody3D.get_angular_velocity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the inverse inertia tensor basis. This is used to calculate the angular acceleration resulting from a torque applied to the [RigidBody3D].
*/
//go:nosplit
func (self class) GetInverseInertiaTensor() gd.Basis { //gd:RigidBody3D.get_inverse_inertia_tensor
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Basis](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_inverse_inertia_tensor, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravityScale(gravity_scale gd.Float) { //gd:RigidBody3D.set_gravity_scale
	var frame = callframe.New()
	callframe.Arg(frame, gravity_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_gravity_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravityScale() gd.Float { //gd:RigidBody3D.get_gravity_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_gravity_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearDampMode(linear_damp_mode gdclass.RigidBody3DDampMode) { //gd:RigidBody3D.set_linear_damp_mode
	var frame = callframe.New()
	callframe.Arg(frame, linear_damp_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_linear_damp_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearDampMode() gdclass.RigidBody3DDampMode { //gd:RigidBody3D.get_linear_damp_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RigidBody3DDampMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_linear_damp_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularDampMode(angular_damp_mode gdclass.RigidBody3DDampMode) { //gd:RigidBody3D.set_angular_damp_mode
	var frame = callframe.New()
	callframe.Arg(frame, angular_damp_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_angular_damp_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularDampMode() gdclass.RigidBody3DDampMode { //gd:RigidBody3D.get_angular_damp_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RigidBody3DDampMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_angular_damp_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearDamp(linear_damp gd.Float) { //gd:RigidBody3D.set_linear_damp
	var frame = callframe.New()
	callframe.Arg(frame, linear_damp)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_linear_damp, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearDamp() gd.Float { //gd:RigidBody3D.get_linear_damp
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_linear_damp, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularDamp(angular_damp gd.Float) { //gd:RigidBody3D.set_angular_damp
	var frame = callframe.New()
	callframe.Arg(frame, angular_damp)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_angular_damp, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularDamp() gd.Float { //gd:RigidBody3D.get_angular_damp
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_angular_damp, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxContactsReported(amount gd.Int) { //gd:RigidBody3D.set_max_contacts_reported
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_max_contacts_reported, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxContactsReported() gd.Int { //gd:RigidBody3D.get_max_contacts_reported
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_max_contacts_reported, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of contacts this body has with other bodies. By default, this returns 0 unless bodies are configured to monitor contacts (see [member contact_monitor]).
[b]Note:[/b] To retrieve the colliding bodies, use [method get_colliding_bodies].
*/
//go:nosplit
func (self class) GetContactCount() gd.Int { //gd:RigidBody3D.get_contact_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_contact_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseCustomIntegrator(enable bool) { //gd:RigidBody3D.set_use_custom_integrator
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_use_custom_integrator, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingCustomIntegrator() bool { //gd:RigidBody3D.is_using_custom_integrator
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_is_using_custom_integrator, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetContactMonitor(enabled bool) { //gd:RigidBody3D.set_contact_monitor
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_contact_monitor, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsContactMonitorEnabled() bool { //gd:RigidBody3D.is_contact_monitor_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_is_contact_monitor_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseContinuousCollisionDetection(enable bool) { //gd:RigidBody3D.set_use_continuous_collision_detection
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_use_continuous_collision_detection, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingContinuousCollisionDetection() bool { //gd:RigidBody3D.is_using_continuous_collision_detection
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_is_using_continuous_collision_detection, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets an axis velocity. The velocity in the given vector axis will be set as the given vector length. This is useful for jumping behavior.
*/
//go:nosplit
func (self class) SetAxisVelocity(axis_velocity gd.Vector3) { //gd:RigidBody3D.set_axis_velocity
	var frame = callframe.New()
	callframe.Arg(frame, axis_velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_axis_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Applies a directional impulse without affecting rotation.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
This is equivalent to using [method apply_impulse] at the body's center of mass.
*/
//go:nosplit
func (self class) ApplyCentralImpulse(impulse gd.Vector3) { //gd:RigidBody3D.apply_central_impulse
	var frame = callframe.New()
	callframe.Arg(frame, impulse)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_apply_central_impulse, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Applies a positioned impulse to the body.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) ApplyImpulse(impulse gd.Vector3, position gd.Vector3) { //gd:RigidBody3D.apply_impulse
	var frame = callframe.New()
	callframe.Arg(frame, impulse)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_apply_impulse, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Applies a rotational impulse to the body without affecting the position.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[b]Note:[/b] [member inertia] is required for this to work. To have [member inertia], an active [CollisionShape3D] must be a child of the node, or you can manually set [member inertia].
*/
//go:nosplit
func (self class) ApplyTorqueImpulse(impulse gd.Vector3) { //gd:RigidBody3D.apply_torque_impulse
	var frame = callframe.New()
	callframe.Arg(frame, impulse)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_apply_torque_impulse, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Applies a directional force without affecting rotation. A force is time dependent and meant to be applied every physics update.
This is equivalent to using [method apply_force] at the body's center of mass.
*/
//go:nosplit
func (self class) ApplyCentralForce(force gd.Vector3) { //gd:RigidBody3D.apply_central_force
	var frame = callframe.New()
	callframe.Arg(frame, force)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_apply_central_force, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Applies a positioned force to the body. A force is time dependent and meant to be applied every physics update.
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) ApplyForce(force gd.Vector3, position gd.Vector3) { //gd:RigidBody3D.apply_force
	var frame = callframe.New()
	callframe.Arg(frame, force)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_apply_force, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Applies a rotational force without affecting position. A force is time dependent and meant to be applied every physics update.
[b]Note:[/b] [member inertia] is required for this to work. To have [member inertia], an active [CollisionShape3D] must be a child of the node, or you can manually set [member inertia].
*/
//go:nosplit
func (self class) ApplyTorque(torque gd.Vector3) { //gd:RigidBody3D.apply_torque
	var frame = callframe.New()
	callframe.Arg(frame, torque)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_apply_torque, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a constant directional force without affecting rotation that keeps being applied over time until cleared with [code]constant_force = Vector3(0, 0, 0)[/code].
This is equivalent to using [method add_constant_force] at the body's center of mass.
*/
//go:nosplit
func (self class) AddConstantCentralForce(force gd.Vector3) { //gd:RigidBody3D.add_constant_central_force
	var frame = callframe.New()
	callframe.Arg(frame, force)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_add_constant_central_force, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a constant positioned force to the body that keeps being applied over time until cleared with [code]constant_force = Vector3(0, 0, 0)[/code].
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) AddConstantForce(force gd.Vector3, position gd.Vector3) { //gd:RigidBody3D.add_constant_force
	var frame = callframe.New()
	callframe.Arg(frame, force)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_add_constant_force, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a constant rotational force without affecting position that keeps being applied over time until cleared with [code]constant_torque = Vector3(0, 0, 0)[/code].
*/
//go:nosplit
func (self class) AddConstantTorque(torque gd.Vector3) { //gd:RigidBody3D.add_constant_torque
	var frame = callframe.New()
	callframe.Arg(frame, torque)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_add_constant_torque, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetConstantForce(force gd.Vector3) { //gd:RigidBody3D.set_constant_force
	var frame = callframe.New()
	callframe.Arg(frame, force)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_constant_force, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetConstantForce() gd.Vector3 { //gd:RigidBody3D.get_constant_force
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_constant_force, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetConstantTorque(torque gd.Vector3) { //gd:RigidBody3D.set_constant_torque
	var frame = callframe.New()
	callframe.Arg(frame, torque)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_constant_torque, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetConstantTorque() gd.Vector3 { //gd:RigidBody3D.get_constant_torque
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_constant_torque, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSleeping(sleeping bool) { //gd:RigidBody3D.set_sleeping
	var frame = callframe.New()
	callframe.Arg(frame, sleeping)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_sleeping, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSleeping() bool { //gd:RigidBody3D.is_sleeping
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_is_sleeping, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCanSleep(able_to_sleep bool) { //gd:RigidBody3D.set_can_sleep
	var frame = callframe.New()
	callframe.Arg(frame, able_to_sleep)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_can_sleep, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAbleToSleep() bool { //gd:RigidBody3D.is_able_to_sleep
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_is_able_to_sleep, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLockRotationEnabled(lock_rotation bool) { //gd:RigidBody3D.set_lock_rotation_enabled
	var frame = callframe.New()
	callframe.Arg(frame, lock_rotation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_lock_rotation_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsLockRotationEnabled() bool { //gd:RigidBody3D.is_lock_rotation_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_is_lock_rotation_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFreezeEnabled(freeze_mode bool) { //gd:RigidBody3D.set_freeze_enabled
	var frame = callframe.New()
	callframe.Arg(frame, freeze_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_freeze_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFreezeEnabled() bool { //gd:RigidBody3D.is_freeze_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_is_freeze_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFreezeMode(freeze_mode gdclass.RigidBody3DFreezeMode) { //gd:RigidBody3D.set_freeze_mode
	var frame = callframe.New()
	callframe.Arg(frame, freeze_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_freeze_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFreezeMode() gdclass.RigidBody3DFreezeMode { //gd:RigidBody3D.get_freeze_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RigidBody3DFreezeMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_freeze_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a list of the bodies colliding with this one. Requires [member contact_monitor] to be set to [code]true[/code] and [member max_contacts_reported] to be set high enough to detect all the collisions.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of collisions is updated once per frame and before the physics step. Consider using signals instead.
*/
//go:nosplit
func (self class) GetCollidingBodies() Array.Contains[[1]gdclass.Node3D] { //gd:RigidBody3D.get_colliding_bodies
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_colliding_bodies, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.Node3D]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}
func (self Instance) OnBodyShapeEntered(cb func(body_rid RID.Any, body [1]gdclass.Node, body_shape_index int, local_shape_index int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("body_shape_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBodyShapeExited(cb func(body_rid RID.Any, body [1]gdclass.Node, body_shape_index int, local_shape_index int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("body_shape_exited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBodyEntered(cb func(body [1]gdclass.Node)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("body_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBodyExited(cb func(body [1]gdclass.Node)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("body_exited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSleepingStateChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("sleeping_state_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsRigidBody3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRigidBody3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPhysicsBody3D() PhysicsBody3D.Advanced {
	return *((*PhysicsBody3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsBody3D() PhysicsBody3D.Instance {
	return *((*PhysicsBody3D.Instance)(unsafe.Pointer(&self)))
}
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
	case "_integrate_forces":
		return reflect.ValueOf(self._integrate_forces)
	default:
		return gd.VirtualByName(PhysicsBody3D.Advanced(self.AsPhysicsBody3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_integrate_forces":
		return reflect.ValueOf(self._integrate_forces)
	default:
		return gd.VirtualByName(PhysicsBody3D.Instance(self.AsPhysicsBody3D()), name)
	}
}
func init() {
	gdclass.Register("RigidBody3D", func(ptr gd.Object) any { return [1]gdclass.RigidBody3D{*(*gdclass.RigidBody3D)(unsafe.Pointer(&ptr))} })
}

type FreezeMode = gdclass.RigidBody3DFreezeMode //gd:RigidBody3D.FreezeMode

const (
	/*Static body freeze mode (default). The body is not affected by gravity and forces. It can be only moved by user code and doesn't collide with other bodies along its path.*/
	FreezeModeStatic FreezeMode = 0
	/*Kinematic body freeze mode. Similar to [constant FREEZE_MODE_STATIC], but collides with other bodies along its path when moved. Useful for a frozen body that needs to be animated.*/
	FreezeModeKinematic FreezeMode = 1
)

type CenterOfMassMode = gdclass.RigidBody3DCenterOfMassMode //gd:RigidBody3D.CenterOfMassMode

const (
	/*In this mode, the body's center of mass is calculated automatically based on its shapes. This assumes that the shapes' origins are also their center of mass.*/
	CenterOfMassModeAuto CenterOfMassMode = 0
	/*In this mode, the body's center of mass is set through [member center_of_mass]. Defaults to the body's origin position.*/
	CenterOfMassModeCustom CenterOfMassMode = 1
)

type DampMode = gdclass.RigidBody3DDampMode //gd:RigidBody3D.DampMode

const (
	/*In this mode, the body's damping value is added to any value set in areas or the default value.*/
	DampModeCombine DampMode = 0
	/*In this mode, the body's damping value replaces any value set in areas or the default value.*/
	DampModeReplace DampMode = 1
)
