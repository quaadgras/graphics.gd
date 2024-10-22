package XRPose

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
XR runtimes often identify multiple locations on devices such as controllers that are spatially tracked.
Orientation, location, linear velocity and angular velocity are all provided for each pose by the XR runtime. This object contains this state of a pose.

*/
type Go [1]classdb.XRPose

/*
Returns the [member transform] with world scale and our reference frame applied. This is the transform used to position [XRNode3D] objects.
*/
func (self Go) GetAdjustedTransform() gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(class(self).GetAdjustedTransform())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.XRPose
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("XRPose"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) HasTrackingData() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetHasTrackingData())
}

func (self Go) SetHasTrackingData(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHasTrackingData(value)
}

func (self Go) Name() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetName(gc).String())
}

func (self Go) SetName(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetName(gc.StringName(value))
}

func (self Go) Transform() gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Transform3D(class(self).GetTransform())
}

func (self Go) SetTransform(value gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTransform(value)
}

func (self Go) LinearVelocity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetLinearVelocity())
}

func (self Go) SetLinearVelocity(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLinearVelocity(value)
}

func (self Go) AngularVelocity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetAngularVelocity())
}

func (self Go) SetAngularVelocity(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAngularVelocity(value)
}

func (self Go) TrackingConfidence() classdb.XRPoseTrackingConfidence {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.XRPoseTrackingConfidence(class(self).GetTrackingConfidence())
}

func (self Go) SetTrackingConfidence(value classdb.XRPoseTrackingConfidence) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTrackingConfidence(value)
}

//go:nosplit
func (self class) SetHasTrackingData(has_tracking_data bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, has_tracking_data)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPose.Bind_set_has_tracking_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHasTrackingData() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPose.Bind_get_has_tracking_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetName(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPose.Bind_set_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetName(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPose.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTransform(transform gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPose.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTransform() gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPose.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [member transform] with world scale and our reference frame applied. This is the transform used to position [XRNode3D] objects.
*/
//go:nosplit
func (self class) GetAdjustedTransform() gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPose.Bind_get_adjusted_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLinearVelocity(velocity gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPose.Bind_set_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLinearVelocity() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPose.Bind_get_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAngularVelocity(velocity gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPose.Bind_set_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAngularVelocity() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPose.Bind_get_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTrackingConfidence(tracking_confidence classdb.XRPoseTrackingConfidence)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tracking_confidence)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPose.Bind_set_tracking_confidence, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTrackingConfidence() classdb.XRPoseTrackingConfidence {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.XRPoseTrackingConfidence](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPose.Bind_get_tracking_confidence, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsXRPose() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRPose() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("XRPose", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type TrackingConfidence = classdb.XRPoseTrackingConfidence

const (
/*No tracking information is available for this pose.*/
	XrTrackingConfidenceNone TrackingConfidence = 0
/*Tracking information may be inaccurate or estimated. For example, with inside out tracking this would indicate a controller may be (partially) obscured.*/
	XrTrackingConfidenceLow TrackingConfidence = 1
/*Tracking information is considered accurate and up to date.*/
	XrTrackingConfidenceHigh TrackingConfidence = 2
)
