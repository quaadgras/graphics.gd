package RDPipelineColorBlendStateAttachment

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
Controls how blending between source and destination fragments is performed when using [RenderingDevice].
For reference, this is how common user-facing blend modes are implemented in Godot's 2D renderer:
[b]Mix:[/b]
[codeblock]
var attachment = RDPipelineColorBlendStateAttachment.new()
attachment.enable_blend = true
attachment.color_blend_op = RenderingDevice.BLEND_OP_ADD
attachment.src_color_blend_factor = RenderingDevice.BLEND_FACTOR_SRC_ALPHA
attachment.dst_color_blend_factor = RenderingDevice.BLEND_FACTOR_ONE_MINUS_SRC_ALPHA
attachment.alpha_blend_op = RenderingDevice.BLEND_OP_ADD
attachment.src_alpha_blend_factor = RenderingDevice.BLEND_FACTOR_ONE
attachment.dst_alpha_blend_factor = RenderingDevice.BLEND_FACTOR_ONE_MINUS_SRC_ALPHA
[/codeblock]
[b]Add:[/b]
[codeblock]
var attachment = RDPipelineColorBlendStateAttachment.new()
attachment.enable_blend = true
attachment.alpha_blend_op = RenderingDevice.BLEND_OP_ADD
attachment.color_blend_op = RenderingDevice.BLEND_OP_ADD
attachment.src_color_blend_factor = RenderingDevice.BLEND_FACTOR_SRC_ALPHA
attachment.dst_color_blend_factor = RenderingDevice.BLEND_FACTOR_ONE
attachment.src_alpha_blend_factor = RenderingDevice.BLEND_FACTOR_SRC_ALPHA
attachment.dst_alpha_blend_factor = RenderingDevice.BLEND_FACTOR_ONE
[/codeblock]
[b]Subtract:[/b]
[codeblock]
var attachment = RDPipelineColorBlendStateAttachment.new()
attachment.enable_blend = true
attachment.alpha_blend_op = RenderingDevice.BLEND_OP_REVERSE_SUBTRACT
attachment.color_blend_op = RenderingDevice.BLEND_OP_REVERSE_SUBTRACT
attachment.src_color_blend_factor = RenderingDevice.BLEND_FACTOR_SRC_ALPHA
attachment.dst_color_blend_factor = RenderingDevice.BLEND_FACTOR_ONE
attachment.src_alpha_blend_factor = RenderingDevice.BLEND_FACTOR_SRC_ALPHA
attachment.dst_alpha_blend_factor = RenderingDevice.BLEND_FACTOR_ONE
[/codeblock]
[b]Multiply:[/b]
[codeblock]
var attachment = RDPipelineColorBlendStateAttachment.new()
attachment.enable_blend = true
attachment.alpha_blend_op = RenderingDevice.BLEND_OP_ADD
attachment.color_blend_op = RenderingDevice.BLEND_OP_ADD
attachment.src_color_blend_factor = RenderingDevice.BLEND_FACTOR_DST_COLOR
attachment.dst_color_blend_factor = RenderingDevice.BLEND_FACTOR_ZERO
attachment.src_alpha_blend_factor = RenderingDevice.BLEND_FACTOR_DST_ALPHA
attachment.dst_alpha_blend_factor = RenderingDevice.BLEND_FACTOR_ZERO
[/codeblock]
[b]Pre-multiplied alpha:[/b]
[codeblock]
var attachment = RDPipelineColorBlendStateAttachment.new()
attachment.enable_blend = true
attachment.alpha_blend_op = RenderingDevice.BLEND_OP_ADD
attachment.color_blend_op = RenderingDevice.BLEND_OP_ADD
attachment.src_color_blend_factor = RenderingDevice.BLEND_FACTOR_ONE
attachment.dst_color_blend_factor = RenderingDevice.BLEND_FACTOR_ONE_MINUS_SRC_ALPHA
attachment.src_alpha_blend_factor = RenderingDevice.BLEND_FACTOR_ONE
attachment.dst_alpha_blend_factor = RenderingDevice.BLEND_FACTOR_ONE_MINUS_SRC_ALPHA
[/codeblock]

*/
type Simple [1]classdb.RDPipelineColorBlendStateAttachment
func (self Simple) SetAsMix() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAsMix()
}
func (self Simple) SetEnableBlend(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnableBlend(p_member)
}
func (self Simple) GetEnableBlend() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetEnableBlend())
}
func (self Simple) SetSrcColorBlendFactor(p_member classdb.RenderingDeviceBlendFactor) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSrcColorBlendFactor(p_member)
}
func (self Simple) GetSrcColorBlendFactor() classdb.RenderingDeviceBlendFactor {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceBlendFactor(Expert(self).GetSrcColorBlendFactor())
}
func (self Simple) SetDstColorBlendFactor(p_member classdb.RenderingDeviceBlendFactor) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDstColorBlendFactor(p_member)
}
func (self Simple) GetDstColorBlendFactor() classdb.RenderingDeviceBlendFactor {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceBlendFactor(Expert(self).GetDstColorBlendFactor())
}
func (self Simple) SetColorBlendOp(p_member classdb.RenderingDeviceBlendOperation) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColorBlendOp(p_member)
}
func (self Simple) GetColorBlendOp() classdb.RenderingDeviceBlendOperation {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceBlendOperation(Expert(self).GetColorBlendOp())
}
func (self Simple) SetSrcAlphaBlendFactor(p_member classdb.RenderingDeviceBlendFactor) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSrcAlphaBlendFactor(p_member)
}
func (self Simple) GetSrcAlphaBlendFactor() classdb.RenderingDeviceBlendFactor {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceBlendFactor(Expert(self).GetSrcAlphaBlendFactor())
}
func (self Simple) SetDstAlphaBlendFactor(p_member classdb.RenderingDeviceBlendFactor) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDstAlphaBlendFactor(p_member)
}
func (self Simple) GetDstAlphaBlendFactor() classdb.RenderingDeviceBlendFactor {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceBlendFactor(Expert(self).GetDstAlphaBlendFactor())
}
func (self Simple) SetAlphaBlendOp(p_member classdb.RenderingDeviceBlendOperation) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAlphaBlendOp(p_member)
}
func (self Simple) GetAlphaBlendOp() classdb.RenderingDeviceBlendOperation {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceBlendOperation(Expert(self).GetAlphaBlendOp())
}
func (self Simple) SetWriteR(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetWriteR(p_member)
}
func (self Simple) GetWriteR() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetWriteR())
}
func (self Simple) SetWriteG(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetWriteG(p_member)
}
func (self Simple) GetWriteG() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetWriteG())
}
func (self Simple) SetWriteB(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetWriteB(p_member)
}
func (self Simple) GetWriteB() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetWriteB())
}
func (self Simple) SetWriteA(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetWriteA(p_member)
}
func (self Simple) GetWriteA() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetWriteA())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RDPipelineColorBlendStateAttachment
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Convenience method to perform standard mix blending with straight (non-premultiplied) alpha. This sets [member enable_blend] to [code]true[/code], [member src_color_blend_factor] to [constant RenderingDevice.BLEND_FACTOR_SRC_ALPHA], [member dst_color_blend_factor] to [constant RenderingDevice.BLEND_FACTOR_ONE_MINUS_SRC_ALPHA], [member src_alpha_blend_factor] to [constant RenderingDevice.BLEND_FACTOR_SRC_ALPHA] and [member dst_alpha_blend_factor] to [constant RenderingDevice.BLEND_FACTOR_ONE_MINUS_SRC_ALPHA].
*/
//go:nosplit
func (self class) SetAsMix()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_set_as_mix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetEnableBlend(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_set_enable_blend, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnableBlend() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_get_enable_blend, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSrcColorBlendFactor(p_member classdb.RenderingDeviceBlendFactor)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_set_src_color_blend_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSrcColorBlendFactor() classdb.RenderingDeviceBlendFactor {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceBlendFactor](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_get_src_color_blend_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDstColorBlendFactor(p_member classdb.RenderingDeviceBlendFactor)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_set_dst_color_blend_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDstColorBlendFactor() classdb.RenderingDeviceBlendFactor {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceBlendFactor](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_get_dst_color_blend_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColorBlendOp(p_member classdb.RenderingDeviceBlendOperation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_set_color_blend_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColorBlendOp() classdb.RenderingDeviceBlendOperation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceBlendOperation](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_get_color_blend_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSrcAlphaBlendFactor(p_member classdb.RenderingDeviceBlendFactor)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_set_src_alpha_blend_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSrcAlphaBlendFactor() classdb.RenderingDeviceBlendFactor {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceBlendFactor](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_get_src_alpha_blend_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDstAlphaBlendFactor(p_member classdb.RenderingDeviceBlendFactor)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_set_dst_alpha_blend_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDstAlphaBlendFactor() classdb.RenderingDeviceBlendFactor {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceBlendFactor](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_get_dst_alpha_blend_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlphaBlendOp(p_member classdb.RenderingDeviceBlendOperation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_set_alpha_blend_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlphaBlendOp() classdb.RenderingDeviceBlendOperation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceBlendOperation](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_get_alpha_blend_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWriteR(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_set_write_r, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWriteR() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_get_write_r, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWriteG(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_set_write_g, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWriteG() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_get_write_g, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWriteB(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_set_write_b, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWriteB() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_get_write_b, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWriteA(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_set_write_a, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWriteA() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineColorBlendStateAttachment.Bind_get_write_a, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsRDPipelineColorBlendStateAttachment() Expert { return self[0].AsRDPipelineColorBlendStateAttachment() }


//go:nosplit
func (self Simple) AsRDPipelineColorBlendStateAttachment() Simple { return self[0].AsRDPipelineColorBlendStateAttachment() }


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
func init() {classdb.Register("RDPipelineColorBlendStateAttachment", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
