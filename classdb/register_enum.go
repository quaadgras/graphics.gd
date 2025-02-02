package classdb

import (
	"iter"
	"reflect"

	gd "graphics.gd/internal"
	"graphics.gd/variant/Enum"
)

var registered_enums = make(map[reflect.Type]iter.Seq2[string, int])

func registerEnumsFor(class gd.StringName, rtype reflect.Type) iter.Seq2[string, int] {
	if gd.Global.ClassDB.RegisterClassIntegerConstant == nil {
		return nil
	}
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
		gd.Global.ClassDB.RegisterClassIntegerConstant(
			gd.Global.ExtensionToken,
			class,
			enumName,
			gd.NewStringName(name),
			int64(value),
			false,
		)
	}
	registered_enums[rtype] = reference.Enum
	return reference.Enum
}
