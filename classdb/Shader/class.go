// Package Shader provides methods for working with Shader object instances.
package Shader

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
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
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
var _ Dictionary.Any
var _ RID.Any

/*
A custom shader program implemented in the Godot shading language, saved with the [code].gdshader[/code] extension.
This class is used by a [ShaderMaterial] and allows you to write your own custom behavior for rendering visual items or updating particle information. For a detailed explanation and usage, please see the tutorials linked below.
*/
type Instance [1]gdclass.Shader

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsShader() Instance
}

/*
Returns the shader mode for the shader.
*/
func (self Instance) GetMode() gdclass.ShaderMode { //gd:Shader.get_mode
	return gdclass.ShaderMode(class(self).GetMode())
}

/*
Sets the default texture to be used with a texture uniform. The default is used if a texture is not set in the [ShaderMaterial].
[b]Note:[/b] [param name] must match the name of the uniform in the code exactly.
[b]Note:[/b] If the sampler array is used use [param index] to access the specified texture.
*/
func (self Instance) SetDefaultTextureParameter(name string, texture [1]gdclass.Texture2D) { //gd:Shader.set_default_texture_parameter
	class(self).SetDefaultTextureParameter(gd.NewStringName(name), texture, gd.Int(0))
}

/*
Returns the texture that is set as default for the specified parameter.
[b]Note:[/b] [param name] must match the name of the uniform in the code exactly.
[b]Note:[/b] If the sampler array is used use [param index] to access the specified texture.
*/
func (self Instance) GetDefaultTextureParameter(name string) [1]gdclass.Texture2D { //gd:Shader.get_default_texture_parameter
	return [1]gdclass.Texture2D(class(self).GetDefaultTextureParameter(gd.NewStringName(name), gd.Int(0)))
}

/*
Get the list of shader uniforms that can be assigned to a [ShaderMaterial], for use with [method ShaderMaterial.set_shader_parameter] and [method ShaderMaterial.get_shader_parameter]. The parameters returned are contained in dictionaries in a similar format to the ones returned by [method Object.get_property_list].
If argument [param get_groups] is true, parameter grouping hints will be provided.
*/
func (self Instance) GetShaderUniformList() []any { //gd:Shader.get_shader_uniform_list
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).GetShaderUniformList(false))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Shader

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Shader"))
	casted := Instance{*(*gdclass.Shader)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Code() string {
	return string(class(self).GetCode().String())
}

func (self Instance) SetCode(value string) {
	class(self).SetCode(gd.NewString(value))
}

/*
Returns the shader mode for the shader.
*/
//go:nosplit
func (self class) GetMode() gdclass.ShaderMode { //gd:Shader.get_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.ShaderMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shader.Bind_get_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCode(code gd.String) { //gd:Shader.set_code
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(code))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shader.Bind_set_code, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCode() gd.String { //gd:Shader.get_code
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shader.Bind_get_code, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the default texture to be used with a texture uniform. The default is used if a texture is not set in the [ShaderMaterial].
[b]Note:[/b] [param name] must match the name of the uniform in the code exactly.
[b]Note:[/b] If the sampler array is used use [param index] to access the specified texture.
*/
//go:nosplit
func (self class) SetDefaultTextureParameter(name gd.StringName, texture [1]gdclass.Texture2D, index gd.Int) { //gd:Shader.set_default_texture_parameter
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shader.Bind_set_default_texture_parameter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the texture that is set as default for the specified parameter.
[b]Note:[/b] [param name] must match the name of the uniform in the code exactly.
[b]Note:[/b] If the sampler array is used use [param index] to access the specified texture.
*/
//go:nosplit
func (self class) GetDefaultTextureParameter(name gd.StringName, index gd.Int) [1]gdclass.Texture2D { //gd:Shader.get_default_texture_parameter
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shader.Bind_get_default_texture_parameter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Get the list of shader uniforms that can be assigned to a [ShaderMaterial], for use with [method ShaderMaterial.set_shader_parameter] and [method ShaderMaterial.get_shader_parameter]. The parameters returned are contained in dictionaries in a similar format to the ones returned by [method Object.get_property_list].
If argument [param get_groups] is true, parameter grouping hints will be provided.
*/
//go:nosplit
func (self class) GetShaderUniformList(get_groups bool) Array.Any { //gd:Shader.get_shader_uniform_list
	var frame = callframe.New()
	callframe.Arg(frame, get_groups)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shader.Bind_get_shader_uniform_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}
func (self class) AsShader() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsShader() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("Shader", func(ptr gd.Object) any { return [1]gdclass.Shader{*(*gdclass.Shader)(unsafe.Pointer(&ptr))} })
}

type Mode = gdclass.ShaderMode //gd:Shader.Mode

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
