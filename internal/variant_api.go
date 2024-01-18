//go:build !generate

package gd

// Type returns the variant's type, similar to [reflect.Kind] but for a variant
// value.
func (variant Variant) Type() VariantType {
	return variant.API.Variants.GetType(variant)
}

// Get returns the value specified by the given key variant and a boolean
// indiciating whether the get operation was valid.
func (variant Variant) Get(ctx Context, key Variant) (val Variant, ok bool) {
	return variant.API.Variants.Get(ctx, variant, key)
}

// Set sets the value specified by the given key variant to the given value
// variant. Returns true if the set operation was valid.
func (variant Variant) Set(key, val Variant) bool {
	return variant.API.Variants.Set(variant, key, val)
}

// Call calls a method on the variant dynamically.
func (variant Variant) Call(ctx Context, method StringName, args ...Variant) (Variant, error) {
	return variant.API.Variants.Call(ctx, variant, method, args...)
}

// Call a static method on a variant type.
func (variant VariantType) Call(ctx Context, method StringName, args ...Variant) (Variant, error) {
	return ctx.API().Variants.CallStatic(ctx, variant, method, args...)
}

// New calls the variant constructor with the given arguments and returns the
// result as a variant.
func (variant VariantType) New(ctx Context, args ...Variant) (Variant, error) {
	return ctx.API().Variants.Construct(ctx, variant, args...)
}

// Iterator returns an iterator for the variant.
func (variant Variant) Iterator(ctx Context) Iterator {
	iter, ok := ctx.API().Variants.IteratorInitialize(ctx, variant)
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
	return variant.API.Variants.Hash(variant)
}

// RecursiveHash returns the hash value of the variant recursively.
func (variant Variant) RecursiveHash(count Int) Int {
	return variant.API.Variants.RecursiveHash(variant, count)
}
