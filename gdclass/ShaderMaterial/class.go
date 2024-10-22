package ShaderMaterial

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Material"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A material that uses a custom [Shader] program to render visual items (canvas items, meshes, skies, fog), or to process particles. Compared to other materials, [ShaderMaterial] gives deeper control over the generated shader code. For more information, see the shaders documentation index below.
Multiple [ShaderMaterial]s can use the same shader and configure different values for the shader uniforms.
[b]Note:[/b] For performance reasons, the [signal Resource.changed] signal is only emitted when the [member Resource.resource_name] changes. Only in editor, it is also emitted for [member shader] changes.

*/
type Go [1]classdb.ShaderMaterial

/*
Changes the value set for this material of a uniform in the shader.
[b]Note:[/b] [param param] is case-sensitive and must match the name of the uniform in the code exactly (not the capitalized name in the inspector).
[b]Note:[/b] Changes to the shader uniform will be effective on all instances using this [ShaderMaterial]. To prevent this, use per-instance uniforms with [method GeometryInstance3D.set_instance_shader_parameter] or duplicate the [ShaderMaterial] resource using [method Resource.duplicate]. Per-instance uniforms allow for better shader reuse and are therefore faster, so they should be preferred over duplicating the [ShaderMaterial] when possible.
*/
func (self Go) SetShaderParameter(param string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShaderParameter(gc.StringName(param), value)
}

/*
Returns the current value set for this material of a uniform in the shader.
*/
func (self Go) GetShaderParameter(param string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).GetShaderParameter(gc, gc.StringName(param)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ShaderMaterial
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ShaderMaterial"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Shader() gdclass.Shader {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Shader(class(self).GetShader(gc))
}

func (self Go) SetShader(value gdclass.Shader) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShader(value)
}

//go:nosplit
func (self class) SetShader(shader gdclass.Shader)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(shader[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShaderMaterial.Bind_set_shader, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShader(ctx gd.Lifetime) gdclass.Shader {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ShaderMaterial.Bind_get_shader, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Shader
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
func (self class) AsShaderMaterial() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsShaderMaterial() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsMaterial() Material.GD { return *((*Material.GD)(unsafe.Pointer(&self))) }
func (self Go) AsMaterial() Material.Go { return *((*Material.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMaterial(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMaterial(), name)
	}
}
func init() {classdb.Register("ShaderMaterial", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
