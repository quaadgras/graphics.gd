package ResourceFormatSaver

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
The engine can save resources when you do it from the editor, or when you use the [ResourceSaver] singleton. This is accomplished thanks to multiple [ResourceFormatSaver]s, each handling its own format and called automatically by the engine.
By default, Godot saves resources as [code].tres[/code] (text-based), [code].res[/code] (binary) or another built-in format, but you can choose to create your own format by extending this class. Be sure to respect the documented return types and values. You should give it a global class name with [code]class_name[/code] for it to be registered. Like built-in ResourceFormatSavers, it will be called automatically when saving resources of its recognized type(s). You may also implement a [ResourceFormatLoader].

	// ResourceFormatSaver methods that can be overridden by a [Class] that extends it.
	type ResourceFormatSaver interface {
		//Saves the given resource object to a file at the target [param path]. [param flags] is a bitmask composed with [enum ResourceSaver.SaverFlags] constants.
		//Returns [constant OK] on success, or an [enum Error] constant in case of failure.
		Save(resource gdclass.Resource, path string, flags int) gd.Error
		//Sets a new UID for the resource at the given [param path]. Returns [constant OK] on success, or an [enum Error] constant in case of failure.
		SetUid(path string, uid int) gd.Error
		//Returns whether the given resource object can be saved by this saver.
		Recognize(resource gdclass.Resource) bool
		//Returns the list of extensions available for saving the resource object, provided it is recognized (see [method _recognize]).
		GetRecognizedExtensions(resource gdclass.Resource) []string
		//Returns [code]true[/code] if this saver handles a given save path and [code]false[/code] otherwise.
		//If this method is not implemented, the default behavior returns whether the path's extension is within the ones provided by [method _get_recognized_extensions].
		RecognizePath(resource gdclass.Resource, path string) bool
	}
*/
type Instance [1]classdb.ResourceFormatSaver

/*
Saves the given resource object to a file at the target [param path]. [param flags] is a bitmask composed with [enum ResourceSaver.SaverFlags] constants.
Returns [constant OK] on success, or an [enum Error] constant in case of failure.
*/
func (Instance) _save(impl func(ptr unsafe.Pointer, resource gdclass.Resource, path string, flags int) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = gdclass.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(path)
		var flags = gd.UnsafeGet[gd.Int](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path.String(), int(flags))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Sets a new UID for the resource at the given [param path]. Returns [constant OK] on success, or an [enum Error] constant in case of failure.
*/
func (Instance) _set_uid(impl func(ptr unsafe.Pointer, path string, uid int) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(path)
		var uid = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), int(uid))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns whether the given resource object can be saved by this saver.
*/
func (Instance) _recognize(impl func(ptr unsafe.Pointer, resource gdclass.Resource) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = gdclass.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the list of extensions available for saving the resource object, provided it is recognized (see [method _recognize]).
*/
func (Instance) _get_recognized_extensions(impl func(ptr unsafe.Pointer, resource gdclass.Resource) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = gdclass.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns [code]true[/code] if this saver handles a given save path and [code]false[/code] otherwise.
If this method is not implemented, the default behavior returns whether the path's extension is within the ones provided by [method _get_recognized_extensions].
*/
func (Instance) _recognize_path(impl func(ptr unsafe.Pointer, resource gdclass.Resource, path string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = gdclass.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(path)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path.String())
		gd.UnsafeSet(p_back, ret)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ResourceFormatSaver

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ResourceFormatSaver"))
	return Instance{classdb.ResourceFormatSaver(object)}
}

/*
Saves the given resource object to a file at the target [param path]. [param flags] is a bitmask composed with [enum ResourceSaver.SaverFlags] constants.
Returns [constant OK] on success, or an [enum Error] constant in case of failure.
*/
func (class) _save(impl func(ptr unsafe.Pointer, resource gdclass.Resource, path gd.String, flags gd.Int) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = gdclass.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		var flags = gd.UnsafeGet[gd.Int](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path, flags)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Sets a new UID for the resource at the given [param path]. Returns [constant OK] on success, or an [enum Error] constant in case of failure.
*/
func (class) _set_uid(impl func(ptr unsafe.Pointer, path gd.String, uid gd.Int) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var uid = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, uid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns whether the given resource object can be saved by this saver.
*/
func (class) _recognize(impl func(ptr unsafe.Pointer, resource gdclass.Resource) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = gdclass.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the list of extensions available for saving the resource object, provided it is recognized (see [method _recognize]).
*/
func (class) _get_recognized_extensions(impl func(ptr unsafe.Pointer, resource gdclass.Resource) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = gdclass.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns [code]true[/code] if this saver handles a given save path and [code]false[/code] otherwise.
If this method is not implemented, the default behavior returns whether the path's extension is within the ones provided by [method _get_recognized_extensions].
*/
func (class) _recognize_path(impl func(ptr unsafe.Pointer, resource gdclass.Resource, path gd.String) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = gdclass.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsResourceFormatSaver() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsResourceFormatSaver() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted        { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted     { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_save":
		return reflect.ValueOf(self._save)
	case "_set_uid":
		return reflect.ValueOf(self._set_uid)
	case "_recognize":
		return reflect.ValueOf(self._recognize)
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	case "_recognize_path":
		return reflect.ValueOf(self._recognize_path)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_save":
		return reflect.ValueOf(self._save)
	case "_set_uid":
		return reflect.ValueOf(self._set_uid)
	case "_recognize":
		return reflect.ValueOf(self._recognize)
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	case "_recognize_path":
		return reflect.ValueOf(self._recognize_path)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("ResourceFormatSaver", func(ptr gd.Object) any { return classdb.ResourceFormatSaver(ptr) })
}
