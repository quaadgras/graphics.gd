// Package CanvasLayer provides methods for working with CanvasLayer object instances.
package CanvasLayer

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Packed"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Transform2D"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ = slices.Delete[[]struct{}, struct{}]

/*
[CanvasItem]-derived nodes that are direct or indirect children of a [CanvasLayer] will be drawn in that layer. The layer is a numeric index that defines the draw order. The default 2D scene renders with index [code]0[/code], so a [CanvasLayer] with index [code]-1[/code] will be drawn below, and a [CanvasLayer] with index [code]1[/code] will be drawn above. This order will hold regardless of the [member CanvasItem.z_index] of the nodes within each layer.
[CanvasLayer]s can be hidden and they can also optionally follow the viewport. This makes them useful for HUDs like health bar overlays (on layers [code]1[/code] and higher) or backgrounds (on layers [code]-1[/code] and lower).
[b]Note:[/b] Embedded [Window]s are placed on layer [code]1024[/code]. [CanvasItem]s on layers [code]1025[/code] and higher appear in front of embedded windows.
[b]Note:[/b] Each [CanvasLayer] is drawn on one specific [Viewport] and cannot be shared between multiple [Viewport]s, see [member custom_viewport]. When using multiple [Viewport]s, for example in a split-screen game, you need create an individual [CanvasLayer] for each [Viewport] you want it to be drawn on.
*/
type Instance [1]gdclass.CanvasLayer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCanvasLayer() Instance
}

/*
Shows any [CanvasItem] under this [CanvasLayer]. This is equivalent to setting [member visible] to [code]true[/code].
*/
func (self Instance) Show() { //gd:CanvasLayer.show
	class(self).Show()
}

/*
Hides any [CanvasItem] under this [CanvasLayer]. This is equivalent to setting [member visible] to [code]false[/code].
*/
func (self Instance) Hide() { //gd:CanvasLayer.hide
	class(self).Hide()
}

/*
Returns the transform from the [CanvasLayer]s coordinate system to the [Viewport]s coordinate system.
*/
func (self Instance) GetFinalTransform() Transform2D.OriginXY { //gd:CanvasLayer.get_final_transform
	return Transform2D.OriginXY(class(self).GetFinalTransform())
}

/*
Returns the RID of the canvas used by this layer.
*/
func (self Instance) GetCanvas() RID.Canvas { //gd:CanvasLayer.get_canvas
	return RID.Canvas(class(self).GetCanvas())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CanvasLayer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CanvasLayer"))
	casted := Instance{*(*gdclass.CanvasLayer)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Layer() int {
	return int(int(class(self).GetLayer()))
}

func (self Instance) SetLayer(value int) {
	class(self).SetLayer(gd.Int(value))
}

func (self Instance) Visible() bool {
	return bool(class(self).IsVisible())
}

func (self Instance) SetVisible(value bool) {
	class(self).SetVisible(value)
}

func (self Instance) Offset() Vector2.XY {
	return Vector2.XY(class(self).GetOffset())
}

func (self Instance) SetOffset(value Vector2.XY) {
	class(self).SetOffset(gd.Vector2(value))
}

func (self Instance) Rotation() Float.X {
	return Float.X(Float.X(class(self).GetRotation()))
}

func (self Instance) SetRotation(value Float.X) {
	class(self).SetRotation(gd.Float(value))
}

func (self Instance) Scale() Vector2.XY {
	return Vector2.XY(class(self).GetScale())
}

func (self Instance) SetScale(value Vector2.XY) {
	class(self).SetScale(gd.Vector2(value))
}

func (self Instance) Transform() Transform2D.OriginXY {
	return Transform2D.OriginXY(class(self).GetTransform())
}

func (self Instance) SetTransform(value Transform2D.OriginXY) {
	class(self).SetTransform(gd.Transform2D(value))
}

func (self Instance) CustomViewport() [1]gdclass.Node {
	return [1]gdclass.Node(class(self).GetCustomViewport())
}

func (self Instance) SetCustomViewport(value [1]gdclass.Node) {
	class(self).SetCustomViewport(value)
}

func (self Instance) FollowViewportEnabled() bool {
	return bool(class(self).IsFollowingViewport())
}

func (self Instance) SetFollowViewportEnabled(value bool) {
	class(self).SetFollowViewport(value)
}

func (self Instance) FollowViewportScale() Float.X {
	return Float.X(Float.X(class(self).GetFollowViewportScale()))
}

func (self Instance) SetFollowViewportScale(value Float.X) {
	class(self).SetFollowViewportScale(gd.Float(value))
}

//go:nosplit
func (self class) SetLayer(layer gd.Int) { //gd:CanvasLayer.set_layer
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_set_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLayer() gd.Int { //gd:CanvasLayer.get_layer
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_get_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisible(visible bool) { //gd:CanvasLayer.set_visible
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_set_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsVisible() bool { //gd:CanvasLayer.is_visible
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_is_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Shows any [CanvasItem] under this [CanvasLayer]. This is equivalent to setting [member visible] to [code]true[/code].
*/
//go:nosplit
func (self class) Show() { //gd:CanvasLayer.show
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_show, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Hides any [CanvasItem] under this [CanvasLayer]. This is equivalent to setting [member visible] to [code]false[/code].
*/
//go:nosplit
func (self class) Hide() { //gd:CanvasLayer.hide
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_hide, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetTransform(transform gd.Transform2D) { //gd:CanvasLayer.set_transform
	var frame = callframe.New()
	callframe.Arg(frame, transform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTransform() gd.Transform2D { //gd:CanvasLayer.get_transform
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the transform from the [CanvasLayer]s coordinate system to the [Viewport]s coordinate system.
*/
//go:nosplit
func (self class) GetFinalTransform() gd.Transform2D { //gd:CanvasLayer.get_final_transform
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_get_final_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOffset(offset gd.Vector2) { //gd:CanvasLayer.set_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOffset() gd.Vector2 { //gd:CanvasLayer.get_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRotation(radians gd.Float) { //gd:CanvasLayer.set_rotation
	var frame = callframe.New()
	callframe.Arg(frame, radians)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_set_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRotation() gd.Float { //gd:CanvasLayer.get_rotation
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_get_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScale(scale gd.Vector2) { //gd:CanvasLayer.set_scale
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_set_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetScale() gd.Vector2 { //gd:CanvasLayer.get_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_get_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFollowViewport(enable bool) { //gd:CanvasLayer.set_follow_viewport
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_set_follow_viewport, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFollowingViewport() bool { //gd:CanvasLayer.is_following_viewport
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_is_following_viewport, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFollowViewportScale(scale gd.Float) { //gd:CanvasLayer.set_follow_viewport_scale
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_set_follow_viewport_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFollowViewportScale() gd.Float { //gd:CanvasLayer.get_follow_viewport_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_get_follow_viewport_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCustomViewport(viewport [1]gdclass.Node) { //gd:CanvasLayer.set_custom_viewport
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(viewport[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_set_custom_viewport, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCustomViewport() [1]gdclass.Node { //gd:CanvasLayer.get_custom_viewport
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_get_custom_viewport, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerMustAssertInstanceID[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the RID of the canvas used by this layer.
*/
//go:nosplit
func (self class) GetCanvas() gd.RID { //gd:CanvasLayer.get_canvas
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasLayer.Bind_get_canvas, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnVisibilityChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("visibility_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsCanvasLayer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCanvasLayer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced      { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance   { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Advanced(self.AsNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Instance(self.AsNode()), name)
	}
}
func init() {
	gdclass.Register("CanvasLayer", func(ptr gd.Object) any { return [1]gdclass.CanvasLayer{*(*gdclass.CanvasLayer)(unsafe.Pointer(&ptr))} })
}
