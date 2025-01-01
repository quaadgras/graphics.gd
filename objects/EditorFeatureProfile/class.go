package EditorFeatureProfile

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
An editor feature profile can be used to disable specific features of the Godot editor. When disabled, the features won't appear in the editor, which makes the editor less cluttered. This is useful in education settings to reduce confusion or when working in a team. For example, artists and level designers could use a feature profile that disables the script editor to avoid accidentally making changes to files they aren't supposed to edit.
To manage editor feature profiles visually, use [b]Editor > Manage Feature Profiles...[/b] at the top of the editor window.
*/
type Instance [1]classdb.EditorFeatureProfile
type Any interface {
	gd.IsClass
	AsEditorFeatureProfile() Instance
}

/*
If [param disable] is [code]true[/code], disables the class specified by [param class_name]. When disabled, the class won't appear in the Create New Node dialog.
*/
func (self Instance) SetDisableClass(class_name string, disable bool) {
	class(self).SetDisableClass(gd.NewStringName(class_name), disable)
}

/*
Returns [code]true[/code] if the class specified by [param class_name] is disabled. When disabled, the class won't appear in the Create New Node dialog.
*/
func (self Instance) IsClassDisabled(class_name string) bool {
	return bool(class(self).IsClassDisabled(gd.NewStringName(class_name)))
}

/*
If [param disable] is [code]true[/code], disables editing for the class specified by [param class_name]. When disabled, the class will still appear in the Create New Node dialog but the Inspector will be read-only when selecting a node that extends the class.
*/
func (self Instance) SetDisableClassEditor(class_name string, disable bool) {
	class(self).SetDisableClassEditor(gd.NewStringName(class_name), disable)
}

/*
Returns [code]true[/code] if editing for the class specified by [param class_name] is disabled. When disabled, the class will still appear in the Create New Node dialog but the Inspector will be read-only when selecting a node that extends the class.
*/
func (self Instance) IsClassEditorDisabled(class_name string) bool {
	return bool(class(self).IsClassEditorDisabled(gd.NewStringName(class_name)))
}

/*
If [param disable] is [code]true[/code], disables editing for [param property] in the class specified by [param class_name]. When a property is disabled, it won't appear in the Inspector when selecting a node that extends the class specified by [param class_name].
*/
func (self Instance) SetDisableClassProperty(class_name string, property string, disable bool) {
	class(self).SetDisableClassProperty(gd.NewStringName(class_name), gd.NewStringName(property), disable)
}

/*
Returns [code]true[/code] if [param property] is disabled in the class specified by [param class_name]. When a property is disabled, it won't appear in the Inspector when selecting a node that extends the class specified by [param class_name].
*/
func (self Instance) IsClassPropertyDisabled(class_name string, property string) bool {
	return bool(class(self).IsClassPropertyDisabled(gd.NewStringName(class_name), gd.NewStringName(property)))
}

/*
If [param disable] is [code]true[/code], disables the editor feature specified in [param feature]. When a feature is disabled, it will disappear from the editor entirely.
*/
func (self Instance) SetDisableFeature(feature classdb.EditorFeatureProfileFeature, disable bool) {
	class(self).SetDisableFeature(feature, disable)
}

/*
Returns [code]true[/code] if the [param feature] is disabled. When a feature is disabled, it will disappear from the editor entirely.
*/
func (self Instance) IsFeatureDisabled(feature classdb.EditorFeatureProfileFeature) bool {
	return bool(class(self).IsFeatureDisabled(feature))
}

/*
Returns the specified [param feature]'s human-readable name.
*/
func (self Instance) GetFeatureName(feature classdb.EditorFeatureProfileFeature) string {
	return string(class(self).GetFeatureName(feature).String())
}

/*
Saves the editor feature profile to a file in JSON format. It can then be imported using the feature profile manager's [b]Import[/b] button or the [method load_from_file] method.
[b]Note:[/b] Feature profiles created via the user interface are saved in the [code]feature_profiles[/code] directory, as a file with the [code].profile[/code] extension. The editor configuration folder can be found by using [method EditorPaths.get_config_dir].
*/
func (self Instance) SaveToFile(path string) error {
	return error(class(self).SaveToFile(gd.NewString(path)))
}

/*
Loads an editor feature profile from a file. The file must follow the JSON format obtained by using the feature profile manager's [b]Export[/b] button or the [method save_to_file] method.
[b]Note:[/b] Feature profiles created via the user interface are loaded from the [code]feature_profiles[/code] directory, as a file with the [code].profile[/code] extension. The editor configuration folder can be found by using [method EditorPaths.get_config_dir].
*/
func (self Instance) LoadFromFile(path string) error {
	return error(class(self).LoadFromFile(gd.NewString(path)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EditorFeatureProfile

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorFeatureProfile"))
	return Instance{classdb.EditorFeatureProfile(object)}
}

/*
If [param disable] is [code]true[/code], disables the class specified by [param class_name]. When disabled, the class won't appear in the Create New Node dialog.
*/
//go:nosplit
func (self class) SetDisableClass(class_name gd.StringName, disable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_name))
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_set_disable_class, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the class specified by [param class_name] is disabled. When disabled, the class won't appear in the Create New Node dialog.
*/
//go:nosplit
func (self class) IsClassDisabled(class_name gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_is_class_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param disable] is [code]true[/code], disables editing for the class specified by [param class_name]. When disabled, the class will still appear in the Create New Node dialog but the Inspector will be read-only when selecting a node that extends the class.
*/
//go:nosplit
func (self class) SetDisableClassEditor(class_name gd.StringName, disable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_name))
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_set_disable_class_editor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if editing for the class specified by [param class_name] is disabled. When disabled, the class will still appear in the Create New Node dialog but the Inspector will be read-only when selecting a node that extends the class.
*/
//go:nosplit
func (self class) IsClassEditorDisabled(class_name gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_is_class_editor_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param disable] is [code]true[/code], disables editing for [param property] in the class specified by [param class_name]. When a property is disabled, it won't appear in the Inspector when selecting a node that extends the class specified by [param class_name].
*/
//go:nosplit
func (self class) SetDisableClassProperty(class_name gd.StringName, property gd.StringName, disable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_name))
	callframe.Arg(frame, pointers.Get(property))
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_set_disable_class_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if [param property] is disabled in the class specified by [param class_name]. When a property is disabled, it won't appear in the Inspector when selecting a node that extends the class specified by [param class_name].
*/
//go:nosplit
func (self class) IsClassPropertyDisabled(class_name gd.StringName, property gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(class_name))
	callframe.Arg(frame, pointers.Get(property))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_is_class_property_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param disable] is [code]true[/code], disables the editor feature specified in [param feature]. When a feature is disabled, it will disappear from the editor entirely.
*/
//go:nosplit
func (self class) SetDisableFeature(feature classdb.EditorFeatureProfileFeature, disable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_set_disable_feature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the [param feature] is disabled. When a feature is disabled, it will disappear from the editor entirely.
*/
//go:nosplit
func (self class) IsFeatureDisabled(feature classdb.EditorFeatureProfileFeature) bool {
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_is_feature_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the specified [param feature]'s human-readable name.
*/
//go:nosplit
func (self class) GetFeatureName(feature classdb.EditorFeatureProfileFeature) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_get_feature_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Saves the editor feature profile to a file in JSON format. It can then be imported using the feature profile manager's [b]Import[/b] button or the [method load_from_file] method.
[b]Note:[/b] Feature profiles created via the user interface are saved in the [code]feature_profiles[/code] directory, as a file with the [code].profile[/code] extension. The editor configuration folder can be found by using [method EditorPaths.get_config_dir].
*/
//go:nosplit
func (self class) SaveToFile(path gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_save_to_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Loads an editor feature profile from a file. The file must follow the JSON format obtained by using the feature profile manager's [b]Export[/b] button or the [method save_to_file] method.
[b]Note:[/b] Feature profiles created via the user interface are loaded from the [code]feature_profiles[/code] directory, as a file with the [code].profile[/code] extension. The editor configuration folder can be found by using [method EditorPaths.get_config_dir].
*/
//go:nosplit
func (self class) LoadFromFile(path gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_load_from_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsEditorFeatureProfile() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorFeatureProfile() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted         { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted      { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("EditorFeatureProfile", func(ptr gd.Object) any { return [1]classdb.EditorFeatureProfile{classdb.EditorFeatureProfile(ptr)} })
}

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

type Error int

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
