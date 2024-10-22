package RDTextureFormat

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
This object is used by [RenderingDevice].

*/
type Go [1]classdb.RDTextureFormat
func (self Go) AddShareableFormat(format classdb.RenderingDeviceDataFormat) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddShareableFormat(format)
}
func (self Go) RemoveShareableFormat(format classdb.RenderingDeviceDataFormat) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveShareableFormat(format)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RDTextureFormat
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RDTextureFormat"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Format() classdb.RenderingDeviceDataFormat {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceDataFormat(class(self).GetFormat())
}

func (self Go) SetFormat(value classdb.RenderingDeviceDataFormat) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFormat(value)
}

func (self Go) Width() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetWidth()))
}

func (self Go) SetWidth(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetWidth(gd.Int(value))
}

func (self Go) Height() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetHeight()))
}

func (self Go) SetHeight(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHeight(gd.Int(value))
}

func (self Go) Depth() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetDepth()))
}

func (self Go) SetDepth(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDepth(gd.Int(value))
}

func (self Go) ArrayLayers() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetArrayLayers()))
}

func (self Go) SetArrayLayers(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetArrayLayers(gd.Int(value))
}

func (self Go) Mipmaps() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMipmaps()))
}

func (self Go) SetMipmaps(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMipmaps(gd.Int(value))
}

func (self Go) TextureType() classdb.RenderingDeviceTextureType {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceTextureType(class(self).GetTextureType())
}

func (self Go) SetTextureType(value classdb.RenderingDeviceTextureType) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextureType(value)
}

func (self Go) Samples() classdb.RenderingDeviceTextureSamples {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceTextureSamples(class(self).GetSamples())
}

func (self Go) SetSamples(value classdb.RenderingDeviceTextureSamples) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSamples(value)
}

func (self Go) UsageBits() classdb.RenderingDeviceTextureUsageBits {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceTextureUsageBits(class(self).GetUsageBits())
}

func (self Go) SetUsageBits(value classdb.RenderingDeviceTextureUsageBits) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUsageBits(value)
}

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
func (self class) AsRDTextureFormat() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRDTextureFormat() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("RDTextureFormat", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
