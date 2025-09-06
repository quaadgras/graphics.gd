//go:build !generate

package gd

import (
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
)

// FIXME/TODO move this into gdextension package?
//

// Copy returns a copy of the variant that will belong to the provided context.
func (variant Variant) Copy() Variant {
	var raw [3]uint64
	gdextension.Host.Variants.Copy(pointers.Get(variant), gdextension.CallReturns[gdextension.Variant](&raw))
	return pointers.New[Variant](raw)
}

// Type returns the variant's type, similar to [reflect.Kind] but for a variant
// value.
func (variant Variant) Type() gdextension.VariantType {
	return gdextension.Host.Variants.Type(pointers.Get(variant))
}

// Get returns the value specified by the given key variant and a boolean
// indiciating whether the get operation was valid.
func (variant Variant) Get(key Variant) (val Variant, ok bool) {
	var raw [3]uint64
	ok = gdextension.Host.Variants.Get.Index(pointers.Get(variant), pointers.Get(key), gdextension.CallReturns[gdextension.Variant](&raw))
	return pointers.New[Variant](raw), ok
}

// Set sets the value specified by the given key variant to the given value
// variant. Returns true if the set operation was valid.
func (variant Variant) Set(key, val Variant) bool {
	return gdextension.Host.Variants.Set.Index(pointers.Get(variant), pointers.Get(key), pointers.Get(val))
}

// Call calls a method on the variant dynamically.
func (variant Variant) Call(method StringName, args ...Variant) (Variant, error) {
	var raw [3]uint64
	var converted []gdextension.Variant
	for i := range args {
		converted = append(converted, gdextension.Variant(pointers.Get(args[i])))
	}
	var err gdextension.CallError
	gdextension.Host.Variants.Call(pointers.Get(variant), pointers.Get(method), gdextension.CallReturns[gdextension.Variant](&raw), len(args), gdextension.CallAccepts[gdextension.Variant](&converted[0]), gdextension.CallReturns[gdextension.CallError](&err))
	return pointers.New[Variant](raw), err.Err()
}

// Iterator returns an iterator for the variant.
func (variant Variant) Iterator() Iterator {
	var err gdextension.CallError
	var raw gdextension.Iterator
	gdextension.Host.Iterators.Make(pointers.Get(variant), gdextension.CallReturns[gdextension.Iterator](&raw), gdextension.CallReturns[gdextension.CallError](&err))
	if err.Type != 0 {
		panic("failed to initialize iterator")
	}
	return Iterator{
		self: variant,
		iter: pointers.New[iterator](raw),
	}
}

// Hash returns the hash value of the variant.
func (variant Variant) Hash() Int {
	var hash int64
	gdextension.Host.Variants.Hash(pointers.Get(variant), gdextension.CallReturns[int64](&hash))
	return hash
}

// RecursiveHash returns the hash value of the variant recursively.
func (variant Variant) RecursiveHash(count Int) Int {
	var hash int64
	gdextension.Host.Variants.Deep.Hash(pointers.Get(variant), count, gdextension.CallReturns[int64](&hash))
	return hash
}
