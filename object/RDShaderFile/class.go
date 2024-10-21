package RDShaderFile

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
Compiled shader file in SPIR-V form.
See also [RDShaderSource]. [RDShaderFile] is only meant to be used with the [RenderingDevice] API. It should not be confused with Godot's own [Shader] resource, which is what Godot's various nodes use for high-level shader programming.

*/
type Simple [1]classdb.RDShaderFile
func (self Simple) SetBytecode(bytecode [1]classdb.RDShaderSPIRV, version string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBytecode(bytecode, gc.StringName(version))
}
func (self Simple) GetSpirv(version string) [1]classdb.RDShaderSPIRV {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.RDShaderSPIRV(Expert(self).GetSpirv(gc, gc.StringName(version)))
}
func (self Simple) GetVersionList() gd.ArrayOf[gd.StringName] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.StringName](Expert(self).GetVersionList(gc))
}
func (self Simple) SetBaseError(error string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBaseError(gc.String(error))
}
func (self Simple) GetBaseError() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetBaseError(gc).String())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RDShaderFile
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Sets the SPIR-V [param bytecode] that will be compiled for the specified [param version].
*/
//go:nosplit
func (self class) SetBytecode(bytecode object.RDShaderSPIRV, version gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bytecode[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(version))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDShaderFile.Bind_set_bytecode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the SPIR-V intermediate representation for the specified shader [param version].
*/
//go:nosplit
func (self class) GetSpirv(ctx gd.Lifetime, version gd.StringName) object.RDShaderSPIRV {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(version))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDShaderFile.Bind_get_spirv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.RDShaderSPIRV
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the list of compiled versions for this shader.
*/
//go:nosplit
func (self class) GetVersionList(ctx gd.Lifetime) gd.ArrayOf[gd.StringName] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDShaderFile.Bind_get_version_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.StringName](ret)
}
//go:nosplit
func (self class) SetBaseError(error gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(error))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDShaderFile.Bind_set_base_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBaseError(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDShaderFile.Bind_get_base_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsRDShaderFile() Expert { return self[0].AsRDShaderFile() }


//go:nosplit
func (self Simple) AsRDShaderFile() Simple { return self[0].AsRDShaderFile() }


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
func init() {classdb.Register("RDShaderFile", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
