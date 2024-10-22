package Texture3D

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
		GetData() gd.ArrayOf[gdclass.Image]
	}

*/
type Go [1]classdb.Texture3D

/*
Called when the [Texture3D]'s format is queried.
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
Called when the [Texture3D]'s width is queried.
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
Called when the [Texture3D]'s height is queried.
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
Called when the [Texture3D]'s depth is queried.
*/
func (Go) _get_depth(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Called when the presence of mipmaps in the [Texture3D] is queried.
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
Called when the [Texture3D]'s data is queried.
*/
func (Go) _get_data(impl func(ptr unsafe.Pointer) gd.ArrayOf[gdclass.Image], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
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
Returns the [Texture3D]'s width in pixels. Width is typically represented by the X axis.
*/
func (self Go) GetWidth() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetWidth()))
}

/*
Returns the [Texture3D]'s height in pixels. Width is typically represented by the Y axis.
*/
func (self Go) GetHeight() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetHeight()))
}

/*
Returns the [Texture3D]'s depth in pixels. Depth is typically represented by the Z axis (a dimension not present in [Texture2D]).
*/
func (self Go) GetDepth() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetDepth()))
}

/*
Returns [code]true[/code] if the [Texture3D] has generated mipmaps.
*/
func (self Go) HasMipmaps() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasMipmaps())
}

/*
Returns the [Texture3D]'s data as an array of [Image]s. Each [Image] represents a [i]slice[/i] of the [Texture3D], with different slices mapping to different depth (Z axis) levels.
*/
func (self Go) GetData() gd.ArrayOf[gdclass.Image] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gdclass.Image](class(self).GetData(gc))
}

/*
Creates a placeholder version of this resource ([PlaceholderTexture3D]).
*/
func (self Go) CreatePlaceholder() gdclass.Resource {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Resource(class(self).CreatePlaceholder(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Texture3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Texture3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Called when the [Texture3D]'s format is queried.
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
Called when the [Texture3D]'s width is queried.
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
Called when the [Texture3D]'s height is queried.
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
Called when the [Texture3D]'s depth is queried.
*/
func (class) _get_depth(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Called when the presence of mipmaps in the [Texture3D] is queried.
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
Called when the [Texture3D]'s data is queried.
*/
func (class) _get_data(impl func(ptr unsafe.Pointer) gd.ArrayOf[gdclass.Image], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Texture3D.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Texture3D]'s width in pixels. Width is typically represented by the X axis.
*/
//go:nosplit
func (self class) GetWidth() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Texture3D.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Texture3D]'s height in pixels. Width is typically represented by the Y axis.
*/
//go:nosplit
func (self class) GetHeight() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Texture3D.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Texture3D]'s depth in pixels. Depth is typically represented by the Z axis (a dimension not present in [Texture2D]).
*/
//go:nosplit
func (self class) GetDepth() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Texture3D.Bind_get_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the [Texture3D] has generated mipmaps.
*/
//go:nosplit
func (self class) HasMipmaps() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Texture3D.Bind_has_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Texture3D]'s data as an array of [Image]s. Each [Image] represents a [i]slice[/i] of the [Texture3D], with different slices mapping to different depth (Z axis) levels.
*/
//go:nosplit
func (self class) GetData(ctx gd.Lifetime) gd.ArrayOf[gdclass.Image] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Texture3D.Bind_get_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.Image](ret)
}
/*
Creates a placeholder version of this resource ([PlaceholderTexture3D]).
*/
//go:nosplit
func (self class) CreatePlaceholder(ctx gd.Lifetime) gdclass.Resource {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Texture3D.Bind_create_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Resource
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsTexture3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture() Texture.GD { return *((*Texture.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture() Texture.Go { return *((*Texture.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_format": return reflect.ValueOf(self._get_format);
	case "_get_width": return reflect.ValueOf(self._get_width);
	case "_get_height": return reflect.ValueOf(self._get_height);
	case "_get_depth": return reflect.ValueOf(self._get_depth);
	case "_has_mipmaps": return reflect.ValueOf(self._has_mipmaps);
	case "_get_data": return reflect.ValueOf(self._get_data);
	default: return gd.VirtualByName(self.AsTexture(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_get_format": return reflect.ValueOf(self._get_format);
	case "_get_width": return reflect.ValueOf(self._get_width);
	case "_get_height": return reflect.ValueOf(self._get_height);
	case "_get_depth": return reflect.ValueOf(self._get_depth);
	case "_has_mipmaps": return reflect.ValueOf(self._has_mipmaps);
	case "_get_data": return reflect.ValueOf(self._get_data);
	default: return gd.VirtualByName(self.AsTexture(), name)
	}
}
func init() {classdb.Register("Texture3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
