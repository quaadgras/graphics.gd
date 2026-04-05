//go:generate go run ./internal/tool/generate
//go:generate go run ./internal/tool/generate/v2
//go:generate go fmt ./...
package gdunsafe

import (
	"math/rand"
	"structs"
	"time"
	"unsafe"

	"graphics.gd/internal/threadsafe"
)

type Int = int64
type Variant [3]uint64
type CallError struct {
	_ structs.HostLayout

	Type     CallErrorType
	Argument int32
	Expected int32
}
type CallErrorType uint32
type ObjectID uint64

type PointerTo[T any] Pointer

type VariadicVariants struct {
	First PointerTo[PointerTo[Variant]]
	Count int
}

var callables threadsafe.Handles[CallableImplementation, CallableID]

type CallableImplementation interface {

	// Call is called whenever the [Callable] is called. It's given
	// a variadic list of arguments and the implementation returns a [Variant]. If the
	// arguments aren't compatible, return a non-zero error.
	Call(VariadicVariants) (Variant, CallError)

	// IsValid is called to verify that the [Callable] is valid. It should return true
	// if the callable is in a valid state (and callable), otherwise false.
	IsValid() bool

	// Hash is called to hash the [Callable]. Identical underlying implementations of a
	// callable should always return the same value.
	Hash() uint32

	// UnsafeString is called when the [Callable] is being converted by the engine into a string.
	// It should return a useful string representation of the callable, or an error.
	UnsafeString() String

	// NumIn is called to determine how many arguments the [Callable] expects to
	// receive. Return -1, if the [Callable] is able to accept an unknown number of arguments.
	NumIn() int

	// Compare is called to compare two different [CustomCallable] values. It should
	// return less than zero, if a < b, zero if a = b and more than zero if a > b.
	Compare(CallableImplementation) int
}

func (ptr Pointer) SetInt32(v int32) {
	ptr.SetUint32(*(*uint32)(unsafe.Pointer(&v)))
}

// just a placeholder for functions that don't need to be implemented
// as they are already available in the Go standard library.

func randomize() { //gd:randomize
	rand.Seed(time.Now().UnixNano())
}

func seed(s int) { //gd:seed
	rand.Seed(int64(s))
}

func rand_from_seed(seed int) *rand.Rand { //gd:rand_from_seed
	return rand.New(rand.NewSource(int64(seed)))
}

func weakref(v any) any { return v } //gd:weakref

// Shape is used to correctly transfer data for unsafe calls into the engine.
type Shape uint64

const (
	ShapeEmpty Shape = iota

	ShapeBytes1
	ShapeBytes2
	ShapeBytes4
	ShapeBytes8
	ShapeBytes4x2
	ShapeBytes4x3
	ShapeBytes8x2
	ShapeBytes4x4
	ShapeBytes8x3
	ShapeBytes4x6
	ShapeBytes4x9
	ShapeBytes4x12
	ShapeBytes4x16
)

const (
	SizeVariant     Shape = ShapeBytes8x3
	SizeBool        Shape = ShapeBytes1
	SizeInt         Shape = ShapeBytes8
	SizeFloat       Shape = ShapeBytes8
	SizeVector2     Shape = ShapeBytes4x2
	SizeVector3     Shape = ShapeBytes4x3
	SizeVector4     Shape = ShapeBytes4x4
	SizeColor       Shape = ShapeBytes4x4
	SizeRect2       Shape = ShapeBytes4x4
	SizeRect2i      Shape = ShapeBytes4x4
	SizeVector2i    Shape = ShapeBytes4x2
	SizeVector3i    Shape = ShapeBytes4x3
	SizeVector4i    Shape = ShapeBytes4x4
	SizeTransform2D Shape = ShapeBytes4x6
	SizeTransform3D Shape = ShapeBytes4x12
	SizePlane       Shape = ShapeBytes4x4
	SizeQuaternion  Shape = ShapeBytes4x4
	SizeAABB        Shape = ShapeBytes4x6
	SizeBasis       Shape = ShapeBytes4x9
	SizeProjection  Shape = ShapeBytes4x16
	SizeRID         Shape = ShapeBytes8
	SizeCallable    Shape = ShapeBytes8x2
	SizeSignal      Shape = ShapeBytes8x2

	SizeCallError Shape = ShapeBytes4x3
)

func (shape Shape) SizeResult() (size int) {
	switch shape & 0xF {
	case ShapeEmpty:
		return 0
	case ShapeBytes1:
		return 1
	case ShapeBytes2:
		return 2
	case ShapeBytes4:
		return 4
	case ShapeBytes8:
		return 8
	case ShapeBytes4x2:
		return 4 * 2
	case ShapeBytes4x3:
		return 4 * 3
	case ShapeBytes8x2:
		return 8 * 2
	case ShapeBytes4x4:
		return 4 * 4
	case ShapeBytes8x3:
		return 8 * 3
	case ShapeBytes4x6:
		return 4 * 6
	case ShapeBytes4x9:
		return 4 * 9
	case ShapeBytes4x12:
		return 4 * 12
	case ShapeBytes4x16:
		return 4 * 16
	default:
		panic("Shape.SizeResult: invalid shape")
	}
}

const (
	CallOK               CallErrorType = iota
	CallInvalidMethod    CallErrorType = 1
	CallInvalidArguments CallErrorType = 2
	CallTooManyArguments CallErrorType = 3
	CallTooFewArguments  CallErrorType = 4
	CallInstanceIsNull   CallErrorType = 5
	CallMethodNotConst   CallErrorType = 6
)

func (err CallError) Err() error {
	if err.Type == CallOK {
		return nil
	}
	return err
}

func (err CallError) Error() string {
	switch err.Type {
	case CallInvalidMethod:
		return "Call Invalid Method"
	case CallInvalidArguments:
		return "Call Invalid Arguments"
	case CallTooManyArguments:
		return "Call Too Many Arguments"
	case CallTooFewArguments:
		return "Call Too Few Arguments"
	case CallInstanceIsNull:
		return "Call Instance Is Null"
	case CallMethodNotConst:
		return "Call Method Not Const"
	default:
		return "Unknown Call Error"
	}
}

// ALIGN_UP aligns a value to the next multiple of align.
func alignUp(value, align uint32) uint32 {
	return (value + (align - 1)) & ^(align - 1)
}

func (shape Shape) Alignment() int {
	switch shape {
	case ShapeEmpty:
		return 0
	case ShapeBytes1:
		return 1
	case ShapeBytes2:
		return 2
	case ShapeBytes4, ShapeBytes4x2, ShapeBytes4x3, ShapeBytes4x4, ShapeBytes4x6, ShapeBytes4x9, ShapeBytes4x12, ShapeBytes4x16:
		return 4
	case ShapeBytes8, ShapeBytes8x2, ShapeBytes8x3:
		return 8
	default:
		panic("Shape.Alignment: invalid shape")
	}
}

func (shape Shape) SizeArguments() (size int) {
	for i := 1; i < 16; i++ {
		var current = (shape >> (i * 4)) & 0xF
		switch current {
		case ShapeEmpty:
			return size
		default:
			size += current.SizeResult()
			size = int(alignUp(uint32(size), uint32(current.Alignment())))
		}
	}
	return
}
