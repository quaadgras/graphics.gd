// Package mat2 provides GPU operations on 2x2 matrices.
package mat2

import (
	"graphics.gd/shaders/internal/gpu"
	"graphics.gd/shaders/vec2"
)

type ColumnMajor gpu.Mat2

// New returns a matrix with the given columns.
func New(x, y vec2.XY) ColumnMajor {
	return ColumnMajor{Columns: [2][2]gpu.Float{{x.X, x.Y}, {y.X, y.Y}}}
}

// Transform returns the vector v transformed by the matrix m.
func Transform(m ColumnMajor, v vec2.XY) vec2.XY { //glsl:*(mat2,vec2)
	return gpu.NewVec2Expression(gpu.Op(m, "*", v))
}

// Column returns column i of the matrix m.
func Column(m ColumnMajor, i int) vec2.XY { //glsl:mat2[int]vec2
	return gpu.NewVec2Expression(gpu.New(gpu.Index{Value: m, Index: i}))
}

func ComponentMul(a, b ColumnMajor) ColumnMajor { //glsl:matrixCompMult(mat2,mat2)mat2
	return gpu.NewMat2Expression(gpu.Fn("matrixCompMult", a, b))
}
func OuterProduct(a, b gpu.Vec2) ColumnMajor { //glsl:outerProduct(vec2,vec2)mat2
	return gpu.NewMat2Expression(gpu.Fn("outerProduct", a, b))
}
func Transpose(a ColumnMajor) ColumnMajor { //glsl:transpose(mat2)mat2
	return gpu.NewMat2Expression(gpu.Fn("transpose", a))
}

func Determinant(a ColumnMajor) gpu.Float { //glsl:determinant(mat2)float
	return gpu.NewFloatExpression(gpu.Fn("determinant", a))
}
func Inverse(a ColumnMajor) ColumnMajor { //glsl:inverse(mat2)mat2
	return gpu.NewMat2Expression(gpu.Fn("inverse", a))
}

func Mul(a, b ColumnMajor) ColumnMajor { //glsl:*(mat2,mat2)
	return gpu.NewMat2Expression(gpu.Op(a, "*", b))
}
