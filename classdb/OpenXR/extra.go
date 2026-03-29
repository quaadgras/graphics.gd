package OpenXR

import (
	"structs"
	"unsafe"
)

// Extension is the common header for all chainable OpenXR structs.
// Every struct in an OpenXR next-pointer chain starts with these two fields.
type Extension struct {
	_ structs.HostLayout

	Type int32
	Next unsafe.Pointer
}

// Vector3f represents a 3D vector in OpenXR.
type Vector3f struct {
	_ structs.HostLayout

	X, Y, Z float32
}

// Quaternionf represents a quaternion in OpenXR.
type Quaternionf struct {
	_ structs.HostLayout

	X, Y, Z, W float32
}

// Posef represents a position and orientation in OpenXR.
type Posef struct {
	_ structs.HostLayout

	Orientation Quaternionf
	Position    Vector3f
}

// EventDataBuffer is the buffer passed to xrPollEvent.
type EventDataBuffer struct {
	_ structs.HostLayout

	Type    int32
	Next    unsafe.Pointer
	Varying [4000]byte
}

// CompositionLayer is the base header for all composition layer types.
type CompositionLayer struct {
	_ structs.HostLayout

	Type       int32
	Next       unsafe.Pointer
	LayerFlags uint64
	Space      uint64
}
