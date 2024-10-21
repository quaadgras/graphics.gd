package StreamPeerTCP

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
A stream peer that handles TCP connections. This object can be used to connect to TCP servers, or also is returned by a TCP server.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.

*/
type Simple [1]classdb.StreamPeerTCP
func (self Simple) Bind(port int, host string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Bind(gd.Int(port), gc.String(host)))
}
func (self Simple) ConnectToHost(host string, port int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).ConnectToHost(gc.String(host), gd.Int(port)))
}
func (self Simple) Poll() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Poll())
}
func (self Simple) GetStatus() classdb.StreamPeerTCPStatus {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.StreamPeerTCPStatus(Expert(self).GetStatus())
}
func (self Simple) GetConnectedHost() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetConnectedHost(gc).String())
}
func (self Simple) GetConnectedPort() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetConnectedPort()))
}
func (self Simple) GetLocalPort() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetLocalPort()))
}
func (self Simple) DisconnectFromHost() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DisconnectFromHost()
}
func (self Simple) SetNoDelay(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNoDelay(enabled)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.StreamPeerTCP
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Opens the TCP socket, and binds it to the specified local address.
This method is generally not needed, and only used to force the subsequent call to [method connect_to_host] to use the specified [param host] and [param port] as source address. This can be desired in some NAT punchthrough techniques, or when forcing the source network interface.
*/
//go:nosplit
func (self class) Bind(port gd.Int, host gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, mmm.Get(host))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerTCP.Bind_bind, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Connects to the specified [code]host:port[/code] pair. A hostname will be resolved if valid. Returns [constant OK] on success.
*/
//go:nosplit
func (self class) ConnectToHost(host gd.String, port gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(host))
	callframe.Arg(frame, port)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerTCP.Bind_connect_to_host, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Poll the socket, updating its state. See [method get_status].
*/
//go:nosplit
func (self class) Poll() int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerTCP.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the status of the connection, see [enum Status].
*/
//go:nosplit
func (self class) GetStatus() classdb.StreamPeerTCPStatus {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.StreamPeerTCPStatus](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerTCP.Bind_get_status, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the IP of this peer.
*/
//go:nosplit
func (self class) GetConnectedHost(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerTCP.Bind_get_connected_host, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the port of this peer.
*/
//go:nosplit
func (self class) GetConnectedPort() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerTCP.Bind_get_connected_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the local port to which this peer is bound.
*/
//go:nosplit
func (self class) GetLocalPort() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerTCP.Bind_get_local_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Disconnects from host.
*/
//go:nosplit
func (self class) DisconnectFromHost()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerTCP.Bind_disconnect_from_host, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [param enabled] is [code]true[/code], packets will be sent immediately. If [param enabled] is [code]false[/code] (the default), packet transfers will be delayed and combined using [url=https://en.wikipedia.org/wiki/Nagle%27s_algorithm]Nagle's algorithm[/url].
[b]Note:[/b] It's recommended to leave this disabled for applications that send large packets or need to transfer a lot of data, as enabling this can decrease the total available bandwidth.
*/
//go:nosplit
func (self class) SetNoDelay(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerTCP.Bind_set_no_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsStreamPeerTCP() Expert { return self[0].AsStreamPeerTCP() }


//go:nosplit
func (self Simple) AsStreamPeerTCP() Simple { return self[0].AsStreamPeerTCP() }


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
func init() {classdb.Register("StreamPeerTCP", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Status = classdb.StreamPeerTCPStatus

const (
/*The initial status of the [StreamPeerTCP]. This is also the status after disconnecting.*/
	StatusNone Status = 0
/*A status representing a [StreamPeerTCP] that is connecting to a host.*/
	StatusConnecting Status = 1
/*A status representing a [StreamPeerTCP] that is connected to a host.*/
	StatusConnected Status = 2
/*A status representing a [StreamPeerTCP] in error state.*/
	StatusError Status = 3
)
