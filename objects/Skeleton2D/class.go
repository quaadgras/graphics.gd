package Skeleton2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Node2D"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"
import "grow.graphics/gd/objects/Resource"
import "grow.graphics/gd/variant/Float"
import "grow.graphics/gd/variant/Transform2D"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
[Skeleton2D] parents a hierarchy of [Bone2D] nodes. It holds a reference to each [Bone2D]'s rest pose and acts as a single point of access to its bones.
To set up different types of inverse kinematics for the given Skeleton2D, a [SkeletonModificationStack2D] should be created. The inverse kinematics be applied by increasing [member SkeletonModificationStack2D.modification_count] and creating the desired number of modifications.
*/
type Instance [1]classdb.Skeleton2D

/*
Returns the number of [Bone2D] nodes in the node hierarchy parented by Skeleton2D.
*/
func (self Instance) GetBoneCount() int {
	return int(int(class(self).GetBoneCount()))
}

/*
Returns a [Bone2D] from the node hierarchy parented by Skeleton2D. The object to return is identified by the parameter [param idx]. Bones are indexed by descending the node hierarchy from top to bottom, adding the children of each branch before moving to the next sibling.
*/
func (self Instance) GetBone(idx int) objects.Bone2D {
	return objects.Bone2D(class(self).GetBone(gd.Int(idx)))
}

/*
Returns the [RID] of a Skeleton2D instance.
*/
func (self Instance) GetSkeleton() Resource.ID {
	return Resource.ID(class(self).GetSkeleton())
}

/*
Sets the [SkeletonModificationStack2D] attached to this skeleton.
*/
func (self Instance) SetModificationStack(modification_stack objects.SkeletonModificationStack2D) {
	class(self).SetModificationStack(modification_stack)
}

/*
Returns the [SkeletonModificationStack2D] attached to this skeleton, if one exists.
*/
func (self Instance) GetModificationStack() objects.SkeletonModificationStack2D {
	return objects.SkeletonModificationStack2D(class(self).GetModificationStack())
}

/*
Executes all the modifications on the [SkeletonModificationStack2D], if the Skeleton2D has one assigned.
*/
func (self Instance) ExecuteModifications(delta Float.X, execution_mode int) {
	class(self).ExecuteModifications(gd.Float(delta), gd.Int(execution_mode))
}

/*
Sets the local pose transform, [param override_pose], for the bone at [param bone_idx].
[param strength] is the interpolation strength that will be used when applying the pose, and [param persistent] determines if the applied pose will remain.
[b]Note:[/b] The pose transform needs to be a local transform relative to the [Bone2D] node at [param bone_idx]!
*/
func (self Instance) SetBoneLocalPoseOverride(bone_idx int, override_pose Transform2D.OriginXY, strength Float.X, persistent bool) {
	class(self).SetBoneLocalPoseOverride(gd.Int(bone_idx), gd.Transform2D(override_pose), gd.Float(strength), persistent)
}

/*
Returns the local pose override transform for [param bone_idx].
*/
func (self Instance) GetBoneLocalPoseOverride(bone_idx int) Transform2D.OriginXY {
	return Transform2D.OriginXY(class(self).GetBoneLocalPoseOverride(gd.Int(bone_idx)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Skeleton2D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Skeleton2D"))
	return Instance{classdb.Skeleton2D(object)}
}

/*
Returns the number of [Bone2D] nodes in the node hierarchy parented by Skeleton2D.
*/
//go:nosplit
func (self class) GetBoneCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton2D.Bind_get_bone_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [Bone2D] from the node hierarchy parented by Skeleton2D. The object to return is identified by the parameter [param idx]. Bones are indexed by descending the node hierarchy from top to bottom, adding the children of each branch before moving to the next sibling.
*/
//go:nosplit
func (self class) GetBone(idx gd.Int) objects.Bone2D {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton2D.Bind_get_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Bone2D{classdb.Bone2D(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the [RID] of a Skeleton2D instance.
*/
//go:nosplit
func (self class) GetSkeleton() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton2D.Bind_get_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [SkeletonModificationStack2D] attached to this skeleton.
*/
//go:nosplit
func (self class) SetModificationStack(modification_stack objects.SkeletonModificationStack2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(modification_stack[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton2D.Bind_set_modification_stack, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [SkeletonModificationStack2D] attached to this skeleton, if one exists.
*/
//go:nosplit
func (self class) GetModificationStack() objects.SkeletonModificationStack2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton2D.Bind_get_modification_stack, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.SkeletonModificationStack2D{classdb.SkeletonModificationStack2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Executes all the modifications on the [SkeletonModificationStack2D], if the Skeleton2D has one assigned.
*/
//go:nosplit
func (self class) ExecuteModifications(delta gd.Float, execution_mode gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, delta)
	callframe.Arg(frame, execution_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton2D.Bind_execute_modifications, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the local pose transform, [param override_pose], for the bone at [param bone_idx].
[param strength] is the interpolation strength that will be used when applying the pose, and [param persistent] determines if the applied pose will remain.
[b]Note:[/b] The pose transform needs to be a local transform relative to the [Bone2D] node at [param bone_idx]!
*/
//go:nosplit
func (self class) SetBoneLocalPoseOverride(bone_idx gd.Int, override_pose gd.Transform2D, strength gd.Float, persistent bool) {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, override_pose)
	callframe.Arg(frame, strength)
	callframe.Arg(frame, persistent)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton2D.Bind_set_bone_local_pose_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the local pose override transform for [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneLocalPoseOverride(bone_idx gd.Int) gd.Transform2D {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton2D.Bind_get_bone_local_pose_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnBoneSetupChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("bone_setup_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsSkeleton2D() Advanced       { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSkeleton2D() Instance    { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced    { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {
	classdb.Register("Skeleton2D", func(ptr gd.Object) any { return classdb.Skeleton2D(ptr) })
}
