package GPUParticles2D

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
2D particle node used to create a variety of particle systems and effects. [GPUParticles2D] features an emitter that generates some number of particles at a given rate.
Use the [member process_material] property to add a [ParticleProcessMaterial] to configure particle appearance and behavior. Alternatively, you can add a [ShaderMaterial] which will be applied to all particles.
2D particles can optionally collide with [LightOccluder2D], but they don't collide with [PhysicsBody2D] nodes.

*/
type Go [1]classdb.GPUParticles2D

/*
Returns a rectangle containing the positions of all existing particles.
[b]Note:[/b] When using threaded rendering this method synchronizes the rendering thread. Calling it often may have a negative impact on performance.
*/
func (self Go) CaptureRect() gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(class(self).CaptureRect())
}

/*
Restarts the particle emission cycle, clearing existing particles. To avoid particles vanishing from the viewport, wait for the [signal finished] signal before calling.
[b]Note:[/b] The [signal finished] signal is only emitted by [member one_shot] emitters.
*/
func (self Go) Restart() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Restart()
}

/*
Emits a single particle. Whether [param xform], [param velocity], [param color] and [param custom] are applied depends on the value of [param flags]. See [enum EmitFlags].
The default ParticleProcessMaterial will overwrite [param color] and use the contents of [param custom] as [code](rotation, age, animation, lifetime)[/code].
*/
func (self Go) EmitParticle(xform gd.Transform2D, velocity gd.Vector2, color gd.Color, custom gd.Color, flags int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).EmitParticle(xform, velocity, color, custom, gd.Int(flags))
}

/*
Sets this node's properties to match a given [CPUParticles2D] node.
*/
func (self Go) ConvertFromParticles(particles gdclass.Node) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ConvertFromParticles(particles)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GPUParticles2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("GPUParticles2D"))
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

func (self Go) AmountRatio() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAmountRatio()))
}

func (self Go) SetAmountRatio(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAmountRatio(gd.Float(value))
}

func (self Go) SubEmitter() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetSubEmitter(gc).String())
}

func (self Go) SetSubEmitter(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSubEmitter(gc.String(value).NodePath(gc))
}

func (self Go) ProcessMaterial() gdclass.Material {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Material(class(self).GetProcessMaterial(gc))
}

func (self Go) SetProcessMaterial(value gdclass.Material) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetProcessMaterial(value)
}

func (self Go) Texture() gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Texture2D(class(self).GetTexture(gc))
}

func (self Go) SetTexture(value gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTexture(value)
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

func (self Go) FixedFps() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetFixedFps()))
}

func (self Go) SetFixedFps(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFixedFps(gd.Int(value))
}

func (self Go) Interpolate() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetInterpolate())
}

func (self Go) SetInterpolate(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInterpolate(value)
}

func (self Go) FractDelta() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetFractionalDelta())
}

func (self Go) SetFractDelta(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFractionalDelta(value)
}

func (self Go) InterpToEnd() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetInterpToEnd()))
}

func (self Go) SetInterpToEnd(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInterpToEnd(gd.Float(value))
}

func (self Go) CollisionBaseSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetCollisionBaseSize()))
}

func (self Go) SetCollisionBaseSize(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionBaseSize(gd.Float(value))
}

func (self Go) VisibilityRect() gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Rect2(class(self).GetVisibilityRect())
}

func (self Go) SetVisibilityRect(value gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibilityRect(value)
}

func (self Go) LocalCoords() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetUseLocalCoordinates())
}

func (self Go) SetLocalCoords(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseLocalCoordinates(value)
}

func (self Go) DrawOrder() classdb.GPUParticles2DDrawOrder {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.GPUParticles2DDrawOrder(class(self).GetDrawOrder())
}

func (self Go) SetDrawOrder(value classdb.GPUParticles2DDrawOrder) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawOrder(value)
}

func (self Go) TrailEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsTrailEnabled())
}

func (self Go) SetTrailEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTrailEnabled(value)
}

func (self Go) TrailLifetime() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTrailLifetime()))
}

func (self Go) SetTrailLifetime(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTrailLifetime(gd.Float(value))
}

func (self Go) TrailSections() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetTrailSections()))
}

func (self Go) SetTrailSections(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTrailSections(gd.Int(value))
}

func (self Go) TrailSectionSubdivisions() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetTrailSectionSubdivisions()))
}

func (self Go) SetTrailSectionSubdivisions(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTrailSectionSubdivisions(gd.Int(value))
}

//go:nosplit
func (self class) SetEmitting(emitting bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, emitting)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_emitting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetAmount(amount gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetLifetime(secs gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_lifetime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetOneShot(secs bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_one_shot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetPreProcessTime(secs gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_pre_process_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetExplosivenessRatio(ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_explosiveness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetRandomnessRatio(ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_randomness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetVisibilityRect(visibility_rect gd.Rect2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, visibility_rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_visibility_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetUseLocalCoordinates(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_use_local_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFixedFps(fps gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fps)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_fixed_fps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFractionalDelta(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_fractional_delta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetInterpolate(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_interpolate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetProcessMaterial(material gdclass.Material)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_process_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetSpeedScale(scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCollisionBaseSize(size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_collision_base_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetInterpToEnd(interp gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, interp)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_interp_to_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEmitting() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_is_emitting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetAmount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetLifetime() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_lifetime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetOneShot() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_one_shot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetPreProcessTime() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_pre_process_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetExplosivenessRatio() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_explosiveness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetRandomnessRatio() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_randomness_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetVisibilityRect() gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_visibility_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetUseLocalCoordinates() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_use_local_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetFixedFps() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_fixed_fps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetFractionalDelta() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_fractional_delta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetInterpolate() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_interpolate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetProcessMaterial(ctx gd.Lifetime) gdclass.Material {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_process_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Material
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetSpeedScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetCollisionBaseSize() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_collision_base_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetInterpToEnd() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_interp_to_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDrawOrder(order classdb.GPUParticles2DDrawOrder)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, order)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_draw_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDrawOrder() classdb.GPUParticles2DDrawOrder {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GPUParticles2DDrawOrder](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_draw_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTexture(ctx gd.Lifetime) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns a rectangle containing the positions of all existing particles.
[b]Note:[/b] When using threaded rendering this method synchronizes the rendering thread. Calling it often may have a negative impact on performance.
*/
//go:nosplit
func (self class) CaptureRect() gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_capture_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Restarts the particle emission cycle, clearing existing particles. To avoid particles vanishing from the viewport, wait for the [signal finished] signal before calling.
[b]Note:[/b] The [signal finished] signal is only emitted by [member one_shot] emitters.
*/
//go:nosplit
func (self class) Restart()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_restart, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetSubEmitter(path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_sub_emitter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSubEmitter(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_sub_emitter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Emits a single particle. Whether [param xform], [param velocity], [param color] and [param custom] are applied depends on the value of [param flags]. See [enum EmitFlags].
The default ParticleProcessMaterial will overwrite [param color] and use the contents of [param custom] as [code](rotation, age, animation, lifetime)[/code].
*/
//go:nosplit
func (self class) EmitParticle(xform gd.Transform2D, velocity gd.Vector2, color gd.Color, custom gd.Color, flags gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, xform)
	callframe.Arg(frame, velocity)
	callframe.Arg(frame, color)
	callframe.Arg(frame, custom)
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_emit_particle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetTrailEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_trail_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetTrailLifetime(secs gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_trail_lifetime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsTrailEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_is_trail_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetTrailLifetime() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_trail_lifetime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTrailSections(sections gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sections)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_trail_sections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTrailSections() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_trail_sections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTrailSectionSubdivisions(subdivisions gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, subdivisions)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_trail_section_subdivisions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTrailSectionSubdivisions() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_trail_section_subdivisions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets this node's properties to match a given [CPUParticles2D] node.
*/
//go:nosplit
func (self class) ConvertFromParticles(particles gdclass.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(particles[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_convert_from_particles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetAmountRatio(ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_set_amount_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAmountRatio() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticles2D.Bind_get_amount_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnFinished(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("finished"), gc.Callable(cb), 0)
}


func (self class) AsGPUParticles2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGPUParticles2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("GPUParticles2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
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
