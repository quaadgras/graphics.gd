// Package Environment provides methods for working with Environment object instances.
package Environment

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
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Color"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
Resource for environment nodes (like [WorldEnvironment]) that define multiple environment operations (such as background [Sky] or [Color], ambient light, fog, depth-of-field...). These parameters affect the final render of the scene. The order of these operations is:
- Depth of Field Blur
- Glow
- Tonemap (Auto Exposure)
- Adjustments
*/
type Instance [1]gdclass.Environment

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEnvironment() Instance
}

/*
Sets the intensity of the glow level [param idx]. A value above [code]0.0[/code] enables the level. Each level relies on the previous level. This means that enabling higher glow levels will slow down the glow effect rendering, even if previous levels aren't enabled.
*/
func (self Instance) SetGlowLevel(idx int, intensity Float.X) { //gd:Environment.set_glow_level
	class(self).SetGlowLevel(gd.Int(idx), gd.Float(intensity))
}

/*
Returns the intensity of the glow level [param idx].
*/
func (self Instance) GetGlowLevel(idx int) Float.X { //gd:Environment.get_glow_level
	return Float.X(Float.X(class(self).GetGlowLevel(gd.Int(idx))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Environment

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Environment"))
	casted := Instance{*(*gdclass.Environment)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) BackgroundMode() gdclass.EnvironmentBGMode {
	return gdclass.EnvironmentBGMode(class(self).GetBackground())
}

func (self Instance) SetBackgroundMode(value gdclass.EnvironmentBGMode) {
	class(self).SetBackground(value)
}

func (self Instance) BackgroundColor() Color.RGBA {
	return Color.RGBA(class(self).GetBgColor())
}

func (self Instance) SetBackgroundColor(value Color.RGBA) {
	class(self).SetBgColor(gd.Color(value))
}

func (self Instance) BackgroundEnergyMultiplier() Float.X {
	return Float.X(Float.X(class(self).GetBgEnergyMultiplier()))
}

func (self Instance) SetBackgroundEnergyMultiplier(value Float.X) {
	class(self).SetBgEnergyMultiplier(gd.Float(value))
}

func (self Instance) BackgroundIntensity() Float.X {
	return Float.X(Float.X(class(self).GetBgIntensity()))
}

func (self Instance) SetBackgroundIntensity(value Float.X) {
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

func (self Instance) Sky() [1]gdclass.Sky {
	return [1]gdclass.Sky(class(self).GetSky())
}

func (self Instance) SetSky(value [1]gdclass.Sky) {
	class(self).SetSky(value)
}

func (self Instance) SkyCustomFov() Float.X {
	return Float.X(Float.X(class(self).GetSkyCustomFov()))
}

func (self Instance) SetSkyCustomFov(value Float.X) {
	class(self).SetSkyCustomFov(gd.Float(value))
}

func (self Instance) SkyRotation() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetSkyRotation())
}

func (self Instance) SetSkyRotation(value Vector3.XYZ) {
	class(self).SetSkyRotation(gd.Vector3(value))
}

func (self Instance) AmbientLightSource() gdclass.EnvironmentAmbientSource {
	return gdclass.EnvironmentAmbientSource(class(self).GetAmbientSource())
}

func (self Instance) SetAmbientLightSource(value gdclass.EnvironmentAmbientSource) {
	class(self).SetAmbientSource(value)
}

func (self Instance) AmbientLightColor() Color.RGBA {
	return Color.RGBA(class(self).GetAmbientLightColor())
}

func (self Instance) SetAmbientLightColor(value Color.RGBA) {
	class(self).SetAmbientLightColor(gd.Color(value))
}

func (self Instance) AmbientLightSkyContribution() Float.X {
	return Float.X(Float.X(class(self).GetAmbientLightSkyContribution()))
}

func (self Instance) SetAmbientLightSkyContribution(value Float.X) {
	class(self).SetAmbientLightSkyContribution(gd.Float(value))
}

func (self Instance) AmbientLightEnergy() Float.X {
	return Float.X(Float.X(class(self).GetAmbientLightEnergy()))
}

func (self Instance) SetAmbientLightEnergy(value Float.X) {
	class(self).SetAmbientLightEnergy(gd.Float(value))
}

func (self Instance) ReflectedLightSource() gdclass.EnvironmentReflectionSource {
	return gdclass.EnvironmentReflectionSource(class(self).GetReflectionSource())
}

func (self Instance) SetReflectedLightSource(value gdclass.EnvironmentReflectionSource) {
	class(self).SetReflectionSource(value)
}

func (self Instance) TonemapMode() gdclass.EnvironmentToneMapper {
	return gdclass.EnvironmentToneMapper(class(self).GetTonemapper())
}

func (self Instance) SetTonemapMode(value gdclass.EnvironmentToneMapper) {
	class(self).SetTonemapper(value)
}

func (self Instance) TonemapExposure() Float.X {
	return Float.X(Float.X(class(self).GetTonemapExposure()))
}

func (self Instance) SetTonemapExposure(value Float.X) {
	class(self).SetTonemapExposure(gd.Float(value))
}

func (self Instance) TonemapWhite() Float.X {
	return Float.X(Float.X(class(self).GetTonemapWhite()))
}

func (self Instance) SetTonemapWhite(value Float.X) {
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

func (self Instance) SsrFadeIn() Float.X {
	return Float.X(Float.X(class(self).GetSsrFadeIn()))
}

func (self Instance) SetSsrFadeIn(value Float.X) {
	class(self).SetSsrFadeIn(gd.Float(value))
}

func (self Instance) SsrFadeOut() Float.X {
	return Float.X(Float.X(class(self).GetSsrFadeOut()))
}

func (self Instance) SetSsrFadeOut(value Float.X) {
	class(self).SetSsrFadeOut(gd.Float(value))
}

func (self Instance) SsrDepthTolerance() Float.X {
	return Float.X(Float.X(class(self).GetSsrDepthTolerance()))
}

func (self Instance) SetSsrDepthTolerance(value Float.X) {
	class(self).SetSsrDepthTolerance(gd.Float(value))
}

func (self Instance) SsaoEnabled() bool {
	return bool(class(self).IsSsaoEnabled())
}

func (self Instance) SetSsaoEnabled(value bool) {
	class(self).SetSsaoEnabled(value)
}

func (self Instance) SsaoRadius() Float.X {
	return Float.X(Float.X(class(self).GetSsaoRadius()))
}

func (self Instance) SetSsaoRadius(value Float.X) {
	class(self).SetSsaoRadius(gd.Float(value))
}

func (self Instance) SsaoIntensity() Float.X {
	return Float.X(Float.X(class(self).GetSsaoIntensity()))
}

func (self Instance) SetSsaoIntensity(value Float.X) {
	class(self).SetSsaoIntensity(gd.Float(value))
}

func (self Instance) SsaoPower() Float.X {
	return Float.X(Float.X(class(self).GetSsaoPower()))
}

func (self Instance) SetSsaoPower(value Float.X) {
	class(self).SetSsaoPower(gd.Float(value))
}

func (self Instance) SsaoDetail() Float.X {
	return Float.X(Float.X(class(self).GetSsaoDetail()))
}

func (self Instance) SetSsaoDetail(value Float.X) {
	class(self).SetSsaoDetail(gd.Float(value))
}

func (self Instance) SsaoHorizon() Float.X {
	return Float.X(Float.X(class(self).GetSsaoHorizon()))
}

func (self Instance) SetSsaoHorizon(value Float.X) {
	class(self).SetSsaoHorizon(gd.Float(value))
}

func (self Instance) SsaoSharpness() Float.X {
	return Float.X(Float.X(class(self).GetSsaoSharpness()))
}

func (self Instance) SetSsaoSharpness(value Float.X) {
	class(self).SetSsaoSharpness(gd.Float(value))
}

func (self Instance) SsaoLightAffect() Float.X {
	return Float.X(Float.X(class(self).GetSsaoDirectLightAffect()))
}

func (self Instance) SetSsaoLightAffect(value Float.X) {
	class(self).SetSsaoDirectLightAffect(gd.Float(value))
}

func (self Instance) SsaoAoChannelAffect() Float.X {
	return Float.X(Float.X(class(self).GetSsaoAoChannelAffect()))
}

func (self Instance) SetSsaoAoChannelAffect(value Float.X) {
	class(self).SetSsaoAoChannelAffect(gd.Float(value))
}

func (self Instance) SsilEnabled() bool {
	return bool(class(self).IsSsilEnabled())
}

func (self Instance) SetSsilEnabled(value bool) {
	class(self).SetSsilEnabled(value)
}

func (self Instance) SsilRadius() Float.X {
	return Float.X(Float.X(class(self).GetSsilRadius()))
}

func (self Instance) SetSsilRadius(value Float.X) {
	class(self).SetSsilRadius(gd.Float(value))
}

func (self Instance) SsilIntensity() Float.X {
	return Float.X(Float.X(class(self).GetSsilIntensity()))
}

func (self Instance) SetSsilIntensity(value Float.X) {
	class(self).SetSsilIntensity(gd.Float(value))
}

func (self Instance) SsilSharpness() Float.X {
	return Float.X(Float.X(class(self).GetSsilSharpness()))
}

func (self Instance) SetSsilSharpness(value Float.X) {
	class(self).SetSsilSharpness(gd.Float(value))
}

func (self Instance) SsilNormalRejection() Float.X {
	return Float.X(Float.X(class(self).GetSsilNormalRejection()))
}

func (self Instance) SetSsilNormalRejection(value Float.X) {
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

func (self Instance) SdfgiBounceFeedback() Float.X {
	return Float.X(Float.X(class(self).GetSdfgiBounceFeedback()))
}

func (self Instance) SetSdfgiBounceFeedback(value Float.X) {
	class(self).SetSdfgiBounceFeedback(gd.Float(value))
}

func (self Instance) SdfgiCascades() int {
	return int(int(class(self).GetSdfgiCascades()))
}

func (self Instance) SetSdfgiCascades(value int) {
	class(self).SetSdfgiCascades(gd.Int(value))
}

func (self Instance) SdfgiMinCellSize() Float.X {
	return Float.X(Float.X(class(self).GetSdfgiMinCellSize()))
}

func (self Instance) SetSdfgiMinCellSize(value Float.X) {
	class(self).SetSdfgiMinCellSize(gd.Float(value))
}

func (self Instance) SdfgiCascade0Distance() Float.X {
	return Float.X(Float.X(class(self).GetSdfgiCascade0Distance()))
}

func (self Instance) SetSdfgiCascade0Distance(value Float.X) {
	class(self).SetSdfgiCascade0Distance(gd.Float(value))
}

func (self Instance) SdfgiMaxDistance() Float.X {
	return Float.X(Float.X(class(self).GetSdfgiMaxDistance()))
}

func (self Instance) SetSdfgiMaxDistance(value Float.X) {
	class(self).SetSdfgiMaxDistance(gd.Float(value))
}

func (self Instance) SdfgiYScale() gdclass.EnvironmentSDFGIYScale {
	return gdclass.EnvironmentSDFGIYScale(class(self).GetSdfgiYScale())
}

func (self Instance) SetSdfgiYScale(value gdclass.EnvironmentSDFGIYScale) {
	class(self).SetSdfgiYScale(value)
}

func (self Instance) SdfgiEnergy() Float.X {
	return Float.X(Float.X(class(self).GetSdfgiEnergy()))
}

func (self Instance) SetSdfgiEnergy(value Float.X) {
	class(self).SetSdfgiEnergy(gd.Float(value))
}

func (self Instance) SdfgiNormalBias() Float.X {
	return Float.X(Float.X(class(self).GetSdfgiNormalBias()))
}

func (self Instance) SetSdfgiNormalBias(value Float.X) {
	class(self).SetSdfgiNormalBias(gd.Float(value))
}

func (self Instance) SdfgiProbeBias() Float.X {
	return Float.X(Float.X(class(self).GetSdfgiProbeBias()))
}

func (self Instance) SetSdfgiProbeBias(value Float.X) {
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

func (self Instance) GlowIntensity() Float.X {
	return Float.X(Float.X(class(self).GetGlowIntensity()))
}

func (self Instance) SetGlowIntensity(value Float.X) {
	class(self).SetGlowIntensity(gd.Float(value))
}

func (self Instance) GlowStrength() Float.X {
	return Float.X(Float.X(class(self).GetGlowStrength()))
}

func (self Instance) SetGlowStrength(value Float.X) {
	class(self).SetGlowStrength(gd.Float(value))
}

func (self Instance) GlowMix() Float.X {
	return Float.X(Float.X(class(self).GetGlowMix()))
}

func (self Instance) SetGlowMix(value Float.X) {
	class(self).SetGlowMix(gd.Float(value))
}

func (self Instance) GlowBloom() Float.X {
	return Float.X(Float.X(class(self).GetGlowBloom()))
}

func (self Instance) SetGlowBloom(value Float.X) {
	class(self).SetGlowBloom(gd.Float(value))
}

func (self Instance) GlowBlendMode() gdclass.EnvironmentGlowBlendMode {
	return gdclass.EnvironmentGlowBlendMode(class(self).GetGlowBlendMode())
}

func (self Instance) SetGlowBlendMode(value gdclass.EnvironmentGlowBlendMode) {
	class(self).SetGlowBlendMode(value)
}

func (self Instance) GlowHdrThreshold() Float.X {
	return Float.X(Float.X(class(self).GetGlowHdrBleedThreshold()))
}

func (self Instance) SetGlowHdrThreshold(value Float.X) {
	class(self).SetGlowHdrBleedThreshold(gd.Float(value))
}

func (self Instance) GlowHdrScale() Float.X {
	return Float.X(Float.X(class(self).GetGlowHdrBleedScale()))
}

func (self Instance) SetGlowHdrScale(value Float.X) {
	class(self).SetGlowHdrBleedScale(gd.Float(value))
}

func (self Instance) GlowHdrLuminanceCap() Float.X {
	return Float.X(Float.X(class(self).GetGlowHdrLuminanceCap()))
}

func (self Instance) SetGlowHdrLuminanceCap(value Float.X) {
	class(self).SetGlowHdrLuminanceCap(gd.Float(value))
}

func (self Instance) GlowMapStrength() Float.X {
	return Float.X(Float.X(class(self).GetGlowMapStrength()))
}

func (self Instance) SetGlowMapStrength(value Float.X) {
	class(self).SetGlowMapStrength(gd.Float(value))
}

func (self Instance) GlowMap() [1]gdclass.Texture {
	return [1]gdclass.Texture(class(self).GetGlowMap())
}

func (self Instance) SetGlowMap(value [1]gdclass.Texture) {
	class(self).SetGlowMap(value)
}

func (self Instance) FogEnabled() bool {
	return bool(class(self).IsFogEnabled())
}

func (self Instance) SetFogEnabled(value bool) {
	class(self).SetFogEnabled(value)
}

func (self Instance) FogMode() gdclass.EnvironmentFogMode {
	return gdclass.EnvironmentFogMode(class(self).GetFogMode())
}

func (self Instance) SetFogMode(value gdclass.EnvironmentFogMode) {
	class(self).SetFogMode(value)
}

func (self Instance) FogLightColor() Color.RGBA {
	return Color.RGBA(class(self).GetFogLightColor())
}

func (self Instance) SetFogLightColor(value Color.RGBA) {
	class(self).SetFogLightColor(gd.Color(value))
}

func (self Instance) FogLightEnergy() Float.X {
	return Float.X(Float.X(class(self).GetFogLightEnergy()))
}

func (self Instance) SetFogLightEnergy(value Float.X) {
	class(self).SetFogLightEnergy(gd.Float(value))
}

func (self Instance) FogSunScatter() Float.X {
	return Float.X(Float.X(class(self).GetFogSunScatter()))
}

func (self Instance) SetFogSunScatter(value Float.X) {
	class(self).SetFogSunScatter(gd.Float(value))
}

func (self Instance) FogDensity() Float.X {
	return Float.X(Float.X(class(self).GetFogDensity()))
}

func (self Instance) SetFogDensity(value Float.X) {
	class(self).SetFogDensity(gd.Float(value))
}

func (self Instance) FogAerialPerspective() Float.X {
	return Float.X(Float.X(class(self).GetFogAerialPerspective()))
}

func (self Instance) SetFogAerialPerspective(value Float.X) {
	class(self).SetFogAerialPerspective(gd.Float(value))
}

func (self Instance) FogSkyAffect() Float.X {
	return Float.X(Float.X(class(self).GetFogSkyAffect()))
}

func (self Instance) SetFogSkyAffect(value Float.X) {
	class(self).SetFogSkyAffect(gd.Float(value))
}

func (self Instance) FogHeight() Float.X {
	return Float.X(Float.X(class(self).GetFogHeight()))
}

func (self Instance) SetFogHeight(value Float.X) {
	class(self).SetFogHeight(gd.Float(value))
}

func (self Instance) FogHeightDensity() Float.X {
	return Float.X(Float.X(class(self).GetFogHeightDensity()))
}

func (self Instance) SetFogHeightDensity(value Float.X) {
	class(self).SetFogHeightDensity(gd.Float(value))
}

func (self Instance) FogDepthCurve() Float.X {
	return Float.X(Float.X(class(self).GetFogDepthCurve()))
}

func (self Instance) SetFogDepthCurve(value Float.X) {
	class(self).SetFogDepthCurve(gd.Float(value))
}

func (self Instance) FogDepthBegin() Float.X {
	return Float.X(Float.X(class(self).GetFogDepthBegin()))
}

func (self Instance) SetFogDepthBegin(value Float.X) {
	class(self).SetFogDepthBegin(gd.Float(value))
}

func (self Instance) FogDepthEnd() Float.X {
	return Float.X(Float.X(class(self).GetFogDepthEnd()))
}

func (self Instance) SetFogDepthEnd(value Float.X) {
	class(self).SetFogDepthEnd(gd.Float(value))
}

func (self Instance) VolumetricFogEnabled() bool {
	return bool(class(self).IsVolumetricFogEnabled())
}

func (self Instance) SetVolumetricFogEnabled(value bool) {
	class(self).SetVolumetricFogEnabled(value)
}

func (self Instance) VolumetricFogDensity() Float.X {
	return Float.X(Float.X(class(self).GetVolumetricFogDensity()))
}

func (self Instance) SetVolumetricFogDensity(value Float.X) {
	class(self).SetVolumetricFogDensity(gd.Float(value))
}

func (self Instance) VolumetricFogAlbedo() Color.RGBA {
	return Color.RGBA(class(self).GetVolumetricFogAlbedo())
}

func (self Instance) SetVolumetricFogAlbedo(value Color.RGBA) {
	class(self).SetVolumetricFogAlbedo(gd.Color(value))
}

func (self Instance) VolumetricFogEmission() Color.RGBA {
	return Color.RGBA(class(self).GetVolumetricFogEmission())
}

func (self Instance) SetVolumetricFogEmission(value Color.RGBA) {
	class(self).SetVolumetricFogEmission(gd.Color(value))
}

func (self Instance) VolumetricFogEmissionEnergy() Float.X {
	return Float.X(Float.X(class(self).GetVolumetricFogEmissionEnergy()))
}

func (self Instance) SetVolumetricFogEmissionEnergy(value Float.X) {
	class(self).SetVolumetricFogEmissionEnergy(gd.Float(value))
}

func (self Instance) VolumetricFogGiInject() Float.X {
	return Float.X(Float.X(class(self).GetVolumetricFogGiInject()))
}

func (self Instance) SetVolumetricFogGiInject(value Float.X) {
	class(self).SetVolumetricFogGiInject(gd.Float(value))
}

func (self Instance) VolumetricFogAnisotropy() Float.X {
	return Float.X(Float.X(class(self).GetVolumetricFogAnisotropy()))
}

func (self Instance) SetVolumetricFogAnisotropy(value Float.X) {
	class(self).SetVolumetricFogAnisotropy(gd.Float(value))
}

func (self Instance) VolumetricFogLength() Float.X {
	return Float.X(Float.X(class(self).GetVolumetricFogLength()))
}

func (self Instance) SetVolumetricFogLength(value Float.X) {
	class(self).SetVolumetricFogLength(gd.Float(value))
}

func (self Instance) VolumetricFogDetailSpread() Float.X {
	return Float.X(Float.X(class(self).GetVolumetricFogDetailSpread()))
}

func (self Instance) SetVolumetricFogDetailSpread(value Float.X) {
	class(self).SetVolumetricFogDetailSpread(gd.Float(value))
}

func (self Instance) VolumetricFogAmbientInject() Float.X {
	return Float.X(Float.X(class(self).GetVolumetricFogAmbientInject()))
}

func (self Instance) SetVolumetricFogAmbientInject(value Float.X) {
	class(self).SetVolumetricFogAmbientInject(gd.Float(value))
}

func (self Instance) VolumetricFogSkyAffect() Float.X {
	return Float.X(Float.X(class(self).GetVolumetricFogSkyAffect()))
}

func (self Instance) SetVolumetricFogSkyAffect(value Float.X) {
	class(self).SetVolumetricFogSkyAffect(gd.Float(value))
}

func (self Instance) VolumetricFogTemporalReprojectionEnabled() bool {
	return bool(class(self).IsVolumetricFogTemporalReprojectionEnabled())
}

func (self Instance) SetVolumetricFogTemporalReprojectionEnabled(value bool) {
	class(self).SetVolumetricFogTemporalReprojectionEnabled(value)
}

func (self Instance) VolumetricFogTemporalReprojectionAmount() Float.X {
	return Float.X(Float.X(class(self).GetVolumetricFogTemporalReprojectionAmount()))
}

func (self Instance) SetVolumetricFogTemporalReprojectionAmount(value Float.X) {
	class(self).SetVolumetricFogTemporalReprojectionAmount(gd.Float(value))
}

func (self Instance) AdjustmentEnabled() bool {
	return bool(class(self).IsAdjustmentEnabled())
}

func (self Instance) SetAdjustmentEnabled(value bool) {
	class(self).SetAdjustmentEnabled(value)
}

func (self Instance) AdjustmentBrightness() Float.X {
	return Float.X(Float.X(class(self).GetAdjustmentBrightness()))
}

func (self Instance) SetAdjustmentBrightness(value Float.X) {
	class(self).SetAdjustmentBrightness(gd.Float(value))
}

func (self Instance) AdjustmentContrast() Float.X {
	return Float.X(Float.X(class(self).GetAdjustmentContrast()))
}

func (self Instance) SetAdjustmentContrast(value Float.X) {
	class(self).SetAdjustmentContrast(gd.Float(value))
}

func (self Instance) AdjustmentSaturation() Float.X {
	return Float.X(Float.X(class(self).GetAdjustmentSaturation()))
}

func (self Instance) SetAdjustmentSaturation(value Float.X) {
	class(self).SetAdjustmentSaturation(gd.Float(value))
}

func (self Instance) AdjustmentColorCorrection() [1]gdclass.Texture {
	return [1]gdclass.Texture(class(self).GetAdjustmentColorCorrection())
}

func (self Instance) SetAdjustmentColorCorrection(value [1]gdclass.Texture) {
	class(self).SetAdjustmentColorCorrection(value)
}

//go:nosplit
func (self class) SetBackground(mode gdclass.EnvironmentBGMode) { //gd:Environment.set_background
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_background, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBackground() gdclass.EnvironmentBGMode { //gd:Environment.get_background
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.EnvironmentBGMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_background, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSky(sky [1]gdclass.Sky) { //gd:Environment.set_sky
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(sky[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sky, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSky() [1]gdclass.Sky { //gd:Environment.get_sky
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sky, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Sky{gd.PointerWithOwnershipTransferredToGo[gdclass.Sky](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkyCustomFov(scale gd.Float) { //gd:Environment.set_sky_custom_fov
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sky_custom_fov, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkyCustomFov() gd.Float { //gd:Environment.get_sky_custom_fov
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sky_custom_fov, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkyRotation(euler_radians gd.Vector3) { //gd:Environment.set_sky_rotation
	var frame = callframe.New()
	callframe.Arg(frame, euler_radians)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sky_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkyRotation() gd.Vector3 { //gd:Environment.get_sky_rotation
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sky_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBgColor(color gd.Color) { //gd:Environment.set_bg_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_bg_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBgColor() gd.Color { //gd:Environment.get_bg_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_bg_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBgEnergyMultiplier(energy gd.Float) { //gd:Environment.set_bg_energy_multiplier
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_bg_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBgEnergyMultiplier() gd.Float { //gd:Environment.get_bg_energy_multiplier
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_bg_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBgIntensity(energy gd.Float) { //gd:Environment.set_bg_intensity
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_bg_intensity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBgIntensity() gd.Float { //gd:Environment.get_bg_intensity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_bg_intensity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCanvasMaxLayer(layer gd.Int) { //gd:Environment.set_canvas_max_layer
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_canvas_max_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCanvasMaxLayer() gd.Int { //gd:Environment.get_canvas_max_layer
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_canvas_max_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCameraFeedId(id gd.Int) { //gd:Environment.set_camera_feed_id
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_camera_feed_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCameraFeedId() gd.Int { //gd:Environment.get_camera_feed_id
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_camera_feed_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAmbientLightColor(color gd.Color) { //gd:Environment.set_ambient_light_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ambient_light_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAmbientLightColor() gd.Color { //gd:Environment.get_ambient_light_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ambient_light_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAmbientSource(source gdclass.EnvironmentAmbientSource) { //gd:Environment.set_ambient_source
	var frame = callframe.New()
	callframe.Arg(frame, source)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ambient_source, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAmbientSource() gdclass.EnvironmentAmbientSource { //gd:Environment.get_ambient_source
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.EnvironmentAmbientSource](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ambient_source, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAmbientLightEnergy(energy gd.Float) { //gd:Environment.set_ambient_light_energy
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ambient_light_energy, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAmbientLightEnergy() gd.Float { //gd:Environment.get_ambient_light_energy
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ambient_light_energy, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAmbientLightSkyContribution(ratio gd.Float) { //gd:Environment.set_ambient_light_sky_contribution
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ambient_light_sky_contribution, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAmbientLightSkyContribution() gd.Float { //gd:Environment.get_ambient_light_sky_contribution
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ambient_light_sky_contribution, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetReflectionSource(source gdclass.EnvironmentReflectionSource) { //gd:Environment.set_reflection_source
	var frame = callframe.New()
	callframe.Arg(frame, source)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_reflection_source, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetReflectionSource() gdclass.EnvironmentReflectionSource { //gd:Environment.get_reflection_source
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.EnvironmentReflectionSource](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_reflection_source, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTonemapper(mode gdclass.EnvironmentToneMapper) { //gd:Environment.set_tonemapper
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_tonemapper, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTonemapper() gdclass.EnvironmentToneMapper { //gd:Environment.get_tonemapper
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.EnvironmentToneMapper](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_tonemapper, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTonemapExposure(exposure gd.Float) { //gd:Environment.set_tonemap_exposure
	var frame = callframe.New()
	callframe.Arg(frame, exposure)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_tonemap_exposure, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTonemapExposure() gd.Float { //gd:Environment.get_tonemap_exposure
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_tonemap_exposure, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTonemapWhite(white gd.Float) { //gd:Environment.set_tonemap_white
	var frame = callframe.New()
	callframe.Arg(frame, white)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_tonemap_white, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTonemapWhite() gd.Float { //gd:Environment.get_tonemap_white
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_tonemap_white, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsrEnabled(enabled bool) { //gd:Environment.set_ssr_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssr_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSsrEnabled() bool { //gd:Environment.is_ssr_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_ssr_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsrMaxSteps(max_steps gd.Int) { //gd:Environment.set_ssr_max_steps
	var frame = callframe.New()
	callframe.Arg(frame, max_steps)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssr_max_steps, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsrMaxSteps() gd.Int { //gd:Environment.get_ssr_max_steps
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssr_max_steps, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsrFadeIn(fade_in gd.Float) { //gd:Environment.set_ssr_fade_in
	var frame = callframe.New()
	callframe.Arg(frame, fade_in)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssr_fade_in, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsrFadeIn() gd.Float { //gd:Environment.get_ssr_fade_in
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssr_fade_in, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsrFadeOut(fade_out gd.Float) { //gd:Environment.set_ssr_fade_out
	var frame = callframe.New()
	callframe.Arg(frame, fade_out)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssr_fade_out, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsrFadeOut() gd.Float { //gd:Environment.get_ssr_fade_out
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssr_fade_out, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsrDepthTolerance(depth_tolerance gd.Float) { //gd:Environment.set_ssr_depth_tolerance
	var frame = callframe.New()
	callframe.Arg(frame, depth_tolerance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssr_depth_tolerance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsrDepthTolerance() gd.Float { //gd:Environment.get_ssr_depth_tolerance
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssr_depth_tolerance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsaoEnabled(enabled bool) { //gd:Environment.set_ssao_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssao_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSsaoEnabled() bool { //gd:Environment.is_ssao_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_ssao_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsaoRadius(radius gd.Float) { //gd:Environment.set_ssao_radius
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssao_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsaoRadius() gd.Float { //gd:Environment.get_ssao_radius
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssao_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsaoIntensity(intensity gd.Float) { //gd:Environment.set_ssao_intensity
	var frame = callframe.New()
	callframe.Arg(frame, intensity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssao_intensity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsaoIntensity() gd.Float { //gd:Environment.get_ssao_intensity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssao_intensity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsaoPower(power gd.Float) { //gd:Environment.set_ssao_power
	var frame = callframe.New()
	callframe.Arg(frame, power)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssao_power, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsaoPower() gd.Float { //gd:Environment.get_ssao_power
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssao_power, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsaoDetail(detail gd.Float) { //gd:Environment.set_ssao_detail
	var frame = callframe.New()
	callframe.Arg(frame, detail)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssao_detail, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsaoDetail() gd.Float { //gd:Environment.get_ssao_detail
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssao_detail, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsaoHorizon(horizon gd.Float) { //gd:Environment.set_ssao_horizon
	var frame = callframe.New()
	callframe.Arg(frame, horizon)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssao_horizon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsaoHorizon() gd.Float { //gd:Environment.get_ssao_horizon
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssao_horizon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsaoSharpness(sharpness gd.Float) { //gd:Environment.set_ssao_sharpness
	var frame = callframe.New()
	callframe.Arg(frame, sharpness)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssao_sharpness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsaoSharpness() gd.Float { //gd:Environment.get_ssao_sharpness
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssao_sharpness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsaoDirectLightAffect(amount gd.Float) { //gd:Environment.set_ssao_direct_light_affect
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssao_direct_light_affect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsaoDirectLightAffect() gd.Float { //gd:Environment.get_ssao_direct_light_affect
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssao_direct_light_affect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsaoAoChannelAffect(amount gd.Float) { //gd:Environment.set_ssao_ao_channel_affect
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssao_ao_channel_affect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsaoAoChannelAffect() gd.Float { //gd:Environment.get_ssao_ao_channel_affect
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssao_ao_channel_affect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsilEnabled(enabled bool) { //gd:Environment.set_ssil_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssil_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSsilEnabled() bool { //gd:Environment.is_ssil_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_ssil_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsilRadius(radius gd.Float) { //gd:Environment.set_ssil_radius
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssil_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsilRadius() gd.Float { //gd:Environment.get_ssil_radius
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssil_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsilIntensity(intensity gd.Float) { //gd:Environment.set_ssil_intensity
	var frame = callframe.New()
	callframe.Arg(frame, intensity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssil_intensity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsilIntensity() gd.Float { //gd:Environment.get_ssil_intensity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssil_intensity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsilSharpness(sharpness gd.Float) { //gd:Environment.set_ssil_sharpness
	var frame = callframe.New()
	callframe.Arg(frame, sharpness)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssil_sharpness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsilSharpness() gd.Float { //gd:Environment.get_ssil_sharpness
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssil_sharpness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSsilNormalRejection(normal_rejection gd.Float) { //gd:Environment.set_ssil_normal_rejection
	var frame = callframe.New()
	callframe.Arg(frame, normal_rejection)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_ssil_normal_rejection, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSsilNormalRejection() gd.Float { //gd:Environment.get_ssil_normal_rejection
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_ssil_normal_rejection, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiEnabled(enabled bool) { //gd:Environment.set_sdfgi_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSdfgiEnabled() bool { //gd:Environment.is_sdfgi_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_sdfgi_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiCascades(amount gd.Int) { //gd:Environment.set_sdfgi_cascades
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_cascades, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfgiCascades() gd.Int { //gd:Environment.get_sdfgi_cascades
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sdfgi_cascades, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiMinCellSize(size gd.Float) { //gd:Environment.set_sdfgi_min_cell_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_min_cell_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfgiMinCellSize() gd.Float { //gd:Environment.get_sdfgi_min_cell_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sdfgi_min_cell_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiMaxDistance(distance gd.Float) { //gd:Environment.set_sdfgi_max_distance
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_max_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfgiMaxDistance() gd.Float { //gd:Environment.get_sdfgi_max_distance
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sdfgi_max_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiCascade0Distance(distance gd.Float) { //gd:Environment.set_sdfgi_cascade0_distance
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_cascade0_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfgiCascade0Distance() gd.Float { //gd:Environment.get_sdfgi_cascade0_distance
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sdfgi_cascade0_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiYScale(scale gdclass.EnvironmentSDFGIYScale) { //gd:Environment.set_sdfgi_y_scale
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_y_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfgiYScale() gdclass.EnvironmentSDFGIYScale { //gd:Environment.get_sdfgi_y_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.EnvironmentSDFGIYScale](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sdfgi_y_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiUseOcclusion(enable bool) { //gd:Environment.set_sdfgi_use_occlusion
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_use_occlusion, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSdfgiUsingOcclusion() bool { //gd:Environment.is_sdfgi_using_occlusion
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_sdfgi_using_occlusion, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiBounceFeedback(amount gd.Float) { //gd:Environment.set_sdfgi_bounce_feedback
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_bounce_feedback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfgiBounceFeedback() gd.Float { //gd:Environment.get_sdfgi_bounce_feedback
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sdfgi_bounce_feedback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiReadSkyLight(enable bool) { //gd:Environment.set_sdfgi_read_sky_light
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_read_sky_light, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSdfgiReadingSkyLight() bool { //gd:Environment.is_sdfgi_reading_sky_light
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_sdfgi_reading_sky_light, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiEnergy(amount gd.Float) { //gd:Environment.set_sdfgi_energy
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_energy, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfgiEnergy() gd.Float { //gd:Environment.get_sdfgi_energy
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sdfgi_energy, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiNormalBias(bias gd.Float) { //gd:Environment.set_sdfgi_normal_bias
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_normal_bias, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfgiNormalBias() gd.Float { //gd:Environment.get_sdfgi_normal_bias
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sdfgi_normal_bias, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfgiProbeBias(bias gd.Float) { //gd:Environment.set_sdfgi_probe_bias
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_sdfgi_probe_bias, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfgiProbeBias() gd.Float { //gd:Environment.get_sdfgi_probe_bias
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_sdfgi_probe_bias, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowEnabled(enabled bool) { //gd:Environment.set_glow_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsGlowEnabled() bool { //gd:Environment.is_glow_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_glow_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the intensity of the glow level [param idx]. A value above [code]0.0[/code] enables the level. Each level relies on the previous level. This means that enabling higher glow levels will slow down the glow effect rendering, even if previous levels aren't enabled.
*/
//go:nosplit
func (self class) SetGlowLevel(idx gd.Int, intensity gd.Float) { //gd:Environment.set_glow_level
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, intensity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_level, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the intensity of the glow level [param idx].
*/
//go:nosplit
func (self class) GetGlowLevel(idx gd.Int) gd.Float { //gd:Environment.get_glow_level
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_level, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowNormalized(normalize bool) { //gd:Environment.set_glow_normalized
	var frame = callframe.New()
	callframe.Arg(frame, normalize)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_normalized, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsGlowNormalized() bool { //gd:Environment.is_glow_normalized
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_glow_normalized, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowIntensity(intensity gd.Float) { //gd:Environment.set_glow_intensity
	var frame = callframe.New()
	callframe.Arg(frame, intensity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_intensity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowIntensity() gd.Float { //gd:Environment.get_glow_intensity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_intensity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowStrength(strength gd.Float) { //gd:Environment.set_glow_strength
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowStrength() gd.Float { //gd:Environment.get_glow_strength
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowMix(mix gd.Float) { //gd:Environment.set_glow_mix
	var frame = callframe.New()
	callframe.Arg(frame, mix)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_mix, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowMix() gd.Float { //gd:Environment.get_glow_mix
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_mix, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowBloom(amount gd.Float) { //gd:Environment.set_glow_bloom
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_bloom, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowBloom() gd.Float { //gd:Environment.get_glow_bloom
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_bloom, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowBlendMode(mode gdclass.EnvironmentGlowBlendMode) { //gd:Environment.set_glow_blend_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_blend_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowBlendMode() gdclass.EnvironmentGlowBlendMode { //gd:Environment.get_glow_blend_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.EnvironmentGlowBlendMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_blend_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowHdrBleedThreshold(threshold gd.Float) { //gd:Environment.set_glow_hdr_bleed_threshold
	var frame = callframe.New()
	callframe.Arg(frame, threshold)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_hdr_bleed_threshold, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowHdrBleedThreshold() gd.Float { //gd:Environment.get_glow_hdr_bleed_threshold
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_hdr_bleed_threshold, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowHdrBleedScale(scale gd.Float) { //gd:Environment.set_glow_hdr_bleed_scale
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_hdr_bleed_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowHdrBleedScale() gd.Float { //gd:Environment.get_glow_hdr_bleed_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_hdr_bleed_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowHdrLuminanceCap(amount gd.Float) { //gd:Environment.set_glow_hdr_luminance_cap
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_hdr_luminance_cap, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowHdrLuminanceCap() gd.Float { //gd:Environment.get_glow_hdr_luminance_cap
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_hdr_luminance_cap, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowMapStrength(strength gd.Float) { //gd:Environment.set_glow_map_strength
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_map_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowMapStrength() gd.Float { //gd:Environment.get_glow_map_strength
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_map_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlowMap(mode [1]gdclass.Texture) { //gd:Environment.set_glow_map
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mode[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_glow_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlowMap() [1]gdclass.Texture { //gd:Environment.get_glow_map
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_glow_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogEnabled(enabled bool) { //gd:Environment.set_fog_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFogEnabled() bool { //gd:Environment.is_fog_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_fog_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogMode(mode gdclass.EnvironmentFogMode) { //gd:Environment.set_fog_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogMode() gdclass.EnvironmentFogMode { //gd:Environment.get_fog_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.EnvironmentFogMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogLightColor(light_color gd.Color) { //gd:Environment.set_fog_light_color
	var frame = callframe.New()
	callframe.Arg(frame, light_color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_light_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogLightColor() gd.Color { //gd:Environment.get_fog_light_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_light_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogLightEnergy(light_energy gd.Float) { //gd:Environment.set_fog_light_energy
	var frame = callframe.New()
	callframe.Arg(frame, light_energy)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_light_energy, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogLightEnergy() gd.Float { //gd:Environment.get_fog_light_energy
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_light_energy, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogSunScatter(sun_scatter gd.Float) { //gd:Environment.set_fog_sun_scatter
	var frame = callframe.New()
	callframe.Arg(frame, sun_scatter)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_sun_scatter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogSunScatter() gd.Float { //gd:Environment.get_fog_sun_scatter
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_sun_scatter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogDensity(density gd.Float) { //gd:Environment.set_fog_density
	var frame = callframe.New()
	callframe.Arg(frame, density)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_density, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogDensity() gd.Float { //gd:Environment.get_fog_density
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_density, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogHeight(height gd.Float) { //gd:Environment.set_fog_height
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogHeight() gd.Float { //gd:Environment.get_fog_height
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogHeightDensity(height_density gd.Float) { //gd:Environment.set_fog_height_density
	var frame = callframe.New()
	callframe.Arg(frame, height_density)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_height_density, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogHeightDensity() gd.Float { //gd:Environment.get_fog_height_density
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_height_density, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogAerialPerspective(aerial_perspective gd.Float) { //gd:Environment.set_fog_aerial_perspective
	var frame = callframe.New()
	callframe.Arg(frame, aerial_perspective)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_aerial_perspective, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogAerialPerspective() gd.Float { //gd:Environment.get_fog_aerial_perspective
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_aerial_perspective, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogSkyAffect(sky_affect gd.Float) { //gd:Environment.set_fog_sky_affect
	var frame = callframe.New()
	callframe.Arg(frame, sky_affect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_sky_affect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogSkyAffect() gd.Float { //gd:Environment.get_fog_sky_affect
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_sky_affect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogDepthCurve(curve gd.Float) { //gd:Environment.set_fog_depth_curve
	var frame = callframe.New()
	callframe.Arg(frame, curve)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_depth_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogDepthCurve() gd.Float { //gd:Environment.get_fog_depth_curve
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_depth_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogDepthBegin(begin gd.Float) { //gd:Environment.set_fog_depth_begin
	var frame = callframe.New()
	callframe.Arg(frame, begin)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_depth_begin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogDepthBegin() gd.Float { //gd:Environment.get_fog_depth_begin
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_depth_begin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFogDepthEnd(end gd.Float) { //gd:Environment.set_fog_depth_end
	var frame = callframe.New()
	callframe.Arg(frame, end)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_fog_depth_end, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFogDepthEnd() gd.Float { //gd:Environment.get_fog_depth_end
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_fog_depth_end, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogEnabled(enabled bool) { //gd:Environment.set_volumetric_fog_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsVolumetricFogEnabled() bool { //gd:Environment.is_volumetric_fog_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_volumetric_fog_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogEmission(color gd.Color) { //gd:Environment.set_volumetric_fog_emission
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_emission, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogEmission() gd.Color { //gd:Environment.get_volumetric_fog_emission
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_emission, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogAlbedo(color gd.Color) { //gd:Environment.set_volumetric_fog_albedo
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_albedo, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogAlbedo() gd.Color { //gd:Environment.get_volumetric_fog_albedo
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_albedo, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogDensity(density gd.Float) { //gd:Environment.set_volumetric_fog_density
	var frame = callframe.New()
	callframe.Arg(frame, density)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_density, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogDensity() gd.Float { //gd:Environment.get_volumetric_fog_density
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_density, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogEmissionEnergy(begin gd.Float) { //gd:Environment.set_volumetric_fog_emission_energy
	var frame = callframe.New()
	callframe.Arg(frame, begin)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_emission_energy, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogEmissionEnergy() gd.Float { //gd:Environment.get_volumetric_fog_emission_energy
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_emission_energy, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogAnisotropy(anisotropy gd.Float) { //gd:Environment.set_volumetric_fog_anisotropy
	var frame = callframe.New()
	callframe.Arg(frame, anisotropy)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_anisotropy, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogAnisotropy() gd.Float { //gd:Environment.get_volumetric_fog_anisotropy
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_anisotropy, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogLength(length gd.Float) { //gd:Environment.set_volumetric_fog_length
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogLength() gd.Float { //gd:Environment.get_volumetric_fog_length
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogDetailSpread(detail_spread gd.Float) { //gd:Environment.set_volumetric_fog_detail_spread
	var frame = callframe.New()
	callframe.Arg(frame, detail_spread)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_detail_spread, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogDetailSpread() gd.Float { //gd:Environment.get_volumetric_fog_detail_spread
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_detail_spread, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogGiInject(gi_inject gd.Float) { //gd:Environment.set_volumetric_fog_gi_inject
	var frame = callframe.New()
	callframe.Arg(frame, gi_inject)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_gi_inject, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogGiInject() gd.Float { //gd:Environment.get_volumetric_fog_gi_inject
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_gi_inject, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogAmbientInject(enabled gd.Float) { //gd:Environment.set_volumetric_fog_ambient_inject
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_ambient_inject, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogAmbientInject() gd.Float { //gd:Environment.get_volumetric_fog_ambient_inject
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_ambient_inject, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogSkyAffect(sky_affect gd.Float) { //gd:Environment.set_volumetric_fog_sky_affect
	var frame = callframe.New()
	callframe.Arg(frame, sky_affect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_sky_affect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogSkyAffect() gd.Float { //gd:Environment.get_volumetric_fog_sky_affect
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_sky_affect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogTemporalReprojectionEnabled(enabled bool) { //gd:Environment.set_volumetric_fog_temporal_reprojection_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_temporal_reprojection_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsVolumetricFogTemporalReprojectionEnabled() bool { //gd:Environment.is_volumetric_fog_temporal_reprojection_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_volumetric_fog_temporal_reprojection_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumetricFogTemporalReprojectionAmount(temporal_reprojection_amount gd.Float) { //gd:Environment.set_volumetric_fog_temporal_reprojection_amount
	var frame = callframe.New()
	callframe.Arg(frame, temporal_reprojection_amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_volumetric_fog_temporal_reprojection_amount, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumetricFogTemporalReprojectionAmount() gd.Float { //gd:Environment.get_volumetric_fog_temporal_reprojection_amount
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_volumetric_fog_temporal_reprojection_amount, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAdjustmentEnabled(enabled bool) { //gd:Environment.set_adjustment_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_adjustment_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAdjustmentEnabled() bool { //gd:Environment.is_adjustment_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_is_adjustment_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAdjustmentBrightness(brightness gd.Float) { //gd:Environment.set_adjustment_brightness
	var frame = callframe.New()
	callframe.Arg(frame, brightness)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_adjustment_brightness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAdjustmentBrightness() gd.Float { //gd:Environment.get_adjustment_brightness
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_adjustment_brightness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAdjustmentContrast(contrast gd.Float) { //gd:Environment.set_adjustment_contrast
	var frame = callframe.New()
	callframe.Arg(frame, contrast)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_adjustment_contrast, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAdjustmentContrast() gd.Float { //gd:Environment.get_adjustment_contrast
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_adjustment_contrast, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAdjustmentSaturation(saturation gd.Float) { //gd:Environment.set_adjustment_saturation
	var frame = callframe.New()
	callframe.Arg(frame, saturation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_adjustment_saturation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAdjustmentSaturation() gd.Float { //gd:Environment.get_adjustment_saturation
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_adjustment_saturation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAdjustmentColorCorrection(color_correction [1]gdclass.Texture) { //gd:Environment.set_adjustment_color_correction
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(color_correction[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_set_adjustment_color_correction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAdjustmentColorCorrection() [1]gdclass.Texture { //gd:Environment.get_adjustment_color_correction
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Environment.Bind_get_adjustment_color_correction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture](r_ret.Get())}
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
	gdclass.Register("Environment", func(ptr gd.Object) any { return [1]gdclass.Environment{*(*gdclass.Environment)(unsafe.Pointer(&ptr))} })
}

type BGMode = gdclass.EnvironmentBGMode //gd:Environment.BGMode

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

type AmbientSource = gdclass.EnvironmentAmbientSource //gd:Environment.AmbientSource

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

type ReflectionSource = gdclass.EnvironmentReflectionSource //gd:Environment.ReflectionSource

const (
	/*Use the background for reflections.*/
	ReflectionSourceBg ReflectionSource = 0
	/*Disable reflections. This provides a slight performance boost over other options.*/
	ReflectionSourceDisabled ReflectionSource = 1
	/*Use the [Sky] for reflections regardless of what the background is.*/
	ReflectionSourceSky ReflectionSource = 2
)

type ToneMapper = gdclass.EnvironmentToneMapper //gd:Environment.ToneMapper

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

type GlowBlendMode = gdclass.EnvironmentGlowBlendMode //gd:Environment.GlowBlendMode

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

type FogMode = gdclass.EnvironmentFogMode //gd:Environment.FogMode

const (
	/*Use a physically-based fog model defined primarily by fog density.*/
	FogModeExponential FogMode = 0
	/*Use a simple fog model defined by start and end positions and a custom curve. While not physically accurate, this model can be useful when you need more artistic control.*/
	FogModeDepth FogMode = 1
)

type SDFGIYScale = gdclass.EnvironmentSDFGIYScale //gd:Environment.SDFGIYScale

const (
	/*Use 50% scale for SDFGI on the Y (vertical) axis. SDFGI cells will be twice as short as they are wide. This allows providing increased GI detail and reduced light leaking with thin floors and ceilings. This is usually the best choice for scenes that don't feature much verticality.*/
	SdfgiYScale50Percent SDFGIYScale = 0
	/*Use 75% scale for SDFGI on the Y (vertical) axis. This is a balance between the 50% and 100% SDFGI Y scales.*/
	SdfgiYScale75Percent SDFGIYScale = 1
	/*Use 100% scale for SDFGI on the Y (vertical) axis. SDFGI cells will be as tall as they are wide. This is usually the best choice for highly vertical scenes. The downside is that light leaking may become more noticeable with thin floors and ceilings.*/
	SdfgiYScale100Percent SDFGIYScale = 2
)
