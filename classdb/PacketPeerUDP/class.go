// Package PacketPeerUDP provides methods for working with PacketPeerUDP object instances.
package PacketPeerUDP

import "unsafe"
import "reflect"
import "slices"
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
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Packed"
import "graphics.gd/classdb/PacketPeer"

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
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ = slices.Delete[[]struct{}, struct{}]

/*
UDP packet peer. Can be used to send raw UDP packets as well as [Variant]s.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
*/
type Instance [1]gdclass.PacketPeerUDP

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPacketPeerUDP() Instance
}

/*
Binds this [PacketPeerUDP] to the specified [param port] and [param bind_address] with a buffer size [param recv_buf_size], allowing it to receive incoming packets.
If [param bind_address] is set to [code]"*"[/code] (default), the peer will be bound on all available addresses (both IPv4 and IPv6).
If [param bind_address] is set to [code]"0.0.0.0"[/code] (for IPv4) or [code]"::"[/code] (for IPv6), the peer will be bound to all available addresses matching that IP type.
If [param bind_address] is set to any valid address (e.g. [code]"192.168.1.101"[/code], [code]"::1"[/code], etc.), the peer will only be bound to the interface with that address (or fail if no interface with the given address exists).
*/
func (self Instance) Bind(port int) error { //gd:PacketPeerUDP.bind
	return error(gd.ToError(class(self).Bind(gd.Int(port), String.New("*"), gd.Int(65536))))
}

/*
Closes the [PacketPeerUDP]'s underlying UDP socket.
*/
func (self Instance) Close() { //gd:PacketPeerUDP.close
	class(self).Close()
}

/*
Waits for a packet to arrive on the bound address. See [method bind].
[b]Note:[/b] [method wait] can't be interrupted once it has been called. This can be worked around by allowing the other party to send a specific "death pill" packet like this:
[codeblocks]
[gdscript]
socket = PacketPeerUDP.new()
# Server
socket.set_dest_address("127.0.0.1", 789)
socket.put_packet("Time to stop".to_ascii_buffer())

# Client
while socket.wait() == OK:

	var data = socket.get_packet().get_string_from_ascii()
	if data == "Time to stop":
	    return

[/gdscript]
[csharp]
var socket = new PacketPeerUdp();
// Server
socket.SetDestAddress("127.0.0.1", 789);
socket.PutPacket("Time to stop".ToAsciiBuffer());

// Client
while (socket.Wait() == OK)

	{
	    string data = socket.GetPacket().GetStringFromASCII();
	    if (data == "Time to stop")
	    {
	        return;
	    }
	}

[/csharp]
[/codeblocks]
*/
func (self Instance) Wait() error { //gd:PacketPeerUDP.wait
	return error(gd.ToError(class(self).Wait()))
}

/*
Returns whether this [PacketPeerUDP] is bound to an address and can receive packets.
*/
func (self Instance) IsBound() bool { //gd:PacketPeerUDP.is_bound
	return bool(class(self).IsBound())
}

/*
Calling this method connects this UDP peer to the given [param host]/[param port] pair. UDP is in reality connectionless, so this option only means that incoming packets from different addresses are automatically discarded, and that outgoing packets are always sent to the connected address (future calls to [method set_dest_address] are not allowed). This method does not send any data to the remote peer, to do that, use [method PacketPeer.put_var] or [method PacketPeer.put_packet] as usual. See also [UDPServer].
[b]Note:[/b] Connecting to the remote peer does not help to protect from malicious attacks like IP spoofing, etc. Think about using an encryption technique like TLS or DTLS if you feel like your application is transferring sensitive information.
*/
func (self Instance) ConnectToHost(host string, port int) error { //gd:PacketPeerUDP.connect_to_host
	return error(gd.ToError(class(self).ConnectToHost(String.New(host), gd.Int(port))))
}

/*
Returns [code]true[/code] if the UDP socket is open and has been connected to a remote address. See [method connect_to_host].
*/
func (self Instance) IsSocketConnected() bool { //gd:PacketPeerUDP.is_socket_connected
	return bool(class(self).IsSocketConnected())
}

/*
Returns the IP of the remote peer that sent the last packet(that was received with [method PacketPeer.get_packet] or [method PacketPeer.get_var]).
*/
func (self Instance) GetPacketIp() string { //gd:PacketPeerUDP.get_packet_ip
	return string(class(self).GetPacketIp().String())
}

/*
Returns the port of the remote peer that sent the last packet(that was received with [method PacketPeer.get_packet] or [method PacketPeer.get_var]).
*/
func (self Instance) GetPacketPort() int { //gd:PacketPeerUDP.get_packet_port
	return int(int(class(self).GetPacketPort()))
}

/*
Returns the local port to which this peer is bound.
*/
func (self Instance) GetLocalPort() int { //gd:PacketPeerUDP.get_local_port
	return int(int(class(self).GetLocalPort()))
}

/*
Sets the destination address and port for sending packets and variables. A hostname will be resolved using DNS if needed.
[b]Note:[/b] [method set_broadcast_enabled] must be enabled before sending packets to a broadcast address (e.g. [code]255.255.255.255[/code]).
*/
func (self Instance) SetDestAddress(host string, port int) error { //gd:PacketPeerUDP.set_dest_address
	return error(gd.ToError(class(self).SetDestAddress(String.New(host), gd.Int(port))))
}

/*
Enable or disable sending of broadcast packets (e.g. [code]set_dest_address("255.255.255.255", 4343)[/code]. This option is disabled by default.
[b]Note:[/b] Some Android devices might require the [code]CHANGE_WIFI_MULTICAST_STATE[/code] permission and this option to be enabled to receive broadcast packets too.
*/
func (self Instance) SetBroadcastEnabled(enabled bool) { //gd:PacketPeerUDP.set_broadcast_enabled
	class(self).SetBroadcastEnabled(enabled)
}

/*
Joins the multicast group specified by [param multicast_address] using the interface identified by [param interface_name].
You can join the same multicast group with multiple interfaces. Use [method IP.get_local_interfaces] to know which are available.
[b]Note:[/b] Some Android devices might require the [code]CHANGE_WIFI_MULTICAST_STATE[/code] permission for multicast to work.
*/
func (self Instance) JoinMulticastGroup(multicast_address string, interface_name string) error { //gd:PacketPeerUDP.join_multicast_group
	return error(gd.ToError(class(self).JoinMulticastGroup(String.New(multicast_address), String.New(interface_name))))
}

/*
Removes the interface identified by [param interface_name] from the multicast group specified by [param multicast_address].
*/
func (self Instance) LeaveMulticastGroup(multicast_address string, interface_name string) error { //gd:PacketPeerUDP.leave_multicast_group
	return error(gd.ToError(class(self).LeaveMulticastGroup(String.New(multicast_address), String.New(interface_name))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PacketPeerUDP

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PacketPeerUDP"))
	casted := Instance{*(*gdclass.PacketPeerUDP)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Binds this [PacketPeerUDP] to the specified [param port] and [param bind_address] with a buffer size [param recv_buf_size], allowing it to receive incoming packets.
If [param bind_address] is set to [code]"*"[/code] (default), the peer will be bound on all available addresses (both IPv4 and IPv6).
If [param bind_address] is set to [code]"0.0.0.0"[/code] (for IPv4) or [code]"::"[/code] (for IPv6), the peer will be bound to all available addresses matching that IP type.
If [param bind_address] is set to any valid address (e.g. [code]"192.168.1.101"[/code], [code]"::1"[/code], etc.), the peer will only be bound to the interface with that address (or fail if no interface with the given address exists).
*/
//go:nosplit
func (self class) Bind(port gd.Int, bind_address String.Readable, recv_buf_size gd.Int) gd.Error { //gd:PacketPeerUDP.bind
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, pointers.Get(gd.InternalString(bind_address)))
	callframe.Arg(frame, recv_buf_size)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerUDP.Bind_bind, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Closes the [PacketPeerUDP]'s underlying UDP socket.
*/
//go:nosplit
func (self class) Close() { //gd:PacketPeerUDP.close
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerUDP.Bind_close, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Waits for a packet to arrive on the bound address. See [method bind].
[b]Note:[/b] [method wait] can't be interrupted once it has been called. This can be worked around by allowing the other party to send a specific "death pill" packet like this:
[codeblocks]
[gdscript]
socket = PacketPeerUDP.new()
# Server
socket.set_dest_address("127.0.0.1", 789)
socket.put_packet("Time to stop".to_ascii_buffer())

# Client
while socket.wait() == OK:
    var data = socket.get_packet().get_string_from_ascii()
    if data == "Time to stop":
        return
[/gdscript]
[csharp]
var socket = new PacketPeerUdp();
// Server
socket.SetDestAddress("127.0.0.1", 789);
socket.PutPacket("Time to stop".ToAsciiBuffer());

// Client
while (socket.Wait() == OK)
{
    string data = socket.GetPacket().GetStringFromASCII();
    if (data == "Time to stop")
    {
        return;
    }
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) Wait() gd.Error { //gd:PacketPeerUDP.wait
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerUDP.Bind_wait, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether this [PacketPeerUDP] is bound to an address and can receive packets.
*/
//go:nosplit
func (self class) IsBound() bool { //gd:PacketPeerUDP.is_bound
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerUDP.Bind_is_bound, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Calling this method connects this UDP peer to the given [param host]/[param port] pair. UDP is in reality connectionless, so this option only means that incoming packets from different addresses are automatically discarded, and that outgoing packets are always sent to the connected address (future calls to [method set_dest_address] are not allowed). This method does not send any data to the remote peer, to do that, use [method PacketPeer.put_var] or [method PacketPeer.put_packet] as usual. See also [UDPServer].
[b]Note:[/b] Connecting to the remote peer does not help to protect from malicious attacks like IP spoofing, etc. Think about using an encryption technique like TLS or DTLS if you feel like your application is transferring sensitive information.
*/
//go:nosplit
func (self class) ConnectToHost(host String.Readable, port gd.Int) gd.Error { //gd:PacketPeerUDP.connect_to_host
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(host)))
	callframe.Arg(frame, port)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerUDP.Bind_connect_to_host, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the UDP socket is open and has been connected to a remote address. See [method connect_to_host].
*/
//go:nosplit
func (self class) IsSocketConnected() bool { //gd:PacketPeerUDP.is_socket_connected
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerUDP.Bind_is_socket_connected, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the IP of the remote peer that sent the last packet(that was received with [method PacketPeer.get_packet] or [method PacketPeer.get_var]).
*/
//go:nosplit
func (self class) GetPacketIp() String.Readable { //gd:PacketPeerUDP.get_packet_ip
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerUDP.Bind_get_packet_ip, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the port of the remote peer that sent the last packet(that was received with [method PacketPeer.get_packet] or [method PacketPeer.get_var]).
*/
//go:nosplit
func (self class) GetPacketPort() gd.Int { //gd:PacketPeerUDP.get_packet_port
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerUDP.Bind_get_packet_port, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the local port to which this peer is bound.
*/
//go:nosplit
func (self class) GetLocalPort() gd.Int { //gd:PacketPeerUDP.get_local_port
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerUDP.Bind_get_local_port, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the destination address and port for sending packets and variables. A hostname will be resolved using DNS if needed.
[b]Note:[/b] [method set_broadcast_enabled] must be enabled before sending packets to a broadcast address (e.g. [code]255.255.255.255[/code]).
*/
//go:nosplit
func (self class) SetDestAddress(host String.Readable, port gd.Int) gd.Error { //gd:PacketPeerUDP.set_dest_address
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(host)))
	callframe.Arg(frame, port)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerUDP.Bind_set_dest_address, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Enable or disable sending of broadcast packets (e.g. [code]set_dest_address("255.255.255.255", 4343)[/code]. This option is disabled by default.
[b]Note:[/b] Some Android devices might require the [code]CHANGE_WIFI_MULTICAST_STATE[/code] permission and this option to be enabled to receive broadcast packets too.
*/
//go:nosplit
func (self class) SetBroadcastEnabled(enabled bool) { //gd:PacketPeerUDP.set_broadcast_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerUDP.Bind_set_broadcast_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Joins the multicast group specified by [param multicast_address] using the interface identified by [param interface_name].
You can join the same multicast group with multiple interfaces. Use [method IP.get_local_interfaces] to know which are available.
[b]Note:[/b] Some Android devices might require the [code]CHANGE_WIFI_MULTICAST_STATE[/code] permission for multicast to work.
*/
//go:nosplit
func (self class) JoinMulticastGroup(multicast_address String.Readable, interface_name String.Readable) gd.Error { //gd:PacketPeerUDP.join_multicast_group
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(multicast_address)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(interface_name)))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerUDP.Bind_join_multicast_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the interface identified by [param interface_name] from the multicast group specified by [param multicast_address].
*/
//go:nosplit
func (self class) LeaveMulticastGroup(multicast_address String.Readable, interface_name String.Readable) gd.Error { //gd:PacketPeerUDP.leave_multicast_group
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(multicast_address)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(interface_name)))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerUDP.Bind_leave_multicast_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPacketPeerUDP() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPacketPeerUDP() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(PacketPeer.Advanced(self.AsPacketPeer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PacketPeer.Instance(self.AsPacketPeer()), name)
	}
}
func init() {
	gdclass.Register("PacketPeerUDP", func(ptr gd.Object) any {
		return [1]gdclass.PacketPeerUDP{*(*gdclass.PacketPeerUDP)(unsafe.Pointer(&ptr))}
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
