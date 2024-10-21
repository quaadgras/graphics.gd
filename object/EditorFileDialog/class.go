package EditorFileDialog

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/ConfirmationDialog"
import "grow.graphics/gd/object/AcceptDialog"
import "grow.graphics/gd/object/Window"
import "grow.graphics/gd/object/Viewport"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[EditorFileDialog] is an enhanced version of [FileDialog] available only to editor plugins. Additional features include list of favorited/recent files and the ability to see files as thumbnails grid instead of list.

*/
type Simple [1]classdb.EditorFileDialog
func (self Simple) ClearFilters() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearFilters()
}
func (self Simple) AddFilter(filter string, description string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddFilter(gc.String(filter), gc.String(description))
}
func (self Simple) SetFilters(filters gd.PackedStringArray) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFilters(filters)
}
func (self Simple) GetFilters() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetFilters(gc))
}
func (self Simple) GetOptionName(option int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetOptionName(gc, gd.Int(option)).String())
}
func (self Simple) GetOptionValues(option int) gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetOptionValues(gc, gd.Int(option)))
}
func (self Simple) GetOptionDefault(option int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetOptionDefault(gd.Int(option))))
}
func (self Simple) SetOptionName(option int, name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOptionName(gd.Int(option), gc.String(name))
}
func (self Simple) SetOptionValues(option int, values gd.PackedStringArray) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOptionValues(gd.Int(option), values)
}
func (self Simple) SetOptionDefault(option int, default_value_index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOptionDefault(gd.Int(option), gd.Int(default_value_index))
}
func (self Simple) SetOptionCount(count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOptionCount(gd.Int(count))
}
func (self Simple) GetOptionCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetOptionCount()))
}
func (self Simple) AddOption(name string, values gd.PackedStringArray, default_value_index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddOption(gc.String(name), values, gd.Int(default_value_index))
}
func (self Simple) GetSelectedOptions() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).GetSelectedOptions(gc))
}
func (self Simple) GetCurrentDir() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetCurrentDir(gc).String())
}
func (self Simple) GetCurrentFile() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetCurrentFile(gc).String())
}
func (self Simple) GetCurrentPath() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetCurrentPath(gc).String())
}
func (self Simple) SetCurrentDir(dir string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCurrentDir(gc.String(dir))
}
func (self Simple) SetCurrentFile(file string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCurrentFile(gc.String(file))
}
func (self Simple) SetCurrentPath(path string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCurrentPath(gc.String(path))
}
func (self Simple) SetFileMode(mode classdb.EditorFileDialogFileMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFileMode(mode)
}
func (self Simple) GetFileMode() classdb.EditorFileDialogFileMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.EditorFileDialogFileMode(Expert(self).GetFileMode())
}
func (self Simple) GetVbox() [1]classdb.VBoxContainer {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.VBoxContainer(Expert(self).GetVbox(gc))
}
func (self Simple) GetLineEdit() [1]classdb.LineEdit {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.LineEdit(Expert(self).GetLineEdit(gc))
}
func (self Simple) SetAccess(access classdb.EditorFileDialogAccess) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAccess(access)
}
func (self Simple) GetAccess() classdb.EditorFileDialogAccess {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.EditorFileDialogAccess(Expert(self).GetAccess())
}
func (self Simple) SetShowHiddenFiles(show bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShowHiddenFiles(show)
}
func (self Simple) IsShowingHiddenFiles() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsShowingHiddenFiles())
}
func (self Simple) SetDisplayMode(mode classdb.EditorFileDialogDisplayMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDisplayMode(mode)
}
func (self Simple) GetDisplayMode() classdb.EditorFileDialogDisplayMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.EditorFileDialogDisplayMode(Expert(self).GetDisplayMode())
}
func (self Simple) SetDisableOverwriteWarning(disable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDisableOverwriteWarning(disable)
}
func (self Simple) IsOverwriteWarningDisabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsOverwriteWarningDisabled())
}
func (self Simple) AddSideMenu(menu [1]classdb.Control, title string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddSideMenu(menu, gc.String(title))
}
func (self Simple) PopupFileDialog() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PopupFileDialog()
}
func (self Simple) Invalidate() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Invalidate()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorFileDialog
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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
func (self class) GetVbox(ctx gd.Lifetime) object.VBoxContainer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_get_vbox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.VBoxContainer
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the LineEdit for the selected file.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetLineEdit(ctx gd.Lifetime) object.LineEdit {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileDialog.Bind_get_line_edit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.LineEdit
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
func (self class) AddSideMenu(menu object.Control, title gd.String)  {
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

//go:nosplit
func (self class) AsEditorFileDialog() Expert { return self[0].AsEditorFileDialog() }


//go:nosplit
func (self Simple) AsEditorFileDialog() Simple { return self[0].AsEditorFileDialog() }


//go:nosplit
func (self class) AsConfirmationDialog() ConfirmationDialog.Expert { return self[0].AsConfirmationDialog() }


//go:nosplit
func (self Simple) AsConfirmationDialog() ConfirmationDialog.Simple { return self[0].AsConfirmationDialog() }


//go:nosplit
func (self class) AsAcceptDialog() AcceptDialog.Expert { return self[0].AsAcceptDialog() }


//go:nosplit
func (self Simple) AsAcceptDialog() AcceptDialog.Simple { return self[0].AsAcceptDialog() }


//go:nosplit
func (self class) AsWindow() Window.Expert { return self[0].AsWindow() }


//go:nosplit
func (self Simple) AsWindow() Window.Simple { return self[0].AsWindow() }


//go:nosplit
func (self class) AsViewport() Viewport.Expert { return self[0].AsViewport() }


//go:nosplit
func (self Simple) AsViewport() Viewport.Simple { return self[0].AsViewport() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

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
func init() {classdb.Register("EditorFileDialog", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
