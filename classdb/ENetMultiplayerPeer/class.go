// Package ENetMultiplayerPeer provides methods for working with ENetMultiplayerPeer object instances.
package ENetMultiplayerPeer

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

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A MultiplayerPeer implementation that should be passed to [member MultiplayerAPI.multiplayer_peer] after being initialized as either a client, server, or mesh. Events can then be handled by connecting to [MultiplayerAPI] signals. See [ENetConnection] for more information on the ENet library wrapper.
[b]Note:[/b] ENet only uses UDP, not TCP. When forwarding the server port to make your server accessible on the public Internet, you only need to forward the server port in UDP. You can use the [UPNP] class to try to forward the server port automatically when starting the server.
*/
type Instance [1]gdclass.ENetMultiplayerPeer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsENetMultiplayerPeer() Instance
}

/*
Create server that listens to connections via [param port]. The port needs to be an available, unused port between 0 and 65535. Note that ports below 1024 are privileged and may require elevated permissions depending on the platform. To change the interface the server listens on, use [method set_bind_ip]. The default IP is the wildcard [code]"*"[/code], which listens on all available interfaces. [param max_clients] is the maximum number of clients that are allowed at once, any number up to 4095 may be used, although the achievable number of simultaneous clients may be far lower and depends on the application. For additional details on the bandwidth parameters, see [method create_client]. Returns [constant OK] if a server was created, [constant ERR_ALREADY_IN_USE] if this ENetMultiplayerPeer instance already has an open connection (in which case you need to call [method MultiplayerPeer.close] first) or [constant ERR_CANT_CREATE] if the server could not be created.
*/
func (self Instance) CreateServer(port int) error {
	return error(class(self).CreateServer(gd.Int(port), gd.Int(32), gd.Int(0), gd.Int(0), gd.Int(0)))
}

/*
Create client that connects to a server at [param address] using specified [param port]. The given address needs to be either a fully qualified domain name (e.g. [code]"www.example.com"[/code]) or an IP address in IPv4 or IPv6 format (e.g. [code]"192.168.1.1"[/code]). The [param port] is the port the server is listening on. The [param channel_count] parameter can be used to specify the number of ENet channels allocated for the connection. The [param in_bandwidth] and [param out_bandwidth] parameters can be used to limit the incoming and outgoing bandwidth to the given number of bytes per second. The default of 0 means unlimited bandwidth. Note that ENet will strategically drop packets on specific sides of a connection between peers to ensure the peer's bandwidth is not overwhelmed. The bandwidth parameters also determine the window size of a connection which limits the amount of reliable packets that may be in transit at any given time. Returns [constant OK] if a client was created, [constant ERR_ALREADY_IN_USE] if this ENetMultiplayerPeer instance already has an open connection (in which case you need to call [method MultiplayerPeer.close] first) or [constant ERR_CANT_CREATE] if the client could not be created. If [param local_port] is specified, the client will also listen to the given port; this is useful for some NAT traversal techniques.
*/
func (self Instance) CreateClient(address string, port int) error {
	return error(class(self).CreateClient(gd.NewString(address), gd.Int(port), gd.Int(0), gd.Int(0), gd.Int(0), gd.Int(0)))
}

/*
Initialize this [MultiplayerPeer] in mesh mode. The provided [param unique_id] will be used as the local peer network unique ID once assigned as the [member MultiplayerAPI.multiplayer_peer]. In the mesh configuration you will need to set up each new peer manually using [ENetConnection] before calling [method add_mesh_peer]. While this technique is more advanced, it allows for better control over the connection process (e.g. when dealing with NAT punch-through) and for better distribution of the network load (which would otherwise be more taxing on the server).
*/
func (self Instance) CreateMesh(unique_id int) error {
	return error(class(self).CreateMesh(gd.Int(unique_id)))
}

/*
Add a new remote peer with the given [param peer_id] connected to the given [param host].
[b]Note:[/b] The [param host] must have exactly one peer in the [constant ENetPacketPeer.STATE_CONNECTED] state.
*/
func (self Instance) AddMeshPeer(peer_id int, host [1]gdclass.ENetConnection) error {
	return error(class(self).AddMeshPeer(gd.Int(peer_id), host))
}

/*
The IP used when creating a server. This is set to the wildcard [code]"*"[/code] by default, which binds to all available interfaces. The given IP needs to be in IPv4 or IPv6 address format, for example: [code]"192.168.1.1"[/code].
*/
func (self Instance) SetBindIp(ip string) {
	class(self).SetBindIp(gd.NewString(ip))
}

/*
Returns the [ENetPacketPeer] associated to the given [param id].
*/
func (self Instance) GetPeer(id int) [1]gdclass.ENetPacketPeer {
	return [1]gdclass.ENetPacketPeer(class(self).GetPeer(gd.Int(id)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ENetMultiplayerPeer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ENetMultiplayerPeer"))
	casted := Instance{*(*gdclass.ENetMultiplayerPeer)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Host() [1]gdclass.ENetConnection {
	return [1]gdclass.ENetConnection(class(self).GetHost())
}

/*
Create server that listens to connections via [param port]. The port needs to be an available, unused port between 0 and 65535. Note that ports below 1024 are privileged and may require elevated permissions depending on the platform. To change the interface the server listens on, use [method set_bind_ip]. The default IP is the wildcard [code]"*"[/code], which listens on all available interfaces. [param max_clients] is the maximum number of clients that are allowed at once, any number up to 4095 may be used, although the achievable number of simultaneous clients may be far lower and depends on the application. For additional details on the bandwidth parameters, see [method create_client]. Returns [constant OK] if a server was created, [constant ERR_ALREADY_IN_USE] if this ENetMultiplayerPeer instance already has an open connection (in which case you need to call [method MultiplayerPeer.close] first) or [constant ERR_CANT_CREATE] if the server could not be created.
*/
//go:nosplit
func (self class) CreateServer(port gd.Int, max_clients gd.Int, max_channels gd.Int, in_bandwidth gd.Int, out_bandwidth gd.Int) error {
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, max_clients)
	callframe.Arg(frame, max_channels)
	callframe.Arg(frame, in_bandwidth)
	callframe.Arg(frame, out_bandwidth)
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetMultiplayerPeer.Bind_create_server, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Create client that connects to a server at [param address] using specified [param port]. The given address needs to be either a fully qualified domain name (e.g. [code]"www.example.com"[/code]) or an IP address in IPv4 or IPv6 format (e.g. [code]"192.168.1.1"[/code]). The [param port] is the port the server is listening on. The [param channel_count] parameter can be used to specify the number of ENet channels allocated for the connection. The [param in_bandwidth] and [param out_bandwidth] parameters can be used to limit the incoming and outgoing bandwidth to the given number of bytes per second. The default of 0 means unlimited bandwidth. Note that ENet will strategically drop packets on specific sides of a connection between peers to ensure the peer's bandwidth is not overwhelmed. The bandwidth parameters also determine the window size of a connection which limits the amount of reliable packets that may be in transit at any given time. Returns [constant OK] if a client was created, [constant ERR_ALREADY_IN_USE] if this ENetMultiplayerPeer instance already has an open connection (in which case you need to call [method MultiplayerPeer.close] first) or [constant ERR_CANT_CREATE] if the client could not be created. If [param local_port] is specified, the client will also listen to the given port; this is useful for some NAT traversal techniques.
*/
//go:nosplit
func (self class) CreateClient(address gd.String, port gd.Int, channel_count gd.Int, in_bandwidth gd.Int, out_bandwidth gd.Int, local_port gd.Int) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(address))
	callframe.Arg(frame, port)
	callframe.Arg(frame, channel_count)
	callframe.Arg(frame, in_bandwidth)
	callframe.Arg(frame, out_bandwidth)
	callframe.Arg(frame, local_port)
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetMultiplayerPeer.Bind_create_client, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Initialize this [MultiplayerPeer] in mesh mode. The provided [param unique_id] will be used as the local peer network unique ID once assigned as the [member MultiplayerAPI.multiplayer_peer]. In the mesh configuration you will need to set up each new peer manually using [ENetConnection] before calling [method add_mesh_peer]. While this technique is more advanced, it allows for better control over the connection process (e.g. when dealing with NAT punch-through) and for better distribution of the network load (which would otherwise be more taxing on the server).
*/
//go:nosplit
func (self class) CreateMesh(unique_id gd.Int) error {
	var frame = callframe.New()
	callframe.Arg(frame, unique_id)
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetMultiplayerPeer.Bind_create_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Add a new remote peer with the given [param peer_id] connected to the given [param host].
[b]Note:[/b] The [param host] must have exactly one peer in the [constant ENetPacketPeer.STATE_CONNECTED] state.
*/
//go:nosplit
func (self class) AddMeshPeer(peer_id gd.Int, host [1]gdclass.ENetConnection) error {
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	callframe.Arg(frame, pointers.Get(host[0])[0])
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetMultiplayerPeer.Bind_add_mesh_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
The IP used when creating a server. This is set to the wildcard [code]"*"[/code] by default, which binds to all available interfaces. The given IP needs to be in IPv4 or IPv6 address format, for example: [code]"192.168.1.1"[/code].
*/
//go:nosplit
func (self class) SetBindIp(ip gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(ip))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetMultiplayerPeer.Bind_set_bind_ip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHost() [1]gdclass.ENetConnection {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetMultiplayerPeer.Bind_get_host, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.ENetConnection{gd.PointerWithOwnershipTransferredToGo[gdclass.ENetConnection](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the [ENetPacketPeer] associated to the given [param id].
*/
//go:nosplit
func (self class) GetPeer(id gd.Int) [1]gdclass.ENetPacketPeer {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetMultiplayerPeer.Bind_get_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.ENetPacketPeer{gd.PointerWithOwnershipTransferredToGo[gdclass.ENetPacketPeer](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsENetMultiplayerPeer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsENetMultiplayerPeer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("ENetMultiplayerPeer", func(ptr gd.Object) any {
		return [1]gdclass.ENetMultiplayerPeer{*(*gdclass.ENetMultiplayerPeer)(unsafe.Pointer(&ptr))}
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
