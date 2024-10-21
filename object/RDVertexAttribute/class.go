package RDVertexAttribute

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This object is used by [RenderingDevice].

*/
type Simple [1]classdb.RDVertexAttribute
func (self Simple) SetLocation(p_member int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLocation(gd.Int(p_member))
}
func (self Simple) GetLocation() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetLocation()))
}
func (self Simple) SetOffset(p_member int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOffset(gd.Int(p_member))
}
func (self Simple) GetOffset() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetOffset()))
}
func (self Simple) SetFormat(p_member classdb.RenderingDeviceDataFormat) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFormat(p_member)
}
func (self Simple) GetFormat() classdb.RenderingDeviceDataFormat {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceDataFormat(Expert(self).GetFormat())
}
func (self Simple) SetStride(p_member int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStride(gd.Int(p_member))
}
func (self Simple) GetStride() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetStride()))
}
func (self Simple) SetFrequency(p_member classdb.RenderingDeviceVertexFrequency) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFrequency(p_member)
}
func (self Simple) GetFrequency() classdb.RenderingDeviceVertexFrequency {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceVertexFrequency(Expert(self).GetFrequency())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RDVertexAttribute
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetLocation(p_member gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDVertexAttribute.Bind_set_location, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLocation() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDVertexAttribute.Bind_get_location, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOffset(p_member gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDVertexAttribute.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOffset() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDVertexAttribute.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFormat(p_member classdb.RenderingDeviceDataFormat)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDVertexAttribute.Bind_set_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFormat() classdb.RenderingDeviceDataFormat {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceDataFormat](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDVertexAttribute.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStride(p_member gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDVertexAttribute.Bind_set_stride, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStride() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDVertexAttribute.Bind_get_stride, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFrequency(p_member classdb.RenderingDeviceVertexFrequency)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDVertexAttribute.Bind_set_frequency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFrequency() classdb.RenderingDeviceVertexFrequency {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceVertexFrequency](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDVertexAttribute.Bind_get_frequency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsRDVertexAttribute() Expert { return self[0].AsRDVertexAttribute() }


//go:nosplit
func (self Simple) AsRDVertexAttribute() Simple { return self[0].AsRDVertexAttribute() }


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
func init() {classdb.Register("RDVertexAttribute", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
