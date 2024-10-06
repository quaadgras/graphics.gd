//go:build !generate

package gd

import (
	"fmt"
	"reflect"

	gd "grow.graphics/gd/internal"
)

// registerSignals registers [Signal[T]] fields as signals emittable by the class,
// when the class is instantiated, the signal field needs to injected into the field
// so that it can be used and emitted.
func registerSignals(godot Lifetime, class StringName, rtype reflect.Type) {
	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		name := field.Name
		if !field.IsExported() {
			continue
		}
		if tag := field.Tag.Get("gd"); tag != "" {
			name = tag
		}
		if field.Type.Implements(reflect.TypeOf([0]gd.IsSignal{}).Elem()) {
			var signalName = godot.StringName(name)
			var ftype = field.Type.Field(1).Type
			if ftype.Kind() != reflect.Func {
				panic("gdextension.RegisterClass: Signal[T] Emit field must be a function")
			}
			if ftype.NumOut() != 0 {
				panic(fmt.Sprintf("gdextension.RegisterClass: %v.%v must not return any values",
					rtype.Name(), name))
			}
			var args []gd.PropertyInfo
			for i := 0; i < ftype.NumIn(); i++ {
				args = append(args, gd.PropertyInfo{
					Type:      variantTypeOf(ftype.In(i)),
					Name:      godot.StringName(fmt.Sprintf("arg%d", i)),
					ClassName: godot.StringName(classNameOf(ftype.In(i))),
				})
			}
			godot.API.ClassDB.RegisterClassSignal(godot.API.ExtensionToken, class, signalName, args)
		}
		if field.Type.Kind() == reflect.Chan && field.Type.ChanDir() == reflect.SendDir {
			var signalName = godot.StringName(name)
			var etype = field.Type.Elem()
			var args []gd.PropertyInfo
			if !(etype.Kind() == reflect.Struct && etype.NumField() == 0) {
				args = append(args, gd.PropertyInfo{
					Type:      variantTypeOf(etype),
					Name:      godot.StringName("event"),
					ClassName: godot.StringName(classNameOf(etype)),
				})
			}
			godot.API.ClassDB.RegisterClassSignal(godot.API.ExtensionToken, class, signalName, args)
		}
	}
}

type signalChan struct {
	signal gd.Signal
	rvalue reflect.Value
}

func manageSignals(godot Lifetime, instance Int, signals []signalChan) {
	var cases = make([]reflect.SelectCase, len(signals))
	for i, signal := range signals {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: signal.rvalue,
		}
	}
	for {
		chosen, value, ok := reflect.Select(cases)
		if !ok {
			cases[chosen] = cases[len(cases)-1]
			cases = cases[:len(cases)-1]
			if len(cases) == 0 {
				break
			}
		}
		signal := signals[chosen]
		godot.Callable(func() {
			tmp := gd.NewLifetime(godot.API)
			if godot.API.Object.GetInstanceFromID(tmp, gd.ObjectID(instance)) == (gd.Object{}) {
				panic("manageSignals: object freed")
			}
			rtype := value.Type()
			if rtype.Kind() == reflect.Struct && rtype.NumField() == 0 {
				signal.signal.Emit()
			} else {
				signal.signal.Emit(tmp.Variant(value.Interface()))
			}
		}).CallDeferred()
	}
}
