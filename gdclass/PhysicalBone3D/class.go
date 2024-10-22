package PhysicalBone3D

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
The [PhysicalBone3D] node is a physics body that can be used to make bones in a [Skeleton3D] react to physics.
[b]Note:[/b] In order to detect physical bones with raycasts, the [member SkeletonModifier3D.active] property of the parent [PhysicalBoneSimulator3D] must be [code]true[/code] and the [Skeleton3D]'s bone must be assigned to [PhysicalBone3D] correctly; it means that [method get_bone_id] should return a valid id ([code]>= 0[/code]).
	// PhysicalBone3D methods that can be overridden by a [Class] that extends it.
	type PhysicalBone3D interface {
		//Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
		IntegrateForces(state gdclass.PhysicsDirectBodyState3D) 
	}

*/
type Go [1]classdb.PhysicalBone3D

/*
Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
*/
func (Go) _integrate_forces(impl func(ptr unsafe.Pointer, state gdclass.PhysicsDirectBodyState3D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var state gdclass.PhysicsDirectBodyState3D
		state[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, state)
		gc.End()
	}
}
func (self Go) ApplyCentralImpulse(impulse gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ApplyCentralImpulse(impulse)
}
func (self Go) ApplyImpulse(impulse gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ApplyImpulse(impulse, gd.Vector3{0, 0, 0})
}
func (self Go) GetSimulatePhysics() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetSimulatePhysics())
}
func (self Go) IsSimulatingPhysics() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsSimulatingPhysics())
}
func (self Go) GetBoneId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetBoneId()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PhysicalBone3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PhysicalBone3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) JointType() classdb.PhysicalBone3DJointType {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.PhysicalBone3DJointType(class(self).GetJointType())
}

func (self Go) SetJointType(value classdb.PhysicalBone3DJointType) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJointType(value)
}

func (self Go) JointOffset() gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Transform3D(class(self).GetJointOffset())
}

func (self Go) SetJointOffset(value gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJointOffset(value)
}

func (self Go) JointRotation() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetJointRotation())
}

func (self Go) SetJointRotation(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJointRotation(value)
}

func (self Go) BodyOffset() gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Transform3D(class(self).GetBodyOffset())
}

func (self Go) SetBodyOffset(value gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBodyOffset(value)
}

func (self Go) Mass() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMass()))
}

func (self Go) SetMass(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMass(gd.Float(value))
}

func (self Go) Friction() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFriction()))
}

func (self Go) SetFriction(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFriction(gd.Float(value))
}

func (self Go) Bounce() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetBounce()))
}

func (self Go) SetBounce(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBounce(gd.Float(value))
}

func (self Go) GravityScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetGravityScale()))
}

func (self Go) SetGravityScale(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGravityScale(gd.Float(value))
}

func (self Go) CustomIntegrator() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsUsingCustomIntegrator())
}

func (self Go) SetCustomIntegrator(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseCustomIntegrator(value)
}

func (self Go) LinearDampMode() classdb.PhysicalBone3DDampMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.PhysicalBone3DDampMode(class(self).GetLinearDampMode())
}

func (self Go) SetLinearDampMode(value classdb.PhysicalBone3DDampMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLinearDampMode(value)
}

func (self Go) LinearDamp() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetLinearDamp()))
}

func (self Go) SetLinearDamp(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLinearDamp(gd.Float(value))
}

func (self Go) AngularDampMode() classdb.PhysicalBone3DDampMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.PhysicalBone3DDampMode(class(self).GetAngularDampMode())
}

func (self Go) SetAngularDampMode(value classdb.PhysicalBone3DDampMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAngularDampMode(value)
}

func (self Go) AngularDamp() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAngularDamp()))
}

func (self Go) SetAngularDamp(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAngularDamp(gd.Float(value))
}

func (self Go) LinearVelocity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetLinearVelocity())
}

func (self Go) SetLinearVelocity(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLinearVelocity(value)
}

func (self Go) AngularVelocity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetAngularVelocity())
}

func (self Go) SetAngularVelocity(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAngularVelocity(value)
}

func (self Go) CanSleep() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsAbleToSleep())
}

func (self Go) SetCanSleep(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCanSleep(value)
}

/*
Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
*/
func (class) _integrate_forces(impl func(ptr unsafe.Pointer, state gdclass.PhysicsDirectBodyState3D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var state gdclass.PhysicsDirectBodyState3D
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
func (self class) AsPhysicalBone3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPhysicalBone3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
	case "_integrate_forces": return reflect.ValueOf(self._integrate_forces);
	default: return gd.VirtualByName(self.AsPhysicsBody3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_integrate_forces": return reflect.ValueOf(self._integrate_forces);
	default: return gd.VirtualByName(self.AsPhysicsBody3D(), name)
	}
}
func init() {classdb.Register("PhysicalBone3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
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
