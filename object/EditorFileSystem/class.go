package EditorFileSystem

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This object holds information of all resources in the filesystem, their types, etc.
[b]Note:[/b] This class shouldn't be instantiated directly. Instead, access the singleton using [method EditorInterface.get_resource_filesystem].

*/
type Simple [1]classdb.EditorFileSystem
func (self Simple) GetFilesystem() [1]classdb.EditorFileSystemDirectory {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.EditorFileSystemDirectory(Expert(self).GetFilesystem(gc))
}
func (self Simple) IsScanning() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsScanning())
}
func (self Simple) GetScanningProgress() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetScanningProgress()))
}
func (self Simple) Scan() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Scan()
}
func (self Simple) ScanSources() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ScanSources()
}
func (self Simple) UpdateFile(path string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UpdateFile(gc.String(path))
}
func (self Simple) GetFilesystemPath(path string) [1]classdb.EditorFileSystemDirectory {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.EditorFileSystemDirectory(Expert(self).GetFilesystemPath(gc, gc.String(path)))
}
func (self Simple) GetFileType(path string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetFileType(gc, gc.String(path)).String())
}
func (self Simple) ReimportFiles(files gd.PackedStringArray) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReimportFiles(files)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorFileSystem
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Gets the root directory object.
*/
//go:nosplit
func (self class) GetFilesystem(ctx gd.Lifetime) object.EditorFileSystemDirectory {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystem.Bind_get_filesystem, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.EditorFileSystemDirectory
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the filesystem is being scanned.
*/
//go:nosplit
func (self class) IsScanning() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystem.Bind_is_scanning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the scan progress for 0 to 1 if the FS is being scanned.
*/
//go:nosplit
func (self class) GetScanningProgress() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystem.Bind_get_scanning_progress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Scan the filesystem for changes.
*/
//go:nosplit
func (self class) Scan()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystem.Bind_scan, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Check if the source of any imported resource changed.
*/
//go:nosplit
func (self class) ScanSources()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystem.Bind_scan_sources, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Add a file in an existing directory, or schedule file information to be updated on editor restart. Can be used to update text files saved by an external program.
This will not import the file. To reimport, call [method reimport_files] or [method scan] methods.
*/
//go:nosplit
func (self class) UpdateFile(path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystem.Bind_update_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a view into the filesystem at [param path].
*/
//go:nosplit
func (self class) GetFilesystemPath(ctx gd.Lifetime, path gd.String) object.EditorFileSystemDirectory {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystem.Bind_get_filesystem_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.EditorFileSystemDirectory
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the resource type of the file, given the full path. This returns a string such as [code]"Resource"[/code] or [code]"GDScript"[/code], [i]not[/i] a file extension such as [code]".gd"[/code].
*/
//go:nosplit
func (self class) GetFileType(ctx gd.Lifetime, path gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystem.Bind_get_file_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Reimports a set of files. Call this if these files or their [code].import[/code] files were directly edited by script or an external program.
If the file type changed or the file was newly created, use [method update_file] or [method scan].
[b]Note:[/b] This function blocks until the import is finished. However, the main loop iteration, including timers and [method Node._process], will occur during the import process due to progress bar updates. Avoid calls to [method reimport_files] or [method scan] while an import is in progress.
*/
//go:nosplit
func (self class) ReimportFiles(files gd.PackedStringArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(files))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystem.Bind_reimport_files, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsEditorFileSystem() Expert { return self[0].AsEditorFileSystem() }


//go:nosplit
func (self Simple) AsEditorFileSystem() Simple { return self[0].AsEditorFileSystem() }


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
func init() {classdb.Register("EditorFileSystem", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
