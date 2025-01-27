// Package RDShaderFile provides methods for working with RDShaderFile object instances.
package RDShaderFile

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/classdb/Resource"

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

/*
Compiled shader file in SPIR-V form.
See also [RDShaderSource]. [RDShaderFile] is only meant to be used with the [RenderingDevice] API. It should not be confused with Godot's own [Shader] resource, which is what Godot's various nodes use for high-level shader programming.
*/
type Instance [1]gdclass.RDShaderFile

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRDShaderFile() Instance
}

/*
Sets the SPIR-V [param bytecode] that will be compiled for the specified [param version].
*/
func (self Instance) SetBytecode(bytecode [1]gdclass.RDShaderSPIRV) { //gd:RDShaderFile.set_bytecode
	class(self).SetBytecode(bytecode, gd.NewStringName(""))
}

/*
Returns the SPIR-V intermediate representation for the specified shader [param version].
*/
func (self Instance) GetSpirv() [1]gdclass.RDShaderSPIRV { //gd:RDShaderFile.get_spirv
	return [1]gdclass.RDShaderSPIRV(class(self).GetSpirv(gd.NewStringName("")))
}

/*
Returns the list of compiled versions for this shader.
*/
func (self Instance) GetVersionList() []string { //gd:RDShaderFile.get_version_list
	return []string(gd.ArrayAs[[]string](gd.InternalArray(class(self).GetVersionList())))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RDShaderFile

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDShaderFile"))
	casted := Instance{*(*gdclass.RDShaderFile)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) BaseError() string {
	return string(class(self).GetBaseError().String())
}

func (self Instance) SetBaseError(value string) {
	class(self).SetBaseError(String.New(value))
}

/*
Sets the SPIR-V [param bytecode] that will be compiled for the specified [param version].
*/
//go:nosplit
func (self class) SetBytecode(bytecode [1]gdclass.RDShaderSPIRV, version gd.StringName) { //gd:RDShaderFile.set_bytecode
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(bytecode[0])[0])
	callframe.Arg(frame, pointers.Get(version))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderFile.Bind_set_bytecode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the SPIR-V intermediate representation for the specified shader [param version].
*/
//go:nosplit
func (self class) GetSpirv(version gd.StringName) [1]gdclass.RDShaderSPIRV { //gd:RDShaderFile.get_spirv
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(version))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderFile.Bind_get_spirv, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.RDShaderSPIRV{gd.PointerWithOwnershipTransferredToGo[gdclass.RDShaderSPIRV](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the list of compiled versions for this shader.
*/
//go:nosplit
func (self class) GetVersionList() Array.Contains[gd.StringName] { //gd:RDShaderFile.get_version_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderFile.Bind_get_version_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.StringName]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBaseError(error String.Readable) { //gd:RDShaderFile.set_base_error
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(error)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderFile.Bind_set_base_error, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBaseError() String.Readable { //gd:RDShaderFile.get_base_error
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDShaderFile.Bind_get_base_error, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
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
	gdclass.Register("RDShaderFile", func(ptr gd.Object) any {
		return [1]gdclass.RDShaderFile{*(*gdclass.RDShaderFile)(unsafe.Pointer(&ptr))}
	})
}
