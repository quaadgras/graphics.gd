package RDTextureFormat

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
type Simple [1]classdb.RDTextureFormat
func (self Simple) SetFormat(p_member classdb.RenderingDeviceDataFormat) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFormat(p_member)
}
func (self Simple) GetFormat() classdb.RenderingDeviceDataFormat {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceDataFormat(Expert(self).GetFormat())
}
func (self Simple) SetWidth(p_member int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetWidth(gd.Int(p_member))
}
func (self Simple) GetWidth() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetWidth()))
}
func (self Simple) SetHeight(p_member int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHeight(gd.Int(p_member))
}
func (self Simple) GetHeight() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetHeight()))
}
func (self Simple) SetDepth(p_member int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDepth(gd.Int(p_member))
}
func (self Simple) GetDepth() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetDepth()))
}
func (self Simple) SetArrayLayers(p_member int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetArrayLayers(gd.Int(p_member))
}
func (self Simple) GetArrayLayers() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetArrayLayers()))
}
func (self Simple) SetMipmaps(p_member int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMipmaps(gd.Int(p_member))
}
func (self Simple) GetMipmaps() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMipmaps()))
}
func (self Simple) SetTextureType(p_member classdb.RenderingDeviceTextureType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTextureType(p_member)
}
func (self Simple) GetTextureType() classdb.RenderingDeviceTextureType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceTextureType(Expert(self).GetTextureType())
}
func (self Simple) SetSamples(p_member classdb.RenderingDeviceTextureSamples) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSamples(p_member)
}
func (self Simple) GetSamples() classdb.RenderingDeviceTextureSamples {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceTextureSamples(Expert(self).GetSamples())
}
func (self Simple) SetUsageBits(p_member classdb.RenderingDeviceTextureUsageBits) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUsageBits(p_member)
}
func (self Simple) GetUsageBits() classdb.RenderingDeviceTextureUsageBits {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceTextureUsageBits(Expert(self).GetUsageBits())
}
func (self Simple) AddShareableFormat(format classdb.RenderingDeviceDataFormat) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddShareableFormat(format)
}
func (self Simple) RemoveShareableFormat(format classdb.RenderingDeviceDataFormat) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveShareableFormat(format)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RDTextureFormat
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetFormat(p_member classdb.RenderingDeviceDataFormat)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_set_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFormat() classdb.RenderingDeviceDataFormat {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceDataFormat](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWidth(p_member gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWidth() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHeight(p_member gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHeight() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDepth(p_member gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_set_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDepth() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_get_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetArrayLayers(p_member gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_set_array_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetArrayLayers() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_get_array_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMipmaps(p_member gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_set_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMipmaps() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_get_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextureType(p_member classdb.RenderingDeviceTextureType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_set_texture_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextureType() classdb.RenderingDeviceTextureType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceTextureType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_get_texture_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSamples(p_member classdb.RenderingDeviceTextureSamples)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_set_samples, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSamples() classdb.RenderingDeviceTextureSamples {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceTextureSamples](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_get_samples, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUsageBits(p_member classdb.RenderingDeviceTextureUsageBits)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_set_usage_bits, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUsageBits() classdb.RenderingDeviceTextureUsageBits {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceTextureUsageBits](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_get_usage_bits, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) AddShareableFormat(format classdb.RenderingDeviceDataFormat)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, format)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_add_shareable_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) RemoveShareableFormat(format classdb.RenderingDeviceDataFormat)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, format)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDTextureFormat.Bind_remove_shareable_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsRDTextureFormat() Expert { return self[0].AsRDTextureFormat() }


//go:nosplit
func (self Simple) AsRDTextureFormat() Simple { return self[0].AsRDTextureFormat() }


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
func init() {classdb.Register("RDTextureFormat", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
