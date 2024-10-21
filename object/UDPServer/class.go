package UDPServer

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Simple [1]classdb.UDPServer
func (self Simple) Listen(port int, bind_address string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Listen(gd.Int(port), gc.String(bind_address)))
}
func (self Simple) Poll() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Poll())
}
func (self Simple) IsConnectionAvailable() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsConnectionAvailable())
}
func (self Simple) GetLocalPort() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetLocalPort()))
}
func (self Simple) IsListening() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsListening())
}
func (self Simple) TakeConnection() [1]classdb.PacketPeerUDP {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PacketPeerUDP(Expert(self).TakeConnection(gc))
}
func (self Simple) Stop() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Stop()
}
func (self Simple) SetMaxPendingConnections(max_pending_connections int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaxPendingConnections(gd.Int(max_pending_connections))
}
func (self Simple) GetMaxPendingConnections() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMaxPendingConnections()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.UDPServer
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Starts the server by opening a UDP socket listening on the given [param port]. You can optionally specify a [param bind_address] to only listen for packets sent to that address. See also [method PacketPeerUDP.bind].
*/
//go:nosplit
func (self class) Listen(port gd.Int, bind_address gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, mmm.Get(bind_address))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UDPServer.Bind_listen, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Call this method at regular intervals (e.g. inside [method Node._process]) to process new packets. And packet from known address/port pair will be delivered to the appropriate [PacketPeerUDP], any packet received from an unknown address/port pair will be added as a pending connection (see [method is_connection_available], [method take_connection]). The maximum number of pending connection is defined via [member max_pending_connections].
*/
//go:nosplit
func (self class) Poll() int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UDPServer.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if a packet with a new address/port combination was received on the socket.
*/
//go:nosplit
func (self class) IsConnectionAvailable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UDPServer.Bind_is_connection_available, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the local port this server is listening to.
*/
//go:nosplit
func (self class) GetLocalPort() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UDPServer.Bind_get_local_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the socket is open and listening on a port.
*/
//go:nosplit
func (self class) IsListening() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UDPServer.Bind_is_listening, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the first pending connection (connected to the appropriate address/port). Will return [code]null[/code] if no new connection is available. See also [method is_connection_available], [method PacketPeerUDP.connect_to_host].
*/
//go:nosplit
func (self class) TakeConnection(ctx gd.Lifetime) object.PacketPeerUDP {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UDPServer.Bind_take_connection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.PacketPeerUDP
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Stops the server, closing the UDP socket if open. Will close all connected [PacketPeerUDP] accepted via [method take_connection] (remote peers will not be notified).
*/
//go:nosplit
func (self class) Stop()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UDPServer.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetMaxPendingConnections(max_pending_connections gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_pending_connections)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UDPServer.Bind_set_max_pending_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxPendingConnections() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UDPServer.Bind_get_max_pending_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsUDPServer() Expert { return self[0].AsUDPServer() }


//go:nosplit
func (self Simple) AsUDPServer() Simple { return self[0].AsUDPServer() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("UDPServer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
