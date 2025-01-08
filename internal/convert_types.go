//go:build !generate

package gd

import (
	"fmt"
	"reflect"
)

func convertVariantToDesiredGoType(value Variant, rtype reflect.Type) (reflect.Value, error) {
	switch rtype.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(Global.Variants.Booleanize(value)).Convert(rtype), nil
	case reflect.Array:
		if rtype.Elem().Implements(reflect.TypeOf([0]IsClass{}).Elem()) {
			if value.Type() != TypeObject {
				return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
			}
			object := [1]Object{LetVariantAsPointerType[Object](value, TypeObject)}
			casted := Global.Object.CastTo(object, Global.ClassDB.GetClassTag(NewStringName(classNameOf(rtype))))
			var result = reflect.New(rtype)
			*(*[1]Object)(result.UnsafePointer()) = casted
			return result.Elem(), nil
		}
		fallthrough
	default:
		return ConvertToDesiredGoType(value.Interface(), rtype)
	}
}

func ConvertToDesiredGoType(value any, rtype reflect.Type) (reflect.Value, error) {
	if reflect.TypeOf(value) == rtype {
		return reflect.ValueOf(value), nil
	}
	variant, ok := value.(Variant)
	if ok {
		return convertVariantToDesiredGoType(variant, rtype)
	}
	switch rtype.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(value).Convert(rtype), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		switch value := value.(type) {
		case Int, Float:
			return reflect.ValueOf(value).Convert(rtype), nil
		default:
			return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
		}
	case reflect.Uint64:
		switch value := value.(type) {
		case RID, Int, Float:
			return reflect.ValueOf(value).Convert(rtype), nil
		default:
			return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
		}
	case reflect.Float32, reflect.Float64:
		switch value := value.(type) {
		case Float, Int:
			return reflect.ValueOf(value).Convert(rtype), nil
		default:
			return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
		}
	case reflect.Complex64, reflect.Complex128:
		switch value := value.(type) {
		case Float:
			return reflect.ValueOf(complex(value, 0)).Convert(rtype), nil
		case Int:
			return reflect.ValueOf(complex(float64(value), 0)).Convert(rtype), nil
		case Vector2:
			return reflect.ValueOf(complex(value.X, value.Y)).Convert(rtype), nil
		default:
			return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
		}
	case reflect.Array:
		if rtype.Elem().Implements(reflect.TypeOf([0]IsClass{}).Elem()) {
			var obj = reflect.New(rtype)
			*(*Object)(obj.UnsafePointer()) = LetVariantAsPointerType[Object](variant, TypeObject)
			return obj.Elem(), nil
		}
		return convertToGoArrayOf(rtype.Elem(), value)
	case reflect.Chan:
		return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
	case reflect.Func:
		return convertToGoFunc(rtype, value)
	case reflect.Map:
		return convertToGoMap(rtype, value)
	case reflect.Pointer:
		switch value {
		case nil:
			return reflect.Zero(rtype), nil
		default:
			var elem = reflect.New(rtype.Elem())
			value, err := ConvertToDesiredGoType(value, rtype.Elem())
			if err != nil {
				return reflect.Value{}, err
			}
			elem.Elem().Set(value)
			return elem, nil
		}
	case reflect.Slice:
		return convertToGoSliceOf(rtype.Elem(), value)
	case reflect.String:
		switch value := value.(type) {
		case String:
			return reflect.ValueOf(value.String()).Convert(rtype), nil
		case StringName:
			return reflect.ValueOf(value.String()).Convert(rtype), nil
		case NodePath:
			return reflect.ValueOf(value.String()).Convert(rtype), nil
		default:
			return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
		}
	case reflect.Struct:
		if rtype.Implements(reflect.TypeOf([0]IsClass{}).Elem()) {
			var obj = reflect.New(rtype)
			*(*[1]Object)(obj.UnsafePointer()) = value.(IsClass).AsObject()
			return obj.Elem(), nil
		}
		return convertToGoStruct(rtype, value)
	default:
		return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
	}
}

func convertToGoMap(rtype reflect.Type, value any) (reflect.Value, error) {
	dictionary, ok := value.(Dictionary)
	if !ok {
		return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
	}
	var mapValue = reflect.MakeMap(rtype)
	for _, key := range dictionary.Keys().Iter() {
		keyValue, err := convertVariantToDesiredGoType(key, rtype.Key())
		if err != nil {
			return reflect.Value{}, err
		}
		valueValue, err := convertVariantToDesiredGoType(dictionary.Index(key), rtype.Elem())
		if err != nil {
			return reflect.Value{}, err
		}
		mapValue.SetMapIndex(keyValue, valueValue)
	}
	return mapValue, nil
}

func convertToGoStruct(rtype reflect.Type, value any) (reflect.Value, error) {
	switch value := value.(type) {
	case IsClass:
		var object = value.(IsClass).AsObject()
		var structure = reflect.New(rtype).Elem()
		for i := 0; i < rtype.NumField(); i++ {
			field := rtype.Field(i)
			name := field.Name
			if tag := field.Tag.Get("gd"); tag != "" {
				name = tag
			}
			fieldValue, err := convertVariantToDesiredGoType(object[0].Get(NewStringName(name)), field.Type)
			if err != nil {
				return reflect.Value{}, err
			}
			structure.Field(i).Set(fieldValue)
		}
		return structure, nil
	case Dictionary:
		var structure = reflect.New(rtype)
		var dictionary = value
		for i := 0; i < rtype.NumField(); i++ {
			field := rtype.Field(i)
			name := field.Name
			if tag := field.Tag.Get("gd"); tag != "" {
				name = tag
			}
			fieldValue, err := convertVariantToDesiredGoType(dictionary.Index(NewVariant(name)), field.Type)
			if err != nil {
				return reflect.Value{}, err
			}
			structure.Field(i).Set(fieldValue)
		}
		return structure, nil
	default:
		return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
	}
}

func convertToGoFunc(rtype reflect.Type, value any) (reflect.Value, error) {
	switch value := value.(type) {
	case Callable:
		callable := value
		return reflect.MakeFunc(rtype, func(args []reflect.Value) []reflect.Value {
			variants := make([]Variant, len(args))
			for i, arg := range args {
				variants[i] = NewVariant(arg.Interface())
			}
			if rtype.NumOut() == 0 {
				callable.Call(variants...)
				return nil
			}
			val, err := convertVariantToDesiredGoType(callable.Call(variants...), rtype.Out(0))
			if err != nil {
				panic("invalid return type")
			}
			return []reflect.Value{
				val,
			}
		}), nil
	case Signal:
		signal := value
		return reflect.MakeFunc(rtype, func(args []reflect.Value) []reflect.Value {
			variants := make([]Variant, len(args))
			for i, arg := range args {
				variants[i] = NewVariant(arg.Interface())
			}
			signal.Emit(variants...)
			return nil
		}), nil
	default:
		return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
	}
}

func convertToGoArrayOf(rtype reflect.Type, value any) (reflect.Value, error) {
	if value, ok := value.(Array); ok {
		var array = reflect.New(rtype).Elem()
		for i := 0; i < rtype.Len(); i++ {
			elem, err := convertVariantToDesiredGoType(value.Index(Int(i)), rtype.Elem())
			if err != nil {
				return reflect.Value{}, err
			}
			array.Index(i).Set(elem)
		}
		return array, nil
	}
	switch rtype.Kind() {
	case reflect.Uint8:
		packed, ok := value.(PackedByteArray)
		if !ok {
			return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
		}
		return reflect.ValueOf(packed.Bytes()).Convert(rtype), nil
	case reflect.Int32:
		packed, ok := value.(PackedInt32Array)
		if !ok {
			return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
		}
		return reflect.ValueOf(packed.AsSlice()).Convert(rtype), nil
	case reflect.Float32:
		packed, ok := value.(PackedFloat32Array)
		if !ok {
			return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
		}
		return reflect.ValueOf(packed.AsSlice()).Convert(rtype), nil
	case reflect.Float64:
		packed, ok := value.(PackedFloat64Array)
		if !ok {
			return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
		}
		return reflect.ValueOf(packed.AsSlice()).Convert(rtype), nil
	default:
		return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
	}
}

func convertToGoSliceOf(rtype reflect.Type, value any) (reflect.Value, error) {
	if value, ok := value.(Array); ok {
		var slice = reflect.MakeSlice(reflect.SliceOf(rtype), int(value.Size()), int(value.Size()))
		for i := 0; i < int(value.Size()); i++ {
			elem, err := convertVariantToDesiredGoType(value.Index(Int(i)), rtype)
			if err != nil {
				return reflect.Value{}, err
			}
			slice.Index(i).Set(elem)
		}
		return slice, nil
	}
	switch rtype.Kind() {
	case reflect.Uint8:
		packed, ok := value.(PackedByteArray)
		if !ok {
			return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
		}
		return reflect.ValueOf(packed.Bytes()).Convert(reflect.SliceOf(rtype)), nil
	case reflect.Int32:
		packed, ok := value.(PackedInt32Array)
		if !ok {
			return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
		}
		return reflect.ValueOf(packed.AsSlice()).Convert(reflect.SliceOf(rtype)), nil
	case reflect.Int64:
		packed, ok := value.(PackedInt64Array)
		if !ok {
			return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
		}
		return reflect.ValueOf(packed.AsSlice()).Convert(reflect.SliceOf(rtype)), nil
	case reflect.Float32:
		packed, ok := value.(PackedFloat32Array)
		if !ok {
			return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
		}
		return reflect.ValueOf(packed.AsSlice()).Convert(reflect.SliceOf(rtype)), nil
	case reflect.Float64:
		packed, ok := value.(PackedFloat64Array)
		if !ok {
			return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
		}
		return reflect.ValueOf(packed.AsSlice()).Convert(reflect.SliceOf(rtype)), nil
	case reflect.String:
		packed, ok := value.(PackedStringArray)
		if !ok {
			return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
		}
		return reflect.ValueOf(packed.Strings()).Convert(reflect.SliceOf(rtype)), nil
	case reflect.Struct:
		switch rtype {
		case reflect.TypeFor[Vector2]():
			packed, ok := value.(PackedVector2Array)
			if !ok {
				return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
			}
			return reflect.ValueOf(packed.AsSlice()).Convert(reflect.SliceOf(rtype)), nil
		case reflect.TypeFor[Vector3]():
			packed, ok := value.(PackedVector3Array)
			if !ok {
				return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
			}
			return reflect.ValueOf(packed.AsSlice()).Convert(reflect.SliceOf(rtype)), nil
		case reflect.TypeFor[Color]():
			packed, ok := value.(PackedColorArray)
			if !ok {
				return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
			}
			return reflect.ValueOf(packed.AsSlice()).Convert(reflect.SliceOf(rtype)), nil
		case reflect.TypeFor[Vector4]():
			packed, ok := value.(PackedVector4Array)
			if !ok {
				return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
			}
			return reflect.ValueOf(packed.AsSlice()).Convert(reflect.SliceOf(rtype)), nil
		default:
			return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
		}
	default:
		return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
	}
}
