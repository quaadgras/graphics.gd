package RenderSceneBuffersRD

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/RenderSceneBuffers"
import "graphics.gd/variant/Vector2i"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This object manages all 3D rendering buffers for the rendering device based renderers. An instance of this object is created for every viewport that has 3D rendering enabled.
All buffers are organized in [b]contexts[/b]. The default context is called [b]render_buffers[/b] and can contain amongst others the color buffer, depth buffer, velocity buffers, VRS density map and MSAA variants of these buffers.
Buffers are only guaranteed to exist during rendering of the viewport.
[b]Note:[/b] This is an internal rendering server object, do not instantiate this from script.
*/
type Instance [1]gdclass.RenderSceneBuffersRD
type Any interface {
	gd.IsClass
	AsRenderSceneBuffersRD() Instance
}

/*
Returns [code]true[/code] if a cached texture exists for this name.
*/
func (self Instance) HasTexture(context string, name string) bool {
	return bool(class(self).HasTexture(gd.NewStringName(context), gd.NewStringName(name)))
}

/*
Create a new texture with the given definition and cache this under the given name. Will return the existing texture if it already exists.
*/
func (self Instance) CreateTexture(context string, name string, data_format gdclass.RenderingDeviceDataFormat, usage_bits int, texture_samples gdclass.RenderingDeviceTextureSamples, size Vector2i.XY, layers int, mipmaps int, unique bool) Resource.ID {
	return Resource.ID(class(self).CreateTexture(gd.NewStringName(context), gd.NewStringName(name), data_format, gd.Int(usage_bits), texture_samples, gd.Vector2i(size), gd.Int(layers), gd.Int(mipmaps), unique))
}

/*
Create a new texture using the given format and view and cache this under the given name. Will return the existing texture if it already exists.
*/
func (self Instance) CreateTextureFromFormat(context string, name string, format [1]gdclass.RDTextureFormat, view [1]gdclass.RDTextureView, unique bool) Resource.ID {
	return Resource.ID(class(self).CreateTextureFromFormat(gd.NewStringName(context), gd.NewStringName(name), format, view, unique))
}

/*
Create a new texture view for an existing texture and cache this under the given view_name. Will return the existing teture view if it already exists. Will error if the source texture doesn't exist.
*/
func (self Instance) CreateTextureView(context string, name string, view_name string, view [1]gdclass.RDTextureView) Resource.ID {
	return Resource.ID(class(self).CreateTextureView(gd.NewStringName(context), gd.NewStringName(name), gd.NewStringName(view_name), view))
}

/*
Returns a cached texture with this name.
*/
func (self Instance) GetTexture(context string, name string) Resource.ID {
	return Resource.ID(class(self).GetTexture(gd.NewStringName(context), gd.NewStringName(name)))
}

/*
Returns the texture format information with which a cached texture was created.
*/
func (self Instance) GetTextureFormat(context string, name string) [1]gdclass.RDTextureFormat {
	return [1]gdclass.RDTextureFormat(class(self).GetTextureFormat(gd.NewStringName(context), gd.NewStringName(name)))
}

/*
Returns a specific slice (layer or mipmap) for a cached texture.
*/
func (self Instance) GetTextureSlice(context string, name string, layer int, mipmap int, layers int, mipmaps int) Resource.ID {
	return Resource.ID(class(self).GetTextureSlice(gd.NewStringName(context), gd.NewStringName(name), gd.Int(layer), gd.Int(mipmap), gd.Int(layers), gd.Int(mipmaps)))
}

/*
Returns a specific view of a slice (layer or mipmap) for a cached texture.
*/
func (self Instance) GetTextureSliceView(context string, name string, layer int, mipmap int, layers int, mipmaps int, view [1]gdclass.RDTextureView) Resource.ID {
	return Resource.ID(class(self).GetTextureSliceView(gd.NewStringName(context), gd.NewStringName(name), gd.Int(layer), gd.Int(mipmap), gd.Int(layers), gd.Int(mipmaps), view))
}

/*
Returns the texture size of a given slice of a cached texture.
*/
func (self Instance) GetTextureSliceSize(context string, name string, mipmap int) Vector2i.XY {
	return Vector2i.XY(class(self).GetTextureSliceSize(gd.NewStringName(context), gd.NewStringName(name), gd.Int(mipmap)))
}

/*
Frees all buffers related to this context.
*/
func (self Instance) ClearContext(context string) {
	class(self).ClearContext(gd.NewStringName(context))
}

/*
Returns the color texture we are rendering 3D content to. If multiview is used this will be a texture array with all views.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
func (self Instance) GetColorTexture() Resource.ID {
	return Resource.ID(class(self).GetColorTexture(false))
}

/*
Returns the specified layer from the color texture we are rendering 3D content to.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
func (self Instance) GetColorLayer(layer int) Resource.ID {
	return Resource.ID(class(self).GetColorLayer(gd.Int(layer), false))
}

/*
Returns the depth texture we are rendering 3D content to. If multiview is used this will be a texture array with all views.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
func (self Instance) GetDepthTexture() Resource.ID {
	return Resource.ID(class(self).GetDepthTexture(false))
}

/*
Returns the specified layer from the depth texture we are rendering 3D content to.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
func (self Instance) GetDepthLayer(layer int) Resource.ID {
	return Resource.ID(class(self).GetDepthLayer(gd.Int(layer), false))
}

/*
Returns the velocity texture we are rendering 3D content to. If multiview is used this will be a texture array with all views.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
func (self Instance) GetVelocityTexture() Resource.ID {
	return Resource.ID(class(self).GetVelocityTexture(false))
}

/*
Returns the specified layer from the velocity texture we are rendering 3D content to.
*/
func (self Instance) GetVelocityLayer(layer int) Resource.ID {
	return Resource.ID(class(self).GetVelocityLayer(gd.Int(layer), false))
}

/*
Returns the render target associated with this buffers object.
*/
func (self Instance) GetRenderTarget() Resource.ID {
	return Resource.ID(class(self).GetRenderTarget())
}

/*
Returns the view count for the associated viewport.
*/
func (self Instance) GetViewCount() int {
	return int(int(class(self).GetViewCount()))
}

/*
Returns the internal size of the render buffer (size before upscaling) with which textures are created by default.
*/
func (self Instance) GetInternalSize() Vector2i.XY {
	return Vector2i.XY(class(self).GetInternalSize())
}

/*
Returns the target size of the render buffer (size after upscaling).
*/
func (self Instance) GetTargetSize() Vector2i.XY {
	return Vector2i.XY(class(self).GetTargetSize())
}

/*
Returns the scaling mode used for upscaling.
*/
func (self Instance) GetScaling3dMode() gdclass.RenderingServerViewportScaling3DMode {
	return gdclass.RenderingServerViewportScaling3DMode(class(self).GetScaling3dMode())
}

/*
Returns the FSR sharpness value used while rendering the 3D content (if [method get_scaling_3d_mode] is an FSR mode).
*/
func (self Instance) GetFsrSharpness() Float.X {
	return Float.X(Float.X(class(self).GetFsrSharpness()))
}

/*
Returns the applied 3D MSAA mode for this viewport.
*/
func (self Instance) GetMsaa3d() gdclass.RenderingServerViewportMSAA {
	return gdclass.RenderingServerViewportMSAA(class(self).GetMsaa3d())
}

/*
Returns the number of MSAA samples used.
*/
func (self Instance) GetTextureSamples() gdclass.RenderingDeviceTextureSamples {
	return gdclass.RenderingDeviceTextureSamples(class(self).GetTextureSamples())
}

/*
Returns the screen-space antialiasing method applied.
*/
func (self Instance) GetScreenSpaceAa() gdclass.RenderingServerViewportScreenSpaceAA {
	return gdclass.RenderingServerViewportScreenSpaceAA(class(self).GetScreenSpaceAa())
}

/*
Returns [code]true[/code] if TAA is enabled.
*/
func (self Instance) GetUseTaa() bool {
	return bool(class(self).GetUseTaa())
}

/*
Returns [code]true[/code] if debanding is enabled.
*/
func (self Instance) GetUseDebanding() bool {
	return bool(class(self).GetUseDebanding())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RenderSceneBuffersRD

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RenderSceneBuffersRD"))
	return Instance{*(*gdclass.RenderSceneBuffersRD)(unsafe.Pointer(&object))}
}

/*
Returns [code]true[/code] if a cached texture exists for this name.
*/
//go:nosplit
func (self class) HasTexture(context gd.StringName, name gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(context))
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_has_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Create a new texture with the given definition and cache this under the given name. Will return the existing texture if it already exists.
*/
//go:nosplit
func (self class) CreateTexture(context gd.StringName, name gd.StringName, data_format gdclass.RenderingDeviceDataFormat, usage_bits gd.Int, texture_samples gdclass.RenderingDeviceTextureSamples, size gd.Vector2i, layers gd.Int, mipmaps gd.Int, unique bool) gd.RID {
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_create_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Create a new texture using the given format and view and cache this under the given name. Will return the existing texture if it already exists.
*/
//go:nosplit
func (self class) CreateTextureFromFormat(context gd.StringName, name gd.StringName, format [1]gdclass.RDTextureFormat, view [1]gdclass.RDTextureView, unique bool) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(context))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(format[0])[0])
	callframe.Arg(frame, pointers.Get(view[0])[0])
	callframe.Arg(frame, unique)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_create_texture_from_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Create a new texture view for an existing texture and cache this under the given view_name. Will return the existing teture view if it already exists. Will error if the source texture doesn't exist.
*/
//go:nosplit
func (self class) CreateTextureView(context gd.StringName, name gd.StringName, view_name gd.StringName, view [1]gdclass.RDTextureView) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(context))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(view_name))
	callframe.Arg(frame, pointers.Get(view[0])[0])
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_create_texture_view, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a cached texture with this name.
*/
//go:nosplit
func (self class) GetTexture(context gd.StringName, name gd.StringName) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(context))
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the texture format information with which a cached texture was created.
*/
//go:nosplit
func (self class) GetTextureFormat(context gd.StringName, name gd.StringName) [1]gdclass.RDTextureFormat {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(context))
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_texture_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.RDTextureFormat{gd.PointerWithOwnershipTransferredToGo[gdclass.RDTextureFormat](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns a specific slice (layer or mipmap) for a cached texture.
*/
//go:nosplit
func (self class) GetTextureSlice(context gd.StringName, name gd.StringName, layer gd.Int, mipmap gd.Int, layers gd.Int, mipmaps gd.Int) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(context))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, layer)
	callframe.Arg(frame, mipmap)
	callframe.Arg(frame, layers)
	callframe.Arg(frame, mipmaps)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_texture_slice, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a specific view of a slice (layer or mipmap) for a cached texture.
*/
//go:nosplit
func (self class) GetTextureSliceView(context gd.StringName, name gd.StringName, layer gd.Int, mipmap gd.Int, layers gd.Int, mipmaps gd.Int, view [1]gdclass.RDTextureView) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(context))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, layer)
	callframe.Arg(frame, mipmap)
	callframe.Arg(frame, layers)
	callframe.Arg(frame, mipmaps)
	callframe.Arg(frame, pointers.Get(view[0])[0])
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_texture_slice_view, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the texture size of a given slice of a cached texture.
*/
//go:nosplit
func (self class) GetTextureSliceSize(context gd.StringName, name gd.StringName, mipmap gd.Int) gd.Vector2i {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(context))
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, mipmap)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_texture_slice_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Frees all buffers related to this context.
*/
//go:nosplit
func (self class) ClearContext(context gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(context))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_clear_context, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the color texture we are rendering 3D content to. If multiview is used this will be a texture array with all views.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
//go:nosplit
func (self class) GetColorTexture(msaa bool) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, msaa)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_color_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the specified layer from the color texture we are rendering 3D content to.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
//go:nosplit
func (self class) GetColorLayer(layer gd.Int, msaa bool) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, msaa)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_color_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the depth texture we are rendering 3D content to. If multiview is used this will be a texture array with all views.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
//go:nosplit
func (self class) GetDepthTexture(msaa bool) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, msaa)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_depth_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the specified layer from the depth texture we are rendering 3D content to.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
//go:nosplit
func (self class) GetDepthLayer(layer gd.Int, msaa bool) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, msaa)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_depth_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the velocity texture we are rendering 3D content to. If multiview is used this will be a texture array with all views.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
//go:nosplit
func (self class) GetVelocityTexture(msaa bool) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, msaa)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_velocity_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the specified layer from the velocity texture we are rendering 3D content to.
*/
//go:nosplit
func (self class) GetVelocityLayer(layer gd.Int, msaa bool) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, msaa)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_velocity_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the render target associated with this buffers object.
*/
//go:nosplit
func (self class) GetRenderTarget() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_render_target, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the view count for the associated viewport.
*/
//go:nosplit
func (self class) GetViewCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_view_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the internal size of the render buffer (size before upscaling) with which textures are created by default.
*/
//go:nosplit
func (self class) GetInternalSize() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_internal_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the target size of the render buffer (size after upscaling).
*/
//go:nosplit
func (self class) GetTargetSize() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_target_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the scaling mode used for upscaling.
*/
//go:nosplit
func (self class) GetScaling3dMode() gdclass.RenderingServerViewportScaling3DMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingServerViewportScaling3DMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_scaling_3d_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the FSR sharpness value used while rendering the 3D content (if [method get_scaling_3d_mode] is an FSR mode).
*/
//go:nosplit
func (self class) GetFsrSharpness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_fsr_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the applied 3D MSAA mode for this viewport.
*/
//go:nosplit
func (self class) GetMsaa3d() gdclass.RenderingServerViewportMSAA {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingServerViewportMSAA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_msaa_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of MSAA samples used.
*/
//go:nosplit
func (self class) GetTextureSamples() gdclass.RenderingDeviceTextureSamples {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceTextureSamples](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_texture_samples, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the screen-space antialiasing method applied.
*/
//go:nosplit
func (self class) GetScreenSpaceAa() gdclass.RenderingServerViewportScreenSpaceAA {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingServerViewportScreenSpaceAA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_screen_space_aa, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if TAA is enabled.
*/
//go:nosplit
func (self class) GetUseTaa() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_use_taa, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if debanding is enabled.
*/
//go:nosplit
func (self class) GetUseDebanding() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersRD.Bind_get_use_debanding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRenderSceneBuffers(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRenderSceneBuffers(), name)
	}
}
func init() {
	gdclass.Register("RenderSceneBuffersRD", func(ptr gd.Object) any {
		return [1]gdclass.RenderSceneBuffersRD{*(*gdclass.RenderSceneBuffersRD)(unsafe.Pointer(&ptr))}
	})
}
