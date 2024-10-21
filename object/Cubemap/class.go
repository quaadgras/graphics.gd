package Cubemap

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/ImageTextureLayered"
import "grow.graphics/gd/object/TextureLayered"
import "grow.graphics/gd/object/Texture"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A cubemap is made of 6 textures organized in layers. They are typically used for faking reflections in 3D rendering (see [ReflectionProbe]). It can be used to make an object look as if it's reflecting its surroundings. This usually delivers much better performance than other reflection methods.
This resource is typically used as a uniform in custom shaders. Few core Godot methods make use of [Cubemap] resources.
To create such a texture file yourself, reimport your image files using the Godot Editor import presets.
[b]Note:[/b] Godot doesn't support using cubemaps in a [PanoramaSkyMaterial]. You can use [url=https://danilw.github.io/GLSL-howto/cubemap_to_panorama_js/cubemap_to_panorama.html]this tool[/url] to convert a cubemap to an equirectangular sky map.

*/
type Simple [1]classdb.Cubemap
func (self Simple) CreatePlaceholder() [1]classdb.Resource {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Resource(Expert(self).CreatePlaceholder(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Cubemap
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Creates a placeholder version of this resource ([PlaceholderCubemap]).
*/
//go:nosplit
func (self class) CreatePlaceholder(ctx gd.Lifetime) object.Resource {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Cubemap.Bind_create_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Resource
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsCubemap() Expert { return self[0].AsCubemap() }


//go:nosplit
func (self Simple) AsCubemap() Simple { return self[0].AsCubemap() }


//go:nosplit
func (self class) AsImageTextureLayered() ImageTextureLayered.Expert { return self[0].AsImageTextureLayered() }


//go:nosplit
func (self Simple) AsImageTextureLayered() ImageTextureLayered.Simple { return self[0].AsImageTextureLayered() }


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
func init() {classdb.Register("Cubemap", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
