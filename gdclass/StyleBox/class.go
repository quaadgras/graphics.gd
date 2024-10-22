package StyleBox

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Go [1]classdb.StyleBox
func (Go) _draw(impl func(ptr unsafe.Pointer, to_canvas_item gd.RID, rect gd.Rect2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var to_canvas_item = gd.UnsafeGet[gd.RID](p_args,0)
		var rect = gd.UnsafeGet[gd.Rect2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, to_canvas_item, rect)
		gc.End()
	}
}
func (Go) _get_draw_rect(impl func(ptr unsafe.Pointer, rect gd.Rect2) gd.Rect2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var rect = gd.UnsafeGet[gd.Rect2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, rect)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Virtual method to be implemented by the user. Returns a custom minimum size that the stylebox must respect when drawing. By default [method get_minimum_size] only takes content margins into account. This method can be overridden to add another size restriction. A combination of the default behavior and the output of this method will be used, to account for both sizes.
*/
func (Go) _get_minimum_size(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _test_mask(impl func(ptr unsafe.Pointer, point gd.Vector2, rect gd.Rect2) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var point = gd.UnsafeGet[gd.Vector2](p_args,0)
		var rect = gd.UnsafeGet[gd.Rect2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, point, rect)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Returns the minimum size that this stylebox can be shrunk to.
*/
func (self Go) GetMinimumSize() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetMinimumSize())
}

/*
Sets the default margin to [param offset] pixels for all sides.
*/
func (self Go) SetContentMarginAll(offset float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetContentMarginAll(gd.Float(offset))
}

/*
Returns the content margin offset for the specified [enum Side].
Positive values reduce size inwards, unlike [Control]'s margin values.
*/
func (self Go) GetMargin(margin gd.Side) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetMargin(margin)))
}

/*
Returns the "offset" of a stylebox. This helper function returns a value equivalent to [code]Vector2(style.get_margin(MARGIN_LEFT), style.get_margin(MARGIN_TOP))[/code].
*/
func (self Go) GetOffset() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetOffset())
}

/*
Draws this stylebox using a canvas item identified by the given [RID].
The [RID] value can either be the result of [method CanvasItem.get_canvas_item] called on an existing [CanvasItem]-derived node, or directly from creating a canvas item in the [RenderingServer] with [method RenderingServer.canvas_item_create].
*/
func (self Go) Draw(canvas_item gd.RID, rect gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Draw(canvas_item, rect)
}

/*
Returns the [CanvasItem] that handles its [constant CanvasItem.NOTIFICATION_DRAW] or [method CanvasItem._draw] callback at this moment.
*/
func (self Go) GetCurrentItemDrawn() gdclass.CanvasItem {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.CanvasItem(class(self).GetCurrentItemDrawn(gc))
}

/*
Test a position in a rectangle, return whether it passes the mask test.
*/
func (self Go) TestMask(point gd.Vector2, rect gd.Rect2) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).TestMask(point, rect))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.StyleBox
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("StyleBox"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) ContentMarginLeft() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetContentMargin(0)))
}

func (self Go) SetContentMarginLeft(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetContentMargin(0, gd.Float(value))
}

func (self Go) ContentMarginTop() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetContentMargin(1)))
}

func (self Go) SetContentMarginTop(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetContentMargin(1, gd.Float(value))
}

func (self Go) ContentMarginRight() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetContentMargin(2)))
}

func (self Go) SetContentMarginRight(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetContentMargin(2, gd.Float(value))
}

func (self Go) ContentMarginBottom() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetContentMargin(3)))
}

func (self Go) SetContentMarginBottom(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetContentMargin(3, gd.Float(value))
}

func (class) _draw(impl func(ptr unsafe.Pointer, to_canvas_item gd.RID, rect gd.Rect2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var to_canvas_item = gd.UnsafeGet[gd.RID](p_args,0)
		var rect = gd.UnsafeGet[gd.Rect2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, to_canvas_item, rect)
		ctx.End()
	}
}

func (class) _get_draw_rect(impl func(ptr unsafe.Pointer, rect gd.Rect2) gd.Rect2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var rect = gd.UnsafeGet[gd.Rect2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, rect)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Virtual method to be implemented by the user. Returns a custom minimum size that the stylebox must respect when drawing. By default [method get_minimum_size] only takes content margins into account. This method can be overridden to add another size restriction. A combination of the default behavior and the output of this method will be used, to account for both sizes.
*/
func (class) _get_minimum_size(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _test_mask(impl func(ptr unsafe.Pointer, point gd.Vector2, rect gd.Rect2) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var point = gd.UnsafeGet[gd.Vector2](p_args,0)
		var rect = gd.UnsafeGet[gd.Rect2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, point, rect)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Returns the minimum size that this stylebox can be shrunk to.
*/
//go:nosplit
func (self class) GetMinimumSize() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBox.Bind_get_minimum_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the default value of the specified [enum Side] to [param offset] pixels.
*/
//go:nosplit
func (self class) SetContentMargin(margin gd.Side, offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBox.Bind_set_content_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the default margin to [param offset] pixels for all sides.
*/
//go:nosplit
func (self class) SetContentMarginAll(offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBox.Bind_set_content_margin_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the default margin of the specified [enum Side].
*/
//go:nosplit
func (self class) GetContentMargin(margin gd.Side) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBox.Bind_get_content_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the content margin offset for the specified [enum Side].
Positive values reduce size inwards, unlike [Control]'s margin values.
*/
//go:nosplit
func (self class) GetMargin(margin gd.Side) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBox.Bind_get_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the "offset" of a stylebox. This helper function returns a value equivalent to [code]Vector2(style.get_margin(MARGIN_LEFT), style.get_margin(MARGIN_TOP))[/code].
*/
//go:nosplit
func (self class) GetOffset() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBox.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Draws this stylebox using a canvas item identified by the given [RID].
The [RID] value can either be the result of [method CanvasItem.get_canvas_item] called on an existing [CanvasItem]-derived node, or directly from creating a canvas item in the [RenderingServer] with [method RenderingServer.canvas_item_create].
*/
//go:nosplit
func (self class) Draw(canvas_item gd.RID, rect gd.Rect2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBox.Bind_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [CanvasItem] that handles its [constant CanvasItem.NOTIFICATION_DRAW] or [method CanvasItem._draw] callback at this moment.
*/
//go:nosplit
func (self class) GetCurrentItemDrawn(ctx gd.Lifetime) gdclass.CanvasItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBox.Bind_get_current_item_drawn, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.CanvasItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Test a position in a rectangle, return whether it passes the mask test.
*/
//go:nosplit
func (self class) TestMask(point gd.Vector2, rect gd.Rect2) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, rect)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBox.Bind_test_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsStyleBox() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsStyleBox() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_draw": return reflect.ValueOf(self._draw);
	case "_get_draw_rect": return reflect.ValueOf(self._get_draw_rect);
	case "_get_minimum_size": return reflect.ValueOf(self._get_minimum_size);
	case "_test_mask": return reflect.ValueOf(self._test_mask);
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_draw": return reflect.ValueOf(self._draw);
	case "_get_draw_rect": return reflect.ValueOf(self._get_draw_rect);
	case "_get_minimum_size": return reflect.ValueOf(self._get_minimum_size);
	case "_test_mask": return reflect.ValueOf(self._test_mask);
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("StyleBox", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
