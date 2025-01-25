// Package TouchScreenButton provides methods for working with TouchScreenButton object instances.
package TouchScreenButton

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
TouchScreenButton allows you to create on-screen buttons for touch devices. It's intended for gameplay use, such as a unit you have to touch to move. Unlike [Button], TouchScreenButton supports multitouch out of the box. Several TouchScreenButtons can be pressed at the same time with touch input.
This node inherits from [Node2D]. Unlike with [Control] nodes, you cannot set anchors on it. If you want to create menus or user interfaces, you may want to use [Button] nodes instead. To make button nodes react to touch events, you can enable the Emulate Mouse option in the Project Settings.
You can configure TouchScreenButton to be visible only on touch devices, helping you develop your game both for desktop and mobile devices.
*/
type Instance [1]gdclass.TouchScreenButton

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTouchScreenButton() Instance
}

/*
Returns [code]true[/code] if this button is currently pressed.
*/
func (self Instance) IsPressed() bool {
	return bool(class(self).IsPressed())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TouchScreenButton

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TouchScreenButton"))
	casted := Instance{*(*gdclass.TouchScreenButton)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) TextureNormal() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTextureNormal())
}

func (self Instance) SetTextureNormal(value [1]gdclass.Texture2D) {
	class(self).SetTextureNormal(value)
}

func (self Instance) TexturePressed() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexturePressed())
}

func (self Instance) SetTexturePressed(value [1]gdclass.Texture2D) {
	class(self).SetTexturePressed(value)
}

func (self Instance) Bitmask() [1]gdclass.BitMap {
	return [1]gdclass.BitMap(class(self).GetBitmask())
}

func (self Instance) SetBitmask(value [1]gdclass.BitMap) {
	class(self).SetBitmask(value)
}

func (self Instance) Shape() [1]gdclass.Shape2D {
	return [1]gdclass.Shape2D(class(self).GetShape())
}

func (self Instance) SetShape(value [1]gdclass.Shape2D) {
	class(self).SetShape(value)
}

func (self Instance) ShapeCentered() bool {
	return bool(class(self).IsShapeCentered())
}

func (self Instance) SetShapeCentered(value bool) {
	class(self).SetShapeCentered(value)
}

func (self Instance) ShapeVisible() bool {
	return bool(class(self).IsShapeVisible())
}

func (self Instance) SetShapeVisible(value bool) {
	class(self).SetShapeVisible(value)
}

func (self Instance) PassbyPress() bool {
	return bool(class(self).IsPassbyPressEnabled())
}

func (self Instance) SetPassbyPress(value bool) {
	class(self).SetPassbyPress(value)
}

func (self Instance) Action() string {
	return string(class(self).GetAction().String())
}

func (self Instance) SetAction(value string) {
	class(self).SetAction(gd.NewString(value))
}

func (self Instance) VisibilityMode() gdclass.TouchScreenButtonVisibilityMode {
	return gdclass.TouchScreenButtonVisibilityMode(class(self).GetVisibilityMode())
}

func (self Instance) SetVisibilityMode(value gdclass.TouchScreenButtonVisibilityMode) {
	class(self).SetVisibilityMode(value)
}

//go:nosplit
func (self class) SetTextureNormal(texture [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TouchScreenButton.Bind_set_texture_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureNormal() [1]gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TouchScreenButton.Bind_get_texture_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTexturePressed(texture [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TouchScreenButton.Bind_set_texture_pressed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexturePressed() [1]gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TouchScreenButton.Bind_get_texture_pressed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBitmask(bitmask [1]gdclass.BitMap) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(bitmask[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TouchScreenButton.Bind_set_bitmask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBitmask() [1]gdclass.BitMap {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TouchScreenButton.Bind_get_bitmask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.BitMap{gd.PointerWithOwnershipTransferredToGo[gdclass.BitMap](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShape(shape [1]gdclass.Shape2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shape[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TouchScreenButton.Bind_set_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShape() [1]gdclass.Shape2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TouchScreenButton.Bind_get_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Shape2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Shape2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShapeCentered(b bool) {
	var frame = callframe.New()
	callframe.Arg(frame, b)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TouchScreenButton.Bind_set_shape_centered, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShapeCentered() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TouchScreenButton.Bind_is_shape_centered, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShapeVisible(b bool) {
	var frame = callframe.New()
	callframe.Arg(frame, b)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TouchScreenButton.Bind_set_shape_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShapeVisible() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TouchScreenButton.Bind_is_shape_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAction(action gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TouchScreenButton.Bind_set_action, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAction() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TouchScreenButton.Bind_get_action, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibilityMode(mode gdclass.TouchScreenButtonVisibilityMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TouchScreenButton.Bind_set_visibility_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityMode() gdclass.TouchScreenButtonVisibilityMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TouchScreenButtonVisibilityMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TouchScreenButton.Bind_get_visibility_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPassbyPress(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TouchScreenButton.Bind_set_passby_press, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPassbyPressEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TouchScreenButton.Bind_is_passby_press_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this button is currently pressed.
*/
//go:nosplit
func (self class) IsPressed() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TouchScreenButton.Bind_is_pressed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnPressed(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("pressed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnReleased(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("released"), gd.NewCallable(cb), 0)
}

func (self class) AsTouchScreenButton() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTouchScreenButton() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced        { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance     { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Advanced(self.AsNode2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Instance(self.AsNode2D()), name)
	}
}
func init() {
	gdclass.Register("TouchScreenButton", func(ptr gd.Object) any {
		return [1]gdclass.TouchScreenButton{*(*gdclass.TouchScreenButton)(unsafe.Pointer(&ptr))}
	})
}

type VisibilityMode = gdclass.TouchScreenButtonVisibilityMode

const (
	/*Always visible.*/
	VisibilityAlways VisibilityMode = 0
	/*Visible on touch screens only.*/
	VisibilityTouchscreenOnly VisibilityMode = 1
)
