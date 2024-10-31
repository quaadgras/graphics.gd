package GLTFDocumentExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Extends the functionality of the [GLTFDocument] class by allowing you to run arbitrary code at various stages of GLTF import or export.
To use, make a new class extending GLTFDocumentExtension, override any methods you need, make an instance of your class, and register it using [method GLTFDocument.register_gltf_document_extension].
[b]Note:[/b] Like GLTFDocument itself, all GLTFDocumentExtension classes must be stateless in order to function properly. If you need to store data, use the [code]set_additional_data[/code] and [code]get_additional_data[/code] methods in [GLTFState] or [GLTFNode].

	// GLTFDocumentExtension methods that can be overridden by a [Class] that extends it.
	type GLTFDocumentExtension interface {
		//Part of the import process. This method is run first, before all other parts of the import process.
		//The return value is used to determine if this [GLTFDocumentExtension] instance should be used for importing a given GLTF file. If [constant OK], the import will use this [GLTFDocumentExtension] instance. If not overridden, [constant OK] is returned.
		ImportPreflight(state gdclass.GLTFState, extensions []string) gd.Error
		//Part of the import process. This method is run after [method _import_preflight] and before [method _parse_node_extensions].
		//Returns an array of the GLTF extensions supported by this GLTFDocumentExtension class. This is used to validate if a GLTF file with required extensions can be loaded.
		GetSupportedExtensions() []string
		//Part of the import process. This method is run after [method _get_supported_extensions] and before [method _import_post_parse].
		//Runs when parsing the node extensions of a GLTFNode. This method can be used to process the extension JSON data into a format that can be used by [method _generate_scene_node]. The return value should be a member of the [enum Error] enum.
		ParseNodeExtensions(state gdclass.GLTFState, gltf_node gdclass.GLTFNode, extensions gd.Dictionary) gd.Error
		//Part of the import process. This method is run after [method _parse_node_extensions] and before [method _parse_texture_json].
		//Runs when parsing image data from a GLTF file. The data could be sourced from a separate file, a URI, or a buffer, and then is passed as a byte array.
		ParseImageData(state gdclass.GLTFState, image_data []byte, mime_type string, ret_image gdclass.Image) gd.Error
		//Returns the file extension to use for saving image data into, for example, [code]".png"[/code]. If defined, when this extension is used to handle images, and the images are saved to a separate file, the image bytes will be copied to a file with this extension. If this is set, there should be a [ResourceImporter] class able to import the file. If not defined or empty, Godot will save the image into a PNG file.
		GetImageFileExtension() string
		//Part of the import process. This method is run after [method _parse_image_data] and before [method _generate_scene_node].
		//Runs when parsing the texture JSON from the GLTF textures array. This can be used to set the source image index to use as the texture.
		ParseTextureJson(state gdclass.GLTFState, texture_json gd.Dictionary, ret_gltf_texture gdclass.GLTFTexture) gd.Error
		//Part of the import process. This method is run after [method _import_post_parse] and before [method _import_node].
		//Runs when generating a Godot scene node from a GLTFNode. The returned node will be added to the scene tree. Multiple nodes can be generated in this step if they are added as a child of the returned node.
		//[b]Note:[/b] The [param scene_parent] parameter may be null if this is the single root node.
		GenerateSceneNode(state gdclass.GLTFState, gltf_node gdclass.GLTFNode, scene_parent gdclass.Node) gdclass.Node3D
		//Part of the import process. This method is run after [method _parse_node_extensions] and before [method _generate_scene_node].
		//This method can be used to modify any of the data imported so far after parsing, before generating the nodes and then running the final per-node import step.
		ImportPostParse(state gdclass.GLTFState) gd.Error
		//Part of the import process. This method is run after [method _generate_scene_node] and before [method _import_post].
		//This method can be used to make modifications to each of the generated Godot scene nodes.
		ImportNode(state gdclass.GLTFState, gltf_node gdclass.GLTFNode, json gd.Dictionary, node gdclass.Node) gd.Error
		//Part of the import process. This method is run last, after all other parts of the import process.
		//This method can be used to modify the final Godot scene generated by the import process.
		ImportPost(state gdclass.GLTFState, root gdclass.Node) gd.Error
		//Part of the export process. This method is run first, before all other parts of the export process.
		//The return value is used to determine if this [GLTFDocumentExtension] instance should be used for exporting a given GLTF file. If [constant OK], the export will use this [GLTFDocumentExtension] instance. If not overridden, [constant OK] is returned.
		ExportPreflight(state gdclass.GLTFState, root gdclass.Node) gd.Error
		//Part of the export process. This method is run after [method _export_preflight] and before [method _export_preserialize].
		//Runs when converting the data from a Godot scene node. This method can be used to process the Godot scene node data into a format that can be used by [method _export_node].
		ConvertSceneNode(state gdclass.GLTFState, gltf_node gdclass.GLTFNode, scene_node gdclass.Node)
		//Part of the export process. This method is run after [method _convert_scene_node] and before [method _get_saveable_image_formats].
		//This method can be used to alter the state before performing serialization. It runs every time when generating a buffer with [method GLTFDocument.generate_buffer] or writing to the file system with [method GLTFDocument.write_to_filesystem].
		ExportPreserialize(state gdclass.GLTFState) gd.Error
		//Part of the export process. This method is run after [method _convert_scene_node] and before [method _export_node].
		//Returns an array of the image formats that can be saved/exported by this extension. This extension will only be selected as the image exporter if the [GLTFDocument]'s [member GLTFDocument.image_format] is in this array. If this [GLTFDocumentExtension] is selected as the image exporter, one of the [method _save_image_at_path] or [method _serialize_image_to_bytes] methods will run next, otherwise [method _export_node] will run next. If the format name contains [code]"Lossy"[/code], the lossy quality slider will be displayed.
		GetSaveableImageFormats() []string
		//Part of the export process. This method is run after [method _get_saveable_image_formats] and before [method _serialize_texture_json].
		//This method is run when embedding images in the GLTF file. When images are saved separately, [method _save_image_at_path] runs instead. Note that these methods only run when this [GLTFDocumentExtension] is selected as the image exporter.
		//This method must set the image MIME type in the [param image_dict] with the [code]"mimeType"[/code] key. For example, for a PNG image, it would be set to [code]"image/png"[/code]. The return value must be a [PackedByteArray] containing the image data.
		SerializeImageToBytes(state gdclass.GLTFState, image gdclass.Image, image_dict gd.Dictionary, image_format string, lossy_quality float64) []byte
		//Part of the export process. This method is run after [method _get_saveable_image_formats] and before [method _serialize_texture_json].
		//This method is run when saving images separately from the GLTF file. When images are embedded, [method _serialize_image_to_bytes] runs instead. Note that these methods only run when this [GLTFDocumentExtension] is selected as the image exporter.
		SaveImageAtPath(state gdclass.GLTFState, image gdclass.Image, file_path string, image_format string, lossy_quality float64) gd.Error
		//Part of the export process. This method is run after [method _save_image_at_path] or [method _serialize_image_to_bytes], and before [method _export_node]. Note that this method only runs when this [GLTFDocumentExtension] is selected as the image exporter.
		//This method can be used to set up the extensions for the texture JSON by editing [param texture_json]. The extension must also be added as used extension with [method GLTFState.add_used_extension], be sure to set [code]required[/code] to [code]true[/code] if you are not providing a fallback.
		SerializeTextureJson(state gdclass.GLTFState, texture_json gd.Dictionary, gltf_texture gdclass.GLTFTexture, image_format string) gd.Error
		//Part of the export process. This method is run after [method _get_saveable_image_formats] and before [method _export_post]. If this [GLTFDocumentExtension] is used for exporting images, this runs after [method _serialize_texture_json].
		//This method can be used to modify the final JSON of each node. Data should be primarily stored in [param gltf_node] prior to serializing the JSON, but the original Godot [param node] is also provided if available. The node may be null if not available, such as when exporting GLTF data not generated from a Godot scene.
		ExportNode(state gdclass.GLTFState, gltf_node gdclass.GLTFNode, json gd.Dictionary, node gdclass.Node) gd.Error
		//Part of the export process. This method is run last, after all other parts of the export process.
		//This method can be used to modify the final JSON of the generated GLTF file.
		ExportPost(state gdclass.GLTFState) gd.Error
	}
*/
type Instance [1]classdb.GLTFDocumentExtension

/*
Part of the import process. This method is run first, before all other parts of the import process.
The return value is used to determine if this [GLTFDocumentExtension] instance should be used for importing a given GLTF file. If [constant OK], the import will use this [GLTFDocumentExtension] instance. If not overridden, [constant OK] is returned.
*/
func (Instance) _import_preflight(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, extensions []string) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var extensions = pointers.New[gd.PackedStringArray](gd.UnsafeGet[[2]uintptr](p_args, 1))
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
func (Instance) _parse_node_extensions(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, gltf_node gdclass.GLTFNode, extensions gd.Dictionary) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var gltf_node = gdclass.GLTFNode{pointers.New[classdb.GLTFNode]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(gltf_node[0])
		var extensions = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 2))
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
func (Instance) _parse_image_data(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, image_data []byte, mime_type string, ret_image gdclass.Image) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var image_data = pointers.New[gd.PackedByteArray](gd.UnsafeGet[[2]uintptr](p_args, 1))
		defer pointers.End(image_data)
		var mime_type = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 2))
		defer pointers.End(mime_type)
		var ret_image = gdclass.Image{pointers.New[classdb.Image]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 3)})}
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
func (Instance) _parse_texture_json(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, texture_json gd.Dictionary, ret_gltf_texture gdclass.GLTFTexture) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var texture_json = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(texture_json)
		var ret_gltf_texture = gdclass.GLTFTexture{pointers.New[classdb.GLTFTexture]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 2)})}
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
func (Instance) _generate_scene_node(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, gltf_node gdclass.GLTFNode, scene_parent gdclass.Node) gdclass.Node3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var gltf_node = gdclass.GLTFNode{pointers.New[classdb.GLTFNode]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(gltf_node[0])
		var scene_parent = gdclass.Node{pointers.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 2)})}
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
func (Instance) _import_post_parse(impl func(ptr unsafe.Pointer, state gdclass.GLTFState) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
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
func (Instance) _import_node(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, gltf_node gdclass.GLTFNode, json gd.Dictionary, node gdclass.Node) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var gltf_node = gdclass.GLTFNode{pointers.New[classdb.GLTFNode]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(gltf_node[0])
		var json = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 2))
		defer pointers.End(json)
		var node = gdclass.Node{pointers.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 3)})}
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
func (Instance) _import_post(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, root gdclass.Node) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var root = gdclass.Node{pointers.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
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
func (Instance) _export_preflight(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, root gdclass.Node) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var root = gdclass.Node{pointers.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
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
func (Instance) _convert_scene_node(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, gltf_node gdclass.GLTFNode, scene_node gdclass.Node)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var gltf_node = gdclass.GLTFNode{pointers.New[classdb.GLTFNode]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(gltf_node[0])
		var scene_node = gdclass.Node{pointers.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 2)})}
		defer pointers.End(scene_node[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, state, gltf_node, scene_node)
	}
}

/*
Part of the export process. This method is run after [method _convert_scene_node] and before [method _get_saveable_image_formats].
This method can be used to alter the state before performing serialization. It runs every time when generating a buffer with [method GLTFDocument.generate_buffer] or writing to the file system with [method GLTFDocument.write_to_filesystem].
*/
func (Instance) _export_preserialize(impl func(ptr unsafe.Pointer, state gdclass.GLTFState) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
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
func (Instance) _serialize_image_to_bytes(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, image gdclass.Image, image_dict gd.Dictionary, image_format string, lossy_quality float64) []byte) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var image = gdclass.Image{pointers.New[classdb.Image]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(image[0])
		var image_dict = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 2))
		defer pointers.End(image_dict)
		var image_format = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 3))
		defer pointers.End(image_format)
		var lossy_quality = gd.UnsafeGet[gd.Float](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, image, image_dict, image_format.String(), float64(lossy_quality))
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
func (Instance) _save_image_at_path(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, image gdclass.Image, file_path string, image_format string, lossy_quality float64) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var image = gdclass.Image{pointers.New[classdb.Image]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(image[0])
		var file_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 2))
		defer pointers.End(file_path)
		var image_format = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 3))
		defer pointers.End(image_format)
		var lossy_quality = gd.UnsafeGet[gd.Float](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, image, file_path.String(), image_format.String(), float64(lossy_quality))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the export process. This method is run after [method _save_image_at_path] or [method _serialize_image_to_bytes], and before [method _export_node]. Note that this method only runs when this [GLTFDocumentExtension] is selected as the image exporter.
This method can be used to set up the extensions for the texture JSON by editing [param texture_json]. The extension must also be added as used extension with [method GLTFState.add_used_extension], be sure to set [code]required[/code] to [code]true[/code] if you are not providing a fallback.
*/
func (Instance) _serialize_texture_json(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, texture_json gd.Dictionary, gltf_texture gdclass.GLTFTexture, image_format string) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var texture_json = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(texture_json)
		var gltf_texture = gdclass.GLTFTexture{pointers.New[classdb.GLTFTexture]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 2)})}
		defer pointers.End(gltf_texture[0])
		var image_format = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 3))
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
func (Instance) _export_node(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, gltf_node gdclass.GLTFNode, json gd.Dictionary, node gdclass.Node) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var gltf_node = gdclass.GLTFNode{pointers.New[classdb.GLTFNode]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(gltf_node[0])
		var json = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 2))
		defer pointers.End(json)
		var node = gdclass.Node{pointers.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 3)})}
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
func (Instance) _export_post(impl func(ptr unsafe.Pointer, state gdclass.GLTFState) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state)
		gd.UnsafeSet(p_back, ret)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.GLTFDocumentExtension

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFDocumentExtension"))
	return Instance{classdb.GLTFDocumentExtension(object)}
}

/*
Part of the import process. This method is run first, before all other parts of the import process.
The return value is used to determine if this [GLTFDocumentExtension] instance should be used for importing a given GLTF file. If [constant OK], the import will use this [GLTFDocumentExtension] instance. If not overridden, [constant OK] is returned.
*/
func (class) _import_preflight(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, extensions gd.PackedStringArray) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var extensions = pointers.New[gd.PackedStringArray](gd.UnsafeGet[[2]uintptr](p_args, 1))
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
func (class) _parse_node_extensions(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, gltf_node gdclass.GLTFNode, extensions gd.Dictionary) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var gltf_node = gdclass.GLTFNode{pointers.New[classdb.GLTFNode]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(gltf_node[0])
		var extensions = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, gltf_node, extensions)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the import process. This method is run after [method _parse_node_extensions] and before [method _parse_texture_json].
Runs when parsing image data from a GLTF file. The data could be sourced from a separate file, a URI, or a buffer, and then is passed as a byte array.
*/
func (class) _parse_image_data(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, image_data gd.PackedByteArray, mime_type gd.String, ret_image gdclass.Image) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var image_data = pointers.New[gd.PackedByteArray](gd.UnsafeGet[[2]uintptr](p_args, 1))
		var mime_type = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 2))
		var ret_image = gdclass.Image{pointers.New[classdb.Image]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 3)})}
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
func (class) _parse_texture_json(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, texture_json gd.Dictionary, ret_gltf_texture gdclass.GLTFTexture) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var texture_json = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 1))
		var ret_gltf_texture = gdclass.GLTFTexture{pointers.New[classdb.GLTFTexture]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 2)})}
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
func (class) _generate_scene_node(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, gltf_node gdclass.GLTFNode, scene_parent gdclass.Node) gdclass.Node3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var gltf_node = gdclass.GLTFNode{pointers.New[classdb.GLTFNode]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(gltf_node[0])
		var scene_parent = gdclass.Node{pointers.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 2)})}
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
func (class) _import_post_parse(impl func(ptr unsafe.Pointer, state gdclass.GLTFState) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
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
func (class) _import_node(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, gltf_node gdclass.GLTFNode, json gd.Dictionary, node gdclass.Node) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var gltf_node = gdclass.GLTFNode{pointers.New[classdb.GLTFNode]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(gltf_node[0])
		var json = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 2))
		var node = gdclass.Node{pointers.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 3)})}
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
func (class) _import_post(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, root gdclass.Node) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var root = gdclass.Node{pointers.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
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
func (class) _export_preflight(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, root gdclass.Node) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var root = gdclass.Node{pointers.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
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
func (class) _convert_scene_node(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, gltf_node gdclass.GLTFNode, scene_node gdclass.Node)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var gltf_node = gdclass.GLTFNode{pointers.New[classdb.GLTFNode]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(gltf_node[0])
		var scene_node = gdclass.Node{pointers.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 2)})}
		defer pointers.End(scene_node[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, state, gltf_node, scene_node)
	}
}

/*
Part of the export process. This method is run after [method _convert_scene_node] and before [method _get_saveable_image_formats].
This method can be used to alter the state before performing serialization. It runs every time when generating a buffer with [method GLTFDocument.generate_buffer] or writing to the file system with [method GLTFDocument.write_to_filesystem].
*/
func (class) _export_preserialize(impl func(ptr unsafe.Pointer, state gdclass.GLTFState) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
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
func (class) _serialize_image_to_bytes(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, image gdclass.Image, image_dict gd.Dictionary, image_format gd.String, lossy_quality gd.Float) gd.PackedByteArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var image = gdclass.Image{pointers.New[classdb.Image]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(image[0])
		var image_dict = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 2))
		var image_format = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 3))
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
func (class) _save_image_at_path(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, image gdclass.Image, file_path gd.String, image_format gd.String, lossy_quality gd.Float) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var image = gdclass.Image{pointers.New[classdb.Image]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(image[0])
		var file_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 2))
		var image_format = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 3))
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
func (class) _serialize_texture_json(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, texture_json gd.Dictionary, gltf_texture gdclass.GLTFTexture, image_format gd.String) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var texture_json = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 1))
		var gltf_texture = gdclass.GLTFTexture{pointers.New[classdb.GLTFTexture]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 2)})}
		defer pointers.End(gltf_texture[0])
		var image_format = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 3))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, state, texture_json, gltf_texture, image_format)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Part of the export process. This method is run after [method _get_saveable_image_formats] and before [method _export_post]. If this [GLTFDocumentExtension] is used for exporting images, this runs after [method _serialize_texture_json].
This method can be used to modify the final JSON of each node. Data should be primarily stored in [param gltf_node] prior to serializing the JSON, but the original Godot [param node] is also provided if available. The node may be null if not available, such as when exporting GLTF data not generated from a Godot scene.
*/
func (class) _export_node(impl func(ptr unsafe.Pointer, state gdclass.GLTFState, gltf_node gdclass.GLTFNode, json gd.Dictionary, node gdclass.Node) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(state[0])
		var gltf_node = gdclass.GLTFNode{pointers.New[classdb.GLTFNode]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(gltf_node[0])
		var json = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 2))
		var node = gdclass.Node{pointers.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 3)})}
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
func (class) _export_post(impl func(ptr unsafe.Pointer, state gdclass.GLTFState) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.GLTFState{pointers.New[classdb.GLTFState]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

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
		return gd.VirtualByName(self.AsResource(), name)
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
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("GLTFDocumentExtension", func(ptr gd.Object) any { return classdb.GLTFDocumentExtension(ptr) })
}
