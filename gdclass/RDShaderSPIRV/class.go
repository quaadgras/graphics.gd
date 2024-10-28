package RDShaderSPIRV

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
[RDShaderSPIRV] represents a [RDShaderFile]'s [url=https://www.khronos.org/spir/]SPIR-V[/url] code for various shader stages, as well as possible compilation error messages. SPIR-V is a low-level intermediate shader representation. This intermediate representation is not used directly by GPUs for rendering, but it can be compiled into binary shaders that GPUs can understand. Unlike compiled shaders, SPIR-V is portable across GPU models and driver versions.
This object is used by [RenderingDevice].

*/
type Go [1]classdb.RDShaderSPIRV
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RDShaderSPIRV
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDShaderSPIRV"))
	return Go{classdb.RDShaderSPIRV(object)}
}

func (self Go) BytecodeVertex() []byte {
		return []byte(class(self).GetStageBytecode(0).Bytes())
}

func (self Go) SetBytecodeVertex(value []byte) {
	class(self).SetStageBytecode(0, gd.NewPackedByteSlice(value))
}

func (self Go) BytecodeFragment() []byte {
		return []byte(class(self).GetStageBytecode(1).Bytes())
}

func (self Go) SetBytecodeFragment(value []byte) {
	class(self).SetStageBytecode(1, gd.NewPackedByteSlice(value))
}

func (self Go) BytecodeTesselationControl() []byte {
		return []byte(class(self).GetStageBytecode(2).Bytes())
}

func (self Go) SetBytecodeTesselationControl(value []byte) {
	class(self).SetStageBytecode(2, gd.NewPackedByteSlice(value))
}

func (self Go) BytecodeTesselationEvaluation() []byte {
		return []byte(class(self).GetStageBytecode(3).Bytes())
}

func (self Go) SetBytecodeTesselationEvaluation(value []byte) {
	class(self).SetStageBytecode(3, gd.NewPackedByteSlice(value))
}

func (self Go) BytecodeCompute() []byte {
		return []byte(class(self).GetStageBytecode(4).Bytes())
}

func (self Go) SetBytecodeCompute(value []byte) {
	class(self).SetStageBytecode(4, gd.NewPackedByteSlice(value))
}

func (self Go) CompileErrorVertex() string {
		return string(class(self).GetStageCompileError(0).String())
}

func (self Go) SetCompileErrorVertex(value string) {
	class(self).SetStageCompileError(0, gd.NewString(value))
}

func (self Go) CompileErrorFragment() string {
		return string(class(self).GetStageCompileError(1).String())
}

func (self Go) SetCompileErrorFragment(value string) {
	class(self).SetStageCompileError(1, gd.NewString(value))
}

func (self Go) CompileErrorTesselationControl() string {
		return string(class(self).GetStageCompileError(2).String())
}

func (self Go) SetCompileErrorTesselationControl(value string) {
	class(self).SetStageCompileError(2, gd.NewString(value))
}

func (self Go) CompileErrorTesselationEvaluation() string {
		return string(class(self).GetStageCompileError(3).String())
}

func (self Go) SetCompileErrorTesselationEvaluation(value string) {
	class(self).SetStageCompileError(3, gd.NewString(value))
}

func (self Go) CompileErrorCompute() string {
		return string(class(self).GetStageCompileError(4).String())
}

func (self Go) SetCompileErrorCompute(value string) {
	class(self).SetStageCompileError(4, gd.NewString(value))
}

/*
Sets the SPIR-V [param bytecode] for the given shader [param stage]. Equivalent to setting one of [member bytecode_compute], [member bytecode_fragment], [member bytecode_tesselation_control], [member bytecode_tesselation_evaluation], [member bytecode_vertex].
*/
//go:nosplit
func (self class) SetStageBytecode(stage classdb.RenderingDeviceShaderStage, bytecode gd.PackedByteArray)  {
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	callframe.Arg(frame, discreet.Get(bytecode))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSPIRV.Bind_set_stage_bytecode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Equivalent to getting one of [member bytecode_compute], [member bytecode_fragment], [member bytecode_tesselation_control], [member bytecode_tesselation_evaluation], [member bytecode_vertex].
*/
//go:nosplit
func (self class) GetStageBytecode(stage classdb.RenderingDeviceShaderStage) gd.PackedByteArray {
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSPIRV.Bind_get_stage_bytecode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the compilation error message for the given shader [param stage] to [param compile_error]. Equivalent to setting one of [member compile_error_compute], [member compile_error_fragment], [member compile_error_tesselation_control], [member compile_error_tesselation_evaluation], [member compile_error_vertex].
*/
//go:nosplit
func (self class) SetStageCompileError(stage classdb.RenderingDeviceShaderStage, compile_error gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	callframe.Arg(frame, discreet.Get(compile_error))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSPIRV.Bind_set_stage_compile_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the compilation error message for the given shader [param stage]. Equivalent to getting one of [member compile_error_compute], [member compile_error_fragment], [member compile_error_tesselation_control], [member compile_error_tesselation_evaluation], [member compile_error_vertex].
*/
//go:nosplit
func (self class) GetStageCompileError(stage classdb.RenderingDeviceShaderStage) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderSPIRV.Bind_get_stage_compile_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsRDShaderSPIRV() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRDShaderSPIRV() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("RDShaderSPIRV", func(ptr gd.Object) any { return classdb.RDShaderSPIRV(ptr) })}
