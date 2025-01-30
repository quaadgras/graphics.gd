// Package VisualShaderNodeCompare provides methods for working with VisualShaderNodeCompare object instances.
package VisualShaderNodeCompare

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
Compares [code]a[/code] and [code]b[/code] of [member type] by [member function]. Returns a boolean scalar. Translates to [code]if[/code] instruction in shader code.
*/
type Instance [1]gdclass.VisualShaderNodeCompare

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeCompare() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeCompare

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeCompare"))
	casted := Instance{*(*gdclass.VisualShaderNodeCompare)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Type() gdclass.VisualShaderNodeCompareComparisonType {
	return gdclass.VisualShaderNodeCompareComparisonType(class(self).GetComparisonType())
}

func (self Instance) SetType(value gdclass.VisualShaderNodeCompareComparisonType) {
	class(self).SetComparisonType(value)
}

func (self Instance) Function() gdclass.VisualShaderNodeCompareFunction {
	return gdclass.VisualShaderNodeCompareFunction(class(self).GetFunction())
}

func (self Instance) SetFunction(value gdclass.VisualShaderNodeCompareFunction) {
	class(self).SetFunction(value)
}

func (self Instance) Condition() gdclass.VisualShaderNodeCompareCondition {
	return gdclass.VisualShaderNodeCompareCondition(class(self).GetCondition())
}

func (self Instance) SetCondition(value gdclass.VisualShaderNodeCompareCondition) {
	class(self).SetCondition(value)
}

//go:nosplit
func (self class) SetComparisonType(atype gdclass.VisualShaderNodeCompareComparisonType) { //gd:VisualShaderNodeCompare.set_comparison_type
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCompare.Bind_set_comparison_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetComparisonType() gdclass.VisualShaderNodeCompareComparisonType { //gd:VisualShaderNodeCompare.get_comparison_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeCompareComparisonType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCompare.Bind_get_comparison_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFunction(fn gdclass.VisualShaderNodeCompareFunction) { //gd:VisualShaderNodeCompare.set_function
	var frame = callframe.New()
	callframe.Arg(frame, fn)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCompare.Bind_set_function, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFunction() gdclass.VisualShaderNodeCompareFunction { //gd:VisualShaderNodeCompare.get_function
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeCompareFunction](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCompare.Bind_get_function, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCondition(condition gdclass.VisualShaderNodeCompareCondition) { //gd:VisualShaderNodeCompare.set_condition
	var frame = callframe.New()
	callframe.Arg(frame, condition)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCompare.Bind_set_condition, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCondition() gdclass.VisualShaderNodeCompareCondition { //gd:VisualShaderNodeCompare.get_condition
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeCompareCondition](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCompare.Bind_get_condition, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeCompare() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShaderNodeCompare() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNode.Advanced(self.AsVisualShaderNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNode.Instance(self.AsVisualShaderNode()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeCompare", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeCompare{*(*gdclass.VisualShaderNodeCompare)(unsafe.Pointer(&ptr))}
	})
}

type ComparisonType = gdclass.VisualShaderNodeCompareComparisonType //gd:VisualShaderNodeCompare.ComparisonType

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

type Function = gdclass.VisualShaderNodeCompareFunction //gd:VisualShaderNodeCompare.Function

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

type Condition = gdclass.VisualShaderNodeCompareCondition //gd:VisualShaderNodeCompare.Condition

const (
	/*The result will be true if all of component in vector satisfy the comparison condition.*/
	CondAll Condition = 0
	/*The result will be true if any of component in vector satisfy the comparison condition.*/
	CondAny Condition = 1
	/*Represents the size of the [enum Condition] enum.*/
	CondMax Condition = 2
)
