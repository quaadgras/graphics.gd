package VisualShaderNodeSwitch

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
Returns an associated value of the [member op_type] type if the provided boolean value is [code]true[/code] or [code]false[/code].

*/
type Simple [1]classdb.VisualShaderNodeSwitch
func (self Simple) SetOpType(atype classdb.VisualShaderNodeSwitchOpType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOpType(atype)
}
func (self Simple) GetOpType() classdb.VisualShaderNodeSwitchOpType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeSwitchOpType(Expert(self).GetOpType())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeSwitch
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetOpType(atype classdb.VisualShaderNodeSwitchOpType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeSwitch.Bind_set_op_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOpType() classdb.VisualShaderNodeSwitchOpType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeSwitchOpType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeSwitch.Bind_get_op_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeSwitch() Expert { return self[0].AsVisualShaderNodeSwitch() }


//go:nosplit
func (self Simple) AsVisualShaderNodeSwitch() Simple { return self[0].AsVisualShaderNodeSwitch() }


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
func init() {classdb.Register("VisualShaderNodeSwitch", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type OpType = classdb.VisualShaderNodeSwitchOpType

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
/*A boolean type.*/
	OpTypeBoolean OpType = 6
/*A transform type.*/
	OpTypeTransform OpType = 7
/*Represents the size of the [enum OpType] enum.*/
	OpTypeMax OpType = 8
)
