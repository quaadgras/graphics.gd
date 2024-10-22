package Skeleton2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node2D"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[Skeleton2D] parents a hierarchy of [Bone2D] nodes. It holds a reference to each [Bone2D]'s rest pose and acts as a single point of access to its bones.
To set up different types of inverse kinematics for the given Skeleton2D, a [SkeletonModificationStack2D] should be created. The inverse kinematics be applied by increasing [member SkeletonModificationStack2D.modification_count] and creating the desired number of modifications.

*/
type Go [1]classdb.Skeleton2D

/*
Returns the number of [Bone2D] nodes in the node hierarchy parented by Skeleton2D.
*/
func (self Go) GetBoneCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetBoneCount()))
}

/*
Returns a [Bone2D] from the node hierarchy parented by Skeleton2D. The object to return is identified by the parameter [param idx]. Bones are indexed by descending the node hierarchy from top to bottom, adding the children of each branch before moving to the next sibling.
*/
func (self Go) GetBone(idx int) gdclass.Bone2D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Bone2D(class(self).GetBone(gc, gd.Int(idx)))
}

/*
Returns the [RID] of a Skeleton2D instance.
*/
func (self Go) GetSkeleton() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetSkeleton())
}

/*
Sets the [SkeletonModificationStack2D] attached to this skeleton.
*/
func (self Go) SetModificationStack(modification_stack gdclass.SkeletonModificationStack2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetModificationStack(modification_stack)
}

/*
Returns the [SkeletonModificationStack2D] attached to this skeleton, if one exists.
*/
func (self Go) GetModificationStack() gdclass.SkeletonModificationStack2D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.SkeletonModificationStack2D(class(self).GetModificationStack(gc))
}

/*
Executes all the modifications on the [SkeletonModificationStack2D], if the Skeleton2D has one assigned.
*/
func (self Go) ExecuteModifications(delta float64, execution_mode int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ExecuteModifications(gd.Float(delta), gd.Int(execution_mode))
}

/*
Sets the local pose transform, [param override_pose], for the bone at [param bone_idx].
[param strength] is the interpolation strength that will be used when applying the pose, and [param persistent] determines if the applied pose will remain.
[b]Note:[/b] The pose transform needs to be a local transform relative to the [Bone2D] node at [param bone_idx]!
*/
func (self Go) SetBoneLocalPoseOverride(bone_idx int, override_pose gd.Transform2D, strength float64, persistent bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBoneLocalPoseOverride(gd.Int(bone_idx), override_pose, gd.Float(strength), persistent)
}

/*
Returns the local pose override transform for [param bone_idx].
*/
func (self Go) GetBoneLocalPoseOverride(bone_idx int) gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(class(self).GetBoneLocalPoseOverride(gd.Int(bone_idx)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Skeleton2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Skeleton2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

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
func (self class) GetBone(ctx gd.Lifetime, idx gd.Int) gdclass.Bone2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton2D.Bind_get_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Bone2D
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
func (self class) SetModificationStack(modification_stack gdclass.SkeletonModificationStack2D)  {
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
func (self class) GetModificationStack(ctx gd.Lifetime) gdclass.SkeletonModificationStack2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Skeleton2D.Bind_get_modification_stack, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.SkeletonModificationStack2D
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
func (self Go) OnBoneSetupChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("bone_setup_changed"), gc.Callable(cb), 0)
}


func (self class) AsSkeleton2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSkeleton2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.GD { return *((*Node2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode2D() Node2D.Go { return *((*Node2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {classdb.Register("Skeleton2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
