// Package VisualShaderNodeFloatParameter provides methods for working with VisualShaderNodeFloatParameter object instances.
package VisualShaderNodeFloatParameter

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
import "graphics.gd/classdb/VisualShaderNodeParameter"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

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

/*
Translated to [code]uniform float[/code] in the shader language.
*/
type Instance [1]gdclass.VisualShaderNodeFloatParameter

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeFloatParameter() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeFloatParameter

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeFloatParameter"))
	casted := Instance{*(*gdclass.VisualShaderNodeFloatParameter)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Hint() gdclass.VisualShaderNodeFloatParameterHint {
	return gdclass.VisualShaderNodeFloatParameterHint(class(self).GetHint())
}

func (self Instance) SetHint(value gdclass.VisualShaderNodeFloatParameterHint) {
	class(self).SetHint(value)
}

func (self Instance) Min() Float.X {
	return Float.X(Float.X(class(self).GetMin()))
}

func (self Instance) SetMin(value Float.X) {
	class(self).SetMin(gd.Float(value))
}

func (self Instance) Max() Float.X {
	return Float.X(Float.X(class(self).GetMax()))
}

func (self Instance) SetMax(value Float.X) {
	class(self).SetMax(gd.Float(value))
}

func (self Instance) Step() Float.X {
	return Float.X(Float.X(class(self).GetStep()))
}

func (self Instance) SetStep(value Float.X) {
	class(self).SetStep(gd.Float(value))
}

func (self Instance) DefaultValueEnabled() bool {
	return bool(class(self).IsDefaultValueEnabled())
}

func (self Instance) SetDefaultValueEnabled(value bool) {
	class(self).SetDefaultValueEnabled(value)
}

func (self Instance) DefaultValue() Float.X {
	return Float.X(Float.X(class(self).GetDefaultValue()))
}

func (self Instance) SetDefaultValue(value Float.X) {
	class(self).SetDefaultValue(gd.Float(value))
}

//go:nosplit
func (self class) SetHint(hint gdclass.VisualShaderNodeFloatParameterHint) { //gd:VisualShaderNodeFloatParameter.set_hint
	var frame = callframe.New()
	callframe.Arg(frame, hint)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFloatParameter.Bind_set_hint, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHint() gdclass.VisualShaderNodeFloatParameterHint { //gd:VisualShaderNodeFloatParameter.get_hint
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeFloatParameterHint](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFloatParameter.Bind_get_hint, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMin(value gd.Float) { //gd:VisualShaderNodeFloatParameter.set_min
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFloatParameter.Bind_set_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMin() gd.Float { //gd:VisualShaderNodeFloatParameter.get_min
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFloatParameter.Bind_get_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMax(value gd.Float) { //gd:VisualShaderNodeFloatParameter.set_max
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFloatParameter.Bind_set_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMax() gd.Float { //gd:VisualShaderNodeFloatParameter.get_max
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFloatParameter.Bind_get_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStep(value gd.Float) { //gd:VisualShaderNodeFloatParameter.set_step
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFloatParameter.Bind_set_step, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStep() gd.Float { //gd:VisualShaderNodeFloatParameter.get_step
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFloatParameter.Bind_get_step, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDefaultValueEnabled(enabled bool) { //gd:VisualShaderNodeFloatParameter.set_default_value_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFloatParameter.Bind_set_default_value_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDefaultValueEnabled() bool { //gd:VisualShaderNodeFloatParameter.is_default_value_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFloatParameter.Bind_is_default_value_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDefaultValue(value gd.Float) { //gd:VisualShaderNodeFloatParameter.set_default_value
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFloatParameter.Bind_set_default_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDefaultValue() gd.Float { //gd:VisualShaderNodeFloatParameter.get_default_value
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFloatParameter.Bind_get_default_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeFloatParameter() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeFloatParameter() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNodeParameter() VisualShaderNodeParameter.Advanced {
	return *((*VisualShaderNodeParameter.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeParameter() VisualShaderNodeParameter.Instance {
	return *((*VisualShaderNodeParameter.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(VisualShaderNodeParameter.Advanced(self.AsVisualShaderNodeParameter()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNodeParameter.Instance(self.AsVisualShaderNodeParameter()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeFloatParameter", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeFloatParameter{*(*gdclass.VisualShaderNodeFloatParameter)(unsafe.Pointer(&ptr))}
	})
}

type Hint = gdclass.VisualShaderNodeFloatParameterHint //gd:VisualShaderNodeFloatParameter.Hint

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
