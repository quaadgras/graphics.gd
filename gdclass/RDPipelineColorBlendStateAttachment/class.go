package RDPipelineColorBlendStateAttachment

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
type Go [1]classdb.RDPipelineColorBlendStateAttachment

/*
Convenience method to perform standard mix blending with straight (non-premultiplied) alpha. This sets [member enable_blend] to [code]true[/code], [member src_color_blend_factor] to [constant RenderingDevice.BLEND_FACTOR_SRC_ALPHA], [member dst_color_blend_factor] to [constant RenderingDevice.BLEND_FACTOR_ONE_MINUS_SRC_ALPHA], [member src_alpha_blend_factor] to [constant RenderingDevice.BLEND_FACTOR_SRC_ALPHA] and [member dst_alpha_blend_factor] to [constant RenderingDevice.BLEND_FACTOR_ONE_MINUS_SRC_ALPHA].
*/
func (self Go) SetAsMix() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAsMix()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RDPipelineColorBlendStateAttachment
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RDPipelineColorBlendStateAttachment"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) EnableBlend() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetEnableBlend())
}

func (self Go) SetEnableBlend(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnableBlend(value)
}

func (self Go) SrcColorBlendFactor() classdb.RenderingDeviceBlendFactor {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceBlendFactor(class(self).GetSrcColorBlendFactor())
}

func (self Go) SetSrcColorBlendFactor(value classdb.RenderingDeviceBlendFactor) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSrcColorBlendFactor(value)
}

func (self Go) DstColorBlendFactor() classdb.RenderingDeviceBlendFactor {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceBlendFactor(class(self).GetDstColorBlendFactor())
}

func (self Go) SetDstColorBlendFactor(value classdb.RenderingDeviceBlendFactor) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDstColorBlendFactor(value)
}

func (self Go) ColorBlendOp() classdb.RenderingDeviceBlendOperation {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceBlendOperation(class(self).GetColorBlendOp())
}

func (self Go) SetColorBlendOp(value classdb.RenderingDeviceBlendOperation) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetColorBlendOp(value)
}

func (self Go) SrcAlphaBlendFactor() classdb.RenderingDeviceBlendFactor {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceBlendFactor(class(self).GetSrcAlphaBlendFactor())
}

func (self Go) SetSrcAlphaBlendFactor(value classdb.RenderingDeviceBlendFactor) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSrcAlphaBlendFactor(value)
}

func (self Go) DstAlphaBlendFactor() classdb.RenderingDeviceBlendFactor {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceBlendFactor(class(self).GetDstAlphaBlendFactor())
}

func (self Go) SetDstAlphaBlendFactor(value classdb.RenderingDeviceBlendFactor) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDstAlphaBlendFactor(value)
}

func (self Go) AlphaBlendOp() classdb.RenderingDeviceBlendOperation {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceBlendOperation(class(self).GetAlphaBlendOp())
}

func (self Go) SetAlphaBlendOp(value classdb.RenderingDeviceBlendOperation) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAlphaBlendOp(value)
}

func (self Go) WriteR() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetWriteR())
}

func (self Go) SetWriteR(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetWriteR(value)
}

func (self Go) WriteG() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetWriteG())
}

func (self Go) SetWriteG(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetWriteG(value)
}

func (self Go) WriteB() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetWriteB())
}

func (self Go) SetWriteB(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetWriteB(value)
}

func (self Go) WriteA() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetWriteA())
}

func (self Go) SetWriteA(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetWriteA(value)
}

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
func (self class) AsRDPipelineColorBlendStateAttachment() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRDPipelineColorBlendStateAttachment() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("RDPipelineColorBlendStateAttachment", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
