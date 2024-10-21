package GLTFState

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Contains all nodes and resources of a GLTF file. This is used by [GLTFDocument] as data storage, which allows [GLTFDocument] and all [GLTFDocumentExtension] classes to remain stateless.
GLTFState can be populated by [GLTFDocument] reading a file or by converting a Godot scene. Then the data can either be used to create a Godot scene or save to a GLTF file. The code that converts to/from a Godot scene can be intercepted at arbitrary points by [GLTFDocumentExtension] classes. This allows for custom data to be stored in the GLTF file or for custom data to be converted to/from Godot nodes.

*/
type Simple [1]classdb.GLTFState
func (self Simple) AddUsedExtension(extension_name string, required bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddUsedExtension(gc.String(extension_name), required)
}
func (self Simple) AppendDataToBuffers(data []byte, deduplication bool) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AppendDataToBuffers(gc.PackedByteSlice(data), deduplication)))
}
func (self Simple) GetJson() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).GetJson(gc))
}
func (self Simple) SetJson(json gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetJson(json)
}
func (self Simple) GetMajorVersion() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMajorVersion()))
}
func (self Simple) SetMajorVersion(major_version int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMajorVersion(gd.Int(major_version))
}
func (self Simple) GetMinorVersion() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMinorVersion()))
}
func (self Simple) SetMinorVersion(minor_version int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMinorVersion(gd.Int(minor_version))
}
func (self Simple) GetCopyright() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetCopyright(gc).String())
}
func (self Simple) SetCopyright(copyright string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCopyright(gc.String(copyright))
}
func (self Simple) GetGlbData() []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).GetGlbData(gc).Bytes())
}
func (self Simple) SetGlbData(glb_data []byte) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlbData(gc.PackedByteSlice(glb_data))
}
func (self Simple) GetUseNamedSkinBinds() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetUseNamedSkinBinds())
}
func (self Simple) SetUseNamedSkinBinds(use_named_skin_binds bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseNamedSkinBinds(use_named_skin_binds)
}
func (self Simple) GetNodes() gd.ArrayOf[[1]classdb.GLTFNode] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.GLTFNode](Expert(self).GetNodes(gc))
}
func (self Simple) SetNodes(nodes gd.ArrayOf[[1]classdb.GLTFNode]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNodes(nodes)
}
func (self Simple) GetBuffers() gd.ArrayOf[gd.PackedByteArray] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.PackedByteArray](Expert(self).GetBuffers(gc))
}
func (self Simple) SetBuffers(buffers gd.ArrayOf[gd.PackedByteArray]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBuffers(buffers)
}
func (self Simple) GetBufferViews() gd.ArrayOf[[1]classdb.GLTFBufferView] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.GLTFBufferView](Expert(self).GetBufferViews(gc))
}
func (self Simple) SetBufferViews(buffer_views gd.ArrayOf[[1]classdb.GLTFBufferView]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBufferViews(buffer_views)
}
func (self Simple) GetAccessors() gd.ArrayOf[[1]classdb.GLTFAccessor] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.GLTFAccessor](Expert(self).GetAccessors(gc))
}
func (self Simple) SetAccessors(accessors gd.ArrayOf[[1]classdb.GLTFAccessor]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAccessors(accessors)
}
func (self Simple) GetMeshes() gd.ArrayOf[[1]classdb.GLTFMesh] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.GLTFMesh](Expert(self).GetMeshes(gc))
}
func (self Simple) SetMeshes(meshes gd.ArrayOf[[1]classdb.GLTFMesh]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMeshes(meshes)
}
func (self Simple) GetAnimationPlayersCount(idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetAnimationPlayersCount(gd.Int(idx))))
}
func (self Simple) GetAnimationPlayer(idx int) [1]classdb.AnimationPlayer {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.AnimationPlayer(Expert(self).GetAnimationPlayer(gc, gd.Int(idx)))
}
func (self Simple) GetMaterials() gd.ArrayOf[[1]classdb.Material] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.Material](Expert(self).GetMaterials(gc))
}
func (self Simple) SetMaterials(materials gd.ArrayOf[[1]classdb.Material]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaterials(materials)
}
func (self Simple) GetSceneName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetSceneName(gc).String())
}
func (self Simple) SetSceneName(scene_name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSceneName(gc.String(scene_name))
}
func (self Simple) GetBasePath() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetBasePath(gc).String())
}
func (self Simple) SetBasePath(base_path string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBasePath(gc.String(base_path))
}
func (self Simple) GetFilename() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetFilename(gc).String())
}
func (self Simple) SetFilename(filename string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFilename(gc.String(filename))
}
func (self Simple) GetRootNodes() gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).GetRootNodes(gc))
}
func (self Simple) SetRootNodes(root_nodes gd.PackedInt32Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRootNodes(root_nodes)
}
func (self Simple) GetTextures() gd.ArrayOf[[1]classdb.GLTFTexture] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.GLTFTexture](Expert(self).GetTextures(gc))
}
func (self Simple) SetTextures(textures gd.ArrayOf[[1]classdb.GLTFTexture]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTextures(textures)
}
func (self Simple) GetTextureSamplers() gd.ArrayOf[[1]classdb.GLTFTextureSampler] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.GLTFTextureSampler](Expert(self).GetTextureSamplers(gc))
}
func (self Simple) SetTextureSamplers(texture_samplers gd.ArrayOf[[1]classdb.GLTFTextureSampler]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTextureSamplers(texture_samplers)
}
func (self Simple) GetImages() gd.ArrayOf[[1]classdb.Texture2D] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.Texture2D](Expert(self).GetImages(gc))
}
func (self Simple) SetImages(images gd.ArrayOf[[1]classdb.Texture2D]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetImages(images)
}
func (self Simple) GetSkins() gd.ArrayOf[[1]classdb.GLTFSkin] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.GLTFSkin](Expert(self).GetSkins(gc))
}
func (self Simple) SetSkins(skins gd.ArrayOf[[1]classdb.GLTFSkin]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSkins(skins)
}
func (self Simple) GetCameras() gd.ArrayOf[[1]classdb.GLTFCamera] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.GLTFCamera](Expert(self).GetCameras(gc))
}
func (self Simple) SetCameras(cameras gd.ArrayOf[[1]classdb.GLTFCamera]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCameras(cameras)
}
func (self Simple) GetLights() gd.ArrayOf[[1]classdb.GLTFLight] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.GLTFLight](Expert(self).GetLights(gc))
}
func (self Simple) SetLights(lights gd.ArrayOf[[1]classdb.GLTFLight]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLights(lights)
}
func (self Simple) GetUniqueNames() gd.ArrayOf[gd.String] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.String](Expert(self).GetUniqueNames(gc))
}
func (self Simple) SetUniqueNames(unique_names gd.ArrayOf[gd.String]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUniqueNames(unique_names)
}
func (self Simple) GetUniqueAnimationNames() gd.ArrayOf[gd.String] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.String](Expert(self).GetUniqueAnimationNames(gc))
}
func (self Simple) SetUniqueAnimationNames(unique_animation_names gd.ArrayOf[gd.String]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUniqueAnimationNames(unique_animation_names)
}
func (self Simple) GetSkeletons() gd.ArrayOf[[1]classdb.GLTFSkeleton] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.GLTFSkeleton](Expert(self).GetSkeletons(gc))
}
func (self Simple) SetSkeletons(skeletons gd.ArrayOf[[1]classdb.GLTFSkeleton]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSkeletons(skeletons)
}
func (self Simple) GetCreateAnimations() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetCreateAnimations())
}
func (self Simple) SetCreateAnimations(create_animations bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCreateAnimations(create_animations)
}
func (self Simple) GetImportAsSkeletonBones() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetImportAsSkeletonBones())
}
func (self Simple) SetImportAsSkeletonBones(import_as_skeleton_bones bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetImportAsSkeletonBones(import_as_skeleton_bones)
}
func (self Simple) GetAnimations() gd.ArrayOf[[1]classdb.GLTFAnimation] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.GLTFAnimation](Expert(self).GetAnimations(gc))
}
func (self Simple) SetAnimations(animations gd.ArrayOf[[1]classdb.GLTFAnimation]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAnimations(animations)
}
func (self Simple) GetSceneNode(idx int) [1]classdb.Node {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Node(Expert(self).GetSceneNode(gc, gd.Int(idx)))
}
func (self Simple) GetNodeIndex(scene_node [1]classdb.Node) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetNodeIndex(scene_node)))
}
func (self Simple) GetAdditionalData(extension_name string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetAdditionalData(gc, gc.StringName(extension_name)))
}
func (self Simple) SetAdditionalData(extension_name string, additional_data gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAdditionalData(gc.StringName(extension_name), additional_data)
}
func (self Simple) GetHandleBinaryImage() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetHandleBinaryImage()))
}
func (self Simple) SetHandleBinaryImage(method int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHandleBinaryImage(gd.Int(method))
}
func (self Simple) SetBakeFps(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBakeFps(gd.Float(value))
}
func (self Simple) GetBakeFps() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetBakeFps()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.GLTFState
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Appends an extension to the list of extensions used by this GLTF file during serialization. If [param required] is true, the extension will also be added to the list of required extensions. Do not run this in [method GLTFDocumentExtension._export_post], as that stage is too late to add extensions. The final list is sorted alphabetically.
*/
//go:nosplit
func (self class) AddUsedExtension(extension_name gd.String, required bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(extension_name))
	callframe.Arg(frame, required)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_add_used_extension, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Appends the given byte array data to the buffers and creates a [GLTFBufferView] for it. The index of the destination [GLTFBufferView] is returned. If [param deduplication] is true, the buffers will first be searched for duplicate data, otherwise new bytes will always be appended.
*/
//go:nosplit
func (self class) AppendDataToBuffers(data gd.PackedByteArray, deduplication bool) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(data))
	callframe.Arg(frame, deduplication)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_append_data_to_buffers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetJson(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_json, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetJson(json gd.Dictionary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(json))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_json, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMajorVersion() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_major_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMajorVersion(major_version gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, major_version)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_major_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMinorVersion() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_minor_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMinorVersion(minor_version gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, minor_version)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_minor_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCopyright(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_copyright, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCopyright(copyright gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(copyright))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_copyright, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlbData(ctx gd.Lifetime) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_glb_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlbData(glb_data gd.PackedByteArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(glb_data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_glb_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUseNamedSkinBinds() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_use_named_skin_binds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseNamedSkinBinds(use_named_skin_binds bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use_named_skin_binds)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_use_named_skin_binds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an array of all [GLTFNode]s in the GLTF file. These are the nodes that [member GLTFNode.children] and [member root_nodes] refer to. This includes nodes that may not be generated in the Godot scene, or nodes that may generate multiple Godot scene nodes.
*/
//go:nosplit
func (self class) GetNodes(ctx gd.Lifetime) gd.ArrayOf[object.GLTFNode] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.GLTFNode](ret)
}
/*
Sets the [GLTFNode]s in the state. These are the nodes that [member GLTFNode.children] and [member root_nodes] refer to. Some of the nodes set here may not be generated in the Godot scene, or may generate multiple Godot scene nodes.
*/
//go:nosplit
func (self class) SetNodes(nodes gd.ArrayOf[object.GLTFNode])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(nodes))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBuffers(ctx gd.Lifetime) gd.ArrayOf[gd.PackedByteArray] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_buffers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.PackedByteArray](ret)
}
//go:nosplit
func (self class) SetBuffers(buffers gd.ArrayOf[gd.PackedByteArray])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(buffers))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_buffers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBufferViews(ctx gd.Lifetime) gd.ArrayOf[object.GLTFBufferView] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_buffer_views, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.GLTFBufferView](ret)
}
//go:nosplit
func (self class) SetBufferViews(buffer_views gd.ArrayOf[object.GLTFBufferView])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(buffer_views))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_buffer_views, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAccessors(ctx gd.Lifetime) gd.ArrayOf[object.GLTFAccessor] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_accessors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.GLTFAccessor](ret)
}
//go:nosplit
func (self class) SetAccessors(accessors gd.ArrayOf[object.GLTFAccessor])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(accessors))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_accessors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an array of all [GLTFMesh]es in the GLTF file. These are the meshes that the [member GLTFNode.mesh] index refers to.
*/
//go:nosplit
func (self class) GetMeshes(ctx gd.Lifetime) gd.ArrayOf[object.GLTFMesh] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_meshes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.GLTFMesh](ret)
}
/*
Sets the [GLTFMesh]es in the state. These are the meshes that the [member GLTFNode.mesh] index refers to.
*/
//go:nosplit
func (self class) SetMeshes(meshes gd.ArrayOf[object.GLTFMesh])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(meshes))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_meshes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of [AnimationPlayer] nodes in this [GLTFState]. These nodes are only used during the export process when converting Godot [AnimationPlayer] nodes to GLTF animations.
*/
//go:nosplit
func (self class) GetAnimationPlayersCount(idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_animation_players_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [AnimationPlayer] node with the given index. These nodes are only used during the export process when converting Godot [AnimationPlayer] nodes to GLTF animations.
*/
//go:nosplit
func (self class) GetAnimationPlayer(ctx gd.Lifetime, idx gd.Int) object.AnimationPlayer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_animation_player, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AnimationPlayer
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetMaterials(ctx gd.Lifetime) gd.ArrayOf[object.Material] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_materials, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.Material](ret)
}
//go:nosplit
func (self class) SetMaterials(materials gd.ArrayOf[object.Material])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(materials))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_materials, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSceneName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_scene_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSceneName(scene_name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(scene_name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_scene_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBasePath(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_base_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBasePath(base_path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(base_path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_base_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFilename(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_filename, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFilename(filename gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(filename))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_filename, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRootNodes(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_root_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRootNodes(root_nodes gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(root_nodes))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_root_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextures(ctx gd.Lifetime) gd.ArrayOf[object.GLTFTexture] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_textures, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.GLTFTexture](ret)
}
//go:nosplit
func (self class) SetTextures(textures gd.ArrayOf[object.GLTFTexture])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(textures))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_textures, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Retrieves the array of texture samplers that are used by the textures contained in the GLTF.
*/
//go:nosplit
func (self class) GetTextureSamplers(ctx gd.Lifetime) gd.ArrayOf[object.GLTFTextureSampler] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_texture_samplers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.GLTFTextureSampler](ret)
}
/*
Sets the array of texture samplers that are used by the textures contained in the GLTF.
*/
//go:nosplit
func (self class) SetTextureSamplers(texture_samplers gd.ArrayOf[object.GLTFTextureSampler])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture_samplers))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_texture_samplers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets the images of the GLTF file as an array of [Texture2D]s. These are the images that the [member GLTFTexture.src_image] index refers to.
*/
//go:nosplit
func (self class) GetImages(ctx gd.Lifetime) gd.ArrayOf[object.Texture2D] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_images, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.Texture2D](ret)
}
/*
Sets the images in the state stored as an array of [Texture2D]s. This can be used during export. These are the images that the [member GLTFTexture.src_image] index refers to.
*/
//go:nosplit
func (self class) SetImages(images gd.ArrayOf[object.Texture2D])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(images))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_images, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an array of all [GLTFSkin]s in the GLTF file. These are the skins that the [member GLTFNode.skin] index refers to.
*/
//go:nosplit
func (self class) GetSkins(ctx gd.Lifetime) gd.ArrayOf[object.GLTFSkin] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_skins, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.GLTFSkin](ret)
}
/*
Sets the [GLTFSkin]s in the state. These are the skins that the [member GLTFNode.skin] index refers to.
*/
//go:nosplit
func (self class) SetSkins(skins gd.ArrayOf[object.GLTFSkin])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(skins))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_skins, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an array of all [GLTFCamera]s in the GLTF file. These are the cameras that the [member GLTFNode.camera] index refers to.
*/
//go:nosplit
func (self class) GetCameras(ctx gd.Lifetime) gd.ArrayOf[object.GLTFCamera] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_cameras, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.GLTFCamera](ret)
}
/*
Sets the [GLTFCamera]s in the state. These are the cameras that the [member GLTFNode.camera] index refers to.
*/
//go:nosplit
func (self class) SetCameras(cameras gd.ArrayOf[object.GLTFCamera])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(cameras))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_cameras, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an array of all [GLTFLight]s in the GLTF file. These are the lights that the [member GLTFNode.light] index refers to.
*/
//go:nosplit
func (self class) GetLights(ctx gd.Lifetime) gd.ArrayOf[object.GLTFLight] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_lights, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.GLTFLight](ret)
}
/*
Sets the [GLTFLight]s in the state. These are the lights that the [member GLTFNode.light] index refers to.
*/
//go:nosplit
func (self class) SetLights(lights gd.ArrayOf[object.GLTFLight])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(lights))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_lights, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an array of unique node names. This is used in both the import process and export process.
*/
//go:nosplit
func (self class) GetUniqueNames(ctx gd.Lifetime) gd.ArrayOf[gd.String] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_unique_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.String](ret)
}
/*
Sets the unique node names in the state. This is used in both the import process and export process.
*/
//go:nosplit
func (self class) SetUniqueNames(unique_names gd.ArrayOf[gd.String])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(unique_names))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_unique_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an array of unique animation names. This is only used during the import process.
*/
//go:nosplit
func (self class) GetUniqueAnimationNames(ctx gd.Lifetime) gd.ArrayOf[gd.String] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_unique_animation_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.String](ret)
}
/*
Sets the unique animation names in the state. This is only used during the import process.
*/
//go:nosplit
func (self class) SetUniqueAnimationNames(unique_animation_names gd.ArrayOf[gd.String])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(unique_animation_names))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_unique_animation_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an array of all [GLTFSkeleton]s in the GLTF file. These are the skeletons that the [member GLTFNode.skeleton] index refers to.
*/
//go:nosplit
func (self class) GetSkeletons(ctx gd.Lifetime) gd.ArrayOf[object.GLTFSkeleton] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_skeletons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.GLTFSkeleton](ret)
}
/*
Sets the [GLTFSkeleton]s in the state. These are the skeletons that the [member GLTFNode.skeleton] index refers to.
*/
//go:nosplit
func (self class) SetSkeletons(skeletons gd.ArrayOf[object.GLTFSkeleton])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(skeletons))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_skeletons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCreateAnimations() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_create_animations, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCreateAnimations(create_animations bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, create_animations)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_create_animations, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetImportAsSkeletonBones() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_import_as_skeleton_bones, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetImportAsSkeletonBones(import_as_skeleton_bones bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, import_as_skeleton_bones)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_import_as_skeleton_bones, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an array of all [GLTFAnimation]s in the GLTF file. When importing, these will be generated as animations in an [AnimationPlayer] node. When exporting, these will be generated from Godot [AnimationPlayer] nodes.
*/
//go:nosplit
func (self class) GetAnimations(ctx gd.Lifetime) gd.ArrayOf[object.GLTFAnimation] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_animations, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.GLTFAnimation](ret)
}
/*
Sets the [GLTFAnimation]s in the state. When importing, these will be generated as animations in an [AnimationPlayer] node. When exporting, these will be generated from Godot [AnimationPlayer] nodes.
*/
//go:nosplit
func (self class) SetAnimations(animations gd.ArrayOf[object.GLTFAnimation])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(animations))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_animations, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the Godot scene node that corresponds to the same index as the [GLTFNode] it was generated from. This is the inverse of [method get_node_index]. Useful during the import process.
[b]Note:[/b] Not every [GLTFNode] will have a scene node generated, and not every generated scene node will have a corresponding [GLTFNode]. If there is no scene node for this [GLTFNode] index, [code]null[/code] is returned.
*/
//go:nosplit
func (self class) GetSceneNode(ctx gd.Lifetime, idx gd.Int) object.Node {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_scene_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Node
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the index of the [GLTFNode] corresponding to this Godot scene node. This is the inverse of [method get_scene_node]. Useful during the export process.
[b]Note:[/b] Not every Godot scene node will have a corresponding [GLTFNode], and not every [GLTFNode] will have a scene node generated. If there is no [GLTFNode] index for this scene node, [code]-1[/code] is returned.
*/
//go:nosplit
func (self class) GetNodeIndex(scene_node object.Node) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(scene_node[0].AsPointer())[0])
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_node_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets additional arbitrary data in this [GLTFState] instance. This can be used to keep per-file state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the return value can be anything you set. If nothing was set, the return value is null.
*/
//go:nosplit
func (self class) GetAdditionalData(ctx gd.Lifetime, extension_name gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(extension_name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_additional_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets additional arbitrary data in this [GLTFState] instance. This can be used to keep per-file state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The first argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the second argument can be anything you want.
*/
//go:nosplit
func (self class) SetAdditionalData(extension_name gd.StringName, additional_data gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(extension_name))
	callframe.Arg(frame, mmm.Get(additional_data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_additional_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHandleBinaryImage() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_handle_binary_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHandleBinaryImage(method gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, method)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_handle_binary_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetBakeFps(value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_bake_fps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBakeFps() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_bake_fps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsGLTFState() Expert { return self[0].AsGLTFState() }


//go:nosplit
func (self Simple) AsGLTFState() Simple { return self[0].AsGLTFState() }


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

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("GLTFState", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
