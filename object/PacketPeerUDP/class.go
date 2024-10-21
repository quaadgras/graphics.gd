package PacketPeerUDP

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/PacketPeer"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
UDP packet peer. Can be used to send raw UDP packets as well as [Variant]s.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.

*/
type Simple [1]classdb.PacketPeerUDP
func (self Simple) Bind(port int, bind_address string, recv_buf_size int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Bind(gd.Int(port), gc.String(bind_address), gd.Int(recv_buf_size)))
}
func (self Simple) Close() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Close()
}
func (self Simple) Wait() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Wait())
}
func (self Simple) IsBound() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsBound())
}
func (self Simple) ConnectToHost(host string, port int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).ConnectToHost(gc.String(host), gd.Int(port)))
}
func (self Simple) IsSocketConnected() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSocketConnected())
}
func (self Simple) GetPacketIp() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetPacketIp(gc).String())
}
func (self Simple) GetPacketPort() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPacketPort()))
}
func (self Simple) GetLocalPort() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetLocalPort()))
}
func (self Simple) SetDestAddress(host string, port int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).SetDestAddress(gc.String(host), gd.Int(port)))
}
func (self Simple) SetBroadcastEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBroadcastEnabled(enabled)
}
func (self Simple) JoinMulticastGroup(multicast_address string, interface_name string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).JoinMulticastGroup(gc.String(multicast_address), gc.String(interface_name)))
}
func (self Simple) LeaveMulticastGroup(multicast_address string, interface_name string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).LeaveMulticastGroup(gc.String(multicast_address), gc.String(interface_name)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PacketPeerUDP
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Binds this [PacketPeerUDP] to the specified [param port] and [param bind_address] with a buffer size [param recv_buf_size], allowing it to receive incoming packets.
If [param bind_address] is set to [code]"*"[/code] (default), the peer will be bound on all available addresses (both IPv4 and IPv6).
If [param bind_address] is set to [code]"0.0.0.0"[/code] (for IPv4) or [code]"::"[/code] (for IPv6), the peer will be bound to all available addresses matching that IP type.
If [param bind_address] is set to any valid address (e.g. [code]"192.168.1.101"[/code], [code]"::1"[/code], etc.), the peer will only be bound to the interface with that address (or fail if no interface with the given address exists).
*/
//go:nosplit
func (self class) Bind(port gd.Int, bind_address gd.String, recv_buf_size gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, mmm.Get(bind_address))
	callframe.Arg(frame, recv_buf_size)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerUDP.Bind_bind, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Closes the [PacketPeerUDP]'s underlying UDP socket.
*/
//go:nosplit
func (self class) Close()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerUDP.Bind_close, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) Wait() int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerUDP.Bind_wait, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether this [PacketPeerUDP] is bound to an address and can receive packets.
*/
//go:nosplit
func (self class) IsBound() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerUDP.Bind_is_bound, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Calling this method connects this UDP peer to the given [param host]/[param port] pair. UDP is in reality connectionless, so this option only means that incoming packets from different addresses are automatically discarded, and that outgoing packets are always sent to the connected address (future calls to [method set_dest_address] are not allowed). This method does not send any data to the remote peer, to do that, use [method PacketPeer.put_var] or [method PacketPeer.put_packet] as usual. See also [UDPServer].
[b]Note:[/b] Connecting to the remote peer does not help to protect from malicious attacks like IP spoofing, etc. Think about using an encryption technique like TLS or DTLS if you feel like your application is transferring sensitive information.
*/
//go:nosplit
func (self class) ConnectToHost(host gd.String, port gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(host))
	callframe.Arg(frame, port)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerUDP.Bind_connect_to_host, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the UDP socket is open and has been connected to a remote address. See [method connect_to_host].
*/
//go:nosplit
func (self class) IsSocketConnected() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerUDP.Bind_is_socket_connected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the IP of the remote peer that sent the last packet(that was received with [method PacketPeer.get_packet] or [method PacketPeer.get_var]).
*/
//go:nosplit
func (self class) GetPacketIp(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerUDP.Bind_get_packet_ip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the port of the remote peer that sent the last packet(that was received with [method PacketPeer.get_packet] or [method PacketPeer.get_var]).
*/
//go:nosplit
func (self class) GetPacketPort() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerUDP.Bind_get_packet_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the local port to which this peer is bound.
*/
//go:nosplit
func (self class) GetLocalPort() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerUDP.Bind_get_local_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the destination address and port for sending packets and variables. A hostname will be resolved using DNS if needed.
[b]Note:[/b] [method set_broadcast_enabled] must be enabled before sending packets to a broadcast address (e.g. [code]255.255.255.255[/code]).
*/
//go:nosplit
func (self class) SetDestAddress(host gd.String, port gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(host))
	callframe.Arg(frame, port)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerUDP.Bind_set_dest_address, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Enable or disable sending of broadcast packets (e.g. [code]set_dest_address("255.255.255.255", 4343)[/code]. This option is disabled by default.
[b]Note:[/b] Some Android devices might require the [code]CHANGE_WIFI_MULTICAST_STATE[/code] permission and this option to be enabled to receive broadcast packets too.
*/
//go:nosplit
func (self class) SetBroadcastEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerUDP.Bind_set_broadcast_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Joins the multicast group specified by [param multicast_address] using the interface identified by [param interface_name].
You can join the same multicast group with multiple interfaces. Use [method IP.get_local_interfaces] to know which are available.
[b]Note:[/b] Some Android devices might require the [code]CHANGE_WIFI_MULTICAST_STATE[/code] permission for multicast to work.
*/
//go:nosplit
func (self class) JoinMulticastGroup(multicast_address gd.String, interface_name gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(multicast_address))
	callframe.Arg(frame, mmm.Get(interface_name))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerUDP.Bind_join_multicast_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the interface identified by [param interface_name] from the multicast group specified by [param multicast_address].
*/
//go:nosplit
func (self class) LeaveMulticastGroup(multicast_address gd.String, interface_name gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(multicast_address))
	callframe.Arg(frame, mmm.Get(interface_name))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerUDP.Bind_leave_multicast_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsPacketPeerUDP() Expert { return self[0].AsPacketPeerUDP() }


//go:nosplit
func (self Simple) AsPacketPeerUDP() Simple { return self[0].AsPacketPeerUDP() }


//go:nosplit
func (self class) AsPacketPeer() PacketPeer.Expert { return self[0].AsPacketPeer() }


//go:nosplit
func (self Simple) AsPacketPeer() PacketPeer.Simple { return self[0].AsPacketPeer() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("PacketPeerUDP", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
