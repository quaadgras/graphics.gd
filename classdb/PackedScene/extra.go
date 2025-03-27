package PackedScene

import (
	"reflect"
	"unsafe"

	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Resource"
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdclass"
	"graphics.gd/variant/Object"
)

// Instantiate an [Instance] of a PackedScene as an object of type T.
// panics if the object is not of the expected type.
func Instantiate[T Node.Any](packed_scene Instance) T {
	return Object.To[T](Node.Instance(packed_scene.Instantiate()))
}

// Is wraps a PackedScene with a type.
type Is[T Node.Any] [1]gdclass.PackedScene

func (self Is[T]) AsPackedScene() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self Is[T]) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self Is[T]) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self Is[T]) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}

func (self Is[T]) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *Is[T]) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Packs the [param path] node, and all owned sub-nodes, into this [PackedScene]. Any existing data will be cleared. See [member Node.owner].
*/
func (self Is[T]) Pack(path T) error { //gd:PackedScene.pack
	return error(gd.ToError(class(self).Pack(path.AsNode())))
}

/*
Instantiates the scene's node hierarchy. Triggers child scene instantiation(s). Triggers a [constant Node.NOTIFICATION_SCENE_INSTANTIATED] notification on the root node.
*/
func (self Is[T]) Instantiate() T { //gd:PackedScene.instantiate
	return Object.To[T](Node.Instance(class(self).Instantiate(0)))
}

/*
Returns [code]true[/code] if the scene file has nodes.
*/
func (self Is[T]) CanInstantiate() bool { //gd:PackedScene.can_instantiate
	return bool(class(self).CanInstantiate())
}

/*
Returns the [SceneState] representing the scene file contents.
*/
func (self Is[T]) GetState() [1]gdclass.SceneState { //gd:PackedScene.get_state
	return [1]gdclass.SceneState(class(self).GetState())
}
