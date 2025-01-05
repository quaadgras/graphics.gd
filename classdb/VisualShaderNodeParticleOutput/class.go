// Package VisualShaderNodeParticleOutput provides methods for working with VisualShaderNodeParticleOutput object instances.
package VisualShaderNodeParticleOutput

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/VisualShaderNodeOutput"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This node defines how particles are emitted. It allows to customize e.g. position and velocity. Available ports are different depending on which function this node is inside (start, process, collision) and whether custom data is enabled.
*/
type Instance [1]gdclass.VisualShaderNodeParticleOutput
type Any interface {
	gd.IsClass
	AsVisualShaderNodeParticleOutput() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeParticleOutput

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeParticleOutput"))
	return Instance{*(*gdclass.VisualShaderNodeParticleOutput)(unsafe.Pointer(&object))}
}

func (self class) AsVisualShaderNodeParticleOutput() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeParticleOutput() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNodeOutput() VisualShaderNodeOutput.Advanced {
	return *((*VisualShaderNodeOutput.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeOutput() VisualShaderNodeOutput.Instance {
	return *((*VisualShaderNodeOutput.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(VisualShaderNodeOutput.Advanced(self.AsVisualShaderNodeOutput()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNodeOutput.Instance(self.AsVisualShaderNodeOutput()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeParticleOutput", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeParticleOutput{*(*gdclass.VisualShaderNodeParticleOutput)(unsafe.Pointer(&ptr))}
	})
}
