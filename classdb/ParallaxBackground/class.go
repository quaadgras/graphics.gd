// Package ParallaxBackground provides methods for working with ParallaxBackground object instances.
package ParallaxBackground

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/CanvasLayer"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
A ParallaxBackground uses one or more [ParallaxLayer] child nodes to create a parallax effect. Each [ParallaxLayer] can move at a different speed using [member ParallaxLayer.motion_offset]. This creates an illusion of depth in a 2D game. If not used with a [Camera2D], you must manually calculate the [member scroll_offset].
[b]Note:[/b] Each [ParallaxBackground] is drawn on one specific [Viewport] and cannot be shared between multiple [Viewport]s, see [member CanvasLayer.custom_viewport]. When using multiple [Viewport]s, for example in a split-screen game, you need create an individual [ParallaxBackground] for each [Viewport] you want it to be drawn on.
*/
type Instance [1]gdclass.ParallaxBackground

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsParallaxBackground() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ParallaxBackground

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ParallaxBackground"))
	casted := Instance{*(*gdclass.ParallaxBackground)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) ScrollOffset() Vector2.XY {
	return Vector2.XY(class(self).GetScrollOffset())
}

func (self Instance) SetScrollOffset(value Vector2.XY) {
	class(self).SetScrollOffset(gd.Vector2(value))
}

func (self Instance) ScrollBaseOffset() Vector2.XY {
	return Vector2.XY(class(self).GetScrollBaseOffset())
}

func (self Instance) SetScrollBaseOffset(value Vector2.XY) {
	class(self).SetScrollBaseOffset(gd.Vector2(value))
}

func (self Instance) ScrollBaseScale() Vector2.XY {
	return Vector2.XY(class(self).GetScrollBaseScale())
}

func (self Instance) SetScrollBaseScale(value Vector2.XY) {
	class(self).SetScrollBaseScale(gd.Vector2(value))
}

func (self Instance) ScrollLimitBegin() Vector2.XY {
	return Vector2.XY(class(self).GetLimitBegin())
}

func (self Instance) SetScrollLimitBegin(value Vector2.XY) {
	class(self).SetLimitBegin(gd.Vector2(value))
}

func (self Instance) ScrollLimitEnd() Vector2.XY {
	return Vector2.XY(class(self).GetLimitEnd())
}

func (self Instance) SetScrollLimitEnd(value Vector2.XY) {
	class(self).SetLimitEnd(gd.Vector2(value))
}

func (self Instance) ScrollIgnoreCameraZoom() bool {
	return bool(class(self).IsIgnoreCameraZoom())
}

func (self Instance) SetScrollIgnoreCameraZoom(value bool) {
	class(self).SetIgnoreCameraZoom(value)
}

//go:nosplit
func (self class) SetScrollOffset(offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParallaxBackground.Bind_set_scroll_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScrollOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParallaxBackground.Bind_get_scroll_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScrollBaseOffset(offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParallaxBackground.Bind_set_scroll_base_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScrollBaseOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParallaxBackground.Bind_get_scroll_base_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScrollBaseScale(scale gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParallaxBackground.Bind_set_scroll_base_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScrollBaseScale() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParallaxBackground.Bind_get_scroll_base_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLimitBegin(offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParallaxBackground.Bind_set_limit_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLimitBegin() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParallaxBackground.Bind_get_limit_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLimitEnd(offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParallaxBackground.Bind_set_limit_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLimitEnd() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParallaxBackground.Bind_get_limit_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIgnoreCameraZoom(ignore bool) {
	var frame = callframe.New()
	callframe.Arg(frame, ignore)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParallaxBackground.Bind_set_ignore_camera_zoom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsIgnoreCameraZoom() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParallaxBackground.Bind_is_ignore_camera_zoom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsParallaxBackground() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsParallaxBackground() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasLayer() CanvasLayer.Advanced {
	return *((*CanvasLayer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasLayer() CanvasLayer.Instance {
	return *((*CanvasLayer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(CanvasLayer.Advanced(self.AsCanvasLayer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(CanvasLayer.Instance(self.AsCanvasLayer()), name)
	}
}
func init() {
	gdclass.Register("ParallaxBackground", func(ptr gd.Object) any {
		return [1]gdclass.ParallaxBackground{*(*gdclass.ParallaxBackground)(unsafe.Pointer(&ptr))}
	})
}
