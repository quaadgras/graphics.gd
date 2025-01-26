// Package EditorExportPlugin provides methods for working with EditorExportPlugin object instances.
package EditorExportPlugin

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

/*
[EditorExportPlugin]s are automatically invoked whenever the user exports the project. Their most common use is to determine what files are being included in the exported project. For each plugin, [method _export_begin] is called at the beginning of the export process and then [method _export_file] is called for each exported file.
To use [EditorExportPlugin], register it using the [method EditorPlugin.add_export_plugin] method first.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=EditorExportPlugin)
*/
type Instance [1]gdclass.EditorExportPlugin

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorExportPlugin() Instance
}
type Interface interface {
	//Virtual method to be overridden by the user. Called for each exported file before [method _customize_resource] and [method _customize_scene]. The arguments can be used to identify the file. [param path] is the path of the file, [param type] is the [Resource] represented by the file (e.g. [PackedScene]), and [param features] is the list of features for the export.
	//Calling [method skip] inside this callback will make the file not included in the export.
	ExportFile(path string, atype string, features []string)
	//Virtual method to be overridden by the user. It is called when the export starts and provides all information about the export. [param features] is the list of features for the export, [param is_debug] is [code]true[/code] for debug builds, [param path] is the target path for the exported project. [param flags] is only used when running a runnable profile, e.g. when using native run on Android.
	ExportBegin(features []string, is_debug bool, path string, flags int)
	//Virtual method to be overridden by the user. Called when the export is finished.
	ExportEnd()
	//Return [code]true[/code] if this plugin will customize resources based on the platform and features used.
	//When enabled, [method _get_customization_configuration_hash] and [method _customize_resource] will be called and must be implemented.
	BeginCustomizeResources(platform [1]gdclass.EditorExportPlatform, features []string) bool
	//Customize a resource. If changes are made to it, return the same or a new resource. Otherwise, return [code]null[/code].
	//The [i]path[/i] argument is only used when customizing an actual file, otherwise this means that this resource is part of another one and it will be empty.
	//Implementing this method is required if [method _begin_customize_resources] returns [code]true[/code].
	CustomizeResource(resource [1]gdclass.Resource, path string) [1]gdclass.Resource
	//Return [code]true[/code] if this plugin will customize scenes based on the platform and features used.
	//When enabled, [method _get_customization_configuration_hash] and [method _customize_scene] will be called and must be implemented.
	BeginCustomizeScenes(platform [1]gdclass.EditorExportPlatform, features []string) bool
	//Customize a scene. If changes are made to it, return the same or a new scene. Otherwise, return [code]null[/code]. If a new scene is returned, it is up to you to dispose of the old one.
	//Implementing this method is required if [method _begin_customize_scenes] returns [code]true[/code].
	CustomizeScene(scene [1]gdclass.Node, path string) [1]gdclass.Node
	//Return a hash based on the configuration passed (for both scenes and resources). This helps keep separate caches for separate export configurations.
	//Implementing this method is required if [method _begin_customize_resources] returns [code]true[/code].
	GetCustomizationConfigurationHash() int
	//This is called when the customization process for scenes ends.
	EndCustomizeScenes()
	//This is called when the customization process for resources ends.
	EndCustomizeResources()
	//Return a list of export options that can be configured for this export plugin.
	//Each element in the return value is a [Dictionary] with the following keys:
	//- [code]option[/code]: A dictionary with the structure documented by [method Object.get_property_list], but all keys are optional.
	//- [code]default_value[/code]: The default value for this option.
	//- [code]update_visibility[/code]: An optional boolean value. If set to [code]true[/code], the preset will emit [signal Object.property_list_changed] when the option is changed.
	GetExportOptions(platform [1]gdclass.EditorExportPlatform) []map[any]any
	//Return a [Dictionary] of override values for export options, that will be used instead of user-provided values. Overridden options will be hidden from the user interface.
	//[codeblock]
	//class MyExportPlugin extends EditorExportPlugin:
	//    func _get_name() -> String:
	//        return "MyExportPlugin"
	//
	//    func _supports_platform(platform) -> bool:
	//        if platform is EditorExportPlatformPC:
	//            # Run on all desktop platforms including Windows, MacOS and Linux.
	//            return true
	//        return false
	//
	//    func _get_export_options_overrides(platform) -> Dictionary:
	//        # Override "Embed PCK" to always be enabled.
	//        return {
	//            "binary_format/embed_pck": true,
	//        }
	//[/codeblock]
	GetExportOptionsOverrides(platform [1]gdclass.EditorExportPlatform) map[any]any
	//Return [code]true[/code], if the result of [method _get_export_options] has changed and the export options of preset corresponding to [param platform] should be updated.
	ShouldUpdateExportOptions(platform [1]gdclass.EditorExportPlatform) bool
	//Check the requirements for the given [param option] and return a non-empty warning string if they are not met.
	//[b]Note:[/b] Use [method get_option] to check the value of the export options.
	GetExportOptionWarning(platform [1]gdclass.EditorExportPlatform, option string) string
	//Return a [PackedStringArray] of additional features this preset, for the given [param platform], should have.
	GetExportFeatures(platform [1]gdclass.EditorExportPlatform, debug bool) []string
	//Return the name identifier of this plugin (for future identification by the exporter). The plugins are sorted by name before exporting.
	//Implementing this method is required.
	GetName() string
	//Return [code]true[/code] if the plugin supports the given [param platform].
	SupportsPlatform(platform [1]gdclass.EditorExportPlatform) bool
	//Virtual method to be overridden by the user. This is called to retrieve the set of Android dependencies provided by this plugin. Each returned Android dependency should have the format of an Android remote binary dependency: [code]org.godot.example:my-plugin:0.0.0[/code]
	//For more information see [url=https://developer.android.com/build/dependencies?agpversion=4.1#dependency-types]Android documentation on dependencies[/url].
	//[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
	GetAndroidDependencies(platform [1]gdclass.EditorExportPlatform, debug bool) []string
	//Virtual method to be overridden by the user. This is called to retrieve the URLs of Maven repositories for the set of Android dependencies provided by this plugin.
	//For more information see [url=https://docs.gradle.org/current/userguide/dependency_management.html#sec:maven_repo]Gradle documentation on dependency management[/url].
	//[b]Note:[/b] Google's Maven repo and the Maven Central repo are already included by default.
	//[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
	GetAndroidDependenciesMavenRepos(platform [1]gdclass.EditorExportPlatform, debug bool) []string
	//Virtual method to be overridden by the user. This is called to retrieve the local paths of the Android libraries archive (AAR) files provided by this plugin.
	//[b]Note:[/b] Relative paths [b]must[/b] be relative to Godot's [code]res://addons/[/code] directory. For example, an AAR file located under [code]res://addons/hello_world_plugin/HelloWorld.release.aar[/code] can be returned as an absolute path using [code]res://addons/hello_world_plugin/HelloWorld.release.aar[/code] or a relative path using [code]hello_world_plugin/HelloWorld.release.aar[/code].
	//[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
	GetAndroidLibraries(platform [1]gdclass.EditorExportPlatform, debug bool) []string
	//Virtual method to be overridden by the user. This is used at export time to update the contents of the [code]activity[/code] element in the generated Android manifest.
	//[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
	GetAndroidManifestActivityElementContents(platform [1]gdclass.EditorExportPlatform, debug bool) string
	//Virtual method to be overridden by the user. This is used at export time to update the contents of the [code]application[/code] element in the generated Android manifest.
	//[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
	GetAndroidManifestApplicationElementContents(platform [1]gdclass.EditorExportPlatform, debug bool) string
	//Virtual method to be overridden by the user. This is used at export time to update the contents of the [code]manifest[/code] element in the generated Android manifest.
	//[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
	GetAndroidManifestElementContents(platform [1]gdclass.EditorExportPlatform, debug bool) string
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) ExportFile(path string, atype string, features []string) { return }
func (self implementation) ExportBegin(features []string, is_debug bool, path string, flags int) {
	return
}
func (self implementation) ExportEnd() { return }
func (self implementation) BeginCustomizeResources(platform [1]gdclass.EditorExportPlatform, features []string) (_ bool) {
	return
}
func (self implementation) CustomizeResource(resource [1]gdclass.Resource, path string) (_ [1]gdclass.Resource) {
	return
}
func (self implementation) BeginCustomizeScenes(platform [1]gdclass.EditorExportPlatform, features []string) (_ bool) {
	return
}
func (self implementation) CustomizeScene(scene [1]gdclass.Node, path string) (_ [1]gdclass.Node) {
	return
}
func (self implementation) GetCustomizationConfigurationHash() (_ int) { return }
func (self implementation) EndCustomizeScenes()                        { return }
func (self implementation) EndCustomizeResources()                     { return }
func (self implementation) GetExportOptions(platform [1]gdclass.EditorExportPlatform) (_ []map[any]any) {
	return
}
func (self implementation) GetExportOptionsOverrides(platform [1]gdclass.EditorExportPlatform) (_ map[any]any) {
	return
}
func (self implementation) ShouldUpdateExportOptions(platform [1]gdclass.EditorExportPlatform) (_ bool) {
	return
}
func (self implementation) GetExportOptionWarning(platform [1]gdclass.EditorExportPlatform, option string) (_ string) {
	return
}
func (self implementation) GetExportFeatures(platform [1]gdclass.EditorExportPlatform, debug bool) (_ []string) {
	return
}
func (self implementation) GetName() (_ string) { return }
func (self implementation) SupportsPlatform(platform [1]gdclass.EditorExportPlatform) (_ bool) {
	return
}
func (self implementation) GetAndroidDependencies(platform [1]gdclass.EditorExportPlatform, debug bool) (_ []string) {
	return
}
func (self implementation) GetAndroidDependenciesMavenRepos(platform [1]gdclass.EditorExportPlatform, debug bool) (_ []string) {
	return
}
func (self implementation) GetAndroidLibraries(platform [1]gdclass.EditorExportPlatform, debug bool) (_ []string) {
	return
}
func (self implementation) GetAndroidManifestActivityElementContents(platform [1]gdclass.EditorExportPlatform, debug bool) (_ string) {
	return
}
func (self implementation) GetAndroidManifestApplicationElementContents(platform [1]gdclass.EditorExportPlatform, debug bool) (_ string) {
	return
}
func (self implementation) GetAndroidManifestElementContents(platform [1]gdclass.EditorExportPlatform, debug bool) (_ string) {
	return
}

/*
Virtual method to be overridden by the user. Called for each exported file before [method _customize_resource] and [method _customize_scene]. The arguments can be used to identify the file. [param path] is the path of the file, [param type] is the [Resource] represented by the file (e.g. [PackedScene]), and [param features] is the list of features for the export.
Calling [method skip] inside this callback will make the file not included in the export.
*/
func (Instance) _export_file(impl func(ptr unsafe.Pointer, path string, atype string, features []string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(path)
		var atype = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		defer pointers.End(atype)
		var features = pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 2))
		defer pointers.End(features)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, path.String(), atype.String(), features.Strings())
	}
}

/*
Virtual method to be overridden by the user. It is called when the export starts and provides all information about the export. [param features] is the list of features for the export, [param is_debug] is [code]true[/code] for debug builds, [param path] is the target path for the exported project. [param flags] is only used when running a runnable profile, e.g. when using native run on Android.
*/
func (Instance) _export_begin(impl func(ptr unsafe.Pointer, features []string, is_debug bool, path string, flags int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var features = pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 0))
		defer pointers.End(features)
		var is_debug = gd.UnsafeGet[bool](p_args, 1)

		var path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		defer pointers.End(path)
		var flags = gd.UnsafeGet[gd.Int](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, features.Strings(), is_debug, path.String(), int(flags))
	}
}

/*
Virtual method to be overridden by the user. Called when the export is finished.
*/
func (Instance) _export_end(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Return [code]true[/code] if this plugin will customize resources based on the platform and features used.
When enabled, [method _get_customization_configuration_hash] and [method _customize_resource] will be called and must be implemented.
*/
func (Instance) _begin_customize_resources(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, features []string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var features = pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 1))
		defer pointers.End(features)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, features.Strings())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Customize a resource. If changes are made to it, return the same or a new resource. Otherwise, return [code]null[/code].
The [i]path[/i] argument is only used when customizing an actual file, otherwise this means that this resource is part of another one and it will be empty.
Implementing this method is required if [method _begin_customize_resources] returns [code]true[/code].
*/
func (Instance) _customize_resource(impl func(ptr unsafe.Pointer, resource [1]gdclass.Resource, path string) [1]gdclass.Resource) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var resource = [1]gdclass.Resource{pointers.New[gdclass.Resource]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(resource[0])
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		defer pointers.End(path)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path.String())
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Return [code]true[/code] if this plugin will customize scenes based on the platform and features used.
When enabled, [method _get_customization_configuration_hash] and [method _customize_scene] will be called and must be implemented.
*/
func (Instance) _begin_customize_scenes(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, features []string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var features = pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 1))
		defer pointers.End(features)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, features.Strings())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Customize a scene. If changes are made to it, return the same or a new scene. Otherwise, return [code]null[/code]. If a new scene is returned, it is up to you to dispose of the old one.
Implementing this method is required if [method _begin_customize_scenes] returns [code]true[/code].
*/
func (Instance) _customize_scene(impl func(ptr unsafe.Pointer, scene [1]gdclass.Node, path string) [1]gdclass.Node) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var scene = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(scene[0])
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		defer pointers.End(path)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, scene, path.String())
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Return a hash based on the configuration passed (for both scenes and resources). This helps keep separate caches for separate export configurations.
Implementing this method is required if [method _begin_customize_resources] returns [code]true[/code].
*/
func (Instance) _get_customization_configuration_hash(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
This is called when the customization process for scenes ends.
*/
func (Instance) _end_customize_scenes(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
This is called when the customization process for resources ends.
*/
func (Instance) _end_customize_resources(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Return a list of export options that can be configured for this export plugin.
Each element in the return value is a [Dictionary] with the following keys:
- [code]option[/code]: A dictionary with the structure documented by [method Object.get_property_list], but all keys are optional.
- [code]default_value[/code]: The default value for this option.
- [code]update_visibility[/code]: An optional boolean value. If set to [code]true[/code], the preset will emit [signal Object.property_list_changed] when the option is changed.
*/
func (Instance) _get_export_options(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform) []map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform)
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[Dictionary.Any]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Return a [Dictionary] of override values for export options, that will be used instead of user-provided values. Overridden options will be hidden from the user interface.
[codeblock]
class MyExportPlugin extends EditorExportPlugin:

	func _get_name() -> String:
	    return "MyExportPlugin"

	func _supports_platform(platform) -> bool:
	    if platform is EditorExportPlatformPC:
	        # Run on all desktop platforms including Windows, MacOS and Linux.
	        return true
	    return false

	func _get_export_options_overrides(platform) -> Dictionary:
	    # Override "Embed PCK" to always be enabled.
	    return {
	        "binary_format/embed_pck": true,
	    }

[/codeblock]
*/
func (Instance) _get_export_options_overrides(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform)
		gd.UnsafeSet(p_back, gd.DictionaryFromMap(ret))
	}
}

/*
Return [code]true[/code], if the result of [method _get_export_options] has changed and the export options of preset corresponding to [param platform] should be updated.
*/
func (Instance) _should_update_export_options(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Check the requirements for the given [param option] and return a non-empty warning string if they are not met.
[b]Note:[/b] Use [method get_option] to check the value of the export options.
*/
func (Instance) _get_export_option_warning(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, option string) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var option = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		defer pointers.End(option)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, option.String())
		ptr, ok := pointers.End(gd.NewString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Return a [PackedStringArray] of additional features this preset, for the given [param platform], should have.
*/
func (Instance) _get_export_features(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, debug bool) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Return the name identifier of this plugin (for future identification by the exporter). The plugins are sorted by name before exporting.
Implementing this method is required.
*/
func (Instance) _get_name(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
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
Return [code]true[/code] if the plugin supports the given [param platform].
*/
func (Instance) _supports_platform(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method to be overridden by the user. This is called to retrieve the set of Android dependencies provided by this plugin. Each returned Android dependency should have the format of an Android remote binary dependency: [code]org.godot.example:my-plugin:0.0.0[/code]
For more information see [url=https://developer.android.com/build/dependencies?agpversion=4.1#dependency-types]Android documentation on dependencies[/url].
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (Instance) _get_android_dependencies(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, debug bool) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to be overridden by the user. This is called to retrieve the URLs of Maven repositories for the set of Android dependencies provided by this plugin.
For more information see [url=https://docs.gradle.org/current/userguide/dependency_management.html#sec:maven_repo]Gradle documentation on dependency management[/url].
[b]Note:[/b] Google's Maven repo and the Maven Central repo are already included by default.
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (Instance) _get_android_dependencies_maven_repos(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, debug bool) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to be overridden by the user. This is called to retrieve the local paths of the Android libraries archive (AAR) files provided by this plugin.
[b]Note:[/b] Relative paths [b]must[/b] be relative to Godot's [code]res://addons/[/code] directory. For example, an AAR file located under [code]res://addons/hello_world_plugin/HelloWorld.release.aar[/code] can be returned as an absolute path using [code]res://addons/hello_world_plugin/HelloWorld.release.aar[/code] or a relative path using [code]hello_world_plugin/HelloWorld.release.aar[/code].
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (Instance) _get_android_libraries(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, debug bool) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to be overridden by the user. This is used at export time to update the contents of the [code]activity[/code] element in the generated Android manifest.
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (Instance) _get_android_manifest_activity_element_contents(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, debug bool) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		ptr, ok := pointers.End(gd.NewString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to be overridden by the user. This is used at export time to update the contents of the [code]application[/code] element in the generated Android manifest.
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (Instance) _get_android_manifest_application_element_contents(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, debug bool) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		ptr, ok := pointers.End(gd.NewString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to be overridden by the user. This is used at export time to update the contents of the [code]manifest[/code] element in the generated Android manifest.
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (Instance) _get_android_manifest_element_contents(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, debug bool) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		ptr, ok := pointers.End(gd.NewString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Adds a shared object or a directory containing only shared objects with the given [param tags] and destination [param path].
[b]Note:[/b] In case of macOS exports, those shared objects will be added to [code]Frameworks[/code] directory of app bundle.
In case of a directory code-sign will error if you place non code object in directory.
*/
func (self Instance) AddSharedObject(path string, tags []string, target string) { //gd:EditorExportPlugin.add_shared_object
	class(self).AddSharedObject(gd.NewString(path), gd.NewPackedStringSlice(tags), gd.NewString(target))
}

/*
Adds a static lib from the given [param path] to the iOS project.
*/
func (self Instance) AddIosProjectStaticLib(path string) { //gd:EditorExportPlugin.add_ios_project_static_lib
	class(self).AddIosProjectStaticLib(gd.NewString(path))
}

/*
Adds a custom file to be exported. [param path] is the virtual path that can be used to load the file, [param file] is the binary data of the file.
When called inside [method _export_file] and [param remap] is [code]true[/code], the current file will not be exported, but instead remapped to this custom file. [param remap] is ignored when called in other places.
[param file] will not be imported, so consider using [method _customize_resource] to remap imported resources.
*/
func (self Instance) AddFile(path string, file []byte, remap bool) { //gd:EditorExportPlugin.add_file
	class(self).AddFile(gd.NewString(path), gd.NewPackedByteSlice(file), remap)
}

/*
Adds a static library (*.a) or dynamic library (*.dylib, *.framework) to Linking Phase in iOS's Xcode project.
*/
func (self Instance) AddIosFramework(path string) { //gd:EditorExportPlugin.add_ios_framework
	class(self).AddIosFramework(gd.NewString(path))
}

/*
Adds a dynamic library (*.dylib, *.framework) to Linking Phase in iOS's Xcode project and embeds it into resulting binary.
[b]Note:[/b] For static libraries (*.a) works in same way as [method add_ios_framework].
[b]Note:[/b] This method should not be used for System libraries as they are already present on the device.
*/
func (self Instance) AddIosEmbeddedFramework(path string) { //gd:EditorExportPlugin.add_ios_embedded_framework
	class(self).AddIosEmbeddedFramework(gd.NewString(path))
}

/*
Adds content for iOS Property List files.
*/
func (self Instance) AddIosPlistContent(plist_content string) { //gd:EditorExportPlugin.add_ios_plist_content
	class(self).AddIosPlistContent(gd.NewString(plist_content))
}

/*
Adds linker flags for the iOS export.
*/
func (self Instance) AddIosLinkerFlags(flags string) { //gd:EditorExportPlugin.add_ios_linker_flags
	class(self).AddIosLinkerFlags(gd.NewString(flags))
}

/*
Adds an iOS bundle file from the given [param path] to the exported project.
*/
func (self Instance) AddIosBundleFile(path string) { //gd:EditorExportPlugin.add_ios_bundle_file
	class(self).AddIosBundleFile(gd.NewString(path))
}

/*
Adds a C++ code to the iOS export. The final code is created from the code appended by each active export plugin.
*/
func (self Instance) AddIosCppCode(code string) { //gd:EditorExportPlugin.add_ios_cpp_code
	class(self).AddIosCppCode(gd.NewString(code))
}

/*
Adds file or directory matching [param path] to [code]PlugIns[/code] directory of macOS app bundle.
[b]Note:[/b] This is useful only for macOS exports.
*/
func (self Instance) AddMacosPluginFile(path string) { //gd:EditorExportPlugin.add_macos_plugin_file
	class(self).AddMacosPluginFile(gd.NewString(path))
}

/*
To be called inside [method _export_file]. Skips the current file, so it's not included in the export.
*/
func (self Instance) Skip() { //gd:EditorExportPlugin.skip
	class(self).Skip()
}

/*
Returns the current value of an export option supplied by [method _get_export_options].
*/
func (self Instance) GetOption(name string) any { //gd:EditorExportPlugin.get_option
	return any(class(self).GetOption(gd.NewStringName(name)).Interface())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorExportPlugin

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorExportPlugin"))
	casted := Instance{*(*gdclass.EditorExportPlugin)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Virtual method to be overridden by the user. Called for each exported file before [method _customize_resource] and [method _customize_scene]. The arguments can be used to identify the file. [param path] is the path of the file, [param type] is the [Resource] represented by the file (e.g. [PackedScene]), and [param features] is the list of features for the export.
Calling [method skip] inside this callback will make the file not included in the export.
*/
func (class) _export_file(impl func(ptr unsafe.Pointer, path gd.String, atype gd.String, features gd.PackedStringArray)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(path)
		var atype = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		defer pointers.End(atype)
		var features = pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 2))
		defer pointers.End(features)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, path, atype, features)
	}
}

/*
Virtual method to be overridden by the user. It is called when the export starts and provides all information about the export. [param features] is the list of features for the export, [param is_debug] is [code]true[/code] for debug builds, [param path] is the target path for the exported project. [param flags] is only used when running a runnable profile, e.g. when using native run on Android.
*/
func (class) _export_begin(impl func(ptr unsafe.Pointer, features gd.PackedStringArray, is_debug bool, path gd.String, flags gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var features = pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 0))
		defer pointers.End(features)
		var is_debug = gd.UnsafeGet[bool](p_args, 1)

		var path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		defer pointers.End(path)
		var flags = gd.UnsafeGet[gd.Int](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, features, is_debug, path, flags)
	}
}

/*
Virtual method to be overridden by the user. Called when the export is finished.
*/
func (class) _export_end(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Return [code]true[/code] if this plugin will customize resources based on the platform and features used.
When enabled, [method _get_customization_configuration_hash] and [method _customize_resource] will be called and must be implemented.
*/
func (class) _begin_customize_resources(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, features gd.PackedStringArray) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var features = pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 1))
		defer pointers.End(features)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, features)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Customize a resource. If changes are made to it, return the same or a new resource. Otherwise, return [code]null[/code].
The [i]path[/i] argument is only used when customizing an actual file, otherwise this means that this resource is part of another one and it will be empty.
Implementing this method is required if [method _begin_customize_resources] returns [code]true[/code].
*/
func (class) _customize_resource(impl func(ptr unsafe.Pointer, resource [1]gdclass.Resource, path gd.String) [1]gdclass.Resource) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var resource = [1]gdclass.Resource{pointers.New[gdclass.Resource]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(resource[0])
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		defer pointers.End(path)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Return [code]true[/code] if this plugin will customize scenes based on the platform and features used.
When enabled, [method _get_customization_configuration_hash] and [method _customize_scene] will be called and must be implemented.
*/
func (class) _begin_customize_scenes(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, features gd.PackedStringArray) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var features = pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 1))
		defer pointers.End(features)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, features)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Customize a scene. If changes are made to it, return the same or a new scene. Otherwise, return [code]null[/code]. If a new scene is returned, it is up to you to dispose of the old one.
Implementing this method is required if [method _begin_customize_scenes] returns [code]true[/code].
*/
func (class) _customize_scene(impl func(ptr unsafe.Pointer, scene [1]gdclass.Node, path gd.String) [1]gdclass.Node) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var scene = [1]gdclass.Node{pointers.New[gdclass.Node]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(scene[0])
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		defer pointers.End(path)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, scene, path)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Return a hash based on the configuration passed (for both scenes and resources). This helps keep separate caches for separate export configurations.
Implementing this method is required if [method _begin_customize_resources] returns [code]true[/code].
*/
func (class) _get_customization_configuration_hash(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
This is called when the customization process for scenes ends.
*/
func (class) _end_customize_scenes(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
This is called when the customization process for resources ends.
*/
func (class) _end_customize_resources(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Return a list of export options that can be configured for this export plugin.
Each element in the return value is a [Dictionary] with the following keys:
- [code]option[/code]: A dictionary with the structure documented by [method Object.get_property_list], but all keys are optional.
- [code]default_value[/code]: The default value for this option.
- [code]update_visibility[/code]: An optional boolean value. If set to [code]true[/code], the preset will emit [signal Object.property_list_changed] when the option is changed.
*/
func (class) _get_export_options(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform) Array.Contains[Dictionary.Any]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Return a [Dictionary] of override values for export options, that will be used instead of user-provided values. Overridden options will be hidden from the user interface.
[codeblock]
class MyExportPlugin extends EditorExportPlugin:

	func _get_name() -> String:
	    return "MyExportPlugin"

	func _supports_platform(platform) -> bool:
	    if platform is EditorExportPlatformPC:
	        # Run on all desktop platforms including Windows, MacOS and Linux.
	        return true
	    return false

	func _get_export_options_overrides(platform) -> Dictionary:
	    # Override "Embed PCK" to always be enabled.
	    return {
	        "binary_format/embed_pck": true,
	    }

[/codeblock]
*/
func (class) _get_export_options_overrides(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Return [code]true[/code], if the result of [method _get_export_options] has changed and the export options of preset corresponding to [param platform] should be updated.
*/
func (class) _should_update_export_options(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Check the requirements for the given [param option] and return a non-empty warning string if they are not met.
[b]Note:[/b] Use [method get_option] to check the value of the export options.
*/
func (class) _get_export_option_warning(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, option gd.String) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var option = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		defer pointers.End(option)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, option)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Return a [PackedStringArray] of additional features this preset, for the given [param platform], should have.
*/
func (class) _get_export_features(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, debug bool) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Return the name identifier of this plugin (for future identification by the exporter). The plugins are sorted by name before exporting.
Implementing this method is required.
*/
func (class) _get_name(impl func(ptr unsafe.Pointer) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
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
Return [code]true[/code] if the plugin supports the given [param platform].
*/
func (class) _supports_platform(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method to be overridden by the user. This is called to retrieve the set of Android dependencies provided by this plugin. Each returned Android dependency should have the format of an Android remote binary dependency: [code]org.godot.example:my-plugin:0.0.0[/code]
For more information see [url=https://developer.android.com/build/dependencies?agpversion=4.1#dependency-types]Android documentation on dependencies[/url].
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (class) _get_android_dependencies(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, debug bool) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to be overridden by the user. This is called to retrieve the URLs of Maven repositories for the set of Android dependencies provided by this plugin.
For more information see [url=https://docs.gradle.org/current/userguide/dependency_management.html#sec:maven_repo]Gradle documentation on dependency management[/url].
[b]Note:[/b] Google's Maven repo and the Maven Central repo are already included by default.
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (class) _get_android_dependencies_maven_repos(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, debug bool) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to be overridden by the user. This is called to retrieve the local paths of the Android libraries archive (AAR) files provided by this plugin.
[b]Note:[/b] Relative paths [b]must[/b] be relative to Godot's [code]res://addons/[/code] directory. For example, an AAR file located under [code]res://addons/hello_world_plugin/HelloWorld.release.aar[/code] can be returned as an absolute path using [code]res://addons/hello_world_plugin/HelloWorld.release.aar[/code] or a relative path using [code]hello_world_plugin/HelloWorld.release.aar[/code].
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (class) _get_android_libraries(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, debug bool) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to be overridden by the user. This is used at export time to update the contents of the [code]activity[/code] element in the generated Android manifest.
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (class) _get_android_manifest_activity_element_contents(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, debug bool) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to be overridden by the user. This is used at export time to update the contents of the [code]application[/code] element in the generated Android manifest.
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (class) _get_android_manifest_application_element_contents(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, debug bool) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to be overridden by the user. This is used at export time to update the contents of the [code]manifest[/code] element in the generated Android manifest.
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (class) _get_android_manifest_element_contents(impl func(ptr unsafe.Pointer, platform [1]gdclass.EditorExportPlatform, debug bool) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var platform = [1]gdclass.EditorExportPlatform{pointers.New[gdclass.EditorExportPlatform]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(platform[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Adds a shared object or a directory containing only shared objects with the given [param tags] and destination [param path].
[b]Note:[/b] In case of macOS exports, those shared objects will be added to [code]Frameworks[/code] directory of app bundle.
In case of a directory code-sign will error if you place non code object in directory.
*/
//go:nosplit
func (self class) AddSharedObject(path gd.String, tags gd.PackedStringArray, target gd.String) { //gd:EditorExportPlugin.add_shared_object
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, pointers.Get(tags))
	callframe.Arg(frame, pointers.Get(target))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorExportPlugin.Bind_add_shared_object, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a static lib from the given [param path] to the iOS project.
*/
//go:nosplit
func (self class) AddIosProjectStaticLib(path gd.String) { //gd:EditorExportPlugin.add_ios_project_static_lib
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorExportPlugin.Bind_add_ios_project_static_lib, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a custom file to be exported. [param path] is the virtual path that can be used to load the file, [param file] is the binary data of the file.
When called inside [method _export_file] and [param remap] is [code]true[/code], the current file will not be exported, but instead remapped to this custom file. [param remap] is ignored when called in other places.
[param file] will not be imported, so consider using [method _customize_resource] to remap imported resources.
*/
//go:nosplit
func (self class) AddFile(path gd.String, file gd.PackedByteArray, remap bool) { //gd:EditorExportPlugin.add_file
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, pointers.Get(file))
	callframe.Arg(frame, remap)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorExportPlugin.Bind_add_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a static library (*.a) or dynamic library (*.dylib, *.framework) to Linking Phase in iOS's Xcode project.
*/
//go:nosplit
func (self class) AddIosFramework(path gd.String) { //gd:EditorExportPlugin.add_ios_framework
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorExportPlugin.Bind_add_ios_framework, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a dynamic library (*.dylib, *.framework) to Linking Phase in iOS's Xcode project and embeds it into resulting binary.
[b]Note:[/b] For static libraries (*.a) works in same way as [method add_ios_framework].
[b]Note:[/b] This method should not be used for System libraries as they are already present on the device.
*/
//go:nosplit
func (self class) AddIosEmbeddedFramework(path gd.String) { //gd:EditorExportPlugin.add_ios_embedded_framework
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorExportPlugin.Bind_add_ios_embedded_framework, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds content for iOS Property List files.
*/
//go:nosplit
func (self class) AddIosPlistContent(plist_content gd.String) { //gd:EditorExportPlugin.add_ios_plist_content
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plist_content))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorExportPlugin.Bind_add_ios_plist_content, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds linker flags for the iOS export.
*/
//go:nosplit
func (self class) AddIosLinkerFlags(flags gd.String) { //gd:EditorExportPlugin.add_ios_linker_flags
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(flags))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorExportPlugin.Bind_add_ios_linker_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds an iOS bundle file from the given [param path] to the exported project.
*/
//go:nosplit
func (self class) AddIosBundleFile(path gd.String) { //gd:EditorExportPlugin.add_ios_bundle_file
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorExportPlugin.Bind_add_ios_bundle_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a C++ code to the iOS export. The final code is created from the code appended by each active export plugin.
*/
//go:nosplit
func (self class) AddIosCppCode(code gd.String) { //gd:EditorExportPlugin.add_ios_cpp_code
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(code))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorExportPlugin.Bind_add_ios_cpp_code, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds file or directory matching [param path] to [code]PlugIns[/code] directory of macOS app bundle.
[b]Note:[/b] This is useful only for macOS exports.
*/
//go:nosplit
func (self class) AddMacosPluginFile(path gd.String) { //gd:EditorExportPlugin.add_macos_plugin_file
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorExportPlugin.Bind_add_macos_plugin_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
To be called inside [method _export_file]. Skips the current file, so it's not included in the export.
*/
//go:nosplit
func (self class) Skip() { //gd:EditorExportPlugin.skip
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorExportPlugin.Bind_skip, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the current value of an export option supplied by [method _get_export_options].
*/
//go:nosplit
func (self class) GetOption(name gd.StringName) gd.Variant { //gd:EditorExportPlugin.get_option
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorExportPlugin.Bind_get_option, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsEditorExportPlugin() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorExportPlugin() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_export_file":
		return reflect.ValueOf(self._export_file)
	case "_export_begin":
		return reflect.ValueOf(self._export_begin)
	case "_export_end":
		return reflect.ValueOf(self._export_end)
	case "_begin_customize_resources":
		return reflect.ValueOf(self._begin_customize_resources)
	case "_customize_resource":
		return reflect.ValueOf(self._customize_resource)
	case "_begin_customize_scenes":
		return reflect.ValueOf(self._begin_customize_scenes)
	case "_customize_scene":
		return reflect.ValueOf(self._customize_scene)
	case "_get_customization_configuration_hash":
		return reflect.ValueOf(self._get_customization_configuration_hash)
	case "_end_customize_scenes":
		return reflect.ValueOf(self._end_customize_scenes)
	case "_end_customize_resources":
		return reflect.ValueOf(self._end_customize_resources)
	case "_get_export_options":
		return reflect.ValueOf(self._get_export_options)
	case "_get_export_options_overrides":
		return reflect.ValueOf(self._get_export_options_overrides)
	case "_should_update_export_options":
		return reflect.ValueOf(self._should_update_export_options)
	case "_get_export_option_warning":
		return reflect.ValueOf(self._get_export_option_warning)
	case "_get_export_features":
		return reflect.ValueOf(self._get_export_features)
	case "_get_name":
		return reflect.ValueOf(self._get_name)
	case "_supports_platform":
		return reflect.ValueOf(self._supports_platform)
	case "_get_android_dependencies":
		return reflect.ValueOf(self._get_android_dependencies)
	case "_get_android_dependencies_maven_repos":
		return reflect.ValueOf(self._get_android_dependencies_maven_repos)
	case "_get_android_libraries":
		return reflect.ValueOf(self._get_android_libraries)
	case "_get_android_manifest_activity_element_contents":
		return reflect.ValueOf(self._get_android_manifest_activity_element_contents)
	case "_get_android_manifest_application_element_contents":
		return reflect.ValueOf(self._get_android_manifest_application_element_contents)
	case "_get_android_manifest_element_contents":
		return reflect.ValueOf(self._get_android_manifest_element_contents)
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_export_file":
		return reflect.ValueOf(self._export_file)
	case "_export_begin":
		return reflect.ValueOf(self._export_begin)
	case "_export_end":
		return reflect.ValueOf(self._export_end)
	case "_begin_customize_resources":
		return reflect.ValueOf(self._begin_customize_resources)
	case "_customize_resource":
		return reflect.ValueOf(self._customize_resource)
	case "_begin_customize_scenes":
		return reflect.ValueOf(self._begin_customize_scenes)
	case "_customize_scene":
		return reflect.ValueOf(self._customize_scene)
	case "_get_customization_configuration_hash":
		return reflect.ValueOf(self._get_customization_configuration_hash)
	case "_end_customize_scenes":
		return reflect.ValueOf(self._end_customize_scenes)
	case "_end_customize_resources":
		return reflect.ValueOf(self._end_customize_resources)
	case "_get_export_options":
		return reflect.ValueOf(self._get_export_options)
	case "_get_export_options_overrides":
		return reflect.ValueOf(self._get_export_options_overrides)
	case "_should_update_export_options":
		return reflect.ValueOf(self._should_update_export_options)
	case "_get_export_option_warning":
		return reflect.ValueOf(self._get_export_option_warning)
	case "_get_export_features":
		return reflect.ValueOf(self._get_export_features)
	case "_get_name":
		return reflect.ValueOf(self._get_name)
	case "_supports_platform":
		return reflect.ValueOf(self._supports_platform)
	case "_get_android_dependencies":
		return reflect.ValueOf(self._get_android_dependencies)
	case "_get_android_dependencies_maven_repos":
		return reflect.ValueOf(self._get_android_dependencies_maven_repos)
	case "_get_android_libraries":
		return reflect.ValueOf(self._get_android_libraries)
	case "_get_android_manifest_activity_element_contents":
		return reflect.ValueOf(self._get_android_manifest_activity_element_contents)
	case "_get_android_manifest_application_element_contents":
		return reflect.ValueOf(self._get_android_manifest_application_element_contents)
	case "_get_android_manifest_element_contents":
		return reflect.ValueOf(self._get_android_manifest_element_contents)
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("EditorExportPlugin", func(ptr gd.Object) any {
		return [1]gdclass.EditorExportPlugin{*(*gdclass.EditorExportPlugin)(unsafe.Pointer(&ptr))}
	})
}
