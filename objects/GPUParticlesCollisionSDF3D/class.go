package GPUParticlesCollisionSDF3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/GPUParticlesCollision3D"
import "graphics.gd/objects/VisualInstance3D"
import "graphics.gd/objects/Node3D"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A baked signed distance field 3D particle collision shape affecting [GPUParticles3D] nodes.
Signed distance fields (SDF) allow for efficiently representing approximate collision shapes for convex and concave objects of any shape. This is more flexible than [GPUParticlesCollisionHeightField3D], but it requires a baking step.
[b]Baking:[/b] The signed distance field texture can be baked by selecting the [GPUParticlesCollisionSDF3D] node in the editor, then clicking [b]Bake SDF[/b] at the top of the 3D viewport. Any [i]visible[/i] [MeshInstance3D]s within the [member size] will be taken into account for baking, regardless of their [member GeometryInstance3D.gi_mode].
[b]Note:[/b] Baking a [GPUParticlesCollisionSDF3D]'s [member texture] is only possible within the editor, as there is no bake method exposed for use in exported projects. However, it's still possible to load pre-baked [Texture3D]s into its [member texture] property in an exported project.
[b]Note:[/b] [member ParticleProcessMaterial.collision_mode] must be [constant ParticleProcessMaterial.COLLISION_RIGID] or [constant ParticleProcessMaterial.COLLISION_HIDE_ON_CONTACT] on the [GPUParticles3D]'s process material for collision to work.
[b]Note:[/b] Particle collision only affects [GPUParticles3D], not [CPUParticles3D].
*/
type Instance [1]classdb.GPUParticlesCollisionSDF3D
type Any interface {
	gd.IsClass
	AsGPUParticlesCollisionSDF3D() Instance
}

/*
Based on [param value], enables or disables the specified layer in the [member bake_mask], given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetBakeMaskValue(layer_number int, value bool) {
	class(self).SetBakeMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member bake_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetBakeMaskValue(layer_number int) bool {
	return bool(class(self).GetBakeMaskValue(gd.Int(layer_number)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.GPUParticlesCollisionSDF3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GPUParticlesCollisionSDF3D"))
	return Instance{*(*classdb.GPUParticlesCollisionSDF3D)(unsafe.Pointer(&object))}
}

func (self Instance) Size() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetSize())
}

func (self Instance) SetSize(value Vector3.XYZ) {
	class(self).SetSize(gd.Vector3(value))
}

func (self Instance) Resolution() classdb.GPUParticlesCollisionSDF3DResolution {
	return classdb.GPUParticlesCollisionSDF3DResolution(class(self).GetResolution())
}

func (self Instance) SetResolution(value classdb.GPUParticlesCollisionSDF3DResolution) {
	class(self).SetResolution(value)
}

func (self Instance) Thickness() Float.X {
	return Float.X(Float.X(class(self).GetThickness()))
}

func (self Instance) SetThickness(value Float.X) {
	class(self).SetThickness(gd.Float(value))
}

func (self Instance) BakeMask() int {
	return int(int(class(self).GetBakeMask()))
}

func (self Instance) SetBakeMask(value int) {
	class(self).SetBakeMask(gd.Int(value))
}

func (self Instance) Texture() objects.Texture3D {
	return objects.Texture3D(class(self).GetTexture())
}

func (self Instance) SetTexture(value objects.Texture3D) {
	class(self).SetTexture(value)
}

//go:nosplit
func (self class) SetSize(size gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSize() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetResolution(resolution classdb.GPUParticlesCollisionSDF3DResolution) {
	var frame = callframe.New()
	callframe.Arg(frame, resolution)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_set_resolution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetResolution() classdb.GPUParticlesCollisionSDF3DResolution {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GPUParticlesCollisionSDF3DResolution](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_get_resolution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTexture(texture objects.Texture3D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() objects.Texture3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture3D{gd.PointerWithOwnershipTransferredToGo[classdb.Texture3D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetThickness(thickness gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, thickness)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_set_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetThickness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_get_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBakeMask(mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_set_bake_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBakeMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_get_bake_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member bake_mask], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetBakeMaskValue(layer_number gd.Int, value bool) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_set_bake_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member bake_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetBakeMaskValue(layer_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_get_bake_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGPUParticlesCollisionSDF3D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGPUParticlesCollisionSDF3D() Instance {
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
		return gd.VirtualByName(self.AsGPUParticlesCollision3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsGPUParticlesCollision3D(), name)
	}
}
func init() {
	classdb.Register("GPUParticlesCollisionSDF3D", func(ptr gd.Object) any {
		return [1]classdb.GPUParticlesCollisionSDF3D{*(*classdb.GPUParticlesCollisionSDF3D)(unsafe.Pointer(&ptr))}
	})
}

type Resolution = classdb.GPUParticlesCollisionSDF3DResolution

const (
	/*Bake a 16×16×16 signed distance field. This is the fastest option, but also the least precise.*/
	Resolution16 Resolution = 0
	/*Bake a 32×32×32 signed distance field.*/
	Resolution32 Resolution = 1
	/*Bake a 64×64×64 signed distance field.*/
	Resolution64 Resolution = 2
	/*Bake a 128×128×128 signed distance field.*/
	Resolution128 Resolution = 3
	/*Bake a 256×256×256 signed distance field.*/
	Resolution256 Resolution = 4
	/*Bake a 512×512×512 signed distance field. This is the slowest option, but also the most precise.*/
	Resolution512 Resolution = 5
	/*Represents the size of the [enum Resolution] enum.*/
	ResolutionMax Resolution = 6
)
