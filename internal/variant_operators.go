//go:build !generate

package gd

import "runtime.link/mmm"

func (variant Variant) evaluateIntoBool(ctx Lifetime, op Operator, other Variant) bool {
	val, ok := mmm.API(variant).Variants.Evaluate(ctx, op, variant, other)
	if !ok {
		panic("variants not comparable")
	}
	var ret = variantAsValueType[bool](val, TypeBool)
	val.Free()
	return ret
}

func (variant Variant) EqualTo(ctx Lifetime, other Variant) bool {
	return variant.evaluateIntoBool(ctx, Equal, other)
}

func (variant Variant) NotEqualTo(ctx Lifetime, other Variant) bool {
	return variant.evaluateIntoBool(ctx, NotEqual, other)
}

func (variant Variant) LessThan(ctx Lifetime, other Variant) bool {
	return variant.evaluateIntoBool(ctx, Less, other)
}

func (variant Variant) AtMost(ctx Lifetime, other Variant) bool {
	return variant.evaluateIntoBool(ctx, LessEqual, other)
}

func (variant Variant) MoreThan(ctx Lifetime, other Variant) bool {
	return variant.evaluateIntoBool(ctx, Greater, other)
}

func (variant Variant) AtLeast(ctx Lifetime, other Variant) bool {
	return variant.evaluateIntoBool(ctx, GreaterEqual, other)
}

func (variant Variant) Add(ctx Lifetime, other Variant) Variant {
	var val, ok = mmm.API(variant).Variants.Evaluate(ctx, Add, variant, other)
	if !ok {
		panic("variants not addable")
	}
	return val
}

func (variant Variant) Subtract(ctx Lifetime, other Variant) Variant {
	var val, ok = mmm.API(variant).Variants.Evaluate(ctx, Subtract, variant, other)
	if !ok {
		panic("variants not subtractable")
	}
	return val
}

func (variant Variant) Multiply(ctx Lifetime, other Variant) Variant {
	var val, ok = mmm.API(variant).Variants.Evaluate(ctx, Multiply, variant, other)
	if !ok {
		panic("variants not multipliable")
	}
	return val
}

func (variant Variant) Divide(ctx Lifetime, other Variant) Variant {
	var val, ok = mmm.API(variant).Variants.Evaluate(ctx, Divide, variant, other)
	if !ok {
		panic("variants not divisible")
	}
	return val
}

func (variant Variant) Negate(ctx Lifetime) Variant {
	var val, ok = mmm.API(variant).Variants.Evaluate(ctx, Negate, variant, mmm.API(variant).Variants.NewNil(ctx))
	if !ok {
		panic("variants not negatable")
	}
	return val
}

func (variant Variant) Positive(ctx Lifetime) Variant {
	var val, ok = mmm.API(variant).Variants.Evaluate(ctx, Positive, variant, mmm.API(variant).Variants.NewNil(ctx))
	if !ok {
		panic("variants not positivable")
	}
	return val
}

func (variant Variant) Modulo(ctx Lifetime, other Variant) Variant {
	var val, ok = mmm.API(variant).Variants.Evaluate(ctx, Module, variant, other)
	if !ok {
		panic("variants not modulable")
	}
	return val
}

func (variant Variant) Power(ctx Lifetime, other Variant) Variant {
	var val, ok = mmm.API(variant).Variants.Evaluate(ctx, Power, variant, other)
	if !ok {
		panic("variants not powerable")
	}
	return val
}

func (variant Variant) ShiftLeft(ctx Lifetime, other Variant) Variant {
	var val, ok = mmm.API(variant).Variants.Evaluate(ctx, ShiftLeft, variant, other)
	if !ok {
		panic("variants not shift leftable")
	}
	return val
}

func (variant Variant) ShiftRight(ctx Lifetime, other Variant) Variant {
	var val, ok = mmm.API(variant).Variants.Evaluate(ctx, ShiftRight, variant, other)
	if !ok {
		panic("variants not shift rightable")
	}
	return val
}

func (variant Variant) BitAnd(ctx Lifetime, other Variant) Variant {
	var val, ok = mmm.API(variant).Variants.Evaluate(ctx, BitAnd, variant, other)
	if !ok {
		panic("variants not bit andable")
	}
	return val
}

func (variant Variant) BitOr(ctx Lifetime, other Variant) Variant {
	var val, ok = mmm.API(variant).Variants.Evaluate(ctx, BitOr, variant, other)
	if !ok {
		panic("variants not bit orable")
	}
	return val
}

func (variant Variant) BitXor(ctx Lifetime, other Variant) Variant {
	var val, ok = mmm.API(variant).Variants.Evaluate(ctx, BitXor, variant, other)
	if !ok {
		panic("variants not bit xorable")
	}
	return val
}

func (variant Variant) BitNegate(ctx Lifetime) Variant {
	var val, ok = mmm.API(variant).Variants.Evaluate(ctx, BitNegate, variant, mmm.API(variant).Variants.NewNil(ctx))
	if !ok {
		panic("variants not bit negatable")
	}
	return val
}

func (variant Variant) LogicalAnd(ctx Lifetime, other Variant) Variant {
	var val, ok = mmm.API(variant).Variants.Evaluate(ctx, LogicalAnd, variant, other)
	if !ok {
		panic("variants not logical andable")
	}
	return val
}

func (variant Variant) LogicalOr(ctx Lifetime, other Variant) Variant {
	var val, ok = mmm.API(variant).Variants.Evaluate(ctx, LogicalOr, variant, other)
	if !ok {
		panic("variants not logical orable")
	}
	return val
}

func (variant Variant) LogicalXor(ctx Lifetime, other Variant) Variant {
	var val, ok = mmm.API(variant).Variants.Evaluate(ctx, LogicalXor, variant, other)
	if !ok {
		panic("variants not logical xorable")
	}
	return val
}

func (variant Variant) LogicalNegate(ctx Lifetime) Variant {
	var val, ok = mmm.API(variant).Variants.Evaluate(ctx, LogicalNegate, variant, mmm.API(variant).Variants.NewNil(ctx))
	if !ok {
		panic("variants not logical negatable")
	}
	return val
}

func (variant Variant) In(ctx Lifetime, other Variant) Variant {
	var val, ok = mmm.API(variant).Variants.Evaluate(ctx, In, variant, other)
	if !ok {
		panic("variants not inable")
	}
	return val
}
