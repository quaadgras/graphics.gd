package OggPacketSequence

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A sequence of Ogg packets.

*/
type Simple [1]classdb.OggPacketSequence
func (self Simple) SetPacketData(packet_data gd.ArrayOf[gd.Array]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPacketData(packet_data)
}
func (self Simple) GetPacketData() gd.ArrayOf[gd.Array] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Array](Expert(self).GetPacketData(gc))
}
func (self Simple) SetPacketGranulePositions(granule_positions gd.PackedInt64Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPacketGranulePositions(granule_positions)
}
func (self Simple) GetPacketGranulePositions() gd.PackedInt64Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt64Array(Expert(self).GetPacketGranulePositions(gc))
}
func (self Simple) SetSamplingRate(sampling_rate float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSamplingRate(gd.Float(sampling_rate))
}
func (self Simple) GetSamplingRate() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSamplingRate()))
}
func (self Simple) GetLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetLength()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.OggPacketSequence
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetPacketData(packet_data gd.ArrayOf[gd.Array])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(packet_data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OggPacketSequence.Bind_set_packet_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPacketData(ctx gd.Lifetime) gd.ArrayOf[gd.Array] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OggPacketSequence.Bind_get_packet_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Array](ret)
}
//go:nosplit
func (self class) SetPacketGranulePositions(granule_positions gd.PackedInt64Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(granule_positions))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OggPacketSequence.Bind_set_packet_granule_positions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPacketGranulePositions(ctx gd.Lifetime) gd.PackedInt64Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OggPacketSequence.Bind_get_packet_granule_positions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt64Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSamplingRate(sampling_rate gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sampling_rate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OggPacketSequence.Bind_set_sampling_rate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSamplingRate() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OggPacketSequence.Bind_get_sampling_rate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
The length of this stream, in seconds.
*/
//go:nosplit
func (self class) GetLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OggPacketSequence.Bind_get_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsOggPacketSequence() Expert { return self[0].AsOggPacketSequence() }


//go:nosplit
func (self Simple) AsOggPacketSequence() Simple { return self[0].AsOggPacketSequence() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("OggPacketSequence", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
