// Package TextParagraph provides methods for working with TextParagraph object instances.
package TextParagraph

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
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Rect2"
import "graphics.gd/variant/Vector2i"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any

/*
Abstraction over [TextServer] for handling a single paragraph of text.
*/
type Instance [1]gdclass.TextParagraph

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTextParagraph() Instance
}

/*
Clears text paragraph (removes text and inline objects).
*/
func (self Instance) Clear() {
	class(self).Clear()
}

/*
Overrides BiDi for the structured text.
Override ranges should cover full source text without overlaps. BiDi algorithm will be used on each range separately.
*/
func (self Instance) SetBidiOverride(override []any) {
	class(self).SetBidiOverride(gd.EngineArrayFromSlice(override))
}

/*
Sets drop cap, overrides previously set drop cap. Drop cap (dropped capital) is a decorative element at the beginning of a paragraph that is larger than the rest of the text.
*/
func (self Instance) SetDropcap(text string, font [1]gdclass.Font, font_size int) bool {
	return bool(class(self).SetDropcap(gd.NewString(text), font, gd.Int(font_size), gd.Rect2(gd.NewRect2(0, 0, 0, 0)), gd.NewString("")))
}

/*
Removes dropcap.
*/
func (self Instance) ClearDropcap() {
	class(self).ClearDropcap()
}

/*
Adds text span and font to draw it.
*/
func (self Instance) AddString(text string, font [1]gdclass.Font, font_size int) bool {
	return bool(class(self).AddString(gd.NewString(text), font, gd.Int(font_size), gd.NewString(""), gd.NewVariant(gd.NewVariant(([1]any{}[0])))))
}

/*
Adds inline object to the text buffer, [param key] must be unique. In the text, object is represented as [param length] object replacement characters.
*/
func (self Instance) AddObject(key any, size Vector2.XY) bool {
	return bool(class(self).AddObject(gd.NewVariant(key), gd.Vector2(size), 5, gd.Int(1), gd.Float(0.0)))
}

/*
Sets new size and alignment of embedded object.
*/
func (self Instance) ResizeObject(key any, size Vector2.XY) bool {
	return bool(class(self).ResizeObject(gd.NewVariant(key), gd.Vector2(size), 5, gd.Float(0.0)))
}

/*
Aligns paragraph to the given tab-stops.
*/
func (self Instance) TabAlign(tab_stops []float32) {
	class(self).TabAlign(gd.NewPackedFloat32Slice(tab_stops))
}

/*
Returns the size of the bounding box of the paragraph, without line breaks.
*/
func (self Instance) GetNonWrappedSize() Vector2.XY {
	return Vector2.XY(class(self).GetNonWrappedSize())
}

/*
Returns the size of the bounding box of the paragraph.
*/
func (self Instance) GetSize() Vector2.XY {
	return Vector2.XY(class(self).GetSize())
}

/*
Returns TextServer full string buffer RID.
*/
func (self Instance) GetRid() Resource.ID {
	return Resource.ID(class(self).GetRid())
}

/*
Returns TextServer line buffer RID.
*/
func (self Instance) GetLineRid(line int) Resource.ID {
	return Resource.ID(class(self).GetLineRid(gd.Int(line)))
}

/*
Returns drop cap text buffer RID.
*/
func (self Instance) GetDropcapRid() Resource.ID {
	return Resource.ID(class(self).GetDropcapRid())
}

/*
Returns number of lines in the paragraph.
*/
func (self Instance) GetLineCount() int {
	return int(int(class(self).GetLineCount()))
}

/*
Returns array of inline objects in the line.
*/
func (self Instance) GetLineObjects(line int) []any {
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).GetLineObjects(gd.Int(line)))))
}

/*
Returns bounding rectangle of the inline object.
*/
func (self Instance) GetLineObjectRect(line int, key any) Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetLineObjectRect(gd.Int(line), gd.NewVariant(key)))
}

/*
Returns size of the bounding box of the line of text. Returned size is rounded up.
*/
func (self Instance) GetLineSize(line int) Vector2.XY {
	return Vector2.XY(class(self).GetLineSize(gd.Int(line)))
}

/*
Returns character range of the line.
*/
func (self Instance) GetLineRange(line int) Vector2i.XY {
	return Vector2i.XY(class(self).GetLineRange(gd.Int(line)))
}

/*
Returns the text line ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
*/
func (self Instance) GetLineAscent(line int) Float.X {
	return Float.X(Float.X(class(self).GetLineAscent(gd.Int(line))))
}

/*
Returns the text line descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
*/
func (self Instance) GetLineDescent(line int) Float.X {
	return Float.X(Float.X(class(self).GetLineDescent(gd.Int(line))))
}

/*
Returns width (for horizontal layout) or height (for vertical) of the line of text.
*/
func (self Instance) GetLineWidth(line int) Float.X {
	return Float.X(Float.X(class(self).GetLineWidth(gd.Int(line))))
}

/*
Returns pixel offset of the underline below the baseline.
*/
func (self Instance) GetLineUnderlinePosition(line int) Float.X {
	return Float.X(Float.X(class(self).GetLineUnderlinePosition(gd.Int(line))))
}

/*
Returns thickness of the underline.
*/
func (self Instance) GetLineUnderlineThickness(line int) Float.X {
	return Float.X(Float.X(class(self).GetLineUnderlineThickness(gd.Int(line))))
}

/*
Returns drop cap bounding box size.
*/
func (self Instance) GetDropcapSize() Vector2.XY {
	return Vector2.XY(class(self).GetDropcapSize())
}

/*
Returns number of lines used by dropcap.
*/
func (self Instance) GetDropcapLines() int {
	return int(int(class(self).GetDropcapLines()))
}

/*
Draw all lines of the text and drop cap into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Instance) Draw(canvas Resource.ID, pos Vector2.XY) {
	class(self).Draw(canvas, gd.Vector2(pos), gd.Color(gd.Color{1, 1, 1, 1}), gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Draw outlines of all lines of the text and drop cap into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Instance) DrawOutline(canvas Resource.ID, pos Vector2.XY) {
	class(self).DrawOutline(canvas, gd.Vector2(pos), gd.Int(1), gd.Color(gd.Color{1, 1, 1, 1}), gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Draw single line of text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Instance) DrawLine(canvas Resource.ID, pos Vector2.XY, line int) {
	class(self).DrawLine(canvas, gd.Vector2(pos), gd.Int(line), gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Draw outline of the single line of text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Instance) DrawLineOutline(canvas Resource.ID, pos Vector2.XY, line int) {
	class(self).DrawLineOutline(canvas, gd.Vector2(pos), gd.Int(line), gd.Int(1), gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Draw drop cap into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Instance) DrawDropcap(canvas Resource.ID, pos Vector2.XY) {
	class(self).DrawDropcap(canvas, gd.Vector2(pos), gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Draw drop cap outline into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Instance) DrawDropcapOutline(canvas Resource.ID, pos Vector2.XY) {
	class(self).DrawDropcapOutline(canvas, gd.Vector2(pos), gd.Int(1), gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Returns caret character offset at the specified coordinates. This function always returns a valid position.
*/
func (self Instance) HitTest(coords Vector2.XY) int {
	return int(int(class(self).HitTest(gd.Vector2(coords))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TextParagraph

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TextParagraph"))
	casted := Instance{*(*gdclass.TextParagraph)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Direction() gdclass.TextServerDirection {
	return gdclass.TextServerDirection(class(self).GetDirection())
}

func (self Instance) SetDirection(value gdclass.TextServerDirection) {
	class(self).SetDirection(value)
}

func (self Instance) CustomPunctuation() string {
	return string(class(self).GetCustomPunctuation().String())
}

func (self Instance) SetCustomPunctuation(value string) {
	class(self).SetCustomPunctuation(gd.NewString(value))
}

func (self Instance) Orientation() gdclass.TextServerOrientation {
	return gdclass.TextServerOrientation(class(self).GetOrientation())
}

func (self Instance) SetOrientation(value gdclass.TextServerOrientation) {
	class(self).SetOrientation(value)
}

func (self Instance) PreserveInvalid() bool {
	return bool(class(self).GetPreserveInvalid())
}

func (self Instance) SetPreserveInvalid(value bool) {
	class(self).SetPreserveInvalid(value)
}

func (self Instance) PreserveControl() bool {
	return bool(class(self).GetPreserveControl())
}

func (self Instance) SetPreserveControl(value bool) {
	class(self).SetPreserveControl(value)
}

func (self Instance) Alignment() HorizontalAlignment {
	return HorizontalAlignment(class(self).GetAlignment())
}

func (self Instance) SetAlignment(value HorizontalAlignment) {
	class(self).SetAlignment(value)
}

func (self Instance) BreakFlags() gdclass.TextServerLineBreakFlag {
	return gdclass.TextServerLineBreakFlag(class(self).GetBreakFlags())
}

func (self Instance) SetBreakFlags(value gdclass.TextServerLineBreakFlag) {
	class(self).SetBreakFlags(value)
}

func (self Instance) JustificationFlags() gdclass.TextServerJustificationFlag {
	return gdclass.TextServerJustificationFlag(class(self).GetJustificationFlags())
}

func (self Instance) SetJustificationFlags(value gdclass.TextServerJustificationFlag) {
	class(self).SetJustificationFlags(value)
}

func (self Instance) TextOverrunBehavior() gdclass.TextServerOverrunBehavior {
	return gdclass.TextServerOverrunBehavior(class(self).GetTextOverrunBehavior())
}

func (self Instance) SetTextOverrunBehavior(value gdclass.TextServerOverrunBehavior) {
	class(self).SetTextOverrunBehavior(value)
}

func (self Instance) EllipsisChar() string {
	return string(class(self).GetEllipsisChar().String())
}

func (self Instance) SetEllipsisChar(value string) {
	class(self).SetEllipsisChar(gd.NewString(value))
}

func (self Instance) Width() Float.X {
	return Float.X(Float.X(class(self).GetWidth()))
}

func (self Instance) SetWidth(value Float.X) {
	class(self).SetWidth(gd.Float(value))
}

func (self Instance) MaxLinesVisible() int {
	return int(int(class(self).GetMaxLinesVisible()))
}

func (self Instance) SetMaxLinesVisible(value int) {
	class(self).SetMaxLinesVisible(gd.Int(value))
}

/*
Clears text paragraph (removes text and inline objects).
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetDirection(direction gdclass.TextServerDirection) {
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDirection() gdclass.TextServerDirection {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCustomPunctuation(custom_punctuation gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(custom_punctuation))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_custom_punctuation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCustomPunctuation() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_custom_punctuation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOrientation(orientation gdclass.TextServerOrientation) {
	var frame = callframe.New()
	callframe.Arg(frame, orientation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_orientation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOrientation() gdclass.TextServerOrientation {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerOrientation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_orientation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPreserveInvalid(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_preserve_invalid, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPreserveInvalid() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_preserve_invalid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPreserveControl(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_preserve_control, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPreserveControl() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_preserve_control, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Overrides BiDi for the structured text.
Override ranges should cover full source text without overlaps. BiDi algorithm will be used on each range separately.
*/
//go:nosplit
func (self class) SetBidiOverride(override Array.Any) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(override)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_bidi_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets drop cap, overrides previously set drop cap. Drop cap (dropped capital) is a decorative element at the beginning of a paragraph that is larger than the rest of the text.
*/
//go:nosplit
func (self class) SetDropcap(text gd.String, font [1]gdclass.Font, font_size gd.Int, dropcap_margins gd.Rect2, language gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, pointers.Get(font[0])[0])
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, dropcap_margins)
	callframe.Arg(frame, pointers.Get(language))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_dropcap, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes dropcap.
*/
//go:nosplit
func (self class) ClearDropcap() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_clear_dropcap, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds text span and font to draw it.
*/
//go:nosplit
func (self class) AddString(text gd.String, font [1]gdclass.Font, font_size gd.Int, language gd.String, meta gd.Variant) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, pointers.Get(font[0])[0])
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, pointers.Get(language))
	callframe.Arg(frame, pointers.Get(meta))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_add_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds inline object to the text buffer, [param key] must be unique. In the text, object is represented as [param length] object replacement characters.
*/
//go:nosplit
func (self class) AddObject(key gd.Variant, size gd.Vector2, inline_align InlineAlignment, length gd.Int, baseline gd.Float) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(key))
	callframe.Arg(frame, size)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, length)
	callframe.Arg(frame, baseline)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_add_object, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets new size and alignment of embedded object.
*/
//go:nosplit
func (self class) ResizeObject(key gd.Variant, size gd.Vector2, inline_align InlineAlignment, baseline gd.Float) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(key))
	callframe.Arg(frame, size)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, baseline)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_resize_object, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlignment(alignment HorizontalAlignment) {
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlignment() HorizontalAlignment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[HorizontalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Aligns paragraph to the given tab-stops.
*/
//go:nosplit
func (self class) TabAlign(tab_stops gd.PackedFloat32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tab_stops))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_tab_align, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetBreakFlags(flags gdclass.TextServerLineBreakFlag) {
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_break_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBreakFlags() gdclass.TextServerLineBreakFlag {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerLineBreakFlag](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_break_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJustificationFlags(flags gdclass.TextServerJustificationFlag) {
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_justification_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetJustificationFlags() gdclass.TextServerJustificationFlag {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerJustificationFlag](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_justification_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextOverrunBehavior(overrun_behavior gdclass.TextServerOverrunBehavior) {
	var frame = callframe.New()
	callframe.Arg(frame, overrun_behavior)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextOverrunBehavior() gdclass.TextServerOverrunBehavior {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerOverrunBehavior](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEllipsisChar(char gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(char))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_ellipsis_char, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEllipsisChar() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_ellipsis_char, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWidth(width gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetWidth() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the size of the bounding box of the paragraph, without line breaks.
*/
//go:nosplit
func (self class) GetNonWrappedSize() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_non_wrapped_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the size of the bounding box of the paragraph.
*/
//go:nosplit
func (self class) GetSize() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns TextServer full string buffer RID.
*/
//go:nosplit
func (self class) GetRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns TextServer line buffer RID.
*/
//go:nosplit
func (self class) GetLineRid(line gd.Int) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns drop cap text buffer RID.
*/
//go:nosplit
func (self class) GetDropcapRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_dropcap_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns number of lines in the paragraph.
*/
//go:nosplit
func (self class) GetLineCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxLinesVisible(max_lines_visible gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, max_lines_visible)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_max_lines_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxLinesVisible() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_max_lines_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns array of inline objects in the line.
*/
//go:nosplit
func (self class) GetLineObjects(line gd.Int) Array.Any {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_objects, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns bounding rectangle of the inline object.
*/
//go:nosplit
func (self class) GetLineObjectRect(line gd.Int, key gd.Variant) gd.Rect2 {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, pointers.Get(key))
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_object_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns size of the bounding box of the line of text. Returned size is rounded up.
*/
//go:nosplit
func (self class) GetLineSize(line gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns character range of the line.
*/
//go:nosplit
func (self class) GetLineRange(line gd.Int) gd.Vector2i {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the text line ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
*/
//go:nosplit
func (self class) GetLineAscent(line gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_ascent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the text line descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
*/
//go:nosplit
func (self class) GetLineDescent(line gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_descent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns width (for horizontal layout) or height (for vertical) of the line of text.
*/
//go:nosplit
func (self class) GetLineWidth(line gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) GetLineUnderlinePosition(line gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_underline_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns thickness of the underline.
*/
//go:nosplit
func (self class) GetLineUnderlineThickness(line gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns drop cap bounding box size.
*/
//go:nosplit
func (self class) GetDropcapSize() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_dropcap_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns number of lines used by dropcap.
*/
//go:nosplit
func (self class) GetDropcapLines() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_dropcap_lines, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Draw all lines of the text and drop cap into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) Draw(canvas gd.RID, pos gd.Vector2, color gd.Color, dc_color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, color)
	callframe.Arg(frame, dc_color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_draw, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draw outlines of all lines of the text and drop cap into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) DrawOutline(canvas gd.RID, pos gd.Vector2, outline_size gd.Int, color gd.Color, dc_color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, outline_size)
	callframe.Arg(frame, color)
	callframe.Arg(frame, dc_color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_draw_outline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draw single line of text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) DrawLine(canvas gd.RID, pos gd.Vector2, line gd.Int, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, line)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_draw_line, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draw outline of the single line of text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) DrawLineOutline(canvas gd.RID, pos gd.Vector2, line gd.Int, outline_size gd.Int, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, line)
	callframe.Arg(frame, outline_size)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_draw_line_outline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draw drop cap into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) DrawDropcap(canvas gd.RID, pos gd.Vector2, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_draw_dropcap, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draw drop cap outline into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) DrawDropcapOutline(canvas gd.RID, pos gd.Vector2, outline_size gd.Int, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, outline_size)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_draw_dropcap_outline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns caret character offset at the specified coordinates. This function always returns a valid position.
*/
//go:nosplit
func (self class) HitTest(coords gd.Vector2) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_hit_test, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTextParagraph() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTextParagraph() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("TextParagraph", func(ptr gd.Object) any {
		return [1]gdclass.TextParagraph{*(*gdclass.TextParagraph)(unsafe.Pointer(&ptr))}
	})
}

type HorizontalAlignment int

const (
	/*Horizontal left alignment, usually for text-derived classes.*/
	HorizontalAlignmentLeft HorizontalAlignment = 0
	/*Horizontal center alignment, usually for text-derived classes.*/
	HorizontalAlignmentCenter HorizontalAlignment = 1
	/*Horizontal right alignment, usually for text-derived classes.*/
	HorizontalAlignmentRight HorizontalAlignment = 2
	/*Expand row to fit width, usually for text-derived classes.*/
	HorizontalAlignmentFill HorizontalAlignment = 3
)

type InlineAlignment int

const (
	/*Aligns the top of the inline object (e.g. image, table) to the position of the text specified by [code]INLINE_ALIGNMENT_TO_*[/code] constant.*/
	InlineAlignmentTopTo InlineAlignment = 0
	/*Aligns the center of the inline object (e.g. image, table) to the position of the text specified by [code]INLINE_ALIGNMENT_TO_*[/code] constant.*/
	InlineAlignmentCenterTo InlineAlignment = 1
	/*Aligns the baseline (user defined) of the inline object (e.g. image, table) to the position of the text specified by [code]INLINE_ALIGNMENT_TO_*[/code] constant.*/
	InlineAlignmentBaselineTo InlineAlignment = 3
	/*Aligns the bottom of the inline object (e.g. image, table) to the position of the text specified by [code]INLINE_ALIGNMENT_TO_*[/code] constant.*/
	InlineAlignmentBottomTo InlineAlignment = 2
	/*Aligns the position of the inline object (e.g. image, table) specified by [code]INLINE_ALIGNMENT_*_TO[/code] constant to the top of the text.*/
	InlineAlignmentToTop InlineAlignment = 0
	/*Aligns the position of the inline object (e.g. image, table) specified by [code]INLINE_ALIGNMENT_*_TO[/code] constant to the center of the text.*/
	InlineAlignmentToCenter InlineAlignment = 4
	/*Aligns the position of the inline object (e.g. image, table) specified by [code]INLINE_ALIGNMENT_*_TO[/code] constant to the baseline of the text.*/
	InlineAlignmentToBaseline InlineAlignment = 8
	/*Aligns inline object (e.g. image, table) to the bottom of the text.*/
	InlineAlignmentToBottom InlineAlignment = 12
	/*Aligns top of the inline object (e.g. image, table) to the top of the text. Equivalent to [code]INLINE_ALIGNMENT_TOP_TO | INLINE_ALIGNMENT_TO_TOP[/code].*/
	InlineAlignmentTop InlineAlignment = 0
	/*Aligns center of the inline object (e.g. image, table) to the center of the text. Equivalent to [code]INLINE_ALIGNMENT_CENTER_TO | INLINE_ALIGNMENT_TO_CENTER[/code].*/
	InlineAlignmentCenter InlineAlignment = 5
	/*Aligns bottom of the inline object (e.g. image, table) to the bottom of the text. Equivalent to [code]INLINE_ALIGNMENT_BOTTOM_TO | INLINE_ALIGNMENT_TO_BOTTOM[/code].*/
	InlineAlignmentBottom InlineAlignment = 14
	/*A bit mask for [code]INLINE_ALIGNMENT_*_TO[/code] alignment constants.*/
	InlineAlignmentImageMask InlineAlignment = 3
	/*A bit mask for [code]INLINE_ALIGNMENT_TO_*[/code] alignment constants.*/
	InlineAlignmentTextMask InlineAlignment = 12
)
