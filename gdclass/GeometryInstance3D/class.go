package GeometryInstance3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualInstance3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Base node for geometry-based visual instances. Shares some common functionality like visibility and custom materials.

*/
type Go [1]classdb.GeometryInstance3D

/*
Set the value of a shader uniform for this instance only ([url=$DOCS_URL/tutorials/shaders/shader_reference/shading_language.html#per-instance-uniforms]per-instance uniform[/url]). See also [method ShaderMaterial.set_shader_parameter] to assign a uniform on all instances using the same [ShaderMaterial].
[b]Note:[/b] For a shader uniform to be assignable on a per-instance basis, it [i]must[/i] be defined with [code]instance uniform ...[/code] rather than [code]uniform ...[/code] in the shader code.
[b]Note:[/b] [param name] is case-sensitive and must match the name of the uniform in the code exactly (not the capitalized name in the inspector).
[b]Note:[/b] Per-instance shader uniforms are currently only available in 3D, so there is no 2D equivalent of this method.
*/
func (self Go) SetInstanceShaderParameter(name string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInstanceShaderParameter(gc.StringName(name), value)
}

/*
Get the value of a shader parameter as set on this instance.
*/
func (self Go) GetInstanceShaderParameter(name string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).GetInstanceShaderParameter(gc, gc.StringName(name)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GeometryInstance3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("GeometryInstance3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) MaterialOverride() gdclass.Material {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Material(class(self).GetMaterialOverride(gc))
}

func (self Go) SetMaterialOverride(value gdclass.Material) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaterialOverride(value)
}

func (self Go) MaterialOverlay() gdclass.Material {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Material(class(self).GetMaterialOverlay(gc))
}

func (self Go) SetMaterialOverlay(value gdclass.Material) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaterialOverlay(value)
}

func (self Go) Transparency() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTransparency()))
}

func (self Go) SetTransparency(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTransparency(gd.Float(value))
}

func (self Go) CastShadow() classdb.GeometryInstance3DShadowCastingSetting {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.GeometryInstance3DShadowCastingSetting(class(self).GetCastShadowsSetting())
}

func (self Go) SetCastShadow(value classdb.GeometryInstance3DShadowCastingSetting) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCastShadowsSetting(value)
}

func (self Go) ExtraCullMargin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetExtraCullMargin()))
}

func (self Go) SetExtraCullMargin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetExtraCullMargin(gd.Float(value))
}

func (self Go) CustomAabb() gd.AABB {
	gc := gd.GarbageCollector(); _ = gc
		return gd.AABB(class(self).GetCustomAabb())
}

func (self Go) SetCustomAabb(value gd.AABB) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCustomAabb(value)
}

func (self Go) LodBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetLodBias()))
}

func (self Go) SetLodBias(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLodBias(gd.Float(value))
}

func (self Go) IgnoreOcclusionCulling() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsIgnoringOcclusionCulling())
}

func (self Go) SetIgnoreOcclusionCulling(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIgnoreOcclusionCulling(value)
}

func (self Go) GiMode() classdb.GeometryInstance3DGIMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.GeometryInstance3DGIMode(class(self).GetGiMode())
}

func (self Go) SetGiMode(value classdb.GeometryInstance3DGIMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGiMode(value)
}

func (self Go) GiLightmapScale() classdb.GeometryInstance3DLightmapScale {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.GeometryInstance3DLightmapScale(class(self).GetLightmapScale())
}

func (self Go) SetGiLightmapScale(value classdb.GeometryInstance3DLightmapScale) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLightmapScale(value)
}

func (self Go) VisibilityRangeBegin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVisibilityRangeBegin()))
}

func (self Go) SetVisibilityRangeBegin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibilityRangeBegin(gd.Float(value))
}

func (self Go) VisibilityRangeBeginMargin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVisibilityRangeBeginMargin()))
}

func (self Go) SetVisibilityRangeBeginMargin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibilityRangeBeginMargin(gd.Float(value))
}

func (self Go) VisibilityRangeEnd() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVisibilityRangeEnd()))
}

func (self Go) SetVisibilityRangeEnd(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibilityRangeEnd(gd.Float(value))
}

func (self Go) VisibilityRangeEndMargin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVisibilityRangeEndMargin()))
}

func (self Go) SetVisibilityRangeEndMargin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibilityRangeEndMargin(gd.Float(value))
}

func (self Go) VisibilityRangeFadeMode() classdb.GeometryInstance3DVisibilityRangeFadeMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.GeometryInstance3DVisibilityRangeFadeMode(class(self).GetVisibilityRangeFadeMode())
}

func (self Go) SetVisibilityRangeFadeMode(value classdb.GeometryInstance3DVisibilityRangeFadeMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibilityRangeFadeMode(value)
}

//go:nosplit
func (self class) SetMaterialOverride(material gdclass.Material)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_set_material_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaterialOverride(ctx gd.Lifetime) gdclass.Material {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_get_material_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Material
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaterialOverlay(material gdclass.Material)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_set_material_overlay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaterialOverlay(ctx gd.Lifetime) gdclass.Material {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_get_material_overlay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Material
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCastShadowsSetting(shadow_casting_setting classdb.GeometryInstance3DShadowCastingSetting)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shadow_casting_setting)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_set_cast_shadows_setting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCastShadowsSetting() classdb.GeometryInstance3DShadowCastingSetting {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GeometryInstance3DShadowCastingSetting](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_get_cast_shadows_setting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLodBias(bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_set_lod_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLodBias() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_get_lod_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTransparency(transparency gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, transparency)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_set_transparency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTransparency() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_get_transparency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisibilityRangeEndMargin(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_set_visibility_range_end_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVisibilityRangeEndMargin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_get_visibility_range_end_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisibilityRangeEnd(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_set_visibility_range_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVisibilityRangeEnd() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_get_visibility_range_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisibilityRangeBeginMargin(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_set_visibility_range_begin_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVisibilityRangeBeginMargin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_get_visibility_range_begin_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisibilityRangeBegin(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_set_visibility_range_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVisibilityRangeBegin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_get_visibility_range_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisibilityRangeFadeMode(mode classdb.GeometryInstance3DVisibilityRangeFadeMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_set_visibility_range_fade_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVisibilityRangeFadeMode() classdb.GeometryInstance3DVisibilityRangeFadeMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GeometryInstance3DVisibilityRangeFadeMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_get_visibility_range_fade_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set the value of a shader uniform for this instance only ([url=$DOCS_URL/tutorials/shaders/shader_reference/shading_language.html#per-instance-uniforms]per-instance uniform[/url]). See also [method ShaderMaterial.set_shader_parameter] to assign a uniform on all instances using the same [ShaderMaterial].
[b]Note:[/b] For a shader uniform to be assignable on a per-instance basis, it [i]must[/i] be defined with [code]instance uniform ...[/code] rather than [code]uniform ...[/code] in the shader code.
[b]Note:[/b] [param name] is case-sensitive and must match the name of the uniform in the code exactly (not the capitalized name in the inspector).
[b]Note:[/b] Per-instance shader uniforms are currently only available in 3D, so there is no 2D equivalent of this method.
*/
//go:nosplit
func (self class) SetInstanceShaderParameter(name gd.StringName, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_set_instance_shader_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Get the value of a shader parameter as set on this instance.
*/
//go:nosplit
func (self class) GetInstanceShaderParameter(ctx gd.Lifetime, name gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_get_instance_shader_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetExtraCullMargin(margin gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_set_extra_cull_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetExtraCullMargin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_get_extra_cull_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLightmapScale(scale classdb.GeometryInstance3DLightmapScale)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_set_lightmap_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLightmapScale() classdb.GeometryInstance3DLightmapScale {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GeometryInstance3DLightmapScale](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_get_lightmap_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGiMode(mode classdb.GeometryInstance3DGIMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_set_gi_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGiMode() classdb.GeometryInstance3DGIMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GeometryInstance3DGIMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_get_gi_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIgnoreOcclusionCulling(ignore_culling bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ignore_culling)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_set_ignore_occlusion_culling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsIgnoringOcclusionCulling() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_is_ignoring_occlusion_culling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCustomAabb(aabb gd.AABB)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, aabb)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_set_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCustomAabb() gd.AABB {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GeometryInstance3D.Bind_get_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGeometryInstance3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGeometryInstance3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualInstance3D() VisualInstance3D.GD { return *((*VisualInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualInstance3D() VisualInstance3D.Go { return *((*VisualInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualInstance3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualInstance3D(), name)
	}
}
func init() {classdb.Register("GeometryInstance3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type ShadowCastingSetting = classdb.GeometryInstance3DShadowCastingSetting

const (
/*Will not cast any shadows. Use this to improve performance for small geometry that is unlikely to cast noticeable shadows (such as debris).*/
	ShadowCastingSettingOff ShadowCastingSetting = 0
/*Will cast shadows from all visible faces in the GeometryInstance3D.
Will take culling into account, so faces not being rendered will not be taken into account when shadow casting.*/
	ShadowCastingSettingOn ShadowCastingSetting = 1
/*Will cast shadows from all visible faces in the GeometryInstance3D.
Will not take culling into account, so all faces will be taken into account when shadow casting.*/
	ShadowCastingSettingDoubleSided ShadowCastingSetting = 2
/*Will only show the shadows casted from this object.
In other words, the actual mesh will not be visible, only the shadows casted from the mesh will be.*/
	ShadowCastingSettingShadowsOnly ShadowCastingSetting = 3
)
type GIMode = classdb.GeometryInstance3DGIMode

const (
/*Disabled global illumination mode. Use for dynamic objects that do not contribute to global illumination (such as characters). When using [VoxelGI] and SDFGI, the geometry will [i]receive[/i] indirect lighting and reflections but the geometry will not be considered in GI baking.*/
	GiModeDisabled GIMode = 0
/*Baked global illumination mode. Use for static objects that contribute to global illumination (such as level geometry). This GI mode is effective when using [VoxelGI], SDFGI and [LightmapGI].*/
	GiModeStatic GIMode = 1
/*Dynamic global illumination mode. Use for dynamic objects that contribute to global illumination. This GI mode is only effective when using [VoxelGI], but it has a higher performance impact than [constant GI_MODE_STATIC]. When using other GI methods, this will act the same as [constant GI_MODE_DISABLED]. When using [LightmapGI], the object will receive indirect lighting using lightmap probes instead of using the baked lightmap texture.*/
	GiModeDynamic GIMode = 2
)
type LightmapScale = classdb.GeometryInstance3DLightmapScale

const (
/*The standard texel density for lightmapping with [LightmapGI].*/
	LightmapScale1x LightmapScale = 0
/*Multiplies texel density by 2× for lightmapping with [LightmapGI]. To ensure consistency in texel density, use this when scaling a mesh by a factor between 1.5 and 3.0.*/
	LightmapScale2x LightmapScale = 1
/*Multiplies texel density by 4× for lightmapping with [LightmapGI]. To ensure consistency in texel density, use this when scaling a mesh by a factor between 3.0 and 6.0.*/
	LightmapScale4x LightmapScale = 2
/*Multiplies texel density by 8× for lightmapping with [LightmapGI]. To ensure consistency in texel density, use this when scaling a mesh by a factor greater than 6.0.*/
	LightmapScale8x LightmapScale = 3
/*Represents the size of the [enum LightmapScale] enum.*/
	LightmapScaleMax LightmapScale = 4
)
type VisibilityRangeFadeMode = classdb.GeometryInstance3DVisibilityRangeFadeMode

const (
/*Will not fade itself nor its visibility dependencies, hysteresis will be used instead. This is the fastest approach to manual LOD, but it can result in noticeable LOD transitions depending on how the LOD meshes are authored. See [member visibility_range_begin] and [member Node3D.visibility_parent] for more information.*/
	VisibilityRangeFadeDisabled VisibilityRangeFadeMode = 0
/*Will fade-out itself when reaching the limits of its own visibility range. This is slower than [constant VISIBILITY_RANGE_FADE_DISABLED], but it can provide smoother transitions. The fading range is determined by [member visibility_range_begin_margin] and [member visibility_range_end_margin].
[b]Note:[/b] Only supported when using the Forward+ rendering method. When using the Mobile or Compatibility rendering method, this mode acts like [constant VISIBILITY_RANGE_FADE_DISABLED] but with hysteresis disabled.*/
	VisibilityRangeFadeSelf VisibilityRangeFadeMode = 1
/*Will fade-in its visibility dependencies (see [member Node3D.visibility_parent]) when reaching the limits of its own visibility range. This is slower than [constant VISIBILITY_RANGE_FADE_DISABLED], but it can provide smoother transitions. The fading range is determined by [member visibility_range_begin_margin] and [member visibility_range_end_margin].
[b]Note:[/b] Only supported when using the Forward+ rendering method. When using the Mobile or Compatibility rendering method, this mode acts like [constant VISIBILITY_RANGE_FADE_DISABLED] but with hysteresis disabled.*/
	VisibilityRangeFadeDependencies VisibilityRangeFadeMode = 2
)
