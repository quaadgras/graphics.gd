package RDShaderSource

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Shader source code in text form.
See also [RDShaderFile]. [RDShaderSource] is only meant to be used with the [RenderingDevice] API. It should not be confused with Godot's own [Shader] resource, which is what Godot's various nodes use for high-level shader programming.
*/
type Instance [1]gdclass.RDShaderSource
type Any interface {
	gd.IsClass
	AsRDShaderSource() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RDShaderSource

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDShaderSource"))
	return Instance{*(*gdclass.RDShaderSource)(unsafe.Pointer(&object))}
}

func (self Instance) SourceVertex() string {
	return string(class(self).GetStageSource(0).String())
}

func (self Instance) SetSourceVertex(value string) {
	class(self).SetStageSource(0, gd.NewString(value))
}

func (self Instance) SourceFragment() string {
	return string(class(self).GetStageSource(1).String())
}

func (self Instance) SetSourceFragment(value string) {
	class(self).SetStageSource(1, gd.NewString(value))
}

func (self Instance) SourceTesselationControl() string {
	return string(class(self).GetStageSource(2).String())
}

func (self Instance) SetSourceTesselationControl(value string) {
	class(self).SetStageSource(2, gd.NewString(value))
}

func (self Instance) SourceTesselationEvaluation() string {
	return string(class(self).GetStageSource(3).String())
}

func (self Instance) SetSourceTesselationEvaluation(value string) {
	class(self).SetStageSource(3, gd.NewString(value))
}

func (self Instance) SourceCompute() string {
	return string(class(self).GetStageSource(4).String())
}

func (self Instance) SetSourceCompute(value string) {
	class(self).SetStageSource(4, gd.NewString(value))
}

func (self Instance) Language() gdclass.RenderingDeviceShaderLanguage {
	return gdclass.RenderingDeviceShaderLanguage(class(self).GetLanguage())
}

func (self Instance) SetLanguage(value gdclass.RenderingDeviceShaderLanguage) {
	class(self).SetLanguage(value)
}

/*
Sets [param source] code for the specified shader [param stage]. Equivalent to setting one of [member source_compute], [member source_fragment], [member source_tesselation_control], [member source_tesselation_evaluation] or [member source_vertex].
*/
//go:nosplit
func (self class) SetStageSource(stage gdclass.RenderingDeviceShaderStage, source gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	callframe.Arg(frame, pointers.Get(source))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSource.Bind_set_stage_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns source code for the specified shader [param stage]. Equivalent to getting one of [member source_compute], [member source_fragment], [member source_tesselation_control], [member source_tesselation_evaluation] or [member source_vertex].
*/
//go:nosplit
func (self class) GetStageSource(stage gdclass.RenderingDeviceShaderStage) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSource.Bind_get_stage_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLanguage(language gdclass.RenderingDeviceShaderLanguage) {
	var frame = callframe.New()
	callframe.Arg(frame, language)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSource.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLanguage() gdclass.RenderingDeviceShaderLanguage {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceShaderLanguage](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSource.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDShaderSource() Advanced     { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRDShaderSource() Instance  { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	gdclass.Register("RDShaderSource", func(ptr gd.Object) any {
		return [1]gdclass.RDShaderSource{*(*gdclass.RDShaderSource)(unsafe.Pointer(&ptr))}
	})
}
