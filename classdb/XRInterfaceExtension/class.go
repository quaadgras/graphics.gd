package XRInterfaceExtension

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/XRInterface"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Float"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Rect2"
import "graphics.gd/variant/Rect2i"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

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
		GetSystemInfo() Dictionary.Any
		//Returns [code]true[/code] if this interface supports this play area mode.
		SupportsPlayAreaMode(mode gdclass.XRInterfacePlayAreaMode) bool
		//Returns the play area mode that sets up our play area.
		GetPlayAreaMode() gdclass.XRInterfacePlayAreaMode
		//Set the play area mode for this interface.
		SetPlayAreaMode(mode gdclass.XRInterfacePlayAreaMode) bool
		//Returns a [PackedVector3Array] that represents the play areas boundaries (if applicable).
		GetPlayArea() []Vector3.XYZ
		//Returns the size of our render target for this interface, this overrides the size of the [Viewport] marked as the xr viewport.
		GetRenderTargetSize() Vector2.XY
		//Returns the number of views this interface requires, 1 for mono, 2 for stereoscopic.
		GetViewCount() int
		//Returns the [Transform3D] that positions the [XRCamera3D] in the world.
		GetCameraTransform() Transform3D.BasisOrigin
		//Returns a [Transform3D] for a given view.
		GetTransformForView(view int, cam_transform Transform3D.BasisOrigin) Transform3D.BasisOrigin
		//Returns the projection matrix for the given view as a [PackedFloat64Array].
		GetProjectionForView(view int, aspect Float.X, z_near Float.X, z_far Float.X) []float64
		GetVrsTexture() Resource.ID
		//Called if this [XRInterfaceExtension] is active before our physics and game process is called. Most XR interfaces will update its [XRPositionalTracker]s at this point in time.
		Process()
		//Called if this [XRInterfaceExtension] is active before rendering starts. Most XR interfaces will sync tracking at this point in time.
		PreRender()
		//Called if this is our primary [XRInterfaceExtension] before we start processing a [Viewport] for every active XR [Viewport], returns [code]true[/code] if that viewport should be rendered. An XR interface may return [code]false[/code] if the user has taken off their headset and we can pause rendering.
		PreDrawViewport(render_target Resource.ID) bool
		//Called after the XR [Viewport] draw logic has completed.
		PostDrawViewport(render_target Resource.ID, screen_rect Rect2.PositionSize)
		//Called if interface is active and queues have been submitted.
		EndFrame()
		//Returns a [PackedStringArray] with tracker names configured by this interface. Note that user configuration can override this list.
		GetSuggestedTrackerNames() []string
		//Returns a [PackedStringArray] with pose names configured by this interface. Note that user configuration can override this list.
		GetSuggestedPoseNames(tracker_name string) []string
		//Returns a [enum XRInterface.TrackingStatus] specifying the current status of our tracking.
		GetTrackingStatus() gdclass.XRInterfaceTrackingStatus
		//Triggers a haptic pulse to be emitted on the specified tracker.
		TriggerHapticPulse(action_name string, tracker_name string, frequency Float.X, amplitude Float.X, duration_sec Float.X, delay_sec Float.X)
		//Return [code]true[/code] if anchor detection is enabled for this interface.
		GetAnchorDetectionIsEnabled() bool
		//Enables anchor detection on this interface if supported.
		SetAnchorDetectionIsEnabled(enabled bool)
		//Returns the camera feed ID for the [CameraFeed] registered with the [CameraServer] that should be presented as the background on an AR capable device (if applicable).
		GetCameraFeedId() int
		//Return color texture into which to render (if applicable).
		GetColorTexture() Resource.ID
		//Return depth texture into which to render (if applicable).
		GetDepthTexture() Resource.ID
		//Return velocity texture into which to render (if applicable).
		GetVelocityTexture() Resource.ID
	}
*/
type Instance [1]gdclass.XRInterfaceExtension
type Any interface {
	gd.IsClass
	AsXRInterfaceExtension() Instance
}

/*
Returns the name of this interface.
*/
func (Instance) _get_name(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewStringName(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns the capabilities of this interface.
*/
func (Instance) _get_capabilities(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Returns [code]true[/code] if this interface has been initialized.
*/
func (Instance) _is_initialized(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Initializes the interface, returns [code]true[/code] on success.
*/
func (Instance) _initialize(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Uninitialize the interface.
*/
func (Instance) _uninitialize(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Returns a [Dictionary] with system information related to this interface.
*/
func (Instance) _get_system_info(impl func(ptr unsafe.Pointer) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns [code]true[/code] if this interface supports this play area mode.
*/
func (Instance) _supports_play_area_mode(impl func(ptr unsafe.Pointer, mode gdclass.XRInterfacePlayAreaMode) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var mode = gd.UnsafeGet[gdclass.XRInterfacePlayAreaMode](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the play area mode that sets up our play area.
*/
func (Instance) _get_play_area_mode(impl func(ptr unsafe.Pointer) gdclass.XRInterfacePlayAreaMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Set the play area mode for this interface.
*/
func (Instance) _set_play_area_mode(impl func(ptr unsafe.Pointer, mode gdclass.XRInterfacePlayAreaMode) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var mode = gd.UnsafeGet[gdclass.XRInterfacePlayAreaMode](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns a [PackedVector3Array] that represents the play areas boundaries (if applicable).
*/
func (Instance) _get_play_area(impl func(ptr unsafe.Pointer) []Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedVector3Slice(*(*[]gd.Vector3)(unsafe.Pointer(&ret))))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns the size of our render target for this interface, this overrides the size of the [Viewport] marked as the xr viewport.
*/
func (Instance) _get_render_target_size(impl func(ptr unsafe.Pointer) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
Returns the number of views this interface requires, 1 for mono, 2 for stereoscopic.
*/
func (Instance) _get_view_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Returns the [Transform3D] that positions the [XRCamera3D] in the world.
*/
func (Instance) _get_camera_transform(impl func(ptr unsafe.Pointer) Transform3D.BasisOrigin) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Transform3D(ret))
	}
}

/*
Returns a [Transform3D] for a given view.
*/
func (Instance) _get_transform_for_view(impl func(ptr unsafe.Pointer, view int, cam_transform Transform3D.BasisOrigin) Transform3D.BasisOrigin) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var view = gd.UnsafeGet[gd.Int](p_args, 0)
		var cam_transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(view), cam_transform)
		gd.UnsafeSet(p_back, gd.Transform3D(ret))
	}
}

/*
Returns the projection matrix for the given view as a [PackedFloat64Array].
*/
func (Instance) _get_projection_for_view(impl func(ptr unsafe.Pointer, view int, aspect Float.X, z_near Float.X, z_far Float.X) []float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var view = gd.UnsafeGet[gd.Int](p_args, 0)
		var aspect = gd.UnsafeGet[gd.Float](p_args, 1)
		var z_near = gd.UnsafeGet[gd.Float](p_args, 2)
		var z_far = gd.UnsafeGet[gd.Float](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(view), Float.X(aspect), Float.X(z_near), Float.X(z_far))
		ptr, ok := pointers.End(gd.NewPackedFloat64Slice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_vrs_texture(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called if this [XRInterfaceExtension] is active before our physics and game process is called. Most XR interfaces will update its [XRPositionalTracker]s at this point in time.
*/
func (Instance) _process(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called if this [XRInterfaceExtension] is active before rendering starts. Most XR interfaces will sync tracking at this point in time.
*/
func (Instance) _pre_render(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called if this is our primary [XRInterfaceExtension] before we start processing a [Viewport] for every active XR [Viewport], returns [code]true[/code] if that viewport should be rendered. An XR interface may return [code]false[/code] if the user has taken off their headset and we can pause rendering.
*/
func (Instance) _pre_draw_viewport(impl func(ptr unsafe.Pointer, render_target Resource.ID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var render_target = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, render_target)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called after the XR [Viewport] draw logic has completed.
*/
func (Instance) _post_draw_viewport(impl func(ptr unsafe.Pointer, render_target Resource.ID, screen_rect Rect2.PositionSize)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var render_target = gd.UnsafeGet[gd.RID](p_args, 0)
		var screen_rect = gd.UnsafeGet[gd.Rect2](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, render_target, screen_rect)
	}
}

/*
Called if interface is active and queues have been submitted.
*/
func (Instance) _end_frame(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Returns a [PackedStringArray] with tracker names configured by this interface. Note that user configuration can override this list.
*/
func (Instance) _get_suggested_tracker_names(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns a [PackedStringArray] with pose names configured by this interface. Note that user configuration can override this list.
*/
func (Instance) _get_suggested_pose_names(impl func(ptr unsafe.Pointer, tracker_name string) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var tracker_name = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(tracker_name)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, tracker_name.String())
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns a [enum XRInterface.TrackingStatus] specifying the current status of our tracking.
*/
func (Instance) _get_tracking_status(impl func(ptr unsafe.Pointer) gdclass.XRInterfaceTrackingStatus) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Triggers a haptic pulse to be emitted on the specified tracker.
*/
func (Instance) _trigger_haptic_pulse(impl func(ptr unsafe.Pointer, action_name string, tracker_name string, frequency Float.X, amplitude Float.X, duration_sec Float.X, delay_sec Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var action_name = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(action_name)
		var tracker_name = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(tracker_name)
		var frequency = gd.UnsafeGet[gd.Float](p_args, 2)
		var amplitude = gd.UnsafeGet[gd.Float](p_args, 3)
		var duration_sec = gd.UnsafeGet[gd.Float](p_args, 4)
		var delay_sec = gd.UnsafeGet[gd.Float](p_args, 5)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, action_name.String(), tracker_name.String(), Float.X(frequency), Float.X(amplitude), Float.X(duration_sec), Float.X(delay_sec))
	}
}

/*
Return [code]true[/code] if anchor detection is enabled for this interface.
*/
func (Instance) _get_anchor_detection_is_enabled(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Enables anchor detection on this interface if supported.
*/
func (Instance) _set_anchor_detection_is_enabled(impl func(ptr unsafe.Pointer, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var enabled = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, enabled)
	}
}

/*
Returns the camera feed ID for the [CameraFeed] registered with the [CameraServer] that should be presented as the background on an AR capable device (if applicable).
*/
func (Instance) _get_camera_feed_id(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Return color texture into which to render (if applicable).
*/
func (Instance) _get_color_texture(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Return depth texture into which to render (if applicable).
*/
func (Instance) _get_depth_texture(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Return velocity texture into which to render (if applicable).
*/
func (Instance) _get_velocity_texture(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (self Instance) GetColorTexture() Resource.ID {
	return Resource.ID(class(self).GetColorTexture())
}
func (self Instance) GetDepthTexture() Resource.ID {
	return Resource.ID(class(self).GetDepthTexture())
}
func (self Instance) GetVelocityTexture() Resource.ID {
	return Resource.ID(class(self).GetVelocityTexture())
}

/*
Blits our render results to screen optionally applying lens distortion. This can only be called while processing [code]_commit_views[/code].
*/
func (self Instance) AddBlit(render_target Resource.ID, src_rect Rect2.PositionSize, dst_rect Rect2i.PositionSize, use_layer bool, layer int, apply_lens_distortion bool, eye_center Vector2.XY, k1 Float.X, k2 Float.X, upscale Float.X, aspect_ratio Float.X) {
	class(self).AddBlit(render_target, gd.Rect2(src_rect), gd.Rect2i(dst_rect), use_layer, gd.Int(layer), apply_lens_distortion, gd.Vector2(eye_center), gd.Float(k1), gd.Float(k2), gd.Float(upscale), gd.Float(aspect_ratio))
}

/*
Returns a valid [RID] for a texture to which we should render the current frame if supported by the interface.
*/
func (self Instance) GetRenderTargetTexture(render_target Resource.ID) Resource.ID {
	return Resource.ID(class(self).GetRenderTargetTexture(render_target))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.XRInterfaceExtension

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRInterfaceExtension"))
	return Instance{*(*gdclass.XRInterfaceExtension)(unsafe.Pointer(&object))}
}

/*
Returns the name of this interface.
*/
func (class) _get_name(impl func(ptr unsafe.Pointer) gd.StringName) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns the capabilities of this interface.
*/
func (class) _get_capabilities(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns [code]true[/code] if this interface has been initialized.
*/
func (class) _is_initialized(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Initializes the interface, returns [code]true[/code] on success.
*/
func (class) _initialize(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Uninitialize the interface.
*/
func (class) _uninitialize(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Returns a [Dictionary] with system information related to this interface.
*/
func (class) _get_system_info(impl func(ptr unsafe.Pointer) gd.Dictionary) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns [code]true[/code] if this interface supports this play area mode.
*/
func (class) _supports_play_area_mode(impl func(ptr unsafe.Pointer, mode gdclass.XRInterfacePlayAreaMode) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var mode = gd.UnsafeGet[gdclass.XRInterfacePlayAreaMode](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the play area mode that sets up our play area.
*/
func (class) _get_play_area_mode(impl func(ptr unsafe.Pointer) gdclass.XRInterfacePlayAreaMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Set the play area mode for this interface.
*/
func (class) _set_play_area_mode(impl func(ptr unsafe.Pointer, mode gdclass.XRInterfacePlayAreaMode) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var mode = gd.UnsafeGet[gdclass.XRInterfacePlayAreaMode](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns a [PackedVector3Array] that represents the play areas boundaries (if applicable).
*/
func (class) _get_play_area(impl func(ptr unsafe.Pointer) gd.PackedVector3Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns the size of our render target for this interface, this overrides the size of the [Viewport] marked as the xr viewport.
*/
func (class) _get_render_target_size(impl func(ptr unsafe.Pointer) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the number of views this interface requires, 1 for mono, 2 for stereoscopic.
*/
func (class) _get_view_count(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the [Transform3D] that positions the [XRCamera3D] in the world.
*/
func (class) _get_camera_transform(impl func(ptr unsafe.Pointer) gd.Transform3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns a [Transform3D] for a given view.
*/
func (class) _get_transform_for_view(impl func(ptr unsafe.Pointer, view gd.Int, cam_transform gd.Transform3D) gd.Transform3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var view = gd.UnsafeGet[gd.Int](p_args, 0)
		var cam_transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, view, cam_transform)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the projection matrix for the given view as a [PackedFloat64Array].
*/
func (class) _get_projection_for_view(impl func(ptr unsafe.Pointer, view gd.Int, aspect gd.Float, z_near gd.Float, z_far gd.Float) gd.PackedFloat64Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var view = gd.UnsafeGet[gd.Int](p_args, 0)
		var aspect = gd.UnsafeGet[gd.Float](p_args, 1)
		var z_near = gd.UnsafeGet[gd.Float](p_args, 2)
		var z_far = gd.UnsafeGet[gd.Float](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, view, aspect, z_near, z_far)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_vrs_texture(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called if this [XRInterfaceExtension] is active before our physics and game process is called. Most XR interfaces will update its [XRPositionalTracker]s at this point in time.
*/
func (class) _process(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called if this [XRInterfaceExtension] is active before rendering starts. Most XR interfaces will sync tracking at this point in time.
*/
func (class) _pre_render(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called if this is our primary [XRInterfaceExtension] before we start processing a [Viewport] for every active XR [Viewport], returns [code]true[/code] if that viewport should be rendered. An XR interface may return [code]false[/code] if the user has taken off their headset and we can pause rendering.
*/
func (class) _pre_draw_viewport(impl func(ptr unsafe.Pointer, render_target gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var render_target = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, render_target)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called after the XR [Viewport] draw logic has completed.
*/
func (class) _post_draw_viewport(impl func(ptr unsafe.Pointer, render_target gd.RID, screen_rect gd.Rect2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var render_target = gd.UnsafeGet[gd.RID](p_args, 0)
		var screen_rect = gd.UnsafeGet[gd.Rect2](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, render_target, screen_rect)
	}
}

/*
Called if interface is active and queues have been submitted.
*/
func (class) _end_frame(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Returns a [PackedStringArray] with tracker names configured by this interface. Note that user configuration can override this list.
*/
func (class) _get_suggested_tracker_names(impl func(ptr unsafe.Pointer) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns a [PackedStringArray] with pose names configured by this interface. Note that user configuration can override this list.
*/
func (class) _get_suggested_pose_names(impl func(ptr unsafe.Pointer, tracker_name gd.StringName) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var tracker_name = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, tracker_name)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns a [enum XRInterface.TrackingStatus] specifying the current status of our tracking.
*/
func (class) _get_tracking_status(impl func(ptr unsafe.Pointer) gdclass.XRInterfaceTrackingStatus) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Triggers a haptic pulse to be emitted on the specified tracker.
*/
func (class) _trigger_haptic_pulse(impl func(ptr unsafe.Pointer, action_name gd.String, tracker_name gd.StringName, frequency gd.Float, amplitude gd.Float, duration_sec gd.Float, delay_sec gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var action_name = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var tracker_name = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 1))
		var frequency = gd.UnsafeGet[gd.Float](p_args, 2)
		var amplitude = gd.UnsafeGet[gd.Float](p_args, 3)
		var duration_sec = gd.UnsafeGet[gd.Float](p_args, 4)
		var delay_sec = gd.UnsafeGet[gd.Float](p_args, 5)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, action_name, tracker_name, frequency, amplitude, duration_sec, delay_sec)
	}
}

/*
Return [code]true[/code] if anchor detection is enabled for this interface.
*/
func (class) _get_anchor_detection_is_enabled(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Enables anchor detection on this interface if supported.
*/
func (class) _set_anchor_detection_is_enabled(impl func(ptr unsafe.Pointer, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var enabled = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, enabled)
	}
}

/*
Returns the camera feed ID for the [CameraFeed] registered with the [CameraServer] that should be presented as the background on an AR capable device (if applicable).
*/
func (class) _get_camera_feed_id(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Return color texture into which to render (if applicable).
*/
func (class) _get_color_texture(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Return depth texture into which to render (if applicable).
*/
func (class) _get_depth_texture(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Return velocity texture into which to render (if applicable).
*/
func (class) _get_velocity_texture(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

//go:nosplit
func (self class) GetColorTexture() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterfaceExtension.Bind_get_color_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetDepthTexture() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterfaceExtension.Bind_get_depth_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetVelocityTexture() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterfaceExtension.Bind_get_velocity_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Blits our render results to screen optionally applying lens distortion. This can only be called while processing [code]_commit_views[/code].
*/
//go:nosplit
func (self class) AddBlit(render_target gd.RID, src_rect gd.Rect2, dst_rect gd.Rect2i, use_layer bool, layer gd.Int, apply_lens_distortion bool, eye_center gd.Vector2, k1 gd.Float, k2 gd.Float, upscale gd.Float, aspect_ratio gd.Float) {
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterfaceExtension.Bind_add_blit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a valid [RID] for a texture to which we should render the current frame if supported by the interface.
*/
//go:nosplit
func (self class) GetRenderTargetTexture(render_target gd.RID) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, render_target)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRInterfaceExtension.Bind_get_render_target_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsXRInterfaceExtension() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsXRInterfaceExtension() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsXRInterface() XRInterface.Advanced {
	return *((*XRInterface.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsXRInterface() XRInterface.Instance {
	return *((*XRInterface.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_name":
		return reflect.ValueOf(self._get_name)
	case "_get_capabilities":
		return reflect.ValueOf(self._get_capabilities)
	case "_is_initialized":
		return reflect.ValueOf(self._is_initialized)
	case "_initialize":
		return reflect.ValueOf(self._initialize)
	case "_uninitialize":
		return reflect.ValueOf(self._uninitialize)
	case "_get_system_info":
		return reflect.ValueOf(self._get_system_info)
	case "_supports_play_area_mode":
		return reflect.ValueOf(self._supports_play_area_mode)
	case "_get_play_area_mode":
		return reflect.ValueOf(self._get_play_area_mode)
	case "_set_play_area_mode":
		return reflect.ValueOf(self._set_play_area_mode)
	case "_get_play_area":
		return reflect.ValueOf(self._get_play_area)
	case "_get_render_target_size":
		return reflect.ValueOf(self._get_render_target_size)
	case "_get_view_count":
		return reflect.ValueOf(self._get_view_count)
	case "_get_camera_transform":
		return reflect.ValueOf(self._get_camera_transform)
	case "_get_transform_for_view":
		return reflect.ValueOf(self._get_transform_for_view)
	case "_get_projection_for_view":
		return reflect.ValueOf(self._get_projection_for_view)
	case "_get_vrs_texture":
		return reflect.ValueOf(self._get_vrs_texture)
	case "_process":
		return reflect.ValueOf(self._process)
	case "_pre_render":
		return reflect.ValueOf(self._pre_render)
	case "_pre_draw_viewport":
		return reflect.ValueOf(self._pre_draw_viewport)
	case "_post_draw_viewport":
		return reflect.ValueOf(self._post_draw_viewport)
	case "_end_frame":
		return reflect.ValueOf(self._end_frame)
	case "_get_suggested_tracker_names":
		return reflect.ValueOf(self._get_suggested_tracker_names)
	case "_get_suggested_pose_names":
		return reflect.ValueOf(self._get_suggested_pose_names)
	case "_get_tracking_status":
		return reflect.ValueOf(self._get_tracking_status)
	case "_trigger_haptic_pulse":
		return reflect.ValueOf(self._trigger_haptic_pulse)
	case "_get_anchor_detection_is_enabled":
		return reflect.ValueOf(self._get_anchor_detection_is_enabled)
	case "_set_anchor_detection_is_enabled":
		return reflect.ValueOf(self._set_anchor_detection_is_enabled)
	case "_get_camera_feed_id":
		return reflect.ValueOf(self._get_camera_feed_id)
	case "_get_color_texture":
		return reflect.ValueOf(self._get_color_texture)
	case "_get_depth_texture":
		return reflect.ValueOf(self._get_depth_texture)
	case "_get_velocity_texture":
		return reflect.ValueOf(self._get_velocity_texture)
	default:
		return gd.VirtualByName(self.AsXRInterface(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_name":
		return reflect.ValueOf(self._get_name)
	case "_get_capabilities":
		return reflect.ValueOf(self._get_capabilities)
	case "_is_initialized":
		return reflect.ValueOf(self._is_initialized)
	case "_initialize":
		return reflect.ValueOf(self._initialize)
	case "_uninitialize":
		return reflect.ValueOf(self._uninitialize)
	case "_get_system_info":
		return reflect.ValueOf(self._get_system_info)
	case "_supports_play_area_mode":
		return reflect.ValueOf(self._supports_play_area_mode)
	case "_get_play_area_mode":
		return reflect.ValueOf(self._get_play_area_mode)
	case "_set_play_area_mode":
		return reflect.ValueOf(self._set_play_area_mode)
	case "_get_play_area":
		return reflect.ValueOf(self._get_play_area)
	case "_get_render_target_size":
		return reflect.ValueOf(self._get_render_target_size)
	case "_get_view_count":
		return reflect.ValueOf(self._get_view_count)
	case "_get_camera_transform":
		return reflect.ValueOf(self._get_camera_transform)
	case "_get_transform_for_view":
		return reflect.ValueOf(self._get_transform_for_view)
	case "_get_projection_for_view":
		return reflect.ValueOf(self._get_projection_for_view)
	case "_get_vrs_texture":
		return reflect.ValueOf(self._get_vrs_texture)
	case "_process":
		return reflect.ValueOf(self._process)
	case "_pre_render":
		return reflect.ValueOf(self._pre_render)
	case "_pre_draw_viewport":
		return reflect.ValueOf(self._pre_draw_viewport)
	case "_post_draw_viewport":
		return reflect.ValueOf(self._post_draw_viewport)
	case "_end_frame":
		return reflect.ValueOf(self._end_frame)
	case "_get_suggested_tracker_names":
		return reflect.ValueOf(self._get_suggested_tracker_names)
	case "_get_suggested_pose_names":
		return reflect.ValueOf(self._get_suggested_pose_names)
	case "_get_tracking_status":
		return reflect.ValueOf(self._get_tracking_status)
	case "_trigger_haptic_pulse":
		return reflect.ValueOf(self._trigger_haptic_pulse)
	case "_get_anchor_detection_is_enabled":
		return reflect.ValueOf(self._get_anchor_detection_is_enabled)
	case "_set_anchor_detection_is_enabled":
		return reflect.ValueOf(self._set_anchor_detection_is_enabled)
	case "_get_camera_feed_id":
		return reflect.ValueOf(self._get_camera_feed_id)
	case "_get_color_texture":
		return reflect.ValueOf(self._get_color_texture)
	case "_get_depth_texture":
		return reflect.ValueOf(self._get_depth_texture)
	case "_get_velocity_texture":
		return reflect.ValueOf(self._get_velocity_texture)
	default:
		return gd.VirtualByName(self.AsXRInterface(), name)
	}
}
func init() {
	gdclass.Register("XRInterfaceExtension", func(ptr gd.Object) any {
		return [1]gdclass.XRInterfaceExtension{*(*gdclass.XRInterfaceExtension)(unsafe.Pointer(&ptr))}
	})
}
