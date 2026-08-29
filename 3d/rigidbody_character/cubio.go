package main

import (
	"graphics.gd/classdb/CenterContainer"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/RigidBody3D"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/ShapeCast3D"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector3"
)

type Cubio struct {
	RigidBody3D.Extension[Cubio]

	ShapeCast ShapeCast3D.Instance     `gd:"ShapeCast3D"`
	Camera    *FollowCamera            `gd:"Target/Camera3D"`
	WinText   CenterContainer.Instance `gd:"WinText"`

	startPosition Vector3.XYZ
}

func (c *Cubio) Ready() {
	c.startPosition = c.AsNode3D().Position()
}

func (c *Cubio) PhysicsProcess(delta Float.X) {
	body := c.AsRigidBody3D()
	node3d := c.AsNode3D()
	if Input.IsActionJustPressed("exit", false) {
		SceneTree.Get(c.AsNode()).Quit()
	}
	if Input.IsActionJustPressed("reset_position", false) || node3d.GlobalPosition().Y < -6.0 {
		// Pressed the reset key or fell off the ground.
		node3d.SetPosition(c.startPosition)
		body.SetLinearVelocity(Vector3.Zero)
		// We teleported the player on the lines above. Reset interpolation
		// to prevent it from interpolating from the old player position
		// to the new position.
		c.AsNode().ResetPhysicsInterpolation()
	}

	var dir Vector3.XYZ
	dir.X = Input.GetAxis("move_left", "move_right")
	dir.Z = Input.GetAxis("move_forward", "move_back")

	// Get the camera's transform basis, but remove the X rotation such
	// that the Y axis is up and Z is horizontal.
	var cam_basis = c.Camera.AsNode3D().GlobalTransform().Basis
	cam_basis = Basis.Rotated(cam_basis, cam_basis.X, -Basis.AsEulerAngles(cam_basis, Angle.OrderYXZ).X)
	dir = Basis.Transform(dir, cam_basis)

	// Air movement.
	body.ApplyCentralImpulse(Vector3.MulX(Vector3.Normalized(dir), 5.0*delta))

	if c.OnGround() {
		// Ground movement (higher acceleration).
		body.ApplyCentralImpulse(Vector3.MulX(Vector3.Normalized(dir), 10.0*delta))

		// Jumping code.
		// It's acceptable to set `linear_velocity` here as it's only set once, rather than continuously.
		// Vertical speed is set (rather than added) to prevent jumping higher than intended
		// if the ShapeCast3D collides for multiple frames.
		if Input.IsActionPressed("jump", false) {
			velocity := body.LinearVelocity()
			velocity.Y = 7
			body.SetLinearVelocity(velocity)
		}
	}
}

// OnGround tests if there is a body below the player.
func (c *Cubio) OnGround() bool {
	return c.ShapeCast.IsColliding()
}

func (c *Cubio) OnTcubeBodyEntered(body Node.Instance) {
	if body.AsObject() == c.AsObject() {
		c.WinText.AsCanvasItem().SetVisible(true)
	}
}
