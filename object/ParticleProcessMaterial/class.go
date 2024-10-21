package ParticleProcessMaterial

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
[ParticleProcessMaterial] defines particle properties and behavior. It is used in the [code]process_material[/code] of the [GPUParticles2D] and [GPUParticles3D] nodes. Some of this material's properties are applied to each particle when emitted, while others can have a [CurveTexture] or a [GradientTexture1D] applied to vary numerical or color values over the lifetime of the particle.

*/
type Simple [1]classdb.ParticleProcessMaterial
func (self Simple) SetDirection(degrees gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDirection(degrees)
}
func (self Simple) GetDirection() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetDirection())
}
func (self Simple) SetInheritVelocityRatio(ratio float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInheritVelocityRatio(gd.Float(ratio))
}
func (self Simple) GetInheritVelocityRatio() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetInheritVelocityRatio()))
}
func (self Simple) SetSpread(degrees float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSpread(gd.Float(degrees))
}
func (self Simple) GetSpread() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSpread()))
}
func (self Simple) SetFlatness(amount float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFlatness(gd.Float(amount))
}
func (self Simple) GetFlatness() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFlatness()))
}
func (self Simple) SetParam(param classdb.ParticleProcessMaterialParameter, value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParam(param, value)
}
func (self Simple) GetParam(param classdb.ParticleProcessMaterialParameter) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetParam(param))
}
func (self Simple) SetParamMin(param classdb.ParticleProcessMaterialParameter, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParamMin(param, gd.Float(value))
}
func (self Simple) GetParamMin(param classdb.ParticleProcessMaterialParameter) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetParamMin(param)))
}
func (self Simple) SetParamMax(param classdb.ParticleProcessMaterialParameter, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParamMax(param, gd.Float(value))
}
func (self Simple) GetParamMax(param classdb.ParticleProcessMaterialParameter) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetParamMax(param)))
}
func (self Simple) SetParamTexture(param classdb.ParticleProcessMaterialParameter, texture [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParamTexture(param, texture)
}
func (self Simple) GetParamTexture(param classdb.ParticleProcessMaterialParameter) [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetParamTexture(gc, param))
}
func (self Simple) SetColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColor(color)
}
func (self Simple) GetColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetColor())
}
func (self Simple) SetColorRamp(ramp [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColorRamp(ramp)
}
func (self Simple) GetColorRamp() [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetColorRamp(gc))
}
func (self Simple) SetAlphaCurve(curve [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAlphaCurve(curve)
}
func (self Simple) GetAlphaCurve() [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetAlphaCurve(gc))
}
func (self Simple) SetEmissionCurve(curve [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmissionCurve(curve)
}
func (self Simple) GetEmissionCurve() [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetEmissionCurve(gc))
}
func (self Simple) SetColorInitialRamp(ramp [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColorInitialRamp(ramp)
}
func (self Simple) GetColorInitialRamp() [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetColorInitialRamp(gc))
}
func (self Simple) SetVelocityLimitCurve(curve [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVelocityLimitCurve(curve)
}
func (self Simple) GetVelocityLimitCurve() [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetVelocityLimitCurve(gc))
}
func (self Simple) SetParticleFlag(particle_flag classdb.ParticleProcessMaterialParticleFlags, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParticleFlag(particle_flag, enable)
}
func (self Simple) GetParticleFlag(particle_flag classdb.ParticleProcessMaterialParticleFlags) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetParticleFlag(particle_flag))
}
func (self Simple) SetVelocityPivot(pivot gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVelocityPivot(pivot)
}
func (self Simple) GetVelocityPivot() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetVelocityPivot())
}
func (self Simple) SetEmissionShape(shape classdb.ParticleProcessMaterialEmissionShape) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmissionShape(shape)
}
func (self Simple) GetEmissionShape() classdb.ParticleProcessMaterialEmissionShape {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ParticleProcessMaterialEmissionShape(Expert(self).GetEmissionShape())
}
func (self Simple) SetEmissionSphereRadius(radius float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmissionSphereRadius(gd.Float(radius))
}
func (self Simple) GetEmissionSphereRadius() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetEmissionSphereRadius()))
}
func (self Simple) SetEmissionBoxExtents(extents gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmissionBoxExtents(extents)
}
func (self Simple) GetEmissionBoxExtents() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetEmissionBoxExtents())
}
func (self Simple) SetEmissionPointTexture(texture [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmissionPointTexture(texture)
}
func (self Simple) GetEmissionPointTexture() [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetEmissionPointTexture(gc))
}
func (self Simple) SetEmissionNormalTexture(texture [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmissionNormalTexture(texture)
}
func (self Simple) GetEmissionNormalTexture() [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetEmissionNormalTexture(gc))
}
func (self Simple) SetEmissionColorTexture(texture [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmissionColorTexture(texture)
}
func (self Simple) GetEmissionColorTexture() [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetEmissionColorTexture(gc))
}
func (self Simple) SetEmissionPointCount(point_count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmissionPointCount(gd.Int(point_count))
}
func (self Simple) GetEmissionPointCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetEmissionPointCount()))
}
func (self Simple) SetEmissionRingAxis(axis gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmissionRingAxis(axis)
}
func (self Simple) GetEmissionRingAxis() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetEmissionRingAxis())
}
func (self Simple) SetEmissionRingHeight(height float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmissionRingHeight(gd.Float(height))
}
func (self Simple) GetEmissionRingHeight() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetEmissionRingHeight()))
}
func (self Simple) SetEmissionRingRadius(radius float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmissionRingRadius(gd.Float(radius))
}
func (self Simple) GetEmissionRingRadius() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetEmissionRingRadius()))
}
func (self Simple) SetEmissionRingInnerRadius(inner_radius float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmissionRingInnerRadius(gd.Float(inner_radius))
}
func (self Simple) GetEmissionRingInnerRadius() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetEmissionRingInnerRadius()))
}
func (self Simple) SetEmissionShapeOffset(emission_shape_offset gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmissionShapeOffset(emission_shape_offset)
}
func (self Simple) GetEmissionShapeOffset() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetEmissionShapeOffset())
}
func (self Simple) SetEmissionShapeScale(emission_shape_scale gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmissionShapeScale(emission_shape_scale)
}
func (self Simple) GetEmissionShapeScale() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetEmissionShapeScale())
}
func (self Simple) GetTurbulenceEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetTurbulenceEnabled())
}
func (self Simple) SetTurbulenceEnabled(turbulence_enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTurbulenceEnabled(turbulence_enabled)
}
func (self Simple) GetTurbulenceNoiseStrength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTurbulenceNoiseStrength()))
}
func (self Simple) SetTurbulenceNoiseStrength(turbulence_noise_strength float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTurbulenceNoiseStrength(gd.Float(turbulence_noise_strength))
}
func (self Simple) GetTurbulenceNoiseScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTurbulenceNoiseScale()))
}
func (self Simple) SetTurbulenceNoiseScale(turbulence_noise_scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTurbulenceNoiseScale(gd.Float(turbulence_noise_scale))
}
func (self Simple) GetTurbulenceNoiseSpeedRandom() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTurbulenceNoiseSpeedRandom()))
}
func (self Simple) SetTurbulenceNoiseSpeedRandom(turbulence_noise_speed_random float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTurbulenceNoiseSpeedRandom(gd.Float(turbulence_noise_speed_random))
}
func (self Simple) GetTurbulenceNoiseSpeed() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetTurbulenceNoiseSpeed())
}
func (self Simple) SetTurbulenceNoiseSpeed(turbulence_noise_speed gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTurbulenceNoiseSpeed(turbulence_noise_speed)
}
func (self Simple) GetGravity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetGravity())
}
func (self Simple) SetGravity(accel_vec gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGravity(accel_vec)
}
func (self Simple) SetLifetimeRandomness(randomness float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLifetimeRandomness(gd.Float(randomness))
}
func (self Simple) GetLifetimeRandomness() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetLifetimeRandomness()))
}
func (self Simple) GetSubEmitterMode() classdb.ParticleProcessMaterialSubEmitterMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ParticleProcessMaterialSubEmitterMode(Expert(self).GetSubEmitterMode())
}
func (self Simple) SetSubEmitterMode(mode classdb.ParticleProcessMaterialSubEmitterMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSubEmitterMode(mode)
}
func (self Simple) GetSubEmitterFrequency() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSubEmitterFrequency()))
}
func (self Simple) SetSubEmitterFrequency(hz float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSubEmitterFrequency(gd.Float(hz))
}
func (self Simple) GetSubEmitterAmountAtEnd() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSubEmitterAmountAtEnd()))
}
func (self Simple) SetSubEmitterAmountAtEnd(amount int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSubEmitterAmountAtEnd(gd.Int(amount))
}
func (self Simple) GetSubEmitterAmountAtCollision() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSubEmitterAmountAtCollision()))
}
func (self Simple) SetSubEmitterAmountAtCollision(amount int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSubEmitterAmountAtCollision(gd.Int(amount))
}
func (self Simple) GetSubEmitterKeepVelocity() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetSubEmitterKeepVelocity())
}
func (self Simple) SetSubEmitterKeepVelocity(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSubEmitterKeepVelocity(enable)
}
func (self Simple) SetAttractorInteractionEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAttractorInteractionEnabled(enabled)
}
func (self Simple) IsAttractorInteractionEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAttractorInteractionEnabled())
}
func (self Simple) SetCollisionMode(mode classdb.ParticleProcessMaterialCollisionMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCollisionMode(mode)
}
func (self Simple) GetCollisionMode() classdb.ParticleProcessMaterialCollisionMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ParticleProcessMaterialCollisionMode(Expert(self).GetCollisionMode())
}
func (self Simple) SetCollisionUseScale(radius bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCollisionUseScale(radius)
}
func (self Simple) IsCollisionUsingScale() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsCollisionUsingScale())
}
func (self Simple) SetCollisionFriction(friction float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCollisionFriction(gd.Float(friction))
}
func (self Simple) GetCollisionFriction() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetCollisionFriction()))
}
func (self Simple) SetCollisionBounce(bounce float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCollisionBounce(gd.Float(bounce))
}
func (self Simple) GetCollisionBounce() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetCollisionBounce()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ParticleProcessMaterial
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetDirection(degrees gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDirection() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInheritVelocityRatio(ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_inherit_velocity_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInheritVelocityRatio() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_inherit_velocity_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSpread(degrees gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_spread, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSpread() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_spread, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFlatness(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_flatness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFlatness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_flatness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the minimum and maximum values of the given [param param] as a vector.
The [code]x[/code] component of the returned vector corresponds to minimum and the [code]y[/code] component corresponds to maximum.
*/
//go:nosplit
func (self class) GetParam(param classdb.ParticleProcessMaterialParameter) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the minimum value range for the given parameter.
*/
//go:nosplit
func (self class) SetParamMin(param classdb.ParticleProcessMaterialParameter, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_param_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the minimum value range for the given parameter.
*/
//go:nosplit
func (self class) GetParamMin(param classdb.ParticleProcessMaterialParameter) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_param_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the maximum value range for the given parameter.
*/
//go:nosplit
func (self class) SetParamMax(param classdb.ParticleProcessMaterialParameter, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_param_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the maximum value range for the given parameter.
*/
//go:nosplit
func (self class) GetParamMax(param classdb.ParticleProcessMaterialParameter) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_param_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [Texture2D] for the specified [enum Parameter].
*/
//go:nosplit
func (self class) SetParamTexture(param classdb.ParticleProcessMaterialParameter, texture object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_param_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Texture2D] used by the specified parameter.
*/
//go:nosplit
func (self class) GetParamTexture(ctx gd.Lifetime, param classdb.ParticleProcessMaterialParameter) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_param_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColorRamp(ramp object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(ramp[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_color_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColorRamp(ctx gd.Lifetime) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_color_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlphaCurve(curve object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(curve[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_alpha_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlphaCurve(ctx gd.Lifetime) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_alpha_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionCurve(curve object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(curve[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_emission_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionCurve(ctx gd.Lifetime) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_emission_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColorInitialRamp(ramp object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(ramp[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_color_initial_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColorInitialRamp(ctx gd.Lifetime) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_color_initial_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVelocityLimitCurve(curve object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(curve[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_velocity_limit_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVelocityLimitCurve(ctx gd.Lifetime) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_velocity_limit_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
If [code]true[/code], enables the specified particle flag. See [enum ParticleFlags] for options.
*/
//go:nosplit
func (self class) SetParticleFlag(particle_flag classdb.ParticleProcessMaterialParticleFlags, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particle_flag)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_particle_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the specified particle flag is enabled. See [enum ParticleFlags] for options.
*/
//go:nosplit
func (self class) GetParticleFlag(particle_flag classdb.ParticleProcessMaterialParticleFlags) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particle_flag)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_particle_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVelocityPivot(pivot gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pivot)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_velocity_pivot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVelocityPivot() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_velocity_pivot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionShape(shape classdb.ParticleProcessMaterialEmissionShape)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_emission_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionShape() classdb.ParticleProcessMaterialEmissionShape {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ParticleProcessMaterialEmissionShape](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_emission_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionSphereRadius(radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_emission_sphere_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionSphereRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_emission_sphere_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionBoxExtents(extents gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, extents)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_emission_box_extents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionBoxExtents() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_emission_box_extents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionPointTexture(texture object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_emission_point_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionPointTexture(ctx gd.Lifetime) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_emission_point_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionNormalTexture(texture object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_emission_normal_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionNormalTexture(ctx gd.Lifetime) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_emission_normal_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionColorTexture(texture object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_emission_color_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionColorTexture(ctx gd.Lifetime) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_emission_color_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionPointCount(point_count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point_count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_emission_point_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionPointCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_emission_point_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionRingAxis(axis gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, axis)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_emission_ring_axis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionRingAxis() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_emission_ring_axis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionRingHeight(height gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_emission_ring_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionRingHeight() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_emission_ring_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionRingRadius(radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_emission_ring_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionRingRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_emission_ring_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionRingInnerRadius(inner_radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, inner_radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_emission_ring_inner_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionRingInnerRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_emission_ring_inner_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionShapeOffset(emission_shape_offset gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, emission_shape_offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_emission_shape_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionShapeOffset() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_emission_shape_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionShapeScale(emission_shape_scale gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, emission_shape_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_emission_shape_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionShapeScale() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_emission_shape_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetTurbulenceEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_turbulence_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTurbulenceEnabled(turbulence_enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, turbulence_enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_turbulence_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTurbulenceNoiseStrength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_turbulence_noise_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTurbulenceNoiseStrength(turbulence_noise_strength gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, turbulence_noise_strength)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_turbulence_noise_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTurbulenceNoiseScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_turbulence_noise_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTurbulenceNoiseScale(turbulence_noise_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, turbulence_noise_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_turbulence_noise_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTurbulenceNoiseSpeedRandom() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_turbulence_noise_speed_random, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTurbulenceNoiseSpeedRandom(turbulence_noise_speed_random gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, turbulence_noise_speed_random)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_turbulence_noise_speed_random, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTurbulenceNoiseSpeed() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_turbulence_noise_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTurbulenceNoiseSpeed(turbulence_noise_speed gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, turbulence_noise_speed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_turbulence_noise_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGravity() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGravity(accel_vec gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, accel_vec)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetLifetimeRandomness(randomness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, randomness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_lifetime_randomness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLifetimeRandomness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_lifetime_randomness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetSubEmitterMode() classdb.ParticleProcessMaterialSubEmitterMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ParticleProcessMaterialSubEmitterMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_sub_emitter_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSubEmitterMode(mode classdb.ParticleProcessMaterialSubEmitterMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_sub_emitter_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSubEmitterFrequency() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_sub_emitter_frequency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSubEmitterFrequency(hz gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hz)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_sub_emitter_frequency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSubEmitterAmountAtEnd() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_sub_emitter_amount_at_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSubEmitterAmountAtEnd(amount gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_sub_emitter_amount_at_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSubEmitterAmountAtCollision() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_sub_emitter_amount_at_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSubEmitterAmountAtCollision(amount gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_sub_emitter_amount_at_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSubEmitterKeepVelocity() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_sub_emitter_keep_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSubEmitterKeepVelocity(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_sub_emitter_keep_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetAttractorInteractionEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_attractor_interaction_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAttractorInteractionEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_is_attractor_interaction_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionMode(mode classdb.ParticleProcessMaterialCollisionMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_collision_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionMode() classdb.ParticleProcessMaterialCollisionMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ParticleProcessMaterialCollisionMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_collision_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionUseScale(radius bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_collision_use_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCollisionUsingScale() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_is_collision_using_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionFriction(friction gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, friction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_collision_friction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionFriction() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_collision_friction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionBounce(bounce gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bounce)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_set_collision_bounce, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionBounce() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ParticleProcessMaterial.Bind_get_collision_bounce, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsParticleProcessMaterial() Expert { return self[0].AsParticleProcessMaterial() }


//go:nosplit
func (self Simple) AsParticleProcessMaterial() Simple { return self[0].AsParticleProcessMaterial() }


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
func init() {classdb.Register("ParticleProcessMaterial", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
