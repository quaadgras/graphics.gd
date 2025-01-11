// Package EditorFileSystem provides methods for working with EditorFileSystem object instances.
package EditorFileSystem

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
This object holds information of all resources in the filesystem, their types, etc.
[b]Note:[/b] This class shouldn't be instantiated directly. Instead, access the singleton using [method EditorInterface.get_resource_filesystem].
*/
type Instance [1]gdclass.EditorFileSystem

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorFileSystem() Instance
}

/*
Gets the root directory object.
*/
func (self Instance) GetFilesystem() [1]gdclass.EditorFileSystemDirectory {
	return [1]gdclass.EditorFileSystemDirectory(class(self).GetFilesystem())
}

/*
Returns [code]true[/code] if the filesystem is being scanned.
*/
func (self Instance) IsScanning() bool {
	return bool(class(self).IsScanning())
}

/*
Returns the scan progress for 0 to 1 if the FS is being scanned.
*/
func (self Instance) GetScanningProgress() Float.X {
	return Float.X(Float.X(class(self).GetScanningProgress()))
}

/*
Scan the filesystem for changes.
*/
func (self Instance) Scan() {
	class(self).Scan()
}

/*
Check if the source of any imported resource changed.
*/
func (self Instance) ScanSources() {
	class(self).ScanSources()
}

/*
Add a file in an existing directory, or schedule file information to be updated on editor restart. Can be used to update text files saved by an external program.
This will not import the file. To reimport, call [method reimport_files] or [method scan] methods.
*/
func (self Instance) UpdateFile(path string) {
	class(self).UpdateFile(gd.NewString(path))
}

/*
Returns a view into the filesystem at [param path].
*/
func (self Instance) GetFilesystemPath(path string) [1]gdclass.EditorFileSystemDirectory {
	return [1]gdclass.EditorFileSystemDirectory(class(self).GetFilesystemPath(gd.NewString(path)))
}

/*
Returns the resource type of the file, given the full path. This returns a string such as [code]"Resource"[/code] or [code]"GDScript"[/code], [i]not[/i] a file extension such as [code]".gd"[/code].
*/
func (self Instance) GetFileType(path string) string {
	return string(class(self).GetFileType(gd.NewString(path)).String())
}

/*
Reimports a set of files. Call this if these files or their [code].import[/code] files were directly edited by script or an external program.
If the file type changed or the file was newly created, use [method update_file] or [method scan].
[b]Note:[/b] This function blocks until the import is finished. However, the main loop iteration, including timers and [method Node._process], will occur during the import process due to progress bar updates. Avoid calls to [method reimport_files] or [method scan] while an import is in progress.
*/
func (self Instance) ReimportFiles(files []string) {
	class(self).ReimportFiles(gd.NewPackedStringSlice(files))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorFileSystem

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorFileSystem"))
	casted := Instance{*(*gdclass.EditorFileSystem)(unsafe.Pointer(&object))}
	return casted
}

/*
Gets the root directory object.
*/
//go:nosplit
func (self class) GetFilesystem() [1]gdclass.EditorFileSystemDirectory {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystem.Bind_get_filesystem, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.EditorFileSystemDirectory{gd.PointerLifetimeBoundTo[gdclass.EditorFileSystemDirectory](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the filesystem is being scanned.
*/
//go:nosplit
func (self class) IsScanning() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystem.Bind_is_scanning, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the scan progress for 0 to 1 if the FS is being scanned.
*/
//go:nosplit
func (self class) GetScanningProgress() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystem.Bind_get_scanning_progress, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Scan the filesystem for changes.
*/
//go:nosplit
func (self class) Scan() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystem.Bind_scan, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Check if the source of any imported resource changed.
*/
//go:nosplit
func (self class) ScanSources() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystem.Bind_scan_sources, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Add a file in an existing directory, or schedule file information to be updated on editor restart. Can be used to update text files saved by an external program.
This will not import the file. To reimport, call [method reimport_files] or [method scan] methods.
*/
//go:nosplit
func (self class) UpdateFile(path gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystem.Bind_update_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a view into the filesystem at [param path].
*/
//go:nosplit
func (self class) GetFilesystemPath(path gd.String) [1]gdclass.EditorFileSystemDirectory {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystem.Bind_get_filesystem_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.EditorFileSystemDirectory{gd.PointerLifetimeBoundTo[gdclass.EditorFileSystemDirectory](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the resource type of the file, given the full path. This returns a string such as [code]"Resource"[/code] or [code]"GDScript"[/code], [i]not[/i] a file extension such as [code]".gd"[/code].
*/
//go:nosplit
func (self class) GetFileType(path gd.String) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystem.Bind_get_file_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Reimports a set of files. Call this if these files or their [code].import[/code] files were directly edited by script or an external program.
If the file type changed or the file was newly created, use [method update_file] or [method scan].
[b]Note:[/b] This function blocks until the import is finished. However, the main loop iteration, including timers and [method Node._process], will occur during the import process due to progress bar updates. Avoid calls to [method reimport_files] or [method scan] while an import is in progress.
*/
//go:nosplit
func (self class) ReimportFiles(files gd.PackedStringArray) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(files))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystem.Bind_reimport_files, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnFilesystemChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("filesystem_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnScriptClassesUpdated(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("script_classes_updated"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSourcesChanged(cb func(exist bool)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("sources_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnResourcesReimporting(cb func(resources []string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("resources_reimporting"), gd.NewCallable(cb), 0)
}

func (self Instance) OnResourcesReimported(cb func(resources []string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("resources_reimported"), gd.NewCallable(cb), 0)
}

func (self Instance) OnResourcesReload(cb func(resources []string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("resources_reload"), gd.NewCallable(cb), 0)
}

func (self class) AsEditorFileSystem() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorFileSystem() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced           { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance        { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Advanced(self.AsNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Instance(self.AsNode()), name)
	}
}
func init() {
	gdclass.Register("EditorFileSystem", func(ptr gd.Object) any {
		return [1]gdclass.EditorFileSystem{*(*gdclass.EditorFileSystem)(unsafe.Pointer(&ptr))}
	})
}
