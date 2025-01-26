// Package VisualShaderNodeParticleEmit provides methods for working with VisualShaderNodeParticleEmit object instances.
package VisualShaderNodeParticleEmit

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

/*
This node internally calls [code]emit_subparticle[/code] shader method. It will emit a particle from the configured sub-emitter and also allows to customize how its emitted. Requires a sub-emitter assigned to the particles node with this shader.
*/
type Instance [1]gdclass.VisualShaderNodeParticleEmit

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeParticleEmit() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeParticleEmit

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeParticleEmit"))
	casted := Instance{*(*gdclass.VisualShaderNodeParticleEmit)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Flags() gdclass.VisualShaderNodeParticleEmitEmitFlags {
	return gdclass.VisualShaderNodeParticleEmitEmitFlags(class(self).GetFlags())
}

func (self Instance) SetFlags(value gdclass.VisualShaderNodeParticleEmitEmitFlags) {
	class(self).SetFlags(value)
}

//go:nosplit
func (self class) SetFlags(flags gdclass.VisualShaderNodeParticleEmitEmitFlags) { //gd:VisualShaderNodeParticleEmit.set_flags
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleEmit.Bind_set_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFlags() gdclass.VisualShaderNodeParticleEmitEmitFlags { //gd:VisualShaderNodeParticleEmit.get_flags
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeParticleEmitEmitFlags](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleEmit.Bind_get_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeParticleEmit() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeParticleEmit() Instance {
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
	gdclass.Register("VisualShaderNodeParticleEmit", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeParticleEmit{*(*gdclass.VisualShaderNodeParticleEmit)(unsafe.Pointer(&ptr))}
	})
}

type EmitFlags = gdclass.VisualShaderNodeParticleEmitEmitFlags //gd:VisualShaderNodeParticleEmit.EmitFlags

const (
	/*If enabled, the particle starts with the position defined by this node.*/
	EmitFlagPosition EmitFlags = 1
	/*If enabled, the particle starts with the rotation and scale defined by this node.*/
	EmitFlagRotScale EmitFlags = 2
	/*If enabled,the particle starts with the velocity defined by this node.*/
	EmitFlagVelocity EmitFlags = 4
	/*If enabled, the particle starts with the color defined by this node.*/
	EmitFlagColor EmitFlags = 8
	/*If enabled, the particle starts with the [code]CUSTOM[/code] data defined by this node.*/
	EmitFlagCustom EmitFlags = 16
)
