package Compositor

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The compositor resource stores attributes used to customize how a [Viewport] is rendered.

*/
type Simple [1]classdb.Compositor
func (self Simple) SetCompositorEffects(compositor_effects gd.ArrayOf[[1]classdb.CompositorEffect]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCompositorEffects(compositor_effects)
}
func (self Simple) GetCompositorEffects() gd.ArrayOf[[1]classdb.CompositorEffect] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.CompositorEffect](Expert(self).GetCompositorEffects(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Compositor
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetCompositorEffects(compositor_effects gd.ArrayOf[object.CompositorEffect])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(compositor_effects))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Compositor.Bind_set_compositor_effects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCompositorEffects(ctx gd.Lifetime) gd.ArrayOf[object.CompositorEffect] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Compositor.Bind_get_compositor_effects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.CompositorEffect](ret)
}

//go:nosplit
func (self class) AsCompositor() Expert { return self[0].AsCompositor() }


//go:nosplit
func (self Simple) AsCompositor() Simple { return self[0].AsCompositor() }


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
func init() {classdb.Register("Compositor", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
