package threadsafe

import "sync"

type Queue[T any] struct {
	mutex sync.Mutex
	slice []T
}

func (q *Queue[T]) Push(v T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.slice = append(q.slice, v)
}

func (q *Queue[T]) Pop() (T, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if len(q.slice) == 0 {
		var zero T
		return zero, false
	}
	value := q.slice[len(q.slice)-1]
	q.slice = q.slice[:len(q.slice)-1]
	return value, true
}
