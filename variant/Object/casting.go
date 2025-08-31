package Object

import (
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
)

// To attempts to cast the given class to T, returning the
// casted value if successful, or panicking if not.
func To[T gd.IsClass](value gd.IsClass) T {
	casted, ok := As[T](value)
	if !ok {
		panic("Object is not of the expected type")
	}
	return casted
}

// Is returns true if the given class is of type T.
func Is[T gd.IsClass](value gd.IsClass) bool {
	_, ok := As[T](value)
	return ok
}

// As attempts to cast the given class to T, returning true
// if the cast was successful.
func As[T gd.IsClass](value gd.IsClass) (T, bool) {
	if value.AsObject() == ([1]gd.Object{}) || value == nil {
		var zero T
		return zero, false
	}
	ext, ok := gd.ExtensionInstanceLookup(gdextension.Object(pointers.Get(value.AsObject()[0])[0])).(T)
	if ok {
		pointers.Lay(value.AsObject()[0])
		return ext, true
	}
	var zero T
	castable, ok := any(&zero).(gd.IsClassCastable)
	if ok {
		return zero, castable.SetObject(value.AsObject())
	}
	return zero, false
}
