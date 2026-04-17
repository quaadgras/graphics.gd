// Package uint provides GPU operations on unsigned integer values.
package uint

import (
	"graphics.gd/shaders/internal/gpu"
)

// X is a unsigned integer on the GPU (represented by an unspecified number of bits).
type X gpu.Uint

// New creates a new [X] value from an integer.
func New[T gpu.AnyUint](x T) X { return gpu.NewUint(x) }

func Add[A, B gpu.AnyUint](a A, b B) X { //glsl:+(uint,uint)
	return gpu.NewUintExpression(gpu.Op(gpu.NewUint(a), "+", gpu.NewUint(b)))
}
func Sub[A, B gpu.AnyUint](a A, b B) X { //glsl:-(uint,uint)
	return gpu.NewUintExpression(gpu.Op(gpu.NewUint(a), "-", gpu.NewUint(b)))
}
func Mul[A, B gpu.AnyUint](a A, b B) X { //glsl:*(uint,uint)
	return gpu.NewUintExpression(gpu.Op(gpu.NewUint(a), "*", gpu.NewUint(b)))
}
func Div[A, B gpu.AnyUint](a A, b B) X { //glsl:/(uint,uint)
	return gpu.NewUintExpression(gpu.Op(gpu.NewUint(a), "/", gpu.NewUint(b)))
}
func Mod[A, B gpu.AnyUint](a A, b B) X {
	return gpu.NewUintExpression(gpu.Op(gpu.NewUint(a), "%", gpu.NewUint(b)))
}

func Eq[A, B gpu.AnyUint](a A, b B) gpu.Bool {
	return gpu.NewBoolExpression(gpu.Op(gpu.NewUint(a), "==", gpu.NewUint(b)))
}
func Neq[A, B gpu.AnyUint](a A, b B) gpu.Bool {
	return gpu.NewBoolExpression(gpu.Op(gpu.NewUint(a), "!=", gpu.NewUint(b)))
}
func Gt[A, B gpu.AnyUint](a A, b B) gpu.Bool {
	return gpu.NewBoolExpression(gpu.Op(gpu.NewUint(a), ">", gpu.NewUint(b)))
}
func Gte[A, B gpu.AnyUint](a A, b B) gpu.Bool {
	return gpu.NewBoolExpression(gpu.Op(gpu.NewUint(a), ">=", gpu.NewUint(b)))
}
func Lt[A, B gpu.AnyUint](a A, b B) gpu.Bool {
	return gpu.NewBoolExpression(gpu.Op(gpu.NewUint(a), "<", gpu.NewUint(b)))
}
func Lte[A, B gpu.AnyUint](a A, b B) gpu.Bool {
	return gpu.NewBoolExpression(gpu.Op(gpu.NewUint(a), "<=", gpu.NewUint(b)))
}

func Abs[A gpu.AnyUint](a A) X {
	return gpu.NewUintExpression(gpu.Fn("abs", gpu.NewUint(a)))
}
func Min[A, B gpu.AnyUint](a A, b B) X { //glsl:min(uint,uint)uint
	return gpu.NewUintExpression(gpu.Fn("min", gpu.NewUint(a), gpu.NewUint(b)))
}
func Max[A, B gpu.AnyUint](a A, b B) X { //glsl:max(uint,uint)uint
	return gpu.NewUintExpression(gpu.Fn("max", gpu.NewUint(a), gpu.NewUint(b)))
}
func Clamp[A, B, C gpu.AnyUint](a A, b B, c C) X { //glsl:clamp(uint,uint,uint)uint
	return gpu.NewUintExpression(gpu.Fn("clamp", gpu.NewUint(a), gpu.NewUint(b), gpu.NewUint(c)))
}
func BitsToFloat[A gpu.AnyUint](a A) gpu.Float { //glsl:uintBitsToFloat(uint)vec
	return gpu.NewFloatExpression(gpu.Fn("uintBitsToFloat", gpu.NewUint(a)))
}

func UnpackHalf2x16(a X) gpu.Vec2 { //glsl:unpackHalf2x16(uint)vec2
	return gpu.NewVec2Expression(gpu.Fn("unpackHalf2x16", a))
}
func UnpackSnorm2x16(a X) gpu.Vec2 { //glsl:unpackSnorm2x16(uint)vec2
	return gpu.NewVec2Expression(gpu.Fn("unpackSnorm2x16", a))
}
func UnpackUnorm2x16(a X) gpu.Vec2 { //glsl:unpackUnorm2x16(uint)vec2
	return gpu.NewVec2Expression(gpu.Fn("unpackUnorm2x16", a))
}
