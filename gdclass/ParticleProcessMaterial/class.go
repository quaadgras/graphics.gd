package ParticleProcessMaterial

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
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
var _ = discreet.Root

/*
[ParticleProcessMaterial] defines particle properties and behavior. It is used in the [code]process_material[/code] of the [GPUParticles2D] and [GPUParticles3D] nodes. Some of this material's properties are applied to each particle when emitted, while others can have a [CurveTexture] or a [GradientTexture1D] applied to vary numerical or color values over the lifetime of the particle.

*/
type Go [1]classdb.ParticleProcessMaterial
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ParticleProcessMaterial
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ParticleProcessMaterial"))
	return Go{classdb.ParticleProcessMaterial(object)}
}

func (self Go) LifetimeRandomness() float64 {
		return float64(float64(class(self).GetLifetimeRandomness()))
}

func (self Go) SetLifetimeRandomness(value float64) {
	class(self).SetLifetimeRandomness(gd.Float(value))
}

func (self Go) ParticleFlagAlignY() bool {
		return bool(class(self).GetParticleFlag(0))
}

func (self Go) SetParticleFlagAlignY(value bool) {
	class(self).SetParticleFlag(0, value)
}

func (self Go) ParticleFlagRotateY() bool {
		return bool(class(self).GetParticleFlag(1))
}

func (self Go) SetParticleFlagRotateY(value bool) {
	class(self).SetParticleFlag(1, value)
}

func (self Go) ParticleFlagDisableZ() bool {
		return bool(class(self).GetParticleFlag(2))
}

func (self Go) SetParticleFlagDisableZ(value bool) {
	class(self).SetParticleFlag(2, value)
}

func (self Go) ParticleFlagDampingAsFriction() bool {
		return bool(class(self).GetParticleFlag(3))
}

func (self Go) SetParticleFlagDampingAsFriction(value bool) {
	class(self).SetParticleFlag(3, value)
}

func (self Go) EmissionShapeOffset() gd.Vector3 {
		return gd.Vector3(class(self).GetEmissionShapeOffset())
}

func (self Go) SetEmissionShapeOffset(value gd.Vector3) {
	class(self).SetEmissionShapeOffset(value)
}

func (self Go) EmissionShapeScale() gd.Vector3 {
		return gd.Vector3(class(self).GetEmissionShapeScale())
}

func (self Go) SetEmissionShapeScale(value gd.Vector3) {
	class(self).SetEmissionShapeScale(value)
}

func (self Go) EmissionShape() classdb.ParticleProcessMaterialEmissionShape {
		return classdb.ParticleProcessMaterialEmissionShape(class(self).GetEmissionShape())
}

func (self Go) SetEmissionShape(value classdb.ParticleProcessMaterialEmissionShape) {
	class(self).SetEmissionShape(value)
}

func (self Go) EmissionSphereRadius() float64 {
		return float64(float64(class(self).GetEmissionSphereRadius()))
}

func (self Go) SetEmissionSphereRadius(value float64) {
	class(self).SetEmissionSphereRadius(gd.Float(value))
}

func (self Go) EmissionBoxExtents() gd.Vector3 {
		return gd.Vector3(class(self).GetEmissionBoxExtents())
}

func (self Go) SetEmissionBoxExtents(value gd.Vector3) {
	class(self).SetEmissionBoxExtents(value)
}

func (self Go) EmissionPointTexture() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetEmissionPointTexture())
}

func (self Go) SetEmissionPointTexture(value gdclass.Texture2D) {
	class(self).SetEmissionPointTexture(value)
}

func (self Go) EmissionNormalTexture() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetEmissionNormalTexture())
}

func (self Go) SetEmissionNormalTexture(value gdclass.Texture2D) {
	class(self).SetEmissionNormalTexture(value)
}

func (self Go) EmissionColorTexture() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetEmissionColorTexture())
}

func (self Go) SetEmissionColorTexture(value gdclass.Texture2D) {
	class(self).SetEmissionColorTexture(value)
}

func (self Go) EmissionPointCount() int {
		return int(int(class(self).GetEmissionPointCount()))
}

func (self Go) SetEmissionPointCount(value int) {
	class(self).SetEmissionPointCount(gd.Int(value))
}

func (self Go) EmissionRingAxis() gd.Vector3 {
		return gd.Vector3(class(self).GetEmissionRingAxis())
}

func (self Go) SetEmissionRingAxis(value gd.Vector3) {
	class(self).SetEmissionRingAxis(value)
}

func (self Go) EmissionRingHeight() float64 {
		return float64(float64(class(self).GetEmissionRingHeight()))
}

func (self Go) SetEmissionRingHeight(value float64) {
	class(self).SetEmissionRingHeight(gd.Float(value))
}

func (self Go) EmissionRingRadius() float64 {
		return float64(float64(class(self).GetEmissionRingRadius()))
}

func (self Go) SetEmissionRingRadius(value float64) {
	class(self).SetEmissionRingRadius(gd.Float(value))
}

func (self Go) EmissionRingInnerRadius() float64 {
		return float64(float64(class(self).GetEmissionRingInnerRadius()))
}

func (self Go) SetEmissionRingInnerRadius(value float64) {
	class(self).SetEmissionRingInnerRadius(gd.Float(value))
}

func (self Go) Angle() gd.Vector2 {
		return gd.Vector2(class(self).GetParam(7))
}

func (self Go) SetAngle(value gd.Vector2) {
	class(self).SetParam(7, value)
}

func (self Go) AngleMin() float64 {
		return float64(float64(class(self).GetParamMin(7)))
}

func (self Go) SetAngleMin(value float64) {
	class(self).SetParamMin(7, gd.Float(value))
}

func (self Go) AngleMax() float64 {
		return float64(float64(class(self).GetParamMax(7)))
}

func (self Go) SetAngleMax(value float64) {
	class(self).SetParamMax(7, gd.Float(value))
}

func (self Go) AngleCurve() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetParamTexture(7))
}

func (self Go) SetAngleCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(7, value)
}

func (self Go) InheritVelocityRatio() float64 {
		return float64(float64(class(self).GetInheritVelocityRatio()))
}

func (self Go) SetInheritVelocityRatio(value float64) {
	class(self).SetInheritVelocityRatio(gd.Float(value))
}

func (self Go) VelocityPivot() gd.Vector3 {
		return gd.Vector3(class(self).GetVelocityPivot())
}

func (self Go) SetVelocityPivot(value gd.Vector3) {
	class(self).SetVelocityPivot(value)
}

func (self Go) Direction() gd.Vector3 {
		return gd.Vector3(class(self).GetDirection())
}

func (self Go) SetDirection(value gd.Vector3) {
	class(self).SetDirection(value)
}

func (self Go) Spread() float64 {
		return float64(float64(class(self).GetSpread()))
}

func (self Go) SetSpread(value float64) {
	class(self).SetSpread(gd.Float(value))
}

func (self Go) Flatness() float64 {
		return float64(float64(class(self).GetFlatness()))
}

func (self Go) SetFlatness(value float64) {
	class(self).SetFlatness(gd.Float(value))
}

func (self Go) InitialVelocity() gd.Vector2 {
		return gd.Vector2(class(self).GetParam(0))
}

func (self Go) SetInitialVelocity(value gd.Vector2) {
	class(self).SetParam(0, value)
}

func (self Go) InitialVelocityMin() float64 {
		return float64(float64(class(self).GetParamMin(0)))
}

func (self Go) SetInitialVelocityMin(value float64) {
	class(self).SetParamMin(0, gd.Float(value))
}

func (self Go) InitialVelocityMax() float64 {
		return float64(float64(class(self).GetParamMax(0)))
}

func (self Go) SetInitialVelocityMax(value float64) {
	class(self).SetParamMax(0, gd.Float(value))
}

func (self Go) AngularVelocity() gd.Vector2 {
		return gd.Vector2(class(self).GetParam(1))
}

func (self Go) SetAngularVelocity(value gd.Vector2) {
	class(self).SetParam(1, value)
}

func (self Go) AngularVelocityMin() float64 {
		return float64(float64(class(self).GetParamMin(1)))
}

func (self Go) SetAngularVelocityMin(value float64) {
	class(self).SetParamMin(1, gd.Float(value))
}

func (self Go) AngularVelocityMax() float64 {
		return float64(float64(class(self).GetParamMax(1)))
}

func (self Go) SetAngularVelocityMax(value float64) {
	class(self).SetParamMax(1, gd.Float(value))
}

func (self Go) AngularVelocityCurve() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetParamTexture(1))
}

func (self Go) SetAngularVelocityCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(1, value)
}

func (self Go) DirectionalVelocity() gd.Vector2 {
		return gd.Vector2(class(self).GetParam(16))
}

func (self Go) SetDirectionalVelocity(value gd.Vector2) {
	class(self).SetParam(16, value)
}

func (self Go) DirectionalVelocityMin() float64 {
		return float64(float64(class(self).GetParamMin(16)))
}

func (self Go) SetDirectionalVelocityMin(value float64) {
	class(self).SetParamMin(16, gd.Float(value))
}

func (self Go) DirectionalVelocityMax() float64 {
		return float64(float64(class(self).GetParamMax(16)))
}

func (self Go) SetDirectionalVelocityMax(value float64) {
	class(self).SetParamMax(16, gd.Float(value))
}

func (self Go) DirectionalVelocityCurve() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetParamTexture(16))
}

func (self Go) SetDirectionalVelocityCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(16, value)
}

func (self Go) OrbitVelocity() gd.Vector2 {
		return gd.Vector2(class(self).GetParam(2))
}

func (self Go) SetOrbitVelocity(value gd.Vector2) {
	class(self).SetParam(2, value)
}

func (self Go) OrbitVelocityMin() float64 {
		return float64(float64(class(self).GetParamMin(2)))
}

func (self Go) SetOrbitVelocityMin(value float64) {
	class(self).SetParamMin(2, gd.Float(value))
}

func (self Go) OrbitVelocityMax() float64 {
		return float64(float64(class(self).GetParamMax(2)))
}

func (self Go) SetOrbitVelocityMax(value float64) {
	class(self).SetParamMax(2, gd.Float(value))
}

func (self Go) OrbitVelocityCurve() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetParamTexture(2))
}

func (self Go) SetOrbitVelocityCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(2, value)
}

func (self Go) RadialVelocity() gd.Vector2 {
		return gd.Vector2(class(self).GetParam(15))
}

func (self Go) SetRadialVelocity(value gd.Vector2) {
	class(self).SetParam(15, value)
}

func (self Go) RadialVelocityMin() float64 {
		return float64(float64(class(self).GetParamMin(15)))
}

func (self Go) SetRadialVelocityMin(value float64) {
	class(self).SetParamMin(15, gd.Float(value))
}

func (self Go) RadialVelocityMax() float64 {
		return float64(float64(class(self).GetParamMax(15)))
}

func (self Go) SetRadialVelocityMax(value float64) {
	class(self).SetParamMax(15, gd.Float(value))
}

func (self Go) RadialVelocityCurve() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetParamTexture(15))
}

func (self Go) SetRadialVelocityCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(15, value)
}

func (self Go) VelocityLimitCurve() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetVelocityLimitCurve())
}

func (self Go) SetVelocityLimitCurve(value gdclass.Texture2D) {
	class(self).SetVelocityLimitCurve(value)
}

func (self Go) Gravity() gd.Vector3 {
		return gd.Vector3(class(self).GetGravity())
}

func (self Go) SetGravity(value gd.Vector3) {
	class(self).SetGravity(value)
}

func (self Go) LinearAccel() gd.Vector2 {
		return gd.Vector2(class(self).GetParam(3))
}

func (self Go) SetLinearAccel(value gd.Vector2) {
	class(self).SetParam(3, value)
}

func (self Go) LinearAccelMin() float64 {
		return float64(float64(class(self).GetParamMin(3)))
}

func (self Go) SetLinearAccelMin(value float64) {
	class(self).SetParamMin(3, gd.Float(value))
}

func (self Go) LinearAccelMax() float64 {
		return float64(float64(class(self).GetParamMax(3)))
}

func (self Go) SetLinearAccelMax(value float64) {
	class(self).SetParamMax(3, gd.Float(value))
}

func (self Go) LinearAccelCurve() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetParamTexture(3))
}

func (self Go) SetLinearAccelCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(3, value)
}

func (self Go) RadialAccel() gd.Vector2 {
		return gd.Vector2(class(self).GetParam(4))
}

func (self Go) SetRadialAccel(value gd.Vector2) {
	class(self).SetParam(4, value)
}

func (self Go) RadialAccelMin() float64 {
		return float64(float64(class(self).GetParamMin(4)))
}

func (self Go) SetRadialAccelMin(value float64) {
	class(self).SetParamMin(4, gd.Float(value))
}

func (self Go) RadialAccelMax() float64 {
		return float64(float64(class(self).GetParamMax(4)))
}

func (self Go) SetRadialAccelMax(value float64) {
	class(self).SetParamMax(4, gd.Float(value))
}

func (self Go) RadialAccelCurve() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetParamTexture(4))
}

func (self Go) SetRadialAccelCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(4, value)
}

func (self Go) TangentialAccel() gd.Vector2 {
		return gd.Vector2(class(self).GetParam(5))
}

func (self Go) SetTangentialAccel(value gd.Vector2) {
	class(self).SetParam(5, value)
}

func (self Go) TangentialAccelMin() float64 {
		return float64(float64(class(self).GetParamMin(5)))
}

func (self Go) SetTangentialAccelMin(value float64) {
	class(self).SetParamMin(5, gd.Float(value))
}

func (self Go) TangentialAccelMax() float64 {
		return float64(float64(class(self).GetParamMax(5)))
}

func (self Go) SetTangentialAccelMax(value float64) {
	class(self).SetParamMax(5, gd.Float(value))
}

func (self Go) TangentialAccelCurve() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetParamTexture(5))
}

func (self Go) SetTangentialAccelCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(5, value)
}

func (self Go) Damping() gd.Vector2 {
		return gd.Vector2(class(self).GetParam(6))
}

func (self Go) SetDamping(value gd.Vector2) {
	class(self).SetParam(6, value)
}

func (self Go) DampingMin() float64 {
		return float64(float64(class(self).GetParamMin(6)))
}

func (self Go) SetDampingMin(value float64) {
	class(self).SetParamMin(6, gd.Float(value))
}

func (self Go) DampingMax() float64 {
		return float64(float64(class(self).GetParamMax(6)))
}

func (self Go) SetDampingMax(value float64) {
	class(self).SetParamMax(6, gd.Float(value))
}

func (self Go) DampingCurve() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetParamTexture(6))
}

func (self Go) SetDampingCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(6, value)
}

func (self Go) AttractorInteractionEnabled() bool {
		return bool(class(self).IsAttractorInteractionEnabled())
}

func (self Go) SetAttractorInteractionEnabled(value bool) {
	class(self).SetAttractorInteractionEnabled(value)
}

func (self Go) Scale() gd.Vector2 {
		return gd.Vector2(class(self).GetParam(8))
}

func (self Go) SetScale(value gd.Vector2) {
	class(self).SetParam(8, value)
}

func (self Go) ScaleMin() float64 {
		return float64(float64(class(self).GetParamMin(8)))
}

func (self Go) SetScaleMin(value float64) {
	class(self).SetParamMin(8, gd.Float(value))
}

func (self Go) ScaleMax() float64 {
		return float64(float64(class(self).GetParamMax(8)))
}

func (self Go) SetScaleMax(value float64) {
	class(self).SetParamMax(8, gd.Float(value))
}

func (self Go) ScaleCurve() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetParamTexture(8))
}

func (self Go) SetScaleCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(8, value)
}

func (self Go) ScaleOverVelocity() gd.Vector2 {
		return gd.Vector2(class(self).GetParam(17))
}

func (self Go) SetScaleOverVelocity(value gd.Vector2) {
	class(self).SetParam(17, value)
}

func (self Go) ScaleOverVelocityMin() float64 {
		return float64(float64(class(self).GetParamMin(17)))
}

func (self Go) SetScaleOverVelocityMin(value float64) {
	class(self).SetParamMin(17, gd.Float(value))
}

func (self Go) ScaleOverVelocityMax() float64 {
		return float64(float64(class(self).GetParamMax(17)))
}

func (self Go) SetScaleOverVelocityMax(value float64) {
	class(self).SetParamMax(17, gd.Float(value))
}

func (self Go) ScaleOverVelocityCurve() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetParamTexture(17))
}

func (self Go) SetScaleOverVelocityCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(17, value)
}

func (self Go) Color() gd.Color {
		return gd.Color(class(self).GetColor())
}

func (self Go) SetColor(value gd.Color) {
	class(self).SetColor(value)
}

func (self Go) ColorRamp() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetColorRamp())
}

func (self Go) SetColorRamp(value gdclass.Texture2D) {
	class(self).SetColorRamp(value)
}

func (self Go) ColorInitialRamp() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetColorInitialRamp())
}

func (self Go) SetColorInitialRamp(value gdclass.Texture2D) {
	class(self).SetColorInitialRamp(value)
}

func (self Go) AlphaCurve() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetAlphaCurve())
}

func (self Go) SetAlphaCurve(value gdclass.Texture2D) {
	class(self).SetAlphaCurve(value)
}

func (self Go) EmissionCurve() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetEmissionCurve())
}

func (self Go) SetEmissionCurve(value gdclass.Texture2D) {
	class(self).SetEmissionCurve(value)
}

func (self Go) HueVariation() gd.Vector2 {
		return gd.Vector2(class(self).GetParam(9))
}

func (self Go) SetHueVariation(value gd.Vector2) {
	class(self).SetParam(9, value)
}

func (self Go) HueVariationMin() float64 {
		return float64(float64(class(self).GetParamMin(9)))
}

func (self Go) SetHueVariationMin(value float64) {
	class(self).SetParamMin(9, gd.Float(value))
}

func (self Go) HueVariationMax() float64 {
		return float64(float64(class(self).GetParamMax(9)))
}

func (self Go) SetHueVariationMax(value float64) {
	class(self).SetParamMax(9, gd.Float(value))
}

func (self Go) HueVariationCurve() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetParamTexture(9))
}

func (self Go) SetHueVariationCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(9, value)
}

func (self Go) AnimSpeed() gd.Vector2 {
		return gd.Vector2(class(self).GetParam(10))
}

func (self Go) SetAnimSpeed(value gd.Vector2) {
	class(self).SetParam(10, value)
}

func (self Go) AnimSpeedMin() float64 {
		return float64(float64(class(self).GetParamMin(10)))
}

func (self Go) SetAnimSpeedMin(value float64) {
	class(self).SetParamMin(10, gd.Float(value))
}

func (self Go) AnimSpeedMax() float64 {
		return float64(float64(class(self).GetParamMax(10)))
}

func (self Go) SetAnimSpeedMax(value float64) {
	class(self).SetParamMax(10, gd.Float(value))
}

func (self Go) AnimSpeedCurve() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetParamTexture(10))
}

func (self Go) SetAnimSpeedCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(10, value)
}

func (self Go) AnimOffset() gd.Vector2 {
		return gd.Vector2(class(self).GetParam(11))
}

func (self Go) SetAnimOffset(value gd.Vector2) {
	class(self).SetParam(11, value)
}

func (self Go) AnimOffsetMin() float64 {
		return float64(float64(class(self).GetParamMin(11)))
}

func (self Go) SetAnimOffsetMin(value float64) {
	class(self).SetParamMin(11, gd.Float(value))
}

func (self Go) AnimOffsetMax() float64 {
		return float64(float64(class(self).GetParamMax(11)))
}

func (self Go) SetAnimOffsetMax(value float64) {
	class(self).SetParamMax(11, gd.Float(value))
}

func (self Go) AnimOffsetCurve() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetParamTexture(11))
}

func (self Go) SetAnimOffsetCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(11, value)
}

func (self Go) TurbulenceEnabled() bool {
		return bool(class(self).GetTurbulenceEnabled())
}

func (self Go) SetTurbulenceEnabled(value bool) {
	class(self).SetTurbulenceEnabled(value)
}

func (self Go) TurbulenceNoiseStrength() float64 {
		return float64(float64(class(self).GetTurbulenceNoiseStrength()))
}

func (self Go) SetTurbulenceNoiseStrength(value float64) {
	class(self).SetTurbulenceNoiseStrength(gd.Float(value))
}

func (self Go) TurbulenceNoiseScale() float64 {
		return float64(float64(class(self).GetTurbulenceNoiseScale()))
}

func (self Go) SetTurbulenceNoiseScale(value float64) {
	class(self).SetTurbulenceNoiseScale(gd.Float(value))
}

func (self Go) TurbulenceNoiseSpeed() gd.Vector3 {
		return gd.Vector3(class(self).GetTurbulenceNoiseSpeed())
}

func (self Go) SetTurbulenceNoiseSpeed(value gd.Vector3) {
	class(self).SetTurbulenceNoiseSpeed(value)
}

func (self Go) TurbulenceNoiseSpeedRandom() float64 {
		return float64(float64(class(self).GetTurbulenceNoiseSpeedRandom()))
}

func (self Go) SetTurbulenceNoiseSpeedRandom(value float64) {
	class(self).SetTurbulenceNoiseSpeedRandom(gd.Float(value))
}

func (self Go) TurbulenceInfluence() gd.Vector2 {
		return gd.Vector2(class(self).GetParam(13))
}

func (self Go) SetTurbulenceInfluence(value gd.Vector2) {
	class(self).SetParam(13, value)
}

func (self Go) TurbulenceInfluenceMin() float64 {
		return float64(float64(class(self).GetParamMin(13)))
}

func (self Go) SetTurbulenceInfluenceMin(value float64) {
	class(self).SetParamMin(13, gd.Float(value))
}

func (self Go) TurbulenceInfluenceMax() float64 {
		return float64(float64(class(self).GetParamMax(13)))
}

func (self Go) SetTurbulenceInfluenceMax(value float64) {
	class(self).SetParamMax(13, gd.Float(value))
}

func (self Go) TurbulenceInitialDisplacement() gd.Vector2 {
		return gd.Vector2(class(self).GetParam(14))
}

func (self Go) SetTurbulenceInitialDisplacement(value gd.Vector2) {
	class(self).SetParam(14, value)
}

func (self Go) TurbulenceInitialDisplacementMin() float64 {
		return float64(float64(class(self).GetParamMin(14)))
}

func (self Go) SetTurbulenceInitialDisplacementMin(value float64) {
	class(self).SetParamMin(14, gd.Float(value))
}

func (self Go) TurbulenceInitialDisplacementMax() float64 {
		return float64(float64(class(self).GetParamMax(14)))
}

func (self Go) SetTurbulenceInitialDisplacementMax(value float64) {
	class(self).SetParamMax(14, gd.Float(value))
}

func (self Go) TurbulenceInfluenceOverLife() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetParamTexture(12))
}

func (self Go) SetTurbulenceInfluenceOverLife(value gdclass.Texture2D) {
	class(self).SetParamTexture(12, value)
}

func (self Go) CollisionMode() classdb.ParticleProcessMaterialCollisionMode {
		return classdb.ParticleProcessMaterialCollisionMode(class(self).GetCollisionMode())
}

func (self Go) SetCollisionMode(value classdb.ParticleProcessMaterialCollisionMode) {
	class(self).SetCollisionMode(value)
}

func (self Go) CollisionFriction() float64 {
		return float64(float64(class(self).GetCollisionFriction()))
}

func (self Go) SetCollisionFriction(value float64) {
	class(self).SetCollisionFriction(gd.Float(value))
}

func (self Go) CollisionBounce() float64 {
		return float64(float64(class(self).GetCollisionBounce()))
}

func (self Go) SetCollisionBounce(value float64) {
	class(self).SetCollisionBounce(gd.Float(value))
}

func (self Go) CollisionUseScale() bool {
		return bool(class(self).IsCollisionUsingScale())
}

func (self Go) SetCollisionUseScale(value bool) {
	class(self).SetCollisionUseScale(value)
}

func (self Go) SubEmitterMode() classdb.ParticleProcessMaterialSubEmitterMode {
		return classdb.ParticleProcessMaterialSubEmitterMode(class(self).GetSubEmitterMode())
}

func (self Go) SetSubEmitterMode(value classdb.ParticleProcessMaterialSubEmitterMode) {
	class(self).SetSubEmitterMode(value)
}

func (self Go) SubEmitterFrequency() float64 {
		return float64(float64(class(self).GetSubEmitterFrequency()))
}

func (self Go) SetSubEmitterFrequency(value float64) {
	class(self).SetSubEmitterFrequency(gd.Float(value))
}

func (self Go) SubEmitterAmountAtEnd() int {
		return int(int(class(self).GetSubEmitterAmountAtEnd()))
}

func (self Go) SetSubEmitterAmountAtEnd(value int) {
	class(self).SetSubEmitterAmountAtEnd(gd.Int(value))
}

func (self Go) SubEmitterAmountAtCollision() int {
		return int(int(class(self).GetSubEmitterAmountAtCollision()))
}

func (self Go) SetSubEmitterAmountAtCollision(value int) {
	class(self).SetSubEmitterAmountAtCollision(gd.Int(value))
}

func (self Go) SubEmitterKeepVelocity() bool {
		return bool(class(self).GetSubEmitterKeepVelocity())
}

func (self Go) SetSubEmitterKeepVelocity(value bool) {
	class(self).SetSubEmitterKeepVelocity(value)
}

//go:nosplit
func (self class) SetDirection(degrees gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDirection() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInheritVelocityRatio(ratio gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_inherit_velocity_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInheritVelocityRatio() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_inherit_velocity_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSpread(degrees gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_spread, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSpread() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_spread, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFlatness(amount gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_flatness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFlatness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_flatness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the minimum and maximum values of the given [param param].
The [code]x[/code] component of the argument vector corresponds to minimum and the [code]y[/code] component corresponds to maximum.
*/
//go:nosplit
func (self class) SetParam(param classdb.ParticleProcessMaterialParameter, value gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the minimum and maximum values of the given [param param] as a vector.
The [code]x[/code] component of the returned vector corresponds to minimum and the [code]y[/code] component corresponds to maximum.
*/
//go:nosplit
func (self class) GetParam(param classdb.ParticleProcessMaterialParameter) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the minimum value range for the given parameter.
*/
//go:nosplit
func (self class) SetParamMin(param classdb.ParticleProcessMaterialParameter, value gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_param_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the minimum value range for the given parameter.
*/
//go:nosplit
func (self class) GetParamMin(param classdb.ParticleProcessMaterialParameter) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_param_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the maximum value range for the given parameter.
*/
//go:nosplit
func (self class) SetParamMax(param classdb.ParticleProcessMaterialParameter, value gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_param_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the maximum value range for the given parameter.
*/
//go:nosplit
func (self class) GetParamMax(param classdb.ParticleProcessMaterialParameter) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_param_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [Texture2D] for the specified [enum Parameter].
*/
//go:nosplit
func (self class) SetParamTexture(param classdb.ParticleProcessMaterialParameter, texture gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, discreet.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_param_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Texture2D] used by the specified parameter.
*/
//go:nosplit
func (self class) GetParamTexture(param classdb.ParticleProcessMaterialParameter) gdclass.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_param_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColor(color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColorRamp(ramp gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(ramp[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_color_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColorRamp() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_color_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlphaCurve(curve gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_alpha_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlphaCurve() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_alpha_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionCurve(curve gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionCurve() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColorInitialRamp(ramp gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(ramp[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_color_initial_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColorInitialRamp() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_color_initial_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVelocityLimitCurve(curve gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_velocity_limit_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVelocityLimitCurve() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_velocity_limit_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
If [code]true[/code], enables the specified particle flag. See [enum ParticleFlags] for options.
*/
//go:nosplit
func (self class) SetParticleFlag(particle_flag classdb.ParticleProcessMaterialParticleFlags, enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, particle_flag)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_particle_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the specified particle flag is enabled. See [enum ParticleFlags] for options.
*/
//go:nosplit
func (self class) GetParticleFlag(particle_flag classdb.ParticleProcessMaterialParticleFlags) bool {
	var frame = callframe.New()
	callframe.Arg(frame, particle_flag)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_particle_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVelocityPivot(pivot gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, pivot)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_velocity_pivot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVelocityPivot() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_velocity_pivot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionShape(shape classdb.ParticleProcessMaterialEmissionShape)  {
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionShape() classdb.ParticleProcessMaterialEmissionShape {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ParticleProcessMaterialEmissionShape](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionSphereRadius(radius gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_sphere_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionSphereRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_sphere_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionBoxExtents(extents gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, extents)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_box_extents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionBoxExtents() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_box_extents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionPointTexture(texture gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_point_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionPointTexture() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_point_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionNormalTexture(texture gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_normal_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionNormalTexture() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_normal_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionColorTexture(texture gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_color_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionColorTexture() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_color_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionPointCount(point_count gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, point_count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_point_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionPointCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_point_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionRingAxis(axis gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, axis)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_ring_axis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionRingAxis() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_ring_axis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionRingHeight(height gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_ring_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionRingHeight() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_ring_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionRingRadius(radius gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_ring_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionRingRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_ring_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionRingInnerRadius(inner_radius gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, inner_radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_ring_inner_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionRingInnerRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_ring_inner_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionShapeOffset(emission_shape_offset gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, emission_shape_offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_shape_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionShapeOffset() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_shape_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionShapeScale(emission_shape_scale gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, emission_shape_scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_shape_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionShapeScale() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_shape_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetTurbulenceEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_turbulence_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTurbulenceEnabled(turbulence_enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, turbulence_enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_turbulence_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTurbulenceNoiseStrength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_turbulence_noise_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTurbulenceNoiseStrength(turbulence_noise_strength gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, turbulence_noise_strength)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_turbulence_noise_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTurbulenceNoiseScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_turbulence_noise_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTurbulenceNoiseScale(turbulence_noise_scale gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, turbulence_noise_scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_turbulence_noise_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTurbulenceNoiseSpeedRandom() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_turbulence_noise_speed_random, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTurbulenceNoiseSpeedRandom(turbulence_noise_speed_random gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, turbulence_noise_speed_random)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_turbulence_noise_speed_random, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTurbulenceNoiseSpeed() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_turbulence_noise_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTurbulenceNoiseSpeed(turbulence_noise_speed gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, turbulence_noise_speed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_turbulence_noise_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGravity() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGravity(accel_vec gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, accel_vec)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetLifetimeRandomness(randomness gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, randomness)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_lifetime_randomness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLifetimeRandomness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_lifetime_randomness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetSubEmitterMode() classdb.ParticleProcessMaterialSubEmitterMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ParticleProcessMaterialSubEmitterMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_sub_emitter_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSubEmitterMode(mode classdb.ParticleProcessMaterialSubEmitterMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_sub_emitter_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSubEmitterFrequency() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_sub_emitter_frequency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSubEmitterFrequency(hz gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, hz)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_sub_emitter_frequency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSubEmitterAmountAtEnd() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_sub_emitter_amount_at_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSubEmitterAmountAtEnd(amount gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_sub_emitter_amount_at_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSubEmitterAmountAtCollision() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_sub_emitter_amount_at_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSubEmitterAmountAtCollision(amount gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_sub_emitter_amount_at_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSubEmitterKeepVelocity() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_sub_emitter_keep_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSubEmitterKeepVelocity(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_sub_emitter_keep_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetAttractorInteractionEnabled(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_attractor_interaction_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAttractorInteractionEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_is_attractor_interaction_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionMode(mode classdb.ParticleProcessMaterialCollisionMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_collision_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionMode() classdb.ParticleProcessMaterialCollisionMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ParticleProcessMaterialCollisionMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_collision_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionUseScale(radius bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_collision_use_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCollisionUsingScale() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_is_collision_using_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionFriction(friction gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, friction)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_collision_friction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionFriction() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_collision_friction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionBounce(bounce gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, bounce)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_collision_bounce, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionBounce() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_collision_bounce, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsParticleProcessMaterial() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsParticleProcessMaterial() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("ParticleProcessMaterial", func(ptr gd.Object) any { return classdb.ParticleProcessMaterial(ptr) })}
type Parameter = classdb.ParticleProcessMaterialParameter

const (
/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set initial velocity properties.*/
	ParamInitialLinearVelocity Parameter = 0
/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set angular velocity properties.*/
	ParamAngularVelocity Parameter = 1
/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set orbital velocity properties.*/
	ParamOrbitVelocity Parameter = 2
/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set linear acceleration properties.*/
	ParamLinearAccel Parameter = 3
/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set radial acceleration properties.*/
	ParamRadialAccel Parameter = 4
/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set tangential acceleration properties.*/
	ParamTangentialAccel Parameter = 5
/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set damping properties.*/
	ParamDamping Parameter = 6
/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set angle properties.*/
	ParamAngle Parameter = 7
/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set scale properties.*/
	ParamScale Parameter = 8
/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set hue variation properties.*/
	ParamHueVariation Parameter = 9
/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set animation speed properties.*/
	ParamAnimSpeed Parameter = 10
/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set animation offset properties.*/
	ParamAnimOffset Parameter = 11
/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set radial velocity properties.*/
	ParamRadialVelocity Parameter = 15
/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set directional velocity properties.*/
	ParamDirectionalVelocity Parameter = 16
/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set scale over velocity properties.*/
	ParamScaleOverVelocity Parameter = 17
/*Represents the size of the [enum Parameter] enum.*/
	ParamMax Parameter = 18
/*Use with [method set_param_min] and [method set_param_max] to set the turbulence minimum und maximum influence on each particles velocity.*/
	ParamTurbVelInfluence Parameter = 13
/*Use with [method set_param_min] and [method set_param_max] to set the turbulence minimum and maximum displacement of the particles spawn position.*/
	ParamTurbInitDisplacement Parameter = 14
/*Use with [method set_param_texture] to set the turbulence influence over the particles life time.*/
	ParamTurbInfluenceOverLife Parameter = 12
)
type ParticleFlags = classdb.ParticleProcessMaterialParticleFlags

const (
/*Use with [method set_particle_flag] to set [member particle_flag_align_y].*/
	ParticleFlagAlignYToVelocity ParticleFlags = 0
/*Use with [method set_particle_flag] to set [member particle_flag_rotate_y].*/
	ParticleFlagRotateY ParticleFlags = 1
/*Use with [method set_particle_flag] to set [member particle_flag_disable_z].*/
	ParticleFlagDisableZ ParticleFlags = 2
	ParticleFlagDampingAsFriction ParticleFlags = 3
/*Represents the size of the [enum ParticleFlags] enum.*/
	ParticleFlagMax ParticleFlags = 4
)
type EmissionShape = classdb.ParticleProcessMaterialEmissionShape

const (
/*All particles will be emitted from a single point.*/
	EmissionShapePoint EmissionShape = 0
/*Particles will be emitted in the volume of a sphere.*/
	EmissionShapeSphere EmissionShape = 1
/*Particles will be emitted on the surface of a sphere.*/
	EmissionShapeSphereSurface EmissionShape = 2
/*Particles will be emitted in the volume of a box.*/
	EmissionShapeBox EmissionShape = 3
/*Particles will be emitted at a position determined by sampling a random point on the [member emission_point_texture]. Particle color will be modulated by [member emission_color_texture].*/
	EmissionShapePoints EmissionShape = 4
/*Particles will be emitted at a position determined by sampling a random point on the [member emission_point_texture]. Particle velocity and rotation will be set based on [member emission_normal_texture]. Particle color will be modulated by [member emission_color_texture].*/
	EmissionShapeDirectedPoints EmissionShape = 5
/*Particles will be emitted in a ring or cylinder.*/
	EmissionShapeRing EmissionShape = 6
/*Represents the size of the [enum EmissionShape] enum.*/
	EmissionShapeMax EmissionShape = 7
)
type SubEmitterMode = classdb.ParticleProcessMaterialSubEmitterMode

const (
	SubEmitterDisabled SubEmitterMode = 0
	SubEmitterConstant SubEmitterMode = 1
	SubEmitterAtEnd SubEmitterMode = 2
	SubEmitterAtCollision SubEmitterMode = 3
/*Represents the size of the [enum SubEmitterMode] enum.*/
	SubEmitterMax SubEmitterMode = 4
)
type CollisionMode = classdb.ParticleProcessMaterialCollisionMode

const (
/*No collision for particles. Particles will go through [GPUParticlesCollision3D] nodes.*/
	CollisionDisabled CollisionMode = 0
/*[RigidBody3D]-style collision for particles using [GPUParticlesCollision3D] nodes.*/
	CollisionRigid CollisionMode = 1
/*Hide particles instantly when colliding with a [GPUParticlesCollision3D] node. This can be combined with a subemitter that uses the [constant COLLISION_RIGID] collision mode to "replace" the parent particle with the subemitter on impact.*/
	CollisionHideOnContact CollisionMode = 2
/*Represents the size of the [enum CollisionMode] enum.*/
	CollisionMax CollisionMode = 3
)
