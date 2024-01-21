package gdextension

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"

	gd "grow.graphics/gd/internal"
	"runtime.link/api/call"
	"runtime.link/mmm"
)

// registerSignals registers [gd.Signal[T]] fields as signals emittable by the class,
// when the class is instantiated, the signal field needs to injected into the field
// so that it can be used and emitted.
func registerSignals(godot gd.Context, class gd.StringName, rtype reflect.Type) {
	ctx := gd.NewContext(godot.API)
	defer ctx.End()

	var pin runtime.Pinner
	defer pin.Unpin()

	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		if field.Type.Implements(reflect.TypeOf([0]gd.IsSignal{}).Elem()) {
			var signalName = godot.StringName(field.Name)
			var ftype = field.Type.Field(1).Type
			if ftype.Kind() != reflect.Func {
				panic("gdextension.RegisterClass: Signal[T] Emit field must be a function")
			}
			if ftype.NumOut() != 0 {
				panic(fmt.Sprintf("gdextension.RegisterClass: %v.%v must not return any values",
					rtype.Name(), field.Name))
			}
			var args []gd.PropertyInfo
			for i := 0; i < ftype.NumIn(); i++ {
				args = append(args, derivePropertyInfo(ctx, &pin, rtype, reflect.StructField{Type: ftype.In(i)}))
			}

			var frame = call.New()
			godot.API.ClassDB.RegisterClassSignal(godot.API.ExtensionToken,
				(gd.StringNamePtr)(unsafe.Pointer(call.Arg(frame, mmm.Get(class)).Uintptr())),
				(gd.StringNamePtr)(unsafe.Pointer(call.Arg(frame, mmm.Get(signalName)).Uintptr())),
				args,
			)
			frame.Free()
		}
	}
}
