package GPUParticlesCollisionHeightField3D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/GPUParticlesCollision3D"
import "grow.graphics/gd/object/VisualInstance3D"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A real-time heightmap-shaped 3D particle collision shape affecting [GPUParticles3D] nodes.
Heightmap shapes allow for efficiently representing collisions for convex and concave objects with a single "floor" (such as terrain). This is less flexible than [GPUParticlesCollisionSDF3D], but it doesn't require a baking step.
[GPUParticlesCollisionHeightField3D] can also be regenerated in real-time when it is moved, when the camera moves, or even continuously. This makes [GPUParticlesCollisionHeightField3D] a good choice for weather effects such as rain and snow and games with highly dynamic geometry. However, this class is limited since heightmaps cannot represent overhangs (e.g. indoors or caves).
[b]Note:[/b] [member ParticleProcessMaterial.collision_mode] must be [code]true[/code] on the [GPUParticles3D]'s process material for collision to work.
[b]Note:[/b] Particle collision only affects [GPUParticles3D], not [CPUParticles3D].

*/
type Simple [1]classdb.GPUParticlesCollisionHeightField3D
func (self Simple) SetSize(size gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSize(size)
}
func (self Simple) GetSize() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetSize())
}
func (self Simple) SetResolution(resolution classdb.GPUParticlesCollisionHeightField3DResolution) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetResolution(resolution)
}
func (self Simple) GetResolution() classdb.GPUParticlesCollisionHeightField3DResolution {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.GPUParticlesCollisionHeightField3DResolution(Expert(self).GetResolution())
}
func (self Simple) SetUpdateMode(update_mode classdb.GPUParticlesCollisionHeightField3DUpdateMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUpdateMode(update_mode)
}
func (self Simple) GetUpdateMode() classdb.GPUParticlesCollisionHeightField3DUpdateMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.GPUParticlesCollisionHeightField3DUpdateMode(Expert(self).GetUpdateMode())
}
func (self Simple) SetFollowCameraEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFollowCameraEnabled(enabled)
}
func (self Simple) IsFollowCameraEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsFollowCameraEnabled())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.GPUParticlesCollisionHeightField3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetSize(size gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticlesCollisionHeightField3D.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSize() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticlesCollisionHeightField3D.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetResolution(resolution classdb.GPUParticlesCollisionHeightField3DResolution)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, resolution)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticlesCollisionHeightField3D.Bind_set_resolution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetResolution() classdb.GPUParticlesCollisionHeightField3DResolution {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GPUParticlesCollisionHeightField3DResolution](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticlesCollisionHeightField3D.Bind_get_resolution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUpdateMode(update_mode classdb.GPUParticlesCollisionHeightField3DUpdateMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, update_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticlesCollisionHeightField3D.Bind_set_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUpdateMode() classdb.GPUParticlesCollisionHeightField3DUpdateMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GPUParticlesCollisionHeightField3DUpdateMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticlesCollisionHeightField3D.Bind_get_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFollowCameraEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticlesCollisionHeightField3D.Bind_set_follow_camera_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFollowCameraEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticlesCollisionHeightField3D.Bind_is_follow_camera_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsGPUParticlesCollisionHeightField3D() Expert { return self[0].AsGPUParticlesCollisionHeightField3D() }


//go:nosplit
func (self Simple) AsGPUParticlesCollisionHeightField3D() Simple { return self[0].AsGPUParticlesCollisionHeightField3D() }


//go:nosplit
func (self class) AsGPUParticlesCollision3D() GPUParticlesCollision3D.Expert { return self[0].AsGPUParticlesCollision3D() }


//go:nosplit
func (self Simple) AsGPUParticlesCollision3D() GPUParticlesCollision3D.Simple { return self[0].AsGPUParticlesCollision3D() }


//go:nosplit
func (self class) AsVisualInstance3D() VisualInstance3D.Expert { return self[0].AsVisualInstance3D() }


//go:nosplit
func (self Simple) AsVisualInstance3D() VisualInstance3D.Simple { return self[0].AsVisualInstance3D() }


//go:nosplit
func (self class) AsNode3D() Node3D.Expert { return self[0].AsNode3D() }


//go:nosplit
func (self Simple) AsNode3D() Node3D.Simple { return self[0].AsNode3D() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("GPUParticlesCollisionHeightField3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Resolution = classdb.GPUParticlesCollisionHeightField3DResolution

const (
/*Generate a 256×256 heightmap. Intended for small-scale scenes, or larger scenes with no distant particles.*/
	Resolution256 Resolution = 0
/*Generate a 512×512 heightmap. Intended for medium-scale scenes, or larger scenes with no distant particles.*/
	Resolution512 Resolution = 1
/*Generate a 1024×1024 heightmap. Intended for large scenes with distant particles.*/
	Resolution1024 Resolution = 2
/*Generate a 2048×2048 heightmap. Intended for very large scenes with distant particles.*/
	Resolution2048 Resolution = 3
/*Generate a 4096×4096 heightmap. Intended for huge scenes with distant particles.*/
	Resolution4096 Resolution = 4
/*Generate a 8192×8192 heightmap. Intended for gigantic scenes with distant particles.*/
	Resolution8192 Resolution = 5
/*Represents the size of the [enum Resolution] enum.*/
	ResolutionMax Resolution = 6
)
type UpdateMode = classdb.GPUParticlesCollisionHeightField3DUpdateMode

const (
/*Only update the heightmap when the [GPUParticlesCollisionHeightField3D] node is moved, or when the camera moves if [member follow_camera_enabled] is [code]true[/code]. An update can be forced by slightly moving the [GPUParticlesCollisionHeightField3D] in any direction, or by calling [method RenderingServer.particles_collision_height_field_update].*/
	UpdateModeWhenMoved UpdateMode = 0
/*Update the heightmap every frame. This has a significant performance cost. This update should only be used when geometry that particles can collide with changes significantly during gameplay.*/
	UpdateModeAlways UpdateMode = 1
)
