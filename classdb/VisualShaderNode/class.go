// Package VisualShaderNode provides methods for working with VisualShaderNode object instances.
package VisualShaderNode

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Visual shader graphs consist of various nodes. Each node in the graph is a separate object and they are represented as a rectangular boxes with title and a set of properties. Each node also has connection ports that allow to connect it to another nodes and control the flow of the shader.
*/
type Instance [1]gdclass.VisualShaderNode

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNode() Instance
}

/*
Returns the input port which should be connected by default when this node is created as a result of dragging a connection from an existing node to the empty space on the graph.
*/
func (self Instance) GetDefaultInputPort(atype gdclass.VisualShaderNodePortType) int {
	return int(int(class(self).GetDefaultInputPort(atype)))
}

/*
Sets the default [param value] for the selected input [param port].
*/
func (self Instance) SetInputPortDefaultValue(port int, value any) {
	class(self).SetInputPortDefaultValue(gd.Int(port), gd.NewVariant(value), gd.NewVariant(gd.NewVariant(([1]any{}[0]))))
}

/*
Returns the default value of the input [param port].
*/
func (self Instance) GetInputPortDefaultValue(port int) any {
	return any(class(self).GetInputPortDefaultValue(gd.Int(port)).Interface())
}

/*
Removes the default value of the input [param port].
*/
func (self Instance) RemoveInputPortDefaultValue(port int) {
	class(self).RemoveInputPortDefaultValue(gd.Int(port))
}

/*
Clears the default input ports value.
*/
func (self Instance) ClearDefaultInputValues() {
	class(self).ClearDefaultInputValues()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNode

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNode"))
	casted := Instance{*(*gdclass.VisualShaderNode)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) OutputPortForPreview() int {
	return int(int(class(self).GetOutputPortForPreview()))
}

func (self Instance) SetOutputPortForPreview(value int) {
	class(self).SetOutputPortForPreview(gd.Int(value))
}

func (self Instance) DefaultInputValues() Array.Any {
	return Array.Any(class(self).GetDefaultInputValues())
}

func (self Instance) SetDefaultInputValues(value Array.Any) {
	class(self).SetDefaultInputValues(value)
}

func (self Instance) LinkedParentGraphFrame() int {
	return int(int(class(self).GetFrame()))
}

func (self Instance) SetLinkedParentGraphFrame(value int) {
	class(self).SetFrame(gd.Int(value))
}

/*
Returns the input port which should be connected by default when this node is created as a result of dragging a connection from an existing node to the empty space on the graph.
*/
//go:nosplit
func (self class) GetDefaultInputPort(atype gdclass.VisualShaderNodePortType) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNode.Bind_get_default_input_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOutputPortForPreview(port gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, port)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNode.Bind_set_output_port_for_preview, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOutputPortForPreview() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNode.Bind_get_output_port_for_preview, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the default [param value] for the selected input [param port].
*/
//go:nosplit
func (self class) SetInputPortDefaultValue(port gd.Int, value gd.Variant, prev_value gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, pointers.Get(value))
	callframe.Arg(frame, pointers.Get(prev_value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNode.Bind_set_input_port_default_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the default value of the input [param port].
*/
//go:nosplit
func (self class) GetInputPortDefaultValue(port gd.Int) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, port)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNode.Bind_get_input_port_default_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Removes the default value of the input [param port].
*/
//go:nosplit
func (self class) RemoveInputPortDefaultValue(port gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, port)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNode.Bind_remove_input_port_default_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Clears the default input ports value.
*/
//go:nosplit
func (self class) ClearDefaultInputValues() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNode.Bind_clear_default_input_values, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the default input ports values using an [Array] of the form [code][index0, value0, index1, value1, ...][/code]. For example: [code][0, Vector3(0, 0, 0), 1, Vector3(0, 0, 0)][/code].
*/
//go:nosplit
func (self class) SetDefaultInputValues(values gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(values))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNode.Bind_set_default_input_values, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns an [Array] containing default values for all of the input ports of the node in the form [code][index0, value0, index1, value1, ...][/code].
*/
//go:nosplit
func (self class) GetDefaultInputValues() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNode.Bind_get_default_input_values, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrame(frame_ gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNode.Bind_set_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrame() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNode.Bind_get_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNode() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShaderNode() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNode", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNode{*(*gdclass.VisualShaderNode)(unsafe.Pointer(&ptr))}
	})
}

type PortType = gdclass.VisualShaderNodePortType

const (
	/*Floating-point scalar. Translated to [code skip-lint]float[/code] type in shader code.*/
	PortTypeScalar PortType = 0
	/*Integer scalar. Translated to [code skip-lint]int[/code] type in shader code.*/
	PortTypeScalarInt PortType = 1
	/*Unsigned integer scalar. Translated to [code skip-lint]uint[/code] type in shader code.*/
	PortTypeScalarUint PortType = 2
	/*2D vector of floating-point values. Translated to [code skip-lint]vec2[/code] type in shader code.*/
	PortTypeVector2d PortType = 3
	/*3D vector of floating-point values. Translated to [code skip-lint]vec3[/code] type in shader code.*/
	PortTypeVector3d PortType = 4
	/*4D vector of floating-point values. Translated to [code skip-lint]vec4[/code] type in shader code.*/
	PortTypeVector4d PortType = 5
	/*Boolean type. Translated to [code skip-lint]bool[/code] type in shader code.*/
	PortTypeBoolean PortType = 6
	/*Transform type. Translated to [code skip-lint]mat4[/code] type in shader code.*/
	PortTypeTransform PortType = 7
	/*Sampler type. Translated to reference of sampler uniform in shader code. Can only be used for input ports in non-uniform nodes.*/
	PortTypeSampler PortType = 8
	/*Represents the size of the [enum PortType] enum.*/
	PortTypeMax PortType = 9
)
