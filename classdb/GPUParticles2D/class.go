// Package GPUParticles2D provides methods for working with GPUParticles2D object instances.
package GPUParticles2D

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
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Rect2"
import "graphics.gd/variant/NodePath"
import "graphics.gd/variant/Transform2D"
import "graphics.gd/variant/Vector2"
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
2D particle node used to create a variety of particle systems and effects. [GPUParticles2D] features an emitter that generates some number of particles at a given rate.
Use the [member process_material] property to add a [ParticleProcessMaterial] to configure particle appearance and behavior. Alternatively, you can add a [ShaderMaterial] which will be applied to all particles.
2D particles can optionally collide with [LightOccluder2D], but they don't collide with [PhysicsBody2D] nodes.
*/
type Instance [1]gdclass.GPUParticles2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGPUParticles2D() Instance
}

/*
Returns a rectangle containing the positions of all existing particles.
[b]Note:[/b] When using threaded rendering this method synchronizes the rendering thread. Calling it often may have a negative impact on performance.
*/
func (self Instance) CaptureRect() Rect2.PositionSize { //gd:GPUParticles2D.capture_rect
	return Rect2.PositionSize(class(self).CaptureRect())
}

/*
Restarts the particle emission cycle, clearing existing particles. To avoid particles vanishing from the viewport, wait for the [signal finished] signal before calling.
[b]Note:[/b] The [signal finished] signal is only emitted by [member one_shot] emitters.
*/
func (self Instance) Restart() { //gd:GPUParticles2D.restart
	class(self).Restart()
}

/*
Emits a single particle. Whether [param xform], [param velocity], [param color] and [param custom] are applied depends on the value of [param flags]. See [enum EmitFlags].
The default ParticleProcessMaterial will overwrite [param color] and use the contents of [param custom] as [code](rotation, age, animation, lifetime)[/code].
*/
func (self Instance) EmitParticle(xform Transform2D.OriginXY, velocity Vector2.XY, color Color.RGBA, custom Color.RGBA, flags int) { //gd:GPUParticles2D.emit_particle
	class(self).EmitParticle(gd.Transform2D(xform), gd.Vector2(velocity), gd.Color(color), gd.Color(custom), gd.Int(flags))
}

/*
Sets this node's properties to match a given [CPUParticles2D] node.
*/
func (self Instance) ConvertFromParticles(particles [1]gdclass.Node) { //gd:GPUParticles2D.convert_from_particles
	class(self).ConvertFromParticles(particles)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GPUParticles2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GPUParticles2D"))
	casted := Instance{*(*gdclass.GPUParticles2D)(unsafe.Pointer(&object))}
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

func (self Instance) ProcessMaterial() [1]gdclass.Material {
	return [1]gdclass.Material(class(self).GetProcessMaterial())
}

func (self Instance) SetProcessMaterial(value [1]gdclass.Material) {
	class(self).SetProcessMaterial(value)
}

func (self Instance) Texture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture())
}

func (self Instance) SetTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(value)
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

func (self Instance) InterpToEnd() Float.X {
	return Float.X(Float.X(class(self).GetInterpToEnd()))
}

func (self Instance) SetInterpToEnd(value Float.X) {
	class(self).SetInterpToEnd(gd.Float(value))
}

func (self Instance) CollisionBaseSize() Float.X {
	return Float.X(Float.X(class(self).GetCollisionBaseSize()))
}

func (self Instance) SetCollisionBaseSize(value Float.X) {
	class(self).SetCollisionBaseSize(gd.Float(value))
}

func (self Instance) VisibilityRect() Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetVisibilityRect())
}

func (self Instance) SetVisibilityRect(value Rect2.PositionSize) {
	class(self).SetVisibilityRect(gd.Rect2(value))
}

func (self Instance) LocalCoords() bool {
	return bool(class(self).GetUseLocalCoordinates())
}

func (self Instance) SetLocalCoords(value bool) {
	class(self).SetUseLocalCoordinates(value)
}

func (self Instance) DrawOrder() gdclass.GPUParticles2DDrawOrder {
	return gdclass.GPUParticles2DDrawOrder(class(self).GetDrawOrder())
}

func (self Instance) SetDrawOrder(value gdclass.GPUParticles2DDrawOrder) {
	class(self).SetDrawOrder(value)
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

func (self Instance) TrailSections() int {
	return int(int(class(self).GetTrailSections()))
}

func (self Instance) SetTrailSections(value int) {
	class(self).SetTrailSections(gd.Int(value))
}

func (self Instance) TrailSectionSubdivisions() int {
	return int(int(class(self).GetTrailSectionSubdivisions()))
}

func (self Instance) SetTrailSectionSubdivisions(value int) {
	class(self).SetTrailSectionSubdivisions(gd.Int(value))
}

//go:nosplit
func (self class) SetEmitting(emitting bool) { //gd:GPUParticles2D.set_emitting
	var frame = callframe.New()
	callframe.Arg(frame, emitting)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_emitting, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetAmount(amount gd.Int) { //gd:GPUParticles2D.set_amount
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_amount, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetLifetime(secs gd.Float) { //gd:GPUParticles2D.set_lifetime
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_lifetime, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetOneShot(secs bool) { //gd:GPUParticles2D.set_one_shot
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_one_shot, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetPreProcessTime(secs gd.Float) { //gd:GPUParticles2D.set_pre_process_time
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_pre_process_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetExplosivenessRatio(ratio gd.Float) { //gd:GPUParticles2D.set_explosiveness_ratio
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_explosiveness_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetRandomnessRatio(ratio gd.Float) { //gd:GPUParticles2D.set_randomness_ratio
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_randomness_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetVisibilityRect(visibility_rect gd.Rect2) { //gd:GPUParticles2D.set_visibility_rect
	var frame = callframe.New()
	callframe.Arg(frame, visibility_rect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_visibility_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetUseLocalCoordinates(enable bool) { //gd:GPUParticles2D.set_use_local_coordinates
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_use_local_coordinates, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFixedFps(fps gd.Int) { //gd:GPUParticles2D.set_fixed_fps
	var frame = callframe.New()
	callframe.Arg(frame, fps)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_fixed_fps, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFractionalDelta(enable bool) { //gd:GPUParticles2D.set_fractional_delta
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_fractional_delta, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetInterpolate(enable bool) { //gd:GPUParticles2D.set_interpolate
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_interpolate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetProcessMaterial(material [1]gdclass.Material) { //gd:GPUParticles2D.set_process_material
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_process_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetSpeedScale(scale gd.Float) { //gd:GPUParticles2D.set_speed_scale
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_speed_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetCollisionBaseSize(size gd.Float) { //gd:GPUParticles2D.set_collision_base_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_collision_base_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetInterpToEnd(interp gd.Float) { //gd:GPUParticles2D.set_interp_to_end
	var frame = callframe.New()
	callframe.Arg(frame, interp)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_interp_to_end, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsEmitting() bool { //gd:GPUParticles2D.is_emitting
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_is_emitting, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetAmount() gd.Int { //gd:GPUParticles2D.get_amount
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_amount, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetLifetime() gd.Float { //gd:GPUParticles2D.get_lifetime
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_lifetime, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetOneShot() bool { //gd:GPUParticles2D.get_one_shot
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_one_shot, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetPreProcessTime() gd.Float { //gd:GPUParticles2D.get_pre_process_time
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_pre_process_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetExplosivenessRatio() gd.Float { //gd:GPUParticles2D.get_explosiveness_ratio
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_explosiveness_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetRandomnessRatio() gd.Float { //gd:GPUParticles2D.get_randomness_ratio
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_randomness_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetVisibilityRect() gd.Rect2 { //gd:GPUParticles2D.get_visibility_rect
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_visibility_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetUseLocalCoordinates() bool { //gd:GPUParticles2D.get_use_local_coordinates
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_use_local_coordinates, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetFixedFps() gd.Int { //gd:GPUParticles2D.get_fixed_fps
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_fixed_fps, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetFractionalDelta() bool { //gd:GPUParticles2D.get_fractional_delta
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_fractional_delta, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetInterpolate() bool { //gd:GPUParticles2D.get_interpolate
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_interpolate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetProcessMaterial() [1]gdclass.Material { //gd:GPUParticles2D.get_process_material
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_process_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Material{gd.PointerWithOwnershipTransferredToGo[gdclass.Material](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetSpeedScale() gd.Float { //gd:GPUParticles2D.get_speed_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_speed_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetCollisionBaseSize() gd.Float { //gd:GPUParticles2D.get_collision_base_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_collision_base_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetInterpToEnd() gd.Float { //gd:GPUParticles2D.get_interp_to_end
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_interp_to_end, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawOrder(order gdclass.GPUParticles2DDrawOrder) { //gd:GPUParticles2D.set_draw_order
	var frame = callframe.New()
	callframe.Arg(frame, order)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_draw_order, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDrawOrder() gdclass.GPUParticles2DDrawOrder { //gd:GPUParticles2D.get_draw_order
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GPUParticles2DDrawOrder](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_draw_order, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTexture(texture [1]gdclass.Texture2D) { //gd:GPUParticles2D.set_texture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() [1]gdclass.Texture2D { //gd:GPUParticles2D.get_texture
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns a rectangle containing the positions of all existing particles.
[b]Note:[/b] When using threaded rendering this method synchronizes the rendering thread. Calling it often may have a negative impact on performance.
*/
//go:nosplit
func (self class) CaptureRect() gd.Rect2 { //gd:GPUParticles2D.capture_rect
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_capture_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Restarts the particle emission cycle, clearing existing particles. To avoid particles vanishing from the viewport, wait for the [signal finished] signal before calling.
[b]Note:[/b] The [signal finished] signal is only emitted by [member one_shot] emitters.
*/
//go:nosplit
func (self class) Restart() { //gd:GPUParticles2D.restart
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_restart, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetSubEmitter(path gd.NodePath) { //gd:GPUParticles2D.set_sub_emitter
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_sub_emitter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubEmitter() gd.NodePath { //gd:GPUParticles2D.get_sub_emitter
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_sub_emitter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

/*
Emits a single particle. Whether [param xform], [param velocity], [param color] and [param custom] are applied depends on the value of [param flags]. See [enum EmitFlags].
The default ParticleProcessMaterial will overwrite [param color] and use the contents of [param custom] as [code](rotation, age, animation, lifetime)[/code].
*/
//go:nosplit
func (self class) EmitParticle(xform gd.Transform2D, velocity gd.Vector2, color gd.Color, custom gd.Color, flags gd.Int) { //gd:GPUParticles2D.emit_particle
	var frame = callframe.New()
	callframe.Arg(frame, xform)
	callframe.Arg(frame, velocity)
	callframe.Arg(frame, color)
	callframe.Arg(frame, custom)
	callframe.Arg(frame, flags)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_emit_particle, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetTrailEnabled(enabled bool) { //gd:GPUParticles2D.set_trail_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_trail_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetTrailLifetime(secs gd.Float) { //gd:GPUParticles2D.set_trail_lifetime
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_trail_lifetime, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsTrailEnabled() bool { //gd:GPUParticles2D.is_trail_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_is_trail_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTrailLifetime() gd.Float { //gd:GPUParticles2D.get_trail_lifetime
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_trail_lifetime, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTrailSections(sections gd.Int) { //gd:GPUParticles2D.set_trail_sections
	var frame = callframe.New()
	callframe.Arg(frame, sections)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_trail_sections, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTrailSections() gd.Int { //gd:GPUParticles2D.get_trail_sections
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_trail_sections, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTrailSectionSubdivisions(subdivisions gd.Int) { //gd:GPUParticles2D.set_trail_section_subdivisions
	var frame = callframe.New()
	callframe.Arg(frame, subdivisions)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_trail_section_subdivisions, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTrailSectionSubdivisions() gd.Int { //gd:GPUParticles2D.get_trail_section_subdivisions
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_trail_section_subdivisions, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets this node's properties to match a given [CPUParticles2D] node.
*/
//go:nosplit
func (self class) ConvertFromParticles(particles [1]gdclass.Node) { //gd:GPUParticles2D.convert_from_particles
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(particles[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_convert_from_particles, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetAmountRatio(ratio gd.Float) { //gd:GPUParticles2D.set_amount_ratio
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_amount_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAmountRatio() gd.Float { //gd:GPUParticles2D.get_amount_ratio
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_amount_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnFinished(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("finished"), gd.NewCallable(cb), 0)
}

func (self class) AsGPUParticles2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGPUParticles2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Node2D.Advanced(self.AsNode2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Instance(self.AsNode2D()), name)
	}
}
func init() {
	gdclass.Register("GPUParticles2D", func(ptr gd.Object) any {
		return [1]gdclass.GPUParticles2D{*(*gdclass.GPUParticles2D)(unsafe.Pointer(&ptr))}
	})
}

type DrawOrder = gdclass.GPUParticles2DDrawOrder //gd:GPUParticles2D.DrawOrder

const (
	/*Particles are drawn in the order emitted.*/
	DrawOrderIndex DrawOrder = 0
	/*Particles are drawn in order of remaining lifetime. In other words, the particle with the highest lifetime is drawn at the front.*/
	DrawOrderLifetime DrawOrder = 1
	/*Particles are drawn in reverse order of remaining lifetime. In other words, the particle with the lowest lifetime is drawn at the front.*/
	DrawOrderReverseLifetime DrawOrder = 2
)

type EmitFlags = gdclass.GPUParticles2DEmitFlags //gd:GPUParticles2D.EmitFlags

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
