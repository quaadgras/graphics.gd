package Resource

import "unsafe"

// ID is used to access a low-level resource by its unique ID. [ID]s are opaque,
// which means they do not grant access to the resource by themselves. They are
// used by the low-level server classes, such as DisplayServer, RenderingServer,
// TextServer, etc.
//
// A low-level resource may correspond to a high-level Resource, such as Texture or Mesh.
//
// Note: [ID]s are only useful during the current session. It won't correspond to
// a similar resource if sent over a network, or loaded from a file at a later time.
type ID uint64

// Int64 returns the [ID] as an int64.
func (id ID) Int64() int64 { //gd:RID.get_id
	return *(*int64)(unsafe.Pointer(&id))
}

// IsValid returns true if the [ID] is not 0.
func (id ID) IsValid() bool { //gd:RID.is_valid
	return id != 0
}
