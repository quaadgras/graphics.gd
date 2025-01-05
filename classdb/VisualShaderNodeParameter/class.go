// Package VisualShaderNodeParameter provides methods for working with VisualShaderNodeParameter object instances.
package VisualShaderNodeParameter

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A parameter represents a variable in the shader which is set externally, i.e. from the [ShaderMaterial]. Parameters are exposed as properties in the [ShaderMaterial] and can be assigned from the Inspector or from a script.
*/
type Instance [1]gdclass.VisualShaderNodeParameter
type Any interface {
	gd.IsClass
	AsVisualShaderNodeParameter() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeParameter

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeParameter"))
	return Instance{*(*gdclass.VisualShaderNodeParameter)(unsafe.Pointer(&object))}
}

func (self Instance) ParameterName() string {
	return string(class(self).GetParameterName().String())
}

func (self Instance) SetParameterName(value string) {
	class(self).SetParameterName(gd.NewString(value))
}

func (self Instance) Qualifier() gdclass.VisualShaderNodeParameterQualifier {
	return gdclass.VisualShaderNodeParameterQualifier(class(self).GetQualifier())
}

func (self Instance) SetQualifier(value gdclass.VisualShaderNodeParameterQualifier) {
	class(self).SetQualifier(value)
}

//go:nosplit
func (self class) SetParameterName(name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParameter.Bind_set_parameter_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetParameterName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParameter.Bind_get_parameter_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetQualifier(qualifier gdclass.VisualShaderNodeParameterQualifier) {
	var frame = callframe.New()
	callframe.Arg(frame, qualifier)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParameter.Bind_set_qualifier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetQualifier() gdclass.VisualShaderNodeParameterQualifier {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeParameterQualifier](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParameter.Bind_get_qualifier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeParameter() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeParameter() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNode() VisualShaderNode.Advanced {
	return *((*VisualShaderNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNode() VisualShaderNode.Instance {
	return *((*VisualShaderNode.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(VisualShaderNode.Advanced(self.AsVisualShaderNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNode.Instance(self.AsVisualShaderNode()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeParameter", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeParameter{*(*gdclass.VisualShaderNodeParameter)(unsafe.Pointer(&ptr))}
	})
}

type Qualifier = gdclass.VisualShaderNodeParameterQualifier

const (
	/*The parameter will be tied to the [ShaderMaterial] using this shader.*/
	QualNone Qualifier = 0
	/*The parameter will use a global value, defined in Project Settings.*/
	QualGlobal Qualifier = 1
	/*The parameter will be tied to the node with attached [ShaderMaterial] using this shader.*/
	QualInstance Qualifier = 2
	/*Represents the size of the [enum Qualifier] enum.*/
	QualMax Qualifier = 3
)
