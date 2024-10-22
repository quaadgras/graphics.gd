package ImageFormatLoaderExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/ImageFormatLoader"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The engine supports multiple image formats out of the box (PNG, SVG, JPEG, WebP to name a few), but you can choose to implement support for additional image formats by extending this class.
Be sure to respect the documented return types and values. You should create an instance of it, and call [method add_format_loader] to register that loader during the initialization phase.
	// ImageFormatLoaderExtension methods that can be overridden by a [Class] that extends it.
	type ImageFormatLoaderExtension interface {
		//Returns the list of file extensions for this image format. Files with the given extensions will be treated as image file and loaded using this class.
		GetRecognizedExtensions() []string
		//Loads the content of [param fileaccess] into the provided [param image].
		LoadImage(image gdclass.Image, fileaccess gdclass.FileAccess, flags classdb.ImageFormatLoaderLoaderFlags, scale float64) gd.Error
	}

*/
type Go [1]classdb.ImageFormatLoaderExtension

/*
Returns the list of file extensions for this image format. Files with the given extensions will be treated as image file and loaded using this class.
*/
func (Go) _get_recognized_extensions(impl func(ptr unsafe.Pointer) []string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.PackedStringSlice(ret)))
		gc.End()
	}
}

/*
Loads the content of [param fileaccess] into the provided [param image].
*/
func (Go) _load_image(impl func(ptr unsafe.Pointer, image gdclass.Image, fileaccess gdclass.FileAccess, flags classdb.ImageFormatLoaderLoaderFlags, scale float64) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var image gdclass.Image
		image[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var fileaccess gdclass.FileAccess
		fileaccess[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,1)}))
		var flags = gd.UnsafeGet[classdb.ImageFormatLoaderLoaderFlags](p_args,2)
		var scale = gd.UnsafeGet[gd.Float](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, image, fileaccess, flags, float64(scale))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Add this format loader to the engine, allowing it to recognize the file extensions returned by [method _get_recognized_extensions].
*/
func (self Go) AddFormatLoader() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddFormatLoader()
}

/*
Remove this format loader from the engine.
*/
func (self Go) RemoveFormatLoader() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveFormatLoader()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ImageFormatLoaderExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ImageFormatLoaderExtension"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns the list of file extensions for this image format. Files with the given extensions will be treated as image file and loaded using this class.
*/
func (class) _get_recognized_extensions(impl func(ptr unsafe.Pointer) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Loads the content of [param fileaccess] into the provided [param image].
*/
func (class) _load_image(impl func(ptr unsafe.Pointer, image gdclass.Image, fileaccess gdclass.FileAccess, flags classdb.ImageFormatLoaderLoaderFlags, scale gd.Float) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var image gdclass.Image
		image[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var fileaccess gdclass.FileAccess
		fileaccess[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,1)}))
		var flags = gd.UnsafeGet[classdb.ImageFormatLoaderLoaderFlags](p_args,2)
		var scale = gd.UnsafeGet[gd.Float](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, image, fileaccess, flags, scale)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Add this format loader to the engine, allowing it to recognize the file extensions returned by [method _get_recognized_extensions].
*/
//go:nosplit
func (self class) AddFormatLoader()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImageFormatLoaderExtension.Bind_add_format_loader, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Remove this format loader from the engine.
*/
//go:nosplit
func (self class) RemoveFormatLoader()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImageFormatLoaderExtension.Bind_remove_format_loader, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsImageFormatLoaderExtension() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsImageFormatLoaderExtension() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsImageFormatLoader() ImageFormatLoader.GD { return *((*ImageFormatLoader.GD)(unsafe.Pointer(&self))) }
func (self Go) AsImageFormatLoader() ImageFormatLoader.Go { return *((*ImageFormatLoader.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_recognized_extensions": return reflect.ValueOf(self._get_recognized_extensions);
	case "_load_image": return reflect.ValueOf(self._load_image);
	default: return gd.VirtualByName(self.AsImageFormatLoader(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_get_recognized_extensions": return reflect.ValueOf(self._get_recognized_extensions);
	case "_load_image": return reflect.ValueOf(self._load_image);
	default: return gd.VirtualByName(self.AsImageFormatLoader(), name)
	}
}
func init() {classdb.Register("ImageFormatLoaderExtension", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
