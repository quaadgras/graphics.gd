package VisualShaderNodeVectorOp

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNodeVectorBase"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A visual shader node for use of vector operators. Operates on vector [code]a[/code] and vector [code]b[/code].

*/
type Go [1]classdb.VisualShaderNodeVectorOp
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeVectorOp
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VisualShaderNodeVectorOp"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Operator() classdb.VisualShaderNodeVectorOpOperator {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.VisualShaderNodeVectorOpOperator(class(self).GetOperator())
}

func (self Go) SetOperator(value classdb.VisualShaderNodeVectorOpOperator) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOperator(value)
}

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
func (self class) AsVisualShaderNodeVectorOp() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeVectorOp() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNodeVectorBase() VisualShaderNodeVectorBase.GD { return *((*VisualShaderNodeVectorBase.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeVectorBase() VisualShaderNodeVectorBase.Go { return *((*VisualShaderNodeVectorBase.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNode() VisualShaderNode.GD { return *((*VisualShaderNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNode() VisualShaderNode.Go { return *((*VisualShaderNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeVectorBase(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeVectorBase(), name)
	}
}
func init() {classdb.Register("VisualShaderNodeVectorOp", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
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
