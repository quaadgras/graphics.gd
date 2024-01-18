package gdextension

import (
	"reflect"
	"runtime"
	"runtime/cgo"
	"strings"
	"unsafe"

	"grow.graphics/gd"
	internal "grow.graphics/gd/internal"
)

func registerMethods(godot gd.Context, class gd.StringName, rtype reflect.Type) {
	var pinner runtime.Pinner
	godot.Defer(func() {
		pinner.Unpin()
	})

	pinner.Pin(&class)

	classTypePtr := reflect.PointerTo(rtype)

	var returnInfo = (*internal.PropertyInfo)(godot.API().Memory.Allocate(unsafe.Sizeof(internal.PropertyInfo{})))

	for i := 0; i < classTypePtr.NumMethod(); i++ {
		method := classTypePtr.Method(i)
		if !method.IsExported() {
			continue
		}
		if method.Type.NumIn() > 2 || method.Type.NumOut() > 0 {
			continue
		}
		if method.Type.In(1) != reflect.TypeOf(internal.Context{}) {
			continue
		}

		var methodName = godot.StringName(strings.ToLower(method.Name))
		pinner.Pin(&methodName)

		var info = internal.ClassMethodInfo{
			Name:            (internal.StringNamePtr)(unsafe.Pointer(&methodName)),
			HasReturnValue:  method.Type.NumOut() > 0,
			ReturnValueInfo: returnInfo,
		}
		info.PointerCall.Set(func(userdata unsafe.Pointer, instance cgo.Handle, args, ret unsafe.Pointer, err *internal.CallError) {
			ctx := internal.NewContext(godot.API())
			var in = make([]reflect.Value, method.Type.NumIn())
			in[0] = reflect.ValueOf(instance.Value())
			in[1] = reflect.ValueOf(ctx)
			var out = method.Func.Call(in)
			if len(out) > 0 {
				reflect.NewAt(method.Type.Out(0), ret).Elem().Set(out[0])
			}
			ctx.Free()
		})
		godot.API().ClassDB.RegisterClassMethod(godot.API().ExtensionToken, (internal.StringNamePtr)(unsafe.Pointer(&class)), &info)
	}
}
