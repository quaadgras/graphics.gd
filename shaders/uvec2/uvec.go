// Package uvec2 provides GPU operations on two-component unsigned integer vectors.
package uvec2

import (
	"reflect"

	"graphics.gd/shaders/internal/gpu"
)

// XY is a two-component vector of unsigned integers on the GPU.
type XY gpu.Vec2u

// New creates a new [XY] value from two integers.
func New[X, Y gpu.AnyUint](x X, y Y) XY { return gpu.NewVec2u(x, y) }

func Add[T gpu.AnyUint | XY](a XY, b T) XY { return gpu.NewVec2uExpression(gpu.Op(a, "+", either(b))) } //glsl:+(uvec2,uvec2)
func Sub[T gpu.AnyUint | XY](a XY, b T) XY { return gpu.NewVec2uExpression(gpu.Op(a, "-", either(b))) } //glsl:-(uvec2,uvec2)
func Mul[T gpu.AnyUint | XY](a XY, b T) XY { return gpu.NewVec2uExpression(gpu.Op(a, "*", either(b))) } //glsl:*(uvec2,uvec2)
func Div[T gpu.AnyUint | XY](a XY, b T) XY { return gpu.NewVec2uExpression(gpu.Op(a, "/", either(b))) } //glsl:/(uvec2,uvec2)

func BitsToFloat(a XY) gpu.Vec2 { return gpu.NewVec2Expression(gpu.Fn("uintBitsToFloat", a)) } //glsl:uintBitsToFloat(uvec2)vec

func either[T gpu.AnyUint | XY](v T) gpu.Evaluator {
	rvalue := reflect.ValueOf(v)
	switch {
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(XY{})):
		return rvalue.Convert(reflect.TypeOf(XY{})).Interface().(XY)
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(gpu.Uint{})):
		return gpu.NewUint(rvalue.Convert(reflect.TypeOf(gpu.Uint{})).Interface().(gpu.Uint))
	default:
		return gpu.NewUint(rvalue.Uint())
	}
}

func Min[B gpu.AnyUint | XY](a XY, b B) XY { //glsl:min(uvec2,uvec2)uvec2 min(uvec2,uint)uvec2
	return gpu.NewVec2uExpression(gpu.Fn("min", a, either(b)))
}
func Max[T gpu.AnyUint | XY](a XY, b T) XY { //glsl:max(uvec2,uint)uvec2 max(uvec2,uvec2)uvec2
	return gpu.NewVec2uExpression(gpu.Fn("max", a, either(b)))
}
func Clamp(a, min, max XY) XY { return gpu.NewVec2uExpression(gpu.Fn("clamp", a, min, max)) } //glsl:clamp(uvec2,uvec2,uvec2)uvec2
func ClampX[A, B gpu.AnyUint](a XY, min A, max B) XY { //glsl:clamp(uvec2,uint,uint)uvec2
	return gpu.NewVec2uExpression(gpu.Fn("clamp", a, gpu.NewUint(min), gpu.NewUint(max)))
}
func Equal(a, b XY) gpu.Vec2b    { return gpu.NewVec2bExpression(gpu.Fn("equal", a, b)) }    //glsl:equal(uvec2,uvec2)bvec
func LessThan(a, b XY) gpu.Vec2b { return gpu.NewVec2bExpression(gpu.Fn("lessThan", a, b)) } //glsl:lessThan(uvec2,uvec2)bvec
func LessThanEqual(a, b XY) gpu.Vec2b { //glsl:lessThanEqual(uvec2,uvec2)bvec
	return gpu.NewVec2bExpression(gpu.Fn("lessThanEqual", a, b))
}
func GreaterThan(a, b XY) gpu.Vec2b { //glsl:greaterThan(uvec2,uvec2)bvec
	return gpu.NewVec2bExpression(gpu.Fn("greaterThan", a, b))
}
func GreaterThanEqual(a, b XY) gpu.Vec2b { //glsl:greaterThanEqual(uvec2,uvec2)bvec
	return gpu.NewVec2bExpression(gpu.Fn("greaterThanEqual", a, b))
}
func NotEqual(a, b XY) gpu.Vec2b { return gpu.NewVec2bExpression(gpu.Fn("notEqual", a, b)) } //glsl:notEqual(uvec2,uvec2)bvec
