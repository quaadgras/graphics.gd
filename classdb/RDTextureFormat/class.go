// Package RDTextureFormat provides methods for working with RDTextureFormat object instances.
package RDTextureFormat

import "unsafe"
import "reflect"
import "slices"
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
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Packed"

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
var _ Path.ToNode
var _ Packed.Bytes
var _ = slices.Delete[[]struct{}, struct{}]

/*
This object is used by [RenderingDevice].
*/
type Instance [1]gdclass.RDTextureFormat

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRDTextureFormat() Instance
}

func (self Instance) AddShareableFormat(format gdclass.RenderingDeviceDataFormat) { //gd:RDTextureFormat.add_shareable_format
	class(self).AddShareableFormat(format)
}
func (self Instance) RemoveShareableFormat(format gdclass.RenderingDeviceDataFormat) { //gd:RDTextureFormat.remove_shareable_format
	class(self).RemoveShareableFormat(format)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RDTextureFormat

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDTextureFormat"))
	casted := Instance{*(*gdclass.RDTextureFormat)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Format() gdclass.RenderingDeviceDataFormat {
	return gdclass.RenderingDeviceDataFormat(class(self).GetFormat())
}

func (self Instance) SetFormat(value gdclass.RenderingDeviceDataFormat) {
	class(self).SetFormat(value)
}

func (self Instance) Width() int {
	return int(int(class(self).GetWidth()))
}

func (self Instance) SetWidth(value int) {
	class(self).SetWidth(gd.Int(value))
}

func (self Instance) Height() int {
	return int(int(class(self).GetHeight()))
}

func (self Instance) SetHeight(value int) {
	class(self).SetHeight(gd.Int(value))
}

func (self Instance) Depth() int {
	return int(int(class(self).GetDepth()))
}

func (self Instance) SetDepth(value int) {
	class(self).SetDepth(gd.Int(value))
}

func (self Instance) ArrayLayers() int {
	return int(int(class(self).GetArrayLayers()))
}

func (self Instance) SetArrayLayers(value int) {
	class(self).SetArrayLayers(gd.Int(value))
}

func (self Instance) Mipmaps() int {
	return int(int(class(self).GetMipmaps()))
}

func (self Instance) SetMipmaps(value int) {
	class(self).SetMipmaps(gd.Int(value))
}

func (self Instance) TextureType() gdclass.RenderingDeviceTextureType {
	return gdclass.RenderingDeviceTextureType(class(self).GetTextureType())
}

func (self Instance) SetTextureType(value gdclass.RenderingDeviceTextureType) {
	class(self).SetTextureType(value)
}

func (self Instance) Samples() gdclass.RenderingDeviceTextureSamples {
	return gdclass.RenderingDeviceTextureSamples(class(self).GetSamples())
}

func (self Instance) SetSamples(value gdclass.RenderingDeviceTextureSamples) {
	class(self).SetSamples(value)
}

func (self Instance) UsageBits() gdclass.RenderingDeviceTextureUsageBits {
	return gdclass.RenderingDeviceTextureUsageBits(class(self).GetUsageBits())
}

func (self Instance) SetUsageBits(value gdclass.RenderingDeviceTextureUsageBits) {
	class(self).SetUsageBits(value)
}

//go:nosplit
func (self class) SetFormat(p_member gdclass.RenderingDeviceDataFormat) { //gd:RDTextureFormat.set_format
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_set_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFormat() gdclass.RenderingDeviceDataFormat { //gd:RDTextureFormat.get_format
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceDataFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWidth(p_member gd.Int) { //gd:RDTextureFormat.set_width
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetWidth() gd.Int { //gd:RDTextureFormat.get_width
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHeight(p_member gd.Int) { //gd:RDTextureFormat.set_height
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHeight() gd.Int { //gd:RDTextureFormat.get_height
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDepth(p_member gd.Int) { //gd:RDTextureFormat.set_depth
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_set_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepth() gd.Int { //gd:RDTextureFormat.get_depth
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_get_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetArrayLayers(p_member gd.Int) { //gd:RDTextureFormat.set_array_layers
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_set_array_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetArrayLayers() gd.Int { //gd:RDTextureFormat.get_array_layers
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_get_array_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMipmaps(p_member gd.Int) { //gd:RDTextureFormat.set_mipmaps
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_set_mipmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMipmaps() gd.Int { //gd:RDTextureFormat.get_mipmaps
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_get_mipmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureType(p_member gdclass.RenderingDeviceTextureType) { //gd:RDTextureFormat.set_texture_type
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_set_texture_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureType() gdclass.RenderingDeviceTextureType { //gd:RDTextureFormat.get_texture_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceTextureType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_get_texture_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSamples(p_member gdclass.RenderingDeviceTextureSamples) { //gd:RDTextureFormat.set_samples
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_set_samples, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSamples() gdclass.RenderingDeviceTextureSamples { //gd:RDTextureFormat.get_samples
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceTextureSamples](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_get_samples, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUsageBits(p_member gdclass.RenderingDeviceTextureUsageBits) { //gd:RDTextureFormat.set_usage_bits
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_set_usage_bits, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUsageBits() gdclass.RenderingDeviceTextureUsageBits { //gd:RDTextureFormat.get_usage_bits
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceTextureUsageBits](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_get_usage_bits, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AddShareableFormat(format gdclass.RenderingDeviceDataFormat) { //gd:RDTextureFormat.add_shareable_format
	var frame = callframe.New()
	callframe.Arg(frame, format)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_add_shareable_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) RemoveShareableFormat(format gdclass.RenderingDeviceDataFormat) { //gd:RDTextureFormat.remove_shareable_format
	var frame = callframe.New()
	callframe.Arg(frame, format)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDTextureFormat.Bind_remove_shareable_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsRDTextureFormat() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRDTextureFormat() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("RDTextureFormat", func(ptr gd.Object) any {
		return [1]gdclass.RDTextureFormat{*(*gdclass.RDTextureFormat)(unsafe.Pointer(&ptr))}
	})
}
