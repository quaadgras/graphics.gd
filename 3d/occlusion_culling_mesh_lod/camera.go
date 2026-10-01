package main

import (
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseMotion"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Euler"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector3"
)

const (
	MouseSensitivity = 0.002
	MoveSpeed        = 1.5
)

type FlyCamera struct {
	Camera3D.Extension[FlyCamera]

	rot      Euler.Radians
	velocity Vector3.XYZ
}

func (c *FlyCamera) Ready() {
	Input.SetMouseMode(Input.MouseModeCaptured)
}

func (c *FlyCamera) Input(input_event InputEvent.Instance) {
	node := c.AsNode3D()
	// Mouse look (only if the mouse is captured).
	if motion, ok := Object.As[InputEventMouseMotion.Instance](input_event); ok && Input.MouseMode() == Input.MouseModeCaptured {
		// Horizontal mouse look.
		c.rot.Y -= Angle.Radians(motion.ScreenRelative().X * MouseSensitivity)
		// Vertical mouse look.
		c.rot.X = Float.Clamp(c.rot.X-Angle.Radians(motion.ScreenRelative().Y*MouseSensitivity), -1.57, 1.57)
		transform := node.Transform()
		transform.Basis = Basis.FromEuler(c.rot, Angle.OrderYXZ)
		node.SetTransform(transform)
	}

	if input_event.IsActionPressed("toggle_mouse_capture") {
		if Input.MouseMode() == Input.MouseModeCaptured {
			Input.SetMouseMode(Input.MouseModeVisible)
		} else {
			Input.SetMouseMode(Input.MouseModeCaptured)
		}
	}
}

func (c *FlyCamera) Process(delta Float.X) {
	node := c.AsNode3D()
	motion := Vector3.XYZ{
		X: Input.GetAxis("move_left", "move_right"),
		Y: 0,
		Z: Input.GetAxis("move_forward", "move_back"),
	}

	// Normalize motion to prevent diagonal movement from being
	// `sqrt(2)` times faster than straight movement.
	motion = Vector3.Normalized(motion)

	c.velocity = Vector3.Add(c.velocity, Vector3.MulX(Basis.Transform(motion, node.Transform().Basis), MoveSpeed*delta))
	c.velocity = Vector3.MulX(c.velocity, 0.85)
	node.SetPosition(Vector3.Add(node.Position(), c.velocity))
}
