package Skeleton3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[Skeleton3D] provides an interface for managing a hierarchy of bones, including pose, rest and animation (see [Animation]). It can also use ragdoll physics.
The overall transform of a bone with respect to the skeleton is determined by bone pose. Bone rest defines the initial transform of the bone pose.
Note that "global pose" below refers to the overall transform of the bone with respect to skeleton, so it is not the actual global/world transform of the bone.

*/
type Simple [1]classdb.Skeleton3D
func (self Simple) AddBone(name string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AddBone(gc.String(name))))
}
func (self Simple) FindBone(name string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).FindBone(gc.String(name))))
}
func (self Simple) GetBoneName(bone_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetBoneName(gc, gd.Int(bone_idx)).String())
}
func (self Simple) SetBoneName(bone_idx int, name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBoneName(gd.Int(bone_idx), gc.String(name))
}
func (self Simple) GetConcatenatedBoneNames() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetConcatenatedBoneNames(gc).String())
}
func (self Simple) GetBoneParent(bone_idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBoneParent(gd.Int(bone_idx))))
}
func (self Simple) SetBoneParent(bone_idx int, parent_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBoneParent(gd.Int(bone_idx), gd.Int(parent_idx))
}
func (self Simple) GetBoneCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBoneCount()))
}
func (self Simple) GetVersion() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetVersion()))
}
func (self Simple) UnparentBoneAndRest(bone_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UnparentBoneAndRest(gd.Int(bone_idx))
}
func (self Simple) GetBoneChildren(bone_idx int) gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).GetBoneChildren(gc, gd.Int(bone_idx)))
}
func (self Simple) GetParentlessBones() gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).GetParentlessBones(gc))
}
func (self Simple) GetBoneRest(bone_idx int) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).GetBoneRest(gd.Int(bone_idx)))
}
func (self Simple) SetBoneRest(bone_idx int, rest gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBoneRest(gd.Int(bone_idx), rest)
}
func (self Simple) GetBoneGlobalRest(bone_idx int) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).GetBoneGlobalRest(gd.Int(bone_idx)))
}
func (self Simple) CreateSkinFromRestTransforms() [1]classdb.Skin {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Skin(Expert(self).CreateSkinFromRestTransforms(gc))
}
func (self Simple) RegisterSkin(skin [1]classdb.Skin) [1]classdb.SkinReference {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.SkinReference(Expert(self).RegisterSkin(gc, skin))
}
func (self Simple) LocalizeRests() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LocalizeRests()
}
func (self Simple) ClearBones() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearBones()
}
func (self Simple) GetBonePose(bone_idx int) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).GetBonePose(gd.Int(bone_idx)))
}
func (self Simple) SetBonePose(bone_idx int, pose gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBonePose(gd.Int(bone_idx), pose)
}
func (self Simple) SetBonePosePosition(bone_idx int, position gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBonePosePosition(gd.Int(bone_idx), position)
}
func (self Simple) SetBonePoseRotation(bone_idx int, rotation gd.Quaternion) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBonePoseRotation(gd.Int(bone_idx), rotation)
}
func (self Simple) SetBonePoseScale(bone_idx int, scale gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBonePoseScale(gd.Int(bone_idx), scale)
}
func (self Simple) GetBonePosePosition(bone_idx int) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetBonePosePosition(gd.Int(bone_idx)))
}
func (self Simple) GetBonePoseRotation(bone_idx int) gd.Quaternion {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Quaternion(Expert(self).GetBonePoseRotation(gd.Int(bone_idx)))
}
func (self Simple) GetBonePoseScale(bone_idx int) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetBonePoseScale(gd.Int(bone_idx)))
}
func (self Simple) ResetBonePose(bone_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ResetBonePose(gd.Int(bone_idx))
}
func (self Simple) ResetBonePoses() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ResetBonePoses()
}
func (self Simple) IsBoneEnabled(bone_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsBoneEnabled(gd.Int(bone_idx)))
}
func (self Simple) SetBoneEnabled(bone_idx int, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBoneEnabled(gd.Int(bone_idx), enabled)
}
func (self Simple) GetBoneGlobalPose(bone_idx int) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).GetBoneGlobalPose(gd.Int(bone_idx)))
}
func (self Simple) SetBoneGlobalPose(bone_idx int, pose gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBoneGlobalPose(gd.Int(bone_idx), pose)
}
func (self Simple) ForceUpdateAllBoneTransforms() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ForceUpdateAllBoneTransforms()
}
func (self Simple) ForceUpdateBoneChildTransform(bone_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ForceUpdateBoneChildTransform(gd.Int(bone_idx))
}
func (self Simple) SetMotionScale(motion_scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMotionScale(gd.Float(motion_scale))
}
func (self Simple) GetMotionScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMotionScale()))
}
func (self Simple) SetShowRestOnly(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShowRestOnly(enabled)
}
func (self Simple) IsShowRestOnly() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsShowRestOnly())
}
func (self Simple) SetModifierCallbackModeProcess(mode classdb.Skeleton3DModifierCallbackModeProcess) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetModifierCallbackModeProcess(mode)
}
func (self Simple) GetModifierCallbackModeProcess() classdb.Skeleton3DModifierCallbackModeProcess {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.Skeleton3DModifierCallbackModeProcess(Expert(self).GetModifierCallbackModeProcess())
}
func (self Simple) ClearBonesGlobalPoseOverride() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearBonesGlobalPoseOverride()
}
func (self Simple) SetBoneGlobalPoseOverride(bone_idx int, pose gd.Transform3D, amount float64, persistent bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBoneGlobalPoseOverride(gd.Int(bone_idx), pose, gd.Float(amount), persistent)
}
func (self Simple) GetBoneGlobalPoseOverride(bone_idx int) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).GetBoneGlobalPoseOverride(gd.Int(bone_idx)))
}
func (self Simple) GetBoneGlobalPoseNoOverride(bone_idx int) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).GetBoneGlobalPoseNoOverride(gd.Int(bone_idx)))
}
func (self Simple) SetAnimatePhysicalBones(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAnimatePhysicalBones(enabled)
}
func (self Simple) GetAnimatePhysicalBones() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetAnimatePhysicalBones())
}
func (self Simple) PhysicalBonesStopSimulation() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PhysicalBonesStopSimulation()
}
func (self Simple) PhysicalBonesStartSimulation(bones gd.ArrayOf[gd.StringName]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PhysicalBonesStartSimulation(bones)
}
func (self Simple) PhysicalBonesAddCollisionException(exception gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PhysicalBonesAddCollisionException(exception)
}
func (self Simple) PhysicalBonesRemoveCollisionException(exception gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PhysicalBonesRemoveCollisionException(exception)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Skeleton3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Adds a new bone with the given name. Returns the new bone's index, or [code]-1[/code] if this method fails.
[b]Note:[/b] Bone names should be unique, non empty, and cannot include the [code]:[/code] and [code]/[/code] characters.
*/
//go:nosplit
func (self class) AddBone(name gd.String) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_add_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the bone index that matches [param name] as its name. Returns [code]-1[/code] if no bone with this name exists.
*/
//go:nosplit
func (self class) FindBone(name gd.String) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_find_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the name of the bone at index [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneName(ctx gd.Lifetime, bone_idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_get_bone_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the bone name, [param name], for the bone at [param bone_idx].
*/
//go:nosplit
func (self class) SetBoneName(bone_idx gd.Int, name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_set_bone_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns all bone names concatenated with commas ([code],[/code]) as a single [StringName].
It is useful to set it as a hint for the enum property.
*/
//go:nosplit
func (self class) GetConcatenatedBoneNames(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_get_concatenated_bone_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the bone index which is the parent of the bone at [param bone_idx]. If -1, then bone has no parent.
[b]Note:[/b] The parent bone returned will always be less than [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneParent(bone_idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_get_bone_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the bone index [param parent_idx] as the parent of the bone at [param bone_idx]. If -1, then bone has no parent.
[b]Note:[/b] [param parent_idx] must be less than [param bone_idx].
*/
//go:nosplit
func (self class) SetBoneParent(bone_idx gd.Int, parent_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, parent_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_set_bone_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of bones in the skeleton.
*/
//go:nosplit
func (self class) GetBoneCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_get_bone_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of times the bone hierarchy has changed within this skeleton, including renames.
The Skeleton version is not serialized: only use within a single instance of Skeleton3D.
Use for invalidating caches in IK solvers and other nodes which process bones.
*/
//go:nosplit
func (self class) GetVersion() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_get_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Unparents the bone at [param bone_idx] and sets its rest position to that of its parent prior to being reset.
*/
//go:nosplit
func (self class) UnparentBoneAndRest(bone_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_unparent_bone_and_rest, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an array containing the bone indexes of all the child node of the passed in bone, [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneChildren(ctx gd.Lifetime, bone_idx gd.Int) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_get_bone_children, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array with all of the bones that are parentless. Another way to look at this is that it returns the indexes of all the bones that are not dependent or modified by other bones in the Skeleton.
*/
//go:nosplit
func (self class) GetParentlessBones(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_get_parentless_bones, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the rest transform for a bone [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneRest(bone_idx gd.Int) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_get_bone_rest, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the rest transform for bone [param bone_idx].
*/
//go:nosplit
func (self class) SetBoneRest(bone_idx gd.Int, rest gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, rest)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_set_bone_rest, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the global rest transform for [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneGlobalRest(bone_idx gd.Int) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_get_bone_global_rest, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) CreateSkinFromRestTransforms(ctx gd.Lifetime) object.Skin {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_create_skin_from_rest_transforms, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Skin
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Binds the given Skin to the Skeleton.
*/
//go:nosplit
func (self class) RegisterSkin(ctx gd.Lifetime, skin object.Skin) object.SkinReference {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(skin[0].AsPointer())[0])
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_register_skin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.SkinReference
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns all bones in the skeleton to their rest poses.
*/
//go:nosplit
func (self class) LocalizeRests()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_localize_rests, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clear all the bones in this skeleton.
*/
//go:nosplit
func (self class) ClearBones()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_clear_bones, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the pose transform of the specified bone.
[b]Note:[/b] This is the pose you set to the skeleton in the process, the final pose can get overridden by modifiers in the deferred process, if you want to access the final pose, use [signal SkeletonModifier3D.modification_processed].
*/
//go:nosplit
func (self class) GetBonePose(bone_idx gd.Int) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_get_bone_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the pose transform, [param pose], for the bone at [param bone_idx].
*/
//go:nosplit
func (self class) SetBonePose(bone_idx gd.Int, pose gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, pose)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_set_bone_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the pose position of the bone at [param bone_idx] to [param position]. [param position] is a [Vector3] describing a position local to the [Skeleton3D] node.
*/
//go:nosplit
func (self class) SetBonePosePosition(bone_idx gd.Int, position gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_set_bone_pose_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the pose rotation of the bone at [param bone_idx] to [param rotation]. [param rotation] is a [Quaternion] describing a rotation in the bone's local coordinate space with respect to the rotation of any parent bones.
*/
//go:nosplit
func (self class) SetBonePoseRotation(bone_idx gd.Int, rotation gd.Quaternion)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, rotation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_set_bone_pose_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the pose scale of the bone at [param bone_idx] to [param scale].
*/
//go:nosplit
func (self class) SetBonePoseScale(bone_idx gd.Int, scale gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_set_bone_pose_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the pose position of the bone at [param bone_idx]. The returned [Vector3] is in the local coordinate space of the [Skeleton3D] node.
*/
//go:nosplit
func (self class) GetBonePosePosition(bone_idx gd.Int) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_get_bone_pose_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the pose rotation of the bone at [param bone_idx]. The returned [Quaternion] is local to the bone with respect to the rotation of any parent bones.
*/
//go:nosplit
func (self class) GetBonePoseRotation(bone_idx gd.Int) gd.Quaternion {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Quaternion](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_get_bone_pose_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the pose scale of the bone at [param bone_idx].
*/
//go:nosplit
func (self class) GetBonePoseScale(bone_idx gd.Int) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_get_bone_pose_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the bone pose to rest for [param bone_idx].
*/
//go:nosplit
func (self class) ResetBonePose(bone_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_reset_bone_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets all bone poses to rests.
*/
//go:nosplit
func (self class) ResetBonePoses()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_reset_bone_poses, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the bone pose for the bone at [param bone_idx] is enabled.
*/
//go:nosplit
func (self class) IsBoneEnabled(bone_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_is_bone_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Disables the pose for the bone at [param bone_idx] if [code]false[/code], enables the bone pose if [code]true[/code].
*/
//go:nosplit
func (self class) SetBoneEnabled(bone_idx gd.Int, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_set_bone_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the overall transform of the specified bone, with respect to the skeleton. Being relative to the skeleton frame, this is not the actual "global" transform of the bone.
[b]Note:[/b] This is the global pose you set to the skeleton in the process, the final global pose can get overridden by modifiers in the deferred process, if you want to access the final global pose, use [signal SkeletonModifier3D.modification_processed].
*/
//go:nosplit
func (self class) GetBoneGlobalPose(bone_idx gd.Int) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_get_bone_global_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the global pose transform, [param pose], for the bone at [param bone_idx].
[b]Note:[/b] If other bone poses have been changed, this method executes a dirty poses recalculation and will cause performance to deteriorate. If you know that multiple global poses will be applied, consider using [method set_bone_pose] with precalculation.
*/
//go:nosplit
func (self class) SetBoneGlobalPose(bone_idx gd.Int, pose gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, pose)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_set_bone_global_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Force updates the bone transforms/poses for all bones in the skeleton.
*/
//go:nosplit
func (self class) ForceUpdateAllBoneTransforms()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_force_update_all_bone_transforms, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Force updates the bone transform for the bone at [param bone_idx] and all of its children.
*/
//go:nosplit
func (self class) ForceUpdateBoneChildTransform(bone_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_force_update_bone_child_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetMotionScale(motion_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, motion_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_set_motion_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMotionScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_get_motion_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowRestOnly(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_set_show_rest_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShowRestOnly() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_is_show_rest_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetModifierCallbackModeProcess(mode classdb.Skeleton3DModifierCallbackModeProcess)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_set_modifier_callback_mode_process, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetModifierCallbackModeProcess() classdb.Skeleton3DModifierCallbackModeProcess {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Skeleton3DModifierCallbackModeProcess](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_get_modifier_callback_mode_process, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the global pose override on all bones in the skeleton.
*/
//go:nosplit
func (self class) ClearBonesGlobalPoseOverride()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_clear_bones_global_pose_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the global pose transform, [param pose], for the bone at [param bone_idx].
[param amount] is the interpolation strength that will be used when applying the pose, and [param persistent] determines if the applied pose will remain.
[b]Note:[/b] The pose transform needs to be a global pose! To convert a world transform from a [Node3D] to a global bone pose, multiply the [method Transform3D.affine_inverse] of the node's [member Node3D.global_transform] by the desired world transform.
*/
//go:nosplit
func (self class) SetBoneGlobalPoseOverride(bone_idx gd.Int, pose gd.Transform3D, amount gd.Float, persistent bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, pose)
	callframe.Arg(frame, amount)
	callframe.Arg(frame, persistent)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_set_bone_global_pose_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the global pose override transform for [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneGlobalPoseOverride(bone_idx gd.Int) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_get_bone_global_pose_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the overall transform of the specified bone, with respect to the skeleton, but without any global pose overrides. Being relative to the skeleton frame, this is not the actual "global" transform of the bone.
*/
//go:nosplit
func (self class) GetBoneGlobalPoseNoOverride(bone_idx gd.Int) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_get_bone_global_pose_no_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAnimatePhysicalBones(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_set_animate_physical_bones, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAnimatePhysicalBones() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_get_animate_physical_bones, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Tells the [PhysicalBone3D] nodes in the Skeleton to stop simulating.
*/
//go:nosplit
func (self class) PhysicalBonesStopSimulation()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_physical_bones_stop_simulation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Tells the [PhysicalBone3D] nodes in the Skeleton to start simulating and reacting to the physics world.
Optionally, a list of bone names can be passed-in, allowing only the passed-in bones to be simulated.
*/
//go:nosplit
func (self class) PhysicalBonesStartSimulation(bones gd.ArrayOf[gd.StringName])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bones))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_physical_bones_start_simulation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
//go:nosplit
func (self class) PhysicalBonesAddCollisionException(exception gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, exception)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_physical_bones_add_collision_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
//go:nosplit
func (self class) PhysicalBonesRemoveCollisionException(exception gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, exception)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton3D.Bind_physical_bones_remove_collision_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsSkeleton3D() Expert { return self[0].AsSkeleton3D() }


//go:nosplit
func (self Simple) AsSkeleton3D() Simple { return self[0].AsSkeleton3D() }


//go:nosplit
func (self class) AsNode3D() Node3D.Expert { return self[0].AsNode3D() }


//go:nosplit
func (self Simple) AsNode3D() Node3D.Simple { return self[0].AsNode3D() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

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
func init() {classdb.Register("Skeleton3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type ModifierCallbackModeProcess = classdb.Skeleton3DModifierCallbackModeProcess

const (
/*Set a flag to process modification during physics frames (see [constant Node.NOTIFICATION_INTERNAL_PHYSICS_PROCESS]).*/
	ModifierCallbackModeProcessPhysics ModifierCallbackModeProcess = 0
/*Set a flag to process modification during process frames (see [constant Node.NOTIFICATION_INTERNAL_PROCESS]).*/
	ModifierCallbackModeProcessIdle ModifierCallbackModeProcess = 1
)
