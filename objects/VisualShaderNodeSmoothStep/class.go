package VisualShaderNodeSmoothStep

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/VisualShaderNode"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Translates to [code]smoothstep(edge0, edge1, x)[/code] in the shader language.
Returns [code]0.0[/code] if [code]x[/code] is smaller than [code]edge0[/code] and [code]1.0[/code] if [code]x[/code] is larger than [code]edge1[/code]. Otherwise, the return value is interpolated between [code]0.0[/code] and [code]1.0[/code] using Hermite polynomials.
*/
type Instance [1]classdb.VisualShaderNodeSmoothStep
type Any interface {
	gd.IsClass
	AsVisualShaderNodeSmoothStep() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShaderNodeSmoothStep

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeSmoothStep"))
	return Instance{classdb.VisualShaderNodeSmoothStep(object)}
}

func (self Instance) OpType() classdb.VisualShaderNodeSmoothStepOpType {
	return classdb.VisualShaderNodeSmoothStepOpType(class(self).GetOpType())
}

func (self Instance) SetOpType(value classdb.VisualShaderNodeSmoothStepOpType) {
	class(self).SetOpType(value)
}

//go:nosplit
func (self class) SetOpType(op_type classdb.VisualShaderNodeSmoothStepOpType) {
	var frame = callframe.New()
	callframe.Arg(frame, op_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeSmoothStep.Bind_set_op_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOpType() classdb.VisualShaderNodeSmoothStepOpType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeSmoothStepOpType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeSmoothStep.Bind_get_op_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeSmoothStep() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeSmoothStep() Instance {
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
	classdb.Register("VisualShaderNodeSmoothStep", func(ptr gd.Object) any {
		return [1]classdb.VisualShaderNodeSmoothStep{classdb.VisualShaderNodeSmoothStep(ptr)}
	})
}

type OpType = classdb.VisualShaderNodeSmoothStepOpType

const (
	/*A floating-point scalar type.*/
	OpTypeScalar OpType = 0
	/*A 2D vector type.*/
	OpTypeVector2d OpType = 1
	/*The [code]x[/code] port uses a 2D vector type. The first two ports use a floating-point scalar type.*/
	OpTypeVector2dScalar OpType = 2
	/*A 3D vector type.*/
	OpTypeVector3d OpType = 3
	/*The [code]x[/code] port uses a 3D vector type. The first two ports use a floating-point scalar type.*/
	OpTypeVector3dScalar OpType = 4
	/*A 4D vector type.*/
	OpTypeVector4d OpType = 5
	/*The [code]a[/code] and [code]b[/code] ports use a 4D vector type. The [code]weight[/code] port uses a scalar type.*/
	OpTypeVector4dScalar OpType = 6
	/*Represents the size of the [enum OpType] enum.*/
	OpTypeMax OpType = 7
)
