//go:build !debug

// Package mmm provides a way to manually manage memory and resource lifetimes with protections against unsafe double-free and use-after-free errors.
package mmm

import (
	"unsafe"
)

type freeable struct {
	prv *freeable
	nxt *freeable
	api unsafe.Pointer // nil if root
	end func(genericPointer)

	// rev highest bit is set if pinned
	// rev second highest bit is set if reference counted.
	rev revision
}

func (f *freeable) record(kind string) {}

func crash(ptr *freeable, name string) {
	panic(name)
}
