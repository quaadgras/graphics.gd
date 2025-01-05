// Package VisualShaderNodeBooleanConstant provides methods for working with VisualShaderNodeBooleanConstant object instances.
package VisualShaderNodeBooleanConstant

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/VisualShaderNodeConstant"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Has only one output port and no inputs.
Translated to [code skip-lint]bool[/code] in the shader language.
*/
type Instance [1]gdclass.VisualShaderNodeBooleanConstant
type Any interface {
	gd.IsClass
	AsVisualShaderNodeBooleanConstant() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeBooleanConstant

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeBooleanConstant"))
	return Instance{*(*gdclass.VisualShaderNodeBooleanConstant)(unsafe.Pointer(&object))}
}

func (self Instance) Constant() bool {
	return bool(class(self).GetConstant())
}

func (self Instance) SetConstant(value bool) {
	class(self).SetConstant(value)
}

//go:nosplit
func (self class) SetConstant(constant bool) {
	var frame = callframe.New()
	callframe.Arg(frame, constant)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeBooleanConstant.Bind_set_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetConstant() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeBooleanConstant.Bind_get_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeBooleanConstant() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeBooleanConstant() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNodeConstant() VisualShaderNodeConstant.Advanced {
	return *((*VisualShaderNodeConstant.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeConstant() VisualShaderNodeConstant.Instance {
	return *((*VisualShaderNodeConstant.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(VisualShaderNodeConstant.Advanced(self.AsVisualShaderNodeConstant()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNodeConstant.Instance(self.AsVisualShaderNodeConstant()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeBooleanConstant", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeBooleanConstant{*(*gdclass.VisualShaderNodeBooleanConstant)(unsafe.Pointer(&ptr))}
	})
}
