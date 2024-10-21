package XRHandModifier3D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/SkeletonModifier3D"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This node uses hand tracking data from an [XRHandTracker] to pose the skeleton of a hand mesh.
Positioning of hands is performed by creating an [XRNode3D] ancestor of the hand mesh driven by the same [XRHandTracker].
The hand tracking position-data is scaled by [member Skeleton3D.motion_scale] when applied to the skeleton, which can be used to adjust the tracked hand to match the scale of the hand model.

*/
type Simple [1]classdb.XRHandModifier3D
func (self Simple) SetHandTracker(tracker_name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHandTracker(gc.StringName(tracker_name))
}
func (self Simple) GetHandTracker() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetHandTracker(gc).String())
}
func (self Simple) SetBoneUpdate(bone_update classdb.XRHandModifier3DBoneUpdate) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBoneUpdate(bone_update)
}
func (self Simple) GetBoneUpdate() classdb.XRHandModifier3DBoneUpdate {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.XRHandModifier3DBoneUpdate(Expert(self).GetBoneUpdate())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.XRHandModifier3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetHandTracker(tracker_name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(tracker_name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRHandModifier3D.Bind_set_hand_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHandTracker(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRHandModifier3D.Bind_get_hand_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBoneUpdate(bone_update classdb.XRHandModifier3DBoneUpdate)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_update)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRHandModifier3D.Bind_set_bone_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBoneUpdate() classdb.XRHandModifier3DBoneUpdate {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.XRHandModifier3DBoneUpdate](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRHandModifier3D.Bind_get_bone_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsXRHandModifier3D() Expert { return self[0].AsXRHandModifier3D() }


//go:nosplit
func (self Simple) AsXRHandModifier3D() Simple { return self[0].AsXRHandModifier3D() }


//go:nosplit
func (self class) AsSkeletonModifier3D() SkeletonModifier3D.Expert { return self[0].AsSkeletonModifier3D() }


//go:nosplit
func (self Simple) AsSkeletonModifier3D() SkeletonModifier3D.Simple { return self[0].AsSkeletonModifier3D() }


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
func init() {classdb.Register("XRHandModifier3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type BoneUpdate = classdb.XRHandModifier3DBoneUpdate

const (
/*The skeleton's bones are fully updated (both position and rotation) to match the tracked bones.*/
	BoneUpdateFull BoneUpdate = 0
/*The skeleton's bones are only rotated to align with the tracked bones, preserving bone length.*/
	BoneUpdateRotationOnly BoneUpdate = 1
/*Represents the size of the [enum BoneUpdate] enum.*/
	BoneUpdateMax BoneUpdate = 2
)
