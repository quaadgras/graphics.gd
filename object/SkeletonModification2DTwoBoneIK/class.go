package SkeletonModification2DTwoBoneIK

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
This [SkeletonModification2D] uses an algorithm typically called TwoBoneIK. This algorithm works by leveraging the law of cosines and the lengths of the bones to figure out what rotation the bones currently have, and what rotation they need to make a complete triangle, where the first bone, the second bone, and the target form the three vertices of the triangle. Because the algorithm works by making a triangle, it can only operate on two bones.
TwoBoneIK is great for arms, legs, and really any joints that can be represented by just two bones that bend to reach a target. This solver is more lightweight than [SkeletonModification2DFABRIK], but gives similar, natural looking results.

*/
type Simple [1]classdb.SkeletonModification2DTwoBoneIK
func (self Simple) SetTargetNode(target_nodepath gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTargetNode(target_nodepath)
}
func (self Simple) GetTargetNode() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetTargetNode(gc))
}
func (self Simple) SetTargetMinimumDistance(minimum_distance float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTargetMinimumDistance(gd.Float(minimum_distance))
}
func (self Simple) GetTargetMinimumDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTargetMinimumDistance()))
}
func (self Simple) SetTargetMaximumDistance(maximum_distance float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTargetMaximumDistance(gd.Float(maximum_distance))
}
func (self Simple) GetTargetMaximumDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTargetMaximumDistance()))
}
func (self Simple) SetFlipBendDirection(flip_direction bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFlipBendDirection(flip_direction)
}
func (self Simple) GetFlipBendDirection() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetFlipBendDirection())
}
func (self Simple) SetJointOneBone2dNode(bone2d_node gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetJointOneBone2dNode(bone2d_node)
}
func (self Simple) GetJointOneBone2dNode() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetJointOneBone2dNode(gc))
}
func (self Simple) SetJointOneBoneIdx(bone_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetJointOneBoneIdx(gd.Int(bone_idx))
}
func (self Simple) GetJointOneBoneIdx() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetJointOneBoneIdx()))
}
func (self Simple) SetJointTwoBone2dNode(bone2d_node gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetJointTwoBone2dNode(bone2d_node)
}
func (self Simple) GetJointTwoBone2dNode() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetJointTwoBone2dNode(gc))
}
func (self Simple) SetJointTwoBoneIdx(bone_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetJointTwoBoneIdx(gd.Int(bone_idx))
}
func (self Simple) GetJointTwoBoneIdx() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetJointTwoBoneIdx()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.SkeletonModification2DTwoBoneIK
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DTwoBoneIK.Bind_set_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTargetNode(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DTwoBoneIK.Bind_get_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTargetMinimumDistance(minimum_distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, minimum_distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DTwoBoneIK.Bind_set_target_minimum_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTargetMinimumDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DTwoBoneIK.Bind_get_target_minimum_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTargetMaximumDistance(maximum_distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, maximum_distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DTwoBoneIK.Bind_set_target_maximum_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTargetMaximumDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DTwoBoneIK.Bind_get_target_maximum_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFlipBendDirection(flip_direction bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flip_direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DTwoBoneIK.Bind_set_flip_bend_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFlipBendDirection() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DTwoBoneIK.Bind_get_flip_bend_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [Bone2D] node that is being used as the first bone in the TwoBoneIK modification.
*/
//go:nosplit
func (self class) SetJointOneBone2dNode(bone2d_node gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bone2d_node))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DTwoBoneIK.Bind_set_joint_one_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Bone2D] node that is being used as the first bone in the TwoBoneIK modification.
*/
//go:nosplit
func (self class) GetJointOneBone2dNode(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DTwoBoneIK.Bind_get_joint_one_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the index of the [Bone2D] node that is being used as the first bone in the TwoBoneIK modification.
*/
//go:nosplit
func (self class) SetJointOneBoneIdx(bone_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DTwoBoneIK.Bind_set_joint_one_bone_idx, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the index of the [Bone2D] node that is being used as the first bone in the TwoBoneIK modification.
*/
//go:nosplit
func (self class) GetJointOneBoneIdx() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DTwoBoneIK.Bind_get_joint_one_bone_idx, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [Bone2D] node that is being used as the second bone in the TwoBoneIK modification.
*/
//go:nosplit
func (self class) SetJointTwoBone2dNode(bone2d_node gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bone2d_node))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DTwoBoneIK.Bind_set_joint_two_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Bone2D] node that is being used as the second bone in the TwoBoneIK modification.
*/
//go:nosplit
func (self class) GetJointTwoBone2dNode(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DTwoBoneIK.Bind_get_joint_two_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the index of the [Bone2D] node that is being used as the second bone in the TwoBoneIK modification.
*/
//go:nosplit
func (self class) SetJointTwoBoneIdx(bone_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DTwoBoneIK.Bind_set_joint_two_bone_idx, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the index of the [Bone2D] node that is being used as the second bone in the TwoBoneIK modification.
*/
//go:nosplit
func (self class) GetJointTwoBoneIdx() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DTwoBoneIK.Bind_get_joint_two_bone_idx, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsSkeletonModification2DTwoBoneIK() Expert { return self[0].AsSkeletonModification2DTwoBoneIK() }


//go:nosplit
func (self Simple) AsSkeletonModification2DTwoBoneIK() Simple { return self[0].AsSkeletonModification2DTwoBoneIK() }


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
func init() {classdb.Register("SkeletonModification2DTwoBoneIK", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
