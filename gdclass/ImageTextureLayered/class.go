package ImageTextureLayered

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/TextureLayered"
import "grow.graphics/gd/gdclass/Texture"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Base class for [Texture2DArray], [Cubemap] and [CubemapArray]. Cannot be used directly, but contains all the functions necessary for accessing the derived resource types. See also [Texture3D].
*/
type Instance [1]classdb.ImageTextureLayered

/*
Creates an [ImageTextureLayered] from an array of [Image]s. See [method Image.create] for the expected data format. The first image decides the width, height, image format and mipmapping setting. The other images [i]must[/i] have the same width, height, image format and mipmapping setting.
Each [Image] represents one [code]layer[/code].
*/
func (self Instance) CreateFromImages(images gd.Array) gd.Error {
	return gd.Error(class(self).CreateFromImages(images))
}

/*
Replaces the existing [Image] data at the given [param layer] with this new image.
The given [Image] must have the same width, height, image format, and mipmapping flag as the rest of the referenced images.
If the image format is unsupported, it will be decompressed and converted to a similar and supported [enum Image.Format].
The update is immediate: it's synchronized with drawing.
*/
func (self Instance) UpdateLayer(image gdclass.Image, layer int) {
	class(self).UpdateLayer(image, gd.Int(layer))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ImageTextureLayered

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ImageTextureLayered"))
	return Instance{classdb.ImageTextureLayered(object)}
}

/*
Creates an [ImageTextureLayered] from an array of [Image]s. See [method Image.create] for the expected data format. The first image decides the width, height, image format and mipmapping setting. The other images [i]must[/i] have the same width, height, image format and mipmapping setting.
Each [Image] represents one [code]layer[/code].
*/
//go:nosplit
func (self class) CreateFromImages(images gd.Array) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(images))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageTextureLayered.Bind_create_from_images, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Replaces the existing [Image] data at the given [param layer] with this new image.
The given [Image] must have the same width, height, image format, and mipmapping flag as the rest of the referenced images.
If the image format is unsupported, it will be decompressed and converted to a similar and supported [enum Image.Format].
The update is immediate: it's synchronized with drawing.
*/
//go:nosplit
func (self class) UpdateLayer(image gdclass.Image, layer gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(image[0])[0])
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageTextureLayered.Bind_update_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsImageTextureLayered() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsImageTextureLayered() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsTextureLayered(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTextureLayered(), name)
	}
}
func init() {
	classdb.Register("ImageTextureLayered", func(ptr gd.Object) any { return classdb.ImageTextureLayered(ptr) })
}
