package XRPositionalTracker

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/XRTracker"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
An instance of this object represents a device that is tracked, such as a controller or anchor point. HMDs aren't represented here as they are handled internally.
As controllers are turned on and the [XRInterface] detects them, instances of this object are automatically added to this list of active tracking objects accessible through the [XRServer].
The [XRNode3D] and [XRAnchor3D] both consume objects of this type and should be used in your project. The positional trackers are just under-the-hood objects that make this all work. These are mostly exposed so that GDExtension-based interfaces can interact with them.

*/
type Go [1]classdb.XRPositionalTracker

/*
Returns [code]true[/code] if the tracker is available and is currently tracking the bound [param name] pose.
*/
func (self Go) HasPose(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasPose(gc.StringName(name)))
}

/*
Returns the current [XRPose] state object for the bound [param name] pose.
*/
func (self Go) GetPose(name string) gdclass.XRPose {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.XRPose(class(self).GetPose(gc, gc.StringName(name)))
}

/*
Marks this pose as invalid, we don't clear the last reported state but it allows users to decide if trackers need to be hidden if we lose tracking or just remain at their last known position.
*/
func (self Go) InvalidatePose(name string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).InvalidatePose(gc.StringName(name))
}

/*
Sets the transform, linear velocity, angular velocity and tracking confidence for the given pose. This method is called by a [XRInterface] implementation and should not be used directly.
*/
func (self Go) SetPose(name string, transform gd.Transform3D, linear_velocity gd.Vector3, angular_velocity gd.Vector3, tracking_confidence classdb.XRPoseTrackingConfidence) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPose(gc.StringName(name), transform, linear_velocity, angular_velocity, tracking_confidence)
}

/*
Returns an input for this tracker. It can return a boolean, float or [Vector2] value depending on whether the input is a button, trigger or thumbstick/thumbpad.
*/
func (self Go) GetInput(name string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).GetInput(gc, gc.StringName(name)))
}

/*
Changes the value for the given input. This method is called by a [XRInterface] implementation and should not be used directly.
*/
func (self Go) SetInput(name string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInput(gc.StringName(name), value)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.XRPositionalTracker
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("XRPositionalTracker"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Profile() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetTrackerProfile(gc).String())
}

func (self Go) SetProfile(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTrackerProfile(gc.String(value))
}

func (self Go) Hand() classdb.XRPositionalTrackerTrackerHand {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.XRPositionalTrackerTrackerHand(class(self).GetTrackerHand())
}

func (self Go) SetHand(value classdb.XRPositionalTrackerTrackerHand) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTrackerHand(value)
}

//go:nosplit
func (self class) GetTrackerProfile(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPositionalTracker.Bind_get_tracker_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTrackerProfile(profile gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(profile))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPositionalTracker.Bind_set_tracker_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTrackerHand() classdb.XRPositionalTrackerTrackerHand {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.XRPositionalTrackerTrackerHand](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPositionalTracker.Bind_get_tracker_hand, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTrackerHand(hand classdb.XRPositionalTrackerTrackerHand)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPositionalTracker.Bind_set_tracker_hand, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the tracker is available and is currently tracking the bound [param name] pose.
*/
//go:nosplit
func (self class) HasPose(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPositionalTracker.Bind_has_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the current [XRPose] state object for the bound [param name] pose.
*/
//go:nosplit
func (self class) GetPose(ctx gd.Lifetime, name gd.StringName) gdclass.XRPose {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPositionalTracker.Bind_get_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.XRPose
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Marks this pose as invalid, we don't clear the last reported state but it allows users to decide if trackers need to be hidden if we lose tracking or just remain at their last known position.
*/
//go:nosplit
func (self class) InvalidatePose(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPositionalTracker.Bind_invalidate_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the transform, linear velocity, angular velocity and tracking confidence for the given pose. This method is called by a [XRInterface] implementation and should not be used directly.
*/
//go:nosplit
func (self class) SetPose(name gd.StringName, transform gd.Transform3D, linear_velocity gd.Vector3, angular_velocity gd.Vector3, tracking_confidence classdb.XRPoseTrackingConfidence)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, transform)
	callframe.Arg(frame, linear_velocity)
	callframe.Arg(frame, angular_velocity)
	callframe.Arg(frame, tracking_confidence)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPositionalTracker.Bind_set_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an input for this tracker. It can return a boolean, float or [Vector2] value depending on whether the input is a button, trigger or thumbstick/thumbpad.
*/
//go:nosplit
func (self class) GetInput(ctx gd.Lifetime, name gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPositionalTracker.Bind_get_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Changes the value for the given input. This method is called by a [XRInterface] implementation and should not be used directly.
*/
//go:nosplit
func (self class) SetInput(name gd.StringName, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPositionalTracker.Bind_set_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnPoseChanged(cb func(pose gdclass.XRPose)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("pose_changed"), gc.Callable(cb), 0)
}


func (self Go) OnPoseLostTracking(cb func(pose gdclass.XRPose)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("pose_lost_tracking"), gc.Callable(cb), 0)
}


func (self Go) OnButtonPressed(cb func(name string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("button_pressed"), gc.Callable(cb), 0)
}


func (self Go) OnButtonReleased(cb func(name string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("button_released"), gc.Callable(cb), 0)
}


func (self Go) OnInputFloatChanged(cb func(name string, value float64)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("input_float_changed"), gc.Callable(cb), 0)
}


func (self Go) OnInputVector2Changed(cb func(name string, vector gd.Vector2)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("input_vector2_changed"), gc.Callable(cb), 0)
}


func (self Go) OnProfileChanged(cb func(role string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("profile_changed"), gc.Callable(cb), 0)
}


func (self class) AsXRPositionalTracker() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRPositionalTracker() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsXRTracker() XRTracker.GD { return *((*XRTracker.GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRTracker() XRTracker.Go { return *((*XRTracker.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsXRTracker(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsXRTracker(), name)
	}
}
func init() {classdb.Register("XRPositionalTracker", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type TrackerHand = classdb.XRPositionalTrackerTrackerHand

const (
/*The hand this tracker is held in is unknown or not applicable.*/
	TrackerHandUnknown TrackerHand = 0
/*This tracker is the left hand controller.*/
	TrackerHandLeft TrackerHand = 1
/*This tracker is the right hand controller.*/
	TrackerHandRight TrackerHand = 2
/*Represents the size of the [enum TrackerHand] enum.*/
	TrackerHandMax TrackerHand = 3
)
