package LightOccluder2D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Node2D"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Occludes light cast by a Light2D, casting shadows. The LightOccluder2D must be provided with an [OccluderPolygon2D] in order for the shadow to be computed.

*/
type Simple [1]classdb.LightOccluder2D
func (self Simple) SetOccluderPolygon(polygon [1]classdb.OccluderPolygon2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOccluderPolygon(polygon)
}
func (self Simple) GetOccluderPolygon() [1]classdb.OccluderPolygon2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.OccluderPolygon2D(Expert(self).GetOccluderPolygon(gc))
}
func (self Simple) SetOccluderLightMask(mask int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOccluderLightMask(gd.Int(mask))
}
func (self Simple) GetOccluderLightMask() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetOccluderLightMask()))
}
func (self Simple) SetAsSdfCollision(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAsSdfCollision(enable)
}
func (self Simple) IsSetAsSdfCollision() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSetAsSdfCollision())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.LightOccluder2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetOccluderPolygon(polygon object.OccluderPolygon2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(polygon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightOccluder2D.Bind_set_occluder_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOccluderPolygon(ctx gd.Lifetime) object.OccluderPolygon2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightOccluder2D.Bind_get_occluder_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.OccluderPolygon2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOccluderLightMask(mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightOccluder2D.Bind_set_occluder_light_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOccluderLightMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightOccluder2D.Bind_get_occluder_light_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAsSdfCollision(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightOccluder2D.Bind_set_as_sdf_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSetAsSdfCollision() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightOccluder2D.Bind_is_set_as_sdf_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsLightOccluder2D() Expert { return self[0].AsLightOccluder2D() }


//go:nosplit
func (self Simple) AsLightOccluder2D() Simple { return self[0].AsLightOccluder2D() }


//go:nosplit
func (self class) AsNode2D() Node2D.Expert { return self[0].AsNode2D() }


//go:nosplit
func (self Simple) AsNode2D() Node2D.Simple { return self[0].AsNode2D() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("LightOccluder2D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
