package WebSocketMultiplayerPeer

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
Base class for WebSocket server and client, allowing them to be used as multiplayer peer for the [MultiplayerAPI].
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.

*/
type Simple [1]classdb.WebSocketMultiplayerPeer
func (self Simple) CreateClient(url string, tls_client_options [1]classdb.TLSOptions) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).CreateClient(gc.String(url), tls_client_options))
}
func (self Simple) CreateServer(port int, bind_address string, tls_server_options [1]classdb.TLSOptions) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).CreateServer(gd.Int(port), gc.String(bind_address), tls_server_options))
}
func (self Simple) GetPeer(peer_id int) [1]classdb.WebSocketPeer {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.WebSocketPeer(Expert(self).GetPeer(gc, gd.Int(peer_id)))
}
func (self Simple) GetPeerAddress(id int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetPeerAddress(gc, gd.Int(id)).String())
}
func (self Simple) GetPeerPort(id int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPeerPort(gd.Int(id))))
}
func (self Simple) GetSupportedProtocols() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetSupportedProtocols(gc))
}
func (self Simple) SetSupportedProtocols(protocols gd.PackedStringArray) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSupportedProtocols(protocols)
}
func (self Simple) GetHandshakeHeaders() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetHandshakeHeaders(gc))
}
func (self Simple) SetHandshakeHeaders(protocols gd.PackedStringArray) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHandshakeHeaders(protocols)
}
func (self Simple) GetInboundBufferSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetInboundBufferSize()))
}
func (self Simple) SetInboundBufferSize(buffer_size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInboundBufferSize(gd.Int(buffer_size))
}
func (self Simple) GetOutboundBufferSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetOutboundBufferSize()))
}
func (self Simple) SetOutboundBufferSize(buffer_size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOutboundBufferSize(gd.Int(buffer_size))
}
func (self Simple) GetHandshakeTimeout() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetHandshakeTimeout()))
}
func (self Simple) SetHandshakeTimeout(timeout float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHandshakeTimeout(gd.Float(timeout))
}
func (self Simple) SetMaxQueuedPackets(max_queued_packets int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaxQueuedPackets(gd.Int(max_queued_packets))
}
func (self Simple) GetMaxQueuedPackets() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMaxQueuedPackets()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.WebSocketMultiplayerPeer
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Starts a new multiplayer client connecting to the given [param url]. TLS certificates will be verified against the hostname when connecting using the [code]wss://[/code] protocol. You can pass the optional [param tls_client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
[b]Note:[/b] It is recommended to specify the scheme part of the URL, i.e. the [param url] should start with either [code]ws://[/code] or [code]wss://[/code].
*/
//go:nosplit
func (self class) CreateClient(url gd.String, tls_client_options object.TLSOptions) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(url))
	callframe.Arg(frame, mmm.Get(tls_client_options[0].AsPointer())[0])
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketMultiplayerPeer.Bind_create_client, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Starts a new multiplayer server listening on the given [param port]. You can optionally specify a [param bind_address], and provide valid [param tls_server_options] to use TLS. See [method TLSOptions.server].
*/
//go:nosplit
func (self class) CreateServer(port gd.Int, bind_address gd.String, tls_server_options object.TLSOptions) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, mmm.Get(bind_address))
	callframe.Arg(frame, mmm.Get(tls_server_options[0].AsPointer())[0])
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketMultiplayerPeer.Bind_create_server, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [WebSocketPeer] associated to the given [param peer_id].
*/
//go:nosplit
func (self class) GetPeer(ctx gd.Lifetime, peer_id gd.Int) object.WebSocketPeer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketMultiplayerPeer.Bind_get_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.WebSocketPeer
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the IP address of the given peer.
*/
//go:nosplit
func (self class) GetPeerAddress(ctx gd.Lifetime, id gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketMultiplayerPeer.Bind_get_peer_address, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the remote port of the given peer.
*/
//go:nosplit
func (self class) GetPeerPort(id gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketMultiplayerPeer.Bind_get_peer_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetSupportedProtocols(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketMultiplayerPeer.Bind_get_supported_protocols, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSupportedProtocols(protocols gd.PackedStringArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(protocols))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketMultiplayerPeer.Bind_set_supported_protocols, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHandshakeHeaders(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketMultiplayerPeer.Bind_get_handshake_headers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHandshakeHeaders(protocols gd.PackedStringArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(protocols))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketMultiplayerPeer.Bind_set_handshake_headers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInboundBufferSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketMultiplayerPeer.Bind_get_inbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInboundBufferSize(buffer_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, buffer_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketMultiplayerPeer.Bind_set_inbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOutboundBufferSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketMultiplayerPeer.Bind_get_outbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOutboundBufferSize(buffer_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, buffer_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketMultiplayerPeer.Bind_set_outbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHandshakeTimeout() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketMultiplayerPeer.Bind_get_handshake_timeout, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHandshakeTimeout(timeout gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, timeout)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketMultiplayerPeer.Bind_set_handshake_timeout, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetMaxQueuedPackets(max_queued_packets gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_queued_packets)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketMultiplayerPeer.Bind_set_max_queued_packets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxQueuedPackets() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketMultiplayerPeer.Bind_get_max_queued_packets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsWebSocketMultiplayerPeer() Expert { return self[0].AsWebSocketMultiplayerPeer() }


//go:nosplit
func (self Simple) AsWebSocketMultiplayerPeer() Simple { return self[0].AsWebSocketMultiplayerPeer() }


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
func init() {classdb.Register("WebSocketMultiplayerPeer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
