package VisualShaderNodeVectorOp

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VisualShaderNodeVectorBase"
import "grow.graphics/gd/object/VisualShaderNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A visual shader node for use of vector operators. Operates on vector [code]a[/code] and vector [code]b[/code].

*/
type Simple [1]classdb.VisualShaderNodeVectorOp
func (self Simple) SetOperator(op classdb.VisualShaderNodeVectorOpOperator) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOperator(op)
}
func (self Simple) GetOperator() classdb.VisualShaderNodeVectorOpOperator {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeVectorOpOperator(Expert(self).GetOperator())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeVectorOp
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetOperator(op classdb.VisualShaderNodeVectorOpOperator)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, op)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeVectorOp.Bind_set_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOperator() classdb.VisualShaderNodeVectorOpOperator {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeVectorOpOperator](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeVectorOp.Bind_get_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeVectorOp() Expert { return self[0].AsVisualShaderNodeVectorOp() }


//go:nosplit
func (self Simple) AsVisualShaderNodeVectorOp() Simple { return self[0].AsVisualShaderNodeVectorOp() }


//go:nosplit
func (self class) AsVisualShaderNodeVectorBase() VisualShaderNodeVectorBase.Expert { return self[0].AsVisualShaderNodeVectorBase() }


//go:nosplit
func (self Simple) AsVisualShaderNodeVectorBase() VisualShaderNodeVectorBase.Simple { return self[0].AsVisualShaderNodeVectorBase() }


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
func init() {classdb.Register("VisualShaderNodeVectorOp", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Operator = classdb.VisualShaderNodeVectorOpOperator

const (
/*Adds two vectors.*/
	OpAdd Operator = 0
/*Subtracts a vector from a vector.*/
	OpSub Operator = 1
/*Multiplies two vectors.*/
	OpMul Operator = 2
/*Divides vector by vector.*/
	OpDiv Operator = 3
/*Returns the remainder of the two vectors.*/
	OpMod Operator = 4
/*Returns the value of the first parameter raised to the power of the second, for each component of the vectors.*/
	OpPow Operator = 5
/*Returns the greater of two values, for each component of the vectors.*/
	OpMax Operator = 6
/*Returns the lesser of two values, for each component of the vectors.*/
	OpMin Operator = 7
/*Calculates the cross product of two vectors.*/
	OpCross Operator = 8
/*Returns the arc-tangent of the parameters.*/
	OpAtan2 Operator = 9
/*Returns the vector that points in the direction of reflection. [code]a[/code] is incident vector and [code]b[/code] is the normal vector.*/
	OpReflect Operator = 10
/*Vector step operator. Returns [code]0.0[/code] if [code]a[/code] is smaller than [code]b[/code] and [code]1.0[/code] otherwise.*/
	OpStep Operator = 11
/*Represents the size of the [enum Operator] enum.*/
	OpEnumSize Operator = 12
)
