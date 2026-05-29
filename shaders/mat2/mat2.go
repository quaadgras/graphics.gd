// Package mat2 provides GPU operations on 2x2 matrices.
package mat2

import (
	"graphics.gd/shaders/internal/gpu"
)

type ColumnMajor gpu.Mat2

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
