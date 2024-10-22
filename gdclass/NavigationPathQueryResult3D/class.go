package NavigationPathQueryResult3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class stores the result of a 3D navigation path query from the [NavigationServer3D].

*/
type Go [1]classdb.NavigationPathQueryResult3D

/*
Reset the result object to its initial state. This is useful to reuse the object across multiple queries.
*/
func (self Go) Reset() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Reset()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.NavigationPathQueryResult3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("NavigationPathQueryResult3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Path() []gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return []gd.Vector3(class(self).GetPath(gc).AsSlice())
}

func (self Go) SetPath(value []gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPath(gc.PackedVector3Slice(value))
}

func (self Go) PathTypes() []int32 {
	gc := gd.GarbageCollector(); _ = gc
		return []int32(class(self).GetPathTypes(gc).AsSlice())
}

func (self Go) SetPathTypes(value []int32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathTypes(gc.PackedInt32Slice(value))
}

func (self Go) PathRids() gd.ArrayOf[gd.RID] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gd.RID](class(self).GetPathRids(gc))
}

func (self Go) SetPathRids(value gd.ArrayOf[gd.RID]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathRids(value)
}

func (self Go) PathOwnerIds() []int64 {
	gc := gd.GarbageCollector(); _ = gc
		return []int64(class(self).GetPathOwnerIds(gc).AsSlice())
}

func (self Go) SetPathOwnerIds(value []int64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathOwnerIds(gc.PackedInt64Slice(value))
}

//go:nosplit
func (self class) SetPath(path gd.PackedVector3Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryResult3D.Bind_set_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPath(ctx gd.Lifetime) gd.PackedVector3Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryResult3D.Bind_get_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector3Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathTypes(path_types gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path_types))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryResult3D.Bind_set_path_types, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathTypes(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryResult3D.Bind_get_path_types, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathRids(path_rids gd.ArrayOf[gd.RID])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path_rids))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryResult3D.Bind_set_path_rids, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathRids(ctx gd.Lifetime) gd.ArrayOf[gd.RID] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryResult3D.Bind_get_path_rids, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.RID](ret)
}
//go:nosplit
func (self class) SetPathOwnerIds(path_owner_ids gd.PackedInt64Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path_owner_ids))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryResult3D.Bind_set_path_owner_ids, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathOwnerIds(ctx gd.Lifetime) gd.PackedInt64Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryResult3D.Bind_get_path_owner_ids, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt64Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Reset the result object to its initial state. This is useful to reuse the object across multiple queries.
*/
//go:nosplit
func (self class) Reset()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryResult3D.Bind_reset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsNavigationPathQueryResult3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsNavigationPathQueryResult3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("NavigationPathQueryResult3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type PathSegmentType = classdb.NavigationPathQueryResult3DPathSegmentType

const (
/*This segment of the path goes through a region.*/
	PathSegmentTypeRegion PathSegmentType = 0
/*This segment of the path goes through a link.*/
	PathSegmentTypeLink PathSegmentType = 1
)
