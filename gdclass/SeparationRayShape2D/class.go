package SeparationRayShape2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Shape2D"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A 2D ray shape, intended for use in physics. Usually used to provide a shape for a [CollisionShape2D]. When a [SeparationRayShape2D] collides with an object, it tries to separate itself from it by moving its endpoint to the collision point. For example, a [SeparationRayShape2D] next to a character can allow it to instantly move up when touching stairs.

*/
type Go [1]classdb.SeparationRayShape2D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SeparationRayShape2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SeparationRayShape2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Length() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetLength()))
}

func (self Go) SetLength(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLength(gd.Float(value))
}

func (self Go) SlideOnSlope() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetSlideOnSlope())
}

func (self Go) SetSlideOnSlope(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSlideOnSlope(value)
}

//go:nosplit
func (self class) SetLength(length gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SeparationRayShape2D.Bind_set_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SeparationRayShape2D.Bind_get_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSlideOnSlope(active bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SeparationRayShape2D.Bind_set_slide_on_slope, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSlideOnSlope() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SeparationRayShape2D.Bind_get_slide_on_slope, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSeparationRayShape2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSeparationRayShape2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsShape2D() Shape2D.GD { return *((*Shape2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsShape2D() Shape2D.Go { return *((*Shape2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsShape2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsShape2D(), name)
	}
}
func init() {classdb.Register("SeparationRayShape2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
