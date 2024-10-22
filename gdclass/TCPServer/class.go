package TCPServer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A TCP server. Listens to connections on a port and returns a [StreamPeerTCP] when it gets an incoming connection.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.

*/
type Go [1]classdb.TCPServer

/*
Listen on the [param port] binding to [param bind_address].
If [param bind_address] is set as [code]"*"[/code] (default), the server will listen on all available addresses (both IPv4 and IPv6).
If [param bind_address] is set as [code]"0.0.0.0"[/code] (for IPv4) or [code]"::"[/code] (for IPv6), the server will listen on all available addresses matching that IP type.
If [param bind_address] is set to any valid address (e.g. [code]"192.168.1.101"[/code], [code]"::1"[/code], etc.), the server will only listen on the interface with that address (or fail if no interface with the given address exists).
*/
func (self Go) Listen(port int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Listen(gd.Int(port), gc.String("*")))
}

/*
Returns [code]true[/code] if a connection is available for taking.
*/
func (self Go) IsConnectionAvailable() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsConnectionAvailable())
}

/*
Returns [code]true[/code] if the server is currently listening for connections.
*/
func (self Go) IsListening() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsListening())
}

/*
Returns the local port this server is listening to.
*/
func (self Go) GetLocalPort() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetLocalPort()))
}

/*
If a connection is available, returns a StreamPeerTCP with the connection.
*/
func (self Go) TakeConnection() gdclass.StreamPeerTCP {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.StreamPeerTCP(class(self).TakeConnection(gc))
}

/*
Stops listening.
*/
func (self Go) Stop() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Stop()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TCPServer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("TCPServer"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Listen on the [param port] binding to [param bind_address].
If [param bind_address] is set as [code]"*"[/code] (default), the server will listen on all available addresses (both IPv4 and IPv6).
If [param bind_address] is set as [code]"0.0.0.0"[/code] (for IPv4) or [code]"::"[/code] (for IPv6), the server will listen on all available addresses matching that IP type.
If [param bind_address] is set to any valid address (e.g. [code]"192.168.1.101"[/code], [code]"::1"[/code], etc.), the server will only listen on the interface with that address (or fail if no interface with the given address exists).
*/
//go:nosplit
func (self class) Listen(port gd.Int, bind_address gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, mmm.Get(bind_address))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TCPServer.Bind_listen, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if a connection is available for taking.
*/
//go:nosplit
func (self class) IsConnectionAvailable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TCPServer.Bind_is_connection_available, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the server is currently listening for connections.
*/
//go:nosplit
func (self class) IsListening() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TCPServer.Bind_is_listening, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the local port this server is listening to.
*/
//go:nosplit
func (self class) GetLocalPort() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TCPServer.Bind_get_local_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If a connection is available, returns a StreamPeerTCP with the connection.
*/
//go:nosplit
func (self class) TakeConnection(ctx gd.Lifetime) gdclass.StreamPeerTCP {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TCPServer.Bind_take_connection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.StreamPeerTCP
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Stops listening.
*/
//go:nosplit
func (self class) Stop()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TCPServer.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsTCPServer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTCPServer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("TCPServer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
