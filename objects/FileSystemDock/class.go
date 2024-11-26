package FileSystemDock

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/VBoxContainer"
import "grow.graphics/gd/objects/BoxContainer"
import "grow.graphics/gd/objects/Container"
import "grow.graphics/gd/objects/Control"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This class is available only in [EditorPlugin]s and can't be instantiated. You can access it using [method EditorInterface.get_file_system_dock].
While [FileSystemDock] doesn't expose any methods for file manipulation, it can listen for various file-related signals.
*/
type Instance [1]classdb.FileSystemDock

/*
Sets the given [param path] as currently selected, ensuring that the selected file/directory is visible.
*/
func (self Instance) NavigateToPath(path string) {
	class(self).NavigateToPath(gd.NewString(path))
}

/*
Registers a new [EditorResourceTooltipPlugin].
*/
func (self Instance) AddResourceTooltipPlugin(plugin objects.EditorResourceTooltipPlugin) {
	class(self).AddResourceTooltipPlugin(plugin)
}

/*
Removes an [EditorResourceTooltipPlugin]. Fails if the plugin wasn't previously added.
*/
func (self Instance) RemoveResourceTooltipPlugin(plugin objects.EditorResourceTooltipPlugin) {
	class(self).RemoveResourceTooltipPlugin(plugin)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.FileSystemDock

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FileSystemDock"))
	return Instance{classdb.FileSystemDock(object)}
}

/*
Sets the given [param path] as currently selected, ensuring that the selected file/directory is visible.
*/
//go:nosplit
func (self class) NavigateToPath(path gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileSystemDock.Bind_navigate_to_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Registers a new [EditorResourceTooltipPlugin].
*/
//go:nosplit
func (self class) AddResourceTooltipPlugin(plugin objects.EditorResourceTooltipPlugin) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileSystemDock.Bind_add_resource_tooltip_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes an [EditorResourceTooltipPlugin]. Fails if the plugin wasn't previously added.
*/
//go:nosplit
func (self class) RemoveResourceTooltipPlugin(plugin objects.EditorResourceTooltipPlugin) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileSystemDock.Bind_remove_resource_tooltip_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Instance) OnInherit(cb func(file string)) {
	self[0].AsObject().Connect(gd.NewStringName("inherit"), gd.NewCallable(cb), 0)
}

func (self Instance) OnInstantiate(cb func(files []string)) {
	self[0].AsObject().Connect(gd.NewStringName("instantiate"), gd.NewCallable(cb), 0)
}

func (self Instance) OnResourceRemoved(cb func(resource objects.Resource)) {
	self[0].AsObject().Connect(gd.NewStringName("resource_removed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnFileRemoved(cb func(file string)) {
	self[0].AsObject().Connect(gd.NewStringName("file_removed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnFolderRemoved(cb func(folder string)) {
	self[0].AsObject().Connect(gd.NewStringName("folder_removed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnFilesMoved(cb func(old_file string, new_file string)) {
	self[0].AsObject().Connect(gd.NewStringName("files_moved"), gd.NewCallable(cb), 0)
}

func (self Instance) OnFolderMoved(cb func(old_folder string, new_folder string)) {
	self[0].AsObject().Connect(gd.NewStringName("folder_moved"), gd.NewCallable(cb), 0)
}

func (self Instance) OnFolderColorChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("folder_color_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnDisplayModeChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("display_mode_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsFileSystemDock() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFileSystemDock() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsVBoxContainer() VBoxContainer.Advanced {
	return *((*VBoxContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVBoxContainer() VBoxContainer.Instance {
	return *((*VBoxContainer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsBoxContainer() BoxContainer.Advanced {
	return *((*BoxContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsBoxContainer() BoxContainer.Instance {
	return *((*BoxContainer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsContainer() Container.Advanced {
	return *((*Container.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsContainer() Container.Instance {
	return *((*Container.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsControl() Control.Advanced { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVBoxContainer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVBoxContainer(), name)
	}
}
func init() {
	classdb.Register("FileSystemDock", func(ptr gd.Object) any { return classdb.FileSystemDock(ptr) })
}
