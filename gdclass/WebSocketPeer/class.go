package WebSocketPeer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/PacketPeer"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class represents WebSocket connection, and can be used as a WebSocket client (RFC 6455-compliant) or as a remote peer of a WebSocket server.
You can send WebSocket binary frames using [method PacketPeer.put_packet], and WebSocket text frames using [method send] (prefer text frames when interacting with text-based API). You can check the frame type of the last packet via [method was_string_packet].
To start a WebSocket client, first call [method connect_to_url], then regularly call [method poll] (e.g. during [Node] process). You can query the socket state via [method get_ready_state], get the number of pending packets using [method PacketPeer.get_available_packet_count], and retrieve them via [method PacketPeer.get_packet].
[codeblocks]
[gdscript]
extends Node

var socket = WebSocketPeer.new()

func _ready():
    socket.connect_to_url("wss://example.com")

func _process(delta):
    socket.poll()
    var state = socket.get_ready_state()
    if state == WebSocketPeer.STATE_OPEN:
        while socket.get_available_packet_count():
            print("Packet: ", socket.get_packet())
    elif state == WebSocketPeer.STATE_CLOSING:
        # Keep polling to achieve proper close.
        pass
    elif state == WebSocketPeer.STATE_CLOSED:
        var code = socket.get_close_code()
        var reason = socket.get_close_reason()
        print("WebSocket closed with code: %d, reason %s. Clean: %s" % [code, reason, code != -1])
        set_process(false) # Stop processing.
[/gdscript]
[/codeblocks]
To use the peer as part of a WebSocket server refer to [method accept_stream] and the online tutorial.

*/
type Go [1]classdb.WebSocketPeer

/*
Connects to the given URL. TLS certificates will be verified against the hostname when connecting using the [code]wss://[/code] protocol. You can pass the optional [param tls_client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
[b]Note:[/b] To avoid mixed content warnings or errors in Web, you may have to use a [param url] that starts with [code]wss://[/code] (secure) instead of [code]ws://[/code]. When doing so, make sure to use the fully qualified domain name that matches the one defined in the server's TLS certificate. Do not connect directly via the IP address for [code]wss://[/code] connections, as it won't match with the TLS certificate.
*/
func (self Go) ConnectToUrl(url string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).ConnectToUrl(gc.String(url), ([1]gdclass.TLSOptions{}[0])))
}

/*
Accepts a peer connection performing the HTTP handshake as a WebSocket server. The [param stream] must be a valid TCP stream retrieved via [method TCPServer.take_connection], or a TLS stream accepted via [method StreamPeerTLS.accept_stream].
[b]Note:[/b] Not supported in Web exports due to browsers' restrictions.
*/
func (self Go) AcceptStream(stream gdclass.StreamPeer) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).AcceptStream(stream))
}

/*
Sends the given [param message] using the desired [param write_mode]. When sending a [String], prefer using [method send_text].
*/
func (self Go) Send(message []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Send(gc.PackedByteSlice(message), 1))
}

/*
Sends the given [param message] using WebSocket text mode. Prefer this method over [method PacketPeer.put_packet] when interacting with third-party text-based API (e.g. when using [JSON] formatted messages).
*/
func (self Go) SendText(message string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).SendText(gc.String(message)))
}

/*
Returns [code]true[/code] if the last received packet was sent as a text payload. See [enum WriteMode].
*/
func (self Go) WasStringPacket() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).WasStringPacket())
}

/*
Updates the connection state and receive incoming packets. Call this function regularly to keep it in a clean state.
*/
func (self Go) Poll() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Poll()
}

/*
Closes this WebSocket connection. [param code] is the status code for the closure (see RFC 6455 section 7.4 for a list of valid status codes). [param reason] is the human readable reason for closing the connection (can be any UTF-8 string that's smaller than 123 bytes). If [param code] is negative, the connection will be closed immediately without notifying the remote peer.
[b]Note:[/b] To achieve a clean close, you will need to keep polling until [constant STATE_CLOSED] is reached.
[b]Note:[/b] The Web export might not support all status codes. Please refer to browser-specific documentation for more details.
*/
func (self Go) Close() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Close(gd.Int(1000), gc.String(""))
}

/*
Returns the IP address of the connected peer.
[b]Note:[/b] Not available in the Web export.
*/
func (self Go) GetConnectedHost() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetConnectedHost(gc).String())
}

/*
Returns the remote port of the connected peer.
[b]Note:[/b] Not available in the Web export.
*/
func (self Go) GetConnectedPort() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetConnectedPort()))
}

/*
Returns the selected WebSocket sub-protocol for this connection or an empty string if the sub-protocol has not been selected yet.
*/
func (self Go) GetSelectedProtocol() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetSelectedProtocol(gc).String())
}

/*
Returns the URL requested by this peer. The URL is derived from the [code]url[/code] passed to [method connect_to_url] or from the HTTP headers when acting as server (i.e. when using [method accept_stream]).
*/
func (self Go) GetRequestedUrl() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetRequestedUrl(gc).String())
}

/*
Disable Nagle's algorithm on the underlying TCP socket (default). See [method StreamPeerTCP.set_no_delay] for more information.
[b]Note:[/b] Not available in the Web export.
*/
func (self Go) SetNoDelay(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNoDelay(enabled)
}

/*
Returns the current amount of data in the outbound websocket buffer. [b]Note:[/b] Web exports use WebSocket.bufferedAmount, while other platforms use an internal buffer.
*/
func (self Go) GetCurrentOutboundBufferedAmount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCurrentOutboundBufferedAmount()))
}

/*
Returns the ready state of the connection. See [enum State].
*/
func (self Go) GetReadyState() classdb.WebSocketPeerState {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.WebSocketPeerState(class(self).GetReadyState())
}

/*
Returns the received WebSocket close frame status code, or [code]-1[/code] when the connection was not cleanly closed. Only call this method when [method get_ready_state] returns [constant STATE_CLOSED].
*/
func (self Go) GetCloseCode() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCloseCode()))
}

/*
Returns the received WebSocket close frame status reason string. Only call this method when [method get_ready_state] returns [constant STATE_CLOSED].
*/
func (self Go) GetCloseReason() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetCloseReason(gc).String())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.WebSocketPeer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("WebSocketPeer"))
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

func (self Go) MaxQueuedPackets() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMaxQueuedPackets()))
}

func (self Go) SetMaxQueuedPackets(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxQueuedPackets(gd.Int(value))
}

/*
Connects to the given URL. TLS certificates will be verified against the hostname when connecting using the [code]wss://[/code] protocol. You can pass the optional [param tls_client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
[b]Note:[/b] To avoid mixed content warnings or errors in Web, you may have to use a [param url] that starts with [code]wss://[/code] (secure) instead of [code]ws://[/code]. When doing so, make sure to use the fully qualified domain name that matches the one defined in the server's TLS certificate. Do not connect directly via the IP address for [code]wss://[/code] connections, as it won't match with the TLS certificate.
*/
//go:nosplit
func (self class) ConnectToUrl(url gd.String, tls_client_options gdclass.TLSOptions) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(url))
	callframe.Arg(frame, mmm.Get(tls_client_options[0].AsPointer())[0])
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_connect_to_url, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Accepts a peer connection performing the HTTP handshake as a WebSocket server. The [param stream] must be a valid TCP stream retrieved via [method TCPServer.take_connection], or a TLS stream accepted via [method StreamPeerTLS.accept_stream].
[b]Note:[/b] Not supported in Web exports due to browsers' restrictions.
*/
//go:nosplit
func (self class) AcceptStream(stream gdclass.StreamPeer) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(stream[0].AsPointer())[0])
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_accept_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sends the given [param message] using the desired [param write_mode]. When sending a [String], prefer using [method send_text].
*/
//go:nosplit
func (self class) Send(message gd.PackedByteArray, write_mode classdb.WebSocketPeerWriteMode) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(message))
	callframe.Arg(frame, write_mode)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_send, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sends the given [param message] using WebSocket text mode. Prefer this method over [method PacketPeer.put_packet] when interacting with third-party text-based API (e.g. when using [JSON] formatted messages).
*/
//go:nosplit
func (self class) SendText(message gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(message))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_send_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the last received packet was sent as a text payload. See [enum WriteMode].
*/
//go:nosplit
func (self class) WasStringPacket() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_was_string_packet, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Updates the connection state and receive incoming packets. Call this function regularly to keep it in a clean state.
*/
//go:nosplit
func (self class) Poll()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Closes this WebSocket connection. [param code] is the status code for the closure (see RFC 6455 section 7.4 for a list of valid status codes). [param reason] is the human readable reason for closing the connection (can be any UTF-8 string that's smaller than 123 bytes). If [param code] is negative, the connection will be closed immediately without notifying the remote peer.
[b]Note:[/b] To achieve a clean close, you will need to keep polling until [constant STATE_CLOSED] is reached.
[b]Note:[/b] The Web export might not support all status codes. Please refer to browser-specific documentation for more details.
*/
//go:nosplit
func (self class) Close(code gd.Int, reason gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, code)
	callframe.Arg(frame, mmm.Get(reason))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_close, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the IP address of the connected peer.
[b]Note:[/b] Not available in the Web export.
*/
//go:nosplit
func (self class) GetConnectedHost(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_get_connected_host, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the remote port of the connected peer.
[b]Note:[/b] Not available in the Web export.
*/
//go:nosplit
func (self class) GetConnectedPort() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_get_connected_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the selected WebSocket sub-protocol for this connection or an empty string if the sub-protocol has not been selected yet.
*/
//go:nosplit
func (self class) GetSelectedProtocol(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_get_selected_protocol, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the URL requested by this peer. The URL is derived from the [code]url[/code] passed to [method connect_to_url] or from the HTTP headers when acting as server (i.e. when using [method accept_stream]).
*/
//go:nosplit
func (self class) GetRequestedUrl(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_get_requested_url, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Disable Nagle's algorithm on the underlying TCP socket (default). See [method StreamPeerTCP.set_no_delay] for more information.
[b]Note:[/b] Not available in the Web export.
*/
//go:nosplit
func (self class) SetNoDelay(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_set_no_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the current amount of data in the outbound websocket buffer. [b]Note:[/b] Web exports use WebSocket.bufferedAmount, while other platforms use an internal buffer.
*/
//go:nosplit
func (self class) GetCurrentOutboundBufferedAmount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_get_current_outbound_buffered_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the ready state of the connection. See [enum State].
*/
//go:nosplit
func (self class) GetReadyState() classdb.WebSocketPeerState {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.WebSocketPeerState](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_get_ready_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the received WebSocket close frame status code, or [code]-1[/code] when the connection was not cleanly closed. Only call this method when [method get_ready_state] returns [constant STATE_CLOSED].
*/
//go:nosplit
func (self class) GetCloseCode() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_get_close_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the received WebSocket close frame status reason string. Only call this method when [method get_ready_state] returns [constant STATE_CLOSED].
*/
//go:nosplit
func (self class) GetCloseReason(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_get_close_reason, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetSupportedProtocols(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_get_supported_protocols, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_set_supported_protocols, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHandshakeHeaders(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_get_handshake_headers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_set_handshake_headers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInboundBufferSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_get_inbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_set_inbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOutboundBufferSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_get_outbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_set_outbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetMaxQueuedPackets(buffer_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, buffer_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_set_max_queued_packets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxQueuedPackets() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebSocketPeer.Bind_get_max_queued_packets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsWebSocketPeer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsWebSocketPeer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsPacketPeer() PacketPeer.GD { return *((*PacketPeer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsPacketPeer() PacketPeer.Go { return *((*PacketPeer.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsPacketPeer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsPacketPeer(), name)
	}
}
func init() {classdb.Register("WebSocketPeer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type WriteMode = classdb.WebSocketPeerWriteMode

const (
/*Specifies that WebSockets messages should be transferred as text payload (only valid UTF-8 is allowed).*/
	WriteModeText WriteMode = 0
/*Specifies that WebSockets messages should be transferred as binary payload (any byte combination is allowed).*/
	WriteModeBinary WriteMode = 1
)
type State = classdb.WebSocketPeerState

const (
/*Socket has been created. The connection is not yet open.*/
	StateConnecting State = 0
/*The connection is open and ready to communicate.*/
	StateOpen State = 1
/*The connection is in the process of closing. This means a close request has been sent to the remote peer but confirmation has not been received.*/
	StateClosing State = 2
/*The connection is closed or couldn't be opened.*/
	StateClosed State = 3
)
