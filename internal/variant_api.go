//go:build !generate

package gd

import (
	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant"
)

// Copy returns a copy of the variant that will belong to the provided context.
func (v Variant) Copy() Variant {
	return pointers.New[Variant](gdunsafe.Variant(pointers.Get(v)).Copy(false))
}

// Type returns the variant's type, similar to [reflect.Kind] but for a variant
// value.
func (v Variant) Type() variant.Type {
	return variant.Type(gdunsafe.Variant(pointers.Get(v)).Type())
}

// Get returns the value specified by the given key variant and a boolean
// indiciating whether the get operation was valid.
func (v Variant) Get(key Variant) (val Variant, ok bool) {
	raw, ok := gdunsafe.Variant(pointers.Get(v)).Lookup(gdunsafe.Variant(pointers.Get(key)))
	return pointers.New[Variant](raw), ok
}

// Set sets the value specified by the given key variant to the given value
// variant. Returns true if the set operation was valid.
func (v Variant) Set(key, val Variant) bool {
	return gdunsafe.Variant(pointers.Get(v)).Insert(gdunsafe.Variant(pointers.Get(key)), gdunsafe.Variant(pointers.Get(val)))
}

// Call calls a method on the variant dynamically.
func (v Variant) Call(method StringName, args ...Variant) (Variant, error) {
	var converted []gdunsafe.Variant
	for i := range args {
		converted = append(converted, gdunsafe.Variant(pointers.Get(args[i])))
	}
	raw, err := gdunsafe.Variant(pointers.Get(v)).Call(gdunsafe.StringName(pointers.Get(method)[0]), converted...)
	if err != (gdunsafe.Error{}) {
		return pointers.New[Variant](raw), err
	}
	return pointers.New[Variant](raw), nil
}

// Iterator returns an iterator for the variant.
func (v Variant) Iterator() Iterator {
	raw, err := gdunsafe.Variant(pointers.Get(v)).Iterator()
	if err != (gdextension.CallError{}) {
		panic("failed to initialize iterator")
	}
	return Iterator{
		self: v,
		iter: pointers.New[iterator](raw),
	}
}

// Hash returns the hash value of the variant.
func (v Variant) Hash() Int { return Int(gdunsafe.Variant(pointers.Get(v)).Hash(0)) }

// RecursiveHash returns the hash value of the variant recursively.
func (v Variant) RecursiveHash(count Int) Int {
	return Int(gdunsafe.Variant(pointers.Get(v)).Hash(int64(count)))
}

// Eval evaluates a binary operator between two variants.
func VariantEval(op gdextension.VariantOperator, a, b gdextension.Variant) (gdextension.Variant, bool) {
	raw, ok := gdunsafe.VariantOperator(op).Evaluate(a, b)
	return raw, ok
}
