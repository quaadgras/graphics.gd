package RDPipelineColorBlendState

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This object is used by [RenderingDevice].

*/
type Go [1]classdb.RDPipelineColorBlendState
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RDPipelineColorBlendState
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RDPipelineColorBlendState"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) EnableLogicOp() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetEnableLogicOp())
}

func (self Go) SetEnableLogicOp(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnableLogicOp(value)
}

func (self Go) LogicOp() classdb.RenderingDeviceLogicOperation {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceLogicOperation(class(self).GetLogicOp())
}

func (self Go) SetLogicOp(value classdb.RenderingDeviceLogicOperation) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLogicOp(value)
}

func (self Go) BlendConstant() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetBlendConstant())
}

func (self Go) SetBlendConstant(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBlendConstant(value)
}

func (self Go) Attachments() gd.ArrayOf[gdclass.RDPipelineColorBlendStateAttachment] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gdclass.RDPipelineColorBlendStateAttachment](class(self).GetAttachments(gc))
}

func (self Go) SetAttachments(value gd.ArrayOf[gdclass.RDPipelineColorBlendStateAttachment]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAttachments(value)
}

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
func (self class) SetAttachments(attachments gd.ArrayOf[gdclass.RDPipelineColorBlendStateAttachment])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(attachments))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendState.Bind_set_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAttachments(ctx gd.Lifetime) gd.ArrayOf[gdclass.RDPipelineColorBlendStateAttachment] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendState.Bind_get_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.RDPipelineColorBlendStateAttachment](ret)
}
func (self class) AsRDPipelineColorBlendState() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRDPipelineColorBlendState() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("RDPipelineColorBlendState", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
