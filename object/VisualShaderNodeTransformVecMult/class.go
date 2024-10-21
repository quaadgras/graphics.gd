package VisualShaderNodeTransformVecMult

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
A multiplication operation on a transform (4×4 matrix) and a vector, with support for different multiplication operators.

*/
type Simple [1]classdb.VisualShaderNodeTransformVecMult
func (self Simple) SetOperator(op classdb.VisualShaderNodeTransformVecMultOperator) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOperator(op)
}
func (self Simple) GetOperator() classdb.VisualShaderNodeTransformVecMultOperator {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeTransformVecMultOperator(Expert(self).GetOperator())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeTransformVecMult
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetOperator(op classdb.VisualShaderNodeTransformVecMultOperator)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, op)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeTransformVecMult.Bind_set_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOperator() classdb.VisualShaderNodeTransformVecMultOperator {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeTransformVecMultOperator](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeTransformVecMult.Bind_get_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeTransformVecMult() Expert { return self[0].AsVisualShaderNodeTransformVecMult() }


//go:nosplit
func (self Simple) AsVisualShaderNodeTransformVecMult() Simple { return self[0].AsVisualShaderNodeTransformVecMult() }


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
func init() {classdb.Register("VisualShaderNodeTransformVecMult", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
