// Package EditorSettings provides methods for working with EditorSettings object instances.
package EditorSettings

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
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any

/*
Object that holds the project-independent editor settings. These settings are generally visible in the [b]Editor > Editor Settings[/b] menu.
Property names use slash delimiters to distinguish sections. Setting values can be of any [Variant] type. It's recommended to use [code]snake_case[/code] for editor settings to be consistent with the Godot editor itself.
Accessing the settings can be done using the following methods, such as:
[codeblocks]
[gdscript]
var settings = EditorInterface.get_editor_settings()
# `settings.set("some/property", 10)` also works as this class overrides `_set()` internally.
settings.set_setting("some/property", 10)
# `settings.get("some/property")` also works as this class overrides `_get()` internally.
settings.get_setting("some/property")
var list_of_settings = settings.get_property_list()
[/gdscript]
[csharp]
EditorSettings settings = EditorInterface.Singleton.GetEditorSettings();
// `settings.set("some/property", value)` also works as this class overrides `_set()` internally.
settings.SetSetting("some/property", Value);
// `settings.get("some/property", value)` also works as this class overrides `_get()` internally.
settings.GetSetting("some/property");
Godot.Collections.Array<Godot.Collections.Dictionary> listOfSettings = settings.GetPropertyList();
[/csharp]
[/codeblocks]
[b]Note:[/b] This class shouldn't be instantiated directly. Instead, access the singleton using [method EditorInterface.get_editor_settings].
*/
type Instance [1]gdclass.EditorSettings

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorSettings() Instance
}

/*
Returns [code]true[/code] if the setting specified by [param name] exists, [code]false[/code] otherwise.
*/
func (self Instance) HasSetting(name string) bool {
	return bool(class(self).HasSetting(gd.NewString(name)))
}

/*
Sets the [param value] of the setting specified by [param name]. This is equivalent to using [method Object.set] on the EditorSettings instance.
*/
func (self Instance) SetSetting(name string, value any) {
	class(self).SetSetting(gd.NewString(name), gd.NewVariant(value))
}

/*
Returns the value of the setting specified by [param name]. This is equivalent to using [method Object.get] on the EditorSettings instance.
*/
func (self Instance) GetSetting(name string) any {
	return any(class(self).GetSetting(gd.NewString(name)).Interface())
}

/*
Erases the setting whose name is specified by [param property].
*/
func (self Instance) Erase(property string) {
	class(self).Erase(gd.NewString(property))
}

/*
Sets the initial value of the setting specified by [param name] to [param value]. This is used to provide a value for the Revert button in the Editor Settings. If [param update_current] is true, the current value of the setting will be set to [param value] as well.
*/
func (self Instance) SetInitialValue(name string, value any, update_current bool) {
	class(self).SetInitialValue(gd.NewStringName(name), gd.NewVariant(value), update_current)
}

/*
Adds a custom property info to a property. The dictionary must contain:
- [code]name[/code]: [String] (the name of the property)
- [code]type[/code]: [int] (see [enum Variant.Type])
- optionally [code]hint[/code]: [int] (see [enum PropertyHint]) and [code]hint_string[/code]: [String]
[b]Example:[/b]
[codeblocks]
[gdscript]
var settings = EditorInterface.get_editor_settings()
settings.set("category/property_name", 0)

	var property_info = {
	    "name": "category/property_name",
	    "type": TYPE_INT,
	    "hint": PROPERTY_HINT_ENUM,
	    "hint_string": "one,two,three"
	}

settings.add_property_info(property_info)
[/gdscript]
[csharp]
var settings = GetEditorInterface().GetEditorSettings();
settings.Set("category/property_name", 0);

var propertyInfo = new Godot.Collections.Dictionary

	{
	    {"name", "category/propertyName"},
	    {"type", Variant.Type.Int},
	    {"hint", PropertyHint.Enum},
	    {"hint_string", "one,two,three"}
	};

settings.AddPropertyInfo(propertyInfo);
[/csharp]
[/codeblocks]
*/
func (self Instance) AddPropertyInfo(info map[any]any) {
	class(self).AddPropertyInfo(gd.NewVariant(info).Interface().(gd.Dictionary))
}

/*
Sets project-specific metadata with the [param section], [param key] and [param data] specified. This metadata is stored outside the project folder and therefore won't be checked into version control. See also [method get_project_metadata].
*/
func (self Instance) SetProjectMetadata(section string, key string, data any) {
	class(self).SetProjectMetadata(gd.NewString(section), gd.NewString(key), gd.NewVariant(data))
}

/*
Returns project-specific metadata for the [param section] and [param key] specified. If the metadata doesn't exist, [param default] will be returned instead. See also [method set_project_metadata].
*/
func (self Instance) GetProjectMetadata(section string, key string) any {
	return any(class(self).GetProjectMetadata(gd.NewString(section), gd.NewString(key), gd.NewVariant(gd.NewVariant(([1]any{}[0])))).Interface())
}

/*
Sets the list of favorite files and directories for this project.
*/
func (self Instance) SetFavorites(dirs []string) {
	class(self).SetFavorites(gd.NewPackedStringSlice(dirs))
}

/*
Returns the list of favorite files and directories for this project.
*/
func (self Instance) GetFavorites() []string {
	return []string(class(self).GetFavorites().Strings())
}

/*
Sets the list of recently visited folders in the file dialog for this project.
*/
func (self Instance) SetRecentDirs(dirs []string) {
	class(self).SetRecentDirs(gd.NewPackedStringSlice(dirs))
}

/*
Returns the list of recently visited folders in the file dialog for this project.
*/
func (self Instance) GetRecentDirs() []string {
	return []string(class(self).GetRecentDirs().Strings())
}

/*
Overrides the built-in editor action [param name] with the input actions defined in [param actions_list].
*/
func (self Instance) SetBuiltinActionOverride(name string, actions_list [][1]gdclass.InputEvent) {
	class(self).SetBuiltinActionOverride(gd.NewString(name), gd.ArrayFromSlice[Array.Contains[[1]gdclass.InputEvent]](actions_list))
}

/*
Checks if any settings with the prefix [param setting_prefix] exist in the set of changed settings. See also [method get_changed_settings].
*/
func (self Instance) CheckChangedSettingsInGroup(setting_prefix string) bool {
	return bool(class(self).CheckChangedSettingsInGroup(gd.NewString(setting_prefix)))
}

/*
Gets an array of the settings which have been changed since the last save. Note that internally [code]changed_settings[/code] is cleared after a successful save, so generally the most appropriate place to use this method is when processing [constant NOTIFICATION_EDITOR_SETTINGS_CHANGED].
*/
func (self Instance) GetChangedSettings() []string {
	return []string(class(self).GetChangedSettings().Strings())
}

/*
Marks the passed editor setting as being changed, see [method get_changed_settings]. Only settings which exist (see [method has_setting]) will be accepted.
*/
func (self Instance) MarkSettingChanged(setting string) {
	class(self).MarkSettingChanged(gd.NewString(setting))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorSettings

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorSettings"))
	casted := Instance{*(*gdclass.EditorSettings)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Returns [code]true[/code] if the setting specified by [param name] exists, [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) HasSetting(name gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSettings.Bind_has_setting, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [param value] of the setting specified by [param name]. This is equivalent to using [method Object.set] on the EditorSettings instance.
*/
//go:nosplit
func (self class) SetSetting(name gd.String, value gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(value))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSettings.Bind_set_setting, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the value of the setting specified by [param name]. This is equivalent to using [method Object.get] on the EditorSettings instance.
*/
//go:nosplit
func (self class) GetSetting(name gd.String) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSettings.Bind_get_setting, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Erases the setting whose name is specified by [param property].
*/
//go:nosplit
func (self class) Erase(property gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(property))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSettings.Bind_erase, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the initial value of the setting specified by [param name] to [param value]. This is used to provide a value for the Revert button in the Editor Settings. If [param update_current] is true, the current value of the setting will be set to [param value] as well.
*/
//go:nosplit
func (self class) SetInitialValue(name gd.StringName, value gd.Variant, update_current bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(value))
	callframe.Arg(frame, update_current)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSettings.Bind_set_initial_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a custom property info to a property. The dictionary must contain:
- [code]name[/code]: [String] (the name of the property)
- [code]type[/code]: [int] (see [enum Variant.Type])
- optionally [code]hint[/code]: [int] (see [enum PropertyHint]) and [code]hint_string[/code]: [String]
[b]Example:[/b]
[codeblocks]
[gdscript]
var settings = EditorInterface.get_editor_settings()
settings.set("category/property_name", 0)

var property_info = {
    "name": "category/property_name",
    "type": TYPE_INT,
    "hint": PROPERTY_HINT_ENUM,
    "hint_string": "one,two,three"
}

settings.add_property_info(property_info)
[/gdscript]
[csharp]
var settings = GetEditorInterface().GetEditorSettings();
settings.Set("category/property_name", 0);

var propertyInfo = new Godot.Collections.Dictionary
{
    {"name", "category/propertyName"},
    {"type", Variant.Type.Int},
    {"hint", PropertyHint.Enum},
    {"hint_string", "one,two,three"}
};

settings.AddPropertyInfo(propertyInfo);
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) AddPropertyInfo(info gd.Dictionary) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(info))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSettings.Bind_add_property_info, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets project-specific metadata with the [param section], [param key] and [param data] specified. This metadata is stored outside the project folder and therefore won't be checked into version control. See also [method get_project_metadata].
*/
//go:nosplit
func (self class) SetProjectMetadata(section gd.String, key gd.String, data gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(section))
	callframe.Arg(frame, pointers.Get(key))
	callframe.Arg(frame, pointers.Get(data))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSettings.Bind_set_project_metadata, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns project-specific metadata for the [param section] and [param key] specified. If the metadata doesn't exist, [param default] will be returned instead. See also [method set_project_metadata].
*/
//go:nosplit
func (self class) GetProjectMetadata(section gd.String, key gd.String, def gd.Variant) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(section))
	callframe.Arg(frame, pointers.Get(key))
	callframe.Arg(frame, pointers.Get(def))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSettings.Bind_get_project_metadata, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the list of favorite files and directories for this project.
*/
//go:nosplit
func (self class) SetFavorites(dirs gd.PackedStringArray) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(dirs))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSettings.Bind_set_favorites, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the list of favorite files and directories for this project.
*/
//go:nosplit
func (self class) GetFavorites() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSettings.Bind_get_favorites, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the list of recently visited folders in the file dialog for this project.
*/
//go:nosplit
func (self class) SetRecentDirs(dirs gd.PackedStringArray) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(dirs))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSettings.Bind_set_recent_dirs, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the list of recently visited folders in the file dialog for this project.
*/
//go:nosplit
func (self class) GetRecentDirs() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSettings.Bind_get_recent_dirs, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Overrides the built-in editor action [param name] with the input actions defined in [param actions_list].
*/
//go:nosplit
func (self class) SetBuiltinActionOverride(name gd.String, actions_list Array.Contains[[1]gdclass.InputEvent]) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(gd.InternalArray(actions_list)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSettings.Bind_set_builtin_action_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Checks if any settings with the prefix [param setting_prefix] exist in the set of changed settings. See also [method get_changed_settings].
*/
//go:nosplit
func (self class) CheckChangedSettingsInGroup(setting_prefix gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(setting_prefix))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSettings.Bind_check_changed_settings_in_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets an array of the settings which have been changed since the last save. Note that internally [code]changed_settings[/code] is cleared after a successful save, so generally the most appropriate place to use this method is when processing [constant NOTIFICATION_EDITOR_SETTINGS_CHANGED].
*/
//go:nosplit
func (self class) GetChangedSettings() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSettings.Bind_get_changed_settings, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Marks the passed editor setting as being changed, see [method get_changed_settings]. Only settings which exist (see [method has_setting]) will be accepted.
*/
//go:nosplit
func (self class) MarkSettingChanged(setting gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(setting))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSettings.Bind_mark_setting_changed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnSettingsChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("settings_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsEditorSettings() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorSettings() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("EditorSettings", func(ptr gd.Object) any {
		return [1]gdclass.EditorSettings{*(*gdclass.EditorSettings)(unsafe.Pointer(&ptr))}
	})
}
