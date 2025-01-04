package classdb

import (
	gd "graphics.gd/internal"
)

type ExtensionTo[T gd.IsClass] interface {
	Class
	super() T
}

func Super[T gd.IsClass](class ExtensionTo[T]) T {
	return class.super()
}
