package FileSystemDock

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VBoxContainer"
import "grow.graphics/gd/gdclass/BoxContainer"
import "grow.graphics/gd/gdclass/Container"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
This class is available only in [EditorPlugin]s and can't be instantiated. You can access it using [method EditorInterface.get_file_system_dock].
While [FileSystemDock] doesn't expose any methods for file manipulation, it can listen for various file-related signals.

*/
type Go [1]classdb.FileSystemDock

/*
Sets the given [param path] as currently selected, ensuring that the selected file/directory is visible.
*/
func (self Go) NavigateToPath(path string) {
	class(self).NavigateToPath(gd.NewString(path))
}

/*
Registers a new [EditorResourceTooltipPlugin].
*/
func (self Go) AddResourceTooltipPlugin(plugin gdclass.EditorResourceTooltipPlugin) {
	class(self).AddResourceTooltipPlugin(plugin)
}

/*
Removes an [EditorResourceTooltipPlugin]. Fails if the plugin wasn't previously added.
*/
func (self Go) RemoveResourceTooltipPlugin(plugin gdclass.EditorResourceTooltipPlugin) {
	class(self).RemoveResourceTooltipPlugin(plugin)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.FileSystemDock
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FileSystemDock"))
	return Go{classdb.FileSystemDock(object)}
}

/*
Sets the given [param path] as currently selected, ensuring that the selected file/directory is visible.
*/
//go:nosplit
func (self class) NavigateToPath(path gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileSystemDock.Bind_navigate_to_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Registers a new [EditorResourceTooltipPlugin].
*/
//go:nosplit
func (self class) AddResourceTooltipPlugin(plugin gdclass.EditorResourceTooltipPlugin)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(plugin[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileSystemDock.Bind_add_resource_tooltip_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes an [EditorResourceTooltipPlugin]. Fails if the plugin wasn't previously added.
*/
//go:nosplit
func (self class) RemoveResourceTooltipPlugin(plugin gdclass.EditorResourceTooltipPlugin)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(plugin[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileSystemDock.Bind_remove_resource_tooltip_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnInherit(cb func(file string)) {
	self[0].AsObject().Connect(gd.NewStringName("inherit"), gd.NewCallable(cb), 0)
}


func (self Go) OnInstantiate(cb func(files []string)) {
	self[0].AsObject().Connect(gd.NewStringName("instantiate"), gd.NewCallable(cb), 0)
}


func (self Go) OnResourceRemoved(cb func(resource gdclass.Resource)) {
	self[0].AsObject().Connect(gd.NewStringName("resource_removed"), gd.NewCallable(cb), 0)
}


func (self Go) OnFileRemoved(cb func(file string)) {
	self[0].AsObject().Connect(gd.NewStringName("file_removed"), gd.NewCallable(cb), 0)
}


func (self Go) OnFolderRemoved(cb func(folder string)) {
	self[0].AsObject().Connect(gd.NewStringName("folder_removed"), gd.NewCallable(cb), 0)
}


func (self Go) OnFilesMoved(cb func(old_file string, new_file string)) {
	self[0].AsObject().Connect(gd.NewStringName("files_moved"), gd.NewCallable(cb), 0)
}


func (self Go) OnFolderMoved(cb func(old_folder string, new_folder string)) {
	self[0].AsObject().Connect(gd.NewStringName("folder_moved"), gd.NewCallable(cb), 0)
}


func (self Go) OnFolderColorChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("folder_color_changed"), gd.NewCallable(cb), 0)
}


func (self Go) OnDisplayModeChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("display_mode_changed"), gd.NewCallable(cb), 0)
}


func (self class) AsFileSystemDock() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsFileSystemDock() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVBoxContainer() VBoxContainer.GD { return *((*VBoxContainer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVBoxContainer() VBoxContainer.Go { return *((*VBoxContainer.Go)(unsafe.Pointer(&self))) }
func (self class) AsBoxContainer() BoxContainer.GD { return *((*BoxContainer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsBoxContainer() BoxContainer.Go { return *((*BoxContainer.Go)(unsafe.Pointer(&self))) }
func (self class) AsContainer() Container.GD { return *((*Container.GD)(unsafe.Pointer(&self))) }
func (self Go) AsContainer() Container.Go { return *((*Container.Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVBoxContainer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVBoxContainer(), name)
	}
}
func init() {classdb.Register("FileSystemDock", func(ptr gd.Object) any { return classdb.FileSystemDock(ptr) })}
