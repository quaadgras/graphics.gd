package EditorFileSystemDirectory

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
A more generalized, low-level variation of the directory concept.

*/
type Simple [1]classdb.EditorFileSystemDirectory
func (self Simple) GetSubdirCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSubdirCount()))
}
func (self Simple) GetSubdir(idx int) [1]classdb.EditorFileSystemDirectory {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.EditorFileSystemDirectory(Expert(self).GetSubdir(gc, gd.Int(idx)))
}
func (self Simple) GetFileCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetFileCount()))
}
func (self Simple) GetFile(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetFile(gc, gd.Int(idx)).String())
}
func (self Simple) GetFilePath(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetFilePath(gc, gd.Int(idx)).String())
}
func (self Simple) GetFileType(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetFileType(gc, gd.Int(idx)).String())
}
func (self Simple) GetFileScriptClassName(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetFileScriptClassName(gc, gd.Int(idx)).String())
}
func (self Simple) GetFileScriptClassExtends(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetFileScriptClassExtends(gc, gd.Int(idx)).String())
}
func (self Simple) GetFileImportIsValid(idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetFileImportIsValid(gd.Int(idx)))
}
func (self Simple) GetName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetName(gc).String())
}
func (self Simple) GetPath() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetPath(gc).String())
}
func (self Simple) GetParent() [1]classdb.EditorFileSystemDirectory {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.EditorFileSystemDirectory(Expert(self).GetParent(gc))
}
func (self Simple) FindFileIndex(name string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).FindFileIndex(gc.String(name))))
}
func (self Simple) FindDirIndex(name string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).FindDirIndex(gc.String(name))))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorFileSystemDirectory
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the number of subdirectories in this directory.
*/
//go:nosplit
func (self class) GetSubdirCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_subdir_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the subdirectory at index [param idx].
*/
//go:nosplit
func (self class) GetSubdir(ctx gd.Lifetime, idx gd.Int) object.EditorFileSystemDirectory {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_subdir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.EditorFileSystemDirectory
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the number of files in this directory.
*/
//go:nosplit
func (self class) GetFileCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_file_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the name of the file at index [param idx].
*/
//go:nosplit
func (self class) GetFile(ctx gd.Lifetime, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the path to the file at index [param idx].
*/
//go:nosplit
func (self class) GetFilePath(ctx gd.Lifetime, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_file_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the resource type of the file at index [param idx]. This returns a string such as [code]"Resource"[/code] or [code]"GDScript"[/code], [i]not[/i] a file extension such as [code]".gd"[/code].
*/
//go:nosplit
func (self class) GetFileType(ctx gd.Lifetime, idx gd.Int) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_file_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the name of the script class defined in the file at index [param idx]. If the file doesn't define a script class using the [code]class_name[/code] syntax, this will return an empty string.
*/
//go:nosplit
func (self class) GetFileScriptClassName(ctx gd.Lifetime, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_file_script_class_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the base class of the script class defined in the file at index [param idx]. If the file doesn't define a script class using the [code]class_name[/code] syntax, this will return an empty string.
*/
//go:nosplit
func (self class) GetFileScriptClassExtends(ctx gd.Lifetime, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_file_script_class_extends, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the file at index [param idx] imported properly.
*/
//go:nosplit
func (self class) GetFileImportIsValid(idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_file_import_is_valid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the name of this directory.
*/
//go:nosplit
func (self class) GetName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the path to this directory.
*/
//go:nosplit
func (self class) GetPath(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the parent directory for this directory or [code]null[/code] if called on a directory at [code]res://[/code] or [code]user://[/code].
*/
//go:nosplit
func (self class) GetParent(ctx gd.Lifetime) object.EditorFileSystemDirectory {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.EditorFileSystemDirectory
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the index of the file with name [param name] or [code]-1[/code] if not found.
*/
//go:nosplit
func (self class) FindFileIndex(name gd.String) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_find_file_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the index of the directory with name [param name] or [code]-1[/code] if not found.
*/
//go:nosplit
func (self class) FindDirIndex(name gd.String) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_find_dir_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsEditorFileSystemDirectory() Expert { return self[0].AsEditorFileSystemDirectory() }


//go:nosplit
func (self Simple) AsEditorFileSystemDirectory() Simple { return self[0].AsEditorFileSystemDirectory() }

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
func init() {classdb.Register("EditorFileSystemDirectory", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
