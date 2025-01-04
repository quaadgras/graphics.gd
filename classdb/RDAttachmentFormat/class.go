// Package RDAttachmentFormat provides methods for working with RDAttachmentFormat object instances.
package RDAttachmentFormat

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This object is used by [RenderingDevice].
*/
type Instance [1]gdclass.RDAttachmentFormat
type Any interface {
	gd.IsClass
	AsRDAttachmentFormat() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RDAttachmentFormat

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDAttachmentFormat"))
	return Instance{*(*gdclass.RDAttachmentFormat)(unsafe.Pointer(&object))}
}

func (self Instance) Format() gdclass.RenderingDeviceDataFormat {
	return gdclass.RenderingDeviceDataFormat(class(self).GetFormat())
}

func (self Instance) SetFormat(value gdclass.RenderingDeviceDataFormat) {
	class(self).SetFormat(value)
}

func (self Instance) Samples() gdclass.RenderingDeviceTextureSamples {
	return gdclass.RenderingDeviceTextureSamples(class(self).GetSamples())
}

func (self Instance) SetSamples(value gdclass.RenderingDeviceTextureSamples) {
	class(self).SetSamples(value)
}

func (self Instance) UsageFlags() int {
	return int(int(class(self).GetUsageFlags()))
}

func (self Instance) SetUsageFlags(value int) {
	class(self).SetUsageFlags(gd.Int(value))
}

//go:nosplit
func (self class) SetFormat(p_member gdclass.RenderingDeviceDataFormat) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDAttachmentFormat.Bind_set_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFormat() gdclass.RenderingDeviceDataFormat {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceDataFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDAttachmentFormat.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSamples(p_member gdclass.RenderingDeviceTextureSamples) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDAttachmentFormat.Bind_set_samples, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSamples() gdclass.RenderingDeviceTextureSamples {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceTextureSamples](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDAttachmentFormat.Bind_get_samples, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUsageFlags(p_member gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDAttachmentFormat.Bind_set_usage_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetUsageFlags() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDAttachmentFormat.Bind_get_usage_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDAttachmentFormat() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRDAttachmentFormat() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted       { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

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
	gdclass.Register("RDAttachmentFormat", func(ptr gd.Object) any {
		return [1]gdclass.RDAttachmentFormat{*(*gdclass.RDAttachmentFormat)(unsafe.Pointer(&ptr))}
	})
}
