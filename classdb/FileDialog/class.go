// Package FileDialog provides methods for working with FileDialog object instances.
package FileDialog

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
import "graphics.gd/classdb/ConfirmationDialog"
import "graphics.gd/classdb/AcceptDialog"
import "graphics.gd/classdb/Window"
import "graphics.gd/classdb/Viewport"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
[FileDialog] is a preset dialog used to choose files and directories in the filesystem. It supports filter masks. [FileDialog] automatically sets its window title according to the [member file_mode]. If you want to use a custom title, disable this by setting [member mode_overrides_title] to [code]false[/code].
*/
type Instance [1]gdclass.FileDialog

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsFileDialog() Instance
}

/*
Clear all the added filters in the dialog.
*/
func (self Instance) ClearFilters() { //gd:FileDialog.clear_filters
	class(self).ClearFilters()
}

/*
Adds a comma-delimited file name [param filter] option to the [FileDialog] with an optional [param description], which restricts what files can be picked.
A [param filter] should be of the form [code]"filename.extension"[/code], where filename and extension can be [code]*[/code] to match any string. Filters starting with [code].[/code] (i.e. empty filenames) are not allowed.
For example, a [param filter] of [code]"*.png, *.jpg"[/code] and a [param description] of [code]"Images"[/code] results in filter text "Images (*.png, *.jpg)".
*/
func (self Instance) AddFilter(filter string) { //gd:FileDialog.add_filter
	class(self).AddFilter(gd.NewString(filter), gd.NewString(""))
}

/*
Returns the name of the [OptionButton] or [CheckBox] with index [param option].
*/
func (self Instance) GetOptionName(option int) string { //gd:FileDialog.get_option_name
	return string(class(self).GetOptionName(gd.Int(option)).String())
}

/*
Returns an array of values of the [OptionButton] with index [param option].
*/
func (self Instance) GetOptionValues(option int) []string { //gd:FileDialog.get_option_values
	return []string(class(self).GetOptionValues(gd.Int(option)).Strings())
}

/*
Returns the default value index of the [OptionButton] or [CheckBox] with index [param option].
*/
func (self Instance) GetOptionDefault(option int) int { //gd:FileDialog.get_option_default
	return int(int(class(self).GetOptionDefault(gd.Int(option))))
}

/*
Sets the name of the [OptionButton] or [CheckBox] with index [param option].
*/
func (self Instance) SetOptionName(option int, name string) { //gd:FileDialog.set_option_name
	class(self).SetOptionName(gd.Int(option), gd.NewString(name))
}

/*
Sets the option values of the [OptionButton] with index [param option].
*/
func (self Instance) SetOptionValues(option int, values []string) { //gd:FileDialog.set_option_values
	class(self).SetOptionValues(gd.Int(option), gd.NewPackedStringSlice(values))
}

/*
Sets the default value index of the [OptionButton] or [CheckBox] with index [param option].
*/
func (self Instance) SetOptionDefault(option int, default_value_index int) { //gd:FileDialog.set_option_default
	class(self).SetOptionDefault(gd.Int(option), gd.Int(default_value_index))
}

/*
Adds an additional [OptionButton] to the file dialog. If [param values] is empty, a [CheckBox] is added instead.
[param default_value_index] should be an index of the value in the [param values]. If [param values] is empty it should be either [code]1[/code] (checked), or [code]0[/code] (unchecked).
*/
func (self Instance) AddOption(name string, values []string, default_value_index int) { //gd:FileDialog.add_option
	class(self).AddOption(gd.NewString(name), gd.NewPackedStringSlice(values), gd.Int(default_value_index))
}

/*
Returns a [Dictionary] with the selected values of the additional [OptionButton]s and/or [CheckBox]es. [Dictionary] keys are names and values are selected value indices.
*/
func (self Instance) GetSelectedOptions() map[any]any { //gd:FileDialog.get_selected_options
	return map[any]any(gd.DictionaryAs[any, any](class(self).GetSelectedOptions()))
}

/*
Returns the vertical box container of the dialog, custom controls can be added to it.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
[b]Note:[/b] Changes to this node are ignored by native file dialogs, use [method add_option] to add custom elements to the dialog instead.
*/
func (self Instance) GetVbox() [1]gdclass.VBoxContainer { //gd:FileDialog.get_vbox
	return [1]gdclass.VBoxContainer(class(self).GetVbox())
}

/*
Returns the LineEdit for the selected file.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Instance) GetLineEdit() [1]gdclass.LineEdit { //gd:FileDialog.get_line_edit
	return [1]gdclass.LineEdit(class(self).GetLineEdit())
}

/*
Clear all currently selected items in the dialog.
*/
func (self Instance) DeselectAll() { //gd:FileDialog.deselect_all
	class(self).DeselectAll()
}

/*
Invalidate and update the current dialog content list.
[b]Note:[/b] This method does nothing on native file dialogs.
*/
func (self Instance) Invalidate() { //gd:FileDialog.invalidate
	class(self).Invalidate()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.FileDialog

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FileDialog"))
	casted := Instance{*(*gdclass.FileDialog)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) ModeOverridesTitle() bool {
	return bool(class(self).IsModeOverridingTitle())
}

func (self Instance) SetModeOverridesTitle(value bool) {
	class(self).SetModeOverridesTitle(value)
}

func (self Instance) FileMode() gdclass.FileDialogFileMode {
	return gdclass.FileDialogFileMode(class(self).GetFileMode())
}

func (self Instance) SetFileMode(value gdclass.FileDialogFileMode) {
	class(self).SetFileMode(value)
}

func (self Instance) Access() gdclass.FileDialogAccess {
	return gdclass.FileDialogAccess(class(self).GetAccess())
}

func (self Instance) SetAccess(value gdclass.FileDialogAccess) {
	class(self).SetAccess(value)
}

func (self Instance) RootSubfolder() string {
	return string(class(self).GetRootSubfolder().String())
}

func (self Instance) SetRootSubfolder(value string) {
	class(self).SetRootSubfolder(gd.NewString(value))
}

func (self Instance) Filters() []string {
	return []string(class(self).GetFilters().Strings())
}

func (self Instance) SetFilters(value []string) {
	class(self).SetFilters(gd.NewPackedStringSlice(value))
}

func (self Instance) OptionCount() int {
	return int(int(class(self).GetOptionCount()))
}

func (self Instance) SetOptionCount(value int) {
	class(self).SetOptionCount(gd.Int(value))
}

func (self Instance) ShowHiddenFiles() bool {
	return bool(class(self).IsShowingHiddenFiles())
}

func (self Instance) SetShowHiddenFiles(value bool) {
	class(self).SetShowHiddenFiles(value)
}

func (self Instance) UseNativeDialog() bool {
	return bool(class(self).GetUseNativeDialog())
}

func (self Instance) SetUseNativeDialog(value bool) {
	class(self).SetUseNativeDialog(value)
}

func (self Instance) CurrentDir() string {
	return string(class(self).GetCurrentDir().String())
}

func (self Instance) SetCurrentDir(value string) {
	class(self).SetCurrentDir(gd.NewString(value))
}

func (self Instance) CurrentFile() string {
	return string(class(self).GetCurrentFile().String())
}

func (self Instance) SetCurrentFile(value string) {
	class(self).SetCurrentFile(gd.NewString(value))
}

func (self Instance) CurrentPath() string {
	return string(class(self).GetCurrentPath().String())
}

func (self Instance) SetCurrentPath(value string) {
	class(self).SetCurrentPath(gd.NewString(value))
}

/*
Clear all the added filters in the dialog.
*/
//go:nosplit
func (self class) ClearFilters() { //gd:FileDialog.clear_filters
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_clear_filters, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a comma-delimited file name [param filter] option to the [FileDialog] with an optional [param description], which restricts what files can be picked.
A [param filter] should be of the form [code]"filename.extension"[/code], where filename and extension can be [code]*[/code] to match any string. Filters starting with [code].[/code] (i.e. empty filenames) are not allowed.
For example, a [param filter] of [code]"*.png, *.jpg"[/code] and a [param description] of [code]"Images"[/code] results in filter text "Images (*.png, *.jpg)".
*/
//go:nosplit
func (self class) AddFilter(filter gd.String, description gd.String) { //gd:FileDialog.add_filter
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(filter))
	callframe.Arg(frame, pointers.Get(description))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_add_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFilters(filters gd.PackedStringArray) { //gd:FileDialog.set_filters
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(filters))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_filters, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFilters() gd.PackedStringArray { //gd:FileDialog.get_filters
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_filters, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the name of the [OptionButton] or [CheckBox] with index [param option].
*/
//go:nosplit
func (self class) GetOptionName(option gd.Int) gd.String { //gd:FileDialog.get_option_name
	var frame = callframe.New()
	callframe.Arg(frame, option)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_option_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an array of values of the [OptionButton] with index [param option].
*/
//go:nosplit
func (self class) GetOptionValues(option gd.Int) gd.PackedStringArray { //gd:FileDialog.get_option_values
	var frame = callframe.New()
	callframe.Arg(frame, option)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_option_values, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the default value index of the [OptionButton] or [CheckBox] with index [param option].
*/
//go:nosplit
func (self class) GetOptionDefault(option gd.Int) gd.Int { //gd:FileDialog.get_option_default
	var frame = callframe.New()
	callframe.Arg(frame, option)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_option_default, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the name of the [OptionButton] or [CheckBox] with index [param option].
*/
//go:nosplit
func (self class) SetOptionName(option gd.Int, name gd.String) { //gd:FileDialog.set_option_name
	var frame = callframe.New()
	callframe.Arg(frame, option)
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_option_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the option values of the [OptionButton] with index [param option].
*/
//go:nosplit
func (self class) SetOptionValues(option gd.Int, values gd.PackedStringArray) { //gd:FileDialog.set_option_values
	var frame = callframe.New()
	callframe.Arg(frame, option)
	callframe.Arg(frame, pointers.Get(values))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_option_values, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the default value index of the [OptionButton] or [CheckBox] with index [param option].
*/
//go:nosplit
func (self class) SetOptionDefault(option gd.Int, default_value_index gd.Int) { //gd:FileDialog.set_option_default
	var frame = callframe.New()
	callframe.Arg(frame, option)
	callframe.Arg(frame, default_value_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_option_default, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetOptionCount(count gd.Int) { //gd:FileDialog.set_option_count
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_option_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOptionCount() gd.Int { //gd:FileDialog.get_option_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_option_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds an additional [OptionButton] to the file dialog. If [param values] is empty, a [CheckBox] is added instead.
[param default_value_index] should be an index of the value in the [param values]. If [param values] is empty it should be either [code]1[/code] (checked), or [code]0[/code] (unchecked).
*/
//go:nosplit
func (self class) AddOption(name gd.String, values gd.PackedStringArray, default_value_index gd.Int) { //gd:FileDialog.add_option
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(values))
	callframe.Arg(frame, default_value_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_add_option, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a [Dictionary] with the selected values of the additional [OptionButton]s and/or [CheckBox]es. [Dictionary] keys are names and values are selected value indices.
*/
//go:nosplit
func (self class) GetSelectedOptions() gd.Dictionary { //gd:FileDialog.get_selected_options
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_selected_options, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetCurrentDir() gd.String { //gd:FileDialog.get_current_dir
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_current_dir, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetCurrentFile() gd.String { //gd:FileDialog.get_current_file
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_current_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetCurrentPath() gd.String { //gd:FileDialog.get_current_path
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_current_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCurrentDir(dir gd.String) { //gd:FileDialog.set_current_dir
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(dir))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_current_dir, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetCurrentFile(file gd.String) { //gd:FileDialog.set_current_file
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(file))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_current_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetCurrentPath(path gd.String) { //gd:FileDialog.set_current_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_current_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetModeOverridesTitle(override bool) { //gd:FileDialog.set_mode_overrides_title
	var frame = callframe.New()
	callframe.Arg(frame, override)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_mode_overrides_title, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsModeOverridingTitle() bool { //gd:FileDialog.is_mode_overriding_title
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_is_mode_overriding_title, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFileMode(mode gdclass.FileDialogFileMode) { //gd:FileDialog.set_file_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_file_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFileMode() gdclass.FileDialogFileMode { //gd:FileDialog.get_file_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.FileDialogFileMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_file_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the vertical box container of the dialog, custom controls can be added to it.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
[b]Note:[/b] Changes to this node are ignored by native file dialogs, use [method add_option] to add custom elements to the dialog instead.
*/
//go:nosplit
func (self class) GetVbox() [1]gdclass.VBoxContainer { //gd:FileDialog.get_vbox
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_vbox, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.VBoxContainer{gd.PointerLifetimeBoundTo[gdclass.VBoxContainer](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the LineEdit for the selected file.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetLineEdit() [1]gdclass.LineEdit { //gd:FileDialog.get_line_edit
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_line_edit, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.LineEdit{gd.PointerLifetimeBoundTo[gdclass.LineEdit](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAccess(access gdclass.FileDialogAccess) { //gd:FileDialog.set_access
	var frame = callframe.New()
	callframe.Arg(frame, access)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_access, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAccess() gdclass.FileDialogAccess { //gd:FileDialog.get_access
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.FileDialogAccess](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_access, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRootSubfolder(dir gd.String) { //gd:FileDialog.set_root_subfolder
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(dir))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_root_subfolder, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRootSubfolder() gd.String { //gd:FileDialog.get_root_subfolder
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_root_subfolder, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShowHiddenFiles(show bool) { //gd:FileDialog.set_show_hidden_files
	var frame = callframe.New()
	callframe.Arg(frame, show)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_show_hidden_files, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShowingHiddenFiles() bool { //gd:FileDialog.is_showing_hidden_files
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_is_showing_hidden_files, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseNativeDialog(native bool) { //gd:FileDialog.set_use_native_dialog
	var frame = callframe.New()
	callframe.Arg(frame, native)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_use_native_dialog, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUseNativeDialog() bool { //gd:FileDialog.get_use_native_dialog
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_use_native_dialog, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clear all currently selected items in the dialog.
*/
//go:nosplit
func (self class) DeselectAll() { //gd:FileDialog.deselect_all
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_deselect_all, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Invalidate and update the current dialog content list.
[b]Note:[/b] This method does nothing on native file dialogs.
*/
//go:nosplit
func (self class) Invalidate() { //gd:FileDialog.invalidate
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_invalidate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnFileSelected(cb func(path string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("file_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnFilesSelected(cb func(paths []string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("files_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnDirSelected(cb func(dir string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("dir_selected"), gd.NewCallable(cb), 0)
}

func (self class) AsFileDialog() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFileDialog() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsConfirmationDialog() ConfirmationDialog.Advanced {
	return *((*ConfirmationDialog.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsConfirmationDialog() ConfirmationDialog.Instance {
	return *((*ConfirmationDialog.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAcceptDialog() AcceptDialog.Advanced {
	return *((*AcceptDialog.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAcceptDialog() AcceptDialog.Instance {
	return *((*AcceptDialog.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsWindow() Window.Advanced    { return *((*Window.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWindow() Window.Instance { return *((*Window.Instance)(unsafe.Pointer(&self))) }
func (self class) AsViewport() Viewport.Advanced {
	return *((*Viewport.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsViewport() Viewport.Instance {
	return *((*Viewport.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(ConfirmationDialog.Advanced(self.AsConfirmationDialog()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(ConfirmationDialog.Instance(self.AsConfirmationDialog()), name)
	}
}
func init() {
	gdclass.Register("FileDialog", func(ptr gd.Object) any { return [1]gdclass.FileDialog{*(*gdclass.FileDialog)(unsafe.Pointer(&ptr))} })
}

type FileMode = gdclass.FileDialogFileMode //gd:FileDialog.FileMode

const (
	/*The dialog allows selecting one, and only one file.*/
	FileModeOpenFile FileMode = 0
	/*The dialog allows selecting multiple files.*/
	FileModeOpenFiles FileMode = 1
	/*The dialog only allows selecting a directory, disallowing the selection of any file.*/
	FileModeOpenDir FileMode = 2
	/*The dialog allows selecting one file or directory.*/
	FileModeOpenAny FileMode = 3
	/*The dialog will warn when a file exists.*/
	FileModeSaveFile FileMode = 4
)

type Access = gdclass.FileDialogAccess //gd:FileDialog.Access

const (
	/*The dialog only allows accessing files under the [Resource] path ([code]res://[/code]).*/
	AccessResources Access = 0
	/*The dialog only allows accessing files under user data path ([code]user://[/code]).*/
	AccessUserdata Access = 1
	/*The dialog allows accessing files on the whole file system.*/
	AccessFilesystem Access = 2
)
