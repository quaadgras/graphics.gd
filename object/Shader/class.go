package Shader

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A custom shader program implemented in the Godot shading language, saved with the [code].gdshader[/code] extension.
This class is used by a [ShaderMaterial] and allows you to write your own custom behavior for rendering visual items or updating particle information. For a detailed explanation and usage, please see the tutorials linked below.

*/
type Simple [1]classdb.Shader
func (self Simple) GetMode() classdb.ShaderMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ShaderMode(Expert(self).GetMode())
}
func (self Simple) SetCode(code string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCode(gc.String(code))
}
func (self Simple) GetCode() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetCode(gc).String())
}
func (self Simple) SetDefaultTextureParameter(name string, texture [1]classdb.Texture2D, index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDefaultTextureParameter(gc.StringName(name), texture, gd.Int(index))
}
func (self Simple) GetDefaultTextureParameter(name string, index int) [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetDefaultTextureParameter(gc, gc.StringName(name), gd.Int(index)))
}
func (self Simple) GetShaderUniformList(get_groups bool) gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).GetShaderUniformList(gc, get_groups))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Shader
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the shader mode for the shader.
*/
//go:nosplit
func (self class) GetMode() classdb.ShaderMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ShaderMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shader.Bind_get_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCode(code gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(code))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shader.Bind_set_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCode(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shader.Bind_get_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the default texture to be used with a texture uniform. The default is used if a texture is not set in the [ShaderMaterial].
[b]Note:[/b] [param name] must match the name of the uniform in the code exactly.
[b]Note:[/b] If the sampler array is used use [param index] to access the specified texture.
*/
//go:nosplit
func (self class) SetDefaultTextureParameter(name gd.StringName, texture object.Texture2D, index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shader.Bind_set_default_texture_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the texture that is set as default for the specified parameter.
[b]Note:[/b] [param name] must match the name of the uniform in the code exactly.
[b]Note:[/b] If the sampler array is used use [param index] to access the specified texture.
*/
//go:nosplit
func (self class) GetDefaultTextureParameter(ctx gd.Lifetime, name gd.StringName, index gd.Int) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shader.Bind_get_default_texture_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Get the list of shader uniforms that can be assigned to a [ShaderMaterial], for use with [method ShaderMaterial.set_shader_parameter] and [method ShaderMaterial.get_shader_parameter]. The parameters returned are contained in dictionaries in a similar format to the ones returned by [method Object.get_property_list].
If argument [param get_groups] is true, parameter grouping hints will be provided.
*/
//go:nosplit
func (self class) GetShaderUniformList(ctx gd.Lifetime, get_groups bool) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, get_groups)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shader.Bind_get_shader_uniform_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsShader() Expert { return self[0].AsShader() }


//go:nosplit
func (self Simple) AsShader() Simple { return self[0].AsShader() }


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
func init() {classdb.Register("Shader", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Mode = classdb.ShaderMode

const (
/*Mode used to draw all 3D objects.*/
	ModeSpatial Mode = 0
/*Mode used to draw all 2D objects.*/
	ModeCanvasItem Mode = 1
/*Mode used to calculate particle information on a per-particle basis. Not used for drawing.*/
	ModeParticles Mode = 2
/*Mode used for drawing skies. Only works with shaders attached to [Sky] objects.*/
	ModeSky Mode = 3
/*Mode used for setting the color and density of volumetric fog effect.*/
	ModeFog Mode = 4
)
