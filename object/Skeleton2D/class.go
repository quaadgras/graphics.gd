package Skeleton2D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Node2D"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[Skeleton2D] parents a hierarchy of [Bone2D] nodes. It holds a reference to each [Bone2D]'s rest pose and acts as a single point of access to its bones.
To set up different types of inverse kinematics for the given Skeleton2D, a [SkeletonModificationStack2D] should be created. The inverse kinematics be applied by increasing [member SkeletonModificationStack2D.modification_count] and creating the desired number of modifications.

*/
type Simple [1]classdb.Skeleton2D
func (self Simple) GetBoneCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBoneCount()))
}
func (self Simple) GetBone(idx int) [1]classdb.Bone2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Bone2D(Expert(self).GetBone(gc, gd.Int(idx)))
}
func (self Simple) GetSkeleton() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetSkeleton())
}
func (self Simple) SetModificationStack(modification_stack [1]classdb.SkeletonModificationStack2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetModificationStack(modification_stack)
}
func (self Simple) GetModificationStack() [1]classdb.SkeletonModificationStack2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.SkeletonModificationStack2D(Expert(self).GetModificationStack(gc))
}
func (self Simple) ExecuteModifications(delta float64, execution_mode int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ExecuteModifications(gd.Float(delta), gd.Int(execution_mode))
}
func (self Simple) SetBoneLocalPoseOverride(bone_idx int, override_pose gd.Transform2D, strength float64, persistent bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBoneLocalPoseOverride(gd.Int(bone_idx), override_pose, gd.Float(strength), persistent)
}
func (self Simple) GetBoneLocalPoseOverride(bone_idx int) gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(Expert(self).GetBoneLocalPoseOverride(gd.Int(bone_idx)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Skeleton2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the number of [Bone2D] nodes in the node hierarchy parented by Skeleton2D.
*/
//go:nosplit
func (self class) GetBoneCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton2D.Bind_get_bone_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [Bone2D] from the node hierarchy parented by Skeleton2D. The object to return is identified by the parameter [param idx]. Bones are indexed by descending the node hierarchy from top to bottom, adding the children of each branch before moving to the next sibling.
*/
//go:nosplit
func (self class) GetBone(ctx gd.Lifetime, idx gd.Int) object.Bone2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton2D.Bind_get_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Bone2D
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [RID] of a Skeleton2D instance.
*/
//go:nosplit
func (self class) GetSkeleton() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton2D.Bind_get_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [SkeletonModificationStack2D] attached to this skeleton.
*/
//go:nosplit
func (self class) SetModificationStack(modification_stack object.SkeletonModificationStack2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(modification_stack[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton2D.Bind_set_modification_stack, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [SkeletonModificationStack2D] attached to this skeleton, if one exists.
*/
//go:nosplit
func (self class) GetModificationStack(ctx gd.Lifetime) object.SkeletonModificationStack2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton2D.Bind_get_modification_stack, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.SkeletonModificationStack2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Executes all the modifications on the [SkeletonModificationStack2D], if the Skeleton2D has one assigned.
*/
//go:nosplit
func (self class) ExecuteModifications(delta gd.Float, execution_mode gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, delta)
	callframe.Arg(frame, execution_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton2D.Bind_execute_modifications, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the local pose transform, [param override_pose], for the bone at [param bone_idx].
[param strength] is the interpolation strength that will be used when applying the pose, and [param persistent] determines if the applied pose will remain.
[b]Note:[/b] The pose transform needs to be a local transform relative to the [Bone2D] node at [param bone_idx]!
*/
//go:nosplit
func (self class) SetBoneLocalPoseOverride(bone_idx gd.Int, override_pose gd.Transform2D, strength gd.Float, persistent bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, override_pose)
	callframe.Arg(frame, strength)
	callframe.Arg(frame, persistent)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton2D.Bind_set_bone_local_pose_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the local pose override transform for [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneLocalPoseOverride(bone_idx gd.Int) gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton2D.Bind_get_bone_local_pose_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsSkeleton2D() Expert { return self[0].AsSkeleton2D() }


//go:nosplit
func (self Simple) AsSkeleton2D() Simple { return self[0].AsSkeleton2D() }


//go:nosplit
func (self class) AsNode2D() Node2D.Expert { return self[0].AsNode2D() }


//go:nosplit
func (self Simple) AsNode2D() Node2D.Simple { return self[0].AsNode2D() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("Skeleton2D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
