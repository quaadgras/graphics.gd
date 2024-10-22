package RenderingDevice

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[RenderingDevice] is an abstraction for working with modern low-level graphics APIs such as Vulkan. Compared to [RenderingServer] (which works with Godot's own rendering subsystems), [RenderingDevice] is much lower-level and allows working more directly with the underlying graphics APIs. [RenderingDevice] is used in Godot to provide support for several modern low-level graphics APIs while reducing the amount of code duplication required. [RenderingDevice] can also be used in your own projects to perform things that are not exposed by [RenderingServer] or high-level nodes, such as using compute shaders.
On startup, Godot creates a global [RenderingDevice] which can be retrieved using [method RenderingServer.get_rendering_device]. This global [RenderingDevice] performs drawing to the screen.
[b]Local RenderingDevices:[/b] Using [method RenderingServer.create_local_rendering_device], you can create "secondary" rendering devices to perform drawing and GPU compute operations on separate threads.
[b]Note:[/b] [RenderingDevice] assumes intermediate knowledge of modern graphics APIs such as Vulkan, Direct3D 12, Metal or WebGPU. These graphics APIs are lower-level than OpenGL or Direct3D 11, requiring you to perform what was previously done by the graphics driver itself. If you have difficulty understanding the concepts used in this class, follow the [url=https://vulkan-tutorial.com/]Vulkan Tutorial[/url] or [url=https://vkguide.dev/]Vulkan Guide[/url]. It's recommended to have existing modern OpenGL or Direct3D 11 knowledge before attempting to learn a low-level graphics API.
[b]Note:[/b] [RenderingDevice] is not available when running in headless mode or when using the Compatibility rendering method.

*/
type Go [1]classdb.RenderingDevice

/*
Creates a new texture. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
[b]Note:[/b] Not to be confused with [method RenderingServer.texture_2d_create], which creates the Godot-specific [Texture2D] resource as opposed to the graphics API's own texture type.
*/
func (self Go) TextureCreate(format gdclass.RDTextureFormat, view gdclass.RDTextureView) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).TextureCreate(format, view, ([1]gd.ArrayOf[gd.PackedByteArray]{}[0])))
}

/*
Creates a shared texture using the specified [param view] and the texture information from [param with_texture].
*/
func (self Go) TextureCreateShared(view gdclass.RDTextureView, with_texture gd.RID) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).TextureCreateShared(view, with_texture))
}

/*
Creates a shared texture using the specified [param view] and the texture information from [param with_texture]'s [param layer] and [param mipmap]. The number of included mipmaps from the original texture can be controlled using the [param mipmaps] parameter. Only relevant for textures with multiple layers, such as 3D textures, texture arrays and cubemaps. For single-layer textures, use [method texture_create_shared]
For 2D textures (which only have one layer), [param layer] must be [code]0[/code].
[b]Note:[/b] Layer slicing is only supported for 2D texture arrays, not 3D textures or cubemaps.
*/
func (self Go) TextureCreateSharedFromSlice(view gdclass.RDTextureView, with_texture gd.RID, layer int, mipmap int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).TextureCreateSharedFromSlice(view, with_texture, gd.Int(layer), gd.Int(mipmap), gd.Int(1), 0))
}

/*
Returns an RID for an existing [param image] ([code]VkImage[/code]) with the given [param type], [param format], [param samples], [param usage_flags], [param width], [param height], [param depth], and [param layers]. This can be used to allow Godot to render onto foreign images.
*/
func (self Go) TextureCreateFromExtension(atype classdb.RenderingDeviceTextureType, format classdb.RenderingDeviceDataFormat, samples classdb.RenderingDeviceTextureSamples, usage_flags classdb.RenderingDeviceTextureUsageBits, image int, width int, height int, depth int, layers int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).TextureCreateFromExtension(atype, format, samples, usage_flags, gd.Int(image), gd.Int(width), gd.Int(height), gd.Int(depth), gd.Int(layers)))
}

/*
Updates texture data with new data, replacing the previous data in place. The updated texture data must have the same dimensions and format. For 2D textures (which only have one layer), [param layer] must be [code]0[/code]. Returns [constant @GlobalScope.OK] if the update was successful, [constant @GlobalScope.ERR_INVALID_PARAMETER] otherwise.
[b]Note:[/b] Updating textures is forbidden during creation of a draw or compute list.
[b]Note:[/b] The existing [param texture] can't be updated while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to update this texture.
[b]Note:[/b] The existing [param texture] requires the [constant TEXTURE_USAGE_CAN_UPDATE_BIT] to be updatable.
*/
func (self Go) TextureUpdate(texture gd.RID, layer int, data []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).TextureUpdate(texture, gd.Int(layer), gc.PackedByteSlice(data)))
}

/*
Returns the [param texture] data for the specified [param layer] as raw binary data. For 2D textures (which only have one layer), [param layer] must be [code]0[/code].
[b]Note:[/b] [param texture] can't be retrieved while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to retrieve this texture. Otherwise, an error is printed and a empty [PackedByteArray] is returned.
[b]Note:[/b] [param texture] requires the [constant TEXTURE_USAGE_CAN_COPY_FROM_BIT] to be retrieved. Otherwise, an error is printed and a empty [PackedByteArray] is returned.
*/
func (self Go) TextureGetData(texture gd.RID, layer int) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(class(self).TextureGetData(gc, texture, gd.Int(layer)).Bytes())
}

/*
Returns [code]true[/code] if the specified [param format] is supported for the given [param usage_flags], [code]false[/code] otherwise.
*/
func (self Go) TextureIsFormatSupportedForUsage(format classdb.RenderingDeviceDataFormat, usage_flags classdb.RenderingDeviceTextureUsageBits) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).TextureIsFormatSupportedForUsage(format, usage_flags))
}

/*
Returns [code]true[/code] if the [param texture] is shared, [code]false[/code] otherwise. See [RDTextureView].
*/
func (self Go) TextureIsShared(texture gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).TextureIsShared(texture))
}

/*
Returns [code]true[/code] if the [param texture] is valid, [code]false[/code] otherwise.
*/
func (self Go) TextureIsValid(texture gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).TextureIsValid(texture))
}

/*
Copies the [param from_texture] to [param to_texture] with the specified [param from_pos], [param to_pos] and [param size] coordinates. The Z axis of the [param from_pos], [param to_pos] and [param size] must be [code]0[/code] for 2-dimensional textures. Source and destination mipmaps/layers must also be specified, with these parameters being [code]0[/code] for textures without mipmaps or single-layer textures. Returns [constant @GlobalScope.OK] if the texture copy was successful or [constant @GlobalScope.ERR_INVALID_PARAMETER] otherwise.
[b]Note:[/b] [param from_texture] texture can't be copied while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to copy this texture.
[b]Note:[/b] [param from_texture] texture requires the [constant TEXTURE_USAGE_CAN_COPY_FROM_BIT] to be retrieved.
[b]Note:[/b] [param to_texture] can't be copied while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to copy this texture.
[b]Note:[/b] [param to_texture] requires the [constant TEXTURE_USAGE_CAN_COPY_TO_BIT] to be retrieved.
[b]Note:[/b] [param from_texture] and [param to_texture] must be of the same type (color or depth).
*/
func (self Go) TextureCopy(from_texture gd.RID, to_texture gd.RID, from_pos gd.Vector3, to_pos gd.Vector3, size gd.Vector3, src_mipmap int, dst_mipmap int, src_layer int, dst_layer int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).TextureCopy(from_texture, to_texture, from_pos, to_pos, size, gd.Int(src_mipmap), gd.Int(dst_mipmap), gd.Int(src_layer), gd.Int(dst_layer)))
}

/*
Clears the specified [param texture] by replacing all of its pixels with the specified [param color]. [param base_mipmap] and [param mipmap_count] determine which mipmaps of the texture are affected by this clear operation, while [param base_layer] and [param layer_count] determine which layers of a 3D texture (or texture array) are affected by this clear operation. For 2D textures (which only have one layer by design), [param base_layer] must be [code]0[/code] and [param layer_count] must be [code]1[/code].
[b]Note:[/b] [param texture] can't be cleared while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to clear this texture.
*/
func (self Go) TextureClear(texture gd.RID, color gd.Color, base_mipmap int, mipmap_count int, base_layer int, layer_count int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).TextureClear(texture, color, gd.Int(base_mipmap), gd.Int(mipmap_count), gd.Int(base_layer), gd.Int(layer_count)))
}

/*
Resolves the [param from_texture] texture onto [param to_texture] with multisample antialiasing enabled. This must be used when rendering a framebuffer for MSAA to work. Returns [constant @GlobalScope.OK] if successful, [constant @GlobalScope.ERR_INVALID_PARAMETER] otherwise.
[b]Note:[/b] [param from_texture] and [param to_texture] textures must have the same dimension, format and type (color or depth).
[b]Note:[/b] [param from_texture] can't be copied while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to resolve this texture.
[b]Note:[/b] [param from_texture] requires the [constant TEXTURE_USAGE_CAN_COPY_FROM_BIT] to be retrieved.
[b]Note:[/b] [param from_texture] must be multisampled and must also be 2D (or a slice of a 3D/cubemap texture).
[b]Note:[/b] [param to_texture] can't be copied while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to resolve this texture.
[b]Note:[/b] [param to_texture] texture requires the [constant TEXTURE_USAGE_CAN_COPY_TO_BIT] to be retrieved.
[b]Note:[/b] [param to_texture] texture must [b]not[/b] be multisampled and must also be 2D (or a slice of a 3D/cubemap texture).
*/
func (self Go) TextureResolveMultisample(from_texture gd.RID, to_texture gd.RID) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).TextureResolveMultisample(from_texture, to_texture))
}

/*
Returns the data format used to create this texture.
*/
func (self Go) TextureGetFormat(texture gd.RID) gdclass.RDTextureFormat {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.RDTextureFormat(class(self).TextureGetFormat(gc, texture))
}

/*
Returns the internal graphics handle for this texture object. For use when communicating with third-party APIs mostly with GDExtension.
[b]Note:[/b] This function returns a [code]uint64_t[/code] which internally maps to a [code]GLuint[/code] (OpenGL) or [code]VkImage[/code] (Vulkan).
*/
func (self Go) TextureGetNativeHandle(texture gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).TextureGetNativeHandle(texture)))
}

/*
Creates a new framebuffer format with the specified [param attachments] and [param view_count]. Returns the new framebuffer's unique framebuffer format ID.
If [param view_count] is greater than or equal to [code]2[/code], enables multiview which is used for VR rendering. This requires support for the Vulkan multiview extension.
*/
func (self Go) FramebufferFormatCreate(attachments gd.ArrayOf[gdclass.RDAttachmentFormat]) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).FramebufferFormatCreate(attachments, gd.Int(1))))
}

/*
Creates a multipass framebuffer format with the specified [param attachments], [param passes] and [param view_count] and returns its ID. If [param view_count] is greater than or equal to [code]2[/code], enables multiview which is used for VR rendering. This requires support for the Vulkan multiview extension.
*/
func (self Go) FramebufferFormatCreateMultipass(attachments gd.ArrayOf[gdclass.RDAttachmentFormat], passes gd.ArrayOf[gdclass.RDFramebufferPass]) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).FramebufferFormatCreateMultipass(attachments, passes, gd.Int(1))))
}

/*
Creates a new empty framebuffer format with the specified number of [param samples] and returns its ID.
*/
func (self Go) FramebufferFormatCreateEmpty() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).FramebufferFormatCreateEmpty(0)))
}

/*
Returns the number of texture samples used for the given framebuffer [param format] ID (returned by [method framebuffer_get_format]).
*/
func (self Go) FramebufferFormatGetTextureSamples(format int) classdb.RenderingDeviceTextureSamples {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceTextureSamples(class(self).FramebufferFormatGetTextureSamples(gd.Int(format), gd.Int(0)))
}

/*
Creates a new framebuffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Go) FramebufferCreate(textures gd.ArrayOf[gd.RID]) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).FramebufferCreate(textures, gd.Int(-1), gd.Int(1)))
}

/*
Creates a new multipass framebuffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Go) FramebufferCreateMultipass(textures gd.ArrayOf[gd.RID], passes gd.ArrayOf[gdclass.RDFramebufferPass]) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).FramebufferCreateMultipass(textures, passes, gd.Int(-1), gd.Int(1)))
}

/*
Creates a new empty framebuffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Go) FramebufferCreateEmpty(size gd.Vector2i) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).FramebufferCreateEmpty(size, 0, gd.Int(-1)))
}

/*
Returns the format ID of the framebuffer specified by the [param framebuffer] RID. This ID is guaranteed to be unique for the same formats and does not need to be freed.
*/
func (self Go) FramebufferGetFormat(framebuffer gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).FramebufferGetFormat(framebuffer)))
}

/*
Returns [code]true[/code] if the framebuffer specified by the [param framebuffer] RID is valid, [code]false[/code] otherwise.
*/
func (self Go) FramebufferIsValid(framebuffer gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).FramebufferIsValid(framebuffer))
}

/*
Creates a new sampler. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Go) SamplerCreate(state gdclass.RDSamplerState) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).SamplerCreate(state))
}

/*
Returns [code]true[/code] if implementation supports using a texture of [param format] with the given [param sampler_filter].
*/
func (self Go) SamplerIsFormatSupportedForFilter(format classdb.RenderingDeviceDataFormat, sampler_filter classdb.RenderingDeviceSamplerFilter) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).SamplerIsFormatSupportedForFilter(format, sampler_filter))
}

/*
It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Go) VertexBufferCreate(size_bytes int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).VertexBufferCreate(gd.Int(size_bytes), gc.PackedByteSlice(([1][]byte{}[0])), false))
}

/*
Creates a new vertex format with the specified [param vertex_descriptions]. Returns a unique vertex format ID corresponding to the newly created vertex format.
*/
func (self Go) VertexFormatCreate(vertex_descriptions gd.ArrayOf[gdclass.RDVertexAttribute]) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).VertexFormatCreate(vertex_descriptions)))
}

/*
Creates a vertex array based on the specified buffers. Optionally, [param offsets] (in bytes) may be defined for each buffer.
*/
func (self Go) VertexArrayCreate(vertex_count int, vertex_format int, src_buffers gd.ArrayOf[gd.RID]) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).VertexArrayCreate(gd.Int(vertex_count), gd.Int(vertex_format), src_buffers, gc.PackedInt64Slice(([1][]int64{}[0]))))
}

/*
Creates a new index buffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Go) IndexBufferCreate(size_indices int, format classdb.RenderingDeviceIndexBufferFormat) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).IndexBufferCreate(gd.Int(size_indices), format, gc.PackedByteSlice(([1][]byte{}[0])), false))
}

/*
Creates a new index array. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Go) IndexArrayCreate(index_buffer gd.RID, index_offset int, index_count int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).IndexArrayCreate(index_buffer, gd.Int(index_offset), gd.Int(index_count)))
}

/*
Compiles a SPIR-V from the shader source code in [param shader_source] and returns the SPIR-V as a [RDShaderSPIRV]. This intermediate language shader is portable across different GPU models and driver versions, but cannot be run directly by GPUs until compiled into a binary shader using [method shader_compile_binary_from_spirv].
If [param allow_cache] is [code]true[/code], make use of the shader cache generated by Godot. This avoids a potentially lengthy shader compilation step if the shader is already in cache. If [param allow_cache] is [code]false[/code], Godot's shader cache is ignored and the shader will always be recompiled.
*/
func (self Go) ShaderCompileSpirvFromSource(shader_source gdclass.RDShaderSource) gdclass.RDShaderSPIRV {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.RDShaderSPIRV(class(self).ShaderCompileSpirvFromSource(gc, shader_source, true))
}

/*
Compiles a binary shader from [param spirv_data] and returns the compiled binary data as a [PackedByteArray]. This compiled shader is specific to the GPU model and driver version used; it will not work on different GPU models or even different driver versions. See also [method shader_compile_spirv_from_source].
[param name] is an optional human-readable name that can be given to the compiled shader for organizational purposes.
*/
func (self Go) ShaderCompileBinaryFromSpirv(spirv_data gdclass.RDShaderSPIRV) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(class(self).ShaderCompileBinaryFromSpirv(gc, spirv_data, gc.String("")).Bytes())
}

/*
Creates a new shader instance from SPIR-V intermediate code. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method. See also [method shader_compile_spirv_from_source] and [method shader_create_from_bytecode].
*/
func (self Go) ShaderCreateFromSpirv(spirv_data gdclass.RDShaderSPIRV) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).ShaderCreateFromSpirv(spirv_data, gc.String("")))
}

/*
Creates a new shader instance from a binary compiled shader. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method. See also [method shader_compile_binary_from_spirv] and [method shader_create_from_spirv].
*/
func (self Go) ShaderCreateFromBytecode(binary_data []byte) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).ShaderCreateFromBytecode(gc.PackedByteSlice(binary_data), ([1]gd.RID{}[0])))
}

/*
Create a placeholder RID by allocating an RID without initializing it for use in [method shader_create_from_bytecode]. This allows you to create an RID for a shader and pass it around, but defer compiling the shader to a later time.
*/
func (self Go) ShaderCreatePlaceholder() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).ShaderCreatePlaceholder())
}

/*
Returns the internal vertex input mask. Internally, the vertex input mask is an unsigned integer consisting of the locations (specified in GLSL via. [code]layout(location = ...)[/code]) of the input variables (specified in GLSL by the [code]in[/code] keyword).
*/
func (self Go) ShaderGetVertexInputAttributeMask(shader gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ShaderGetVertexInputAttributeMask(shader)))
}

/*
Creates a new uniform buffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Go) UniformBufferCreate(size_bytes int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).UniformBufferCreate(gd.Int(size_bytes), gc.PackedByteSlice(([1][]byte{}[0]))))
}

/*
Creates a [url=https://vkguide.dev/docs/chapter-4/storage_buffers/]storage buffer[/url] with the specified [param data] and [param usage]. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Go) StorageBufferCreate(size_bytes int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).StorageBufferCreate(gd.Int(size_bytes), gc.PackedByteSlice(([1][]byte{}[0])), 0))
}

/*
Creates a new texture buffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Go) TextureBufferCreate(size_bytes int, format classdb.RenderingDeviceDataFormat) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).TextureBufferCreate(gd.Int(size_bytes), format, gc.PackedByteSlice(([1][]byte{}[0]))))
}

/*
Creates a new uniform set. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Go) UniformSetCreate(uniforms gd.ArrayOf[gdclass.RDUniform], shader gd.RID, shader_set int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).UniformSetCreate(uniforms, shader, gd.Int(shader_set)))
}

/*
Checks if the [param uniform_set] is valid, i.e. is owned.
*/
func (self Go) UniformSetIsValid(uniform_set gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).UniformSetIsValid(uniform_set))
}

/*
Copies [param size] bytes from the [param src_buffer] at [param src_offset] into [param dst_buffer] at [param dst_offset].
Prints an error if:
- [param size] exceeds the size of either [param src_buffer] or [param dst_buffer] at their corresponding offsets
- a draw list is currently active (created by [method draw_list_begin])
- a compute list is currently active (created by [method compute_list_begin])
*/
func (self Go) BufferCopy(src_buffer gd.RID, dst_buffer gd.RID, src_offset int, dst_offset int, size int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).BufferCopy(src_buffer, dst_buffer, gd.Int(src_offset), gd.Int(dst_offset), gd.Int(size)))
}

/*
Updates a region of [param size_bytes] bytes, starting at [param offset], in the buffer, with the specified [param data].
Prints an error if:
- the region specified by [param offset] + [param size_bytes] exceeds the buffer
- a draw list is currently active (created by [method draw_list_begin])
- a compute list is currently active (created by [method compute_list_begin])
*/
func (self Go) BufferUpdate(buffer gd.RID, offset int, size_bytes int, data []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).BufferUpdate(buffer, gd.Int(offset), gd.Int(size_bytes), gc.PackedByteSlice(data)))
}

/*
Clears the contents of the [param buffer], clearing [param size_bytes] bytes, starting at [param offset].
Prints an error if:
- the size isn't a multiple of four
- the region specified by [param offset] + [param size_bytes] exceeds the buffer
- a draw list is currently active (created by [method draw_list_begin])
- a compute list is currently active (created by [method compute_list_begin])
*/
func (self Go) BufferClear(buffer gd.RID, offset int, size_bytes int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).BufferClear(buffer, gd.Int(offset), gd.Int(size_bytes)))
}

/*
Returns a copy of the data of the specified [param buffer], optionally [param offset_bytes] and [param size_bytes] can be set to copy only a portion of the buffer.
*/
func (self Go) BufferGetData(buffer gd.RID) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(class(self).BufferGetData(gc, buffer, gd.Int(0), gd.Int(0)).Bytes())
}

/*
Creates a new render pipeline. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Go) RenderPipelineCreate(shader gd.RID, framebuffer_format int, vertex_format int, primitive classdb.RenderingDeviceRenderPrimitive, rasterization_state gdclass.RDPipelineRasterizationState, multisample_state gdclass.RDPipelineMultisampleState, stencil_state gdclass.RDPipelineDepthStencilState, color_blend_state gdclass.RDPipelineColorBlendState) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).RenderPipelineCreate(shader, gd.Int(framebuffer_format), gd.Int(vertex_format), primitive, rasterization_state, multisample_state, stencil_state, color_blend_state, 0, gd.Int(0), ([1]gd.ArrayOf[gdclass.RDPipelineSpecializationConstant]{}[0])))
}

/*
Returns [code]true[/code] if the render pipeline specified by the [param render_pipeline] RID is valid, [code]false[/code] otherwise.
*/
func (self Go) RenderPipelineIsValid(render_pipeline gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).RenderPipelineIsValid(render_pipeline))
}

/*
Creates a new compute pipeline. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
func (self Go) ComputePipelineCreate(shader gd.RID) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).ComputePipelineCreate(shader, ([1]gd.ArrayOf[gdclass.RDPipelineSpecializationConstant]{}[0])))
}

/*
Returns [code]true[/code] if the compute pipeline specified by the [param compute_pipeline] RID is valid, [code]false[/code] otherwise.
*/
func (self Go) ComputePipelineIsValid(compute_pipeline gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).ComputePipelineIsValid(compute_pipeline))
}

/*
Returns the window width matching the graphics API context for the given window ID (in pixels). Despite the parameter being named [param screen], this returns the [i]window[/i] size. See also [method screen_get_height].
[b]Note:[/b] Only the main [RenderingDevice] returned by [method RenderingServer.get_rendering_device] has a width. If called on a local [RenderingDevice], this method prints an error and returns [constant INVALID_ID].
*/
func (self Go) ScreenGetWidth() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ScreenGetWidth(gd.Int(0))))
}

/*
Returns the window height matching the graphics API context for the given window ID (in pixels). Despite the parameter being named [param screen], this returns the [i]window[/i] size. See also [method screen_get_width].
[b]Note:[/b] Only the main [RenderingDevice] returned by [method RenderingServer.get_rendering_device] has a height. If called on a local [RenderingDevice], this method prints an error and returns [constant INVALID_ID].
*/
func (self Go) ScreenGetHeight() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ScreenGetHeight(gd.Int(0))))
}

/*
Returns the framebuffer format of the given screen.
[b]Note:[/b] Only the main [RenderingDevice] returned by [method RenderingServer.get_rendering_device] has a format. If called on a local [RenderingDevice], this method prints an error and returns [constant INVALID_ID].
*/
func (self Go) ScreenGetFramebufferFormat() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ScreenGetFramebufferFormat(gd.Int(0))))
}

/*
High-level variant of [method draw_list_begin], with the parameters automatically being adjusted for drawing onto the window specified by the [param screen] ID.
[b]Note:[/b] Cannot be used with local RenderingDevices, as these don't have a screen. If called on a local RenderingDevice, [method draw_list_begin_for_screen] returns [constant INVALID_ID].
*/
func (self Go) DrawListBeginForScreen() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).DrawListBeginForScreen(gd.Int(0), gd.Color{0, 0, 0, 1})))
}

/*
Starts a list of raster drawing commands created with the [code]draw_*[/code] methods. The returned value should be passed to other [code]draw_list_*[/code] functions.
Multiple draw lists cannot be created at the same time; you must finish the previous draw list first using [method draw_list_end].
A simple drawing operation might look like this (code is not a complete example):
[codeblock]
var rd = RenderingDevice.new()
var clear_colors = PackedColorArray([Color(0, 0, 0, 0), Color(0, 0, 0, 0), Color(0, 0, 0, 0)])
var draw_list = rd.draw_list_begin(framebuffers[i], RenderingDevice.INITIAL_ACTION_CLEAR, RenderingDevice.FINAL_ACTION_READ, RenderingDevice.INITIAL_ACTION_CLEAR, RenderingDevice.FINAL_ACTION_DISCARD, clear_colors)

# Draw opaque.
rd.draw_list_bind_render_pipeline(draw_list, raster_pipeline)
rd.draw_list_bind_uniform_set(draw_list, raster_base_uniform, 0)
rd.draw_list_set_push_constant(draw_list, raster_push_constant, raster_push_constant.size())
rd.draw_list_draw(draw_list, false, 1, slice_triangle_count[i] * 3)
# Draw wire.
rd.draw_list_bind_render_pipeline(draw_list, raster_pipeline_wire)
rd.draw_list_bind_uniform_set(draw_list, raster_base_uniform, 0)
rd.draw_list_set_push_constant(draw_list, raster_push_constant, raster_push_constant.size())
rd.draw_list_draw(draw_list, false, 1, slice_triangle_count[i] * 3)

rd.draw_list_end()
[/codeblock]
*/
func (self Go) DrawListBegin(framebuffer gd.RID, initial_color_action classdb.RenderingDeviceInitialAction, final_color_action classdb.RenderingDeviceFinalAction, initial_depth_action classdb.RenderingDeviceInitialAction, final_depth_action classdb.RenderingDeviceFinalAction) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).DrawListBegin(framebuffer, initial_color_action, final_color_action, initial_depth_action, final_depth_action, gc.PackedColorSlice(([1][]gd.Color{}[0])), gd.Float(1.0), gd.Int(0), gd.NewRect2(0, 0, 0, 0))))
}

/*
This method does nothing and always returns an empty [PackedInt64Array].
*/
func (self Go) DrawListBeginSplit(framebuffer gd.RID, splits int, initial_color_action classdb.RenderingDeviceInitialAction, final_color_action classdb.RenderingDeviceFinalAction, initial_depth_action classdb.RenderingDeviceInitialAction, final_depth_action classdb.RenderingDeviceFinalAction) []int64 {
	gc := gd.GarbageCollector(); _ = gc
	return []int64(class(self).DrawListBeginSplit(gc, framebuffer, gd.Int(splits), initial_color_action, final_color_action, initial_depth_action, final_depth_action, gc.PackedColorSlice(([1][]gd.Color{}[0])), gd.Float(1.0), gd.Int(0), gd.NewRect2(0, 0, 0, 0), ([1]gd.ArrayOf[gd.RID]{}[0])).AsSlice())
}

/*
Sets blend constants for the specified [param draw_list] to [param color]. Blend constants are used only if the graphics pipeline is created with [constant DYNAMIC_STATE_BLEND_CONSTANTS] flag set.
*/
func (self Go) DrawListSetBlendConstants(draw_list int, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawListSetBlendConstants(gd.Int(draw_list), color)
}

/*
Binds [param render_pipeline] to the specified [param draw_list].
*/
func (self Go) DrawListBindRenderPipeline(draw_list int, render_pipeline gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawListBindRenderPipeline(gd.Int(draw_list), render_pipeline)
}

/*
Binds [param uniform_set] to the specified [param draw_list]. A [param set_index] must also be specified, which is an identifier starting from [code]0[/code] that must match the one expected by the draw list.
*/
func (self Go) DrawListBindUniformSet(draw_list int, uniform_set gd.RID, set_index int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawListBindUniformSet(gd.Int(draw_list), uniform_set, gd.Int(set_index))
}

/*
Binds [param vertex_array] to the specified [param draw_list].
*/
func (self Go) DrawListBindVertexArray(draw_list int, vertex_array gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawListBindVertexArray(gd.Int(draw_list), vertex_array)
}

/*
Binds [param index_array] to the specified [param draw_list].
*/
func (self Go) DrawListBindIndexArray(draw_list int, index_array gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawListBindIndexArray(gd.Int(draw_list), index_array)
}

/*
Sets the push constant data to [param buffer] for the specified [param draw_list]. The shader determines how this binary data is used. The buffer's size in bytes must also be specified in [param size_bytes] (this can be obtained by calling the [method PackedByteArray.size] method on the passed [param buffer]).
*/
func (self Go) DrawListSetPushConstant(draw_list int, buffer []byte, size_bytes int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawListSetPushConstant(gd.Int(draw_list), gc.PackedByteSlice(buffer), gd.Int(size_bytes))
}

/*
Submits [param draw_list] for rendering on the GPU. This is the raster equivalent to [method compute_list_dispatch].
*/
func (self Go) DrawListDraw(draw_list int, use_indices bool, instances int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawListDraw(gd.Int(draw_list), use_indices, gd.Int(instances), gd.Int(0))
}

/*
Creates a scissor rectangle and enables it for the specified [param draw_list]. Scissor rectangles are used for clipping by discarding fragments that fall outside a specified rectangular portion of the screen. See also [method draw_list_disable_scissor].
[b]Note:[/b] The specified [param rect] is automatically intersected with the screen's dimensions, which means it cannot exceed the screen's dimensions.
*/
func (self Go) DrawListEnableScissor(draw_list int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawListEnableScissor(gd.Int(draw_list), gd.NewRect2(0, 0, 0, 0))
}

/*
Removes and disables the scissor rectangle for the specified [param draw_list]. See also [method draw_list_enable_scissor].
*/
func (self Go) DrawListDisableScissor(draw_list int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawListDisableScissor(gd.Int(draw_list))
}

/*
Switches to the next draw pass.
*/
func (self Go) DrawListSwitchToNextPass() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).DrawListSwitchToNextPass()))
}

/*
This method does nothing and always returns an empty [PackedInt64Array].
*/
func (self Go) DrawListSwitchToNextPassSplit(splits int) []int64 {
	gc := gd.GarbageCollector(); _ = gc
	return []int64(class(self).DrawListSwitchToNextPassSplit(gc, gd.Int(splits)).AsSlice())
}

/*
Finishes a list of raster drawing commands created with the [code]draw_*[/code] methods.
*/
func (self Go) DrawListEnd() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawListEnd()
}

/*
Starts a list of compute commands created with the [code]compute_*[/code] methods. The returned value should be passed to other [code]compute_list_*[/code] functions.
Multiple compute lists cannot be created at the same time; you must finish the previous compute list first using [method compute_list_end].
A simple compute operation might look like this (code is not a complete example):
[codeblock]
var rd = RenderingDevice.new()
var compute_list = rd.compute_list_begin()

rd.compute_list_bind_compute_pipeline(compute_list, compute_shader_dilate_pipeline)
rd.compute_list_bind_uniform_set(compute_list, compute_base_uniform_set, 0)
rd.compute_list_bind_uniform_set(compute_list, dilate_uniform_set, 1)

for i in atlas_slices:
    rd.compute_list_set_push_constant(compute_list, push_constant, push_constant.size())
    rd.compute_list_dispatch(compute_list, group_size.x, group_size.y, group_size.z)
    # No barrier, let them run all together.

rd.compute_list_end()
[/codeblock]
*/
func (self Go) ComputeListBegin() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ComputeListBegin()))
}

/*
Tells the GPU what compute pipeline to use when processing the compute list. If the shader has changed since the last time this function was called, Godot will unbind all descriptor sets and will re-bind them inside [method compute_list_dispatch].
*/
func (self Go) ComputeListBindComputePipeline(compute_list int, compute_pipeline gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ComputeListBindComputePipeline(gd.Int(compute_list), compute_pipeline)
}

/*
Sets the push constant data to [param buffer] for the specified [param compute_list]. The shader determines how this binary data is used. The buffer's size in bytes must also be specified in [param size_bytes] (this can be obtained by calling the [method PackedByteArray.size] method on the passed [param buffer]).
*/
func (self Go) ComputeListSetPushConstant(compute_list int, buffer []byte, size_bytes int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ComputeListSetPushConstant(gd.Int(compute_list), gc.PackedByteSlice(buffer), gd.Int(size_bytes))
}

/*
Binds the [param uniform_set] to this [param compute_list]. Godot ensures that all textures in the uniform set have the correct Vulkan access masks. If Godot had to change access masks of textures, it will raise a Vulkan image memory barrier.
*/
func (self Go) ComputeListBindUniformSet(compute_list int, uniform_set gd.RID, set_index int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ComputeListBindUniformSet(gd.Int(compute_list), uniform_set, gd.Int(set_index))
}

/*
Submits the compute list for processing on the GPU. This is the compute equivalent to [method draw_list_draw].
*/
func (self Go) ComputeListDispatch(compute_list int, x_groups int, y_groups int, z_groups int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ComputeListDispatch(gd.Int(compute_list), gd.Int(x_groups), gd.Int(y_groups), gd.Int(z_groups))
}

/*
Submits the compute list for processing on the GPU with the given group counts stored in the [param buffer] at [param offset]. Buffer must have been created with [constant STORAGE_BUFFER_USAGE_DISPATCH_INDIRECT] flag.
*/
func (self Go) ComputeListDispatchIndirect(compute_list int, buffer gd.RID, offset int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ComputeListDispatchIndirect(gd.Int(compute_list), buffer, gd.Int(offset))
}

/*
Raises a Vulkan compute barrier in the specified [param compute_list].
*/
func (self Go) ComputeListAddBarrier(compute_list int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ComputeListAddBarrier(gd.Int(compute_list))
}

/*
Finishes a list of compute commands created with the [code]compute_*[/code] methods.
*/
func (self Go) ComputeListEnd() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ComputeListEnd()
}

/*
Tries to free an object in the RenderingDevice. To avoid memory leaks, this should be called after using an object as memory management does not occur automatically when using RenderingDevice directly.
*/
func (self Go) FreeRid(rid gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FreeRid(rid)
}

/*
Creates a timestamp marker with the specified [param name]. This is used for performance reporting with the [method get_captured_timestamp_cpu_time], [method get_captured_timestamp_gpu_time] and [method get_captured_timestamp_name] methods.
*/
func (self Go) CaptureTimestamp(name string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).CaptureTimestamp(gc.String(name))
}

/*
Returns the total number of timestamps (rendering steps) available for profiling.
*/
func (self Go) GetCapturedTimestampsCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCapturedTimestampsCount()))
}

/*
Returns the index of the last frame rendered that has rendering timestamps available for querying.
*/
func (self Go) GetCapturedTimestampsFrame() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCapturedTimestampsFrame()))
}

/*
Returns the timestamp in GPU time for the rendering step specified by [param index] (in microseconds since the engine started). See also [method get_captured_timestamp_cpu_time] and [method capture_timestamp].
*/
func (self Go) GetCapturedTimestampGpuTime(index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCapturedTimestampGpuTime(gd.Int(index))))
}

/*
Returns the timestamp in CPU time for the rendering step specified by [param index] (in microseconds since the engine started). See also [method get_captured_timestamp_gpu_time] and [method capture_timestamp].
*/
func (self Go) GetCapturedTimestampCpuTime(index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCapturedTimestampCpuTime(gd.Int(index))))
}

/*
Returns the timestamp's name for the rendering step specified by [param index]. See also [method capture_timestamp].
*/
func (self Go) GetCapturedTimestampName(index int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetCapturedTimestampName(gc, gd.Int(index)).String())
}

/*
Returns the value of the specified [param limit]. This limit varies depending on the current graphics hardware (and sometimes the driver version). If the given limit is exceeded, rendering errors will occur.
Limits for various graphics hardware can be found in the [url=https://vulkan.gpuinfo.org/]Vulkan Hardware Database[/url].
*/
func (self Go) LimitGet(limit classdb.RenderingDeviceLimit) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).LimitGet(limit)))
}

/*
Returns the frame count kept by the graphics API. Higher values result in higher input lag, but with more consistent throughput. For the main [RenderingDevice], frames are cycled (usually 3 with triple-buffered V-Sync enabled). However, local [RenderingDevice]s only have 1 frame.
*/
func (self Go) GetFrameDelay() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetFrameDelay()))
}

/*
Pushes the frame setup and draw command buffers then marks the local device as currently processing (which allows calling [method sync]).
[b]Note:[/b] Only available in local RenderingDevices.
*/
func (self Go) Submit() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Submit()
}

/*
Forces a synchronization between the CPU and GPU, which may be required in certain cases. Only call this when needed, as CPU-GPU synchronization has a performance cost.
[b]Note:[/b] Only available in local RenderingDevices.
[b]Note:[/b] [method sync] can only be called after a [method submit].
*/
func (self Go) Sync() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Sync()
}

/*
This method does nothing.
*/
func (self Go) Barrier() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Barrier(32767, 32767)
}

/*
This method does nothing.
*/
func (self Go) FullBarrier() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FullBarrier()
}

/*
Create a new local [RenderingDevice]. This is most useful for performing compute operations on the GPU independently from the rest of the engine.
*/
func (self Go) CreateLocalDevice() gdclass.RenderingDevice {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.RenderingDevice(class(self).CreateLocalDevice(gc))
}

/*
Sets the resource name for [param id] to [param name]. This is used for debugging with third-party tools such as [url=https://renderdoc.org/]RenderDoc[/url].
The following types of resources can be named: texture, sampler, vertex buffer, index buffer, uniform buffer, texture buffer, storage buffer, uniform set buffer, shader, render pipeline and compute pipeline. Framebuffers cannot be named. Attempting to name an incompatible resource type will print an error.
[b]Note:[/b] Resource names are only set when the engine runs in verbose mode ([method OS.is_stdout_verbose] = [code]true[/code]), or when using an engine build compiled with the [code]dev_mode=yes[/code] SCons option. The graphics driver must also support the [code]VK_EXT_DEBUG_UTILS_EXTENSION_NAME[/code] Vulkan extension for named resources to work.
*/
func (self Go) SetResourceName(id gd.RID, name string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetResourceName(id, gc.String(name))
}

/*
Create a command buffer debug label region that can be displayed in third-party tools such as [url=https://renderdoc.org/]RenderDoc[/url]. All regions must be ended with a [method draw_command_end_label] call. When viewed from the linear series of submissions to a single queue, calls to [method draw_command_begin_label] and [method draw_command_end_label] must be matched and balanced.
The [code]VK_EXT_DEBUG_UTILS_EXTENSION_NAME[/code] Vulkan extension must be available and enabled for command buffer debug label region to work. See also [method draw_command_end_label].
*/
func (self Go) DrawCommandBeginLabel(name string, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawCommandBeginLabel(gc.String(name), color)
}

/*
This method does nothing.
*/
func (self Go) DrawCommandInsertLabel(name string, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawCommandInsertLabel(gc.String(name), color)
}

/*
Ends the command buffer debug label region started by a [method draw_command_begin_label] call.
*/
func (self Go) DrawCommandEndLabel() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawCommandEndLabel()
}

/*
Returns the vendor of the video adapter (e.g. "NVIDIA Corporation"). Equivalent to [method RenderingServer.get_video_adapter_vendor]. See also [method get_device_name].
*/
func (self Go) GetDeviceVendorName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetDeviceVendorName(gc).String())
}

/*
Returns the name of the video adapter (e.g. "GeForce GTX 1080/PCIe/SSE2"). Equivalent to [method RenderingServer.get_video_adapter_name]. See also [method get_device_vendor_name].
*/
func (self Go) GetDeviceName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetDeviceName(gc).String())
}

/*
Returns the universally unique identifier for the pipeline cache. This is used to cache shader files on disk, which avoids shader recompilations on subsequent engine runs. This UUID varies depending on the graphics card model, but also the driver version. Therefore, updating graphics drivers will invalidate the shader cache.
*/
func (self Go) GetDevicePipelineCacheUuid() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetDevicePipelineCacheUuid(gc).String())
}

/*
Returns the memory usage in bytes corresponding to the given [param type]. When using Vulkan, these statistics are calculated by [url=https://github.com/GPUOpen-LibrariesAndSDKs/VulkanMemoryAllocator]Vulkan Memory Allocator[/url].
*/
func (self Go) GetMemoryUsage(atype classdb.RenderingDeviceMemoryType) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetMemoryUsage(atype)))
}

/*
Returns the unique identifier of the driver [param resource] for the specified [param rid]. Some driver resource types ignore the specified [param rid] (see [enum DriverResource] descriptions). [param index] is always ignored but must be specified anyway.
*/
func (self Go) GetDriverResource(resource classdb.RenderingDeviceDriverResource, rid gd.RID, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetDriverResource(resource, rid, gd.Int(index))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RenderingDevice
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RenderingDevice"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Creates a new texture. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
[b]Note:[/b] Not to be confused with [method RenderingServer.texture_2d_create], which creates the Godot-specific [Texture2D] resource as opposed to the graphics API's own texture type.
*/
//go:nosplit
func (self class) TextureCreate(format gdclass.RDTextureFormat, view gdclass.RDTextureView, data gd.ArrayOf[gd.PackedByteArray]) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(format[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(view[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(data))
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_texture_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a shared texture using the specified [param view] and the texture information from [param with_texture].
*/
//go:nosplit
func (self class) TextureCreateShared(view gdclass.RDTextureView, with_texture gd.RID) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(view[0].AsPointer())[0])
	callframe.Arg(frame, with_texture)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_texture_create_shared, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a shared texture using the specified [param view] and the texture information from [param with_texture]'s [param layer] and [param mipmap]. The number of included mipmaps from the original texture can be controlled using the [param mipmaps] parameter. Only relevant for textures with multiple layers, such as 3D textures, texture arrays and cubemaps. For single-layer textures, use [method texture_create_shared]
For 2D textures (which only have one layer), [param layer] must be [code]0[/code].
[b]Note:[/b] Layer slicing is only supported for 2D texture arrays, not 3D textures or cubemaps.
*/
//go:nosplit
func (self class) TextureCreateSharedFromSlice(view gdclass.RDTextureView, with_texture gd.RID, layer gd.Int, mipmap gd.Int, mipmaps gd.Int, slice_type classdb.RenderingDeviceTextureSliceType) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(view[0].AsPointer())[0])
	callframe.Arg(frame, with_texture)
	callframe.Arg(frame, layer)
	callframe.Arg(frame, mipmap)
	callframe.Arg(frame, mipmaps)
	callframe.Arg(frame, slice_type)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_texture_create_shared_from_slice, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an RID for an existing [param image] ([code]VkImage[/code]) with the given [param type], [param format], [param samples], [param usage_flags], [param width], [param height], [param depth], and [param layers]. This can be used to allow Godot to render onto foreign images.
*/
//go:nosplit
func (self class) TextureCreateFromExtension(atype classdb.RenderingDeviceTextureType, format classdb.RenderingDeviceDataFormat, samples classdb.RenderingDeviceTextureSamples, usage_flags classdb.RenderingDeviceTextureUsageBits, image gd.Int, width gd.Int, height gd.Int, depth gd.Int, layers gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, format)
	callframe.Arg(frame, samples)
	callframe.Arg(frame, usage_flags)
	callframe.Arg(frame, image)
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, depth)
	callframe.Arg(frame, layers)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_texture_create_from_extension, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Updates texture data with new data, replacing the previous data in place. The updated texture data must have the same dimensions and format. For 2D textures (which only have one layer), [param layer] must be [code]0[/code]. Returns [constant @GlobalScope.OK] if the update was successful, [constant @GlobalScope.ERR_INVALID_PARAMETER] otherwise.
[b]Note:[/b] Updating textures is forbidden during creation of a draw or compute list.
[b]Note:[/b] The existing [param texture] can't be updated while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to update this texture.
[b]Note:[/b] The existing [param texture] requires the [constant TEXTURE_USAGE_CAN_UPDATE_BIT] to be updatable.
*/
//go:nosplit
func (self class) TextureUpdate(texture gd.RID, layer gd.Int, data gd.PackedByteArray) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	callframe.Arg(frame, layer)
	callframe.Arg(frame, mmm.Get(data))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_texture_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [param texture] data for the specified [param layer] as raw binary data. For 2D textures (which only have one layer), [param layer] must be [code]0[/code].
[b]Note:[/b] [param texture] can't be retrieved while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to retrieve this texture. Otherwise, an error is printed and a empty [PackedByteArray] is returned.
[b]Note:[/b] [param texture] requires the [constant TEXTURE_USAGE_CAN_COPY_FROM_BIT] to be retrieved. Otherwise, an error is printed and a empty [PackedByteArray] is returned.
*/
//go:nosplit
func (self class) TextureGetData(ctx gd.Lifetime, texture gd.RID, layer gd.Int) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_texture_get_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the specified [param format] is supported for the given [param usage_flags], [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) TextureIsFormatSupportedForUsage(format classdb.RenderingDeviceDataFormat, usage_flags classdb.RenderingDeviceTextureUsageBits) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, format)
	callframe.Arg(frame, usage_flags)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_texture_is_format_supported_for_usage, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the [param texture] is shared, [code]false[/code] otherwise. See [RDTextureView].
*/
//go:nosplit
func (self class) TextureIsShared(texture gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_texture_is_shared, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the [param texture] is valid, [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) TextureIsValid(texture gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_texture_is_valid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Copies the [param from_texture] to [param to_texture] with the specified [param from_pos], [param to_pos] and [param size] coordinates. The Z axis of the [param from_pos], [param to_pos] and [param size] must be [code]0[/code] for 2-dimensional textures. Source and destination mipmaps/layers must also be specified, with these parameters being [code]0[/code] for textures without mipmaps or single-layer textures. Returns [constant @GlobalScope.OK] if the texture copy was successful or [constant @GlobalScope.ERR_INVALID_PARAMETER] otherwise.
[b]Note:[/b] [param from_texture] texture can't be copied while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to copy this texture.
[b]Note:[/b] [param from_texture] texture requires the [constant TEXTURE_USAGE_CAN_COPY_FROM_BIT] to be retrieved.
[b]Note:[/b] [param to_texture] can't be copied while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to copy this texture.
[b]Note:[/b] [param to_texture] requires the [constant TEXTURE_USAGE_CAN_COPY_TO_BIT] to be retrieved.
[b]Note:[/b] [param from_texture] and [param to_texture] must be of the same type (color or depth).
*/
//go:nosplit
func (self class) TextureCopy(from_texture gd.RID, to_texture gd.RID, from_pos gd.Vector3, to_pos gd.Vector3, size gd.Vector3, src_mipmap gd.Int, dst_mipmap gd.Int, src_layer gd.Int, dst_layer gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_texture)
	callframe.Arg(frame, to_texture)
	callframe.Arg(frame, from_pos)
	callframe.Arg(frame, to_pos)
	callframe.Arg(frame, size)
	callframe.Arg(frame, src_mipmap)
	callframe.Arg(frame, dst_mipmap)
	callframe.Arg(frame, src_layer)
	callframe.Arg(frame, dst_layer)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_texture_copy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Clears the specified [param texture] by replacing all of its pixels with the specified [param color]. [param base_mipmap] and [param mipmap_count] determine which mipmaps of the texture are affected by this clear operation, while [param base_layer] and [param layer_count] determine which layers of a 3D texture (or texture array) are affected by this clear operation. For 2D textures (which only have one layer by design), [param base_layer] must be [code]0[/code] and [param layer_count] must be [code]1[/code].
[b]Note:[/b] [param texture] can't be cleared while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to clear this texture.
*/
//go:nosplit
func (self class) TextureClear(texture gd.RID, color gd.Color, base_mipmap gd.Int, mipmap_count gd.Int, base_layer gd.Int, layer_count gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	callframe.Arg(frame, color)
	callframe.Arg(frame, base_mipmap)
	callframe.Arg(frame, mipmap_count)
	callframe.Arg(frame, base_layer)
	callframe.Arg(frame, layer_count)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_texture_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Resolves the [param from_texture] texture onto [param to_texture] with multisample antialiasing enabled. This must be used when rendering a framebuffer for MSAA to work. Returns [constant @GlobalScope.OK] if successful, [constant @GlobalScope.ERR_INVALID_PARAMETER] otherwise.
[b]Note:[/b] [param from_texture] and [param to_texture] textures must have the same dimension, format and type (color or depth).
[b]Note:[/b] [param from_texture] can't be copied while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to resolve this texture.
[b]Note:[/b] [param from_texture] requires the [constant TEXTURE_USAGE_CAN_COPY_FROM_BIT] to be retrieved.
[b]Note:[/b] [param from_texture] must be multisampled and must also be 2D (or a slice of a 3D/cubemap texture).
[b]Note:[/b] [param to_texture] can't be copied while a draw list that uses it as part of a framebuffer is being created. Ensure the draw list is finalized (and that the color/depth texture using it is not set to [constant FINAL_ACTION_CONTINUE]) to resolve this texture.
[b]Note:[/b] [param to_texture] texture requires the [constant TEXTURE_USAGE_CAN_COPY_TO_BIT] to be retrieved.
[b]Note:[/b] [param to_texture] texture must [b]not[/b] be multisampled and must also be 2D (or a slice of a 3D/cubemap texture).
*/
//go:nosplit
func (self class) TextureResolveMultisample(from_texture gd.RID, to_texture gd.RID) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_texture)
	callframe.Arg(frame, to_texture)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_texture_resolve_multisample, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the data format used to create this texture.
*/
//go:nosplit
func (self class) TextureGetFormat(ctx gd.Lifetime, texture gd.RID) gdclass.RDTextureFormat {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_texture_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.RDTextureFormat
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the internal graphics handle for this texture object. For use when communicating with third-party APIs mostly with GDExtension.
[b]Note:[/b] This function returns a [code]uint64_t[/code] which internally maps to a [code]GLuint[/code] (OpenGL) or [code]VkImage[/code] (Vulkan).
*/
//go:nosplit
func (self class) TextureGetNativeHandle(texture gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_texture_get_native_handle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new framebuffer format with the specified [param attachments] and [param view_count]. Returns the new framebuffer's unique framebuffer format ID.
If [param view_count] is greater than or equal to [code]2[/code], enables multiview which is used for VR rendering. This requires support for the Vulkan multiview extension.
*/
//go:nosplit
func (self class) FramebufferFormatCreate(attachments gd.ArrayOf[gdclass.RDAttachmentFormat], view_count gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(attachments))
	callframe.Arg(frame, view_count)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_framebuffer_format_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a multipass framebuffer format with the specified [param attachments], [param passes] and [param view_count] and returns its ID. If [param view_count] is greater than or equal to [code]2[/code], enables multiview which is used for VR rendering. This requires support for the Vulkan multiview extension.
*/
//go:nosplit
func (self class) FramebufferFormatCreateMultipass(attachments gd.ArrayOf[gdclass.RDAttachmentFormat], passes gd.ArrayOf[gdclass.RDFramebufferPass], view_count gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(attachments))
	callframe.Arg(frame, mmm.Get(passes))
	callframe.Arg(frame, view_count)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_framebuffer_format_create_multipass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new empty framebuffer format with the specified number of [param samples] and returns its ID.
*/
//go:nosplit
func (self class) FramebufferFormatCreateEmpty(samples classdb.RenderingDeviceTextureSamples) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, samples)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_framebuffer_format_create_empty, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of texture samples used for the given framebuffer [param format] ID (returned by [method framebuffer_get_format]).
*/
//go:nosplit
func (self class) FramebufferFormatGetTextureSamples(format gd.Int, render_pass gd.Int) classdb.RenderingDeviceTextureSamples {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, format)
	callframe.Arg(frame, render_pass)
	var r_ret = callframe.Ret[classdb.RenderingDeviceTextureSamples](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_framebuffer_format_get_texture_samples, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new framebuffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) FramebufferCreate(textures gd.ArrayOf[gd.RID], validate_with_format gd.Int, view_count gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(textures))
	callframe.Arg(frame, validate_with_format)
	callframe.Arg(frame, view_count)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_framebuffer_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new multipass framebuffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) FramebufferCreateMultipass(textures gd.ArrayOf[gd.RID], passes gd.ArrayOf[gdclass.RDFramebufferPass], validate_with_format gd.Int, view_count gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(textures))
	callframe.Arg(frame, mmm.Get(passes))
	callframe.Arg(frame, validate_with_format)
	callframe.Arg(frame, view_count)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_framebuffer_create_multipass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new empty framebuffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) FramebufferCreateEmpty(size gd.Vector2i, samples classdb.RenderingDeviceTextureSamples, validate_with_format gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, samples)
	callframe.Arg(frame, validate_with_format)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_framebuffer_create_empty, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the format ID of the framebuffer specified by the [param framebuffer] RID. This ID is guaranteed to be unique for the same formats and does not need to be freed.
*/
//go:nosplit
func (self class) FramebufferGetFormat(framebuffer gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, framebuffer)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_framebuffer_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the framebuffer specified by the [param framebuffer] RID is valid, [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) FramebufferIsValid(framebuffer gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, framebuffer)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_framebuffer_is_valid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new sampler. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) SamplerCreate(state gdclass.RDSamplerState) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(state[0].AsPointer())[0])
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_sampler_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if implementation supports using a texture of [param format] with the given [param sampler_filter].
*/
//go:nosplit
func (self class) SamplerIsFormatSupportedForFilter(format classdb.RenderingDeviceDataFormat, sampler_filter classdb.RenderingDeviceSamplerFilter) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, format)
	callframe.Arg(frame, sampler_filter)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_sampler_is_format_supported_for_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) VertexBufferCreate(size_bytes gd.Int, data gd.PackedByteArray, use_as_storage bool) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size_bytes)
	callframe.Arg(frame, mmm.Get(data))
	callframe.Arg(frame, use_as_storage)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_vertex_buffer_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new vertex format with the specified [param vertex_descriptions]. Returns a unique vertex format ID corresponding to the newly created vertex format.
*/
//go:nosplit
func (self class) VertexFormatCreate(vertex_descriptions gd.ArrayOf[gdclass.RDVertexAttribute]) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(vertex_descriptions))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_vertex_format_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a vertex array based on the specified buffers. Optionally, [param offsets] (in bytes) may be defined for each buffer.
*/
//go:nosplit
func (self class) VertexArrayCreate(vertex_count gd.Int, vertex_format gd.Int, src_buffers gd.ArrayOf[gd.RID], offsets gd.PackedInt64Array) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, vertex_count)
	callframe.Arg(frame, vertex_format)
	callframe.Arg(frame, mmm.Get(src_buffers))
	callframe.Arg(frame, mmm.Get(offsets))
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_vertex_array_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new index buffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) IndexBufferCreate(size_indices gd.Int, format classdb.RenderingDeviceIndexBufferFormat, data gd.PackedByteArray, use_restart_indices bool) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size_indices)
	callframe.Arg(frame, format)
	callframe.Arg(frame, mmm.Get(data))
	callframe.Arg(frame, use_restart_indices)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_index_buffer_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new index array. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) IndexArrayCreate(index_buffer gd.RID, index_offset gd.Int, index_count gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index_buffer)
	callframe.Arg(frame, index_offset)
	callframe.Arg(frame, index_count)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_index_array_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Compiles a SPIR-V from the shader source code in [param shader_source] and returns the SPIR-V as a [RDShaderSPIRV]. This intermediate language shader is portable across different GPU models and driver versions, but cannot be run directly by GPUs until compiled into a binary shader using [method shader_compile_binary_from_spirv].
If [param allow_cache] is [code]true[/code], make use of the shader cache generated by Godot. This avoids a potentially lengthy shader compilation step if the shader is already in cache. If [param allow_cache] is [code]false[/code], Godot's shader cache is ignored and the shader will always be recompiled.
*/
//go:nosplit
func (self class) ShaderCompileSpirvFromSource(ctx gd.Lifetime, shader_source gdclass.RDShaderSource, allow_cache bool) gdclass.RDShaderSPIRV {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(shader_source[0].AsPointer())[0])
	callframe.Arg(frame, allow_cache)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_shader_compile_spirv_from_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.RDShaderSPIRV
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Compiles a binary shader from [param spirv_data] and returns the compiled binary data as a [PackedByteArray]. This compiled shader is specific to the GPU model and driver version used; it will not work on different GPU models or even different driver versions. See also [method shader_compile_spirv_from_source].
[param name] is an optional human-readable name that can be given to the compiled shader for organizational purposes.
*/
//go:nosplit
func (self class) ShaderCompileBinaryFromSpirv(ctx gd.Lifetime, spirv_data gdclass.RDShaderSPIRV, name gd.String) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(spirv_data[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_shader_compile_binary_from_spirv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Creates a new shader instance from SPIR-V intermediate code. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method. See also [method shader_compile_spirv_from_source] and [method shader_create_from_bytecode].
*/
//go:nosplit
func (self class) ShaderCreateFromSpirv(spirv_data gdclass.RDShaderSPIRV, name gd.String) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(spirv_data[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_shader_create_from_spirv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new shader instance from a binary compiled shader. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method. See also [method shader_compile_binary_from_spirv] and [method shader_create_from_spirv].
*/
//go:nosplit
func (self class) ShaderCreateFromBytecode(binary_data gd.PackedByteArray, placeholder_rid gd.RID) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(binary_data))
	callframe.Arg(frame, placeholder_rid)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_shader_create_from_bytecode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Create a placeholder RID by allocating an RID without initializing it for use in [method shader_create_from_bytecode]. This allows you to create an RID for a shader and pass it around, but defer compiling the shader to a later time.
*/
//go:nosplit
func (self class) ShaderCreatePlaceholder() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_shader_create_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the internal vertex input mask. Internally, the vertex input mask is an unsigned integer consisting of the locations (specified in GLSL via. [code]layout(location = ...)[/code]) of the input variables (specified in GLSL by the [code]in[/code] keyword).
*/
//go:nosplit
func (self class) ShaderGetVertexInputAttributeMask(shader gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shader)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_shader_get_vertex_input_attribute_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new uniform buffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) UniformBufferCreate(size_bytes gd.Int, data gd.PackedByteArray) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size_bytes)
	callframe.Arg(frame, mmm.Get(data))
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_uniform_buffer_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a [url=https://vkguide.dev/docs/chapter-4/storage_buffers/]storage buffer[/url] with the specified [param data] and [param usage]. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) StorageBufferCreate(size_bytes gd.Int, data gd.PackedByteArray, usage classdb.RenderingDeviceStorageBufferUsage) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size_bytes)
	callframe.Arg(frame, mmm.Get(data))
	callframe.Arg(frame, usage)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_storage_buffer_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new texture buffer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) TextureBufferCreate(size_bytes gd.Int, format classdb.RenderingDeviceDataFormat, data gd.PackedByteArray) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size_bytes)
	callframe.Arg(frame, format)
	callframe.Arg(frame, mmm.Get(data))
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_texture_buffer_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new uniform set. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) UniformSetCreate(uniforms gd.ArrayOf[gdclass.RDUniform], shader gd.RID, shader_set gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(uniforms))
	callframe.Arg(frame, shader)
	callframe.Arg(frame, shader_set)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_uniform_set_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Checks if the [param uniform_set] is valid, i.e. is owned.
*/
//go:nosplit
func (self class) UniformSetIsValid(uniform_set gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, uniform_set)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_uniform_set_is_valid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Copies [param size] bytes from the [param src_buffer] at [param src_offset] into [param dst_buffer] at [param dst_offset].
Prints an error if:
- [param size] exceeds the size of either [param src_buffer] or [param dst_buffer] at their corresponding offsets
- a draw list is currently active (created by [method draw_list_begin])
- a compute list is currently active (created by [method compute_list_begin])
*/
//go:nosplit
func (self class) BufferCopy(src_buffer gd.RID, dst_buffer gd.RID, src_offset gd.Int, dst_offset gd.Int, size gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, src_buffer)
	callframe.Arg(frame, dst_buffer)
	callframe.Arg(frame, src_offset)
	callframe.Arg(frame, dst_offset)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_buffer_copy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Updates a region of [param size_bytes] bytes, starting at [param offset], in the buffer, with the specified [param data].
Prints an error if:
- the region specified by [param offset] + [param size_bytes] exceeds the buffer
- a draw list is currently active (created by [method draw_list_begin])
- a compute list is currently active (created by [method compute_list_begin])
*/
//go:nosplit
func (self class) BufferUpdate(buffer gd.RID, offset gd.Int, size_bytes gd.Int, data gd.PackedByteArray) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, buffer)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, size_bytes)
	callframe.Arg(frame, mmm.Get(data))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_buffer_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Clears the contents of the [param buffer], clearing [param size_bytes] bytes, starting at [param offset].
Prints an error if:
- the size isn't a multiple of four
- the region specified by [param offset] + [param size_bytes] exceeds the buffer
- a draw list is currently active (created by [method draw_list_begin])
- a compute list is currently active (created by [method compute_list_begin])
*/
//go:nosplit
func (self class) BufferClear(buffer gd.RID, offset gd.Int, size_bytes gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, buffer)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, size_bytes)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_buffer_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a copy of the data of the specified [param buffer], optionally [param offset_bytes] and [param size_bytes] can be set to copy only a portion of the buffer.
*/
//go:nosplit
func (self class) BufferGetData(ctx gd.Lifetime, buffer gd.RID, offset_bytes gd.Int, size_bytes gd.Int) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, buffer)
	callframe.Arg(frame, offset_bytes)
	callframe.Arg(frame, size_bytes)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_buffer_get_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Creates a new render pipeline. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) RenderPipelineCreate(shader gd.RID, framebuffer_format gd.Int, vertex_format gd.Int, primitive classdb.RenderingDeviceRenderPrimitive, rasterization_state gdclass.RDPipelineRasterizationState, multisample_state gdclass.RDPipelineMultisampleState, stencil_state gdclass.RDPipelineDepthStencilState, color_blend_state gdclass.RDPipelineColorBlendState, dynamic_state_flags classdb.RenderingDevicePipelineDynamicStateFlags, for_render_pass gd.Int, specialization_constants gd.ArrayOf[gdclass.RDPipelineSpecializationConstant]) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shader)
	callframe.Arg(frame, framebuffer_format)
	callframe.Arg(frame, vertex_format)
	callframe.Arg(frame, primitive)
	callframe.Arg(frame, mmm.Get(rasterization_state[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(multisample_state[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(stencil_state[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(color_blend_state[0].AsPointer())[0])
	callframe.Arg(frame, dynamic_state_flags)
	callframe.Arg(frame, for_render_pass)
	callframe.Arg(frame, mmm.Get(specialization_constants))
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_render_pipeline_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the render pipeline specified by the [param render_pipeline] RID is valid, [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) RenderPipelineIsValid(render_pipeline gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, render_pipeline)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_render_pipeline_is_valid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new compute pipeline. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingDevice's [method free_rid] method.
*/
//go:nosplit
func (self class) ComputePipelineCreate(shader gd.RID, specialization_constants gd.ArrayOf[gdclass.RDPipelineSpecializationConstant]) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shader)
	callframe.Arg(frame, mmm.Get(specialization_constants))
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_compute_pipeline_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the compute pipeline specified by the [param compute_pipeline] RID is valid, [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) ComputePipelineIsValid(compute_pipeline gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, compute_pipeline)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_compute_pipeline_is_valid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the window width matching the graphics API context for the given window ID (in pixels). Despite the parameter being named [param screen], this returns the [i]window[/i] size. See also [method screen_get_height].
[b]Note:[/b] Only the main [RenderingDevice] returned by [method RenderingServer.get_rendering_device] has a width. If called on a local [RenderingDevice], this method prints an error and returns [constant INVALID_ID].
*/
//go:nosplit
func (self class) ScreenGetWidth(screen gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_screen_get_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the window height matching the graphics API context for the given window ID (in pixels). Despite the parameter being named [param screen], this returns the [i]window[/i] size. See also [method screen_get_width].
[b]Note:[/b] Only the main [RenderingDevice] returned by [method RenderingServer.get_rendering_device] has a height. If called on a local [RenderingDevice], this method prints an error and returns [constant INVALID_ID].
*/
//go:nosplit
func (self class) ScreenGetHeight(screen gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_screen_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the framebuffer format of the given screen.
[b]Note:[/b] Only the main [RenderingDevice] returned by [method RenderingServer.get_rendering_device] has a format. If called on a local [RenderingDevice], this method prints an error and returns [constant INVALID_ID].
*/
//go:nosplit
func (self class) ScreenGetFramebufferFormat(screen gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_screen_get_framebuffer_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
High-level variant of [method draw_list_begin], with the parameters automatically being adjusted for drawing onto the window specified by the [param screen] ID.
[b]Note:[/b] Cannot be used with local RenderingDevices, as these don't have a screen. If called on a local RenderingDevice, [method draw_list_begin_for_screen] returns [constant INVALID_ID].
*/
//go:nosplit
func (self class) DrawListBeginForScreen(screen gd.Int, clear_color gd.Color) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	callframe.Arg(frame, clear_color)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_draw_list_begin_for_screen, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Starts a list of raster drawing commands created with the [code]draw_*[/code] methods. The returned value should be passed to other [code]draw_list_*[/code] functions.
Multiple draw lists cannot be created at the same time; you must finish the previous draw list first using [method draw_list_end].
A simple drawing operation might look like this (code is not a complete example):
[codeblock]
var rd = RenderingDevice.new()
var clear_colors = PackedColorArray([Color(0, 0, 0, 0), Color(0, 0, 0, 0), Color(0, 0, 0, 0)])
var draw_list = rd.draw_list_begin(framebuffers[i], RenderingDevice.INITIAL_ACTION_CLEAR, RenderingDevice.FINAL_ACTION_READ, RenderingDevice.INITIAL_ACTION_CLEAR, RenderingDevice.FINAL_ACTION_DISCARD, clear_colors)

# Draw opaque.
rd.draw_list_bind_render_pipeline(draw_list, raster_pipeline)
rd.draw_list_bind_uniform_set(draw_list, raster_base_uniform, 0)
rd.draw_list_set_push_constant(draw_list, raster_push_constant, raster_push_constant.size())
rd.draw_list_draw(draw_list, false, 1, slice_triangle_count[i] * 3)
# Draw wire.
rd.draw_list_bind_render_pipeline(draw_list, raster_pipeline_wire)
rd.draw_list_bind_uniform_set(draw_list, raster_base_uniform, 0)
rd.draw_list_set_push_constant(draw_list, raster_push_constant, raster_push_constant.size())
rd.draw_list_draw(draw_list, false, 1, slice_triangle_count[i] * 3)

rd.draw_list_end()
[/codeblock]
*/
//go:nosplit
func (self class) DrawListBegin(framebuffer gd.RID, initial_color_action classdb.RenderingDeviceInitialAction, final_color_action classdb.RenderingDeviceFinalAction, initial_depth_action classdb.RenderingDeviceInitialAction, final_depth_action classdb.RenderingDeviceFinalAction, clear_color_values gd.PackedColorArray, clear_depth gd.Float, clear_stencil gd.Int, region gd.Rect2) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, framebuffer)
	callframe.Arg(frame, initial_color_action)
	callframe.Arg(frame, final_color_action)
	callframe.Arg(frame, initial_depth_action)
	callframe.Arg(frame, final_depth_action)
	callframe.Arg(frame, mmm.Get(clear_color_values))
	callframe.Arg(frame, clear_depth)
	callframe.Arg(frame, clear_stencil)
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_draw_list_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
This method does nothing and always returns an empty [PackedInt64Array].
*/
//go:nosplit
func (self class) DrawListBeginSplit(ctx gd.Lifetime, framebuffer gd.RID, splits gd.Int, initial_color_action classdb.RenderingDeviceInitialAction, final_color_action classdb.RenderingDeviceFinalAction, initial_depth_action classdb.RenderingDeviceInitialAction, final_depth_action classdb.RenderingDeviceFinalAction, clear_color_values gd.PackedColorArray, clear_depth gd.Float, clear_stencil gd.Int, region gd.Rect2, storage_textures gd.ArrayOf[gd.RID]) gd.PackedInt64Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, framebuffer)
	callframe.Arg(frame, splits)
	callframe.Arg(frame, initial_color_action)
	callframe.Arg(frame, final_color_action)
	callframe.Arg(frame, initial_depth_action)
	callframe.Arg(frame, final_depth_action)
	callframe.Arg(frame, mmm.Get(clear_color_values))
	callframe.Arg(frame, clear_depth)
	callframe.Arg(frame, clear_stencil)
	callframe.Arg(frame, region)
	callframe.Arg(frame, mmm.Get(storage_textures))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_draw_list_begin_split, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt64Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets blend constants for the specified [param draw_list] to [param color]. Blend constants are used only if the graphics pipeline is created with [constant DYNAMIC_STATE_BLEND_CONSTANTS] flag set.
*/
//go:nosplit
func (self class) DrawListSetBlendConstants(draw_list gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, draw_list)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_draw_list_set_blend_constants, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Binds [param render_pipeline] to the specified [param draw_list].
*/
//go:nosplit
func (self class) DrawListBindRenderPipeline(draw_list gd.Int, render_pipeline gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, draw_list)
	callframe.Arg(frame, render_pipeline)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_draw_list_bind_render_pipeline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Binds [param uniform_set] to the specified [param draw_list]. A [param set_index] must also be specified, which is an identifier starting from [code]0[/code] that must match the one expected by the draw list.
*/
//go:nosplit
func (self class) DrawListBindUniformSet(draw_list gd.Int, uniform_set gd.RID, set_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, draw_list)
	callframe.Arg(frame, uniform_set)
	callframe.Arg(frame, set_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_draw_list_bind_uniform_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Binds [param vertex_array] to the specified [param draw_list].
*/
//go:nosplit
func (self class) DrawListBindVertexArray(draw_list gd.Int, vertex_array gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, draw_list)
	callframe.Arg(frame, vertex_array)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_draw_list_bind_vertex_array, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Binds [param index_array] to the specified [param draw_list].
*/
//go:nosplit
func (self class) DrawListBindIndexArray(draw_list gd.Int, index_array gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, draw_list)
	callframe.Arg(frame, index_array)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_draw_list_bind_index_array, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the push constant data to [param buffer] for the specified [param draw_list]. The shader determines how this binary data is used. The buffer's size in bytes must also be specified in [param size_bytes] (this can be obtained by calling the [method PackedByteArray.size] method on the passed [param buffer]).
*/
//go:nosplit
func (self class) DrawListSetPushConstant(draw_list gd.Int, buffer gd.PackedByteArray, size_bytes gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, draw_list)
	callframe.Arg(frame, mmm.Get(buffer))
	callframe.Arg(frame, size_bytes)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_draw_list_set_push_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Submits [param draw_list] for rendering on the GPU. This is the raster equivalent to [method compute_list_dispatch].
*/
//go:nosplit
func (self class) DrawListDraw(draw_list gd.Int, use_indices bool, instances gd.Int, procedural_vertex_count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, draw_list)
	callframe.Arg(frame, use_indices)
	callframe.Arg(frame, instances)
	callframe.Arg(frame, procedural_vertex_count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_draw_list_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a scissor rectangle and enables it for the specified [param draw_list]. Scissor rectangles are used for clipping by discarding fragments that fall outside a specified rectangular portion of the screen. See also [method draw_list_disable_scissor].
[b]Note:[/b] The specified [param rect] is automatically intersected with the screen's dimensions, which means it cannot exceed the screen's dimensions.
*/
//go:nosplit
func (self class) DrawListEnableScissor(draw_list gd.Int, rect gd.Rect2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, draw_list)
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_draw_list_enable_scissor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes and disables the scissor rectangle for the specified [param draw_list]. See also [method draw_list_enable_scissor].
*/
//go:nosplit
func (self class) DrawListDisableScissor(draw_list gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, draw_list)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_draw_list_disable_scissor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Switches to the next draw pass.
*/
//go:nosplit
func (self class) DrawListSwitchToNextPass() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_draw_list_switch_to_next_pass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
This method does nothing and always returns an empty [PackedInt64Array].
*/
//go:nosplit
func (self class) DrawListSwitchToNextPassSplit(ctx gd.Lifetime, splits gd.Int) gd.PackedInt64Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, splits)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_draw_list_switch_to_next_pass_split, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt64Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Finishes a list of raster drawing commands created with the [code]draw_*[/code] methods.
*/
//go:nosplit
func (self class) DrawListEnd()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_draw_list_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Starts a list of compute commands created with the [code]compute_*[/code] methods. The returned value should be passed to other [code]compute_list_*[/code] functions.
Multiple compute lists cannot be created at the same time; you must finish the previous compute list first using [method compute_list_end].
A simple compute operation might look like this (code is not a complete example):
[codeblock]
var rd = RenderingDevice.new()
var compute_list = rd.compute_list_begin()

rd.compute_list_bind_compute_pipeline(compute_list, compute_shader_dilate_pipeline)
rd.compute_list_bind_uniform_set(compute_list, compute_base_uniform_set, 0)
rd.compute_list_bind_uniform_set(compute_list, dilate_uniform_set, 1)

for i in atlas_slices:
    rd.compute_list_set_push_constant(compute_list, push_constant, push_constant.size())
    rd.compute_list_dispatch(compute_list, group_size.x, group_size.y, group_size.z)
    # No barrier, let them run all together.

rd.compute_list_end()
[/codeblock]
*/
//go:nosplit
func (self class) ComputeListBegin() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_compute_list_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Tells the GPU what compute pipeline to use when processing the compute list. If the shader has changed since the last time this function was called, Godot will unbind all descriptor sets and will re-bind them inside [method compute_list_dispatch].
*/
//go:nosplit
func (self class) ComputeListBindComputePipeline(compute_list gd.Int, compute_pipeline gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, compute_list)
	callframe.Arg(frame, compute_pipeline)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_compute_list_bind_compute_pipeline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the push constant data to [param buffer] for the specified [param compute_list]. The shader determines how this binary data is used. The buffer's size in bytes must also be specified in [param size_bytes] (this can be obtained by calling the [method PackedByteArray.size] method on the passed [param buffer]).
*/
//go:nosplit
func (self class) ComputeListSetPushConstant(compute_list gd.Int, buffer gd.PackedByteArray, size_bytes gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, compute_list)
	callframe.Arg(frame, mmm.Get(buffer))
	callframe.Arg(frame, size_bytes)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_compute_list_set_push_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Binds the [param uniform_set] to this [param compute_list]. Godot ensures that all textures in the uniform set have the correct Vulkan access masks. If Godot had to change access masks of textures, it will raise a Vulkan image memory barrier.
*/
//go:nosplit
func (self class) ComputeListBindUniformSet(compute_list gd.Int, uniform_set gd.RID, set_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, compute_list)
	callframe.Arg(frame, uniform_set)
	callframe.Arg(frame, set_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_compute_list_bind_uniform_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Submits the compute list for processing on the GPU. This is the compute equivalent to [method draw_list_draw].
*/
//go:nosplit
func (self class) ComputeListDispatch(compute_list gd.Int, x_groups gd.Int, y_groups gd.Int, z_groups gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, compute_list)
	callframe.Arg(frame, x_groups)
	callframe.Arg(frame, y_groups)
	callframe.Arg(frame, z_groups)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_compute_list_dispatch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Submits the compute list for processing on the GPU with the given group counts stored in the [param buffer] at [param offset]. Buffer must have been created with [constant STORAGE_BUFFER_USAGE_DISPATCH_INDIRECT] flag.
*/
//go:nosplit
func (self class) ComputeListDispatchIndirect(compute_list gd.Int, buffer gd.RID, offset gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, compute_list)
	callframe.Arg(frame, buffer)
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_compute_list_dispatch_indirect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Raises a Vulkan compute barrier in the specified [param compute_list].
*/
//go:nosplit
func (self class) ComputeListAddBarrier(compute_list gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, compute_list)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_compute_list_add_barrier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Finishes a list of compute commands created with the [code]compute_*[/code] methods.
*/
//go:nosplit
func (self class) ComputeListEnd()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_compute_list_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Tries to free an object in the RenderingDevice. To avoid memory leaks, this should be called after using an object as memory management does not occur automatically when using RenderingDevice directly.
*/
//go:nosplit
func (self class) FreeRid(rid gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_free_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a timestamp marker with the specified [param name]. This is used for performance reporting with the [method get_captured_timestamp_cpu_time], [method get_captured_timestamp_gpu_time] and [method get_captured_timestamp_name] methods.
*/
//go:nosplit
func (self class) CaptureTimestamp(name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_capture_timestamp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the total number of timestamps (rendering steps) available for profiling.
*/
//go:nosplit
func (self class) GetCapturedTimestampsCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_get_captured_timestamps_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the index of the last frame rendered that has rendering timestamps available for querying.
*/
//go:nosplit
func (self class) GetCapturedTimestampsFrame() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_get_captured_timestamps_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the timestamp in GPU time for the rendering step specified by [param index] (in microseconds since the engine started). See also [method get_captured_timestamp_cpu_time] and [method capture_timestamp].
*/
//go:nosplit
func (self class) GetCapturedTimestampGpuTime(index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_get_captured_timestamp_gpu_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the timestamp in CPU time for the rendering step specified by [param index] (in microseconds since the engine started). See also [method get_captured_timestamp_gpu_time] and [method capture_timestamp].
*/
//go:nosplit
func (self class) GetCapturedTimestampCpuTime(index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_get_captured_timestamp_cpu_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the timestamp's name for the rendering step specified by [param index]. See also [method capture_timestamp].
*/
//go:nosplit
func (self class) GetCapturedTimestampName(ctx gd.Lifetime, index gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_get_captured_timestamp_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the value of the specified [param limit]. This limit varies depending on the current graphics hardware (and sometimes the driver version). If the given limit is exceeded, rendering errors will occur.
Limits for various graphics hardware can be found in the [url=https://vulkan.gpuinfo.org/]Vulkan Hardware Database[/url].
*/
//go:nosplit
func (self class) LimitGet(limit classdb.RenderingDeviceLimit) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, limit)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_limit_get, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the frame count kept by the graphics API. Higher values result in higher input lag, but with more consistent throughput. For the main [RenderingDevice], frames are cycled (usually 3 with triple-buffered V-Sync enabled). However, local [RenderingDevice]s only have 1 frame.
*/
//go:nosplit
func (self class) GetFrameDelay() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_get_frame_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Pushes the frame setup and draw command buffers then marks the local device as currently processing (which allows calling [method sync]).
[b]Note:[/b] Only available in local RenderingDevices.
*/
//go:nosplit
func (self class) Submit()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_submit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Forces a synchronization between the CPU and GPU, which may be required in certain cases. Only call this when needed, as CPU-GPU synchronization has a performance cost.
[b]Note:[/b] Only available in local RenderingDevices.
[b]Note:[/b] [method sync] can only be called after a [method submit].
*/
//go:nosplit
func (self class) Sync()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_sync, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
This method does nothing.
*/
//go:nosplit
func (self class) Barrier(from classdb.RenderingDeviceBarrierMask, to classdb.RenderingDeviceBarrierMask)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_barrier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
This method does nothing.
*/
//go:nosplit
func (self class) FullBarrier()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_full_barrier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Create a new local [RenderingDevice]. This is most useful for performing compute operations on the GPU independently from the rest of the engine.
*/
//go:nosplit
func (self class) CreateLocalDevice(ctx gd.Lifetime) gdclass.RenderingDevice {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_create_local_device, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.RenderingDevice
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the resource name for [param id] to [param name]. This is used for debugging with third-party tools such as [url=https://renderdoc.org/]RenderDoc[/url].
The following types of resources can be named: texture, sampler, vertex buffer, index buffer, uniform buffer, texture buffer, storage buffer, uniform set buffer, shader, render pipeline and compute pipeline. Framebuffers cannot be named. Attempting to name an incompatible resource type will print an error.
[b]Note:[/b] Resource names are only set when the engine runs in verbose mode ([method OS.is_stdout_verbose] = [code]true[/code]), or when using an engine build compiled with the [code]dev_mode=yes[/code] SCons option. The graphics driver must also support the [code]VK_EXT_DEBUG_UTILS_EXTENSION_NAME[/code] Vulkan extension for named resources to work.
*/
//go:nosplit
func (self class) SetResourceName(id gd.RID, name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_set_resource_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Create a command buffer debug label region that can be displayed in third-party tools such as [url=https://renderdoc.org/]RenderDoc[/url]. All regions must be ended with a [method draw_command_end_label] call. When viewed from the linear series of submissions to a single queue, calls to [method draw_command_begin_label] and [method draw_command_end_label] must be matched and balanced.
The [code]VK_EXT_DEBUG_UTILS_EXTENSION_NAME[/code] Vulkan extension must be available and enabled for command buffer debug label region to work. See also [method draw_command_end_label].
*/
//go:nosplit
func (self class) DrawCommandBeginLabel(name gd.String, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_draw_command_begin_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
This method does nothing.
*/
//go:nosplit
func (self class) DrawCommandInsertLabel(name gd.String, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_draw_command_insert_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Ends the command buffer debug label region started by a [method draw_command_begin_label] call.
*/
//go:nosplit
func (self class) DrawCommandEndLabel()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_draw_command_end_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the vendor of the video adapter (e.g. "NVIDIA Corporation"). Equivalent to [method RenderingServer.get_video_adapter_vendor]. See also [method get_device_name].
*/
//go:nosplit
func (self class) GetDeviceVendorName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_get_device_vendor_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the name of the video adapter (e.g. "GeForce GTX 1080/PCIe/SSE2"). Equivalent to [method RenderingServer.get_video_adapter_name]. See also [method get_device_vendor_name].
*/
//go:nosplit
func (self class) GetDeviceName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_get_device_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the universally unique identifier for the pipeline cache. This is used to cache shader files on disk, which avoids shader recompilations on subsequent engine runs. This UUID varies depending on the graphics card model, but also the driver version. Therefore, updating graphics drivers will invalidate the shader cache.
*/
//go:nosplit
func (self class) GetDevicePipelineCacheUuid(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_get_device_pipeline_cache_uuid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the memory usage in bytes corresponding to the given [param type]. When using Vulkan, these statistics are calculated by [url=https://github.com/GPUOpen-LibrariesAndSDKs/VulkanMemoryAllocator]Vulkan Memory Allocator[/url].
*/
//go:nosplit
func (self class) GetMemoryUsage(atype classdb.RenderingDeviceMemoryType) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_get_memory_usage, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the unique identifier of the driver [param resource] for the specified [param rid]. Some driver resource types ignore the specified [param rid] (see [enum DriverResource] descriptions). [param index] is always ignored but must be specified anyway.
*/
//go:nosplit
func (self class) GetDriverResource(resource classdb.RenderingDeviceDriverResource, rid gd.RID, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, resource)
	callframe.Arg(frame, rid)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingDevice.Bind_get_driver_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRenderingDevice() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRenderingDevice() Go { return *((*Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("RenderingDevice", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type DeviceType = classdb.RenderingDeviceDeviceType

const (
/*Rendering device type does not match any of the other enum values or is unknown.*/
	DeviceTypeOther DeviceType = 0
/*Rendering device is an integrated GPU, which is typically [i](but not always)[/i] slower than dedicated GPUs ([constant DEVICE_TYPE_DISCRETE_GPU]). On Android and iOS, the rendering device type is always considered to be [constant DEVICE_TYPE_INTEGRATED_GPU].*/
	DeviceTypeIntegratedGpu DeviceType = 1
/*Rendering device is a dedicated GPU, which is typically [i](but not always)[/i] faster than integrated GPUs ([constant DEVICE_TYPE_INTEGRATED_GPU]).*/
	DeviceTypeDiscreteGpu DeviceType = 2
/*Rendering device is an emulated GPU in a virtual environment. This is typically much slower than the host GPU, which means the expected performance level on a dedicated GPU will be roughly equivalent to [constant DEVICE_TYPE_INTEGRATED_GPU]. Virtual machine GPU passthrough (such as VFIO) will not report the device type as [constant DEVICE_TYPE_VIRTUAL_GPU]. Instead, the host GPU's device type will be reported as if the GPU was not emulated.*/
	DeviceTypeVirtualGpu DeviceType = 3
/*Rendering device is provided by software emulation (such as Lavapipe or [url=https://github.com/google/swiftshader]SwiftShader[/url]). This is the slowest kind of rendering device available; it's typically much slower than [constant DEVICE_TYPE_INTEGRATED_GPU].*/
	DeviceTypeCpu DeviceType = 4
/*Represents the size of the [enum DeviceType] enum.*/
	DeviceTypeMax DeviceType = 5
)
type DriverResource = classdb.RenderingDeviceDriverResource

const (
/*Specific device object based on a physical device.
- Vulkan: Vulkan device driver resource ([code]VkDevice[/code]). ([code]rid[/code] argument doesn't apply.)*/
	DriverResourceLogicalDevice DriverResource = 0
/*Physical device the specific logical device is based on.
- Vulkan: [code]VkDevice[/code]. ([code]rid[/code] argument doesn't apply.)*/
	DriverResourcePhysicalDevice DriverResource = 1
/*Top-most graphics API entry object.
- Vulkan: [code]VkInstance[/code]. ([code]rid[/code] argument doesn't apply.)*/
	DriverResourceTopmostObject DriverResource = 2
/*The main graphics-compute command queue.
- Vulkan: [code]VkQueue[/code]. ([code]rid[/code] argument doesn't apply.)*/
	DriverResourceCommandQueue DriverResource = 3
/*The specific family the main queue belongs to.
- Vulkan: the queue family index, an [code]uint32_t[/code]. ([code]rid[/code] argument doesn't apply.)*/
	DriverResourceQueueFamily DriverResource = 4
/*- Vulkan: [code]VkImage[/code].*/
	DriverResourceTexture DriverResource = 5
/*The view of an owned or shared texture.
- Vulkan: [code]VkImageView[/code].*/
	DriverResourceTextureView DriverResource = 6
/*The native id of the data format of the texture.
- Vulkan: [code]VkFormat[/code].*/
	DriverResourceTextureDataFormat DriverResource = 7
/*- Vulkan: [code]VkSampler[/code].*/
	DriverResourceSampler DriverResource = 8
/*- Vulkan: [code]VkDescriptorSet[/code].*/
	DriverResourceUniformSet DriverResource = 9
/*Buffer of any kind of (storage, vertex, etc.).
- Vulkan: [code]VkBuffer[/code].*/
	DriverResourceBuffer DriverResource = 10
/*- Vulkan: [code]VkPipeline[/code].*/
	DriverResourceComputePipeline DriverResource = 11
/*- Vulkan: [code]VkPipeline[/code].*/
	DriverResourceRenderPipeline DriverResource = 12
	DriverResourceVulkanDevice DriverResource = 0
	DriverResourceVulkanPhysicalDevice DriverResource = 1
	DriverResourceVulkanInstance DriverResource = 2
	DriverResourceVulkanQueue DriverResource = 3
	DriverResourceVulkanQueueFamilyIndex DriverResource = 4
	DriverResourceVulkanImage DriverResource = 5
	DriverResourceVulkanImageView DriverResource = 6
	DriverResourceVulkanImageNativeTextureFormat DriverResource = 7
	DriverResourceVulkanSampler DriverResource = 8
	DriverResourceVulkanDescriptorSet DriverResource = 9
	DriverResourceVulkanBuffer DriverResource = 10
	DriverResourceVulkanComputePipeline DriverResource = 11
	DriverResourceVulkanRenderPipeline DriverResource = 12
)
type DataFormat = classdb.RenderingDeviceDataFormat

const (
/*4-bit-per-channel red/green channel data format, packed into 8 bits. Values are in the [code][0.0, 1.0][/code] range.
[b]Note:[/b] More information on all data formats can be found on the [url=https://registry.khronos.org/vulkan/specs/1.1/html/vkspec.html#_identification_of_formats]Identification of formats[/url] section of the Vulkan specification, as well as the [url=https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html]VkFormat[/url] enum.*/
	DataFormatR4g4UnormPack8 DataFormat = 0
/*4-bit-per-channel red/green/blue/alpha channel data format, packed into 16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR4g4b4a4UnormPack16 DataFormat = 1
/*4-bit-per-channel blue/green/red/alpha channel data format, packed into 16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatB4g4r4a4UnormPack16 DataFormat = 2
/*Red/green/blue channel data format with 5 bits of red, 6 bits of green and 5 bits of blue, packed into 16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR5g6b5UnormPack16 DataFormat = 3
/*Blue/green/red channel data format with 5 bits of blue, 6 bits of green and 5 bits of red, packed into 16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatB5g6r5UnormPack16 DataFormat = 4
/*Red/green/blue/alpha channel data format with 5 bits of red, 6 bits of green, 5 bits of blue and 1 bit of alpha, packed into 16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR5g5b5a1UnormPack16 DataFormat = 5
/*Blue/green/red/alpha channel data format with 5 bits of blue, 6 bits of green, 5 bits of red and 1 bit of alpha, packed into 16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatB5g5r5a1UnormPack16 DataFormat = 6
/*Alpha/red/green/blue channel data format with 1 bit of alpha, 5 bits of red, 6 bits of green and 5 bits of blue, packed into 16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatA1r5g5b5UnormPack16 DataFormat = 7
/*8-bit-per-channel unsigned floating-point red channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR8Unorm DataFormat = 8
/*8-bit-per-channel signed floating-point red channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatR8Snorm DataFormat = 9
/*8-bit-per-channel unsigned floating-point red channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 255.0][/code] range.*/
	DataFormatR8Uscaled DataFormat = 10
/*8-bit-per-channel signed floating-point red channel data format with scaled value (value is converted from integer to float). Values are in the [code][-127.0, 127.0][/code] range.*/
	DataFormatR8Sscaled DataFormat = 11
/*8-bit-per-channel unsigned integer red channel data format. Values are in the [code][0, 255][/code] range.*/
	DataFormatR8Uint DataFormat = 12
/*8-bit-per-channel signed integer red channel data format. Values are in the [code][-127, 127][/code] range.*/
	DataFormatR8Sint DataFormat = 13
/*8-bit-per-channel unsigned floating-point red channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR8Srgb DataFormat = 14
/*8-bit-per-channel unsigned floating-point red/green channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR8g8Unorm DataFormat = 15
/*8-bit-per-channel signed floating-point red/green channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatR8g8Snorm DataFormat = 16
/*8-bit-per-channel unsigned floating-point red/green channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 255.0][/code] range.*/
	DataFormatR8g8Uscaled DataFormat = 17
/*8-bit-per-channel signed floating-point red/green channel data format with scaled value (value is converted from integer to float). Values are in the [code][-127.0, 127.0][/code] range.*/
	DataFormatR8g8Sscaled DataFormat = 18
/*8-bit-per-channel unsigned integer red/green channel data format. Values are in the [code][0, 255][/code] range.*/
	DataFormatR8g8Uint DataFormat = 19
/*8-bit-per-channel signed integer red/green channel data format. Values are in the [code][-127, 127][/code] range.*/
	DataFormatR8g8Sint DataFormat = 20
/*8-bit-per-channel unsigned floating-point red/green channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR8g8Srgb DataFormat = 21
/*8-bit-per-channel unsigned floating-point red/green/blue channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR8g8b8Unorm DataFormat = 22
/*8-bit-per-channel signed floating-point red/green/blue channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatR8g8b8Snorm DataFormat = 23
/*8-bit-per-channel unsigned floating-point red/green/blue channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 255.0][/code] range.*/
	DataFormatR8g8b8Uscaled DataFormat = 24
/*8-bit-per-channel signed floating-point red/green/blue channel data format with scaled value (value is converted from integer to float). Values are in the [code][-127.0, 127.0][/code] range.*/
	DataFormatR8g8b8Sscaled DataFormat = 25
/*8-bit-per-channel unsigned integer red/green/blue channel data format. Values are in the [code][0, 255][/code] range.*/
	DataFormatR8g8b8Uint DataFormat = 26
/*8-bit-per-channel signed integer red/green/blue channel data format. Values are in the [code][-127, 127][/code] range.*/
	DataFormatR8g8b8Sint DataFormat = 27
/*8-bit-per-channel unsigned floating-point red/green/blue/blue channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR8g8b8Srgb DataFormat = 28
/*8-bit-per-channel unsigned floating-point blue/green/red channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatB8g8r8Unorm DataFormat = 29
/*8-bit-per-channel signed floating-point blue/green/red channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatB8g8r8Snorm DataFormat = 30
/*8-bit-per-channel unsigned floating-point blue/green/red channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 255.0][/code] range.*/
	DataFormatB8g8r8Uscaled DataFormat = 31
/*8-bit-per-channel signed floating-point blue/green/red channel data format with scaled value (value is converted from integer to float). Values are in the [code][-127.0, 127.0][/code] range.*/
	DataFormatB8g8r8Sscaled DataFormat = 32
/*8-bit-per-channel unsigned integer blue/green/red channel data format. Values are in the [code][0, 255][/code] range.*/
	DataFormatB8g8r8Uint DataFormat = 33
/*8-bit-per-channel signed integer blue/green/red channel data format. Values are in the [code][-127, 127][/code] range.*/
	DataFormatB8g8r8Sint DataFormat = 34
/*8-bit-per-channel unsigned floating-point blue/green/red data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatB8g8r8Srgb DataFormat = 35
/*8-bit-per-channel unsigned floating-point red/green/blue/alpha channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR8g8b8a8Unorm DataFormat = 36
/*8-bit-per-channel signed floating-point red/green/blue/alpha channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatR8g8b8a8Snorm DataFormat = 37
/*8-bit-per-channel unsigned floating-point red/green/blue/alpha channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 255.0][/code] range.*/
	DataFormatR8g8b8a8Uscaled DataFormat = 38
/*8-bit-per-channel signed floating-point red/green/blue/alpha channel data format with scaled value (value is converted from integer to float). Values are in the [code][-127.0, 127.0][/code] range.*/
	DataFormatR8g8b8a8Sscaled DataFormat = 39
/*8-bit-per-channel unsigned integer red/green/blue/alpha channel data format. Values are in the [code][0, 255][/code] range.*/
	DataFormatR8g8b8a8Uint DataFormat = 40
/*8-bit-per-channel signed integer red/green/blue/alpha channel data format. Values are in the [code][-127, 127][/code] range.*/
	DataFormatR8g8b8a8Sint DataFormat = 41
/*8-bit-per-channel unsigned floating-point red/green/blue/alpha channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR8g8b8a8Srgb DataFormat = 42
/*8-bit-per-channel unsigned floating-point blue/green/red/alpha channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatB8g8r8a8Unorm DataFormat = 43
/*8-bit-per-channel signed floating-point blue/green/red/alpha channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatB8g8r8a8Snorm DataFormat = 44
/*8-bit-per-channel unsigned floating-point blue/green/red/alpha channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 255.0][/code] range.*/
	DataFormatB8g8r8a8Uscaled DataFormat = 45
/*8-bit-per-channel signed floating-point blue/green/red/alpha channel data format with scaled value (value is converted from integer to float). Values are in the [code][-127.0, 127.0][/code] range.*/
	DataFormatB8g8r8a8Sscaled DataFormat = 46
/*8-bit-per-channel unsigned integer blue/green/red/alpha channel data format. Values are in the [code][0, 255][/code] range.*/
	DataFormatB8g8r8a8Uint DataFormat = 47
/*8-bit-per-channel signed integer blue/green/red/alpha channel data format. Values are in the [code][-127, 127][/code] range.*/
	DataFormatB8g8r8a8Sint DataFormat = 48
/*8-bit-per-channel unsigned floating-point blue/green/red/alpha channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatB8g8r8a8Srgb DataFormat = 49
/*8-bit-per-channel unsigned floating-point alpha/red/green/blue channel data format with normalized value, packed in 32 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatA8b8g8r8UnormPack32 DataFormat = 50
/*8-bit-per-channel signed floating-point alpha/red/green/blue channel data format with normalized value, packed in 32 bits. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatA8b8g8r8SnormPack32 DataFormat = 51
/*8-bit-per-channel unsigned floating-point alpha/red/green/blue channel data format with scaled value (value is converted from integer to float), packed in 32 bits. Values are in the [code][0.0, 255.0][/code] range.*/
	DataFormatA8b8g8r8UscaledPack32 DataFormat = 52
/*8-bit-per-channel signed floating-point alpha/red/green/blue channel data format with scaled value (value is converted from integer to float), packed in 32 bits. Values are in the [code][-127.0, 127.0][/code] range.*/
	DataFormatA8b8g8r8SscaledPack32 DataFormat = 53
/*8-bit-per-channel unsigned integer alpha/red/green/blue channel data format, packed in 32 bits. Values are in the [code][0, 255][/code] range.*/
	DataFormatA8b8g8r8UintPack32 DataFormat = 54
/*8-bit-per-channel signed integer alpha/red/green/blue channel data format, packed in 32 bits. Values are in the [code][-127, 127][/code] range.*/
	DataFormatA8b8g8r8SintPack32 DataFormat = 55
/*8-bit-per-channel unsigned floating-point alpha/red/green/blue channel data format with normalized value and non-linear sRGB encoding, packed in 32 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatA8b8g8r8SrgbPack32 DataFormat = 56
/*Unsigned floating-point alpha/red/green/blue channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of red, 10 bits of green and 10 bits of blue. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatA2r10g10b10UnormPack32 DataFormat = 57
/*Signed floating-point alpha/red/green/blue channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of red, 10 bits of green and 10 bits of blue. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatA2r10g10b10SnormPack32 DataFormat = 58
/*Unsigned floating-point alpha/red/green/blue channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of red, 10 bits of green and 10 bits of blue. Values are in the [code][0.0, 1023.0][/code] range for red/green/blue and [code][0.0, 3.0][/code] for alpha.*/
	DataFormatA2r10g10b10UscaledPack32 DataFormat = 59
/*Signed floating-point alpha/red/green/blue channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of red, 10 bits of green and 10 bits of blue. Values are in the [code][-511.0, 511.0][/code] range for red/green/blue and [code][-1.0, 1.0][/code] for alpha.*/
	DataFormatA2r10g10b10SscaledPack32 DataFormat = 60
/*Unsigned integer alpha/red/green/blue channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of red, 10 bits of green and 10 bits of blue. Values are in the [code][0, 1023][/code] range for red/green/blue and [code][0, 3][/code] for alpha.*/
	DataFormatA2r10g10b10UintPack32 DataFormat = 61
/*Signed integer alpha/red/green/blue channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of red, 10 bits of green and 10 bits of blue. Values are in the [code][-511, 511][/code] range for red/green/blue and [code][-1, 1][/code] for alpha.*/
	DataFormatA2r10g10b10SintPack32 DataFormat = 62
/*Unsigned floating-point alpha/blue/green/red channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of blue, 10 bits of green and 10 bits of red. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatA2b10g10r10UnormPack32 DataFormat = 63
/*Signed floating-point alpha/blue/green/red channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of blue, 10 bits of green and 10 bits of red. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatA2b10g10r10SnormPack32 DataFormat = 64
/*Unsigned floating-point alpha/blue/green/red channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of blue, 10 bits of green and 10 bits of red. Values are in the [code][0.0, 1023.0][/code] range for blue/green/red and [code][0.0, 3.0][/code] for alpha.*/
	DataFormatA2b10g10r10UscaledPack32 DataFormat = 65
/*Signed floating-point alpha/blue/green/red channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of blue, 10 bits of green and 10 bits of red. Values are in the [code][-511.0, 511.0][/code] range for blue/green/red and [code][-1.0, 1.0][/code] for alpha.*/
	DataFormatA2b10g10r10SscaledPack32 DataFormat = 66
/*Unsigned integer alpha/blue/green/red channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of blue, 10 bits of green and 10 bits of red. Values are in the [code][0, 1023][/code] range for blue/green/red and [code][0, 3][/code] for alpha.*/
	DataFormatA2b10g10r10UintPack32 DataFormat = 67
/*Signed integer alpha/blue/green/red channel data format with normalized value, packed in 32 bits. Format contains 2 bits of alpha, 10 bits of blue, 10 bits of green and 10 bits of red. Values are in the [code][-511, 511][/code] range for blue/green/red and [code][-1, 1][/code] for alpha.*/
	DataFormatA2b10g10r10SintPack32 DataFormat = 68
/*16-bit-per-channel unsigned floating-point red channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR16Unorm DataFormat = 69
/*16-bit-per-channel signed floating-point red channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatR16Snorm DataFormat = 70
/*16-bit-per-channel unsigned floating-point red channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 65535.0][/code] range.*/
	DataFormatR16Uscaled DataFormat = 71
/*16-bit-per-channel signed floating-point red channel data format with scaled value (value is converted from integer to float). Values are in the [code][-32767.0, 32767.0][/code] range.*/
	DataFormatR16Sscaled DataFormat = 72
/*16-bit-per-channel unsigned integer red channel data format. Values are in the [code][0.0, 65535][/code] range.*/
	DataFormatR16Uint DataFormat = 73
/*16-bit-per-channel signed integer red channel data format. Values are in the [code][-32767, 32767][/code] range.*/
	DataFormatR16Sint DataFormat = 74
/*16-bit-per-channel signed floating-point red channel data format with the value stored as-is.*/
	DataFormatR16Sfloat DataFormat = 75
/*16-bit-per-channel unsigned floating-point red/green channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR16g16Unorm DataFormat = 76
/*16-bit-per-channel signed floating-point red/green channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatR16g16Snorm DataFormat = 77
/*16-bit-per-channel unsigned floating-point red/green channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 65535.0][/code] range.*/
	DataFormatR16g16Uscaled DataFormat = 78
/*16-bit-per-channel signed floating-point red/green channel data format with scaled value (value is converted from integer to float). Values are in the [code][-32767.0, 32767.0][/code] range.*/
	DataFormatR16g16Sscaled DataFormat = 79
/*16-bit-per-channel unsigned integer red/green channel data format. Values are in the [code][0.0, 65535][/code] range.*/
	DataFormatR16g16Uint DataFormat = 80
/*16-bit-per-channel signed integer red/green channel data format. Values are in the [code][-32767, 32767][/code] range.*/
	DataFormatR16g16Sint DataFormat = 81
/*16-bit-per-channel signed floating-point red/green channel data format with the value stored as-is.*/
	DataFormatR16g16Sfloat DataFormat = 82
/*16-bit-per-channel unsigned floating-point red/green/blue channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR16g16b16Unorm DataFormat = 83
/*16-bit-per-channel signed floating-point red/green/blue channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatR16g16b16Snorm DataFormat = 84
/*16-bit-per-channel unsigned floating-point red/green/blue channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 65535.0][/code] range.*/
	DataFormatR16g16b16Uscaled DataFormat = 85
/*16-bit-per-channel signed floating-point red/green/blue channel data format with scaled value (value is converted from integer to float). Values are in the [code][-32767.0, 32767.0][/code] range.*/
	DataFormatR16g16b16Sscaled DataFormat = 86
/*16-bit-per-channel unsigned integer red/green/blue channel data format. Values are in the [code][0.0, 65535][/code] range.*/
	DataFormatR16g16b16Uint DataFormat = 87
/*16-bit-per-channel signed integer red/green/blue channel data format. Values are in the [code][-32767, 32767][/code] range.*/
	DataFormatR16g16b16Sint DataFormat = 88
/*16-bit-per-channel signed floating-point red/green/blue channel data format with the value stored as-is.*/
	DataFormatR16g16b16Sfloat DataFormat = 89
/*16-bit-per-channel unsigned floating-point red/green/blue/alpha channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR16g16b16a16Unorm DataFormat = 90
/*16-bit-per-channel signed floating-point red/green/blue/alpha channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range.*/
	DataFormatR16g16b16a16Snorm DataFormat = 91
/*16-bit-per-channel unsigned floating-point red/green/blue/alpha channel data format with scaled value (value is converted from integer to float). Values are in the [code][0.0, 65535.0][/code] range.*/
	DataFormatR16g16b16a16Uscaled DataFormat = 92
/*16-bit-per-channel signed floating-point red/green/blue/alpha channel data format with scaled value (value is converted from integer to float). Values are in the [code][-32767.0, 32767.0][/code] range.*/
	DataFormatR16g16b16a16Sscaled DataFormat = 93
/*16-bit-per-channel unsigned integer red/green/blue/alpha channel data format. Values are in the [code][0.0, 65535][/code] range.*/
	DataFormatR16g16b16a16Uint DataFormat = 94
/*16-bit-per-channel signed integer red/green/blue/alpha channel data format. Values are in the [code][-32767, 32767][/code] range.*/
	DataFormatR16g16b16a16Sint DataFormat = 95
/*16-bit-per-channel signed floating-point red/green/blue/alpha channel data format with the value stored as-is.*/
	DataFormatR16g16b16a16Sfloat DataFormat = 96
/*32-bit-per-channel unsigned integer red channel data format. Values are in the [code][0, 2^32 - 1][/code] range.*/
	DataFormatR32Uint DataFormat = 97
/*32-bit-per-channel signed integer red channel data format. Values are in the [code][2^31 + 1, 2^31 - 1][/code] range.*/
	DataFormatR32Sint DataFormat = 98
/*32-bit-per-channel signed floating-point red channel data format with the value stored as-is.*/
	DataFormatR32Sfloat DataFormat = 99
/*32-bit-per-channel unsigned integer red/green channel data format. Values are in the [code][0, 2^32 - 1][/code] range.*/
	DataFormatR32g32Uint DataFormat = 100
/*32-bit-per-channel signed integer red/green channel data format. Values are in the [code][2^31 + 1, 2^31 - 1][/code] range.*/
	DataFormatR32g32Sint DataFormat = 101
/*32-bit-per-channel signed floating-point red/green channel data format with the value stored as-is.*/
	DataFormatR32g32Sfloat DataFormat = 102
/*32-bit-per-channel unsigned integer red/green/blue channel data format. Values are in the [code][0, 2^32 - 1][/code] range.*/
	DataFormatR32g32b32Uint DataFormat = 103
/*32-bit-per-channel signed integer red/green/blue channel data format. Values are in the [code][2^31 + 1, 2^31 - 1][/code] range.*/
	DataFormatR32g32b32Sint DataFormat = 104
/*32-bit-per-channel signed floating-point red/green/blue channel data format with the value stored as-is.*/
	DataFormatR32g32b32Sfloat DataFormat = 105
/*32-bit-per-channel unsigned integer red/green/blue/alpha channel data format. Values are in the [code][0, 2^32 - 1][/code] range.*/
	DataFormatR32g32b32a32Uint DataFormat = 106
/*32-bit-per-channel signed integer red/green/blue/alpha channel data format. Values are in the [code][2^31 + 1, 2^31 - 1][/code] range.*/
	DataFormatR32g32b32a32Sint DataFormat = 107
/*32-bit-per-channel signed floating-point red/green/blue/alpha channel data format with the value stored as-is.*/
	DataFormatR32g32b32a32Sfloat DataFormat = 108
/*64-bit-per-channel unsigned integer red channel data format. Values are in the [code][0, 2^64 - 1][/code] range.*/
	DataFormatR64Uint DataFormat = 109
/*64-bit-per-channel signed integer red channel data format. Values are in the [code][2^63 + 1, 2^63 - 1][/code] range.*/
	DataFormatR64Sint DataFormat = 110
/*64-bit-per-channel signed floating-point red channel data format with the value stored as-is.*/
	DataFormatR64Sfloat DataFormat = 111
/*64-bit-per-channel unsigned integer red/green channel data format. Values are in the [code][0, 2^64 - 1][/code] range.*/
	DataFormatR64g64Uint DataFormat = 112
/*64-bit-per-channel signed integer red/green channel data format. Values are in the [code][2^63 + 1, 2^63 - 1][/code] range.*/
	DataFormatR64g64Sint DataFormat = 113
/*64-bit-per-channel signed floating-point red/green channel data format with the value stored as-is.*/
	DataFormatR64g64Sfloat DataFormat = 114
/*64-bit-per-channel unsigned integer red/green/blue channel data format. Values are in the [code][0, 2^64 - 1][/code] range.*/
	DataFormatR64g64b64Uint DataFormat = 115
/*64-bit-per-channel signed integer red/green/blue channel data format. Values are in the [code][2^63 + 1, 2^63 - 1][/code] range.*/
	DataFormatR64g64b64Sint DataFormat = 116
/*64-bit-per-channel signed floating-point red/green/blue channel data format with the value stored as-is.*/
	DataFormatR64g64b64Sfloat DataFormat = 117
/*64-bit-per-channel unsigned integer red/green/blue/alpha channel data format. Values are in the [code][0, 2^64 - 1][/code] range.*/
	DataFormatR64g64b64a64Uint DataFormat = 118
/*64-bit-per-channel signed integer red/green/blue/alpha channel data format. Values are in the [code][2^63 + 1, 2^63 - 1][/code] range.*/
	DataFormatR64g64b64a64Sint DataFormat = 119
/*64-bit-per-channel signed floating-point red/green/blue/alpha channel data format with the value stored as-is.*/
	DataFormatR64g64b64a64Sfloat DataFormat = 120
/*Unsigned floating-point blue/green/red data format with the value stored as-is, packed in 32 bits. The format's precision is 10 bits of blue channel, 11 bits of green channel and 11 bits of red channel.*/
	DataFormatB10g11r11UfloatPack32 DataFormat = 121
/*Unsigned floating-point exposure/blue/green/red data format with the value stored as-is, packed in 32 bits. The format's precision is 5 bits of exposure, 9 bits of blue channel, 9 bits of green channel and 9 bits of red channel.*/
	DataFormatE5b9g9r9UfloatPack32 DataFormat = 122
/*16-bit unsigned floating-point depth data format with normalized value. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatD16Unorm DataFormat = 123
/*24-bit unsigned floating-point depth data format with normalized value, plus 8 unused bits, packed in 32 bits. Values for depth are in the [code][0.0, 1.0][/code] range.*/
	DataFormatX8D24UnormPack32 DataFormat = 124
/*32-bit signed floating-point depth data format with the value stored as-is.*/
	DataFormatD32Sfloat DataFormat = 125
/*8-bit unsigned integer stencil data format.*/
	DataFormatS8Uint DataFormat = 126
/*16-bit unsigned floating-point depth data format with normalized value, plus 8 bits of stencil in unsigned integer format. Values for depth are in the [code][0.0, 1.0][/code] range. Values for stencil are in the [code][0, 255][/code] range.*/
	DataFormatD16UnormS8Uint DataFormat = 127
/*24-bit unsigned floating-point depth data format with normalized value, plus 8 bits of stencil in unsigned integer format. Values for depth are in the [code][0.0, 1.0][/code] range. Values for stencil are in the [code][0, 255][/code] range.*/
	DataFormatD24UnormS8Uint DataFormat = 128
/*32-bit signed floating-point depth data format with the value stored as-is, plus 8 bits of stencil in unsigned integer format. Values for stencil are in the [code][0, 255][/code] range.*/
	DataFormatD32SfloatS8Uint DataFormat = 129
/*VRAM-compressed unsigned red/green/blue channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 5 bits of red channel, 6 bits of green channel and 5 bits of blue channel. Using BC1 texture compression (also known as S3TC DXT1).*/
	DataFormatBc1RgbUnormBlock DataFormat = 130
/*VRAM-compressed unsigned red/green/blue channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 5 bits of red channel, 6 bits of green channel and 5 bits of blue channel. Using BC1 texture compression (also known as S3TC DXT1).*/
	DataFormatBc1RgbSrgbBlock DataFormat = 131
/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 5 bits of red channel, 6 bits of green channel, 5 bits of blue channel and 1 bit of alpha channel. Using BC1 texture compression (also known as S3TC DXT1).*/
	DataFormatBc1RgbaUnormBlock DataFormat = 132
/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 5 bits of red channel, 6 bits of green channel, 5 bits of blue channel and 1 bit of alpha channel. Using BC1 texture compression (also known as S3TC DXT1).*/
	DataFormatBc1RgbaSrgbBlock DataFormat = 133
/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 5 bits of red channel, 6 bits of green channel, 5 bits of blue channel and 4 bits of alpha channel. Using BC2 texture compression (also known as S3TC DXT3).*/
	DataFormatBc2UnormBlock DataFormat = 134
/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 5 bits of red channel, 6 bits of green channel, 5 bits of blue channel and 4 bits of alpha channel. Using BC2 texture compression (also known as S3TC DXT3).*/
	DataFormatBc2SrgbBlock DataFormat = 135
/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 5 bits of red channel, 6 bits of green channel, 5 bits of blue channel and 8 bits of alpha channel. Using BC3 texture compression (also known as S3TC DXT5).*/
	DataFormatBc3UnormBlock DataFormat = 136
/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 5 bits of red channel, 6 bits of green channel, 5 bits of blue channel and 8 bits of alpha channel. Using BC3 texture compression (also known as S3TC DXT5).*/
	DataFormatBc3SrgbBlock DataFormat = 137
/*VRAM-compressed unsigned red channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 8 bits of red channel. Using BC4 texture compression.*/
	DataFormatBc4UnormBlock DataFormat = 138
/*VRAM-compressed signed red channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range. The format's precision is 8 bits of red channel. Using BC4 texture compression.*/
	DataFormatBc4SnormBlock DataFormat = 139
/*VRAM-compressed unsigned red/green channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. The format's precision is 8 bits of red channel and 8 bits of green channel. Using BC5 texture compression (also known as S3TC RGTC).*/
	DataFormatBc5UnormBlock DataFormat = 140
/*VRAM-compressed signed red/green channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range. The format's precision is 8 bits of red channel and 8 bits of green channel. Using BC5 texture compression (also known as S3TC RGTC).*/
	DataFormatBc5SnormBlock DataFormat = 141
/*VRAM-compressed unsigned red/green/blue channel data format with the floating-point value stored as-is. The format's precision is between 10 and 13 bits for the red/green/blue channels. Using BC6H texture compression (also known as BPTC HDR).*/
	DataFormatBc6hUfloatBlock DataFormat = 142
/*VRAM-compressed signed red/green/blue channel data format with the floating-point value stored as-is. The format's precision is between 10 and 13 bits for the red/green/blue channels. Using BC6H texture compression (also known as BPTC HDR).*/
	DataFormatBc6hSfloatBlock DataFormat = 143
/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. The format's precision is between 4 and 7 bits for the red/green/blue channels and between 0 and 8 bits for the alpha channel. Also known as BPTC LDR.*/
	DataFormatBc7UnormBlock DataFormat = 144
/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range. The format's precision is between 4 and 7 bits for the red/green/blue channels and between 0 and 8 bits for the alpha channel. Also known as BPTC LDR.*/
	DataFormatBc7SrgbBlock DataFormat = 145
/*VRAM-compressed unsigned red/green/blue channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. Using ETC2 texture compression.*/
	DataFormatEtc2R8g8b8UnormBlock DataFormat = 146
/*VRAM-compressed unsigned red/green/blue channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range. Using ETC2 texture compression.*/
	DataFormatEtc2R8g8b8SrgbBlock DataFormat = 147
/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. Red/green/blue use 8 bit of precision each, with alpha using 1 bit of precision. Using ETC2 texture compression.*/
	DataFormatEtc2R8g8b8a1UnormBlock DataFormat = 148
/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range. Red/green/blue use 8 bit of precision each, with alpha using 1 bit of precision. Using ETC2 texture compression.*/
	DataFormatEtc2R8g8b8a1SrgbBlock DataFormat = 149
/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. Red/green/blue use 8 bits of precision each, with alpha using 8 bits of precision. Using ETC2 texture compression.*/
	DataFormatEtc2R8g8b8a8UnormBlock DataFormat = 150
/*VRAM-compressed unsigned red/green/blue/alpha channel data format with normalized value and non-linear sRGB encoding. Values are in the [code][0.0, 1.0][/code] range. Red/green/blue use 8 bits of precision each, with alpha using 8 bits of precision. Using ETC2 texture compression.*/
	DataFormatEtc2R8g8b8a8SrgbBlock DataFormat = 151
/*11-bit VRAM-compressed unsigned red channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. Using ETC2 texture compression.*/
	DataFormatEacR11UnormBlock DataFormat = 152
/*11-bit VRAM-compressed signed red channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range. Using ETC2 texture compression.*/
	DataFormatEacR11SnormBlock DataFormat = 153
/*11-bit VRAM-compressed unsigned red/green channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. Using ETC2 texture compression.*/
	DataFormatEacR11g11UnormBlock DataFormat = 154
/*11-bit VRAM-compressed signed red/green channel data format with normalized value. Values are in the [code][-1.0, 1.0][/code] range. Using ETC2 texture compression.*/
	DataFormatEacR11g11SnormBlock DataFormat = 155
/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 4×4 blocks (highest quality). Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc4x4UnormBlock DataFormat = 156
/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 4×4 blocks (highest quality). Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc4x4SrgbBlock DataFormat = 157
/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 5×4 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc5x4UnormBlock DataFormat = 158
/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 5×4 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc5x4SrgbBlock DataFormat = 159
/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 5×5 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc5x5UnormBlock DataFormat = 160
/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 5×5 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc5x5SrgbBlock DataFormat = 161
/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 6×5 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc6x5UnormBlock DataFormat = 162
/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 6×5 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc6x5SrgbBlock DataFormat = 163
/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 6×6 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc6x6UnormBlock DataFormat = 164
/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 6×6 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc6x6SrgbBlock DataFormat = 165
/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 8×5 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc8x5UnormBlock DataFormat = 166
/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 8×5 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc8x5SrgbBlock DataFormat = 167
/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 8×6 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc8x6UnormBlock DataFormat = 168
/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 8×6 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc8x6SrgbBlock DataFormat = 169
/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 8×8 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc8x8UnormBlock DataFormat = 170
/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 8×8 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc8x8SrgbBlock DataFormat = 171
/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 10×5 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc10x5UnormBlock DataFormat = 172
/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 10×5 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc10x5SrgbBlock DataFormat = 173
/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 10×6 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc10x6UnormBlock DataFormat = 174
/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 10×6 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc10x6SrgbBlock DataFormat = 175
/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 10×8 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc10x8UnormBlock DataFormat = 176
/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 10×8 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc10x8SrgbBlock DataFormat = 177
/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 10×10 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc10x10UnormBlock DataFormat = 178
/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 10×10 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc10x10SrgbBlock DataFormat = 179
/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 12×10 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc12x10UnormBlock DataFormat = 180
/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 12×10 blocks. Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc12x10SrgbBlock DataFormat = 181
/*VRAM-compressed unsigned floating-point data format with normalized value, packed in 12 blocks (lowest quality). Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc12x12UnormBlock DataFormat = 182
/*VRAM-compressed unsigned floating-point data format with normalized value and non-linear sRGB encoding, packed in 12 blocks (lowest quality). Values are in the [code][0.0, 1.0][/code] range. Using ASTC compression.*/
	DataFormatAstc12x12SrgbBlock DataFormat = 183
/*8-bit-per-channel unsigned floating-point green/blue/red channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG8b8g8r8422Unorm DataFormat = 184
/*8-bit-per-channel unsigned floating-point blue/green/red channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatB8g8r8g8422Unorm DataFormat = 185
/*8-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, stored across 3 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal and vertical resolution (i.e. 2×2 adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG8B8R83plane420Unorm DataFormat = 186
/*8-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, stored across 2 separate planes (green + blue/red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal and vertical resolution (i.e. 2×2 adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG8B8r82plane420Unorm DataFormat = 187
/*8-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, stored across 2 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG8B8R83plane422Unorm DataFormat = 188
/*8-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, stored across 2 separate planes (green + blue/red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG8B8r82plane422Unorm DataFormat = 189
/*8-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, stored across 3 separate planes. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatG8B8R83plane444Unorm DataFormat = 190
/*10-bit-per-channel unsigned floating-point red channel data with normalized value, plus 6 unused bits, packed in 16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR10x6UnormPack16 DataFormat = 191
/*10-bit-per-channel unsigned floating-point red/green channel data with normalized value, plus 6 unused bits after each channel, packed in 2×16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR10x6g10x6Unorm2pack16 DataFormat = 192
/*10-bit-per-channel unsigned floating-point red/green/blue/alpha channel data with normalized value, plus 6 unused bits after each channel, packed in 4×16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR10x6g10x6b10x6a10x6Unorm4pack16 DataFormat = 193
/*10-bit-per-channel unsigned floating-point green/blue/green/red channel data with normalized value, plus 6 unused bits after each channel, packed in 4×16 bits. Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel). The green channel is listed twice, but contains different values to allow it to be represented at full resolution.*/
	DataFormatG10x6b10x6g10x6r10x6422Unorm4pack16 DataFormat = 194
/*10-bit-per-channel unsigned floating-point blue/green/red/green channel data with normalized value, plus 6 unused bits after each channel, packed in 4×16 bits. Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel). The green channel is listed twice, but contains different values to allow it to be represented at full resolution.*/
	DataFormatB10x6g10x6r10x6g10x6422Unorm4pack16 DataFormat = 195
/*10-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 2 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal and vertical resolution (i.e. 2×2 adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG10x6B10x6R10x63plane420Unorm3pack16 DataFormat = 196
/*10-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 2 separate planes (green + blue/red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal and vertical resolution (i.e. 2×2 adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG10x6B10x6r10x62plane420Unorm3pack16 DataFormat = 197
/*10-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 3 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG10x6B10x6R10x63plane422Unorm3pack16 DataFormat = 198
/*10-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 3 separate planes (green + blue/red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG10x6B10x6r10x62plane422Unorm3pack16 DataFormat = 199
/*10-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 3 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatG10x6B10x6R10x63plane444Unorm3pack16 DataFormat = 200
/*12-bit-per-channel unsigned floating-point red channel data with normalized value, plus 6 unused bits, packed in 16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR12x4UnormPack16 DataFormat = 201
/*12-bit-per-channel unsigned floating-point red/green channel data with normalized value, plus 6 unused bits after each channel, packed in 2×16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR12x4g12x4Unorm2pack16 DataFormat = 202
/*12-bit-per-channel unsigned floating-point red/green/blue/alpha channel data with normalized value, plus 6 unused bits after each channel, packed in 4×16 bits. Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatR12x4g12x4b12x4a12x4Unorm4pack16 DataFormat = 203
/*12-bit-per-channel unsigned floating-point green/blue/green/red channel data with normalized value, plus 6 unused bits after each channel, packed in 4×16 bits. Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel). The green channel is listed twice, but contains different values to allow it to be represented at full resolution.*/
	DataFormatG12x4b12x4g12x4r12x4422Unorm4pack16 DataFormat = 204
/*12-bit-per-channel unsigned floating-point blue/green/red/green channel data with normalized value, plus 6 unused bits after each channel, packed in 4×16 bits. Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel). The green channel is listed twice, but contains different values to allow it to be represented at full resolution.*/
	DataFormatB12x4g12x4r12x4g12x4422Unorm4pack16 DataFormat = 205
/*12-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 2 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal and vertical resolution (i.e. 2×2 adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG12x4B12x4R12x43plane420Unorm3pack16 DataFormat = 206
/*12-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 2 separate planes (green + blue/red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal and vertical resolution (i.e. 2×2 adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG12x4B12x4r12x42plane420Unorm3pack16 DataFormat = 207
/*12-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 3 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG12x4B12x4R12x43plane422Unorm3pack16 DataFormat = 208
/*12-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 3 separate planes (green + blue/red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG12x4B12x4r12x42plane422Unorm3pack16 DataFormat = 209
/*12-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Packed in 3×16 bits and stored across 3 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatG12x4B12x4R12x43plane444Unorm3pack16 DataFormat = 210
/*16-bit-per-channel unsigned floating-point green/blue/red channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG16b16g16r16422Unorm DataFormat = 211
/*16-bit-per-channel unsigned floating-point blue/green/red channel data format with normalized value. Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatB16g16r16g16422Unorm DataFormat = 212
/*16-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Stored across 2 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal and vertical resolution (i.e. 2×2 adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG16B16R163plane420Unorm DataFormat = 213
/*16-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Stored across 2 separate planes (green + blue/red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal and vertical resolution (i.e. 2×2 adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG16B16r162plane420Unorm DataFormat = 214
/*16-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Stored across 3 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG16B16R163plane422Unorm DataFormat = 215
/*16-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Stored across 3 separate planes (green + blue/red). Values are in the [code][0.0, 1.0][/code] range. Blue and red channel data is stored at halved horizontal resolution (i.e. 2 horizontally adjacent pixels will share the same value for the blue/red channel).*/
	DataFormatG16B16r162plane422Unorm DataFormat = 216
/*16-bit-per-channel unsigned floating-point green/blue/red channel data with normalized value, plus 6 unused bits after each channel. Stored across 3 separate planes (green + blue + red). Values are in the [code][0.0, 1.0][/code] range.*/
	DataFormatG16B16R163plane444Unorm DataFormat = 217
/*Represents the size of the [enum DataFormat] enum.*/
	DataFormatMax DataFormat = 218
)
type BarrierMask = classdb.RenderingDeviceBarrierMask

const (
/*Vertex shader barrier mask.*/
	BarrierMaskVertex BarrierMask = 1
/*Fragment shader barrier mask.*/
	BarrierMaskFragment BarrierMask = 8
/*Compute barrier mask.*/
	BarrierMaskCompute BarrierMask = 2
/*Transfer barrier mask.*/
	BarrierMaskTransfer BarrierMask = 4
/*Raster barrier mask (vertex and fragment). Equivalent to [code]BARRIER_MASK_VERTEX | BARRIER_MASK_FRAGMENT[/code].*/
	BarrierMaskRaster BarrierMask = 9
/*Barrier mask for all types (vertex, fragment, compute, transfer).*/
	BarrierMaskAllBarriers BarrierMask = 32767
/*No barrier for any type.*/
	BarrierMaskNoBarrier BarrierMask = 32768
)
type TextureType = classdb.RenderingDeviceTextureType

const (
/*1-dimensional texture.*/
	TextureType1d TextureType = 0
/*2-dimensional texture.*/
	TextureType2d TextureType = 1
/*3-dimensional texture.*/
	TextureType3d TextureType = 2
/*[Cubemap] texture.*/
	TextureTypeCube TextureType = 3
/*Array of 1-dimensional textures.*/
	TextureType1dArray TextureType = 4
/*Array of 2-dimensional textures.*/
	TextureType2dArray TextureType = 5
/*Array of [Cubemap] textures.*/
	TextureTypeCubeArray TextureType = 6
/*Represents the size of the [enum TextureType] enum.*/
	TextureTypeMax TextureType = 7
)
type TextureSamples = classdb.RenderingDeviceTextureSamples

const (
/*Perform 1 texture sample (this is the fastest but lowest-quality for antialiasing).*/
	TextureSamples1 TextureSamples = 0
/*Perform 2 texture samples.*/
	TextureSamples2 TextureSamples = 1
/*Perform 4 texture samples.*/
	TextureSamples4 TextureSamples = 2
/*Perform 8 texture samples. Not supported on mobile GPUs (including Apple Silicon).*/
	TextureSamples8 TextureSamples = 3
/*Perform 16 texture samples. Not supported on mobile GPUs and many desktop GPUs.*/
	TextureSamples16 TextureSamples = 4
/*Perform 32 texture samples. Not supported on most GPUs.*/
	TextureSamples32 TextureSamples = 5
/*Perform 64 texture samples (this is the slowest but highest-quality for antialiasing). Not supported on most GPUs.*/
	TextureSamples64 TextureSamples = 6
/*Represents the size of the [enum TextureSamples] enum.*/
	TextureSamplesMax TextureSamples = 7
)
type TextureUsageBits = classdb.RenderingDeviceTextureUsageBits

const (
/*Texture can be sampled.*/
	TextureUsageSamplingBit TextureUsageBits = 1
/*Texture can be used as a color attachment in a framebuffer.*/
	TextureUsageColorAttachmentBit TextureUsageBits = 2
/*Texture can be used as a depth/stencil attachment in a framebuffer.*/
	TextureUsageDepthStencilAttachmentBit TextureUsageBits = 4
/*Texture can be used as a [url=https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storageimage]storage image[/url].*/
	TextureUsageStorageBit TextureUsageBits = 8
/*Texture can be used as a [url=https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storageimage]storage image[/url] with support for atomic operations.*/
	TextureUsageStorageAtomicBit TextureUsageBits = 16
/*Texture can be read back on the CPU using [method texture_get_data] faster than without this bit, since it is always kept in the system memory.*/
	TextureUsageCpuReadBit TextureUsageBits = 32
/*Texture can be updated using [method texture_update].*/
	TextureUsageCanUpdateBit TextureUsageBits = 64
/*Texture can be a source for [method texture_copy].*/
	TextureUsageCanCopyFromBit TextureUsageBits = 128
/*Texture can be a destination for [method texture_copy].*/
	TextureUsageCanCopyToBit TextureUsageBits = 256
/*Texture can be used as a [url=https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-inputattachment]input attachment[/url] in a framebuffer.*/
	TextureUsageInputAttachmentBit TextureUsageBits = 512
)
type TextureSwizzle = classdb.RenderingDeviceTextureSwizzle

const (
/*Return the sampled value as-is.*/
	TextureSwizzleIdentity TextureSwizzle = 0
/*Always return [code]0.0[/code] when sampling.*/
	TextureSwizzleZero TextureSwizzle = 1
/*Always return [code]1.0[/code] when sampling.*/
	TextureSwizzleOne TextureSwizzle = 2
/*Sample the red color channel.*/
	TextureSwizzleR TextureSwizzle = 3
/*Sample the green color channel.*/
	TextureSwizzleG TextureSwizzle = 4
/*Sample the blue color channel.*/
	TextureSwizzleB TextureSwizzle = 5
/*Sample the alpha channel.*/
	TextureSwizzleA TextureSwizzle = 6
/*Represents the size of the [enum TextureSwizzle] enum.*/
	TextureSwizzleMax TextureSwizzle = 7
)
type TextureSliceType = classdb.RenderingDeviceTextureSliceType

const (
/*2-dimensional texture slice.*/
	TextureSlice2d TextureSliceType = 0
/*Cubemap texture slice.*/
	TextureSliceCubemap TextureSliceType = 1
/*3-dimensional texture slice.*/
	TextureSlice3d TextureSliceType = 2
)
type SamplerFilter = classdb.RenderingDeviceSamplerFilter

const (
/*Nearest-neighbor sampler filtering. Sampling at higher resolutions than the source will result in a pixelated look.*/
	SamplerFilterNearest SamplerFilter = 0
/*Bilinear sampler filtering. Sampling at higher resolutions than the source will result in a blurry look.*/
	SamplerFilterLinear SamplerFilter = 1
)
type SamplerRepeatMode = classdb.RenderingDeviceSamplerRepeatMode

const (
/*Sample with repeating enabled.*/
	SamplerRepeatModeRepeat SamplerRepeatMode = 0
/*Sample with mirrored repeating enabled. When sampling outside the [code][0.0, 1.0][/code] range, return a mirrored version of the sampler. This mirrored version is mirrored again if sampling further away, with the pattern repeating indefinitely.*/
	SamplerRepeatModeMirroredRepeat SamplerRepeatMode = 1
/*Sample with repeating disabled. When sampling outside the [code][0.0, 1.0][/code] range, return the color of the last pixel on the edge.*/
	SamplerRepeatModeClampToEdge SamplerRepeatMode = 2
/*Sample with repeating disabled. When sampling outside the [code][0.0, 1.0][/code] range, return the specified [member RDSamplerState.border_color].*/
	SamplerRepeatModeClampToBorder SamplerRepeatMode = 3
/*Sample with mirrored repeating enabled, but only once. When sampling in the [code][-1.0, 0.0][/code] range, return a mirrored version of the sampler. When sampling outside the [code][-1.0, 1.0][/code] range, return the color of the last pixel on the edge.*/
	SamplerRepeatModeMirrorClampToEdge SamplerRepeatMode = 4
/*Represents the size of the [enum SamplerRepeatMode] enum.*/
	SamplerRepeatModeMax SamplerRepeatMode = 5
)
type SamplerBorderColor = classdb.RenderingDeviceSamplerBorderColor

const (
/*Return a floating-point transparent black color when sampling outside the [code][0.0, 1.0][/code] range. Only effective if the sampler repeat mode is [constant SAMPLER_REPEAT_MODE_CLAMP_TO_BORDER].*/
	SamplerBorderColorFloatTransparentBlack SamplerBorderColor = 0
/*Return a integer transparent black color when sampling outside the [code][0.0, 1.0][/code] range. Only effective if the sampler repeat mode is [constant SAMPLER_REPEAT_MODE_CLAMP_TO_BORDER].*/
	SamplerBorderColorIntTransparentBlack SamplerBorderColor = 1
/*Return a floating-point opaque black color when sampling outside the [code][0.0, 1.0][/code] range. Only effective if the sampler repeat mode is [constant SAMPLER_REPEAT_MODE_CLAMP_TO_BORDER].*/
	SamplerBorderColorFloatOpaqueBlack SamplerBorderColor = 2
/*Return a integer opaque black color when sampling outside the [code][0.0, 1.0][/code] range. Only effective if the sampler repeat mode is [constant SAMPLER_REPEAT_MODE_CLAMP_TO_BORDER].*/
	SamplerBorderColorIntOpaqueBlack SamplerBorderColor = 3
/*Return a floating-point opaque white color when sampling outside the [code][0.0, 1.0][/code] range. Only effective if the sampler repeat mode is [constant SAMPLER_REPEAT_MODE_CLAMP_TO_BORDER].*/
	SamplerBorderColorFloatOpaqueWhite SamplerBorderColor = 4
/*Return a integer opaque white color when sampling outside the [code][0.0, 1.0][/code] range. Only effective if the sampler repeat mode is [constant SAMPLER_REPEAT_MODE_CLAMP_TO_BORDER].*/
	SamplerBorderColorIntOpaqueWhite SamplerBorderColor = 5
/*Represents the size of the [enum SamplerBorderColor] enum.*/
	SamplerBorderColorMax SamplerBorderColor = 6
)
type VertexFrequency = classdb.RenderingDeviceVertexFrequency

const (
/*Vertex attribute addressing is a function of the vertex. This is used to specify the rate at which vertex attributes are pulled from buffers.*/
	VertexFrequencyVertex VertexFrequency = 0
/*Vertex attribute addressing is a function of the instance index. This is used to specify the rate at which vertex attributes are pulled from buffers.*/
	VertexFrequencyInstance VertexFrequency = 1
)
type IndexBufferFormat = classdb.RenderingDeviceIndexBufferFormat

const (
/*Index buffer in 16-bit unsigned integer format. This limits the maximum index that can be specified to [code]65535[/code].*/
	IndexBufferFormatUint16 IndexBufferFormat = 0
/*Index buffer in 32-bit unsigned integer format. This limits the maximum index that can be specified to [code]4294967295[/code].*/
	IndexBufferFormatUint32 IndexBufferFormat = 1
)
type StorageBufferUsage = classdb.RenderingDeviceStorageBufferUsage

const (
	StorageBufferUsageDispatchIndirect StorageBufferUsage = 1
)
type UniformType = classdb.RenderingDeviceUniformType

const (
/*Sampler uniform.*/
	UniformTypeSampler UniformType = 0
/*Sampler uniform with a texture.*/
	UniformTypeSamplerWithTexture UniformType = 1
/*Texture uniform.*/
	UniformTypeTexture UniformType = 2
/*Image uniform.*/
	UniformTypeImage UniformType = 3
/*Texture buffer uniform.*/
	UniformTypeTextureBuffer UniformType = 4
/*Sampler uniform with a texture buffer.*/
	UniformTypeSamplerWithTextureBuffer UniformType = 5
/*Image buffer uniform.*/
	UniformTypeImageBuffer UniformType = 6
/*Uniform buffer uniform.*/
	UniformTypeUniformBuffer UniformType = 7
/*[url=https://vkguide.dev/docs/chapter-4/storage_buffers/]Storage buffer[/url] uniform.*/
	UniformTypeStorageBuffer UniformType = 8
/*Input attachment uniform.*/
	UniformTypeInputAttachment UniformType = 9
/*Represents the size of the [enum UniformType] enum.*/
	UniformTypeMax UniformType = 10
)
type RenderPrimitive = classdb.RenderingDeviceRenderPrimitive

const (
/*Point rendering primitive (with constant size, regardless of distance from camera).*/
	RenderPrimitivePoints RenderPrimitive = 0
/*Line list rendering primitive. Lines are drawn separated from each other.*/
	RenderPrimitiveLines RenderPrimitive = 1
/*[url=https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-line-lists-with-adjacency]Line list rendering primitive with adjacency.[/url]
[b]Note:[/b] Adjacency is only useful with geometry shaders, which Godot does not expose.*/
	RenderPrimitiveLinesWithAdjacency RenderPrimitive = 2
/*Line strip rendering primitive. Lines drawn are connected to the previous vertex.*/
	RenderPrimitiveLinestrips RenderPrimitive = 3
/*[url=https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-line-strips-with-adjacency]Line strip rendering primitive with adjacency.[/url]
[b]Note:[/b] Adjacency is only useful with geometry shaders, which Godot does not expose.*/
	RenderPrimitiveLinestripsWithAdjacency RenderPrimitive = 4
/*Triangle list rendering primitive. Triangles are drawn separated from each other.*/
	RenderPrimitiveTriangles RenderPrimitive = 5
/*[url=https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-triangle-lists-with-adjacency]Triangle list rendering primitive with adjacency.[/url]
 [b]Note:[/b] Adjacency is only useful with geometry shaders, which Godot does not expose.*/
	RenderPrimitiveTrianglesWithAdjacency RenderPrimitive = 6
/*Triangle strip rendering primitive. Triangles drawn are connected to the previous triangle.*/
	RenderPrimitiveTriangleStrips RenderPrimitive = 7
/*[url=https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-triangle-strips-with-adjacency]Triangle strip rendering primitive with adjacency.[/url]
[b]Note:[/b] Adjacency is only useful with geometry shaders, which Godot does not expose.*/
	RenderPrimitiveTriangleStripsWithAjacency RenderPrimitive = 8
/*Triangle strip rendering primitive with [i]primitive restart[/i] enabled. Triangles drawn are connected to the previous triangle, but a primitive restart index can be specified before drawing to create a second triangle strip after the specified index.
[b]Note:[/b] Only compatible with indexed draws.*/
	RenderPrimitiveTriangleStripsWithRestartIndex RenderPrimitive = 9
/*Tessellation patch rendering primitive. Only useful with tessellation shaders, which can be used to deform these patches.*/
	RenderPrimitiveTesselationPatch RenderPrimitive = 10
/*Represents the size of the [enum RenderPrimitive] enum.*/
	RenderPrimitiveMax RenderPrimitive = 11
)
type PolygonCullMode = classdb.RenderingDevicePolygonCullMode

const (
/*Do not use polygon front face or backface culling.*/
	PolygonCullDisabled PolygonCullMode = 0
/*Use polygon frontface culling (faces pointing towards the camera are hidden).*/
	PolygonCullFront PolygonCullMode = 1
/*Use polygon backface culling (faces pointing away from the camera are hidden).*/
	PolygonCullBack PolygonCullMode = 2
)
type PolygonFrontFace = classdb.RenderingDevicePolygonFrontFace

const (
/*Clockwise winding order to determine which face of a polygon is its front face.*/
	PolygonFrontFaceClockwise PolygonFrontFace = 0
/*Counter-clockwise winding order to determine which face of a polygon is its front face.*/
	PolygonFrontFaceCounterClockwise PolygonFrontFace = 1
)
type StencilOperation = classdb.RenderingDeviceStencilOperation

const (
/*Keep the current stencil value.*/
	StencilOpKeep StencilOperation = 0
/*Set the stencil value to [code]0[/code].*/
	StencilOpZero StencilOperation = 1
/*Replace the existing stencil value with the new one.*/
	StencilOpReplace StencilOperation = 2
/*Increment the existing stencil value and clamp to the maximum representable unsigned value if reached. Stencil bits are considered as an unsigned integer.*/
	StencilOpIncrementAndClamp StencilOperation = 3
/*Decrement the existing stencil value and clamp to the minimum value if reached. Stencil bits are considered as an unsigned integer.*/
	StencilOpDecrementAndClamp StencilOperation = 4
/*Bitwise-invert the existing stencil value.*/
	StencilOpInvert StencilOperation = 5
/*Increment the stencil value and wrap around to [code]0[/code] if reaching the maximum representable unsigned. Stencil bits are considered as an unsigned integer.*/
	StencilOpIncrementAndWrap StencilOperation = 6
/*Decrement the stencil value and wrap around to the maximum representable unsigned if reaching the minimum. Stencil bits are considered as an unsigned integer.*/
	StencilOpDecrementAndWrap StencilOperation = 7
/*Represents the size of the [enum StencilOperation] enum.*/
	StencilOpMax StencilOperation = 8
)
type CompareOperator = classdb.RenderingDeviceCompareOperator

const (
/*"Never" comparison (opposite of [constant COMPARE_OP_ALWAYS]).*/
	CompareOpNever CompareOperator = 0
/*"Less than" comparison.*/
	CompareOpLess CompareOperator = 1
/*"Equal" comparison.*/
	CompareOpEqual CompareOperator = 2
/*"Less than or equal" comparison.*/
	CompareOpLessOrEqual CompareOperator = 3
/*"Greater than" comparison.*/
	CompareOpGreater CompareOperator = 4
/*"Not equal" comparison.*/
	CompareOpNotEqual CompareOperator = 5
/*"Greater than or equal" comparison.*/
	CompareOpGreaterOrEqual CompareOperator = 6
/*"Always" comparison (opposite of [constant COMPARE_OP_NEVER]).*/
	CompareOpAlways CompareOperator = 7
/*Represents the size of the [enum CompareOperator] enum.*/
	CompareOpMax CompareOperator = 8
)
type LogicOperation = classdb.RenderingDeviceLogicOperation

const (
/*Clear logic operation (result is always [code]0[/code]). See also [constant LOGIC_OP_SET].*/
	LogicOpClear LogicOperation = 0
/*AND logic operation.*/
	LogicOpAnd LogicOperation = 1
/*AND logic operation with the [i]destination[/i] operand being inverted. See also [constant LOGIC_OP_AND_INVERTED].*/
	LogicOpAndReverse LogicOperation = 2
/*Copy logic operation (keeps the [i]source[/i] value as-is). See also [constant LOGIC_OP_COPY_INVERTED] and [constant LOGIC_OP_NO_OP].*/
	LogicOpCopy LogicOperation = 3
/*AND logic operation with the [i]source[/i] operand being inverted. See also [constant LOGIC_OP_AND_REVERSE].*/
	LogicOpAndInverted LogicOperation = 4
/*No-op logic operation (keeps the [i]destination[/i] value as-is). See also [constant LOGIC_OP_COPY].*/
	LogicOpNoOp LogicOperation = 5
/*Exclusive or (XOR) logic operation.*/
	LogicOpXor LogicOperation = 6
/*OR logic operation.*/
	LogicOpOr LogicOperation = 7
/*Not-OR (NOR) logic operation.*/
	LogicOpNor LogicOperation = 8
/*Not-XOR (XNOR) logic operation.*/
	LogicOpEquivalent LogicOperation = 9
/*Invert logic operation.*/
	LogicOpInvert LogicOperation = 10
/*OR logic operation with the [i]destination[/i] operand being inverted. See also [constant LOGIC_OP_OR_REVERSE].*/
	LogicOpOrReverse LogicOperation = 11
/*NOT logic operation (inverts the value). See also [constant LOGIC_OP_COPY].*/
	LogicOpCopyInverted LogicOperation = 12
/*OR logic operation with the [i]source[/i] operand being inverted. See also [constant LOGIC_OP_OR_REVERSE].*/
	LogicOpOrInverted LogicOperation = 13
/*Not-AND (NAND) logic operation.*/
	LogicOpNand LogicOperation = 14
/*SET logic operation (result is always [code]1[/code]). See also [constant LOGIC_OP_CLEAR].*/
	LogicOpSet LogicOperation = 15
/*Represents the size of the [enum LogicOperation] enum.*/
	LogicOpMax LogicOperation = 16
)
type BlendFactor = classdb.RenderingDeviceBlendFactor

const (
/*Constant [code]0.0[/code] blend factor.*/
	BlendFactorZero BlendFactor = 0
/*Constant [code]1.0[/code] blend factor.*/
	BlendFactorOne BlendFactor = 1
/*Color blend factor is [code]source color[/code]. Alpha blend factor is [code]source alpha[/code].*/
	BlendFactorSrcColor BlendFactor = 2
/*Color blend factor is [code]1.0 - source color[/code]. Alpha blend factor is [code]1.0 - source alpha[/code].*/
	BlendFactorOneMinusSrcColor BlendFactor = 3
/*Color blend factor is [code]destination color[/code]. Alpha blend factor is [code]destination alpha[/code].*/
	BlendFactorDstColor BlendFactor = 4
/*Color blend factor is [code]1.0 - destination color[/code]. Alpha blend factor is [code]1.0 - destination alpha[/code].*/
	BlendFactorOneMinusDstColor BlendFactor = 5
/*Color and alpha blend factor is [code]source alpha[/code].*/
	BlendFactorSrcAlpha BlendFactor = 6
/*Color and alpha blend factor is [code]1.0 - source alpha[/code].*/
	BlendFactorOneMinusSrcAlpha BlendFactor = 7
/*Color and alpha blend factor is [code]destination alpha[/code].*/
	BlendFactorDstAlpha BlendFactor = 8
/*Color and alpha blend factor is [code]1.0 - destination alpha[/code].*/
	BlendFactorOneMinusDstAlpha BlendFactor = 9
/*Color blend factor is [code]blend constant color[/code]. Alpha blend factor is [code]blend constant alpha[/code] (see [method draw_list_set_blend_constants]).*/
	BlendFactorConstantColor BlendFactor = 10
/*Color blend factor is [code]1.0 - blend constant color[/code]. Alpha blend factor is [code]1.0 - blend constant alpha[/code] (see [method draw_list_set_blend_constants]).*/
	BlendFactorOneMinusConstantColor BlendFactor = 11
/*Color and alpha blend factor is [code]blend constant alpha[/code] (see [method draw_list_set_blend_constants]).*/
	BlendFactorConstantAlpha BlendFactor = 12
/*Color and alpha blend factor is [code]1.0 - blend constant alpha[/code] (see [method draw_list_set_blend_constants]).*/
	BlendFactorOneMinusConstantAlpha BlendFactor = 13
/*Color blend factor is [code]min(source alpha, 1.0 - destination alpha)[/code]. Alpha blend factor is [code]1.0[/code].*/
	BlendFactorSrcAlphaSaturate BlendFactor = 14
/*Color blend factor is [code]second source color[/code]. Alpha blend factor is [code]second source alpha[/code]. Only relevant for dual-source blending.*/
	BlendFactorSrc1Color BlendFactor = 15
/*Color blend factor is [code]1.0 - second source color[/code]. Alpha blend factor is [code]1.0 - second source alpha[/code]. Only relevant for dual-source blending.*/
	BlendFactorOneMinusSrc1Color BlendFactor = 16
/*Color and alpha blend factor is [code]second source alpha[/code]. Only relevant for dual-source blending.*/
	BlendFactorSrc1Alpha BlendFactor = 17
/*Color and alpha blend factor is [code]1.0 - second source alpha[/code]. Only relevant for dual-source blending.*/
	BlendFactorOneMinusSrc1Alpha BlendFactor = 18
/*Represents the size of the [enum BlendFactor] enum.*/
	BlendFactorMax BlendFactor = 19
)
type BlendOperation = classdb.RenderingDeviceBlendOperation

const (
/*Additive blending operation ([code]source + destination[/code]).*/
	BlendOpAdd BlendOperation = 0
/*Subtractive blending operation ([code]source - destination[/code]).*/
	BlendOpSubtract BlendOperation = 1
/*Reverse subtractive blending operation ([code]destination - source[/code]).*/
	BlendOpReverseSubtract BlendOperation = 2
/*Minimum blending operation (keep the lowest value of the two).*/
	BlendOpMinimum BlendOperation = 3
/*Maximum blending operation (keep the highest value of the two).*/
	BlendOpMaximum BlendOperation = 4
/*Represents the size of the [enum BlendOperation] enum.*/
	BlendOpMax BlendOperation = 5
)
type PipelineDynamicStateFlags = classdb.RenderingDevicePipelineDynamicStateFlags

const (
/*Allows dynamically changing the width of rendering lines.*/
	DynamicStateLineWidth PipelineDynamicStateFlags = 1
/*Allows dynamically changing the depth bias.*/
	DynamicStateDepthBias PipelineDynamicStateFlags = 2
	DynamicStateBlendConstants PipelineDynamicStateFlags = 4
	DynamicStateDepthBounds PipelineDynamicStateFlags = 8
	DynamicStateStencilCompareMask PipelineDynamicStateFlags = 16
	DynamicStateStencilWriteMask PipelineDynamicStateFlags = 32
	DynamicStateStencilReference PipelineDynamicStateFlags = 64
)
type InitialAction = classdb.RenderingDeviceInitialAction

const (
/*Load the previous contents of the framebuffer.*/
	InitialActionLoad InitialAction = 0
/*Clear the whole framebuffer or its specified region.*/
	InitialActionClear InitialAction = 1
/*Ignore the previous contents of the framebuffer. This is the fastest option if you'll overwrite all of the pixels and don't need to read any of them.*/
	InitialActionDiscard InitialAction = 2
/*Represents the size of the [enum InitialAction] enum.*/
	InitialActionMax InitialAction = 3
	InitialActionClearRegion InitialAction = 1
	InitialActionClearRegionContinue InitialAction = 1
	InitialActionKeep InitialAction = 0
	InitialActionDrop InitialAction = 2
	InitialActionContinue InitialAction = 0
)
type FinalAction = classdb.RenderingDeviceFinalAction

const (
/*Store the result of the draw list in the framebuffer. This is generally what you want to do.*/
	FinalActionStore FinalAction = 0
/*Discard the contents of the framebuffer. This is the fastest option if you don't need to use the results of the draw list.*/
	FinalActionDiscard FinalAction = 1
/*Represents the size of the [enum FinalAction] enum.*/
	FinalActionMax FinalAction = 2
	FinalActionRead FinalAction = 0
	FinalActionContinue FinalAction = 0
)
type ShaderStage = classdb.RenderingDeviceShaderStage

const (
/*Vertex shader stage. This can be used to manipulate vertices from a shader (but not create new vertices).*/
	ShaderStageVertex ShaderStage = 0
/*Fragment shader stage (called "pixel shader" in Direct3D). This can be used to manipulate pixels from a shader.*/
	ShaderStageFragment ShaderStage = 1
/*Tessellation control shader stage. This can be used to create additional geometry from a shader.*/
	ShaderStageTesselationControl ShaderStage = 2
/*Tessellation evaluation shader stage. This can be used to create additional geometry from a shader.*/
	ShaderStageTesselationEvaluation ShaderStage = 3
/*Compute shader stage. This can be used to run arbitrary computing tasks in a shader, performing them on the GPU instead of the CPU.*/
	ShaderStageCompute ShaderStage = 4
/*Represents the size of the [enum ShaderStage] enum.*/
	ShaderStageMax ShaderStage = 5
/*Vertex shader stage bit (see also [constant SHADER_STAGE_VERTEX]).*/
	ShaderStageVertexBit ShaderStage = 1
/*Fragment shader stage bit (see also [constant SHADER_STAGE_FRAGMENT]).*/
	ShaderStageFragmentBit ShaderStage = 2
/*Tessellation control shader stage bit (see also [constant SHADER_STAGE_TESSELATION_CONTROL]).*/
	ShaderStageTesselationControlBit ShaderStage = 4
/*Tessellation evaluation shader stage bit (see also [constant SHADER_STAGE_TESSELATION_EVALUATION]).*/
	ShaderStageTesselationEvaluationBit ShaderStage = 8
/*Compute shader stage bit (see also [constant SHADER_STAGE_COMPUTE]).*/
	ShaderStageComputeBit ShaderStage = 16
)
type ShaderLanguage = classdb.RenderingDeviceShaderLanguage

const (
/*Khronos' GLSL shading language (used natively by OpenGL and Vulkan). This is the language used for core Godot shaders.*/
	ShaderLanguageGlsl ShaderLanguage = 0
/*Microsoft's High-Level Shading Language (used natively by Direct3D, but can also be used in Vulkan).*/
	ShaderLanguageHlsl ShaderLanguage = 1
)
type PipelineSpecializationConstantType = classdb.RenderingDevicePipelineSpecializationConstantType

const (
/*Boolean specialization constant.*/
	PipelineSpecializationConstantTypeBool PipelineSpecializationConstantType = 0
/*Integer specialization constant.*/
	PipelineSpecializationConstantTypeInt PipelineSpecializationConstantType = 1
/*Floating-point specialization constant.*/
	PipelineSpecializationConstantTypeFloat PipelineSpecializationConstantType = 2
)
type Limit = classdb.RenderingDeviceLimit

const (
/*Maximum number of uniform sets that can be bound at a given time.*/
	LimitMaxBoundUniformSets Limit = 0
/*Maximum number of color framebuffer attachments that can be used at a given time.*/
	LimitMaxFramebufferColorAttachments Limit = 1
/*Maximum number of textures that can be used per uniform set.*/
	LimitMaxTexturesPerUniformSet Limit = 2
/*Maximum number of samplers that can be used per uniform set.*/
	LimitMaxSamplersPerUniformSet Limit = 3
/*Maximum number of [url=https://vkguide.dev/docs/chapter-4/storage_buffers/]storage buffers[/url] per uniform set.*/
	LimitMaxStorageBuffersPerUniformSet Limit = 4
/*Maximum number of storage images per uniform set.*/
	LimitMaxStorageImagesPerUniformSet Limit = 5
/*Maximum number of uniform buffers per uniform set.*/
	LimitMaxUniformBuffersPerUniformSet Limit = 6
/*Maximum index for an indexed draw command.*/
	LimitMaxDrawIndexedIndex Limit = 7
/*Maximum height of a framebuffer (in pixels).*/
	LimitMaxFramebufferHeight Limit = 8
/*Maximum width of a framebuffer (in pixels).*/
	LimitMaxFramebufferWidth Limit = 9
/*Maximum number of texture array layers.*/
	LimitMaxTextureArrayLayers Limit = 10
/*Maximum supported 1-dimensional texture size (in pixels on a single axis).*/
	LimitMaxTextureSize1d Limit = 11
/*Maximum supported 2-dimensional texture size (in pixels on a single axis).*/
	LimitMaxTextureSize2d Limit = 12
/*Maximum supported 3-dimensional texture size (in pixels on a single axis).*/
	LimitMaxTextureSize3d Limit = 13
/*Maximum supported cubemap texture size (in pixels on a single axis of a single face).*/
	LimitMaxTextureSizeCube Limit = 14
/*Maximum number of textures per shader stage.*/
	LimitMaxTexturesPerShaderStage Limit = 15
/*Maximum number of samplers per shader stage.*/
	LimitMaxSamplersPerShaderStage Limit = 16
/*Maximum number of [url=https://vkguide.dev/docs/chapter-4/storage_buffers/]storage buffers[/url] per shader stage.*/
	LimitMaxStorageBuffersPerShaderStage Limit = 17
/*Maximum number of storage images per shader stage.*/
	LimitMaxStorageImagesPerShaderStage Limit = 18
/*Maximum number of uniform buffers per uniform set.*/
	LimitMaxUniformBuffersPerShaderStage Limit = 19
/*Maximum size of a push constant. A lot of devices are limited to 128 bytes, so try to avoid exceeding 128 bytes in push constants to ensure compatibility even if your GPU is reporting a higher value.*/
	LimitMaxPushConstantSize Limit = 20
/*Maximum size of a uniform buffer.*/
	LimitMaxUniformBufferSize Limit = 21
/*Maximum vertex input attribute offset.*/
	LimitMaxVertexInputAttributeOffset Limit = 22
/*Maximum number of vertex input attributes.*/
	LimitMaxVertexInputAttributes Limit = 23
/*Maximum number of vertex input bindings.*/
	LimitMaxVertexInputBindings Limit = 24
/*Maximum vertex input binding stride.*/
	LimitMaxVertexInputBindingStride Limit = 25
/*Minimum uniform buffer offset alignment.*/
	LimitMinUniformBufferOffsetAlignment Limit = 26
/*Maximum shared memory size for compute shaders.*/
	LimitMaxComputeSharedMemorySize Limit = 27
/*Maximum number of workgroups for compute shaders on the X axis.*/
	LimitMaxComputeWorkgroupCountX Limit = 28
/*Maximum number of workgroups for compute shaders on the Y axis.*/
	LimitMaxComputeWorkgroupCountY Limit = 29
/*Maximum number of workgroups for compute shaders on the Z axis.*/
	LimitMaxComputeWorkgroupCountZ Limit = 30
/*Maximum number of workgroup invocations for compute shaders.*/
	LimitMaxComputeWorkgroupInvocations Limit = 31
/*Maximum workgroup size for compute shaders on the X axis.*/
	LimitMaxComputeWorkgroupSizeX Limit = 32
/*Maximum workgroup size for compute shaders on the Y axis.*/
	LimitMaxComputeWorkgroupSizeY Limit = 33
/*Maximum workgroup size for compute shaders on the Z axis.*/
	LimitMaxComputeWorkgroupSizeZ Limit = 34
/*Maximum viewport width (in pixels).*/
	LimitMaxViewportDimensionsX Limit = 35
/*Maximum viewport height (in pixels).*/
	LimitMaxViewportDimensionsY Limit = 36
)
type MemoryType = classdb.RenderingDeviceMemoryType

const (
/*Memory taken by textures.*/
	MemoryTextures MemoryType = 0
/*Memory taken by buffers.*/
	MemoryBuffers MemoryType = 1
/*Total memory taken. This is greater than the sum of [constant MEMORY_TEXTURES] and [constant MEMORY_BUFFERS], as it also includes miscellaneous memory usage.*/
	MemoryTotal MemoryType = 2
)
