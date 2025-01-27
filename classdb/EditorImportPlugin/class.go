// Package EditorImportPlugin provides methods for working with EditorImportPlugin object instances.
package EditorImportPlugin

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
import "graphics.gd/variant/String"
import "graphics.gd/classdb/ResourceImporter"
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
var _ String.Readable

/*
[EditorImportPlugin]s provide a way to extend the editor's resource import functionality. Use them to import resources from custom files or to provide alternatives to the editor's existing importers.
EditorImportPlugins work by associating with specific file extensions and a resource type. See [method _get_recognized_extensions] and [method _get_resource_type]. They may optionally specify some import presets that affect the import process. EditorImportPlugins are responsible for creating the resources and saving them in the [code].godot/imported[/code] directory (see [member ProjectSettings.application/config/use_hidden_project_data_directory]).
Below is an example EditorImportPlugin that imports a [Mesh] from a file with the extension ".special" or ".spec":
[codeblocks]
[gdscript]
@tool
extends EditorImportPlugin

func _get_importer_name():

	return "my.special.plugin"

func _get_visible_name():

	return "Special Mesh"

func _get_recognized_extensions():

	return ["special", "spec"]

func _get_save_extension():

	return "mesh"

func _get_resource_type():

	return "Mesh"

func _get_preset_count():

	return 1

func _get_preset_name(preset_index):

	return "Default"

func _get_import_options(path, preset_index):

	return [{"name": "my_option", "default_value": false}]

func _import(source_file, save_path, options, platform_variants, gen_files):

	var file = FileAccess.open(source_file, FileAccess.READ)
	if file == null:
	    return FAILED
	var mesh = ArrayMesh.new()
	# Fill the Mesh with data read in "file", left as an exercise to the reader.

	var filename = save_path + "." + _get_save_extension()
	return ResourceSaver.save(mesh, filename)

[/gdscript]
[csharp]
using Godot;

public partial class MySpecialPlugin : EditorImportPlugin

	{
	    public override string _GetImporterName()
	    {
	        return "my.special.plugin";
	    }

	    public override string _GetVisibleName()
	    {
	        return "Special Mesh";
	    }

	    public override string[] _GetRecognizedExtensions()
	    {
	        return new string[] { "special", "spec" };
	    }

	    public override string _GetSaveExtension()
	    {
	        return "mesh";
	    }

	    public override string _GetResourceType()
	    {
	        return "Mesh";
	    }

	    public override int _GetPresetCount()
	    {
	        return 1;
	    }

	    public override string _GetPresetName(int presetIndex)
	    {
	        return "Default";
	    }

	    public override Godot.Collections.Array<Godot.Collections.Dictionary> _GetImportOptions(string path, int presetIndex)
	    {
	        return new Godot.Collections.Array<Godot.Collections.Dictionary>
	        {
	            new Godot.Collections.Dictionary
	            {
	                { "name", "myOption" },
	                { "default_value", false },
	            }
	        };
	    }

	    public override Error _Import(string sourceFile, string savePath, Godot.Collections.Dictionary options, Godot.Collections.Array<string> platformVariants, Godot.Collections.Array<string> genFiles)
	    {
	        using var file = FileAccess.Open(sourceFile, FileAccess.ModeFlags.Read);
	        if (file.GetError() != Error.Ok)
	        {
	            return Error.Failed;
	        }

	        var mesh = new ArrayMesh();
	        // Fill the Mesh with data read in "file", left as an exercise to the reader.
	        string filename = $"{savePath}.{_GetSaveExtension()}";
	        return ResourceSaver.Save(mesh, filename);
	    }
	}

[/csharp]
[/codeblocks]
To use [EditorImportPlugin], register it using the [method EditorPlugin.add_import_plugin] method first.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=EditorImportPlugin)
*/
type Instance [1]gdclass.EditorImportPlugin

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorImportPlugin() Instance
}
type Interface interface {
	//Gets the unique name of the importer.
	GetImporterName() string
	//Gets the name to display in the import window. You should choose this name as a continuation to "Import as", e.g. "Import as Special Mesh".
	GetVisibleName() string
	//Gets the number of initial presets defined by the plugin. Use [method _get_import_options] to get the default options for the preset and [method _get_preset_name] to get the name of the preset.
	GetPresetCount() int
	//Gets the name of the options preset at this index.
	GetPresetName(preset_index int) string
	//Gets the list of file extensions to associate with this loader (case-insensitive). e.g. [code]["obj"][/code].
	GetRecognizedExtensions() []string
	//Gets the options and default values for the preset at this index. Returns an Array of Dictionaries with the following keys: [code]name[/code], [code]default_value[/code], [code]property_hint[/code] (optional), [code]hint_string[/code] (optional), [code]usage[/code] (optional).
	GetImportOptions(path string, preset_index int) []map[any]any
	//Gets the extension used to save this resource in the [code].godot/imported[/code] directory (see [member ProjectSettings.application/config/use_hidden_project_data_directory]).
	GetSaveExtension() string
	//Gets the Godot resource type associated with this loader. e.g. [code]"Mesh"[/code] or [code]"Animation"[/code].
	GetResourceType() string
	//Gets the priority of this plugin for the recognized extension. Higher priority plugins will be preferred. The default priority is [code]1.0[/code].
	GetPriority() Float.X
	//Gets the order of this importer to be run when importing resources. Importers with [i]lower[/i] import orders will be called first, and higher values will be called later. Use this to ensure the importer runs after the dependencies are already imported. The default import order is [code]0[/code] unless overridden by a specific importer. See [enum ResourceImporter.ImportOrder] for some predefined values.
	GetImportOrder() int
	//This method can be overridden to hide specific import options if conditions are met. This is mainly useful for hiding options that depend on others if one of them is disabled. For example:
	//[codeblocks]
	//[gdscript]
	//func _get_option_visibility(option, options):
	//    # Only show the lossy quality setting if the compression mode is set to "Lossy".
	//    if option == "compress/lossy_quality" and options.has("compress/mode"):
	//        return int(options["compress/mode"]) == COMPRESS_LOSSY # This is a constant that you set
	//
	//    return true
	//[/gdscript]
	//[csharp]
	//public void _GetOptionVisibility(string option, Godot.Collections.Dictionary options)
	//{
	//    // Only show the lossy quality setting if the compression mode is set to "Lossy".
	//    if (option == "compress/lossy_quality" && options.ContainsKey("compress/mode"))
	//    {
	//        return (int)options["compress/mode"] == CompressLossy; // This is a constant you set
	//    }
	//
	//    return true;
	//}
	//[/csharp]
	//[/codeblocks]
	//Returns [code]true[/code] to make all options always visible.
	GetOptionVisibility(path string, option_name string, options map[any]any) bool
	//Imports [param source_file] into [param save_path] with the import [param options] specified. The [param platform_variants] and [param gen_files] arrays will be modified by this function.
	//This method must be overridden to do the actual importing work. See this class' description for an example of overriding this method.
	Import(source_file string, save_path string, options map[any]any, platform_variants []string, gen_files []string) error
	//Tells whether this importer can be run in parallel on threads, or, on the contrary, it's only safe for the editor to call it from the main thread, for one file at a time.
	//If this method is not overridden, it will return [code]true[/code] by default (i.e., safe for parallel importing).
	CanImportThreaded() bool
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetImporterName() (_ string)                                      { return }
func (self implementation) GetVisibleName() (_ string)                                       { return }
func (self implementation) GetPresetCount() (_ int)                                          { return }
func (self implementation) GetPresetName(preset_index int) (_ string)                        { return }
func (self implementation) GetRecognizedExtensions() (_ []string)                            { return }
func (self implementation) GetImportOptions(path string, preset_index int) (_ []map[any]any) { return }
func (self implementation) GetSaveExtension() (_ string)                                     { return }
func (self implementation) GetResourceType() (_ string)                                      { return }
func (self implementation) GetPriority() (_ Float.X)                                         { return }
func (self implementation) GetImportOrder() (_ int)                                          { return }
func (self implementation) GetOptionVisibility(path string, option_name string, options map[any]any) (_ bool) {
	return
}
func (self implementation) Import(source_file string, save_path string, options map[any]any, platform_variants []string, gen_files []string) (_ error) {
	return
}
func (self implementation) CanImportThreaded() (_ bool) { return }

/*
Gets the unique name of the importer.
*/
func (Instance) _get_importer_name(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Gets the name to display in the import window. You should choose this name as a continuation to "Import as", e.g. "Import as Special Mesh".
*/
func (Instance) _get_visible_name(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Gets the number of initial presets defined by the plugin. Use [method _get_import_options] to get the default options for the preset and [method _get_preset_name] to get the name of the preset.
*/
func (Instance) _get_preset_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Gets the name of the options preset at this index.
*/
func (Instance) _get_preset_name(impl func(ptr unsafe.Pointer, preset_index int) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset_index = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(preset_index))
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Gets the list of file extensions to associate with this loader (case-insensitive). e.g. [code]["obj"][/code].
*/
func (Instance) _get_recognized_extensions(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
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
Gets the options and default values for the preset at this index. Returns an Array of Dictionaries with the following keys: [code]name[/code], [code]default_value[/code], [code]property_hint[/code] (optional), [code]hint_string[/code] (optional), [code]usage[/code] (optional).
*/
func (Instance) _get_import_options(impl func(ptr unsafe.Pointer, path string, preset_index int) []map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		var preset_index = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), int(preset_index))
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[Dictionary.Any]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Gets the extension used to save this resource in the [code].godot/imported[/code] directory (see [member ProjectSettings.application/config/use_hidden_project_data_directory]).
*/
func (Instance) _get_save_extension(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Gets the Godot resource type associated with this loader. e.g. [code]"Mesh"[/code] or [code]"Animation"[/code].
*/
func (Instance) _get_resource_type(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Gets the priority of this plugin for the recognized extension. Higher priority plugins will be preferred. The default priority is [code]1.0[/code].
*/
func (Instance) _get_priority(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Gets the order of this importer to be run when importing resources. Importers with [i]lower[/i] import orders will be called first, and higher values will be called later. Use this to ensure the importer runs after the dependencies are already imported. The default import order is [code]0[/code] unless overridden by a specific importer. See [enum ResourceImporter.ImportOrder] for some predefined values.
*/
func (Instance) _get_import_order(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
This method can be overridden to hide specific import options if conditions are met. This is mainly useful for hiding options that depend on others if one of them is disabled. For example:
[codeblocks]
[gdscript]
func _get_option_visibility(option, options):

	# Only show the lossy quality setting if the compression mode is set to "Lossy".
	if option == "compress/lossy_quality" and options.has("compress/mode"):
	    return int(options["compress/mode"]) == COMPRESS_LOSSY # This is a constant that you set

	return true

[/gdscript]
[csharp]
public void _GetOptionVisibility(string option, Godot.Collections.Dictionary options)

	{
	    // Only show the lossy quality setting if the compression mode is set to "Lossy".
	    if (option == "compress/lossy_quality" && options.ContainsKey("compress/mode"))
	    {
	        return (int)options["compress/mode"] == CompressLossy; // This is a constant you set
	    }

	    return true;
	}

[/csharp]
[/codeblocks]
Returns [code]true[/code] to make all options always visible.
*/
func (Instance) _get_option_visibility(impl func(ptr unsafe.Pointer, path string, option_name string, options map[any]any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		var option_name = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		defer pointers.End(option_name)
		var options = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalDictionary(options))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), option_name.String(), gd.DictionaryAs[map[any]any](options))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Imports [param source_file] into [param save_path] with the import [param options] specified. The [param platform_variants] and [param gen_files] arrays will be modified by this function.
This method must be overridden to do the actual importing work. See this class' description for an example of overriding this method.
*/
func (Instance) _import(impl func(ptr unsafe.Pointer, source_file string, save_path string, options map[any]any, platform_variants []string, gen_files []string) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var source_file = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(source_file))
		var save_path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(save_path))
		var options = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalDictionary(options))
		var platform_variants = Array.Through(gd.ArrayProxy[String.Readable]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 3))))
		defer pointers.End(gd.InternalArray(platform_variants))
		var gen_files = Array.Through(gd.ArrayProxy[String.Readable]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 4))))
		defer pointers.End(gd.InternalArray(gen_files))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, source_file.String(), save_path.String(), gd.DictionaryAs[map[any]any](options), gd.ArrayAs[[]string](gd.InternalArray(platform_variants)), gd.ArrayAs[[]string](gd.InternalArray(gen_files)))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Tells whether this importer can be run in parallel on threads, or, on the contrary, it's only safe for the editor to call it from the main thread, for one file at a time.
If this method is not overridden, it will return [code]true[/code] by default (i.e., safe for parallel importing).
*/
func (Instance) _can_import_threaded(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
This function can only be called during the [method _import] callback and it allows manually importing resources from it. This is useful when the imported file generates external resources that require importing (as example, images). Custom parameters for the ".import" file can be passed via the [param custom_options]. Additionally, in cases where multiple importers can handle a file, the [param custom_importer] can be specified to force a specific one. This function performs a resource import and returns immediately with a success or error code. [param generator_parameters] defines optional extra metadata which will be stored as [code skip-lint]generator_parameters[/code] in the [code]remap[/code] section of the [code].import[/code] file, for example to store a md5 hash of the source data.
*/
func (self Instance) AppendImportExternalResource(path string) error { //gd:EditorImportPlugin.append_import_external_resource
	return error(gd.ToError(class(self).AppendImportExternalResource(String.New(path), Dictionary.Nil, String.New(""), gd.NewVariant(gd.NewVariant(([1]any{}[0]))))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorImportPlugin

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorImportPlugin"))
	casted := Instance{*(*gdclass.EditorImportPlugin)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Gets the unique name of the importer.
*/
func (class) _get_importer_name(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Gets the name to display in the import window. You should choose this name as a continuation to "Import as", e.g. "Import as Special Mesh".
*/
func (class) _get_visible_name(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Gets the number of initial presets defined by the plugin. Use [method _get_import_options] to get the default options for the preset and [method _get_preset_name] to get the name of the preset.
*/
func (class) _get_preset_count(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Gets the name of the options preset at this index.
*/
func (class) _get_preset_name(impl func(ptr unsafe.Pointer, preset_index gd.Int) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset_index = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset_index)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Gets the list of file extensions to associate with this loader (case-insensitive). e.g. [code]["obj"][/code].
*/
func (class) _get_recognized_extensions(impl func(ptr unsafe.Pointer) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
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
Gets the options and default values for the preset at this index. Returns an Array of Dictionaries with the following keys: [code]name[/code], [code]default_value[/code], [code]property_hint[/code] (optional), [code]hint_string[/code] (optional), [code]usage[/code] (optional).
*/
func (class) _get_import_options(impl func(ptr unsafe.Pointer, path String.Readable, preset_index gd.Int) Array.Contains[Dictionary.Any]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		var preset_index = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, preset_index)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Gets the extension used to save this resource in the [code].godot/imported[/code] directory (see [member ProjectSettings.application/config/use_hidden_project_data_directory]).
*/
func (class) _get_save_extension(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Gets the Godot resource type associated with this loader. e.g. [code]"Mesh"[/code] or [code]"Animation"[/code].
*/
func (class) _get_resource_type(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Gets the priority of this plugin for the recognized extension. Higher priority plugins will be preferred. The default priority is [code]1.0[/code].
*/
func (class) _get_priority(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Gets the order of this importer to be run when importing resources. Importers with [i]lower[/i] import orders will be called first, and higher values will be called later. Use this to ensure the importer runs after the dependencies are already imported. The default import order is [code]0[/code] unless overridden by a specific importer. See [enum ResourceImporter.ImportOrder] for some predefined values.
*/
func (class) _get_import_order(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
This method can be overridden to hide specific import options if conditions are met. This is mainly useful for hiding options that depend on others if one of them is disabled. For example:
[codeblocks]
[gdscript]
func _get_option_visibility(option, options):

	# Only show the lossy quality setting if the compression mode is set to "Lossy".
	if option == "compress/lossy_quality" and options.has("compress/mode"):
	    return int(options["compress/mode"]) == COMPRESS_LOSSY # This is a constant that you set

	return true

[/gdscript]
[csharp]
public void _GetOptionVisibility(string option, Godot.Collections.Dictionary options)

	{
	    // Only show the lossy quality setting if the compression mode is set to "Lossy".
	    if (option == "compress/lossy_quality" && options.ContainsKey("compress/mode"))
	    {
	        return (int)options["compress/mode"] == CompressLossy; // This is a constant you set
	    }

	    return true;
	}

[/csharp]
[/codeblocks]
Returns [code]true[/code] to make all options always visible.
*/
func (class) _get_option_visibility(impl func(ptr unsafe.Pointer, path String.Readable, option_name gd.StringName, options Dictionary.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		var option_name = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		defer pointers.End(option_name)
		var options = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalDictionary(options))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, option_name, options)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Imports [param source_file] into [param save_path] with the import [param options] specified. The [param platform_variants] and [param gen_files] arrays will be modified by this function.
This method must be overridden to do the actual importing work. See this class' description for an example of overriding this method.
*/
func (class) _import(impl func(ptr unsafe.Pointer, source_file String.Readable, save_path String.Readable, options Dictionary.Any, platform_variants Array.Contains[String.Readable], gen_files Array.Contains[String.Readable]) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var source_file = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(source_file))
		var save_path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(save_path))
		var options = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalDictionary(options))
		var platform_variants = Array.Through(gd.ArrayProxy[String.Readable]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 3))))
		defer pointers.End(gd.InternalArray(platform_variants))
		var gen_files = Array.Through(gd.ArrayProxy[String.Readable]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 4))))
		defer pointers.End(gd.InternalArray(gen_files))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, source_file, save_path, options, platform_variants, gen_files)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Tells whether this importer can be run in parallel on threads, or, on the contrary, it's only safe for the editor to call it from the main thread, for one file at a time.
If this method is not overridden, it will return [code]true[/code] by default (i.e., safe for parallel importing).
*/
func (class) _can_import_threaded(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
This function can only be called during the [method _import] callback and it allows manually importing resources from it. This is useful when the imported file generates external resources that require importing (as example, images). Custom parameters for the ".import" file can be passed via the [param custom_options]. Additionally, in cases where multiple importers can handle a file, the [param custom_importer] can be specified to force a specific one. This function performs a resource import and returns immediately with a success or error code. [param generator_parameters] defines optional extra metadata which will be stored as [code skip-lint]generator_parameters[/code] in the [code]remap[/code] section of the [code].import[/code] file, for example to store a md5 hash of the source data.
*/
//go:nosplit
func (self class) AppendImportExternalResource(path String.Readable, custom_options Dictionary.Any, custom_importer String.Readable, generator_parameters gd.Variant) gd.Error { //gd:EditorImportPlugin.append_import_external_resource
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, pointers.Get(gd.InternalDictionary(custom_options)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(custom_importer)))
	callframe.Arg(frame, pointers.Get(generator_parameters))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorImportPlugin.Bind_append_import_external_resource, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsEditorImportPlugin() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorImportPlugin() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResourceImporter() ResourceImporter.Advanced {
	return *((*ResourceImporter.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResourceImporter() ResourceImporter.Instance {
	return *((*ResourceImporter.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_importer_name":
		return reflect.ValueOf(self._get_importer_name)
	case "_get_visible_name":
		return reflect.ValueOf(self._get_visible_name)
	case "_get_preset_count":
		return reflect.ValueOf(self._get_preset_count)
	case "_get_preset_name":
		return reflect.ValueOf(self._get_preset_name)
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	case "_get_import_options":
		return reflect.ValueOf(self._get_import_options)
	case "_get_save_extension":
		return reflect.ValueOf(self._get_save_extension)
	case "_get_resource_type":
		return reflect.ValueOf(self._get_resource_type)
	case "_get_priority":
		return reflect.ValueOf(self._get_priority)
	case "_get_import_order":
		return reflect.ValueOf(self._get_import_order)
	case "_get_option_visibility":
		return reflect.ValueOf(self._get_option_visibility)
	case "_import":
		return reflect.ValueOf(self._import)
	case "_can_import_threaded":
		return reflect.ValueOf(self._can_import_threaded)
	default:
		return gd.VirtualByName(ResourceImporter.Advanced(self.AsResourceImporter()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_importer_name":
		return reflect.ValueOf(self._get_importer_name)
	case "_get_visible_name":
		return reflect.ValueOf(self._get_visible_name)
	case "_get_preset_count":
		return reflect.ValueOf(self._get_preset_count)
	case "_get_preset_name":
		return reflect.ValueOf(self._get_preset_name)
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	case "_get_import_options":
		return reflect.ValueOf(self._get_import_options)
	case "_get_save_extension":
		return reflect.ValueOf(self._get_save_extension)
	case "_get_resource_type":
		return reflect.ValueOf(self._get_resource_type)
	case "_get_priority":
		return reflect.ValueOf(self._get_priority)
	case "_get_import_order":
		return reflect.ValueOf(self._get_import_order)
	case "_get_option_visibility":
		return reflect.ValueOf(self._get_option_visibility)
	case "_import":
		return reflect.ValueOf(self._import)
	case "_can_import_threaded":
		return reflect.ValueOf(self._can_import_threaded)
	default:
		return gd.VirtualByName(ResourceImporter.Instance(self.AsResourceImporter()), name)
	}
}
func init() {
	gdclass.Register("EditorImportPlugin", func(ptr gd.Object) any {
		return [1]gdclass.EditorImportPlugin{*(*gdclass.EditorImportPlugin)(unsafe.Pointer(&ptr))}
	})
}

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
