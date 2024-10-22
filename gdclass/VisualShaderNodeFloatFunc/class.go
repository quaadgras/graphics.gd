package VisualShaderNodeFloatFunc

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
Accept a floating-point scalar ([code]x[/code]) to the input port and transform it according to [member function].

*/
type Go [1]classdb.VisualShaderNodeFloatFunc
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeFloatFunc
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VisualShaderNodeFloatFunc"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Function() classdb.VisualShaderNodeFloatFuncFunction {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.VisualShaderNodeFloatFuncFunction(class(self).GetFunction())
}

func (self Go) SetFunction(value classdb.VisualShaderNodeFloatFuncFunction) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFunction(value)
}

//go:nosplit
func (self class) SetFunction(fn classdb.VisualShaderNodeFloatFuncFunction)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fn)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFloatFunc.Bind_set_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFunction() classdb.VisualShaderNodeFloatFuncFunction {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeFloatFuncFunction](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFloatFunc.Bind_get_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeFloatFunc() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeFloatFunc() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("VisualShaderNodeFloatFunc", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type Function = classdb.VisualShaderNodeFloatFuncFunction

const (
/*Returns the sine of the parameter. Translates to [code]sin(x)[/code] in the Godot Shader Language.*/
	FuncSin Function = 0
/*Returns the cosine of the parameter. Translates to [code]cos(x)[/code] in the Godot Shader Language.*/
	FuncCos Function = 1
/*Returns the tangent of the parameter. Translates to [code]tan(x)[/code] in the Godot Shader Language.*/
	FuncTan Function = 2
/*Returns the arc-sine of the parameter. Translates to [code]asin(x)[/code] in the Godot Shader Language.*/
	FuncAsin Function = 3
/*Returns the arc-cosine of the parameter. Translates to [code]acos(x)[/code] in the Godot Shader Language.*/
	FuncAcos Function = 4
/*Returns the arc-tangent of the parameter. Translates to [code]atan(x)[/code] in the Godot Shader Language.*/
	FuncAtan Function = 5
/*Returns the hyperbolic sine of the parameter. Translates to [code]sinh(x)[/code] in the Godot Shader Language.*/
	FuncSinh Function = 6
/*Returns the hyperbolic cosine of the parameter. Translates to [code]cosh(x)[/code] in the Godot Shader Language.*/
	FuncCosh Function = 7
/*Returns the hyperbolic tangent of the parameter. Translates to [code]tanh(x)[/code] in the Godot Shader Language.*/
	FuncTanh Function = 8
/*Returns the natural logarithm of the parameter. Translates to [code]log(x)[/code] in the Godot Shader Language.*/
	FuncLog Function = 9
/*Returns the natural exponentiation of the parameter. Translates to [code]exp(x)[/code] in the Godot Shader Language.*/
	FuncExp Function = 10
/*Returns the square root of the parameter. Translates to [code]sqrt(x)[/code] in the Godot Shader Language.*/
	FuncSqrt Function = 11
/*Returns the absolute value of the parameter. Translates to [code]abs(x)[/code] in the Godot Shader Language.*/
	FuncAbs Function = 12
/*Extracts the sign of the parameter. Translates to [code]sign(x)[/code] in the Godot Shader Language.*/
	FuncSign Function = 13
/*Finds the nearest integer less than or equal to the parameter. Translates to [code]floor(x)[/code] in the Godot Shader Language.*/
	FuncFloor Function = 14
/*Finds the nearest integer to the parameter. Translates to [code]round(x)[/code] in the Godot Shader Language.*/
	FuncRound Function = 15
/*Finds the nearest integer that is greater than or equal to the parameter. Translates to [code]ceil(x)[/code] in the Godot Shader Language.*/
	FuncCeil Function = 16
/*Computes the fractional part of the argument. Translates to [code]fract(x)[/code] in the Godot Shader Language.*/
	FuncFract Function = 17
/*Clamps the value between [code]0.0[/code] and [code]1.0[/code] using [code]min(max(x, 0.0), 1.0)[/code].*/
	FuncSaturate Function = 18
/*Negates the [code]x[/code] using [code]-(x)[/code].*/
	FuncNegate Function = 19
/*Returns the arc-hyperbolic-cosine of the parameter. Translates to [code]acosh(x)[/code] in the Godot Shader Language.*/
	FuncAcosh Function = 20
/*Returns the arc-hyperbolic-sine of the parameter. Translates to [code]asinh(x)[/code] in the Godot Shader Language.*/
	FuncAsinh Function = 21
/*Returns the arc-hyperbolic-tangent of the parameter. Translates to [code]atanh(x)[/code] in the Godot Shader Language.*/
	FuncAtanh Function = 22
/*Convert a quantity in radians to degrees. Translates to [code]degrees(x)[/code] in the Godot Shader Language.*/
	FuncDegrees Function = 23
/*Returns 2 raised by the power of the parameter. Translates to [code]exp2(x)[/code] in the Godot Shader Language.*/
	FuncExp2 Function = 24
/*Returns the inverse of the square root of the parameter. Translates to [code]inversesqrt(x)[/code] in the Godot Shader Language.*/
	FuncInverseSqrt Function = 25
/*Returns the base 2 logarithm of the parameter. Translates to [code]log2(x)[/code] in the Godot Shader Language.*/
	FuncLog2 Function = 26
/*Convert a quantity in degrees to radians. Translates to [code]radians(x)[/code] in the Godot Shader Language.*/
	FuncRadians Function = 27
/*Finds reciprocal value of dividing 1 by [code]x[/code] (i.e. [code]1 / x[/code]).*/
	FuncReciprocal Function = 28
/*Finds the nearest even integer to the parameter. Translates to [code]roundEven(x)[/code] in the Godot Shader Language.*/
	FuncRoundeven Function = 29
/*Returns a value equal to the nearest integer to [code]x[/code] whose absolute value is not larger than the absolute value of [code]x[/code]. Translates to [code]trunc(x)[/code] in the Godot Shader Language.*/
	FuncTrunc Function = 30
/*Subtracts scalar [code]x[/code] from 1 (i.e. [code]1 - x[/code]).*/
	FuncOneminus Function = 31
/*Represents the size of the [enum Function] enum.*/
	FuncMax Function = 32
)
