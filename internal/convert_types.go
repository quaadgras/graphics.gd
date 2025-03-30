//go:build !generate

package gd

import (
	"fmt"
	"reflect"
	"slices"

	"graphics.gd/internal/pointers"
	VariantPkg "graphics.gd/variant"
	ArrayType "graphics.gd/variant/Array"
	DictionaryType "graphics.gd/variant/Dictionary"
	"graphics.gd/variant/Enum"
	PackedType "graphics.gd/variant/Packed"
	"graphics.gd/variant/Path"
	StringType "graphics.gd/variant/String"
	"runtime.link/api/xray"
)

func convertVariantToDesiredGoType(value Variant, rtype reflect.Type) (reflect.Value, error) {
	if value.Type() == TypeNil {
		return reflect.Zero(rtype), nil
	}
	switch rtype {
	case reflect.TypeFor[any]():
		return reflect.ValueOf(value.Interface()), nil
	case reflect.TypeFor[VariantPkg.Any]():
		return reflect.ValueOf(VariantPkg.New(value.Interface())), nil
	}
	switch rtype.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(Global.Variants.Booleanize(value)).Convert(rtype), nil
	case reflect.Array:
		if rtype.Elem().Implements(reflect.TypeOf([0]IsClass{}).Elem()) {
			if value.Type() != TypeObject {
				return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
			}
			object := [1]Object{LetVariantAsPointerType[Object](value, TypeObject)}
			casted := Global.Object.CastTo(object, Global.ClassDB.GetClassTag(NewStringName(classNameOf(rtype))))
			var result = reflect.New(rtype)
			*(*[1]Object)(result.UnsafePointer()) = casted
			return result.Elem(), nil
		}
		fallthrough
	default:
		val, err := ConvertToDesiredGoType(value.Interface(), rtype)
		if err != nil {
			return reflect.Value{}, xray.New(err)
		}
		return val, nil
	}
}

func VariantAs[T any](value Variant) T {
	result, err := ConvertToDesiredGoType(value, reflect.TypeFor[T]())
	if err != nil {
		panic(fmt.Sprintf("cannot convert %T to %s: %v", value, reflect.TypeFor[T](), err))
	}
	return result.Interface().(T)
}

func ConvertToDesiredGoType(value any, rtype reflect.Type) (reflect.Value, error) {
	if reflect.TypeOf(value) == rtype {
		return reflect.ValueOf(value), nil
	}
	if reflect.TypeOf(value).ConvertibleTo(rtype) {
		return reflect.ValueOf(value).Convert(rtype), nil
	}
	variant, ok := value.(Variant)
	if ok {
		val, err := convertVariantToDesiredGoType(variant, rtype)
		if err != nil {
			return reflect.Value{}, xray.New(err)
		}
		return val, err
	}
	switch rtype.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(value).Convert(rtype), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		switch value := value.(type) {
		case Int, Float:
			return reflect.ValueOf(value).Convert(rtype), nil
		default:
			return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
		}
	case reflect.Uint64:
		switch value := value.(type) {
		case RID, Int, Float:
			return reflect.ValueOf(value).Convert(rtype), nil
		default:
			return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
		}
	case reflect.Float32, reflect.Float64:
		switch value := value.(type) {
		case Float, Int:
			return reflect.ValueOf(value).Convert(rtype), nil
		default:
			return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
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
			return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
		}
	case reflect.Array:
		if rtype.Elem().Implements(reflect.TypeOf([0]IsClass{}).Elem()) {
			var obj = reflect.New(rtype)
			*(*Object)(obj.UnsafePointer()) = LetVariantAsPointerType[Object](variant, TypeObject)
			return obj.Elem(), nil
		}
		val, err := convertToGoArrayOf(rtype.Elem(), value)
		if err != nil {
			return reflect.Value{}, xray.New(err)
		}
		return val, nil
	case reflect.Chan:
		return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
	case reflect.Func:
		return convertToGoFunc(rtype, value)
	case reflect.Map:
		return convertToGoMap(rtype, value)
	case reflect.Pointer:
		switch value {
		case nil:
			return reflect.Zero(rtype), nil
		default:
			if rtype.Implements(reflect.TypeOf([0]IsClass{}).Elem()) {
				object, ok := value.(IsClass)
				if !ok {
					return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
				}
				native, ok := ExtensionInstances.Load(pointers.Get(object.AsObject()[0])[0])
				if ok {
					return reflect.ValueOf(native), nil
				}
				return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
			} else {
				var elem = reflect.New(rtype.Elem())
				value, err := ConvertToDesiredGoType(value, rtype.Elem())
				if err != nil {
					return reflect.Value{}, xray.New(err)
				}
				elem.Elem().Set(value)
				return elem, nil
			}
		}
	case reflect.Slice:
		val, err := convertToGoSliceOf(rtype.Elem(), value)
		if err != nil {
			return reflect.Value{}, xray.New(err)
		}
		return val, nil
	case reflect.String:
		if value == nil {
			return reflect.Zero(rtype), nil
		}
		switch value := value.(type) {
		case String:
			return reflect.ValueOf(value.String()).Convert(rtype), nil
		case StringName:
			return reflect.ValueOf(value.String()).Convert(rtype), nil
		case NodePath:
			return reflect.ValueOf(value.String()).Convert(rtype), nil
		case Path.ToNode:
			return reflect.ValueOf(value.String()).Convert(rtype), nil
		case StringType.Readable:
			return reflect.ValueOf(value.String()).Convert(rtype), nil
		case StringType.Name:
			return reflect.ValueOf(value.String()).Convert(rtype), nil
		default:
			return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to stringy %s", value, rtype))
		}
	case reflect.Struct:
		if rtype.Implements(reflect.TypeOf([0]IsClass{}).Elem()) {
			var obj = reflect.New(rtype)
			*(*[1]Object)(obj.UnsafePointer()) = value.(IsClass).AsObject()
			return obj.Elem(), nil
		}
		val, err := convertToGoStruct(rtype, value)
		if err != nil {
			return reflect.Value{}, xray.New(err)
		}
		return val, nil
	default:
		return reflect.Value{}, fmt.Errorf("cannot convert %T to %s", value, rtype)
	}
}

func convertToGoMap(rtype reflect.Type, value any) (reflect.Value, error) {
	switch dictionary := value.(type) {
	case DictionaryType.Any:
		var mapValue = reflect.MakeMap(rtype)
		for _, key := range dictionary.Iter() {
			keyValue, err := convertVariantToDesiredGoType(NewVariant(key), rtype.Key())
			if err != nil {
				return reflect.Value{}, xray.New(err)
			}
			valueValue, err := convertVariantToDesiredGoType(NewVariant(dictionary.Index(key)), rtype.Elem())
			if err != nil {
				return reflect.Value{}, xray.New(err)
			}
			mapValue.SetMapIndex(keyValue, valueValue)
		}
		return mapValue, nil
	case Dictionary:
		var mapValue = reflect.MakeMap(rtype)
		for _, key := range dictionary.Keys().Iter() {
			keyValue, err := convertVariantToDesiredGoType(NewVariant(key), rtype.Key())
			if err != nil {
				return reflect.Value{}, xray.New(err)
			}
			valueValue, err := convertVariantToDesiredGoType(NewVariant(dictionary.Index(key)), rtype.Elem())
			if err != nil {
				return reflect.Value{}, xray.New(err)
			}
			mapValue.SetMapIndex(keyValue, valueValue)
		}
		return mapValue, nil
	default:
		return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
	}
}

func convertToGoStruct(rtype reflect.Type, value any) (reflect.Value, error) {
	i64, isInt64 := value.(Int)
	if reflect.PointerTo(rtype).Implements(reflect.TypeFor[Enum.Pointer]()) && isInt64 {
		enum := reflect.New(rtype)
		enum.Interface().(Enum.Pointer).SetInt(int(i64))
		return enum.Elem(), nil
	}
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
				return reflect.Value{}, xray.New(err)
			}
			structure.Field(i).Set(fieldValue)
		}
		return structure, nil
	case Dictionary:
		var structure = reflect.New(rtype).Elem()
		var dictionary = value
		for i := 0; i < rtype.NumField(); i++ {
			field := rtype.Field(i)
			name := field.Name
			if tag := field.Tag.Get("gd"); tag != "" {
				name = tag
			}
			fieldValue, err := convertVariantToDesiredGoType(dictionary.Index(NewVariant(name)), field.Type)
			if err != nil {
				return reflect.Value{}, xray.New(err)
			}
			structure.Field(i).Set(fieldValue)
		}
		return structure, nil
	case ArrayType.Any:
		if reflect.PointerTo(rtype).Implements(reflect.TypeFor[ArrayType.Pointer]()) {
			var obj = reflect.New(rtype)
			obj.Interface().(ArrayType.Pointer).SetAny(value)
			return obj.Elem(), nil
		}
		return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
	case StringType.Readable:
		if rtype.Kind() == reflect.String {
			return reflect.ValueOf(value.String()), nil
		}
		return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
	case Path.ToNode:
		if rtype.Kind() == reflect.String {
			return reflect.ValueOf(value.String()), nil
		}
		return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
	case StringType.Name:
		if rtype.Kind() == reflect.String {
			return reflect.ValueOf(value.String()), nil
		}
		return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
	case String:
		if rtype.ConvertibleTo(reflect.TypeFor[StringType.Readable]()) {
			return reflect.ValueOf(StringType.New(value.String())).Convert(rtype), nil
		}
		return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
	case StringName:
		if reflect.TypeFor[StringType.Readable]().ConvertibleTo(rtype) {
			return reflect.ValueOf(StringType.Via(StringNameProxy{}, pointers.Pack(value))).Convert(rtype), nil
		}
		return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
	case NodePath:
		if reflect.TypeFor[StringType.Readable]().ConvertibleTo(rtype) {
			return reflect.ValueOf(StringType.Via(NodePathProxy{}, pointers.Pack(value))).Convert(rtype), nil
		}
		return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
	case DictionaryType.Any:
		if reflect.PointerTo(rtype).Implements(reflect.TypeFor[DictionaryType.Pointer]()) {
			var obj = reflect.New(rtype)
			obj.Interface().(DictionaryType.Pointer).SetAny(value)
			return obj.Elem(), nil
		}
		var structure = reflect.New(rtype).Elem()
		var dictionary = value
		for i := 0; i < rtype.NumField(); i++ {
			field := rtype.Field(i)
			name := field.Name
			if tag := field.Tag.Get("gd"); tag != "" {
				name = tag
			}
			fieldValue, err := convertVariantToDesiredGoType(NewVariant(dictionary.Index(VariantPkg.New(name))), field.Type)
			if err != nil {
				return reflect.Value{}, xray.New(err)
			}
			structure.Field(i).Set(fieldValue)
		}
		return structure, nil
	default:
		return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to struct %s", value, rtype))
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
		return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
	}
}

func convertToGoArrayOf(rtype reflect.Type, value any) (reflect.Value, error) {
	if value, ok := value.(ArrayType.Any); ok {
		var internalArray = InternalArray(value)
		var array = reflect.New(rtype).Elem()
		for i := 0; i < rtype.Len(); i++ {
			elem, err := convertVariantToDesiredGoType(internalArray.Index(Int(i)), rtype.Elem())
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
			return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
		}
		return reflect.ValueOf(packed.Bytes()).Convert(rtype), nil
	case reflect.Int32:
		packed, ok := value.(PackedInt32Array)
		if !ok {
			return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
		}
		return reflect.ValueOf(packed.AsSlice()).Convert(rtype), nil
	case reflect.Float32:
		packed, ok := value.(PackedFloat32Array)
		if !ok {
			return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
		}
		return reflect.ValueOf(packed.AsSlice()).Convert(rtype), nil
	case reflect.Float64:
		packed, ok := value.(PackedFloat64Array)
		if !ok {
			return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
		}
		return reflect.ValueOf(packed.AsSlice()).Convert(rtype), nil
	default:
		return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
	}
}

func convertToGoSliceOf(rtype reflect.Type, value any) (reflect.Value, error) {
	if value, ok := value.(ArrayType.Any); ok {
		var array = InternalArray(value)
		var slice = reflect.MakeSlice(reflect.SliceOf(rtype), int(array.Size()), int(array.Size()))
		for i := 0; i < int(array.Size()); i++ {
			elem, err := convertVariantToDesiredGoType(array.Index(Int(i)), rtype)
			if err != nil {
				return reflect.Value{}, xray.New(err)
			}
			slice.Index(i).Set(elem)
		}
		return slice, nil
	}
	if value, ok := value.(Array); ok {
		var slice = reflect.MakeSlice(reflect.SliceOf(rtype), int(value.Size()), int(value.Size()))
		for i := 0; i < int(value.Size()); i++ {
			elem, err := convertVariantToDesiredGoType(value.Index(Int(i)), rtype)
			if err != nil {
				return reflect.Value{}, xray.New(err)
			}
			slice.Index(i).Set(elem)
		}
		return slice, nil
	}
	switch rtype.Kind() {
	case reflect.Uint8:
		packed, ok := value.(PackedType.Bytes)
		if !ok {
			return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
		}
		return reflect.ValueOf(packed.Bytes()).Convert(reflect.SliceOf(rtype)), nil
	case reflect.Int32:
		packed, ok := value.(PackedType.Array[int32])
		if !ok {
			return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
		}
		return reflect.ValueOf(slices.Collect(packed.Values())).Convert(reflect.SliceOf(rtype)), nil
	case reflect.Int64:
		packed, ok := value.(PackedType.Array[int64])
		if !ok {
			return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
		}
		return reflect.ValueOf(slices.Collect(packed.Values())).Convert(reflect.SliceOf(rtype)), nil
	case reflect.Float32:
		packed, ok := value.(PackedType.Array[float32])
		if !ok {
			return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
		}
		return reflect.ValueOf(slices.Collect(packed.Values())).Convert(reflect.SliceOf(rtype)), nil
	case reflect.Float64:
		packed, ok := value.(PackedType.Array[float64])
		if !ok {
			return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
		}
		return reflect.ValueOf(slices.Collect(packed.Values())).Convert(reflect.SliceOf(rtype)), nil
	case reflect.String:
		if value == nil {
			return reflect.Zero(rtype), nil
		}
		packed, ok := value.(PackedType.Strings)
		if !ok {
			return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
		}
		return reflect.ValueOf(packed.Strings()).Convert(reflect.SliceOf(rtype)), nil
	case reflect.Struct:
		switch rtype {
		case reflect.TypeFor[Vector2]():
			packed, ok := value.(PackedType.Array[Vector2])
			if !ok {
				return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
			}
			return reflect.ValueOf(slices.Collect(packed.Values())).Convert(reflect.SliceOf(rtype)), nil
		case reflect.TypeFor[Vector3]():
			packed, ok := value.(PackedType.Array[Vector3])
			if !ok {
				return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
			}
			return reflect.ValueOf(slices.Collect(packed.Values())).Convert(reflect.SliceOf(rtype)), nil
		case reflect.TypeFor[Color]():
			packed, ok := value.(PackedType.Array[Color])
			if !ok {
				return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
			}
			return reflect.ValueOf(slices.Collect(packed.Values())).Convert(reflect.SliceOf(rtype)), nil
		case reflect.TypeFor[Vector4]():
			packed, ok := value.(PackedType.Array[Vector4])
			if !ok {
				return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
			}
			return reflect.ValueOf(slices.Collect(packed.Values())).Convert(reflect.SliceOf(rtype)), nil
		default:
			return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
		}
	default:
		return reflect.Value{}, xray.New(fmt.Errorf("cannot convert %T to %s", value, rtype))
	}
}
