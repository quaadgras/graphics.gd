package RDShaderFile

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Compiled shader file in SPIR-V form.
See also [RDShaderSource]. [RDShaderFile] is only meant to be used with the [RenderingDevice] API. It should not be confused with Godot's own [Shader] resource, which is what Godot's various nodes use for high-level shader programming.
*/
type Instance [1]classdb.RDShaderFile
type Any interface {
	gd.IsClass
	AsRDShaderFile() Instance
}

/*
Sets the SPIR-V [param bytecode] that will be compiled for the specified [param version].
*/
func (self Instance) SetBytecode(bytecode objects.RDShaderSPIRV) {
	class(self).SetBytecode(bytecode, gd.NewStringName(""))
}

/*
Returns the SPIR-V intermediate representation for the specified shader [param version].
*/
func (self Instance) GetSpirv() objects.RDShaderSPIRV {
	return objects.RDShaderSPIRV(class(self).GetSpirv(gd.NewStringName("")))
}

/*
Returns the list of compiled versions for this shader.
*/
func (self Instance) GetVersionList() gd.Array {
	return gd.Array(class(self).GetVersionList())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RDShaderFile

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDShaderFile"))
	return Instance{classdb.RDShaderFile(object)}
}

func (self Instance) BaseError() string {
	return string(class(self).GetBaseError().String())
}

func (self Instance) SetBaseError(value string) {
	class(self).SetBaseError(gd.NewString(value))
}

/*
Sets the SPIR-V [param bytecode] that will be compiled for the specified [param version].
*/
//go:nosplit
func (self class) SetBytecode(bytecode objects.RDShaderSPIRV, version gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(bytecode[0])[0])
	callframe.Arg(frame, pointers.Get(version))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderFile.Bind_set_bytecode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the SPIR-V intermediate representation for the specified shader [param version].
*/
//go:nosplit
func (self class) GetSpirv(version gd.StringName) objects.RDShaderSPIRV {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(version))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderFile.Bind_get_spirv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.RDShaderSPIRV{classdb.RDShaderSPIRV(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
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
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBaseError(error gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(error))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderFile.Bind_set_base_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBaseError() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderFile.Bind_get_base_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsRDShaderFile() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRDShaderFile() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("RDShaderFile", func(ptr gd.Object) any { return [1]classdb.RDShaderFile{classdb.RDShaderFile(ptr)} })
}
