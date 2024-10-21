package SkeletonProfile

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This resource is used in [EditorScenePostImport]. Some parameters are referring to bones in [Skeleton3D], [Skin], [Animation], and some other nodes are rewritten based on the parameters of [SkeletonProfile].
[b]Note:[/b] These parameters need to be set only when creating a custom profile. In [SkeletonProfileHumanoid], they are defined internally as read-only values.

*/
type Simple [1]classdb.SkeletonProfile
func (self Simple) SetRootBone(bone_name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRootBone(gc.StringName(bone_name))
}
func (self Simple) GetRootBone() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetRootBone(gc).String())
}
func (self Simple) SetScaleBaseBone(bone_name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetScaleBaseBone(gc.StringName(bone_name))
}
func (self Simple) GetScaleBaseBone() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetScaleBaseBone(gc).String())
}
func (self Simple) SetGroupSize(size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGroupSize(gd.Int(size))
}
func (self Simple) GetGroupSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetGroupSize()))
}
func (self Simple) GetGroupName(group_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetGroupName(gc, gd.Int(group_idx)).String())
}
func (self Simple) SetGroupName(group_idx int, group_name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGroupName(gd.Int(group_idx), gc.StringName(group_name))
}
func (self Simple) GetTexture(group_idx int) [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetTexture(gc, gd.Int(group_idx)))
}
func (self Simple) SetTexture(group_idx int, texture [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTexture(gd.Int(group_idx), texture)
}
func (self Simple) SetBoneSize(size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBoneSize(gd.Int(size))
}
func (self Simple) GetBoneSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBoneSize()))
}
func (self Simple) FindBone(bone_name string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).FindBone(gc.StringName(bone_name))))
}
func (self Simple) GetBoneName(bone_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetBoneName(gc, gd.Int(bone_idx)).String())
}
func (self Simple) SetBoneName(bone_idx int, bone_name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBoneName(gd.Int(bone_idx), gc.StringName(bone_name))
}
func (self Simple) GetBoneParent(bone_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetBoneParent(gc, gd.Int(bone_idx)).String())
}
func (self Simple) SetBoneParent(bone_idx int, bone_parent string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBoneParent(gd.Int(bone_idx), gc.StringName(bone_parent))
}
func (self Simple) GetTailDirection(bone_idx int) classdb.SkeletonProfileTailDirection {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.SkeletonProfileTailDirection(Expert(self).GetTailDirection(gd.Int(bone_idx)))
}
func (self Simple) SetTailDirection(bone_idx int, tail_direction classdb.SkeletonProfileTailDirection) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTailDirection(gd.Int(bone_idx), tail_direction)
}
func (self Simple) GetBoneTail(bone_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetBoneTail(gc, gd.Int(bone_idx)).String())
}
func (self Simple) SetBoneTail(bone_idx int, bone_tail string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBoneTail(gd.Int(bone_idx), gc.StringName(bone_tail))
}
func (self Simple) GetReferencePose(bone_idx int) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).GetReferencePose(gd.Int(bone_idx)))
}
func (self Simple) SetReferencePose(bone_idx int, bone_name gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetReferencePose(gd.Int(bone_idx), bone_name)
}
func (self Simple) GetHandleOffset(bone_idx int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetHandleOffset(gd.Int(bone_idx)))
}
func (self Simple) SetHandleOffset(bone_idx int, handle_offset gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHandleOffset(gd.Int(bone_idx), handle_offset)
}
func (self Simple) GetGroup(bone_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetGroup(gc, gd.Int(bone_idx)).String())
}
func (self Simple) SetGroup(bone_idx int, group string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGroup(gd.Int(bone_idx), gc.StringName(group))
}
func (self Simple) IsRequired(bone_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsRequired(gd.Int(bone_idx)))
}
func (self Simple) SetRequired(bone_idx int, required bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRequired(gd.Int(bone_idx), required)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.SkeletonProfile
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetRootBone(bone_name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bone_name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_set_root_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRootBone(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_get_root_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScaleBaseBone(bone_name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bone_name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_set_scale_base_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetScaleBaseBone(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_get_scale_base_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGroupSize(size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_set_group_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGroupSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_get_group_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the name of the group at [param group_idx] that will be the drawing group in the [BoneMap] editor.
*/
//go:nosplit
func (self class) GetGroupName(ctx gd.Lifetime, group_idx gd.Int) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, group_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_get_group_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the name of the group at [param group_idx] that will be the drawing group in the [BoneMap] editor.
*/
//go:nosplit
func (self class) SetGroupName(group_idx gd.Int, group_name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, group_idx)
	callframe.Arg(frame, mmm.Get(group_name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_set_group_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the texture of the group at [param group_idx] that will be the drawing group background image in the [BoneMap] editor.
*/
//go:nosplit
func (self class) GetTexture(ctx gd.Lifetime, group_idx gd.Int) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, group_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the texture of the group at [param group_idx] that will be the drawing group background image in the [BoneMap] editor.
*/
//go:nosplit
func (self class) SetTexture(group_idx gd.Int, texture object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, group_idx)
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetBoneSize(size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_set_bone_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBoneSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_get_bone_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the bone index that matches [param bone_name] as its name.
*/
//go:nosplit
func (self class) FindBone(bone_name gd.StringName) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bone_name))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_find_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the name of the bone at [param bone_idx] that will be the key name in the [BoneMap].
In the retargeting process, the returned bone name is the bone name of the target skeleton.
*/
//go:nosplit
func (self class) GetBoneName(ctx gd.Lifetime, bone_idx gd.Int) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_get_bone_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the name of the bone at [param bone_idx] that will be the key name in the [BoneMap].
In the retargeting process, the setting bone name is the bone name of the target skeleton.
*/
//go:nosplit
func (self class) SetBoneName(bone_idx gd.Int, bone_name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, mmm.Get(bone_name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_set_bone_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the name of the bone which is the parent to the bone at [param bone_idx]. The result is empty if the bone has no parent.
*/
//go:nosplit
func (self class) GetBoneParent(ctx gd.Lifetime, bone_idx gd.Int) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_get_bone_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the bone with name [param bone_parent] as the parent of the bone at [param bone_idx]. If an empty string is passed, then the bone has no parent.
*/
//go:nosplit
func (self class) SetBoneParent(bone_idx gd.Int, bone_parent gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, mmm.Get(bone_parent))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_set_bone_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the tail direction of the bone at [param bone_idx].
*/
//go:nosplit
func (self class) GetTailDirection(bone_idx gd.Int) classdb.SkeletonProfileTailDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[classdb.SkeletonProfileTailDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_get_tail_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the tail direction of the bone at [param bone_idx].
[b]Note:[/b] This only specifies the method of calculation. The actual coordinates required should be stored in an external skeleton, so the calculation itself needs to be done externally.
*/
//go:nosplit
func (self class) SetTailDirection(bone_idx gd.Int, tail_direction classdb.SkeletonProfileTailDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, tail_direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_set_tail_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the name of the bone which is the tail of the bone at [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneTail(ctx gd.Lifetime, bone_idx gd.Int) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_get_bone_tail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the bone with name [param bone_tail] as the tail of the bone at [param bone_idx].
*/
//go:nosplit
func (self class) SetBoneTail(bone_idx gd.Int, bone_tail gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, mmm.Get(bone_tail))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_set_bone_tail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the reference pose transform for bone [param bone_idx].
*/
//go:nosplit
func (self class) GetReferencePose(bone_idx gd.Int) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_get_reference_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the reference pose transform for bone [param bone_idx].
*/
//go:nosplit
func (self class) SetReferencePose(bone_idx gd.Int, bone_name gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, bone_name)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_set_reference_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the offset of the bone at [param bone_idx] that will be the button position in the [BoneMap] editor.
This is the offset with origin at the top left corner of the square.
*/
//go:nosplit
func (self class) GetHandleOffset(bone_idx gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_get_handle_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the offset of the bone at [param bone_idx] that will be the button position in the [BoneMap] editor.
This is the offset with origin at the top left corner of the square.
*/
//go:nosplit
func (self class) SetHandleOffset(bone_idx gd.Int, handle_offset gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, handle_offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_set_handle_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the group of the bone at [param bone_idx].
*/
//go:nosplit
func (self class) GetGroup(ctx gd.Lifetime, bone_idx gd.Int) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_get_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the group of the bone at [param bone_idx].
*/
//go:nosplit
func (self class) SetGroup(bone_idx gd.Int, group gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, mmm.Get(group))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_set_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the bone at [param bone_idx] is required for retargeting.
This value is used by the bone map editor. If this method returns [code]true[/code], and no bone is assigned, the handle color will be red on the bone map editor.
*/
//go:nosplit
func (self class) IsRequired(bone_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_is_required, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the required status for bone [param bone_idx] to [param required].
*/
//go:nosplit
func (self class) SetRequired(bone_idx gd.Int, required bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, required)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonProfile.Bind_set_required, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsSkeletonProfile() Expert { return self[0].AsSkeletonProfile() }


//go:nosplit
func (self Simple) AsSkeletonProfile() Simple { return self[0].AsSkeletonProfile() }


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
func init() {classdb.Register("SkeletonProfile", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type TailDirection = classdb.SkeletonProfileTailDirection

const (
/*Direction to the average coordinates of bone children.*/
	TailDirectionAverageChildren TailDirection = 0
/*Direction to the coordinates of specified bone child.*/
	TailDirectionSpecificChild TailDirection = 1
/*Direction is not calculated.*/
	TailDirectionEnd TailDirection = 2
)
