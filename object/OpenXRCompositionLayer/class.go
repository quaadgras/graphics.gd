package OpenXRCompositionLayer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Composition layers allow 2D viewports to be displayed inside of the headset by the XR compositor through special projections that retain their quality. This allows for rendering clear text while keeping the layer at a native resolution.
[b]Note:[/b] If the OpenXR runtime doesn't support the given composition layer type, a fallback mesh can be generated with a [ViewportTexture], in order to emulate the composition layer.

*/
type Simple [1]classdb.OpenXRCompositionLayer
func (self Simple) SetLayerViewport(viewport [1]classdb.SubViewport) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLayerViewport(viewport)
}
func (self Simple) GetLayerViewport() [1]classdb.SubViewport {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.SubViewport(Expert(self).GetLayerViewport(gc))
}
func (self Simple) SetEnableHolePunch(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnableHolePunch(enable)
}
func (self Simple) GetEnableHolePunch() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetEnableHolePunch())
}
func (self Simple) SetSortOrder(order int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSortOrder(gd.Int(order))
}
func (self Simple) GetSortOrder() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSortOrder()))
}
func (self Simple) SetAlphaBlend(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAlphaBlend(enabled)
}
func (self Simple) GetAlphaBlend() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetAlphaBlend())
}
func (self Simple) IsNativelySupported() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsNativelySupported())
}
func (self Simple) IntersectsRay(origin gd.Vector3, direction gd.Vector3) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).IntersectsRay(origin, direction))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.OpenXRCompositionLayer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetLayerViewport(viewport object.SubViewport)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(viewport[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayer.Bind_set_layer_viewport, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLayerViewport(ctx gd.Lifetime) object.SubViewport {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayer.Bind_get_layer_viewport, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.SubViewport
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnableHolePunch(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayer.Bind_set_enable_hole_punch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnableHolePunch() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayer.Bind_get_enable_hole_punch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSortOrder(order gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, order)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayer.Bind_set_sort_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSortOrder() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayer.Bind_get_sort_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlphaBlend(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayer.Bind_set_alpha_blend, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlphaBlend() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayer.Bind_get_alpha_blend, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns true if the OpenXR runtime natively supports this composition layer type.
[b]Note:[/b] This will only return an accurate result after the OpenXR session has started.
*/
//go:nosplit
func (self class) IsNativelySupported() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayer.Bind_is_natively_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns UV coordinates where the given ray intersects with the composition layer. [param origin] and [param direction] must be in global space.
Returns [code]Vector2(-1.0, -1.0)[/code] if the ray doesn't intersect.
*/
//go:nosplit
func (self class) IntersectsRay(origin gd.Vector3, direction gd.Vector3) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, origin)
	callframe.Arg(frame, direction)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayer.Bind_intersects_ray, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsOpenXRCompositionLayer() Expert { return self[0].AsOpenXRCompositionLayer() }


//go:nosplit
func (self Simple) AsOpenXRCompositionLayer() Simple { return self[0].AsOpenXRCompositionLayer() }


//go:nosplit
func (self class) AsNode3D() Node3D.Expert { return self[0].AsNode3D() }


//go:nosplit
func (self Simple) AsNode3D() Node3D.Simple { return self[0].AsNode3D() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

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
func init() {classdb.Register("OpenXRCompositionLayer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
