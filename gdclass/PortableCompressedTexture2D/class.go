package PortableCompressedTexture2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Texture2D"
import "grow.graphics/gd/gdclass/Texture"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
This class allows storing compressed textures as self contained (not imported) resources.
For 2D usage (compressed on disk, uncompressed on VRAM), the lossy and lossless modes are recommended. For 3D usage (compressed on VRAM) it depends on the target platform.
If you intend to only use desktop, S3TC or BPTC are recommended. For only mobile, ETC2 is recommended.
For portable, self contained 3D textures that work on both desktop and mobile, Basis Universal is recommended (although it has a small quality cost and longer compression time as a tradeoff).
This resource is intended to be created from code.

*/
type Go [1]classdb.PortableCompressedTexture2D

/*
Initializes the compressed texture from a base image. The compression mode must be provided.
[param normal_map] is recommended to ensure optimum quality if this image will be used as a normal map.
If lossy compression is requested, the quality setting can optionally be provided. This maps to Lossy WebP compression quality.
*/
func (self Go) CreateFromImage(image gdclass.Image, compression_mode classdb.PortableCompressedTexture2DCompressionMode) {
	class(self).CreateFromImage(image, compression_mode, false, gd.Float(0.8))
}

/*
Return the image format used (valid after initialized).
*/
func (self Go) GetFormat() classdb.ImageFormat {
	return classdb.ImageFormat(class(self).GetFormat())
}

/*
Return the compression mode used (valid after initialized).
*/
func (self Go) GetCompressionMode() classdb.PortableCompressedTexture2DCompressionMode {
	return classdb.PortableCompressedTexture2DCompressionMode(class(self).GetCompressionMode())
}

/*
Overrides the flag globally for all textures of this type. This is used primarily by the editor.
*/
func (self Go) SetKeepAllCompressedBuffers(keep bool) {
	class(self).SetKeepAllCompressedBuffers(keep)
}

/*
Return whether the flag is overridden for all textures of this type.
*/
func (self Go) IsKeepingAllCompressedBuffers() bool {
	return bool(class(self).IsKeepingAllCompressedBuffers())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PortableCompressedTexture2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PortableCompressedTexture2D"))
	return Go{classdb.PortableCompressedTexture2D(object)}
}

func (self Go) SizeOverride() gd.Vector2 {
		return gd.Vector2(class(self).GetSizeOverride())
}

func (self Go) SetSizeOverride(value gd.Vector2) {
	class(self).SetSizeOverride(value)
}

func (self Go) KeepCompressedBuffer() bool {
		return bool(class(self).IsKeepingCompressedBuffer())
}

func (self Go) SetKeepCompressedBuffer(value bool) {
	class(self).SetKeepCompressedBuffer(value)
}

/*
Initializes the compressed texture from a base image. The compression mode must be provided.
[param normal_map] is recommended to ensure optimum quality if this image will be used as a normal map.
If lossy compression is requested, the quality setting can optionally be provided. This maps to Lossy WebP compression quality.
*/
//go:nosplit
func (self class) CreateFromImage(image gdclass.Image, compression_mode classdb.PortableCompressedTexture2DCompressionMode, normal_map bool, lossy_quality gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(image[0])[0])
	callframe.Arg(frame, compression_mode)
	callframe.Arg(frame, normal_map)
	callframe.Arg(frame, lossy_quality)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PortableCompressedTexture2D.Bind_create_from_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Return the image format used (valid after initialized).
*/
//go:nosplit
func (self class) GetFormat() classdb.ImageFormat {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ImageFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PortableCompressedTexture2D.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Return the compression mode used (valid after initialized).
*/
//go:nosplit
func (self class) GetCompressionMode() classdb.PortableCompressedTexture2DCompressionMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.PortableCompressedTexture2DCompressionMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PortableCompressedTexture2D.Bind_get_compression_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSizeOverride(size gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PortableCompressedTexture2D.Bind_set_size_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSizeOverride() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PortableCompressedTexture2D.Bind_get_size_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetKeepCompressedBuffer(keep bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, keep)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PortableCompressedTexture2D.Bind_set_keep_compressed_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsKeepingCompressedBuffer() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PortableCompressedTexture2D.Bind_is_keeping_compressed_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Overrides the flag globally for all textures of this type. This is used primarily by the editor.
*/
//go:nosplit
func (self class) SetKeepAllCompressedBuffers(keep bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, keep)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PortableCompressedTexture2D.Bind_set_keep_all_compressed_buffers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Return whether the flag is overridden for all textures of this type.
*/
//go:nosplit
func (self class) IsKeepingAllCompressedBuffers() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PortableCompressedTexture2D.Bind_is_keeping_all_compressed_buffers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPortableCompressedTexture2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPortableCompressedTexture2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture2D() Texture2D.GD { return *((*Texture2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture2D() Texture2D.Go { return *((*Texture2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture() Texture.GD { return *((*Texture.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture() Texture.Go { return *((*Texture.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTexture2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTexture2D(), name)
	}
}
func init() {classdb.Register("PortableCompressedTexture2D", func(ptr gd.Object) any { return classdb.PortableCompressedTexture2D(ptr) })}
type CompressionMode = classdb.PortableCompressedTexture2DCompressionMode

const (
	CompressionModeLossless CompressionMode = 0
	CompressionModeLossy CompressionMode = 1
	CompressionModeBasisUniversal CompressionMode = 2
	CompressionModeS3tc CompressionMode = 3
	CompressionModeEtc2 CompressionMode = 4
	CompressionModeBptc CompressionMode = 5
)
