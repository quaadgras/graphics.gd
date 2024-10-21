package OpenXRInterface

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/XRInterface"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The OpenXR interface allows Godot to interact with OpenXR runtimes and make it possible to create XR experiences and games.
Due to the needs of OpenXR this interface works slightly different than other plugin based XR interfaces. It needs to be initialized when Godot starts. You need to enable OpenXR, settings for this can be found in your games project settings under the XR heading. You do need to mark a viewport for use with XR in order for Godot to know which render result should be output to the headset.

*/
type Simple [1]classdb.OpenXRInterface
func (self Simple) GetDisplayRefreshRate() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDisplayRefreshRate()))
}
func (self Simple) SetDisplayRefreshRate(refresh_rate float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDisplayRefreshRate(gd.Float(refresh_rate))
}
func (self Simple) GetRenderTargetSizeMultiplier() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetRenderTargetSizeMultiplier()))
}
func (self Simple) SetRenderTargetSizeMultiplier(multiplier float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRenderTargetSizeMultiplier(gd.Float(multiplier))
}
func (self Simple) IsFoveationSupported() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsFoveationSupported())
}
func (self Simple) GetFoveationLevel() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetFoveationLevel()))
}
func (self Simple) SetFoveationLevel(foveation_level int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFoveationLevel(gd.Int(foveation_level))
}
func (self Simple) GetFoveationDynamic() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetFoveationDynamic())
}
func (self Simple) SetFoveationDynamic(foveation_dynamic bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFoveationDynamic(foveation_dynamic)
}
func (self Simple) IsActionSetActive(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsActionSetActive(gc.String(name)))
}
func (self Simple) SetActionSetActive(name string, active bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetActionSetActive(gc.String(name), active)
}
func (self Simple) GetActionSets() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).GetActionSets(gc))
}
func (self Simple) GetAvailableDisplayRefreshRates() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).GetAvailableDisplayRefreshRates(gc))
}
func (self Simple) SetMotionRange(hand classdb.OpenXRInterfaceHand, motion_range classdb.OpenXRInterfaceHandMotionRange) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMotionRange(hand, motion_range)
}
func (self Simple) GetMotionRange(hand classdb.OpenXRInterfaceHand) classdb.OpenXRInterfaceHandMotionRange {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.OpenXRInterfaceHandMotionRange(Expert(self).GetMotionRange(hand))
}
func (self Simple) GetHandTrackingSource(hand classdb.OpenXRInterfaceHand) classdb.OpenXRInterfaceHandTrackedSource {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.OpenXRInterfaceHandTrackedSource(Expert(self).GetHandTrackingSource(hand))
}
func (self Simple) GetHandJointFlags(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) classdb.OpenXRInterfaceHandJointFlags {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.OpenXRInterfaceHandJointFlags(Expert(self).GetHandJointFlags(hand, joint))
}
func (self Simple) GetHandJointRotation(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) gd.Quaternion {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Quaternion(Expert(self).GetHandJointRotation(hand, joint))
}
func (self Simple) GetHandJointPosition(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetHandJointPosition(hand, joint))
}
func (self Simple) GetHandJointRadius(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetHandJointRadius(hand, joint)))
}
func (self Simple) GetHandJointLinearVelocity(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetHandJointLinearVelocity(hand, joint))
}
func (self Simple) GetHandJointAngularVelocity(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetHandJointAngularVelocity(hand, joint))
}
func (self Simple) IsHandTrackingSupported() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsHandTrackingSupported())
}
func (self Simple) IsHandInteractionSupported() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsHandInteractionSupported())
}
func (self Simple) IsEyeGazeInteractionSupported() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsEyeGazeInteractionSupported())
}
func (self Simple) GetVrsMinRadius() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVrsMinRadius()))
}
func (self Simple) SetVrsMinRadius(radius float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVrsMinRadius(gd.Float(radius))
}
func (self Simple) GetVrsStrength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVrsStrength()))
}
func (self Simple) SetVrsStrength(strength float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVrsStrength(gd.Float(strength))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.OpenXRInterface
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) GetDisplayRefreshRate() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_display_refresh_rate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDisplayRefreshRate(refresh_rate gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, refresh_rate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_set_display_refresh_rate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRenderTargetSizeMultiplier() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_render_target_size_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRenderTargetSizeMultiplier(multiplier gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multiplier)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_set_render_target_size_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if OpenXR's foveation extension is supported, the interface must be initialized before this returns a valid value.
[b]Note:[/b] This feature is only available on the compatibility renderer and currently only available on some stand alone headsets. For Vulkan set [member Viewport.vrs_mode] to [code]VRS_XR[/code] on desktop.
*/
//go:nosplit
func (self class) IsFoveationSupported() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_is_foveation_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetFoveationLevel() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_foveation_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFoveationLevel(foveation_level gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, foveation_level)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_set_foveation_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFoveationDynamic() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_foveation_dynamic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFoveationDynamic(foveation_dynamic bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, foveation_dynamic)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_set_foveation_dynamic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the given action set is active.
*/
//go:nosplit
func (self class) IsActionSetActive(name gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_is_action_set_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the given action set as active or inactive.
*/
//go:nosplit
func (self class) SetActionSetActive(name gd.String, active bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_set_action_set_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a list of action sets registered with Godot (loaded from the action map at runtime).
*/
//go:nosplit
func (self class) GetActionSets(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_action_sets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns display refresh rates supported by the current HMD. Only returned if this feature is supported by the OpenXR runtime and after the interface has been initialized.
*/
//go:nosplit
func (self class) GetAvailableDisplayRefreshRates(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_available_display_refresh_rates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
If handtracking is enabled and motion range is supported, sets the currently configured motion range for [param hand] to [param motion_range].
*/
//go:nosplit
func (self class) SetMotionRange(hand classdb.OpenXRInterfaceHand, motion_range classdb.OpenXRInterfaceHandMotionRange)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, motion_range)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_set_motion_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If handtracking is enabled and motion range is supported, gets the currently configured motion range for [param hand].
*/
//go:nosplit
func (self class) GetMotionRange(hand classdb.OpenXRInterfaceHand) classdb.OpenXRInterfaceHandMotionRange {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	var r_ret = callframe.Ret[classdb.OpenXRInterfaceHandMotionRange](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_motion_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If handtracking is enabled and hand tracking source is supported, gets the source of the hand tracking data for [param hand].
*/
//go:nosplit
func (self class) GetHandTrackingSource(hand classdb.OpenXRInterfaceHand) classdb.OpenXRInterfaceHandTrackedSource {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	var r_ret = callframe.Ret[classdb.OpenXRInterfaceHandTrackedSource](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_hand_tracking_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If handtracking is enabled, returns flags that inform us of the validity of the tracking data.
*/
//go:nosplit
func (self class) GetHandJointFlags(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) classdb.OpenXRInterfaceHandJointFlags {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[classdb.OpenXRInterfaceHandJointFlags](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_hand_joint_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If handtracking is enabled, returns the rotation of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR.
*/
//go:nosplit
func (self class) GetHandJointRotation(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) gd.Quaternion {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Quaternion](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_hand_joint_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If handtracking is enabled, returns the position of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is relative to [XROrigin3D] without worldscale applied!
*/
//go:nosplit
func (self class) GetHandJointPosition(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_hand_joint_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If handtracking is enabled, returns the radius of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is without worldscale applied!
*/
//go:nosplit
func (self class) GetHandJointRadius(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_hand_joint_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If handtracking is enabled, returns the linear velocity of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is relative to [XROrigin3D] without worldscale applied!
*/
//go:nosplit
func (self class) GetHandJointLinearVelocity(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_hand_joint_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If handtracking is enabled, returns the angular velocity of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is relative to [XROrigin3D]!
*/
//go:nosplit
func (self class) GetHandJointAngularVelocity(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_hand_joint_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if OpenXR's hand tracking is supported and enabled.
[b]Note:[/b] This only returns a valid value after OpenXR has been initialized.
*/
//go:nosplit
func (self class) IsHandTrackingSupported() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_is_hand_tracking_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if OpenXR's hand interaction profile is supported and enabled.
[b]Note:[/b] This only returns a valid value after OpenXR has been initialized.
*/
//go:nosplit
func (self class) IsHandInteractionSupported() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_is_hand_interaction_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the capabilities of the eye gaze interaction extension.
[b]Note:[/b] This only returns a valid value after OpenXR has been initialized.
*/
//go:nosplit
func (self class) IsEyeGazeInteractionSupported() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_is_eye_gaze_interaction_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetVrsMinRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_vrs_min_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVrsMinRadius(radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_set_vrs_min_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVrsStrength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_vrs_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVrsStrength(strength gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_set_vrs_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsOpenXRInterface() Expert { return self[0].AsOpenXRInterface() }


//go:nosplit
func (self Simple) AsOpenXRInterface() Simple { return self[0].AsOpenXRInterface() }


//go:nosplit
func (self class) AsXRInterface() XRInterface.Expert { return self[0].AsXRInterface() }


//go:nosplit
func (self Simple) AsXRInterface() XRInterface.Simple { return self[0].AsXRInterface() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("OpenXRInterface", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Hand = classdb.OpenXRInterfaceHand

const (
/*Left hand.*/
	HandLeft Hand = 0
/*Right hand.*/
	HandRight Hand = 1
/*Maximum value for the hand enum.*/
	HandMax Hand = 2
)
type HandMotionRange = classdb.OpenXRInterfaceHandMotionRange

const (
/*Full hand range, if user closes their hands, we make a full fist.*/
	HandMotionRangeUnobstructed HandMotionRange = 0
/*Conform to controller, if user closes their hands, the tracked data conforms to the shape of the controller.*/
	HandMotionRangeConformToController HandMotionRange = 1
/*Maximum value for the motion range enum.*/
	HandMotionRangeMax HandMotionRange = 2
)
type HandTrackedSource = classdb.OpenXRInterfaceHandTrackedSource

const (
/*The source of hand tracking data is unknown (the extension is likely unsupported).*/
	HandTrackedSourceUnknown HandTrackedSource = 0
/*The source of hand tracking is unobstructed, this means that an accurate method of hand tracking is used, e.g. optical hand tracking, data gloves, etc.*/
	HandTrackedSourceUnobstructed HandTrackedSource = 1
/*The source of hand tracking is a controller, bone positions are inferred from controller inputs.*/
	HandTrackedSourceController HandTrackedSource = 2
/*Maximum value for the hand tracked source enum.*/
	HandTrackedSourceMax HandTrackedSource = 3
)
type HandJoints = classdb.OpenXRInterfaceHandJoints

const (
/*Palm joint.*/
	HandJointPalm HandJoints = 0
/*Wrist joint.*/
	HandJointWrist HandJoints = 1
/*Thumb metacarpal joint.*/
	HandJointThumbMetacarpal HandJoints = 2
/*Thumb proximal joint.*/
	HandJointThumbProximal HandJoints = 3
/*Thumb distal joint.*/
	HandJointThumbDistal HandJoints = 4
/*Thumb tip joint.*/
	HandJointThumbTip HandJoints = 5
/*Index metacarpal joint.*/
	HandJointIndexMetacarpal HandJoints = 6
/*Index proximal joint.*/
	HandJointIndexProximal HandJoints = 7
/*Index intermediate joint.*/
	HandJointIndexIntermediate HandJoints = 8
/*Index distal joint.*/
	HandJointIndexDistal HandJoints = 9
/*Index tip joint.*/
	HandJointIndexTip HandJoints = 10
/*Middle metacarpal joint.*/
	HandJointMiddleMetacarpal HandJoints = 11
/*Middle proximal joint.*/
	HandJointMiddleProximal HandJoints = 12
/*Middle intermediate joint.*/
	HandJointMiddleIntermediate HandJoints = 13
/*Middle distal joint.*/
	HandJointMiddleDistal HandJoints = 14
/*Middle tip joint.*/
	HandJointMiddleTip HandJoints = 15
/*Ring metacarpal joint.*/
	HandJointRingMetacarpal HandJoints = 16
/*Ring proximal joint.*/
	HandJointRingProximal HandJoints = 17
/*Ring intermediate joint.*/
	HandJointRingIntermediate HandJoints = 18
/*Ring distal joint.*/
	HandJointRingDistal HandJoints = 19
/*Ring tip joint.*/
	HandJointRingTip HandJoints = 20
/*Little metacarpal joint.*/
	HandJointLittleMetacarpal HandJoints = 21
/*Little proximal joint.*/
	HandJointLittleProximal HandJoints = 22
/*Little intermediate joint.*/
	HandJointLittleIntermediate HandJoints = 23
/*Little distal joint.*/
	HandJointLittleDistal HandJoints = 24
/*Little tip joint.*/
	HandJointLittleTip HandJoints = 25
/*Maximum value for the hand joint enum.*/
	HandJointMax HandJoints = 26
)
type HandJointFlags = classdb.OpenXRInterfaceHandJointFlags

const (
/*No flags are set.*/
	HandJointNone HandJointFlags = 0
/*If set, the orientation data is valid, otherwise, the orientation data is unreliable and should not be used.*/
	HandJointOrientationValid HandJointFlags = 1
/*If set, the orientation data comes from tracking data, otherwise, the orientation data contains predicted data.*/
	HandJointOrientationTracked HandJointFlags = 2
/*If set, the positional data is valid, otherwise, the positional data is unreliable and should not be used.*/
	HandJointPositionValid HandJointFlags = 4
/*If set, the positional data comes from tracking data, otherwise, the positional data contains predicted data.*/
	HandJointPositionTracked HandJointFlags = 8
/*If set, our linear velocity data is valid, otherwise, the linear velocity data is unreliable and should not be used.*/
	HandJointLinearVelocityValid HandJointFlags = 16
/*If set, our angular velocity data is valid, otherwise, the angular velocity data is unreliable and should not be used.*/
	HandJointAngularVelocityValid HandJointFlags = 32
)
