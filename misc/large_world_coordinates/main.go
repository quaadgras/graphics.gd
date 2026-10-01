package main

import (
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/CheckButton"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseButton"
	"graphics.gd/classdb/InputEventMouseMotion"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/OS"
	"graphics.gd/classdb/RichTextLabel"
	"graphics.gd/classdb/VBoxContainer"
	"graphics.gd/startup"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Euler"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector3"
)

const (
	RotSpeed    = 0.003
	ZoomSpeed   = 0.5
	MainButtons = Input.MouseButtonMaskLeft | Input.MouseButtonMaskMiddle | Input.MouseButtonMaskRight
)

type Controls struct {
	VBoxContainer.Extension[Controls] `gd:"Controls"`

	Camera       Camera3D.Instance `gd:"../Move/CameraHolder/RotationX/Camera3D"`
	CameraHolder Node3D.Instance   `gd:"../Move/CameraHolder"`
	RotationX    Node3D.Instance   `gd:"../Move/CameraHolder/RotationX"`
	NodeToMove   Node3D.Instance   `gd:"../Move"`

	HelpLabel   Label.Instance         `gd:"%HelpLabel"`
	Coordinates RichTextLabel.Instance `gd:"%Coordinates"`
	IncrementX  CheckButton.Instance   `gd:"%IncrementX"`
	IncrementY  CheckButton.Instance   `gd:"%IncrementY"`
	IncrementZ  CheckButton.Instance   `gd:"%IncrementZ"`

	zoom Float.X
	rotX Angle.Radians
	rotY Angle.Radians
}

func (c *Controls) Ready() {
	c.zoom = c.Camera.AsNode3D().Position().Z
	c.rotX = c.RotationX.Rotation().X
	c.rotY = c.CameraHolder.Rotation().Y

	if OS.HasFeature("double") {
		c.HelpLabel.SetText("Double precision is enabled in this engine build.\nNo shaking should occur at high coordinate levels\n(±65,536 or more on any axis).")
		c.HelpLabel.AsControl().AddThemeColorOverride("font_color", Color.RGBA{0.667, 1, 0.667, 1})
	}
}

func (c *Controls) Process(delta Float.X) {
	position := c.NodeToMove.Position()
	c.Coordinates.SetText(fmt.Sprintf(
		"X: [color=#fb9]%f[/color]\nY: [color=#bfa]%f[/color]\nZ: [color=#9cf]%f[/color]",
		position.X, position.Y, position.Z))
	if c.IncrementX.AsBaseButton().ButtonPressed() {
		position.X += 10_000 * delta
	}
	if c.IncrementY.AsBaseButton().ButtonPressed() {
		position.Y += 100_000 * delta
	}
	if c.IncrementZ.AsBaseButton().ButtonPressed() {
		position.Z += 1_000_000 * delta
	}
	c.NodeToMove.SetPosition(position)
}

func (c *Controls) Input(event InputEvent.Instance) {
	if button, ok := Object.As[InputEventMouseButton.Instance](event); ok {
		if button.ButtonIndex() == Input.MouseButtonWheelUp {
			c.zoom -= ZoomSpeed
		}
		if button.ButtonIndex() == Input.MouseButtonWheelDown {
			c.zoom += ZoomSpeed
		}
		c.zoom = Float.Clamp(c.zoom, 4, 15)
		position := c.Camera.AsNode3D().Position()
		position.Z = c.zoom
		c.Camera.AsNode3D().SetPosition(position)
	}

	if motion, ok := Object.As[InputEventMouseMotion.Instance](event); ok && motion.AsInputEventMouse().ButtonMask()&MainButtons != 0 {
		// Use ScreenRelative to make mouse sensitivity independent of viewport resolution.
		relative := motion.ScreenRelative()
		c.rotY -= Angle.Radians(relative.X * RotSpeed)
		c.rotX -= Angle.Radians(relative.Y * RotSpeed)
		c.rotX = Angle.Radians(Float.Clamp(Float.X(c.rotX), -1.4, 0.16))

		holder := c.CameraHolder.Transform()
		holder.Basis = Basis.FromEuler(Euler.Radians{Y: c.rotY}, Angle.OrderYXZ)
		c.CameraHolder.SetTransform(holder)

		rotation := c.RotationX.Transform()
		rotation.Basis = Basis.FromEuler(Euler.Radians{X: c.rotX}, Angle.OrderYXZ)
		c.RotationX.SetTransform(rotation)
	}
}

func (c *Controls) OnGoToButtonPressed(xPosition int) {
	if xPosition == 0 {
		// Reset all coordinates, not just X.
		c.NodeToMove.SetPosition(Vector3.Zero)
	} else {
		position := c.NodeToMove.Position()
		position.X = Float.X(xPosition)
		c.NodeToMove.SetPosition(position)
	}
}

func (c *Controls) OnOpenDocumentationPressed() {
	OS.ShellOpen("https://docs.godotengine.org/en/latest/tutorials/physics/large_world_coordinates.html")
}

func main() {
	classdb.Register[Controls]()
	startup.Scene()
}
