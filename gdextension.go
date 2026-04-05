//go:generate go run ./internal/tool/generate
//go:generate go run ./internal/tool/generate/v2
//go:generate go fmt ./...
package gdunsafe

import (
	"math/rand"
	"time"
	"unsafe"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/threadsafe"
)

type Int = int64
type Variant = gdextension.Variant
type CallError = gdextension.CallError
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
