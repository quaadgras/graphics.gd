package VisualShaderNodeFloatOp

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
Applies [member operator] to two floating-point inputs: [code]a[/code] and [code]b[/code].

*/
type Simple [1]classdb.VisualShaderNodeFloatOp
func (self Simple) SetOperator(op classdb.VisualShaderNodeFloatOpOperator) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOperator(op)
}
func (self Simple) GetOperator() classdb.VisualShaderNodeFloatOpOperator {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeFloatOpOperator(Expert(self).GetOperator())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeFloatOp
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetOperator(op classdb.VisualShaderNodeFloatOpOperator)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, op)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFloatOp.Bind_set_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOperator() classdb.VisualShaderNodeFloatOpOperator {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeFloatOpOperator](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFloatOp.Bind_get_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeFloatOp() Expert { return self[0].AsVisualShaderNodeFloatOp() }


//go:nosplit
func (self Simple) AsVisualShaderNodeFloatOp() Simple { return self[0].AsVisualShaderNodeFloatOp() }


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
func init() {classdb.Register("VisualShaderNodeFloatOp", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Operator = classdb.VisualShaderNodeFloatOpOperator

const (
/*Sums two numbers using [code]a + b[/code].*/
	OpAdd Operator = 0
/*Subtracts two numbers using [code]a - b[/code].*/
	OpSub Operator = 1
/*Multiplies two numbers using [code]a * b[/code].*/
	OpMul Operator = 2
/*Divides two numbers using [code]a / b[/code].*/
	OpDiv Operator = 3
/*Calculates the remainder of two numbers. Translates to [code]mod(a, b)[/code] in the Godot Shader Language.*/
	OpMod Operator = 4
/*Raises the [code]a[/code] to the power of [code]b[/code]. Translates to [code]pow(a, b)[/code] in the Godot Shader Language.*/
	OpPow Operator = 5
/*Returns the greater of two numbers. Translates to [code]max(a, b)[/code] in the Godot Shader Language.*/
	OpMax Operator = 6
/*Returns the lesser of two numbers. Translates to [code]min(a, b)[/code] in the Godot Shader Language.*/
	OpMin Operator = 7
/*Returns the arc-tangent of the parameters. Translates to [code]atan(a, b)[/code] in the Godot Shader Language.*/
	OpAtan2 Operator = 8
/*Generates a step function by comparing [code]b[/code](x) to [code]a[/code](edge). Returns 0.0 if [code]x[/code] is smaller than [code]edge[/code] and otherwise 1.0. Translates to [code]step(a, b)[/code] in the Godot Shader Language.*/
	OpStep Operator = 9
/*Represents the size of the [enum Operator] enum.*/
	OpEnumSize Operator = 10
)
