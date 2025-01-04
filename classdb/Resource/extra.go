package Resource

import (
	"reflect"
	"strings"
	"unsafe"

	"graphics.gd/classdb/ResourceLoader"
	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
)

// ID that uniquely identifies a resource.
type ID = gd.RID

// Path to a resource, typically prefixed with "res://"
type Path string

// Load behaves like the builtin "load" function in GDScript.
func Load[T Any](path Path) T {
	resource := Instance(ResourceLoader.Load(string(path)))
	result, ok := as[T](resource)
	if !ok {
		panic("Resource is not of the expected type")
	}
	return result
}

// As attempts to cast the given class to T, returning true
// if the cast was successful.
func as[T gd.IsClass](value gd.IsClass) (T, bool) {
	ext, ok := gd.ExtensionInstances.Load(pointers.Get(value.AsObject())[0])
	if ok {
		if ref, ok := ext.(T); ok {
			return ref, true
		}
	}
	var zero T
	var rtype = reflect.TypeFor[T]()
	if rtype.Kind() == reflect.Pointer {
		return zero, false
	}
	var classtag = gd.Global.ClassDB.GetClassTag(gd.NewStringName(nameOf(rtype)))
	casted := gd.Global.Object.CastTo(value.AsObject(), classtag)
	if casted != (gd.Object{}) && pointers.Get(casted) != ([3]uintptr{}) {
		return (*(*T)(unsafe.Pointer(&casted))), true
	}
	return zero, false
}

func nameOf(rtype reflect.Type) string {
	if rtype.Kind() == reflect.Ptr || rtype.Kind() == reflect.Array {
		return nameOf(rtype.Elem())
	}
	if rtype.Kind() == reflect.Struct && rtype.NumField() > 0 {
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
