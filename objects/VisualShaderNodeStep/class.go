package VisualShaderNodeStep

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
var _ = pointers.Cycle

/*
Translates to [code]step(edge, x)[/code] in the shader language.
Returns [code]0.0[/code] if [code]x[/code] is smaller than [code]edge[/code] and [code]1.0[/code] otherwise.
*/
type Instance [1]classdb.VisualShaderNodeStep
type Any interface {
	gd.IsClass
	AsVisualShaderNodeStep() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShaderNodeStep

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeStep"))
	return Instance{*(*classdb.VisualShaderNodeStep)(unsafe.Pointer(&object))}
}

func (self Instance) OpType() classdb.VisualShaderNodeStepOpType {
	return classdb.VisualShaderNodeStepOpType(class(self).GetOpType())
}

func (self Instance) SetOpType(value classdb.VisualShaderNodeStepOpType) {
	class(self).SetOpType(value)
}

//go:nosplit
func (self class) SetOpType(op_type classdb.VisualShaderNodeStepOpType) {
	var frame = callframe.New()
	callframe.Arg(frame, op_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeStep.Bind_set_op_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOpType() classdb.VisualShaderNodeStepOpType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeStepOpType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeStep.Bind_get_op_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeStep() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShaderNodeStep() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("VisualShaderNodeStep", func(ptr gd.Object) any {
		return [1]classdb.VisualShaderNodeStep{*(*classdb.VisualShaderNodeStep)(unsafe.Pointer(&ptr))}
	})
}

type OpType = classdb.VisualShaderNodeStepOpType

const (
	/*A floating-point scalar type.*/
	OpTypeScalar OpType = 0
	/*A 2D vector type.*/
	OpTypeVector2d OpType = 1
	/*The [code]x[/code] port uses a 2D vector type, while the [code]edge[/code] port uses a floating-point scalar type.*/
	OpTypeVector2dScalar OpType = 2
	/*A 3D vector type.*/
	OpTypeVector3d OpType = 3
	/*The [code]x[/code] port uses a 3D vector type, while the [code]edge[/code] port uses a floating-point scalar type.*/
	OpTypeVector3dScalar OpType = 4
	/*A 4D vector type.*/
	OpTypeVector4d OpType = 5
	/*The [code]a[/code] and [code]b[/code] ports use a 4D vector type. The [code]weight[/code] port uses a scalar type.*/
	OpTypeVector4dScalar OpType = 6
	/*Represents the size of the [enum OpType] enum.*/
	OpTypeMax OpType = 7
)
