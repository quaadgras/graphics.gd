// Package SceneMultiplayer provides methods for working with SceneMultiplayer object instances.
package SceneMultiplayer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/classdb/MultiplayerAPI"
import "graphics.gd/variant/NodePath"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any

/*
This class is the default implementation of [MultiplayerAPI], used to provide multiplayer functionalities in Godot Engine.
This implementation supports RPCs via [method Node.rpc] and [method Node.rpc_id] and requires [method MultiplayerAPI.rpc] to be passed a [Node] (it will fail for other object types).
This implementation additionally provide [SceneTree] replication via the [MultiplayerSpawner] and [MultiplayerSynchronizer] nodes, and the [SceneReplicationConfig] resource.
[b]Note:[/b] The high-level multiplayer API protocol is an implementation detail and isn't meant to be used by non-Godot servers. It may change without notice.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
*/
type Instance [1]gdclass.SceneMultiplayer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSceneMultiplayer() Instance
}

/*
Clears the current SceneMultiplayer network state (you shouldn't call this unless you know what you are doing).
*/
func (self Instance) Clear() { //gd:SceneMultiplayer.clear
	class(self).Clear()
}

/*
Disconnects the peer identified by [param id], removing it from the list of connected peers, and closing the underlying connection with it.
*/
func (self Instance) DisconnectPeer(id int) { //gd:SceneMultiplayer.disconnect_peer
	class(self).DisconnectPeer(gd.Int(id))
}

/*
Returns the IDs of the peers currently trying to authenticate with this [MultiplayerAPI].
*/
func (self Instance) GetAuthenticatingPeers() []int32 { //gd:SceneMultiplayer.get_authenticating_peers
	return []int32(class(self).GetAuthenticatingPeers().AsSlice())
}

/*
Sends the specified [param data] to the remote peer identified by [param id] as part of an authentication message. This can be used to authenticate peers, and control when [signal MultiplayerAPI.peer_connected] is emitted (and the remote peer accepted as one of the connected peers).
*/
func (self Instance) SendAuth(id int, data []byte) error { //gd:SceneMultiplayer.send_auth
	return error(gd.ToError(class(self).SendAuth(gd.Int(id), gd.NewPackedByteSlice(data))))
}

/*
Mark the authentication step as completed for the remote peer identified by [param id]. The [signal MultiplayerAPI.peer_connected] signal will be emitted for this peer once the remote side also completes the authentication. No further authentication messages are expected to be received from this peer.
If a peer disconnects before completing authentication, either due to a network issue, the [member auth_timeout] expiring, or manually calling [method disconnect_peer], the [signal peer_authentication_failed] signal will be emitted instead of [signal MultiplayerAPI.peer_disconnected].
*/
func (self Instance) CompleteAuth(id int) error { //gd:SceneMultiplayer.complete_auth
	return error(gd.ToError(class(self).CompleteAuth(gd.Int(id))))
}

/*
Sends the given raw [param bytes] to a specific peer identified by [param id] (see [method MultiplayerPeer.set_target_peer]). Default ID is [code]0[/code], i.e. broadcast to all peers.
*/
func (self Instance) SendBytes(bytes []byte) error { //gd:SceneMultiplayer.send_bytes
	return error(gd.ToError(class(self).SendBytes(gd.NewPackedByteSlice(bytes), gd.Int(0), 2, gd.Int(0))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SceneMultiplayer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SceneMultiplayer"))
	casted := Instance{*(*gdclass.SceneMultiplayer)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) RootPath() NodePath.String {
	return NodePath.String(class(self).GetRootPath().String())
}

func (self Instance) SetRootPath(value NodePath.String) {
	class(self).SetRootPath(gd.NewString(string(value)).NodePath())
}

func (self Instance) AuthCallback() Callable.Function {
	return Callable.Function(class(self).GetAuthCallback())
}

func (self Instance) SetAuthCallback(value Callable.Function) {
	class(self).SetAuthCallback(Callable.New(value))
}

func (self Instance) AuthTimeout() Float.X {
	return Float.X(Float.X(class(self).GetAuthTimeout()))
}

func (self Instance) SetAuthTimeout(value Float.X) {
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
func (self class) SetRootPath(path gd.NodePath) { //gd:SceneMultiplayer.set_root_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_set_root_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRootPath() gd.NodePath { //gd:SceneMultiplayer.get_root_path
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_get_root_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

/*
Clears the current SceneMultiplayer network state (you shouldn't call this unless you know what you are doing).
*/
//go:nosplit
func (self class) Clear() { //gd:SceneMultiplayer.clear
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Disconnects the peer identified by [param id], removing it from the list of connected peers, and closing the underlying connection with it.
*/
//go:nosplit
func (self class) DisconnectPeer(id gd.Int) { //gd:SceneMultiplayer.disconnect_peer
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_disconnect_peer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the IDs of the peers currently trying to authenticate with this [MultiplayerAPI].
*/
//go:nosplit
func (self class) GetAuthenticatingPeers() gd.PackedInt32Array { //gd:SceneMultiplayer.get_authenticating_peers
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_get_authenticating_peers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sends the specified [param data] to the remote peer identified by [param id] as part of an authentication message. This can be used to authenticate peers, and control when [signal MultiplayerAPI.peer_connected] is emitted (and the remote peer accepted as one of the connected peers).
*/
//go:nosplit
func (self class) SendAuth(id gd.Int, data gd.PackedByteArray) gd.Error { //gd:SceneMultiplayer.send_auth
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(data))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_send_auth, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Mark the authentication step as completed for the remote peer identified by [param id]. The [signal MultiplayerAPI.peer_connected] signal will be emitted for this peer once the remote side also completes the authentication. No further authentication messages are expected to be received from this peer.
If a peer disconnects before completing authentication, either due to a network issue, the [member auth_timeout] expiring, or manually calling [method disconnect_peer], the [signal peer_authentication_failed] signal will be emitted instead of [signal MultiplayerAPI.peer_disconnected].
*/
//go:nosplit
func (self class) CompleteAuth(id gd.Int) gd.Error { //gd:SceneMultiplayer.complete_auth
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_complete_auth, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAuthCallback(callback Callable.Function) { //gd:SceneMultiplayer.set_auth_callback
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_set_auth_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAuthCallback() Callable.Function { //gd:SceneMultiplayer.get_auth_callback
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_get_auth_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAuthTimeout(timeout gd.Float) { //gd:SceneMultiplayer.set_auth_timeout
	var frame = callframe.New()
	callframe.Arg(frame, timeout)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_set_auth_timeout, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAuthTimeout() gd.Float { //gd:SceneMultiplayer.get_auth_timeout
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_get_auth_timeout, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRefuseNewConnections(refuse bool) { //gd:SceneMultiplayer.set_refuse_new_connections
	var frame = callframe.New()
	callframe.Arg(frame, refuse)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_set_refuse_new_connections, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsRefusingNewConnections() bool { //gd:SceneMultiplayer.is_refusing_new_connections
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_is_refusing_new_connections, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowObjectDecoding(enable bool) { //gd:SceneMultiplayer.set_allow_object_decoding
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_set_allow_object_decoding, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsObjectDecodingAllowed() bool { //gd:SceneMultiplayer.is_object_decoding_allowed
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_is_object_decoding_allowed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetServerRelayEnabled(enabled bool) { //gd:SceneMultiplayer.set_server_relay_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_set_server_relay_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsServerRelayEnabled() bool { //gd:SceneMultiplayer.is_server_relay_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_is_server_relay_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sends the given raw [param bytes] to a specific peer identified by [param id] (see [method MultiplayerPeer.set_target_peer]). Default ID is [code]0[/code], i.e. broadcast to all peers.
*/
//go:nosplit
func (self class) SendBytes(bytes gd.PackedByteArray, id gd.Int, mode gdclass.MultiplayerPeerTransferMode, channel gd.Int) gd.Error { //gd:SceneMultiplayer.send_bytes
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(bytes))
	callframe.Arg(frame, id)
	callframe.Arg(frame, mode)
	callframe.Arg(frame, channel)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_send_bytes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetMaxSyncPacketSize() gd.Int { //gd:SceneMultiplayer.get_max_sync_packet_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_get_max_sync_packet_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxSyncPacketSize(size gd.Int) { //gd:SceneMultiplayer.set_max_sync_packet_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_set_max_sync_packet_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxDeltaPacketSize() gd.Int { //gd:SceneMultiplayer.get_max_delta_packet_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_get_max_delta_packet_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxDeltaPacketSize(size gd.Int) { //gd:SceneMultiplayer.set_max_delta_packet_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneMultiplayer.Bind_set_max_delta_packet_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnPeerAuthenticating(cb func(id int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("peer_authenticating"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPeerAuthenticationFailed(cb func(id int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("peer_authentication_failed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPeerPacket(cb func(id int, packet []byte)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("peer_packet"), gd.NewCallable(cb), 0)
}

func (self class) AsSceneMultiplayer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSceneMultiplayer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsMultiplayerAPI() MultiplayerAPI.Advanced {
	return *((*MultiplayerAPI.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsMultiplayerAPI() MultiplayerAPI.Instance {
	return *((*MultiplayerAPI.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(MultiplayerAPI.Advanced(self.AsMultiplayerAPI()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(MultiplayerAPI.Instance(self.AsMultiplayerAPI()), name)
	}
}
func init() {
	gdclass.Register("SceneMultiplayer", func(ptr gd.Object) any {
		return [1]gdclass.SceneMultiplayer{*(*gdclass.SceneMultiplayer)(unsafe.Pointer(&ptr))}
	})
}

type Error = gd.Error //gd:Error

const (
	/*Methods that return [enum Error] return [constant OK] when no error occurred.
	  Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
	  [b]Example:[/b]
	  [codeblock]
	  var error = method_that_returns_error()
	  if error != OK:
	      printerr("Failure!")

	  # Or, alternatively:
	  if error:
	      printerr("Still failing!")
	  [/codeblock]
	  [b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
	/*Generic error.*/
	Failed Error = 1
	/*Unavailable error.*/
	ErrUnavailable Error = 2
	/*Unconfigured error.*/
	ErrUnconfigured Error = 3
	/*Unauthorized error.*/
	ErrUnauthorized Error = 4
	/*Parameter range error.*/
	ErrParameterRangeError Error = 5
	/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
	/*File: Not found error.*/
	ErrFileNotFound Error = 7
	/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
	/*File: Bad path error.*/
	ErrFileBadPath Error = 9
	/*File: No permission error.*/
	ErrFileNoPermission Error = 10
	/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
	/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
	/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
	/*File: Can't read error.*/
	ErrFileCantRead Error = 14
	/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
	/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
	/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
	/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
	/*Can't open error.*/
	ErrCantOpen Error = 19
	/*Can't create error.*/
	ErrCantCreate Error = 20
	/*Query failed error.*/
	ErrQueryFailed Error = 21
	/*Already in use error.*/
	ErrAlreadyInUse Error = 22
	/*Locked error.*/
	ErrLocked Error = 23
	/*Timeout error.*/
	ErrTimeout Error = 24
	/*Can't connect error.*/
	ErrCantConnect Error = 25
	/*Can't resolve error.*/
	ErrCantResolve Error = 26
	/*Connection error.*/
	ErrConnectionError Error = 27
	/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
	/*Can't fork process error.*/
	ErrCantFork Error = 29
	/*Invalid data error.*/
	ErrInvalidData Error = 30
	/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
	/*Already exists error.*/
	ErrAlreadyExists Error = 32
	/*Does not exist error.*/
	ErrDoesNotExist Error = 33
	/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
	/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
	/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
	/*Method not found error.*/
	ErrMethodNotFound Error = 37
	/*Linking failed error.*/
	ErrLinkFailed Error = 38
	/*Script failed error.*/
	ErrScriptFailed Error = 39
	/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
	/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
	/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
	/*Parse error.*/
	ErrParseError Error = 43
	/*Busy error.*/
	ErrBusy Error = 44
	/*Skip error.*/
	ErrSkip Error = 45
	/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
	/*Bug error, caused by an implementation issue in the method.
	  [b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
	/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
)
