package VisualShaderNodeParticleMultiplyByAxisAngle

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
This node helps to multiply a position input vector by rotation using specific axis. Intended to work with emitters.

*/
type Simple [1]classdb.VisualShaderNodeParticleMultiplyByAxisAngle
func (self Simple) SetDegreesMode(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDegreesMode(enabled)
}
func (self Simple) IsDegreesMode() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDegreesMode())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeParticleMultiplyByAxisAngle
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetDegreesMode(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParticleMultiplyByAxisAngle.Bind_set_degrees_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDegreesMode() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParticleMultiplyByAxisAngle.Bind_is_degrees_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeParticleMultiplyByAxisAngle() Expert { return self[0].AsVisualShaderNodeParticleMultiplyByAxisAngle() }


//go:nosplit
func (self Simple) AsVisualShaderNodeParticleMultiplyByAxisAngle() Simple { return self[0].AsVisualShaderNodeParticleMultiplyByAxisAngle() }


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
func init() {classdb.Register("VisualShaderNodeParticleMultiplyByAxisAngle", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
