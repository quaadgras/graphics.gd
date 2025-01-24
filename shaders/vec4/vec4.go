// Package vec4 provides GPU operations on four-component floating-point vectors.
package vec4

import (
	"reflect"

	"graphics.gd/shaders/internal/gpu"
)

// XYZW is a four-component vector of floating-point values on the GPU.
type XYZW gpu.Vec4

// RGBA is a four-component vector of floating-point values on the GPU.
type RGBA gpu.RGBA

func New[X, Y, Z, W gpu.AnyFloat](x X, y Y, z Z, w W) XYZW {
	return gpu.NewVec4(x, y, z, w)
}

func Add[T gpu.AnyFloat | XYZW](a XYZW, b T) XYZW { //glsl:+(vec4,vec4)
	return gpu.NewVec4Expression(gpu.Op(either(a), "+", either(b)))
}
func Sub[T gpu.AnyFloat | XYZW](a XYZW, b T) XYZW { //glsl:-(vec4,vec4)
	return gpu.NewVec4Expression(gpu.Op(either(a), "-", either(b)))
}
func Mul[T gpu.AnyFloat | XYZW](a XYZW, b T) XYZW { //glsl:*(vec4,vec4)
	return gpu.NewVec4Expression(gpu.Op(either(a), "*", either(b)))
}
func Div[T gpu.AnyFloat | XYZW](a XYZW, b T) XYZW { //glsl:/(vec4,vec4)
	return gpu.NewVec4Expression(gpu.Op(either(a), "/", either(b)))
}
func Neg(a XYZW) XYZW { return gpu.NewVec4Expression(gpu.Op(nil, "-", a)) } //glsl:-(vec4)

func either[T gpu.AnyFloat | XYZW](v T) gpu.Evaluator {
	rvalue := reflect.ValueOf(v)
	switch {
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(XYZW{})):
		return rvalue.Convert(reflect.TypeOf(XYZW{})).Interface().(XYZW)
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(gpu.Float{})):
		return gpu.NewFloat(rvalue.Convert(reflect.TypeOf(gpu.Float{})).Interface().(gpu.Float))
	default:
		return gpu.NewFloat(rvalue.Float())
	}
}

func FloatBitsToInt[A gpu.AnyFloat](a A) gpu.Vec4i { //glsl:floatBitsToInt(vec4)ivec
	return gpu.NewVec4iExpression(gpu.Fn("floatBitsToInt", gpu.NewFloat(a)))
}
func FloatBitsToUint[A gpu.AnyFloat](a A) gpu.Vec4u { //glsl:floatBitsToUint(vec4)uvec
	return gpu.NewVec4uExpression(gpu.Fn("floatBitsToUint", gpu.NewFloat(a)))
}

func Abs(a XYZW) XYZW      { return gpu.NewVec4Expression(gpu.Fn("abs", a)) }     //glsl:abs(vec4)vec4
func Acos(a XYZW) XYZW     { return gpu.NewVec4Expression(gpu.Fn("acos", a)) }    //glsl:acos(vec4)vec4
func Acosh(a XYZW) XYZW    { return gpu.NewVec4Expression(gpu.Fn("acosh", a)) }   //glsl:acosh(vec4)vec4
func Asin(a XYZW) XYZW     { return gpu.NewVec4Expression(gpu.Fn("asin", a)) }    //glsl:asin(vec4)vec4
func Asinh(a XYZW) XYZW    { return gpu.NewVec4Expression(gpu.Fn("asinh", a)) }   //glsl:asinh(vec4)vec4
func Atan2(a, b XYZW) XYZW { return gpu.NewVec4Expression(gpu.Fn("atan", a, b)) } //glsl:atan(vec4,vec4)vec4
func Atan(a XYZW) XYZW     { return gpu.NewVec4Expression(gpu.Fn("atan", a)) }    //glsl:atan(vec4)vec4
func Atanh(a XYZW) XYZW    { return gpu.NewVec4Expression(gpu.Fn("atanh", a)) }   //glsl:atanh(vec4)vec4
func Ceil(a XYZW) XYZW     { return gpu.NewVec4Expression(gpu.Fn("ceil", a)) }    //glsl:ceil(vec4)vec4
func Clamp(a, min, max XYZW) XYZW { //glsl:clamp(vec4,vec4,vec4)vec4
	return gpu.NewVec4Expression(gpu.Fn("clamp", a, min, max))
}
func ClampX[A, B gpu.AnyFloat](a XYZW, min A, max B) XYZW { //glsl:clamp(vec4,float,float)vec4
	return gpu.NewVec4Expression(gpu.Fn("clamp", a, gpu.NewFloat(min), gpu.NewFloat(max)))
}
func Cos(a XYZW) XYZW  { return gpu.NewVec4Expression(gpu.Fn("cos", a)) }  //glsl:cos(vec4)vec4
func Cosh(a XYZW) XYZW { return gpu.NewVec4Expression(gpu.Fn("cosh", a)) } //glsl:cosh(vec4)vec4
func DFdx(a XYZW) XYZW { return gpu.NewVec4Expression(gpu.Fn("dFdx", a)) } //glsl:dFdx(vec4)vec4
func DFdy(a XYZW) XYZW { return gpu.NewVec4Expression(gpu.Fn("dFdy", a)) } //glsl:dFdy(vec4)vec4
func Degrees(a XYZW) XYZW { //glsl:degrees(vec4)vec4
	return gpu.NewVec4Expression(gpu.Fn("degrees", a))
}
func Exp(a XYZW) XYZW  { return gpu.NewVec4Expression(gpu.Fn("exp", a)) }  //glsl:exp(vec4)vec4
func Exp2(a XYZW) XYZW { return gpu.NewVec4Expression(gpu.Fn("exp2", a)) } //glsl:exp2(vec4)vec4
func FaceForward(a, b, n XYZW) XYZW { //glsl:faceforward(vec4,vec4,vec4)vec4
	return gpu.NewVec4Expression(gpu.Fn("faceforward", a, b, n))
}
func Floor(a XYZW) XYZW  { return gpu.NewVec4Expression(gpu.Fn("floor", a)) }  //glsl:floor(vec4)vec4
func Fract(a XYZW) XYZW  { return gpu.NewVec4Expression(gpu.Fn("fract", a)) }  //glsl:fract(vec4)vec4
func Fwidth(a XYZW) XYZW { return gpu.NewVec4Expression(gpu.Fn("fwidth", a)) } //glsl:fwidth(vec4)vec4
func InverseSqrt(a XYZW) XYZW { //glsl:inversesqrt(vec4)vec4
	return gpu.NewVec4Expression(gpu.Fn("inversesqrt", a))
}
func Log(a XYZW) XYZW  { return gpu.NewVec4Expression(gpu.Fn("log", a)) }  //glsl:log(vec4)vec4
func Log2(a XYZW) XYZW { return gpu.NewVec4Expression(gpu.Fn("log2", a)) } //glsl:log2(vec4)vec4
func Max[T gpu.AnyFloat | XYZW](a XYZW, b T) XYZW { //glsl:max(vec4,float)vec4 max(vec4,vec4)vec4
	return gpu.NewVec4Expression(gpu.Fn("max", a, either(b)))
}
func Min[T gpu.AnyFloat | XYZW](a XYZW, b T) XYZW { //glsl:min(vec4,float)vec4 min(vec4,vec4)vec4
	return gpu.NewVec4Expression(gpu.Fn("min", a, either(b)))
}
func Mix[T gpu.AnyFloat | XYZW](a XYZW, b, t T) XYZW { //glsl:mix(vec4,float,float)vec4 mix(vec4,vec4,vec4)vec4
	return gpu.NewVec4Expression(gpu.Fn("mix", a, either(b), either(t)))
}
func Mod[T gpu.AnyFloat | XYZW](a XYZW, b T) XYZW { //glsl:mod(vec4,vec4)vec4 mod(vec4,float)vec4
	return gpu.NewVec4Expression(gpu.Fn("mod", a, either(b)))
}
func Modf(a XYZW) (XYZW, XYZW) { //glsl:modf(vec4,out[vec4])vec4
	out := gpu.NewVec4Expression(gpu.Out("vec4"))
	return out, gpu.NewVec4Expression(gpu.Fn("modf", a, out))
}
func Normalize(a XYZW) XYZW { return gpu.NewVec4Expression(gpu.Fn("normalize", a)) } //glsl:normalize(vec4)vec4
func Pow(a, b XYZW) XYZW    { return gpu.NewVec4Expression(gpu.Fn("pow", a, b)) }    //glsl:pow(vec4,vec4)vec4
func Radians(a XYZW) XYZW { //glsl:radians(vec4)vec4
	return gpu.NewVec4Expression(gpu.Fn("radians", a))
}
func Reflect(a, b XYZW) XYZW { //glsl:reflect(vec4,vec4)vec4
	return gpu.NewVec4Expression(gpu.Fn("reflect", a, b))
}
func Refract[T gpu.AnyFloat](a, b XYZW, c T) XYZW { //glsl:refract(vec4,vec4,float)vec4
	return gpu.NewVec4Expression(gpu.Fn("refract", a, b, gpu.NewFloat(c)))
}
func Round(a XYZW) XYZW     { return gpu.NewVec4Expression(gpu.Fn("round", a)) }     //glsl:round(vec4)vec4
func RoundEven(a XYZW) XYZW { return gpu.NewVec4Expression(gpu.Fn("roundEven", a)) } //glsl:roundEven(vec4)vec4
func Sign(a XYZW) XYZW      { return gpu.NewVec4Expression(gpu.Fn("sign", a)) }      //glsl:sign(vec4)vec4
func Sin(a XYZW) XYZW       { return gpu.NewVec4Expression(gpu.Fn("sin", a)) }       //glsl:sin(vec4)vec4
func Sinh(a XYZW) XYZW      { return gpu.NewVec4Expression(gpu.Fn("sinh", a)) }      //glsl:sinh(vec4)vec4
func SmoothStep[T gpu.AnyFloat | XYZW](a, b T, x XYZW) XYZW { //glsl:smoothstep(float,float,vec4)vec4 smoothstep(vec4,vec4,vec4)vec4
	return gpu.NewVec4Expression(gpu.Fn("smoothstep", either(a), either(b), x))
}
func Sqrt(a XYZW) XYZW { return gpu.NewVec4Expression(gpu.Fn("sqrt", a)) } //glsl:sqrt(vec4)vec4
func Step[T gpu.AnyFloat | XYZW](a T, x XYZW) XYZW { //glsl:step(float,vec4)vec4 step(vec4,vec4)vec4
	return gpu.NewVec4Expression(gpu.Fn("step", either(a), x))
}
func Tan(a XYZW) XYZW   { return gpu.NewVec4Expression(gpu.Fn("tan", a)) }   //glsl:tan(vec4)vec4
func Tanh(a XYZW) XYZW  { return gpu.NewVec4Expression(gpu.Fn("tanh", a)) }  //glsl:tanh(vec4)vec4
func Trunc(a XYZW) XYZW { return gpu.NewVec4Expression(gpu.Fn("trunc", a)) } //glsl:trunc(vec4)vec4
func Equal(a, b XYZW) gpu.Vec4b { //glsl:equal(vec4,vec4)bvec
	return gpu.NewVec4bExpression(gpu.Fn("equal", a, b))
}
func GreaterThan(a, b XYZW) gpu.Vec4b { //glsl:greaterThan(vec4,vec4)bvec
	return gpu.NewVec4bExpression(gpu.Fn("greaterThan", a, b))
}
func GreaterThanEqual(a, b XYZW) gpu.Vec4b { //glsl:greaterThanEqual(vec4,vec4)bvec
	return gpu.NewVec4bExpression(gpu.Fn("greaterThanEqual", a, b))
}
func LessThan(a, b XYZW) gpu.Vec4b { //glsl:lessThan(vec4,vec4)bvec
	return gpu.NewVec4bExpression(gpu.Fn("lessThan", a, b))
}
func LessThanEqual(a, b XYZW) gpu.Vec4b { //glsl:lessThanEqual(vec4,vec4)bvec
	return gpu.NewVec4bExpression(gpu.Fn("lessThanEqual", a, b))
}
func NotEqual(a, b XYZW) gpu.Vec4b { //glsl:notEqual(vec4,vec4)bvec
	return gpu.NewVec4bExpression(gpu.Fn("notEqual", a, b))
}
func IsInf(a XYZW) gpu.Vec4b { //glsl:isinf(vec4)bvec
	return gpu.NewVec4bExpression(gpu.Fn("isinf", a))
}
func IsNaN(a XYZW) gpu.Vec4b { //glsl:isnan(vec4)bvec
	return gpu.NewVec4bExpression(gpu.Fn("isnan", a))
}
func Distance(a, b XYZW) gpu.Float { //glsl:distance(vec4,vec4)float
	return gpu.NewFloatExpression(gpu.Fn("distance", a, b))
}
func Dot(a, b XYZW) gpu.Float { //glsl:dot(vec4,vec4)float
	return gpu.NewFloatExpression(gpu.Fn("dot", a, b))
}
func Length(a XYZW) gpu.Float { //glsl:length(vec4)float
	return gpu.NewFloatExpression(gpu.Fn("length", a))
}
