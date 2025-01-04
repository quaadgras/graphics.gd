// Package VisualShaderNodeColorFunc provides methods for working with VisualShaderNodeColorFunc object instances.
package VisualShaderNodeColorFunc

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Accept a [Color] to the input port and transform it according to [member function].
*/
type Instance [1]gdclass.VisualShaderNodeColorFunc
type Any interface {
	gd.IsClass
	AsVisualShaderNodeColorFunc() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeColorFunc

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeColorFunc"))
	return Instance{*(*gdclass.VisualShaderNodeColorFunc)(unsafe.Pointer(&object))}
}

func (self Instance) Function() gdclass.VisualShaderNodeColorFuncFunction {
	return gdclass.VisualShaderNodeColorFuncFunction(class(self).GetFunction())
}

func (self Instance) SetFunction(value gdclass.VisualShaderNodeColorFuncFunction) {
	class(self).SetFunction(value)
}

//go:nosplit
func (self class) SetFunction(fn gdclass.VisualShaderNodeColorFuncFunction) {
	var frame = callframe.New()
	callframe.Arg(frame, fn)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeColorFunc.Bind_set_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFunction() gdclass.VisualShaderNodeColorFuncFunction {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeColorFuncFunction](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeColorFunc.Bind_get_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeColorFunc() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeColorFunc() Instance {
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
	gdclass.Register("VisualShaderNodeColorFunc", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeColorFunc{*(*gdclass.VisualShaderNodeColorFunc)(unsafe.Pointer(&ptr))}
	})
}

type Function = gdclass.VisualShaderNodeColorFuncFunction

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
