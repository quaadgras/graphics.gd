package RDShaderSPIRV

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[RDShaderSPIRV] represents a [RDShaderFile]'s [url=https://www.khronos.org/spir/]SPIR-V[/url] code for various shader stages, as well as possible compilation error messages. SPIR-V is a low-level intermediate shader representation. This intermediate representation is not used directly by GPUs for rendering, but it can be compiled into binary shaders that GPUs can understand. Unlike compiled shaders, SPIR-V is portable across GPU models and driver versions.
This object is used by [RenderingDevice].

*/
type Simple [1]classdb.RDShaderSPIRV
func (self Simple) SetStageBytecode(stage classdb.RenderingDeviceShaderStage, bytecode []byte) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStageBytecode(stage, gc.PackedByteSlice(bytecode))
}
func (self Simple) GetStageBytecode(stage classdb.RenderingDeviceShaderStage) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).GetStageBytecode(gc, stage).Bytes())
}
func (self Simple) SetStageCompileError(stage classdb.RenderingDeviceShaderStage, compile_error string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStageCompileError(stage, gc.String(compile_error))
}
func (self Simple) GetStageCompileError(stage classdb.RenderingDeviceShaderStage) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetStageCompileError(gc, stage).String())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RDShaderSPIRV
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsRDShaderSPIRV() Expert { return self[0].AsRDShaderSPIRV() }


//go:nosplit
func (self Simple) AsRDShaderSPIRV() Simple { return self[0].AsRDShaderSPIRV() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("RDShaderSPIRV", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
