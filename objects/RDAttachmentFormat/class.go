package RDAttachmentFormat

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This object is used by [RenderingDevice].
*/
type Instance [1]classdb.RDAttachmentFormat
type Any interface {
	gd.IsClass
	AsRDAttachmentFormat() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RDAttachmentFormat

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDAttachmentFormat"))
	return Instance{classdb.RDAttachmentFormat(object)}
}

func (self Instance) Format() classdb.RenderingDeviceDataFormat {
	return classdb.RenderingDeviceDataFormat(class(self).GetFormat())
}

func (self Instance) SetFormat(value classdb.RenderingDeviceDataFormat) {
	class(self).SetFormat(value)
}

func (self Instance) Samples() classdb.RenderingDeviceTextureSamples {
	return classdb.RenderingDeviceTextureSamples(class(self).GetSamples())
}

func (self Instance) SetSamples(value classdb.RenderingDeviceTextureSamples) {
	class(self).SetSamples(value)
}

func (self Instance) UsageFlags() int {
	return int(int(class(self).GetUsageFlags()))
}

func (self Instance) SetUsageFlags(value int) {
	class(self).SetUsageFlags(gd.Int(value))
}

//go:nosplit
func (self class) SetFormat(p_member classdb.RenderingDeviceDataFormat) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDAttachmentFormat.Bind_set_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFormat() classdb.RenderingDeviceDataFormat {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceDataFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDAttachmentFormat.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSamples(p_member classdb.RenderingDeviceTextureSamples) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDAttachmentFormat.Bind_set_samples, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSamples() classdb.RenderingDeviceTextureSamples {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceTextureSamples](frame)
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
	classdb.Register("RDAttachmentFormat", func(ptr gd.Object) any { return classdb.RDAttachmentFormat(ptr) })
}
