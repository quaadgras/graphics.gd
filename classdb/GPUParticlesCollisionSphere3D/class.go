// Package GPUParticlesCollisionSphere3D provides methods for working with GPUParticlesCollisionSphere3D object instances.
package GPUParticlesCollisionSphere3D

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
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A sphere-shaped 3D particle collision shape affecting [GPUParticles3D] nodes.
Particle collision shapes work in real-time and can be moved, rotated and scaled during gameplay. Unlike attractors, non-uniform scaling of collision shapes is [i]not[/i] supported.
[b]Note:[/b] [member ParticleProcessMaterial.collision_mode] must be [constant ParticleProcessMaterial.COLLISION_RIGID] or [constant ParticleProcessMaterial.COLLISION_HIDE_ON_CONTACT] on the [GPUParticles3D]'s process material for collision to work.
[b]Note:[/b] Particle collision only affects [GPUParticles3D], not [CPUParticles3D].
*/
type Instance [1]gdclass.GPUParticlesCollisionSphere3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGPUParticlesCollisionSphere3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GPUParticlesCollisionSphere3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GPUParticlesCollisionSphere3D"))
	casted := Instance{*(*gdclass.GPUParticlesCollisionSphere3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Radius() Float.X {
	return Float.X(Float.X(class(self).GetRadius()))
}

func (self Instance) SetRadius(value Float.X) {
	class(self).SetRadius(gd.Float(value))
}

//go:nosplit
func (self class) SetRadius(radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSphere3D.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSphere3D.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGPUParticlesCollisionSphere3D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGPUParticlesCollisionSphere3D() Instance {
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
	gdclass.Register("GPUParticlesCollisionSphere3D", func(ptr gd.Object) any {
		return [1]gdclass.GPUParticlesCollisionSphere3D{*(*gdclass.GPUParticlesCollisionSphere3D)(unsafe.Pointer(&ptr))}
	})
}
