package classdb

import (
	"fmt"
	"reflect"
	"strings"

	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Signal"
	"graphics.gd/variant/String"
)

// registerSignals registers [Signal.Any], [Signal.Pair], [Signal.Trio] and [Signal.Quad] fields as signals
// emittable by the class, when the class is instantiated, the signal field needs to injected into the field
// so that it can be used and emitted.
func registerSignals(class gd.StringName, rtype reflect.Type) {
	for _, field := range reflect.VisibleFields(rtype) {
		name := String.ToSnakeCase(field.Name)
		if !field.IsExported() {
			continue
		}
		if tag := field.Tag.Get("gd"); tag != "" {
			name = tag
		}
		name = strings.TrimSuffix(name, ")")
		name, args, _ := strings.Cut(name, "(")
		argNames := strings.Split(args, ",")
		if reflect.PointerTo(field.Type).Implements(reflect.TypeFor[Signal.Pointer]()) {
			var signalName = gd.NewStringName(name)
			var emit, ok = field.Type.MethodByName("Emit")
			if !ok {
				panic("gdextension.RegisterClass: Signal[T] Emit method not found")
			}
			ftype := emit.Type
			if ftype.Kind() != reflect.Func {
				panic("gdextension.RegisterClass: Signal[T] Emit field must be a function")
			}
			if ftype.NumOut() != 0 {
				panic(fmt.Sprintf("gdextension.RegisterClass: %v.%v must not return any values",
					rtype.Name(), name))
			}
			var args []gd.PropertyInfo
			for i := 1; i < ftype.NumIn(); i++ {
				vtype, ok := gd.VariantTypeOf(ftype.In(i))
				if ok {
					name := fmt.Sprintf("arg%d", i)
					if i-1 < len(argNames) {
						name = argNames[i-1]
					}
					args = append(args, gd.PropertyInfo{
						Type:      vtype,
						Name:      gd.NewStringName(name),
						ClassName: gd.NewStringName(nameOf(ftype.In(i))),
					})
				}
			}
			gd.Global.ClassDB.RegisterClassSignal(gd.Global.ExtensionToken, class, signalName, args)
		}
		if field.Type.Kind() == reflect.Chan && field.Type.ChanDir() == reflect.SendDir {
			var signalName = gd.NewStringName(name)
			var etype = field.Type.Elem()
			var args []gd.PropertyInfo
			if etype.Kind() == reflect.Func {
				for i := 0; i < etype.NumOut(); i++ {
					arg := etype.Out(i)
					vtype, ok := gd.VariantTypeOf(arg)
					if ok {
						name := fmt.Sprintf("arg%d", i)
						if i < len(argNames) {
							name = argNames[i]
						}
						args = append(args, gd.PropertyInfo{
							Type:      vtype,
							Name:      gd.NewStringName(name),
							ClassName: gd.NewStringName(nameOf(arg)),
						})
					}
				}
			} else if !(etype.Kind() == reflect.Struct && etype.NumField() == 0) {
				vtype, ok := gd.VariantTypeOf(etype)
				if ok {
					name := fmt.Sprintf("event")
					if len(argNames) > 0 {
						name = argNames[0]
					}
					args = append(args, gd.PropertyInfo{
						Type:      vtype,
						Name:      gd.NewStringName(name),
						ClassName: gd.NewStringName(nameOf(etype)),
					})
				}
			}
			gd.Global.ClassDB.RegisterClassSignal(gd.Global.ExtensionToken, class, signalName, args)
		}
	}
}

type signalChan struct {
	signal gd.Signal
	rvalue reflect.Value
}

func manageSignals(instance gd.Int, signals []signalChan) {
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
		gd.NewCallable(func() {
			lookup := gd.Global.Object.GetInstanceFromID(gd.ObjectID(instance))
			if lookup == ([1]gd.Object{}) {
				panic("manageSignals: object freed")
			}
			pointers.End(lookup[0])
			rtype := value.Type()
			if rtype.Kind() == reflect.Struct && rtype.NumField() == 0 {
				signal.signal.Emit()
			} else if rtype.Kind() == reflect.Func {
				var args []gd.Variant
				results := value.Call(nil)
				for _, result := range results {
					args = append(args, gd.NewVariant(result.Interface()))
				}
				signal.signal.Emit(args...)
			} else {
				signal.signal.Emit(gd.NewVariant(value.Interface()))
			}
		}).CallDeferred()
	}
}
