// Package rgba provides a constructor for vec4.RGBA values.
package rgba

import (
	"graphics.gd/shaders/internal/gpu"
	"graphics.gd/shaders/vec4"
)

func New[R, G, B, A gpu.AnyFloat](r R, g G, b B, a A) vec4.RGBA {
	return gpu.NewRGBA(r, g, b, a)
}

/*func Add(a, b vec4.RGBA) vec4.RGBA { return vec4.RGBA{Expression: dsl.Op(a, "+", b)} }
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
	}*/
