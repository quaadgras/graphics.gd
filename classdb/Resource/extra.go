package Resource

import (
	"reflect"
	"strings"
	"unsafe"

	"graphics.gd/classdb/ResourceLoader"
	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Path"
	"graphics.gd/variant/String"
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

var preloaded_resources []gd.RefCounted

var startup []func()

func init() {
	preload := func() {
		for _, f := range startup {
			f()
		}
		gd.RegisterCleanup(func() {
			for _, resource := range preloaded_resources {
				if resource.Unreference() {
					gd.Global.Object.Destroy(resource.AsObject())
				}
			}
		})
	}
	Callable.Defer(Callable.New(preload))
}

// Load behaves like the builtin "load" function in GDScript. It can also be used to preload a resource if called before
// startup.
func Load[T Any, P string | Path.ToResource](path_to_resource P) T {
	path := Path.ToResource(String.New(path_to_resource))
	if !gd.Linked {
		var placeholder T
		*(*gd.Object)(unsafe.Pointer(&placeholder)) = pointers.Add[gd.Object]([3]uint64{})
		preloaded_resources = append(preloaded_resources, *(*gd.RefCounted)(unsafe.Pointer(&placeholder)))
		startup = append(startup, func() {
			resource := Instance(ResourceLoader.Load(path.String(), ""))
			result, ok := as[T](resource)
			if !ok {
				panic("Resource \"" + path.String() + "\" is " + resource.AsObject()[0].GetClass().String() + " not " + reflect.TypeFor[T]().String())
			}
			raw, ok := pointers.End(result.AsObject()[0])
			if ok {
				pointers.Set(*(*gd.Object)(unsafe.Pointer(&placeholder)), raw)
			}
		})
		return placeholder
	}
	resource := Instance(ResourceLoader.Load(path.String(), ""))
	result, ok := as[T](resource)
	if !ok {
		panic("Resource \"" + path.String() + "\" is " + resource.AsObject()[0].GetClass().String() + " not " + reflect.TypeFor[T]().String())
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
