package VisualShaderNodeParameter

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VisualShaderNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A parameter represents a variable in the shader which is set externally, i.e. from the [ShaderMaterial]. Parameters are exposed as properties in the [ShaderMaterial] and can be assigned from the Inspector or from a script.

*/
type Simple [1]classdb.VisualShaderNodeParameter
func (self Simple) SetParameterName(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParameterName(gc.String(name))
}
func (self Simple) GetParameterName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetParameterName(gc).String())
}
func (self Simple) SetQualifier(qualifier classdb.VisualShaderNodeParameterQualifier) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetQualifier(qualifier)
}
func (self Simple) GetQualifier() classdb.VisualShaderNodeParameterQualifier {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeParameterQualifier(Expert(self).GetQualifier())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeParameter
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsVisualShaderNodeParameter() Expert { return self[0].AsVisualShaderNodeParameter() }


//go:nosplit
func (self Simple) AsVisualShaderNodeParameter() Simple { return self[0].AsVisualShaderNodeParameter() }


//go:nosplit
func (self class) AsVisualShaderNode() VisualShaderNode.Expert { return self[0].AsVisualShaderNode() }


//go:nosplit
func (self Simple) AsVisualShaderNode() VisualShaderNode.Simple { return self[0].AsVisualShaderNode() }


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
func init() {classdb.Register("VisualShaderNodeParameter", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
