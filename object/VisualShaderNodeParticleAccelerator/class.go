package VisualShaderNodeParticleAccelerator

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VisualShaderNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Particle accelerator can be used in "process" step of particle shader. It will accelerate the particles. Connect it to the Velocity output port.

*/
type Simple [1]classdb.VisualShaderNodeParticleAccelerator
func (self Simple) SetMode(mode classdb.VisualShaderNodeParticleAcceleratorMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMode(mode)
}
func (self Simple) GetMode() classdb.VisualShaderNodeParticleAcceleratorMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeParticleAcceleratorMode(Expert(self).GetMode())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeParticleAccelerator
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetMode(mode classdb.VisualShaderNodeParticleAcceleratorMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParticleAccelerator.Bind_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMode() classdb.VisualShaderNodeParticleAcceleratorMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeParticleAcceleratorMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParticleAccelerator.Bind_get_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeParticleAccelerator() Expert { return self[0].AsVisualShaderNodeParticleAccelerator() }


//go:nosplit
func (self Simple) AsVisualShaderNodeParticleAccelerator() Simple { return self[0].AsVisualShaderNodeParticleAccelerator() }


//go:nosplit
func (self class) AsVisualShaderNode() VisualShaderNode.Expert { return self[0].AsVisualShaderNode() }


//go:nosplit
func (self Simple) AsVisualShaderNode() VisualShaderNode.Simple { return self[0].AsVisualShaderNode() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("VisualShaderNodeParticleAccelerator", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Mode = classdb.VisualShaderNodeParticleAcceleratorMode

const (
/*The particles will be accelerated based on their velocity.*/
	ModeLinear Mode = 0
/*The particles will be accelerated towards or away from the center.*/
	ModeRadial Mode = 1
/*The particles will be accelerated tangentially to the radius vector from center to their position.*/
	ModeTangential Mode = 2
/*Represents the size of the [enum Mode] enum.*/
	ModeMax Mode = 3
)
