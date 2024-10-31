package TCPServer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
A TCP server. Listens to connections on a port and returns a [StreamPeerTCP] when it gets an incoming connection.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
*/
type Instance [1]classdb.TCPServer

/*
Listen on the [param port] binding to [param bind_address].
If [param bind_address] is set as [code]"*"[/code] (default), the server will listen on all available addresses (both IPv4 and IPv6).
If [param bind_address] is set as [code]"0.0.0.0"[/code] (for IPv4) or [code]"::"[/code] (for IPv6), the server will listen on all available addresses matching that IP type.
If [param bind_address] is set to any valid address (e.g. [code]"192.168.1.101"[/code], [code]"::1"[/code], etc.), the server will only listen on the interface with that address (or fail if no interface with the given address exists).
*/
func (self Instance) Listen(port int) gd.Error {
	return gd.Error(class(self).Listen(gd.Int(port), gd.NewString("*")))
}

/*
Returns [code]true[/code] if a connection is available for taking.
*/
func (self Instance) IsConnectionAvailable() bool {
	return bool(class(self).IsConnectionAvailable())
}

/*
Returns [code]true[/code] if the server is currently listening for connections.
*/
func (self Instance) IsListening() bool {
	return bool(class(self).IsListening())
}

/*
Returns the local port this server is listening to.
*/
func (self Instance) GetLocalPort() int {
	return int(int(class(self).GetLocalPort()))
}

/*
If a connection is available, returns a StreamPeerTCP with the connection.
*/
func (self Instance) TakeConnection() gdclass.StreamPeerTCP {
	return gdclass.StreamPeerTCP(class(self).TakeConnection())
}

/*
Stops listening.
*/
func (self Instance) Stop() {
	class(self).Stop()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.TCPServer

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TCPServer"))
	return Instance{classdb.TCPServer(object)}
}

/*
Listen on the [param port] binding to [param bind_address].
If [param bind_address] is set as [code]"*"[/code] (default), the server will listen on all available addresses (both IPv4 and IPv6).
If [param bind_address] is set as [code]"0.0.0.0"[/code] (for IPv4) or [code]"::"[/code] (for IPv6), the server will listen on all available addresses matching that IP type.
If [param bind_address] is set to any valid address (e.g. [code]"192.168.1.101"[/code], [code]"::1"[/code], etc.), the server will only listen on the interface with that address (or fail if no interface with the given address exists).
*/
//go:nosplit
func (self class) Listen(port gd.Int, bind_address gd.String) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, pointers.Get(bind_address))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TCPServer.Bind_listen, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if a connection is available for taking.
*/
//go:nosplit
func (self class) IsConnectionAvailable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TCPServer.Bind_is_connection_available, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the server is currently listening for connections.
*/
//go:nosplit
func (self class) IsListening() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TCPServer.Bind_is_listening, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the local port this server is listening to.
*/
//go:nosplit
func (self class) GetLocalPort() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TCPServer.Bind_get_local_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If a connection is available, returns a StreamPeerTCP with the connection.
*/
//go:nosplit
func (self class) TakeConnection() gdclass.StreamPeerTCP {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TCPServer.Bind_take_connection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.StreamPeerTCP{classdb.StreamPeerTCP(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Stops listening.
*/
//go:nosplit
func (self class) Stop() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TCPServer.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsTCPServer() Advanced          { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTCPServer() Instance       { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() { classdb.Register("TCPServer", func(ptr gd.Object) any { return classdb.TCPServer(ptr) }) }
