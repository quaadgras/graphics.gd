// Package CPUParticles3D provides methods for working with CPUParticles3D object instances.
package CPUParticles3D

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
import "graphics.gd/variant/Dictionary"
import "graphics.gd/classdb/GeometryInstance3D"
import "graphics.gd/classdb/VisualInstance3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/AABB"
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
var _ Dictionary.Any

/*
CPU-based 3D particle node used to create a variety of particle systems and effects.
See also [GPUParticles3D], which provides the same functionality with hardware acceleration, but may not run on older devices.
*/
type Instance [1]gdclass.CPUParticles3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCPUParticles3D() Instance
}

/*
Restarts the particle emitter.
*/
func (self Instance) Restart() { //gd:CPUParticles3D.restart
	class(self).Restart()
}

/*
Sets this node's properties to match a given [GPUParticles3D] node with an assigned [ParticleProcessMaterial].
*/
func (self Instance) ConvertFromParticles(particles [1]gdclass.Node) { //gd:CPUParticles3D.convert_from_particles
	class(self).ConvertFromParticles(particles)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CPUParticles3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CPUParticles3D"))
	casted := Instance{*(*gdclass.CPUParticles3D)(unsafe.Pointer(&object))}
	return casted
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

func (self Instance) Lifetime() Float.X {
	return Float.X(Float.X(class(self).GetLifetime()))
}

func (self Instance) SetLifetime(value Float.X) {
	class(self).SetLifetime(gd.Float(value))
}

func (self Instance) OneShot() bool {
	return bool(class(self).GetOneShot())
}

func (self Instance) SetOneShot(value bool) {
	class(self).SetOneShot(value)
}

func (self Instance) Preprocess() Float.X {
	return Float.X(Float.X(class(self).GetPreProcessTime()))
}

func (self Instance) SetPreprocess(value Float.X) {
	class(self).SetPreProcessTime(gd.Float(value))
}

func (self Instance) SpeedScale() Float.X {
	return Float.X(Float.X(class(self).GetSpeedScale()))
}

func (self Instance) SetSpeedScale(value Float.X) {
	class(self).SetSpeedScale(gd.Float(value))
}

func (self Instance) Explosiveness() Float.X {
	return Float.X(Float.X(class(self).GetExplosivenessRatio()))
}

func (self Instance) SetExplosiveness(value Float.X) {
	class(self).SetExplosivenessRatio(gd.Float(value))
}

func (self Instance) Randomness() Float.X {
	return Float.X(Float.X(class(self).GetRandomnessRatio()))
}

func (self Instance) SetRandomness(value Float.X) {
	class(self).SetRandomnessRatio(gd.Float(value))
}

func (self Instance) LifetimeRandomness() Float.X {
	return Float.X(Float.X(class(self).GetLifetimeRandomness()))
}

func (self Instance) SetLifetimeRandomness(value Float.X) {
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

func (self Instance) VisibilityAabb() AABB.PositionSize {
	return AABB.PositionSize(class(self).GetVisibilityAabb())
}

func (self Instance) SetVisibilityAabb(value AABB.PositionSize) {
	class(self).SetVisibilityAabb(gd.AABB(value))
}

func (self Instance) LocalCoords() bool {
	return bool(class(self).GetUseLocalCoordinates())
}

func (self Instance) SetLocalCoords(value bool) {
	class(self).SetUseLocalCoordinates(value)
}

func (self Instance) DrawOrder() gdclass.CPUParticles3DDrawOrder {
	return gdclass.CPUParticles3DDrawOrder(class(self).GetDrawOrder())
}

func (self Instance) SetDrawOrder(value gdclass.CPUParticles3DDrawOrder) {
	class(self).SetDrawOrder(value)
}

func (self Instance) Mesh() [1]gdclass.Mesh {
	return [1]gdclass.Mesh(class(self).GetMesh())
}

func (self Instance) SetMesh(value [1]gdclass.Mesh) {
	class(self).SetMesh(value)
}

func (self Instance) EmissionShape() gdclass.CPUParticles3DEmissionShape {
	return gdclass.CPUParticles3DEmissionShape(class(self).GetEmissionShape())
}

func (self Instance) SetEmissionShape(value gdclass.CPUParticles3DEmissionShape) {
	class(self).SetEmissionShape(value)
}

func (self Instance) EmissionSphereRadius() Float.X {
	return Float.X(Float.X(class(self).GetEmissionSphereRadius()))
}

func (self Instance) SetEmissionSphereRadius(value Float.X) {
	class(self).SetEmissionSphereRadius(gd.Float(value))
}

func (self Instance) EmissionBoxExtents() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetEmissionBoxExtents())
}

func (self Instance) SetEmissionBoxExtents(value Vector3.XYZ) {
	class(self).SetEmissionBoxExtents(gd.Vector3(value))
}

func (self Instance) EmissionPoints() []Vector3.XYZ {
	return []Vector3.XYZ(class(self).GetEmissionPoints().AsSlice())
}

func (self Instance) SetEmissionPoints(value []Vector3.XYZ) {
	class(self).SetEmissionPoints(gd.NewPackedVector3Slice(*(*[]gd.Vector3)(unsafe.Pointer(&value))))
}

func (self Instance) EmissionNormals() []Vector3.XYZ {
	return []Vector3.XYZ(class(self).GetEmissionNormals().AsSlice())
}

func (self Instance) SetEmissionNormals(value []Vector3.XYZ) {
	class(self).SetEmissionNormals(gd.NewPackedVector3Slice(*(*[]gd.Vector3)(unsafe.Pointer(&value))))
}

func (self Instance) EmissionColors() []Color.RGBA {
	return []Color.RGBA(class(self).GetEmissionColors().AsSlice())
}

func (self Instance) SetEmissionColors(value []Color.RGBA) {
	class(self).SetEmissionColors(gd.NewPackedColorSlice(*(*[]gd.Color)(unsafe.Pointer(&value))))
}

func (self Instance) EmissionRingAxis() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetEmissionRingAxis())
}

func (self Instance) SetEmissionRingAxis(value Vector3.XYZ) {
	class(self).SetEmissionRingAxis(gd.Vector3(value))
}

func (self Instance) EmissionRingHeight() Float.X {
	return Float.X(Float.X(class(self).GetEmissionRingHeight()))
}

func (self Instance) SetEmissionRingHeight(value Float.X) {
	class(self).SetEmissionRingHeight(gd.Float(value))
}

func (self Instance) EmissionRingRadius() Float.X {
	return Float.X(Float.X(class(self).GetEmissionRingRadius()))
}

func (self Instance) SetEmissionRingRadius(value Float.X) {
	class(self).SetEmissionRingRadius(gd.Float(value))
}

func (self Instance) EmissionRingInnerRadius() Float.X {
	return Float.X(Float.X(class(self).GetEmissionRingInnerRadius()))
}

func (self Instance) SetEmissionRingInnerRadius(value Float.X) {
	class(self).SetEmissionRingInnerRadius(gd.Float(value))
}

func (self Instance) ParticleFlagAlignY() bool {
	return bool(class(self).GetParticleFlag(0))
}

func (self Instance) SetParticleFlagAlignY(value bool) {
	class(self).SetParticleFlag(0, value)
}

func (self Instance) ParticleFlagRotateY() bool {
	return bool(class(self).GetParticleFlag(1))
}

func (self Instance) SetParticleFlagRotateY(value bool) {
	class(self).SetParticleFlag(1, value)
}

func (self Instance) ParticleFlagDisableZ() bool {
	return bool(class(self).GetParticleFlag(2))
}

func (self Instance) SetParticleFlagDisableZ(value bool) {
	class(self).SetParticleFlag(2, value)
}

func (self Instance) Direction() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetDirection())
}

func (self Instance) SetDirection(value Vector3.XYZ) {
	class(self).SetDirection(gd.Vector3(value))
}

func (self Instance) Spread() Float.X {
	return Float.X(Float.X(class(self).GetSpread()))
}

func (self Instance) SetSpread(value Float.X) {
	class(self).SetSpread(gd.Float(value))
}

func (self Instance) Flatness() Float.X {
	return Float.X(Float.X(class(self).GetFlatness()))
}

func (self Instance) SetFlatness(value Float.X) {
	class(self).SetFlatness(gd.Float(value))
}

func (self Instance) Gravity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetGravity())
}

func (self Instance) SetGravity(value Vector3.XYZ) {
	class(self).SetGravity(gd.Vector3(value))
}

func (self Instance) InitialVelocityMin() Float.X {
	return Float.X(Float.X(class(self).GetParamMin(0)))
}

func (self Instance) SetInitialVelocityMin(value Float.X) {
	class(self).SetParamMin(0, gd.Float(value))
}

func (self Instance) InitialVelocityMax() Float.X {
	return Float.X(Float.X(class(self).GetParamMax(0)))
}

func (self Instance) SetInitialVelocityMax(value Float.X) {
	class(self).SetParamMax(0, gd.Float(value))
}

func (self Instance) AngularVelocityMin() Float.X {
	return Float.X(Float.X(class(self).GetParamMin(1)))
}

func (self Instance) SetAngularVelocityMin(value Float.X) {
	class(self).SetParamMin(1, gd.Float(value))
}

func (self Instance) AngularVelocityMax() Float.X {
	return Float.X(Float.X(class(self).GetParamMax(1)))
}

func (self Instance) SetAngularVelocityMax(value Float.X) {
	class(self).SetParamMax(1, gd.Float(value))
}

func (self Instance) AngularVelocityCurve() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetParamCurve(1))
}

func (self Instance) SetAngularVelocityCurve(value [1]gdclass.Curve) {
	class(self).SetParamCurve(1, value)
}

func (self Instance) OrbitVelocityMin() Float.X {
	return Float.X(Float.X(class(self).GetParamMin(2)))
}

func (self Instance) SetOrbitVelocityMin(value Float.X) {
	class(self).SetParamMin(2, gd.Float(value))
}

func (self Instance) OrbitVelocityMax() Float.X {
	return Float.X(Float.X(class(self).GetParamMax(2)))
}

func (self Instance) SetOrbitVelocityMax(value Float.X) {
	class(self).SetParamMax(2, gd.Float(value))
}

func (self Instance) OrbitVelocityCurve() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetParamCurve(2))
}

func (self Instance) SetOrbitVelocityCurve(value [1]gdclass.Curve) {
	class(self).SetParamCurve(2, value)
}

func (self Instance) LinearAccelMin() Float.X {
	return Float.X(Float.X(class(self).GetParamMin(3)))
}

func (self Instance) SetLinearAccelMin(value Float.X) {
	class(self).SetParamMin(3, gd.Float(value))
}

func (self Instance) LinearAccelMax() Float.X {
	return Float.X(Float.X(class(self).GetParamMax(3)))
}

func (self Instance) SetLinearAccelMax(value Float.X) {
	class(self).SetParamMax(3, gd.Float(value))
}

func (self Instance) LinearAccelCurve() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetParamCurve(3))
}

func (self Instance) SetLinearAccelCurve(value [1]gdclass.Curve) {
	class(self).SetParamCurve(3, value)
}

func (self Instance) RadialAccelMin() Float.X {
	return Float.X(Float.X(class(self).GetParamMin(4)))
}

func (self Instance) SetRadialAccelMin(value Float.X) {
	class(self).SetParamMin(4, gd.Float(value))
}

func (self Instance) RadialAccelMax() Float.X {
	return Float.X(Float.X(class(self).GetParamMax(4)))
}

func (self Instance) SetRadialAccelMax(value Float.X) {
	class(self).SetParamMax(4, gd.Float(value))
}

func (self Instance) RadialAccelCurve() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetParamCurve(4))
}

func (self Instance) SetRadialAccelCurve(value [1]gdclass.Curve) {
	class(self).SetParamCurve(4, value)
}

func (self Instance) TangentialAccelMin() Float.X {
	return Float.X(Float.X(class(self).GetParamMin(5)))
}

func (self Instance) SetTangentialAccelMin(value Float.X) {
	class(self).SetParamMin(5, gd.Float(value))
}

func (self Instance) TangentialAccelMax() Float.X {
	return Float.X(Float.X(class(self).GetParamMax(5)))
}

func (self Instance) SetTangentialAccelMax(value Float.X) {
	class(self).SetParamMax(5, gd.Float(value))
}

func (self Instance) TangentialAccelCurve() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetParamCurve(5))
}

func (self Instance) SetTangentialAccelCurve(value [1]gdclass.Curve) {
	class(self).SetParamCurve(5, value)
}

func (self Instance) DampingMin() Float.X {
	return Float.X(Float.X(class(self).GetParamMin(6)))
}

func (self Instance) SetDampingMin(value Float.X) {
	class(self).SetParamMin(6, gd.Float(value))
}

func (self Instance) DampingMax() Float.X {
	return Float.X(Float.X(class(self).GetParamMax(6)))
}

func (self Instance) SetDampingMax(value Float.X) {
	class(self).SetParamMax(6, gd.Float(value))
}

func (self Instance) DampingCurve() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetParamCurve(6))
}

func (self Instance) SetDampingCurve(value [1]gdclass.Curve) {
	class(self).SetParamCurve(6, value)
}

func (self Instance) AngleMin() Float.X {
	return Float.X(Float.X(class(self).GetParamMin(7)))
}

func (self Instance) SetAngleMin(value Float.X) {
	class(self).SetParamMin(7, gd.Float(value))
}

func (self Instance) AngleMax() Float.X {
	return Float.X(Float.X(class(self).GetParamMax(7)))
}

func (self Instance) SetAngleMax(value Float.X) {
	class(self).SetParamMax(7, gd.Float(value))
}

func (self Instance) AngleCurve() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetParamCurve(7))
}

func (self Instance) SetAngleCurve(value [1]gdclass.Curve) {
	class(self).SetParamCurve(7, value)
}

func (self Instance) ScaleAmountMin() Float.X {
	return Float.X(Float.X(class(self).GetParamMin(8)))
}

func (self Instance) SetScaleAmountMin(value Float.X) {
	class(self).SetParamMin(8, gd.Float(value))
}

func (self Instance) ScaleAmountMax() Float.X {
	return Float.X(Float.X(class(self).GetParamMax(8)))
}

func (self Instance) SetScaleAmountMax(value Float.X) {
	class(self).SetParamMax(8, gd.Float(value))
}

func (self Instance) ScaleAmountCurve() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetParamCurve(8))
}

func (self Instance) SetScaleAmountCurve(value [1]gdclass.Curve) {
	class(self).SetParamCurve(8, value)
}

func (self Instance) SplitScale() bool {
	return bool(class(self).GetSplitScale())
}

func (self Instance) SetSplitScale(value bool) {
	class(self).SetSplitScale(value)
}

func (self Instance) ScaleCurveX() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetScaleCurveX())
}

func (self Instance) SetScaleCurveX(value [1]gdclass.Curve) {
	class(self).SetScaleCurveX(value)
}

func (self Instance) ScaleCurveY() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetScaleCurveY())
}

func (self Instance) SetScaleCurveY(value [1]gdclass.Curve) {
	class(self).SetScaleCurveY(value)
}

func (self Instance) ScaleCurveZ() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetScaleCurveZ())
}

func (self Instance) SetScaleCurveZ(value [1]gdclass.Curve) {
	class(self).SetScaleCurveZ(value)
}

func (self Instance) Color() Color.RGBA {
	return Color.RGBA(class(self).GetColor())
}

func (self Instance) SetColor(value Color.RGBA) {
	class(self).SetColor(gd.Color(value))
}

func (self Instance) ColorRamp() [1]gdclass.Gradient {
	return [1]gdclass.Gradient(class(self).GetColorRamp())
}

func (self Instance) SetColorRamp(value [1]gdclass.Gradient) {
	class(self).SetColorRamp(value)
}

func (self Instance) ColorInitialRamp() [1]gdclass.Gradient {
	return [1]gdclass.Gradient(class(self).GetColorInitialRamp())
}

func (self Instance) SetColorInitialRamp(value [1]gdclass.Gradient) {
	class(self).SetColorInitialRamp(value)
}

func (self Instance) HueVariationMin() Float.X {
	return Float.X(Float.X(class(self).GetParamMin(9)))
}

func (self Instance) SetHueVariationMin(value Float.X) {
	class(self).SetParamMin(9, gd.Float(value))
}

func (self Instance) HueVariationMax() Float.X {
	return Float.X(Float.X(class(self).GetParamMax(9)))
}

func (self Instance) SetHueVariationMax(value Float.X) {
	class(self).SetParamMax(9, gd.Float(value))
}

func (self Instance) HueVariationCurve() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetParamCurve(9))
}

func (self Instance) SetHueVariationCurve(value [1]gdclass.Curve) {
	class(self).SetParamCurve(9, value)
}

func (self Instance) AnimSpeedMin() Float.X {
	return Float.X(Float.X(class(self).GetParamMin(10)))
}

func (self Instance) SetAnimSpeedMin(value Float.X) {
	class(self).SetParamMin(10, gd.Float(value))
}

func (self Instance) AnimSpeedMax() Float.X {
	return Float.X(Float.X(class(self).GetParamMax(10)))
}

func (self Instance) SetAnimSpeedMax(value Float.X) {
	class(self).SetParamMax(10, gd.Float(value))
}

func (self Instance) AnimSpeedCurve() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetParamCurve(10))
}

func (self Instance) SetAnimSpeedCurve(value [1]gdclass.Curve) {
	class(self).SetParamCurve(10, value)
}

func (self Instance) AnimOffsetMin() Float.X {
	return Float.X(Float.X(class(self).GetParamMin(11)))
}

func (self Instance) SetAnimOffsetMin(value Float.X) {
	class(self).SetParamMin(11, gd.Float(value))
}

func (self Instance) AnimOffsetMax() Float.X {
	return Float.X(Float.X(class(self).GetParamMax(11)))
}

func (self Instance) SetAnimOffsetMax(value Float.X) {
	class(self).SetParamMax(11, gd.Float(value))
}

func (self Instance) AnimOffsetCurve() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetParamCurve(11))
}

func (self Instance) SetAnimOffsetCurve(value [1]gdclass.Curve) {
	class(self).SetParamCurve(11, value)
}

//go:nosplit
func (self class) SetEmitting(emitting bool) { //gd:CPUParticles3D.set_emitting
	var frame = callframe.New()
	callframe.Arg(frame, emitting)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_emitting, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetAmount(amount gd.Int) { //gd:CPUParticles3D.set_amount
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_amount, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetLifetime(secs gd.Float) { //gd:CPUParticles3D.set_lifetime
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_lifetime, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetOneShot(enable bool) { //gd:CPUParticles3D.set_one_shot
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_one_shot, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetPreProcessTime(secs gd.Float) { //gd:CPUParticles3D.set_pre_process_time
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_pre_process_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetExplosivenessRatio(ratio gd.Float) { //gd:CPUParticles3D.set_explosiveness_ratio
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_explosiveness_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetRandomnessRatio(ratio gd.Float) { //gd:CPUParticles3D.set_randomness_ratio
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_randomness_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetVisibilityAabb(aabb gd.AABB) { //gd:CPUParticles3D.set_visibility_aabb
	var frame = callframe.New()
	callframe.Arg(frame, aabb)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_visibility_aabb, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetLifetimeRandomness(random gd.Float) { //gd:CPUParticles3D.set_lifetime_randomness
	var frame = callframe.New()
	callframe.Arg(frame, random)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_lifetime_randomness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetUseLocalCoordinates(enable bool) { //gd:CPUParticles3D.set_use_local_coordinates
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_use_local_coordinates, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFixedFps(fps gd.Int) { //gd:CPUParticles3D.set_fixed_fps
	var frame = callframe.New()
	callframe.Arg(frame, fps)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_fixed_fps, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFractionalDelta(enable bool) { //gd:CPUParticles3D.set_fractional_delta
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_fractional_delta, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetSpeedScale(scale gd.Float) { //gd:CPUParticles3D.set_speed_scale
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_speed_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsEmitting() bool { //gd:CPUParticles3D.is_emitting
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_is_emitting, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetAmount() gd.Int { //gd:CPUParticles3D.get_amount
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_amount, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetLifetime() gd.Float { //gd:CPUParticles3D.get_lifetime
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_lifetime, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetOneShot() bool { //gd:CPUParticles3D.get_one_shot
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_one_shot, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetPreProcessTime() gd.Float { //gd:CPUParticles3D.get_pre_process_time
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_pre_process_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetExplosivenessRatio() gd.Float { //gd:CPUParticles3D.get_explosiveness_ratio
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_explosiveness_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetRandomnessRatio() gd.Float { //gd:CPUParticles3D.get_randomness_ratio
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_randomness_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetVisibilityAabb() gd.AABB { //gd:CPUParticles3D.get_visibility_aabb
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_visibility_aabb, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetLifetimeRandomness() gd.Float { //gd:CPUParticles3D.get_lifetime_randomness
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_lifetime_randomness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetUseLocalCoordinates() bool { //gd:CPUParticles3D.get_use_local_coordinates
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_use_local_coordinates, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetFixedFps() gd.Int { //gd:CPUParticles3D.get_fixed_fps
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_fixed_fps, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetFractionalDelta() bool { //gd:CPUParticles3D.get_fractional_delta
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_fractional_delta, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetSpeedScale() gd.Float { //gd:CPUParticles3D.get_speed_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_speed_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawOrder(order gdclass.CPUParticles3DDrawOrder) { //gd:CPUParticles3D.set_draw_order
	var frame = callframe.New()
	callframe.Arg(frame, order)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_draw_order, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDrawOrder() gdclass.CPUParticles3DDrawOrder { //gd:CPUParticles3D.get_draw_order
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CPUParticles3DDrawOrder](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_draw_order, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMesh(mesh [1]gdclass.Mesh) { //gd:CPUParticles3D.set_mesh
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMesh() [1]gdclass.Mesh { //gd:CPUParticles3D.get_mesh
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Mesh{gd.PointerWithOwnershipTransferredToGo[gdclass.Mesh](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Restarts the particle emitter.
*/
//go:nosplit
func (self class) Restart() { //gd:CPUParticles3D.restart
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_restart, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetDirection(direction gd.Vector3) { //gd:CPUParticles3D.set_direction
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDirection() gd.Vector3 { //gd:CPUParticles3D.get_direction
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpread(degrees gd.Float) { //gd:CPUParticles3D.set_spread
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_spread, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpread() gd.Float { //gd:CPUParticles3D.get_spread
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_spread, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlatness(amount gd.Float) { //gd:CPUParticles3D.set_flatness
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_flatness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFlatness() gd.Float { //gd:CPUParticles3D.get_flatness
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_flatness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the minimum value for the given parameter.
*/
//go:nosplit
func (self class) SetParamMin(param gdclass.CPUParticles3DParameter, value gd.Float) { //gd:CPUParticles3D.set_param_min
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_param_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the minimum value range for the given parameter.
*/
//go:nosplit
func (self class) GetParamMin(param gdclass.CPUParticles3DParameter) gd.Float { //gd:CPUParticles3D.get_param_min
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_param_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the maximum value for the given parameter.
*/
//go:nosplit
func (self class) SetParamMax(param gdclass.CPUParticles3DParameter, value gd.Float) { //gd:CPUParticles3D.set_param_max
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_param_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the maximum value range for the given parameter.
*/
//go:nosplit
func (self class) GetParamMax(param gdclass.CPUParticles3DParameter) gd.Float { //gd:CPUParticles3D.get_param_max
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_param_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [Curve] of the parameter specified by [enum Parameter].
*/
//go:nosplit
func (self class) SetParamCurve(param gdclass.CPUParticles3DParameter, curve [1]gdclass.Curve) { //gd:CPUParticles3D.set_param_curve
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_param_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [Curve] of the parameter specified by [enum Parameter].
*/
//go:nosplit
func (self class) GetParamCurve(param gdclass.CPUParticles3DParameter) [1]gdclass.Curve { //gd:CPUParticles3D.get_param_curve
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_param_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Curve{gd.PointerWithOwnershipTransferredToGo[gdclass.Curve](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColor(color gd.Color) { //gd:CPUParticles3D.set_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetColor() gd.Color { //gd:CPUParticles3D.get_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColorRamp(ramp [1]gdclass.Gradient) { //gd:CPUParticles3D.set_color_ramp
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(ramp[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_color_ramp, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetColorRamp() [1]gdclass.Gradient { //gd:CPUParticles3D.get_color_ramp
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_color_ramp, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Gradient{gd.PointerWithOwnershipTransferredToGo[gdclass.Gradient](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColorInitialRamp(ramp [1]gdclass.Gradient) { //gd:CPUParticles3D.set_color_initial_ramp
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(ramp[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_color_initial_ramp, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetColorInitialRamp() [1]gdclass.Gradient { //gd:CPUParticles3D.get_color_initial_ramp
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_color_initial_ramp, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Gradient{gd.PointerWithOwnershipTransferredToGo[gdclass.Gradient](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Enables or disables the given particle flag (see [enum ParticleFlags] for options).
*/
//go:nosplit
func (self class) SetParticleFlag(particle_flag gdclass.CPUParticles3DParticleFlags, enable bool) { //gd:CPUParticles3D.set_particle_flag
	var frame = callframe.New()
	callframe.Arg(frame, particle_flag)
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_particle_flag, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the enabled state of the given particle flag (see [enum ParticleFlags] for options).
*/
//go:nosplit
func (self class) GetParticleFlag(particle_flag gdclass.CPUParticles3DParticleFlags) bool { //gd:CPUParticles3D.get_particle_flag
	var frame = callframe.New()
	callframe.Arg(frame, particle_flag)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_particle_flag, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionShape(shape gdclass.CPUParticles3DEmissionShape) { //gd:CPUParticles3D.set_emission_shape
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_emission_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionShape() gdclass.CPUParticles3DEmissionShape { //gd:CPUParticles3D.get_emission_shape
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CPUParticles3DEmissionShape](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_emission_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionSphereRadius(radius gd.Float) { //gd:CPUParticles3D.set_emission_sphere_radius
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_emission_sphere_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionSphereRadius() gd.Float { //gd:CPUParticles3D.get_emission_sphere_radius
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_emission_sphere_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionBoxExtents(extents gd.Vector3) { //gd:CPUParticles3D.set_emission_box_extents
	var frame = callframe.New()
	callframe.Arg(frame, extents)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_emission_box_extents, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionBoxExtents() gd.Vector3 { //gd:CPUParticles3D.get_emission_box_extents
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_emission_box_extents, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionPoints(array gd.PackedVector3Array) { //gd:CPUParticles3D.set_emission_points
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(array))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_emission_points, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionPoints() gd.PackedVector3Array { //gd:CPUParticles3D.get_emission_points
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_emission_points, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedVector3Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionNormals(array gd.PackedVector3Array) { //gd:CPUParticles3D.set_emission_normals
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(array))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_emission_normals, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionNormals() gd.PackedVector3Array { //gd:CPUParticles3D.get_emission_normals
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_emission_normals, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedVector3Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionColors(array gd.PackedColorArray) { //gd:CPUParticles3D.set_emission_colors
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(array))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_emission_colors, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionColors() gd.PackedColorArray { //gd:CPUParticles3D.get_emission_colors
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_emission_colors, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedColorArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionRingAxis(axis gd.Vector3) { //gd:CPUParticles3D.set_emission_ring_axis
	var frame = callframe.New()
	callframe.Arg(frame, axis)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_emission_ring_axis, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionRingAxis() gd.Vector3 { //gd:CPUParticles3D.get_emission_ring_axis
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_emission_ring_axis, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionRingHeight(height gd.Float) { //gd:CPUParticles3D.set_emission_ring_height
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_emission_ring_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionRingHeight() gd.Float { //gd:CPUParticles3D.get_emission_ring_height
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_emission_ring_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionRingRadius(radius gd.Float) { //gd:CPUParticles3D.set_emission_ring_radius
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_emission_ring_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionRingRadius() gd.Float { //gd:CPUParticles3D.get_emission_ring_radius
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_emission_ring_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionRingInnerRadius(inner_radius gd.Float) { //gd:CPUParticles3D.set_emission_ring_inner_radius
	var frame = callframe.New()
	callframe.Arg(frame, inner_radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_emission_ring_inner_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionRingInnerRadius() gd.Float { //gd:CPUParticles3D.get_emission_ring_inner_radius
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_emission_ring_inner_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetGravity() gd.Vector3 { //gd:CPUParticles3D.get_gravity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravity(accel_vec gd.Vector3) { //gd:CPUParticles3D.set_gravity
	var frame = callframe.New()
	callframe.Arg(frame, accel_vec)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_gravity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSplitScale() bool { //gd:CPUParticles3D.get_split_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_split_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSplitScale(split_scale bool) { //gd:CPUParticles3D.set_split_scale
	var frame = callframe.New()
	callframe.Arg(frame, split_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_split_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetScaleCurveX() [1]gdclass.Curve { //gd:CPUParticles3D.get_scale_curve_x
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_scale_curve_x, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Curve{gd.PointerWithOwnershipTransferredToGo[gdclass.Curve](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScaleCurveX(scale_curve [1]gdclass.Curve) { //gd:CPUParticles3D.set_scale_curve_x
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scale_curve[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_scale_curve_x, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetScaleCurveY() [1]gdclass.Curve { //gd:CPUParticles3D.get_scale_curve_y
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_scale_curve_y, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Curve{gd.PointerWithOwnershipTransferredToGo[gdclass.Curve](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScaleCurveY(scale_curve [1]gdclass.Curve) { //gd:CPUParticles3D.set_scale_curve_y
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scale_curve[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_scale_curve_y, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetScaleCurveZ() [1]gdclass.Curve { //gd:CPUParticles3D.get_scale_curve_z
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_get_scale_curve_z, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Curve{gd.PointerWithOwnershipTransferredToGo[gdclass.Curve](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScaleCurveZ(scale_curve [1]gdclass.Curve) { //gd:CPUParticles3D.set_scale_curve_z
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scale_curve[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_set_scale_curve_z, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets this node's properties to match a given [GPUParticles3D] node with an assigned [ParticleProcessMaterial].
*/
//go:nosplit
func (self class) ConvertFromParticles(particles [1]gdclass.Node) { //gd:CPUParticles3D.convert_from_particles
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(particles[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CPUParticles3D.Bind_convert_from_particles, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnFinished(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("finished"), gd.NewCallable(cb), 0)
}

func (self class) AsCPUParticles3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCPUParticles3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsGeometryInstance3D() GeometryInstance3D.Advanced {
	return *((*GeometryInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGeometryInstance3D() GeometryInstance3D.Instance {
	return *((*GeometryInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualInstance3D() VisualInstance3D.Advanced {
	return *((*VisualInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualInstance3D() VisualInstance3D.Instance {
	return *((*VisualInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(GeometryInstance3D.Advanced(self.AsGeometryInstance3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(GeometryInstance3D.Instance(self.AsGeometryInstance3D()), name)
	}
}
func init() {
	gdclass.Register("CPUParticles3D", func(ptr gd.Object) any {
		return [1]gdclass.CPUParticles3D{*(*gdclass.CPUParticles3D)(unsafe.Pointer(&ptr))}
	})
}

type DrawOrder = gdclass.CPUParticles3DDrawOrder //gd:CPUParticles3D.DrawOrder

const (
	/*Particles are drawn in the order emitted.*/
	DrawOrderIndex DrawOrder = 0
	/*Particles are drawn in order of remaining lifetime. In other words, the particle with the highest lifetime is drawn at the front.*/
	DrawOrderLifetime DrawOrder = 1
	/*Particles are drawn in order of depth.*/
	DrawOrderViewDepth DrawOrder = 2
)

type Parameter = gdclass.CPUParticles3DParameter //gd:CPUParticles3D.Parameter

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

type ParticleFlags = gdclass.CPUParticles3DParticleFlags //gd:CPUParticles3D.ParticleFlags

const (
	/*Use with [method set_particle_flag] to set [member particle_flag_align_y].*/
	ParticleFlagAlignYToVelocity ParticleFlags = 0
	/*Use with [method set_particle_flag] to set [member particle_flag_rotate_y].*/
	ParticleFlagRotateY ParticleFlags = 1
	/*Use with [method set_particle_flag] to set [member particle_flag_disable_z].*/
	ParticleFlagDisableZ ParticleFlags = 2
	/*Represents the size of the [enum ParticleFlags] enum.*/
	ParticleFlagMax ParticleFlags = 3
)

type EmissionShape = gdclass.CPUParticles3DEmissionShape //gd:CPUParticles3D.EmissionShape

const (
	/*All particles will be emitted from a single point.*/
	EmissionShapePoint EmissionShape = 0
	/*Particles will be emitted in the volume of a sphere.*/
	EmissionShapeSphere EmissionShape = 1
	/*Particles will be emitted on the surface of a sphere.*/
	EmissionShapeSphereSurface EmissionShape = 2
	/*Particles will be emitted in the volume of a box.*/
	EmissionShapeBox EmissionShape = 3
	/*Particles will be emitted at a position chosen randomly among [member emission_points]. Particle color will be modulated by [member emission_colors].*/
	EmissionShapePoints EmissionShape = 4
	/*Particles will be emitted at a position chosen randomly among [member emission_points]. Particle velocity and rotation will be set based on [member emission_normals]. Particle color will be modulated by [member emission_colors].*/
	EmissionShapeDirectedPoints EmissionShape = 5
	/*Particles will be emitted in a ring or cylinder.*/
	EmissionShapeRing EmissionShape = 6
	/*Represents the size of the [enum EmissionShape] enum.*/
	EmissionShapeMax EmissionShape = 7
)
