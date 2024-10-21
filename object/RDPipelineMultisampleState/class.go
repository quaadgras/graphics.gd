package RDPipelineMultisampleState

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
[RDPipelineMultisampleState] is used to control how multisample or supersample antialiasing is being performed when rendering using [RenderingDevice].

*/
type Simple [1]classdb.RDPipelineMultisampleState
func (self Simple) SetSampleCount(p_member classdb.RenderingDeviceTextureSamples) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSampleCount(p_member)
}
func (self Simple) GetSampleCount() classdb.RenderingDeviceTextureSamples {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceTextureSamples(Expert(self).GetSampleCount())
}
func (self Simple) SetEnableSampleShading(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnableSampleShading(p_member)
}
func (self Simple) GetEnableSampleShading() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetEnableSampleShading())
}
func (self Simple) SetMinSampleShading(p_member float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMinSampleShading(gd.Float(p_member))
}
func (self Simple) GetMinSampleShading() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMinSampleShading()))
}
func (self Simple) SetEnableAlphaToCoverage(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnableAlphaToCoverage(p_member)
}
func (self Simple) GetEnableAlphaToCoverage() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetEnableAlphaToCoverage())
}
func (self Simple) SetEnableAlphaToOne(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnableAlphaToOne(p_member)
}
func (self Simple) GetEnableAlphaToOne() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetEnableAlphaToOne())
}
func (self Simple) SetSampleMasks(masks gd.ArrayOf[gd.Int]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSampleMasks(masks)
}
func (self Simple) GetSampleMasks() gd.ArrayOf[gd.Int] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Int](Expert(self).GetSampleMasks(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RDPipelineMultisampleState
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsRDPipelineMultisampleState() Expert { return self[0].AsRDPipelineMultisampleState() }


//go:nosplit
func (self Simple) AsRDPipelineMultisampleState() Simple { return self[0].AsRDPipelineMultisampleState() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("RDPipelineMultisampleState", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
