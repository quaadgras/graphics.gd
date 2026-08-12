package Resource

import (
	"reflect"
	"unsafe"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdreference"
	"graphics.gd/variant/Path"
	"graphics.gd/variant/String"
)

// Library returns a *T whose class-typed fields refer to the resources named
// by their `gd:"res://..."` tags. Safe to call when declaring a package-level
// variable, before startup: no engine calls are made until a field is first
// used, at which point the resource it refers to is loaded (and kept loaded
// until shutdown). Panics on first use if the resource is not of the field's
// type.
//
//	var PNG = Resource.Library[struct {
//		Bullet Texture2D.Instance `gd:"res://bullet.png"`
//	}]()
func Library[T any]() *T {
	var library = new(T)
	rvalue := reflect.ValueOf(library).Elem()
	rtype := rvalue.Type()
	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		path, ok := field.Tag.Lookup("gd")
		if !ok || path == "" || path == "-" || !field.IsExported() {
			continue
		}
		if field.Type.Size() != unsafe.Sizeof(gdreference.Object{}) || !reflect.PointerTo(field.Type).Implements(reflect.TypeFor[gd.IsClassCastable]()) {
			panic("Resource.Library field " + field.Name + " is " + field.Type.String() + ", expected a class instance type")
		}
		fieldType := field.Type
		resolver := gdreference.NewResolver(func(gdextension.ObjectID) gdextension.Object {
			return loadLibraryResource(path, fieldType)
		})
		*(*gdreference.Object)(unsafe.Pointer(rvalue.Field(i).Addr().Pointer())) = gdreference.DeferObject(0, resolver)
	}
	return library
}

func loadLibraryResource(path string, rtype reflect.Type) gdextension.Object {
	resource := Instance(load(String.Unicode(Path.ToResource(String.New(path))), String.New(""), 1))
	if resource == (Instance{}) {
		return 0
	}
	if castable, ok := reflect.New(rtype).Interface().(gd.IsClassCastable); ok && !castable.SetObject(resource.AsObject()) {
		panic("Resource \"" + path + "\" is " + gd.ObjectGetClass(resource.AsObject()[0]).String() + " not " + rtype.String())
	}
	raw, owned := gdreference.EndObject(resource.AsObject()[0])
	if owned && raw != 0 {
		preloaded_resources = append(preloaded_resources, gd.RefCounted(gdreference.RawObject(raw)))
	}
	return raw
}
