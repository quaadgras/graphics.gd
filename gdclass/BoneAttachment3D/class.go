package BoneAttachment3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
This node selects a bone in a [Skeleton3D] and attaches to it. This means that the [BoneAttachment3D] node will either dynamically copy or override the 3D transform of the selected bone.

*/
type Go [1]classdb.BoneAttachment3D

/*
A function that is called automatically when the [Skeleton3D] is updated. This function is where the [BoneAttachment3D] node updates its position so it is correctly bound when it is [i]not[/i] set to override the bone pose.
*/
func (self Go) OnSkeletonUpdate() {
	class(self).OnSkeletonUpdate()
}

/*
Sets whether the BoneAttachment3D node will use an external [Skeleton3D] node rather than attempting to use its parent node as the [Skeleton3D]. When set to [code]true[/code], the BoneAttachment3D node will use the external [Skeleton3D] node set in [method set_external_skeleton].
*/
func (self Go) SetUseExternalSkeleton(use_external_skeleton bool) {
	class(self).SetUseExternalSkeleton(use_external_skeleton)
}

/*
Returns whether the BoneAttachment3D node is using an external [Skeleton3D] rather than attempting to use its parent node as the [Skeleton3D].
*/
func (self Go) GetUseExternalSkeleton() bool {
	return bool(class(self).GetUseExternalSkeleton())
}

/*
Sets the [NodePath] to the external skeleton that the BoneAttachment3D node should use. See [method set_use_external_skeleton] to enable the external [Skeleton3D] node.
*/
func (self Go) SetExternalSkeleton(external_skeleton string) {
	class(self).SetExternalSkeleton(gd.NewString(external_skeleton).NodePath())
}

/*
Returns the [NodePath] to the external [Skeleton3D] node, if one has been set.
*/
func (self Go) GetExternalSkeleton() string {
	return string(class(self).GetExternalSkeleton().String())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.BoneAttachment3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("BoneAttachment3D"))
	return Go{classdb.BoneAttachment3D(object)}
}

func (self Go) BoneName() string {
		return string(class(self).GetBoneName().String())
}

func (self Go) SetBoneName(value string) {
	class(self).SetBoneName(gd.NewString(value))
}

func (self Go) BoneIdx() int {
		return int(int(class(self).GetBoneIdx()))
}

func (self Go) SetBoneIdx(value int) {
	class(self).SetBoneIdx(gd.Int(value))
}

func (self Go) OverridePose() bool {
		return bool(class(self).GetOverridePose())
}

func (self Go) SetOverridePose(value bool) {
	class(self).SetOverridePose(value)
}

//go:nosplit
func (self class) SetBoneName(bone_name gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(bone_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneAttachment3D.Bind_set_bone_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBoneName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneAttachment3D.Bind_get_bone_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBoneIdx(bone_idx gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneAttachment3D.Bind_set_bone_idx, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBoneIdx() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneAttachment3D.Bind_get_bone_idx, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
A function that is called automatically when the [Skeleton3D] is updated. This function is where the [BoneAttachment3D] node updates its position so it is correctly bound when it is [i]not[/i] set to override the bone pose.
*/
//go:nosplit
func (self class) OnSkeletonUpdate()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneAttachment3D.Bind_on_skeleton_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetOverridePose(override_pose bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, override_pose)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneAttachment3D.Bind_set_override_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOverridePose() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneAttachment3D.Bind_get_override_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the BoneAttachment3D node will use an external [Skeleton3D] node rather than attempting to use its parent node as the [Skeleton3D]. When set to [code]true[/code], the BoneAttachment3D node will use the external [Skeleton3D] node set in [method set_external_skeleton].
*/
//go:nosplit
func (self class) SetUseExternalSkeleton(use_external_skeleton bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, use_external_skeleton)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneAttachment3D.Bind_set_use_external_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the BoneAttachment3D node is using an external [Skeleton3D] rather than attempting to use its parent node as the [Skeleton3D].
*/
//go:nosplit
func (self class) GetUseExternalSkeleton() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneAttachment3D.Bind_get_use_external_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [NodePath] to the external skeleton that the BoneAttachment3D node should use. See [method set_use_external_skeleton] to enable the external [Skeleton3D] node.
*/
//go:nosplit
func (self class) SetExternalSkeleton(external_skeleton gd.NodePath)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(external_skeleton))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneAttachment3D.Bind_set_external_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [NodePath] to the external [Skeleton3D] node, if one has been set.
*/
//go:nosplit
func (self class) GetExternalSkeleton() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneAttachment3D.Bind_get_external_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsBoneAttachment3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsBoneAttachment3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {classdb.Register("BoneAttachment3D", func(ptr gd.Object) any { return classdb.BoneAttachment3D(ptr) })}
