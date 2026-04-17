package gd

import (
	"iter"

	gdunsafe "graphics.gd"
	VariantPkg "graphics.gd/variant"
	ArrayVariant "graphics.gd/variant/Array"
)

func IntsCollectAs[T, S ~int | ~int64 | ~int32](seq iter.Seq[S]) []T {
	var result = make([]T, 0)
	for value := range seq {
		result = append(result, T(value))
	}
	return result
}

func NewArray() gdunsafe.Array {
	return builtin.creation.Array[0](0, nil)
}

func ArrayAs[S []T, T any](array gdunsafe.Array) []T {
	var size = builtin.Array.size.Call(array, struct{}{})
	var result = make([]T, size)
	for i := 0; i < int(size); i++ {
		result[i] = VariantAs[T](array.Index(i))
	}
	return result
}

func ArrayFromSlice[T ArrayVariant.Contains[A], A, B any](slice []B) T {
	var array = NewArray()
	array.Resize(len(slice))
	for i, value := range slice {
		array.SetIndex(i, VariantAs[A](NewVariant(VariantPkg.New(value))))
	}
	return T(array)
}

func EngineArrayFromSlice[T any](slice []T) ArrayVariant.Any {
	var array = ArrayVariant.Through(NewArrayProxy[VariantPkg.Any]())
	array.Resize(len(slice))
	for i, value := range slice {
		array.SetIndex(i, VariantPkg.New(value))
	}
	return array
}
