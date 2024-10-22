package SceneMultiplayer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/MultiplayerAPI"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class is the default implementation of [MultiplayerAPI], used to provide multiplayer functionalities in Godot Engine.
This implementation supports RPCs via [method Node.rpc] and [method Node.rpc_id] and requires [method MultiplayerAPI.rpc] to be passed a [Node] (it will fail for other object types).
This implementation additionally provide [SceneTree] replication via the [MultiplayerSpawner] and [MultiplayerSynchronizer] nodes, and the [SceneReplicationConfig] resource.
[b]Note:[/b] The high-level multiplayer API protocol is an implementation detail and isn't meant to be used by non-Godot servers. It may change without notice.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.

*/
type Go [1]classdb.SceneMultiplayer

/*
Clears the current SceneMultiplayer network state (you shouldn't call this unless you know what you are doing).
*/
func (self Go) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Clear()
}

/*
Disconnects the peer identified by [param id], removing it from the list of connected peers, and closing the underlying connection with it.
*/
func (self Go) DisconnectPeer(id int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DisconnectPeer(gd.Int(id))
}

/*
Returns the IDs of the peers currently trying to authenticate with this [MultiplayerAPI].
*/
func (self Go) GetAuthenticatingPeers() []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).GetAuthenticatingPeers(gc).AsSlice())
}

/*
Sends the specified [param data] to the remote peer identified by [param id] as part of an authentication message. This can be used to authenticate peers, and control when [signal MultiplayerAPI.peer_connected] is emitted (and the remote peer accepted as one of the connected peers).
*/
func (self Go) SendAuth(id int, data []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).SendAuth(gd.Int(id), gc.PackedByteSlice(data)))
}

/*
Mark the authentication step as completed for the remote peer identified by [param id]. The [signal MultiplayerAPI.peer_connected] signal will be emitted for this peer once the remote side also completes the authentication. No further authentication messages are expected to be received from this peer.
If a peer disconnects before completing authentication, either due to a network issue, the [member auth_timeout] expiring, or manually calling [method disconnect_peer], the [signal peer_authentication_failed] signal will be emitted instead of [signal MultiplayerAPI.peer_disconnected].
*/
func (self Go) CompleteAuth(id int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).CompleteAuth(gd.Int(id)))
}

/*
Sends the given raw [param bytes] to a specific peer identified by [param id] (see [method MultiplayerPeer.set_target_peer]). Default ID is [code]0[/code], i.e. broadcast to all peers.
*/
func (self Go) SendBytes(bytes []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).SendBytes(gc.PackedByteSlice(bytes), gd.Int(0), 2, gd.Int(0)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SceneMultiplayer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SceneMultiplayer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) RootPath() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetRootPath(gc).String())
}

func (self Go) SetRootPath(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRootPath(gc.String(value).NodePath(gc))
}

func (self Go) AuthCallback() gd.Callable {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Callable(class(self).GetAuthCallback(gc))
}

func (self Go) SetAuthCallback(value gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAuthCallback(value)
}

func (self Go) AuthTimeout() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAuthTimeout()))
}

func (self Go) SetAuthTimeout(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAuthTimeout(gd.Float(value))
}

func (self Go) AllowObjectDecoding() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsObjectDecodingAllowed())
}

func (self Go) SetAllowObjectDecoding(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAllowObjectDecoding(value)
}

func (self Go) RefuseNewConnections() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsRefusingNewConnections())
}

func (self Go) SetRefuseNewConnections(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRefuseNewConnections(value)
}

func (self Go) ServerRelay() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsServerRelayEnabled())
}

func (self Go) SetServerRelay(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetServerRelayEnabled(value)
}

func (self Go) MaxSyncPacketSize() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMaxSyncPacketSize()))
}

func (self Go) SetMaxSyncPacketSize(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxSyncPacketSize(gd.Int(value))
}

func (self Go) MaxDeltaPacketSize() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMaxDeltaPacketSize()))
}

func (self Go) SetMaxDeltaPacketSize(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxDeltaPacketSize(gd.Int(value))
}

//go:nosplit
func (self class) SetRootPath(path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_set_root_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRootPath(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_get_root_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Clears the current SceneMultiplayer network state (you shouldn't call this unless you know what you are doing).
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Disconnects the peer identified by [param id], removing it from the list of connected peers, and closing the underlying connection with it.
*/
//go:nosplit
func (self class) DisconnectPeer(id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_disconnect_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the IDs of the peers currently trying to authenticate with this [MultiplayerAPI].
*/
//go:nosplit
func (self class) GetAuthenticatingPeers(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_get_authenticating_peers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sends the specified [param data] to the remote peer identified by [param id] as part of an authentication message. This can be used to authenticate peers, and control when [signal MultiplayerAPI.peer_connected] is emitted (and the remote peer accepted as one of the connected peers).
*/
//go:nosplit
func (self class) SendAuth(id gd.Int, data gd.PackedByteArray) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, mmm.Get(data))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_send_auth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Mark the authentication step as completed for the remote peer identified by [param id]. The [signal MultiplayerAPI.peer_connected] signal will be emitted for this peer once the remote side also completes the authentication. No further authentication messages are expected to be received from this peer.
If a peer disconnects before completing authentication, either due to a network issue, the [member auth_timeout] expiring, or manually calling [method disconnect_peer], the [signal peer_authentication_failed] signal will be emitted instead of [signal MultiplayerAPI.peer_disconnected].
*/
//go:nosplit
func (self class) CompleteAuth(id gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_complete_auth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAuthCallback(callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_set_auth_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAuthCallback(ctx gd.Lifetime) gd.Callable {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_get_auth_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Callable](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAuthTimeout(timeout gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, timeout)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_set_auth_timeout, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAuthTimeout() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_get_auth_timeout, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRefuseNewConnections(refuse bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, refuse)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_set_refuse_new_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsRefusingNewConnections() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_is_refusing_new_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAllowObjectDecoding(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_set_allow_object_decoding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsObjectDecodingAllowed() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_is_object_decoding_allowed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetServerRelayEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_set_server_relay_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsServerRelayEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_is_server_relay_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sends the given raw [param bytes] to a specific peer identified by [param id] (see [method MultiplayerPeer.set_target_peer]). Default ID is [code]0[/code], i.e. broadcast to all peers.
*/
//go:nosplit
func (self class) SendBytes(bytes gd.PackedByteArray, id gd.Int, mode classdb.MultiplayerPeerTransferMode, channel gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bytes))
	callframe.Arg(frame, id)
	callframe.Arg(frame, mode)
	callframe.Arg(frame, channel)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_send_bytes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetMaxSyncPacketSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_get_max_sync_packet_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxSyncPacketSize(size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_set_max_sync_packet_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxDeltaPacketSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_get_max_delta_packet_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxDeltaPacketSize(size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneMultiplayer.Bind_set_max_delta_packet_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnPeerAuthenticating(cb func(id int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("peer_authenticating"), gc.Callable(cb), 0)
}


func (self Go) OnPeerAuthenticationFailed(cb func(id int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("peer_authentication_failed"), gc.Callable(cb), 0)
}


func (self Go) OnPeerPacket(cb func(id int, packet []byte)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("peer_packet"), gc.Callable(cb), 0)
}


func (self class) AsSceneMultiplayer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSceneMultiplayer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsMultiplayerAPI() MultiplayerAPI.GD { return *((*MultiplayerAPI.GD)(unsafe.Pointer(&self))) }
func (self Go) AsMultiplayerAPI() MultiplayerAPI.Go { return *((*MultiplayerAPI.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMultiplayerAPI(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMultiplayerAPI(), name)
	}
}
func init() {classdb.Register("SceneMultiplayer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
