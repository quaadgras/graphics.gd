package SkeletonModification2DCCDIK

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/SkeletonModification2D"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This [SkeletonModification2D] uses an algorithm called Cyclic Coordinate Descent Inverse Kinematics, or CCDIK, to manipulate a chain of bones in a [Skeleton2D] so it reaches a defined target.
CCDIK works by rotating a set of bones, typically called a "bone chain", on a single axis. Each bone is rotated to face the target from the tip (by default), which over a chain of bones allow it to rotate properly to reach the target. Because the bones only rotate on a single axis, CCDIK [i]can[/i] look more robotic than other IK solvers.
[b]Note:[/b] The CCDIK modifier has [code]ccdik_joints[/code], which are the data objects that hold the data for each joint in the CCDIK chain. This is different from a bone! CCDIK joints hold the data needed for each bone in the bone chain used by CCDIK.
CCDIK also fully supports angle constraints, allowing for more control over how a solution is met.

*/
type Go [1]classdb.SkeletonModification2DCCDIK

/*
Sets the [Bone2D] node assigned to the CCDIK joint at [param joint_idx].
*/
func (self Go) SetCcdikJointBone2dNode(joint_idx int, bone2d_nodepath string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCcdikJointBone2dNode(gd.Int(joint_idx), gc.String(bone2d_nodepath).NodePath(gc))
}

/*
Returns the [Bone2D] node assigned to the CCDIK joint at [param joint_idx].
*/
func (self Go) GetCcdikJointBone2dNode(joint_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetCcdikJointBone2dNode(gc, gd.Int(joint_idx)).String())
}

/*
Sets the bone index, [param bone_idx], of the CCDIK joint at [param joint_idx]. When possible, this will also update the [code]bone2d_node[/code] of the CCDIK joint based on data provided by the linked skeleton.
*/
func (self Go) SetCcdikJointBoneIndex(joint_idx int, bone_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCcdikJointBoneIndex(gd.Int(joint_idx), gd.Int(bone_idx))
}

/*
Returns the index of the [Bone2D] node assigned to the CCDIK joint at [param joint_idx].
*/
func (self Go) GetCcdikJointBoneIndex(joint_idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCcdikJointBoneIndex(gd.Int(joint_idx))))
}

/*
Sets whether the joint at [param joint_idx] is set to rotate from the joint, [code]true[/code], or to rotate from the tip, [code]false[/code].
*/
func (self Go) SetCcdikJointRotateFromJoint(joint_idx int, rotate_from_joint bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCcdikJointRotateFromJoint(gd.Int(joint_idx), rotate_from_joint)
}

/*
Returns whether the joint at [param joint_idx] is set to rotate from the joint, [code]true[/code], or to rotate from the tip, [code]false[/code]. The default is to rotate from the tip.
*/
func (self Go) GetCcdikJointRotateFromJoint(joint_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetCcdikJointRotateFromJoint(gd.Int(joint_idx)))
}

/*
Determines whether angle constraints on the CCDIK joint at [param joint_idx] are enabled. When [code]true[/code], constraints will be enabled and taken into account when solving.
*/
func (self Go) SetCcdikJointEnableConstraint(joint_idx int, enable_constraint bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCcdikJointEnableConstraint(gd.Int(joint_idx), enable_constraint)
}

/*
Returns whether angle constraints on the CCDIK joint at [param joint_idx] are enabled.
*/
func (self Go) GetCcdikJointEnableConstraint(joint_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetCcdikJointEnableConstraint(gd.Int(joint_idx)))
}

/*
Sets the minimum angle constraint for the joint at [param joint_idx].
*/
func (self Go) SetCcdikJointConstraintAngleMin(joint_idx int, angle_min float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCcdikJointConstraintAngleMin(gd.Int(joint_idx), gd.Float(angle_min))
}

/*
Returns the minimum angle constraint for the joint at [param joint_idx].
*/
func (self Go) GetCcdikJointConstraintAngleMin(joint_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetCcdikJointConstraintAngleMin(gd.Int(joint_idx))))
}

/*
Sets the maximum angle constraint for the joint at [param joint_idx].
*/
func (self Go) SetCcdikJointConstraintAngleMax(joint_idx int, angle_max float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCcdikJointConstraintAngleMax(gd.Int(joint_idx), gd.Float(angle_max))
}

/*
Returns the maximum angle constraint for the joint at [param joint_idx].
*/
func (self Go) GetCcdikJointConstraintAngleMax(joint_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetCcdikJointConstraintAngleMax(gd.Int(joint_idx))))
}

/*
Sets whether the CCDIK joint at [param joint_idx] uses an inverted joint constraint.
An inverted joint constraint only constraints the CCDIK joint to the angles [i]outside of[/i] the inputted minimum and maximum angles. For this reason, it is referred to as an inverted joint constraint, as it constraints the joint to the outside of the inputted values.
*/
func (self Go) SetCcdikJointConstraintAngleInvert(joint_idx int, invert bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCcdikJointConstraintAngleInvert(gd.Int(joint_idx), invert)
}

/*
Returns whether the CCDIK joint at [param joint_idx] uses an inverted joint constraint. See [method set_ccdik_joint_constraint_angle_invert] for details.
*/
func (self Go) GetCcdikJointConstraintAngleInvert(joint_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetCcdikJointConstraintAngleInvert(gd.Int(joint_idx)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SkeletonModification2DCCDIK
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SkeletonModification2DCCDIK"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) TargetNodepath() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetTargetNode(gc).String())
}

func (self Go) SetTargetNodepath(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTargetNode(gc.String(value).NodePath(gc))
}

func (self Go) TipNodepath() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetTipNode(gc).String())
}

func (self Go) SetTipNodepath(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTipNode(gc.String(value).NodePath(gc))
}

func (self Go) CcdikDataChainLength() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCcdikDataChainLength()))
}

func (self Go) SetCcdikDataChainLength(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCcdikDataChainLength(gd.Int(value))
}

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
func (self class) AsSkeletonModification2DCCDIK() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSkeletonModification2DCCDIK() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsSkeletonModification2D() SkeletonModification2D.GD { return *((*SkeletonModification2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsSkeletonModification2D() SkeletonModification2D.Go { return *((*SkeletonModification2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsSkeletonModification2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsSkeletonModification2D(), name)
	}
}
func init() {classdb.Register("SkeletonModification2DCCDIK", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
