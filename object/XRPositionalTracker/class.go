package XRPositionalTracker

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/XRTracker"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
An instance of this object represents a device that is tracked, such as a controller or anchor point. HMDs aren't represented here as they are handled internally.
As controllers are turned on and the [XRInterface] detects them, instances of this object are automatically added to this list of active tracking objects accessible through the [XRServer].
The [XRNode3D] and [XRAnchor3D] both consume objects of this type and should be used in your project. The positional trackers are just under-the-hood objects that make this all work. These are mostly exposed so that GDExtension-based interfaces can interact with them.

*/
type Simple [1]classdb.XRPositionalTracker
func (self Simple) GetTrackerProfile() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetTrackerProfile(gc).String())
}
func (self Simple) SetTrackerProfile(profile string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTrackerProfile(gc.String(profile))
}
func (self Simple) GetTrackerHand() classdb.XRPositionalTrackerTrackerHand {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.XRPositionalTrackerTrackerHand(Expert(self).GetTrackerHand())
}
func (self Simple) SetTrackerHand(hand classdb.XRPositionalTrackerTrackerHand) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTrackerHand(hand)
}
func (self Simple) HasPose(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasPose(gc.StringName(name)))
}
func (self Simple) GetPose(name string) [1]classdb.XRPose {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.XRPose(Expert(self).GetPose(gc, gc.StringName(name)))
}
func (self Simple) InvalidatePose(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InvalidatePose(gc.StringName(name))
}
func (self Simple) SetPose(name string, transform gd.Transform3D, linear_velocity gd.Vector3, angular_velocity gd.Vector3, tracking_confidence classdb.XRPoseTrackingConfidence) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPose(gc.StringName(name), transform, linear_velocity, angular_velocity, tracking_confidence)
}
func (self Simple) GetInput(name string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetInput(gc, gc.StringName(name)))
}
func (self Simple) SetInput(name string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInput(gc.StringName(name), value)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.XRPositionalTracker
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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
func (self class) GetPose(ctx gd.Lifetime, name gd.StringName) object.XRPose {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRPositionalTracker.Bind_get_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.XRPose
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

//go:nosplit
func (self class) AsXRPositionalTracker() Expert { return self[0].AsXRPositionalTracker() }


//go:nosplit
func (self Simple) AsXRPositionalTracker() Simple { return self[0].AsXRPositionalTracker() }


//go:nosplit
func (self class) AsXRTracker() XRTracker.Expert { return self[0].AsXRTracker() }


//go:nosplit
func (self Simple) AsXRTracker() XRTracker.Simple { return self[0].AsXRTracker() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("XRPositionalTracker", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
