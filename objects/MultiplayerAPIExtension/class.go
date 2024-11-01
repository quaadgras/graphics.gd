package MultiplayerAPIExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/MultiplayerAPI"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
This class can be used to augment or replace the default [MultiplayerAPI] implementation via script or extensions.
The following example augment the default implementation ([SceneMultiplayer]) by logging every RPC being made, and every object being configured for replication.
[codeblocks]
[gdscript]
extends MultiplayerAPIExtension
class_name LogMultiplayer

# We want to augment the default SceneMultiplayer.
var base_multiplayer = SceneMultiplayer.new()

func _init():

	# Just passthrough base signals (copied to var to avoid cyclic reference)
	var cts = connected_to_server
	var cf = connection_failed
	var pc = peer_connected
	var pd = peer_disconnected
	base_multiplayer.connected_to_server.connect(func(): cts.emit())
	base_multiplayer.connection_failed.connect(func(): cf.emit())
	base_multiplayer.peer_connected.connect(func(id): pc.emit(id))
	base_multiplayer.peer_disconnected.connect(func(id): pd.emit(id))

func _poll():

	return base_multiplayer.poll()

# Log RPC being made and forward it to the default multiplayer.
func _rpc(peer: int, object: Object, method: StringName, args: Array) -> Error:

	print("Got RPC for %d: %s::%s(%s)" % [peer, object, method, args])
	return base_multiplayer.rpc(peer, object, method, args)

# Log configuration add. E.g. root path (nullptr, NodePath), replication (Node, Spawner|Synchronizer), custom.
func _object_configuration_add(object, config: Variant) -> Error:

	if config is MultiplayerSynchronizer:
	    print("Adding synchronization configuration for %s. Synchronizer: %s" % [object, config])
	elif config is MultiplayerSpawner:
	    print("Adding node %s to the spawn list. Spawner: %s" % [object, config])
	return base_multiplayer.object_configuration_add(object, config)

# Log configuration remove. E.g. root path (nullptr, NodePath), replication (Node, Spawner|Synchronizer), custom.
func _object_configuration_remove(object, config: Variant) -> Error:

	if config is MultiplayerSynchronizer:
	    print("Removing synchronization configuration for %s. Synchronizer: %s" % [object, config])
	elif config is MultiplayerSpawner:
	    print("Removing node %s from the spawn list. Spawner: %s" % [object, config])
	return base_multiplayer.object_configuration_remove(object, config)

# These can be optional, but in our case we want to augment SceneMultiplayer, so forward everything.
func _set_multiplayer_peer(p_peer: MultiplayerPeer):

	base_multiplayer.multiplayer_peer = p_peer

func _get_multiplayer_peer() -> MultiplayerPeer:

	return base_multiplayer.multiplayer_peer

func _get_unique_id() -> int:

	return base_multiplayer.get_unique_id()

func _get_peer_ids() -> PackedInt32Array:

	return base_multiplayer.get_peers()

[/gdscript]
[/codeblocks]
Then in your main scene or in an autoload call [method SceneTree.set_multiplayer] to start using your custom [MultiplayerAPI]:
[codeblocks]
[gdscript]
# autoload.gd
func _enter_tree():

	# Sets our custom multiplayer as the main one in SceneTree.

get_tree().set_multiplayer(LogMultiplayer.new())
[/gdscript]
[/codeblocks]
Native extensions can alternatively use the [method MultiplayerAPI.set_default_interface] method during initialization to configure themselves as the default implementation.

	// MultiplayerAPIExtension methods that can be overridden by a [Class] that extends it.
	type MultiplayerAPIExtension interface {
		//Callback for [method MultiplayerAPI.poll].
		Poll() gd.Error
		//Called when the [member MultiplayerAPI.multiplayer_peer] is set.
		SetMultiplayerPeer(multiplayer_peer objects.MultiplayerPeer)
		//Called when the [member MultiplayerAPI.multiplayer_peer] is retrieved.
		GetMultiplayerPeer() objects.MultiplayerPeer
		//Callback for [method MultiplayerAPI.get_unique_id].
		GetUniqueId() int
		//Callback for [method MultiplayerAPI.get_peers].
		GetPeerIds() []int32
		//Callback for [method MultiplayerAPI.rpc].
		Rpc(peer int, obj gd.Object, method string, args gd.Array) gd.Error
		//Callback for [method MultiplayerAPI.get_remote_sender_id].
		GetRemoteSenderId() int
		//Callback for [method MultiplayerAPI.object_configuration_add].
		ObjectConfigurationAdd(obj gd.Object, configuration gd.Variant) gd.Error
		//Callback for [method MultiplayerAPI.object_configuration_remove].
		ObjectConfigurationRemove(obj gd.Object, configuration gd.Variant) gd.Error
	}
*/
type Instance [1]classdb.MultiplayerAPIExtension

/*
Callback for [method MultiplayerAPI.poll].
*/
func (Instance) _poll(impl func(ptr unsafe.Pointer) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [member MultiplayerAPI.multiplayer_peer] is set.
*/
func (Instance) _set_multiplayer_peer(impl func(ptr unsafe.Pointer, multiplayer_peer objects.MultiplayerPeer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var multiplayer_peer = objects.MultiplayerPeer{pointers.New[classdb.MultiplayerPeer]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(multiplayer_peer[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, multiplayer_peer)
	}
}

/*
Called when the [member MultiplayerAPI.multiplayer_peer] is retrieved.
*/
func (Instance) _get_multiplayer_peer(impl func(ptr unsafe.Pointer) objects.MultiplayerPeer) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Callback for [method MultiplayerAPI.get_unique_id].
*/
func (Instance) _get_unique_id(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Callback for [method MultiplayerAPI.get_peers].
*/
func (Instance) _get_peer_ids(impl func(ptr unsafe.Pointer) []int32) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedInt32Slice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Callback for [method MultiplayerAPI.rpc].
*/
func (Instance) _rpc(impl func(ptr unsafe.Pointer, peer int, obj gd.Object, method string, args gd.Array) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var peer = gd.UnsafeGet[gd.Int](p_args, 0)
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})
		defer pointers.End(obj)
		var method = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 2))
		defer pointers.End(method)
		var args = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 3))
		defer pointers.End(args)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(peer), obj, method.String(), args)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Callback for [method MultiplayerAPI.get_remote_sender_id].
*/
func (Instance) _get_remote_sender_id(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Callback for [method MultiplayerAPI.object_configuration_add].
*/
func (Instance) _object_configuration_add(impl func(ptr unsafe.Pointer, obj gd.Object, configuration gd.Variant) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		var configuration = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 1))
		defer pointers.End(configuration)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, configuration)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Callback for [method MultiplayerAPI.object_configuration_remove].
*/
func (Instance) _object_configuration_remove(impl func(ptr unsafe.Pointer, obj gd.Object, configuration gd.Variant) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		var configuration = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 1))
		defer pointers.End(configuration)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, configuration)
		gd.UnsafeSet(p_back, ret)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.MultiplayerAPIExtension

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MultiplayerAPIExtension"))
	return Instance{classdb.MultiplayerAPIExtension(object)}
}

/*
Callback for [method MultiplayerAPI.poll].
*/
func (class) _poll(impl func(ptr unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [member MultiplayerAPI.multiplayer_peer] is set.
*/
func (class) _set_multiplayer_peer(impl func(ptr unsafe.Pointer, multiplayer_peer objects.MultiplayerPeer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var multiplayer_peer = objects.MultiplayerPeer{pointers.New[classdb.MultiplayerPeer]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(multiplayer_peer[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, multiplayer_peer)
	}
}

/*
Called when the [member MultiplayerAPI.multiplayer_peer] is retrieved.
*/
func (class) _get_multiplayer_peer(impl func(ptr unsafe.Pointer) objects.MultiplayerPeer) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Callback for [method MultiplayerAPI.get_unique_id].
*/
func (class) _get_unique_id(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Callback for [method MultiplayerAPI.get_peers].
*/
func (class) _get_peer_ids(impl func(ptr unsafe.Pointer) gd.PackedInt32Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Callback for [method MultiplayerAPI.rpc].
*/
func (class) _rpc(impl func(ptr unsafe.Pointer, peer gd.Int, obj gd.Object, method gd.StringName, args gd.Array) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var peer = gd.UnsafeGet[gd.Int](p_args, 0)
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})
		defer pointers.End(obj)
		var method = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 2))
		var args = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 3))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, peer, obj, method, args)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Callback for [method MultiplayerAPI.get_remote_sender_id].
*/
func (class) _get_remote_sender_id(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Callback for [method MultiplayerAPI.object_configuration_add].
*/
func (class) _object_configuration_add(impl func(ptr unsafe.Pointer, obj gd.Object, configuration gd.Variant) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		var configuration = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, configuration)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Callback for [method MultiplayerAPI.object_configuration_remove].
*/
func (class) _object_configuration_remove(impl func(ptr unsafe.Pointer, obj gd.Object, configuration gd.Variant) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		var configuration = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, configuration)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsMultiplayerAPIExtension() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMultiplayerAPIExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsMultiplayerAPI() MultiplayerAPI.Advanced {
	return *((*MultiplayerAPI.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsMultiplayerAPI() MultiplayerAPI.Instance {
	return *((*MultiplayerAPI.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_poll":
		return reflect.ValueOf(self._poll)
	case "_set_multiplayer_peer":
		return reflect.ValueOf(self._set_multiplayer_peer)
	case "_get_multiplayer_peer":
		return reflect.ValueOf(self._get_multiplayer_peer)
	case "_get_unique_id":
		return reflect.ValueOf(self._get_unique_id)
	case "_get_peer_ids":
		return reflect.ValueOf(self._get_peer_ids)
	case "_rpc":
		return reflect.ValueOf(self._rpc)
	case "_get_remote_sender_id":
		return reflect.ValueOf(self._get_remote_sender_id)
	case "_object_configuration_add":
		return reflect.ValueOf(self._object_configuration_add)
	case "_object_configuration_remove":
		return reflect.ValueOf(self._object_configuration_remove)
	default:
		return gd.VirtualByName(self.AsMultiplayerAPI(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_poll":
		return reflect.ValueOf(self._poll)
	case "_set_multiplayer_peer":
		return reflect.ValueOf(self._set_multiplayer_peer)
	case "_get_multiplayer_peer":
		return reflect.ValueOf(self._get_multiplayer_peer)
	case "_get_unique_id":
		return reflect.ValueOf(self._get_unique_id)
	case "_get_peer_ids":
		return reflect.ValueOf(self._get_peer_ids)
	case "_rpc":
		return reflect.ValueOf(self._rpc)
	case "_get_remote_sender_id":
		return reflect.ValueOf(self._get_remote_sender_id)
	case "_object_configuration_add":
		return reflect.ValueOf(self._object_configuration_add)
	case "_object_configuration_remove":
		return reflect.ValueOf(self._object_configuration_remove)
	default:
		return gd.VirtualByName(self.AsMultiplayerAPI(), name)
	}
}
func init() {
	classdb.Register("MultiplayerAPIExtension", func(ptr gd.Object) any { return classdb.MultiplayerAPIExtension(ptr) })
}
