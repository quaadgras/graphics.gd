package WebXRInterface

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/XRInterface"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
WebXR is an open standard that allows creating VR and AR applications that run in the web browser.
As such, this interface is only available when running in Web exports.
WebXR supports a wide range of devices, from the very capable (like Valve Index, HTC Vive, Oculus Rift and Quest) down to the much less capable (like Google Cardboard, Oculus Go, GearVR, or plain smartphones).
Since WebXR is based on JavaScript, it makes extensive use of callbacks, which means that [WebXRInterface] is forced to use signals, where other XR interfaces would instead use functions that return a result immediately. This makes [WebXRInterface] quite a bit more complicated to initialize than other XR interfaces.
Here's the minimum code required to start an immersive VR session:
[codeblock]
extends Node3D

var webxr_interface
var vr_supported = false

func _ready():
    # We assume this node has a button as a child.
    # This button is for the user to consent to entering immersive VR mode.
    $Button.pressed.connect(self._on_button_pressed)

    webxr_interface = XRServer.find_interface("WebXR")
    if webxr_interface:
        # WebXR uses a lot of asynchronous callbacks, so we connect to various
        # signals in order to receive them.
        webxr_interface.session_supported.connect(self._webxr_session_supported)
        webxr_interface.session_started.connect(self._webxr_session_started)
        webxr_interface.session_ended.connect(self._webxr_session_ended)
        webxr_interface.session_failed.connect(self._webxr_session_failed)

        # This returns immediately - our _webxr_session_supported() method
        # (which we connected to the "session_supported" signal above) will
        # be called sometime later to let us know if it's supported or not.
        webxr_interface.is_session_supported("immersive-vr")

func _webxr_session_supported(session_mode, supported):
    if session_mode == 'immersive-vr':
        vr_supported = supported

func _on_button_pressed():
    if not vr_supported:
        OS.alert("Your browser doesn't support VR")
        return

    # We want an immersive VR session, as opposed to AR ('immersive-ar') or a
    # simple 3DoF viewer ('viewer').
    webxr_interface.session_mode = 'immersive-vr'
    # 'bounded-floor' is room scale, 'local-floor' is a standing or sitting
    # experience (it puts you 1.6m above the ground if you have 3DoF headset),
    # whereas as 'local' puts you down at the XROrigin.
    # This list means it'll first try to request 'bounded-floor', then
    # fallback on 'local-floor' and ultimately 'local', if nothing else is
    # supported.
    webxr_interface.requested_reference_space_types = 'bounded-floor, local-floor, local'
    # In order to use 'local-floor' or 'bounded-floor' we must also
    # mark the features as required or optional. By including 'hand-tracking'
    # as an optional feature, it will be enabled if supported.
    webxr_interface.required_features = 'local-floor'
    webxr_interface.optional_features = 'bounded-floor, hand-tracking'

    # This will return false if we're unable to even request the session,
    # however, it can still fail asynchronously later in the process, so we
    # only know if it's really succeeded or failed when our
    # _webxr_session_started() or _webxr_session_failed() methods are called.
    if not webxr_interface.initialize():
        OS.alert("Failed to initialize")
        return

func _webxr_session_started():
    $Button.visible = false
    # This tells Godot to start rendering to the headset.
    get_viewport().use_xr = true
    # This will be the reference space type you ultimately got, out of the
    # types that you requested above. This is useful if you want the game to
    # work a little differently in 'bounded-floor' versus 'local-floor'.
    print("Reference space type: ", webxr_interface.reference_space_type)
    # This will be the list of features that were successfully enabled
    # (except on browsers that don't support this property).
    print("Enabled features: ", webxr_interface.enabled_features)

func _webxr_session_ended():
    $Button.visible = true
    # If the user exits immersive mode, then we tell Godot to render to the web
    # page again.
    get_viewport().use_xr = false

func _webxr_session_failed(message):
    OS.alert("Failed to initialize: " + message)
[/codeblock]
There are a couple ways to handle "controller" input:
- Using [XRController3D] nodes and their [signal XRController3D.button_pressed] and [signal XRController3D.button_released] signals. This is how controllers are typically handled in XR apps in Godot, however, this will only work with advanced VR controllers like the Oculus Touch or Index controllers, for example.
- Using the [signal select], [signal squeeze] and related signals. This method will work for both advanced VR controllers, and non-traditional input sources like a tap on the screen, a spoken voice command or a button press on the device itself.
You can use both methods to allow your game or app to support a wider or narrower set of devices and input methods, or to allow more advanced interactions with more advanced devices.

*/
type Go [1]classdb.WebXRInterface

/*
Checks if the given [param session_mode] is supported by the user's browser.
Possible values come from [url=https://developer.mozilla.org/en-US/docs/Web/API/XRSessionMode]WebXR's XRSessionMode[/url], including: [code]"immersive-vr"[/code], [code]"immersive-ar"[/code], and [code]"inline"[/code].
This method returns nothing, instead it emits the [signal session_supported] signal with the result.
*/
func (self Go) IsSessionSupported(session_mode string) {
	class(self).IsSessionSupported(gd.NewString(session_mode))
}

/*
Returns [code]true[/code] if there is an active input source with the given [param input_source_id].
*/
func (self Go) IsInputSourceActive(input_source_id int) bool {
	return bool(class(self).IsInputSourceActive(gd.Int(input_source_id)))
}

/*
Gets an [XRControllerTracker] for the given [param input_source_id].
In the context of WebXR, an input source can be an advanced VR controller like the Oculus Touch or Index controllers, or even a tap on the screen, a spoken voice command or a button press on the device itself. When a non-traditional input source is used, interpret the position and orientation of the [XRPositionalTracker] as a ray pointing at the object the user wishes to interact with.
Use this method to get information about the input source that triggered one of these signals:
- [signal selectstart]
- [signal select]
- [signal selectend]
- [signal squeezestart]
- [signal squeeze]
- [signal squeezestart]
*/
func (self Go) GetInputSourceTracker(input_source_id int) gdclass.XRControllerTracker {
	return gdclass.XRControllerTracker(class(self).GetInputSourceTracker(gd.Int(input_source_id)))
}

/*
Returns the target ray mode for the given [param input_source_id].
This can help interpret the input coming from that input source. See [url=https://developer.mozilla.org/en-US/docs/Web/API/XRInputSource/targetRayMode]XRInputSource.targetRayMode[/url] for more information.
*/
func (self Go) GetInputSourceTargetRayMode(input_source_id int) classdb.WebXRInterfaceTargetRayMode {
	return classdb.WebXRInterfaceTargetRayMode(class(self).GetInputSourceTargetRayMode(gd.Int(input_source_id)))
}

/*
Returns the display refresh rate for the current HMD. Not supported on all HMDs and browsers. It may not report an accurate value until after using [method set_display_refresh_rate].
*/
func (self Go) GetDisplayRefreshRate() float64 {
	return float64(float64(class(self).GetDisplayRefreshRate()))
}

/*
Sets the display refresh rate for the current HMD. Not supported on all HMDs and browsers. It won't take effect right away until after [signal display_refresh_rate_changed] is emitted.
*/
func (self Go) SetDisplayRefreshRate(refresh_rate float64) {
	class(self).SetDisplayRefreshRate(gd.Float(refresh_rate))
}

/*
Returns display refresh rates supported by the current HMD. Only returned if this feature is supported by the web browser and after the interface has been initialized.
*/
func (self Go) GetAvailableDisplayRefreshRates() gd.Array {
	return gd.Array(class(self).GetAvailableDisplayRefreshRates())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.WebXRInterface
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WebXRInterface"))
	return Go{classdb.WebXRInterface(object)}
}

func (self Go) SessionMode() string {
		return string(class(self).GetSessionMode().String())
}

func (self Go) SetSessionMode(value string) {
	class(self).SetSessionMode(gd.NewString(value))
}

func (self Go) RequiredFeatures() string {
		return string(class(self).GetRequiredFeatures().String())
}

func (self Go) SetRequiredFeatures(value string) {
	class(self).SetRequiredFeatures(gd.NewString(value))
}

func (self Go) OptionalFeatures() string {
		return string(class(self).GetOptionalFeatures().String())
}

func (self Go) SetOptionalFeatures(value string) {
	class(self).SetOptionalFeatures(gd.NewString(value))
}

func (self Go) RequestedReferenceSpaceTypes() string {
		return string(class(self).GetRequestedReferenceSpaceTypes().String())
}

func (self Go) SetRequestedReferenceSpaceTypes(value string) {
	class(self).SetRequestedReferenceSpaceTypes(gd.NewString(value))
}

func (self Go) ReferenceSpaceType() string {
		return string(class(self).GetReferenceSpaceType().String())
}

func (self Go) EnabledFeatures() string {
		return string(class(self).GetEnabledFeatures().String())
}

func (self Go) VisibilityState() string {
		return string(class(self).GetVisibilityState().String())
}

/*
Checks if the given [param session_mode] is supported by the user's browser.
Possible values come from [url=https://developer.mozilla.org/en-US/docs/Web/API/XRSessionMode]WebXR's XRSessionMode[/url], including: [code]"immersive-vr"[/code], [code]"immersive-ar"[/code], and [code]"inline"[/code].
This method returns nothing, instead it emits the [signal session_supported] signal with the result.
*/
//go:nosplit
func (self class) IsSessionSupported(session_mode gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(session_mode))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebXRInterface.Bind_is_session_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetSessionMode(session_mode gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(session_mode))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebXRInterface.Bind_set_session_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSessionMode() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebXRInterface.Bind_get_session_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRequiredFeatures(required_features gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(required_features))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebXRInterface.Bind_set_required_features, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRequiredFeatures() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebXRInterface.Bind_get_required_features, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOptionalFeatures(optional_features gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(optional_features))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebXRInterface.Bind_set_optional_features, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOptionalFeatures() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebXRInterface.Bind_get_optional_features, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetReferenceSpaceType() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebXRInterface.Bind_get_reference_space_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetEnabledFeatures() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebXRInterface.Bind_get_enabled_features, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRequestedReferenceSpaceTypes(requested_reference_space_types gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(requested_reference_space_types))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebXRInterface.Bind_set_requested_reference_space_types, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRequestedReferenceSpaceTypes() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebXRInterface.Bind_get_requested_reference_space_types, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is an active input source with the given [param input_source_id].
*/
//go:nosplit
func (self class) IsInputSourceActive(input_source_id gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, input_source_id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebXRInterface.Bind_is_input_source_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets an [XRControllerTracker] for the given [param input_source_id].
In the context of WebXR, an input source can be an advanced VR controller like the Oculus Touch or Index controllers, or even a tap on the screen, a spoken voice command or a button press on the device itself. When a non-traditional input source is used, interpret the position and orientation of the [XRPositionalTracker] as a ray pointing at the object the user wishes to interact with.
Use this method to get information about the input source that triggered one of these signals:
- [signal selectstart]
- [signal select]
- [signal selectend]
- [signal squeezestart]
- [signal squeeze]
- [signal squeezestart]
*/
//go:nosplit
func (self class) GetInputSourceTracker(input_source_id gd.Int) gdclass.XRControllerTracker {
	var frame = callframe.New()
	callframe.Arg(frame, input_source_id)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebXRInterface.Bind_get_input_source_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.XRControllerTracker{classdb.XRControllerTracker(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Returns the target ray mode for the given [param input_source_id].
This can help interpret the input coming from that input source. See [url=https://developer.mozilla.org/en-US/docs/Web/API/XRInputSource/targetRayMode]XRInputSource.targetRayMode[/url] for more information.
*/
//go:nosplit
func (self class) GetInputSourceTargetRayMode(input_source_id gd.Int) classdb.WebXRInterfaceTargetRayMode {
	var frame = callframe.New()
	callframe.Arg(frame, input_source_id)
	var r_ret = callframe.Ret[classdb.WebXRInterfaceTargetRayMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebXRInterface.Bind_get_input_source_target_ray_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetVisibilityState() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebXRInterface.Bind_get_visibility_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the display refresh rate for the current HMD. Not supported on all HMDs and browsers. It may not report an accurate value until after using [method set_display_refresh_rate].
*/
//go:nosplit
func (self class) GetDisplayRefreshRate() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebXRInterface.Bind_get_display_refresh_rate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the display refresh rate for the current HMD. Not supported on all HMDs and browsers. It won't take effect right away until after [signal display_refresh_rate_changed] is emitted.
*/
//go:nosplit
func (self class) SetDisplayRefreshRate(refresh_rate gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, refresh_rate)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebXRInterface.Bind_set_display_refresh_rate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns display refresh rates supported by the current HMD. Only returned if this feature is supported by the web browser and after the interface has been initialized.
*/
//go:nosplit
func (self class) GetAvailableDisplayRefreshRates() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebXRInterface.Bind_get_available_display_refresh_rates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self Go) OnSessionSupported(cb func(session_mode string, supported bool)) {
	self[0].AsObject().Connect(gd.NewStringName("session_supported"), gd.NewCallable(cb), 0)
}


func (self Go) OnSessionStarted(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("session_started"), gd.NewCallable(cb), 0)
}


func (self Go) OnSessionEnded(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("session_ended"), gd.NewCallable(cb), 0)
}


func (self Go) OnSessionFailed(cb func(message string)) {
	self[0].AsObject().Connect(gd.NewStringName("session_failed"), gd.NewCallable(cb), 0)
}


func (self Go) OnSelectstart(cb func(input_source_id int)) {
	self[0].AsObject().Connect(gd.NewStringName("selectstart"), gd.NewCallable(cb), 0)
}


func (self Go) OnSelect(cb func(input_source_id int)) {
	self[0].AsObject().Connect(gd.NewStringName("select"), gd.NewCallable(cb), 0)
}


func (self Go) OnSelectend(cb func(input_source_id int)) {
	self[0].AsObject().Connect(gd.NewStringName("selectend"), gd.NewCallable(cb), 0)
}


func (self Go) OnSqueezestart(cb func(input_source_id int)) {
	self[0].AsObject().Connect(gd.NewStringName("squeezestart"), gd.NewCallable(cb), 0)
}


func (self Go) OnSqueeze(cb func(input_source_id int)) {
	self[0].AsObject().Connect(gd.NewStringName("squeeze"), gd.NewCallable(cb), 0)
}


func (self Go) OnSqueezeend(cb func(input_source_id int)) {
	self[0].AsObject().Connect(gd.NewStringName("squeezeend"), gd.NewCallable(cb), 0)
}


func (self Go) OnVisibilityStateChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("visibility_state_changed"), gd.NewCallable(cb), 0)
}


func (self Go) OnReferenceSpaceReset(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("reference_space_reset"), gd.NewCallable(cb), 0)
}


func (self Go) OnDisplayRefreshRateChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("display_refresh_rate_changed"), gd.NewCallable(cb), 0)
}


func (self class) AsWebXRInterface() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsWebXRInterface() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsXRInterface() XRInterface.GD { return *((*XRInterface.GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRInterface() XRInterface.Go { return *((*XRInterface.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsXRInterface(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsXRInterface(), name)
	}
}
func init() {classdb.Register("WebXRInterface", func(ptr gd.Object) any { return classdb.WebXRInterface(ptr) })}
type TargetRayMode = classdb.WebXRInterfaceTargetRayMode

const (
/*We don't know the the target ray mode.*/
	TargetRayModeUnknown TargetRayMode = 0
/*Target ray originates at the viewer's eyes and points in the direction they are looking.*/
	TargetRayModeGaze TargetRayMode = 1
/*Target ray from a handheld pointer, most likely a VR touch controller.*/
	TargetRayModeTrackedPointer TargetRayMode = 2
/*Target ray from touch screen, mouse or other tactile input device.*/
	TargetRayModeScreen TargetRayMode = 3
)
