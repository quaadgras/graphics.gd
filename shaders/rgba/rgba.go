// Package rgba provides a constructor for vec4.RGBA values.
package rgba

import (
	"reflect"

	"graphics.gd/shaders/internal/gpu"
	"graphics.gd/shaders/vec4"
)

func New[R, G, B, A gpu.AnyFloat](r R, g G, b B, a A) vec4.RGBA {
	return gpu.NewRGBA(r, g, b, a)
}

func Add[T gpu.AnyFloat | vec4.RGBA](a vec4.RGBA, b T) vec4.RGBA { //glsl:+(vec4,vec4)
	return gpu.NewRGBAExpression(gpu.Op(either(a), "+", either(b)))
}
func Sub[T gpu.AnyFloat | vec4.RGBA](a vec4.RGBA, b T) vec4.RGBA { //glsl:-(vec4,vec4)
	return gpu.NewRGBAExpression(gpu.Op(either(a), "-", either(b)))
}
func Mul[T gpu.AnyFloat | vec4.RGBA](a vec4.RGBA, b T) vec4.RGBA { //glsl:*(vec4,vec4)
	return gpu.NewRGBAExpression(gpu.Op(either(a), "*", either(b)))
}
func Div[T gpu.AnyFloat | vec4.RGBA](a vec4.RGBA, b T) vec4.RGBA { //glsl:/(vec4,vec4)
	return gpu.NewRGBAExpression(gpu.Op(either(a), "/", either(b)))
}
func Neg(a vec4.RGBA) vec4.RGBA { return gpu.NewRGBAExpression(gpu.Op(nil, "-", a)) } //glsl:-(vec4)

func either[T gpu.AnyFloat | vec4.RGBA](v T) gpu.Evaluator {
	rvalue := reflect.ValueOf(v)
	switch {
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(vec4.RGBA{})):
		return rvalue.Convert(reflect.TypeOf(vec4.RGBA{})).Interface().(vec4.RGBA)
	case rvalue.Type().ConvertibleTo(reflect.TypeOf(gpu.Float{})):
		return gpu.NewFloat(rvalue.Convert(reflect.TypeOf(gpu.Float{})).Interface().(gpu.Float))
	default:
		return gpu.NewFloat(rvalue.Float())
	}
}
