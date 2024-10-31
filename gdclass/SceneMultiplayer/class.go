package SceneMultiplayer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/MultiplayerAPI"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
This class is the default implementation of [MultiplayerAPI], used to provide multiplayer functionalities in Godot Engine.
This implementation supports RPCs via [method Node.rpc] and [method Node.rpc_id] and requires [method MultiplayerAPI.rpc] to be passed a [Node] (it will fail for other object types).
This implementation additionally provide [SceneTree] replication via the [MultiplayerSpawner] and [MultiplayerSynchronizer] nodes, and the [SceneReplicationConfig] resource.
[b]Note:[/b] The high-level multiplayer API protocol is an implementation detail and isn't meant to be used by non-Godot servers. It may change without notice.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
*/
type Instance [1]classdb.SceneMultiplayer

/*
Clears the current SceneMultiplayer network state (you shouldn't call this unless you know what you are doing).
*/
func (self Instance) Clear() {
	class(self).Clear()
}

/*
Disconnects the peer identified by [param id], removing it from the list of connected peers, and closing the underlying connection with it.
*/
func (self Instance) DisconnectPeer(id int) {
	class(self).DisconnectPeer(gd.Int(id))
}

/*
Returns the IDs of the peers currently trying to authenticate with this [MultiplayerAPI].
*/
func (self Instance) GetAuthenticatingPeers() []int32 {
	return []int32(class(self).GetAuthenticatingPeers().AsSlice())
}

/*
Sends the specified [param data] to the remote peer identified by [param id] as part of an authentication message. This can be used to authenticate peers, and control when [signal MultiplayerAPI.peer_connected] is emitted (and the remote peer accepted as one of the connected peers).
*/
func (self Instance) SendAuth(id int, data []byte) gd.Error {
	return gd.Error(class(self).SendAuth(gd.Int(id), gd.NewPackedByteSlice(data)))
}

/*
Mark the authentication step as completed for the remote peer identified by [param id]. The [signal MultiplayerAPI.peer_connected] signal will be emitted for this peer once the remote side also completes the authentication. No further authentication messages are expected to be received from this peer.
If a peer disconnects before completing authentication, either due to a network issue, the [member auth_timeout] expiring, or manually calling [method disconnect_peer], the [signal peer_authentication_failed] signal will be emitted instead of [signal MultiplayerAPI.peer_disconnected].
*/
func (self Instance) CompleteAuth(id int) gd.Error {
	return gd.Error(class(self).CompleteAuth(gd.Int(id)))
}

/*
Sends the given raw [param bytes] to a specific peer identified by [param id] (see [method MultiplayerPeer.set_target_peer]). Default ID is [code]0[/code], i.e. broadcast to all peers.
*/
func (self Instance) SendBytes(bytes []byte) gd.Error {
	return gd.Error(class(self).SendBytes(gd.NewPackedByteSlice(bytes), gd.Int(0), 2, gd.Int(0)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.SceneMultiplayer

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SceneMultiplayer"))
	return Instance{classdb.SceneMultiplayer(object)}
}

func (self Instance) RootPath() string {
	return string(class(self).GetRootPath().String())
}

func (self Instance) SetRootPath(value string) {
	class(self).SetRootPath(gd.NewString(value).NodePath())
}

func (self Instance) AuthCallback() gd.Callable {
	return gd.Callable(class(self).GetAuthCallback())
}

func (self Instance) SetAuthCallback(value gd.Callable) {
	class(self).SetAuthCallback(value)
}

func (self Instance) AuthTimeout() float64 {
	return float64(float64(class(self).GetAuthTimeout()))
}

func (self Instance) SetAuthTimeout(value float64) {
	class(self).SetAuthTimeout(gd.Float(value))
}

func (self Instance) AllowObjectDecoding() bool {
	return bool(class(self).IsObjectDecodingAllowed())
}

func (self Instance) SetAllowObjectDecoding(value bool) {
	class(self).SetAllowObjectDecoding(value)
}

func (self Instance) RefuseNewConnections() bool {
	return bool(class(self).IsRefusingNewConnections())
}

func (self Instance) SetRefuseNewConnections(value bool) {
	class(self).SetRefuseNewConnections(value)
}

func (self Instance) ServerRelay() bool {
	return bool(class(self).IsServerRelayEnabled())
}

func (self Instance) SetServerRelay(value bool) {
	class(self).SetServerRelayEnabled(value)
}

func (self Instance) MaxSyncPacketSize() int {
	return int(int(class(self).GetMaxSyncPacketSize()))
}

func (self Instance) SetMaxSyncPacketSize(value int) {
	class(self).SetMaxSyncPacketSize(gd.Int(value))
}

func (self Instance) MaxDeltaPacketSize() int {
	return int(int(class(self).GetMaxDeltaPacketSize()))
}

func (self Instance) SetMaxDeltaPacketSize(value int) {
	class(self).SetMaxDeltaPacketSize(gd.Int(value))
}

//go:nosplit
func (self class) SetRootPath(path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_set_root_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRootPath() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_get_root_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

/*
Clears the current SceneMultiplayer network state (you shouldn't call this unless you know what you are doing).
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Disconnects the peer identified by [param id], removing it from the list of connected peers, and closing the underlying connection with it.
*/
//go:nosplit
func (self class) DisconnectPeer(id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_disconnect_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the IDs of the peers currently trying to authenticate with this [MultiplayerAPI].
*/
//go:nosplit
func (self class) GetAuthenticatingPeers() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_get_authenticating_peers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sends the specified [param data] to the remote peer identified by [param id] as part of an authentication message. This can be used to authenticate peers, and control when [signal MultiplayerAPI.peer_connected] is emitted (and the remote peer accepted as one of the connected peers).
*/
//go:nosplit
func (self class) SendAuth(id gd.Int, data gd.PackedByteArray) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(data))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_send_auth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_complete_auth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAuthCallback(callback gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(callback))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_set_auth_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAuthCallback() gd.Callable {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_get_auth_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Callable](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAuthTimeout(timeout gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, timeout)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_set_auth_timeout, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAuthTimeout() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_get_auth_timeout, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRefuseNewConnections(refuse bool) {
	var frame = callframe.New()
	callframe.Arg(frame, refuse)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_set_refuse_new_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsRefusingNewConnections() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_is_refusing_new_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowObjectDecoding(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_set_allow_object_decoding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsObjectDecodingAllowed() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_is_object_decoding_allowed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetServerRelayEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_set_server_relay_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsServerRelayEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_is_server_relay_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sends the given raw [param bytes] to a specific peer identified by [param id] (see [method MultiplayerPeer.set_target_peer]). Default ID is [code]0[/code], i.e. broadcast to all peers.
*/
//go:nosplit
func (self class) SendBytes(bytes gd.PackedByteArray, id gd.Int, mode classdb.MultiplayerPeerTransferMode, channel gd.Int) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(bytes))
	callframe.Arg(frame, id)
	callframe.Arg(frame, mode)
	callframe.Arg(frame, channel)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_send_bytes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetMaxSyncPacketSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_get_max_sync_packet_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxSyncPacketSize(size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_set_max_sync_packet_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxDeltaPacketSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_get_max_delta_packet_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxDeltaPacketSize(size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_set_max_delta_packet_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Instance) OnPeerAuthenticating(cb func(id int)) {
	self[0].AsObject().Connect(gd.NewStringName("peer_authenticating"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPeerAuthenticationFailed(cb func(id int)) {
	self[0].AsObject().Connect(gd.NewStringName("peer_authentication_failed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPeerPacket(cb func(id int, packet []byte)) {
	self[0].AsObject().Connect(gd.NewStringName("peer_packet"), gd.NewCallable(cb), 0)
}

func (self class) AsSceneMultiplayer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSceneMultiplayer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	default:
		return gd.VirtualByName(self.AsMultiplayerAPI(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsMultiplayerAPI(), name)
	}
}
func init() {
	classdb.Register("SceneMultiplayer", func(ptr gd.Object) any { return classdb.SceneMultiplayer(ptr) })
}
