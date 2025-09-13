// Package loosely provides loose any-to-any type conversions with support for all variant types.
package loosely

import (
	"fmt"
	"reflect"
	"unsafe"

	"graphics.gd/variant"
	"graphics.gd/variant/Array"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Dictionary"
	"graphics.gd/variant/RID"
)

func To[T any](value any) T {
	result, ok := As[T](value)
	if !ok {
		panic(fmt.Sprintf("cannot convert %T to %T", value, result))
	}
	return result
}

func Is[T any](value any) bool {
	_, ok := As[T](value)
	return ok
}

func As[T any](value any) (T, bool) {
	if result, ok := value.(T); ok {
		return result, true
	}
	var result T
	return result, into(reflect.ValueOf(&result), reflect.ValueOf(value))
}

type (
	iZero          interface{ IsZero() bool }
	iInterface     interface{ Interface() any }
	iUnsafePointer interface{ UnsafePointer() unsafe.Pointer }
	iArray         interface{ Any() Array.Any }
	iCallable      interface{ Callable() Callable.Function }
	iDictionary    interface{ Any() Dictionary.Any }
	iRID           interface{ RID() RID.Any }
)

func Into(ptr any, value any) bool {
	return into(reflect.ValueOf(ptr), reflect.ValueOf(value))
}

func into(ptr reflect.Value, value reflect.Value) bool {
	if variant, ok := reflect.TypeAssert[iInterface](value); ok {
		value = reflect.ValueOf(variant.Interface())
	}
	if reflect.TypeOf(ptr).Kind() != reflect.Pointer {
		return false
	}
	rtype := reflect.TypeOf(ptr).Elem()
	rdata := reflect.ValueOf(ptr).Elem()
	vtype := reflect.TypeOf(value)
	vdata := reflect.ValueOf(value)
	if vtype.Kind() == reflect.Array && vtype.Len() == 1 {
		vtype = vtype.Elem()
		vdata = vdata.Index(0)
	}
	if vtype.ConvertibleTo(rtype) {
		rdata.Set(vdata.Convert(rtype))
		return true
	}
	switch rtype.Kind() {
	case reflect.Bool:
		if truthy, ok := reflect.TypeAssert[iZero](vdata); ok {
			rdata.SetBool(!truthy.IsZero())
			return true
		}
		rdata.SetBool(!vdata.IsZero())
		return true
	case reflect.Complex64, reflect.Complex128:
		if vtype.Kind() == reflect.Array && vtype.Elem().ConvertibleTo(reflect.TypeFor[float64]()) {
			rdata.SetComplex(complex(vdata.Index(0).Convert(reflect.TypeFor[float64]()).Float(), vdata.Index(1).Convert(reflect.TypeFor[float64]()).Float()))
			return true
		}
		if vtype.Kind() == reflect.Struct && vtype.NumField() == 2 && vtype.Field(0).Type.ConvertibleTo(reflect.TypeFor[float64]()) && vtype.Field(1).Type.ConvertibleTo(reflect.TypeFor[float64]()) {
			rdata.SetComplex(complex(vdata.Field(0).Convert(reflect.TypeFor[float64]()).Float(), vdata.Field(1).Convert(reflect.TypeFor[float64]()).Float()))
			return true
		}
	case reflect.Array:
		if rtype.Len() == 1 {
			return into(rdata.Index(0).Addr(), vdata)
		}
		if vtype.Kind() == reflect.Array && vtype.Elem().ConvertibleTo(rtype.Elem()) && vtype.Len() <= rtype.Len() {
			for i := 0; i < vdata.Len(); i++ {
				rdata.Index(i).Set(vdata.Index(i).Convert(rtype.Elem()))
			}
			return true
		}
		if array, ok := reflect.TypeAssert[iArray](vdata); ok {
			array := array.Any()
			for i := 0; i < vdata.Len(); i++ {
				if !into(rdata.Index(i).Addr(), reflect.ValueOf(array.Index(i))) {
					return false
				}
			}
			return true
		}
		if vtype.Kind() == reflect.Struct && vtype.NumField() <= rtype.Len() {
			for i := 0; i < vtype.NumField(); i++ {
				if !into(rdata.Index(i).Addr(), vdata.Field(i)) {
					return false
				}
			}
			return true
		}
	case reflect.Func:
		if callable, ok := reflect.TypeAssert[Callable.Function](vdata); ok {
			if rtype.NumOut() <= 1 {
				rdata.Set(reflect.MakeFunc(rtype, func(args []reflect.Value) (results []reflect.Value) {
					var arguments = make([]variant.Any, len(args))
					for i, arg := range args {
						arguments[i] = variant.New(arg.Interface())
					}
					if rtype.NumOut() == 0 {
						callable.Call(arguments...)
						return nil
					}
					var result = reflect.New(rtype.Out(0))
					if !into(result, reflect.ValueOf(callable.Call(arguments...))) {
						panic("failed to loosely convert function result from " + vtype.String() + " to " + rtype.String())
					}
					return []reflect.Value{reflect.ValueOf(result)}
				}))
				return true
			}
		}
	case reflect.Map:
		if vtype.Kind() == reflect.Map {
			values := vdata.MapRange()
			for values.Next() {
				var index = reflect.New(rtype.Key())
				if !into(index, values.Key()) {
					return false
				}
				var value = reflect.New(rtype.Elem())
				if !into(value, values.Value()) {
					return false
				}
				rdata.SetMapIndex(index.Elem(), value.Elem())
			}
		}
		if vtype.Kind() == reflect.Struct {
			for i := 0; i < vtype.NumField(); i++ {
				var index = reflect.New(rtype.Key())
				var field = vtype.Field(i)
				var key = field.Name
				if val, ok := field.Tag.Lookup("gd"); ok {
					key = val
				}
				if !into(index, reflect.ValueOf(key)) {
					return false
				}
				var value = reflect.New(rtype.Elem())
				if !into(value, vdata.Field(i)) {
					return false
				}
				rdata.SetMapIndex(index.Elem(), value.Elem())
			}
		}
		if vtype.Kind() == reflect.Array {
			for i := 0; i < vtype.Len(); i++ {
				var index = reflect.New(rtype.Key())
				if !into(index, reflect.ValueOf(i)) {
					return false
				}
				var value = reflect.New(rtype.Elem())
				if !into(value, vdata.Index(i)) {
					return false
				}
				rdata.SetMapIndex(index.Elem(), value.Elem())
			}
		}
	}
	return false
}
