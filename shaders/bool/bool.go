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

// Select returns ifTrue when cond is true and ifFalse otherwise. Both values
// are computed, as with the ?: operator in the shading language, so this is
// for choosing between values, not for skipping work.
func Select[T gpu.Evaluator, C gpu.AnyBool](cond C, ifTrue, ifFalse T) T {
	return gpu.If(New(cond), ifTrue, ifFalse)
}

// Evaluate returns which way the condition goes, so that it can be used with a
// Go if statement:
//
//	color := rgb.New(0, 0, 0)
//	if float.Gt(density, 0.0).Evaluate() {
//		color = lighting(density) // only computed inside of clouds
//	} else {
//		color = background
//	}
//
// The pipeline function is run once for each way that the conditions it
// evaluates can go, and the values that differ between these runs are computed
// inside of an if statement on the GPU, so only the work of the branch taken is
// done. Pipeline functions must therefore do the same thing each time they run
// with the same outcomes. Every condition evaluated doubles the number of runs
// that follow it, so to choose between two values that are cheap to compute,
// use [Select] instead.
//
// Evaluate can only be called by pipeline functions and the bodies of loops,
// while the shader is being compiled.
func (x X) Evaluate() bool {
	if gpu.IsLiteral(x) {
		return x.X
	}
	if gpu.Decide == nil {
		panic("shaders: Evaluate can only be called while a shader is being compiled")
	}
	return gpu.Decide(gpu.NewBool(x))
}
