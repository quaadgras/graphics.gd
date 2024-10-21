package ImageTexture

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Texture2D"
import "grow.graphics/gd/object/Texture"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A [Texture2D] based on an [Image]. For an image to be displayed, an [ImageTexture] has to be created from it using the [method create_from_image] method:
[codeblock]
var image = Image.load_from_file("res://icon.svg")
var texture = ImageTexture.create_from_image(image)
$Sprite2D.texture = texture
[/codeblock]
This way, textures can be created at run-time by loading images both from within the editor and externally.
[b]Warning:[/b] Prefer to load imported textures with [method @GDScript.load] over loading them from within the filesystem dynamically with [method Image.load], as it may not work in exported projects:
[codeblock]
var texture = load("res://icon.svg")
$Sprite2D.texture = texture
[/codeblock]
This is because images have to be imported as a [CompressedTexture2D] first to be loaded with [method @GDScript.load]. If you'd still like to load an image file just like any other [Resource], import it as an [Image] resource instead, and then load it normally using the [method @GDScript.load] method.
[b]Note:[/b] The image can be retrieved from an imported texture using the [method Texture2D.get_image] method, which returns a copy of the image:
[codeblock]
var texture = load("res://icon.svg")
var image: Image = texture.get_image()
[/codeblock]
An [ImageTexture] is not meant to be operated from within the editor interface directly, and is mostly useful for rendering images on screen dynamically via code. If you need to generate images procedurally from within the editor, consider saving and importing images as custom texture resources implementing a new [EditorImportPlugin].
[b]Note:[/b] The maximum texture size is 16384×16384 pixels due to graphics hardware limitations.

*/
type Simple [1]classdb.ImageTexture
func (self Simple) CreateFromImage(image [1]classdb.Image) [1]classdb.ImageTexture {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.ImageTexture(Expert(self).CreateFromImage(gc, image))
}
func (self Simple) GetFormat() classdb.ImageFormat {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ImageFormat(Expert(self).GetFormat())
}
func (self Simple) SetImage(image [1]classdb.Image) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetImage(image)
}
func (self Simple) Update(image [1]classdb.Image) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Update(image)
}
func (self Simple) SetSizeOverride(size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSizeOverride(size)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ImageTexture
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Creates a new [ImageTexture] and initializes it by allocating and setting the data from an [Image].
*/
//go:nosplit
func (self class) CreateFromImage(ctx gd.Lifetime, image object.Image) object.ImageTexture {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(image[0].AsPointer())[0])
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.ImageTexture.Bind_create_from_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.ImageTexture
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the format of the texture, one of [enum Image.Format].
*/
//go:nosplit
func (self class) GetFormat() classdb.ImageFormat {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ImageFormat](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImageTexture.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Replaces the texture's data with a new [Image]. This will re-allocate new memory for the texture.
If you want to update the image, but don't need to change its parameters (format, size), use [method update] instead for better performance.
*/
//go:nosplit
func (self class) SetImage(image object.Image)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(image[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImageTexture.Bind_set_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Replaces the texture's data with a new [Image].
[b]Note:[/b] The texture has to be created using [method create_from_image] or initialized first with the [method set_image] method before it can be updated. The new image dimensions, format, and mipmaps configuration should match the existing texture's image configuration.
Use this method over [method set_image] if you need to update the texture frequently, which is faster than allocating additional memory for a new texture each time.
*/
//go:nosplit
func (self class) Update(image object.Image)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(image[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImageTexture.Bind_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Resizes the texture to the specified dimensions.
*/
//go:nosplit
func (self class) SetSizeOverride(size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImageTexture.Bind_set_size_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsImageTexture() Expert { return self[0].AsImageTexture() }


//go:nosplit
func (self Simple) AsImageTexture() Simple { return self[0].AsImageTexture() }


//go:nosplit
func (self class) AsTexture2D() Texture2D.Expert { return self[0].AsTexture2D() }


//go:nosplit
func (self Simple) AsTexture2D() Texture2D.Simple { return self[0].AsTexture2D() }


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
func init() {classdb.Register("ImageTexture", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
