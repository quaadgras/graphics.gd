package VisualShaderNodeCompare

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
Compares [code]a[/code] and [code]b[/code] of [member type] by [member function]. Returns a boolean scalar. Translates to [code]if[/code] instruction in shader code.

*/
type Simple [1]classdb.VisualShaderNodeCompare
func (self Simple) SetComparisonType(atype classdb.VisualShaderNodeCompareComparisonType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetComparisonType(atype)
}
func (self Simple) GetComparisonType() classdb.VisualShaderNodeCompareComparisonType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeCompareComparisonType(Expert(self).GetComparisonType())
}
func (self Simple) SetFunction(fn classdb.VisualShaderNodeCompareFunction) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFunction(fn)
}
func (self Simple) GetFunction() classdb.VisualShaderNodeCompareFunction {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeCompareFunction(Expert(self).GetFunction())
}
func (self Simple) SetCondition(condition classdb.VisualShaderNodeCompareCondition) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCondition(condition)
}
func (self Simple) GetCondition() classdb.VisualShaderNodeCompareCondition {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeCompareCondition(Expert(self).GetCondition())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeCompare
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetComparisonType(atype classdb.VisualShaderNodeCompareComparisonType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeCompare.Bind_set_comparison_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetComparisonType() classdb.VisualShaderNodeCompareComparisonType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeCompareComparisonType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeCompare.Bind_get_comparison_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFunction(fn classdb.VisualShaderNodeCompareFunction)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fn)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeCompare.Bind_set_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFunction() classdb.VisualShaderNodeCompareFunction {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeCompareFunction](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeCompare.Bind_get_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCondition(condition classdb.VisualShaderNodeCompareCondition)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, condition)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeCompare.Bind_set_condition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCondition() classdb.VisualShaderNodeCompareCondition {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeCompareCondition](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeCompare.Bind_get_condition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeCompare() Expert { return self[0].AsVisualShaderNodeCompare() }


//go:nosplit
func (self Simple) AsVisualShaderNodeCompare() Simple { return self[0].AsVisualShaderNodeCompare() }


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
func init() {classdb.Register("VisualShaderNodeCompare", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type ComparisonType = classdb.VisualShaderNodeCompareComparisonType

const (
/*A floating-point scalar.*/
	CtypeScalar ComparisonType = 0
/*An integer scalar.*/
	CtypeScalarInt ComparisonType = 1
/*An unsigned integer scalar.*/
	CtypeScalarUint ComparisonType = 2
/*A 2D vector type.*/
	CtypeVector2d ComparisonType = 3
/*A 3D vector type.*/
	CtypeVector3d ComparisonType = 4
/*A 4D vector type.*/
	CtypeVector4d ComparisonType = 5
/*A boolean type.*/
	CtypeBoolean ComparisonType = 6
/*A transform ([code]mat4[/code]) type.*/
	CtypeTransform ComparisonType = 7
/*Represents the size of the [enum ComparisonType] enum.*/
	CtypeMax ComparisonType = 8
)
type Function = classdb.VisualShaderNodeCompareFunction

const (
/*Comparison for equality ([code]a == b[/code]).*/
	FuncEqual Function = 0
/*Comparison for inequality ([code]a != b[/code]).*/
	FuncNotEqual Function = 1
/*Comparison for greater than ([code]a > b[/code]). Cannot be used if [member type] set to [constant CTYPE_BOOLEAN] or [constant CTYPE_TRANSFORM].*/
	FuncGreaterThan Function = 2
/*Comparison for greater than or equal ([code]a >= b[/code]). Cannot be used if [member type] set to [constant CTYPE_BOOLEAN] or [constant CTYPE_TRANSFORM].*/
	FuncGreaterThanEqual Function = 3
/*Comparison for less than ([code]a < b[/code]). Cannot be used if [member type] set to [constant CTYPE_BOOLEAN] or [constant CTYPE_TRANSFORM].*/
	FuncLessThan Function = 4
/*Comparison for less than or equal ([code]a <= b[/code]). Cannot be used if [member type] set to [constant CTYPE_BOOLEAN] or [constant CTYPE_TRANSFORM].*/
	FuncLessThanEqual Function = 5
/*Represents the size of the [enum Function] enum.*/
	FuncMax Function = 6
)
type Condition = classdb.VisualShaderNodeCompareCondition

const (
/*The result will be true if all of component in vector satisfy the comparison condition.*/
	CondAll Condition = 0
/*The result will be true if any of component in vector satisfy the comparison condition.*/
	CondAny Condition = 1
/*Represents the size of the [enum Condition] enum.*/
	CondMax Condition = 2
)
