package RDShaderSource

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
Shader source code in text form.
See also [RDShaderFile]. [RDShaderSource] is only meant to be used with the [RenderingDevice] API. It should not be confused with Godot's own [Shader] resource, which is what Godot's various nodes use for high-level shader programming.

*/
type Simple [1]classdb.RDShaderSource
func (self Simple) SetStageSource(stage classdb.RenderingDeviceShaderStage, source string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStageSource(stage, gc.String(source))
}
func (self Simple) GetStageSource(stage classdb.RenderingDeviceShaderStage) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetStageSource(gc, stage).String())
}
func (self Simple) SetLanguage(language classdb.RenderingDeviceShaderLanguage) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLanguage(language)
}
func (self Simple) GetLanguage() classdb.RenderingDeviceShaderLanguage {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceShaderLanguage(Expert(self).GetLanguage())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RDShaderSource
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Sets [param source] code for the specified shader [param stage]. Equivalent to setting one of [member source_compute], [member source_fragment], [member source_tesselation_control], [member source_tesselation_evaluation] or [member source_vertex].
*/
//go:nosplit
func (self class) SetStageSource(stage classdb.RenderingDeviceShaderStage, source gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	callframe.Arg(frame, mmm.Get(source))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDShaderSource.Bind_set_stage_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns source code for the specified shader [param stage]. Equivalent to getting one of [member source_compute], [member source_fragment], [member source_tesselation_control], [member source_tesselation_evaluation] or [member source_vertex].
*/
//go:nosplit
func (self class) GetStageSource(ctx gd.Lifetime, stage classdb.RenderingDeviceShaderStage) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDShaderSource.Bind_get_stage_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLanguage(language classdb.RenderingDeviceShaderLanguage)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, language)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDShaderSource.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLanguage() classdb.RenderingDeviceShaderLanguage {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceShaderLanguage](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDShaderSource.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsRDShaderSource() Expert { return self[0].AsRDShaderSource() }


//go:nosplit
func (self Simple) AsRDShaderSource() Simple { return self[0].AsRDShaderSource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("RDShaderSource", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
