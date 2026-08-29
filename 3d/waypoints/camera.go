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
	MoveSpeed        = 0.65
)

type Camera struct {
	Camera3D.Extension[Camera]

	rot      Euler.Radians
	velocity Vector3.XYZ
}

func (c *Camera) Ready() {
	Input.SetMouseMode(Input.MouseModeCaptured)
}

func (c *Camera) Input(event InputEvent.Instance) {
	// Mouse look (only if the mouse is captured).
	if motion, ok := Object.As[InputEventMouseMotion.Instance](event); ok && Input.MouseMode() == Input.MouseModeCaptured {
		// Horizontal mouse look.
		c.rot.Y -= Angle.Radians(motion.ScreenRelative().X * MouseSensitivity)
		// Vertical mouse look.
		c.rot.X = Angle.Radians(Float.Clamp(Float.X(c.rot.X)-motion.ScreenRelative().Y*MouseSensitivity, -1.57, 1.57))
		transform := c.AsNode3D().Transform()
		transform.Basis = Basis.FromEuler(c.rot, Angle.OrderYXZ)
		c.AsNode3D().SetTransform(transform)
	}
	if event.IsActionPressed("toggle_mouse_capture") {
		if Input.MouseMode() == Input.MouseModeCaptured {
			Input.SetMouseMode(Input.MouseModeVisible)
		} else {
			Input.SetMouseMode(Input.MouseModeCaptured)
		}
	}
}

func (c *Camera) Process(delta Float.X) {
	motion := Vector3.XYZ{
		X: Input.GetAxis("move_left", "move_right"),
		Y: 0,
		Z: Input.GetAxis("move_forward", "move_back"),
	}
	// Normalize motion to prevent diagonal movement from being
	// `sqrt(2)` times faster than straight movement.
	motion = Vector3.Normalized(motion)

	node3d := c.AsNode3D()
	c.velocity = Vector3.Add(c.velocity, Vector3.MulX(Basis.Transform(motion, node3d.Transform().Basis), MoveSpeed*delta))
	c.velocity = Vector3.MulX(c.velocity, 0.85)
	node3d.SetPosition(Vector3.Add(node3d.Position(), c.velocity))
}
