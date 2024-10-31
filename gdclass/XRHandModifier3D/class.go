package XRHandModifier3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/SkeletonModifier3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
This node uses hand tracking data from an [XRHandTracker] to pose the skeleton of a hand mesh.
Positioning of hands is performed by creating an [XRNode3D] ancestor of the hand mesh driven by the same [XRHandTracker].
The hand tracking position-data is scaled by [member Skeleton3D.motion_scale] when applied to the skeleton, which can be used to adjust the tracked hand to match the scale of the hand model.
*/
type Instance [1]classdb.XRHandModifier3D

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.XRHandModifier3D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRHandModifier3D"))
	return Instance{classdb.XRHandModifier3D(object)}
}

func (self Instance) HandTracker() string {
	return string(class(self).GetHandTracker().String())
}

func (self Instance) SetHandTracker(value string) {
	class(self).SetHandTracker(gd.NewStringName(value))
}

func (self Instance) BoneUpdate() classdb.XRHandModifier3DBoneUpdate {
	return classdb.XRHandModifier3DBoneUpdate(class(self).GetBoneUpdate())
}

func (self Instance) SetBoneUpdate(value classdb.XRHandModifier3DBoneUpdate) {
	class(self).SetBoneUpdate(value)
}

//go:nosplit
func (self class) SetHandTracker(tracker_name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tracker_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandModifier3D.Bind_set_hand_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHandTracker() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandModifier3D.Bind_get_hand_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBoneUpdate(bone_update classdb.XRHandModifier3DBoneUpdate) {
	var frame = callframe.New()
	callframe.Arg(frame, bone_update)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandModifier3D.Bind_set_bone_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBoneUpdate() classdb.XRHandModifier3DBoneUpdate {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.XRHandModifier3DBoneUpdate](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandModifier3D.Bind_get_bone_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsXRHandModifier3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsXRHandModifier3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsSkeletonModifier3D() SkeletonModifier3D.Advanced {
	return *((*SkeletonModifier3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSkeletonModifier3D() SkeletonModifier3D.Instance {
	return *((*SkeletonModifier3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsSkeletonModifier3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsSkeletonModifier3D(), name)
	}
}
func init() {
	classdb.Register("XRHandModifier3D", func(ptr gd.Object) any { return classdb.XRHandModifier3D(ptr) })
}

type BoneUpdate = classdb.XRHandModifier3DBoneUpdate

const (
	/*The skeleton's bones are fully updated (both position and rotation) to match the tracked bones.*/
	BoneUpdateFull BoneUpdate = 0
	/*The skeleton's bones are only rotated to align with the tracked bones, preserving bone length.*/
	BoneUpdateRotationOnly BoneUpdate = 1
	/*Represents the size of the [enum BoneUpdate] enum.*/
	BoneUpdateMax BoneUpdate = 2
)
