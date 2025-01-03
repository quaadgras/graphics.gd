package VisualShaderNodeTransformVecMult

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/VisualShaderNode"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A multiplication operation on a transform (4×4 matrix) and a vector, with support for different multiplication operators.
*/
type Instance [1]classdb.VisualShaderNodeTransformVecMult
type Any interface {
	gd.IsClass
	AsVisualShaderNodeTransformVecMult() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShaderNodeTransformVecMult

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeTransformVecMult"))
	return Instance{*(*classdb.VisualShaderNodeTransformVecMult)(unsafe.Pointer(&object))}
}

func (self Instance) Operator() classdb.VisualShaderNodeTransformVecMultOperator {
	return classdb.VisualShaderNodeTransformVecMultOperator(class(self).GetOperator())
}

func (self Instance) SetOperator(value classdb.VisualShaderNodeTransformVecMultOperator) {
	class(self).SetOperator(value)
}

//go:nosplit
func (self class) SetOperator(op classdb.VisualShaderNodeTransformVecMultOperator) {
	var frame = callframe.New()
	callframe.Arg(frame, op)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTransformVecMult.Bind_set_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOperator() classdb.VisualShaderNodeTransformVecMultOperator {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeTransformVecMultOperator](frame)
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}
func init() {
	classdb.Register("VisualShaderNodeTransformVecMult", func(ptr gd.Object) any {
		return [1]classdb.VisualShaderNodeTransformVecMult{*(*classdb.VisualShaderNodeTransformVecMult)(unsafe.Pointer(&ptr))}
	})
}

type Operator = classdb.VisualShaderNodeTransformVecMultOperator

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
