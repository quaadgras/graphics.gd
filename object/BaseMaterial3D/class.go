package BaseMaterial3D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
This class serves as a default material with a wide variety of rendering features and properties without the need to write shader code. See the tutorial below for details.

*/
type Simple [1]classdb.BaseMaterial3D
func (self Simple) SetAlbedo(albedo gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAlbedo(albedo)
}
func (self Simple) GetAlbedo() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetAlbedo())
}
func (self Simple) SetTransparency(transparency classdb.BaseMaterial3DTransparency) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTransparency(transparency)
}
func (self Simple) GetTransparency() classdb.BaseMaterial3DTransparency {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseMaterial3DTransparency(Expert(self).GetTransparency())
}
func (self Simple) SetAlphaAntialiasing(alpha_aa classdb.BaseMaterial3DAlphaAntiAliasing) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAlphaAntialiasing(alpha_aa)
}
func (self Simple) GetAlphaAntialiasing() classdb.BaseMaterial3DAlphaAntiAliasing {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseMaterial3DAlphaAntiAliasing(Expert(self).GetAlphaAntialiasing())
}
func (self Simple) SetAlphaAntialiasingEdge(edge float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAlphaAntialiasingEdge(gd.Float(edge))
}
func (self Simple) GetAlphaAntialiasingEdge() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAlphaAntialiasingEdge()))
}
func (self Simple) SetShadingMode(shading_mode classdb.BaseMaterial3DShadingMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShadingMode(shading_mode)
}
func (self Simple) GetShadingMode() classdb.BaseMaterial3DShadingMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseMaterial3DShadingMode(Expert(self).GetShadingMode())
}
func (self Simple) SetSpecular(specular float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSpecular(gd.Float(specular))
}
func (self Simple) GetSpecular() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSpecular()))
}
func (self Simple) SetMetallic(metallic float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMetallic(gd.Float(metallic))
}
func (self Simple) GetMetallic() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMetallic()))
}
func (self Simple) SetRoughness(roughness float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRoughness(gd.Float(roughness))
}
func (self Simple) GetRoughness() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetRoughness()))
}
func (self Simple) SetEmission(emission gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmission(emission)
}
func (self Simple) GetEmission() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetEmission())
}
func (self Simple) SetEmissionEnergyMultiplier(emission_energy_multiplier float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmissionEnergyMultiplier(gd.Float(emission_energy_multiplier))
}
func (self Simple) GetEmissionEnergyMultiplier() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetEmissionEnergyMultiplier()))
}
func (self Simple) SetEmissionIntensity(emission_energy_multiplier float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmissionIntensity(gd.Float(emission_energy_multiplier))
}
func (self Simple) GetEmissionIntensity() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetEmissionIntensity()))
}
func (self Simple) SetNormalScale(normal_scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNormalScale(gd.Float(normal_scale))
}
func (self Simple) GetNormalScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetNormalScale()))
}
func (self Simple) SetRim(rim float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRim(gd.Float(rim))
}
func (self Simple) GetRim() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetRim()))
}
func (self Simple) SetRimTint(rim_tint float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRimTint(gd.Float(rim_tint))
}
func (self Simple) GetRimTint() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetRimTint()))
}
func (self Simple) SetClearcoat(clearcoat float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetClearcoat(gd.Float(clearcoat))
}
func (self Simple) GetClearcoat() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetClearcoat()))
}
func (self Simple) SetClearcoatRoughness(clearcoat_roughness float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetClearcoatRoughness(gd.Float(clearcoat_roughness))
}
func (self Simple) GetClearcoatRoughness() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetClearcoatRoughness()))
}
func (self Simple) SetAnisotropy(anisotropy float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAnisotropy(gd.Float(anisotropy))
}
func (self Simple) GetAnisotropy() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAnisotropy()))
}
func (self Simple) SetHeightmapScale(heightmap_scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHeightmapScale(gd.Float(heightmap_scale))
}
func (self Simple) GetHeightmapScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetHeightmapScale()))
}
func (self Simple) SetSubsurfaceScatteringStrength(strength float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSubsurfaceScatteringStrength(gd.Float(strength))
}
func (self Simple) GetSubsurfaceScatteringStrength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSubsurfaceScatteringStrength()))
}
func (self Simple) SetTransmittanceColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTransmittanceColor(color)
}
func (self Simple) GetTransmittanceColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetTransmittanceColor())
}
func (self Simple) SetTransmittanceDepth(depth float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTransmittanceDepth(gd.Float(depth))
}
func (self Simple) GetTransmittanceDepth() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTransmittanceDepth()))
}
func (self Simple) SetTransmittanceBoost(boost float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTransmittanceBoost(gd.Float(boost))
}
func (self Simple) GetTransmittanceBoost() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTransmittanceBoost()))
}
func (self Simple) SetBacklight(backlight gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBacklight(backlight)
}
func (self Simple) GetBacklight() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetBacklight())
}
func (self Simple) SetRefraction(refraction float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRefraction(gd.Float(refraction))
}
func (self Simple) GetRefraction() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetRefraction()))
}
func (self Simple) SetPointSize(point_size float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPointSize(gd.Float(point_size))
}
func (self Simple) GetPointSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetPointSize()))
}
func (self Simple) SetDetailUv(detail_uv classdb.BaseMaterial3DDetailUV) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDetailUv(detail_uv)
}
func (self Simple) GetDetailUv() classdb.BaseMaterial3DDetailUV {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseMaterial3DDetailUV(Expert(self).GetDetailUv())
}
func (self Simple) SetBlendMode(blend_mode classdb.BaseMaterial3DBlendMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBlendMode(blend_mode)
}
func (self Simple) GetBlendMode() classdb.BaseMaterial3DBlendMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseMaterial3DBlendMode(Expert(self).GetBlendMode())
}
func (self Simple) SetDepthDrawMode(depth_draw_mode classdb.BaseMaterial3DDepthDrawMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDepthDrawMode(depth_draw_mode)
}
func (self Simple) GetDepthDrawMode() classdb.BaseMaterial3DDepthDrawMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseMaterial3DDepthDrawMode(Expert(self).GetDepthDrawMode())
}
func (self Simple) SetCullMode(cull_mode classdb.BaseMaterial3DCullMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCullMode(cull_mode)
}
func (self Simple) GetCullMode() classdb.BaseMaterial3DCullMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseMaterial3DCullMode(Expert(self).GetCullMode())
}
func (self Simple) SetDiffuseMode(diffuse_mode classdb.BaseMaterial3DDiffuseMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDiffuseMode(diffuse_mode)
}
func (self Simple) GetDiffuseMode() classdb.BaseMaterial3DDiffuseMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseMaterial3DDiffuseMode(Expert(self).GetDiffuseMode())
}
func (self Simple) SetSpecularMode(specular_mode classdb.BaseMaterial3DSpecularMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSpecularMode(specular_mode)
}
func (self Simple) GetSpecularMode() classdb.BaseMaterial3DSpecularMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseMaterial3DSpecularMode(Expert(self).GetSpecularMode())
}
func (self Simple) SetFlag(flag classdb.BaseMaterial3DFlags, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFlag(flag, enable)
}
func (self Simple) GetFlag(flag classdb.BaseMaterial3DFlags) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetFlag(flag))
}
func (self Simple) SetTextureFilter(mode classdb.BaseMaterial3DTextureFilter) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTextureFilter(mode)
}
func (self Simple) GetTextureFilter() classdb.BaseMaterial3DTextureFilter {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseMaterial3DTextureFilter(Expert(self).GetTextureFilter())
}
func (self Simple) SetFeature(feature classdb.BaseMaterial3DFeature, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFeature(feature, enable)
}
func (self Simple) GetFeature(feature classdb.BaseMaterial3DFeature) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetFeature(feature))
}
func (self Simple) SetTexture(param classdb.BaseMaterial3DTextureParam, texture [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTexture(param, texture)
}
func (self Simple) GetTexture(param classdb.BaseMaterial3DTextureParam) [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetTexture(gc, param))
}
func (self Simple) SetDetailBlendMode(detail_blend_mode classdb.BaseMaterial3DBlendMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDetailBlendMode(detail_blend_mode)
}
func (self Simple) GetDetailBlendMode() classdb.BaseMaterial3DBlendMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseMaterial3DBlendMode(Expert(self).GetDetailBlendMode())
}
func (self Simple) SetUv1Scale(scale gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUv1Scale(scale)
}
func (self Simple) GetUv1Scale() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetUv1Scale())
}
func (self Simple) SetUv1Offset(offset gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUv1Offset(offset)
}
func (self Simple) GetUv1Offset() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetUv1Offset())
}
func (self Simple) SetUv1TriplanarBlendSharpness(sharpness float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUv1TriplanarBlendSharpness(gd.Float(sharpness))
}
func (self Simple) GetUv1TriplanarBlendSharpness() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetUv1TriplanarBlendSharpness()))
}
func (self Simple) SetUv2Scale(scale gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUv2Scale(scale)
}
func (self Simple) GetUv2Scale() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetUv2Scale())
}
func (self Simple) SetUv2Offset(offset gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUv2Offset(offset)
}
func (self Simple) GetUv2Offset() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetUv2Offset())
}
func (self Simple) SetUv2TriplanarBlendSharpness(sharpness float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUv2TriplanarBlendSharpness(gd.Float(sharpness))
}
func (self Simple) GetUv2TriplanarBlendSharpness() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetUv2TriplanarBlendSharpness()))
}
func (self Simple) SetBillboardMode(mode classdb.BaseMaterial3DBillboardMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBillboardMode(mode)
}
func (self Simple) GetBillboardMode() classdb.BaseMaterial3DBillboardMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseMaterial3DBillboardMode(Expert(self).GetBillboardMode())
}
func (self Simple) SetParticlesAnimHFrames(frames int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParticlesAnimHFrames(gd.Int(frames))
}
func (self Simple) GetParticlesAnimHFrames() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetParticlesAnimHFrames()))
}
func (self Simple) SetParticlesAnimVFrames(frames int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParticlesAnimVFrames(gd.Int(frames))
}
func (self Simple) GetParticlesAnimVFrames() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetParticlesAnimVFrames()))
}
func (self Simple) SetParticlesAnimLoop(loop bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParticlesAnimLoop(loop)
}
func (self Simple) GetParticlesAnimLoop() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetParticlesAnimLoop())
}
func (self Simple) SetHeightmapDeepParallax(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHeightmapDeepParallax(enable)
}
func (self Simple) IsHeightmapDeepParallaxEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsHeightmapDeepParallaxEnabled())
}
func (self Simple) SetHeightmapDeepParallaxMinLayers(layer int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHeightmapDeepParallaxMinLayers(gd.Int(layer))
}
func (self Simple) GetHeightmapDeepParallaxMinLayers() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetHeightmapDeepParallaxMinLayers()))
}
func (self Simple) SetHeightmapDeepParallaxMaxLayers(layer int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHeightmapDeepParallaxMaxLayers(gd.Int(layer))
}
func (self Simple) GetHeightmapDeepParallaxMaxLayers() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetHeightmapDeepParallaxMaxLayers()))
}
func (self Simple) SetHeightmapDeepParallaxFlipTangent(flip bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHeightmapDeepParallaxFlipTangent(flip)
}
func (self Simple) GetHeightmapDeepParallaxFlipTangent() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetHeightmapDeepParallaxFlipTangent())
}
func (self Simple) SetHeightmapDeepParallaxFlipBinormal(flip bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHeightmapDeepParallaxFlipBinormal(flip)
}
func (self Simple) GetHeightmapDeepParallaxFlipBinormal() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetHeightmapDeepParallaxFlipBinormal())
}
func (self Simple) SetGrow(amount float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGrow(gd.Float(amount))
}
func (self Simple) GetGrow() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetGrow()))
}
func (self Simple) SetEmissionOperator(operator classdb.BaseMaterial3DEmissionOperator) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmissionOperator(operator)
}
func (self Simple) GetEmissionOperator() classdb.BaseMaterial3DEmissionOperator {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseMaterial3DEmissionOperator(Expert(self).GetEmissionOperator())
}
func (self Simple) SetAoLightAffect(amount float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAoLightAffect(gd.Float(amount))
}
func (self Simple) GetAoLightAffect() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAoLightAffect()))
}
func (self Simple) SetAlphaScissorThreshold(threshold float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAlphaScissorThreshold(gd.Float(threshold))
}
func (self Simple) GetAlphaScissorThreshold() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAlphaScissorThreshold()))
}
func (self Simple) SetAlphaHashScale(threshold float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAlphaHashScale(gd.Float(threshold))
}
func (self Simple) GetAlphaHashScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAlphaHashScale()))
}
func (self Simple) SetGrowEnabled(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGrowEnabled(enable)
}
func (self Simple) IsGrowEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsGrowEnabled())
}
func (self Simple) SetMetallicTextureChannel(channel classdb.BaseMaterial3DTextureChannel) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMetallicTextureChannel(channel)
}
func (self Simple) GetMetallicTextureChannel() classdb.BaseMaterial3DTextureChannel {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseMaterial3DTextureChannel(Expert(self).GetMetallicTextureChannel())
}
func (self Simple) SetRoughnessTextureChannel(channel classdb.BaseMaterial3DTextureChannel) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRoughnessTextureChannel(channel)
}
func (self Simple) GetRoughnessTextureChannel() classdb.BaseMaterial3DTextureChannel {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseMaterial3DTextureChannel(Expert(self).GetRoughnessTextureChannel())
}
func (self Simple) SetAoTextureChannel(channel classdb.BaseMaterial3DTextureChannel) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAoTextureChannel(channel)
}
func (self Simple) GetAoTextureChannel() classdb.BaseMaterial3DTextureChannel {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseMaterial3DTextureChannel(Expert(self).GetAoTextureChannel())
}
func (self Simple) SetRefractionTextureChannel(channel classdb.BaseMaterial3DTextureChannel) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRefractionTextureChannel(channel)
}
func (self Simple) GetRefractionTextureChannel() classdb.BaseMaterial3DTextureChannel {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseMaterial3DTextureChannel(Expert(self).GetRefractionTextureChannel())
}
func (self Simple) SetProximityFadeEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetProximityFadeEnabled(enabled)
}
func (self Simple) IsProximityFadeEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsProximityFadeEnabled())
}
func (self Simple) SetProximityFadeDistance(distance float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetProximityFadeDistance(gd.Float(distance))
}
func (self Simple) GetProximityFadeDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetProximityFadeDistance()))
}
func (self Simple) SetMsdfPixelRange(arange float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMsdfPixelRange(gd.Float(arange))
}
func (self Simple) GetMsdfPixelRange() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMsdfPixelRange()))
}
func (self Simple) SetMsdfOutlineSize(size float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMsdfOutlineSize(gd.Float(size))
}
func (self Simple) GetMsdfOutlineSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMsdfOutlineSize()))
}
func (self Simple) SetDistanceFade(mode classdb.BaseMaterial3DDistanceFadeMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDistanceFade(mode)
}
func (self Simple) GetDistanceFade() classdb.BaseMaterial3DDistanceFadeMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseMaterial3DDistanceFadeMode(Expert(self).GetDistanceFade())
}
func (self Simple) SetDistanceFadeMaxDistance(distance float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDistanceFadeMaxDistance(gd.Float(distance))
}
func (self Simple) GetDistanceFadeMaxDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDistanceFadeMaxDistance()))
}
func (self Simple) SetDistanceFadeMinDistance(distance float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDistanceFadeMinDistance(gd.Float(distance))
}
func (self Simple) GetDistanceFadeMinDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDistanceFadeMinDistance()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.BaseMaterial3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetAlbedo(albedo gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, albedo)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_albedo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlbedo() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_albedo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTransparency(transparency classdb.BaseMaterial3DTransparency)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, transparency)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_transparency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTransparency() classdb.BaseMaterial3DTransparency {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DTransparency](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_transparency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlphaAntialiasing(alpha_aa classdb.BaseMaterial3DAlphaAntiAliasing)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, alpha_aa)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_alpha_antialiasing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlphaAntialiasing() classdb.BaseMaterial3DAlphaAntiAliasing {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DAlphaAntiAliasing](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_alpha_antialiasing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlphaAntialiasingEdge(edge gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, edge)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_alpha_antialiasing_edge, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlphaAntialiasingEdge() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_alpha_antialiasing_edge, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShadingMode(shading_mode classdb.BaseMaterial3DShadingMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shading_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_shading_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShadingMode() classdb.BaseMaterial3DShadingMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DShadingMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_shading_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSpecular(specular gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, specular)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_specular, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSpecular() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_specular, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMetallic(metallic gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, metallic)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_metallic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMetallic() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_metallic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRoughness(roughness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, roughness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_roughness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRoughness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_roughness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmission(emission gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, emission)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_emission, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmission() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_emission, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionEnergyMultiplier(emission_energy_multiplier gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, emission_energy_multiplier)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_emission_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionEnergyMultiplier() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_emission_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionIntensity(emission_energy_multiplier gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, emission_energy_multiplier)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_emission_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionIntensity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_emission_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNormalScale(normal_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, normal_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_normal_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNormalScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_normal_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRim(rim gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rim)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_rim, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRim() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_rim, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRimTint(rim_tint gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rim_tint)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_rim_tint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRimTint() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_rim_tint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetClearcoat(clearcoat gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, clearcoat)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_clearcoat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetClearcoat() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_clearcoat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetClearcoatRoughness(clearcoat_roughness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, clearcoat_roughness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_clearcoat_roughness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetClearcoatRoughness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_clearcoat_roughness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAnisotropy(anisotropy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, anisotropy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_anisotropy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAnisotropy() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_anisotropy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHeightmapScale(heightmap_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, heightmap_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_heightmap_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHeightmapScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_heightmap_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSubsurfaceScatteringStrength(strength gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_subsurface_scattering_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSubsurfaceScatteringStrength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_subsurface_scattering_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTransmittanceColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_transmittance_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTransmittanceColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_transmittance_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTransmittanceDepth(depth gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, depth)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_transmittance_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTransmittanceDepth() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_transmittance_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTransmittanceBoost(boost gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, boost)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_transmittance_boost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTransmittanceBoost() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_transmittance_boost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBacklight(backlight gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, backlight)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_backlight, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBacklight() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_backlight, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRefraction(refraction gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, refraction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_refraction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRefraction() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_refraction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPointSize(point_size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_point_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPointSize() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_point_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDetailUv(detail_uv classdb.BaseMaterial3DDetailUV)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, detail_uv)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_detail_uv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDetailUv() classdb.BaseMaterial3DDetailUV {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DDetailUV](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_detail_uv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBlendMode(blend_mode classdb.BaseMaterial3DBlendMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, blend_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBlendMode() classdb.BaseMaterial3DBlendMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DBlendMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDepthDrawMode(depth_draw_mode classdb.BaseMaterial3DDepthDrawMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, depth_draw_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_depth_draw_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDepthDrawMode() classdb.BaseMaterial3DDepthDrawMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DDepthDrawMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_depth_draw_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCullMode(cull_mode classdb.BaseMaterial3DCullMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cull_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_cull_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCullMode() classdb.BaseMaterial3DCullMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DCullMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_cull_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDiffuseMode(diffuse_mode classdb.BaseMaterial3DDiffuseMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, diffuse_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_diffuse_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDiffuseMode() classdb.BaseMaterial3DDiffuseMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DDiffuseMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_diffuse_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSpecularMode(specular_mode classdb.BaseMaterial3DSpecularMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, specular_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_specular_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSpecularMode() classdb.BaseMaterial3DSpecularMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DSpecularMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_specular_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], enables the specified flag. Flags are optional behavior that can be turned on and off. Only one flag can be enabled at a time with this function, the flag enumerators cannot be bit-masked together to enable or disable multiple flags at once. Flags can also be enabled by setting the corresponding member to [code]true[/code]. See [enum Flags] enumerator for options.
*/
//go:nosplit
func (self class) SetFlag(flag classdb.BaseMaterial3DFlags, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code], if the specified flag is enabled. See [enum Flags] enumerator for options.
*/
//go:nosplit
func (self class) GetFlag(flag classdb.BaseMaterial3DFlags) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextureFilter(mode classdb.BaseMaterial3DTextureFilter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_texture_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextureFilter() classdb.BaseMaterial3DTextureFilter {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DTextureFilter](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_texture_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], enables the specified [enum Feature]. Many features that are available in [BaseMaterial3D]s need to be enabled before use. This way the cost for using the feature is only incurred when specified. Features can also be enabled by setting the corresponding member to [code]true[/code].
*/
//go:nosplit
func (self class) SetFeature(feature classdb.BaseMaterial3DFeature, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_feature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code], if the specified [enum Feature] is enabled.
*/
//go:nosplit
func (self class) GetFeature(feature classdb.BaseMaterial3DFeature) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_feature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the texture for the slot specified by [param param]. See [enum TextureParam] for available slots.
*/
//go:nosplit
func (self class) SetTexture(param classdb.BaseMaterial3DTextureParam, texture object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Texture2D] associated with the specified [enum TextureParam].
*/
//go:nosplit
func (self class) GetTexture(ctx gd.Lifetime, param classdb.BaseMaterial3DTextureParam) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDetailBlendMode(detail_blend_mode classdb.BaseMaterial3DBlendMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, detail_blend_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_detail_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDetailBlendMode() classdb.BaseMaterial3DBlendMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DBlendMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_detail_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUv1Scale(scale gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_uv1_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUv1Scale() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_uv1_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUv1Offset(offset gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_uv1_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUv1Offset() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_uv1_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUv1TriplanarBlendSharpness(sharpness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sharpness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_uv1_triplanar_blend_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUv1TriplanarBlendSharpness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_uv1_triplanar_blend_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUv2Scale(scale gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_uv2_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUv2Scale() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_uv2_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUv2Offset(offset gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_uv2_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUv2Offset() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_uv2_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUv2TriplanarBlendSharpness(sharpness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sharpness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_uv2_triplanar_blend_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUv2TriplanarBlendSharpness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_uv2_triplanar_blend_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBillboardMode(mode classdb.BaseMaterial3DBillboardMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_billboard_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBillboardMode() classdb.BaseMaterial3DBillboardMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DBillboardMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_billboard_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetParticlesAnimHFrames(frames gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, frames)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_particles_anim_h_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParticlesAnimHFrames() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_particles_anim_h_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetParticlesAnimVFrames(frames gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, frames)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_particles_anim_v_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParticlesAnimVFrames() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_particles_anim_v_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetParticlesAnimLoop(loop bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, loop)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_particles_anim_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParticlesAnimLoop() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_particles_anim_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHeightmapDeepParallax(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_heightmap_deep_parallax, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsHeightmapDeepParallaxEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_is_heightmap_deep_parallax_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHeightmapDeepParallaxMinLayers(layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_heightmap_deep_parallax_min_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHeightmapDeepParallaxMinLayers() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_heightmap_deep_parallax_min_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHeightmapDeepParallaxMaxLayers(layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_heightmap_deep_parallax_max_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHeightmapDeepParallaxMaxLayers() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_heightmap_deep_parallax_max_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHeightmapDeepParallaxFlipTangent(flip bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flip)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_heightmap_deep_parallax_flip_tangent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHeightmapDeepParallaxFlipTangent() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_heightmap_deep_parallax_flip_tangent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHeightmapDeepParallaxFlipBinormal(flip bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flip)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_heightmap_deep_parallax_flip_binormal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHeightmapDeepParallaxFlipBinormal() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_heightmap_deep_parallax_flip_binormal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGrow(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_grow, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGrow() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_grow, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionOperator(operator classdb.BaseMaterial3DEmissionOperator)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, operator)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_emission_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionOperator() classdb.BaseMaterial3DEmissionOperator {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DEmissionOperator](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_emission_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAoLightAffect(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_ao_light_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAoLightAffect() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_ao_light_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlphaScissorThreshold(threshold gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, threshold)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_alpha_scissor_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlphaScissorThreshold() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_alpha_scissor_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlphaHashScale(threshold gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, threshold)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_alpha_hash_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlphaHashScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_alpha_hash_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGrowEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_grow_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsGrowEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_is_grow_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMetallicTextureChannel(channel classdb.BaseMaterial3DTextureChannel)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, channel)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_metallic_texture_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMetallicTextureChannel() classdb.BaseMaterial3DTextureChannel {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DTextureChannel](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_metallic_texture_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRoughnessTextureChannel(channel classdb.BaseMaterial3DTextureChannel)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, channel)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_roughness_texture_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRoughnessTextureChannel() classdb.BaseMaterial3DTextureChannel {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DTextureChannel](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_roughness_texture_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAoTextureChannel(channel classdb.BaseMaterial3DTextureChannel)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, channel)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_ao_texture_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAoTextureChannel() classdb.BaseMaterial3DTextureChannel {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DTextureChannel](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_ao_texture_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRefractionTextureChannel(channel classdb.BaseMaterial3DTextureChannel)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, channel)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_refraction_texture_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRefractionTextureChannel() classdb.BaseMaterial3DTextureChannel {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DTextureChannel](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_refraction_texture_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetProximityFadeEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_proximity_fade_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsProximityFadeEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_is_proximity_fade_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetProximityFadeDistance(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_proximity_fade_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetProximityFadeDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_proximity_fade_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMsdfPixelRange(arange gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, arange)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMsdfPixelRange() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMsdfOutlineSize(size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_msdf_outline_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMsdfOutlineSize() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_msdf_outline_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDistanceFade(mode classdb.BaseMaterial3DDistanceFadeMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_distance_fade, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDistanceFade() classdb.BaseMaterial3DDistanceFadeMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DDistanceFadeMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_distance_fade, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDistanceFadeMaxDistance(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_distance_fade_max_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDistanceFadeMaxDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_distance_fade_max_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDistanceFadeMinDistance(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_set_distance_fade_min_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDistanceFadeMinDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseMaterial3D.Bind_get_distance_fade_min_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsBaseMaterial3D() Expert { return self[0].AsBaseMaterial3D() }


//go:nosplit
func (self Simple) AsBaseMaterial3D() Simple { return self[0].AsBaseMaterial3D() }


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
func init() {classdb.Register("BaseMaterial3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type TextureParam = classdb.BaseMaterial3DTextureParam

const (
/*Texture specifying per-pixel color.*/
	TextureAlbedo TextureParam = 0
/*Texture specifying per-pixel metallic value.*/
	TextureMetallic TextureParam = 1
/*Texture specifying per-pixel roughness value.*/
	TextureRoughness TextureParam = 2
/*Texture specifying per-pixel emission color.*/
	TextureEmission TextureParam = 3
/*Texture specifying per-pixel normal vector.*/
	TextureNormal TextureParam = 4
/*Texture specifying per-pixel rim value.*/
	TextureRim TextureParam = 5
/*Texture specifying per-pixel clearcoat value.*/
	TextureClearcoat TextureParam = 6
/*Texture specifying per-pixel flowmap direction for use with [member anisotropy].*/
	TextureFlowmap TextureParam = 7
/*Texture specifying per-pixel ambient occlusion value.*/
	TextureAmbientOcclusion TextureParam = 8
/*Texture specifying per-pixel height.*/
	TextureHeightmap TextureParam = 9
/*Texture specifying per-pixel subsurface scattering.*/
	TextureSubsurfaceScattering TextureParam = 10
/*Texture specifying per-pixel transmittance for subsurface scattering.*/
	TextureSubsurfaceTransmittance TextureParam = 11
/*Texture specifying per-pixel backlight color.*/
	TextureBacklight TextureParam = 12
/*Texture specifying per-pixel refraction strength.*/
	TextureRefraction TextureParam = 13
/*Texture specifying per-pixel detail mask blending value.*/
	TextureDetailMask TextureParam = 14
/*Texture specifying per-pixel detail color.*/
	TextureDetailAlbedo TextureParam = 15
/*Texture specifying per-pixel detail normal.*/
	TextureDetailNormal TextureParam = 16
/*Texture holding ambient occlusion, roughness, and metallic.*/
	TextureOrm TextureParam = 17
/*Represents the size of the [enum TextureParam] enum.*/
	TextureMax TextureParam = 18
)
type TextureFilter = classdb.BaseMaterial3DTextureFilter

const (
/*The texture filter reads from the nearest pixel only. This makes the texture look pixelated from up close, and grainy from a distance (due to mipmaps not being sampled).*/
	TextureFilterNearest TextureFilter = 0
/*The texture filter blends between the nearest 4 pixels. This makes the texture look smooth from up close, and grainy from a distance (due to mipmaps not being sampled).*/
	TextureFilterLinear TextureFilter = 1
/*The texture filter reads from the nearest pixel and blends between the nearest 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]). This makes the texture look pixelated from up close, and smooth from a distance.*/
	TextureFilterNearestWithMipmaps TextureFilter = 2
/*The texture filter blends between the nearest 4 pixels and between the nearest 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]). This makes the texture look smooth from up close, and smooth from a distance.*/
	TextureFilterLinearWithMipmaps TextureFilter = 3
/*The texture filter reads from the nearest pixel and blends between 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]) based on the angle between the surface and the camera view. This makes the texture look pixelated from up close, and smooth from a distance. Anisotropic filtering improves texture quality on surfaces that are almost in line with the camera, but is slightly slower. The anisotropic filtering level can be changed by adjusting [member ProjectSettings.rendering/textures/default_filters/anisotropic_filtering_level].*/
	TextureFilterNearestWithMipmapsAnisotropic TextureFilter = 4
/*The texture filter blends between the nearest 4 pixels and blends between 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]) based on the angle between the surface and the camera view. This makes the texture look smooth from up close, and smooth from a distance. Anisotropic filtering improves texture quality on surfaces that are almost in line with the camera, but is slightly slower. The anisotropic filtering level can be changed by adjusting [member ProjectSettings.rendering/textures/default_filters/anisotropic_filtering_level].*/
	TextureFilterLinearWithMipmapsAnisotropic TextureFilter = 5
/*Represents the size of the [enum TextureFilter] enum.*/
	TextureFilterMax TextureFilter = 6
)
type DetailUV = classdb.BaseMaterial3DDetailUV

const (
/*Use [code]UV[/code] with the detail texture.*/
	DetailUv1 DetailUV = 0
/*Use [code]UV2[/code] with the detail texture.*/
	DetailUv2 DetailUV = 1
)
type Transparency = classdb.BaseMaterial3DTransparency

const (
/*The material will not use transparency. This is the fastest to render.*/
	TransparencyDisabled Transparency = 0
/*The material will use the texture's alpha values for transparency. This is the slowest to render, and disables shadow casting.*/
	TransparencyAlpha Transparency = 1
/*The material will cut off all values below a threshold, the rest will remain opaque. The opaque portions will be rendered in the depth prepass. This is faster to render than alpha blending, but slower than opaque rendering. This also supports casting shadows.*/
	TransparencyAlphaScissor Transparency = 2
/*The material will cut off all values below a spatially-deterministic threshold, the rest will remain opaque. This is faster to render than alpha blending, but slower than opaque rendering. This also supports casting shadows. Alpha hashing is suited for hair rendering.*/
	TransparencyAlphaHash Transparency = 3
/*The material will use the texture's alpha value for transparency, but will discard fragments with an alpha of less than 0.99 during the depth prepass and fragments with an alpha less than 0.1 during the shadow pass. This also supports casting shadows.*/
	TransparencyAlphaDepthPrePass Transparency = 4
/*Represents the size of the [enum Transparency] enum.*/
	TransparencyMax Transparency = 5
)
type ShadingMode = classdb.BaseMaterial3DShadingMode

const (
/*The object will not receive shadows. This is the fastest to render, but it disables all interactions with lights.*/
	ShadingModeUnshaded ShadingMode = 0
/*The object will be shaded per pixel. Useful for realistic shading effects.*/
	ShadingModePerPixel ShadingMode = 1
/*The object will be shaded per vertex. Useful when you want cheaper shaders and do not care about visual quality. Not implemented yet (this mode will act like [constant SHADING_MODE_PER_PIXEL]).*/
	ShadingModePerVertex ShadingMode = 2
/*Represents the size of the [enum ShadingMode] enum.*/
	ShadingModeMax ShadingMode = 3
)
type Feature = classdb.BaseMaterial3DFeature

const (
/*Constant for setting [member emission_enabled].*/
	FeatureEmission Feature = 0
/*Constant for setting [member normal_enabled].*/
	FeatureNormalMapping Feature = 1
/*Constant for setting [member rim_enabled].*/
	FeatureRim Feature = 2
/*Constant for setting [member clearcoat_enabled].*/
	FeatureClearcoat Feature = 3
/*Constant for setting [member anisotropy_enabled].*/
	FeatureAnisotropy Feature = 4
/*Constant for setting [member ao_enabled].*/
	FeatureAmbientOcclusion Feature = 5
/*Constant for setting [member heightmap_enabled].*/
	FeatureHeightMapping Feature = 6
/*Constant for setting [member subsurf_scatter_enabled].*/
	FeatureSubsurfaceScattering Feature = 7
/*Constant for setting [member subsurf_scatter_transmittance_enabled].*/
	FeatureSubsurfaceTransmittance Feature = 8
/*Constant for setting [member backlight_enabled].*/
	FeatureBacklight Feature = 9
/*Constant for setting [member refraction_enabled].*/
	FeatureRefraction Feature = 10
/*Constant for setting [member detail_enabled].*/
	FeatureDetail Feature = 11
/*Represents the size of the [enum Feature] enum.*/
	FeatureMax Feature = 12
)
type BlendMode = classdb.BaseMaterial3DBlendMode

const (
/*Default blend mode. The color of the object is blended over the background based on the object's alpha value.*/
	BlendModeMix BlendMode = 0
/*The color of the object is added to the background.*/
	BlendModeAdd BlendMode = 1
/*The color of the object is subtracted from the background.*/
	BlendModeSub BlendMode = 2
/*The color of the object is multiplied by the background.*/
	BlendModeMul BlendMode = 3
/*The color of the object is added to the background and the alpha channel is used to mask out the background. This is effectively a hybrid of the blend mix and add modes, useful for effects like fire where you want the flame to add but the smoke to mix. By default, this works with unshaded materials using premultiplied textures. For shaded materials, use the [code]PREMUL_ALPHA_FACTOR[/code] built-in so that lighting can be modulated as well.*/
	BlendModePremultAlpha BlendMode = 4
)
type AlphaAntiAliasing = classdb.BaseMaterial3DAlphaAntiAliasing

const (
/*Disables Alpha AntiAliasing for the material.*/
	AlphaAntialiasingOff AlphaAntiAliasing = 0
/*Enables AlphaToCoverage. Alpha values in the material are passed to the AntiAliasing sample mask.*/
	AlphaAntialiasingAlphaToCoverage AlphaAntiAliasing = 1
/*Enables AlphaToCoverage and forces all non-zero alpha values to [code]1[/code]. Alpha values in the material are passed to the AntiAliasing sample mask.*/
	AlphaAntialiasingAlphaToCoverageAndToOne AlphaAntiAliasing = 2
)
type DepthDrawMode = classdb.BaseMaterial3DDepthDrawMode

const (
/*Default depth draw mode. Depth is drawn only for opaque objects during the opaque prepass (if any) and during the opaque pass.*/
	DepthDrawOpaqueOnly DepthDrawMode = 0
/*Objects will write to depth during the opaque and the transparent passes. Transparent objects that are close to the camera may obscure other transparent objects behind them.
[b]Note:[/b] This does not influence whether transparent objects are included in the depth prepass or not. For that, see [enum Transparency].*/
	DepthDrawAlways DepthDrawMode = 1
/*Objects will not write their depth to the depth buffer, even during the depth prepass (if enabled).*/
	DepthDrawDisabled DepthDrawMode = 2
)
type CullMode = classdb.BaseMaterial3DCullMode

const (
/*Default cull mode. The back of the object is culled when not visible. Back face triangles will be culled when facing the camera. This results in only the front side of triangles being drawn. For closed-surface meshes, this means that only the exterior of the mesh will be visible.*/
	CullBack CullMode = 0
/*Front face triangles will be culled when facing the camera. This results in only the back side of triangles being drawn. For closed-surface meshes, this means that the interior of the mesh will be drawn instead of the exterior.*/
	CullFront CullMode = 1
/*No face culling is performed; both the front face and back face will be visible.*/
	CullDisabled CullMode = 2
)
type Flags = classdb.BaseMaterial3DFlags

const (
/*Disables the depth test, so this object is drawn on top of all others drawn before it. This puts the object in the transparent draw pass where it is sorted based on distance to camera. Objects drawn after it in the draw order may cover it. This also disables writing to depth.*/
	FlagDisableDepthTest Flags = 0
/*Set [code]ALBEDO[/code] to the per-vertex color specified in the mesh.*/
	FlagAlbedoFromVertexColor Flags = 1
/*Vertex colors are considered to be stored in sRGB color space and are converted to linear color space during rendering. See also [member vertex_color_is_srgb].
[b]Note:[/b] Only effective when using the Forward+ and Mobile rendering methods.*/
	FlagSrgbVertexColor Flags = 2
/*Uses point size to alter the size of primitive points. Also changes the albedo texture lookup to use [code]POINT_COORD[/code] instead of [code]UV[/code].*/
	FlagUsePointSize Flags = 3
/*Object is scaled by depth so that it always appears the same size on screen.*/
	FlagFixedSize Flags = 4
/*Shader will keep the scale set for the mesh. Otherwise the scale is lost when billboarding. Only applies when [member billboard_mode] is [constant BILLBOARD_ENABLED].*/
	FlagBillboardKeepScale Flags = 5
/*Use triplanar texture lookup for all texture lookups that would normally use [code]UV[/code].*/
	FlagUv1UseTriplanar Flags = 6
/*Use triplanar texture lookup for all texture lookups that would normally use [code]UV2[/code].*/
	FlagUv2UseTriplanar Flags = 7
/*Use triplanar texture lookup for all texture lookups that would normally use [code]UV[/code].*/
	FlagUv1UseWorldTriplanar Flags = 8
/*Use triplanar texture lookup for all texture lookups that would normally use [code]UV2[/code].*/
	FlagUv2UseWorldTriplanar Flags = 9
/*Use [code]UV2[/code] coordinates to look up from the [member ao_texture].*/
	FlagAoOnUv2 Flags = 10
/*Use [code]UV2[/code] coordinates to look up from the [member emission_texture].*/
	FlagEmissionOnUv2 Flags = 11
/*Forces the shader to convert albedo from sRGB space to linear space. See also [member albedo_texture_force_srgb].*/
	FlagAlbedoTextureForceSrgb Flags = 12
/*Disables receiving shadows from other objects.*/
	FlagDontReceiveShadows Flags = 13
/*Disables receiving ambient light.*/
	FlagDisableAmbientLight Flags = 14
/*Enables the shadow to opacity feature.*/
	FlagUseShadowToOpacity Flags = 15
/*Enables the texture to repeat when UV coordinates are outside the 0-1 range. If using one of the linear filtering modes, this can result in artifacts at the edges of a texture when the sampler filters across the edges of the texture.*/
	FlagUseTextureRepeat Flags = 16
/*Invert values read from a depth texture to convert them to height values (heightmap).*/
	FlagInvertHeightmap Flags = 17
/*Enables the skin mode for subsurface scattering which is used to improve the look of subsurface scattering when used for human skin.*/
	FlagSubsurfaceModeSkin Flags = 18
/*Enables parts of the shader required for [GPUParticles3D] trails to function. This also requires using a mesh with appropriate skinning, such as [RibbonTrailMesh] or [TubeTrailMesh]. Enabling this feature outside of materials used in [GPUParticles3D] meshes will break material rendering.*/
	FlagParticleTrailsMode Flags = 19
/*Enables multichannel signed distance field rendering shader.*/
	FlagAlbedoTextureMsdf Flags = 20
/*Disables receiving depth-based or volumetric fog.*/
	FlagDisableFog Flags = 21
/*Represents the size of the [enum Flags] enum.*/
	FlagMax Flags = 22
)
type DiffuseMode = classdb.BaseMaterial3DDiffuseMode

const (
/*Default diffuse scattering algorithm.*/
	DiffuseBurley DiffuseMode = 0
/*Diffuse scattering ignores roughness.*/
	DiffuseLambert DiffuseMode = 1
/*Extends Lambert to cover more than 90 degrees when roughness increases.*/
	DiffuseLambertWrap DiffuseMode = 2
/*Uses a hard cut for lighting, with smoothing affected by roughness.*/
	DiffuseToon DiffuseMode = 3
)
type SpecularMode = classdb.BaseMaterial3DSpecularMode

const (
/*Default specular blob.*/
	SpecularSchlickGgx SpecularMode = 0
/*Toon blob which changes size based on roughness.*/
	SpecularToon SpecularMode = 1
/*No specular blob. This is slightly faster to render than other specular modes.*/
	SpecularDisabled SpecularMode = 2
)
type BillboardMode = classdb.BaseMaterial3DBillboardMode

const (
/*Billboard mode is disabled.*/
	BillboardDisabled BillboardMode = 0
/*The object's Z axis will always face the camera.*/
	BillboardEnabled BillboardMode = 1
/*The object's X axis will always face the camera.*/
	BillboardFixedY BillboardMode = 2
/*Used for particle systems when assigned to [GPUParticles3D] and [CPUParticles3D] nodes (flipbook animation). Enables [code]particles_anim_*[/code] properties.
The [member ParticleProcessMaterial.anim_speed_min] or [member CPUParticles3D.anim_speed_min] should also be set to a value bigger than zero for the animation to play.*/
	BillboardParticles BillboardMode = 3
)
type TextureChannel = classdb.BaseMaterial3DTextureChannel

const (
/*Used to read from the red channel of a texture.*/
	TextureChannelRed TextureChannel = 0
/*Used to read from the green channel of a texture.*/
	TextureChannelGreen TextureChannel = 1
/*Used to read from the blue channel of a texture.*/
	TextureChannelBlue TextureChannel = 2
/*Used to read from the alpha channel of a texture.*/
	TextureChannelAlpha TextureChannel = 3
/*Used to read from the linear (non-perceptual) average of the red, green and blue channels of a texture.*/
	TextureChannelGrayscale TextureChannel = 4
)
type EmissionOperator = classdb.BaseMaterial3DEmissionOperator

const (
/*Adds the emission color to the color from the emission texture.*/
	EmissionOpAdd EmissionOperator = 0
/*Multiplies the emission color by the color from the emission texture.*/
	EmissionOpMultiply EmissionOperator = 1
)
type DistanceFadeMode = classdb.BaseMaterial3DDistanceFadeMode

const (
/*Do not use distance fade.*/
	DistanceFadeDisabled DistanceFadeMode = 0
/*Smoothly fades the object out based on each pixel's distance from the camera using the alpha channel.*/
	DistanceFadePixelAlpha DistanceFadeMode = 1
/*Smoothly fades the object out based on each pixel's distance from the camera using a dithering approach. Dithering discards pixels based on a set pattern to smoothly fade without enabling transparency. On certain hardware, this can be faster than [constant DISTANCE_FADE_PIXEL_ALPHA].*/
	DistanceFadePixelDither DistanceFadeMode = 2
/*Smoothly fades the object out based on the object's distance from the camera using a dithering approach. Dithering discards pixels based on a set pattern to smoothly fade without enabling transparency. On certain hardware, this can be faster than [constant DISTANCE_FADE_PIXEL_ALPHA] and [constant DISTANCE_FADE_PIXEL_DITHER].*/
	DistanceFadeObjectDither DistanceFadeMode = 3
)
