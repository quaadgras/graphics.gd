package TextLine

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Abstraction over [TextServer] for handling a single line of text.

*/
type Go [1]classdb.TextLine

/*
Clears text line (removes text and inline objects).
*/
func (self Go) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Clear()
}

/*
Overrides BiDi for the structured text.
Override ranges should cover full source text without overlaps. BiDi algorithm will be used on each range separately.
*/
func (self Go) SetBidiOverride(override gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBidiOverride(override)
}

/*
Adds text span and font to draw it.
*/
func (self Go) AddString(text string, font gdclass.Font, font_size int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).AddString(gc.String(text), font, gd.Int(font_size), gc.String(""), gc.Variant(([1]gd.Variant{}[0]))))
}

/*
Adds inline object to the text buffer, [param key] must be unique. In the text, object is represented as [param length] object replacement characters.
*/
func (self Go) AddObject(key gd.Variant, size gd.Vector2) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).AddObject(key, size, 5, gd.Int(1), gd.Float(0.0)))
}

/*
Sets new size and alignment of embedded object.
*/
func (self Go) ResizeObject(key gd.Variant, size gd.Vector2) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).ResizeObject(key, size, 5, gd.Float(0.0)))
}

/*
Aligns text to the given tab-stops.
*/
func (self Go) TabAlign(tab_stops []float32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).TabAlign(gc.PackedFloat32Slice(tab_stops))
}

/*
Returns array of inline objects.
*/
func (self Go) GetObjects() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(class(self).GetObjects(gc))
}

/*
Returns bounding rectangle of the inline object.
*/
func (self Go) GetObjectRect(key gd.Variant) gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(class(self).GetObjectRect(key))
}

/*
Returns size of the bounding box of the text.
*/
func (self Go) GetSize() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetSize())
}

/*
Returns TextServer buffer RID.
*/
func (self Go) GetRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetRid())
}

/*
Returns the text ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
*/
func (self Go) GetLineAscent() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetLineAscent()))
}

/*
Returns the text descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
*/
func (self Go) GetLineDescent() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetLineDescent()))
}

/*
Returns width (for horizontal layout) or height (for vertical) of the text.
*/
func (self Go) GetLineWidth() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetLineWidth()))
}

/*
Returns pixel offset of the underline below the baseline.
*/
func (self Go) GetLineUnderlinePosition() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetLineUnderlinePosition()))
}

/*
Returns thickness of the underline.
*/
func (self Go) GetLineUnderlineThickness() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetLineUnderlineThickness()))
}

/*
Draw text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Go) Draw(canvas gd.RID, pos gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Draw(canvas, pos, gd.Color{1, 1, 1, 1})
}

/*
Draw text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Go) DrawOutline(canvas gd.RID, pos gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawOutline(canvas, pos, gd.Int(1), gd.Color{1, 1, 1, 1})
}

/*
Returns caret character offset at the specified pixel offset at the baseline. This function always returns a valid position.
*/
func (self Go) HitTest(coords float64) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).HitTest(gd.Float(coords))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TextLine
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("TextLine"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Direction() classdb.TextServerDirection {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextServerDirection(class(self).GetDirection())
}

func (self Go) SetDirection(value classdb.TextServerDirection) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDirection(value)
}

func (self Go) Orientation() classdb.TextServerOrientation {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextServerOrientation(class(self).GetOrientation())
}

func (self Go) SetOrientation(value classdb.TextServerOrientation) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOrientation(value)
}

func (self Go) PreserveInvalid() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetPreserveInvalid())
}

func (self Go) SetPreserveInvalid(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPreserveInvalid(value)
}

func (self Go) PreserveControl() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetPreserveControl())
}

func (self Go) SetPreserveControl(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPreserveControl(value)
}

func (self Go) Width() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetWidth()))
}

func (self Go) SetWidth(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetWidth(gd.Float(value))
}

func (self Go) Alignment() gd.HorizontalAlignment {
	gc := gd.GarbageCollector(); _ = gc
		return gd.HorizontalAlignment(class(self).GetHorizontalAlignment())
}

func (self Go) SetAlignment(value gd.HorizontalAlignment) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHorizontalAlignment(value)
}

func (self Go) Flags() classdb.TextServerJustificationFlag {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextServerJustificationFlag(class(self).GetFlags())
}

func (self Go) SetFlags(value classdb.TextServerJustificationFlag) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFlags(value)
}

func (self Go) TextOverrunBehavior() classdb.TextServerOverrunBehavior {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextServerOverrunBehavior(class(self).GetTextOverrunBehavior())
}

func (self Go) SetTextOverrunBehavior(value classdb.TextServerOverrunBehavior) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextOverrunBehavior(value)
}

func (self Go) EllipsisChar() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetEllipsisChar(gc).String())
}

func (self Go) SetEllipsisChar(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEllipsisChar(gc.String(value))
}

/*
Clears text line (removes text and inline objects).
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetDirection(direction classdb.TextServerDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_set_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDirection() classdb.TextServerDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_get_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOrientation(orientation classdb.TextServerOrientation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, orientation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_set_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOrientation() classdb.TextServerOrientation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerOrientation](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_get_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPreserveInvalid(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_set_preserve_invalid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPreserveInvalid() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_get_preserve_invalid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPreserveControl(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_set_preserve_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPreserveControl() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_get_preserve_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Overrides BiDi for the structured text.
Override ranges should cover full source text without overlaps. BiDi algorithm will be used on each range separately.
*/
//go:nosplit
func (self class) SetBidiOverride(override gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(override))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_set_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds text span and font to draw it.
*/
//go:nosplit
func (self class) AddString(text gd.String, font gdclass.Font, font_size gd.Int, language gd.String, meta gd.Variant) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, mmm.Get(language))
	callframe.Arg(frame, mmm.Get(meta))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_add_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds inline object to the text buffer, [param key] must be unique. In the text, object is represented as [param length] object replacement characters.
*/
//go:nosplit
func (self class) AddObject(key gd.Variant, size gd.Vector2, inline_align gd.InlineAlignment, length gd.Int, baseline gd.Float) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(key))
	callframe.Arg(frame, size)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, length)
	callframe.Arg(frame, baseline)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_add_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets new size and alignment of embedded object.
*/
//go:nosplit
func (self class) ResizeObject(key gd.Variant, size gd.Vector2, inline_align gd.InlineAlignment, baseline gd.Float) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(key))
	callframe.Arg(frame, size)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, baseline)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_resize_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWidth(width gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWidth() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHorizontalAlignment(alignment gd.HorizontalAlignment)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_set_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHorizontalAlignment() gd.HorizontalAlignment {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.HorizontalAlignment](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_get_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Aligns text to the given tab-stops.
*/
//go:nosplit
func (self class) TabAlign(tab_stops gd.PackedFloat32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(tab_stops))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_tab_align, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFlags(flags classdb.TextServerJustificationFlag)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_set_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFlags() classdb.TextServerJustificationFlag {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerJustificationFlag](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_get_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextOverrunBehavior(overrun_behavior classdb.TextServerOverrunBehavior)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, overrun_behavior)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_set_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextOverrunBehavior() classdb.TextServerOverrunBehavior {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerOverrunBehavior](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_get_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEllipsisChar(char gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(char))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_set_ellipsis_char, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEllipsisChar(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_get_ellipsis_char, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns array of inline objects.
*/
//go:nosplit
func (self class) GetObjects(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_get_objects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns bounding rectangle of the inline object.
*/
//go:nosplit
func (self class) GetObjectRect(key gd.Variant) gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(key))
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_get_object_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns size of the bounding box of the text.
*/
//go:nosplit
func (self class) GetSize() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns TextServer buffer RID.
*/
//go:nosplit
func (self class) GetRid() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the text ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
*/
//go:nosplit
func (self class) GetLineAscent() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_get_line_ascent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the text descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
*/
//go:nosplit
func (self class) GetLineDescent() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_get_line_descent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns width (for horizontal layout) or height (for vertical) of the text.
*/
//go:nosplit
func (self class) GetLineWidth() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_get_line_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) GetLineUnderlinePosition() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_get_line_underline_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns thickness of the underline.
*/
//go:nosplit
func (self class) GetLineUnderlineThickness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_get_line_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Draw text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) Draw(canvas gd.RID, pos gd.Vector2, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draw text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) DrawOutline(canvas gd.RID, pos gd.Vector2, outline_size gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, outline_size)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_draw_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns caret character offset at the specified pixel offset at the baseline. This function always returns a valid position.
*/
//go:nosplit
func (self class) HitTest(coords gd.Float) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextLine.Bind_hit_test, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTextLine() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTextLine() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("TextLine", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
