// Package ENetConnection provides methods for working with ENetConnection object instances.
package ENetConnection

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
ENet's purpose is to provide a relatively thin, simple and robust network communication layer on top of UDP (User Datagram Protocol).
*/
type Instance [1]gdclass.ENetConnection
type Any interface {
	gd.IsClass
	AsENetConnection() Instance
}

/*
Creates an ENetHost bound to the given [param bind_address] and [param bind_port] that allows up to [param max_peers] connected peers, each allocating up to [param max_channels] channels, optionally limiting bandwidth to [param in_bandwidth] and [param out_bandwidth] (if greater than zero).
[b]Note:[/b] It is necessary to create a host in both client and server in order to establish a connection.
*/
func (self Instance) CreateHostBound(bind_address string, bind_port int) error {
	return error(class(self).CreateHostBound(gd.NewString(bind_address), gd.Int(bind_port), gd.Int(32), gd.Int(0), gd.Int(0), gd.Int(0)))
}

/*
Creates an ENetHost that allows up to [param max_peers] connected peers, each allocating up to [param max_channels] channels, optionally limiting bandwidth to [param in_bandwidth] and [param out_bandwidth] (if greater than zero).
This method binds a random available dynamic UDP port on the host machine at the [i]unspecified[/i] address. Use [method create_host_bound] to specify the address and port.
[b]Note:[/b] It is necessary to create a host in both client and server in order to establish a connection.
*/
func (self Instance) CreateHost() error {
	return error(class(self).CreateHost(gd.Int(32), gd.Int(0), gd.Int(0), gd.Int(0)))
}

/*
Destroys the host and all resources associated with it.
*/
func (self Instance) Destroy() {
	class(self).Destroy()
}

/*
Initiates a connection to a foreign [param address] using the specified [param port] and allocating the requested [param channels]. Optional [param data] can be passed during connection in the form of a 32 bit integer.
[b]Note:[/b] You must call either [method create_host] or [method create_host_bound] on both ends before calling this method.
*/
func (self Instance) ConnectToHost(address string, port int) [1]gdclass.ENetPacketPeer {
	return [1]gdclass.ENetPacketPeer(class(self).ConnectToHost(gd.NewString(address), gd.Int(port), gd.Int(0), gd.Int(0)))
}

/*
Waits for events on this connection and shuttles packets between the host and its peers, with the given [param timeout] (in milliseconds). The returned [Array] will have 4 elements. An [enum EventType], the [ENetPacketPeer] which generated the event, the event associated data (if any), the event associated channel (if any). If the generated event is [constant EVENT_RECEIVE], the received packet will be queued to the associated [ENetPacketPeer].
Call this function regularly to handle connections, disconnections, and to receive new packets.
[b]Note:[/b] This method must be called on both ends involved in the event (sending and receiving hosts).
*/
func (self Instance) Service() Array.Any {
	return Array.Any(class(self).Service(gd.Int(0)))
}

/*
Sends any queued packets on the host specified to its designated peers.
*/
func (self Instance) Flush() {
	class(self).Flush()
}

/*
Adjusts the bandwidth limits of a host.
*/
func (self Instance) BandwidthLimit() {
	class(self).BandwidthLimit(gd.Int(0), gd.Int(0))
}

/*
Limits the maximum allowed channels of future incoming connections.
*/
func (self Instance) ChannelLimit(limit int) {
	class(self).ChannelLimit(gd.Int(limit))
}

/*
Queues a [param packet] to be sent to all peers associated with the host over the specified [param channel]. See [ENetPacketPeer] [code]FLAG_*[/code] constants for available packet flags.
*/
func (self Instance) Broadcast(channel int, packet []byte, flags int) {
	class(self).Broadcast(gd.Int(channel), gd.NewPackedByteSlice(packet), gd.Int(flags))
}

/*
Sets the compression method used for network packets. These have different tradeoffs of compression speed versus bandwidth, you may need to test which one works best for your use case if you use compression at all.
[b]Note:[/b] Most games' network design involve sending many small packets frequently (smaller than 4 KB each). If in doubt, it is recommended to keep the default compression algorithm as it works best on these small packets.
[b]Note:[/b] The compression mode must be set to the same value on both the server and all its clients. Clients will fail to connect if the compression mode set on the client differs from the one set on the server.
*/
func (self Instance) Compress(mode gdclass.ENetConnectionCompressionMode) {
	class(self).Compress(mode)
}

/*
Configure this ENetHost to use the custom Godot extension allowing DTLS encryption for ENet servers. Call this right after [method create_host_bound] to have ENet expect peers to connect using DTLS. See [method TLSOptions.server].
*/
func (self Instance) DtlsServerSetup(server_options [1]gdclass.TLSOptions) error {
	return error(class(self).DtlsServerSetup(server_options))
}

/*
Configure this ENetHost to use the custom Godot extension allowing DTLS encryption for ENet clients. Call this before [method connect_to_host] to have ENet connect using DTLS validating the server certificate against [param hostname]. You can pass the optional [param client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
*/
func (self Instance) DtlsClientSetup(hostname string) error {
	return error(class(self).DtlsClientSetup(gd.NewString(hostname), [1][1]gdclass.TLSOptions{}[0]))
}

/*
Configures the DTLS server to automatically drop new connections.
[b]Note:[/b] This method is only relevant after calling [method dtls_server_setup].
*/
func (self Instance) RefuseNewConnections(refuse bool) {
	class(self).RefuseNewConnections(refuse)
}

/*
Returns and resets host statistics. See [enum HostStatistic] for more info.
*/
func (self Instance) PopStatistic(statistic gdclass.ENetConnectionHostStatistic) Float.X {
	return Float.X(Float.X(class(self).PopStatistic(statistic)))
}

/*
Returns the maximum number of channels allowed for connected peers.
*/
func (self Instance) GetMaxChannels() int {
	return int(int(class(self).GetMaxChannels()))
}

/*
Returns the local port to which this peer is bound.
*/
func (self Instance) GetLocalPort() int {
	return int(int(class(self).GetLocalPort()))
}

/*
Returns the list of peers associated with this host.
[b]Note:[/b] This list might include some peers that are not fully connected or are still being disconnected.
*/
func (self Instance) GetPeers() gd.Array {
	return gd.Array(class(self).GetPeers())
}

/*
Sends a [param packet] toward a destination from the address and port currently bound by this ENetConnection instance.
This is useful as it serves to establish entries in NAT routing tables on all devices between this bound instance and the public facing internet, allowing a prospective client's connection packets to be routed backward through the NAT device(s) between the public internet and this host.
This requires forward knowledge of a prospective client's address and communication port as seen by the public internet - after any NAT devices have handled their connection request. This information can be obtained by a [url=https://en.wikipedia.org/wiki/STUN]STUN[/url] service, and must be handed off to your host by an entity that is not the prospective client. This will never work for a client behind a Symmetric NAT due to the nature of the Symmetric NAT routing algorithm, as their IP and Port cannot be known beforehand.
*/
func (self Instance) SocketSend(destination_address string, destination_port int, packet []byte) {
	class(self).SocketSend(gd.NewString(destination_address), gd.Int(destination_port), gd.NewPackedByteSlice(packet))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ENetConnection

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ENetConnection"))
	return Instance{*(*gdclass.ENetConnection)(unsafe.Pointer(&object))}
}

/*
Creates an ENetHost bound to the given [param bind_address] and [param bind_port] that allows up to [param max_peers] connected peers, each allocating up to [param max_channels] channels, optionally limiting bandwidth to [param in_bandwidth] and [param out_bandwidth] (if greater than zero).
[b]Note:[/b] It is necessary to create a host in both client and server in order to establish a connection.
*/
//go:nosplit
func (self class) CreateHostBound(bind_address gd.String, bind_port gd.Int, max_peers gd.Int, max_channels gd.Int, in_bandwidth gd.Int, out_bandwidth gd.Int) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(bind_address))
	callframe.Arg(frame, bind_port)
	callframe.Arg(frame, max_peers)
	callframe.Arg(frame, max_channels)
	callframe.Arg(frame, in_bandwidth)
	callframe.Arg(frame, out_bandwidth)
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetConnection.Bind_create_host_bound, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) CreateHost(max_peers gd.Int, max_channels gd.Int, in_bandwidth gd.Int, out_bandwidth gd.Int) error {
	var frame = callframe.New()
	callframe.Arg(frame, max_peers)
	callframe.Arg(frame, max_channels)
	callframe.Arg(frame, in_bandwidth)
	callframe.Arg(frame, out_bandwidth)
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetConnection.Bind_create_host, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Destroys the host and all resources associated with it.
*/
//go:nosplit
func (self class) Destroy() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetConnection.Bind_destroy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Initiates a connection to a foreign [param address] using the specified [param port] and allocating the requested [param channels]. Optional [param data] can be passed during connection in the form of a 32 bit integer.
[b]Note:[/b] You must call either [method create_host] or [method create_host_bound] on both ends before calling this method.
*/
//go:nosplit
func (self class) ConnectToHost(address gd.String, port gd.Int, channels gd.Int, data gd.Int) [1]gdclass.ENetPacketPeer {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(address))
	callframe.Arg(frame, port)
	callframe.Arg(frame, channels)
	callframe.Arg(frame, data)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetConnection.Bind_connect_to_host, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.ENetPacketPeer{gd.PointerWithOwnershipTransferredToGo[gdclass.ENetPacketPeer](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Waits for events on this connection and shuttles packets between the host and its peers, with the given [param timeout] (in milliseconds). The returned [Array] will have 4 elements. An [enum EventType], the [ENetPacketPeer] which generated the event, the event associated data (if any), the event associated channel (if any). If the generated event is [constant EVENT_RECEIVE], the received packet will be queued to the associated [ENetPacketPeer].
Call this function regularly to handle connections, disconnections, and to receive new packets.
[b]Note:[/b] This method must be called on both ends involved in the event (sending and receiving hosts).
*/
//go:nosplit
func (self class) Service(timeout gd.Int) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, timeout)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetConnection.Bind_service, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sends any queued packets on the host specified to its designated peers.
*/
//go:nosplit
func (self class) Flush() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetConnection.Bind_flush, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adjusts the bandwidth limits of a host.
*/
//go:nosplit
func (self class) BandwidthLimit(in_bandwidth gd.Int, out_bandwidth gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, in_bandwidth)
	callframe.Arg(frame, out_bandwidth)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetConnection.Bind_bandwidth_limit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Limits the maximum allowed channels of future incoming connections.
*/
//go:nosplit
func (self class) ChannelLimit(limit gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, limit)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetConnection.Bind_channel_limit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Queues a [param packet] to be sent to all peers associated with the host over the specified [param channel]. See [ENetPacketPeer] [code]FLAG_*[/code] constants for available packet flags.
*/
//go:nosplit
func (self class) Broadcast(channel gd.Int, packet gd.PackedByteArray, flags gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, channel)
	callframe.Arg(frame, pointers.Get(packet))
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetConnection.Bind_broadcast, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the compression method used for network packets. These have different tradeoffs of compression speed versus bandwidth, you may need to test which one works best for your use case if you use compression at all.
[b]Note:[/b] Most games' network design involve sending many small packets frequently (smaller than 4 KB each). If in doubt, it is recommended to keep the default compression algorithm as it works best on these small packets.
[b]Note:[/b] The compression mode must be set to the same value on both the server and all its clients. Clients will fail to connect if the compression mode set on the client differs from the one set on the server.
*/
//go:nosplit
func (self class) Compress(mode gdclass.ENetConnectionCompressionMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetConnection.Bind_compress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Configure this ENetHost to use the custom Godot extension allowing DTLS encryption for ENet servers. Call this right after [method create_host_bound] to have ENet expect peers to connect using DTLS. See [method TLSOptions.server].
*/
//go:nosplit
func (self class) DtlsServerSetup(server_options [1]gdclass.TLSOptions) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(server_options[0])[0])
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetConnection.Bind_dtls_server_setup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Configure this ENetHost to use the custom Godot extension allowing DTLS encryption for ENet clients. Call this before [method connect_to_host] to have ENet connect using DTLS validating the server certificate against [param hostname]. You can pass the optional [param client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
*/
//go:nosplit
func (self class) DtlsClientSetup(hostname gd.String, client_options [1]gdclass.TLSOptions) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(hostname))
	callframe.Arg(frame, pointers.Get(client_options[0])[0])
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetConnection.Bind_dtls_client_setup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Configures the DTLS server to automatically drop new connections.
[b]Note:[/b] This method is only relevant after calling [method dtls_server_setup].
*/
//go:nosplit
func (self class) RefuseNewConnections(refuse bool) {
	var frame = callframe.New()
	callframe.Arg(frame, refuse)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetConnection.Bind_refuse_new_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns and resets host statistics. See [enum HostStatistic] for more info.
*/
//go:nosplit
func (self class) PopStatistic(statistic gdclass.ENetConnectionHostStatistic) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, statistic)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetConnection.Bind_pop_statistic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the maximum number of channels allowed for connected peers.
*/
//go:nosplit
func (self class) GetMaxChannels() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetConnection.Bind_get_max_channels, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the local port to which this peer is bound.
*/
//go:nosplit
func (self class) GetLocalPort() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetConnection.Bind_get_local_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the list of peers associated with this host.
[b]Note:[/b] This list might include some peers that are not fully connected or are still being disconnected.
*/
//go:nosplit
func (self class) GetPeers() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetConnection.Bind_get_peers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sends a [param packet] toward a destination from the address and port currently bound by this ENetConnection instance.
This is useful as it serves to establish entries in NAT routing tables on all devices between this bound instance and the public facing internet, allowing a prospective client's connection packets to be routed backward through the NAT device(s) between the public internet and this host.
This requires forward knowledge of a prospective client's address and communication port as seen by the public internet - after any NAT devices have handled their connection request. This information can be obtained by a [url=https://en.wikipedia.org/wiki/STUN]STUN[/url] service, and must be handed off to your host by an entity that is not the prospective client. This will never work for a client behind a Symmetric NAT due to the nature of the Symmetric NAT routing algorithm, as their IP and Port cannot be known beforehand.
*/
//go:nosplit
func (self class) SocketSend(destination_address gd.String, destination_port gd.Int, packet gd.PackedByteArray) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(destination_address))
	callframe.Arg(frame, destination_port)
	callframe.Arg(frame, pointers.Get(packet))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetConnection.Bind_socket_send, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsENetConnection() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsENetConnection() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("ENetConnection", func(ptr gd.Object) any {
		return [1]gdclass.ENetConnection{*(*gdclass.ENetConnection)(unsafe.Pointer(&ptr))}
	})
}

type CompressionMode = gdclass.ENetConnectionCompressionMode

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

type EventType = gdclass.ENetConnectionEventType

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

type HostStatistic = gdclass.ENetConnectionHostStatistic

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
