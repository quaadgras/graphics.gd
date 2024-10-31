package ShaderMaterial

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Material"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
A material that uses a custom [Shader] program to render visual items (canvas items, meshes, skies, fog), or to process particles. Compared to other materials, [ShaderMaterial] gives deeper control over the generated shader code. For more information, see the shaders documentation index below.
Multiple [ShaderMaterial]s can use the same shader and configure different values for the shader uniforms.
[b]Note:[/b] For performance reasons, the [signal Resource.changed] signal is only emitted when the [member Resource.resource_name] changes. Only in editor, it is also emitted for [member shader] changes.
*/
type Instance [1]classdb.ShaderMaterial

/*
Changes the value set for this material of a uniform in the shader.
[b]Note:[/b] [param param] is case-sensitive and must match the name of the uniform in the code exactly (not the capitalized name in the inspector).
[b]Note:[/b] Changes to the shader uniform will be effective on all instances using this [ShaderMaterial]. To prevent this, use per-instance uniforms with [method GeometryInstance3D.set_instance_shader_parameter] or duplicate the [ShaderMaterial] resource using [method Resource.duplicate]. Per-instance uniforms allow for better shader reuse and are therefore faster, so they should be preferred over duplicating the [ShaderMaterial] when possible.
*/
func (self Instance) SetShaderParameter(param string, value gd.Variant) {
	class(self).SetShaderParameter(gd.NewStringName(param), value)
}

/*
Returns the current value set for this material of a uniform in the shader.
*/
func (self Instance) GetShaderParameter(param string) gd.Variant {
	return gd.Variant(class(self).GetShaderParameter(gd.NewStringName(param)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ShaderMaterial

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ShaderMaterial"))
	return Instance{classdb.ShaderMaterial(object)}
}

func (self Instance) Shader() gdclass.Shader {
	return gdclass.Shader(class(self).GetShader())
}

func (self Instance) SetShader(value gdclass.Shader) {
	class(self).SetShader(value)
}

//go:nosplit
func (self class) SetShader(shader gdclass.Shader) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shader[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShaderMaterial.Bind_set_shader, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetShader() gdclass.Shader {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShaderMaterial.Bind_get_shader, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Shader{classdb.Shader(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Changes the value set for this material of a uniform in the shader.
[b]Note:[/b] [param param] is case-sensitive and must match the name of the uniform in the code exactly (not the capitalized name in the inspector).
[b]Note:[/b] Changes to the shader uniform will be effective on all instances using this [ShaderMaterial]. To prevent this, use per-instance uniforms with [method GeometryInstance3D.set_instance_shader_parameter] or duplicate the [ShaderMaterial] resource using [method Resource.duplicate]. Per-instance uniforms allow for better shader reuse and are therefore faster, so they should be preferred over duplicating the [ShaderMaterial] when possible.
*/
//go:nosplit
func (self class) SetShaderParameter(param gd.StringName, value gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(param))
	callframe.Arg(frame, pointers.Get(value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShaderMaterial.Bind_set_shader_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the current value set for this material of a uniform in the shader.
*/
//go:nosplit
func (self class) GetShaderParameter(param gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(param))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShaderMaterial.Bind_get_shader_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsShaderMaterial() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsShaderMaterial() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsMaterial() Material.Advanced {
	return *((*Material.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsMaterial() Material.Instance {
	return *((*Material.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsMaterial(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsMaterial(), name)
	}
}
func init() {
	classdb.Register("ShaderMaterial", func(ptr gd.Object) any { return classdb.ShaderMaterial(ptr) })
}
