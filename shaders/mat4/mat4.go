// Package mat4 provides GPU operations on 4x4 matrices.
package mat4

import (
	"graphics.gd/shaders/internal/gpu"
)

type ColumnMajor gpu.Mat4

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
