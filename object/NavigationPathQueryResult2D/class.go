package NavigationPathQueryResult2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class stores the result of a 2D navigation path query from the [NavigationServer2D].

*/
type Simple [1]classdb.NavigationPathQueryResult2D
func (self Simple) SetPath(path gd.PackedVector2Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPath(path)
}
func (self Simple) GetPath() gd.PackedVector2Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector2Array(Expert(self).GetPath(gc))
}
func (self Simple) SetPathTypes(path_types gd.PackedInt32Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPathTypes(path_types)
}
func (self Simple) GetPathTypes() gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).GetPathTypes(gc))
}
func (self Simple) SetPathRids(path_rids gd.ArrayOf[gd.RID]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPathRids(path_rids)
}
func (self Simple) GetPathRids() gd.ArrayOf[gd.RID] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.RID](Expert(self).GetPathRids(gc))
}
func (self Simple) SetPathOwnerIds(path_owner_ids gd.PackedInt64Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPathOwnerIds(path_owner_ids)
}
func (self Simple) GetPathOwnerIds() gd.PackedInt64Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt64Array(Expert(self).GetPathOwnerIds(gc))
}
func (self Simple) Reset() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Reset()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.NavigationPathQueryResult2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetPath(path gd.PackedVector2Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryResult2D.Bind_set_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPath(ctx gd.Lifetime) gd.PackedVector2Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryResult2D.Bind_get_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector2Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathTypes(path_types gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path_types))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryResult2D.Bind_set_path_types, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathTypes(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryResult2D.Bind_get_path_types, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryResult2D.Bind_set_path_rids, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathRids(ctx gd.Lifetime) gd.ArrayOf[gd.RID] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryResult2D.Bind_get_path_rids, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryResult2D.Bind_set_path_owner_ids, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathOwnerIds(ctx gd.Lifetime) gd.PackedInt64Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryResult2D.Bind_get_path_owner_ids, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryResult2D.Bind_reset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsNavigationPathQueryResult2D() Expert { return self[0].AsNavigationPathQueryResult2D() }


//go:nosplit
func (self Simple) AsNavigationPathQueryResult2D() Simple { return self[0].AsNavigationPathQueryResult2D() }


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
func init() {classdb.Register("NavigationPathQueryResult2D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type PathSegmentType = classdb.NavigationPathQueryResult2DPathSegmentType

const (
/*This segment of the path goes through a region.*/
	PathSegmentTypeRegion PathSegmentType = 0
/*This segment of the path goes through a link.*/
	PathSegmentTypeLink PathSegmentType = 1
)
