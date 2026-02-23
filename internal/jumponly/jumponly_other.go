//go:build !(go1.26 && (amd64 || arm64))

package jumponly

import (
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/noescape"
)

// PtrcallFn is unused on unsupported platforms.
var PtrcallFn uintptr

// Call falls back to noescape.Call on unsupported platforms.
func Call[T any](object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	return noescape.Call[T](object, method, shape, args)
}

// CallStatic falls back to noescape.CallStatic on unsupported platforms.
func CallStatic[T any](method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	return noescape.CallStatic[T](method, shape, args)
}
