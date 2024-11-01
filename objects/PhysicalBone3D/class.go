package PhysicalBone3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/PhysicsBody3D"
import "grow.graphics/gd/objects/CollisionObject3D"
import "grow.graphics/gd/objects/Node3D"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
The [PhysicalBone3D] node is a physics body that can be used to make bones in a [Skeleton3D] react to physics.
[b]Note:[/b] In order to detect physical bones with raycasts, the [member SkeletonModifier3D.active] property of the parent [PhysicalBoneSimulator3D] must be [code]true[/code] and the [Skeleton3D]'s bone must be assigned to [PhysicalBone3D] correctly; it means that [method get_bone_id] should return a valid id ([code]>= 0[/code]).

	// PhysicalBone3D methods that can be overridden by a [Class] that extends it.
	type PhysicalBone3D interface {
		//Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
		IntegrateForces(state objects.PhysicsDirectBodyState3D)
	}
*/
type Instance [1]classdb.PhysicalBone3D

/*
Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
*/
func (Instance) _integrate_forces(impl func(ptr unsafe.Pointer, state objects.PhysicsDirectBodyState3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = objects.PhysicsDirectBodyState3D{pointers.New[classdb.PhysicsDirectBodyState3D]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, state)
	}
}
func (self Instance) ApplyCentralImpulse(impulse gd.Vector3) {
	class(self).ApplyCentralImpulse(impulse)
}
func (self Instance) ApplyImpulse(impulse gd.Vector3) {
	class(self).ApplyImpulse(impulse, gd.Vector3{0, 0, 0})
}
func (self Instance) GetSimulatePhysics() bool {
	return bool(class(self).GetSimulatePhysics())
}
func (self Instance) IsSimulatingPhysics() bool {
	return bool(class(self).IsSimulatingPhysics())
}
func (self Instance) GetBoneId() int {
	return int(int(class(self).GetBoneId()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.PhysicalBone3D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicalBone3D"))
	return Instance{classdb.PhysicalBone3D(object)}
}

func (self Instance) JointType() classdb.PhysicalBone3DJointType {
	return classdb.PhysicalBone3DJointType(class(self).GetJointType())
}

func (self Instance) SetJointType(value classdb.PhysicalBone3DJointType) {
	class(self).SetJointType(value)
}

func (self Instance) JointOffset() gd.Transform3D {
	return gd.Transform3D(class(self).GetJointOffset())
}

func (self Instance) SetJointOffset(value gd.Transform3D) {
	class(self).SetJointOffset(value)
}

func (self Instance) JointRotation() gd.Vector3 {
	return gd.Vector3(class(self).GetJointRotation())
}

func (self Instance) SetJointRotation(value gd.Vector3) {
	class(self).SetJointRotation(value)
}

func (self Instance) BodyOffset() gd.Transform3D {
	return gd.Transform3D(class(self).GetBodyOffset())
}

func (self Instance) SetBodyOffset(value gd.Transform3D) {
	class(self).SetBodyOffset(value)
}

func (self Instance) Mass() float64 {
	return float64(float64(class(self).GetMass()))
}

func (self Instance) SetMass(value float64) {
	class(self).SetMass(gd.Float(value))
}

func (self Instance) Friction() float64 {
	return float64(float64(class(self).GetFriction()))
}

func (self Instance) SetFriction(value float64) {
	class(self).SetFriction(gd.Float(value))
}

func (self Instance) Bounce() float64 {
	return float64(float64(class(self).GetBounce()))
}

func (self Instance) SetBounce(value float64) {
	class(self).SetBounce(gd.Float(value))
}

func (self Instance) GravityScale() float64 {
	return float64(float64(class(self).GetGravityScale()))
}

func (self Instance) SetGravityScale(value float64) {
	class(self).SetGravityScale(gd.Float(value))
}

func (self Instance) CustomIntegrator() bool {
	return bool(class(self).IsUsingCustomIntegrator())
}

func (self Instance) SetCustomIntegrator(value bool) {
	class(self).SetUseCustomIntegrator(value)
}

func (self Instance) LinearDampMode() classdb.PhysicalBone3DDampMode {
	return classdb.PhysicalBone3DDampMode(class(self).GetLinearDampMode())
}

func (self Instance) SetLinearDampMode(value classdb.PhysicalBone3DDampMode) {
	class(self).SetLinearDampMode(value)
}

func (self Instance) LinearDamp() float64 {
	return float64(float64(class(self).GetLinearDamp()))
}

func (self Instance) SetLinearDamp(value float64) {
	class(self).SetLinearDamp(gd.Float(value))
}

func (self Instance) AngularDampMode() classdb.PhysicalBone3DDampMode {
	return classdb.PhysicalBone3DDampMode(class(self).GetAngularDampMode())
}

func (self Instance) SetAngularDampMode(value classdb.PhysicalBone3DDampMode) {
	class(self).SetAngularDampMode(value)
}

func (self Instance) AngularDamp() float64 {
	return float64(float64(class(self).GetAngularDamp()))
}

func (self Instance) SetAngularDamp(value float64) {
	class(self).SetAngularDamp(gd.Float(value))
}

func (self Instance) LinearVelocity() gd.Vector3 {
	return gd.Vector3(class(self).GetLinearVelocity())
}

func (self Instance) SetLinearVelocity(value gd.Vector3) {
	class(self).SetLinearVelocity(value)
}

func (self Instance) AngularVelocity() gd.Vector3 {
	return gd.Vector3(class(self).GetAngularVelocity())
}

func (self Instance) SetAngularVelocity(value gd.Vector3) {
	class(self).SetAngularVelocity(value)
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
func (class) _integrate_forces(impl func(ptr unsafe.Pointer, state objects.PhysicsDirectBodyState3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = objects.PhysicsDirectBodyState3D{pointers.New[classdb.PhysicsDirectBodyState3D]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, state)
	}
}

//go:nosplit
func (self class) ApplyCentralImpulse(impulse gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, impulse)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_apply_central_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) ApplyImpulse(impulse gd.Vector3, position gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, impulse)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_apply_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetJointType(joint_type classdb.PhysicalBone3DJointType) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_joint_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetJointType() classdb.PhysicalBone3DJointType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.PhysicalBone3DJointType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_joint_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJointOffset(offset gd.Transform3D) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_joint_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetJointOffset() gd.Transform3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_joint_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJointRotation(euler gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, euler)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_joint_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetJointRotation() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_joint_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBodyOffset(offset gd.Transform3D) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_body_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBodyOffset() gd.Transform3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_body_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetSimulatePhysics() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_simulate_physics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) IsSimulatingPhysics() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_is_simulating_physics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetBoneId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_bone_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMass(mass gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, mass)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMass() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFriction(friction gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, friction)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_friction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFriction() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_friction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBounce(bounce gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, bounce)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_bounce, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBounce() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_bounce, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravityScale(gravity_scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, gravity_scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_gravity_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravityScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_gravity_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearDampMode(linear_damp_mode classdb.PhysicalBone3DDampMode) {
	var frame = callframe.New()
	callframe.Arg(frame, linear_damp_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_linear_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearDampMode() classdb.PhysicalBone3DDampMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.PhysicalBone3DDampMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_linear_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularDampMode(angular_damp_mode classdb.PhysicalBone3DDampMode) {
	var frame = callframe.New()
	callframe.Arg(frame, angular_damp_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_angular_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularDampMode() classdb.PhysicalBone3DDampMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.PhysicalBone3DDampMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_angular_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearDamp(linear_damp gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, linear_damp)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_linear_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearDamp() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_linear_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularDamp(angular_damp gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, angular_damp)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_angular_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularDamp() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_angular_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearVelocity(linear_velocity gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, linear_velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearVelocity() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularVelocity(angular_velocity gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, angular_velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularVelocity() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseCustomIntegrator(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_use_custom_integrator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingCustomIntegrator() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_is_using_custom_integrator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCanSleep(able_to_sleep bool) {
	var frame = callframe.New()
	callframe.Arg(frame, able_to_sleep)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_can_sleep, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsAbleToSleep() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_is_able_to_sleep, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
		return gd.VirtualByName(self.AsPhysicsBody3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_integrate_forces":
		return reflect.ValueOf(self._integrate_forces)
	default:
		return gd.VirtualByName(self.AsPhysicsBody3D(), name)
	}
}
func init() {
	classdb.Register("PhysicalBone3D", func(ptr gd.Object) any { return classdb.PhysicalBone3D(ptr) })
}

type DampMode = classdb.PhysicalBone3DDampMode

const (
	/*In this mode, the body's damping value is added to any value set in areas or the default value.*/
	DampModeCombine DampMode = 0
	/*In this mode, the body's damping value replaces any value set in areas or the default value.*/
	DampModeReplace DampMode = 1
)

type JointType = classdb.PhysicalBone3DJointType

const (
	JointTypeNone   JointType = 0
	JointTypePin    JointType = 1
	JointTypeCone   JointType = 2
	JointTypeHinge  JointType = 3
	JointTypeSlider JointType = 4
	JointType6dof   JointType = 5
)
