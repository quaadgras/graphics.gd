// Package VisualShaderNodeUVFunc provides methods for working with VisualShaderNodeUVFunc object instances.
package VisualShaderNodeUVFunc

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
UV functions are similar to [Vector2] functions, but the input port of this node uses the shader's UV value by default.
*/
type Instance [1]gdclass.VisualShaderNodeUVFunc

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeUVFunc() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeUVFunc

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeUVFunc"))
	casted := Instance{*(*gdclass.VisualShaderNodeUVFunc)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Function() gdclass.VisualShaderNodeUVFuncFunction {
	return gdclass.VisualShaderNodeUVFuncFunction(class(self).GetFunction())
}

func (self Instance) SetFunction(value gdclass.VisualShaderNodeUVFuncFunction) {
	class(self).SetFunction(value)
}

//go:nosplit
func (self class) SetFunction(fn gdclass.VisualShaderNodeUVFuncFunction) {
	var frame = callframe.New()
	callframe.Arg(frame, fn)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeUVFunc.Bind_set_function, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFunction() gdclass.VisualShaderNodeUVFuncFunction {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeUVFuncFunction](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeUVFunc.Bind_get_function, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeUVFunc() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShaderNodeUVFunc() Instance {
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
	gdclass.Register("VisualShaderNodeUVFunc", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeUVFunc{*(*gdclass.VisualShaderNodeUVFunc)(unsafe.Pointer(&ptr))}
	})
}

type Function = gdclass.VisualShaderNodeUVFuncFunction

const (
	/*Translates [code]uv[/code] by using [code]scale[/code] and [code]offset[/code] values using the following formula: [code]uv = uv + offset * scale[/code]. [code]uv[/code] port is connected to [code]UV[/code] built-in by default.*/
	FuncPanning Function = 0
	/*Scales [code]uv[/code] by using [code]scale[/code] and [code]pivot[/code] values using the following formula: [code]uv = (uv - pivot) * scale + pivot[/code]. [code]uv[/code] port is connected to [code]UV[/code] built-in by default.*/
	FuncScaling Function = 1
	/*Represents the size of the [enum Function] enum.*/
	FuncMax Function = 2
)
