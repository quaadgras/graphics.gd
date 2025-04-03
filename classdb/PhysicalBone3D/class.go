// Package PhysicalBone3D provides methods for working with PhysicalBone3D object instances.
package PhysicalBone3D

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/CollisionObject3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/PhysicsBody3D"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Vector3"

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
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
The [PhysicalBone3D] node is a physics body that can be used to make bones in a [Skeleton3D] react to physics.
[b]Note:[/b] In order to detect physical bones with raycasts, the [member SkeletonModifier3D.active] property of the parent [PhysicalBoneSimulator3D] must be [code]true[/code] and the [Skeleton3D]'s bone must be assigned to [PhysicalBone3D] correctly; it means that [method get_bone_id] should return a valid id ([code]>= 0[/code]).

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=PhysicalBone3D)
*/
type Instance [1]gdclass.PhysicalBone3D
type Expanded [1]gdclass.PhysicalBone3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPhysicalBone3D() Instance
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
Applies a directional impulse without affecting rotation.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_integrate_forces" functions otherwise).
This is equivalent to using [method apply_impulse] at the body's center of mass.
*/
func (self Instance) ApplyCentralImpulse(impulse Vector3.XYZ) { //gd:PhysicalBone3D.apply_central_impulse
	Advanced(self).ApplyCentralImpulse(Vector3.XYZ(impulse))
}

/*
Applies a positioned impulse to the PhysicsBone3D.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_integrate_forces" functions otherwise).
[param position] is the offset from the PhysicsBone3D origin in global coordinates.
*/
func (self Instance) ApplyImpulse(impulse Vector3.XYZ) { //gd:PhysicalBone3D.apply_impulse
	Advanced(self).ApplyImpulse(Vector3.XYZ(impulse), Vector3.XYZ(gd.Vector3{0, 0, 0}))
}

/*
Applies a positioned impulse to the PhysicsBone3D.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_integrate_forces" functions otherwise).
[param position] is the offset from the PhysicsBone3D origin in global coordinates.
*/
func (self Expanded) ApplyImpulse(impulse Vector3.XYZ, position Vector3.XYZ) { //gd:PhysicalBone3D.apply_impulse
	Advanced(self).ApplyImpulse(Vector3.XYZ(impulse), Vector3.XYZ(position))
}

/*
Returns [code]true[/code] if the PhysicsBone3D is allowed to simulate physics.
*/
func (self Instance) GetSimulatePhysics() bool { //gd:PhysicalBone3D.get_simulate_physics
	return bool(Advanced(self).GetSimulatePhysics())
}

/*
Returns [code]true[/code] if the PhysicsBone3D is currently simulating physics.
*/
func (self Instance) IsSimulatingPhysics() bool { //gd:PhysicalBone3D.is_simulating_physics
	return bool(Advanced(self).IsSimulatingPhysics())
}

/*
Returns the unique identifier of the PhysicsBone3D.
*/
func (self Instance) GetBoneId() int { //gd:PhysicalBone3D.get_bone_id
	return int(int(Advanced(self).GetBoneId()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicalBone3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicalBone3D"))
	casted := Instance{*(*gdclass.PhysicalBone3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) JointType() gdclass.PhysicalBone3DJointType {
	return gdclass.PhysicalBone3DJointType(class(self).GetJointType())
}

func (self Instance) SetJointType(value gdclass.PhysicalBone3DJointType) {
	class(self).SetJointType(value)
}

func (self Instance) JointOffset() Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetJointOffset())
}

func (self Instance) SetJointOffset(value Transform3D.BasisOrigin) {
	class(self).SetJointOffset(Transform3D.BasisOrigin(value))
}

func (self Instance) JointRotation() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetJointRotation())
}

func (self Instance) SetJointRotation(value Vector3.XYZ) {
	class(self).SetJointRotation(Vector3.XYZ(value))
}

func (self Instance) BodyOffset() Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetBodyOffset())
}

func (self Instance) SetBodyOffset(value Transform3D.BasisOrigin) {
	class(self).SetBodyOffset(Transform3D.BasisOrigin(value))
}

func (self Instance) Mass() Float.X {
	return Float.X(Float.X(class(self).GetMass()))
}

func (self Instance) SetMass(value Float.X) {
	class(self).SetMass(float64(value))
}

func (self Instance) Friction() Float.X {
	return Float.X(Float.X(class(self).GetFriction()))
}

func (self Instance) SetFriction(value Float.X) {
	class(self).SetFriction(float64(value))
}

func (self Instance) Bounce() Float.X {
	return Float.X(Float.X(class(self).GetBounce()))
}

func (self Instance) SetBounce(value Float.X) {
	class(self).SetBounce(float64(value))
}

func (self Instance) GravityScale() Float.X {
	return Float.X(Float.X(class(self).GetGravityScale()))
}

func (self Instance) SetGravityScale(value Float.X) {
	class(self).SetGravityScale(float64(value))
}

func (self Instance) CustomIntegrator() bool {
	return bool(class(self).IsUsingCustomIntegrator())
}

func (self Instance) SetCustomIntegrator(value bool) {
	class(self).SetUseCustomIntegrator(value)
}

func (self Instance) LinearDampMode() gdclass.PhysicalBone3DDampMode {
	return gdclass.PhysicalBone3DDampMode(class(self).GetLinearDampMode())
}

func (self Instance) SetLinearDampMode(value gdclass.PhysicalBone3DDampMode) {
	class(self).SetLinearDampMode(value)
}

func (self Instance) LinearDamp() Float.X {
	return Float.X(Float.X(class(self).GetLinearDamp()))
}

func (self Instance) SetLinearDamp(value Float.X) {
	class(self).SetLinearDamp(float64(value))
}

func (self Instance) AngularDampMode() gdclass.PhysicalBone3DDampMode {
	return gdclass.PhysicalBone3DDampMode(class(self).GetAngularDampMode())
}

func (self Instance) SetAngularDampMode(value gdclass.PhysicalBone3DDampMode) {
	class(self).SetAngularDampMode(value)
}

func (self Instance) AngularDamp() Float.X {
	return Float.X(Float.X(class(self).GetAngularDamp()))
}

func (self Instance) SetAngularDamp(value Float.X) {
	class(self).SetAngularDamp(float64(value))
}

func (self Instance) LinearVelocity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetLinearVelocity())
}

func (self Instance) SetLinearVelocity(value Vector3.XYZ) {
	class(self).SetLinearVelocity(Vector3.XYZ(value))
}

func (self Instance) AngularVelocity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetAngularVelocity())
}

func (self Instance) SetAngularVelocity(value Vector3.XYZ) {
	class(self).SetAngularVelocity(Vector3.XYZ(value))
}

func (self Instance) CanSleep() bool {
	return bool(class(self).IsAbleToSleep())
}

func (self Instance) SetCanSleep(value bool) {
	class(self).SetCanSleep(value)
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

/*
Applies a directional impulse without affecting rotation.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_integrate_forces" functions otherwise).
This is equivalent to using [method apply_impulse] at the body's center of mass.
*/
//go:nosplit
func (self class) ApplyCentralImpulse(impulse Vector3.XYZ) { //gd:PhysicalBone3D.apply_central_impulse
	var frame = callframe.New()
	callframe.Arg(frame, impulse)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_apply_central_impulse, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Applies a positioned impulse to the PhysicsBone3D.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_integrate_forces" functions otherwise).
[param position] is the offset from the PhysicsBone3D origin in global coordinates.
*/
//go:nosplit
func (self class) ApplyImpulse(impulse Vector3.XYZ, position Vector3.XYZ) { //gd:PhysicalBone3D.apply_impulse
	var frame = callframe.New()
	callframe.Arg(frame, impulse)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_apply_impulse, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetJointType(joint_type gdclass.PhysicalBone3DJointType) { //gd:PhysicalBone3D.set_joint_type
	var frame = callframe.New()
	callframe.Arg(frame, joint_type)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_joint_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetJointType() gdclass.PhysicalBone3DJointType { //gd:PhysicalBone3D.get_joint_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.PhysicalBone3DJointType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_joint_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJointOffset(offset Transform3D.BasisOrigin) { //gd:PhysicalBone3D.set_joint_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_joint_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetJointOffset() Transform3D.BasisOrigin { //gd:PhysicalBone3D.get_joint_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_joint_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJointRotation(euler Vector3.XYZ) { //gd:PhysicalBone3D.set_joint_rotation
	var frame = callframe.New()
	callframe.Arg(frame, euler)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_joint_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetJointRotation() Vector3.XYZ { //gd:PhysicalBone3D.get_joint_rotation
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_joint_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBodyOffset(offset Transform3D.BasisOrigin) { //gd:PhysicalBone3D.set_body_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_body_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBodyOffset() Transform3D.BasisOrigin { //gd:PhysicalBone3D.get_body_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_body_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the PhysicsBone3D is allowed to simulate physics.
*/
//go:nosplit
func (self class) GetSimulatePhysics() bool { //gd:PhysicalBone3D.get_simulate_physics
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_simulate_physics, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the PhysicsBone3D is currently simulating physics.
*/
//go:nosplit
func (self class) IsSimulatingPhysics() bool { //gd:PhysicalBone3D.is_simulating_physics
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_is_simulating_physics, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the unique identifier of the PhysicsBone3D.
*/
//go:nosplit
func (self class) GetBoneId() int64 { //gd:PhysicalBone3D.get_bone_id
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_bone_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMass(mass float64) { //gd:PhysicalBone3D.set_mass
	var frame = callframe.New()
	callframe.Arg(frame, mass)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_mass, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMass() float64 { //gd:PhysicalBone3D.get_mass
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_mass, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFriction(friction float64) { //gd:PhysicalBone3D.set_friction
	var frame = callframe.New()
	callframe.Arg(frame, friction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_friction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFriction() float64 { //gd:PhysicalBone3D.get_friction
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_friction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBounce(bounce float64) { //gd:PhysicalBone3D.set_bounce
	var frame = callframe.New()
	callframe.Arg(frame, bounce)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_bounce, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBounce() float64 { //gd:PhysicalBone3D.get_bounce
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_bounce, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravityScale(gravity_scale float64) { //gd:PhysicalBone3D.set_gravity_scale
	var frame = callframe.New()
	callframe.Arg(frame, gravity_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_gravity_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravityScale() float64 { //gd:PhysicalBone3D.get_gravity_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_gravity_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearDampMode(linear_damp_mode gdclass.PhysicalBone3DDampMode) { //gd:PhysicalBone3D.set_linear_damp_mode
	var frame = callframe.New()
	callframe.Arg(frame, linear_damp_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_linear_damp_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearDampMode() gdclass.PhysicalBone3DDampMode { //gd:PhysicalBone3D.get_linear_damp_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.PhysicalBone3DDampMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_linear_damp_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularDampMode(angular_damp_mode gdclass.PhysicalBone3DDampMode) { //gd:PhysicalBone3D.set_angular_damp_mode
	var frame = callframe.New()
	callframe.Arg(frame, angular_damp_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_angular_damp_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularDampMode() gdclass.PhysicalBone3DDampMode { //gd:PhysicalBone3D.get_angular_damp_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.PhysicalBone3DDampMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_angular_damp_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearDamp(linear_damp float64) { //gd:PhysicalBone3D.set_linear_damp
	var frame = callframe.New()
	callframe.Arg(frame, linear_damp)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_linear_damp, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearDamp() float64 { //gd:PhysicalBone3D.get_linear_damp
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_linear_damp, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularDamp(angular_damp float64) { //gd:PhysicalBone3D.set_angular_damp
	var frame = callframe.New()
	callframe.Arg(frame, angular_damp)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_angular_damp, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularDamp() float64 { //gd:PhysicalBone3D.get_angular_damp
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_angular_damp, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearVelocity(linear_velocity Vector3.XYZ) { //gd:PhysicalBone3D.set_linear_velocity
	var frame = callframe.New()
	callframe.Arg(frame, linear_velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearVelocity() Vector3.XYZ { //gd:PhysicalBone3D.get_linear_velocity
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularVelocity(angular_velocity Vector3.XYZ) { //gd:PhysicalBone3D.set_angular_velocity
	var frame = callframe.New()
	callframe.Arg(frame, angular_velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularVelocity() Vector3.XYZ { //gd:PhysicalBone3D.get_angular_velocity
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseCustomIntegrator(enable bool) { //gd:PhysicalBone3D.set_use_custom_integrator
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_use_custom_integrator, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingCustomIntegrator() bool { //gd:PhysicalBone3D.is_using_custom_integrator
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_is_using_custom_integrator, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCanSleep(able_to_sleep bool) { //gd:PhysicalBone3D.set_can_sleep
	var frame = callframe.New()
	callframe.Arg(frame, able_to_sleep)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_can_sleep, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAbleToSleep() bool { //gd:PhysicalBone3D.is_able_to_sleep
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_is_able_to_sleep, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPhysicalBone3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPhysicalBone3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("PhysicalBone3D", func(ptr gd.Object) any {
		return [1]gdclass.PhysicalBone3D{*(*gdclass.PhysicalBone3D)(unsafe.Pointer(&ptr))}
	})
}

type DampMode = gdclass.PhysicalBone3DDampMode //gd:PhysicalBone3D.DampMode

const (
	/*In this mode, the body's damping value is added to any value set in areas or the default value.*/
	DampModeCombine DampMode = 0
	/*In this mode, the body's damping value replaces any value set in areas or the default value.*/
	DampModeReplace DampMode = 1
)

type JointType = gdclass.PhysicalBone3DJointType //gd:PhysicalBone3D.JointType

const (
	/*No joint is applied to the PhysicsBone3D.*/
	JointTypeNone JointType = 0
	/*A pin joint is applied to the PhysicsBone3D.*/
	JointTypePin JointType = 1
	/*A cone joint is applied to the PhysicsBone3D.*/
	JointTypeCone JointType = 2
	/*A hinge joint is applied to the PhysicsBone3D.*/
	JointTypeHinge JointType = 3
	/*A slider joint is applied to the PhysicsBone3D.*/
	JointTypeSlider JointType = 4
	/*A 6 degrees of freedom joint is applied to the PhysicsBone3D.*/
	JointType6dof JointType = 5
)
