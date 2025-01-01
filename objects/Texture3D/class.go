package Texture3D

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
Base class for [ImageTexture3D] and [CompressedTexture3D]. Cannot be used directly, but contains all the functions necessary for accessing the derived resource types. [Texture3D] is the base class for all 3-dimensional texture types. See also [TextureLayered].
All images need to have the same width, height and number of mipmap levels.
To create such a texture file yourself, reimport your image files using the Godot Editor import presets.

	// Texture3D methods that can be overridden by a [Class] that extends it.
	type Texture3D interface {
		//Called when the [Texture3D]'s format is queried.
		GetFormat() classdb.ImageFormat
		//Called when the [Texture3D]'s width is queried.
		GetWidth() int
		//Called when the [Texture3D]'s height is queried.
		GetHeight() int
		//Called when the [Texture3D]'s depth is queried.
		GetDepth() int
		//Called when the presence of mipmaps in the [Texture3D] is queried.
		HasMipmaps() bool
		//Called when the [Texture3D]'s data is queried.
		GetData() gd.Array
	}
*/
type Instance [1]classdb.Texture3D
type Any interface {
	gd.IsClass
	AsTexture3D() Instance
}

/*
Called when the [Texture3D]'s format is queried.
*/
func (Instance) _get_format(impl func(ptr unsafe.Pointer) classdb.ImageFormat) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [Texture3D]'s width is queried.
*/
func (Instance) _get_width(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the [Texture3D]'s height is queried.
*/
func (Instance) _get_height(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the [Texture3D]'s depth is queried.
*/
func (Instance) _get_depth(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the presence of mipmaps in the [Texture3D] is queried.
*/
func (Instance) _has_mipmaps(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [Texture3D]'s data is queried.
*/
func (Instance) _get_data(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
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
Returns the current format being used by this texture. See [enum Image.Format] for details.
*/
func (self Instance) GetFormat() classdb.ImageFormat {
	return classdb.ImageFormat(class(self).GetFormat())
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
func (self Instance) GetData() gd.Array {
	return gd.Array(class(self).GetData())
}

/*
Creates a placeholder version of this resource ([PlaceholderTexture3D]).
*/
func (self Instance) CreatePlaceholder() objects.Resource {
	return objects.Resource(class(self).CreatePlaceholder())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Texture3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Texture3D"))
	return Instance{classdb.Texture3D(object)}
}

/*
Called when the [Texture3D]'s format is queried.
*/
func (class) _get_format(impl func(ptr unsafe.Pointer) classdb.ImageFormat) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [Texture3D]'s width is queried.
*/
func (class) _get_width(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [Texture3D]'s height is queried.
*/
func (class) _get_height(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [Texture3D]'s depth is queried.
*/
func (class) _get_depth(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the presence of mipmaps in the [Texture3D] is queried.
*/
func (class) _has_mipmaps(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [Texture3D]'s data is queried.
*/
func (class) _get_data(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
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
Returns the current format being used by this texture. See [enum Image.Format] for details.
*/
//go:nosplit
func (self class) GetFormat() classdb.ImageFormat {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ImageFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture3D.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture3D.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture3D.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture3D.Bind_get_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture3D.Bind_has_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Texture3D]'s data as an array of [Image]s. Each [Image] represents a [i]slice[/i] of the [Texture3D], with different slices mapping to different depth (Z axis) levels.
*/
//go:nosplit
func (self class) GetData() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture3D.Bind_get_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Creates a placeholder version of this resource ([PlaceholderTexture3D]).
*/
//go:nosplit
func (self class) CreatePlaceholder() objects.Resource {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture3D.Bind_create_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Resource{classdb.Resource(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

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
		return gd.VirtualByName(self.AsTexture(), name)
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
		return gd.VirtualByName(self.AsTexture(), name)
	}
}
func init() {
	classdb.Register("Texture3D", func(ptr gd.Object) any { return [1]classdb.Texture3D{classdb.Texture3D(ptr)} })
}
