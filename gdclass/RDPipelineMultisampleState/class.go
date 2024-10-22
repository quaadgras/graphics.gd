package RDPipelineMultisampleState

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
[RDPipelineMultisampleState] is used to control how multisample or supersample antialiasing is being performed when rendering using [RenderingDevice].

*/
type Go [1]classdb.RDPipelineMultisampleState
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RDPipelineMultisampleState
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RDPipelineMultisampleState"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) SampleCount() classdb.RenderingDeviceTextureSamples {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceTextureSamples(class(self).GetSampleCount())
}

func (self Go) SetSampleCount(value classdb.RenderingDeviceTextureSamples) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSampleCount(value)
}

func (self Go) EnableSampleShading() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetEnableSampleShading())
}

func (self Go) SetEnableSampleShading(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnableSampleShading(value)
}

func (self Go) MinSampleShading() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMinSampleShading()))
}

func (self Go) SetMinSampleShading(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMinSampleShading(gd.Float(value))
}

func (self Go) EnableAlphaToCoverage() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetEnableAlphaToCoverage())
}

func (self Go) SetEnableAlphaToCoverage(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnableAlphaToCoverage(value)
}

func (self Go) EnableAlphaToOne() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetEnableAlphaToOne())
}

func (self Go) SetEnableAlphaToOne(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnableAlphaToOne(value)
}

func (self Go) SampleMasks() gd.ArrayOf[gd.Int] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gd.Int](class(self).GetSampleMasks(gc))
}

func (self Go) SetSampleMasks(value gd.ArrayOf[gd.Int]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSampleMasks(value)
}

//go:nosplit
func (self class) SetSampleCount(p_member classdb.RenderingDeviceTextureSamples)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineMultisampleState.Bind_set_sample_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSampleCount() classdb.RenderingDeviceTextureSamples {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceTextureSamples](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineMultisampleState.Bind_get_sample_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnableSampleShading(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineMultisampleState.Bind_set_enable_sample_shading, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnableSampleShading() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineMultisampleState.Bind_get_enable_sample_shading, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMinSampleShading(p_member gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineMultisampleState.Bind_set_min_sample_shading, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMinSampleShading() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineMultisampleState.Bind_get_min_sample_shading, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnableAlphaToCoverage(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineMultisampleState.Bind_set_enable_alpha_to_coverage, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnableAlphaToCoverage() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineMultisampleState.Bind_get_enable_alpha_to_coverage, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnableAlphaToOne(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineMultisampleState.Bind_set_enable_alpha_to_one, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnableAlphaToOne() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineMultisampleState.Bind_get_enable_alpha_to_one, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSampleMasks(masks gd.ArrayOf[gd.Int])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(masks))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineMultisampleState.Bind_set_sample_masks, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSampleMasks(ctx gd.Lifetime) gd.ArrayOf[gd.Int] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineMultisampleState.Bind_get_sample_masks, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Int](ret)
}
func (self class) AsRDPipelineMultisampleState() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRDPipelineMultisampleState() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("RDPipelineMultisampleState", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
