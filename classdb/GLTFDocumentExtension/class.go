// Package GLTFDocumentExtension provides methods for working with GLTFDocumentExtension object instances.
package GLTFDocumentExtension

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Extends the functionality of the [GLTFDocument] class by allowing you to run arbitrary code at various stages of GLTF import or export.
To use, make a new class extending GLTFDocumentExtension, override any methods you need, make an instance of your class, and register it using [method GLTFDocument.register_gltf_document_extension].
[b]Note:[/b] Like GLTFDocument itself, all GLTFDocumentExtension classes must be stateless in order to function properly. If you need to store data, use the [code]set_additional_data[/code] and [code]get_additional_data[/code] methods in [GLTFState] or [GLTFNode].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=GLTFDocumentExtension)
*/
type Instance [1]gdclass.GLTFDocumentExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGLTFDocumentExtension() Instance
}
type Interface interface {
	//Part of the import process. This method is run first, before all other parts of the import process.
	//The return value is used to determine if this [GLTFDocumentExtension] instance should be used for importing a given GLTF file. If [constant OK], the import will use this [GLTFDocumentExtension] instance. If not overridden, [constant OK] is returned.
	ImportPreflight(state [1]gdclass.GLTFState, extensions []string) error
	//Part of the import process. This method is run after [method _import_preflight] and before [method _parse_node_extensions].
	//Returns an array of the GLTF extensions supported by this GLTFDocumentExtension class. This is used to validate if a GLTF file with required extensions can be loaded.
	GetSupportedExtensions() []string
	//Part of the import process. This method is run after [method _get_supported_extensions] and before [method _import_post_parse].
	//Runs when parsing the node extensions of a GLTFNode. This method can be used to process the extension JSON data into a format that can be used by [method _generate_scene_node]. The return value should be a member of the [enum Error] enum.
	ParseNodeExtensions(state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, extensions Dictionary.Any) error
	//Part of the import process. This method is run after [method _parse_node_extensions] and before [method _parse_texture_json].
	//Runs when parsing image data from a GLTF file. The data could be sourced from a separate file, a URI, or a buffer, and then is passed as a byte array.
	ParseImageData(state [1]gdclass.GLTFState, image_data []byte, mime_type string, ret_image [1]gdclass.Image) error
	//Returns the file extension to use for saving image data into, for example, [code]".png"[/code]. If defined, when this extension is used to handle images, and the images are saved to a separate file, the image bytes will be copied to a file with this extension. If this is set, there should be a [ResourceImporter] class able to import the file. If not defined or empty, Godot will save the image into a PNG file.
	GetImageFileExtension() string
	//Part of the import process. This method is run after [method _parse_image_data] and before [method _generate_scene_node].
	//Runs when parsing the texture JSON from the GLTF textures array. This can be used to set the source image index to use as the texture.
	ParseTextureJson(state [1]gdclass.GLTFState, texture_json Dictionary.Any, ret_gltf_texture [1]gdclass.GLTFTexture) error
	//Part of the import process. This method is run after [method _import_post_parse] and before [method _import_node].
	//Runs when generating a Godot scene node from a GLTFNode. The returned node will be added to the scene tree. Multiple nodes can be generated in this step if they are added as a child of the returned node.
	//[b]Note:[/b] The [param scene_parent] parameter may be null if this is the single root node.
	GenerateSceneNode(state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, scene_parent [1]gdclass.Node) [1]gdclass.Node3D
	//Part of the import process. This method is run after [method _parse_node_extensions] and before [method _generate_scene_node].
	//This method can be used to modify any of the data imported so far after parsing, before generating the nodes and then running the final per-node import step.
	ImportPostParse(state [1]gdclass.GLTFState) error
	//Part of the import process. This method is run after [method _generate_scene_node] and before [method _import_post].
	//This method can be used to make modifications to each of the generated Godot scene nodes.
	ImportNode(state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, json Dictionary.Any, node [1]gdclass.Node) error
	//Part of the import process. This method is run last, after all other parts of the import process.
	//This method can be used to modify the final Godot scene generated by the import process.
	ImportPost(state [1]gdclass.GLTFState, root [1]gdclass.Node) error
	//Part of the export process. This method is run first, before all other parts of the export process.
	//The return value is used to determine if this [GLTFDocumentExtension] instance should be used for exporting a given GLTF file. If [constant OK], the export will use this [GLTFDocumentExtension] instance. If not overridden, [constant OK] is returned.
	ExportPreflight(state [1]gdclass.GLTFState, root [1]gdclass.Node) error
	//Part of the export process. This method is run after [method _export_preflight] and before [method _export_preserialize].
	//Runs when converting the data from a Godot scene node. This method can be used to process the Godot scene node data into a format that can be used by [method _export_node].
	ConvertSceneNode(state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, scene_node [1]gdclass.Node)
	//Part of the export process. This method is run after [method _convert_scene_node] and before [method _get_saveable_image_formats].
	//This method can be used to alter the state before performing serialization. It runs every time when generating a buffer with [method GLTFDocument.generate_buffer] or writing to the file system with [method GLTFDocument.write_to_filesystem].
	ExportPreserialize(state [1]gdclass.GLTFState) error
	//Part of the export process. This method is run after [method _convert_scene_node] and before [method _export_node].
	//Returns an array of the image formats that can be saved/exported by this extension. This extension will only be selected as the image exporter if the [GLTFDocument]'s [member GLTFDocument.image_format] is in this array. If this [GLTFDocumentExtension] is selected as the image exporter, one of the [method _save_image_at_path] or [method _serialize_image_to_bytes] methods will run next, otherwise [method _export_node] will run next. If the format name contains [code]"Lossy"[/code], the lossy quality slider will be displayed.
	GetSaveableImageFormats() []string
	//Part of the export process. This method is run after [method _get_saveable_image_formats] and before [method _serialize_texture_json].
	//This method is run when embedding images in the GLTF file. When images are saved separately, [method _save_image_at_path] runs instead. Note that these methods only run when this [GLTFDocumentExtension] is selected as the image exporter.
	//This method must set the image MIME type in the [param image_dict] with the [code]"mimeType"[/code] key. For example, for a PNG image, it would be set to [code]"image/png"[/code]. The return value must be a [PackedByteArray] containing the image data.
	SerializeImageToBytes(state [1]gdclass.GLTFState, image [1]gdclass.Image, image_dict Dictionary.Any, image_format string, lossy_quality Float.X) []byte
	//Part of the export process. This method is run after [method _get_saveable_image_formats] and before [method _serialize_texture_json].
	//This method is run when saving images separately from the GLTF file. When images are embedded, [method _serialize_image_to_bytes] runs instead. Note that these methods only run when this [GLTFDocumentExtension] is selected as the image exporter.
	SaveImageAtPath(state [1]gdclass.GLTFState, image [1]gdclass.Image, file_path string, image_format string, lossy_quality Float.X) error
	//Part of the export process. This method is run after [method _save_image_at_path] or [method _serialize_image_to_bytes], and before [method _export_node]. Note that this method only runs when this [GLTFDocumentExtension] is selected as the image exporter.
	//This method can be used to set up the extensions for the texture JSON by editing [param texture_json]. The extension must also be added as used extension with [method GLTFState.add_used_extension], be sure to set [code]required[/code] to [code]true[/code] if you are not providing a fallback.
	SerializeTextureJson(state [1]gdclass.GLTFState, texture_json Dictionary.Any, gltf_texture [1]gdclass.GLTFTexture, image_format string) error
	//Part of the export process. This method is run after [method _get_saveable_image_formats] and before [method _export_post]. If this [GLTFDocumentExtension] is used for exporting images, this runs after [method _serialize_texture_json].
	//This method can be used to modify the final JSON of each node. Data should be primarily stored in [param gltf_node] prior to serializing the JSON, but the original Godot [param node] is also provided if available. The node may be null if not available, such as when exporting GLTF data not generated from a Godot scene.
	ExportNode(state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, json Dictionary.Any, node [1]gdclass.Node) error
	//Part of the export process. This method is run last, after all other parts of the export process.
	//This method can be used to modify the final JSON of the generated GLTF file.
	ExportPost(state [1]gdclass.GLTFState) error
}

// Implementation implements [Interface] with empty methods.
type Implementation struct{}

func (self Implementation) ImportPreflight(state [1]gdclass.GLTFState, extensions []string) (_ error) {
	return
}
func (self Implementation) GetSupportedExtensions() (_ []string) { return }
func (self Implementation) ParseNodeExtensions(state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, extensions Dictionary.Any) (_ error) {
	return
}
func (self Implementation) ParseImageData(state [1]gdclass.GLTFState, image_data []byte, mime_type string, ret_image [1]gdclass.Image) (_ error) {
	return
}
func (self Implementation) GetImageFileExtension() (_ string) { return }
func (self Implementation) ParseTextureJson(state [1]gdclass.GLTFState, texture_json Dictionary.Any, ret_gltf_texture [1]gdclass.GLTFTexture) (_ error) {
	return
}
func (self Implementation) GenerateSceneNode(state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, scene_parent [1]gdclass.Node) (_ [1]gdclass.Node3D) {
	return
}
func (self Implementation) ImportPostParse(state [1]gdclass.GLTFState) (_ error) { return }
func (self Implementation) ImportNode(state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, json Dictionary.Any, node [1]gdclass.Node) (_ error) {
	return
}
func (self Implementation) ImportPost(state [1]gdclass.GLTFState, root [1]gdclass.Node) (_ error) {
	return
}
func (self Implementation) ExportPreflight(state [1]gdclass.GLTFState, root [1]gdclass.Node) (_ error) {
	return
}
func (self Implementation) ConvertSceneNode(state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, scene_node [1]gdclass.Node) {
	return
}
func (self Implementation) ExportPreserialize(state [1]gdclass.GLTFState) (_ error) { return }
func (self Implementation) GetSaveableImageFormats() (_ []string)                   { return }
func (self Implementation) SerializeImageToBytes(state [1]gdclass.GLTFState, image [1]gdclass.Image, image_dict Dictionary.Any, image_format string, lossy_quality Float.X) (_ []byte) {
	return
}
func (self Implementation) SaveImageAtPath(state [1]gdclass.GLTFState, image [1]gdclass.Image, file_path string, image_format string, lossy_quality Float.X) (_ error) {
	return
}
func (self Implementation) SerializeTextureJson(state [1]gdclass.GLTFState, texture_json Dictionary.Any, gltf_texture [1]gdclass.GLTFTexture, image_format string) (_ error) {
	return
}
func (self Implementation) ExportNode(state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, json Dictionary.Any, node [1]gdclass.Node) (_ error) {
	return
}
func (self Implementation) ExportPost(state [1]gdclass.GLTFState) (_ error) { return }

/*
Part of the import process. This method is run first, before all other parts of the import process.
The return value is used to determine if this [GLTFDocumentExtension] instance should be used for importing a given GLTF file. If [constant OK], the import will use this [GLTFDocumentExtension] instance. If not overridden, [constant OK] is returned.
*/
func (Instance) _import_preflight(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, extensions []string) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(state[0])
		var extensions = pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 1))
		defer pointers.End(extensions)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, extensions.Strings())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the import process. This method is run after [method _import_preflight] and before [method _parse_node_extensions].
Returns an array of the GLTF extensions supported by this GLTFDocumentExtension class. This is used to validate if a GLTF file with required extensions can be loaded.
*/
func (Instance) _get_supported_extensions(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Part of the import process. This method is run after [method _get_supported_extensions] and before [method _import_post_parse].
Runs when parsing the node extensions of a GLTFNode. This method can be used to process the extension JSON data into a format that can be used by [method _generate_scene_node]. The return value should be a member of the [enum Error] enum.
*/
func (Instance) _parse_node_extensions(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, extensions Dictionary.Any) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(state[0])
		var gltf_node = [1]gdclass.GLTFNode{pointers.New[gdclass.GLTFNode]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 1))})}
		defer pointers.End(gltf_node[0])
		var extensions = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		defer pointers.End(extensions)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, gltf_node, extensions)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the import process. This method is run after [method _parse_node_extensions] and before [method _parse_texture_json].
Runs when parsing image data from a GLTF file. The data could be sourced from a separate file, a URI, or a buffer, and then is passed as a byte array.
*/
func (Instance) _parse_image_data(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, image_data []byte, mime_type string, ret_image [1]gdclass.Image) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(state[0])
		var image_data = pointers.New[gd.PackedByteArray](gd.UnsafeGet[gd.PackedPointers](p_args, 1))
		defer pointers.End(image_data)
		var mime_type = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		defer pointers.End(mime_type)
		var ret_image = [1]gdclass.Image{pointers.New[gdclass.Image]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 3))})}
		defer pointers.End(ret_image[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, image_data.Bytes(), mime_type.String(), ret_image)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the file extension to use for saving image data into, for example, [code]".png"[/code]. If defined, when this extension is used to handle images, and the images are saved to a separate file, the image bytes will be copied to a file with this extension. If this is set, there should be a [ResourceImporter] class able to import the file. If not defined or empty, Godot will save the image into a PNG file.
*/
func (Instance) _get_image_file_extension(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Part of the import process. This method is run after [method _parse_image_data] and before [method _generate_scene_node].
Runs when parsing the texture JSON from the GLTF textures array. This can be used to set the source image index to use as the texture.
*/
func (Instance) _parse_texture_json(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, texture_json Dictionary.Any, ret_gltf_texture [1]gdclass.GLTFTexture) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(state[0])
		var texture_json = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		defer pointers.End(texture_json)
		var ret_gltf_texture = [1]gdclass.GLTFTexture{pointers.New[gdclass.GLTFTexture]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 2))})}
		defer pointers.End(ret_gltf_texture[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, texture_json, ret_gltf_texture)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the import process. This method is run after [method _import_post_parse] and before [method _import_node].
Runs when generating a Godot scene node from a GLTFNode. The returned node will be added to the scene tree. Multiple nodes can be generated in this step if they are added as a child of the returned node.
[b]Note:[/b] The [param scene_parent] parameter may be null if this is the single root node.
*/
func (Instance) _generate_scene_node(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, scene_parent [1]gdclass.Node) [1]gdclass.Node3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(state[0])
		var gltf_node = [1]gdclass.GLTFNode{pointers.New[gdclass.GLTFNode]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 1))})}
		defer pointers.End(gltf_node[0])
		var scene_parent = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 2))})}
		defer pointers.End(scene_parent[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, gltf_node, scene_parent)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Part of the import process. This method is run after [method _parse_node_extensions] and before [method _generate_scene_node].
This method can be used to modify any of the data imported so far after parsing, before generating the nodes and then running the final per-node import step.
*/
func (Instance) _import_post_parse(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(state[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the import process. This method is run after [method _generate_scene_node] and before [method _import_post].
This method can be used to make modifications to each of the generated Godot scene nodes.
*/
func (Instance) _import_node(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, json Dictionary.Any, node [1]gdclass.Node) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(state[0])
		var gltf_node = [1]gdclass.GLTFNode{pointers.New[gdclass.GLTFNode]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 1))})}
		defer pointers.End(gltf_node[0])
		var json = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		defer pointers.End(json)
		var node = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 3))})}
		defer pointers.End(node[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, gltf_node, json, node)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the import process. This method is run last, after all other parts of the import process.
This method can be used to modify the final Godot scene generated by the import process.
*/
func (Instance) _import_post(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, root [1]gdclass.Node) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(state[0])
		var root = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 1))})}
		defer pointers.End(root[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, root)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the export process. This method is run first, before all other parts of the export process.
The return value is used to determine if this [GLTFDocumentExtension] instance should be used for exporting a given GLTF file. If [constant OK], the export will use this [GLTFDocumentExtension] instance. If not overridden, [constant OK] is returned.
*/
func (Instance) _export_preflight(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, root [1]gdclass.Node) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(state[0])
		var root = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 1))})}
		defer pointers.End(root[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, root)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the export process. This method is run after [method _export_preflight] and before [method _export_preserialize].
Runs when converting the data from a Godot scene node. This method can be used to process the Godot scene node data into a format that can be used by [method _export_node].
*/
func (Instance) _convert_scene_node(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, scene_node [1]gdclass.Node)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(state[0])
		var gltf_node = [1]gdclass.GLTFNode{pointers.New[gdclass.GLTFNode]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 1))})}
		defer pointers.End(gltf_node[0])
		var scene_node = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 2))})}
		defer pointers.End(scene_node[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, state, gltf_node, scene_node)
	}
}

/*
Part of the export process. This method is run after [method _convert_scene_node] and before [method _get_saveable_image_formats].
This method can be used to alter the state before performing serialization. It runs every time when generating a buffer with [method GLTFDocument.generate_buffer] or writing to the file system with [method GLTFDocument.write_to_filesystem].
*/
func (Instance) _export_preserialize(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(state[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the export process. This method is run after [method _convert_scene_node] and before [method _export_node].
Returns an array of the image formats that can be saved/exported by this extension. This extension will only be selected as the image exporter if the [GLTFDocument]'s [member GLTFDocument.image_format] is in this array. If this [GLTFDocumentExtension] is selected as the image exporter, one of the [method _save_image_at_path] or [method _serialize_image_to_bytes] methods will run next, otherwise [method _export_node] will run next. If the format name contains [code]"Lossy"[/code], the lossy quality slider will be displayed.
*/
func (Instance) _get_saveable_image_formats(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Part of the export process. This method is run after [method _get_saveable_image_formats] and before [method _serialize_texture_json].
This method is run when embedding images in the GLTF file. When images are saved separately, [method _save_image_at_path] runs instead. Note that these methods only run when this [GLTFDocumentExtension] is selected as the image exporter.
This method must set the image MIME type in the [param image_dict] with the [code]"mimeType"[/code] key. For example, for a PNG image, it would be set to [code]"image/png"[/code]. The return value must be a [PackedByteArray] containing the image data.
*/
func (Instance) _serialize_image_to_bytes(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, image [1]gdclass.Image, image_dict Dictionary.Any, image_format string, lossy_quality Float.X) []byte) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(state[0])
		var image = [1]gdclass.Image{pointers.New[gdclass.Image]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 1))})}
		defer pointers.End(image[0])
		var image_dict = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		defer pointers.End(image_dict)
		var image_format = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 3))
		defer pointers.End(image_format)
		var lossy_quality = gd.UnsafeGet[gd.Float](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, image, image_dict, image_format.String(), Float.X(lossy_quality))
		ptr, ok := pointers.End(gd.NewPackedByteSlice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Part of the export process. This method is run after [method _get_saveable_image_formats] and before [method _serialize_texture_json].
This method is run when saving images separately from the GLTF file. When images are embedded, [method _serialize_image_to_bytes] runs instead. Note that these methods only run when this [GLTFDocumentExtension] is selected as the image exporter.
*/
func (Instance) _save_image_at_path(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, image [1]gdclass.Image, file_path string, image_format string, lossy_quality Float.X) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(state[0])
		var image = [1]gdclass.Image{pointers.New[gdclass.Image]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 1))})}
		defer pointers.End(image[0])
		var file_path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		defer pointers.End(file_path)
		var image_format = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 3))
		defer pointers.End(image_format)
		var lossy_quality = gd.UnsafeGet[gd.Float](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, image, file_path.String(), image_format.String(), Float.X(lossy_quality))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the export process. This method is run after [method _save_image_at_path] or [method _serialize_image_to_bytes], and before [method _export_node]. Note that this method only runs when this [GLTFDocumentExtension] is selected as the image exporter.
This method can be used to set up the extensions for the texture JSON by editing [param texture_json]. The extension must also be added as used extension with [method GLTFState.add_used_extension], be sure to set [code]required[/code] to [code]true[/code] if you are not providing a fallback.
*/
func (Instance) _serialize_texture_json(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, texture_json Dictionary.Any, gltf_texture [1]gdclass.GLTFTexture, image_format string) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(state[0])
		var texture_json = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		defer pointers.End(texture_json)
		var gltf_texture = [1]gdclass.GLTFTexture{pointers.New[gdclass.GLTFTexture]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 2))})}
		defer pointers.End(gltf_texture[0])
		var image_format = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 3))
		defer pointers.End(image_format)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, texture_json, gltf_texture, image_format.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the export process. This method is run after [method _get_saveable_image_formats] and before [method _export_post]. If this [GLTFDocumentExtension] is used for exporting images, this runs after [method _serialize_texture_json].
This method can be used to modify the final JSON of each node. Data should be primarily stored in [param gltf_node] prior to serializing the JSON, but the original Godot [param node] is also provided if available. The node may be null if not available, such as when exporting GLTF data not generated from a Godot scene.
*/
func (Instance) _export_node(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, json Dictionary.Any, node [1]gdclass.Node) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(state[0])
		var gltf_node = [1]gdclass.GLTFNode{pointers.New[gdclass.GLTFNode]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 1))})}
		defer pointers.End(gltf_node[0])
		var json = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		defer pointers.End(json)
		var node = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 3))})}
		defer pointers.End(node[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, gltf_node, json, node)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the export process. This method is run last, after all other parts of the export process.
This method can be used to modify the final JSON of the generated GLTF file.
*/
func (Instance) _export_post(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(state[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state)
		gd.UnsafeSet(p_back, ret)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GLTFDocumentExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFDocumentExtension"))
	casted := Instance{*(*gdclass.GLTFDocumentExtension)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Part of the import process. This method is run first, before all other parts of the import process.
The return value is used to determine if this [GLTFDocumentExtension] instance should be used for importing a given GLTF file. If [constant OK], the import will use this [GLTFDocumentExtension] instance. If not overridden, [constant OK] is returned.
*/
func (class) _import_preflight(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, extensions gd.PackedStringArray) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(state[0])
		var extensions = pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, extensions)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the import process. This method is run after [method _import_preflight] and before [method _parse_node_extensions].
Returns an array of the GLTF extensions supported by this GLTFDocumentExtension class. This is used to validate if a GLTF file with required extensions can be loaded.
*/
func (class) _get_supported_extensions(impl func(ptr unsafe.Pointer) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Part of the import process. This method is run after [method _get_supported_extensions] and before [method _import_post_parse].
Runs when parsing the node extensions of a GLTFNode. This method can be used to process the extension JSON data into a format that can be used by [method _generate_scene_node]. The return value should be a member of the [enum Error] enum.
*/
func (class) _parse_node_extensions(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, extensions gd.Dictionary) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(state[0])
		var gltf_node = [1]gdclass.GLTFNode{pointers.New[gdclass.GLTFNode]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}
		defer pointers.End(gltf_node[0])
		var extensions = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, gltf_node, extensions)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the import process. This method is run after [method _parse_node_extensions] and before [method _parse_texture_json].
Runs when parsing image data from a GLTF file. The data could be sourced from a separate file, a URI, or a buffer, and then is passed as a byte array.
*/
func (class) _parse_image_data(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, image_data gd.PackedByteArray, mime_type gd.String, ret_image [1]gdclass.Image) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(state[0])
		var image_data = pointers.New[gd.PackedByteArray](gd.UnsafeGet[gd.PackedPointers](p_args, 1))
		var mime_type = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		var ret_image = [1]gdclass.Image{pointers.New[gdclass.Image]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 3))})}
		defer pointers.End(ret_image[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, image_data, mime_type, ret_image)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the file extension to use for saving image data into, for example, [code]".png"[/code]. If defined, when this extension is used to handle images, and the images are saved to a separate file, the image bytes will be copied to a file with this extension. If this is set, there should be a [ResourceImporter] class able to import the file. If not defined or empty, Godot will save the image into a PNG file.
*/
func (class) _get_image_file_extension(impl func(ptr unsafe.Pointer) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Part of the import process. This method is run after [method _parse_image_data] and before [method _generate_scene_node].
Runs when parsing the texture JSON from the GLTF textures array. This can be used to set the source image index to use as the texture.
*/
func (class) _parse_texture_json(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, texture_json gd.Dictionary, ret_gltf_texture [1]gdclass.GLTFTexture) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(state[0])
		var texture_json = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		var ret_gltf_texture = [1]gdclass.GLTFTexture{pointers.New[gdclass.GLTFTexture]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 2))})}
		defer pointers.End(ret_gltf_texture[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, texture_json, ret_gltf_texture)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the import process. This method is run after [method _import_post_parse] and before [method _import_node].
Runs when generating a Godot scene node from a GLTFNode. The returned node will be added to the scene tree. Multiple nodes can be generated in this step if they are added as a child of the returned node.
[b]Note:[/b] The [param scene_parent] parameter may be null if this is the single root node.
*/
func (class) _generate_scene_node(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, scene_parent [1]gdclass.Node) [1]gdclass.Node3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(state[0])
		var gltf_node = [1]gdclass.GLTFNode{pointers.New[gdclass.GLTFNode]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}
		defer pointers.End(gltf_node[0])
		var scene_parent = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 2))})}
		defer pointers.End(scene_parent[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, gltf_node, scene_parent)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Part of the import process. This method is run after [method _parse_node_extensions] and before [method _generate_scene_node].
This method can be used to modify any of the data imported so far after parsing, before generating the nodes and then running the final per-node import step.
*/
func (class) _import_post_parse(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(state[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the import process. This method is run after [method _generate_scene_node] and before [method _import_post].
This method can be used to make modifications to each of the generated Godot scene nodes.
*/
func (class) _import_node(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, json gd.Dictionary, node [1]gdclass.Node) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(state[0])
		var gltf_node = [1]gdclass.GLTFNode{pointers.New[gdclass.GLTFNode]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}
		defer pointers.End(gltf_node[0])
		var json = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		var node = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 3))})}
		defer pointers.End(node[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, gltf_node, json, node)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the import process. This method is run last, after all other parts of the import process.
This method can be used to modify the final Godot scene generated by the import process.
*/
func (class) _import_post(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, root [1]gdclass.Node) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(state[0])
		var root = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}
		defer pointers.End(root[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, root)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the export process. This method is run first, before all other parts of the export process.
The return value is used to determine if this [GLTFDocumentExtension] instance should be used for exporting a given GLTF file. If [constant OK], the export will use this [GLTFDocumentExtension] instance. If not overridden, [constant OK] is returned.
*/
func (class) _export_preflight(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, root [1]gdclass.Node) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(state[0])
		var root = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}
		defer pointers.End(root[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, root)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the export process. This method is run after [method _export_preflight] and before [method _export_preserialize].
Runs when converting the data from a Godot scene node. This method can be used to process the Godot scene node data into a format that can be used by [method _export_node].
*/
func (class) _convert_scene_node(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, scene_node [1]gdclass.Node)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(state[0])
		var gltf_node = [1]gdclass.GLTFNode{pointers.New[gdclass.GLTFNode]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}
		defer pointers.End(gltf_node[0])
		var scene_node = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 2))})}
		defer pointers.End(scene_node[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, state, gltf_node, scene_node)
	}
}

/*
Part of the export process. This method is run after [method _convert_scene_node] and before [method _get_saveable_image_formats].
This method can be used to alter the state before performing serialization. It runs every time when generating a buffer with [method GLTFDocument.generate_buffer] or writing to the file system with [method GLTFDocument.write_to_filesystem].
*/
func (class) _export_preserialize(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(state[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the export process. This method is run after [method _convert_scene_node] and before [method _export_node].
Returns an array of the image formats that can be saved/exported by this extension. This extension will only be selected as the image exporter if the [GLTFDocument]'s [member GLTFDocument.image_format] is in this array. If this [GLTFDocumentExtension] is selected as the image exporter, one of the [method _save_image_at_path] or [method _serialize_image_to_bytes] methods will run next, otherwise [method _export_node] will run next. If the format name contains [code]"Lossy"[/code], the lossy quality slider will be displayed.
*/
func (class) _get_saveable_image_formats(impl func(ptr unsafe.Pointer) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Part of the export process. This method is run after [method _get_saveable_image_formats] and before [method _serialize_texture_json].
This method is run when embedding images in the GLTF file. When images are saved separately, [method _save_image_at_path] runs instead. Note that these methods only run when this [GLTFDocumentExtension] is selected as the image exporter.
This method must set the image MIME type in the [param image_dict] with the [code]"mimeType"[/code] key. For example, for a PNG image, it would be set to [code]"image/png"[/code]. The return value must be a [PackedByteArray] containing the image data.
*/
func (class) _serialize_image_to_bytes(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, image [1]gdclass.Image, image_dict gd.Dictionary, image_format gd.String, lossy_quality gd.Float) gd.PackedByteArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(state[0])
		var image = [1]gdclass.Image{pointers.New[gdclass.Image]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}
		defer pointers.End(image[0])
		var image_dict = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		var image_format = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 3))
		var lossy_quality = gd.UnsafeGet[gd.Float](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, image, image_dict, image_format, lossy_quality)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Part of the export process. This method is run after [method _get_saveable_image_formats] and before [method _serialize_texture_json].
This method is run when saving images separately from the GLTF file. When images are embedded, [method _serialize_image_to_bytes] runs instead. Note that these methods only run when this [GLTFDocumentExtension] is selected as the image exporter.
*/
func (class) _save_image_at_path(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, image [1]gdclass.Image, file_path gd.String, image_format gd.String, lossy_quality gd.Float) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(state[0])
		var image = [1]gdclass.Image{pointers.New[gdclass.Image]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}
		defer pointers.End(image[0])
		var file_path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		var image_format = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 3))
		var lossy_quality = gd.UnsafeGet[gd.Float](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, image, file_path, image_format, lossy_quality)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the export process. This method is run after [method _save_image_at_path] or [method _serialize_image_to_bytes], and before [method _export_node]. Note that this method only runs when this [GLTFDocumentExtension] is selected as the image exporter.
This method can be used to set up the extensions for the texture JSON by editing [param texture_json]. The extension must also be added as used extension with [method GLTFState.add_used_extension], be sure to set [code]required[/code] to [code]true[/code] if you are not providing a fallback.
*/
func (class) _serialize_texture_json(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, texture_json gd.Dictionary, gltf_texture [1]gdclass.GLTFTexture, image_format gd.String) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(state[0])
		var texture_json = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		var gltf_texture = [1]gdclass.GLTFTexture{pointers.New[gdclass.GLTFTexture]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 2))})}
		defer pointers.End(gltf_texture[0])
		var image_format = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 3))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, texture_json, gltf_texture, image_format)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the export process. This method is run after [method _get_saveable_image_formats] and before [method _export_post]. If this [GLTFDocumentExtension] is used for exporting images, this runs after [method _serialize_texture_json].
This method can be used to modify the final JSON of each node. Data should be primarily stored in [param gltf_node] prior to serializing the JSON, but the original Godot [param node] is also provided if available. The node may be null if not available, such as when exporting GLTF data not generated from a Godot scene.
*/
func (class) _export_node(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState, gltf_node [1]gdclass.GLTFNode, json gd.Dictionary, node [1]gdclass.Node) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(state[0])
		var gltf_node = [1]gdclass.GLTFNode{pointers.New[gdclass.GLTFNode]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}
		defer pointers.End(gltf_node[0])
		var json = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		var node = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 3))})}
		defer pointers.End(node[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, gltf_node, json, node)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the export process. This method is run last, after all other parts of the export process.
This method can be used to modify the final JSON of the generated GLTF file.
*/
func (class) _export_post(impl func(ptr unsafe.Pointer, state [1]gdclass.GLTFState) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = [1]gdclass.GLTFState{pointers.New[gdclass.GLTFState]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(state[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsGLTFDocumentExtension() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGLTFDocumentExtension() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	case "_import_preflight":
		return reflect.ValueOf(self._import_preflight)
	case "_get_supported_extensions":
		return reflect.ValueOf(self._get_supported_extensions)
	case "_parse_node_extensions":
		return reflect.ValueOf(self._parse_node_extensions)
	case "_parse_image_data":
		return reflect.ValueOf(self._parse_image_data)
	case "_get_image_file_extension":
		return reflect.ValueOf(self._get_image_file_extension)
	case "_parse_texture_json":
		return reflect.ValueOf(self._parse_texture_json)
	case "_generate_scene_node":
		return reflect.ValueOf(self._generate_scene_node)
	case "_import_post_parse":
		return reflect.ValueOf(self._import_post_parse)
	case "_import_node":
		return reflect.ValueOf(self._import_node)
	case "_import_post":
		return reflect.ValueOf(self._import_post)
	case "_export_preflight":
		return reflect.ValueOf(self._export_preflight)
	case "_convert_scene_node":
		return reflect.ValueOf(self._convert_scene_node)
	case "_export_preserialize":
		return reflect.ValueOf(self._export_preserialize)
	case "_get_saveable_image_formats":
		return reflect.ValueOf(self._get_saveable_image_formats)
	case "_serialize_image_to_bytes":
		return reflect.ValueOf(self._serialize_image_to_bytes)
	case "_save_image_at_path":
		return reflect.ValueOf(self._save_image_at_path)
	case "_serialize_texture_json":
		return reflect.ValueOf(self._serialize_texture_json)
	case "_export_node":
		return reflect.ValueOf(self._export_node)
	case "_export_post":
		return reflect.ValueOf(self._export_post)
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_import_preflight":
		return reflect.ValueOf(self._import_preflight)
	case "_get_supported_extensions":
		return reflect.ValueOf(self._get_supported_extensions)
	case "_parse_node_extensions":
		return reflect.ValueOf(self._parse_node_extensions)
	case "_parse_image_data":
		return reflect.ValueOf(self._parse_image_data)
	case "_get_image_file_extension":
		return reflect.ValueOf(self._get_image_file_extension)
	case "_parse_texture_json":
		return reflect.ValueOf(self._parse_texture_json)
	case "_generate_scene_node":
		return reflect.ValueOf(self._generate_scene_node)
	case "_import_post_parse":
		return reflect.ValueOf(self._import_post_parse)
	case "_import_node":
		return reflect.ValueOf(self._import_node)
	case "_import_post":
		return reflect.ValueOf(self._import_post)
	case "_export_preflight":
		return reflect.ValueOf(self._export_preflight)
	case "_convert_scene_node":
		return reflect.ValueOf(self._convert_scene_node)
	case "_export_preserialize":
		return reflect.ValueOf(self._export_preserialize)
	case "_get_saveable_image_formats":
		return reflect.ValueOf(self._get_saveable_image_formats)
	case "_serialize_image_to_bytes":
		return reflect.ValueOf(self._serialize_image_to_bytes)
	case "_save_image_at_path":
		return reflect.ValueOf(self._save_image_at_path)
	case "_serialize_texture_json":
		return reflect.ValueOf(self._serialize_texture_json)
	case "_export_node":
		return reflect.ValueOf(self._export_node)
	case "_export_post":
		return reflect.ValueOf(self._export_post)
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("GLTFDocumentExtension", func(ptr gd.Object) any {
		return [1]gdclass.GLTFDocumentExtension{*(*gdclass.GLTFDocumentExtension)(unsafe.Pointer(&ptr))}
	})
}

type Error int

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
