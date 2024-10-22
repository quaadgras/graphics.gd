package EditorFileDialog

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
var _ mmm.Lifetime

/*
[EditorFileDialog] is an enhanced version of [FileDialog] available only to editor plugins. Additional features include list of favorited/recent files and the ability to see files as thumbnails grid instead of list.

*/
type Go [1]classdb.EditorFileDialog

/*
Removes all filters except for "All Files (*)".
*/
func (self Go) ClearFilters() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearFilters()
}

/*
Adds a comma-delimited file name [param filter] option to the [EditorFileDialog] with an optional [param description], which restricts what files can be picked.
A [param filter] should be of the form [code]"filename.extension"[/code], where filename and extension can be [code]*[/code] to match any string. Filters starting with [code].[/code] (i.e. empty filenames) are not allowed.
For example, a [param filter] of [code]"*.tscn, *.scn"[/code] and a [param description] of [code]"Scenes"[/code] results in filter text "Scenes (*.tscn, *.scn)".
*/
func (self Go) AddFilter(filter string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddFilter(gc.String(filter), gc.String(""))
}

/*
Returns the name of the [OptionButton] or [CheckBox] with index [param option].
*/
func (self Go) GetOptionName(option int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetOptionName(gc, gd.Int(option)).String())
}

/*
Returns an array of values of the [OptionButton] with index [param option].
*/
func (self Go) GetOptionValues(option int) []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetOptionValues(gc, gd.Int(option)).Strings(gc))
}

/*
Returns the default value index of the [OptionButton] or [CheckBox] with index [param option].
*/
func (self Go) GetOptionDefault(option int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetOptionDefault(gd.Int(option))))
}

/*
Sets the name of the [OptionButton] or [CheckBox] with index [param option].
*/
func (self Go) SetOptionName(option int, name string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOptionName(gd.Int(option), gc.String(name))
}

/*
Sets the option values of the [OptionButton] with index [param option].
*/
func (self Go) SetOptionValues(option int, values []string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOptionValues(gd.Int(option), gc.PackedStringSlice(values))
}

/*
Sets the default value index of the [OptionButton] or [CheckBox] with index [param option].
*/
func (self Go) SetOptionDefault(option int, default_value_index int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOptionDefault(gd.Int(option), gd.Int(default_value_index))
}

/*
Adds an additional [OptionButton] to the file dialog. If [param values] is empty, a [CheckBox] is added instead.
[param default_value_index] should be an index of the value in the [param values]. If [param values] is empty it should be either [code]1[/code] (checked), or [code]0[/code] (unchecked).
*/
func (self Go) AddOption(name string, values []string, default_value_index int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddOption(gc.String(name), gc.PackedStringSlice(values), gd.Int(default_value_index))
}

/*
Returns a [Dictionary] with the selected values of the additional [OptionButton]s and/or [CheckBox]es. [Dictionary] keys are names and values are selected value indices.
*/
func (self Go) GetSelectedOptions() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).GetSelectedOptions(gc))
}

/*
Returns the [VBoxContainer] used to display the file system.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Go) GetVbox() gdclass.VBoxContainer {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.VBoxContainer(class(self).GetVbox(gc))
}

/*
Returns the LineEdit for the selected file.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Go) GetLineEdit() gdclass.LineEdit {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.LineEdit(class(self).GetLineEdit(gc))
}

/*
Adds the given [param menu] to the side of the file dialog with the given [param title] text on top. Only one side menu is allowed.
*/
func (self Go) AddSideMenu(menu gdclass.Control) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddSideMenu(menu, gc.String(""))
}

/*
Shows the [EditorFileDialog] at the default size and position for file dialogs in the editor, and selects the file name if there is a current file.
*/
func (self Go) PopupFileDialog() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PopupFileDialog()
}

/*
Notify the [EditorFileDialog] that its view of the data is no longer accurate. Updates the view contents on next view update.
*/
func (self Go) Invalidate() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Invalidate()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorFileDialog
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("EditorFileDialog"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Access() classdb.EditorFileDialogAccess {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.EditorFileDialogAccess(class(self).GetAccess())
}

func (self Go) SetAccess(value classdb.EditorFileDialogAccess) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAccess(value)
}

func (self Go) DisplayMode() classdb.EditorFileDialogDisplayMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.EditorFileDialogDisplayMode(class(self).GetDisplayMode())
}

func (self Go) SetDisplayMode(value classdb.EditorFileDialogDisplayMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDisplayMode(value)
}

func (self Go) FileMode() classdb.EditorFileDialogFileMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.EditorFileDialogFileMode(class(self).GetFileMode())
}

func (self Go) SetFileMode(value classdb.EditorFileDialogFileMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFileMode(value)
}

func (self Go) CurrentDir() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetCurrentDir(gc).String())
}

func (self Go) SetCurrentDir(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCurrentDir(gc.String(value))
}

func (self Go) CurrentFile() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetCurrentFile(gc).String())
}

func (self Go) SetCurrentFile(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCurrentFile(gc.String(value))
}

func (self Go) CurrentPath() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetCurrentPath(gc).String())
}

func (self Go) SetCurrentPath(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCurrentPath(gc.String(value))
}

func (self Go) Filters() []string {
	gc := gd.GarbageCollector(); _ = gc
		return []string(class(self).GetFilters(gc).Strings(gc))
}

func (self Go) SetFilters(value []string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFilters(gc.PackedStringSlice(value))
}

func (self Go) OptionCount() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetOptionCount()))
}

func (self Go) SetOptionCount(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOptionCount(gd.Int(value))
}

func (self Go) ShowHiddenFiles() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsShowingHiddenFiles())
}

func (self Go) SetShowHiddenFiles(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShowHiddenFiles(value)
}

func (self Go) DisableOverwriteWarning() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsOverwriteWarningDisabled())
}

func (self Go) SetDisableOverwriteWarning(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDisableOverwriteWarning(value)
}

/*
Removes all filters except for "All Files (*)".
*/
//go:nosplit
func (self class) ClearFilters()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_clear_filters, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a comma-delimited file name [param filter] option to the [EditorFileDialog] with an optional [param description], which restricts what files can be picked.
A [param filter] should be of the form [code]"filename.extension"[/code], where filename and extension can be [code]*[/code] to match any string. Filters starting with [code].[/code] (i.e. empty filenames) are not allowed.
For example, a [param filter] of [code]"*.tscn, *.scn"[/code] and a [param description] of [code]"Scenes"[/code] results in filter text "Scenes (*.tscn, *.scn)".
*/
//go:nosplit
func (self class) AddFilter(filter gd.String, description gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(filter))
	callframe.Arg(frame, mmm.Get(description))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_add_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFilters(filters gd.PackedStringArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(filters))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_set_filters, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFilters(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_get_filters, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the name of the [OptionButton] or [CheckBox] with index [param option].
*/
//go:nosplit
func (self class) GetOptionName(ctx gd.Lifetime, option gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, option)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_get_option_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array of values of the [OptionButton] with index [param option].
*/
//go:nosplit
func (self class) GetOptionValues(ctx gd.Lifetime, option gd.Int) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, option)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_get_option_values, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the default value index of the [OptionButton] or [CheckBox] with index [param option].
*/
//go:nosplit
func (self class) GetOptionDefault(option gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, option)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_get_option_default, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the name of the [OptionButton] or [CheckBox] with index [param option].
*/
//go:nosplit
func (self class) SetOptionName(option gd.Int, name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, option)
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_set_option_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the option values of the [OptionButton] with index [param option].
*/
//go:nosplit
func (self class) SetOptionValues(option gd.Int, values gd.PackedStringArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, option)
	callframe.Arg(frame, mmm.Get(values))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_set_option_values, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the default value index of the [OptionButton] or [CheckBox] with index [param option].
*/
//go:nosplit
func (self class) SetOptionDefault(option gd.Int, default_value_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, option)
	callframe.Arg(frame, default_value_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_set_option_default, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetOptionCount(count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_set_option_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOptionCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_get_option_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(values))
	callframe.Arg(frame, default_value_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_add_option, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [Dictionary] with the selected values of the additional [OptionButton]s and/or [CheckBox]es. [Dictionary] keys are names and values are selected value indices.
*/
//go:nosplit
func (self class) GetSelectedOptions(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_get_selected_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetCurrentDir(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_get_current_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetCurrentFile(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_get_current_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetCurrentPath(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_get_current_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCurrentDir(dir gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(dir))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_set_current_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCurrentFile(file gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(file))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_set_current_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCurrentPath(path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_set_current_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFileMode(mode classdb.EditorFileDialogFileMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_set_file_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFileMode() classdb.EditorFileDialogFileMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EditorFileDialogFileMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_get_file_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [VBoxContainer] used to display the file system.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetVbox(ctx gd.Lifetime) gdclass.VBoxContainer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_get_vbox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.VBoxContainer
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the LineEdit for the selected file.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetLineEdit(ctx gd.Lifetime) gdclass.LineEdit {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_get_line_edit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.LineEdit
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAccess(access classdb.EditorFileDialogAccess)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, access)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_set_access, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAccess() classdb.EditorFileDialogAccess {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EditorFileDialogAccess](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_get_access, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShowHiddenFiles(show bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, show)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_set_show_hidden_files, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShowingHiddenFiles() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_is_showing_hidden_files, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDisplayMode(mode classdb.EditorFileDialogDisplayMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_set_display_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDisplayMode() classdb.EditorFileDialogDisplayMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EditorFileDialogDisplayMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_get_display_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDisableOverwriteWarning(disable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_set_disable_overwrite_warning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsOverwriteWarningDisabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_is_overwrite_warning_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds the given [param menu] to the side of the file dialog with the given [param title] text on top. Only one side menu is allowed.
*/
//go:nosplit
func (self class) AddSideMenu(menu gdclass.Control, title gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(menu[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(title))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_add_side_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Shows the [EditorFileDialog] at the default size and position for file dialogs in the editor, and selects the file name if there is a current file.
*/
//go:nosplit
func (self class) PopupFileDialog()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_popup_file_dialog, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Notify the [EditorFileDialog] that its view of the data is no longer accurate. Updates the view contents on next view update.
*/
//go:nosplit
func (self class) Invalidate()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_invalidate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnFileSelected(cb func(path string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("file_selected"), gc.Callable(cb), 0)
}


func (self Go) OnFilesSelected(cb func(paths []string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("files_selected"), gc.Callable(cb), 0)
}


func (self Go) OnDirSelected(cb func(dir string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("dir_selected"), gc.Callable(cb), 0)
}


func (self class) AsEditorFileDialog() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorFileDialog() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("EditorFileDialog", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type FileMode = classdb.EditorFileDialogFileMode

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
type Access = classdb.EditorFileDialogAccess

const (
/*The [EditorFileDialog] can only view [code]res://[/code] directory contents.*/
	AccessResources Access = 0
/*The [EditorFileDialog] can only view [code]user://[/code] directory contents.*/
	AccessUserdata Access = 1
/*The [EditorFileDialog] can view the entire local file system.*/
	AccessFilesystem Access = 2
)
type DisplayMode = classdb.EditorFileDialogDisplayMode

const (
/*The [EditorFileDialog] displays resources as thumbnails.*/
	DisplayThumbnails DisplayMode = 0
/*The [EditorFileDialog] displays resources as a list of filenames.*/
	DisplayList DisplayMode = 1
)
