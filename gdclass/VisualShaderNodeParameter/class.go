package VisualShaderNodeParameter

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A parameter represents a variable in the shader which is set externally, i.e. from the [ShaderMaterial]. Parameters are exposed as properties in the [ShaderMaterial] and can be assigned from the Inspector or from a script.

*/
type Go [1]classdb.VisualShaderNodeParameter
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeParameter
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VisualShaderNodeParameter"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) ParameterName() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetParameterName(gc).String())
}

func (self Go) SetParameterName(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParameterName(gc.String(value))
}

func (self Go) Qualifier() classdb.VisualShaderNodeParameterQualifier {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.VisualShaderNodeParameterQualifier(class(self).GetQualifier())
}

func (self Go) SetQualifier(value classdb.VisualShaderNodeParameterQualifier) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetQualifier(value)
}

//go:nosplit
func (self class) SetParameterName(name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParameter.Bind_set_parameter_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParameterName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParameter.Bind_get_parameter_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetQualifier(qualifier classdb.VisualShaderNodeParameterQualifier)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, qualifier)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParameter.Bind_set_qualifier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetQualifier() classdb.VisualShaderNodeParameterQualifier {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeParameterQualifier](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParameter.Bind_get_qualifier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeParameter() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeParameter() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNode() VisualShaderNode.GD { return *((*VisualShaderNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNode() VisualShaderNode.Go { return *((*VisualShaderNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}
func init() {classdb.Register("VisualShaderNodeParameter", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type Qualifier = classdb.VisualShaderNodeParameterQualifier

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
