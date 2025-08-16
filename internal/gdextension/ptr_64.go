//go:build !js

package gdextension

import "unsafe"

type Pointer = uintptr
type Returns[T any] unsafe.Pointer
type Accepts[T any] unsafe.Pointer
