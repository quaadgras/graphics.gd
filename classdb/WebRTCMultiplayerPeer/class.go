// Package WebRTCMultiplayerPeer provides methods for working with WebRTCMultiplayerPeer object instances.
package WebRTCMultiplayerPeer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/MultiplayerPeer"
import "graphics.gd/classdb/PacketPeer"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Dictionary"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This class constructs a full mesh of [WebRTCPeerConnection] (one connection for each peer) that can be used as a [member MultiplayerAPI.multiplayer_peer].
You can add each [WebRTCPeerConnection] via [method add_peer] or remove them via [method remove_peer]. Peers must be added in [constant WebRTCPeerConnection.STATE_NEW] state to allow it to create the appropriate channels. This class will not create offers nor set descriptions, it will only poll them, and notify connections and disconnections.
When creating the peer via [method create_client] or [method create_server] the [method MultiplayerPeer.is_server_relay_supported] method will return [code]true[/code] enabling peer exchange and packet relaying when supported by the [MultiplayerAPI] implementation.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
*/
type Instance [1]gdclass.WebRTCMultiplayerPeer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsWebRTCMultiplayerPeer() Instance
}

/*
Initialize the multiplayer peer as a server (with unique ID of [code]1[/code]). This mode enables [method MultiplayerPeer.is_server_relay_supported], allowing the upper [MultiplayerAPI] layer to perform peer exchange and packet relaying.
You can optionally specify a [param channels_config] array of [enum MultiplayerPeer.TransferMode] which will be used to create extra channels (WebRTC only supports one transfer mode per channel).
*/
func (self Instance) CreateServer() error {
	return error(class(self).CreateServer([1]Array.Any{}[0]))
}

/*
Initialize the multiplayer peer as a client with the given [param peer_id] (must be between 2 and 2147483647). In this mode, you should only call [method add_peer] once and with [param peer_id] of [code]1[/code]. This mode enables [method MultiplayerPeer.is_server_relay_supported], allowing the upper [MultiplayerAPI] layer to perform peer exchange and packet relaying.
You can optionally specify a [param channels_config] array of [enum MultiplayerPeer.TransferMode] which will be used to create extra channels (WebRTC only supports one transfer mode per channel).
*/
func (self Instance) CreateClient(peer_id int) error {
	return error(class(self).CreateClient(gd.Int(peer_id), [1]Array.Any{}[0]))
}

/*
Initialize the multiplayer peer as a mesh (i.e. all peers connect to each other) with the given [param peer_id] (must be between 1 and 2147483647).
*/
func (self Instance) CreateMesh(peer_id int) error {
	return error(class(self).CreateMesh(gd.Int(peer_id), [1]Array.Any{}[0]))
}

/*
Add a new peer to the mesh with the given [param peer_id]. The [WebRTCPeerConnection] must be in state [constant WebRTCPeerConnection.STATE_NEW].
Three channels will be created for reliable, unreliable, and ordered transport. The value of [param unreliable_lifetime] will be passed to the [code]"maxPacketLifetime"[/code] option when creating unreliable and ordered channels (see [method WebRTCPeerConnection.create_data_channel]).
*/
func (self Instance) AddPeer(peer [1]gdclass.WebRTCPeerConnection, peer_id int) error {
	return error(class(self).AddPeer(peer, gd.Int(peer_id), gd.Int(1)))
}

/*
Remove the peer with given [param peer_id] from the mesh. If the peer was connected, and [signal MultiplayerPeer.peer_connected] was emitted for it, then [signal MultiplayerPeer.peer_disconnected] will be emitted.
*/
func (self Instance) RemovePeer(peer_id int) {
	class(self).RemovePeer(gd.Int(peer_id))
}

/*
Returns [code]true[/code] if the given [param peer_id] is in the peers map (it might not be connected though).
*/
func (self Instance) HasPeer(peer_id int) bool {
	return bool(class(self).HasPeer(gd.Int(peer_id)))
}

/*
Returns a dictionary representation of the peer with given [param peer_id] with three keys. [code]"connection"[/code] containing the [WebRTCPeerConnection] to this peer, [code]"channels"[/code] an array of three [WebRTCDataChannel], and [code]"connected"[/code] a boolean representing if the peer connection is currently connected (all three channels are open).
*/
func (self Instance) GetPeer(peer_id int) Dictionary.Any {
	return Dictionary.Any(class(self).GetPeer(gd.Int(peer_id)))
}

/*
Returns a dictionary which keys are the peer ids and values the peer representation as in [method get_peer].
*/
func (self Instance) GetPeers() Dictionary.Any {
	return Dictionary.Any(class(self).GetPeers())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.WebRTCMultiplayerPeer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WebRTCMultiplayerPeer"))
	casted := Instance{*(*gdclass.WebRTCMultiplayerPeer)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Initialize the multiplayer peer as a server (with unique ID of [code]1[/code]). This mode enables [method MultiplayerPeer.is_server_relay_supported], allowing the upper [MultiplayerAPI] layer to perform peer exchange and packet relaying.
You can optionally specify a [param channels_config] array of [enum MultiplayerPeer.TransferMode] which will be used to create extra channels (WebRTC only supports one transfer mode per channel).
*/
//go:nosplit
func (self class) CreateServer(channels_config gd.Array) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(channels_config))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCMultiplayerPeer.Bind_create_server, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Initialize the multiplayer peer as a client with the given [param peer_id] (must be between 2 and 2147483647). In this mode, you should only call [method add_peer] once and with [param peer_id] of [code]1[/code]. This mode enables [method MultiplayerPeer.is_server_relay_supported], allowing the upper [MultiplayerAPI] layer to perform peer exchange and packet relaying.
You can optionally specify a [param channels_config] array of [enum MultiplayerPeer.TransferMode] which will be used to create extra channels (WebRTC only supports one transfer mode per channel).
*/
//go:nosplit
func (self class) CreateClient(peer_id gd.Int, channels_config gd.Array) error {
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	callframe.Arg(frame, pointers.Get(channels_config))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCMultiplayerPeer.Bind_create_client, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Initialize the multiplayer peer as a mesh (i.e. all peers connect to each other) with the given [param peer_id] (must be between 1 and 2147483647).
*/
//go:nosplit
func (self class) CreateMesh(peer_id gd.Int, channels_config gd.Array) error {
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	callframe.Arg(frame, pointers.Get(channels_config))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCMultiplayerPeer.Bind_create_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Add a new peer to the mesh with the given [param peer_id]. The [WebRTCPeerConnection] must be in state [constant WebRTCPeerConnection.STATE_NEW].
Three channels will be created for reliable, unreliable, and ordered transport. The value of [param unreliable_lifetime] will be passed to the [code]"maxPacketLifetime"[/code] option when creating unreliable and ordered channels (see [method WebRTCPeerConnection.create_data_channel]).
*/
//go:nosplit
func (self class) AddPeer(peer [1]gdclass.WebRTCPeerConnection, peer_id gd.Int, unreliable_lifetime gd.Int) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(peer[0])[0])
	callframe.Arg(frame, peer_id)
	callframe.Arg(frame, unreliable_lifetime)
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCMultiplayerPeer.Bind_add_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Remove the peer with given [param peer_id] from the mesh. If the peer was connected, and [signal MultiplayerPeer.peer_connected] was emitted for it, then [signal MultiplayerPeer.peer_disconnected] will be emitted.
*/
//go:nosplit
func (self class) RemovePeer(peer_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCMultiplayerPeer.Bind_remove_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the given [param peer_id] is in the peers map (it might not be connected though).
*/
//go:nosplit
func (self class) HasPeer(peer_id gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCMultiplayerPeer.Bind_has_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a dictionary representation of the peer with given [param peer_id] with three keys. [code]"connection"[/code] containing the [WebRTCPeerConnection] to this peer, [code]"channels"[/code] an array of three [WebRTCDataChannel], and [code]"connected"[/code] a boolean representing if the peer connection is currently connected (all three channels are open).
*/
//go:nosplit
func (self class) GetPeer(peer_id gd.Int) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCMultiplayerPeer.Bind_get_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a dictionary which keys are the peer ids and values the peer representation as in [method get_peer].
*/
//go:nosplit
func (self class) GetPeers() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCMultiplayerPeer.Bind_get_peers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsWebRTCMultiplayerPeer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWebRTCMultiplayerPeer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsMultiplayerPeer() MultiplayerPeer.Advanced {
	return *((*MultiplayerPeer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsMultiplayerPeer() MultiplayerPeer.Instance {
	return *((*MultiplayerPeer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsPacketPeer() PacketPeer.Advanced {
	return *((*PacketPeer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPacketPeer() PacketPeer.Instance {
	return *((*PacketPeer.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(MultiplayerPeer.Advanced(self.AsMultiplayerPeer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(MultiplayerPeer.Instance(self.AsMultiplayerPeer()), name)
	}
}
func init() {
	gdclass.Register("WebRTCMultiplayerPeer", func(ptr gd.Object) any {
		return [1]gdclass.WebRTCMultiplayerPeer{*(*gdclass.WebRTCMultiplayerPeer)(unsafe.Pointer(&ptr))}
	})
}

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
