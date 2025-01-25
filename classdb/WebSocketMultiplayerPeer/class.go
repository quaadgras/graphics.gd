// Package WebSocketMultiplayerPeer provides methods for working with WebSocketMultiplayerPeer object instances.
package WebSocketMultiplayerPeer

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
import "graphics.gd/classdb/MultiplayerPeer"
import "graphics.gd/classdb/PacketPeer"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any

/*
Base class for WebSocket server and client, allowing them to be used as multiplayer peer for the [MultiplayerAPI].
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
*/
type Instance [1]gdclass.WebSocketMultiplayerPeer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsWebSocketMultiplayerPeer() Instance
}

/*
Starts a new multiplayer client connecting to the given [param url]. TLS certificates will be verified against the hostname when connecting using the [code]wss://[/code] protocol. You can pass the optional [param tls_client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
[b]Note:[/b] It is recommended to specify the scheme part of the URL, i.e. the [param url] should start with either [code]ws://[/code] or [code]wss://[/code].
*/
func (self Instance) CreateClient(url string) error {
	return error(gd.ToError(class(self).CreateClient(gd.NewString(url), [1][1]gdclass.TLSOptions{}[0])))
}

/*
Starts a new multiplayer server listening on the given [param port]. You can optionally specify a [param bind_address], and provide valid [param tls_server_options] to use TLS. See [method TLSOptions.server].
*/
func (self Instance) CreateServer(port int) error {
	return error(gd.ToError(class(self).CreateServer(gd.Int(port), gd.NewString("*"), [1][1]gdclass.TLSOptions{}[0])))
}

/*
Returns the [WebSocketPeer] associated to the given [param peer_id].
*/
func (self Instance) GetPeer(peer_id int) [1]gdclass.WebSocketPeer {
	return [1]gdclass.WebSocketPeer(class(self).GetPeer(gd.Int(peer_id)))
}

/*
Returns the IP address of the given peer.
*/
func (self Instance) GetPeerAddress(id int) string {
	return string(class(self).GetPeerAddress(gd.Int(id)).String())
}

/*
Returns the remote port of the given peer.
*/
func (self Instance) GetPeerPort(id int) int {
	return int(int(class(self).GetPeerPort(gd.Int(id))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.WebSocketMultiplayerPeer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WebSocketMultiplayerPeer"))
	casted := Instance{*(*gdclass.WebSocketMultiplayerPeer)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) SupportedProtocols() []string {
	return []string(class(self).GetSupportedProtocols().Strings())
}

func (self Instance) SetSupportedProtocols(value []string) {
	class(self).SetSupportedProtocols(gd.NewPackedStringSlice(value))
}

func (self Instance) HandshakeHeaders() []string {
	return []string(class(self).GetHandshakeHeaders().Strings())
}

func (self Instance) SetHandshakeHeaders(value []string) {
	class(self).SetHandshakeHeaders(gd.NewPackedStringSlice(value))
}

func (self Instance) InboundBufferSize() int {
	return int(int(class(self).GetInboundBufferSize()))
}

func (self Instance) SetInboundBufferSize(value int) {
	class(self).SetInboundBufferSize(gd.Int(value))
}

func (self Instance) OutboundBufferSize() int {
	return int(int(class(self).GetOutboundBufferSize()))
}

func (self Instance) SetOutboundBufferSize(value int) {
	class(self).SetOutboundBufferSize(gd.Int(value))
}

func (self Instance) HandshakeTimeout() Float.X {
	return Float.X(Float.X(class(self).GetHandshakeTimeout()))
}

func (self Instance) SetHandshakeTimeout(value Float.X) {
	class(self).SetHandshakeTimeout(gd.Float(value))
}

func (self Instance) MaxQueuedPackets() int {
	return int(int(class(self).GetMaxQueuedPackets()))
}

func (self Instance) SetMaxQueuedPackets(value int) {
	class(self).SetMaxQueuedPackets(gd.Int(value))
}

/*
Starts a new multiplayer client connecting to the given [param url]. TLS certificates will be verified against the hostname when connecting using the [code]wss://[/code] protocol. You can pass the optional [param tls_client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
[b]Note:[/b] It is recommended to specify the scheme part of the URL, i.e. the [param url] should start with either [code]ws://[/code] or [code]wss://[/code].
*/
//go:nosplit
func (self class) CreateClient(url gd.String, tls_client_options [1]gdclass.TLSOptions) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(url))
	callframe.Arg(frame, pointers.Get(tls_client_options[0])[0])
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_create_client, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Starts a new multiplayer server listening on the given [param port]. You can optionally specify a [param bind_address], and provide valid [param tls_server_options] to use TLS. See [method TLSOptions.server].
*/
//go:nosplit
func (self class) CreateServer(port gd.Int, bind_address gd.String, tls_server_options [1]gdclass.TLSOptions) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, pointers.Get(bind_address))
	callframe.Arg(frame, pointers.Get(tls_server_options[0])[0])
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_create_server, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [WebSocketPeer] associated to the given [param peer_id].
*/
//go:nosplit
func (self class) GetPeer(peer_id gd.Int) [1]gdclass.WebSocketPeer {
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_get_peer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.WebSocketPeer{gd.PointerWithOwnershipTransferredToGo[gdclass.WebSocketPeer](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the IP address of the given peer.
*/
//go:nosplit
func (self class) GetPeerAddress(id gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_get_peer_address, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the remote port of the given peer.
*/
//go:nosplit
func (self class) GetPeerPort(id gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_get_peer_port, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetSupportedProtocols() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_get_supported_protocols, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSupportedProtocols(protocols gd.PackedStringArray) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(protocols))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_set_supported_protocols, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHandshakeHeaders() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_get_handshake_headers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHandshakeHeaders(protocols gd.PackedStringArray) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(protocols))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_set_handshake_headers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInboundBufferSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_get_inbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInboundBufferSize(buffer_size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, buffer_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_set_inbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOutboundBufferSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_get_outbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOutboundBufferSize(buffer_size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, buffer_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_set_outbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHandshakeTimeout() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_get_handshake_timeout, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHandshakeTimeout(timeout gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, timeout)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_set_handshake_timeout, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetMaxQueuedPackets(max_queued_packets gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, max_queued_packets)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_set_max_queued_packets, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxQueuedPackets() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_get_max_queued_packets, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsWebSocketMultiplayerPeer() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWebSocketMultiplayerPeer() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
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
	gdclass.Register("WebSocketMultiplayerPeer", func(ptr gd.Object) any {
		return [1]gdclass.WebSocketMultiplayerPeer{*(*gdclass.WebSocketMultiplayerPeer)(unsafe.Pointer(&ptr))}
	})
}

type Error = gd.Error

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
