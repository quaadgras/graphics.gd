package gd

import (
	"fmt"
	"iter"
	"reflect"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	VariantPkg "graphics.gd/variant"
	DictionaryType "graphics.gd/variant/Dictionary"
)

func (d Dictionary) Index(key Variant) Variant {
	var raw [3]uint64
	gdextension.Host.Dictionaries.Get(pointers.Get(d), pointers.Get(key), gdextension.CallReturns[gdextension.Variant](&raw))
	return pointers.Raw[Variant](raw).Copy()
}

func (d Dictionary) SetIndex(key Variant, value Variant) {
	gdextension.Host.Dictionaries.Set(pointers.Get(d), pointers.Get(key), pointers.Cut(value.Copy(), true))
}

func (d Dictionary) Free() {
	if ptr, ok := pointers.End(d); ok {
		gdextension.Free(gdextension.TypeDictionary, &ptr)
	}
}

func NewDictionary() Dictionary {
	return pointers.New[Dictionary](gdextension.Make[gdextension.Dictionary](Global.typeset.creation.Dictionary[0], 0, nil))
}

func InternalDictionary[K comparable, V any](dict DictionaryType.Map[K, V]) Dictionary {
	_, state := DictionaryType.As(dict, NewDictionaryProxy[K, V])
	return pointers.Load[Dictionary](state)
}

func DictionaryFromMap[V any](val V) DictionaryType.Any {
	converted := NewVariant(val).Interface()
	if converted == nil {
		return DictionaryType.New[VariantPkg.Any, VariantPkg.Any]()
	}
	return converted.(DictionaryType.Any)
}

func NewDictionaryProxy[K comparable, V any]() (DictionaryProxy[K, V], complex128) {
	var dict = NewDictionary()
	var pack = pointers.Pack(dict)
	return DictionaryProxy[K, V]{}, pack
}

func DictionaryAs[T any](dictionary DictionaryType.Any) T {
	_, state := DictionaryType.As(dictionary, NewDictionaryProxy[VariantPkg.Any, VariantPkg.Any])
	dict := pointers.Load[Dictionary](state)
	result, err := ConvertToDesiredGoType(dict, reflect.TypeFor[T]())
	if err != nil {
		panic(fmt.Sprintf("could not convert dictionary to desired go type: %v", err))
	}
	return result.Interface().(T)
}

type DictionaryProxy[K comparable, V any] struct{}

func (DictionaryProxy[K, V]) Index(state complex128, key K) V {
	dict := pointers.Load[Dictionary](state)
	val := dict.Index(NewVariant(key))
	value, err := convertVariantToDesiredGoType(val, reflect.TypeFor[V]())
	if err != nil {
		panic(fmt.Sprintf("could not convert variant to desired go type: %v", err))
	}
	return value.Interface().(V)
}
func (DictionaryProxy[K, V]) Has(state complex128, key K) bool {
	return bool(pointers.Load[Dictionary](state).Has(NewVariant(key)))
}
func (DictionaryProxy[K, V]) Lookup(state complex128, key K) (V, bool) {
	k := NewVariant(key)
	ok := pointers.Load[Dictionary](state).Has(k)
	if !ok {
		return [1]V{}[0], false
	}
	value := pointers.Load[Dictionary](state).Index(k)
	result, err := convertVariantToDesiredGoType(value, reflect.TypeFor[V]())
	if err != nil {
		panic(fmt.Sprintf("could not convert variant to desired go type: %v", err))
	}
	return result.Interface().(V), true
}
func (DictionaryProxy[K, V]) SetIndex(state complex128, key K, val V) {
	pointers.Load[Dictionary](state).SetIndex(NewVariant(key), NewVariant(val))
}
func (DictionaryProxy[K, V]) Clear(state complex128) {
	pointers.Load[Dictionary](state).Clear()
}
func (DictionaryProxy[K, V]) Sort(state complex128, less func(K, K) bool) {
	pointers.Load[Dictionary](state).Sort()
}
func (DictionaryProxy[K, V]) Iter(state complex128) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		dict := pointers.Load[Dictionary](state)
		keys := dict.Keys()
		for _, key := range keys.Iter() {
			value := dict.Index(key)
			k, err := convertVariantToDesiredGoType(key, reflect.TypeFor[K]())
			if err != nil {
				panic(fmt.Sprintf("could not convert variant to desired go type: %v", err))
			}
			v, err := convertVariantToDesiredGoType(value, reflect.TypeFor[V]())
			if err != nil {
				panic(fmt.Sprintf("could not convert variant to desired go type: %v", err))
			}
			if !yield(k.Interface().(K), v.Interface().(V)) {
				break
			}
		}
	}
}
func (DictionaryProxy[K, V]) Erase(state complex128, key K) bool {
	return bool(pointers.Load[Dictionary](state).Erase(NewVariant(key)))
}
func (DictionaryProxy[K, V]) Hash(state complex128) uint32 {
	return uint32(pointers.Load[Dictionary](state).Hash())
}
func (DictionaryProxy[K, V]) Len(state complex128) int {
	return int(pointers.Load[Dictionary](state).Size())
}
func (DictionaryProxy[K, V]) IsReadOnly(state complex128) bool {
	return bool(pointers.Load[Dictionary](state).IsReadOnly())
}
func (DictionaryProxy[K, V]) MakeReadOnly(state complex128) {
	pointers.Load[Dictionary](state).MakeReadOnly()
}
func (DictionaryProxy[K, V]) Any(state complex128) DictionaryType.Any {
	return DictionaryType.Through(DictionaryProxy[VariantPkg.Any, VariantPkg.Any]{}, state)
}
