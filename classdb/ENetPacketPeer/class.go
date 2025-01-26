// Package ENetPacketPeer provides methods for working with ENetPacketPeer object instances.
package ENetPacketPeer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/classdb/PacketPeer"
import "graphics.gd/variant/Float"

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

/*
A PacketPeer implementation representing a peer of an [ENetConnection].
This class cannot be instantiated directly but can be retrieved during [method ENetConnection.service] or via [method ENetConnection.get_peers].
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
*/
type Instance [1]gdclass.ENetPacketPeer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsENetPacketPeer() Instance
}

/*
Request a disconnection from a peer. An [constant ENetConnection.EVENT_DISCONNECT] will be generated during [method ENetConnection.service] once the disconnection is complete.
*/
func (self Instance) PeerDisconnect() { //gd:ENetPacketPeer.peer_disconnect
	class(self).PeerDisconnect(gd.Int(0))
}

/*
Request a disconnection from a peer, but only after all queued outgoing packets are sent. An [constant ENetConnection.EVENT_DISCONNECT] will be generated during [method ENetConnection.service] once the disconnection is complete.
*/
func (self Instance) PeerDisconnectLater() { //gd:ENetPacketPeer.peer_disconnect_later
	class(self).PeerDisconnectLater(gd.Int(0))
}

/*
Force an immediate disconnection from a peer. No [constant ENetConnection.EVENT_DISCONNECT] will be generated. The foreign peer is not guaranteed to receive the disconnect notification, and is reset immediately upon return from this function.
*/
func (self Instance) PeerDisconnectNow() { //gd:ENetPacketPeer.peer_disconnect_now
	class(self).PeerDisconnectNow(gd.Int(0))
}

/*
Sends a ping request to a peer. ENet automatically pings all connected peers at regular intervals, however, this function may be called to ensure more frequent ping requests.
*/
func (self Instance) Ping() { //gd:ENetPacketPeer.ping
	class(self).Ping()
}

/*
Sets the [param ping_interval] in milliseconds at which pings will be sent to a peer. Pings are used both to monitor the liveness of the connection and also to dynamically adjust the throttle during periods of low traffic so that the throttle has reasonable responsiveness during traffic spikes. The default ping interval is [code]500[/code] milliseconds.
*/
func (self Instance) PingInterval(ping_interval int) { //gd:ENetPacketPeer.ping_interval
	class(self).PingInterval(gd.Int(ping_interval))
}

/*
Forcefully disconnects a peer. The foreign host represented by the peer is not notified of the disconnection and will timeout on its connection to the local host.
*/
func (self Instance) Reset() { //gd:ENetPacketPeer.reset
	class(self).Reset()
}

/*
Queues a [param packet] to be sent over the specified [param channel]. See [code]FLAG_*[/code] constants for available packet flags.
*/
func (self Instance) Send(channel int, packet []byte, flags int) error { //gd:ENetPacketPeer.send
	return error(gd.ToError(class(self).Send(gd.Int(channel), gd.NewPackedByteSlice(packet), gd.Int(flags))))
}

/*
Configures throttle parameter for a peer.
Unreliable packets are dropped by ENet in response to the varying conditions of the Internet connection to the peer. The throttle represents a probability that an unreliable packet should not be dropped and thus sent by ENet to the peer. By measuring fluctuations in round trip times of reliable packets over the specified [param interval], ENet will either increase the probability by the amount specified in the [param acceleration] parameter, or decrease it by the amount specified in the [param deceleration] parameter (both are ratios to [constant PACKET_THROTTLE_SCALE]).
When the throttle has a value of [constant PACKET_THROTTLE_SCALE], no unreliable packets are dropped by ENet, and so 100% of all unreliable packets will be sent.
When the throttle has a value of [code]0[/code], all unreliable packets are dropped by ENet, and so 0% of all unreliable packets will be sent.
Intermediate values for the throttle represent intermediate probabilities between 0% and 100% of unreliable packets being sent. The bandwidth limits of the local and foreign hosts are taken into account to determine a sensible limit for the throttle probability above which it should not raise even in the best of conditions.
*/
func (self Instance) ThrottleConfigure(interval int, acceleration int, deceleration int) { //gd:ENetPacketPeer.throttle_configure
	class(self).ThrottleConfigure(gd.Int(interval), gd.Int(acceleration), gd.Int(deceleration))
}

/*
Sets the timeout parameters for a peer. The timeout parameters control how and when a peer will timeout from a failure to acknowledge reliable traffic. Timeout values are expressed in milliseconds.
The [param timeout] is a factor that, multiplied by a value based on the average round trip time, will determine the timeout limit for a reliable packet. When that limit is reached, the timeout will be doubled, and the peer will be disconnected if that limit has reached [param timeout_min]. The [param timeout_max] parameter, on the other hand, defines a fixed timeout for which any packet must be acknowledged or the peer will be dropped.
*/
func (self Instance) SetTimeout(timeout int, timeout_min int, timeout_max int) { //gd:ENetPacketPeer.set_timeout
	class(self).SetTimeout(gd.Int(timeout), gd.Int(timeout_min), gd.Int(timeout_max))
}

/*
Returns the IP address of this peer.
*/
func (self Instance) GetRemoteAddress() string { //gd:ENetPacketPeer.get_remote_address
	return string(class(self).GetRemoteAddress().String())
}

/*
Returns the remote port of this peer.
*/
func (self Instance) GetRemotePort() int { //gd:ENetPacketPeer.get_remote_port
	return int(int(class(self).GetRemotePort()))
}

/*
Returns the requested [param statistic] for this peer. See [enum PeerStatistic].
*/
func (self Instance) GetStatistic(statistic gdclass.ENetPacketPeerPeerStatistic) Float.X { //gd:ENetPacketPeer.get_statistic
	return Float.X(Float.X(class(self).GetStatistic(statistic)))
}

/*
Returns the current peer state. See [enum PeerState].
*/
func (self Instance) GetState() gdclass.ENetPacketPeerPeerState { //gd:ENetPacketPeer.get_state
	return gdclass.ENetPacketPeerPeerState(class(self).GetState())
}

/*
Returns the number of channels allocated for communication with peer.
*/
func (self Instance) GetChannels() int { //gd:ENetPacketPeer.get_channels
	return int(int(class(self).GetChannels()))
}

/*
Returns [code]true[/code] if the peer is currently active (i.e. the associated [ENetConnection] is still valid).
*/
func (self Instance) IsActive() bool { //gd:ENetPacketPeer.is_active
	return bool(class(self).IsActive())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ENetPacketPeer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ENetPacketPeer"))
	casted := Instance{*(*gdclass.ENetPacketPeer)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Request a disconnection from a peer. An [constant ENetConnection.EVENT_DISCONNECT] will be generated during [method ENetConnection.service] once the disconnection is complete.
*/
//go:nosplit
func (self class) PeerDisconnect(data gd.Int) { //gd:ENetPacketPeer.peer_disconnect
	var frame = callframe.New()
	callframe.Arg(frame, data)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetPacketPeer.Bind_peer_disconnect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Request a disconnection from a peer, but only after all queued outgoing packets are sent. An [constant ENetConnection.EVENT_DISCONNECT] will be generated during [method ENetConnection.service] once the disconnection is complete.
*/
//go:nosplit
func (self class) PeerDisconnectLater(data gd.Int) { //gd:ENetPacketPeer.peer_disconnect_later
	var frame = callframe.New()
	callframe.Arg(frame, data)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetPacketPeer.Bind_peer_disconnect_later, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Force an immediate disconnection from a peer. No [constant ENetConnection.EVENT_DISCONNECT] will be generated. The foreign peer is not guaranteed to receive the disconnect notification, and is reset immediately upon return from this function.
*/
//go:nosplit
func (self class) PeerDisconnectNow(data gd.Int) { //gd:ENetPacketPeer.peer_disconnect_now
	var frame = callframe.New()
	callframe.Arg(frame, data)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetPacketPeer.Bind_peer_disconnect_now, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sends a ping request to a peer. ENet automatically pings all connected peers at regular intervals, however, this function may be called to ensure more frequent ping requests.
*/
//go:nosplit
func (self class) Ping() { //gd:ENetPacketPeer.ping
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetPacketPeer.Bind_ping, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the [param ping_interval] in milliseconds at which pings will be sent to a peer. Pings are used both to monitor the liveness of the connection and also to dynamically adjust the throttle during periods of low traffic so that the throttle has reasonable responsiveness during traffic spikes. The default ping interval is [code]500[/code] milliseconds.
*/
//go:nosplit
func (self class) PingInterval(ping_interval gd.Int) { //gd:ENetPacketPeer.ping_interval
	var frame = callframe.New()
	callframe.Arg(frame, ping_interval)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetPacketPeer.Bind_ping_interval, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Forcefully disconnects a peer. The foreign host represented by the peer is not notified of the disconnection and will timeout on its connection to the local host.
*/
//go:nosplit
func (self class) Reset() { //gd:ENetPacketPeer.reset
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetPacketPeer.Bind_reset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Queues a [param packet] to be sent over the specified [param channel]. See [code]FLAG_*[/code] constants for available packet flags.
*/
//go:nosplit
func (self class) Send(channel gd.Int, packet gd.PackedByteArray, flags gd.Int) gd.Error { //gd:ENetPacketPeer.send
	var frame = callframe.New()
	callframe.Arg(frame, channel)
	callframe.Arg(frame, pointers.Get(packet))
	callframe.Arg(frame, flags)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetPacketPeer.Bind_send, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Configures throttle parameter for a peer.
Unreliable packets are dropped by ENet in response to the varying conditions of the Internet connection to the peer. The throttle represents a probability that an unreliable packet should not be dropped and thus sent by ENet to the peer. By measuring fluctuations in round trip times of reliable packets over the specified [param interval], ENet will either increase the probability by the amount specified in the [param acceleration] parameter, or decrease it by the amount specified in the [param deceleration] parameter (both are ratios to [constant PACKET_THROTTLE_SCALE]).
When the throttle has a value of [constant PACKET_THROTTLE_SCALE], no unreliable packets are dropped by ENet, and so 100% of all unreliable packets will be sent.
When the throttle has a value of [code]0[/code], all unreliable packets are dropped by ENet, and so 0% of all unreliable packets will be sent.
Intermediate values for the throttle represent intermediate probabilities between 0% and 100% of unreliable packets being sent. The bandwidth limits of the local and foreign hosts are taken into account to determine a sensible limit for the throttle probability above which it should not raise even in the best of conditions.
*/
//go:nosplit
func (self class) ThrottleConfigure(interval gd.Int, acceleration gd.Int, deceleration gd.Int) { //gd:ENetPacketPeer.throttle_configure
	var frame = callframe.New()
	callframe.Arg(frame, interval)
	callframe.Arg(frame, acceleration)
	callframe.Arg(frame, deceleration)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetPacketPeer.Bind_throttle_configure, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the timeout parameters for a peer. The timeout parameters control how and when a peer will timeout from a failure to acknowledge reliable traffic. Timeout values are expressed in milliseconds.
The [param timeout] is a factor that, multiplied by a value based on the average round trip time, will determine the timeout limit for a reliable packet. When that limit is reached, the timeout will be doubled, and the peer will be disconnected if that limit has reached [param timeout_min]. The [param timeout_max] parameter, on the other hand, defines a fixed timeout for which any packet must be acknowledged or the peer will be dropped.
*/
//go:nosplit
func (self class) SetTimeout(timeout gd.Int, timeout_min gd.Int, timeout_max gd.Int) { //gd:ENetPacketPeer.set_timeout
	var frame = callframe.New()
	callframe.Arg(frame, timeout)
	callframe.Arg(frame, timeout_min)
	callframe.Arg(frame, timeout_max)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetPacketPeer.Bind_set_timeout, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the IP address of this peer.
*/
//go:nosplit
func (self class) GetRemoteAddress() gd.String { //gd:ENetPacketPeer.get_remote_address
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetPacketPeer.Bind_get_remote_address, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the remote port of this peer.
*/
//go:nosplit
func (self class) GetRemotePort() gd.Int { //gd:ENetPacketPeer.get_remote_port
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetPacketPeer.Bind_get_remote_port, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the requested [param statistic] for this peer. See [enum PeerStatistic].
*/
//go:nosplit
func (self class) GetStatistic(statistic gdclass.ENetPacketPeerPeerStatistic) gd.Float { //gd:ENetPacketPeer.get_statistic
	var frame = callframe.New()
	callframe.Arg(frame, statistic)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetPacketPeer.Bind_get_statistic, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current peer state. See [enum PeerState].
*/
//go:nosplit
func (self class) GetState() gdclass.ENetPacketPeerPeerState { //gd:ENetPacketPeer.get_state
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.ENetPacketPeerPeerState](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetPacketPeer.Bind_get_state, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of channels allocated for communication with peer.
*/
//go:nosplit
func (self class) GetChannels() gd.Int { //gd:ENetPacketPeer.get_channels
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetPacketPeer.Bind_get_channels, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the peer is currently active (i.e. the associated [ENetConnection] is still valid).
*/
//go:nosplit
func (self class) IsActive() bool { //gd:ENetPacketPeer.is_active
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetPacketPeer.Bind_is_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsENetPacketPeer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsENetPacketPeer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(PacketPeer.Advanced(self.AsPacketPeer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PacketPeer.Instance(self.AsPacketPeer()), name)
	}
}
func init() {
	gdclass.Register("ENetPacketPeer", func(ptr gd.Object) any {
		return [1]gdclass.ENetPacketPeer{*(*gdclass.ENetPacketPeer)(unsafe.Pointer(&ptr))}
	})
}

type PeerState = gdclass.ENetPacketPeerPeerState //gd:ENetPacketPeer.PeerState

const (
	/*The peer is disconnected.*/
	StateDisconnected PeerState = 0
	/*The peer is currently attempting to connect.*/
	StateConnecting PeerState = 1
	/*The peer has acknowledged the connection request.*/
	StateAcknowledgingConnect PeerState = 2
	/*The peer is currently connecting.*/
	StateConnectionPending PeerState = 3
	/*The peer has successfully connected, but is not ready to communicate with yet ([constant STATE_CONNECTED]).*/
	StateConnectionSucceeded PeerState = 4
	/*The peer is currently connected and ready to communicate with.*/
	StateConnected PeerState = 5
	/*The peer is slated to disconnect after it has no more outgoing packets to send.*/
	StateDisconnectLater PeerState = 6
	/*The peer is currently disconnecting.*/
	StateDisconnecting PeerState = 7
	/*The peer has acknowledged the disconnection request.*/
	StateAcknowledgingDisconnect PeerState = 8
	/*The peer has lost connection, but is not considered truly disconnected (as the peer didn't acknowledge the disconnection request).*/
	StateZombie PeerState = 9
)

type PeerStatistic = gdclass.ENetPacketPeerPeerStatistic //gd:ENetPacketPeer.PeerStatistic

const (
	/*Mean packet loss of reliable packets as a ratio with respect to the [constant PACKET_LOSS_SCALE].*/
	PeerPacketLoss PeerStatistic = 0
	/*Packet loss variance.*/
	PeerPacketLossVariance PeerStatistic = 1
	/*The time at which packet loss statistics were last updated (in milliseconds since the connection started). The interval for packet loss statistics updates is 10 seconds, and at least one packet must have been sent since the last statistics update.*/
	PeerPacketLossEpoch PeerStatistic = 2
	/*Mean packet round trip time for reliable packets.*/
	PeerRoundTripTime PeerStatistic = 3
	/*Variance of the mean round trip time.*/
	PeerRoundTripTimeVariance PeerStatistic = 4
	/*Last recorded round trip time for a reliable packet.*/
	PeerLastRoundTripTime PeerStatistic = 5
	/*Variance of the last trip time recorded.*/
	PeerLastRoundTripTimeVariance PeerStatistic = 6
	/*The peer's current throttle status.*/
	PeerPacketThrottle PeerStatistic = 7
	/*The maximum number of unreliable packets that should not be dropped. This value is always greater than or equal to [code]1[/code]. The initial value is equal to [constant PACKET_THROTTLE_SCALE].*/
	PeerPacketThrottleLimit PeerStatistic = 8
	/*Internal value used to increment the packet throttle counter. The value is hardcoded to [code]7[/code] and cannot be changed. You probably want to look at [constant PEER_PACKET_THROTTLE_ACCELERATION] instead.*/
	PeerPacketThrottleCounter PeerStatistic = 9
	/*The time at which throttle statistics were last updated (in milliseconds since the connection started). The interval for throttle statistics updates is [constant PEER_PACKET_THROTTLE_INTERVAL].*/
	PeerPacketThrottleEpoch PeerStatistic = 10
	/*The throttle's acceleration factor. Higher values will make ENet adapt to fluctuating network conditions faster, causing unrelaible packets to be sent [i]more[/i] often. The default value is [code]2[/code].*/
	PeerPacketThrottleAcceleration PeerStatistic = 11
	/*The throttle's deceleration factor. Higher values will make ENet adapt to fluctuating network conditions faster, causing unrelaible packets to be sent [i]less[/i] often. The default value is [code]2[/code].*/
	PeerPacketThrottleDeceleration PeerStatistic = 12
	/*The interval over which the lowest mean round trip time should be measured for use by the throttle mechanism (in milliseconds). The default value is [code]5000[/code].*/
	PeerPacketThrottleInterval PeerStatistic = 13
)

type Error = gd.Error //gd:Error

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
