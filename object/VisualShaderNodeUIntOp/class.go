package VisualShaderNodeUIntOp

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
Applies [member operator] to two unsigned integer inputs: [code]a[/code] and [code]b[/code].

*/
type Simple [1]classdb.VisualShaderNodeUIntOp
func (self Simple) SetOperator(op classdb.VisualShaderNodeUIntOpOperator) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOperator(op)
}
func (self Simple) GetOperator() classdb.VisualShaderNodeUIntOpOperator {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeUIntOpOperator(Expert(self).GetOperator())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeUIntOp
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetOperator(op classdb.VisualShaderNodeUIntOpOperator)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, op)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeUIntOp.Bind_set_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOperator() classdb.VisualShaderNodeUIntOpOperator {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeUIntOpOperator](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeUIntOp.Bind_get_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeUIntOp() Expert { return self[0].AsVisualShaderNodeUIntOp() }


//go:nosplit
func (self Simple) AsVisualShaderNodeUIntOp() Simple { return self[0].AsVisualShaderNodeUIntOp() }


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
func init() {classdb.Register("VisualShaderNodeUIntOp", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Operator = classdb.VisualShaderNodeUIntOpOperator

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
