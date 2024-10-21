package SkeletonModification2DLookAt

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
This [SkeletonModification2D] rotates a bone to look a target. This is extremely helpful for moving character's head to look at the player, rotating a turret to look at a target, or any other case where you want to make a bone rotate towards something quickly and easily.

*/
type Simple [1]classdb.SkeletonModification2DLookAt
func (self Simple) SetBone2dNode(bone2d_nodepath gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBone2dNode(bone2d_nodepath)
}
func (self Simple) GetBone2dNode() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetBone2dNode(gc))
}
func (self Simple) SetBoneIndex(bone_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBoneIndex(gd.Int(bone_idx))
}
func (self Simple) GetBoneIndex() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBoneIndex()))
}
func (self Simple) SetTargetNode(target_nodepath gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTargetNode(target_nodepath)
}
func (self Simple) GetTargetNode() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetTargetNode(gc))
}
func (self Simple) SetAdditionalRotation(rotation float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAdditionalRotation(gd.Float(rotation))
}
func (self Simple) GetAdditionalRotation() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAdditionalRotation()))
}
func (self Simple) SetEnableConstraint(enable_constraint bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnableConstraint(enable_constraint)
}
func (self Simple) GetEnableConstraint() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetEnableConstraint())
}
func (self Simple) SetConstraintAngleMin(angle_min float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetConstraintAngleMin(gd.Float(angle_min))
}
func (self Simple) GetConstraintAngleMin() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetConstraintAngleMin()))
}
func (self Simple) SetConstraintAngleMax(angle_max float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetConstraintAngleMax(gd.Float(angle_max))
}
func (self Simple) GetConstraintAngleMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetConstraintAngleMax()))
}
func (self Simple) SetConstraintAngleInvert(invert bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetConstraintAngleInvert(invert)
}
func (self Simple) GetConstraintAngleInvert() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetConstraintAngleInvert())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.SkeletonModification2DLookAt
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetBone2dNode(bone2d_nodepath gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bone2d_nodepath))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DLookAt.Bind_set_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBone2dNode(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DLookAt.Bind_get_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBoneIndex(bone_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DLookAt.Bind_set_bone_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBoneIndex() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DLookAt.Bind_get_bone_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTargetNode(target_nodepath gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(target_nodepath))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DLookAt.Bind_set_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTargetNode(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DLookAt.Bind_get_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the amount of additional rotation that is to be applied after executing the modification. This allows for offsetting the results by the inputted rotation amount.
*/
//go:nosplit
func (self class) SetAdditionalRotation(rotation gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rotation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DLookAt.Bind_set_additional_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the amount of additional rotation that is applied after the LookAt modification executes.
*/
//go:nosplit
func (self class) GetAdditionalRotation() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DLookAt.Bind_get_additional_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether this modification will use constraints or not. When [code]true[/code], constraints will be applied when solving the LookAt modification.
*/
//go:nosplit
func (self class) SetEnableConstraint(enable_constraint bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable_constraint)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DLookAt.Bind_set_enable_constraint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the LookAt modification is using constraints.
*/
//go:nosplit
func (self class) GetEnableConstraint() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DLookAt.Bind_get_enable_constraint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the constraint's minimum allowed angle.
*/
//go:nosplit
func (self class) SetConstraintAngleMin(angle_min gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angle_min)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DLookAt.Bind_set_constraint_angle_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the constraint's minimum allowed angle.
*/
//go:nosplit
func (self class) GetConstraintAngleMin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DLookAt.Bind_get_constraint_angle_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the constraint's maximum allowed angle.
*/
//go:nosplit
func (self class) SetConstraintAngleMax(angle_max gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angle_max)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DLookAt.Bind_set_constraint_angle_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the constraint's maximum allowed angle.
*/
//go:nosplit
func (self class) GetConstraintAngleMax() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DLookAt.Bind_get_constraint_angle_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
When [code]true[/code], the modification will use an inverted joint constraint.
An inverted joint constraint only constraints the [Bone2D] to the angles [i]outside of[/i] the inputted minimum and maximum angles. For this reason, it is referred to as an inverted joint constraint, as it constraints the joint to the outside of the inputted values.
*/
//go:nosplit
func (self class) SetConstraintAngleInvert(invert bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, invert)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DLookAt.Bind_set_constraint_angle_invert, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the constraints to this modification are inverted or not.
*/
//go:nosplit
func (self class) GetConstraintAngleInvert() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DLookAt.Bind_get_constraint_angle_invert, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsSkeletonModification2DLookAt() Expert { return self[0].AsSkeletonModification2DLookAt() }


//go:nosplit
func (self Simple) AsSkeletonModification2DLookAt() Simple { return self[0].AsSkeletonModification2DLookAt() }


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
func init() {classdb.Register("SkeletonModification2DLookAt", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
