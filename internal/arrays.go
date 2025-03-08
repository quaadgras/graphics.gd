package gd

import (
	"fmt"
	"iter"
	"reflect"

	"graphics.gd/internal/callframe"
	"graphics.gd/internal/pointers"
	VariantPkg "graphics.gd/variant"
	ArrayVariant "graphics.gd/variant/Array"
)

func (a Array) Index(index int64) Variant {
	return Global.Array.Index(a, index)
}

func (a Array) SetIndex(index int64, value Variant) {
	Global.Array.SetIndex(a, index, value)
	pointers.End(value)
}

func (a Array) isArray() {}

func (a Array) end() { pointers.End(a) }

func (a Array) Free() {
	if ptr, ok := pointers.End(a); ok {
		var frame = callframe.New()
		Global.typeset.destruct.Array(callframe.Arg(frame, ptr).Addr())
		frame.Free()
	}
}

func (a Array) Iter() iter.Seq2[int64, Variant] {
	return func(yield func(int64, Variant) bool) {
		for i := int64(0); i < a.Size(); i++ {
			if !yield(i, a.Index(i)) {
				break
			}
		}
	}
}

func NewArray() Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]EnginePointer](frame)
	Global.typeset.creation.Array[0](r_ret.Addr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[Array](raw)
}

func ArrayAs[S []T, T any](array Array) []T {
	var result = make([]T, array.Size())
	for i := 0; i < int(array.Size()); i++ {
		result[i] = VariantAs[T](array.Index(int64(i)))
	}
	return result
}

func InternalArray[T any](array ArrayVariant.Contains[T]) Array {
	_, state := ArrayVariant.As(array, NewArrayProxy[T])
	return pointers.Load[Array](state)
}

func ArrayFromSlice[T ArrayVariant.Contains[A], A, B any](slice []B) T {
	var array = ArrayVariant.Through(NewArrayProxy[A]())
	array.Resize(len(slice))
	for i, value := range slice {
		array.SetIndex(i, VariantPkg.New(value).Interface().(A))
	}
	return T(array)
}

func EngineArrayFromSlice(slice []any) ArrayVariant.Any {
	var array = ArrayVariant.Through(NewArrayProxy[VariantPkg.Any]())
	array.Resize(len(slice))
	for i, value := range slice {
		array.SetIndex(i, VariantPkg.New(value))
	}
	return array
}

func NewArrayProxy[T any]() (ArrayProxy[T], complex128) {
	var array = NewArray()
	var pack = pointers.Pack(array)
	return ArrayProxy[T]{}, pack
}

type ArrayProxy[T any] struct{}

func (ArrayProxy[T]) Any(state complex128) ArrayVariant.Any {
	return ArrayVariant.Through(ArrayProxy[VariantPkg.Any]{}, state)
}

func (ArrayProxy[T]) Resize(state complex128, i int) {
	pointers.Load[Array](state).Resize(int64(i))
}
func (ArrayProxy[T]) Index(state complex128, i int) T {
	value, err := convertVariantToDesiredGoType(pointers.Load[Array](state).Index(int64(i)), reflect.TypeFor[T]())
	if err != nil {
		panic(fmt.Sprintf("could not convert variant to desired go type: %v", err))
	}
	return value.Interface().(T)
}
func (ArrayProxy[T]) SetIndex(state complex128, i int, val T) {
	pointers.Load[Array](state).SetIndex(int64(i), NewVariant(val))
}
func (ArrayProxy[T]) Len(state complex128) int {
	return int(pointers.Load[Array](state).Size())
}
func (ArrayProxy[T]) IsReadOnly(state complex128) bool {
	return bool(pointers.Load[Array](state).IsReadOnly())
}
func (ArrayProxy[T]) MakeReadOnly(state complex128) {
	pointers.Load[Array](state).MakeReadOnly()
}
