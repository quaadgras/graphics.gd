package XRInterfaceExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/XRInterface"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
External XR interface plugins should inherit from this class.
	// XRInterfaceExtension methods that can be overridden by a [Class] that extends it.
	type XRInterfaceExtension interface {
		//Returns the name of this interface.
		GetName() string
		//Returns the capabilities of this interface.
		GetCapabilities() int
		//Returns [code]true[/code] if this interface has been initialized.
		IsInitialized() bool
		//Initializes the interface, returns [code]true[/code] on success.
		Initialize() bool
		//Uninitialize the interface.
		Uninitialize() 
		//Returns a [Dictionary] with system information related to this interface.
		GetSystemInfo() gd.Dictionary
		//Returns [code]true[/code] if this interface supports this play area mode.
		SupportsPlayAreaMode(mode classdb.XRInterfacePlayAreaMode) bool
		//Returns the play area mode that sets up our play area.
		GetPlayAreaMode() classdb.XRInterfacePlayAreaMode
		//Set the play area mode for this interface.
		SetPlayAreaMode(mode classdb.XRInterfacePlayAreaMode) bool
		//Returns a [PackedVector3Array] that represents the play areas boundaries (if applicable).
		GetPlayArea() []gd.Vector3
		//Returns the size of our render target for this interface, this overrides the size of the [Viewport] marked as the xr viewport.
		GetRenderTargetSize() gd.Vector2
		//Returns the number of views this interface requires, 1 for mono, 2 for stereoscopic.
		GetViewCount() int
		//Returns the [Transform3D] that positions the [XRCamera3D] in the world.
		GetCameraTransform() gd.Transform3D
		//Returns a [Transform3D] for a given view.
		GetTransformForView(view int, cam_transform gd.Transform3D) gd.Transform3D
		//Returns the projection matrix for the given view as a [PackedFloat64Array].
		GetProjectionForView(view int, aspect float64, z_near float64, z_far float64) []float64
		GetVrsTexture() gd.RID
		//Called if this [XRInterfaceExtension] is active before our physics and game process is called. Most XR interfaces will update its [XRPositionalTracker]s at this point in time.
		Process() 
		//Called if this [XRInterfaceExtension] is active before rendering starts. Most XR interfaces will sync tracking at this point in time.
		PreRender() 
		//Called if this is our primary [XRInterfaceExtension] before we start processing a [Viewport] for every active XR [Viewport], returns [code]true[/code] if that viewport should be rendered. An XR interface may return [code]false[/code] if the user has taken off their headset and we can pause rendering.
		PreDrawViewport(render_target gd.RID) bool
		//Called after the XR [Viewport] draw logic has completed.
		PostDrawViewport(render_target gd.RID, screen_rect gd.Rect2) 
		//Called if interface is active and queues have been submitted.
		EndFrame() 
		//Returns a [PackedStringArray] with tracker names configured by this interface. Note that user configuration can override this list.
		GetSuggestedTrackerNames() []string
		//Returns a [PackedStringArray] with pose names configured by this interface. Note that user configuration can override this list.
		GetSuggestedPoseNames(tracker_name string) []string
		//Returns a [enum XRInterface.TrackingStatus] specifying the current status of our tracking.
		GetTrackingStatus() classdb.XRInterfaceTrackingStatus
		//Triggers a haptic pulse to be emitted on the specified tracker.
		TriggerHapticPulse(action_name string, tracker_name string, frequency float64, amplitude float64, duration_sec float64, delay_sec float64) 
		//Return [code]true[/code] if anchor detection is enabled for this interface.
		GetAnchorDetectionIsEnabled() bool
		//Enables anchor detection on this interface if supported.
		SetAnchorDetectionIsEnabled(enabled bool) 
		//Returns the camera feed ID for the [CameraFeed] registered with the [CameraServer] that should be presented as the background on an AR capable device (if applicable).
		GetCameraFeedId() int
		//Return color texture into which to render (if applicable).
		GetColorTexture() gd.RID
		//Return depth texture into which to render (if applicable).
		GetDepthTexture() gd.RID
		//Return velocity texture into which to render (if applicable).
		GetVelocityTexture() gd.RID
	}

*/
type Go [1]classdb.XRInterfaceExtension

/*
Returns the name of this interface.
*/
func (Go) _get_name(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.StringName(ret)))
		gc.End()
	}
}

/*
Returns the capabilities of this interface.
*/
func (Go) _get_capabilities(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Returns [code]true[/code] if this interface has been initialized.
*/
func (Go) _is_initialized(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Initializes the interface, returns [code]true[/code] on success.
*/
func (Go) _initialize(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Uninitialize the interface.
*/
func (Go) _uninitialize(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Returns a [Dictionary] with system information related to this interface.
*/
func (Go) _get_system_info(impl func(ptr unsafe.Pointer) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}

/*
Returns [code]true[/code] if this interface supports this play area mode.
*/
func (Go) _supports_play_area_mode(impl func(ptr unsafe.Pointer, mode classdb.XRInterfacePlayAreaMode) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var mode = gd.UnsafeGet[classdb.XRInterfacePlayAreaMode](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Returns the play area mode that sets up our play area.
*/
func (Go) _get_play_area_mode(impl func(ptr unsafe.Pointer) classdb.XRInterfacePlayAreaMode, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Set the play area mode for this interface.
*/
func (Go) _set_play_area_mode(impl func(ptr unsafe.Pointer, mode classdb.XRInterfacePlayAreaMode) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var mode = gd.UnsafeGet[classdb.XRInterfacePlayAreaMode](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Returns a [PackedVector3Array] that represents the play areas boundaries (if applicable).
*/
func (Go) _get_play_area(impl func(ptr unsafe.Pointer) []gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.PackedVector3Slice(ret)))
		gc.End()
	}
}

/*
Returns the size of our render target for this interface, this overrides the size of the [Viewport] marked as the xr viewport.
*/
func (Go) _get_render_target_size(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Returns the number of views this interface requires, 1 for mono, 2 for stereoscopic.
*/
func (Go) _get_view_count(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Returns the [Transform3D] that positions the [XRCamera3D] in the world.
*/
func (Go) _get_camera_transform(impl func(ptr unsafe.Pointer) gd.Transform3D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Returns a [Transform3D] for a given view.
*/
func (Go) _get_transform_for_view(impl func(ptr unsafe.Pointer, view int, cam_transform gd.Transform3D) gd.Transform3D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var view = gd.UnsafeGet[gd.Int](p_args,0)
		var cam_transform = gd.UnsafeGet[gd.Transform3D](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(view), cam_transform)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Returns the projection matrix for the given view as a [PackedFloat64Array].
*/
func (Go) _get_projection_for_view(impl func(ptr unsafe.Pointer, view int, aspect float64, z_near float64, z_far float64) []float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var view = gd.UnsafeGet[gd.Int](p_args,0)
		var aspect = gd.UnsafeGet[gd.Float](p_args,1)
		var z_near = gd.UnsafeGet[gd.Float](p_args,2)
		var z_far = gd.UnsafeGet[gd.Float](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(view), float64(aspect), float64(z_near), float64(z_far))
		gd.UnsafeSet(p_back, mmm.End(gc.PackedFloat64Slice(ret)))
		gc.End()
	}
}
func (Go) _get_vrs_texture(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Called if this [XRInterfaceExtension] is active before our physics and game process is called. Most XR interfaces will update its [XRPositionalTracker]s at this point in time.
*/
func (Go) _process(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called if this [XRInterfaceExtension] is active before rendering starts. Most XR interfaces will sync tracking at this point in time.
*/
func (Go) _pre_render(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called if this is our primary [XRInterfaceExtension] before we start processing a [Viewport] for every active XR [Viewport], returns [code]true[/code] if that viewport should be rendered. An XR interface may return [code]false[/code] if the user has taken off their headset and we can pause rendering.
*/
func (Go) _pre_draw_viewport(impl func(ptr unsafe.Pointer, render_target gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var render_target = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, render_target)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Called after the XR [Viewport] draw logic has completed.
*/
func (Go) _post_draw_viewport(impl func(ptr unsafe.Pointer, render_target gd.RID, screen_rect gd.Rect2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var render_target = gd.UnsafeGet[gd.RID](p_args,0)
		var screen_rect = gd.UnsafeGet[gd.Rect2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, render_target, screen_rect)
		gc.End()
	}
}

/*
Called if interface is active and queues have been submitted.
*/
func (Go) _end_frame(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Returns a [PackedStringArray] with tracker names configured by this interface. Note that user configuration can override this list.
*/
func (Go) _get_suggested_tracker_names(impl func(ptr unsafe.Pointer) []string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.PackedStringSlice(ret)))
		gc.End()
	}
}

/*
Returns a [PackedStringArray] with pose names configured by this interface. Note that user configuration can override this list.
*/
func (Go) _get_suggested_pose_names(impl func(ptr unsafe.Pointer, tracker_name string) []string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var tracker_name = mmm.Let[gd.StringName](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, tracker_name.String())
		gd.UnsafeSet(p_back, mmm.End(gc.PackedStringSlice(ret)))
		gc.End()
	}
}

/*
Returns a [enum XRInterface.TrackingStatus] specifying the current status of our tracking.
*/
func (Go) _get_tracking_status(impl func(ptr unsafe.Pointer) classdb.XRInterfaceTrackingStatus, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Triggers a haptic pulse to be emitted on the specified tracker.
*/
func (Go) _trigger_haptic_pulse(impl func(ptr unsafe.Pointer, action_name string, tracker_name string, frequency float64, amplitude float64, duration_sec float64, delay_sec float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var action_name = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var tracker_name = mmm.Let[gd.StringName](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		var frequency = gd.UnsafeGet[gd.Float](p_args,2)
		var amplitude = gd.UnsafeGet[gd.Float](p_args,3)
		var duration_sec = gd.UnsafeGet[gd.Float](p_args,4)
		var delay_sec = gd.UnsafeGet[gd.Float](p_args,5)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, action_name.String(), tracker_name.String(), float64(frequency), float64(amplitude), float64(duration_sec), float64(delay_sec))
		gc.End()
	}
}

/*
Return [code]true[/code] if anchor detection is enabled for this interface.
*/
func (Go) _get_anchor_detection_is_enabled(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Enables anchor detection on this interface if supported.
*/
func (Go) _set_anchor_detection_is_enabled(impl func(ptr unsafe.Pointer, enabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var enabled = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, enabled)
		gc.End()
	}
}

/*
Returns the camera feed ID for the [CameraFeed] registered with the [CameraServer] that should be presented as the background on an AR capable device (if applicable).
*/
func (Go) _get_camera_feed_id(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Return color texture into which to render (if applicable).
*/
func (Go) _get_color_texture(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Return depth texture into which to render (if applicable).
*/
func (Go) _get_depth_texture(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Return velocity texture into which to render (if applicable).
*/
func (Go) _get_velocity_texture(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (self Go) GetColorTexture() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetColorTexture())
}
func (self Go) GetDepthTexture() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetDepthTexture())
}
func (self Go) GetVelocityTexture() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetVelocityTexture())
}

/*
Blits our render results to screen optionally applying lens distortion. This can only be called while processing [code]_commit_views[/code].
*/
func (self Go) AddBlit(render_target gd.RID, src_rect gd.Rect2, dst_rect gd.Rect2i, use_layer bool, layer int, apply_lens_distortion bool, eye_center gd.Vector2, k1 float64, k2 float64, upscale float64, aspect_ratio float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddBlit(render_target, src_rect, dst_rect, use_layer, gd.Int(layer), apply_lens_distortion, eye_center, gd.Float(k1), gd.Float(k2), gd.Float(upscale), gd.Float(aspect_ratio))
}

/*
Returns a valid [RID] for a texture to which we should render the current frame if supported by the interface.
*/
func (self Go) GetRenderTargetTexture(render_target gd.RID) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetRenderTargetTexture(render_target))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.XRInterfaceExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("XRInterfaceExtension"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns the name of this interface.
*/
func (class) _get_name(impl func(ptr unsafe.Pointer) gd.StringName, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Returns the capabilities of this interface.
*/
func (class) _get_capabilities(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Returns [code]true[/code] if this interface has been initialized.
*/
func (class) _is_initialized(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Initializes the interface, returns [code]true[/code] on success.
*/
func (class) _initialize(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Uninitialize the interface.
*/
func (class) _uninitialize(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Returns a [Dictionary] with system information related to this interface.
*/
func (class) _get_system_info(impl func(ptr unsafe.Pointer) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Returns [code]true[/code] if this interface supports this play area mode.
*/
func (class) _supports_play_area_mode(impl func(ptr unsafe.Pointer, mode classdb.XRInterfacePlayAreaMode) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var mode = gd.UnsafeGet[classdb.XRInterfacePlayAreaMode](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Returns the play area mode that sets up our play area.
*/
func (class) _get_play_area_mode(impl func(ptr unsafe.Pointer) classdb.XRInterfacePlayAreaMode, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Set the play area mode for this interface.
*/
func (class) _set_play_area_mode(impl func(ptr unsafe.Pointer, mode classdb.XRInterfacePlayAreaMode) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var mode = gd.UnsafeGet[classdb.XRInterfacePlayAreaMode](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Returns a [PackedVector3Array] that represents the play areas boundaries (if applicable).
*/
func (class) _get_play_area(impl func(ptr unsafe.Pointer) gd.PackedVector3Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Returns the size of our render target for this interface, this overrides the size of the [Viewport] marked as the xr viewport.
*/
func (class) _get_render_target_size(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Returns the number of views this interface requires, 1 for mono, 2 for stereoscopic.
*/
func (class) _get_view_count(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Returns the [Transform3D] that positions the [XRCamera3D] in the world.
*/
func (class) _get_camera_transform(impl func(ptr unsafe.Pointer) gd.Transform3D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Returns a [Transform3D] for a given view.
*/
func (class) _get_transform_for_view(impl func(ptr unsafe.Pointer, view gd.Int, cam_transform gd.Transform3D) gd.Transform3D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var view = gd.UnsafeGet[gd.Int](p_args,0)
		var cam_transform = gd.UnsafeGet[gd.Transform3D](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, view, cam_transform)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Returns the projection matrix for the given view as a [PackedFloat64Array].
*/
func (class) _get_projection_for_view(impl func(ptr unsafe.Pointer, view gd.Int, aspect gd.Float, z_near gd.Float, z_far gd.Float) gd.PackedFloat64Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var view = gd.UnsafeGet[gd.Int](p_args,0)
		var aspect = gd.UnsafeGet[gd.Float](p_args,1)
		var z_near = gd.UnsafeGet[gd.Float](p_args,2)
		var z_far = gd.UnsafeGet[gd.Float](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, view, aspect, z_near, z_far)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _get_vrs_texture(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Called if this [XRInterfaceExtension] is active before our physics and game process is called. Most XR interfaces will update its [XRPositionalTracker]s at this point in time.
*/
func (class) _process(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called if this [XRInterfaceExtension] is active before rendering starts. Most XR interfaces will sync tracking at this point in time.
*/
func (class) _pre_render(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called if this is our primary [XRInterfaceExtension] before we start processing a [Viewport] for every active XR [Viewport], returns [code]true[/code] if that viewport should be rendered. An XR interface may return [code]false[/code] if the user has taken off their headset and we can pause rendering.
*/
func (class) _pre_draw_viewport(impl func(ptr unsafe.Pointer, render_target gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var render_target = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, render_target)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Called after the XR [Viewport] draw logic has completed.
*/
func (class) _post_draw_viewport(impl func(ptr unsafe.Pointer, render_target gd.RID, screen_rect gd.Rect2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var render_target = gd.UnsafeGet[gd.RID](p_args,0)
		var screen_rect = gd.UnsafeGet[gd.Rect2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, render_target, screen_rect)
		ctx.End()
	}
}

/*
Called if interface is active and queues have been submitted.
*/
func (class) _end_frame(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Returns a [PackedStringArray] with tracker names configured by this interface. Note that user configuration can override this list.
*/
func (class) _get_suggested_tracker_names(impl func(ptr unsafe.Pointer) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Returns a [PackedStringArray] with pose names configured by this interface. Note that user configuration can override this list.
*/
func (class) _get_suggested_pose_names(impl func(ptr unsafe.Pointer, tracker_name gd.StringName) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var tracker_name = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, tracker_name)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Returns a [enum XRInterface.TrackingStatus] specifying the current status of our tracking.
*/
func (class) _get_tracking_status(impl func(ptr unsafe.Pointer) classdb.XRInterfaceTrackingStatus, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Triggers a haptic pulse to be emitted on the specified tracker.
*/
func (class) _trigger_haptic_pulse(impl func(ptr unsafe.Pointer, action_name gd.String, tracker_name gd.StringName, frequency gd.Float, amplitude gd.Float, duration_sec gd.Float, delay_sec gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var action_name = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var tracker_name = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var frequency = gd.UnsafeGet[gd.Float](p_args,2)
		var amplitude = gd.UnsafeGet[gd.Float](p_args,3)
		var duration_sec = gd.UnsafeGet[gd.Float](p_args,4)
		var delay_sec = gd.UnsafeGet[gd.Float](p_args,5)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, action_name, tracker_name, frequency, amplitude, duration_sec, delay_sec)
		ctx.End()
	}
}

/*
Return [code]true[/code] if anchor detection is enabled for this interface.
*/
func (class) _get_anchor_detection_is_enabled(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Enables anchor detection on this interface if supported.
*/
func (class) _set_anchor_detection_is_enabled(impl func(ptr unsafe.Pointer, enabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var enabled = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, enabled)
		ctx.End()
	}
}

/*
Returns the camera feed ID for the [CameraFeed] registered with the [CameraServer] that should be presented as the background on an AR capable device (if applicable).
*/
func (class) _get_camera_feed_id(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Return color texture into which to render (if applicable).
*/
func (class) _get_color_texture(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Return depth texture into which to render (if applicable).
*/
func (class) _get_depth_texture(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Return velocity texture into which to render (if applicable).
*/
func (class) _get_velocity_texture(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

//go:nosplit
func (self class) GetColorTexture() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterfaceExtension.Bind_get_color_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetDepthTexture() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterfaceExtension.Bind_get_depth_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetVelocityTexture() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterfaceExtension.Bind_get_velocity_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Blits our render results to screen optionally applying lens distortion. This can only be called while processing [code]_commit_views[/code].
*/
//go:nosplit
func (self class) AddBlit(render_target gd.RID, src_rect gd.Rect2, dst_rect gd.Rect2i, use_layer bool, layer gd.Int, apply_lens_distortion bool, eye_center gd.Vector2, k1 gd.Float, k2 gd.Float, upscale gd.Float, aspect_ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, render_target)
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, dst_rect)
	callframe.Arg(frame, use_layer)
	callframe.Arg(frame, layer)
	callframe.Arg(frame, apply_lens_distortion)
	callframe.Arg(frame, eye_center)
	callframe.Arg(frame, k1)
	callframe.Arg(frame, k2)
	callframe.Arg(frame, upscale)
	callframe.Arg(frame, aspect_ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterfaceExtension.Bind_add_blit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a valid [RID] for a texture to which we should render the current frame if supported by the interface.
*/
//go:nosplit
func (self class) GetRenderTargetTexture(render_target gd.RID) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, render_target)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRInterfaceExtension.Bind_get_render_target_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsXRInterfaceExtension() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRInterfaceExtension() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsXRInterface() XRInterface.GD { return *((*XRInterface.GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRInterface() XRInterface.Go { return *((*XRInterface.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_name": return reflect.ValueOf(self._get_name);
	case "_get_capabilities": return reflect.ValueOf(self._get_capabilities);
	case "_is_initialized": return reflect.ValueOf(self._is_initialized);
	case "_initialize": return reflect.ValueOf(self._initialize);
	case "_uninitialize": return reflect.ValueOf(self._uninitialize);
	case "_get_system_info": return reflect.ValueOf(self._get_system_info);
	case "_supports_play_area_mode": return reflect.ValueOf(self._supports_play_area_mode);
	case "_get_play_area_mode": return reflect.ValueOf(self._get_play_area_mode);
	case "_set_play_area_mode": return reflect.ValueOf(self._set_play_area_mode);
	case "_get_play_area": return reflect.ValueOf(self._get_play_area);
	case "_get_render_target_size": return reflect.ValueOf(self._get_render_target_size);
	case "_get_view_count": return reflect.ValueOf(self._get_view_count);
	case "_get_camera_transform": return reflect.ValueOf(self._get_camera_transform);
	case "_get_transform_for_view": return reflect.ValueOf(self._get_transform_for_view);
	case "_get_projection_for_view": return reflect.ValueOf(self._get_projection_for_view);
	case "_get_vrs_texture": return reflect.ValueOf(self._get_vrs_texture);
	case "_process": return reflect.ValueOf(self._process);
	case "_pre_render": return reflect.ValueOf(self._pre_render);
	case "_pre_draw_viewport": return reflect.ValueOf(self._pre_draw_viewport);
	case "_post_draw_viewport": return reflect.ValueOf(self._post_draw_viewport);
	case "_end_frame": return reflect.ValueOf(self._end_frame);
	case "_get_suggested_tracker_names": return reflect.ValueOf(self._get_suggested_tracker_names);
	case "_get_suggested_pose_names": return reflect.ValueOf(self._get_suggested_pose_names);
	case "_get_tracking_status": return reflect.ValueOf(self._get_tracking_status);
	case "_trigger_haptic_pulse": return reflect.ValueOf(self._trigger_haptic_pulse);
	case "_get_anchor_detection_is_enabled": return reflect.ValueOf(self._get_anchor_detection_is_enabled);
	case "_set_anchor_detection_is_enabled": return reflect.ValueOf(self._set_anchor_detection_is_enabled);
	case "_get_camera_feed_id": return reflect.ValueOf(self._get_camera_feed_id);
	case "_get_color_texture": return reflect.ValueOf(self._get_color_texture);
	case "_get_depth_texture": return reflect.ValueOf(self._get_depth_texture);
	case "_get_velocity_texture": return reflect.ValueOf(self._get_velocity_texture);
	default: return gd.VirtualByName(self.AsXRInterface(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_get_name": return reflect.ValueOf(self._get_name);
	case "_get_capabilities": return reflect.ValueOf(self._get_capabilities);
	case "_is_initialized": return reflect.ValueOf(self._is_initialized);
	case "_initialize": return reflect.ValueOf(self._initialize);
	case "_uninitialize": return reflect.ValueOf(self._uninitialize);
	case "_get_system_info": return reflect.ValueOf(self._get_system_info);
	case "_supports_play_area_mode": return reflect.ValueOf(self._supports_play_area_mode);
	case "_get_play_area_mode": return reflect.ValueOf(self._get_play_area_mode);
	case "_set_play_area_mode": return reflect.ValueOf(self._set_play_area_mode);
	case "_get_play_area": return reflect.ValueOf(self._get_play_area);
	case "_get_render_target_size": return reflect.ValueOf(self._get_render_target_size);
	case "_get_view_count": return reflect.ValueOf(self._get_view_count);
	case "_get_camera_transform": return reflect.ValueOf(self._get_camera_transform);
	case "_get_transform_for_view": return reflect.ValueOf(self._get_transform_for_view);
	case "_get_projection_for_view": return reflect.ValueOf(self._get_projection_for_view);
	case "_get_vrs_texture": return reflect.ValueOf(self._get_vrs_texture);
	case "_process": return reflect.ValueOf(self._process);
	case "_pre_render": return reflect.ValueOf(self._pre_render);
	case "_pre_draw_viewport": return reflect.ValueOf(self._pre_draw_viewport);
	case "_post_draw_viewport": return reflect.ValueOf(self._post_draw_viewport);
	case "_end_frame": return reflect.ValueOf(self._end_frame);
	case "_get_suggested_tracker_names": return reflect.ValueOf(self._get_suggested_tracker_names);
	case "_get_suggested_pose_names": return reflect.ValueOf(self._get_suggested_pose_names);
	case "_get_tracking_status": return reflect.ValueOf(self._get_tracking_status);
	case "_trigger_haptic_pulse": return reflect.ValueOf(self._trigger_haptic_pulse);
	case "_get_anchor_detection_is_enabled": return reflect.ValueOf(self._get_anchor_detection_is_enabled);
	case "_set_anchor_detection_is_enabled": return reflect.ValueOf(self._set_anchor_detection_is_enabled);
	case "_get_camera_feed_id": return reflect.ValueOf(self._get_camera_feed_id);
	case "_get_color_texture": return reflect.ValueOf(self._get_color_texture);
	case "_get_depth_texture": return reflect.ValueOf(self._get_depth_texture);
	case "_get_velocity_texture": return reflect.ValueOf(self._get_velocity_texture);
	default: return gd.VirtualByName(self.AsXRInterface(), name)
	}
}
func init() {classdb.Register("XRInterfaceExtension", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
