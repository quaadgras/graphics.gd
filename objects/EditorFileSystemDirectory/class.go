package EditorFileSystemDirectory

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
A more generalized, low-level variation of the directory concept.
*/
type Instance [1]classdb.EditorFileSystemDirectory

/*
Returns the number of subdirectories in this directory.
*/
func (self Instance) GetSubdirCount() int {
	return int(int(class(self).GetSubdirCount()))
}

/*
Returns the subdirectory at index [param idx].
*/
func (self Instance) GetSubdir(idx int) objects.EditorFileSystemDirectory {
	return objects.EditorFileSystemDirectory(class(self).GetSubdir(gd.Int(idx)))
}

/*
Returns the number of files in this directory.
*/
func (self Instance) GetFileCount() int {
	return int(int(class(self).GetFileCount()))
}

/*
Returns the name of the file at index [param idx].
*/
func (self Instance) GetFile(idx int) string {
	return string(class(self).GetFile(gd.Int(idx)).String())
}

/*
Returns the path to the file at index [param idx].
*/
func (self Instance) GetFilePath(idx int) string {
	return string(class(self).GetFilePath(gd.Int(idx)).String())
}

/*
Returns the resource type of the file at index [param idx]. This returns a string such as [code]"Resource"[/code] or [code]"GDScript"[/code], [i]not[/i] a file extension such as [code]".gd"[/code].
*/
func (self Instance) GetFileType(idx int) string {
	return string(class(self).GetFileType(gd.Int(idx)).String())
}

/*
Returns the name of the script class defined in the file at index [param idx]. If the file doesn't define a script class using the [code]class_name[/code] syntax, this will return an empty string.
*/
func (self Instance) GetFileScriptClassName(idx int) string {
	return string(class(self).GetFileScriptClassName(gd.Int(idx)).String())
}

/*
Returns the base class of the script class defined in the file at index [param idx]. If the file doesn't define a script class using the [code]class_name[/code] syntax, this will return an empty string.
*/
func (self Instance) GetFileScriptClassExtends(idx int) string {
	return string(class(self).GetFileScriptClassExtends(gd.Int(idx)).String())
}

/*
Returns [code]true[/code] if the file at index [param idx] imported properly.
*/
func (self Instance) GetFileImportIsValid(idx int) bool {
	return bool(class(self).GetFileImportIsValid(gd.Int(idx)))
}

/*
Returns the name of this directory.
*/
func (self Instance) GetName() string {
	return string(class(self).GetName().String())
}

/*
Returns the path to this directory.
*/
func (self Instance) GetPath() string {
	return string(class(self).GetPath().String())
}

/*
Returns the parent directory for this directory or [code]null[/code] if called on a directory at [code]res://[/code] or [code]user://[/code].
*/
func (self Instance) GetParent() objects.EditorFileSystemDirectory {
	return objects.EditorFileSystemDirectory(class(self).GetParent())
}

/*
Returns the index of the file with name [param name] or [code]-1[/code] if not found.
*/
func (self Instance) FindFileIndex(name string) int {
	return int(int(class(self).FindFileIndex(gd.NewString(name))))
}

/*
Returns the index of the directory with name [param name] or [code]-1[/code] if not found.
*/
func (self Instance) FindDirIndex(name string) int {
	return int(int(class(self).FindDirIndex(gd.NewString(name))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EditorFileSystemDirectory

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorFileSystemDirectory"))
	return Instance{classdb.EditorFileSystemDirectory(object)}
}

/*
Returns the number of subdirectories in this directory.
*/
//go:nosplit
func (self class) GetSubdirCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_subdir_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the subdirectory at index [param idx].
*/
//go:nosplit
func (self class) GetSubdir(idx gd.Int) objects.EditorFileSystemDirectory {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_subdir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.EditorFileSystemDirectory{classdb.EditorFileSystemDirectory(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the number of files in this directory.
*/
//go:nosplit
func (self class) GetFileCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_file_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the name of the file at index [param idx].
*/
//go:nosplit
func (self class) GetFile(idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the path to the file at index [param idx].
*/
//go:nosplit
func (self class) GetFilePath(idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_file_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the resource type of the file at index [param idx]. This returns a string such as [code]"Resource"[/code] or [code]"GDScript"[/code], [i]not[/i] a file extension such as [code]".gd"[/code].
*/
//go:nosplit
func (self class) GetFileType(idx gd.Int) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_file_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the name of the script class defined in the file at index [param idx]. If the file doesn't define a script class using the [code]class_name[/code] syntax, this will return an empty string.
*/
//go:nosplit
func (self class) GetFileScriptClassName(idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_file_script_class_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the base class of the script class defined in the file at index [param idx]. If the file doesn't define a script class using the [code]class_name[/code] syntax, this will return an empty string.
*/
//go:nosplit
func (self class) GetFileScriptClassExtends(idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_file_script_class_extends, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the file at index [param idx] imported properly.
*/
//go:nosplit
func (self class) GetFileImportIsValid(idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_file_import_is_valid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the name of this directory.
*/
//go:nosplit
func (self class) GetName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the path to this directory.
*/
//go:nosplit
func (self class) GetPath() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the parent directory for this directory or [code]null[/code] if called on a directory at [code]res://[/code] or [code]user://[/code].
*/
//go:nosplit
func (self class) GetParent() objects.EditorFileSystemDirectory {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.EditorFileSystemDirectory{classdb.EditorFileSystemDirectory(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the index of the file with name [param name] or [code]-1[/code] if not found.
*/
//go:nosplit
func (self class) FindFileIndex(name gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_find_file_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the index of the directory with name [param name] or [code]-1[/code] if not found.
*/
//go:nosplit
func (self class) FindDirIndex(name gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_find_dir_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsEditorFileSystemDirectory() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorFileSystemDirectory() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("EditorFileSystemDirectory", func(ptr gd.Object) any { return classdb.EditorFileSystemDirectory(ptr) })
}
