package VisualShaderNodeTransformVecMult

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A multiplication operation on a transform (4×4 matrix) and a vector, with support for different multiplication operators.

*/
type Go [1]classdb.VisualShaderNodeTransformVecMult
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeTransformVecMult
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VisualShaderNodeTransformVecMult"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Operator() classdb.VisualShaderNodeTransformVecMultOperator {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.VisualShaderNodeTransformVecMultOperator(class(self).GetOperator())
}

func (self Go) SetOperator(value classdb.VisualShaderNodeTransformVecMultOperator) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOperator(value)
}

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
func (self class) AsVisualShaderNodeTransformVecMult() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeTransformVecMult() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNode() VisualShaderNode.GD { return *((*VisualShaderNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNode() VisualShaderNode.Go { return *((*VisualShaderNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}
func init() {classdb.Register("VisualShaderNodeTransformVecMult", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
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
