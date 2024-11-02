package UDPServer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A simple server that opens a UDP socket and returns connected [PacketPeerUDP] upon receiving new packets. See also [method PacketPeerUDP.connect_to_host].
After starting the server ([method listen]), you will need to [method poll] it at regular intervals (e.g. inside [method Node._process]) for it to process new packets, delivering them to the appropriate [PacketPeerUDP], and taking new connections.
Below a small example of how it can be used:
[codeblocks]
[gdscript]
# server_node.gd
class_name ServerNode
extends Node

var server := UDPServer.new()
var peers = []

func _ready():

	server.listen(4242)

func _process(delta):

	server.poll() # Important!
	if server.is_connection_available():
	    var peer: PacketPeerUDP = server.take_connection()
	    var packet = peer.get_packet()
	    print("Accepted peer: %s:%s" % [peer.get_packet_ip(), peer.get_packet_port()])
	    print("Received data: %s" % [packet.get_string_from_utf8()])
	    # Reply so it knows we received the message.
	    peer.put_packet(packet)
	    # Keep a reference so we can keep contacting the remote peer.
	    peers.append(peer)

	for i in range(0, peers.size()):
	    pass # Do something with the connected peers.

[/gdscript]
[csharp]
// ServerNode.cs
using Godot;
using System.Collections.Generic;

public partial class ServerNode : Node

	{
	    private UdpServer _server = new UdpServer();
	    private List<PacketPeerUdp> _peers  = new List<PacketPeerUdp>();

	    public override void _Ready()
	    {
	        _server.Listen(4242);
	    }

	    public override void _Process(double delta)
	    {
	        _server.Poll(); // Important!
	        if (_server.IsConnectionAvailable())
	        {
	            PacketPeerUdp peer = _server.TakeConnection();
	            byte[] packet = peer.GetPacket();
	            GD.Print($"Accepted Peer: {peer.GetPacketIP()}:{peer.GetPacketPort()}");
	            GD.Print($"Received Data: {packet.GetStringFromUtf8()}");
	            // Reply so it knows we received the message.
	            peer.PutPacket(packet);
	            // Keep a reference so we can keep contacting the remote peer.
	            _peers.Add(peer);
	        }
	        foreach (var peer in _peers)
	        {
	            // Do something with the peers.
	        }
	    }
	}

[/csharp]
[/codeblocks]
[codeblocks]
[gdscript]
# client_node.gd
class_name ClientNode
extends Node

var udp := PacketPeerUDP.new()
var connected = false

func _ready():

	udp.connect_to_host("127.0.0.1", 4242)

func _process(delta):

	if !connected:
	    # Try to contact server
	    udp.put_packet("The answer is... 42!".to_utf8_buffer())
	if udp.get_available_packet_count() > 0:
	    print("Connected: %s" % udp.get_packet().get_string_from_utf8())
	    connected = true

[/gdscript]
[csharp]
// ClientNode.cs
using Godot;

public partial class ClientNode : Node

	{
	    private PacketPeerUdp _udp = new PacketPeerUdp();
	    private bool _connected = false;

	    public override void _Ready()
	    {
	        _udp.ConnectToHost("127.0.0.1", 4242);
	    }

	    public override void _Process(double delta)
	    {
	        if (!_connected)
	        {
	            // Try to contact server
	            _udp.PutPacket("The Answer Is..42!".ToUtf8Buffer());
	        }
	        if (_udp.GetAvailablePacketCount() > 0)
	        {
	            GD.Print($"Connected: {_udp.GetPacket().GetStringFromUtf8()}");
	            _connected = true;
	        }
	    }
	}

[/csharp]
[/codeblocks]
*/
type Instance [1]classdb.UDPServer

/*
Starts the server by opening a UDP socket listening on the given [param port]. You can optionally specify a [param bind_address] to only listen for packets sent to that address. See also [method PacketPeerUDP.bind].
*/
func (self Instance) Listen(port int) error {
	return error(class(self).Listen(gd.Int(port), gd.NewString("*")))
}

/*
Call this method at regular intervals (e.g. inside [method Node._process]) to process new packets. And packet from known address/port pair will be delivered to the appropriate [PacketPeerUDP], any packet received from an unknown address/port pair will be added as a pending connection (see [method is_connection_available], [method take_connection]). The maximum number of pending connection is defined via [member max_pending_connections].
*/
func (self Instance) Poll() error {
	return error(class(self).Poll())
}

/*
Returns [code]true[/code] if a packet with a new address/port combination was received on the socket.
*/
func (self Instance) IsConnectionAvailable() bool {
	return bool(class(self).IsConnectionAvailable())
}

/*
Returns the local port this server is listening to.
*/
func (self Instance) GetLocalPort() int {
	return int(int(class(self).GetLocalPort()))
}

/*
Returns [code]true[/code] if the socket is open and listening on a port.
*/
func (self Instance) IsListening() bool {
	return bool(class(self).IsListening())
}

/*
Returns the first pending connection (connected to the appropriate address/port). Will return [code]null[/code] if no new connection is available. See also [method is_connection_available], [method PacketPeerUDP.connect_to_host].
*/
func (self Instance) TakeConnection() objects.PacketPeerUDP {
	return objects.PacketPeerUDP(class(self).TakeConnection())
}

/*
Stops the server, closing the UDP socket if open. Will close all connected [PacketPeerUDP] accepted via [method take_connection] (remote peers will not be notified).
*/
func (self Instance) Stop() {
	class(self).Stop()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.UDPServer

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("UDPServer"))
	return Instance{classdb.UDPServer(object)}
}

func (self Instance) MaxPendingConnections() int {
	return int(int(class(self).GetMaxPendingConnections()))
}

func (self Instance) SetMaxPendingConnections(value int) {
	class(self).SetMaxPendingConnections(gd.Int(value))
}

/*
Starts the server by opening a UDP socket listening on the given [param port]. You can optionally specify a [param bind_address] to only listen for packets sent to that address. See also [method PacketPeerUDP.bind].
*/
//go:nosplit
func (self class) Listen(port gd.Int, bind_address gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, pointers.Get(bind_address))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UDPServer.Bind_listen, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Call this method at regular intervals (e.g. inside [method Node._process]) to process new packets. And packet from known address/port pair will be delivered to the appropriate [PacketPeerUDP], any packet received from an unknown address/port pair will be added as a pending connection (see [method is_connection_available], [method take_connection]). The maximum number of pending connection is defined via [member max_pending_connections].
*/
//go:nosplit
func (self class) Poll() error {
	var frame = callframe.New()
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UDPServer.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if a packet with a new address/port combination was received on the socket.
*/
//go:nosplit
func (self class) IsConnectionAvailable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UDPServer.Bind_is_connection_available, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the local port this server is listening to.
*/
//go:nosplit
func (self class) GetLocalPort() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UDPServer.Bind_get_local_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the socket is open and listening on a port.
*/
//go:nosplit
func (self class) IsListening() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UDPServer.Bind_is_listening, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the first pending connection (connected to the appropriate address/port). Will return [code]null[/code] if no new connection is available. See also [method is_connection_available], [method PacketPeerUDP.connect_to_host].
*/
//go:nosplit
func (self class) TakeConnection() objects.PacketPeerUDP {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UDPServer.Bind_take_connection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.PacketPeerUDP{classdb.PacketPeerUDP(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Stops the server, closing the UDP socket if open. Will close all connected [PacketPeerUDP] accepted via [method take_connection] (remote peers will not be notified).
*/
//go:nosplit
func (self class) Stop() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UDPServer.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetMaxPendingConnections(max_pending_connections gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, max_pending_connections)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UDPServer.Bind_set_max_pending_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxPendingConnections() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UDPServer.Bind_get_max_pending_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsUDPServer() Advanced          { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsUDPServer() Instance       { return *((*Instance)(unsafe.Pointer(&self))) }
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
func init() { classdb.Register("UDPServer", func(ptr gd.Object) any { return classdb.UDPServer(ptr) }) }

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
