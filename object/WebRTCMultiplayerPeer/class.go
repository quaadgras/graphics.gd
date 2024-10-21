package WebRTCMultiplayerPeer

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
This class constructs a full mesh of [WebRTCPeerConnection] (one connection for each peer) that can be used as a [member MultiplayerAPI.multiplayer_peer].
You can add each [WebRTCPeerConnection] via [method add_peer] or remove them via [method remove_peer]. Peers must be added in [constant WebRTCPeerConnection.STATE_NEW] state to allow it to create the appropriate channels. This class will not create offers nor set descriptions, it will only poll them, and notify connections and disconnections.
When creating the peer via [method create_client] or [method create_server] the [method MultiplayerPeer.is_server_relay_supported] method will return [code]true[/code] enabling peer exchange and packet relaying when supported by the [MultiplayerAPI] implementation.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.

*/
type Simple [1]classdb.WebRTCMultiplayerPeer
func (self Simple) CreateServer(channels_config gd.Array) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).CreateServer(channels_config))
}
func (self Simple) CreateClient(peer_id int, channels_config gd.Array) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).CreateClient(gd.Int(peer_id), channels_config))
}
func (self Simple) CreateMesh(peer_id int, channels_config gd.Array) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).CreateMesh(gd.Int(peer_id), channels_config))
}
func (self Simple) AddPeer(peer [1]classdb.WebRTCPeerConnection, peer_id int, unreliable_lifetime int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).AddPeer(peer, gd.Int(peer_id), gd.Int(unreliable_lifetime)))
}
func (self Simple) RemovePeer(peer_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemovePeer(gd.Int(peer_id))
}
func (self Simple) HasPeer(peer_id int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasPeer(gd.Int(peer_id)))
}
func (self Simple) GetPeer(peer_id int) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).GetPeer(gc, gd.Int(peer_id)))
}
func (self Simple) GetPeers() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).GetPeers(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.WebRTCMultiplayerPeer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Initialize the multiplayer peer as a server (with unique ID of [code]1[/code]). This mode enables [method MultiplayerPeer.is_server_relay_supported], allowing the upper [MultiplayerAPI] layer to perform peer exchange and packet relaying.
You can optionally specify a [param channels_config] array of [enum MultiplayerPeer.TransferMode] which will be used to create extra channels (WebRTC only supports one transfer mode per channel).
*/
//go:nosplit
func (self class) CreateServer(channels_config gd.Array) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(channels_config))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCMultiplayerPeer.Bind_create_server, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Initialize the multiplayer peer as a client with the given [param peer_id] (must be between 2 and 2147483647). In this mode, you should only call [method add_peer] once and with [param peer_id] of [code]1[/code]. This mode enables [method MultiplayerPeer.is_server_relay_supported], allowing the upper [MultiplayerAPI] layer to perform peer exchange and packet relaying.
You can optionally specify a [param channels_config] array of [enum MultiplayerPeer.TransferMode] which will be used to create extra channels (WebRTC only supports one transfer mode per channel).
*/
//go:nosplit
func (self class) CreateClient(peer_id gd.Int, channels_config gd.Array) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	callframe.Arg(frame, mmm.Get(channels_config))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCMultiplayerPeer.Bind_create_client, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Initialize the multiplayer peer as a mesh (i.e. all peers connect to each other) with the given [param peer_id] (must be between 1 and 2147483647).
*/
//go:nosplit
func (self class) CreateMesh(peer_id gd.Int, channels_config gd.Array) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	callframe.Arg(frame, mmm.Get(channels_config))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCMultiplayerPeer.Bind_create_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Add a new peer to the mesh with the given [param peer_id]. The [WebRTCPeerConnection] must be in state [constant WebRTCPeerConnection.STATE_NEW].
Three channels will be created for reliable, unreliable, and ordered transport. The value of [param unreliable_lifetime] will be passed to the [code]"maxPacketLifetime"[/code] option when creating unreliable and ordered channels (see [method WebRTCPeerConnection.create_data_channel]).
*/
//go:nosplit
func (self class) AddPeer(peer object.WebRTCPeerConnection, peer_id gd.Int, unreliable_lifetime gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(peer[0].AsPointer())[0])
	callframe.Arg(frame, peer_id)
	callframe.Arg(frame, unreliable_lifetime)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCMultiplayerPeer.Bind_add_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Remove the peer with given [param peer_id] from the mesh. If the peer was connected, and [signal MultiplayerPeer.peer_connected] was emitted for it, then [signal MultiplayerPeer.peer_disconnected] will be emitted.
*/
//go:nosplit
func (self class) RemovePeer(peer_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCMultiplayerPeer.Bind_remove_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the given [param peer_id] is in the peers map (it might not be connected though).
*/
//go:nosplit
func (self class) HasPeer(peer_id gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCMultiplayerPeer.Bind_has_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a dictionary representation of the peer with given [param peer_id] with three keys. [code]"connection"[/code] containing the [WebRTCPeerConnection] to this peer, [code]"channels"[/code] an array of three [WebRTCDataChannel], and [code]"connected"[/code] a boolean representing if the peer connection is currently connected (all three channels are open).
*/
//go:nosplit
func (self class) GetPeer(ctx gd.Lifetime, peer_id gd.Int) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCMultiplayerPeer.Bind_get_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a dictionary which keys are the peer ids and values the peer representation as in [method get_peer].
*/
//go:nosplit
func (self class) GetPeers(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCMultiplayerPeer.Bind_get_peers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsWebRTCMultiplayerPeer() Expert { return self[0].AsWebRTCMultiplayerPeer() }


//go:nosplit
func (self Simple) AsWebRTCMultiplayerPeer() Simple { return self[0].AsWebRTCMultiplayerPeer() }


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
func init() {classdb.Register("WebRTCMultiplayerPeer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
