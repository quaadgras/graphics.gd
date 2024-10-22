package TextureCubemapRD

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/TextureLayeredRD"
import "grow.graphics/gd/gdclass/TextureLayered"
import "grow.graphics/gd/gdclass/Texture"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This texture class allows you to use a cubemap texture created directly on the [RenderingDevice] as a texture for materials, meshes, etc.

*/
type Go [1]classdb.TextureCubemapRD
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TextureCubemapRD
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("TextureCubemapRD"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsTextureCubemapRD() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTextureCubemapRD() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTextureLayeredRD() TextureLayeredRD.GD { return *((*TextureLayeredRD.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTextureLayeredRD() TextureLayeredRD.Go { return *((*TextureLayeredRD.Go)(unsafe.Pointer(&self))) }
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
	default: return gd.VirtualByName(self.AsTextureLayeredRD(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTextureLayeredRD(), name)
	}
}
func init() {classdb.Register("TextureCubemapRD", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
