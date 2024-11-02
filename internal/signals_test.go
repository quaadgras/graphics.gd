//go:build !generate

package gd_test

import (
	"testing"

	"grow.graphics/gd/gdextension"
	"grow.graphics/gd/objects/Node2D"
)

type CustomSignal struct {
	gdextension.Class[CustomSignal, Node2D.Instance]

	HealthChanged chan<- func() (old, new int)

	Health int
}

func (c *CustomSignal) Ready() {
	if c.Health == 0 {
		c.Health = 10
	}
}

func (c *CustomSignal) TakeDamage(amount int) {
	oldHealth := c.Health
	c.Health -= amount
	c.HealthChanged <- func() (int, int) { return oldHealth, c.Health }
}

func TestSignals(t *testing.T) {
	gdextension.Register[CustomSignal]()

	custom := new(CustomSignal)
	custom.HealthChanged = make(chan func() (int, int), 1)
	custom.TakeDamage(10)
}
