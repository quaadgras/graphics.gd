// Package XRInterface provides methods for working with XRInterface object instances.
package XRInterface

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Projection"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Vector3"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
This class needs to be implemented to make an AR or VR platform available to Godot and these should be implemented as C++ modules or GDExtension modules. Part of the interface is exposed to GDScript so you can detect, enable and configure an AR or VR platform.
Interfaces should be written in such a way that simply enabling them will give us a working setup. You can query the available interfaces through [XRServer].
*/
type Instance [1]gdclass.XRInterface

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsXRInterface() Instance
}

/*
Returns the name of this interface ([code]"OpenXR"[/code], [code]"OpenVR"[/code], [code]"OpenHMD"[/code], [code]"ARKit"[/code], etc.).
*/
func (self Instance) GetName() string { //gd:XRInterface.get_name
	return string(class(self).GetName().String())
}

/*
Returns a combination of [enum Capabilities] flags providing information about the capabilities of this interface.
*/
func (self Instance) GetCapabilities() int { //gd:XRInterface.get_capabilities
	return int(int(class(self).GetCapabilities()))
}

/*
Returns [code]true[/code] if this interface has been initialized.
*/
func (self Instance) IsInitialized() bool { //gd:XRInterface.is_initialized
	return bool(class(self).IsInitialized())
}

/*
Call this to initialize this interface. The first interface that is initialized is identified as the primary interface and it will be used for rendering output.
After initializing the interface you want to use you then need to enable the AR/VR mode of a viewport and rendering should commence.
[b]Note:[/b] You must enable the XR mode on the main viewport for any device that uses the main output of Godot, such as for mobile VR.
If you do this for a platform that handles its own output (such as OpenVR) Godot will show just one eye without distortion on screen. Alternatively, you can add a separate viewport node to your scene and enable AR/VR on that viewport. It will be used to output to the HMD, leaving you free to do anything you like in the main window, such as using a separate camera as a spectator camera or rendering something completely different.
While currently not used, you can activate additional interfaces. You may wish to do this if you want to track controllers from other platforms. However, at this point in time only one interface can render to an HMD.
*/
func (self Instance) Initialize() bool { //gd:XRInterface.initialize
	return bool(class(self).Initialize())
}

/*
Turns the interface off.
*/
func (self Instance) Uninitialize() { //gd:XRInterface.uninitialize
	class(self).Uninitialize()
}

/*
Returns a [Dictionary] with extra system info. Interfaces are expected to return [code]XRRuntimeName[/code] and [code]XRRuntimeVersion[/code] providing info about the used XR runtime. Additional entries may be provided specific to an interface.
[b]Note:[/b]This information may only be available after [method initialize] was successfully called.
*/
func (self Instance) GetSystemInfo() map[string]interface{} { //gd:XRInterface.get_system_info
	return map[string]interface{}(gd.DictionaryAs[map[string]interface{}](class(self).GetSystemInfo()))
}

/*
If supported, returns the status of our tracking. This will allow you to provide feedback to the user whether there are issues with positional tracking.
*/
func (self Instance) GetTrackingStatus() gdclass.XRInterfaceTrackingStatus { //gd:XRInterface.get_tracking_status
	return gdclass.XRInterfaceTrackingStatus(class(self).GetTrackingStatus())
}

/*
Returns the resolution at which we should render our intermediate results before things like lens distortion are applied by the VR platform.
*/
func (self Instance) GetRenderTargetSize() Vector2.XY { //gd:XRInterface.get_render_target_size
	return Vector2.XY(class(self).GetRenderTargetSize())
}

/*
Returns the number of views that need to be rendered for this device. 1 for Monoscopic, 2 for Stereoscopic.
*/
func (self Instance) GetViewCount() int { //gd:XRInterface.get_view_count
	return int(int(class(self).GetViewCount()))
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
func (self Instance) TriggerHapticPulse(action_name string, tracker_name string, frequency Float.X, amplitude Float.X, duration_sec Float.X, delay_sec Float.X) { //gd:XRInterface.trigger_haptic_pulse
	class(self).TriggerHapticPulse(String.New(action_name), String.Name(String.New(tracker_name)), float64(frequency), float64(amplitude), float64(duration_sec), float64(delay_sec))
}

/*
Call this to find out if a given play area mode is supported by this interface.
*/
func (self Instance) SupportsPlayAreaMode(mode gdclass.XRInterfacePlayAreaMode) bool { //gd:XRInterface.supports_play_area_mode
	return bool(class(self).SupportsPlayAreaMode(mode))
}

/*
Returns an array of vectors that represent the physical play area mapped to the virtual space around the [XROrigin3D] point. The points form a convex polygon that can be used to react to or visualize the play area. This returns an empty array if this feature is not supported or if the information is not yet available.
*/
func (self Instance) GetPlayArea() []Vector3.XYZ { //gd:XRInterface.get_play_area
	return []Vector3.XYZ(slices.Collect(class(self).GetPlayArea().Values()))
}

/*
If this is an AR interface that requires displaying a camera feed as the background, this method returns the feed ID in the [CameraServer] for this interface.
*/
func (self Instance) GetCameraFeedId() int { //gd:XRInterface.get_camera_feed_id
	return int(int(class(self).GetCameraFeedId()))
}

/*
Returns [code]true[/code] if this interface supports passthrough.
*/
func (self Instance) IsPassthroughSupported() bool { //gd:XRInterface.is_passthrough_supported
	return bool(class(self).IsPassthroughSupported())
}

/*
Returns [code]true[/code] if passthrough is enabled.
*/
func (self Instance) IsPassthroughEnabled() bool { //gd:XRInterface.is_passthrough_enabled
	return bool(class(self).IsPassthroughEnabled())
}

/*
Starts passthrough, will return [code]false[/code] if passthrough couldn't be started.
[b]Note:[/b] The viewport used for XR must have a transparent background, otherwise passthrough may not properly render.
*/
func (self Instance) StartPassthrough() bool { //gd:XRInterface.start_passthrough
	return bool(class(self).StartPassthrough())
}

/*
Stops passthrough.
*/
func (self Instance) StopPassthrough() { //gd:XRInterface.stop_passthrough
	class(self).StopPassthrough()
}

/*
Returns the transform for a view/eye.
[param view] is the view/eye index.
[param cam_transform] is the transform that maps device coordinates to scene coordinates, typically the [member Node3D.global_transform] of the current XROrigin3D.
*/
func (self Instance) GetTransformForView(view int, cam_transform Transform3D.BasisOrigin) Transform3D.BasisOrigin { //gd:XRInterface.get_transform_for_view
	return Transform3D.BasisOrigin(class(self).GetTransformForView(int64(view), Transform3D.BasisOrigin(cam_transform)))
}

/*
Returns the projection matrix for a view/eye.
*/
func (self Instance) GetProjectionForView(view int, aspect Float.X, near Float.X, far Float.X) Projection.XYZW { //gd:XRInterface.get_projection_for_view
	return Projection.XYZW(class(self).GetProjectionForView(int64(view), float64(aspect), float64(near), float64(far)))
}

/*
Returns the an array of supported environment blend modes, see [enum XRInterface.EnvironmentBlendMode].
*/
func (self Instance) GetSupportedEnvironmentBlendModes() []any { //gd:XRInterface.get_supported_environment_blend_modes
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).GetSupportedEnvironmentBlendModes())))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.XRInterface

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRInterface"))
	casted := Instance{*(*gdclass.XRInterface)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) InterfaceIsPrimary() bool {
	return bool(class(self).IsPrimary())
}

func (self Instance) SetInterfaceIsPrimary(value bool) {
	class(self).SetPrimary(value)
}

func (self Instance) XrPlayAreaMode() gdclass.XRInterfacePlayAreaMode {
	return gdclass.XRInterfacePlayAreaMode(class(self).GetPlayAreaMode())
}

func (self Instance) SetXrPlayAreaMode(value gdclass.XRInterfacePlayAreaMode) {
	class(self).SetPlayAreaMode(value)
}

func (self Instance) EnvironmentBlendMode() gdclass.XRInterfaceEnvironmentBlendMode {
	return gdclass.XRInterfaceEnvironmentBlendMode(class(self).GetEnvironmentBlendMode())
}

func (self Instance) SetEnvironmentBlendMode(value gdclass.XRInterfaceEnvironmentBlendMode) {
	class(self).SetEnvironmentBlendMode(value)
}

func (self Instance) ArIsAnchorDetectionEnabled() bool {
	return bool(class(self).GetAnchorDetectionIsEnabled())
}

func (self Instance) SetArIsAnchorDetectionEnabled(value bool) {
	class(self).SetAnchorDetectionIsEnabled(value)
}

/*
Returns the name of this interface ([code]"OpenXR"[/code], [code]"OpenVR"[/code], [code]"OpenHMD"[/code], [code]"ARKit"[/code], etc.).
*/
//go:nosplit
func (self class) GetName() String.Name { //gd:XRInterface.get_name
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns a combination of [enum Capabilities] flags providing information about the capabilities of this interface.
*/
//go:nosplit
func (self class) GetCapabilities() int64 { //gd:XRInterface.get_capabilities
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_get_capabilities, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) IsPrimary() bool { //gd:XRInterface.is_primary
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_is_primary, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPrimary(primary bool) { //gd:XRInterface.set_primary
	var frame = callframe.New()
	callframe.Arg(frame, primary)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_set_primary, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if this interface has been initialized.
*/
//go:nosplit
func (self class) IsInitialized() bool { //gd:XRInterface.is_initialized
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_is_initialized, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) Initialize() bool { //gd:XRInterface.initialize
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_initialize, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Turns the interface off.
*/
//go:nosplit
func (self class) Uninitialize() { //gd:XRInterface.uninitialize
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_uninitialize, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a [Dictionary] with extra system info. Interfaces are expected to return [code]XRRuntimeName[/code] and [code]XRRuntimeVersion[/code] providing info about the used XR runtime. Additional entries may be provided specific to an interface.
[b]Note:[/b]This information may only be available after [method initialize] was successfully called.
*/
//go:nosplit
func (self class) GetSystemInfo() Dictionary.Any { //gd:XRInterface.get_system_info
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_get_system_info, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
If supported, returns the status of our tracking. This will allow you to provide feedback to the user whether there are issues with positional tracking.
*/
//go:nosplit
func (self class) GetTrackingStatus() gdclass.XRInterfaceTrackingStatus { //gd:XRInterface.get_tracking_status
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.XRInterfaceTrackingStatus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_get_tracking_status, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the resolution at which we should render our intermediate results before things like lens distortion are applied by the VR platform.
*/
//go:nosplit
func (self class) GetRenderTargetSize() Vector2.XY { //gd:XRInterface.get_render_target_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_get_render_target_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of views that need to be rendered for this device. 1 for Monoscopic, 2 for Stereoscopic.
*/
//go:nosplit
func (self class) GetViewCount() int64 { //gd:XRInterface.get_view_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_get_view_count, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) TriggerHapticPulse(action_name String.Readable, tracker_name String.Name, frequency float64, amplitude float64, duration_sec float64, delay_sec float64) { //gd:XRInterface.trigger_haptic_pulse
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(action_name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(tracker_name)))
	callframe.Arg(frame, frequency)
	callframe.Arg(frame, amplitude)
	callframe.Arg(frame, duration_sec)
	callframe.Arg(frame, delay_sec)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_trigger_haptic_pulse, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Call this to find out if a given play area mode is supported by this interface.
*/
//go:nosplit
func (self class) SupportsPlayAreaMode(mode gdclass.XRInterfacePlayAreaMode) bool { //gd:XRInterface.supports_play_area_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_supports_play_area_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetPlayAreaMode() gdclass.XRInterfacePlayAreaMode { //gd:XRInterface.get_play_area_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.XRInterfacePlayAreaMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_get_play_area_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the active play area mode, will return [code]false[/code] if the mode can't be used with this interface.
[b]Note:[/b] Changing this after the interface has already been initialized can be jarring for the player, so it's recommended to recenter on the HMD with [method XRServer.center_on_hmd] (if switching to [constant XRInterface.XR_PLAY_AREA_STAGE]) or make the switch during a scene change.
*/
//go:nosplit
func (self class) SetPlayAreaMode(mode gdclass.XRInterfacePlayAreaMode) bool { //gd:XRInterface.set_play_area_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_set_play_area_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array of vectors that represent the physical play area mapped to the virtual space around the [XROrigin3D] point. The points form a convex polygon that can be used to react to or visualize the play area. This returns an empty array if this feature is not supported or if the information is not yet available.
*/
//go:nosplit
func (self class) GetPlayArea() Packed.Array[Vector3.XYZ] { //gd:XRInterface.get_play_area
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_get_play_area, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector3.XYZ](Array.Through(gd.PackedProxy[gd.PackedVector3Array, Vector3.XYZ]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetAnchorDetectionIsEnabled() bool { //gd:XRInterface.get_anchor_detection_is_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_get_anchor_detection_is_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAnchorDetectionIsEnabled(enable bool) { //gd:XRInterface.set_anchor_detection_is_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_set_anchor_detection_is_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
If this is an AR interface that requires displaying a camera feed as the background, this method returns the feed ID in the [CameraServer] for this interface.
*/
//go:nosplit
func (self class) GetCameraFeedId() int64 { //gd:XRInterface.get_camera_feed_id
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_get_camera_feed_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this interface supports passthrough.
*/
//go:nosplit
func (self class) IsPassthroughSupported() bool { //gd:XRInterface.is_passthrough_supported
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_is_passthrough_supported, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if passthrough is enabled.
*/
//go:nosplit
func (self class) IsPassthroughEnabled() bool { //gd:XRInterface.is_passthrough_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_is_passthrough_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Starts passthrough, will return [code]false[/code] if passthrough couldn't be started.
[b]Note:[/b] The viewport used for XR must have a transparent background, otherwise passthrough may not properly render.
*/
//go:nosplit
func (self class) StartPassthrough() bool { //gd:XRInterface.start_passthrough
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_start_passthrough, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Stops passthrough.
*/
//go:nosplit
func (self class) StopPassthrough() { //gd:XRInterface.stop_passthrough
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_stop_passthrough, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the transform for a view/eye.
[param view] is the view/eye index.
[param cam_transform] is the transform that maps device coordinates to scene coordinates, typically the [member Node3D.global_transform] of the current XROrigin3D.
*/
//go:nosplit
func (self class) GetTransformForView(view int64, cam_transform Transform3D.BasisOrigin) Transform3D.BasisOrigin { //gd:XRInterface.get_transform_for_view
	var frame = callframe.New()
	callframe.Arg(frame, view)
	callframe.Arg(frame, cam_transform)
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_get_transform_for_view, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the projection matrix for a view/eye.
*/
//go:nosplit
func (self class) GetProjectionForView(view int64, aspect float64, near float64, far float64) Projection.XYZW { //gd:XRInterface.get_projection_for_view
	var frame = callframe.New()
	callframe.Arg(frame, view)
	callframe.Arg(frame, aspect)
	callframe.Arg(frame, near)
	callframe.Arg(frame, far)
	var r_ret = callframe.Ret[Projection.XYZW](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_get_projection_for_view, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the an array of supported environment blend modes, see [enum XRInterface.EnvironmentBlendMode].
*/
//go:nosplit
func (self class) GetSupportedEnvironmentBlendModes() Array.Any { //gd:XRInterface.get_supported_environment_blend_modes
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_get_supported_environment_blend_modes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
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
func (self class) SetEnvironmentBlendMode(mode gdclass.XRInterfaceEnvironmentBlendMode) bool { //gd:XRInterface.set_environment_blend_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_set_environment_blend_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetEnvironmentBlendMode() gdclass.XRInterfaceEnvironmentBlendMode { //gd:XRInterface.get_environment_blend_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.XRInterfaceEnvironmentBlendMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterface.Bind_get_environment_blend_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnPlayAreaChanged(cb func(mode int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("play_area_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsXRInterface() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsXRInterface() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("XRInterface", func(ptr gd.Object) any { return [1]gdclass.XRInterface{*(*gdclass.XRInterface)(unsafe.Pointer(&ptr))} })
}

type Capabilities = gdclass.XRInterfaceCapabilities //gd:XRInterface.Capabilities

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

type TrackingStatus = gdclass.XRInterfaceTrackingStatus //gd:XRInterface.TrackingStatus

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

type PlayAreaMode = gdclass.XRInterfacePlayAreaMode //gd:XRInterface.PlayAreaMode

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

type EnvironmentBlendMode = gdclass.XRInterfaceEnvironmentBlendMode //gd:XRInterface.EnvironmentBlendMode

const (
	/*Opaque blend mode. This is typically used for VR devices.*/
	XrEnvBlendModeOpaque EnvironmentBlendMode = 0
	/*Additive blend mode. This is typically used for AR devices or VR devices with passthrough.*/
	XrEnvBlendModeAdditive EnvironmentBlendMode = 1
	/*Alpha blend mode. This is typically used for AR or VR devices with passthrough capabilities. The alpha channel controls how much of the passthrough is visible. Alpha of 0.0 means the passthrough is visible and this pixel works in ADDITIVE mode. Alpha of 1.0 means that the passthrough is not visible and this pixel works in OPAQUE mode.*/
	XrEnvBlendModeAlphaBlend EnvironmentBlendMode = 2
)
