// Package GLTFDocument provides methods for working with GLTFDocument object instances.
package GLTFDocument

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
GLTFDocument supports reading data from a glTF file, buffer, or Godot scene. This data can then be written to the filesystem, buffer, or used to create a Godot scene.
All of the data in a GLTF scene is stored in the [GLTFState] class. GLTFDocument processes state objects, but does not contain any scene data itself. GLTFDocument has member variables to store export configuration settings such as the image format, but is otherwise stateless. Multiple scenes can be processed with the same settings using the same GLTFDocument object and different [GLTFState] objects.
GLTFDocument can be extended with arbitrary functionality by extending the [GLTFDocumentExtension] class and registering it with GLTFDocument via [method register_gltf_document_extension]. This allows for custom data to be imported and exported.
*/
type Instance [1]gdclass.GLTFDocument

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGLTFDocument() Instance
}

/*
Takes a path to a GLTF file and imports the data at that file path to the given [GLTFState] object through the [param state] parameter.
[b]Note:[/b] The [param base_path] tells [method append_from_file] where to find dependencies and can be empty.
*/
func (self Instance) AppendFromFile(path string, state [1]gdclass.GLTFState) error { //gd:GLTFDocument.append_from_file
	return error(gd.ToError(class(self).AppendFromFile(gd.NewString(path), state, gd.Int(0), gd.NewString(""))))
}

/*
Takes a [PackedByteArray] defining a GLTF and imports the data to the given [GLTFState] object through the [param state] parameter.
[b]Note:[/b] The [param base_path] tells [method append_from_buffer] where to find dependencies and can be empty.
*/
func (self Instance) AppendFromBuffer(bytes []byte, base_path string, state [1]gdclass.GLTFState) error { //gd:GLTFDocument.append_from_buffer
	return error(gd.ToError(class(self).AppendFromBuffer(gd.NewPackedByteSlice(bytes), gd.NewString(base_path), state, gd.Int(0))))
}

/*
Takes a Godot Engine scene node and exports it and its descendants to the given [GLTFState] object through the [param state] parameter.
*/
func (self Instance) AppendFromScene(node [1]gdclass.Node, state [1]gdclass.GLTFState) error { //gd:GLTFDocument.append_from_scene
	return error(gd.ToError(class(self).AppendFromScene(node, state, gd.Int(0))))
}

/*
Takes a [GLTFState] object through the [param state] parameter and returns a Godot Engine scene node.
The [param bake_fps] parameter overrides the bake_fps in [param state].
*/
func (self Instance) GenerateScene(state [1]gdclass.GLTFState) [1]gdclass.Node { //gd:GLTFDocument.generate_scene
	return [1]gdclass.Node(class(self).GenerateScene(state, gd.Float(30), false, true))
}

/*
Takes a [GLTFState] object through the [param state] parameter and returns a GLTF [PackedByteArray].
*/
func (self Instance) GenerateBuffer(state [1]gdclass.GLTFState) []byte { //gd:GLTFDocument.generate_buffer
	return []byte(class(self).GenerateBuffer(state).Bytes())
}

/*
Takes a [GLTFState] object through the [param state] parameter and writes a glTF file to the filesystem.
[b]Note:[/b] The extension of the glTF file determines if it is a .glb binary file or a .gltf text file.
*/
func (self Instance) WriteToFilesystem(state [1]gdclass.GLTFState, path string) error { //gd:GLTFDocument.write_to_filesystem
	return error(gd.ToError(class(self).WriteToFilesystem(state, gd.NewString(path))))
}

/*
Registers the given [GLTFDocumentExtension] instance with GLTFDocument. If [param first_priority] is true, this extension will be run first. Otherwise, it will be run last.
[b]Note:[/b] Like GLTFDocument itself, all GLTFDocumentExtension classes must be stateless in order to function properly. If you need to store data, use the [code]set_additional_data[/code] and [code]get_additional_data[/code] methods in [GLTFState] or [GLTFNode].
*/
func RegisterGltfDocumentExtension(extension [1]gdclass.GLTFDocumentExtension) { //gd:GLTFDocument.register_gltf_document_extension
	self := Instance{}
	class(self).RegisterGltfDocumentExtension(extension, false)
}

/*
Unregisters the given [GLTFDocumentExtension] instance.
*/
func UnregisterGltfDocumentExtension(extension [1]gdclass.GLTFDocumentExtension) { //gd:GLTFDocument.unregister_gltf_document_extension
	self := Instance{}
	class(self).UnregisterGltfDocumentExtension(extension)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GLTFDocument

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFDocument"))
	casted := Instance{*(*gdclass.GLTFDocument)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) ImageFormat() string {
	return string(class(self).GetImageFormat().String())
}

func (self Instance) SetImageFormat(value string) {
	class(self).SetImageFormat(gd.NewString(value))
}

func (self Instance) LossyQuality() Float.X {
	return Float.X(Float.X(class(self).GetLossyQuality()))
}

func (self Instance) SetLossyQuality(value Float.X) {
	class(self).SetLossyQuality(gd.Float(value))
}

func (self Instance) RootNodeMode() gdclass.GLTFDocumentRootNodeMode {
	return gdclass.GLTFDocumentRootNodeMode(class(self).GetRootNodeMode())
}

func (self Instance) SetRootNodeMode(value gdclass.GLTFDocumentRootNodeMode) {
	class(self).SetRootNodeMode(value)
}

//go:nosplit
func (self class) SetImageFormat(image_format gd.String) { //gd:GLTFDocument.set_image_format
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(image_format))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFDocument.Bind_set_image_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetImageFormat() gd.String { //gd:GLTFDocument.get_image_format
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFDocument.Bind_get_image_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLossyQuality(lossy_quality gd.Float) { //gd:GLTFDocument.set_lossy_quality
	var frame = callframe.New()
	callframe.Arg(frame, lossy_quality)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFDocument.Bind_set_lossy_quality, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLossyQuality() gd.Float { //gd:GLTFDocument.get_lossy_quality
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFDocument.Bind_get_lossy_quality, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRootNodeMode(root_node_mode gdclass.GLTFDocumentRootNodeMode) { //gd:GLTFDocument.set_root_node_mode
	var frame = callframe.New()
	callframe.Arg(frame, root_node_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFDocument.Bind_set_root_node_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRootNodeMode() gdclass.GLTFDocumentRootNodeMode { //gd:GLTFDocument.get_root_node_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GLTFDocumentRootNodeMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFDocument.Bind_get_root_node_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Takes a path to a GLTF file and imports the data at that file path to the given [GLTFState] object through the [param state] parameter.
[b]Note:[/b] The [param base_path] tells [method append_from_file] where to find dependencies and can be empty.
*/
//go:nosplit
func (self class) AppendFromFile(path gd.String, state [1]gdclass.GLTFState, flags gd.Int, base_path gd.String) gd.Error { //gd:GLTFDocument.append_from_file
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, pointers.Get(state[0])[0])
	callframe.Arg(frame, flags)
	callframe.Arg(frame, pointers.Get(base_path))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFDocument.Bind_append_from_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Takes a [PackedByteArray] defining a GLTF and imports the data to the given [GLTFState] object through the [param state] parameter.
[b]Note:[/b] The [param base_path] tells [method append_from_buffer] where to find dependencies and can be empty.
*/
//go:nosplit
func (self class) AppendFromBuffer(bytes gd.PackedByteArray, base_path gd.String, state [1]gdclass.GLTFState, flags gd.Int) gd.Error { //gd:GLTFDocument.append_from_buffer
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(bytes))
	callframe.Arg(frame, pointers.Get(base_path))
	callframe.Arg(frame, pointers.Get(state[0])[0])
	callframe.Arg(frame, flags)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFDocument.Bind_append_from_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Takes a Godot Engine scene node and exports it and its descendants to the given [GLTFState] object through the [param state] parameter.
*/
//go:nosplit
func (self class) AppendFromScene(node [1]gdclass.Node, state [1]gdclass.GLTFState, flags gd.Int) gd.Error { //gd:GLTFDocument.append_from_scene
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	callframe.Arg(frame, pointers.Get(state[0])[0])
	callframe.Arg(frame, flags)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFDocument.Bind_append_from_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Takes a [GLTFState] object through the [param state] parameter and returns a Godot Engine scene node.
The [param bake_fps] parameter overrides the bake_fps in [param state].
*/
//go:nosplit
func (self class) GenerateScene(state [1]gdclass.GLTFState, bake_fps gd.Float, trimming bool, remove_immutable_tracks bool) [1]gdclass.Node { //gd:GLTFDocument.generate_scene
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(state[0])[0])
	callframe.Arg(frame, bake_fps)
	callframe.Arg(frame, trimming)
	callframe.Arg(frame, remove_immutable_tracks)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFDocument.Bind_generate_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerWithOwnershipTransferredToGo[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Takes a [GLTFState] object through the [param state] parameter and returns a GLTF [PackedByteArray].
*/
//go:nosplit
func (self class) GenerateBuffer(state [1]gdclass.GLTFState) gd.PackedByteArray { //gd:GLTFDocument.generate_buffer
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(state[0])[0])
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFDocument.Bind_generate_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Takes a [GLTFState] object through the [param state] parameter and writes a glTF file to the filesystem.
[b]Note:[/b] The extension of the glTF file determines if it is a .glb binary file or a .gltf text file.
*/
//go:nosplit
func (self class) WriteToFilesystem(state [1]gdclass.GLTFState, path gd.String) gd.Error { //gd:GLTFDocument.write_to_filesystem
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(state[0])[0])
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFDocument.Bind_write_to_filesystem, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Registers the given [GLTFDocumentExtension] instance with GLTFDocument. If [param first_priority] is true, this extension will be run first. Otherwise, it will be run last.
[b]Note:[/b] Like GLTFDocument itself, all GLTFDocumentExtension classes must be stateless in order to function properly. If you need to store data, use the [code]set_additional_data[/code] and [code]get_additional_data[/code] methods in [GLTFState] or [GLTFNode].
*/
//go:nosplit
func (self class) RegisterGltfDocumentExtension(extension [1]gdclass.GLTFDocumentExtension, first_priority bool) { //gd:GLTFDocument.register_gltf_document_extension
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension[0])[0])
	callframe.Arg(frame, first_priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFDocument.Bind_register_gltf_document_extension, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Unregisters the given [GLTFDocumentExtension] instance.
*/
//go:nosplit
func (self class) UnregisterGltfDocumentExtension(extension [1]gdclass.GLTFDocumentExtension) { //gd:GLTFDocument.unregister_gltf_document_extension
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFDocument.Bind_unregister_gltf_document_extension, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsGLTFDocument() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGLTFDocument() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("GLTFDocument", func(ptr gd.Object) any {
		return [1]gdclass.GLTFDocument{*(*gdclass.GLTFDocument)(unsafe.Pointer(&ptr))}
	})
}

type RootNodeMode = gdclass.GLTFDocumentRootNodeMode //gd:GLTFDocument.RootNodeMode

const (
	/*Treat the Godot scene's root node as the root node of the glTF file, and mark it as the single root node via the [code]GODOT_single_root[/code] glTF extension. This will be parsed the same as [constant ROOT_NODE_MODE_KEEP_ROOT] if the implementation does not support [code]GODOT_single_root[/code].*/
	RootNodeModeSingleRoot RootNodeMode = 0
	/*Treat the Godot scene's root node as the root node of the glTF file, but do not mark it as anything special. An extra root node will be generated when importing into Godot. This uses only vanilla glTF features. This is equivalent to the behavior in Godot 4.1 and earlier.*/
	RootNodeModeKeepRoot RootNodeMode = 1
	/*Treat the Godot scene's root node as the name of the glTF scene, and add all of its children as root nodes of the glTF file. This uses only vanilla glTF features. This avoids an extra root node, but only the name of the Godot scene's root node will be preserved, as it will not be saved as a node.*/
	RootNodeModeMultiRoot RootNodeMode = 2
)

type Error = gd.Error //gd:Error

const (
	/*Methods that return [enum Error] return [constant OK] when no error occurred.
	  Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
	  [b]Example:[/b]
	  [codeblock]
	  var error = method_that_returns_error()
	  if error != OK:
	      printerr("Failure!")

	  # Or, alternatively:
	  if error:
	      printerr("Still failing!")
	  [/codeblock]
	  [b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
	/*Generic error.*/
	Failed Error = 1
	/*Unavailable error.*/
	ErrUnavailable Error = 2
	/*Unconfigured error.*/
	ErrUnconfigured Error = 3
	/*Unauthorized error.*/
	ErrUnauthorized Error = 4
	/*Parameter range error.*/
	ErrParameterRangeError Error = 5
	/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
	/*File: Not found error.*/
	ErrFileNotFound Error = 7
	/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
	/*File: Bad path error.*/
	ErrFileBadPath Error = 9
	/*File: No permission error.*/
	ErrFileNoPermission Error = 10
	/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
	/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
	/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
	/*File: Can't read error.*/
	ErrFileCantRead Error = 14
	/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
	/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
	/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
	/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
	/*Can't open error.*/
	ErrCantOpen Error = 19
	/*Can't create error.*/
	ErrCantCreate Error = 20
	/*Query failed error.*/
	ErrQueryFailed Error = 21
	/*Already in use error.*/
	ErrAlreadyInUse Error = 22
	/*Locked error.*/
	ErrLocked Error = 23
	/*Timeout error.*/
	ErrTimeout Error = 24
	/*Can't connect error.*/
	ErrCantConnect Error = 25
	/*Can't resolve error.*/
	ErrCantResolve Error = 26
	/*Connection error.*/
	ErrConnectionError Error = 27
	/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
	/*Can't fork process error.*/
	ErrCantFork Error = 29
	/*Invalid data error.*/
	ErrInvalidData Error = 30
	/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
	/*Already exists error.*/
	ErrAlreadyExists Error = 32
	/*Does not exist error.*/
	ErrDoesNotExist Error = 33
	/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
	/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
	/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
	/*Method not found error.*/
	ErrMethodNotFound Error = 37
	/*Linking failed error.*/
	ErrLinkFailed Error = 38
	/*Script failed error.*/
	ErrScriptFailed Error = 39
	/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
	/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
	/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
	/*Parse error.*/
	ErrParseError Error = 43
	/*Busy error.*/
	ErrBusy Error = 44
	/*Skip error.*/
	ErrSkip Error = 45
	/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
	/*Bug error, caused by an implementation issue in the method.
	  [b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
	/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
)
