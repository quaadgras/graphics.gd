//go:build !generate

package gd

import "runtime.link/mmm"

func (variant Variant) evaluateIntoBool(op Operator, other Variant) bool {
	var ctx Context
	ctx.godotContext = mmm.NewContextWith[API](nil, variant.API)
	val, ok := variant.API.Variants.Evaluate(ctx, op, variant, other)
	if !ok {
		panic("variants not comparable")
	}
	var ret = val.asBool()
	val.Free()
	return ret
}

func (variant Variant) EqualTo(other Variant) bool {
	return variant.evaluateIntoBool(Equal, other)
}

func (variant Variant) NotEqualTo(other Variant) bool {
	return variant.evaluateIntoBool(NotEqual, other)
}

func (variant Variant) LessThan(other Variant) bool {
	return variant.evaluateIntoBool(Less, other)
}

func (variant Variant) AtMost(other Variant) bool {
	return variant.evaluateIntoBool(LessEqual, other)
}

func (variant Variant) MoreThan(other Variant) bool {
	return variant.evaluateIntoBool(Greater, other)
}

func (variant Variant) AtLeast(other Variant) bool {
	return variant.evaluateIntoBool(GreaterEqual, other)
}

func (variant Variant) Add(ctx Context, other Variant) Variant {
	var val, ok = variant.API.Variants.Evaluate(ctx, Add, variant, other)
	if !ok {
		panic("variants not addable")
	}
	return val
}

func (variant Variant) Subtract(ctx Context, other Variant) Variant {
	var val, ok = variant.API.Variants.Evaluate(ctx, Subtract, variant, other)
	if !ok {
		panic("variants not subtractable")
	}
	return val
}

func (variant Variant) Multiply(ctx Context, other Variant) Variant {
	var val, ok = variant.API.Variants.Evaluate(ctx, Multiply, variant, other)
	if !ok {
		panic("variants not multipliable")
	}
	return val
}

func (variant Variant) Divide(ctx Context, other Variant) Variant {
	var val, ok = variant.API.Variants.Evaluate(ctx, Divide, variant, other)
	if !ok {
		panic("variants not divisible")
	}
	return val
}

func (variant Variant) Negate(ctx Context) Variant {
	var val, ok = variant.API.Variants.Evaluate(ctx, Negate, variant, variant.API.Variants.NewNil(ctx))
	if !ok {
		panic("variants not negatable")
	}
	return val
}

func (variant Variant) Positive(ctx Context) Variant {
	var val, ok = variant.API.Variants.Evaluate(ctx, Positive, variant, variant.API.Variants.NewNil(ctx))
	if !ok {
		panic("variants not positivable")
	}
	return val
}

func (variant Variant) Modulo(ctx Context, other Variant) Variant {
	var val, ok = variant.API.Variants.Evaluate(ctx, Module, variant, other)
	if !ok {
		panic("variants not modulable")
	}
	return val
}

func (variant Variant) Power(ctx Context, other Variant) Variant {
	var val, ok = variant.API.Variants.Evaluate(ctx, Power, variant, other)
	if !ok {
		panic("variants not powerable")
	}
	return val
}

func (variant Variant) ShiftLeft(ctx Context, other Variant) Variant {
	var val, ok = variant.API.Variants.Evaluate(ctx, ShiftLeft, variant, other)
	if !ok {
		panic("variants not shift leftable")
	}
	return val
}

func (variant Variant) ShiftRight(ctx Context, other Variant) Variant {
	var val, ok = variant.API.Variants.Evaluate(ctx, ShiftRight, variant, other)
	if !ok {
		panic("variants not shift rightable")
	}
	return val
}

func (variant Variant) BitAnd(ctx Context, other Variant) Variant {
	var val, ok = variant.API.Variants.Evaluate(ctx, BitAnd, variant, other)
	if !ok {
		panic("variants not bit andable")
	}
	return val
}

func (variant Variant) BitOr(ctx Context, other Variant) Variant {
	var val, ok = variant.API.Variants.Evaluate(ctx, BitOr, variant, other)
	if !ok {
		panic("variants not bit orable")
	}
	return val
}

func (variant Variant) BitXor(ctx Context, other Variant) Variant {
	var val, ok = variant.API.Variants.Evaluate(ctx, BitXor, variant, other)
	if !ok {
		panic("variants not bit xorable")
	}
	return val
}

func (variant Variant) BitNegate(ctx Context) Variant {
	var val, ok = variant.API.Variants.Evaluate(ctx, BitNegate, variant, variant.API.Variants.NewNil(ctx))
	if !ok {
		panic("variants not bit negatable")
	}
	return val
}

func (variant Variant) LogicalAnd(ctx Context, other Variant) Variant {
	var val, ok = variant.API.Variants.Evaluate(ctx, LogicalAnd, variant, other)
	if !ok {
		panic("variants not logical andable")
	}
	return val
}

func (variant Variant) LogicalOr(ctx Context, other Variant) Variant {
	var val, ok = variant.API.Variants.Evaluate(ctx, LogicalOr, variant, other)
	if !ok {
		panic("variants not logical orable")
	}
	return val
}

func (variant Variant) LogicalXor(ctx Context, other Variant) Variant {
	var val, ok = variant.API.Variants.Evaluate(ctx, LogicalXor, variant, other)
	if !ok {
		panic("variants not logical xorable")
	}
	return val
}

func (variant Variant) LogicalNegate(ctx Context) Variant {
	var val, ok = variant.API.Variants.Evaluate(ctx, LogicalNegate, variant, variant.API.Variants.NewNil(ctx))
	if !ok {
		panic("variants not logical negatable")
	}
	return val
}

func (variant Variant) In(ctx Context, other Variant) Variant {
	var val, ok = variant.API.Variants.Evaluate(ctx, In, variant, other)
	if !ok {
		panic("variants not inable")
	}
	return val
}
