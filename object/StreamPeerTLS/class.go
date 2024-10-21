package StreamPeerTLS

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/StreamPeer"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A stream peer that handles TLS connections. This object can be used to connect to a TLS server or accept a single TLS client connection.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.

*/
type Simple [1]classdb.StreamPeerTLS
func (self Simple) Poll() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Poll()
}
func (self Simple) AcceptStream(stream [1]classdb.StreamPeer, server_options [1]classdb.TLSOptions) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).AcceptStream(stream, server_options))
}
func (self Simple) ConnectToStream(stream [1]classdb.StreamPeer, common_name string, client_options [1]classdb.TLSOptions) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).ConnectToStream(stream, gc.String(common_name), client_options))
}
func (self Simple) GetStatus() classdb.StreamPeerTLSStatus {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.StreamPeerTLSStatus(Expert(self).GetStatus())
}
func (self Simple) GetStream() [1]classdb.StreamPeer {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.StreamPeer(Expert(self).GetStream(gc))
}
func (self Simple) DisconnectFromStream() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DisconnectFromStream()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.StreamPeerTLS
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Poll the connection to check for incoming bytes. Call this right before [method StreamPeer.get_available_bytes] for it to work properly.
*/
//go:nosplit
func (self class) Poll()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerTLS.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Accepts a peer connection as a server using the given [param server_options]. See [method TLSOptions.server].
*/
//go:nosplit
func (self class) AcceptStream(stream object.StreamPeer, server_options object.TLSOptions) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(stream[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(server_options[0].AsPointer())[0])
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerTLS.Bind_accept_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Connects to a peer using an underlying [StreamPeer] [param stream] and verifying the remote certificate is correctly signed for the given [param common_name]. You can pass the optional [param client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
*/
//go:nosplit
func (self class) ConnectToStream(stream object.StreamPeer, common_name gd.String, client_options object.TLSOptions) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(stream[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(common_name))
	callframe.Arg(frame, mmm.Get(client_options[0].AsPointer())[0])
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerTLS.Bind_connect_to_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the status of the connection. See [enum Status] for values.
*/
//go:nosplit
func (self class) GetStatus() classdb.StreamPeerTLSStatus {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.StreamPeerTLSStatus](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerTLS.Bind_get_status, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the underlying [StreamPeer] connection, used in [method accept_stream] or [method connect_to_stream].
*/
//go:nosplit
func (self class) GetStream(ctx gd.Lifetime) object.StreamPeer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerTLS.Bind_get_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.StreamPeer
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Disconnects from host.
*/
//go:nosplit
func (self class) DisconnectFromStream()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerTLS.Bind_disconnect_from_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsStreamPeerTLS() Expert { return self[0].AsStreamPeerTLS() }


//go:nosplit
func (self Simple) AsStreamPeerTLS() Simple { return self[0].AsStreamPeerTLS() }


//go:nosplit
func (self class) AsStreamPeer() StreamPeer.Expert { return self[0].AsStreamPeer() }


//go:nosplit
func (self Simple) AsStreamPeer() StreamPeer.Simple { return self[0].AsStreamPeer() }


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
func init() {classdb.Register("StreamPeerTLS", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Status = classdb.StreamPeerTLSStatus

const (
/*A status representing a [StreamPeerTLS] that is disconnected.*/
	StatusDisconnected Status = 0
/*A status representing a [StreamPeerTLS] during handshaking.*/
	StatusHandshaking Status = 1
/*A status representing a [StreamPeerTLS] that is connected to a host.*/
	StatusConnected Status = 2
/*A status representing a [StreamPeerTLS] in error state.*/
	StatusError Status = 3
/*An error status that shows a mismatch in the TLS certificate domain presented by the host and the domain requested for validation.*/
	StatusErrorHostnameMismatch Status = 4
)
