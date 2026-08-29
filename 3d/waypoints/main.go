package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/DirectionalLight3D"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/startup"
	"graphics.gd/variant/Object"
)

type Main struct {
	Node3D.Extension[Main]

	Sun DirectionalLight3D.Instance
}

func (m *Main) Ready() {
	if RenderingServer.GetCurrentRenderingMethod() == "gl_compatibility" {
		// Use PCF13 shadow filtering to improve quality (Medium maps to PCF5 instead).
		RenderingServer.DirectionalSoftShadowFilterSetQuality(RenderingServer.ShadowQualitySoftHigh)

		// Darken the light's energy to compensate for sRGB blending (without affecting sky rendering).
		m.Sun.SetSkyMode(DirectionalLight3D.SkyModeSkyOnly)
		new_light, _ := Object.As[DirectionalLight3D.Instance](m.Sun.AsNode().Duplicate())
		new_light.AsLight3D().SetLightEnergy(0.35)
		new_light.SetSkyMode(DirectionalLight3D.SkyModeLightOnly)
		m.AsNode().AddChild(new_light.AsNode())
	}
}

func main() {
	classdb.Register[Main]()
	classdb.Register[Camera]()
	classdb.Register[Waypoint](NewWaypoint)
	startup.Scene()
}
