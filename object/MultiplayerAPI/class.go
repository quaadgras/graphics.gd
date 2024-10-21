package MultiplayerAPI

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Base class for high-level multiplayer API implementations. See also [MultiplayerPeer].
By default, [SceneTree] has a reference to an implementation of this class and uses it to provide multiplayer capabilities (i.e. RPCs) across the whole scene.
It is possible to override the MultiplayerAPI instance used by specific tree branches by calling the [method SceneTree.set_multiplayer] method, effectively allowing to run both client and server in the same scene.
It is also possible to extend or replace the default implementation via scripting or native extensions. See [MultiplayerAPIExtension] for details about extensions, [SceneMultiplayer] for the details about the default implementation.

*/
type Simple [1]classdb.MultiplayerAPI
func (self Simple) HasMultiplayerPeer() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasMultiplayerPeer())
}
func (self Simple) GetMultiplayerPeer() [1]classdb.MultiplayerPeer {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.MultiplayerPeer(Expert(self).GetMultiplayerPeer(gc))
}
func (self Simple) SetMultiplayerPeer(peer [1]classdb.MultiplayerPeer) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMultiplayerPeer(peer)
}
func (self Simple) GetUniqueId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetUniqueId()))
}
func (self Simple) IsServer() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsServer())
}
func (self Simple) GetRemoteSenderId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetRemoteSenderId()))
}
func (self Simple) Poll() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Poll())
}
func (self Simple) Rpc(peer int, obj gd.Object, method string, arguments gd.Array) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Rpc(gd.Int(peer), obj, gc.StringName(method), arguments))
}
func (self Simple) ObjectConfigurationAdd(obj gd.Object, configuration gd.Variant) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).ObjectConfigurationAdd(obj, configuration))
}
func (self Simple) ObjectConfigurationRemove(obj gd.Object, configuration gd.Variant) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).ObjectConfigurationRemove(obj, configuration))
}
func (self Simple) GetPeers() gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).GetPeers(gc))
}
func (self Simple) SetDefaultInterface(interface_name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDefaultInterface(gc, gc.StringName(interface_name))
}
func (self Simple) GetDefaultInterface() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetDefaultInterface(gc).String())
}
func (self Simple) CreateDefaultInterface() [1]classdb.MultiplayerAPI {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.MultiplayerAPI(Expert(self).CreateDefaultInterface(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.MultiplayerAPI
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns [code]true[/code] if there is a [member multiplayer_peer] set.
*/
//go:nosplit
func (self class) HasMultiplayerPeer() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerAPI.Bind_has_multiplayer_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetMultiplayerPeer(ctx gd.Lifetime) object.MultiplayerPeer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerAPI.Bind_get_multiplayer_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.MultiplayerPeer
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMultiplayerPeer(peer object.MultiplayerPeer)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(peer[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerAPI.Bind_set_multiplayer_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the unique peer ID of this MultiplayerAPI's [member multiplayer_peer].
*/
//go:nosplit
func (self class) GetUniqueId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerAPI.Bind_get_unique_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this MultiplayerAPI's [member multiplayer_peer] is valid and in server mode (listening for connections).
*/
//go:nosplit
func (self class) IsServer() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerAPI.Bind_is_server, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the sender's peer ID for the RPC currently being executed.
[b]Note:[/b] This method returns [code]0[/code] when called outside of an RPC. As such, the original peer ID may be lost when code execution is delayed (such as with GDScript's [code]await[/code] keyword).
*/
//go:nosplit
func (self class) GetRemoteSenderId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerAPI.Bind_get_remote_sender_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Method used for polling the MultiplayerAPI. You only need to worry about this if you set [member SceneTree.multiplayer_poll] to [code]false[/code]. By default, [SceneTree] will poll its MultiplayerAPI(s) for you.
[b]Note:[/b] This method results in RPCs being called, so they will be executed in the same context of this function (e.g. [code]_process[/code], [code]physics[/code], [Thread]).
*/
//go:nosplit
func (self class) Poll() int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerAPI.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sends an RPC to the target [param peer]. The given [param method] will be called on the remote [param object] with the provided [param arguments]. The RPC may also be called locally depending on the implementation and RPC configuration. See [method Node.rpc] and [method Node.rpc_config].
[b]Note:[/b] Prefer using [method Node.rpc], [method Node.rpc_id], or [code]my_method.rpc(peer, arg1, arg2, ...)[/code] (in GDScript), since they are faster. This method is mostly useful in conjunction with [MultiplayerAPIExtension] when augmenting or replacing the multiplayer capabilities.
*/
//go:nosplit
func (self class) Rpc(peer gd.Int, obj gd.Object, method gd.StringName, arguments gd.Array) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, peer)
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	callframe.Arg(frame, mmm.Get(method))
	callframe.Arg(frame, mmm.Get(arguments))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerAPI.Bind_rpc, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Notifies the MultiplayerAPI of a new [param configuration] for the given [param object]. This method is used internally by [SceneTree] to configure the root path for this MultiplayerAPI (passing [code]null[/code] and a valid [NodePath] as [param configuration]). This method can be further used by MultiplayerAPI implementations to provide additional features, refer to specific implementation (e.g. [SceneMultiplayer]) for details on how they use it.
[b]Note:[/b] This method is mostly relevant when extending or overriding the MultiplayerAPI behavior via [MultiplayerAPIExtension].
*/
//go:nosplit
func (self class) ObjectConfigurationAdd(obj gd.Object, configuration gd.Variant) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	callframe.Arg(frame, mmm.Get(configuration))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerAPI.Bind_object_configuration_add, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Notifies the MultiplayerAPI to remove a [param configuration] for the given [param object]. This method is used internally by [SceneTree] to configure the root path for this MultiplayerAPI (passing [code]null[/code] and an empty [NodePath] as [param configuration]). This method can be further used by MultiplayerAPI implementations to provide additional features, refer to specific implementation (e.g. [SceneMultiplayer]) for details on how they use it.
[b]Note:[/b] This method is mostly relevant when extending or overriding the MultiplayerAPI behavior via [MultiplayerAPIExtension].
*/
//go:nosplit
func (self class) ObjectConfigurationRemove(obj gd.Object, configuration gd.Variant) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	callframe.Arg(frame, mmm.Get(configuration))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerAPI.Bind_object_configuration_remove, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the peer IDs of all connected peers of this MultiplayerAPI's [member multiplayer_peer].
*/
//go:nosplit
func (self class) GetPeers(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerAPI.Bind_get_peers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the default MultiplayerAPI implementation class. This method can be used by modules and extensions to configure which implementation will be used by [SceneTree] when the engine starts.
*/
//go:nosplit
func (self class) SetDefaultInterface(ctx gd.Lifetime, interface_name gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(interface_name))
	var r_ret callframe.Nil
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.MultiplayerAPI.Bind_set_default_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the default MultiplayerAPI implementation class name. This is usually [code]"SceneMultiplayer"[/code] when [SceneMultiplayer] is available. See [method set_default_interface].
*/
//go:nosplit
func (self class) GetDefaultInterface(ctx gd.Lifetime) gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.MultiplayerAPI.Bind_get_default_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a new instance of the default MultiplayerAPI.
*/
//go:nosplit
func (self class) CreateDefaultInterface(ctx gd.Lifetime) object.MultiplayerAPI {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.MultiplayerAPI.Bind_create_default_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.MultiplayerAPI
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsMultiplayerAPI() Expert { return self[0].AsMultiplayerAPI() }


//go:nosplit
func (self Simple) AsMultiplayerAPI() Simple { return self[0].AsMultiplayerAPI() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("MultiplayerAPI", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type RPCMode = classdb.MultiplayerAPIRPCMode

const (
/*Used with [method Node.rpc_config] to disable a method or property for all RPC calls, making it unavailable. Default for all methods.*/
	RpcModeDisabled RPCMode = 0
/*Used with [method Node.rpc_config] to set a method to be callable remotely by any peer. Analogous to the [code]@rpc("any_peer")[/code] annotation. Calls are accepted from all remote peers, no matter if they are node's authority or not.*/
	RpcModeAnyPeer RPCMode = 1
/*Used with [method Node.rpc_config] to set a method to be callable remotely only by the current multiplayer authority (which is the server by default). Analogous to the [code]@rpc("authority")[/code] annotation. See [method Node.set_multiplayer_authority].*/
	RpcModeAuthority RPCMode = 2
)
