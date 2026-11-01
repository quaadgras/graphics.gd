// Package mat3 provides GPU operations on 3x3 matrices.
package mat3

import (
	"graphics.gd/shaders/internal/gpu"
	"graphics.gd/shaders/vec3"
)

type ColumnMajor gpu.Mat3

// New returns a matrix with the given columns.
func New(x, y, z vec3.XYZ) ColumnMajor {
	return ColumnMajor{Columns: [3][3]gpu.Float{{x.X, x.Y, x.Z}, {y.X, y.Y, y.Z}, {z.X, z.Y, z.Z}}}
}

// Transform returns the vector v transformed by the matrix m.
func Transform(m ColumnMajor, v vec3.XYZ) vec3.XYZ { //glsl:*(mat3,vec3)
	return gpu.NewVec3Expression(gpu.Op(m, "*", v))
}

// Column returns column i of the matrix m.
func Column(m ColumnMajor, i int) vec3.XYZ { //glsl:mat3[int]vec3
	return gpu.NewVec3Expression(gpu.New(gpu.Index{Value: m, Index: i}))
}

func ComponentMul(a, b ColumnMajor) ColumnMajor { //glsl:matrixCompMult(mat3,mat3)mat3
	return gpu.NewMat3Expression(gpu.Fn("matrixCompMult", a, b))
}
func OuterProduct(a, b gpu.Vec3) ColumnMajor { //glsl:outerProduct(vec3,vec3)mat3
	return gpu.NewMat3Expression(gpu.Fn("outerProduct", a, b))
}
func Transpose(a ColumnMajor) ColumnMajor { //glsl:transpose(mat3)mat3
	return gpu.NewMat3Expression(gpu.Fn("transpose", a))
}

func Determinant(a ColumnMajor) gpu.Float { //glsl:determinant(mat3)float
	return gpu.NewFloatExpression(gpu.Fn("determinant", a))
}
func Inverse(a ColumnMajor) ColumnMajor { //glsl:inverse(mat3)mat3
	return gpu.NewMat3Expression(gpu.Fn("inverse", a))
}

func Mul(a, b ColumnMajor) ColumnMajor { //glsl:*(mat3,mat3)
	return gpu.NewMat3Expression(gpu.Op(a, "*", b))
}
