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

// TestSignalDisconnect tests that a callable created with Callable.New can be
// disconnected from a Godot signal after being connected, using the same
// Callable.Function value for both operations. This is the bug reported in
// https://github.com/quaadgras/graphics.gd/discussions/263
func TestSignalDisconnect(t *testing.T) {
	custom := new(CustomStringSignals)
	signal := gd.NewSignalOf(custom.AsObject(), gd.NewStringName("on_string"))

	var callCount int
	do := func(s string) {
		callCount++
	}
	mycb := Callable.New(do)

	// Connect via InternalCallable (this is what SignalProxy.Attach does)
	c1 := gd.InternalCallable(mycb)
	if err := signal.Connect(c1, 0); err != 0 {
		t.Fatalf("Failed to connect signal: %d", err)
	}

	// Emit and verify the callable fires
	signal.Emit(gd.NewVariant("test"))
	if callCount != 1 {
		t.Fatalf("Expected 1 call after connect, got %d", callCount)
	}

	// Disconnect via InternalCallable (this is what SignalProxy.Remove does)
	c2 := gd.InternalCallable(mycb)
	signal.Disconnect(c2)

	// Emit again and verify the callable no longer fires
	signal.Emit(gd.NewVariant("test2"))
	if callCount != 1 {
		t.Fatalf("Expected callCount to remain 1 after disconnect, got %d", callCount)
	}
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
	Callable.Cycle()
	if triggered != 4 {
		t.Fatalf("Expected 4 triggers, got %d", triggered)
	}

}
