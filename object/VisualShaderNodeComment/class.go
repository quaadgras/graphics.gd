package VisualShaderNodeComment

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VisualShaderNodeFrame"
import "grow.graphics/gd/object/VisualShaderNodeResizableBase"
import "grow.graphics/gd/object/VisualShaderNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This node was replaced by [VisualShaderNodeFrame] and only exists to preserve compatibility. In the [VisualShader] editor it behaves exactly like [VisualShaderNodeFrame].

*/
type Simple [1]classdb.VisualShaderNodeComment
func (self Simple) SetDescription(description string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDescription(gc.String(description))
}
func (self Simple) GetDescription() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetDescription(gc).String())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeComment
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetDescription(description gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(description))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeComment.Bind_set_description, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDescription(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeComment.Bind_get_description, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeComment() Expert { return self[0].AsVisualShaderNodeComment() }


//go:nosplit
func (self Simple) AsVisualShaderNodeComment() Simple { return self[0].AsVisualShaderNodeComment() }


//go:nosplit
func (self class) AsVisualShaderNodeFrame() VisualShaderNodeFrame.Expert { return self[0].AsVisualShaderNodeFrame() }


//go:nosplit
func (self Simple) AsVisualShaderNodeFrame() VisualShaderNodeFrame.Simple { return self[0].AsVisualShaderNodeFrame() }


//go:nosplit
func (self class) AsVisualShaderNodeResizableBase() VisualShaderNodeResizableBase.Expert { return self[0].AsVisualShaderNodeResizableBase() }


//go:nosplit
func (self Simple) AsVisualShaderNodeResizableBase() VisualShaderNodeResizableBase.Simple { return self[0].AsVisualShaderNodeResizableBase() }


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
func init() {classdb.Register("VisualShaderNodeComment", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
