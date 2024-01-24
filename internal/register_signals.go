//go:build !generate

package gd

import (
	"fmt"
	"reflect"
)

// registerSignals registers [Signal[T]] fields as signals emittable by the class,
// when the class is instantiated, the signal field needs to injected into the field
// so that it can be used and emitted.
func registerSignals(godot Context, class StringName, rtype reflect.Type) {
	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		if field.Type.Implements(reflect.TypeOf([0]IsSignal{}).Elem()) {
			var signalName = godot.StringName(field.Name)
			var ftype = field.Type.Field(1).Type
			if ftype.Kind() != reflect.Func {
				panic("gdextension.RegisterClass: Signal[T] Emit field must be a function")
			}
			if ftype.NumOut() != 0 {
				panic(fmt.Sprintf("gdextension.RegisterClass: %v.%v must not return any values",
					rtype.Name(), field.Name))
			}
			var args []PropertyInfo
			for i := 0; i < ftype.NumIn(); i++ {
				args = append(args, PropertyInfo{
					Type:      variantTypeOf(ftype.In(i)),
					Name:      godot.StringName(fmt.Sprintf("arg%d", i)),
					ClassName: godot.StringName(classNameOf(ftype.In(i))),
				})
			}
			godot.API.ClassDB.RegisterClassSignal(godot.API.ExtensionToken, class, signalName, args)
		}
	}
}
