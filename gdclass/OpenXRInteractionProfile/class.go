package OpenXRInteractionProfile

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This object stores suggested bindings for an interaction profile. Interaction profiles define the metadata for a tracked XR device such as an XR controller.
For more information see the [url=https://www.khronos.org/registry/OpenXR/specs/1.0/html/xrspec.html#semantic-path-interaction-profiles]interaction profiles info in the OpenXR specification[/url].

*/
type Go [1]classdb.OpenXRInteractionProfile

/*
Get the number of bindings in this interaction profile.
*/
func (self Go) GetBindingCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetBindingCount()))
}

/*
Retrieve the binding at this index.
*/
func (self Go) GetBinding(index int) gdclass.OpenXRIPBinding {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.OpenXRIPBinding(class(self).GetBinding(gc, gd.Int(index)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.OpenXRInteractionProfile
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("OpenXRInteractionProfile"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) InteractionProfilePath() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetInteractionProfilePath(gc).String())
}

func (self Go) SetInteractionProfilePath(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInteractionProfilePath(gc.String(value))
}

func (self Go) Bindings() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Array(class(self).GetBindings(gc))
}

func (self Go) SetBindings(value gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBindings(value)
}

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
func (self class) GetBinding(ctx gd.Lifetime, index gd.Int) gdclass.OpenXRIPBinding {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInteractionProfile.Bind_get_binding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.OpenXRIPBinding
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
func (self class) AsOpenXRInteractionProfile() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsOpenXRInteractionProfile() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("OpenXRInteractionProfile", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
