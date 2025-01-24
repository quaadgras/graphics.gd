// Package bool provides GPU operations on boolean values.
package bool

import "graphics.gd/shaders/internal/gpu"

// X is a boolean on the GPU, can only contain true or false.
type X gpu.Bool

// New creates a new [X] value from a boolean.
func New[T gpu.AnyBool](x T) X {
	return gpu.NewBool(x)
}

func Not[A gpu.AnyBool](a A) X         { return gpu.NewBoolExpression(gpu.Op(nil, "!", New(a))) }
func And[A, B gpu.AnyBool](a A, b B) X { return gpu.NewBoolExpression(gpu.Op(New(a), "&&", New(b))) }
func Or[A, B gpu.AnyBool](a A, b B) X  { return gpu.NewBoolExpression(gpu.Op(New(a), "||", New(b))) }
func Eq[A, B gpu.AnyBool](a A, b B) X  { return gpu.NewBoolExpression(gpu.Op(New(a), "==", New(b))) }
func Neq[A, B gpu.AnyBool](a A, b B) X { return gpu.NewBoolExpression(gpu.Op(New(a), "!=", New(b))) }

func Mix[A, B gpu.AnyFloat, T gpu.AnyBool](a A, b B, t T) gpu.Float { //glsl:mix(float,float,bvec)float
	return gpu.NewFloatExpression(gpu.New(gpu.Ternary{
		If: New(t),
		A:  gpu.NewFloat(b),
		B:  gpu.NewFloat(a),
	}))
}
