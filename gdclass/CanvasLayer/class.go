package CanvasLayer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[CanvasItem]-derived nodes that are direct or indirect children of a [CanvasLayer] will be drawn in that layer. The layer is a numeric index that defines the draw order. The default 2D scene renders with index [code]0[/code], so a [CanvasLayer] with index [code]-1[/code] will be drawn below, and a [CanvasLayer] with index [code]1[/code] will be drawn above. This order will hold regardless of the [member CanvasItem.z_index] of the nodes within each layer.
[CanvasLayer]s can be hidden and they can also optionally follow the viewport. This makes them useful for HUDs like health bar overlays (on layers [code]1[/code] and higher) or backgrounds (on layers [code]-1[/code] and lower).
[b]Note:[/b] Embedded [Window]s are placed on layer [code]1024[/code]. [CanvasItem]s on layers [code]1025[/code] and higher appear in front of embedded windows.
[b]Note:[/b] Each [CanvasLayer] is drawn on one specific [Viewport] and cannot be shared between multiple [Viewport]s, see [member custom_viewport]. When using multiple [Viewport]s, for example in a split-screen game, you need create an individual [CanvasLayer] for each [Viewport] you want it to be drawn on.

*/
type Go [1]classdb.CanvasLayer

/*
Shows any [CanvasItem] under this [CanvasLayer]. This is equivalent to setting [member visible] to [code]true[/code].
*/
func (self Go) Show() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Show()
}

/*
Hides any [CanvasItem] under this [CanvasLayer]. This is equivalent to setting [member visible] to [code]false[/code].
*/
func (self Go) Hide() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Hide()
}

/*
Returns the transform from the [CanvasLayer]s coordinate system to the [Viewport]s coordinate system.
*/
func (self Go) GetFinalTransform() gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(class(self).GetFinalTransform())
}

/*
Returns the RID of the canvas used by this layer.
*/
func (self Go) GetCanvas() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetCanvas())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CanvasLayer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("CanvasLayer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Layer() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetLayer()))
}

func (self Go) SetLayer(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLayer(gd.Int(value))
}

func (self Go) Visible() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsVisible())
}

func (self Go) SetVisible(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisible(value)
}

func (self Go) Offset() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetOffset())
}

func (self Go) SetOffset(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOffset(value)
}

func (self Go) Rotation() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetRotation()))
}

func (self Go) SetRotation(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRotation(gd.Float(value))
}

func (self Go) Scale() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetScale())
}

func (self Go) SetScale(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetScale(value)
}

func (self Go) Transform() gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Transform2D(class(self).GetTransform())
}

func (self Go) SetTransform(value gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTransform(value)
}

func (self Go) CustomViewport() gdclass.Node {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Node(class(self).GetCustomViewport(gc))
}

func (self Go) SetCustomViewport(value gdclass.Node) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCustomViewport(value)
}

func (self Go) FollowViewportEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsFollowingViewport())
}

func (self Go) SetFollowViewportEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFollowViewport(value)
}

func (self Go) FollowViewportScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFollowViewportScale()))
}

func (self Go) SetFollowViewportScale(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFollowViewportScale(gd.Float(value))
}

//go:nosplit
func (self class) SetLayer(layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_set_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLayer() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_get_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisible(visible bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_set_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsVisible() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_is_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Shows any [CanvasItem] under this [CanvasLayer]. This is equivalent to setting [member visible] to [code]true[/code].
*/
//go:nosplit
func (self class) Show()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_show, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Hides any [CanvasItem] under this [CanvasLayer]. This is equivalent to setting [member visible] to [code]false[/code].
*/
//go:nosplit
func (self class) Hide()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_hide, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetTransform(transform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTransform() gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the transform from the [CanvasLayer]s coordinate system to the [Viewport]s coordinate system.
*/
//go:nosplit
func (self class) GetFinalTransform() gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_get_final_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOffset(offset gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOffset() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRotation(radians gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radians)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_set_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRotation() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_get_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScale(scale gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_set_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetScale() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_get_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFollowViewport(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_set_follow_viewport, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFollowingViewport() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_is_following_viewport, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFollowViewportScale(scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_set_follow_viewport_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFollowViewportScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_get_follow_viewport_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCustomViewport(viewport gdclass.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(viewport[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_set_custom_viewport, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCustomViewport(ctx gd.Lifetime) gdclass.Node {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_get_custom_viewport, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Node
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the RID of the canvas used by this layer.
*/
//go:nosplit
func (self class) GetCanvas() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasLayer.Bind_get_canvas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnVisibilityChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("visibility_changed"), gc.Callable(cb), 0)
}


func (self class) AsCanvasLayer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasLayer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}
func init() {classdb.Register("CanvasLayer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
