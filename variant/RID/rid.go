package RID

import (
	gd "grow.graphics/gd/internal"
	"grow.graphics/gd/objects/Resource"
)

// New allocates a unique ID which can be used by the implementation to construct an RID.
// This is used mainly from native extensions to implement servers.
func New() Resource.ID { //gd:rid_allocate_id
	return gd.RidFromInt64(gd.RidAllocateId())
}

// Int64 returns a resource ID from the given int64.
func Int64(id int64) Resource.ID { //gd:rid_from_int64
	return gd.RidFromInt64(id)
}
