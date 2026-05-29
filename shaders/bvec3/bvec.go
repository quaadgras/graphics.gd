// Pacakge bvec3 provides GPU operations on three-component boolean vectors.
package bvec3

import "graphics.gd/shaders/internal/gpu"

// XYZ is a three-component vector of booleans on the GPU.
type XYZ gpu.Vec3b

// New creates a new [XY] value from three booleans.
func New[X, Y, Z gpu.AnyBool](x X, y Y, z Z) XYZ { return gpu.NewVec3b(x, y, z) }

func And(a, b XYZ) XYZ { return gpu.NewVec3bExpression(gpu.Op(a, "&&", b)) }
func Or(a, b XYZ) XYZ  { return gpu.NewVec3bExpression(gpu.Op(a, "||", b)) }
func Eq(a, b XYZ) XYZ  { return gpu.NewVec3bExpression(gpu.Op(a, "==", b)) }
func Neq(a, b XYZ) XYZ { return gpu.NewVec3bExpression(gpu.Op(a, "!=", b)) }

func Mix(a, b gpu.Vec3, t XYZ) gpu.Vec3 { //glsl:mix(vec3,vec3,bvec)vec3
	return gpu.NewVec3Expression(gpu.Fn("mix", a, b, t))
}
func All(a XYZ) gpu.Bool { return gpu.NewBoolExpression(gpu.Fn("all", a)) }  //glsl:all(bvec3)bool
func Any(a XYZ) gpu.Bool { return gpu.NewBoolExpression(gpu.Fn("any", a)) }  //glsl:any(bvec3)bool
func Not(a XYZ) XYZ      { return gpu.NewVec3bExpression(gpu.Fn("not", a)) } //glsl:not(bvec3)bvec3
