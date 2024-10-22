package RDShaderSPIRV

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RDShaderSPIRV"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) BytecodeVertex() []byte {
	gc := gd.GarbageCollector(); _ = gc
		return []byte(class(self).GetStageBytecode(gc,0).Bytes())
}

func (self Go) SetBytecodeVertex(value []byte) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStageBytecode(0, gc.PackedByteSlice(value))
}

func (self Go) BytecodeFragment() []byte {
	gc := gd.GarbageCollector(); _ = gc
		return []byte(class(self).GetStageBytecode(gc,1).Bytes())
}

func (self Go) SetBytecodeFragment(value []byte) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStageBytecode(1, gc.PackedByteSlice(value))
}

func (self Go) BytecodeTesselationControl() []byte {
	gc := gd.GarbageCollector(); _ = gc
		return []byte(class(self).GetStageBytecode(gc,2).Bytes())
}

func (self Go) SetBytecodeTesselationControl(value []byte) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStageBytecode(2, gc.PackedByteSlice(value))
}

func (self Go) BytecodeTesselationEvaluation() []byte {
	gc := gd.GarbageCollector(); _ = gc
		return []byte(class(self).GetStageBytecode(gc,3).Bytes())
}

func (self Go) SetBytecodeTesselationEvaluation(value []byte) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStageBytecode(3, gc.PackedByteSlice(value))
}

func (self Go) BytecodeCompute() []byte {
	gc := gd.GarbageCollector(); _ = gc
		return []byte(class(self).GetStageBytecode(gc,4).Bytes())
}

func (self Go) SetBytecodeCompute(value []byte) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStageBytecode(4, gc.PackedByteSlice(value))
}

func (self Go) CompileErrorVertex() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetStageCompileError(gc,0).String())
}

func (self Go) SetCompileErrorVertex(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStageCompileError(0, gc.String(value))
}

func (self Go) CompileErrorFragment() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetStageCompileError(gc,1).String())
}

func (self Go) SetCompileErrorFragment(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStageCompileError(1, gc.String(value))
}

func (self Go) CompileErrorTesselationControl() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetStageCompileError(gc,2).String())
}

func (self Go) SetCompileErrorTesselationControl(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStageCompileError(2, gc.String(value))
}

func (self Go) CompileErrorTesselationEvaluation() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetStageCompileError(gc,3).String())
}

func (self Go) SetCompileErrorTesselationEvaluation(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStageCompileError(3, gc.String(value))
}

func (self Go) CompileErrorCompute() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetStageCompileError(gc,4).String())
}

func (self Go) SetCompileErrorCompute(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStageCompileError(4, gc.String(value))
}

/*
Sets the SPIR-V [param bytecode] for the given shader [param stage]. Equivalent to setting one of [member bytecode_compute], [member bytecode_fragment], [member bytecode_tesselation_control], [member bytecode_tesselation_evaluation], [member bytecode_vertex].
*/
//go:nosplit
func (self class) SetStageBytecode(stage classdb.RenderingDeviceShaderStage, bytecode gd.PackedByteArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	callframe.Arg(frame, mmm.Get(bytecode))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDShaderSPIRV.Bind_set_stage_bytecode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Equivalent to getting one of [member bytecode_compute], [member bytecode_fragment], [member bytecode_tesselation_control], [member bytecode_tesselation_evaluation], [member bytecode_vertex].
*/
//go:nosplit
func (self class) GetStageBytecode(ctx gd.Lifetime, stage classdb.RenderingDeviceShaderStage) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDShaderSPIRV.Bind_get_stage_bytecode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the compilation error message for the given shader [param stage] to [param compile_error]. Equivalent to setting one of [member compile_error_compute], [member compile_error_fragment], [member compile_error_tesselation_control], [member compile_error_tesselation_evaluation], [member compile_error_vertex].
*/
//go:nosplit
func (self class) SetStageCompileError(stage classdb.RenderingDeviceShaderStage, compile_error gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	callframe.Arg(frame, mmm.Get(compile_error))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDShaderSPIRV.Bind_set_stage_compile_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the compilation error message for the given shader [param stage]. Equivalent to getting one of [member compile_error_compute], [member compile_error_fragment], [member compile_error_tesselation_control], [member compile_error_tesselation_evaluation], [member compile_error_vertex].
*/
//go:nosplit
func (self class) GetStageCompileError(ctx gd.Lifetime, stage classdb.RenderingDeviceShaderStage) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stage)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDShaderSPIRV.Bind_get_stage_compile_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
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
func init() {classdb.Register("RDShaderSPIRV", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
