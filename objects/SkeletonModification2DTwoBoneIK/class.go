package SkeletonModification2DTwoBoneIK

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/SkeletonModification2D"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
This [SkeletonModification2D] uses an algorithm typically called TwoBoneIK. This algorithm works by leveraging the law of cosines and the lengths of the bones to figure out what rotation the bones currently have, and what rotation they need to make a complete triangle, where the first bone, the second bone, and the target form the three vertices of the triangle. Because the algorithm works by making a triangle, it can only operate on two bones.
TwoBoneIK is great for arms, legs, and really any joints that can be represented by just two bones that bend to reach a target. This solver is more lightweight than [SkeletonModification2DFABRIK], but gives similar, natural looking results.
*/
type Instance [1]classdb.SkeletonModification2DTwoBoneIK

/*
Sets the [Bone2D] node that is being used as the first bone in the TwoBoneIK modification.
*/
func (self Instance) SetJointOneBone2dNode(bone2d_node string) {
	class(self).SetJointOneBone2dNode(gd.NewString(bone2d_node).NodePath())
}

/*
Returns the [Bone2D] node that is being used as the first bone in the TwoBoneIK modification.
*/
func (self Instance) GetJointOneBone2dNode() string {
	return string(class(self).GetJointOneBone2dNode().String())
}

/*
Sets the index of the [Bone2D] node that is being used as the first bone in the TwoBoneIK modification.
*/
func (self Instance) SetJointOneBoneIdx(bone_idx int) {
	class(self).SetJointOneBoneIdx(gd.Int(bone_idx))
}

/*
Returns the index of the [Bone2D] node that is being used as the first bone in the TwoBoneIK modification.
*/
func (self Instance) GetJointOneBoneIdx() int {
	return int(int(class(self).GetJointOneBoneIdx()))
}

/*
Sets the [Bone2D] node that is being used as the second bone in the TwoBoneIK modification.
*/
func (self Instance) SetJointTwoBone2dNode(bone2d_node string) {
	class(self).SetJointTwoBone2dNode(gd.NewString(bone2d_node).NodePath())
}

/*
Returns the [Bone2D] node that is being used as the second bone in the TwoBoneIK modification.
*/
func (self Instance) GetJointTwoBone2dNode() string {
	return string(class(self).GetJointTwoBone2dNode().String())
}

/*
Sets the index of the [Bone2D] node that is being used as the second bone in the TwoBoneIK modification.
*/
func (self Instance) SetJointTwoBoneIdx(bone_idx int) {
	class(self).SetJointTwoBoneIdx(gd.Int(bone_idx))
}

/*
Returns the index of the [Bone2D] node that is being used as the second bone in the TwoBoneIK modification.
*/
func (self Instance) GetJointTwoBoneIdx() int {
	return int(int(class(self).GetJointTwoBoneIdx()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.SkeletonModification2DTwoBoneIK

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SkeletonModification2DTwoBoneIK"))
	return Instance{classdb.SkeletonModification2DTwoBoneIK(object)}
}

func (self Instance) TargetNodepath() string {
	return string(class(self).GetTargetNode().String())
}

func (self Instance) SetTargetNodepath(value string) {
	class(self).SetTargetNode(gd.NewString(value).NodePath())
}

func (self Instance) TargetMinimumDistance() float64 {
	return float64(float64(class(self).GetTargetMinimumDistance()))
}

func (self Instance) SetTargetMinimumDistance(value float64) {
	class(self).SetTargetMinimumDistance(gd.Float(value))
}

func (self Instance) TargetMaximumDistance() float64 {
	return float64(float64(class(self).GetTargetMaximumDistance()))
}

func (self Instance) SetTargetMaximumDistance(value float64) {
	class(self).SetTargetMaximumDistance(gd.Float(value))
}

func (self Instance) FlipBendDirection() bool {
	return bool(class(self).GetFlipBendDirection())
}

func (self Instance) SetFlipBendDirection(value bool) {
	class(self).SetFlipBendDirection(value)
}

//go:nosplit
func (self class) SetTargetNode(target_nodepath gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(target_nodepath))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DTwoBoneIK.Bind_set_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTargetNode() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DTwoBoneIK.Bind_get_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTargetMinimumDistance(minimum_distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, minimum_distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DTwoBoneIK.Bind_set_target_minimum_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTargetMinimumDistance() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DTwoBoneIK.Bind_get_target_minimum_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTargetMaximumDistance(maximum_distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, maximum_distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DTwoBoneIK.Bind_set_target_maximum_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTargetMaximumDistance() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DTwoBoneIK.Bind_get_target_maximum_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlipBendDirection(flip_direction bool) {
	var frame = callframe.New()
	callframe.Arg(frame, flip_direction)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DTwoBoneIK.Bind_set_flip_bend_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFlipBendDirection() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DTwoBoneIK.Bind_get_flip_bend_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [Bone2D] node that is being used as the first bone in the TwoBoneIK modification.
*/
//go:nosplit
func (self class) SetJointOneBone2dNode(bone2d_node gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(bone2d_node))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DTwoBoneIK.Bind_set_joint_one_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [Bone2D] node that is being used as the first bone in the TwoBoneIK modification.
*/
//go:nosplit
func (self class) GetJointOneBone2dNode() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DTwoBoneIK.Bind_get_joint_one_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the index of the [Bone2D] node that is being used as the first bone in the TwoBoneIK modification.
*/
//go:nosplit
func (self class) SetJointOneBoneIdx(bone_idx gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DTwoBoneIK.Bind_set_joint_one_bone_idx, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the index of the [Bone2D] node that is being used as the first bone in the TwoBoneIK modification.
*/
//go:nosplit
func (self class) GetJointOneBoneIdx() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DTwoBoneIK.Bind_get_joint_one_bone_idx, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [Bone2D] node that is being used as the second bone in the TwoBoneIK modification.
*/
//go:nosplit
func (self class) SetJointTwoBone2dNode(bone2d_node gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(bone2d_node))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DTwoBoneIK.Bind_set_joint_two_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [Bone2D] node that is being used as the second bone in the TwoBoneIK modification.
*/
//go:nosplit
func (self class) GetJointTwoBone2dNode() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DTwoBoneIK.Bind_get_joint_two_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the index of the [Bone2D] node that is being used as the second bone in the TwoBoneIK modification.
*/
//go:nosplit
func (self class) SetJointTwoBoneIdx(bone_idx gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DTwoBoneIK.Bind_set_joint_two_bone_idx, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the index of the [Bone2D] node that is being used as the second bone in the TwoBoneIK modification.
*/
//go:nosplit
func (self class) GetJointTwoBoneIdx() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DTwoBoneIK.Bind_get_joint_two_bone_idx, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSkeletonModification2DTwoBoneIK() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSkeletonModification2DTwoBoneIK() Instance {
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
	classdb.Register("SkeletonModification2DTwoBoneIK", func(ptr gd.Object) any { return classdb.SkeletonModification2DTwoBoneIK(ptr) })
}
