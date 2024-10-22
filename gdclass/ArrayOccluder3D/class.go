package ArrayOccluder3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Occluder3D"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[ArrayOccluder3D] stores an arbitrary 3D polygon shape that can be used by the engine's occlusion culling system. This is analogous to [ArrayMesh], but for occluders.
See [OccluderInstance3D]'s documentation for instructions on setting up occlusion culling.

*/
type Go [1]classdb.ArrayOccluder3D

/*
Sets [member indices] and [member vertices], while updating the final occluder only once after both values are set.
*/
func (self Go) SetArrays(vertices []gd.Vector3, indices []int32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetArrays(gc.PackedVector3Slice(vertices), gc.PackedInt32Slice(indices))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ArrayOccluder3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ArrayOccluder3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Sets [member indices] and [member vertices], while updating the final occluder only once after both values are set.
*/
//go:nosplit
func (self class) SetArrays(vertices gd.PackedVector3Array, indices gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(vertices))
	callframe.Arg(frame, mmm.Get(indices))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ArrayOccluder3D.Bind_set_arrays, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetVertices(vertices gd.PackedVector3Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(vertices))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ArrayOccluder3D.Bind_set_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetIndices(indices gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(indices))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ArrayOccluder3D.Bind_set_indices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsArrayOccluder3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsArrayOccluder3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsOccluder3D() Occluder3D.GD { return *((*Occluder3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsOccluder3D() Occluder3D.Go { return *((*Occluder3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsOccluder3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsOccluder3D(), name)
	}
}
func init() {classdb.Register("ArrayOccluder3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
