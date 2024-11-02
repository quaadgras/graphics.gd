package Texture2DArrayRD

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/TextureLayeredRD"
import "grow.graphics/gd/objects/TextureLayered"
import "grow.graphics/gd/objects/Texture"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This texture array class allows you to use a 2D array texture created directly on the [RenderingDevice] as a texture for materials, meshes, etc.
*/
type Instance [1]classdb.Texture2DArrayRD

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Texture2DArrayRD

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Texture2DArrayRD"))
	return Instance{classdb.Texture2DArrayRD(object)}
}

func (self class) AsTexture2DArrayRD() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTexture2DArrayRD() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTextureLayeredRD() TextureLayeredRD.Advanced {
	return *((*TextureLayeredRD.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTextureLayeredRD() TextureLayeredRD.Instance {
	return *((*TextureLayeredRD.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsTextureLayeredRD(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTextureLayeredRD(), name)
	}
}
func init() {
	classdb.Register("Texture2DArrayRD", func(ptr gd.Object) any { return classdb.Texture2DArrayRD(ptr) })
}
