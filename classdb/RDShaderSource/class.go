// Package RDShaderSource provides methods for working with RDShaderSource object instances.
package RDShaderSource

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
Shader source code in text form.
See also [RDShaderFile]. [RDShaderSource] is only meant to be used with the [RenderingDevice] API. It should not be confused with Godot's own [Shader] resource, which is what Godot's various nodes use for high-level shader programming.
*/
type Instance [1]gdclass.RDShaderSource

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRDShaderSource() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RDShaderSource

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDShaderSource"))
	casted := Instance{*(*gdclass.RDShaderSource)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) SourceVertex() string {
	return string(class(self).GetStageSource(0).String())
}

func (self Instance) SetSourceVertex(value string) {
	class(self).SetStageSource(0, String.New(value))
}

func (self Instance) SourceFragment() string {
	return string(class(self).GetStageSource(1).String())
}

func (self Instance) SetSourceFragment(value string) {
	class(self).SetStageSource(1, String.New(value))
}

func (self Instance) SourceTesselationControl() string {
	return string(class(self).GetStageSource(2).String())
}

func (self Instance) SetSourceTesselationControl(value string) {
	class(self).SetStageSource(2, String.New(value))
}

func (self Instance) SourceTesselationEvaluation() string {
	return string(class(self).GetStageSource(3).String())
}

func (self Instance) SetSourceTesselationEvaluation(value string) {
	class(self).SetStageSource(3, String.New(value))
}

func (self Instance) SourceCompute() string {
	return string(class(self).GetStageSource(4).String())
}

func (self Instance) SetSourceCompute(value string) {
	class(self).SetStageSource(4, String.New(value))
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
func (self class) SetStageSource(stage gdclass.RenderingDeviceShaderStage, source String.Readable) { //gd:RDShaderSource.set_stage_source
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	callframe.Arg(frame, pointers.Get(gd.InternalString(source)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSource.Bind_set_stage_source, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns source code for the specified shader [param stage]. Equivalent to getting one of [member source_compute], [member source_fragment], [member source_tesselation_control], [member source_tesselation_evaluation] or [member source_vertex].
*/
//go:nosplit
func (self class) GetStageSource(stage gdclass.RenderingDeviceShaderStage) String.Readable { //gd:RDShaderSource.get_stage_source
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSource.Bind_get_stage_source, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLanguage(language gdclass.RenderingDeviceShaderLanguage) { //gd:RDShaderSource.set_language
	var frame = callframe.New()
	callframe.Arg(frame, language)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSource.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLanguage() gdclass.RenderingDeviceShaderLanguage { //gd:RDShaderSource.get_language
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceShaderLanguage](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSource.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDShaderSource() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRDShaderSource() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("RDShaderSource", func(ptr gd.Object) any {
		return [1]gdclass.RDShaderSource{*(*gdclass.RDShaderSource)(unsafe.Pointer(&ptr))}
	})
}
