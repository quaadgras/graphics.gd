package VisualShaderNodeTextureParameter

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/VisualShaderNodeParameter"
import "graphics.gd/objects/VisualShaderNode"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Performs a lookup operation on the texture provided as a uniform for the shader.
*/
type Instance [1]classdb.VisualShaderNodeTextureParameter
type Any interface {
	gd.IsClass
	AsVisualShaderNodeTextureParameter() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShaderNodeTextureParameter

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeTextureParameter"))
	return Instance{classdb.VisualShaderNodeTextureParameter(object)}
}

func (self Instance) TextureType() classdb.VisualShaderNodeTextureParameterTextureType {
	return classdb.VisualShaderNodeTextureParameterTextureType(class(self).GetTextureType())
}

func (self Instance) SetTextureType(value classdb.VisualShaderNodeTextureParameterTextureType) {
	class(self).SetTextureType(value)
}

func (self Instance) ColorDefault() classdb.VisualShaderNodeTextureParameterColorDefault {
	return classdb.VisualShaderNodeTextureParameterColorDefault(class(self).GetColorDefault())
}

func (self Instance) SetColorDefault(value classdb.VisualShaderNodeTextureParameterColorDefault) {
	class(self).SetColorDefault(value)
}

func (self Instance) TextureFilter() classdb.VisualShaderNodeTextureParameterTextureFilter {
	return classdb.VisualShaderNodeTextureParameterTextureFilter(class(self).GetTextureFilter())
}

func (self Instance) SetTextureFilter(value classdb.VisualShaderNodeTextureParameterTextureFilter) {
	class(self).SetTextureFilter(value)
}

func (self Instance) TextureRepeat() classdb.VisualShaderNodeTextureParameterTextureRepeat {
	return classdb.VisualShaderNodeTextureParameterTextureRepeat(class(self).GetTextureRepeat())
}

func (self Instance) SetTextureRepeat(value classdb.VisualShaderNodeTextureParameterTextureRepeat) {
	class(self).SetTextureRepeat(value)
}

func (self Instance) TextureSource() classdb.VisualShaderNodeTextureParameterTextureSource {
	return classdb.VisualShaderNodeTextureParameterTextureSource(class(self).GetTextureSource())
}

func (self Instance) SetTextureSource(value classdb.VisualShaderNodeTextureParameterTextureSource) {
	class(self).SetTextureSource(value)
}

//go:nosplit
func (self class) SetTextureType(atype classdb.VisualShaderNodeTextureParameterTextureType) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_set_texture_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureType() classdb.VisualShaderNodeTextureParameterTextureType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeTextureParameterTextureType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_get_texture_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColorDefault(color classdb.VisualShaderNodeTextureParameterColorDefault) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_set_color_default, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetColorDefault() classdb.VisualShaderNodeTextureParameterColorDefault {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeTextureParameterColorDefault](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_get_color_default, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureFilter(filter classdb.VisualShaderNodeTextureParameterTextureFilter) {
	var frame = callframe.New()
	callframe.Arg(frame, filter)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_set_texture_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureFilter() classdb.VisualShaderNodeTextureParameterTextureFilter {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeTextureParameterTextureFilter](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_get_texture_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureRepeat(repeat classdb.VisualShaderNodeTextureParameterTextureRepeat) {
	var frame = callframe.New()
	callframe.Arg(frame, repeat)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_set_texture_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureRepeat() classdb.VisualShaderNodeTextureParameterTextureRepeat {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeTextureParameterTextureRepeat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_get_texture_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureSource(source classdb.VisualShaderNodeTextureParameterTextureSource) {
	var frame = callframe.New()
	callframe.Arg(frame, source)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_set_texture_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureSource() classdb.VisualShaderNodeTextureParameterTextureSource {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeTextureParameterTextureSource](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTextureParameter.Bind_get_texture_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNodeParameter(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNodeParameter(), name)
	}
}
func init() {
	classdb.Register("VisualShaderNodeTextureParameter", func(ptr gd.Object) any {
		return [1]classdb.VisualShaderNodeTextureParameter{classdb.VisualShaderNodeTextureParameter(ptr)}
	})
}

type TextureType = classdb.VisualShaderNodeTextureParameterTextureType

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

type ColorDefault = classdb.VisualShaderNodeTextureParameterColorDefault

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

type TextureFilter = classdb.VisualShaderNodeTextureParameterTextureFilter

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

type TextureRepeat = classdb.VisualShaderNodeTextureParameterTextureRepeat

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

type TextureSource = classdb.VisualShaderNodeTextureParameterTextureSource

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
