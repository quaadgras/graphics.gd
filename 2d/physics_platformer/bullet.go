package main

import (
	"graphics.gd/classdb/AnimationPlayer"
	"graphics.gd/classdb/RigidBody2D"
	"graphics.gd/classdb/Timer"
)

type Bullet struct {
	RigidBody2D.Extension[Bullet]

	Timer           Timer.Instance
	AnimationPlayer AnimationPlayer.Instance

	disabled bool
}

func (b *Bullet) Ready() {
	b.Timer.Start()
}

func (b *Bullet) Disable() {
	if b.disabled {
		return
	}
	b.AnimationPlayer.MoreArgs().Play("shutdown", -1, 1, false)
	b.disabled = true
}
