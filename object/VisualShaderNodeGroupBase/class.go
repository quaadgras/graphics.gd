package VisualShaderNodeGroupBase

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VisualShaderNodeResizableBase"
import "grow.graphics/gd/object/VisualShaderNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Currently, has no direct usage, use the derived classes instead.

*/
type Simple [1]classdb.VisualShaderNodeGroupBase
func (self Simple) SetInputs(inputs string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInputs(gc.String(inputs))
}
func (self Simple) GetInputs() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetInputs(gc).String())
}
func (self Simple) SetOutputs(outputs string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOutputs(gc.String(outputs))
}
func (self Simple) GetOutputs() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetOutputs(gc).String())
}
func (self Simple) IsValidPortName(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsValidPortName(gc.String(name)))
}
func (self Simple) AddInputPort(id int, atype int, name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddInputPort(gd.Int(id), gd.Int(atype), gc.String(name))
}
func (self Simple) RemoveInputPort(id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveInputPort(gd.Int(id))
}
func (self Simple) GetInputPortCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetInputPortCount()))
}
func (self Simple) HasInputPort(id int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasInputPort(gd.Int(id)))
}
func (self Simple) ClearInputPorts() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearInputPorts()
}
func (self Simple) AddOutputPort(id int, atype int, name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddOutputPort(gd.Int(id), gd.Int(atype), gc.String(name))
}
func (self Simple) RemoveOutputPort(id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveOutputPort(gd.Int(id))
}
func (self Simple) GetOutputPortCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetOutputPortCount()))
}
func (self Simple) HasOutputPort(id int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasOutputPort(gd.Int(id)))
}
func (self Simple) ClearOutputPorts() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearOutputPorts()
}
func (self Simple) SetInputPortName(id int, name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInputPortName(gd.Int(id), gc.String(name))
}
func (self Simple) SetInputPortType(id int, atype int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInputPortType(gd.Int(id), gd.Int(atype))
}
func (self Simple) SetOutputPortName(id int, name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOutputPortName(gd.Int(id), gc.String(name))
}
func (self Simple) SetOutputPortType(id int, atype int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOutputPortType(gd.Int(id), gd.Int(atype))
}
func (self Simple) GetFreeInputPortId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetFreeInputPortId()))
}
func (self Simple) GetFreeOutputPortId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetFreeOutputPortId()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeGroupBase
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Defines all input ports using a [String] formatted as a colon-separated list: [code]id,type,name;[/code] (see [method add_input_port]).
*/
//go:nosplit
func (self class) SetInputs(inputs gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(inputs))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_set_inputs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [String] description of the input ports as a colon-separated list using the format [code]id,type,name;[/code] (see [method add_input_port]).
*/
//go:nosplit
func (self class) GetInputs(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_get_inputs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Defines all output ports using a [String] formatted as a colon-separated list: [code]id,type,name;[/code] (see [method add_output_port]).
*/
//go:nosplit
func (self class) SetOutputs(outputs gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(outputs))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_set_outputs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [String] description of the output ports as a colon-separated list using the format [code]id,type,name;[/code] (see [method add_output_port]).
*/
//go:nosplit
func (self class) GetOutputs(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_get_outputs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the specified port name does not override an existed port name and is valid within the shader.
*/
//go:nosplit
func (self class) IsValidPortName(name gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_is_valid_port_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds an input port with the specified [param type] (see [enum VisualShaderNode.PortType]) and [param name].
*/
//go:nosplit
func (self class) AddInputPort(id gd.Int, atype gd.Int, name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, atype)
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_add_input_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the specified input port.
*/
//go:nosplit
func (self class) RemoveInputPort(id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_remove_input_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of input ports in use. Alternative for [method get_free_input_port_id].
*/
//go:nosplit
func (self class) GetInputPortCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_get_input_port_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the specified input port exists.
*/
//go:nosplit
func (self class) HasInputPort(id gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_has_input_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes all previously specified input ports.
*/
//go:nosplit
func (self class) ClearInputPorts()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_clear_input_ports, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds an output port with the specified [param type] (see [enum VisualShaderNode.PortType]) and [param name].
*/
//go:nosplit
func (self class) AddOutputPort(id gd.Int, atype gd.Int, name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, atype)
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_add_output_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the specified output port.
*/
//go:nosplit
func (self class) RemoveOutputPort(id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_remove_output_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of output ports in use. Alternative for [method get_free_output_port_id].
*/
//go:nosplit
func (self class) GetOutputPortCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_get_output_port_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the specified output port exists.
*/
//go:nosplit
func (self class) HasOutputPort(id gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_has_output_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes all previously specified output ports.
*/
//go:nosplit
func (self class) ClearOutputPorts()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_clear_output_ports, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Renames the specified input port.
*/
//go:nosplit
func (self class) SetInputPortName(id gd.Int, name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_set_input_port_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the specified input port's type (see [enum VisualShaderNode.PortType]).
*/
//go:nosplit
func (self class) SetInputPortType(id gd.Int, atype gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_set_input_port_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Renames the specified output port.
*/
//go:nosplit
func (self class) SetOutputPortName(id gd.Int, name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_set_output_port_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the specified output port's type (see [enum VisualShaderNode.PortType]).
*/
//go:nosplit
func (self class) SetOutputPortType(id gd.Int, atype gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_set_output_port_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a free input port ID which can be used in [method add_input_port].
*/
//go:nosplit
func (self class) GetFreeInputPortId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_get_free_input_port_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a free output port ID which can be used in [method add_output_port].
*/
//go:nosplit
func (self class) GetFreeOutputPortId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeGroupBase.Bind_get_free_output_port_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeGroupBase() Expert { return self[0].AsVisualShaderNodeGroupBase() }


//go:nosplit
func (self Simple) AsVisualShaderNodeGroupBase() Simple { return self[0].AsVisualShaderNodeGroupBase() }


//go:nosplit
func (self class) AsVisualShaderNodeResizableBase() VisualShaderNodeResizableBase.Expert { return self[0].AsVisualShaderNodeResizableBase() }


//go:nosplit
func (self Simple) AsVisualShaderNodeResizableBase() VisualShaderNodeResizableBase.Simple { return self[0].AsVisualShaderNodeResizableBase() }


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
func init() {classdb.Register("VisualShaderNodeGroupBase", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
