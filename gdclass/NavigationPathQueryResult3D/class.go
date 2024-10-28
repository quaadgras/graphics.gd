package NavigationPathQueryResult3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
This class stores the result of a 3D navigation path query from the [NavigationServer3D].

*/
type Go [1]classdb.NavigationPathQueryResult3D

/*
Reset the result object to its initial state. This is useful to reuse the object across multiple queries.
*/
func (self Go) Reset() {
	class(self).Reset()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.NavigationPathQueryResult3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NavigationPathQueryResult3D"))
	return Go{classdb.NavigationPathQueryResult3D(object)}
}

func (self Go) Path() []gd.Vector3 {
		return []gd.Vector3(class(self).GetPath().AsSlice())
}

func (self Go) SetPath(value []gd.Vector3) {
	class(self).SetPath(gd.NewPackedVector3Slice(value))
}

func (self Go) PathTypes() []int32 {
		return []int32(class(self).GetPathTypes().AsSlice())
}

func (self Go) SetPathTypes(value []int32) {
	class(self).SetPathTypes(gd.NewPackedInt32Slice(value))
}

func (self Go) PathRids() gd.Array {
		return gd.Array(class(self).GetPathRids())
}

func (self Go) SetPathRids(value gd.Array) {
	class(self).SetPathRids(value)
}

func (self Go) PathOwnerIds() []int64 {
		return []int64(class(self).GetPathOwnerIds().AsSlice())
}

func (self Go) SetPathOwnerIds(value []int64) {
	class(self).SetPathOwnerIds(gd.NewPackedInt64Slice(value))
}

//go:nosplit
func (self class) SetPath(path gd.PackedVector3Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryResult3D.Bind_set_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPath() gd.PackedVector3Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryResult3D.Bind_get_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedVector3Array](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathTypes(path_types gd.PackedInt32Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(path_types))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryResult3D.Bind_set_path_types, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathTypes() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryResult3D.Bind_get_path_types, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathRids(path_rids gd.Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(path_rids))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryResult3D.Bind_set_path_rids, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathRids() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryResult3D.Bind_get_path_rids, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathOwnerIds(path_owner_ids gd.PackedInt64Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(path_owner_ids))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryResult3D.Bind_set_path_owner_ids, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathOwnerIds() gd.PackedInt64Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryResult3D.Bind_get_path_owner_ids, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedInt64Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Reset the result object to its initial state. This is useful to reuse the object across multiple queries.
*/
//go:nosplit
func (self class) Reset()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryResult3D.Bind_reset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func init() {classdb.Register("NavigationPathQueryResult3D", func(ptr gd.Object) any { return classdb.NavigationPathQueryResult3D(ptr) })}
type PathSegmentType = classdb.NavigationPathQueryResult3DPathSegmentType

const (
/*This segment of the path goes through a region.*/
	PathSegmentTypeRegion PathSegmentType = 0
/*This segment of the path goes through a link.*/
	PathSegmentTypeLink PathSegmentType = 1
)
