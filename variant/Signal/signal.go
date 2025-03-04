// Package Signal provides a type representing an event stream or pub/sub message queue.
package Signal

import (
	"iter"

	"graphics.gd/variant"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/String"
)

// Any is a multi-producer, multi-consumer channel. Unlike with standard channels, consumers
// are not guaranteed to be running on a different goroutine, so emitting a signal may or may not
// be synchronous or asynchronous depending on the underlying [API] implementation. Any number of
// [variant.Any] values can be sent in a single send operation.
//
// Also known as a Pub/Sub and/or message queue.
type Any struct {
	state complex128
	proxy API
}

func (any *Any) SetAny(other Any) { *any = other }

type Pointer interface {
	SetAny(Any)
}

// Nil reference value, do not modify.
var Nil Any

// Consumer represents a connection between a signal and a callable.
type Consumer struct {
	Callable Callable.Function `gd:"callable"`
	Flags    Flags             `gd:"flags"`
}

type Flags int

const (
	Deferred Flags = 1 << iota // the consumer will be notifed asynchronously.
	Persist                    // the consumer will be serialized.
	OneShot                    // the consumer will be removed after the first signal.
	Weak                       // the consumer can be garbage collected.
)

// Attach connects this signal to the specified [Callable.Function]
// A signal can only be connected once to the same [Callable.Function].
// If the signal is already connected, returns [Error.InvalidParameter]
func (signal *Any) Attach(consumer Callable.Function, flags ...Flags) error { //gd:Signal.connect
	var f Flags
	for _, f := range flags {
		f |= f
	}
	if signal.proxy == nil {
		signal.proxy = &localFirst{}
	}
	return signal.proxy.Attach(signal.state, consumer, f)
}

// Remove the specified [Callable.Function] from this signal.
func (signal *Any) Remove(consumer Callable.Function) { //gd:Signal.disconnect
	if signal.proxy != nil {
		signal.proxy.Remove(signal.state, consumer)
	}
}

// HasConnections returns true if any Callable is connected to this signal.
func (signal Any) HasConnections() bool { //gd:Signal.has_connections
	for range signal.Consumers() {
		return true
	}
	return false
}

// Emit emits this signal. All consumers to this signal will be notified.
func (signal Any) Emit(values ...variant.Any) { //gd:Signal.emit
	if signal.proxy != nil {
		signal.proxy.Emit(signal.state, values...)
	}
}

// Consumers returns a sequence on all of the connected consumers.
func (signal Any) Consumers() iter.Seq[Consumer] { //gd:Signal.get_connections
	if signal.proxy != nil {
		return signal.proxy.Consumers(signal.state)
	}
	return func(func(Consumer) bool) {}
}

// Name returns the name of this signal.
func (signal Any) Name() String.Readable { //gd:Signal.get_name
	if signal.proxy != nil {
		return signal.proxy.Name(signal.state)
	}
	return String.New()
}

// Emitter returns the originater associated with this signal.
func (signal Any) Emitter() variant.Any { //gd:Signal.get_object Signal.get_object_id
	if signal.proxy != nil {
		return signal.proxy.Emitter(signal.state)
	}
	return variant.Nil
}

// Has returns true if the specified callable is connected to this signal.
func (signal Any) Has(fn Callable.Function) bool { //gd:Signal.is_connected
	for c := range signal.Consumers() {
		if c.Callable == fn {
			return true
		}
	}
	return false
}

// IsNull returns true if this signal is not connected to any callable.
func (signal Any) IsNull() bool { //gd:Signal.is_null
	return signal == Nil
}

// Solo value that can be signaled, add this as a field inside a [classdb.Extension]
// to register it as a signal.
type Solo[A any] struct {
	Any
}

// Emit the value to all connected signal handlers. Safe to call from any goroutine.
func (signal Solo[A]) Emit(a A) {
	Callable.Defer(Callable.New(func() {
		signal.Any.Emit(variant.New(a))
	}))
}

// Pair of values that can be signaled, add this as a field inside a [classdb.Extension]
// to register it as a signal.
type Pair[A, B any] struct {
	Any
}

// Emit the pair of values to all connected signal handlers. Safe to call from any goroutine.
func (signal Pair[A, B]) Emit(a A, b B) {
	Callable.Defer(Callable.New(func() {
		signal.Any.Emit(variant.New(a), variant.New(b))
	}))
}

// Trio of values that can be signaled, add this as a field inside a [classdb.Extension]
// to register it as a signal.
type Trio[A, B, C any] struct {
	Any
}

// Emit the pair of values to all connected signal handlers. Safe to call from any goroutine.
func (signal Trio[A, B, C]) Emit(a A, b B, c C) {
	Callable.Defer(Callable.New(func() {
		signal.Any.Emit(variant.New(a), variant.New(b), variant.New(c))
	}))
}

// Quad of values that can be signaled, add this as a field inside a [classdb.Extension]
// to register it as a signal.
type Quad[A, B, C, D any] struct {
	Any
}

// Emit the pair of values to all connected signal handlers. Safe to call from any goroutine.
func (signal Quad[A, B, C, D]) Emit(a A, b B, c C, d D) {
	Callable.Defer(Callable.New(func() {
		signal.Any.Emit(variant.New(a), variant.New(b), variant.New(c), variant.New(d))
	}))
}

// Quin of values that can be signaled, add this as a field inside a [classdb.Extension]
// to register it as a signal.
type Quin[A, B, C, D, E any] struct {
	Any
}

// Emit the pair of values to all connected signal handlers. Safe to call from any goroutine.
// This function is safe to call from any goroutine.
func (signal Quin[A, B, C, D, E]) Emit(a A, b B, c C, d D, e E) {
	Callable.Defer(Callable.New(func() {
		signal.Any.Emit(variant.New(a), variant.New(b), variant.New(c), variant.New(d), variant.New(e))
	}))
}

// Hexa of values that can be signaled, add this as a field inside a [classdb.Extension]
// to register it as a signal.
type Hexa[A, B, C, D, E, F any] struct {
	Any
}

// Emit the pair of values to all connected signal handlers. Safe to call from any goroutine.
// This function is safe to call from any goroutine.
func (signal Hexa[A, B, C, D, E, F]) Emit(a A, b B, c C, d D, e E, f F) {
	Callable.Defer(Callable.New(func() {
		signal.Any.Emit(variant.New(a), variant.New(b), variant.New(c), variant.New(d), variant.New(e), variant.New(f))
	}))
}
