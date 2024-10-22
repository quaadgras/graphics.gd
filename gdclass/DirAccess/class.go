package DirAccess

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class is used to manage directories and their content, even outside of the project folder.
[DirAccess] can't be instantiated directly. Instead it is created with a static method that takes a path for which it will be opened.
Most of the methods have a static alternative that can be used without creating a [DirAccess]. Static methods only support absolute paths (including [code]res://[/code] and [code]user://[/code]).
[codeblock]
# Standard
var dir = DirAccess.open("user://levels")
dir.make_dir("world1")
# Static
DirAccess.make_dir_absolute("user://levels/world1")
[/codeblock]
[b]Note:[/b] Many resources types are imported (e.g. textures or sound files), and their source asset will not be included in the exported game, as only the imported version is used. Use [ResourceLoader] to access imported resources.
Here is an example on how to iterate through the files of a directory:
[codeblocks]
[gdscript]
func dir_contents(path):
    var dir = DirAccess.open(path)
    if dir:
        dir.list_dir_begin()
        var file_name = dir.get_next()
        while file_name != "":
            if dir.current_is_dir():
                print("Found directory: " + file_name)
            else:
                print("Found file: " + file_name)
            file_name = dir.get_next()
    else:
        print("An error occurred when trying to access the path.")
[/gdscript]
[csharp]
public void DirContents(string path)
{
    using var dir = DirAccess.Open(path);
    if (dir != null)
    {
        dir.ListDirBegin();
        string fileName = dir.GetNext();
        while (fileName != "")
        {
            if (dir.CurrentIsDir())
            {
                GD.Print($"Found directory: {fileName}");
            }
            else
            {
                GD.Print($"Found file: {fileName}");
            }
            fileName = dir.GetNext();
        }
    }
    else
    {
        GD.Print("An error occurred when trying to access the path.");
    }
}
[/csharp]
[/codeblocks]

*/
type Go [1]classdb.DirAccess

/*
Creates a new [DirAccess] object and opens an existing directory of the filesystem. The [param path] argument can be within the project tree ([code]res://folder[/code]), the user directory ([code]user://folder[/code]) or an absolute path of the user filesystem (e.g. [code]/tmp/folder[/code] or [code]C:\tmp\folder[/code]).
Returns [code]null[/code] if opening the directory failed. You can use [method get_open_error] to check the error that occurred.
*/
func (self Go) Open(path string) gdclass.DirAccess {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.DirAccess(class(self).Open(gc, gc.String(path)))
}

/*
Returns the result of the last [method open] call in the current thread.
*/
func (self Go) GetOpenError() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).GetOpenError(gc))
}

/*
Initializes the stream used to list all files and directories using the [method get_next] function, closing the currently opened stream if needed. Once the stream has been processed, it should typically be closed with [method list_dir_end].
Affected by [member include_hidden] and [member include_navigational].
[b]Note:[/b] The order of files and directories returned by this method is not deterministic, and can vary between operating systems. If you want a list of all files or folders sorted alphabetically, use [method get_files] or [method get_directories].
*/
func (self Go) ListDirBegin() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).ListDirBegin())
}

/*
Returns the next element (file or directory) in the current directory.
The name of the file or directory is returned (and not its full path). Once the stream has been fully processed, the method returns an empty [String] and closes the stream automatically (i.e. [method list_dir_end] would not be mandatory in such a case).
*/
func (self Go) GetNext() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetNext(gc).String())
}

/*
Returns whether the current item processed with the last [method get_next] call is a directory ([code].[/code] and [code]..[/code] are considered directories).
*/
func (self Go) CurrentIsDir() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).CurrentIsDir())
}

/*
Closes the current stream opened with [method list_dir_begin] (whether it has been fully processed with [method get_next] does not matter).
*/
func (self Go) ListDirEnd() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ListDirEnd()
}

/*
Returns a [PackedStringArray] containing filenames of the directory contents, excluding directories. The array is sorted alphabetically.
Affected by [member include_hidden].
[b]Note:[/b] When used on a [code]res://[/code] path in an exported project, only the files actually included in the PCK at the given folder level are returned. In practice, this means that since imported resources are stored in a top-level [code].godot/[/code] folder, only paths to [code]*.gd[/code] and [code]*.import[/code] files are returned (plus a few files such as [code]project.godot[/code] or [code]project.binary[/code] and the project icon). In an exported project, the list of returned files will also vary depending on whether [member ProjectSettings.editor/export/convert_text_resources_to_binary] is [code]true[/code].
*/
func (self Go) GetFiles() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetFiles(gc).Strings(gc))
}

/*
Returns a [PackedStringArray] containing filenames of the directory contents, excluding directories, at the given [param path]. The array is sorted alphabetically.
Use [method get_files] if you want more control of what gets included.
*/
func (self Go) GetFilesAt(path string) []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetFilesAt(gc, gc.String(path)).Strings(gc))
}

/*
Returns a [PackedStringArray] containing filenames of the directory contents, excluding files. The array is sorted alphabetically.
Affected by [member include_hidden] and [member include_navigational].
*/
func (self Go) GetDirectories() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetDirectories(gc).Strings(gc))
}

/*
Returns a [PackedStringArray] containing filenames of the directory contents, excluding files, at the given [param path]. The array is sorted alphabetically.
Use [method get_directories] if you want more control of what gets included.
*/
func (self Go) GetDirectoriesAt(path string) []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetDirectoriesAt(gc, gc.String(path)).Strings(gc))
}

/*
On Windows, returns the number of drives (partitions) mounted on the current filesystem.
On macOS, returns the number of mounted volumes.
On Linux, returns the number of mounted volumes and GTK 3 bookmarks.
On other platforms, the method returns 0.
*/
func (self Go) GetDriveCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetDriveCount(gc)))
}

/*
On Windows, returns the name of the drive (partition) passed as an argument (e.g. [code]C:[/code]).
On macOS, returns the path to the mounted volume passed as an argument.
On Linux, returns the path to the mounted volume or GTK 3 bookmark passed as an argument.
On other platforms, or if the requested drive does not exist, the method returns an empty String.
*/
func (self Go) GetDriveName(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetDriveName(gc, gd.Int(idx)).String())
}

/*
Returns the currently opened directory's drive index. See [method get_drive_name] to convert returned index to the name of the drive.
*/
func (self Go) GetCurrentDrive() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCurrentDrive()))
}

/*
Changes the currently opened directory to the one passed as an argument. The argument can be relative to the current directory (e.g. [code]newdir[/code] or [code]../newdir[/code]), or an absolute path (e.g. [code]/tmp/newdir[/code] or [code]res://somedir/newdir[/code]).
Returns one of the [enum Error] code constants ([constant OK] on success).
[b]Note:[/b] The new directory must be within the same scope, e.g. when you had opened a directory inside [code]res://[/code], you can't change it to [code]user://[/code] directory. If you need to open a directory in another access scope, use [method open] to create a new instance instead.
*/
func (self Go) ChangeDir(to_dir string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).ChangeDir(gc.String(to_dir)))
}

/*
Returns the absolute path to the currently opened directory (e.g. [code]res://folder[/code] or [code]C:\tmp\folder[/code]).
*/
func (self Go) GetCurrentDir() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetCurrentDir(gc, true).String())
}

/*
Creates a directory. The argument can be relative to the current directory, or an absolute path. The target directory should be placed in an already existing directory (to create the full path recursively, see [method make_dir_recursive]).
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
func (self Go) MakeDir(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).MakeDir(gc.String(path)))
}

/*
Static version of [method make_dir]. Supports only absolute paths.
*/
func (self Go) MakeDirAbsolute(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).MakeDirAbsolute(gc, gc.String(path)))
}

/*
Creates a target directory and all necessary intermediate directories in its path, by calling [method make_dir] recursively. The argument can be relative to the current directory, or an absolute path.
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
func (self Go) MakeDirRecursive(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).MakeDirRecursive(gc.String(path)))
}

/*
Static version of [method make_dir_recursive]. Supports only absolute paths.
*/
func (self Go) MakeDirRecursiveAbsolute(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).MakeDirRecursiveAbsolute(gc, gc.String(path)))
}

/*
Returns whether the target file exists. The argument can be relative to the current directory, or an absolute path.
For a static equivalent, use [method FileAccess.file_exists].
*/
func (self Go) FileExists(path string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).FileExists(gc.String(path)))
}

/*
Returns whether the target directory exists. The argument can be relative to the current directory, or an absolute path.
*/
func (self Go) DirExists(path string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).DirExists(gc.String(path)))
}

/*
Static version of [method dir_exists]. Supports only absolute paths.
*/
func (self Go) DirExistsAbsolute(path string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).DirExistsAbsolute(gc, gc.String(path)))
}

/*
Returns the available space on the current directory's disk, in bytes. Returns [code]0[/code] if the platform-specific method to query the available space fails.
*/
func (self Go) GetSpaceLeft() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSpaceLeft()))
}

/*
Copies the [param from] file to the [param to] destination. Both arguments should be paths to files, either relative or absolute. If the destination file exists and is not access-protected, it will be overwritten.
If [param chmod_flags] is different than [code]-1[/code], the Unix permissions for the destination path will be set to the provided value, if available on the current operating system.
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
func (self Go) Copy(from string, to string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Copy(gc.String(from), gc.String(to), gd.Int(-1)))
}

/*
Static version of [method copy]. Supports only absolute paths.
*/
func (self Go) CopyAbsolute(from string, to string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).CopyAbsolute(gc, gc.String(from), gc.String(to), gd.Int(-1)))
}

/*
Renames (move) the [param from] file or directory to the [param to] destination. Both arguments should be paths to files or directories, either relative or absolute. If the destination file or directory exists and is not access-protected, it will be overwritten.
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
func (self Go) Rename(from string, to string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Rename(gc.String(from), gc.String(to)))
}

/*
Static version of [method rename]. Supports only absolute paths.
*/
func (self Go) RenameAbsolute(from string, to string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).RenameAbsolute(gc, gc.String(from), gc.String(to)))
}

/*
Permanently deletes the target file or an empty directory. The argument can be relative to the current directory, or an absolute path. If the target directory is not empty, the operation will fail.
If you don't want to delete the file/directory permanently, use [method OS.move_to_trash] instead.
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
func (self Go) Remove(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Remove(gc.String(path)))
}

/*
Static version of [method remove]. Supports only absolute paths.
*/
func (self Go) RemoveAbsolute(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).RemoveAbsolute(gc, gc.String(path)))
}

/*
Returns [code]true[/code] if the file or directory is a symbolic link, directory junction, or other reparse point.
[b]Note:[/b] This method is implemented on macOS, Linux, and Windows.
*/
func (self Go) IsLink(path string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsLink(gc.String(path)))
}

/*
Returns target of the symbolic link.
[b]Note:[/b] This method is implemented on macOS, Linux, and Windows.
*/
func (self Go) ReadLink(path string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).ReadLink(gc, gc.String(path)).String())
}

/*
Creates symbolic link between files or folders.
[b]Note:[/b] On Windows, this method works only if the application is running with elevated privileges or Developer Mode is enabled.
[b]Note:[/b] This method is implemented on macOS, Linux, and Windows.
*/
func (self Go) CreateLink(source string, target string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).CreateLink(gc.String(source), gc.String(target)))
}

/*
Returns [code]true[/code] if the file system or directory use case sensitive file names.
[b]Note:[/b] This method is implemented on macOS, Linux (for EXT4 and F2FS filesystems only) and Windows. On other platforms, it always returns [code]true[/code].
*/
func (self Go) IsCaseSensitive(path string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsCaseSensitive(gc.String(path)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.DirAccess
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("DirAccess"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) IncludeNavigational() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetIncludeNavigational())
}

func (self Go) SetIncludeNavigational(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIncludeNavigational(value)
}

func (self Go) IncludeHidden() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetIncludeHidden())
}

func (self Go) SetIncludeHidden(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIncludeHidden(value)
}

/*
Creates a new [DirAccess] object and opens an existing directory of the filesystem. The [param path] argument can be within the project tree ([code]res://folder[/code]), the user directory ([code]user://folder[/code]) or an absolute path of the user filesystem (e.g. [code]/tmp/folder[/code] or [code]C:\tmp\folder[/code]).
Returns [code]null[/code] if opening the directory failed. You can use [method get_open_error] to check the error that occurred.
*/
//go:nosplit
func (self class) Open(ctx gd.Lifetime, path gd.String) gdclass.DirAccess {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.DirAccess.Bind_open, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.DirAccess
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the result of the last [method open] call in the current thread.
*/
//go:nosplit
func (self class) GetOpenError(ctx gd.Lifetime) int64 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.DirAccess.Bind_get_open_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Initializes the stream used to list all files and directories using the [method get_next] function, closing the currently opened stream if needed. Once the stream has been processed, it should typically be closed with [method list_dir_end].
Affected by [member include_hidden] and [member include_navigational].
[b]Note:[/b] The order of files and directories returned by this method is not deterministic, and can vary between operating systems. If you want a list of all files or folders sorted alphabetically, use [method get_files] or [method get_directories].
*/
//go:nosplit
func (self class) ListDirBegin() int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_list_dir_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the next element (file or directory) in the current directory.
The name of the file or directory is returned (and not its full path). Once the stream has been fully processed, the method returns an empty [String] and closes the stream automatically (i.e. [method list_dir_end] would not be mandatory in such a case).
*/
//go:nosplit
func (self class) GetNext(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_get_next, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns whether the current item processed with the last [method get_next] call is a directory ([code].[/code] and [code]..[/code] are considered directories).
*/
//go:nosplit
func (self class) CurrentIsDir() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_current_is_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Closes the current stream opened with [method list_dir_begin] (whether it has been fully processed with [method get_next] does not matter).
*/
//go:nosplit
func (self class) ListDirEnd()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_list_dir_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [PackedStringArray] containing filenames of the directory contents, excluding directories. The array is sorted alphabetically.
Affected by [member include_hidden].
[b]Note:[/b] When used on a [code]res://[/code] path in an exported project, only the files actually included in the PCK at the given folder level are returned. In practice, this means that since imported resources are stored in a top-level [code].godot/[/code] folder, only paths to [code]*.gd[/code] and [code]*.import[/code] files are returned (plus a few files such as [code]project.godot[/code] or [code]project.binary[/code] and the project icon). In an exported project, the list of returned files will also vary depending on whether [member ProjectSettings.editor/export/convert_text_resources_to_binary] is [code]true[/code].
*/
//go:nosplit
func (self class) GetFiles(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_get_files, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a [PackedStringArray] containing filenames of the directory contents, excluding directories, at the given [param path]. The array is sorted alphabetically.
Use [method get_files] if you want more control of what gets included.
*/
//go:nosplit
func (self class) GetFilesAt(ctx gd.Lifetime, path gd.String) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.DirAccess.Bind_get_files_at, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a [PackedStringArray] containing filenames of the directory contents, excluding files. The array is sorted alphabetically.
Affected by [member include_hidden] and [member include_navigational].
*/
//go:nosplit
func (self class) GetDirectories(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_get_directories, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a [PackedStringArray] containing filenames of the directory contents, excluding files, at the given [param path]. The array is sorted alphabetically.
Use [method get_directories] if you want more control of what gets included.
*/
//go:nosplit
func (self class) GetDirectoriesAt(ctx gd.Lifetime, path gd.String) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.DirAccess.Bind_get_directories_at, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
On Windows, returns the number of drives (partitions) mounted on the current filesystem.
On macOS, returns the number of mounted volumes.
On Linux, returns the number of mounted volumes and GTK 3 bookmarks.
On other platforms, the method returns 0.
*/
//go:nosplit
func (self class) GetDriveCount(ctx gd.Lifetime) gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.DirAccess.Bind_get_drive_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
On Windows, returns the name of the drive (partition) passed as an argument (e.g. [code]C:[/code]).
On macOS, returns the path to the mounted volume passed as an argument.
On Linux, returns the path to the mounted volume or GTK 3 bookmark passed as an argument.
On other platforms, or if the requested drive does not exist, the method returns an empty String.
*/
//go:nosplit
func (self class) GetDriveName(ctx gd.Lifetime, idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.DirAccess.Bind_get_drive_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the currently opened directory's drive index. See [method get_drive_name] to convert returned index to the name of the drive.
*/
//go:nosplit
func (self class) GetCurrentDrive() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_get_current_drive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Changes the currently opened directory to the one passed as an argument. The argument can be relative to the current directory (e.g. [code]newdir[/code] or [code]../newdir[/code]), or an absolute path (e.g. [code]/tmp/newdir[/code] or [code]res://somedir/newdir[/code]).
Returns one of the [enum Error] code constants ([constant OK] on success).
[b]Note:[/b] The new directory must be within the same scope, e.g. when you had opened a directory inside [code]res://[/code], you can't change it to [code]user://[/code] directory. If you need to open a directory in another access scope, use [method open] to create a new instance instead.
*/
//go:nosplit
func (self class) ChangeDir(to_dir gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(to_dir))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_change_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the absolute path to the currently opened directory (e.g. [code]res://folder[/code] or [code]C:\tmp\folder[/code]).
*/
//go:nosplit
func (self class) GetCurrentDir(ctx gd.Lifetime, include_drive bool) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, include_drive)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_get_current_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Creates a directory. The argument can be relative to the current directory, or an absolute path. The target directory should be placed in an already existing directory (to create the full path recursively, see [method make_dir_recursive]).
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
//go:nosplit
func (self class) MakeDir(path gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_make_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Static version of [method make_dir]. Supports only absolute paths.
*/
//go:nosplit
func (self class) MakeDirAbsolute(ctx gd.Lifetime, path gd.String) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.DirAccess.Bind_make_dir_absolute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a target directory and all necessary intermediate directories in its path, by calling [method make_dir] recursively. The argument can be relative to the current directory, or an absolute path.
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
//go:nosplit
func (self class) MakeDirRecursive(path gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_make_dir_recursive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Static version of [method make_dir_recursive]. Supports only absolute paths.
*/
//go:nosplit
func (self class) MakeDirRecursiveAbsolute(ctx gd.Lifetime, path gd.String) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.DirAccess.Bind_make_dir_recursive_absolute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether the target file exists. The argument can be relative to the current directory, or an absolute path.
For a static equivalent, use [method FileAccess.file_exists].
*/
//go:nosplit
func (self class) FileExists(path gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_file_exists, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether the target directory exists. The argument can be relative to the current directory, or an absolute path.
*/
//go:nosplit
func (self class) DirExists(path gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_dir_exists, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Static version of [method dir_exists]. Supports only absolute paths.
*/
//go:nosplit
func (self class) DirExistsAbsolute(ctx gd.Lifetime, path gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.DirAccess.Bind_dir_exists_absolute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the available space on the current directory's disk, in bytes. Returns [code]0[/code] if the platform-specific method to query the available space fails.
*/
//go:nosplit
func (self class) GetSpaceLeft() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_get_space_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Copies the [param from] file to the [param to] destination. Both arguments should be paths to files, either relative or absolute. If the destination file exists and is not access-protected, it will be overwritten.
If [param chmod_flags] is different than [code]-1[/code], the Unix permissions for the destination path will be set to the provided value, if available on the current operating system.
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
//go:nosplit
func (self class) Copy(from gd.String, to gd.String, chmod_flags gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(from))
	callframe.Arg(frame, mmm.Get(to))
	callframe.Arg(frame, chmod_flags)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_copy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Static version of [method copy]. Supports only absolute paths.
*/
//go:nosplit
func (self class) CopyAbsolute(ctx gd.Lifetime, from gd.String, to gd.String, chmod_flags gd.Int) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(from))
	callframe.Arg(frame, mmm.Get(to))
	callframe.Arg(frame, chmod_flags)
	var r_ret = callframe.Ret[int64](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.DirAccess.Bind_copy_absolute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Renames (move) the [param from] file or directory to the [param to] destination. Both arguments should be paths to files or directories, either relative or absolute. If the destination file or directory exists and is not access-protected, it will be overwritten.
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
//go:nosplit
func (self class) Rename(from gd.String, to gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(from))
	callframe.Arg(frame, mmm.Get(to))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_rename, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Static version of [method rename]. Supports only absolute paths.
*/
//go:nosplit
func (self class) RenameAbsolute(ctx gd.Lifetime, from gd.String, to gd.String) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(from))
	callframe.Arg(frame, mmm.Get(to))
	var r_ret = callframe.Ret[int64](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.DirAccess.Bind_rename_absolute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Permanently deletes the target file or an empty directory. The argument can be relative to the current directory, or an absolute path. If the target directory is not empty, the operation will fail.
If you don't want to delete the file/directory permanently, use [method OS.move_to_trash] instead.
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
//go:nosplit
func (self class) Remove(path gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_remove, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Static version of [method remove]. Supports only absolute paths.
*/
//go:nosplit
func (self class) RemoveAbsolute(ctx gd.Lifetime, path gd.String) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.DirAccess.Bind_remove_absolute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the file or directory is a symbolic link, directory junction, or other reparse point.
[b]Note:[/b] This method is implemented on macOS, Linux, and Windows.
*/
//go:nosplit
func (self class) IsLink(path gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_is_link, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns target of the symbolic link.
[b]Note:[/b] This method is implemented on macOS, Linux, and Windows.
*/
//go:nosplit
func (self class) ReadLink(ctx gd.Lifetime, path gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_read_link, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Creates symbolic link between files or folders.
[b]Note:[/b] On Windows, this method works only if the application is running with elevated privileges or Developer Mode is enabled.
[b]Note:[/b] This method is implemented on macOS, Linux, and Windows.
*/
//go:nosplit
func (self class) CreateLink(source gd.String, target gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(source))
	callframe.Arg(frame, mmm.Get(target))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_create_link, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIncludeNavigational(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_set_include_navigational, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIncludeNavigational() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_get_include_navigational, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIncludeHidden(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_set_include_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIncludeHidden() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_get_include_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the file system or directory use case sensitive file names.
[b]Note:[/b] This method is implemented on macOS, Linux (for EXT4 and F2FS filesystems only) and Windows. On other platforms, it always returns [code]true[/code].
*/
//go:nosplit
func (self class) IsCaseSensitive(path gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.DirAccess.Bind_is_case_sensitive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsDirAccess() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsDirAccess() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("DirAccess", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
