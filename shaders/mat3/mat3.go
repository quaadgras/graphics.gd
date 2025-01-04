package mat3

import (
	"graphics.gd/shaders/internal/dsl"
	"graphics.gd/shaders/vec1"
)

type expression = dsl.Expression

type ColumnMajor struct {
	expression

	Columns [3][3]vec1.X
}

func ComponentMul(a, b ColumnMajor) ColumnMajor {
	return ColumnMajor{expression: dsl.Fn("matrixCompMult", a, b)}
}
func OuterProduct(a, b vec1.X) ColumnMajor {
	return ColumnMajor{expression: dsl.Fn("outerProduct", a, b)}
}
func Transpose(a ColumnMajor) ColumnMajor {
	return ColumnMajor{expression: dsl.Fn("transpose", a)}
}

func Determinant(a ColumnMajor) vec1.X {
	return vec1.X{Expression: dsl.Fn("determinant", a)}
}
func Inverse(a ColumnMajor) ColumnMajor {
	return ColumnMajor{expression: dsl.Fn("inverse", a)}
}

func Mul(a, b ColumnMajor) ColumnMajor {
	return ColumnMajor{expression: dsl.Op(a, "*", b)}
}
