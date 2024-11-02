package EditorPaths

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This editor-only singleton returns OS-specific paths to various data folders and files. It can be used in editor plugins to ensure files are saved in the correct location on each operating system.
[b]Note:[/b] This singleton is not accessible in exported projects. Attempting to access it in an exported project will result in a script error as the singleton won't be declared. To prevent script errors in exported projects, use [method Engine.has_singleton] to check whether the singleton is available before using it.
[b]Note:[/b] On the Linux/BSD platform, Godot complies with the [url=https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html]XDG Base Directory Specification[/url]. You can override environment variables following the specification to change the editor and project data paths.
*/
type Instance [1]classdb.EditorPaths

/*
Returns the absolute path to the user's data folder. This folder should be used for [i]persistent[/i] user data files such as installed export templates.
[b]Default paths per platform:[/b]
[codeblock lang=text]
- Windows: %APPDATA%\Godot\                    (same as `get_config_dir()`)
- macOS: ~/Library/Application Support/Godot/  (same as `get_config_dir()`)
- Linux: ~/.local/share/godot/
[/codeblock]
*/
func (self Instance) GetDataDir() string {
	return string(class(self).GetDataDir().String())
}

/*
Returns the absolute path to the user's configuration folder. This folder should be used for [i]persistent[/i] user configuration files.
[b]Default paths per platform:[/b]
[codeblock lang=text]
- Windows: %APPDATA%\Godot\                    (same as `get_data_dir()`)
- macOS: ~/Library/Application Support/Godot/  (same as `get_data_dir()`)
- Linux: ~/.config/godot/
[/codeblock]
*/
func (self Instance) GetConfigDir() string {
	return string(class(self).GetConfigDir().String())
}

/*
Returns the absolute path to the user's cache folder. This folder should be used for temporary data that can be removed safely whenever the editor is closed (such as generated resource thumbnails).
[b]Default paths per platform:[/b]
[codeblock lang=text]
- Windows: %LOCALAPPDATA%\Godot\
- macOS: ~/Library/Caches/Godot/
- Linux: ~/.cache/godot/
[/codeblock]
*/
func (self Instance) GetCacheDir() string {
	return string(class(self).GetCacheDir().String())
}

/*
Returns [code]true[/code] if the editor is marked as self-contained, [code]false[/code] otherwise. When self-contained mode is enabled, user configuration, data and cache files are saved in an [code]editor_data/[/code] folder next to the editor binary. This makes portable usage easier and ensures the Godot editor minimizes file writes outside its own folder. Self-contained mode is not available for exported projects.
Self-contained mode can be enabled by creating a file named [code]._sc_[/code] or [code]_sc_[/code] in the same folder as the editor binary or macOS .app bundle while the editor is not running. See also [method get_self_contained_file].
[b]Note:[/b] On macOS, quarantine flag should be manually removed before using self-contained mode, see [url=https://docs.godotengine.org/en/stable/tutorials/export/running_on_macos.html]Running on macOS[/url].
[b]Note:[/b] On macOS, placing [code]_sc_[/code] or any other file inside .app bundle will break digital signature and make it non-portable, consider placing it in the same folder as the .app bundle instead.
[b]Note:[/b] The Steam release of Godot uses self-contained mode by default.
*/
func (self Instance) IsSelfContained() bool {
	return bool(class(self).IsSelfContained())
}

/*
Returns the absolute path to the self-contained file that makes the current Godot editor instance be considered as self-contained. Returns an empty string if the current Godot editor instance isn't self-contained. See also [method is_self_contained].
*/
func (self Instance) GetSelfContainedFile() string {
	return string(class(self).GetSelfContainedFile().String())
}

/*
Returns the project-specific editor settings path. Projects all have a unique subdirectory inside the settings path where project-specific editor settings are saved.
*/
func (self Instance) GetProjectSettingsDir() string {
	return string(class(self).GetProjectSettingsDir().String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EditorPaths

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorPaths"))
	return Instance{classdb.EditorPaths(object)}
}

/*
Returns the absolute path to the user's data folder. This folder should be used for [i]persistent[/i] user data files such as installed export templates.
[b]Default paths per platform:[/b]
[codeblock lang=text]
- Windows: %APPDATA%\Godot\                    (same as `get_config_dir()`)
- macOS: ~/Library/Application Support/Godot/  (same as `get_config_dir()`)
- Linux: ~/.local/share/godot/
[/codeblock]
*/
//go:nosplit
func (self class) GetDataDir() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPaths.Bind_get_data_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the absolute path to the user's configuration folder. This folder should be used for [i]persistent[/i] user configuration files.
[b]Default paths per platform:[/b]
[codeblock lang=text]
- Windows: %APPDATA%\Godot\                    (same as `get_data_dir()`)
- macOS: ~/Library/Application Support/Godot/  (same as `get_data_dir()`)
- Linux: ~/.config/godot/
[/codeblock]
*/
//go:nosplit
func (self class) GetConfigDir() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPaths.Bind_get_config_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the absolute path to the user's cache folder. This folder should be used for temporary data that can be removed safely whenever the editor is closed (such as generated resource thumbnails).
[b]Default paths per platform:[/b]
[codeblock lang=text]
- Windows: %LOCALAPPDATA%\Godot\
- macOS: ~/Library/Caches/Godot/
- Linux: ~/.cache/godot/
[/codeblock]
*/
//go:nosplit
func (self class) GetCacheDir() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPaths.Bind_get_cache_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the editor is marked as self-contained, [code]false[/code] otherwise. When self-contained mode is enabled, user configuration, data and cache files are saved in an [code]editor_data/[/code] folder next to the editor binary. This makes portable usage easier and ensures the Godot editor minimizes file writes outside its own folder. Self-contained mode is not available for exported projects.
Self-contained mode can be enabled by creating a file named [code]._sc_[/code] or [code]_sc_[/code] in the same folder as the editor binary or macOS .app bundle while the editor is not running. See also [method get_self_contained_file].
[b]Note:[/b] On macOS, quarantine flag should be manually removed before using self-contained mode, see [url=https://docs.godotengine.org/en/stable/tutorials/export/running_on_macos.html]Running on macOS[/url].
[b]Note:[/b] On macOS, placing [code]_sc_[/code] or any other file inside .app bundle will break digital signature and make it non-portable, consider placing it in the same folder as the .app bundle instead.
[b]Note:[/b] The Steam release of Godot uses self-contained mode by default.
*/
//go:nosplit
func (self class) IsSelfContained() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPaths.Bind_is_self_contained, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the absolute path to the self-contained file that makes the current Godot editor instance be considered as self-contained. Returns an empty string if the current Godot editor instance isn't self-contained. See also [method is_self_contained].
*/
//go:nosplit
func (self class) GetSelfContainedFile() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPaths.Bind_get_self_contained_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the project-specific editor settings path. Projects all have a unique subdirectory inside the settings path where project-specific editor settings are saved.
*/
//go:nosplit
func (self class) GetProjectSettingsDir() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPaths.Bind_get_project_settings_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsEditorPaths() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorPaths() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("EditorPaths", func(ptr gd.Object) any { return classdb.EditorPaths(ptr) })
}
