package Array

import "graphics.gd/variant"

// Proxy can be implemented to provide a foreign-managed array representation. This can be useful
// when you want to access an array with its implementation hosted in another programming language.
type Proxy[T any] interface {
	Any() Any
	Resize(int)
	Index(int) T
	SetIndex(int, T)
	Len() int
	IsReadOnly() bool
	MakeReadOnly()
}

// State is used as user data when working with a foreign-managed array so that the implementation
// does not require any Go-managed allocations for a new array.
type State [16]byte

// Through returns a new array that accesses the underlying data of the array through the given
// [Proxy].
func Through[T any](proxy Proxy[T], state State) Contains[T] {
	return Contains[T]{proxy: proxy}
}

// As converts the array into a foreign array representation, by reconstructing the array via the
// available proxy methods. Panics if the array is already being proxied through a different
// implementation, as reference semantics would no longer be preserved. The allocation function is
// required so that a new proxy can be constructed if necessary, otherwise the existing proxy and
// state is returned.
func As[P Proxy[T], T any](array Contains[T], alloc func() (P, State)) (P, State) {
	existing, ok := array.proxy.(P)
	if ok {
		return existing, array.state
	}
	if array.proxy != nil {
		panic("array is already proxied")
	}
	proxy, state := alloc()
	proxy.Resize(array.Len())
	for i, v := range array.Iter() {
		proxy.SetIndex(i, v)
	}
	return proxy, state
}

// arrays are always backed by a local Go slice by default but can be proxied on-demand to a foreign
// array representation.
type localFirst[T any] struct {
	slice []T
	state State
	proxy Proxy[T]
}

func (p *localFirst[T]) Any() Any {
	if p.proxy != nil {
		return p.proxy.Any()
	}
	return Any{proxy: anyView{localFirst: p}}
}
func (p *localFirst[T]) Index(i int) T {
	if p.proxy != nil {
		return p.proxy.Index(i)
	}
	return p.slice[i%len(p.slice)]
}
func (p *localFirst[T]) IndexAny(i int) variant.Any { return variant.New(p.Index(i)) }
func (p *localFirst[T]) SetIndex(i int, v T) {
	if p.proxy != nil {
		p.proxy.SetIndex(i, v)
		return
	}
	if p.state[0] == 1 {
		panic("array is read-only")
	}
	p.slice[i%len(p.slice)] = v
}
func (p *localFirst[T]) SetIndexAny(i int, v variant.Any) { p.SetIndex(i, v.Interface().(T)) }
func (p *localFirst[T]) Len() int {
	if p.proxy != nil {
		return p.proxy.Len()
	}
	return len(p.slice)
}
func (p *localFirst[T]) Resize(size int) {
	if p.proxy != nil {
		p.proxy.Resize(size)
		return
	}
	if size < len(p.slice) {
		p.slice = p.slice[:size]
	} else {
		p.slice = append(p.slice, make([]T, size-len(p.slice))...)
	}
}
func (p *localFirst[T]) IsReadOnly() bool {
	if p.proxy != nil {
		return p.proxy.IsReadOnly()
	}
	return p.state[0] == 1
}
func (p *localFirst[T]) MakeReadOnly() {
	if p.proxy != nil {
		p.proxy.MakeReadOnly()
		return
	}
	p.state[0] = 1
}

type anyView struct {
	localFirst interface {
		IndexAny(int) variant.Any
		SetIndexAny(int, variant.Any)
		Len() int
		Resize(int)
		IsReadOnly() bool
		MakeReadOnly()
	}
}

func (a anyView) Any() Any                      { return Any{proxy: a} }
func (a anyView) Index(i int) variant.Any       { return a.localFirst.IndexAny(i) }
func (a anyView) SetIndex(i int, v variant.Any) { a.localFirst.SetIndexAny(i, v) }
func (a anyView) Len() int                      { return a.localFirst.Len() }
func (a anyView) Resize(size int)               { a.localFirst.Resize(size) }
func (a anyView) IsReadOnly() bool              { return a.localFirst.IsReadOnly() }
func (a anyView) MakeReadOnly()                 { a.localFirst.MakeReadOnly() }
