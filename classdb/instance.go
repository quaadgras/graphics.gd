package classdb

import (
	"reflect"
	"strings"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/threadsafe"
)

var instances threadsafe.Handles[*instanceImplementation, gdextension.ExtensionInstanceID]

func nameOf(rtype reflect.Type) string {
	if rtype.Kind() == reflect.Array {
		return rtype.Elem().Name()
	}
	if rtype.Kind() == reflect.Pointer {
		return nameOf(rtype.Elem())
	}
	isClass := reflect.PointerTo(rtype).Implements(reflect.TypeFor[gd.IsClass]()) || rtype.Implements(reflect.TypeFor[gd.IsClass]())
	if rtype.Kind() == reflect.Struct && rtype.NumField() > 0 && isClass {
		if rtype.Field(0).Anonymous {
			if rename, ok := rtype.Field(0).Tag.Lookup("gd"); ok {
				return rename
			}
			if rtype.Name() == "" || !rtype.Implements(reflect.TypeFor[gdclass.Interface]()) {
				return nameOf(rtype.Field(0).Type)
			}
		}
		return strings.TrimPrefix(rtype.Name(), "class")
	}
	return ""
}
