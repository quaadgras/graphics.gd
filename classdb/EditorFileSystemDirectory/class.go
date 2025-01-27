// Package EditorFileSystemDirectory provides methods for working with EditorFileSystemDirectory object instances.
package EditorFileSystemDirectory

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

/*
A more generalized, low-level variation of the directory concept.
*/
type Instance [1]gdclass.EditorFileSystemDirectory

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorFileSystemDirectory() Instance
}

/*
Returns the number of subdirectories in this directory.
*/
func (self Instance) GetSubdirCount() int { //gd:EditorFileSystemDirectory.get_subdir_count
	return int(int(class(self).GetSubdirCount()))
}

/*
Returns the subdirectory at index [param idx].
*/
func (self Instance) GetSubdir(idx int) [1]gdclass.EditorFileSystemDirectory { //gd:EditorFileSystemDirectory.get_subdir
	return [1]gdclass.EditorFileSystemDirectory(class(self).GetSubdir(gd.Int(idx)))
}

/*
Returns the number of files in this directory.
*/
func (self Instance) GetFileCount() int { //gd:EditorFileSystemDirectory.get_file_count
	return int(int(class(self).GetFileCount()))
}

/*
Returns the name of the file at index [param idx].
*/
func (self Instance) GetFile(idx int) string { //gd:EditorFileSystemDirectory.get_file
	return string(class(self).GetFile(gd.Int(idx)).String())
}

/*
Returns the path to the file at index [param idx].
*/
func (self Instance) GetFilePath(idx int) string { //gd:EditorFileSystemDirectory.get_file_path
	return string(class(self).GetFilePath(gd.Int(idx)).String())
}

/*
Returns the resource type of the file at index [param idx]. This returns a string such as [code]"Resource"[/code] or [code]"GDScript"[/code], [i]not[/i] a file extension such as [code]".gd"[/code].
*/
func (self Instance) GetFileType(idx int) string { //gd:EditorFileSystemDirectory.get_file_type
	return string(class(self).GetFileType(gd.Int(idx)).String())
}

/*
Returns the name of the script class defined in the file at index [param idx]. If the file doesn't define a script class using the [code]class_name[/code] syntax, this will return an empty string.
*/
func (self Instance) GetFileScriptClassName(idx int) string { //gd:EditorFileSystemDirectory.get_file_script_class_name
	return string(class(self).GetFileScriptClassName(gd.Int(idx)).String())
}

/*
Returns the base class of the script class defined in the file at index [param idx]. If the file doesn't define a script class using the [code]class_name[/code] syntax, this will return an empty string.
*/
func (self Instance) GetFileScriptClassExtends(idx int) string { //gd:EditorFileSystemDirectory.get_file_script_class_extends
	return string(class(self).GetFileScriptClassExtends(gd.Int(idx)).String())
}

/*
Returns [code]true[/code] if the file at index [param idx] imported properly.
*/
func (self Instance) GetFileImportIsValid(idx int) bool { //gd:EditorFileSystemDirectory.get_file_import_is_valid
	return bool(class(self).GetFileImportIsValid(gd.Int(idx)))
}

/*
Returns the name of this directory.
*/
func (self Instance) GetName() string { //gd:EditorFileSystemDirectory.get_name
	return string(class(self).GetName().String())
}

/*
Returns the path to this directory.
*/
func (self Instance) GetPath() string { //gd:EditorFileSystemDirectory.get_path
	return string(class(self).GetPath().String())
}

/*
Returns the parent directory for this directory or [code]null[/code] if called on a directory at [code]res://[/code] or [code]user://[/code].
*/
func (self Instance) GetParent() [1]gdclass.EditorFileSystemDirectory { //gd:EditorFileSystemDirectory.get_parent
	return [1]gdclass.EditorFileSystemDirectory(class(self).GetParent())
}

/*
Returns the index of the file with name [param name] or [code]-1[/code] if not found.
*/
func (self Instance) FindFileIndex(name string) int { //gd:EditorFileSystemDirectory.find_file_index
	return int(int(class(self).FindFileIndex(String.New(name))))
}

/*
Returns the index of the directory with name [param name] or [code]-1[/code] if not found.
*/
func (self Instance) FindDirIndex(name string) int { //gd:EditorFileSystemDirectory.find_dir_index
	return int(int(class(self).FindDirIndex(String.New(name))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorFileSystemDirectory

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorFileSystemDirectory"))
	casted := Instance{*(*gdclass.EditorFileSystemDirectory)(unsafe.Pointer(&object))}
	return casted
}

/*
Returns the number of subdirectories in this directory.
*/
//go:nosplit
func (self class) GetSubdirCount() gd.Int { //gd:EditorFileSystemDirectory.get_subdir_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_subdir_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the subdirectory at index [param idx].
*/
//go:nosplit
func (self class) GetSubdir(idx gd.Int) [1]gdclass.EditorFileSystemDirectory { //gd:EditorFileSystemDirectory.get_subdir
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_subdir, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.EditorFileSystemDirectory{gd.PointerMustAssertInstanceID[gdclass.EditorFileSystemDirectory](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the number of files in this directory.
*/
//go:nosplit
func (self class) GetFileCount() gd.Int { //gd:EditorFileSystemDirectory.get_file_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_file_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the name of the file at index [param idx].
*/
//go:nosplit
func (self class) GetFile(idx gd.Int) String.Readable { //gd:EditorFileSystemDirectory.get_file
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the path to the file at index [param idx].
*/
//go:nosplit
func (self class) GetFilePath(idx gd.Int) String.Readable { //gd:EditorFileSystemDirectory.get_file_path
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_file_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the resource type of the file at index [param idx]. This returns a string such as [code]"Resource"[/code] or [code]"GDScript"[/code], [i]not[/i] a file extension such as [code]".gd"[/code].
*/
//go:nosplit
func (self class) GetFileType(idx gd.Int) gd.StringName { //gd:EditorFileSystemDirectory.get_file_type
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_file_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the name of the script class defined in the file at index [param idx]. If the file doesn't define a script class using the [code]class_name[/code] syntax, this will return an empty string.
*/
//go:nosplit
func (self class) GetFileScriptClassName(idx gd.Int) String.Readable { //gd:EditorFileSystemDirectory.get_file_script_class_name
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_file_script_class_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the base class of the script class defined in the file at index [param idx]. If the file doesn't define a script class using the [code]class_name[/code] syntax, this will return an empty string.
*/
//go:nosplit
func (self class) GetFileScriptClassExtends(idx gd.Int) String.Readable { //gd:EditorFileSystemDirectory.get_file_script_class_extends
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_file_script_class_extends, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the file at index [param idx] imported properly.
*/
//go:nosplit
func (self class) GetFileImportIsValid(idx gd.Int) bool { //gd:EditorFileSystemDirectory.get_file_import_is_valid
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_file_import_is_valid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the name of this directory.
*/
//go:nosplit
func (self class) GetName() String.Readable { //gd:EditorFileSystemDirectory.get_name
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the path to this directory.
*/
//go:nosplit
func (self class) GetPath() String.Readable { //gd:EditorFileSystemDirectory.get_path
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the parent directory for this directory or [code]null[/code] if called on a directory at [code]res://[/code] or [code]user://[/code].
*/
//go:nosplit
func (self class) GetParent() [1]gdclass.EditorFileSystemDirectory { //gd:EditorFileSystemDirectory.get_parent
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_get_parent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.EditorFileSystemDirectory{gd.PointerMustAssertInstanceID[gdclass.EditorFileSystemDirectory](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the index of the file with name [param name] or [code]-1[/code] if not found.
*/
//go:nosplit
func (self class) FindFileIndex(name String.Readable) gd.Int { //gd:EditorFileSystemDirectory.find_file_index
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_find_file_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the index of the directory with name [param name] or [code]-1[/code] if not found.
*/
//go:nosplit
func (self class) FindDirIndex(name String.Readable) gd.Int { //gd:EditorFileSystemDirectory.find_dir_index
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystemDirectory.Bind_find_dir_index, self.AsObject(), frame.Array(0), r_ret.Addr())
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
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Instance(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("EditorFileSystemDirectory", func(ptr gd.Object) any {
		return [1]gdclass.EditorFileSystemDirectory{*(*gdclass.EditorFileSystemDirectory)(unsafe.Pointer(&ptr))}
	})
}
