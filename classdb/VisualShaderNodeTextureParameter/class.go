// Package VisualShaderNodeTextureParameter provides methods for working with VisualShaderNodeTextureParameter object instances.
package VisualShaderNodeTextureParameter

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/classdb/VisualShaderNodeParameter"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
Performs a lookup operation on the texture provided as a uniform for the shader.
*/
type Instance [1]gdclass.VisualShaderNodeTextureParameter

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeTextureParameter() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeTextureParameter

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeTextureParameter"))
	casted := Instance{*(*gdclass.VisualShaderNodeTextureParameter)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) TextureType() gdclass.VisualShaderNodeTextureParameterTextureType {
	return gdclass.VisualShaderNodeTextureParameterTextureType(class(self).GetTextureType())
}

func (self Instance) SetTextureType(value gdclass.VisualShaderNodeTextureParameterTextureType) {
	class(self).SetTextureType(value)
}

func (self Instance) ColorDefault() gdclass.VisualShaderNodeTextureParameterColorDefault {
	return gdclass.VisualShaderNodeTextureParameterColorDefault(class(self).GetColorDefault())
}

func (self Instance) SetColorDefault(value gdclass.VisualShaderNodeTextureParameterColorDefault) {
	class(self).SetColorDefault(value)
}

func (self Instance) TextureFilter() gdclass.VisualShaderNodeTextureParameterTextureFilter {
	return gdclass.VisualShaderNodeTextureParameterTextureFilter(class(self).GetTextureFilter())
}

func (self Instance) SetTextureFilter(value gdclass.VisualShaderNodeTextureParameterTextureFilter) {
	class(self).SetTextureFilter(value)
}

func (self Instance) TextureRepeat() gdclass.VisualShaderNodeTextureParameterTextureRepeat {
	return gdclass.VisualShaderNodeTextureParameterTextureRepeat(class(self).GetTextureRepeat())
}

func (self Instance) SetTextureRepeat(value gdclass.VisualShaderNodeTextureParameterTextureRepeat) {
	class(self).SetTextureRepeat(value)
}

func (self Instance) TextureSource() gdclass.VisualShaderNodeTextureParameterTextureSource {
	return gdclass.VisualShaderNodeTextureParameterTextureSource(class(self).GetTextureSource())
}

func (self Instance) SetTextureSource(value gdclass.VisualShaderNodeTextureParameterTextureSource) {
	class(self).SetTextureSource(value)
}

//go:nosplit
func (self class) SetTextureType(atype gdclass.VisualShaderNodeTextureParameterTextureType) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_set_texture_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureType() gdclass.VisualShaderNodeTextureParameterTextureType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeTextureParameterTextureType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_get_texture_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColorDefault(color gdclass.VisualShaderNodeTextureParameterColorDefault) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_set_color_default, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetColorDefault() gdclass.VisualShaderNodeTextureParameterColorDefault {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeTextureParameterColorDefault](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_get_color_default, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureFilter(filter gdclass.VisualShaderNodeTextureParameterTextureFilter) {
	var frame = callframe.New()
	callframe.Arg(frame, filter)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_set_texture_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureFilter() gdclass.VisualShaderNodeTextureParameterTextureFilter {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeTextureParameterTextureFilter](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_get_texture_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureRepeat(repeat gdclass.VisualShaderNodeTextureParameterTextureRepeat) {
	var frame = callframe.New()
	callframe.Arg(frame, repeat)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_set_texture_repeat, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureRepeat() gdclass.VisualShaderNodeTextureParameterTextureRepeat {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeTextureParameterTextureRepeat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_get_texture_repeat, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureSource(source gdclass.VisualShaderNodeTextureParameterTextureSource) {
	var frame = callframe.New()
	callframe.Arg(frame, source)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_set_texture_source, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureSource() gdclass.VisualShaderNodeTextureParameterTextureSource {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeTextureParameterTextureSource](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_get_texture_source, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeTextureParameter() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeTextureParameter() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNodeParameter() VisualShaderNodeParameter.Advanced {
	return *((*VisualShaderNodeParameter.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeParameter() VisualShaderNodeParameter.Instance {
	return *((*VisualShaderNodeParameter.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNode() VisualShaderNode.Advanced {
	return *((*VisualShaderNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNode() VisualShaderNode.Instance {
	return *((*VisualShaderNode.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(VisualShaderNodeParameter.Advanced(self.AsVisualShaderNodeParameter()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNodeParameter.Instance(self.AsVisualShaderNodeParameter()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeTextureParameter", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeTextureParameter{*(*gdclass.VisualShaderNodeTextureParameter)(unsafe.Pointer(&ptr))}
	})
}

type TextureType = gdclass.VisualShaderNodeTextureParameterTextureType

const (
	/*No hints are added to the uniform declaration.*/
	TypeData TextureType = 0
	/*Adds [code]source_color[/code] as hint to the uniform declaration for proper sRGB to linear conversion.*/
	TypeColor TextureType = 1
	/*Adds [code]hint_normal[/code] as hint to the uniform declaration, which internally converts the texture for proper usage as normal map.*/
	TypeNormalMap TextureType = 2
	/*Adds [code]hint_anisotropy[/code] as hint to the uniform declaration to use for a flowmap.*/
	TypeAnisotropy TextureType = 3
	/*Represents the size of the [enum TextureType] enum.*/
	TypeMax TextureType = 4
)

type ColorDefault = gdclass.VisualShaderNodeTextureParameterColorDefault

const (
	/*Defaults to fully opaque white color.*/
	ColorDefaultWhite ColorDefault = 0
	/*Defaults to fully opaque black color.*/
	ColorDefaultBlack ColorDefault = 1
	/*Defaults to fully transparent black color.*/
	ColorDefaultTransparent ColorDefault = 2
	/*Represents the size of the [enum ColorDefault] enum.*/
	ColorDefaultMax ColorDefault = 3
)

type TextureFilter = gdclass.VisualShaderNodeTextureParameterTextureFilter

const (
	/*Sample the texture using the filter determined by the node this shader is attached to.*/
	FilterDefault TextureFilter = 0
	/*The texture filter reads from the nearest pixel only. This makes the texture look pixelated from up close, and grainy from a distance (due to mipmaps not being sampled).*/
	FilterNearest TextureFilter = 1
	/*The texture filter blends between the nearest 4 pixels. This makes the texture look smooth from up close, and grainy from a distance (due to mipmaps not being sampled).*/
	FilterLinear TextureFilter = 2
	/*The texture filter reads from the nearest pixel and blends between the nearest 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]). This makes the texture look pixelated from up close, and smooth from a distance.
	  Use this for non-pixel art textures that may be viewed at a low scale (e.g. due to [Camera2D] zoom or sprite scaling), as mipmaps are important to smooth out pixels that are smaller than on-screen pixels.*/
	FilterNearestMipmap TextureFilter = 3
	/*The texture filter blends between the nearest 4 pixels and between the nearest 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]). This makes the texture look smooth from up close, and smooth from a distance.
	  Use this for non-pixel art textures that may be viewed at a low scale (e.g. due to [Camera2D] zoom or sprite scaling), as mipmaps are important to smooth out pixels that are smaller than on-screen pixels.*/
	FilterLinearMipmap TextureFilter = 4
	/*The texture filter reads from the nearest pixel and blends between 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]) based on the angle between the surface and the camera view. This makes the texture look pixelated from up close, and smooth from a distance. Anisotropic filtering improves texture quality on surfaces that are almost in line with the camera, but is slightly slower. The anisotropic filtering level can be changed by adjusting [member ProjectSettings.rendering/textures/default_filters/anisotropic_filtering_level].
	  [b]Note:[/b] This texture filter is rarely useful in 2D projects. [constant FILTER_NEAREST_MIPMAP] is usually more appropriate in this case.*/
	FilterNearestMipmapAnisotropic TextureFilter = 5
	/*The texture filter blends between the nearest 4 pixels and blends between 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]) based on the angle between the surface and the camera view. This makes the texture look smooth from up close, and smooth from a distance. Anisotropic filtering improves texture quality on surfaces that are almost in line with the camera, but is slightly slower. The anisotropic filtering level can be changed by adjusting [member ProjectSettings.rendering/textures/default_filters/anisotropic_filtering_level].
	  [b]Note:[/b] This texture filter is rarely useful in 2D projects. [constant FILTER_LINEAR_MIPMAP] is usually more appropriate in this case.*/
	FilterLinearMipmapAnisotropic TextureFilter = 6
	/*Represents the size of the [enum TextureFilter] enum.*/
	FilterMax TextureFilter = 7
)

type TextureRepeat = gdclass.VisualShaderNodeTextureParameterTextureRepeat

const (
	/*Sample the texture using the repeat mode determined by the node this shader is attached to.*/
	RepeatDefault TextureRepeat = 0
	/*Texture will repeat normally.*/
	RepeatEnabled TextureRepeat = 1
	/*Texture will not repeat.*/
	RepeatDisabled TextureRepeat = 2
	/*Represents the size of the [enum TextureRepeat] enum.*/
	RepeatMax TextureRepeat = 3
)

type TextureSource = gdclass.VisualShaderNodeTextureParameterTextureSource

const (
	/*The texture source is not specified in the shader.*/
	SourceNone TextureSource = 0
	/*The texture source is the screen texture which captures all opaque objects drawn this frame.*/
	SourceScreen TextureSource = 1
	/*The texture source is the depth texture from the depth prepass.*/
	SourceDepth TextureSource = 2
	/*The texture source is the normal-roughness buffer from the depth prepass.*/
	SourceNormalRoughness TextureSource = 3
	/*Represents the size of the [enum TextureSource] enum.*/
	SourceMax TextureSource = 4
)
