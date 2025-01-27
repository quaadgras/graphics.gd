package Signal

import (
	"iter"
	"slices"
	"sync"

	"graphics.gd/variant"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Error"
	"graphics.gd/variant/String"
)

type API interface {
	Attach(complex128, Callable.Function, Flags) error
	Remove(complex128, Callable.Function)
	Name(complex128) String.Readable
	Consumers(complex128) iter.Seq[Consumer]
	Emit(complex128, ...variant.Any)
	Emitter(complex128) variant.Any
}

// Proxy can be used to transform the underlying encoding and implementation of an [Any]
// signal into a specific type. The alloc function should return a new implementation of
// the desired type, and the internal state. This may be cached in the existing implementation
// and, if so, will be passed to the check function so that it can decide whether to reuse
// the cache or not.
//
// The uint64 can be used to avoid unnecessary allocations, such that T can be zero-sized if
// the implementation does not require more than 64bits of state.
func Proxy[T API](signal Any, reuse func(T, complex128) bool, alloc func() (T, complex128)) (T, complex128) {
	if signal.proxy == nil {
		return alloc()
	}
	already, ok := signal.proxy.(T)
	if ok && reuse(already, signal.state) {
		return already, signal.state
	}
	local, ok := signal.proxy.(*localFirst)
	if !ok {
		b, state := alloc()
		for consumer := range signal.Consumers() {
			b.Attach(state, consumer.Callable, consumer.Flags)
		}
	}
	local.mutex.Lock()
	defer local.mutex.Unlock()
	b, state := alloc()
	for _, consumer := range local.consumers {
		b.Attach(state, consumer.Callable, consumer.Flags)
	}
	local.proxy = b
	local.state = state
	local.consumers = nil
	return b, state
}

// Via creates a new signal from the specified [API] implementation and state.
func Via(impl API, state complex128) Any {
	return Any{proxy: impl, state: state}
}

// localFirst is a Go allocated [Any] signal, which can convert itself into a foreign [Any] implementation
// on demand.
type localFirst struct {
	mutex     sync.RWMutex
	consumers []Consumer
	state     complex128
	proxy     API
}

func (signal *localFirst) Attach(_ complex128, fn Callable.Function, flags Flags) error {
	signal.mutex.Lock()
	defer signal.mutex.Unlock()
	if signal.proxy != nil {
		return signal.proxy.Attach(signal.state, fn, flags)
	}
	consumer := Consumer{Callable: fn, Flags: flags}
	if slices.Contains(signal.consumers, consumer) {
		return Error.InvalidParameter
	}
	signal.consumers = append(signal.consumers, consumer)
	return nil
}

func (signal *localFirst) Remove(_ complex128, fn Callable.Function) {
	signal.mutex.Lock()
	defer signal.mutex.Unlock()
	if signal.proxy != nil {
		signal.proxy.Remove(signal.state, fn)
		return
	}
	signal.consumers = slices.DeleteFunc(signal.consumers, func(consumer Consumer) bool {
		return consumer.Callable == fn
	})
}

func (signal *localFirst) Emit(_ complex128, values ...variant.Any) {
	signal.mutex.RLock()
	defer signal.mutex.RUnlock()
	if signal.proxy != nil {
		signal.proxy.Emit(signal.state, values...)
		return
	}
	var removes []Callable.Function
	for _, consumer := range signal.consumers {
		switch {
		case consumer.Flags&Deferred != 0:
			Callable.Defer(consumer.Callable, values...)
		default:
			consumer.Callable.Call(values...)
		}
		switch {
		case consumer.Flags&OneShot != 0:
			removes = append(removes, consumer.Callable)
		}
	}
}

func (signal *localFirst) Name(_ complex128) String.Readable {
	if signal.proxy != nil {
		return signal.proxy.Name(signal.state)
	}
	return String.New("anonymous")
}

func (signal *localFirst) Consumers(_ complex128) iter.Seq[Consumer] {
	signal.mutex.RLock()
	defer signal.mutex.RUnlock()
	if signal.proxy != nil {
		return signal.proxy.Consumers(signal.state)
	}
	return func(fn func(Consumer) bool) {
		for _, consumer := range signal.consumers {
			if !fn(consumer) {
				return
			}
		}
	}
}

func (signal *localFirst) Emitter(_ complex128) variant.Any {
	if signal.proxy != nil {
		return signal.proxy.Emitter(signal.state)
	}
	return variant.Nil
}
