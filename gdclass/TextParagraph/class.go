package TextParagraph

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
Abstraction over [TextServer] for handling a single paragraph of text.

*/
type Go [1]classdb.TextParagraph

/*
Clears text paragraph (removes text and inline objects).
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
Sets drop cap, overrides previously set drop cap. Drop cap (dropped capital) is a decorative element at the beginning of a paragraph that is larger than the rest of the text.
*/
func (self Go) SetDropcap(text string, font gdclass.Font, font_size int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).SetDropcap(gc.String(text), font, gd.Int(font_size), gd.NewRect2(0, 0, 0, 0), gc.String("")))
}

/*
Removes dropcap.
*/
func (self Go) ClearDropcap() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearDropcap()
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
Aligns paragraph to the given tab-stops.
*/
func (self Go) TabAlign(tab_stops []float32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).TabAlign(gc.PackedFloat32Slice(tab_stops))
}

/*
Returns the size of the bounding box of the paragraph, without line breaks.
*/
func (self Go) GetNonWrappedSize() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetNonWrappedSize())
}

/*
Returns the size of the bounding box of the paragraph.
*/
func (self Go) GetSize() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetSize())
}

/*
Returns TextServer full string buffer RID.
*/
func (self Go) GetRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetRid())
}

/*
Returns TextServer line buffer RID.
*/
func (self Go) GetLineRid(line int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetLineRid(gd.Int(line)))
}

/*
Returns drop cap text buffer RID.
*/
func (self Go) GetDropcapRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetDropcapRid())
}

/*
Returns number of lines in the paragraph.
*/
func (self Go) GetLineCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetLineCount()))
}

/*
Returns array of inline objects in the line.
*/
func (self Go) GetLineObjects(line int) gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(class(self).GetLineObjects(gc, gd.Int(line)))
}

/*
Returns bounding rectangle of the inline object.
*/
func (self Go) GetLineObjectRect(line int, key gd.Variant) gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(class(self).GetLineObjectRect(gd.Int(line), key))
}

/*
Returns size of the bounding box of the line of text. Returned size is rounded up.
*/
func (self Go) GetLineSize(line int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetLineSize(gd.Int(line)))
}

/*
Returns character range of the line.
*/
func (self Go) GetLineRange(line int) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(class(self).GetLineRange(gd.Int(line)))
}

/*
Returns the text line ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
*/
func (self Go) GetLineAscent(line int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetLineAscent(gd.Int(line))))
}

/*
Returns the text line descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
*/
func (self Go) GetLineDescent(line int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetLineDescent(gd.Int(line))))
}

/*
Returns width (for horizontal layout) or height (for vertical) of the line of text.
*/
func (self Go) GetLineWidth(line int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetLineWidth(gd.Int(line))))
}

/*
Returns pixel offset of the underline below the baseline.
*/
func (self Go) GetLineUnderlinePosition(line int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetLineUnderlinePosition(gd.Int(line))))
}

/*
Returns thickness of the underline.
*/
func (self Go) GetLineUnderlineThickness(line int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetLineUnderlineThickness(gd.Int(line))))
}

/*
Returns drop cap bounding box size.
*/
func (self Go) GetDropcapSize() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetDropcapSize())
}

/*
Returns number of lines used by dropcap.
*/
func (self Go) GetDropcapLines() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetDropcapLines()))
}

/*
Draw all lines of the text and drop cap into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Go) Draw(canvas gd.RID, pos gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Draw(canvas, pos, gd.Color{1, 1, 1, 1}, gd.Color{1, 1, 1, 1})
}

/*
Draw outlines of all lines of the text and drop cap into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Go) DrawOutline(canvas gd.RID, pos gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawOutline(canvas, pos, gd.Int(1), gd.Color{1, 1, 1, 1}, gd.Color{1, 1, 1, 1})
}

/*
Draw single line of text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Go) DrawLine(canvas gd.RID, pos gd.Vector2, line int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawLine(canvas, pos, gd.Int(line), gd.Color{1, 1, 1, 1})
}

/*
Draw outline of the single line of text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Go) DrawLineOutline(canvas gd.RID, pos gd.Vector2, line int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawLineOutline(canvas, pos, gd.Int(line), gd.Int(1), gd.Color{1, 1, 1, 1})
}

/*
Draw drop cap into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Go) DrawDropcap(canvas gd.RID, pos gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawDropcap(canvas, pos, gd.Color{1, 1, 1, 1})
}

/*
Draw drop cap outline into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Go) DrawDropcapOutline(canvas gd.RID, pos gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawDropcapOutline(canvas, pos, gd.Int(1), gd.Color{1, 1, 1, 1})
}

/*
Returns caret character offset at the specified coordinates. This function always returns a valid position.
*/
func (self Go) HitTest(coords gd.Vector2) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).HitTest(coords)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TextParagraph
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("TextParagraph"))
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

func (self Go) CustomPunctuation() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetCustomPunctuation(gc).String())
}

func (self Go) SetCustomPunctuation(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCustomPunctuation(gc.String(value))
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

func (self Go) Alignment() gd.HorizontalAlignment {
	gc := gd.GarbageCollector(); _ = gc
		return gd.HorizontalAlignment(class(self).GetAlignment())
}

func (self Go) SetAlignment(value gd.HorizontalAlignment) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAlignment(value)
}

func (self Go) BreakFlags() classdb.TextServerLineBreakFlag {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextServerLineBreakFlag(class(self).GetBreakFlags())
}

func (self Go) SetBreakFlags(value classdb.TextServerLineBreakFlag) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBreakFlags(value)
}

func (self Go) JustificationFlags() classdb.TextServerJustificationFlag {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextServerJustificationFlag(class(self).GetJustificationFlags())
}

func (self Go) SetJustificationFlags(value classdb.TextServerJustificationFlag) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJustificationFlags(value)
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

func (self Go) Width() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetWidth()))
}

func (self Go) SetWidth(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetWidth(gd.Float(value))
}

func (self Go) MaxLinesVisible() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMaxLinesVisible()))
}

func (self Go) SetMaxLinesVisible(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxLinesVisible(gd.Int(value))
}

/*
Clears text paragraph (removes text and inline objects).
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetDirection(direction classdb.TextServerDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_set_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDirection() classdb.TextServerDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCustomPunctuation(custom_punctuation gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(custom_punctuation))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_set_custom_punctuation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCustomPunctuation(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_custom_punctuation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOrientation(orientation classdb.TextServerOrientation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, orientation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_set_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOrientation() classdb.TextServerOrientation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerOrientation](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_set_preserve_invalid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPreserveInvalid() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_preserve_invalid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_set_preserve_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPreserveControl() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_preserve_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_set_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets drop cap, overrides previously set drop cap. Drop cap (dropped capital) is a decorative element at the beginning of a paragraph that is larger than the rest of the text.
*/
//go:nosplit
func (self class) SetDropcap(text gd.String, font gdclass.Font, font_size gd.Int, dropcap_margins gd.Rect2, language gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, dropcap_margins)
	callframe.Arg(frame, mmm.Get(language))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_set_dropcap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes dropcap.
*/
//go:nosplit
func (self class) ClearDropcap()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_clear_dropcap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_add_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_add_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_resize_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlignment(alignment gd.HorizontalAlignment)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_set_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlignment() gd.HorizontalAlignment {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.HorizontalAlignment](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Aligns paragraph to the given tab-stops.
*/
//go:nosplit
func (self class) TabAlign(tab_stops gd.PackedFloat32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(tab_stops))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_tab_align, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetBreakFlags(flags classdb.TextServerLineBreakFlag)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_set_break_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBreakFlags() classdb.TextServerLineBreakFlag {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerLineBreakFlag](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_break_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetJustificationFlags(flags classdb.TextServerJustificationFlag)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_set_justification_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetJustificationFlags() classdb.TextServerJustificationFlag {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerJustificationFlag](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_justification_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_set_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextOverrunBehavior() classdb.TextServerOverrunBehavior {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerOverrunBehavior](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_set_ellipsis_char, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEllipsisChar(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_ellipsis_char, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWidth(width gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWidth() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the size of the bounding box of the paragraph, without line breaks.
*/
//go:nosplit
func (self class) GetNonWrappedSize() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_non_wrapped_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the size of the bounding box of the paragraph.
*/
//go:nosplit
func (self class) GetSize() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns TextServer full string buffer RID.
*/
//go:nosplit
func (self class) GetRid() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns TextServer line buffer RID.
*/
//go:nosplit
func (self class) GetLineRid(line gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_line_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns drop cap text buffer RID.
*/
//go:nosplit
func (self class) GetDropcapRid() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_dropcap_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns number of lines in the paragraph.
*/
//go:nosplit
func (self class) GetLineCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_line_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxLinesVisible(max_lines_visible gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_lines_visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_set_max_lines_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxLinesVisible() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_max_lines_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns array of inline objects in the line.
*/
//go:nosplit
func (self class) GetLineObjects(ctx gd.Lifetime, line gd.Int) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_line_objects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns bounding rectangle of the inline object.
*/
//go:nosplit
func (self class) GetLineObjectRect(line gd.Int, key gd.Variant) gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, mmm.Get(key))
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_line_object_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns size of the bounding box of the line of text. Returned size is rounded up.
*/
//go:nosplit
func (self class) GetLineSize(line gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_line_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns character range of the line.
*/
//go:nosplit
func (self class) GetLineRange(line gd.Int) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_line_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the text line ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
*/
//go:nosplit
func (self class) GetLineAscent(line gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_line_ascent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the text line descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
*/
//go:nosplit
func (self class) GetLineDescent(line gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_line_descent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns width (for horizontal layout) or height (for vertical) of the line of text.
*/
//go:nosplit
func (self class) GetLineWidth(line gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_line_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) GetLineUnderlinePosition(line gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_line_underline_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns thickness of the underline.
*/
//go:nosplit
func (self class) GetLineUnderlineThickness(line gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_line_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns drop cap bounding box size.
*/
//go:nosplit
func (self class) GetDropcapSize() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_dropcap_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns number of lines used by dropcap.
*/
//go:nosplit
func (self class) GetDropcapLines() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_get_dropcap_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Draw all lines of the text and drop cap into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) Draw(canvas gd.RID, pos gd.Vector2, color gd.Color, dc_color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, color)
	callframe.Arg(frame, dc_color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draw outlines of all lines of the text and drop cap into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) DrawOutline(canvas gd.RID, pos gd.Vector2, outline_size gd.Int, color gd.Color, dc_color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, outline_size)
	callframe.Arg(frame, color)
	callframe.Arg(frame, dc_color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_draw_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draw single line of text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) DrawLine(canvas gd.RID, pos gd.Vector2, line gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, line)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_draw_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draw outline of the single line of text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) DrawLineOutline(canvas gd.RID, pos gd.Vector2, line gd.Int, outline_size gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, line)
	callframe.Arg(frame, outline_size)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_draw_line_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draw drop cap into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) DrawDropcap(canvas gd.RID, pos gd.Vector2, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_draw_dropcap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draw drop cap outline into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) DrawDropcapOutline(canvas gd.RID, pos gd.Vector2, outline_size gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, outline_size)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_draw_dropcap_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns caret character offset at the specified coordinates. This function always returns a valid position.
*/
//go:nosplit
func (self class) HitTest(coords gd.Vector2) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextParagraph.Bind_hit_test, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTextParagraph() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTextParagraph() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("TextParagraph", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
