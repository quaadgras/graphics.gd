package PhysicalBone3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/PhysicsBody3D"
import "graphics.gd/classdb/CollisionObject3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The [PhysicalBone3D] node is a physics body that can be used to make bones in a [Skeleton3D] react to physics.
[b]Note:[/b] In order to detect physical bones with raycasts, the [member SkeletonModifier3D.active] property of the parent [PhysicalBoneSimulator3D] must be [code]true[/code] and the [Skeleton3D]'s bone must be assigned to [PhysicalBone3D] correctly; it means that [method get_bone_id] should return a valid id ([code]>= 0[/code]).

	// PhysicalBone3D methods that can be overridden by a [Class] that extends it.
	type PhysicalBone3D interface {
		//Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
		IntegrateForces(state [1]gdclass.PhysicsDirectBodyState3D)
	}
*/
type Instance [1]gdclass.PhysicalBone3D
type Any interface {
	gd.IsClass
	AsPhysicalBone3D() Instance
}

/*
Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
*/
func (Instance) _integrate_forces(impl func(ptr unsafe.Pointer, state [1]gdclass.PhysicsDirectBodyState3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.PhysicsDirectBodyState3D{pointers.New[gdclass.PhysicsDirectBodyState3D]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, state)
	}
}
func (self Instance) ApplyCentralImpulse(impulse Vector3.XYZ) {
	class(self).ApplyCentralImpulse(gd.Vector3(impulse))
}
func (self Instance) ApplyImpulse(impulse Vector3.XYZ) {
	class(self).ApplyImpulse(gd.Vector3(impulse), gd.Vector3(gd.Vector3{0, 0, 0}))
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
type class [1]gdclass.PhysicalBone3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicalBone3D"))
	return Instance{*(*gdclass.PhysicalBone3D)(unsafe.Pointer(&object))}
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
	class(self).SetJointOffset(gd.Transform3D(value))
}

func (self Instance) JointRotation() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetJointRotation())
}

func (self Instance) SetJointRotation(value Vector3.XYZ) {
	class(self).SetJointRotation(gd.Vector3(value))
}

func (self Instance) BodyOffset() Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetBodyOffset())
}

func (self Instance) SetBodyOffset(value Transform3D.BasisOrigin) {
	class(self).SetBodyOffset(gd.Transform3D(value))
}

func (self Instance) Mass() Float.X {
	return Float.X(Float.X(class(self).GetMass()))
}

func (self Instance) SetMass(value Float.X) {
	class(self).SetMass(gd.Float(value))
}

func (self Instance) Friction() Float.X {
	return Float.X(Float.X(class(self).GetFriction()))
}

func (self Instance) SetFriction(value Float.X) {
	class(self).SetFriction(gd.Float(value))
}

func (self Instance) Bounce() Float.X {
	return Float.X(Float.X(class(self).GetBounce()))
}

func (self Instance) SetBounce(value Float.X) {
	class(self).SetBounce(gd.Float(value))
}

func (self Instance) GravityScale() Float.X {
	return Float.X(Float.X(class(self).GetGravityScale()))
}

func (self Instance) SetGravityScale(value Float.X) {
	class(self).SetGravityScale(gd.Float(value))
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
	class(self).SetLinearDamp(gd.Float(value))
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
	class(self).SetAngularDamp(gd.Float(value))
}

func (self Instance) LinearVelocity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetLinearVelocity())
}

func (self Instance) SetLinearVelocity(value Vector3.XYZ) {
	class(self).SetLinearVelocity(gd.Vector3(value))
}

func (self Instance) AngularVelocity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetAngularVelocity())
}

func (self Instance) SetAngularVelocity(value Vector3.XYZ) {
	class(self).SetAngularVelocity(gd.Vector3(value))
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
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.PhysicsDirectBodyState3D{pointers.New[gdclass.PhysicsDirectBodyState3D]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
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
func (self class) SetJointType(joint_type gdclass.PhysicalBone3DJointType) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_joint_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetJointType() gdclass.PhysicalBone3DJointType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.PhysicalBone3DJointType](frame)
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
func (self class) SetLinearDampMode(linear_damp_mode gdclass.PhysicalBone3DDampMode) {
	var frame = callframe.New()
	callframe.Arg(frame, linear_damp_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_linear_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearDampMode() gdclass.PhysicalBone3DDampMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.PhysicalBone3DDampMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_get_linear_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularDampMode(angular_damp_mode gdclass.PhysicalBone3DDampMode) {
	var frame = callframe.New()
	callframe.Arg(frame, angular_damp_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone3D.Bind_set_angular_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularDampMode() gdclass.PhysicalBone3DDampMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.PhysicalBone3DDampMode](frame)
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
	gdclass.Register("PhysicalBone3D", func(ptr gd.Object) any {
		return [1]gdclass.PhysicalBone3D{*(*gdclass.PhysicalBone3D)(unsafe.Pointer(&ptr))}
	})
}

type DampMode = gdclass.PhysicalBone3DDampMode

const (
	/*In this mode, the body's damping value is added to any value set in areas or the default value.*/
	DampModeCombine DampMode = 0
	/*In this mode, the body's damping value replaces any value set in areas or the default value.*/
	DampModeReplace DampMode = 1
)

type JointType = gdclass.PhysicalBone3DJointType

const (
	JointTypeNone   JointType = 0
	JointTypePin    JointType = 1
	JointTypeCone   JointType = 2
	JointTypeHinge  JointType = 3
	JointTypeSlider JointType = 4
	JointType6dof   JointType = 5
)
