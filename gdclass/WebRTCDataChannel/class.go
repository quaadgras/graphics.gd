package WebRTCDataChannel

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/PacketPeer"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

type Go [1]classdb.WebRTCDataChannel

/*
Reserved, but not used for now.
*/
func (self Go) Poll() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Poll())
}

/*
Closes this data channel, notifying the other peer.
*/
func (self Go) Close() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Close()
}

/*
Returns [code]true[/code] if the last received packet was transferred as text. See [member write_mode].
*/
func (self Go) WasStringPacket() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).WasStringPacket())
}

/*
Returns the current state of this channel, see [enum ChannelState].
*/
func (self Go) GetReadyState() classdb.WebRTCDataChannelChannelState {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.WebRTCDataChannelChannelState(class(self).GetReadyState())
}

/*
Returns the label assigned to this channel during creation.
*/
func (self Go) GetLabel() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetLabel(gc).String())
}

/*
Returns [code]true[/code] if this channel was created with ordering enabled (default).
*/
func (self Go) IsOrdered() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsOrdered())
}

/*
Returns the ID assigned to this channel during creation (or auto-assigned during negotiation).
If the channel is not negotiated out-of-band the ID will only be available after the connection is established (will return [code]65535[/code] until then).
*/
func (self Go) GetId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetId()))
}

/*
Returns the [code]maxPacketLifeTime[/code] value assigned to this channel during creation.
Will be [code]65535[/code] if not specified.
*/
func (self Go) GetMaxPacketLifeTime() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetMaxPacketLifeTime()))
}

/*
Returns the [code]maxRetransmits[/code] value assigned to this channel during creation.
Will be [code]65535[/code] if not specified.
*/
func (self Go) GetMaxRetransmits() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetMaxRetransmits()))
}

/*
Returns the sub-protocol assigned to this channel during creation. An empty string if not specified.
*/
func (self Go) GetProtocol() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetProtocol(gc).String())
}

/*
Returns [code]true[/code] if this channel was created with out-of-band configuration.
*/
func (self Go) IsNegotiated() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsNegotiated())
}

/*
Returns the number of bytes currently queued to be sent over this channel.
*/
func (self Go) GetBufferedAmount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetBufferedAmount()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.WebRTCDataChannel
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("WebRTCDataChannel"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) WriteMode() classdb.WebRTCDataChannelWriteMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.WebRTCDataChannelWriteMode(class(self).GetWriteMode())
}

func (self Go) SetWriteMode(value classdb.WebRTCDataChannelWriteMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetWriteMode(value)
}

/*
Reserved, but not used for now.
*/
//go:nosplit
func (self class) Poll() int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCDataChannel.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Closes this data channel, notifying the other peer.
*/
//go:nosplit
func (self class) Close()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCDataChannel.Bind_close, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the last received packet was transferred as text. See [member write_mode].
*/
//go:nosplit
func (self class) WasStringPacket() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCDataChannel.Bind_was_string_packet, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWriteMode(write_mode classdb.WebRTCDataChannelWriteMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, write_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCDataChannel.Bind_set_write_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWriteMode() classdb.WebRTCDataChannelWriteMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.WebRTCDataChannelWriteMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCDataChannel.Bind_get_write_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the current state of this channel, see [enum ChannelState].
*/
//go:nosplit
func (self class) GetReadyState() classdb.WebRTCDataChannelChannelState {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.WebRTCDataChannelChannelState](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCDataChannel.Bind_get_ready_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the label assigned to this channel during creation.
*/
//go:nosplit
func (self class) GetLabel(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCDataChannel.Bind_get_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this channel was created with ordering enabled (default).
*/
//go:nosplit
func (self class) IsOrdered() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCDataChannel.Bind_is_ordered, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the ID assigned to this channel during creation (or auto-assigned during negotiation).
If the channel is not negotiated out-of-band the ID will only be available after the connection is established (will return [code]65535[/code] until then).
*/
//go:nosplit
func (self class) GetId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCDataChannel.Bind_get_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [code]maxPacketLifeTime[/code] value assigned to this channel during creation.
Will be [code]65535[/code] if not specified.
*/
//go:nosplit
func (self class) GetMaxPacketLifeTime() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCDataChannel.Bind_get_max_packet_life_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [code]maxRetransmits[/code] value assigned to this channel during creation.
Will be [code]65535[/code] if not specified.
*/
//go:nosplit
func (self class) GetMaxRetransmits() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCDataChannel.Bind_get_max_retransmits, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the sub-protocol assigned to this channel during creation. An empty string if not specified.
*/
//go:nosplit
func (self class) GetProtocol(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCDataChannel.Bind_get_protocol, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this channel was created with out-of-band configuration.
*/
//go:nosplit
func (self class) IsNegotiated() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCDataChannel.Bind_is_negotiated, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of bytes currently queued to be sent over this channel.
*/
//go:nosplit
func (self class) GetBufferedAmount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.WebRTCDataChannel.Bind_get_buffered_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsWebRTCDataChannel() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsWebRTCDataChannel() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsPacketPeer() PacketPeer.GD { return *((*PacketPeer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsPacketPeer() PacketPeer.Go { return *((*PacketPeer.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsPacketPeer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsPacketPeer(), name)
	}
}
func init() {classdb.Register("WebRTCDataChannel", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type WriteMode = classdb.WebRTCDataChannelWriteMode

const (
/*Tells the channel to send data over this channel as text. An external peer (non-Godot) would receive this as a string.*/
	WriteModeText WriteMode = 0
/*Tells the channel to send data over this channel as binary. An external peer (non-Godot) would receive this as array buffer or blob.*/
	WriteModeBinary WriteMode = 1
)
type ChannelState = classdb.WebRTCDataChannelChannelState

const (
/*The channel was created, but it's still trying to connect.*/
	StateConnecting ChannelState = 0
/*The channel is currently open, and data can flow over it.*/
	StateOpen ChannelState = 1
/*The channel is being closed, no new messages will be accepted, but those already in queue will be flushed.*/
	StateClosing ChannelState = 2
/*The channel was closed, or connection failed.*/
	StateClosed ChannelState = 3
)
