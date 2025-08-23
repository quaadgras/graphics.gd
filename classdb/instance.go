package classdb

import (
	"reflect"
	"strings"

	gd "graphics.gd/internal"
)

func nameOf(rtype reflect.Type) string {
	if rtype.Kind() == reflect.Ptr || rtype.Kind() == reflect.Array {
		return nameOf(rtype.Elem())
	}
	isClass := reflect.PointerTo(rtype).Implements(reflect.TypeFor[gd.IsClass]()) || rtype.Implements(reflect.TypeFor[gd.IsClass]())
	if rtype.Kind() == reflect.Struct && rtype.NumField() > 0 && isClass {
		if rtype.Field(0).Anonymous {
			if rename, ok := rtype.Field(0).Tag.Lookup("gd"); ok {
				return rename
			}
			if rtype.Name() == "" {
				return nameOf(rtype.Field(0).Type)
			}
		}
		return strings.TrimPrefix(rtype.Name(), "class")
	}
	return ""
}
