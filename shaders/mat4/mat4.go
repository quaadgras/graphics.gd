// Package mat4 provides GPU operations on 4x4 matrices.
package mat4

import (
	"graphics.gd/shaders/internal/gpu"
	"graphics.gd/shaders/vec4"
)

type ColumnMajor gpu.Mat4

// New returns a matrix with the given columns.
func New(x, y, z, w vec4.XYZW) ColumnMajor {
	return ColumnMajor{Columns: [4][4]gpu.Float{{x.X, x.Y, x.Z, x.W}, {y.X, y.Y, y.Z, y.W}, {z.X, z.Y, z.Z, z.W}, {w.X, w.Y, w.Z, w.W}}}
}

// Transform returns the vector v transformed by the matrix m.
func Transform(m ColumnMajor, v vec4.XYZW) vec4.XYZW { //glsl:*(mat4,vec4)
	return gpu.NewVec4Expression(gpu.Op(m, "*", v))
}

// Column returns column i of the matrix m.
func Column(m ColumnMajor, i int) vec4.XYZW { //glsl:mat4[int]vec4
	return gpu.NewVec4Expression(gpu.New(gpu.Index{Value: m, Index: i}))
}

func ComponentMul(a, b ColumnMajor) ColumnMajor { //glsl:matrixCompMult(mat4,mat4)mat4
	return gpu.NewMat4Expression(gpu.Fn("matrixCompMult", a, b))
}

func OuterProduct(a, b gpu.Vec4) ColumnMajor { //glsl:outerProduct(vec4,vec4)mat4
	return gpu.NewMat4Expression(gpu.Fn("outerProduct", a, b))
}
func Transpose(a ColumnMajor) ColumnMajor { //glsl:transpose(mat4)mat4
	return gpu.NewMat4Expression(gpu.Fn("transpose", a))
}

func Determinant(a ColumnMajor) gpu.Float { //glsl:determinant(mat4)float
	return gpu.NewFloatExpression(gpu.Fn("determinant", a))
}
func Inverse(a ColumnMajor) ColumnMajor { //glsl:inverse(mat4)mat4
	return gpu.NewMat4Expression(gpu.Fn("inverse", a))
}

func Mul(a, b ColumnMajor) ColumnMajor { //glsl:*(mat4,mat4)
	return gpu.NewMat4Expression(gpu.Op(a, "*", b))
}
