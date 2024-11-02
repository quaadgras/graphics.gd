package RDPipelineColorBlendState

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/variant/Color"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This object is used by [RenderingDevice].
*/
type Instance [1]classdb.RDPipelineColorBlendState

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RDPipelineColorBlendState

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDPipelineColorBlendState"))
	return Instance{classdb.RDPipelineColorBlendState(object)}
}

func (self Instance) EnableLogicOp() bool {
	return bool(class(self).GetEnableLogicOp())
}

func (self Instance) SetEnableLogicOp(value bool) {
	class(self).SetEnableLogicOp(value)
}

func (self Instance) LogicOp() classdb.RenderingDeviceLogicOperation {
	return classdb.RenderingDeviceLogicOperation(class(self).GetLogicOp())
}

func (self Instance) SetLogicOp(value classdb.RenderingDeviceLogicOperation) {
	class(self).SetLogicOp(value)
}

func (self Instance) BlendConstant() Color.RGBA {
	return Color.RGBA(class(self).GetBlendConstant())
}

func (self Instance) SetBlendConstant(value Color.RGBA) {
	class(self).SetBlendConstant(gd.Color(value))
}

func (self Instance) Attachments() gd.Array {
	return gd.Array(class(self).GetAttachments())
}

func (self Instance) SetAttachments(value gd.Array) {
	class(self).SetAttachments(value)
}

//go:nosplit
func (self class) SetEnableLogicOp(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendState.Bind_set_enable_logic_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableLogicOp() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendState.Bind_get_enable_logic_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLogicOp(p_member classdb.RenderingDeviceLogicOperation) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendState.Bind_set_logic_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLogicOp() classdb.RenderingDeviceLogicOperation {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceLogicOperation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendState.Bind_get_logic_op, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBlendConstant(p_member gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendState.Bind_set_blend_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBlendConstant() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendState.Bind_get_blend_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAttachments(attachments gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(attachments))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendState.Bind_set_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAttachments() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendState.Bind_get_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsRDPipelineColorBlendState() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRDPipelineColorBlendState() Instance {
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
	classdb.Register("RDPipelineColorBlendState", func(ptr gd.Object) any { return classdb.RDPipelineColorBlendState(ptr) })
}
