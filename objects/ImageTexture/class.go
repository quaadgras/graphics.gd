package ImageTexture

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Texture2D"
import "graphics.gd/objects/Texture"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Vector2i"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

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
type Instance [1]classdb.ImageTexture
type Any interface {
	gd.IsClass
	AsImageTexture() Instance
}

/*
Creates a new [ImageTexture] and initializes it by allocating and setting the data from an [Image].
*/
func CreateFromImage(image objects.Image) objects.ImageTexture {
	self := Instance{}
	return objects.ImageTexture(class(self).CreateFromImage(image))
}

/*
Returns the format of the texture, one of [enum Image.Format].
*/
func (self Instance) GetFormat() classdb.ImageFormat {
	return classdb.ImageFormat(class(self).GetFormat())
}

/*
Replaces the texture's data with a new [Image]. This will re-allocate new memory for the texture.
If you want to update the image, but don't need to change its parameters (format, size), use [method update] instead for better performance.
*/
func (self Instance) SetImage(image objects.Image) {
	class(self).SetImage(image)
}

/*
Replaces the texture's data with a new [Image].
[b]Note:[/b] The texture has to be created using [method create_from_image] or initialized first with the [method set_image] method before it can be updated. The new image dimensions, format, and mipmaps configuration should match the existing texture's image configuration.
Use this method over [method set_image] if you need to update the texture frequently, which is faster than allocating additional memory for a new texture each time.
*/
func (self Instance) Update(image objects.Image) {
	class(self).Update(image)
}

/*
Resizes the texture to the specified dimensions.
*/
func (self Instance) SetSizeOverride(size Vector2i.XY) {
	class(self).SetSizeOverride(gd.Vector2i(size))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ImageTexture

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ImageTexture"))
	return Instance{classdb.ImageTexture(object)}
}

/*
Creates a new [ImageTexture] and initializes it by allocating and setting the data from an [Image].
*/
//go:nosplit
func (self class) CreateFromImage(image objects.Image) objects.ImageTexture {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(image[0])[0])
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageTexture.Bind_create_from_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.ImageTexture{classdb.ImageTexture(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the format of the texture, one of [enum Image.Format].
*/
//go:nosplit
func (self class) GetFormat() classdb.ImageFormat {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ImageFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageTexture.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Replaces the texture's data with a new [Image]. This will re-allocate new memory for the texture.
If you want to update the image, but don't need to change its parameters (format, size), use [method update] instead for better performance.
*/
//go:nosplit
func (self class) SetImage(image objects.Image) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(image[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageTexture.Bind_set_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Replaces the texture's data with a new [Image].
[b]Note:[/b] The texture has to be created using [method create_from_image] or initialized first with the [method set_image] method before it can be updated. The new image dimensions, format, and mipmaps configuration should match the existing texture's image configuration.
Use this method over [method set_image] if you need to update the texture frequently, which is faster than allocating additional memory for a new texture each time.
*/
//go:nosplit
func (self class) Update(image objects.Image) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(image[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageTexture.Bind_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Resizes the texture to the specified dimensions.
*/
//go:nosplit
func (self class) SetSizeOverride(size gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageTexture.Bind_set_size_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsImageTexture() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsImageTexture() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTexture2D() Texture2D.Advanced {
	return *((*Texture2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTexture2D() Texture2D.Instance {
	return *((*Texture2D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsTexture2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTexture2D(), name)
	}
}
func init() {
	classdb.Register("ImageTexture", func(ptr gd.Object) any { return [1]classdb.ImageTexture{classdb.ImageTexture(ptr)} })
}
