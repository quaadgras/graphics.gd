package PacketPeerDTLS

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
This class represents a DTLS peer connection. It can be used to connect to a DTLS server, and is returned by [method DTLSServer.take_connection].
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
[b]Warning:[/b] TLS certificate revocation and certificate pinning are currently not supported. Revoked certificates are accepted as long as they are otherwise valid. If this is a concern, you may want to use automatically managed certificates with a short validity period.

*/
type Simple [1]classdb.PacketPeerDTLS
func (self Simple) Poll() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Poll()
}
func (self Simple) ConnectToPeer(packet_peer [1]classdb.PacketPeerUDP, hostname string, client_options [1]classdb.TLSOptions) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).ConnectToPeer(packet_peer, gc.String(hostname), client_options))
}
func (self Simple) GetStatus() classdb.PacketPeerDTLSStatus {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.PacketPeerDTLSStatus(Expert(self).GetStatus())
}
func (self Simple) DisconnectFromPeer() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DisconnectFromPeer()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PacketPeerDTLS
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Poll the connection to check for incoming packets. Call this frequently to update the status and keep the connection working.
*/
//go:nosplit
func (self class) Poll()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerDTLS.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Connects a [param packet_peer] beginning the DTLS handshake using the underlying [PacketPeerUDP] which must be connected (see [method PacketPeerUDP.connect_to_host]). You can optionally specify the [param client_options] to be used while verifying the TLS connections. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
*/
//go:nosplit
func (self class) ConnectToPeer(packet_peer object.PacketPeerUDP, hostname gd.String, client_options object.TLSOptions) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(packet_peer[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(hostname))
	callframe.Arg(frame, mmm.Get(client_options[0].AsPointer())[0])
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerDTLS.Bind_connect_to_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the status of the connection. See [enum Status] for values.
*/
//go:nosplit
func (self class) GetStatus() classdb.PacketPeerDTLSStatus {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.PacketPeerDTLSStatus](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerDTLS.Bind_get_status, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Disconnects this peer, terminating the DTLS session.
*/
//go:nosplit
func (self class) DisconnectFromPeer()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerDTLS.Bind_disconnect_from_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsPacketPeerDTLS() Expert { return self[0].AsPacketPeerDTLS() }


//go:nosplit
func (self Simple) AsPacketPeerDTLS() Simple { return self[0].AsPacketPeerDTLS() }


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
func init() {classdb.Register("PacketPeerDTLS", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Status = classdb.PacketPeerDTLSStatus

const (
/*A status representing a [PacketPeerDTLS] that is disconnected.*/
	StatusDisconnected Status = 0
/*A status representing a [PacketPeerDTLS] that is currently performing the handshake with a remote peer.*/
	StatusHandshaking Status = 1
/*A status representing a [PacketPeerDTLS] that is connected to a remote peer.*/
	StatusConnected Status = 2
/*A status representing a [PacketPeerDTLS] in a generic error state.*/
	StatusError Status = 3
/*An error status that shows a mismatch in the DTLS certificate domain presented by the host and the domain requested for validation.*/
	StatusErrorHostnameMismatch Status = 4
)
