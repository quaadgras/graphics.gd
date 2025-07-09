package gd

import (
	"reflect"

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

// TransferOwnershipToEngine releases the ownership of the given value. Use it on non-RefCounted
// Object return values passed back to the engine.
//
// used to fix cases of https://github.com/grow-graphics/gd/issues/147
func TransferOwnershipToEngine(val any) {
	rtype := reflect.TypeOf(val)
	switch rtype.Kind() {
	case reflect.Array:
		if rtype.Size() == 1 {
			class, ok := val.(IsClass)
			if ok {
				pointers.End(class.AsObject()[0])
			}
		}
	}
}
