package VisualShaderNode

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Visual shader graphs consist of various nodes. Each node in the graph is a separate object and they are represented as a rectangular boxes with title and a set of properties. Each node also has connection ports that allow to connect it to another nodes and control the flow of the shader.

*/
type Go [1]classdb.VisualShaderNode

/*
Returns the input port which should be connected by default when this node is created as a result of dragging a connection from an existing node to the empty space on the graph.
*/
func (self Go) GetDefaultInputPort(atype classdb.VisualShaderNodePortType) int {
	return int(int(class(self).GetDefaultInputPort(atype)))
}

/*
Sets the default [param value] for the selected input [param port].
*/
func (self Go) SetInputPortDefaultValue(port int, value gd.Variant) {
	class(self).SetInputPortDefaultValue(gd.Int(port), value, gd.NewVariant(([1]gd.Variant{}[0])))
}

/*
Returns the default value of the input [param port].
*/
func (self Go) GetInputPortDefaultValue(port int) gd.Variant {
	return gd.Variant(class(self).GetInputPortDefaultValue(gd.Int(port)))
}

/*
Removes the default value of the input [param port].
*/
func (self Go) RemoveInputPortDefaultValue(port int) {
	class(self).RemoveInputPortDefaultValue(gd.Int(port))
}

/*
Clears the default input ports value.
*/
func (self Go) ClearDefaultInputValues() {
	class(self).ClearDefaultInputValues()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNode
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNode"))
	return Go{classdb.VisualShaderNode(object)}
}

func (self Go) OutputPortForPreview() int {
		return int(int(class(self).GetOutputPortForPreview()))
}

func (self Go) SetOutputPortForPreview(value int) {
	class(self).SetOutputPortForPreview(gd.Int(value))
}

func (self Go) DefaultInputValues() gd.Array {
		return gd.Array(class(self).GetDefaultInputValues())
}

func (self Go) SetDefaultInputValues(value gd.Array) {
	class(self).SetDefaultInputValues(value)
}

func (self Go) LinkedParentGraphFrame() int {
		return int(int(class(self).GetFrame()))
}

func (self Go) SetLinkedParentGraphFrame(value int) {
	class(self).SetFrame(gd.Int(value))
}

/*
Returns the input port which should be connected by default when this node is created as a result of dragging a connection from an existing node to the empty space on the graph.
*/
//go:nosplit
func (self class) GetDefaultInputPort(atype classdb.VisualShaderNodePortType) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNode.Bind_get_default_input_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOutputPortForPreview(port gd.Int)  {
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
func (self class) SetInputPortDefaultValue(port gd.Int, value gd.Variant, prev_value gd.Variant)  {
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, discreet.Get(value))
	callframe.Arg(frame, discreet.Get(prev_value))
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
	var ret = discreet.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
/*
Removes the default value of the input [param port].
*/
//go:nosplit
func (self class) RemoveInputPortDefaultValue(port gd.Int)  {
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
func (self class) ClearDefaultInputValues()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNode.Bind_clear_default_input_values, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the default input ports values using an [Array] of the form [code][index0, value0, index1, value1, ...][/code]. For example: [code][0, Vector3(0, 0, 0), 1, Vector3(0, 0, 0)][/code].
*/
//go:nosplit
func (self class) SetDefaultInputValues(values gd.Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(values))
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
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFrame(frame_ gd.Int)  {
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
func (self class) AsVisualShaderNode() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNode() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("VisualShaderNode", func(ptr gd.Object) any { return classdb.VisualShaderNode(ptr) })}
type PortType = classdb.VisualShaderNodePortType

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
