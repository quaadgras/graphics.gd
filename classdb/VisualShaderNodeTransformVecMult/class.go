// Package VisualShaderNodeTransformVecMult provides methods for working with VisualShaderNodeTransformVecMult object instances.
package VisualShaderNodeTransformVecMult

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
A multiplication operation on a transform (4×4 matrix) and a vector, with support for different multiplication operators.
*/
type Instance [1]gdclass.VisualShaderNodeTransformVecMult
type Any interface {
	gd.IsClass
	AsVisualShaderNodeTransformVecMult() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeTransformVecMult

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeTransformVecMult"))
	return Instance{*(*gdclass.VisualShaderNodeTransformVecMult)(unsafe.Pointer(&object))}
}

func (self Instance) Operator() gdclass.VisualShaderNodeTransformVecMultOperator {
	return gdclass.VisualShaderNodeTransformVecMultOperator(class(self).GetOperator())
}

func (self Instance) SetOperator(value gdclass.VisualShaderNodeTransformVecMultOperator) {
	class(self).SetOperator(value)
}

//go:nosplit
func (self class) SetOperator(op gdclass.VisualShaderNodeTransformVecMultOperator) {
	var frame = callframe.New()
	callframe.Arg(frame, op)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTransformVecMult.Bind_set_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOperator() gdclass.VisualShaderNodeTransformVecMultOperator {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeTransformVecMultOperator](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTransformVecMult.Bind_get_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeTransformVecMult() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeTransformVecMult() Instance {
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
	gdclass.Register("VisualShaderNodeTransformVecMult", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeTransformVecMult{*(*gdclass.VisualShaderNodeTransformVecMult)(unsafe.Pointer(&ptr))}
	})
}

type Operator = gdclass.VisualShaderNodeTransformVecMultOperator

const (
	/*Multiplies transform [code]a[/code] by the vector [code]b[/code].*/
	OpAxb Operator = 0
	/*Multiplies vector [code]b[/code] by the transform [code]a[/code].*/
	OpBxa Operator = 1
	/*Multiplies transform [code]a[/code] by the vector [code]b[/code], skipping the last row and column of the transform.*/
	Op3x3Axb Operator = 2
	/*Multiplies vector [code]b[/code] by the transform [code]a[/code], skipping the last row and column of the transform.*/
	Op3x3Bxa Operator = 3
	/*Represents the size of the [enum Operator] enum.*/
	OpMax Operator = 4
)
