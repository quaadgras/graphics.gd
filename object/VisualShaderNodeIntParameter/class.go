package VisualShaderNodeIntParameter

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VisualShaderNodeParameter"
import "grow.graphics/gd/object/VisualShaderNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A [VisualShaderNodeParameter] of type [int]. Offers additional customization for range of accepted values.

*/
type Simple [1]classdb.VisualShaderNodeIntParameter
func (self Simple) SetHint(hint classdb.VisualShaderNodeIntParameterHint) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHint(hint)
}
func (self Simple) GetHint() classdb.VisualShaderNodeIntParameterHint {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeIntParameterHint(Expert(self).GetHint())
}
func (self Simple) SetMin(value int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMin(gd.Int(value))
}
func (self Simple) GetMin() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMin()))
}
func (self Simple) SetMax(value int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMax(gd.Int(value))
}
func (self Simple) GetMax() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMax()))
}
func (self Simple) SetStep(value int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStep(gd.Int(value))
}
func (self Simple) GetStep() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetStep()))
}
func (self Simple) SetDefaultValueEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDefaultValueEnabled(enabled)
}
func (self Simple) IsDefaultValueEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDefaultValueEnabled())
}
func (self Simple) SetDefaultValue(value int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDefaultValue(gd.Int(value))
}
func (self Simple) GetDefaultValue() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetDefaultValue()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeIntParameter
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetHint(hint classdb.VisualShaderNodeIntParameterHint)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hint)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeIntParameter.Bind_set_hint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHint() classdb.VisualShaderNodeIntParameterHint {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeIntParameterHint](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeIntParameter.Bind_get_hint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMin(value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeIntParameter.Bind_set_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMin() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeIntParameter.Bind_get_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMax(value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeIntParameter.Bind_set_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMax() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeIntParameter.Bind_get_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStep(value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeIntParameter.Bind_set_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStep() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeIntParameter.Bind_get_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeIntParameter.Bind_set_default_value_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDefaultValueEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeIntParameter.Bind_is_default_value_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDefaultValue(value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeIntParameter.Bind_set_default_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDefaultValue() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeIntParameter.Bind_get_default_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeIntParameter() Expert { return self[0].AsVisualShaderNodeIntParameter() }


//go:nosplit
func (self Simple) AsVisualShaderNodeIntParameter() Simple { return self[0].AsVisualShaderNodeIntParameter() }


//go:nosplit
func (self class) AsVisualShaderNodeParameter() VisualShaderNodeParameter.Expert { return self[0].AsVisualShaderNodeParameter() }


//go:nosplit
func (self Simple) AsVisualShaderNodeParameter() VisualShaderNodeParameter.Simple { return self[0].AsVisualShaderNodeParameter() }


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
func init() {classdb.Register("VisualShaderNodeIntParameter", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Hint = classdb.VisualShaderNodeIntParameterHint

const (
/*The parameter will not constrain its value.*/
	HintNone Hint = 0
/*The parameter's value must be within the specified [member min]/[member max] range.*/
	HintRange Hint = 1
/*The parameter's value must be within the specified range, with the given [member step] between values.*/
	HintRangeStep Hint = 2
/*Represents the size of the [enum Hint] enum.*/
	HintMax Hint = 3
)
