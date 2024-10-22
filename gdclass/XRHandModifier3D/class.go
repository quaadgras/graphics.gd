package XRHandModifier3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
var _ mmm.Lifetime

/*
This node uses hand tracking data from an [XRHandTracker] to pose the skeleton of a hand mesh.
Positioning of hands is performed by creating an [XRNode3D] ancestor of the hand mesh driven by the same [XRHandTracker].
The hand tracking position-data is scaled by [member Skeleton3D.motion_scale] when applied to the skeleton, which can be used to adjust the tracked hand to match the scale of the hand model.

*/
type Go [1]classdb.XRHandModifier3D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.XRHandModifier3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("XRHandModifier3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) HandTracker() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetHandTracker(gc).String())
}

func (self Go) SetHandTracker(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHandTracker(gc.StringName(value))
}

func (self Go) BoneUpdate() classdb.XRHandModifier3DBoneUpdate {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.XRHandModifier3DBoneUpdate(class(self).GetBoneUpdate())
}

func (self Go) SetBoneUpdate(value classdb.XRHandModifier3DBoneUpdate) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBoneUpdate(value)
}

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
func (self class) AsXRHandModifier3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRHandModifier3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("XRHandModifier3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type BoneUpdate = classdb.XRHandModifier3DBoneUpdate

const (
/*The skeleton's bones are fully updated (both position and rotation) to match the tracked bones.*/
	BoneUpdateFull BoneUpdate = 0
/*The skeleton's bones are only rotated to align with the tracked bones, preserving bone length.*/
	BoneUpdateRotationOnly BoneUpdate = 1
/*Represents the size of the [enum BoneUpdate] enum.*/
	BoneUpdateMax BoneUpdate = 2
)
