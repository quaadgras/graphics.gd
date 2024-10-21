package PhysicalBone3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/PhysicsBody3D"
import "grow.graphics/gd/object/CollisionObject3D"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The [PhysicalBone3D] node is a physics body that can be used to make bones in a [Skeleton3D] react to physics.
[b]Note:[/b] In order to detect physical bones with raycasts, the [member SkeletonModifier3D.active] property of the parent [PhysicalBoneSimulator3D] must be [code]true[/code] and the [Skeleton3D]'s bone must be assigned to [PhysicalBone3D] correctly; it means that [method get_bone_id] should return a valid id ([code]>= 0[/code]).
	// PhysicalBone3D methods that can be overridden by a [Class] that extends it.
	type PhysicalBone3D interface {
		//Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
		IntegrateForces(state object.PhysicsDirectBodyState3D) 
	}

*/
type Simple [1]classdb.PhysicalBone3D
func (Simple) _integrate_forces(impl func(ptr unsafe.Pointer, state [1]classdb.PhysicsDirectBodyState3D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var state [1]classdb.PhysicsDirectBodyState3D
		state[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, state)
		gc.End()
	}
}
func (self Simple) ApplyCentralImpulse(impulse gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ApplyCentralImpulse(impulse)
}
func (self Simple) ApplyImpulse(impulse gd.Vector3, position gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ApplyImpulse(impulse, position)
}
func (self Simple) SetJointType(joint_type classdb.PhysicalBone3DJointType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetJointType(joint_type)
}
func (self Simple) GetJointType() classdb.PhysicalBone3DJointType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.PhysicalBone3DJointType(Expert(self).GetJointType())
}
func (self Simple) SetJointOffset(offset gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetJointOffset(offset)
}
func (self Simple) GetJointOffset() gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).GetJointOffset())
}
func (self Simple) SetJointRotation(euler gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetJointRotation(euler)
}
func (self Simple) GetJointRotation() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetJointRotation())
}
func (self Simple) SetBodyOffset(offset gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBodyOffset(offset)
}
func (self Simple) GetBodyOffset() gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).GetBodyOffset())
}
func (self Simple) GetSimulatePhysics() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetSimulatePhysics())
}
func (self Simple) IsSimulatingPhysics() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSimulatingPhysics())
}
func (self Simple) GetBoneId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBoneId()))
}
func (self Simple) SetMass(mass float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMass(gd.Float(mass))
}
func (self Simple) GetMass() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMass()))
}
func (self Simple) SetFriction(friction float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFriction(gd.Float(friction))
}
func (self Simple) GetFriction() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFriction()))
}
func (self Simple) SetBounce(bounce float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBounce(gd.Float(bounce))
}
func (self Simple) GetBounce() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetBounce()))
}
func (self Simple) SetGravityScale(gravity_scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGravityScale(gd.Float(gravity_scale))
}
func (self Simple) GetGravityScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetGravityScale()))
}
func (self Simple) SetLinearDampMode(linear_damp_mode classdb.PhysicalBone3DDampMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLinearDampMode(linear_damp_mode)
}
func (self Simple) GetLinearDampMode() classdb.PhysicalBone3DDampMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.PhysicalBone3DDampMode(Expert(self).GetLinearDampMode())
}
func (self Simple) SetAngularDampMode(angular_damp_mode classdb.PhysicalBone3DDampMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAngularDampMode(angular_damp_mode)
}
func (self Simple) GetAngularDampMode() classdb.PhysicalBone3DDampMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.PhysicalBone3DDampMode(Expert(self).GetAngularDampMode())
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
func (self Simple) SetLinearVelocity(linear_velocity gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLinearVelocity(linear_velocity)
}
func (self Simple) GetLinearVelocity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetLinearVelocity())
}
func (self Simple) SetAngularVelocity(angular_velocity gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAngularVelocity(angular_velocity)
}
func (self Simple) GetAngularVelocity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetAngularVelocity())
}
func (self Simple) SetUseCustomIntegrator(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseCustomIntegrator(enable)
}
func (self Simple) IsUsingCustomIntegrator() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUsingCustomIntegrator())
}
func (self Simple) SetCanSleep(able_to_sleep bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCanSleep(able_to_sleep)
}
func (self Simple) IsAbleToSleep() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAbleToSleep())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PhysicalBone3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
*/
func (class) _integrate_forces(impl func(ptr unsafe.Pointer, state object.PhysicsDirectBodyState3D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var state object.PhysicsDirectBodyState3D
		state[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, state)
		ctx.End()
	}
}

//go:nosplit
func (self class) ApplyCentralImpulse(impulse gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, impulse)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_apply_central_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) ApplyImpulse(impulse gd.Vector3, position gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, impulse)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_apply_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetJointType(joint_type classdb.PhysicalBone3DJointType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_set_joint_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetJointType() classdb.PhysicalBone3DJointType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.PhysicalBone3DJointType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_get_joint_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetJointOffset(offset gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_set_joint_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetJointOffset() gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_get_joint_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetJointRotation(euler gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, euler)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_set_joint_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetJointRotation() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_get_joint_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBodyOffset(offset gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_set_body_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBodyOffset() gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_get_body_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetSimulatePhysics() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_get_simulate_physics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) IsSimulatingPhysics() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_is_simulating_physics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetBoneId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_get_bone_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMass(mass gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mass)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_set_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMass() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_get_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFriction(friction gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, friction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_set_friction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFriction() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_get_friction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBounce(bounce gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bounce)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_set_bounce, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBounce() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_get_bounce, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGravityScale(gravity_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gravity_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_set_gravity_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGravityScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_get_gravity_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLinearDampMode(linear_damp_mode classdb.PhysicalBone3DDampMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, linear_damp_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_set_linear_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLinearDampMode() classdb.PhysicalBone3DDampMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.PhysicalBone3DDampMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_get_linear_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAngularDampMode(angular_damp_mode classdb.PhysicalBone3DDampMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angular_damp_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_set_angular_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAngularDampMode() classdb.PhysicalBone3DDampMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.PhysicalBone3DDampMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_get_angular_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_set_linear_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLinearDamp() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_get_linear_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_set_angular_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAngularDamp() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_get_angular_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLinearVelocity(linear_velocity gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, linear_velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_set_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLinearVelocity() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_get_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAngularVelocity(angular_velocity gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angular_velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_set_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAngularVelocity() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_get_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_set_use_custom_integrator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingCustomIntegrator() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_is_using_custom_integrator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_set_can_sleep, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAbleToSleep() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicalBone3D.Bind_is_able_to_sleep, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsPhysicalBone3D() Expert { return self[0].AsPhysicalBone3D() }


//go:nosplit
func (self Simple) AsPhysicalBone3D() Simple { return self[0].AsPhysicalBone3D() }


//go:nosplit
func (self class) AsPhysicsBody3D() PhysicsBody3D.Expert { return self[0].AsPhysicsBody3D() }


//go:nosplit
func (self Simple) AsPhysicsBody3D() PhysicsBody3D.Simple { return self[0].AsPhysicsBody3D() }


//go:nosplit
func (self class) AsCollisionObject3D() CollisionObject3D.Expert { return self[0].AsCollisionObject3D() }


//go:nosplit
func (self Simple) AsCollisionObject3D() CollisionObject3D.Simple { return self[0].AsCollisionObject3D() }


//go:nosplit
func (self class) AsNode3D() Node3D.Expert { return self[0].AsNode3D() }


//go:nosplit
func (self Simple) AsNode3D() Node3D.Simple { return self[0].AsNode3D() }


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
func init() {classdb.Register("PhysicalBone3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type DampMode = classdb.PhysicalBone3DDampMode

const (
/*In this mode, the body's damping value is added to any value set in areas or the default value.*/
	DampModeCombine DampMode = 0
/*In this mode, the body's damping value replaces any value set in areas or the default value.*/
	DampModeReplace DampMode = 1
)
type JointType = classdb.PhysicalBone3DJointType

const (
	JointTypeNone JointType = 0
	JointTypePin JointType = 1
	JointTypeCone JointType = 2
	JointTypeHinge JointType = 3
	JointTypeSlider JointType = 4
	JointType6dof JointType = 5
)
