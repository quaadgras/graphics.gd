package OggPacketSequence

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A sequence of Ogg packets.

*/
type Go [1]classdb.OggPacketSequence

/*
The length of this stream, in seconds.
*/
func (self Go) GetLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetLength()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.OggPacketSequence
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("OggPacketSequence"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) PacketData() gd.ArrayOf[gd.Array] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gd.Array](class(self).GetPacketData(gc))
}

func (self Go) SetPacketData(value gd.ArrayOf[gd.Array]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPacketData(value)
}

func (self Go) GranulePositions() []int64 {
	gc := gd.GarbageCollector(); _ = gc
		return []int64(class(self).GetPacketGranulePositions(gc).AsSlice())
}

func (self Go) SetGranulePositions(value []int64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPacketGranulePositions(gc.PackedInt64Slice(value))
}

func (self Go) SamplingRate() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSamplingRate()))
}

func (self Go) SetSamplingRate(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSamplingRate(gd.Float(value))
}

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
func (self class) AsOggPacketSequence() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsOggPacketSequence() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("OggPacketSequence", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
