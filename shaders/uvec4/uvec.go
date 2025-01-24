// Package uvec4 provides GPU operations on four-component unsigned integer vectors.
package uvec4

import (
	"reflect"

	"graphics.gd/shaders/internal/gpu"
)

// XYZW is a four-component vector of unsigned integers on the GPU.
type XYZW gpu.Vec4u

// New creates a new [XYZW] value from two integers.
func New[X, Y, Z, W gpu.AnyUint](x X, y Y, z Z, w W) XYZW { return gpu.NewVec4u(x, y, z, w) }

func either[T gpu.AnyUint | XYZW](v T) gpu.Evaluator {
	rvalue := reflect.ValueOf(v)
	switch {
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(XYZW{})):
		return rvalue.Convert(reflect.TypeOf(XYZW{})).Interface().(XYZW)
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(gpu.Uint{})):
		return gpu.NewUint(rvalue.Convert(reflect.TypeOf(gpu.Uint{})).Interface().(gpu.Uint))
	default:
		return gpu.NewUint(rvalue.Uint())
	}
}

func Add[T gpu.AnyUint | XYZW](a XYZW, b T) XYZW { //glsl:+(uvec4,uvec4)
	return gpu.NewVec4uExpression(gpu.Op(a, "+", either(b)))
}
func Sub[T gpu.AnyUint | XYZW](a XYZW, b T) XYZW { //glsl:-(uvec4,uvec4)
	return gpu.NewVec4uExpression(gpu.Op(a, "-", either(b)))
}
func Mul[T gpu.AnyUint | XYZW](a XYZW, b T) XYZW { //glsl:*(uvec4,uvec4)
	return gpu.NewVec4uExpression(gpu.Op(a, "*", either(b)))
}
func Div[T gpu.AnyUint | XYZW](a XYZW, b T) XYZW { //glsl:/(uvec4,uvec4)
	return gpu.NewVec4uExpression(gpu.Op(a, "/", either(b)))
}

func BitsToFloat(a XYZW) gpu.Vec4 { return gpu.NewVec4Expression(gpu.Fn("uintBitsToFloat", a)) } //glsl:uintBitsToFloat(uvec4)vec

func Min[T gpu.AnyUint | XYZW](a XYZW, b T) XYZW { //glsl:min(uvec4,uvec4)uvec4 min(uvec4,uint)uvec4
	return gpu.NewVec4uExpression(gpu.Fn("min", a, either(b)))
}
func Max[T gpu.AnyUint | XYZW](a XYZW, b T) XYZW { //glsl:max(uvec4,uint)uvec4 max(uvec4,uvec4)uvec4
	return gpu.NewVec4uExpression(gpu.Fn("max", a, either(b)))
}
func Clamp(a, min, max XYZW) XYZW { return gpu.NewVec4uExpression(gpu.Fn("clamp", a, min, max)) } //glsl:clamp(uvec4,uvec4,uvec4)uvec4
func ClampX[A, B gpu.AnyUint](a XYZW, min A, max B) XYZW { //glsl:clamp(uvec4,uint,uint)uvec4
	return gpu.NewVec4uExpression(gpu.Fn("clamp", a, gpu.NewUint(min), gpu.NewUint(max)))
}
func Equal(a, b XYZW) gpu.Vec4b    { return gpu.NewVec4bExpression(gpu.Fn("equal", a, b)) }    //glsl:equal(uvec4,uvec4)bvec
func LessThan(a, b XYZW) gpu.Vec4b { return gpu.NewVec4bExpression(gpu.Fn("lessThan", a, b)) } //glsl:lessThan(uvec4,uvec4)bvec
func LessThanEqual(a, b XYZW) gpu.Vec4b { //glsl:lessThanEqual(uvec4,uvec4)bvec
	return gpu.NewVec4bExpression(gpu.Fn("lessThanEqual", a, b))
}
func GreaterThan(a, b XYZW) gpu.Vec4b { //glsl:greaterThan(uvec4,uvec4)bvec
	return gpu.NewVec4bExpression(gpu.Fn("greaterThan", a, b))
}
func GreaterThanEqual(a, b XYZW) gpu.Vec4b { //glsl:greaterThanEqual(uvec4,uvec4)bvec
	return gpu.NewVec4bExpression(gpu.Fn("greaterThanEqual", a, b))
}
func NotEqual(a, b XYZW) gpu.Vec4b { return gpu.NewVec4bExpression(gpu.Fn("notEqual", a, b)) } //glsl:notEqual(uvec4,uvec4)bvec
