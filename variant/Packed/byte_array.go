// Package Packed provides functions for working with packed arrays.
package Packed

import (
	"grow.graphics/gd"
)

// ByteArray is an array specifically designed to hold bytes. Packs data tightly, so it saves memory for
// large array sizes.
//
// ByteArray also provides methods to encode/decode various types to/from bytes. The way values are
// encoded is an implementation detail and shouldn't be relied upon when interacting with external apps.
type ByteArray struct {
	Array[ByteArray, byte, gd.Int, gd.PackedByteArray, *gd.PackedByteArray]
}

func (array ByteArray) conv(c byte) gd.Int { return gd.Int(c) }
func (array ByteArray) wrap(c gd.Int) byte { return byte(c) }

func (array ByteArray) make(local []byte, proxy gd.PackedByteArray) ByteArray {
	array.local = local
	array.proxy = proxy
	return array
}

func (array ByteArray) less(a, b byte) bool {
	return a < b
}
