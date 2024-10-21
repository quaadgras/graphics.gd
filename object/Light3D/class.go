package Light3D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VisualInstance3D"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Light3D is the [i]abstract[/i] base class for light nodes. As it can't be instantiated, it shouldn't be used directly. Other types of light nodes inherit from it. Light3D contains the common variables and parameters used for lighting.

*/
type Simple [1]classdb.Light3D
func (self Simple) SetEditorOnly(editor_only bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEditorOnly(editor_only)
}
func (self Simple) IsEditorOnly() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsEditorOnly())
}
func (self Simple) SetParam(param classdb.Light3DParam, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParam(param, gd.Float(value))
}
func (self Simple) GetParam(param classdb.Light3DParam) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetParam(param)))
}
func (self Simple) SetShadow(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShadow(enabled)
}
func (self Simple) HasShadow() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasShadow())
}
func (self Simple) SetNegative(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNegative(enabled)
}
func (self Simple) IsNegative() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsNegative())
}
func (self Simple) SetCullMask(cull_mask int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCullMask(gd.Int(cull_mask))
}
func (self Simple) GetCullMask() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCullMask()))
}
func (self Simple) SetEnableDistanceFade(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnableDistanceFade(enable)
}
func (self Simple) IsDistanceFadeEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDistanceFadeEnabled())
}
func (self Simple) SetDistanceFadeBegin(distance float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDistanceFadeBegin(gd.Float(distance))
}
func (self Simple) GetDistanceFadeBegin() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDistanceFadeBegin()))
}
func (self Simple) SetDistanceFadeShadow(distance float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDistanceFadeShadow(gd.Float(distance))
}
func (self Simple) GetDistanceFadeShadow() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDistanceFadeShadow()))
}
func (self Simple) SetDistanceFadeLength(distance float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDistanceFadeLength(gd.Float(distance))
}
func (self Simple) GetDistanceFadeLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDistanceFadeLength()))
}
func (self Simple) SetColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColor(color)
}
func (self Simple) GetColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetColor())
}
func (self Simple) SetShadowReverseCullFace(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShadowReverseCullFace(enable)
}
func (self Simple) GetShadowReverseCullFace() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetShadowReverseCullFace())
}
func (self Simple) SetBakeMode(bake_mode classdb.Light3DBakeMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBakeMode(bake_mode)
}
func (self Simple) GetBakeMode() classdb.Light3DBakeMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.Light3DBakeMode(Expert(self).GetBakeMode())
}
func (self Simple) SetProjector(projector [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetProjector(projector)
}
func (self Simple) GetProjector() [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetProjector(gc))
}
func (self Simple) SetTemperature(temperature float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTemperature(gd.Float(temperature))
}
func (self Simple) GetTemperature() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTemperature()))
}
func (self Simple) GetCorrelatedColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetCorrelatedColor())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Light3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetEditorOnly(editor_only bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, editor_only)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_editor_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEditorOnly() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_is_editor_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the value of the specified [enum Light3D.Param] parameter.
*/
//go:nosplit
func (self class) SetParam(param classdb.Light3DParam, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the specified [enum Light3D.Param] parameter.
*/
//go:nosplit
func (self class) GetParam(param classdb.Light3DParam) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShadow(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_shadow, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) HasShadow() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_has_shadow, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNegative(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_negative, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsNegative() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_is_negative, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCullMask(cull_mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cull_mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCullMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnableDistanceFade(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_enable_distance_fade, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDistanceFadeEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_is_distance_fade_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDistanceFadeBegin(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_distance_fade_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDistanceFadeBegin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_distance_fade_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDistanceFadeShadow(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_distance_fade_shadow, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDistanceFadeShadow() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_distance_fade_shadow, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDistanceFadeLength(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_distance_fade_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDistanceFadeLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_distance_fade_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShadowReverseCullFace(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_shadow_reverse_cull_face, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShadowReverseCullFace() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_shadow_reverse_cull_face, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBakeMode(bake_mode classdb.Light3DBakeMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bake_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_bake_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBakeMode() classdb.Light3DBakeMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Light3DBakeMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_bake_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetProjector(projector object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(projector[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_projector, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetProjector(ctx gd.Lifetime) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_projector, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTemperature(temperature gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, temperature)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_temperature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTemperature() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_temperature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Color] of an idealized blackbody at the given [member light_temperature]. This value is calculated internally based on the [member light_temperature]. This [Color] is multiplied by [member light_color] before being sent to the [RenderingServer].
*/
//go:nosplit
func (self class) GetCorrelatedColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_correlated_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsLight3D() Expert { return self[0].AsLight3D() }


//go:nosplit
func (self Simple) AsLight3D() Simple { return self[0].AsLight3D() }


//go:nosplit
func (self class) AsVisualInstance3D() VisualInstance3D.Expert { return self[0].AsVisualInstance3D() }


//go:nosplit
func (self Simple) AsVisualInstance3D() VisualInstance3D.Simple { return self[0].AsVisualInstance3D() }


//go:nosplit
func (self class) AsNode3D() Node3D.Expert { return self[0].AsNode3D() }


//go:nosplit
func (self Simple) AsNode3D() Node3D.Simple { return self[0].AsNode3D() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("Light3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Param = classdb.Light3DParam

const (
/*Constant for accessing [member light_energy].*/
	ParamEnergy Param = 0
/*Constant for accessing [member light_indirect_energy].*/
	ParamIndirectEnergy Param = 1
/*Constant for accessing [member light_volumetric_fog_energy].*/
	ParamVolumetricFogEnergy Param = 2
/*Constant for accessing [member light_specular].*/
	ParamSpecular Param = 3
/*Constant for accessing [member OmniLight3D.omni_range] or [member SpotLight3D.spot_range].*/
	ParamRange Param = 4
/*Constant for accessing [member light_size].*/
	ParamSize Param = 5
/*Constant for accessing [member OmniLight3D.omni_attenuation] or [member SpotLight3D.spot_attenuation].*/
	ParamAttenuation Param = 6
/*Constant for accessing [member SpotLight3D.spot_angle].*/
	ParamSpotAngle Param = 7
/*Constant for accessing [member SpotLight3D.spot_angle_attenuation].*/
	ParamSpotAttenuation Param = 8
/*Constant for accessing [member DirectionalLight3D.directional_shadow_max_distance].*/
	ParamShadowMaxDistance Param = 9
/*Constant for accessing [member DirectionalLight3D.directional_shadow_split_1].*/
	ParamShadowSplit1Offset Param = 10
/*Constant for accessing [member DirectionalLight3D.directional_shadow_split_2].*/
	ParamShadowSplit2Offset Param = 11
/*Constant for accessing [member DirectionalLight3D.directional_shadow_split_3].*/
	ParamShadowSplit3Offset Param = 12
/*Constant for accessing [member DirectionalLight3D.directional_shadow_fade_start].*/
	ParamShadowFadeStart Param = 13
/*Constant for accessing [member shadow_normal_bias].*/
	ParamShadowNormalBias Param = 14
/*Constant for accessing [member shadow_bias].*/
	ParamShadowBias Param = 15
/*Constant for accessing [member DirectionalLight3D.directional_shadow_pancake_size].*/
	ParamShadowPancakeSize Param = 16
/*Constant for accessing [member shadow_opacity].*/
	ParamShadowOpacity Param = 17
/*Constant for accessing [member shadow_blur].*/
	ParamShadowBlur Param = 18
/*Constant for accessing [member shadow_transmittance_bias].*/
	ParamTransmittanceBias Param = 19
/*Constant for accessing [member light_intensity_lumens] and [member light_intensity_lux]. Only used when [member ProjectSettings.rendering/lights_and_shadows/use_physical_light_units] is [code]true[/code].*/
	ParamIntensity Param = 20
/*Represents the size of the [enum Param] enum.*/
	ParamMax Param = 21
)
type BakeMode = classdb.Light3DBakeMode

const (
/*Light is ignored when baking. This is the fastest mode, but the light will be taken into account when baking global illumination. This mode should generally be used for dynamic lights that change quickly, as the effect of global illumination is less noticeable on those lights.
[b]Note:[/b] Hiding a light does [i]not[/i] affect baking [LightmapGI]. Hiding a light will still affect baking [VoxelGI] and SDFGI (see [member Environment.sdfgi_enabled]).*/
	BakeDisabled BakeMode = 0
/*Light is taken into account in static baking ([VoxelGI], [LightmapGI], SDFGI ([member Environment.sdfgi_enabled])). The light can be moved around or modified, but its global illumination will not update in real-time. This is suitable for subtle changes (such as flickering torches), but generally not large changes such as toggling a light on and off.
[b]Note:[/b] The light is not baked in [LightmapGI] if [member editor_only] is [code]true[/code].*/
	BakeStatic BakeMode = 1
/*Light is taken into account in dynamic baking ([VoxelGI] and SDFGI ([member Environment.sdfgi_enabled]) only). The light can be moved around or modified with global illumination updating in real-time. The light's global illumination appearance will be slightly different compared to [constant BAKE_STATIC]. This has a greater performance cost compared to [constant BAKE_STATIC]. When using SDFGI, the update speed of dynamic lights is affected by [member ProjectSettings.rendering/global_illumination/sdfgi/frames_to_update_lights].*/
	BakeDynamic BakeMode = 2
)
