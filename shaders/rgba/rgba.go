// Package rgba provides a constructor for vec4.RGBA values.
package rgba

import (
	"graphics.gd/shaders/internal/dsl"
	"graphics.gd/shaders/vec1"
	"graphics.gd/shaders/vec4"
)

type float interface {
	float64 | int | uint | vec1.X | vec1.T[int] | vec1.T[uint]
}

func toVec1[T float](x T) vec1.X {
	switch x := any(x).(type) {
	case float64:
		return vec1.New(float64(x))
	case int:
		return vec1.New(float64(x))
	case uint:
		return vec1.New(float64(x))
	case vec1.X:
		return x
	case vec1.T[int]:
		return vec1.X{
			Expression: dsl.Fn("float", x),
		}
	case vec1.T[uint]:
		return vec1.X{
			Expression: dsl.Fn("float", x),
		}
	default:
		panic("unreachable")
	}
}

func New[R, G, B, A float](r R, g G, b B, a A) vec4.RGBA {
	return vec4.RGBA{
		R: toVec1(r),
		G: toVec1(g),
		B: toVec1(b),
		A: toVec1(a),
	}
}

func Add(a, b vec4.RGBA) vec4.RGBA { return vec4.RGBA{Expression: dsl.Op(a, "+", b)} }
func Sub(a, b vec4.RGBA) vec4.RGBA { return vec4.RGBA{Expression: dsl.Op(a, "-", b)} }
func Mul(a, b vec4.RGBA) vec4.RGBA { return vec4.RGBA{Expression: dsl.Op(a, "*", b)} }
func Div(a, b vec4.RGBA) vec4.RGBA { return vec4.RGBA{Expression: dsl.Op(a, "/", b)} }

func AddX(a vec4.RGBA, b vec1.X) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Op(a, "+", b)}
}

func SubX(a vec4.RGBA, b vec1.X) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Op(a, "-", b)}
}

func MulX(a vec4.RGBA, b vec1.X) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Op(a, "*", b)}
}

func DivX(a vec4.RGBA, b vec1.X) vec4.RGBA {
	return vec4.RGBA{Expression: dsl.Op(a, "/", b)}
}
