// Package GPUParticlesCollisionHeightField3D provides methods for working with GPUParticlesCollisionHeightField3D object instances.
package GPUParticlesCollisionHeightField3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/GPUParticlesCollision3D"
import "graphics.gd/classdb/VisualInstance3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector3"

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
A real-time heightmap-shaped 3D particle collision shape affecting [GPUParticles3D] nodes.
Heightmap shapes allow for efficiently representing collisions for convex and concave objects with a single "floor" (such as terrain). This is less flexible than [GPUParticlesCollisionSDF3D], but it doesn't require a baking step.
[GPUParticlesCollisionHeightField3D] can also be regenerated in real-time when it is moved, when the camera moves, or even continuously. This makes [GPUParticlesCollisionHeightField3D] a good choice for weather effects such as rain and snow and games with highly dynamic geometry. However, this class is limited since heightmaps cannot represent overhangs (e.g. indoors or caves).
[b]Note:[/b] [member ParticleProcessMaterial.collision_mode] must be [code]true[/code] on the [GPUParticles3D]'s process material for collision to work.
[b]Note:[/b] Particle collision only affects [GPUParticles3D], not [CPUParticles3D].
*/
type Instance [1]gdclass.GPUParticlesCollisionHeightField3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGPUParticlesCollisionHeightField3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GPUParticlesCollisionHeightField3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GPUParticlesCollisionHeightField3D"))
	casted := Instance{*(*gdclass.GPUParticlesCollisionHeightField3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Size() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetSize())
}

func (self Instance) SetSize(value Vector3.XYZ) {
	class(self).SetSize(gd.Vector3(value))
}

func (self Instance) Resolution() gdclass.GPUParticlesCollisionHeightField3DResolution {
	return gdclass.GPUParticlesCollisionHeightField3DResolution(class(self).GetResolution())
}

func (self Instance) SetResolution(value gdclass.GPUParticlesCollisionHeightField3DResolution) {
	class(self).SetResolution(value)
}

func (self Instance) UpdateMode() gdclass.GPUParticlesCollisionHeightField3DUpdateMode {
	return gdclass.GPUParticlesCollisionHeightField3DUpdateMode(class(self).GetUpdateMode())
}

func (self Instance) SetUpdateMode(value gdclass.GPUParticlesCollisionHeightField3DUpdateMode) {
	class(self).SetUpdateMode(value)
}

func (self Instance) FollowCameraEnabled() bool {
	return bool(class(self).IsFollowCameraEnabled())
}

func (self Instance) SetFollowCameraEnabled(value bool) {
	class(self).SetFollowCameraEnabled(value)
}

//go:nosplit
func (self class) SetSize(size gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionHeightField3D.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSize() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionHeightField3D.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetResolution(resolution gdclass.GPUParticlesCollisionHeightField3DResolution) {
	var frame = callframe.New()
	callframe.Arg(frame, resolution)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionHeightField3D.Bind_set_resolution, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetResolution() gdclass.GPUParticlesCollisionHeightField3DResolution {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GPUParticlesCollisionHeightField3DResolution](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionHeightField3D.Bind_get_resolution, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUpdateMode(update_mode gdclass.GPUParticlesCollisionHeightField3DUpdateMode) {
	var frame = callframe.New()
	callframe.Arg(frame, update_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionHeightField3D.Bind_set_update_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUpdateMode() gdclass.GPUParticlesCollisionHeightField3DUpdateMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GPUParticlesCollisionHeightField3DUpdateMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionHeightField3D.Bind_get_update_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFollowCameraEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionHeightField3D.Bind_set_follow_camera_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFollowCameraEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionHeightField3D.Bind_is_follow_camera_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGPUParticlesCollisionHeightField3D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGPUParticlesCollisionHeightField3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsGPUParticlesCollision3D() GPUParticlesCollision3D.Advanced {
	return *((*GPUParticlesCollision3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGPUParticlesCollision3D() GPUParticlesCollision3D.Instance {
	return *((*GPUParticlesCollision3D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(GPUParticlesCollision3D.Advanced(self.AsGPUParticlesCollision3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(GPUParticlesCollision3D.Instance(self.AsGPUParticlesCollision3D()), name)
	}
}
func init() {
	gdclass.Register("GPUParticlesCollisionHeightField3D", func(ptr gd.Object) any {
		return [1]gdclass.GPUParticlesCollisionHeightField3D{*(*gdclass.GPUParticlesCollisionHeightField3D)(unsafe.Pointer(&ptr))}
	})
}

type Resolution = gdclass.GPUParticlesCollisionHeightField3DResolution

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

type UpdateMode = gdclass.GPUParticlesCollisionHeightField3DUpdateMode

const (
	/*Only update the heightmap when the [GPUParticlesCollisionHeightField3D] node is moved, or when the camera moves if [member follow_camera_enabled] is [code]true[/code]. An update can be forced by slightly moving the [GPUParticlesCollisionHeightField3D] in any direction, or by calling [method RenderingServer.particles_collision_height_field_update].*/
	UpdateModeWhenMoved UpdateMode = 0
	/*Update the heightmap every frame. This has a significant performance cost. This update should only be used when geometry that particles can collide with changes significantly during gameplay.*/
	UpdateModeAlways UpdateMode = 1
)
