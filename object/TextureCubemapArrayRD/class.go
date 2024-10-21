package TextureCubemapArrayRD

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/TextureLayeredRD"
import "grow.graphics/gd/object/TextureLayered"
import "grow.graphics/gd/object/Texture"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This texture class allows you to use a cubemap array texture created directly on the [RenderingDevice] as a texture for materials, meshes, etc.

*/
type Simple [1]classdb.TextureCubemapArrayRD
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.TextureCubemapArrayRD
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsTextureCubemapArrayRD() Expert { return self[0].AsTextureCubemapArrayRD() }


//go:nosplit
func (self Simple) AsTextureCubemapArrayRD() Simple { return self[0].AsTextureCubemapArrayRD() }


//go:nosplit
func (self class) AsTextureLayeredRD() TextureLayeredRD.Expert { return self[0].AsTextureLayeredRD() }


//go:nosplit
func (self Simple) AsTextureLayeredRD() TextureLayeredRD.Simple { return self[0].AsTextureLayeredRD() }


//go:nosplit
func (self class) AsTextureLayered() TextureLayered.Expert { return self[0].AsTextureLayered() }


//go:nosplit
func (self Simple) AsTextureLayered() TextureLayered.Simple { return self[0].AsTextureLayered() }


//go:nosplit
func (self class) AsTexture() Texture.Expert { return self[0].AsTexture() }


//go:nosplit
func (self Simple) AsTexture() Texture.Simple { return self[0].AsTexture() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("TextureCubemapArrayRD", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
