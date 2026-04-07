package gdextension

import gdunsafe "graphics.gd"

// Call a static method on a variant type.
func (variant VariantType) Call(method StringName, args ...Variant) (Variant, error) {
	converted := make([]gdunsafe.Variant, len(args))
	for i, a := range args {
		converted[i] = gdunsafe.Variant(a)
	}
	raw, callErr := gdunsafe.VariantType(variant).StaticCall(gdunsafe.StringName(method[0]), converted...)
	return Variant(raw), callErr
}

// New calls the variant constructor with the given arguments and returns the
// result as a variant.
func (variant VariantType) New(args ...Variant) (Variant, error) {
	converted := make([]gdunsafe.Variant, len(args))
	for i, a := range args {
		converted[i] = gdunsafe.Variant(a)
	}
	raw, callErr := gdunsafe.VariantType(variant).Make(converted...)
	return Variant(raw), callErr
}
