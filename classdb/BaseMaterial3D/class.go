// Package BaseMaterial3D provides methods for working with BaseMaterial3D object instances.
package BaseMaterial3D

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Material"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector3"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
This class serves as a default material with a wide variety of rendering features and properties without the need to write shader code. See the tutorial below for details.
*/
type Instance [1]gdclass.BaseMaterial3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsBaseMaterial3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.BaseMaterial3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("BaseMaterial3D"))
	casted := Instance{*(*gdclass.BaseMaterial3D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Transparency() gdclass.BaseMaterial3DTransparency {
	return gdclass.BaseMaterial3DTransparency(class(self).GetTransparency())
}

func (self Instance) SetTransparency(value gdclass.BaseMaterial3DTransparency) {
	class(self).SetTransparency(value)
}

func (self Instance) AlphaScissorThreshold() Float.X {
	return Float.X(Float.X(class(self).GetAlphaScissorThreshold()))
}

func (self Instance) SetAlphaScissorThreshold(value Float.X) {
	class(self).SetAlphaScissorThreshold(float64(value))
}

func (self Instance) AlphaHashScale() Float.X {
	return Float.X(Float.X(class(self).GetAlphaHashScale()))
}

func (self Instance) SetAlphaHashScale(value Float.X) {
	class(self).SetAlphaHashScale(float64(value))
}

func (self Instance) AlphaAntialiasingMode() gdclass.BaseMaterial3DAlphaAntiAliasing {
	return gdclass.BaseMaterial3DAlphaAntiAliasing(class(self).GetAlphaAntialiasing())
}

func (self Instance) SetAlphaAntialiasingMode(value gdclass.BaseMaterial3DAlphaAntiAliasing) {
	class(self).SetAlphaAntialiasing(value)
}

func (self Instance) AlphaAntialiasingEdge() Float.X {
	return Float.X(Float.X(class(self).GetAlphaAntialiasingEdge()))
}

func (self Instance) SetAlphaAntialiasingEdge(value Float.X) {
	class(self).SetAlphaAntialiasingEdge(float64(value))
}

func (self Instance) BlendMode() gdclass.BaseMaterial3DBlendMode {
	return gdclass.BaseMaterial3DBlendMode(class(self).GetBlendMode())
}

func (self Instance) SetBlendMode(value gdclass.BaseMaterial3DBlendMode) {
	class(self).SetBlendMode(value)
}

func (self Instance) CullMode() gdclass.BaseMaterial3DCullMode {
	return gdclass.BaseMaterial3DCullMode(class(self).GetCullMode())
}

func (self Instance) SetCullMode(value gdclass.BaseMaterial3DCullMode) {
	class(self).SetCullMode(value)
}

func (self Instance) DepthDrawMode() gdclass.BaseMaterial3DDepthDrawMode {
	return gdclass.BaseMaterial3DDepthDrawMode(class(self).GetDepthDrawMode())
}

func (self Instance) SetDepthDrawMode(value gdclass.BaseMaterial3DDepthDrawMode) {
	class(self).SetDepthDrawMode(value)
}

func (self Instance) NoDepthTest() bool {
	return bool(class(self).GetFlag(0))
}

func (self Instance) SetNoDepthTest(value bool) {
	class(self).SetFlag(0, value)
}

func (self Instance) ShadingMode() gdclass.BaseMaterial3DShadingMode {
	return gdclass.BaseMaterial3DShadingMode(class(self).GetShadingMode())
}

func (self Instance) SetShadingMode(value gdclass.BaseMaterial3DShadingMode) {
	class(self).SetShadingMode(value)
}

func (self Instance) DiffuseMode() gdclass.BaseMaterial3DDiffuseMode {
	return gdclass.BaseMaterial3DDiffuseMode(class(self).GetDiffuseMode())
}

func (self Instance) SetDiffuseMode(value gdclass.BaseMaterial3DDiffuseMode) {
	class(self).SetDiffuseMode(value)
}

func (self Instance) SpecularMode() gdclass.BaseMaterial3DSpecularMode {
	return gdclass.BaseMaterial3DSpecularMode(class(self).GetSpecularMode())
}

func (self Instance) SetSpecularMode(value gdclass.BaseMaterial3DSpecularMode) {
	class(self).SetSpecularMode(value)
}

func (self Instance) DisableAmbientLight() bool {
	return bool(class(self).GetFlag(14))
}

func (self Instance) SetDisableAmbientLight(value bool) {
	class(self).SetFlag(14, value)
}

func (self Instance) DisableFog() bool {
	return bool(class(self).GetFlag(21))
}

func (self Instance) SetDisableFog(value bool) {
	class(self).SetFlag(21, value)
}

func (self Instance) VertexColorUseAsAlbedo() bool {
	return bool(class(self).GetFlag(1))
}

func (self Instance) SetVertexColorUseAsAlbedo(value bool) {
	class(self).SetFlag(1, value)
}

func (self Instance) VertexColorIsSrgb() bool {
	return bool(class(self).GetFlag(2))
}

func (self Instance) SetVertexColorIsSrgb(value bool) {
	class(self).SetFlag(2, value)
}

func (self Instance) AlbedoColor() Color.RGBA {
	return Color.RGBA(class(self).GetAlbedo())
}

func (self Instance) SetAlbedoColor(value Color.RGBA) {
	class(self).SetAlbedo(Color.RGBA(value))
}

func (self Instance) AlbedoTexture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(0))
}

func (self Instance) SetAlbedoTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(0, value)
}

func (self Instance) AlbedoTextureForceSrgb() bool {
	return bool(class(self).GetFlag(12))
}

func (self Instance) SetAlbedoTextureForceSrgb(value bool) {
	class(self).SetFlag(12, value)
}

func (self Instance) AlbedoTextureMsdf() bool {
	return bool(class(self).GetFlag(20))
}

func (self Instance) SetAlbedoTextureMsdf(value bool) {
	class(self).SetFlag(20, value)
}

func (self Instance) OrmTexture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(17))
}

func (self Instance) SetOrmTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(17, value)
}

func (self Instance) Metallic() Float.X {
	return Float.X(Float.X(class(self).GetMetallic()))
}

func (self Instance) SetMetallic(value Float.X) {
	class(self).SetMetallic(float64(value))
}

func (self Instance) MetallicSpecular() Float.X {
	return Float.X(Float.X(class(self).GetSpecular()))
}

func (self Instance) SetMetallicSpecular(value Float.X) {
	class(self).SetSpecular(float64(value))
}

func (self Instance) MetallicTexture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(1))
}

func (self Instance) SetMetallicTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(1, value)
}

func (self Instance) MetallicTextureChannel() gdclass.BaseMaterial3DTextureChannel {
	return gdclass.BaseMaterial3DTextureChannel(class(self).GetMetallicTextureChannel())
}

func (self Instance) SetMetallicTextureChannel(value gdclass.BaseMaterial3DTextureChannel) {
	class(self).SetMetallicTextureChannel(value)
}

func (self Instance) Roughness() Float.X {
	return Float.X(Float.X(class(self).GetRoughness()))
}

func (self Instance) SetRoughness(value Float.X) {
	class(self).SetRoughness(float64(value))
}

func (self Instance) RoughnessTexture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(2))
}

func (self Instance) SetRoughnessTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(2, value)
}

func (self Instance) RoughnessTextureChannel() gdclass.BaseMaterial3DTextureChannel {
	return gdclass.BaseMaterial3DTextureChannel(class(self).GetRoughnessTextureChannel())
}

func (self Instance) SetRoughnessTextureChannel(value gdclass.BaseMaterial3DTextureChannel) {
	class(self).SetRoughnessTextureChannel(value)
}

func (self Instance) EmissionEnabled() bool {
	return bool(class(self).GetFeature(0))
}

func (self Instance) SetEmissionEnabled(value bool) {
	class(self).SetFeature(0, value)
}

func (self Instance) Emission() Color.RGBA {
	return Color.RGBA(class(self).GetEmission())
}

func (self Instance) SetEmission(value Color.RGBA) {
	class(self).SetEmission(Color.RGBA(value))
}

func (self Instance) EmissionEnergyMultiplier() Float.X {
	return Float.X(Float.X(class(self).GetEmissionEnergyMultiplier()))
}

func (self Instance) SetEmissionEnergyMultiplier(value Float.X) {
	class(self).SetEmissionEnergyMultiplier(float64(value))
}

func (self Instance) EmissionIntensity() Float.X {
	return Float.X(Float.X(class(self).GetEmissionIntensity()))
}

func (self Instance) SetEmissionIntensity(value Float.X) {
	class(self).SetEmissionIntensity(float64(value))
}

func (self Instance) EmissionOperator() gdclass.BaseMaterial3DEmissionOperator {
	return gdclass.BaseMaterial3DEmissionOperator(class(self).GetEmissionOperator())
}

func (self Instance) SetEmissionOperator(value gdclass.BaseMaterial3DEmissionOperator) {
	class(self).SetEmissionOperator(value)
}

func (self Instance) EmissionOnUv2() bool {
	return bool(class(self).GetFlag(11))
}

func (self Instance) SetEmissionOnUv2(value bool) {
	class(self).SetFlag(11, value)
}

func (self Instance) EmissionTexture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(3))
}

func (self Instance) SetEmissionTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(3, value)
}

func (self Instance) NormalEnabled() bool {
	return bool(class(self).GetFeature(1))
}

func (self Instance) SetNormalEnabled(value bool) {
	class(self).SetFeature(1, value)
}

func (self Instance) NormalScale() Float.X {
	return Float.X(Float.X(class(self).GetNormalScale()))
}

func (self Instance) SetNormalScale(value Float.X) {
	class(self).SetNormalScale(float64(value))
}

func (self Instance) NormalTexture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(4))
}

func (self Instance) SetNormalTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(4, value)
}

func (self Instance) RimEnabled() bool {
	return bool(class(self).GetFeature(2))
}

func (self Instance) SetRimEnabled(value bool) {
	class(self).SetFeature(2, value)
}

func (self Instance) Rim() Float.X {
	return Float.X(Float.X(class(self).GetRim()))
}

func (self Instance) SetRim(value Float.X) {
	class(self).SetRim(float64(value))
}

func (self Instance) RimTint() Float.X {
	return Float.X(Float.X(class(self).GetRimTint()))
}

func (self Instance) SetRimTint(value Float.X) {
	class(self).SetRimTint(float64(value))
}

func (self Instance) RimTexture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(5))
}

func (self Instance) SetRimTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(5, value)
}

func (self Instance) ClearcoatEnabled() bool {
	return bool(class(self).GetFeature(3))
}

func (self Instance) SetClearcoatEnabled(value bool) {
	class(self).SetFeature(3, value)
}

func (self Instance) Clearcoat() Float.X {
	return Float.X(Float.X(class(self).GetClearcoat()))
}

func (self Instance) SetClearcoat(value Float.X) {
	class(self).SetClearcoat(float64(value))
}

func (self Instance) ClearcoatRoughness() Float.X {
	return Float.X(Float.X(class(self).GetClearcoatRoughness()))
}

func (self Instance) SetClearcoatRoughness(value Float.X) {
	class(self).SetClearcoatRoughness(float64(value))
}

func (self Instance) ClearcoatTexture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(6))
}

func (self Instance) SetClearcoatTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(6, value)
}

func (self Instance) AnisotropyEnabled() bool {
	return bool(class(self).GetFeature(4))
}

func (self Instance) SetAnisotropyEnabled(value bool) {
	class(self).SetFeature(4, value)
}

func (self Instance) Anisotropy() Float.X {
	return Float.X(Float.X(class(self).GetAnisotropy()))
}

func (self Instance) SetAnisotropy(value Float.X) {
	class(self).SetAnisotropy(float64(value))
}

func (self Instance) AnisotropyFlowmap() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(7))
}

func (self Instance) SetAnisotropyFlowmap(value [1]gdclass.Texture2D) {
	class(self).SetTexture(7, value)
}

func (self Instance) AoEnabled() bool {
	return bool(class(self).GetFeature(5))
}

func (self Instance) SetAoEnabled(value bool) {
	class(self).SetFeature(5, value)
}

func (self Instance) AoLightAffect() Float.X {
	return Float.X(Float.X(class(self).GetAoLightAffect()))
}

func (self Instance) SetAoLightAffect(value Float.X) {
	class(self).SetAoLightAffect(float64(value))
}

func (self Instance) AoTexture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(8))
}

func (self Instance) SetAoTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(8, value)
}

func (self Instance) AoOnUv2() bool {
	return bool(class(self).GetFlag(10))
}

func (self Instance) SetAoOnUv2(value bool) {
	class(self).SetFlag(10, value)
}

func (self Instance) AoTextureChannel() gdclass.BaseMaterial3DTextureChannel {
	return gdclass.BaseMaterial3DTextureChannel(class(self).GetAoTextureChannel())
}

func (self Instance) SetAoTextureChannel(value gdclass.BaseMaterial3DTextureChannel) {
	class(self).SetAoTextureChannel(value)
}

func (self Instance) HeightmapEnabled() bool {
	return bool(class(self).GetFeature(6))
}

func (self Instance) SetHeightmapEnabled(value bool) {
	class(self).SetFeature(6, value)
}

func (self Instance) HeightmapScale() Float.X {
	return Float.X(Float.X(class(self).GetHeightmapScale()))
}

func (self Instance) SetHeightmapScale(value Float.X) {
	class(self).SetHeightmapScale(float64(value))
}

func (self Instance) HeightmapDeepParallax() bool {
	return bool(class(self).IsHeightmapDeepParallaxEnabled())
}

func (self Instance) SetHeightmapDeepParallax(value bool) {
	class(self).SetHeightmapDeepParallax(value)
}

func (self Instance) HeightmapMinLayers() int {
	return int(int(class(self).GetHeightmapDeepParallaxMinLayers()))
}

func (self Instance) SetHeightmapMinLayers(value int) {
	class(self).SetHeightmapDeepParallaxMinLayers(int64(value))
}

func (self Instance) HeightmapMaxLayers() int {
	return int(int(class(self).GetHeightmapDeepParallaxMaxLayers()))
}

func (self Instance) SetHeightmapMaxLayers(value int) {
	class(self).SetHeightmapDeepParallaxMaxLayers(int64(value))
}

func (self Instance) HeightmapFlipTangent() bool {
	return bool(class(self).GetHeightmapDeepParallaxFlipTangent())
}

func (self Instance) SetHeightmapFlipTangent(value bool) {
	class(self).SetHeightmapDeepParallaxFlipTangent(value)
}

func (self Instance) HeightmapFlipBinormal() bool {
	return bool(class(self).GetHeightmapDeepParallaxFlipBinormal())
}

func (self Instance) SetHeightmapFlipBinormal(value bool) {
	class(self).SetHeightmapDeepParallaxFlipBinormal(value)
}

func (self Instance) HeightmapTexture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(9))
}

func (self Instance) SetHeightmapTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(9, value)
}

func (self Instance) HeightmapFlipTexture() bool {
	return bool(class(self).GetFlag(17))
}

func (self Instance) SetHeightmapFlipTexture(value bool) {
	class(self).SetFlag(17, value)
}

func (self Instance) SubsurfScatterEnabled() bool {
	return bool(class(self).GetFeature(7))
}

func (self Instance) SetSubsurfScatterEnabled(value bool) {
	class(self).SetFeature(7, value)
}

func (self Instance) SubsurfScatterStrength() Float.X {
	return Float.X(Float.X(class(self).GetSubsurfaceScatteringStrength()))
}

func (self Instance) SetSubsurfScatterStrength(value Float.X) {
	class(self).SetSubsurfaceScatteringStrength(float64(value))
}

func (self Instance) SubsurfScatterSkinMode() bool {
	return bool(class(self).GetFlag(18))
}

func (self Instance) SetSubsurfScatterSkinMode(value bool) {
	class(self).SetFlag(18, value)
}

func (self Instance) SubsurfScatterTexture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(10))
}

func (self Instance) SetSubsurfScatterTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(10, value)
}

func (self Instance) SubsurfScatterTransmittanceEnabled() bool {
	return bool(class(self).GetFeature(8))
}

func (self Instance) SetSubsurfScatterTransmittanceEnabled(value bool) {
	class(self).SetFeature(8, value)
}

func (self Instance) SubsurfScatterTransmittanceColor() Color.RGBA {
	return Color.RGBA(class(self).GetTransmittanceColor())
}

func (self Instance) SetSubsurfScatterTransmittanceColor(value Color.RGBA) {
	class(self).SetTransmittanceColor(Color.RGBA(value))
}

func (self Instance) SubsurfScatterTransmittanceTexture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(11))
}

func (self Instance) SetSubsurfScatterTransmittanceTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(11, value)
}

func (self Instance) SubsurfScatterTransmittanceDepth() Float.X {
	return Float.X(Float.X(class(self).GetTransmittanceDepth()))
}

func (self Instance) SetSubsurfScatterTransmittanceDepth(value Float.X) {
	class(self).SetTransmittanceDepth(float64(value))
}

func (self Instance) SubsurfScatterTransmittanceBoost() Float.X {
	return Float.X(Float.X(class(self).GetTransmittanceBoost()))
}

func (self Instance) SetSubsurfScatterTransmittanceBoost(value Float.X) {
	class(self).SetTransmittanceBoost(float64(value))
}

func (self Instance) BacklightEnabled() bool {
	return bool(class(self).GetFeature(9))
}

func (self Instance) SetBacklightEnabled(value bool) {
	class(self).SetFeature(9, value)
}

func (self Instance) Backlight() Color.RGBA {
	return Color.RGBA(class(self).GetBacklight())
}

func (self Instance) SetBacklight(value Color.RGBA) {
	class(self).SetBacklight(Color.RGBA(value))
}

func (self Instance) BacklightTexture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(12))
}

func (self Instance) SetBacklightTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(12, value)
}

func (self Instance) RefractionEnabled() bool {
	return bool(class(self).GetFeature(10))
}

func (self Instance) SetRefractionEnabled(value bool) {
	class(self).SetFeature(10, value)
}

func (self Instance) RefractionScale() Float.X {
	return Float.X(Float.X(class(self).GetRefraction()))
}

func (self Instance) SetRefractionScale(value Float.X) {
	class(self).SetRefraction(float64(value))
}

func (self Instance) RefractionTexture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(13))
}

func (self Instance) SetRefractionTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(13, value)
}

func (self Instance) RefractionTextureChannel() gdclass.BaseMaterial3DTextureChannel {
	return gdclass.BaseMaterial3DTextureChannel(class(self).GetRefractionTextureChannel())
}

func (self Instance) SetRefractionTextureChannel(value gdclass.BaseMaterial3DTextureChannel) {
	class(self).SetRefractionTextureChannel(value)
}

func (self Instance) DetailEnabled() bool {
	return bool(class(self).GetFeature(11))
}

func (self Instance) SetDetailEnabled(value bool) {
	class(self).SetFeature(11, value)
}

func (self Instance) DetailMask() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(14))
}

func (self Instance) SetDetailMask(value [1]gdclass.Texture2D) {
	class(self).SetTexture(14, value)
}

func (self Instance) DetailBlendMode() gdclass.BaseMaterial3DBlendMode {
	return gdclass.BaseMaterial3DBlendMode(class(self).GetDetailBlendMode())
}

func (self Instance) SetDetailBlendMode(value gdclass.BaseMaterial3DBlendMode) {
	class(self).SetDetailBlendMode(value)
}

func (self Instance) DetailUvLayer() gdclass.BaseMaterial3DDetailUV {
	return gdclass.BaseMaterial3DDetailUV(class(self).GetDetailUv())
}

func (self Instance) SetDetailUvLayer(value gdclass.BaseMaterial3DDetailUV) {
	class(self).SetDetailUv(value)
}

func (self Instance) DetailAlbedo() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(15))
}

func (self Instance) SetDetailAlbedo(value [1]gdclass.Texture2D) {
	class(self).SetTexture(15, value)
}

func (self Instance) DetailNormal() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(16))
}

func (self Instance) SetDetailNormal(value [1]gdclass.Texture2D) {
	class(self).SetTexture(16, value)
}

func (self Instance) Uv1Scale() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetUv1Scale())
}

func (self Instance) SetUv1Scale(value Vector3.XYZ) {
	class(self).SetUv1Scale(Vector3.XYZ(value))
}

func (self Instance) Uv1Offset() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetUv1Offset())
}

func (self Instance) SetUv1Offset(value Vector3.XYZ) {
	class(self).SetUv1Offset(Vector3.XYZ(value))
}

func (self Instance) Uv1Triplanar() bool {
	return bool(class(self).GetFlag(6))
}

func (self Instance) SetUv1Triplanar(value bool) {
	class(self).SetFlag(6, value)
}

func (self Instance) Uv1TriplanarSharpness() Float.X {
	return Float.X(Float.X(class(self).GetUv1TriplanarBlendSharpness()))
}

func (self Instance) SetUv1TriplanarSharpness(value Float.X) {
	class(self).SetUv1TriplanarBlendSharpness(float64(value))
}

func (self Instance) Uv1WorldTriplanar() bool {
	return bool(class(self).GetFlag(8))
}

func (self Instance) SetUv1WorldTriplanar(value bool) {
	class(self).SetFlag(8, value)
}

func (self Instance) Uv2Scale() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetUv2Scale())
}

func (self Instance) SetUv2Scale(value Vector3.XYZ) {
	class(self).SetUv2Scale(Vector3.XYZ(value))
}

func (self Instance) Uv2Offset() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetUv2Offset())
}

func (self Instance) SetUv2Offset(value Vector3.XYZ) {
	class(self).SetUv2Offset(Vector3.XYZ(value))
}

func (self Instance) Uv2Triplanar() bool {
	return bool(class(self).GetFlag(7))
}

func (self Instance) SetUv2Triplanar(value bool) {
	class(self).SetFlag(7, value)
}

func (self Instance) Uv2TriplanarSharpness() Float.X {
	return Float.X(Float.X(class(self).GetUv2TriplanarBlendSharpness()))
}

func (self Instance) SetUv2TriplanarSharpness(value Float.X) {
	class(self).SetUv2TriplanarBlendSharpness(float64(value))
}

func (self Instance) Uv2WorldTriplanar() bool {
	return bool(class(self).GetFlag(9))
}

func (self Instance) SetUv2WorldTriplanar(value bool) {
	class(self).SetFlag(9, value)
}

func (self Instance) TextureFilter() gdclass.BaseMaterial3DTextureFilter {
	return gdclass.BaseMaterial3DTextureFilter(class(self).GetTextureFilter())
}

func (self Instance) SetTextureFilter(value gdclass.BaseMaterial3DTextureFilter) {
	class(self).SetTextureFilter(value)
}

func (self Instance) TextureRepeat() bool {
	return bool(class(self).GetFlag(16))
}

func (self Instance) SetTextureRepeat(value bool) {
	class(self).SetFlag(16, value)
}

func (self Instance) DisableReceiveShadows() bool {
	return bool(class(self).GetFlag(13))
}

func (self Instance) SetDisableReceiveShadows(value bool) {
	class(self).SetFlag(13, value)
}

func (self Instance) ShadowToOpacity() bool {
	return bool(class(self).GetFlag(15))
}

func (self Instance) SetShadowToOpacity(value bool) {
	class(self).SetFlag(15, value)
}

func (self Instance) BillboardMode() gdclass.BaseMaterial3DBillboardMode {
	return gdclass.BaseMaterial3DBillboardMode(class(self).GetBillboardMode())
}

func (self Instance) SetBillboardMode(value gdclass.BaseMaterial3DBillboardMode) {
	class(self).SetBillboardMode(value)
}

func (self Instance) BillboardKeepScale() bool {
	return bool(class(self).GetFlag(5))
}

func (self Instance) SetBillboardKeepScale(value bool) {
	class(self).SetFlag(5, value)
}

func (self Instance) ParticlesAnimHFrames() int {
	return int(int(class(self).GetParticlesAnimHFrames()))
}

func (self Instance) SetParticlesAnimHFrames(value int) {
	class(self).SetParticlesAnimHFrames(int64(value))
}

func (self Instance) ParticlesAnimVFrames() int {
	return int(int(class(self).GetParticlesAnimVFrames()))
}

func (self Instance) SetParticlesAnimVFrames(value int) {
	class(self).SetParticlesAnimVFrames(int64(value))
}

func (self Instance) ParticlesAnimLoop() bool {
	return bool(class(self).GetParticlesAnimLoop())
}

func (self Instance) SetParticlesAnimLoop(value bool) {
	class(self).SetParticlesAnimLoop(value)
}

func (self Instance) Grow() bool {
	return bool(class(self).IsGrowEnabled())
}

func (self Instance) SetGrow(value bool) {
	class(self).SetGrowEnabled(value)
}

func (self Instance) GrowAmount() Float.X {
	return Float.X(Float.X(class(self).GetGrow()))
}

func (self Instance) SetGrowAmount(value Float.X) {
	class(self).SetGrow(float64(value))
}

func (self Instance) FixedSize() bool {
	return bool(class(self).GetFlag(4))
}

func (self Instance) SetFixedSize(value bool) {
	class(self).SetFlag(4, value)
}

func (self Instance) UsePointSize() bool {
	return bool(class(self).GetFlag(3))
}

func (self Instance) SetUsePointSize(value bool) {
	class(self).SetFlag(3, value)
}

func (self Instance) PointSize() Float.X {
	return Float.X(Float.X(class(self).GetPointSize()))
}

func (self Instance) SetPointSize(value Float.X) {
	class(self).SetPointSize(float64(value))
}

func (self Instance) UseParticleTrails() bool {
	return bool(class(self).GetFlag(19))
}

func (self Instance) SetUseParticleTrails(value bool) {
	class(self).SetFlag(19, value)
}

func (self Instance) ProximityFadeEnabled() bool {
	return bool(class(self).IsProximityFadeEnabled())
}

func (self Instance) SetProximityFadeEnabled(value bool) {
	class(self).SetProximityFadeEnabled(value)
}

func (self Instance) ProximityFadeDistance() Float.X {
	return Float.X(Float.X(class(self).GetProximityFadeDistance()))
}

func (self Instance) SetProximityFadeDistance(value Float.X) {
	class(self).SetProximityFadeDistance(float64(value))
}

func (self Instance) MsdfPixelRange() Float.X {
	return Float.X(Float.X(class(self).GetMsdfPixelRange()))
}

func (self Instance) SetMsdfPixelRange(value Float.X) {
	class(self).SetMsdfPixelRange(float64(value))
}

func (self Instance) MsdfOutlineSize() Float.X {
	return Float.X(Float.X(class(self).GetMsdfOutlineSize()))
}

func (self Instance) SetMsdfOutlineSize(value Float.X) {
	class(self).SetMsdfOutlineSize(float64(value))
}

func (self Instance) DistanceFadeMode() gdclass.BaseMaterial3DDistanceFadeMode {
	return gdclass.BaseMaterial3DDistanceFadeMode(class(self).GetDistanceFade())
}

func (self Instance) SetDistanceFadeMode(value gdclass.BaseMaterial3DDistanceFadeMode) {
	class(self).SetDistanceFade(value)
}

func (self Instance) DistanceFadeMinDistance() Float.X {
	return Float.X(Float.X(class(self).GetDistanceFadeMinDistance()))
}

func (self Instance) SetDistanceFadeMinDistance(value Float.X) {
	class(self).SetDistanceFadeMinDistance(float64(value))
}

func (self Instance) DistanceFadeMaxDistance() Float.X {
	return Float.X(Float.X(class(self).GetDistanceFadeMaxDistance()))
}

func (self Instance) SetDistanceFadeMaxDistance(value Float.X) {
	class(self).SetDistanceFadeMaxDistance(float64(value))
}

//go:nosplit
func (self class) SetAlbedo(albedo Color.RGBA) { //gd:BaseMaterial3D.set_albedo
	var frame = callframe.New()
	callframe.Arg(frame, albedo)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_albedo, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlbedo() Color.RGBA { //gd:BaseMaterial3D.get_albedo
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_albedo, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTransparency(transparency gdclass.BaseMaterial3DTransparency) { //gd:BaseMaterial3D.set_transparency
	var frame = callframe.New()
	callframe.Arg(frame, transparency)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_transparency, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTransparency() gdclass.BaseMaterial3DTransparency { //gd:BaseMaterial3D.get_transparency
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DTransparency](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_transparency, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlphaAntialiasing(alpha_aa gdclass.BaseMaterial3DAlphaAntiAliasing) { //gd:BaseMaterial3D.set_alpha_antialiasing
	var frame = callframe.New()
	callframe.Arg(frame, alpha_aa)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_alpha_antialiasing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlphaAntialiasing() gdclass.BaseMaterial3DAlphaAntiAliasing { //gd:BaseMaterial3D.get_alpha_antialiasing
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DAlphaAntiAliasing](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_alpha_antialiasing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlphaAntialiasingEdge(edge float64) { //gd:BaseMaterial3D.set_alpha_antialiasing_edge
	var frame = callframe.New()
	callframe.Arg(frame, edge)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_alpha_antialiasing_edge, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlphaAntialiasingEdge() float64 { //gd:BaseMaterial3D.get_alpha_antialiasing_edge
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_alpha_antialiasing_edge, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShadingMode(shading_mode gdclass.BaseMaterial3DShadingMode) { //gd:BaseMaterial3D.set_shading_mode
	var frame = callframe.New()
	callframe.Arg(frame, shading_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_shading_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShadingMode() gdclass.BaseMaterial3DShadingMode { //gd:BaseMaterial3D.get_shading_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DShadingMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_shading_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpecular(specular float64) { //gd:BaseMaterial3D.set_specular
	var frame = callframe.New()
	callframe.Arg(frame, specular)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_specular, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpecular() float64 { //gd:BaseMaterial3D.get_specular
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_specular, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMetallic(metallic float64) { //gd:BaseMaterial3D.set_metallic
	var frame = callframe.New()
	callframe.Arg(frame, metallic)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_metallic, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMetallic() float64 { //gd:BaseMaterial3D.get_metallic
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_metallic, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRoughness(roughness float64) { //gd:BaseMaterial3D.set_roughness
	var frame = callframe.New()
	callframe.Arg(frame, roughness)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_roughness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRoughness() float64 { //gd:BaseMaterial3D.get_roughness
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_roughness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmission(emission Color.RGBA) { //gd:BaseMaterial3D.set_emission
	var frame = callframe.New()
	callframe.Arg(frame, emission)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_emission, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmission() Color.RGBA { //gd:BaseMaterial3D.get_emission
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_emission, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionEnergyMultiplier(emission_energy_multiplier float64) { //gd:BaseMaterial3D.set_emission_energy_multiplier
	var frame = callframe.New()
	callframe.Arg(frame, emission_energy_multiplier)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_emission_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionEnergyMultiplier() float64 { //gd:BaseMaterial3D.get_emission_energy_multiplier
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_emission_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionIntensity(emission_energy_multiplier float64) { //gd:BaseMaterial3D.set_emission_intensity
	var frame = callframe.New()
	callframe.Arg(frame, emission_energy_multiplier)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_emission_intensity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionIntensity() float64 { //gd:BaseMaterial3D.get_emission_intensity
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_emission_intensity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNormalScale(normal_scale float64) { //gd:BaseMaterial3D.set_normal_scale
	var frame = callframe.New()
	callframe.Arg(frame, normal_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_normal_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNormalScale() float64 { //gd:BaseMaterial3D.get_normal_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_normal_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRim(rim float64) { //gd:BaseMaterial3D.set_rim
	var frame = callframe.New()
	callframe.Arg(frame, rim)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_rim, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRim() float64 { //gd:BaseMaterial3D.get_rim
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_rim, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRimTint(rim_tint float64) { //gd:BaseMaterial3D.set_rim_tint
	var frame = callframe.New()
	callframe.Arg(frame, rim_tint)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_rim_tint, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRimTint() float64 { //gd:BaseMaterial3D.get_rim_tint
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_rim_tint, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetClearcoat(clearcoat float64) { //gd:BaseMaterial3D.set_clearcoat
	var frame = callframe.New()
	callframe.Arg(frame, clearcoat)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_clearcoat, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetClearcoat() float64 { //gd:BaseMaterial3D.get_clearcoat
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_clearcoat, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetClearcoatRoughness(clearcoat_roughness float64) { //gd:BaseMaterial3D.set_clearcoat_roughness
	var frame = callframe.New()
	callframe.Arg(frame, clearcoat_roughness)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_clearcoat_roughness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetClearcoatRoughness() float64 { //gd:BaseMaterial3D.get_clearcoat_roughness
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_clearcoat_roughness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAnisotropy(anisotropy float64) { //gd:BaseMaterial3D.set_anisotropy
	var frame = callframe.New()
	callframe.Arg(frame, anisotropy)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_anisotropy, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAnisotropy() float64 { //gd:BaseMaterial3D.get_anisotropy
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_anisotropy, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHeightmapScale(heightmap_scale float64) { //gd:BaseMaterial3D.set_heightmap_scale
	var frame = callframe.New()
	callframe.Arg(frame, heightmap_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_heightmap_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHeightmapScale() float64 { //gd:BaseMaterial3D.get_heightmap_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_heightmap_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubsurfaceScatteringStrength(strength float64) { //gd:BaseMaterial3D.set_subsurface_scattering_strength
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_subsurface_scattering_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubsurfaceScatteringStrength() float64 { //gd:BaseMaterial3D.get_subsurface_scattering_strength
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_subsurface_scattering_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTransmittanceColor(color Color.RGBA) { //gd:BaseMaterial3D.set_transmittance_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_transmittance_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTransmittanceColor() Color.RGBA { //gd:BaseMaterial3D.get_transmittance_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_transmittance_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTransmittanceDepth(depth float64) { //gd:BaseMaterial3D.set_transmittance_depth
	var frame = callframe.New()
	callframe.Arg(frame, depth)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_transmittance_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTransmittanceDepth() float64 { //gd:BaseMaterial3D.get_transmittance_depth
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_transmittance_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTransmittanceBoost(boost float64) { //gd:BaseMaterial3D.set_transmittance_boost
	var frame = callframe.New()
	callframe.Arg(frame, boost)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_transmittance_boost, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTransmittanceBoost() float64 { //gd:BaseMaterial3D.get_transmittance_boost
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_transmittance_boost, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBacklight(backlight Color.RGBA) { //gd:BaseMaterial3D.set_backlight
	var frame = callframe.New()
	callframe.Arg(frame, backlight)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_backlight, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBacklight() Color.RGBA { //gd:BaseMaterial3D.get_backlight
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_backlight, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRefraction(refraction float64) { //gd:BaseMaterial3D.set_refraction
	var frame = callframe.New()
	callframe.Arg(frame, refraction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_refraction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRefraction() float64 { //gd:BaseMaterial3D.get_refraction
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_refraction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPointSize(point_size float64) { //gd:BaseMaterial3D.set_point_size
	var frame = callframe.New()
	callframe.Arg(frame, point_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_point_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPointSize() float64 { //gd:BaseMaterial3D.get_point_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_point_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDetailUv(detail_uv gdclass.BaseMaterial3DDetailUV) { //gd:BaseMaterial3D.set_detail_uv
	var frame = callframe.New()
	callframe.Arg(frame, detail_uv)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_detail_uv, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDetailUv() gdclass.BaseMaterial3DDetailUV { //gd:BaseMaterial3D.get_detail_uv
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DDetailUV](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_detail_uv, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBlendMode(blend_mode gdclass.BaseMaterial3DBlendMode) { //gd:BaseMaterial3D.set_blend_mode
	var frame = callframe.New()
	callframe.Arg(frame, blend_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_blend_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBlendMode() gdclass.BaseMaterial3DBlendMode { //gd:BaseMaterial3D.get_blend_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DBlendMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_blend_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDepthDrawMode(depth_draw_mode gdclass.BaseMaterial3DDepthDrawMode) { //gd:BaseMaterial3D.set_depth_draw_mode
	var frame = callframe.New()
	callframe.Arg(frame, depth_draw_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_depth_draw_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepthDrawMode() gdclass.BaseMaterial3DDepthDrawMode { //gd:BaseMaterial3D.get_depth_draw_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DDepthDrawMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_depth_draw_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCullMode(cull_mode gdclass.BaseMaterial3DCullMode) { //gd:BaseMaterial3D.set_cull_mode
	var frame = callframe.New()
	callframe.Arg(frame, cull_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_cull_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCullMode() gdclass.BaseMaterial3DCullMode { //gd:BaseMaterial3D.get_cull_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DCullMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_cull_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDiffuseMode(diffuse_mode gdclass.BaseMaterial3DDiffuseMode) { //gd:BaseMaterial3D.set_diffuse_mode
	var frame = callframe.New()
	callframe.Arg(frame, diffuse_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_diffuse_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDiffuseMode() gdclass.BaseMaterial3DDiffuseMode { //gd:BaseMaterial3D.get_diffuse_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DDiffuseMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_diffuse_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpecularMode(specular_mode gdclass.BaseMaterial3DSpecularMode) { //gd:BaseMaterial3D.set_specular_mode
	var frame = callframe.New()
	callframe.Arg(frame, specular_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_specular_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpecularMode() gdclass.BaseMaterial3DSpecularMode { //gd:BaseMaterial3D.get_specular_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DSpecularMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_specular_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [code]true[/code], enables the specified flag. Flags are optional behavior that can be turned on and off. Only one flag can be enabled at a time with this function, the flag enumerators cannot be bit-masked together to enable or disable multiple flags at once. Flags can also be enabled by setting the corresponding member to [code]true[/code]. See [enum Flags] enumerator for options.
*/
//go:nosplit
func (self class) SetFlag(flag gdclass.BaseMaterial3DFlags, enable bool) { //gd:BaseMaterial3D.set_flag
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_flag, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code], if the specified flag is enabled. See [enum Flags] enumerator for options.
*/
//go:nosplit
func (self class) GetFlag(flag gdclass.BaseMaterial3DFlags) bool { //gd:BaseMaterial3D.get_flag
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_flag, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureFilter(mode gdclass.BaseMaterial3DTextureFilter) { //gd:BaseMaterial3D.set_texture_filter
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_texture_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureFilter() gdclass.BaseMaterial3DTextureFilter { //gd:BaseMaterial3D.get_texture_filter
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DTextureFilter](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_texture_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [code]true[/code], enables the specified [enum Feature]. Many features that are available in [BaseMaterial3D]s need to be enabled before use. This way the cost for using the feature is only incurred when specified. Features can also be enabled by setting the corresponding member to [code]true[/code].
*/
//go:nosplit
func (self class) SetFeature(feature gdclass.BaseMaterial3DFeature, enable bool) { //gd:BaseMaterial3D.set_feature
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_feature, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code], if the specified [enum Feature] is enabled.
*/
//go:nosplit
func (self class) GetFeature(feature gdclass.BaseMaterial3DFeature) bool { //gd:BaseMaterial3D.get_feature
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_feature, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the texture for the slot specified by [param param]. See [enum TextureParam] for available slots.
*/
//go:nosplit
func (self class) SetTexture(param gdclass.BaseMaterial3DTextureParam, texture [1]gdclass.Texture2D) { //gd:BaseMaterial3D.set_texture
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [Texture2D] associated with the specified [enum TextureParam].
*/
//go:nosplit
func (self class) GetTexture(param gdclass.BaseMaterial3DTextureParam) [1]gdclass.Texture2D { //gd:BaseMaterial3D.get_texture
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDetailBlendMode(detail_blend_mode gdclass.BaseMaterial3DBlendMode) { //gd:BaseMaterial3D.set_detail_blend_mode
	var frame = callframe.New()
	callframe.Arg(frame, detail_blend_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_detail_blend_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDetailBlendMode() gdclass.BaseMaterial3DBlendMode { //gd:BaseMaterial3D.get_detail_blend_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DBlendMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_detail_blend_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUv1Scale(scale Vector3.XYZ) { //gd:BaseMaterial3D.set_uv1_scale
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_uv1_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUv1Scale() Vector3.XYZ { //gd:BaseMaterial3D.get_uv1_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_uv1_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUv1Offset(offset Vector3.XYZ) { //gd:BaseMaterial3D.set_uv1_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_uv1_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUv1Offset() Vector3.XYZ { //gd:BaseMaterial3D.get_uv1_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_uv1_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUv1TriplanarBlendSharpness(sharpness float64) { //gd:BaseMaterial3D.set_uv1_triplanar_blend_sharpness
	var frame = callframe.New()
	callframe.Arg(frame, sharpness)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_uv1_triplanar_blend_sharpness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUv1TriplanarBlendSharpness() float64 { //gd:BaseMaterial3D.get_uv1_triplanar_blend_sharpness
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_uv1_triplanar_blend_sharpness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUv2Scale(scale Vector3.XYZ) { //gd:BaseMaterial3D.set_uv2_scale
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_uv2_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUv2Scale() Vector3.XYZ { //gd:BaseMaterial3D.get_uv2_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_uv2_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUv2Offset(offset Vector3.XYZ) { //gd:BaseMaterial3D.set_uv2_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_uv2_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUv2Offset() Vector3.XYZ { //gd:BaseMaterial3D.get_uv2_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_uv2_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUv2TriplanarBlendSharpness(sharpness float64) { //gd:BaseMaterial3D.set_uv2_triplanar_blend_sharpness
	var frame = callframe.New()
	callframe.Arg(frame, sharpness)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_uv2_triplanar_blend_sharpness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUv2TriplanarBlendSharpness() float64 { //gd:BaseMaterial3D.get_uv2_triplanar_blend_sharpness
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_uv2_triplanar_blend_sharpness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBillboardMode(mode gdclass.BaseMaterial3DBillboardMode) { //gd:BaseMaterial3D.set_billboard_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_billboard_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBillboardMode() gdclass.BaseMaterial3DBillboardMode { //gd:BaseMaterial3D.get_billboard_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DBillboardMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_billboard_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetParticlesAnimHFrames(frames int64) { //gd:BaseMaterial3D.set_particles_anim_h_frames
	var frame = callframe.New()
	callframe.Arg(frame, frames)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_particles_anim_h_frames, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetParticlesAnimHFrames() int64 { //gd:BaseMaterial3D.get_particles_anim_h_frames
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_particles_anim_h_frames, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetParticlesAnimVFrames(frames int64) { //gd:BaseMaterial3D.set_particles_anim_v_frames
	var frame = callframe.New()
	callframe.Arg(frame, frames)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_particles_anim_v_frames, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetParticlesAnimVFrames() int64 { //gd:BaseMaterial3D.get_particles_anim_v_frames
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_particles_anim_v_frames, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetParticlesAnimLoop(loop bool) { //gd:BaseMaterial3D.set_particles_anim_loop
	var frame = callframe.New()
	callframe.Arg(frame, loop)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_particles_anim_loop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetParticlesAnimLoop() bool { //gd:BaseMaterial3D.get_particles_anim_loop
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_particles_anim_loop, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHeightmapDeepParallax(enable bool) { //gd:BaseMaterial3D.set_heightmap_deep_parallax
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_heightmap_deep_parallax, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsHeightmapDeepParallaxEnabled() bool { //gd:BaseMaterial3D.is_heightmap_deep_parallax_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_is_heightmap_deep_parallax_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHeightmapDeepParallaxMinLayers(layer int64) { //gd:BaseMaterial3D.set_heightmap_deep_parallax_min_layers
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_heightmap_deep_parallax_min_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHeightmapDeepParallaxMinLayers() int64 { //gd:BaseMaterial3D.get_heightmap_deep_parallax_min_layers
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_heightmap_deep_parallax_min_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHeightmapDeepParallaxMaxLayers(layer int64) { //gd:BaseMaterial3D.set_heightmap_deep_parallax_max_layers
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_heightmap_deep_parallax_max_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHeightmapDeepParallaxMaxLayers() int64 { //gd:BaseMaterial3D.get_heightmap_deep_parallax_max_layers
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_heightmap_deep_parallax_max_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHeightmapDeepParallaxFlipTangent(flip bool) { //gd:BaseMaterial3D.set_heightmap_deep_parallax_flip_tangent
	var frame = callframe.New()
	callframe.Arg(frame, flip)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_heightmap_deep_parallax_flip_tangent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHeightmapDeepParallaxFlipTangent() bool { //gd:BaseMaterial3D.get_heightmap_deep_parallax_flip_tangent
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_heightmap_deep_parallax_flip_tangent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHeightmapDeepParallaxFlipBinormal(flip bool) { //gd:BaseMaterial3D.set_heightmap_deep_parallax_flip_binormal
	var frame = callframe.New()
	callframe.Arg(frame, flip)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_heightmap_deep_parallax_flip_binormal, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHeightmapDeepParallaxFlipBinormal() bool { //gd:BaseMaterial3D.get_heightmap_deep_parallax_flip_binormal
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_heightmap_deep_parallax_flip_binormal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGrow(amount float64) { //gd:BaseMaterial3D.set_grow
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_grow, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGrow() float64 { //gd:BaseMaterial3D.get_grow
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_grow, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionOperator(operator gdclass.BaseMaterial3DEmissionOperator) { //gd:BaseMaterial3D.set_emission_operator
	var frame = callframe.New()
	callframe.Arg(frame, operator)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_emission_operator, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionOperator() gdclass.BaseMaterial3DEmissionOperator { //gd:BaseMaterial3D.get_emission_operator
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DEmissionOperator](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_emission_operator, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAoLightAffect(amount float64) { //gd:BaseMaterial3D.set_ao_light_affect
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_ao_light_affect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAoLightAffect() float64 { //gd:BaseMaterial3D.get_ao_light_affect
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_ao_light_affect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlphaScissorThreshold(threshold float64) { //gd:BaseMaterial3D.set_alpha_scissor_threshold
	var frame = callframe.New()
	callframe.Arg(frame, threshold)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_alpha_scissor_threshold, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlphaScissorThreshold() float64 { //gd:BaseMaterial3D.get_alpha_scissor_threshold
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_alpha_scissor_threshold, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlphaHashScale(threshold float64) { //gd:BaseMaterial3D.set_alpha_hash_scale
	var frame = callframe.New()
	callframe.Arg(frame, threshold)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_alpha_hash_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlphaHashScale() float64 { //gd:BaseMaterial3D.get_alpha_hash_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_alpha_hash_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGrowEnabled(enable bool) { //gd:BaseMaterial3D.set_grow_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_grow_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsGrowEnabled() bool { //gd:BaseMaterial3D.is_grow_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_is_grow_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMetallicTextureChannel(channel gdclass.BaseMaterial3DTextureChannel) { //gd:BaseMaterial3D.set_metallic_texture_channel
	var frame = callframe.New()
	callframe.Arg(frame, channel)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_metallic_texture_channel, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMetallicTextureChannel() gdclass.BaseMaterial3DTextureChannel { //gd:BaseMaterial3D.get_metallic_texture_channel
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DTextureChannel](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_metallic_texture_channel, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRoughnessTextureChannel(channel gdclass.BaseMaterial3DTextureChannel) { //gd:BaseMaterial3D.set_roughness_texture_channel
	var frame = callframe.New()
	callframe.Arg(frame, channel)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_roughness_texture_channel, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRoughnessTextureChannel() gdclass.BaseMaterial3DTextureChannel { //gd:BaseMaterial3D.get_roughness_texture_channel
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DTextureChannel](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_roughness_texture_channel, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAoTextureChannel(channel gdclass.BaseMaterial3DTextureChannel) { //gd:BaseMaterial3D.set_ao_texture_channel
	var frame = callframe.New()
	callframe.Arg(frame, channel)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_ao_texture_channel, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAoTextureChannel() gdclass.BaseMaterial3DTextureChannel { //gd:BaseMaterial3D.get_ao_texture_channel
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DTextureChannel](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_ao_texture_channel, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRefractionTextureChannel(channel gdclass.BaseMaterial3DTextureChannel) { //gd:BaseMaterial3D.set_refraction_texture_channel
	var frame = callframe.New()
	callframe.Arg(frame, channel)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_refraction_texture_channel, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRefractionTextureChannel() gdclass.BaseMaterial3DTextureChannel { //gd:BaseMaterial3D.get_refraction_texture_channel
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DTextureChannel](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_refraction_texture_channel, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProximityFadeEnabled(enabled bool) { //gd:BaseMaterial3D.set_proximity_fade_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_proximity_fade_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsProximityFadeEnabled() bool { //gd:BaseMaterial3D.is_proximity_fade_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_is_proximity_fade_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProximityFadeDistance(distance float64) { //gd:BaseMaterial3D.set_proximity_fade_distance
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_proximity_fade_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetProximityFadeDistance() float64 { //gd:BaseMaterial3D.get_proximity_fade_distance
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_proximity_fade_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMsdfPixelRange(arange float64) { //gd:BaseMaterial3D.set_msdf_pixel_range
	var frame = callframe.New()
	callframe.Arg(frame, arange)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMsdfPixelRange() float64 { //gd:BaseMaterial3D.get_msdf_pixel_range
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMsdfOutlineSize(size float64) { //gd:BaseMaterial3D.set_msdf_outline_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_msdf_outline_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMsdfOutlineSize() float64 { //gd:BaseMaterial3D.get_msdf_outline_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_msdf_outline_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDistanceFade(mode gdclass.BaseMaterial3DDistanceFadeMode) { //gd:BaseMaterial3D.set_distance_fade
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_distance_fade, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDistanceFade() gdclass.BaseMaterial3DDistanceFadeMode { //gd:BaseMaterial3D.get_distance_fade
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DDistanceFadeMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_distance_fade, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDistanceFadeMaxDistance(distance float64) { //gd:BaseMaterial3D.set_distance_fade_max_distance
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_distance_fade_max_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDistanceFadeMaxDistance() float64 { //gd:BaseMaterial3D.get_distance_fade_max_distance
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_distance_fade_max_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDistanceFadeMinDistance(distance float64) { //gd:BaseMaterial3D.set_distance_fade_min_distance
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_set_distance_fade_min_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDistanceFadeMinDistance() float64 { //gd:BaseMaterial3D.get_distance_fade_min_distance
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseMaterial3D.Bind_get_distance_fade_min_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsBaseMaterial3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsBaseMaterial3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Material.Advanced(self.AsMaterial()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Material.Instance(self.AsMaterial()), name)
	}
}
func init() {
	gdclass.Register("BaseMaterial3D", func(ptr gd.Object) any {
		return [1]gdclass.BaseMaterial3D{*(*gdclass.BaseMaterial3D)(unsafe.Pointer(&ptr))}
	})
}

type TextureParam = gdclass.BaseMaterial3DTextureParam //gd:BaseMaterial3D.TextureParam

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

type TextureFilter = gdclass.BaseMaterial3DTextureFilter //gd:BaseMaterial3D.TextureFilter

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

type DetailUV = gdclass.BaseMaterial3DDetailUV //gd:BaseMaterial3D.DetailUV

const (
	/*Use [code]UV[/code] with the detail texture.*/
	DetailUv1 DetailUV = 0
	/*Use [code]UV2[/code] with the detail texture.*/
	DetailUv2 DetailUV = 1
)

type Transparency = gdclass.BaseMaterial3DTransparency //gd:BaseMaterial3D.Transparency

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

type ShadingMode = gdclass.BaseMaterial3DShadingMode //gd:BaseMaterial3D.ShadingMode

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

type Feature = gdclass.BaseMaterial3DFeature //gd:BaseMaterial3D.Feature

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

type BlendMode = gdclass.BaseMaterial3DBlendMode //gd:BaseMaterial3D.BlendMode

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

type AlphaAntiAliasing = gdclass.BaseMaterial3DAlphaAntiAliasing //gd:BaseMaterial3D.AlphaAntiAliasing

const (
	/*Disables Alpha AntiAliasing for the material.*/
	AlphaAntialiasingOff AlphaAntiAliasing = 0
	/*Enables AlphaToCoverage. Alpha values in the material are passed to the AntiAliasing sample mask.*/
	AlphaAntialiasingAlphaToCoverage AlphaAntiAliasing = 1
	/*Enables AlphaToCoverage and forces all non-zero alpha values to [code]1[/code]. Alpha values in the material are passed to the AntiAliasing sample mask.*/
	AlphaAntialiasingAlphaToCoverageAndToOne AlphaAntiAliasing = 2
)

type DepthDrawMode = gdclass.BaseMaterial3DDepthDrawMode //gd:BaseMaterial3D.DepthDrawMode

const (
	/*Default depth draw mode. Depth is drawn only for opaque objects during the opaque prepass (if any) and during the opaque pass.*/
	DepthDrawOpaqueOnly DepthDrawMode = 0
	/*Objects will write to depth during the opaque and the transparent passes. Transparent objects that are close to the camera may obscure other transparent objects behind them.
	  [b]Note:[/b] This does not influence whether transparent objects are included in the depth prepass or not. For that, see [enum Transparency].*/
	DepthDrawAlways DepthDrawMode = 1
	/*Objects will not write their depth to the depth buffer, even during the depth prepass (if enabled).*/
	DepthDrawDisabled DepthDrawMode = 2
)

type CullMode = gdclass.BaseMaterial3DCullMode //gd:BaseMaterial3D.CullMode

const (
	/*Default cull mode. The back of the object is culled when not visible. Back face triangles will be culled when facing the camera. This results in only the front side of triangles being drawn. For closed-surface meshes, this means that only the exterior of the mesh will be visible.*/
	CullBack CullMode = 0
	/*Front face triangles will be culled when facing the camera. This results in only the back side of triangles being drawn. For closed-surface meshes, this means that the interior of the mesh will be drawn instead of the exterior.*/
	CullFront CullMode = 1
	/*No face culling is performed; both the front face and back face will be visible.*/
	CullDisabled CullMode = 2
)

type Flags = gdclass.BaseMaterial3DFlags //gd:BaseMaterial3D.Flags

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

type DiffuseMode = gdclass.BaseMaterial3DDiffuseMode //gd:BaseMaterial3D.DiffuseMode

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

type SpecularMode = gdclass.BaseMaterial3DSpecularMode //gd:BaseMaterial3D.SpecularMode

const (
	/*Default specular blob.*/
	SpecularSchlickGgx SpecularMode = 0
	/*Toon blob which changes size based on roughness.*/
	SpecularToon SpecularMode = 1
	/*No specular blob. This is slightly faster to render than other specular modes.*/
	SpecularDisabled SpecularMode = 2
)

type BillboardMode = gdclass.BaseMaterial3DBillboardMode //gd:BaseMaterial3D.BillboardMode

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

type TextureChannel = gdclass.BaseMaterial3DTextureChannel //gd:BaseMaterial3D.TextureChannel

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

type EmissionOperator = gdclass.BaseMaterial3DEmissionOperator //gd:BaseMaterial3D.EmissionOperator

const (
	/*Adds the emission color to the color from the emission texture.*/
	EmissionOpAdd EmissionOperator = 0
	/*Multiplies the emission color by the color from the emission texture.*/
	EmissionOpMultiply EmissionOperator = 1
)

type DistanceFadeMode = gdclass.BaseMaterial3DDistanceFadeMode //gd:BaseMaterial3D.DistanceFadeMode

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
