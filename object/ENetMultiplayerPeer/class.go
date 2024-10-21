package ENetMultiplayerPeer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/MultiplayerPeer"
import "grow.graphics/gd/object/PacketPeer"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A MultiplayerPeer implementation that should be passed to [member MultiplayerAPI.multiplayer_peer] after being initialized as either a client, server, or mesh. Events can then be handled by connecting to [MultiplayerAPI] signals. See [ENetConnection] for more information on the ENet library wrapper.
[b]Note:[/b] ENet only uses UDP, not TCP. When forwarding the server port to make your server accessible on the public Internet, you only need to forward the server port in UDP. You can use the [UPNP] class to try to forward the server port automatically when starting the server.

*/
type Simple [1]classdb.ENetMultiplayerPeer
func (self Simple) CreateServer(port int, max_clients int, max_channels int, in_bandwidth int, out_bandwidth int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).CreateServer(gd.Int(port), gd.Int(max_clients), gd.Int(max_channels), gd.Int(in_bandwidth), gd.Int(out_bandwidth)))
}
func (self Simple) CreateClient(address string, port int, channel_count int, in_bandwidth int, out_bandwidth int, local_port int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).CreateClient(gc.String(address), gd.Int(port), gd.Int(channel_count), gd.Int(in_bandwidth), gd.Int(out_bandwidth), gd.Int(local_port)))
}
func (self Simple) CreateMesh(unique_id int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).CreateMesh(gd.Int(unique_id)))
}
func (self Simple) AddMeshPeer(peer_id int, host [1]classdb.ENetConnection) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).AddMeshPeer(gd.Int(peer_id), host))
}
func (self Simple) SetBindIp(ip string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBindIp(gc.String(ip))
}
func (self Simple) GetHost() [1]classdb.ENetConnection {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.ENetConnection(Expert(self).GetHost(gc))
}
func (self Simple) GetPeer(id int) [1]classdb.ENetPacketPeer {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.ENetPacketPeer(Expert(self).GetPeer(gc, gd.Int(id)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ENetMultiplayerPeer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Create server that listens to connections via [param port]. The port needs to be an available, unused port between 0 and 65535. Note that ports below 1024 are privileged and may require elevated permissions depending on the platform. To change the interface the server listens on, use [method set_bind_ip]. The default IP is the wildcard [code]"*"[/code], which listens on all available interfaces. [param max_clients] is the maximum number of clients that are allowed at once, any number up to 4095 may be used, although the achievable number of simultaneous clients may be far lower and depends on the application. For additional details on the bandwidth parameters, see [method create_client]. Returns [constant OK] if a server was created, [constant ERR_ALREADY_IN_USE] if this ENetMultiplayerPeer instance already has an open connection (in which case you need to call [method MultiplayerPeer.close] first) or [constant ERR_CANT_CREATE] if the server could not be created.
*/
//go:nosplit
func (self class) CreateServer(port gd.Int, max_clients gd.Int, max_channels gd.Int, in_bandwidth gd.Int, out_bandwidth gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, max_clients)
	callframe.Arg(frame, max_channels)
	callframe.Arg(frame, in_bandwidth)
	callframe.Arg(frame, out_bandwidth)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetMultiplayerPeer.Bind_create_server, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Create client that connects to a server at [param address] using specified [param port]. The given address needs to be either a fully qualified domain name (e.g. [code]"www.example.com"[/code]) or an IP address in IPv4 or IPv6 format (e.g. [code]"192.168.1.1"[/code]). The [param port] is the port the server is listening on. The [param channel_count] parameter can be used to specify the number of ENet channels allocated for the connection. The [param in_bandwidth] and [param out_bandwidth] parameters can be used to limit the incoming and outgoing bandwidth to the given number of bytes per second. The default of 0 means unlimited bandwidth. Note that ENet will strategically drop packets on specific sides of a connection between peers to ensure the peer's bandwidth is not overwhelmed. The bandwidth parameters also determine the window size of a connection which limits the amount of reliable packets that may be in transit at any given time. Returns [constant OK] if a client was created, [constant ERR_ALREADY_IN_USE] if this ENetMultiplayerPeer instance already has an open connection (in which case you need to call [method MultiplayerPeer.close] first) or [constant ERR_CANT_CREATE] if the client could not be created. If [param local_port] is specified, the client will also listen to the given port; this is useful for some NAT traversal techniques.
*/
//go:nosplit
func (self class) CreateClient(address gd.String, port gd.Int, channel_count gd.Int, in_bandwidth gd.Int, out_bandwidth gd.Int, local_port gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(address))
	callframe.Arg(frame, port)
	callframe.Arg(frame, channel_count)
	callframe.Arg(frame, in_bandwidth)
	callframe.Arg(frame, out_bandwidth)
	callframe.Arg(frame, local_port)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetMultiplayerPeer.Bind_create_client, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Initialize this [MultiplayerPeer] in mesh mode. The provided [param unique_id] will be used as the local peer network unique ID once assigned as the [member MultiplayerAPI.multiplayer_peer]. In the mesh configuration you will need to set up each new peer manually using [ENetConnection] before calling [method add_mesh_peer]. While this technique is more advanced, it allows for better control over the connection process (e.g. when dealing with NAT punch-through) and for better distribution of the network load (which would otherwise be more taxing on the server).
*/
//go:nosplit
func (self class) CreateMesh(unique_id gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, unique_id)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetMultiplayerPeer.Bind_create_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Add a new remote peer with the given [param peer_id] connected to the given [param host].
[b]Note:[/b] The [param host] must have exactly one peer in the [constant ENetPacketPeer.STATE_CONNECTED] state.
*/
//go:nosplit
func (self class) AddMeshPeer(peer_id gd.Int, host object.ENetConnection) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	callframe.Arg(frame, mmm.Get(host[0].AsPointer())[0])
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetMultiplayerPeer.Bind_add_mesh_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
The IP used when creating a server. This is set to the wildcard [code]"*"[/code] by default, which binds to all available interfaces. The given IP needs to be in IPv4 or IPv6 address format, for example: [code]"192.168.1.1"[/code].
*/
//go:nosplit
func (self class) SetBindIp(ip gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(ip))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetMultiplayerPeer.Bind_set_bind_ip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHost(ctx gd.Lifetime) object.ENetConnection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetMultiplayerPeer.Bind_get_host, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.ENetConnection
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [ENetPacketPeer] associated to the given [param id].
*/
//go:nosplit
func (self class) GetPeer(ctx gd.Lifetime, id gd.Int) object.ENetPacketPeer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetMultiplayerPeer.Bind_get_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.ENetPacketPeer
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsENetMultiplayerPeer() Expert { return self[0].AsENetMultiplayerPeer() }


//go:nosplit
func (self Simple) AsENetMultiplayerPeer() Simple { return self[0].AsENetMultiplayerPeer() }


//go:nosplit
func (self class) AsMultiplayerPeer() MultiplayerPeer.Expert { return self[0].AsMultiplayerPeer() }


//go:nosplit
func (self Simple) AsMultiplayerPeer() MultiplayerPeer.Simple { return self[0].AsMultiplayerPeer() }


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
func init() {classdb.Register("ENetMultiplayerPeer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
