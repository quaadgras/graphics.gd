// Package OpenXRCompositionLayer provides methods for working with OpenXRCompositionLayer object instances.
package OpenXRCompositionLayer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Vector2"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Composition layers allow 2D viewports to be displayed inside of the headset by the XR compositor through special projections that retain their quality. This allows for rendering clear text while keeping the layer at a native resolution.
[b]Note:[/b] If the OpenXR runtime doesn't support the given composition layer type, a fallback mesh can be generated with a [ViewportTexture], in order to emulate the composition layer.
*/
type Instance [1]gdclass.OpenXRCompositionLayer
type Any interface {
	gd.IsClass
	AsOpenXRCompositionLayer() Instance
}

/*
Returns true if the OpenXR runtime natively supports this composition layer type.
[b]Note:[/b] This will only return an accurate result after the OpenXR session has started.
*/
func (self Instance) IsNativelySupported() bool {
	return bool(class(self).IsNativelySupported())
}

/*
Returns UV coordinates where the given ray intersects with the composition layer. [param origin] and [param direction] must be in global space.
Returns [code]Vector2(-1.0, -1.0)[/code] if the ray doesn't intersect.
*/
func (self Instance) IntersectsRay(origin Vector3.XYZ, direction Vector3.XYZ) Vector2.XY {
	return Vector2.XY(class(self).IntersectsRay(gd.Vector3(origin), gd.Vector3(direction)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.OpenXRCompositionLayer

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRCompositionLayer"))
	return Instance{*(*gdclass.OpenXRCompositionLayer)(unsafe.Pointer(&object))}
}

func (self Instance) LayerViewport() [1]gdclass.SubViewport {
	return [1]gdclass.SubViewport(class(self).GetLayerViewport())
}

func (self Instance) SetLayerViewport(value [1]gdclass.SubViewport) {
	class(self).SetLayerViewport(value)
}

func (self Instance) SortOrder() int {
	return int(int(class(self).GetSortOrder()))
}

func (self Instance) SetSortOrder(value int) {
	class(self).SetSortOrder(gd.Int(value))
}

func (self Instance) AlphaBlend() bool {
	return bool(class(self).GetAlphaBlend())
}

func (self Instance) SetAlphaBlend(value bool) {
	class(self).SetAlphaBlend(value)
}

func (self Instance) EnableHolePunch() bool {
	return bool(class(self).GetEnableHolePunch())
}

func (self Instance) SetEnableHolePunch(value bool) {
	class(self).SetEnableHolePunch(value)
}

//go:nosplit
func (self class) SetLayerViewport(viewport [1]gdclass.SubViewport) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(viewport[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayer.Bind_set_layer_viewport, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLayerViewport() [1]gdclass.SubViewport {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayer.Bind_get_layer_viewport, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.SubViewport{gd.PointerWithOwnershipTransferredToGo[gdclass.SubViewport](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableHolePunch(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayer.Bind_set_enable_hole_punch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableHolePunch() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayer.Bind_get_enable_hole_punch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSortOrder(order gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, order)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayer.Bind_set_sort_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSortOrder() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayer.Bind_get_sort_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlphaBlend(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayer.Bind_set_alpha_blend, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlphaBlend() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayer.Bind_get_alpha_blend, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayer.Bind_is_natively_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, origin)
	callframe.Arg(frame, direction)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayer.Bind_intersects_ray, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsOpenXRCompositionLayer() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOpenXRCompositionLayer() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {
	gdclass.Register("OpenXRCompositionLayer", func(ptr gd.Object) any {
		return [1]gdclass.OpenXRCompositionLayer{*(*gdclass.OpenXRCompositionLayer)(unsafe.Pointer(&ptr))}
	})
}
