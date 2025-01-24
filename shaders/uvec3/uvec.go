// Package uvec3 provides GPU operations on three-component unsigned integer vectors.
package uvec3

import (
	"reflect"

	"graphics.gd/shaders/internal/gpu"
)

// XYZ is a three-component vector of unsigned integers on the GPU.
type XYZ gpu.Vec3u

// New creates a new [XYZ] value from three integers.
func New[X, Y, Z gpu.AnyUint](x X, y Y, z Z) XYZ { return gpu.NewVec3u(x, y, z) }

func either[T gpu.AnyUint | XYZ](v T) gpu.Evaluator {
	rvalue := reflect.ValueOf(v)
	switch {
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(XYZ{})):
		return rvalue.Convert(reflect.TypeOf(XYZ{})).Interface().(XYZ)
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(gpu.Uint{})):
		return gpu.NewUint(rvalue.Convert(reflect.TypeOf(gpu.Uint{})).Interface().(gpu.Uint))
	default:
		return gpu.NewUint(rvalue.Uint())
	}
}

func Add[T gpu.AnyUint | XYZ](a XYZ, b T) XYZ { //glsl:+(uvec3,uvec3)
	return gpu.NewVec3uExpression(gpu.Op(a, "+", either(b)))
}
func Sub[T gpu.AnyUint | XYZ](a XYZ, b T) XYZ { //glsl:-(uvec3,uvec3)
	return gpu.NewVec3uExpression(gpu.Op(a, "-", either(b)))
}
func Mul[T gpu.AnyUint | XYZ](a XYZ, b T) XYZ { //glsl:*(uvec3,uvec3)
	return gpu.NewVec3uExpression(gpu.Op(a, "*", either(b)))
}
func Div[T gpu.AnyUint | XYZ](a XYZ, b T) XYZ { //glsl:/(uvec3,uvec3)
	return gpu.NewVec3uExpression(gpu.Op(a, "/", either(b)))
}

func BitsToFloat(a XYZ) gpu.Vec3 { return gpu.NewVec3Expression(gpu.Fn("uintBitsToFloat", a)) } //glsl:uintBitsToFloat(uvec3)vec

func Min[T XYZ | gpu.AnyUint](a XYZ, b T) XYZ { //glsl:min(uvec3,uint)uvec3 min(uvec3,uvec3)uvec3
	return gpu.NewVec3uExpression(gpu.Fn("min", a, either(b)))
}
func Max[T gpu.AnyUint](a XYZ, b T) XYZ { //glsl:max(uvec3,uint)uvec3 max(uvec3,uvec3)uvec3
	return gpu.NewVec3uExpression(gpu.Fn("max", a, either(b)))
}
func Clamp(a, min, max XYZ) XYZ { return gpu.NewVec3uExpression(gpu.Fn("clamp", a, min, max)) } //glsl:clamp(uvec3,uvec3,uvec3)uvec3
func ClampX[A, B gpu.AnyUint](a XYZ, min A, max B) XYZ { //glsl:clamp(uvec3,uint,uint)uvec3
	return gpu.NewVec3uExpression(gpu.Fn("clamp", a, gpu.NewUint(min), gpu.NewUint(max)))
}
func Equal(a, b XYZ) gpu.Vec2b    { return gpu.NewVec2bExpression(gpu.Fn("equal", a, b)) }    //glsl:equal(uvec3,uvec3)bvec
func LessThan(a, b XYZ) gpu.Vec3b { return gpu.NewVec3bExpression(gpu.Fn("lessThan", a, b)) } //glsl:lessThan(uvec3,uvec3)bvec
func LessThanEqual(a, b XYZ) gpu.Vec3b { //glsl:lessThanEqual(uvec3,uvec3)bvec
	return gpu.NewVec3bExpression(gpu.Fn("lessThanEqual", a, b))
}
func GreaterThan(a, b XYZ) gpu.Vec3b { //glsl:greaterThan(uvec3,uvec3)bvec
	return gpu.NewVec3bExpression(gpu.Fn("greaterThan", a, b))
}
func GreaterThanEqual(a, b XYZ) gpu.Vec3b { //glsl:greaterThanEqual(uvec3,uvec3)bvec
	return gpu.NewVec3bExpression(gpu.Fn("greaterThanEqual", a, b))
}
func NotEqual(a, b XYZ) gpu.Vec3b { return gpu.NewVec3bExpression(gpu.Fn("notEqual", a, b)) } //glsl:notEqual(uvec3,uvec3)bvec
