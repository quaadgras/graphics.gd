package main

import (
	"graphics.gd/classdb/AnimationPlayer"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/RigidBody2D"
	"graphics.gd/variant/Object"
)

type Bullet struct {
	RigidBody2D.Extension[Bullet]

	AnimationPlayer AnimationPlayer.Instance
}

func (b *Bullet) Destroy() {
	b.AnimationPlayer.MoreArgs().Play("destroy", -1, 1, false)
}

func (b *Bullet) OnBodyEntered(body Node.Instance) {
	if enemy, ok := Object.As[*Enemy](body); ok {
		enemy.Destroy()
	}
}
