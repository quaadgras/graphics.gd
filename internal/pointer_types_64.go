//go:build amd64 || arm64

package gd

import (
	"graphics.gd/internal/pointers"
)

type gdptr uint64

type Variant pointers.Trio[Variant]
type Signal pointers.Pair[Signal]
type Callable pointers.Pair[Callable]

type VariantPointers = [3]uintptr
type SignalPointers = [2]uintptr
type CallablePointers = [2]uintptr
