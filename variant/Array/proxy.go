package Array

import (
	"graphics.gd/variant"
)

// Proxy can be implemented to provide a foreign-managed array representation. This can be useful
// when you want to access an array with its implementation hosted in another programming language.
type Proxy[T any] interface {
	Any(State) Any
	Resize(State, int)
	Index(State, int) T
	SetIndex(State, int, T)
	Len(State) int
	IsReadOnly(State) bool
	MakeReadOnly(State)
}

// State is used as user data when working with a foreign-managed array so that the implementation
// does not require any Go-managed allocations for a new array.
type State [16]byte

// Through returns a new array that accesses the underlying data of the array through the given
// [Proxy].
func Through[T any](proxy Proxy[T], state State) Contains[T] {
	return Contains[T]{proxy: proxy, state: state}
}

// As converts the array into a foreign array representation, by reconstructing the array via the
// available proxy methods. Panics if the array is already being proxied through a different
// implementation, as reference semantics would no longer be preserved. The allocation function is
// required so that a new proxy can be constructed if necessary, otherwise the existing proxy and
// state is returned.
func As[P Proxy[T], T any](array Contains[T], alloc func() (P, State)) (P, State) {
	if array.proxy == nil {
		return alloc()
	}
	existing, ok := array.proxy.(P)
	if ok {
		return existing, array.state
	}
	local, ok := array.proxy.(*localFirst[T])
	if !ok {
		view, ok := any(array.proxy).(anyView)
		if ok {
			proxy, state := alloc()
			proxy.Resize(state, view.Len(array.state))
			for i := 0; i < view.Len(array.state); i++ {
				proxy.SetIndex(state, i, any(view.Index(array.state, i)).(T))
			}
			view.pass(any(proxy).(Proxy[variant.Any]), state)
			return proxy, state
		}
		panic("array is already proxied")
	}
	proxy, state := alloc()
	proxy.Resize(state, local.Len(array.state))
	for i := 0; i < local.Len(array.state); i++ {
		proxy.SetIndex(state, i, local.Index(array.state, i))
	}
	local.slice = nil
	local.proxy = proxy
	local.state = state
	return proxy, state
}

// arrays are always backed by a local Go slice by default but can be proxied on-demand to a foreign
// array representation.
type localFirst[T any] struct {
	slice []T
	state State
	proxy Proxy[T]
}

func (p *localFirst[T]) Any(State) Any {
	if p.proxy != nil {
		return p.proxy.Any(p.state)
	}
	return Any{proxy: anyView{localFirst: p}}
}
func (p *localFirst[T]) Index(_ State, i int) T {
	if p.proxy != nil {
		return p.proxy.Index(p.state, i)
	}
	return p.slice[i%len(p.slice)]
}
func (p *localFirst[T]) IndexAny(i int) variant.Any { return variant.New(p.Index(p.state, i)) }
func (p *localFirst[T]) SetIndex(_ State, i int, v T) {
	if p.proxy != nil {
		p.proxy.SetIndex(p.state, i, v)
		return
	}
	if p.state[0] == 1 {
		panic("array is read-only")
	}
	p.slice[i%len(p.slice)] = v
}
func (p *localFirst[T]) SetIndexAny(i int, v variant.Any) { p.SetIndex(p.state, i, v.Interface().(T)) }
func (p *localFirst[T]) Len(State) int {
	if p.proxy != nil {
		return p.proxy.Len(p.state)
	}
	return len(p.slice)
}
func (p *localFirst[T]) Resize(_ State, size int) {
	if p.proxy != nil {
		p.proxy.Resize(p.state, size)
		return
	}
	if size < len(p.slice) {
		p.slice = p.slice[:size]
	} else {
		p.slice = append(p.slice, make([]T, size-len(p.slice))...)
	}
}
func (p *localFirst[T]) IsReadOnly(State) bool {
	if p.proxy != nil {
		return p.proxy.IsReadOnly(p.state)
	}
	return p.state[0] == 1
}
func (p *localFirst[T]) MakeReadOnly(State) {
	if p.proxy != nil {
		p.proxy.MakeReadOnly(p.state)
		return
	}
	p.state[0] = 1
}
func (p *localFirst[T]) pass(proxy Proxy[variant.Any], state State) {
	if p.proxy != nil {
		panic("array is already proxied")
	}
	p.proxy = typedView[T]{proxy}
	p.state = state
	p.slice = nil
}

type typedView[T any] struct {
	Proxy[variant.Any]
}

func (t typedView[T]) Index(state State, i int) T { return variant.As[T](t.Proxy.Index(state, i)) }
func (t typedView[T]) SetIndex(state State, i int, v T) {
	t.Proxy.SetIndex(state, i, variant.New(v))
}

func (array *Contains[T]) SetAny(a Any) {
	array.proxy = typedView[T]{a.proxy}
	array.state = a.state
}

type anyView struct {
	localFirst interface {
		IndexAny(int) variant.Any
		SetIndexAny(int, variant.Any)
		Len(State) int
		Resize(State, int)
		IsReadOnly(State) bool
		MakeReadOnly(State)
		pass(Proxy[variant.Any], State)
	}
}

func (a anyView) Any(State) Any                              { return Any{proxy: a} }
func (a anyView) Index(_ State, i int) variant.Any           { return a.localFirst.IndexAny(i) }
func (a anyView) SetIndex(_ State, i int, v variant.Any)     { a.localFirst.SetIndexAny(i, v) }
func (a anyView) Len(state State) int                        { return a.localFirst.Len(state) }
func (a anyView) Resize(state State, size int)               { a.localFirst.Resize(state, size) }
func (a anyView) IsReadOnly(state State) bool                { return a.localFirst.IsReadOnly(state) }
func (a anyView) MakeReadOnly(state State)                   { a.localFirst.MakeReadOnly(state) }
func (a anyView) pass(proxy Proxy[variant.Any], state State) { a.localFirst.pass(proxy, state) }
