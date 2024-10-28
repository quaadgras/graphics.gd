package EditorSceneFormatImporter

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
[EditorSceneFormatImporter] allows to define an importer script for a third-party 3D format.
To use [EditorSceneFormatImporter], register it using the [method EditorPlugin.add_scene_format_importer_plugin] method first.
	// EditorSceneFormatImporter methods that can be overridden by a [Class] that extends it.
	type EditorSceneFormatImporter interface {
		GetImportFlags() int
		GetExtensions() []string
		ImportScene(path string, flags int, options gd.Dictionary) gd.Object
		GetImportOptions(path string) 
		GetOptionVisibility(path string, for_animation bool, option string) gd.Variant
	}

*/
type Go [1]classdb.EditorSceneFormatImporter
func (Go) _get_import_flags(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Go) _get_extensions(impl func(ptr unsafe.Pointer) []string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
ptr, ok := discreet.End(gd.NewPackedStringSlice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Go) _import_scene(impl func(ptr unsafe.Pointer, path string, flags int, options gd.Dictionary) gd.Object, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		defer discreet.End(path)
		var flags = gd.UnsafeGet[gd.Int](p_args,1)
		var options = discreet.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args,2))
		defer discreet.End(options)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), int(flags), options)
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Go) _get_import_options(impl func(ptr unsafe.Pointer, path string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		defer discreet.End(path)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, path.String())
	}
}
func (Go) _get_option_visibility(impl func(ptr unsafe.Pointer, path string, for_animation bool, option string) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		defer discreet.End(path)
		var for_animation = gd.UnsafeGet[bool](p_args,1)
		var option = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,2))
		defer discreet.End(option)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), for_animation, option.String())
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorSceneFormatImporter
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorSceneFormatImporter"))
	return Go{classdb.EditorSceneFormatImporter(object)}
}

func (class) _get_import_flags(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_extensions(impl func(ptr unsafe.Pointer) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _import_scene(impl func(ptr unsafe.Pointer, path gd.String, flags gd.Int, options gd.Dictionary) gd.Object, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		var flags = gd.UnsafeGet[gd.Int](p_args,1)
		var options = discreet.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, flags, options)
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_import_options(impl func(ptr unsafe.Pointer, path gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, path)
	}
}

func (class) _get_option_visibility(impl func(ptr unsafe.Pointer, path gd.String, for_animation bool, option gd.String) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		var for_animation = gd.UnsafeGet[bool](p_args,1)
		var option = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, for_animation, option)
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (self class) AsEditorSceneFormatImporter() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorSceneFormatImporter() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_import_flags": return reflect.ValueOf(self._get_import_flags);
	case "_get_extensions": return reflect.ValueOf(self._get_extensions);
	case "_import_scene": return reflect.ValueOf(self._import_scene);
	case "_get_import_options": return reflect.ValueOf(self._get_import_options);
	case "_get_option_visibility": return reflect.ValueOf(self._get_option_visibility);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_get_import_flags": return reflect.ValueOf(self._get_import_flags);
	case "_get_extensions": return reflect.ValueOf(self._get_extensions);
	case "_import_scene": return reflect.ValueOf(self._import_scene);
	case "_get_import_options": return reflect.ValueOf(self._get_import_options);
	case "_get_option_visibility": return reflect.ValueOf(self._get_option_visibility);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("EditorSceneFormatImporter", func(ptr gd.Object) any { return classdb.EditorSceneFormatImporter(ptr) })}
