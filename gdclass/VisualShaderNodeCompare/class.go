package VisualShaderNodeCompare

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
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
var _ = discreet.Root

/*
Compares [code]a[/code] and [code]b[/code] of [member type] by [member function]. Returns a boolean scalar. Translates to [code]if[/code] instruction in shader code.

*/
type Go [1]classdb.VisualShaderNodeCompare
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeCompare
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeCompare"))
	return Go{classdb.VisualShaderNodeCompare(object)}
}

func (self Go) Type() classdb.VisualShaderNodeCompareComparisonType {
		return classdb.VisualShaderNodeCompareComparisonType(class(self).GetComparisonType())
}

func (self Go) SetType(value classdb.VisualShaderNodeCompareComparisonType) {
	class(self).SetComparisonType(value)
}

func (self Go) Function() classdb.VisualShaderNodeCompareFunction {
		return classdb.VisualShaderNodeCompareFunction(class(self).GetFunction())
}

func (self Go) SetFunction(value classdb.VisualShaderNodeCompareFunction) {
	class(self).SetFunction(value)
}

func (self Go) Condition() classdb.VisualShaderNodeCompareCondition {
		return classdb.VisualShaderNodeCompareCondition(class(self).GetCondition())
}

func (self Go) SetCondition(value classdb.VisualShaderNodeCompareCondition) {
	class(self).SetCondition(value)
}

//go:nosplit
func (self class) SetComparisonType(atype classdb.VisualShaderNodeCompareComparisonType)  {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCompare.Bind_set_comparison_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetComparisonType() classdb.VisualShaderNodeCompareComparisonType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeCompareComparisonType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCompare.Bind_get_comparison_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFunction(fn classdb.VisualShaderNodeCompareFunction)  {
	var frame = callframe.New()
	callframe.Arg(frame, fn)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCompare.Bind_set_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFunction() classdb.VisualShaderNodeCompareFunction {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeCompareFunction](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCompare.Bind_get_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCondition(condition classdb.VisualShaderNodeCompareCondition)  {
	var frame = callframe.New()
	callframe.Arg(frame, condition)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCompare.Bind_set_condition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCondition() classdb.VisualShaderNodeCompareCondition {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeCompareCondition](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCompare.Bind_get_condition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeCompare() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeCompare() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("VisualShaderNodeCompare", func(ptr gd.Object) any { return classdb.VisualShaderNodeCompare(ptr) })}
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
