package PortableCompressedTexture2D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Texture2D"
import "grow.graphics/gd/object/Texture"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class allows storing compressed textures as self contained (not imported) resources.
For 2D usage (compressed on disk, uncompressed on VRAM), the lossy and lossless modes are recommended. For 3D usage (compressed on VRAM) it depends on the target platform.
If you intend to only use desktop, S3TC or BPTC are recommended. For only mobile, ETC2 is recommended.
For portable, self contained 3D textures that work on both desktop and mobile, Basis Universal is recommended (although it has a small quality cost and longer compression time as a tradeoff).
This resource is intended to be created from code.

*/
type Simple [1]classdb.PortableCompressedTexture2D
func (self Simple) CreateFromImage(image [1]classdb.Image, compression_mode classdb.PortableCompressedTexture2DCompressionMode, normal_map bool, lossy_quality float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CreateFromImage(image, compression_mode, normal_map, gd.Float(lossy_quality))
}
func (self Simple) GetFormat() classdb.ImageFormat {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ImageFormat(Expert(self).GetFormat())
}
func (self Simple) GetCompressionMode() classdb.PortableCompressedTexture2DCompressionMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.PortableCompressedTexture2DCompressionMode(Expert(self).GetCompressionMode())
}
func (self Simple) SetSizeOverride(size gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSizeOverride(size)
}
func (self Simple) GetSizeOverride() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetSizeOverride())
}
func (self Simple) SetKeepCompressedBuffer(keep bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetKeepCompressedBuffer(keep)
}
func (self Simple) IsKeepingCompressedBuffer() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsKeepingCompressedBuffer())
}
func (self Simple) SetKeepAllCompressedBuffers(keep bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetKeepAllCompressedBuffers(gc, keep)
}
func (self Simple) IsKeepingAllCompressedBuffers() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsKeepingAllCompressedBuffers(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PortableCompressedTexture2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Initializes the compressed texture from a base image. The compression mode must be provided.
[param normal_map] is recommended to ensure optimum quality if this image will be used as a normal map.
If lossy compression is requested, the quality setting can optionally be provided. This maps to Lossy WebP compression quality.
*/
//go:nosplit
func (self class) CreateFromImage(image object.Image, compression_mode classdb.PortableCompressedTexture2DCompressionMode, normal_map bool, lossy_quality gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(image[0].AsPointer())[0])
	callframe.Arg(frame, compression_mode)
	callframe.Arg(frame, normal_map)
	callframe.Arg(frame, lossy_quality)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PortableCompressedTexture2D.Bind_create_from_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Return the image format used (valid after initialized).
*/
//go:nosplit
func (self class) GetFormat() classdb.ImageFormat {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ImageFormat](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PortableCompressedTexture2D.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Return the compression mode used (valid after initialized).
*/
//go:nosplit
func (self class) GetCompressionMode() classdb.PortableCompressedTexture2DCompressionMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.PortableCompressedTexture2DCompressionMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PortableCompressedTexture2D.Bind_get_compression_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSizeOverride(size gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PortableCompressedTexture2D.Bind_set_size_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSizeOverride() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PortableCompressedTexture2D.Bind_get_size_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetKeepCompressedBuffer(keep bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, keep)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PortableCompressedTexture2D.Bind_set_keep_compressed_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsKeepingCompressedBuffer() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PortableCompressedTexture2D.Bind_is_keeping_compressed_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Overrides the flag globally for all textures of this type. This is used primarily by the editor.
*/
//go:nosplit
func (self class) SetKeepAllCompressedBuffers(ctx gd.Lifetime, keep bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, keep)
	var r_ret callframe.Nil
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.PortableCompressedTexture2D.Bind_set_keep_all_compressed_buffers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Return whether the flag is overridden for all textures of this type.
*/
//go:nosplit
func (self class) IsKeepingAllCompressedBuffers(ctx gd.Lifetime) bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.PortableCompressedTexture2D.Bind_is_keeping_all_compressed_buffers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsPortableCompressedTexture2D() Expert { return self[0].AsPortableCompressedTexture2D() }


//go:nosplit
func (self Simple) AsPortableCompressedTexture2D() Simple { return self[0].AsPortableCompressedTexture2D() }


//go:nosplit
func (self class) AsTexture2D() Texture2D.Expert { return self[0].AsTexture2D() }


//go:nosplit
func (self Simple) AsTexture2D() Texture2D.Simple { return self[0].AsTexture2D() }


//go:nosplit
func (self class) AsTexture() Texture.Expert { return self[0].AsTexture() }


//go:nosplit
func (self Simple) AsTexture() Texture.Simple { return self[0].AsTexture() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("PortableCompressedTexture2D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type CompressionMode = classdb.PortableCompressedTexture2DCompressionMode

const (
	CompressionModeLossless CompressionMode = 0
	CompressionModeLossy CompressionMode = 1
	CompressionModeBasisUniversal CompressionMode = 2
	CompressionModeS3tc CompressionMode = 3
	CompressionModeEtc2 CompressionMode = 4
	CompressionModeBptc CompressionMode = 5
)
