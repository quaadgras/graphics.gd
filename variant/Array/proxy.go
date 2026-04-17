package Array

import (
	"graphics.gd/variant"
)

type Implementation[T any] interface {
	AsArrayAny() Any
	Resize(int)
	Index(int) T
	SetIndex(int, T)
	Len() int
	IsReadOnly() bool
	MakeReadOnly()
	Free()
}

// As converts the array into a foreign array representation, by reconstructing the array via the
// available proxy methods. Panics if the array is already being proxied through a different
// implementation, as reference semantics would no longer be preserved. The allocation function is
// required so that a new proxy can be constructed if necessary, otherwise the existing proxy is used.
func As[P Implementation[T], T any](array Contains[T], alloc func() P) P {
	if array.Implementation == nil {
		return alloc()
	}
	existing, ok := array.Implementation.(P)
	if ok {
		return existing
	}
	local, ok := array.Implementation.(*localFirst[T])
	if !ok {
		view, ok := any(array.Implementation).(anyView)
		if ok {
			proxy := alloc()
			proxy.Resize(view.Len())
			for i := 0; i < view.Len(); i++ {
				proxy.SetIndex(i, variant.As[T](view.Index(i)))
			}
			if view.IsReadOnly() {
				proxy.MakeReadOnly()
			}
			view.pass(any(proxy).(Implementation[variant.Any]))
			return proxy
		}
		panic("array is already proxied")
	}
	proxy := alloc()
	proxy.Resize(local.Len())
	for i := 0; i < local.Len(); i++ {
		proxy.SetIndex(i, local.Index(i))
	}
	if local.IsReadOnly() {
		proxy.MakeReadOnly()
	}
	return proxy
}

// arrays are always backed by a local Go slice by default but can be proxied on-demand to a foreign
// array representation.
type localFirst[T any] struct {
	slice []T
	write bool
	proxy Implementation[T]
}

func (p *localFirst[T]) Free() {
	if p.proxy != nil {
		p.proxy.Free()
	}
}

func (p *localFirst[T]) AsArrayAny() Any {
	if p.proxy != nil {
		return p.proxy.AsArrayAny()
	}
	return Any{Implementation: anyView{localFirst: p}}
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
	if !p.write {
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
func (p *localFirst[T]) IsReadOnly() bool { return !p.write }
func (p *localFirst[T]) MakeReadOnly()    { p.write = false }

func (p *localFirst[T]) pass(proxy Implementation[variant.Any]) {
	if p.proxy != nil {
		panic("array is already proxied")
	}
	p.proxy = typedView[T]{proxy}
	p.slice = nil
}

type typedView[T any] struct {
	Implementation[variant.Any]
}

func (t typedView[T]) Index(i int) T { return variant.As[T](t.Implementation.Index(i)) }
func (t typedView[T]) SetIndex(i int, v T) {
	t.Implementation.SetIndex(i, variant.New(v))
}

func (array *Contains[T]) SetAny(a Any) {
	array.Implementation = typedView[T]{a.Implementation}
}

type anyView struct {
	localFirst interface {
		IndexAny(int) variant.Any
		SetIndexAny(int, variant.Any)
		Len() int
		Resize(int)
		IsReadOnly() bool
		MakeReadOnly()
		Free()
		pass(Implementation[variant.Any])
	}
}

func (a anyView) AsArrayAny() Any                        { return Any{Implementation: a} }
func (a anyView) Index(i int) variant.Any                { return a.localFirst.IndexAny(i) }
func (a anyView) SetIndex(i int, v variant.Any)          { a.localFirst.SetIndexAny(i, v) }
func (a anyView) Len() int                               { return a.localFirst.Len() }
func (a anyView) Resize(size int)                        { a.localFirst.Resize(size) }
func (a anyView) IsReadOnly() bool                       { return a.localFirst.IsReadOnly() }
func (a anyView) MakeReadOnly()                          { a.localFirst.MakeReadOnly() }
func (a anyView) pass(proxy Implementation[variant.Any]) { a.localFirst.pass(proxy) }
func (a anyView) Free()                                  { a.localFirst.Free() }
