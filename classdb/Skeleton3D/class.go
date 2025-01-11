// Package Skeleton3D provides methods for working with Skeleton3D object instances.
package Skeleton3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Quaternion"
import "graphics.gd/variant/Float"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

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
func (self Instance) AddBone(name string) int {
	return int(int(class(self).AddBone(gd.NewString(name))))
}

/*
Returns the bone index that matches [param name] as its name. Returns [code]-1[/code] if no bone with this name exists.
*/
func (self Instance) FindBone(name string) int {
	return int(int(class(self).FindBone(gd.NewString(name))))
}

/*
Returns the name of the bone at index [param bone_idx].
*/
func (self Instance) GetBoneName(bone_idx int) string {
	return string(class(self).GetBoneName(gd.Int(bone_idx)).String())
}

/*
Sets the bone name, [param name], for the bone at [param bone_idx].
*/
func (self Instance) SetBoneName(bone_idx int, name string) {
	class(self).SetBoneName(gd.Int(bone_idx), gd.NewString(name))
}

/*
Returns all bone names concatenated with commas ([code],[/code]) as a single [StringName].
It is useful to set it as a hint for the enum property.
*/
func (self Instance) GetConcatenatedBoneNames() string {
	return string(class(self).GetConcatenatedBoneNames().String())
}

/*
Returns the bone index which is the parent of the bone at [param bone_idx]. If -1, then bone has no parent.
[b]Note:[/b] The parent bone returned will always be less than [param bone_idx].
*/
func (self Instance) GetBoneParent(bone_idx int) int {
	return int(int(class(self).GetBoneParent(gd.Int(bone_idx))))
}

/*
Sets the bone index [param parent_idx] as the parent of the bone at [param bone_idx]. If -1, then bone has no parent.
[b]Note:[/b] [param parent_idx] must be less than [param bone_idx].
*/
func (self Instance) SetBoneParent(bone_idx int, parent_idx int) {
	class(self).SetBoneParent(gd.Int(bone_idx), gd.Int(parent_idx))
}

/*
Returns the number of bones in the skeleton.
*/
func (self Instance) GetBoneCount() int {
	return int(int(class(self).GetBoneCount()))
}

/*
Returns the number of times the bone hierarchy has changed within this skeleton, including renames.
The Skeleton version is not serialized: only use within a single instance of Skeleton3D.
Use for invalidating caches in IK solvers and other nodes which process bones.
*/
func (self Instance) GetVersion() int {
	return int(int(class(self).GetVersion()))
}

/*
Unparents the bone at [param bone_idx] and sets its rest position to that of its parent prior to being reset.
*/
func (self Instance) UnparentBoneAndRest(bone_idx int) {
	class(self).UnparentBoneAndRest(gd.Int(bone_idx))
}

/*
Returns an array containing the bone indexes of all the child node of the passed in bone, [param bone_idx].
*/
func (self Instance) GetBoneChildren(bone_idx int) []int32 {
	return []int32(class(self).GetBoneChildren(gd.Int(bone_idx)).AsSlice())
}

/*
Returns an array with all of the bones that are parentless. Another way to look at this is that it returns the indexes of all the bones that are not dependent or modified by other bones in the Skeleton.
*/
func (self Instance) GetParentlessBones() []int32 {
	return []int32(class(self).GetParentlessBones().AsSlice())
}

/*
Returns the rest transform for a bone [param bone_idx].
*/
func (self Instance) GetBoneRest(bone_idx int) Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetBoneRest(gd.Int(bone_idx)))
}

/*
Sets the rest transform for bone [param bone_idx].
*/
func (self Instance) SetBoneRest(bone_idx int, rest Transform3D.BasisOrigin) {
	class(self).SetBoneRest(gd.Int(bone_idx), gd.Transform3D(rest))
}

/*
Returns the global rest transform for [param bone_idx].
*/
func (self Instance) GetBoneGlobalRest(bone_idx int) Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetBoneGlobalRest(gd.Int(bone_idx)))
}
func (self Instance) CreateSkinFromRestTransforms() [1]gdclass.Skin {
	return [1]gdclass.Skin(class(self).CreateSkinFromRestTransforms())
}

/*
Binds the given Skin to the Skeleton.
*/
func (self Instance) RegisterSkin(skin [1]gdclass.Skin) [1]gdclass.SkinReference {
	return [1]gdclass.SkinReference(class(self).RegisterSkin(skin))
}

/*
Returns all bones in the skeleton to their rest poses.
*/
func (self Instance) LocalizeRests() {
	class(self).LocalizeRests()
}

/*
Clear all the bones in this skeleton.
*/
func (self Instance) ClearBones() {
	class(self).ClearBones()
}

/*
Returns the pose transform of the specified bone.
[b]Note:[/b] This is the pose you set to the skeleton in the process, the final pose can get overridden by modifiers in the deferred process, if you want to access the final pose, use [signal SkeletonModifier3D.modification_processed].
*/
func (self Instance) GetBonePose(bone_idx int) Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetBonePose(gd.Int(bone_idx)))
}

/*
Sets the pose transform, [param pose], for the bone at [param bone_idx].
*/
func (self Instance) SetBonePose(bone_idx int, pose Transform3D.BasisOrigin) {
	class(self).SetBonePose(gd.Int(bone_idx), gd.Transform3D(pose))
}

/*
Sets the pose position of the bone at [param bone_idx] to [param position]. [param position] is a [Vector3] describing a position local to the [Skeleton3D] node.
*/
func (self Instance) SetBonePosePosition(bone_idx int, position Vector3.XYZ) {
	class(self).SetBonePosePosition(gd.Int(bone_idx), gd.Vector3(position))
}

/*
Sets the pose rotation of the bone at [param bone_idx] to [param rotation]. [param rotation] is a [Quaternion] describing a rotation in the bone's local coordinate space with respect to the rotation of any parent bones.
*/
func (self Instance) SetBonePoseRotation(bone_idx int, rotation Quaternion.IJKX) {
	class(self).SetBonePoseRotation(gd.Int(bone_idx), gd.Quaternion(rotation))
}

/*
Sets the pose scale of the bone at [param bone_idx] to [param scale].
*/
func (self Instance) SetBonePoseScale(bone_idx int, scale Vector3.XYZ) {
	class(self).SetBonePoseScale(gd.Int(bone_idx), gd.Vector3(scale))
}

/*
Returns the pose position of the bone at [param bone_idx]. The returned [Vector3] is in the local coordinate space of the [Skeleton3D] node.
*/
func (self Instance) GetBonePosePosition(bone_idx int) Vector3.XYZ {
	return Vector3.XYZ(class(self).GetBonePosePosition(gd.Int(bone_idx)))
}

/*
Returns the pose rotation of the bone at [param bone_idx]. The returned [Quaternion] is local to the bone with respect to the rotation of any parent bones.
*/
func (self Instance) GetBonePoseRotation(bone_idx int) Quaternion.IJKX {
	return Quaternion.IJKX(class(self).GetBonePoseRotation(gd.Int(bone_idx)))
}

/*
Returns the pose scale of the bone at [param bone_idx].
*/
func (self Instance) GetBonePoseScale(bone_idx int) Vector3.XYZ {
	return Vector3.XYZ(class(self).GetBonePoseScale(gd.Int(bone_idx)))
}

/*
Sets the bone pose to rest for [param bone_idx].
*/
func (self Instance) ResetBonePose(bone_idx int) {
	class(self).ResetBonePose(gd.Int(bone_idx))
}

/*
Sets all bone poses to rests.
*/
func (self Instance) ResetBonePoses() {
	class(self).ResetBonePoses()
}

/*
Returns whether the bone pose for the bone at [param bone_idx] is enabled.
*/
func (self Instance) IsBoneEnabled(bone_idx int) bool {
	return bool(class(self).IsBoneEnabled(gd.Int(bone_idx)))
}

/*
Disables the pose for the bone at [param bone_idx] if [code]false[/code], enables the bone pose if [code]true[/code].
*/
func (self Instance) SetBoneEnabled(bone_idx int) {
	class(self).SetBoneEnabled(gd.Int(bone_idx), true)
}

/*
Returns the overall transform of the specified bone, with respect to the skeleton. Being relative to the skeleton frame, this is not the actual "global" transform of the bone.
[b]Note:[/b] This is the global pose you set to the skeleton in the process, the final global pose can get overridden by modifiers in the deferred process, if you want to access the final global pose, use [signal SkeletonModifier3D.modification_processed].
*/
func (self Instance) GetBoneGlobalPose(bone_idx int) Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetBoneGlobalPose(gd.Int(bone_idx)))
}

/*
Sets the global pose transform, [param pose], for the bone at [param bone_idx].
[b]Note:[/b] If other bone poses have been changed, this method executes a dirty poses recalculation and will cause performance to deteriorate. If you know that multiple global poses will be applied, consider using [method set_bone_pose] with precalculation.
*/
func (self Instance) SetBoneGlobalPose(bone_idx int, pose Transform3D.BasisOrigin) {
	class(self).SetBoneGlobalPose(gd.Int(bone_idx), gd.Transform3D(pose))
}

/*
Force updates the bone transforms/poses for all bones in the skeleton.
*/
func (self Instance) ForceUpdateAllBoneTransforms() {
	class(self).ForceUpdateAllBoneTransforms()
}

/*
Force updates the bone transform for the bone at [param bone_idx] and all of its children.
*/
func (self Instance) ForceUpdateBoneChildTransform(bone_idx int) {
	class(self).ForceUpdateBoneChildTransform(gd.Int(bone_idx))
}

/*
Removes the global pose override on all bones in the skeleton.
*/
func (self Instance) ClearBonesGlobalPoseOverride() {
	class(self).ClearBonesGlobalPoseOverride()
}

/*
Sets the global pose transform, [param pose], for the bone at [param bone_idx].
[param amount] is the interpolation strength that will be used when applying the pose, and [param persistent] determines if the applied pose will remain.
[b]Note:[/b] The pose transform needs to be a global pose! To convert a world transform from a [Node3D] to a global bone pose, multiply the [method Transform3D.affine_inverse] of the node's [member Node3D.global_transform] by the desired world transform.
*/
func (self Instance) SetBoneGlobalPoseOverride(bone_idx int, pose Transform3D.BasisOrigin, amount Float.X) {
	class(self).SetBoneGlobalPoseOverride(gd.Int(bone_idx), gd.Transform3D(pose), gd.Float(amount), false)
}

/*
Returns the global pose override transform for [param bone_idx].
*/
func (self Instance) GetBoneGlobalPoseOverride(bone_idx int) Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetBoneGlobalPoseOverride(gd.Int(bone_idx)))
}

/*
Returns the overall transform of the specified bone, with respect to the skeleton, but without any global pose overrides. Being relative to the skeleton frame, this is not the actual "global" transform of the bone.
*/
func (self Instance) GetBoneGlobalPoseNoOverride(bone_idx int) Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetBoneGlobalPoseNoOverride(gd.Int(bone_idx)))
}

/*
Tells the [PhysicalBone3D] nodes in the Skeleton to stop simulating.
*/
func (self Instance) PhysicalBonesStopSimulation() {
	class(self).PhysicalBonesStopSimulation()
}

/*
Tells the [PhysicalBone3D] nodes in the Skeleton to start simulating and reacting to the physics world.
Optionally, a list of bone names can be passed-in, allowing only the passed-in bones to be simulated.
*/
func (self Instance) PhysicalBonesStartSimulation() {
	class(self).PhysicalBonesStartSimulation([1]gd.Array{}[0])
}

/*
Adds a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
func (self Instance) PhysicalBonesAddCollisionException(exception Resource.ID) {
	class(self).PhysicalBonesAddCollisionException(exception)
}

/*
Removes a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
func (self Instance) PhysicalBonesRemoveCollisionException(exception Resource.ID) {
	class(self).PhysicalBonesRemoveCollisionException(exception)
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
	class(self).SetMotionScale(gd.Float(value))
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
func (self class) AddBone(name gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_add_bone, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the bone index that matches [param name] as its name. Returns [code]-1[/code] if no bone with this name exists.
*/
//go:nosplit
func (self class) FindBone(name gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_find_bone, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the name of the bone at index [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneName(bone_idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the bone name, [param name], for the bone at [param bone_idx].
*/
//go:nosplit
func (self class) SetBoneName(bone_idx gd.Int, name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_bone_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns all bone names concatenated with commas ([code],[/code]) as a single [StringName].
It is useful to set it as a hint for the enum property.
*/
//go:nosplit
func (self class) GetConcatenatedBoneNames() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_concatenated_bone_names, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the bone index which is the parent of the bone at [param bone_idx]. If -1, then bone has no parent.
[b]Note:[/b] The parent bone returned will always be less than [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneParent(bone_idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
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
func (self class) SetBoneParent(bone_idx gd.Int, parent_idx gd.Int) {
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
func (self class) GetBoneCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
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
func (self class) GetVersion() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_version, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Unparents the bone at [param bone_idx] and sets its rest position to that of its parent prior to being reset.
*/
//go:nosplit
func (self class) UnparentBoneAndRest(bone_idx gd.Int) {
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
func (self class) GetBoneChildren(bone_idx gd.Int) gd.PackedInt32Array {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_children, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an array with all of the bones that are parentless. Another way to look at this is that it returns the indexes of all the bones that are not dependent or modified by other bones in the Skeleton.
*/
//go:nosplit
func (self class) GetParentlessBones() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_parentless_bones, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the rest transform for a bone [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneRest(bone_idx gd.Int) gd.Transform3D {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_rest, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the rest transform for bone [param bone_idx].
*/
//go:nosplit
func (self class) SetBoneRest(bone_idx gd.Int, rest gd.Transform3D) {
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
func (self class) GetBoneGlobalRest(bone_idx gd.Int) gd.Transform3D {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_global_rest, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) CreateSkinFromRestTransforms() [1]gdclass.Skin {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_create_skin_from_rest_transforms, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Skin{gd.PointerWithOwnershipTransferredToGo[gdclass.Skin](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Binds the given Skin to the Skeleton.
*/
//go:nosplit
func (self class) RegisterSkin(skin [1]gdclass.Skin) [1]gdclass.SkinReference {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(skin[0])[0])
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_register_skin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.SkinReference{gd.PointerWithOwnershipTransferredToGo[gdclass.SkinReference](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns all bones in the skeleton to their rest poses.
*/
//go:nosplit
func (self class) LocalizeRests() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_localize_rests, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Clear all the bones in this skeleton.
*/
//go:nosplit
func (self class) ClearBones() {
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
func (self class) GetBonePose(bone_idx gd.Int) gd.Transform3D {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_pose, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the pose transform, [param pose], for the bone at [param bone_idx].
*/
//go:nosplit
func (self class) SetBonePose(bone_idx gd.Int, pose gd.Transform3D) {
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
func (self class) SetBonePosePosition(bone_idx gd.Int, position gd.Vector3) {
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
func (self class) SetBonePoseRotation(bone_idx gd.Int, rotation gd.Quaternion) {
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
func (self class) SetBonePoseScale(bone_idx gd.Int, scale gd.Vector3) {
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
func (self class) GetBonePosePosition(bone_idx gd.Int) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_pose_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the pose rotation of the bone at [param bone_idx]. The returned [Quaternion] is local to the bone with respect to the rotation of any parent bones.
*/
//go:nosplit
func (self class) GetBonePoseRotation(bone_idx gd.Int) gd.Quaternion {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Quaternion](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_pose_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the pose scale of the bone at [param bone_idx].
*/
//go:nosplit
func (self class) GetBonePoseScale(bone_idx gd.Int) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_pose_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the bone pose to rest for [param bone_idx].
*/
//go:nosplit
func (self class) ResetBonePose(bone_idx gd.Int) {
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
func (self class) ResetBonePoses() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_reset_bone_poses, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether the bone pose for the bone at [param bone_idx] is enabled.
*/
//go:nosplit
func (self class) IsBoneEnabled(bone_idx gd.Int) bool {
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
func (self class) SetBoneEnabled(bone_idx gd.Int, enabled bool) {
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
func (self class) GetBoneGlobalPose(bone_idx gd.Int) gd.Transform3D {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
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
func (self class) SetBoneGlobalPose(bone_idx gd.Int, pose gd.Transform3D) {
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
func (self class) ForceUpdateAllBoneTransforms() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_force_update_all_bone_transforms, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Force updates the bone transform for the bone at [param bone_idx] and all of its children.
*/
//go:nosplit
func (self class) ForceUpdateBoneChildTransform(bone_idx gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_force_update_bone_child_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetMotionScale(motion_scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, motion_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_motion_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMotionScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_motion_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShowRestOnly(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_show_rest_only, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShowRestOnly() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_is_show_rest_only, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetModifierCallbackModeProcess(mode gdclass.Skeleton3DModifierCallbackModeProcess) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_modifier_callback_mode_process, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetModifierCallbackModeProcess() gdclass.Skeleton3DModifierCallbackModeProcess {
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
func (self class) ClearBonesGlobalPoseOverride() {
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
func (self class) SetBoneGlobalPoseOverride(bone_idx gd.Int, pose gd.Transform3D, amount gd.Float, persistent bool) {
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
func (self class) GetBoneGlobalPoseOverride(bone_idx gd.Int) gd.Transform3D {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_global_pose_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the overall transform of the specified bone, with respect to the skeleton, but without any global pose overrides. Being relative to the skeleton frame, this is not the actual "global" transform of the bone.
*/
//go:nosplit
func (self class) GetBoneGlobalPoseNoOverride(bone_idx gd.Int) gd.Transform3D {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_get_bone_global_pose_no_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAnimatePhysicalBones(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_set_animate_physical_bones, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAnimatePhysicalBones() bool {
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
func (self class) PhysicalBonesStopSimulation() {
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
func (self class) PhysicalBonesStartSimulation(bones gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(bones))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_physical_bones_start_simulation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
//go:nosplit
func (self class) PhysicalBonesAddCollisionException(exception gd.RID) {
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
func (self class) PhysicalBonesRemoveCollisionException(exception gd.RID) {
	var frame = callframe.New()
	callframe.Arg(frame, exception)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton3D.Bind_physical_bones_remove_collision_exception, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
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

type ModifierCallbackModeProcess = gdclass.Skeleton3DModifierCallbackModeProcess

const (
	/*Set a flag to process modification during physics frames (see [constant Node.NOTIFICATION_INTERNAL_PHYSICS_PROCESS]).*/
	ModifierCallbackModeProcessPhysics ModifierCallbackModeProcess = 0
	/*Set a flag to process modification during process frames (see [constant Node.NOTIFICATION_INTERNAL_PROCESS]).*/
	ModifierCallbackModeProcessIdle ModifierCallbackModeProcess = 1
)
