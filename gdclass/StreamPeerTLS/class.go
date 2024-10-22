package StreamPeerTLS

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/StreamPeer"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A stream peer that handles TLS connections. This object can be used to connect to a TLS server or accept a single TLS client connection.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.

*/
type Go [1]classdb.StreamPeerTLS

/*
Poll the connection to check for incoming bytes. Call this right before [method StreamPeer.get_available_bytes] for it to work properly.
*/
func (self Go) Poll() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Poll()
}

/*
Accepts a peer connection as a server using the given [param server_options]. See [method TLSOptions.server].
*/
func (self Go) AcceptStream(stream gdclass.StreamPeer, server_options gdclass.TLSOptions) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).AcceptStream(stream, server_options))
}

/*
Connects to a peer using an underlying [StreamPeer] [param stream] and verifying the remote certificate is correctly signed for the given [param common_name]. You can pass the optional [param client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
*/
func (self Go) ConnectToStream(stream gdclass.StreamPeer, common_name string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).ConnectToStream(stream, gc.String(common_name), ([1]gdclass.TLSOptions{}[0])))
}

/*
Returns the status of the connection. See [enum Status] for values.
*/
func (self Go) GetStatus() classdb.StreamPeerTLSStatus {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.StreamPeerTLSStatus(class(self).GetStatus())
}

/*
Returns the underlying [StreamPeer] connection, used in [method accept_stream] or [method connect_to_stream].
*/
func (self Go) GetStream() gdclass.StreamPeer {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.StreamPeer(class(self).GetStream(gc))
}

/*
Disconnects from host.
*/
func (self Go) DisconnectFromStream() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DisconnectFromStream()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.StreamPeerTLS
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("StreamPeerTLS"))
	return *(*Go)(unsafe.Pointer(&object))
}

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
func (self class) AcceptStream(stream gdclass.StreamPeer, server_options gdclass.TLSOptions) int64 {
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
func (self class) ConnectToStream(stream gdclass.StreamPeer, common_name gd.String, client_options gdclass.TLSOptions) int64 {
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
func (self class) GetStream(ctx gd.Lifetime) gdclass.StreamPeer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerTLS.Bind_get_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.StreamPeer
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
func (self class) AsStreamPeerTLS() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsStreamPeerTLS() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsStreamPeer() StreamPeer.GD { return *((*StreamPeer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsStreamPeer() StreamPeer.Go { return *((*StreamPeer.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsStreamPeer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsStreamPeer(), name)
	}
}
func init() {classdb.Register("StreamPeerTLS", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
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
