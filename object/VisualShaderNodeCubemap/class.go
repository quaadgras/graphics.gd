package VisualShaderNodeCubemap

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VisualShaderNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Translated to [code]texture(cubemap, vec3)[/code] in the shader language. Returns a color vector and alpha channel as scalar.

*/
type Simple [1]classdb.VisualShaderNodeCubemap
func (self Simple) SetSource(value classdb.VisualShaderNodeCubemapSource) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSource(value)
}
func (self Simple) GetSource() classdb.VisualShaderNodeCubemapSource {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeCubemapSource(Expert(self).GetSource())
}
func (self Simple) SetCubeMap(value [1]classdb.Cubemap) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCubeMap(value)
}
func (self Simple) GetCubeMap() [1]classdb.Cubemap {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Cubemap(Expert(self).GetCubeMap(gc))
}
func (self Simple) SetTextureType(value classdb.VisualShaderNodeCubemapTextureType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTextureType(value)
}
func (self Simple) GetTextureType() classdb.VisualShaderNodeCubemapTextureType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeCubemapTextureType(Expert(self).GetTextureType())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeCubemap
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetSource(value classdb.VisualShaderNodeCubemapSource)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeCubemap.Bind_set_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSource() classdb.VisualShaderNodeCubemapSource {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeCubemapSource](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeCubemap.Bind_get_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCubeMap(value object.Cubemap)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(value[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeCubemap.Bind_set_cube_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCubeMap(ctx gd.Lifetime) object.Cubemap {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeCubemap.Bind_get_cube_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Cubemap
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextureType(value classdb.VisualShaderNodeCubemapTextureType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeCubemap.Bind_set_texture_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextureType() classdb.VisualShaderNodeCubemapTextureType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeCubemapTextureType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeCubemap.Bind_get_texture_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeCubemap() Expert { return self[0].AsVisualShaderNodeCubemap() }


//go:nosplit
func (self Simple) AsVisualShaderNodeCubemap() Simple { return self[0].AsVisualShaderNodeCubemap() }


//go:nosplit
func (self class) AsVisualShaderNode() VisualShaderNode.Expert { return self[0].AsVisualShaderNode() }


//go:nosplit
func (self Simple) AsVisualShaderNode() VisualShaderNode.Simple { return self[0].AsVisualShaderNode() }


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
func init() {classdb.Register("VisualShaderNodeCubemap", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
