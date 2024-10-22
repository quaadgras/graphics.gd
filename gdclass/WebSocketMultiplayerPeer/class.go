package WebSocketMultiplayerPeer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/MultiplayerPeer"
import "grow.graphics/gd/gdclass/PacketPeer"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Base class for WebSocket server and client, allowing them to be used as multiplayer peer for the [MultiplayerAPI].
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.

*/
type Go [1]classdb.WebSocketMultiplayerPeer

/*
Starts a new multiplayer client connecting to the given [param url]. TLS certificates will be verified against the hostname when connecting using the [code]wss://[/code] protocol. You can pass the optional [param tls_client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
[b]Note:[/b] It is recommended to specify the scheme part of the URL, i.e. the [param url] should start with either [code]ws://[/code] or [code]wss://[/code].
*/
func (self Go) CreateClient(url string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).CreateClient(gc.String(url), ([1]gdclass.TLSOptions{}[0])))
}

/*
Starts a new multiplayer server listening on the given [param port]. You can optionally specify a [param bind_address], and provide valid [param tls_server_options] to use TLS. See [method TLSOptions.server].
*/
func (self Go) CreateServer(port int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).CreateServer(gd.Int(port), gc.String("*"), ([1]gdclass.TLSOptions{}[0])))
}

/*
Returns the [WebSocketPeer] associated to the given [param peer_id].
*/
func (self Go) GetPeer(peer_id int) gdclass.WebSocketPeer {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.WebSocketPeer(class(self).GetPeer(gc, gd.Int(peer_id)))
}

/*
Returns the IP address of the given peer.
*/
func (self Go) GetPeerAddress(id int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetPeerAddress(gc, gd.Int(id)).String())
}

/*
Returns the remote port of the given peer.
*/
func (self Go) GetPeerPort(id int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetPeerPort(gd.Int(id))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.WebSocketMultiplayerPeer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("WebSocketMultiplayerPeer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) SupportedProtocols() []string {
	gc := gd.GarbageCollector(); _ = gc
		return []string(class(self).GetSupportedProtocols(gc).Strings(gc))
}

func (self Go) SetSupportedProtocols(value []string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSupportedProtocols(gc.PackedStringSlice(value))
}

func (self Go) HandshakeHeaders() []string {
	gc := gd.GarbageCollector(); _ = gc
		return []string(class(self).GetHandshakeHeaders(gc).Strings(gc))
}

func (self Go) SetHandshakeHeaders(value []string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHandshakeHeaders(gc.PackedStringSlice(value))
}

func (self Go) InboundBufferSize() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetInboundBufferSize()))
}

func (self Go) SetInboundBufferSize(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInboundBufferSize(gd.Int(value))
}

func (self Go) OutboundBufferSize() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetOutboundBufferSize()))
}

func (self Go) SetOutboundBufferSize(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOutboundBufferSize(gd.Int(value))
}

func (self Go) HandshakeTimeout() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetHandshakeTimeout()))
}

func (self Go) SetHandshakeTimeout(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHandshakeTimeout(gd.Float(value))
}

func (self Go) MaxQueuedPackets() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMaxQueuedPackets()))
}

func (self Go) SetMaxQueuedPackets(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxQueuedPackets(gd.Int(value))
}

/*
Starts a new multiplayer client connecting to the given [param url]. TLS certificates will be verified against the hostname when connecting using the [code]wss://[/code] protocol. You can pass the optional [param tls_client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
[b]Note:[/b] It is recommended to specify the scheme part of the URL, i.e. the [param url] should start with either [code]ws://[/code] or [code]wss://[/code].
*/
//go:nosplit
func (self class) CreateClient(url gd.String, tls_client_options gdclass.TLSOptions) int64 {
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
func (self class) CreateServer(port gd.Int, bind_address gd.String, tls_server_options gdclass.TLSOptions) int64 {
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
func (self class) GetPeer(ctx gd.Lifetime, peer_id gd.Int) gdclass.WebSocketPeer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketMultiplayerPeer.Bind_get_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.WebSocketPeer
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
func (self class) AsWebSocketMultiplayerPeer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsWebSocketMultiplayerPeer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsMultiplayerPeer() MultiplayerPeer.GD { return *((*MultiplayerPeer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsMultiplayerPeer() MultiplayerPeer.Go { return *((*MultiplayerPeer.Go)(unsafe.Pointer(&self))) }
func (self class) AsPacketPeer() PacketPeer.GD { return *((*PacketPeer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsPacketPeer() PacketPeer.Go { return *((*PacketPeer.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMultiplayerPeer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMultiplayerPeer(), name)
	}
}
func init() {classdb.Register("WebSocketMultiplayerPeer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
