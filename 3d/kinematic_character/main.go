package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/CenterContainer"
	"graphics.gd/classdb/CharacterBody3D"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/PhysicsBody3D"
	"graphics.gd/classdb/ProjectSettings"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/startup"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector3"
)

const (
	MaxSpeed     = 3.5
	JumpSpeed    = 6.5
	Acceleration = 4
	Deceleration = 4
)

// Cubio is the player-controlled cube.
type Cubio struct {
	CharacterBody3D.Extension[Cubio] `gd:"Cubio"`

	Camera  *FollowCamera            `gd:"Target/Camera3D"`
	WinText CenterContainer.Instance `gd:"WinText"`

	gravity       Float.X
	startPosition Vector3.XYZ
}

func (c *Cubio) Ready() {
	c.gravity = -Float.X(ProjectSettings.GetSetting("physics/3d/default_gravity", 9.8).(float64))
	c.startPosition = c.AsNode3D().Position()
}

func (c *Cubio) PhysicsProcess(delta Float.X) {
	body := c.AsCharacterBody3D()
	node := c.AsNode3D()
	if Input.IsActionJustPressed("exit", false) {
		SceneTree.Get(c.AsNode()).Quit()
	}
	if Input.IsActionJustPressed("reset_position", false) || node.GlobalPosition().Y < -6.0 {
		// Pressed the reset key or fell off the ground.
		node.SetPosition(c.startPosition)
		body.SetVelocity(Vector3.Zero)
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
	camBasis := c.Camera.AsNode3D().GlobalTransform().Basis
	camBasis = Basis.Rotated(camBasis, camBasis.X, -Basis.AsEulerAngles(camBasis, 0).X)
	dir = Basis.Transform(dir, camBasis)

	// Limit the input to a length of 1. LengthSquared is faster to check than Length.
	if Vector3.LengthSquared(dir) > 1 {
		dir = Vector3.DivX(dir, Vector3.Length(dir))
	}

	// Apply gravity.
	velocity := body.Velocity()
	velocity.Y += delta * c.gravity

	// Using only the horizontal velocity, interpolate towards the input.
	hvel := velocity
	hvel.Y = 0

	target := Vector3.MulX(dir, MaxSpeed)
	var acceleration Float.X
	if Vector3.Dot(dir, hvel) > 0 {
		acceleration = Acceleration
	} else {
		acceleration = Deceleration
	}
	hvel = Vector3.Lerp(hvel, target, acceleration*delta)

	// Assign hvel's values back to velocity, and then move.
	velocity.X = hvel.X
	velocity.Z = hvel.Z
	body.SetVelocity(velocity)
	body.MoveAndSlide()

	// Jumping code. IsOnFloor must come after MoveAndSlide.
	if body.IsOnFloor() && Input.IsActionPressed("jump", false) {
		velocity = body.Velocity()
		velocity.Y = JumpSpeed
		body.SetVelocity(velocity)
	}
}

// OnTcubeBodyEntered is connected to the Princess area's body_entered signal.
func (c *Cubio) OnTcubeBodyEntered(body PhysicsBody3D.Instance) {
	if body.AsObject() == c.AsObject() {
		c.WinText.AsCanvasItem().Show()
	}
}

func main() {
	classdb.Register[Cubio]()
	classdb.Register[FollowCamera](NewFollowCamera)
	startup.Scene()
}
