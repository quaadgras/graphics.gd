package VisualShaderNodeIs

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
Returns the boolean result of the comparison between [code]INF[/code] or [code]NaN[/code] and a scalar parameter.
*/
type Instance [1]classdb.VisualShaderNodeIs

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShaderNodeIs

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeIs"))
	return Instance{classdb.VisualShaderNodeIs(object)}
}

func (self Instance) Function() classdb.VisualShaderNodeIsFunction {
	return classdb.VisualShaderNodeIsFunction(class(self).GetFunction())
}

func (self Instance) SetFunction(value classdb.VisualShaderNodeIsFunction) {
	class(self).SetFunction(value)
}

//go:nosplit
func (self class) SetFunction(fn classdb.VisualShaderNodeIsFunction) {
	var frame = callframe.New()
	callframe.Arg(frame, fn)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeIs.Bind_set_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFunction() classdb.VisualShaderNodeIsFunction {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeIsFunction](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeIs.Bind_get_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeIs() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShaderNodeIs() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("VisualShaderNodeIs", func(ptr gd.Object) any { return classdb.VisualShaderNodeIs(ptr) })
}

type Function = classdb.VisualShaderNodeIsFunction

const (
	/*Comparison with [code]INF[/code] (Infinity).*/
	FuncIsInf Function = 0
	/*Comparison with [code]NaN[/code] (Not a Number; indicates invalid numeric results, such as division by zero).*/
	FuncIsNan Function = 1
	/*Represents the size of the [enum Function] enum.*/
	FuncMax Function = 2
)
