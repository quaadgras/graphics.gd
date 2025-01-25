// Package WebRTCPeerConnection provides methods for working with WebRTCPeerConnection object instances.
package WebRTCPeerConnection

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

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
A WebRTC connection between the local computer and a remote peer. Provides an interface to connect, maintain and monitor the connection.
Setting up a WebRTC connection between two peers may not seem a trivial task, but it can be broken down into 3 main steps:
- The peer that wants to initiate the connection ([code]A[/code] from now on) creates an offer and send it to the other peer ([code]B[/code] from now on).
- [code]B[/code] receives the offer, generate and answer, and sends it to [code]A[/code]).
- [code]A[/code] and [code]B[/code] then generates and exchange ICE candidates with each other.
After these steps, the connection should become connected. Keep on reading or look into the tutorial for more information.
*/
type Instance [1]gdclass.WebRTCPeerConnection

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsWebRTCPeerConnection() Instance
}

/*
Sets the [param extension_class] as the default [WebRTCPeerConnectionExtension] returned when creating a new [WebRTCPeerConnection].
*/
func SetDefaultExtension(extension_class string) { //gd:WebRTCPeerConnection.set_default_extension
	self := Instance{}
	class(self).SetDefaultExtension(gd.NewStringName(extension_class))
}

/*
Re-initialize this peer connection, closing any previously active connection, and going back to state [constant STATE_NEW]. A dictionary of [param configuration] options can be passed to configure the peer connection.
Valid [param configuration] options are:
[codeblock]

	{
	    "iceServers": [
	        {
	            "urls": [ "stun:stun.example.com:3478" ], # One or more STUN servers.
	        },
	        {
	            "urls": [ "turn:turn.example.com:3478" ], # One or more TURN servers.
	            "username": "a_username", # Optional username for the TURN server.
	            "credential": "a_password", # Optional password for the TURN server.
	        }
	    ]
	}

[/codeblock]
*/
func (self Instance) Initialize() error { //gd:WebRTCPeerConnection.initialize
	return error(gd.ToError(class(self).Initialize(gd.NewVariant([1]map[any]any{}[0]).Interface().(gd.Dictionary))))
}

/*
Returns a new [WebRTCDataChannel] (or [code]null[/code] on failure) with given [param label] and optionally configured via the [param options] dictionary. This method can only be called when the connection is in state [constant STATE_NEW].
There are two ways to create a working data channel: either call [method create_data_channel] on only one of the peer and listen to [signal data_channel_received] on the other, or call [method create_data_channel] on both peers, with the same values, and the [code]"negotiated"[/code] option set to [code]true[/code].
Valid [param options] are:
[codeblock]

	{
	    "negotiated": true, # When set to true (default off), means the channel is negotiated out of band. "id" must be set too. "data_channel_received" will not be called.
	    "id": 1, # When "negotiated" is true this value must also be set to the same value on both peer.

	    # Only one of maxRetransmits and maxPacketLifeTime can be specified, not both. They make the channel unreliable (but also better at real time).
	    "maxRetransmits": 1, # Specify the maximum number of attempt the peer will make to retransmits packets if they are not acknowledged.
	    "maxPacketLifeTime": 100, # Specify the maximum amount of time before giving up retransmitions of unacknowledged packets (in milliseconds).
	    "ordered": true, # When in unreliable mode (i.e. either "maxRetransmits" or "maxPacketLifetime" is set), "ordered" (true by default) specify if packet ordering is to be enforced.

	    "protocol": "my-custom-protocol", # A custom sub-protocol string for this channel.
	}

[/codeblock]
[b]Note:[/b] You must keep a reference to channels created this way, or it will be closed.
*/
func (self Instance) CreateDataChannel(label string) [1]gdclass.WebRTCDataChannel { //gd:WebRTCPeerConnection.create_data_channel
	return [1]gdclass.WebRTCDataChannel(class(self).CreateDataChannel(gd.NewString(label), gd.NewVariant([1]map[any]any{}[0]).Interface().(gd.Dictionary)))
}

/*
Creates a new SDP offer to start a WebRTC connection with a remote peer. At least one [WebRTCDataChannel] must have been created before calling this method.
If this functions returns [constant OK], [signal session_description_created] will be called when the session is ready to be sent.
*/
func (self Instance) CreateOffer() error { //gd:WebRTCPeerConnection.create_offer
	return error(gd.ToError(class(self).CreateOffer()))
}

/*
Sets the SDP description of the local peer. This should be called in response to [signal session_description_created].
After calling this function the peer will start emitting [signal ice_candidate_created] (unless an [enum Error] different from [constant OK] is returned).
*/
func (self Instance) SetLocalDescription(atype string, sdp string) error { //gd:WebRTCPeerConnection.set_local_description
	return error(gd.ToError(class(self).SetLocalDescription(gd.NewString(atype), gd.NewString(sdp))))
}

/*
Sets the SDP description of the remote peer. This should be called with the values generated by a remote peer and received over the signaling server.
If [param type] is [code]"offer"[/code] the peer will emit [signal session_description_created] with the appropriate answer.
If [param type] is [code]"answer"[/code] the peer will start emitting [signal ice_candidate_created].
*/
func (self Instance) SetRemoteDescription(atype string, sdp string) error { //gd:WebRTCPeerConnection.set_remote_description
	return error(gd.ToError(class(self).SetRemoteDescription(gd.NewString(atype), gd.NewString(sdp))))
}

/*
Add an ice candidate generated by a remote peer (and received over the signaling server). See [signal ice_candidate_created].
*/
func (self Instance) AddIceCandidate(media string, index int, name string) error { //gd:WebRTCPeerConnection.add_ice_candidate
	return error(gd.ToError(class(self).AddIceCandidate(gd.NewString(media), gd.Int(index), gd.NewString(name))))
}

/*
Call this method frequently (e.g. in [method Node._process] or [method Node._physics_process]) to properly receive signals.
*/
func (self Instance) Poll() error { //gd:WebRTCPeerConnection.poll
	return error(gd.ToError(class(self).Poll()))
}

/*
Close the peer connection and all data channels associated with it.
[b]Note:[/b] You cannot reuse this object for a new connection unless you call [method initialize].
*/
func (self Instance) Close() { //gd:WebRTCPeerConnection.close
	class(self).Close()
}

/*
Returns the connection state. See [enum ConnectionState].
*/
func (self Instance) GetConnectionState() gdclass.WebRTCPeerConnectionConnectionState { //gd:WebRTCPeerConnection.get_connection_state
	return gdclass.WebRTCPeerConnectionConnectionState(class(self).GetConnectionState())
}

/*
Returns the ICE [enum GatheringState] of the connection. This lets you detect, for example, when collection of ICE candidates has finished.
*/
func (self Instance) GetGatheringState() gdclass.WebRTCPeerConnectionGatheringState { //gd:WebRTCPeerConnection.get_gathering_state
	return gdclass.WebRTCPeerConnectionGatheringState(class(self).GetGatheringState())
}

/*
Returns the signaling state on the local end of the connection while connecting or reconnecting to another peer.
*/
func (self Instance) GetSignalingState() gdclass.WebRTCPeerConnectionSignalingState { //gd:WebRTCPeerConnection.get_signaling_state
	return gdclass.WebRTCPeerConnectionSignalingState(class(self).GetSignalingState())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.WebRTCPeerConnection

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WebRTCPeerConnection"))
	casted := Instance{*(*gdclass.WebRTCPeerConnection)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Sets the [param extension_class] as the default [WebRTCPeerConnectionExtension] returned when creating a new [WebRTCPeerConnection].
*/
//go:nosplit
func (self class) SetDefaultExtension(extension_class gd.StringName) { //gd:WebRTCPeerConnection.set_default_extension
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension_class))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCPeerConnection.Bind_set_default_extension, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Re-initialize this peer connection, closing any previously active connection, and going back to state [constant STATE_NEW]. A dictionary of [param configuration] options can be passed to configure the peer connection.
Valid [param configuration] options are:
[codeblock]
{
    "iceServers": [
        {
            "urls": [ "stun:stun.example.com:3478" ], # One or more STUN servers.
        },
        {
            "urls": [ "turn:turn.example.com:3478" ], # One or more TURN servers.
            "username": "a_username", # Optional username for the TURN server.
            "credential": "a_password", # Optional password for the TURN server.
        }
    ]
}
[/codeblock]
*/
//go:nosplit
func (self class) Initialize(configuration gd.Dictionary) gd.Error { //gd:WebRTCPeerConnection.initialize
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(configuration))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCPeerConnection.Bind_initialize, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a new [WebRTCDataChannel] (or [code]null[/code] on failure) with given [param label] and optionally configured via the [param options] dictionary. This method can only be called when the connection is in state [constant STATE_NEW].
There are two ways to create a working data channel: either call [method create_data_channel] on only one of the peer and listen to [signal data_channel_received] on the other, or call [method create_data_channel] on both peers, with the same values, and the [code]"negotiated"[/code] option set to [code]true[/code].
Valid [param options] are:
[codeblock]
{
    "negotiated": true, # When set to true (default off), means the channel is negotiated out of band. "id" must be set too. "data_channel_received" will not be called.
    "id": 1, # When "negotiated" is true this value must also be set to the same value on both peer.

    # Only one of maxRetransmits and maxPacketLifeTime can be specified, not both. They make the channel unreliable (but also better at real time).
    "maxRetransmits": 1, # Specify the maximum number of attempt the peer will make to retransmits packets if they are not acknowledged.
    "maxPacketLifeTime": 100, # Specify the maximum amount of time before giving up retransmitions of unacknowledged packets (in milliseconds).
    "ordered": true, # When in unreliable mode (i.e. either "maxRetransmits" or "maxPacketLifetime" is set), "ordered" (true by default) specify if packet ordering is to be enforced.

    "protocol": "my-custom-protocol", # A custom sub-protocol string for this channel.
}
[/codeblock]
[b]Note:[/b] You must keep a reference to channels created this way, or it will be closed.
*/
//go:nosplit
func (self class) CreateDataChannel(label gd.String, options gd.Dictionary) [1]gdclass.WebRTCDataChannel { //gd:WebRTCPeerConnection.create_data_channel
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, pointers.Get(options))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCPeerConnection.Bind_create_data_channel, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.WebRTCDataChannel{gd.PointerWithOwnershipTransferredToGo[gdclass.WebRTCDataChannel](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates a new SDP offer to start a WebRTC connection with a remote peer. At least one [WebRTCDataChannel] must have been created before calling this method.
If this functions returns [constant OK], [signal session_description_created] will be called when the session is ready to be sent.
*/
//go:nosplit
func (self class) CreateOffer() gd.Error { //gd:WebRTCPeerConnection.create_offer
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCPeerConnection.Bind_create_offer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the SDP description of the local peer. This should be called in response to [signal session_description_created].
After calling this function the peer will start emitting [signal ice_candidate_created] (unless an [enum Error] different from [constant OK] is returned).
*/
//go:nosplit
func (self class) SetLocalDescription(atype gd.String, sdp gd.String) gd.Error { //gd:WebRTCPeerConnection.set_local_description
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(atype))
	callframe.Arg(frame, pointers.Get(sdp))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCPeerConnection.Bind_set_local_description, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the SDP description of the remote peer. This should be called with the values generated by a remote peer and received over the signaling server.
If [param type] is [code]"offer"[/code] the peer will emit [signal session_description_created] with the appropriate answer.
If [param type] is [code]"answer"[/code] the peer will start emitting [signal ice_candidate_created].
*/
//go:nosplit
func (self class) SetRemoteDescription(atype gd.String, sdp gd.String) gd.Error { //gd:WebRTCPeerConnection.set_remote_description
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(atype))
	callframe.Arg(frame, pointers.Get(sdp))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCPeerConnection.Bind_set_remote_description, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Add an ice candidate generated by a remote peer (and received over the signaling server). See [signal ice_candidate_created].
*/
//go:nosplit
func (self class) AddIceCandidate(media gd.String, index gd.Int, name gd.String) gd.Error { //gd:WebRTCPeerConnection.add_ice_candidate
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(media))
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCPeerConnection.Bind_add_ice_candidate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Call this method frequently (e.g. in [method Node._process] or [method Node._physics_process]) to properly receive signals.
*/
//go:nosplit
func (self class) Poll() gd.Error { //gd:WebRTCPeerConnection.poll
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCPeerConnection.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Close the peer connection and all data channels associated with it.
[b]Note:[/b] You cannot reuse this object for a new connection unless you call [method initialize].
*/
//go:nosplit
func (self class) Close() { //gd:WebRTCPeerConnection.close
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCPeerConnection.Bind_close, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the connection state. See [enum ConnectionState].
*/
//go:nosplit
func (self class) GetConnectionState() gdclass.WebRTCPeerConnectionConnectionState { //gd:WebRTCPeerConnection.get_connection_state
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.WebRTCPeerConnectionConnectionState](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCPeerConnection.Bind_get_connection_state, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the ICE [enum GatheringState] of the connection. This lets you detect, for example, when collection of ICE candidates has finished.
*/
//go:nosplit
func (self class) GetGatheringState() gdclass.WebRTCPeerConnectionGatheringState { //gd:WebRTCPeerConnection.get_gathering_state
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.WebRTCPeerConnectionGatheringState](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCPeerConnection.Bind_get_gathering_state, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the signaling state on the local end of the connection while connecting or reconnecting to another peer.
*/
//go:nosplit
func (self class) GetSignalingState() gdclass.WebRTCPeerConnectionSignalingState { //gd:WebRTCPeerConnection.get_signaling_state
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.WebRTCPeerConnectionSignalingState](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCPeerConnection.Bind_get_signaling_state, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnSessionDescriptionCreated(cb func(atype string, sdp string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("session_description_created"), gd.NewCallable(cb), 0)
}

func (self Instance) OnIceCandidateCreated(cb func(media string, index int, name string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("ice_candidate_created"), gd.NewCallable(cb), 0)
}

func (self Instance) OnDataChannelReceived(cb func(channel [1]gdclass.WebRTCDataChannel)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("data_channel_received"), gd.NewCallable(cb), 0)
}

func (self class) AsWebRTCPeerConnection() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWebRTCPeerConnection() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("WebRTCPeerConnection", func(ptr gd.Object) any {
		return [1]gdclass.WebRTCPeerConnection{*(*gdclass.WebRTCPeerConnection)(unsafe.Pointer(&ptr))}
	})
}

type ConnectionState = gdclass.WebRTCPeerConnectionConnectionState //gd:WebRTCPeerConnection.ConnectionState

const (
	/*The connection is new, data channels and an offer can be created in this state.*/
	StateNew ConnectionState = 0
	/*The peer is connecting, ICE is in progress, none of the transports has failed.*/
	StateConnecting ConnectionState = 1
	/*The peer is connected, all ICE transports are connected.*/
	StateConnected ConnectionState = 2
	/*At least one ICE transport is disconnected.*/
	StateDisconnected ConnectionState = 3
	/*One or more of the ICE transports failed.*/
	StateFailed ConnectionState = 4
	/*The peer connection is closed (after calling [method close] for example).*/
	StateClosed ConnectionState = 5
)

type GatheringState = gdclass.WebRTCPeerConnectionGatheringState //gd:WebRTCPeerConnection.GatheringState

const (
	/*The peer connection was just created and hasn't done any networking yet.*/
	GatheringStateNew GatheringState = 0
	/*The ICE agent is in the process of gathering candidates for the connection.*/
	GatheringStateGathering GatheringState = 1
	/*The ICE agent has finished gathering candidates. If something happens that requires collecting new candidates, such as a new interface being added or the addition of a new ICE server, the state will revert to gathering to gather those candidates.*/
	GatheringStateComplete GatheringState = 2
)

type SignalingState = gdclass.WebRTCPeerConnectionSignalingState //gd:WebRTCPeerConnection.SignalingState

const (
	/*There is no ongoing exchange of offer and answer underway. This may mean that the [WebRTCPeerConnection] is new ([constant STATE_NEW]) or that negotiation is complete and a connection has been established ([constant STATE_CONNECTED]).*/
	SignalingStateStable SignalingState = 0
	/*The local peer has called [method set_local_description], passing in SDP representing an offer (usually created by calling [method create_offer]), and the offer has been applied successfully.*/
	SignalingStateHaveLocalOffer SignalingState = 1
	/*The remote peer has created an offer and used the signaling server to deliver it to the local peer, which has set the offer as the remote description by calling [method set_remote_description].*/
	SignalingStateHaveRemoteOffer SignalingState = 2
	/*The offer sent by the remote peer has been applied and an answer has been created and applied by calling [method set_local_description]. This provisional answer describes the supported media formats and so forth, but may not have a complete set of ICE candidates included. Further candidates will be delivered separately later.*/
	SignalingStateHaveLocalPranswer SignalingState = 3
	/*A provisional answer has been received and successfully applied in response to an offer previously sent and established by calling [method set_local_description].*/
	SignalingStateHaveRemotePranswer SignalingState = 4
	/*The [WebRTCPeerConnection] has been closed.*/
	SignalingStateClosed SignalingState = 5
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
