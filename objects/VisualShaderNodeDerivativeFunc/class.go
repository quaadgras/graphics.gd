package VisualShaderNodeDerivativeFunc

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/VisualShaderNode"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
This node is only available in [code]Fragment[/code] and [code]Light[/code] visual shaders.
*/
type Instance [1]classdb.VisualShaderNodeDerivativeFunc

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShaderNodeDerivativeFunc

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeDerivativeFunc"))
	return Instance{classdb.VisualShaderNodeDerivativeFunc(object)}
}

func (self Instance) OpType() classdb.VisualShaderNodeDerivativeFuncOpType {
	return classdb.VisualShaderNodeDerivativeFuncOpType(class(self).GetOpType())
}

func (self Instance) SetOpType(value classdb.VisualShaderNodeDerivativeFuncOpType) {
	class(self).SetOpType(value)
}

func (self Instance) Function() classdb.VisualShaderNodeDerivativeFuncFunction {
	return classdb.VisualShaderNodeDerivativeFuncFunction(class(self).GetFunction())
}

func (self Instance) SetFunction(value classdb.VisualShaderNodeDerivativeFuncFunction) {
	class(self).SetFunction(value)
}

func (self Instance) Precision() classdb.VisualShaderNodeDerivativeFuncPrecision {
	return classdb.VisualShaderNodeDerivativeFuncPrecision(class(self).GetPrecision())
}

func (self Instance) SetPrecision(value classdb.VisualShaderNodeDerivativeFuncPrecision) {
	class(self).SetPrecision(value)
}

//go:nosplit
func (self class) SetOpType(atype classdb.VisualShaderNodeDerivativeFuncOpType) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeDerivativeFunc.Bind_set_op_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOpType() classdb.VisualShaderNodeDerivativeFuncOpType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeDerivativeFuncOpType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeDerivativeFunc.Bind_get_op_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFunction(fn classdb.VisualShaderNodeDerivativeFuncFunction) {
	var frame = callframe.New()
	callframe.Arg(frame, fn)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeDerivativeFunc.Bind_set_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFunction() classdb.VisualShaderNodeDerivativeFuncFunction {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeDerivativeFuncFunction](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeDerivativeFunc.Bind_get_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPrecision(precision classdb.VisualShaderNodeDerivativeFuncPrecision) {
	var frame = callframe.New()
	callframe.Arg(frame, precision)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeDerivativeFunc.Bind_set_precision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPrecision() classdb.VisualShaderNodeDerivativeFuncPrecision {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeDerivativeFuncPrecision](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeDerivativeFunc.Bind_get_precision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	classdb.Register("VisualShaderNodeDerivativeFunc", func(ptr gd.Object) any { return classdb.VisualShaderNodeDerivativeFunc(ptr) })
}

type OpType = classdb.VisualShaderNodeDerivativeFuncOpType

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

type Function = classdb.VisualShaderNodeDerivativeFuncFunction

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

type Precision = classdb.VisualShaderNodeDerivativeFuncPrecision

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
