package RenderingServer

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The rendering server is the API backend for everything visible. The whole scene system mounts on it to display. The rendering server is completely opaque: the internals are entirely implementation-specific and cannot be accessed.
The rendering server can be used to bypass the scene/[Node] system entirely. This can improve performance in cases where the scene system is the bottleneck, but won't improve performance otherwise (for instance, if the GPU is already fully utilized).
Resources are created using the [code]*_create[/code] functions. These functions return [RID]s which are not references to the objects themselves, but opaque [i]pointers[/i] towards these objects.
All objects are drawn to a viewport. You can use the [Viewport] attached to the [SceneTree] or you can create one yourself with [method viewport_create]. When using a custom scenario or canvas, the scenario or canvas needs to be attached to the viewport using [method viewport_set_scenario] or [method viewport_attach_canvas].
[b]Scenarios:[/b] In 3D, all visual objects must be associated with a scenario. The scenario is a visual representation of the world. If accessing the rendering server from a running game, the scenario can be accessed from the scene tree from any [Node3D] node with [method Node3D.get_world_3d]. Otherwise, a scenario can be created with [method scenario_create].
Similarly, in 2D, a canvas is needed to draw all canvas items.
[b]3D:[/b] In 3D, all visible objects are comprised of a resource and an instance. A resource can be a mesh, a particle system, a light, or any other 3D object. In order to be visible resources must be attached to an instance using [method instance_set_base]. The instance must also be attached to the scenario using [method instance_set_scenario] in order to be visible. RenderingServer methods that don't have a prefix are usually 3D-specific (but not always).
[b]2D:[/b] In 2D, all visible objects are some form of canvas item. In order to be visible, a canvas item needs to be the child of a canvas attached to a viewport, or it needs to be the child of another canvas item that is eventually attached to the canvas. 2D-specific RenderingServer methods generally start with [code]canvas_*[/code].
[b]Headless mode:[/b] Starting the engine with the [code]--headless[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url] disables all rendering and window management functions. Most functions from [RenderingServer] will return dummy values in this case.

*/
type Simple [1]classdb.RenderingServer
func (self Simple) Texture2dCreate(image [1]classdb.Image) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).Texture2dCreate(image))
}
func (self Simple) Texture2dLayeredCreate(layers gd.ArrayOf[[1]classdb.Image], layered_type classdb.RenderingServerTextureLayeredType) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).Texture2dLayeredCreate(layers, layered_type))
}
func (self Simple) Texture3dCreate(format classdb.ImageFormat, width int, height int, depth int, mipmaps bool, data gd.ArrayOf[[1]classdb.Image]) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).Texture3dCreate(format, gd.Int(width), gd.Int(height), gd.Int(depth), mipmaps, data))
}
func (self Simple) TextureProxyCreate(base gd.RID) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).TextureProxyCreate(base))
}
func (self Simple) Texture2dUpdate(texture gd.RID, image [1]classdb.Image, layer int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Texture2dUpdate(texture, image, gd.Int(layer))
}
func (self Simple) Texture3dUpdate(texture gd.RID, data gd.ArrayOf[[1]classdb.Image]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Texture3dUpdate(texture, data)
}
func (self Simple) TextureProxyUpdate(texture gd.RID, proxy_to gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).TextureProxyUpdate(texture, proxy_to)
}
func (self Simple) Texture2dPlaceholderCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).Texture2dPlaceholderCreate())
}
func (self Simple) Texture2dLayeredPlaceholderCreate(layered_type classdb.RenderingServerTextureLayeredType) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).Texture2dLayeredPlaceholderCreate(layered_type))
}
func (self Simple) Texture3dPlaceholderCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).Texture3dPlaceholderCreate())
}
func (self Simple) Texture2dGet(texture gd.RID) [1]classdb.Image {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Image(Expert(self).Texture2dGet(gc, texture))
}
func (self Simple) Texture2dLayerGet(texture gd.RID, layer int) [1]classdb.Image {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Image(Expert(self).Texture2dLayerGet(gc, texture, gd.Int(layer)))
}
func (self Simple) Texture3dGet(texture gd.RID) gd.ArrayOf[[1]classdb.Image] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.Image](Expert(self).Texture3dGet(gc, texture))
}
func (self Simple) TextureReplace(texture gd.RID, by_texture gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).TextureReplace(texture, by_texture)
}
func (self Simple) TextureSetSizeOverride(texture gd.RID, width int, height int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).TextureSetSizeOverride(texture, gd.Int(width), gd.Int(height))
}
func (self Simple) TextureSetPath(texture gd.RID, path string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).TextureSetPath(texture, gc.String(path))
}
func (self Simple) TextureGetPath(texture gd.RID) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).TextureGetPath(gc, texture).String())
}
func (self Simple) TextureGetFormat(texture gd.RID) classdb.ImageFormat {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ImageFormat(Expert(self).TextureGetFormat(texture))
}
func (self Simple) TextureSetForceRedrawIfVisible(texture gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).TextureSetForceRedrawIfVisible(texture, enable)
}
func (self Simple) TextureRdCreate(rd_texture gd.RID, layer_type classdb.RenderingServerTextureLayeredType) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).TextureRdCreate(rd_texture, layer_type))
}
func (self Simple) TextureGetRdTexture(texture gd.RID, srgb bool) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).TextureGetRdTexture(texture, srgb))
}
func (self Simple) TextureGetNativeHandle(texture gd.RID, srgb bool) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).TextureGetNativeHandle(texture, srgb)))
}
func (self Simple) ShaderCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).ShaderCreate())
}
func (self Simple) ShaderSetCode(shader gd.RID, code string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ShaderSetCode(shader, gc.String(code))
}
func (self Simple) ShaderSetPathHint(shader gd.RID, path string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ShaderSetPathHint(shader, gc.String(path))
}
func (self Simple) ShaderGetCode(shader gd.RID) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).ShaderGetCode(gc, shader).String())
}
func (self Simple) GetShaderParameterList(shader gd.RID) gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](Expert(self).GetShaderParameterList(gc, shader))
}
func (self Simple) ShaderGetParameterDefault(shader gd.RID, name string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).ShaderGetParameterDefault(gc, shader, gc.StringName(name)))
}
func (self Simple) ShaderSetDefaultTextureParameter(shader gd.RID, name string, texture gd.RID, index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ShaderSetDefaultTextureParameter(shader, gc.StringName(name), texture, gd.Int(index))
}
func (self Simple) ShaderGetDefaultTextureParameter(shader gd.RID, name string, index int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).ShaderGetDefaultTextureParameter(shader, gc.StringName(name), gd.Int(index)))
}
func (self Simple) MaterialCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).MaterialCreate())
}
func (self Simple) MaterialSetShader(shader_material gd.RID, shader gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MaterialSetShader(shader_material, shader)
}
func (self Simple) MaterialSetParam(material gd.RID, parameter string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MaterialSetParam(material, gc.StringName(parameter), value)
}
func (self Simple) MaterialGetParam(material gd.RID, parameter string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).MaterialGetParam(gc, material, gc.StringName(parameter)))
}
func (self Simple) MaterialSetRenderPriority(material gd.RID, priority int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MaterialSetRenderPriority(material, gd.Int(priority))
}
func (self Simple) MaterialSetNextPass(material gd.RID, next_material gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MaterialSetNextPass(material, next_material)
}
func (self Simple) MeshCreateFromSurfaces(surfaces gd.ArrayOf[gd.Dictionary], blend_shape_count int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).MeshCreateFromSurfaces(surfaces, gd.Int(blend_shape_count)))
}
func (self Simple) MeshCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).MeshCreate())
}
func (self Simple) MeshSurfaceGetFormatOffset(format classdb.RenderingServerArrayFormat, vertex_count int, array_index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).MeshSurfaceGetFormatOffset(format, gd.Int(vertex_count), gd.Int(array_index))))
}
func (self Simple) MeshSurfaceGetFormatVertexStride(format classdb.RenderingServerArrayFormat, vertex_count int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).MeshSurfaceGetFormatVertexStride(format, gd.Int(vertex_count))))
}
func (self Simple) MeshSurfaceGetFormatNormalTangentStride(format classdb.RenderingServerArrayFormat, vertex_count int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).MeshSurfaceGetFormatNormalTangentStride(format, gd.Int(vertex_count))))
}
func (self Simple) MeshSurfaceGetFormatAttributeStride(format classdb.RenderingServerArrayFormat, vertex_count int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).MeshSurfaceGetFormatAttributeStride(format, gd.Int(vertex_count))))
}
func (self Simple) MeshSurfaceGetFormatSkinStride(format classdb.RenderingServerArrayFormat, vertex_count int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).MeshSurfaceGetFormatSkinStride(format, gd.Int(vertex_count))))
}
func (self Simple) MeshAddSurface(mesh gd.RID, surface gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MeshAddSurface(mesh, surface)
}
func (self Simple) MeshAddSurfaceFromArrays(mesh gd.RID, primitive classdb.RenderingServerPrimitiveType, arrays gd.Array, blend_shapes gd.Array, lods gd.Dictionary, compress_format classdb.RenderingServerArrayFormat) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MeshAddSurfaceFromArrays(mesh, primitive, arrays, blend_shapes, lods, compress_format)
}
func (self Simple) MeshGetBlendShapeCount(mesh gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).MeshGetBlendShapeCount(mesh)))
}
func (self Simple) MeshSetBlendShapeMode(mesh gd.RID, mode classdb.RenderingServerBlendShapeMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MeshSetBlendShapeMode(mesh, mode)
}
func (self Simple) MeshGetBlendShapeMode(mesh gd.RID) classdb.RenderingServerBlendShapeMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingServerBlendShapeMode(Expert(self).MeshGetBlendShapeMode(mesh))
}
func (self Simple) MeshSurfaceSetMaterial(mesh gd.RID, surface int, material gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MeshSurfaceSetMaterial(mesh, gd.Int(surface), material)
}
func (self Simple) MeshSurfaceGetMaterial(mesh gd.RID, surface int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).MeshSurfaceGetMaterial(mesh, gd.Int(surface)))
}
func (self Simple) MeshGetSurface(mesh gd.RID, surface int) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).MeshGetSurface(gc, mesh, gd.Int(surface)))
}
func (self Simple) MeshSurfaceGetArrays(mesh gd.RID, surface int) gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).MeshSurfaceGetArrays(gc, mesh, gd.Int(surface)))
}
func (self Simple) MeshSurfaceGetBlendShapeArrays(mesh gd.RID, surface int) gd.ArrayOf[gd.Array] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Array](Expert(self).MeshSurfaceGetBlendShapeArrays(gc, mesh, gd.Int(surface)))
}
func (self Simple) MeshGetSurfaceCount(mesh gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).MeshGetSurfaceCount(mesh)))
}
func (self Simple) MeshSetCustomAabb(mesh gd.RID, aabb gd.AABB) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MeshSetCustomAabb(mesh, aabb)
}
func (self Simple) MeshGetCustomAabb(mesh gd.RID) gd.AABB {
	gc := gd.GarbageCollector(); _ = gc
	return gd.AABB(Expert(self).MeshGetCustomAabb(mesh))
}
func (self Simple) MeshClear(mesh gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MeshClear(mesh)
}
func (self Simple) MeshSurfaceUpdateVertexRegion(mesh gd.RID, surface int, offset int, data []byte) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MeshSurfaceUpdateVertexRegion(mesh, gd.Int(surface), gd.Int(offset), gc.PackedByteSlice(data))
}
func (self Simple) MeshSurfaceUpdateAttributeRegion(mesh gd.RID, surface int, offset int, data []byte) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MeshSurfaceUpdateAttributeRegion(mesh, gd.Int(surface), gd.Int(offset), gc.PackedByteSlice(data))
}
func (self Simple) MeshSurfaceUpdateSkinRegion(mesh gd.RID, surface int, offset int, data []byte) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MeshSurfaceUpdateSkinRegion(mesh, gd.Int(surface), gd.Int(offset), gc.PackedByteSlice(data))
}
func (self Simple) MeshSetShadowMesh(mesh gd.RID, shadow_mesh gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MeshSetShadowMesh(mesh, shadow_mesh)
}
func (self Simple) MultimeshCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).MultimeshCreate())
}
func (self Simple) MultimeshAllocateData(multimesh gd.RID, instances int, transform_format classdb.RenderingServerMultimeshTransformFormat, color_format bool, custom_data_format bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MultimeshAllocateData(multimesh, gd.Int(instances), transform_format, color_format, custom_data_format)
}
func (self Simple) MultimeshGetInstanceCount(multimesh gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).MultimeshGetInstanceCount(multimesh)))
}
func (self Simple) MultimeshSetMesh(multimesh gd.RID, mesh gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MultimeshSetMesh(multimesh, mesh)
}
func (self Simple) MultimeshInstanceSetTransform(multimesh gd.RID, index int, transform gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MultimeshInstanceSetTransform(multimesh, gd.Int(index), transform)
}
func (self Simple) MultimeshInstanceSetTransform2d(multimesh gd.RID, index int, transform gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MultimeshInstanceSetTransform2d(multimesh, gd.Int(index), transform)
}
func (self Simple) MultimeshInstanceSetColor(multimesh gd.RID, index int, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MultimeshInstanceSetColor(multimesh, gd.Int(index), color)
}
func (self Simple) MultimeshInstanceSetCustomData(multimesh gd.RID, index int, custom_data gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MultimeshInstanceSetCustomData(multimesh, gd.Int(index), custom_data)
}
func (self Simple) MultimeshGetMesh(multimesh gd.RID) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).MultimeshGetMesh(multimesh))
}
func (self Simple) MultimeshGetAabb(multimesh gd.RID) gd.AABB {
	gc := gd.GarbageCollector(); _ = gc
	return gd.AABB(Expert(self).MultimeshGetAabb(multimesh))
}
func (self Simple) MultimeshSetCustomAabb(multimesh gd.RID, aabb gd.AABB) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MultimeshSetCustomAabb(multimesh, aabb)
}
func (self Simple) MultimeshGetCustomAabb(multimesh gd.RID) gd.AABB {
	gc := gd.GarbageCollector(); _ = gc
	return gd.AABB(Expert(self).MultimeshGetCustomAabb(multimesh))
}
func (self Simple) MultimeshInstanceGetTransform(multimesh gd.RID, index int) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).MultimeshInstanceGetTransform(multimesh, gd.Int(index)))
}
func (self Simple) MultimeshInstanceGetTransform2d(multimesh gd.RID, index int) gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(Expert(self).MultimeshInstanceGetTransform2d(multimesh, gd.Int(index)))
}
func (self Simple) MultimeshInstanceGetColor(multimesh gd.RID, index int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).MultimeshInstanceGetColor(multimesh, gd.Int(index)))
}
func (self Simple) MultimeshInstanceGetCustomData(multimesh gd.RID, index int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).MultimeshInstanceGetCustomData(multimesh, gd.Int(index)))
}
func (self Simple) MultimeshSetVisibleInstances(multimesh gd.RID, visible int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MultimeshSetVisibleInstances(multimesh, gd.Int(visible))
}
func (self Simple) MultimeshGetVisibleInstances(multimesh gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).MultimeshGetVisibleInstances(multimesh)))
}
func (self Simple) MultimeshSetBuffer(multimesh gd.RID, buffer gd.PackedFloat32Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MultimeshSetBuffer(multimesh, buffer)
}
func (self Simple) MultimeshGetBuffer(multimesh gd.RID) gd.PackedFloat32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedFloat32Array(Expert(self).MultimeshGetBuffer(gc, multimesh))
}
func (self Simple) SkeletonCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).SkeletonCreate())
}
func (self Simple) SkeletonAllocateData(skeleton gd.RID, bones int, is_2d_skeleton bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SkeletonAllocateData(skeleton, gd.Int(bones), is_2d_skeleton)
}
func (self Simple) SkeletonGetBoneCount(skeleton gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).SkeletonGetBoneCount(skeleton)))
}
func (self Simple) SkeletonBoneSetTransform(skeleton gd.RID, bone int, transform gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SkeletonBoneSetTransform(skeleton, gd.Int(bone), transform)
}
func (self Simple) SkeletonBoneGetTransform(skeleton gd.RID, bone int) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).SkeletonBoneGetTransform(skeleton, gd.Int(bone)))
}
func (self Simple) SkeletonBoneSetTransform2d(skeleton gd.RID, bone int, transform gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SkeletonBoneSetTransform2d(skeleton, gd.Int(bone), transform)
}
func (self Simple) SkeletonBoneGetTransform2d(skeleton gd.RID, bone int) gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(Expert(self).SkeletonBoneGetTransform2d(skeleton, gd.Int(bone)))
}
func (self Simple) SkeletonSetBaseTransform2d(skeleton gd.RID, base_transform gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SkeletonSetBaseTransform2d(skeleton, base_transform)
}
func (self Simple) DirectionalLightCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).DirectionalLightCreate())
}
func (self Simple) OmniLightCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).OmniLightCreate())
}
func (self Simple) SpotLightCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).SpotLightCreate())
}
func (self Simple) LightSetColor(light gd.RID, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightSetColor(light, color)
}
func (self Simple) LightSetParam(light gd.RID, param classdb.RenderingServerLightParam, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightSetParam(light, param, gd.Float(value))
}
func (self Simple) LightSetShadow(light gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightSetShadow(light, enabled)
}
func (self Simple) LightSetProjector(light gd.RID, texture gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightSetProjector(light, texture)
}
func (self Simple) LightSetNegative(light gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightSetNegative(light, enable)
}
func (self Simple) LightSetCullMask(light gd.RID, mask int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightSetCullMask(light, gd.Int(mask))
}
func (self Simple) LightSetDistanceFade(decal gd.RID, enabled bool, begin float64, shadow float64, length float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightSetDistanceFade(decal, enabled, gd.Float(begin), gd.Float(shadow), gd.Float(length))
}
func (self Simple) LightSetReverseCullFaceMode(light gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightSetReverseCullFaceMode(light, enabled)
}
func (self Simple) LightSetBakeMode(light gd.RID, bake_mode classdb.RenderingServerLightBakeMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightSetBakeMode(light, bake_mode)
}
func (self Simple) LightSetMaxSdfgiCascade(light gd.RID, cascade int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightSetMaxSdfgiCascade(light, gd.Int(cascade))
}
func (self Simple) LightOmniSetShadowMode(light gd.RID, mode classdb.RenderingServerLightOmniShadowMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightOmniSetShadowMode(light, mode)
}
func (self Simple) LightDirectionalSetShadowMode(light gd.RID, mode classdb.RenderingServerLightDirectionalShadowMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightDirectionalSetShadowMode(light, mode)
}
func (self Simple) LightDirectionalSetBlendSplits(light gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightDirectionalSetBlendSplits(light, enable)
}
func (self Simple) LightDirectionalSetSkyMode(light gd.RID, mode classdb.RenderingServerLightDirectionalSkyMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightDirectionalSetSkyMode(light, mode)
}
func (self Simple) LightProjectorsSetFilter(filter classdb.RenderingServerLightProjectorFilter) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightProjectorsSetFilter(filter)
}
func (self Simple) PositionalSoftShadowFilterSetQuality(quality classdb.RenderingServerShadowQuality) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PositionalSoftShadowFilterSetQuality(quality)
}
func (self Simple) DirectionalSoftShadowFilterSetQuality(quality classdb.RenderingServerShadowQuality) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DirectionalSoftShadowFilterSetQuality(quality)
}
func (self Simple) DirectionalShadowAtlasSetSize(size int, is_16bits bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DirectionalShadowAtlasSetSize(gd.Int(size), is_16bits)
}
func (self Simple) ReflectionProbeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).ReflectionProbeCreate())
}
func (self Simple) ReflectionProbeSetUpdateMode(probe gd.RID, mode classdb.RenderingServerReflectionProbeUpdateMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReflectionProbeSetUpdateMode(probe, mode)
}
func (self Simple) ReflectionProbeSetIntensity(probe gd.RID, intensity float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReflectionProbeSetIntensity(probe, gd.Float(intensity))
}
func (self Simple) ReflectionProbeSetAmbientMode(probe gd.RID, mode classdb.RenderingServerReflectionProbeAmbientMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReflectionProbeSetAmbientMode(probe, mode)
}
func (self Simple) ReflectionProbeSetAmbientColor(probe gd.RID, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReflectionProbeSetAmbientColor(probe, color)
}
func (self Simple) ReflectionProbeSetAmbientEnergy(probe gd.RID, energy float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReflectionProbeSetAmbientEnergy(probe, gd.Float(energy))
}
func (self Simple) ReflectionProbeSetMaxDistance(probe gd.RID, distance float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReflectionProbeSetMaxDistance(probe, gd.Float(distance))
}
func (self Simple) ReflectionProbeSetSize(probe gd.RID, size gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReflectionProbeSetSize(probe, size)
}
func (self Simple) ReflectionProbeSetOriginOffset(probe gd.RID, offset gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReflectionProbeSetOriginOffset(probe, offset)
}
func (self Simple) ReflectionProbeSetAsInterior(probe gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReflectionProbeSetAsInterior(probe, enable)
}
func (self Simple) ReflectionProbeSetEnableBoxProjection(probe gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReflectionProbeSetEnableBoxProjection(probe, enable)
}
func (self Simple) ReflectionProbeSetEnableShadows(probe gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReflectionProbeSetEnableShadows(probe, enable)
}
func (self Simple) ReflectionProbeSetCullMask(probe gd.RID, layers int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReflectionProbeSetCullMask(probe, gd.Int(layers))
}
func (self Simple) ReflectionProbeSetReflectionMask(probe gd.RID, layers int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReflectionProbeSetReflectionMask(probe, gd.Int(layers))
}
func (self Simple) ReflectionProbeSetResolution(probe gd.RID, resolution int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReflectionProbeSetResolution(probe, gd.Int(resolution))
}
func (self Simple) ReflectionProbeSetMeshLodThreshold(probe gd.RID, pixels float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReflectionProbeSetMeshLodThreshold(probe, gd.Float(pixels))
}
func (self Simple) DecalCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).DecalCreate())
}
func (self Simple) DecalSetSize(decal gd.RID, size gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DecalSetSize(decal, size)
}
func (self Simple) DecalSetTexture(decal gd.RID, atype classdb.RenderingServerDecalTexture, texture gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DecalSetTexture(decal, atype, texture)
}
func (self Simple) DecalSetEmissionEnergy(decal gd.RID, energy float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DecalSetEmissionEnergy(decal, gd.Float(energy))
}
func (self Simple) DecalSetAlbedoMix(decal gd.RID, albedo_mix float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DecalSetAlbedoMix(decal, gd.Float(albedo_mix))
}
func (self Simple) DecalSetModulate(decal gd.RID, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DecalSetModulate(decal, color)
}
func (self Simple) DecalSetCullMask(decal gd.RID, mask int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DecalSetCullMask(decal, gd.Int(mask))
}
func (self Simple) DecalSetDistanceFade(decal gd.RID, enabled bool, begin float64, length float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DecalSetDistanceFade(decal, enabled, gd.Float(begin), gd.Float(length))
}
func (self Simple) DecalSetFade(decal gd.RID, above float64, below float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DecalSetFade(decal, gd.Float(above), gd.Float(below))
}
func (self Simple) DecalSetNormalFade(decal gd.RID, fade float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DecalSetNormalFade(decal, gd.Float(fade))
}
func (self Simple) DecalsSetFilter(filter classdb.RenderingServerDecalFilter) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DecalsSetFilter(filter)
}
func (self Simple) GiSetUseHalfResolution(half_resolution bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GiSetUseHalfResolution(half_resolution)
}
func (self Simple) VoxelGiCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).VoxelGiCreate())
}
func (self Simple) VoxelGiAllocateData(voxel_gi gd.RID, to_cell_xform gd.Transform3D, aabb gd.AABB, octree_size gd.Vector3i, octree_cells []byte, data_cells []byte, distance_field []byte, level_counts gd.PackedInt32Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).VoxelGiAllocateData(voxel_gi, to_cell_xform, aabb, octree_size, gc.PackedByteSlice(octree_cells), gc.PackedByteSlice(data_cells), gc.PackedByteSlice(distance_field), level_counts)
}
func (self Simple) VoxelGiGetOctreeSize(voxel_gi gd.RID) gd.Vector3i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3i(Expert(self).VoxelGiGetOctreeSize(voxel_gi))
}
func (self Simple) VoxelGiGetOctreeCells(voxel_gi gd.RID) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).VoxelGiGetOctreeCells(gc, voxel_gi).Bytes())
}
func (self Simple) VoxelGiGetDataCells(voxel_gi gd.RID) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).VoxelGiGetDataCells(gc, voxel_gi).Bytes())
}
func (self Simple) VoxelGiGetDistanceField(voxel_gi gd.RID) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).VoxelGiGetDistanceField(gc, voxel_gi).Bytes())
}
func (self Simple) VoxelGiGetLevelCounts(voxel_gi gd.RID) gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).VoxelGiGetLevelCounts(gc, voxel_gi))
}
func (self Simple) VoxelGiGetToCellXform(voxel_gi gd.RID) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).VoxelGiGetToCellXform(voxel_gi))
}
func (self Simple) VoxelGiSetDynamicRange(voxel_gi gd.RID, arange float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).VoxelGiSetDynamicRange(voxel_gi, gd.Float(arange))
}
func (self Simple) VoxelGiSetPropagation(voxel_gi gd.RID, amount float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).VoxelGiSetPropagation(voxel_gi, gd.Float(amount))
}
func (self Simple) VoxelGiSetEnergy(voxel_gi gd.RID, energy float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).VoxelGiSetEnergy(voxel_gi, gd.Float(energy))
}
func (self Simple) VoxelGiSetBakedExposureNormalization(voxel_gi gd.RID, baked_exposure float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).VoxelGiSetBakedExposureNormalization(voxel_gi, gd.Float(baked_exposure))
}
func (self Simple) VoxelGiSetBias(voxel_gi gd.RID, bias float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).VoxelGiSetBias(voxel_gi, gd.Float(bias))
}
func (self Simple) VoxelGiSetNormalBias(voxel_gi gd.RID, bias float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).VoxelGiSetNormalBias(voxel_gi, gd.Float(bias))
}
func (self Simple) VoxelGiSetInterior(voxel_gi gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).VoxelGiSetInterior(voxel_gi, enable)
}
func (self Simple) VoxelGiSetUseTwoBounces(voxel_gi gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).VoxelGiSetUseTwoBounces(voxel_gi, enable)
}
func (self Simple) VoxelGiSetQuality(quality classdb.RenderingServerVoxelGIQuality) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).VoxelGiSetQuality(quality)
}
func (self Simple) LightmapCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).LightmapCreate())
}
func (self Simple) LightmapSetTextures(lightmap gd.RID, light gd.RID, uses_sh bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightmapSetTextures(lightmap, light, uses_sh)
}
func (self Simple) LightmapSetProbeBounds(lightmap gd.RID, bounds gd.AABB) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightmapSetProbeBounds(lightmap, bounds)
}
func (self Simple) LightmapSetProbeInterior(lightmap gd.RID, interior bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightmapSetProbeInterior(lightmap, interior)
}
func (self Simple) LightmapSetProbeCaptureData(lightmap gd.RID, points gd.PackedVector3Array, point_sh gd.PackedColorArray, tetrahedra gd.PackedInt32Array, bsp_tree gd.PackedInt32Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightmapSetProbeCaptureData(lightmap, points, point_sh, tetrahedra, bsp_tree)
}
func (self Simple) LightmapGetProbeCapturePoints(lightmap gd.RID) gd.PackedVector3Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector3Array(Expert(self).LightmapGetProbeCapturePoints(gc, lightmap))
}
func (self Simple) LightmapGetProbeCaptureSh(lightmap gd.RID) gd.PackedColorArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedColorArray(Expert(self).LightmapGetProbeCaptureSh(gc, lightmap))
}
func (self Simple) LightmapGetProbeCaptureTetrahedra(lightmap gd.RID) gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).LightmapGetProbeCaptureTetrahedra(gc, lightmap))
}
func (self Simple) LightmapGetProbeCaptureBspTree(lightmap gd.RID) gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).LightmapGetProbeCaptureBspTree(gc, lightmap))
}
func (self Simple) LightmapSetBakedExposureNormalization(lightmap gd.RID, baked_exposure float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightmapSetBakedExposureNormalization(lightmap, gd.Float(baked_exposure))
}
func (self Simple) LightmapSetProbeCaptureUpdateSpeed(speed float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LightmapSetProbeCaptureUpdateSpeed(gd.Float(speed))
}
func (self Simple) ParticlesCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).ParticlesCreate())
}
func (self Simple) ParticlesSetMode(particles gd.RID, mode classdb.RenderingServerParticlesMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetMode(particles, mode)
}
func (self Simple) ParticlesSetEmitting(particles gd.RID, emitting bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetEmitting(particles, emitting)
}
func (self Simple) ParticlesGetEmitting(particles gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).ParticlesGetEmitting(particles))
}
func (self Simple) ParticlesSetAmount(particles gd.RID, amount int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetAmount(particles, gd.Int(amount))
}
func (self Simple) ParticlesSetAmountRatio(particles gd.RID, ratio float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetAmountRatio(particles, gd.Float(ratio))
}
func (self Simple) ParticlesSetLifetime(particles gd.RID, lifetime float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetLifetime(particles, gd.Float(lifetime))
}
func (self Simple) ParticlesSetOneShot(particles gd.RID, one_shot bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetOneShot(particles, one_shot)
}
func (self Simple) ParticlesSetPreProcessTime(particles gd.RID, time float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetPreProcessTime(particles, gd.Float(time))
}
func (self Simple) ParticlesSetExplosivenessRatio(particles gd.RID, ratio float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetExplosivenessRatio(particles, gd.Float(ratio))
}
func (self Simple) ParticlesSetRandomnessRatio(particles gd.RID, ratio float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetRandomnessRatio(particles, gd.Float(ratio))
}
func (self Simple) ParticlesSetInterpToEnd(particles gd.RID, factor float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetInterpToEnd(particles, gd.Float(factor))
}
func (self Simple) ParticlesSetEmitterVelocity(particles gd.RID, velocity gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetEmitterVelocity(particles, velocity)
}
func (self Simple) ParticlesSetCustomAabb(particles gd.RID, aabb gd.AABB) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetCustomAabb(particles, aabb)
}
func (self Simple) ParticlesSetSpeedScale(particles gd.RID, scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetSpeedScale(particles, gd.Float(scale))
}
func (self Simple) ParticlesSetUseLocalCoordinates(particles gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetUseLocalCoordinates(particles, enable)
}
func (self Simple) ParticlesSetProcessMaterial(particles gd.RID, material gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetProcessMaterial(particles, material)
}
func (self Simple) ParticlesSetFixedFps(particles gd.RID, fps int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetFixedFps(particles, gd.Int(fps))
}
func (self Simple) ParticlesSetInterpolate(particles gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetInterpolate(particles, enable)
}
func (self Simple) ParticlesSetFractionalDelta(particles gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetFractionalDelta(particles, enable)
}
func (self Simple) ParticlesSetCollisionBaseSize(particles gd.RID, size float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetCollisionBaseSize(particles, gd.Float(size))
}
func (self Simple) ParticlesSetTransformAlign(particles gd.RID, align classdb.RenderingServerParticlesTransformAlign) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetTransformAlign(particles, align)
}
func (self Simple) ParticlesSetTrails(particles gd.RID, enable bool, length_sec float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetTrails(particles, enable, gd.Float(length_sec))
}
func (self Simple) ParticlesSetTrailBindPoses(particles gd.RID, bind_poses gd.ArrayOf[gd.Transform3D]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetTrailBindPoses(particles, bind_poses)
}
func (self Simple) ParticlesIsInactive(particles gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).ParticlesIsInactive(particles))
}
func (self Simple) ParticlesRequestProcess(particles gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesRequestProcess(particles)
}
func (self Simple) ParticlesRestart(particles gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesRestart(particles)
}
func (self Simple) ParticlesSetSubemitter(particles gd.RID, subemitter_particles gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetSubemitter(particles, subemitter_particles)
}
func (self Simple) ParticlesEmit(particles gd.RID, transform gd.Transform3D, velocity gd.Vector3, color gd.Color, custom gd.Color, emit_flags int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesEmit(particles, transform, velocity, color, custom, gd.Int(emit_flags))
}
func (self Simple) ParticlesSetDrawOrder(particles gd.RID, order classdb.RenderingServerParticlesDrawOrder) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetDrawOrder(particles, order)
}
func (self Simple) ParticlesSetDrawPasses(particles gd.RID, count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetDrawPasses(particles, gd.Int(count))
}
func (self Simple) ParticlesSetDrawPassMesh(particles gd.RID, pass int, mesh gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetDrawPassMesh(particles, gd.Int(pass), mesh)
}
func (self Simple) ParticlesGetCurrentAabb(particles gd.RID) gd.AABB {
	gc := gd.GarbageCollector(); _ = gc
	return gd.AABB(Expert(self).ParticlesGetCurrentAabb(particles))
}
func (self Simple) ParticlesSetEmissionTransform(particles gd.RID, transform gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesSetEmissionTransform(particles, transform)
}
func (self Simple) ParticlesCollisionCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).ParticlesCollisionCreate())
}
func (self Simple) ParticlesCollisionSetCollisionType(particles_collision gd.RID, atype classdb.RenderingServerParticlesCollisionType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesCollisionSetCollisionType(particles_collision, atype)
}
func (self Simple) ParticlesCollisionSetCullMask(particles_collision gd.RID, mask int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesCollisionSetCullMask(particles_collision, gd.Int(mask))
}
func (self Simple) ParticlesCollisionSetSphereRadius(particles_collision gd.RID, radius float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesCollisionSetSphereRadius(particles_collision, gd.Float(radius))
}
func (self Simple) ParticlesCollisionSetBoxExtents(particles_collision gd.RID, extents gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesCollisionSetBoxExtents(particles_collision, extents)
}
func (self Simple) ParticlesCollisionSetAttractorStrength(particles_collision gd.RID, strength float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesCollisionSetAttractorStrength(particles_collision, gd.Float(strength))
}
func (self Simple) ParticlesCollisionSetAttractorDirectionality(particles_collision gd.RID, amount float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesCollisionSetAttractorDirectionality(particles_collision, gd.Float(amount))
}
func (self Simple) ParticlesCollisionSetAttractorAttenuation(particles_collision gd.RID, curve float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesCollisionSetAttractorAttenuation(particles_collision, gd.Float(curve))
}
func (self Simple) ParticlesCollisionSetFieldTexture(particles_collision gd.RID, texture gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesCollisionSetFieldTexture(particles_collision, texture)
}
func (self Simple) ParticlesCollisionHeightFieldUpdate(particles_collision gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesCollisionHeightFieldUpdate(particles_collision)
}
func (self Simple) ParticlesCollisionSetHeightFieldResolution(particles_collision gd.RID, resolution classdb.RenderingServerParticlesCollisionHeightfieldResolution) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParticlesCollisionSetHeightFieldResolution(particles_collision, resolution)
}
func (self Simple) FogVolumeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).FogVolumeCreate())
}
func (self Simple) FogVolumeSetShape(fog_volume gd.RID, shape classdb.RenderingServerFogVolumeShape) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).FogVolumeSetShape(fog_volume, shape)
}
func (self Simple) FogVolumeSetSize(fog_volume gd.RID, size gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).FogVolumeSetSize(fog_volume, size)
}
func (self Simple) FogVolumeSetMaterial(fog_volume gd.RID, material gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).FogVolumeSetMaterial(fog_volume, material)
}
func (self Simple) VisibilityNotifierCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).VisibilityNotifierCreate())
}
func (self Simple) VisibilityNotifierSetAabb(notifier gd.RID, aabb gd.AABB) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).VisibilityNotifierSetAabb(notifier, aabb)
}
func (self Simple) VisibilityNotifierSetCallbacks(notifier gd.RID, enter_callable gd.Callable, exit_callable gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).VisibilityNotifierSetCallbacks(notifier, enter_callable, exit_callable)
}
func (self Simple) OccluderCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).OccluderCreate())
}
func (self Simple) OccluderSetMesh(occluder gd.RID, vertices gd.PackedVector3Array, indices gd.PackedInt32Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).OccluderSetMesh(occluder, vertices, indices)
}
func (self Simple) CameraCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).CameraCreate())
}
func (self Simple) CameraSetPerspective(camera gd.RID, fovy_degrees float64, z_near float64, z_far float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CameraSetPerspective(camera, gd.Float(fovy_degrees), gd.Float(z_near), gd.Float(z_far))
}
func (self Simple) CameraSetOrthogonal(camera gd.RID, size float64, z_near float64, z_far float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CameraSetOrthogonal(camera, gd.Float(size), gd.Float(z_near), gd.Float(z_far))
}
func (self Simple) CameraSetFrustum(camera gd.RID, size float64, offset gd.Vector2, z_near float64, z_far float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CameraSetFrustum(camera, gd.Float(size), offset, gd.Float(z_near), gd.Float(z_far))
}
func (self Simple) CameraSetTransform(camera gd.RID, transform gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CameraSetTransform(camera, transform)
}
func (self Simple) CameraSetCullMask(camera gd.RID, layers int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CameraSetCullMask(camera, gd.Int(layers))
}
func (self Simple) CameraSetEnvironment(camera gd.RID, env gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CameraSetEnvironment(camera, env)
}
func (self Simple) CameraSetCameraAttributes(camera gd.RID, effects gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CameraSetCameraAttributes(camera, effects)
}
func (self Simple) CameraSetCompositor(camera gd.RID, compositor gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CameraSetCompositor(camera, compositor)
}
func (self Simple) CameraSetUseVerticalAspect(camera gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CameraSetUseVerticalAspect(camera, enable)
}
func (self Simple) ViewportCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).ViewportCreate())
}
func (self Simple) ViewportSetUseXr(viewport gd.RID, use_xr bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetUseXr(viewport, use_xr)
}
func (self Simple) ViewportSetSize(viewport gd.RID, width int, height int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetSize(viewport, gd.Int(width), gd.Int(height))
}
func (self Simple) ViewportSetActive(viewport gd.RID, active bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetActive(viewport, active)
}
func (self Simple) ViewportSetParentViewport(viewport gd.RID, parent_viewport gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetParentViewport(viewport, parent_viewport)
}
func (self Simple) ViewportAttachToScreen(viewport gd.RID, rect gd.Rect2, screen int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportAttachToScreen(viewport, rect, gd.Int(screen))
}
func (self Simple) ViewportSetRenderDirectToScreen(viewport gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetRenderDirectToScreen(viewport, enabled)
}
func (self Simple) ViewportSetCanvasCullMask(viewport gd.RID, canvas_cull_mask int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetCanvasCullMask(viewport, gd.Int(canvas_cull_mask))
}
func (self Simple) ViewportSetScaling3dMode(viewport gd.RID, scaling_3d_mode classdb.RenderingServerViewportScaling3DMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetScaling3dMode(viewport, scaling_3d_mode)
}
func (self Simple) ViewportSetScaling3dScale(viewport gd.RID, scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetScaling3dScale(viewport, gd.Float(scale))
}
func (self Simple) ViewportSetFsrSharpness(viewport gd.RID, sharpness float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetFsrSharpness(viewport, gd.Float(sharpness))
}
func (self Simple) ViewportSetTextureMipmapBias(viewport gd.RID, mipmap_bias float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetTextureMipmapBias(viewport, gd.Float(mipmap_bias))
}
func (self Simple) ViewportSetUpdateMode(viewport gd.RID, update_mode classdb.RenderingServerViewportUpdateMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetUpdateMode(viewport, update_mode)
}
func (self Simple) ViewportGetUpdateMode(viewport gd.RID) classdb.RenderingServerViewportUpdateMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingServerViewportUpdateMode(Expert(self).ViewportGetUpdateMode(viewport))
}
func (self Simple) ViewportSetClearMode(viewport gd.RID, clear_mode classdb.RenderingServerViewportClearMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetClearMode(viewport, clear_mode)
}
func (self Simple) ViewportGetRenderTarget(viewport gd.RID) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).ViewportGetRenderTarget(viewport))
}
func (self Simple) ViewportGetTexture(viewport gd.RID) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).ViewportGetTexture(viewport))
}
func (self Simple) ViewportSetDisable3d(viewport gd.RID, disable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetDisable3d(viewport, disable)
}
func (self Simple) ViewportSetDisable2d(viewport gd.RID, disable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetDisable2d(viewport, disable)
}
func (self Simple) ViewportSetEnvironmentMode(viewport gd.RID, mode classdb.RenderingServerViewportEnvironmentMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetEnvironmentMode(viewport, mode)
}
func (self Simple) ViewportAttachCamera(viewport gd.RID, camera gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportAttachCamera(viewport, camera)
}
func (self Simple) ViewportSetScenario(viewport gd.RID, scenario gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetScenario(viewport, scenario)
}
func (self Simple) ViewportAttachCanvas(viewport gd.RID, canvas gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportAttachCanvas(viewport, canvas)
}
func (self Simple) ViewportRemoveCanvas(viewport gd.RID, canvas gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportRemoveCanvas(viewport, canvas)
}
func (self Simple) ViewportSetSnap2dTransformsToPixel(viewport gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetSnap2dTransformsToPixel(viewport, enabled)
}
func (self Simple) ViewportSetSnap2dVerticesToPixel(viewport gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetSnap2dVerticesToPixel(viewport, enabled)
}
func (self Simple) ViewportSetDefaultCanvasItemTextureFilter(viewport gd.RID, filter classdb.RenderingServerCanvasItemTextureFilter) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetDefaultCanvasItemTextureFilter(viewport, filter)
}
func (self Simple) ViewportSetDefaultCanvasItemTextureRepeat(viewport gd.RID, repeat classdb.RenderingServerCanvasItemTextureRepeat) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetDefaultCanvasItemTextureRepeat(viewport, repeat)
}
func (self Simple) ViewportSetCanvasTransform(viewport gd.RID, canvas gd.RID, offset gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetCanvasTransform(viewport, canvas, offset)
}
func (self Simple) ViewportSetCanvasStacking(viewport gd.RID, canvas gd.RID, layer int, sublayer int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetCanvasStacking(viewport, canvas, gd.Int(layer), gd.Int(sublayer))
}
func (self Simple) ViewportSetTransparentBackground(viewport gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetTransparentBackground(viewport, enabled)
}
func (self Simple) ViewportSetGlobalCanvasTransform(viewport gd.RID, transform gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetGlobalCanvasTransform(viewport, transform)
}
func (self Simple) ViewportSetSdfOversizeAndScale(viewport gd.RID, oversize classdb.RenderingServerViewportSDFOversize, scale classdb.RenderingServerViewportSDFScale) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetSdfOversizeAndScale(viewport, oversize, scale)
}
func (self Simple) ViewportSetPositionalShadowAtlasSize(viewport gd.RID, size int, use_16_bits bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetPositionalShadowAtlasSize(viewport, gd.Int(size), use_16_bits)
}
func (self Simple) ViewportSetPositionalShadowAtlasQuadrantSubdivision(viewport gd.RID, quadrant int, subdivision int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetPositionalShadowAtlasQuadrantSubdivision(viewport, gd.Int(quadrant), gd.Int(subdivision))
}
func (self Simple) ViewportSetMsaa3d(viewport gd.RID, msaa classdb.RenderingServerViewportMSAA) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetMsaa3d(viewport, msaa)
}
func (self Simple) ViewportSetMsaa2d(viewport gd.RID, msaa classdb.RenderingServerViewportMSAA) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetMsaa2d(viewport, msaa)
}
func (self Simple) ViewportSetUseHdr2d(viewport gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetUseHdr2d(viewport, enabled)
}
func (self Simple) ViewportSetScreenSpaceAa(viewport gd.RID, mode classdb.RenderingServerViewportScreenSpaceAA) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetScreenSpaceAa(viewport, mode)
}
func (self Simple) ViewportSetUseTaa(viewport gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetUseTaa(viewport, enable)
}
func (self Simple) ViewportSetUseDebanding(viewport gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetUseDebanding(viewport, enable)
}
func (self Simple) ViewportSetUseOcclusionCulling(viewport gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetUseOcclusionCulling(viewport, enable)
}
func (self Simple) ViewportSetOcclusionRaysPerThread(rays_per_thread int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetOcclusionRaysPerThread(gd.Int(rays_per_thread))
}
func (self Simple) ViewportSetOcclusionCullingBuildQuality(quality classdb.RenderingServerViewportOcclusionCullingBuildQuality) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetOcclusionCullingBuildQuality(quality)
}
func (self Simple) ViewportGetRenderInfo(viewport gd.RID, atype classdb.RenderingServerViewportRenderInfoType, info classdb.RenderingServerViewportRenderInfo) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).ViewportGetRenderInfo(viewport, atype, info)))
}
func (self Simple) ViewportSetDebugDraw(viewport gd.RID, draw classdb.RenderingServerViewportDebugDraw) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetDebugDraw(viewport, draw)
}
func (self Simple) ViewportSetMeasureRenderTime(viewport gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetMeasureRenderTime(viewport, enable)
}
func (self Simple) ViewportGetMeasuredRenderTimeCpu(viewport gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).ViewportGetMeasuredRenderTimeCpu(viewport)))
}
func (self Simple) ViewportGetMeasuredRenderTimeGpu(viewport gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).ViewportGetMeasuredRenderTimeGpu(viewport)))
}
func (self Simple) ViewportSetVrsMode(viewport gd.RID, mode classdb.RenderingServerViewportVRSMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetVrsMode(viewport, mode)
}
func (self Simple) ViewportSetVrsUpdateMode(viewport gd.RID, mode classdb.RenderingServerViewportVRSUpdateMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetVrsUpdateMode(viewport, mode)
}
func (self Simple) ViewportSetVrsTexture(viewport gd.RID, texture gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ViewportSetVrsTexture(viewport, texture)
}
func (self Simple) SkyCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).SkyCreate())
}
func (self Simple) SkySetRadianceSize(sky gd.RID, radiance_size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SkySetRadianceSize(sky, gd.Int(radiance_size))
}
func (self Simple) SkySetMode(sky gd.RID, mode classdb.RenderingServerSkyMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SkySetMode(sky, mode)
}
func (self Simple) SkySetMaterial(sky gd.RID, material gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SkySetMaterial(sky, material)
}
func (self Simple) SkyBakePanorama(sky gd.RID, energy float64, bake_irradiance bool, size gd.Vector2i) [1]classdb.Image {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Image(Expert(self).SkyBakePanorama(gc, sky, gd.Float(energy), bake_irradiance, size))
}
func (self Simple) CompositorEffectCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).CompositorEffectCreate())
}
func (self Simple) CompositorEffectSetEnabled(effect gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CompositorEffectSetEnabled(effect, enabled)
}
func (self Simple) CompositorEffectSetCallback(effect gd.RID, callback_type classdb.RenderingServerCompositorEffectCallbackType, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CompositorEffectSetCallback(effect, callback_type, callback)
}
func (self Simple) CompositorEffectSetFlag(effect gd.RID, flag classdb.RenderingServerCompositorEffectFlags, set bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CompositorEffectSetFlag(effect, flag, set)
}
func (self Simple) CompositorCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).CompositorCreate())
}
func (self Simple) CompositorSetCompositorEffects(compositor gd.RID, effects gd.ArrayOf[gd.RID]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CompositorSetCompositorEffects(compositor, effects)
}
func (self Simple) EnvironmentCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).EnvironmentCreate())
}
func (self Simple) EnvironmentSetBackground(env gd.RID, bg classdb.RenderingServerEnvironmentBG) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetBackground(env, bg)
}
func (self Simple) EnvironmentSetSky(env gd.RID, sky gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetSky(env, sky)
}
func (self Simple) EnvironmentSetSkyCustomFov(env gd.RID, scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetSkyCustomFov(env, gd.Float(scale))
}
func (self Simple) EnvironmentSetSkyOrientation(env gd.RID, orientation gd.Basis) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetSkyOrientation(env, orientation)
}
func (self Simple) EnvironmentSetBgColor(env gd.RID, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetBgColor(env, color)
}
func (self Simple) EnvironmentSetBgEnergy(env gd.RID, multiplier float64, exposure_value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetBgEnergy(env, gd.Float(multiplier), gd.Float(exposure_value))
}
func (self Simple) EnvironmentSetCanvasMaxLayer(env gd.RID, max_layer int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetCanvasMaxLayer(env, gd.Int(max_layer))
}
func (self Simple) EnvironmentSetAmbientLight(env gd.RID, color gd.Color, ambient classdb.RenderingServerEnvironmentAmbientSource, energy float64, sky_contibution float64, reflection_source classdb.RenderingServerEnvironmentReflectionSource) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetAmbientLight(env, color, ambient, gd.Float(energy), gd.Float(sky_contibution), reflection_source)
}
func (self Simple) EnvironmentSetGlow(env gd.RID, enable bool, levels gd.PackedFloat32Array, intensity float64, strength float64, mix float64, bloom_threshold float64, blend_mode classdb.RenderingServerEnvironmentGlowBlendMode, hdr_bleed_threshold float64, hdr_bleed_scale float64, hdr_luminance_cap float64, glow_map_strength float64, glow_map gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetGlow(env, enable, levels, gd.Float(intensity), gd.Float(strength), gd.Float(mix), gd.Float(bloom_threshold), blend_mode, gd.Float(hdr_bleed_threshold), gd.Float(hdr_bleed_scale), gd.Float(hdr_luminance_cap), gd.Float(glow_map_strength), glow_map)
}
func (self Simple) EnvironmentSetTonemap(env gd.RID, tone_mapper classdb.RenderingServerEnvironmentToneMapper, exposure float64, white float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetTonemap(env, tone_mapper, gd.Float(exposure), gd.Float(white))
}
func (self Simple) EnvironmentSetAdjustment(env gd.RID, enable bool, brightness float64, contrast float64, saturation float64, use_1d_color_correction bool, color_correction gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetAdjustment(env, enable, gd.Float(brightness), gd.Float(contrast), gd.Float(saturation), use_1d_color_correction, color_correction)
}
func (self Simple) EnvironmentSetSsr(env gd.RID, enable bool, max_steps int, fade_in float64, fade_out float64, depth_tolerance float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetSsr(env, enable, gd.Int(max_steps), gd.Float(fade_in), gd.Float(fade_out), gd.Float(depth_tolerance))
}
func (self Simple) EnvironmentSetSsao(env gd.RID, enable bool, radius float64, intensity float64, power float64, detail float64, horizon float64, sharpness float64, light_affect float64, ao_channel_affect float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetSsao(env, enable, gd.Float(radius), gd.Float(intensity), gd.Float(power), gd.Float(detail), gd.Float(horizon), gd.Float(sharpness), gd.Float(light_affect), gd.Float(ao_channel_affect))
}
func (self Simple) EnvironmentSetFog(env gd.RID, enable bool, light_color gd.Color, light_energy float64, sun_scatter float64, density float64, height float64, height_density float64, aerial_perspective float64, sky_affect float64, fog_mode classdb.RenderingServerEnvironmentFogMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetFog(env, enable, light_color, gd.Float(light_energy), gd.Float(sun_scatter), gd.Float(density), gd.Float(height), gd.Float(height_density), gd.Float(aerial_perspective), gd.Float(sky_affect), fog_mode)
}
func (self Simple) EnvironmentSetSdfgi(env gd.RID, enable bool, cascades int, min_cell_size float64, y_scale classdb.RenderingServerEnvironmentSDFGIYScale, use_occlusion bool, bounce_feedback float64, read_sky bool, energy float64, normal_bias float64, probe_bias float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetSdfgi(env, enable, gd.Int(cascades), gd.Float(min_cell_size), y_scale, use_occlusion, gd.Float(bounce_feedback), read_sky, gd.Float(energy), gd.Float(normal_bias), gd.Float(probe_bias))
}
func (self Simple) EnvironmentSetVolumetricFog(env gd.RID, enable bool, density float64, albedo gd.Color, emission gd.Color, emission_energy float64, anisotropy float64, length float64, p_detail_spread float64, gi_inject float64, temporal_reprojection bool, temporal_reprojection_amount float64, ambient_inject float64, sky_affect float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetVolumetricFog(env, enable, gd.Float(density), albedo, emission, gd.Float(emission_energy), gd.Float(anisotropy), gd.Float(length), gd.Float(p_detail_spread), gd.Float(gi_inject), temporal_reprojection, gd.Float(temporal_reprojection_amount), gd.Float(ambient_inject), gd.Float(sky_affect))
}
func (self Simple) EnvironmentGlowSetUseBicubicUpscale(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentGlowSetUseBicubicUpscale(enable)
}
func (self Simple) EnvironmentSetSsrRoughnessQuality(quality classdb.RenderingServerEnvironmentSSRRoughnessQuality) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetSsrRoughnessQuality(quality)
}
func (self Simple) EnvironmentSetSsaoQuality(quality classdb.RenderingServerEnvironmentSSAOQuality, half_size bool, adaptive_target float64, blur_passes int, fadeout_from float64, fadeout_to float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetSsaoQuality(quality, half_size, gd.Float(adaptive_target), gd.Int(blur_passes), gd.Float(fadeout_from), gd.Float(fadeout_to))
}
func (self Simple) EnvironmentSetSsilQuality(quality classdb.RenderingServerEnvironmentSSILQuality, half_size bool, adaptive_target float64, blur_passes int, fadeout_from float64, fadeout_to float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetSsilQuality(quality, half_size, gd.Float(adaptive_target), gd.Int(blur_passes), gd.Float(fadeout_from), gd.Float(fadeout_to))
}
func (self Simple) EnvironmentSetSdfgiRayCount(ray_count classdb.RenderingServerEnvironmentSDFGIRayCount) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetSdfgiRayCount(ray_count)
}
func (self Simple) EnvironmentSetSdfgiFramesToConverge(frames classdb.RenderingServerEnvironmentSDFGIFramesToConverge) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetSdfgiFramesToConverge(frames)
}
func (self Simple) EnvironmentSetSdfgiFramesToUpdateLight(frames classdb.RenderingServerEnvironmentSDFGIFramesToUpdateLight) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetSdfgiFramesToUpdateLight(frames)
}
func (self Simple) EnvironmentSetVolumetricFogVolumeSize(size int, depth int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetVolumetricFogVolumeSize(gd.Int(size), gd.Int(depth))
}
func (self Simple) EnvironmentSetVolumetricFogFilterActive(active bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnvironmentSetVolumetricFogFilterActive(active)
}
func (self Simple) EnvironmentBakePanorama(environment gd.RID, bake_irradiance bool, size gd.Vector2i) [1]classdb.Image {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Image(Expert(self).EnvironmentBakePanorama(gc, environment, bake_irradiance, size))
}
func (self Simple) ScreenSpaceRoughnessLimiterSetActive(enable bool, amount float64, limit float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ScreenSpaceRoughnessLimiterSetActive(enable, gd.Float(amount), gd.Float(limit))
}
func (self Simple) SubSurfaceScatteringSetQuality(quality classdb.RenderingServerSubSurfaceScatteringQuality) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SubSurfaceScatteringSetQuality(quality)
}
func (self Simple) SubSurfaceScatteringSetScale(scale float64, depth_scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SubSurfaceScatteringSetScale(gd.Float(scale), gd.Float(depth_scale))
}
func (self Simple) CameraAttributesCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).CameraAttributesCreate())
}
func (self Simple) CameraAttributesSetDofBlurQuality(quality classdb.RenderingServerDOFBlurQuality, use_jitter bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CameraAttributesSetDofBlurQuality(quality, use_jitter)
}
func (self Simple) CameraAttributesSetDofBlurBokehShape(shape classdb.RenderingServerDOFBokehShape) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CameraAttributesSetDofBlurBokehShape(shape)
}
func (self Simple) CameraAttributesSetDofBlur(camera_attributes gd.RID, far_enable bool, far_distance float64, far_transition float64, near_enable bool, near_distance float64, near_transition float64, amount float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CameraAttributesSetDofBlur(camera_attributes, far_enable, gd.Float(far_distance), gd.Float(far_transition), near_enable, gd.Float(near_distance), gd.Float(near_transition), gd.Float(amount))
}
func (self Simple) CameraAttributesSetExposure(camera_attributes gd.RID, multiplier float64, normalization float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CameraAttributesSetExposure(camera_attributes, gd.Float(multiplier), gd.Float(normalization))
}
func (self Simple) CameraAttributesSetAutoExposure(camera_attributes gd.RID, enable bool, min_sensitivity float64, max_sensitivity float64, speed float64, scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CameraAttributesSetAutoExposure(camera_attributes, enable, gd.Float(min_sensitivity), gd.Float(max_sensitivity), gd.Float(speed), gd.Float(scale))
}
func (self Simple) ScenarioCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).ScenarioCreate())
}
func (self Simple) ScenarioSetEnvironment(scenario gd.RID, environment gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ScenarioSetEnvironment(scenario, environment)
}
func (self Simple) ScenarioSetFallbackEnvironment(scenario gd.RID, environment gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ScenarioSetFallbackEnvironment(scenario, environment)
}
func (self Simple) ScenarioSetCameraAttributes(scenario gd.RID, effects gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ScenarioSetCameraAttributes(scenario, effects)
}
func (self Simple) ScenarioSetCompositor(scenario gd.RID, compositor gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ScenarioSetCompositor(scenario, compositor)
}
func (self Simple) InstanceCreate2(base gd.RID, scenario gd.RID) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).InstanceCreate2(base, scenario))
}
func (self Simple) InstanceCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).InstanceCreate())
}
func (self Simple) InstanceSetBase(instance gd.RID, base gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceSetBase(instance, base)
}
func (self Simple) InstanceSetScenario(instance gd.RID, scenario gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceSetScenario(instance, scenario)
}
func (self Simple) InstanceSetLayerMask(instance gd.RID, mask int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceSetLayerMask(instance, gd.Int(mask))
}
func (self Simple) InstanceSetPivotData(instance gd.RID, sorting_offset float64, use_aabb_center bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceSetPivotData(instance, gd.Float(sorting_offset), use_aabb_center)
}
func (self Simple) InstanceSetTransform(instance gd.RID, transform gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceSetTransform(instance, transform)
}
func (self Simple) InstanceAttachObjectInstanceId(instance gd.RID, id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceAttachObjectInstanceId(instance, gd.Int(id))
}
func (self Simple) InstanceSetBlendShapeWeight(instance gd.RID, shape int, weight float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceSetBlendShapeWeight(instance, gd.Int(shape), gd.Float(weight))
}
func (self Simple) InstanceSetSurfaceOverrideMaterial(instance gd.RID, surface int, material gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceSetSurfaceOverrideMaterial(instance, gd.Int(surface), material)
}
func (self Simple) InstanceSetVisible(instance gd.RID, visible bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceSetVisible(instance, visible)
}
func (self Simple) InstanceGeometrySetTransparency(instance gd.RID, transparency float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceGeometrySetTransparency(instance, gd.Float(transparency))
}
func (self Simple) InstanceSetCustomAabb(instance gd.RID, aabb gd.AABB) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceSetCustomAabb(instance, aabb)
}
func (self Simple) InstanceAttachSkeleton(instance gd.RID, skeleton gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceAttachSkeleton(instance, skeleton)
}
func (self Simple) InstanceSetExtraVisibilityMargin(instance gd.RID, margin float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceSetExtraVisibilityMargin(instance, gd.Float(margin))
}
func (self Simple) InstanceSetVisibilityParent(instance gd.RID, parent gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceSetVisibilityParent(instance, parent)
}
func (self Simple) InstanceSetIgnoreCulling(instance gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceSetIgnoreCulling(instance, enabled)
}
func (self Simple) InstanceGeometrySetFlag(instance gd.RID, flag classdb.RenderingServerInstanceFlags, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceGeometrySetFlag(instance, flag, enabled)
}
func (self Simple) InstanceGeometrySetCastShadowsSetting(instance gd.RID, shadow_casting_setting classdb.RenderingServerShadowCastingSetting) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceGeometrySetCastShadowsSetting(instance, shadow_casting_setting)
}
func (self Simple) InstanceGeometrySetMaterialOverride(instance gd.RID, material gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceGeometrySetMaterialOverride(instance, material)
}
func (self Simple) InstanceGeometrySetMaterialOverlay(instance gd.RID, material gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceGeometrySetMaterialOverlay(instance, material)
}
func (self Simple) InstanceGeometrySetVisibilityRange(instance gd.RID, min float64, max float64, min_margin float64, max_margin float64, fade_mode classdb.RenderingServerVisibilityRangeFadeMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceGeometrySetVisibilityRange(instance, gd.Float(min), gd.Float(max), gd.Float(min_margin), gd.Float(max_margin), fade_mode)
}
func (self Simple) InstanceGeometrySetLightmap(instance gd.RID, lightmap gd.RID, lightmap_uv_scale gd.Rect2, lightmap_slice int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceGeometrySetLightmap(instance, lightmap, lightmap_uv_scale, gd.Int(lightmap_slice))
}
func (self Simple) InstanceGeometrySetLodBias(instance gd.RID, lod_bias float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceGeometrySetLodBias(instance, gd.Float(lod_bias))
}
func (self Simple) InstanceGeometrySetShaderParameter(instance gd.RID, parameter string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstanceGeometrySetShaderParameter(instance, gc.StringName(parameter), value)
}
func (self Simple) InstanceGeometryGetShaderParameter(instance gd.RID, parameter string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).InstanceGeometryGetShaderParameter(gc, instance, gc.StringName(parameter)))
}
func (self Simple) InstanceGeometryGetShaderParameterDefaultValue(instance gd.RID, parameter string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).InstanceGeometryGetShaderParameterDefaultValue(gc, instance, gc.StringName(parameter)))
}
func (self Simple) InstanceGeometryGetShaderParameterList(instance gd.RID) gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](Expert(self).InstanceGeometryGetShaderParameterList(gc, instance))
}
func (self Simple) InstancesCullAabb(aabb gd.AABB, scenario gd.RID) gd.PackedInt64Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt64Array(Expert(self).InstancesCullAabb(gc, aabb, scenario))
}
func (self Simple) InstancesCullRay(from gd.Vector3, to gd.Vector3, scenario gd.RID) gd.PackedInt64Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt64Array(Expert(self).InstancesCullRay(gc, from, to, scenario))
}
func (self Simple) InstancesCullConvex(convex gd.ArrayOf[gd.Plane], scenario gd.RID) gd.PackedInt64Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt64Array(Expert(self).InstancesCullConvex(gc, convex, scenario))
}
func (self Simple) BakeRenderUv2(base gd.RID, material_overrides gd.ArrayOf[gd.RID], image_size gd.Vector2i) gd.ArrayOf[[1]classdb.Image] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.Image](Expert(self).BakeRenderUv2(gc, base, material_overrides, image_size))
}
func (self Simple) CanvasCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).CanvasCreate())
}
func (self Simple) CanvasSetItemMirroring(canvas gd.RID, item gd.RID, mirroring gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasSetItemMirroring(canvas, item, mirroring)
}
func (self Simple) CanvasSetItemRepeat(item gd.RID, repeat_size gd.Vector2, repeat_times int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasSetItemRepeat(item, repeat_size, gd.Int(repeat_times))
}
func (self Simple) CanvasSetModulate(canvas gd.RID, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasSetModulate(canvas, color)
}
func (self Simple) CanvasSetDisableScale(disable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasSetDisableScale(disable)
}
func (self Simple) CanvasTextureCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).CanvasTextureCreate())
}
func (self Simple) CanvasTextureSetChannel(canvas_texture gd.RID, channel classdb.RenderingServerCanvasTextureChannel, texture gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasTextureSetChannel(canvas_texture, channel, texture)
}
func (self Simple) CanvasTextureSetShadingParameters(canvas_texture gd.RID, base_color gd.Color, shininess float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasTextureSetShadingParameters(canvas_texture, base_color, gd.Float(shininess))
}
func (self Simple) CanvasTextureSetTextureFilter(canvas_texture gd.RID, filter classdb.RenderingServerCanvasItemTextureFilter) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasTextureSetTextureFilter(canvas_texture, filter)
}
func (self Simple) CanvasTextureSetTextureRepeat(canvas_texture gd.RID, repeat classdb.RenderingServerCanvasItemTextureRepeat) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasTextureSetTextureRepeat(canvas_texture, repeat)
}
func (self Simple) CanvasItemCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).CanvasItemCreate())
}
func (self Simple) CanvasItemSetParent(item gd.RID, parent gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetParent(item, parent)
}
func (self Simple) CanvasItemSetDefaultTextureFilter(item gd.RID, filter classdb.RenderingServerCanvasItemTextureFilter) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetDefaultTextureFilter(item, filter)
}
func (self Simple) CanvasItemSetDefaultTextureRepeat(item gd.RID, repeat classdb.RenderingServerCanvasItemTextureRepeat) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetDefaultTextureRepeat(item, repeat)
}
func (self Simple) CanvasItemSetVisible(item gd.RID, visible bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetVisible(item, visible)
}
func (self Simple) CanvasItemSetLightMask(item gd.RID, mask int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetLightMask(item, gd.Int(mask))
}
func (self Simple) CanvasItemSetVisibilityLayer(item gd.RID, visibility_layer int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetVisibilityLayer(item, gd.Int(visibility_layer))
}
func (self Simple) CanvasItemSetTransform(item gd.RID, transform gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetTransform(item, transform)
}
func (self Simple) CanvasItemSetClip(item gd.RID, clip bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetClip(item, clip)
}
func (self Simple) CanvasItemSetDistanceFieldMode(item gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetDistanceFieldMode(item, enabled)
}
func (self Simple) CanvasItemSetCustomRect(item gd.RID, use_custom_rect bool, rect gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetCustomRect(item, use_custom_rect, rect)
}
func (self Simple) CanvasItemSetModulate(item gd.RID, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetModulate(item, color)
}
func (self Simple) CanvasItemSetSelfModulate(item gd.RID, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetSelfModulate(item, color)
}
func (self Simple) CanvasItemSetDrawBehindParent(item gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetDrawBehindParent(item, enabled)
}
func (self Simple) CanvasItemSetInterpolated(item gd.RID, interpolated bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetInterpolated(item, interpolated)
}
func (self Simple) CanvasItemResetPhysicsInterpolation(item gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemResetPhysicsInterpolation(item)
}
func (self Simple) CanvasItemTransformPhysicsInterpolation(item gd.RID, transform gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemTransformPhysicsInterpolation(item, transform)
}
func (self Simple) CanvasItemAddLine(item gd.RID, from gd.Vector2, to gd.Vector2, color gd.Color, width float64, antialiased bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemAddLine(item, from, to, color, gd.Float(width), antialiased)
}
func (self Simple) CanvasItemAddPolyline(item gd.RID, points gd.PackedVector2Array, colors gd.PackedColorArray, width float64, antialiased bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemAddPolyline(item, points, colors, gd.Float(width), antialiased)
}
func (self Simple) CanvasItemAddMultiline(item gd.RID, points gd.PackedVector2Array, colors gd.PackedColorArray, width float64, antialiased bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemAddMultiline(item, points, colors, gd.Float(width), antialiased)
}
func (self Simple) CanvasItemAddRect(item gd.RID, rect gd.Rect2, color gd.Color, antialiased bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemAddRect(item, rect, color, antialiased)
}
func (self Simple) CanvasItemAddCircle(item gd.RID, pos gd.Vector2, radius float64, color gd.Color, antialiased bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemAddCircle(item, pos, gd.Float(radius), color, antialiased)
}
func (self Simple) CanvasItemAddTextureRect(item gd.RID, rect gd.Rect2, texture gd.RID, tile bool, modulate gd.Color, transpose bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemAddTextureRect(item, rect, texture, tile, modulate, transpose)
}
func (self Simple) CanvasItemAddMsdfTextureRectRegion(item gd.RID, rect gd.Rect2, texture gd.RID, src_rect gd.Rect2, modulate gd.Color, outline_size int, px_range float64, scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemAddMsdfTextureRectRegion(item, rect, texture, src_rect, modulate, gd.Int(outline_size), gd.Float(px_range), gd.Float(scale))
}
func (self Simple) CanvasItemAddLcdTextureRectRegion(item gd.RID, rect gd.Rect2, texture gd.RID, src_rect gd.Rect2, modulate gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemAddLcdTextureRectRegion(item, rect, texture, src_rect, modulate)
}
func (self Simple) CanvasItemAddTextureRectRegion(item gd.RID, rect gd.Rect2, texture gd.RID, src_rect gd.Rect2, modulate gd.Color, transpose bool, clip_uv bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemAddTextureRectRegion(item, rect, texture, src_rect, modulate, transpose, clip_uv)
}
func (self Simple) CanvasItemAddNinePatch(item gd.RID, rect gd.Rect2, source gd.Rect2, texture gd.RID, topleft gd.Vector2, bottomright gd.Vector2, x_axis_mode classdb.RenderingServerNinePatchAxisMode, y_axis_mode classdb.RenderingServerNinePatchAxisMode, draw_center bool, modulate gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemAddNinePatch(item, rect, source, texture, topleft, bottomright, x_axis_mode, y_axis_mode, draw_center, modulate)
}
func (self Simple) CanvasItemAddPrimitive(item gd.RID, points gd.PackedVector2Array, colors gd.PackedColorArray, uvs gd.PackedVector2Array, texture gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemAddPrimitive(item, points, colors, uvs, texture)
}
func (self Simple) CanvasItemAddPolygon(item gd.RID, points gd.PackedVector2Array, colors gd.PackedColorArray, uvs gd.PackedVector2Array, texture gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemAddPolygon(item, points, colors, uvs, texture)
}
func (self Simple) CanvasItemAddTriangleArray(item gd.RID, indices gd.PackedInt32Array, points gd.PackedVector2Array, colors gd.PackedColorArray, uvs gd.PackedVector2Array, bones gd.PackedInt32Array, weights gd.PackedFloat32Array, texture gd.RID, count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemAddTriangleArray(item, indices, points, colors, uvs, bones, weights, texture, gd.Int(count))
}
func (self Simple) CanvasItemAddMesh(item gd.RID, mesh gd.RID, transform gd.Transform2D, modulate gd.Color, texture gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemAddMesh(item, mesh, transform, modulate, texture)
}
func (self Simple) CanvasItemAddMultimesh(item gd.RID, mesh gd.RID, texture gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemAddMultimesh(item, mesh, texture)
}
func (self Simple) CanvasItemAddParticles(item gd.RID, particles gd.RID, texture gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemAddParticles(item, particles, texture)
}
func (self Simple) CanvasItemAddSetTransform(item gd.RID, transform gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemAddSetTransform(item, transform)
}
func (self Simple) CanvasItemAddClipIgnore(item gd.RID, ignore bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemAddClipIgnore(item, ignore)
}
func (self Simple) CanvasItemAddAnimationSlice(item gd.RID, animation_length float64, slice_begin float64, slice_end float64, offset float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemAddAnimationSlice(item, gd.Float(animation_length), gd.Float(slice_begin), gd.Float(slice_end), gd.Float(offset))
}
func (self Simple) CanvasItemSetSortChildrenByY(item gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetSortChildrenByY(item, enabled)
}
func (self Simple) CanvasItemSetZIndex(item gd.RID, z_index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetZIndex(item, gd.Int(z_index))
}
func (self Simple) CanvasItemSetZAsRelativeToParent(item gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetZAsRelativeToParent(item, enabled)
}
func (self Simple) CanvasItemSetCopyToBackbuffer(item gd.RID, enabled bool, rect gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetCopyToBackbuffer(item, enabled, rect)
}
func (self Simple) CanvasItemClear(item gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemClear(item)
}
func (self Simple) CanvasItemSetDrawIndex(item gd.RID, index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetDrawIndex(item, gd.Int(index))
}
func (self Simple) CanvasItemSetMaterial(item gd.RID, material gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetMaterial(item, material)
}
func (self Simple) CanvasItemSetUseParentMaterial(item gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetUseParentMaterial(item, enabled)
}
func (self Simple) CanvasItemSetVisibilityNotifier(item gd.RID, enable bool, area gd.Rect2, enter_callable gd.Callable, exit_callable gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetVisibilityNotifier(item, enable, area, enter_callable, exit_callable)
}
func (self Simple) CanvasItemSetCanvasGroupMode(item gd.RID, mode classdb.RenderingServerCanvasGroupMode, clear_margin float64, fit_empty bool, fit_margin float64, blur_mipmaps bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasItemSetCanvasGroupMode(item, mode, gd.Float(clear_margin), fit_empty, gd.Float(fit_margin), blur_mipmaps)
}
func (self Simple) DebugCanvasItemGetRect(item gd.RID) gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(Expert(self).DebugCanvasItemGetRect(item))
}
func (self Simple) CanvasLightCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).CanvasLightCreate())
}
func (self Simple) CanvasLightAttachToCanvas(light gd.RID, canvas gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightAttachToCanvas(light, canvas)
}
func (self Simple) CanvasLightSetEnabled(light gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightSetEnabled(light, enabled)
}
func (self Simple) CanvasLightSetTextureScale(light gd.RID, scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightSetTextureScale(light, gd.Float(scale))
}
func (self Simple) CanvasLightSetTransform(light gd.RID, transform gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightSetTransform(light, transform)
}
func (self Simple) CanvasLightSetTexture(light gd.RID, texture gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightSetTexture(light, texture)
}
func (self Simple) CanvasLightSetTextureOffset(light gd.RID, offset gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightSetTextureOffset(light, offset)
}
func (self Simple) CanvasLightSetColor(light gd.RID, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightSetColor(light, color)
}
func (self Simple) CanvasLightSetHeight(light gd.RID, height float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightSetHeight(light, gd.Float(height))
}
func (self Simple) CanvasLightSetEnergy(light gd.RID, energy float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightSetEnergy(light, gd.Float(energy))
}
func (self Simple) CanvasLightSetZRange(light gd.RID, min_z int, max_z int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightSetZRange(light, gd.Int(min_z), gd.Int(max_z))
}
func (self Simple) CanvasLightSetLayerRange(light gd.RID, min_layer int, max_layer int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightSetLayerRange(light, gd.Int(min_layer), gd.Int(max_layer))
}
func (self Simple) CanvasLightSetItemCullMask(light gd.RID, mask int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightSetItemCullMask(light, gd.Int(mask))
}
func (self Simple) CanvasLightSetItemShadowCullMask(light gd.RID, mask int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightSetItemShadowCullMask(light, gd.Int(mask))
}
func (self Simple) CanvasLightSetMode(light gd.RID, mode classdb.RenderingServerCanvasLightMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightSetMode(light, mode)
}
func (self Simple) CanvasLightSetShadowEnabled(light gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightSetShadowEnabled(light, enabled)
}
func (self Simple) CanvasLightSetShadowFilter(light gd.RID, filter classdb.RenderingServerCanvasLightShadowFilter) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightSetShadowFilter(light, filter)
}
func (self Simple) CanvasLightSetShadowColor(light gd.RID, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightSetShadowColor(light, color)
}
func (self Simple) CanvasLightSetShadowSmooth(light gd.RID, smooth float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightSetShadowSmooth(light, gd.Float(smooth))
}
func (self Simple) CanvasLightSetBlendMode(light gd.RID, mode classdb.RenderingServerCanvasLightBlendMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightSetBlendMode(light, mode)
}
func (self Simple) CanvasLightSetInterpolated(light gd.RID, interpolated bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightSetInterpolated(light, interpolated)
}
func (self Simple) CanvasLightResetPhysicsInterpolation(light gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightResetPhysicsInterpolation(light)
}
func (self Simple) CanvasLightTransformPhysicsInterpolation(light gd.RID, transform gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightTransformPhysicsInterpolation(light, transform)
}
func (self Simple) CanvasLightOccluderCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).CanvasLightOccluderCreate())
}
func (self Simple) CanvasLightOccluderAttachToCanvas(occluder gd.RID, canvas gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightOccluderAttachToCanvas(occluder, canvas)
}
func (self Simple) CanvasLightOccluderSetEnabled(occluder gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightOccluderSetEnabled(occluder, enabled)
}
func (self Simple) CanvasLightOccluderSetPolygon(occluder gd.RID, polygon gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightOccluderSetPolygon(occluder, polygon)
}
func (self Simple) CanvasLightOccluderSetAsSdfCollision(occluder gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightOccluderSetAsSdfCollision(occluder, enable)
}
func (self Simple) CanvasLightOccluderSetTransform(occluder gd.RID, transform gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightOccluderSetTransform(occluder, transform)
}
func (self Simple) CanvasLightOccluderSetLightMask(occluder gd.RID, mask int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightOccluderSetLightMask(occluder, gd.Int(mask))
}
func (self Simple) CanvasLightOccluderSetInterpolated(occluder gd.RID, interpolated bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightOccluderSetInterpolated(occluder, interpolated)
}
func (self Simple) CanvasLightOccluderResetPhysicsInterpolation(occluder gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightOccluderResetPhysicsInterpolation(occluder)
}
func (self Simple) CanvasLightOccluderTransformPhysicsInterpolation(occluder gd.RID, transform gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasLightOccluderTransformPhysicsInterpolation(occluder, transform)
}
func (self Simple) CanvasOccluderPolygonCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).CanvasOccluderPolygonCreate())
}
func (self Simple) CanvasOccluderPolygonSetShape(occluder_polygon gd.RID, shape gd.PackedVector2Array, closed bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasOccluderPolygonSetShape(occluder_polygon, shape, closed)
}
func (self Simple) CanvasOccluderPolygonSetCullMode(occluder_polygon gd.RID, mode classdb.RenderingServerCanvasOccluderPolygonCullMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasOccluderPolygonSetCullMode(occluder_polygon, mode)
}
func (self Simple) CanvasSetShadowTextureSize(size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CanvasSetShadowTextureSize(gd.Int(size))
}
func (self Simple) GlobalShaderParameterAdd(name string, atype classdb.RenderingServerGlobalShaderParameterType, default_value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalShaderParameterAdd(gc.StringName(name), atype, default_value)
}
func (self Simple) GlobalShaderParameterRemove(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalShaderParameterRemove(gc.StringName(name))
}
func (self Simple) GlobalShaderParameterGetList() gd.ArrayOf[gd.StringName] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.StringName](Expert(self).GlobalShaderParameterGetList(gc))
}
func (self Simple) GlobalShaderParameterSet(name string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalShaderParameterSet(gc.StringName(name), value)
}
func (self Simple) GlobalShaderParameterSetOverride(name string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GlobalShaderParameterSetOverride(gc.StringName(name), value)
}
func (self Simple) GlobalShaderParameterGet(name string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GlobalShaderParameterGet(gc, gc.StringName(name)))
}
func (self Simple) GlobalShaderParameterGetType(name string) classdb.RenderingServerGlobalShaderParameterType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingServerGlobalShaderParameterType(Expert(self).GlobalShaderParameterGetType(gc.StringName(name)))
}
func (self Simple) FreeRid(rid gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).FreeRid(rid)
}
func (self Simple) RequestFrameDrawnCallback(callable gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RequestFrameDrawnCallback(callable)
}
func (self Simple) HasChanged() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasChanged())
}
func (self Simple) GetRenderingInfo(info classdb.RenderingServerRenderingInfo) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetRenderingInfo(info)))
}
func (self Simple) GetVideoAdapterName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetVideoAdapterName(gc).String())
}
func (self Simple) GetVideoAdapterVendor() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetVideoAdapterVendor(gc).String())
}
func (self Simple) GetVideoAdapterType() classdb.RenderingDeviceDeviceType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceDeviceType(Expert(self).GetVideoAdapterType())
}
func (self Simple) GetVideoAdapterApiVersion() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetVideoAdapterApiVersion(gc).String())
}
func (self Simple) MakeSphereMesh(latitudes int, longitudes int, radius float64) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).MakeSphereMesh(gd.Int(latitudes), gd.Int(longitudes), gd.Float(radius)))
}
func (self Simple) GetTestCube() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetTestCube())
}
func (self Simple) GetTestTexture() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetTestTexture())
}
func (self Simple) GetWhiteTexture() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetWhiteTexture())
}
func (self Simple) SetBootImage(image [1]classdb.Image, color gd.Color, scale bool, use_filter bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBootImage(image, color, scale, use_filter)
}
func (self Simple) GetDefaultClearColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetDefaultClearColor())
}
func (self Simple) SetDefaultClearColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDefaultClearColor(color)
}
func (self Simple) HasOsFeature(feature string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasOsFeature(gc.String(feature)))
}
func (self Simple) SetDebugGenerateWireframes(generate bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDebugGenerateWireframes(generate)
}
func (self Simple) IsRenderLoopEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsRenderLoopEnabled())
}
func (self Simple) SetRenderLoopEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRenderLoopEnabled(enabled)
}
func (self Simple) GetFrameSetupTimeCpu() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFrameSetupTimeCpu()))
}
func (self Simple) ForceSync() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ForceSync()
}
func (self Simple) ForceDraw(swap_buffers bool, frame_step float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ForceDraw(swap_buffers, gd.Float(frame_step))
}
func (self Simple) GetRenderingDevice() [1]classdb.RenderingDevice {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.RenderingDevice(Expert(self).GetRenderingDevice(gc))
}
func (self Simple) CreateLocalRenderingDevice() [1]classdb.RenderingDevice {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.RenderingDevice(Expert(self).CreateLocalRenderingDevice(gc))
}
func (self Simple) IsOnRenderThread() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsOnRenderThread())
}
func (self Simple) CallOnRenderThread(callable gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CallOnRenderThread(callable)
}
func (self Simple) HasFeature(feature classdb.RenderingServerFeatures) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasFeature(feature))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RenderingServer
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Creates a 2-dimensional texture and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]texture_2d_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
[b]Note:[/b] The equivalent resource is [Texture2D].
[b]Note:[/b] Not to be confused with [method RenderingDevice.texture_create], which creates the graphics API's own texture type as opposed to the Godot-specific [Texture2D] resource.
*/
//go:nosplit
func (self class) Texture2dCreate(image object.Image) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(image[0].AsPointer())[0])
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_2d_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a 2-dimensional layered texture and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]texture_2d_layered_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
[b]Note:[/b] The equivalent resource is [TextureLayered].
*/
//go:nosplit
func (self class) Texture2dLayeredCreate(layers gd.ArrayOf[object.Image], layered_type classdb.RenderingServerTextureLayeredType) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(layers))
	callframe.Arg(frame, layered_type)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_2d_layered_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
[b]Note:[/b] The equivalent resource is [Texture3D].
*/
//go:nosplit
func (self class) Texture3dCreate(format classdb.ImageFormat, width gd.Int, height gd.Int, depth gd.Int, mipmaps bool, data gd.ArrayOf[object.Image]) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, format)
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, depth)
	callframe.Arg(frame, mipmaps)
	callframe.Arg(frame, mmm.Get(data))
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_3d_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
This method does nothing and always returns an invalid [RID].
*/
//go:nosplit
func (self class) TextureProxyCreate(base gd.RID) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, base)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_proxy_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Updates the texture specified by the [param texture] [RID] with the data in [param image]. A [param layer] must also be specified, which should be [code]0[/code] when updating a single-layer texture ([Texture2D]).
[b]Note:[/b] The [param image] must have the same width, height and format as the current [param texture] data. Otherwise, an error will be printed and the original texture won't be modified. If you need to use different width, height or format, use [method texture_replace] instead.
*/
//go:nosplit
func (self class) Texture2dUpdate(texture gd.RID, image object.Image, layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	callframe.Arg(frame, mmm.Get(image[0].AsPointer())[0])
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_2d_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Updates the texture specified by the [param texture] [RID]'s data with the data in [param data]. All the texture's layers must be replaced at once.
[b]Note:[/b] The [param texture] must have the same width, height, depth and format as the current texture data. Otherwise, an error will be printed and the original texture won't be modified. If you need to use different width, height, depth or format, use [method texture_replace] instead.
*/
//go:nosplit
func (self class) Texture3dUpdate(texture gd.RID, data gd.ArrayOf[object.Image])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	callframe.Arg(frame, mmm.Get(data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_3d_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
This method does nothing.
*/
//go:nosplit
func (self class) TextureProxyUpdate(texture gd.RID, proxy_to gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	callframe.Arg(frame, proxy_to)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_proxy_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a placeholder for a 2-dimensional layered texture and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]texture_2d_layered_*[/code] RenderingServer functions, although it does nothing when used. See also [method texture_2d_layered_placeholder_create]
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
[b]Note:[/b] The equivalent resource is [PlaceholderTexture2D].
*/
//go:nosplit
func (self class) Texture2dPlaceholderCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_2d_placeholder_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a placeholder for a 2-dimensional layered texture and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]texture_2d_layered_*[/code] RenderingServer functions, although it does nothing when used. See also [method texture_2d_placeholder_create].
[b]Note:[/b] The equivalent resource is [PlaceholderTextureLayered].
*/
//go:nosplit
func (self class) Texture2dLayeredPlaceholderCreate(layered_type classdb.RenderingServerTextureLayeredType) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layered_type)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_2d_layered_placeholder_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a placeholder for a 3-dimensional texture and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]texture_3d_*[/code] RenderingServer functions, although it does nothing when used.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
[b]Note:[/b] The equivalent resource is [PlaceholderTexture3D].
*/
//go:nosplit
func (self class) Texture3dPlaceholderCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_3d_placeholder_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an [Image] instance from the given [param texture] [RID].
Example of getting the test texture from [method get_test_texture] and applying it to a [Sprite2D] node:
[codeblock]
var texture_rid = RenderingServer.get_test_texture()
var texture = ImageTexture.create_from_image(RenderingServer.texture_2d_get(texture_rid))
$Sprite2D.texture = texture
[/codeblock]
*/
//go:nosplit
func (self class) Texture2dGet(ctx gd.Lifetime, texture gd.RID) object.Image {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_2d_get, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns an [Image] instance from the given [param texture] [RID] and [param layer].
*/
//go:nosplit
func (self class) Texture2dLayerGet(ctx gd.Lifetime, texture gd.RID, layer gd.Int) object.Image {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_2d_layer_get, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns 3D texture data as an array of [Image]s for the specified texture [RID].
*/
//go:nosplit
func (self class) Texture3dGet(ctx gd.Lifetime, texture gd.RID) gd.ArrayOf[object.Image] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_3d_get, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.Image](ret)
}
/*
Replaces [param texture]'s texture data by the texture specified by the [param by_texture] RID, without changing [param texture]'s RID.
*/
//go:nosplit
func (self class) TextureReplace(texture gd.RID, by_texture gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	callframe.Arg(frame, by_texture)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_replace, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) TextureSetSizeOverride(texture gd.RID, width gd.Int, height gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_set_size_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) TextureSetPath(texture gd.RID, path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_set_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) TextureGetPath(ctx gd.Lifetime, texture gd.RID) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_get_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the format for the texture.
*/
//go:nosplit
func (self class) TextureGetFormat(texture gd.RID) classdb.ImageFormat {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	var r_ret = callframe.Ret[classdb.ImageFormat](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) TextureSetForceRedrawIfVisible(texture gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_set_force_redraw_if_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a new texture object based on a texture created directly on the [RenderingDevice]. If the texture contains layers, [param layer_type] is used to define the layer type.
*/
//go:nosplit
func (self class) TextureRdCreate(rd_texture gd.RID, layer_type classdb.RenderingServerTextureLayeredType) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rd_texture)
	callframe.Arg(frame, layer_type)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_rd_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a texture [RID] that can be used with [RenderingDevice].
*/
//go:nosplit
func (self class) TextureGetRdTexture(texture gd.RID, srgb bool) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	callframe.Arg(frame, srgb)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_get_rd_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the internal graphics handle for this texture object. For use when communicating with third-party APIs mostly with GDExtension.
[b]Note:[/b] This function returns a [code]uint64_t[/code] which internally maps to a [code]GLuint[/code] (OpenGL) or [code]VkImage[/code] (Vulkan).
*/
//go:nosplit
func (self class) TextureGetNativeHandle(texture gd.RID, srgb bool) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture)
	callframe.Arg(frame, srgb)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_texture_get_native_handle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates an empty shader and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]shader_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
[b]Note:[/b] The equivalent resource is [Shader].
*/
//go:nosplit
func (self class) ShaderCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_shader_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the shader's source code (which triggers recompilation after being changed).
*/
//go:nosplit
func (self class) ShaderSetCode(shader gd.RID, code gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shader)
	callframe.Arg(frame, mmm.Get(code))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_shader_set_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the path hint for the specified shader. This should generally match the [Shader] resource's [member Resource.resource_path].
*/
//go:nosplit
func (self class) ShaderSetPathHint(shader gd.RID, path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shader)
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_shader_set_path_hint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a shader's source code as a string.
*/
//go:nosplit
func (self class) ShaderGetCode(ctx gd.Lifetime, shader gd.RID) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shader)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_shader_get_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the parameters of a shader.
*/
//go:nosplit
func (self class) GetShaderParameterList(ctx gd.Lifetime, shader gd.RID) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shader)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_get_shader_parameter_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Returns the default value for the specified shader uniform. This is usually the value written in the shader source code.
*/
//go:nosplit
func (self class) ShaderGetParameterDefault(ctx gd.Lifetime, shader gd.RID, name gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shader)
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_shader_get_parameter_default, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets a shader's default texture. Overwrites the texture given by name.
[b]Note:[/b] If the sampler array is used use [param index] to access the specified texture.
*/
//go:nosplit
func (self class) ShaderSetDefaultTextureParameter(shader gd.RID, name gd.StringName, texture gd.RID, index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shader)
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, texture)
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_shader_set_default_texture_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a default texture from a shader searched by name.
[b]Note:[/b] If the sampler array is used use [param index] to access the specified texture.
*/
//go:nosplit
func (self class) ShaderGetDefaultTextureParameter(shader gd.RID, name gd.StringName, index gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shader)
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_shader_get_default_texture_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates an empty material and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]material_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
[b]Note:[/b] The equivalent resource is [Material].
*/
//go:nosplit
func (self class) MaterialCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_material_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a shader material's shader.
*/
//go:nosplit
func (self class) MaterialSetShader(shader_material gd.RID, shader gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shader_material)
	callframe.Arg(frame, shader)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_material_set_shader, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a material's parameter.
*/
//go:nosplit
func (self class) MaterialSetParam(material gd.RID, parameter gd.StringName, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, material)
	callframe.Arg(frame, mmm.Get(parameter))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_material_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of a certain material's parameter.
*/
//go:nosplit
func (self class) MaterialGetParam(ctx gd.Lifetime, material gd.RID, parameter gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, material)
	callframe.Arg(frame, mmm.Get(parameter))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_material_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets a material's render priority.
*/
//go:nosplit
func (self class) MaterialSetRenderPriority(material gd.RID, priority gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, material)
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_material_set_render_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets an object's next material.
*/
//go:nosplit
func (self class) MaterialSetNextPass(material gd.RID, next_material gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, material)
	callframe.Arg(frame, next_material)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_material_set_next_pass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) MeshCreateFromSurfaces(surfaces gd.ArrayOf[gd.Dictionary], blend_shape_count gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(surfaces))
	callframe.Arg(frame, blend_shape_count)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_create_from_surfaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new mesh and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]mesh_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
To place in a scene, attach this mesh to an instance using [method instance_set_base] using the returned RID.
[b]Note:[/b] The equivalent resource is [Mesh].
*/
//go:nosplit
func (self class) MeshCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the offset of a given attribute by [param array_index] in the start of its respective buffer.
*/
//go:nosplit
func (self class) MeshSurfaceGetFormatOffset(format classdb.RenderingServerArrayFormat, vertex_count gd.Int, array_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, format)
	callframe.Arg(frame, vertex_count)
	callframe.Arg(frame, array_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_surface_get_format_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the stride of the vertex positions for a mesh with given [param format]. Note importantly that vertex positions are stored consecutively and are not interleaved with the other attributes in the vertex buffer (normals and tangents).
*/
//go:nosplit
func (self class) MeshSurfaceGetFormatVertexStride(format classdb.RenderingServerArrayFormat, vertex_count gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, format)
	callframe.Arg(frame, vertex_count)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_surface_get_format_vertex_stride, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the stride of the combined normals and tangents for a mesh with given [param format]. Note importantly that, while normals and tangents are in the vertex buffer with vertices, they are only interleaved with each other and so have a different stride than vertex positions.
*/
//go:nosplit
func (self class) MeshSurfaceGetFormatNormalTangentStride(format classdb.RenderingServerArrayFormat, vertex_count gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, format)
	callframe.Arg(frame, vertex_count)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_surface_get_format_normal_tangent_stride, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the stride of the attribute buffer for a mesh with given [param format].
*/
//go:nosplit
func (self class) MeshSurfaceGetFormatAttributeStride(format classdb.RenderingServerArrayFormat, vertex_count gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, format)
	callframe.Arg(frame, vertex_count)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_surface_get_format_attribute_stride, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the stride of the skin buffer for a mesh with given [param format].
*/
//go:nosplit
func (self class) MeshSurfaceGetFormatSkinStride(format classdb.RenderingServerArrayFormat, vertex_count gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, format)
	callframe.Arg(frame, vertex_count)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_surface_get_format_skin_stride, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) MeshAddSurface(mesh gd.RID, surface gd.Dictionary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mesh)
	callframe.Arg(frame, mmm.Get(surface))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_add_surface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) MeshAddSurfaceFromArrays(mesh gd.RID, primitive classdb.RenderingServerPrimitiveType, arrays gd.Array, blend_shapes gd.Array, lods gd.Dictionary, compress_format classdb.RenderingServerArrayFormat)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mesh)
	callframe.Arg(frame, primitive)
	callframe.Arg(frame, mmm.Get(arrays))
	callframe.Arg(frame, mmm.Get(blend_shapes))
	callframe.Arg(frame, mmm.Get(lods))
	callframe.Arg(frame, compress_format)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_add_surface_from_arrays, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a mesh's blend shape count.
*/
//go:nosplit
func (self class) MeshGetBlendShapeCount(mesh gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mesh)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_get_blend_shape_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a mesh's blend shape mode.
*/
//go:nosplit
func (self class) MeshSetBlendShapeMode(mesh gd.RID, mode classdb.RenderingServerBlendShapeMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mesh)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_set_blend_shape_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a mesh's blend shape mode.
*/
//go:nosplit
func (self class) MeshGetBlendShapeMode(mesh gd.RID) classdb.RenderingServerBlendShapeMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mesh)
	var r_ret = callframe.Ret[classdb.RenderingServerBlendShapeMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_get_blend_shape_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a mesh's surface's material.
*/
//go:nosplit
func (self class) MeshSurfaceSetMaterial(mesh gd.RID, surface gd.Int, material gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mesh)
	callframe.Arg(frame, surface)
	callframe.Arg(frame, material)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_surface_set_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a mesh's surface's material.
*/
//go:nosplit
func (self class) MeshSurfaceGetMaterial(mesh gd.RID, surface gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mesh)
	callframe.Arg(frame, surface)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_surface_get_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) MeshGetSurface(ctx gd.Lifetime, mesh gd.RID, surface gd.Int) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mesh)
	callframe.Arg(frame, surface)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_get_surface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a mesh's surface's buffer arrays.
*/
//go:nosplit
func (self class) MeshSurfaceGetArrays(ctx gd.Lifetime, mesh gd.RID, surface gd.Int) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mesh)
	callframe.Arg(frame, surface)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_surface_get_arrays, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a mesh's surface's arrays for blend shapes.
*/
//go:nosplit
func (self class) MeshSurfaceGetBlendShapeArrays(ctx gd.Lifetime, mesh gd.RID, surface gd.Int) gd.ArrayOf[gd.Array] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mesh)
	callframe.Arg(frame, surface)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_surface_get_blend_shape_arrays, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Array](ret)
}
/*
Returns a mesh's number of surfaces.
*/
//go:nosplit
func (self class) MeshGetSurfaceCount(mesh gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mesh)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_get_surface_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a mesh's custom aabb.
*/
//go:nosplit
func (self class) MeshSetCustomAabb(mesh gd.RID, aabb gd.AABB)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mesh)
	callframe.Arg(frame, aabb)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_set_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a mesh's custom aabb.
*/
//go:nosplit
func (self class) MeshGetCustomAabb(mesh gd.RID) gd.AABB {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mesh)
	var r_ret = callframe.Ret[gd.AABB](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_get_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes all surfaces from a mesh.
*/
//go:nosplit
func (self class) MeshClear(mesh gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mesh)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) MeshSurfaceUpdateVertexRegion(mesh gd.RID, surface gd.Int, offset gd.Int, data gd.PackedByteArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mesh)
	callframe.Arg(frame, surface)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, mmm.Get(data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_surface_update_vertex_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) MeshSurfaceUpdateAttributeRegion(mesh gd.RID, surface gd.Int, offset gd.Int, data gd.PackedByteArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mesh)
	callframe.Arg(frame, surface)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, mmm.Get(data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_surface_update_attribute_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) MeshSurfaceUpdateSkinRegion(mesh gd.RID, surface gd.Int, offset gd.Int, data gd.PackedByteArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mesh)
	callframe.Arg(frame, surface)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, mmm.Get(data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_surface_update_skin_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) MeshSetShadowMesh(mesh gd.RID, shadow_mesh gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mesh)
	callframe.Arg(frame, shadow_mesh)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_mesh_set_shadow_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a new multimesh on the RenderingServer and returns an [RID] handle. This RID will be used in all [code]multimesh_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
To place in a scene, attach this multimesh to an instance using [method instance_set_base] using the returned RID.
[b]Note:[/b] The equivalent resource is [MultiMesh].
*/
//go:nosplit
func (self class) MultimeshCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) MultimeshAllocateData(multimesh gd.RID, instances gd.Int, transform_format classdb.RenderingServerMultimeshTransformFormat, color_format bool, custom_data_format bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multimesh)
	callframe.Arg(frame, instances)
	callframe.Arg(frame, transform_format)
	callframe.Arg(frame, color_format)
	callframe.Arg(frame, custom_data_format)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_allocate_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of instances allocated for this multimesh.
*/
//go:nosplit
func (self class) MultimeshGetInstanceCount(multimesh gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multimesh)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_get_instance_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the mesh to be drawn by the multimesh. Equivalent to [member MultiMesh.mesh].
*/
//go:nosplit
func (self class) MultimeshSetMesh(multimesh gd.RID, mesh gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multimesh)
	callframe.Arg(frame, mesh)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_set_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [Transform3D] for this instance. Equivalent to [method MultiMesh.set_instance_transform].
*/
//go:nosplit
func (self class) MultimeshInstanceSetTransform(multimesh gd.RID, index gd.Int, transform gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multimesh)
	callframe.Arg(frame, index)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_instance_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [Transform2D] for this instance. For use when multimesh is used in 2D. Equivalent to [method MultiMesh.set_instance_transform_2d].
*/
//go:nosplit
func (self class) MultimeshInstanceSetTransform2d(multimesh gd.RID, index gd.Int, transform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multimesh)
	callframe.Arg(frame, index)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_instance_set_transform_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the color by which this instance will be modulated. Equivalent to [method MultiMesh.set_instance_color].
*/
//go:nosplit
func (self class) MultimeshInstanceSetColor(multimesh gd.RID, index gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multimesh)
	callframe.Arg(frame, index)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_instance_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the custom data for this instance. Custom data is passed as a [Color], but is interpreted as a [code]vec4[/code] in the shader. Equivalent to [method MultiMesh.set_instance_custom_data].
*/
//go:nosplit
func (self class) MultimeshInstanceSetCustomData(multimesh gd.RID, index gd.Int, custom_data gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multimesh)
	callframe.Arg(frame, index)
	callframe.Arg(frame, custom_data)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_instance_set_custom_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the RID of the mesh that will be used in drawing this multimesh.
*/
//go:nosplit
func (self class) MultimeshGetMesh(multimesh gd.RID) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multimesh)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_get_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Calculates and returns the axis-aligned bounding box that encloses all instances within the multimesh.
*/
//go:nosplit
func (self class) MultimeshGetAabb(multimesh gd.RID) gd.AABB {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multimesh)
	var r_ret = callframe.Ret[gd.AABB](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_get_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the custom AABB for this MultiMesh resource.
*/
//go:nosplit
func (self class) MultimeshSetCustomAabb(multimesh gd.RID, aabb gd.AABB)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multimesh)
	callframe.Arg(frame, aabb)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_set_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the custom AABB defined for this MultiMesh resource.
*/
//go:nosplit
func (self class) MultimeshGetCustomAabb(multimesh gd.RID) gd.AABB {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multimesh)
	var r_ret = callframe.Ret[gd.AABB](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_get_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Transform3D] of the specified instance.
*/
//go:nosplit
func (self class) MultimeshInstanceGetTransform(multimesh gd.RID, index gd.Int) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multimesh)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_instance_get_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Transform2D] of the specified instance. For use when the multimesh is set to use 2D transforms.
*/
//go:nosplit
func (self class) MultimeshInstanceGetTransform2d(multimesh gd.RID, index gd.Int) gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multimesh)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_instance_get_transform_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the color by which the specified instance will be modulated.
*/
//go:nosplit
func (self class) MultimeshInstanceGetColor(multimesh gd.RID, index gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multimesh)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_instance_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the custom data associated with the specified instance.
*/
//go:nosplit
func (self class) MultimeshInstanceGetCustomData(multimesh gd.RID, index gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multimesh)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_instance_get_custom_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the number of instances visible at a given time. If -1, all instances that have been allocated are drawn. Equivalent to [member MultiMesh.visible_instance_count].
*/
//go:nosplit
func (self class) MultimeshSetVisibleInstances(multimesh gd.RID, visible gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multimesh)
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_set_visible_instances, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of visible instances for this multimesh.
*/
//go:nosplit
func (self class) MultimeshGetVisibleInstances(multimesh gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multimesh)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_get_visible_instances, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set the entire data to use for drawing the [param multimesh] at once to [param buffer] (such as instance transforms and colors). [param buffer]'s size must match the number of instances multiplied by the per-instance data size (which depends on the enabled MultiMesh fields). Otherwise, an error message is printed and nothing is rendered. See also [method multimesh_get_buffer].
The per-instance data size and expected data order is:
[codeblock lang=text]
2D:
  - Position: 8 floats (8 floats for Transform2D)
  - Position + Vertex color: 12 floats (8 floats for Transform2D, 4 floats for Color)
  - Position + Custom data: 12 floats (8 floats for Transform2D, 4 floats of custom data)
  - Position + Vertex color + Custom data: 16 floats (8 floats for Transform2D, 4 floats for Color, 4 floats of custom data)
3D:
  - Position: 12 floats (12 floats for Transform3D)
  - Position + Vertex color: 16 floats (12 floats for Transform3D, 4 floats for Color)
  - Position + Custom data: 16 floats (12 floats for Transform3D, 4 floats of custom data)
  - Position + Vertex color + Custom data: 20 floats (12 floats for Transform3D, 4 floats for Color, 4 floats of custom data)
[/codeblock]
*/
//go:nosplit
func (self class) MultimeshSetBuffer(multimesh gd.RID, buffer gd.PackedFloat32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multimesh)
	callframe.Arg(frame, mmm.Get(buffer))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_set_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the MultiMesh data (such as instance transforms, colors, etc.). See [method multimesh_set_buffer] for details on the returned data.
[b]Note:[/b] If the buffer is in the engine's internal cache, it will have to be fetched from GPU memory and possibly decompressed. This means [method multimesh_get_buffer] is potentially a slow operation and should be avoided whenever possible.
*/
//go:nosplit
func (self class) MultimeshGetBuffer(ctx gd.Lifetime, multimesh gd.RID) gd.PackedFloat32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multimesh)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_multimesh_get_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedFloat32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Creates a skeleton and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]skeleton_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
*/
//go:nosplit
func (self class) SkeletonCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_skeleton_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SkeletonAllocateData(skeleton gd.RID, bones gd.Int, is_2d_skeleton bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, skeleton)
	callframe.Arg(frame, bones)
	callframe.Arg(frame, is_2d_skeleton)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_skeleton_allocate_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of bones allocated for this skeleton.
*/
//go:nosplit
func (self class) SkeletonGetBoneCount(skeleton gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, skeleton)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_skeleton_get_bone_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [Transform3D] for a specific bone of this skeleton.
*/
//go:nosplit
func (self class) SkeletonBoneSetTransform(skeleton gd.RID, bone gd.Int, transform gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, skeleton)
	callframe.Arg(frame, bone)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_skeleton_bone_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Transform3D] set for a specific bone of this skeleton.
*/
//go:nosplit
func (self class) SkeletonBoneGetTransform(skeleton gd.RID, bone gd.Int) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, skeleton)
	callframe.Arg(frame, bone)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_skeleton_bone_get_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [Transform2D] for a specific bone of this skeleton.
*/
//go:nosplit
func (self class) SkeletonBoneSetTransform2d(skeleton gd.RID, bone gd.Int, transform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, skeleton)
	callframe.Arg(frame, bone)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_skeleton_bone_set_transform_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Transform2D] set for a specific bone of this skeleton.
*/
//go:nosplit
func (self class) SkeletonBoneGetTransform2d(skeleton gd.RID, bone gd.Int) gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, skeleton)
	callframe.Arg(frame, bone)
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_skeleton_bone_get_transform_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SkeletonSetBaseTransform2d(skeleton gd.RID, base_transform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, skeleton)
	callframe.Arg(frame, base_transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_skeleton_set_base_transform_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a directional light and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID can be used in most [code]light_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
To place in a scene, attach this directional light to an instance using [method instance_set_base] using the returned RID.
[b]Note:[/b] The equivalent node is [DirectionalLight3D].
*/
//go:nosplit
func (self class) DirectionalLightCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_directional_light_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new omni light and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID can be used in most [code]light_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
To place in a scene, attach this omni light to an instance using [method instance_set_base] using the returned RID.
[b]Note:[/b] The equivalent node is [OmniLight3D].
*/
//go:nosplit
func (self class) OmniLightCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_omni_light_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a spot light and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID can be used in most [code]light_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
To place in a scene, attach this spot light to an instance using [method instance_set_base] using the returned RID.
*/
//go:nosplit
func (self class) SpotLightCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_spot_light_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the color of the light. Equivalent to [member Light3D.light_color].
*/
//go:nosplit
func (self class) LightSetColor(light gd.RID, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_light_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the specified 3D light parameter. See [enum LightParam] for options. Equivalent to [method Light3D.set_param].
*/
//go:nosplit
func (self class) LightSetParam(light gd.RID, param classdb.RenderingServerLightParam, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_light_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], light will cast shadows. Equivalent to [member Light3D.shadow_enabled].
*/
//go:nosplit
func (self class) LightSetShadow(light gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_light_set_shadow, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the projector texture to use for the specified 3D light. Equivalent to [member Light3D.light_projector].
*/
//go:nosplit
func (self class) LightSetProjector(light gd.RID, texture gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, texture)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_light_set_projector, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], the 3D light will subtract light instead of adding light. Equivalent to [member Light3D.light_negative].
*/
//go:nosplit
func (self class) LightSetNegative(light gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_light_set_negative, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the cull mask for this 3D light. Lights only affect objects in the selected layers. Equivalent to [member Light3D.light_cull_mask].
*/
//go:nosplit
func (self class) LightSetCullMask(light gd.RID, mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_light_set_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the distance fade for this 3D light. This acts as a form of level of detail (LOD) and can be used to improve performance. Equivalent to [member Light3D.distance_fade_enabled], [member Light3D.distance_fade_begin], [member Light3D.distance_fade_shadow], and [member Light3D.distance_fade_length].
*/
//go:nosplit
func (self class) LightSetDistanceFade(decal gd.RID, enabled bool, begin gd.Float, shadow gd.Float, length gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, decal)
	callframe.Arg(frame, enabled)
	callframe.Arg(frame, begin)
	callframe.Arg(frame, shadow)
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_light_set_distance_fade, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], reverses the backface culling of the mesh. This can be useful when you have a flat mesh that has a light behind it. If you need to cast a shadow on both sides of the mesh, set the mesh to use double-sided shadows with [method instance_geometry_set_cast_shadows_setting]. Equivalent to [member Light3D.shadow_reverse_cull_face].
*/
//go:nosplit
func (self class) LightSetReverseCullFaceMode(light gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_light_set_reverse_cull_face_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the bake mode to use for the specified 3D light. Equivalent to [member Light3D.light_bake_mode].
*/
//go:nosplit
func (self class) LightSetBakeMode(light gd.RID, bake_mode classdb.RenderingServerLightBakeMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, bake_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_light_set_bake_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the maximum SDFGI cascade in which the 3D light's indirect lighting is rendered. Higher values allow the light to be rendered in SDFGI further away from the camera.
*/
//go:nosplit
func (self class) LightSetMaxSdfgiCascade(light gd.RID, cascade gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, cascade)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_light_set_max_sdfgi_cascade, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets whether to use a dual paraboloid or a cubemap for the shadow map. Dual paraboloid is faster but may suffer from artifacts. Equivalent to [member OmniLight3D.omni_shadow_mode].
*/
//go:nosplit
func (self class) LightOmniSetShadowMode(light gd.RID, mode classdb.RenderingServerLightOmniShadowMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_light_omni_set_shadow_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the shadow mode for this directional light. Equivalent to [member DirectionalLight3D.directional_shadow_mode]. See [enum LightDirectionalShadowMode] for options.
*/
//go:nosplit
func (self class) LightDirectionalSetShadowMode(light gd.RID, mode classdb.RenderingServerLightDirectionalShadowMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_light_directional_set_shadow_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], this directional light will blend between shadow map splits resulting in a smoother transition between them. Equivalent to [member DirectionalLight3D.directional_shadow_blend_splits].
*/
//go:nosplit
func (self class) LightDirectionalSetBlendSplits(light gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_light_directional_set_blend_splits, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], this light will not be used for anything except sky shaders. Use this for lights that impact your sky shader that you may want to hide from affecting the rest of the scene. For example, you may want to enable this when the sun in your sky shader falls below the horizon.
*/
//go:nosplit
func (self class) LightDirectionalSetSkyMode(light gd.RID, mode classdb.RenderingServerLightDirectionalSkyMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_light_directional_set_sky_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the texture filter mode to use when rendering light projectors. This parameter is global and cannot be set on a per-light basis.
*/
//go:nosplit
func (self class) LightProjectorsSetFilter(filter classdb.RenderingServerLightProjectorFilter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, filter)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_light_projectors_set_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the filter quality for omni and spot light shadows in 3D. See also [member ProjectSettings.rendering/lights_and_shadows/positional_shadow/soft_shadow_filter_quality]. This parameter is global and cannot be set on a per-viewport basis.
*/
//go:nosplit
func (self class) PositionalSoftShadowFilterSetQuality(quality classdb.RenderingServerShadowQuality)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, quality)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_positional_soft_shadow_filter_set_quality, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the filter [param quality] for directional light shadows in 3D. See also [member ProjectSettings.rendering/lights_and_shadows/directional_shadow/soft_shadow_filter_quality]. This parameter is global and cannot be set on a per-viewport basis.
*/
//go:nosplit
func (self class) DirectionalSoftShadowFilterSetQuality(quality classdb.RenderingServerShadowQuality)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, quality)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_directional_soft_shadow_filter_set_quality, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [param size] of the directional light shadows in 3D. See also [member ProjectSettings.rendering/lights_and_shadows/directional_shadow/size]. This parameter is global and cannot be set on a per-viewport basis.
*/
//go:nosplit
func (self class) DirectionalShadowAtlasSetSize(size gd.Int, is_16bits bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, is_16bits)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_directional_shadow_atlas_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a reflection probe and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]reflection_probe_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
To place in a scene, attach this reflection probe to an instance using [method instance_set_base] using the returned RID.
[b]Note:[/b] The equivalent node is [ReflectionProbe].
*/
//go:nosplit
func (self class) ReflectionProbeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_reflection_probe_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets how often the reflection probe updates. Can either be once or every frame. See [enum ReflectionProbeUpdateMode] for options.
*/
//go:nosplit
func (self class) ReflectionProbeSetUpdateMode(probe gd.RID, mode classdb.RenderingServerReflectionProbeUpdateMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, probe)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_reflection_probe_set_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the intensity of the reflection probe. Intensity modulates the strength of the reflection. Equivalent to [member ReflectionProbe.intensity].
*/
//go:nosplit
func (self class) ReflectionProbeSetIntensity(probe gd.RID, intensity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, probe)
	callframe.Arg(frame, intensity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_reflection_probe_set_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the reflection probe's ambient light mode. Equivalent to [member ReflectionProbe.ambient_mode].
*/
//go:nosplit
func (self class) ReflectionProbeSetAmbientMode(probe gd.RID, mode classdb.RenderingServerReflectionProbeAmbientMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, probe)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_reflection_probe_set_ambient_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the reflection probe's custom ambient light color. Equivalent to [member ReflectionProbe.ambient_color].
*/
//go:nosplit
func (self class) ReflectionProbeSetAmbientColor(probe gd.RID, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, probe)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_reflection_probe_set_ambient_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the reflection probe's custom ambient light energy. Equivalent to [member ReflectionProbe.ambient_color_energy].
*/
//go:nosplit
func (self class) ReflectionProbeSetAmbientEnergy(probe gd.RID, energy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, probe)
	callframe.Arg(frame, energy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_reflection_probe_set_ambient_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the max distance away from the probe an object can be before it is culled. Equivalent to [member ReflectionProbe.max_distance].
*/
//go:nosplit
func (self class) ReflectionProbeSetMaxDistance(probe gd.RID, distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, probe)
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_reflection_probe_set_max_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the size of the area that the reflection probe will capture. Equivalent to [member ReflectionProbe.size].
*/
//go:nosplit
func (self class) ReflectionProbeSetSize(probe gd.RID, size gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, probe)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_reflection_probe_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the origin offset to be used when this reflection probe is in box project mode. Equivalent to [member ReflectionProbe.origin_offset].
*/
//go:nosplit
func (self class) ReflectionProbeSetOriginOffset(probe gd.RID, offset gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, probe)
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_reflection_probe_set_origin_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], reflections will ignore sky contribution. Equivalent to [member ReflectionProbe.interior].
*/
//go:nosplit
func (self class) ReflectionProbeSetAsInterior(probe gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, probe)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_reflection_probe_set_as_interior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], uses box projection. This can make reflections look more correct in certain situations. Equivalent to [member ReflectionProbe.box_projection].
*/
//go:nosplit
func (self class) ReflectionProbeSetEnableBoxProjection(probe gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, probe)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_reflection_probe_set_enable_box_projection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], computes shadows in the reflection probe. This makes the reflection much slower to compute. Equivalent to [member ReflectionProbe.enable_shadows].
*/
//go:nosplit
func (self class) ReflectionProbeSetEnableShadows(probe gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, probe)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_reflection_probe_set_enable_shadows, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the render cull mask for this reflection probe. Only instances with a matching layer will be reflected by this probe. Equivalent to [member ReflectionProbe.cull_mask].
*/
//go:nosplit
func (self class) ReflectionProbeSetCullMask(probe gd.RID, layers gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, probe)
	callframe.Arg(frame, layers)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_reflection_probe_set_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the render reflection mask for this reflection probe. Only instances with a matching layer will have reflections applied from this probe. Equivalent to [member ReflectionProbe.reflection_mask].
*/
//go:nosplit
func (self class) ReflectionProbeSetReflectionMask(probe gd.RID, layers gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, probe)
	callframe.Arg(frame, layers)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_reflection_probe_set_reflection_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the resolution to use when rendering the specified reflection probe. The [param resolution] is specified for each cubemap face: for instance, specifying [code]512[/code] will allocate 6 faces of 512×512 each (plus mipmaps for roughness levels).
*/
//go:nosplit
func (self class) ReflectionProbeSetResolution(probe gd.RID, resolution gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, probe)
	callframe.Arg(frame, resolution)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_reflection_probe_set_resolution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the mesh level of detail to use in the reflection probe rendering. Higher values will use less detailed versions of meshes that have LOD variations generated, which can improve performance. Equivalent to [member ReflectionProbe.mesh_lod_threshold].
*/
//go:nosplit
func (self class) ReflectionProbeSetMeshLodThreshold(probe gd.RID, pixels gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, probe)
	callframe.Arg(frame, pixels)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_reflection_probe_set_mesh_lod_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a decal and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]decal_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
To place in a scene, attach this decal to an instance using [method instance_set_base] using the returned RID.
[b]Note:[/b] The equivalent node is [Decal].
*/
//go:nosplit
func (self class) DecalCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_decal_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [param size] of the decal specified by the [param decal] RID. Equivalent to [member Decal.size].
*/
//go:nosplit
func (self class) DecalSetSize(decal gd.RID, size gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, decal)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_decal_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [param texture] in the given texture [param type] slot for the specified decal. Equivalent to [method Decal.set_texture].
*/
//go:nosplit
func (self class) DecalSetTexture(decal gd.RID, atype classdb.RenderingServerDecalTexture, texture gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, decal)
	callframe.Arg(frame, atype)
	callframe.Arg(frame, texture)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_decal_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the emission [param energy] in the decal specified by the [param decal] RID. Equivalent to [member Decal.emission_energy].
*/
//go:nosplit
func (self class) DecalSetEmissionEnergy(decal gd.RID, energy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, decal)
	callframe.Arg(frame, energy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_decal_set_emission_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [param albedo_mix] in the decal specified by the [param decal] RID. Equivalent to [member Decal.albedo_mix].
*/
//go:nosplit
func (self class) DecalSetAlbedoMix(decal gd.RID, albedo_mix gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, decal)
	callframe.Arg(frame, albedo_mix)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_decal_set_albedo_mix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the color multiplier in the decal specified by the [param decal] RID to [param color]. Equivalent to [member Decal.modulate].
*/
//go:nosplit
func (self class) DecalSetModulate(decal gd.RID, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, decal)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_decal_set_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the cull [param mask] in the decal specified by the [param decal] RID. Equivalent to [member Decal.cull_mask].
*/
//go:nosplit
func (self class) DecalSetCullMask(decal gd.RID, mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, decal)
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_decal_set_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the distance fade parameters in the decal specified by the [param decal] RID. Equivalent to [member Decal.distance_fade_enabled], [member Decal.distance_fade_begin] and [member Decal.distance_fade_length].
*/
//go:nosplit
func (self class) DecalSetDistanceFade(decal gd.RID, enabled bool, begin gd.Float, length gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, decal)
	callframe.Arg(frame, enabled)
	callframe.Arg(frame, begin)
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_decal_set_distance_fade, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the upper fade ([param above]) and lower fade ([param below]) in the decal specified by the [param decal] RID. Equivalent to [member Decal.upper_fade] and [member Decal.lower_fade].
*/
//go:nosplit
func (self class) DecalSetFade(decal gd.RID, above gd.Float, below gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, decal)
	callframe.Arg(frame, above)
	callframe.Arg(frame, below)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_decal_set_fade, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the normal [param fade] in the decal specified by the [param decal] RID. Equivalent to [member Decal.normal_fade].
*/
//go:nosplit
func (self class) DecalSetNormalFade(decal gd.RID, fade gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, decal)
	callframe.Arg(frame, fade)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_decal_set_normal_fade, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the texture [param filter] mode to use when rendering decals. This parameter is global and cannot be set on a per-decal basis.
*/
//go:nosplit
func (self class) DecalsSetFilter(filter classdb.RenderingServerDecalFilter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, filter)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_decals_set_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [param half_resolution] is [code]true[/code], renders [VoxelGI] and SDFGI ([member Environment.sdfgi_enabled]) buffers at halved resolution on each axis (e.g. 960×540 when the viewport size is 1920×1080). This improves performance significantly when VoxelGI or SDFGI is enabled, at the cost of artifacts that may be visible on polygon edges. The loss in quality becomes less noticeable as the viewport resolution increases. [LightmapGI] rendering is not affected by this setting. Equivalent to [member ProjectSettings.rendering/global_illumination/gi/use_half_resolution].
*/
//go:nosplit
func (self class) GiSetUseHalfResolution(half_resolution bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, half_resolution)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_gi_set_use_half_resolution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a new voxel-based global illumination object and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]voxel_gi_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
[b]Note:[/b] The equivalent node is [VoxelGI].
*/
//go:nosplit
func (self class) VoxelGiCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_voxel_gi_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) VoxelGiAllocateData(voxel_gi gd.RID, to_cell_xform gd.Transform3D, aabb gd.AABB, octree_size gd.Vector3i, octree_cells gd.PackedByteArray, data_cells gd.PackedByteArray, distance_field gd.PackedByteArray, level_counts gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voxel_gi)
	callframe.Arg(frame, to_cell_xform)
	callframe.Arg(frame, aabb)
	callframe.Arg(frame, octree_size)
	callframe.Arg(frame, mmm.Get(octree_cells))
	callframe.Arg(frame, mmm.Get(data_cells))
	callframe.Arg(frame, mmm.Get(distance_field))
	callframe.Arg(frame, mmm.Get(level_counts))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_voxel_gi_allocate_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) VoxelGiGetOctreeSize(voxel_gi gd.RID) gd.Vector3i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voxel_gi)
	var r_ret = callframe.Ret[gd.Vector3i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_voxel_gi_get_octree_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) VoxelGiGetOctreeCells(ctx gd.Lifetime, voxel_gi gd.RID) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voxel_gi)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_voxel_gi_get_octree_cells, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) VoxelGiGetDataCells(ctx gd.Lifetime, voxel_gi gd.RID) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voxel_gi)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_voxel_gi_get_data_cells, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) VoxelGiGetDistanceField(ctx gd.Lifetime, voxel_gi gd.RID) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voxel_gi)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_voxel_gi_get_distance_field, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) VoxelGiGetLevelCounts(ctx gd.Lifetime, voxel_gi gd.RID) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voxel_gi)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_voxel_gi_get_level_counts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) VoxelGiGetToCellXform(voxel_gi gd.RID) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voxel_gi)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_voxel_gi_get_to_cell_xform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [member VoxelGIData.dynamic_range] value to use on the specified [param voxel_gi]'s [RID].
*/
//go:nosplit
func (self class) VoxelGiSetDynamicRange(voxel_gi gd.RID, arange gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voxel_gi)
	callframe.Arg(frame, arange)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_voxel_gi_set_dynamic_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [member VoxelGIData.propagation] value to use on the specified [param voxel_gi]'s [RID].
*/
//go:nosplit
func (self class) VoxelGiSetPropagation(voxel_gi gd.RID, amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voxel_gi)
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_voxel_gi_set_propagation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [member VoxelGIData.energy] value to use on the specified [param voxel_gi]'s [RID].
*/
//go:nosplit
func (self class) VoxelGiSetEnergy(voxel_gi gd.RID, energy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voxel_gi)
	callframe.Arg(frame, energy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_voxel_gi_set_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Used to inform the renderer what exposure normalization value was used while baking the voxel gi. This value will be used and modulated at run time to ensure that the voxel gi maintains a consistent level of exposure even if the scene-wide exposure normalization is changed at run time. For more information see [method camera_attributes_set_exposure].
*/
//go:nosplit
func (self class) VoxelGiSetBakedExposureNormalization(voxel_gi gd.RID, baked_exposure gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voxel_gi)
	callframe.Arg(frame, baked_exposure)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_voxel_gi_set_baked_exposure_normalization, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [member VoxelGIData.bias] value to use on the specified [param voxel_gi]'s [RID].
*/
//go:nosplit
func (self class) VoxelGiSetBias(voxel_gi gd.RID, bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voxel_gi)
	callframe.Arg(frame, bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_voxel_gi_set_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [member VoxelGIData.normal_bias] value to use on the specified [param voxel_gi]'s [RID].
*/
//go:nosplit
func (self class) VoxelGiSetNormalBias(voxel_gi gd.RID, bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voxel_gi)
	callframe.Arg(frame, bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_voxel_gi_set_normal_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [member VoxelGIData.interior] value to use on the specified [param voxel_gi]'s [RID].
*/
//go:nosplit
func (self class) VoxelGiSetInterior(voxel_gi gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voxel_gi)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_voxel_gi_set_interior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [member VoxelGIData.use_two_bounces] value to use on the specified [param voxel_gi]'s [RID].
*/
//go:nosplit
func (self class) VoxelGiSetUseTwoBounces(voxel_gi gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voxel_gi)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_voxel_gi_set_use_two_bounces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [member ProjectSettings.rendering/global_illumination/voxel_gi/quality] value to use when rendering. This parameter is global and cannot be set on a per-VoxelGI basis.
*/
//go:nosplit
func (self class) VoxelGiSetQuality(quality classdb.RenderingServerVoxelGIQuality)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, quality)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_voxel_gi_set_quality, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a new lightmap global illumination instance and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]lightmap_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
[b]Note:[/b] The equivalent node is [LightmapGI].
*/
//go:nosplit
func (self class) LightmapCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_lightmap_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set the textures on the given [param lightmap] GI instance to the texture array pointed to by the [param light] RID. If the lightmap texture was baked with [member LightmapGI.directional] set to [code]true[/code], then [param uses_sh] must also be [code]true[/code].
*/
//go:nosplit
func (self class) LightmapSetTextures(lightmap gd.RID, light gd.RID, uses_sh bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, lightmap)
	callframe.Arg(frame, light)
	callframe.Arg(frame, uses_sh)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_lightmap_set_textures, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) LightmapSetProbeBounds(lightmap gd.RID, bounds gd.AABB)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, lightmap)
	callframe.Arg(frame, bounds)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_lightmap_set_probe_bounds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) LightmapSetProbeInterior(lightmap gd.RID, interior bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, lightmap)
	callframe.Arg(frame, interior)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_lightmap_set_probe_interior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) LightmapSetProbeCaptureData(lightmap gd.RID, points gd.PackedVector3Array, point_sh gd.PackedColorArray, tetrahedra gd.PackedInt32Array, bsp_tree gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, lightmap)
	callframe.Arg(frame, mmm.Get(points))
	callframe.Arg(frame, mmm.Get(point_sh))
	callframe.Arg(frame, mmm.Get(tetrahedra))
	callframe.Arg(frame, mmm.Get(bsp_tree))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_lightmap_set_probe_capture_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) LightmapGetProbeCapturePoints(ctx gd.Lifetime, lightmap gd.RID) gd.PackedVector3Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, lightmap)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_lightmap_get_probe_capture_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector3Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) LightmapGetProbeCaptureSh(ctx gd.Lifetime, lightmap gd.RID) gd.PackedColorArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, lightmap)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_lightmap_get_probe_capture_sh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedColorArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) LightmapGetProbeCaptureTetrahedra(ctx gd.Lifetime, lightmap gd.RID) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, lightmap)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_lightmap_get_probe_capture_tetrahedra, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) LightmapGetProbeCaptureBspTree(ctx gd.Lifetime, lightmap gd.RID) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, lightmap)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_lightmap_get_probe_capture_bsp_tree, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Used to inform the renderer what exposure normalization value was used while baking the lightmap. This value will be used and modulated at run time to ensure that the lightmap maintains a consistent level of exposure even if the scene-wide exposure normalization is changed at run time. For more information see [method camera_attributes_set_exposure].
*/
//go:nosplit
func (self class) LightmapSetBakedExposureNormalization(lightmap gd.RID, baked_exposure gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, lightmap)
	callframe.Arg(frame, baked_exposure)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_lightmap_set_baked_exposure_normalization, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) LightmapSetProbeCaptureUpdateSpeed(speed gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, speed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_lightmap_set_probe_capture_update_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a GPU-based particle system and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]particles_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
To place in a scene, attach these particles to an instance using [method instance_set_base] using the returned RID.
[b]Note:[/b] The equivalent nodes are [GPUParticles2D] and [GPUParticles3D].
[b]Note:[/b] All [code]particles_*[/code] methods only apply to GPU-based particles, not CPU-based particles. [CPUParticles2D] and [CPUParticles3D] do not have equivalent RenderingServer functions available, as these use [MultiMeshInstance2D] and [MultiMeshInstance3D] under the hood (see [code]multimesh_*[/code] methods).
*/
//go:nosplit
func (self class) ParticlesCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the GPU particles specified by the [param particles] RID should be rendered in 2D or 3D according to [param mode].
*/
//go:nosplit
func (self class) ParticlesSetMode(particles gd.RID, mode classdb.RenderingServerParticlesMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], particles will emit over time. Setting to false does not reset the particles, but only stops their emission. Equivalent to [member GPUParticles3D.emitting].
*/
//go:nosplit
func (self class) ParticlesSetEmitting(particles gd.RID, emitting bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, emitting)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_emitting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if particles are currently set to emitting.
*/
//go:nosplit
func (self class) ParticlesGetEmitting(particles gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_get_emitting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the number of particles to be drawn and allocates the memory for them. Equivalent to [member GPUParticles3D.amount].
*/
//go:nosplit
func (self class) ParticlesSetAmount(particles gd.RID, amount gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the amount ratio for particles to be emitted. Equivalent to [member GPUParticles3D.amount_ratio].
*/
//go:nosplit
func (self class) ParticlesSetAmountRatio(particles gd.RID, ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_amount_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the lifetime of each particle in the system. Equivalent to [member GPUParticles3D.lifetime].
*/
//go:nosplit
func (self class) ParticlesSetLifetime(particles gd.RID, lifetime gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, lifetime)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_lifetime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], particles will emit once and then stop. Equivalent to [member GPUParticles3D.one_shot].
*/
//go:nosplit
func (self class) ParticlesSetOneShot(particles gd.RID, one_shot bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, one_shot)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_one_shot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the preprocess time for the particles' animation. This lets you delay starting an animation until after the particles have begun emitting. Equivalent to [member GPUParticles3D.preprocess].
*/
//go:nosplit
func (self class) ParticlesSetPreProcessTime(particles gd.RID, time gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, time)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_pre_process_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the explosiveness ratio. Equivalent to [member GPUParticles3D.explosiveness].
*/
//go:nosplit
func (self class) ParticlesSetExplosivenessRatio(particles gd.RID, ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_explosiveness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the emission randomness ratio. This randomizes the emission of particles within their phase. Equivalent to [member GPUParticles3D.randomness].
*/
//go:nosplit
func (self class) ParticlesSetRandomnessRatio(particles gd.RID, ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_randomness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the value that informs a [ParticleProcessMaterial] to rush all particles towards the end of their lifetime.
*/
//go:nosplit
func (self class) ParticlesSetInterpToEnd(particles gd.RID, factor gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, factor)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_interp_to_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the velocity of a particle node, that will be used by [member ParticleProcessMaterial.inherit_velocity_ratio].
*/
//go:nosplit
func (self class) ParticlesSetEmitterVelocity(particles gd.RID, velocity gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_emitter_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a custom axis-aligned bounding box for the particle system. Equivalent to [member GPUParticles3D.visibility_aabb].
*/
//go:nosplit
func (self class) ParticlesSetCustomAabb(particles gd.RID, aabb gd.AABB)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, aabb)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the speed scale of the particle system. Equivalent to [member GPUParticles3D.speed_scale].
*/
//go:nosplit
func (self class) ParticlesSetSpeedScale(particles gd.RID, scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], particles use local coordinates. If [code]false[/code] they use global coordinates. Equivalent to [member GPUParticles3D.local_coords].
*/
//go:nosplit
func (self class) ParticlesSetUseLocalCoordinates(particles gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_use_local_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the material for processing the particles.
[b]Note:[/b] This is not the material used to draw the materials. Equivalent to [member GPUParticles3D.process_material].
*/
//go:nosplit
func (self class) ParticlesSetProcessMaterial(particles gd.RID, material gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, material)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_process_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the frame rate that the particle system rendering will be fixed to. Equivalent to [member GPUParticles3D.fixed_fps].
*/
//go:nosplit
func (self class) ParticlesSetFixedFps(particles gd.RID, fps gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, fps)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_fixed_fps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) ParticlesSetInterpolate(particles gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_interpolate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], uses fractional delta which smooths the movement of the particles. Equivalent to [member GPUParticles3D.fract_delta].
*/
//go:nosplit
func (self class) ParticlesSetFractionalDelta(particles gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_fractional_delta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) ParticlesSetCollisionBaseSize(particles gd.RID, size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_collision_base_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) ParticlesSetTransformAlign(particles gd.RID, align classdb.RenderingServerParticlesTransformAlign)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, align)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_transform_align, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [param enable] is [code]true[/code], enables trails for the [param particles] with the specified [param length_sec] in seconds. Equivalent to [member GPUParticles3D.trail_enabled] and [member GPUParticles3D.trail_lifetime].
*/
//go:nosplit
func (self class) ParticlesSetTrails(particles gd.RID, enable bool, length_sec gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, enable)
	callframe.Arg(frame, length_sec)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_trails, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) ParticlesSetTrailBindPoses(particles gd.RID, bind_poses gd.ArrayOf[gd.Transform3D])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, mmm.Get(bind_poses))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_trail_bind_poses, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if particles are not emitting and particles are set to inactive.
*/
//go:nosplit
func (self class) ParticlesIsInactive(particles gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_is_inactive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Add particle system to list of particle systems that need to be updated. Update will take place on the next frame, or on the next call to [method instances_cull_aabb], [method instances_cull_convex], or [method instances_cull_ray].
*/
//go:nosplit
func (self class) ParticlesRequestProcess(particles gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_request_process, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Reset the particles on the next update. Equivalent to [method GPUParticles3D.restart].
*/
//go:nosplit
func (self class) ParticlesRestart(particles gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_restart, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) ParticlesSetSubemitter(particles gd.RID, subemitter_particles gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, subemitter_particles)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_subemitter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Manually emits particles from the [param particles] instance.
*/
//go:nosplit
func (self class) ParticlesEmit(particles gd.RID, transform gd.Transform3D, velocity gd.Vector3, color gd.Color, custom gd.Color, emit_flags gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, transform)
	callframe.Arg(frame, velocity)
	callframe.Arg(frame, color)
	callframe.Arg(frame, custom)
	callframe.Arg(frame, emit_flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_emit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the draw order of the particles to one of the named enums from [enum ParticlesDrawOrder]. See [enum ParticlesDrawOrder] for options. Equivalent to [member GPUParticles3D.draw_order].
*/
//go:nosplit
func (self class) ParticlesSetDrawOrder(particles gd.RID, order classdb.RenderingServerParticlesDrawOrder)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, order)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_draw_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the number of draw passes to use. Equivalent to [member GPUParticles3D.draw_passes].
*/
//go:nosplit
func (self class) ParticlesSetDrawPasses(particles gd.RID, count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_draw_passes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the mesh to be used for the specified draw pass. Equivalent to [member GPUParticles3D.draw_pass_1], [member GPUParticles3D.draw_pass_2], [member GPUParticles3D.draw_pass_3], and [member GPUParticles3D.draw_pass_4].
*/
//go:nosplit
func (self class) ParticlesSetDrawPassMesh(particles gd.RID, pass gd.Int, mesh gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, pass)
	callframe.Arg(frame, mesh)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_draw_pass_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Calculates and returns the axis-aligned bounding box that contains all the particles. Equivalent to [method GPUParticles3D.capture_aabb].
*/
//go:nosplit
func (self class) ParticlesGetCurrentAabb(particles gd.RID) gd.AABB {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	var r_ret = callframe.Ret[gd.AABB](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_get_current_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [Transform3D] that will be used by the particles when they first emit.
*/
//go:nosplit
func (self class) ParticlesSetEmissionTransform(particles gd.RID, transform gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_set_emission_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a new 3D GPU particle collision or attractor and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID can be used in most [code]particles_collision_*[/code] RenderingServer functions.
[b]Note:[/b] The equivalent nodes are [GPUParticlesCollision3D] and [GPUParticlesAttractor3D].
*/
//go:nosplit
func (self class) ParticlesCollisionCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_collision_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the collision or attractor shape [param type] for the 3D GPU particles collision or attractor specified by the [param particles_collision] RID.
*/
//go:nosplit
func (self class) ParticlesCollisionSetCollisionType(particles_collision gd.RID, atype classdb.RenderingServerParticlesCollisionType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles_collision)
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_collision_set_collision_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the cull [param mask] for the 3D GPU particles collision or attractor specified by the [param particles_collision] RID. Equivalent to [member GPUParticlesCollision3D.cull_mask] or [member GPUParticlesAttractor3D.cull_mask] depending on the [param particles_collision] type.
*/
//go:nosplit
func (self class) ParticlesCollisionSetCullMask(particles_collision gd.RID, mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles_collision)
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_collision_set_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [param radius] for the 3D GPU particles sphere collision or attractor specified by the [param particles_collision] RID. Equivalent to [member GPUParticlesCollisionSphere3D.radius] or [member GPUParticlesAttractorSphere3D.radius] depending on the [param particles_collision] type.
*/
//go:nosplit
func (self class) ParticlesCollisionSetSphereRadius(particles_collision gd.RID, radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles_collision)
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_collision_set_sphere_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [param extents] for the 3D GPU particles collision by the [param particles_collision] RID. Equivalent to [member GPUParticlesCollisionBox3D.size], [member GPUParticlesCollisionSDF3D.size], [member GPUParticlesCollisionHeightField3D.size], [member GPUParticlesAttractorBox3D.size] or [member GPUParticlesAttractorVectorField3D.size] depending on the [param particles_collision] type.
*/
//go:nosplit
func (self class) ParticlesCollisionSetBoxExtents(particles_collision gd.RID, extents gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles_collision)
	callframe.Arg(frame, extents)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_collision_set_box_extents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [param strength] for the 3D GPU particles attractor specified by the [param particles_collision] RID. Only used for attractors, not colliders. Equivalent to [member GPUParticlesAttractor3D.strength].
*/
//go:nosplit
func (self class) ParticlesCollisionSetAttractorStrength(particles_collision gd.RID, strength gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles_collision)
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_collision_set_attractor_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the directionality [param amount] for the 3D GPU particles attractor specified by the [param particles_collision] RID. Only used for attractors, not colliders. Equivalent to [member GPUParticlesAttractor3D.directionality].
*/
//go:nosplit
func (self class) ParticlesCollisionSetAttractorDirectionality(particles_collision gd.RID, amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles_collision)
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_collision_set_attractor_directionality, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the attenuation [param curve] for the 3D GPU particles attractor specified by the [param particles_collision] RID. Only used for attractors, not colliders. Equivalent to [member GPUParticlesAttractor3D.attenuation].
*/
//go:nosplit
func (self class) ParticlesCollisionSetAttractorAttenuation(particles_collision gd.RID, curve gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles_collision)
	callframe.Arg(frame, curve)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_collision_set_attractor_attenuation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the signed distance field [param texture] for the 3D GPU particles collision specified by the [param particles_collision] RID. Equivalent to [member GPUParticlesCollisionSDF3D.texture] or [member GPUParticlesAttractorVectorField3D.texture] depending on the [param particles_collision] type.
*/
//go:nosplit
func (self class) ParticlesCollisionSetFieldTexture(particles_collision gd.RID, texture gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles_collision)
	callframe.Arg(frame, texture)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_collision_set_field_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Requests an update for the 3D GPU particle collision heightfield. This may be automatically called by the 3D GPU particle collision heightfield depending on its [member GPUParticlesCollisionHeightField3D.update_mode].
*/
//go:nosplit
func (self class) ParticlesCollisionHeightFieldUpdate(particles_collision gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles_collision)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_collision_height_field_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the heightmap [param resolution] for the 3D GPU particles heightfield collision specified by the [param particles_collision] RID. Equivalent to [member GPUParticlesCollisionHeightField3D.resolution].
*/
//go:nosplit
func (self class) ParticlesCollisionSetHeightFieldResolution(particles_collision gd.RID, resolution classdb.RenderingServerParticlesCollisionHeightfieldResolution)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles_collision)
	callframe.Arg(frame, resolution)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_particles_collision_set_height_field_resolution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a new fog volume and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]fog_volume_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
[b]Note:[/b] The equivalent node is [FogVolume].
*/
//go:nosplit
func (self class) FogVolumeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_fog_volume_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the shape of the fog volume to either [constant RenderingServer.FOG_VOLUME_SHAPE_ELLIPSOID], [constant RenderingServer.FOG_VOLUME_SHAPE_CONE], [constant RenderingServer.FOG_VOLUME_SHAPE_CYLINDER], [constant RenderingServer.FOG_VOLUME_SHAPE_BOX] or [constant RenderingServer.FOG_VOLUME_SHAPE_WORLD].
*/
//go:nosplit
func (self class) FogVolumeSetShape(fog_volume gd.RID, shape classdb.RenderingServerFogVolumeShape)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fog_volume)
	callframe.Arg(frame, shape)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_fog_volume_set_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the size of the fog volume when shape is [constant RenderingServer.FOG_VOLUME_SHAPE_ELLIPSOID], [constant RenderingServer.FOG_VOLUME_SHAPE_CONE], [constant RenderingServer.FOG_VOLUME_SHAPE_CYLINDER] or [constant RenderingServer.FOG_VOLUME_SHAPE_BOX].
*/
//go:nosplit
func (self class) FogVolumeSetSize(fog_volume gd.RID, size gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fog_volume)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_fog_volume_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [Material] of the fog volume. Can be either a [FogMaterial] or a custom [ShaderMaterial].
*/
//go:nosplit
func (self class) FogVolumeSetMaterial(fog_volume gd.RID, material gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fog_volume)
	callframe.Arg(frame, material)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_fog_volume_set_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a new 3D visibility notifier object and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]visibility_notifier_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
To place in a scene, attach this mesh to an instance using [method instance_set_base] using the returned RID.
[b]Note:[/b] The equivalent node is [VisibleOnScreenNotifier3D].
*/
//go:nosplit
func (self class) VisibilityNotifierCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_visibility_notifier_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) VisibilityNotifierSetAabb(notifier gd.RID, aabb gd.AABB)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, notifier)
	callframe.Arg(frame, aabb)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_visibility_notifier_set_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) VisibilityNotifierSetCallbacks(notifier gd.RID, enter_callable gd.Callable, exit_callable gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, notifier)
	callframe.Arg(frame, mmm.Get(enter_callable))
	callframe.Arg(frame, mmm.Get(exit_callable))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_visibility_notifier_set_callbacks, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates an occluder instance and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]occluder_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
[b]Note:[/b] The equivalent resource is [Occluder3D] (not to be confused with the [OccluderInstance3D] node).
*/
//go:nosplit
func (self class) OccluderCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_occluder_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the mesh data for the given occluder RID, which controls the shape of the occlusion culling that will be performed.
*/
//go:nosplit
func (self class) OccluderSetMesh(occluder gd.RID, vertices gd.PackedVector3Array, indices gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, occluder)
	callframe.Arg(frame, mmm.Get(vertices))
	callframe.Arg(frame, mmm.Get(indices))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_occluder_set_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a 3D camera and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]camera_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
[b]Note:[/b] The equivalent node is [Camera3D].
*/
//go:nosplit
func (self class) CameraCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_camera_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets camera to use perspective projection. Objects on the screen becomes smaller when they are far away.
*/
//go:nosplit
func (self class) CameraSetPerspective(camera gd.RID, fovy_degrees gd.Float, z_near gd.Float, z_far gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, camera)
	callframe.Arg(frame, fovy_degrees)
	callframe.Arg(frame, z_near)
	callframe.Arg(frame, z_far)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_camera_set_perspective, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets camera to use orthogonal projection, also known as orthographic projection. Objects remain the same size on the screen no matter how far away they are.
*/
//go:nosplit
func (self class) CameraSetOrthogonal(camera gd.RID, size gd.Float, z_near gd.Float, z_far gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, camera)
	callframe.Arg(frame, size)
	callframe.Arg(frame, z_near)
	callframe.Arg(frame, z_far)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_camera_set_orthogonal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets camera to use frustum projection. This mode allows adjusting the [param offset] argument to create "tilted frustum" effects.
*/
//go:nosplit
func (self class) CameraSetFrustum(camera gd.RID, size gd.Float, offset gd.Vector2, z_near gd.Float, z_far gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, camera)
	callframe.Arg(frame, size)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, z_near)
	callframe.Arg(frame, z_far)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_camera_set_frustum, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets [Transform3D] of camera.
*/
//go:nosplit
func (self class) CameraSetTransform(camera gd.RID, transform gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, camera)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_camera_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the cull mask associated with this camera. The cull mask describes which 3D layers are rendered by this camera. Equivalent to [member Camera3D.cull_mask].
*/
//go:nosplit
func (self class) CameraSetCullMask(camera gd.RID, layers gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, camera)
	callframe.Arg(frame, layers)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_camera_set_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the environment used by this camera. Equivalent to [member Camera3D.environment].
*/
//go:nosplit
func (self class) CameraSetEnvironment(camera gd.RID, env gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, camera)
	callframe.Arg(frame, env)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_camera_set_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the camera_attributes created with [method camera_attributes_create] to the given camera.
*/
//go:nosplit
func (self class) CameraSetCameraAttributes(camera gd.RID, effects gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, camera)
	callframe.Arg(frame, effects)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_camera_set_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the compositor used by this camera. Equivalent to [member Camera3D.compositor].
*/
//go:nosplit
func (self class) CameraSetCompositor(camera gd.RID, compositor gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, camera)
	callframe.Arg(frame, compositor)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_camera_set_compositor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], preserves the horizontal aspect ratio which is equivalent to [constant Camera3D.KEEP_WIDTH]. If [code]false[/code], preserves the vertical aspect ratio which is equivalent to [constant Camera3D.KEEP_HEIGHT].
*/
//go:nosplit
func (self class) CameraSetUseVerticalAspect(camera gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, camera)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_camera_set_use_vertical_aspect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates an empty viewport and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]viewport_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
[b]Note:[/b] The equivalent node is [Viewport].
*/
//go:nosplit
func (self class) ViewportCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], the viewport uses augmented or virtual reality technologies. See [XRInterface].
*/
//go:nosplit
func (self class) ViewportSetUseXr(viewport gd.RID, use_xr bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, use_xr)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_use_xr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the viewport's width and height in pixels.
*/
//go:nosplit
func (self class) ViewportSetSize(viewport gd.RID, width gd.Int, height gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], sets the viewport active, else sets it inactive.
*/
//go:nosplit
func (self class) ViewportSetActive(viewport gd.RID, active bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the viewport's parent to the viewport specified by the [param parent_viewport] RID.
*/
//go:nosplit
func (self class) ViewportSetParentViewport(viewport gd.RID, parent_viewport gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, parent_viewport)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_parent_viewport, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Copies the viewport to a region of the screen specified by [param rect]. If [method viewport_set_render_direct_to_screen] is [code]true[/code], then the viewport does not use a framebuffer and the contents of the viewport are rendered directly to screen. However, note that the root viewport is drawn last, therefore it will draw over the screen. Accordingly, you must set the root viewport to an area that does not cover the area that you have attached this viewport to.
For example, you can set the root viewport to not render at all with the following code:
FIXME: The method seems to be non-existent.
[codeblocks]
[gdscript]
func _ready():
    get_viewport().set_attach_to_screen_rect(Rect2())
    $Viewport.set_attach_to_screen_rect(Rect2(0, 0, 600, 600))
[/gdscript]
[/codeblocks]
Using this can result in significant optimization, especially on lower-end devices. However, it comes at the cost of having to manage your viewports manually. For further optimization, see [method viewport_set_render_direct_to_screen].
*/
//go:nosplit
func (self class) ViewportAttachToScreen(viewport gd.RID, rect gd.Rect2, screen gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, rect)
	callframe.Arg(frame, screen)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_attach_to_screen, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], render the contents of the viewport directly to screen. This allows a low-level optimization where you can skip drawing a viewport to the root viewport. While this optimization can result in a significant increase in speed (especially on older devices), it comes at a cost of usability. When this is enabled, you cannot read from the viewport or from the screen_texture. You also lose the benefit of certain window settings, such as the various stretch modes. Another consequence to be aware of is that in 2D the rendering happens in window coordinates, so if you have a viewport that is double the size of the window, and you set this, then only the portion that fits within the window will be drawn, no automatic scaling is possible, even if your game scene is significantly larger than the window size.
*/
//go:nosplit
func (self class) ViewportSetRenderDirectToScreen(viewport gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_render_direct_to_screen, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the rendering mask associated with this [Viewport]. Only [CanvasItem] nodes with a matching rendering visibility layer will be rendered by this [Viewport].
*/
//go:nosplit
func (self class) ViewportSetCanvasCullMask(viewport gd.RID, canvas_cull_mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, canvas_cull_mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_canvas_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the 3D resolution scaling mode. Bilinear scaling renders at different resolution to either undersample or supersample the viewport. FidelityFX Super Resolution 1.0, abbreviated to FSR, is an upscaling technology that produces high quality images at fast framerates by using a spatially aware upscaling algorithm. FSR is slightly more expensive than bilinear, but it produces significantly higher image quality. FSR should be used where possible.
*/
//go:nosplit
func (self class) ViewportSetScaling3dMode(viewport gd.RID, scaling_3d_mode classdb.RenderingServerViewportScaling3DMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, scaling_3d_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_scaling_3d_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Scales the 3D render buffer based on the viewport size uses an image filter specified in [enum ViewportScaling3DMode] to scale the output image to the full viewport size. Values lower than [code]1.0[/code] can be used to speed up 3D rendering at the cost of quality (undersampling). Values greater than [code]1.0[/code] are only valid for bilinear mode and can be used to improve 3D rendering quality at a high performance cost (supersampling). See also [enum ViewportMSAA] for multi-sample antialiasing, which is significantly cheaper but only smoothens the edges of polygons.
When using FSR upscaling, AMD recommends exposing the following values as preset options to users "Ultra Quality: 0.77", "Quality: 0.67", "Balanced: 0.59", "Performance: 0.5" instead of exposing the entire scale.
*/
//go:nosplit
func (self class) ViewportSetScaling3dScale(viewport gd.RID, scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_scaling_3d_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Determines how sharp the upscaled image will be when using the FSR upscaling mode. Sharpness halves with every whole number. Values go from 0.0 (sharpest) to 2.0. Values above 2.0 won't make a visible difference.
*/
//go:nosplit
func (self class) ViewportSetFsrSharpness(viewport gd.RID, sharpness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, sharpness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_fsr_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Affects the final texture sharpness by reading from a lower or higher mipmap (also called "texture LOD bias"). Negative values make mipmapped textures sharper but grainier when viewed at a distance, while positive values make mipmapped textures blurrier (even when up close). To get sharper textures at a distance without introducing too much graininess, set this between [code]-0.75[/code] and [code]0.0[/code]. Enabling temporal antialiasing ([member ProjectSettings.rendering/anti_aliasing/quality/use_taa]) can help reduce the graininess visible when using negative mipmap bias.
[b]Note:[/b] When the 3D scaling mode is set to FSR 1.0, this value is used to adjust the automatic mipmap bias which is calculated internally based on the scale factor. The formula for this is [code]-log2(1.0 / scale) + mipmap_bias[/code].
*/
//go:nosplit
func (self class) ViewportSetTextureMipmapBias(viewport gd.RID, mipmap_bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, mipmap_bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_texture_mipmap_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets when the viewport should be updated. See [enum ViewportUpdateMode] constants for options.
*/
//go:nosplit
func (self class) ViewportSetUpdateMode(viewport gd.RID, update_mode classdb.RenderingServerViewportUpdateMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, update_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the viewport's update mode. See [enum ViewportUpdateMode] constants for options.
[b]Warning:[/b] Calling this from any thread other than the rendering thread will be detrimental to performance.
*/
//go:nosplit
func (self class) ViewportGetUpdateMode(viewport gd.RID) classdb.RenderingServerViewportUpdateMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	var r_ret = callframe.Ret[classdb.RenderingServerViewportUpdateMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_get_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the clear mode of a viewport. See [enum ViewportClearMode] for options.
*/
//go:nosplit
func (self class) ViewportSetClearMode(viewport gd.RID, clear_mode classdb.RenderingServerViewportClearMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, clear_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_clear_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the render target for the viewport.
*/
//go:nosplit
func (self class) ViewportGetRenderTarget(viewport gd.RID) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_get_render_target, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the viewport's last rendered frame.
*/
//go:nosplit
func (self class) ViewportGetTexture(viewport gd.RID) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], the viewport's 3D elements are not rendered.
*/
//go:nosplit
func (self class) ViewportSetDisable3d(viewport gd.RID, disable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_disable_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], the viewport's canvas (i.e. 2D and GUI elements) is not rendered.
*/
//go:nosplit
func (self class) ViewportSetDisable2d(viewport gd.RID, disable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_disable_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the viewport's environment mode which allows enabling or disabling rendering of 3D environment over 2D canvas. When disabled, 2D will not be affected by the environment. When enabled, 2D will be affected by the environment if the environment background mode is [constant ENV_BG_CANVAS]. The default behavior is to inherit the setting from the viewport's parent. If the topmost parent is also set to [constant VIEWPORT_ENVIRONMENT_INHERIT], then the behavior will be the same as if it was set to [constant VIEWPORT_ENVIRONMENT_ENABLED].
*/
//go:nosplit
func (self class) ViewportSetEnvironmentMode(viewport gd.RID, mode classdb.RenderingServerViewportEnvironmentMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_environment_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a viewport's camera.
*/
//go:nosplit
func (self class) ViewportAttachCamera(viewport gd.RID, camera gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, camera)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_attach_camera, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a viewport's scenario. The scenario contains information about environment information, reflection atlas, etc.
*/
//go:nosplit
func (self class) ViewportSetScenario(viewport gd.RID, scenario gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, scenario)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_scenario, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a viewport's canvas.
*/
//go:nosplit
func (self class) ViewportAttachCanvas(viewport gd.RID, canvas gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, canvas)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_attach_canvas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Detaches a viewport from a canvas.
*/
//go:nosplit
func (self class) ViewportRemoveCanvas(viewport gd.RID, canvas gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, canvas)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_remove_canvas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], canvas item transforms (i.e. origin position) are snapped to the nearest pixel when rendering. This can lead to a crisper appearance at the cost of less smooth movement, especially when [Camera2D] smoothing is enabled. Equivalent to [member ProjectSettings.rendering/2d/snap/snap_2d_transforms_to_pixel].
*/
//go:nosplit
func (self class) ViewportSetSnap2dTransformsToPixel(viewport gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_snap_2d_transforms_to_pixel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], canvas item vertices (i.e. polygon points) are snapped to the nearest pixel when rendering. This can lead to a crisper appearance at the cost of less smooth movement, especially when [Camera2D] smoothing is enabled. Equivalent to [member ProjectSettings.rendering/2d/snap/snap_2d_vertices_to_pixel].
*/
//go:nosplit
func (self class) ViewportSetSnap2dVerticesToPixel(viewport gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_snap_2d_vertices_to_pixel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the default texture filtering mode for the specified [param viewport] RID. See [enum CanvasItemTextureFilter] for options.
*/
//go:nosplit
func (self class) ViewportSetDefaultCanvasItemTextureFilter(viewport gd.RID, filter classdb.RenderingServerCanvasItemTextureFilter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, filter)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_default_canvas_item_texture_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the default texture repeat mode for the specified [param viewport] RID. See [enum CanvasItemTextureRepeat] for options.
*/
//go:nosplit
func (self class) ViewportSetDefaultCanvasItemTextureRepeat(viewport gd.RID, repeat classdb.RenderingServerCanvasItemTextureRepeat)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, repeat)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_default_canvas_item_texture_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the transformation of a viewport's canvas.
*/
//go:nosplit
func (self class) ViewportSetCanvasTransform(viewport gd.RID, canvas gd.RID, offset gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_canvas_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the stacking order for a viewport's canvas.
[param layer] is the actual canvas layer, while [param sublayer] specifies the stacking order of the canvas among those in the same layer.
*/
//go:nosplit
func (self class) ViewportSetCanvasStacking(viewport gd.RID, canvas gd.RID, layer gd.Int, sublayer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, layer)
	callframe.Arg(frame, sublayer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_canvas_stacking, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], the viewport renders its background as transparent.
*/
//go:nosplit
func (self class) ViewportSetTransparentBackground(viewport gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_transparent_background, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the viewport's global transformation matrix.
*/
//go:nosplit
func (self class) ViewportSetGlobalCanvasTransform(viewport gd.RID, transform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_global_canvas_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the viewport's 2D signed distance field [member ProjectSettings.rendering/2d/sdf/oversize] and [member ProjectSettings.rendering/2d/sdf/scale]. This is used when sampling the signed distance field in [CanvasItem] shaders as well as [GPUParticles2D] collision. This is [i]not[/i] used by SDFGI in 3D rendering.
*/
//go:nosplit
func (self class) ViewportSetSdfOversizeAndScale(viewport gd.RID, oversize classdb.RenderingServerViewportSDFOversize, scale classdb.RenderingServerViewportSDFScale)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, oversize)
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_sdf_oversize_and_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [param size] of the shadow atlas's images (used for omni and spot lights) on the viewport specified by the [param viewport] RID. The value is rounded up to the nearest power of 2. If [param use_16_bits] is [code]true[/code], use 16 bits for the omni/spot shadow depth map. Enabling this results in shadows having less precision and may result in shadow acne, but can lead to performance improvements on some devices.
[b]Note:[/b] If this is set to [code]0[/code], no positional shadows will be visible at all. This can improve performance significantly on low-end systems by reducing both the CPU and GPU load (as fewer draw calls are needed to draw the scene without shadows).
*/
//go:nosplit
func (self class) ViewportSetPositionalShadowAtlasSize(viewport gd.RID, size gd.Int, use_16_bits bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, size)
	callframe.Arg(frame, use_16_bits)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_positional_shadow_atlas_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the number of subdivisions to use in the specified shadow atlas [param quadrant] for omni and spot shadows. See also [method Viewport.set_positional_shadow_atlas_quadrant_subdiv].
*/
//go:nosplit
func (self class) ViewportSetPositionalShadowAtlasQuadrantSubdivision(viewport gd.RID, quadrant gd.Int, subdivision gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, quadrant)
	callframe.Arg(frame, subdivision)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_positional_shadow_atlas_quadrant_subdivision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the multisample anti-aliasing mode for 3D on the specified [param viewport] RID. See [enum ViewportMSAA] for options.
*/
//go:nosplit
func (self class) ViewportSetMsaa3d(viewport gd.RID, msaa classdb.RenderingServerViewportMSAA)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, msaa)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_msaa_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the multisample anti-aliasing mode for 2D/Canvas on the specified [param viewport] RID. See [enum ViewportMSAA] for options.
*/
//go:nosplit
func (self class) ViewportSetMsaa2d(viewport gd.RID, msaa classdb.RenderingServerViewportMSAA)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, msaa)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_msaa_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], 2D rendering will use a high dynamic range (HDR) format framebuffer matching the bit depth of the 3D framebuffer. When using the Forward+ renderer this will be an [code]RGBA16[/code] framebuffer, while when using the Mobile renderer it will be an [code]RGB10_A2[/code] framebuffer. Additionally, 2D rendering will take place in linear color space and will be converted to sRGB space immediately before blitting to the screen (if the Viewport is attached to the screen). Practically speaking, this means that the end result of the Viewport will not be clamped into the [code]0-1[/code] range and can be used in 3D rendering without color space adjustments. This allows 2D rendering to take advantage of effects requiring high dynamic range (e.g. 2D glow) as well as substantially improves the appearance of effects requiring highly detailed gradients. This setting has the same effect as [member Viewport.use_hdr_2d].
[b]Note:[/b] This setting will have no effect when using the GL Compatibility renderer as the GL Compatibility renderer always renders in low dynamic range for performance reasons.
*/
//go:nosplit
func (self class) ViewportSetUseHdr2d(viewport gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_use_hdr_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the viewport's screen-space antialiasing mode.
*/
//go:nosplit
func (self class) ViewportSetScreenSpaceAa(viewport gd.RID, mode classdb.RenderingServerViewportScreenSpaceAA)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_screen_space_aa, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], use Temporal Anti-Aliasing. Equivalent to [member ProjectSettings.rendering/anti_aliasing/quality/use_taa].
*/
//go:nosplit
func (self class) ViewportSetUseTaa(viewport gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_use_taa, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], enables debanding on the specified viewport. Equivalent to [member ProjectSettings.rendering/anti_aliasing/quality/use_debanding].
*/
//go:nosplit
func (self class) ViewportSetUseDebanding(viewport gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_use_debanding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], enables occlusion culling on the specified viewport. Equivalent to [member ProjectSettings.rendering/occlusion_culling/use_occlusion_culling].
*/
//go:nosplit
func (self class) ViewportSetUseOcclusionCulling(viewport gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_use_occlusion_culling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [member ProjectSettings.rendering/occlusion_culling/occlusion_rays_per_thread] to use for occlusion culling. This parameter is global and cannot be set on a per-viewport basis.
*/
//go:nosplit
func (self class) ViewportSetOcclusionRaysPerThread(rays_per_thread gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rays_per_thread)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_occlusion_rays_per_thread, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [member ProjectSettings.rendering/occlusion_culling/bvh_build_quality] to use for occlusion culling. This parameter is global and cannot be set on a per-viewport basis.
*/
//go:nosplit
func (self class) ViewportSetOcclusionCullingBuildQuality(quality classdb.RenderingServerViewportOcclusionCullingBuildQuality)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, quality)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_occlusion_culling_build_quality, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a statistic about the rendering engine which can be used for performance profiling. This is separated into render pass [param type]s, each of them having the same [param info]s you can query (different passes will return different values). See [enum RenderingServer.ViewportRenderInfoType] for a list of render pass types and [enum RenderingServer.ViewportRenderInfo] for a list of information that can be queried.
See also [method get_rendering_info], which returns global information across all viewports.
[b]Note:[/b] Viewport rendering information is not available until at least 2 frames have been rendered by the engine. If rendering information is not available, [method viewport_get_render_info] returns [code]0[/code]. To print rendering information in [code]_ready()[/code] successfully, use the following:
[codeblock]
func _ready():
    for _i in 2:
        await get_tree().process_frame

    print(
            RenderingServer.viewport_get_render_info(get_viewport().get_viewport_rid(),
            RenderingServer.VIEWPORT_RENDER_INFO_TYPE_VISIBLE,
            RenderingServer.VIEWPORT_RENDER_INFO_DRAW_CALLS_IN_FRAME)
    )
[/codeblock]
*/
//go:nosplit
func (self class) ViewportGetRenderInfo(viewport gd.RID, atype classdb.RenderingServerViewportRenderInfoType, info classdb.RenderingServerViewportRenderInfo) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, atype)
	callframe.Arg(frame, info)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_get_render_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the debug draw mode of a viewport. See [enum ViewportDebugDraw] for options.
*/
//go:nosplit
func (self class) ViewportSetDebugDraw(viewport gd.RID, draw classdb.RenderingServerViewportDebugDraw)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, draw)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_debug_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the measurement for the given [param viewport] RID (obtained using [method Viewport.get_viewport_rid]). Once enabled, [method viewport_get_measured_render_time_cpu] and [method viewport_get_measured_render_time_gpu] will return values greater than [code]0.0[/code] when queried with the given [param viewport].
*/
//go:nosplit
func (self class) ViewportSetMeasureRenderTime(viewport gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_measure_render_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the CPU time taken to render the last frame in milliseconds. This [i]only[/i] includes time spent in rendering-related operations; scripts' [code]_process[/code] functions and other engine subsystems are not included in this readout. To get a complete readout of CPU time spent to render the scene, sum the render times of all viewports that are drawn every frame plus [method get_frame_setup_time_cpu]. Unlike [method Engine.get_frames_per_second], this method will accurately reflect CPU utilization even if framerate is capped via V-Sync or [member Engine.max_fps]. See also [method viewport_get_measured_render_time_gpu].
[b]Note:[/b] Requires measurements to be enabled on the specified [param viewport] using [method viewport_set_measure_render_time]. Otherwise, this method returns [code]0.0[/code].
*/
//go:nosplit
func (self class) ViewportGetMeasuredRenderTimeCpu(viewport gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_get_measured_render_time_cpu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the GPU time taken to render the last frame in milliseconds. To get a complete readout of GPU time spent to render the scene, sum the render times of all viewports that are drawn every frame. Unlike [method Engine.get_frames_per_second], this method accurately reflects GPU utilization even if framerate is capped via V-Sync or [member Engine.max_fps]. See also [method viewport_get_measured_render_time_gpu].
[b]Note:[/b] Requires measurements to be enabled on the specified [param viewport] using [method viewport_set_measure_render_time]. Otherwise, this method returns [code]0.0[/code].
[b]Note:[/b] When GPU utilization is low enough during a certain period of time, GPUs will decrease their power state (which in turn decreases core and memory clock speeds). This can cause the reported GPU time to increase if GPU utilization is kept low enough by a framerate cap (compared to what it would be at the GPU's highest power state). Keep this in mind when benchmarking using [method viewport_get_measured_render_time_gpu]. This behavior can be overridden in the graphics driver settings at the cost of higher power usage.
*/
//go:nosplit
func (self class) ViewportGetMeasuredRenderTimeGpu(viewport gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_get_measured_render_time_gpu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the Variable Rate Shading (VRS) mode for the viewport. If the GPU does not support VRS, this property is ignored. Equivalent to [member ProjectSettings.rendering/vrs/mode].
*/
//go:nosplit
func (self class) ViewportSetVrsMode(viewport gd.RID, mode classdb.RenderingServerViewportVRSMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_vrs_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the update mode for Variable Rate Shading (VRS) for the viewport. VRS requires the input texture to be converted to the format usable by the VRS method supported by the hardware. The update mode defines how often this happens. If the GPU does not support VRS, or VRS is not enabled, this property is ignored.
If set to [constant RenderingServer.VIEWPORT_VRS_UPDATE_ONCE], the input texture is copied once and the mode is changed to [constant RenderingServer.VIEWPORT_VRS_UPDATE_DISABLED].
*/
//go:nosplit
func (self class) ViewportSetVrsUpdateMode(viewport gd.RID, mode classdb.RenderingServerViewportVRSUpdateMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_vrs_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
The texture to use when the VRS mode is set to [constant RenderingServer.VIEWPORT_VRS_TEXTURE]. Equivalent to [member ProjectSettings.rendering/vrs/texture].
*/
//go:nosplit
func (self class) ViewportSetVrsTexture(viewport gd.RID, texture gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, viewport)
	callframe.Arg(frame, texture)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_viewport_set_vrs_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates an empty sky and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]sky_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
*/
//go:nosplit
func (self class) SkyCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_sky_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [param radiance_size] of the sky specified by the [param sky] RID (in pixels). Equivalent to [member Sky.radiance_size].
*/
//go:nosplit
func (self class) SkySetRadianceSize(sky gd.RID, radiance_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sky)
	callframe.Arg(frame, radiance_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_sky_set_radiance_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the process [param mode] of the sky specified by the [param sky] RID. Equivalent to [member Sky.process_mode].
*/
//go:nosplit
func (self class) SkySetMode(sky gd.RID, mode classdb.RenderingServerSkyMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sky)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_sky_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the material that the sky uses to render the background, ambient and reflection maps.
*/
//go:nosplit
func (self class) SkySetMaterial(sky gd.RID, material gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sky)
	callframe.Arg(frame, material)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_sky_set_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Generates and returns an [Image] containing the radiance map for the specified [param sky] RID. This supports built-in sky material and custom sky shaders. If [param bake_irradiance] is [code]true[/code], the irradiance map is saved instead of the radiance map. The radiance map is used to render reflected light, while the irradiance map is used to render ambient light. See also [method environment_bake_panorama].
[b]Note:[/b] The image is saved in linear color space without any tonemapping performed, which means it will look too dark if viewed directly in an image editor. [param energy] values above [code]1.0[/code] can be used to brighten the resulting image.
[b]Note:[/b] [param size] should be a 2:1 aspect ratio for the generated panorama to have square pixels. For radiance maps, there is no point in using a height greater than [member Sky.radiance_size], as it won't increase detail. Irradiance maps only contain low-frequency data, so there is usually no point in going past a size of 128×64 pixels when saving an irradiance map.
*/
//go:nosplit
func (self class) SkyBakePanorama(ctx gd.Lifetime, sky gd.RID, energy gd.Float, bake_irradiance bool, size gd.Vector2i) object.Image {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sky)
	callframe.Arg(frame, energy)
	callframe.Arg(frame, bake_irradiance)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_sky_bake_panorama, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a new rendering effect and adds it to the RenderingServer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
*/
//go:nosplit
func (self class) CompositorEffectCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_compositor_effect_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Enables/disables this rendering effect.
*/
//go:nosplit
func (self class) CompositorEffectSetEnabled(effect gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, effect)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_compositor_effect_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the callback type ([param callback_type]) and callback method([param callback]) for this rendering effect.
*/
//go:nosplit
func (self class) CompositorEffectSetCallback(effect gd.RID, callback_type classdb.RenderingServerCompositorEffectCallbackType, callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, effect)
	callframe.Arg(frame, callback_type)
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_compositor_effect_set_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the flag ([param flag]) for this rendering effect to [code]true[/code] or [code]false[/code] ([param set]).
*/
//go:nosplit
func (self class) CompositorEffectSetFlag(effect gd.RID, flag classdb.RenderingServerCompositorEffectFlags, set bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, effect)
	callframe.Arg(frame, flag)
	callframe.Arg(frame, set)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_compositor_effect_set_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a new compositor and adds it to the RenderingServer. It can be accessed with the RID that is returned.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
*/
//go:nosplit
func (self class) CompositorCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_compositor_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the compositor effects for the specified compositor RID. [param effects] should be an array containing RIDs created with [method compositor_effect_create].
*/
//go:nosplit
func (self class) CompositorSetCompositorEffects(compositor gd.RID, effects gd.ArrayOf[gd.RID])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, compositor)
	callframe.Arg(frame, mmm.Get(effects))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_compositor_set_compositor_effects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates an environment and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]environment_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
[b]Note:[/b] The equivalent resource is [Environment].
*/
//go:nosplit
func (self class) EnvironmentCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the environment's background mode. Equivalent to [member Environment.background_mode].
*/
//go:nosplit
func (self class) EnvironmentSetBackground(env gd.RID, bg classdb.RenderingServerEnvironmentBG)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, env)
	callframe.Arg(frame, bg)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_background, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [Sky] to be used as the environment's background when using [i]BGMode[/i] sky. Equivalent to [member Environment.sky].
*/
//go:nosplit
func (self class) EnvironmentSetSky(env gd.RID, sky gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, env)
	callframe.Arg(frame, sky)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_sky, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a custom field of view for the background [Sky]. Equivalent to [member Environment.sky_custom_fov].
*/
//go:nosplit
func (self class) EnvironmentSetSkyCustomFov(env gd.RID, scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, env)
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_sky_custom_fov, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the rotation of the background [Sky] expressed as a [Basis]. Equivalent to [member Environment.sky_rotation], where the rotation vector is used to construct the [Basis].
*/
//go:nosplit
func (self class) EnvironmentSetSkyOrientation(env gd.RID, orientation gd.Basis)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, env)
	callframe.Arg(frame, orientation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_sky_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Color displayed for clear areas of the scene. Only effective if using the [constant ENV_BG_COLOR] background mode.
*/
//go:nosplit
func (self class) EnvironmentSetBgColor(env gd.RID, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, env)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the intensity of the background color.
*/
//go:nosplit
func (self class) EnvironmentSetBgEnergy(env gd.RID, multiplier gd.Float, exposure_value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, env)
	callframe.Arg(frame, multiplier)
	callframe.Arg(frame, exposure_value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_bg_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the maximum layer to use if using Canvas background mode.
*/
//go:nosplit
func (self class) EnvironmentSetCanvasMaxLayer(env gd.RID, max_layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, env)
	callframe.Arg(frame, max_layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_canvas_max_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the values to be used for ambient light rendering. See [Environment] for more details.
*/
//go:nosplit
func (self class) EnvironmentSetAmbientLight(env gd.RID, color gd.Color, ambient classdb.RenderingServerEnvironmentAmbientSource, energy gd.Float, sky_contibution gd.Float, reflection_source classdb.RenderingServerEnvironmentReflectionSource)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, env)
	callframe.Arg(frame, color)
	callframe.Arg(frame, ambient)
	callframe.Arg(frame, energy)
	callframe.Arg(frame, sky_contibution)
	callframe.Arg(frame, reflection_source)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_ambient_light, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Configures glow for the specified environment RID. See [code]glow_*[/code] properties in [Environment] for more information.
*/
//go:nosplit
func (self class) EnvironmentSetGlow(env gd.RID, enable bool, levels gd.PackedFloat32Array, intensity gd.Float, strength gd.Float, mix gd.Float, bloom_threshold gd.Float, blend_mode classdb.RenderingServerEnvironmentGlowBlendMode, hdr_bleed_threshold gd.Float, hdr_bleed_scale gd.Float, hdr_luminance_cap gd.Float, glow_map_strength gd.Float, glow_map gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, env)
	callframe.Arg(frame, enable)
	callframe.Arg(frame, mmm.Get(levels))
	callframe.Arg(frame, intensity)
	callframe.Arg(frame, strength)
	callframe.Arg(frame, mix)
	callframe.Arg(frame, bloom_threshold)
	callframe.Arg(frame, blend_mode)
	callframe.Arg(frame, hdr_bleed_threshold)
	callframe.Arg(frame, hdr_bleed_scale)
	callframe.Arg(frame, hdr_luminance_cap)
	callframe.Arg(frame, glow_map_strength)
	callframe.Arg(frame, glow_map)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_glow, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the variables to be used with the "tonemap" post-process effect. See [Environment] for more details.
*/
//go:nosplit
func (self class) EnvironmentSetTonemap(env gd.RID, tone_mapper classdb.RenderingServerEnvironmentToneMapper, exposure gd.Float, white gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, env)
	callframe.Arg(frame, tone_mapper)
	callframe.Arg(frame, exposure)
	callframe.Arg(frame, white)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_tonemap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the values to be used with the "adjustments" post-process effect. See [Environment] for more details.
*/
//go:nosplit
func (self class) EnvironmentSetAdjustment(env gd.RID, enable bool, brightness gd.Float, contrast gd.Float, saturation gd.Float, use_1d_color_correction bool, color_correction gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, env)
	callframe.Arg(frame, enable)
	callframe.Arg(frame, brightness)
	callframe.Arg(frame, contrast)
	callframe.Arg(frame, saturation)
	callframe.Arg(frame, use_1d_color_correction)
	callframe.Arg(frame, color_correction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_adjustment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the variables to be used with the screen-space reflections (SSR) post-process effect. See [Environment] for more details.
*/
//go:nosplit
func (self class) EnvironmentSetSsr(env gd.RID, enable bool, max_steps gd.Int, fade_in gd.Float, fade_out gd.Float, depth_tolerance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, env)
	callframe.Arg(frame, enable)
	callframe.Arg(frame, max_steps)
	callframe.Arg(frame, fade_in)
	callframe.Arg(frame, fade_out)
	callframe.Arg(frame, depth_tolerance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_ssr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the variables to be used with the screen-space ambient occlusion (SSAO) post-process effect. See [Environment] for more details.
*/
//go:nosplit
func (self class) EnvironmentSetSsao(env gd.RID, enable bool, radius gd.Float, intensity gd.Float, power gd.Float, detail gd.Float, horizon gd.Float, sharpness gd.Float, light_affect gd.Float, ao_channel_affect gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, env)
	callframe.Arg(frame, enable)
	callframe.Arg(frame, radius)
	callframe.Arg(frame, intensity)
	callframe.Arg(frame, power)
	callframe.Arg(frame, detail)
	callframe.Arg(frame, horizon)
	callframe.Arg(frame, sharpness)
	callframe.Arg(frame, light_affect)
	callframe.Arg(frame, ao_channel_affect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_ssao, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Configures fog for the specified environment RID. See [code]fog_*[/code] properties in [Environment] for more information.
*/
//go:nosplit
func (self class) EnvironmentSetFog(env gd.RID, enable bool, light_color gd.Color, light_energy gd.Float, sun_scatter gd.Float, density gd.Float, height gd.Float, height_density gd.Float, aerial_perspective gd.Float, sky_affect gd.Float, fog_mode classdb.RenderingServerEnvironmentFogMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, env)
	callframe.Arg(frame, enable)
	callframe.Arg(frame, light_color)
	callframe.Arg(frame, light_energy)
	callframe.Arg(frame, sun_scatter)
	callframe.Arg(frame, density)
	callframe.Arg(frame, height)
	callframe.Arg(frame, height_density)
	callframe.Arg(frame, aerial_perspective)
	callframe.Arg(frame, sky_affect)
	callframe.Arg(frame, fog_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_fog, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Configures signed distance field global illumination for the specified environment RID. See [code]sdfgi_*[/code] properties in [Environment] for more information.
*/
//go:nosplit
func (self class) EnvironmentSetSdfgi(env gd.RID, enable bool, cascades gd.Int, min_cell_size gd.Float, y_scale classdb.RenderingServerEnvironmentSDFGIYScale, use_occlusion bool, bounce_feedback gd.Float, read_sky bool, energy gd.Float, normal_bias gd.Float, probe_bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, env)
	callframe.Arg(frame, enable)
	callframe.Arg(frame, cascades)
	callframe.Arg(frame, min_cell_size)
	callframe.Arg(frame, y_scale)
	callframe.Arg(frame, use_occlusion)
	callframe.Arg(frame, bounce_feedback)
	callframe.Arg(frame, read_sky)
	callframe.Arg(frame, energy)
	callframe.Arg(frame, normal_bias)
	callframe.Arg(frame, probe_bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_sdfgi, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the variables to be used with the volumetric fog post-process effect. See [Environment] for more details.
*/
//go:nosplit
func (self class) EnvironmentSetVolumetricFog(env gd.RID, enable bool, density gd.Float, albedo gd.Color, emission gd.Color, emission_energy gd.Float, anisotropy gd.Float, length gd.Float, p_detail_spread gd.Float, gi_inject gd.Float, temporal_reprojection bool, temporal_reprojection_amount gd.Float, ambient_inject gd.Float, sky_affect gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, env)
	callframe.Arg(frame, enable)
	callframe.Arg(frame, density)
	callframe.Arg(frame, albedo)
	callframe.Arg(frame, emission)
	callframe.Arg(frame, emission_energy)
	callframe.Arg(frame, anisotropy)
	callframe.Arg(frame, length)
	callframe.Arg(frame, p_detail_spread)
	callframe.Arg(frame, gi_inject)
	callframe.Arg(frame, temporal_reprojection)
	callframe.Arg(frame, temporal_reprojection_amount)
	callframe.Arg(frame, ambient_inject)
	callframe.Arg(frame, sky_affect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_volumetric_fog, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [param enable] is [code]true[/code], enables bicubic upscaling for glow which improves quality at the cost of performance. Equivalent to [member ProjectSettings.rendering/environment/glow/upscale_mode].
*/
//go:nosplit
func (self class) EnvironmentGlowSetUseBicubicUpscale(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_glow_set_use_bicubic_upscale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) EnvironmentSetSsrRoughnessQuality(quality classdb.RenderingServerEnvironmentSSRRoughnessQuality)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, quality)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_ssr_roughness_quality, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the quality level of the screen-space ambient occlusion (SSAO) post-process effect. See [Environment] for more details.
*/
//go:nosplit
func (self class) EnvironmentSetSsaoQuality(quality classdb.RenderingServerEnvironmentSSAOQuality, half_size bool, adaptive_target gd.Float, blur_passes gd.Int, fadeout_from gd.Float, fadeout_to gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, quality)
	callframe.Arg(frame, half_size)
	callframe.Arg(frame, adaptive_target)
	callframe.Arg(frame, blur_passes)
	callframe.Arg(frame, fadeout_from)
	callframe.Arg(frame, fadeout_to)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_ssao_quality, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the quality level of the screen-space indirect lighting (SSIL) post-process effect. See [Environment] for more details.
*/
//go:nosplit
func (self class) EnvironmentSetSsilQuality(quality classdb.RenderingServerEnvironmentSSILQuality, half_size bool, adaptive_target gd.Float, blur_passes gd.Int, fadeout_from gd.Float, fadeout_to gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, quality)
	callframe.Arg(frame, half_size)
	callframe.Arg(frame, adaptive_target)
	callframe.Arg(frame, blur_passes)
	callframe.Arg(frame, fadeout_from)
	callframe.Arg(frame, fadeout_to)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_ssil_quality, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the number of rays to throw per frame when computing signed distance field global illumination. Equivalent to [member ProjectSettings.rendering/global_illumination/sdfgi/probe_ray_count].
*/
//go:nosplit
func (self class) EnvironmentSetSdfgiRayCount(ray_count classdb.RenderingServerEnvironmentSDFGIRayCount)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ray_count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_sdfgi_ray_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the number of frames to use for converging signed distance field global illumination. Equivalent to [member ProjectSettings.rendering/global_illumination/sdfgi/frames_to_converge].
*/
//go:nosplit
func (self class) EnvironmentSetSdfgiFramesToConverge(frames classdb.RenderingServerEnvironmentSDFGIFramesToConverge)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, frames)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_sdfgi_frames_to_converge, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the update speed for dynamic lights' indirect lighting when computing signed distance field global illumination. Equivalent to [member ProjectSettings.rendering/global_illumination/sdfgi/frames_to_update_lights].
*/
//go:nosplit
func (self class) EnvironmentSetSdfgiFramesToUpdateLight(frames classdb.RenderingServerEnvironmentSDFGIFramesToUpdateLight)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, frames)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_sdfgi_frames_to_update_light, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the resolution of the volumetric fog's froxel buffer. [param size] is modified by the screen's aspect ratio and then used to set the width and height of the buffer. While [param depth] is directly used to set the depth of the buffer.
*/
//go:nosplit
func (self class) EnvironmentSetVolumetricFogVolumeSize(size gd.Int, depth gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, depth)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_volumetric_fog_volume_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Enables filtering of the volumetric fog scattering buffer. This results in much smoother volumes with very few under-sampling artifacts.
*/
//go:nosplit
func (self class) EnvironmentSetVolumetricFogFilterActive(active bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_set_volumetric_fog_filter_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Generates and returns an [Image] containing the radiance map for the specified [param environment] RID's sky. This supports built-in sky material and custom sky shaders. If [param bake_irradiance] is [code]true[/code], the irradiance map is saved instead of the radiance map. The radiance map is used to render reflected light, while the irradiance map is used to render ambient light. See also [method sky_bake_panorama].
[b]Note:[/b] The image is saved in linear color space without any tonemapping performed, which means it will look too dark if viewed directly in an image editor.
[b]Note:[/b] [param size] should be a 2:1 aspect ratio for the generated panorama to have square pixels. For radiance maps, there is no point in using a height greater than [member Sky.radiance_size], as it won't increase detail. Irradiance maps only contain low-frequency data, so there is usually no point in going past a size of 128×64 pixels when saving an irradiance map.
*/
//go:nosplit
func (self class) EnvironmentBakePanorama(ctx gd.Lifetime, environment gd.RID, bake_irradiance bool, size gd.Vector2i) object.Image {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, environment)
	callframe.Arg(frame, bake_irradiance)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_environment_bake_panorama, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the screen-space roughness limiter parameters, such as whether it should be enabled and its thresholds. Equivalent to [member ProjectSettings.rendering/anti_aliasing/screen_space_roughness_limiter/enabled], [member ProjectSettings.rendering/anti_aliasing/screen_space_roughness_limiter/amount] and [member ProjectSettings.rendering/anti_aliasing/screen_space_roughness_limiter/limit].
*/
//go:nosplit
func (self class) ScreenSpaceRoughnessLimiterSetActive(enable bool, amount gd.Float, limit gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	callframe.Arg(frame, amount)
	callframe.Arg(frame, limit)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_screen_space_roughness_limiter_set_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets [member ProjectSettings.rendering/environment/subsurface_scattering/subsurface_scattering_quality] to use when rendering materials that have subsurface scattering enabled.
*/
//go:nosplit
func (self class) SubSurfaceScatteringSetQuality(quality classdb.RenderingServerSubSurfaceScatteringQuality)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, quality)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_sub_surface_scattering_set_quality, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [member ProjectSettings.rendering/environment/subsurface_scattering/subsurface_scattering_scale] and [member ProjectSettings.rendering/environment/subsurface_scattering/subsurface_scattering_depth_scale] to use when rendering materials that have subsurface scattering enabled.
*/
//go:nosplit
func (self class) SubSurfaceScatteringSetScale(scale gd.Float, depth_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	callframe.Arg(frame, depth_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_sub_surface_scattering_set_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a camera attributes object and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]camera_attributes_[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
[b]Note:[/b] The equivalent resource is [CameraAttributes].
*/
//go:nosplit
func (self class) CameraAttributesCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_camera_attributes_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the quality level of the DOF blur effect to one of the options in [enum DOFBlurQuality]. [param use_jitter] can be used to jitter samples taken during the blur pass to hide artifacts at the cost of looking more fuzzy.
*/
//go:nosplit
func (self class) CameraAttributesSetDofBlurQuality(quality classdb.RenderingServerDOFBlurQuality, use_jitter bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, quality)
	callframe.Arg(frame, use_jitter)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_camera_attributes_set_dof_blur_quality, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the shape of the DOF bokeh pattern. Different shapes may be used to achieve artistic effect, or to meet performance targets. For more detail on available options see [enum DOFBokehShape].
*/
//go:nosplit
func (self class) CameraAttributesSetDofBlurBokehShape(shape classdb.RenderingServerDOFBokehShape)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_camera_attributes_set_dof_blur_bokeh_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the parameters to use with the DOF blur effect. These parameters take on the same meaning as their counterparts in [CameraAttributesPractical].
*/
//go:nosplit
func (self class) CameraAttributesSetDofBlur(camera_attributes gd.RID, far_enable bool, far_distance gd.Float, far_transition gd.Float, near_enable bool, near_distance gd.Float, near_transition gd.Float, amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, camera_attributes)
	callframe.Arg(frame, far_enable)
	callframe.Arg(frame, far_distance)
	callframe.Arg(frame, far_transition)
	callframe.Arg(frame, near_enable)
	callframe.Arg(frame, near_distance)
	callframe.Arg(frame, near_transition)
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_camera_attributes_set_dof_blur, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the exposure values that will be used by the renderers. The normalization amount is used to bake a given Exposure Value (EV) into rendering calculations to reduce the dynamic range of the scene.
The normalization factor can be calculated from exposure value (EV100) as follows:
[codeblock]
func get_exposure_normalization(ev100: float):
    return 1.0 / (pow(2.0, ev100) * 1.2)
[/codeblock]
The exposure value can be calculated from aperture (in f-stops), shutter speed (in seconds), and sensitivity (in ISO) as follows:
[codeblock]
func get_exposure(aperture: float, shutter_speed: float, sensitivity: float):
    return log((aperture * aperture) / shutter_speed * (100.0 / sensitivity)) / log(2)
[/codeblock]
*/
//go:nosplit
func (self class) CameraAttributesSetExposure(camera_attributes gd.RID, multiplier gd.Float, normalization gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, camera_attributes)
	callframe.Arg(frame, multiplier)
	callframe.Arg(frame, normalization)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_camera_attributes_set_exposure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the parameters to use with the auto-exposure effect. These parameters take on the same meaning as their counterparts in [CameraAttributes] and [CameraAttributesPractical].
*/
//go:nosplit
func (self class) CameraAttributesSetAutoExposure(camera_attributes gd.RID, enable bool, min_sensitivity gd.Float, max_sensitivity gd.Float, speed gd.Float, scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, camera_attributes)
	callframe.Arg(frame, enable)
	callframe.Arg(frame, min_sensitivity)
	callframe.Arg(frame, max_sensitivity)
	callframe.Arg(frame, speed)
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_camera_attributes_set_auto_exposure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a scenario and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]scenario_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
The scenario is the 3D world that all the visual instances exist in.
*/
//go:nosplit
func (self class) ScenarioCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_scenario_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the environment that will be used with this scenario. See also [Environment].
*/
//go:nosplit
func (self class) ScenarioSetEnvironment(scenario gd.RID, environment gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scenario)
	callframe.Arg(frame, environment)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_scenario_set_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the fallback environment to be used by this scenario. The fallback environment is used if no environment is set. Internally, this is used by the editor to provide a default environment.
*/
//go:nosplit
func (self class) ScenarioSetFallbackEnvironment(scenario gd.RID, environment gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scenario)
	callframe.Arg(frame, environment)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_scenario_set_fallback_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the camera attributes ([param effects]) that will be used with this scenario. See also [CameraAttributes].
*/
//go:nosplit
func (self class) ScenarioSetCameraAttributes(scenario gd.RID, effects gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scenario)
	callframe.Arg(frame, effects)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_scenario_set_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the compositor ([param compositor]) that will be used with this scenario. See also [Compositor].
*/
//go:nosplit
func (self class) ScenarioSetCompositor(scenario gd.RID, compositor gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scenario)
	callframe.Arg(frame, compositor)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_scenario_set_compositor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a visual instance, adds it to the RenderingServer, and sets both base and scenario. It can be accessed with the RID that is returned. This RID will be used in all [code]instance_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method. This is a shorthand for using [method instance_create] and setting the base and scenario manually.
*/
//go:nosplit
func (self class) InstanceCreate2(base gd.RID, scenario gd.RID) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, base)
	callframe.Arg(frame, scenario)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_create2, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a visual instance and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]instance_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
An instance is a way of placing a 3D object in the scenario. Objects like particles, meshes, reflection probes and decals need to be associated with an instance to be visible in the scenario using [method instance_set_base].
[b]Note:[/b] The equivalent node is [VisualInstance3D].
*/
//go:nosplit
func (self class) InstanceCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the base of the instance. A base can be any of the 3D objects that are created in the RenderingServer that can be displayed. For example, any of the light types, mesh, multimesh, particle system, reflection probe, decal, lightmap, voxel GI and visibility notifiers are all types that can be set as the base of an instance in order to be displayed in the scenario.
*/
//go:nosplit
func (self class) InstanceSetBase(instance gd.RID, base gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, base)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_set_base, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the scenario that the instance is in. The scenario is the 3D world that the objects will be displayed in.
*/
//go:nosplit
func (self class) InstanceSetScenario(instance gd.RID, scenario gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, scenario)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_set_scenario, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the render layers that this instance will be drawn to. Equivalent to [member VisualInstance3D.layers].
*/
//go:nosplit
func (self class) InstanceSetLayerMask(instance gd.RID, mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_set_layer_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the sorting offset and switches between using the bounding box or instance origin for depth sorting.
*/
//go:nosplit
func (self class) InstanceSetPivotData(instance gd.RID, sorting_offset gd.Float, use_aabb_center bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, sorting_offset)
	callframe.Arg(frame, use_aabb_center)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_set_pivot_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the world space transform of the instance. Equivalent to [member Node3D.global_transform].
*/
//go:nosplit
func (self class) InstanceSetTransform(instance gd.RID, transform gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Attaches a unique Object ID to instance. Object ID must be attached to instance for proper culling with [method instances_cull_aabb], [method instances_cull_convex], and [method instances_cull_ray].
*/
//go:nosplit
func (self class) InstanceAttachObjectInstanceId(instance gd.RID, id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_attach_object_instance_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the weight for a given blend shape associated with this instance.
*/
//go:nosplit
func (self class) InstanceSetBlendShapeWeight(instance gd.RID, shape gd.Int, weight gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, shape)
	callframe.Arg(frame, weight)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_set_blend_shape_weight, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the override material of a specific surface. Equivalent to [method MeshInstance3D.set_surface_override_material].
*/
//go:nosplit
func (self class) InstanceSetSurfaceOverrideMaterial(instance gd.RID, surface gd.Int, material gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, surface)
	callframe.Arg(frame, material)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_set_surface_override_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets whether an instance is drawn or not. Equivalent to [member Node3D.visible].
*/
//go:nosplit
func (self class) InstanceSetVisible(instance gd.RID, visible bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_set_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the transparency for the given geometry instance. Equivalent to [member GeometryInstance3D.transparency].
A transparency of [code]0.0[/code] is fully opaque, while [code]1.0[/code] is fully transparent. Values greater than [code]0.0[/code] (exclusive) will force the geometry's materials to go through the transparent pipeline, which is slower to render and can exhibit rendering issues due to incorrect transparency sorting. However, unlike using a transparent material, setting [param transparency] to a value greater than [code]0.0[/code] (exclusive) will [i]not[/i] disable shadow rendering.
In spatial shaders, [code]1.0 - transparency[/code] is set as the default value of the [code]ALPHA[/code] built-in.
[b]Note:[/b] [param transparency] is clamped between [code]0.0[/code] and [code]1.0[/code], so this property cannot be used to make transparent materials more opaque than they originally are.
*/
//go:nosplit
func (self class) InstanceGeometrySetTransparency(instance gd.RID, transparency gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, transparency)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_geometry_set_transparency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a custom AABB to use when culling objects from the view frustum. Equivalent to setting [member GeometryInstance3D.custom_aabb].
*/
//go:nosplit
func (self class) InstanceSetCustomAabb(instance gd.RID, aabb gd.AABB)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, aabb)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_set_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Attaches a skeleton to an instance. Removes the previous skeleton from the instance.
*/
//go:nosplit
func (self class) InstanceAttachSkeleton(instance gd.RID, skeleton gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, skeleton)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_attach_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a margin to increase the size of the AABB when culling objects from the view frustum. This allows you to avoid culling objects that fall outside the view frustum. Equivalent to [member GeometryInstance3D.extra_cull_margin].
*/
//go:nosplit
func (self class) InstanceSetExtraVisibilityMargin(instance gd.RID, margin gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, margin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_set_extra_visibility_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the visibility parent for the given instance. Equivalent to [member Node3D.visibility_parent].
*/
//go:nosplit
func (self class) InstanceSetVisibilityParent(instance gd.RID, parent gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, parent)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_set_visibility_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], ignores both frustum and occlusion culling on the specified 3D geometry instance. This is not the same as [member GeometryInstance3D.ignore_occlusion_culling], which only ignores occlusion culling and leaves frustum culling intact.
*/
//go:nosplit
func (self class) InstanceSetIgnoreCulling(instance gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_set_ignore_culling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the flag for a given [enum InstanceFlags]. See [enum InstanceFlags] for more details.
*/
//go:nosplit
func (self class) InstanceGeometrySetFlag(instance gd.RID, flag classdb.RenderingServerInstanceFlags, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, flag)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_geometry_set_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the shadow casting setting to one of [enum ShadowCastingSetting]. Equivalent to [member GeometryInstance3D.cast_shadow].
*/
//go:nosplit
func (self class) InstanceGeometrySetCastShadowsSetting(instance gd.RID, shadow_casting_setting classdb.RenderingServerShadowCastingSetting)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, shadow_casting_setting)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_geometry_set_cast_shadows_setting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a material that will override the material for all surfaces on the mesh associated with this instance. Equivalent to [member GeometryInstance3D.material_override].
*/
//go:nosplit
func (self class) InstanceGeometrySetMaterialOverride(instance gd.RID, material gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, material)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_geometry_set_material_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a material that will be rendered for all surfaces on top of active materials for the mesh associated with this instance. Equivalent to [member GeometryInstance3D.material_overlay].
*/
//go:nosplit
func (self class) InstanceGeometrySetMaterialOverlay(instance gd.RID, material gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, material)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_geometry_set_material_overlay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the visibility range values for the given geometry instance. Equivalent to [member GeometryInstance3D.visibility_range_begin] and related properties.
*/
//go:nosplit
func (self class) InstanceGeometrySetVisibilityRange(instance gd.RID, min gd.Float, max gd.Float, min_margin gd.Float, max_margin gd.Float, fade_mode classdb.RenderingServerVisibilityRangeFadeMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, min)
	callframe.Arg(frame, max)
	callframe.Arg(frame, min_margin)
	callframe.Arg(frame, max_margin)
	callframe.Arg(frame, fade_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_geometry_set_visibility_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the lightmap GI instance to use for the specified 3D geometry instance. The lightmap UV scale for the specified instance (equivalent to [member GeometryInstance3D.gi_lightmap_scale]) and lightmap atlas slice must also be specified.
*/
//go:nosplit
func (self class) InstanceGeometrySetLightmap(instance gd.RID, lightmap gd.RID, lightmap_uv_scale gd.Rect2, lightmap_slice gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, lightmap)
	callframe.Arg(frame, lightmap_uv_scale)
	callframe.Arg(frame, lightmap_slice)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_geometry_set_lightmap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the level of detail bias to use when rendering the specified 3D geometry instance. Higher values result in higher detail from further away. Equivalent to [member GeometryInstance3D.lod_bias].
*/
//go:nosplit
func (self class) InstanceGeometrySetLodBias(instance gd.RID, lod_bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, lod_bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_geometry_set_lod_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the per-instance shader uniform on the specified 3D geometry instance. Equivalent to [method GeometryInstance3D.set_instance_shader_parameter].
*/
//go:nosplit
func (self class) InstanceGeometrySetShaderParameter(instance gd.RID, parameter gd.StringName, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, mmm.Get(parameter))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_geometry_set_shader_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the per-instance shader uniform from the specified 3D geometry instance. Equivalent to [method GeometryInstance3D.get_instance_shader_parameter].
[b]Note:[/b] Per-instance shader parameter names are case-sensitive.
*/
//go:nosplit
func (self class) InstanceGeometryGetShaderParameter(ctx gd.Lifetime, instance gd.RID, parameter gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, mmm.Get(parameter))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_geometry_get_shader_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the default value of the per-instance shader uniform from the specified 3D geometry instance. Equivalent to [method GeometryInstance3D.get_instance_shader_parameter].
*/
//go:nosplit
func (self class) InstanceGeometryGetShaderParameterDefaultValue(ctx gd.Lifetime, instance gd.RID, parameter gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, mmm.Get(parameter))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_geometry_get_shader_parameter_default_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a dictionary of per-instance shader uniform names of the per-instance shader uniform from the specified 3D geometry instance. The returned dictionary is in PropertyInfo format, with the keys [code]name[/code], [code]class_name[/code], [code]type[/code], [code]hint[/code], [code]hint_string[/code] and [code]usage[/code]. Equivalent to [method GeometryInstance3D.get_instance_shader_parameter].
*/
//go:nosplit
func (self class) InstanceGeometryGetShaderParameterList(ctx gd.Lifetime, instance gd.RID) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instance_geometry_get_shader_parameter_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Returns an array of object IDs intersecting with the provided AABB. Only 3D nodes that inherit from [VisualInstance3D] are considered, such as [MeshInstance3D] or [DirectionalLight3D]. Use [method @GlobalScope.instance_from_id] to obtain the actual nodes. A scenario RID must be provided, which is available in the [World3D] you want to query. This forces an update for all resources queued to update.
[b]Warning:[/b] This function is primarily intended for editor usage. For in-game use cases, prefer physics collision.
*/
//go:nosplit
func (self class) InstancesCullAabb(ctx gd.Lifetime, aabb gd.AABB, scenario gd.RID) gd.PackedInt64Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, aabb)
	callframe.Arg(frame, scenario)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instances_cull_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt64Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array of object IDs intersecting with the provided 3D ray. Only 3D nodes that inherit from [VisualInstance3D] are considered, such as [MeshInstance3D] or [DirectionalLight3D]. Use [method @GlobalScope.instance_from_id] to obtain the actual nodes. A scenario RID must be provided, which is available in the [World3D] you want to query. This forces an update for all resources queued to update.
[b]Warning:[/b] This function is primarily intended for editor usage. For in-game use cases, prefer physics collision.
*/
//go:nosplit
func (self class) InstancesCullRay(ctx gd.Lifetime, from gd.Vector3, to gd.Vector3, scenario gd.RID) gd.PackedInt64Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	callframe.Arg(frame, scenario)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instances_cull_ray, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt64Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array of object IDs intersecting with the provided convex shape. Only 3D nodes that inherit from [VisualInstance3D] are considered, such as [MeshInstance3D] or [DirectionalLight3D]. Use [method @GlobalScope.instance_from_id] to obtain the actual nodes. A scenario RID must be provided, which is available in the [World3D] you want to query. This forces an update for all resources queued to update.
[b]Warning:[/b] This function is primarily intended for editor usage. For in-game use cases, prefer physics collision.
*/
//go:nosplit
func (self class) InstancesCullConvex(ctx gd.Lifetime, convex gd.ArrayOf[gd.Plane], scenario gd.RID) gd.PackedInt64Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(convex))
	callframe.Arg(frame, scenario)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_instances_cull_convex, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt64Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Bakes the material data of the Mesh passed in the [param base] parameter with optional [param material_overrides] to a set of [Image]s of size [param image_size]. Returns an array of [Image]s containing material properties as specified in [enum BakeChannels].
*/
//go:nosplit
func (self class) BakeRenderUv2(ctx gd.Lifetime, base gd.RID, material_overrides gd.ArrayOf[gd.RID], image_size gd.Vector2i) gd.ArrayOf[object.Image] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, base)
	callframe.Arg(frame, mmm.Get(material_overrides))
	callframe.Arg(frame, image_size)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_bake_render_uv2, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.Image](ret)
}
/*
Creates a canvas and returns the assigned [RID]. It can be accessed with the RID that is returned. This RID will be used in all [code]canvas_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
Canvas has no [Resource] or [Node] equivalent.
*/
//go:nosplit
func (self class) CanvasCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
A copy of the canvas item will be drawn with a local offset of the mirroring [Vector2].
*/
//go:nosplit
func (self class) CanvasSetItemMirroring(canvas gd.RID, item gd.RID, mirroring gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, item)
	callframe.Arg(frame, mirroring)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_set_item_mirroring, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
A copy of the canvas item will be drawn with a local offset of the [param repeat_size] by the number of times of the [param repeat_times]. As the [param repeat_times] increases, the copies will spread away from the origin texture.
*/
//go:nosplit
func (self class) CanvasSetItemRepeat(item gd.RID, repeat_size gd.Vector2, repeat_times gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, repeat_size)
	callframe.Arg(frame, repeat_times)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_set_item_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Modulates all colors in the given canvas.
*/
//go:nosplit
func (self class) CanvasSetModulate(canvas gd.RID, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_set_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) CanvasSetDisableScale(disable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_set_disable_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a canvas texture and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]canvas_texture_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method. See also [method texture_2d_create].
[b]Note:[/b] The equivalent resource is [CanvasTexture] and is only meant to be used in 2D rendering, not 3D.
*/
//go:nosplit
func (self class) CanvasTextureCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_texture_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [param channel]'s [param texture] for the canvas texture specified by the [param canvas_texture] RID. Equivalent to [member CanvasTexture.diffuse_texture], [member CanvasTexture.normal_texture] and [member CanvasTexture.specular_texture].
*/
//go:nosplit
func (self class) CanvasTextureSetChannel(canvas_texture gd.RID, channel classdb.RenderingServerCanvasTextureChannel, texture gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas_texture)
	callframe.Arg(frame, channel)
	callframe.Arg(frame, texture)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_texture_set_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [param base_color] and [param shininess] to use for the canvas texture specified by the [param canvas_texture] RID. Equivalent to [member CanvasTexture.specular_color] and [member CanvasTexture.specular_shininess].
*/
//go:nosplit
func (self class) CanvasTextureSetShadingParameters(canvas_texture gd.RID, base_color gd.Color, shininess gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas_texture)
	callframe.Arg(frame, base_color)
	callframe.Arg(frame, shininess)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_texture_set_shading_parameters, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the texture [param filter] mode to use for the canvas texture specified by the [param canvas_texture] RID.
*/
//go:nosplit
func (self class) CanvasTextureSetTextureFilter(canvas_texture gd.RID, filter classdb.RenderingServerCanvasItemTextureFilter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas_texture)
	callframe.Arg(frame, filter)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_texture_set_texture_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the texture [param repeat] mode to use for the canvas texture specified by the [param canvas_texture] RID.
*/
//go:nosplit
func (self class) CanvasTextureSetTextureRepeat(canvas_texture gd.RID, repeat classdb.RenderingServerCanvasItemTextureRepeat)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas_texture)
	callframe.Arg(frame, repeat)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_texture_set_texture_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a new CanvasItem instance and returns its [RID]. It can be accessed with the RID that is returned. This RID will be used in all [code]canvas_item_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
[b]Note:[/b] The equivalent node is [CanvasItem].
*/
//go:nosplit
func (self class) CanvasItemCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a parent [CanvasItem] to the [CanvasItem]. The item will inherit transform, modulation and visibility from its parent, like [CanvasItem] nodes in the scene tree.
*/
//go:nosplit
func (self class) CanvasItemSetParent(item gd.RID, parent gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, parent)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the default texture filter mode for the canvas item specified by the [param item] RID. Equivalent to [member CanvasItem.texture_filter].
*/
//go:nosplit
func (self class) CanvasItemSetDefaultTextureFilter(item gd.RID, filter classdb.RenderingServerCanvasItemTextureFilter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, filter)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_default_texture_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the default texture repeat mode for the canvas item specified by the [param item] RID. Equivalent to [member CanvasItem.texture_repeat].
*/
//go:nosplit
func (self class) CanvasItemSetDefaultTextureRepeat(item gd.RID, repeat classdb.RenderingServerCanvasItemTextureRepeat)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, repeat)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_default_texture_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the visibility of the [CanvasItem].
*/
//go:nosplit
func (self class) CanvasItemSetVisible(item gd.RID, visible bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the light [param mask] for the canvas item specified by the [param item] RID. Equivalent to [member CanvasItem.light_mask].
*/
//go:nosplit
func (self class) CanvasItemSetLightMask(item gd.RID, mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_light_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the rendering visibility layer associated with this [CanvasItem]. Only [Viewport] nodes with a matching rendering mask will render this [CanvasItem].
*/
//go:nosplit
func (self class) CanvasItemSetVisibilityLayer(item gd.RID, visibility_layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, visibility_layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_visibility_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [param transform] of the canvas item specified by the [param item] RID. This affects where and how the item will be drawn. Child canvas items' transforms are multiplied by their parent's transform. Equivalent to [member Node2D.transform].
*/
//go:nosplit
func (self class) CanvasItemSetTransform(item gd.RID, transform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [param clip] is [code]true[/code], makes the canvas item specified by the [param item] RID not draw anything outside of its rect's coordinates. This clipping is fast, but works only with axis-aligned rectangles. This means that rotation is ignored by the clipping rectangle. For more advanced clipping shapes, use [method canvas_item_set_canvas_group_mode] instead.
[b]Note:[/b] The equivalent node functionality is found in [member Label.clip_text], [RichTextLabel] (always enabled) and more.
*/
//go:nosplit
func (self class) CanvasItemSetClip(item gd.RID, clip bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, clip)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_clip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [param enabled] is [code]true[/code], enables multichannel signed distance field rendering mode for the canvas item specified by the [param item] RID. This is meant to be used for font rendering, or with specially generated images using [url=https://github.com/Chlumsky/msdfgen]msdfgen[/url].
*/
//go:nosplit
func (self class) CanvasItemSetDistanceFieldMode(item gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_distance_field_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [param use_custom_rect] is [code]true[/code], sets the custom visibility rectangle (used for culling) to [param rect] for the canvas item specified by [param item]. Setting a custom visibility rect can reduce CPU load when drawing lots of 2D instances. If [param use_custom_rect] is [code]false[/code], automatically computes a visibility rectangle based on the canvas item's draw commands.
*/
//go:nosplit
func (self class) CanvasItemSetCustomRect(item gd.RID, use_custom_rect bool, rect gd.Rect2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, use_custom_rect)
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_custom_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Multiplies the color of the canvas item specified by the [param item] RID, while affecting its children. See also [method canvas_item_set_self_modulate]. Equivalent to [member CanvasItem.modulate].
*/
//go:nosplit
func (self class) CanvasItemSetModulate(item gd.RID, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Multiplies the color of the canvas item specified by the [param item] RID, without affecting its children. See also [method canvas_item_set_modulate]. Equivalent to [member CanvasItem.self_modulate].
*/
//go:nosplit
func (self class) CanvasItemSetSelfModulate(item gd.RID, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_self_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [param enabled] is [code]true[/code], draws the canvas item specified by the [param item] RID behind its parent. Equivalent to [member CanvasItem.show_behind_parent].
*/
//go:nosplit
func (self class) CanvasItemSetDrawBehindParent(item gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_draw_behind_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [param interpolated] is [code]true[/code], turns on physics interpolation for the canvas item.
*/
//go:nosplit
func (self class) CanvasItemSetInterpolated(item gd.RID, interpolated bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, interpolated)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_interpolated, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Prevents physics interpolation for the current physics tick.
This is useful when moving a canvas item to a new location, to give an instantaneous change rather than interpolation from the previous location.
*/
//go:nosplit
func (self class) CanvasItemResetPhysicsInterpolation(item gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_reset_physics_interpolation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Transforms both the current and previous stored transform for a canvas item.
This allows transforming a canvas item without creating a "glitch" in the interpolation, which is particularly useful for large worlds utilizing a shifting origin.
*/
//go:nosplit
func (self class) CanvasItemTransformPhysicsInterpolation(item gd.RID, transform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_transform_physics_interpolation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a line on the [CanvasItem] pointed to by the [param item] [RID]. See also [method CanvasItem.draw_line].
*/
//go:nosplit
func (self class) CanvasItemAddLine(item gd.RID, from gd.Vector2, to gd.Vector2, color gd.Color, width gd.Float, antialiased bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	callframe.Arg(frame, color)
	callframe.Arg(frame, width)
	callframe.Arg(frame, antialiased)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_add_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a 2D polyline on the [CanvasItem] pointed to by the [param item] [RID]. See also [method CanvasItem.draw_polyline] and [method CanvasItem.draw_polyline_colors].
*/
//go:nosplit
func (self class) CanvasItemAddPolyline(item gd.RID, points gd.PackedVector2Array, colors gd.PackedColorArray, width gd.Float, antialiased bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, mmm.Get(points))
	callframe.Arg(frame, mmm.Get(colors))
	callframe.Arg(frame, width)
	callframe.Arg(frame, antialiased)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_add_polyline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a 2D multiline on the [CanvasItem] pointed to by the [param item] [RID]. See also [method CanvasItem.draw_multiline] and [method CanvasItem.draw_multiline_colors].
*/
//go:nosplit
func (self class) CanvasItemAddMultiline(item gd.RID, points gd.PackedVector2Array, colors gd.PackedColorArray, width gd.Float, antialiased bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, mmm.Get(points))
	callframe.Arg(frame, mmm.Get(colors))
	callframe.Arg(frame, width)
	callframe.Arg(frame, antialiased)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_add_multiline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a rectangle on the [CanvasItem] pointed to by the [param item] [RID]. See also [method CanvasItem.draw_rect].
*/
//go:nosplit
func (self class) CanvasItemAddRect(item gd.RID, rect gd.Rect2, color gd.Color, antialiased bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, rect)
	callframe.Arg(frame, color)
	callframe.Arg(frame, antialiased)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_add_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a circle on the [CanvasItem] pointed to by the [param item] [RID]. See also [method CanvasItem.draw_circle].
*/
//go:nosplit
func (self class) CanvasItemAddCircle(item gd.RID, pos gd.Vector2, radius gd.Float, color gd.Color, antialiased bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, radius)
	callframe.Arg(frame, color)
	callframe.Arg(frame, antialiased)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_add_circle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a 2D textured rectangle on the [CanvasItem] pointed to by the [param item] [RID]. See also [method CanvasItem.draw_texture_rect] and [method Texture2D.draw_rect].
*/
//go:nosplit
func (self class) CanvasItemAddTextureRect(item gd.RID, rect gd.Rect2, texture gd.RID, tile bool, modulate gd.Color, transpose bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, rect)
	callframe.Arg(frame, texture)
	callframe.Arg(frame, tile)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, transpose)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_add_texture_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
See also [method CanvasItem.draw_msdf_texture_rect_region].
*/
//go:nosplit
func (self class) CanvasItemAddMsdfTextureRectRegion(item gd.RID, rect gd.Rect2, texture gd.RID, src_rect gd.Rect2, modulate gd.Color, outline_size gd.Int, px_range gd.Float, scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, rect)
	callframe.Arg(frame, texture)
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, outline_size)
	callframe.Arg(frame, px_range)
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_add_msdf_texture_rect_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
See also [method CanvasItem.draw_lcd_texture_rect_region].
*/
//go:nosplit
func (self class) CanvasItemAddLcdTextureRectRegion(item gd.RID, rect gd.Rect2, texture gd.RID, src_rect gd.Rect2, modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, rect)
	callframe.Arg(frame, texture)
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_add_lcd_texture_rect_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws the specified region of a 2D textured rectangle on the [CanvasItem] pointed to by the [param item] [RID]. See also [method CanvasItem.draw_texture_rect_region] and [method Texture2D.draw_rect_region].
*/
//go:nosplit
func (self class) CanvasItemAddTextureRectRegion(item gd.RID, rect gd.Rect2, texture gd.RID, src_rect gd.Rect2, modulate gd.Color, transpose bool, clip_uv bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, rect)
	callframe.Arg(frame, texture)
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, transpose)
	callframe.Arg(frame, clip_uv)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_add_texture_rect_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a nine-patch rectangle on the [CanvasItem] pointed to by the [param item] [RID].
*/
//go:nosplit
func (self class) CanvasItemAddNinePatch(item gd.RID, rect gd.Rect2, source gd.Rect2, texture gd.RID, topleft gd.Vector2, bottomright gd.Vector2, x_axis_mode classdb.RenderingServerNinePatchAxisMode, y_axis_mode classdb.RenderingServerNinePatchAxisMode, draw_center bool, modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, rect)
	callframe.Arg(frame, source)
	callframe.Arg(frame, texture)
	callframe.Arg(frame, topleft)
	callframe.Arg(frame, bottomright)
	callframe.Arg(frame, x_axis_mode)
	callframe.Arg(frame, y_axis_mode)
	callframe.Arg(frame, draw_center)
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_add_nine_patch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a 2D primitive on the [CanvasItem] pointed to by the [param item] [RID]. See also [method CanvasItem.draw_primitive].
*/
//go:nosplit
func (self class) CanvasItemAddPrimitive(item gd.RID, points gd.PackedVector2Array, colors gd.PackedColorArray, uvs gd.PackedVector2Array, texture gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, mmm.Get(points))
	callframe.Arg(frame, mmm.Get(colors))
	callframe.Arg(frame, mmm.Get(uvs))
	callframe.Arg(frame, texture)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_add_primitive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a 2D polygon on the [CanvasItem] pointed to by the [param item] [RID]. If you need more flexibility (such as being able to use bones), use [method canvas_item_add_triangle_array] instead. See also [method CanvasItem.draw_polygon].
*/
//go:nosplit
func (self class) CanvasItemAddPolygon(item gd.RID, points gd.PackedVector2Array, colors gd.PackedColorArray, uvs gd.PackedVector2Array, texture gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, mmm.Get(points))
	callframe.Arg(frame, mmm.Get(colors))
	callframe.Arg(frame, mmm.Get(uvs))
	callframe.Arg(frame, texture)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_add_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a triangle array on the [CanvasItem] pointed to by the [param item] [RID]. This is internally used by [Line2D] and [StyleBoxFlat] for rendering. [method canvas_item_add_triangle_array] is highly flexible, but more complex to use than [method canvas_item_add_polygon].
[b]Note:[/b] [param count] is unused and can be left unspecified.
*/
//go:nosplit
func (self class) CanvasItemAddTriangleArray(item gd.RID, indices gd.PackedInt32Array, points gd.PackedVector2Array, colors gd.PackedColorArray, uvs gd.PackedVector2Array, bones gd.PackedInt32Array, weights gd.PackedFloat32Array, texture gd.RID, count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, mmm.Get(indices))
	callframe.Arg(frame, mmm.Get(points))
	callframe.Arg(frame, mmm.Get(colors))
	callframe.Arg(frame, mmm.Get(uvs))
	callframe.Arg(frame, mmm.Get(bones))
	callframe.Arg(frame, mmm.Get(weights))
	callframe.Arg(frame, texture)
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_add_triangle_array, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a mesh created with [method mesh_create] with given [param transform], [param modulate] color, and [param texture]. This is used internally by [MeshInstance2D].
*/
//go:nosplit
func (self class) CanvasItemAddMesh(item gd.RID, mesh gd.RID, transform gd.Transform2D, modulate gd.Color, texture gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, mesh)
	callframe.Arg(frame, transform)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, texture)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_add_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws a 2D [MultiMesh] on the [CanvasItem] pointed to by the [param item] [RID]. See also [method CanvasItem.draw_multimesh].
*/
//go:nosplit
func (self class) CanvasItemAddMultimesh(item gd.RID, mesh gd.RID, texture gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, mesh)
	callframe.Arg(frame, texture)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_add_multimesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws particles on the [CanvasItem] pointed to by the [param item] [RID].
*/
//go:nosplit
func (self class) CanvasItemAddParticles(item gd.RID, particles gd.RID, texture gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, particles)
	callframe.Arg(frame, texture)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_add_particles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a [Transform2D] that will be used to transform subsequent canvas item commands.
*/
//go:nosplit
func (self class) CanvasItemAddSetTransform(item gd.RID, transform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_add_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [param ignore] is [code]true[/code], ignore clipping on items drawn with this canvas item until this is called again with [param ignore] set to false.
*/
//go:nosplit
func (self class) CanvasItemAddClipIgnore(item gd.RID, ignore bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, ignore)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_add_clip_ignore, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Subsequent drawing commands will be ignored unless they fall within the specified animation slice. This is a faster way to implement animations that loop on background rather than redrawing constantly.
*/
//go:nosplit
func (self class) CanvasItemAddAnimationSlice(item gd.RID, animation_length gd.Float, slice_begin gd.Float, slice_end gd.Float, offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, animation_length)
	callframe.Arg(frame, slice_begin)
	callframe.Arg(frame, slice_end)
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_add_animation_slice, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [param enabled] is [code]true[/code], child nodes with the lowest Y position are drawn before those with a higher Y position. Y-sorting only affects children that inherit from the canvas item specified by the [param item] RID, not the canvas item itself. Equivalent to [member CanvasItem.y_sort_enabled].
*/
//go:nosplit
func (self class) CanvasItemSetSortChildrenByY(item gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_sort_children_by_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [CanvasItem]'s Z index, i.e. its draw order (lower indexes are drawn first).
*/
//go:nosplit
func (self class) CanvasItemSetZIndex(item gd.RID, z_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, z_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_z_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If this is enabled, the Z index of the parent will be added to the children's Z index.
*/
//go:nosplit
func (self class) CanvasItemSetZAsRelativeToParent(item gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_z_as_relative_to_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [CanvasItem] to copy a rect to the backbuffer.
*/
//go:nosplit
func (self class) CanvasItemSetCopyToBackbuffer(item gd.RID, enabled bool, rect gd.Rect2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, enabled)
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_copy_to_backbuffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears the [CanvasItem] and removes all commands in it.
*/
//go:nosplit
func (self class) CanvasItemClear(item gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the index for the [CanvasItem].
*/
//go:nosplit
func (self class) CanvasItemSetDrawIndex(item gd.RID, index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_draw_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a new [param material] to the canvas item specified by the [param item] RID. Equivalent to [member CanvasItem.material].
*/
//go:nosplit
func (self class) CanvasItemSetMaterial(item gd.RID, material gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, material)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets if the [CanvasItem] uses its parent's material.
*/
//go:nosplit
func (self class) CanvasItemSetUseParentMaterial(item gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_use_parent_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the given [CanvasItem] as visibility notifier. [param area] defines the area of detecting visibility. [param enter_callable] is called when the [CanvasItem] enters the screen, [param exit_callable] is called when the [CanvasItem] exits the screen. If [param enable] is [code]false[/code], the item will no longer function as notifier.
This method can be used to manually mimic [VisibleOnScreenNotifier2D].
*/
//go:nosplit
func (self class) CanvasItemSetVisibilityNotifier(item gd.RID, enable bool, area gd.Rect2, enter_callable gd.Callable, exit_callable gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, enable)
	callframe.Arg(frame, area)
	callframe.Arg(frame, mmm.Get(enter_callable))
	callframe.Arg(frame, mmm.Get(exit_callable))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_visibility_notifier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the canvas group mode used during 2D rendering for the canvas item specified by the [param item] RID. For faster but more limited clipping, use [method canvas_item_set_clip] instead.
[b]Note:[/b] The equivalent node functionality is found in [CanvasGroup] and [member CanvasItem.clip_children].
*/
//go:nosplit
func (self class) CanvasItemSetCanvasGroupMode(item gd.RID, mode classdb.RenderingServerCanvasGroupMode, clear_margin gd.Float, fit_empty bool, fit_margin gd.Float, blur_mipmaps bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	callframe.Arg(frame, mode)
	callframe.Arg(frame, clear_margin)
	callframe.Arg(frame, fit_empty)
	callframe.Arg(frame, fit_margin)
	callframe.Arg(frame, blur_mipmaps)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_item_set_canvas_group_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the bounding rectangle for a canvas item in local space, as calculated by the renderer. This bound is used internally for culling.
[b]Warning:[/b] This function is intended for debugging in the editor, and will pass through and return a zero [Rect2] in exported projects.
*/
//go:nosplit
func (self class) DebugCanvasItemGetRect(item gd.RID) gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, item)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_debug_canvas_item_get_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a canvas light and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]canvas_light_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
[b]Note:[/b] The equivalent node is [Light2D].
*/
//go:nosplit
func (self class) CanvasLightCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Attaches the canvas light to the canvas. Removes it from its previous canvas.
*/
//go:nosplit
func (self class) CanvasLightAttachToCanvas(light gd.RID, canvas gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, canvas)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_attach_to_canvas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Enables or disables a canvas light.
*/
//go:nosplit
func (self class) CanvasLightSetEnabled(light gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the scale factor of a [PointLight2D]'s texture. Equivalent to [member PointLight2D.texture_scale].
*/
//go:nosplit
func (self class) CanvasLightSetTextureScale(light gd.RID, scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_set_texture_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the canvas light's [Transform2D].
*/
//go:nosplit
func (self class) CanvasLightSetTransform(light gd.RID, transform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the texture to be used by a [PointLight2D]. Equivalent to [member PointLight2D.texture].
*/
//go:nosplit
func (self class) CanvasLightSetTexture(light gd.RID, texture gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, texture)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the offset of a [PointLight2D]'s texture. Equivalent to [member PointLight2D.offset].
*/
//go:nosplit
func (self class) CanvasLightSetTextureOffset(light gd.RID, offset gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_set_texture_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the color for a light.
*/
//go:nosplit
func (self class) CanvasLightSetColor(light gd.RID, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a canvas light's height.
*/
//go:nosplit
func (self class) CanvasLightSetHeight(light gd.RID, height gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_set_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a canvas light's energy.
*/
//go:nosplit
func (self class) CanvasLightSetEnergy(light gd.RID, energy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, energy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_set_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the Z range of objects that will be affected by this light. Equivalent to [member Light2D.range_z_min] and [member Light2D.range_z_max].
*/
//go:nosplit
func (self class) CanvasLightSetZRange(light gd.RID, min_z gd.Int, max_z gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, min_z)
	callframe.Arg(frame, max_z)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_set_z_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
The layer range that gets rendered with this light.
*/
//go:nosplit
func (self class) CanvasLightSetLayerRange(light gd.RID, min_layer gd.Int, max_layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, min_layer)
	callframe.Arg(frame, max_layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_set_layer_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
The light mask. See [LightOccluder2D] for more information on light masks.
*/
//go:nosplit
func (self class) CanvasLightSetItemCullMask(light gd.RID, mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_set_item_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
The binary mask used to determine which layers this canvas light's shadows affects. See [LightOccluder2D] for more information on light masks.
*/
//go:nosplit
func (self class) CanvasLightSetItemShadowCullMask(light gd.RID, mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_set_item_shadow_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
The mode of the light, see [enum CanvasLightMode] constants.
*/
//go:nosplit
func (self class) CanvasLightSetMode(light gd.RID, mode classdb.RenderingServerCanvasLightMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Enables or disables the canvas light's shadow.
*/
//go:nosplit
func (self class) CanvasLightSetShadowEnabled(light gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_set_shadow_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the canvas light's shadow's filter, see [enum CanvasLightShadowFilter] constants.
*/
//go:nosplit
func (self class) CanvasLightSetShadowFilter(light gd.RID, filter classdb.RenderingServerCanvasLightShadowFilter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, filter)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_set_shadow_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the color of the canvas light's shadow.
*/
//go:nosplit
func (self class) CanvasLightSetShadowColor(light gd.RID, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_set_shadow_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Smoothens the shadow. The lower, the smoother.
*/
//go:nosplit
func (self class) CanvasLightSetShadowSmooth(light gd.RID, smooth gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, smooth)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_set_shadow_smooth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the blend mode for the given canvas light. See [enum CanvasLightBlendMode] for options. Equivalent to [member Light2D.blend_mode].
*/
//go:nosplit
func (self class) CanvasLightSetBlendMode(light gd.RID, mode classdb.RenderingServerCanvasLightBlendMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_set_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [param interpolated] is [code]true[/code], turns on physics interpolation for the canvas light.
*/
//go:nosplit
func (self class) CanvasLightSetInterpolated(light gd.RID, interpolated bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, interpolated)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_set_interpolated, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Prevents physics interpolation for the current physics tick.
This is useful when moving a canvas item to a new location, to give an instantaneous change rather than interpolation from the previous location.
*/
//go:nosplit
func (self class) CanvasLightResetPhysicsInterpolation(light gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_reset_physics_interpolation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Transforms both the current and previous stored transform for a canvas light.
This allows transforming a light without creating a "glitch" in the interpolation, which is is particularly useful for large worlds utilizing a shifting origin.
*/
//go:nosplit
func (self class) CanvasLightTransformPhysicsInterpolation(light gd.RID, transform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_transform_physics_interpolation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a light occluder and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]canvas_light_occluder_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
[b]Note:[/b] The equivalent node is [LightOccluder2D].
*/
//go:nosplit
func (self class) CanvasLightOccluderCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_occluder_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Attaches a light occluder to the canvas. Removes it from its previous canvas.
*/
//go:nosplit
func (self class) CanvasLightOccluderAttachToCanvas(occluder gd.RID, canvas gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, occluder)
	callframe.Arg(frame, canvas)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_occluder_attach_to_canvas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Enables or disables light occluder.
*/
//go:nosplit
func (self class) CanvasLightOccluderSetEnabled(occluder gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, occluder)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_occluder_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a light occluder's polygon.
*/
//go:nosplit
func (self class) CanvasLightOccluderSetPolygon(occluder gd.RID, polygon gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, occluder)
	callframe.Arg(frame, polygon)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_occluder_set_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) CanvasLightOccluderSetAsSdfCollision(occluder gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, occluder)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_occluder_set_as_sdf_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a light occluder's [Transform2D].
*/
//go:nosplit
func (self class) CanvasLightOccluderSetTransform(occluder gd.RID, transform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, occluder)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_occluder_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
The light mask. See [LightOccluder2D] for more information on light masks.
*/
//go:nosplit
func (self class) CanvasLightOccluderSetLightMask(occluder gd.RID, mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, occluder)
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_occluder_set_light_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [param interpolated] is [code]true[/code], turns on physics interpolation for the light occluder.
*/
//go:nosplit
func (self class) CanvasLightOccluderSetInterpolated(occluder gd.RID, interpolated bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, occluder)
	callframe.Arg(frame, interpolated)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_occluder_set_interpolated, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Prevents physics interpolation for the current physics tick.
This is useful when moving an occluder to a new location, to give an instantaneous change rather than interpolation from the previous location.
*/
//go:nosplit
func (self class) CanvasLightOccluderResetPhysicsInterpolation(occluder gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, occluder)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_occluder_reset_physics_interpolation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Transforms both the current and previous stored transform for a light occluder.
This allows transforming an occluder without creating a "glitch" in the interpolation, which is particularly useful for large worlds utilizing a shifting origin.
*/
//go:nosplit
func (self class) CanvasLightOccluderTransformPhysicsInterpolation(occluder gd.RID, transform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, occluder)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_light_occluder_transform_physics_interpolation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a new light occluder polygon and adds it to the RenderingServer. It can be accessed with the RID that is returned. This RID will be used in all [code]canvas_occluder_polygon_*[/code] RenderingServer functions.
Once finished with your RID, you will want to free the RID using the RenderingServer's [method free_rid] method.
[b]Note:[/b] The equivalent resource is [OccluderPolygon2D].
*/
//go:nosplit
func (self class) CanvasOccluderPolygonCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_occluder_polygon_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the shape of the occluder polygon.
*/
//go:nosplit
func (self class) CanvasOccluderPolygonSetShape(occluder_polygon gd.RID, shape gd.PackedVector2Array, closed bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, occluder_polygon)
	callframe.Arg(frame, mmm.Get(shape))
	callframe.Arg(frame, closed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_occluder_polygon_set_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets an occluder polygons cull mode. See [enum CanvasOccluderPolygonCullMode] constants.
*/
//go:nosplit
func (self class) CanvasOccluderPolygonSetCullMode(occluder_polygon gd.RID, mode classdb.RenderingServerCanvasOccluderPolygonCullMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, occluder_polygon)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_occluder_polygon_set_cull_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [member ProjectSettings.rendering/2d/shadow_atlas/size] to use for [Light2D] shadow rendering (in pixels). The value is rounded up to the nearest power of 2.
*/
//go:nosplit
func (self class) CanvasSetShadowTextureSize(size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_canvas_set_shadow_texture_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a new global shader uniform.
[b]Note:[/b] Global shader parameter names are case-sensitive.
*/
//go:nosplit
func (self class) GlobalShaderParameterAdd(name gd.StringName, atype classdb.RenderingServerGlobalShaderParameterType, default_value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, atype)
	callframe.Arg(frame, mmm.Get(default_value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_global_shader_parameter_add, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the global shader uniform specified by [param name].
*/
//go:nosplit
func (self class) GlobalShaderParameterRemove(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_global_shader_parameter_remove, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the list of global shader uniform names.
[b]Note:[/b] [method global_shader_parameter_get] has a large performance penalty as the rendering thread needs to synchronize with the calling thread, which is slow. Do not use this method during gameplay to avoid stuttering. If you need to read values in a script after setting them, consider creating an autoload where you store the values you need to query at the same time you're setting them as global parameters.
*/
//go:nosplit
func (self class) GlobalShaderParameterGetList(ctx gd.Lifetime) gd.ArrayOf[gd.StringName] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_global_shader_parameter_get_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.StringName](ret)
}
/*
Sets the global shader uniform [param name] to [param value].
*/
//go:nosplit
func (self class) GlobalShaderParameterSet(name gd.StringName, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_global_shader_parameter_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Overrides the global shader uniform [param name] with [param value]. Equivalent to the [ShaderGlobalsOverride] node.
*/
//go:nosplit
func (self class) GlobalShaderParameterSetOverride(name gd.StringName, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_global_shader_parameter_set_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the global shader uniform specified by [param name].
[b]Note:[/b] [method global_shader_parameter_get] has a large performance penalty as the rendering thread needs to synchronize with the calling thread, which is slow. Do not use this method during gameplay to avoid stuttering. If you need to read values in a script after setting them, consider creating an autoload where you store the values you need to query at the same time you're setting them as global parameters.
*/
//go:nosplit
func (self class) GlobalShaderParameterGet(ctx gd.Lifetime, name gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_global_shader_parameter_get, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the type associated to the global shader uniform specified by [param name].
[b]Note:[/b] [method global_shader_parameter_get] has a large performance penalty as the rendering thread needs to synchronize with the calling thread, which is slow. Do not use this method during gameplay to avoid stuttering. If you need to read values in a script after setting them, consider creating an autoload where you store the values you need to query at the same time you're setting them as global parameters.
*/
//go:nosplit
func (self class) GlobalShaderParameterGetType(name gd.StringName) classdb.RenderingServerGlobalShaderParameterType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[classdb.RenderingServerGlobalShaderParameterType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_global_shader_parameter_get_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Tries to free an object in the RenderingServer. To avoid memory leaks, this should be called after using an object as memory management does not occur automatically when using RenderingServer directly.
*/
//go:nosplit
func (self class) FreeRid(rid gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_free_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Schedules a callback to the given callable after a frame has been drawn.
*/
//go:nosplit
func (self class) RequestFrameDrawnCallback(callable gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(callable))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_request_frame_drawn_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if changes have been made to the RenderingServer's data. [method force_draw] is usually called if this happens.
*/
//go:nosplit
func (self class) HasChanged() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_has_changed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a statistic about the rendering engine which can be used for performance profiling. See [enum RenderingServer.RenderingInfo] for a list of values that can be queried. See also [method viewport_get_render_info], which returns information specific to a viewport.
[b]Note:[/b] Only 3D rendering is currently taken into account by some of these values, such as the number of draw calls.
[b]Note:[/b] Rendering information is not available until at least 2 frames have been rendered by the engine. If rendering information is not available, [method get_rendering_info] returns [code]0[/code]. To print rendering information in [code]_ready()[/code] successfully, use the following:
[codeblock]
func _ready():
    for _i in 2:
        await get_tree().process_frame

    print(RenderingServer.get_rendering_info(RENDERING_INFO_TOTAL_DRAW_CALLS_IN_FRAME))
[/codeblock]
*/
//go:nosplit
func (self class) GetRenderingInfo(info classdb.RenderingServerRenderingInfo) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, info)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_get_rendering_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the name of the video adapter (e.g. "GeForce GTX 1080/PCIe/SSE2").
[b]Note:[/b] When running a headless or server binary, this function returns an empty string.
[b]Note:[/b] On the web platform, some browsers such as Firefox may report a different, fixed GPU name such as "GeForce GTX 980" (regardless of the user's actual GPU model). This is done to make fingerprinting more difficult.
*/
//go:nosplit
func (self class) GetVideoAdapterName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_get_video_adapter_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the vendor of the video adapter (e.g. "NVIDIA Corporation").
[b]Note:[/b] When running a headless or server binary, this function returns an empty string.
*/
//go:nosplit
func (self class) GetVideoAdapterVendor(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_get_video_adapter_vendor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the type of the video adapter. Since dedicated graphics cards from a given generation will [i]usually[/i] be significantly faster than integrated graphics made in the same generation, the device type can be used as a basis for automatic graphics settings adjustment. However, this is not always true, so make sure to provide users with a way to manually override graphics settings.
[b]Note:[/b] When using the OpenGL backend or when running in headless mode, this function always returns [constant RenderingDevice.DEVICE_TYPE_OTHER].
*/
//go:nosplit
func (self class) GetVideoAdapterType() classdb.RenderingDeviceDeviceType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceDeviceType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_get_video_adapter_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the version of the graphics video adapter [i]currently in use[/i] (e.g. "1.2.189" for Vulkan, "3.3.0 NVIDIA 510.60.02" for OpenGL). This version may be different from the actual latest version supported by the hardware, as Godot may not always request the latest version. See also [method OS.get_video_adapter_driver_info].
[b]Note:[/b] When running a headless or server binary, this function returns an empty string.
*/
//go:nosplit
func (self class) GetVideoAdapterApiVersion(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_get_video_adapter_api_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a mesh of a sphere with the given number of horizontal subdivisions, vertical subdivisions and radius. See also [method get_test_cube].
*/
//go:nosplit
func (self class) MakeSphereMesh(latitudes gd.Int, longitudes gd.Int, radius gd.Float) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, latitudes)
	callframe.Arg(frame, longitudes)
	callframe.Arg(frame, radius)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_make_sphere_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the RID of the test cube. This mesh will be created and returned on the first call to [method get_test_cube], then it will be cached for subsequent calls. See also [method make_sphere_mesh].
*/
//go:nosplit
func (self class) GetTestCube() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_get_test_cube, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the RID of a 256×256 texture with a testing pattern on it (in [constant Image.FORMAT_RGB8] format). This texture will be created and returned on the first call to [method get_test_texture], then it will be cached for subsequent calls. See also [method get_white_texture].
Example of getting the test texture and applying it to a [Sprite2D] node:
[codeblock]
var texture_rid = RenderingServer.get_test_texture()
var texture = ImageTexture.create_from_image(RenderingServer.texture_2d_get(texture_rid))
$Sprite2D.texture = texture
[/codeblock]
*/
//go:nosplit
func (self class) GetTestTexture() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_get_test_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the ID of a 4×4 white texture (in [constant Image.FORMAT_RGB8] format). This texture will be created and returned on the first call to [method get_white_texture], then it will be cached for subsequent calls. See also [method get_test_texture].
Example of getting the white texture and applying it to a [Sprite2D] node:
[codeblock]
var texture_rid = RenderingServer.get_white_texture()
var texture = ImageTexture.create_from_image(RenderingServer.texture_2d_get(texture_rid))
$Sprite2D.texture = texture
[/codeblock]
*/
//go:nosplit
func (self class) GetWhiteTexture() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_get_white_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a boot image. The color defines the background color. If [param scale] is [code]true[/code], the image will be scaled to fit the screen size. If [param use_filter] is [code]true[/code], the image will be scaled with linear interpolation. If [param use_filter] is [code]false[/code], the image will be scaled with nearest-neighbor interpolation.
*/
//go:nosplit
func (self class) SetBootImage(image object.Image, color gd.Color, scale bool, use_filter bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(image[0].AsPointer())[0])
	callframe.Arg(frame, color)
	callframe.Arg(frame, scale)
	callframe.Arg(frame, use_filter)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_set_boot_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the default clear color which is used when a specific clear color has not been selected. See also [method set_default_clear_color].
*/
//go:nosplit
func (self class) GetDefaultClearColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_get_default_clear_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the default clear color which is used when a specific clear color has not been selected. See also [method get_default_clear_color].
*/
//go:nosplit
func (self class) SetDefaultClearColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_set_default_clear_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the OS supports a certain [param feature]. Features might be [code]s3tc[/code], [code]etc[/code], and [code]etc2[/code].
*/
//go:nosplit
func (self class) HasOsFeature(feature gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(feature))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_has_os_feature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
This method is currently unimplemented and does nothing if called with [param generate] set to [code]true[/code].
*/
//go:nosplit
func (self class) SetDebugGenerateWireframes(generate bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, generate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_set_debug_generate_wireframes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsRenderLoopEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_is_render_loop_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRenderLoopEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_set_render_loop_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the time taken to setup rendering on the CPU in milliseconds. This value is shared across all viewports and does [i]not[/i] require [method viewport_set_measure_render_time] to be enabled on a viewport to be queried. See also [method viewport_get_measured_render_time_cpu].
*/
//go:nosplit
func (self class) GetFrameSetupTimeCpu() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_get_frame_setup_time_cpu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Forces a synchronization between the CPU and GPU, which may be required in certain cases. Only call this when needed, as CPU-GPU synchronization has a performance cost.
*/
//go:nosplit
func (self class) ForceSync()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_force_sync, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Forces redrawing of all viewports at once. Must be called from the main thread.
*/
//go:nosplit
func (self class) ForceDraw(swap_buffers bool, frame_step gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, swap_buffers)
	callframe.Arg(frame, frame_step)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_force_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the global RenderingDevice.
[b]Note:[/b] When using the OpenGL backend or when running in headless mode, this function always returns [code]null[/code].
*/
//go:nosplit
func (self class) GetRenderingDevice(ctx gd.Lifetime) object.RenderingDevice {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_get_rendering_device, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.RenderingDevice
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a RenderingDevice that can be used to do draw and compute operations on a separate thread. Cannot draw to the screen nor share data with the global RenderingDevice.
[b]Note:[/b] When using the OpenGL backend or when running in headless mode, this function always returns [code]null[/code].
*/
//go:nosplit
func (self class) CreateLocalRenderingDevice(ctx gd.Lifetime) object.RenderingDevice {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_create_local_rendering_device, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.RenderingDevice
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if our code is currently executing on the rendering thread.
*/
//go:nosplit
func (self class) IsOnRenderThread() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_is_on_render_thread, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
As the RenderingServer actual logic may run on an separate thread, accessing its internals from the main (or any other) thread will result in errors. To make it easier to run code that can safely access the rendering internals (such as [RenderingDevice] and similar RD classes), push a callable via this function so it will be executed on the render thread.
*/
//go:nosplit
func (self class) CallOnRenderThread(callable gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(callable))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_call_on_render_thread, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
This method does nothing and always returns [code]false[/code].
*/
//go:nosplit
func (self class) HasFeature(feature classdb.RenderingServerFeatures) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RenderingServer.Bind_has_feature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsRenderingServer() Expert { return self[0].AsRenderingServer() }


//go:nosplit
func (self Simple) AsRenderingServer() Simple { return self[0].AsRenderingServer() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("RenderingServer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type TextureLayeredType = classdb.RenderingServerTextureLayeredType

const (
/*Array of 2-dimensional textures (see [Texture2DArray]).*/
	TextureLayered2dArray TextureLayeredType = 0
/*Cubemap texture (see [Cubemap]).*/
	TextureLayeredCubemap TextureLayeredType = 1
/*Array of cubemap textures (see [CubemapArray]).*/
	TextureLayeredCubemapArray TextureLayeredType = 2
)
type CubeMapLayer = classdb.RenderingServerCubeMapLayer

const (
/*Left face of a [Cubemap].*/
	CubemapLayerLeft CubeMapLayer = 0
/*Right face of a [Cubemap].*/
	CubemapLayerRight CubeMapLayer = 1
/*Bottom face of a [Cubemap].*/
	CubemapLayerBottom CubeMapLayer = 2
/*Top face of a [Cubemap].*/
	CubemapLayerTop CubeMapLayer = 3
/*Front face of a [Cubemap].*/
	CubemapLayerFront CubeMapLayer = 4
/*Back face of a [Cubemap].*/
	CubemapLayerBack CubeMapLayer = 5
)
type ShaderMode = classdb.RenderingServerShaderMode

const (
/*Shader is a 3D shader.*/
	ShaderSpatial ShaderMode = 0
/*Shader is a 2D shader.*/
	ShaderCanvasItem ShaderMode = 1
/*Shader is a particle shader (can be used in both 2D and 3D).*/
	ShaderParticles ShaderMode = 2
/*Shader is a 3D sky shader.*/
	ShaderSky ShaderMode = 3
/*Shader is a 3D fog shader.*/
	ShaderFog ShaderMode = 4
/*Represents the size of the [enum ShaderMode] enum.*/
	ShaderMax ShaderMode = 5
)
type ArrayType = classdb.RenderingServerArrayType

const (
/*Array is a vertex position array.*/
	ArrayVertex ArrayType = 0
/*Array is a normal array.*/
	ArrayNormal ArrayType = 1
/*Array is a tangent array.*/
	ArrayTangent ArrayType = 2
/*Array is a vertex color array.*/
	ArrayColor ArrayType = 3
/*Array is a UV coordinates array.*/
	ArrayTexUv ArrayType = 4
/*Array is a UV coordinates array for the second set of UV coordinates.*/
	ArrayTexUv2 ArrayType = 5
/*Array is a custom data array for the first set of custom data.*/
	ArrayCustom0 ArrayType = 6
/*Array is a custom data array for the second set of custom data.*/
	ArrayCustom1 ArrayType = 7
/*Array is a custom data array for the third set of custom data.*/
	ArrayCustom2 ArrayType = 8
/*Array is a custom data array for the fourth set of custom data.*/
	ArrayCustom3 ArrayType = 9
/*Array contains bone information.*/
	ArrayBones ArrayType = 10
/*Array is weight information.*/
	ArrayWeights ArrayType = 11
/*Array is an index array.*/
	ArrayIndex ArrayType = 12
/*Represents the size of the [enum ArrayType] enum.*/
	ArrayMax ArrayType = 13
)
type ArrayCustomFormat = classdb.RenderingServerArrayCustomFormat

const (
/*Custom data array contains 8-bit-per-channel red/green/blue/alpha color data. Values are normalized, unsigned floating-point in the [code][0.0, 1.0][/code] range.*/
	ArrayCustomRgba8Unorm ArrayCustomFormat = 0
/*Custom data array contains 8-bit-per-channel red/green/blue/alpha color data. Values are normalized, signed floating-point in the [code][-1.0, 1.0][/code] range.*/
	ArrayCustomRgba8Snorm ArrayCustomFormat = 1
/*Custom data array contains 16-bit-per-channel red/green color data. Values are floating-point in half precision.*/
	ArrayCustomRgHalf ArrayCustomFormat = 2
/*Custom data array contains 16-bit-per-channel red/green/blue/alpha color data. Values are floating-point in half precision.*/
	ArrayCustomRgbaHalf ArrayCustomFormat = 3
/*Custom data array contains 32-bit-per-channel red color data. Values are floating-point in single precision.*/
	ArrayCustomRFloat ArrayCustomFormat = 4
/*Custom data array contains 32-bit-per-channel red/green color data. Values are floating-point in single precision.*/
	ArrayCustomRgFloat ArrayCustomFormat = 5
/*Custom data array contains 32-bit-per-channel red/green/blue color data. Values are floating-point in single precision.*/
	ArrayCustomRgbFloat ArrayCustomFormat = 6
/*Custom data array contains 32-bit-per-channel red/green/blue/alpha color data. Values are floating-point in single precision.*/
	ArrayCustomRgbaFloat ArrayCustomFormat = 7
/*Represents the size of the [enum ArrayCustomFormat] enum.*/
	ArrayCustomMax ArrayCustomFormat = 8
)
type ArrayFormat = classdb.RenderingServerArrayFormat

const (
/*Flag used to mark a vertex position array.*/
	ArrayFormatVertex ArrayFormat = 1
/*Flag used to mark a normal array.*/
	ArrayFormatNormal ArrayFormat = 2
/*Flag used to mark a tangent array.*/
	ArrayFormatTangent ArrayFormat = 4
/*Flag used to mark a vertex color array.*/
	ArrayFormatColor ArrayFormat = 8
/*Flag used to mark a UV coordinates array.*/
	ArrayFormatTexUv ArrayFormat = 16
/*Flag used to mark a UV coordinates array for the second UV coordinates.*/
	ArrayFormatTexUv2 ArrayFormat = 32
/*Flag used to mark an array of custom per-vertex data for the first set of custom data.*/
	ArrayFormatCustom0 ArrayFormat = 64
/*Flag used to mark an array of custom per-vertex data for the second set of custom data.*/
	ArrayFormatCustom1 ArrayFormat = 128
/*Flag used to mark an array of custom per-vertex data for the third set of custom data.*/
	ArrayFormatCustom2 ArrayFormat = 256
/*Flag used to mark an array of custom per-vertex data for the fourth set of custom data.*/
	ArrayFormatCustom3 ArrayFormat = 512
/*Flag used to mark a bone information array.*/
	ArrayFormatBones ArrayFormat = 1024
/*Flag used to mark a weights array.*/
	ArrayFormatWeights ArrayFormat = 2048
/*Flag used to mark an index array.*/
	ArrayFormatIndex ArrayFormat = 4096
	ArrayFormatBlendShapeMask ArrayFormat = 7
	ArrayFormatCustomBase ArrayFormat = 13
	ArrayFormatCustomBits ArrayFormat = 3
	ArrayFormatCustom0Shift ArrayFormat = 13
	ArrayFormatCustom1Shift ArrayFormat = 16
	ArrayFormatCustom2Shift ArrayFormat = 19
	ArrayFormatCustom3Shift ArrayFormat = 22
	ArrayFormatCustomMask ArrayFormat = 7
	ArrayCompressFlagsBase ArrayFormat = 25
/*Flag used to mark that the array contains 2D vertices.*/
	ArrayFlagUse2dVertices ArrayFormat = 33554432
	ArrayFlagUseDynamicUpdate ArrayFormat = 67108864
/*Flag used to mark that the array uses 8 bone weights instead of 4.*/
	ArrayFlagUse8BoneWeights ArrayFormat = 134217728
/*Flag used to mark that the mesh does not have a vertex array and instead will infer vertex positions in the shader using indices and other information.*/
	ArrayFlagUsesEmptyVertexArray ArrayFormat = 268435456
/*Flag used to mark that a mesh is using compressed attributes (vertices, normals, tangents, UVs). When this form of compression is enabled, vertex positions will be packed into an RGBA16UNORM attribute and scaled in the vertex shader. The normal and tangent will be packed into an RG16UNORM representing an axis, and a 16-bit float stored in the A-channel of the vertex. UVs will use 16-bit normalized floats instead of full 32-bit signed floats. When using this compression mode you must use either vertices, normals, and tangents or only vertices. You cannot use normals without tangents. Importers will automatically enable this compression if they can.*/
	ArrayFlagCompressAttributes ArrayFormat = 536870912
/*Flag used to mark the start of the bits used to store the mesh version.*/
	ArrayFlagFormatVersionBase ArrayFormat = 35
/*Flag used to shift a mesh format int to bring the version into the lowest digits.*/
	ArrayFlagFormatVersionShift ArrayFormat = 35
/*Flag used to record the format used by prior mesh versions before the introduction of a version.*/
	ArrayFlagFormatVersion1 ArrayFormat = 0
/*Flag used to record the second iteration of the mesh version flag. The primary difference between this and [constant ARRAY_FLAG_FORMAT_VERSION_1] is that this version supports [constant ARRAY_FLAG_COMPRESS_ATTRIBUTES] and in this version vertex positions are de-interleaved from normals and tangents.*/
	ArrayFlagFormatVersion2 ArrayFormat = 34359738368
/*Flag used to record the current version that the engine expects. Currently this is the same as [constant ARRAY_FLAG_FORMAT_VERSION_2].*/
	ArrayFlagFormatCurrentVersion ArrayFormat = 34359738368
/*Flag used to isolate the bits used for mesh version after using [constant ARRAY_FLAG_FORMAT_VERSION_SHIFT] to shift them into place.*/
	ArrayFlagFormatVersionMask ArrayFormat = 255
)
type PrimitiveType = classdb.RenderingServerPrimitiveType

const (
/*Primitive to draw consists of points.*/
	PrimitivePoints PrimitiveType = 0
/*Primitive to draw consists of lines.*/
	PrimitiveLines PrimitiveType = 1
/*Primitive to draw consists of a line strip from start to end.*/
	PrimitiveLineStrip PrimitiveType = 2
/*Primitive to draw consists of triangles.*/
	PrimitiveTriangles PrimitiveType = 3
/*Primitive to draw consists of a triangle strip (the last 3 vertices are always combined to make a triangle).*/
	PrimitiveTriangleStrip PrimitiveType = 4
/*Represents the size of the [enum PrimitiveType] enum.*/
	PrimitiveMax PrimitiveType = 5
)
type BlendShapeMode = classdb.RenderingServerBlendShapeMode

const (
/*Blend shapes are normalized.*/
	BlendShapeModeNormalized BlendShapeMode = 0
/*Blend shapes are relative to base weight.*/
	BlendShapeModeRelative BlendShapeMode = 1
)
type MultimeshTransformFormat = classdb.RenderingServerMultimeshTransformFormat

const (
/*Use [Transform2D] to store MultiMesh transform.*/
	MultimeshTransform2d MultimeshTransformFormat = 0
/*Use [Transform3D] to store MultiMesh transform.*/
	MultimeshTransform3d MultimeshTransformFormat = 1
)
type LightProjectorFilter = classdb.RenderingServerLightProjectorFilter

const (
/*Nearest-neighbor filter for light projectors (use for pixel art light projectors). No mipmaps are used for rendering, which means light projectors at a distance will look sharp but grainy. This has roughly the same performance cost as using mipmaps.*/
	LightProjectorFilterNearest LightProjectorFilter = 0
/*Linear filter for light projectors (use for non-pixel art light projectors). No mipmaps are used for rendering, which means light projectors at a distance will look smooth but blurry. This has roughly the same performance cost as using mipmaps.*/
	LightProjectorFilterLinear LightProjectorFilter = 1
/*Nearest-neighbor filter for light projectors (use for pixel art light projectors). Isotropic mipmaps are used for rendering, which means light projectors at a distance will look smooth but blurry. This has roughly the same performance cost as not using mipmaps.*/
	LightProjectorFilterNearestMipmaps LightProjectorFilter = 2
/*Linear filter for light projectors (use for non-pixel art light projectors). Isotropic mipmaps are used for rendering, which means light projectors at a distance will look smooth but blurry. This has roughly the same performance cost as not using mipmaps.*/
	LightProjectorFilterLinearMipmaps LightProjectorFilter = 3
/*Nearest-neighbor filter for light projectors (use for pixel art light projectors). Anisotropic mipmaps are used for rendering, which means light projectors at a distance will look smooth and sharp when viewed from oblique angles. This looks better compared to isotropic mipmaps, but is slower. The level of anisotropic filtering is defined by [member ProjectSettings.rendering/textures/default_filters/anisotropic_filtering_level].*/
	LightProjectorFilterNearestMipmapsAnisotropic LightProjectorFilter = 4
/*Linear filter for light projectors (use for non-pixel art light projectors). Anisotropic mipmaps are used for rendering, which means light projectors at a distance will look smooth and sharp when viewed from oblique angles. This looks better compared to isotropic mipmaps, but is slower. The level of anisotropic filtering is defined by [member ProjectSettings.rendering/textures/default_filters/anisotropic_filtering_level].*/
	LightProjectorFilterLinearMipmapsAnisotropic LightProjectorFilter = 5
)
type LightType = classdb.RenderingServerLightType

const (
/*Directional (sun/moon) light (see [DirectionalLight3D]).*/
	LightDirectional LightType = 0
/*Omni light (see [OmniLight3D]).*/
	LightOmni LightType = 1
/*Spot light (see [SpotLight3D]).*/
	LightSpot LightType = 2
)
type LightParam = classdb.RenderingServerLightParam

const (
/*The light's energy multiplier.*/
	LightParamEnergy LightParam = 0
/*The light's indirect energy multiplier (final indirect energy is [constant LIGHT_PARAM_ENERGY] * [constant LIGHT_PARAM_INDIRECT_ENERGY]).*/
	LightParamIndirectEnergy LightParam = 1
/*The light's volumetric fog energy multiplier (final volumetric fog energy is [constant LIGHT_PARAM_ENERGY] * [constant LIGHT_PARAM_VOLUMETRIC_FOG_ENERGY]).*/
	LightParamVolumetricFogEnergy LightParam = 2
/*The light's influence on specularity.*/
	LightParamSpecular LightParam = 3
/*The light's range.*/
	LightParamRange LightParam = 4
/*The size of the light when using spot light or omni light. The angular size of the light when using directional light.*/
	LightParamSize LightParam = 5
/*The light's attenuation.*/
	LightParamAttenuation LightParam = 6
/*The spotlight's angle.*/
	LightParamSpotAngle LightParam = 7
/*The spotlight's attenuation.*/
	LightParamSpotAttenuation LightParam = 8
/*The maximum distance for shadow splits. Increasing this value will make directional shadows visible from further away, at the cost of lower overall shadow detail and performance (since more objects need to be included in the directional shadow rendering).*/
	LightParamShadowMaxDistance LightParam = 9
/*Proportion of shadow atlas occupied by the first split.*/
	LightParamShadowSplit1Offset LightParam = 10
/*Proportion of shadow atlas occupied by the second split.*/
	LightParamShadowSplit2Offset LightParam = 11
/*Proportion of shadow atlas occupied by the third split. The fourth split occupies the rest.*/
	LightParamShadowSplit3Offset LightParam = 12
/*Proportion of shadow max distance where the shadow will start to fade out.*/
	LightParamShadowFadeStart LightParam = 13
/*Normal bias used to offset shadow lookup by object normal. Can be used to fix self-shadowing artifacts.*/
	LightParamShadowNormalBias LightParam = 14
/*Bias the shadow lookup to fix self-shadowing artifacts.*/
	LightParamShadowBias LightParam = 15
/*Sets the size of the directional shadow pancake. The pancake offsets the start of the shadow's camera frustum to provide a higher effective depth resolution for the shadow. However, a high pancake size can cause artifacts in the shadows of large objects that are close to the edge of the frustum. Reducing the pancake size can help. Setting the size to [code]0[/code] turns off the pancaking effect.*/
	LightParamShadowPancakeSize LightParam = 16
/*The light's shadow opacity. Values lower than [code]1.0[/code] make the light appear through shadows. This can be used to fake global illumination at a low performance cost.*/
	LightParamShadowOpacity LightParam = 17
/*Blurs the edges of the shadow. Can be used to hide pixel artifacts in low resolution shadow maps. A high value can make shadows appear grainy and can cause other unwanted artifacts. Try to keep as near default as possible.*/
	LightParamShadowBlur LightParam = 18
	LightParamTransmittanceBias LightParam = 19
/*Constant representing the intensity of the light, measured in Lumens when dealing with a [SpotLight3D] or [OmniLight3D], or measured in Lux with a [DirectionalLight3D]. Only used when [member ProjectSettings.rendering/lights_and_shadows/use_physical_light_units] is [code]true[/code].*/
	LightParamIntensity LightParam = 20
/*Represents the size of the [enum LightParam] enum.*/
	LightParamMax LightParam = 21
)
type LightBakeMode = classdb.RenderingServerLightBakeMode

const (
/*Light is ignored when baking. This is the fastest mode, but the light will be taken into account when baking global illumination. This mode should generally be used for dynamic lights that change quickly, as the effect of global illumination is less noticeable on those lights.*/
	LightBakeDisabled LightBakeMode = 0
/*Light is taken into account in static baking ([VoxelGI], [LightmapGI], SDFGI ([member Environment.sdfgi_enabled])). The light can be moved around or modified, but its global illumination will not update in real-time. This is suitable for subtle changes (such as flickering torches), but generally not large changes such as toggling a light on and off.*/
	LightBakeStatic LightBakeMode = 1
/*Light is taken into account in dynamic baking ([VoxelGI] and SDFGI ([member Environment.sdfgi_enabled]) only). The light can be moved around or modified with global illumination updating in real-time. The light's global illumination appearance will be slightly different compared to [constant LIGHT_BAKE_STATIC]. This has a greater performance cost compared to [constant LIGHT_BAKE_STATIC]. When using SDFGI, the update speed of dynamic lights is affected by [member ProjectSettings.rendering/global_illumination/sdfgi/frames_to_update_lights].*/
	LightBakeDynamic LightBakeMode = 2
)
type LightOmniShadowMode = classdb.RenderingServerLightOmniShadowMode

const (
/*Use a dual paraboloid shadow map for omni lights.*/
	LightOmniShadowDualParaboloid LightOmniShadowMode = 0
/*Use a cubemap shadow map for omni lights. Slower but better quality than dual paraboloid.*/
	LightOmniShadowCube LightOmniShadowMode = 1
)
type LightDirectionalShadowMode = classdb.RenderingServerLightDirectionalShadowMode

const (
/*Use orthogonal shadow projection for directional light.*/
	LightDirectionalShadowOrthogonal LightDirectionalShadowMode = 0
/*Use 2 splits for shadow projection when using directional light.*/
	LightDirectionalShadowParallel2Splits LightDirectionalShadowMode = 1
/*Use 4 splits for shadow projection when using directional light.*/
	LightDirectionalShadowParallel4Splits LightDirectionalShadowMode = 2
)
type LightDirectionalSkyMode = classdb.RenderingServerLightDirectionalSkyMode

const (
/*Use DirectionalLight3D in both sky rendering and scene lighting.*/
	LightDirectionalSkyModeLightAndSky LightDirectionalSkyMode = 0
/*Only use DirectionalLight3D in scene lighting.*/
	LightDirectionalSkyModeLightOnly LightDirectionalSkyMode = 1
/*Only use DirectionalLight3D in sky rendering.*/
	LightDirectionalSkyModeSkyOnly LightDirectionalSkyMode = 2
)
type ShadowQuality = classdb.RenderingServerShadowQuality

const (
/*Lowest shadow filtering quality (fastest). Soft shadows are not available with this quality setting, which means the [member Light3D.shadow_blur] property is ignored if [member Light3D.light_size] and [member Light3D.light_angular_distance] is [code]0.0[/code].
[b]Note:[/b] The variable shadow blur performed by [member Light3D.light_size] and [member Light3D.light_angular_distance] is still effective when using hard shadow filtering. In this case, [member Light3D.shadow_blur] [i]is[/i] taken into account. However, the results will not be blurred, instead the blur amount is treated as a maximum radius for the penumbra.*/
	ShadowQualityHard ShadowQuality = 0
/*Very low shadow filtering quality (faster). When using this quality setting, [member Light3D.shadow_blur] is automatically multiplied by 0.75× to avoid introducing too much noise. This division only applies to lights whose [member Light3D.light_size] or [member Light3D.light_angular_distance] is [code]0.0[/code]).*/
	ShadowQualitySoftVeryLow ShadowQuality = 1
/*Low shadow filtering quality (fast).*/
	ShadowQualitySoftLow ShadowQuality = 2
/*Medium low shadow filtering quality (average).*/
	ShadowQualitySoftMedium ShadowQuality = 3
/*High low shadow filtering quality (slow). When using this quality setting, [member Light3D.shadow_blur] is automatically multiplied by 1.5× to better make use of the high sample count. This increased blur also improves the stability of dynamic object shadows. This multiplier only applies to lights whose [member Light3D.light_size] or [member Light3D.light_angular_distance] is [code]0.0[/code]).*/
	ShadowQualitySoftHigh ShadowQuality = 4
/*Highest low shadow filtering quality (slowest). When using this quality setting, [member Light3D.shadow_blur] is automatically multiplied by 2× to better make use of the high sample count. This increased blur also improves the stability of dynamic object shadows. This multiplier only applies to lights whose [member Light3D.light_size] or [member Light3D.light_angular_distance] is [code]0.0[/code]).*/
	ShadowQualitySoftUltra ShadowQuality = 5
/*Represents the size of the [enum ShadowQuality] enum.*/
	ShadowQualityMax ShadowQuality = 6
)
type ReflectionProbeUpdateMode = classdb.RenderingServerReflectionProbeUpdateMode

const (
/*Reflection probe will update reflections once and then stop.*/
	ReflectionProbeUpdateOnce ReflectionProbeUpdateMode = 0
/*Reflection probe will update each frame. This mode is necessary to capture moving objects.*/
	ReflectionProbeUpdateAlways ReflectionProbeUpdateMode = 1
)
type ReflectionProbeAmbientMode = classdb.RenderingServerReflectionProbeAmbientMode

const (
/*Do not apply any ambient lighting inside the reflection probe's box defined by its size.*/
	ReflectionProbeAmbientDisabled ReflectionProbeAmbientMode = 0
/*Apply automatically-sourced environment lighting inside the reflection probe's box defined by its size.*/
	ReflectionProbeAmbientEnvironment ReflectionProbeAmbientMode = 1
/*Apply custom ambient lighting inside the reflection probe's box defined by its size. See [method reflection_probe_set_ambient_color] and [method reflection_probe_set_ambient_energy].*/
	ReflectionProbeAmbientColor ReflectionProbeAmbientMode = 2
)
type DecalTexture = classdb.RenderingServerDecalTexture

const (
/*Albedo texture slot in a decal ([member Decal.texture_albedo]).*/
	DecalTextureAlbedo DecalTexture = 0
/*Normal map texture slot in a decal ([member Decal.texture_normal]).*/
	DecalTextureNormal DecalTexture = 1
/*Occlusion/Roughness/Metallic texture slot in a decal ([member Decal.texture_orm]).*/
	DecalTextureOrm DecalTexture = 2
/*Emission texture slot in a decal ([member Decal.texture_emission]).*/
	DecalTextureEmission DecalTexture = 3
/*Represents the size of the [enum DecalTexture] enum.*/
	DecalTextureMax DecalTexture = 4
)
type DecalFilter = classdb.RenderingServerDecalFilter

const (
/*Nearest-neighbor filter for decals (use for pixel art decals). No mipmaps are used for rendering, which means decals at a distance will look sharp but grainy. This has roughly the same performance cost as using mipmaps.*/
	DecalFilterNearest DecalFilter = 0
/*Linear filter for decals (use for non-pixel art decals). No mipmaps are used for rendering, which means decals at a distance will look smooth but blurry. This has roughly the same performance cost as using mipmaps.*/
	DecalFilterLinear DecalFilter = 1
/*Nearest-neighbor filter for decals (use for pixel art decals). Isotropic mipmaps are used for rendering, which means decals at a distance will look smooth but blurry. This has roughly the same performance cost as not using mipmaps.*/
	DecalFilterNearestMipmaps DecalFilter = 2
/*Linear filter for decals (use for non-pixel art decals). Isotropic mipmaps are used for rendering, which means decals at a distance will look smooth but blurry. This has roughly the same performance cost as not using mipmaps.*/
	DecalFilterLinearMipmaps DecalFilter = 3
/*Nearest-neighbor filter for decals (use for pixel art decals). Anisotropic mipmaps are used for rendering, which means decals at a distance will look smooth and sharp when viewed from oblique angles. This looks better compared to isotropic mipmaps, but is slower. The level of anisotropic filtering is defined by [member ProjectSettings.rendering/textures/default_filters/anisotropic_filtering_level].*/
	DecalFilterNearestMipmapsAnisotropic DecalFilter = 4
/*Linear filter for decals (use for non-pixel art decals). Anisotropic mipmaps are used for rendering, which means decals at a distance will look smooth and sharp when viewed from oblique angles. This looks better compared to isotropic mipmaps, but is slower. The level of anisotropic filtering is defined by [member ProjectSettings.rendering/textures/default_filters/anisotropic_filtering_level].*/
	DecalFilterLinearMipmapsAnisotropic DecalFilter = 5
)
type VoxelGIQuality = classdb.RenderingServerVoxelGIQuality

const (
/*Low [VoxelGI] rendering quality using 4 cones.*/
	VoxelGiQualityLow VoxelGIQuality = 0
/*High [VoxelGI] rendering quality using 6 cones.*/
	VoxelGiQualityHigh VoxelGIQuality = 1
)
type ParticlesMode = classdb.RenderingServerParticlesMode

const (
/*2D particles.*/
	ParticlesMode2d ParticlesMode = 0
/*3D particles.*/
	ParticlesMode3d ParticlesMode = 1
)
type ParticlesTransformAlign = classdb.RenderingServerParticlesTransformAlign

const (
	ParticlesTransformAlignDisabled ParticlesTransformAlign = 0
	ParticlesTransformAlignZBillboard ParticlesTransformAlign = 1
	ParticlesTransformAlignYToVelocity ParticlesTransformAlign = 2
	ParticlesTransformAlignZBillboardYToVelocity ParticlesTransformAlign = 3
)
type ParticlesDrawOrder = classdb.RenderingServerParticlesDrawOrder

const (
/*Draw particles in the order that they appear in the particles array.*/
	ParticlesDrawOrderIndex ParticlesDrawOrder = 0
/*Sort particles based on their lifetime. In other words, the particle with the highest lifetime is drawn at the front.*/
	ParticlesDrawOrderLifetime ParticlesDrawOrder = 1
/*Sort particles based on the inverse of their lifetime. In other words, the particle with the lowest lifetime is drawn at the front.*/
	ParticlesDrawOrderReverseLifetime ParticlesDrawOrder = 2
/*Sort particles based on their distance to the camera.*/
	ParticlesDrawOrderViewDepth ParticlesDrawOrder = 3
)
type ParticlesCollisionType = classdb.RenderingServerParticlesCollisionType

const (
	ParticlesCollisionTypeSphereAttract ParticlesCollisionType = 0
	ParticlesCollisionTypeBoxAttract ParticlesCollisionType = 1
	ParticlesCollisionTypeVectorFieldAttract ParticlesCollisionType = 2
	ParticlesCollisionTypeSphereCollide ParticlesCollisionType = 3
	ParticlesCollisionTypeBoxCollide ParticlesCollisionType = 4
	ParticlesCollisionTypeSdfCollide ParticlesCollisionType = 5
	ParticlesCollisionTypeHeightfieldCollide ParticlesCollisionType = 6
)
type ParticlesCollisionHeightfieldResolution = classdb.RenderingServerParticlesCollisionHeightfieldResolution

const (
	ParticlesCollisionHeightfieldResolution256 ParticlesCollisionHeightfieldResolution = 0
	ParticlesCollisionHeightfieldResolution512 ParticlesCollisionHeightfieldResolution = 1
	ParticlesCollisionHeightfieldResolution1024 ParticlesCollisionHeightfieldResolution = 2
	ParticlesCollisionHeightfieldResolution2048 ParticlesCollisionHeightfieldResolution = 3
	ParticlesCollisionHeightfieldResolution4096 ParticlesCollisionHeightfieldResolution = 4
	ParticlesCollisionHeightfieldResolution8192 ParticlesCollisionHeightfieldResolution = 5
/*Represents the size of the [enum ParticlesCollisionHeightfieldResolution] enum.*/
	ParticlesCollisionHeightfieldResolutionMax ParticlesCollisionHeightfieldResolution = 6
)
type FogVolumeShape = classdb.RenderingServerFogVolumeShape

const (
/*[FogVolume] will be shaped like an ellipsoid (stretched sphere).*/
	FogVolumeShapeEllipsoid FogVolumeShape = 0
/*[FogVolume] will be shaped like a cone pointing upwards (in local coordinates). The cone's angle is set automatically to fill the size. The cone will be adjusted to fit within the size. Rotate the [FogVolume] node to reorient the cone. Non-uniform scaling via size is not supported (scale the [FogVolume] node instead).*/
	FogVolumeShapeCone FogVolumeShape = 1
/*[FogVolume] will be shaped like an upright cylinder (in local coordinates). Rotate the [FogVolume] node to reorient the cylinder. The cylinder will be adjusted to fit within the size. Non-uniform scaling via size is not supported (scale the [FogVolume] node instead).*/
	FogVolumeShapeCylinder FogVolumeShape = 2
/*[FogVolume] will be shaped like a box.*/
	FogVolumeShapeBox FogVolumeShape = 3
/*[FogVolume] will have no shape, will cover the whole world and will not be culled.*/
	FogVolumeShapeWorld FogVolumeShape = 4
/*Represents the size of the [enum FogVolumeShape] enum.*/
	FogVolumeShapeMax FogVolumeShape = 5
)
type ViewportScaling3DMode = classdb.RenderingServerViewportScaling3DMode

const (
/*Use bilinear scaling for the viewport's 3D buffer. The amount of scaling can be set using [member Viewport.scaling_3d_scale]. Values less than [code]1.0[/code] will result in undersampling while values greater than [code]1.0[/code] will result in supersampling. A value of [code]1.0[/code] disables scaling.*/
	ViewportScaling3dModeBilinear ViewportScaling3DMode = 0
/*Use AMD FidelityFX Super Resolution 1.0 upscaling for the viewport's 3D buffer. The amount of scaling can be set using [member Viewport.scaling_3d_scale]. Values less than [code]1.0[/code] will be result in the viewport being upscaled using FSR. Values greater than [code]1.0[/code] are not supported and bilinear downsampling will be used instead. A value of [code]1.0[/code] disables scaling.*/
	ViewportScaling3dModeFsr ViewportScaling3DMode = 1
/*Use AMD FidelityFX Super Resolution 2.2 upscaling for the viewport's 3D buffer. The amount of scaling can be set using [member Viewport.scaling_3d_scale]. Values less than [code]1.0[/code] will be result in the viewport being upscaled using FSR2. Values greater than [code]1.0[/code] are not supported and bilinear downsampling will be used instead. A value of [code]1.0[/code] will use FSR2 at native resolution as a TAA solution.*/
	ViewportScaling3dModeFsr2 ViewportScaling3DMode = 2
/*Represents the size of the [enum ViewportScaling3DMode] enum.*/
	ViewportScaling3dModeMax ViewportScaling3DMode = 3
)
type ViewportUpdateMode = classdb.RenderingServerViewportUpdateMode

const (
/*Do not update the viewport's render target.*/
	ViewportUpdateDisabled ViewportUpdateMode = 0
/*Update the viewport's render target once, then switch to [constant VIEWPORT_UPDATE_DISABLED].*/
	ViewportUpdateOnce ViewportUpdateMode = 1
/*Update the viewport's render target only when it is visible. This is the default value.*/
	ViewportUpdateWhenVisible ViewportUpdateMode = 2
/*Update the viewport's render target only when its parent is visible.*/
	ViewportUpdateWhenParentVisible ViewportUpdateMode = 3
/*Always update the viewport's render target.*/
	ViewportUpdateAlways ViewportUpdateMode = 4
)
type ViewportClearMode = classdb.RenderingServerViewportClearMode

const (
/*Always clear the viewport's render target before drawing.*/
	ViewportClearAlways ViewportClearMode = 0
/*Never clear the viewport's render target.*/
	ViewportClearNever ViewportClearMode = 1
/*Clear the viewport's render target on the next frame, then switch to [constant VIEWPORT_CLEAR_NEVER].*/
	ViewportClearOnlyNextFrame ViewportClearMode = 2
)
type ViewportEnvironmentMode = classdb.RenderingServerViewportEnvironmentMode

const (
/*Disable rendering of 3D environment over 2D canvas.*/
	ViewportEnvironmentDisabled ViewportEnvironmentMode = 0
/*Enable rendering of 3D environment over 2D canvas.*/
	ViewportEnvironmentEnabled ViewportEnvironmentMode = 1
/*Inherit enable/disable value from parent. If the topmost parent is also set to [constant VIEWPORT_ENVIRONMENT_INHERIT], then this has the same behavior as [constant VIEWPORT_ENVIRONMENT_ENABLED].*/
	ViewportEnvironmentInherit ViewportEnvironmentMode = 2
/*Represents the size of the [enum ViewportEnvironmentMode] enum.*/
	ViewportEnvironmentMax ViewportEnvironmentMode = 3
)
type ViewportSDFOversize = classdb.RenderingServerViewportSDFOversize

const (
/*Do not oversize the 2D signed distance field. Occluders may disappear when touching the viewport's edges, and [GPUParticles3D] collision may stop working earlier than intended. This has the lowest GPU requirements.*/
	ViewportSdfOversize100Percent ViewportSDFOversize = 0
/*2D signed distance field covers 20% of the viewport's size outside the viewport on each side (top, right, bottom, left).*/
	ViewportSdfOversize120Percent ViewportSDFOversize = 1
/*2D signed distance field covers 50% of the viewport's size outside the viewport on each side (top, right, bottom, left).*/
	ViewportSdfOversize150Percent ViewportSDFOversize = 2
/*2D signed distance field covers 100% of the viewport's size outside the viewport on each side (top, right, bottom, left). This has the highest GPU requirements.*/
	ViewportSdfOversize200Percent ViewportSDFOversize = 3
/*Represents the size of the [enum ViewportSDFOversize] enum.*/
	ViewportSdfOversizeMax ViewportSDFOversize = 4
)
type ViewportSDFScale = classdb.RenderingServerViewportSDFScale

const (
/*Full resolution 2D signed distance field scale. This has the highest GPU requirements.*/
	ViewportSdfScale100Percent ViewportSDFScale = 0
/*Half resolution 2D signed distance field scale on each axis (25% of the viewport pixel count).*/
	ViewportSdfScale50Percent ViewportSDFScale = 1
/*Quarter resolution 2D signed distance field scale on each axis (6.25% of the viewport pixel count). This has the lowest GPU requirements.*/
	ViewportSdfScale25Percent ViewportSDFScale = 2
/*Represents the size of the [enum ViewportSDFScale] enum.*/
	ViewportSdfScaleMax ViewportSDFScale = 3
)
type ViewportMSAA = classdb.RenderingServerViewportMSAA

const (
/*Multisample antialiasing for 3D is disabled. This is the default value, and also the fastest setting.*/
	ViewportMsaaDisabled ViewportMSAA = 0
/*Multisample antialiasing uses 2 samples per pixel for 3D. This has a moderate impact on performance.*/
	ViewportMsaa2x ViewportMSAA = 1
/*Multisample antialiasing uses 4 samples per pixel for 3D. This has a high impact on performance.*/
	ViewportMsaa4x ViewportMSAA = 2
/*Multisample antialiasing uses 8 samples per pixel for 3D. This has a very high impact on performance. Likely unsupported on low-end and older hardware.*/
	ViewportMsaa8x ViewportMSAA = 3
/*Represents the size of the [enum ViewportMSAA] enum.*/
	ViewportMsaaMax ViewportMSAA = 4
)
type ViewportScreenSpaceAA = classdb.RenderingServerViewportScreenSpaceAA

const (
/*Do not perform any antialiasing in the full screen post-process.*/
	ViewportScreenSpaceAaDisabled ViewportScreenSpaceAA = 0
/*Use fast approximate antialiasing. FXAA is a popular screen-space antialiasing method, which is fast but will make the image look blurry, especially at lower resolutions. It can still work relatively well at large resolutions such as 1440p and 4K.*/
	ViewportScreenSpaceAaFxaa ViewportScreenSpaceAA = 1
/*Represents the size of the [enum ViewportScreenSpaceAA] enum.*/
	ViewportScreenSpaceAaMax ViewportScreenSpaceAA = 2
)
type ViewportOcclusionCullingBuildQuality = classdb.RenderingServerViewportOcclusionCullingBuildQuality

const (
/*Low occlusion culling BVH build quality (as defined by Embree). Results in the lowest CPU usage, but least effective culling.*/
	ViewportOcclusionBuildQualityLow ViewportOcclusionCullingBuildQuality = 0
/*Medium occlusion culling BVH build quality (as defined by Embree).*/
	ViewportOcclusionBuildQualityMedium ViewportOcclusionCullingBuildQuality = 1
/*High occlusion culling BVH build quality (as defined by Embree). Results in the highest CPU usage, but most effective culling.*/
	ViewportOcclusionBuildQualityHigh ViewportOcclusionCullingBuildQuality = 2
)
type ViewportRenderInfo = classdb.RenderingServerViewportRenderInfo

const (
/*Number of objects drawn in a single frame.*/
	ViewportRenderInfoObjectsInFrame ViewportRenderInfo = 0
/*Number of points, lines, or triangles drawn in a single frame.*/
	ViewportRenderInfoPrimitivesInFrame ViewportRenderInfo = 1
/*Number of draw calls during this frame.*/
	ViewportRenderInfoDrawCallsInFrame ViewportRenderInfo = 2
/*Represents the size of the [enum ViewportRenderInfo] enum.*/
	ViewportRenderInfoMax ViewportRenderInfo = 3
)
type ViewportRenderInfoType = classdb.RenderingServerViewportRenderInfoType

const (
/*Visible render pass (excluding shadows).*/
	ViewportRenderInfoTypeVisible ViewportRenderInfoType = 0
/*Shadow render pass. Objects will be rendered several times depending on the number of amounts of lights with shadows and the number of directional shadow splits.*/
	ViewportRenderInfoTypeShadow ViewportRenderInfoType = 1
/*Canvas item rendering. This includes all 2D rendering.*/
	ViewportRenderInfoTypeCanvas ViewportRenderInfoType = 2
/*Represents the size of the [enum ViewportRenderInfoType] enum.*/
	ViewportRenderInfoTypeMax ViewportRenderInfoType = 3
)
type ViewportDebugDraw = classdb.RenderingServerViewportDebugDraw

const (
/*Debug draw is disabled. Default setting.*/
	ViewportDebugDrawDisabled ViewportDebugDraw = 0
/*Objects are displayed without light information.*/
	ViewportDebugDrawUnshaded ViewportDebugDraw = 1
/*Objects are displayed with only light information.*/
	ViewportDebugDrawLighting ViewportDebugDraw = 2
/*Objects are displayed semi-transparent with additive blending so you can see where they are drawing over top of one another. A higher overdraw (represented by brighter colors) means you are wasting performance on drawing pixels that are being hidden behind others.
[b]Note:[/b] When using this debug draw mode, custom shaders will be ignored. This means vertex displacement won't be visible anymore.*/
	ViewportDebugDrawOverdraw ViewportDebugDraw = 3
/*Debug draw draws objects in wireframe.*/
	ViewportDebugDrawWireframe ViewportDebugDraw = 4
/*Normal buffer is drawn instead of regular scene so you can see the per-pixel normals that will be used by post-processing effects.*/
	ViewportDebugDrawNormalBuffer ViewportDebugDraw = 5
/*Objects are displayed with only the albedo value from [VoxelGI]s.*/
	ViewportDebugDrawVoxelGiAlbedo ViewportDebugDraw = 6
/*Objects are displayed with only the lighting value from [VoxelGI]s.*/
	ViewportDebugDrawVoxelGiLighting ViewportDebugDraw = 7
/*Objects are displayed with only the emission color from [VoxelGI]s.*/
	ViewportDebugDrawVoxelGiEmission ViewportDebugDraw = 8
/*Draws the shadow atlas that stores shadows from [OmniLight3D]s and [SpotLight3D]s in the upper left quadrant of the [Viewport].*/
	ViewportDebugDrawShadowAtlas ViewportDebugDraw = 9
/*Draws the shadow atlas that stores shadows from [DirectionalLight3D]s in the upper left quadrant of the [Viewport].
The slice of the camera frustum related to the shadow map cascade is superimposed to visualize coverage. The color of each slice matches the colors used for [constant VIEWPORT_DEBUG_DRAW_PSSM_SPLITS]. When shadow cascades are blended the overlap is taken into account when drawing the frustum slices.
The last cascade shows all frustum slices to illustrate the coverage of all slices.*/
	ViewportDebugDrawDirectionalShadowAtlas ViewportDebugDraw = 10
/*Draws the estimated scene luminance. This is a 1×1 texture that is generated when autoexposure is enabled to control the scene's exposure.*/
	ViewportDebugDrawSceneLuminance ViewportDebugDraw = 11
/*Draws the screen space ambient occlusion texture instead of the scene so that you can clearly see how it is affecting objects. In order for this display mode to work, you must have [member Environment.ssao_enabled] set in your [WorldEnvironment].*/
	ViewportDebugDrawSsao ViewportDebugDraw = 12
/*Draws the screen space indirect lighting texture instead of the scene so that you can clearly see how it is affecting objects. In order for this display mode to work, you must have [member Environment.ssil_enabled] set in your [WorldEnvironment].*/
	ViewportDebugDrawSsil ViewportDebugDraw = 13
/*Colors each PSSM split for the [DirectionalLight3D]s in the scene a different color so you can see where the splits are. In order they will be colored red, green, blue, yellow.*/
	ViewportDebugDrawPssmSplits ViewportDebugDraw = 14
/*Draws the decal atlas that stores decal textures from [Decal]s.*/
	ViewportDebugDrawDecalAtlas ViewportDebugDraw = 15
/*Draws SDFGI cascade data. This is the data structure that is used to bounce lighting against and create reflections.*/
	ViewportDebugDrawSdfgi ViewportDebugDraw = 16
/*Draws SDFGI probe data. This is the data structure that is used to give indirect lighting dynamic objects moving within the scene.*/
	ViewportDebugDrawSdfgiProbes ViewportDebugDraw = 17
/*Draws the global illumination buffer ([VoxelGI] or SDFGI).*/
	ViewportDebugDrawGiBuffer ViewportDebugDraw = 18
/*Disable mesh LOD. All meshes are drawn with full detail, which can be used to compare performance.*/
	ViewportDebugDrawDisableLod ViewportDebugDraw = 19
/*Draws the [OmniLight3D] cluster. Clustering determines where lights are positioned in screen-space, which allows the engine to only process these portions of the screen for lighting.*/
	ViewportDebugDrawClusterOmniLights ViewportDebugDraw = 20
/*Draws the [SpotLight3D] cluster. Clustering determines where lights are positioned in screen-space, which allows the engine to only process these portions of the screen for lighting.*/
	ViewportDebugDrawClusterSpotLights ViewportDebugDraw = 21
/*Draws the [Decal] cluster. Clustering determines where decals are positioned in screen-space, which allows the engine to only process these portions of the screen for decals.*/
	ViewportDebugDrawClusterDecals ViewportDebugDraw = 22
/*Draws the [ReflectionProbe] cluster. Clustering determines where reflection probes are positioned in screen-space, which allows the engine to only process these portions of the screen for reflection probes.*/
	ViewportDebugDrawClusterReflectionProbes ViewportDebugDraw = 23
/*Draws the occlusion culling buffer. This low-resolution occlusion culling buffer is rasterized on the CPU and is used to check whether instances are occluded by other objects.*/
	ViewportDebugDrawOccluders ViewportDebugDraw = 24
/*Draws the motion vectors buffer. This is used by temporal antialiasing to correct for motion that occurs during gameplay.*/
	ViewportDebugDrawMotionVectors ViewportDebugDraw = 25
/*Internal buffer is drawn instead of regular scene so you can see the per-pixel output that will be used by post-processing effects.*/
	ViewportDebugDrawInternalBuffer ViewportDebugDraw = 26
)
type ViewportVRSMode = classdb.RenderingServerViewportVRSMode

const (
/*Variable rate shading is disabled.*/
	ViewportVrsDisabled ViewportVRSMode = 0
/*Variable rate shading uses a texture. Note, for stereoscopic use a texture atlas with a texture for each view.*/
	ViewportVrsTexture ViewportVRSMode = 1
/*Variable rate shading texture is supplied by the primary [XRInterface]. Note that this may override the update mode.*/
	ViewportVrsXr ViewportVRSMode = 2
/*Represents the size of the [enum ViewportVRSMode] enum.*/
	ViewportVrsMax ViewportVRSMode = 3
)
type ViewportVRSUpdateMode = classdb.RenderingServerViewportVRSUpdateMode

const (
/*The input texture for variable rate shading will not be processed.*/
	ViewportVrsUpdateDisabled ViewportVRSUpdateMode = 0
/*The input texture for variable rate shading will be processed once.*/
	ViewportVrsUpdateOnce ViewportVRSUpdateMode = 1
/*The input texture for variable rate shading will be processed each frame.*/
	ViewportVrsUpdateAlways ViewportVRSUpdateMode = 2
/*Represents the size of the [enum ViewportVRSUpdateMode] enum.*/
	ViewportVrsUpdateMax ViewportVRSUpdateMode = 3
)
type SkyMode = classdb.RenderingServerSkyMode

const (
/*Automatically selects the appropriate process mode based on your sky shader. If your shader uses [code]TIME[/code] or [code]POSITION[/code], this will use [constant SKY_MODE_REALTIME]. If your shader uses any of the [code]LIGHT_*[/code] variables or any custom uniforms, this uses [constant SKY_MODE_INCREMENTAL]. Otherwise, this defaults to [constant SKY_MODE_QUALITY].*/
	SkyModeAutomatic SkyMode = 0
/*Uses high quality importance sampling to process the radiance map. In general, this results in much higher quality than [constant SKY_MODE_REALTIME] but takes much longer to generate. This should not be used if you plan on changing the sky at runtime. If you are finding that the reflection is not blurry enough and is showing sparkles or fireflies, try increasing [member ProjectSettings.rendering/reflections/sky_reflections/ggx_samples].*/
	SkyModeQuality SkyMode = 1
/*Uses the same high quality importance sampling to process the radiance map as [constant SKY_MODE_QUALITY], but updates over several frames. The number of frames is determined by [member ProjectSettings.rendering/reflections/sky_reflections/roughness_layers]. Use this when you need highest quality radiance maps, but have a sky that updates slowly.*/
	SkyModeIncremental SkyMode = 2
/*Uses the fast filtering algorithm to process the radiance map. In general this results in lower quality, but substantially faster run times. If you need better quality, but still need to update the sky every frame, consider turning on [member ProjectSettings.rendering/reflections/sky_reflections/fast_filter_high_quality].
[b]Note:[/b] The fast filtering algorithm is limited to 256×256 cubemaps, so [method sky_set_radiance_size] must be set to [code]256[/code]. Otherwise, a warning is printed and the overridden radiance size is ignored.*/
	SkyModeRealtime SkyMode = 3
)
type CompositorEffectFlags = classdb.RenderingServerCompositorEffectFlags

const (
/*The rendering effect requires the color buffer to be resolved if MSAA is enabled.*/
	CompositorEffectFlagAccessResolvedColor CompositorEffectFlags = 1
/*The rendering effect requires the depth buffer to be resolved if MSAA is enabled.*/
	CompositorEffectFlagAccessResolvedDepth CompositorEffectFlags = 2
/*The rendering effect requires motion vectors to be produced.*/
	CompositorEffectFlagNeedsMotionVectors CompositorEffectFlags = 4
/*The rendering effect requires normals and roughness g-buffer to be produced (Forward+ only).*/
	CompositorEffectFlagNeedsRoughness CompositorEffectFlags = 8
/*The rendering effect requires specular data to be separated out (Forward+ only).*/
	CompositorEffectFlagNeedsSeparateSpecular CompositorEffectFlags = 16
)
type CompositorEffectCallbackType = classdb.RenderingServerCompositorEffectCallbackType

const (
/*The callback is called before our opaque rendering pass, but after depth prepass (if applicable).*/
	CompositorEffectCallbackTypePreOpaque CompositorEffectCallbackType = 0
/*The callback is called after our opaque rendering pass, but before our sky is rendered.*/
	CompositorEffectCallbackTypePostOpaque CompositorEffectCallbackType = 1
/*The callback is called after our sky is rendered, but before our back buffers are created (and if enabled, before subsurface scattering and/or screen space reflections).*/
	CompositorEffectCallbackTypePostSky CompositorEffectCallbackType = 2
/*The callback is called before our transparent rendering pass, but after our sky is rendered and we've created our back buffers.*/
	CompositorEffectCallbackTypePreTransparent CompositorEffectCallbackType = 3
/*The callback is called after our transparent rendering pass, but before any build in post effects and output to our render target.*/
	CompositorEffectCallbackTypePostTransparent CompositorEffectCallbackType = 4
	CompositorEffectCallbackTypeAny CompositorEffectCallbackType = -1
)
type EnvironmentBG = classdb.RenderingServerEnvironmentBG

const (
/*Use the clear color as background.*/
	EnvBgClearColor EnvironmentBG = 0
/*Use a specified color as the background.*/
	EnvBgColor EnvironmentBG = 1
/*Use a sky resource for the background.*/
	EnvBgSky EnvironmentBG = 2
/*Use a specified canvas layer as the background. This can be useful for instantiating a 2D scene in a 3D world.*/
	EnvBgCanvas EnvironmentBG = 3
/*Do not clear the background, use whatever was rendered last frame as the background.*/
	EnvBgKeep EnvironmentBG = 4
/*Displays a camera feed in the background.*/
	EnvBgCameraFeed EnvironmentBG = 5
/*Represents the size of the [enum EnvironmentBG] enum.*/
	EnvBgMax EnvironmentBG = 6
)
type EnvironmentAmbientSource = classdb.RenderingServerEnvironmentAmbientSource

const (
/*Gather ambient light from whichever source is specified as the background.*/
	EnvAmbientSourceBg EnvironmentAmbientSource = 0
/*Disable ambient light.*/
	EnvAmbientSourceDisabled EnvironmentAmbientSource = 1
/*Specify a specific [Color] for ambient light.*/
	EnvAmbientSourceColor EnvironmentAmbientSource = 2
/*Gather ambient light from the [Sky] regardless of what the background is.*/
	EnvAmbientSourceSky EnvironmentAmbientSource = 3
)
type EnvironmentReflectionSource = classdb.RenderingServerEnvironmentReflectionSource

const (
/*Use the background for reflections.*/
	EnvReflectionSourceBg EnvironmentReflectionSource = 0
/*Disable reflections.*/
	EnvReflectionSourceDisabled EnvironmentReflectionSource = 1
/*Use the [Sky] for reflections regardless of what the background is.*/
	EnvReflectionSourceSky EnvironmentReflectionSource = 2
)
type EnvironmentGlowBlendMode = classdb.RenderingServerEnvironmentGlowBlendMode

const (
/*Additive glow blending mode. Mostly used for particles, glows (bloom), lens flare, bright sources.*/
	EnvGlowBlendModeAdditive EnvironmentGlowBlendMode = 0
/*Screen glow blending mode. Increases brightness, used frequently with bloom.*/
	EnvGlowBlendModeScreen EnvironmentGlowBlendMode = 1
/*Soft light glow blending mode. Modifies contrast, exposes shadows and highlights (vivid bloom).*/
	EnvGlowBlendModeSoftlight EnvironmentGlowBlendMode = 2
/*Replace glow blending mode. Replaces all pixels' color by the glow value. This can be used to simulate a full-screen blur effect by tweaking the glow parameters to match the original image's brightness.*/
	EnvGlowBlendModeReplace EnvironmentGlowBlendMode = 3
/*Mixes the glow with the underlying color to avoid increasing brightness as much while still maintaining a glow effect.*/
	EnvGlowBlendModeMix EnvironmentGlowBlendMode = 4
)
type EnvironmentFogMode = classdb.RenderingServerEnvironmentFogMode

const (
/*Use a physically-based fog model defined primarily by fog density.*/
	EnvFogModeExponential EnvironmentFogMode = 0
/*Use a simple fog model defined by start and end positions and a custom curve. While not physically accurate, this model can be useful when you need more artistic control.*/
	EnvFogModeDepth EnvironmentFogMode = 1
)
type EnvironmentToneMapper = classdb.RenderingServerEnvironmentToneMapper

const (
/*Output color as they came in. This can cause bright lighting to look blown out, with noticeable clipping in the output colors.*/
	EnvToneMapperLinear EnvironmentToneMapper = 0
/*Use the Reinhard tonemapper. Performs a variation on rendered pixels' colors by this formula: [code]color = color / (1 + color)[/code]. This avoids clipping bright highlights, but the resulting image can look a bit dull.*/
	EnvToneMapperReinhard EnvironmentToneMapper = 1
/*Use the filmic tonemapper. This avoids clipping bright highlights, with a resulting image that usually looks more vivid than [constant ENV_TONE_MAPPER_REINHARD].*/
	EnvToneMapperFilmic EnvironmentToneMapper = 2
/*Use the Academy Color Encoding System tonemapper. ACES is slightly more expensive than other options, but it handles bright lighting in a more realistic fashion by desaturating it as it becomes brighter. ACES typically has a more contrasted output compared to [constant ENV_TONE_MAPPER_REINHARD] and [constant ENV_TONE_MAPPER_FILMIC].
[b]Note:[/b] This tonemapping operator is called "ACES Fitted" in Godot 3.x.*/
	EnvToneMapperAces EnvironmentToneMapper = 3
)
type EnvironmentSSRRoughnessQuality = classdb.RenderingServerEnvironmentSSRRoughnessQuality

const (
/*Lowest quality of roughness filter for screen-space reflections. Rough materials will not have blurrier screen-space reflections compared to smooth (non-rough) materials. This is the fastest option.*/
	EnvSsrRoughnessQualityDisabled EnvironmentSSRRoughnessQuality = 0
/*Low quality of roughness filter for screen-space reflections.*/
	EnvSsrRoughnessQualityLow EnvironmentSSRRoughnessQuality = 1
/*Medium quality of roughness filter for screen-space reflections.*/
	EnvSsrRoughnessQualityMedium EnvironmentSSRRoughnessQuality = 2
/*High quality of roughness filter for screen-space reflections. This is the slowest option.*/
	EnvSsrRoughnessQualityHigh EnvironmentSSRRoughnessQuality = 3
)
type EnvironmentSSAOQuality = classdb.RenderingServerEnvironmentSSAOQuality

const (
/*Lowest quality of screen-space ambient occlusion.*/
	EnvSsaoQualityVeryLow EnvironmentSSAOQuality = 0
/*Low quality screen-space ambient occlusion.*/
	EnvSsaoQualityLow EnvironmentSSAOQuality = 1
/*Medium quality screen-space ambient occlusion.*/
	EnvSsaoQualityMedium EnvironmentSSAOQuality = 2
/*High quality screen-space ambient occlusion.*/
	EnvSsaoQualityHigh EnvironmentSSAOQuality = 3
/*Highest quality screen-space ambient occlusion. Uses the adaptive target setting which can be dynamically adjusted to smoothly balance performance and visual quality.*/
	EnvSsaoQualityUltra EnvironmentSSAOQuality = 4
)
type EnvironmentSSILQuality = classdb.RenderingServerEnvironmentSSILQuality

const (
/*Lowest quality of screen-space indirect lighting.*/
	EnvSsilQualityVeryLow EnvironmentSSILQuality = 0
/*Low quality screen-space indirect lighting.*/
	EnvSsilQualityLow EnvironmentSSILQuality = 1
/*High quality screen-space indirect lighting.*/
	EnvSsilQualityMedium EnvironmentSSILQuality = 2
/*High quality screen-space indirect lighting.*/
	EnvSsilQualityHigh EnvironmentSSILQuality = 3
/*Highest quality screen-space indirect lighting. Uses the adaptive target setting which can be dynamically adjusted to smoothly balance performance and visual quality.*/
	EnvSsilQualityUltra EnvironmentSSILQuality = 4
)
type EnvironmentSDFGIYScale = classdb.RenderingServerEnvironmentSDFGIYScale

const (
/*Use 50% scale for SDFGI on the Y (vertical) axis. SDFGI cells will be twice as short as they are wide. This allows providing increased GI detail and reduced light leaking with thin floors and ceilings. This is usually the best choice for scenes that don't feature much verticality.*/
	EnvSdfgiYScale50Percent EnvironmentSDFGIYScale = 0
/*Use 75% scale for SDFGI on the Y (vertical) axis. This is a balance between the 50% and 100% SDFGI Y scales.*/
	EnvSdfgiYScale75Percent EnvironmentSDFGIYScale = 1
/*Use 100% scale for SDFGI on the Y (vertical) axis. SDFGI cells will be as tall as they are wide. This is usually the best choice for highly vertical scenes. The downside is that light leaking may become more noticeable with thin floors and ceilings.*/
	EnvSdfgiYScale100Percent EnvironmentSDFGIYScale = 2
)
type EnvironmentSDFGIRayCount = classdb.RenderingServerEnvironmentSDFGIRayCount

const (
/*Throw 4 rays per frame when converging SDFGI. This has the lowest GPU requirements, but creates the most noisy result.*/
	EnvSdfgiRayCount4 EnvironmentSDFGIRayCount = 0
/*Throw 8 rays per frame when converging SDFGI.*/
	EnvSdfgiRayCount8 EnvironmentSDFGIRayCount = 1
/*Throw 16 rays per frame when converging SDFGI.*/
	EnvSdfgiRayCount16 EnvironmentSDFGIRayCount = 2
/*Throw 32 rays per frame when converging SDFGI.*/
	EnvSdfgiRayCount32 EnvironmentSDFGIRayCount = 3
/*Throw 64 rays per frame when converging SDFGI.*/
	EnvSdfgiRayCount64 EnvironmentSDFGIRayCount = 4
/*Throw 96 rays per frame when converging SDFGI. This has high GPU requirements.*/
	EnvSdfgiRayCount96 EnvironmentSDFGIRayCount = 5
/*Throw 128 rays per frame when converging SDFGI. This has very high GPU requirements, but creates the least noisy result.*/
	EnvSdfgiRayCount128 EnvironmentSDFGIRayCount = 6
/*Represents the size of the [enum EnvironmentSDFGIRayCount] enum.*/
	EnvSdfgiRayCountMax EnvironmentSDFGIRayCount = 7
)
type EnvironmentSDFGIFramesToConverge = classdb.RenderingServerEnvironmentSDFGIFramesToConverge

const (
/*Converge SDFGI over 5 frames. This is the most responsive, but creates the most noisy result with a given ray count.*/
	EnvSdfgiConvergeIn5Frames EnvironmentSDFGIFramesToConverge = 0
/*Configure SDFGI to fully converge over 10 frames.*/
	EnvSdfgiConvergeIn10Frames EnvironmentSDFGIFramesToConverge = 1
/*Configure SDFGI to fully converge over 15 frames.*/
	EnvSdfgiConvergeIn15Frames EnvironmentSDFGIFramesToConverge = 2
/*Configure SDFGI to fully converge over 20 frames.*/
	EnvSdfgiConvergeIn20Frames EnvironmentSDFGIFramesToConverge = 3
/*Configure SDFGI to fully converge over 25 frames.*/
	EnvSdfgiConvergeIn25Frames EnvironmentSDFGIFramesToConverge = 4
/*Configure SDFGI to fully converge over 30 frames. This is the least responsive, but creates the least noisy result with a given ray count.*/
	EnvSdfgiConvergeIn30Frames EnvironmentSDFGIFramesToConverge = 5
/*Represents the size of the [enum EnvironmentSDFGIFramesToConverge] enum.*/
	EnvSdfgiConvergeMax EnvironmentSDFGIFramesToConverge = 6
)
type EnvironmentSDFGIFramesToUpdateLight = classdb.RenderingServerEnvironmentSDFGIFramesToUpdateLight

const (
/*Update indirect light from dynamic lights in SDFGI over 1 frame. This is the most responsive, but has the highest GPU requirements.*/
	EnvSdfgiUpdateLightIn1Frame EnvironmentSDFGIFramesToUpdateLight = 0
/*Update indirect light from dynamic lights in SDFGI over 2 frames.*/
	EnvSdfgiUpdateLightIn2Frames EnvironmentSDFGIFramesToUpdateLight = 1
/*Update indirect light from dynamic lights in SDFGI over 4 frames.*/
	EnvSdfgiUpdateLightIn4Frames EnvironmentSDFGIFramesToUpdateLight = 2
/*Update indirect light from dynamic lights in SDFGI over 8 frames.*/
	EnvSdfgiUpdateLightIn8Frames EnvironmentSDFGIFramesToUpdateLight = 3
/*Update indirect light from dynamic lights in SDFGI over 16 frames. This is the least responsive, but has the lowest GPU requirements.*/
	EnvSdfgiUpdateLightIn16Frames EnvironmentSDFGIFramesToUpdateLight = 4
/*Represents the size of the [enum EnvironmentSDFGIFramesToUpdateLight] enum.*/
	EnvSdfgiUpdateLightMax EnvironmentSDFGIFramesToUpdateLight = 5
)
type SubSurfaceScatteringQuality = classdb.RenderingServerSubSurfaceScatteringQuality

const (
/*Disables subsurface scattering entirely, even on materials that have [member BaseMaterial3D.subsurf_scatter_enabled] set to [code]true[/code]. This has the lowest GPU requirements.*/
	SubSurfaceScatteringQualityDisabled SubSurfaceScatteringQuality = 0
/*Low subsurface scattering quality.*/
	SubSurfaceScatteringQualityLow SubSurfaceScatteringQuality = 1
/*Medium subsurface scattering quality.*/
	SubSurfaceScatteringQualityMedium SubSurfaceScatteringQuality = 2
/*High subsurface scattering quality. This has the highest GPU requirements.*/
	SubSurfaceScatteringQualityHigh SubSurfaceScatteringQuality = 3
)
type DOFBokehShape = classdb.RenderingServerDOFBokehShape

const (
/*Calculate the DOF blur using a box filter. The fastest option, but results in obvious lines in blur pattern.*/
	DofBokehBox DOFBokehShape = 0
/*Calculates DOF blur using a hexagon shaped filter.*/
	DofBokehHexagon DOFBokehShape = 1
/*Calculates DOF blur using a circle shaped filter. Best quality and most realistic, but slowest. Use only for areas where a lot of performance can be dedicated to post-processing (e.g. cutscenes).*/
	DofBokehCircle DOFBokehShape = 2
)
type DOFBlurQuality = classdb.RenderingServerDOFBlurQuality

const (
/*Lowest quality DOF blur. This is the fastest setting, but you may be able to see filtering artifacts.*/
	DofBlurQualityVeryLow DOFBlurQuality = 0
/*Low quality DOF blur.*/
	DofBlurQualityLow DOFBlurQuality = 1
/*Medium quality DOF blur.*/
	DofBlurQualityMedium DOFBlurQuality = 2
/*Highest quality DOF blur. Results in the smoothest looking blur by taking the most samples, but is also significantly slower.*/
	DofBlurQualityHigh DOFBlurQuality = 3
)
type InstanceType = classdb.RenderingServerInstanceType

const (
/*The instance does not have a type.*/
	InstanceNone InstanceType = 0
/*The instance is a mesh.*/
	InstanceMesh InstanceType = 1
/*The instance is a multimesh.*/
	InstanceMultimesh InstanceType = 2
/*The instance is a particle emitter.*/
	InstanceParticles InstanceType = 3
/*The instance is a GPUParticles collision shape.*/
	InstanceParticlesCollision InstanceType = 4
/*The instance is a light.*/
	InstanceLight InstanceType = 5
/*The instance is a reflection probe.*/
	InstanceReflectionProbe InstanceType = 6
/*The instance is a decal.*/
	InstanceDecal InstanceType = 7
/*The instance is a VoxelGI.*/
	InstanceVoxelGi InstanceType = 8
/*The instance is a lightmap.*/
	InstanceLightmap InstanceType = 9
/*The instance is an occlusion culling occluder.*/
	InstanceOccluder InstanceType = 10
/*The instance is a visible on-screen notifier.*/
	InstanceVisiblityNotifier InstanceType = 11
/*The instance is a fog volume.*/
	InstanceFogVolume InstanceType = 12
/*Represents the size of the [enum InstanceType] enum.*/
	InstanceMax InstanceType = 13
/*A combination of the flags of geometry instances (mesh, multimesh, immediate and particles).*/
	InstanceGeometryMask InstanceType = 14
)
type InstanceFlags = classdb.RenderingServerInstanceFlags

const (
/*Allows the instance to be used in baked lighting.*/
	InstanceFlagUseBakedLight InstanceFlags = 0
/*Allows the instance to be used with dynamic global illumination.*/
	InstanceFlagUseDynamicGi InstanceFlags = 1
/*When set, manually requests to draw geometry on next frame.*/
	InstanceFlagDrawNextFrameIfVisible InstanceFlags = 2
/*Always draw, even if the instance would be culled by occlusion culling. Does not affect view frustum culling.*/
	InstanceFlagIgnoreOcclusionCulling InstanceFlags = 3
/*Represents the size of the [enum InstanceFlags] enum.*/
	InstanceFlagMax InstanceFlags = 4
)
type ShadowCastingSetting = classdb.RenderingServerShadowCastingSetting

const (
/*Disable shadows from this instance.*/
	ShadowCastingSettingOff ShadowCastingSetting = 0
/*Cast shadows from this instance.*/
	ShadowCastingSettingOn ShadowCastingSetting = 1
/*Disable backface culling when rendering the shadow of the object. This is slightly slower but may result in more correct shadows.*/
	ShadowCastingSettingDoubleSided ShadowCastingSetting = 2
/*Only render the shadows from the object. The object itself will not be drawn.*/
	ShadowCastingSettingShadowsOnly ShadowCastingSetting = 3
)
type VisibilityRangeFadeMode = classdb.RenderingServerVisibilityRangeFadeMode

const (
/*Disable visibility range fading for the given instance.*/
	VisibilityRangeFadeDisabled VisibilityRangeFadeMode = 0
/*Fade-out the given instance when it approaches its visibility range limits.*/
	VisibilityRangeFadeSelf VisibilityRangeFadeMode = 1
/*Fade-in the given instance's dependencies when reaching its visibility range limits.*/
	VisibilityRangeFadeDependencies VisibilityRangeFadeMode = 2
)
type BakeChannels = classdb.RenderingServerBakeChannels

const (
/*Index of [Image] in array of [Image]s returned by [method bake_render_uv2]. Image uses [constant Image.FORMAT_RGBA8] and contains albedo color in the [code].rgb[/code] channels and alpha in the [code].a[/code] channel.*/
	BakeChannelAlbedoAlpha BakeChannels = 0
/*Index of [Image] in array of [Image]s returned by [method bake_render_uv2]. Image uses [constant Image.FORMAT_RGBA8] and contains the per-pixel normal of the object in the [code].rgb[/code] channels and nothing in the [code].a[/code] channel. The per-pixel normal is encoded as [code]normal * 0.5 + 0.5[/code].*/
	BakeChannelNormal BakeChannels = 1
/*Index of [Image] in array of [Image]s returned by [method bake_render_uv2]. Image uses [constant Image.FORMAT_RGBA8] and contains ambient occlusion (from material and decals only) in the [code].r[/code] channel, roughness in the [code].g[/code] channel, metallic in the [code].b[/code] channel and sub surface scattering amount in the [code].a[/code] channel.*/
	BakeChannelOrm BakeChannels = 2
/*Index of [Image] in array of [Image]s returned by [method bake_render_uv2]. Image uses [constant Image.FORMAT_RGBAH] and contains emission color in the [code].rgb[/code] channels and nothing in the [code].a[/code] channel.*/
	BakeChannelEmission BakeChannels = 3
)
type CanvasTextureChannel = classdb.RenderingServerCanvasTextureChannel

const (
/*Diffuse canvas texture ([member CanvasTexture.diffuse_texture]).*/
	CanvasTextureChannelDiffuse CanvasTextureChannel = 0
/*Normal map canvas texture ([member CanvasTexture.normal_texture]).*/
	CanvasTextureChannelNormal CanvasTextureChannel = 1
/*Specular map canvas texture ([member CanvasTexture.specular_texture]).*/
	CanvasTextureChannelSpecular CanvasTextureChannel = 2
)
type NinePatchAxisMode = classdb.RenderingServerNinePatchAxisMode

const (
/*The nine patch gets stretched where needed.*/
	NinePatchStretch NinePatchAxisMode = 0
/*The nine patch gets filled with tiles where needed.*/
	NinePatchTile NinePatchAxisMode = 1
/*The nine patch gets filled with tiles where needed and stretches them a bit if needed.*/
	NinePatchTileFit NinePatchAxisMode = 2
)
type CanvasItemTextureFilter = classdb.RenderingServerCanvasItemTextureFilter

const (
/*Uses the default filter mode for this [Viewport].*/
	CanvasItemTextureFilterDefault CanvasItemTextureFilter = 0
/*The texture filter reads from the nearest pixel only. This makes the texture look pixelated from up close, and grainy from a distance (due to mipmaps not being sampled).*/
	CanvasItemTextureFilterNearest CanvasItemTextureFilter = 1
/*The texture filter blends between the nearest 4 pixels. This makes the texture look smooth from up close, and grainy from a distance (due to mipmaps not being sampled).*/
	CanvasItemTextureFilterLinear CanvasItemTextureFilter = 2
/*The texture filter reads from the nearest pixel and blends between the nearest 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]). This makes the texture look pixelated from up close, and smooth from a distance.
Use this for non-pixel art textures that may be viewed at a low scale (e.g. due to [Camera2D] zoom or sprite scaling), as mipmaps are important to smooth out pixels that are smaller than on-screen pixels.*/
	CanvasItemTextureFilterNearestWithMipmaps CanvasItemTextureFilter = 3
/*The texture filter blends between the nearest 4 pixels and between the nearest 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]). This makes the texture look smooth from up close, and smooth from a distance.
Use this for non-pixel art textures that may be viewed at a low scale (e.g. due to [Camera2D] zoom or sprite scaling), as mipmaps are important to smooth out pixels that are smaller than on-screen pixels.*/
	CanvasItemTextureFilterLinearWithMipmaps CanvasItemTextureFilter = 4
/*The texture filter reads from the nearest pixel and blends between 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]) based on the angle between the surface and the camera view. This makes the texture look pixelated from up close, and smooth from a distance. Anisotropic filtering improves texture quality on surfaces that are almost in line with the camera, but is slightly slower. The anisotropic filtering level can be changed by adjusting [member ProjectSettings.rendering/textures/default_filters/anisotropic_filtering_level].
[b]Note:[/b] This texture filter is rarely useful in 2D projects. [constant CANVAS_ITEM_TEXTURE_FILTER_NEAREST_WITH_MIPMAPS] is usually more appropriate in this case.*/
	CanvasItemTextureFilterNearestWithMipmapsAnisotropic CanvasItemTextureFilter = 5
/*The texture filter blends between the nearest 4 pixels and blends between 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]) based on the angle between the surface and the camera view. This makes the texture look smooth from up close, and smooth from a distance. Anisotropic filtering improves texture quality on surfaces that are almost in line with the camera, but is slightly slower. The anisotropic filtering level can be changed by adjusting [member ProjectSettings.rendering/textures/default_filters/anisotropic_filtering_level].
[b]Note:[/b] This texture filter is rarely useful in 2D projects. [constant CANVAS_ITEM_TEXTURE_FILTER_LINEAR_WITH_MIPMAPS] is usually more appropriate in this case.*/
	CanvasItemTextureFilterLinearWithMipmapsAnisotropic CanvasItemTextureFilter = 6
/*Max value for [enum CanvasItemTextureFilter] enum.*/
	CanvasItemTextureFilterMax CanvasItemTextureFilter = 7
)
type CanvasItemTextureRepeat = classdb.RenderingServerCanvasItemTextureRepeat

const (
/*Uses the default repeat mode for this [Viewport].*/
	CanvasItemTextureRepeatDefault CanvasItemTextureRepeat = 0
/*Disables textures repeating. Instead, when reading UVs outside the 0-1 range, the value will be clamped to the edge of the texture, resulting in a stretched out look at the borders of the texture.*/
	CanvasItemTextureRepeatDisabled CanvasItemTextureRepeat = 1
/*Enables the texture to repeat when UV coordinates are outside the 0-1 range. If using one of the linear filtering modes, this can result in artifacts at the edges of a texture when the sampler filters across the edges of the texture.*/
	CanvasItemTextureRepeatEnabled CanvasItemTextureRepeat = 2
/*Flip the texture when repeating so that the edge lines up instead of abruptly changing.*/
	CanvasItemTextureRepeatMirror CanvasItemTextureRepeat = 3
/*Max value for [enum CanvasItemTextureRepeat] enum.*/
	CanvasItemTextureRepeatMax CanvasItemTextureRepeat = 4
)
type CanvasGroupMode = classdb.RenderingServerCanvasGroupMode

const (
/*Child draws over parent and is not clipped.*/
	CanvasGroupModeDisabled CanvasGroupMode = 0
/*Parent is used for the purposes of clipping only. Child is clipped to the parent's visible area, parent is not drawn.*/
	CanvasGroupModeClipOnly CanvasGroupMode = 1
/*Parent is used for clipping child, but parent is also drawn underneath child as normal before clipping child to its visible area.*/
	CanvasGroupModeClipAndDraw CanvasGroupMode = 2
	CanvasGroupModeTransparent CanvasGroupMode = 3
)
type CanvasLightMode = classdb.RenderingServerCanvasLightMode

const (
/*2D point light (see [PointLight2D]).*/
	CanvasLightModePoint CanvasLightMode = 0
/*2D directional (sun/moon) light (see [DirectionalLight2D]).*/
	CanvasLightModeDirectional CanvasLightMode = 1
)
type CanvasLightBlendMode = classdb.RenderingServerCanvasLightBlendMode

const (
/*Adds light color additive to the canvas.*/
	CanvasLightBlendModeAdd CanvasLightBlendMode = 0
/*Adds light color subtractive to the canvas.*/
	CanvasLightBlendModeSub CanvasLightBlendMode = 1
/*The light adds color depending on transparency.*/
	CanvasLightBlendModeMix CanvasLightBlendMode = 2
)
type CanvasLightShadowFilter = classdb.RenderingServerCanvasLightShadowFilter

const (
/*Do not apply a filter to canvas light shadows.*/
	CanvasLightFilterNone CanvasLightShadowFilter = 0
/*Use PCF5 filtering to filter canvas light shadows.*/
	CanvasLightFilterPcf5 CanvasLightShadowFilter = 1
/*Use PCF13 filtering to filter canvas light shadows.*/
	CanvasLightFilterPcf13 CanvasLightShadowFilter = 2
/*Max value of the [enum CanvasLightShadowFilter] enum.*/
	CanvasLightFilterMax CanvasLightShadowFilter = 3
)
type CanvasOccluderPolygonCullMode = classdb.RenderingServerCanvasOccluderPolygonCullMode

const (
/*Culling of the canvas occluder is disabled.*/
	CanvasOccluderPolygonCullDisabled CanvasOccluderPolygonCullMode = 0
/*Culling of the canvas occluder is clockwise.*/
	CanvasOccluderPolygonCullClockwise CanvasOccluderPolygonCullMode = 1
/*Culling of the canvas occluder is counterclockwise.*/
	CanvasOccluderPolygonCullCounterClockwise CanvasOccluderPolygonCullMode = 2
)
type GlobalShaderParameterType = classdb.RenderingServerGlobalShaderParameterType

const (
/*Boolean global shader parameter ([code]global uniform bool ...[/code]).*/
	GlobalVarTypeBool GlobalShaderParameterType = 0
/*2-dimensional boolean vector global shader parameter ([code]global uniform bvec2 ...[/code]).*/
	GlobalVarTypeBvec2 GlobalShaderParameterType = 1
/*3-dimensional boolean vector global shader parameter ([code]global uniform bvec3 ...[/code]).*/
	GlobalVarTypeBvec3 GlobalShaderParameterType = 2
/*4-dimensional boolean vector global shader parameter ([code]global uniform bvec4 ...[/code]).*/
	GlobalVarTypeBvec4 GlobalShaderParameterType = 3
/*Integer global shader parameter ([code]global uniform int ...[/code]).*/
	GlobalVarTypeInt GlobalShaderParameterType = 4
/*2-dimensional integer vector global shader parameter ([code]global uniform ivec2 ...[/code]).*/
	GlobalVarTypeIvec2 GlobalShaderParameterType = 5
/*3-dimensional integer vector global shader parameter ([code]global uniform ivec3 ...[/code]).*/
	GlobalVarTypeIvec3 GlobalShaderParameterType = 6
/*4-dimensional integer vector global shader parameter ([code]global uniform ivec4 ...[/code]).*/
	GlobalVarTypeIvec4 GlobalShaderParameterType = 7
/*2-dimensional integer rectangle global shader parameter ([code]global uniform ivec4 ...[/code]). Equivalent to [constant GLOBAL_VAR_TYPE_IVEC4] in shader code, but exposed as a [Rect2i] in the editor UI.*/
	GlobalVarTypeRect2i GlobalShaderParameterType = 8
/*Unsigned integer global shader parameter ([code]global uniform uint ...[/code]).*/
	GlobalVarTypeUint GlobalShaderParameterType = 9
/*2-dimensional unsigned integer vector global shader parameter ([code]global uniform uvec2 ...[/code]).*/
	GlobalVarTypeUvec2 GlobalShaderParameterType = 10
/*3-dimensional unsigned integer vector global shader parameter ([code]global uniform uvec3 ...[/code]).*/
	GlobalVarTypeUvec3 GlobalShaderParameterType = 11
/*4-dimensional unsigned integer vector global shader parameter ([code]global uniform uvec4 ...[/code]).*/
	GlobalVarTypeUvec4 GlobalShaderParameterType = 12
/*Single-precision floating-point global shader parameter ([code]global uniform float ...[/code]).*/
	GlobalVarTypeFloat GlobalShaderParameterType = 13
/*2-dimensional floating-point vector global shader parameter ([code]global uniform vec2 ...[/code]).*/
	GlobalVarTypeVec2 GlobalShaderParameterType = 14
/*3-dimensional floating-point vector global shader parameter ([code]global uniform vec3 ...[/code]).*/
	GlobalVarTypeVec3 GlobalShaderParameterType = 15
/*4-dimensional floating-point vector global shader parameter ([code]global uniform vec4 ...[/code]).*/
	GlobalVarTypeVec4 GlobalShaderParameterType = 16
/*Color global shader parameter ([code]global uniform vec4 ...[/code]). Equivalent to [constant GLOBAL_VAR_TYPE_VEC4] in shader code, but exposed as a [Color] in the editor UI.*/
	GlobalVarTypeColor GlobalShaderParameterType = 17
/*2-dimensional floating-point rectangle global shader parameter ([code]global uniform vec4 ...[/code]). Equivalent to [constant GLOBAL_VAR_TYPE_VEC4] in shader code, but exposed as a [Rect2] in the editor UI.*/
	GlobalVarTypeRect2 GlobalShaderParameterType = 18
/*2×2 matrix global shader parameter ([code]global uniform mat2 ...[/code]). Exposed as a [PackedInt32Array] in the editor UI.*/
	GlobalVarTypeMat2 GlobalShaderParameterType = 19
/*3×3 matrix global shader parameter ([code]global uniform mat3 ...[/code]). Exposed as a [Basis] in the editor UI.*/
	GlobalVarTypeMat3 GlobalShaderParameterType = 20
/*4×4 matrix global shader parameter ([code]global uniform mat4 ...[/code]). Exposed as a [Projection] in the editor UI.*/
	GlobalVarTypeMat4 GlobalShaderParameterType = 21
/*2-dimensional transform global shader parameter ([code]global uniform mat2x3 ...[/code]). Exposed as a [Transform2D] in the editor UI.*/
	GlobalVarTypeTransform2d GlobalShaderParameterType = 22
/*3-dimensional transform global shader parameter ([code]global uniform mat3x4 ...[/code]). Exposed as a [Transform3D] in the editor UI.*/
	GlobalVarTypeTransform GlobalShaderParameterType = 23
/*2D sampler global shader parameter ([code]global uniform sampler2D ...[/code]). Exposed as a [Texture2D] in the editor UI.*/
	GlobalVarTypeSampler2d GlobalShaderParameterType = 24
/*2D sampler array global shader parameter ([code]global uniform sampler2DArray ...[/code]). Exposed as a [Texture2DArray] in the editor UI.*/
	GlobalVarTypeSampler2darray GlobalShaderParameterType = 25
/*3D sampler global shader parameter ([code]global uniform sampler3D ...[/code]). Exposed as a [Texture3D] in the editor UI.*/
	GlobalVarTypeSampler3d GlobalShaderParameterType = 26
/*Cubemap sampler global shader parameter ([code]global uniform samplerCube ...[/code]). Exposed as a [Cubemap] in the editor UI.*/
	GlobalVarTypeSamplercube GlobalShaderParameterType = 27
/*Represents the size of the [enum GlobalShaderParameterType] enum.*/
	GlobalVarTypeMax GlobalShaderParameterType = 28
)
type RenderingInfo = classdb.RenderingServerRenderingInfo

const (
/*Number of objects rendered in the current 3D scene. This varies depending on camera position and rotation.*/
	RenderingInfoTotalObjectsInFrame RenderingInfo = 0
/*Number of points, lines, or triangles rendered in the current 3D scene. This varies depending on camera position and rotation.*/
	RenderingInfoTotalPrimitivesInFrame RenderingInfo = 1
/*Number of draw calls performed to render in the current 3D scene. This varies depending on camera position and rotation.*/
	RenderingInfoTotalDrawCallsInFrame RenderingInfo = 2
/*Texture memory used (in bytes).*/
	RenderingInfoTextureMemUsed RenderingInfo = 3
/*Buffer memory used (in bytes). This includes vertex data, uniform buffers, and many miscellaneous buffer types used internally.*/
	RenderingInfoBufferMemUsed RenderingInfo = 4
/*Video memory used (in bytes). When using the Forward+ or mobile rendering backends, this is always greater than the sum of [constant RENDERING_INFO_TEXTURE_MEM_USED] and [constant RENDERING_INFO_BUFFER_MEM_USED], since there is miscellaneous data not accounted for by those two metrics. When using the GL Compatibility backend, this is equal to the sum of [constant RENDERING_INFO_TEXTURE_MEM_USED] and [constant RENDERING_INFO_BUFFER_MEM_USED].*/
	RenderingInfoVideoMemUsed RenderingInfo = 5
)
type Features = classdb.RenderingServerFeatures

const (
	FeatureShaders Features = 0
	FeatureMultithreaded Features = 1
)
