package XRNode3D

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
This node can be bound to a specific pose of a [XRPositionalTracker] and will automatically have its [member Node3D.transform] updated by the [XRServer]. Nodes of this type must be added as children of the [XROrigin3D] node.

*/
type Go [1]classdb.XRNode3D

/*
Returns [code]true[/code] if the [member tracker] has been registered and the [member pose] is being tracked.
*/
func (self Go) GetIsActive() bool {
	return bool(class(self).GetIsActive())
}

/*
Returns [code]true[/code] if the [member tracker] has current tracking data for the [member pose] being tracked.
*/
func (self Go) GetHasTrackingData() bool {
	return bool(class(self).GetHasTrackingData())
}

/*
Returns the [XRPose] containing the current state of the pose being tracked. This gives access to additional properties of this pose.
*/
func (self Go) GetPose() gdclass.XRPose {
	return gdclass.XRPose(class(self).GetPose())
}

/*
Triggers a haptic pulse on a device associated with this interface.
[param action_name] is the name of the action for this pulse.
[param frequency] is the frequency of the pulse, set to [code]0.0[/code] to have the system use a default frequency.
[param amplitude] is the amplitude of the pulse between [code]0.0[/code] and [code]1.0[/code].
[param duration_sec] is the duration of the pulse in seconds.
[param delay_sec] is a delay in seconds before the pulse is given.
*/
func (self Go) TriggerHapticPulse(action_name string, frequency float64, amplitude float64, duration_sec float64, delay_sec float64) {
	class(self).TriggerHapticPulse(gd.NewString(action_name), gd.Float(frequency), gd.Float(amplitude), gd.Float(duration_sec), gd.Float(delay_sec))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.XRNode3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRNode3D"))
	return Go{classdb.XRNode3D(object)}
}

func (self Go) Tracker() string {
		return string(class(self).GetTracker().String())
}

func (self Go) SetTracker(value string) {
	class(self).SetTracker(gd.NewStringName(value))
}

func (self Go) Pose() string {
		return string(class(self).GetPoseName().String())
}

func (self Go) SetPose(value string) {
	class(self).SetPoseName(gd.NewStringName(value))
}

func (self Go) ShowWhenTracked() bool {
		return bool(class(self).GetShowWhenTracked())
}

func (self Go) SetShowWhenTracked(value bool) {
	class(self).SetShowWhenTracked(value)
}

//go:nosplit
func (self class) SetTracker(tracker_name gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(tracker_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRNode3D.Bind_set_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTracker() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRNode3D.Bind_get_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPoseName(pose gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(pose))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRNode3D.Bind_set_pose_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPoseName() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRNode3D.Bind_get_pose_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowWhenTracked(show bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, show)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRNode3D.Bind_set_show_when_tracked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShowWhenTracked() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRNode3D.Bind_get_show_when_tracked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the [member tracker] has been registered and the [member pose] is being tracked.
*/
//go:nosplit
func (self class) GetIsActive() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRNode3D.Bind_get_is_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the [member tracker] has current tracking data for the [member pose] being tracked.
*/
//go:nosplit
func (self class) GetHasTrackingData() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRNode3D.Bind_get_has_tracking_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [XRPose] containing the current state of the pose being tracked. This gives access to additional properties of this pose.
*/
//go:nosplit
func (self class) GetPose() gdclass.XRPose {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRNode3D.Bind_get_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.XRPose{classdb.XRPose(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
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
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(action_name))
	callframe.Arg(frame, frequency)
	callframe.Arg(frame, amplitude)
	callframe.Arg(frame, duration_sec)
	callframe.Arg(frame, delay_sec)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRNode3D.Bind_trigger_haptic_pulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnTrackingChanged(cb func(tracking bool)) {
	self[0].AsObject().Connect(gd.NewStringName("tracking_changed"), gd.NewCallable(cb), 0)
}


func (self class) AsXRNode3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRNode3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("XRNode3D", func(ptr gd.Object) any { return classdb.XRNode3D(ptr) })}
