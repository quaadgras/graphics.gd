package vec4

import (
	"graphics.gd/shaders/internal/dsl"
	"graphics.gd/shaders/vec1"
)

type XYZW = T[float64]

type T[T bool | int | uint | float64] struct {
	dsl.Expression

	X vec1.T[T]
	Y vec1.T[T]
	Z vec1.T[T]
	W vec1.T[T]
}

type RGBA struct {
	dsl.Expression

	R vec1.X
	G vec1.X
	B vec1.X
	A vec1.X
}

func New[X vec1.Any](x, y, z, w X) T[X] {
	return T[X]{
		X: vec1.New(x),
		Y: vec1.New(y),
		Z: vec1.New(z),
		W: vec1.New(w),
	}
}

func Add[X int | uint | float64](a, b T[X]) T[X] { return T[X]{Expression: dsl.Op(a, "+", b)} }
func Sub[X int | uint | float64](a, b T[X]) T[X] { return T[X]{Expression: dsl.Op(a, "-", b)} }
func Mul[X int | uint | float64](a, b T[X]) T[X] { return T[X]{Expression: dsl.Op(a, "*", b)} }
func Div[X int | uint | float64](a, b T[X]) T[X] { return T[X]{Expression: dsl.Op(a, "/", b)} }
func Mod[X int | uint](a, b T[X]) T[X]           { return T[X]{Expression: dsl.Op(a, "%", b)} }

func Not(a T[bool]) T[bool]                         { return T[bool]{Expression: dsl.Op(nil, "!", a)} }
func And(a, b T[bool]) T[bool]                      { return T[bool]{Expression: dsl.Op(a, "&&", b)} }
func Or(a, b T[bool]) T[bool]                       { return T[bool]{Expression: dsl.Op(a, "||", b)} }
func Eq[X vec1.Any](a, b T[X]) T[bool]              { return T[bool]{Expression: dsl.Op(a, "==", b)} }
func Neq[X vec1.Any](a, b T[X]) T[bool]             { return T[bool]{Expression: dsl.Op(a, "!=", b)} }
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

func AddX[X int | uint | float64](a T[X], b vec1.T[X]) T[X] {
	return T[X]{Expression: dsl.Op(a, "+", b)}
}

func SubX[X int | uint | float64](a T[X], b vec1.T[X]) T[X] {
	return T[X]{Expression: dsl.Op(a, "-", b)}
}

func MulX[X int | uint | float64](a T[X], b vec1.T[X]) T[X] {
	return T[X]{Expression: dsl.Op(a, "*", b)}
}

func DivX[X int | uint | float64](a T[X], b vec1.T[X]) T[X] {
	return T[X]{Expression: dsl.Op(a, "/", b)}
}

func ModX[X int | uint](a T[X], b vec1.T[X]) T[X] {
	return T[X]{Expression: dsl.Op(a, "%", b)}
}

func PackUnorm4x8(v XYZW) vec1.X {
	return vec1.X{Expression: dsl.Fn("packUnorm4x8", v)}
}
func UnpackUnorm4x8(v vec1.X) XYZW {
	return XYZW{Expression: dsl.Fn("unpackUnorm4x8", v)}
}
func PackSnorm4x8(v XYZW) vec1.X {
	return vec1.X{Expression: dsl.Fn("packSnorm4x8", v)}
}
func UnpackSnorm4x8(v vec1.X) XYZW {
	return XYZW{Expression: dsl.Fn("unpackSnorm4x8", v)}
}

func Sin(x T[float64]) T[float64]                { return T[float64]{Expression: dsl.Fn("sin", x)} }
func Cos(x T[float64]) T[float64]                { return T[float64]{Expression: dsl.Fn("cos", x)} }
func Tan(x T[float64]) T[float64]                { return T[float64]{Expression: dsl.Fn("tan", x)} }
func Asin(x T[float64]) T[float64]               { return T[float64]{Expression: dsl.Fn("asin", x)} }
func Acos(x T[float64]) T[float64]               { return T[float64]{Expression: dsl.Fn("acos", x)} }
func Atan(x T[float64]) T[float64]               { return T[float64]{Expression: dsl.Fn("atan", x)} }
func Atan2(y, x T[float64]) T[float64]           { return T[float64]{Expression: dsl.Fn("atan", y, x)} }
func Sinh(x T[float64]) T[float64]               { return T[float64]{Expression: dsl.Fn("sinh", x)} }
func Cosh(x T[float64]) T[float64]               { return T[float64]{Expression: dsl.Fn("cosh", x)} }
func Tanh(x T[float64]) T[float64]               { return T[float64]{Expression: dsl.Fn("tanh", x)} }
func Asinh(x T[float64]) T[float64]              { return T[float64]{Expression: dsl.Fn("asinh", x)} }
func Acosh(x T[float64]) T[float64]              { return T[float64]{Expression: dsl.Fn("acosh", x)} }
func Atanh(x T[float64]) T[float64]              { return T[float64]{Expression: dsl.Fn("atanh", x)} }
func Pow(x, y T[float64]) T[float64]             { return T[float64]{Expression: dsl.Fn("pow", x, y)} }
func Exp(x T[float64]) T[float64]                { return T[float64]{Expression: dsl.Fn("exp", x)} }
func Exp2(x T[float64]) T[float64]               { return T[float64]{Expression: dsl.Fn("exp2", x)} }
func Log(x T[float64]) T[float64]                { return T[float64]{Expression: dsl.Fn("log", x)} }
func Log2(x T[float64]) T[float64]               { return T[float64]{Expression: dsl.Fn("log2", x)} }
func Sqrt(x T[float64]) T[float64]               { return T[float64]{Expression: dsl.Fn("sqrt", x)} }
func Inversesqrt(x T[float64]) T[float64]        { return T[float64]{Expression: dsl.Fn("inversesqrt", x)} }
func Abs[X int | float64](x T[X]) T[X]           { return T[X]{Expression: dsl.Fn("abs", x)} }
func Sign[X int | float64](x T[X]) T[X]          { return T[X]{Expression: dsl.Fn("sign", x)} }
func Floor(x T[float64]) T[float64]              { return T[float64]{Expression: dsl.Fn("floor", x)} }
func Round(x T[float64]) T[float64]              { return T[float64]{Expression: dsl.Fn("round", x)} }
func RoundEven(x T[float64]) T[float64]          { return T[float64]{Expression: dsl.Fn("roundEven", x)} }
func Trunc(x T[float64]) T[float64]              { return T[float64]{Expression: dsl.Fn("trunc", x)} }
func Ceil(x T[float64]) T[float64]               { return T[float64]{Expression: dsl.Fn("ceil", x)} }
func Fract(x T[float64]) T[float64]              { return T[float64]{Expression: dsl.Fn("fract", x)} }
func Min[X uint | int | float64](a, b T[X]) T[X] { return T[X]{Expression: dsl.Fn("min", a, b)} }
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
func IsNaN(x T[float64]) T[bool] { return T[bool]{Expression: dsl.Fn("isnan", x)} }
func IsInf(x T[float64]) T[bool] { return T[bool]{Expression: dsl.Fn("isinf", x)} }
func FloatBitsToInt(x T[float64]) T[int] {
	return T[int]{Expression: dsl.Fn("floatBitsToInt", x)}
}
func FloatBitsToUint(x T[float64]) T[uint] {
	return T[uint]{Expression: dsl.Fn("floatBitsToUint", x)}
}
func IntBitsToFloat(x T[int]) T[float64] {
	return T[float64]{Expression: dsl.Fn("intBitsToFloat", x)}
}
func UintBitsToFloat(x T[uint]) T[float64] {
	return T[float64]{Expression: dsl.Fn("uintBitsToFloat", x)}
}
func Length(x T[float64]) T[float64] {
	return T[float64]{Expression: dsl.Fn("length", x)}
}
func Distance(a, b T[float64]) T[float64] {
	return T[float64]{Expression: dsl.Fn("distance", a, b)}
}
func Normalize(x T[float64]) T[float64] {
	return T[float64]{Expression: dsl.Fn("normalize", x)}
}
func Reflect(i, n T[float64]) T[float64] {
	return T[float64]{Expression: dsl.Fn("reflect", i, n)}
}
func Refract(i, n T[float64], eta T[float64]) T[float64] {
	return T[float64]{Expression: dsl.Fn("refract", i, n, eta)}
}
func Faceforward(n, i, nref T[float64]) T[float64] {
	return T[float64]{Expression: dsl.Fn("faceforward", n, i, nref)}
}
func DFdx(x T[float64]) T[float64] {
	return T[float64]{Expression: dsl.Fn("dFdx", x)}
}
func DFdxCoarse(x T[float64]) T[float64] {
	return T[float64]{Expression: dsl.Fn("dFdxCoarse", x)}
}
func DFdxFine(x T[float64]) T[float64] {
	return T[float64]{Expression: dsl.Fn("dFdxFine", x)}
}
func DFdy(x T[float64]) T[float64] {
	return T[float64]{Expression: dsl.Fn("dFdy", x)}
}
func DFdyCoarse(x T[float64]) T[float64] {
	return T[float64]{Expression: dsl.Fn("dFdyCoarse", x)}
}
func DFdyFine(x T[float64]) T[float64] {
	return T[float64]{Expression: dsl.Fn("dFdyFine", x)}
}
func Fwidth(x T[float64]) T[float64] {
	return T[float64]{Expression: dsl.Fn("fwidth", x)}
}
func FwidthCoarse(x T[float64]) T[float64] {
	return T[float64]{Expression: dsl.Fn("fwidthCoarse", x)}
}
func FwidthFine(x T[float64]) T[float64] {
	return T[float64]{Expression: dsl.Fn("fwidthFine", x)}
}
func BitfieldExtract[X int | uint](value T[X], offset, bits vec1.T[X]) T[X] {
	return T[X]{Expression: dsl.Fn("bitfieldExtract", value, offset, bits)}
}
func BitfieldInsert[X int | uint](base, insert T[X], offset, bits vec1.T[X]) T[X] {
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

func Dot(a, b XYZW) XYZW {
	return XYZW{Expression: dsl.Fn("dot", a, b)}
}
