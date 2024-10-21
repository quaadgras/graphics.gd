package VisualShaderNodeParticleEmit

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
This node internally calls [code]emit_subparticle[/code] shader method. It will emit a particle from the configured sub-emitter and also allows to customize how its emitted. Requires a sub-emitter assigned to the particles node with this shader.

*/
type Simple [1]classdb.VisualShaderNodeParticleEmit
func (self Simple) SetFlags(flags classdb.VisualShaderNodeParticleEmitEmitFlags) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFlags(flags)
}
func (self Simple) GetFlags() classdb.VisualShaderNodeParticleEmitEmitFlags {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeParticleEmitEmitFlags(Expert(self).GetFlags())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeParticleEmit
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetFlags(flags classdb.VisualShaderNodeParticleEmitEmitFlags)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParticleEmit.Bind_set_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFlags() classdb.VisualShaderNodeParticleEmitEmitFlags {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeParticleEmitEmitFlags](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParticleEmit.Bind_get_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeParticleEmit() Expert { return self[0].AsVisualShaderNodeParticleEmit() }


//go:nosplit
func (self Simple) AsVisualShaderNodeParticleEmit() Simple { return self[0].AsVisualShaderNodeParticleEmit() }


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

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("VisualShaderNodeParticleEmit", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
