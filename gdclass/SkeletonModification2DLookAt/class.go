package SkeletonModification2DLookAt

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
This [SkeletonModification2D] rotates a bone to look a target. This is extremely helpful for moving character's head to look at the player, rotating a turret to look at a target, or any other case where you want to make a bone rotate towards something quickly and easily.

*/
type Go [1]classdb.SkeletonModification2DLookAt

/*
Sets the amount of additional rotation that is to be applied after executing the modification. This allows for offsetting the results by the inputted rotation amount.
*/
func (self Go) SetAdditionalRotation(rotation float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAdditionalRotation(gd.Float(rotation))
}

/*
Returns the amount of additional rotation that is applied after the LookAt modification executes.
*/
func (self Go) GetAdditionalRotation() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetAdditionalRotation()))
}

/*
Sets whether this modification will use constraints or not. When [code]true[/code], constraints will be applied when solving the LookAt modification.
*/
func (self Go) SetEnableConstraint(enable_constraint bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnableConstraint(enable_constraint)
}

/*
Returns [code]true[/code] if the LookAt modification is using constraints.
*/
func (self Go) GetEnableConstraint() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetEnableConstraint())
}

/*
Sets the constraint's minimum allowed angle.
*/
func (self Go) SetConstraintAngleMin(angle_min float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetConstraintAngleMin(gd.Float(angle_min))
}

/*
Returns the constraint's minimum allowed angle.
*/
func (self Go) GetConstraintAngleMin() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetConstraintAngleMin()))
}

/*
Sets the constraint's maximum allowed angle.
*/
func (self Go) SetConstraintAngleMax(angle_max float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetConstraintAngleMax(gd.Float(angle_max))
}

/*
Returns the constraint's maximum allowed angle.
*/
func (self Go) GetConstraintAngleMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetConstraintAngleMax()))
}

/*
When [code]true[/code], the modification will use an inverted joint constraint.
An inverted joint constraint only constraints the [Bone2D] to the angles [i]outside of[/i] the inputted minimum and maximum angles. For this reason, it is referred to as an inverted joint constraint, as it constraints the joint to the outside of the inputted values.
*/
func (self Go) SetConstraintAngleInvert(invert bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetConstraintAngleInvert(invert)
}

/*
Returns whether the constraints to this modification are inverted or not.
*/
func (self Go) GetConstraintAngleInvert() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetConstraintAngleInvert())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SkeletonModification2DLookAt
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SkeletonModification2DLookAt"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) BoneIndex() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetBoneIndex()))
}

func (self Go) SetBoneIndex(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBoneIndex(gd.Int(value))
}

func (self Go) Bone2dNode() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetBone2dNode(gc).String())
}

func (self Go) SetBone2dNode(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBone2dNode(gc.String(value).NodePath(gc))
}

func (self Go) TargetNodepath() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetTargetNode(gc).String())
}

func (self Go) SetTargetNodepath(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTargetNode(gc.String(value).NodePath(gc))
}

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
func (self class) AsSkeletonModification2DLookAt() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSkeletonModification2DLookAt() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("SkeletonModification2DLookAt", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
