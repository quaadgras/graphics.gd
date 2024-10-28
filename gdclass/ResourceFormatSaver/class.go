package ResourceFormatSaver

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

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
type Go [1]classdb.ResourceFormatSaver

/*
Saves the given resource object to a file at the target [param path]. [param flags] is a bitmask composed with [enum ResourceSaver.SaverFlags] constants.
Returns [constant OK] on success, or an [enum Error] constant in case of failure.
*/
func (Go) _save(impl func(ptr unsafe.Pointer, resource gdclass.Resource, path string, flags int) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = gdclass.Resource{discreet.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})}
		defer discreet.End(resource[0])
		var path = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,1))
		defer discreet.End(path)
		var flags = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path.String(), int(flags))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Sets a new UID for the resource at the given [param path]. Returns [constant OK] on success, or an [enum Error] constant in case of failure.
*/
func (Go) _set_uid(impl func(ptr unsafe.Pointer, path string, uid int) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		defer discreet.End(path)
		var uid = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), int(uid))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns whether the given resource object can be saved by this saver.
*/
func (Go) _recognize(impl func(ptr unsafe.Pointer, resource gdclass.Resource) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = gdclass.Resource{discreet.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})}
		defer discreet.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the list of extensions available for saving the resource object, provided it is recognized (see [method _recognize]).
*/
func (Go) _get_recognized_extensions(impl func(ptr unsafe.Pointer, resource gdclass.Resource) []string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = gdclass.Resource{discreet.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})}
		defer discreet.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
ptr, ok := discreet.End(gd.NewPackedStringSlice(ret))
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
func (Go) _recognize_path(impl func(ptr unsafe.Pointer, resource gdclass.Resource, path string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = gdclass.Resource{discreet.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})}
		defer discreet.End(resource[0])
		var path = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,1))
		defer discreet.End(path)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path.String())
		gd.UnsafeSet(p_back, ret)
	}
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ResourceFormatSaver
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ResourceFormatSaver"))
	return Go{classdb.ResourceFormatSaver(object)}
}

/*
Saves the given resource object to a file at the target [param path]. [param flags] is a bitmask composed with [enum ResourceSaver.SaverFlags] constants.
Returns [constant OK] on success, or an [enum Error] constant in case of failure.
*/
func (class) _save(impl func(ptr unsafe.Pointer, resource gdclass.Resource, path gd.String, flags gd.Int) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = gdclass.Resource{discreet.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})}
		defer discreet.End(resource[0])
		var path = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,1))
		var flags = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path, flags)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Sets a new UID for the resource at the given [param path]. Returns [constant OK] on success, or an [enum Error] constant in case of failure.
*/
func (class) _set_uid(impl func(ptr unsafe.Pointer, path gd.String, uid gd.Int) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		var uid = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, uid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns whether the given resource object can be saved by this saver.
*/
func (class) _recognize(impl func(ptr unsafe.Pointer, resource gdclass.Resource) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = gdclass.Resource{discreet.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})}
		defer discreet.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the list of extensions available for saving the resource object, provided it is recognized (see [method _recognize]).
*/
func (class) _get_recognized_extensions(impl func(ptr unsafe.Pointer, resource gdclass.Resource) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = gdclass.Resource{discreet.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})}
		defer discreet.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
ptr, ok := discreet.End(ret)
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
func (class) _recognize_path(impl func(ptr unsafe.Pointer, resource gdclass.Resource, path gd.String) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = gdclass.Resource{discreet.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})}
		defer discreet.End(resource[0])
		var path = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsResourceFormatSaver() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsResourceFormatSaver() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_save": return reflect.ValueOf(self._save);
	case "_set_uid": return reflect.ValueOf(self._set_uid);
	case "_recognize": return reflect.ValueOf(self._recognize);
	case "_get_recognized_extensions": return reflect.ValueOf(self._get_recognized_extensions);
	case "_recognize_path": return reflect.ValueOf(self._recognize_path);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_save": return reflect.ValueOf(self._save);
	case "_set_uid": return reflect.ValueOf(self._set_uid);
	case "_recognize": return reflect.ValueOf(self._recognize);
	case "_get_recognized_extensions": return reflect.ValueOf(self._get_recognized_extensions);
	case "_recognize_path": return reflect.ValueOf(self._recognize_path);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("ResourceFormatSaver", func(ptr gd.Object) any { return classdb.ResourceFormatSaver(ptr) })}
