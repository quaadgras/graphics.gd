package WebSocketPeer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/PacketPeer"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

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
type Instance [1]classdb.WebSocketPeer

/*
Connects to the given URL. TLS certificates will be verified against the hostname when connecting using the [code]wss://[/code] protocol. You can pass the optional [param tls_client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
[b]Note:[/b] To avoid mixed content warnings or errors in Web, you may have to use a [param url] that starts with [code]wss://[/code] (secure) instead of [code]ws://[/code]. When doing so, make sure to use the fully qualified domain name that matches the one defined in the server's TLS certificate. Do not connect directly via the IP address for [code]wss://[/code] connections, as it won't match with the TLS certificate.
*/
func (self Instance) ConnectToUrl(url string) gd.Error {
	return gd.Error(class(self).ConnectToUrl(gd.NewString(url), ([1]gdclass.TLSOptions{}[0])))
}

/*
Accepts a peer connection performing the HTTP handshake as a WebSocket server. The [param stream] must be a valid TCP stream retrieved via [method TCPServer.take_connection], or a TLS stream accepted via [method StreamPeerTLS.accept_stream].
[b]Note:[/b] Not supported in Web exports due to browsers' restrictions.
*/
func (self Instance) AcceptStream(stream gdclass.StreamPeer) gd.Error {
	return gd.Error(class(self).AcceptStream(stream))
}

/*
Sends the given [param message] using the desired [param write_mode]. When sending a [String], prefer using [method send_text].
*/
func (self Instance) Send(message []byte) gd.Error {
	return gd.Error(class(self).Send(gd.NewPackedByteSlice(message), 1))
}

/*
Sends the given [param message] using WebSocket text mode. Prefer this method over [method PacketPeer.put_packet] when interacting with third-party text-based API (e.g. when using [JSON] formatted messages).
*/
func (self Instance) SendText(message string) gd.Error {
	return gd.Error(class(self).SendText(gd.NewString(message)))
}

/*
Returns [code]true[/code] if the last received packet was sent as a text payload. See [enum WriteMode].
*/
func (self Instance) WasStringPacket() bool {
	return bool(class(self).WasStringPacket())
}

/*
Updates the connection state and receive incoming packets. Call this function regularly to keep it in a clean state.
*/
func (self Instance) Poll() {
	class(self).Poll()
}

/*
Closes this WebSocket connection. [param code] is the status code for the closure (see RFC 6455 section 7.4 for a list of valid status codes). [param reason] is the human readable reason for closing the connection (can be any UTF-8 string that's smaller than 123 bytes). If [param code] is negative, the connection will be closed immediately without notifying the remote peer.
[b]Note:[/b] To achieve a clean close, you will need to keep polling until [constant STATE_CLOSED] is reached.
[b]Note:[/b] The Web export might not support all status codes. Please refer to browser-specific documentation for more details.
*/
func (self Instance) Close() {
	class(self).Close(gd.Int(1000), gd.NewString(""))
}

/*
Returns the IP address of the connected peer.
[b]Note:[/b] Not available in the Web export.
*/
func (self Instance) GetConnectedHost() string {
	return string(class(self).GetConnectedHost().String())
}

/*
Returns the remote port of the connected peer.
[b]Note:[/b] Not available in the Web export.
*/
func (self Instance) GetConnectedPort() int {
	return int(int(class(self).GetConnectedPort()))
}

/*
Returns the selected WebSocket sub-protocol for this connection or an empty string if the sub-protocol has not been selected yet.
*/
func (self Instance) GetSelectedProtocol() string {
	return string(class(self).GetSelectedProtocol().String())
}

/*
Returns the URL requested by this peer. The URL is derived from the [code]url[/code] passed to [method connect_to_url] or from the HTTP headers when acting as server (i.e. when using [method accept_stream]).
*/
func (self Instance) GetRequestedUrl() string {
	return string(class(self).GetRequestedUrl().String())
}

/*
Disable Nagle's algorithm on the underlying TCP socket (default). See [method StreamPeerTCP.set_no_delay] for more information.
[b]Note:[/b] Not available in the Web export.
*/
func (self Instance) SetNoDelay(enabled bool) {
	class(self).SetNoDelay(enabled)
}

/*
Returns the current amount of data in the outbound websocket buffer. [b]Note:[/b] Web exports use WebSocket.bufferedAmount, while other platforms use an internal buffer.
*/
func (self Instance) GetCurrentOutboundBufferedAmount() int {
	return int(int(class(self).GetCurrentOutboundBufferedAmount()))
}

/*
Returns the ready state of the connection. See [enum State].
*/
func (self Instance) GetReadyState() classdb.WebSocketPeerState {
	return classdb.WebSocketPeerState(class(self).GetReadyState())
}

/*
Returns the received WebSocket close frame status code, or [code]-1[/code] when the connection was not cleanly closed. Only call this method when [method get_ready_state] returns [constant STATE_CLOSED].
*/
func (self Instance) GetCloseCode() int {
	return int(int(class(self).GetCloseCode()))
}

/*
Returns the received WebSocket close frame status reason string. Only call this method when [method get_ready_state] returns [constant STATE_CLOSED].
*/
func (self Instance) GetCloseReason() string {
	return string(class(self).GetCloseReason().String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.WebSocketPeer

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WebSocketPeer"))
	return Instance{classdb.WebSocketPeer(object)}
}

func (self Instance) SupportedProtocols() []string {
	return []string(class(self).GetSupportedProtocols().Strings())
}

func (self Instance) SetSupportedProtocols(value []string) {
	class(self).SetSupportedProtocols(gd.NewPackedStringSlice(value))
}

func (self Instance) HandshakeHeaders() []string {
	return []string(class(self).GetHandshakeHeaders().Strings())
}

func (self Instance) SetHandshakeHeaders(value []string) {
	class(self).SetHandshakeHeaders(gd.NewPackedStringSlice(value))
}

func (self Instance) InboundBufferSize() int {
	return int(int(class(self).GetInboundBufferSize()))
}

func (self Instance) SetInboundBufferSize(value int) {
	class(self).SetInboundBufferSize(gd.Int(value))
}

func (self Instance) OutboundBufferSize() int {
	return int(int(class(self).GetOutboundBufferSize()))
}

func (self Instance) SetOutboundBufferSize(value int) {
	class(self).SetOutboundBufferSize(gd.Int(value))
}

func (self Instance) MaxQueuedPackets() int {
	return int(int(class(self).GetMaxQueuedPackets()))
}

func (self Instance) SetMaxQueuedPackets(value int) {
	class(self).SetMaxQueuedPackets(gd.Int(value))
}

/*
Connects to the given URL. TLS certificates will be verified against the hostname when connecting using the [code]wss://[/code] protocol. You can pass the optional [param tls_client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
[b]Note:[/b] To avoid mixed content warnings or errors in Web, you may have to use a [param url] that starts with [code]wss://[/code] (secure) instead of [code]ws://[/code]. When doing so, make sure to use the fully qualified domain name that matches the one defined in the server's TLS certificate. Do not connect directly via the IP address for [code]wss://[/code] connections, as it won't match with the TLS certificate.
*/
//go:nosplit
func (self class) ConnectToUrl(url gd.String, tls_client_options gdclass.TLSOptions) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(url))
	callframe.Arg(frame, pointers.Get(tls_client_options[0])[0])
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_connect_to_url, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(stream[0])[0])
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_accept_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sends the given [param message] using the desired [param write_mode]. When sending a [String], prefer using [method send_text].
*/
//go:nosplit
func (self class) Send(message gd.PackedByteArray, write_mode classdb.WebSocketPeerWriteMode) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(message))
	callframe.Arg(frame, write_mode)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_send, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sends the given [param message] using WebSocket text mode. Prefer this method over [method PacketPeer.put_packet] when interacting with third-party text-based API (e.g. when using [JSON] formatted messages).
*/
//go:nosplit
func (self class) SendText(message gd.String) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(message))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_send_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the last received packet was sent as a text payload. See [enum WriteMode].
*/
//go:nosplit
func (self class) WasStringPacket() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_was_string_packet, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Updates the connection state and receive incoming packets. Call this function regularly to keep it in a clean state.
*/
//go:nosplit
func (self class) Poll() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Closes this WebSocket connection. [param code] is the status code for the closure (see RFC 6455 section 7.4 for a list of valid status codes). [param reason] is the human readable reason for closing the connection (can be any UTF-8 string that's smaller than 123 bytes). If [param code] is negative, the connection will be closed immediately without notifying the remote peer.
[b]Note:[/b] To achieve a clean close, you will need to keep polling until [constant STATE_CLOSED] is reached.
[b]Note:[/b] The Web export might not support all status codes. Please refer to browser-specific documentation for more details.
*/
//go:nosplit
func (self class) Close(code gd.Int, reason gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, code)
	callframe.Arg(frame, pointers.Get(reason))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_close, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the IP address of the connected peer.
[b]Note:[/b] Not available in the Web export.
*/
//go:nosplit
func (self class) GetConnectedHost() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_get_connected_host, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the remote port of the connected peer.
[b]Note:[/b] Not available in the Web export.
*/
//go:nosplit
func (self class) GetConnectedPort() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_get_connected_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the selected WebSocket sub-protocol for this connection or an empty string if the sub-protocol has not been selected yet.
*/
//go:nosplit
func (self class) GetSelectedProtocol() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_get_selected_protocol, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the URL requested by this peer. The URL is derived from the [code]url[/code] passed to [method connect_to_url] or from the HTTP headers when acting as server (i.e. when using [method accept_stream]).
*/
//go:nosplit
func (self class) GetRequestedUrl() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_get_requested_url, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Disable Nagle's algorithm on the underlying TCP socket (default). See [method StreamPeerTCP.set_no_delay] for more information.
[b]Note:[/b] Not available in the Web export.
*/
//go:nosplit
func (self class) SetNoDelay(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_set_no_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the current amount of data in the outbound websocket buffer. [b]Note:[/b] Web exports use WebSocket.bufferedAmount, while other platforms use an internal buffer.
*/
//go:nosplit
func (self class) GetCurrentOutboundBufferedAmount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_get_current_outbound_buffered_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the ready state of the connection. See [enum State].
*/
//go:nosplit
func (self class) GetReadyState() classdb.WebSocketPeerState {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.WebSocketPeerState](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_get_ready_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the received WebSocket close frame status code, or [code]-1[/code] when the connection was not cleanly closed. Only call this method when [method get_ready_state] returns [constant STATE_CLOSED].
*/
//go:nosplit
func (self class) GetCloseCode() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_get_close_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the received WebSocket close frame status reason string. Only call this method when [method get_ready_state] returns [constant STATE_CLOSED].
*/
//go:nosplit
func (self class) GetCloseReason() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_get_close_reason, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetSupportedProtocols() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_get_supported_protocols, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSupportedProtocols(protocols gd.PackedStringArray) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(protocols))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_set_supported_protocols, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHandshakeHeaders() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_get_handshake_headers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHandshakeHeaders(protocols gd.PackedStringArray) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(protocols))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_set_handshake_headers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetInboundBufferSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_get_inbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInboundBufferSize(buffer_size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, buffer_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_set_inbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOutboundBufferSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_get_outbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOutboundBufferSize(buffer_size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, buffer_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_set_outbound_buffer_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetMaxQueuedPackets(buffer_size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, buffer_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_set_max_queued_packets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxQueuedPackets() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebSocketPeer.Bind_get_max_queued_packets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsWebSocketPeer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWebSocketPeer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPacketPeer() PacketPeer.Advanced {
	return *((*PacketPeer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPacketPeer() PacketPeer.Instance {
	return *((*PacketPeer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsPacketPeer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsPacketPeer(), name)
	}
}
func init() {
	classdb.Register("WebSocketPeer", func(ptr gd.Object) any { return classdb.WebSocketPeer(ptr) })
}

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
