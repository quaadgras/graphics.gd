package VisualShaderNodeIntFunc

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/VisualShaderNode"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Accept an integer scalar ([code]x[/code]) to the input port and transform it according to [member function].
*/
type Instance [1]classdb.VisualShaderNodeIntFunc

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShaderNodeIntFunc

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeIntFunc"))
	return Instance{classdb.VisualShaderNodeIntFunc(object)}
}

func (self Instance) Function() classdb.VisualShaderNodeIntFuncFunction {
	return classdb.VisualShaderNodeIntFuncFunction(class(self).GetFunction())
}

func (self Instance) SetFunction(value classdb.VisualShaderNodeIntFuncFunction) {
	class(self).SetFunction(value)
}

//go:nosplit
func (self class) SetFunction(fn classdb.VisualShaderNodeIntFuncFunction) {
	var frame = callframe.New()
	callframe.Arg(frame, fn)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeIntFunc.Bind_set_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFunction() classdb.VisualShaderNodeIntFuncFunction {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeIntFuncFunction](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeIntFunc.Bind_get_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeIntFunc() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShaderNodeIntFunc() Instance {
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}
func init() {
	classdb.Register("VisualShaderNodeIntFunc", func(ptr gd.Object) any { return classdb.VisualShaderNodeIntFunc(ptr) })
}

type Function = classdb.VisualShaderNodeIntFuncFunction

const (
	/*Returns the absolute value of the parameter. Translates to [code]abs(x)[/code] in the Godot Shader Language.*/
	FuncAbs Function = 0
	/*Negates the [code]x[/code] using [code]-(x)[/code].*/
	FuncNegate Function = 1
	/*Extracts the sign of the parameter. Translates to [code]sign(x)[/code] in the Godot Shader Language.*/
	FuncSign Function = 2
	/*Returns the result of bitwise [code]NOT[/code] operation on the integer. Translates to [code]~a[/code] in the Godot Shader Language.*/
	FuncBitwiseNot Function = 3
	/*Represents the size of the [enum Function] enum.*/
	FuncMax Function = 4
)
