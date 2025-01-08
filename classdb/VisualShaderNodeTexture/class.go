// Package VisualShaderNodeTexture provides methods for working with VisualShaderNodeTexture object instances.
package VisualShaderNodeTexture

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Performs a lookup operation on the provided texture, with support for multiple texture sources to choose from.
*/
type Instance [1]gdclass.VisualShaderNodeTexture

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeTexture() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeTexture

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeTexture"))
	casted := Instance{*(*gdclass.VisualShaderNodeTexture)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Source() gdclass.VisualShaderNodeTextureSource {
	return gdclass.VisualShaderNodeTextureSource(class(self).GetSource())
}

func (self Instance) SetSource(value gdclass.VisualShaderNodeTextureSource) {
	class(self).SetSource(value)
}

func (self Instance) Texture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture())
}

func (self Instance) SetTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(value)
}

func (self Instance) TextureType() gdclass.VisualShaderNodeTextureTextureType {
	return gdclass.VisualShaderNodeTextureTextureType(class(self).GetTextureType())
}

func (self Instance) SetTextureType(value gdclass.VisualShaderNodeTextureTextureType) {
	class(self).SetTextureType(value)
}

//go:nosplit
func (self class) SetSource(value gdclass.VisualShaderNodeTextureSource) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTexture.Bind_set_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSource() gdclass.VisualShaderNodeTextureSource {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeTextureSource](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTexture.Bind_get_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTexture(value [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(value[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTexture.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() [1]gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTexture.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureType(value gdclass.VisualShaderNodeTextureTextureType) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTexture.Bind_set_texture_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureType() gdclass.VisualShaderNodeTextureTextureType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeTextureTextureType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTexture.Bind_get_texture_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeTexture() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShaderNodeTexture() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(VisualShaderNode.Advanced(self.AsVisualShaderNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNode.Instance(self.AsVisualShaderNode()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeTexture", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeTexture{*(*gdclass.VisualShaderNodeTexture)(unsafe.Pointer(&ptr))}
	})
}

type Source = gdclass.VisualShaderNodeTextureSource

const (
	/*Use the texture given as an argument for this function.*/
	SourceTexture Source = 0
	/*Use the current viewport's texture as the source.*/
	SourceScreen Source = 1
	/*Use the texture from this shader's texture built-in (e.g. a texture of a [Sprite2D]).*/
	Source2dTexture Source = 2
	/*Use the texture from this shader's normal map built-in.*/
	Source2dNormal Source = 3
	/*Use the depth texture captured during the depth prepass. Only available when the depth prepass is used (i.e. in spatial shaders and in the forward_plus or gl_compatibility renderers).*/
	SourceDepth Source = 4
	/*Use the texture provided in the input port for this function.*/
	SourcePort Source = 5
	/*Use the normal buffer captured during the depth prepass. Only available when the normal-roughness buffer is available (i.e. in spatial shaders and in the forward_plus renderer).*/
	Source3dNormal Source = 6
	/*Use the roughness buffer captured during the depth prepass. Only available when the normal-roughness buffer is available (i.e. in spatial shaders and in the forward_plus renderer).*/
	SourceRoughness Source = 7
	/*Represents the size of the [enum Source] enum.*/
	SourceMax Source = 8
)

type TextureType = gdclass.VisualShaderNodeTextureTextureType

const (
	/*No hints are added to the uniform declaration.*/
	TypeData TextureType = 0
	/*Adds [code]source_color[/code] as hint to the uniform declaration for proper sRGB to linear conversion.*/
	TypeColor TextureType = 1
	/*Adds [code]hint_normal[/code] as hint to the uniform declaration, which internally converts the texture for proper usage as normal map.*/
	TypeNormalMap TextureType = 2
	/*Represents the size of the [enum TextureType] enum.*/
	TypeMax TextureType = 3
)
