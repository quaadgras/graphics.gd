// Package PlaceholderCubemap provides methods for working with PlaceholderCubemap object instances.
package PlaceholderCubemap

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
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/classdb/PlaceholderTextureLayered"
import "graphics.gd/classdb/TextureLayered"
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
var _ Callable.Function
var _ Dictionary.Any

/*
This class replaces a [Cubemap] or a [Cubemap]-derived class in 2 conditions:
- In dedicated server mode, where the image data shouldn't affect game logic. This allows reducing the exported PCK's size significantly.
- When the [Cubemap]-derived class is missing, for example when using a different engine version.
[b]Note:[/b] This class is not intended for rendering or for use in shaders. Operations like calculating UV are not guaranteed to work.
*/
type Instance [1]gdclass.PlaceholderCubemap

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPlaceholderCubemap() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PlaceholderCubemap

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PlaceholderCubemap"))
	casted := Instance{*(*gdclass.PlaceholderCubemap)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self class) AsPlaceholderCubemap() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPlaceholderCubemap() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPlaceholderTextureLayered() PlaceholderTextureLayered.Advanced {
	return *((*PlaceholderTextureLayered.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPlaceholderTextureLayered() PlaceholderTextureLayered.Instance {
	return *((*PlaceholderTextureLayered.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(PlaceholderTextureLayered.Advanced(self.AsPlaceholderTextureLayered()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PlaceholderTextureLayered.Instance(self.AsPlaceholderTextureLayered()), name)
	}
}
func init() {
	gdclass.Register("PlaceholderCubemap", func(ptr gd.Object) any {
		return [1]gdclass.PlaceholderCubemap{*(*gdclass.PlaceholderCubemap)(unsafe.Pointer(&ptr))}
	})
}
