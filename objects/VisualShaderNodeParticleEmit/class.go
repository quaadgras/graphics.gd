package VisualShaderNodeParticleEmit

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
This node internally calls [code]emit_subparticle[/code] shader method. It will emit a particle from the configured sub-emitter and also allows to customize how its emitted. Requires a sub-emitter assigned to the particles node with this shader.
*/
type Instance [1]classdb.VisualShaderNodeParticleEmit
type Any interface {
	gd.IsClass
	AsVisualShaderNodeParticleEmit() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShaderNodeParticleEmit

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeParticleEmit"))
	return Instance{classdb.VisualShaderNodeParticleEmit(object)}
}

func (self Instance) Flags() classdb.VisualShaderNodeParticleEmitEmitFlags {
	return classdb.VisualShaderNodeParticleEmitEmitFlags(class(self).GetFlags())
}

func (self Instance) SetFlags(value classdb.VisualShaderNodeParticleEmitEmitFlags) {
	class(self).SetFlags(value)
}

//go:nosplit
func (self class) SetFlags(flags classdb.VisualShaderNodeParticleEmitEmitFlags) {
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleEmit.Bind_set_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFlags() classdb.VisualShaderNodeParticleEmitEmitFlags {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeParticleEmitEmitFlags](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleEmit.Bind_get_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	classdb.Register("VisualShaderNodeParticleEmit", func(ptr gd.Object) any {
		return [1]classdb.VisualShaderNodeParticleEmit{classdb.VisualShaderNodeParticleEmit(ptr)}
	})
}

type EmitFlags = classdb.VisualShaderNodeParticleEmitEmitFlags

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
