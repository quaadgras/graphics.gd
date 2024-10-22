package GLTFState

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Contains all nodes and resources of a GLTF file. This is used by [GLTFDocument] as data storage, which allows [GLTFDocument] and all [GLTFDocumentExtension] classes to remain stateless.
GLTFState can be populated by [GLTFDocument] reading a file or by converting a Godot scene. Then the data can either be used to create a Godot scene or save to a GLTF file. The code that converts to/from a Godot scene can be intercepted at arbitrary points by [GLTFDocumentExtension] classes. This allows for custom data to be stored in the GLTF file or for custom data to be converted to/from Godot nodes.

*/
type Go [1]classdb.GLTFState

/*
Appends an extension to the list of extensions used by this GLTF file during serialization. If [param required] is true, the extension will also be added to the list of required extensions. Do not run this in [method GLTFDocumentExtension._export_post], as that stage is too late to add extensions. The final list is sorted alphabetically.
*/
func (self Go) AddUsedExtension(extension_name string, required bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddUsedExtension(gc.String(extension_name), required)
}

/*
Appends the given byte array data to the buffers and creates a [GLTFBufferView] for it. The index of the destination [GLTFBufferView] is returned. If [param deduplication] is true, the buffers will first be searched for duplicate data, otherwise new bytes will always be appended.
*/
func (self Go) AppendDataToBuffers(data []byte, deduplication bool) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).AppendDataToBuffers(gc.PackedByteSlice(data), deduplication)))
}

/*
Returns the number of [AnimationPlayer] nodes in this [GLTFState]. These nodes are only used during the export process when converting Godot [AnimationPlayer] nodes to GLTF animations.
*/
func (self Go) GetAnimationPlayersCount(idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetAnimationPlayersCount(gd.Int(idx))))
}

/*
Returns the [AnimationPlayer] node with the given index. These nodes are only used during the export process when converting Godot [AnimationPlayer] nodes to GLTF animations.
*/
func (self Go) GetAnimationPlayer(idx int) gdclass.AnimationPlayer {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.AnimationPlayer(class(self).GetAnimationPlayer(gc, gd.Int(idx)))
}

/*
Returns the Godot scene node that corresponds to the same index as the [GLTFNode] it was generated from. This is the inverse of [method get_node_index]. Useful during the import process.
[b]Note:[/b] Not every [GLTFNode] will have a scene node generated, and not every generated scene node will have a corresponding [GLTFNode]. If there is no scene node for this [GLTFNode] index, [code]null[/code] is returned.
*/
func (self Go) GetSceneNode(idx int) gdclass.Node {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Node(class(self).GetSceneNode(gc, gd.Int(idx)))
}

/*
Returns the index of the [GLTFNode] corresponding to this Godot scene node. This is the inverse of [method get_scene_node]. Useful during the export process.
[b]Note:[/b] Not every Godot scene node will have a corresponding [GLTFNode], and not every [GLTFNode] will have a scene node generated. If there is no [GLTFNode] index for this scene node, [code]-1[/code] is returned.
*/
func (self Go) GetNodeIndex(scene_node gdclass.Node) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetNodeIndex(scene_node)))
}

/*
Gets additional arbitrary data in this [GLTFState] instance. This can be used to keep per-file state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the return value can be anything you set. If nothing was set, the return value is null.
*/
func (self Go) GetAdditionalData(extension_name string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).GetAdditionalData(gc, gc.StringName(extension_name)))
}

/*
Sets additional arbitrary data in this [GLTFState] instance. This can be used to keep per-file state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The first argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the second argument can be anything you want.
*/
func (self Go) SetAdditionalData(extension_name string, additional_data gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAdditionalData(gc.StringName(extension_name), additional_data)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GLTFState
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("GLTFState"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Json() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Dictionary(class(self).GetJson(gc))
}

func (self Go) SetJson(value gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJson(value)
}

func (self Go) MajorVersion() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMajorVersion()))
}

func (self Go) SetMajorVersion(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMajorVersion(gd.Int(value))
}

func (self Go) MinorVersion() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMinorVersion()))
}

func (self Go) SetMinorVersion(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMinorVersion(gd.Int(value))
}

func (self Go) Copyright() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetCopyright(gc).String())
}

func (self Go) SetCopyright(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCopyright(gc.String(value))
}

func (self Go) GlbData() []byte {
	gc := gd.GarbageCollector(); _ = gc
		return []byte(class(self).GetGlbData(gc).Bytes())
}

func (self Go) SetGlbData(value []byte) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlbData(gc.PackedByteSlice(value))
}

func (self Go) UseNamedSkinBinds() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetUseNamedSkinBinds())
}

func (self Go) SetUseNamedSkinBinds(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseNamedSkinBinds(value)
}

func (self Go) Nodes() gd.ArrayOf[gdclass.GLTFNode] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gdclass.GLTFNode](class(self).GetNodes(gc))
}

func (self Go) SetNodes(value gd.ArrayOf[gdclass.GLTFNode]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNodes(value)
}

func (self Go) Buffers() gd.ArrayOf[gd.PackedByteArray] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gd.PackedByteArray](class(self).GetBuffers(gc))
}

func (self Go) SetBuffers(value gd.ArrayOf[gd.PackedByteArray]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBuffers(value)
}

func (self Go) BufferViews() gd.ArrayOf[gdclass.GLTFBufferView] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gdclass.GLTFBufferView](class(self).GetBufferViews(gc))
}

func (self Go) SetBufferViews(value gd.ArrayOf[gdclass.GLTFBufferView]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBufferViews(value)
}

func (self Go) Accessors() gd.ArrayOf[gdclass.GLTFAccessor] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gdclass.GLTFAccessor](class(self).GetAccessors(gc))
}

func (self Go) SetAccessors(value gd.ArrayOf[gdclass.GLTFAccessor]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAccessors(value)
}

func (self Go) Meshes() gd.ArrayOf[gdclass.GLTFMesh] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gdclass.GLTFMesh](class(self).GetMeshes(gc))
}

func (self Go) SetMeshes(value gd.ArrayOf[gdclass.GLTFMesh]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMeshes(value)
}

func (self Go) Materials() gd.ArrayOf[gdclass.Material] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gdclass.Material](class(self).GetMaterials(gc))
}

func (self Go) SetMaterials(value gd.ArrayOf[gdclass.Material]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaterials(value)
}

func (self Go) SceneName() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetSceneName(gc).String())
}

func (self Go) SetSceneName(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSceneName(gc.String(value))
}

func (self Go) BasePath() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetBasePath(gc).String())
}

func (self Go) SetBasePath(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBasePath(gc.String(value))
}

func (self Go) Filename() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetFilename(gc).String())
}

func (self Go) SetFilename(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFilename(gc.String(value))
}

func (self Go) RootNodes() []int32 {
	gc := gd.GarbageCollector(); _ = gc
		return []int32(class(self).GetRootNodes(gc).AsSlice())
}

func (self Go) SetRootNodes(value []int32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRootNodes(gc.PackedInt32Slice(value))
}

func (self Go) Textures() gd.ArrayOf[gdclass.GLTFTexture] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gdclass.GLTFTexture](class(self).GetTextures(gc))
}

func (self Go) SetTextures(value gd.ArrayOf[gdclass.GLTFTexture]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextures(value)
}

func (self Go) TextureSamplers() gd.ArrayOf[gdclass.GLTFTextureSampler] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gdclass.GLTFTextureSampler](class(self).GetTextureSamplers(gc))
}

func (self Go) SetTextureSamplers(value gd.ArrayOf[gdclass.GLTFTextureSampler]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextureSamplers(value)
}

func (self Go) Images() gd.ArrayOf[gdclass.Texture2D] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gdclass.Texture2D](class(self).GetImages(gc))
}

func (self Go) SetImages(value gd.ArrayOf[gdclass.Texture2D]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetImages(value)
}

func (self Go) Skins() gd.ArrayOf[gdclass.GLTFSkin] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gdclass.GLTFSkin](class(self).GetSkins(gc))
}

func (self Go) SetSkins(value gd.ArrayOf[gdclass.GLTFSkin]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSkins(value)
}

func (self Go) Cameras() gd.ArrayOf[gdclass.GLTFCamera] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gdclass.GLTFCamera](class(self).GetCameras(gc))
}

func (self Go) SetCameras(value gd.ArrayOf[gdclass.GLTFCamera]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCameras(value)
}

func (self Go) Lights() gd.ArrayOf[gdclass.GLTFLight] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gdclass.GLTFLight](class(self).GetLights(gc))
}

func (self Go) SetLights(value gd.ArrayOf[gdclass.GLTFLight]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLights(value)
}

func (self Go) UniqueNames() gd.ArrayOf[gd.String] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gd.String](class(self).GetUniqueNames(gc))
}

func (self Go) SetUniqueNames(value gd.ArrayOf[gd.String]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUniqueNames(value)
}

func (self Go) UniqueAnimationNames() gd.ArrayOf[gd.String] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gd.String](class(self).GetUniqueAnimationNames(gc))
}

func (self Go) SetUniqueAnimationNames(value gd.ArrayOf[gd.String]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUniqueAnimationNames(value)
}

func (self Go) Skeletons() gd.ArrayOf[gdclass.GLTFSkeleton] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gdclass.GLTFSkeleton](class(self).GetSkeletons(gc))
}

func (self Go) SetSkeletons(value gd.ArrayOf[gdclass.GLTFSkeleton]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSkeletons(value)
}

func (self Go) CreateAnimations() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetCreateAnimations())
}

func (self Go) SetCreateAnimations(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCreateAnimations(value)
}

func (self Go) ImportAsSkeletonBones() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetImportAsSkeletonBones())
}

func (self Go) SetImportAsSkeletonBones(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetImportAsSkeletonBones(value)
}

func (self Go) Animations() gd.ArrayOf[gdclass.GLTFAnimation] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gdclass.GLTFAnimation](class(self).GetAnimations(gc))
}

func (self Go) SetAnimations(value gd.ArrayOf[gdclass.GLTFAnimation]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAnimations(value)
}

func (self Go) HandleBinaryImage() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetHandleBinaryImage()))
}

func (self Go) SetHandleBinaryImage(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHandleBinaryImage(gd.Int(value))
}

func (self Go) BakeFps() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetBakeFps()))
}

func (self Go) SetBakeFps(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBakeFps(gd.Float(value))
}

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
func (self class) GetNodes(ctx gd.Lifetime) gd.ArrayOf[gdclass.GLTFNode] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.GLTFNode](ret)
}
/*
Sets the [GLTFNode]s in the state. These are the nodes that [member GLTFNode.children] and [member root_nodes] refer to. Some of the nodes set here may not be generated in the Godot scene, or may generate multiple Godot scene nodes.
*/
//go:nosplit
func (self class) SetNodes(nodes gd.ArrayOf[gdclass.GLTFNode])  {
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
func (self class) GetBufferViews(ctx gd.Lifetime) gd.ArrayOf[gdclass.GLTFBufferView] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_buffer_views, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.GLTFBufferView](ret)
}
//go:nosplit
func (self class) SetBufferViews(buffer_views gd.ArrayOf[gdclass.GLTFBufferView])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(buffer_views))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_set_buffer_views, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAccessors(ctx gd.Lifetime) gd.ArrayOf[gdclass.GLTFAccessor] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_accessors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.GLTFAccessor](ret)
}
//go:nosplit
func (self class) SetAccessors(accessors gd.ArrayOf[gdclass.GLTFAccessor])  {
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
func (self class) GetMeshes(ctx gd.Lifetime) gd.ArrayOf[gdclass.GLTFMesh] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_meshes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.GLTFMesh](ret)
}
/*
Sets the [GLTFMesh]es in the state. These are the meshes that the [member GLTFNode.mesh] index refers to.
*/
//go:nosplit
func (self class) SetMeshes(meshes gd.ArrayOf[gdclass.GLTFMesh])  {
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
func (self class) GetAnimationPlayer(ctx gd.Lifetime, idx gd.Int) gdclass.AnimationPlayer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_animation_player, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.AnimationPlayer
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetMaterials(ctx gd.Lifetime) gd.ArrayOf[gdclass.Material] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_materials, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.Material](ret)
}
//go:nosplit
func (self class) SetMaterials(materials gd.ArrayOf[gdclass.Material])  {
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
func (self class) GetTextures(ctx gd.Lifetime) gd.ArrayOf[gdclass.GLTFTexture] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_textures, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.GLTFTexture](ret)
}
//go:nosplit
func (self class) SetTextures(textures gd.ArrayOf[gdclass.GLTFTexture])  {
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
func (self class) GetTextureSamplers(ctx gd.Lifetime) gd.ArrayOf[gdclass.GLTFTextureSampler] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_texture_samplers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.GLTFTextureSampler](ret)
}
/*
Sets the array of texture samplers that are used by the textures contained in the GLTF.
*/
//go:nosplit
func (self class) SetTextureSamplers(texture_samplers gd.ArrayOf[gdclass.GLTFTextureSampler])  {
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
func (self class) GetImages(ctx gd.Lifetime) gd.ArrayOf[gdclass.Texture2D] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_images, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.Texture2D](ret)
}
/*
Sets the images in the state stored as an array of [Texture2D]s. This can be used during export. These are the images that the [member GLTFTexture.src_image] index refers to.
*/
//go:nosplit
func (self class) SetImages(images gd.ArrayOf[gdclass.Texture2D])  {
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
func (self class) GetSkins(ctx gd.Lifetime) gd.ArrayOf[gdclass.GLTFSkin] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_skins, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.GLTFSkin](ret)
}
/*
Sets the [GLTFSkin]s in the state. These are the skins that the [member GLTFNode.skin] index refers to.
*/
//go:nosplit
func (self class) SetSkins(skins gd.ArrayOf[gdclass.GLTFSkin])  {
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
func (self class) GetCameras(ctx gd.Lifetime) gd.ArrayOf[gdclass.GLTFCamera] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_cameras, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.GLTFCamera](ret)
}
/*
Sets the [GLTFCamera]s in the state. These are the cameras that the [member GLTFNode.camera] index refers to.
*/
//go:nosplit
func (self class) SetCameras(cameras gd.ArrayOf[gdclass.GLTFCamera])  {
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
func (self class) GetLights(ctx gd.Lifetime) gd.ArrayOf[gdclass.GLTFLight] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_lights, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.GLTFLight](ret)
}
/*
Sets the [GLTFLight]s in the state. These are the lights that the [member GLTFNode.light] index refers to.
*/
//go:nosplit
func (self class) SetLights(lights gd.ArrayOf[gdclass.GLTFLight])  {
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
func (self class) GetSkeletons(ctx gd.Lifetime) gd.ArrayOf[gdclass.GLTFSkeleton] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_skeletons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.GLTFSkeleton](ret)
}
/*
Sets the [GLTFSkeleton]s in the state. These are the skeletons that the [member GLTFNode.skeleton] index refers to.
*/
//go:nosplit
func (self class) SetSkeletons(skeletons gd.ArrayOf[gdclass.GLTFSkeleton])  {
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
func (self class) GetAnimations(ctx gd.Lifetime) gd.ArrayOf[gdclass.GLTFAnimation] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_animations, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.GLTFAnimation](ret)
}
/*
Sets the [GLTFAnimation]s in the state. When importing, these will be generated as animations in an [AnimationPlayer] node. When exporting, these will be generated from Godot [AnimationPlayer] nodes.
*/
//go:nosplit
func (self class) SetAnimations(animations gd.ArrayOf[gdclass.GLTFAnimation])  {
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
func (self class) GetSceneNode(ctx gd.Lifetime, idx gd.Int) gdclass.Node {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFState.Bind_get_scene_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Node
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the index of the [GLTFNode] corresponding to this Godot scene node. This is the inverse of [method get_scene_node]. Useful during the export process.
[b]Note:[/b] Not every Godot scene node will have a corresponding [GLTFNode], and not every [GLTFNode] will have a scene node generated. If there is no [GLTFNode] index for this scene node, [code]-1[/code] is returned.
*/
//go:nosplit
func (self class) GetNodeIndex(scene_node gdclass.Node) gd.Int {
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
func (self class) AsGLTFState() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGLTFState() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("GLTFState", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
