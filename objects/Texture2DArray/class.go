package Texture2DArray

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/ImageTextureLayered"
import "graphics.gd/objects/TextureLayered"
import "graphics.gd/objects/Texture"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A Texture2DArray is different from a Texture3D: The Texture2DArray does not support trilinear interpolation between the [Image]s, i.e. no blending. See also [Cubemap] and [CubemapArray], which are texture arrays with specialized cubemap functions.
A Texture2DArray is also different from an [AtlasTexture]: In a Texture2DArray, all images are treated separately. In an atlas, the regions (i.e. the single images) can be of different sizes. Furthermore, you usually need to add a padding around the regions, to prevent accidental UV mapping to more than one region. The same goes for mipmapping: Mipmap chains are handled separately for each layer. In an atlas, the slicing has to be done manually in the fragment shader.
To create such a texture file yourself, reimport your image files using the Godot Editor import presets.
*/
type Instance [1]classdb.Texture2DArray
type Any interface {
	gd.IsClass
	AsTexture2DArray() Instance
}

/*
Creates a placeholder version of this resource ([PlaceholderTexture2DArray]).
*/
func (self Instance) CreatePlaceholder() objects.Resource {
	return objects.Resource(class(self).CreatePlaceholder())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Texture2DArray

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Texture2DArray"))
	return Instance{*(*classdb.Texture2DArray)(unsafe.Pointer(&object))}
}

/*
Creates a placeholder version of this resource ([PlaceholderTexture2DArray]).
*/
//go:nosplit
func (self class) CreatePlaceholder() objects.Resource {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2DArray.Bind_create_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Resource{gd.PointerWithOwnershipTransferredToGo[classdb.Resource](r_ret.Get())}
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsImageTextureLayered(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsImageTextureLayered(), name)
	}
}
func init() {
	classdb.Register("Texture2DArray", func(ptr gd.Object) any {
		return [1]classdb.Texture2DArray{*(*classdb.Texture2DArray)(unsafe.Pointer(&ptr))}
	})
}
