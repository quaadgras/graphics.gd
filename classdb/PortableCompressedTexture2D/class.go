// Package PortableCompressedTexture2D provides methods for working with PortableCompressedTexture2D object instances.
package PortableCompressedTexture2D

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/classdb/Texture"
import "graphics.gd/classdb/Texture2D"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector2"

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
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
This class allows storing compressed textures as self contained (not imported) resources.
For 2D usage (compressed on disk, uncompressed on VRAM), the lossy and lossless modes are recommended. For 3D usage (compressed on VRAM) it depends on the target platform.
If you intend to only use desktop, S3TC or BPTC are recommended. For only mobile, ETC2 is recommended.
For portable, self contained 3D textures that work on both desktop and mobile, Basis Universal is recommended (although it has a small quality cost and longer compression time as a tradeoff).
This resource is intended to be created from code.
*/
type Instance [1]gdclass.PortableCompressedTexture2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPortableCompressedTexture2D() Instance
}

/*
Initializes the compressed texture from a base image. The compression mode must be provided.
[param normal_map] is recommended to ensure optimum quality if this image will be used as a normal map.
If lossy compression is requested, the quality setting can optionally be provided. This maps to Lossy WebP compression quality.
*/
func (self Instance) CreateFromImage(image [1]gdclass.Image, compression_mode gdclass.PortableCompressedTexture2DCompressionMode) { //gd:PortableCompressedTexture2D.create_from_image
	class(self).CreateFromImage(image, compression_mode, false, float64(0.8))
}

/*
Return the image format used (valid after initialized).
*/
func (self Instance) GetFormat() gdclass.ImageFormat { //gd:PortableCompressedTexture2D.get_format
	return gdclass.ImageFormat(class(self).GetFormat())
}

/*
Return the compression mode used (valid after initialized).
*/
func (self Instance) GetCompressionMode() gdclass.PortableCompressedTexture2DCompressionMode { //gd:PortableCompressedTexture2D.get_compression_mode
	return gdclass.PortableCompressedTexture2DCompressionMode(class(self).GetCompressionMode())
}

/*
Overrides the flag globally for all textures of this type. This is used primarily by the editor.
*/
func SetKeepAllCompressedBuffers(keep bool) { //gd:PortableCompressedTexture2D.set_keep_all_compressed_buffers
	self := Instance{}
	class(self).SetKeepAllCompressedBuffers(keep)
}

/*
Return whether the flag is overridden for all textures of this type.
*/
func IsKeepingAllCompressedBuffers() bool { //gd:PortableCompressedTexture2D.is_keeping_all_compressed_buffers
	self := Instance{}
	return bool(class(self).IsKeepingAllCompressedBuffers())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PortableCompressedTexture2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PortableCompressedTexture2D"))
	casted := Instance{*(*gdclass.PortableCompressedTexture2D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) SizeOverride() Vector2.XY {
	return Vector2.XY(class(self).GetSizeOverride())
}

func (self Instance) SetSizeOverride(value Vector2.XY) {
	class(self).SetSizeOverride(Vector2.XY(value))
}

func (self Instance) KeepCompressedBuffer() bool {
	return bool(class(self).IsKeepingCompressedBuffer())
}

func (self Instance) SetKeepCompressedBuffer(value bool) {
	class(self).SetKeepCompressedBuffer(value)
}

/*
Initializes the compressed texture from a base image. The compression mode must be provided.
[param normal_map] is recommended to ensure optimum quality if this image will be used as a normal map.
If lossy compression is requested, the quality setting can optionally be provided. This maps to Lossy WebP compression quality.
*/
//go:nosplit
func (self class) CreateFromImage(image [1]gdclass.Image, compression_mode gdclass.PortableCompressedTexture2DCompressionMode, normal_map bool, lossy_quality float64) { //gd:PortableCompressedTexture2D.create_from_image
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(image[0])[0])
	callframe.Arg(frame, compression_mode)
	callframe.Arg(frame, normal_map)
	callframe.Arg(frame, lossy_quality)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PortableCompressedTexture2D.Bind_create_from_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Return the image format used (valid after initialized).
*/
//go:nosplit
func (self class) GetFormat() gdclass.ImageFormat { //gd:PortableCompressedTexture2D.get_format
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.ImageFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PortableCompressedTexture2D.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Return the compression mode used (valid after initialized).
*/
//go:nosplit
func (self class) GetCompressionMode() gdclass.PortableCompressedTexture2DCompressionMode { //gd:PortableCompressedTexture2D.get_compression_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.PortableCompressedTexture2DCompressionMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PortableCompressedTexture2D.Bind_get_compression_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSizeOverride(size Vector2.XY) { //gd:PortableCompressedTexture2D.set_size_override
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PortableCompressedTexture2D.Bind_set_size_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSizeOverride() Vector2.XY { //gd:PortableCompressedTexture2D.get_size_override
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PortableCompressedTexture2D.Bind_get_size_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetKeepCompressedBuffer(keep bool) { //gd:PortableCompressedTexture2D.set_keep_compressed_buffer
	var frame = callframe.New()
	callframe.Arg(frame, keep)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PortableCompressedTexture2D.Bind_set_keep_compressed_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsKeepingCompressedBuffer() bool { //gd:PortableCompressedTexture2D.is_keeping_compressed_buffer
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PortableCompressedTexture2D.Bind_is_keeping_compressed_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Overrides the flag globally for all textures of this type. This is used primarily by the editor.
*/
//go:nosplit
func (self class) SetKeepAllCompressedBuffers(keep bool) { //gd:PortableCompressedTexture2D.set_keep_all_compressed_buffers
	var frame = callframe.New()
	callframe.Arg(frame, keep)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PortableCompressedTexture2D.Bind_set_keep_all_compressed_buffers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Return whether the flag is overridden for all textures of this type.
*/
//go:nosplit
func (self class) IsKeepingAllCompressedBuffers() bool { //gd:PortableCompressedTexture2D.is_keeping_all_compressed_buffers
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PortableCompressedTexture2D.Bind_is_keeping_all_compressed_buffers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPortableCompressedTexture2D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPortableCompressedTexture2D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsTexture2D() Texture2D.Advanced {
	return *((*Texture2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTexture2D() Texture2D.Instance {
	return *((*Texture2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsTexture() Texture.Advanced { return *((*Texture.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTexture() Texture.Instance {
	return *((*Texture.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Texture2D.Advanced(self.AsTexture2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Texture2D.Instance(self.AsTexture2D()), name)
	}
}
func init() {
	gdclass.Register("PortableCompressedTexture2D", func(ptr gd.Object) any {
		return [1]gdclass.PortableCompressedTexture2D{*(*gdclass.PortableCompressedTexture2D)(unsafe.Pointer(&ptr))}
	})
}

type CompressionMode = gdclass.PortableCompressedTexture2DCompressionMode //gd:PortableCompressedTexture2D.CompressionMode

const (
	CompressionModeLossless       CompressionMode = 0
	CompressionModeLossy          CompressionMode = 1
	CompressionModeBasisUniversal CompressionMode = 2
	CompressionModeS3tc           CompressionMode = 3
	CompressionModeEtc2           CompressionMode = 4
	CompressionModeBptc           CompressionMode = 5
)
