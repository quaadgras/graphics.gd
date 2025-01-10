// Package VisualShaderNodeTransformFunc provides methods for working with VisualShaderNodeTransformFunc object instances.
package VisualShaderNodeTransformFunc

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
Computes an inverse or transpose function on the provided [Transform3D].
*/
type Instance [1]gdclass.VisualShaderNodeTransformFunc

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeTransformFunc() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeTransformFunc

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeTransformFunc"))
	casted := Instance{*(*gdclass.VisualShaderNodeTransformFunc)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Function() gdclass.VisualShaderNodeTransformFuncFunction {
	return gdclass.VisualShaderNodeTransformFuncFunction(class(self).GetFunction())
}

func (self Instance) SetFunction(value gdclass.VisualShaderNodeTransformFuncFunction) {
	class(self).SetFunction(value)
}

//go:nosplit
func (self class) SetFunction(fn gdclass.VisualShaderNodeTransformFuncFunction) {
	var frame = callframe.New()
	callframe.Arg(frame, fn)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTransformFunc.Bind_set_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFunction() gdclass.VisualShaderNodeTransformFuncFunction {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeTransformFuncFunction](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTransformFunc.Bind_get_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeTransformFunc() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeTransformFunc() Instance {
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
	gdclass.Register("VisualShaderNodeTransformFunc", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeTransformFunc{*(*gdclass.VisualShaderNodeTransformFunc)(unsafe.Pointer(&ptr))}
	})
}

type Function = gdclass.VisualShaderNodeTransformFuncFunction

const (
	/*Perform the inverse operation on the [Transform3D] matrix.*/
	FuncInverse Function = 0
	/*Perform the transpose operation on the [Transform3D] matrix.*/
	FuncTranspose Function = 1
	/*Represents the size of the [enum Function] enum.*/
	FuncMax Function = 2
)
