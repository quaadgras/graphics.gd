package CubemapArray

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
[CubemapArray]s are made of an array of [Cubemap]s. Like [Cubemap]s, they are made of multiple textures, the amount of which must be divisible by 6 (one for each face of the cube).
The primary benefit of [CubemapArray]s is that they can be accessed in shader code using a single texture reference. In other words, you can pass multiple [Cubemap]s into a shader using a single [CubemapArray]. [Cubemap]s are allocated in adjacent cache regions on the GPU, which makes [CubemapArray]s the most efficient way to store multiple [Cubemap]s.
[b]Note:[/b] Godot uses [CubemapArray]s internally for many effects, including the [Sky] if you set [member ProjectSettings.rendering/reflections/sky_reflections/texture_array_reflections] to [code]true[/code]. To create such a texture file yourself, reimport your image files using the import presets of the File System dock.
[b]Note:[/b] [CubemapArray] is not supported in the OpenGL 3 rendering backend.

*/
type Simple [1]classdb.CubemapArray
func (self Simple) CreatePlaceholder() [1]classdb.Resource {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Resource(Expert(self).CreatePlaceholder(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.CubemapArray
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Creates a placeholder version of this resource ([PlaceholderCubemapArray]).
*/
//go:nosplit
func (self class) CreatePlaceholder(ctx gd.Lifetime) object.Resource {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CubemapArray.Bind_create_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Resource
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsCubemapArray() Expert { return self[0].AsCubemapArray() }


//go:nosplit
func (self Simple) AsCubemapArray() Simple { return self[0].AsCubemapArray() }


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
func init() {classdb.Register("CubemapArray", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
