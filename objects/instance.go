package objects

import (
	gd "grow.graphics/gd/internal"
)

// ID that uniquely identifies an object.
type ID gd.ObjectID

// IsValid returns true if the Object that corresponds to id is a valid object
// (e.g. has not been deleted from memory). All Objects have a unique instance ID.
func (id ID) IsValid() bool { //gd:is_instance_id_valid is_instance_valid
	return bool(gd.IsInstanceIdValid(gd.Int(id)))
}

// Get returns the Object that corresponds to id. All Objects have a unique instance ID.
func Get(id ID) gd.Object { //gd:instance_from_id
	return gd.InstanceFromId(gd.Int(id))
}

// New returns a new empty object instance.
func New() gd.Object {
	return gd.Global.ClassDB.ConstructObject(gd.NewStringName("Object"))
}
