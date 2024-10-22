package GDExtensionManager

import "unsafe"
import "sync"
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
The GDExtensionManager loads, initializes, and keeps track of all available [GDExtension] libraries in the project.
[b]Note:[/b] Do not worry about GDExtension unless you know what you are doing.

*/
var self gdclass.GDExtensionManager
var once sync.Once
func singleton() {
	gc := gd.Static
	obj := gc.API.Object.GetSingleton(gc, gc.API.Singletons.GDExtensionManager)
	self = *(*gdclass.GDExtensionManager)(unsafe.Pointer(&obj))
}

/*
Loads an extension by absolute file path. The [param path] needs to point to a valid [GDExtension]. Returns [constant LOAD_STATUS_OK] if successful.
*/
func LoadExtension(path string) classdb.GDExtensionManagerLoadStatus {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return classdb.GDExtensionManagerLoadStatus(class(self).LoadExtension(gc.String(path)))
}

/*
Reloads the extension at the given file path. The [param path] needs to point to a valid [GDExtension], otherwise this method may return either [constant LOAD_STATUS_NOT_LOADED] or [constant LOAD_STATUS_FAILED].
[b]Note:[/b] You can only reload extensions in the editor. In release builds, this method always fails and returns [constant LOAD_STATUS_FAILED].
*/
func ReloadExtension(path string) classdb.GDExtensionManagerLoadStatus {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return classdb.GDExtensionManagerLoadStatus(class(self).ReloadExtension(gc.String(path)))
}

/*
Unloads an extension by file path. The [param path] needs to point to an already loaded [GDExtension], otherwise this method returns [constant LOAD_STATUS_NOT_LOADED].
*/
func UnloadExtension(path string) classdb.GDExtensionManagerLoadStatus {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return classdb.GDExtensionManagerLoadStatus(class(self).UnloadExtension(gc.String(path)))
}

/*
Returns [code]true[/code] if the extension at the given file [param path] has already been loaded successfully. See also [method get_loaded_extensions].
*/
func IsExtensionLoaded(path string) bool {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return bool(class(self).IsExtensionLoaded(gc.String(path)))
}

/*
Returns the file paths of all currently loaded extensions.
*/
func GetLoadedExtensions() []string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return []string(class(self).GetLoadedExtensions(gc).Strings(gc))
}

/*
Returns the [GDExtension] at the given file [param path], or [code]null[/code] if it has not been loaded or does not exist.
*/
func GetExtension(path string) gdclass.GDExtension {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.GDExtension(class(self).GetExtension(gc, gc.String(path)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func GD() class { once.Do(singleton); return self }
type class [1]classdb.GDExtensionManager
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Loads an extension by absolute file path. The [param path] needs to point to a valid [GDExtension]. Returns [constant LOAD_STATUS_OK] if successful.
*/
//go:nosplit
func (self class) LoadExtension(path gd.String) classdb.GDExtensionManagerLoadStatus {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[classdb.GDExtensionManagerLoadStatus](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GDExtensionManager.Bind_load_extension, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Reloads the extension at the given file path. The [param path] needs to point to a valid [GDExtension], otherwise this method may return either [constant LOAD_STATUS_NOT_LOADED] or [constant LOAD_STATUS_FAILED].
[b]Note:[/b] You can only reload extensions in the editor. In release builds, this method always fails and returns [constant LOAD_STATUS_FAILED].
*/
//go:nosplit
func (self class) ReloadExtension(path gd.String) classdb.GDExtensionManagerLoadStatus {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[classdb.GDExtensionManagerLoadStatus](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GDExtensionManager.Bind_reload_extension, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Unloads an extension by file path. The [param path] needs to point to an already loaded [GDExtension], otherwise this method returns [constant LOAD_STATUS_NOT_LOADED].
*/
//go:nosplit
func (self class) UnloadExtension(path gd.String) classdb.GDExtensionManagerLoadStatus {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[classdb.GDExtensionManagerLoadStatus](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GDExtensionManager.Bind_unload_extension, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the extension at the given file [param path] has already been loaded successfully. See also [method get_loaded_extensions].
*/
//go:nosplit
func (self class) IsExtensionLoaded(path gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GDExtensionManager.Bind_is_extension_loaded, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the file paths of all currently loaded extensions.
*/
//go:nosplit
func (self class) GetLoadedExtensions(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GDExtensionManager.Bind_get_loaded_extensions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the [GDExtension] at the given file [param path], or [code]null[/code] if it has not been loaded or does not exist.
*/
//go:nosplit
func (self class) GetExtension(ctx gd.Lifetime, path gd.String) gdclass.GDExtension {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GDExtensionManager.Bind_get_extension, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.GDExtension
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func OnExtensionsReloaded(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("extensions_reloaded"), gc.Callable(cb), 0)
}


func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("GDExtensionManager", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type LoadStatus = classdb.GDExtensionManagerLoadStatus

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
