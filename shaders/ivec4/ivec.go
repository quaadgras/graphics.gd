// Package ivec4 provides GPU operations on four-component signed integer vectors.
package ivec4

import (
	"reflect"

	"graphics.gd/shaders/internal/gpu"
)

// XYZW is a four-component vector of signed integers on the GPU.
type XYZW gpu.Vec4i

// New creates a new [XYZW] value from two integers.
func New[X, Y, Z, W gpu.AnyInt](x X, y Y, z Z, w W) XYZW { return gpu.NewVec4i(x, y, z, w) }

func either[T gpu.AnyInt | XYZW](v T) gpu.Evaluator {
	rvalue := reflect.ValueOf(v)
	switch {
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(XYZW{})):
		return rvalue.Convert(reflect.TypeOf(XYZW{})).Interface().(XYZW)
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(gpu.Int{})):
		return gpu.NewInt(rvalue.Convert(reflect.TypeOf(gpu.Int{})).Interface().(gpu.Int))
	default:
		return gpu.NewInt(rvalue.Int())
	}
}

func Add[T gpu.AnyInt | XYZW](a XYZW, b T) XYZW { //glsl:+(ivec4,ivec4)
	return gpu.NewVec4iExpression(gpu.Op(a, "+", either(b)))
}
func Sub[T gpu.AnyInt | XYZW](a XYZW, b T) XYZW { //glsl:-(ivec4,ivec4)
	return gpu.NewVec4iExpression(gpu.Op(a, "-", either(b)))
}
func Mul[T gpu.AnyInt | XYZW](a XYZW, b T) XYZW { //glsl:*(ivec4,ivec4)
	return gpu.NewVec4iExpression(gpu.Op(a, "*", either(b)))
}
func Div[T gpu.AnyInt | XYZW](a XYZW, b T) XYZW { //glsl:/(ivec4,ivec4)
	return gpu.NewVec4iExpression(gpu.Op(a, "/", either(b)))
}
func Neg(a XYZW) XYZW { return gpu.NewVec4iExpression(gpu.Op(nil, "-", a)) } //glsl:-(ivec4)

func BitsToFloat(a XYZW) gpu.Vec4 { return gpu.NewVec4Expression(gpu.Fn("intBitsToFloat", a)) } //glsl:intBitsToFloat(ivec4)vec

func Abs(a XYZW) XYZW  { return gpu.NewVec4iExpression(gpu.Fn("abs", a)) }  //glsl:abs(ivec4)ivec4
func Sign(a XYZW) XYZW { return gpu.NewVec4iExpression(gpu.Fn("sign", a)) } //glsl:sign(ivec4)ivec4
func Min[T gpu.AnyInt | XYZW](a XYZW, b T) XYZW { //glsl:min(ivec4,ivec4)ivec4 min(ivec4,int)ivec4
	return gpu.NewVec4iExpression(gpu.Fn("min", a, either(b)))
}
func Max[T gpu.AnyInt | XYZW](a XYZW, b T) XYZW { //glsl:max(ivec4,int)ivec4 max(ivec4,ivec4)ivec4
	return gpu.NewVec4iExpression(gpu.Fn("max", a, either(b)))
}
func Clamp[T XYZW | gpu.AnyInt](a XYZW, min, max T) XYZW { //glsl:clamp(ivec4,ivec4,ivec4)ivec4 clamp(ivec4,int,int)ivec4
	return gpu.NewVec4iExpression(gpu.Fn("clamp", a, either(min), either(max)))
}
func Equal(a, b XYZW) gpu.Vec4b    { return gpu.NewVec4bExpression(gpu.Fn("equal", a, b)) }    //glsl:equal(ivec4,ivec4)bvec
func LessThan(a, b XYZW) gpu.Vec4b { return gpu.NewVec4bExpression(gpu.Fn("lessThan", a, b)) } //glsl:lessThan(ivec4,ivec4)bvec
func LessThanEqual(a, b XYZW) gpu.Vec4b { //glsl:lessThanEqual(ivec4,ivec4)bvec
	return gpu.NewVec4bExpression(gpu.Fn("lessThanEqual", a, b))
}
func GreaterThan(a, b XYZW) gpu.Vec4b { //glsl:greaterThan(ivec4,ivec4)bvec
	return gpu.NewVec4bExpression(gpu.Fn("greaterThan", a, b))
}
func GreaterThanEqual(a, b XYZW) gpu.Vec4b { //glsl:greaterThanEqual(ivec4,ivec4)bvec
	return gpu.NewVec4bExpression(gpu.Fn("greaterThanEqual", a, b))
}
func NotEqual(a, b XYZW) gpu.Vec4b { return gpu.NewVec4bExpression(gpu.Fn("notEqual", a, b)) } //glsl:notEqual(ivec4,ivec4)bvec
