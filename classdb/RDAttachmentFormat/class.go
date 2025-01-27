// Package RDAttachmentFormat provides methods for working with RDAttachmentFormat object instances.
package RDAttachmentFormat

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable

/*
This object is used by [RenderingDevice].
*/
type Instance [1]gdclass.RDAttachmentFormat

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRDAttachmentFormat() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RDAttachmentFormat

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDAttachmentFormat"))
	casted := Instance{*(*gdclass.RDAttachmentFormat)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
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
func (self class) SetFormat(p_member gdclass.RenderingDeviceDataFormat) { //gd:RDAttachmentFormat.set_format
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDAttachmentFormat.Bind_set_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFormat() gdclass.RenderingDeviceDataFormat { //gd:RDAttachmentFormat.get_format
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceDataFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDAttachmentFormat.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSamples(p_member gdclass.RenderingDeviceTextureSamples) { //gd:RDAttachmentFormat.set_samples
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDAttachmentFormat.Bind_set_samples, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSamples() gdclass.RenderingDeviceTextureSamples { //gd:RDAttachmentFormat.get_samples
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceTextureSamples](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDAttachmentFormat.Bind_get_samples, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUsageFlags(p_member gd.Int) { //gd:RDAttachmentFormat.set_usage_flags
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDAttachmentFormat.Bind_set_usage_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUsageFlags() gd.Int { //gd:RDAttachmentFormat.get_usage_flags
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDAttachmentFormat.Bind_get_usage_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDAttachmentFormat() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRDAttachmentFormat() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("RDAttachmentFormat", func(ptr gd.Object) any {
		return [1]gdclass.RDAttachmentFormat{*(*gdclass.RDAttachmentFormat)(unsafe.Pointer(&ptr))}
	})
}
