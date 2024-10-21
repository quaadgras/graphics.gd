package BoneAttachment3D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
This node selects a bone in a [Skeleton3D] and attaches to it. This means that the [BoneAttachment3D] node will either dynamically copy or override the 3D transform of the selected bone.

*/
type Simple [1]classdb.BoneAttachment3D
func (self Simple) SetBoneName(bone_name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBoneName(gc.String(bone_name))
}
func (self Simple) GetBoneName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetBoneName(gc).String())
}
func (self Simple) SetBoneIdx(bone_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBoneIdx(gd.Int(bone_idx))
}
func (self Simple) GetBoneIdx() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBoneIdx()))
}
func (self Simple) OnSkeletonUpdate() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).OnSkeletonUpdate()
}
func (self Simple) SetOverridePose(override_pose bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOverridePose(override_pose)
}
func (self Simple) GetOverridePose() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetOverridePose())
}
func (self Simple) SetUseExternalSkeleton(use_external_skeleton bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseExternalSkeleton(use_external_skeleton)
}
func (self Simple) GetUseExternalSkeleton() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetUseExternalSkeleton())
}
func (self Simple) SetExternalSkeleton(external_skeleton gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetExternalSkeleton(external_skeleton)
}
func (self Simple) GetExternalSkeleton() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetExternalSkeleton(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.BoneAttachment3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetBoneName(bone_name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bone_name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoneAttachment3D.Bind_set_bone_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBoneName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoneAttachment3D.Bind_get_bone_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBoneIdx(bone_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoneAttachment3D.Bind_set_bone_idx, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBoneIdx() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoneAttachment3D.Bind_get_bone_idx, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
A function that is called automatically when the [Skeleton3D] is updated. This function is where the [BoneAttachment3D] node updates its position so it is correctly bound when it is [i]not[/i] set to override the bone pose.
*/
//go:nosplit
func (self class) OnSkeletonUpdate()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoneAttachment3D.Bind_on_skeleton_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetOverridePose(override_pose bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, override_pose)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoneAttachment3D.Bind_set_override_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOverridePose() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoneAttachment3D.Bind_get_override_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the BoneAttachment3D node will use an external [Skeleton3D] node rather than attempting to use its parent node as the [Skeleton3D]. When set to [code]true[/code], the BoneAttachment3D node will use the external [Skeleton3D] node set in [method set_external_skeleton].
*/
//go:nosplit
func (self class) SetUseExternalSkeleton(use_external_skeleton bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use_external_skeleton)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoneAttachment3D.Bind_set_use_external_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the BoneAttachment3D node is using an external [Skeleton3D] rather than attempting to use its parent node as the [Skeleton3D].
*/
//go:nosplit
func (self class) GetUseExternalSkeleton() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoneAttachment3D.Bind_get_use_external_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [NodePath] to the external skeleton that the BoneAttachment3D node should use. See [method set_use_external_skeleton] to enable the external [Skeleton3D] node.
*/
//go:nosplit
func (self class) SetExternalSkeleton(external_skeleton gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(external_skeleton))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoneAttachment3D.Bind_set_external_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [NodePath] to the external [Skeleton3D] node, if one has been set.
*/
//go:nosplit
func (self class) GetExternalSkeleton(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BoneAttachment3D.Bind_get_external_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsBoneAttachment3D() Expert { return self[0].AsBoneAttachment3D() }


//go:nosplit
func (self Simple) AsBoneAttachment3D() Simple { return self[0].AsBoneAttachment3D() }


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
func init() {classdb.Register("BoneAttachment3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
