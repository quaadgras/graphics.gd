// Package StyleBox provides methods for working with StyleBox object instances.
package StyleBox

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
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Rect2"
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
[StyleBox] is an abstract base class for drawing stylized boxes for UI elements. It is used for panels, buttons, [LineEdit] backgrounds, [Tree] backgrounds, etc. and also for testing a transparency mask for pointer signals. If mask test fails on a [StyleBox] assigned as mask to a control, clicks and motion signals will go through it to the one below.
[b]Note:[/b] For control nodes that have [i]Theme Properties[/i], the [code]focus[/code] [StyleBox] is displayed over the [code]normal[/code], [code]hover[/code] or [code]pressed[/code] [StyleBox]. This makes the [code]focus[/code] [StyleBox] more reusable across different nodes.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=StyleBox)
*/
type Instance [1]gdclass.StyleBox

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsStyleBox() Instance
}
type Interface interface {
	Draw(to_canvas_item RID.Any, rect Rect2.PositionSize)
	GetDrawRect(rect Rect2.PositionSize) Rect2.PositionSize
	//Virtual method to be implemented by the user. Returns a custom minimum size that the stylebox must respect when drawing. By default [method get_minimum_size] only takes content margins into account. This method can be overridden to add another size restriction. A combination of the default behavior and the output of this method will be used, to account for both sizes.
	GetMinimumSize() Vector2.XY
	TestMask(point Vector2.XY, rect Rect2.PositionSize) bool
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) Draw(to_canvas_item RID.Any, rect Rect2.PositionSize)        { return }
func (self implementation) GetDrawRect(rect Rect2.PositionSize) (_ Rect2.PositionSize)  { return }
func (self implementation) GetMinimumSize() (_ Vector2.XY)                              { return }
func (self implementation) TestMask(point Vector2.XY, rect Rect2.PositionSize) (_ bool) { return }
func (Instance) _draw(impl func(ptr unsafe.Pointer, to_canvas_item RID.Any, rect Rect2.PositionSize)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var to_canvas_item = gd.UnsafeGet[gd.RID](p_args, 0)

		var rect = gd.UnsafeGet[gd.Rect2](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, to_canvas_item, rect)
	}
}
func (Instance) _get_draw_rect(impl func(ptr unsafe.Pointer, rect Rect2.PositionSize) Rect2.PositionSize) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var rect = gd.UnsafeGet[gd.Rect2](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, rect)
		gd.UnsafeSet(p_back, gd.Rect2(ret))
	}
}

/*
Virtual method to be implemented by the user. Returns a custom minimum size that the stylebox must respect when drawing. By default [method get_minimum_size] only takes content margins into account. This method can be overridden to add another size restriction. A combination of the default behavior and the output of this method will be used, to account for both sizes.
*/
func (Instance) _get_minimum_size(impl func(ptr unsafe.Pointer) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}
func (Instance) _test_mask(impl func(ptr unsafe.Pointer, point Vector2.XY, rect Rect2.PositionSize) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
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
func (self Instance) GetMinimumSize() Vector2.XY { //gd:StyleBox.get_minimum_size
	return Vector2.XY(class(self).GetMinimumSize())
}

/*
Sets the default margin to [param offset] pixels for all sides.
*/
func (self Instance) SetContentMarginAll(offset Float.X) { //gd:StyleBox.set_content_margin_all
	class(self).SetContentMarginAll(gd.Float(offset))
}

/*
Returns the content margin offset for the specified [enum Side].
Positive values reduce size inwards, unlike [Control]'s margin values.
*/
func (self Instance) GetMargin(margin Side) Float.X { //gd:StyleBox.get_margin
	return Float.X(Float.X(class(self).GetMargin(margin)))
}

/*
Returns the "offset" of a stylebox. This helper function returns a value equivalent to [code]Vector2(style.get_margin(MARGIN_LEFT), style.get_margin(MARGIN_TOP))[/code].
*/
func (self Instance) GetOffset() Vector2.XY { //gd:StyleBox.get_offset
	return Vector2.XY(class(self).GetOffset())
}

/*
Draws this stylebox using a canvas item identified by the given [RID].
The [RID] value can either be the result of [method CanvasItem.get_canvas_item] called on an existing [CanvasItem]-derived node, or directly from creating a canvas item in the [RenderingServer] with [method RenderingServer.canvas_item_create].
*/
func (self Instance) Draw(canvas_item RID.CanvasItem, rect Rect2.PositionSize) { //gd:StyleBox.draw
	class(self).Draw(gd.RID(canvas_item), gd.Rect2(rect))
}

/*
Returns the [CanvasItem] that handles its [constant CanvasItem.NOTIFICATION_DRAW] or [method CanvasItem._draw] callback at this moment.
*/
func (self Instance) GetCurrentItemDrawn() [1]gdclass.CanvasItem { //gd:StyleBox.get_current_item_drawn
	return [1]gdclass.CanvasItem(class(self).GetCurrentItemDrawn())
}

/*
Test a position in a rectangle, return whether it passes the mask test.
*/
func (self Instance) TestMask(point Vector2.XY, rect Rect2.PositionSize) bool { //gd:StyleBox.test_mask
	return bool(class(self).TestMask(gd.Vector2(point), gd.Rect2(rect)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.StyleBox

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StyleBox"))
	casted := Instance{*(*gdclass.StyleBox)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) ContentMarginLeft() Float.X {
	return Float.X(Float.X(class(self).GetContentMargin(0)))
}

func (self Instance) SetContentMarginLeft(value Float.X) {
	class(self).SetContentMargin(0, gd.Float(value))
}

func (self Instance) ContentMarginTop() Float.X {
	return Float.X(Float.X(class(self).GetContentMargin(1)))
}

func (self Instance) SetContentMarginTop(value Float.X) {
	class(self).SetContentMargin(1, gd.Float(value))
}

func (self Instance) ContentMarginRight() Float.X {
	return Float.X(Float.X(class(self).GetContentMargin(2)))
}

func (self Instance) SetContentMarginRight(value Float.X) {
	class(self).SetContentMargin(2, gd.Float(value))
}

func (self Instance) ContentMarginBottom() Float.X {
	return Float.X(Float.X(class(self).GetContentMargin(3)))
}

func (self Instance) SetContentMarginBottom(value Float.X) {
	class(self).SetContentMargin(3, gd.Float(value))
}

func (class) _draw(impl func(ptr unsafe.Pointer, to_canvas_item gd.RID, rect gd.Rect2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var to_canvas_item = gd.UnsafeGet[gd.RID](p_args, 0)

		var rect = gd.UnsafeGet[gd.Rect2](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, to_canvas_item, rect)
	}
}

func (class) _get_draw_rect(impl func(ptr unsafe.Pointer, rect gd.Rect2) gd.Rect2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
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
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _test_mask(impl func(ptr unsafe.Pointer, point gd.Vector2, rect gd.Rect2) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
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
func (self class) GetMinimumSize() gd.Vector2 { //gd:StyleBox.get_minimum_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBox.Bind_get_minimum_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the default value of the specified [enum Side] to [param offset] pixels.
*/
//go:nosplit
func (self class) SetContentMargin(margin Side, offset gd.Float) { //gd:StyleBox.set_content_margin
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBox.Bind_set_content_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the default margin to [param offset] pixels for all sides.
*/
//go:nosplit
func (self class) SetContentMarginAll(offset gd.Float) { //gd:StyleBox.set_content_margin_all
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBox.Bind_set_content_margin_all, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the default margin of the specified [enum Side].
*/
//go:nosplit
func (self class) GetContentMargin(margin Side) gd.Float { //gd:StyleBox.get_content_margin
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBox.Bind_get_content_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the content margin offset for the specified [enum Side].
Positive values reduce size inwards, unlike [Control]'s margin values.
*/
//go:nosplit
func (self class) GetMargin(margin Side) gd.Float { //gd:StyleBox.get_margin
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBox.Bind_get_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the "offset" of a stylebox. This helper function returns a value equivalent to [code]Vector2(style.get_margin(MARGIN_LEFT), style.get_margin(MARGIN_TOP))[/code].
*/
//go:nosplit
func (self class) GetOffset() gd.Vector2 { //gd:StyleBox.get_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBox.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Draws this stylebox using a canvas item identified by the given [RID].
The [RID] value can either be the result of [method CanvasItem.get_canvas_item] called on an existing [CanvasItem]-derived node, or directly from creating a canvas item in the [RenderingServer] with [method RenderingServer.canvas_item_create].
*/
//go:nosplit
func (self class) Draw(canvas_item gd.RID, rect gd.Rect2) { //gd:StyleBox.draw
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, rect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBox.Bind_draw, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [CanvasItem] that handles its [constant CanvasItem.NOTIFICATION_DRAW] or [method CanvasItem._draw] callback at this moment.
*/
//go:nosplit
func (self class) GetCurrentItemDrawn() [1]gdclass.CanvasItem { //gd:StyleBox.get_current_item_drawn
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBox.Bind_get_current_item_drawn, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.CanvasItem{gd.PointerMustAssertInstanceID[gdclass.CanvasItem](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Test a position in a rectangle, return whether it passes the mask test.
*/
//go:nosplit
func (self class) TestMask(point gd.Vector2, rect gd.Rect2) bool { //gd:StyleBox.test_mask
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, rect)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBox.Bind_test_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
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
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("StyleBox", func(ptr gd.Object) any { return [1]gdclass.StyleBox{*(*gdclass.StyleBox)(unsafe.Pointer(&ptr))} })
}

type Side int

const (
	/*Left side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideLeft Side = 0
	/*Top side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideTop Side = 1
	/*Right side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideRight Side = 2
	/*Bottom side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideBottom Side = 3
)
