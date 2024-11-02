package GLTFState

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"
import "grow.graphics/gd/variant/Dictionary"
import "grow.graphics/gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Contains all nodes and resources of a GLTF file. This is used by [GLTFDocument] as data storage, which allows [GLTFDocument] and all [GLTFDocumentExtension] classes to remain stateless.
GLTFState can be populated by [GLTFDocument] reading a file or by converting a Godot scene. Then the data can either be used to create a Godot scene or save to a GLTF file. The code that converts to/from a Godot scene can be intercepted at arbitrary points by [GLTFDocumentExtension] classes. This allows for custom data to be stored in the GLTF file or for custom data to be converted to/from Godot nodes.
*/
type Instance [1]classdb.GLTFState

/*
Appends an extension to the list of extensions used by this GLTF file during serialization. If [param required] is true, the extension will also be added to the list of required extensions. Do not run this in [method GLTFDocumentExtension._export_post], as that stage is too late to add extensions. The final list is sorted alphabetically.
*/
func (self Instance) AddUsedExtension(extension_name string, required bool) {
	class(self).AddUsedExtension(gd.NewString(extension_name), required)
}

/*
Appends the given byte array data to the buffers and creates a [GLTFBufferView] for it. The index of the destination [GLTFBufferView] is returned. If [param deduplication] is true, the buffers will first be searched for duplicate data, otherwise new bytes will always be appended.
*/
func (self Instance) AppendDataToBuffers(data []byte, deduplication bool) int {
	return int(int(class(self).AppendDataToBuffers(gd.NewPackedByteSlice(data), deduplication)))
}

/*
Returns the number of [AnimationPlayer] nodes in this [GLTFState]. These nodes are only used during the export process when converting Godot [AnimationPlayer] nodes to GLTF animations.
*/
func (self Instance) GetAnimationPlayersCount(idx int) int {
	return int(int(class(self).GetAnimationPlayersCount(gd.Int(idx))))
}

/*
Returns the [AnimationPlayer] node with the given index. These nodes are only used during the export process when converting Godot [AnimationPlayer] nodes to GLTF animations.
*/
func (self Instance) GetAnimationPlayer(idx int) objects.AnimationPlayer {
	return objects.AnimationPlayer(class(self).GetAnimationPlayer(gd.Int(idx)))
}

/*
Returns the Godot scene node that corresponds to the same index as the [GLTFNode] it was generated from. This is the inverse of [method get_node_index]. Useful during the import process.
[b]Note:[/b] Not every [GLTFNode] will have a scene node generated, and not every generated scene node will have a corresponding [GLTFNode]. If there is no scene node for this [GLTFNode] index, [code]null[/code] is returned.
*/
func (self Instance) GetSceneNode(idx int) objects.Node {
	return objects.Node(class(self).GetSceneNode(gd.Int(idx)))
}

/*
Returns the index of the [GLTFNode] corresponding to this Godot scene node. This is the inverse of [method get_scene_node]. Useful during the export process.
[b]Note:[/b] Not every Godot scene node will have a corresponding [GLTFNode], and not every [GLTFNode] will have a scene node generated. If there is no [GLTFNode] index for this scene node, [code]-1[/code] is returned.
*/
func (self Instance) GetNodeIndex(scene_node objects.Node) int {
	return int(int(class(self).GetNodeIndex(scene_node)))
}

/*
Gets additional arbitrary data in this [GLTFState] instance. This can be used to keep per-file state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the return value can be anything you set. If nothing was set, the return value is null.
*/
func (self Instance) GetAdditionalData(extension_name string) any {
	return any(class(self).GetAdditionalData(gd.NewStringName(extension_name)).Interface())
}

/*
Sets additional arbitrary data in this [GLTFState] instance. This can be used to keep per-file state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The first argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the second argument can be anything you want.
*/
func (self Instance) SetAdditionalData(extension_name string, additional_data any) {
	class(self).SetAdditionalData(gd.NewStringName(extension_name), gd.NewVariant(additional_data))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.GLTFState

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFState"))
	return Instance{classdb.GLTFState(object)}
}

func (self Instance) Json() Dictionary.Any {
	return Dictionary.Any(class(self).GetJson())
}

func (self Instance) SetJson(value Dictionary.Any) {
	class(self).SetJson(value)
}

func (self Instance) MajorVersion() int {
	return int(int(class(self).GetMajorVersion()))
}

func (self Instance) SetMajorVersion(value int) {
	class(self).SetMajorVersion(gd.Int(value))
}

func (self Instance) MinorVersion() int {
	return int(int(class(self).GetMinorVersion()))
}

func (self Instance) SetMinorVersion(value int) {
	class(self).SetMinorVersion(gd.Int(value))
}

func (self Instance) Copyright() string {
	return string(class(self).GetCopyright().String())
}

func (self Instance) SetCopyright(value string) {
	class(self).SetCopyright(gd.NewString(value))
}

func (self Instance) GlbData() []byte {
	return []byte(class(self).GetGlbData().Bytes())
}

func (self Instance) SetGlbData(value []byte) {
	class(self).SetGlbData(gd.NewPackedByteSlice(value))
}

func (self Instance) UseNamedSkinBinds() bool {
	return bool(class(self).GetUseNamedSkinBinds())
}

func (self Instance) SetUseNamedSkinBinds(value bool) {
	class(self).SetUseNamedSkinBinds(value)
}

func (self Instance) Nodes() gd.Array {
	return gd.Array(class(self).GetNodes())
}

func (self Instance) SetNodes(value gd.Array) {
	class(self).SetNodes(value)
}

func (self Instance) Buffers() gd.Array {
	return gd.Array(class(self).GetBuffers())
}

func (self Instance) SetBuffers(value gd.Array) {
	class(self).SetBuffers(value)
}

func (self Instance) BufferViews() gd.Array {
	return gd.Array(class(self).GetBufferViews())
}

func (self Instance) SetBufferViews(value gd.Array) {
	class(self).SetBufferViews(value)
}

func (self Instance) Accessors() gd.Array {
	return gd.Array(class(self).GetAccessors())
}

func (self Instance) SetAccessors(value gd.Array) {
	class(self).SetAccessors(value)
}

func (self Instance) Meshes() gd.Array {
	return gd.Array(class(self).GetMeshes())
}

func (self Instance) SetMeshes(value gd.Array) {
	class(self).SetMeshes(value)
}

func (self Instance) Materials() gd.Array {
	return gd.Array(class(self).GetMaterials())
}

func (self Instance) SetMaterials(value gd.Array) {
	class(self).SetMaterials(value)
}

func (self Instance) SceneName() string {
	return string(class(self).GetSceneName().String())
}

func (self Instance) SetSceneName(value string) {
	class(self).SetSceneName(gd.NewString(value))
}

func (self Instance) BasePath() string {
	return string(class(self).GetBasePath().String())
}

func (self Instance) SetBasePath(value string) {
	class(self).SetBasePath(gd.NewString(value))
}

func (self Instance) Filename() string {
	return string(class(self).GetFilename().String())
}

func (self Instance) SetFilename(value string) {
	class(self).SetFilename(gd.NewString(value))
}

func (self Instance) RootNodes() []int32 {
	return []int32(class(self).GetRootNodes().AsSlice())
}

func (self Instance) SetRootNodes(value []int32) {
	class(self).SetRootNodes(gd.NewPackedInt32Slice(value))
}

func (self Instance) Textures() gd.Array {
	return gd.Array(class(self).GetTextures())
}

func (self Instance) SetTextures(value gd.Array) {
	class(self).SetTextures(value)
}

func (self Instance) TextureSamplers() gd.Array {
	return gd.Array(class(self).GetTextureSamplers())
}

func (self Instance) SetTextureSamplers(value gd.Array) {
	class(self).SetTextureSamplers(value)
}

func (self Instance) Images() gd.Array {
	return gd.Array(class(self).GetImages())
}

func (self Instance) SetImages(value gd.Array) {
	class(self).SetImages(value)
}

func (self Instance) Skins() gd.Array {
	return gd.Array(class(self).GetSkins())
}

func (self Instance) SetSkins(value gd.Array) {
	class(self).SetSkins(value)
}

func (self Instance) Cameras() gd.Array {
	return gd.Array(class(self).GetCameras())
}

func (self Instance) SetCameras(value gd.Array) {
	class(self).SetCameras(value)
}

func (self Instance) Lights() gd.Array {
	return gd.Array(class(self).GetLights())
}

func (self Instance) SetLights(value gd.Array) {
	class(self).SetLights(value)
}

func (self Instance) UniqueNames() gd.Array {
	return gd.Array(class(self).GetUniqueNames())
}

func (self Instance) SetUniqueNames(value gd.Array) {
	class(self).SetUniqueNames(value)
}

func (self Instance) UniqueAnimationNames() gd.Array {
	return gd.Array(class(self).GetUniqueAnimationNames())
}

func (self Instance) SetUniqueAnimationNames(value gd.Array) {
	class(self).SetUniqueAnimationNames(value)
}

func (self Instance) Skeletons() gd.Array {
	return gd.Array(class(self).GetSkeletons())
}

func (self Instance) SetSkeletons(value gd.Array) {
	class(self).SetSkeletons(value)
}

func (self Instance) CreateAnimations() bool {
	return bool(class(self).GetCreateAnimations())
}

func (self Instance) SetCreateAnimations(value bool) {
	class(self).SetCreateAnimations(value)
}

func (self Instance) ImportAsSkeletonBones() bool {
	return bool(class(self).GetImportAsSkeletonBones())
}

func (self Instance) SetImportAsSkeletonBones(value bool) {
	class(self).SetImportAsSkeletonBones(value)
}

func (self Instance) Animations() gd.Array {
	return gd.Array(class(self).GetAnimations())
}

func (self Instance) SetAnimations(value gd.Array) {
	class(self).SetAnimations(value)
}

func (self Instance) HandleBinaryImage() int {
	return int(int(class(self).GetHandleBinaryImage()))
}

func (self Instance) SetHandleBinaryImage(value int) {
	class(self).SetHandleBinaryImage(gd.Int(value))
}

func (self Instance) BakeFps() Float.X {
	return Float.X(Float.X(class(self).GetBakeFps()))
}

func (self Instance) SetBakeFps(value Float.X) {
	class(self).SetBakeFps(gd.Float(value))
}

/*
Appends an extension to the list of extensions used by this GLTF file during serialization. If [param required] is true, the extension will also be added to the list of required extensions. Do not run this in [method GLTFDocumentExtension._export_post], as that stage is too late to add extensions. The final list is sorted alphabetically.
*/
//go:nosplit
func (self class) AddUsedExtension(extension_name gd.String, required bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension_name))
	callframe.Arg(frame, required)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_add_used_extension, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Appends the given byte array data to the buffers and creates a [GLTFBufferView] for it. The index of the destination [GLTFBufferView] is returned. If [param deduplication] is true, the buffers will first be searched for duplicate data, otherwise new bytes will always be appended.
*/
//go:nosplit
func (self class) AppendDataToBuffers(data gd.PackedByteArray, deduplication bool) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(data))
	callframe.Arg(frame, deduplication)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_append_data_to_buffers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetJson() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_json, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJson(json gd.Dictionary) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(json))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_json, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMajorVersion() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_major_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMajorVersion(major_version gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, major_version)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_major_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinorVersion() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_minor_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinorVersion(minor_version gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, minor_version)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_minor_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCopyright() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_copyright, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCopyright(copyright gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(copyright))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_copyright, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlbData() gd.PackedByteArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_glb_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlbData(glb_data gd.PackedByteArray) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(glb_data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_glb_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetUseNamedSkinBinds() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_use_named_skin_binds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseNamedSkinBinds(use_named_skin_binds bool) {
	var frame = callframe.New()
	callframe.Arg(frame, use_named_skin_binds)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_use_named_skin_binds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns an array of all [GLTFNode]s in the GLTF file. These are the nodes that [member GLTFNode.children] and [member root_nodes] refer to. This includes nodes that may not be generated in the Godot scene, or nodes that may generate multiple Godot scene nodes.
*/
//go:nosplit
func (self class) GetNodes() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the [GLTFNode]s in the state. These are the nodes that [member GLTFNode.children] and [member root_nodes] refer to. Some of the nodes set here may not be generated in the Godot scene, or may generate multiple Godot scene nodes.
*/
//go:nosplit
func (self class) SetNodes(nodes gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(nodes))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBuffers() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_buffers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBuffers(buffers gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(buffers))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_buffers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBufferViews() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_buffer_views, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBufferViews(buffer_views gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(buffer_views))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_buffer_views, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAccessors() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_accessors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAccessors(accessors gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(accessors))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_accessors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns an array of all [GLTFMesh]es in the GLTF file. These are the meshes that the [member GLTFNode.mesh] index refers to.
*/
//go:nosplit
func (self class) GetMeshes() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_meshes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the [GLTFMesh]es in the state. These are the meshes that the [member GLTFNode.mesh] index refers to.
*/
//go:nosplit
func (self class) SetMeshes(meshes gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(meshes))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_meshes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the number of [AnimationPlayer] nodes in this [GLTFState]. These nodes are only used during the export process when converting Godot [AnimationPlayer] nodes to GLTF animations.
*/
//go:nosplit
func (self class) GetAnimationPlayersCount(idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_animation_players_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [AnimationPlayer] node with the given index. These nodes are only used during the export process when converting Godot [AnimationPlayer] nodes to GLTF animations.
*/
//go:nosplit
func (self class) GetAnimationPlayer(idx gd.Int) objects.AnimationPlayer {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_animation_player, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.AnimationPlayer{classdb.AnimationPlayer(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetMaterials() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_materials, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaterials(materials gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(materials))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_materials, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSceneName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_scene_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSceneName(scene_name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scene_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_scene_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBasePath() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_base_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBasePath(base_path gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(base_path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_base_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFilename() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_filename, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFilename(filename gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(filename))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_filename, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRootNodes() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_root_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRootNodes(root_nodes gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(root_nodes))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_root_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextures() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_textures, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextures(textures gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(textures))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_textures, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Retrieves the array of texture samplers that are used by the textures contained in the GLTF.
*/
//go:nosplit
func (self class) GetTextureSamplers() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_texture_samplers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the array of texture samplers that are used by the textures contained in the GLTF.
*/
//go:nosplit
func (self class) SetTextureSamplers(texture_samplers gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture_samplers))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_texture_samplers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Gets the images of the GLTF file as an array of [Texture2D]s. These are the images that the [member GLTFTexture.src_image] index refers to.
*/
//go:nosplit
func (self class) GetImages() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_images, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the images in the state stored as an array of [Texture2D]s. This can be used during export. These are the images that the [member GLTFTexture.src_image] index refers to.
*/
//go:nosplit
func (self class) SetImages(images gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(images))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_images, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns an array of all [GLTFSkin]s in the GLTF file. These are the skins that the [member GLTFNode.skin] index refers to.
*/
//go:nosplit
func (self class) GetSkins() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_skins, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the [GLTFSkin]s in the state. These are the skins that the [member GLTFNode.skin] index refers to.
*/
//go:nosplit
func (self class) SetSkins(skins gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(skins))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_skins, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns an array of all [GLTFCamera]s in the GLTF file. These are the cameras that the [member GLTFNode.camera] index refers to.
*/
//go:nosplit
func (self class) GetCameras() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_cameras, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the [GLTFCamera]s in the state. These are the cameras that the [member GLTFNode.camera] index refers to.
*/
//go:nosplit
func (self class) SetCameras(cameras gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(cameras))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_cameras, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns an array of all [GLTFLight]s in the GLTF file. These are the lights that the [member GLTFNode.light] index refers to.
*/
//go:nosplit
func (self class) GetLights() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_lights, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the [GLTFLight]s in the state. These are the lights that the [member GLTFNode.light] index refers to.
*/
//go:nosplit
func (self class) SetLights(lights gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(lights))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_lights, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns an array of unique node names. This is used in both the import process and export process.
*/
//go:nosplit
func (self class) GetUniqueNames() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_unique_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the unique node names in the state. This is used in both the import process and export process.
*/
//go:nosplit
func (self class) SetUniqueNames(unique_names gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(unique_names))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_unique_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns an array of unique animation names. This is only used during the import process.
*/
//go:nosplit
func (self class) GetUniqueAnimationNames() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_unique_animation_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the unique animation names in the state. This is only used during the import process.
*/
//go:nosplit
func (self class) SetUniqueAnimationNames(unique_animation_names gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(unique_animation_names))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_unique_animation_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns an array of all [GLTFSkeleton]s in the GLTF file. These are the skeletons that the [member GLTFNode.skeleton] index refers to.
*/
//go:nosplit
func (self class) GetSkeletons() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_skeletons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the [GLTFSkeleton]s in the state. These are the skeletons that the [member GLTFNode.skeleton] index refers to.
*/
//go:nosplit
func (self class) SetSkeletons(skeletons gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(skeletons))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_skeletons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCreateAnimations() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_create_animations, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCreateAnimations(create_animations bool) {
	var frame = callframe.New()
	callframe.Arg(frame, create_animations)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_create_animations, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetImportAsSkeletonBones() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_import_as_skeleton_bones, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetImportAsSkeletonBones(import_as_skeleton_bones bool) {
	var frame = callframe.New()
	callframe.Arg(frame, import_as_skeleton_bones)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_import_as_skeleton_bones, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns an array of all [GLTFAnimation]s in the GLTF file. When importing, these will be generated as animations in an [AnimationPlayer] node. When exporting, these will be generated from Godot [AnimationPlayer] nodes.
*/
//go:nosplit
func (self class) GetAnimations() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_animations, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the [GLTFAnimation]s in the state. When importing, these will be generated as animations in an [AnimationPlayer] node. When exporting, these will be generated from Godot [AnimationPlayer] nodes.
*/
//go:nosplit
func (self class) SetAnimations(animations gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(animations))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_animations, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the Godot scene node that corresponds to the same index as the [GLTFNode] it was generated from. This is the inverse of [method get_node_index]. Useful during the import process.
[b]Note:[/b] Not every [GLTFNode] will have a scene node generated, and not every generated scene node will have a corresponding [GLTFNode]. If there is no scene node for this [GLTFNode] index, [code]null[/code] is returned.
*/
//go:nosplit
func (self class) GetSceneNode(idx gd.Int) objects.Node {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_scene_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Node{classdb.Node(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the index of the [GLTFNode] corresponding to this Godot scene node. This is the inverse of [method get_scene_node]. Useful during the export process.
[b]Note:[/b] Not every Godot scene node will have a corresponding [GLTFNode], and not every [GLTFNode] will have a scene node generated. If there is no [GLTFNode] index for this scene node, [code]-1[/code] is returned.
*/
//go:nosplit
func (self class) GetNodeIndex(scene_node objects.Node) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scene_node[0])[0])
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_node_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets additional arbitrary data in this [GLTFState] instance. This can be used to keep per-file state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the return value can be anything you set. If nothing was set, the return value is null.
*/
//go:nosplit
func (self class) GetAdditionalData(extension_name gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension_name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_additional_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets additional arbitrary data in this [GLTFState] instance. This can be used to keep per-file state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The first argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the second argument can be anything you want.
*/
//go:nosplit
func (self class) SetAdditionalData(extension_name gd.StringName, additional_data gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension_name))
	callframe.Arg(frame, pointers.Get(additional_data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_additional_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHandleBinaryImage() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_handle_binary_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHandleBinaryImage(method gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, method)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_handle_binary_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetBakeFps(value gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_bake_fps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBakeFps() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_bake_fps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGLTFState() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGLTFState() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() { classdb.Register("GLTFState", func(ptr gd.Object) any { return classdb.GLTFState(ptr) }) }
