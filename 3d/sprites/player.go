package main

import (
	"graphics.gd/classdb/AnimatedSprite3D"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/Sprite3D"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector3"
)

// SpritePlayer moves an AnimatedSprite3D around with the move_* actions.
type SpritePlayer struct {
	Node3D.Extension[SpritePlayer]

	Sprite AnimatedSprite3D.Instance `gd:"AnimatedSprite3D"`

	MoveSpeed Float.X

	velocity Vector3.XYZ
}

func NewSpritePlayer() *SpritePlayer { return &SpritePlayer{MoveSpeed: 5.0} }

func (p *SpritePlayer) Process(delta Float.X) {
	input := Vector2.XY{
		X: Input.GetActionStrength("move_right", false) - Input.GetActionStrength("move_left", false),
		Y: Input.GetActionStrength("move_back", false) - Input.GetActionStrength("move_forward", false),
	}

	if Vector2.Length(input) > 0 {
		input = Vector2.Normalized(input)
		// Update player's position in 3D space.
		p.velocity = Vector3.MulX(Vector3.XYZ{X: input.X, Y: 0, Z: input.Y}, p.MoveSpeed)
		p.AsNode3D().Translate(Vector3.MulX(p.velocity, delta))

		// Play corresponding animation.
		var animation string
		switch {
		case Float.Abs(input.X) > Float.Abs(input.Y) && input.X > 0:
			animation = "walk_right"
		case Float.Abs(input.X) > Float.Abs(input.Y):
			animation = "walk_left"
		case input.Y > 0:
			animation = "walk_down"
		default:
			animation = "walk_up"
		}
		p.Sprite.MoreArgs().Play(animation, 1.0, false)
	} else {
		p.Sprite.Stop()
	}
}

// SpriteRotate spins a Sprite3D around its Y axis.
type SpriteRotate struct {
	Sprite3D.Extension[SpriteRotate]

	SpeedDeg Float.X
}

func NewSpriteRotate() *SpriteRotate { return &SpriteRotate{SpeedDeg: 90.0} }

func (s *SpriteRotate) Process(delta Float.X) {
	s.AsNode3D().RotateY(Angle.InRadians(Angle.Degrees(s.SpeedDeg * delta)))
}
