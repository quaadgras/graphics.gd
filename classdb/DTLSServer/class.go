// Package DTLSServer provides methods for working with DTLSServer object instances.
package DTLSServer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This class is used to store the state of a DTLS server. Upon [method setup] it converts connected [PacketPeerUDP] to [PacketPeerDTLS] accepting them via [method take_connection] as DTLS clients. Under the hood, this class is used to store the DTLS state and cookies of the server. The reason of why the state and cookies are needed is outside of the scope of this documentation.
Below a small example of how to use it:
[codeblocks]
[gdscript]
# server_node.gd
extends Node

var dtls := DTLSServer.new()
var server := UDPServer.new()
var peers = []

func _ready():

	server.listen(4242)
	var key = load("key.key") # Your private key.
	var cert = load("cert.crt") # Your X509 certificate.
	dtls.setup(key, cert)

func _process(delta):

	while server.is_connection_available():
	    var peer: PacketPeerUDP = server.take_connection()
	    var dtls_peer: PacketPeerDTLS = dtls.take_connection(peer)
	    if dtls_peer.get_status() != PacketPeerDTLS.STATUS_HANDSHAKING:
	        continue # It is normal that 50% of the connections fails due to cookie exchange.
	    print("Peer connected!")
	    peers.append(dtls_peer)

	for p in peers:
	    p.poll() # Must poll to update the state.
	    if p.get_status() == PacketPeerDTLS.STATUS_CONNECTED:
	        while p.get_available_packet_count() > 0:
	            print("Received message from client: %s" % p.get_packet().get_string_from_utf8())
	            p.put_packet("Hello DTLS client".to_utf8_buffer())

[/gdscript]
[csharp]
// ServerNode.cs
using Godot;

public partial class ServerNode : Node

	{
	    private DtlsServer _dtls = new DtlsServer();
	    private UdpServer _server = new UdpServer();
	    private Godot.Collections.Array<PacketPeerDtls> _peers = new Godot.Collections.Array<PacketPeerDtls>();

	    public override void _Ready()
	    {
	        _server.Listen(4242);
	        var key = GD.Load<CryptoKey>("key.key"); // Your private key.
	        var cert = GD.Load<X509Certificate>("cert.crt"); // Your X509 certificate.
	        _dtls.Setup(key, cert);
	    }

	    public override void _Process(double delta)
	    {
	        while (Server.IsConnectionAvailable())
	        {
	            PacketPeerUdp peer = _server.TakeConnection();
	            PacketPeerDtls dtlsPeer = _dtls.TakeConnection(peer);
	            if (dtlsPeer.GetStatus() != PacketPeerDtls.Status.Handshaking)
	            {
	                continue; // It is normal that 50% of the connections fails due to cookie exchange.
	            }
	            GD.Print("Peer connected!");
	            _peers.Add(dtlsPeer);
	        }

	        foreach (var p in _peers)
	        {
	            p.Poll(); // Must poll to update the state.
	            if (p.GetStatus() == PacketPeerDtls.Status.Connected)
	            {
	                while (p.GetAvailablePacketCount() > 0)
	                {
	                    GD.Print($"Received Message From Client: {p.GetPacket().GetStringFromUtf8()}");
	                    p.PutPacket("Hello DTLS Client".ToUtf8Buffer());
	                }
	            }
	        }
	    }
	}

[/csharp]
[/codeblocks]
[codeblocks]
[gdscript]
# client_node.gd
extends Node

var dtls := PacketPeerDTLS.new()
var udp := PacketPeerUDP.new()
var connected = false

func _ready():

	udp.connect_to_host("127.0.0.1", 4242)
	dtls.connect_to_peer(udp, false) # Use true in production for certificate validation!

func _process(delta):

	dtls.poll()
	if dtls.get_status() == PacketPeerDTLS.STATUS_CONNECTED:
	    if !connected:
	        # Try to contact server
	        dtls.put_packet("The answer is... 42!".to_utf8_buffer())
	    while dtls.get_available_packet_count() > 0:
	        print("Connected: %s" % dtls.get_packet().get_string_from_utf8())
	        connected = true

[/gdscript]
[csharp]
// ClientNode.cs
using Godot;
using System.Text;

public partial class ClientNode : Node

	{
	    private PacketPeerDtls _dtls = new PacketPeerDtls();
	    private PacketPeerUdp _udp = new PacketPeerUdp();
	    private bool _connected = false;

	    public override void _Ready()
	    {
	        _udp.ConnectToHost("127.0.0.1", 4242);
	        _dtls.ConnectToPeer(_udp, validateCerts: false); // Use true in production for certificate validation!
	    }

	    public override void _Process(double delta)
	    {
	        _dtls.Poll();
	        if (_dtls.GetStatus() == PacketPeerDtls.Status.Connected)
	        {
	            if (!_connected)
	            {
	                // Try to contact server
	                _dtls.PutPacket("The Answer Is..42!".ToUtf8Buffer());
	            }
	            while (_dtls.GetAvailablePacketCount() > 0)
	            {
	                GD.Print($"Connected: {_dtls.GetPacket().GetStringFromUtf8()}");
	                _connected = true;
	            }
	        }
	    }
	}

[/csharp]
[/codeblocks]
*/
type Instance [1]gdclass.DTLSServer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsDTLSServer() Instance
}

/*
Setup the DTLS server to use the given [param server_options]. See [method TLSOptions.server].
*/
func (self Instance) Setup(server_options [1]gdclass.TLSOptions) error {
	return error(gd.ToError(class(self).Setup(server_options)))
}

/*
Try to initiate the DTLS handshake with the given [param udp_peer] which must be already connected (see [method PacketPeerUDP.connect_to_host]).
[b]Note:[/b] You must check that the state of the return PacketPeerUDP is [constant PacketPeerDTLS.STATUS_HANDSHAKING], as it is normal that 50% of the new connections will be invalid due to cookie exchange.
*/
func (self Instance) TakeConnection(udp_peer [1]gdclass.PacketPeerUDP) [1]gdclass.PacketPeerDTLS {
	return [1]gdclass.PacketPeerDTLS(class(self).TakeConnection(udp_peer))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.DTLSServer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("DTLSServer"))
	casted := Instance{*(*gdclass.DTLSServer)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Setup the DTLS server to use the given [param server_options]. See [method TLSOptions.server].
*/
//go:nosplit
func (self class) Setup(server_options [1]gdclass.TLSOptions) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(server_options[0])[0])
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DTLSServer.Bind_setup, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Try to initiate the DTLS handshake with the given [param udp_peer] which must be already connected (see [method PacketPeerUDP.connect_to_host]).
[b]Note:[/b] You must check that the state of the return PacketPeerUDP is [constant PacketPeerDTLS.STATUS_HANDSHAKING], as it is normal that 50% of the new connections will be invalid due to cookie exchange.
*/
//go:nosplit
func (self class) TakeConnection(udp_peer [1]gdclass.PacketPeerUDP) [1]gdclass.PacketPeerDTLS {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(udp_peer[0])[0])
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DTLSServer.Bind_take_connection, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PacketPeerDTLS{gd.PointerWithOwnershipTransferredToGo[gdclass.PacketPeerDTLS](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsDTLSServer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsDTLSServer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("DTLSServer", func(ptr gd.Object) any { return [1]gdclass.DTLSServer{*(*gdclass.DTLSServer)(unsafe.Pointer(&ptr))} })
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
