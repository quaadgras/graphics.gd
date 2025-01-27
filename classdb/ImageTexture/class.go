// Package ImageTexture provides methods for working with ImageTexture object instances.
package ImageTexture

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Packed"
import "graphics.gd/classdb/Texture2D"
import "graphics.gd/classdb/Texture"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2i"

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
var _ = slices.Delete[[]struct{}, struct{}]

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
type Instance [1]gdclass.ImageTexture

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsImageTexture() Instance
}

/*
Creates a new [ImageTexture] and initializes it by allocating and setting the data from an [Image].
*/
func CreateFromImage(image [1]gdclass.Image) [1]gdclass.ImageTexture { //gd:ImageTexture.create_from_image
	self := Instance{}
	return [1]gdclass.ImageTexture(class(self).CreateFromImage(image))
}

/*
Returns the format of the texture, one of [enum Image.Format].
*/
func (self Instance) GetFormat() gdclass.ImageFormat { //gd:ImageTexture.get_format
	return gdclass.ImageFormat(class(self).GetFormat())
}

/*
Replaces the texture's data with a new [Image]. This will re-allocate new memory for the texture.
If you want to update the image, but don't need to change its parameters (format, size), use [method update] instead for better performance.
*/
func (self Instance) SetImage(image [1]gdclass.Image) { //gd:ImageTexture.set_image
	class(self).SetImage(image)
}

/*
Replaces the texture's data with a new [Image].
[b]Note:[/b] The texture has to be created using [method create_from_image] or initialized first with the [method set_image] method before it can be updated. The new image dimensions, format, and mipmaps configuration should match the existing texture's image configuration.
Use this method over [method set_image] if you need to update the texture frequently, which is faster than allocating additional memory for a new texture each time.
*/
func (self Instance) Update(image [1]gdclass.Image) { //gd:ImageTexture.update
	class(self).Update(image)
}

/*
Resizes the texture to the specified dimensions.
*/
func (self Instance) SetSizeOverride(size Vector2i.XY) { //gd:ImageTexture.set_size_override
	class(self).SetSizeOverride(gd.Vector2i(size))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ImageTexture

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ImageTexture"))
	casted := Instance{*(*gdclass.ImageTexture)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Creates a new [ImageTexture] and initializes it by allocating and setting the data from an [Image].
*/
//go:nosplit
func (self class) CreateFromImage(image [1]gdclass.Image) [1]gdclass.ImageTexture { //gd:ImageTexture.create_from_image
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(image[0])[0])
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageTexture.Bind_create_from_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.ImageTexture{gd.PointerWithOwnershipTransferredToGo[gdclass.ImageTexture](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the format of the texture, one of [enum Image.Format].
*/
//go:nosplit
func (self class) GetFormat() gdclass.ImageFormat { //gd:ImageTexture.get_format
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.ImageFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageTexture.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Replaces the texture's data with a new [Image]. This will re-allocate new memory for the texture.
If you want to update the image, but don't need to change its parameters (format, size), use [method update] instead for better performance.
*/
//go:nosplit
func (self class) SetImage(image [1]gdclass.Image) { //gd:ImageTexture.set_image
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(image[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageTexture.Bind_set_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Replaces the texture's data with a new [Image].
[b]Note:[/b] The texture has to be created using [method create_from_image] or initialized first with the [method set_image] method before it can be updated. The new image dimensions, format, and mipmaps configuration should match the existing texture's image configuration.
Use this method over [method set_image] if you need to update the texture frequently, which is faster than allocating additional memory for a new texture each time.
*/
//go:nosplit
func (self class) Update(image [1]gdclass.Image) { //gd:ImageTexture.update
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(image[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageTexture.Bind_update, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Resizes the texture to the specified dimensions.
*/
//go:nosplit
func (self class) SetSizeOverride(size gd.Vector2i) { //gd:ImageTexture.set_size_override
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageTexture.Bind_set_size_override, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Texture2D.Advanced(self.AsTexture2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Texture2D.Instance(self.AsTexture2D()), name)
	}
}
func init() {
	gdclass.Register("ImageTexture", func(ptr gd.Object) any {
		return [1]gdclass.ImageTexture{*(*gdclass.ImageTexture)(unsafe.Pointer(&ptr))}
	})
}
