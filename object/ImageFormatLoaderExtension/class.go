package ImageFormatLoaderExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/ImageFormatLoader"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The engine supports multiple image formats out of the box (PNG, SVG, JPEG, WebP to name a few), but you can choose to implement support for additional image formats by extending this class.
Be sure to respect the documented return types and values. You should create an instance of it, and call [method add_format_loader] to register that loader during the initialization phase.
	// ImageFormatLoaderExtension methods that can be overridden by a [Class] that extends it.
	type ImageFormatLoaderExtension interface {
		//Returns the list of file extensions for this image format. Files with the given extensions will be treated as image file and loaded using this class.
		GetRecognizedExtensions() gd.PackedStringArray
		//Loads the content of [param fileaccess] into the provided [param image].
		LoadImage(image object.Image, fileaccess object.FileAccess, flags classdb.ImageFormatLoaderLoaderFlags, scale gd.Float) int64
	}

*/
type Simple [1]classdb.ImageFormatLoaderExtension
func (Simple) _get_recognized_extensions(impl func(ptr unsafe.Pointer) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _load_image(impl func(ptr unsafe.Pointer, image [1]classdb.Image, fileaccess [1]classdb.FileAccess, flags classdb.ImageFormatLoaderLoaderFlags, scale float64) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var image [1]classdb.Image
		image[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var fileaccess [1]classdb.FileAccess
		fileaccess[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,1)}))
		var flags = gd.UnsafeGet[classdb.ImageFormatLoaderLoaderFlags](p_args,2)
		var scale = gd.UnsafeGet[gd.Float](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, image, fileaccess, flags, float64(scale))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (self Simple) AddFormatLoader() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddFormatLoader()
}
func (self Simple) RemoveFormatLoader() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveFormatLoader()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ImageFormatLoaderExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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
func (class) _load_image(impl func(ptr unsafe.Pointer, image object.Image, fileaccess object.FileAccess, flags classdb.ImageFormatLoaderLoaderFlags, scale gd.Float) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var image object.Image
		image[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var fileaccess object.FileAccess
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

//go:nosplit
func (self class) AsImageFormatLoaderExtension() Expert { return self[0].AsImageFormatLoaderExtension() }


//go:nosplit
func (self Simple) AsImageFormatLoaderExtension() Simple { return self[0].AsImageFormatLoaderExtension() }


//go:nosplit
func (self class) AsImageFormatLoader() ImageFormatLoader.Expert { return self[0].AsImageFormatLoader() }


//go:nosplit
func (self Simple) AsImageFormatLoader() ImageFormatLoader.Simple { return self[0].AsImageFormatLoader() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_recognized_extensions": return reflect.ValueOf(self._get_recognized_extensions);
	case "_load_image": return reflect.ValueOf(self._load_image);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_get_recognized_extensions": return reflect.ValueOf(self._get_recognized_extensions);
	case "_load_image": return reflect.ValueOf(self._load_image);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("ImageFormatLoaderExtension", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
