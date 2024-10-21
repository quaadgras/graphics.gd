package RDPipelineColorBlendState

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This object is used by [RenderingDevice].

*/
type Simple [1]classdb.RDPipelineColorBlendState
func (self Simple) SetEnableLogicOp(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnableLogicOp(p_member)
}
func (self Simple) GetEnableLogicOp() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetEnableLogicOp())
}
func (self Simple) SetLogicOp(p_member classdb.RenderingDeviceLogicOperation) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLogicOp(p_member)
}
func (self Simple) GetLogicOp() classdb.RenderingDeviceLogicOperation {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceLogicOperation(Expert(self).GetLogicOp())
}
func (self Simple) SetBlendConstant(p_member gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBlendConstant(p_member)
}
func (self Simple) GetBlendConstant() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetBlendConstant())
}
func (self Simple) SetAttachments(attachments gd.ArrayOf[[1]classdb.RDPipelineColorBlendStateAttachment]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAttachments(attachments)
}
func (self Simple) GetAttachments() gd.ArrayOf[[1]classdb.RDPipelineColorBlendStateAttachment] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.RDPipelineColorBlendStateAttachment](Expert(self).GetAttachments(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RDPipelineColorBlendState
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetEnableLogicOp(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendState.Bind_set_enable_logic_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnableLogicOp() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendState.Bind_get_enable_logic_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLogicOp(p_member classdb.RenderingDeviceLogicOperation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendState.Bind_set_logic_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLogicOp() classdb.RenderingDeviceLogicOperation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceLogicOperation](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendState.Bind_get_logic_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBlendConstant(p_member gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendState.Bind_set_blend_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBlendConstant() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendState.Bind_get_blend_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAttachments(attachments gd.ArrayOf[object.RDPipelineColorBlendStateAttachment])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(attachments))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendState.Bind_set_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAttachments(ctx gd.Lifetime) gd.ArrayOf[object.RDPipelineColorBlendStateAttachment] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendState.Bind_get_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.RDPipelineColorBlendStateAttachment](ret)
}

//go:nosplit
func (self class) AsRDPipelineColorBlendState() Expert { return self[0].AsRDPipelineColorBlendState() }


//go:nosplit
func (self Simple) AsRDPipelineColorBlendState() Simple { return self[0].AsRDPipelineColorBlendState() }


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
func init() {classdb.Register("RDPipelineColorBlendState", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
