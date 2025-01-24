// Package vec3 provides GPU operations on three-component floating-point vectors.
package vec3

import (
	"reflect"

	"graphics.gd/shaders/internal/gpu"
)

// XYZ is a three-component vector of floating-point values on the GPU.
type XYZ gpu.Vec3

// RGB is a three-component vector of red, green, and blue color values on the GPU.
type RGB gpu.RGB

func New[X, Y, Z gpu.AnyFloat](x X, y Y, z Z) XYZ { return gpu.NewVec3(x, y, z) }

func Add[T gpu.AnyFloat | XYZ](a XYZ, b T) XYZ { //glsl:+(vec3,vec3)
	return gpu.NewVec3Expression(gpu.Op(either(a), "+", either(b)))
}
func Sub[T gpu.AnyFloat | XYZ](a XYZ, b T) XYZ { //glsl:-(vec3,vec3)
	return gpu.NewVec3Expression(gpu.Op(either(a), "-", either(b)))
}
func Mul[T gpu.AnyFloat | XYZ](a XYZ, b T) XYZ { //glsl:*(vec3,vec3)
	return gpu.NewVec3Expression(gpu.Op(either(a), "*", either(b)))
}
func Div[T gpu.AnyFloat | XYZ](a XYZ, b T) XYZ { //glsl:/(vec3,vec3)
	return gpu.NewVec3Expression(gpu.Op(either(a), "/", either(b)))
}
func Neg(a XYZ) XYZ { return gpu.NewVec3Expression(gpu.Op(nil, "-", a)) } //glsl:-(vec3)

func either[T gpu.AnyFloat | XYZ](v T) gpu.Evaluator {
	rvalue := reflect.ValueOf(v)
	switch {
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(XYZ{})):
		return rvalue.Convert(reflect.TypeOf(XYZ{})).Interface().(XYZ)
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

func Abs(a XYZ) XYZ      { return gpu.NewVec3Expression(gpu.Fn("abs", a)) }     //glsl:abs(vec3)vec3
func Acos(a XYZ) XYZ     { return gpu.NewVec3Expression(gpu.Fn("acos", a)) }    //glsl:acos(vec3)vec3
func Acosh(a XYZ) XYZ    { return gpu.NewVec3Expression(gpu.Fn("acosh", a)) }   //glsl:acosh(vec3)vec3
func Asin(a XYZ) XYZ     { return gpu.NewVec3Expression(gpu.Fn("asin", a)) }    //glsl:asin(vec3)vec3
func Asinh(a XYZ) XYZ    { return gpu.NewVec3Expression(gpu.Fn("asinh", a)) }   //glsl:asinh(vec3)vec3
func Atan2(a, b XYZ) XYZ { return gpu.NewVec3Expression(gpu.Fn("atan", a, b)) } //glsl:atan(vec3,vec3)vec3
func Atan(a XYZ) XYZ     { return gpu.NewVec3Expression(gpu.Fn("atan", a)) }    //glsl:atan(vec3)vec3
func Atanh(a XYZ) XYZ    { return gpu.NewVec3Expression(gpu.Fn("atanh", a)) }   //glsl:atanh(vec3)vec3
func Ceil(a XYZ) XYZ     { return gpu.NewVec3Expression(gpu.Fn("ceil", a)) }    //glsl:ceil(vec3)vec3
func Clmap(a, min, max XYZ) XYZ { //glsl:clamp(vec3,vec3,vec3)vec3
	return gpu.NewVec3Expression(gpu.Fn("clamp", a, min, max))
}
func ClampX[A, B gpu.AnyFloat](a XYZ, min A, max B) XYZ { //glsl:clamp(vec3,float,float)vec3
	return gpu.NewVec3Expression(gpu.Fn("clamp", a, gpu.NewFloat(min), gpu.NewFloat(max)))
}
func Cos(a XYZ) XYZ     { return gpu.NewVec3Expression(gpu.Fn("cos", a)) }     //glsl:cos(vec3)vec3
func Cosh(a XYZ) XYZ    { return gpu.NewVec3Expression(gpu.Fn("cosh", a)) }    //glsl:cosh(vec3)vec3
func DFdx(a XYZ) XYZ    { return gpu.NewVec3Expression(gpu.Fn("dFdx", a)) }    //glsl:dFdx(vec3)vec3
func DFdy(a XYZ) XYZ    { return gpu.NewVec3Expression(gpu.Fn("dFdy", a)) }    //glsl:dFdy(vec3)vec3
func Degrees(a XYZ) XYZ { return gpu.NewVec3Expression(gpu.Fn("degrees", a)) } //glsl:degrees(vec3)vec3
func Exp(a XYZ) XYZ     { return gpu.NewVec3Expression(gpu.Fn("exp", a)) }     //glsl:exp(vec3)vec3
func Exp2(a XYZ) XYZ    { return gpu.NewVec3Expression(gpu.Fn("exp2", a)) }    //glsl:exp2(vec3)vec3
func FaceForward(a, b, n XYZ) XYZ { //glsl:faceforward(vec3,vec3,vec3)vec3
	return gpu.NewVec3Expression(gpu.Fn("faceforward", a, b, n))
}
func Floor(a XYZ) XYZ       { return gpu.NewVec3Expression(gpu.Fn("floor", a)) }       //glsl:floor(vec3)vec3
func Fract(a XYZ) XYZ       { return gpu.NewVec3Expression(gpu.Fn("fract", a)) }       //glsl:fract(vec3)vec3
func Fwidth(a XYZ) XYZ      { return gpu.NewVec3Expression(gpu.Fn("fwidth", a)) }      //glsl:fwidth(vec3)vec3
func InverseSqrt(a XYZ) XYZ { return gpu.NewVec3Expression(gpu.Fn("inversesqrt", a)) } //glsl:inversesqrt(vec3)vec3
func Log(a XYZ) XYZ         { return gpu.NewVec3Expression(gpu.Fn("log", a)) }         //glsl:log(vec3)vec3
func Log2(a XYZ) XYZ        { return gpu.NewVec3Expression(gpu.Fn("log2", a)) }        //glsl:log2(vec3)vec3
func Max[T gpu.AnyFloat | XYZ](a XYZ, b T) XYZ { //glsl:max(vec3,float)vec3 max(vec3,vec3)vec3
	return gpu.NewVec3Expression(gpu.Fn("max", a, either(b)))
}
func Min[T gpu.AnyFloat | XYZ](a XYZ, b T) XYZ { //glsl:min(vec3,float)vec3 min(vec3,vec3)vec3
	return gpu.NewVec3Expression(gpu.Fn("min", a, either(b)))
}
func Mix[T gpu.AnyFloat | XYZ](a XYZ, b, t T) XYZ { //glsl:mix(vec3,float,float)vec3 mix(vec3,vec3,vec3)vec3
	return gpu.NewVec3Expression(gpu.Fn("mix", a, either(b), either(t)))
}
func Mod[T gpu.AnyFloat | XYZ](a XYZ, b T) XYZ { //glsl:mod(vec3,vec3)vec3 mod(vec3,float)vec3
	return gpu.NewVec3Expression(gpu.Fn("mod", a, either(b)))
}
func Modf(a XYZ) (XYZ, XYZ) { //glsl:modf(vec3,out[vec3])vec3
	out := gpu.NewVec3Expression(gpu.Out("vec3"))
	return out, gpu.NewVec3Expression(gpu.Fn("modf", a, out))
}
func Normalize(a XYZ) XYZ { return gpu.NewVec3Expression(gpu.Fn("normalize", a)) } //glsl:normalize(vec3)vec3
func Pow(a, b XYZ) XYZ    { return gpu.NewVec3Expression(gpu.Fn("pow", a, b)) }    //glsl:pow(vec3,vec3)vec3
func Radians(a XYZ) XYZ { //glsl:radians(vec3)vec3
	return gpu.NewVec3Expression(gpu.Fn("radians", a))
}
func Reflect(a, b XYZ) XYZ { //glsl:reflect(vec3,vec3)vec3
	return gpu.NewVec3Expression(gpu.Fn("reflect", a, b))
}
func Refract[T gpu.AnyFloat](a, b XYZ, eta T) XYZ { //glsl:refract(vec3,vec3,float)vec3
	return gpu.NewVec3Expression(gpu.Fn("refract", a, b, gpu.NewFloat(eta)))
}
func Round(a XYZ) XYZ     { return gpu.NewVec3Expression(gpu.Fn("round", a)) }     //glsl:round(vec3)vec3
func RoundEven(a XYZ) XYZ { return gpu.NewVec3Expression(gpu.Fn("roundEven", a)) } //glsl:roundEven(vec3)vec3
func Sign(a XYZ) XYZ      { return gpu.NewVec3Expression(gpu.Fn("sign", a)) }      //glsl:sign(vec3)vec3
func Sin(a XYZ) XYZ       { return gpu.NewVec3Expression(gpu.Fn("sin", a)) }       //glsl:sin(vec3)vec3
func Sinh(a XYZ) XYZ      { return gpu.NewVec3Expression(gpu.Fn("sinh", a)) }      //glsl:sinh(vec3)vec3
func SmoothStep[T gpu.AnyFloat | XYZ](a, b T, x XYZ) XYZ { //glsl:smoothstep(float,float,vec3)vec3 smoothstep(vec3,vec3,vec3)vec3
	return gpu.NewVec3Expression(gpu.Fn("smoothstep", either(a), either(b), x))
}
func Sqrt(a XYZ) XYZ { return gpu.NewVec3Expression(gpu.Fn("sqrt", a)) } //glsl:sqrt(vec3)vec3
func Step[T gpu.AnyFloat | XYZ](a T, x XYZ) XYZ { //glsl:step(float,vec3)vec3 step(vec3,vec3)vec3
	return gpu.NewVec3Expression(gpu.Fn("step", either(a), x))
}
func Tan(a XYZ) XYZ               { return gpu.NewVec3Expression(gpu.Fn("tan", a)) }          //glsl:tan(vec3)vec3
func Tanh(a XYZ) XYZ              { return gpu.NewVec3Expression(gpu.Fn("tanh", a)) }         //glsl:tanh(vec3)vec3
func Trunc(a XYZ) XYZ             { return gpu.NewVec3Expression(gpu.Fn("trunc", a)) }        //glsl:trunc(vec3)vec3
func Equal(a, b XYZ) gpu.Vec3b    { return gpu.NewVec3bExpression(gpu.Fn("equal", a, b)) }    //glsl:equal(vec3,vec3)bvec
func LessThan(a, b XYZ) gpu.Vec3b { return gpu.NewVec3bExpression(gpu.Fn("lessThan", a, b)) } //glsl:lessThan(vec3,vec3)bvec
func GreaterThan(a, b XYZ) gpu.Vec3b { //glsl:greaterThan(vec3,vec3)bvec
	return gpu.NewVec3bExpression(gpu.Fn("greaterThan", a, b))
}
func LessThanEqual(a, b XYZ) gpu.Vec3b { //glsl:lessThanEqual(vec3,vec3)bvec
	return gpu.NewVec3bExpression(gpu.Fn("lessThanEqual", a, b))
}
func GreaterThanEqual(a, b XYZ) gpu.Vec3b { //glsl:greaterThanEqual(vec3,vec3)bvec
	return gpu.NewVec3bExpression(gpu.Fn("greaterThanEqual", a, b))
}
func NotEqual(a, b XYZ) gpu.Vec3b { return gpu.NewVec3bExpression(gpu.Fn("notEqual", a, b)) } //glsl:notEqual(vec3,vec3)bvec
func IsInf(a XYZ) gpu.Vec3b       { return gpu.NewVec3bExpression(gpu.Fn("isinf", a)) }       //glsl:isinf(vec3)bvec
func IsNaN(a XYZ) gpu.Vec3b       { return gpu.NewVec3bExpression(gpu.Fn("isnan", a)) }       //glsl:isnan(vec3)bvec
func Distance(a, b XYZ) gpu.Float { //glsl:distance(vec3,vec3)float
	return gpu.NewFloatExpression(gpu.Fn("distance", a, b))
}
func Dot(a, b XYZ) gpu.Float { //glsl:dot(vec3,vec3)float
	return gpu.NewFloatExpression(gpu.Fn("dot", a, b))
}
func Length(a XYZ) gpu.Float { return gpu.NewFloatExpression(gpu.Fn("length", a)) } //glsl:length(vec3)float
func Cross(a, b, c XYZ) XYZ { //glsl:cross(vec3,vec3)vec3
	return gpu.NewVec3Expression(gpu.Fn("cross", a, b, c))
}
