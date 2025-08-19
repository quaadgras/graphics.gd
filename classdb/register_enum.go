package classdb

import (
	"iter"
	"reflect"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Enum"
)

var registered_enums = make(map[reflect.Type]iter.Seq2[string, int])

func registerEnumsFor(class gd.StringName, rtype reflect.Type) iter.Seq2[string, int] {
	if class, ok := registered_enums[rtype]; ok {
		return class
	}
	if !rtype.Implements(reflect.TypeFor[Enum.Any]()) {
		return nil
	}
	var enumName = gd.NewStringName(rtype.Name())
	defer enumName.Free()
	var reference = reflect.New(rtype).Interface().(Enum.Any)
	for name, value := range reference.Enum {
		gdextension.Host.ClassDB.Register.Constant(
			gdextension.StringName(pointers.Get(class)[0]),
			gdextension.StringName(pointers.Get(enumName)[0]),
			gdextension.StringName(pointers.Get(gd.NewStringName(name))[0]),
			int64(value),
			false,
		)
	}
	registered_enums[rtype] = reference.Enum
	return reference.Enum
}
