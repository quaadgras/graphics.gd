package SkeletonModification2DCCDIK

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
This [SkeletonModification2D] uses an algorithm called Cyclic Coordinate Descent Inverse Kinematics, or CCDIK, to manipulate a chain of bones in a [Skeleton2D] so it reaches a defined target.
CCDIK works by rotating a set of bones, typically called a "bone chain", on a single axis. Each bone is rotated to face the target from the tip (by default), which over a chain of bones allow it to rotate properly to reach the target. Because the bones only rotate on a single axis, CCDIK [i]can[/i] look more robotic than other IK solvers.
[b]Note:[/b] The CCDIK modifier has [code]ccdik_joints[/code], which are the data objects that hold the data for each joint in the CCDIK chain. This is different from a bone! CCDIK joints hold the data needed for each bone in the bone chain used by CCDIK.
CCDIK also fully supports angle constraints, allowing for more control over how a solution is met.

*/
type Simple [1]classdb.SkeletonModification2DCCDIK
func (self Simple) SetTargetNode(target_nodepath gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTargetNode(target_nodepath)
}
func (self Simple) GetTargetNode() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetTargetNode(gc))
}
func (self Simple) SetTipNode(tip_nodepath gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTipNode(tip_nodepath)
}
func (self Simple) GetTipNode() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetTipNode(gc))
}
func (self Simple) SetCcdikDataChainLength(length int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCcdikDataChainLength(gd.Int(length))
}
func (self Simple) GetCcdikDataChainLength() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCcdikDataChainLength()))
}
func (self Simple) SetCcdikJointBone2dNode(joint_idx int, bone2d_nodepath gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCcdikJointBone2dNode(gd.Int(joint_idx), bone2d_nodepath)
}
func (self Simple) GetCcdikJointBone2dNode(joint_idx int) gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetCcdikJointBone2dNode(gc, gd.Int(joint_idx)))
}
func (self Simple) SetCcdikJointBoneIndex(joint_idx int, bone_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCcdikJointBoneIndex(gd.Int(joint_idx), gd.Int(bone_idx))
}
func (self Simple) GetCcdikJointBoneIndex(joint_idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCcdikJointBoneIndex(gd.Int(joint_idx))))
}
func (self Simple) SetCcdikJointRotateFromJoint(joint_idx int, rotate_from_joint bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCcdikJointRotateFromJoint(gd.Int(joint_idx), rotate_from_joint)
}
func (self Simple) GetCcdikJointRotateFromJoint(joint_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetCcdikJointRotateFromJoint(gd.Int(joint_idx)))
}
func (self Simple) SetCcdikJointEnableConstraint(joint_idx int, enable_constraint bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCcdikJointEnableConstraint(gd.Int(joint_idx), enable_constraint)
}
func (self Simple) GetCcdikJointEnableConstraint(joint_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetCcdikJointEnableConstraint(gd.Int(joint_idx)))
}
func (self Simple) SetCcdikJointConstraintAngleMin(joint_idx int, angle_min float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCcdikJointConstraintAngleMin(gd.Int(joint_idx), gd.Float(angle_min))
}
func (self Simple) GetCcdikJointConstraintAngleMin(joint_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetCcdikJointConstraintAngleMin(gd.Int(joint_idx))))
}
func (self Simple) SetCcdikJointConstraintAngleMax(joint_idx int, angle_max float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCcdikJointConstraintAngleMax(gd.Int(joint_idx), gd.Float(angle_max))
}
func (self Simple) GetCcdikJointConstraintAngleMax(joint_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetCcdikJointConstraintAngleMax(gd.Int(joint_idx))))
}
func (self Simple) SetCcdikJointConstraintAngleInvert(joint_idx int, invert bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCcdikJointConstraintAngleInvert(gd.Int(joint_idx), invert)
}
func (self Simple) GetCcdikJointConstraintAngleInvert(joint_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetCcdikJointConstraintAngleInvert(gd.Int(joint_idx)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.SkeletonModification2DCCDIK
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetTargetNode(target_nodepath gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(target_nodepath))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_set_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTargetNode(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_get_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTipNode(tip_nodepath gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(tip_nodepath))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_set_tip_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTipNode(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_get_tip_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCcdikDataChainLength(length gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_set_ccdik_data_chain_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCcdikDataChainLength() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_get_ccdik_data_chain_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [Bone2D] node assigned to the CCDIK joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetCcdikJointBone2dNode(joint_idx gd.Int, bone2d_nodepath gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, mmm.Get(bone2d_nodepath))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_set_ccdik_joint_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Bone2D] node assigned to the CCDIK joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetCcdikJointBone2dNode(ctx gd.Lifetime, joint_idx gd.Int) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_get_ccdik_joint_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the bone index, [param bone_idx], of the CCDIK joint at [param joint_idx]. When possible, this will also update the [code]bone2d_node[/code] of the CCDIK joint based on data provided by the linked skeleton.
*/
//go:nosplit
func (self class) SetCcdikJointBoneIndex(joint_idx gd.Int, bone_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, bone_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_set_ccdik_joint_bone_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the index of the [Bone2D] node assigned to the CCDIK joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetCcdikJointBoneIndex(joint_idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_get_ccdik_joint_bone_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the joint at [param joint_idx] is set to rotate from the joint, [code]true[/code], or to rotate from the tip, [code]false[/code].
*/
//go:nosplit
func (self class) SetCcdikJointRotateFromJoint(joint_idx gd.Int, rotate_from_joint bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, rotate_from_joint)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_set_ccdik_joint_rotate_from_joint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the joint at [param joint_idx] is set to rotate from the joint, [code]true[/code], or to rotate from the tip, [code]false[/code]. The default is to rotate from the tip.
*/
//go:nosplit
func (self class) GetCcdikJointRotateFromJoint(joint_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_get_ccdik_joint_rotate_from_joint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Determines whether angle constraints on the CCDIK joint at [param joint_idx] are enabled. When [code]true[/code], constraints will be enabled and taken into account when solving.
*/
//go:nosplit
func (self class) SetCcdikJointEnableConstraint(joint_idx gd.Int, enable_constraint bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, enable_constraint)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_set_ccdik_joint_enable_constraint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether angle constraints on the CCDIK joint at [param joint_idx] are enabled.
*/
//go:nosplit
func (self class) GetCcdikJointEnableConstraint(joint_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_get_ccdik_joint_enable_constraint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the minimum angle constraint for the joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetCcdikJointConstraintAngleMin(joint_idx gd.Int, angle_min gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, angle_min)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_set_ccdik_joint_constraint_angle_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the minimum angle constraint for the joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetCcdikJointConstraintAngleMin(joint_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_get_ccdik_joint_constraint_angle_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the maximum angle constraint for the joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetCcdikJointConstraintAngleMax(joint_idx gd.Int, angle_max gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, angle_max)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_set_ccdik_joint_constraint_angle_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the maximum angle constraint for the joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetCcdikJointConstraintAngleMax(joint_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_get_ccdik_joint_constraint_angle_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the CCDIK joint at [param joint_idx] uses an inverted joint constraint.
An inverted joint constraint only constraints the CCDIK joint to the angles [i]outside of[/i] the inputted minimum and maximum angles. For this reason, it is referred to as an inverted joint constraint, as it constraints the joint to the outside of the inputted values.
*/
//go:nosplit
func (self class) SetCcdikJointConstraintAngleInvert(joint_idx gd.Int, invert bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, invert)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_set_ccdik_joint_constraint_angle_invert, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the CCDIK joint at [param joint_idx] uses an inverted joint constraint. See [method set_ccdik_joint_constraint_angle_invert] for details.
*/
//go:nosplit
func (self class) GetCcdikJointConstraintAngleInvert(joint_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DCCDIK.Bind_get_ccdik_joint_constraint_angle_invert, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsSkeletonModification2DCCDIK() Expert { return self[0].AsSkeletonModification2DCCDIK() }


//go:nosplit
func (self Simple) AsSkeletonModification2DCCDIK() Simple { return self[0].AsSkeletonModification2DCCDIK() }


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
func init() {classdb.Register("SkeletonModification2DCCDIK", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
