// Package GDExtensionManager provides methods for working with GDExtensionManager object instances.
package GDExtensionManager

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The GDExtensionManager loads, initializes, and keeps track of all available [GDExtension] libraries in the project.
[b]Note:[/b] Do not worry about GDExtension unless you know what you are doing.
*/
var self [1]gdclass.GDExtensionManager
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.GDExtensionManager)
	self = *(*[1]gdclass.GDExtensionManager)(unsafe.Pointer(&obj))
}

/*
Loads an extension by absolute file path. The [param path] needs to point to a valid [GDExtension]. Returns [constant LOAD_STATUS_OK] if successful.
*/
func LoadExtension(path string) gdclass.GDExtensionManagerLoadStatus {
	once.Do(singleton)
	return gdclass.GDExtensionManagerLoadStatus(class(self).LoadExtension(gd.NewString(path)))
}

/*
Reloads the extension at the given file path. The [param path] needs to point to a valid [GDExtension], otherwise this method may return either [constant LOAD_STATUS_NOT_LOADED] or [constant LOAD_STATUS_FAILED].
[b]Note:[/b] You can only reload extensions in the editor. In release builds, this method always fails and returns [constant LOAD_STATUS_FAILED].
*/
func ReloadExtension(path string) gdclass.GDExtensionManagerLoadStatus {
	once.Do(singleton)
	return gdclass.GDExtensionManagerLoadStatus(class(self).ReloadExtension(gd.NewString(path)))
}

/*
Unloads an extension by file path. The [param path] needs to point to an already loaded [GDExtension], otherwise this method returns [constant LOAD_STATUS_NOT_LOADED].
*/
func UnloadExtension(path string) gdclass.GDExtensionManagerLoadStatus {
	once.Do(singleton)
	return gdclass.GDExtensionManagerLoadStatus(class(self).UnloadExtension(gd.NewString(path)))
}

/*
Returns [code]true[/code] if the extension at the given file [param path] has already been loaded successfully. See also [method get_loaded_extensions].
*/
func IsExtensionLoaded(path string) bool {
	once.Do(singleton)
	return bool(class(self).IsExtensionLoaded(gd.NewString(path)))
}

/*
Returns the file paths of all currently loaded extensions.
*/
func GetLoadedExtensions() []string {
	once.Do(singleton)
	return []string(class(self).GetLoadedExtensions().Strings())
}

/*
Returns the [GDExtension] at the given file [param path], or [code]null[/code] if it has not been loaded or does not exist.
*/
func GetExtension(path string) [1]gdclass.GDExtension {
	once.Do(singleton)
	return [1]gdclass.GDExtension(class(self).GetExtension(gd.NewString(path)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.GDExtensionManager

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Loads an extension by absolute file path. The [param path] needs to point to a valid [GDExtension]. Returns [constant LOAD_STATUS_OK] if successful.
*/
//go:nosplit
func (self class) LoadExtension(path gd.String) gdclass.GDExtensionManagerLoadStatus {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[gdclass.GDExtensionManagerLoadStatus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GDExtensionManager.Bind_load_extension, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Reloads the extension at the given file path. The [param path] needs to point to a valid [GDExtension], otherwise this method may return either [constant LOAD_STATUS_NOT_LOADED] or [constant LOAD_STATUS_FAILED].
[b]Note:[/b] You can only reload extensions in the editor. In release builds, this method always fails and returns [constant LOAD_STATUS_FAILED].
*/
//go:nosplit
func (self class) ReloadExtension(path gd.String) gdclass.GDExtensionManagerLoadStatus {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[gdclass.GDExtensionManagerLoadStatus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GDExtensionManager.Bind_reload_extension, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Unloads an extension by file path. The [param path] needs to point to an already loaded [GDExtension], otherwise this method returns [constant LOAD_STATUS_NOT_LOADED].
*/
//go:nosplit
func (self class) UnloadExtension(path gd.String) gdclass.GDExtensionManagerLoadStatus {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[gdclass.GDExtensionManagerLoadStatus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GDExtensionManager.Bind_unload_extension, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the extension at the given file [param path] has already been loaded successfully. See also [method get_loaded_extensions].
*/
//go:nosplit
func (self class) IsExtensionLoaded(path gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GDExtensionManager.Bind_is_extension_loaded, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the file paths of all currently loaded extensions.
*/
//go:nosplit
func (self class) GetLoadedExtensions() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GDExtensionManager.Bind_get_loaded_extensions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the [GDExtension] at the given file [param path], or [code]null[/code] if it has not been loaded or does not exist.
*/
//go:nosplit
func (self class) GetExtension(path gd.String) [1]gdclass.GDExtension {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GDExtensionManager.Bind_get_extension, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.GDExtension{gd.PointerWithOwnershipTransferredToGo[gdclass.GDExtension](r_ret.Get())}
	frame.Free()
	return ret
}
func OnExtensionsReloaded(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("extensions_reloaded"), gd.NewCallable(cb), 0)
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("GDExtensionManager", func(ptr gd.Object) any {
		return [1]gdclass.GDExtensionManager{*(*gdclass.GDExtensionManager)(unsafe.Pointer(&ptr))}
	})
}

type LoadStatus = gdclass.GDExtensionManagerLoadStatus

const (
	/*The extension has loaded successfully.*/
	LoadStatusOk LoadStatus = 0
	/*The extension has failed to load, possibly because it does not exist or has missing dependencies.*/
	LoadStatusFailed LoadStatus = 1
	/*The extension has already been loaded.*/
	LoadStatusAlreadyLoaded LoadStatus = 2
	/*The extension has not been loaded.*/
	LoadStatusNotLoaded LoadStatus = 3
	/*The extension requires the application to restart to fully load.*/
	LoadStatusNeedsRestart LoadStatus = 4
)
