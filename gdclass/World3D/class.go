package World3D

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
Class that has everything pertaining to a world: A physics space, a visual scenario, and a sound space. 3D nodes register their resources into the current 3D world.

*/
type Go [1]classdb.World3D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.World3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("World3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Environment() gdclass.Environment {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Environment(class(self).GetEnvironment(gc))
}

func (self Go) SetEnvironment(value gdclass.Environment) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnvironment(value)
}

func (self Go) FallbackEnvironment() gdclass.Environment {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Environment(class(self).GetFallbackEnvironment(gc))
}

func (self Go) SetFallbackEnvironment(value gdclass.Environment) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFallbackEnvironment(value)
}

func (self Go) CameraAttributes() gdclass.CameraAttributes {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.CameraAttributes(class(self).GetCameraAttributes(gc))
}

func (self Go) SetCameraAttributes(value gdclass.CameraAttributes) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCameraAttributes(value)
}

func (self Go) Space() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
		return gd.RID(class(self).GetSpace())
}

func (self Go) NavigationMap() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
		return gd.RID(class(self).GetNavigationMap())
}

func (self Go) Scenario() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
		return gd.RID(class(self).GetScenario())
}

func (self Go) DirectSpaceState() gdclass.PhysicsDirectSpaceState3D {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.PhysicsDirectSpaceState3D(class(self).GetDirectSpaceState(gc))
}

//go:nosplit
func (self class) GetSpace() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.World3D.Bind_get_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetNavigationMap() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.World3D.Bind_get_navigation_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetScenario() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.World3D.Bind_get_scenario, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnvironment(env gdclass.Environment)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(env[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.World3D.Bind_set_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnvironment(ctx gd.Lifetime) gdclass.Environment {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.World3D.Bind_get_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Environment
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFallbackEnvironment(env gdclass.Environment)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(env[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.World3D.Bind_set_fallback_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFallbackEnvironment(ctx gd.Lifetime) gdclass.Environment {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.World3D.Bind_get_fallback_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Environment
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCameraAttributes(attributes gdclass.CameraAttributes)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(attributes[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.World3D.Bind_set_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCameraAttributes(ctx gd.Lifetime) gdclass.CameraAttributes {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.World3D.Bind_get_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.CameraAttributes
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetDirectSpaceState(ctx gd.Lifetime) gdclass.PhysicsDirectSpaceState3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.World3D.Bind_get_direct_space_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.PhysicsDirectSpaceState3D
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsWorld3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsWorld3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("World3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
