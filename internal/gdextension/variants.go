package gdextension

func (op VariantOperator) Eval(a, b Variant) (Variant, bool) {
	var raw Variant
	ok := Host.Variants.Eval(op, a, b, CallReturns[Variant](&raw))
	return raw, ok
}

// Call a static method on a variant type.
func (variant VariantType) Call(method StringName, args ...Variant) (Variant, error) {
	var err CallError
	var raw Variant
	Host.Builtin.Types.Call(variant, method, CallReturns[Variant](&raw), len(args), CallAccepts[Variant](&args[0]), CallReturns[CallError](&err))
	return raw, err
}

// New calls the variant constructor with the given arguments and returns the
// result as a variant.
func (variant VariantType) New(args ...Variant) (Variant, error) {
	var err CallError
	var raw Variant
	Host.Builtin.Types.Make(variant, CallReturns[Variant](&raw), len(args), CallAccepts[Variant](&args[0]), CallReturns[CallError](&err))
	return raw, err
}
