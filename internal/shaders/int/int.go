// Package int provides GPU operations on signed integer values.
package int

import (
	"graphics.gd/shaders/internal/gpu"
)

// X is a signed integer on the GPU (represented by an unspecified number of bits).
type X gpu.Int

// New creates a new [X] value from an integer.
func New[T gpu.AnyInt](x T) X { return gpu.NewInt(x) }

func Add[A, B gpu.AnyInt](a A, b B) X { //glsl:+(int,int)
	return gpu.NewIntExpression(gpu.Op(gpu.NewInt(a), "+", gpu.NewInt(b)))
}
func Sub[A, B gpu.AnyInt](a A, b B) X { //glsl:-(int,int)
	return gpu.NewIntExpression(gpu.Op(gpu.NewInt(a), "-", gpu.NewInt(b)))
}
func Mul[A, B gpu.AnyInt](a A, b B) X { //glsl:*(int,int)
	return gpu.NewIntExpression(gpu.Op(gpu.NewInt(a), "*", gpu.NewInt(b)))
}
func Div[A, B gpu.AnyInt](a A, b B) X { //glsl:/(int,int)
	return gpu.NewIntExpression(gpu.Op(gpu.NewInt(a), "/", gpu.NewInt(b)))
}
func Mod[A, B gpu.AnyInt](a A, b B) X {
	return gpu.NewIntExpression(gpu.Op(gpu.NewInt(a), "%", gpu.NewInt(b)))
}
func Neg[A gpu.AnyInt](a A) X { //glsl:-(int)
	return gpu.NewIntExpression(gpu.Op(nil, "-", gpu.NewInt(a)))
}

func Eq[A, B gpu.AnyInt](a A, b B) gpu.Bool {
	return gpu.NewBoolExpression(gpu.Op(gpu.NewInt(a), "==", gpu.NewInt(b)))
}
func Neq[A, B gpu.AnyInt](a A, b B) gpu.Bool {
	return gpu.NewBoolExpression(gpu.Op(gpu.NewInt(a), "!=", gpu.NewInt(b)))
}
func Gt[A, B gpu.AnyInt](a A, b B) gpu.Bool {
	return gpu.NewBoolExpression(gpu.Op(gpu.NewInt(a), ">", gpu.NewInt(b)))
}
func Gte[A, B gpu.AnyInt](a A, b B) gpu.Bool {
	return gpu.NewBoolExpression(gpu.Op(gpu.NewInt(a), ">=", gpu.NewInt(b)))
}
func Lt[A, B gpu.AnyInt](a A, b B) gpu.Bool {
	return gpu.NewBoolExpression(gpu.Op(gpu.NewInt(a), "<", gpu.NewInt(b)))
}
func Lte[A, B gpu.AnyInt](a A, b B) gpu.Bool {
	return gpu.NewBoolExpression(gpu.Op(gpu.NewInt(a), "<=", gpu.NewInt(b)))
}

func Abs[A gpu.AnyInt](a A) X { //glsl:abs(int)int
	return gpu.NewIntExpression(gpu.Fn("abs", gpu.NewInt(a)))
}
func Sign[A gpu.AnyInt](a A) X { //glsl:sign(int)int
	return gpu.NewIntExpression(gpu.Fn("sign", gpu.NewInt(a)))
}
func Min[A, B gpu.AnyInt](a A, b B) X { //glsl:min(int,int)int
	return gpu.NewIntExpression(gpu.Fn("min", gpu.NewInt(a), gpu.NewInt(b)))
}
func Max[A, B gpu.AnyInt](a A, b B) X { //glsl:max(int,int)int
	return gpu.NewIntExpression(gpu.Fn("max", gpu.NewInt(a), gpu.NewInt(b)))
}
func Clamp[A, B, C gpu.AnyInt](a A, b B, c C) X { //glsl:clamp(int,int,int)int
	return gpu.NewIntExpression(gpu.Fn("clamp", gpu.NewInt(a), gpu.NewInt(b), gpu.NewInt(c)))
}
func BitsToFloat[A gpu.AnyInt](a A) gpu.Float { //glsl:intBitsToFloat(int)vec
	return gpu.NewFloatExpression(gpu.Fn("intBitsToFloat", gpu.NewInt(a)))
}
