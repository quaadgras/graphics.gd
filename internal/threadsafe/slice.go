package threadsafe

import (
	"iter"
	"sync"
)

type Slice[T any] struct {
	mutex sync.RWMutex
	slice []T
}

func (slice *Slice[T]) Append(value T) int {
	slice.mutex.Lock()
	defer slice.mutex.Unlock()
	slice.slice = append(slice.slice, value)
	return len(slice.slice) - 1
}

func (slice *Slice[T]) Index(index int) T {
	slice.mutex.RLock()
	defer slice.mutex.RUnlock()
	return slice.slice[index]
}

func (slice *Slice[T]) SetIndex(index int, value T) {
	slice.mutex.Lock()
	defer slice.mutex.Unlock()
	slice.slice[index] = value
}

func (slice *Slice[T]) Values() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		slice.mutex.RLock()
		defer slice.mutex.RUnlock()
		for i, v := range slice.slice {
			if !yield(i, v) {
				return
			}
		}
	}
}
