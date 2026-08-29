package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/DirectionalLight3D"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseButton"
	"graphics.gd/classdb/InputEventMouseMotion"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/classdb/WorldEnvironment"
	"graphics.gd/startup"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Euler"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/String"
)

const (
	RotSpeed    = 0.003
	ZoomSpeed   = 0.125
	MainButtons = Input.MouseButtonMaskLeft | Input.MouseButtonMaskRight | Input.MouseButtonMaskMiddle
)

type CSG struct {
	WorldEnvironment.Extension[CSG]

	DirectionalLight3D DirectionalLight3D.Instance
	Testers            Node3D.Instance
	CameraHolder       Node3D.Instance   // Has a position and rotates on Y.
	RotationX          Node3D.Instance   `gd:"CameraHolder/RotationX"`
	Camera             Camera3D.Instance `gd:"CameraHolder/RotationX/Camera3D"`
	TestName           Label.Instance
	Previous           Button.Instance
	Next               Button.Instance

	testerIndex    int
	rotX           Angle.Radians // This must be kept in sync with RotationX.
	rotY           Angle.Radians // This must be kept in sync with CameraHolder.
	cameraDistance Float.X
}

func NewCSG() *CSG {
	return &CSG{
		rotX:           -Angle.Tau / 16,
		rotY:           Angle.Tau / 8,
		cameraDistance: 4.0,
	}
}

func (c *CSG) Ready() {
	if RenderingServer.GetCurrentRenderingMethod() == "gl_compatibility" {
		// Darken the light's energy to compensate for sRGB blending (without affecting sky rendering).
		c.DirectionalLight3D.SetSkyMode(DirectionalLight3D.SkyModeSkyOnly)
		new_light, _ := Object.As[DirectionalLight3D.Instance](c.DirectionalLight3D.AsNode().Duplicate())
		new_light.AsLight3D().SetLightEnergy(0.3)
		new_light.SetSkyMode(DirectionalLight3D.SkyModeLightOnly)
		c.AsNode().AddChild(new_light.AsNode())
	}
	c.updateCameraRotation()
	c.UpdateGui()
}

func (c *CSG) updateCameraRotation() {
	var holder = c.CameraHolder.Transform()
	holder.Basis = Basis.FromEuler(Euler.Radians{Y: c.rotY}, Angle.OrderYXZ)
	c.CameraHolder.SetTransform(holder)
	var rotation_x = c.RotationX.Transform()
	rotation_x.Basis = Basis.FromEuler(Euler.Radians{X: c.rotX}, Angle.OrderYXZ)
	c.RotationX.SetTransform(rotation_x)
}

func (c *CSG) UnhandledInput(input_event InputEvent.Instance) {
	if input_event.IsActionPressed("ui_left") {
		c.OnPreviousPressed()
	}
	if input_event.IsActionPressed("ui_right") {
		c.OnNextPressed()
	}

	if button, ok := Object.As[InputEventMouseButton.Instance](input_event); ok {
		if button.ButtonIndex() == Input.MouseButtonWheelUp {
			c.cameraDistance -= ZoomSpeed
		}
		if button.ButtonIndex() == Input.MouseButtonWheelDown {
			c.cameraDistance += ZoomSpeed
		}
		c.cameraDistance = Float.Clamp(c.cameraDistance, 1.5, 6)
	}

	if motion, ok := Object.As[InputEventMouseMotion.Instance](input_event); ok && motion.AsInputEventMouse().ButtonMask()&MainButtons != 0 {
		// Use `screen_relative` to make mouse sensitivity independent of viewport resolution.
		var relative_motion = motion.ScreenRelative()
		c.rotY -= Angle.Radians(relative_motion.X * RotSpeed)
		c.rotX -= Angle.Radians(relative_motion.Y * RotSpeed)
		c.rotX = Float.Clamp(c.rotX, -1.57, 0)
		c.updateCameraRotation()
	}
}

func (c *CSG) Process(delta Float.X) {
	current_tester, _ := Object.As[Node3D.Instance](c.Testers.AsNode().GetChild(c.testerIndex))
	// This code assumes CameraHolder's X and Y coordinates are already correct.
	var current_position = c.CameraHolder.GlobalTransform().Origin.Z
	var target_position = current_tester.GlobalTransform().Origin.Z
	var holder = c.CameraHolder.GlobalTransform()
	holder.Origin.Z = Float.Lerp(current_position, target_position, 3*delta)
	c.CameraHolder.SetGlobalTransform(holder)
	var camera_position = c.Camera.AsNode3D().Position()
	camera_position.Z = Float.Lerp(camera_position.Z, c.cameraDistance, 10*delta)
	c.Camera.AsNode3D().SetPosition(camera_position)
}

func (c *CSG) OnPreviousPressed() {
	c.testerIndex = max(0, c.testerIndex-1)
	c.UpdateGui()
}

func (c *CSG) OnNextPressed() {
	c.testerIndex = min(c.testerIndex+1, c.Testers.AsNode().GetChildCount()-1)
	c.UpdateGui()
}

func (c *CSG) UpdateGui() {
	c.TestName.SetText(String.Capitalize(c.Testers.AsNode().GetChild(c.testerIndex).Name()))
	c.Previous.AsBaseButton().SetDisabled(c.testerIndex == 0)
	c.Next.AsBaseButton().SetDisabled(c.testerIndex == c.Testers.AsNode().GetChildCount()-1)
}

func main() {
	classdb.Register[CSG](NewCSG)
	startup.Scene()
}
