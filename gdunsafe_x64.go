//go:build amd64 || arm64 || wasipv1

package gdunsafe

type Pointer uintptr
type MutablePointer uintptr
type PackedArray[T Packable] [2]uint64

const (
	ShapeString      Shape = bytes8   // shape of a [String]
	ShapeObject      Shape = bytes8   // shape of an [Object]
	ShapeArray       Shape = bytes8   // shape of an [Array]
	ShapePackedArray Shape = bytes8x2 // shape of a [PackedArray[T]]
	ShapeDictionary  Shape = bytes8   // shape of a [Dictionary]
	ShapeStringName  Shape = bytes8   // shape of a [StringName]
	ShapeNodePath    Shape = bytes8   // shape of a [NodePath]
	ShapePointer     Shape = bytes8   // shape of a [Pointer]
)
