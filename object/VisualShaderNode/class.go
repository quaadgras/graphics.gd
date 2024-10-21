package VisualShaderNode

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Visual shader graphs consist of various nodes. Each node in the graph is a separate object and they are represented as a rectangular boxes with title and a set of properties. Each node also has connection ports that allow to connect it to another nodes and control the flow of the shader.

*/
type Simple [1]classdb.VisualShaderNode
func (self Simple) GetDefaultInputPort(atype classdb.VisualShaderNodePortType) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetDefaultInputPort(atype)))
}
func (self Simple) SetOutputPortForPreview(port int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOutputPortForPreview(gd.Int(port))
}
func (self Simple) GetOutputPortForPreview() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetOutputPortForPreview()))
}
func (self Simple) SetInputPortDefaultValue(port int, value gd.Variant, prev_value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInputPortDefaultValue(gd.Int(port), value, prev_value)
}
func (self Simple) GetInputPortDefaultValue(port int) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetInputPortDefaultValue(gc, gd.Int(port)))
}
func (self Simple) RemoveInputPortDefaultValue(port int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveInputPortDefaultValue(gd.Int(port))
}
func (self Simple) ClearDefaultInputValues() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearDefaultInputValues()
}
func (self Simple) SetDefaultInputValues(values gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDefaultInputValues(values)
}
func (self Simple) GetDefaultInputValues() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).GetDefaultInputValues(gc))
}
func (self Simple) SetFrame(frame_ int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFrame(gd.Int(frame_))
}
func (self Simple) GetFrame() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetFrame()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNode
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the input port which should be connected by default when this node is created as a result of dragging a connection from an existing node to the empty space on the graph.
*/
//go:nosplit
func (self class) GetDefaultInputPort(atype classdb.VisualShaderNodePortType) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNode.Bind_get_default_input_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOutputPortForPreview(port gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNode.Bind_set_output_port_for_preview, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOutputPortForPreview() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNode.Bind_get_output_port_for_preview, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the default [param value] for the selected input [param port].
*/
//go:nosplit
func (self class) SetInputPortDefaultValue(port gd.Int, value gd.Variant, prev_value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, mmm.Get(value))
	callframe.Arg(frame, mmm.Get(prev_value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNode.Bind_set_input_port_default_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the default value of the input [param port].
*/
//go:nosplit
func (self class) GetInputPortDefaultValue(ctx gd.Lifetime, port gd.Int) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNode.Bind_get_input_port_default_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Removes the default value of the input [param port].
*/
//go:nosplit
func (self class) RemoveInputPortDefaultValue(port gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNode.Bind_remove_input_port_default_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears the default input ports value.
*/
//go:nosplit
func (self class) ClearDefaultInputValues()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNode.Bind_clear_default_input_values, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the default input ports values using an [Array] of the form [code][index0, value0, index1, value1, ...][/code]. For example: [code][0, Vector3(0, 0, 0), 1, Vector3(0, 0, 0)][/code].
*/
//go:nosplit
func (self class) SetDefaultInputValues(values gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(values))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNode.Bind_set_default_input_values, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an [Array] containing default values for all of the input ports of the node in the form [code][index0, value0, index1, value1, ...][/code].
*/
//go:nosplit
func (self class) GetDefaultInputValues(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNode.Bind_get_default_input_values, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFrame(frame_ gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNode.Bind_set_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFrame() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNode.Bind_get_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNode() Expert { return self[0].AsVisualShaderNode() }


//go:nosplit
func (self Simple) AsVisualShaderNode() Simple { return self[0].AsVisualShaderNode() }


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
func init() {classdb.Register("VisualShaderNode", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
