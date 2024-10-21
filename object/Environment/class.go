package Environment

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
Resource for environment nodes (like [WorldEnvironment]) that define multiple environment operations (such as background [Sky] or [Color], ambient light, fog, depth-of-field...). These parameters affect the final render of the scene. The order of these operations is:
- Depth of Field Blur
- Glow
- Tonemap (Auto Exposure)
- Adjustments

*/
type Simple [1]classdb.Environment
func (self Simple) SetBackground(mode classdb.EnvironmentBGMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBackground(mode)
}
func (self Simple) GetBackground() classdb.EnvironmentBGMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.EnvironmentBGMode(Expert(self).GetBackground())
}
func (self Simple) SetSky(sky [1]classdb.Sky) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSky(sky)
}
func (self Simple) GetSky() [1]classdb.Sky {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Sky(Expert(self).GetSky(gc))
}
func (self Simple) SetSkyCustomFov(scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSkyCustomFov(gd.Float(scale))
}
func (self Simple) GetSkyCustomFov() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSkyCustomFov()))
}
func (self Simple) SetSkyRotation(euler_radians gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSkyRotation(euler_radians)
}
func (self Simple) GetSkyRotation() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetSkyRotation())
}
func (self Simple) SetBgColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBgColor(color)
}
func (self Simple) GetBgColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetBgColor())
}
func (self Simple) SetBgEnergyMultiplier(energy float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBgEnergyMultiplier(gd.Float(energy))
}
func (self Simple) GetBgEnergyMultiplier() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetBgEnergyMultiplier()))
}
func (self Simple) SetBgIntensity(energy float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBgIntensity(gd.Float(energy))
}
func (self Simple) GetBgIntensity() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetBgIntensity()))
}
func (self Simple) SetCanvasMaxLayer(layer int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCanvasMaxLayer(gd.Int(layer))
}
func (self Simple) GetCanvasMaxLayer() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCanvasMaxLayer()))
}
func (self Simple) SetCameraFeedId(id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCameraFeedId(gd.Int(id))
}
func (self Simple) GetCameraFeedId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCameraFeedId()))
}
func (self Simple) SetAmbientLightColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAmbientLightColor(color)
}
func (self Simple) GetAmbientLightColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetAmbientLightColor())
}
func (self Simple) SetAmbientSource(source classdb.EnvironmentAmbientSource) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAmbientSource(source)
}
func (self Simple) GetAmbientSource() classdb.EnvironmentAmbientSource {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.EnvironmentAmbientSource(Expert(self).GetAmbientSource())
}
func (self Simple) SetAmbientLightEnergy(energy float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAmbientLightEnergy(gd.Float(energy))
}
func (self Simple) GetAmbientLightEnergy() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAmbientLightEnergy()))
}
func (self Simple) SetAmbientLightSkyContribution(ratio float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAmbientLightSkyContribution(gd.Float(ratio))
}
func (self Simple) GetAmbientLightSkyContribution() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAmbientLightSkyContribution()))
}
func (self Simple) SetReflectionSource(source classdb.EnvironmentReflectionSource) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetReflectionSource(source)
}
func (self Simple) GetReflectionSource() classdb.EnvironmentReflectionSource {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.EnvironmentReflectionSource(Expert(self).GetReflectionSource())
}
func (self Simple) SetTonemapper(mode classdb.EnvironmentToneMapper) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTonemapper(mode)
}
func (self Simple) GetTonemapper() classdb.EnvironmentToneMapper {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.EnvironmentToneMapper(Expert(self).GetTonemapper())
}
func (self Simple) SetTonemapExposure(exposure float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTonemapExposure(gd.Float(exposure))
}
func (self Simple) GetTonemapExposure() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTonemapExposure()))
}
func (self Simple) SetTonemapWhite(white float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTonemapWhite(gd.Float(white))
}
func (self Simple) GetTonemapWhite() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTonemapWhite()))
}
func (self Simple) SetSsrEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSsrEnabled(enabled)
}
func (self Simple) IsSsrEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSsrEnabled())
}
func (self Simple) SetSsrMaxSteps(max_steps int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSsrMaxSteps(gd.Int(max_steps))
}
func (self Simple) GetSsrMaxSteps() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSsrMaxSteps()))
}
func (self Simple) SetSsrFadeIn(fade_in float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSsrFadeIn(gd.Float(fade_in))
}
func (self Simple) GetSsrFadeIn() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSsrFadeIn()))
}
func (self Simple) SetSsrFadeOut(fade_out float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSsrFadeOut(gd.Float(fade_out))
}
func (self Simple) GetSsrFadeOut() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSsrFadeOut()))
}
func (self Simple) SetSsrDepthTolerance(depth_tolerance float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSsrDepthTolerance(gd.Float(depth_tolerance))
}
func (self Simple) GetSsrDepthTolerance() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSsrDepthTolerance()))
}
func (self Simple) SetSsaoEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSsaoEnabled(enabled)
}
func (self Simple) IsSsaoEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSsaoEnabled())
}
func (self Simple) SetSsaoRadius(radius float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSsaoRadius(gd.Float(radius))
}
func (self Simple) GetSsaoRadius() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSsaoRadius()))
}
func (self Simple) SetSsaoIntensity(intensity float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSsaoIntensity(gd.Float(intensity))
}
func (self Simple) GetSsaoIntensity() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSsaoIntensity()))
}
func (self Simple) SetSsaoPower(power float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSsaoPower(gd.Float(power))
}
func (self Simple) GetSsaoPower() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSsaoPower()))
}
func (self Simple) SetSsaoDetail(detail float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSsaoDetail(gd.Float(detail))
}
func (self Simple) GetSsaoDetail() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSsaoDetail()))
}
func (self Simple) SetSsaoHorizon(horizon float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSsaoHorizon(gd.Float(horizon))
}
func (self Simple) GetSsaoHorizon() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSsaoHorizon()))
}
func (self Simple) SetSsaoSharpness(sharpness float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSsaoSharpness(gd.Float(sharpness))
}
func (self Simple) GetSsaoSharpness() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSsaoSharpness()))
}
func (self Simple) SetSsaoDirectLightAffect(amount float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSsaoDirectLightAffect(gd.Float(amount))
}
func (self Simple) GetSsaoDirectLightAffect() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSsaoDirectLightAffect()))
}
func (self Simple) SetSsaoAoChannelAffect(amount float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSsaoAoChannelAffect(gd.Float(amount))
}
func (self Simple) GetSsaoAoChannelAffect() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSsaoAoChannelAffect()))
}
func (self Simple) SetSsilEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSsilEnabled(enabled)
}
func (self Simple) IsSsilEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSsilEnabled())
}
func (self Simple) SetSsilRadius(radius float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSsilRadius(gd.Float(radius))
}
func (self Simple) GetSsilRadius() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSsilRadius()))
}
func (self Simple) SetSsilIntensity(intensity float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSsilIntensity(gd.Float(intensity))
}
func (self Simple) GetSsilIntensity() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSsilIntensity()))
}
func (self Simple) SetSsilSharpness(sharpness float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSsilSharpness(gd.Float(sharpness))
}
func (self Simple) GetSsilSharpness() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSsilSharpness()))
}
func (self Simple) SetSsilNormalRejection(normal_rejection float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSsilNormalRejection(gd.Float(normal_rejection))
}
func (self Simple) GetSsilNormalRejection() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSsilNormalRejection()))
}
func (self Simple) SetSdfgiEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSdfgiEnabled(enabled)
}
func (self Simple) IsSdfgiEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSdfgiEnabled())
}
func (self Simple) SetSdfgiCascades(amount int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSdfgiCascades(gd.Int(amount))
}
func (self Simple) GetSdfgiCascades() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSdfgiCascades()))
}
func (self Simple) SetSdfgiMinCellSize(size float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSdfgiMinCellSize(gd.Float(size))
}
func (self Simple) GetSdfgiMinCellSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSdfgiMinCellSize()))
}
func (self Simple) SetSdfgiMaxDistance(distance float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSdfgiMaxDistance(gd.Float(distance))
}
func (self Simple) GetSdfgiMaxDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSdfgiMaxDistance()))
}
func (self Simple) SetSdfgiCascade0Distance(distance float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSdfgiCascade0Distance(gd.Float(distance))
}
func (self Simple) GetSdfgiCascade0Distance() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSdfgiCascade0Distance()))
}
func (self Simple) SetSdfgiYScale(scale classdb.EnvironmentSDFGIYScale) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSdfgiYScale(scale)
}
func (self Simple) GetSdfgiYScale() classdb.EnvironmentSDFGIYScale {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.EnvironmentSDFGIYScale(Expert(self).GetSdfgiYScale())
}
func (self Simple) SetSdfgiUseOcclusion(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSdfgiUseOcclusion(enable)
}
func (self Simple) IsSdfgiUsingOcclusion() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSdfgiUsingOcclusion())
}
func (self Simple) SetSdfgiBounceFeedback(amount float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSdfgiBounceFeedback(gd.Float(amount))
}
func (self Simple) GetSdfgiBounceFeedback() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSdfgiBounceFeedback()))
}
func (self Simple) SetSdfgiReadSkyLight(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSdfgiReadSkyLight(enable)
}
func (self Simple) IsSdfgiReadingSkyLight() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSdfgiReadingSkyLight())
}
func (self Simple) SetSdfgiEnergy(amount float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSdfgiEnergy(gd.Float(amount))
}
func (self Simple) GetSdfgiEnergy() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSdfgiEnergy()))
}
func (self Simple) SetSdfgiNormalBias(bias float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSdfgiNormalBias(gd.Float(bias))
}
func (self Simple) GetSdfgiNormalBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSdfgiNormalBias()))
}
func (self Simple) SetSdfgiProbeBias(bias float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSdfgiProbeBias(gd.Float(bias))
}
func (self Simple) GetSdfgiProbeBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSdfgiProbeBias()))
}
func (self Simple) SetGlowEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlowEnabled(enabled)
}
func (self Simple) IsGlowEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsGlowEnabled())
}
func (self Simple) SetGlowLevel(idx int, intensity float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlowLevel(gd.Int(idx), gd.Float(intensity))
}
func (self Simple) GetGlowLevel(idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetGlowLevel(gd.Int(idx))))
}
func (self Simple) SetGlowNormalized(normalize bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlowNormalized(normalize)
}
func (self Simple) IsGlowNormalized() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsGlowNormalized())
}
func (self Simple) SetGlowIntensity(intensity float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlowIntensity(gd.Float(intensity))
}
func (self Simple) GetGlowIntensity() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetGlowIntensity()))
}
func (self Simple) SetGlowStrength(strength float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlowStrength(gd.Float(strength))
}
func (self Simple) GetGlowStrength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetGlowStrength()))
}
func (self Simple) SetGlowMix(mix float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlowMix(gd.Float(mix))
}
func (self Simple) GetGlowMix() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetGlowMix()))
}
func (self Simple) SetGlowBloom(amount float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlowBloom(gd.Float(amount))
}
func (self Simple) GetGlowBloom() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetGlowBloom()))
}
func (self Simple) SetGlowBlendMode(mode classdb.EnvironmentGlowBlendMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlowBlendMode(mode)
}
func (self Simple) GetGlowBlendMode() classdb.EnvironmentGlowBlendMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.EnvironmentGlowBlendMode(Expert(self).GetGlowBlendMode())
}
func (self Simple) SetGlowHdrBleedThreshold(threshold float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlowHdrBleedThreshold(gd.Float(threshold))
}
func (self Simple) GetGlowHdrBleedThreshold() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetGlowHdrBleedThreshold()))
}
func (self Simple) SetGlowHdrBleedScale(scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlowHdrBleedScale(gd.Float(scale))
}
func (self Simple) GetGlowHdrBleedScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetGlowHdrBleedScale()))
}
func (self Simple) SetGlowHdrLuminanceCap(amount float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlowHdrLuminanceCap(gd.Float(amount))
}
func (self Simple) GetGlowHdrLuminanceCap() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetGlowHdrLuminanceCap()))
}
func (self Simple) SetGlowMapStrength(strength float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlowMapStrength(gd.Float(strength))
}
func (self Simple) GetGlowMapStrength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetGlowMapStrength()))
}
func (self Simple) SetGlowMap(mode [1]classdb.Texture) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlowMap(mode)
}
func (self Simple) GetGlowMap() [1]classdb.Texture {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture(Expert(self).GetGlowMap(gc))
}
func (self Simple) SetFogEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFogEnabled(enabled)
}
func (self Simple) IsFogEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsFogEnabled())
}
func (self Simple) SetFogMode(mode classdb.EnvironmentFogMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFogMode(mode)
}
func (self Simple) GetFogMode() classdb.EnvironmentFogMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.EnvironmentFogMode(Expert(self).GetFogMode())
}
func (self Simple) SetFogLightColor(light_color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFogLightColor(light_color)
}
func (self Simple) GetFogLightColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetFogLightColor())
}
func (self Simple) SetFogLightEnergy(light_energy float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFogLightEnergy(gd.Float(light_energy))
}
func (self Simple) GetFogLightEnergy() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFogLightEnergy()))
}
func (self Simple) SetFogSunScatter(sun_scatter float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFogSunScatter(gd.Float(sun_scatter))
}
func (self Simple) GetFogSunScatter() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFogSunScatter()))
}
func (self Simple) SetFogDensity(density float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFogDensity(gd.Float(density))
}
func (self Simple) GetFogDensity() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFogDensity()))
}
func (self Simple) SetFogHeight(height float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFogHeight(gd.Float(height))
}
func (self Simple) GetFogHeight() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFogHeight()))
}
func (self Simple) SetFogHeightDensity(height_density float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFogHeightDensity(gd.Float(height_density))
}
func (self Simple) GetFogHeightDensity() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFogHeightDensity()))
}
func (self Simple) SetFogAerialPerspective(aerial_perspective float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFogAerialPerspective(gd.Float(aerial_perspective))
}
func (self Simple) GetFogAerialPerspective() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFogAerialPerspective()))
}
func (self Simple) SetFogSkyAffect(sky_affect float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFogSkyAffect(gd.Float(sky_affect))
}
func (self Simple) GetFogSkyAffect() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFogSkyAffect()))
}
func (self Simple) SetFogDepthCurve(curve float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFogDepthCurve(gd.Float(curve))
}
func (self Simple) GetFogDepthCurve() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFogDepthCurve()))
}
func (self Simple) SetFogDepthBegin(begin float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFogDepthBegin(gd.Float(begin))
}
func (self Simple) GetFogDepthBegin() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFogDepthBegin()))
}
func (self Simple) SetFogDepthEnd(end float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFogDepthEnd(gd.Float(end))
}
func (self Simple) GetFogDepthEnd() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFogDepthEnd()))
}
func (self Simple) SetVolumetricFogEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVolumetricFogEnabled(enabled)
}
func (self Simple) IsVolumetricFogEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsVolumetricFogEnabled())
}
func (self Simple) SetVolumetricFogEmission(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVolumetricFogEmission(color)
}
func (self Simple) GetVolumetricFogEmission() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetVolumetricFogEmission())
}
func (self Simple) SetVolumetricFogAlbedo(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVolumetricFogAlbedo(color)
}
func (self Simple) GetVolumetricFogAlbedo() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetVolumetricFogAlbedo())
}
func (self Simple) SetVolumetricFogDensity(density float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVolumetricFogDensity(gd.Float(density))
}
func (self Simple) GetVolumetricFogDensity() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVolumetricFogDensity()))
}
func (self Simple) SetVolumetricFogEmissionEnergy(begin float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVolumetricFogEmissionEnergy(gd.Float(begin))
}
func (self Simple) GetVolumetricFogEmissionEnergy() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVolumetricFogEmissionEnergy()))
}
func (self Simple) SetVolumetricFogAnisotropy(anisotropy float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVolumetricFogAnisotropy(gd.Float(anisotropy))
}
func (self Simple) GetVolumetricFogAnisotropy() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVolumetricFogAnisotropy()))
}
func (self Simple) SetVolumetricFogLength(length float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVolumetricFogLength(gd.Float(length))
}
func (self Simple) GetVolumetricFogLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVolumetricFogLength()))
}
func (self Simple) SetVolumetricFogDetailSpread(detail_spread float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVolumetricFogDetailSpread(gd.Float(detail_spread))
}
func (self Simple) GetVolumetricFogDetailSpread() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVolumetricFogDetailSpread()))
}
func (self Simple) SetVolumetricFogGiInject(gi_inject float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVolumetricFogGiInject(gd.Float(gi_inject))
}
func (self Simple) GetVolumetricFogGiInject() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVolumetricFogGiInject()))
}
func (self Simple) SetVolumetricFogAmbientInject(enabled float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVolumetricFogAmbientInject(gd.Float(enabled))
}
func (self Simple) GetVolumetricFogAmbientInject() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVolumetricFogAmbientInject()))
}
func (self Simple) SetVolumetricFogSkyAffect(sky_affect float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVolumetricFogSkyAffect(gd.Float(sky_affect))
}
func (self Simple) GetVolumetricFogSkyAffect() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVolumetricFogSkyAffect()))
}
func (self Simple) SetVolumetricFogTemporalReprojectionEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVolumetricFogTemporalReprojectionEnabled(enabled)
}
func (self Simple) IsVolumetricFogTemporalReprojectionEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsVolumetricFogTemporalReprojectionEnabled())
}
func (self Simple) SetVolumetricFogTemporalReprojectionAmount(temporal_reprojection_amount float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVolumetricFogTemporalReprojectionAmount(gd.Float(temporal_reprojection_amount))
}
func (self Simple) GetVolumetricFogTemporalReprojectionAmount() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVolumetricFogTemporalReprojectionAmount()))
}
func (self Simple) SetAdjustmentEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAdjustmentEnabled(enabled)
}
func (self Simple) IsAdjustmentEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAdjustmentEnabled())
}
func (self Simple) SetAdjustmentBrightness(brightness float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAdjustmentBrightness(gd.Float(brightness))
}
func (self Simple) GetAdjustmentBrightness() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAdjustmentBrightness()))
}
func (self Simple) SetAdjustmentContrast(contrast float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAdjustmentContrast(gd.Float(contrast))
}
func (self Simple) GetAdjustmentContrast() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAdjustmentContrast()))
}
func (self Simple) SetAdjustmentSaturation(saturation float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAdjustmentSaturation(gd.Float(saturation))
}
func (self Simple) GetAdjustmentSaturation() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAdjustmentSaturation()))
}
func (self Simple) SetAdjustmentColorCorrection(color_correction [1]classdb.Texture) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAdjustmentColorCorrection(color_correction)
}
func (self Simple) GetAdjustmentColorCorrection() [1]classdb.Texture {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture(Expert(self).GetAdjustmentColorCorrection(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Environment
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetBackground(mode classdb.EnvironmentBGMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_background, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBackground() classdb.EnvironmentBGMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentBGMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_background, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSky(sky object.Sky)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(sky[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sky, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSky(ctx gd.Lifetime) object.Sky {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sky, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Sky
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkyCustomFov(scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sky_custom_fov, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkyCustomFov() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sky_custom_fov, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkyRotation(euler_radians gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, euler_radians)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sky_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkyRotation() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sky_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBgColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBgColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBgEnergyMultiplier(energy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_bg_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBgEnergyMultiplier() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_bg_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBgIntensity(energy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_bg_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBgIntensity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_bg_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCanvasMaxLayer(layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_canvas_max_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCanvasMaxLayer() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_canvas_max_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCameraFeedId(id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_camera_feed_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCameraFeedId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_camera_feed_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAmbientLightColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ambient_light_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAmbientLightColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ambient_light_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAmbientSource(source classdb.EnvironmentAmbientSource)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, source)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ambient_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAmbientSource() classdb.EnvironmentAmbientSource {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentAmbientSource](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ambient_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAmbientLightEnergy(energy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ambient_light_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAmbientLightEnergy() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ambient_light_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAmbientLightSkyContribution(ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ambient_light_sky_contribution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAmbientLightSkyContribution() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ambient_light_sky_contribution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetReflectionSource(source classdb.EnvironmentReflectionSource)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, source)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_reflection_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetReflectionSource() classdb.EnvironmentReflectionSource {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentReflectionSource](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_reflection_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTonemapper(mode classdb.EnvironmentToneMapper)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_tonemapper, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTonemapper() classdb.EnvironmentToneMapper {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentToneMapper](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_tonemapper, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTonemapExposure(exposure gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, exposure)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_tonemap_exposure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTonemapExposure() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_tonemap_exposure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTonemapWhite(white gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, white)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_tonemap_white, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTonemapWhite() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_tonemap_white, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsrEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssr_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSsrEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_ssr_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsrMaxSteps(max_steps gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_steps)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssr_max_steps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsrMaxSteps() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssr_max_steps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsrFadeIn(fade_in gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fade_in)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssr_fade_in, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsrFadeIn() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssr_fade_in, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsrFadeOut(fade_out gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fade_out)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssr_fade_out, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsrFadeOut() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssr_fade_out, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsrDepthTolerance(depth_tolerance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, depth_tolerance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssr_depth_tolerance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsrDepthTolerance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssr_depth_tolerance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsaoEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssao_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSsaoEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_ssao_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsaoRadius(radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssao_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsaoRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssao_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsaoIntensity(intensity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, intensity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssao_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsaoIntensity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssao_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsaoPower(power gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, power)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssao_power, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsaoPower() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssao_power, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsaoDetail(detail gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, detail)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssao_detail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsaoDetail() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssao_detail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsaoHorizon(horizon gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, horizon)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssao_horizon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsaoHorizon() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssao_horizon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsaoSharpness(sharpness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sharpness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssao_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsaoSharpness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssao_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsaoDirectLightAffect(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssao_direct_light_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsaoDirectLightAffect() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssao_direct_light_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsaoAoChannelAffect(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssao_ao_channel_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsaoAoChannelAffect() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssao_ao_channel_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsilEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssil_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSsilEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_ssil_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsilRadius(radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssil_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsilRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssil_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsilIntensity(intensity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, intensity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssil_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsilIntensity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssil_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsilSharpness(sharpness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sharpness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssil_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsilSharpness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssil_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsilNormalRejection(normal_rejection gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, normal_rejection)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssil_normal_rejection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsilNormalRejection() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssil_normal_rejection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSdfgiEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_sdfgi_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiCascades(amount gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_cascades, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfgiCascades() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sdfgi_cascades, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiMinCellSize(size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_min_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfgiMinCellSize() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sdfgi_min_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiMaxDistance(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_max_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfgiMaxDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sdfgi_max_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiCascade0Distance(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_cascade0_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfgiCascade0Distance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sdfgi_cascade0_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiYScale(scale classdb.EnvironmentSDFGIYScale)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_y_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfgiYScale() classdb.EnvironmentSDFGIYScale {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentSDFGIYScale](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sdfgi_y_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiUseOcclusion(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_use_occlusion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSdfgiUsingOcclusion() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_sdfgi_using_occlusion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiBounceFeedback(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_bounce_feedback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfgiBounceFeedback() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sdfgi_bounce_feedback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiReadSkyLight(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_read_sky_light, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSdfgiReadingSkyLight() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_sdfgi_reading_sky_light, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiEnergy(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfgiEnergy() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sdfgi_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiNormalBias(bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_normal_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfgiNormalBias() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sdfgi_normal_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiProbeBias(bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_probe_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfgiProbeBias() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sdfgi_probe_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsGlowEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_glow_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the intensity of the glow level [param idx]. A value above [code]0.0[/code] enables the level. Each level relies on the previous level. This means that enabling higher glow levels will slow down the glow effect rendering, even if previous levels aren't enabled.
*/
//go:nosplit
func (self class) SetGlowLevel(idx gd.Int, intensity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, intensity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the intensity of the glow level [param idx].
*/
//go:nosplit
func (self class) GetGlowLevel(idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowNormalized(normalize bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, normalize)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_normalized, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsGlowNormalized() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_glow_normalized, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowIntensity(intensity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, intensity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowIntensity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowStrength(strength gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowStrength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowMix(mix gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mix)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_mix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowMix() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_mix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowBloom(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_bloom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowBloom() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_bloom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowBlendMode(mode classdb.EnvironmentGlowBlendMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowBlendMode() classdb.EnvironmentGlowBlendMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentGlowBlendMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowHdrBleedThreshold(threshold gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, threshold)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_hdr_bleed_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowHdrBleedThreshold() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_hdr_bleed_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowHdrBleedScale(scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_hdr_bleed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowHdrBleedScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_hdr_bleed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowHdrLuminanceCap(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_hdr_luminance_cap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowHdrLuminanceCap() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_hdr_luminance_cap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowMapStrength(strength gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_map_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowMapStrength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_map_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowMap(mode object.Texture)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(mode[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowMap(ctx gd.Lifetime) object.Texture {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFogEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_fog_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogMode(mode classdb.EnvironmentFogMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogMode() classdb.EnvironmentFogMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentFogMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogLightColor(light_color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light_color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_light_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogLightColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_light_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogLightEnergy(light_energy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light_energy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_light_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogLightEnergy() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_light_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogSunScatter(sun_scatter gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sun_scatter)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_sun_scatter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogSunScatter() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_sun_scatter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogDensity(density gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, density)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogDensity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogHeight(height gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogHeight() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogHeightDensity(height_density gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, height_density)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_height_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogHeightDensity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_height_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogAerialPerspective(aerial_perspective gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, aerial_perspective)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_aerial_perspective, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogAerialPerspective() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_aerial_perspective, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogSkyAffect(sky_affect gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sky_affect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_sky_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogSkyAffect() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_sky_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogDepthCurve(curve gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, curve)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_depth_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogDepthCurve() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_depth_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogDepthBegin(begin gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, begin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_depth_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogDepthBegin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_depth_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogDepthEnd(end gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, end)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_depth_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogDepthEnd() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_depth_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsVolumetricFogEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_volumetric_fog_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogEmission(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_emission, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogEmission() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_emission, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogAlbedo(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_albedo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogAlbedo() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_albedo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogDensity(density gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, density)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogDensity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogEmissionEnergy(begin gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, begin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_emission_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogEmissionEnergy() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_emission_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogAnisotropy(anisotropy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, anisotropy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_anisotropy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogAnisotropy() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_anisotropy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogLength(length gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogDetailSpread(detail_spread gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, detail_spread)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_detail_spread, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogDetailSpread() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_detail_spread, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogGiInject(gi_inject gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gi_inject)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_gi_inject, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogGiInject() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_gi_inject, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogAmbientInject(enabled gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_ambient_inject, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogAmbientInject() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_ambient_inject, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogSkyAffect(sky_affect gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sky_affect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_sky_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogSkyAffect() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_sky_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogTemporalReprojectionEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_temporal_reprojection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsVolumetricFogTemporalReprojectionEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_volumetric_fog_temporal_reprojection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogTemporalReprojectionAmount(temporal_reprojection_amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, temporal_reprojection_amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_temporal_reprojection_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogTemporalReprojectionAmount() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_temporal_reprojection_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAdjustmentEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_adjustment_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAdjustmentEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_adjustment_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAdjustmentBrightness(brightness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, brightness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_adjustment_brightness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAdjustmentBrightness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_adjustment_brightness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAdjustmentContrast(contrast gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, contrast)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_adjustment_contrast, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAdjustmentContrast() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_adjustment_contrast, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAdjustmentSaturation(saturation gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, saturation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_adjustment_saturation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAdjustmentSaturation() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_adjustment_saturation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAdjustmentColorCorrection(color_correction object.Texture)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(color_correction[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_adjustment_color_correction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAdjustmentColorCorrection(ctx gd.Lifetime) object.Texture {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_adjustment_color_correction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsEnvironment() Expert { return self[0].AsEnvironment() }


//go:nosplit
func (self Simple) AsEnvironment() Simple { return self[0].AsEnvironment() }


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
func init() {classdb.Register("Environment", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type BGMode = classdb.EnvironmentBGMode

const (
/*Clears the background using the clear color defined in [member ProjectSettings.rendering/environment/defaults/default_clear_color].*/
	BgClearColor BGMode = 0
/*Clears the background using a custom clear color.*/
	BgColor BGMode = 1
/*Displays a user-defined sky in the background.*/
	BgSky BGMode = 2
/*Displays a [CanvasLayer] in the background.*/
	BgCanvas BGMode = 3
/*Keeps on screen every pixel drawn in the background. This is the fastest background mode, but it can only be safely used in fully-interior scenes (no visible sky or sky reflections). If enabled in a scene where the background is visible, "ghost trail" artifacts will be visible when moving the camera.*/
	BgKeep BGMode = 4
/*Displays a camera feed in the background.*/
	BgCameraFeed BGMode = 5
/*Represents the size of the [enum BGMode] enum.*/
	BgMax BGMode = 6
)
type AmbientSource = classdb.EnvironmentAmbientSource

const (
/*Gather ambient light from whichever source is specified as the background.*/
	AmbientSourceBg AmbientSource = 0
/*Disable ambient light. This provides a slight performance boost over [constant AMBIENT_SOURCE_SKY].*/
	AmbientSourceDisabled AmbientSource = 1
/*Specify a specific [Color] for ambient light. This provides a slight performance boost over [constant AMBIENT_SOURCE_SKY].*/
	AmbientSourceColor AmbientSource = 2
/*Gather ambient light from the [Sky] regardless of what the background is.*/
	AmbientSourceSky AmbientSource = 3
)
type ReflectionSource = classdb.EnvironmentReflectionSource

const (
/*Use the background for reflections.*/
	ReflectionSourceBg ReflectionSource = 0
/*Disable reflections. This provides a slight performance boost over other options.*/
	ReflectionSourceDisabled ReflectionSource = 1
/*Use the [Sky] for reflections regardless of what the background is.*/
	ReflectionSourceSky ReflectionSource = 2
)
type ToneMapper = classdb.EnvironmentToneMapper

const (
/*Linear tonemapper operator. Reads the linear data and passes it on unmodified. This can cause bright lighting to look blown out, with noticeable clipping in the output colors.*/
	ToneMapperLinear ToneMapper = 0
/*Reinhardt tonemapper operator. Performs a variation on rendered pixels' colors by this formula: [code]color = color / (1 + color)[/code]. This avoids clipping bright highlights, but the resulting image can look a bit dull.*/
	ToneMapperReinhardt ToneMapper = 1
/*Filmic tonemapper operator. This avoids clipping bright highlights, with a resulting image that usually looks more vivid than [constant TONE_MAPPER_REINHARDT].*/
	ToneMapperFilmic ToneMapper = 2
/*Use the Academy Color Encoding System tonemapper. ACES is slightly more expensive than other options, but it handles bright lighting in a more realistic fashion by desaturating it as it becomes brighter. ACES typically has a more contrasted output compared to [constant TONE_MAPPER_REINHARDT] and [constant TONE_MAPPER_FILMIC].
[b]Note:[/b] This tonemapping operator is called "ACES Fitted" in Godot 3.x.*/
	ToneMapperAces ToneMapper = 3
)
type GlowBlendMode = classdb.EnvironmentGlowBlendMode

const (
/*Additive glow blending mode. Mostly used for particles, glows (bloom), lens flare, bright sources.*/
	GlowBlendModeAdditive GlowBlendMode = 0
/*Screen glow blending mode. Increases brightness, used frequently with bloom.*/
	GlowBlendModeScreen GlowBlendMode = 1
/*Soft light glow blending mode. Modifies contrast, exposes shadows and highlights (vivid bloom).*/
	GlowBlendModeSoftlight GlowBlendMode = 2
/*Replace glow blending mode. Replaces all pixels' color by the glow value. This can be used to simulate a full-screen blur effect by tweaking the glow parameters to match the original image's brightness.*/
	GlowBlendModeReplace GlowBlendMode = 3
/*Mixes the glow with the underlying color to avoid increasing brightness as much while still maintaining a glow effect.*/
	GlowBlendModeMix GlowBlendMode = 4
)
type FogMode = classdb.EnvironmentFogMode

const (
/*Use a physically-based fog model defined primarily by fog density.*/
	FogModeExponential FogMode = 0
/*Use a simple fog model defined by start and end positions and a custom curve. While not physically accurate, this model can be useful when you need more artistic control.*/
	FogModeDepth FogMode = 1
)
type SDFGIYScale = classdb.EnvironmentSDFGIYScale

const (
/*Use 50% scale for SDFGI on the Y (vertical) axis. SDFGI cells will be twice as short as they are wide. This allows providing increased GI detail and reduced light leaking with thin floors and ceilings. This is usually the best choice for scenes that don't feature much verticality.*/
	SdfgiYScale50Percent SDFGIYScale = 0
/*Use 75% scale for SDFGI on the Y (vertical) axis. This is a balance between the 50% and 100% SDFGI Y scales.*/
	SdfgiYScale75Percent SDFGIYScale = 1
/*Use 100% scale for SDFGI on the Y (vertical) axis. SDFGI cells will be as tall as they are wide. This is usually the best choice for highly vertical scenes. The downside is that light leaking may become more noticeable with thin floors and ceilings.*/
	SdfgiYScale100Percent SDFGIYScale = 2
)
