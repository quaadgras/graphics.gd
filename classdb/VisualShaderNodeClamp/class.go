// Package VisualShaderNodeClamp provides methods for working with VisualShaderNodeClamp object instances.
package VisualShaderNodeClamp

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

/*
Constrains a value to lie between [code]min[/code] and [code]max[/code] values.
*/
type Instance [1]gdclass.VisualShaderNodeClamp

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeClamp() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeClamp

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeClamp"))
	casted := Instance{*(*gdclass.VisualShaderNodeClamp)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) OpType() gdclass.VisualShaderNodeClampOpType {
	return gdclass.VisualShaderNodeClampOpType(class(self).GetOpType())
}

func (self Instance) SetOpType(value gdclass.VisualShaderNodeClampOpType) {
	class(self).SetOpType(value)
}

//go:nosplit
func (self class) SetOpType(op_type gdclass.VisualShaderNodeClampOpType) {
	var frame = callframe.New()
	callframe.Arg(frame, op_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeClamp.Bind_set_op_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOpType() gdclass.VisualShaderNodeClampOpType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeClampOpType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeClamp.Bind_get_op_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeClamp() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShaderNodeClamp() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("VisualShaderNodeClamp", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeClamp{*(*gdclass.VisualShaderNodeClamp)(unsafe.Pointer(&ptr))}
	})
}

type OpType = gdclass.VisualShaderNodeClampOpType

const (
	/*A floating-point scalar.*/
	OpTypeFloat OpType = 0
	/*An integer scalar.*/
	OpTypeInt OpType = 1
	/*An unsigned integer scalar.*/
	OpTypeUint OpType = 2
	/*A 2D vector type.*/
	OpTypeVector2d OpType = 3
	/*A 3D vector type.*/
	OpTypeVector3d OpType = 4
	/*A 4D vector type.*/
	OpTypeVector4d OpType = 5
	/*Represents the size of the [enum OpType] enum.*/
	OpTypeMax OpType = 6
)
