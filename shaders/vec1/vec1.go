package vec1

import "graphics.gd/shaders/internal/dsl"

type X = T[float64]

type T[T bool | int | uint | float64] struct {
	dsl.Expression

	X T
}

type Any interface {
	float64 | int | bool | uint
}

func New[X Any](x X) T[X] {
	return T[X]{X: x}
}
func Add[X int | uint | float64](a, b T[X]) T[X] { return T[X]{Expression: dsl.Op(a, "+", b)} }
func Sub[X int | uint | float64](a, b T[X]) T[X] { return T[X]{Expression: dsl.Op(a, "-", b)} }
func Mul[X int | uint | float64](a, b T[X]) T[X] { return T[X]{Expression: dsl.Op(a, "*", b)} }
func Div[X int | uint | float64](a, b T[X]) T[X] { return T[X]{Expression: dsl.Op(a, "/", b)} }
func Mod[X int | uint](a, b T[X]) T[X]           { return T[X]{Expression: dsl.Op(a, "%", b)} }

func Not(a T[bool]) T[bool]                         { return T[bool]{Expression: dsl.Op(nil, "!", a)} }
func And(a, b T[bool]) T[bool]                      { return T[bool]{Expression: dsl.Op(a, "&&", b)} }
func Or(a, b T[bool]) T[bool]                       { return T[bool]{Expression: dsl.Op(a, "||", b)} }
func Eq[X Any](a, b T[X]) T[bool]                   { return T[bool]{Expression: dsl.Op(a, "==", b)} }
func Neq[X Any](a, b T[X]) T[bool]                  { return T[bool]{Expression: dsl.Op(a, "!=", b)} }
func Gt[X int | uint | float64](a, b T[X]) T[bool]  { return T[bool]{Expression: dsl.Op(a, ">", b)} }
func Gte[X int | uint | float64](a, b T[X]) T[bool] { return T[bool]{Expression: dsl.Op(a, ">=", b)} }
func Lt[X int | uint | float64](a, b T[X]) T[bool]  { return T[bool]{Expression: dsl.Op(a, "<", b)} }
func Lte[X int | uint | float64](a, b T[X]) T[bool] { return T[bool]{Expression: dsl.Op(a, "<=", b)} }

func Neg[X int | uint | float64](a T[X]) T[X] { return T[X]{Expression: dsl.Op(nil, "-", a)} }

func BitwiseNot[X int | uint](a T[X]) T[X]           { return T[X]{Expression: dsl.Op(nil, "~", a)} }
func BitwiseAnd[X int | uint](a, b T[X]) T[X]        { return T[X]{Expression: dsl.Op(a, "&", b)} }
func BitwiseOr[X int | uint](a, b T[X]) T[X]         { return T[X]{Expression: dsl.Op(a, "|", b)} }
func BitwiseXor[X int | uint](a, b T[X]) T[X]        { return T[X]{Expression: dsl.Op(a, "^", b)} }
func BitwiseShiftLeft[X int | uint](a, b T[X]) T[X]  { return T[X]{Expression: dsl.Op(a, "<<", b)} }
func BitwiseShiftRight[X int | uint](a, b T[X]) T[X] { return T[X]{Expression: dsl.Op(a, ">>", b)} }

func If[X Any](cond T[bool], a, b T[X]) T[X] {
	return T[X]{
		Expression: dsl.New(dsl.Ternary{
			If: cond,
			A:  a,
			B:  b,
		}),
	}
}

func Radians(degrees X) X                        { return X{Expression: dsl.Fn("radians", degrees)} }
func Degrees(radians X) X                        { return X{Expression: dsl.Fn("degrees", radians)} }
func Sin(x X) X                                  { return X{Expression: dsl.Fn("sin", x)} }
func Cos(x X) X                                  { return X{Expression: dsl.Fn("cos", x)} }
func Tan(x X) X                                  { return X{Expression: dsl.Fn("tan", x)} }
func Asin(x X) X                                 { return X{Expression: dsl.Fn("asin", x)} }
func Acos(x X) X                                 { return X{Expression: dsl.Fn("acos", x)} }
func Atan(x X) X                                 { return X{Expression: dsl.Fn("atan", x)} }
func Atan2(y, x X) X                             { return X{Expression: dsl.Fn("atan", y, x)} }
func Sinh(x X) X                                 { return X{Expression: dsl.Fn("sinh", x)} }
func Cosh(x X) X                                 { return X{Expression: dsl.Fn("cosh", x)} }
func Tanh(x X) X                                 { return X{Expression: dsl.Fn("tanh", x)} }
func Asinh(x X) X                                { return X{Expression: dsl.Fn("asinh", x)} }
func Acosh(x X) X                                { return X{Expression: dsl.Fn("acosh", x)} }
func Atanh(x X) X                                { return X{Expression: dsl.Fn("atanh", x)} }
func Pow(x, y X) X                               { return X{Expression: dsl.Fn("pow", x, y)} }
func Exp(x X) X                                  { return X{Expression: dsl.Fn("exp", x)} }
func Exp2(x X) X                                 { return X{Expression: dsl.Fn("exp2", x)} }
func Log(x X) X                                  { return X{Expression: dsl.Fn("log", x)} }
func Log2(x X) X                                 { return X{Expression: dsl.Fn("log2", x)} }
func Sqrt(x X) X                                 { return X{Expression: dsl.Fn("sqrt", x)} }
func Inversesqrt(x X) X                          { return X{Expression: dsl.Fn("inversesqrt", x)} }
func Abs[X int | float64](x T[X]) T[X]           { return T[X]{Expression: dsl.Fn("abs", x)} }
func Sign[X int | float64](x T[X]) T[X]          { return T[X]{Expression: dsl.Fn("sign", x)} }
func Floor(x X) X                                { return X{Expression: dsl.Fn("floor", x)} }
func Round(x X) X                                { return X{Expression: dsl.Fn("round", x)} }
func RoundEven(x X) X                            { return X{Expression: dsl.Fn("roundEven", x)} }
func Trunc(x X) X                                { return X{Expression: dsl.Fn("trunc", x)} }
func Ceil(x X) X                                 { return X{Expression: dsl.Fn("ceil", x)} }
func Fract(x X) X                                { return X{Expression: dsl.Fn("fract", x)} }
func Min[X uint | int | float64](a, b T[X]) T[X] { return T[X]{Expression: dsl.Fn("min", a, b)} }
func Modf(a, b X) X {
	return X{Expression: dsl.Fn("mod", a, b)}
}
func Max[X uint | int | float64](a, b T[X]) T[X] { return T[X]{Expression: dsl.Fn("max", a, b)} }
func Clamp[X uint | int | float64](x, min, max T[X]) T[X] {
	return T[X]{Expression: dsl.Fn("clamp", x, min, max)}
}
func Mix[X uint | int | float64](a, b, t T[X]) T[X] { return T[X]{Expression: dsl.Fn("mix", a, b, t)} }
func FMA[X uint | int | float64](a, b, c T[X]) T[X] { return T[X]{Expression: dsl.Fn("fma", a, b, c)} }
func Step[X uint | int | float64](edge, x T[X]) T[X] {
	return T[X]{Expression: dsl.Fn("step", edge, x)}
}
func Smoothstep[X uint | int | float64](edge0, edge1, x T[X]) T[X] {
	return T[X]{Expression: dsl.Fn("smoothstep", edge0, edge1, x)}
}
func IsNaN(x X) T[bool] { return T[bool]{Expression: dsl.Fn("isnan", x)} }
func IsInf(x X) T[bool] { return T[bool]{Expression: dsl.Fn("isinf", x)} }
func FloatBitsToInt(x X) T[int] {
	return T[int]{Expression: dsl.Fn("floatBitsToInt", x)}
}
func FloatBitsToUint(x X) T[uint] {
	return T[uint]{Expression: dsl.Fn("floatBitsToUint", x)}
}
func IntBitsToFloat(x T[int]) X {
	return X{Expression: dsl.Fn("intBitsToFloat", x)}
}
func UintBitsToFloat(x T[uint]) X {
	return X{Expression: dsl.Fn("uintBitsToFloat", x)}
}
func Length(x X) X {
	return X{Expression: dsl.Fn("length", x)}
}
func Distance(a, b X) X {
	return X{Expression: dsl.Fn("distance", a, b)}
}
func Normalize(x X) X {
	return X{Expression: dsl.Fn("normalize", x)}
}
func Reflect(i, n X) X {
	return X{Expression: dsl.Fn("reflect", i, n)}
}
func Refract(i, n X, eta X) X {
	return X{Expression: dsl.Fn("refract", i, n, eta)}
}
func Faceforward(n, i, nref X) X {
	return X{Expression: dsl.Fn("faceforward", n, i, nref)}
}
func DFdx(x X) X {
	return X{Expression: dsl.Fn("dFdx", x)}
}
func DFdxCoarse(x X) X {
	return X{Expression: dsl.Fn("dFdxCoarse", x)}
}
func DFdxFine(x X) X {
	return X{Expression: dsl.Fn("dFdxFine", x)}
}
func DFdy(x X) X {
	return X{Expression: dsl.Fn("dFdy", x)}
}
func DFdyCoarse(x X) X {
	return X{Expression: dsl.Fn("dFdyCoarse", x)}
}
func DFdyFine(x X) X {
	return X{Expression: dsl.Fn("dFdyFine", x)}
}
func Fwidth(x X) X {
	return X{Expression: dsl.Fn("fwidth", x)}
}
func FwidthCoarse(x X) X {
	return X{Expression: dsl.Fn("fwidthCoarse", x)}
}
func FwidthFine(x X) X {
	return X{Expression: dsl.Fn("fwidthFine", x)}
}
func BitfieldExtract[X int | uint](value T[X], offset, bits T[X]) T[X] {
	return T[X]{Expression: dsl.Fn("bitfieldExtract", value, offset, bits)}
}
func BitfieldInsert[X int | uint](base, insert T[X], offset, bits T[X]) T[X] {
	return T[X]{Expression: dsl.Fn("bitfieldInsert", base, insert, offset, bits)}
}
func BitfieldReverse[X int | uint](value T[X]) T[X] {
	return T[X]{Expression: dsl.Fn("bitfieldReverse", value)}
}
func BitCount[X int | uint](value T[X]) T[X] {
	return T[X]{Expression: dsl.Fn("bitCount", value)}
}
func FindLSB[X int | uint](value T[X]) T[X] {
	return T[X]{Expression: dsl.Fn("findLSB", value)}
}
func FindMSB[X int | uint](value T[X]) T[X] {
	return T[X]{Expression: dsl.Fn("findMSB", value)}
}
