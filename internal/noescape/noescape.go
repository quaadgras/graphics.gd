package noescape

import (
	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
)

type MethodForClass gdextension.MethodForClass
type Variant = gdunsafe.Variant

func CallStatic[T any](method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	return Call[T](0, method, shape, args)
}
