// Package EditorExportPlatformExtension provides methods for working with EditorExportPlatformExtension object instances.
package EditorExportPlatformExtension

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/EditorExportPlatform"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

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
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
External [EditorExportPlatform] implementations should inherit from this class.
To use [EditorExportPlatform], register it using the [method EditorPlugin.add_export_platform] method first.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=EditorExportPlatformExtension)
*/
type Instance [1]gdclass.EditorExportPlatformExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorExportPlatformExtension() Instance
}
type Interface interface {
	//[b]Required.[/b]
	//Returns array of platform specific features for the specified [param preset].
	GetPresetFeatures(preset [1]gdclass.EditorExportPreset) []string
	//[b]Optional.[/b]
	//Returns [code]true[/code] if specified file is a valid executable (native executable or script) for the target platform.
	IsExecutable(path string) bool
	//[b]Optional.[/b]
	//Returns a property list, as an [Array] of dictionaries. Each [Dictionary] must at least contain the [code]name: StringName[/code] and [code]type: Variant.Type[/code] entries.
	//Additionally, the following keys are supported:
	//- [code]hint: PropertyHint[/code]
	//- [code]hint_string: String[/code]
	//- [code]usage: PropertyUsageFlags[/code]
	//- [code]class_name: StringName[/code]
	//- [code]default_value: Variant[/code], default value of the property.
	//- [code]update_visibility: bool[/code], if set to [code]true[/code], [method _get_export_option_visibility] is called for each property when this property is changed.
	//- [code]required: bool[/code], if set to [code]true[/code], this property warnings are critical, and should be resolved to make export possible. This value is a hint for the [method _has_valid_export_configuration] implementation, and not used by the engine directly.
	//See also [method Object._get_property_list].
	GetExportOptions() []map[any]any
	//[b]Optional.[/b]
	//Returns [code]true[/code] if export options list is changed and presets should be updated.
	ShouldUpdateExportOptions() bool
	//[b]Optional.[/b]
	//Validates [param option] and returns visibility for the specified [param preset]. Default implementation return [code]true[/code] for all options.
	GetExportOptionVisibility(preset [1]gdclass.EditorExportPreset, option string) bool
	//[b]Optional.[/b]
	//Validates [param option] and returns warning message for the specified [param preset]. Default implementation return empty string for all options.
	GetExportOptionWarning(preset [1]gdclass.EditorExportPreset, option string) string
	//[b]Required.[/b]
	//Returns target OS name.
	GetOsName() string
	//[b]Required.[/b]
	//Returns export platform name.
	GetName() string
	//[b]Required.[/b]
	//Returns platform logo displayed in the export dialog, logo should be 32x32 adjusted to the current editor scale, see [method EditorInterface.get_editor_scale].
	GetLogo() [1]gdclass.Texture2D
	//[b]Optional.[/b]
	//Returns [code]true[/code] if one-click deploy options are changed and editor interface should be updated.
	PollExport() bool
	//[b]Optional.[/b]
	//Returns number one-click deploy devices (or other one-click option displayed in the menu).
	GetOptionsCount() int
	//[b]Optional.[/b]
	//Returns tooltip of the one-click deploy menu button.
	GetOptionsTooltip() string
	//[b]Optional.[/b]
	//Returns one-click deploy menu item icon for the specified [param device], icon should be 16x16 adjusted to the current editor scale, see [method EditorInterface.get_editor_scale].
	GetOptionIcon(device int) [1]gdclass.ImageTexture
	//[b]Optional.[/b]
	//Returns one-click deploy menu item label for the specified [param device].
	GetOptionLabel(device int) string
	//[b]Optional.[/b]
	//Returns one-click deploy menu item tooltip for the specified [param device].
	GetOptionTooltip(device int) string
	//[b]Optional.[/b]
	//Returns device architecture for one-click deploy.
	GetDeviceArchitecture(device int) string
	//[b]Optional.[/b]
	//Called by the editor before platform is unregistered.
	Cleanup()
	//[b]Optional.[/b]
	//This method is called when [param device] one-click deploy menu option is selected.
	//Implementation should export project to a temporary location, upload and run it on the specific [param device], or perform another action associated with the menu item.
	Run(preset [1]gdclass.EditorExportPreset, device int, debug_flags gdclass.EditorExportPlatformDebugFlags) error
	//[b]Optional.[/b]
	//Returns icon of the one-click deploy menu button, icon should be 16x16 adjusted to the current editor scale, see [method EditorInterface.get_editor_scale].
	GetRunIcon() [1]gdclass.Texture2D
	//[b]Optional.[/b]
	//Returns [code]true[/code], if specified [param preset] is valid and can be exported. Use [method set_config_error] and [method set_config_missing_templates] to set error details.
	//Usual implementation can call [method _has_valid_export_configuration] and [method _has_valid_project_configuration] to determine if export is possible.
	CanExport(preset [1]gdclass.EditorExportPreset, debug bool) bool
	//[b]Required.[/b]
	//Returns [code]true[/code] if export configuration is valid.
	HasValidExportConfiguration(preset [1]gdclass.EditorExportPreset, debug bool) bool
	//[b]Required.[/b]
	//Returns [code]true[/code] if project configuration is valid.
	HasValidProjectConfiguration(preset [1]gdclass.EditorExportPreset) bool
	//[b]Required.[/b]
	//Returns array of supported binary extensions for the full project export.
	GetBinaryExtensions(preset [1]gdclass.EditorExportPreset) []string
	//[b]Required.[/b]
	//Creates a full project at [param path] for the specified [param preset].
	//This method is called when "Export" button is pressed in the export dialog.
	//This method implementation can call [method EditorExportPlatform.save_pack] or [method EditorExportPlatform.save_zip] to use default PCK/ZIP export process, or calls [method EditorExportPlatform.export_project_files] and implement custom callback for processing each exported file.
	ExportProject(preset [1]gdclass.EditorExportPreset, debug bool, path string, flags gdclass.EditorExportPlatformDebugFlags) error
	//[b]Optional.[/b]
	//Creates a PCK archive at [param path] for the specified [param preset].
	//This method is called when "Export PCK/ZIP" button is pressed in the export dialog, with "Export as Patch" disabled, and PCK is selected as a file type.
	ExportPack(preset [1]gdclass.EditorExportPreset, debug bool, path string, flags gdclass.EditorExportPlatformDebugFlags) error
	//[b]Optional.[/b]
	//Create a ZIP archive at [param path] for the specified [param preset].
	//This method is called when "Export PCK/ZIP" button is pressed in the export dialog, with "Export as Patch" disabled, and ZIP is selected as a file type.
	ExportZip(preset [1]gdclass.EditorExportPreset, debug bool, path string, flags gdclass.EditorExportPlatformDebugFlags) error
	//[b]Optional.[/b]
	//Creates a patch PCK archive at [param path] for the specified [param preset], containing only the files that have changed since the last patch.
	//This method is called when "Export PCK/ZIP" button is pressed in the export dialog, with "Export as Patch" enabled, and PCK is selected as a file type.
	//[b]Note:[/b] The patches provided in [param patches] have already been loaded when this method is called and are merely provided as context. When empty the patches defined in the export preset have been loaded instead.
	ExportPackPatch(preset [1]gdclass.EditorExportPreset, debug bool, path string, patches []string, flags gdclass.EditorExportPlatformDebugFlags) error
	//[b]Optional.[/b]
	//Create a ZIP archive at [param path] for the specified [param preset], containing only the files that have changed since the last patch.
	//This method is called when "Export PCK/ZIP" button is pressed in the export dialog, with "Export as Patch" enabled, and ZIP is selected as a file type.
	//[b]Note:[/b] The patches provided in [param patches] have already been loaded when this method is called and are merely provided as context. When empty the patches defined in the export preset have been loaded instead.
	ExportZipPatch(preset [1]gdclass.EditorExportPreset, debug bool, path string, patches []string, flags gdclass.EditorExportPlatformDebugFlags) error
	//[b]Required.[/b]
	//Returns array of platform specific features.
	GetPlatformFeatures() []string
	//[b]Optional.[/b]
	//Returns protocol used for remote debugging. Default implementation return [code]tcp://[/code].
	GetDebugProtocol() string
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetPresetFeatures(preset [1]gdclass.EditorExportPreset) (_ []string) {
	return
}
func (self implementation) IsExecutable(path string) (_ bool)   { return }
func (self implementation) GetExportOptions() (_ []map[any]any) { return }
func (self implementation) ShouldUpdateExportOptions() (_ bool) { return }
func (self implementation) GetExportOptionVisibility(preset [1]gdclass.EditorExportPreset, option string) (_ bool) {
	return
}
func (self implementation) GetExportOptionWarning(preset [1]gdclass.EditorExportPreset, option string) (_ string) {
	return
}
func (self implementation) GetOsName() (_ string)                                { return }
func (self implementation) GetName() (_ string)                                  { return }
func (self implementation) GetLogo() (_ [1]gdclass.Texture2D)                    { return }
func (self implementation) PollExport() (_ bool)                                 { return }
func (self implementation) GetOptionsCount() (_ int)                             { return }
func (self implementation) GetOptionsTooltip() (_ string)                        { return }
func (self implementation) GetOptionIcon(device int) (_ [1]gdclass.ImageTexture) { return }
func (self implementation) GetOptionLabel(device int) (_ string)                 { return }
func (self implementation) GetOptionTooltip(device int) (_ string)               { return }
func (self implementation) GetDeviceArchitecture(device int) (_ string)          { return }
func (self implementation) Cleanup()                                             { return }
func (self implementation) Run(preset [1]gdclass.EditorExportPreset, device int, debug_flags gdclass.EditorExportPlatformDebugFlags) (_ error) {
	return
}
func (self implementation) GetRunIcon() (_ [1]gdclass.Texture2D) { return }
func (self implementation) CanExport(preset [1]gdclass.EditorExportPreset, debug bool) (_ bool) {
	return
}
func (self implementation) HasValidExportConfiguration(preset [1]gdclass.EditorExportPreset, debug bool) (_ bool) {
	return
}
func (self implementation) HasValidProjectConfiguration(preset [1]gdclass.EditorExportPreset) (_ bool) {
	return
}
func (self implementation) GetBinaryExtensions(preset [1]gdclass.EditorExportPreset) (_ []string) {
	return
}
func (self implementation) ExportProject(preset [1]gdclass.EditorExportPreset, debug bool, path string, flags gdclass.EditorExportPlatformDebugFlags) (_ error) {
	return
}
func (self implementation) ExportPack(preset [1]gdclass.EditorExportPreset, debug bool, path string, flags gdclass.EditorExportPlatformDebugFlags) (_ error) {
	return
}
func (self implementation) ExportZip(preset [1]gdclass.EditorExportPreset, debug bool, path string, flags gdclass.EditorExportPlatformDebugFlags) (_ error) {
	return
}
func (self implementation) ExportPackPatch(preset [1]gdclass.EditorExportPreset, debug bool, path string, patches []string, flags gdclass.EditorExportPlatformDebugFlags) (_ error) {
	return
}
func (self implementation) ExportZipPatch(preset [1]gdclass.EditorExportPreset, debug bool, path string, patches []string, flags gdclass.EditorExportPlatformDebugFlags) (_ error) {
	return
}
func (self implementation) GetPlatformFeatures() (_ []string) { return }
func (self implementation) GetDebugProtocol() (_ string)      { return }

/*
[b]Required.[/b]
Returns array of platform specific features for the specified [param preset].
*/
func (Instance) _get_preset_features(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset)
		ptr, ok := pointers.End(gd.InternalPackedStrings(Packed.MakeStrings(ret...)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if specified file is a valid executable (native executable or script) for the target platform.
*/
func (Instance) _is_executable(impl func(ptr unsafe.Pointer, path string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns a property list, as an [Array] of dictionaries. Each [Dictionary] must at least contain the [code]name: StringName[/code] and [code]type: Variant.Type[/code] entries.
Additionally, the following keys are supported:
- [code]hint: PropertyHint[/code]
- [code]hint_string: String[/code]
- [code]usage: PropertyUsageFlags[/code]
- [code]class_name: StringName[/code]
- [code]default_value: Variant[/code], default value of the property.
- [code]update_visibility: bool[/code], if set to [code]true[/code], [method _get_export_option_visibility] is called for each property when this property is changed.
- [code]required: bool[/code], if set to [code]true[/code], this property warnings are critical, and should be resolved to make export possible. This value is a hint for the [method _has_valid_export_configuration] implementation, and not used by the engine directly.
See also [method Object._get_property_list].
*/
func (Instance) _get_export_options(impl func(ptr unsafe.Pointer) []map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[Dictionary.Any]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if export options list is changed and presets should be updated.
*/
func (Instance) _should_update_export_options(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Validates [param option] and returns visibility for the specified [param preset]. Default implementation return [code]true[/code] for all options.
*/
func (Instance) _get_export_option_visibility(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, option string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var option = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(option))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, option.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Validates [param option] and returns warning message for the specified [param preset]. Default implementation return empty string for all options.
*/
func (Instance) _get_export_option_warning(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, option string) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var option = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1)))))
		defer pointers.End(gd.InternalStringName(option))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, option.String())
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Returns target OS name.
*/
func (Instance) _get_os_name(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
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
[b]Required.[/b]
Returns export platform name.
*/
func (Instance) _get_name(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
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
[b]Required.[/b]
Returns platform logo displayed in the export dialog, logo should be 32x32 adjusted to the current editor scale, see [method EditorInterface.get_editor_scale].
*/
func (Instance) _get_logo(impl func(ptr unsafe.Pointer) [1]gdclass.Texture2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if one-click deploy options are changed and editor interface should be updated.
*/
func (Instance) _poll_export(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns number one-click deploy devices (or other one-click option displayed in the menu).
*/
func (Instance) _get_options_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
[b]Optional.[/b]
Returns tooltip of the one-click deploy menu button.
*/
func (Instance) _get_options_tooltip(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
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
[b]Optional.[/b]
Returns one-click deploy menu item icon for the specified [param device], icon should be 16x16 adjusted to the current editor scale, see [method EditorInterface.get_editor_scale].
*/
func (Instance) _get_option_icon(impl func(ptr unsafe.Pointer, device int) [1]gdclass.ImageTexture) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var device = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(device))
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns one-click deploy menu item label for the specified [param device].
*/
func (Instance) _get_option_label(impl func(ptr unsafe.Pointer, device int) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var device = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(device))
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns one-click deploy menu item tooltip for the specified [param device].
*/
func (Instance) _get_option_tooltip(impl func(ptr unsafe.Pointer, device int) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var device = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(device))
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns device architecture for one-click deploy.
*/
func (Instance) _get_device_architecture(impl func(ptr unsafe.Pointer, device int) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var device = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(device))
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Called by the editor before platform is unregistered.
*/
func (Instance) _cleanup(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
[b]Optional.[/b]
This method is called when [param device] one-click deploy menu option is selected.
Implementation should export project to a temporary location, upload and run it on the specific [param device], or perform another action associated with the menu item.
*/
func (Instance) _run(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, device int, debug_flags gdclass.EditorExportPlatformDebugFlags) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var device = gd.UnsafeGet[int64](p_args, 1)

		var debug_flags = gd.UnsafeGet[gdclass.EditorExportPlatformDebugFlags](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, int(device), debug_flags)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns icon of the one-click deploy menu button, icon should be 16x16 adjusted to the current editor scale, see [method EditorInterface.get_editor_scale].
*/
func (Instance) _get_run_icon(impl func(ptr unsafe.Pointer) [1]gdclass.Texture2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code], if specified [param preset] is valid and can be exported. Use [method set_config_error] and [method set_config_missing_templates] to set error details.
Usual implementation can call [method _has_valid_export_configuration] and [method _has_valid_project_configuration] to determine if export is possible.
*/
func (Instance) _can_export(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, debug bool) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, debug)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns [code]true[/code] if export configuration is valid.
*/
func (Instance) _has_valid_export_configuration(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, debug bool) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, debug)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns [code]true[/code] if project configuration is valid.
*/
func (Instance) _has_valid_project_configuration(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns array of supported binary extensions for the full project export.
*/
func (Instance) _get_binary_extensions(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset)
		ptr, ok := pointers.End(gd.InternalPackedStrings(Packed.MakeStrings(ret...)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Creates a full project at [param path] for the specified [param preset].
This method is called when "Export" button is pressed in the export dialog.
This method implementation can call [method EditorExportPlatform.save_pack] or [method EditorExportPlatform.save_zip] to use default PCK/ZIP export process, or calls [method EditorExportPlatform.export_project_files] and implement custom callback for processing each exported file.
*/
func (Instance) _export_project(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, debug bool, path string, flags gdclass.EditorExportPlatformDebugFlags) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(path))
		var flags = gd.UnsafeGet[gdclass.EditorExportPlatformDebugFlags](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, debug, path.String(), flags)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Creates a PCK archive at [param path] for the specified [param preset].
This method is called when "Export PCK/ZIP" button is pressed in the export dialog, with "Export as Patch" disabled, and PCK is selected as a file type.
*/
func (Instance) _export_pack(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, debug bool, path string, flags gdclass.EditorExportPlatformDebugFlags) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(path))
		var flags = gd.UnsafeGet[gdclass.EditorExportPlatformDebugFlags](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, debug, path.String(), flags)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Create a ZIP archive at [param path] for the specified [param preset].
This method is called when "Export PCK/ZIP" button is pressed in the export dialog, with "Export as Patch" disabled, and ZIP is selected as a file type.
*/
func (Instance) _export_zip(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, debug bool, path string, flags gdclass.EditorExportPlatformDebugFlags) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(path))
		var flags = gd.UnsafeGet[gdclass.EditorExportPlatformDebugFlags](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, debug, path.String(), flags)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Creates a patch PCK archive at [param path] for the specified [param preset], containing only the files that have changed since the last patch.
This method is called when "Export PCK/ZIP" button is pressed in the export dialog, with "Export as Patch" enabled, and PCK is selected as a file type.
[b]Note:[/b] The patches provided in [param patches] have already been loaded when this method is called and are merely provided as context. When empty the patches defined in the export preset have been loaded instead.
*/
func (Instance) _export_pack_patch(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, debug bool, path string, patches []string, flags gdclass.EditorExportPlatformDebugFlags) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(path))
		var patches = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 3)))))
		defer pointers.End(gd.InternalPackedStrings(patches))
		var flags = gd.UnsafeGet[gdclass.EditorExportPlatformDebugFlags](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, debug, path.String(), patches.Strings(), flags)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Create a ZIP archive at [param path] for the specified [param preset], containing only the files that have changed since the last patch.
This method is called when "Export PCK/ZIP" button is pressed in the export dialog, with "Export as Patch" enabled, and ZIP is selected as a file type.
[b]Note:[/b] The patches provided in [param patches] have already been loaded when this method is called and are merely provided as context. When empty the patches defined in the export preset have been loaded instead.
*/
func (Instance) _export_zip_patch(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, debug bool, path string, patches []string, flags gdclass.EditorExportPlatformDebugFlags) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(path))
		var patches = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 3)))))
		defer pointers.End(gd.InternalPackedStrings(patches))
		var flags = gd.UnsafeGet[gdclass.EditorExportPlatformDebugFlags](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, debug, path.String(), patches.Strings(), flags)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Returns array of platform specific features.
*/
func (Instance) _get_platform_features(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalPackedStrings(Packed.MakeStrings(ret...)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns protocol used for remote debugging. Default implementation return [code]tcp://[/code].
*/
func (Instance) _get_debug_protocol(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
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
Sets current configuration error message text. This method should be called only from the [method _can_export], [method _has_valid_export_configuration], or [method _has_valid_project_configuration] implementations.
*/
func (self Instance) SetConfigError(error_text string) { //gd:EditorExportPlatformExtension.set_config_error
	class(self).SetConfigError(String.New(error_text))
}

/*
Returns current configuration error message text. This method should be called only from the [method _can_export], [method _has_valid_export_configuration], or [method _has_valid_project_configuration] implementations.
*/
func (self Instance) GetConfigError() string { //gd:EditorExportPlatformExtension.get_config_error
	return string(class(self).GetConfigError().String())
}

/*
Set to [code]true[/code] is export templates are missing from the current configuration. This method should be called only from the [method _can_export], [method _has_valid_export_configuration], or [method _has_valid_project_configuration] implementations.
*/
func (self Instance) SetConfigMissingTemplates(missing_templates bool) { //gd:EditorExportPlatformExtension.set_config_missing_templates
	class(self).SetConfigMissingTemplates(missing_templates)
}

/*
Returns [code]true[/code] is export templates are missing from the current configuration. This method should be called only from the [method _can_export], [method _has_valid_export_configuration], or [method _has_valid_project_configuration] implementations.
*/
func (self Instance) GetConfigMissingTemplates() bool { //gd:EditorExportPlatformExtension.get_config_missing_templates
	return bool(class(self).GetConfigMissingTemplates())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorExportPlatformExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorExportPlatformExtension"))
	casted := Instance{*(*gdclass.EditorExportPlatformExtension)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
[b]Required.[/b]
Returns array of platform specific features for the specified [param preset].
*/
func (class) _get_preset_features(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset) Packed.Strings) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset)
		ptr, ok := pointers.End(gd.InternalPackedStrings(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if specified file is a valid executable (native executable or script) for the target platform.
*/
func (class) _is_executable(impl func(ptr unsafe.Pointer, path String.Readable) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns a property list, as an [Array] of dictionaries. Each [Dictionary] must at least contain the [code]name: StringName[/code] and [code]type: Variant.Type[/code] entries.
Additionally, the following keys are supported:
- [code]hint: PropertyHint[/code]
- [code]hint_string: String[/code]
- [code]usage: PropertyUsageFlags[/code]
- [code]class_name: StringName[/code]
- [code]default_value: Variant[/code], default value of the property.
- [code]update_visibility: bool[/code], if set to [code]true[/code], [method _get_export_option_visibility] is called for each property when this property is changed.
- [code]required: bool[/code], if set to [code]true[/code], this property warnings are critical, and should be resolved to make export possible. This value is a hint for the [method _has_valid_export_configuration] implementation, and not used by the engine directly.
See also [method Object._get_property_list].
*/
func (class) _get_export_options(impl func(ptr unsafe.Pointer) Array.Contains[Dictionary.Any]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if export options list is changed and presets should be updated.
*/
func (class) _should_update_export_options(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Validates [param option] and returns visibility for the specified [param preset]. Default implementation return [code]true[/code] for all options.
*/
func (class) _get_export_option_visibility(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, option String.Readable) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var option = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(option))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, option)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Validates [param option] and returns warning message for the specified [param preset]. Default implementation return empty string for all options.
*/
func (class) _get_export_option_warning(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, option String.Name) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var option = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1)))))
		defer pointers.End(gd.InternalStringName(option))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, option)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Returns target OS name.
*/
func (class) _get_os_name(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
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
[b]Required.[/b]
Returns export platform name.
*/
func (class) _get_name(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
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
[b]Required.[/b]
Returns platform logo displayed in the export dialog, logo should be 32x32 adjusted to the current editor scale, see [method EditorInterface.get_editor_scale].
*/
func (class) _get_logo(impl func(ptr unsafe.Pointer) [1]gdclass.Texture2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if one-click deploy options are changed and editor interface should be updated.
*/
func (class) _poll_export(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns number one-click deploy devices (or other one-click option displayed in the menu).
*/
func (class) _get_options_count(impl func(ptr unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns tooltip of the one-click deploy menu button.
*/
func (class) _get_options_tooltip(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
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
[b]Optional.[/b]
Returns one-click deploy menu item icon for the specified [param device], icon should be 16x16 adjusted to the current editor scale, see [method EditorInterface.get_editor_scale].
*/
func (class) _get_option_icon(impl func(ptr unsafe.Pointer, device int64) [1]gdclass.ImageTexture) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var device = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, device)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns one-click deploy menu item label for the specified [param device].
*/
func (class) _get_option_label(impl func(ptr unsafe.Pointer, device int64) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var device = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, device)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns one-click deploy menu item tooltip for the specified [param device].
*/
func (class) _get_option_tooltip(impl func(ptr unsafe.Pointer, device int64) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var device = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, device)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns device architecture for one-click deploy.
*/
func (class) _get_device_architecture(impl func(ptr unsafe.Pointer, device int64) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var device = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, device)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Called by the editor before platform is unregistered.
*/
func (class) _cleanup(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
[b]Optional.[/b]
This method is called when [param device] one-click deploy menu option is selected.
Implementation should export project to a temporary location, upload and run it on the specific [param device], or perform another action associated with the menu item.
*/
func (class) _run(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, device int64, debug_flags gdclass.EditorExportPlatformDebugFlags) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var device = gd.UnsafeGet[int64](p_args, 1)

		var debug_flags = gd.UnsafeGet[gdclass.EditorExportPlatformDebugFlags](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, device, debug_flags)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns icon of the one-click deploy menu button, icon should be 16x16 adjusted to the current editor scale, see [method EditorInterface.get_editor_scale].
*/
func (class) _get_run_icon(impl func(ptr unsafe.Pointer) [1]gdclass.Texture2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code], if specified [param preset] is valid and can be exported. Use [method set_config_error] and [method set_config_missing_templates] to set error details.
Usual implementation can call [method _has_valid_export_configuration] and [method _has_valid_project_configuration] to determine if export is possible.
*/
func (class) _can_export(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, debug bool) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, debug)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns [code]true[/code] if export configuration is valid.
*/
func (class) _has_valid_export_configuration(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, debug bool) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, debug)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns [code]true[/code] if project configuration is valid.
*/
func (class) _has_valid_project_configuration(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns array of supported binary extensions for the full project export.
*/
func (class) _get_binary_extensions(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset) Packed.Strings) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset)
		ptr, ok := pointers.End(gd.InternalPackedStrings(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Creates a full project at [param path] for the specified [param preset].
This method is called when "Export" button is pressed in the export dialog.
This method implementation can call [method EditorExportPlatform.save_pack] or [method EditorExportPlatform.save_zip] to use default PCK/ZIP export process, or calls [method EditorExportPlatform.export_project_files] and implement custom callback for processing each exported file.
*/
func (class) _export_project(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, debug bool, path String.Readable, flags gdclass.EditorExportPlatformDebugFlags) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(path))
		var flags = gd.UnsafeGet[gdclass.EditorExportPlatformDebugFlags](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, debug, path, flags)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Creates a PCK archive at [param path] for the specified [param preset].
This method is called when "Export PCK/ZIP" button is pressed in the export dialog, with "Export as Patch" disabled, and PCK is selected as a file type.
*/
func (class) _export_pack(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, debug bool, path String.Readable, flags gdclass.EditorExportPlatformDebugFlags) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(path))
		var flags = gd.UnsafeGet[gdclass.EditorExportPlatformDebugFlags](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, debug, path, flags)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Create a ZIP archive at [param path] for the specified [param preset].
This method is called when "Export PCK/ZIP" button is pressed in the export dialog, with "Export as Patch" disabled, and ZIP is selected as a file type.
*/
func (class) _export_zip(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, debug bool, path String.Readable, flags gdclass.EditorExportPlatformDebugFlags) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(path))
		var flags = gd.UnsafeGet[gdclass.EditorExportPlatformDebugFlags](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, debug, path, flags)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Creates a patch PCK archive at [param path] for the specified [param preset], containing only the files that have changed since the last patch.
This method is called when "Export PCK/ZIP" button is pressed in the export dialog, with "Export as Patch" enabled, and PCK is selected as a file type.
[b]Note:[/b] The patches provided in [param patches] have already been loaded when this method is called and are merely provided as context. When empty the patches defined in the export preset have been loaded instead.
*/
func (class) _export_pack_patch(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, debug bool, path String.Readable, patches Packed.Strings, flags gdclass.EditorExportPlatformDebugFlags) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(path))
		var patches = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 3)))))
		defer pointers.End(gd.InternalPackedStrings(patches))
		var flags = gd.UnsafeGet[gdclass.EditorExportPlatformDebugFlags](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, debug, path, patches, flags)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Create a ZIP archive at [param path] for the specified [param preset], containing only the files that have changed since the last patch.
This method is called when "Export PCK/ZIP" button is pressed in the export dialog, with "Export as Patch" enabled, and ZIP is selected as a file type.
[b]Note:[/b] The patches provided in [param patches] have already been loaded when this method is called and are merely provided as context. When empty the patches defined in the export preset have been loaded instead.
*/
func (class) _export_zip_patch(impl func(ptr unsafe.Pointer, preset [1]gdclass.EditorExportPreset, debug bool, path String.Readable, patches Packed.Strings, flags gdclass.EditorExportPlatformDebugFlags) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var preset = [1]gdclass.EditorExportPreset{pointers.New[gdclass.EditorExportPreset]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(preset[0])
		var debug = gd.UnsafeGet[bool](p_args, 1)

		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(path))
		var patches = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 3)))))
		defer pointers.End(gd.InternalPackedStrings(patches))
		var flags = gd.UnsafeGet[gdclass.EditorExportPlatformDebugFlags](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, preset, debug, path, patches, flags)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Returns array of platform specific features.
*/
func (class) _get_platform_features(impl func(ptr unsafe.Pointer) Packed.Strings) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalPackedStrings(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns protocol used for remote debugging. Default implementation return [code]tcp://[/code].
*/
func (class) _get_debug_protocol(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
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
Sets current configuration error message text. This method should be called only from the [method _can_export], [method _has_valid_export_configuration], or [method _has_valid_project_configuration] implementations.
*/
//go:nosplit
func (self class) SetConfigError(error_text String.Readable) { //gd:EditorExportPlatformExtension.set_config_error
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(error_text)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorExportPlatformExtension.Bind_set_config_error, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns current configuration error message text. This method should be called only from the [method _can_export], [method _has_valid_export_configuration], or [method _has_valid_project_configuration] implementations.
*/
//go:nosplit
func (self class) GetConfigError() String.Readable { //gd:EditorExportPlatformExtension.get_config_error
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorExportPlatformExtension.Bind_get_config_error, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Set to [code]true[/code] is export templates are missing from the current configuration. This method should be called only from the [method _can_export], [method _has_valid_export_configuration], or [method _has_valid_project_configuration] implementations.
*/
//go:nosplit
func (self class) SetConfigMissingTemplates(missing_templates bool) { //gd:EditorExportPlatformExtension.set_config_missing_templates
	var frame = callframe.New()
	callframe.Arg(frame, missing_templates)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorExportPlatformExtension.Bind_set_config_missing_templates, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] is export templates are missing from the current configuration. This method should be called only from the [method _can_export], [method _has_valid_export_configuration], or [method _has_valid_project_configuration] implementations.
*/
//go:nosplit
func (self class) GetConfigMissingTemplates() bool { //gd:EditorExportPlatformExtension.get_config_missing_templates
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorExportPlatformExtension.Bind_get_config_missing_templates, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsEditorExportPlatformExtension() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorExportPlatformExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsEditorExportPlatform() EditorExportPlatform.Advanced {
	return *((*EditorExportPlatform.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorExportPlatform() EditorExportPlatform.Instance {
	return *((*EditorExportPlatform.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_preset_features":
		return reflect.ValueOf(self._get_preset_features)
	case "_is_executable":
		return reflect.ValueOf(self._is_executable)
	case "_get_export_options":
		return reflect.ValueOf(self._get_export_options)
	case "_should_update_export_options":
		return reflect.ValueOf(self._should_update_export_options)
	case "_get_export_option_visibility":
		return reflect.ValueOf(self._get_export_option_visibility)
	case "_get_export_option_warning":
		return reflect.ValueOf(self._get_export_option_warning)
	case "_get_os_name":
		return reflect.ValueOf(self._get_os_name)
	case "_get_name":
		return reflect.ValueOf(self._get_name)
	case "_get_logo":
		return reflect.ValueOf(self._get_logo)
	case "_poll_export":
		return reflect.ValueOf(self._poll_export)
	case "_get_options_count":
		return reflect.ValueOf(self._get_options_count)
	case "_get_options_tooltip":
		return reflect.ValueOf(self._get_options_tooltip)
	case "_get_option_icon":
		return reflect.ValueOf(self._get_option_icon)
	case "_get_option_label":
		return reflect.ValueOf(self._get_option_label)
	case "_get_option_tooltip":
		return reflect.ValueOf(self._get_option_tooltip)
	case "_get_device_architecture":
		return reflect.ValueOf(self._get_device_architecture)
	case "_cleanup":
		return reflect.ValueOf(self._cleanup)
	case "_run":
		return reflect.ValueOf(self._run)
	case "_get_run_icon":
		return reflect.ValueOf(self._get_run_icon)
	case "_can_export":
		return reflect.ValueOf(self._can_export)
	case "_has_valid_export_configuration":
		return reflect.ValueOf(self._has_valid_export_configuration)
	case "_has_valid_project_configuration":
		return reflect.ValueOf(self._has_valid_project_configuration)
	case "_get_binary_extensions":
		return reflect.ValueOf(self._get_binary_extensions)
	case "_export_project":
		return reflect.ValueOf(self._export_project)
	case "_export_pack":
		return reflect.ValueOf(self._export_pack)
	case "_export_zip":
		return reflect.ValueOf(self._export_zip)
	case "_export_pack_patch":
		return reflect.ValueOf(self._export_pack_patch)
	case "_export_zip_patch":
		return reflect.ValueOf(self._export_zip_patch)
	case "_get_platform_features":
		return reflect.ValueOf(self._get_platform_features)
	case "_get_debug_protocol":
		return reflect.ValueOf(self._get_debug_protocol)
	default:
		return gd.VirtualByName(EditorExportPlatform.Advanced(self.AsEditorExportPlatform()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_preset_features":
		return reflect.ValueOf(self._get_preset_features)
	case "_is_executable":
		return reflect.ValueOf(self._is_executable)
	case "_get_export_options":
		return reflect.ValueOf(self._get_export_options)
	case "_should_update_export_options":
		return reflect.ValueOf(self._should_update_export_options)
	case "_get_export_option_visibility":
		return reflect.ValueOf(self._get_export_option_visibility)
	case "_get_export_option_warning":
		return reflect.ValueOf(self._get_export_option_warning)
	case "_get_os_name":
		return reflect.ValueOf(self._get_os_name)
	case "_get_name":
		return reflect.ValueOf(self._get_name)
	case "_get_logo":
		return reflect.ValueOf(self._get_logo)
	case "_poll_export":
		return reflect.ValueOf(self._poll_export)
	case "_get_options_count":
		return reflect.ValueOf(self._get_options_count)
	case "_get_options_tooltip":
		return reflect.ValueOf(self._get_options_tooltip)
	case "_get_option_icon":
		return reflect.ValueOf(self._get_option_icon)
	case "_get_option_label":
		return reflect.ValueOf(self._get_option_label)
	case "_get_option_tooltip":
		return reflect.ValueOf(self._get_option_tooltip)
	case "_get_device_architecture":
		return reflect.ValueOf(self._get_device_architecture)
	case "_cleanup":
		return reflect.ValueOf(self._cleanup)
	case "_run":
		return reflect.ValueOf(self._run)
	case "_get_run_icon":
		return reflect.ValueOf(self._get_run_icon)
	case "_can_export":
		return reflect.ValueOf(self._can_export)
	case "_has_valid_export_configuration":
		return reflect.ValueOf(self._has_valid_export_configuration)
	case "_has_valid_project_configuration":
		return reflect.ValueOf(self._has_valid_project_configuration)
	case "_get_binary_extensions":
		return reflect.ValueOf(self._get_binary_extensions)
	case "_export_project":
		return reflect.ValueOf(self._export_project)
	case "_export_pack":
		return reflect.ValueOf(self._export_pack)
	case "_export_zip":
		return reflect.ValueOf(self._export_zip)
	case "_export_pack_patch":
		return reflect.ValueOf(self._export_pack_patch)
	case "_export_zip_patch":
		return reflect.ValueOf(self._export_zip_patch)
	case "_get_platform_features":
		return reflect.ValueOf(self._get_platform_features)
	case "_get_debug_protocol":
		return reflect.ValueOf(self._get_debug_protocol)
	default:
		return gd.VirtualByName(EditorExportPlatform.Instance(self.AsEditorExportPlatform()), name)
	}
}
func init() {
	gdclass.Register("EditorExportPlatformExtension", func(ptr gd.Object) any {
		return [1]gdclass.EditorExportPlatformExtension{*(*gdclass.EditorExportPlatformExtension)(unsafe.Pointer(&ptr))}
	})
}
