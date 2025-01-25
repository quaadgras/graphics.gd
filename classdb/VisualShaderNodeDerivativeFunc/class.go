// Package VisualShaderNodeDerivativeFunc provides methods for working with VisualShaderNodeDerivativeFunc object instances.
package VisualShaderNodeDerivativeFunc

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
This node is only available in [code]Fragment[/code] and [code]Light[/code] visual shaders.
*/
type Instance [1]gdclass.VisualShaderNodeDerivativeFunc

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeDerivativeFunc() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeDerivativeFunc

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeDerivativeFunc"))
	casted := Instance{*(*gdclass.VisualShaderNodeDerivativeFunc)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) OpType() gdclass.VisualShaderNodeDerivativeFuncOpType {
	return gdclass.VisualShaderNodeDerivativeFuncOpType(class(self).GetOpType())
}

func (self Instance) SetOpType(value gdclass.VisualShaderNodeDerivativeFuncOpType) {
	class(self).SetOpType(value)
}

func (self Instance) Function() gdclass.VisualShaderNodeDerivativeFuncFunction {
	return gdclass.VisualShaderNodeDerivativeFuncFunction(class(self).GetFunction())
}

func (self Instance) SetFunction(value gdclass.VisualShaderNodeDerivativeFuncFunction) {
	class(self).SetFunction(value)
}

func (self Instance) Precision() gdclass.VisualShaderNodeDerivativeFuncPrecision {
	return gdclass.VisualShaderNodeDerivativeFuncPrecision(class(self).GetPrecision())
}

func (self Instance) SetPrecision(value gdclass.VisualShaderNodeDerivativeFuncPrecision) {
	class(self).SetPrecision(value)
}

//go:nosplit
func (self class) SetOpType(atype gdclass.VisualShaderNodeDerivativeFuncOpType) { //gd:VisualShaderNodeDerivativeFunc.set_op_type
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeDerivativeFunc.Bind_set_op_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOpType() gdclass.VisualShaderNodeDerivativeFuncOpType { //gd:VisualShaderNodeDerivativeFunc.get_op_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeDerivativeFuncOpType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeDerivativeFunc.Bind_get_op_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFunction(fn gdclass.VisualShaderNodeDerivativeFuncFunction) { //gd:VisualShaderNodeDerivativeFunc.set_function
	var frame = callframe.New()
	callframe.Arg(frame, fn)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeDerivativeFunc.Bind_set_function, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFunction() gdclass.VisualShaderNodeDerivativeFuncFunction { //gd:VisualShaderNodeDerivativeFunc.get_function
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeDerivativeFuncFunction](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeDerivativeFunc.Bind_get_function, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPrecision(precision gdclass.VisualShaderNodeDerivativeFuncPrecision) { //gd:VisualShaderNodeDerivativeFunc.set_precision
	var frame = callframe.New()
	callframe.Arg(frame, precision)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeDerivativeFunc.Bind_set_precision, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPrecision() gdclass.VisualShaderNodeDerivativeFuncPrecision { //gd:VisualShaderNodeDerivativeFunc.get_precision
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeDerivativeFuncPrecision](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeDerivativeFunc.Bind_get_precision, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeDerivativeFunc() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeDerivativeFunc() Instance {
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
	gdclass.Register("VisualShaderNodeDerivativeFunc", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeDerivativeFunc{*(*gdclass.VisualShaderNodeDerivativeFunc)(unsafe.Pointer(&ptr))}
	})
}

type OpType = gdclass.VisualShaderNodeDerivativeFuncOpType //gd:VisualShaderNodeDerivativeFunc.OpType

const (
	/*A floating-point scalar.*/
	OpTypeScalar OpType = 0
	/*A 2D vector type.*/
	OpTypeVector2d OpType = 1
	/*A 3D vector type.*/
	OpTypeVector3d OpType = 2
	/*A 4D vector type.*/
	OpTypeVector4d OpType = 3
	/*Represents the size of the [enum OpType] enum.*/
	OpTypeMax OpType = 4
)

type Function = gdclass.VisualShaderNodeDerivativeFuncFunction //gd:VisualShaderNodeDerivativeFunc.Function

const (
	/*Sum of absolute derivative in [code]x[/code] and [code]y[/code].*/
	FuncSum Function = 0
	/*Derivative in [code]x[/code] using local differencing.*/
	FuncX Function = 1
	/*Derivative in [code]y[/code] using local differencing.*/
	FuncY Function = 2
	/*Represents the size of the [enum Function] enum.*/
	FuncMax Function = 3
)

type Precision = gdclass.VisualShaderNodeDerivativeFuncPrecision //gd:VisualShaderNodeDerivativeFunc.Precision

const (
	/*No precision is specified, the GPU driver is allowed to use whatever level of precision it chooses. This is the default option and is equivalent to using [code]dFdx()[/code] or [code]dFdy()[/code] in text shaders.*/
	PrecisionNone Precision = 0
	/*The derivative will be calculated using the current fragment's neighbors (which may not include the current fragment). This tends to be faster than using [constant PRECISION_FINE], but may not be suitable when more precision is needed. This is equivalent to using [code]dFdxCoarse()[/code] or [code]dFdyCoarse()[/code] in text shaders.*/
	PrecisionCoarse Precision = 1
	/*The derivative will be calculated using the current fragment and its immediate neighbors. This tends to be slower than using [constant PRECISION_COARSE], but may be necessary when more precision is needed. This is equivalent to using [code]dFdxFine()[/code] or [code]dFdyFine()[/code] in text shaders.*/
	PrecisionFine Precision = 2
	/*Represents the size of the [enum Precision] enum.*/
	PrecisionMax Precision = 3
)
