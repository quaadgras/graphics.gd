// Package RDFramebufferPass provides methods for working with RDFramebufferPass object instances.
package RDFramebufferPass

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This class contains the list of attachment descriptions for a framebuffer pass. Each points with an index to a previously supplied list of texture attachments.
Multipass framebuffers can optimize some configurations in mobile. On desktop, they provide little to no advantage.
This object is used by [RenderingDevice].
*/
type Instance [1]gdclass.RDFramebufferPass

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRDFramebufferPass() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RDFramebufferPass

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDFramebufferPass"))
	casted := Instance{*(*gdclass.RDFramebufferPass)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) ColorAttachments() []int32 {
	return []int32(class(self).GetColorAttachments().AsSlice())
}

func (self Instance) SetColorAttachments(value []int32) {
	class(self).SetColorAttachments(gd.NewPackedInt32Slice(value))
}

func (self Instance) InputAttachments() []int32 {
	return []int32(class(self).GetInputAttachments().AsSlice())
}

func (self Instance) SetInputAttachments(value []int32) {
	class(self).SetInputAttachments(gd.NewPackedInt32Slice(value))
}

func (self Instance) ResolveAttachments() []int32 {
	return []int32(class(self).GetResolveAttachments().AsSlice())
}

func (self Instance) SetResolveAttachments(value []int32) {
	class(self).SetResolveAttachments(gd.NewPackedInt32Slice(value))
}

func (self Instance) PreserveAttachments() []int32 {
	return []int32(class(self).GetPreserveAttachments().AsSlice())
}

func (self Instance) SetPreserveAttachments(value []int32) {
	class(self).SetPreserveAttachments(gd.NewPackedInt32Slice(value))
}

func (self Instance) DepthAttachment() int {
	return int(int(class(self).GetDepthAttachment()))
}

func (self Instance) SetDepthAttachment(value int) {
	class(self).SetDepthAttachment(gd.Int(value))
}

//go:nosplit
func (self class) SetColorAttachments(p_member gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(p_member))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_set_color_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetColorAttachments() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_get_color_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInputAttachments(p_member gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(p_member))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_set_input_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetInputAttachments() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_get_input_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetResolveAttachments(p_member gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(p_member))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_set_resolve_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetResolveAttachments() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_get_resolve_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPreserveAttachments(p_member gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(p_member))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_set_preserve_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPreserveAttachments() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_get_preserve_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDepthAttachment(p_member gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_set_depth_attachment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepthAttachment() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_get_depth_attachment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDFramebufferPass() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRDFramebufferPass() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("RDFramebufferPass", func(ptr gd.Object) any {
		return [1]gdclass.RDFramebufferPass{*(*gdclass.RDFramebufferPass)(unsafe.Pointer(&ptr))}
	})
}
