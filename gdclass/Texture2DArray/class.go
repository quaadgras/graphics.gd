package Texture2DArray

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/ImageTextureLayered"
import "grow.graphics/gd/gdclass/TextureLayered"
import "grow.graphics/gd/gdclass/Texture"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A Texture2DArray is different from a Texture3D: The Texture2DArray does not support trilinear interpolation between the [Image]s, i.e. no blending. See also [Cubemap] and [CubemapArray], which are texture arrays with specialized cubemap functions.
A Texture2DArray is also different from an [AtlasTexture]: In a Texture2DArray, all images are treated separately. In an atlas, the regions (i.e. the single images) can be of different sizes. Furthermore, you usually need to add a padding around the regions, to prevent accidental UV mapping to more than one region. The same goes for mipmapping: Mipmap chains are handled separately for each layer. In an atlas, the slicing has to be done manually in the fragment shader.
To create such a texture file yourself, reimport your image files using the Godot Editor import presets.

*/
type Go [1]classdb.Texture2DArray

/*
Creates a placeholder version of this resource ([PlaceholderTexture2DArray]).
*/
func (self Go) CreatePlaceholder() gdclass.Resource {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Resource(class(self).CreatePlaceholder(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Texture2DArray
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Texture2DArray"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Creates a placeholder version of this resource ([PlaceholderTexture2DArray]).
*/
//go:nosplit
func (self class) CreatePlaceholder(ctx gd.Lifetime) gdclass.Resource {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Texture2DArray.Bind_create_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Resource
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsTexture2DArray() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture2DArray() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsImageTextureLayered() ImageTextureLayered.GD { return *((*ImageTextureLayered.GD)(unsafe.Pointer(&self))) }
func (self Go) AsImageTextureLayered() ImageTextureLayered.Go { return *((*ImageTextureLayered.Go)(unsafe.Pointer(&self))) }
func (self class) AsTextureLayered() TextureLayered.GD { return *((*TextureLayered.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTextureLayered() TextureLayered.Go { return *((*TextureLayered.Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture() Texture.GD { return *((*Texture.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture() Texture.Go { return *((*Texture.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsImageTextureLayered(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsImageTextureLayered(), name)
	}
}
func init() {classdb.Register("Texture2DArray", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
