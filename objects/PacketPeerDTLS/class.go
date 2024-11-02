package PacketPeerDTLS

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/PacketPeer"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This class represents a DTLS peer connection. It can be used to connect to a DTLS server, and is returned by [method DTLSServer.take_connection].
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
[b]Warning:[/b] TLS certificate revocation and certificate pinning are currently not supported. Revoked certificates are accepted as long as they are otherwise valid. If this is a concern, you may want to use automatically managed certificates with a short validity period.
*/
type Instance [1]classdb.PacketPeerDTLS

/*
Poll the connection to check for incoming packets. Call this frequently to update the status and keep the connection working.
*/
func (self Instance) Poll() {
	class(self).Poll()
}

/*
Connects a [param packet_peer] beginning the DTLS handshake using the underlying [PacketPeerUDP] which must be connected (see [method PacketPeerUDP.connect_to_host]). You can optionally specify the [param client_options] to be used while verifying the TLS connections. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
*/
func (self Instance) ConnectToPeer(packet_peer objects.PacketPeerUDP, hostname string) error {
	return error(class(self).ConnectToPeer(packet_peer, gd.NewString(hostname), ([1]objects.TLSOptions{}[0])))
}

/*
Returns the status of the connection. See [enum Status] for values.
*/
func (self Instance) GetStatus() classdb.PacketPeerDTLSStatus {
	return classdb.PacketPeerDTLSStatus(class(self).GetStatus())
}

/*
Disconnects this peer, terminating the DTLS session.
*/
func (self Instance) DisconnectFromPeer() {
	class(self).DisconnectFromPeer()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.PacketPeerDTLS

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PacketPeerDTLS"))
	return Instance{classdb.PacketPeerDTLS(object)}
}

/*
Poll the connection to check for incoming packets. Call this frequently to update the status and keep the connection working.
*/
//go:nosplit
func (self class) Poll() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerDTLS.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Connects a [param packet_peer] beginning the DTLS handshake using the underlying [PacketPeerUDP] which must be connected (see [method PacketPeerUDP.connect_to_host]). You can optionally specify the [param client_options] to be used while verifying the TLS connections. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
*/
//go:nosplit
func (self class) ConnectToPeer(packet_peer objects.PacketPeerUDP, hostname gd.String, client_options objects.TLSOptions) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(packet_peer[0])[0])
	callframe.Arg(frame, pointers.Get(hostname))
	callframe.Arg(frame, pointers.Get(client_options[0])[0])
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerDTLS.Bind_connect_to_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the status of the connection. See [enum Status] for values.
*/
//go:nosplit
func (self class) GetStatus() classdb.PacketPeerDTLSStatus {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.PacketPeerDTLSStatus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerDTLS.Bind_get_status, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Disconnects this peer, terminating the DTLS session.
*/
//go:nosplit
func (self class) DisconnectFromPeer() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerDTLS.Bind_disconnect_from_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsPacketPeerDTLS() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPacketPeerDTLS() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPacketPeer() PacketPeer.Advanced {
	return *((*PacketPeer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPacketPeer() PacketPeer.Instance {
	return *((*PacketPeer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsPacketPeer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsPacketPeer(), name)
	}
}
func init() {
	classdb.Register("PacketPeerDTLS", func(ptr gd.Object) any { return classdb.PacketPeerDTLS(ptr) })
}

type Status = classdb.PacketPeerDTLSStatus

const (
	/*A status representing a [PacketPeerDTLS] that is disconnected.*/
	StatusDisconnected Status = 0
	/*A status representing a [PacketPeerDTLS] that is currently performing the handshake with a remote peer.*/
	StatusHandshaking Status = 1
	/*A status representing a [PacketPeerDTLS] that is connected to a remote peer.*/
	StatusConnected Status = 2
	/*A status representing a [PacketPeerDTLS] in a generic error state.*/
	StatusError Status = 3
	/*An error status that shows a mismatch in the DTLS certificate domain presented by the host and the domain requested for validation.*/
	StatusErrorHostnameMismatch Status = 4
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
