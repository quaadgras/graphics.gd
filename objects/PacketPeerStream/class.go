package PacketPeerStream

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/PacketPeer"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
PacketStreamPeer provides a wrapper for working using packets over a stream. This allows for using packet based code with StreamPeers. PacketPeerStream implements a custom protocol over the StreamPeer, so the user should not read or write to the wrapped StreamPeer directly.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
*/
type Instance [1]classdb.PacketPeerStream

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.PacketPeerStream

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PacketPeerStream"))
	return Instance{classdb.PacketPeerStream(object)}
}

func (self Instance) InputBufferMaxSize() int {
	return int(int(class(self).GetInputBufferMaxSize()))
}

func (self Instance) SetInputBufferMaxSize(value int) {
	class(self).SetInputBufferMaxSize(gd.Int(value))
}

func (self Instance) OutputBufferMaxSize() int {
	return int(int(class(self).GetOutputBufferMaxSize()))
}

func (self Instance) SetOutputBufferMaxSize(value int) {
	class(self).SetOutputBufferMaxSize(gd.Int(value))
}

func (self Instance) StreamPeer() objects.StreamPeer {
	return objects.StreamPeer(class(self).GetStreamPeer())
}

func (self Instance) SetStreamPeer(value objects.StreamPeer) {
	class(self).SetStreamPeer(value)
}

//go:nosplit
func (self class) SetStreamPeer(peer objects.StreamPeer) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(peer[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerStream.Bind_set_stream_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetStreamPeer() objects.StreamPeer {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerStream.Bind_get_stream_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.StreamPeer{classdb.StreamPeer(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInputBufferMaxSize(max_size_bytes gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, max_size_bytes)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerStream.Bind_set_input_buffer_max_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetOutputBufferMaxSize(max_size_bytes gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, max_size_bytes)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerStream.Bind_set_output_buffer_max_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetInputBufferMaxSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerStream.Bind_get_input_buffer_max_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetOutputBufferMaxSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerStream.Bind_get_output_buffer_max_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPacketPeerStream() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPacketPeerStream() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPacketPeer() PacketPeer.Advanced {
	return *((*PacketPeer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPacketPeer() PacketPeer.Instance {
	return *((*PacketPeer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsPacketPeer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsPacketPeer(), name)
	}
}
func init() {
	classdb.Register("PacketPeerStream", func(ptr gd.Object) any { return classdb.PacketPeerStream(ptr) })
}
