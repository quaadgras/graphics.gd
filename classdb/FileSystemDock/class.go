// Package FileSystemDock provides methods for working with FileSystemDock object instances.
package FileSystemDock

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
import "graphics.gd/classdb/VBoxContainer"
import "graphics.gd/classdb/BoxContainer"
import "graphics.gd/classdb/Container"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"

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

/*
This class is available only in [EditorPlugin]s and can't be instantiated. You can access it using [method EditorInterface.get_file_system_dock].
While [FileSystemDock] doesn't expose any methods for file manipulation, it can listen for various file-related signals.
*/
type Instance [1]gdclass.FileSystemDock

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsFileSystemDock() Instance
}

/*
Sets the given [param path] as currently selected, ensuring that the selected file/directory is visible.
*/
func (self Instance) NavigateToPath(path string) { //gd:FileSystemDock.navigate_to_path
	class(self).NavigateToPath(gd.NewString(path))
}

/*
Registers a new [EditorResourceTooltipPlugin].
*/
func (self Instance) AddResourceTooltipPlugin(plugin [1]gdclass.EditorResourceTooltipPlugin) { //gd:FileSystemDock.add_resource_tooltip_plugin
	class(self).AddResourceTooltipPlugin(plugin)
}

/*
Removes an [EditorResourceTooltipPlugin]. Fails if the plugin wasn't previously added.
*/
func (self Instance) RemoveResourceTooltipPlugin(plugin [1]gdclass.EditorResourceTooltipPlugin) { //gd:FileSystemDock.remove_resource_tooltip_plugin
	class(self).RemoveResourceTooltipPlugin(plugin)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.FileSystemDock

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FileSystemDock"))
	casted := Instance{*(*gdclass.FileSystemDock)(unsafe.Pointer(&object))}
	return casted
}

/*
Sets the given [param path] as currently selected, ensuring that the selected file/directory is visible.
*/
//go:nosplit
func (self class) NavigateToPath(path gd.String) { //gd:FileSystemDock.navigate_to_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileSystemDock.Bind_navigate_to_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Registers a new [EditorResourceTooltipPlugin].
*/
//go:nosplit
func (self class) AddResourceTooltipPlugin(plugin [1]gdclass.EditorResourceTooltipPlugin) { //gd:FileSystemDock.add_resource_tooltip_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileSystemDock.Bind_add_resource_tooltip_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes an [EditorResourceTooltipPlugin]. Fails if the plugin wasn't previously added.
*/
//go:nosplit
func (self class) RemoveResourceTooltipPlugin(plugin [1]gdclass.EditorResourceTooltipPlugin) { //gd:FileSystemDock.remove_resource_tooltip_plugin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileSystemDock.Bind_remove_resource_tooltip_plugin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnInherit(cb func(file string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("inherit"), gd.NewCallable(cb), 0)
}

func (self Instance) OnInstantiate(cb func(files []string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("instantiate"), gd.NewCallable(cb), 0)
}

func (self Instance) OnResourceRemoved(cb func(resource [1]gdclass.Resource)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("resource_removed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnFileRemoved(cb func(file string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("file_removed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnFolderRemoved(cb func(folder string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("folder_removed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnFilesMoved(cb func(old_file string, new_file string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("files_moved"), gd.NewCallable(cb), 0)
}

func (self Instance) OnFolderMoved(cb func(old_folder string, new_folder string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("folder_moved"), gd.NewCallable(cb), 0)
}

func (self Instance) OnFolderColorChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("folder_color_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnDisplayModeChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("display_mode_changed"), gd.NewCallable(cb), 0)
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
		return gd.VirtualByName(VBoxContainer.Advanced(self.AsVBoxContainer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VBoxContainer.Instance(self.AsVBoxContainer()), name)
	}
}
func init() {
	gdclass.Register("FileSystemDock", func(ptr gd.Object) any {
		return [1]gdclass.FileSystemDock{*(*gdclass.FileSystemDock)(unsafe.Pointer(&ptr))}
	})
}
