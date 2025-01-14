// Package ProjectSettings provides methods for working with ProjectSettings object instances.
package ProjectSettings

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Dictionary"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Stores variables that can be accessed from everywhere. Use [method get_setting], [method set_setting] or [method has_setting] to access them. Variables stored in [code]project.godot[/code] are also loaded into [ProjectSettings], making this object very useful for reading custom game configuration options.
When naming a Project Settings property, use the full path to the setting including the category. For example, [code]"application/config/name"[/code] for the project name. Category and property names can be viewed in the Project Settings dialog.
[b]Feature tags:[/b] Project settings can be overridden for specific platforms and configurations (debug, release, ...) using [url=$DOCS_URL/tutorials/export/feature_tags.html]feature tags[/url].
[b]Overriding:[/b] Any project setting can be overridden by creating a file named [code]override.cfg[/code] in the project's root directory. This can also be used in exported projects by placing this file in the same directory as the project binary. Overriding will still take the base project settings' [url=$DOCS_URL/tutorials/export/feature_tags.html]feature tags[/url] in account. Therefore, make sure to [i]also[/i] override the setting with the desired feature tags if you want them to override base project settings on all platforms and configurations.
*/
var self [1]gdclass.ProjectSettings
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.ProjectSettings)
	self = *(*[1]gdclass.ProjectSettings)(unsafe.Pointer(&obj))
}

/*
Returns [code]true[/code] if a configuration value is present.
*/
func HasSetting(name string) bool {
	once.Do(singleton)
	return bool(class(self).HasSetting(gd.NewString(name)))
}

/*
Sets the value of a setting.
[b]Example:[/b]
[codeblocks]
[gdscript]
ProjectSettings.set_setting("application/config/name", "Example")
[/gdscript]
[csharp]
ProjectSettings.SetSetting("application/config/name", "Example");
[/csharp]
[/codeblocks]
This can also be used to erase custom project settings. To do this change the setting value to [code]null[/code].
*/
func SetSetting(name string, value any) {
	once.Do(singleton)
	class(self).SetSetting(gd.NewString(name), gd.NewVariant(value))
}

/*
Returns the value of the setting identified by [param name]. If the setting doesn't exist and [param default_value] is specified, the value of [param default_value] is returned. Otherwise, [code]null[/code] is returned.
[b]Example:[/b]
[codeblocks]
[gdscript]
print(ProjectSettings.get_setting("application/config/name"))
print(ProjectSettings.get_setting("application/config/custom_description", "No description specified."))
[/gdscript]
[csharp]
GD.Print(ProjectSettings.GetSetting("application/config/name"));
GD.Print(ProjectSettings.GetSetting("application/config/custom_description", "No description specified."));
[/csharp]
[/codeblocks]
[b]Note:[/b] This method doesn't take potential feature overrides into account automatically. Use [method get_setting_with_override] to handle seamlessly.
*/
func GetSetting(name string) any {
	once.Do(singleton)
	return any(class(self).GetSetting(gd.NewString(name), gd.NewVariant(gd.NewVariant(([1]any{}[0])))).Interface())
}

/*
Similar to [method get_setting], but applies feature tag overrides if any exists and is valid.
[b]Example:[/b]
If the following setting override exists "application/config/name.windows", and the following code is executed:
[codeblocks]
[gdscript]
print(ProjectSettings.get_setting_with_override("application/config/name"))
[/gdscript]
[csharp]
GD.Print(ProjectSettings.GetSettingWithOverride("application/config/name"));
[/csharp]
[/codeblocks]
Then the overridden setting will be returned instead if the project is running on the [i]Windows[/i] operating system.
*/
func GetSettingWithOverride(name string) any {
	once.Do(singleton)
	return any(class(self).GetSettingWithOverride(gd.NewStringName(name)).Interface())
}

/*
Returns an [Array] of registered global classes. Each global class is represented as a [Dictionary] that contains the following entries:
- [code]base[/code] is a name of the base class;
- [code]class[/code] is a name of the registered global class;
- [code]icon[/code] is a path to a custom icon of the global class, if it has any;
- [code]language[/code] is a name of a programming language in which the global class is written;
- [code]path[/code] is a path to a file containing the global class.
[b]Note:[/b] Both the script and the icon paths are local to the project filesystem, i.e. they start with [code]res://[/code].
*/
func GetGlobalClassList() gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).GetGlobalClassList())
}

/*
Sets the order of a configuration value (influences when saved to the config file).
*/
func SetOrder(name string, position int) {
	once.Do(singleton)
	class(self).SetOrder(gd.NewString(name), gd.Int(position))
}

/*
Returns the order of a configuration value (influences when saved to the config file).
*/
func GetOrder(name string) int {
	once.Do(singleton)
	return int(int(class(self).GetOrder(gd.NewString(name))))
}

/*
Sets the specified setting's initial value. This is the value the setting reverts to.
*/
func SetInitialValue(name string, value any) {
	once.Do(singleton)
	class(self).SetInitialValue(gd.NewString(name), gd.NewVariant(value))
}

/*
Defines if the specified setting is considered basic or advanced. Basic settings will always be shown in the project settings. Advanced settings will only be shown if the user enables the "Advanced Settings" option.
*/
func SetAsBasic(name string, basic bool) {
	once.Do(singleton)
	class(self).SetAsBasic(gd.NewString(name), basic)
}

/*
Defines if the specified setting is considered internal. An internal setting won't show up in the Project Settings dialog. This is mostly useful for addons that need to store their own internal settings without exposing them directly to the user.
*/
func SetAsInternal(name string, internal_ bool) {
	once.Do(singleton)
	class(self).SetAsInternal(gd.NewString(name), internal_)
}

/*
Adds a custom property info to a property. The dictionary must contain:
- [code]"name"[/code]: [String] (the property's name)
- [code]"type"[/code]: [int] (see [enum Variant.Type])
- optionally [code]"hint"[/code]: [int] (see [enum PropertyHint]) and [code]"hint_string"[/code]: [String]
[b]Example:[/b]
[codeblocks]
[gdscript]
ProjectSettings.set("category/property_name", 0)

	var property_info = {
	    "name": "category/property_name",
	    "type": TYPE_INT,
	    "hint": PROPERTY_HINT_ENUM,
	    "hint_string": "one,two,three"
	}

ProjectSettings.add_property_info(property_info)
[/gdscript]
[csharp]
ProjectSettings.Singleton.Set("category/property_name", 0);

var propertyInfo = new Godot.Collections.Dictionary

	{
	    {"name", "category/propertyName"},
	    {"type", (int)Variant.Type.Int},
	    {"hint", (int)PropertyHint.Enum},
	    {"hint_string", "one,two,three"},
	};

ProjectSettings.AddPropertyInfo(propertyInfo);
[/csharp]
[/codeblocks]
*/
func AddPropertyInfo(hint Dictionary.Any) {
	once.Do(singleton)
	class(self).AddPropertyInfo(hint)
}

/*
Sets whether a setting requires restarting the editor to properly take effect.
[b]Note:[/b] This is just a hint to display to the user that the editor must be restarted for changes to take effect. Enabling [method set_restart_if_changed] does [i]not[/i] delay the setting being set when changed.
*/
func SetRestartIfChanged(name string, restart bool) {
	once.Do(singleton)
	class(self).SetRestartIfChanged(gd.NewString(name), restart)
}

/*
Clears the whole configuration (not recommended, may break things).
*/
func Clear(name string) {
	once.Do(singleton)
	class(self).Clear(gd.NewString(name))
}

/*
Returns the localized path (starting with [code]res://[/code]) corresponding to the absolute, native OS [param path]. See also [method globalize_path].
*/
func LocalizePath(path string) string {
	once.Do(singleton)
	return string(class(self).LocalizePath(gd.NewString(path)).String())
}

/*
Returns the absolute, native OS path corresponding to the localized [param path] (starting with [code]res://[/code] or [code]user://[/code]). The returned path will vary depending on the operating system and user preferences. See [url=$DOCS_URL/tutorials/io/data_paths.html]File paths in Godot projects[/url] to see what those paths convert to. See also [method localize_path].
[b]Note:[/b] [method globalize_path] with [code]res://[/code] will not work in an exported project. Instead, prepend the executable's base directory to the path when running from an exported project:
[codeblock]
var path = ""
if OS.has_feature("editor"):

	# Running from an editor binary.
	# `path` will contain the absolute path to `hello.txt` located in the project root.
	path = ProjectSettings.globalize_path("res://hello.txt")

else:

	# Running from an exported project.
	# `path` will contain the absolute path to `hello.txt` next to the executable.
	# This is *not* identical to using `ProjectSettings.globalize_path()` with a `res://` path,
	# but is close enough in spirit.
	path = OS.get_executable_path().get_base_dir().path_join("hello.txt")

[/codeblock]
*/
func GlobalizePath(path string) string {
	once.Do(singleton)
	return string(class(self).GlobalizePath(gd.NewString(path)).String())
}

/*
Saves the configuration to the [code]project.godot[/code] file.
[b]Note:[/b] This method is intended to be used by editor plugins, as modified [ProjectSettings] can't be loaded back in the running app. If you want to change project settings in exported projects, use [method save_custom] to save [code]override.cfg[/code] file.
*/
func Save() error {
	once.Do(singleton)
	return error(class(self).Save())
}

/*
Loads the contents of the .pck or .zip file specified by [param pack] into the resource filesystem ([code]res://[/code]). Returns [code]true[/code] on success.
[b]Note:[/b] If a file from [param pack] shares the same path as a file already in the resource filesystem, any attempts to load that file will use the file from [param pack] unless [param replace_files] is set to [code]false[/code].
[b]Note:[/b] The optional [param offset] parameter can be used to specify the offset in bytes to the start of the resource pack. This is only supported for .pck files.
*/
func LoadResourcePack(pack string) bool {
	once.Do(singleton)
	return bool(class(self).LoadResourcePack(gd.NewString(pack), true, gd.Int(0)))
}

/*
Saves the configuration to a custom file. The file extension must be [code].godot[/code] (to save in text-based [ConfigFile] format) or [code].binary[/code] (to save in binary format). You can also save [code]override.cfg[/code] file, which is also text, but can be used in exported projects unlike other formats.
*/
func SaveCustom(file string) error {
	once.Do(singleton)
	return error(class(self).SaveCustom(gd.NewString(file)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.ProjectSettings

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Returns [code]true[/code] if a configuration value is present.
*/
//go:nosplit
func (self class) HasSetting(name gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_has_setting, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the value of a setting.
[b]Example:[/b]
[codeblocks]
[gdscript]
ProjectSettings.set_setting("application/config/name", "Example")
[/gdscript]
[csharp]
ProjectSettings.SetSetting("application/config/name", "Example");
[/csharp]
[/codeblocks]
This can also be used to erase custom project settings. To do this change the setting value to [code]null[/code].
*/
//go:nosplit
func (self class) SetSetting(name gd.String, value gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(value))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_set_setting, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the value of the setting identified by [param name]. If the setting doesn't exist and [param default_value] is specified, the value of [param default_value] is returned. Otherwise, [code]null[/code] is returned.
[b]Example:[/b]
[codeblocks]
[gdscript]
print(ProjectSettings.get_setting("application/config/name"))
print(ProjectSettings.get_setting("application/config/custom_description", "No description specified."))
[/gdscript]
[csharp]
GD.Print(ProjectSettings.GetSetting("application/config/name"));
GD.Print(ProjectSettings.GetSetting("application/config/custom_description", "No description specified."));
[/csharp]
[/codeblocks]
[b]Note:[/b] This method doesn't take potential feature overrides into account automatically. Use [method get_setting_with_override] to handle seamlessly.
*/
//go:nosplit
func (self class) GetSetting(name gd.String, default_value gd.Variant) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(default_value))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_get_setting, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Similar to [method get_setting], but applies feature tag overrides if any exists and is valid.
[b]Example:[/b]
If the following setting override exists "application/config/name.windows", and the following code is executed:
[codeblocks]
[gdscript]
print(ProjectSettings.get_setting_with_override("application/config/name"))
[/gdscript]
[csharp]
GD.Print(ProjectSettings.GetSettingWithOverride("application/config/name"));
[/csharp]
[/codeblocks]
Then the overridden setting will be returned instead if the project is running on the [i]Windows[/i] operating system.
*/
//go:nosplit
func (self class) GetSettingWithOverride(name gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_get_setting_with_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an [Array] of registered global classes. Each global class is represented as a [Dictionary] that contains the following entries:
- [code]base[/code] is a name of the base class;
- [code]class[/code] is a name of the registered global class;
- [code]icon[/code] is a path to a custom icon of the global class, if it has any;
- [code]language[/code] is a name of a programming language in which the global class is written;
- [code]path[/code] is a path to a file containing the global class.
[b]Note:[/b] Both the script and the icon paths are local to the project filesystem, i.e. they start with [code]res://[/code].
*/
//go:nosplit
func (self class) GetGlobalClassList() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_get_global_class_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the order of a configuration value (influences when saved to the config file).
*/
//go:nosplit
func (self class) SetOrder(name gd.String, position gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_set_order, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the order of a configuration value (influences when saved to the config file).
*/
//go:nosplit
func (self class) GetOrder(name gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_get_order, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the specified setting's initial value. This is the value the setting reverts to.
*/
//go:nosplit
func (self class) SetInitialValue(name gd.String, value gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(value))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_set_initial_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Defines if the specified setting is considered basic or advanced. Basic settings will always be shown in the project settings. Advanced settings will only be shown if the user enables the "Advanced Settings" option.
*/
//go:nosplit
func (self class) SetAsBasic(name gd.String, basic bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, basic)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_set_as_basic, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Defines if the specified setting is considered internal. An internal setting won't show up in the Project Settings dialog. This is mostly useful for addons that need to store their own internal settings without exposing them directly to the user.
*/
//go:nosplit
func (self class) SetAsInternal(name gd.String, internal_ bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, internal_)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_set_as_internal, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a custom property info to a property. The dictionary must contain:
- [code]"name"[/code]: [String] (the property's name)
- [code]"type"[/code]: [int] (see [enum Variant.Type])
- optionally [code]"hint"[/code]: [int] (see [enum PropertyHint]) and [code]"hint_string"[/code]: [String]
[b]Example:[/b]
[codeblocks]
[gdscript]
ProjectSettings.set("category/property_name", 0)

var property_info = {
    "name": "category/property_name",
    "type": TYPE_INT,
    "hint": PROPERTY_HINT_ENUM,
    "hint_string": "one,two,three"
}

ProjectSettings.add_property_info(property_info)
[/gdscript]
[csharp]
ProjectSettings.Singleton.Set("category/property_name", 0);

var propertyInfo = new Godot.Collections.Dictionary
{
    {"name", "category/propertyName"},
    {"type", (int)Variant.Type.Int},
    {"hint", (int)PropertyHint.Enum},
    {"hint_string", "one,two,three"},
};

ProjectSettings.AddPropertyInfo(propertyInfo);
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) AddPropertyInfo(hint gd.Dictionary) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(hint))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_add_property_info, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets whether a setting requires restarting the editor to properly take effect.
[b]Note:[/b] This is just a hint to display to the user that the editor must be restarted for changes to take effect. Enabling [method set_restart_if_changed] does [i]not[/i] delay the setting being set when changed.
*/
//go:nosplit
func (self class) SetRestartIfChanged(name gd.String, restart bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, restart)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_set_restart_if_changed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Clears the whole configuration (not recommended, may break things).
*/
//go:nosplit
func (self class) Clear(name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the localized path (starting with [code]res://[/code]) corresponding to the absolute, native OS [param path]. See also [method globalize_path].
*/
//go:nosplit
func (self class) LocalizePath(path gd.String) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_localize_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the absolute, native OS path corresponding to the localized [param path] (starting with [code]res://[/code] or [code]user://[/code]). The returned path will vary depending on the operating system and user preferences. See [url=$DOCS_URL/tutorials/io/data_paths.html]File paths in Godot projects[/url] to see what those paths convert to. See also [method localize_path].
[b]Note:[/b] [method globalize_path] with [code]res://[/code] will not work in an exported project. Instead, prepend the executable's base directory to the path when running from an exported project:
[codeblock]
var path = ""
if OS.has_feature("editor"):
    # Running from an editor binary.
    # `path` will contain the absolute path to `hello.txt` located in the project root.
    path = ProjectSettings.globalize_path("res://hello.txt")
else:
    # Running from an exported project.
    # `path` will contain the absolute path to `hello.txt` next to the executable.
    # This is *not* identical to using `ProjectSettings.globalize_path()` with a `res://` path,
    # but is close enough in spirit.
    path = OS.get_executable_path().get_base_dir().path_join("hello.txt")
[/codeblock]
*/
//go:nosplit
func (self class) GlobalizePath(path gd.String) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_globalize_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Saves the configuration to the [code]project.godot[/code] file.
[b]Note:[/b] This method is intended to be used by editor plugins, as modified [ProjectSettings] can't be loaded back in the running app. If you want to change project settings in exported projects, use [method save_custom] to save [code]override.cfg[/code] file.
*/
//go:nosplit
func (self class) Save() gd.Error {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_save, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Loads the contents of the .pck or .zip file specified by [param pack] into the resource filesystem ([code]res://[/code]). Returns [code]true[/code] on success.
[b]Note:[/b] If a file from [param pack] shares the same path as a file already in the resource filesystem, any attempts to load that file will use the file from [param pack] unless [param replace_files] is set to [code]false[/code].
[b]Note:[/b] The optional [param offset] parameter can be used to specify the offset in bytes to the start of the resource pack. This is only supported for .pck files.
*/
//go:nosplit
func (self class) LoadResourcePack(pack gd.String, replace_files bool, offset gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(pack))
	callframe.Arg(frame, replace_files)
	callframe.Arg(frame, offset)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_load_resource_pack, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Saves the configuration to a custom file. The file extension must be [code].godot[/code] (to save in text-based [ConfigFile] format) or [code].binary[/code] (to save in binary format). You can also save [code]override.cfg[/code] file, which is also text, but can be used in exported projects unlike other formats.
*/
//go:nosplit
func (self class) SaveCustom(file gd.String) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(file))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_save_custom, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func OnSettingsChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("settings_changed"), gd.NewCallable(cb), 0)
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("ProjectSettings", func(ptr gd.Object) any {
		return [1]gdclass.ProjectSettings{*(*gdclass.ProjectSettings)(unsafe.Pointer(&ptr))}
	})
}

type Error = gd.Error

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
