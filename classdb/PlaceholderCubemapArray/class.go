// Package PlaceholderCubemapArray provides methods for working with PlaceholderCubemapArray object instances.
package PlaceholderCubemapArray

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/PlaceholderTextureLayered"
import "graphics.gd/classdb/TextureLayered"
import "graphics.gd/classdb/Texture"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This class replaces a [CubemapArray] or a [CubemapArray]-derived class in 2 conditions:
- In dedicated server mode, where the image data shouldn't affect game logic. This allows reducing the exported PCK's size significantly.
- When the [CubemapArray]-derived class is missing, for example when using a different engine version.
[b]Note:[/b] This class is not intended for rendering or for use in shaders. Operations like calculating UV are not guaranteed to work.
*/
type Instance [1]gdclass.PlaceholderCubemapArray
type Any interface {
	gd.IsClass
	AsPlaceholderCubemapArray() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PlaceholderCubemapArray

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PlaceholderCubemapArray"))
	return Instance{*(*gdclass.PlaceholderCubemapArray)(unsafe.Pointer(&object))}
}

func (self class) AsPlaceholderCubemapArray() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPlaceholderCubemapArray() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsPlaceholderTextureLayered(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsPlaceholderTextureLayered(), name)
	}
}
func init() {
	gdclass.Register("PlaceholderCubemapArray", func(ptr gd.Object) any {
		return [1]gdclass.PlaceholderCubemapArray{*(*gdclass.PlaceholderCubemapArray)(unsafe.Pointer(&ptr))}
	})
}
