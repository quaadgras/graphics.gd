// Package EditorFileDialog provides methods for working with EditorFileDialog object instances.
package EditorFileDialog

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
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
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
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode

/*
[EditorFileDialog] is an enhanced version of [FileDialog] available only to editor plugins. Additional features include list of favorited/recent files and the ability to see files as thumbnails grid instead of list.
*/
type Instance [1]gdclass.EditorFileDialog

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorFileDialog() Instance
}

/*
Removes all filters except for "All Files (*)".
*/
func (self Instance) ClearFilters() { //gd:EditorFileDialog.clear_filters
	class(self).ClearFilters()
}

/*
Adds a comma-delimited file name [param filter] option to the [EditorFileDialog] with an optional [param description], which restricts what files can be picked.
A [param filter] should be of the form [code]"filename.extension"[/code], where filename and extension can be [code]*[/code] to match any string. Filters starting with [code].[/code] (i.e. empty filenames) are not allowed.
For example, a [param filter] of [code]"*.tscn, *.scn"[/code] and a [param description] of [code]"Scenes"[/code] results in filter text "Scenes (*.tscn, *.scn)".
*/
func (self Instance) AddFilter(filter string) { //gd:EditorFileDialog.add_filter
	class(self).AddFilter(String.New(filter), String.New(""))
}

/*
Returns the name of the [OptionButton] or [CheckBox] with index [param option].
*/
func (self Instance) GetOptionName(option int) string { //gd:EditorFileDialog.get_option_name
	return string(class(self).GetOptionName(gd.Int(option)).String())
}

/*
Returns an array of values of the [OptionButton] with index [param option].
*/
func (self Instance) GetOptionValues(option int) []string { //gd:EditorFileDialog.get_option_values
	return []string(class(self).GetOptionValues(gd.Int(option)).Strings())
}

/*
Returns the default value index of the [OptionButton] or [CheckBox] with index [param option].
*/
func (self Instance) GetOptionDefault(option int) int { //gd:EditorFileDialog.get_option_default
	return int(int(class(self).GetOptionDefault(gd.Int(option))))
}

/*
Sets the name of the [OptionButton] or [CheckBox] with index [param option].
*/
func (self Instance) SetOptionName(option int, name string) { //gd:EditorFileDialog.set_option_name
	class(self).SetOptionName(gd.Int(option), String.New(name))
}

/*
Sets the option values of the [OptionButton] with index [param option].
*/
func (self Instance) SetOptionValues(option int, values []string) { //gd:EditorFileDialog.set_option_values
	class(self).SetOptionValues(gd.Int(option), gd.NewPackedStringSlice(values))
}

/*
Sets the default value index of the [OptionButton] or [CheckBox] with index [param option].
*/
func (self Instance) SetOptionDefault(option int, default_value_index int) { //gd:EditorFileDialog.set_option_default
	class(self).SetOptionDefault(gd.Int(option), gd.Int(default_value_index))
}

/*
Adds an additional [OptionButton] to the file dialog. If [param values] is empty, a [CheckBox] is added instead.
[param default_value_index] should be an index of the value in the [param values]. If [param values] is empty it should be either [code]1[/code] (checked), or [code]0[/code] (unchecked).
*/
func (self Instance) AddOption(name string, values []string, default_value_index int) { //gd:EditorFileDialog.add_option
	class(self).AddOption(String.New(name), gd.NewPackedStringSlice(values), gd.Int(default_value_index))
}

/*
Returns a [Dictionary] with the selected values of the additional [OptionButton]s and/or [CheckBox]es. [Dictionary] keys are names and values are selected value indices.
*/
func (self Instance) GetSelectedOptions() map[string]int { //gd:EditorFileDialog.get_selected_options
	return map[string]int(gd.DictionaryAs[map[string]int](class(self).GetSelectedOptions()))
}

/*
Returns the [VBoxContainer] used to display the file system.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Instance) GetVbox() [1]gdclass.VBoxContainer { //gd:EditorFileDialog.get_vbox
	return [1]gdclass.VBoxContainer(class(self).GetVbox())
}

/*
Returns the LineEdit for the selected file.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Instance) GetLineEdit() [1]gdclass.LineEdit { //gd:EditorFileDialog.get_line_edit
	return [1]gdclass.LineEdit(class(self).GetLineEdit())
}

/*
Adds the given [param menu] to the side of the file dialog with the given [param title] text on top. Only one side menu is allowed.
*/
func (self Instance) AddSideMenu(menu [1]gdclass.Control) { //gd:EditorFileDialog.add_side_menu
	class(self).AddSideMenu(menu, String.New(""))
}

/*
Shows the [EditorFileDialog] at the default size and position for file dialogs in the editor, and selects the file name if there is a current file.
*/
func (self Instance) PopupFileDialog() { //gd:EditorFileDialog.popup_file_dialog
	class(self).PopupFileDialog()
}

/*
Notify the [EditorFileDialog] that its view of the data is no longer accurate. Updates the view contents on next view update.
*/
func (self Instance) Invalidate() { //gd:EditorFileDialog.invalidate
	class(self).Invalidate()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorFileDialog

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorFileDialog"))
	casted := Instance{*(*gdclass.EditorFileDialog)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Access() gdclass.EditorFileDialogAccess {
	return gdclass.EditorFileDialogAccess(class(self).GetAccess())
}

func (self Instance) SetAccess(value gdclass.EditorFileDialogAccess) {
	class(self).SetAccess(value)
}

func (self Instance) DisplayMode() gdclass.EditorFileDialogDisplayMode {
	return gdclass.EditorFileDialogDisplayMode(class(self).GetDisplayMode())
}

func (self Instance) SetDisplayMode(value gdclass.EditorFileDialogDisplayMode) {
	class(self).SetDisplayMode(value)
}

func (self Instance) FileMode() gdclass.EditorFileDialogFileMode {
	return gdclass.EditorFileDialogFileMode(class(self).GetFileMode())
}

func (self Instance) SetFileMode(value gdclass.EditorFileDialogFileMode) {
	class(self).SetFileMode(value)
}

func (self Instance) CurrentDir() string {
	return string(class(self).GetCurrentDir().String())
}

func (self Instance) SetCurrentDir(value string) {
	class(self).SetCurrentDir(String.New(value))
}

func (self Instance) CurrentFile() string {
	return string(class(self).GetCurrentFile().String())
}

func (self Instance) SetCurrentFile(value string) {
	class(self).SetCurrentFile(String.New(value))
}

func (self Instance) CurrentPath() string {
	return string(class(self).GetCurrentPath().String())
}

func (self Instance) SetCurrentPath(value string) {
	class(self).SetCurrentPath(String.New(value))
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

func (self Instance) DisableOverwriteWarning() bool {
	return bool(class(self).IsOverwriteWarningDisabled())
}

func (self Instance) SetDisableOverwriteWarning(value bool) {
	class(self).SetDisableOverwriteWarning(value)
}

/*
Removes all filters except for "All Files (*)".
*/
//go:nosplit
func (self class) ClearFilters() { //gd:EditorFileDialog.clear_filters
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_clear_filters, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a comma-delimited file name [param filter] option to the [EditorFileDialog] with an optional [param description], which restricts what files can be picked.
A [param filter] should be of the form [code]"filename.extension"[/code], where filename and extension can be [code]*[/code] to match any string. Filters starting with [code].[/code] (i.e. empty filenames) are not allowed.
For example, a [param filter] of [code]"*.tscn, *.scn"[/code] and a [param description] of [code]"Scenes"[/code] results in filter text "Scenes (*.tscn, *.scn)".
*/
//go:nosplit
func (self class) AddFilter(filter String.Readable, description String.Readable) { //gd:EditorFileDialog.add_filter
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(filter)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(description)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_add_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFilters(filters gd.PackedStringArray) { //gd:EditorFileDialog.set_filters
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(filters))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_set_filters, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFilters() gd.PackedStringArray { //gd:EditorFileDialog.get_filters
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_get_filters, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the name of the [OptionButton] or [CheckBox] with index [param option].
*/
//go:nosplit
func (self class) GetOptionName(option gd.Int) String.Readable { //gd:EditorFileDialog.get_option_name
	var frame = callframe.New()
	callframe.Arg(frame, option)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_get_option_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an array of values of the [OptionButton] with index [param option].
*/
//go:nosplit
func (self class) GetOptionValues(option gd.Int) gd.PackedStringArray { //gd:EditorFileDialog.get_option_values
	var frame = callframe.New()
	callframe.Arg(frame, option)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_get_option_values, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the default value index of the [OptionButton] or [CheckBox] with index [param option].
*/
//go:nosplit
func (self class) GetOptionDefault(option gd.Int) gd.Int { //gd:EditorFileDialog.get_option_default
	var frame = callframe.New()
	callframe.Arg(frame, option)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_get_option_default, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the name of the [OptionButton] or [CheckBox] with index [param option].
*/
//go:nosplit
func (self class) SetOptionName(option gd.Int, name String.Readable) { //gd:EditorFileDialog.set_option_name
	var frame = callframe.New()
	callframe.Arg(frame, option)
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_set_option_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the option values of the [OptionButton] with index [param option].
*/
//go:nosplit
func (self class) SetOptionValues(option gd.Int, values gd.PackedStringArray) { //gd:EditorFileDialog.set_option_values
	var frame = callframe.New()
	callframe.Arg(frame, option)
	callframe.Arg(frame, pointers.Get(values))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_set_option_values, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the default value index of the [OptionButton] or [CheckBox] with index [param option].
*/
//go:nosplit
func (self class) SetOptionDefault(option gd.Int, default_value_index gd.Int) { //gd:EditorFileDialog.set_option_default
	var frame = callframe.New()
	callframe.Arg(frame, option)
	callframe.Arg(frame, default_value_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_set_option_default, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetOptionCount(count gd.Int) { //gd:EditorFileDialog.set_option_count
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_set_option_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOptionCount() gd.Int { //gd:EditorFileDialog.get_option_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_get_option_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds an additional [OptionButton] to the file dialog. If [param values] is empty, a [CheckBox] is added instead.
[param default_value_index] should be an index of the value in the [param values]. If [param values] is empty it should be either [code]1[/code] (checked), or [code]0[/code] (unchecked).
*/
//go:nosplit
func (self class) AddOption(name String.Readable, values gd.PackedStringArray, default_value_index gd.Int) { //gd:EditorFileDialog.add_option
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	callframe.Arg(frame, pointers.Get(values))
	callframe.Arg(frame, default_value_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_add_option, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a [Dictionary] with the selected values of the additional [OptionButton]s and/or [CheckBox]es. [Dictionary] keys are names and values are selected value indices.
*/
//go:nosplit
func (self class) GetSelectedOptions() Dictionary.Any { //gd:EditorFileDialog.get_selected_options
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_get_selected_options, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetCurrentDir() String.Readable { //gd:EditorFileDialog.get_current_dir
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_get_current_dir, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetCurrentFile() String.Readable { //gd:EditorFileDialog.get_current_file
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_get_current_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetCurrentPath() String.Readable { //gd:EditorFileDialog.get_current_path
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_get_current_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCurrentDir(dir String.Readable) { //gd:EditorFileDialog.set_current_dir
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(dir)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_set_current_dir, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetCurrentFile(file String.Readable) { //gd:EditorFileDialog.set_current_file
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(file)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_set_current_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetCurrentPath(path String.Readable) { //gd:EditorFileDialog.set_current_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_set_current_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFileMode(mode gdclass.EditorFileDialogFileMode) { //gd:EditorFileDialog.set_file_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_set_file_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFileMode() gdclass.EditorFileDialogFileMode { //gd:EditorFileDialog.get_file_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.EditorFileDialogFileMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_get_file_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [VBoxContainer] used to display the file system.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetVbox() [1]gdclass.VBoxContainer { //gd:EditorFileDialog.get_vbox
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_get_vbox, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.VBoxContainer{gd.PointerLifetimeBoundTo[gdclass.VBoxContainer](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the LineEdit for the selected file.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetLineEdit() [1]gdclass.LineEdit { //gd:EditorFileDialog.get_line_edit
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_get_line_edit, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.LineEdit{gd.PointerLifetimeBoundTo[gdclass.LineEdit](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAccess(access gdclass.EditorFileDialogAccess) { //gd:EditorFileDialog.set_access
	var frame = callframe.New()
	callframe.Arg(frame, access)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_set_access, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAccess() gdclass.EditorFileDialogAccess { //gd:EditorFileDialog.get_access
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.EditorFileDialogAccess](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_get_access, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShowHiddenFiles(show bool) { //gd:EditorFileDialog.set_show_hidden_files
	var frame = callframe.New()
	callframe.Arg(frame, show)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_set_show_hidden_files, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShowingHiddenFiles() bool { //gd:EditorFileDialog.is_showing_hidden_files
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_is_showing_hidden_files, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisplayMode(mode gdclass.EditorFileDialogDisplayMode) { //gd:EditorFileDialog.set_display_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_set_display_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDisplayMode() gdclass.EditorFileDialogDisplayMode { //gd:EditorFileDialog.get_display_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.EditorFileDialogDisplayMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_get_display_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisableOverwriteWarning(disable bool) { //gd:EditorFileDialog.set_disable_overwrite_warning
	var frame = callframe.New()
	callframe.Arg(frame, disable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_set_disable_overwrite_warning, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsOverwriteWarningDisabled() bool { //gd:EditorFileDialog.is_overwrite_warning_disabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_is_overwrite_warning_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds the given [param menu] to the side of the file dialog with the given [param title] text on top. Only one side menu is allowed.
*/
//go:nosplit
func (self class) AddSideMenu(menu [1]gdclass.Control, title String.Readable) { //gd:EditorFileDialog.add_side_menu
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(menu[0].AsObject()[0]))
	callframe.Arg(frame, pointers.Get(gd.InternalString(title)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_add_side_menu, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Shows the [EditorFileDialog] at the default size and position for file dialogs in the editor, and selects the file name if there is a current file.
*/
//go:nosplit
func (self class) PopupFileDialog() { //gd:EditorFileDialog.popup_file_dialog
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_popup_file_dialog, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Notify the [EditorFileDialog] that its view of the data is no longer accurate. Updates the view contents on next view update.
*/
//go:nosplit
func (self class) Invalidate() { //gd:EditorFileDialog.invalidate
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileDialog.Bind_invalidate, self.AsObject(), frame.Array(0), r_ret.Addr())
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

func (self class) AsEditorFileDialog() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorFileDialog() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("EditorFileDialog", func(ptr gd.Object) any {
		return [1]gdclass.EditorFileDialog{*(*gdclass.EditorFileDialog)(unsafe.Pointer(&ptr))}
	})
}

type FileMode = gdclass.EditorFileDialogFileMode //gd:EditorFileDialog.FileMode

const (
	/*The [EditorFileDialog] can select only one file. Accepting the window will open the file.*/
	FileModeOpenFile FileMode = 0
	/*The [EditorFileDialog] can select multiple files. Accepting the window will open all files.*/
	FileModeOpenFiles FileMode = 1
	/*The [EditorFileDialog] can select only one directory. Accepting the window will open the directory.*/
	FileModeOpenDir FileMode = 2
	/*The [EditorFileDialog] can select a file or directory. Accepting the window will open it.*/
	FileModeOpenAny FileMode = 3
	/*The [EditorFileDialog] can select only one file. Accepting the window will save the file.*/
	FileModeSaveFile FileMode = 4
)

type Access = gdclass.EditorFileDialogAccess //gd:EditorFileDialog.Access

const (
	/*The [EditorFileDialog] can only view [code]res://[/code] directory contents.*/
	AccessResources Access = 0
	/*The [EditorFileDialog] can only view [code]user://[/code] directory contents.*/
	AccessUserdata Access = 1
	/*The [EditorFileDialog] can view the entire local file system.*/
	AccessFilesystem Access = 2
)

type DisplayMode = gdclass.EditorFileDialogDisplayMode //gd:EditorFileDialog.DisplayMode

const (
	/*The [EditorFileDialog] displays resources as thumbnails.*/
	DisplayThumbnails DisplayMode = 0
	/*The [EditorFileDialog] displays resources as a list of filenames.*/
	DisplayList DisplayMode = 1
)
