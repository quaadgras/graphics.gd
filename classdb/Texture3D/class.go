// Package Texture3D provides methods for working with Texture3D object instances.
package Texture3D

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
import "graphics.gd/classdb/Texture"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any

/*
Base class for [ImageTexture3D] and [CompressedTexture3D]. Cannot be used directly, but contains all the functions necessary for accessing the derived resource types. [Texture3D] is the base class for all 3-dimensional texture types. See also [TextureLayered].
All images need to have the same width, height and number of mipmap levels.
To create such a texture file yourself, reimport your image files using the Godot Editor import presets.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=Texture3D)
*/
type Instance [1]gdclass.Texture3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTexture3D() Instance
}
type Interface interface {
	//Called when the [Texture3D]'s format is queried.
	GetFormat() gdclass.ImageFormat
	//Called when the [Texture3D]'s width is queried.
	GetWidth() int
	//Called when the [Texture3D]'s height is queried.
	GetHeight() int
	//Called when the [Texture3D]'s depth is queried.
	GetDepth() int
	//Called when the presence of mipmaps in the [Texture3D] is queried.
	HasMipmaps() bool
	//Called when the [Texture3D]'s data is queried.
	GetData() [][1]gdclass.Image
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetFormat() (_ gdclass.ImageFormat) { return }
func (self implementation) GetWidth() (_ int)                  { return }
func (self implementation) GetHeight() (_ int)                 { return }
func (self implementation) GetDepth() (_ int)                  { return }
func (self implementation) HasMipmaps() (_ bool)               { return }
func (self implementation) GetData() (_ [][1]gdclass.Image)    { return }

/*
Called when the [Texture3D]'s format is queried.
*/
func (Instance) _get_format(impl func(ptr unsafe.Pointer) gdclass.ImageFormat) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [Texture3D]'s width is queried.
*/
func (Instance) _get_width(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the [Texture3D]'s height is queried.
*/
func (Instance) _get_height(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the [Texture3D]'s depth is queried.
*/
func (Instance) _get_depth(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the presence of mipmaps in the [Texture3D] is queried.
*/
func (Instance) _has_mipmaps(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [Texture3D]'s data is queried.
*/
func (Instance) _get_data(impl func(ptr unsafe.Pointer) [][1]gdclass.Image) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[[1]gdclass.Image]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns the current format being used by this texture. See [enum Image.Format] for details.
*/
func (self Instance) GetFormat() gdclass.ImageFormat {
	return gdclass.ImageFormat(class(self).GetFormat())
}

/*
Returns the [Texture3D]'s width in pixels. Width is typically represented by the X axis.
*/
func (self Instance) GetWidth() int {
	return int(int(class(self).GetWidth()))
}

/*
Returns the [Texture3D]'s height in pixels. Width is typically represented by the Y axis.
*/
func (self Instance) GetHeight() int {
	return int(int(class(self).GetHeight()))
}

/*
Returns the [Texture3D]'s depth in pixels. Depth is typically represented by the Z axis (a dimension not present in [Texture2D]).
*/
func (self Instance) GetDepth() int {
	return int(int(class(self).GetDepth()))
}

/*
Returns [code]true[/code] if the [Texture3D] has generated mipmaps.
*/
func (self Instance) HasMipmaps() bool {
	return bool(class(self).HasMipmaps())
}

/*
Returns the [Texture3D]'s data as an array of [Image]s. Each [Image] represents a [i]slice[/i] of the [Texture3D], with different slices mapping to different depth (Z axis) levels.
*/
func (self Instance) GetData() [][1]gdclass.Image {
	return [][1]gdclass.Image(gd.ArrayAs[[][1]gdclass.Image](gd.InternalArray(class(self).GetData())))
}

/*
Creates a placeholder version of this resource ([PlaceholderTexture3D]).
*/
func (self Instance) CreatePlaceholder() [1]gdclass.Resource {
	return [1]gdclass.Resource(class(self).CreatePlaceholder())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Texture3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Texture3D"))
	casted := Instance{*(*gdclass.Texture3D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Called when the [Texture3D]'s format is queried.
*/
func (class) _get_format(impl func(ptr unsafe.Pointer) gdclass.ImageFormat) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [Texture3D]'s width is queried.
*/
func (class) _get_width(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [Texture3D]'s height is queried.
*/
func (class) _get_height(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [Texture3D]'s depth is queried.
*/
func (class) _get_depth(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the presence of mipmaps in the [Texture3D] is queried.
*/
func (class) _has_mipmaps(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [Texture3D]'s data is queried.
*/
func (class) _get_data(impl func(ptr unsafe.Pointer) Array.Contains[[1]gdclass.Image]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalArray(ret))

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
func (self class) GetFormat() gdclass.ImageFormat {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.ImageFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture3D.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Texture3D]'s width in pixels. Width is typically represented by the X axis.
*/
//go:nosplit
func (self class) GetWidth() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture3D.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Texture3D]'s height in pixels. Width is typically represented by the Y axis.
*/
//go:nosplit
func (self class) GetHeight() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture3D.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Texture3D]'s depth in pixels. Depth is typically represented by the Z axis (a dimension not present in [Texture2D]).
*/
//go:nosplit
func (self class) GetDepth() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture3D.Bind_get_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the [Texture3D] has generated mipmaps.
*/
//go:nosplit
func (self class) HasMipmaps() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture3D.Bind_has_mipmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Texture3D]'s data as an array of [Image]s. Each [Image] represents a [i]slice[/i] of the [Texture3D], with different slices mapping to different depth (Z axis) levels.
*/
//go:nosplit
func (self class) GetData() Array.Contains[[1]gdclass.Image] {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture3D.Bind_get_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.Image]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Creates a placeholder version of this resource ([PlaceholderTexture3D]).
*/
//go:nosplit
func (self class) CreatePlaceholder() [1]gdclass.Resource {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture3D.Bind_create_placeholder, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Resource{gd.PointerWithOwnershipTransferredToGo[gdclass.Resource](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsTexture3D() Advanced       { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTexture3D() Instance    { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTexture() Texture.Advanced { return *((*Texture.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTexture() Texture.Instance {
	return *((*Texture.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_format":
		return reflect.ValueOf(self._get_format)
	case "_get_width":
		return reflect.ValueOf(self._get_width)
	case "_get_height":
		return reflect.ValueOf(self._get_height)
	case "_get_depth":
		return reflect.ValueOf(self._get_depth)
	case "_has_mipmaps":
		return reflect.ValueOf(self._has_mipmaps)
	case "_get_data":
		return reflect.ValueOf(self._get_data)
	default:
		return gd.VirtualByName(Texture.Advanced(self.AsTexture()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_format":
		return reflect.ValueOf(self._get_format)
	case "_get_width":
		return reflect.ValueOf(self._get_width)
	case "_get_height":
		return reflect.ValueOf(self._get_height)
	case "_get_depth":
		return reflect.ValueOf(self._get_depth)
	case "_has_mipmaps":
		return reflect.ValueOf(self._has_mipmaps)
	case "_get_data":
		return reflect.ValueOf(self._get_data)
	default:
		return gd.VirtualByName(Texture.Instance(self.AsTexture()), name)
	}
}
func init() {
	gdclass.Register("Texture3D", func(ptr gd.Object) any { return [1]gdclass.Texture3D{*(*gdclass.Texture3D)(unsafe.Pointer(&ptr))} })
}
