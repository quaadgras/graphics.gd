package VisualShaderNodeIntOp

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Applies [member operator] to two integer inputs: [code]a[/code] and [code]b[/code].
*/
type Instance [1]classdb.VisualShaderNodeIntOp

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShaderNodeIntOp

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeIntOp"))
	return Instance{classdb.VisualShaderNodeIntOp(object)}
}

func (self Instance) Operator() classdb.VisualShaderNodeIntOpOperator {
	return classdb.VisualShaderNodeIntOpOperator(class(self).GetOperator())
}

func (self Instance) SetOperator(value classdb.VisualShaderNodeIntOpOperator) {
	class(self).SetOperator(value)
}

//go:nosplit
func (self class) SetOperator(op classdb.VisualShaderNodeIntOpOperator) {
	var frame = callframe.New()
	callframe.Arg(frame, op)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeIntOp.Bind_set_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOperator() classdb.VisualShaderNodeIntOpOperator {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeIntOpOperator](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeIntOp.Bind_get_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeIntOp() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShaderNodeIntOp() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("VisualShaderNodeIntOp", func(ptr gd.Object) any { return classdb.VisualShaderNodeIntOp(ptr) })
}

type Operator = classdb.VisualShaderNodeIntOpOperator

const (
	/*Sums two numbers using [code]a + b[/code].*/
	OpAdd Operator = 0
	/*Subtracts two numbers using [code]a - b[/code].*/
	OpSub Operator = 1
	/*Multiplies two numbers using [code]a * b[/code].*/
	OpMul Operator = 2
	/*Divides two numbers using [code]a / b[/code].*/
	OpDiv Operator = 3
	/*Calculates the remainder of two numbers using [code]a % b[/code].*/
	OpMod Operator = 4
	/*Returns the greater of two numbers. Translates to [code]max(a, b)[/code] in the Godot Shader Language.*/
	OpMax Operator = 5
	/*Returns the lesser of two numbers. Translates to [code]max(a, b)[/code] in the Godot Shader Language.*/
	OpMin Operator = 6
	/*Returns the result of bitwise [code]AND[/code] operation on the integer. Translates to [code]a & b[/code] in the Godot Shader Language.*/
	OpBitwiseAnd Operator = 7
	/*Returns the result of bitwise [code]OR[/code] operation for two integers. Translates to [code]a | b[/code] in the Godot Shader Language.*/
	OpBitwiseOr Operator = 8
	/*Returns the result of bitwise [code]XOR[/code] operation for two integers. Translates to [code]a ^ b[/code] in the Godot Shader Language.*/
	OpBitwiseXor Operator = 9
	/*Returns the result of bitwise left shift operation on the integer. Translates to [code]a << b[/code] in the Godot Shader Language.*/
	OpBitwiseLeftShift Operator = 10
	/*Returns the result of bitwise right shift operation on the integer. Translates to [code]a >> b[/code] in the Godot Shader Language.*/
	OpBitwiseRightShift Operator = 11
	/*Represents the size of the [enum Operator] enum.*/
	OpEnumSize Operator = 12
)
