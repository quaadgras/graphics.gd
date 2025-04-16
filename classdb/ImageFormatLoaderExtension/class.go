// Package ImageFormatLoaderExtension provides methods for working with ImageFormatLoaderExtension object instances.
package ImageFormatLoaderExtension

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/ImageFormatLoader"
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
The engine supports multiple image formats out of the box (PNG, SVG, JPEG, WebP to name a few), but you can choose to implement support for additional image formats by extending this class.
Be sure to respect the documented return types and values. You should create an instance of it, and call [method add_format_loader] to register that loader during the initialization phase.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=ImageFormatLoaderExtension)
*/
type Instance [1]gdclass.ImageFormatLoaderExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsImageFormatLoaderExtension() Instance
}
type Interface interface {
	//Returns the list of file extensions for this image format. Files with the given extensions will be treated as image file and loaded using this class.
	GetRecognizedExtensions() []string
	//Loads the content of [param fileaccess] into the provided [param image].
	LoadImage(image [1]gdclass.Image, fileaccess [1]gdclass.FileAccess, flags gdclass.ImageFormatLoaderLoaderFlags, scale Float.X) error
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetRecognizedExtensions() (_ []string) { return }
func (self implementation) LoadImage(image [1]gdclass.Image, fileaccess [1]gdclass.FileAccess, flags gdclass.ImageFormatLoaderLoaderFlags, scale Float.X) (_ error) {
	return
}

/*
Returns the list of file extensions for this image format. Files with the given extensions will be treated as image file and loaded using this class.
*/
func (Instance) _get_recognized_extensions(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalPackedStrings(Packed.MakeStrings(ret...)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Loads the content of [param fileaccess] into the provided [param image].
*/
func (Instance) _load_image(impl func(ptr unsafe.Pointer, image [1]gdclass.Image, fileaccess [1]gdclass.FileAccess, flags gdclass.ImageFormatLoaderLoaderFlags, scale Float.X) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var image = [1]gdclass.Image{pointers.New[gdclass.Image]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(image[0])
		var fileaccess = [1]gdclass.FileAccess{pointers.New[gdclass.FileAccess]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}

		defer pointers.End(fileaccess[0])
		var flags = gd.UnsafeGet[gdclass.ImageFormatLoaderLoaderFlags](p_args, 2)
		var scale = gd.UnsafeGet[float64](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, image, fileaccess, flags, Float.X(scale))
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Add this format loader to the engine, allowing it to recognize the file extensions returned by [method _get_recognized_extensions].
*/
func (self Instance) AddFormatLoader() { //gd:ImageFormatLoaderExtension.add_format_loader
	Advanced(self).AddFormatLoader()
}

/*
Remove this format loader from the engine.
*/
func (self Instance) RemoveFormatLoader() { //gd:ImageFormatLoaderExtension.remove_format_loader
	Advanced(self).RemoveFormatLoader()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ImageFormatLoaderExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ImageFormatLoaderExtension"))
	casted := Instance{*(*gdclass.ImageFormatLoaderExtension)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Returns the list of file extensions for this image format. Files with the given extensions will be treated as image file and loaded using this class.
*/
func (class) _get_recognized_extensions(impl func(ptr unsafe.Pointer) Packed.Strings) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalPackedStrings(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Loads the content of [param fileaccess] into the provided [param image].
*/
func (class) _load_image(impl func(ptr unsafe.Pointer, image [1]gdclass.Image, fileaccess [1]gdclass.FileAccess, flags gdclass.ImageFormatLoaderLoaderFlags, scale float64) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var image = [1]gdclass.Image{pointers.New[gdclass.Image]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(image[0])
		var fileaccess = [1]gdclass.FileAccess{pointers.New[gdclass.FileAccess]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}

		defer pointers.End(fileaccess[0])
		var flags = gd.UnsafeGet[gdclass.ImageFormatLoaderLoaderFlags](p_args, 2)
		var scale = gd.UnsafeGet[float64](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, image, fileaccess, flags, scale)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Add this format loader to the engine, allowing it to recognize the file extensions returned by [method _get_recognized_extensions].
*/
//go:nosplit
func (self class) AddFormatLoader() { //gd:ImageFormatLoaderExtension.add_format_loader
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageFormatLoaderExtension.Bind_add_format_loader, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Remove this format loader from the engine.
*/
//go:nosplit
func (self class) RemoveFormatLoader() { //gd:ImageFormatLoaderExtension.remove_format_loader
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageFormatLoaderExtension.Bind_remove_format_loader, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	case "_load_image":
		return reflect.ValueOf(self._load_image)
	default:
		return gd.VirtualByName(ImageFormatLoader.Advanced(self.AsImageFormatLoader()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	case "_load_image":
		return reflect.ValueOf(self._load_image)
	default:
		return gd.VirtualByName(ImageFormatLoader.Instance(self.AsImageFormatLoader()), name)
	}
}
func init() {
	gdclass.Register("ImageFormatLoaderExtension", func(ptr gd.Object) any {
		return [1]gdclass.ImageFormatLoaderExtension{*(*gdclass.ImageFormatLoaderExtension)(unsafe.Pointer(&ptr))}
	})
}
