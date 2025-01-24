// Pacakge bvec2 provides GPU operations on two-component boolean vectors.
package bvec2

import "graphics.gd/shaders/internal/gpu"

// XY is a two-component vector of booleans on the GPU.
type XY gpu.Vec2b

// New creates a new [XY] value from two booleans.
func New[X, Y gpu.AnyBool](x X, y Y) XY { return gpu.NewVec2b(x, y) }

func And(a, b XY) XY { return gpu.NewVec2bExpression(gpu.Op(a, "&&", b)) }
func Or(a, b XY) XY  { return gpu.NewVec2bExpression(gpu.Op(a, "||", b)) }
func Eq(a, b XY) XY  { return gpu.NewVec2bExpression(gpu.Op(a, "==", b)) }
func Neq(a, b XY) XY { return gpu.NewVec2bExpression(gpu.Op(a, "!=", b)) }

func Mix(a, b gpu.Vec2, t XY) gpu.Vec2 { //glsl:mix(vec2,vec2,bvec)vec2
	return gpu.NewVec2Expression(gpu.Fn("mix", a, b, t))
}
func All(a XY) gpu.Bool { return gpu.NewBoolExpression(gpu.Fn("all", a)) }  //glsl:all(bvec2)bool
func Any(a XY) gpu.Bool { return gpu.NewBoolExpression(gpu.Fn("any", a)) }  //glsl:any(bvec2)bool
func Not(a XY) XY       { return gpu.NewVec2bExpression(gpu.Fn("not", a)) } //glsl:not(bvec2)bvec2
