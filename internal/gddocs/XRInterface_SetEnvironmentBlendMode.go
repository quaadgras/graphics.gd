/*
func _ready():
	var xr_interface = XRServer.find_interface("OpenXR")
	if xr_interface and xr_interface.is_initialized():
		var vp = get_viewport()
		vp.use_xr = true
		var acceptable_modes = [XRInterface.XR_ENV_BLEND_MODE_OPAQUE, XRInterface.XR_ENV_BLEND_MODE_ADDITIVE]
		var modes = xr_interface.get_supported_environment_blend_modes()
		for mode in acceptable_modes:
			if mode in modes:
				xr_interface.set_environment_blend_mode(mode)
				break
*/

package main

import (
	"slices"

	"graphics.gd/classdb/Viewport"
	"graphics.gd/classdb/XRInterface"
	"graphics.gd/classdb/XRServer"
)

func XRInterface_SetEnvironmentBlendMode() {
	var xrInterface = XRServer.FindInterface("OpenXR")
	if xrInterface != XRInterface.Nil && xrInterface.IsInitialized() {
		var vp = Viewport.Get(node)
		vp.SetUseXr(true)
		var acceptableModes = []XRInterface.EnvironmentBlendMode{XRInterface.XrEnvBlendModeOpaque, XRInterface.XrEnvBlendModeAdditive}
		var modes = xrInterface.GetSupportedEnvironmentBlendModes()
		for _, mode := range acceptableModes {
			if slices.Contains(modes, mode) {
				xrInterface.SetEnvironmentBlendMode(mode)
				break
			}
		}
	}
}
