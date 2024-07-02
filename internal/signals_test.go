//go:build !generate

package gd_test

import (
	"testing"

	"grow.graphics/gd"
	internal "grow.graphics/gd/internal"
)

type CustomSignal struct {
	gd.Class[CustomSignal, gd.Node2D]

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
	godot := internal.NewContext(API)
	defer godot.End()

	gd.Register[CustomSignal](godot)

	custom := gd.Create(godot, new(CustomSignal))
	custom.TakeDamage(godot, 10)
}
