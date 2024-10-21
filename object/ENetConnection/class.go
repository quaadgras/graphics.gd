package ENetConnection

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
ENet's purpose is to provide a relatively thin, simple and robust network communication layer on top of UDP (User Datagram Protocol).

*/
type Simple [1]classdb.ENetConnection
func (self Simple) CreateHostBound(bind_address string, bind_port int, max_peers int, max_channels int, in_bandwidth int, out_bandwidth int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).CreateHostBound(gc.String(bind_address), gd.Int(bind_port), gd.Int(max_peers), gd.Int(max_channels), gd.Int(in_bandwidth), gd.Int(out_bandwidth)))
}
func (self Simple) CreateHost(max_peers int, max_channels int, in_bandwidth int, out_bandwidth int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).CreateHost(gd.Int(max_peers), gd.Int(max_channels), gd.Int(in_bandwidth), gd.Int(out_bandwidth)))
}
func (self Simple) Destroy() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Destroy()
}
func (self Simple) ConnectToHost(address string, port int, channels int, data int) [1]classdb.ENetPacketPeer {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.ENetPacketPeer(Expert(self).ConnectToHost(gc, gc.String(address), gd.Int(port), gd.Int(channels), gd.Int(data)))
}
func (self Simple) Service(timeout int) gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).Service(gc, gd.Int(timeout)))
}
func (self Simple) Flush() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Flush()
}
func (self Simple) BandwidthLimit(in_bandwidth int, out_bandwidth int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BandwidthLimit(gd.Int(in_bandwidth), gd.Int(out_bandwidth))
}
func (self Simple) ChannelLimit(limit int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ChannelLimit(gd.Int(limit))
}
func (self Simple) Broadcast(channel int, packet []byte, flags int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Broadcast(gd.Int(channel), gc.PackedByteSlice(packet), gd.Int(flags))
}
func (self Simple) Compress(mode classdb.ENetConnectionCompressionMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Compress(mode)
}
func (self Simple) DtlsServerSetup(server_options [1]classdb.TLSOptions) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).DtlsServerSetup(server_options))
}
func (self Simple) DtlsClientSetup(hostname string, client_options [1]classdb.TLSOptions) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).DtlsClientSetup(gc.String(hostname), client_options))
}
func (self Simple) RefuseNewConnections(refuse bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RefuseNewConnections(refuse)
}
func (self Simple) PopStatistic(statistic classdb.ENetConnectionHostStatistic) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).PopStatistic(statistic)))
}
func (self Simple) GetMaxChannels() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMaxChannels()))
}
func (self Simple) GetLocalPort() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetLocalPort()))
}
func (self Simple) GetPeers() gd.ArrayOf[[1]classdb.ENetPacketPeer] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.ENetPacketPeer](Expert(self).GetPeers(gc))
}
func (self Simple) SocketSend(destination_address string, destination_port int, packet []byte) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SocketSend(gc.String(destination_address), gd.Int(destination_port), gc.PackedByteSlice(packet))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ENetConnection
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Creates an ENetHost bound to the given [param bind_address] and [param bind_port] that allows up to [param max_peers] connected peers, each allocating up to [param max_channels] channels, optionally limiting bandwidth to [param in_bandwidth] and [param out_bandwidth] (if greater than zero).
[b]Note:[/b] It is necessary to create a host in both client and server in order to establish a connection.
*/
//go:nosplit
func (self class) CreateHostBound(bind_address gd.String, bind_port gd.Int, max_peers gd.Int, max_channels gd.Int, in_bandwidth gd.Int, out_bandwidth gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bind_address))
	callframe.Arg(frame, bind_port)
	callframe.Arg(frame, max_peers)
	callframe.Arg(frame, max_channels)
	callframe.Arg(frame, in_bandwidth)
	callframe.Arg(frame, out_bandwidth)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetConnection.Bind_create_host_bound, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates an ENetHost that allows up to [param max_peers] connected peers, each allocating up to [param max_channels] channels, optionally limiting bandwidth to [param in_bandwidth] and [param out_bandwidth] (if greater than zero).
This method binds a random available dynamic UDP port on the host machine at the [i]unspecified[/i] address. Use [method create_host_bound] to specify the address and port.
[b]Note:[/b] It is necessary to create a host in both client and server in order to establish a connection.
*/
//go:nosplit
func (self class) CreateHost(max_peers gd.Int, max_channels gd.Int, in_bandwidth gd.Int, out_bandwidth gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_peers)
	callframe.Arg(frame, max_channels)
	callframe.Arg(frame, in_bandwidth)
	callframe.Arg(frame, out_bandwidth)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetConnection.Bind_create_host, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Destroys the host and all resources associated with it.
*/
//go:nosplit
func (self class) Destroy()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetConnection.Bind_destroy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Initiates a connection to a foreign [param address] using the specified [param port] and allocating the requested [param channels]. Optional [param data] can be passed during connection in the form of a 32 bit integer.
[b]Note:[/b] You must call either [method create_host] or [method create_host_bound] on both ends before calling this method.
*/
//go:nosplit
func (self class) ConnectToHost(ctx gd.Lifetime, address gd.String, port gd.Int, channels gd.Int, data gd.Int) object.ENetPacketPeer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(address))
	callframe.Arg(frame, port)
	callframe.Arg(frame, channels)
	callframe.Arg(frame, data)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetConnection.Bind_connect_to_host, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.ENetPacketPeer
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Waits for events on this connection and shuttles packets between the host and its peers, with the given [param timeout] (in milliseconds). The returned [Array] will have 4 elements. An [enum EventType], the [ENetPacketPeer] which generated the event, the event associated data (if any), the event associated channel (if any). If the generated event is [constant EVENT_RECEIVE], the received packet will be queued to the associated [ENetPacketPeer].
Call this function regularly to handle connections, disconnections, and to receive new packets.
[b]Note:[/b] This method must be called on both ends involved in the event (sending and receiving hosts).
*/
//go:nosplit
func (self class) Service(ctx gd.Lifetime, timeout gd.Int) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, timeout)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetConnection.Bind_service, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sends any queued packets on the host specified to its designated peers.
*/
//go:nosplit
func (self class) Flush()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetConnection.Bind_flush, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adjusts the bandwidth limits of a host.
*/
//go:nosplit
func (self class) BandwidthLimit(in_bandwidth gd.Int, out_bandwidth gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, in_bandwidth)
	callframe.Arg(frame, out_bandwidth)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetConnection.Bind_bandwidth_limit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Limits the maximum allowed channels of future incoming connections.
*/
//go:nosplit
func (self class) ChannelLimit(limit gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, limit)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetConnection.Bind_channel_limit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Queues a [param packet] to be sent to all peers associated with the host over the specified [param channel]. See [ENetPacketPeer] [code]FLAG_*[/code] constants for available packet flags.
*/
//go:nosplit
func (self class) Broadcast(channel gd.Int, packet gd.PackedByteArray, flags gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, channel)
	callframe.Arg(frame, mmm.Get(packet))
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetConnection.Bind_broadcast, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the compression method used for network packets. These have different tradeoffs of compression speed versus bandwidth, you may need to test which one works best for your use case if you use compression at all.
[b]Note:[/b] Most games' network design involve sending many small packets frequently (smaller than 4 KB each). If in doubt, it is recommended to keep the default compression algorithm as it works best on these small packets.
[b]Note:[/b] The compression mode must be set to the same value on both the server and all its clients. Clients will fail to connect if the compression mode set on the client differs from the one set on the server.
*/
//go:nosplit
func (self class) Compress(mode classdb.ENetConnectionCompressionMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetConnection.Bind_compress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Configure this ENetHost to use the custom Godot extension allowing DTLS encryption for ENet servers. Call this right after [method create_host_bound] to have ENet expect peers to connect using DTLS. See [method TLSOptions.server].
*/
//go:nosplit
func (self class) DtlsServerSetup(server_options object.TLSOptions) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(server_options[0].AsPointer())[0])
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetConnection.Bind_dtls_server_setup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Configure this ENetHost to use the custom Godot extension allowing DTLS encryption for ENet clients. Call this before [method connect_to_host] to have ENet connect using DTLS validating the server certificate against [param hostname]. You can pass the optional [param client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
*/
//go:nosplit
func (self class) DtlsClientSetup(hostname gd.String, client_options object.TLSOptions) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(hostname))
	callframe.Arg(frame, mmm.Get(client_options[0].AsPointer())[0])
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetConnection.Bind_dtls_client_setup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Configures the DTLS server to automatically drop new connections.
[b]Note:[/b] This method is only relevant after calling [method dtls_server_setup].
*/
//go:nosplit
func (self class) RefuseNewConnections(refuse bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, refuse)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetConnection.Bind_refuse_new_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns and resets host statistics. See [enum HostStatistic] for more info.
*/
//go:nosplit
func (self class) PopStatistic(statistic classdb.ENetConnectionHostStatistic) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, statistic)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetConnection.Bind_pop_statistic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the maximum number of channels allowed for connected peers.
*/
//go:nosplit
func (self class) GetMaxChannels() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetConnection.Bind_get_max_channels, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetConnection.Bind_get_local_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the list of peers associated with this host.
[b]Note:[/b] This list might include some peers that are not fully connected or are still being disconnected.
*/
//go:nosplit
func (self class) GetPeers(ctx gd.Lifetime) gd.ArrayOf[object.ENetPacketPeer] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetConnection.Bind_get_peers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.ENetPacketPeer](ret)
}
/*
Sends a [param packet] toward a destination from the address and port currently bound by this ENetConnection instance.
This is useful as it serves to establish entries in NAT routing tables on all devices between this bound instance and the public facing internet, allowing a prospective client's connection packets to be routed backward through the NAT device(s) between the public internet and this host.
This requires forward knowledge of a prospective client's address and communication port as seen by the public internet - after any NAT devices have handled their connection request. This information can be obtained by a [url=https://en.wikipedia.org/wiki/STUN]STUN[/url] service, and must be handed off to your host by an entity that is not the prospective client. This will never work for a client behind a Symmetric NAT due to the nature of the Symmetric NAT routing algorithm, as their IP and Port cannot be known beforehand.
*/
//go:nosplit
func (self class) SocketSend(destination_address gd.String, destination_port gd.Int, packet gd.PackedByteArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(destination_address))
	callframe.Arg(frame, destination_port)
	callframe.Arg(frame, mmm.Get(packet))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetConnection.Bind_socket_send, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsENetConnection() Expert { return self[0].AsENetConnection() }


//go:nosplit
func (self Simple) AsENetConnection() Simple { return self[0].AsENetConnection() }


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
func init() {classdb.Register("ENetConnection", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type CompressionMode = classdb.ENetConnectionCompressionMode

const (
/*No compression. This uses the most bandwidth, but has the upside of requiring the fewest CPU resources. This option may also be used to make network debugging using tools like Wireshark easier.*/
	CompressNone CompressionMode = 0
/*ENet's built-in range encoding. Works well on small packets, but is not the most efficient algorithm on packets larger than 4 KB.*/
	CompressRangeCoder CompressionMode = 1
/*[url=https://fastlz.org/]FastLZ[/url] compression. This option uses less CPU resources compared to [constant COMPRESS_ZLIB], at the expense of using more bandwidth.*/
	CompressFastlz CompressionMode = 2
/*[url=https://www.zlib.net/]Zlib[/url] compression. This option uses less bandwidth compared to [constant COMPRESS_FASTLZ], at the expense of using more CPU resources.*/
	CompressZlib CompressionMode = 3
/*[url=https://facebook.github.io/zstd/]Zstandard[/url] compression. Note that this algorithm is not very efficient on packets smaller than 4 KB. Therefore, it's recommended to use other compression algorithms in most cases.*/
	CompressZstd CompressionMode = 4
)
type EventType = classdb.ENetConnectionEventType

const (
/*An error occurred during [method service]. You will likely need to [method destroy] the host and recreate it.*/
	EventError EventType = -1
/*No event occurred within the specified time limit.*/
	EventNone EventType = 0
/*A connection request initiated by enet_host_connect has completed. The array will contain the peer which successfully connected.*/
	EventConnect EventType = 1
/*A peer has disconnected. This event is generated on a successful completion of a disconnect initiated by [method ENetPacketPeer.peer_disconnect], if a peer has timed out, or if a connection request initialized by [method connect_to_host] has timed out. The array will contain the peer which disconnected. The data field contains user supplied data describing the disconnection, or 0, if none is available.*/
	EventDisconnect EventType = 2
/*A packet has been received from a peer. The array will contain the peer which sent the packet and the channel number upon which the packet was received. The received packet will be queued to the associated [ENetPacketPeer].*/
	EventReceive EventType = 3
)
type HostStatistic = classdb.ENetConnectionHostStatistic

const (
/*Total data sent.*/
	HostTotalSentData HostStatistic = 0
/*Total UDP packets sent.*/
	HostTotalSentPackets HostStatistic = 1
/*Total data received.*/
	HostTotalReceivedData HostStatistic = 2
/*Total UDP packets received.*/
	HostTotalReceivedPackets HostStatistic = 3
)
