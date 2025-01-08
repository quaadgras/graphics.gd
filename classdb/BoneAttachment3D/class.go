// Package BoneAttachment3D provides methods for working with BoneAttachment3D object instances.
package BoneAttachment3D

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
import "graphics.gd/variant/NodePath"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This node selects a bone in a [Skeleton3D] and attaches to it. This means that the [BoneAttachment3D] node will either dynamically copy or override the 3D transform of the selected bone.
*/
type Instance [1]gdclass.BoneAttachment3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsBoneAttachment3D() Instance
}

/*
A function that is called automatically when the [Skeleton3D] is updated. This function is where the [BoneAttachment3D] node updates its position so it is correctly bound when it is [i]not[/i] set to override the bone pose.
*/
func (self Instance) OnSkeletonUpdate() {
	class(self).OnSkeletonUpdate()
}

/*
Sets whether the BoneAttachment3D node will use an external [Skeleton3D] node rather than attempting to use its parent node as the [Skeleton3D]. When set to [code]true[/code], the BoneAttachment3D node will use the external [Skeleton3D] node set in [method set_external_skeleton].
*/
func (self Instance) SetUseExternalSkeleton(use_external_skeleton bool) {
	class(self).SetUseExternalSkeleton(use_external_skeleton)
}

/*
Returns whether the BoneAttachment3D node is using an external [Skeleton3D] rather than attempting to use its parent node as the [Skeleton3D].
*/
func (self Instance) GetUseExternalSkeleton() bool {
	return bool(class(self).GetUseExternalSkeleton())
}

/*
Sets the [NodePath] to the external skeleton that the BoneAttachment3D node should use. See [method set_use_external_skeleton] to enable the external [Skeleton3D] node.
*/
func (self Instance) SetExternalSkeleton(external_skeleton NodePath.String) {
	class(self).SetExternalSkeleton(gd.NewString(string(external_skeleton)).NodePath())
}

/*
Returns the [NodePath] to the external [Skeleton3D] node, if one has been set.
*/
func (self Instance) GetExternalSkeleton() NodePath.String {
	return NodePath.String(class(self).GetExternalSkeleton().String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.BoneAttachment3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("BoneAttachment3D"))
	casted := Instance{*(*gdclass.BoneAttachment3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) BoneName() string {
	return string(class(self).GetBoneName().String())
}

func (self Instance) SetBoneName(value string) {
	class(self).SetBoneName(gd.NewString(value))
}

func (self Instance) BoneIdx() int {
	return int(int(class(self).GetBoneIdx()))
}

func (self Instance) SetBoneIdx(value int) {
	class(self).SetBoneIdx(gd.Int(value))
}

func (self Instance) OverridePose() bool {
	return bool(class(self).GetOverridePose())
}

func (self Instance) SetOverridePose(value bool) {
	class(self).SetOverridePose(value)
}

//go:nosplit
func (self class) SetBoneName(bone_name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(bone_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneAttachment3D.Bind_set_bone_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBoneName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneAttachment3D.Bind_get_bone_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBoneIdx(bone_idx gd.Int) {
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
func (self class) OnSkeletonUpdate() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneAttachment3D.Bind_on_skeleton_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetOverridePose(override_pose bool) {
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
func (self class) SetUseExternalSkeleton(use_external_skeleton bool) {
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
func (self class) SetExternalSkeleton(external_skeleton gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(external_skeleton))
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
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsBoneAttachment3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsBoneAttachment3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced       { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance    { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced           { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance        { return *((*Node.Instance)(unsafe.Pointer(&self))) }

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
	gdclass.Register("BoneAttachment3D", func(ptr gd.Object) any {
		return [1]gdclass.BoneAttachment3D{*(*gdclass.BoneAttachment3D)(unsafe.Pointer(&ptr))}
	})
}
