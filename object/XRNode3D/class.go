package XRNode3D

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
This node can be bound to a specific pose of a [XRPositionalTracker] and will automatically have its [member Node3D.transform] updated by the [XRServer]. Nodes of this type must be added as children of the [XROrigin3D] node.

*/
type Simple [1]classdb.XRNode3D
func (self Simple) SetTracker(tracker_name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTracker(gc.StringName(tracker_name))
}
func (self Simple) GetTracker() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetTracker(gc).String())
}
func (self Simple) SetPoseName(pose string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPoseName(gc.StringName(pose))
}
func (self Simple) GetPoseName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetPoseName(gc).String())
}
func (self Simple) SetShowWhenTracked(show bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShowWhenTracked(show)
}
func (self Simple) GetShowWhenTracked() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetShowWhenTracked())
}
func (self Simple) GetIsActive() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetIsActive())
}
func (self Simple) GetHasTrackingData() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetHasTrackingData())
}
func (self Simple) GetPose() [1]classdb.XRPose {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.XRPose(Expert(self).GetPose(gc))
}
func (self Simple) TriggerHapticPulse(action_name string, frequency float64, amplitude float64, duration_sec float64, delay_sec float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).TriggerHapticPulse(gc.String(action_name), gd.Float(frequency), gd.Float(amplitude), gd.Float(duration_sec), gd.Float(delay_sec))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.XRNode3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetTracker(tracker_name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(tracker_name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRNode3D.Bind_set_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTracker(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRNode3D.Bind_get_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPoseName(pose gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(pose))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRNode3D.Bind_set_pose_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPoseName(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRNode3D.Bind_get_pose_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowWhenTracked(show bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, show)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRNode3D.Bind_set_show_when_tracked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShowWhenTracked() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRNode3D.Bind_get_show_when_tracked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the [member tracker] has been registered and the [member pose] is being tracked.
*/
//go:nosplit
func (self class) GetIsActive() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRNode3D.Bind_get_is_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the [member tracker] has current tracking data for the [member pose] being tracked.
*/
//go:nosplit
func (self class) GetHasTrackingData() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRNode3D.Bind_get_has_tracking_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [XRPose] containing the current state of the pose being tracked. This gives access to additional properties of this pose.
*/
//go:nosplit
func (self class) GetPose(ctx gd.Lifetime) object.XRPose {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRNode3D.Bind_get_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.XRPose
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Triggers a haptic pulse on a device associated with this interface.
[param action_name] is the name of the action for this pulse.
[param frequency] is the frequency of the pulse, set to [code]0.0[/code] to have the system use a default frequency.
[param amplitude] is the amplitude of the pulse between [code]0.0[/code] and [code]1.0[/code].
[param duration_sec] is the duration of the pulse in seconds.
[param delay_sec] is a delay in seconds before the pulse is given.
*/
//go:nosplit
func (self class) TriggerHapticPulse(action_name gd.String, frequency gd.Float, amplitude gd.Float, duration_sec gd.Float, delay_sec gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action_name))
	callframe.Arg(frame, frequency)
	callframe.Arg(frame, amplitude)
	callframe.Arg(frame, duration_sec)
	callframe.Arg(frame, delay_sec)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRNode3D.Bind_trigger_haptic_pulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsXRNode3D() Expert { return self[0].AsXRNode3D() }


//go:nosplit
func (self Simple) AsXRNode3D() Simple { return self[0].AsXRNode3D() }


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
func init() {classdb.Register("XRNode3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
