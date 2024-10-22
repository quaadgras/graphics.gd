package CPUParticles2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node2D"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
CPU-based 2D particle node used to create a variety of particle systems and effects.
See also [GPUParticles2D], which provides the same functionality with hardware acceleration, but may not run on older devices.

*/
type Go [1]classdb.CPUParticles2D

/*
Restarts the particle emitter.
*/
func (self Go) Restart() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Restart()
}

/*
Sets this node's properties to match a given [GPUParticles2D] node with an assigned [ParticleProcessMaterial].
*/
func (self Go) ConvertFromParticles(particles gdclass.Node) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ConvertFromParticles(particles)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CPUParticles2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("CPUParticles2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Emitting() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsEmitting())
}

func (self Go) SetEmitting(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEmitting(value)
}

func (self Go) Amount() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetAmount()))
}

func (self Go) SetAmount(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAmount(gd.Int(value))
}

func (self Go) Lifetime() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetLifetime()))
}

func (self Go) SetLifetime(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLifetime(gd.Float(value))
}

func (self Go) OneShot() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetOneShot())
}

func (self Go) SetOneShot(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOneShot(value)
}

func (self Go) Preprocess() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetPreProcessTime()))
}

func (self Go) SetPreprocess(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPreProcessTime(gd.Float(value))
}

func (self Go) SpeedScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSpeedScale()))
}

func (self Go) SetSpeedScale(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSpeedScale(gd.Float(value))
}

func (self Go) Explosiveness() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetExplosivenessRatio()))
}

func (self Go) SetExplosiveness(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetExplosivenessRatio(gd.Float(value))
}

func (self Go) Randomness() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetRandomnessRatio()))
}

func (self Go) SetRandomness(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRandomnessRatio(gd.Float(value))
}

func (self Go) LifetimeRandomness() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetLifetimeRandomness()))
}

func (self Go) SetLifetimeRandomness(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLifetimeRandomness(gd.Float(value))
}

func (self Go) FixedFps() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetFixedFps()))
}

func (self Go) SetFixedFps(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFixedFps(gd.Int(value))
}

func (self Go) FractDelta() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetFractionalDelta())
}

func (self Go) SetFractDelta(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFractionalDelta(value)
}

func (self Go) LocalCoords() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetUseLocalCoordinates())
}

func (self Go) SetLocalCoords(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseLocalCoordinates(value)
}

func (self Go) DrawOrder() classdb.CPUParticles2DDrawOrder {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.CPUParticles2DDrawOrder(class(self).GetDrawOrder())
}

func (self Go) SetDrawOrder(value classdb.CPUParticles2DDrawOrder) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawOrder(value)
}

func (self Go) Texture() gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Texture2D(class(self).GetTexture(gc))
}

func (self Go) SetTexture(value gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTexture(value)
}

func (self Go) EmissionShape() classdb.CPUParticles2DEmissionShape {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.CPUParticles2DEmissionShape(class(self).GetEmissionShape())
}

func (self Go) SetEmissionShape(value classdb.CPUParticles2DEmissionShape) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEmissionShape(value)
}

func (self Go) EmissionSphereRadius() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetEmissionSphereRadius()))
}

func (self Go) SetEmissionSphereRadius(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEmissionSphereRadius(gd.Float(value))
}

func (self Go) EmissionRectExtents() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetEmissionRectExtents())
}

func (self Go) SetEmissionRectExtents(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEmissionRectExtents(value)
}

func (self Go) EmissionPoints() []gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return []gd.Vector2(class(self).GetEmissionPoints(gc).AsSlice())
}

func (self Go) SetEmissionPoints(value []gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEmissionPoints(gc.PackedVector2Slice(value))
}

func (self Go) EmissionNormals() []gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return []gd.Vector2(class(self).GetEmissionNormals(gc).AsSlice())
}

func (self Go) SetEmissionNormals(value []gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEmissionNormals(gc.PackedVector2Slice(value))
}

func (self Go) EmissionColors() []gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return []gd.Color(class(self).GetEmissionColors(gc).AsSlice())
}

func (self Go) SetEmissionColors(value []gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEmissionColors(gc.PackedColorSlice(value))
}

func (self Go) ParticleFlagAlignY() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetParticleFlag(0))
}

func (self Go) SetParticleFlagAlignY(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParticleFlag(0, value)
}

func (self Go) Direction() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetDirection())
}

func (self Go) SetDirection(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDirection(value)
}

func (self Go) Spread() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSpread()))
}

func (self Go) SetSpread(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSpread(gd.Float(value))
}

func (self Go) Gravity() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetGravity())
}

func (self Go) SetGravity(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGravity(value)
}

func (self Go) InitialVelocityMin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMin(0)))
}

func (self Go) SetInitialVelocityMin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMin(0, gd.Float(value))
}

func (self Go) InitialVelocityMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMax(0)))
}

func (self Go) SetInitialVelocityMax(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMax(0, gd.Float(value))
}

func (self Go) AngularVelocityMin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMin(1)))
}

func (self Go) SetAngularVelocityMin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMin(1, gd.Float(value))
}

func (self Go) AngularVelocityMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMax(1)))
}

func (self Go) SetAngularVelocityMax(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMax(1, gd.Float(value))
}

func (self Go) AngularVelocityCurve() gdclass.Curve {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Curve(class(self).GetParamCurve(gc,1))
}

func (self Go) SetAngularVelocityCurve(value gdclass.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamCurve(1, value)
}

func (self Go) OrbitVelocityMin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMin(2)))
}

func (self Go) SetOrbitVelocityMin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMin(2, gd.Float(value))
}

func (self Go) OrbitVelocityMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMax(2)))
}

func (self Go) SetOrbitVelocityMax(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMax(2, gd.Float(value))
}

func (self Go) OrbitVelocityCurve() gdclass.Curve {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Curve(class(self).GetParamCurve(gc,2))
}

func (self Go) SetOrbitVelocityCurve(value gdclass.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamCurve(2, value)
}

func (self Go) LinearAccelMin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMin(3)))
}

func (self Go) SetLinearAccelMin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMin(3, gd.Float(value))
}

func (self Go) LinearAccelMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMax(3)))
}

func (self Go) SetLinearAccelMax(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMax(3, gd.Float(value))
}

func (self Go) LinearAccelCurve() gdclass.Curve {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Curve(class(self).GetParamCurve(gc,3))
}

func (self Go) SetLinearAccelCurve(value gdclass.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamCurve(3, value)
}

func (self Go) RadialAccelMin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMin(4)))
}

func (self Go) SetRadialAccelMin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMin(4, gd.Float(value))
}

func (self Go) RadialAccelMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMax(4)))
}

func (self Go) SetRadialAccelMax(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMax(4, gd.Float(value))
}

func (self Go) RadialAccelCurve() gdclass.Curve {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Curve(class(self).GetParamCurve(gc,4))
}

func (self Go) SetRadialAccelCurve(value gdclass.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamCurve(4, value)
}

func (self Go) TangentialAccelMin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMin(5)))
}

func (self Go) SetTangentialAccelMin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMin(5, gd.Float(value))
}

func (self Go) TangentialAccelMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMax(5)))
}

func (self Go) SetTangentialAccelMax(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMax(5, gd.Float(value))
}

func (self Go) TangentialAccelCurve() gdclass.Curve {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Curve(class(self).GetParamCurve(gc,5))
}

func (self Go) SetTangentialAccelCurve(value gdclass.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamCurve(5, value)
}

func (self Go) DampingMin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMin(6)))
}

func (self Go) SetDampingMin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMin(6, gd.Float(value))
}

func (self Go) DampingMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMax(6)))
}

func (self Go) SetDampingMax(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMax(6, gd.Float(value))
}

func (self Go) DampingCurve() gdclass.Curve {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Curve(class(self).GetParamCurve(gc,6))
}

func (self Go) SetDampingCurve(value gdclass.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamCurve(6, value)
}

func (self Go) AngleMin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMin(7)))
}

func (self Go) SetAngleMin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMin(7, gd.Float(value))
}

func (self Go) AngleMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMax(7)))
}

func (self Go) SetAngleMax(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMax(7, gd.Float(value))
}

func (self Go) AngleCurve() gdclass.Curve {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Curve(class(self).GetParamCurve(gc,7))
}

func (self Go) SetAngleCurve(value gdclass.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamCurve(7, value)
}

func (self Go) ScaleAmountMin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMin(8)))
}

func (self Go) SetScaleAmountMin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMin(8, gd.Float(value))
}

func (self Go) ScaleAmountMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMax(8)))
}

func (self Go) SetScaleAmountMax(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMax(8, gd.Float(value))
}

func (self Go) ScaleAmountCurve() gdclass.Curve {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Curve(class(self).GetParamCurve(gc,8))
}

func (self Go) SetScaleAmountCurve(value gdclass.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamCurve(8, value)
}

func (self Go) SplitScale() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetSplitScale())
}

func (self Go) SetSplitScale(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSplitScale(value)
}

func (self Go) ScaleCurveX() gdclass.Curve {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Curve(class(self).GetScaleCurveX(gc))
}

func (self Go) SetScaleCurveX(value gdclass.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetScaleCurveX(value)
}

func (self Go) ScaleCurveY() gdclass.Curve {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Curve(class(self).GetScaleCurveY(gc))
}

func (self Go) SetScaleCurveY(value gdclass.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetScaleCurveY(value)
}

func (self Go) Color() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetColor())
}

func (self Go) SetColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetColor(value)
}

func (self Go) ColorRamp() gdclass.Gradient {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Gradient(class(self).GetColorRamp(gc))
}

func (self Go) SetColorRamp(value gdclass.Gradient) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetColorRamp(value)
}

func (self Go) ColorInitialRamp() gdclass.Gradient {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Gradient(class(self).GetColorInitialRamp(gc))
}

func (self Go) SetColorInitialRamp(value gdclass.Gradient) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetColorInitialRamp(value)
}

func (self Go) HueVariationMin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMin(9)))
}

func (self Go) SetHueVariationMin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMin(9, gd.Float(value))
}

func (self Go) HueVariationMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMax(9)))
}

func (self Go) SetHueVariationMax(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMax(9, gd.Float(value))
}

func (self Go) HueVariationCurve() gdclass.Curve {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Curve(class(self).GetParamCurve(gc,9))
}

func (self Go) SetHueVariationCurve(value gdclass.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamCurve(9, value)
}

func (self Go) AnimSpeedMin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMin(10)))
}

func (self Go) SetAnimSpeedMin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMin(10, gd.Float(value))
}

func (self Go) AnimSpeedMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMax(10)))
}

func (self Go) SetAnimSpeedMax(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMax(10, gd.Float(value))
}

func (self Go) AnimSpeedCurve() gdclass.Curve {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Curve(class(self).GetParamCurve(gc,10))
}

func (self Go) SetAnimSpeedCurve(value gdclass.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamCurve(10, value)
}

func (self Go) AnimOffsetMin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMin(11)))
}

func (self Go) SetAnimOffsetMin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMin(11, gd.Float(value))
}

func (self Go) AnimOffsetMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParamMax(11)))
}

func (self Go) SetAnimOffsetMax(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamMax(11, gd.Float(value))
}

func (self Go) AnimOffsetCurve() gdclass.Curve {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Curve(class(self).GetParamCurve(gc,11))
}

func (self Go) SetAnimOffsetCurve(value gdclass.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParamCurve(11, value)
}

//go:nosplit
func (self class) SetEmitting(emitting bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, emitting)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_emitting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetAmount(amount gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetLifetime(secs gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_lifetime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetOneShot(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_one_shot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetPreProcessTime(secs gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_pre_process_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetExplosivenessRatio(ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_explosiveness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetRandomnessRatio(ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_randomness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetLifetimeRandomness(random gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, random)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_lifetime_randomness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetUseLocalCoordinates(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_use_local_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFixedFps(fps gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fps)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_fixed_fps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFractionalDelta(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_fractional_delta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetSpeedScale(scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEmitting() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_is_emitting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetAmount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetLifetime() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_lifetime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetOneShot() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_one_shot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetPreProcessTime() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_pre_process_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetExplosivenessRatio() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_explosiveness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetRandomnessRatio() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_randomness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetLifetimeRandomness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_lifetime_randomness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetUseLocalCoordinates() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_use_local_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetFixedFps() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_fixed_fps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetFractionalDelta() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_fractional_delta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetSpeedScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDrawOrder(order classdb.CPUParticles2DDrawOrder)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, order)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_draw_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDrawOrder() classdb.CPUParticles2DDrawOrder {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CPUParticles2DDrawOrder](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_draw_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTexture(texture gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTexture(ctx gd.Lifetime) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Restarts the particle emitter.
*/
//go:nosplit
func (self class) Restart()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_restart, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetDirection(direction gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDirection() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSpread(spread gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, spread)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_spread, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSpread() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_spread, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the minimum value for the given parameter.
*/
//go:nosplit
func (self class) SetParamMin(param classdb.CPUParticles2DParameter, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_param_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the minimum value range for the given parameter.
*/
//go:nosplit
func (self class) GetParamMin(param classdb.CPUParticles2DParameter) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_param_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the maximum value for the given parameter.
*/
//go:nosplit
func (self class) SetParamMax(param classdb.CPUParticles2DParameter, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_param_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the maximum value range for the given parameter.
*/
//go:nosplit
func (self class) GetParamMax(param classdb.CPUParticles2DParameter) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_param_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [Curve] of the parameter specified by [enum Parameter].
*/
//go:nosplit
func (self class) SetParamCurve(param classdb.CPUParticles2DParameter, curve gdclass.Curve)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, mmm.Get(curve[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_param_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Curve] of the parameter specified by [enum Parameter].
*/
//go:nosplit
func (self class) GetParamCurve(ctx gd.Lifetime, param classdb.CPUParticles2DParameter) gdclass.Curve {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_param_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Curve
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColorRamp(ramp gdclass.Gradient)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(ramp[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_color_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColorRamp(ctx gd.Lifetime) gdclass.Gradient {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_color_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Gradient
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColorInitialRamp(ramp gdclass.Gradient)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(ramp[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_color_initial_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColorInitialRamp(ctx gd.Lifetime) gdclass.Gradient {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_color_initial_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Gradient
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Enables or disables the given flag (see [enum ParticleFlags] for options).
*/
//go:nosplit
func (self class) SetParticleFlag(particle_flag classdb.CPUParticles2DParticleFlags, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particle_flag)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_particle_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the enabled state of the given flag (see [enum ParticleFlags] for options).
*/
//go:nosplit
func (self class) GetParticleFlag(particle_flag classdb.CPUParticles2DParticleFlags) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particle_flag)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_particle_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionShape(shape classdb.CPUParticles2DEmissionShape)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_emission_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionShape() classdb.CPUParticles2DEmissionShape {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CPUParticles2DEmissionShape](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_emission_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_emission_sphere_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionSphereRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_emission_sphere_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionRectExtents(extents gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, extents)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_emission_rect_extents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionRectExtents() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_emission_rect_extents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionPoints(array gd.PackedVector2Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(array))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_emission_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionPoints(ctx gd.Lifetime) gd.PackedVector2Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_emission_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector2Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionNormals(array gd.PackedVector2Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(array))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_emission_normals, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionNormals(ctx gd.Lifetime) gd.PackedVector2Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_emission_normals, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector2Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionColors(array gd.PackedColorArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(array))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_emission_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionColors(ctx gd.Lifetime) gd.PackedColorArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_emission_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedColorArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetGravity() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGravity(accel_vec gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, accel_vec)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSplitScale() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_split_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSplitScale(split_scale bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, split_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_split_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetScaleCurveX(ctx gd.Lifetime) gdclass.Curve {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_scale_curve_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Curve
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScaleCurveX(scale_curve gdclass.Curve)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(scale_curve[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_scale_curve_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetScaleCurveY(ctx gd.Lifetime) gdclass.Curve {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_get_scale_curve_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Curve
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScaleCurveY(scale_curve gdclass.Curve)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(scale_curve[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_set_scale_curve_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets this node's properties to match a given [GPUParticles2D] node with an assigned [ParticleProcessMaterial].
*/
//go:nosplit
func (self class) ConvertFromParticles(particles gdclass.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(particles[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CPUParticles2D.Bind_convert_from_particles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnFinished(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("finished"), gc.Callable(cb), 0)
}


func (self class) AsCPUParticles2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCPUParticles2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.GD { return *((*Node2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode2D() Node2D.Go { return *((*Node2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {classdb.Register("CPUParticles2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
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
