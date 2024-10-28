package FileDialog

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/ConfirmationDialog"
import "grow.graphics/gd/gdclass/AcceptDialog"
import "grow.graphics/gd/gdclass/Window"
import "grow.graphics/gd/gdclass/Viewport"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
[FileDialog] is a preset dialog used to choose files and directories in the filesystem. It supports filter masks. [FileDialog] automatically sets its window title according to the [member file_mode]. If you want to use a custom title, disable this by setting [member mode_overrides_title] to [code]false[/code].

*/
type Go [1]classdb.FileDialog

/*
Clear all the added filters in the dialog.
*/
func (self Go) ClearFilters() {
	class(self).ClearFilters()
}

/*
Adds a comma-delimited file name [param filter] option to the [FileDialog] with an optional [param description], which restricts what files can be picked.
A [param filter] should be of the form [code]"filename.extension"[/code], where filename and extension can be [code]*[/code] to match any string. Filters starting with [code].[/code] (i.e. empty filenames) are not allowed.
For example, a [param filter] of [code]"*.png, *.jpg"[/code] and a [param description] of [code]"Images"[/code] results in filter text "Images (*.png, *.jpg)".
*/
func (self Go) AddFilter(filter string) {
	class(self).AddFilter(gd.NewString(filter), gd.NewString(""))
}

/*
Returns the name of the [OptionButton] or [CheckBox] with index [param option].
*/
func (self Go) GetOptionName(option int) string {
	return string(class(self).GetOptionName(gd.Int(option)).String())
}

/*
Returns an array of values of the [OptionButton] with index [param option].
*/
func (self Go) GetOptionValues(option int) []string {
	return []string(class(self).GetOptionValues(gd.Int(option)).Strings())
}

/*
Returns the default value index of the [OptionButton] or [CheckBox] with index [param option].
*/
func (self Go) GetOptionDefault(option int) int {
	return int(int(class(self).GetOptionDefault(gd.Int(option))))
}

/*
Sets the name of the [OptionButton] or [CheckBox] with index [param option].
*/
func (self Go) SetOptionName(option int, name string) {
	class(self).SetOptionName(gd.Int(option), gd.NewString(name))
}

/*
Sets the option values of the [OptionButton] with index [param option].
*/
func (self Go) SetOptionValues(option int, values []string) {
	class(self).SetOptionValues(gd.Int(option), gd.NewPackedStringSlice(values))
}

/*
Sets the default value index of the [OptionButton] or [CheckBox] with index [param option].
*/
func (self Go) SetOptionDefault(option int, default_value_index int) {
	class(self).SetOptionDefault(gd.Int(option), gd.Int(default_value_index))
}

/*
Adds an additional [OptionButton] to the file dialog. If [param values] is empty, a [CheckBox] is added instead.
[param default_value_index] should be an index of the value in the [param values]. If [param values] is empty it should be either [code]1[/code] (checked), or [code]0[/code] (unchecked).
*/
func (self Go) AddOption(name string, values []string, default_value_index int) {
	class(self).AddOption(gd.NewString(name), gd.NewPackedStringSlice(values), gd.Int(default_value_index))
}

/*
Returns a [Dictionary] with the selected values of the additional [OptionButton]s and/or [CheckBox]es. [Dictionary] keys are names and values are selected value indices.
*/
func (self Go) GetSelectedOptions() gd.Dictionary {
	return gd.Dictionary(class(self).GetSelectedOptions())
}

/*
Returns the vertical box container of the dialog, custom controls can be added to it.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
[b]Note:[/b] Changes to this node are ignored by native file dialogs, use [method add_option] to add custom elements to the dialog instead.
*/
func (self Go) GetVbox() gdclass.VBoxContainer {
	return gdclass.VBoxContainer(class(self).GetVbox())
}

/*
Returns the LineEdit for the selected file.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Go) GetLineEdit() gdclass.LineEdit {
	return gdclass.LineEdit(class(self).GetLineEdit())
}

/*
Clear all currently selected items in the dialog.
*/
func (self Go) DeselectAll() {
	class(self).DeselectAll()
}

/*
Invalidate and update the current dialog content list.
[b]Note:[/b] This method does nothing on native file dialogs.
*/
func (self Go) Invalidate() {
	class(self).Invalidate()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.FileDialog
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FileDialog"))
	return Go{classdb.FileDialog(object)}
}

func (self Go) ModeOverridesTitle() bool {
		return bool(class(self).IsModeOverridingTitle())
}

func (self Go) SetModeOverridesTitle(value bool) {
	class(self).SetModeOverridesTitle(value)
}

func (self Go) FileMode() classdb.FileDialogFileMode {
		return classdb.FileDialogFileMode(class(self).GetFileMode())
}

func (self Go) SetFileMode(value classdb.FileDialogFileMode) {
	class(self).SetFileMode(value)
}

func (self Go) Access() classdb.FileDialogAccess {
		return classdb.FileDialogAccess(class(self).GetAccess())
}

func (self Go) SetAccess(value classdb.FileDialogAccess) {
	class(self).SetAccess(value)
}

func (self Go) RootSubfolder() string {
		return string(class(self).GetRootSubfolder().String())
}

func (self Go) SetRootSubfolder(value string) {
	class(self).SetRootSubfolder(gd.NewString(value))
}

func (self Go) Filters() []string {
		return []string(class(self).GetFilters().Strings())
}

func (self Go) SetFilters(value []string) {
	class(self).SetFilters(gd.NewPackedStringSlice(value))
}

func (self Go) OptionCount() int {
		return int(int(class(self).GetOptionCount()))
}

func (self Go) SetOptionCount(value int) {
	class(self).SetOptionCount(gd.Int(value))
}

func (self Go) ShowHiddenFiles() bool {
		return bool(class(self).IsShowingHiddenFiles())
}

func (self Go) SetShowHiddenFiles(value bool) {
	class(self).SetShowHiddenFiles(value)
}

func (self Go) UseNativeDialog() bool {
		return bool(class(self).GetUseNativeDialog())
}

func (self Go) SetUseNativeDialog(value bool) {
	class(self).SetUseNativeDialog(value)
}

func (self Go) CurrentDir() string {
		return string(class(self).GetCurrentDir().String())
}

func (self Go) SetCurrentDir(value string) {
	class(self).SetCurrentDir(gd.NewString(value))
}

func (self Go) CurrentFile() string {
		return string(class(self).GetCurrentFile().String())
}

func (self Go) SetCurrentFile(value string) {
	class(self).SetCurrentFile(gd.NewString(value))
}

func (self Go) CurrentPath() string {
		return string(class(self).GetCurrentPath().String())
}

func (self Go) SetCurrentPath(value string) {
	class(self).SetCurrentPath(gd.NewString(value))
}

/*
Clear all the added filters in the dialog.
*/
//go:nosplit
func (self class) ClearFilters()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_clear_filters, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a comma-delimited file name [param filter] option to the [FileDialog] with an optional [param description], which restricts what files can be picked.
A [param filter] should be of the form [code]"filename.extension"[/code], where filename and extension can be [code]*[/code] to match any string. Filters starting with [code].[/code] (i.e. empty filenames) are not allowed.
For example, a [param filter] of [code]"*.png, *.jpg"[/code] and a [param description] of [code]"Images"[/code] results in filter text "Images (*.png, *.jpg)".
*/
//go:nosplit
func (self class) AddFilter(filter gd.String, description gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(filter))
	callframe.Arg(frame, discreet.Get(description))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_add_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFilters(filters gd.PackedStringArray)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(filters))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_filters, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFilters() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_filters, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the name of the [OptionButton] or [CheckBox] with index [param option].
*/
//go:nosplit
func (self class) GetOptionName(option gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, option)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_option_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array of values of the [OptionButton] with index [param option].
*/
//go:nosplit
func (self class) GetOptionValues(option gd.Int) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, option)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_option_values, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the default value index of the [OptionButton] or [CheckBox] with index [param option].
*/
//go:nosplit
func (self class) GetOptionDefault(option gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, option)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_option_default, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the name of the [OptionButton] or [CheckBox] with index [param option].
*/
//go:nosplit
func (self class) SetOptionName(option gd.Int, name gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, option)
	callframe.Arg(frame, discreet.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_option_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the option values of the [OptionButton] with index [param option].
*/
//go:nosplit
func (self class) SetOptionValues(option gd.Int, values gd.PackedStringArray)  {
	var frame = callframe.New()
	callframe.Arg(frame, option)
	callframe.Arg(frame, discreet.Get(values))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_option_values, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the default value index of the [OptionButton] or [CheckBox] with index [param option].
*/
//go:nosplit
func (self class) SetOptionDefault(option gd.Int, default_value_index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, option)
	callframe.Arg(frame, default_value_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_option_default, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetOptionCount(count gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_option_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOptionCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_option_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds an additional [OptionButton] to the file dialog. If [param values] is empty, a [CheckBox] is added instead.
[param default_value_index] should be an index of the value in the [param values]. If [param values] is empty it should be either [code]1[/code] (checked), or [code]0[/code] (unchecked).
*/
//go:nosplit
func (self class) AddOption(name gd.String, values gd.PackedStringArray, default_value_index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	callframe.Arg(frame, discreet.Get(values))
	callframe.Arg(frame, default_value_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_add_option, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [Dictionary] with the selected values of the additional [OptionButton]s and/or [CheckBox]es. [Dictionary] keys are names and values are selected value indices.
*/
//go:nosplit
func (self class) GetSelectedOptions() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_selected_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetCurrentDir() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_current_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetCurrentFile() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_current_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetCurrentPath() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_current_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCurrentDir(dir gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(dir))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_current_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCurrentFile(file gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(file))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_current_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCurrentPath(path gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_current_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetModeOverridesTitle(override bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, override)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_mode_overrides_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsModeOverridingTitle() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_is_mode_overriding_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFileMode(mode classdb.FileDialogFileMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_file_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFileMode() classdb.FileDialogFileMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.FileDialogFileMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_file_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) GetVbox() gdclass.VBoxContainer {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_vbox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.VBoxContainer{classdb.VBoxContainer(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Returns the LineEdit for the selected file.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetLineEdit() gdclass.LineEdit {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_line_edit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.LineEdit{classdb.LineEdit(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAccess(access classdb.FileDialogAccess)  {
	var frame = callframe.New()
	callframe.Arg(frame, access)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_access, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAccess() classdb.FileDialogAccess {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.FileDialogAccess](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_access, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRootSubfolder(dir gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(dir))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_root_subfolder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRootSubfolder() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_root_subfolder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowHiddenFiles(show bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, show)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_show_hidden_files, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShowingHiddenFiles() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_is_showing_hidden_files, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseNativeDialog(native bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, native)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_set_use_native_dialog, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUseNativeDialog() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_get_use_native_dialog, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Clear all currently selected items in the dialog.
*/
//go:nosplit
func (self class) DeselectAll()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_deselect_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Invalidate and update the current dialog content list.
[b]Note:[/b] This method does nothing on native file dialogs.
*/
//go:nosplit
func (self class) Invalidate()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileDialog.Bind_invalidate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnFileSelected(cb func(path string)) {
	self[0].AsObject().Connect(gd.NewStringName("file_selected"), gd.NewCallable(cb), 0)
}


func (self Go) OnFilesSelected(cb func(paths []string)) {
	self[0].AsObject().Connect(gd.NewStringName("files_selected"), gd.NewCallable(cb), 0)
}


func (self Go) OnDirSelected(cb func(dir string)) {
	self[0].AsObject().Connect(gd.NewStringName("dir_selected"), gd.NewCallable(cb), 0)
}


func (self class) AsFileDialog() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsFileDialog() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsConfirmationDialog() ConfirmationDialog.GD { return *((*ConfirmationDialog.GD)(unsafe.Pointer(&self))) }
func (self Go) AsConfirmationDialog() ConfirmationDialog.Go { return *((*ConfirmationDialog.Go)(unsafe.Pointer(&self))) }
func (self class) AsAcceptDialog() AcceptDialog.GD { return *((*AcceptDialog.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAcceptDialog() AcceptDialog.Go { return *((*AcceptDialog.Go)(unsafe.Pointer(&self))) }
func (self class) AsWindow() Window.GD { return *((*Window.GD)(unsafe.Pointer(&self))) }
func (self Go) AsWindow() Window.Go { return *((*Window.Go)(unsafe.Pointer(&self))) }
func (self class) AsViewport() Viewport.GD { return *((*Viewport.GD)(unsafe.Pointer(&self))) }
func (self Go) AsViewport() Viewport.Go { return *((*Viewport.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsConfirmationDialog(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsConfirmationDialog(), name)
	}
}
func init() {classdb.Register("FileDialog", func(ptr gd.Object) any { return classdb.FileDialog(ptr) })}
type FileMode = classdb.FileDialogFileMode

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
type Access = classdb.FileDialogAccess

const (
/*The dialog only allows accessing files under the [Resource] path ([code]res://[/code]).*/
	AccessResources Access = 0
/*The dialog only allows accessing files under user data path ([code]user://[/code]).*/
	AccessUserdata Access = 1
/*The dialog allows accessing files on the whole file system.*/
	AccessFilesystem Access = 2
)
