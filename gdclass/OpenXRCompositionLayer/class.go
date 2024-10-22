package OpenXRCompositionLayer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Composition layers allow 2D viewports to be displayed inside of the headset by the XR compositor through special projections that retain their quality. This allows for rendering clear text while keeping the layer at a native resolution.
[b]Note:[/b] If the OpenXR runtime doesn't support the given composition layer type, a fallback mesh can be generated with a [ViewportTexture], in order to emulate the composition layer.

*/
type Go [1]classdb.OpenXRCompositionLayer

/*
Returns true if the OpenXR runtime natively supports this composition layer type.
[b]Note:[/b] This will only return an accurate result after the OpenXR session has started.
*/
func (self Go) IsNativelySupported() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsNativelySupported())
}

/*
Returns UV coordinates where the given ray intersects with the composition layer. [param origin] and [param direction] must be in global space.
Returns [code]Vector2(-1.0, -1.0)[/code] if the ray doesn't intersect.
*/
func (self Go) IntersectsRay(origin gd.Vector3, direction gd.Vector3) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).IntersectsRay(origin, direction))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.OpenXRCompositionLayer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("OpenXRCompositionLayer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) LayerViewport() gdclass.SubViewport {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.SubViewport(class(self).GetLayerViewport(gc))
}

func (self Go) SetLayerViewport(value gdclass.SubViewport) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLayerViewport(value)
}

func (self Go) SortOrder() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSortOrder()))
}

func (self Go) SetSortOrder(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSortOrder(gd.Int(value))
}

func (self Go) AlphaBlend() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetAlphaBlend())
}

func (self Go) SetAlphaBlend(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAlphaBlend(value)
}

func (self Go) EnableHolePunch() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetEnableHolePunch())
}

func (self Go) SetEnableHolePunch(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnableHolePunch(value)
}

//go:nosplit
func (self class) SetLayerViewport(viewport gdclass.SubViewport)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(viewport[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayer.Bind_set_layer_viewport, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLayerViewport(ctx gd.Lifetime) gdclass.SubViewport {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRCompositionLayer.Bind_get_layer_viewport, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.SubViewport
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
func (self class) AsOpenXRCompositionLayer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsOpenXRCompositionLayer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {classdb.Register("OpenXRCompositionLayer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
