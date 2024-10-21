package EditorImportPlugin

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/ResourceImporter"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
	// EditorImportPlugin methods that can be overridden by a [Class] that extends it.
	type EditorImportPlugin interface {
		//Gets the unique name of the importer.
		GetImporterName() gd.String
		//Gets the name to display in the import window. You should choose this name as a continuation to "Import as", e.g. "Import as Special Mesh".
		GetVisibleName() gd.String
		//Gets the number of initial presets defined by the plugin. Use [method _get_import_options] to get the default options for the preset and [method _get_preset_name] to get the name of the preset.
		GetPresetCount() gd.Int
		//Gets the name of the options preset at this index.
		GetPresetName(preset_index gd.Int) gd.String
		//Gets the list of file extensions to associate with this loader (case-insensitive). e.g. [code]["obj"][/code].
		GetRecognizedExtensions() gd.PackedStringArray
		//Gets the options and default values for the preset at this index. Returns an Array of Dictionaries with the following keys: [code]name[/code], [code]default_value[/code], [code]property_hint[/code] (optional), [code]hint_string[/code] (optional), [code]usage[/code] (optional).
		GetImportOptions(path gd.String, preset_index gd.Int) gd.ArrayOf[gd.Dictionary]
		//Gets the extension used to save this resource in the [code].godot/imported[/code] directory (see [member ProjectSettings.application/config/use_hidden_project_data_directory]).
		GetSaveExtension() gd.String
		//Gets the Godot resource type associated with this loader. e.g. [code]"Mesh"[/code] or [code]"Animation"[/code].
		GetResourceType() gd.String
		//Gets the priority of this plugin for the recognized extension. Higher priority plugins will be preferred. The default priority is [code]1.0[/code].
		GetPriority() gd.Float
		//Gets the order of this importer to be run when importing resources. Importers with [i]lower[/i] import orders will be called first, and higher values will be called later. Use this to ensure the importer runs after the dependencies are already imported. The default import order is [code]0[/code] unless overridden by a specific importer. See [enum ResourceImporter.ImportOrder] for some predefined values.
		GetImportOrder() gd.Int
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
		GetOptionVisibility(path gd.String, option_name gd.StringName, options gd.Dictionary) bool
		//Imports [param source_file] into [param save_path] with the import [param options] specified. The [param platform_variants] and [param gen_files] arrays will be modified by this function.
		//This method must be overridden to do the actual importing work. See this class' description for an example of overriding this method.
		Import(source_file gd.String, save_path gd.String, options gd.Dictionary, platform_variants gd.ArrayOf[gd.String], gen_files gd.ArrayOf[gd.String]) int64
		//Tells whether this importer can be run in parallel on threads, or, on the contrary, it's only safe for the editor to call it from the main thread, for one file at a time.
		//If this method is not overridden, it will return [code]true[/code] by default (i.e., safe for parallel importing).
		CanImportThreaded() bool
	}

*/
type Simple [1]classdb.EditorImportPlugin
func (self Simple) AppendImportExternalResource(path string, custom_options gd.Dictionary, custom_importer string, generator_parameters gd.Variant) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).AppendImportExternalResource(gc.String(path), custom_options, gc.String(custom_importer), generator_parameters))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorImportPlugin
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Gets the unique name of the importer.
*/
func (class) _get_importer_name(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Gets the name to display in the import window. You should choose this name as a continuation to "Import as", e.g. "Import as Special Mesh".
*/
func (class) _get_visible_name(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Gets the number of initial presets defined by the plugin. Use [method _get_import_options] to get the default options for the preset and [method _get_preset_name] to get the name of the preset.
*/
func (class) _get_preset_count(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Gets the name of the options preset at this index.
*/
func (class) _get_preset_name(impl func(ptr unsafe.Pointer, preset_index gd.Int) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var preset_index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset_index)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Gets the list of file extensions to associate with this loader (case-insensitive). e.g. [code]["obj"][/code].
*/
func (class) _get_recognized_extensions(impl func(ptr unsafe.Pointer) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Gets the options and default values for the preset at this index. Returns an Array of Dictionaries with the following keys: [code]name[/code], [code]default_value[/code], [code]property_hint[/code] (optional), [code]hint_string[/code] (optional), [code]usage[/code] (optional).
*/
func (class) _get_import_options(impl func(ptr unsafe.Pointer, path gd.String, preset_index gd.Int) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var preset_index = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, preset_index)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

/*
Gets the extension used to save this resource in the [code].godot/imported[/code] directory (see [member ProjectSettings.application/config/use_hidden_project_data_directory]).
*/
func (class) _get_save_extension(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Gets the Godot resource type associated with this loader. e.g. [code]"Mesh"[/code] or [code]"Animation"[/code].
*/
func (class) _get_resource_type(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Gets the priority of this plugin for the recognized extension. Higher priority plugins will be preferred. The default priority is [code]1.0[/code].
*/
func (class) _get_priority(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Gets the order of this importer to be run when importing resources. Importers with [i]lower[/i] import orders will be called first, and higher values will be called later. Use this to ensure the importer runs after the dependencies are already imported. The default import order is [code]0[/code] unless overridden by a specific importer. See [enum ResourceImporter.ImportOrder] for some predefined values.
*/
func (class) _get_import_order(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
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
func (class) _get_option_visibility(impl func(ptr unsafe.Pointer, path gd.String, option_name gd.StringName, options gd.Dictionary) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var option_name = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var options = mmm.Let[gd.Dictionary](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, option_name, options)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Imports [param source_file] into [param save_path] with the import [param options] specified. The [param platform_variants] and [param gen_files] arrays will be modified by this function.
This method must be overridden to do the actual importing work. See this class' description for an example of overriding this method.
*/
func (class) _import(impl func(ptr unsafe.Pointer, source_file gd.String, save_path gd.String, options gd.Dictionary, platform_variants gd.ArrayOf[gd.String], gen_files gd.ArrayOf[gd.String]) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var source_file = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var save_path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var options = mmm.Let[gd.Dictionary](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,2))
		var platform_variants = gd.TypedArray[gd.String](mmm.Let[gd.Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,3)))
		var gen_files = gd.TypedArray[gd.String](mmm.Let[gd.Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,4)))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, source_file, save_path, options, platform_variants, gen_files)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Tells whether this importer can be run in parallel on threads, or, on the contrary, it's only safe for the editor to call it from the main thread, for one file at a time.
If this method is not overridden, it will return [code]true[/code] by default (i.e., safe for parallel importing).
*/
func (class) _can_import_threaded(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
This function can only be called during the [method _import] callback and it allows manually importing resources from it. This is useful when the imported file generates external resources that require importing (as example, images). Custom parameters for the ".import" file can be passed via the [param custom_options]. Additionally, in cases where multiple importers can handle a file, the [param custom_importer] can be specified to force a specific one. This function performs a resource import and returns immediately with a success or error code. [param generator_parameters] defines optional extra metadata which will be stored as [code skip-lint]generator_parameters[/code] in the [code]remap[/code] section of the [code].import[/code] file, for example to store a md5 hash of the source data.
*/
//go:nosplit
func (self class) AppendImportExternalResource(path gd.String, custom_options gd.Dictionary, custom_importer gd.String, generator_parameters gd.Variant) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, mmm.Get(custom_options))
	callframe.Arg(frame, mmm.Get(custom_importer))
	callframe.Arg(frame, mmm.Get(generator_parameters))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorImportPlugin.Bind_append_import_external_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsEditorImportPlugin() Expert { return self[0].AsEditorImportPlugin() }


//go:nosplit
func (self Simple) AsEditorImportPlugin() Simple { return self[0].AsEditorImportPlugin() }


//go:nosplit
func (self class) AsResourceImporter() ResourceImporter.Expert { return self[0].AsResourceImporter() }


//go:nosplit
func (self Simple) AsResourceImporter() ResourceImporter.Simple { return self[0].AsResourceImporter() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_importer_name": return reflect.ValueOf(self._get_importer_name);
	case "_get_visible_name": return reflect.ValueOf(self._get_visible_name);
	case "_get_preset_count": return reflect.ValueOf(self._get_preset_count);
	case "_get_preset_name": return reflect.ValueOf(self._get_preset_name);
	case "_get_recognized_extensions": return reflect.ValueOf(self._get_recognized_extensions);
	case "_get_import_options": return reflect.ValueOf(self._get_import_options);
	case "_get_save_extension": return reflect.ValueOf(self._get_save_extension);
	case "_get_resource_type": return reflect.ValueOf(self._get_resource_type);
	case "_get_priority": return reflect.ValueOf(self._get_priority);
	case "_get_import_order": return reflect.ValueOf(self._get_import_order);
	case "_get_option_visibility": return reflect.ValueOf(self._get_option_visibility);
	case "_import": return reflect.ValueOf(self._import);
	case "_can_import_threaded": return reflect.ValueOf(self._can_import_threaded);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("EditorImportPlugin", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
