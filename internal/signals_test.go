//go:build !generate

package gd_test

import (
	"testing"

	"graphics.gd/classdb/Node2D"
	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Signal"
)

type CustomSignal struct {
	Node2D.Extension[CustomSignal]

	HealthChanged chan<- func() (old, new int)

	GenericSignal Signal.Solo[int]

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
	custom := new(CustomSignal)
	custom.HealthChanged = make(chan func() (int, int), 1)
	custom.TakeDamage(10)
}

type CustomStringSignals struct {
	Node2D.Extension[CustomStringSignals]

	StringSignal Signal.Solo[string] `gd:"on_string(s)"`
}

func TestSignalString(t *testing.T) {
	custom := new(CustomStringSignals)
	signal := gd.NewSignalOf(custom.AsObject(), gd.NewStringName("on_string"))
	var triggered int
	if err := signal.Connect(gd.NewCallable(func(s string) {
		if s != "Hello World" {
			t.Fail()
		}
		triggered++
	}), 0); err != 0 {
		t.Fatalf("Failed to connect signal: %d", err)
	}
	custom.StringSignal.Attach(Callable.New(func(s string) {
		if s != "Hello World" {
			t.Fail()
		}
		triggered++
	}))
	Signal.Via(gd.SignalProxy{}, pointers.Pack(signal)).Emit(variant.New("Hello World"))
	signal.Emit(gd.NewVariant("Hello World"))
	if triggered != 4 {
		t.Fatalf("Expected 4 triggers, got %d", triggered)
	}

}
