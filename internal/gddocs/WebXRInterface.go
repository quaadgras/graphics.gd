/*
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
*/

package main

import (
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/OS"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/classdb/WebXRInterface"
	"graphics.gd/classdb/XRServer"
	"graphics.gd/variant/Object"
)

type NodeWebXR struct {
	Node3D.Extension[NodeWebXR]

	Button Button.Instance

	webxr_interface WebXRInterface.Instance
	vr_supported    bool
}

func (n *NodeWebXR) Ready() {
	n.Button.AsBaseButton().OnPressed(func() {
		if !n.vr_supported {
			OS.Alert("Your browser doesn't support VR")
			return
		}
		// We want an immersive VR session, as opposed to AR ('immersive-ar') or a
		// simple 3DoF viewer ('viewer').
		n.webxr_interface.SetSessionMode("immersive-vr")
		// 'bounded-floor' is room scale, 'local-floor' is a standing or sitting
		// experience (it puts you 1.6m above the ground if you have 3DoF headset),
		// whereas as 'local' puts you down at the XROrigin.
		// This list means it'll first try to request 'bounded-floor', then
		// fallback on 'local-floor' and ultimately 'local', if nothing else is
		// supported.
		n.webxr_interface.SetRequestedReferenceSpaceTypes("bounded-floor, local-floor, local")
		// In order to use 'local-floor' or 'bounded-floor' we must also
		// mark the features as required or optional. By including 'hand-tracking'
		// as an optional feature, it will be enabled if supported.
		n.webxr_interface.SetRequiredFeatures("local-floor")
		n.webxr_interface.SetOptionalFeatures("bounded-floor, hand-tracking")

		// This will return false if we're unable to even request the session,
		// however, it can still fail asynchronously later in the process, so we
		// only know if it's really succeeded or failed when our
		// _webxr_session_started() or _webxr_session_failed() methods are called.
		if !n.webxr_interface.AsXRInterface().Initialize() {
			OS.Alert("Failed to initialize")
			return
		}
	})

	n.webxr_interface = Object.To[WebXRInterface.Instance](XRServer.FindInterface("WebXR"))
	if n.webxr_interface != WebXRInterface.Nil {
		// WebXR uses a lot of asynchronous callbacks, so we connect to various
		// signals in order to receive them.
		n.webxr_interface.OnSessionSupported(func(session_mode string, supported bool) {
			if session_mode == "immersive-vr" {
				n.vr_supported = supported
			}
		})
		n.webxr_interface.OnSessionStarted(func() {
			n.Button.AsCanvasItem().SetVisible(false)
			// This tells Godot to start rendering to the headset.
			Viewport.Get(n.AsNode()).SetUseXr(true)
			// This will be the reference space type you ultimately got, out of the
			// types that you requested above. This is useful if you want the game to
			// work a little differently in 'bounded-floor' versus 'local-floor'.
			print("Reference space type: ", n.webxr_interface.ReferenceSpaceType())
			// This will be the list of features that were successfully enabled
			// (except on browsers that don't support this property).
			print("Enabled features: ", n.webxr_interface.EnabledFeatures())
		})
		n.webxr_interface.OnSessionEnded(func() {
			n.Button.AsCanvasItem().SetVisible(true)
			// If the user exits immersive mode, then we tell Godot to render to the web
			// page again.
			Viewport.Get(n.AsNode()).SetUseXr(false)
		})
		n.webxr_interface.OnSessionFailed(func(message string) {
			OS.Alert("Failed to initialize: " + message)
		})
	}
}
