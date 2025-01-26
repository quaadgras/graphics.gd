// Package RenderSceneBuffersRD provides methods for working with RenderSceneBuffersRD object instances.
package RenderSceneBuffersRD

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
import "graphics.gd/classdb/RenderSceneBuffers"
import "graphics.gd/variant/Vector2i"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
This object manages all 3D rendering buffers for the rendering device based renderers. An instance of this object is created for every viewport that has 3D rendering enabled.
All buffers are organized in [b]contexts[/b]. The default context is called [b]render_buffers[/b] and can contain amongst others the color buffer, depth buffer, velocity buffers, VRS density map and MSAA variants of these buffers.
Buffers are only guaranteed to exist during rendering of the viewport.
[b]Note:[/b] This is an internal rendering server object, do not instantiate this from script.
*/
type Instance [1]gdclass.RenderSceneBuffersRD

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRenderSceneBuffersRD() Instance
}

/*
Returns [code]true[/code] if a cached texture exists for this name.
*/
func (self Instance) HasTexture(context string, name string) bool { //gd:RenderSceneBuffersRD.has_texture
	return bool(class(self).HasTexture(gd.NewStringName(context), gd.NewStringName(name)))
}

/*
Create a new texture with the given definition and cache this under the given name. Will return the existing texture if it already exists.
*/
func (self Instance) CreateTexture(context string, name string, data_format gdclass.RenderingDeviceDataFormat, usage_bits int, texture_samples gdclass.RenderingDeviceTextureSamples, size Vector2i.XY, layers int, mipmaps int, unique bool) RID.Texture { //gd:RenderSceneBuffersRD.create_texture
	return RID.Texture(class(self).CreateTexture(gd.NewStringName(context), gd.NewStringName(name), data_format, gd.Int(usage_bits), texture_samples, gd.Vector2i(size), gd.Int(layers), gd.Int(mipmaps), unique))
}

/*
Create a new texture using the given format and view and cache this under the given name. Will return the existing texture if it already exists.
*/
func (self Instance) CreateTextureFromFormat(context string, name string, format [1]gdclass.RDTextureFormat, view [1]gdclass.RDTextureView, unique bool) RID.Texture { //gd:RenderSceneBuffersRD.create_texture_from_format
	return RID.Texture(class(self).CreateTextureFromFormat(gd.NewStringName(context), gd.NewStringName(name), format, view, unique))
}

/*
Create a new texture view for an existing texture and cache this under the given view_name. Will return the existing teture view if it already exists. Will error if the source texture doesn't exist.
*/
func (self Instance) CreateTextureView(context string, name string, view_name string, view [1]gdclass.RDTextureView) RID.Texture { //gd:RenderSceneBuffersRD.create_texture_view
	return RID.Texture(class(self).CreateTextureView(gd.NewStringName(context), gd.NewStringName(name), gd.NewStringName(view_name), view))
}

/*
Returns a cached texture with this name.
*/
func (self Instance) GetTexture(context string, name string) RID.Texture { //gd:RenderSceneBuffersRD.get_texture
	return RID.Texture(class(self).GetTexture(gd.NewStringName(context), gd.NewStringName(name)))
}

/*
Returns the texture format information with which a cached texture was created.
*/
func (self Instance) GetTextureFormat(context string, name string) [1]gdclass.RDTextureFormat { //gd:RenderSceneBuffersRD.get_texture_format
	return [1]gdclass.RDTextureFormat(class(self).GetTextureFormat(gd.NewStringName(context), gd.NewStringName(name)))
}

/*
Returns a specific slice (layer or mipmap) for a cached texture.
*/
func (self Instance) GetTextureSlice(context string, name string, layer int, mipmap int, layers int, mipmaps int) RID.Texture { //gd:RenderSceneBuffersRD.get_texture_slice
	return RID.Texture(class(self).GetTextureSlice(gd.NewStringName(context), gd.NewStringName(name), gd.Int(layer), gd.Int(mipmap), gd.Int(layers), gd.Int(mipmaps)))
}

/*
Returns a specific view of a slice (layer or mipmap) for a cached texture.
*/
func (self Instance) GetTextureSliceView(context string, name string, layer int, mipmap int, layers int, mipmaps int, view [1]gdclass.RDTextureView) RID.Texture { //gd:RenderSceneBuffersRD.get_texture_slice_view
	return RID.Texture(class(self).GetTextureSliceView(gd.NewStringName(context), gd.NewStringName(name), gd.Int(layer), gd.Int(mipmap), gd.Int(layers), gd.Int(mipmaps), view))
}

/*
Returns the texture size of a given slice of a cached texture.
*/
func (self Instance) GetTextureSliceSize(context string, name string, mipmap int) Vector2i.XY { //gd:RenderSceneBuffersRD.get_texture_slice_size
	return Vector2i.XY(class(self).GetTextureSliceSize(gd.NewStringName(context), gd.NewStringName(name), gd.Int(mipmap)))
}

/*
Frees all buffers related to this context.
*/
func (self Instance) ClearContext(context string) { //gd:RenderSceneBuffersRD.clear_context
	class(self).ClearContext(gd.NewStringName(context))
}

/*
Returns the color texture we are rendering 3D content to. If multiview is used this will be a texture array with all views.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
func (self Instance) GetColorTexture() RID.Texture { //gd:RenderSceneBuffersRD.get_color_texture
	return RID.Texture(class(self).GetColorTexture(false))
}

/*
Returns the specified layer from the color texture we are rendering 3D content to.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
func (self Instance) GetColorLayer(layer int) RID.Texture { //gd:RenderSceneBuffersRD.get_color_layer
	return RID.Texture(class(self).GetColorLayer(gd.Int(layer), false))
}

/*
Returns the depth texture we are rendering 3D content to. If multiview is used this will be a texture array with all views.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
func (self Instance) GetDepthTexture() RID.Texture { //gd:RenderSceneBuffersRD.get_depth_texture
	return RID.Texture(class(self).GetDepthTexture(false))
}

/*
Returns the specified layer from the depth texture we are rendering 3D content to.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
func (self Instance) GetDepthLayer(layer int) RID.Texture { //gd:RenderSceneBuffersRD.get_depth_layer
	return RID.Texture(class(self).GetDepthLayer(gd.Int(layer), false))
}

/*
Returns the velocity texture we are rendering 3D content to. If multiview is used this will be a texture array with all views.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
func (self Instance) GetVelocityTexture() RID.Texture { //gd:RenderSceneBuffersRD.get_velocity_texture
	return RID.Texture(class(self).GetVelocityTexture(false))
}

/*
Returns the specified layer from the velocity texture we are rendering 3D content to.
*/
func (self Instance) GetVelocityLayer(layer int) RID.Texture { //gd:RenderSceneBuffersRD.get_velocity_layer
	return RID.Texture(class(self).GetVelocityLayer(gd.Int(layer), false))
}

/*
Returns the render target associated with this buffers object.
*/
func (self Instance) GetRenderTarget() RID.Framebuffer { //gd:RenderSceneBuffersRD.get_render_target
	return RID.Framebuffer(class(self).GetRenderTarget())
}

/*
Returns the view count for the associated viewport.
*/
func (self Instance) GetViewCount() int { //gd:RenderSceneBuffersRD.get_view_count
	return int(int(class(self).GetViewCount()))
}

/*
Returns the internal size of the render buffer (size before upscaling) with which textures are created by default.
*/
func (self Instance) GetInternalSize() Vector2i.XY { //gd:RenderSceneBuffersRD.get_internal_size
	return Vector2i.XY(class(self).GetInternalSize())
}

/*
Returns the target size of the render buffer (size after upscaling).
*/
func (self Instance) GetTargetSize() Vector2i.XY { //gd:RenderSceneBuffersRD.get_target_size
	return Vector2i.XY(class(self).GetTargetSize())
}

/*
Returns the scaling mode used for upscaling.
*/
func (self Instance) GetScaling3dMode() gdclass.RenderingServerViewportScaling3DMode { //gd:RenderSceneBuffersRD.get_scaling_3d_mode
	return gdclass.RenderingServerViewportScaling3DMode(class(self).GetScaling3dMode())
}

/*
Returns the FSR sharpness value used while rendering the 3D content (if [method get_scaling_3d_mode] is an FSR mode).
*/
func (self Instance) GetFsrSharpness() Float.X { //gd:RenderSceneBuffersRD.get_fsr_sharpness
	return Float.X(Float.X(class(self).GetFsrSharpness()))
}

/*
Returns the applied 3D MSAA mode for this viewport.
*/
func (self Instance) GetMsaa3d() gdclass.RenderingServerViewportMSAA { //gd:RenderSceneBuffersRD.get_msaa_3d
	return gdclass.RenderingServerViewportMSAA(class(self).GetMsaa3d())
}

/*
Returns the number of MSAA samples used.
*/
func (self Instance) GetTextureSamples() gdclass.RenderingDeviceTextureSamples { //gd:RenderSceneBuffersRD.get_texture_samples
	return gdclass.RenderingDeviceTextureSamples(class(self).GetTextureSamples())
}

/*
Returns the screen-space antialiasing method applied.
*/
func (self Instance) GetScreenSpaceAa() gdclass.RenderingServerViewportScreenSpaceAA { //gd:RenderSceneBuffersRD.get_screen_space_aa
	return gdclass.RenderingServerViewportScreenSpaceAA(class(self).GetScreenSpaceAa())
}

/*
Returns [code]true[/code] if TAA is enabled.
*/
func (self Instance) GetUseTaa() bool { //gd:RenderSceneBuffersRD.get_use_taa
	return bool(class(self).GetUseTaa())
}

/*
Returns [code]true[/code] if debanding is enabled.
*/
func (self Instance) GetUseDebanding() bool { //gd:RenderSceneBuffersRD.get_use_debanding
	return bool(class(self).GetUseDebanding())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RenderSceneBuffersRD

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RenderSceneBuffersRD"))
	casted := Instance{*(*gdclass.RenderSceneBuffersRD)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Returns [code]true[/code] if a cached texture exists for this name.
*/
//go:nosplit
func (self class) HasTexture(context gd.StringName, name gd.StringName) bool { //gd:RenderSceneBuffersRD.has_texture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(context))
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_has_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Create a new texture with the given definition and cache this under the given name. Will return the existing texture if it already exists.
*/
//go:nosplit
func (self class) CreateTexture(context gd.StringName, name gd.StringName, data_format gdclass.RenderingDeviceDataFormat, usage_bits gd.Int, texture_samples gdclass.RenderingDeviceTextureSamples, size gd.Vector2i, layers gd.Int, mipmaps gd.Int, unique bool) gd.RID { //gd:RenderSceneBuffersRD.create_texture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(context))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, data_format)
	callframe.Arg(frame, usage_bits)
	callframe.Arg(frame, texture_samples)
	callframe.Arg(frame, size)
	callframe.Arg(frame, layers)
	callframe.Arg(frame, mipmaps)
	callframe.Arg(frame, unique)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_create_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Create a new texture using the given format and view and cache this under the given name. Will return the existing texture if it already exists.
*/
//go:nosplit
func (self class) CreateTextureFromFormat(context gd.StringName, name gd.StringName, format [1]gdclass.RDTextureFormat, view [1]gdclass.RDTextureView, unique bool) gd.RID { //gd:RenderSceneBuffersRD.create_texture_from_format
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(context))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(format[0])[0])
	callframe.Arg(frame, pointers.Get(view[0])[0])
	callframe.Arg(frame, unique)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_create_texture_from_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Create a new texture view for an existing texture and cache this under the given view_name. Will return the existing teture view if it already exists. Will error if the source texture doesn't exist.
*/
//go:nosplit
func (self class) CreateTextureView(context gd.StringName, name gd.StringName, view_name gd.StringName, view [1]gdclass.RDTextureView) gd.RID { //gd:RenderSceneBuffersRD.create_texture_view
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(context))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(view_name))
	callframe.Arg(frame, pointers.Get(view[0])[0])
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_create_texture_view, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a cached texture with this name.
*/
//go:nosplit
func (self class) GetTexture(context gd.StringName, name gd.StringName) gd.RID { //gd:RenderSceneBuffersRD.get_texture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(context))
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the texture format information with which a cached texture was created.
*/
//go:nosplit
func (self class) GetTextureFormat(context gd.StringName, name gd.StringName) [1]gdclass.RDTextureFormat { //gd:RenderSceneBuffersRD.get_texture_format
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(context))
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_texture_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.RDTextureFormat{gd.PointerWithOwnershipTransferredToGo[gdclass.RDTextureFormat](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns a specific slice (layer or mipmap) for a cached texture.
*/
//go:nosplit
func (self class) GetTextureSlice(context gd.StringName, name gd.StringName, layer gd.Int, mipmap gd.Int, layers gd.Int, mipmaps gd.Int) gd.RID { //gd:RenderSceneBuffersRD.get_texture_slice
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(context))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, layer)
	callframe.Arg(frame, mipmap)
	callframe.Arg(frame, layers)
	callframe.Arg(frame, mipmaps)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_texture_slice, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a specific view of a slice (layer or mipmap) for a cached texture.
*/
//go:nosplit
func (self class) GetTextureSliceView(context gd.StringName, name gd.StringName, layer gd.Int, mipmap gd.Int, layers gd.Int, mipmaps gd.Int, view [1]gdclass.RDTextureView) gd.RID { //gd:RenderSceneBuffersRD.get_texture_slice_view
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(context))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, layer)
	callframe.Arg(frame, mipmap)
	callframe.Arg(frame, layers)
	callframe.Arg(frame, mipmaps)
	callframe.Arg(frame, pointers.Get(view[0])[0])
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_texture_slice_view, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the texture size of a given slice of a cached texture.
*/
//go:nosplit
func (self class) GetTextureSliceSize(context gd.StringName, name gd.StringName, mipmap gd.Int) gd.Vector2i { //gd:RenderSceneBuffersRD.get_texture_slice_size
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(context))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, mipmap)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_texture_slice_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Frees all buffers related to this context.
*/
//go:nosplit
func (self class) ClearContext(context gd.StringName) { //gd:RenderSceneBuffersRD.clear_context
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(context))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_clear_context, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the color texture we are rendering 3D content to. If multiview is used this will be a texture array with all views.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
//go:nosplit
func (self class) GetColorTexture(msaa bool) gd.RID { //gd:RenderSceneBuffersRD.get_color_texture
	var frame = callframe.New()
	callframe.Arg(frame, msaa)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_color_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the specified layer from the color texture we are rendering 3D content to.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
//go:nosplit
func (self class) GetColorLayer(layer gd.Int, msaa bool) gd.RID { //gd:RenderSceneBuffersRD.get_color_layer
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, msaa)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_color_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the depth texture we are rendering 3D content to. If multiview is used this will be a texture array with all views.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
//go:nosplit
func (self class) GetDepthTexture(msaa bool) gd.RID { //gd:RenderSceneBuffersRD.get_depth_texture
	var frame = callframe.New()
	callframe.Arg(frame, msaa)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_depth_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the specified layer from the depth texture we are rendering 3D content to.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
//go:nosplit
func (self class) GetDepthLayer(layer gd.Int, msaa bool) gd.RID { //gd:RenderSceneBuffersRD.get_depth_layer
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, msaa)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_depth_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the velocity texture we are rendering 3D content to. If multiview is used this will be a texture array with all views.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
//go:nosplit
func (self class) GetVelocityTexture(msaa bool) gd.RID { //gd:RenderSceneBuffersRD.get_velocity_texture
	var frame = callframe.New()
	callframe.Arg(frame, msaa)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_velocity_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the specified layer from the velocity texture we are rendering 3D content to.
*/
//go:nosplit
func (self class) GetVelocityLayer(layer gd.Int, msaa bool) gd.RID { //gd:RenderSceneBuffersRD.get_velocity_layer
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, msaa)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_velocity_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the render target associated with this buffers object.
*/
//go:nosplit
func (self class) GetRenderTarget() gd.RID { //gd:RenderSceneBuffersRD.get_render_target
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_render_target, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the view count for the associated viewport.
*/
//go:nosplit
func (self class) GetViewCount() gd.Int { //gd:RenderSceneBuffersRD.get_view_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_view_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the internal size of the render buffer (size before upscaling) with which textures are created by default.
*/
//go:nosplit
func (self class) GetInternalSize() gd.Vector2i { //gd:RenderSceneBuffersRD.get_internal_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_internal_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the target size of the render buffer (size after upscaling).
*/
//go:nosplit
func (self class) GetTargetSize() gd.Vector2i { //gd:RenderSceneBuffersRD.get_target_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_target_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the scaling mode used for upscaling.
*/
//go:nosplit
func (self class) GetScaling3dMode() gdclass.RenderingServerViewportScaling3DMode { //gd:RenderSceneBuffersRD.get_scaling_3d_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingServerViewportScaling3DMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_scaling_3d_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the FSR sharpness value used while rendering the 3D content (if [method get_scaling_3d_mode] is an FSR mode).
*/
//go:nosplit
func (self class) GetFsrSharpness() gd.Float { //gd:RenderSceneBuffersRD.get_fsr_sharpness
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_fsr_sharpness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the applied 3D MSAA mode for this viewport.
*/
//go:nosplit
func (self class) GetMsaa3d() gdclass.RenderingServerViewportMSAA { //gd:RenderSceneBuffersRD.get_msaa_3d
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingServerViewportMSAA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_msaa_3d, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of MSAA samples used.
*/
//go:nosplit
func (self class) GetTextureSamples() gdclass.RenderingDeviceTextureSamples { //gd:RenderSceneBuffersRD.get_texture_samples
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceTextureSamples](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_texture_samples, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the screen-space antialiasing method applied.
*/
//go:nosplit
func (self class) GetScreenSpaceAa() gdclass.RenderingServerViewportScreenSpaceAA { //gd:RenderSceneBuffersRD.get_screen_space_aa
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingServerViewportScreenSpaceAA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_screen_space_aa, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if TAA is enabled.
*/
//go:nosplit
func (self class) GetUseTaa() bool { //gd:RenderSceneBuffersRD.get_use_taa
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_use_taa, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if debanding is enabled.
*/
//go:nosplit
func (self class) GetUseDebanding() bool { //gd:RenderSceneBuffersRD.get_use_debanding
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_use_debanding, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRenderSceneBuffersRD() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRenderSceneBuffersRD() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRenderSceneBuffers() RenderSceneBuffers.Advanced {
	return *((*RenderSceneBuffers.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRenderSceneBuffers() RenderSceneBuffers.Instance {
	return *((*RenderSceneBuffers.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(RenderSceneBuffers.Advanced(self.AsRenderSceneBuffers()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RenderSceneBuffers.Instance(self.AsRenderSceneBuffers()), name)
	}
}
func init() {
	gdclass.Register("RenderSceneBuffersRD", func(ptr gd.Object) any {
		return [1]gdclass.RenderSceneBuffersRD{*(*gdclass.RenderSceneBuffersRD)(unsafe.Pointer(&ptr))}
	})
}
