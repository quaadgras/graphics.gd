package EditorFeatureProfile

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
An editor feature profile can be used to disable specific features of the Godot editor. When disabled, the features won't appear in the editor, which makes the editor less cluttered. This is useful in education settings to reduce confusion or when working in a team. For example, artists and level designers could use a feature profile that disables the script editor to avoid accidentally making changes to files they aren't supposed to edit.
To manage editor feature profiles visually, use [b]Editor > Manage Feature Profiles...[/b] at the top of the editor window.

*/
type Simple [1]classdb.EditorFeatureProfile
func (self Simple) SetDisableClass(class_name string, disable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDisableClass(gc.StringName(class_name), disable)
}
func (self Simple) IsClassDisabled(class_name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsClassDisabled(gc.StringName(class_name)))
}
func (self Simple) SetDisableClassEditor(class_name string, disable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDisableClassEditor(gc.StringName(class_name), disable)
}
func (self Simple) IsClassEditorDisabled(class_name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsClassEditorDisabled(gc.StringName(class_name)))
}
func (self Simple) SetDisableClassProperty(class_name string, property string, disable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDisableClassProperty(gc.StringName(class_name), gc.StringName(property), disable)
}
func (self Simple) IsClassPropertyDisabled(class_name string, property string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsClassPropertyDisabled(gc.StringName(class_name), gc.StringName(property)))
}
func (self Simple) SetDisableFeature(feature classdb.EditorFeatureProfileFeature, disable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDisableFeature(feature, disable)
}
func (self Simple) IsFeatureDisabled(feature classdb.EditorFeatureProfileFeature) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsFeatureDisabled(feature))
}
func (self Simple) GetFeatureName(feature classdb.EditorFeatureProfileFeature) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetFeatureName(gc, feature).String())
}
func (self Simple) SaveToFile(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).SaveToFile(gc.String(path)))
}
func (self Simple) LoadFromFile(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).LoadFromFile(gc.String(path)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorFeatureProfile
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
If [param disable] is [code]true[/code], disables the class specified by [param class_name]. When disabled, the class won't appear in the Create New Node dialog.
*/
//go:nosplit
func (self class) SetDisableClass(class_name gd.StringName, disable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_name))
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFeatureProfile.Bind_set_disable_class, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the class specified by [param class_name] is disabled. When disabled, the class won't appear in the Create New Node dialog.
*/
//go:nosplit
func (self class) IsClassDisabled(class_name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFeatureProfile.Bind_is_class_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param disable] is [code]true[/code], disables editing for the class specified by [param class_name]. When disabled, the class will still appear in the Create New Node dialog but the Inspector will be read-only when selecting a node that extends the class.
*/
//go:nosplit
func (self class) SetDisableClassEditor(class_name gd.StringName, disable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_name))
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFeatureProfile.Bind_set_disable_class_editor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if editing for the class specified by [param class_name] is disabled. When disabled, the class will still appear in the Create New Node dialog but the Inspector will be read-only when selecting a node that extends the class.
*/
//go:nosplit
func (self class) IsClassEditorDisabled(class_name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFeatureProfile.Bind_is_class_editor_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param disable] is [code]true[/code], disables editing for [param property] in the class specified by [param class_name]. When a property is disabled, it won't appear in the Inspector when selecting a node that extends the class specified by [param class_name].
*/
//go:nosplit
func (self class) SetDisableClassProperty(class_name gd.StringName, property gd.StringName, disable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_name))
	callframe.Arg(frame, mmm.Get(property))
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFeatureProfile.Bind_set_disable_class_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if [param property] is disabled in the class specified by [param class_name]. When a property is disabled, it won't appear in the Inspector when selecting a node that extends the class specified by [param class_name].
*/
//go:nosplit
func (self class) IsClassPropertyDisabled(class_name gd.StringName, property gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(class_name))
	callframe.Arg(frame, mmm.Get(property))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFeatureProfile.Bind_is_class_property_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param disable] is [code]true[/code], disables the editor feature specified in [param feature]. When a feature is disabled, it will disappear from the editor entirely.
*/
//go:nosplit
func (self class) SetDisableFeature(feature classdb.EditorFeatureProfileFeature, disable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFeatureProfile.Bind_set_disable_feature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the [param feature] is disabled. When a feature is disabled, it will disappear from the editor entirely.
*/
//go:nosplit
func (self class) IsFeatureDisabled(feature classdb.EditorFeatureProfileFeature) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFeatureProfile.Bind_is_feature_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the specified [param feature]'s human-readable name.
*/
//go:nosplit
func (self class) GetFeatureName(ctx gd.Lifetime, feature classdb.EditorFeatureProfileFeature) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFeatureProfile.Bind_get_feature_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Saves the editor feature profile to a file in JSON format. It can then be imported using the feature profile manager's [b]Import[/b] button or the [method load_from_file] method.
[b]Note:[/b] Feature profiles created via the user interface are saved in the [code]feature_profiles[/code] directory, as a file with the [code].profile[/code] extension. The editor configuration folder can be found by using [method EditorPaths.get_config_dir].
*/
//go:nosplit
func (self class) SaveToFile(path gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFeatureProfile.Bind_save_to_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Loads an editor feature profile from a file. The file must follow the JSON format obtained by using the feature profile manager's [b]Export[/b] button or the [method save_to_file] method.
[b]Note:[/b] Feature profiles created via the user interface are loaded from the [code]feature_profiles[/code] directory, as a file with the [code].profile[/code] extension. The editor configuration folder can be found by using [method EditorPaths.get_config_dir].
*/
//go:nosplit
func (self class) LoadFromFile(path gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFeatureProfile.Bind_load_from_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsEditorFeatureProfile() Expert { return self[0].AsEditorFeatureProfile() }


//go:nosplit
func (self Simple) AsEditorFeatureProfile() Simple { return self[0].AsEditorFeatureProfile() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("EditorFeatureProfile", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Feature = classdb.EditorFeatureProfileFeature

const (
/*The 3D editor. If this feature is disabled, the 3D editor won't display but 3D nodes will still display in the Create New Node dialog.*/
	Feature3d Feature = 0
/*The Script tab, which contains the script editor and class reference browser. If this feature is disabled, the Script tab won't display.*/
	FeatureScript Feature = 1
/*The AssetLib tab. If this feature is disabled, the AssetLib tab won't display.*/
	FeatureAssetLib Feature = 2
/*Scene tree editing. If this feature is disabled, the Scene tree dock will still be visible but will be read-only.*/
	FeatureSceneTree Feature = 3
/*The Node dock. If this feature is disabled, signals and groups won't be visible and modifiable from the editor.*/
	FeatureNodeDock Feature = 4
/*The FileSystem dock. If this feature is disabled, the FileSystem dock won't be visible.*/
	FeatureFilesystemDock Feature = 5
/*The Import dock. If this feature is disabled, the Import dock won't be visible.*/
	FeatureImportDock Feature = 6
/*The History dock. If this feature is disabled, the History dock won't be visible.*/
	FeatureHistoryDock Feature = 7
/*Represents the size of the [enum Feature] enum.*/
	FeatureMax Feature = 8
)
