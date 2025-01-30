// Package ResourceLoader provides methods for working with ResourceLoader object instances.
package ResourceLoader

import "unsafe"
import "sync"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

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
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
A singleton used to load resource files from the filesystem.
It uses the many [ResourceFormatLoader] classes registered in the engine (either built-in or from a plugin) to load files into memory and convert them to a format that can be used by the engine.
[b]Note:[/b] You have to import the files into the engine first to load them using [method load]. If you want to load [Image]s at run-time, you may use [method Image.load]. If you want to import audio files, you can use the snippet described in [member AudioStreamMP3.data].
*/
var self [1]gdclass.ResourceLoader
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.ResourceLoader)
	self = *(*[1]gdclass.ResourceLoader)(unsafe.Pointer(&obj))
}

/*
Loads the resource using threads. If [param use_sub_threads] is [code]true[/code], multiple threads will be used to load the resource, which makes loading faster, but may affect the main thread (and thus cause game slowdowns).
The [param cache_mode] property defines whether and how the cache should be used or updated when loading the resource. See [enum CacheMode] for details.
*/
func LoadThreadedRequest(path string) error { //gd:ResourceLoader.load_threaded_request
	once.Do(singleton)
	return error(gd.ToError(class(self).LoadThreadedRequest(String.New(path), String.New(""), false, 1)))
}

/*
Returns the status of a threaded loading operation started with [method load_threaded_request] for the resource at [param path]. See [enum ThreadLoadStatus] for possible return values.
An array variable can optionally be passed via [param progress], and will return a one-element array containing the percentage of completion of the threaded loading.
[b]Note:[/b] The recommended way of using this method is to call it during different frames (e.g., in [method Node._process], instead of a loop).
*/
func LoadThreadedGetStatus(path string) gdclass.ResourceLoaderThreadLoadStatus { //gd:ResourceLoader.load_threaded_get_status
	once.Do(singleton)
	return gdclass.ResourceLoaderThreadLoadStatus(class(self).LoadThreadedGetStatus(String.New(path), Array.Nil))
}

/*
Returns the resource loaded by [method load_threaded_request].
If this is called before the loading thread is done (i.e. [method load_threaded_get_status] is not [constant THREAD_LOAD_LOADED]), the calling thread will be blocked until the resource has finished loading. However, it's recommended to use [method load_threaded_get_status] to known when the load has actually completed.
*/
func LoadThreadedGet(path string) [1]gdclass.Resource { //gd:ResourceLoader.load_threaded_get
	once.Do(singleton)
	return [1]gdclass.Resource(class(self).LoadThreadedGet(String.New(path)))
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
func Load(path string) [1]gdclass.Resource { //gd:ResourceLoader.load
	once.Do(singleton)
	return [1]gdclass.Resource(class(self).Load(String.New(path), String.New(""), 1))
}

/*
Returns the list of recognized extensions for a resource type.
*/
func GetRecognizedExtensionsForType(atype string) []string { //gd:ResourceLoader.get_recognized_extensions_for_type
	once.Do(singleton)
	return []string(class(self).GetRecognizedExtensionsForType(String.New(atype)).Strings())
}

/*
Registers a new [ResourceFormatLoader]. The ResourceLoader will use the ResourceFormatLoader as described in [method load].
This method is performed implicitly for ResourceFormatLoaders written in GDScript (see [ResourceFormatLoader] for more information).
*/
func AddResourceFormatLoader(format_loader [1]gdclass.ResourceFormatLoader) { //gd:ResourceLoader.add_resource_format_loader
	once.Do(singleton)
	class(self).AddResourceFormatLoader(format_loader, false)
}

/*
Unregisters the given [ResourceFormatLoader].
*/
func RemoveResourceFormatLoader(format_loader [1]gdclass.ResourceFormatLoader) { //gd:ResourceLoader.remove_resource_format_loader
	once.Do(singleton)
	class(self).RemoveResourceFormatLoader(format_loader)
}

/*
Changes the behavior on missing sub-resources. The default behavior is to abort loading.
*/
func SetAbortOnMissingResources(abort bool) { //gd:ResourceLoader.set_abort_on_missing_resources
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
func GetDependencies(path string) []string { //gd:ResourceLoader.get_dependencies
	once.Do(singleton)
	return []string(class(self).GetDependencies(String.New(path)).Strings())
}

/*
Returns whether a cached resource is available for the given [param path].
Once a resource has been loaded by the engine, it is cached in memory for faster access, and future calls to the [method load] method will use the cached version. The cached resource can be overridden by using [method Resource.take_over_path] on a new resource for that same path.
*/
func HasCached(path string) bool { //gd:ResourceLoader.has_cached
	once.Do(singleton)
	return bool(class(self).HasCached(String.New(path)))
}

/*
Returns whether a recognized resource exists for the given [param path].
An optional [param type_hint] can be used to further specify the [Resource] type that should be handled by the [ResourceFormatLoader]. Anything that inherits from [Resource] can be used as a type hint, for example [Image].
[b]Note:[/b] If you use [method Resource.take_over_path], this method will return [code]true[/code] for the taken path even if the resource wasn't saved (i.e. exists only in resource cache).
*/
func Exists(path string) bool { //gd:ResourceLoader.exists
	once.Do(singleton)
	return bool(class(self).Exists(String.New(path), String.New("")))
}

/*
Returns the ID associated with a given resource path, or [code]-1[/code] when no such ID exists.
*/
func GetResourceUid(path string) int { //gd:ResourceLoader.get_resource_uid
	once.Do(singleton)
	return int(int(class(self).GetResourceUid(String.New(path))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.ResourceLoader

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Loads the resource using threads. If [param use_sub_threads] is [code]true[/code], multiple threads will be used to load the resource, which makes loading faster, but may affect the main thread (and thus cause game slowdowns).
The [param cache_mode] property defines whether and how the cache should be used or updated when loading the resource. See [enum CacheMode] for details.
*/
//go:nosplit
func (self class) LoadThreadedRequest(path String.Readable, type_hint String.Readable, use_sub_threads bool, cache_mode gdclass.ResourceLoaderCacheMode) Error.Code { //gd:ResourceLoader.load_threaded_request
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(type_hint)))
	callframe.Arg(frame, use_sub_threads)
	callframe.Arg(frame, cache_mode)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_load_threaded_request, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the status of a threaded loading operation started with [method load_threaded_request] for the resource at [param path]. See [enum ThreadLoadStatus] for possible return values.
An array variable can optionally be passed via [param progress], and will return a one-element array containing the percentage of completion of the threaded loading.
[b]Note:[/b] The recommended way of using this method is to call it during different frames (e.g., in [method Node._process], instead of a loop).
*/
//go:nosplit
func (self class) LoadThreadedGetStatus(path String.Readable, progress Array.Any) gdclass.ResourceLoaderThreadLoadStatus { //gd:ResourceLoader.load_threaded_get_status
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, pointers.Get(gd.InternalArray(progress)))
	var r_ret = callframe.Ret[gdclass.ResourceLoaderThreadLoadStatus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_load_threaded_get_status, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the resource loaded by [method load_threaded_request].
If this is called before the loading thread is done (i.e. [method load_threaded_get_status] is not [constant THREAD_LOAD_LOADED]), the calling thread will be blocked until the resource has finished loading. However, it's recommended to use [method load_threaded_get_status] to known when the load has actually completed.
*/
//go:nosplit
func (self class) LoadThreadedGet(path String.Readable) [1]gdclass.Resource { //gd:ResourceLoader.load_threaded_get
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_load_threaded_get, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Resource{gd.PointerWithOwnershipTransferredToGo[gdclass.Resource](r_ret.Get())}
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
func (self class) Load(path String.Readable, type_hint String.Readable, cache_mode gdclass.ResourceLoaderCacheMode) [1]gdclass.Resource { //gd:ResourceLoader.load
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(type_hint)))
	callframe.Arg(frame, cache_mode)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_load, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Resource{gd.PointerWithOwnershipTransferredToGo[gdclass.Resource](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the list of recognized extensions for a resource type.
*/
//go:nosplit
func (self class) GetRecognizedExtensionsForType(atype String.Readable) Packed.Strings { //gd:ResourceLoader.get_recognized_extensions_for_type
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(atype)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_get_recognized_extensions_for_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Registers a new [ResourceFormatLoader]. The ResourceLoader will use the ResourceFormatLoader as described in [method load].
This method is performed implicitly for ResourceFormatLoaders written in GDScript (see [ResourceFormatLoader] for more information).
*/
//go:nosplit
func (self class) AddResourceFormatLoader(format_loader [1]gdclass.ResourceFormatLoader, at_front bool) { //gd:ResourceLoader.add_resource_format_loader
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(format_loader[0])[0])
	callframe.Arg(frame, at_front)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_add_resource_format_loader, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Unregisters the given [ResourceFormatLoader].
*/
//go:nosplit
func (self class) RemoveResourceFormatLoader(format_loader [1]gdclass.ResourceFormatLoader) { //gd:ResourceLoader.remove_resource_format_loader
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(format_loader[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_remove_resource_format_loader, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Changes the behavior on missing sub-resources. The default behavior is to abort loading.
*/
//go:nosplit
func (self class) SetAbortOnMissingResources(abort bool) { //gd:ResourceLoader.set_abort_on_missing_resources
	var frame = callframe.New()
	callframe.Arg(frame, abort)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_set_abort_on_missing_resources, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) GetDependencies(path String.Readable) Packed.Strings { //gd:ResourceLoader.get_dependencies
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_get_dependencies, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns whether a cached resource is available for the given [param path].
Once a resource has been loaded by the engine, it is cached in memory for faster access, and future calls to the [method load] method will use the cached version. The cached resource can be overridden by using [method Resource.take_over_path] on a new resource for that same path.
*/
//go:nosplit
func (self class) HasCached(path String.Readable) bool { //gd:ResourceLoader.has_cached
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_has_cached, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) Exists(path String.Readable, type_hint String.Readable) bool { //gd:ResourceLoader.exists
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(type_hint)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_exists, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the ID associated with a given resource path, or [code]-1[/code] when no such ID exists.
*/
//go:nosplit
func (self class) GetResourceUid(path String.Readable) int64 { //gd:ResourceLoader.get_resource_uid
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceLoader.Bind_get_resource_uid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("ResourceLoader", func(ptr gd.Object) any {
		return [1]gdclass.ResourceLoader{*(*gdclass.ResourceLoader)(unsafe.Pointer(&ptr))}
	})
}

type ThreadLoadStatus = gdclass.ResourceLoaderThreadLoadStatus //gd:ResourceLoader.ThreadLoadStatus

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

type CacheMode = gdclass.ResourceLoaderCacheMode //gd:ResourceLoader.CacheMode

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
