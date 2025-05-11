// Package rgb provides GPU operations on three-component floating-point colors.
package rgb

import (
	"reflect"

	"graphics.gd/shaders/internal/gpu"
	"graphics.gd/shaders/vec3"
)

func New[X, Y, Z gpu.AnyFloat](x X, y Y, z Z) vec3.RGB { return gpu.NewRGB(x, y, z) }

func Add[T gpu.AnyFloat | vec3.RGB](a vec3.RGB, b T) vec3.RGB { //glsl:+(vec3,vec3)
	return gpu.NewRGBExpression(gpu.Op(either(a), "+", either(b)))
}
func Sub[T gpu.AnyFloat | vec3.RGB](a vec3.RGB, b T) vec3.RGB { //glsl:-(vec3,vec3)
	return gpu.NewRGBExpression(gpu.Op(either(a), "-", either(b)))
}
func Mul[T gpu.AnyFloat | vec3.RGB](a vec3.RGB, b T) vec3.RGB { //glsl:*(vec3,vec3)
	return gpu.NewRGBExpression(gpu.Op(either(a), "*", either(b)))
}
func Div[T gpu.AnyFloat | vec3.RGB](a vec3.RGB, b T) vec3.RGB { //glsl:/(vec3,vec3)
	return gpu.NewRGBExpression(gpu.Op(either(a), "/", either(b)))
}
func Neg(a vec3.RGB) vec3.RGB { return gpu.NewRGBExpression(gpu.Op(nil, "-", a)) } //glsl:-(vec3)

func either[T gpu.AnyFloat | vec3.RGB](v T) gpu.Evaluator {
	rvalue := reflect.ValueOf(v)
	switch {
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(vec3.RGB{})):
		return rvalue.Convert(reflect.TypeOf(vec3.RGB{})).Interface().(vec3.RGB)
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(gpu.Float{})):
		return gpu.NewFloat(rvalue.Convert(reflect.TypeOf(gpu.Float{})).Interface().(gpu.Float))
	default:
		return gpu.NewFloat(rvalue.Float())
	}
}

func FloatBitsToInt[A gpu.AnyFloat](a A) gpu.Vec3i { //glsl:floatBitsToInt(vec3)ivec
	return gpu.NewVec3iExpression(gpu.Fn("floatBitsToInt", gpu.NewFloat(a)))
}
func FloatBitsToUint[A gpu.AnyFloat](a A) gpu.Vec3u { //glsl:floatBitsToUint(vec3)uvec
	return gpu.NewVec3uExpression(gpu.Fn("floatBitsToUint", gpu.NewFloat(a)))
}

func Abs(a vec3.RGB) vec3.RGB      { return gpu.NewRGBExpression(gpu.Fn("abs", a)) }     //glsl:abs(vec3)vec3
func Acos(a vec3.RGB) vec3.RGB     { return gpu.NewRGBExpression(gpu.Fn("acos", a)) }    //glsl:acos(vec3)vec3
func Acosh(a vec3.RGB) vec3.RGB    { return gpu.NewRGBExpression(gpu.Fn("acosh", a)) }   //glsl:acosh(vec3)vec3
func Asin(a vec3.RGB) vec3.RGB     { return gpu.NewRGBExpression(gpu.Fn("asin", a)) }    //glsl:asin(vec3)vec3
func Asinh(a vec3.RGB) vec3.RGB    { return gpu.NewRGBExpression(gpu.Fn("asinh", a)) }   //glsl:asinh(vec3)vec3
func Atan2(a, b vec3.RGB) vec3.RGB { return gpu.NewRGBExpression(gpu.Fn("atan", a, b)) } //glsl:atan(vec3,vec3)vec3
func Atan(a vec3.RGB) vec3.RGB     { return gpu.NewRGBExpression(gpu.Fn("atan", a)) }    //glsl:atan(vec3)vec3
func Atanh(a vec3.RGB) vec3.RGB    { return gpu.NewRGBExpression(gpu.Fn("atanh", a)) }   //glsl:atanh(vec3)vec3
func Ceil(a vec3.RGB) vec3.RGB     { return gpu.NewRGBExpression(gpu.Fn("ceil", a)) }    //glsl:ceil(vec3)vec3
func Clmap(a, min, max vec3.RGB) vec3.RGB { //glsl:clamp(vec3,vec3,vec3)vec3
	return gpu.NewRGBExpression(gpu.Fn("clamp", a, min, max))
}
func ClampX[A, B gpu.AnyFloat](a vec3.RGB, min A, max B) vec3.RGB { //glsl:clamp(vec3,float,float)vec3
	return gpu.NewRGBExpression(gpu.Fn("clamp", a, gpu.NewFloat(min), gpu.NewFloat(max)))
}
func Cos(a vec3.RGB) vec3.RGB     { return gpu.NewRGBExpression(gpu.Fn("cos", a)) }     //glsl:cos(vec3)vec3
func Cosh(a vec3.RGB) vec3.RGB    { return gpu.NewRGBExpression(gpu.Fn("cosh", a)) }    //glsl:cosh(vec3)vec3
func DFdx(a vec3.RGB) vec3.RGB    { return gpu.NewRGBExpression(gpu.Fn("dFdx", a)) }    //glsl:dFdx(vec3)vec3
func DFdy(a vec3.RGB) vec3.RGB    { return gpu.NewRGBExpression(gpu.Fn("dFdy", a)) }    //glsl:dFdy(vec3)vec3
func Degrees(a vec3.RGB) vec3.RGB { return gpu.NewRGBExpression(gpu.Fn("degrees", a)) } //glsl:degrees(vec3)vec3
func Exp(a vec3.RGB) vec3.RGB     { return gpu.NewRGBExpression(gpu.Fn("exp", a)) }     //glsl:exp(vec3)vec3
func Exp2(a vec3.RGB) vec3.RGB    { return gpu.NewRGBExpression(gpu.Fn("exp2", a)) }    //glsl:exp2(vec3)vec3
func FaceForward(a, b, n vec3.RGB) vec3.RGB { //glsl:faceforward(vec3,vec3,vec3)vec3
	return gpu.NewRGBExpression(gpu.Fn("faceforward", a, b, n))
}
func Floor(a vec3.RGB) vec3.RGB       { return gpu.NewRGBExpression(gpu.Fn("floor", a)) }       //glsl:floor(vec3)vec3
func Fract(a vec3.RGB) vec3.RGB       { return gpu.NewRGBExpression(gpu.Fn("fract", a)) }       //glsl:fract(vec3)vec3
func Fwidth(a vec3.RGB) vec3.RGB      { return gpu.NewRGBExpression(gpu.Fn("fwidth", a)) }      //glsl:fwidth(vec3)vec3
func InverseSqrt(a vec3.RGB) vec3.RGB { return gpu.NewRGBExpression(gpu.Fn("inversesqrt", a)) } //glsl:inversesqrt(vec3)vec3
func Log(a vec3.RGB) vec3.RGB         { return gpu.NewRGBExpression(gpu.Fn("log", a)) }         //glsl:log(vec3)vec3
func Log2(a vec3.RGB) vec3.RGB        { return gpu.NewRGBExpression(gpu.Fn("log2", a)) }        //glsl:log2(vec3)vec3
func Max[T gpu.AnyFloat | vec3.RGB](a vec3.RGB, b T) vec3.RGB { //glsl:max(vec3,float)vec3 max(vec3,vec3)vec3
	return gpu.NewRGBExpression(gpu.Fn("max", a, either(b)))
}
func Min[T gpu.AnyFloat | vec3.RGB](a vec3.RGB, b T) vec3.RGB { //glsl:min(vec3,float)vec3 min(vec3,vec3)vec3
	return gpu.NewRGBExpression(gpu.Fn("min", a, either(b)))
}
func Mix[T gpu.AnyFloat | vec3.RGB](a vec3.RGB, b, t T) vec3.RGB { //glsl:mix(vec3,float,float)vec3 mix(vec3,vec3,vec3)vec3
	return gpu.NewRGBExpression(gpu.Fn("mix", a, either(b), either(t)))
}
func Mod[T gpu.AnyFloat | vec3.RGB](a vec3.RGB, b T) vec3.RGB { //glsl:mod(vec3,vec3)vec3 mod(vec3,float)vec3
	return gpu.NewRGBExpression(gpu.Fn("mod", a, either(b)))
}
func Modf(a vec3.RGB) (vec3.RGB, vec3.RGB) { //glsl:modf(vec3,out[vec3])vec3
	out := gpu.NewRGBExpression(gpu.Out("vec3"))
	return out, gpu.NewRGBExpression(gpu.Fn("modf", a, out))
}
func Normalize(a vec3.RGB) vec3.RGB { return gpu.NewRGBExpression(gpu.Fn("normalize", a)) } //glsl:normalize(vec3)vec3
func Pow(a, b vec3.RGB) vec3.RGB    { return gpu.NewRGBExpression(gpu.Fn("pow", a, b)) }    //glsl:pow(vec3,vec3)vec3
func Radians(a vec3.RGB) vec3.RGB { //glsl:radians(vec3)vec3
	return gpu.NewRGBExpression(gpu.Fn("radians", a))
}
func Reflect(a, b vec3.RGB) vec3.RGB { //glsl:reflect(vec3,vec3)vec3
	return gpu.NewRGBExpression(gpu.Fn("reflect", a, b))
}
func Refract[T gpu.AnyFloat](a, b vec3.RGB, eta T) vec3.RGB { //glsl:refract(vec3,vec3,float)vec3
	return gpu.NewRGBExpression(gpu.Fn("refract", a, b, gpu.NewFloat(eta)))
}
func Round(a vec3.RGB) vec3.RGB     { return gpu.NewRGBExpression(gpu.Fn("round", a)) }     //glsl:round(vec3)vec3
func RoundEven(a vec3.RGB) vec3.RGB { return gpu.NewRGBExpression(gpu.Fn("roundEven", a)) } //glsl:roundEven(vec3)vec3
func Sign(a vec3.RGB) vec3.RGB      { return gpu.NewRGBExpression(gpu.Fn("sign", a)) }      //glsl:sign(vec3)vec3
func Sin(a vec3.RGB) vec3.RGB       { return gpu.NewRGBExpression(gpu.Fn("sin", a)) }       //glsl:sin(vec3)vec3
func Sinh(a vec3.RGB) vec3.RGB      { return gpu.NewRGBExpression(gpu.Fn("sinh", a)) }      //glsl:sinh(vec3)vec3
func SmoothStep[T gpu.AnyFloat | vec3.RGB](a, b T, x vec3.RGB) vec3.RGB { //glsl:smoothstep(float,float,vec3)vec3 smoothstep(vec3,vec3,vec3)vec3
	return gpu.NewRGBExpression(gpu.Fn("smoothstep", either(a), either(b), x))
}
func Sqrt(a vec3.RGB) vec3.RGB { return gpu.NewRGBExpression(gpu.Fn("sqrt", a)) } //glsl:sqrt(vec3)vec3
func Step[T gpu.AnyFloat | vec3.RGB](a T, x vec3.RGB) vec3.RGB { //glsl:step(float,vec3)vec3 step(vec3,vec3)vec3
	return gpu.NewRGBExpression(gpu.Fn("step", either(a), x))
}
func Tan(a vec3.RGB) vec3.RGB          { return gpu.NewRGBExpression(gpu.Fn("tan", a)) }           //glsl:tan(vec3)vec3
func Tanh(a vec3.RGB) vec3.RGB         { return gpu.NewRGBExpression(gpu.Fn("tanh", a)) }          //glsl:tanh(vec3)vec3
func Trunc(a vec3.RGB) vec3.RGB        { return gpu.NewRGBExpression(gpu.Fn("trunc", a)) }         //glsl:trunc(vec3)vec3
func Equal(a, b vec3.RGB) gpu.Vec3b    { return gpu.NewVec3bExpression(gpu.Fn("equal", a, b)) }    //glsl:equal(vec3,vec3)bvec
func LessThan(a, b vec3.RGB) gpu.Vec3b { return gpu.NewVec3bExpression(gpu.Fn("lessThan", a, b)) } //glsl:lessThan(vec3,vec3)bvec
func GreaterThan(a, b vec3.RGB) gpu.Vec3b { //glsl:greaterThan(vec3,vec3)bvec
	return gpu.NewVec3bExpression(gpu.Fn("greaterThan", a, b))
}
func LessThanEqual(a, b vec3.RGB) gpu.Vec3b { //glsl:lessThanEqual(vec3,vec3)bvec
	return gpu.NewVec3bExpression(gpu.Fn("lessThanEqual", a, b))
}
func GreaterThanEqual(a, b vec3.RGB) gpu.Vec3b { //glsl:greaterThanEqual(vec3,vec3)bvec
	return gpu.NewVec3bExpression(gpu.Fn("greaterThanEqual", a, b))
}
func NotEqual(a, b vec3.RGB) gpu.Vec3b { return gpu.NewVec3bExpression(gpu.Fn("notEqual", a, b)) } //glsl:notEqual(vec3,vec3)bvec
func IsInf(a vec3.RGB) gpu.Vec3b       { return gpu.NewVec3bExpression(gpu.Fn("isinf", a)) }       //glsl:isinf(vec3)bvec
func IsNaN(a vec3.RGB) gpu.Vec3b       { return gpu.NewVec3bExpression(gpu.Fn("isnan", a)) }       //glsl:isnan(vec3)bvec
func Distance(a, b vec3.RGB) gpu.Float { //glsl:distance(vec3,vec3)float
	return gpu.NewFloatExpression(gpu.Fn("distance", a, b))
}
func Dot(a, b vec3.RGB) gpu.Float { //glsl:dot(vec3,vec3)float
	return gpu.NewFloatExpression(gpu.Fn("dot", a, b))
}
func Length(a vec3.RGB) gpu.Float { return gpu.NewFloatExpression(gpu.Fn("length", a)) } //glsl:length(vec3)float
func Cross(a, b, c vec3.RGB) vec3.RGB { //glsl:cross(vec3,vec3)vec3
	return gpu.NewRGBExpression(gpu.Fn("cross", a, b, c))
}
