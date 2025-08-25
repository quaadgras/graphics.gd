package classdb

import "graphics.gd/variant"

// Trampoline can be used to efficiently call methods on a class without producing any allocations.
type Trampoline[T any] struct{}

// MakeTrampoline can be used to create a trampoline for a method. T should be a func type matching
// the signature of the method being optimized.
func MakeTrampoline[T, M any](jump func(instance *T, method M, args variant.Arguments)) Trampoline[T] {
	return Trampoline[T]{}
}
