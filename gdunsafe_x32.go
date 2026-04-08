//go:build wasm32 || js

package gdunsafe

type PackedArray[T Packable] [2]uint32

const (
	ShapeString      Shape = bytes4   // shape of a [String]
	ShapeObject      Shape = bytes4   // shape of an [Object]
	ShapeArray       Shape = bytes4   // shape of an [Array]
	ShapePackedArray Shape = bytes4x2 // shape of a [PackedArray[T]]
	ShapeDictionary  Shape = bytes4   // shape of a [Dictionary]
	ShapeStringName  Shape = bytes4   // shape of a [StringName]
	ShapeNodePath    Shape = bytes4   // shape of a [NodePath]
	ShapePointer     Shape = bytes4   // shape of a [Pointer]
)
