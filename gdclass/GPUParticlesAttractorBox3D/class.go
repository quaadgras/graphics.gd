package GPUParticlesAttractorBox3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/GPUParticlesAttractor3D"
import "grow.graphics/gd/gdclass/VisualInstance3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A box-shaped attractor that influences particles from [GPUParticles3D] nodes. Can be used to attract particles towards its origin, or to push them away from its origin.
Particle attractors work in real-time and can be moved, rotated and scaled during gameplay. Unlike collision shapes, non-uniform scaling of attractors is also supported.
[b]Note:[/b] Particle attractors only affect [GPUParticles3D], not [CPUParticles3D].

*/
type Go [1]classdb.GPUParticlesAttractorBox3D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GPUParticlesAttractorBox3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("GPUParticlesAttractorBox3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Size() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetSize())
}

func (self Go) SetSize(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSize(value)
}

//go:nosplit
func (self class) SetSize(size gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticlesAttractorBox3D.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSize() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GPUParticlesAttractorBox3D.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGPUParticlesAttractorBox3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGPUParticlesAttractorBox3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsGPUParticlesAttractor3D() GPUParticlesAttractor3D.GD { return *((*GPUParticlesAttractor3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsGPUParticlesAttractor3D() GPUParticlesAttractor3D.Go { return *((*GPUParticlesAttractor3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualInstance3D() VisualInstance3D.GD { return *((*VisualInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualInstance3D() VisualInstance3D.Go { return *((*VisualInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsGPUParticlesAttractor3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsGPUParticlesAttractor3D(), name)
	}
}
func init() {classdb.Register("GPUParticlesAttractorBox3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
