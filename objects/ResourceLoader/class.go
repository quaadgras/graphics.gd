package ResourceLoader

import "unsafe"
import "sync"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
A singleton used to load resource files from the filesystem.
It uses the many [ResourceFormatLoader] classes registered in the engine (either built-in or from a plugin) to load files into memory and convert them to a format that can be used by the engine.
[b]Note:[/b] You have to import the files into the engine first to load them using [method load]. If you want to load [Image]s at run-time, you may use [method Image.load]. If you want to import audio files, you can use the snippet described in [member AudioStreamMP3.data].
*/
var self objects.ResourceLoader
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.ResourceLoader)
	self = *(*objects.ResourceLoader)(unsafe.Pointer(&obj))
}

/*
Loads the resource using threads. If [param use_sub_threads] is [code]true[/code], multiple threads will be used to load the resource, which makes loading faster, but may affect the main thread (and thus cause game slowdowns).
The [param cache_mode] property defines whether and how the cache should be used or updated when loading the resource. See [enum CacheMode] for details.
*/
func LoadThreadedRequest(path string) gd.Error {
	once.Do(singleton)
	return gd.Error(class(self).LoadThreadedRequest(gd.NewString(path), gd.NewString(""), false, 1))
}

/*
Returns the status of a threaded loading operation started with [method load_threaded_request] for the resource at [param path]. See [enum ThreadLoadStatus] for possible return values.
An array variable can optionally be passed via [param progress], and will return a one-element array containing the percentage of completion of the threaded loading.
[b]Note:[/b] The recommended way of using this method is to call it during different frames (e.g., in [method Node._process], instead of a loop).
*/
func LoadThreadedGetStatus(path string) classdb.ResourceLoaderThreadLoadStatus {
	once.Do(singleton)
	return classdb.ResourceLoaderThreadLoadStatus(class(self).LoadThreadedGetStatus(gd.NewString(path), ([1]gd.Array{}[0])))
}

/*
Returns the resource loaded by [method load_threaded_request].
If this is called before the loading thread is done (i.e. [method load_threaded_get_status] is not [constant THREAD_LOAD_LOADED]), the calling thread will be blocked until the resource has finished loading. However, it's recommended to use [method load_threaded_get_status] to known when the load has actually completed.
*/
func LoadThreadedGet(path string) objects.Resource {
	once.Do(singleton)
	return objects.Resource(class(self).LoadThreadedGet(gd.NewString(path)))
}

/*
Loads a resource at the given [param path], caching the result for further access.
The registered [ResourceFormatLoader]s are queried sequentially to find the first one which can handle the file's extension, and then attempt loading. If loading fails, the remaining ResourceFormatLoaders are also attempted.
An optional [param type_hint] can be used to further specify the [Resource] type that should be handled by the [ResourceFormatLoader]. Anything that inherits from [Resource] can be used as a type hint, for example [Image].
The [param cache_mode] property defines whether and how the cache should be used or updated when loading the resource. See [enum CacheMode] for details.
Returns an empty resource if no [ResourceFormatLoader] could handle the file, and prints an error if no file is found at the specified path.
GDScript has a simplified [method @GDScript.load] built-in method which can be used in most situations, leaving the use of [ResourceLoader] for more advanced scenarios.
[b]Note:[/b] If [member ProjectSettings.editor/export/convert_text_resources_to_binary] is [code]true[/code], [method @GDScript.load] will not be able to read converted files in an exported project. If you rely on run-time loading of files present within the PCK, set [member ProjectSettings.editor/export/convert_text_resources_to_binary] to [code]false[/code].
[b]Note:[/b] Relative paths will be prefixed with [code]"res://"[/code] before loading, to avoid unexpected results make sure your paths are absolute.
*/
func Load(path string) objects.Resource {
	once.Do(singleton)
	return objects.Resource(class(self).Load(gd.NewString(path), gd.NewString(""), 1))
}

/*
Returns the list of recognized extensions for a resource type.
*/
func GetRecognizedExtensionsForType(atype string) []string {
	once.Do(singleton)
	return []string(class(self).GetRecognizedExtensionsForType(gd.NewString(atype)).Strings())
}

/*
Registers a new [ResourceFormatLoader]. The ResourceLoader will use the ResourceFormatLoader as described in [method load].
This method is performed implicitly for ResourceFormatLoaders written in GDScript (see [ResourceFormatLoader] for more information).
*/
func AddResourceFormatLoader(format_loader objects.ResourceFormatLoader) {
	once.Do(singleton)
	class(self).AddResourceFormatLoader(format_loader, false)
}

/*
Unregisters the given [ResourceFormatLoader].
*/
func RemoveResourceFormatLoader(format_loader objects.ResourceFormatLoader) {
	once.Do(singleton)
	class(self).RemoveResourceFormatLoader(format_loader)
}

/*
Changes the behavior on missing sub-resources. The default behavior is to abort loading.
*/
func SetAbortOnMissingResources(abort bool) {
	once.Do(singleton)
	class(self).SetAbortOnMissingResources(abort)
}

/*
Returns the dependencies for the resource at the given [param path].
[b]Note:[/b] The dependencies are returned with slices separated by [code]::[/code]. You can use [method String.get_slice] to get their components.
[codeblock]
for dep in ResourceLoader.get_dependencies(path):

	print(dep.get_slice("::", 0)) # Prints UID.
	print(dep.get_slice("::", 2)) # Prints path.

[/codeblock]
*/
func GetDependencies(path string) []string {
	once.Do(singleton)
	return []string(class(self).GetDependencies(gd.NewString(path)).Strings())
}

/*
Returns whether a cached resource is available for the given [param path].
Once a resource has been loaded by the engine, it is cached in memory for faster access, and future calls to the [method load] method will use the cached version. The cached resource can be overridden by using [method Resource.take_over_path] on a new resource for that same path.
*/
func HasCached(path string) bool {
	once.Do(singleton)
	return bool(class(self).HasCached(gd.NewString(path)))
}

/*
Returns whether a recognized resource exists for the given [param path].
An optional [param type_hint] can be used to further specify the [Resource] type that should be handled by the [ResourceFormatLoader]. Anything that inherits from [Resource] can be used as a type hint, for example [Image].
[b]Note:[/b] If you use [method Resource.take_over_path], this method will return [code]true[/code] for the taken path even if the resource wasn't saved (i.e. exists only in resource cache).
*/
func Exists(path string) bool {
	once.Do(singleton)
	return bool(class(self).Exists(gd.NewString(path), gd.NewString("")))
}

/*
Returns the ID associated with a given resource path, or [code]-1[/code] when no such ID exists.
*/
func GetResourceUid(path string) int {
	once.Do(singleton)
	return int(int(class(self).GetResourceUid(gd.NewString(path))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]classdb.ResourceLoader

func (self class) AsObject() gd.Object { return self[0].AsObject() }

/*
Loads the resource using threads. If [param use_sub_threads] is [code]true[/code], multiple threads will be used to load the resource, which makes loading faster, but may affect the main thread (and thus cause game slowdowns).
The [param cache_mode] property defines whether and how the cache should be used or updated when loading the resource. See [enum CacheMode] for details.
*/
//go:nosplit
func (self class) LoadThreadedRequest(path gd.String, type_hint gd.String, use_sub_threads bool, cache_mode classdb.ResourceLoaderCacheMode) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, pointers.Get(type_hint))
	callframe.Arg(frame, use_sub_threads)
	callframe.Arg(frame, cache_mode)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_load_threaded_request, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the status of a threaded loading operation started with [method load_threaded_request] for the resource at [param path]. See [enum ThreadLoadStatus] for possible return values.
An array variable can optionally be passed via [param progress], and will return a one-element array containing the percentage of completion of the threaded loading.
[b]Note:[/b] The recommended way of using this method is to call it during different frames (e.g., in [method Node._process], instead of a loop).
*/
//go:nosplit
func (self class) LoadThreadedGetStatus(path gd.String, progress gd.Array) classdb.ResourceLoaderThreadLoadStatus {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, pointers.Get(progress))
	var r_ret = callframe.Ret[classdb.ResourceLoaderThreadLoadStatus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_load_threaded_get_status, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the resource loaded by [method load_threaded_request].
If this is called before the loading thread is done (i.e. [method load_threaded_get_status] is not [constant THREAD_LOAD_LOADED]), the calling thread will be blocked until the resource has finished loading. However, it's recommended to use [method load_threaded_get_status] to known when the load has actually completed.
*/
//go:nosplit
func (self class) LoadThreadedGet(path gd.String) objects.Resource {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_load_threaded_get, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Resource{classdb.Resource(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Loads a resource at the given [param path], caching the result for further access.
The registered [ResourceFormatLoader]s are queried sequentially to find the first one which can handle the file's extension, and then attempt loading. If loading fails, the remaining ResourceFormatLoaders are also attempted.
An optional [param type_hint] can be used to further specify the [Resource] type that should be handled by the [ResourceFormatLoader]. Anything that inherits from [Resource] can be used as a type hint, for example [Image].
The [param cache_mode] property defines whether and how the cache should be used or updated when loading the resource. See [enum CacheMode] for details.
Returns an empty resource if no [ResourceFormatLoader] could handle the file, and prints an error if no file is found at the specified path.
GDScript has a simplified [method @GDScript.load] built-in method which can be used in most situations, leaving the use of [ResourceLoader] for more advanced scenarios.
[b]Note:[/b] If [member ProjectSettings.editor/export/convert_text_resources_to_binary] is [code]true[/code], [method @GDScript.load] will not be able to read converted files in an exported project. If you rely on run-time loading of files present within the PCK, set [member ProjectSettings.editor/export/convert_text_resources_to_binary] to [code]false[/code].
[b]Note:[/b] Relative paths will be prefixed with [code]"res://"[/code] before loading, to avoid unexpected results make sure your paths are absolute.
*/
//go:nosplit
func (self class) Load(path gd.String, type_hint gd.String, cache_mode classdb.ResourceLoaderCacheMode) objects.Resource {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, pointers.Get(type_hint))
	callframe.Arg(frame, cache_mode)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_load, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Resource{classdb.Resource(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the list of recognized extensions for a resource type.
*/
//go:nosplit
func (self class) GetRecognizedExtensionsForType(atype gd.String) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(atype))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_get_recognized_extensions_for_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Registers a new [ResourceFormatLoader]. The ResourceLoader will use the ResourceFormatLoader as described in [method load].
This method is performed implicitly for ResourceFormatLoaders written in GDScript (see [ResourceFormatLoader] for more information).
*/
//go:nosplit
func (self class) AddResourceFormatLoader(format_loader objects.ResourceFormatLoader, at_front bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(format_loader[0])[0])
	callframe.Arg(frame, at_front)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_add_resource_format_loader, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Unregisters the given [ResourceFormatLoader].
*/
//go:nosplit
func (self class) RemoveResourceFormatLoader(format_loader objects.ResourceFormatLoader) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(format_loader[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_remove_resource_format_loader, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Changes the behavior on missing sub-resources. The default behavior is to abort loading.
*/
//go:nosplit
func (self class) SetAbortOnMissingResources(abort bool) {
	var frame = callframe.New()
	callframe.Arg(frame, abort)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_set_abort_on_missing_resources, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the dependencies for the resource at the given [param path].
[b]Note:[/b] The dependencies are returned with slices separated by [code]::[/code]. You can use [method String.get_slice] to get their components.
[codeblock]
for dep in ResourceLoader.get_dependencies(path):
    print(dep.get_slice("::", 0)) # Prints UID.
    print(dep.get_slice("::", 2)) # Prints path.
[/codeblock]
*/
//go:nosplit
func (self class) GetDependencies(path gd.String) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_get_dependencies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns whether a cached resource is available for the given [param path].
Once a resource has been loaded by the engine, it is cached in memory for faster access, and future calls to the [method load] method will use the cached version. The cached resource can be overridden by using [method Resource.take_over_path] on a new resource for that same path.
*/
//go:nosplit
func (self class) HasCached(path gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_has_cached, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether a recognized resource exists for the given [param path].
An optional [param type_hint] can be used to further specify the [Resource] type that should be handled by the [ResourceFormatLoader]. Anything that inherits from [Resource] can be used as a type hint, for example [Image].
[b]Note:[/b] If you use [method Resource.take_over_path], this method will return [code]true[/code] for the taken path even if the resource wasn't saved (i.e. exists only in resource cache).
*/
//go:nosplit
func (self class) Exists(path gd.String, type_hint gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, pointers.Get(type_hint))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_exists, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the ID associated with a given resource path, or [code]-1[/code] when no such ID exists.
*/
//go:nosplit
func (self class) GetResourceUid(path gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_get_resource_uid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("ResourceLoader", func(ptr gd.Object) any { return classdb.ResourceLoader(ptr) })
}

type ThreadLoadStatus = classdb.ResourceLoaderThreadLoadStatus

const (
	/*The resource is invalid, or has not been loaded with [method load_threaded_request].*/
	ThreadLoadInvalidResource ThreadLoadStatus = 0
	/*The resource is still being loaded.*/
	ThreadLoadInProgress ThreadLoadStatus = 1
	/*Some error occurred during loading and it failed.*/
	ThreadLoadFailed ThreadLoadStatus = 2
	/*The resource was loaded successfully and can be accessed via [method load_threaded_get].*/
	ThreadLoadLoaded ThreadLoadStatus = 3
)

type CacheMode = classdb.ResourceLoaderCacheMode

const (
	/*Neither the main resource (the one requested to be loaded) nor any of its subresources are retrieved from cache nor stored into it. Dependencies (external resources) are loaded with [constant CACHE_MODE_REUSE].*/
	CacheModeIgnore CacheMode = 0
	/*The main resource (the one requested to be loaded), its subresources, and its dependencies (external resources) are retrieved from cache if present, instead of loaded. Those not cached are loaded and then stored into the cache. The same rules are propagated recursively down the tree of dependencies (external resources).*/
	CacheModeReuse CacheMode = 1
	/*Like [constant CACHE_MODE_REUSE], but the cache is checked for the main resource (the one requested to be loaded) as well as for each of its subresources. Those already in the cache, as long as the loaded and cached types match, have their data refreshed from storage into the already existing instances. Otherwise, they are recreated as completely new objects.*/
	CacheModeReplace CacheMode = 2
	/*Like [constant CACHE_MODE_IGNORE], but propagated recursively down the tree of dependencies (external resources).*/
	CacheModeIgnoreDeep CacheMode = 3
	/*Like [constant CACHE_MODE_REPLACE], but propagated recursively down the tree of dependencies (external resources).*/
	CacheModeReplaceDeep CacheMode = 4
)
