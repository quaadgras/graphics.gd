/*
var interface = XRServer.find_interface("Native mobile")
if interface and interface.initialize():
    get_viewport().use_xr = true
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/classdb/XRInterface"
	"graphics.gd/classdb/XRServer"
)

func ExampleMobileVR(node Node.Instance) {
	var XR = XRServer.FindInterface("Native mobile")
	if XR != XRInterface.Nil && XR.Initialize() {
		Viewport.Get(node).SetUseXr(true)
	}
}
