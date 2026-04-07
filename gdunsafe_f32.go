//go:build !precision_double

package gdunsafe

// Variant is the raw representation for a variant value in the engine.
// It should be destroyed with [Variant.Free] when no longer in use.
type Variant [3]uint64

const (
	bytes8x3  Shape = 9 + iota // shape for three 8-byte values
	bytes4x6                   // shape for six 4-byte values
	bytes4x9                   // shape for nine 4-byte values
	bytes4x12                  // shape for twelve 4-byte values
	bytes4x16                  // shape for sixteen 4-byte values
)

const (
	ShapeVariant     Shape = bytes8x3  // shape of a [Variant]
	ShapeVector2     Shape = bytes4x2  // shape of a [Vector2.XY]
	ShapeVector3     Shape = bytes4x3  // shape of a [Vector3.XYZ]
	ShapeVector4     Shape = bytes4x4  // shape of a [Vector4.XYZW]
	ShapeRect2       Shape = bytes4x4  // shape of a [Rect2.PositionSize]
	ShapeTransform2D Shape = bytes4x6  // shape of a [Transform2D.OriginXY]
	ShapeTransform3D Shape = bytes4x12 // shape of a [Transform3D.BasisOrigin]
	ShapePlane       Shape = bytes4x4  // shape of a [Plane.NormalD]
	ShapeQuaternion  Shape = bytes4x4  // shape of a [Quaternion.IJKL]
	ShapeAABB        Shape = bytes4x6  // shape of a [AABB.PositionSize]
	ShapeBasis       Shape = bytes4x9  // shape of a [Basis.XYZ]
	ShapeProjection  Shape = bytes4x16 // shape of a [Projection.XYZW]
)

func (shape Shape) SizeResult() (size int) {
	switch shape & 0xF {
	case empty:
		return 0
	case bytes1:
		return 1
	case bytes2:
		return 2
	case bytes4:
		return 4
	case bytes8:
		return 8
	case bytes4x2:
		return 4 * 2
	case bytes4x3:
		return 4 * 3
	case bytes8x2:
		return 8 * 2
	case bytes4x4:
		return 4 * 4
	case bytes8x3:
		return 8 * 3
	case bytes4x6:
		return 4 * 6
	case bytes4x9:
		return 4 * 9
	case bytes4x12:
		return 4 * 12
	case bytes4x16:
		return 4 * 16
	default:
		panic("Shape.SizeResult: invalid shape")
	}
}

// Alignment returns the memory alignment for a [Shape].
func (shape Shape) Alignment() int {
	switch shape {
	case empty:
		return 0
	case bytes1:
		return 1
	case bytes2:
		return 2
	case bytes4, bytes4x2, bytes4x3, bytes4x4, bytes4x6, bytes4x9, bytes4x12, bytes4x16:
		return 4
	case bytes8, bytes8x2, bytes8x3:
		return 8
	default:
		panic("Shape.Alignment: invalid shape")
	}
}
