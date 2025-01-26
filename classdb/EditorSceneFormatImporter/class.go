// Package EditorSceneFormatImporter provides methods for working with EditorSceneFormatImporter object instances.
package EditorSceneFormatImporter

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
import "graphics.gd/variant/RID"

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

/*
[EditorSceneFormatImporter] allows to define an importer script for a third-party 3D format.
To use [EditorSceneFormatImporter], register it using the [method EditorPlugin.add_scene_format_importer_plugin] method first.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=EditorSceneFormatImporter)
*/
type Instance [1]gdclass.EditorSceneFormatImporter

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorSceneFormatImporter() Instance
}
type Interface interface {
	GetImportFlags() int
	GetExtensions() []string
	ImportScene(path string, flags int, options map[any]any) Object.Instance
	GetImportOptions(path string)
	GetOptionVisibility(path string, for_animation bool, option string) any
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetImportFlags() (_ int)     { return }
func (self implementation) GetExtensions() (_ []string) { return }
func (self implementation) ImportScene(path string, flags int, options map[any]any) (_ Object.Instance) {
	return
}
func (self implementation) GetImportOptions(path string) { return }
func (self implementation) GetOptionVisibility(path string, for_animation bool, option string) (_ any) {
	return
}
func (Instance) _get_import_flags(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _get_extensions(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _import_scene(impl func(ptr unsafe.Pointer, path string, flags int, options map[any]any) Object.Instance) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(path)
		var flags = gd.UnsafeGet[gd.Int](p_args, 1)

		var options = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalDictionary(options))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), int(flags), gd.DictionaryAs[map[any]any](options))
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_import_options(impl func(ptr unsafe.Pointer, path string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(path)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, path.String())
	}
}
func (Instance) _get_option_visibility(impl func(ptr unsafe.Pointer, path string, for_animation bool, option string) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(path)
		var for_animation = gd.UnsafeGet[bool](p_args, 1)

		var option = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		defer pointers.End(option)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), for_animation, option.String())
		ptr, ok := pointers.End(gd.NewVariant(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorSceneFormatImporter

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorSceneFormatImporter"))
	casted := Instance{*(*gdclass.EditorSceneFormatImporter)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (class) _get_import_flags(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_extensions(impl func(ptr unsafe.Pointer) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _import_scene(impl func(ptr unsafe.Pointer, path gd.String, flags gd.Int, options Dictionary.Any) [1]gd.Object) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(path)
		var flags = gd.UnsafeGet[gd.Int](p_args, 1)

		var options = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalDictionary(options))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, flags, options)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_import_options(impl func(ptr unsafe.Pointer, path gd.String)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(path)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, path)
	}
}

func (class) _get_option_visibility(impl func(ptr unsafe.Pointer, path gd.String, for_animation bool, option gd.String) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(path)
		var for_animation = gd.UnsafeGet[bool](p_args, 1)

		var option = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		defer pointers.End(option)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, for_animation, option)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (self class) AsEditorSceneFormatImporter() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorSceneFormatImporter() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_import_flags":
		return reflect.ValueOf(self._get_import_flags)
	case "_get_extensions":
		return reflect.ValueOf(self._get_extensions)
	case "_import_scene":
		return reflect.ValueOf(self._import_scene)
	case "_get_import_options":
		return reflect.ValueOf(self._get_import_options)
	case "_get_option_visibility":
		return reflect.ValueOf(self._get_option_visibility)
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_import_flags":
		return reflect.ValueOf(self._get_import_flags)
	case "_get_extensions":
		return reflect.ValueOf(self._get_extensions)
	case "_import_scene":
		return reflect.ValueOf(self._import_scene)
	case "_get_import_options":
		return reflect.ValueOf(self._get_import_options)
	case "_get_option_visibility":
		return reflect.ValueOf(self._get_option_visibility)
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("EditorSceneFormatImporter", func(ptr gd.Object) any {
		return [1]gdclass.EditorSceneFormatImporter{*(*gdclass.EditorSceneFormatImporter)(unsafe.Pointer(&ptr))}
	})
}
