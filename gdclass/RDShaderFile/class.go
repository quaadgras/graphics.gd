package RDShaderFile

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
Compiled shader file in SPIR-V form.
See also [RDShaderSource]. [RDShaderFile] is only meant to be used with the [RenderingDevice] API. It should not be confused with Godot's own [Shader] resource, which is what Godot's various nodes use for high-level shader programming.

*/
type Go [1]classdb.RDShaderFile

/*
Sets the SPIR-V [param bytecode] that will be compiled for the specified [param version].
*/
func (self Go) SetBytecode(bytecode gdclass.RDShaderSPIRV) {
	class(self).SetBytecode(bytecode, gd.NewStringName(""))
}

/*
Returns the SPIR-V intermediate representation for the specified shader [param version].
*/
func (self Go) GetSpirv() gdclass.RDShaderSPIRV {
	return gdclass.RDShaderSPIRV(class(self).GetSpirv(gd.NewStringName("")))
}

/*
Returns the list of compiled versions for this shader.
*/
func (self Go) GetVersionList() gd.Array {
	return gd.Array(class(self).GetVersionList())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RDShaderFile
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDShaderFile"))
	return Go{classdb.RDShaderFile(object)}
}

func (self Go) BaseError() string {
		return string(class(self).GetBaseError().String())
}

func (self Go) SetBaseError(value string) {
	class(self).SetBaseError(gd.NewString(value))
}

/*
Sets the SPIR-V [param bytecode] that will be compiled for the specified [param version].
*/
//go:nosplit
func (self class) SetBytecode(bytecode gdclass.RDShaderSPIRV, version gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(bytecode[0])[0])
	callframe.Arg(frame, discreet.Get(version))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderFile.Bind_set_bytecode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the SPIR-V intermediate representation for the specified shader [param version].
*/
//go:nosplit
func (self class) GetSpirv(version gd.StringName) gdclass.RDShaderSPIRV {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(version))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderFile.Bind_get_spirv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.RDShaderSPIRV{classdb.RDShaderSPIRV(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Returns the list of compiled versions for this shader.
*/
//go:nosplit
func (self class) GetVersionList() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderFile.Bind_get_version_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBaseError(error gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(error))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderFile.Bind_set_base_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBaseError() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderFile.Bind_get_base_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
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
func init() {classdb.Register("RDShaderFile", func(ptr gd.Object) any { return classdb.RDShaderFile(ptr) })}
