// Package RDShaderSPIRV provides methods for working with RDShaderSPIRV object instances.
package RDShaderSPIRV

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
[RDShaderSPIRV] represents a [RDShaderFile]'s [url=https://www.khronos.org/spir/]SPIR-V[/url] code for various shader stages, as well as possible compilation error messages. SPIR-V is a low-level intermediate shader representation. This intermediate representation is not used directly by GPUs for rendering, but it can be compiled into binary shaders that GPUs can understand. Unlike compiled shaders, SPIR-V is portable across GPU models and driver versions.
This object is used by [RenderingDevice].
*/
type Instance [1]gdclass.RDShaderSPIRV

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRDShaderSPIRV() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RDShaderSPIRV

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDShaderSPIRV"))
	casted := Instance{*(*gdclass.RDShaderSPIRV)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) BytecodeVertex() []byte {
	return []byte(class(self).GetStageBytecode(0).Bytes())
}

func (self Instance) SetBytecodeVertex(value []byte) {
	class(self).SetStageBytecode(0, gd.NewPackedByteSlice(value))
}

func (self Instance) BytecodeFragment() []byte {
	return []byte(class(self).GetStageBytecode(1).Bytes())
}

func (self Instance) SetBytecodeFragment(value []byte) {
	class(self).SetStageBytecode(1, gd.NewPackedByteSlice(value))
}

func (self Instance) BytecodeTesselationControl() []byte {
	return []byte(class(self).GetStageBytecode(2).Bytes())
}

func (self Instance) SetBytecodeTesselationControl(value []byte) {
	class(self).SetStageBytecode(2, gd.NewPackedByteSlice(value))
}

func (self Instance) BytecodeTesselationEvaluation() []byte {
	return []byte(class(self).GetStageBytecode(3).Bytes())
}

func (self Instance) SetBytecodeTesselationEvaluation(value []byte) {
	class(self).SetStageBytecode(3, gd.NewPackedByteSlice(value))
}

func (self Instance) BytecodeCompute() []byte {
	return []byte(class(self).GetStageBytecode(4).Bytes())
}

func (self Instance) SetBytecodeCompute(value []byte) {
	class(self).SetStageBytecode(4, gd.NewPackedByteSlice(value))
}

func (self Instance) CompileErrorVertex() string {
	return string(class(self).GetStageCompileError(0).String())
}

func (self Instance) SetCompileErrorVertex(value string) {
	class(self).SetStageCompileError(0, gd.NewString(value))
}

func (self Instance) CompileErrorFragment() string {
	return string(class(self).GetStageCompileError(1).String())
}

func (self Instance) SetCompileErrorFragment(value string) {
	class(self).SetStageCompileError(1, gd.NewString(value))
}

func (self Instance) CompileErrorTesselationControl() string {
	return string(class(self).GetStageCompileError(2).String())
}

func (self Instance) SetCompileErrorTesselationControl(value string) {
	class(self).SetStageCompileError(2, gd.NewString(value))
}

func (self Instance) CompileErrorTesselationEvaluation() string {
	return string(class(self).GetStageCompileError(3).String())
}

func (self Instance) SetCompileErrorTesselationEvaluation(value string) {
	class(self).SetStageCompileError(3, gd.NewString(value))
}

func (self Instance) CompileErrorCompute() string {
	return string(class(self).GetStageCompileError(4).String())
}

func (self Instance) SetCompileErrorCompute(value string) {
	class(self).SetStageCompileError(4, gd.NewString(value))
}

/*
Sets the SPIR-V [param bytecode] for the given shader [param stage]. Equivalent to setting one of [member bytecode_compute], [member bytecode_fragment], [member bytecode_tesselation_control], [member bytecode_tesselation_evaluation], [member bytecode_vertex].
*/
//go:nosplit
func (self class) SetStageBytecode(stage gdclass.RenderingDeviceShaderStage, bytecode gd.PackedByteArray) {
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	callframe.Arg(frame, pointers.Get(bytecode))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSPIRV.Bind_set_stage_bytecode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Equivalent to getting one of [member bytecode_compute], [member bytecode_fragment], [member bytecode_tesselation_control], [member bytecode_tesselation_evaluation], [member bytecode_vertex].
*/
//go:nosplit
func (self class) GetStageBytecode(stage gdclass.RenderingDeviceShaderStage) gd.PackedByteArray {
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSPIRV.Bind_get_stage_bytecode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the compilation error message for the given shader [param stage] to [param compile_error]. Equivalent to setting one of [member compile_error_compute], [member compile_error_fragment], [member compile_error_tesselation_control], [member compile_error_tesselation_evaluation], [member compile_error_vertex].
*/
//go:nosplit
func (self class) SetStageCompileError(stage gdclass.RenderingDeviceShaderStage, compile_error gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	callframe.Arg(frame, pointers.Get(compile_error))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSPIRV.Bind_set_stage_compile_error, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the compilation error message for the given shader [param stage]. Equivalent to getting one of [member compile_error_compute], [member compile_error_fragment], [member compile_error_tesselation_control], [member compile_error_tesselation_evaluation], [member compile_error_vertex].
*/
//go:nosplit
func (self class) GetStageCompileError(stage gdclass.RenderingDeviceShaderStage) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSPIRV.Bind_get_stage_compile_error, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsRDShaderSPIRV() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRDShaderSPIRV() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("RDShaderSPIRV", func(ptr gd.Object) any {
		return [1]gdclass.RDShaderSPIRV{*(*gdclass.RDShaderSPIRV)(unsafe.Pointer(&ptr))}
	})
}
