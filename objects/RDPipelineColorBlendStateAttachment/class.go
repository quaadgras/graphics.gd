package RDPipelineColorBlendStateAttachment

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

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
type Instance [1]classdb.RDPipelineColorBlendStateAttachment

/*
Convenience method to perform standard mix blending with straight (non-premultiplied) alpha. This sets [member enable_blend] to [code]true[/code], [member src_color_blend_factor] to [constant RenderingDevice.BLEND_FACTOR_SRC_ALPHA], [member dst_color_blend_factor] to [constant RenderingDevice.BLEND_FACTOR_ONE_MINUS_SRC_ALPHA], [member src_alpha_blend_factor] to [constant RenderingDevice.BLEND_FACTOR_SRC_ALPHA] and [member dst_alpha_blend_factor] to [constant RenderingDevice.BLEND_FACTOR_ONE_MINUS_SRC_ALPHA].
*/
func (self Instance) SetAsMix() {
	class(self).SetAsMix()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RDPipelineColorBlendStateAttachment

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDPipelineColorBlendStateAttachment"))
	return Instance{classdb.RDPipelineColorBlendStateAttachment(object)}
}

func (self Instance) EnableBlend() bool {
	return bool(class(self).GetEnableBlend())
}

func (self Instance) SetEnableBlend(value bool) {
	class(self).SetEnableBlend(value)
}

func (self Instance) SrcColorBlendFactor() classdb.RenderingDeviceBlendFactor {
	return classdb.RenderingDeviceBlendFactor(class(self).GetSrcColorBlendFactor())
}

func (self Instance) SetSrcColorBlendFactor(value classdb.RenderingDeviceBlendFactor) {
	class(self).SetSrcColorBlendFactor(value)
}

func (self Instance) DstColorBlendFactor() classdb.RenderingDeviceBlendFactor {
	return classdb.RenderingDeviceBlendFactor(class(self).GetDstColorBlendFactor())
}

func (self Instance) SetDstColorBlendFactor(value classdb.RenderingDeviceBlendFactor) {
	class(self).SetDstColorBlendFactor(value)
}

func (self Instance) ColorBlendOp() classdb.RenderingDeviceBlendOperation {
	return classdb.RenderingDeviceBlendOperation(class(self).GetColorBlendOp())
}

func (self Instance) SetColorBlendOp(value classdb.RenderingDeviceBlendOperation) {
	class(self).SetColorBlendOp(value)
}

func (self Instance) SrcAlphaBlendFactor() classdb.RenderingDeviceBlendFactor {
	return classdb.RenderingDeviceBlendFactor(class(self).GetSrcAlphaBlendFactor())
}

func (self Instance) SetSrcAlphaBlendFactor(value classdb.RenderingDeviceBlendFactor) {
	class(self).SetSrcAlphaBlendFactor(value)
}

func (self Instance) DstAlphaBlendFactor() classdb.RenderingDeviceBlendFactor {
	return classdb.RenderingDeviceBlendFactor(class(self).GetDstAlphaBlendFactor())
}

func (self Instance) SetDstAlphaBlendFactor(value classdb.RenderingDeviceBlendFactor) {
	class(self).SetDstAlphaBlendFactor(value)
}

func (self Instance) AlphaBlendOp() classdb.RenderingDeviceBlendOperation {
	return classdb.RenderingDeviceBlendOperation(class(self).GetAlphaBlendOp())
}

func (self Instance) SetAlphaBlendOp(value classdb.RenderingDeviceBlendOperation) {
	class(self).SetAlphaBlendOp(value)
}

func (self Instance) WriteR() bool {
	return bool(class(self).GetWriteR())
}

func (self Instance) SetWriteR(value bool) {
	class(self).SetWriteR(value)
}

func (self Instance) WriteG() bool {
	return bool(class(self).GetWriteG())
}

func (self Instance) SetWriteG(value bool) {
	class(self).SetWriteG(value)
}

func (self Instance) WriteB() bool {
	return bool(class(self).GetWriteB())
}

func (self Instance) SetWriteB(value bool) {
	class(self).SetWriteB(value)
}

func (self Instance) WriteA() bool {
	return bool(class(self).GetWriteA())
}

func (self Instance) SetWriteA(value bool) {
	class(self).SetWriteA(value)
}

/*
Convenience method to perform standard mix blending with straight (non-premultiplied) alpha. This sets [member enable_blend] to [code]true[/code], [member src_color_blend_factor] to [constant RenderingDevice.BLEND_FACTOR_SRC_ALPHA], [member dst_color_blend_factor] to [constant RenderingDevice.BLEND_FACTOR_ONE_MINUS_SRC_ALPHA], [member src_alpha_blend_factor] to [constant RenderingDevice.BLEND_FACTOR_SRC_ALPHA] and [member dst_alpha_blend_factor] to [constant RenderingDevice.BLEND_FACTOR_ONE_MINUS_SRC_ALPHA].
*/
//go:nosplit
func (self class) SetAsMix() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_set_as_mix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetEnableBlend(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_set_enable_blend, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableBlend() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_get_enable_blend, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSrcColorBlendFactor(p_member classdb.RenderingDeviceBlendFactor) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_set_src_color_blend_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSrcColorBlendFactor() classdb.RenderingDeviceBlendFactor {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceBlendFactor](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_get_src_color_blend_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDstColorBlendFactor(p_member classdb.RenderingDeviceBlendFactor) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_set_dst_color_blend_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDstColorBlendFactor() classdb.RenderingDeviceBlendFactor {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceBlendFactor](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_get_dst_color_blend_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColorBlendOp(p_member classdb.RenderingDeviceBlendOperation) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_set_color_blend_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetColorBlendOp() classdb.RenderingDeviceBlendOperation {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceBlendOperation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_get_color_blend_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSrcAlphaBlendFactor(p_member classdb.RenderingDeviceBlendFactor) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_set_src_alpha_blend_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSrcAlphaBlendFactor() classdb.RenderingDeviceBlendFactor {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceBlendFactor](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_get_src_alpha_blend_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDstAlphaBlendFactor(p_member classdb.RenderingDeviceBlendFactor) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_set_dst_alpha_blend_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDstAlphaBlendFactor() classdb.RenderingDeviceBlendFactor {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceBlendFactor](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_get_dst_alpha_blend_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlphaBlendOp(p_member classdb.RenderingDeviceBlendOperation) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_set_alpha_blend_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlphaBlendOp() classdb.RenderingDeviceBlendOperation {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceBlendOperation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_get_alpha_blend_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWriteR(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_set_write_r, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetWriteR() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_get_write_r, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWriteG(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_set_write_g, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetWriteG() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_get_write_g, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWriteB(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_set_write_b, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetWriteB() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_get_write_b, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWriteA(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_set_write_a, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetWriteA() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendStateAttachment.Bind_get_write_a, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDPipelineColorBlendStateAttachment() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRDPipelineColorBlendStateAttachment() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("RDPipelineColorBlendStateAttachment", func(ptr gd.Object) any { return classdb.RDPipelineColorBlendStateAttachment(ptr) })
}
