package startup

import (
	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
)

var unpins []func()

// Pin the given value to the lifetime of the program, preventing the underlying pointer from
// expiring, can only be called during initialization.
func Pin[T any](val T) T {
	switch val := any(val).(type) {
	case gd.IsClass:
		obj := val.AsObject()
		pointers.Pin(obj[0])
		unpins = append(unpins, func() {
			obj[0].Free()
		})
	}
	return val
}
