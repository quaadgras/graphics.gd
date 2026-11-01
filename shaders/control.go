package shaders

import (
	"reflect"

	"graphics.gd/shaders/internal/gpu"
)

// eachValue calls fn with a pointer to each GPU value in value, which is either
// a GPU value itself, or a struct or array that contains them.
func eachValue(value reflect.Value, fn func(ptr reflect.Value)) {
	if isGPUValue(value.Type()) {
		fn(value.Addr())
		return
	}
	switch value.Kind() {
	case reflect.Struct:
		for field, rvalue := range value.Fields() {
			if field.IsExported() {
				eachValue(rvalue, fn)
			}
		}
	case reflect.Array:
		for i := range value.Len() {
			eachValue(value.Index(i), fn)
		}
	}
}

// bind makes the GPU value that ptr points to read the variable l.
func bind(ptr reflect.Value, l *gpu.Local) {
	gpu.Set(ptr.Interface().(gpu.Pointer), gpu.New(l))
}
