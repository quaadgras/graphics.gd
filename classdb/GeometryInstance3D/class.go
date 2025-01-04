// Package GeometryInstance3D provides methods for working with GeometryInstance3D object instances.
package GeometryInstance3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/VisualInstance3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/AABB"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Base node for geometry-based visual instances. Shares some common functionality like visibility and custom materials.
*/
type Instance [1]gdclass.GeometryInstance3D
type Any interface {
	gd.IsClass
	AsGeometryInstance3D() Instance
}

/*
Set the value of a shader uniform for this instance only ([url=$DOCS_URL/tutorials/shaders/shader_reference/shading_language.html#per-instance-uniforms]per-instance uniform[/url]). See also [method ShaderMaterial.set_shader_parameter] to assign a uniform on all instances using the same [ShaderMaterial].
[b]Note:[/b] For a shader uniform to be assignable on a per-instance basis, it [i]must[/i] be defined with [code]instance uniform ...[/code] rather than [code]uniform ...[/code] in the shader code.
[b]Note:[/b] [param name] is case-sensitive and must match the name of the uniform in the code exactly (not the capitalized name in the inspector).
[b]Note:[/b] Per-instance shader uniforms are currently only available in 3D, so there is no 2D equivalent of this method.
*/
func (self Instance) SetInstanceShaderParameter(name string, value any) {
	class(self).SetInstanceShaderParameter(gd.NewStringName(name), gd.NewVariant(value))
}

/*
Get the value of a shader parameter as set on this instance.
*/
func (self Instance) GetInstanceShaderParameter(name string) any {
	return any(class(self).GetInstanceShaderParameter(gd.NewStringName(name)).Interface())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GeometryInstance3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GeometryInstance3D"))
	return Instance{*(*gdclass.GeometryInstance3D)(unsafe.Pointer(&object))}
}

func (self Instance) MaterialOverride() [1]gdclass.Material {
	return [1]gdclass.Material(class(self).GetMaterialOverride())
}

func (self Instance) SetMaterialOverride(value [1]gdclass.Material) {
	class(self).SetMaterialOverride(value)
}

func (self Instance) MaterialOverlay() [1]gdclass.Material {
	return [1]gdclass.Material(class(self).GetMaterialOverlay())
}

func (self Instance) SetMaterialOverlay(value [1]gdclass.Material) {
	class(self).SetMaterialOverlay(value)
}

func (self Instance) Transparency() Float.X {
	return Float.X(Float.X(class(self).GetTransparency()))
}

func (self Instance) SetTransparency(value Float.X) {
	class(self).SetTransparency(gd.Float(value))
}

func (self Instance) CastShadow() gdclass.GeometryInstance3DShadowCastingSetting {
	return gdclass.GeometryInstance3DShadowCastingSetting(class(self).GetCastShadowsSetting())
}

func (self Instance) SetCastShadow(value gdclass.GeometryInstance3DShadowCastingSetting) {
	class(self).SetCastShadowsSetting(value)
}

func (self Instance) ExtraCullMargin() Float.X {
	return Float.X(Float.X(class(self).GetExtraCullMargin()))
}

func (self Instance) SetExtraCullMargin(value Float.X) {
	class(self).SetExtraCullMargin(gd.Float(value))
}

func (self Instance) CustomAabb() AABB.PositionSize {
	return AABB.PositionSize(class(self).GetCustomAabb())
}

func (self Instance) SetCustomAabb(value AABB.PositionSize) {
	class(self).SetCustomAabb(gd.AABB(value))
}

func (self Instance) LodBias() Float.X {
	return Float.X(Float.X(class(self).GetLodBias()))
}

func (self Instance) SetLodBias(value Float.X) {
	class(self).SetLodBias(gd.Float(value))
}

func (self Instance) IgnoreOcclusionCulling() bool {
	return bool(class(self).IsIgnoringOcclusionCulling())
}

func (self Instance) SetIgnoreOcclusionCulling(value bool) {
	class(self).SetIgnoreOcclusionCulling(value)
}

func (self Instance) GiMode() gdclass.GeometryInstance3DGIMode {
	return gdclass.GeometryInstance3DGIMode(class(self).GetGiMode())
}

func (self Instance) SetGiMode(value gdclass.GeometryInstance3DGIMode) {
	class(self).SetGiMode(value)
}

func (self Instance) GiLightmapScale() gdclass.GeometryInstance3DLightmapScale {
	return gdclass.GeometryInstance3DLightmapScale(class(self).GetLightmapScale())
}

func (self Instance) SetGiLightmapScale(value gdclass.GeometryInstance3DLightmapScale) {
	class(self).SetLightmapScale(value)
}

func (self Instance) VisibilityRangeBegin() Float.X {
	return Float.X(Float.X(class(self).GetVisibilityRangeBegin()))
}

func (self Instance) SetVisibilityRangeBegin(value Float.X) {
	class(self).SetVisibilityRangeBegin(gd.Float(value))
}

func (self Instance) VisibilityRangeBeginMargin() Float.X {
	return Float.X(Float.X(class(self).GetVisibilityRangeBeginMargin()))
}

func (self Instance) SetVisibilityRangeBeginMargin(value Float.X) {
	class(self).SetVisibilityRangeBeginMargin(gd.Float(value))
}

func (self Instance) VisibilityRangeEnd() Float.X {
	return Float.X(Float.X(class(self).GetVisibilityRangeEnd()))
}

func (self Instance) SetVisibilityRangeEnd(value Float.X) {
	class(self).SetVisibilityRangeEnd(gd.Float(value))
}

func (self Instance) VisibilityRangeEndMargin() Float.X {
	return Float.X(Float.X(class(self).GetVisibilityRangeEndMargin()))
}

func (self Instance) SetVisibilityRangeEndMargin(value Float.X) {
	class(self).SetVisibilityRangeEndMargin(gd.Float(value))
}

func (self Instance) VisibilityRangeFadeMode() gdclass.GeometryInstance3DVisibilityRangeFadeMode {
	return gdclass.GeometryInstance3DVisibilityRangeFadeMode(class(self).GetVisibilityRangeFadeMode())
}

func (self Instance) SetVisibilityRangeFadeMode(value gdclass.GeometryInstance3DVisibilityRangeFadeMode) {
	class(self).SetVisibilityRangeFadeMode(value)
}

//go:nosplit
func (self class) SetMaterialOverride(material [1]gdclass.Material) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_set_material_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaterialOverride() [1]gdclass.Material {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_get_material_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Material{gd.PointerWithOwnershipTransferredToGo[gdclass.Material](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaterialOverlay(material [1]gdclass.Material) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_set_material_overlay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaterialOverlay() [1]gdclass.Material {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_get_material_overlay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Material{gd.PointerWithOwnershipTransferredToGo[gdclass.Material](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCastShadowsSetting(shadow_casting_setting gdclass.GeometryInstance3DShadowCastingSetting) {
	var frame = callframe.New()
	callframe.Arg(frame, shadow_casting_setting)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_set_cast_shadows_setting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCastShadowsSetting() gdclass.GeometryInstance3DShadowCastingSetting {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GeometryInstance3DShadowCastingSetting](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_get_cast_shadows_setting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLodBias(bias gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_set_lod_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLodBias() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_get_lod_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTransparency(transparency gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, transparency)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_set_transparency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTransparency() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_get_transparency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibilityRangeEndMargin(distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_set_visibility_range_end_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityRangeEndMargin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_get_visibility_range_end_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibilityRangeEnd(distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_set_visibility_range_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityRangeEnd() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_get_visibility_range_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibilityRangeBeginMargin(distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_set_visibility_range_begin_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityRangeBeginMargin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_get_visibility_range_begin_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibilityRangeBegin(distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_set_visibility_range_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityRangeBegin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_get_visibility_range_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibilityRangeFadeMode(mode gdclass.GeometryInstance3DVisibilityRangeFadeMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_set_visibility_range_fade_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityRangeFadeMode() gdclass.GeometryInstance3DVisibilityRangeFadeMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GeometryInstance3DVisibilityRangeFadeMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_get_visibility_range_fade_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) SetInstanceShaderParameter(name gd.StringName, value gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_set_instance_shader_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Get the value of a shader parameter as set on this instance.
*/
//go:nosplit
func (self class) GetInstanceShaderParameter(name gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_get_instance_shader_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExtraCullMargin(margin gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_set_extra_cull_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetExtraCullMargin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_get_extra_cull_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLightmapScale(scale gdclass.GeometryInstance3DLightmapScale) {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_set_lightmap_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLightmapScale() gdclass.GeometryInstance3DLightmapScale {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GeometryInstance3DLightmapScale](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_get_lightmap_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGiMode(mode gdclass.GeometryInstance3DGIMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_set_gi_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGiMode() gdclass.GeometryInstance3DGIMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GeometryInstance3DGIMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_get_gi_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIgnoreOcclusionCulling(ignore_culling bool) {
	var frame = callframe.New()
	callframe.Arg(frame, ignore_culling)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_set_ignore_occlusion_culling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsIgnoringOcclusionCulling() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_is_ignoring_occlusion_culling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCustomAabb(aabb gd.AABB) {
	var frame = callframe.New()
	callframe.Arg(frame, aabb)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_set_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCustomAabb() gd.AABB {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GeometryInstance3D.Bind_get_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGeometryInstance3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGeometryInstance3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsVisualInstance3D() VisualInstance3D.Advanced {
	return *((*VisualInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualInstance3D() VisualInstance3D.Instance {
	return *((*VisualInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualInstance3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualInstance3D(), name)
	}
}
func init() {
	gdclass.Register("GeometryInstance3D", func(ptr gd.Object) any {
		return [1]gdclass.GeometryInstance3D{*(*gdclass.GeometryInstance3D)(unsafe.Pointer(&ptr))}
	})
}

type ShadowCastingSetting = gdclass.GeometryInstance3DShadowCastingSetting

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

type GIMode = gdclass.GeometryInstance3DGIMode

const (
	/*Disabled global illumination mode. Use for dynamic objects that do not contribute to global illumination (such as characters). When using [VoxelGI] and SDFGI, the geometry will [i]receive[/i] indirect lighting and reflections but the geometry will not be considered in GI baking.*/
	GiModeDisabled GIMode = 0
	/*Baked global illumination mode. Use for static objects that contribute to global illumination (such as level geometry). This GI mode is effective when using [VoxelGI], SDFGI and [LightmapGI].*/
	GiModeStatic GIMode = 1
	/*Dynamic global illumination mode. Use for dynamic objects that contribute to global illumination. This GI mode is only effective when using [VoxelGI], but it has a higher performance impact than [constant GI_MODE_STATIC]. When using other GI methods, this will act the same as [constant GI_MODE_DISABLED]. When using [LightmapGI], the object will receive indirect lighting using lightmap probes instead of using the baked lightmap texture.*/
	GiModeDynamic GIMode = 2
)

type LightmapScale = gdclass.GeometryInstance3DLightmapScale

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

type VisibilityRangeFadeMode = gdclass.GeometryInstance3DVisibilityRangeFadeMode

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
