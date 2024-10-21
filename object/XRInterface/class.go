package XRInterface

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class needs to be implemented to make an AR or VR platform available to Godot and these should be implemented as C++ modules or GDExtension modules. Part of the interface is exposed to GDScript so you can detect, enable and configure an AR or VR platform.
Interfaces should be written in such a way that simply enabling them will give us a working setup. You can query the available interfaces through [XRServer].

*/
type Simple [1]classdb.XRInterface
func (self Simple) GetName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetName(gc).String())
}
func (self Simple) GetCapabilities() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCapabilities()))
}
func (self Simple) IsPrimary() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsPrimary())
}
func (self Simple) SetPrimary(primary bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPrimary(primary)
}
func (self Simple) IsInitialized() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsInitialized())
}
func (self Simple) Initialize() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).Initialize())
}
func (self Simple) Uninitialize() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Uninitialize()
}
func (self Simple) GetSystemInfo() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).GetSystemInfo(gc))
}
func (self Simple) GetTrackingStatus() classdb.XRInterfaceTrackingStatus {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.XRInterfaceTrackingStatus(Expert(self).GetTrackingStatus())
}
func (self Simple) GetRenderTargetSize() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetRenderTargetSize())
}
func (self Simple) GetViewCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetViewCount()))
}
func (self Simple) TriggerHapticPulse(action_name string, tracker_name string, frequency float64, amplitude float64, duration_sec float64, delay_sec float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).TriggerHapticPulse(gc.String(action_name), gc.StringName(tracker_name), gd.Float(frequency), gd.Float(amplitude), gd.Float(duration_sec), gd.Float(delay_sec))
}
func (self Simple) SupportsPlayAreaMode(mode classdb.XRInterfacePlayAreaMode) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).SupportsPlayAreaMode(mode))
}
func (self Simple) GetPlayAreaMode() classdb.XRInterfacePlayAreaMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.XRInterfacePlayAreaMode(Expert(self).GetPlayAreaMode())
}
func (self Simple) SetPlayAreaMode(mode classdb.XRInterfacePlayAreaMode) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).SetPlayAreaMode(mode))
}
func (self Simple) GetPlayArea() gd.PackedVector3Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector3Array(Expert(self).GetPlayArea(gc))
}
func (self Simple) GetAnchorDetectionIsEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetAnchorDetectionIsEnabled())
}
func (self Simple) SetAnchorDetectionIsEnabled(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAnchorDetectionIsEnabled(enable)
}
func (self Simple) GetCameraFeedId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCameraFeedId()))
}
func (self Simple) IsPassthroughSupported() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsPassthroughSupported())
}
func (self Simple) IsPassthroughEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsPassthroughEnabled())
}
func (self Simple) StartPassthrough() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).StartPassthrough())
}
func (self Simple) StopPassthrough() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).StopPassthrough()
}
func (self Simple) GetTransformForView(view int, cam_transform gd.Transform3D) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).GetTransformForView(gd.Int(view), cam_transform))
}
func (self Simple) GetProjectionForView(view int, aspect float64, near float64, far float64) gd.Projection {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Projection(Expert(self).GetProjectionForView(gd.Int(view), gd.Float(aspect), gd.Float(near), gd.Float(far)))
}
func (self Simple) GetSupportedEnvironmentBlendModes() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).GetSupportedEnvironmentBlendModes(gc))
}
func (self Simple) SetEnvironmentBlendMode(mode classdb.XRInterfaceEnvironmentBlendMode) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).SetEnvironmentBlendMode(mode))
}
func (self Simple) GetEnvironmentBlendMode() classdb.XRInterfaceEnvironmentBlendMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.XRInterfaceEnvironmentBlendMode(Expert(self).GetEnvironmentBlendMode())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.XRInterface
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the name of this interface ([code]"OpenXR"[/code], [code]"OpenVR"[/code], [code]"OpenHMD"[/code], [code]"ARKit"[/code], etc.).
*/
//go:nosplit
func (self class) GetName(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a combination of [enum Capabilities] flags providing information about the capabilities of this interface.
*/
//go:nosplit
func (self class) GetCapabilities() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_get_capabilities, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) IsPrimary() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_is_primary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPrimary(primary bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, primary)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_set_primary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if this interface has been initialized.
*/
//go:nosplit
func (self class) IsInitialized() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_is_initialized, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Call this to initialize this interface. The first interface that is initialized is identified as the primary interface and it will be used for rendering output.
After initializing the interface you want to use you then need to enable the AR/VR mode of a viewport and rendering should commence.
[b]Note:[/b] You must enable the XR mode on the main viewport for any device that uses the main output of Godot, such as for mobile VR.
If you do this for a platform that handles its own output (such as OpenVR) Godot will show just one eye without distortion on screen. Alternatively, you can add a separate viewport node to your scene and enable AR/VR on that viewport. It will be used to output to the HMD, leaving you free to do anything you like in the main window, such as using a separate camera as a spectator camera or rendering something completely different.
While currently not used, you can activate additional interfaces. You may wish to do this if you want to track controllers from other platforms. However, at this point in time only one interface can render to an HMD.
*/
//go:nosplit
func (self class) Initialize() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_initialize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Turns the interface off.
*/
//go:nosplit
func (self class) Uninitialize()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_uninitialize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [Dictionary] with extra system info. Interfaces are expected to return [code]XRRuntimeName[/code] and [code]XRRuntimeVersion[/code] providing info about the used XR runtime. Additional entries may be provided specific to an interface.
[b]Note:[/b]This information may only be available after [method initialize] was successfully called.
*/
//go:nosplit
func (self class) GetSystemInfo(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_get_system_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
If supported, returns the status of our tracking. This will allow you to provide feedback to the user whether there are issues with positional tracking.
*/
//go:nosplit
func (self class) GetTrackingStatus() classdb.XRInterfaceTrackingStatus {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.XRInterfaceTrackingStatus](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_get_tracking_status, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the resolution at which we should render our intermediate results before things like lens distortion are applied by the VR platform.
*/
//go:nosplit
func (self class) GetRenderTargetSize() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_get_render_target_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of views that need to be rendered for this device. 1 for Monoscopic, 2 for Stereoscopic.
*/
//go:nosplit
func (self class) GetViewCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_get_view_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Triggers a haptic pulse on a device associated with this interface.
[param action_name] is the name of the action for this pulse.
[param tracker_name] is optional and can be used to direct the pulse to a specific device provided that device is bound to this haptic.
[param frequency] is the frequency of the pulse, set to [code]0.0[/code] to have the system use a default frequency.
[param amplitude] is the amplitude of the pulse between [code]0.0[/code] and [code]1.0[/code].
[param duration_sec] is the duration of the pulse in seconds.
[param delay_sec] is a delay in seconds before the pulse is given.
*/
//go:nosplit
func (self class) TriggerHapticPulse(action_name gd.String, tracker_name gd.StringName, frequency gd.Float, amplitude gd.Float, duration_sec gd.Float, delay_sec gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action_name))
	callframe.Arg(frame, mmm.Get(tracker_name))
	callframe.Arg(frame, frequency)
	callframe.Arg(frame, amplitude)
	callframe.Arg(frame, duration_sec)
	callframe.Arg(frame, delay_sec)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_trigger_haptic_pulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Call this to find out if a given play area mode is supported by this interface.
*/
//go:nosplit
func (self class) SupportsPlayAreaMode(mode classdb.XRInterfacePlayAreaMode) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_supports_play_area_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetPlayAreaMode() classdb.XRInterfacePlayAreaMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.XRInterfacePlayAreaMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_get_play_area_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the active play area mode, will return [code]false[/code] if the mode can't be used with this interface.
[b]Note:[/b] Changing this after the interface has already been initialized can be jarring for the player, so it's recommended to recenter on the HMD with [method XRServer.center_on_hmd] (if switching to [constant XRInterface.XR_PLAY_AREA_STAGE]) or make the switch during a scene change.
*/
//go:nosplit
func (self class) SetPlayAreaMode(mode classdb.XRInterfacePlayAreaMode) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_set_play_area_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array of vectors that represent the physical play area mapped to the virtual space around the [XROrigin3D] point. The points form a convex polygon that can be used to react to or visualize the play area. This returns an empty array if this feature is not supported or if the information is not yet available.
*/
//go:nosplit
func (self class) GetPlayArea(ctx gd.Lifetime) gd.PackedVector3Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_get_play_area, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector3Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetAnchorDetectionIsEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_get_anchor_detection_is_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAnchorDetectionIsEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_set_anchor_detection_is_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If this is an AR interface that requires displaying a camera feed as the background, this method returns the feed ID in the [CameraServer] for this interface.
*/
//go:nosplit
func (self class) GetCameraFeedId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_get_camera_feed_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this interface supports passthrough.
*/
//go:nosplit
func (self class) IsPassthroughSupported() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_is_passthrough_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if passthrough is enabled.
*/
//go:nosplit
func (self class) IsPassthroughEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_is_passthrough_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Starts passthrough, will return [code]false[/code] if passthrough couldn't be started.
[b]Note:[/b] The viewport used for XR must have a transparent background, otherwise passthrough may not properly render.
*/
//go:nosplit
func (self class) StartPassthrough() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_start_passthrough, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Stops passthrough.
*/
//go:nosplit
func (self class) StopPassthrough()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_stop_passthrough, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the transform for a view/eye.
[param view] is the view/eye index.
[param cam_transform] is the transform that maps device coordinates to scene coordinates, typically the [member Node3D.global_transform] of the current XROrigin3D.
*/
//go:nosplit
func (self class) GetTransformForView(view gd.Int, cam_transform gd.Transform3D) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, view)
	callframe.Arg(frame, cam_transform)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_get_transform_for_view, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the projection matrix for a view/eye.
*/
//go:nosplit
func (self class) GetProjectionForView(view gd.Int, aspect gd.Float, near gd.Float, far gd.Float) gd.Projection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, view)
	callframe.Arg(frame, aspect)
	callframe.Arg(frame, near)
	callframe.Arg(frame, far)
	var r_ret = callframe.Ret[gd.Projection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_get_projection_for_view, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the an array of supported environment blend modes, see [enum XRInterface.EnvironmentBlendMode].
*/
//go:nosplit
func (self class) GetSupportedEnvironmentBlendModes(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_get_supported_environment_blend_modes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the active environment blend mode.
[param mode] is the environment blend mode starting with the next frame.
[b]Note:[/b] Not all runtimes support all environment blend modes, so it is important to check this at startup. For example:
[codeblock]
func _ready():
    var xr_interface: XRInterface = XRServer.find_interface("OpenXR")
    if xr_interface and xr_interface.is_initialized():
        var vp: Viewport = get_viewport()
        vp.use_xr = true
        var acceptable_modes = [XRInterface.XR_ENV_BLEND_MODE_OPAQUE, XRInterface.XR_ENV_BLEND_MODE_ADDITIVE]
        var modes = xr_interface.get_supported_environment_blend_modes()
        for mode in acceptable_modes:
            if mode in modes:
                xr_interface.set_environment_blend_mode(mode)
                break
[/codeblock]
*/
//go:nosplit
func (self class) SetEnvironmentBlendMode(mode classdb.XRInterfaceEnvironmentBlendMode) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_set_environment_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetEnvironmentBlendMode() classdb.XRInterfaceEnvironmentBlendMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.XRInterfaceEnvironmentBlendMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterface.Bind_get_environment_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsXRInterface() Expert { return self[0].AsXRInterface() }


//go:nosplit
func (self Simple) AsXRInterface() Simple { return self[0].AsXRInterface() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("XRInterface", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Capabilities = classdb.XRInterfaceCapabilities

const (
/*No XR capabilities.*/
	XrNone Capabilities = 0
/*This interface can work with normal rendering output (non-HMD based AR).*/
	XrMono Capabilities = 1
/*This interface supports stereoscopic rendering.*/
	XrStereo Capabilities = 2
/*This interface supports quad rendering (not yet supported by Godot).*/
	XrQuad Capabilities = 4
/*This interface supports VR.*/
	XrVr Capabilities = 8
/*This interface supports AR (video background and real world tracking).*/
	XrAr Capabilities = 16
/*This interface outputs to an external device. If the main viewport is used, the on screen output is an unmodified buffer of either the left or right eye (stretched if the viewport size is not changed to the same aspect ratio of [method get_render_target_size]). Using a separate viewport node frees up the main viewport for other purposes.*/
	XrExternal Capabilities = 32
)
type TrackingStatus = classdb.XRInterfaceTrackingStatus

const (
/*Tracking is behaving as expected.*/
	XrNormalTracking TrackingStatus = 0
/*Tracking is hindered by excessive motion (the player is moving faster than tracking can keep up).*/
	XrExcessiveMotion TrackingStatus = 1
/*Tracking is hindered by insufficient features, it's too dark (for camera-based tracking), player is blocked, etc.*/
	XrInsufficientFeatures TrackingStatus = 2
/*We don't know the status of the tracking or this interface does not provide feedback.*/
	XrUnknownTracking TrackingStatus = 3
/*Tracking is not functional (camera not plugged in or obscured, lighthouses turned off, etc.).*/
	XrNotTracking TrackingStatus = 4
)
type PlayAreaMode = classdb.XRInterfacePlayAreaMode

const (
/*Play area mode not set or not available.*/
	XrPlayAreaUnknown PlayAreaMode = 0
/*Play area only supports orientation tracking, no positional tracking, area will center around player.*/
	XrPlayArea3dof PlayAreaMode = 1
/*Player is in seated position, limited positional tracking, fixed guardian around player.*/
	XrPlayAreaSitting PlayAreaMode = 2
/*Player is free to move around, full positional tracking.*/
	XrPlayAreaRoomscale PlayAreaMode = 3
/*Same as [constant XR_PLAY_AREA_ROOMSCALE] but origin point is fixed to the center of the physical space. In this mode, system-level recentering may be disabled, requiring the use of [method XRServer.center_on_hmd].*/
	XrPlayAreaStage PlayAreaMode = 4
)
type EnvironmentBlendMode = classdb.XRInterfaceEnvironmentBlendMode

const (
/*Opaque blend mode. This is typically used for VR devices.*/
	XrEnvBlendModeOpaque EnvironmentBlendMode = 0
/*Additive blend mode. This is typically used for AR devices or VR devices with passthrough.*/
	XrEnvBlendModeAdditive EnvironmentBlendMode = 1
/*Alpha blend mode. This is typically used for AR or VR devices with passthrough capabilities. The alpha channel controls how much of the passthrough is visible. Alpha of 0.0 means the passthrough is visible and this pixel works in ADDITIVE mode. Alpha of 1.0 means that the passthrough is not visible and this pixel works in OPAQUE mode.*/
	XrEnvBlendModeAlphaBlend EnvironmentBlendMode = 2
)
