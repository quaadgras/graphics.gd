package OccluderPolygon2D

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
Editor facility that helps you draw a 2D polygon used as resource for [LightOccluder2D].

*/
type Go [1]classdb.OccluderPolygon2D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.OccluderPolygon2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("OccluderPolygon2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Closed() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsClosed())
}

func (self Go) SetClosed(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetClosed(value)
}

func (self Go) CullMode() classdb.OccluderPolygon2DCullMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.OccluderPolygon2DCullMode(class(self).GetCullMode())
}

func (self Go) SetCullMode(value classdb.OccluderPolygon2DCullMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCullMode(value)
}

func (self Go) Polygon() []gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return []gd.Vector2(class(self).GetPolygon(gc).AsSlice())
}

func (self Go) SetPolygon(value []gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPolygon(gc.PackedVector2Slice(value))
}

//go:nosplit
func (self class) SetClosed(closed bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, closed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OccluderPolygon2D.Bind_set_closed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsClosed() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OccluderPolygon2D.Bind_is_closed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCullMode(cull_mode classdb.OccluderPolygon2DCullMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cull_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OccluderPolygon2D.Bind_set_cull_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCullMode() classdb.OccluderPolygon2DCullMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.OccluderPolygon2DCullMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OccluderPolygon2D.Bind_get_cull_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPolygon(polygon gd.PackedVector2Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(polygon))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OccluderPolygon2D.Bind_set_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPolygon(ctx gd.Lifetime) gd.PackedVector2Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OccluderPolygon2D.Bind_get_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector2Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsOccluderPolygon2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsOccluderPolygon2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("OccluderPolygon2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type CullMode = classdb.OccluderPolygon2DCullMode

const (
/*Culling is disabled. See [member cull_mode].*/
	CullDisabled CullMode = 0
/*Culling is performed in the clockwise direction. See [member cull_mode].*/
	CullClockwise CullMode = 1
/*Culling is performed in the counterclockwise direction. See [member cull_mode].*/
	CullCounterClockwise CullMode = 2
)
