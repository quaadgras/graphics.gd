package threadsafe

import (
	"iter"
	"sync/atomic"
)

type Slice[T any] struct {
	data atomic.Pointer[[]T]
}

func (s *Slice[T]) Append(value T) int {
	for {
		oldp := s.data.Load()
		var old []T
		if oldp != nil {
			old = *oldp
		}
		next := make([]T, len(old)+1)
		copy(next, old)
		next[len(old)] = value
		if s.data.CompareAndSwap(oldp, &next) {
			return len(next) - 1
		}
	}
}

func (s *Slice[T]) Index(index int) T {
	return (*s.data.Load())[index]
}

func (s *Slice[T]) SetIndex(index int, value T) {
	for {
		oldp := s.data.Load()
		old := *oldp
		next := make([]T, len(old))
		copy(next, old)
		next[index] = value
		if s.data.CompareAndSwap(oldp, &next) {
			return
		}
	}
}

func (s *Slice[T]) Values() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		p := s.data.Load()
		if p == nil {
			return
		}
		for i, v := range *p {
			if !yield(i, v) {
				return
			}
		}
	}
}
