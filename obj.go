//go:build !generate

package gd

import (
	"unsafe"
)

type cache struct {
	args []unsafe.Pointer
	byte []byte
}
