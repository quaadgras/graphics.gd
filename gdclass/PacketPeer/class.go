package PacketPeer

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
PacketPeer is an abstraction and base class for packet-based protocols (such as UDP). It provides an API for sending and receiving packets both as raw data or variables. This makes it easy to transfer data over a protocol, without having to encode data as low-level bytes or having to worry about network ordering.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.

*/
type Go [1]classdb.PacketPeer

/*
Gets a Variant. If [param allow_objects] is [code]true[/code], decoding objects is allowed.
Internally, this uses the same decoding mechanism as the [method @GlobalScope.bytes_to_var] method.
[b]Warning:[/b] Deserialized objects can contain code which gets executed. Do not use this option if the serialized object comes from untrusted sources to avoid potential security threats such as remote code execution.
*/
func (self Go) GetVar() gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).GetVar(gc, false))
}

/*
Sends a [Variant] as a packet. If [param full_objects] is [code]true[/code], encoding objects is allowed (and can potentially include code).
Internally, this uses the same encoding mechanism as the [method @GlobalScope.var_to_bytes] method.
*/
func (self Go) PutVar(v gd.Variant) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).PutVar(v, false))
}

/*
Gets a raw packet.
*/
func (self Go) GetPacket() []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(class(self).GetPacket(gc).Bytes())
}

/*
Sends a raw packet.
*/
func (self Go) PutPacket(buffer []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).PutPacket(gc.PackedByteSlice(buffer)))
}

/*
Returns the error state of the last packet received (via [method get_packet] and [method get_var]).
*/
func (self Go) GetPacketError() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).GetPacketError())
}

/*
Returns the number of packets currently available in the ring-buffer.
*/
func (self Go) GetAvailablePacketCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetAvailablePacketCount()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PacketPeer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PacketPeer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) EncodeBufferMaxSize() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetEncodeBufferMaxSize()))
}

func (self Go) SetEncodeBufferMaxSize(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEncodeBufferMaxSize(gd.Int(value))
}

/*
Gets a Variant. If [param allow_objects] is [code]true[/code], decoding objects is allowed.
Internally, this uses the same decoding mechanism as the [method @GlobalScope.bytes_to_var] method.
[b]Warning:[/b] Deserialized objects can contain code which gets executed. Do not use this option if the serialized object comes from untrusted sources to avoid potential security threats such as remote code execution.
*/
//go:nosplit
func (self class) GetVar(ctx gd.Lifetime, allow_objects bool) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, allow_objects)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeer.Bind_get_var, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sends a [Variant] as a packet. If [param full_objects] is [code]true[/code], encoding objects is allowed (and can potentially include code).
Internally, this uses the same encoding mechanism as the [method @GlobalScope.var_to_bytes] method.
*/
//go:nosplit
func (self class) PutVar(v gd.Variant, full_objects bool) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(v))
	callframe.Arg(frame, full_objects)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeer.Bind_put_var, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets a raw packet.
*/
//go:nosplit
func (self class) GetPacket(ctx gd.Lifetime) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeer.Bind_get_packet, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sends a raw packet.
*/
//go:nosplit
func (self class) PutPacket(buffer gd.PackedByteArray) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(buffer))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeer.Bind_put_packet, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the error state of the last packet received (via [method get_packet] and [method get_var]).
*/
//go:nosplit
func (self class) GetPacketError() int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeer.Bind_get_packet_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of packets currently available in the ring-buffer.
*/
//go:nosplit
func (self class) GetAvailablePacketCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeer.Bind_get_available_packet_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetEncodeBufferMaxSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeer.Bind_get_encode_buffer_max_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEncodeBufferMaxSize(max_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeer.Bind_set_encode_buffer_max_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsPacketPeer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPacketPeer() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("PacketPeer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
