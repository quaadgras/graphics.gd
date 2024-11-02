package Environment

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Resource for environment nodes (like [WorldEnvironment]) that define multiple environment operations (such as background [Sky] or [Color], ambient light, fog, depth-of-field...). These parameters affect the final render of the scene. The order of these operations is:
- Depth of Field Blur
- Glow
- Tonemap (Auto Exposure)
- Adjustments
*/
type Instance [1]classdb.Environment

/*
Sets the intensity of the glow level [param idx]. A value above [code]0.0[/code] enables the level. Each level relies on the previous level. This means that enabling higher glow levels will slow down the glow effect rendering, even if previous levels aren't enabled.
*/
func (self Instance) SetGlowLevel(idx int, intensity float64) {
	class(self).SetGlowLevel(gd.Int(idx), gd.Float(intensity))
}

/*
Returns the intensity of the glow level [param idx].
*/
func (self Instance) GetGlowLevel(idx int) float64 {
	return float64(float64(class(self).GetGlowLevel(gd.Int(idx))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Environment

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Environment"))
	return Instance{classdb.Environment(object)}
}

func (self Instance) BackgroundMode() classdb.EnvironmentBGMode {
	return classdb.EnvironmentBGMode(class(self).GetBackground())
}

func (self Instance) SetBackgroundMode(value classdb.EnvironmentBGMode) {
	class(self).SetBackground(value)
}

func (self Instance) BackgroundColor() gd.Color {
	return gd.Color(class(self).GetBgColor())
}

func (self Instance) SetBackgroundColor(value gd.Color) {
	class(self).SetBgColor(value)
}

func (self Instance) BackgroundEnergyMultiplier() float64 {
	return float64(float64(class(self).GetBgEnergyMultiplier()))
}

func (self Instance) SetBackgroundEnergyMultiplier(value float64) {
	class(self).SetBgEnergyMultiplier(gd.Float(value))
}

func (self Instance) BackgroundIntensity() float64 {
	return float64(float64(class(self).GetBgIntensity()))
}

func (self Instance) SetBackgroundIntensity(value float64) {
	class(self).SetBgIntensity(gd.Float(value))
}

func (self Instance) BackgroundCanvasMaxLayer() int {
	return int(int(class(self).GetCanvasMaxLayer()))
}

func (self Instance) SetBackgroundCanvasMaxLayer(value int) {
	class(self).SetCanvasMaxLayer(gd.Int(value))
}

func (self Instance) BackgroundCameraFeedId() int {
	return int(int(class(self).GetCameraFeedId()))
}

func (self Instance) SetBackgroundCameraFeedId(value int) {
	class(self).SetCameraFeedId(gd.Int(value))
}

func (self Instance) Sky() objects.Sky {
	return objects.Sky(class(self).GetSky())
}

func (self Instance) SetSky(value objects.Sky) {
	class(self).SetSky(value)
}

func (self Instance) SkyCustomFov() float64 {
	return float64(float64(class(self).GetSkyCustomFov()))
}

func (self Instance) SetSkyCustomFov(value float64) {
	class(self).SetSkyCustomFov(gd.Float(value))
}

func (self Instance) SkyRotation() gd.Vector3 {
	return gd.Vector3(class(self).GetSkyRotation())
}

func (self Instance) SetSkyRotation(value gd.Vector3) {
	class(self).SetSkyRotation(value)
}

func (self Instance) AmbientLightSource() classdb.EnvironmentAmbientSource {
	return classdb.EnvironmentAmbientSource(class(self).GetAmbientSource())
}

func (self Instance) SetAmbientLightSource(value classdb.EnvironmentAmbientSource) {
	class(self).SetAmbientSource(value)
}

func (self Instance) AmbientLightColor() gd.Color {
	return gd.Color(class(self).GetAmbientLightColor())
}

func (self Instance) SetAmbientLightColor(value gd.Color) {
	class(self).SetAmbientLightColor(value)
}

func (self Instance) AmbientLightSkyContribution() float64 {
	return float64(float64(class(self).GetAmbientLightSkyContribution()))
}

func (self Instance) SetAmbientLightSkyContribution(value float64) {
	class(self).SetAmbientLightSkyContribution(gd.Float(value))
}

func (self Instance) AmbientLightEnergy() float64 {
	return float64(float64(class(self).GetAmbientLightEnergy()))
}

func (self Instance) SetAmbientLightEnergy(value float64) {
	class(self).SetAmbientLightEnergy(gd.Float(value))
}

func (self Instance) ReflectedLightSource() classdb.EnvironmentReflectionSource {
	return classdb.EnvironmentReflectionSource(class(self).GetReflectionSource())
}

func (self Instance) SetReflectedLightSource(value classdb.EnvironmentReflectionSource) {
	class(self).SetReflectionSource(value)
}

func (self Instance) TonemapMode() classdb.EnvironmentToneMapper {
	return classdb.EnvironmentToneMapper(class(self).GetTonemapper())
}

func (self Instance) SetTonemapMode(value classdb.EnvironmentToneMapper) {
	class(self).SetTonemapper(value)
}

func (self Instance) TonemapExposure() float64 {
	return float64(float64(class(self).GetTonemapExposure()))
}

func (self Instance) SetTonemapExposure(value float64) {
	class(self).SetTonemapExposure(gd.Float(value))
}

func (self Instance) TonemapWhite() float64 {
	return float64(float64(class(self).GetTonemapWhite()))
}

func (self Instance) SetTonemapWhite(value float64) {
	class(self).SetTonemapWhite(gd.Float(value))
}

func (self Instance) SsrEnabled() bool {
	return bool(class(self).IsSsrEnabled())
}

func (self Instance) SetSsrEnabled(value bool) {
	class(self).SetSsrEnabled(value)
}

func (self Instance) SsrMaxSteps() int {
	return int(int(class(self).GetSsrMaxSteps()))
}

func (self Instance) SetSsrMaxSteps(value int) {
	class(self).SetSsrMaxSteps(gd.Int(value))
}

func (self Instance) SsrFadeIn() float64 {
	return float64(float64(class(self).GetSsrFadeIn()))
}

func (self Instance) SetSsrFadeIn(value float64) {
	class(self).SetSsrFadeIn(gd.Float(value))
}

func (self Instance) SsrFadeOut() float64 {
	return float64(float64(class(self).GetSsrFadeOut()))
}

func (self Instance) SetSsrFadeOut(value float64) {
	class(self).SetSsrFadeOut(gd.Float(value))
}

func (self Instance) SsrDepthTolerance() float64 {
	return float64(float64(class(self).GetSsrDepthTolerance()))
}

func (self Instance) SetSsrDepthTolerance(value float64) {
	class(self).SetSsrDepthTolerance(gd.Float(value))
}

func (self Instance) SsaoEnabled() bool {
	return bool(class(self).IsSsaoEnabled())
}

func (self Instance) SetSsaoEnabled(value bool) {
	class(self).SetSsaoEnabled(value)
}

func (self Instance) SsaoRadius() float64 {
	return float64(float64(class(self).GetSsaoRadius()))
}

func (self Instance) SetSsaoRadius(value float64) {
	class(self).SetSsaoRadius(gd.Float(value))
}

func (self Instance) SsaoIntensity() float64 {
	return float64(float64(class(self).GetSsaoIntensity()))
}

func (self Instance) SetSsaoIntensity(value float64) {
	class(self).SetSsaoIntensity(gd.Float(value))
}

func (self Instance) SsaoPower() float64 {
	return float64(float64(class(self).GetSsaoPower()))
}

func (self Instance) SetSsaoPower(value float64) {
	class(self).SetSsaoPower(gd.Float(value))
}

func (self Instance) SsaoDetail() float64 {
	return float64(float64(class(self).GetSsaoDetail()))
}

func (self Instance) SetSsaoDetail(value float64) {
	class(self).SetSsaoDetail(gd.Float(value))
}

func (self Instance) SsaoHorizon() float64 {
	return float64(float64(class(self).GetSsaoHorizon()))
}

func (self Instance) SetSsaoHorizon(value float64) {
	class(self).SetSsaoHorizon(gd.Float(value))
}

func (self Instance) SsaoSharpness() float64 {
	return float64(float64(class(self).GetSsaoSharpness()))
}

func (self Instance) SetSsaoSharpness(value float64) {
	class(self).SetSsaoSharpness(gd.Float(value))
}

func (self Instance) SsaoLightAffect() float64 {
	return float64(float64(class(self).GetSsaoDirectLightAffect()))
}

func (self Instance) SetSsaoLightAffect(value float64) {
	class(self).SetSsaoDirectLightAffect(gd.Float(value))
}

func (self Instance) SsaoAoChannelAffect() float64 {
	return float64(float64(class(self).GetSsaoAoChannelAffect()))
}

func (self Instance) SetSsaoAoChannelAffect(value float64) {
	class(self).SetSsaoAoChannelAffect(gd.Float(value))
}

func (self Instance) SsilEnabled() bool {
	return bool(class(self).IsSsilEnabled())
}

func (self Instance) SetSsilEnabled(value bool) {
	class(self).SetSsilEnabled(value)
}

func (self Instance) SsilRadius() float64 {
	return float64(float64(class(self).GetSsilRadius()))
}

func (self Instance) SetSsilRadius(value float64) {
	class(self).SetSsilRadius(gd.Float(value))
}

func (self Instance) SsilIntensity() float64 {
	return float64(float64(class(self).GetSsilIntensity()))
}

func (self Instance) SetSsilIntensity(value float64) {
	class(self).SetSsilIntensity(gd.Float(value))
}

func (self Instance) SsilSharpness() float64 {
	return float64(float64(class(self).GetSsilSharpness()))
}

func (self Instance) SetSsilSharpness(value float64) {
	class(self).SetSsilSharpness(gd.Float(value))
}

func (self Instance) SsilNormalRejection() float64 {
	return float64(float64(class(self).GetSsilNormalRejection()))
}

func (self Instance) SetSsilNormalRejection(value float64) {
	class(self).SetSsilNormalRejection(gd.Float(value))
}

func (self Instance) SdfgiEnabled() bool {
	return bool(class(self).IsSdfgiEnabled())
}

func (self Instance) SetSdfgiEnabled(value bool) {
	class(self).SetSdfgiEnabled(value)
}

func (self Instance) SdfgiUseOcclusion() bool {
	return bool(class(self).IsSdfgiUsingOcclusion())
}

func (self Instance) SetSdfgiUseOcclusion(value bool) {
	class(self).SetSdfgiUseOcclusion(value)
}

func (self Instance) SdfgiReadSkyLight() bool {
	return bool(class(self).IsSdfgiReadingSkyLight())
}

func (self Instance) SetSdfgiReadSkyLight(value bool) {
	class(self).SetSdfgiReadSkyLight(value)
}

func (self Instance) SdfgiBounceFeedback() float64 {
	return float64(float64(class(self).GetSdfgiBounceFeedback()))
}

func (self Instance) SetSdfgiBounceFeedback(value float64) {
	class(self).SetSdfgiBounceFeedback(gd.Float(value))
}

func (self Instance) SdfgiCascades() int {
	return int(int(class(self).GetSdfgiCascades()))
}

func (self Instance) SetSdfgiCascades(value int) {
	class(self).SetSdfgiCascades(gd.Int(value))
}

func (self Instance) SdfgiMinCellSize() float64 {
	return float64(float64(class(self).GetSdfgiMinCellSize()))
}

func (self Instance) SetSdfgiMinCellSize(value float64) {
	class(self).SetSdfgiMinCellSize(gd.Float(value))
}

func (self Instance) SdfgiCascade0Distance() float64 {
	return float64(float64(class(self).GetSdfgiCascade0Distance()))
}

func (self Instance) SetSdfgiCascade0Distance(value float64) {
	class(self).SetSdfgiCascade0Distance(gd.Float(value))
}

func (self Instance) SdfgiMaxDistance() float64 {
	return float64(float64(class(self).GetSdfgiMaxDistance()))
}

func (self Instance) SetSdfgiMaxDistance(value float64) {
	class(self).SetSdfgiMaxDistance(gd.Float(value))
}

func (self Instance) SdfgiYScale() classdb.EnvironmentSDFGIYScale {
	return classdb.EnvironmentSDFGIYScale(class(self).GetSdfgiYScale())
}

func (self Instance) SetSdfgiYScale(value classdb.EnvironmentSDFGIYScale) {
	class(self).SetSdfgiYScale(value)
}

func (self Instance) SdfgiEnergy() float64 {
	return float64(float64(class(self).GetSdfgiEnergy()))
}

func (self Instance) SetSdfgiEnergy(value float64) {
	class(self).SetSdfgiEnergy(gd.Float(value))
}

func (self Instance) SdfgiNormalBias() float64 {
	return float64(float64(class(self).GetSdfgiNormalBias()))
}

func (self Instance) SetSdfgiNormalBias(value float64) {
	class(self).SetSdfgiNormalBias(gd.Float(value))
}

func (self Instance) SdfgiProbeBias() float64 {
	return float64(float64(class(self).GetSdfgiProbeBias()))
}

func (self Instance) SetSdfgiProbeBias(value float64) {
	class(self).SetSdfgiProbeBias(gd.Float(value))
}

func (self Instance) GlowEnabled() bool {
	return bool(class(self).IsGlowEnabled())
}

func (self Instance) SetGlowEnabled(value bool) {
	class(self).SetGlowEnabled(value)
}

func (self Instance) GlowNormalized() bool {
	return bool(class(self).IsGlowNormalized())
}

func (self Instance) SetGlowNormalized(value bool) {
	class(self).SetGlowNormalized(value)
}

func (self Instance) GlowIntensity() float64 {
	return float64(float64(class(self).GetGlowIntensity()))
}

func (self Instance) SetGlowIntensity(value float64) {
	class(self).SetGlowIntensity(gd.Float(value))
}

func (self Instance) GlowStrength() float64 {
	return float64(float64(class(self).GetGlowStrength()))
}

func (self Instance) SetGlowStrength(value float64) {
	class(self).SetGlowStrength(gd.Float(value))
}

func (self Instance) GlowMix() float64 {
	return float64(float64(class(self).GetGlowMix()))
}

func (self Instance) SetGlowMix(value float64) {
	class(self).SetGlowMix(gd.Float(value))
}

func (self Instance) GlowBloom() float64 {
	return float64(float64(class(self).GetGlowBloom()))
}

func (self Instance) SetGlowBloom(value float64) {
	class(self).SetGlowBloom(gd.Float(value))
}

func (self Instance) GlowBlendMode() classdb.EnvironmentGlowBlendMode {
	return classdb.EnvironmentGlowBlendMode(class(self).GetGlowBlendMode())
}

func (self Instance) SetGlowBlendMode(value classdb.EnvironmentGlowBlendMode) {
	class(self).SetGlowBlendMode(value)
}

func (self Instance) GlowHdrThreshold() float64 {
	return float64(float64(class(self).GetGlowHdrBleedThreshold()))
}

func (self Instance) SetGlowHdrThreshold(value float64) {
	class(self).SetGlowHdrBleedThreshold(gd.Float(value))
}

func (self Instance) GlowHdrScale() float64 {
	return float64(float64(class(self).GetGlowHdrBleedScale()))
}

func (self Instance) SetGlowHdrScale(value float64) {
	class(self).SetGlowHdrBleedScale(gd.Float(value))
}

func (self Instance) GlowHdrLuminanceCap() float64 {
	return float64(float64(class(self).GetGlowHdrLuminanceCap()))
}

func (self Instance) SetGlowHdrLuminanceCap(value float64) {
	class(self).SetGlowHdrLuminanceCap(gd.Float(value))
}

func (self Instance) GlowMapStrength() float64 {
	return float64(float64(class(self).GetGlowMapStrength()))
}

func (self Instance) SetGlowMapStrength(value float64) {
	class(self).SetGlowMapStrength(gd.Float(value))
}

func (self Instance) GlowMap() objects.Texture {
	return objects.Texture(class(self).GetGlowMap())
}

func (self Instance) SetGlowMap(value objects.Texture) {
	class(self).SetGlowMap(value)
}

func (self Instance) FogEnabled() bool {
	return bool(class(self).IsFogEnabled())
}

func (self Instance) SetFogEnabled(value bool) {
	class(self).SetFogEnabled(value)
}

func (self Instance) FogMode() classdb.EnvironmentFogMode {
	return classdb.EnvironmentFogMode(class(self).GetFogMode())
}

func (self Instance) SetFogMode(value classdb.EnvironmentFogMode) {
	class(self).SetFogMode(value)
}

func (self Instance) FogLightColor() gd.Color {
	return gd.Color(class(self).GetFogLightColor())
}

func (self Instance) SetFogLightColor(value gd.Color) {
	class(self).SetFogLightColor(value)
}

func (self Instance) FogLightEnergy() float64 {
	return float64(float64(class(self).GetFogLightEnergy()))
}

func (self Instance) SetFogLightEnergy(value float64) {
	class(self).SetFogLightEnergy(gd.Float(value))
}

func (self Instance) FogSunScatter() float64 {
	return float64(float64(class(self).GetFogSunScatter()))
}

func (self Instance) SetFogSunScatter(value float64) {
	class(self).SetFogSunScatter(gd.Float(value))
}

func (self Instance) FogDensity() float64 {
	return float64(float64(class(self).GetFogDensity()))
}

func (self Instance) SetFogDensity(value float64) {
	class(self).SetFogDensity(gd.Float(value))
}

func (self Instance) FogAerialPerspective() float64 {
	return float64(float64(class(self).GetFogAerialPerspective()))
}

func (self Instance) SetFogAerialPerspective(value float64) {
	class(self).SetFogAerialPerspective(gd.Float(value))
}

func (self Instance) FogSkyAffect() float64 {
	return float64(float64(class(self).GetFogSkyAffect()))
}

func (self Instance) SetFogSkyAffect(value float64) {
	class(self).SetFogSkyAffect(gd.Float(value))
}

func (self Instance) FogHeight() float64 {
	return float64(float64(class(self).GetFogHeight()))
}

func (self Instance) SetFogHeight(value float64) {
	class(self).SetFogHeight(gd.Float(value))
}

func (self Instance) FogHeightDensity() float64 {
	return float64(float64(class(self).GetFogHeightDensity()))
}

func (self Instance) SetFogHeightDensity(value float64) {
	class(self).SetFogHeightDensity(gd.Float(value))
}

func (self Instance) FogDepthCurve() float64 {
	return float64(float64(class(self).GetFogDepthCurve()))
}

func (self Instance) SetFogDepthCurve(value float64) {
	class(self).SetFogDepthCurve(gd.Float(value))
}

func (self Instance) FogDepthBegin() float64 {
	return float64(float64(class(self).GetFogDepthBegin()))
}

func (self Instance) SetFogDepthBegin(value float64) {
	class(self).SetFogDepthBegin(gd.Float(value))
}

func (self Instance) FogDepthEnd() float64 {
	return float64(float64(class(self).GetFogDepthEnd()))
}

func (self Instance) SetFogDepthEnd(value float64) {
	class(self).SetFogDepthEnd(gd.Float(value))
}

func (self Instance) VolumetricFogEnabled() bool {
	return bool(class(self).IsVolumetricFogEnabled())
}

func (self Instance) SetVolumetricFogEnabled(value bool) {
	class(self).SetVolumetricFogEnabled(value)
}

func (self Instance) VolumetricFogDensity() float64 {
	return float64(float64(class(self).GetVolumetricFogDensity()))
}

func (self Instance) SetVolumetricFogDensity(value float64) {
	class(self).SetVolumetricFogDensity(gd.Float(value))
}

func (self Instance) VolumetricFogAlbedo() gd.Color {
	return gd.Color(class(self).GetVolumetricFogAlbedo())
}

func (self Instance) SetVolumetricFogAlbedo(value gd.Color) {
	class(self).SetVolumetricFogAlbedo(value)
}

func (self Instance) VolumetricFogEmission() gd.Color {
	return gd.Color(class(self).GetVolumetricFogEmission())
}

func (self Instance) SetVolumetricFogEmission(value gd.Color) {
	class(self).SetVolumetricFogEmission(value)
}

func (self Instance) VolumetricFogEmissionEnergy() float64 {
	return float64(float64(class(self).GetVolumetricFogEmissionEnergy()))
}

func (self Instance) SetVolumetricFogEmissionEnergy(value float64) {
	class(self).SetVolumetricFogEmissionEnergy(gd.Float(value))
}

func (self Instance) VolumetricFogGiInject() float64 {
	return float64(float64(class(self).GetVolumetricFogGiInject()))
}

func (self Instance) SetVolumetricFogGiInject(value float64) {
	class(self).SetVolumetricFogGiInject(gd.Float(value))
}

func (self Instance) VolumetricFogAnisotropy() float64 {
	return float64(float64(class(self).GetVolumetricFogAnisotropy()))
}

func (self Instance) SetVolumetricFogAnisotropy(value float64) {
	class(self).SetVolumetricFogAnisotropy(gd.Float(value))
}

func (self Instance) VolumetricFogLength() float64 {
	return float64(float64(class(self).GetVolumetricFogLength()))
}

func (self Instance) SetVolumetricFogLength(value float64) {
	class(self).SetVolumetricFogLength(gd.Float(value))
}

func (self Instance) VolumetricFogDetailSpread() float64 {
	return float64(float64(class(self).GetVolumetricFogDetailSpread()))
}

func (self Instance) SetVolumetricFogDetailSpread(value float64) {
	class(self).SetVolumetricFogDetailSpread(gd.Float(value))
}

func (self Instance) VolumetricFogAmbientInject() float64 {
	return float64(float64(class(self).GetVolumetricFogAmbientInject()))
}

func (self Instance) SetVolumetricFogAmbientInject(value float64) {
	class(self).SetVolumetricFogAmbientInject(gd.Float(value))
}

func (self Instance) VolumetricFogSkyAffect() float64 {
	return float64(float64(class(self).GetVolumetricFogSkyAffect()))
}

func (self Instance) SetVolumetricFogSkyAffect(value float64) {
	class(self).SetVolumetricFogSkyAffect(gd.Float(value))
}

func (self Instance) VolumetricFogTemporalReprojectionEnabled() bool {
	return bool(class(self).IsVolumetricFogTemporalReprojectionEnabled())
}

func (self Instance) SetVolumetricFogTemporalReprojectionEnabled(value bool) {
	class(self).SetVolumetricFogTemporalReprojectionEnabled(value)
}

func (self Instance) VolumetricFogTemporalReprojectionAmount() float64 {
	return float64(float64(class(self).GetVolumetricFogTemporalReprojectionAmount()))
}

func (self Instance) SetVolumetricFogTemporalReprojectionAmount(value float64) {
	class(self).SetVolumetricFogTemporalReprojectionAmount(gd.Float(value))
}

func (self Instance) AdjustmentEnabled() bool {
	return bool(class(self).IsAdjustmentEnabled())
}

func (self Instance) SetAdjustmentEnabled(value bool) {
	class(self).SetAdjustmentEnabled(value)
}

func (self Instance) AdjustmentBrightness() float64 {
	return float64(float64(class(self).GetAdjustmentBrightness()))
}

func (self Instance) SetAdjustmentBrightness(value float64) {
	class(self).SetAdjustmentBrightness(gd.Float(value))
}

func (self Instance) AdjustmentContrast() float64 {
	return float64(float64(class(self).GetAdjustmentContrast()))
}

func (self Instance) SetAdjustmentContrast(value float64) {
	class(self).SetAdjustmentContrast(gd.Float(value))
}

func (self Instance) AdjustmentSaturation() float64 {
	return float64(float64(class(self).GetAdjustmentSaturation()))
}

func (self Instance) SetAdjustmentSaturation(value float64) {
	class(self).SetAdjustmentSaturation(gd.Float(value))
}

func (self Instance) AdjustmentColorCorrection() objects.Texture {
	return objects.Texture(class(self).GetAdjustmentColorCorrection())
}

func (self Instance) SetAdjustmentColorCorrection(value objects.Texture) {
	class(self).SetAdjustmentColorCorrection(value)
}

//go:nosplit
func (self class) SetBackground(mode classdb.EnvironmentBGMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_background, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBackground() classdb.EnvironmentBGMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentBGMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_background, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSky(sky objects.Sky) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(sky[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sky, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSky() objects.Sky {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sky, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Sky{classdb.Sky(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkyCustomFov(scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sky_custom_fov, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkyCustomFov() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sky_custom_fov, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkyRotation(euler_radians gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, euler_radians)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sky_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkyRotation() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sky_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBgColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBgColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBgEnergyMultiplier(energy gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_bg_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBgEnergyMultiplier() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_bg_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBgIntensity(energy gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_bg_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBgIntensity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_bg_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCanvasMaxLayer(layer gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_canvas_max_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCanvasMaxLayer() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_canvas_max_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCameraFeedId(id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_camera_feed_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCameraFeedId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_camera_feed_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAmbientLightColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ambient_light_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAmbientLightColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ambient_light_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAmbientSource(source classdb.EnvironmentAmbientSource) {
	var frame = callframe.New()
	callframe.Arg(frame, source)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ambient_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAmbientSource() classdb.EnvironmentAmbientSource {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentAmbientSource](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ambient_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAmbientLightEnergy(energy gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ambient_light_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAmbientLightEnergy() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ambient_light_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAmbientLightSkyContribution(ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ambient_light_sky_contribution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAmbientLightSkyContribution() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ambient_light_sky_contribution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetReflectionSource(source classdb.EnvironmentReflectionSource) {
	var frame = callframe.New()
	callframe.Arg(frame, source)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_reflection_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetReflectionSource() classdb.EnvironmentReflectionSource {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentReflectionSource](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_reflection_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTonemapper(mode classdb.EnvironmentToneMapper) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_tonemapper, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTonemapper() classdb.EnvironmentToneMapper {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentToneMapper](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_tonemapper, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTonemapExposure(exposure gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, exposure)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_tonemap_exposure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTonemapExposure() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_tonemap_exposure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTonemapWhite(white gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, white)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_tonemap_white, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTonemapWhite() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_tonemap_white, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsrEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssr_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSsrEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_ssr_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsrMaxSteps(max_steps gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, max_steps)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssr_max_steps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsrMaxSteps() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssr_max_steps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsrFadeIn(fade_in gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, fade_in)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssr_fade_in, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsrFadeIn() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssr_fade_in, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsrFadeOut(fade_out gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, fade_out)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssr_fade_out, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsrFadeOut() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssr_fade_out, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsrDepthTolerance(depth_tolerance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, depth_tolerance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssr_depth_tolerance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsrDepthTolerance() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssr_depth_tolerance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsaoEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssao_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSsaoEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_ssao_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsaoRadius(radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssao_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsaoRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssao_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsaoIntensity(intensity gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, intensity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssao_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsaoIntensity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssao_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsaoPower(power gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, power)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssao_power, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsaoPower() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssao_power, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsaoDetail(detail gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, detail)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssao_detail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsaoDetail() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssao_detail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsaoHorizon(horizon gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, horizon)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssao_horizon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsaoHorizon() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssao_horizon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsaoSharpness(sharpness gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, sharpness)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssao_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsaoSharpness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssao_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsaoDirectLightAffect(amount gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssao_direct_light_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsaoDirectLightAffect() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssao_direct_light_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsaoAoChannelAffect(amount gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssao_ao_channel_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsaoAoChannelAffect() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssao_ao_channel_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsilEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssil_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSsilEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_ssil_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsilRadius(radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssil_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsilRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssil_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsilIntensity(intensity gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, intensity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssil_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsilIntensity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssil_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsilSharpness(sharpness gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, sharpness)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssil_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsilSharpness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssil_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsilNormalRejection(normal_rejection gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, normal_rejection)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssil_normal_rejection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsilNormalRejection() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssil_normal_rejection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSdfgiEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_sdfgi_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiCascades(amount gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_cascades, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfgiCascades() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sdfgi_cascades, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiMinCellSize(size gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_min_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfgiMinCellSize() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sdfgi_min_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiMaxDistance(distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_max_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfgiMaxDistance() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sdfgi_max_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiCascade0Distance(distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_cascade0_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfgiCascade0Distance() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sdfgi_cascade0_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiYScale(scale classdb.EnvironmentSDFGIYScale) {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_y_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfgiYScale() classdb.EnvironmentSDFGIYScale {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentSDFGIYScale](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sdfgi_y_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiUseOcclusion(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_use_occlusion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSdfgiUsingOcclusion() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_sdfgi_using_occlusion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiBounceFeedback(amount gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_bounce_feedback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfgiBounceFeedback() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sdfgi_bounce_feedback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiReadSkyLight(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_read_sky_light, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSdfgiReadingSkyLight() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_sdfgi_reading_sky_light, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiEnergy(amount gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfgiEnergy() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sdfgi_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiNormalBias(bias gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_normal_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfgiNormalBias() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sdfgi_normal_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiProbeBias(bias gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_probe_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfgiProbeBias() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sdfgi_probe_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsGlowEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_glow_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the intensity of the glow level [param idx]. A value above [code]0.0[/code] enables the level. Each level relies on the previous level. This means that enabling higher glow levels will slow down the glow effect rendering, even if previous levels aren't enabled.
*/
//go:nosplit
func (self class) SetGlowLevel(idx gd.Int, intensity gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, intensity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the intensity of the glow level [param idx].
*/
//go:nosplit
func (self class) GetGlowLevel(idx gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowNormalized(normalize bool) {
	var frame = callframe.New()
	callframe.Arg(frame, normalize)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_normalized, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsGlowNormalized() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_glow_normalized, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowIntensity(intensity gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, intensity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowIntensity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowStrength(strength gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowStrength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowMix(mix gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, mix)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_mix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowMix() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_mix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowBloom(amount gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_bloom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowBloom() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_bloom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowBlendMode(mode classdb.EnvironmentGlowBlendMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowBlendMode() classdb.EnvironmentGlowBlendMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentGlowBlendMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowHdrBleedThreshold(threshold gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, threshold)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_hdr_bleed_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowHdrBleedThreshold() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_hdr_bleed_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowHdrBleedScale(scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_hdr_bleed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowHdrBleedScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_hdr_bleed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowHdrLuminanceCap(amount gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_hdr_luminance_cap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowHdrLuminanceCap() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_hdr_luminance_cap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowMapStrength(strength gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_map_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowMapStrength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_map_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowMap(mode objects.Texture) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mode[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowMap() objects.Texture {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture{classdb.Texture(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsFogEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_fog_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogMode(mode classdb.EnvironmentFogMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogMode() classdb.EnvironmentFogMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentFogMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogLightColor(light_color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, light_color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_light_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogLightColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_light_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogLightEnergy(light_energy gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, light_energy)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_light_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogLightEnergy() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_light_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogSunScatter(sun_scatter gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, sun_scatter)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_sun_scatter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogSunScatter() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_sun_scatter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogDensity(density gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, density)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogDensity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogHeight(height gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogHeight() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogHeightDensity(height_density gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, height_density)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_height_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogHeightDensity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_height_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogAerialPerspective(aerial_perspective gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, aerial_perspective)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_aerial_perspective, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogAerialPerspective() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_aerial_perspective, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogSkyAffect(sky_affect gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, sky_affect)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_sky_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogSkyAffect() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_sky_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogDepthCurve(curve gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, curve)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_depth_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogDepthCurve() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_depth_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogDepthBegin(begin gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, begin)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_depth_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogDepthBegin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_depth_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogDepthEnd(end gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, end)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_depth_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogDepthEnd() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_depth_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsVolumetricFogEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_volumetric_fog_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogEmission(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_emission, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogEmission() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_emission, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogAlbedo(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_albedo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogAlbedo() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_albedo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogDensity(density gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, density)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogDensity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogEmissionEnergy(begin gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, begin)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_emission_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogEmissionEnergy() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_emission_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogAnisotropy(anisotropy gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, anisotropy)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_anisotropy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogAnisotropy() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_anisotropy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogLength(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogDetailSpread(detail_spread gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, detail_spread)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_detail_spread, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogDetailSpread() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_detail_spread, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogGiInject(gi_inject gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, gi_inject)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_gi_inject, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogGiInject() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_gi_inject, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogAmbientInject(enabled gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_ambient_inject, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogAmbientInject() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_ambient_inject, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogSkyAffect(sky_affect gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, sky_affect)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_sky_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogSkyAffect() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_sky_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogTemporalReprojectionEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_temporal_reprojection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsVolumetricFogTemporalReprojectionEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_volumetric_fog_temporal_reprojection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogTemporalReprojectionAmount(temporal_reprojection_amount gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, temporal_reprojection_amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_temporal_reprojection_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogTemporalReprojectionAmount() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_temporal_reprojection_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAdjustmentEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_adjustment_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsAdjustmentEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_adjustment_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAdjustmentBrightness(brightness gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, brightness)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_adjustment_brightness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAdjustmentBrightness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_adjustment_brightness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAdjustmentContrast(contrast gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, contrast)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_adjustment_contrast, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAdjustmentContrast() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_adjustment_contrast, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAdjustmentSaturation(saturation gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, saturation)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_adjustment_saturation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAdjustmentSaturation() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_adjustment_saturation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAdjustmentColorCorrection(color_correction objects.Texture) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(color_correction[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_adjustment_color_correction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAdjustmentColorCorrection() objects.Texture {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_adjustment_color_correction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture{classdb.Texture(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsEnvironment() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEnvironment() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("Environment", func(ptr gd.Object) any { return classdb.Environment(ptr) })
}

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
