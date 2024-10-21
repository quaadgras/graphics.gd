package SkeletonModification2DJiggle

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/SkeletonModification2D"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This modification moves a series of bones, typically called a bone chain, towards a target. What makes this modification special is that it calculates the velocity and acceleration for each bone in the bone chain, and runs a very light physics-like calculation using the inputted values. This allows the bones to overshoot the target and "jiggle" around. It can be configured to act more like a spring, or sway around like cloth might.
This modification is useful for adding additional motion to things like hair, the edges of clothing, and more. It has several settings to that allow control over how the joint moves when the target moves.
[b]Note:[/b] The Jiggle modifier has [code]jiggle_joints[/code], which are the data objects that hold the data for each joint in the Jiggle chain. This is different from than [Bone2D] nodes! Jiggle joints hold the data needed for each [Bone2D] in the bone chain used by the Jiggle modification.

*/
type Simple [1]classdb.SkeletonModification2DJiggle
func (self Simple) SetTargetNode(target_nodepath gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTargetNode(target_nodepath)
}
func (self Simple) GetTargetNode() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetTargetNode(gc))
}
func (self Simple) SetJiggleDataChainLength(length int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetJiggleDataChainLength(gd.Int(length))
}
func (self Simple) GetJiggleDataChainLength() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetJiggleDataChainLength()))
}
func (self Simple) SetStiffness(stiffness float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStiffness(gd.Float(stiffness))
}
func (self Simple) GetStiffness() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetStiffness()))
}
func (self Simple) SetMass(mass float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMass(gd.Float(mass))
}
func (self Simple) GetMass() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMass()))
}
func (self Simple) SetDamping(damping float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDamping(gd.Float(damping))
}
func (self Simple) GetDamping() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDamping()))
}
func (self Simple) SetUseGravity(use_gravity bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseGravity(use_gravity)
}
func (self Simple) GetUseGravity() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetUseGravity())
}
func (self Simple) SetGravity(gravity gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGravity(gravity)
}
func (self Simple) GetGravity() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetGravity())
}
func (self Simple) SetUseColliders(use_colliders bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseColliders(use_colliders)
}
func (self Simple) GetUseColliders() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetUseColliders())
}
func (self Simple) SetCollisionMask(collision_mask int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCollisionMask(gd.Int(collision_mask))
}
func (self Simple) GetCollisionMask() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCollisionMask()))
}
func (self Simple) SetJiggleJointBone2dNode(joint_idx int, bone2d_node gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetJiggleJointBone2dNode(gd.Int(joint_idx), bone2d_node)
}
func (self Simple) GetJiggleJointBone2dNode(joint_idx int) gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetJiggleJointBone2dNode(gc, gd.Int(joint_idx)))
}
func (self Simple) SetJiggleJointBoneIndex(joint_idx int, bone_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetJiggleJointBoneIndex(gd.Int(joint_idx), gd.Int(bone_idx))
}
func (self Simple) GetJiggleJointBoneIndex(joint_idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetJiggleJointBoneIndex(gd.Int(joint_idx))))
}
func (self Simple) SetJiggleJointOverride(joint_idx int, override bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetJiggleJointOverride(gd.Int(joint_idx), override)
}
func (self Simple) GetJiggleJointOverride(joint_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetJiggleJointOverride(gd.Int(joint_idx)))
}
func (self Simple) SetJiggleJointStiffness(joint_idx int, stiffness float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetJiggleJointStiffness(gd.Int(joint_idx), gd.Float(stiffness))
}
func (self Simple) GetJiggleJointStiffness(joint_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetJiggleJointStiffness(gd.Int(joint_idx))))
}
func (self Simple) SetJiggleJointMass(joint_idx int, mass float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetJiggleJointMass(gd.Int(joint_idx), gd.Float(mass))
}
func (self Simple) GetJiggleJointMass(joint_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetJiggleJointMass(gd.Int(joint_idx))))
}
func (self Simple) SetJiggleJointDamping(joint_idx int, damping float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetJiggleJointDamping(gd.Int(joint_idx), gd.Float(damping))
}
func (self Simple) GetJiggleJointDamping(joint_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetJiggleJointDamping(gd.Int(joint_idx))))
}
func (self Simple) SetJiggleJointUseGravity(joint_idx int, use_gravity bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetJiggleJointUseGravity(gd.Int(joint_idx), use_gravity)
}
func (self Simple) GetJiggleJointUseGravity(joint_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetJiggleJointUseGravity(gd.Int(joint_idx)))
}
func (self Simple) SetJiggleJointGravity(joint_idx int, gravity gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetJiggleJointGravity(gd.Int(joint_idx), gravity)
}
func (self Simple) GetJiggleJointGravity(joint_idx int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetJiggleJointGravity(gd.Int(joint_idx)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.SkeletonModification2DJiggle
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetTargetNode(target_nodepath gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(target_nodepath))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTargetNode(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetJiggleDataChainLength(length gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_jiggle_data_chain_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetJiggleDataChainLength() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_jiggle_data_chain_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStiffness(stiffness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stiffness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStiffness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMass() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDamping(damping gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, damping)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_damping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDamping() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_damping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseGravity(use_gravity bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use_gravity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_use_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUseGravity() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_use_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGravity(gravity gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gravity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGravity() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], the Jiggle modifier will take colliders into account, keeping them from entering into these collision objects.
*/
//go:nosplit
func (self class) SetUseColliders(use_colliders bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use_colliders)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_use_colliders, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the jiggle modifier is taking physics colliders into account when solving.
*/
//go:nosplit
func (self class) GetUseColliders() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_use_colliders, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the collision mask that the Jiggle modifier will use when reacting to colliders, if the Jiggle modifier is set to take colliders into account.
*/
//go:nosplit
func (self class) SetCollisionMask(collision_mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the collision mask used by the Jiggle modifier when collisions are enabled.
*/
//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointBone2dNode(joint_idx gd.Int, bone2d_node gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, mmm.Get(bone2d_node))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointBone2dNode(ctx gd.Lifetime, joint_idx gd.Int) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the bone index, [param bone_idx], of the Jiggle joint at [param joint_idx]. When possible, this will also update the [code]bone2d_node[/code] of the Jiggle joint based on data provided by the linked skeleton.
*/
//go:nosplit
func (self class) SetJiggleJointBoneIndex(joint_idx gd.Int, bone_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, bone_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_bone_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the index of the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointBoneIndex(joint_idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_bone_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the Jiggle joint at [param joint_idx] should override the default Jiggle joint settings. Setting this to [code]true[/code] will make the joint use its own settings rather than the default ones attached to the modification.
*/
//go:nosplit
func (self class) SetJiggleJointOverride(joint_idx gd.Int, override bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, override)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a boolean that indicates whether the joint at [param joint_idx] is overriding the default Jiggle joint data defined in the modification.
*/
//go:nosplit
func (self class) GetJiggleJointOverride(joint_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the of stiffness of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointStiffness(joint_idx gd.Int, stiffness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, stiffness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the stiffness of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointStiffness(joint_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the of mass of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointMass(joint_idx gd.Int, mass gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, mass)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the amount of mass of the jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointMass(joint_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the amount of damping of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointDamping(joint_idx gd.Int, damping gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, damping)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_damping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the amount of damping of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointDamping(joint_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_damping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the Jiggle joint at [param joint_idx] should use gravity.
*/
//go:nosplit
func (self class) SetJiggleJointUseGravity(joint_idx gd.Int, use_gravity bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, use_gravity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_use_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a boolean that indicates whether the joint at [param joint_idx] is using gravity or not.
*/
//go:nosplit
func (self class) GetJiggleJointUseGravity(joint_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_use_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the gravity vector of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointGravity(joint_idx gd.Int, gravity gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, gravity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [Vector2] representing the amount of gravity the Jiggle joint at [param joint_idx] is influenced by.
*/
//go:nosplit
func (self class) GetJiggleJointGravity(joint_idx gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsSkeletonModification2DJiggle() Expert { return self[0].AsSkeletonModification2DJiggle() }


//go:nosplit
func (self Simple) AsSkeletonModification2DJiggle() Simple { return self[0].AsSkeletonModification2DJiggle() }


//go:nosplit
func (self class) AsSkeletonModification2D() SkeletonModification2D.Expert { return self[0].AsSkeletonModification2D() }


//go:nosplit
func (self Simple) AsSkeletonModification2D() SkeletonModification2D.Simple { return self[0].AsSkeletonModification2D() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("SkeletonModification2DJiggle", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
