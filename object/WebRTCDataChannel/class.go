package WebRTCDataChannel

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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

type Simple [1]classdb.WebRTCDataChannel
func (self Simple) Poll() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Poll())
}
func (self Simple) Close() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Close()
}
func (self Simple) WasStringPacket() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).WasStringPacket())
}
func (self Simple) SetWriteMode(write_mode classdb.WebRTCDataChannelWriteMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetWriteMode(write_mode)
}
func (self Simple) GetWriteMode() classdb.WebRTCDataChannelWriteMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.WebRTCDataChannelWriteMode(Expert(self).GetWriteMode())
}
func (self Simple) GetReadyState() classdb.WebRTCDataChannelChannelState {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.WebRTCDataChannelChannelState(Expert(self).GetReadyState())
}
func (self Simple) GetLabel() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetLabel(gc).String())
}
func (self Simple) IsOrdered() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsOrdered())
}
func (self Simple) GetId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetId()))
}
func (self Simple) GetMaxPacketLifeTime() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMaxPacketLifeTime()))
}
func (self Simple) GetMaxRetransmits() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMaxRetransmits()))
}
func (self Simple) GetProtocol() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetProtocol(gc).String())
}
func (self Simple) IsNegotiated() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsNegotiated())
}
func (self Simple) GetBufferedAmount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBufferedAmount()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.WebRTCDataChannel
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsWebRTCDataChannel() Expert { return self[0].AsWebRTCDataChannel() }


//go:nosplit
func (self Simple) AsWebRTCDataChannel() Simple { return self[0].AsWebRTCDataChannel() }


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
func init() {classdb.Register("WebRTCDataChannel", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
