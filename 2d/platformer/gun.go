package main

import (
	"graphics.gd/classdb/AudioStreamPlayer2D"
	"graphics.gd/classdb/Marker2D"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/Timer"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
)

const BulletVelocity = 850.0

var BulletScene = Resource.Load[PackedScene.Is[*Bullet]]("res://player/bullet.tscn")

// Gun represents a weapon that spawns and shoots bullets.
// The Cooldown timer controls the cooldown duration between shots.
type Gun struct {
	Marker2D.Extension[Gun]

	SoundShoot AudioStreamPlayer2D.Instance `gd:"Shoot"`
	Timer      Timer.Instance               `gd:"Cooldown"`
}

// Shoot is only called by Player.
func (g *Gun) Shoot(direction Float.X) bool {
	if !g.Timer.IsStopped() {
		return false
	}
	bullet := PackedScene.Instantiate[*Bullet](PackedScene.Instance(BulletScene))
	bullet.AsNode2D().SetGlobalPosition(g.AsNode2D().GlobalPosition())
	bullet.AsRigidBody2D().SetLinearVelocity(Vector2.XY{X: direction * BulletVelocity, Y: 0})

	bullet.AsCanvasItem().SetTopLevel(true)
	g.AsNode().AddChild(bullet.AsNode())
	g.SoundShoot.Play()
	g.Timer.Start()
	return true
}
