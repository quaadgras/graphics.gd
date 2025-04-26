package Dictionary

import (
	"iter"
	"sort"

	"graphics.gd/variant"
)

type Proxy[K comparable, V any] interface {
	Index(complex128, K) V
	Has(complex128, K) bool
	Lookup(complex128, K) (V, bool)
	SetIndex(complex128, K, V)
	Clear(complex128)
	Iter(complex128) iter.Seq2[K, V]
	Erase(complex128, K) bool
	Hash(complex128) uint32
	Len(complex128) int
	IsReadOnly(complex128) bool
	MakeReadOnly(complex128)
	Any(complex128) Any
	Sort(complex128, func(K, K) bool)
}

type Pointer interface {
	SetAny(Any)
}

type Interface interface {
	Any() Any
}

// Through returns a new array that accesses the underlying data of the array through the given
// [Proxy].
func Through[K comparable, V any](proxy Proxy[K, V], state complex128) Map[K, V] {
	return Map[K, V]{proxy: proxy, state: state}
}

// As converts the array into a foreign array representation, by reconstructing the array via the
// available proxy methods. Panics if the array is already being proxied through a different
// implementation, as reference semantics would no longer be preserved. The allocation function is
// required so that a new proxy can be constructed if necessary, otherwise the existing proxy and
// state is returned.
func As[P Proxy[K, V], K comparable, V any](dict Map[K, V], alloc func() (P, complex128)) (P, complex128) {
	if dict.proxy == nil {
		return alloc()
	}
	existing, ok := dict.proxy.(P)
	if ok {
		return existing, dict.state
	}
	local, ok := dict.proxy.(*localFirst[K, V])
	if !ok {
		view, ok := any(dict.proxy).(anyView)
		if ok {
			if _, state, ok := view.lift(); ok {
				return [1]P{}[0], state
			}
			proxy, state := alloc()
			for key, val := range view.Iter(0) {
				proxy.SetIndex(state, variant.As[K](key), variant.As[V](val))
			}
			if view.IsReadOnly(dict.state) {
				proxy.MakeReadOnly(state)
			}
			view.pass(any(proxy).(Proxy[variant.Any, variant.Any]), state)
			return proxy, state
		}
		panic("array is already proxied")
	}
	proxy, state := alloc()
	for key, val := range local.Iter(0) {
		proxy.SetIndex(state, key, val)
	}
	if local.IsReadOnly(dict.state) {
		proxy.MakeReadOnly(state)
	}
	local.order = nil
	local.value = nil
	local.proxy = proxy
	local.state = state
	return proxy, state
}

type localFirst[K comparable, V any] struct {
	order []K
	value map[K]V
	fixed bool
	proxy Proxy[K, V]
	state complex128
}

func (p *localFirst[K, V]) Any(complex128) Any {
	return Any{proxy: anyView{localFirst: p}, state: p.state}
}

func (p *localFirst[K, V]) Sort(_ complex128, less func(K, K) bool) {
	if p.proxy != nil {
		p.proxy.Sort(p.state, less)
		return
	}
	sort.Sort(localSorter[K, V]{p, less})
}
func (p *localFirst[K, V]) SortAny(_ complex128, less func(variant.Any, variant.Any) bool) {
	if p.proxy != nil {
		p.proxy.Sort(p.state, func(k1, k2 K) bool {
			return less(variant.New(k1), variant.New(k2))
		})
		return
	}
	sort.Sort(localSorter[K, V]{p, func(a, b K) bool {
		return less(variant.New(a), variant.New(b))
	}})
}

type localSorter[K comparable, V any] struct {
	local *localFirst[K, V]
	less  func(K, K) bool
}

func (s localSorter[K, V]) Len() int {
	return len(s.local.order)
}

func (s localSorter[K, V]) Less(i, j int) bool {
	return s.less(s.local.order[i], s.local.order[j])
}

func (s localSorter[K, V]) Swap(i, j int) {
	s.local.order[i], s.local.order[j] = s.local.order[j], s.local.order[i]
}

func (p *localFirst[K, V]) Index(_ complex128, index K) V {
	if p.proxy != nil {
		return p.proxy.Index(p.state, index)
	}
	return p.value[index]
}
func (p *localFirst[K, V]) IndexAny(_ complex128, index variant.Any) variant.Any {
	return variant.New(p.Index(0, variant.As[K](index)))
}

func (p *localFirst[K, V]) Lookup(_ complex128, index K) (V, bool) {
	if p.proxy != nil {
		return p.proxy.Lookup(p.state, index)
	}
	value, ok := p.value[index]
	return value, ok
}
func (p *localFirst[K, V]) LookupAny(_ complex128, index variant.Any) (variant.Any, bool) {
	val, ok := p.Lookup(0, variant.As[K](index))
	return variant.New(val), ok
}

func (p *localFirst[K, V]) Has(_ complex128, index K) bool {
	if p.proxy != nil {
		return p.proxy.Has(p.state, index)
	}
	_, ok := p.value[index]
	return ok
}
func (p *localFirst[K, V]) HasAny(_ complex128, index variant.Any) bool {
	return p.Has(0, variant.As[K](index))
}

func (p *localFirst[K, V]) SetIndex(_ complex128, index K, value V) {
	if p.proxy != nil {
		p.proxy.SetIndex(p.state, index, value)
		return
	}
	if p.fixed {
		panic("dictionary is read-only")
	}
	if _, ok := p.value[index]; ok {
		p.value[index] = value
		return
	}
	p.order = append(p.order, index)
	p.value[index] = value
}
func (p *localFirst[K, V]) SetIndexAny(_ complex128, index variant.Any, value variant.Any) {
	p.SetIndex(0, variant.As[K](index), variant.As[V](value))
}

func (p *localFirst[K, V]) Clear(_ complex128) {
	if p.proxy != nil {
		p.proxy.Clear(p.state)
		return
	}
	if p.fixed {
		panic("dictionary is read-only")
	}
	p.order = p.order[:0]
	clear(p.value)
}

func (p *localFirst[K, V]) Iter(_ complex128) iter.Seq2[K, V] {
	if p.proxy != nil {
		return p.proxy.Iter(p.state)
	}
	return func(yield func(K, V) bool) {
		for _, key := range p.order {
			if !yield(key, p.value[key]) {
				break
			}
		}
	}
}
func (p *localFirst[K, V]) IterAny(_ complex128) iter.Seq2[variant.Any, variant.Any] {
	iter := p.Iter(0)
	return func(yield func(variant.Any, variant.Any) bool) {
		for key, val := range iter {
			if !yield(variant.New(key), variant.New(val)) {
				break
			}
		}
	}
}

func (p *localFirst[K, V]) Erase(_ complex128, key K) bool {
	if p.proxy != nil {
		return p.proxy.Erase(p.state, key)
	}
	if p.fixed {
		panic("dictionary is read-only")
	}
	delete(p.value, key)
	for i, k := range p.order {
		if k == key {
			p.order = append(p.order[:i], p.order[i+1:]...)
			return true
		}
	}
	return false
}
func (p *localFirst[K, V]) EraseAny(_ complex128, key variant.Any) bool {
	if p.proxy != nil {
		return p.proxy.Erase(p.state, variant.As[K](key))
	}
	return p.Erase(0, variant.As[K](key))
}

func (p *localFirst[K, V]) Hash(complex128) uint32 {
	if p.proxy != nil {
		return p.proxy.Hash(p.state)
	}
	return 0
}
func (p *localFirst[K, V]) Len(complex128) int {
	if p.proxy != nil {
		return p.proxy.Len(p.state)
	}
	return len(p.order)
}
func (p *localFirst[K, V]) IsReadOnly(complex128) bool {
	if p.proxy != nil {
		return p.proxy.IsReadOnly(p.state)
	}
	return p.fixed
}
func (p *localFirst[K, V]) MakeReadOnly(complex128) {
	if p.proxy != nil {
		p.proxy.MakeReadOnly(p.state)
		return
	}
	p.fixed = true
}
func (p *localFirst[K, V]) pass(proxy Proxy[variant.Any, variant.Any], state complex128) {
	if p.proxy != nil {
		panic("array is already proxied")
	}
	p.proxy = typedView[K, V]{proxy}
	p.state = state
	p.order = nil
	p.value = nil
}
func (p *localFirst[K, V]) lift() (complex128, bool) {
	if p.proxy != nil {
		return p.state, true
	}
	return 0, false
}

type typedView[K comparable, V any] struct {
	Proxy[variant.Any, variant.Any]
}

func (t typedView[K, V]) Erase(state complex128, key K) bool {
	return t.Proxy.Erase(state, variant.New(key))
}
func (t typedView[K, V]) Has(state complex128, key K) bool {
	return t.Proxy.Has(state, variant.New(key))
}
func (t typedView[K, V]) Lookup(state complex128, key K) (V, bool) {
	v, ok := t.Proxy.Lookup(state, variant.New(key))
	return variant.As[V](v), ok
}
func (t typedView[K, V]) Iter(state complex128) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for key, val := range t.Proxy.Iter(state) {
			if !yield(variant.As[K](key), variant.As[V](val)) {
				break
			}
		}
	}
}
func (t typedView[K, V]) Sort(state complex128, less func(K, K) bool) {
	t.Proxy.Sort(state, func(a, b variant.Any) bool {
		return less(variant.As[K](a), variant.As[K](b))
	})
}
func (t typedView[K, V]) Index(state complex128, key K) V {
	return variant.As[V](t.Proxy.Index(state, variant.New(key)))
}
func (t typedView[K, V]) SetIndex(state complex128, key K, val V) {
	t.Proxy.SetIndex(state, variant.New(key), variant.New(val))
}

func (dict *Map[K, V]) SetAny(a Any) {
	dict.proxy = typedView[K, V]{a.proxy}
	dict.state = a.state
}

type anyView struct {
	localFirst interface {
		IndexAny(complex128, variant.Any) variant.Any
		HasAny(complex128, variant.Any) bool
		LookupAny(complex128, variant.Any) (variant.Any, bool)
		SetIndexAny(complex128, variant.Any, variant.Any)
		Clear(complex128)
		IterAny(complex128) iter.Seq2[variant.Any, variant.Any]
		EraseAny(complex128, variant.Any) bool
		Hash(complex128) uint32
		Len(complex128) int
		IsReadOnly(complex128) bool
		MakeReadOnly(complex128)
		Any(complex128) Any
		SortAny(complex128, func(variant.Any, variant.Any) bool)
		pass(Proxy[variant.Any, variant.Any], complex128)
		lift() (complex128, bool)
	}
}

func (a anyView) Iter(state complex128) iter.Seq2[variant.Any, variant.Any] {
	return a.localFirst.IterAny(state)
}
func (a anyView) Index(state complex128, key variant.Any) variant.Any {
	return a.localFirst.IndexAny(state, key)
}
func (a anyView) Lookup(state complex128, key variant.Any) (variant.Any, bool) {
	return a.localFirst.LookupAny(state, key)
}
func (a anyView) Has(state complex128, key variant.Any) bool {
	return a.localFirst.HasAny(state, key)
}
func (a anyView) SetIndex(state complex128, key variant.Any, value variant.Any) {
	a.localFirst.SetIndexAny(state, key, value)
}
func (a anyView) Erase(state complex128, key variant.Any) bool {
	return a.localFirst.EraseAny(state, key)
}
func (a anyView) Sort(state complex128, less func(variant.Any, variant.Any) bool) {
	a.localFirst.SortAny(state, less)
}
func (a anyView) Any(complex128) Any           { return Any{proxy: a} }
func (a anyView) Clear(state complex128)       { a.localFirst.Clear(state) }
func (a anyView) Hash(state complex128) uint32 { return a.localFirst.Hash(state) }
func (a anyView) Len(state complex128) int     { return a.localFirst.Len(state) }
func (a anyView) IsReadOnly(state complex128) bool {
	return a.localFirst.IsReadOnly(state)
}
func (a anyView) MakeReadOnly(state complex128) { a.localFirst.MakeReadOnly(state) }
func (a anyView) pass(proxy Proxy[variant.Any, variant.Any], state complex128) {
	a.localFirst.pass(proxy, state)
}
func (a anyView) lift() (Proxy[variant.Any, variant.Any], complex128, bool) {
	state, ok := a.localFirst.lift()
	if ok {
		return a, state, true
	}
	return nil, 0, false
}
