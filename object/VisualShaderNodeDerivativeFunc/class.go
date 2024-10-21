package VisualShaderNodeDerivativeFunc

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
This node is only available in [code]Fragment[/code] and [code]Light[/code] visual shaders.

*/
type Simple [1]classdb.VisualShaderNodeDerivativeFunc
func (self Simple) SetOpType(atype classdb.VisualShaderNodeDerivativeFuncOpType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOpType(atype)
}
func (self Simple) GetOpType() classdb.VisualShaderNodeDerivativeFuncOpType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeDerivativeFuncOpType(Expert(self).GetOpType())
}
func (self Simple) SetFunction(fn classdb.VisualShaderNodeDerivativeFuncFunction) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFunction(fn)
}
func (self Simple) GetFunction() classdb.VisualShaderNodeDerivativeFuncFunction {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeDerivativeFuncFunction(Expert(self).GetFunction())
}
func (self Simple) SetPrecision(precision classdb.VisualShaderNodeDerivativeFuncPrecision) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPrecision(precision)
}
func (self Simple) GetPrecision() classdb.VisualShaderNodeDerivativeFuncPrecision {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeDerivativeFuncPrecision(Expert(self).GetPrecision())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeDerivativeFunc
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetOpType(atype classdb.VisualShaderNodeDerivativeFuncOpType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeDerivativeFunc.Bind_set_op_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOpType() classdb.VisualShaderNodeDerivativeFuncOpType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeDerivativeFuncOpType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeDerivativeFunc.Bind_get_op_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFunction(fn classdb.VisualShaderNodeDerivativeFuncFunction)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fn)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeDerivativeFunc.Bind_set_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFunction() classdb.VisualShaderNodeDerivativeFuncFunction {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeDerivativeFuncFunction](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeDerivativeFunc.Bind_get_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPrecision(precision classdb.VisualShaderNodeDerivativeFuncPrecision)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, precision)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeDerivativeFunc.Bind_set_precision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPrecision() classdb.VisualShaderNodeDerivativeFuncPrecision {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeDerivativeFuncPrecision](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeDerivativeFunc.Bind_get_precision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeDerivativeFunc() Expert { return self[0].AsVisualShaderNodeDerivativeFunc() }


//go:nosplit
func (self Simple) AsVisualShaderNodeDerivativeFunc() Simple { return self[0].AsVisualShaderNodeDerivativeFunc() }


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
func init() {classdb.Register("VisualShaderNodeDerivativeFunc", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
