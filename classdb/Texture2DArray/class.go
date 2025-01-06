// Package Texture2DArray provides methods for working with Texture2DArray object instances.
package Texture2DArray

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/ImageTextureLayered"
import "graphics.gd/classdb/TextureLayered"
import "graphics.gd/classdb/Texture"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A Texture2DArray is different from a Texture3D: The Texture2DArray does not support trilinear interpolation between the [Image]s, i.e. no blending. See also [Cubemap] and [CubemapArray], which are texture arrays with specialized cubemap functions.
A Texture2DArray is also different from an [AtlasTexture]: In a Texture2DArray, all images are treated separately. In an atlas, the regions (i.e. the single images) can be of different sizes. Furthermore, you usually need to add a padding around the regions, to prevent accidental UV mapping to more than one region. The same goes for mipmapping: Mipmap chains are handled separately for each layer. In an atlas, the slicing has to be done manually in the fragment shader.
To create such a texture file yourself, reimport your image files using the Godot Editor import presets.
*/
type Instance [1]gdclass.Texture2DArray
type Any interface {
	gd.IsClass
	AsTexture2DArray() Instance
}

/*
Creates a placeholder version of this resource ([PlaceholderTexture2DArray]).
*/
func (self Instance) CreatePlaceholder() [1]gdclass.Resource {
	return [1]gdclass.Resource(class(self).CreatePlaceholder())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Texture2DArray

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Texture2DArray"))
	casted := Instance{*(*gdclass.Texture2DArray)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Creates a placeholder version of this resource ([PlaceholderTexture2DArray]).
*/
//go:nosplit
func (self class) CreatePlaceholder() [1]gdclass.Resource {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2DArray.Bind_create_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Resource{gd.PointerWithOwnershipTransferredToGo[gdclass.Resource](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsTexture2DArray() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTexture2DArray() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsImageTextureLayered() ImageTextureLayered.Advanced {
	return *((*ImageTextureLayered.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsImageTextureLayered() ImageTextureLayered.Instance {
	return *((*ImageTextureLayered.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsTextureLayered() TextureLayered.Advanced {
	return *((*TextureLayered.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTextureLayered() TextureLayered.Instance {
	return *((*TextureLayered.Instance)(unsafe.Pointer(&self)))
}
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
	default:
		return gd.VirtualByName(ImageTextureLayered.Advanced(self.AsImageTextureLayered()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(ImageTextureLayered.Instance(self.AsImageTextureLayered()), name)
	}
}
func init() {
	gdclass.Register("Texture2DArray", func(ptr gd.Object) any {
		return [1]gdclass.Texture2DArray{*(*gdclass.Texture2DArray)(unsafe.Pointer(&ptr))}
	})
}
