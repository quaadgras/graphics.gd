package RenderSceneBuffersRD

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/RenderSceneBuffers"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This object manages all 3D rendering buffers for the rendering device based renderers. An instance of this object is created for every viewport that has 3D rendering enabled.
All buffers are organized in [b]contexts[/b]. The default context is called [b]render_buffers[/b] and can contain amongst others the color buffer, depth buffer, velocity buffers, VRS density map and MSAA variants of these buffers.
Buffers are only guaranteed to exist during rendering of the viewport.
[b]Note:[/b] This is an internal rendering server object, do not instantiate this from script.

*/
type Go [1]classdb.RenderSceneBuffersRD

/*
Returns [code]true[/code] if a cached texture exists for this name.
*/
func (self Go) HasTexture(context string, name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasTexture(gc.StringName(context), gc.StringName(name)))
}

/*
Create a new texture with the given definition and cache this under the given name. Will return the existing texture if it already exists.
*/
func (self Go) CreateTexture(context string, name string, data_format classdb.RenderingDeviceDataFormat, usage_bits int, texture_samples classdb.RenderingDeviceTextureSamples, size gd.Vector2i, layers int, mipmaps int, unique bool) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).CreateTexture(gc.StringName(context), gc.StringName(name), data_format, gd.Int(usage_bits), texture_samples, size, gd.Int(layers), gd.Int(mipmaps), unique))
}

/*
Create a new texture using the given format and view and cache this under the given name. Will return the existing texture if it already exists.
*/
func (self Go) CreateTextureFromFormat(context string, name string, format gdclass.RDTextureFormat, view gdclass.RDTextureView, unique bool) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).CreateTextureFromFormat(gc.StringName(context), gc.StringName(name), format, view, unique))
}

/*
Create a new texture view for an existing texture and cache this under the given view_name. Will return the existing teture view if it already exists. Will error if the source texture doesn't exist.
*/
func (self Go) CreateTextureView(context string, name string, view_name string, view gdclass.RDTextureView) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).CreateTextureView(gc.StringName(context), gc.StringName(name), gc.StringName(view_name), view))
}

/*
Returns a cached texture with this name.
*/
func (self Go) GetTexture(context string, name string) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetTexture(gc.StringName(context), gc.StringName(name)))
}

/*
Returns the texture format information with which a cached texture was created.
*/
func (self Go) GetTextureFormat(context string, name string) gdclass.RDTextureFormat {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.RDTextureFormat(class(self).GetTextureFormat(gc, gc.StringName(context), gc.StringName(name)))
}

/*
Returns a specific slice (layer or mipmap) for a cached texture.
*/
func (self Go) GetTextureSlice(context string, name string, layer int, mipmap int, layers int, mipmaps int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetTextureSlice(gc.StringName(context), gc.StringName(name), gd.Int(layer), gd.Int(mipmap), gd.Int(layers), gd.Int(mipmaps)))
}

/*
Returns a specific view of a slice (layer or mipmap) for a cached texture.
*/
func (self Go) GetTextureSliceView(context string, name string, layer int, mipmap int, layers int, mipmaps int, view gdclass.RDTextureView) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetTextureSliceView(gc.StringName(context), gc.StringName(name), gd.Int(layer), gd.Int(mipmap), gd.Int(layers), gd.Int(mipmaps), view))
}

/*
Returns the texture size of a given slice of a cached texture.
*/
func (self Go) GetTextureSliceSize(context string, name string, mipmap int) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(class(self).GetTextureSliceSize(gc.StringName(context), gc.StringName(name), gd.Int(mipmap)))
}

/*
Frees all buffers related to this context.
*/
func (self Go) ClearContext(context string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearContext(gc.StringName(context))
}

/*
Returns the color texture we are rendering 3D content to. If multiview is used this will be a texture array with all views.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
func (self Go) GetColorTexture() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetColorTexture(false))
}

/*
Returns the specified layer from the color texture we are rendering 3D content to.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
func (self Go) GetColorLayer(layer int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetColorLayer(gd.Int(layer), false))
}

/*
Returns the depth texture we are rendering 3D content to. If multiview is used this will be a texture array with all views.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
func (self Go) GetDepthTexture() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetDepthTexture(false))
}

/*
Returns the specified layer from the depth texture we are rendering 3D content to.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
func (self Go) GetDepthLayer(layer int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetDepthLayer(gd.Int(layer), false))
}

/*
Returns the velocity texture we are rendering 3D content to. If multiview is used this will be a texture array with all views.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
func (self Go) GetVelocityTexture() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetVelocityTexture(false))
}

/*
Returns the specified layer from the velocity texture we are rendering 3D content to.
*/
func (self Go) GetVelocityLayer(layer int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetVelocityLayer(gd.Int(layer), false))
}

/*
Returns the render target associated with this buffers object.
*/
func (self Go) GetRenderTarget() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetRenderTarget())
}

/*
Returns the view count for the associated viewport.
*/
func (self Go) GetViewCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetViewCount()))
}

/*
Returns the internal size of the render buffer (size before upscaling) with which textures are created by default.
*/
func (self Go) GetInternalSize() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(class(self).GetInternalSize())
}

/*
Returns the target size of the render buffer (size after upscaling).
*/
func (self Go) GetTargetSize() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(class(self).GetTargetSize())
}

/*
Returns the scaling mode used for upscaling.
*/
func (self Go) GetScaling3dMode() classdb.RenderingServerViewportScaling3DMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingServerViewportScaling3DMode(class(self).GetScaling3dMode())
}

/*
Returns the FSR sharpness value used while rendering the 3D content (if [method get_scaling_3d_mode] is an FSR mode).
*/
func (self Go) GetFsrSharpness() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetFsrSharpness()))
}

/*
Returns the applied 3D MSAA mode for this viewport.
*/
func (self Go) GetMsaa3d() classdb.RenderingServerViewportMSAA {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingServerViewportMSAA(class(self).GetMsaa3d())
}

/*
Returns the number of MSAA samples used.
*/
func (self Go) GetTextureSamples() classdb.RenderingDeviceTextureSamples {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceTextureSamples(class(self).GetTextureSamples())
}

/*
Returns the screen-space antialiasing method applied.
*/
func (self Go) GetScreenSpaceAa() classdb.RenderingServerViewportScreenSpaceAA {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingServerViewportScreenSpaceAA(class(self).GetScreenSpaceAa())
}

/*
Returns [code]true[/code] if TAA is enabled.
*/
func (self Go) GetUseTaa() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetUseTaa())
}

/*
Returns [code]true[/code] if debanding is enabled.
*/
func (self Go) GetUseDebanding() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetUseDebanding())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RenderSceneBuffersRD
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RenderSceneBuffersRD"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns [code]true[/code] if a cached texture exists for this name.
*/
//go:nosplit
func (self class) HasTexture(context gd.StringName, name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(context))
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_has_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Create a new texture with the given definition and cache this under the given name. Will return the existing texture if it already exists.
*/
//go:nosplit
func (self class) CreateTexture(context gd.StringName, name gd.StringName, data_format classdb.RenderingDeviceDataFormat, usage_bits gd.Int, texture_samples classdb.RenderingDeviceTextureSamples, size gd.Vector2i, layers gd.Int, mipmaps gd.Int, unique bool) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(context))
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, data_format)
	callframe.Arg(frame, usage_bits)
	callframe.Arg(frame, texture_samples)
	callframe.Arg(frame, size)
	callframe.Arg(frame, layers)
	callframe.Arg(frame, mipmaps)
	callframe.Arg(frame, unique)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_create_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Create a new texture using the given format and view and cache this under the given name. Will return the existing texture if it already exists.
*/
//go:nosplit
func (self class) CreateTextureFromFormat(context gd.StringName, name gd.StringName, format gdclass.RDTextureFormat, view gdclass.RDTextureView, unique bool) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(context))
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(format[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(view[0].AsPointer())[0])
	callframe.Arg(frame, unique)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_create_texture_from_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Create a new texture view for an existing texture and cache this under the given view_name. Will return the existing teture view if it already exists. Will error if the source texture doesn't exist.
*/
//go:nosplit
func (self class) CreateTextureView(context gd.StringName, name gd.StringName, view_name gd.StringName, view gdclass.RDTextureView) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(context))
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(view_name))
	callframe.Arg(frame, mmm.Get(view[0].AsPointer())[0])
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_create_texture_view, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a cached texture with this name.
*/
//go:nosplit
func (self class) GetTexture(context gd.StringName, name gd.StringName) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(context))
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the texture format information with which a cached texture was created.
*/
//go:nosplit
func (self class) GetTextureFormat(ctx gd.Lifetime, context gd.StringName, name gd.StringName) gdclass.RDTextureFormat {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(context))
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_texture_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.RDTextureFormat
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns a specific slice (layer or mipmap) for a cached texture.
*/
//go:nosplit
func (self class) GetTextureSlice(context gd.StringName, name gd.StringName, layer gd.Int, mipmap gd.Int, layers gd.Int, mipmaps gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(context))
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, layer)
	callframe.Arg(frame, mipmap)
	callframe.Arg(frame, layers)
	callframe.Arg(frame, mipmaps)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_texture_slice, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a specific view of a slice (layer or mipmap) for a cached texture.
*/
//go:nosplit
func (self class) GetTextureSliceView(context gd.StringName, name gd.StringName, layer gd.Int, mipmap gd.Int, layers gd.Int, mipmaps gd.Int, view gdclass.RDTextureView) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(context))
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, layer)
	callframe.Arg(frame, mipmap)
	callframe.Arg(frame, layers)
	callframe.Arg(frame, mipmaps)
	callframe.Arg(frame, mmm.Get(view[0].AsPointer())[0])
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_texture_slice_view, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the texture size of a given slice of a cached texture.
*/
//go:nosplit
func (self class) GetTextureSliceSize(context gd.StringName, name gd.StringName, mipmap gd.Int) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(context))
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mipmap)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_texture_slice_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Frees all buffers related to this context.
*/
//go:nosplit
func (self class) ClearContext(context gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(context))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_clear_context, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the color texture we are rendering 3D content to. If multiview is used this will be a texture array with all views.
If [param msaa] is [b]true[/b] and MSAA is enabled, this returns the MSAA variant of the buffer.
*/
//go:nosplit
func (self class) GetColorTexture(msaa bool) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, msaa)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_color_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, msaa)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_color_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, msaa)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_depth_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, msaa)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_depth_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, msaa)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_velocity_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the specified layer from the velocity texture we are rendering 3D content to.
*/
//go:nosplit
func (self class) GetVelocityLayer(layer gd.Int, msaa bool) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, msaa)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_velocity_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the render target associated with this buffers object.
*/
//go:nosplit
func (self class) GetRenderTarget() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_render_target, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the view count for the associated viewport.
*/
//go:nosplit
func (self class) GetViewCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_view_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the internal size of the render buffer (size before upscaling) with which textures are created by default.
*/
//go:nosplit
func (self class) GetInternalSize() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_internal_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the target size of the render buffer (size after upscaling).
*/
//go:nosplit
func (self class) GetTargetSize() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_target_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the scaling mode used for upscaling.
*/
//go:nosplit
func (self class) GetScaling3dMode() classdb.RenderingServerViewportScaling3DMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingServerViewportScaling3DMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_scaling_3d_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the FSR sharpness value used while rendering the 3D content (if [method get_scaling_3d_mode] is an FSR mode).
*/
//go:nosplit
func (self class) GetFsrSharpness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_fsr_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the applied 3D MSAA mode for this viewport.
*/
//go:nosplit
func (self class) GetMsaa3d() classdb.RenderingServerViewportMSAA {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingServerViewportMSAA](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_msaa_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of MSAA samples used.
*/
//go:nosplit
func (self class) GetTextureSamples() classdb.RenderingDeviceTextureSamples {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceTextureSamples](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_texture_samples, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the screen-space antialiasing method applied.
*/
//go:nosplit
func (self class) GetScreenSpaceAa() classdb.RenderingServerViewportScreenSpaceAA {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingServerViewportScreenSpaceAA](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_screen_space_aa, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if TAA is enabled.
*/
//go:nosplit
func (self class) GetUseTaa() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_use_taa, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if debanding is enabled.
*/
//go:nosplit
func (self class) GetUseDebanding() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderSceneBuffersRD.Bind_get_use_debanding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRenderSceneBuffersRD() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRenderSceneBuffersRD() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRenderSceneBuffers() RenderSceneBuffers.GD { return *((*RenderSceneBuffers.GD)(unsafe.Pointer(&self))) }
func (self Go) AsRenderSceneBuffers() RenderSceneBuffers.Go { return *((*RenderSceneBuffers.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRenderSceneBuffers(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRenderSceneBuffers(), name)
	}
}
func init() {classdb.Register("RenderSceneBuffersRD", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
