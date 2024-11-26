package PhysicsDirectBodyState3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/variant/Vector3"
import "grow.graphics/gd/variant/Float"
import "grow.graphics/gd/variant/Basis"
import "grow.graphics/gd/variant/Transform3D"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Provides direct access to a physics body in the [PhysicsServer3D], allowing safe changes to physics properties. This object is passed via the direct state callback of [RigidBody3D], and is intended for changing the direct state of that body. See [method RigidBody3D._integrate_forces].
*/
type Instance [1]classdb.PhysicsDirectBodyState3D

/*
Returns the body's velocity at the given relative position, including both translation and rotation.
*/
func (self Instance) GetVelocityAtLocalPosition(local_position Vector3.XYZ) Vector3.XYZ {
	return Vector3.XYZ(class(self).GetVelocityAtLocalPosition(gd.Vector3(local_position)))
}

/*
Applies a directional impulse without affecting rotation.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
This is equivalent to using [method apply_impulse] at the body's center of mass.
*/
func (self Instance) ApplyCentralImpulse() {
	class(self).ApplyCentralImpulse(gd.Vector3(gd.Vector3{0, 0, 0}))
}

/*
Applies a positioned impulse to the body.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[param position] is the offset from the body origin in global coordinates.
*/
func (self Instance) ApplyImpulse(impulse Vector3.XYZ) {
	class(self).ApplyImpulse(gd.Vector3(impulse), gd.Vector3(gd.Vector3{0, 0, 0}))
}

/*
Applies a rotational impulse to the body without affecting the position.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[b]Note:[/b] [member inverse_inertia] is required for this to work. To have [member inverse_inertia], an active [CollisionShape3D] must be a child of the node, or you can manually set [member inverse_inertia].
*/
func (self Instance) ApplyTorqueImpulse(impulse Vector3.XYZ) {
	class(self).ApplyTorqueImpulse(gd.Vector3(impulse))
}

/*
Applies a directional force without affecting rotation. A force is time dependent and meant to be applied every physics update.
This is equivalent to using [method apply_force] at the body's center of mass.
*/
func (self Instance) ApplyCentralForce() {
	class(self).ApplyCentralForce(gd.Vector3(gd.Vector3{0, 0, 0}))
}

/*
Applies a positioned force to the body. A force is time dependent and meant to be applied every physics update.
[param position] is the offset from the body origin in global coordinates.
*/
func (self Instance) ApplyForce(force Vector3.XYZ) {
	class(self).ApplyForce(gd.Vector3(force), gd.Vector3(gd.Vector3{0, 0, 0}))
}

/*
Applies a rotational force without affecting position. A force is time dependent and meant to be applied every physics update.
[b]Note:[/b] [member inverse_inertia] is required for this to work. To have [member inverse_inertia], an active [CollisionShape3D] must be a child of the node, or you can manually set [member inverse_inertia].
*/
func (self Instance) ApplyTorque(torque Vector3.XYZ) {
	class(self).ApplyTorque(gd.Vector3(torque))
}

/*
Adds a constant directional force without affecting rotation that keeps being applied over time until cleared with [code]constant_force = Vector3(0, 0, 0)[/code].
This is equivalent to using [method add_constant_force] at the body's center of mass.
*/
func (self Instance) AddConstantCentralForce() {
	class(self).AddConstantCentralForce(gd.Vector3(gd.Vector3{0, 0, 0}))
}

/*
Adds a constant positioned force to the body that keeps being applied over time until cleared with [code]constant_force = Vector3(0, 0, 0)[/code].
[param position] is the offset from the body origin in global coordinates.
*/
func (self Instance) AddConstantForce(force Vector3.XYZ) {
	class(self).AddConstantForce(gd.Vector3(force), gd.Vector3(gd.Vector3{0, 0, 0}))
}

/*
Adds a constant rotational force without affecting position that keeps being applied over time until cleared with [code]constant_torque = Vector3(0, 0, 0)[/code].
*/
func (self Instance) AddConstantTorque(torque Vector3.XYZ) {
	class(self).AddConstantTorque(gd.Vector3(torque))
}

/*
Sets the body's total constant positional forces applied during each physics update.
See [method add_constant_force] and [method add_constant_central_force].
*/
func (self Instance) SetConstantForce(force Vector3.XYZ) {
	class(self).SetConstantForce(gd.Vector3(force))
}

/*
Returns the body's total constant positional forces applied during each physics update.
See [method add_constant_force] and [method add_constant_central_force].
*/
func (self Instance) GetConstantForce() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetConstantForce())
}

/*
Sets the body's total constant rotational forces applied during each physics update.
See [method add_constant_torque].
*/
func (self Instance) SetConstantTorque(torque Vector3.XYZ) {
	class(self).SetConstantTorque(gd.Vector3(torque))
}

/*
Returns the body's total constant rotational forces applied during each physics update.
See [method add_constant_torque].
*/
func (self Instance) GetConstantTorque() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetConstantTorque())
}

/*
Returns the number of contacts this body has with other bodies.
[b]Note:[/b] By default, this returns 0 unless bodies are configured to monitor contacts. See [member RigidBody3D.contact_monitor].
*/
func (self Instance) GetContactCount() int {
	return int(int(class(self).GetContactCount()))
}

/*
Returns the position of the contact point on the body in the global coordinate system.
*/
func (self Instance) GetContactLocalPosition(contact_idx int) Vector3.XYZ {
	return Vector3.XYZ(class(self).GetContactLocalPosition(gd.Int(contact_idx)))
}

/*
Returns the local normal at the contact point.
*/
func (self Instance) GetContactLocalNormal(contact_idx int) Vector3.XYZ {
	return Vector3.XYZ(class(self).GetContactLocalNormal(gd.Int(contact_idx)))
}

/*
Impulse created by the contact.
*/
func (self Instance) GetContactImpulse(contact_idx int) Vector3.XYZ {
	return Vector3.XYZ(class(self).GetContactImpulse(gd.Int(contact_idx)))
}

/*
Returns the local shape index of the collision.
*/
func (self Instance) GetContactLocalShape(contact_idx int) int {
	return int(int(class(self).GetContactLocalShape(gd.Int(contact_idx))))
}

/*
Returns the linear velocity vector at the body's contact point.
*/
func (self Instance) GetContactLocalVelocityAtPosition(contact_idx int) Vector3.XYZ {
	return Vector3.XYZ(class(self).GetContactLocalVelocityAtPosition(gd.Int(contact_idx)))
}

/*
Returns the collider's [RID].
*/
func (self Instance) GetContactCollider(contact_idx int) Resource.ID {
	return Resource.ID(class(self).GetContactCollider(gd.Int(contact_idx)))
}

/*
Returns the position of the contact point on the collider in the global coordinate system.
*/
func (self Instance) GetContactColliderPosition(contact_idx int) Vector3.XYZ {
	return Vector3.XYZ(class(self).GetContactColliderPosition(gd.Int(contact_idx)))
}

/*
Returns the collider's object id.
*/
func (self Instance) GetContactColliderId(contact_idx int) int {
	return int(int(class(self).GetContactColliderId(gd.Int(contact_idx))))
}

/*
Returns the collider object.
*/
func (self Instance) GetContactColliderObject(contact_idx int) gd.Object {
	return gd.Object(class(self).GetContactColliderObject(gd.Int(contact_idx)))
}

/*
Returns the collider's shape index.
*/
func (self Instance) GetContactColliderShape(contact_idx int) int {
	return int(int(class(self).GetContactColliderShape(gd.Int(contact_idx))))
}

/*
Returns the linear velocity vector at the collider's contact point.
*/
func (self Instance) GetContactColliderVelocityAtPosition(contact_idx int) Vector3.XYZ {
	return Vector3.XYZ(class(self).GetContactColliderVelocityAtPosition(gd.Int(contact_idx)))
}

/*
Updates the body's linear and angular velocity by applying gravity and damping for the equivalent of one physics tick.
*/
func (self Instance) IntegrateForces() {
	class(self).IntegrateForces()
}

/*
Returns the current state of the space, useful for queries.
*/
func (self Instance) GetSpaceState() objects.PhysicsDirectSpaceState3D {
	return objects.PhysicsDirectSpaceState3D(class(self).GetSpaceState())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.PhysicsDirectBodyState3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsDirectBodyState3D"))
	return Instance{classdb.PhysicsDirectBodyState3D(object)}
}

func (self Instance) Step() Float.X {
	return Float.X(Float.X(class(self).GetStep()))
}

func (self Instance) InverseMass() Float.X {
	return Float.X(Float.X(class(self).GetInverseMass()))
}

func (self Instance) TotalAngularDamp() Float.X {
	return Float.X(Float.X(class(self).GetTotalAngularDamp()))
}

func (self Instance) TotalLinearDamp() Float.X {
	return Float.X(Float.X(class(self).GetTotalLinearDamp()))
}

func (self Instance) InverseInertia() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetInverseInertia())
}

func (self Instance) InverseInertiaTensor() Basis.XYZ {
	return Basis.XYZ(class(self).GetInverseInertiaTensor())
}

func (self Instance) TotalGravity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetTotalGravity())
}

func (self Instance) CenterOfMass() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetCenterOfMass())
}

func (self Instance) CenterOfMassLocal() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetCenterOfMassLocal())
}

func (self Instance) PrincipalInertiaAxes() Basis.XYZ {
	return Basis.XYZ(class(self).GetPrincipalInertiaAxes())
}

func (self Instance) AngularVelocity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetAngularVelocity())
}

func (self Instance) SetAngularVelocity(value Vector3.XYZ) {
	class(self).SetAngularVelocity(gd.Vector3(value))
}

func (self Instance) LinearVelocity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetLinearVelocity())
}

func (self Instance) SetLinearVelocity(value Vector3.XYZ) {
	class(self).SetLinearVelocity(gd.Vector3(value))
}

func (self Instance) Sleeping() bool {
	return bool(class(self).IsSleeping())
}

func (self Instance) SetSleeping(value bool) {
	class(self).SetSleepState(value)
}

func (self Instance) Transform() Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetTransform())
}

func (self Instance) SetTransform(value Transform3D.BasisOrigin) {
	class(self).SetTransform(gd.Transform3D(value))
}

//go:nosplit
func (self class) GetTotalGravity() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_total_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTotalLinearDamp() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_total_linear_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTotalAngularDamp() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_total_angular_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetCenterOfMass() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_center_of_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetCenterOfMassLocal() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_center_of_mass_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetPrincipalInertiaAxes() gd.Basis {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Basis](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_principal_inertia_axes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetInverseMass() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_inverse_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetInverseInertia() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_inverse_inertia, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetInverseInertiaTensor() gd.Basis {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Basis](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_inverse_inertia_tensor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearVelocity(velocity gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_set_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearVelocity() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularVelocity(velocity gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_set_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularVelocity() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTransform(transform gd.Transform3D) {
	var frame = callframe.New()
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTransform() gd.Transform3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the body's velocity at the given relative position, including both translation and rotation.
*/
//go:nosplit
func (self class) GetVelocityAtLocalPosition(local_position gd.Vector3) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, local_position)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_velocity_at_local_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Applies a directional impulse without affecting rotation.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
This is equivalent to using [method apply_impulse] at the body's center of mass.
*/
//go:nosplit
func (self class) ApplyCentralImpulse(impulse gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, impulse)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_apply_central_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Applies a positioned impulse to the body.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) ApplyImpulse(impulse gd.Vector3, position gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, impulse)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_apply_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Applies a rotational impulse to the body without affecting the position.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[b]Note:[/b] [member inverse_inertia] is required for this to work. To have [member inverse_inertia], an active [CollisionShape3D] must be a child of the node, or you can manually set [member inverse_inertia].
*/
//go:nosplit
func (self class) ApplyTorqueImpulse(impulse gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, impulse)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_apply_torque_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Applies a directional force without affecting rotation. A force is time dependent and meant to be applied every physics update.
This is equivalent to using [method apply_force] at the body's center of mass.
*/
//go:nosplit
func (self class) ApplyCentralForce(force gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_apply_central_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Applies a positioned force to the body. A force is time dependent and meant to be applied every physics update.
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) ApplyForce(force gd.Vector3, position gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, force)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_apply_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Applies a rotational force without affecting position. A force is time dependent and meant to be applied every physics update.
[b]Note:[/b] [member inverse_inertia] is required for this to work. To have [member inverse_inertia], an active [CollisionShape3D] must be a child of the node, or you can manually set [member inverse_inertia].
*/
//go:nosplit
func (self class) ApplyTorque(torque gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_apply_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a constant directional force without affecting rotation that keeps being applied over time until cleared with [code]constant_force = Vector3(0, 0, 0)[/code].
This is equivalent to using [method add_constant_force] at the body's center of mass.
*/
//go:nosplit
func (self class) AddConstantCentralForce(force gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_add_constant_central_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a constant positioned force to the body that keeps being applied over time until cleared with [code]constant_force = Vector3(0, 0, 0)[/code].
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) AddConstantForce(force gd.Vector3, position gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, force)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_add_constant_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a constant rotational force without affecting position that keeps being applied over time until cleared with [code]constant_torque = Vector3(0, 0, 0)[/code].
*/
//go:nosplit
func (self class) AddConstantTorque(torque gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_add_constant_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the body's total constant positional forces applied during each physics update.
See [method add_constant_force] and [method add_constant_central_force].
*/
//go:nosplit
func (self class) SetConstantForce(force gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_set_constant_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the body's total constant positional forces applied during each physics update.
See [method add_constant_force] and [method add_constant_central_force].
*/
//go:nosplit
func (self class) GetConstantForce() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_constant_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the body's total constant rotational forces applied during each physics update.
See [method add_constant_torque].
*/
//go:nosplit
func (self class) SetConstantTorque(torque gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_set_constant_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the body's total constant rotational forces applied during each physics update.
See [method add_constant_torque].
*/
//go:nosplit
func (self class) GetConstantTorque() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_constant_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSleepState(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_set_sleep_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSleeping() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_is_sleeping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of contacts this body has with other bodies.
[b]Note:[/b] By default, this returns 0 unless bodies are configured to monitor contacts. See [member RigidBody3D.contact_monitor].
*/
//go:nosplit
func (self class) GetContactCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_contact_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the position of the contact point on the body in the global coordinate system.
*/
//go:nosplit
func (self class) GetContactLocalPosition(contact_idx gd.Int) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, contact_idx)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_contact_local_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the local normal at the contact point.
*/
//go:nosplit
func (self class) GetContactLocalNormal(contact_idx gd.Int) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, contact_idx)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_contact_local_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Impulse created by the contact.
*/
//go:nosplit
func (self class) GetContactImpulse(contact_idx gd.Int) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, contact_idx)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_contact_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the local shape index of the collision.
*/
//go:nosplit
func (self class) GetContactLocalShape(contact_idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, contact_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_contact_local_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the linear velocity vector at the body's contact point.
*/
//go:nosplit
func (self class) GetContactLocalVelocityAtPosition(contact_idx gd.Int) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, contact_idx)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_contact_local_velocity_at_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the collider's [RID].
*/
//go:nosplit
func (self class) GetContactCollider(contact_idx gd.Int) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, contact_idx)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_contact_collider, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the position of the contact point on the collider in the global coordinate system.
*/
//go:nosplit
func (self class) GetContactColliderPosition(contact_idx gd.Int) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, contact_idx)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_contact_collider_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the collider's object id.
*/
//go:nosplit
func (self class) GetContactColliderId(contact_idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, contact_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_contact_collider_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the collider object.
*/
//go:nosplit
func (self class) GetContactColliderObject(contact_idx gd.Int) gd.Object {
	var frame = callframe.New()
	callframe.Arg(frame, contact_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_contact_collider_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gd.PointerWithOwnershipTransferredToGo(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the collider's shape index.
*/
//go:nosplit
func (self class) GetContactColliderShape(contact_idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, contact_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_contact_collider_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the linear velocity vector at the collider's contact point.
*/
//go:nosplit
func (self class) GetContactColliderVelocityAtPosition(contact_idx gd.Int) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, contact_idx)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_contact_collider_velocity_at_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetStep() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Updates the body's linear and angular velocity by applying gravity and damping for the equivalent of one physics tick.
*/
//go:nosplit
func (self class) IntegrateForces() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_integrate_forces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the current state of the space, useful for queries.
*/
//go:nosplit
func (self class) GetSpaceState() objects.PhysicsDirectSpaceState3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectBodyState3D.Bind_get_space_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.PhysicsDirectSpaceState3D{classdb.PhysicsDirectSpaceState3D(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsPhysicsDirectBodyState3D() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPhysicsDirectBodyState3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("PhysicsDirectBodyState3D", func(ptr gd.Object) any { return classdb.PhysicsDirectBodyState3D(ptr) })
}
