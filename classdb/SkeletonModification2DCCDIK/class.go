package SkeletonModification2DCCDIK

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/SkeletonModification2D"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This [SkeletonModification2D] uses an algorithm called Cyclic Coordinate Descent Inverse Kinematics, or CCDIK, to manipulate a chain of bones in a [Skeleton2D] so it reaches a defined target.
CCDIK works by rotating a set of bones, typically called a "bone chain", on a single axis. Each bone is rotated to face the target from the tip (by default), which over a chain of bones allow it to rotate properly to reach the target. Because the bones only rotate on a single axis, CCDIK [i]can[/i] look more robotic than other IK solvers.
[b]Note:[/b] The CCDIK modifier has [code]ccdik_joints[/code], which are the data objects that hold the data for each joint in the CCDIK chain. This is different from a bone! CCDIK joints hold the data needed for each bone in the bone chain used by CCDIK.
CCDIK also fully supports angle constraints, allowing for more control over how a solution is met.
*/
type Instance [1]gdclass.SkeletonModification2DCCDIK
type Any interface {
	gd.IsClass
	AsSkeletonModification2DCCDIK() Instance
}

/*
Sets the [Bone2D] node assigned to the CCDIK joint at [param joint_idx].
*/
func (self Instance) SetCcdikJointBone2dNode(joint_idx int, bone2d_nodepath Path.String) {
	class(self).SetCcdikJointBone2dNode(gd.Int(joint_idx), gd.NewString(string(bone2d_nodepath)).NodePath())
}

/*
Returns the [Bone2D] node assigned to the CCDIK joint at [param joint_idx].
*/
func (self Instance) GetCcdikJointBone2dNode(joint_idx int) Path.String {
	return Path.String(class(self).GetCcdikJointBone2dNode(gd.Int(joint_idx)).String())
}

/*
Sets the bone index, [param bone_idx], of the CCDIK joint at [param joint_idx]. When possible, this will also update the [code]bone2d_node[/code] of the CCDIK joint based on data provided by the linked skeleton.
*/
func (self Instance) SetCcdikJointBoneIndex(joint_idx int, bone_idx int) {
	class(self).SetCcdikJointBoneIndex(gd.Int(joint_idx), gd.Int(bone_idx))
}

/*
Returns the index of the [Bone2D] node assigned to the CCDIK joint at [param joint_idx].
*/
func (self Instance) GetCcdikJointBoneIndex(joint_idx int) int {
	return int(int(class(self).GetCcdikJointBoneIndex(gd.Int(joint_idx))))
}

/*
Sets whether the joint at [param joint_idx] is set to rotate from the joint, [code]true[/code], or to rotate from the tip, [code]false[/code].
*/
func (self Instance) SetCcdikJointRotateFromJoint(joint_idx int, rotate_from_joint bool) {
	class(self).SetCcdikJointRotateFromJoint(gd.Int(joint_idx), rotate_from_joint)
}

/*
Returns whether the joint at [param joint_idx] is set to rotate from the joint, [code]true[/code], or to rotate from the tip, [code]false[/code]. The default is to rotate from the tip.
*/
func (self Instance) GetCcdikJointRotateFromJoint(joint_idx int) bool {
	return bool(class(self).GetCcdikJointRotateFromJoint(gd.Int(joint_idx)))
}

/*
Determines whether angle constraints on the CCDIK joint at [param joint_idx] are enabled. When [code]true[/code], constraints will be enabled and taken into account when solving.
*/
func (self Instance) SetCcdikJointEnableConstraint(joint_idx int, enable_constraint bool) {
	class(self).SetCcdikJointEnableConstraint(gd.Int(joint_idx), enable_constraint)
}

/*
Returns whether angle constraints on the CCDIK joint at [param joint_idx] are enabled.
*/
func (self Instance) GetCcdikJointEnableConstraint(joint_idx int) bool {
	return bool(class(self).GetCcdikJointEnableConstraint(gd.Int(joint_idx)))
}

/*
Sets the minimum angle constraint for the joint at [param joint_idx].
*/
func (self Instance) SetCcdikJointConstraintAngleMin(joint_idx int, angle_min Float.X) {
	class(self).SetCcdikJointConstraintAngleMin(gd.Int(joint_idx), gd.Float(angle_min))
}

/*
Returns the minimum angle constraint for the joint at [param joint_idx].
*/
func (self Instance) GetCcdikJointConstraintAngleMin(joint_idx int) Float.X {
	return Float.X(Float.X(class(self).GetCcdikJointConstraintAngleMin(gd.Int(joint_idx))))
}

/*
Sets the maximum angle constraint for the joint at [param joint_idx].
*/
func (self Instance) SetCcdikJointConstraintAngleMax(joint_idx int, angle_max Float.X) {
	class(self).SetCcdikJointConstraintAngleMax(gd.Int(joint_idx), gd.Float(angle_max))
}

/*
Returns the maximum angle constraint for the joint at [param joint_idx].
*/
func (self Instance) GetCcdikJointConstraintAngleMax(joint_idx int) Float.X {
	return Float.X(Float.X(class(self).GetCcdikJointConstraintAngleMax(gd.Int(joint_idx))))
}

/*
Sets whether the CCDIK joint at [param joint_idx] uses an inverted joint constraint.
An inverted joint constraint only constraints the CCDIK joint to the angles [i]outside of[/i] the inputted minimum and maximum angles. For this reason, it is referred to as an inverted joint constraint, as it constraints the joint to the outside of the inputted values.
*/
func (self Instance) SetCcdikJointConstraintAngleInvert(joint_idx int, invert bool) {
	class(self).SetCcdikJointConstraintAngleInvert(gd.Int(joint_idx), invert)
}

/*
Returns whether the CCDIK joint at [param joint_idx] uses an inverted joint constraint. See [method set_ccdik_joint_constraint_angle_invert] for details.
*/
func (self Instance) GetCcdikJointConstraintAngleInvert(joint_idx int) bool {
	return bool(class(self).GetCcdikJointConstraintAngleInvert(gd.Int(joint_idx)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SkeletonModification2DCCDIK

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SkeletonModification2DCCDIK"))
	return Instance{*(*gdclass.SkeletonModification2DCCDIK)(unsafe.Pointer(&object))}
}

func (self Instance) TargetNodepath() Path.String {
	return Path.String(class(self).GetTargetNode().String())
}

func (self Instance) SetTargetNodepath(value Path.String) {
	class(self).SetTargetNode(gd.NewString(string(value)).NodePath())
}

func (self Instance) TipNodepath() Path.String {
	return Path.String(class(self).GetTipNode().String())
}

func (self Instance) SetTipNodepath(value Path.String) {
	class(self).SetTipNode(gd.NewString(string(value)).NodePath())
}

func (self Instance) CcdikDataChainLength() int {
	return int(int(class(self).GetCcdikDataChainLength()))
}

func (self Instance) SetCcdikDataChainLength(value int) {
	class(self).SetCcdikDataChainLength(gd.Int(value))
}

//go:nosplit
func (self class) SetTargetNode(target_nodepath gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(target_nodepath))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_set_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTargetNode() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_get_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTipNode(tip_nodepath gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tip_nodepath))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_set_tip_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTipNode() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_get_tip_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCcdikDataChainLength(length gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_set_ccdik_data_chain_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCcdikDataChainLength() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_get_ccdik_data_chain_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [Bone2D] node assigned to the CCDIK joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetCcdikJointBone2dNode(joint_idx gd.Int, bone2d_nodepath gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, pointers.Get(bone2d_nodepath))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_set_ccdik_joint_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [Bone2D] node assigned to the CCDIK joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetCcdikJointBone2dNode(joint_idx gd.Int) gd.NodePath {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_get_ccdik_joint_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the bone index, [param bone_idx], of the CCDIK joint at [param joint_idx]. When possible, this will also update the [code]bone2d_node[/code] of the CCDIK joint based on data provided by the linked skeleton.
*/
//go:nosplit
func (self class) SetCcdikJointBoneIndex(joint_idx gd.Int, bone_idx gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, bone_idx)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_set_ccdik_joint_bone_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the index of the [Bone2D] node assigned to the CCDIK joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetCcdikJointBoneIndex(joint_idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_get_ccdik_joint_bone_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether the joint at [param joint_idx] is set to rotate from the joint, [code]true[/code], or to rotate from the tip, [code]false[/code].
*/
//go:nosplit
func (self class) SetCcdikJointRotateFromJoint(joint_idx gd.Int, rotate_from_joint bool) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, rotate_from_joint)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_set_ccdik_joint_rotate_from_joint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether the joint at [param joint_idx] is set to rotate from the joint, [code]true[/code], or to rotate from the tip, [code]false[/code]. The default is to rotate from the tip.
*/
//go:nosplit
func (self class) GetCcdikJointRotateFromJoint(joint_idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_get_ccdik_joint_rotate_from_joint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Determines whether angle constraints on the CCDIK joint at [param joint_idx] are enabled. When [code]true[/code], constraints will be enabled and taken into account when solving.
*/
//go:nosplit
func (self class) SetCcdikJointEnableConstraint(joint_idx gd.Int, enable_constraint bool) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, enable_constraint)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_set_ccdik_joint_enable_constraint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether angle constraints on the CCDIK joint at [param joint_idx] are enabled.
*/
//go:nosplit
func (self class) GetCcdikJointEnableConstraint(joint_idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_get_ccdik_joint_enable_constraint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the minimum angle constraint for the joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetCcdikJointConstraintAngleMin(joint_idx gd.Int, angle_min gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, angle_min)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_set_ccdik_joint_constraint_angle_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the minimum angle constraint for the joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetCcdikJointConstraintAngleMin(joint_idx gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_get_ccdik_joint_constraint_angle_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the maximum angle constraint for the joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetCcdikJointConstraintAngleMax(joint_idx gd.Int, angle_max gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, angle_max)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_set_ccdik_joint_constraint_angle_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the maximum angle constraint for the joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetCcdikJointConstraintAngleMax(joint_idx gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_get_ccdik_joint_constraint_angle_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether the CCDIK joint at [param joint_idx] uses an inverted joint constraint.
An inverted joint constraint only constraints the CCDIK joint to the angles [i]outside of[/i] the inputted minimum and maximum angles. For this reason, it is referred to as an inverted joint constraint, as it constraints the joint to the outside of the inputted values.
*/
//go:nosplit
func (self class) SetCcdikJointConstraintAngleInvert(joint_idx gd.Int, invert bool) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, invert)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_set_ccdik_joint_constraint_angle_invert, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether the CCDIK joint at [param joint_idx] uses an inverted joint constraint. See [method set_ccdik_joint_constraint_angle_invert] for details.
*/
//go:nosplit
func (self class) GetCcdikJointConstraintAngleInvert(joint_idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DCCDIK.Bind_get_ccdik_joint_constraint_angle_invert, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSkeletonModification2DCCDIK() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSkeletonModification2DCCDIK() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsSkeletonModification2D() SkeletonModification2D.Advanced {
	return *((*SkeletonModification2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSkeletonModification2D() SkeletonModification2D.Instance {
	return *((*SkeletonModification2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsSkeletonModification2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsSkeletonModification2D(), name)
	}
}
func init() {
	gdclass.Register("SkeletonModification2DCCDIK", func(ptr gd.Object) any {
		return [1]gdclass.SkeletonModification2DCCDIK{*(*gdclass.SkeletonModification2DCCDIK)(unsafe.Pointer(&ptr))}
	})
}
