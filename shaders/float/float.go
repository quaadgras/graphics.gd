// Package float provides GPU operations on floating-point values.
package float

import "graphics.gd/shaders/internal/gpu"

// X is a floating-point value on the GPU (represented by an unspecified number of bits).
type X gpu.Float

// New creates a new [X] value from a float.
func New[T gpu.AnyFloat](x T) X { return gpu.NewFloat(x) }

func Add[A, B gpu.AnyFloat](a A, b B) X { //glsl:+(float,float)
	return gpu.NewFloatExpression(gpu.Op(gpu.NewFloat(a), "+", gpu.NewFloat(b)))
}
func Sub[A, B gpu.AnyFloat](a A, b B) X { //glsl:-(float,float)
	return gpu.NewFloatExpression(gpu.Op(gpu.NewFloat(a), "-", gpu.NewFloat(b)))
}
func Mul[A, B gpu.AnyFloat](a A, b B) X { //glsl:*(float,float)
	return gpu.NewFloatExpression(gpu.Op(gpu.NewFloat(a), "*", gpu.NewFloat(b)))
}
func Div[A, B gpu.AnyFloat](a A, b B) X { //glsl:/(float,float)
	return gpu.NewFloatExpression(gpu.Op(gpu.NewFloat(a), "/", gpu.NewFloat(b)))
}
func Neg[A gpu.AnyFloat](a A) X { //glsl:-(float)
	return gpu.NewFloatExpression(gpu.Op(nil, "-", gpu.NewFloat(a)))
}

func Eq[A, B gpu.AnyFloat](a A, b B) gpu.Bool {
	return gpu.NewBoolExpression(gpu.Op(gpu.NewFloat(a), "==", gpu.NewFloat(b)))
}
func Neq[A, B gpu.AnyFloat](a A, b B) gpu.Bool {
	return gpu.NewBoolExpression(gpu.Op(gpu.NewFloat(a), "!=", gpu.NewFloat(b)))
}
func Gt[A, B gpu.AnyFloat](a A, b B) gpu.Bool {
	return gpu.NewBoolExpression(gpu.Op(gpu.NewFloat(a), ">", gpu.NewFloat(b)))
}
func Gte[A, B gpu.AnyFloat](a A, b B) gpu.Bool {
	return gpu.NewBoolExpression(gpu.Op(gpu.NewFloat(a), ">=", gpu.NewFloat(b)))
}
func Lt[A, B gpu.AnyFloat](a A, b B) gpu.Bool {
	return gpu.NewBoolExpression(gpu.Op(gpu.NewFloat(a), "<", gpu.NewFloat(b)))
}
func Lte[A, B gpu.AnyFloat](a A, b B) gpu.Bool {
	return gpu.NewBoolExpression(gpu.Op(gpu.NewFloat(a), "<=", gpu.NewFloat(b)))
}

func BitsToInt[A gpu.AnyFloat](a A) gpu.Int { //glsl:floatBitsToInt(float)ivec
	return gpu.NewIntExpression(gpu.Fn("floatBitsToInt", gpu.NewFloat(a)))
}
func BitsToUint[A gpu.AnyFloat](a A) gpu.Uint { //glsl:floatBitsToUint(float)uvec
	return gpu.NewUintExpression(gpu.Fn("floatBitsToUint", gpu.NewFloat(a)))
}

func Abs[A gpu.AnyFloat](a A) X   { return gpu.NewFloatExpression(gpu.Fn("abs", gpu.NewFloat(a))) }   //glsl:abs(float)float
func Acos[A gpu.AnyFloat](a A) X  { return gpu.NewFloatExpression(gpu.Fn("acos", gpu.NewFloat(a))) }  //glsl:acos(float)float
func Acosh[A gpu.AnyFloat](a A) X { return gpu.NewFloatExpression(gpu.Fn("acosh", gpu.NewFloat(a))) } //glsl:acosh(float)float
func Asin[A gpu.AnyFloat](a A) X  { return gpu.NewFloatExpression(gpu.Fn("asin", gpu.NewFloat(a))) }  //glsl:asin(float)float
func Asinh[A gpu.AnyFloat](a A) X { return gpu.NewFloatExpression(gpu.Fn("asinh", gpu.NewFloat(a))) } //glsl:asinh(float)float
func Atan2[A, B gpu.AnyFloat](a A, b B) X { //glsl:atan(float,float)float
	return gpu.NewFloatExpression(gpu.Fn("atan", gpu.NewFloat(a), gpu.NewFloat(b)))
}
func Atan[A gpu.AnyFloat](a A) X  { return gpu.NewFloatExpression(gpu.Fn("atan", gpu.NewFloat(a))) }  //glsl:atan(float)float
func Atanh[A gpu.AnyFloat](a A) X { return gpu.NewFloatExpression(gpu.Fn("atanh", gpu.NewFloat(a))) } //glsl:atanh(float)float
func Ceil[A gpu.AnyFloat](a A) X  { return gpu.NewFloatExpression(gpu.Fn("ceil", gpu.NewFloat(a))) }  //glsl:ceil(float)float
func Clamp[A, B, C gpu.AnyFloat](a A, b B, c C) X { //glsl:clamp(float,float,float)float
	return gpu.NewFloatExpression(gpu.Fn("clamp", gpu.NewFloat(a), gpu.NewFloat(b), gpu.NewFloat(c)))
}
func Cos[A gpu.AnyFloat](a A) X  { return gpu.NewFloatExpression(gpu.Fn("cos", gpu.NewFloat(a))) }  //glsl:cos(float)float
func Cosh[A gpu.AnyFloat](a A) X { return gpu.NewFloatExpression(gpu.Fn("cosh", gpu.NewFloat(a))) } //glsl:cosh(float)float
func DFdx[A gpu.AnyFloat](a A) X { return gpu.NewFloatExpression(gpu.Fn("dFdx", gpu.NewFloat(a))) } //glsl:dFdx(float)float
func DFdy[A gpu.AnyFloat](a A) X { return gpu.NewFloatExpression(gpu.Fn("dFdy", gpu.NewFloat(a))) } //glsl:dFdy(float)float
func Degrees[A gpu.AnyFloat](a A) X { //glsl:degrees(float)float
	return gpu.NewFloatExpression(gpu.Fn("degrees", gpu.NewFloat(a)))
}
func Exp[A gpu.AnyFloat](a A) X  { return gpu.NewFloatExpression(gpu.Fn("exp", gpu.NewFloat(a))) }  //glsl:exp(float)float
func Exp2[A gpu.AnyFloat](a A) X { return gpu.NewFloatExpression(gpu.Fn("exp2", gpu.NewFloat(a))) } //glsl:exp2(float)float
func FaceFoward[A, B, C gpu.AnyFloat](a A, b B, c C) X { //glsl:faceforward(float,float,float)float
	return gpu.NewFloatExpression(gpu.Fn("faceforward", gpu.NewFloat(a), gpu.NewFloat(b), gpu.NewFloat(c)))
}
func Floor[A gpu.AnyFloat](a A) X  { return gpu.NewFloatExpression(gpu.Fn("floor", gpu.NewFloat(a))) }  //glsl:floor(float)float
func Fract[A gpu.AnyFloat](a A) X  { return gpu.NewFloatExpression(gpu.Fn("fract", gpu.NewFloat(a))) }  //glsl:fract(float)float
func Fwidth[A gpu.AnyFloat](a A) X { return gpu.NewFloatExpression(gpu.Fn("fwidth", gpu.NewFloat(a))) } //glsl:fwidth(float)float
func InverseSqrt[A gpu.AnyFloat](a A) X { //glsl:inversesqrt(float)float
	return gpu.NewFloatExpression(gpu.Fn("inverseSqrt", gpu.NewFloat(a)))
}
func Log[A gpu.AnyFloat](a A) X  { return gpu.NewFloatExpression(gpu.Fn("log", gpu.NewFloat(a))) }  //glsl:log(float)float
func Log2[A gpu.AnyFloat](a A) X { return gpu.NewFloatExpression(gpu.Fn("log2", gpu.NewFloat(a))) } //glsl:log2(float)float
func Max[A, B gpu.AnyFloat](a A, b B) X { //glsl:max(float,float)float
	return gpu.NewFloatExpression(gpu.Fn("max", gpu.NewFloat(a), gpu.NewFloat(b)))
}
func Min[A, B gpu.AnyFloat](a A, b B) X { //glsl:min(float,float)float
	return gpu.NewFloatExpression(gpu.Fn("min", gpu.NewFloat(a), gpu.NewFloat(b)))
}
func Mix[A, B, C gpu.AnyFloat](a A, b B, c C) X { //glsl:mix(float,float,float)float
	return gpu.NewFloatExpression(gpu.Fn("mix", gpu.NewFloat(a), gpu.NewFloat(b), gpu.NewFloat(c)))
}
func Mod[A, B gpu.AnyFloat](a A, b B) X { //glsl:mod(float,float)float
	return gpu.NewFloatExpression(gpu.Fn("mod", gpu.NewFloat(a), gpu.NewFloat(b)))
}
func Modf(a X) (X, X) { //glsl:modf(float,out[float])float
	out := gpu.NewFloatExpression(gpu.Out("float"))
	return out, gpu.NewFloatExpression(gpu.Fn("modf", a, out))
}
func Normalize(a X) X { //glsl:normalize(float)float
	return gpu.NewFloatExpression(gpu.Fn("normalize", a))
}
func Pow[A, B gpu.AnyFloat](a A, b B) X { //glsl:pow(float,float)float
	return gpu.NewFloatExpression(gpu.Fn("pow", gpu.NewFloat(a), gpu.NewFloat(b)))
}
func Radians[A gpu.AnyFloat](a A) X { //glsl:radians(float)float
	return gpu.NewFloatExpression(gpu.Fn("radians", gpu.NewFloat(a)))
}
func Reflect[A, B gpu.AnyFloat](a A, b B) X { //glsl:reflect(float,float)float
	return gpu.NewFloatExpression(gpu.Fn("reflect", gpu.NewFloat(a), gpu.NewFloat(b)))
}
func Refract[A, B, C gpu.AnyFloat](a A, b B, c C) X { //glsl:refract(float,float,float)float
	return gpu.NewFloatExpression(gpu.Fn("refract", gpu.NewFloat(a), gpu.NewFloat(b), gpu.NewFloat(c)))
}
func Round[A gpu.AnyFloat](a A) X { //glsl:round(float)float
	return gpu.NewFloatExpression(gpu.Fn("round", gpu.NewFloat(a)))
}
func RoundEven[A gpu.AnyFloat](a A) X { //glsl:roundEven(float)float
	return gpu.NewFloatExpression(gpu.Fn("roundeven", gpu.NewFloat(a)))
}
func Sign[A gpu.AnyFloat](a A) X { //glsl:sign(float)float
	return gpu.NewFloatExpression(gpu.Fn("sign", gpu.NewFloat(a)))
}
func Sin[A gpu.AnyFloat](a A) X  { return gpu.NewFloatExpression(gpu.Fn("sin", gpu.NewFloat(a))) }  //glsl:sin(float)float
func Sinh[A gpu.AnyFloat](a A) X { return gpu.NewFloatExpression(gpu.Fn("sinh", gpu.NewFloat(a))) } //glsl:sinh(float)float
func SmoothStep[A, B, C gpu.AnyFloat](a A, b B, c C) X { //glsl:smoothstep(float,float,float)float
	return gpu.NewFloatExpression(gpu.Fn("smoothstep", gpu.NewFloat(a), gpu.NewFloat(b), gpu.NewFloat(c)))
}
func Sqrt[A gpu.AnyFloat](a A) X { return gpu.NewFloatExpression(gpu.Fn("sqrt", gpu.NewFloat(a))) } //glsl:sqrt(float)float
func Step[A, B gpu.AnyFloat](a A, b B) X { //glsl:step(float,float)float
	return gpu.NewFloatExpression(gpu.Fn("step", gpu.NewFloat(a), gpu.NewFloat(b)))
}
func Tan[A gpu.AnyFloat](a A) X   { return gpu.NewFloatExpression(gpu.Fn("tan", gpu.NewFloat(a))) }   //glsl:tan(float)float
func Tanh[A gpu.AnyFloat](a A) X  { return gpu.NewFloatExpression(gpu.Fn("tanh", gpu.NewFloat(a))) }  //glsl:tanh(float)float
func Trunc[A gpu.AnyFloat](a A) X { return gpu.NewFloatExpression(gpu.Fn("trunc", gpu.NewFloat(a))) } //glsl:trunc(float)float
func IsInf[A gpu.AnyFloat](a A) gpu.Bool { //glsl:isinf(float)bvec
	return gpu.NewBoolExpression(gpu.Fn("isinf", gpu.NewFloat(a)))
}
func IsNaN[A gpu.AnyFloat](a A) gpu.Bool { //glsl:isnan(float)bvec
	return gpu.NewBoolExpression(gpu.Fn("isnan", gpu.NewFloat(a)))
}
func Distance[A, B gpu.AnyFloat](a A, b B) X { //glsl:distance(float,float)float
	return gpu.NewFloatExpression(gpu.Fn("distance", gpu.NewFloat(a), gpu.NewFloat(b)))
}
func Dot[A, B gpu.AnyFloat](a A, b B) X { //glsl:dot(float,float)float
	return gpu.NewFloatExpression(gpu.Fn("dot", gpu.NewFloat(a), gpu.NewFloat(b)))
}
func Length[A gpu.AnyFloat](a A) X { return gpu.NewFloatExpression(gpu.Fn("length", gpu.NewFloat(a))) } //glsl:length(float)float
