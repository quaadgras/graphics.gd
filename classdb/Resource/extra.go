package Resource

import (
	"reflect"
	"strings"
	"unsafe"

	"graphics.gd/classdb/ResourceLoader"
	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Path"
)

// ID that uniquely identifies a resource.
type ID = gd.RID

// New allocates a unique ID which can be used by the implementation to construct an RID.
// This is used mainly from native extensions to implement servers.
func AllocateID() ID { //gd:rid_allocate_id
	return gd.RidFromInt64(gd.RidAllocateId())
}

// Int64 returns a resource ID from the given int64.
func Int64(id int64) ID { //gd:rid_from_int64
	return gd.RidFromInt64(gd.Int(id))
}

// Load behaves like the builtin "load" function in GDScript.
func Load[T Any](path Path.ToResource) T {
	resource := Instance(ResourceLoader.Load(path.String()))
	result, ok := as[T](resource)
	if !ok {
		panic("Resource is not of the expected type")
	}
	return result
}

// As attempts to cast the given class to T, returning true
// if the cast was successful.
func as[T gd.IsClass](value gd.IsClass) (T, bool) {
	ext, ok := gd.ExtensionInstances.Load(pointers.Get(value.AsObject()[0])[0])
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
	if casted != ([1]gd.Object{}) && pointers.Get(casted[0]) != ([3]uint64{}) {
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
