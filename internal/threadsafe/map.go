package threadsafe

import (
	"iter"
	"sync"
)

type Map[K comparable, V any] sync.Map

func (m *Map[K, V]) Lookup(key K) (value V, ok bool) {
	v, ok := (*sync.Map)(m).Load(key)
	if !ok {
		var zero V
		return zero, false
	}
	return v.(V), true
}

func (m *Map[K, V]) Insert(key K, value V) {
	(*sync.Map)(m).Store(key, value)
}

func (m *Map[K, V]) Remove(key K) {
	(*sync.Map)(m).Delete(key)
}

func (m *Map[K, V]) Iter() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		(*sync.Map)(m).Range(func(key, value any) bool {
			return yield(key.(K), value.(V))
		})
	}
}
