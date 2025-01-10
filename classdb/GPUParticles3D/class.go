// Package GPUParticles3D provides methods for working with GPUParticles3D object instances.
package GPUParticles3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/GeometryInstance3D"
import "graphics.gd/classdb/VisualInstance3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/AABB"
import "graphics.gd/variant/NodePath"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Color"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
3D particle node used to create a variety of particle systems and effects. [GPUParticles3D] features an emitter that generates some number of particles at a given rate.
Use [member process_material] to add a [ParticleProcessMaterial] to configure particle appearance and behavior. Alternatively, you can add a [ShaderMaterial] which will be applied to all particles.
*/
type Instance [1]gdclass.GPUParticles3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGPUParticles3D() Instance
}

/*
Restarts the particle emission cycle, clearing existing particles. To avoid particles vanishing from the viewport, wait for the [signal finished] signal before calling.
[b]Note:[/b] The [signal finished] signal is only emitted by [member one_shot] emitters.
*/
func (self Instance) Restart() {
	class(self).Restart()
}

/*
Returns the axis-aligned bounding box that contains all the particles that are active in the current frame.
*/
func (self Instance) CaptureAabb() AABB.PositionSize {
	return AABB.PositionSize(class(self).CaptureAabb())
}

/*
Emits a single particle. Whether [param xform], [param velocity], [param color] and [param custom] are applied depends on the value of [param flags]. See [enum EmitFlags].
The default ParticleProcessMaterial will overwrite [param color] and use the contents of [param custom] as [code](rotation, age, animation, lifetime)[/code].
*/
func (self Instance) EmitParticle(xform Transform3D.BasisOrigin, velocity Vector3.XYZ, color Color.RGBA, custom Color.RGBA, flags int) {
	class(self).EmitParticle(gd.Transform3D(xform), gd.Vector3(velocity), gd.Color(color), gd.Color(custom), gd.Int(flags))
}

/*
Sets this node's properties to match a given [CPUParticles3D] node.
*/
func (self Instance) ConvertFromParticles(particles [1]gdclass.Node) {
	class(self).ConvertFromParticles(particles)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GPUParticles3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GPUParticles3D"))
	casted := Instance{*(*gdclass.GPUParticles3D)(unsafe.Pointer(&object))}
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

func (self Instance) AmountRatio() Float.X {
	return Float.X(Float.X(class(self).GetAmountRatio()))
}

func (self Instance) SetAmountRatio(value Float.X) {
	class(self).SetAmountRatio(gd.Float(value))
}

func (self Instance) SubEmitter() NodePath.String {
	return NodePath.String(class(self).GetSubEmitter().String())
}

func (self Instance) SetSubEmitter(value NodePath.String) {
	class(self).SetSubEmitter(gd.NewString(string(value)).NodePath())
}

func (self Instance) Lifetime() Float.X {
	return Float.X(Float.X(class(self).GetLifetime()))
}

func (self Instance) SetLifetime(value Float.X) {
	class(self).SetLifetime(gd.Float(value))
}

func (self Instance) InterpToEnd() Float.X {
	return Float.X(Float.X(class(self).GetInterpToEnd()))
}

func (self Instance) SetInterpToEnd(value Float.X) {
	class(self).SetInterpToEnd(gd.Float(value))
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

func (self Instance) FixedFps() int {
	return int(int(class(self).GetFixedFps()))
}

func (self Instance) SetFixedFps(value int) {
	class(self).SetFixedFps(gd.Int(value))
}

func (self Instance) Interpolate() bool {
	return bool(class(self).GetInterpolate())
}

func (self Instance) SetInterpolate(value bool) {
	class(self).SetInterpolate(value)
}

func (self Instance) FractDelta() bool {
	return bool(class(self).GetFractionalDelta())
}

func (self Instance) SetFractDelta(value bool) {
	class(self).SetFractionalDelta(value)
}

func (self Instance) CollisionBaseSize() Float.X {
	return Float.X(Float.X(class(self).GetCollisionBaseSize()))
}

func (self Instance) SetCollisionBaseSize(value Float.X) {
	class(self).SetCollisionBaseSize(gd.Float(value))
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

func (self Instance) DrawOrder() gdclass.GPUParticles3DDrawOrder {
	return gdclass.GPUParticles3DDrawOrder(class(self).GetDrawOrder())
}

func (self Instance) SetDrawOrder(value gdclass.GPUParticles3DDrawOrder) {
	class(self).SetDrawOrder(value)
}

func (self Instance) TransformAlign() gdclass.GPUParticles3DTransformAlign {
	return gdclass.GPUParticles3DTransformAlign(class(self).GetTransformAlign())
}

func (self Instance) SetTransformAlign(value gdclass.GPUParticles3DTransformAlign) {
	class(self).SetTransformAlign(value)
}

func (self Instance) TrailEnabled() bool {
	return bool(class(self).IsTrailEnabled())
}

func (self Instance) SetTrailEnabled(value bool) {
	class(self).SetTrailEnabled(value)
}

func (self Instance) TrailLifetime() Float.X {
	return Float.X(Float.X(class(self).GetTrailLifetime()))
}

func (self Instance) SetTrailLifetime(value Float.X) {
	class(self).SetTrailLifetime(gd.Float(value))
}

func (self Instance) ProcessMaterial() [1]gdclass.Material {
	return [1]gdclass.Material(class(self).GetProcessMaterial())
}

func (self Instance) SetProcessMaterial(value [1]gdclass.Material) {
	class(self).SetProcessMaterial(value)
}

func (self Instance) DrawPasses() int {
	return int(int(class(self).GetDrawPasses()))
}

func (self Instance) SetDrawPasses(value int) {
	class(self).SetDrawPasses(gd.Int(value))
}

func (self Instance) DrawPass1() [1]gdclass.Mesh {
	return [1]gdclass.Mesh(class(self).GetDrawPassMesh(0))
}

func (self Instance) SetDrawPass1(value [1]gdclass.Mesh) {
	class(self).SetDrawPassMesh(0, value)
}

func (self Instance) DrawPass2() [1]gdclass.Mesh {
	return [1]gdclass.Mesh(class(self).GetDrawPassMesh(1))
}

func (self Instance) SetDrawPass2(value [1]gdclass.Mesh) {
	class(self).SetDrawPassMesh(1, value)
}

func (self Instance) DrawPass3() [1]gdclass.Mesh {
	return [1]gdclass.Mesh(class(self).GetDrawPassMesh(2))
}

func (self Instance) SetDrawPass3(value [1]gdclass.Mesh) {
	class(self).SetDrawPassMesh(2, value)
}

func (self Instance) DrawPass4() [1]gdclass.Mesh {
	return [1]gdclass.Mesh(class(self).GetDrawPassMesh(3))
}

func (self Instance) SetDrawPass4(value [1]gdclass.Mesh) {
	class(self).SetDrawPassMesh(3, value)
}

func (self Instance) DrawSkin() [1]gdclass.Skin {
	return [1]gdclass.Skin(class(self).GetSkin())
}

func (self Instance) SetDrawSkin(value [1]gdclass.Skin) {
	class(self).SetSkin(value)
}

//go:nosplit
func (self class) SetEmitting(emitting bool) {
	var frame = callframe.New()
	callframe.Arg(frame, emitting)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_emitting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetAmount(amount gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetLifetime(secs gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_lifetime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetOneShot(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_one_shot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetPreProcessTime(secs gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_pre_process_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetExplosivenessRatio(ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_explosiveness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetRandomnessRatio(ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_randomness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetVisibilityAabb(aabb gd.AABB) {
	var frame = callframe.New()
	callframe.Arg(frame, aabb)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_visibility_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetUseLocalCoordinates(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_use_local_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetFixedFps(fps gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, fps)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_fixed_fps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetFractionalDelta(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_fractional_delta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetInterpolate(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_interpolate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetProcessMaterial(material [1]gdclass.Material) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_process_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetSpeedScale(scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetCollisionBaseSize(size gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_collision_base_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetInterpToEnd(interp gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, interp)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_interp_to_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsEmitting() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_is_emitting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetAmount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetLifetime() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_lifetime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetOneShot() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_one_shot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetPreProcessTime() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_pre_process_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetExplosivenessRatio() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_explosiveness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetRandomnessRatio() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_randomness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetVisibilityAabb() gd.AABB {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_visibility_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetUseLocalCoordinates() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_use_local_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetFixedFps() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_fixed_fps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetFractionalDelta() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_fractional_delta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetInterpolate() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_interpolate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetProcessMaterial() [1]gdclass.Material {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_process_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Material{gd.PointerWithOwnershipTransferredToGo[gdclass.Material](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetSpeedScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetCollisionBaseSize() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_collision_base_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetInterpToEnd() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_interp_to_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawOrder(order gdclass.GPUParticles3DDrawOrder) {
	var frame = callframe.New()
	callframe.Arg(frame, order)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_draw_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDrawOrder() gdclass.GPUParticles3DDrawOrder {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GPUParticles3DDrawOrder](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_draw_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawPasses(passes gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, passes)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_draw_passes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the [Mesh] that is drawn at index [param pass].
*/
//go:nosplit
func (self class) SetDrawPassMesh(pass gd.Int, mesh [1]gdclass.Mesh) {
	var frame = callframe.New()
	callframe.Arg(frame, pass)
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_draw_pass_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDrawPasses() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_draw_passes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Mesh] that is drawn at index [param pass].
*/
//go:nosplit
func (self class) GetDrawPassMesh(pass gd.Int) [1]gdclass.Mesh {
	var frame = callframe.New()
	callframe.Arg(frame, pass)
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_draw_pass_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Mesh{gd.PointerWithOwnershipTransferredToGo[gdclass.Mesh](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkin(skin [1]gdclass.Skin) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(skin[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_skin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkin() [1]gdclass.Skin {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_skin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Skin{gd.PointerWithOwnershipTransferredToGo[gdclass.Skin](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Restarts the particle emission cycle, clearing existing particles. To avoid particles vanishing from the viewport, wait for the [signal finished] signal before calling.
[b]Note:[/b] The [signal finished] signal is only emitted by [member one_shot] emitters.
*/
//go:nosplit
func (self class) Restart() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_restart, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the axis-aligned bounding box that contains all the particles that are active in the current frame.
*/
//go:nosplit
func (self class) CaptureAabb() gd.AABB {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_capture_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubEmitter(path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_sub_emitter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubEmitter() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_sub_emitter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

/*
Emits a single particle. Whether [param xform], [param velocity], [param color] and [param custom] are applied depends on the value of [param flags]. See [enum EmitFlags].
The default ParticleProcessMaterial will overwrite [param color] and use the contents of [param custom] as [code](rotation, age, animation, lifetime)[/code].
*/
//go:nosplit
func (self class) EmitParticle(xform gd.Transform3D, velocity gd.Vector3, color gd.Color, custom gd.Color, flags gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, xform)
	callframe.Arg(frame, velocity)
	callframe.Arg(frame, color)
	callframe.Arg(frame, custom)
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_emit_particle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetTrailEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_trail_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetTrailLifetime(secs gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_trail_lifetime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsTrailEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_is_trail_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTrailLifetime() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_trail_lifetime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTransformAlign(align gdclass.GPUParticles3DTransformAlign) {
	var frame = callframe.New()
	callframe.Arg(frame, align)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_transform_align, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTransformAlign() gdclass.GPUParticles3DTransformAlign {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GPUParticles3DTransformAlign](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_transform_align, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets this node's properties to match a given [CPUParticles3D] node.
*/
//go:nosplit
func (self class) ConvertFromParticles(particles [1]gdclass.Node) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(particles[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_convert_from_particles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetAmountRatio(ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_set_amount_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAmountRatio() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles3D.Bind_get_amount_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnFinished(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("finished"), gd.NewCallable(cb), 0)
}

func (self class) AsGPUParticles3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGPUParticles3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("GPUParticles3D", func(ptr gd.Object) any {
		return [1]gdclass.GPUParticles3D{*(*gdclass.GPUParticles3D)(unsafe.Pointer(&ptr))}
	})
}

type DrawOrder = gdclass.GPUParticles3DDrawOrder

const (
	/*Particles are drawn in the order emitted.*/
	DrawOrderIndex DrawOrder = 0
	/*Particles are drawn in order of remaining lifetime. In other words, the particle with the highest lifetime is drawn at the front.*/
	DrawOrderLifetime DrawOrder = 1
	/*Particles are drawn in reverse order of remaining lifetime. In other words, the particle with the lowest lifetime is drawn at the front.*/
	DrawOrderReverseLifetime DrawOrder = 2
	/*Particles are drawn in order of depth.*/
	DrawOrderViewDepth DrawOrder = 3
)

type EmitFlags = gdclass.GPUParticles3DEmitFlags

const (
	/*Particle starts at the specified position.*/
	EmitFlagPosition EmitFlags = 1
	/*Particle starts with specified rotation and scale.*/
	EmitFlagRotationScale EmitFlags = 2
	/*Particle starts with the specified velocity vector, which defines the emission direction and speed.*/
	EmitFlagVelocity EmitFlags = 4
	/*Particle starts with specified color.*/
	EmitFlagColor EmitFlags = 8
	/*Particle starts with specified [code]CUSTOM[/code] data.*/
	EmitFlagCustom EmitFlags = 16
)

type TransformAlign = gdclass.GPUParticles3DTransformAlign

const (
	TransformAlignDisabled              TransformAlign = 0
	TransformAlignZBillboard            TransformAlign = 1
	TransformAlignYToVelocity           TransformAlign = 2
	TransformAlignZBillboardYToVelocity TransformAlign = 3
)
