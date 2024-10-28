//go:build !generate

package gd

func (variant Variant) evaluateIntoBool(op Operator, other Variant) bool {
	val, ok := Global.Variants.Evaluate(op, variant, other)
	if !ok {
		panic("variants not comparable")
	}
	var ret = variantAsValueType[bool](val, TypeBool)
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

func (variant Variant) Add(other Variant) Variant {
	var val, ok = Global.Variants.Evaluate(Add, variant, other)
	if !ok {
		panic("variants not addable")
	}
	return val
}

func (variant Variant) Subtract(other Variant) Variant {
	var val, ok = Global.Variants.Evaluate(Subtract, variant, other)
	if !ok {
		panic("variants not subtractable")
	}
	return val
}

func (variant Variant) Multiply(other Variant) Variant {
	var val, ok = Global.Variants.Evaluate(Multiply, variant, other)
	if !ok {
		panic("variants not multipliable")
	}
	return val
}

func (variant Variant) Divide(other Variant) Variant {
	var val, ok = Global.Variants.Evaluate(Divide, variant, other)
	if !ok {
		panic("variants not divisible")
	}
	return val
}

func (variant Variant) Negate() Variant {
	var val, ok = Global.Variants.Evaluate(Negate, variant, Global.Variants.NewNil())
	if !ok {
		panic("variants not negatable")
	}
	return val
}

func (variant Variant) Positive() Variant {
	var val, ok = Global.Variants.Evaluate(Positive, variant, Global.Variants.NewNil())
	if !ok {
		panic("variants not positivable")
	}
	return val
}

func (variant Variant) Modulo(other Variant) Variant {
	var val, ok = Global.Variants.Evaluate(Module, variant, other)
	if !ok {
		panic("variants not modulable")
	}
	return val
}

func (variant Variant) Power(other Variant) Variant {
	var val, ok = Global.Variants.Evaluate(Power, variant, other)
	if !ok {
		panic("variants not powerable")
	}
	return val
}

func (variant Variant) ShiftLeft(other Variant) Variant {
	var val, ok = Global.Variants.Evaluate(ShiftLeft, variant, other)
	if !ok {
		panic("variants not shift leftable")
	}
	return val
}

func (variant Variant) ShiftRight(other Variant) Variant {
	var val, ok = Global.Variants.Evaluate(ShiftRight, variant, other)
	if !ok {
		panic("variants not shift rightable")
	}
	return val
}

func (variant Variant) BitAnd(other Variant) Variant {
	var val, ok = Global.Variants.Evaluate(BitAnd, variant, other)
	if !ok {
		panic("variants not bit andable")
	}
	return val
}

func (variant Variant) BitOr(other Variant) Variant {
	var val, ok = Global.Variants.Evaluate(BitOr, variant, other)
	if !ok {
		panic("variants not bit orable")
	}
	return val
}

func (variant Variant) BitXor(other Variant) Variant {
	var val, ok = Global.Variants.Evaluate(BitXor, variant, other)
	if !ok {
		panic("variants not bit xorable")
	}
	return val
}

func (variant Variant) BitNegate() Variant {
	var val, ok = Global.Variants.Evaluate(BitNegate, variant, Global.Variants.NewNil())
	if !ok {
		panic("variants not bit negatable")
	}
	return val
}

func (variant Variant) LogicalAnd(other Variant) Variant {
	var val, ok = Global.Variants.Evaluate(LogicalAnd, variant, other)
	if !ok {
		panic("variants not logical andable")
	}
	return val
}

func (variant Variant) LogicalOr(other Variant) Variant {
	var val, ok = Global.Variants.Evaluate(LogicalOr, variant, other)
	if !ok {
		panic("variants not logical orable")
	}
	return val
}

func (variant Variant) LogicalXor(other Variant) Variant {
	var val, ok = Global.Variants.Evaluate(LogicalXor, variant, other)
	if !ok {
		panic("variants not logical xorable")
	}
	return val
}

func (variant Variant) LogicalNegate() Variant {
	var val, ok = Global.Variants.Evaluate(LogicalNegate, variant, Global.Variants.NewNil())
	if !ok {
		panic("variants not logical negatable")
	}
	return val
}

func (variant Variant) In(other Variant) Variant {
	var val, ok = Global.Variants.Evaluate(In, variant, other)
	if !ok {
		panic("variants not inable")
	}
	return val
}
