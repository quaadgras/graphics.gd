package ResourceFormatLoader

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Godot loads resources in the editor or in exported games using ResourceFormatLoaders. They are queried automatically via the [ResourceLoader] singleton, or when a resource with internal dependencies is loaded. Each file type may load as a different resource type, so multiple ResourceFormatLoaders are registered in the engine.
Extending this class allows you to define your own loader. Be sure to respect the documented return types and values. You should give it a global class name with [code]class_name[/code] for it to be registered. Like built-in ResourceFormatLoaders, it will be called automatically when loading resources of its handled type(s). You may also implement a [ResourceFormatSaver].
[b]Note:[/b] You can also extend [EditorImportPlugin] if the resource type you need exists but Godot is unable to load its format. Choosing one way over another depends on if the format is suitable or not for the final exported game. For example, it's better to import [code].png[/code] textures as [code].ctex[/code] ([CompressedTexture2D]) first, so they can be loaded with better efficiency on the graphics card.

	// ResourceFormatLoader methods that can be overridden by a [Class] that extends it.
	type ResourceFormatLoader interface {
		//Gets the list of extensions for files this loader is able to read.
		GetRecognizedExtensions() []string
		//Tells whether or not this loader should load a resource from its resource path for a given type.
		//If it is not implemented, the default behavior returns whether the path's extension is within the ones provided by [method _get_recognized_extensions], and if the type is within the ones provided by [method _get_resource_type].
		RecognizePath(path string, atype string) bool
		//Tells which resource class this loader can load.
		//[b]Note:[/b] Custom resource types defined by scripts aren't known by the [ClassDB], so you might just handle [code]"Resource"[/code] for them.
		HandlesType(atype string) bool
		//Gets the class name of the resource associated with the given path. If the loader cannot handle it, it should return [code]""[/code].
		//[b]Note:[/b] Custom resource types defined by scripts aren't known by the [ClassDB], so you might just return [code]"Resource"[/code] for them.
		GetResourceType(path string) string
		//Returns the script class name associated with the [Resource] under the given [param path]. If the resource has no script or the script isn't a named class, it should return [code]""[/code].
		GetResourceScriptClass(path string) string
		GetResourceUid(path string) int
		//If implemented, gets the dependencies of a given resource. If [param add_types] is [code]true[/code], paths should be appended [code]::TypeName[/code], where [code]TypeName[/code] is the class name of the dependency.
		//[b]Note:[/b] Custom resource types defined by scripts aren't known by the [ClassDB], so you might just return [code]"Resource"[/code] for them.
		GetDependencies(path string, add_types bool) []string
		//If implemented, renames dependencies within the given resource and saves it. [param renames] is a dictionary [code]{ String => String }[/code] mapping old dependency paths to new paths.
		//Returns [constant OK] on success, or an [enum Error] constant in case of failure.
		RenameDependencies(path string, renames gd.Dictionary) gd.Error
		Exists(path string) bool
		GetClassesUsed(path string) []string
		//Loads a resource when the engine finds this loader to be compatible. If the loaded resource is the result of an import, [param original_path] will target the source file. Returns a [Resource] object on success, or an [enum Error] constant in case of failure.
		//The [param cache_mode] property defines whether and how the cache should be used or updated when loading the resource. See [enum CacheMode] for details.
		Load(path string, original_path string, use_sub_threads bool, cache_mode int) gd.Variant
	}
*/
type Instance [1]classdb.ResourceFormatLoader

/*
Gets the list of extensions for files this loader is able to read.
*/
func (Instance) _get_recognized_extensions(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Tells whether or not this loader should load a resource from its resource path for a given type.
If it is not implemented, the default behavior returns whether the path's extension is within the ones provided by [method _get_recognized_extensions], and if the type is within the ones provided by [method _get_resource_type].
*/
func (Instance) _recognize_path(impl func(ptr unsafe.Pointer, path string, atype string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(path)
		var atype = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(atype)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), atype.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Tells which resource class this loader can load.
[b]Note:[/b] Custom resource types defined by scripts aren't known by the [ClassDB], so you might just handle [code]"Resource"[/code] for them.
*/
func (Instance) _handles_type(impl func(ptr unsafe.Pointer, atype string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var atype = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(atype)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, atype.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Gets the class name of the resource associated with the given path. If the loader cannot handle it, it should return [code]""[/code].
[b]Note:[/b] Custom resource types defined by scripts aren't known by the [ClassDB], so you might just return [code]"Resource"[/code] for them.
*/
func (Instance) _get_resource_type(impl func(ptr unsafe.Pointer, path string) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(path)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String())
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns the script class name associated with the [Resource] under the given [param path]. If the resource has no script or the script isn't a named class, it should return [code]""[/code].
*/
func (Instance) _get_resource_script_class(impl func(ptr unsafe.Pointer, path string) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(path)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String())
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_resource_uid(impl func(ptr unsafe.Pointer, path string) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(path)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String())
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
If implemented, gets the dependencies of a given resource. If [param add_types] is [code]true[/code], paths should be appended [code]::TypeName[/code], where [code]TypeName[/code] is the class name of the dependency.
[b]Note:[/b] Custom resource types defined by scripts aren't known by the [ClassDB], so you might just return [code]"Resource"[/code] for them.
*/
func (Instance) _get_dependencies(impl func(ptr unsafe.Pointer, path string, add_types bool) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(path)
		var add_types = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), add_types)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
If implemented, renames dependencies within the given resource and saves it. [param renames] is a dictionary [code]{ String => String }[/code] mapping old dependency paths to new paths.
Returns [constant OK] on success, or an [enum Error] constant in case of failure.
*/
func (Instance) _rename_dependencies(impl func(ptr unsafe.Pointer, path string, renames gd.Dictionary) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(path)
		var renames = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(renames)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), renames)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _exists(impl func(ptr unsafe.Pointer, path string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(path)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String())
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_classes_used(impl func(ptr unsafe.Pointer, path string) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(path)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String())
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Loads a resource when the engine finds this loader to be compatible. If the loaded resource is the result of an import, [param original_path] will target the source file. Returns a [Resource] object on success, or an [enum Error] constant in case of failure.
The [param cache_mode] property defines whether and how the cache should be used or updated when loading the resource. See [enum CacheMode] for details.
*/
func (Instance) _load(impl func(ptr unsafe.Pointer, path string, original_path string, use_sub_threads bool, cache_mode int) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(path)
		var original_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(original_path)
		var use_sub_threads = gd.UnsafeGet[bool](p_args, 2)
		var cache_mode = gd.UnsafeGet[gd.Int](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), original_path.String(), use_sub_threads, int(cache_mode))
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ResourceFormatLoader

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ResourceFormatLoader"))
	return Instance{classdb.ResourceFormatLoader(object)}
}

/*
Gets the list of extensions for files this loader is able to read.
*/
func (class) _get_recognized_extensions(impl func(ptr unsafe.Pointer) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Tells whether or not this loader should load a resource from its resource path for a given type.
If it is not implemented, the default behavior returns whether the path's extension is within the ones provided by [method _get_recognized_extensions], and if the type is within the ones provided by [method _get_resource_type].
*/
func (class) _recognize_path(impl func(ptr unsafe.Pointer, path gd.String, atype gd.StringName) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var atype = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, atype)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Tells which resource class this loader can load.
[b]Note:[/b] Custom resource types defined by scripts aren't known by the [ClassDB], so you might just handle [code]"Resource"[/code] for them.
*/
func (class) _handles_type(impl func(ptr unsafe.Pointer, atype gd.StringName) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var atype = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, atype)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Gets the class name of the resource associated with the given path. If the loader cannot handle it, it should return [code]""[/code].
[b]Note:[/b] Custom resource types defined by scripts aren't known by the [ClassDB], so you might just return [code]"Resource"[/code] for them.
*/
func (class) _get_resource_type(impl func(ptr unsafe.Pointer, path gd.String) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns the script class name associated with the [Resource] under the given [param path]. If the resource has no script or the script isn't a named class, it should return [code]""[/code].
*/
func (class) _get_resource_script_class(impl func(ptr unsafe.Pointer, path gd.String) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_resource_uid(impl func(ptr unsafe.Pointer, path gd.String) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
If implemented, gets the dependencies of a given resource. If [param add_types] is [code]true[/code], paths should be appended [code]::TypeName[/code], where [code]TypeName[/code] is the class name of the dependency.
[b]Note:[/b] Custom resource types defined by scripts aren't known by the [ClassDB], so you might just return [code]"Resource"[/code] for them.
*/
func (class) _get_dependencies(impl func(ptr unsafe.Pointer, path gd.String, add_types bool) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var add_types = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, add_types)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
If implemented, renames dependencies within the given resource and saves it. [param renames] is a dictionary [code]{ String => String }[/code] mapping old dependency paths to new paths.
Returns [constant OK] on success, or an [enum Error] constant in case of failure.
*/
func (class) _rename_dependencies(impl func(ptr unsafe.Pointer, path gd.String, renames gd.Dictionary) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var renames = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, renames)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _exists(impl func(ptr unsafe.Pointer, path gd.String) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_classes_used(impl func(ptr unsafe.Pointer, path gd.String) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Loads a resource when the engine finds this loader to be compatible. If the loaded resource is the result of an import, [param original_path] will target the source file. Returns a [Resource] object on success, or an [enum Error] constant in case of failure.
The [param cache_mode] property defines whether and how the cache should be used or updated when loading the resource. See [enum CacheMode] for details.
*/
func (class) _load(impl func(ptr unsafe.Pointer, path gd.String, original_path gd.String, use_sub_threads bool, cache_mode gd.Int) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var original_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		var use_sub_threads = gd.UnsafeGet[bool](p_args, 2)
		var cache_mode = gd.UnsafeGet[gd.Int](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, original_path, use_sub_threads, cache_mode)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (self class) AsResourceFormatLoader() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsResourceFormatLoader() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted         { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted      { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	case "_recognize_path":
		return reflect.ValueOf(self._recognize_path)
	case "_handles_type":
		return reflect.ValueOf(self._handles_type)
	case "_get_resource_type":
		return reflect.ValueOf(self._get_resource_type)
	case "_get_resource_script_class":
		return reflect.ValueOf(self._get_resource_script_class)
	case "_get_resource_uid":
		return reflect.ValueOf(self._get_resource_uid)
	case "_get_dependencies":
		return reflect.ValueOf(self._get_dependencies)
	case "_rename_dependencies":
		return reflect.ValueOf(self._rename_dependencies)
	case "_exists":
		return reflect.ValueOf(self._exists)
	case "_get_classes_used":
		return reflect.ValueOf(self._get_classes_used)
	case "_load":
		return reflect.ValueOf(self._load)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	case "_recognize_path":
		return reflect.ValueOf(self._recognize_path)
	case "_handles_type":
		return reflect.ValueOf(self._handles_type)
	case "_get_resource_type":
		return reflect.ValueOf(self._get_resource_type)
	case "_get_resource_script_class":
		return reflect.ValueOf(self._get_resource_script_class)
	case "_get_resource_uid":
		return reflect.ValueOf(self._get_resource_uid)
	case "_get_dependencies":
		return reflect.ValueOf(self._get_dependencies)
	case "_rename_dependencies":
		return reflect.ValueOf(self._rename_dependencies)
	case "_exists":
		return reflect.ValueOf(self._exists)
	case "_get_classes_used":
		return reflect.ValueOf(self._get_classes_used)
	case "_load":
		return reflect.ValueOf(self._load)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("ResourceFormatLoader", func(ptr gd.Object) any { return classdb.ResourceFormatLoader(ptr) })
}

type CacheMode = classdb.ResourceFormatLoaderCacheMode

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
