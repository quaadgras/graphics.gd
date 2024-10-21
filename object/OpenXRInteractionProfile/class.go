package OpenXRInteractionProfile

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
This object stores suggested bindings for an interaction profile. Interaction profiles define the metadata for a tracked XR device such as an XR controller.
For more information see the [url=https://www.khronos.org/registry/OpenXR/specs/1.0/html/xrspec.html#semantic-path-interaction-profiles]interaction profiles info in the OpenXR specification[/url].

*/
type Simple [1]classdb.OpenXRInteractionProfile
func (self Simple) SetInteractionProfilePath(interaction_profile_path string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInteractionProfilePath(gc.String(interaction_profile_path))
}
func (self Simple) GetInteractionProfilePath() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetInteractionProfilePath(gc).String())
}
func (self Simple) GetBindingCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBindingCount()))
}
func (self Simple) GetBinding(index int) [1]classdb.OpenXRIPBinding {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.OpenXRIPBinding(Expert(self).GetBinding(gc, gd.Int(index)))
}
func (self Simple) SetBindings(bindings gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBindings(bindings)
}
func (self Simple) GetBindings() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).GetBindings(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.OpenXRInteractionProfile
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetInteractionProfilePath(interaction_profile_path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(interaction_profile_path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInteractionProfile.Bind_set_interaction_profile_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInteractionProfilePath(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInteractionProfile.Bind_get_interaction_profile_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Get the number of bindings in this interaction profile.
*/
//go:nosplit
func (self class) GetBindingCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInteractionProfile.Bind_get_binding_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Retrieve the binding at this index.
*/
//go:nosplit
func (self class) GetBinding(ctx gd.Lifetime, index gd.Int) object.OpenXRIPBinding {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInteractionProfile.Bind_get_binding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.OpenXRIPBinding
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBindings(bindings gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bindings))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInteractionProfile.Bind_set_bindings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBindings(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInteractionProfile.Bind_get_bindings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsOpenXRInteractionProfile() Expert { return self[0].AsOpenXRInteractionProfile() }


//go:nosplit
func (self Simple) AsOpenXRInteractionProfile() Simple { return self[0].AsOpenXRInteractionProfile() }


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
func init() {classdb.Register("OpenXRInteractionProfile", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
