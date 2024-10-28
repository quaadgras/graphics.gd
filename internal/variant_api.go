//go:build !generate

package gd

// Copy returns a copy of the variant that will belong to the provided context.
func (variant Variant) Copy() Variant {
	return Global.Variants.NewCopy(variant)
}

// Type returns the variant's type, similar to [reflect.Kind] but for a variant
// value.
func (variant Variant) Type() VariantType {
	return Global.Variants.GetType(variant)
}

// Get returns the value specified by the given key variant and a boolean
// indiciating whether the get operation was valid.
func (variant Variant) Get(key Variant) (val Variant, ok bool) {
	return Global.Variants.Get(variant, key)
}

// Set sets the value specified by the given key variant to the given value
// variant. Returns true if the set operation was valid.
func (variant Variant) Set(key, val Variant) bool {
	return Global.Variants.Set(variant, key, val)
}

// Call calls a method on the variant dynamically.
func (variant Variant) Call(method StringName, args ...Variant) (Variant, error) {
	return Global.Variants.Call(variant, method, args...)
}

// Call a static method on a variant type.
func (variant VariantType) Call(method StringName, args ...Variant) (Variant, error) {
	return Global.Variants.CallStatic(variant, method, args...)
}

// New calls the variant constructor with the given arguments and returns the
// result as a variant.
func (variant VariantType) New(args ...Variant) (Variant, error) {
	return Global.Variants.Construct(variant, args...)
}

// Iterator returns an iterator for the variant.
func (variant Variant) Iterator() Iterator {
	iter, ok := Global.Variants.IteratorInitialize(variant)
	if !ok {
		panic("failed to initialize iterator")
	}
	return Iterator{
		self: variant,
		iter: iter,
	}
}

// Hash returns the hash value of the variant.
func (variant Variant) Hash() Int {
	return Global.Variants.Hash(variant)
}

// RecursiveHash returns the hash value of the variant recursively.
func (variant Variant) RecursiveHash(count Int) Int {
	return Global.Variants.RecursiveHash(variant, count)
}
