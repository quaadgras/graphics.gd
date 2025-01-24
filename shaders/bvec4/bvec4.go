// Pacakge bvec4 provides GPU operations on four-component boolean vectors.
package bvec4

import "graphics.gd/shaders/internal/gpu"

// XYZW is a four-component vector of booleans on the GPU.
type XYZW gpu.Vec4b

// New creates a new [XY] value from three booleans.
func New[X, Y, Z, W gpu.AnyBool](x X, y Y, z Z, w W) XYZW { return gpu.NewVec4b(x, y, z, w) }

func And(a, b XYZW) XYZW { return gpu.NewVec4bExpression(gpu.Op(a, "&&", b)) }
func Or(a, b XYZW) XYZW  { return gpu.NewVec4bExpression(gpu.Op(a, "||", b)) }
func Eq(a, b XYZW) XYZW  { return gpu.NewVec4bExpression(gpu.Op(a, "==", b)) }
func Neq(a, b XYZW) XYZW { return gpu.NewVec4bExpression(gpu.Op(a, "!=", b)) }

func Mix(a, b gpu.Vec4, t XYZW) gpu.Vec4 { //glsl:mix(vec4,vec4,bvec)vec4
	return gpu.NewVec4Expression(gpu.Fn("mix", a, b, t))
}
func All(a XYZW) gpu.Bool { return gpu.NewBoolExpression(gpu.Fn("all", a)) }  //glsl:all(bvec4)bool
func Any(a XYZW) gpu.Bool { return gpu.NewBoolExpression(gpu.Fn("any", a)) }  //glsl:any(bvec4)bool
func Not(a XYZW) XYZW     { return gpu.NewVec4bExpression(gpu.Fn("not", a)) } //glsl:not(bvec4)bvec4
