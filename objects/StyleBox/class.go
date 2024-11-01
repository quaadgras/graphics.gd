package StyleBox

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
[StyleBox] is an abstract base class for drawing stylized boxes for UI elements. It is used for panels, buttons, [LineEdit] backgrounds, [Tree] backgrounds, etc. and also for testing a transparency mask for pointer signals. If mask test fails on a [StyleBox] assigned as mask to a control, clicks and motion signals will go through it to the one below.
[b]Note:[/b] For control nodes that have [i]Theme Properties[/i], the [code]focus[/code] [StyleBox] is displayed over the [code]normal[/code], [code]hover[/code] or [code]pressed[/code] [StyleBox]. This makes the [code]focus[/code] [StyleBox] more reusable across different nodes.

	// StyleBox methods that can be overridden by a [Class] that extends it.
	type StyleBox interface {
		Draw(to_canvas_item gd.RID, rect gd.Rect2)
		GetDrawRect(rect gd.Rect2) gd.Rect2
		//Virtual method to be implemented by the user. Returns a custom minimum size that the stylebox must respect when drawing. By default [method get_minimum_size] only takes content margins into account. This method can be overridden to add another size restriction. A combination of the default behavior and the output of this method will be used, to account for both sizes.
		GetMinimumSize() gd.Vector2
		TestMask(point gd.Vector2, rect gd.Rect2) bool
	}
*/
type Instance [1]classdb.StyleBox

func (Instance) _draw(impl func(ptr unsafe.Pointer, to_canvas_item gd.RID, rect gd.Rect2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var to_canvas_item = gd.UnsafeGet[gd.RID](p_args, 0)
		var rect = gd.UnsafeGet[gd.Rect2](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, to_canvas_item, rect)
	}
}
func (Instance) _get_draw_rect(impl func(ptr unsafe.Pointer, rect gd.Rect2) gd.Rect2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var rect = gd.UnsafeGet[gd.Rect2](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, rect)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method to be implemented by the user. Returns a custom minimum size that the stylebox must respect when drawing. By default [method get_minimum_size] only takes content margins into account. This method can be overridden to add another size restriction. A combination of the default behavior and the output of this method will be used, to account for both sizes.
*/
func (Instance) _get_minimum_size(impl func(ptr unsafe.Pointer) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _test_mask(impl func(ptr unsafe.Pointer, point gd.Vector2, rect gd.Rect2) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var point = gd.UnsafeGet[gd.Vector2](p_args, 0)
		var rect = gd.UnsafeGet[gd.Rect2](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, point, rect)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the minimum size that this stylebox can be shrunk to.
*/
func (self Instance) GetMinimumSize() gd.Vector2 {
	return gd.Vector2(class(self).GetMinimumSize())
}

/*
Sets the default margin to [param offset] pixels for all sides.
*/
func (self Instance) SetContentMarginAll(offset float64) {
	class(self).SetContentMarginAll(gd.Float(offset))
}

/*
Returns the content margin offset for the specified [enum Side].
Positive values reduce size inwards, unlike [Control]'s margin values.
*/
func (self Instance) GetMargin(margin gdconst.Side) float64 {
	return float64(float64(class(self).GetMargin(margin)))
}

/*
Returns the "offset" of a stylebox. This helper function returns a value equivalent to [code]Vector2(style.get_margin(MARGIN_LEFT), style.get_margin(MARGIN_TOP))[/code].
*/
func (self Instance) GetOffset() gd.Vector2 {
	return gd.Vector2(class(self).GetOffset())
}

/*
Draws this stylebox using a canvas item identified by the given [RID].
The [RID] value can either be the result of [method CanvasItem.get_canvas_item] called on an existing [CanvasItem]-derived node, or directly from creating a canvas item in the [RenderingServer] with [method RenderingServer.canvas_item_create].
*/
func (self Instance) Draw(canvas_item gd.RID, rect gd.Rect2) {
	class(self).Draw(canvas_item, rect)
}

/*
Returns the [CanvasItem] that handles its [constant CanvasItem.NOTIFICATION_DRAW] or [method CanvasItem._draw] callback at this moment.
*/
func (self Instance) GetCurrentItemDrawn() objects.CanvasItem {
	return objects.CanvasItem(class(self).GetCurrentItemDrawn())
}

/*
Test a position in a rectangle, return whether it passes the mask test.
*/
func (self Instance) TestMask(point gd.Vector2, rect gd.Rect2) bool {
	return bool(class(self).TestMask(point, rect))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.StyleBox

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StyleBox"))
	return Instance{classdb.StyleBox(object)}
}

func (self Instance) ContentMarginLeft() float64 {
	return float64(float64(class(self).GetContentMargin(0)))
}

func (self Instance) SetContentMarginLeft(value float64) {
	class(self).SetContentMargin(0, gd.Float(value))
}

func (self Instance) ContentMarginTop() float64 {
	return float64(float64(class(self).GetContentMargin(1)))
}

func (self Instance) SetContentMarginTop(value float64) {
	class(self).SetContentMargin(1, gd.Float(value))
}

func (self Instance) ContentMarginRight() float64 {
	return float64(float64(class(self).GetContentMargin(2)))
}

func (self Instance) SetContentMarginRight(value float64) {
	class(self).SetContentMargin(2, gd.Float(value))
}

func (self Instance) ContentMarginBottom() float64 {
	return float64(float64(class(self).GetContentMargin(3)))
}

func (self Instance) SetContentMarginBottom(value float64) {
	class(self).SetContentMargin(3, gd.Float(value))
}

func (class) _draw(impl func(ptr unsafe.Pointer, to_canvas_item gd.RID, rect gd.Rect2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var to_canvas_item = gd.UnsafeGet[gd.RID](p_args, 0)
		var rect = gd.UnsafeGet[gd.Rect2](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, to_canvas_item, rect)
	}
}

func (class) _get_draw_rect(impl func(ptr unsafe.Pointer, rect gd.Rect2) gd.Rect2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var rect = gd.UnsafeGet[gd.Rect2](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, rect)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method to be implemented by the user. Returns a custom minimum size that the stylebox must respect when drawing. By default [method get_minimum_size] only takes content margins into account. This method can be overridden to add another size restriction. A combination of the default behavior and the output of this method will be used, to account for both sizes.
*/
func (class) _get_minimum_size(impl func(ptr unsafe.Pointer) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _test_mask(impl func(ptr unsafe.Pointer, point gd.Vector2, rect gd.Rect2) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var point = gd.UnsafeGet[gd.Vector2](p_args, 0)
		var rect = gd.UnsafeGet[gd.Rect2](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, point, rect)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the minimum size that this stylebox can be shrunk to.
*/
//go:nosplit
func (self class) GetMinimumSize() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBox.Bind_get_minimum_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the default value of the specified [enum Side] to [param offset] pixels.
*/
//go:nosplit
func (self class) SetContentMargin(margin gdconst.Side, offset gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBox.Bind_set_content_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the default margin to [param offset] pixels for all sides.
*/
//go:nosplit
func (self class) SetContentMarginAll(offset gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBox.Bind_set_content_margin_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the default margin of the specified [enum Side].
*/
//go:nosplit
func (self class) GetContentMargin(margin gdconst.Side) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBox.Bind_get_content_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the content margin offset for the specified [enum Side].
Positive values reduce size inwards, unlike [Control]'s margin values.
*/
//go:nosplit
func (self class) GetMargin(margin gdconst.Side) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBox.Bind_get_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the "offset" of a stylebox. This helper function returns a value equivalent to [code]Vector2(style.get_margin(MARGIN_LEFT), style.get_margin(MARGIN_TOP))[/code].
*/
//go:nosplit
func (self class) GetOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBox.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Draws this stylebox using a canvas item identified by the given [RID].
The [RID] value can either be the result of [method CanvasItem.get_canvas_item] called on an existing [CanvasItem]-derived node, or directly from creating a canvas item in the [RenderingServer] with [method RenderingServer.canvas_item_create].
*/
//go:nosplit
func (self class) Draw(canvas_item gd.RID, rect gd.Rect2) {
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBox.Bind_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [CanvasItem] that handles its [constant CanvasItem.NOTIFICATION_DRAW] or [method CanvasItem._draw] callback at this moment.
*/
//go:nosplit
func (self class) GetCurrentItemDrawn() objects.CanvasItem {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBox.Bind_get_current_item_drawn, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.CanvasItem{classdb.CanvasItem(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Test a position in a rectangle, return whether it passes the mask test.
*/
//go:nosplit
func (self class) TestMask(point gd.Vector2, rect gd.Rect2) bool {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, rect)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBox.Bind_test_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsStyleBox() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStyleBox() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_draw":
		return reflect.ValueOf(self._draw)
	case "_get_draw_rect":
		return reflect.ValueOf(self._get_draw_rect)
	case "_get_minimum_size":
		return reflect.ValueOf(self._get_minimum_size)
	case "_test_mask":
		return reflect.ValueOf(self._test_mask)
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_draw":
		return reflect.ValueOf(self._draw)
	case "_get_draw_rect":
		return reflect.ValueOf(self._get_draw_rect)
	case "_get_minimum_size":
		return reflect.ValueOf(self._get_minimum_size)
	case "_test_mask":
		return reflect.ValueOf(self._test_mask)
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() { classdb.Register("StyleBox", func(ptr gd.Object) any { return classdb.StyleBox(ptr) }) }
