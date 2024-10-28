package RDShaderSource

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Shader source code in text form.
See also [RDShaderFile]. [RDShaderSource] is only meant to be used with the [RenderingDevice] API. It should not be confused with Godot's own [Shader] resource, which is what Godot's various nodes use for high-level shader programming.

*/
type Go [1]classdb.RDShaderSource
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RDShaderSource
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDShaderSource"))
	return Go{classdb.RDShaderSource(object)}
}

func (self Go) SourceVertex() string {
		return string(class(self).GetStageSource(0).String())
}

func (self Go) SetSourceVertex(value string) {
	class(self).SetStageSource(0, gd.NewString(value))
}

func (self Go) SourceFragment() string {
		return string(class(self).GetStageSource(1).String())
}

func (self Go) SetSourceFragment(value string) {
	class(self).SetStageSource(1, gd.NewString(value))
}

func (self Go) SourceTesselationControl() string {
		return string(class(self).GetStageSource(2).String())
}

func (self Go) SetSourceTesselationControl(value string) {
	class(self).SetStageSource(2, gd.NewString(value))
}

func (self Go) SourceTesselationEvaluation() string {
		return string(class(self).GetStageSource(3).String())
}

func (self Go) SetSourceTesselationEvaluation(value string) {
	class(self).SetStageSource(3, gd.NewString(value))
}

func (self Go) SourceCompute() string {
		return string(class(self).GetStageSource(4).String())
}

func (self Go) SetSourceCompute(value string) {
	class(self).SetStageSource(4, gd.NewString(value))
}

func (self Go) Language() classdb.RenderingDeviceShaderLanguage {
		return classdb.RenderingDeviceShaderLanguage(class(self).GetLanguage())
}

func (self Go) SetLanguage(value classdb.RenderingDeviceShaderLanguage) {
	class(self).SetLanguage(value)
}

/*
Sets [param source] code for the specified shader [param stage]. Equivalent to setting one of [member source_compute], [member source_fragment], [member source_tesselation_control], [member source_tesselation_evaluation] or [member source_vertex].
*/
//go:nosplit
func (self class) SetStageSource(stage classdb.RenderingDeviceShaderStage, source gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	callframe.Arg(frame, discreet.Get(source))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSource.Bind_set_stage_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns source code for the specified shader [param stage]. Equivalent to getting one of [member source_compute], [member source_fragment], [member source_tesselation_control], [member source_tesselation_evaluation] or [member source_vertex].
*/
//go:nosplit
func (self class) GetStageSource(stage classdb.RenderingDeviceShaderStage) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSource.Bind_get_stage_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLanguage(language classdb.RenderingDeviceShaderLanguage)  {
	var frame = callframe.New()
	callframe.Arg(frame, language)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSource.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLanguage() classdb.RenderingDeviceShaderLanguage {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceShaderLanguage](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSource.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDShaderSource() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRDShaderSource() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("RDShaderSource", func(ptr gd.Object) any { return classdb.RDShaderSource(ptr) })}
