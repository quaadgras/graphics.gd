package gd

import (
	"graphics.gd/internal/pointers"
)

// All pointer types from the engine need to be defined here.
// Each pointer of a given shape (number of uintptrs) must implement
// a Free method and have a unique [pointers.Type] array size, this
// is required for the pointers garbage collector to call the correct
// Free method.

type Object pointers.Trio[Object]
type RefCounted pointers.Trio[Object]

type enginePointer = EnginePointer
type packedPointers = PackedPointers
type VariantPointers = [3]uint64
