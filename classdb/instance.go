package classdb

import (
	"reflect"
	"strings"
	"unsafe"

	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
)

// ID that uniquely identifies an object.
type ID gd.ObjectID

// IsValid returns true if the Object that corresponds to id is a valid object
// (e.g. has not been deleted from memory). All Objects have a unique instance ID.
func (id ID) IsValid() bool { //gd:is_instance_id_valid is_instance_valid
	return bool(gd.IsInstanceIdValid(gd.Int(id)))
}

// Get returns the Object that corresponds to id. All Objects have a unique instance ID.
func Get(id ID) gd.Object { //gd:instance_from_id
	return gd.InstanceFromId(gd.Int(id))
}

// As attempts to cast the given class to T, returning true
// if the cast was successful.
func As[T gd.IsClass](value gd.IsClass) (T, bool) {
	ext, ok := gd.ExtensionInstances.Load(pointers.Get(value.AsObject()[0])[0])
	if ok {
		if ref, ok := ext.(T); ok {
			pointers.Lay(value.AsObject()[0])
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
	if casted != ([1]gd.Object{}) && pointers.Get(casted[0]) != ([3]uint64{}) {
		return (*(*T)(unsafe.Pointer(&casted))), true
	}
	return zero, false
}

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
