// Package Cubemap provides methods for working with Cubemap object instances.
package Cubemap

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/ImageTextureLayered"
import "graphics.gd/classdb/Resource"
import "graphics.gd/classdb/Texture"
import "graphics.gd/classdb/TextureLayered"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

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
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
A cubemap is made of 6 textures organized in layers. They are typically used for faking reflections in 3D rendering (see [ReflectionProbe]). It can be used to make an object look as if it's reflecting its surroundings. This usually delivers much better performance than other reflection methods.
This resource is typically used as a uniform in custom shaders. Few core Godot methods make use of [Cubemap] resources.
To create such a texture file yourself, reimport your image files using the Godot Editor import presets. To create a Cubemap from code, use [method ImageTextureLayered.create_from_images] on an instance of the Cubemap class.
The expected image order is X+, X-, Y+, Y-, Z+, Z- (in Godot's coordinate system, so Y+ is "up" and Z- is "forward"). You can use one of the following templates as a base:
- [url=https://raw.githubusercontent.com/godotengine/godot-docs/master/tutorials/assets_pipeline/img/cubemap_template_2x3.webp]2×3 cubemap template (default layout option)[/url]
- [url=https://raw.githubusercontent.com/godotengine/godot-docs/master/tutorials/assets_pipeline/img/cubemap_template_3x2.webp]3×2 cubemap template[/url]
- [url=https://raw.githubusercontent.com/godotengine/godot-docs/master/tutorials/assets_pipeline/img/cubemap_template_1x6.webp]1×6 cubemap template[/url]
- [url=https://raw.githubusercontent.com/godotengine/godot-docs/master/tutorials/assets_pipeline/img/cubemap_template_6x1.webp]6×1 cubemap template[/url]
[b]Note:[/b] Godot doesn't support using cubemaps in a [PanoramaSkyMaterial]. To use a cubemap as a skybox, convert the default [PanoramaSkyMaterial] to a [ShaderMaterial] using the [b]Convert to ShaderMaterial[/b] resource dropdown option, then replace its code with the following:
[codeblock lang=text]
shader_type sky;

uniform samplerCube source_panorama : filter_linear, source_color, hint_default_black;
uniform float exposure : hint_range(0, 128) = 1.0;

	void sky() {
	    // If importing a cubemap from another engine, you may need to flip one of the `EYEDIR` components below
	    // by replacing it with `-EYEDIR`.
	    vec3 eyedir = vec3(EYEDIR.x, EYEDIR.y, EYEDIR.z);
	    COLOR = texture(source_panorama, eyedir).rgb * exposure;
	}

[/codeblock]
After replacing the shader code and saving, specify the imported Cubemap resource in the Shader Parameters section of the ShaderMaterial in the inspector.
Alternatively, you can use [url=https://danilw.github.io/GLSL-howto/cubemap_to_panorama_js/cubemap_to_panorama.html]this tool[/url] to convert a cubemap to an equirectangular sky map and use [PanoramaSkyMaterial] as usual.
*/
type Instance [1]gdclass.Cubemap

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCubemap() Instance
}

/*
Creates a placeholder version of this resource ([PlaceholderCubemap]).
*/
func (self Instance) CreatePlaceholder() [1]gdclass.Resource { //gd:Cubemap.create_placeholder
	return [1]gdclass.Resource(Advanced(self).CreatePlaceholder())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Cubemap

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Cubemap"))
	casted := Instance{*(*gdclass.Cubemap)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Creates a placeholder version of this resource ([PlaceholderCubemap]).
*/
//go:nosplit
func (self class) CreatePlaceholder() [1]gdclass.Resource { //gd:Cubemap.create_placeholder
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Cubemap.Bind_create_placeholder, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Resource{gd.PointerWithOwnershipTransferredToGo[gdclass.Resource](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsCubemap() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCubemap() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsImageTextureLayered() ImageTextureLayered.Advanced {
	return *((*ImageTextureLayered.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsImageTextureLayered() ImageTextureLayered.Instance {
	return *((*ImageTextureLayered.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(ImageTextureLayered.Advanced(self.AsImageTextureLayered()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(ImageTextureLayered.Instance(self.AsImageTextureLayered()), name)
	}
}
func init() {
	gdclass.Register("Cubemap", func(ptr gd.Object) any { return [1]gdclass.Cubemap{*(*gdclass.Cubemap)(unsafe.Pointer(&ptr))} })
}
