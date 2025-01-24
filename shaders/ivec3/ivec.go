// Package ivec3 provides GPU operations on three-component signed integer vectors.
package ivec3

import (
	"reflect"

	"graphics.gd/shaders/internal/gpu"
)

// XYZ is a three-component vector of signed integers on the GPU.
type XYZ gpu.Vec3i

// New creates a new [XYZ] value from three integers.
func New[X, Y, Z gpu.AnyInt](x X, y Y, z Z) XYZ { return gpu.NewVec3i(x, y, z) }

func either[T gpu.AnyInt | XYZ](v T) gpu.Evaluator {
	rvalue := reflect.ValueOf(v)
	switch {
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(XYZ{})):
		return rvalue.Convert(reflect.TypeOf(XYZ{})).Interface().(XYZ)
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(gpu.Int{})):
		return gpu.NewInt(rvalue.Convert(reflect.TypeOf(gpu.Int{})).Interface().(gpu.Int))
	default:
		return gpu.NewInt(rvalue.Int())
	}
}

func Add[T gpu.AnyInt | XYZ](a XYZ, b T) XYZ { //glsl:+(ivec3,ivec3)
	return gpu.NewVec3iExpression(gpu.Op(a, "+", either(b)))
}
func Sub[T gpu.AnyInt | XYZ](a XYZ, b T) XYZ { //glsl:-(ivec3,ivec3)
	return gpu.NewVec3iExpression(gpu.Op(a, "-", either(b)))
}
func Mul[T gpu.AnyInt | XYZ](a XYZ, b T) XYZ { //glsl:*(ivec3,ivec3)
	return gpu.NewVec3iExpression(gpu.Op(a, "*", either(b)))
}
func Div[T gpu.AnyInt | XYZ](a XYZ, b T) XYZ { //glsl:/(ivec3,ivec3)
	return gpu.NewVec3iExpression(gpu.Op(a, "/", either(b)))
}

func Neg(a XYZ) XYZ { return gpu.NewVec3iExpression(gpu.Op(nil, "-", a)) } //glsl:-(ivec3)

func BitsToFloat(a XYZ) gpu.Vec3 { return gpu.NewVec3Expression(gpu.Fn("intBitsToFloat", a)) } //glsl:intBitsToFloat(ivec3)vec

func Abs(a XYZ) XYZ  { return gpu.NewVec3iExpression(gpu.Fn("abs", a)) }  //glsl:abs(ivec3)ivec3
func Sign(a XYZ) XYZ { return gpu.NewVec3iExpression(gpu.Fn("sign", a)) } //glsl:sign(ivec3)ivec3
func Min[T XYZ | gpu.AnyInt](a XYZ, b T) XYZ { //glsl:min(ivec3,int)ivec3 min(ivec3,ivec3)ivec3
	return gpu.NewVec3iExpression(gpu.Fn("min", a, either(b)))
}
func Max[T gpu.AnyInt](a XYZ, b T) XYZ { //glsl:max(ivec3,int)ivec3 max(ivec3,ivec3)ivec3
	return gpu.NewVec3iExpression(gpu.Fn("max", a, either(b)))
}
func Clamp[T XYZ | gpu.AnyInt](a XYZ, min, max T) XYZ { //glsl:clamp(ivec3,ivec3,ivec3)ivec3 clamp(ivec3,int,int)ivec3
	return gpu.NewVec3iExpression(gpu.Fn("clamp", a, either(min), either(max)))
}
func Equal(a, b XYZ) gpu.Vec3b    { return gpu.NewVec3bExpression(gpu.Fn("equal", a, b)) }    //glsl:equal(ivec3,ivec3)bvec
func LessThan(a, b XYZ) gpu.Vec3b { return gpu.NewVec3bExpression(gpu.Fn("lessThan", a, b)) } //glsl:lessThan(ivec3,ivec3)bvec
func LessThanEqual(a, b XYZ) gpu.Vec3b { //glsl:lessThanEqual(ivec3,ivec3)bvec
	return gpu.NewVec3bExpression(gpu.Fn("lessThanEqual", a, b))
}
func GreaterThan(a, b XYZ) gpu.Vec3b { //glsl:greaterThan(ivec3,ivec3)bvec
	return gpu.NewVec3bExpression(gpu.Fn("greaterThan", a, b))
}
func GreaterThanEqual(a, b XYZ) gpu.Vec3b { //glsl:greaterThanEqual(ivec3,ivec3)bvec
	return gpu.NewVec3bExpression(gpu.Fn("greaterThanEqual", a, b))
}
func NotEqual(a, b XYZ) gpu.Vec3b { return gpu.NewVec3bExpression(gpu.Fn("notEqual", a, b)) } //glsl:notEqual(ivec3,ivec3)bvec
