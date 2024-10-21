package MultiplayerPeer

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/PacketPeer"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Manages the connection with one or more remote peers acting as server or client and assigning unique IDs to each of them. See also [MultiplayerAPI].
[b]Note:[/b] The [MultiplayerAPI] protocol is an implementation detail and isn't meant to be used by non-Godot servers. It may change without notice.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.

*/
type Simple [1]classdb.MultiplayerPeer
func (self Simple) SetTransferChannel(channel int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTransferChannel(gd.Int(channel))
}
func (self Simple) GetTransferChannel() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTransferChannel()))
}
func (self Simple) SetTransferMode(mode classdb.MultiplayerPeerTransferMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTransferMode(mode)
}
func (self Simple) GetTransferMode() classdb.MultiplayerPeerTransferMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.MultiplayerPeerTransferMode(Expert(self).GetTransferMode())
}
func (self Simple) SetTargetPeer(id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTargetPeer(gd.Int(id))
}
func (self Simple) GetPacketPeer() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPacketPeer()))
}
func (self Simple) GetPacketChannel() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPacketChannel()))
}
func (self Simple) GetPacketMode() classdb.MultiplayerPeerTransferMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.MultiplayerPeerTransferMode(Expert(self).GetPacketMode())
}
func (self Simple) Poll() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Poll()
}
func (self Simple) Close() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Close()
}
func (self Simple) DisconnectPeer(peer int, force bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DisconnectPeer(gd.Int(peer), force)
}
func (self Simple) GetConnectionStatus() classdb.MultiplayerPeerConnectionStatus {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.MultiplayerPeerConnectionStatus(Expert(self).GetConnectionStatus())
}
func (self Simple) GetUniqueId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetUniqueId()))
}
func (self Simple) GenerateUniqueId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GenerateUniqueId()))
}
func (self Simple) SetRefuseNewConnections(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRefuseNewConnections(enable)
}
func (self Simple) IsRefusingNewConnections() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsRefusingNewConnections())
}
func (self Simple) IsServerRelaySupported() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsServerRelaySupported())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.MultiplayerPeer
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetTransferChannel(channel gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, channel)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_set_transfer_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTransferChannel() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_get_transfer_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTransferMode(mode classdb.MultiplayerPeerTransferMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_set_transfer_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTransferMode() classdb.MultiplayerPeerTransferMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.MultiplayerPeerTransferMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_get_transfer_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the peer to which packets will be sent.
The [param id] can be one of: [constant TARGET_PEER_BROADCAST] to send to all connected peers, [constant TARGET_PEER_SERVER] to send to the peer acting as server, a valid peer ID to send to that specific peer, a negative peer ID to send to all peers except that one. By default, the target peer is [constant TARGET_PEER_BROADCAST].
*/
//go:nosplit
func (self class) SetTargetPeer(id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_set_target_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the ID of the [MultiplayerPeer] who sent the next available packet. See [method PacketPeer.get_available_packet_count].
*/
//go:nosplit
func (self class) GetPacketPeer() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_get_packet_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the channel over which the next available packet was received. See [method PacketPeer.get_available_packet_count].
*/
//go:nosplit
func (self class) GetPacketChannel() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_get_packet_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the transfer mode the remote peer used to send the next available packet. See [method PacketPeer.get_available_packet_count].
*/
//go:nosplit
func (self class) GetPacketMode() classdb.MultiplayerPeerTransferMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.MultiplayerPeerTransferMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_get_packet_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Waits up to 1 second to receive a new network event.
*/
//go:nosplit
func (self class) Poll()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Immediately close the multiplayer peer returning to the state [constant CONNECTION_DISCONNECTED]. Connected peers will be dropped without emitting [signal peer_disconnected].
*/
//go:nosplit
func (self class) Close()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_close, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Disconnects the given [param peer] from this host. If [param force] is [code]true[/code] the [signal peer_disconnected] signal will not be emitted for this peer.
*/
//go:nosplit
func (self class) DisconnectPeer(peer gd.Int, force bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, peer)
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_disconnect_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the current state of the connection. See [enum ConnectionStatus].
*/
//go:nosplit
func (self class) GetConnectionStatus() classdb.MultiplayerPeerConnectionStatus {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.MultiplayerPeerConnectionStatus](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_get_connection_status, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the ID of this [MultiplayerPeer].
*/
//go:nosplit
func (self class) GetUniqueId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_get_unique_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a randomly generated integer that can be used as a network unique ID.
*/
//go:nosplit
func (self class) GenerateUniqueId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_generate_unique_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRefuseNewConnections(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_set_refuse_new_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsRefusingNewConnections() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_is_refusing_new_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns true if the server can act as a relay in the current configuration (i.e. if the higher level [MultiplayerAPI] should notify connected clients of other peers, and implement a relay protocol to allow communication between them).
*/
//go:nosplit
func (self class) IsServerRelaySupported() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_is_server_relay_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsMultiplayerPeer() Expert { return self[0].AsMultiplayerPeer() }


//go:nosplit
func (self Simple) AsMultiplayerPeer() Simple { return self[0].AsMultiplayerPeer() }


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
func init() {classdb.Register("MultiplayerPeer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type ConnectionStatus = classdb.MultiplayerPeerConnectionStatus

const (
/*The MultiplayerPeer is disconnected.*/
	ConnectionDisconnected ConnectionStatus = 0
/*The MultiplayerPeer is currently connecting to a server.*/
	ConnectionConnecting ConnectionStatus = 1
/*This MultiplayerPeer is connected.*/
	ConnectionConnected ConnectionStatus = 2
)
type TransferMode = classdb.MultiplayerPeerTransferMode

const (
/*Packets are not acknowledged, no resend attempts are made for lost packets. Packets may arrive in any order. Potentially faster than [constant TRANSFER_MODE_UNRELIABLE_ORDERED]. Use for non-critical data, and always consider whether the order matters.*/
	TransferModeUnreliable TransferMode = 0
/*Packets are not acknowledged, no resend attempts are made for lost packets. Packets are received in the order they were sent in. Potentially faster than [constant TRANSFER_MODE_RELIABLE]. Use for non-critical data or data that would be outdated if received late due to resend attempt(s) anyway, for example movement and positional data.*/
	TransferModeUnreliableOrdered TransferMode = 1
/*Packets must be received and resend attempts should be made until the packets are acknowledged. Packets must be received in the order they were sent in. Most reliable transfer mode, but potentially the slowest due to the overhead. Use for critical data that must be transmitted and arrive in order, for example an ability being triggered or a chat message. Consider carefully if the information really is critical, and use sparingly.*/
	TransferModeReliable TransferMode = 2
)
