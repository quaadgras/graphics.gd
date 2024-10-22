package Shape3D

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
Abstract base class for all 3D shapes, intended for use in physics.
[b]Performance:[/b] Primitive shapes, especially [SphereShape3D], are fast to check collisions against. [ConvexPolygonShape3D] and [HeightMapShape3D] are slower, and [ConcavePolygonShape3D] is the slowest.

*/
type Go [1]classdb.Shape3D

/*
Returns the [ArrayMesh] used to draw the debug collision for this [Shape3D].
*/
func (self Go) GetDebugMesh() gdclass.ArrayMesh {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.ArrayMesh(class(self).GetDebugMesh(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Shape3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Shape3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) CustomSolverBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetCustomSolverBias()))
}

func (self Go) SetCustomSolverBias(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCustomSolverBias(gd.Float(value))
}

func (self Go) Margin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMargin()))
}

func (self Go) SetMargin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMargin(gd.Float(value))
}

//go:nosplit
func (self class) SetCustomSolverBias(bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shape3D.Bind_set_custom_solver_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCustomSolverBias() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shape3D.Bind_get_custom_solver_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMargin(margin gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shape3D.Bind_set_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMargin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shape3D.Bind_get_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [ArrayMesh] used to draw the debug collision for this [Shape3D].
*/
//go:nosplit
func (self class) GetDebugMesh(ctx gd.Lifetime) gdclass.ArrayMesh {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shape3D.Bind_get_debug_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.ArrayMesh
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsShape3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsShape3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("Shape3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
