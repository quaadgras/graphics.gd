// Package GLTFState provides methods for working with GLTFState object instances.
package GLTFState

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
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/classdb/Resource"
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
var _ Dictionary.Any
var _ RID.Any

/*
Contains all nodes and resources of a GLTF file. This is used by [GLTFDocument] as data storage, which allows [GLTFDocument] and all [GLTFDocumentExtension] classes to remain stateless.
GLTFState can be populated by [GLTFDocument] reading a file or by converting a Godot scene. Then the data can either be used to create a Godot scene or save to a GLTF file. The code that converts to/from a Godot scene can be intercepted at arbitrary points by [GLTFDocumentExtension] classes. This allows for custom data to be stored in the GLTF file or for custom data to be converted to/from Godot nodes.
*/
type Instance [1]gdclass.GLTFState

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGLTFState() Instance
}

/*
Appends an extension to the list of extensions used by this GLTF file during serialization. If [param required] is true, the extension will also be added to the list of required extensions. Do not run this in [method GLTFDocumentExtension._export_post], as that stage is too late to add extensions. The final list is sorted alphabetically.
*/
func (self Instance) AddUsedExtension(extension_name string, required bool) { //gd:GLTFState.add_used_extension
	class(self).AddUsedExtension(gd.NewString(extension_name), required)
}

/*
Appends the given byte array data to the buffers and creates a [GLTFBufferView] for it. The index of the destination [GLTFBufferView] is returned. If [param deduplication] is true, the buffers will first be searched for duplicate data, otherwise new bytes will always be appended.
*/
func (self Instance) AppendDataToBuffers(data []byte, deduplication bool) int { //gd:GLTFState.append_data_to_buffers
	return int(int(class(self).AppendDataToBuffers(gd.NewPackedByteSlice(data), deduplication)))
}

/*
Returns the number of [AnimationPlayer] nodes in this [GLTFState]. These nodes are only used during the export process when converting Godot [AnimationPlayer] nodes to GLTF animations.
*/
func (self Instance) GetAnimationPlayersCount(idx int) int { //gd:GLTFState.get_animation_players_count
	return int(int(class(self).GetAnimationPlayersCount(gd.Int(idx))))
}

/*
Returns the [AnimationPlayer] node with the given index. These nodes are only used during the export process when converting Godot [AnimationPlayer] nodes to GLTF animations.
*/
func (self Instance) GetAnimationPlayer(idx int) [1]gdclass.AnimationPlayer { //gd:GLTFState.get_animation_player
	return [1]gdclass.AnimationPlayer(class(self).GetAnimationPlayer(gd.Int(idx)))
}

/*
Returns the Godot scene node that corresponds to the same index as the [GLTFNode] it was generated from. This is the inverse of [method get_node_index]. Useful during the import process.
[b]Note:[/b] Not every [GLTFNode] will have a scene node generated, and not every generated scene node will have a corresponding [GLTFNode]. If there is no scene node for this [GLTFNode] index, [code]null[/code] is returned.
*/
func (self Instance) GetSceneNode(idx int) [1]gdclass.Node { //gd:GLTFState.get_scene_node
	return [1]gdclass.Node(class(self).GetSceneNode(gd.Int(idx)))
}

/*
Returns the index of the [GLTFNode] corresponding to this Godot scene node. This is the inverse of [method get_scene_node]. Useful during the export process.
[b]Note:[/b] Not every Godot scene node will have a corresponding [GLTFNode], and not every [GLTFNode] will have a scene node generated. If there is no [GLTFNode] index for this scene node, [code]-1[/code] is returned.
*/
func (self Instance) GetNodeIndex(scene_node [1]gdclass.Node) int { //gd:GLTFState.get_node_index
	return int(int(class(self).GetNodeIndex(scene_node)))
}

/*
Gets additional arbitrary data in this [GLTFState] instance. This can be used to keep per-file state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the return value can be anything you set. If nothing was set, the return value is null.
*/
func (self Instance) GetAdditionalData(extension_name string) any { //gd:GLTFState.get_additional_data
	return any(class(self).GetAdditionalData(gd.NewStringName(extension_name)).Interface())
}

/*
Sets additional arbitrary data in this [GLTFState] instance. This can be used to keep per-file state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The first argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the second argument can be anything you want.
*/
func (self Instance) SetAdditionalData(extension_name string, additional_data any) { //gd:GLTFState.set_additional_data
	class(self).SetAdditionalData(gd.NewStringName(extension_name), gd.NewVariant(additional_data))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GLTFState

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFState"))
	casted := Instance{*(*gdclass.GLTFState)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Json() map[any]any {
	return map[any]any(gd.DictionaryAs[map[any]any](class(self).GetJson()))
}

func (self Instance) SetJson(value map[any]any) {
	class(self).SetJson(gd.DictionaryFromMap(value))
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

func (self Instance) Nodes() [][1]gdclass.GLTFNode {
	return [][1]gdclass.GLTFNode(gd.ArrayAs[[][1]gdclass.GLTFNode](gd.InternalArray(class(self).GetNodes())))
}

func (self Instance) SetNodes(value [][1]gdclass.GLTFNode) {
	class(self).SetNodes(gd.ArrayFromSlice[Array.Contains[[1]gdclass.GLTFNode]](value))
}

func (self Instance) Buffers() [][]byte {
	return [][]byte(gd.ArrayAs[[][]byte](gd.InternalArray(class(self).GetBuffers())))
}

func (self Instance) SetBuffers(value [][]byte) {
	class(self).SetBuffers(gd.ArrayFromSlice[Array.Contains[gd.PackedByteArray]](value))
}

func (self Instance) BufferViews() [][1]gdclass.GLTFBufferView {
	return [][1]gdclass.GLTFBufferView(gd.ArrayAs[[][1]gdclass.GLTFBufferView](gd.InternalArray(class(self).GetBufferViews())))
}

func (self Instance) SetBufferViews(value [][1]gdclass.GLTFBufferView) {
	class(self).SetBufferViews(gd.ArrayFromSlice[Array.Contains[[1]gdclass.GLTFBufferView]](value))
}

func (self Instance) Accessors() [][1]gdclass.GLTFAccessor {
	return [][1]gdclass.GLTFAccessor(gd.ArrayAs[[][1]gdclass.GLTFAccessor](gd.InternalArray(class(self).GetAccessors())))
}

func (self Instance) SetAccessors(value [][1]gdclass.GLTFAccessor) {
	class(self).SetAccessors(gd.ArrayFromSlice[Array.Contains[[1]gdclass.GLTFAccessor]](value))
}

func (self Instance) Meshes() [][1]gdclass.GLTFMesh {
	return [][1]gdclass.GLTFMesh(gd.ArrayAs[[][1]gdclass.GLTFMesh](gd.InternalArray(class(self).GetMeshes())))
}

func (self Instance) SetMeshes(value [][1]gdclass.GLTFMesh) {
	class(self).SetMeshes(gd.ArrayFromSlice[Array.Contains[[1]gdclass.GLTFMesh]](value))
}

func (self Instance) Materials() [][1]gdclass.Material {
	return [][1]gdclass.Material(gd.ArrayAs[[][1]gdclass.Material](gd.InternalArray(class(self).GetMaterials())))
}

func (self Instance) SetMaterials(value [][1]gdclass.Material) {
	class(self).SetMaterials(gd.ArrayFromSlice[Array.Contains[[1]gdclass.Material]](value))
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

func (self Instance) Textures() [][1]gdclass.GLTFTexture {
	return [][1]gdclass.GLTFTexture(gd.ArrayAs[[][1]gdclass.GLTFTexture](gd.InternalArray(class(self).GetTextures())))
}

func (self Instance) SetTextures(value [][1]gdclass.GLTFTexture) {
	class(self).SetTextures(gd.ArrayFromSlice[Array.Contains[[1]gdclass.GLTFTexture]](value))
}

func (self Instance) TextureSamplers() [][1]gdclass.GLTFTextureSampler {
	return [][1]gdclass.GLTFTextureSampler(gd.ArrayAs[[][1]gdclass.GLTFTextureSampler](gd.InternalArray(class(self).GetTextureSamplers())))
}

func (self Instance) SetTextureSamplers(value [][1]gdclass.GLTFTextureSampler) {
	class(self).SetTextureSamplers(gd.ArrayFromSlice[Array.Contains[[1]gdclass.GLTFTextureSampler]](value))
}

func (self Instance) Images() [][1]gdclass.Texture2D {
	return [][1]gdclass.Texture2D(gd.ArrayAs[[][1]gdclass.Texture2D](gd.InternalArray(class(self).GetImages())))
}

func (self Instance) SetImages(value [][1]gdclass.Texture2D) {
	class(self).SetImages(gd.ArrayFromSlice[Array.Contains[[1]gdclass.Texture2D]](value))
}

func (self Instance) Skins() [][1]gdclass.GLTFSkin {
	return [][1]gdclass.GLTFSkin(gd.ArrayAs[[][1]gdclass.GLTFSkin](gd.InternalArray(class(self).GetSkins())))
}

func (self Instance) SetSkins(value [][1]gdclass.GLTFSkin) {
	class(self).SetSkins(gd.ArrayFromSlice[Array.Contains[[1]gdclass.GLTFSkin]](value))
}

func (self Instance) Cameras() [][1]gdclass.GLTFCamera {
	return [][1]gdclass.GLTFCamera(gd.ArrayAs[[][1]gdclass.GLTFCamera](gd.InternalArray(class(self).GetCameras())))
}

func (self Instance) SetCameras(value [][1]gdclass.GLTFCamera) {
	class(self).SetCameras(gd.ArrayFromSlice[Array.Contains[[1]gdclass.GLTFCamera]](value))
}

func (self Instance) Lights() [][1]gdclass.GLTFLight {
	return [][1]gdclass.GLTFLight(gd.ArrayAs[[][1]gdclass.GLTFLight](gd.InternalArray(class(self).GetLights())))
}

func (self Instance) SetLights(value [][1]gdclass.GLTFLight) {
	class(self).SetLights(gd.ArrayFromSlice[Array.Contains[[1]gdclass.GLTFLight]](value))
}

func (self Instance) UniqueNames() []string {
	return []string(gd.ArrayAs[[]string](gd.InternalArray(class(self).GetUniqueNames())))
}

func (self Instance) SetUniqueNames(value []string) {
	class(self).SetUniqueNames(gd.ArrayFromSlice[Array.Contains[gd.String]](value))
}

func (self Instance) UniqueAnimationNames() []string {
	return []string(gd.ArrayAs[[]string](gd.InternalArray(class(self).GetUniqueAnimationNames())))
}

func (self Instance) SetUniqueAnimationNames(value []string) {
	class(self).SetUniqueAnimationNames(gd.ArrayFromSlice[Array.Contains[gd.String]](value))
}

func (self Instance) Skeletons() [][1]gdclass.GLTFSkeleton {
	return [][1]gdclass.GLTFSkeleton(gd.ArrayAs[[][1]gdclass.GLTFSkeleton](gd.InternalArray(class(self).GetSkeletons())))
}

func (self Instance) SetSkeletons(value [][1]gdclass.GLTFSkeleton) {
	class(self).SetSkeletons(gd.ArrayFromSlice[Array.Contains[[1]gdclass.GLTFSkeleton]](value))
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

func (self Instance) Animations() [][1]gdclass.GLTFAnimation {
	return [][1]gdclass.GLTFAnimation(gd.ArrayAs[[][1]gdclass.GLTFAnimation](gd.InternalArray(class(self).GetAnimations())))
}

func (self Instance) SetAnimations(value [][1]gdclass.GLTFAnimation) {
	class(self).SetAnimations(gd.ArrayFromSlice[Array.Contains[[1]gdclass.GLTFAnimation]](value))
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
func (self class) AddUsedExtension(extension_name gd.String, required bool) { //gd:GLTFState.add_used_extension
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension_name))
	callframe.Arg(frame, required)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_add_used_extension, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Appends the given byte array data to the buffers and creates a [GLTFBufferView] for it. The index of the destination [GLTFBufferView] is returned. If [param deduplication] is true, the buffers will first be searched for duplicate data, otherwise new bytes will always be appended.
*/
//go:nosplit
func (self class) AppendDataToBuffers(data gd.PackedByteArray, deduplication bool) gd.Int { //gd:GLTFState.append_data_to_buffers
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(data))
	callframe.Arg(frame, deduplication)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_append_data_to_buffers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetJson() Dictionary.Any { //gd:GLTFState.get_json
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_json, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJson(json Dictionary.Any) { //gd:GLTFState.set_json
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalDictionary(json)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_json, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMajorVersion() gd.Int { //gd:GLTFState.get_major_version
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_major_version, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMajorVersion(major_version gd.Int) { //gd:GLTFState.set_major_version
	var frame = callframe.New()
	callframe.Arg(frame, major_version)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_major_version, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinorVersion() gd.Int { //gd:GLTFState.get_minor_version
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_minor_version, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinorVersion(minor_version gd.Int) { //gd:GLTFState.set_minor_version
	var frame = callframe.New()
	callframe.Arg(frame, minor_version)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_minor_version, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCopyright() gd.String { //gd:GLTFState.get_copyright
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_copyright, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCopyright(copyright gd.String) { //gd:GLTFState.set_copyright
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(copyright))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_copyright, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlbData() gd.PackedByteArray { //gd:GLTFState.get_glb_data
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_glb_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlbData(glb_data gd.PackedByteArray) { //gd:GLTFState.set_glb_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(glb_data))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_glb_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUseNamedSkinBinds() bool { //gd:GLTFState.get_use_named_skin_binds
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_use_named_skin_binds, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseNamedSkinBinds(use_named_skin_binds bool) { //gd:GLTFState.set_use_named_skin_binds
	var frame = callframe.New()
	callframe.Arg(frame, use_named_skin_binds)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_use_named_skin_binds, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an array of all [GLTFNode]s in the GLTF file. These are the nodes that [member GLTFNode.children] and [member root_nodes] refer to. This includes nodes that may not be generated in the Godot scene, or nodes that may generate multiple Godot scene nodes.
*/
//go:nosplit
func (self class) GetNodes() Array.Contains[[1]gdclass.GLTFNode] { //gd:GLTFState.get_nodes
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_nodes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.GLTFNode]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the [GLTFNode]s in the state. These are the nodes that [member GLTFNode.children] and [member root_nodes] refer to. Some of the nodes set here may not be generated in the Godot scene, or may generate multiple Godot scene nodes.
*/
//go:nosplit
func (self class) SetNodes(nodes Array.Contains[[1]gdclass.GLTFNode]) { //gd:GLTFState.set_nodes
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(nodes)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_nodes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBuffers() Array.Contains[gd.PackedByteArray] { //gd:GLTFState.get_buffers
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_buffers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.PackedByteArray]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBuffers(buffers Array.Contains[gd.PackedByteArray]) { //gd:GLTFState.set_buffers
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(buffers)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_buffers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBufferViews() Array.Contains[[1]gdclass.GLTFBufferView] { //gd:GLTFState.get_buffer_views
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_buffer_views, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.GLTFBufferView]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBufferViews(buffer_views Array.Contains[[1]gdclass.GLTFBufferView]) { //gd:GLTFState.set_buffer_views
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(buffer_views)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_buffer_views, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAccessors() Array.Contains[[1]gdclass.GLTFAccessor] { //gd:GLTFState.get_accessors
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_accessors, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.GLTFAccessor]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAccessors(accessors Array.Contains[[1]gdclass.GLTFAccessor]) { //gd:GLTFState.set_accessors
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(accessors)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_accessors, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an array of all [GLTFMesh]es in the GLTF file. These are the meshes that the [member GLTFNode.mesh] index refers to.
*/
//go:nosplit
func (self class) GetMeshes() Array.Contains[[1]gdclass.GLTFMesh] { //gd:GLTFState.get_meshes
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_meshes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.GLTFMesh]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the [GLTFMesh]es in the state. These are the meshes that the [member GLTFNode.mesh] index refers to.
*/
//go:nosplit
func (self class) SetMeshes(meshes Array.Contains[[1]gdclass.GLTFMesh]) { //gd:GLTFState.set_meshes
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(meshes)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_meshes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the number of [AnimationPlayer] nodes in this [GLTFState]. These nodes are only used during the export process when converting Godot [AnimationPlayer] nodes to GLTF animations.
*/
//go:nosplit
func (self class) GetAnimationPlayersCount(idx gd.Int) gd.Int { //gd:GLTFState.get_animation_players_count
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_animation_players_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [AnimationPlayer] node with the given index. These nodes are only used during the export process when converting Godot [AnimationPlayer] nodes to GLTF animations.
*/
//go:nosplit
func (self class) GetAnimationPlayer(idx gd.Int) [1]gdclass.AnimationPlayer { //gd:GLTFState.get_animation_player
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_animation_player, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.AnimationPlayer{gd.PointerMustAssertInstanceID[gdclass.AnimationPlayer](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetMaterials() Array.Contains[[1]gdclass.Material] { //gd:GLTFState.get_materials
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_materials, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.Material]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaterials(materials Array.Contains[[1]gdclass.Material]) { //gd:GLTFState.set_materials
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(materials)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_materials, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSceneName() gd.String { //gd:GLTFState.get_scene_name
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_scene_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSceneName(scene_name gd.String) { //gd:GLTFState.set_scene_name
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scene_name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_scene_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBasePath() gd.String { //gd:GLTFState.get_base_path
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_base_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBasePath(base_path gd.String) { //gd:GLTFState.set_base_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(base_path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_base_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFilename() gd.String { //gd:GLTFState.get_filename
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_filename, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFilename(filename gd.String) { //gd:GLTFState.set_filename
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(filename))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_filename, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRootNodes() gd.PackedInt32Array { //gd:GLTFState.get_root_nodes
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_root_nodes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRootNodes(root_nodes gd.PackedInt32Array) { //gd:GLTFState.set_root_nodes
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(root_nodes))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_root_nodes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextures() Array.Contains[[1]gdclass.GLTFTexture] { //gd:GLTFState.get_textures
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_textures, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.GLTFTexture]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextures(textures Array.Contains[[1]gdclass.GLTFTexture]) { //gd:GLTFState.set_textures
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(textures)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_textures, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Retrieves the array of texture samplers that are used by the textures contained in the GLTF.
*/
//go:nosplit
func (self class) GetTextureSamplers() Array.Contains[[1]gdclass.GLTFTextureSampler] { //gd:GLTFState.get_texture_samplers
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_texture_samplers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.GLTFTextureSampler]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the array of texture samplers that are used by the textures contained in the GLTF.
*/
//go:nosplit
func (self class) SetTextureSamplers(texture_samplers Array.Contains[[1]gdclass.GLTFTextureSampler]) { //gd:GLTFState.set_texture_samplers
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(texture_samplers)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_texture_samplers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Gets the images of the GLTF file as an array of [Texture2D]s. These are the images that the [member GLTFTexture.src_image] index refers to.
*/
//go:nosplit
func (self class) GetImages() Array.Contains[[1]gdclass.Texture2D] { //gd:GLTFState.get_images
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_images, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.Texture2D]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the images in the state stored as an array of [Texture2D]s. This can be used during export. These are the images that the [member GLTFTexture.src_image] index refers to.
*/
//go:nosplit
func (self class) SetImages(images Array.Contains[[1]gdclass.Texture2D]) { //gd:GLTFState.set_images
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(images)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_images, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an array of all [GLTFSkin]s in the GLTF file. These are the skins that the [member GLTFNode.skin] index refers to.
*/
//go:nosplit
func (self class) GetSkins() Array.Contains[[1]gdclass.GLTFSkin] { //gd:GLTFState.get_skins
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_skins, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.GLTFSkin]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the [GLTFSkin]s in the state. These are the skins that the [member GLTFNode.skin] index refers to.
*/
//go:nosplit
func (self class) SetSkins(skins Array.Contains[[1]gdclass.GLTFSkin]) { //gd:GLTFState.set_skins
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(skins)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_skins, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an array of all [GLTFCamera]s in the GLTF file. These are the cameras that the [member GLTFNode.camera] index refers to.
*/
//go:nosplit
func (self class) GetCameras() Array.Contains[[1]gdclass.GLTFCamera] { //gd:GLTFState.get_cameras
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_cameras, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.GLTFCamera]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the [GLTFCamera]s in the state. These are the cameras that the [member GLTFNode.camera] index refers to.
*/
//go:nosplit
func (self class) SetCameras(cameras Array.Contains[[1]gdclass.GLTFCamera]) { //gd:GLTFState.set_cameras
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(cameras)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_cameras, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an array of all [GLTFLight]s in the GLTF file. These are the lights that the [member GLTFNode.light] index refers to.
*/
//go:nosplit
func (self class) GetLights() Array.Contains[[1]gdclass.GLTFLight] { //gd:GLTFState.get_lights
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_lights, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.GLTFLight]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the [GLTFLight]s in the state. These are the lights that the [member GLTFNode.light] index refers to.
*/
//go:nosplit
func (self class) SetLights(lights Array.Contains[[1]gdclass.GLTFLight]) { //gd:GLTFState.set_lights
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(lights)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_lights, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an array of unique node names. This is used in both the import process and export process.
*/
//go:nosplit
func (self class) GetUniqueNames() Array.Contains[gd.String] { //gd:GLTFState.get_unique_names
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_unique_names, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.String]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the unique node names in the state. This is used in both the import process and export process.
*/
//go:nosplit
func (self class) SetUniqueNames(unique_names Array.Contains[gd.String]) { //gd:GLTFState.set_unique_names
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(unique_names)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_unique_names, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an array of unique animation names. This is only used during the import process.
*/
//go:nosplit
func (self class) GetUniqueAnimationNames() Array.Contains[gd.String] { //gd:GLTFState.get_unique_animation_names
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_unique_animation_names, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.String]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the unique animation names in the state. This is only used during the import process.
*/
//go:nosplit
func (self class) SetUniqueAnimationNames(unique_animation_names Array.Contains[gd.String]) { //gd:GLTFState.set_unique_animation_names
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(unique_animation_names)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_unique_animation_names, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an array of all [GLTFSkeleton]s in the GLTF file. These are the skeletons that the [member GLTFNode.skeleton] index refers to.
*/
//go:nosplit
func (self class) GetSkeletons() Array.Contains[[1]gdclass.GLTFSkeleton] { //gd:GLTFState.get_skeletons
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_skeletons, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.GLTFSkeleton]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the [GLTFSkeleton]s in the state. These are the skeletons that the [member GLTFNode.skeleton] index refers to.
*/
//go:nosplit
func (self class) SetSkeletons(skeletons Array.Contains[[1]gdclass.GLTFSkeleton]) { //gd:GLTFState.set_skeletons
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(skeletons)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_skeletons, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCreateAnimations() bool { //gd:GLTFState.get_create_animations
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_create_animations, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCreateAnimations(create_animations bool) { //gd:GLTFState.set_create_animations
	var frame = callframe.New()
	callframe.Arg(frame, create_animations)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_create_animations, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetImportAsSkeletonBones() bool { //gd:GLTFState.get_import_as_skeleton_bones
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_import_as_skeleton_bones, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetImportAsSkeletonBones(import_as_skeleton_bones bool) { //gd:GLTFState.set_import_as_skeleton_bones
	var frame = callframe.New()
	callframe.Arg(frame, import_as_skeleton_bones)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_import_as_skeleton_bones, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an array of all [GLTFAnimation]s in the GLTF file. When importing, these will be generated as animations in an [AnimationPlayer] node. When exporting, these will be generated from Godot [AnimationPlayer] nodes.
*/
//go:nosplit
func (self class) GetAnimations() Array.Contains[[1]gdclass.GLTFAnimation] { //gd:GLTFState.get_animations
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_animations, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.GLTFAnimation]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the [GLTFAnimation]s in the state. When importing, these will be generated as animations in an [AnimationPlayer] node. When exporting, these will be generated from Godot [AnimationPlayer] nodes.
*/
//go:nosplit
func (self class) SetAnimations(animations Array.Contains[[1]gdclass.GLTFAnimation]) { //gd:GLTFState.set_animations
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(animations)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_animations, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the Godot scene node that corresponds to the same index as the [GLTFNode] it was generated from. This is the inverse of [method get_node_index]. Useful during the import process.
[b]Note:[/b] Not every [GLTFNode] will have a scene node generated, and not every generated scene node will have a corresponding [GLTFNode]. If there is no scene node for this [GLTFNode] index, [code]null[/code] is returned.
*/
//go:nosplit
func (self class) GetSceneNode(idx gd.Int) [1]gdclass.Node { //gd:GLTFState.get_scene_node
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_scene_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerMustAssertInstanceID[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the index of the [GLTFNode] corresponding to this Godot scene node. This is the inverse of [method get_scene_node]. Useful during the export process.
[b]Note:[/b] Not every Godot scene node will have a corresponding [GLTFNode], and not every [GLTFNode] will have a scene node generated. If there is no [GLTFNode] index for this scene node, [code]-1[/code] is returned.
*/
//go:nosplit
func (self class) GetNodeIndex(scene_node [1]gdclass.Node) gd.Int { //gd:GLTFState.get_node_index
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scene_node[0])[0])
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_node_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets additional arbitrary data in this [GLTFState] instance. This can be used to keep per-file state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the return value can be anything you set. If nothing was set, the return value is null.
*/
//go:nosplit
func (self class) GetAdditionalData(extension_name gd.StringName) gd.Variant { //gd:GLTFState.get_additional_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension_name))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_additional_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets additional arbitrary data in this [GLTFState] instance. This can be used to keep per-file state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The first argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the second argument can be anything you want.
*/
//go:nosplit
func (self class) SetAdditionalData(extension_name gd.StringName, additional_data gd.Variant) { //gd:GLTFState.set_additional_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension_name))
	callframe.Arg(frame, pointers.Get(additional_data))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_additional_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHandleBinaryImage() gd.Int { //gd:GLTFState.get_handle_binary_image
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_handle_binary_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHandleBinaryImage(method gd.Int) { //gd:GLTFState.set_handle_binary_image
	var frame = callframe.New()
	callframe.Arg(frame, method)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_handle_binary_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetBakeFps(value gd.Float) { //gd:GLTFState.set_bake_fps
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_set_bake_fps, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBakeFps() gd.Float { //gd:GLTFState.get_bake_fps
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFState.Bind_get_bake_fps, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("GLTFState", func(ptr gd.Object) any { return [1]gdclass.GLTFState{*(*gdclass.GLTFState)(unsafe.Pointer(&ptr))} })
}
