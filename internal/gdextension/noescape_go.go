//go:build !cgo && !wasm

package gdextension

func Call[T any](object Object, method MethodForClass, shape Shape, args any) T {
	panic("not implemented")
}

func (method MethodForClass) Call(self Object, args ...Variant) (Variant, error) {
	panic("not implemented")
}
