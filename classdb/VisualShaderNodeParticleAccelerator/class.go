// Package VisualShaderNodeParticleAccelerator provides methods for working with VisualShaderNodeParticleAccelerator object instances.
package VisualShaderNodeParticleAccelerator

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any

/*
Particle accelerator can be used in "process" step of particle shader. It will accelerate the particles. Connect it to the Velocity output port.
*/
type Instance [1]gdclass.VisualShaderNodeParticleAccelerator

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeParticleAccelerator() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeParticleAccelerator

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeParticleAccelerator"))
	casted := Instance{*(*gdclass.VisualShaderNodeParticleAccelerator)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Mode() gdclass.VisualShaderNodeParticleAcceleratorMode {
	return gdclass.VisualShaderNodeParticleAcceleratorMode(class(self).GetMode())
}

func (self Instance) SetMode(value gdclass.VisualShaderNodeParticleAcceleratorMode) {
	class(self).SetMode(value)
}

//go:nosplit
func (self class) SetMode(mode gdclass.VisualShaderNodeParticleAcceleratorMode) { //gd:VisualShaderNodeParticleAccelerator.set_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleAccelerator.Bind_set_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMode() gdclass.VisualShaderNodeParticleAcceleratorMode { //gd:VisualShaderNodeParticleAccelerator.get_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeParticleAcceleratorMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleAccelerator.Bind_get_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNode.Advanced(self.AsVisualShaderNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNode.Instance(self.AsVisualShaderNode()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeParticleAccelerator", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeParticleAccelerator{*(*gdclass.VisualShaderNodeParticleAccelerator)(unsafe.Pointer(&ptr))}
	})
}

type Mode = gdclass.VisualShaderNodeParticleAcceleratorMode //gd:VisualShaderNodeParticleAccelerator.Mode

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
