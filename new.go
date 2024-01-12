//go:build !generate

package gd

func New[T any](owner any) T {
	var zero T
	return zero
}
