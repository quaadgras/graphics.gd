package variant

import (
	"reflect"
)

// As coerces a variant to a specific type.
func As[T any](variant Any) T {
	al, ok := any(variant).(T)
	if ok {
		return al
	}
	ez, ok := variant.Interface().(T)
	if ok {
		return ez
	}
	var val = reflect.New(reflect.TypeFor[T]()).Elem()
	switch val.Kind() {
	case reflect.Bool:
		val.SetBool(variant.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val.SetInt(int64(variant.Int()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val.SetUint(uint64(variant.Uint()))
	case reflect.Float32, reflect.Float64:
		val.SetFloat(variant.Float64())
	case reflect.String:
		val.SetString(variant.String())
	case reflect.Complex64, reflect.Complex128:
		val.SetComplex(variant.Complex128())
	case reflect.Interface:
		if reflect.TypeFor[T]() == reflect.TypeFor[any]() {
			val.Set(reflect.ValueOf(variant.Interface()))
		}
	}
	return val.Interface().(T)
}
