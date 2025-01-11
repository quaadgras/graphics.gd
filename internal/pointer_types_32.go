//go:build i386 || wasm

package gd

import (
	"graphics.gd/internal/pointers"
)

type gdptr uint32

type Variant pointers.Hexa[Variant]
type Signal pointers.Quad[Signal]
type Callable pointers.Quad[Callable]

type VariantPointers = [6]uintptr
type SignalPointers = [4]uintptr
type CallablePointers = [4]uintptr
