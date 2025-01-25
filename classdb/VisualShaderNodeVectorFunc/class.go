// Package VisualShaderNodeVectorFunc provides methods for working with VisualShaderNodeVectorFunc object instances.
package VisualShaderNodeVectorFunc

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/classdb/VisualShaderNodeVectorBase"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
A visual shader node able to perform different functions using vectors.
*/
type Instance [1]gdclass.VisualShaderNodeVectorFunc

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeVectorFunc() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeVectorFunc

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeVectorFunc"))
	casted := Instance{*(*gdclass.VisualShaderNodeVectorFunc)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Function() gdclass.VisualShaderNodeVectorFuncFunction {
	return gdclass.VisualShaderNodeVectorFuncFunction(class(self).GetFunction())
}

func (self Instance) SetFunction(value gdclass.VisualShaderNodeVectorFuncFunction) {
	class(self).SetFunction(value)
}

//go:nosplit
func (self class) SetFunction(fn gdclass.VisualShaderNodeVectorFuncFunction) {
	var frame = callframe.New()
	callframe.Arg(frame, fn)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeVectorFunc.Bind_set_function, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFunction() gdclass.VisualShaderNodeVectorFuncFunction {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeVectorFuncFunction](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeVectorFunc.Bind_get_function, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeVectorFunc() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeVectorFunc() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNodeVectorBase() VisualShaderNodeVectorBase.Advanced {
	return *((*VisualShaderNodeVectorBase.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeVectorBase() VisualShaderNodeVectorBase.Instance {
	return *((*VisualShaderNodeVectorBase.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(VisualShaderNodeVectorBase.Advanced(self.AsVisualShaderNodeVectorBase()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNodeVectorBase.Instance(self.AsVisualShaderNodeVectorBase()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeVectorFunc", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeVectorFunc{*(*gdclass.VisualShaderNodeVectorFunc)(unsafe.Pointer(&ptr))}
	})
}

type Function = gdclass.VisualShaderNodeVectorFuncFunction

const (
	/*Normalizes the vector so that it has a length of [code]1[/code] but points in the same direction.*/
	FuncNormalize Function = 0
	/*Clamps the value between [code]0.0[/code] and [code]1.0[/code].*/
	FuncSaturate Function = 1
	/*Returns the opposite value of the parameter.*/
	FuncNegate Function = 2
	/*Returns [code]1/vector[/code].*/
	FuncReciprocal Function = 3
	/*Returns the absolute value of the parameter.*/
	FuncAbs Function = 4
	/*Returns the arc-cosine of the parameter.*/
	FuncAcos Function = 5
	/*Returns the inverse hyperbolic cosine of the parameter.*/
	FuncAcosh Function = 6
	/*Returns the arc-sine of the parameter.*/
	FuncAsin Function = 7
	/*Returns the inverse hyperbolic sine of the parameter.*/
	FuncAsinh Function = 8
	/*Returns the arc-tangent of the parameter.*/
	FuncAtan Function = 9
	/*Returns the inverse hyperbolic tangent of the parameter.*/
	FuncAtanh Function = 10
	/*Finds the nearest integer that is greater than or equal to the parameter.*/
	FuncCeil Function = 11
	/*Returns the cosine of the parameter.*/
	FuncCos Function = 12
	/*Returns the hyperbolic cosine of the parameter.*/
	FuncCosh Function = 13
	/*Converts a quantity in radians to degrees.*/
	FuncDegrees Function = 14
	/*Base-e Exponential.*/
	FuncExp Function = 15
	/*Base-2 Exponential.*/
	FuncExp2 Function = 16
	/*Finds the nearest integer less than or equal to the parameter.*/
	FuncFloor Function = 17
	/*Computes the fractional part of the argument.*/
	FuncFract Function = 18
	/*Returns the inverse of the square root of the parameter.*/
	FuncInverseSqrt Function = 19
	/*Natural logarithm.*/
	FuncLog Function = 20
	/*Base-2 logarithm.*/
	FuncLog2 Function = 21
	/*Converts a quantity in degrees to radians.*/
	FuncRadians Function = 22
	/*Finds the nearest integer to the parameter.*/
	FuncRound Function = 23
	/*Finds the nearest even integer to the parameter.*/
	FuncRoundeven Function = 24
	/*Extracts the sign of the parameter, i.e. returns [code]-1[/code] if the parameter is negative, [code]1[/code] if it's positive and [code]0[/code] otherwise.*/
	FuncSign Function = 25
	/*Returns the sine of the parameter.*/
	FuncSin Function = 26
	/*Returns the hyperbolic sine of the parameter.*/
	FuncSinh Function = 27
	/*Returns the square root of the parameter.*/
	FuncSqrt Function = 28
	/*Returns the tangent of the parameter.*/
	FuncTan Function = 29
	/*Returns the hyperbolic tangent of the parameter.*/
	FuncTanh Function = 30
	/*Returns a value equal to the nearest integer to the parameter whose absolute value is not larger than the absolute value of the parameter.*/
	FuncTrunc Function = 31
	/*Returns [code]1.0 - vector[/code].*/
	FuncOneminus Function = 32
	/*Represents the size of the [enum Function] enum.*/
	FuncMax Function = 33
)
