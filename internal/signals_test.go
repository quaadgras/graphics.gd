//go:build !generate

package gd_test

import (
	"testing"

	"grow.graphics/gd"
	internal "grow.graphics/gd/internal"
	"grow.graphics/gd/object/Node2D"
)

type CustomSignal struct {
	gd.Class[CustomSignal, Node2D.Expert]

	HealthChanged gd.SignalAs[func(gd.Int, gd.Int)]

	Health gd.Int
}

func (c *CustomSignal) Ready() {
	if c.Health == 0 {
		c.Health = 10
	}
}

func (c *CustomSignal) TakeDamage(godot gd.Lifetime, amount gd.Int) {
	oldHealth := c.Health
	c.Health -= amount
	c.HealthChanged.Emit(oldHealth, c.Health)
}

func TestSignals(t *testing.T) {
	godot := internal.NewLifetime(API)
	defer godot.End()

	gd.Register[CustomSignal](godot)

	custom := gd.New[CustomSignal](godot)
	custom.TakeDamage(godot, 10)
}
