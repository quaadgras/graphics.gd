package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/CharacterBody2D"
	"graphics.gd/classdb/Input"
	"graphics.gd/startup"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
)

type Troll struct {
	CharacterBody2D.Extension[Troll]
}

const (
	MotionSpeed    = 30
	FrictionFactor = 0.89
)

var Tan30Deg = Angle.Tan(Angle.InRadians(30))

func (t *Troll) PhysicsProcess(delta Float.X) {
	var motion Vector2.XY
	motion.X = Input.GetAxis("move_left", "move_right")
	motion.Y = Input.GetAxis("move_up", "move_down")
	motion.Y *= Tan30Deg
	velocity := t.AsCharacterBody2D().Velocity()
	velocity = Vector2.Add(velocity, Vector2.MulX(Vector2.Normalized(motion), MotionSpeed))
	velocity = Vector2.MulX(velocity, FrictionFactor)
	t.AsCharacterBody2D().SetVelocity(velocity)
	t.AsCharacterBody2D().MoveAndSlide()
}

func main() {
	classdb.Register[Troll]()
	startup.Scene()
}
