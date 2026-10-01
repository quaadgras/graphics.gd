package main

import (
	"graphics.gd/classdb/AnimationPlayer"
	"graphics.gd/classdb/Area2D"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/variant/Object"
)

type Coin struct {
	Area2D.Extension[Coin]

	AnimationPlayer AnimationPlayer.Instance

	taken bool
}

func (c *Coin) OnBodyEnter(body Node2D.Instance) {
	if _, isPlayer := Object.As[*Player](body); !c.taken && isPlayer {
		c.AnimationPlayer.MoreArgs().Play("taken", -1, 1, false)
	}
}
