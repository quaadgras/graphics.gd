package defined

import (
	gd "graphics.gd/internal"
)

type ExtensionTo[T gd.IsClass] interface {
	Class
	super() T
}
