package GPUParticlesCollisionHeightField3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/GPUParticlesCollision3D"
import "grow.graphics/gd/gdclass/VisualInstance3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
A real-time heightmap-shaped 3D particle collision shape affecting [GPUParticles3D] nodes.
Heightmap shapes allow for efficiently representing collisions for convex and concave objects with a single "floor" (such as terrain). This is less flexible than [GPUParticlesCollisionSDF3D], but it doesn't require a baking step.
[GPUParticlesCollisionHeightField3D] can also be regenerated in real-time when it is moved, when the camera moves, or even continuously. This makes [GPUParticlesCollisionHeightField3D] a good choice for weather effects such as rain and snow and games with highly dynamic geometry. However, this class is limited since heightmaps cannot represent overhangs (e.g. indoors or caves).
[b]Note:[/b] [member ParticleProcessMaterial.collision_mode] must be [code]true[/code] on the [GPUParticles3D]'s process material for collision to work.
[b]Note:[/b] Particle collision only affects [GPUParticles3D], not [CPUParticles3D].

*/
type Go [1]classdb.GPUParticlesCollisionHeightField3D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GPUParticlesCollisionHeightField3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GPUParticlesCollisionHeightField3D"))
	return Go{classdb.GPUParticlesCollisionHeightField3D(object)}
}

func (self Go) Size() gd.Vector3 {
		return gd.Vector3(class(self).GetSize())
}

func (self Go) SetSize(value gd.Vector3) {
	class(self).SetSize(value)
}

func (self Go) Resolution() classdb.GPUParticlesCollisionHeightField3DResolution {
		return classdb.GPUParticlesCollisionHeightField3DResolution(class(self).GetResolution())
}

func (self Go) SetResolution(value classdb.GPUParticlesCollisionHeightField3DResolution) {
	class(self).SetResolution(value)
}

func (self Go) UpdateMode() classdb.GPUParticlesCollisionHeightField3DUpdateMode {
		return classdb.GPUParticlesCollisionHeightField3DUpdateMode(class(self).GetUpdateMode())
}

func (self Go) SetUpdateMode(value classdb.GPUParticlesCollisionHeightField3DUpdateMode) {
	class(self).SetUpdateMode(value)
}

func (self Go) FollowCameraEnabled() bool {
		return bool(class(self).IsFollowCameraEnabled())
}

func (self Go) SetFollowCameraEnabled(value bool) {
	class(self).SetFollowCameraEnabled(value)
}

//go:nosplit
func (self class) SetSize(size gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionHeightField3D.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSize() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionHeightField3D.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetResolution(resolution classdb.GPUParticlesCollisionHeightField3DResolution)  {
	var frame = callframe.New()
	callframe.Arg(frame, resolution)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionHeightField3D.Bind_set_resolution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetResolution() classdb.GPUParticlesCollisionHeightField3DResolution {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GPUParticlesCollisionHeightField3DResolution](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionHeightField3D.Bind_get_resolution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUpdateMode(update_mode classdb.GPUParticlesCollisionHeightField3DUpdateMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, update_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionHeightField3D.Bind_set_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUpdateMode() classdb.GPUParticlesCollisionHeightField3DUpdateMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GPUParticlesCollisionHeightField3DUpdateMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionHeightField3D.Bind_get_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFollowCameraEnabled(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionHeightField3D.Bind_set_follow_camera_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFollowCameraEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionHeightField3D.Bind_is_follow_camera_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGPUParticlesCollisionHeightField3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGPUParticlesCollisionHeightField3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsGPUParticlesCollision3D() GPUParticlesCollision3D.GD { return *((*GPUParticlesCollision3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsGPUParticlesCollision3D() GPUParticlesCollision3D.Go { return *((*GPUParticlesCollision3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualInstance3D() VisualInstance3D.GD { return *((*VisualInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualInstance3D() VisualInstance3D.Go { return *((*VisualInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsGPUParticlesCollision3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsGPUParticlesCollision3D(), name)
	}
}
func init() {classdb.Register("GPUParticlesCollisionHeightField3D", func(ptr gd.Object) any { return classdb.GPUParticlesCollisionHeightField3D(ptr) })}
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
