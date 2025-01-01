package VisualShaderNodeColorParameter

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/VisualShaderNodeParameter"
import "graphics.gd/objects/VisualShaderNode"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Color"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Translated to [code]uniform vec4[/code] in the shader language.
*/
type Instance [1]classdb.VisualShaderNodeColorParameter
type Any interface {
	gd.IsClass
	AsVisualShaderNodeColorParameter() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShaderNodeColorParameter

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeColorParameter"))
	return Instance{classdb.VisualShaderNodeColorParameter(object)}
}

func (self Instance) DefaultValueEnabled() bool {
	return bool(class(self).IsDefaultValueEnabled())
}

func (self Instance) SetDefaultValueEnabled(value bool) {
	class(self).SetDefaultValueEnabled(value)
}

func (self Instance) DefaultValue() Color.RGBA {
	return Color.RGBA(class(self).GetDefaultValue())
}

func (self Instance) SetDefaultValue(value Color.RGBA) {
	class(self).SetDefaultValue(gd.Color(value))
}

//go:nosplit
func (self class) SetDefaultValueEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeColorParameter.Bind_set_default_value_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDefaultValueEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeColorParameter.Bind_is_default_value_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDefaultValue(value gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeColorParameter.Bind_set_default_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDefaultValue() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeColorParameter.Bind_get_default_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeColorParameter() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeColorParameter() Instance {
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNodeParameter(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNodeParameter(), name)
	}
}
func init() {
	classdb.Register("VisualShaderNodeColorParameter", func(ptr gd.Object) any {
		return [1]classdb.VisualShaderNodeColorParameter{classdb.VisualShaderNodeColorParameter(ptr)}
	})
}
