//go:build !generate

package gd

import (
	"errors"
	"reflect"
)

var ErrInvalidParameter = errors.New("invalid parameter")

func ConvertVariantTypeToDesiredGoType(value Variant, rtype reflect.Type) (reflect.Value, error) {
	switch rtype.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(Global.Variants.Booleanize(value)).Convert(rtype), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		switch value.Type() {
		case TypeInt:
			return reflect.ValueOf(value.Interface().(Int)).Convert(rtype), nil
		case TypeFloat:
			return reflect.ValueOf(int(value.Interface().(Float))).Convert(rtype), nil
		default:
			return reflect.Value{}, ErrInvalidParameter
		}
	case reflect.Uint64:
		switch value.Type() {
		case TypeRID:
			return reflect.ValueOf(value.Interface().(RID)).Convert(rtype), nil
		case TypeInt:
			return reflect.ValueOf(uint64(value.Interface().(Int))).Convert(rtype), nil
		case TypeFloat:
			return reflect.ValueOf(uint64(value.Interface().(Float))).Convert(rtype), nil
		default:
			return reflect.Value{}, ErrInvalidParameter
		}
	case reflect.Float32, reflect.Float64:
		switch value.Type() {
		case TypeFloat:
			return reflect.ValueOf(value.Interface().(Float)).Convert(rtype), nil
		case TypeInt:
			return reflect.ValueOf(float64(value.Interface().(Int))).Convert(rtype), nil
		default:
			return reflect.Value{}, ErrInvalidParameter
		}
	case reflect.Complex64, reflect.Complex128:
		switch value.Type() {
		case TypeFloat:
			return reflect.ValueOf(complex(value.Interface().(Float), 0)).Convert(rtype), nil
		case TypeInt:
			return reflect.ValueOf(complex(float64(value.Interface().(Int)), 0)).Convert(rtype), nil
		case TypeVector2:
			vec2 := value.Interface().(Vector2)
			return reflect.ValueOf(complex(vec2.X, vec2.Y)).Convert(rtype), nil
		default:
			return reflect.Value{}, ErrInvalidParameter
		}
	case reflect.Array:
		return convertVariantTypeToArrayOf(rtype.Elem(), value)
	case reflect.Chan:
		return reflect.Value{}, ErrInvalidParameter
	case reflect.Func:
		return convertVariantTypeToFunc(rtype, value)
	case reflect.Map:
		return convertVariantTypeToMap(rtype, value)
	case reflect.Pointer:
		switch value.Type() {
		case TypeNil:
			return reflect.Zero(rtype), nil
		default:
			var elem = reflect.New(rtype.Elem())
			value, err := ConvertVariantTypeToDesiredGoType(value, rtype.Elem())
			if err != nil {
				return reflect.Value{}, err
			}
			elem.Elem().Set(value)
			return elem, nil
		}
	case reflect.Slice:
		return convertVariantTypeToSliceOf(rtype.Elem(), value)
	case reflect.String:
		switch value.Type() {
		case TypeString:
			return reflect.ValueOf(value.Interface().(String).String()).Convert(rtype), nil
		case TypeStringName:
			return reflect.ValueOf(value.Interface().(StringName).String()).Convert(rtype), nil
		default:
			return reflect.Value{}, ErrInvalidParameter
		}
	case reflect.Struct:
		if rtype.Implements(reflect.TypeOf([0]IsClass{}).Elem()) {
			var obj = reflect.New(rtype)
			*(*Object)(obj.UnsafePointer()) = LetVariantAsPointerType[Object](value, TypeObject)
			return obj.Elem(), nil
		}
		return convertVariantTypeToStruct(rtype, value)
	default:
		return reflect.Value{}, ErrInvalidParameter
	}
}

func convertVariantTypeToMap(rtype reflect.Type, value Variant) (reflect.Value, error) {
	if value.Type() != TypeDictionary {
		return reflect.Value{}, ErrInvalidParameter
	}
	var dictionary = value.Interface().(Dictionary)
	var mapValue = reflect.MakeMap(rtype)
	for _, key := range dictionary.Keys().Iter() {
		keyValue, err := ConvertVariantTypeToDesiredGoType(key, rtype.Key())
		if err != nil {
			return reflect.Value{}, err
		}
		valueValue, err := ConvertVariantTypeToDesiredGoType(dictionary.Index(key), rtype.Elem())
		if err != nil {
			return reflect.Value{}, err
		}
		mapValue.SetMapIndex(keyValue, valueValue)
	}
	return mapValue, nil
}

func convertVariantTypeToStruct(rtype reflect.Type, value Variant) (reflect.Value, error) {
	switch value.Type() {
	case TypeObject:
		var object = value.Interface().(IsClass).AsObject()
		var structure = reflect.New(rtype).Elem()
		for i := 0; i < rtype.NumField(); i++ {
			field := rtype.Field(i)
			name := field.Name
			if tag := field.Tag.Get("gd"); tag != "" {
				name = tag
			}
			fieldValue, err := ConvertVariantTypeToDesiredGoType(object.Get(NewStringName(name)), field.Type)
			if err != nil {
				return reflect.Value{}, err
			}
			structure.Field(i).Set(fieldValue)
		}
		return structure, nil
	case TypeDictionary:
		var structure = reflect.New(rtype)
		var dictionary = value.Interface().(Dictionary)
		for i := 0; i < rtype.NumField(); i++ {
			field := rtype.Field(i)
			name := field.Name
			if tag := field.Tag.Get("gd"); tag != "" {
				name = tag
			}
			fieldValue, err := ConvertVariantTypeToDesiredGoType(dictionary.Index(NewVariant(name)), field.Type)
			if err != nil {
				return reflect.Value{}, err
			}
			structure.Field(i).Set(fieldValue)
		}
		return structure, nil
	default:
		return reflect.Value{}, ErrInvalidParameter
	}
}

func convertVariantTypeToFunc(rtype reflect.Type, value Variant) (reflect.Value, error) {
	switch value.Type() {
	case TypeCallable:
		callable := value.Interface().(Callable)
		return reflect.MakeFunc(rtype, func(args []reflect.Value) []reflect.Value {
			variants := make([]Variant, len(args))
			for i, arg := range args {
				variants[i] = NewVariant(arg.Interface())
			}
			if rtype.NumOut() == 0 {
				callable.Call(variants...)
				return nil
			}
			val, err := ConvertVariantTypeToDesiredGoType(callable.Call(variants...), rtype.Out(0))
			if err != nil {
				panic("invalid return type")
			}
			return []reflect.Value{
				val,
			}
		}), nil
	case TypeSignal:
		signal := value.Interface().(Signal)
		return reflect.MakeFunc(rtype, func(args []reflect.Value) []reflect.Value {
			variants := make([]Variant, len(args))
			for i, arg := range args {
				variants[i] = NewVariant(arg.Interface())
			}
			signal.Emit(variants...)
			return nil
		}), nil
	default:
		return reflect.Value{}, ErrInvalidParameter
	}
}

func convertVariantTypeToArrayOf(rtype reflect.Type, value Variant) (reflect.Value, error) {
	if value.Type() == TypeArray {
		var array = reflect.New(rtype).Elem()
		var godot = value.Interface().(Array)
		for i := 0; i < rtype.Len(); i++ {
			elem, err := ConvertVariantTypeToDesiredGoType(godot.Index(Int(i)), rtype.Elem())
			if err != nil {
				return reflect.Value{}, err
			}
			array.Index(i).Set(elem)
		}
		return array, nil
	}
	switch rtype.Kind() {
	case reflect.Uint8:
		if value.Type() != TypePackedByteArray {
			return reflect.Value{}, ErrInvalidParameter
		}
		return reflect.ValueOf(value.Interface().(PackedByteArray).Bytes()).Convert(rtype), nil
	case reflect.Int32:
		if value.Type() != TypePackedInt32Array {
			return reflect.Value{}, ErrInvalidParameter
		}
		return reflect.ValueOf(value.Interface().(PackedInt32Array).AsSlice()).Convert(rtype), nil
	case reflect.Float32:
		if value.Type() != TypePackedFloat32Array {
			return reflect.Value{}, ErrInvalidParameter
		}
		return reflect.ValueOf(value.Interface().(PackedFloat32Array).AsSlice()).Convert(rtype), nil
	case reflect.Float64:
		if value.Type() != TypePackedFloat64Array {
			return reflect.Value{}, ErrInvalidParameter
		}
		return reflect.ValueOf(value.Interface().(PackedFloat64Array).AsSlice()).Convert(rtype), nil
	default:
		return reflect.Value{}, ErrInvalidParameter
	}
}

func convertVariantTypeToSliceOf(rtype reflect.Type, value Variant) (reflect.Value, error) {
	if value.Type() == TypeArray {
		var godot = value.Interface().(Array)
		var slice = reflect.MakeSlice(reflect.SliceOf(rtype), int(godot.Size()), int(godot.Size()))
		for i := 0; i < int(godot.Size()); i++ {
			elem, err := ConvertVariantTypeToDesiredGoType(godot.Index(Int(i)), rtype)
			if err != nil {
				return reflect.Value{}, err
			}
			slice.Index(i).Set(elem)
		}
		return slice, nil
	}
	switch rtype.Kind() {
	case reflect.Uint8:
		if value.Type() != TypePackedByteArray {
			return reflect.Value{}, ErrInvalidParameter
		}
		return reflect.ValueOf(value.Interface().(PackedByteArray).Bytes()).Convert(rtype), nil
	case reflect.Int32:
		if value.Type() != TypePackedInt32Array {
			return reflect.Value{}, ErrInvalidParameter
		}
		return reflect.ValueOf(value.Interface().(PackedInt32Array).AsSlice()).Convert(rtype), nil
	case reflect.Float32:
		if value.Type() != TypePackedFloat32Array {
			return reflect.Value{}, ErrInvalidParameter
		}
		return reflect.ValueOf(value.Interface().(PackedFloat32Array).AsSlice()).Convert(rtype), nil
	case reflect.Float64:
		if value.Type() != TypePackedFloat64Array {
			return reflect.Value{}, ErrInvalidParameter
		}
		return reflect.ValueOf(value.Interface().(PackedFloat64Array).AsSlice()).Convert(rtype), nil
	case reflect.String:
		if value.Type() != TypePackedStringArray {
			return reflect.Value{}, ErrInvalidParameter
		}
		return reflect.ValueOf(value.Interface().(PackedStringArray).AsSlice()).Convert(rtype), nil
	case reflect.Struct:
		switch rtype {
		case reflect.TypeFor[Vector2]():
			if value.Type() != TypePackedVector2Array {
				return reflect.Value{}, ErrInvalidParameter
			}
			return reflect.ValueOf(value.Interface().(PackedVector2Array).AsSlice()).Convert(rtype), nil
		case reflect.TypeFor[Vector3]():
			if value.Type() != TypePackedVector3Array {
				return reflect.Value{}, ErrInvalidParameter
			}
			return reflect.ValueOf(value.Interface().(PackedVector3Array).AsSlice()).Convert(rtype), nil
		case reflect.TypeFor[Color]():
			if value.Type() != TypePackedColorArray {
				return reflect.Value{}, ErrInvalidParameter
			}
			return reflect.ValueOf(value.Interface().(PackedColorArray).AsSlice()).Convert(rtype), nil
		case reflect.TypeFor[Vector4]():
			if value.Type() != TypePackedVector4Array {
				return reflect.Value{}, ErrInvalidParameter
			}
			return reflect.ValueOf(value.Interface().(PackedVector4Array).AsSlice()).Convert(rtype), nil
		default:
			return reflect.Value{}, ErrInvalidParameter
		}
	default:
		return reflect.Value{}, ErrInvalidParameter
	}
}
