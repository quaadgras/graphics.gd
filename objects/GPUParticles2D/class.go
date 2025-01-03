package GPUParticles2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Node2D"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Rect2"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Transform2D"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Color"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
2D particle node used to create a variety of particle systems and effects. [GPUParticles2D] features an emitter that generates some number of particles at a given rate.
Use the [member process_material] property to add a [ParticleProcessMaterial] to configure particle appearance and behavior. Alternatively, you can add a [ShaderMaterial] which will be applied to all particles.
2D particles can optionally collide with [LightOccluder2D], but they don't collide with [PhysicsBody2D] nodes.
*/
type Instance [1]classdb.GPUParticles2D
type Any interface {
	gd.IsClass
	AsGPUParticles2D() Instance
}

/*
Returns a rectangle containing the positions of all existing particles.
[b]Note:[/b] When using threaded rendering this method synchronizes the rendering thread. Calling it often may have a negative impact on performance.
*/
func (self Instance) CaptureRect() Rect2.PositionSize {
	return Rect2.PositionSize(class(self).CaptureRect())
}

/*
Restarts the particle emission cycle, clearing existing particles. To avoid particles vanishing from the viewport, wait for the [signal finished] signal before calling.
[b]Note:[/b] The [signal finished] signal is only emitted by [member one_shot] emitters.
*/
func (self Instance) Restart() {
	class(self).Restart()
}

/*
Emits a single particle. Whether [param xform], [param velocity], [param color] and [param custom] are applied depends on the value of [param flags]. See [enum EmitFlags].
The default ParticleProcessMaterial will overwrite [param color] and use the contents of [param custom] as [code](rotation, age, animation, lifetime)[/code].
*/
func (self Instance) EmitParticle(xform Transform2D.OriginXY, velocity Vector2.XY, color Color.RGBA, custom Color.RGBA, flags int) {
	class(self).EmitParticle(gd.Transform2D(xform), gd.Vector2(velocity), gd.Color(color), gd.Color(custom), gd.Int(flags))
}

/*
Sets this node's properties to match a given [CPUParticles2D] node.
*/
func (self Instance) ConvertFromParticles(particles objects.Node) {
	class(self).ConvertFromParticles(particles)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.GPUParticles2D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GPUParticles2D"))
	return Instance{*(*classdb.GPUParticles2D)(unsafe.Pointer(&object))}
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

func (self Instance) SubEmitter() Path.String {
	return Path.String(class(self).GetSubEmitter().String())
}

func (self Instance) SetSubEmitter(value Path.String) {
	class(self).SetSubEmitter(gd.NewString(string(value)).NodePath())
}

func (self Instance) ProcessMaterial() objects.Material {
	return objects.Material(class(self).GetProcessMaterial())
}

func (self Instance) SetProcessMaterial(value objects.Material) {
	class(self).SetProcessMaterial(value)
}

func (self Instance) Texture() objects.Texture2D {
	return objects.Texture2D(class(self).GetTexture())
}

func (self Instance) SetTexture(value objects.Texture2D) {
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

func (self Instance) DrawOrder() classdb.GPUParticles2DDrawOrder {
	return classdb.GPUParticles2DDrawOrder(class(self).GetDrawOrder())
}

func (self Instance) SetDrawOrder(value classdb.GPUParticles2DDrawOrder) {
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
func (self class) SetEmitting(emitting bool) {
	var frame = callframe.New()
	callframe.Arg(frame, emitting)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_emitting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetAmount(amount gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetLifetime(secs gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_lifetime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetOneShot(secs bool) {
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_one_shot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetPreProcessTime(secs gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_pre_process_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetExplosivenessRatio(ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_explosiveness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetRandomnessRatio(ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_randomness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetVisibilityRect(visibility_rect gd.Rect2) {
	var frame = callframe.New()
	callframe.Arg(frame, visibility_rect)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_visibility_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetUseLocalCoordinates(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_use_local_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetFixedFps(fps gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, fps)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_fixed_fps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetFractionalDelta(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_fractional_delta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetInterpolate(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_interpolate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetProcessMaterial(material objects.Material) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_process_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetSpeedScale(scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetCollisionBaseSize(size gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_collision_base_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetInterpToEnd(interp gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, interp)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_interp_to_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsEmitting() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_is_emitting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetAmount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetLifetime() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_lifetime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetOneShot() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_one_shot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetPreProcessTime() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_pre_process_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetExplosivenessRatio() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_explosiveness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetRandomnessRatio() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_randomness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetVisibilityRect() gd.Rect2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_visibility_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetUseLocalCoordinates() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_use_local_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetFixedFps() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_fixed_fps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetFractionalDelta() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_fractional_delta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetInterpolate() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_interpolate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetProcessMaterial() objects.Material {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_process_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Material{gd.PointerWithOwnershipTransferredToGo[classdb.Material](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetSpeedScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetCollisionBaseSize() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_collision_base_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetInterpToEnd() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_interp_to_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawOrder(order classdb.GPUParticles2DDrawOrder) {
	var frame = callframe.New()
	callframe.Arg(frame, order)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_draw_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDrawOrder() classdb.GPUParticles2DDrawOrder {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GPUParticles2DDrawOrder](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_draw_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTexture(texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{gd.PointerWithOwnershipTransferredToGo[classdb.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns a rectangle containing the positions of all existing particles.
[b]Note:[/b] When using threaded rendering this method synchronizes the rendering thread. Calling it often may have a negative impact on performance.
*/
//go:nosplit
func (self class) CaptureRect() gd.Rect2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_capture_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_restart, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetSubEmitter(path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_sub_emitter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubEmitter() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_sub_emitter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

/*
Emits a single particle. Whether [param xform], [param velocity], [param color] and [param custom] are applied depends on the value of [param flags]. See [enum EmitFlags].
The default ParticleProcessMaterial will overwrite [param color] and use the contents of [param custom] as [code](rotation, age, animation, lifetime)[/code].
*/
//go:nosplit
func (self class) EmitParticle(xform gd.Transform2D, velocity gd.Vector2, color gd.Color, custom gd.Color, flags gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, xform)
	callframe.Arg(frame, velocity)
	callframe.Arg(frame, color)
	callframe.Arg(frame, custom)
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_emit_particle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetTrailEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_trail_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetTrailLifetime(secs gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_trail_lifetime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsTrailEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_is_trail_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTrailLifetime() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_trail_lifetime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTrailSections(sections gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, sections)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_trail_sections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTrailSections() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_trail_sections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTrailSectionSubdivisions(subdivisions gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, subdivisions)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_trail_section_subdivisions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTrailSectionSubdivisions() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_trail_section_subdivisions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets this node's properties to match a given [CPUParticles2D] node.
*/
//go:nosplit
func (self class) ConvertFromParticles(particles objects.Node) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(particles[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_convert_from_particles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetAmountRatio(ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_set_amount_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAmountRatio() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticles2D.Bind_get_amount_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnFinished(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("finished"), gd.NewCallable(cb), 0)
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
	classdb.Register("GPUParticles2D", func(ptr gd.Object) any {
		return [1]classdb.GPUParticles2D{*(*classdb.GPUParticles2D)(unsafe.Pointer(&ptr))}
	})
}

type DrawOrder = classdb.GPUParticles2DDrawOrder

const (
	/*Particles are drawn in the order emitted.*/
	DrawOrderIndex DrawOrder = 0
	/*Particles are drawn in order of remaining lifetime. In other words, the particle with the highest lifetime is drawn at the front.*/
	DrawOrderLifetime DrawOrder = 1
	/*Particles are drawn in reverse order of remaining lifetime. In other words, the particle with the lowest lifetime is drawn at the front.*/
	DrawOrderReverseLifetime DrawOrder = 2
)

type EmitFlags = classdb.GPUParticles2DEmitFlags

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
