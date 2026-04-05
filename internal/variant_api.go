//go:build !generate

package gd

import (
	"unsafe"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
)

// Copy returns a copy of the variant that will belong to the provided context.
func (variant Variant) Copy() Variant {
	return pointers.New[Variant](gdunsafe.Variant(pointers.Get(variant)).Copy())
}

// Type returns the variant's type, similar to [reflect.Kind] but for a variant
// value.
func (variant Variant) Type() gdextension.VariantType {
	return gdextension.VariantType(gdunsafe.Variant(pointers.Get(variant)).Type())
}

// Get returns the value specified by the given key variant and a boolean
// indiciating whether the get operation was valid.
func (variant Variant) Get(key Variant) (val Variant, ok bool) {
	raw, ok := gdunsafe.Variant(pointers.Get(variant)).GetIndex(gdunsafe.Variant(pointers.Get(key)))
	return pointers.New[Variant](raw), ok
}

// Set sets the value specified by the given key variant to the given value
// variant. Returns true if the set operation was valid.
func (variant Variant) Set(key, val Variant) bool {
	return gdunsafe.Variant(pointers.Get(variant)).SetIndex(gdunsafe.Variant(pointers.Get(key)), gdunsafe.Variant(pointers.Get(val)))
}

// Call calls a method on the variant dynamically.
func (variant Variant) Call(method StringName, args ...Variant) (Variant, error) {
	var converted []gdunsafe.Variant
	for i := range args {
		converted = append(converted, gdunsafe.Variant(pointers.Get(args[i])))
	}
	raw, err := gdunsafe.Variant(pointers.Get(variant)).VariantCall(gdunsafe.StringName(pointers.Get(method)[0]), converted...)
	return pointers.New[Variant](raw), err.Err()
}

// Iterator returns an iterator for the variant.
func (variant Variant) Iterator() Iterator {
	var err gdextension.CallError
	var raw gdextension.Iterator
	gdunsafe.Variant(pointers.Get(variant)).IteratorMake(unsafe.Pointer(&raw), unsafe.Pointer(&err))
	if err.Type != 0 {
		panic("failed to initialize iterator")
	}
	return Iterator{
		self: variant,
		iter: pointers.New[iterator](raw),
	}
}

// Hash returns the hash value of the variant.
func (variant Variant) Hash() Int { return Int(gdunsafe.Variant(pointers.Get(variant)).Hash()) }

// RecursiveHash returns the hash value of the variant recursively.
func (variant Variant) RecursiveHash(count Int) Int {
	return Int(gdunsafe.Variant(pointers.Get(variant)).DeepHash(gdunsafe.Int(count)))
}

// Eval evaluates a binary operator between two variants.
func VariantEval(op gdextension.VariantOperator, a, b gdextension.Variant) (gdextension.Variant, bool) {
	raw, ok := gdunsafe.VariantEval(gdunsafe.VariantOperator(op), gdunsafe.Variant(a), gdunsafe.Variant(b))
	return raw, ok
}
