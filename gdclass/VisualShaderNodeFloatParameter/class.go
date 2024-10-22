package VisualShaderNodeFloatParameter

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNodeParameter"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Translated to [code]uniform float[/code] in the shader language.

*/
type Go [1]classdb.VisualShaderNodeFloatParameter
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeFloatParameter
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VisualShaderNodeFloatParameter"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Hint() classdb.VisualShaderNodeFloatParameterHint {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.VisualShaderNodeFloatParameterHint(class(self).GetHint())
}

func (self Go) SetHint(value classdb.VisualShaderNodeFloatParameterHint) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHint(value)
}

func (self Go) Min() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMin()))
}

func (self Go) SetMin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMin(gd.Float(value))
}

func (self Go) Max() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMax()))
}

func (self Go) SetMax(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMax(gd.Float(value))
}

func (self Go) Step() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetStep()))
}

func (self Go) SetStep(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStep(gd.Float(value))
}

func (self Go) DefaultValueEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDefaultValueEnabled())
}

func (self Go) SetDefaultValueEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDefaultValueEnabled(value)
}

func (self Go) DefaultValue() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDefaultValue()))
}

func (self Go) SetDefaultValue(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDefaultValue(gd.Float(value))
}

//go:nosplit
func (self class) SetHint(hint classdb.VisualShaderNodeFloatParameterHint)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hint)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFloatParameter.Bind_set_hint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHint() classdb.VisualShaderNodeFloatParameterHint {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeFloatParameterHint](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFloatParameter.Bind_get_hint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMin(value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFloatParameter.Bind_set_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFloatParameter.Bind_get_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMax(value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFloatParameter.Bind_set_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMax() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFloatParameter.Bind_get_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStep(value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFloatParameter.Bind_set_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStep() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFloatParameter.Bind_get_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDefaultValueEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFloatParameter.Bind_set_default_value_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDefaultValueEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFloatParameter.Bind_is_default_value_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDefaultValue(value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFloatParameter.Bind_set_default_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDefaultValue() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFloatParameter.Bind_get_default_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeFloatParameter() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeFloatParameter() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNodeParameter() VisualShaderNodeParameter.GD { return *((*VisualShaderNodeParameter.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeParameter() VisualShaderNodeParameter.Go { return *((*VisualShaderNodeParameter.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNode() VisualShaderNode.GD { return *((*VisualShaderNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNode() VisualShaderNode.Go { return *((*VisualShaderNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeParameter(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeParameter(), name)
	}
}
func init() {classdb.Register("VisualShaderNodeFloatParameter", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type Hint = classdb.VisualShaderNodeFloatParameterHint

const (
/*No hint used.*/
	HintNone Hint = 0
/*A range hint for scalar value, which limits possible input values between [member min] and [member max]. Translated to [code]hint_range(min, max)[/code] in shader code.*/
	HintRange Hint = 1
/*A range hint for scalar value with step, which limits possible input values between [member min] and [member max], with a step (increment) of [member step]). Translated to [code]hint_range(min, max, step)[/code] in shader code.*/
	HintRangeStep Hint = 2
/*Represents the size of the [enum Hint] enum.*/
	HintMax Hint = 3
)
