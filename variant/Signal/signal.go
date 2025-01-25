// Package Signal provides a type representing an event stream or pub/sub message queue.
package Signal

import (
	"sync"

	"graphics.gd/classdb/Engine"
	gd "graphics.gd/internal"
	"graphics.gd/variant/Callable"
)

// Chan is a multi-producer, multi-consumer channel, when T is a value, it denotes the
// value being sent on the channel, when T is a function, the results of calling that
// function are passed on the channel instead.
type Chan[T any] struct {
	_ [0]sync.Mutex //nocopy

	owner any
	topic string
	funcs []Callable.Function
	proxy *gd.Signal
}

type Any = gd.Signal

const ErrInvalidParameter = Engine.ErrInvalidParameter

// Attach connects this signal to the specified [Callable.Func]
// A signal can only be connected once to the same [Callable.Func].
// If the signal is already connected, returns [ErrInvalidParameter]
func (c *Chan[T]) Attach(fn Callable.Function) error { //gd:Signal.connect
	if c.proxy != nil {
		//return c.proxy.C(fn)
	}
	for _, f := range c.funcs {
		if f == fn {
			return Engine.ErrInvalidParameter
		}
	}
	c.funcs = append(c.funcs, fn)
	return nil
}

// Remove disconnects this signal from the specified [Callable.Func].
func (c *Chan[T]) Remove(fn Callable.Function) { //gd:Signal.disconnect
	if c.proxy != nil {
		//if c.proxy.Has(fn) {
		//	c.proxy.Remove(fn)
		//}
		return
	}
	for i, f := range c.funcs {
		if f == fn {
			c.funcs = append(c.funcs[:i], c.funcs[i+1:]...)
			return
		}
	}
}

// Emit emits this signal. All Callables connected to this signal will be triggered.
func (c *Chan[T]) Emit(signal T) { //gd:Signal.emit
	if c.proxy != nil {
		//c.proxy.Emit(signal)
		return
	}
	for _, f := range c.funcs {
		f.Call()
	}
}

// Callables returns a slice of callables for this signal.
func (c *Chan[T]) Callables() []Callable.Function { //gd:Signal.get_connections
	if c.proxy != nil {
		//return c.proxy.Callables()
	}
	return c.funcs
}

// Name returns the name of this signal.
func (c *Chan[T]) Name() string { //gd:Signal.get_name
	if c.proxy != nil {
		//return c.proxy.Name()
	}
	return c.topic
}

// Emitter returns the object that emits this signal.
func (c *Chan[T]) Emitter() any { //gd:Signal.get_object Signal.get_object_id
	if c.proxy != nil {
		//return c.proxy.Emitter()
	}
	return c.owner
}

// Has returns true if the specified callable is connected to this signal.
func (c *Chan[T]) Has(fn Callable.Function) bool { //gd:Signal.is_connected
	if c.proxy != nil {
		//return c.proxy.Has(fn)
	}
	for _, f := range c.funcs {
		if f == fn {
			return true
		}
	}
	return false
}

// IsNull returns true if this signal is not connected to any callable.
func (c *Chan[T]) IsNull() bool { //gd:Signal.is_null
	return c.owner == nil && c.proxy == nil && len(c.funcs) == 0
}

// Proxy converts the signal into a proxied signal and returns the proxy.
func (c *Chan[T]) Proxy(freed bool, alloc func(any, string) gd.Signal) *gd.Signal {
	if c.proxy != nil {
		//return c.proxy
	}
	if c.IsNull() {
		return nil
	}
	//proxy := alloc(c.owner, c.topic)
	//for _, f := range c.funcs {
	//	proxy.Attach(f)
	//}
	c.funcs = nil
	//c.proxy = proxy
	//return proxy
	return nil
}

// Reset clears all connections to this signal.
func (c *Chan[T]) Reset() {
	if c.proxy != nil {
		c.proxy.Free()
	}
	*c = Chan[T]{}
}

// Solo value that can be signaled, add this as a field inside a [classdb.Extension]
// to register it as a signal.
type Solo[A any] struct {
	Any
}

// Emit the value to all connected signal handlers. Safe to call from any goroutine.
func (signal Solo[A]) Emit(a A) {
	Callable.New(func() {
		signal.Any.Emit(gd.NewVariant(a))
	})
}

// Pair of values that can be signaled, add this as a field inside a [classdb.Extension]
// to register it as a signal.
type Pair[A, B any] struct {
	Any
}

// Emit the pair of values to all connected signal handlers. Safe to call from any goroutine.
func (signal Pair[A, B]) Emit(a A, b B) {
	Callable.Defer(Callable.New(func() {
		signal.Any.Emit(gd.NewVariant(a), gd.NewVariant(b))
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
		signal.Any.Emit(gd.NewVariant(a), gd.NewVariant(b), gd.NewVariant(c))
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
		signal.Any.Emit(gd.NewVariant(a), gd.NewVariant(b), gd.NewVariant(c), gd.NewVariant(d))
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
		signal.Any.Emit(gd.NewVariant(a), gd.NewVariant(b), gd.NewVariant(c), gd.NewVariant(d), gd.NewVariant(e))
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
		signal.Any.Emit(gd.NewVariant(a), gd.NewVariant(b), gd.NewVariant(c), gd.NewVariant(d), gd.NewVariant(e), gd.NewVariant(f))
	}))
}
