package VisualShaderNodeColorFunc

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
Accept a [Color] to the input port and transform it according to [member function].

*/
type Simple [1]classdb.VisualShaderNodeColorFunc
func (self Simple) SetFunction(fn classdb.VisualShaderNodeColorFuncFunction) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFunction(fn)
}
func (self Simple) GetFunction() classdb.VisualShaderNodeColorFuncFunction {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeColorFuncFunction(Expert(self).GetFunction())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeColorFunc
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetFunction(fn classdb.VisualShaderNodeColorFuncFunction)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fn)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeColorFunc.Bind_set_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFunction() classdb.VisualShaderNodeColorFuncFunction {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeColorFuncFunction](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeColorFunc.Bind_get_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeColorFunc() Expert { return self[0].AsVisualShaderNodeColorFunc() }


//go:nosplit
func (self Simple) AsVisualShaderNodeColorFunc() Simple { return self[0].AsVisualShaderNodeColorFunc() }


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
func init() {classdb.Register("VisualShaderNodeColorFunc", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Function = classdb.VisualShaderNodeColorFuncFunction

const (
/*Converts the color to grayscale using the following formula:
[codeblock]
vec3 c = input;
float max1 = max(c.r, c.g);
float max2 = max(max1, c.b);
float max3 = max(max1, max2);
return vec3(max3, max3, max3);
[/codeblock]*/
	FuncGrayscale Function = 0
/*Converts HSV vector to RGB equivalent.*/
	FuncHsv2rgb Function = 1
/*Converts RGB vector to HSV equivalent.*/
	FuncRgb2hsv Function = 2
/*Applies sepia tone effect using the following formula:
[codeblock]
vec3 c = input;
float r = (c.r * 0.393) + (c.g * 0.769) + (c.b * 0.189);
float g = (c.r * 0.349) + (c.g * 0.686) + (c.b * 0.168);
float b = (c.r * 0.272) + (c.g * 0.534) + (c.b * 0.131);
return vec3(r, g, b);
[/codeblock]*/
	FuncSepia Function = 3
/*Represents the size of the [enum Function] enum.*/
	FuncMax Function = 4
)
