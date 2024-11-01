package ImageFormatLoaderExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/ImageFormatLoader"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
The engine supports multiple image formats out of the box (PNG, SVG, JPEG, WebP to name a few), but you can choose to implement support for additional image formats by extending this class.
Be sure to respect the documented return types and values. You should create an instance of it, and call [method add_format_loader] to register that loader during the initialization phase.

	// ImageFormatLoaderExtension methods that can be overridden by a [Class] that extends it.
	type ImageFormatLoaderExtension interface {
		//Returns the list of file extensions for this image format. Files with the given extensions will be treated as image file and loaded using this class.
		GetRecognizedExtensions() []string
		//Loads the content of [param fileaccess] into the provided [param image].
		LoadImage(image objects.Image, fileaccess objects.FileAccess, flags classdb.ImageFormatLoaderLoaderFlags, scale float64) gd.Error
	}
*/
type Instance [1]classdb.ImageFormatLoaderExtension

/*
Returns the list of file extensions for this image format. Files with the given extensions will be treated as image file and loaded using this class.
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
Loads the content of [param fileaccess] into the provided [param image].
*/
func (Instance) _load_image(impl func(ptr unsafe.Pointer, image objects.Image, fileaccess objects.FileAccess, flags classdb.ImageFormatLoaderLoaderFlags, scale float64) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var image = objects.Image{pointers.New[classdb.Image]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(image[0])
		var fileaccess = objects.FileAccess{pointers.New[classdb.FileAccess]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(fileaccess[0])
		var flags = gd.UnsafeGet[classdb.ImageFormatLoaderLoaderFlags](p_args, 2)
		var scale = gd.UnsafeGet[gd.Float](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, image, fileaccess, flags, float64(scale))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Add this format loader to the engine, allowing it to recognize the file extensions returned by [method _get_recognized_extensions].
*/
func (self Instance) AddFormatLoader() {
	class(self).AddFormatLoader()
}

/*
Remove this format loader from the engine.
*/
func (self Instance) RemoveFormatLoader() {
	class(self).RemoveFormatLoader()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ImageFormatLoaderExtension

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ImageFormatLoaderExtension"))
	return Instance{classdb.ImageFormatLoaderExtension(object)}
}

/*
Returns the list of file extensions for this image format. Files with the given extensions will be treated as image file and loaded using this class.
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
Loads the content of [param fileaccess] into the provided [param image].
*/
func (class) _load_image(impl func(ptr unsafe.Pointer, image objects.Image, fileaccess objects.FileAccess, flags classdb.ImageFormatLoaderLoaderFlags, scale gd.Float) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var image = objects.Image{pointers.New[classdb.Image]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(image[0])
		var fileaccess = objects.FileAccess{pointers.New[classdb.FileAccess]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(fileaccess[0])
		var flags = gd.UnsafeGet[classdb.ImageFormatLoaderLoaderFlags](p_args, 2)
		var scale = gd.UnsafeGet[gd.Float](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, image, fileaccess, flags, scale)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Add this format loader to the engine, allowing it to recognize the file extensions returned by [method _get_recognized_extensions].
*/
//go:nosplit
func (self class) AddFormatLoader() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageFormatLoaderExtension.Bind_add_format_loader, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Remove this format loader from the engine.
*/
//go:nosplit
func (self class) RemoveFormatLoader() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageFormatLoaderExtension.Bind_remove_format_loader, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsImageFormatLoaderExtension() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsImageFormatLoaderExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsImageFormatLoader() ImageFormatLoader.Advanced {
	return *((*ImageFormatLoader.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsImageFormatLoader() ImageFormatLoader.Instance {
	return *((*ImageFormatLoader.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	case "_load_image":
		return reflect.ValueOf(self._load_image)
	default:
		return gd.VirtualByName(self.AsImageFormatLoader(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	case "_load_image":
		return reflect.ValueOf(self._load_image)
	default:
		return gd.VirtualByName(self.AsImageFormatLoader(), name)
	}
}
func init() {
	classdb.Register("ImageFormatLoaderExtension", func(ptr gd.Object) any { return classdb.ImageFormatLoaderExtension(ptr) })
}
