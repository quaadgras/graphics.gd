// Package Skeleton3D provides methods for working with Skeleton3D object instances.
package Skeleton3D

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Quaternion"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Vector3"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
[Skeleton3D] provides an interface for managing a hierarchy of bones, including pose, rest and animation (see [Animation]). It can also use ragdoll physics.
The overall transform of a bone with respect to the skeleton is determined by bone pose. Bone rest defines the initial transform of the bone pose.
Note that "global pose" below refers to the overall transform of the bone with respect to skeleton, so it is not the actual global/world transform of the bone.
*/
type Instance [1]gdclass.Skeleton3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSkeleton3D() Instance
}

/*
Adds a new bone with the given name. Returns the new bone's index, or [code]-1[/code] if this method fails.
[b]Note:[/b] Bone names should be unique, non empty, and cannot include the [code]:[/code] and [code]/[/code] characters.
*/
func (self Instance) AddBone(name string) int { //gd:Skeleton3D.add_bone
	return int(int(class(self).AddBone(String.New(name))))
}

/*
Returns the bone index that matches [param name] as its name. Returns [code]-1[/code] if no bone with this name exists.
*/
func (self Instance) FindBone(name string) int { //gd:Skeleton3D.find_bone
	return int(int(class(self).FindBone(String.New(name))))
}

/*
Returns the name of the bone at index [param bone_idx].
*/
func (self Instance) GetBoneName(bone_idx int) string { //gd:Skeleton3D.get_bone_name
	return string(class(self).GetBoneName(int64(bone_idx)).String())
}

/*
Sets the bone name, [param name], for the bone at [param bone_idx].
*/
func (self Instance) SetBoneName(bone_idx int, name string) { //gd:Skeleton3D.set_bone_name
	class(self).SetBoneName(int64(bone_idx), String.New(name))
}

/*
Returns bone metadata for [param bone_idx] with [param key].
*/
func (self Instance) GetBoneMeta(bone_idx int, key string) any { //gd:Skeleton3D.get_bone_meta
	return any(class(self).GetBoneMeta(int64(bone_idx), String.Name(String.New(key))).Interface())
}

/*
Returns a list of all metadata keys for [param bone_idx].
*/
func (self Instance) GetBoneMetaList(bone_idx int) []string { //gd:Skeleton3D.get_bone_meta_list
	return []string(gd.ArrayAs[[]string](gd.InternalArray(class(self).GetBoneMetaList(int64(bone_idx)))))
}

/*
Returns whether there exists any bone metadata for [param bone_idx] with key [param key].
*/
func (self Instance) HasBoneMeta(bone_idx int, key string) bool { //gd:Skeleton3D.has_bone_meta
	return bool(class(self).HasBoneMeta(int64(bone_idx), String.Name(String.New(key))))
}

/*
Sets bone metadata for [param bone_idx], will set the [param key] meta to [param value].
*/
func (self Instance) SetBoneMeta(bone_idx int, key string, value any) { //gd:Skeleton3D.set_bone_meta
	class(self).SetBoneMeta(int64(bone_idx), String.Name(String.New(key)), variant.New(value))
}

/*
Returns all bone names concatenated with commas ([code],[/code]) as a single [StringName].
It is useful to set it as a hint for the enum property.
*/
func (self Instance) GetConcatenatedBoneNames() string { //gd:Skeleton3D.get_concatenated_bone_names
	return string(class(self).GetConcatenatedBoneNames().String())
}

/*
Returns the bone index which is the parent of the bone at [param bone_idx]. If -1, then bone has no parent.
[b]Note:[/b] The parent bone returned will always be less than [param bone_idx].
*/
func (self Instance) GetBoneParent(bone_idx int) int { //gd:Skeleton3D.get_bone_parent
	return int(int(class(self).GetBoneParent(int64(bone_idx))))
}

/*
Sets the bone index [param parent_idx] as the parent of the bone at [param bone_idx]. If -1, then bone has no parent.
[b]Note:[/b] [param parent_idx] must be less than [param bone_idx].
*/
func (self Instance) SetBoneParent(bone_idx int, parent_idx int) { //gd:Skeleton3D.set_bone_parent
	class(self).SetBoneParent(int64(bone_idx), int64(parent_idx))
}

/*
Returns the number of bones in the skeleton.
*/
func (self Instance) GetBoneCount() int { //gd:Skeleton3D.get_bone_count
	return int(int(class(self).GetBoneCount()))
}

/*
Returns the number of times the bone hierarchy has changed within this skeleton, including renames.
The Skeleton version is not serialized: only use within a single instance of Skeleton3D.
Use for invalidating caches in IK solvers and other nodes which process bones.
*/
func (self Instance) GetVersion() int { //gd:Skeleton3D.get_version
	return int(int(class(self).GetVersion()))
}

/*
Unparents the bone at [param bone_idx] and sets its rest position to that of its parent prior to being reset.
*/
func (self Instance) UnparentBoneAndRest(bone_idx int) { //gd:Skeleton3D.unparent_bone_and_rest
	class(self).UnparentBoneAndRest(int64(bone_idx))
}

/*
Returns an array containing the bone indexes of all the child node of the passed in bone, [param bone_idx].
*/
func (self Instance) GetBoneChildren(bone_idx int) []int32 { //gd:Skeleton3D.get_bone_children
	return []int32(slices.Collect(class(self).GetBoneChildren(int64(bone_idx)).Values()))
}

/*
Returns an array with all of the bones that are parentless. Another way to look at this is that it returns the indexes of all the bones that are not dependent or modified by other bones in the Skeleton.
*/
func (self Instance) GetParentlessBones() []int32 { //gd:Skeleton3D.get_parentless_bones
	return []int32(slices.Collect(class(self).GetParentlessBones().Values()))
}

/*
Returns the rest transform for a bone [param bone_idx].
*/
func (self Instance) GetBoneRest(bone_idx int) Transform3D.BasisOrigin { //gd:Skeleton3D.get_bone_rest
	return Transform3D.BasisOrigin(class(self).GetBoneRest(int64(bone_idx)))
}

/*
Sets the rest transform for bone [param bone_idx].
*/
func (self Instance) SetBoneRest(bone_idx int, rest Transform3D.BasisOrigin) { //gd:Skeleton3D.set_bone_rest
	class(self).SetBoneRest(int64(bone_idx), Transform3D.BasisOrigin(rest))
}

/*
Returns the global rest transform for [param bone_idx].
*/
func (self Instance) GetBoneGlobalRest(bone_idx int) Transform3D.BasisOrigin { //gd:Skeleton3D.get_bone_global_rest
	return Transform3D.BasisOrigin(class(self).GetBoneGlobalRest(int64(bone_idx)))
}
func (self Instance) CreateSkinFromRestTransforms() [1]gdclass.Skin { //gd:Skeleton3D.create_skin_from_rest_transforms
	return [1]gdclass.Skin(class(self).CreateSkinFromRestTransforms())
}

/*
Binds the given Skin to the Skeleton.
*/
func (self Instance) RegisterSkin(skin [1]gdclass.Skin) [1]gdclass.SkinReference { //gd:Skeleton3D.register_skin
	return [1]gdclass.SkinReference(class(self).RegisterSkin(skin))
}

/*
Returns all bones in the skeleton to their rest poses.
*/
func (self Instance) LocalizeRests() { //gd:Skeleton3D.localize_rests
	class(self).LocalizeRests()
}

/*
Clear all the bones in this skeleton.
*/
func (self Instance) ClearBones() { //gd:Skeleton3D.clear_bones
	class(self).ClearBones()
}

/*
Returns the pose transform of the specified bone.
[b]Note:[/b] This is the pose you set to the skeleton in the process, the final pose can get overridden by modifiers in the deferred process, if you want to access the final pose, use [signal SkeletonModifier3D.modification_processed].
*/
func (self Instance) GetBonePose(bone_idx int) Transform3D.BasisOrigin { //gd:Skeleton3D.get_bone_pose
	return Transform3D.BasisOrigin(class(self).GetBonePose(int64(bone_idx)))
}

/*
Sets the pose transform, [param pose], for the bone at [param bone_idx].
*/
func (self Instance) SetBonePose(bone_idx int, pose Transform3D.BasisOrigin) { //gd:Skeleton3D.set_bone_pose
	class(self).SetBonePose(int64(bone_idx), Transform3D.BasisOrigin(pose))
}

/*
Sets the pose position of the bone at [param bone_idx] to [param position]. [param position] is a [Vector3] describing a position local to the [Skeleton3D] node.
*/
func (self Instance) SetBonePosePosition(bone_idx int, position Vector3.XYZ) { //gd:Skeleton3D.set_bone_pose_position
	class(self).SetBonePosePosition(int64(bone_idx), Vector3.XYZ(position))
}

/*
Sets the pose rotation of the bone at [param bone_idx] to [param rotation]. [param rotation] is a [Quaternion] describing a rotation in the bone's local coordinate space with respect to the rotation of any parent bones.
*/
func (self Instance) SetBonePoseRotation(bone_idx int, rotation Quaternion.IJKX) { //gd:Skeleton3D.set_bone_pose_rotation
	class(self).SetBonePoseRotation(int64(bone_idx), rotation)
}

/*
Sets the pose scale of the bone at [param bone_idx] to [param scale].
*/
func (self Instance) SetBonePoseScale(bone_idx int, scale Vector3.XYZ) { //gd:Skeleton3D.set_bone_pose_scale
	class(self).SetBonePoseScale(int64(bone_idx), Vector3.XYZ(scale))
}

/*
Returns the pose position of the bone at [param bone_idx]. The returned [Vector3] is in the local coordinate space of the [Skeleton3D] node.
*/
func (self Instance) GetBonePosePosition(bone_idx int) Vector3.XYZ { //gd:Skeleton3D.get_bone_pose_position
	return Vector3.XYZ(class(self).GetBonePosePosition(int64(bone_idx)))
}

/*
Returns the pose rotation of the bone at [param bone_idx]. The returned [Quaternion] is local to the bone with respect to the rotation of any parent bones.
*/
func (self Instance) GetBonePoseRotation(bone_idx int) Quaternion.IJKX { //gd:Skeleton3D.get_bone_pose_rotation
	return Quaternion.IJKX(class(self).GetBonePoseRotation(int64(bone_idx)))
}

/*
Returns the pose scale of the bone at [param bone_idx].
*/
func (self Instance) GetBonePoseScale(bone_idx int) Vector3.XYZ { //gd:Skeleton3D.get_bone_pose_scale
	return Vector3.XYZ(class(self).GetBonePoseScale(int64(bone_idx)))
}

/*
Sets the bone pose to rest for [param bone_idx].
*/
func (self Instance) ResetBonePose(bone_idx int) { //gd:Skeleton3D.reset_bone_pose
	class(self).ResetBonePose(int64(bone_idx))
}

/*
Sets all bone poses to rests.
*/
func (self Instance) ResetBonePoses() { //gd:Skeleton3D.reset_bone_poses
	class(self).ResetBonePoses()
}

/*
Returns whether the bone pose for the bone at [param bone_idx] is enabled.
*/
func (self Instance) IsBoneEnabled(bone_idx int) bool { //gd:Skeleton3D.is_bone_enabled
	return bool(class(self).IsBoneEnabled(int64(bone_idx)))
}

/*
Disables the pose for the bone at [param bone_idx] if [code]false[/code], enables the bone pose if [code]true[/code].
*/
func (self Instance) SetBoneEnabled(bone_idx int) { //gd:Skeleton3D.set_bone_enabled
	class(self).SetBoneEnabled(int64(bone_idx), true)
}

/*
Returns the overall transform of the specified bone, with respect to the skeleton. Being relative to the skeleton frame, this is not the actual "global" transform of the bone.
[b]Note:[/b] This is the global pose you set to the skeleton in the process, the final global pose can get overridden by modifiers in the deferred process, if you want to access the final global pose, use [signal SkeletonModifier3D.modification_processed].
*/
func (self Instance) GetBoneGlobalPose(bone_idx int) Transform3D.BasisOrigin { //gd:Skeleton3D.get_bone_global_pose
	return Transform3D.BasisOrigin(class(self).GetBoneGlobalPose(int64(bone_idx)))
}

/*
Sets the global pose transform, [param pose], for the bone at [param bone_idx].
[b]Note:[/b] If other bone poses have been changed, this method executes a dirty poses recalculation and will cause performance to deteriorate. If you know that multiple global poses will be applied, consider using [method set_bone_pose] with precalculation.
*/
func (self Instance) SetBoneGlobalPose(bone_idx int, pose Transform3D.BasisOrigin) { //gd:Skeleton3D.set_bone_global_pose
	class(self).SetBoneGlobalPose(int64(bone_idx), Transform3D.BasisOrigin(pose))
}

/*
Force updates the bone transforms/poses for all bones in the skeleton.
*/
func (self Instance) ForceUpdateAllBoneTransforms() { //gd:Skeleton3D.force_update_all_bone_transforms
	class(self).ForceUpdateAllBoneTransforms()
}

/*
Force updates the bone transform for the bone at [param bone_idx] and all of its children.
*/
func (self Instance) ForceUpdateBoneChildTransform(bone_idx int) { //gd:Skeleton3D.force_update_bone_child_transform
	class(self).ForceUpdateBoneChildTransform(int64(bone_idx))
}

/*
Removes the global pose override on all bones in the skeleton.
*/
func (self Instance) ClearBonesGlobalPoseOverride() { //gd:Skeleton3D.clear_bones_global_pose_override
	class(self).ClearBonesGlobalPoseOverride()
}

/*
Sets the global pose transform, [param pose], for the bone at [param bone_idx].
[param amount] is the interpolation strength that will be used when applying the pose, and [param persistent] determines if the applied pose will remain.
[b]Note:[/b] The pose transform needs to be a global pose! To convert a world transform from a [Node3D] to a global bone pose, multiply the [method Transform3D.affine_inverse] of the node's [member Node3D.global_transform] by the desired world transform.
*/
func (self Instance) SetBoneGlobalPoseOverride(bone_idx int, pose Transform3D.BasisOrigin, amount Float.X) { //gd:Skeleton3D.set_bone_global_pose_override
	class(self).SetBoneGlobalPoseOverride(int64(bone_idx), Transform3D.BasisOrigin(pose), float64(amount), false)
}

/*
Returns the global pose override transform for [param bone_idx].
*/
func (self Instance) GetBoneGlobalPoseOverride(bone_idx int) Transform3D.BasisOrigin { //gd:Skeleton3D.get_bone_global_pose_override
	return Transform3D.BasisOrigin(class(self).GetBoneGlobalPoseOverride(int64(bone_idx)))
}

/*
Returns the overall transform of the specified bone, with respect to the skeleton, but without any global pose overrides. Being relative to the skeleton frame, this is not the actual "global" transform of the bone.
*/
func (self Instance) GetBoneGlobalPoseNoOverride(bone_idx int) Transform3D.BasisOrigin { //gd:Skeleton3D.get_bone_global_pose_no_override
	return Transform3D.BasisOrigin(class(self).GetBoneGlobalPoseNoOverride(int64(bone_idx)))
}

/*
Tells the [PhysicalBone3D] nodes in the Skeleton to stop simulating.
*/
func (self Instance) PhysicalBonesStopSimulation() { //gd:Skeleton3D.physical_bones_stop_simulation
	class(self).PhysicalBonesStopSimulation()
}

/*
Tells the [PhysicalBone3D] nodes in the Skeleton to start simulating and reacting to the physics world.
Optionally, a list of bone names can be passed-in, allowing only the passed-in bones to be simulated.
*/
func (self Instance) PhysicalBonesStartSimulation() { //gd:Skeleton3D.physical_bones_start_simulation
	class(self).PhysicalBonesStartSimulation(gd.ArrayFromSlice[Array.Contains[String.Name]]([1][]string{}[0]))
}

/*
Adds a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
func (self Instance) PhysicalBonesAddCollisionException(exception RID.Body3D) { //gd:Skeleton3D.physical_bones_add_collision_exception
	class(self).PhysicalBonesAddCollisionException(RID.Any(exception))
}

/*
Removes a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
func (self Instance) PhysicalBonesRemoveCollisionException(exception RID.Body3D) { //gd:Skeleton3D.physical_bones_remove_collision_exception
	class(self).PhysicalBonesRemoveCollisionException(RID.Any(exception))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Skeleton3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Skeleton3D"))
	casted := Instance{*(*gdclass.Skeleton3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) MotionScale() Float.X {
	return Float.X(Float.X(class(self).GetMotionScale()))
}

func (self Instance) SetMotionScale(value Float.X) {
	class(self).SetMotionScale(float64(value))
}

func (self Instance) ShowRestOnly() bool {
	return bool(class(self).IsShowRestOnly())
}

func (self Instance) SetShowRestOnly(value bool) {
	class(self).SetShowRestOnly(value)
}

func (self Instance) ModifierCallbackModeProcess() gdclass.Skeleton3DModifierCallbackModeProcess {
	return gdclass.Skeleton3DModifierCallbackModeProcess(class(self).GetModifierCallbackModeProcess())
}

func (self Instance) SetModifierCallbackModeProcess(value gdclass.Skeleton3DModifierCallbackModeProcess) {
	class(self).SetModifierCallbackModeProcess(value)
}

func (self Instance) AnimatePhysicalBones() bool {
	return bool(class(self).GetAnimatePhysicalBones())
}

func (self Instance) SetAnimatePhysicalBones(value bool) {
	class(self).SetAnimatePhysicalBones(value)
}

/*
Adds a new bone with the given name. Returns the new bone's index, or [code]-1[/code] if this method fails.
[b]Note:[/b] Bone names should be unique, non empty, and cannot include the [code]:[/code] and [code]/[/code] characters.
*/
//go:nosplit
func (self class) AddBone(name String.Readable) int64 { //gd:Skeleton3D.add_bone
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_add_bone, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the bone index that matches [param name] as its name. Returns [code]-1[/code] if no bone with this name exists.
*/
//go:nosplit
func (self class) FindBone(name String.Readable) int64 { //gd:Skeleton3D.find_bone
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_find_bone, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the name of the bone at index [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneName(bone_idx int64) String.Readable { //gd:Skeleton3D.get_bone_name
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the bone name, [param name], for the bone at [param bone_idx].
*/
//go:nosplit
func (self class) SetBoneName(bone_idx int64, name String.Readable) { //gd:Skeleton3D.set_bone_name
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_bone_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns bone metadata for [param bone_idx] with [param key].
*/
//go:nosplit
func (self class) GetBoneMeta(bone_idx int64, key String.Name) variant.Any { //gd:Skeleton3D.get_bone_meta
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(key)))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_meta, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns a list of all metadata keys for [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneMetaList(bone_idx int64) Array.Contains[String.Name] { //gd:Skeleton3D.get_bone_meta_list
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_meta_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[String.Name]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns whether there exists any bone metadata for [param bone_idx] with key [param key].
*/
//go:nosplit
func (self class) HasBoneMeta(bone_idx int64, key String.Name) bool { //gd:Skeleton3D.has_bone_meta
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(key)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_has_bone_meta, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets bone metadata for [param bone_idx], will set the [param key] meta to [param value].
*/
//go:nosplit
func (self class) SetBoneMeta(bone_idx int64, key String.Name, value variant.Any) { //gd:Skeleton3D.set_bone_meta
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(key)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(value)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_bone_meta, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns all bone names concatenated with commas ([code],[/code]) as a single [StringName].
It is useful to set it as a hint for the enum property.
*/
//go:nosplit
func (self class) GetConcatenatedBoneNames() String.Name { //gd:Skeleton3D.get_concatenated_bone_names
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_concatenated_bone_names, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the bone index which is the parent of the bone at [param bone_idx]. If -1, then bone has no parent.
[b]Note:[/b] The parent bone returned will always be less than [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneParent(bone_idx int64) int64 { //gd:Skeleton3D.get_bone_parent
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_parent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the bone index [param parent_idx] as the parent of the bone at [param bone_idx]. If -1, then bone has no parent.
[b]Note:[/b] [param parent_idx] must be less than [param bone_idx].
*/
//go:nosplit
func (self class) SetBoneParent(bone_idx int64, parent_idx int64) { //gd:Skeleton3D.set_bone_parent
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, parent_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_bone_parent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the number of bones in the skeleton.
*/
//go:nosplit
func (self class) GetBoneCount() int64 { //gd:Skeleton3D.get_bone_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_count, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) GetVersion() int64 { //gd:Skeleton3D.get_version
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_version, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Unparents the bone at [param bone_idx] and sets its rest position to that of its parent prior to being reset.
*/
//go:nosplit
func (self class) UnparentBoneAndRest(bone_idx int64) { //gd:Skeleton3D.unparent_bone_and_rest
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_unparent_bone_and_rest, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an array containing the bone indexes of all the child node of the passed in bone, [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneChildren(bone_idx int64) Packed.Array[int32] { //gd:Skeleton3D.get_bone_children
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_children, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns an array with all of the bones that are parentless. Another way to look at this is that it returns the indexes of all the bones that are not dependent or modified by other bones in the Skeleton.
*/
//go:nosplit
func (self class) GetParentlessBones() Packed.Array[int32] { //gd:Skeleton3D.get_parentless_bones
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_parentless_bones, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the rest transform for a bone [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneRest(bone_idx int64) Transform3D.BasisOrigin { //gd:Skeleton3D.get_bone_rest
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_rest, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the rest transform for bone [param bone_idx].
*/
//go:nosplit
func (self class) SetBoneRest(bone_idx int64, rest Transform3D.BasisOrigin) { //gd:Skeleton3D.set_bone_rest
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, rest)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_bone_rest, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the global rest transform for [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneGlobalRest(bone_idx int64) Transform3D.BasisOrigin { //gd:Skeleton3D.get_bone_global_rest
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_global_rest, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) CreateSkinFromRestTransforms() [1]gdclass.Skin { //gd:Skeleton3D.create_skin_from_rest_transforms
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_create_skin_from_rest_transforms, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Skin{gd.PointerWithOwnershipTransferredToGo[gdclass.Skin](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Binds the given Skin to the Skeleton.
*/
//go:nosplit
func (self class) RegisterSkin(skin [1]gdclass.Skin) [1]gdclass.SkinReference { //gd:Skeleton3D.register_skin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(skin[0])[0])
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_register_skin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.SkinReference{gd.PointerWithOwnershipTransferredToGo[gdclass.SkinReference](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns all bones in the skeleton to their rest poses.
*/
//go:nosplit
func (self class) LocalizeRests() { //gd:Skeleton3D.localize_rests
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_localize_rests, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Clear all the bones in this skeleton.
*/
//go:nosplit
func (self class) ClearBones() { //gd:Skeleton3D.clear_bones
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_clear_bones, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the pose transform of the specified bone.
[b]Note:[/b] This is the pose you set to the skeleton in the process, the final pose can get overridden by modifiers in the deferred process, if you want to access the final pose, use [signal SkeletonModifier3D.modification_processed].
*/
//go:nosplit
func (self class) GetBonePose(bone_idx int64) Transform3D.BasisOrigin { //gd:Skeleton3D.get_bone_pose
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_pose, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the pose transform, [param pose], for the bone at [param bone_idx].
*/
//go:nosplit
func (self class) SetBonePose(bone_idx int64, pose Transform3D.BasisOrigin) { //gd:Skeleton3D.set_bone_pose
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, pose)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_bone_pose, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the pose position of the bone at [param bone_idx] to [param position]. [param position] is a [Vector3] describing a position local to the [Skeleton3D] node.
*/
//go:nosplit
func (self class) SetBonePosePosition(bone_idx int64, position Vector3.XYZ) { //gd:Skeleton3D.set_bone_pose_position
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_bone_pose_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the pose rotation of the bone at [param bone_idx] to [param rotation]. [param rotation] is a [Quaternion] describing a rotation in the bone's local coordinate space with respect to the rotation of any parent bones.
*/
//go:nosplit
func (self class) SetBonePoseRotation(bone_idx int64, rotation Quaternion.IJKX) { //gd:Skeleton3D.set_bone_pose_rotation
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, rotation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_bone_pose_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the pose scale of the bone at [param bone_idx] to [param scale].
*/
//go:nosplit
func (self class) SetBonePoseScale(bone_idx int64, scale Vector3.XYZ) { //gd:Skeleton3D.set_bone_pose_scale
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_bone_pose_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the pose position of the bone at [param bone_idx]. The returned [Vector3] is in the local coordinate space of the [Skeleton3D] node.
*/
//go:nosplit
func (self class) GetBonePosePosition(bone_idx int64) Vector3.XYZ { //gd:Skeleton3D.get_bone_pose_position
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_pose_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the pose rotation of the bone at [param bone_idx]. The returned [Quaternion] is local to the bone with respect to the rotation of any parent bones.
*/
//go:nosplit
func (self class) GetBonePoseRotation(bone_idx int64) Quaternion.IJKX { //gd:Skeleton3D.get_bone_pose_rotation
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[Quaternion.IJKX](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_pose_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the pose scale of the bone at [param bone_idx].
*/
//go:nosplit
func (self class) GetBonePoseScale(bone_idx int64) Vector3.XYZ { //gd:Skeleton3D.get_bone_pose_scale
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_pose_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the bone pose to rest for [param bone_idx].
*/
//go:nosplit
func (self class) ResetBonePose(bone_idx int64) { //gd:Skeleton3D.reset_bone_pose
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_reset_bone_pose, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets all bone poses to rests.
*/
//go:nosplit
func (self class) ResetBonePoses() { //gd:Skeleton3D.reset_bone_poses
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_reset_bone_poses, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether the bone pose for the bone at [param bone_idx] is enabled.
*/
//go:nosplit
func (self class) IsBoneEnabled(bone_idx int64) bool { //gd:Skeleton3D.is_bone_enabled
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_is_bone_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Disables the pose for the bone at [param bone_idx] if [code]false[/code], enables the bone pose if [code]true[/code].
*/
//go:nosplit
func (self class) SetBoneEnabled(bone_idx int64, enabled bool) { //gd:Skeleton3D.set_bone_enabled
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_bone_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the overall transform of the specified bone, with respect to the skeleton. Being relative to the skeleton frame, this is not the actual "global" transform of the bone.
[b]Note:[/b] This is the global pose you set to the skeleton in the process, the final global pose can get overridden by modifiers in the deferred process, if you want to access the final global pose, use [signal SkeletonModifier3D.modification_processed].
*/
//go:nosplit
func (self class) GetBoneGlobalPose(bone_idx int64) Transform3D.BasisOrigin { //gd:Skeleton3D.get_bone_global_pose
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_global_pose, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the global pose transform, [param pose], for the bone at [param bone_idx].
[b]Note:[/b] If other bone poses have been changed, this method executes a dirty poses recalculation and will cause performance to deteriorate. If you know that multiple global poses will be applied, consider using [method set_bone_pose] with precalculation.
*/
//go:nosplit
func (self class) SetBoneGlobalPose(bone_idx int64, pose Transform3D.BasisOrigin) { //gd:Skeleton3D.set_bone_global_pose
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, pose)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_bone_global_pose, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Force updates the bone transforms/poses for all bones in the skeleton.
*/
//go:nosplit
func (self class) ForceUpdateAllBoneTransforms() { //gd:Skeleton3D.force_update_all_bone_transforms
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_force_update_all_bone_transforms, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Force updates the bone transform for the bone at [param bone_idx] and all of its children.
*/
//go:nosplit
func (self class) ForceUpdateBoneChildTransform(bone_idx int64) { //gd:Skeleton3D.force_update_bone_child_transform
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_force_update_bone_child_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetMotionScale(motion_scale float64) { //gd:Skeleton3D.set_motion_scale
	var frame = callframe.New()
	callframe.Arg(frame, motion_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_motion_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMotionScale() float64 { //gd:Skeleton3D.get_motion_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_motion_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShowRestOnly(enabled bool) { //gd:Skeleton3D.set_show_rest_only
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_show_rest_only, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShowRestOnly() bool { //gd:Skeleton3D.is_show_rest_only
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_is_show_rest_only, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetModifierCallbackModeProcess(mode gdclass.Skeleton3DModifierCallbackModeProcess) { //gd:Skeleton3D.set_modifier_callback_mode_process
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_modifier_callback_mode_process, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetModifierCallbackModeProcess() gdclass.Skeleton3DModifierCallbackModeProcess { //gd:Skeleton3D.get_modifier_callback_mode_process
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Skeleton3DModifierCallbackModeProcess](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_modifier_callback_mode_process, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the global pose override on all bones in the skeleton.
*/
//go:nosplit
func (self class) ClearBonesGlobalPoseOverride() { //gd:Skeleton3D.clear_bones_global_pose_override
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_clear_bones_global_pose_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the global pose transform, [param pose], for the bone at [param bone_idx].
[param amount] is the interpolation strength that will be used when applying the pose, and [param persistent] determines if the applied pose will remain.
[b]Note:[/b] The pose transform needs to be a global pose! To convert a world transform from a [Node3D] to a global bone pose, multiply the [method Transform3D.affine_inverse] of the node's [member Node3D.global_transform] by the desired world transform.
*/
//go:nosplit
func (self class) SetBoneGlobalPoseOverride(bone_idx int64, pose Transform3D.BasisOrigin, amount float64, persistent bool) { //gd:Skeleton3D.set_bone_global_pose_override
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, pose)
	callframe.Arg(frame, amount)
	callframe.Arg(frame, persistent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_bone_global_pose_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the global pose override transform for [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneGlobalPoseOverride(bone_idx int64) Transform3D.BasisOrigin { //gd:Skeleton3D.get_bone_global_pose_override
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_global_pose_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the overall transform of the specified bone, with respect to the skeleton, but without any global pose overrides. Being relative to the skeleton frame, this is not the actual "global" transform of the bone.
*/
//go:nosplit
func (self class) GetBoneGlobalPoseNoOverride(bone_idx int64) Transform3D.BasisOrigin { //gd:Skeleton3D.get_bone_global_pose_no_override
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_global_pose_no_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAnimatePhysicalBones(enabled bool) { //gd:Skeleton3D.set_animate_physical_bones
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_animate_physical_bones, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAnimatePhysicalBones() bool { //gd:Skeleton3D.get_animate_physical_bones
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_animate_physical_bones, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Tells the [PhysicalBone3D] nodes in the Skeleton to stop simulating.
*/
//go:nosplit
func (self class) PhysicalBonesStopSimulation() { //gd:Skeleton3D.physical_bones_stop_simulation
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_physical_bones_stop_simulation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Tells the [PhysicalBone3D] nodes in the Skeleton to start simulating and reacting to the physics world.
Optionally, a list of bone names can be passed-in, allowing only the passed-in bones to be simulated.
*/
//go:nosplit
func (self class) PhysicalBonesStartSimulation(bones Array.Contains[String.Name]) { //gd:Skeleton3D.physical_bones_start_simulation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(bones)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_physical_bones_start_simulation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
//go:nosplit
func (self class) PhysicalBonesAddCollisionException(exception RID.Any) { //gd:Skeleton3D.physical_bones_add_collision_exception
	var frame = callframe.New()
	callframe.Arg(frame, exception)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_physical_bones_add_collision_exception, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
//go:nosplit
func (self class) PhysicalBonesRemoveCollisionException(exception RID.Any) { //gd:Skeleton3D.physical_bones_remove_collision_exception
	var frame = callframe.New()
	callframe.Arg(frame, exception)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_physical_bones_remove_collision_exception, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnRestUpdated(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("rest_updated"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPoseUpdated(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("pose_updated"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSkeletonUpdated(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("skeleton_updated"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBoneEnabledChanged(cb func(bone_idx int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("bone_enabled_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBoneListChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("bone_list_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnShowRestOnlyChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("show_rest_only_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsSkeleton3D() Advanced       { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSkeleton3D() Instance    { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Advanced(self.AsNode3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Instance(self.AsNode3D()), name)
	}
}
func init() {
	gdclass.Register("Skeleton3D", func(ptr gd.Object) any { return [1]gdclass.Skeleton3D{*(*gdclass.Skeleton3D)(unsafe.Pointer(&ptr))} })
}

type ModifierCallbackModeProcess = gdclass.Skeleton3DModifierCallbackModeProcess //gd:Skeleton3D.ModifierCallbackModeProcess

const (
	/*Set a flag to process modification during physics frames (see [constant Node.NOTIFICATION_INTERNAL_PHYSICS_PROCESS]).*/
	ModifierCallbackModeProcessPhysics ModifierCallbackModeProcess = 0
	/*Set a flag to process modification during process frames (see [constant Node.NOTIFICATION_INTERNAL_PROCESS]).*/
	ModifierCallbackModeProcessIdle ModifierCallbackModeProcess = 1
)
