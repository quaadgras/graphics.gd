package GPUParticlesAttractorBox3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/GPUParticlesAttractor3D"
import "graphics.gd/objects/VisualInstance3D"
import "graphics.gd/objects/Node3D"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Vector3"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A box-shaped attractor that influences particles from [GPUParticles3D] nodes. Can be used to attract particles towards its origin, or to push them away from its origin.
Particle attractors work in real-time and can be moved, rotated and scaled during gameplay. Unlike collision shapes, non-uniform scaling of attractors is also supported.
[b]Note:[/b] Particle attractors only affect [GPUParticles3D], not [CPUParticles3D].
*/
type Instance [1]classdb.GPUParticlesAttractorBox3D
type Any interface {
	gd.IsClass
	AsGPUParticlesAttractorBox3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.GPUParticlesAttractorBox3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GPUParticlesAttractorBox3D"))
	return Instance{*(*classdb.GPUParticlesAttractorBox3D)(unsafe.Pointer(&object))}
}

func (self Instance) Size() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetSize())
}

func (self Instance) SetSize(value Vector3.XYZ) {
	class(self).SetSize(gd.Vector3(value))
}

//go:nosplit
func (self class) SetSize(size gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesAttractorBox3D.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSize() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesAttractorBox3D.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGPUParticlesAttractorBox3D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGPUParticlesAttractorBox3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsGPUParticlesAttractor3D() GPUParticlesAttractor3D.Advanced {
	return *((*GPUParticlesAttractor3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGPUParticlesAttractor3D() GPUParticlesAttractor3D.Instance {
	return *((*GPUParticlesAttractor3D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsGPUParticlesAttractor3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsGPUParticlesAttractor3D(), name)
	}
}
func init() {
	classdb.Register("GPUParticlesAttractorBox3D", func(ptr gd.Object) any {
		return [1]classdb.GPUParticlesAttractorBox3D{*(*classdb.GPUParticlesAttractorBox3D)(unsafe.Pointer(&ptr))}
	})
}
