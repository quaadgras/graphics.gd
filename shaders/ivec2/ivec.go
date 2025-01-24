// Package ivec2 provides GPU operations on two-component signed integer vectors.
package ivec2

import (
	"reflect"

	"graphics.gd/shaders/internal/gpu"
)

// XY is a two-component vector of signed integers on the GPU.
type XY gpu.Vec2i

// New creates a new [XY] value from two integers.
func New[X, Y gpu.AnyInt](x X, y Y) XY { return gpu.NewVec2i(x, y) }

func Add[T gpu.AnyInt | XY](a XY, b T) XY { return gpu.NewVec2iExpression(gpu.Op(a, "+", either(b))) } //glsl:+(ivec2,ivec2)
func Sub[T gpu.AnyInt | XY](a XY, b T) XY { return gpu.NewVec2iExpression(gpu.Op(a, "-", either(b))) } //glsl:-(ivec2,ivec2)
func Mul[T gpu.AnyInt | XY](a XY, b T) XY { return gpu.NewVec2iExpression(gpu.Op(a, "*", either(b))) } //glsl:*(ivec2,ivec2)
func Div[T gpu.AnyInt | XY](a XY, b T) XY { return gpu.NewVec2iExpression(gpu.Op(a, "/", either(b))) } //glsl:/(ivec2,ivec2)

func Neg(a XY) XY { return gpu.NewVec2iExpression(gpu.Op(nil, "-", a)) } //glsl:-(ivec2)

func BitsToFloat(a XY) gpu.Vec2 { return gpu.NewVec2Expression(gpu.Fn("intBitsToFloat", a)) } //glsl:intBitsToFloat(ivec2)vec

func either[T gpu.AnyInt | XY](v T) gpu.Evaluator {
	rvalue := reflect.ValueOf(v)
	switch {
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(XY{})):
		return rvalue.Convert(reflect.TypeOf(XY{})).Interface().(XY)
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(gpu.Int{})):
		return gpu.NewInt(rvalue.Convert(reflect.TypeOf(gpu.Int{})).Interface().(gpu.Int))
	default:
		return gpu.NewInt(rvalue.Int())
	}
}

func Abs(a XY) XY  { return gpu.NewVec2iExpression(gpu.Fn("abs", a)) }  //glsl:abs(ivec2)ivec2
func Sign(a XY) XY { return gpu.NewVec2iExpression(gpu.Fn("sign", a)) } //glsl:sign(ivec2)ivec2
func Min[B gpu.AnyInt | XY](a XY, b B) XY { //glsl:min(ivec2,ivec2)ivec2 min(ivec2,int)ivec2
	return gpu.NewVec2iExpression(gpu.Fn("min", a, either(b)))
}
func Max[T gpu.AnyInt | XY](a XY, b T) XY { //glsl:max(ivec2,int)ivec2 max(ivec2,ivec2)ivec2
	return gpu.NewVec2iExpression(gpu.Fn("max", a, either(b)))
}
func Clamp[T XY | gpu.AnyInt](a XY, min, max T) XY { //glsl:clamp(ivec2,int,int)ivec2 clamp(ivec2,ivec2,ivec2)ivec2
	return gpu.NewVec2iExpression(gpu.Fn("clamp", a, either(min), either(max)))
}
func Equal(a, b XY) gpu.Vec2b    { return gpu.NewVec2bExpression(gpu.Fn("equal", a, b)) }    //glsl:equal(ivec2,ivec2)bvec
func LessThan(a, b XY) gpu.Vec2b { return gpu.NewVec2bExpression(gpu.Fn("lessThan", a, b)) } //glsl:lessThan(ivec2,ivec2)bvec
func LessThanEqual(a, b XY) gpu.Vec2b { //glsl:lessThanEqual(ivec2,ivec2)bvec
	return gpu.NewVec2bExpression(gpu.Fn("lessThanEqual", a, b))
}
func GreaterThan(a, b XY) gpu.Vec2b { //glsl:greaterThan(ivec2,ivec2)bvec
	return gpu.NewVec2bExpression(gpu.Fn("greaterThan", a, b))
}
func GreaterThanEqual(a, b XY) gpu.Vec2b { //glsl:greaterThanEqual(ivec2,ivec2)bvec
	return gpu.NewVec2bExpression(gpu.Fn("greaterThanEqual", a, b))
}
func NotEqual(a, b XY) gpu.Vec2b { return gpu.NewVec2bExpression(gpu.Fn("notEqual", a, b)) } //glsl:notEqual(ivec2,ivec2)bvec
