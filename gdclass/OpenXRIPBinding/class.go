package OpenXRIPBinding

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
This binding resource binds an [OpenXRAction] to inputs or outputs. As most controllers have left hand and right versions that are handled by the same interaction profile we can specify multiple bindings. For instance an action "Fire" could be bound to both "/user/hand/left/input/trigger" and "/user/hand/right/input/trigger".

*/
type Go [1]classdb.OpenXRIPBinding

/*
Get the number of input/output paths in this binding.
*/
func (self Go) GetPathCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetPathCount()))
}

/*
Returns [code]true[/code] if this input/output path is part of this binding.
*/
func (self Go) HasPath(path string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasPath(gc.String(path)))
}

/*
Add an input/output path to this binding.
*/
func (self Go) AddPath(path string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddPath(gc.String(path))
}

/*
Removes this input/output path from this binding.
*/
func (self Go) RemovePath(path string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemovePath(gc.String(path))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.OpenXRIPBinding
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("OpenXRIPBinding"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Action() gdclass.OpenXRAction {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.OpenXRAction(class(self).GetAction(gc))
}

func (self Go) SetAction(value gdclass.OpenXRAction) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAction(value)
}

func (self Go) Paths() []string {
	gc := gd.GarbageCollector(); _ = gc
		return []string(class(self).GetPaths(gc).Strings(gc))
}

func (self Go) SetPaths(value []string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPaths(gc.PackedStringSlice(value))
}

//go:nosplit
func (self class) SetAction(action gdclass.OpenXRAction)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRIPBinding.Bind_set_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAction(ctx gd.Lifetime) gdclass.OpenXRAction {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRIPBinding.Bind_get_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.OpenXRAction
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Get the number of input/output paths in this binding.
*/
//go:nosplit
func (self class) GetPathCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRIPBinding.Bind_get_path_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPaths(paths gd.PackedStringArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(paths))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRIPBinding.Bind_set_paths, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPaths(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRIPBinding.Bind_get_paths, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this input/output path is part of this binding.
*/
//go:nosplit
func (self class) HasPath(path gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRIPBinding.Bind_has_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Add an input/output path to this binding.
*/
//go:nosplit
func (self class) AddPath(path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRIPBinding.Bind_add_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes this input/output path from this binding.
*/
//go:nosplit
func (self class) RemovePath(path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRIPBinding.Bind_remove_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsOpenXRIPBinding() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsOpenXRIPBinding() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("OpenXRIPBinding", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
