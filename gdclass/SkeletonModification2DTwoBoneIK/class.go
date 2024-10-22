package SkeletonModification2DTwoBoneIK

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
This [SkeletonModification2D] uses an algorithm typically called TwoBoneIK. This algorithm works by leveraging the law of cosines and the lengths of the bones to figure out what rotation the bones currently have, and what rotation they need to make a complete triangle, where the first bone, the second bone, and the target form the three vertices of the triangle. Because the algorithm works by making a triangle, it can only operate on two bones.
TwoBoneIK is great for arms, legs, and really any joints that can be represented by just two bones that bend to reach a target. This solver is more lightweight than [SkeletonModification2DFABRIK], but gives similar, natural looking results.

*/
type Go [1]classdb.SkeletonModification2DTwoBoneIK

/*
Sets the [Bone2D] node that is being used as the first bone in the TwoBoneIK modification.
*/
func (self Go) SetJointOneBone2dNode(bone2d_node string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJointOneBone2dNode(gc.String(bone2d_node).NodePath(gc))
}

/*
Returns the [Bone2D] node that is being used as the first bone in the TwoBoneIK modification.
*/
func (self Go) GetJointOneBone2dNode() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetJointOneBone2dNode(gc).String())
}

/*
Sets the index of the [Bone2D] node that is being used as the first bone in the TwoBoneIK modification.
*/
func (self Go) SetJointOneBoneIdx(bone_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJointOneBoneIdx(gd.Int(bone_idx))
}

/*
Returns the index of the [Bone2D] node that is being used as the first bone in the TwoBoneIK modification.
*/
func (self Go) GetJointOneBoneIdx() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetJointOneBoneIdx()))
}

/*
Sets the [Bone2D] node that is being used as the second bone in the TwoBoneIK modification.
*/
func (self Go) SetJointTwoBone2dNode(bone2d_node string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJointTwoBone2dNode(gc.String(bone2d_node).NodePath(gc))
}

/*
Returns the [Bone2D] node that is being used as the second bone in the TwoBoneIK modification.
*/
func (self Go) GetJointTwoBone2dNode() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetJointTwoBone2dNode(gc).String())
}

/*
Sets the index of the [Bone2D] node that is being used as the second bone in the TwoBoneIK modification.
*/
func (self Go) SetJointTwoBoneIdx(bone_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJointTwoBoneIdx(gd.Int(bone_idx))
}

/*
Returns the index of the [Bone2D] node that is being used as the second bone in the TwoBoneIK modification.
*/
func (self Go) GetJointTwoBoneIdx() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetJointTwoBoneIdx()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SkeletonModification2DTwoBoneIK
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SkeletonModification2DTwoBoneIK"))
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

func (self Go) TargetMinimumDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTargetMinimumDistance()))
}

func (self Go) SetTargetMinimumDistance(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTargetMinimumDistance(gd.Float(value))
}

func (self Go) TargetMaximumDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTargetMaximumDistance()))
}

func (self Go) SetTargetMaximumDistance(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTargetMaximumDistance(gd.Float(value))
}

func (self Go) FlipBendDirection() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetFlipBendDirection())
}

func (self Go) SetFlipBendDirection(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFlipBendDirection(value)
}

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
func (self class) AsSkeletonModification2DTwoBoneIK() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSkeletonModification2DTwoBoneIK() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("SkeletonModification2DTwoBoneIK", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
