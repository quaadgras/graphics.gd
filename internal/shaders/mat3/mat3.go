// Package mat3 provides GPU operations on 3x3 matrices.
package mat3

import (
	"graphics.gd/shaders/internal/gpu"
)

type ColumnMajor gpu.Mat3

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
