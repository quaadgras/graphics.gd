package main

import (
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/Environment"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseMotion"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/ProjectSettings"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/startup"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Euler"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector3"
)

const (
	MouseSensitivity = 0.002
	MoveSpeed        = 0.6
)

type FogCamera struct {
	Camera3D.Extension[FogCamera]

	Label Label.Instance

	volumetricFogVolumeSize  int
	volumetricFogVolumeDepth int

	rot      Euler.Radians
	velocity Vector3.XYZ
}

func settingInt(name string) int {
	switch v := ProjectSettings.GetSetting(name, 0).(type) {
	case int:
		return v
	case int64:
		return int(v)
	case float64:
		return int(v)
	}
	return 0
}

func (c *FogCamera) environment() Environment.Instance {
	return c.AsNode3D().GetWorld3d().Environment()
}

func (c *FogCamera) Ready() {
	c.volumetricFogVolumeSize = settingInt("rendering/environment/volumetric_fog/volume_size")
	c.volumetricFogVolumeDepth = settingInt("rendering/environment/volumetric_fog/volume_depth")
	Input.SetMouseMode(Input.MouseModeCaptured)
	c.UpdateLabel()
}

func (c *FogCamera) Process(delta Float.X) {
	node := c.AsNode3D()
	motion := Vector3.XYZ{
		X: Input.GetActionStrength("move_right", false) - Input.GetActionStrength("move_left", false),
		Y: 0,
		Z: Input.GetActionStrength("move_back", false) - Input.GetActionStrength("move_forward", false),
	}

	// Normalize motion to prevent diagonal movement from being
	// `sqrt(2)` times faster than straight movement.
	motion = Vector3.Normalized(motion)

	c.velocity = Vector3.Add(c.velocity, Vector3.MulX(Basis.Transform(motion, node.Transform().Basis), MoveSpeed*delta))
	c.velocity = Vector3.MulX(c.velocity, 0.85)
	node.SetPosition(Vector3.Add(node.Position(), c.velocity))
}

func (c *FogCamera) Input(event InputEvent.Instance) {
	node := c.AsNode3D()
	// Mouse look (only if the mouse is captured).
	if motion, ok := Object.As[InputEventMouseMotion.Instance](event); ok && Input.MouseMode() == Input.MouseModeCaptured {
		// Horizontal mouse look.
		c.rot.Y -= Angle.Radians(motion.ScreenRelative().X * MouseSensitivity)
		// Vertical mouse look.
		c.rot.X = Float.Clamp(c.rot.X-Angle.Radians(motion.ScreenRelative().Y*MouseSensitivity), -1.57, 1.57)
		transform := node.Transform()
		transform.Basis = Basis.FromEuler(c.rot, Angle.OrderYXZ)
		node.SetTransform(transform)
	}

	if event.IsActionPressed("toggle_mouse_capture") {
		if Input.MouseMode() == Input.MouseModeCaptured {
			Input.SetMouseMode(Input.MouseModeVisible)
		} else {
			Input.SetMouseMode(Input.MouseModeCaptured)
		}
	}

	env := c.environment()
	switch {
	case event.IsActionPressed("toggle_temporal_reprojection"):
		env.SetVolumetricFogTemporalReprojectionEnabled(!env.VolumetricFogTemporalReprojectionEnabled())
		c.UpdateLabel()
	case event.IsActionPressed("increase_temporal_reprojection"):
		env.SetVolumetricFogTemporalReprojectionAmount(Float.Clamp(env.VolumetricFogTemporalReprojectionAmount()+0.01, 0.5, 0.99))
		c.UpdateLabel()
	case event.IsActionPressed("decrease_temporal_reprojection"):
		env.SetVolumetricFogTemporalReprojectionAmount(Float.Clamp(env.VolumetricFogTemporalReprojectionAmount()-0.01, 0.5, 0.99))
		c.UpdateLabel()
	case event.IsActionPressed("increase_fog_density"):
		env.SetVolumetricFogDensity(Float.Clamp(env.VolumetricFogDensity()+0.01, 0.0, 1.0))
		c.UpdateLabel()
	case event.IsActionPressed("decrease_fog_density"):
		env.SetVolumetricFogDensity(Float.Clamp(env.VolumetricFogDensity()-0.01, 0.0, 1.0))
		c.UpdateLabel()
	case event.IsActionPressed("increase_volumetric_fog_quality"):
		c.volumetricFogVolumeSize = min(max(c.volumetricFogVolumeSize+16, 16), 384)
		c.volumetricFogVolumeDepth = min(max(c.volumetricFogVolumeDepth+16, 16), 384)
		RenderingServer.EnvironmentSetVolumetricFogVolumeSize(c.volumetricFogVolumeSize, c.volumetricFogVolumeDepth)
		c.UpdateLabel()
	case event.IsActionPressed("decrease_volumetric_fog_quality"):
		c.volumetricFogVolumeSize = min(max(c.volumetricFogVolumeSize-16, 16), 384)
		c.volumetricFogVolumeDepth = min(max(c.volumetricFogVolumeDepth-16, 16), 384)
		RenderingServer.EnvironmentSetVolumetricFogVolumeSize(c.volumetricFogVolumeSize, c.volumetricFogVolumeDepth)
		c.UpdateLabel()
	}
}

func (c *FogCamera) UpdateLabel() {
	env := c.environment()
	if env.VolumetricFogTemporalReprojectionEnabled() {
		c.Label.SetText(fmt.Sprintf("Fog density: %.2f\nTemporal reprojection: Enabled\nTemporal reprojection strength: %.2f\nVolumetric fog quality: %d×%d×%d",
			env.VolumetricFogDensity(),
			env.VolumetricFogTemporalReprojectionAmount(),
			c.volumetricFogVolumeSize,
			c.volumetricFogVolumeSize,
			c.volumetricFogVolumeDepth,
		))
	} else {
		c.Label.SetText(fmt.Sprintf("Fog density: %.2f\nTemporal reprojection: Disabled\nVolumetric fog quality: %d×%d×%d",
			env.VolumetricFogDensity(),
			c.volumetricFogVolumeSize,
			c.volumetricFogVolumeSize,
			c.volumetricFogVolumeDepth,
		))
	}
}

func main() {
	classdb.Register[FogCamera]()
	startup.Scene()
}
