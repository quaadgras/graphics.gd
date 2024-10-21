package ShaderMaterial

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Material"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A material that uses a custom [Shader] program to render visual items (canvas items, meshes, skies, fog), or to process particles. Compared to other materials, [ShaderMaterial] gives deeper control over the generated shader code. For more information, see the shaders documentation index below.
Multiple [ShaderMaterial]s can use the same shader and configure different values for the shader uniforms.
[b]Note:[/b] For performance reasons, the [signal Resource.changed] signal is only emitted when the [member Resource.resource_name] changes. Only in editor, it is also emitted for [member shader] changes.

*/
type Simple [1]classdb.ShaderMaterial
func (self Simple) SetShader(shader [1]classdb.Shader) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShader(shader)
}
func (self Simple) GetShader() [1]classdb.Shader {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Shader(Expert(self).GetShader(gc))
}
func (self Simple) SetShaderParameter(param string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShaderParameter(gc.StringName(param), value)
}
func (self Simple) GetShaderParameter(param string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetShaderParameter(gc, gc.StringName(param)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ShaderMaterial
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetShader(shader object.Shader)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(shader[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShaderMaterial.Bind_set_shader, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShader(ctx gd.Lifetime) object.Shader {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShaderMaterial.Bind_get_shader, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Shader
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Changes the value set for this material of a uniform in the shader.
[b]Note:[/b] [param param] is case-sensitive and must match the name of the uniform in the code exactly (not the capitalized name in the inspector).
[b]Note:[/b] Changes to the shader uniform will be effective on all instances using this [ShaderMaterial]. To prevent this, use per-instance uniforms with [method GeometryInstance3D.set_instance_shader_parameter] or duplicate the [ShaderMaterial] resource using [method Resource.duplicate]. Per-instance uniforms allow for better shader reuse and are therefore faster, so they should be preferred over duplicating the [ShaderMaterial] when possible.
*/
//go:nosplit
func (self class) SetShaderParameter(param gd.StringName, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(param))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShaderMaterial.Bind_set_shader_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the current value set for this material of a uniform in the shader.
*/
//go:nosplit
func (self class) GetShaderParameter(ctx gd.Lifetime, param gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(param))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShaderMaterial.Bind_get_shader_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsShaderMaterial() Expert { return self[0].AsShaderMaterial() }


//go:nosplit
func (self Simple) AsShaderMaterial() Simple { return self[0].AsShaderMaterial() }


//go:nosplit
func (self class) AsMaterial() Material.Expert { return self[0].AsMaterial() }


//go:nosplit
func (self Simple) AsMaterial() Material.Simple { return self[0].AsMaterial() }


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
func init() {classdb.Register("ShaderMaterial", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
