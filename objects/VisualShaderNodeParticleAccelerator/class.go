package VisualShaderNodeParticleAccelerator

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/VisualShaderNode"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Particle accelerator can be used in "process" step of particle shader. It will accelerate the particles. Connect it to the Velocity output port.
*/
type Instance [1]classdb.VisualShaderNodeParticleAccelerator
type Any interface {
	gd.IsClass
	AsVisualShaderNodeParticleAccelerator() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShaderNodeParticleAccelerator

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeParticleAccelerator"))
	return Instance{classdb.VisualShaderNodeParticleAccelerator(object)}
}

func (self Instance) Mode() classdb.VisualShaderNodeParticleAcceleratorMode {
	return classdb.VisualShaderNodeParticleAcceleratorMode(class(self).GetMode())
}

func (self Instance) SetMode(value classdb.VisualShaderNodeParticleAcceleratorMode) {
	class(self).SetMode(value)
}

//go:nosplit
func (self class) SetMode(mode classdb.VisualShaderNodeParticleAcceleratorMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleAccelerator.Bind_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMode() classdb.VisualShaderNodeParticleAcceleratorMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeParticleAcceleratorMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleAccelerator.Bind_get_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeParticleAccelerator() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeParticleAccelerator() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNode() VisualShaderNode.Advanced {
	return *((*VisualShaderNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNode() VisualShaderNode.Instance {
	return *((*VisualShaderNode.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}
func init() {
	classdb.Register("VisualShaderNodeParticleAccelerator", func(ptr gd.Object) any { return classdb.VisualShaderNodeParticleAccelerator(ptr) })
}

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
