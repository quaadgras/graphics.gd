// Package vec2 provides GPU operations on two-component floating-point vectors.
package vec2

import (
	"reflect"

	"graphics.gd/shaders/internal/gpu"
)

// XY is a two-component vector of floating-point values on the GPU.
type XY gpu.Vec2

func New[X, Y gpu.AnyFloat](x X, y Y) XY {
	return gpu.NewVec2(x, y)
}

func Add[T gpu.AnyFloat | XY](a XY, b T) XY { //glsl:+(vec2,vec2)
	return gpu.NewVec2Expression(gpu.Op(either(a), "+", either(b)))
}
func Sub[T gpu.AnyFloat | XY](a XY, b T) XY { //glsl:-(vec2,vec2)
	return gpu.NewVec2Expression(gpu.Op(either(a), "-", either(b)))
}
func Mul[T gpu.AnyFloat | XY](a XY, b T) XY { //glsl:*(vec2,vec2)
	return gpu.NewVec2Expression(gpu.Op(either(a), "*", either(b)))
}
func Div[T gpu.AnyFloat | XY](a XY, b T) XY { //glsl:/(vec2,vec2)
	return gpu.NewVec2Expression(gpu.Op(either(a), "/", either(b)))
}
func Neg(a XY) XY { return gpu.NewVec2Expression(gpu.Op(nil, "-", a)) } //glsl:-(vec2)

func either[T gpu.AnyFloat | XY](v T) gpu.Evaluator {
	rvalue := reflect.ValueOf(v)
	switch {
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(XY{})):
		return rvalue.Convert(reflect.TypeOf(XY{})).Interface().(XY)
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(gpu.Float{})):
		return gpu.NewFloat(rvalue.Convert(reflect.TypeOf(gpu.Float{})).Interface().(gpu.Float))
	default:
		return gpu.NewFloat(rvalue.Float())
	}
}

func FloatBitsToInt[A gpu.AnyFloat](a A) gpu.Vec2i { //glsl:floatBitsToInt(vec2)ivec
	return gpu.NewVec2iExpression(gpu.Fn("floatBitsToInt", gpu.NewFloat(a)))
}
func FloatBitsToUint[A gpu.AnyFloat](a A) gpu.Vec2u { //glsl:floatBitsToUint(vec2)uvec
	return gpu.NewVec2uExpression(gpu.Fn("floatBitsToUint", gpu.NewFloat(a)))
}
func PackHalf2x16(a XY) gpu.Uint { //glsl:packHalf2x16(vec2)uint
	return gpu.NewUintExpression(gpu.Fn("packHalf2x16", a))
}
func PackSnorm2x16(a XY) gpu.Uint { //glsl:packSnorm2x16(vec2)uint
	return gpu.NewUintExpression(gpu.Fn("packSnorm2x16", a))
}
func PackUnorm2x16(a XY) gpu.Uint { //glsl:packUnorm2x16(vec2)uint
	return gpu.NewUintExpression(gpu.Fn("packUnorm2x16", a))
}

func Abs(a XY) XY      { return gpu.NewVec2Expression(gpu.Fn("abs", a)) }     //glsl:abs(vec2)vec2
func Acos(a XY) XY     { return gpu.NewVec2Expression(gpu.Fn("acos", a)) }    //glsl:acos(vec2)vec2
func Acosh(a XY) XY    { return gpu.NewVec2Expression(gpu.Fn("acosh", a)) }   //glsl:acosh(vec2)vec2
func Asin(a XY) XY     { return gpu.NewVec2Expression(gpu.Fn("asin", a)) }    //glsl:asin(vec2)vec2
func Asinh(a XY) XY    { return gpu.NewVec2Expression(gpu.Fn("asinh", a)) }   //glsl:asinh(vec2)vec2
func Atan2(a, b XY) XY { return gpu.NewVec2Expression(gpu.Fn("atan", a, b)) } //glsl:atan(vec2,vec2)vec2
func Atan(a XY) XY     { return gpu.NewVec2Expression(gpu.Fn("atan", a)) }    //glsl:atan(vec2)vec2
func Atanh(a XY) XY    { return gpu.NewVec2Expression(gpu.Fn("atanh", a)) }   //glsl:atanh(vec2)vec2
func Ceil(a XY) XY     { return gpu.NewVec2Expression(gpu.Fn("ceil", a)) }    //glsl:ceil(vec2)vec2
func Clamp(a, min, max XY) XY { //glsl:clamp(vec2,vec2,vec2)vec2
	return gpu.NewVec2Expression(gpu.Fn("clamp", a, min, max))
}
func ClampX[A, B gpu.AnyFloat](a XY, min A, max B) XY { //glsl:clamp(vec2,float,float)vec2
	return gpu.NewVec2Expression(gpu.Fn("clamp", a, gpu.NewFloat(min), gpu.NewFloat(max)))
}
func Cos(a XY) XY     { return gpu.NewVec2Expression(gpu.Fn("cos", a)) }     //glsl:cos(vec2)vec2
func Cosh(a XY) XY    { return gpu.NewVec2Expression(gpu.Fn("cosh", a)) }    //glsl:cosh(vec2)vec2
func DFdx(a XY) XY    { return gpu.NewVec2Expression(gpu.Fn("dFdx", a)) }    //glsl:dFdx(vec2)vec2
func DFdy(a XY) XY    { return gpu.NewVec2Expression(gpu.Fn("dFdy", a)) }    //glsl:dFdy(vec2)vec2
func Degrees(a XY) XY { return gpu.NewVec2Expression(gpu.Fn("degrees", a)) } //glsl:degrees(vec2)vec2
func Exp(a XY) XY     { return gpu.NewVec2Expression(gpu.Fn("exp", a)) }     //glsl:exp(vec2)vec2
func Exp2(a XY) XY    { return gpu.NewVec2Expression(gpu.Fn("exp2", a)) }    //glsl:exp2(vec2)vec2
func FaceForward(a, b, n XY) XY { //glsl:faceforward(vec2,vec2,vec2)vec2
	return gpu.NewVec2Expression(gpu.Fn("faceforward", a, b, n))
}
func Floor(a XY) XY       { return gpu.NewVec2Expression(gpu.Fn("floor", a)) }       //glsl:floor(vec2)vec2
func Fract(a XY) XY       { return gpu.NewVec2Expression(gpu.Fn("fract", a)) }       //glsl:fract(vec2)vec2
func Fwidth(a XY) XY      { return gpu.NewVec2Expression(gpu.Fn("fwidth", a)) }      //glsl:fwidth(vec2)vec2
func InverseSqrt(a XY) XY { return gpu.NewVec2Expression(gpu.Fn("inversesqrt", a)) } //glsl:inversesqrt(vec2)vec2
func Log(a XY) XY         { return gpu.NewVec2Expression(gpu.Fn("log", a)) }         //glsl:log(vec2)vec2
func Log2(a XY) XY        { return gpu.NewVec2Expression(gpu.Fn("log2", a)) }        //glsl:log2(vec2)vec2
func Max[T gpu.AnyFloat | XY](a XY, b T) XY { //glsl:max(vec2,float)vec2 max(vec2,vec2)vec2
	return gpu.NewVec2Expression(gpu.Fn("max", a, either(b)))
}
func Min[T gpu.AnyFloat | XY](a XY, b T) XY { //glsl:min(vec2,float)vec2 min(vec2,vec2)vec2
	return gpu.NewVec2Expression(gpu.Fn("min", a, either(b)))
}
func Mix[T gpu.AnyFloat | XY](a XY, b T, t T) XY { //glsl:mix(vec2,vec2,vec2)vec2 mix(vec2,float,float)vec2
	return gpu.NewVec2Expression(gpu.Fn("mix", a, either(b), either(t)))
}
func Mod[T gpu.AnyFloat | XY](a XY, b T) XY { //glsl:mod(vec2,vec2)vec2 mod(vec2,float)vec2
	return gpu.NewVec2Expression(gpu.Fn("mod", a, either(b)))
}
func Modf(a XY) (XY, XY) { //glsl:modf(vec2,out[vec2])vec2
	out := gpu.NewVec2Expression(gpu.Out("vec2"))
	return out, gpu.NewVec2Expression(gpu.Fn("modf", a, out))
}
func Normalize(a XY) XY { return gpu.NewVec2Expression(gpu.Fn("normalize", a)) } //glsl:normalize(vec2)vec2
func Pow(a, b XY) XY    { return gpu.NewVec2Expression(gpu.Fn("pow", a, b)) }    //glsl:pow(vec2,vec2)vec2
func Radians(a XY) XY { //glsl:radians(vec2)vec2
	return gpu.NewVec2Expression(gpu.Fn("radians", a))
}
func Reflect(a, b XY) XY { //glsl:reflect(vec2,vec2)vec2
	return gpu.NewVec2Expression(gpu.Fn("reflect", a, b))
}
func Refract[T gpu.AnyFloat](a, b XY, c T) XY { //glsl:refract(vec2,vec2,float)vec2
	return gpu.NewVec2Expression(gpu.Fn("refract", a, b, gpu.NewFloat(c)))
}
func Round(a XY) XY     { return gpu.NewVec2Expression(gpu.Fn("round", a)) }     //glsl:round(vec2)vec2
func RoundEven(a XY) XY { return gpu.NewVec2Expression(gpu.Fn("roundEven", a)) } //glsl:roundEven(vec2)vec2
func Sign(a XY) XY      { return gpu.NewVec2Expression(gpu.Fn("sign", a)) }      //glsl:sign(vec2)vec2
func Sin(a XY) XY       { return gpu.NewVec2Expression(gpu.Fn("sin", a)) }       //glsl:sin(vec2)vec2
func Sinh(a XY) XY      { return gpu.NewVec2Expression(gpu.Fn("sinh", a)) }      //glsl:sinh(vec2)vec2
func Smoothstep[T gpu.AnyFloat | XY](a, b T, x XY) XY { //glsl:smoothstep(vec2,vec2,vec2)vec2 smoothstep(float,float,vec2)vec2
	return gpu.NewVec2Expression(gpu.Fn("smoothstep", either(a), either(b), x))
}
func Sqrt(a XY) XY { return gpu.NewVec2Expression(gpu.Fn("sqrt", a)) } //glsl:sqrt(vec2)vec2
func Step[T gpu.AnyFloat | XY](a T, x XY) XY { //glsl:step(float,vec2)vec2 step(vec2,vec2)vec2
	return gpu.NewVec2Expression(gpu.Fn("step", either(a), either(x)))
}
func Tan(a XY) XY                   { return gpu.NewVec2Expression(gpu.Fn("tan", a)) }             //glsl:tan(vec2)vec2
func Tanh(a XY) XY                  { return gpu.NewVec2Expression(gpu.Fn("tanh", a)) }            //glsl:tanh(vec2)vec2
func Trunc(a XY) XY                 { return gpu.NewVec2Expression(gpu.Fn("trunc", a)) }           //glsl:trunc(vec2)vec2
func Equal(a, b XY) gpu.Vec2b       { return gpu.NewVec2bExpression(gpu.Fn("equal", a, b)) }       //glsl:equal(vec2,vec2)bvec
func GreaterThan(a, b XY) gpu.Vec2b { return gpu.NewVec2bExpression(gpu.Fn("greaterThan", a, b)) } //glsl:greaterThan(vec2,vec2)bvec
func LessThan(a, b XY) gpu.Vec2b    { return gpu.NewVec2bExpression(gpu.Fn("lessThan", a, b)) }    //glsl:lessThan(vec2,vec2)bvec
func LessThanEqual(a, b XY) gpu.Vec2b { //glsl:lessThanEqual(vec2,vec2)bvec
	return gpu.NewVec2bExpression(gpu.Fn("lessThanEqual", a, b))
}
func NotEqual(a, b XY) gpu.Vec2b { return gpu.NewVec2bExpression(gpu.Fn("notEqual", a, b)) } //glsl:notEqual(vec2,vec2)bvec
func GreaterThanEqual(a, b XY) gpu.Vec2b { //glsl:greaterThanEqual(vec2,vec2)bvec
	return gpu.NewVec2bExpression(gpu.Fn("greaterThanEqual", a, b))
}
func IsInf(a XY) gpu.Vec2b { return gpu.NewVec2bExpression(gpu.Fn("isinf", a)) } //glsl:isinf(vec2)bvec
func IsNaN(a XY) gpu.Vec2b { return gpu.NewVec2bExpression(gpu.Fn("isnan", a)) } //glsl:isnan(vec2)bvec
func Distance(a, b XY) gpu.Float { //glsl:distance(vec2,vec2)float
	return gpu.NewFloatExpression(gpu.Fn("distance", a, b))
}
func Dot(a, b XY) gpu.Float { //glsl:dot(vec2,vec2)float
	return gpu.NewFloatExpression(gpu.Fn("dot", a, b))
}
func Length(a XY) gpu.Float { return gpu.NewFloatExpression(gpu.Fn("length", a)) } //glsl:length(vec2)float
