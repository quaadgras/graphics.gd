package PlaceholderCubemapArray

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/PlaceholderTextureLayered"
import "grow.graphics/gd/gdclass/TextureLayered"
import "grow.graphics/gd/gdclass/Texture"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class replaces a [CubemapArray] or a [CubemapArray]-derived class in 2 conditions:
- In dedicated server mode, where the image data shouldn't affect game logic. This allows reducing the exported PCK's size significantly.
- When the [CubemapArray]-derived class is missing, for example when using a different engine version.
[b]Note:[/b] This class is not intended for rendering or for use in shaders. Operations like calculating UV are not guaranteed to work.

*/
type Go [1]classdb.PlaceholderCubemapArray
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PlaceholderCubemapArray
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PlaceholderCubemapArray"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsPlaceholderCubemapArray() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPlaceholderCubemapArray() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsPlaceholderTextureLayered() PlaceholderTextureLayered.GD { return *((*PlaceholderTextureLayered.GD)(unsafe.Pointer(&self))) }
func (self Go) AsPlaceholderTextureLayered() PlaceholderTextureLayered.Go { return *((*PlaceholderTextureLayered.Go)(unsafe.Pointer(&self))) }
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
	default: return gd.VirtualByName(self.AsPlaceholderTextureLayered(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsPlaceholderTextureLayered(), name)
	}
}
func init() {classdb.Register("PlaceholderCubemapArray", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
