package gd

import (
	"fmt"
	"iter"
	"reflect"
	"runtime"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/noescape"
	"graphics.gd/internal/pointers"
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

func (a Array) Index(index int64) Variant {
	var raw [3]uint64
	// Array.Get is a shallow memcpy of the element header (see gd_array_get);
	// it does not add a reference. Copy() performs a real variant_new_copy so
	// the returned Variant owns an independent value that is safe to Free.
	gdextension.Host.Array.Get(pointers.Get(a), int(index), gdextension.CallReturns[gdextension.Variant](&raw[0]))
	return pointers.Raw[Variant](raw).Copy()
}

func (a Array) SetIndex(index int64, value Variant) {
	raw, _ := pointers.End(value.Copy())
	gdextension.Host.Array.Set(pointers.Get(a), int(index), raw)
}

func (a Array) Free() {
	if ptr, ok := pointers.End(a); ok {
		noescape.Free(gdextension.TypeArray, &ptr)
	}
}

func (a Array) Iter() iter.Seq2[int64, Variant] {
	return func(yield func(int64, Variant) bool) {
		anchor, array := anchorTracked(a)
		defer runtime.KeepAlive(anchor)
		size := array.Size()
		for i := int64(0); i < size; i++ {
			if !yield(i, array.Index(i)) {
				break
			}
		}
	}
}

func NewArray() Array {
	return pointers.New[Array](noescape.Make[gdextension.Array](builtin.creation.Array[0], 0, nil))
}

func ArrayAs[S []T, T any](array Array) []T {
	anchor, array := anchorTracked(array)
	defer runtime.KeepAlive(anchor)
	var size = int(array.Size())
	var result = make([]T, size)
	for i := 0; i < size; i++ {
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

func NewArrayProxy[T any]() (ArrayProxy[T], complex128) {
	return WrapArray[T](NewArray())
}

// ArrayProxy is engine-backed proxy state for an array, the anchor carries
// the GC lifetime of goroutine-created arrays — see anchors.go.
type ArrayProxy[T any] struct {
	anchor *Array
}

func (proxy ArrayProxy[T]) Any(state complex128) ArrayVariant.Any {
	return ArrayVariant.Through(ArrayProxy[VariantPkg.Any]{anchor: proxy.anchor}, state)
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
