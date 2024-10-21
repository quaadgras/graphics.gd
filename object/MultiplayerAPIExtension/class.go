package MultiplayerAPIExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/MultiplayerAPI"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
		Poll() int64
		//Called when the [member MultiplayerAPI.multiplayer_peer] is set.
		SetMultiplayerPeer(multiplayer_peer object.MultiplayerPeer) 
		//Called when the [member MultiplayerAPI.multiplayer_peer] is retrieved.
		GetMultiplayerPeer() object.MultiplayerPeer
		//Callback for [method MultiplayerAPI.get_unique_id].
		GetUniqueId() gd.Int
		//Callback for [method MultiplayerAPI.get_peers].
		GetPeerIds() gd.PackedInt32Array
		//Callback for [method MultiplayerAPI.rpc].
		Rpc(peer gd.Int, obj gd.Object, method gd.StringName, args gd.Array) int64
		//Callback for [method MultiplayerAPI.get_remote_sender_id].
		GetRemoteSenderId() gd.Int
		//Callback for [method MultiplayerAPI.object_configuration_add].
		ObjectConfigurationAdd(obj gd.Object, configuration gd.Variant) int64
		//Callback for [method MultiplayerAPI.object_configuration_remove].
		ObjectConfigurationRemove(obj gd.Object, configuration gd.Variant) int64
	}

*/
type Simple [1]classdb.MultiplayerAPIExtension
func (Simple) _poll(impl func(ptr unsafe.Pointer) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _set_multiplayer_peer(impl func(ptr unsafe.Pointer, multiplayer_peer [1]classdb.MultiplayerPeer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var multiplayer_peer [1]classdb.MultiplayerPeer
		multiplayer_peer[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, multiplayer_peer)
		gc.End()
	}
}
func (Simple) _get_multiplayer_peer(impl func(ptr unsafe.Pointer) [1]classdb.MultiplayerPeer, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _get_unique_id(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _get_peer_ids(impl func(ptr unsafe.Pointer) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _rpc(impl func(ptr unsafe.Pointer, peer int, obj gd.Object, method string, args gd.Array) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var peer = gd.UnsafeGet[gd.Int](p_args,0)
		var obj gd.Object
		obj.SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,1)}))
		var method = mmm.Let[gd.StringName](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,2))
		var args = mmm.Let[gd.Array](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,3))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(peer), obj, method.String(), args)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _get_remote_sender_id(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _object_configuration_add(impl func(ptr unsafe.Pointer, obj gd.Object, configuration gd.Variant) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var obj gd.Object
		obj.SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var configuration = mmm.Let[gd.Variant](gc.Lifetime, gc.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, configuration)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _object_configuration_remove(impl func(ptr unsafe.Pointer, obj gd.Object, configuration gd.Variant) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var obj gd.Object
		obj.SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var configuration = mmm.Let[gd.Variant](gc.Lifetime, gc.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, configuration)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.MultiplayerAPIExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Callback for [method MultiplayerAPI.poll].
*/
func (class) _poll(impl func(ptr unsafe.Pointer) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Called when the [member MultiplayerAPI.multiplayer_peer] is set.
*/
func (class) _set_multiplayer_peer(impl func(ptr unsafe.Pointer, multiplayer_peer object.MultiplayerPeer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var multiplayer_peer object.MultiplayerPeer
		multiplayer_peer[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, multiplayer_peer)
		ctx.End()
	}
}

/*
Called when the [member MultiplayerAPI.multiplayer_peer] is retrieved.
*/
func (class) _get_multiplayer_peer(impl func(ptr unsafe.Pointer) object.MultiplayerPeer, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}

/*
Callback for [method MultiplayerAPI.get_unique_id].
*/
func (class) _get_unique_id(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Callback for [method MultiplayerAPI.get_peers].
*/
func (class) _get_peer_ids(impl func(ptr unsafe.Pointer) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Callback for [method MultiplayerAPI.rpc].
*/
func (class) _rpc(impl func(ptr unsafe.Pointer, peer gd.Int, obj gd.Object, method gd.StringName, args gd.Array) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var peer = gd.UnsafeGet[gd.Int](p_args,0)
		var obj gd.Object
		obj.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,1)}))
		var method = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,2))
		var args = mmm.Let[gd.Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,3))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, peer, obj, method, args)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Callback for [method MultiplayerAPI.get_remote_sender_id].
*/
func (class) _get_remote_sender_id(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Callback for [method MultiplayerAPI.object_configuration_add].
*/
func (class) _object_configuration_add(impl func(ptr unsafe.Pointer, obj gd.Object, configuration gd.Variant) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var obj gd.Object
		obj.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var configuration = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, configuration)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Callback for [method MultiplayerAPI.object_configuration_remove].
*/
func (class) _object_configuration_remove(impl func(ptr unsafe.Pointer, obj gd.Object, configuration gd.Variant) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var obj gd.Object
		obj.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var configuration = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, configuration)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}


//go:nosplit
func (self class) AsMultiplayerAPIExtension() Expert { return self[0].AsMultiplayerAPIExtension() }


//go:nosplit
func (self Simple) AsMultiplayerAPIExtension() Simple { return self[0].AsMultiplayerAPIExtension() }


//go:nosplit
func (self class) AsMultiplayerAPI() MultiplayerAPI.Expert { return self[0].AsMultiplayerAPI() }


//go:nosplit
func (self Simple) AsMultiplayerAPI() MultiplayerAPI.Simple { return self[0].AsMultiplayerAPI() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_poll": return reflect.ValueOf(self._poll);
	case "_set_multiplayer_peer": return reflect.ValueOf(self._set_multiplayer_peer);
	case "_get_multiplayer_peer": return reflect.ValueOf(self._get_multiplayer_peer);
	case "_get_unique_id": return reflect.ValueOf(self._get_unique_id);
	case "_get_peer_ids": return reflect.ValueOf(self._get_peer_ids);
	case "_rpc": return reflect.ValueOf(self._rpc);
	case "_get_remote_sender_id": return reflect.ValueOf(self._get_remote_sender_id);
	case "_object_configuration_add": return reflect.ValueOf(self._object_configuration_add);
	case "_object_configuration_remove": return reflect.ValueOf(self._object_configuration_remove);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_poll": return reflect.ValueOf(self._poll);
	case "_set_multiplayer_peer": return reflect.ValueOf(self._set_multiplayer_peer);
	case "_get_multiplayer_peer": return reflect.ValueOf(self._get_multiplayer_peer);
	case "_get_unique_id": return reflect.ValueOf(self._get_unique_id);
	case "_get_peer_ids": return reflect.ValueOf(self._get_peer_ids);
	case "_rpc": return reflect.ValueOf(self._rpc);
	case "_get_remote_sender_id": return reflect.ValueOf(self._get_remote_sender_id);
	case "_object_configuration_add": return reflect.ValueOf(self._object_configuration_add);
	case "_object_configuration_remove": return reflect.ValueOf(self._object_configuration_remove);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("MultiplayerAPIExtension", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
