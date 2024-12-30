package MultiplayerAPI

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/variant/Array"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Base class for high-level multiplayer API implementations. See also [MultiplayerPeer].
By default, [SceneTree] has a reference to an implementation of this class and uses it to provide multiplayer capabilities (i.e. RPCs) across the whole scene.
It is possible to override the MultiplayerAPI instance used by specific tree branches by calling the [method SceneTree.set_multiplayer] method, effectively allowing to run both client and server in the same scene.
It is also possible to extend or replace the default implementation via scripting or native extensions. See [MultiplayerAPIExtension] for details about extensions, [SceneMultiplayer] for the details about the default implementation.
*/
type Instance [1]classdb.MultiplayerAPI
type Any interface {
	gd.IsClass
	AsMultiplayerAPI() Instance
}

/*
Returns [code]true[/code] if there is a [member multiplayer_peer] set.
*/
func (self Instance) HasMultiplayerPeer() bool {
	return bool(class(self).HasMultiplayerPeer())
}

/*
Returns the unique peer ID of this MultiplayerAPI's [member multiplayer_peer].
*/
func (self Instance) GetUniqueId() int {
	return int(int(class(self).GetUniqueId()))
}

/*
Returns [code]true[/code] if this MultiplayerAPI's [member multiplayer_peer] is valid and in server mode (listening for connections).
*/
func (self Instance) IsServer() bool {
	return bool(class(self).IsServer())
}

/*
Returns the sender's peer ID for the RPC currently being executed.
[b]Note:[/b] This method returns [code]0[/code] when called outside of an RPC. As such, the original peer ID may be lost when code execution is delayed (such as with GDScript's [code]await[/code] keyword).
*/
func (self Instance) GetRemoteSenderId() int {
	return int(int(class(self).GetRemoteSenderId()))
}

/*
Method used for polling the MultiplayerAPI. You only need to worry about this if you set [member SceneTree.multiplayer_poll] to [code]false[/code]. By default, [SceneTree] will poll its MultiplayerAPI(s) for you.
[b]Note:[/b] This method results in RPCs being called, so they will be executed in the same context of this function (e.g. [code]_process[/code], [code]physics[/code], [Thread]).
*/
func (self Instance) Poll() error {
	return error(class(self).Poll())
}

/*
Sends an RPC to the target [param peer]. The given [param method] will be called on the remote [param object] with the provided [param arguments]. The RPC may also be called locally depending on the implementation and RPC configuration. See [method Node.rpc] and [method Node.rpc_config].
[b]Note:[/b] Prefer using [method Node.rpc], [method Node.rpc_id], or [code]my_method.rpc(peer, arg1, arg2, ...)[/code] (in GDScript), since they are faster. This method is mostly useful in conjunction with [MultiplayerAPIExtension] when augmenting or replacing the multiplayer capabilities.
*/
func (self Instance) Rpc(peer int, obj gd.Object, method string) error {
	return error(class(self).Rpc(gd.Int(peer), obj, gd.NewStringName(method), [1]Array.Any{}[0]))
}

/*
Notifies the MultiplayerAPI of a new [param configuration] for the given [param object]. This method is used internally by [SceneTree] to configure the root path for this MultiplayerAPI (passing [code]null[/code] and a valid [NodePath] as [param configuration]). This method can be further used by MultiplayerAPI implementations to provide additional features, refer to specific implementation (e.g. [SceneMultiplayer]) for details on how they use it.
[b]Note:[/b] This method is mostly relevant when extending or overriding the MultiplayerAPI behavior via [MultiplayerAPIExtension].
*/
func (self Instance) ObjectConfigurationAdd(obj gd.Object, configuration any) error {
	return error(class(self).ObjectConfigurationAdd(obj, gd.NewVariant(configuration)))
}

/*
Notifies the MultiplayerAPI to remove a [param configuration] for the given [param object]. This method is used internally by [SceneTree] to configure the root path for this MultiplayerAPI (passing [code]null[/code] and an empty [NodePath] as [param configuration]). This method can be further used by MultiplayerAPI implementations to provide additional features, refer to specific implementation (e.g. [SceneMultiplayer]) for details on how they use it.
[b]Note:[/b] This method is mostly relevant when extending or overriding the MultiplayerAPI behavior via [MultiplayerAPIExtension].
*/
func (self Instance) ObjectConfigurationRemove(obj gd.Object, configuration any) error {
	return error(class(self).ObjectConfigurationRemove(obj, gd.NewVariant(configuration)))
}

/*
Returns the peer IDs of all connected peers of this MultiplayerAPI's [member multiplayer_peer].
*/
func (self Instance) GetPeers() []int32 {
	return []int32(class(self).GetPeers().AsSlice())
}

/*
Sets the default MultiplayerAPI implementation class. This method can be used by modules and extensions to configure which implementation will be used by [SceneTree] when the engine starts.
*/
func SetDefaultInterface(interface_name string) {
	self := MultiplayerAPI{}
	class(self).SetDefaultInterface(gd.NewStringName(interface_name))
}

/*
Returns the default MultiplayerAPI implementation class name. This is usually [code]"SceneMultiplayer"[/code] when [SceneMultiplayer] is available. See [method set_default_interface].
*/
func GetDefaultInterface() string {
	self := MultiplayerAPI{}
	return string(class(self).GetDefaultInterface().String())
}

/*
Returns a new instance of the default MultiplayerAPI.
*/
func CreateDefaultInterface() objects.MultiplayerAPI {
	self := MultiplayerAPI{}
	return objects.MultiplayerAPI(class(self).CreateDefaultInterface())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.MultiplayerAPI

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MultiplayerAPI"))
	return Instance{classdb.MultiplayerAPI(object)}
}

func (self Instance) MultiplayerPeer() objects.MultiplayerPeer {
	return objects.MultiplayerPeer(class(self).GetMultiplayerPeer())
}

func (self Instance) SetMultiplayerPeer(value objects.MultiplayerPeer) {
	class(self).SetMultiplayerPeer(value)
}

/*
Returns [code]true[/code] if there is a [member multiplayer_peer] set.
*/
//go:nosplit
func (self class) HasMultiplayerPeer() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_has_multiplayer_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetMultiplayerPeer() objects.MultiplayerPeer {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_get_multiplayer_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.MultiplayerPeer{classdb.MultiplayerPeer(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMultiplayerPeer(peer objects.MultiplayerPeer) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(peer[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_set_multiplayer_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the unique peer ID of this MultiplayerAPI's [member multiplayer_peer].
*/
//go:nosplit
func (self class) GetUniqueId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_get_unique_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this MultiplayerAPI's [member multiplayer_peer] is valid and in server mode (listening for connections).
*/
//go:nosplit
func (self class) IsServer() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_is_server, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_get_remote_sender_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Method used for polling the MultiplayerAPI. You only need to worry about this if you set [member SceneTree.multiplayer_poll] to [code]false[/code]. By default, [SceneTree] will poll its MultiplayerAPI(s) for you.
[b]Note:[/b] This method results in RPCs being called, so they will be executed in the same context of this function (e.g. [code]_process[/code], [code]physics[/code], [Thread]).
*/
//go:nosplit
func (self class) Poll() error {
	var frame = callframe.New()
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sends an RPC to the target [param peer]. The given [param method] will be called on the remote [param object] with the provided [param arguments]. The RPC may also be called locally depending on the implementation and RPC configuration. See [method Node.rpc] and [method Node.rpc_config].
[b]Note:[/b] Prefer using [method Node.rpc], [method Node.rpc_id], or [code]my_method.rpc(peer, arg1, arg2, ...)[/code] (in GDScript), since they are faster. This method is mostly useful in conjunction with [MultiplayerAPIExtension] when augmenting or replacing the multiplayer capabilities.
*/
//go:nosplit
func (self class) Rpc(peer gd.Int, obj gd.Object, method gd.StringName, arguments gd.Array) error {
	var frame = callframe.New()
	callframe.Arg(frame, peer)
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(obj))
	callframe.Arg(frame, pointers.Get(method))
	callframe.Arg(frame, pointers.Get(arguments))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_rpc, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Notifies the MultiplayerAPI of a new [param configuration] for the given [param object]. This method is used internally by [SceneTree] to configure the root path for this MultiplayerAPI (passing [code]null[/code] and a valid [NodePath] as [param configuration]). This method can be further used by MultiplayerAPI implementations to provide additional features, refer to specific implementation (e.g. [SceneMultiplayer]) for details on how they use it.
[b]Note:[/b] This method is mostly relevant when extending or overriding the MultiplayerAPI behavior via [MultiplayerAPIExtension].
*/
//go:nosplit
func (self class) ObjectConfigurationAdd(obj gd.Object, configuration gd.Variant) error {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(obj))
	callframe.Arg(frame, pointers.Get(configuration))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_object_configuration_add, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Notifies the MultiplayerAPI to remove a [param configuration] for the given [param object]. This method is used internally by [SceneTree] to configure the root path for this MultiplayerAPI (passing [code]null[/code] and an empty [NodePath] as [param configuration]). This method can be further used by MultiplayerAPI implementations to provide additional features, refer to specific implementation (e.g. [SceneMultiplayer]) for details on how they use it.
[b]Note:[/b] This method is mostly relevant when extending or overriding the MultiplayerAPI behavior via [MultiplayerAPIExtension].
*/
//go:nosplit
func (self class) ObjectConfigurationRemove(obj gd.Object, configuration gd.Variant) error {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(obj))
	callframe.Arg(frame, pointers.Get(configuration))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_object_configuration_remove, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the peer IDs of all connected peers of this MultiplayerAPI's [member multiplayer_peer].
*/
//go:nosplit
func (self class) GetPeers() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_get_peers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the default MultiplayerAPI implementation class. This method can be used by modules and extensions to configure which implementation will be used by [SceneTree] when the engine starts.
*/
//go:nosplit
func (self class) SetDefaultInterface(interface_name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(interface_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_set_default_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the default MultiplayerAPI implementation class name. This is usually [code]"SceneMultiplayer"[/code] when [SceneMultiplayer] is available. See [method set_default_interface].
*/
//go:nosplit
func (self class) GetDefaultInterface() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_get_default_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a new instance of the default MultiplayerAPI.
*/
//go:nosplit
func (self class) CreateDefaultInterface() objects.MultiplayerAPI {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_create_default_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.MultiplayerAPI{classdb.MultiplayerAPI(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self Instance) OnPeerConnected(cb func(id int)) {
	self[0].AsObject().Connect(gd.NewStringName("peer_connected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPeerDisconnected(cb func(id int)) {
	self[0].AsObject().Connect(gd.NewStringName("peer_disconnected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnConnectedToServer(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("connected_to_server"), gd.NewCallable(cb), 0)
}

func (self Instance) OnConnectionFailed(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("connection_failed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnServerDisconnected(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("server_disconnected"), gd.NewCallable(cb), 0)
}

func (self class) AsMultiplayerAPI() Advanced     { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMultiplayerAPI() Instance  { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("MultiplayerAPI", func(ptr gd.Object) any { return classdb.MultiplayerAPI(ptr) })
}

type RPCMode = classdb.MultiplayerAPIRPCMode

const (
	/*Used with [method Node.rpc_config] to disable a method or property for all RPC calls, making it unavailable. Default for all methods.*/
	RpcModeDisabled RPCMode = 0
	/*Used with [method Node.rpc_config] to set a method to be callable remotely by any peer. Analogous to the [code]@rpc("any_peer")[/code] annotation. Calls are accepted from all remote peers, no matter if they are node's authority or not.*/
	RpcModeAnyPeer RPCMode = 1
	/*Used with [method Node.rpc_config] to set a method to be callable remotely only by the current multiplayer authority (which is the server by default). Analogous to the [code]@rpc("authority")[/code] annotation. See [method Node.set_multiplayer_authority].*/
	RpcModeAuthority RPCMode = 2
)

type Error int

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
