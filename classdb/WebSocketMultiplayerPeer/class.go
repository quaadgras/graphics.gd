// Package WebSocketMultiplayerPeer provides methods for working with WebSocketMultiplayerPeer object instances.
package WebSocketMultiplayerPeer

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/MultiplayerPeer"
import "graphics.gd/classdb/PacketPeer"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

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
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
Base class for WebSocket server and client, allowing them to be used as multiplayer peer for the [MultiplayerAPI].
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
*/
type Instance [1]gdclass.WebSocketMultiplayerPeer
type Expanded [1]gdclass.WebSocketMultiplayerPeer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsWebSocketMultiplayerPeer() Instance
}

/*
Starts a new multiplayer client connecting to the given [param url]. TLS certificates will be verified against the hostname when connecting using the [code]wss://[/code] protocol. You can pass the optional [param tls_client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
[b]Note:[/b] It is recommended to specify the scheme part of the URL, i.e. the [param url] should start with either [code]ws://[/code] or [code]wss://[/code].
*/
func (self Instance) CreateClient(url string) error { //gd:WebSocketMultiplayerPeer.create_client
	return error(gd.ToError(Advanced(self).CreateClient(String.New(url), [1][1]gdclass.TLSOptions{}[0])))
}

/*
Starts a new multiplayer client connecting to the given [param url]. TLS certificates will be verified against the hostname when connecting using the [code]wss://[/code] protocol. You can pass the optional [param tls_client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
[b]Note:[/b] It is recommended to specify the scheme part of the URL, i.e. the [param url] should start with either [code]ws://[/code] or [code]wss://[/code].
*/
func (self Expanded) CreateClient(url string, tls_client_options [1]gdclass.TLSOptions) error { //gd:WebSocketMultiplayerPeer.create_client
	return error(gd.ToError(Advanced(self).CreateClient(String.New(url), tls_client_options)))
}

/*
Starts a new multiplayer server listening on the given [param port]. You can optionally specify a [param bind_address], and provide valid [param tls_server_options] to use TLS. See [method TLSOptions.server].
*/
func (self Instance) CreateServer(port int) error { //gd:WebSocketMultiplayerPeer.create_server
	return error(gd.ToError(Advanced(self).CreateServer(int64(port), String.New("*"), [1][1]gdclass.TLSOptions{}[0])))
}

/*
Starts a new multiplayer server listening on the given [param port]. You can optionally specify a [param bind_address], and provide valid [param tls_server_options] to use TLS. See [method TLSOptions.server].
*/
func (self Expanded) CreateServer(port int, bind_address string, tls_server_options [1]gdclass.TLSOptions) error { //gd:WebSocketMultiplayerPeer.create_server
	return error(gd.ToError(Advanced(self).CreateServer(int64(port), String.New(bind_address), tls_server_options)))
}

/*
Returns the [WebSocketPeer] associated to the given [param peer_id].
*/
func (self Instance) GetPeer(peer_id int) [1]gdclass.WebSocketPeer { //gd:WebSocketMultiplayerPeer.get_peer
	return [1]gdclass.WebSocketPeer(Advanced(self).GetPeer(int64(peer_id)))
}

/*
Returns the IP address of the given peer.
*/
func (self Instance) GetPeerAddress(id int) string { //gd:WebSocketMultiplayerPeer.get_peer_address
	return string(Advanced(self).GetPeerAddress(int64(id)).String())
}

/*
Returns the remote port of the given peer.
*/
func (self Instance) GetPeerPort(id int) int { //gd:WebSocketMultiplayerPeer.get_peer_port
	return int(int(Advanced(self).GetPeerPort(int64(id))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.WebSocketMultiplayerPeer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WebSocketMultiplayerPeer"))
	casted := Instance{*(*gdclass.WebSocketMultiplayerPeer)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) SupportedProtocols() []string {
	return []string(class(self).GetSupportedProtocols().Strings())
}

func (self Instance) SetSupportedProtocols(value []string) {
	class(self).SetSupportedProtocols(Packed.MakeStrings(value...))
}

func (self Instance) HandshakeHeaders() []string {
	return []string(class(self).GetHandshakeHeaders().Strings())
}

func (self Instance) SetHandshakeHeaders(value []string) {
	class(self).SetHandshakeHeaders(Packed.MakeStrings(value...))
}

func (self Instance) InboundBufferSize() int {
	return int(int(class(self).GetInboundBufferSize()))
}

func (self Instance) SetInboundBufferSize(value int) {
	class(self).SetInboundBufferSize(int64(value))
}

func (self Instance) OutboundBufferSize() int {
	return int(int(class(self).GetOutboundBufferSize()))
}

func (self Instance) SetOutboundBufferSize(value int) {
	class(self).SetOutboundBufferSize(int64(value))
}

func (self Instance) HandshakeTimeout() Float.X {
	return Float.X(Float.X(class(self).GetHandshakeTimeout()))
}

func (self Instance) SetHandshakeTimeout(value Float.X) {
	class(self).SetHandshakeTimeout(float64(value))
}

func (self Instance) MaxQueuedPackets() int {
	return int(int(class(self).GetMaxQueuedPackets()))
}

func (self Instance) SetMaxQueuedPackets(value int) {
	class(self).SetMaxQueuedPackets(int64(value))
}

/*
Starts a new multiplayer client connecting to the given [param url]. TLS certificates will be verified against the hostname when connecting using the [code]wss://[/code] protocol. You can pass the optional [param tls_client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
[b]Note:[/b] It is recommended to specify the scheme part of the URL, i.e. the [param url] should start with either [code]ws://[/code] or [code]wss://[/code].
*/
//go:nosplit
func (self class) CreateClient(url String.Readable, tls_client_options [1]gdclass.TLSOptions) Error.Code { //gd:WebSocketMultiplayerPeer.create_client
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(url)))
	callframe.Arg(frame, pointers.Get(tls_client_options[0])[0])
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_create_client, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Starts a new multiplayer server listening on the given [param port]. You can optionally specify a [param bind_address], and provide valid [param tls_server_options] to use TLS. See [method TLSOptions.server].
*/
//go:nosplit
func (self class) CreateServer(port int64, bind_address String.Readable, tls_server_options [1]gdclass.TLSOptions) Error.Code { //gd:WebSocketMultiplayerPeer.create_server
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, pointers.Get(gd.InternalString(bind_address)))
	callframe.Arg(frame, pointers.Get(tls_server_options[0])[0])
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_create_server, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the [WebSocketPeer] associated to the given [param peer_id].
*/
//go:nosplit
func (self class) GetPeer(peer_id int64) [1]gdclass.WebSocketPeer { //gd:WebSocketMultiplayerPeer.get_peer
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_get_peer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.WebSocketPeer{gd.PointerWithOwnershipTransferredToGo[gdclass.WebSocketPeer](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the IP address of the given peer.
*/
//go:nosplit
func (self class) GetPeerAddress(id int64) String.Readable { //gd:WebSocketMultiplayerPeer.get_peer_address
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_get_peer_address, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the remote port of the given peer.
*/
//go:nosplit
func (self class) GetPeerPort(id int64) int64 { //gd:WebSocketMultiplayerPeer.get_peer_port
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_get_peer_port, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetSupportedProtocols() Packed.Strings { //gd:WebSocketMultiplayerPeer.get_supported_protocols
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_get_supported_protocols, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSupportedProtocols(protocols Packed.Strings) { //gd:WebSocketMultiplayerPeer.set_supported_protocols
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPackedStrings(protocols)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_set_supported_protocols, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHandshakeHeaders() Packed.Strings { //gd:WebSocketMultiplayerPeer.get_handshake_headers
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_get_handshake_headers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHandshakeHeaders(protocols Packed.Strings) { //gd:WebSocketMultiplayerPeer.set_handshake_headers
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPackedStrings(protocols)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_set_handshake_headers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInboundBufferSize() int64 { //gd:WebSocketMultiplayerPeer.get_inbound_buffer_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_get_inbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInboundBufferSize(buffer_size int64) { //gd:WebSocketMultiplayerPeer.set_inbound_buffer_size
	var frame = callframe.New()
	callframe.Arg(frame, buffer_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_set_inbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOutboundBufferSize() int64 { //gd:WebSocketMultiplayerPeer.get_outbound_buffer_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_get_outbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOutboundBufferSize(buffer_size int64) { //gd:WebSocketMultiplayerPeer.set_outbound_buffer_size
	var frame = callframe.New()
	callframe.Arg(frame, buffer_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_set_outbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHandshakeTimeout() float64 { //gd:WebSocketMultiplayerPeer.get_handshake_timeout
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_get_handshake_timeout, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHandshakeTimeout(timeout float64) { //gd:WebSocketMultiplayerPeer.set_handshake_timeout
	var frame = callframe.New()
	callframe.Arg(frame, timeout)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_set_handshake_timeout, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetMaxQueuedPackets(max_queued_packets int64) { //gd:WebSocketMultiplayerPeer.set_max_queued_packets
	var frame = callframe.New()
	callframe.Arg(frame, max_queued_packets)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_set_max_queued_packets, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxQueuedPackets() int64 { //gd:WebSocketMultiplayerPeer.get_max_queued_packets
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketMultiplayerPeer.Bind_get_max_queued_packets, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsWebSocketMultiplayerPeer() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWebSocketMultiplayerPeer() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsMultiplayerPeer() MultiplayerPeer.Advanced {
	return *((*MultiplayerPeer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsMultiplayerPeer() MultiplayerPeer.Instance {
	return *((*MultiplayerPeer.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(MultiplayerPeer.Advanced(self.AsMultiplayerPeer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(MultiplayerPeer.Instance(self.AsMultiplayerPeer()), name)
	}
}
func init() {
	gdclass.Register("WebSocketMultiplayerPeer", func(ptr gd.Object) any {
		return [1]gdclass.WebSocketMultiplayerPeer{*(*gdclass.WebSocketMultiplayerPeer)(unsafe.Pointer(&ptr))}
	})
}
