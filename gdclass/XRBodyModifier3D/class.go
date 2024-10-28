package XRBodyModifier3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/SkeletonModifier3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
This node uses body tracking data from an [XRBodyTracker] to pose the skeleton of a body mesh.
Positioning of the body is performed by creating an [XRNode3D] ancestor of the body mesh driven by the same [XRBodyTracker].
The body tracking position-data is scaled by [member Skeleton3D.motion_scale] when applied to the skeleton, which can be used to adjust the tracked body to match the scale of the body model.

*/
type Go [1]classdb.XRBodyModifier3D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.XRBodyModifier3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRBodyModifier3D"))
	return Go{classdb.XRBodyModifier3D(object)}
}

func (self Go) BodyTracker() string {
		return string(class(self).GetBodyTracker().String())
}

func (self Go) SetBodyTracker(value string) {
	class(self).SetBodyTracker(gd.NewStringName(value))
}

func (self Go) BodyUpdate() classdb.XRBodyModifier3DBodyUpdate {
		return classdb.XRBodyModifier3DBodyUpdate(class(self).GetBodyUpdate())
}

func (self Go) SetBodyUpdate(value classdb.XRBodyModifier3DBodyUpdate) {
	class(self).SetBodyUpdate(value)
}

func (self Go) BoneUpdate() classdb.XRBodyModifier3DBoneUpdate {
		return classdb.XRBodyModifier3DBoneUpdate(class(self).GetBoneUpdate())
}

func (self Go) SetBoneUpdate(value classdb.XRBodyModifier3DBoneUpdate) {
	class(self).SetBoneUpdate(value)
}

//go:nosplit
func (self class) SetBodyTracker(tracker_name gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(tracker_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRBodyModifier3D.Bind_set_body_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBodyTracker() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRBodyModifier3D.Bind_get_body_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBodyUpdate(body_update classdb.XRBodyModifier3DBodyUpdate)  {
	var frame = callframe.New()
	callframe.Arg(frame, body_update)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRBodyModifier3D.Bind_set_body_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBodyUpdate() classdb.XRBodyModifier3DBodyUpdate {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.XRBodyModifier3DBodyUpdate](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRBodyModifier3D.Bind_get_body_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBoneUpdate(bone_update classdb.XRBodyModifier3DBoneUpdate)  {
	var frame = callframe.New()
	callframe.Arg(frame, bone_update)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRBodyModifier3D.Bind_set_bone_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBoneUpdate() classdb.XRBodyModifier3DBoneUpdate {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.XRBodyModifier3DBoneUpdate](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRBodyModifier3D.Bind_get_bone_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsXRBodyModifier3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRBodyModifier3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsSkeletonModifier3D() SkeletonModifier3D.GD { return *((*SkeletonModifier3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsSkeletonModifier3D() SkeletonModifier3D.Go { return *((*SkeletonModifier3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsSkeletonModifier3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsSkeletonModifier3D(), name)
	}
}
func init() {classdb.Register("XRBodyModifier3D", func(ptr gd.Object) any { return classdb.XRBodyModifier3D(ptr) })}
type BodyUpdate = classdb.XRBodyModifier3DBodyUpdate

const (
/*The skeleton's upper body joints are updated.*/
	BodyUpdateUpperBody BodyUpdate = 1
/*The skeleton's lower body joints are updated.*/
	BodyUpdateLowerBody BodyUpdate = 2
/*The skeleton's hand joints are updated.*/
	BodyUpdateHands BodyUpdate = 4
)
type BoneUpdate = classdb.XRBodyModifier3DBoneUpdate

const (
/*The skeleton's bones are fully updated (both position and rotation) to match the tracked bones.*/
	BoneUpdateFull BoneUpdate = 0
/*The skeleton's bones are only rotated to align with the tracked bones, preserving bone length.*/
	BoneUpdateRotationOnly BoneUpdate = 1
/*Represents the size of the [enum BoneUpdate] enum.*/
	BoneUpdateMax BoneUpdate = 2
)
