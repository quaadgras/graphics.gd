package OpenXRCompositionLayerEquirect

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/OpenXRCompositionLayer"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
An OpenXR composition layer that allows rendering a [SubViewport] on an internal slice of a sphere.

*/
type Go [1]classdb.OpenXRCompositionLayerEquirect
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.OpenXRCompositionLayerEquirect
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("OpenXRCompositionLayerEquirect"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Radius() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetRadius()))
}

func (self Go) SetRadius(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRadius(gd.Float(value))
}

func (self Go) CentralHorizontalAngle() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetCentralHorizontalAngle()))
}

func (self Go) SetCentralHorizontalAngle(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCentralHorizontalAngle(gd.Float(value))
}

func (self Go) UpperVerticalAngle() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetUpperVerticalAngle()))
}

func (self Go) SetUpperVerticalAngle(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUpperVerticalAngle(gd.Float(value))
}

func (self Go) LowerVerticalAngle() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetLowerVerticalAngle()))
}

func (self Go) SetLowerVerticalAngle(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLowerVerticalAngle(gd.Float(value))
}

func (self Go) FallbackSegments() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetFallbackSegments()))
}

func (self Go) SetFallbackSegments(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFallbackSegments(gd.Int(value))
}

//go:nosplit
func (self class) SetRadius(radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayerEquirect.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayerEquirect.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCentralHorizontalAngle(angle gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angle)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayerEquirect.Bind_set_central_horizontal_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCentralHorizontalAngle() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayerEquirect.Bind_get_central_horizontal_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUpperVerticalAngle(angle gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angle)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayerEquirect.Bind_set_upper_vertical_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUpperVerticalAngle() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayerEquirect.Bind_get_upper_vertical_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLowerVerticalAngle(angle gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angle)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayerEquirect.Bind_set_lower_vertical_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLowerVerticalAngle() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayerEquirect.Bind_get_lower_vertical_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFallbackSegments(segments gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, segments)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayerEquirect.Bind_set_fallback_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFallbackSegments() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayerEquirect.Bind_get_fallback_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsOpenXRCompositionLayerEquirect() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsOpenXRCompositionLayerEquirect() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsOpenXRCompositionLayer() OpenXRCompositionLayer.GD { return *((*OpenXRCompositionLayer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsOpenXRCompositionLayer() OpenXRCompositionLayer.Go { return *((*OpenXRCompositionLayer.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsOpenXRCompositionLayer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsOpenXRCompositionLayer(), name)
	}
}
func init() {classdb.Register("OpenXRCompositionLayerEquirect", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
