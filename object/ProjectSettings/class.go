package ProjectSettings

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Stores variables that can be accessed from everywhere. Use [method get_setting], [method set_setting] or [method has_setting] to access them. Variables stored in [code]project.godot[/code] are also loaded into [ProjectSettings], making this object very useful for reading custom game configuration options.
When naming a Project Settings property, use the full path to the setting including the category. For example, [code]"application/config/name"[/code] for the project name. Category and property names can be viewed in the Project Settings dialog.
[b]Feature tags:[/b] Project settings can be overridden for specific platforms and configurations (debug, release, ...) using [url=$DOCS_URL/tutorials/export/feature_tags.html]feature tags[/url].
[b]Overriding:[/b] Any project setting can be overridden by creating a file named [code]override.cfg[/code] in the project's root directory. This can also be used in exported projects by placing this file in the same directory as the project binary. Overriding will still take the base project settings' [url=$DOCS_URL/tutorials/export/feature_tags.html]feature tags[/url] in account. Therefore, make sure to [i]also[/i] override the setting with the desired feature tags if you want them to override base project settings on all platforms and configurations.

*/
type Simple [1]classdb.ProjectSettings
func (self Simple) HasSetting(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasSetting(gc.String(name)))
}
func (self Simple) SetSetting(name string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSetting(gc.String(name), value)
}
func (self Simple) GetSetting(name string, default_value gd.Variant) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetSetting(gc, gc.String(name), default_value))
}
func (self Simple) GetSettingWithOverride(name string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetSettingWithOverride(gc, gc.StringName(name)))
}
func (self Simple) GetGlobalClassList() gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](Expert(self).GetGlobalClassList(gc))
}
func (self Simple) SetOrder(name string, position int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOrder(gc.String(name), gd.Int(position))
}
func (self Simple) GetOrder(name string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetOrder(gc.String(name))))
}
func (self Simple) SetInitialValue(name string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInitialValue(gc.String(name), value)
}
func (self Simple) SetAsBasic(name string, basic bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAsBasic(gc.String(name), basic)
}
func (self Simple) SetAsInternal(name string, internal_ bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAsInternal(gc.String(name), internal_)
}
func (self Simple) AddPropertyInfo(hint gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddPropertyInfo(hint)
}
func (self Simple) SetRestartIfChanged(name string, restart bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRestartIfChanged(gc.String(name), restart)
}
func (self Simple) Clear(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear(gc.String(name))
}
func (self Simple) LocalizePath(path string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).LocalizePath(gc, gc.String(path)).String())
}
func (self Simple) GlobalizePath(path string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GlobalizePath(gc, gc.String(path)).String())
}
func (self Simple) Save() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Save())
}
func (self Simple) LoadResourcePack(pack string, replace_files bool, offset int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).LoadResourcePack(gc.String(pack), replace_files, gd.Int(offset)))
}
func (self Simple) SaveCustom(file string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).SaveCustom(gc.String(file)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ProjectSettings
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns [code]true[/code] if a configuration value is present.
*/
//go:nosplit
func (self class) HasSetting(name gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProjectSettings.Bind_has_setting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) SetSetting(name gd.String, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProjectSettings.Bind_set_setting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) GetSetting(ctx gd.Lifetime, name gd.String, default_value gd.Variant) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(default_value))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProjectSettings.Bind_get_setting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
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
func (self class) GetSettingWithOverride(ctx gd.Lifetime, name gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProjectSettings.Bind_get_setting_with_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
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
func (self class) GetGlobalClassList(ctx gd.Lifetime) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProjectSettings.Bind_get_global_class_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Sets the order of a configuration value (influences when saved to the config file).
*/
//go:nosplit
func (self class) SetOrder(name gd.String, position gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProjectSettings.Bind_set_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the order of a configuration value (influences when saved to the config file).
*/
//go:nosplit
func (self class) GetOrder(name gd.String) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProjectSettings.Bind_get_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the specified setting's initial value. This is the value the setting reverts to.
*/
//go:nosplit
func (self class) SetInitialValue(name gd.String, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProjectSettings.Bind_set_initial_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Defines if the specified setting is considered basic or advanced. Basic settings will always be shown in the project settings. Advanced settings will only be shown if the user enables the "Advanced Settings" option.
*/
//go:nosplit
func (self class) SetAsBasic(name gd.String, basic bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, basic)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProjectSettings.Bind_set_as_basic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Defines if the specified setting is considered internal. An internal setting won't show up in the Project Settings dialog. This is mostly useful for addons that need to store their own internal settings without exposing them directly to the user.
*/
//go:nosplit
func (self class) SetAsInternal(name gd.String, internal_ bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, internal_)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProjectSettings.Bind_set_as_internal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AddPropertyInfo(hint gd.Dictionary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(hint))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProjectSettings.Bind_add_property_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets whether a setting requires restarting the editor to properly take effect.
[b]Note:[/b] This is just a hint to display to the user that the editor must be restarted for changes to take effect. Enabling [method set_restart_if_changed] does [i]not[/i] delay the setting being set when changed.
*/
//go:nosplit
func (self class) SetRestartIfChanged(name gd.String, restart bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, restart)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProjectSettings.Bind_set_restart_if_changed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears the whole configuration (not recommended, may break things).
*/
//go:nosplit
func (self class) Clear(name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProjectSettings.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the localized path (starting with [code]res://[/code]) corresponding to the absolute, native OS [param path]. See also [method globalize_path].
*/
//go:nosplit
func (self class) LocalizePath(ctx gd.Lifetime, path gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProjectSettings.Bind_localize_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
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
func (self class) GlobalizePath(ctx gd.Lifetime, path gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProjectSettings.Bind_globalize_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Saves the configuration to the [code]project.godot[/code] file.
[b]Note:[/b] This method is intended to be used by editor plugins, as modified [ProjectSettings] can't be loaded back in the running app. If you want to change project settings in exported projects, use [method save_custom] to save [code]override.cfg[/code] file.
*/
//go:nosplit
func (self class) Save() int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProjectSettings.Bind_save, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(pack))
	callframe.Arg(frame, replace_files)
	callframe.Arg(frame, offset)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProjectSettings.Bind_load_resource_pack, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Saves the configuration to a custom file. The file extension must be [code].godot[/code] (to save in text-based [ConfigFile] format) or [code].binary[/code] (to save in binary format). You can also save [code]override.cfg[/code] file, which is also text, but can be used in exported projects unlike other formats.
*/
//go:nosplit
func (self class) SaveCustom(file gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(file))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProjectSettings.Bind_save_custom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsProjectSettings() Expert { return self[0].AsProjectSettings() }


//go:nosplit
func (self Simple) AsProjectSettings() Simple { return self[0].AsProjectSettings() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("ProjectSettings", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
