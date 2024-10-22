package TextureLayered

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Texture"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Base class for [ImageTextureLayered] and [CompressedTextureLayered]. Cannot be used directly, but contains all the functions necessary for accessing the derived resource types. See also [Texture3D].
Data is set on a per-layer basis. For [Texture2DArray]s, the layer specifies the array layer.
All images need to have the same width, height and number of mipmap levels.
A [TextureLayered] can be loaded with [method ResourceLoader.load].
Internally, Godot maps these files to their respective counterparts in the target rendering driver (Vulkan, OpenGL3).
	// TextureLayered methods that can be overridden by a [Class] that extends it.
	type TextureLayered interface {
		//Called when the [TextureLayered]'s format is queried.
		GetFormat() classdb.ImageFormat
		//Called when the layers' type in the [TextureLayered] is queried.
		GetLayeredType() int
		//Called when the [TextureLayered]'s width queried.
		GetWidth() int
		//Called when the [TextureLayered]'s height is queried.
		GetHeight() int
		//Called when the number of layers in the [TextureLayered] is queried.
		GetLayers() int
		//Called when the presence of mipmaps in the [TextureLayered] is queried.
		HasMipmaps() bool
		//Called when the data for a layer in the [TextureLayered] is queried.
		GetLayerData(layer_index int) gdclass.Image
	}

*/
type Go [1]classdb.TextureLayered

/*
Called when the [TextureLayered]'s format is queried.
*/
func (Go) _get_format(impl func(ptr unsafe.Pointer) classdb.ImageFormat, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Called when the layers' type in the [TextureLayered] is queried.
*/
func (Go) _get_layered_type(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Called when the [TextureLayered]'s width queried.
*/
func (Go) _get_width(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Called when the [TextureLayered]'s height is queried.
*/
func (Go) _get_height(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Called when the number of layers in the [TextureLayered] is queried.
*/
func (Go) _get_layers(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Called when the presence of mipmaps in the [TextureLayered] is queried.
*/
func (Go) _has_mipmaps(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Called when the data for a layer in the [TextureLayered] is queried.
*/
func (Go) _get_layer_data(impl func(ptr unsafe.Pointer, layer_index int) gdclass.Image, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var layer_index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(layer_index))
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		gc.End()
	}
}

/*
Returns the current format being used by this texture. See [enum Image.Format] for details.
*/
func (self Go) GetFormat() classdb.ImageFormat {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ImageFormat(class(self).GetFormat())
}

/*
Returns the [TextureLayered]'s type. The type determines how the data is accessed, with cubemaps having special types.
*/
func (self Go) GetLayeredType() classdb.TextureLayeredLayeredType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextureLayeredLayeredType(class(self).GetLayeredType())
}

/*
Returns the width of the texture in pixels. Width is typically represented by the X axis.
*/
func (self Go) GetWidth() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetWidth()))
}

/*
Returns the height of the texture in pixels. Height is typically represented by the Y axis.
*/
func (self Go) GetHeight() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetHeight()))
}

/*
Returns the number of referenced [Image]s.
*/
func (self Go) GetLayers() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetLayers()))
}

/*
Returns [code]true[/code] if the layers have generated mipmaps.
*/
func (self Go) HasMipmaps() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasMipmaps())
}

/*
Returns an [Image] resource with the data from specified [param layer].
*/
func (self Go) GetLayerData(layer int) gdclass.Image {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Image(class(self).GetLayerData(gc, gd.Int(layer)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TextureLayered
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("TextureLayered"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Called when the [TextureLayered]'s format is queried.
*/
func (class) _get_format(impl func(ptr unsafe.Pointer) classdb.ImageFormat, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Called when the layers' type in the [TextureLayered] is queried.
*/
func (class) _get_layered_type(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Called when the [TextureLayered]'s width queried.
*/
func (class) _get_width(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Called when the [TextureLayered]'s height is queried.
*/
func (class) _get_height(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Called when the number of layers in the [TextureLayered] is queried.
*/
func (class) _get_layers(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Called when the presence of mipmaps in the [TextureLayered] is queried.
*/
func (class) _has_mipmaps(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Called when the data for a layer in the [TextureLayered] is queried.
*/
func (class) _get_layer_data(impl func(ptr unsafe.Pointer, layer_index gd.Int) gdclass.Image, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var layer_index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, layer_index)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}

/*
Returns the current format being used by this texture. See [enum Image.Format] for details.
*/
//go:nosplit
func (self class) GetFormat() classdb.ImageFormat {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ImageFormat](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureLayered.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [TextureLayered]'s type. The type determines how the data is accessed, with cubemaps having special types.
*/
//go:nosplit
func (self class) GetLayeredType() classdb.TextureLayeredLayeredType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextureLayeredLayeredType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureLayered.Bind_get_layered_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the width of the texture in pixels. Width is typically represented by the X axis.
*/
//go:nosplit
func (self class) GetWidth() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureLayered.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the height of the texture in pixels. Height is typically represented by the Y axis.
*/
//go:nosplit
func (self class) GetHeight() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureLayered.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of referenced [Image]s.
*/
//go:nosplit
func (self class) GetLayers() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureLayered.Bind_get_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the layers have generated mipmaps.
*/
//go:nosplit
func (self class) HasMipmaps() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureLayered.Bind_has_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an [Image] resource with the data from specified [param layer].
*/
//go:nosplit
func (self class) GetLayerData(ctx gd.Lifetime, layer gd.Int) gdclass.Image {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextureLayered.Bind_get_layer_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsTextureLayered() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTextureLayered() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture() Texture.GD { return *((*Texture.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture() Texture.Go { return *((*Texture.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_format": return reflect.ValueOf(self._get_format);
	case "_get_layered_type": return reflect.ValueOf(self._get_layered_type);
	case "_get_width": return reflect.ValueOf(self._get_width);
	case "_get_height": return reflect.ValueOf(self._get_height);
	case "_get_layers": return reflect.ValueOf(self._get_layers);
	case "_has_mipmaps": return reflect.ValueOf(self._has_mipmaps);
	case "_get_layer_data": return reflect.ValueOf(self._get_layer_data);
	default: return gd.VirtualByName(self.AsTexture(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_get_format": return reflect.ValueOf(self._get_format);
	case "_get_layered_type": return reflect.ValueOf(self._get_layered_type);
	case "_get_width": return reflect.ValueOf(self._get_width);
	case "_get_height": return reflect.ValueOf(self._get_height);
	case "_get_layers": return reflect.ValueOf(self._get_layers);
	case "_has_mipmaps": return reflect.ValueOf(self._has_mipmaps);
	case "_get_layer_data": return reflect.ValueOf(self._get_layer_data);
	default: return gd.VirtualByName(self.AsTexture(), name)
	}
}
func init() {classdb.Register("TextureLayered", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type LayeredType = classdb.TextureLayeredLayeredType

const (
/*Texture is a generic [Texture2DArray].*/
	LayeredType2dArray LayeredType = 0
/*Texture is a [Cubemap], with each side in its own layer (6 in total).*/
	LayeredTypeCubemap LayeredType = 1
/*Texture is a [CubemapArray], with each cubemap being made of 6 layers.*/
	LayeredTypeCubemapArray LayeredType = 2
)
