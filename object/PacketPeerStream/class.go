package PacketPeerStream

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
PacketStreamPeer provides a wrapper for working using packets over a stream. This allows for using packet based code with StreamPeers. PacketPeerStream implements a custom protocol over the StreamPeer, so the user should not read or write to the wrapped StreamPeer directly.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.

*/
type Simple [1]classdb.PacketPeerStream
func (self Simple) SetStreamPeer(peer [1]classdb.StreamPeer) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStreamPeer(peer)
}
func (self Simple) GetStreamPeer() [1]classdb.StreamPeer {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.StreamPeer(Expert(self).GetStreamPeer(gc))
}
func (self Simple) SetInputBufferMaxSize(max_size_bytes int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInputBufferMaxSize(gd.Int(max_size_bytes))
}
func (self Simple) SetOutputBufferMaxSize(max_size_bytes int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOutputBufferMaxSize(gd.Int(max_size_bytes))
}
func (self Simple) GetInputBufferMaxSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetInputBufferMaxSize()))
}
func (self Simple) GetOutputBufferMaxSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetOutputBufferMaxSize()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PacketPeerStream
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetStreamPeer(peer object.StreamPeer)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(peer[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerStream.Bind_set_stream_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStreamPeer(ctx gd.Lifetime) object.StreamPeer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerStream.Bind_get_stream_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.StreamPeer
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInputBufferMaxSize(max_size_bytes gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_size_bytes)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerStream.Bind_set_input_buffer_max_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetOutputBufferMaxSize(max_size_bytes gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_size_bytes)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerStream.Bind_set_output_buffer_max_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInputBufferMaxSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerStream.Bind_get_input_buffer_max_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetOutputBufferMaxSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerStream.Bind_get_output_buffer_max_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsPacketPeerStream() Expert { return self[0].AsPacketPeerStream() }


//go:nosplit
func (self Simple) AsPacketPeerStream() Simple { return self[0].AsPacketPeerStream() }


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
func init() {classdb.Register("PacketPeerStream", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
