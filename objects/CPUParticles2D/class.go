package CPUParticles2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Node2D"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
CPU-based 2D particle node used to create a variety of particle systems and effects.
See also [GPUParticles2D], which provides the same functionality with hardware acceleration, but may not run on older devices.
*/
type Instance [1]classdb.CPUParticles2D

/*
Restarts the particle emitter.
*/
func (self Instance) Restart() {
	class(self).Restart()
}

/*
Sets this node's properties to match a given [GPUParticles2D] node with an assigned [ParticleProcessMaterial].
*/
func (self Instance) ConvertFromParticles(particles objects.Node) {
	class(self).ConvertFromParticles(particles)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.CPUParticles2D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CPUParticles2D"))
	return Instance{classdb.CPUParticles2D(object)}
}

func (self Instance) Emitting() bool {
	return bool(class(self).IsEmitting())
}

func (self Instance) SetEmitting(value bool) {
	class(self).SetEmitting(value)
}

func (self Instance) Amount() int {
	return int(int(class(self).GetAmount()))
}

func (self Instance) SetAmount(value int) {
	class(self).SetAmount(gd.Int(value))
}

func (self Instance) Lifetime() float64 {
	return float64(float64(class(self).GetLifetime()))
}

func (self Instance) SetLifetime(value float64) {
	class(self).SetLifetime(gd.Float(value))
}

func (self Instance) OneShot() bool {
	return bool(class(self).GetOneShot())
}

func (self Instance) SetOneShot(value bool) {
	class(self).SetOneShot(value)
}

func (self Instance) Preprocess() float64 {
	return float64(float64(class(self).GetPreProcessTime()))
}

func (self Instance) SetPreprocess(value float64) {
	class(self).SetPreProcessTime(gd.Float(value))
}

func (self Instance) SpeedScale() float64 {
	return float64(float64(class(self).GetSpeedScale()))
}

func (self Instance) SetSpeedScale(value float64) {
	class(self).SetSpeedScale(gd.Float(value))
}

func (self Instance) Explosiveness() float64 {
	return float64(float64(class(self).GetExplosivenessRatio()))
}

func (self Instance) SetExplosiveness(value float64) {
	class(self).SetExplosivenessRatio(gd.Float(value))
}

func (self Instance) Randomness() float64 {
	return float64(float64(class(self).GetRandomnessRatio()))
}

func (self Instance) SetRandomness(value float64) {
	class(self).SetRandomnessRatio(gd.Float(value))
}

func (self Instance) LifetimeRandomness() float64 {
	return float64(float64(class(self).GetLifetimeRandomness()))
}

func (self Instance) SetLifetimeRandomness(value float64) {
	class(self).SetLifetimeRandomness(gd.Float(value))
}

func (self Instance) FixedFps() int {
	return int(int(class(self).GetFixedFps()))
}

func (self Instance) SetFixedFps(value int) {
	class(self).SetFixedFps(gd.Int(value))
}

func (self Instance) FractDelta() bool {
	return bool(class(self).GetFractionalDelta())
}

func (self Instance) SetFractDelta(value bool) {
	class(self).SetFractionalDelta(value)
}

func (self Instance) LocalCoords() bool {
	return bool(class(self).GetUseLocalCoordinates())
}

func (self Instance) SetLocalCoords(value bool) {
	class(self).SetUseLocalCoordinates(value)
}

func (self Instance) DrawOrder() classdb.CPUParticles2DDrawOrder {
	return classdb.CPUParticles2DDrawOrder(class(self).GetDrawOrder())
}

func (self Instance) SetDrawOrder(value classdb.CPUParticles2DDrawOrder) {
	class(self).SetDrawOrder(value)
}

func (self Instance) Texture() objects.Texture2D {
	return objects.Texture2D(class(self).GetTexture())
}

func (self Instance) SetTexture(value objects.Texture2D) {
	class(self).SetTexture(value)
}

func (self Instance) EmissionShape() classdb.CPUParticles2DEmissionShape {
	return classdb.CPUParticles2DEmissionShape(class(self).GetEmissionShape())
}

func (self Instance) SetEmissionShape(value classdb.CPUParticles2DEmissionShape) {
	class(self).SetEmissionShape(value)
}

func (self Instance) EmissionSphereRadius() float64 {
	return float64(float64(class(self).GetEmissionSphereRadius()))
}

func (self Instance) SetEmissionSphereRadius(value float64) {
	class(self).SetEmissionSphereRadius(gd.Float(value))
}

func (self Instance) EmissionRectExtents() gd.Vector2 {
	return gd.Vector2(class(self).GetEmissionRectExtents())
}

func (self Instance) SetEmissionRectExtents(value gd.Vector2) {
	class(self).SetEmissionRectExtents(value)
}

func (self Instance) EmissionPoints() []gd.Vector2 {
	return []gd.Vector2(class(self).GetEmissionPoints().AsSlice())
}

func (self Instance) SetEmissionPoints(value []gd.Vector2) {
	class(self).SetEmissionPoints(gd.NewPackedVector2Slice(value))
}

func (self Instance) EmissionNormals() []gd.Vector2 {
	return []gd.Vector2(class(self).GetEmissionNormals().AsSlice())
}

func (self Instance) SetEmissionNormals(value []gd.Vector2) {
	class(self).SetEmissionNormals(gd.NewPackedVector2Slice(value))
}

func (self Instance) EmissionColors() []gd.Color {
	return []gd.Color(class(self).GetEmissionColors().AsSlice())
}

func (self Instance) SetEmissionColors(value []gd.Color) {
	class(self).SetEmissionColors(gd.NewPackedColorSlice(value))
}

func (self Instance) ParticleFlagAlignY() bool {
	return bool(class(self).GetParticleFlag(0))
}

func (self Instance) SetParticleFlagAlignY(value bool) {
	class(self).SetParticleFlag(0, value)
}

func (self Instance) Direction() gd.Vector2 {
	return gd.Vector2(class(self).GetDirection())
}

func (self Instance) SetDirection(value gd.Vector2) {
	class(self).SetDirection(value)
}

func (self Instance) Spread() float64 {
	return float64(float64(class(self).GetSpread()))
}

func (self Instance) SetSpread(value float64) {
	class(self).SetSpread(gd.Float(value))
}

func (self Instance) Gravity() gd.Vector2 {
	return gd.Vector2(class(self).GetGravity())
}

func (self Instance) SetGravity(value gd.Vector2) {
	class(self).SetGravity(value)
}

func (self Instance) InitialVelocityMin() float64 {
	return float64(float64(class(self).GetParamMin(0)))
}

func (self Instance) SetInitialVelocityMin(value float64) {
	class(self).SetParamMin(0, gd.Float(value))
}

func (self Instance) InitialVelocityMax() float64 {
	return float64(float64(class(self).GetParamMax(0)))
}

func (self Instance) SetInitialVelocityMax(value float64) {
	class(self).SetParamMax(0, gd.Float(value))
}

func (self Instance) AngularVelocityMin() float64 {
	return float64(float64(class(self).GetParamMin(1)))
}

func (self Instance) SetAngularVelocityMin(value float64) {
	class(self).SetParamMin(1, gd.Float(value))
}

func (self Instance) AngularVelocityMax() float64 {
	return float64(float64(class(self).GetParamMax(1)))
}

func (self Instance) SetAngularVelocityMax(value float64) {
	class(self).SetParamMax(1, gd.Float(value))
}

func (self Instance) AngularVelocityCurve() objects.Curve {
	return objects.Curve(class(self).GetParamCurve(1))
}

func (self Instance) SetAngularVelocityCurve(value objects.Curve) {
	class(self).SetParamCurve(1, value)
}

func (self Instance) OrbitVelocityMin() float64 {
	return float64(float64(class(self).GetParamMin(2)))
}

func (self Instance) SetOrbitVelocityMin(value float64) {
	class(self).SetParamMin(2, gd.Float(value))
}

func (self Instance) OrbitVelocityMax() float64 {
	return float64(float64(class(self).GetParamMax(2)))
}

func (self Instance) SetOrbitVelocityMax(value float64) {
	class(self).SetParamMax(2, gd.Float(value))
}

func (self Instance) OrbitVelocityCurve() objects.Curve {
	return objects.Curve(class(self).GetParamCurve(2))
}

func (self Instance) SetOrbitVelocityCurve(value objects.Curve) {
	class(self).SetParamCurve(2, value)
}

func (self Instance) LinearAccelMin() float64 {
	return float64(float64(class(self).GetParamMin(3)))
}

func (self Instance) SetLinearAccelMin(value float64) {
	class(self).SetParamMin(3, gd.Float(value))
}

func (self Instance) LinearAccelMax() float64 {
	return float64(float64(class(self).GetParamMax(3)))
}

func (self Instance) SetLinearAccelMax(value float64) {
	class(self).SetParamMax(3, gd.Float(value))
}

func (self Instance) LinearAccelCurve() objects.Curve {
	return objects.Curve(class(self).GetParamCurve(3))
}

func (self Instance) SetLinearAccelCurve(value objects.Curve) {
	class(self).SetParamCurve(3, value)
}

func (self Instance) RadialAccelMin() float64 {
	return float64(float64(class(self).GetParamMin(4)))
}

func (self Instance) SetRadialAccelMin(value float64) {
	class(self).SetParamMin(4, gd.Float(value))
}

func (self Instance) RadialAccelMax() float64 {
	return float64(float64(class(self).GetParamMax(4)))
}

func (self Instance) SetRadialAccelMax(value float64) {
	class(self).SetParamMax(4, gd.Float(value))
}

func (self Instance) RadialAccelCurve() objects.Curve {
	return objects.Curve(class(self).GetParamCurve(4))
}

func (self Instance) SetRadialAccelCurve(value objects.Curve) {
	class(self).SetParamCurve(4, value)
}

func (self Instance) TangentialAccelMin() float64 {
	return float64(float64(class(self).GetParamMin(5)))
}

func (self Instance) SetTangentialAccelMin(value float64) {
	class(self).SetParamMin(5, gd.Float(value))
}

func (self Instance) TangentialAccelMax() float64 {
	return float64(float64(class(self).GetParamMax(5)))
}

func (self Instance) SetTangentialAccelMax(value float64) {
	class(self).SetParamMax(5, gd.Float(value))
}

func (self Instance) TangentialAccelCurve() objects.Curve {
	return objects.Curve(class(self).GetParamCurve(5))
}

func (self Instance) SetTangentialAccelCurve(value objects.Curve) {
	class(self).SetParamCurve(5, value)
}

func (self Instance) DampingMin() float64 {
	return float64(float64(class(self).GetParamMin(6)))
}

func (self Instance) SetDampingMin(value float64) {
	class(self).SetParamMin(6, gd.Float(value))
}

func (self Instance) DampingMax() float64 {
	return float64(float64(class(self).GetParamMax(6)))
}

func (self Instance) SetDampingMax(value float64) {
	class(self).SetParamMax(6, gd.Float(value))
}

func (self Instance) DampingCurve() objects.Curve {
	return objects.Curve(class(self).GetParamCurve(6))
}

func (self Instance) SetDampingCurve(value objects.Curve) {
	class(self).SetParamCurve(6, value)
}

func (self Instance) AngleMin() float64 {
	return float64(float64(class(self).GetParamMin(7)))
}

func (self Instance) SetAngleMin(value float64) {
	class(self).SetParamMin(7, gd.Float(value))
}

func (self Instance) AngleMax() float64 {
	return float64(float64(class(self).GetParamMax(7)))
}

func (self Instance) SetAngleMax(value float64) {
	class(self).SetParamMax(7, gd.Float(value))
}

func (self Instance) AngleCurve() objects.Curve {
	return objects.Curve(class(self).GetParamCurve(7))
}

func (self Instance) SetAngleCurve(value objects.Curve) {
	class(self).SetParamCurve(7, value)
}

func (self Instance) ScaleAmountMin() float64 {
	return float64(float64(class(self).GetParamMin(8)))
}

func (self Instance) SetScaleAmountMin(value float64) {
	class(self).SetParamMin(8, gd.Float(value))
}

func (self Instance) ScaleAmountMax() float64 {
	return float64(float64(class(self).GetParamMax(8)))
}

func (self Instance) SetScaleAmountMax(value float64) {
	class(self).SetParamMax(8, gd.Float(value))
}

func (self Instance) ScaleAmountCurve() objects.Curve {
	return objects.Curve(class(self).GetParamCurve(8))
}

func (self Instance) SetScaleAmountCurve(value objects.Curve) {
	class(self).SetParamCurve(8, value)
}

func (self Instance) SplitScale() bool {
	return bool(class(self).GetSplitScale())
}

func (self Instance) SetSplitScale(value bool) {
	class(self).SetSplitScale(value)
}

func (self Instance) ScaleCurveX() objects.Curve {
	return objects.Curve(class(self).GetScaleCurveX())
}

func (self Instance) SetScaleCurveX(value objects.Curve) {
	class(self).SetScaleCurveX(value)
}

func (self Instance) ScaleCurveY() objects.Curve {
	return objects.Curve(class(self).GetScaleCurveY())
}

func (self Instance) SetScaleCurveY(value objects.Curve) {
	class(self).SetScaleCurveY(value)
}

func (self Instance) Color() gd.Color {
	return gd.Color(class(self).GetColor())
}

func (self Instance) SetColor(value gd.Color) {
	class(self).SetColor(value)
}

func (self Instance) ColorRamp() objects.Gradient {
	return objects.Gradient(class(self).GetColorRamp())
}

func (self Instance) SetColorRamp(value objects.Gradient) {
	class(self).SetColorRamp(value)
}

func (self Instance) ColorInitialRamp() objects.Gradient {
	return objects.Gradient(class(self).GetColorInitialRamp())
}

func (self Instance) SetColorInitialRamp(value objects.Gradient) {
	class(self).SetColorInitialRamp(value)
}

func (self Instance) HueVariationMin() float64 {
	return float64(float64(class(self).GetParamMin(9)))
}

func (self Instance) SetHueVariationMin(value float64) {
	class(self).SetParamMin(9, gd.Float(value))
}

func (self Instance) HueVariationMax() float64 {
	return float64(float64(class(self).GetParamMax(9)))
}

func (self Instance) SetHueVariationMax(value float64) {
	class(self).SetParamMax(9, gd.Float(value))
}

func (self Instance) HueVariationCurve() objects.Curve {
	return objects.Curve(class(self).GetParamCurve(9))
}

func (self Instance) SetHueVariationCurve(value objects.Curve) {
	class(self).SetParamCurve(9, value)
}

func (self Instance) AnimSpeedMin() float64 {
	return float64(float64(class(self).GetParamMin(10)))
}

func (self Instance) SetAnimSpeedMin(value float64) {
	class(self).SetParamMin(10, gd.Float(value))
}

func (self Instance) AnimSpeedMax() float64 {
	return float64(float64(class(self).GetParamMax(10)))
}

func (self Instance) SetAnimSpeedMax(value float64) {
	class(self).SetParamMax(10, gd.Float(value))
}

func (self Instance) AnimSpeedCurve() objects.Curve {
	return objects.Curve(class(self).GetParamCurve(10))
}

func (self Instance) SetAnimSpeedCurve(value objects.Curve) {
	class(self).SetParamCurve(10, value)
}

func (self Instance) AnimOffsetMin() float64 {
	return float64(float64(class(self).GetParamMin(11)))
}

func (self Instance) SetAnimOffsetMin(value float64) {
	class(self).SetParamMin(11, gd.Float(value))
}

func (self Instance) AnimOffsetMax() float64 {
	return float64(float64(class(self).GetParamMax(11)))
}

func (self Instance) SetAnimOffsetMax(value float64) {
	class(self).SetParamMax(11, gd.Float(value))
}

func (self Instance) AnimOffsetCurve() objects.Curve {
	return objects.Curve(class(self).GetParamCurve(11))
}

func (self Instance) SetAnimOffsetCurve(value objects.Curve) {
	class(self).SetParamCurve(11, value)
}

//go:nosplit
func (self class) SetEmitting(emitting bool) {
	var frame = callframe.New()
	callframe.Arg(frame, emitting)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_emitting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetAmount(amount gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetLifetime(secs gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_lifetime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetOneShot(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_one_shot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetPreProcessTime(secs gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_pre_process_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetExplosivenessRatio(ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_explosiveness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetRandomnessRatio(ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_randomness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetLifetimeRandomness(random gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, random)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_lifetime_randomness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetUseLocalCoordinates(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_use_local_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetFixedFps(fps gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, fps)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_fixed_fps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetFractionalDelta(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_fractional_delta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetSpeedScale(scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsEmitting() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_is_emitting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetAmount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetLifetime() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_lifetime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetOneShot() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_one_shot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetPreProcessTime() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_pre_process_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetExplosivenessRatio() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_explosiveness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetRandomnessRatio() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_randomness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetLifetimeRandomness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_lifetime_randomness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetUseLocalCoordinates() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_use_local_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetFixedFps() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_fixed_fps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetFractionalDelta() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_fractional_delta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetSpeedScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawOrder(order classdb.CPUParticles2DDrawOrder) {
	var frame = callframe.New()
	callframe.Arg(frame, order)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_draw_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDrawOrder() classdb.CPUParticles2DDrawOrder {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CPUParticles2DDrawOrder](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_draw_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTexture(texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Restarts the particle emitter.
*/
//go:nosplit
func (self class) Restart() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_restart, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetDirection(direction gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDirection() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpread(spread gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, spread)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_spread, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpread() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_spread, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the minimum value for the given parameter.
*/
//go:nosplit
func (self class) SetParamMin(param classdb.CPUParticles2DParameter, value gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_param_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the minimum value range for the given parameter.
*/
//go:nosplit
func (self class) GetParamMin(param classdb.CPUParticles2DParameter) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_param_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the maximum value for the given parameter.
*/
//go:nosplit
func (self class) SetParamMax(param classdb.CPUParticles2DParameter, value gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_param_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the maximum value range for the given parameter.
*/
//go:nosplit
func (self class) GetParamMax(param classdb.CPUParticles2DParameter) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_param_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [Curve] of the parameter specified by [enum Parameter].
*/
//go:nosplit
func (self class) SetParamCurve(param classdb.CPUParticles2DParameter, curve objects.Curve) {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_param_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [Curve] of the parameter specified by [enum Parameter].
*/
//go:nosplit
func (self class) GetParamCurve(param classdb.CPUParticles2DParameter) objects.Curve {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_param_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Curve{classdb.Curve(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColorRamp(ramp objects.Gradient) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(ramp[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_color_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetColorRamp() objects.Gradient {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_color_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Gradient{classdb.Gradient(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColorInitialRamp(ramp objects.Gradient) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(ramp[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_color_initial_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetColorInitialRamp() objects.Gradient {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_color_initial_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Gradient{classdb.Gradient(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Enables or disables the given flag (see [enum ParticleFlags] for options).
*/
//go:nosplit
func (self class) SetParticleFlag(particle_flag classdb.CPUParticles2DParticleFlags, enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, particle_flag)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_particle_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the enabled state of the given flag (see [enum ParticleFlags] for options).
*/
//go:nosplit
func (self class) GetParticleFlag(particle_flag classdb.CPUParticles2DParticleFlags) bool {
	var frame = callframe.New()
	callframe.Arg(frame, particle_flag)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_particle_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionShape(shape classdb.CPUParticles2DEmissionShape) {
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_emission_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionShape() classdb.CPUParticles2DEmissionShape {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CPUParticles2DEmissionShape](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_emission_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionSphereRadius(radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_emission_sphere_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionSphereRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_emission_sphere_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionRectExtents(extents gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, extents)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_emission_rect_extents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionRectExtents() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_emission_rect_extents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionPoints(array gd.PackedVector2Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(array))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_emission_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionPoints() gd.PackedVector2Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_emission_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionNormals(array gd.PackedVector2Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(array))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_emission_normals, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionNormals() gd.PackedVector2Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_emission_normals, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionColors(array gd.PackedColorArray) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(array))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_emission_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionColors() gd.PackedColorArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_emission_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedColorArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetGravity() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravity(accel_vec gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, accel_vec)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSplitScale() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_split_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSplitScale(split_scale bool) {
	var frame = callframe.New()
	callframe.Arg(frame, split_scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_split_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScaleCurveX() objects.Curve {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_scale_curve_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Curve{classdb.Curve(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScaleCurveX(scale_curve objects.Curve) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scale_curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_scale_curve_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScaleCurveY() objects.Curve {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_get_scale_curve_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Curve{classdb.Curve(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScaleCurveY(scale_curve objects.Curve) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scale_curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_set_scale_curve_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets this node's properties to match a given [GPUParticles2D] node with an assigned [ParticleProcessMaterial].
*/
//go:nosplit
func (self class) ConvertFromParticles(particles objects.Node) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(particles[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles2D.Bind_convert_from_particles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Instance) OnFinished(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("finished"), gd.NewCallable(cb), 0)
}

func (self class) AsCPUParticles2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCPUParticles2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced     { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance  { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {
	classdb.Register("CPUParticles2D", func(ptr gd.Object) any { return classdb.CPUParticles2D(ptr) })
}

type DrawOrder = classdb.CPUParticles2DDrawOrder

const (
	/*Particles are drawn in the order emitted.*/
	DrawOrderIndex DrawOrder = 0
	/*Particles are drawn in order of remaining lifetime. In other words, the particle with the highest lifetime is drawn at the front.*/
	DrawOrderLifetime DrawOrder = 1
)

type Parameter = classdb.CPUParticles2DParameter

const (
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_curve] to set initial velocity properties.*/
	ParamInitialLinearVelocity Parameter = 0
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_curve] to set angular velocity properties.*/
	ParamAngularVelocity Parameter = 1
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_curve] to set orbital velocity properties.*/
	ParamOrbitVelocity Parameter = 2
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_curve] to set linear acceleration properties.*/
	ParamLinearAccel Parameter = 3
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_curve] to set radial acceleration properties.*/
	ParamRadialAccel Parameter = 4
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_curve] to set tangential acceleration properties.*/
	ParamTangentialAccel Parameter = 5
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_curve] to set damping properties.*/
	ParamDamping Parameter = 6
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_curve] to set angle properties.*/
	ParamAngle Parameter = 7
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_curve] to set scale properties.*/
	ParamScale Parameter = 8
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_curve] to set hue variation properties.*/
	ParamHueVariation Parameter = 9
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_curve] to set animation speed properties.*/
	ParamAnimSpeed Parameter = 10
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_curve] to set animation offset properties.*/
	ParamAnimOffset Parameter = 11
	/*Represents the size of the [enum Parameter] enum.*/
	ParamMax Parameter = 12
)

type ParticleFlags = classdb.CPUParticles2DParticleFlags

const (
	/*Use with [method set_particle_flag] to set [member particle_flag_align_y].*/
	ParticleFlagAlignYToVelocity ParticleFlags = 0
	/*Present for consistency with 3D particle nodes, not used in 2D.*/
	ParticleFlagRotateY ParticleFlags = 1
	/*Present for consistency with 3D particle nodes, not used in 2D.*/
	ParticleFlagDisableZ ParticleFlags = 2
	/*Represents the size of the [enum ParticleFlags] enum.*/
	ParticleFlagMax ParticleFlags = 3
)

type EmissionShape = classdb.CPUParticles2DEmissionShape

const (
	/*All particles will be emitted from a single point.*/
	EmissionShapePoint EmissionShape = 0
	/*Particles will be emitted in the volume of a sphere flattened to two dimensions.*/
	EmissionShapeSphere EmissionShape = 1
	/*Particles will be emitted on the surface of a sphere flattened to two dimensions.*/
	EmissionShapeSphereSurface EmissionShape = 2
	/*Particles will be emitted in the area of a rectangle.*/
	EmissionShapeRectangle EmissionShape = 3
	/*Particles will be emitted at a position chosen randomly among [member emission_points]. Particle color will be modulated by [member emission_colors].*/
	EmissionShapePoints EmissionShape = 4
	/*Particles will be emitted at a position chosen randomly among [member emission_points]. Particle velocity and rotation will be set based on [member emission_normals]. Particle color will be modulated by [member emission_colors].*/
	EmissionShapeDirectedPoints EmissionShape = 5
	/*Represents the size of the [enum EmissionShape] enum.*/
	EmissionShapeMax EmissionShape = 6
)
