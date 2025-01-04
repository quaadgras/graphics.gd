package OggPacketSequence

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A sequence of Ogg packets.
*/
type Instance [1]gdclass.OggPacketSequence
type Any interface {
	gd.IsClass
	AsOggPacketSequence() Instance
}

/*
The length of this stream, in seconds.
*/
func (self Instance) GetLength() Float.X {
	return Float.X(Float.X(class(self).GetLength()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.OggPacketSequence

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OggPacketSequence"))
	return Instance{*(*gdclass.OggPacketSequence)(unsafe.Pointer(&object))}
}

func (self Instance) PacketData() gd.Array {
	return gd.Array(class(self).GetPacketData())
}

func (self Instance) SetPacketData(value gd.Array) {
	class(self).SetPacketData(value)
}

func (self Instance) GranulePositions() []int64 {
	return []int64(class(self).GetPacketGranulePositions().AsSlice())
}

func (self Instance) SetGranulePositions(value []int64) {
	class(self).SetPacketGranulePositions(gd.NewPackedInt64Slice(value))
}

func (self Instance) SamplingRate() Float.X {
	return Float.X(Float.X(class(self).GetSamplingRate()))
}

func (self Instance) SetSamplingRate(value Float.X) {
	class(self).SetSamplingRate(gd.Float(value))
}

//go:nosplit
func (self class) SetPacketData(packet_data gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(packet_data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OggPacketSequence.Bind_set_packet_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPacketData() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OggPacketSequence.Bind_get_packet_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPacketGranulePositions(granule_positions gd.PackedInt64Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(granule_positions))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OggPacketSequence.Bind_set_packet_granule_positions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPacketGranulePositions() gd.PackedInt64Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OggPacketSequence.Bind_get_packet_granule_positions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt64Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSamplingRate(sampling_rate gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, sampling_rate)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OggPacketSequence.Bind_set_sampling_rate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSamplingRate() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OggPacketSequence.Bind_get_sampling_rate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
The length of this stream, in seconds.
*/
//go:nosplit
func (self class) GetLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OggPacketSequence.Bind_get_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsOggPacketSequence() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOggPacketSequence() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	gdclass.Register("OggPacketSequence", func(ptr gd.Object) any {
		return [1]gdclass.OggPacketSequence{*(*gdclass.OggPacketSequence)(unsafe.Pointer(&ptr))}
	})
}
