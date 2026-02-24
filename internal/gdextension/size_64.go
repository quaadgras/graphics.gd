//go:build amd64 || arm64 || wasm

package gdextension

import (
	"reflect"
	"unsafe"
)

func SizeOf[T AnyVariant]() Shape {
	type (
		bytes1 = uint8
		bytes2 = uint16
		bytes4 = uint32
		bytes8 = uint64
	)
	switch unsafe.Sizeof([1]T{}[0])<<8 + unsafe.Alignof([1]T{}[0]) {
	case 0:
		return ShapeEmpty
	case unsafe.Sizeof([1]bytes1{})<<8 + unsafe.Alignof([1]bytes1{}):
		return ShapeBytes1
	case unsafe.Sizeof([1]bytes2{})<<8 + unsafe.Alignof([1]bytes2{}):
		return ShapeBytes2
	case unsafe.Sizeof([1]bytes4{})<<8 + unsafe.Alignof([1]bytes4{}):
		return ShapeBytes4
	case unsafe.Sizeof([1]bytes8{})<<8 + unsafe.Alignof([1]bytes8{}):
		return ShapeBytes8
	case unsafe.Sizeof([2]bytes4{})<<8 + unsafe.Alignof([2]bytes4{}):
		return ShapeBytes4x2
	case unsafe.Sizeof([3]bytes4{})<<8 + unsafe.Alignof([3]bytes4{}):
		return ShapeBytes4x3
	case unsafe.Sizeof([2]bytes8{})<<8 + unsafe.Alignof([2]bytes8{}):
		return ShapeBytes8x2
	case unsafe.Sizeof([4]bytes4{})<<8 + unsafe.Alignof([4]bytes4{}):
		return ShapeBytes4x4
	case unsafe.Sizeof([3]bytes8{})<<8 + unsafe.Alignof([3]bytes8{}):
		return ShapeBytes8x3
	case unsafe.Sizeof([6]bytes4{})<<8 + unsafe.Alignof([6]bytes4{}):
		return ShapeBytes4x6
	case unsafe.Sizeof([9]bytes4{})<<8 + unsafe.Alignof([9]bytes4{}):
		return ShapeBytes4x9
	case unsafe.Sizeof([12]bytes4{})<<8 + unsafe.Alignof([12]bytes4{}):
		return ShapeBytes4x12
	case unsafe.Sizeof([16]bytes4{})<<8 + unsafe.Alignof([16]bytes4{}):
		return ShapeBytes4x16
	default:
		panic("SizeOf: unsupported type " + reflect.TypeFor[T]().String())
	}
}
