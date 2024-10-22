package VisualShaderNodeTransformOp

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
Applies [member operator] to two transform (4×4 matrices) inputs.

*/
type Go [1]classdb.VisualShaderNodeTransformOp
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeTransformOp
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VisualShaderNodeTransformOp"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Operator() classdb.VisualShaderNodeTransformOpOperator {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.VisualShaderNodeTransformOpOperator(class(self).GetOperator())
}

func (self Go) SetOperator(value classdb.VisualShaderNodeTransformOpOperator) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOperator(value)
}

//go:nosplit
func (self class) SetOperator(op classdb.VisualShaderNodeTransformOpOperator)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, op)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeTransformOp.Bind_set_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOperator() classdb.VisualShaderNodeTransformOpOperator {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeTransformOpOperator](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeTransformOp.Bind_get_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeTransformOp() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeTransformOp() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("VisualShaderNodeTransformOp", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type Operator = classdb.VisualShaderNodeTransformOpOperator

const (
/*Multiplies transform [code]a[/code] by the transform [code]b[/code].*/
	OpAxb Operator = 0
/*Multiplies transform [code]b[/code] by the transform [code]a[/code].*/
	OpBxa Operator = 1
/*Performs a component-wise multiplication of transform [code]a[/code] by the transform [code]b[/code].*/
	OpAxbComp Operator = 2
/*Performs a component-wise multiplication of transform [code]b[/code] by the transform [code]a[/code].*/
	OpBxaComp Operator = 3
/*Adds two transforms.*/
	OpAdd Operator = 4
/*Subtracts the transform [code]a[/code] from the transform [code]b[/code].*/
	OpAMinusB Operator = 5
/*Subtracts the transform [code]b[/code] from the transform [code]a[/code].*/
	OpBMinusA Operator = 6
/*Divides the transform [code]a[/code] by the transform [code]b[/code].*/
	OpADivB Operator = 7
/*Divides the transform [code]b[/code] by the transform [code]a[/code].*/
	OpBDivA Operator = 8
/*Represents the size of the [enum Operator] enum.*/
	OpMax Operator = 9
)
