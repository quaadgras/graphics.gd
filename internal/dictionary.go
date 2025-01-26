package gd

import (
	"fmt"
	"iter"
	"reflect"

	"graphics.gd/internal/callframe"
	"graphics.gd/internal/pointers"
	VariantPkg "graphics.gd/variant"
	DictionaryType "graphics.gd/variant/Dictionary"
)

func (d Dictionary) Index(key Variant) Variant {
	return Global.Dictionary.Index(d, key)
}

func (d Dictionary) SetIndex(key Variant, value Variant) {
	Global.Dictionary.SetIndex(d, key, value)
}

func (d Dictionary) Free() {
	if ptr, ok := pointers.End(d); ok {
		var frame = callframe.New()
		Global.typeset.destruct.Dictionary(callframe.Arg(frame, ptr).Addr())
		frame.Free()
	}
}

func NewDictionary() Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]EnginePointer](frame)
	Global.typeset.creation.Dictionary[0](r_ret.Addr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[Dictionary](raw)
}

func InternalDictionary[K comparable, V any](dict DictionaryType.Map[K, V]) Dictionary {
	_, state := DictionaryType.As(dict, NewDictionaryProxy[K, V])
	return pointers.Load[Dictionary](state)
}

func DictionaryFromMap[K comparable, V any](m map[K]V) DictionaryType.Any {
	var dict = DictionaryType.Through(NewDictionaryProxy[VariantPkg.Any, VariantPkg.Any]())
	for k, value := range m {
		dict.SetIndex(VariantPkg.New(k), VariantPkg.New(value))
	}
	return dict
}

func NewDictionaryProxy[K comparable, V any]() (DictionaryProxy[K, V], complex128) {
	var dict = NewDictionary()
	var pack = pointers.Pack(dict)
	return DictionaryProxy[K, V]{}, pack
}

func DictionaryAs[M map[K]V, K comparable, V any](dictionary DictionaryType.Any) M {
	_, state := DictionaryType.As(dictionary, NewDictionaryProxy[VariantPkg.Any, VariantPkg.Any])
	dict := pointers.Load[Dictionary](state)
	var result = make(map[K]V)
	for _, key := range dict.Keys().Iter() {
		result[VariantAs[K](key)] = VariantAs[V](dict.Index(key))
	}
	return result
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
