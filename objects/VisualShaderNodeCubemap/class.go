package VisualShaderNodeCubemap

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/VisualShaderNode"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Translated to [code]texture(cubemap, vec3)[/code] in the shader language. Returns a color vector and alpha channel as scalar.
*/
type Instance [1]classdb.VisualShaderNodeCubemap

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShaderNodeCubemap

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeCubemap"))
	return Instance{classdb.VisualShaderNodeCubemap(object)}
}

func (self Instance) Source() classdb.VisualShaderNodeCubemapSource {
	return classdb.VisualShaderNodeCubemapSource(class(self).GetSource())
}

func (self Instance) SetSource(value classdb.VisualShaderNodeCubemapSource) {
	class(self).SetSource(value)
}

func (self Instance) CubeMap() objects.Cubemap {
	return objects.Cubemap(class(self).GetCubeMap())
}

func (self Instance) SetCubeMap(value objects.Cubemap) {
	class(self).SetCubeMap(value)
}

func (self Instance) TextureType() classdb.VisualShaderNodeCubemapTextureType {
	return classdb.VisualShaderNodeCubemapTextureType(class(self).GetTextureType())
}

func (self Instance) SetTextureType(value classdb.VisualShaderNodeCubemapTextureType) {
	class(self).SetTextureType(value)
}

//go:nosplit
func (self class) SetSource(value classdb.VisualShaderNodeCubemapSource) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCubemap.Bind_set_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSource() classdb.VisualShaderNodeCubemapSource {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeCubemapSource](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCubemap.Bind_get_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCubeMap(value objects.Cubemap) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(value[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCubemap.Bind_set_cube_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCubeMap() objects.Cubemap {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCubemap.Bind_get_cube_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Cubemap{classdb.Cubemap(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureType(value classdb.VisualShaderNodeCubemapTextureType) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCubemap.Bind_set_texture_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureType() classdb.VisualShaderNodeCubemapTextureType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeCubemapTextureType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCubemap.Bind_get_texture_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeCubemap() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShaderNodeCubemap() Instance {
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}
func init() {
	classdb.Register("VisualShaderNodeCubemap", func(ptr gd.Object) any { return classdb.VisualShaderNodeCubemap(ptr) })
}

type Source = classdb.VisualShaderNodeCubemapSource

const (
	/*Use the [Cubemap] set via [member cube_map]. If this is set to [member source], the [code]samplerCube[/code] port is ignored.*/
	SourceTexture Source = 0
	/*Use the [Cubemap] sampler reference passed via the [code]samplerCube[/code] port. If this is set to [member source], the [member cube_map] texture is ignored.*/
	SourcePort Source = 1
	/*Represents the size of the [enum Source] enum.*/
	SourceMax Source = 2
)

type TextureType = classdb.VisualShaderNodeCubemapTextureType

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
