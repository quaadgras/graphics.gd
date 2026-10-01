package main

import (
	"graphics.gd/classdb/AnimationPlayer"
	"graphics.gd/classdb/Area2D"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/variant/Object"
)

// Coin is a collectible that disappears when the player touches it.
type Coin struct {
	Area2D.Extension[Coin]

	AnimationPlayer AnimationPlayer.Instance
}

func (c *Coin) OnBodyEntered(body Node2D.Instance) {
	c.AnimationPlayer.MoreArgs().Play("picked", -1, 1, false)
	if player, ok := Object.As[*Player](body); ok {
		player.CoinCollected.Emit()
	}
}
