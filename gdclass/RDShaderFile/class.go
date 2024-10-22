package RDShaderFile

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
Compiled shader file in SPIR-V form.
See also [RDShaderSource]. [RDShaderFile] is only meant to be used with the [RenderingDevice] API. It should not be confused with Godot's own [Shader] resource, which is what Godot's various nodes use for high-level shader programming.

*/
type Go [1]classdb.RDShaderFile

/*
Sets the SPIR-V [param bytecode] that will be compiled for the specified [param version].
*/
func (self Go) SetBytecode(bytecode gdclass.RDShaderSPIRV) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBytecode(bytecode, gc.StringName(""))
}

/*
Returns the SPIR-V intermediate representation for the specified shader [param version].
*/
func (self Go) GetSpirv() gdclass.RDShaderSPIRV {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.RDShaderSPIRV(class(self).GetSpirv(gc, gc.StringName("")))
}

/*
Returns the list of compiled versions for this shader.
*/
func (self Go) GetVersionList() gd.ArrayOf[gd.StringName] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.StringName](class(self).GetVersionList(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RDShaderFile
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RDShaderFile"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) BaseError() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetBaseError(gc).String())
}

func (self Go) SetBaseError(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBaseError(gc.String(value))
}

/*
Sets the SPIR-V [param bytecode] that will be compiled for the specified [param version].
*/
//go:nosplit
func (self class) SetBytecode(bytecode gdclass.RDShaderSPIRV, version gd.StringName)  {
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
func (self class) GetSpirv(ctx gd.Lifetime, version gd.StringName) gdclass.RDShaderSPIRV {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(version))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDShaderFile.Bind_get_spirv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.RDShaderSPIRV
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
func (self class) AsRDShaderFile() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRDShaderFile() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("RDShaderFile", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
