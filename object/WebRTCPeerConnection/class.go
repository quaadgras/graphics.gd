package WebRTCPeerConnection

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
A WebRTC connection between the local computer and a remote peer. Provides an interface to connect, maintain and monitor the connection.
Setting up a WebRTC connection between two peers may not seem a trivial task, but it can be broken down into 3 main steps:
- The peer that wants to initiate the connection ([code]A[/code] from now on) creates an offer and send it to the other peer ([code]B[/code] from now on).
- [code]B[/code] receives the offer, generate and answer, and sends it to [code]A[/code]).
- [code]A[/code] and [code]B[/code] then generates and exchange ICE candidates with each other.
After these steps, the connection should become connected. Keep on reading or look into the tutorial for more information.

*/
type Simple [1]classdb.WebRTCPeerConnection
func (self Simple) SetDefaultExtension(extension_class string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDefaultExtension(gc, gc.StringName(extension_class))
}
func (self Simple) Initialize(configuration gd.Dictionary) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Initialize(configuration))
}
func (self Simple) CreateDataChannel(label string, options gd.Dictionary) [1]classdb.WebRTCDataChannel {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.WebRTCDataChannel(Expert(self).CreateDataChannel(gc, gc.String(label), options))
}
func (self Simple) CreateOffer() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).CreateOffer())
}
func (self Simple) SetLocalDescription(atype string, sdp string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).SetLocalDescription(gc.String(atype), gc.String(sdp)))
}
func (self Simple) SetRemoteDescription(atype string, sdp string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).SetRemoteDescription(gc.String(atype), gc.String(sdp)))
}
func (self Simple) AddIceCandidate(media string, index int, name string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).AddIceCandidate(gc.String(media), gd.Int(index), gc.String(name)))
}
func (self Simple) Poll() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Poll())
}
func (self Simple) Close() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Close()
}
func (self Simple) GetConnectionState() classdb.WebRTCPeerConnectionConnectionState {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.WebRTCPeerConnectionConnectionState(Expert(self).GetConnectionState())
}
func (self Simple) GetGatheringState() classdb.WebRTCPeerConnectionGatheringState {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.WebRTCPeerConnectionGatheringState(Expert(self).GetGatheringState())
}
func (self Simple) GetSignalingState() classdb.WebRTCPeerConnectionSignalingState {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.WebRTCPeerConnectionSignalingState(Expert(self).GetSignalingState())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.WebRTCPeerConnection
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Sets the [param extension_class] as the default [WebRTCPeerConnectionExtension] returned when creating a new [WebRTCPeerConnection].
*/
//go:nosplit
func (self class) SetDefaultExtension(ctx gd.Lifetime, extension_class gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(extension_class))
	var r_ret callframe.Nil
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.WebRTCPeerConnection.Bind_set_default_extension, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) Initialize(configuration gd.Dictionary) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(configuration))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCPeerConnection.Bind_initialize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) CreateDataChannel(ctx gd.Lifetime, label gd.String, options gd.Dictionary) object.WebRTCDataChannel {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, mmm.Get(options))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCPeerConnection.Bind_create_data_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.WebRTCDataChannel
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a new SDP offer to start a WebRTC connection with a remote peer. At least one [WebRTCDataChannel] must have been created before calling this method.
If this functions returns [constant OK], [signal session_description_created] will be called when the session is ready to be sent.
*/
//go:nosplit
func (self class) CreateOffer() int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCPeerConnection.Bind_create_offer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the SDP description of the local peer. This should be called in response to [signal session_description_created].
After calling this function the peer will start emitting [signal ice_candidate_created] (unless an [enum Error] different from [constant OK] is returned).
*/
//go:nosplit
func (self class) SetLocalDescription(atype gd.String, sdp gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(atype))
	callframe.Arg(frame, mmm.Get(sdp))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCPeerConnection.Bind_set_local_description, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) SetRemoteDescription(atype gd.String, sdp gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(atype))
	callframe.Arg(frame, mmm.Get(sdp))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCPeerConnection.Bind_set_remote_description, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Add an ice candidate generated by a remote peer (and received over the signaling server). See [signal ice_candidate_created].
*/
//go:nosplit
func (self class) AddIceCandidate(media gd.String, index gd.Int, name gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(media))
	callframe.Arg(frame, index)
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCPeerConnection.Bind_add_ice_candidate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Call this method frequently (e.g. in [method Node._process] or [method Node._physics_process]) to properly receive signals.
*/
//go:nosplit
func (self class) Poll() int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCPeerConnection.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Close the peer connection and all data channels associated with it.
[b]Note:[/b] You cannot reuse this object for a new connection unless you call [method initialize].
*/
//go:nosplit
func (self class) Close()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCPeerConnection.Bind_close, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the connection state. See [enum ConnectionState].
*/
//go:nosplit
func (self class) GetConnectionState() classdb.WebRTCPeerConnectionConnectionState {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.WebRTCPeerConnectionConnectionState](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCPeerConnection.Bind_get_connection_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the ICE [enum GatheringState] of the connection. This lets you detect, for example, when collection of ICE candidates has finished.
*/
//go:nosplit
func (self class) GetGatheringState() classdb.WebRTCPeerConnectionGatheringState {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.WebRTCPeerConnectionGatheringState](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCPeerConnection.Bind_get_gathering_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the signaling state on the local end of the connection while connecting or reconnecting to another peer.
*/
//go:nosplit
func (self class) GetSignalingState() classdb.WebRTCPeerConnectionSignalingState {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.WebRTCPeerConnectionSignalingState](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCPeerConnection.Bind_get_signaling_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsWebRTCPeerConnection() Expert { return self[0].AsWebRTCPeerConnection() }


//go:nosplit
func (self Simple) AsWebRTCPeerConnection() Simple { return self[0].AsWebRTCPeerConnection() }


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
func init() {classdb.Register("WebRTCPeerConnection", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type ConnectionState = classdb.WebRTCPeerConnectionConnectionState

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
type GatheringState = classdb.WebRTCPeerConnectionGatheringState

const (
/*The peer connection was just created and hasn't done any networking yet.*/
	GatheringStateNew GatheringState = 0
/*The ICE agent is in the process of gathering candidates for the connection.*/
	GatheringStateGathering GatheringState = 1
/*The ICE agent has finished gathering candidates. If something happens that requires collecting new candidates, such as a new interface being added or the addition of a new ICE server, the state will revert to gathering to gather those candidates.*/
	GatheringStateComplete GatheringState = 2
)
type SignalingState = classdb.WebRTCPeerConnectionSignalingState

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
