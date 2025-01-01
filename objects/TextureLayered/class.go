package TextureLayered

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Texture"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

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
		GetLayerData(layer_index int) objects.Image
	}
*/
type Instance [1]classdb.TextureLayered
type Any interface {
	gd.IsClass
	AsTextureLayered() Instance
}

/*
Called when the [TextureLayered]'s format is queried.
*/
func (Instance) _get_format(impl func(ptr unsafe.Pointer) classdb.ImageFormat) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the layers' type in the [TextureLayered] is queried.
*/
func (Instance) _get_layered_type(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the [TextureLayered]'s width queried.
*/
func (Instance) _get_width(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the [TextureLayered]'s height is queried.
*/
func (Instance) _get_height(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the number of layers in the [TextureLayered] is queried.
*/
func (Instance) _get_layers(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the presence of mipmaps in the [TextureLayered] is queried.
*/
func (Instance) _has_mipmaps(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the data for a layer in the [TextureLayered] is queried.
*/
func (Instance) _get_layer_data(impl func(ptr unsafe.Pointer, layer_index int) objects.Image) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var layer_index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(layer_index))
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns the current format being used by this texture. See [enum Image.Format] for details.
*/
func (self Instance) GetFormat() classdb.ImageFormat {
	return classdb.ImageFormat(class(self).GetFormat())
}

/*
Returns the [TextureLayered]'s type. The type determines how the data is accessed, with cubemaps having special types.
*/
func (self Instance) GetLayeredType() classdb.TextureLayeredLayeredType {
	return classdb.TextureLayeredLayeredType(class(self).GetLayeredType())
}

/*
Returns the width of the texture in pixels. Width is typically represented by the X axis.
*/
func (self Instance) GetWidth() int {
	return int(int(class(self).GetWidth()))
}

/*
Returns the height of the texture in pixels. Height is typically represented by the Y axis.
*/
func (self Instance) GetHeight() int {
	return int(int(class(self).GetHeight()))
}

/*
Returns the number of referenced [Image]s.
*/
func (self Instance) GetLayers() int {
	return int(int(class(self).GetLayers()))
}

/*
Returns [code]true[/code] if the layers have generated mipmaps.
*/
func (self Instance) HasMipmaps() bool {
	return bool(class(self).HasMipmaps())
}

/*
Returns an [Image] resource with the data from specified [param layer].
*/
func (self Instance) GetLayerData(layer int) objects.Image {
	return objects.Image(class(self).GetLayerData(gd.Int(layer)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.TextureLayered

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TextureLayered"))
	return Instance{classdb.TextureLayered(object)}
}

/*
Called when the [TextureLayered]'s format is queried.
*/
func (class) _get_format(impl func(ptr unsafe.Pointer) classdb.ImageFormat) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the layers' type in the [TextureLayered] is queried.
*/
func (class) _get_layered_type(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [TextureLayered]'s width queried.
*/
func (class) _get_width(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [TextureLayered]'s height is queried.
*/
func (class) _get_height(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the number of layers in the [TextureLayered] is queried.
*/
func (class) _get_layers(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the presence of mipmaps in the [TextureLayered] is queried.
*/
func (class) _has_mipmaps(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the data for a layer in the [TextureLayered] is queried.
*/
func (class) _get_layer_data(impl func(ptr unsafe.Pointer, layer_index gd.Int) objects.Image) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var layer_index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, layer_index)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns the current format being used by this texture. See [enum Image.Format] for details.
*/
//go:nosplit
func (self class) GetFormat() classdb.ImageFormat {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ImageFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureLayered.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [TextureLayered]'s type. The type determines how the data is accessed, with cubemaps having special types.
*/
//go:nosplit
func (self class) GetLayeredType() classdb.TextureLayeredLayeredType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextureLayeredLayeredType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureLayered.Bind_get_layered_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the width of the texture in pixels. Width is typically represented by the X axis.
*/
//go:nosplit
func (self class) GetWidth() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureLayered.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the height of the texture in pixels. Height is typically represented by the Y axis.
*/
//go:nosplit
func (self class) GetHeight() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureLayered.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of referenced [Image]s.
*/
//go:nosplit
func (self class) GetLayers() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureLayered.Bind_get_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the layers have generated mipmaps.
*/
//go:nosplit
func (self class) HasMipmaps() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureLayered.Bind_has_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an [Image] resource with the data from specified [param layer].
*/
//go:nosplit
func (self class) GetLayerData(layer gd.Int) objects.Image {
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureLayered.Bind_get_layer_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Image{classdb.Image(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsTextureLayered() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTextureLayered() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTexture() Texture.Advanced   { return *((*Texture.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTexture() Texture.Instance {
	return *((*Texture.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_format":
		return reflect.ValueOf(self._get_format)
	case "_get_layered_type":
		return reflect.ValueOf(self._get_layered_type)
	case "_get_width":
		return reflect.ValueOf(self._get_width)
	case "_get_height":
		return reflect.ValueOf(self._get_height)
	case "_get_layers":
		return reflect.ValueOf(self._get_layers)
	case "_has_mipmaps":
		return reflect.ValueOf(self._has_mipmaps)
	case "_get_layer_data":
		return reflect.ValueOf(self._get_layer_data)
	default:
		return gd.VirtualByName(self.AsTexture(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_format":
		return reflect.ValueOf(self._get_format)
	case "_get_layered_type":
		return reflect.ValueOf(self._get_layered_type)
	case "_get_width":
		return reflect.ValueOf(self._get_width)
	case "_get_height":
		return reflect.ValueOf(self._get_height)
	case "_get_layers":
		return reflect.ValueOf(self._get_layers)
	case "_has_mipmaps":
		return reflect.ValueOf(self._has_mipmaps)
	case "_get_layer_data":
		return reflect.ValueOf(self._get_layer_data)
	default:
		return gd.VirtualByName(self.AsTexture(), name)
	}
}
func init() {
	classdb.Register("TextureLayered", func(ptr gd.Object) any { return [1]classdb.TextureLayered{classdb.TextureLayered(ptr)} })
}

type LayeredType = classdb.TextureLayeredLayeredType

const (
	/*Texture is a generic [Texture2DArray].*/
	LayeredType2dArray LayeredType = 0
	/*Texture is a [Cubemap], with each side in its own layer (6 in total).*/
	LayeredTypeCubemap LayeredType = 1
	/*Texture is a [CubemapArray], with each cubemap being made of 6 layers.*/
	LayeredTypeCubemapArray LayeredType = 2
)
