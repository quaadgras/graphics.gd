package PlaceholderCubemapArray

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/PlaceholderTextureLayered"
import "grow.graphics/gd/object/TextureLayered"
import "grow.graphics/gd/object/Texture"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class replaces a [CubemapArray] or a [CubemapArray]-derived class in 2 conditions:
- In dedicated server mode, where the image data shouldn't affect game logic. This allows reducing the exported PCK's size significantly.
- When the [CubemapArray]-derived class is missing, for example when using a different engine version.
[b]Note:[/b] This class is not intended for rendering or for use in shaders. Operations like calculating UV are not guaranteed to work.

*/
type Simple [1]classdb.PlaceholderCubemapArray
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PlaceholderCubemapArray
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsPlaceholderCubemapArray() Expert { return self[0].AsPlaceholderCubemapArray() }


//go:nosplit
func (self Simple) AsPlaceholderCubemapArray() Simple { return self[0].AsPlaceholderCubemapArray() }


//go:nosplit
func (self class) AsPlaceholderTextureLayered() PlaceholderTextureLayered.Expert { return self[0].AsPlaceholderTextureLayered() }


//go:nosplit
func (self Simple) AsPlaceholderTextureLayered() PlaceholderTextureLayered.Simple { return self[0].AsPlaceholderTextureLayered() }


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

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("PlaceholderCubemapArray", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
