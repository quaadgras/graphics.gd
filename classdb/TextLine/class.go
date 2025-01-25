// Package TextLine provides methods for working with TextLine object instances.
package TextLine

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
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Rect2"
import "graphics.gd/classdb/Resource"

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
Abstraction over [TextServer] for handling a single line of text.
*/
type Instance [1]gdclass.TextLine

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTextLine() Instance
}

/*
Clears text line (removes text and inline objects).
*/
func (self Instance) Clear() { //gd:TextLine.clear
	class(self).Clear()
}

/*
Overrides BiDi for the structured text.
Override ranges should cover full source text without overlaps. BiDi algorithm will be used on each range separately.
*/
func (self Instance) SetBidiOverride(override []any) { //gd:TextLine.set_bidi_override
	class(self).SetBidiOverride(gd.EngineArrayFromSlice(override))
}

/*
Adds text span and font to draw it.
*/
func (self Instance) AddString(text string, font [1]gdclass.Font, font_size int) bool { //gd:TextLine.add_string
	return bool(class(self).AddString(gd.NewString(text), font, gd.Int(font_size), gd.NewString(""), gd.NewVariant(gd.NewVariant(([1]any{}[0])))))
}

/*
Adds inline object to the text buffer, [param key] must be unique. In the text, object is represented as [param length] object replacement characters.
*/
func (self Instance) AddObject(key any, size Vector2.XY) bool { //gd:TextLine.add_object
	return bool(class(self).AddObject(gd.NewVariant(key), gd.Vector2(size), 5, gd.Int(1), gd.Float(0.0)))
}

/*
Sets new size and alignment of embedded object.
*/
func (self Instance) ResizeObject(key any, size Vector2.XY) bool { //gd:TextLine.resize_object
	return bool(class(self).ResizeObject(gd.NewVariant(key), gd.Vector2(size), 5, gd.Float(0.0)))
}

/*
Aligns text to the given tab-stops.
*/
func (self Instance) TabAlign(tab_stops []float32) { //gd:TextLine.tab_align
	class(self).TabAlign(gd.NewPackedFloat32Slice(tab_stops))
}

/*
Returns array of inline objects.
*/
func (self Instance) GetObjects() []any { //gd:TextLine.get_objects
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).GetObjects())))
}

/*
Returns bounding rectangle of the inline object.
*/
func (self Instance) GetObjectRect(key any) Rect2.PositionSize { //gd:TextLine.get_object_rect
	return Rect2.PositionSize(class(self).GetObjectRect(gd.NewVariant(key)))
}

/*
Returns size of the bounding box of the text.
*/
func (self Instance) GetSize() Vector2.XY { //gd:TextLine.get_size
	return Vector2.XY(class(self).GetSize())
}

/*
Returns TextServer buffer RID.
*/
func (self Instance) GetRid() Resource.ID { //gd:TextLine.get_rid
	return Resource.ID(class(self).GetRid())
}

/*
Returns the text ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
*/
func (self Instance) GetLineAscent() Float.X { //gd:TextLine.get_line_ascent
	return Float.X(Float.X(class(self).GetLineAscent()))
}

/*
Returns the text descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
*/
func (self Instance) GetLineDescent() Float.X { //gd:TextLine.get_line_descent
	return Float.X(Float.X(class(self).GetLineDescent()))
}

/*
Returns width (for horizontal layout) or height (for vertical) of the text.
*/
func (self Instance) GetLineWidth() Float.X { //gd:TextLine.get_line_width
	return Float.X(Float.X(class(self).GetLineWidth()))
}

/*
Returns pixel offset of the underline below the baseline.
*/
func (self Instance) GetLineUnderlinePosition() Float.X { //gd:TextLine.get_line_underline_position
	return Float.X(Float.X(class(self).GetLineUnderlinePosition()))
}

/*
Returns thickness of the underline.
*/
func (self Instance) GetLineUnderlineThickness() Float.X { //gd:TextLine.get_line_underline_thickness
	return Float.X(Float.X(class(self).GetLineUnderlineThickness()))
}

/*
Draw text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Instance) Draw(canvas Resource.ID, pos Vector2.XY) { //gd:TextLine.draw
	class(self).Draw(canvas, gd.Vector2(pos), gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Draw text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Instance) DrawOutline(canvas Resource.ID, pos Vector2.XY) { //gd:TextLine.draw_outline
	class(self).DrawOutline(canvas, gd.Vector2(pos), gd.Int(1), gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Returns caret character offset at the specified pixel offset at the baseline. This function always returns a valid position.
*/
func (self Instance) HitTest(coords Float.X) int { //gd:TextLine.hit_test
	return int(int(class(self).HitTest(gd.Float(coords))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TextLine

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TextLine"))
	casted := Instance{*(*gdclass.TextLine)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Direction() gdclass.TextServerDirection {
	return gdclass.TextServerDirection(class(self).GetDirection())
}

func (self Instance) SetDirection(value gdclass.TextServerDirection) {
	class(self).SetDirection(value)
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

func (self Instance) Width() Float.X {
	return Float.X(Float.X(class(self).GetWidth()))
}

func (self Instance) SetWidth(value Float.X) {
	class(self).SetWidth(gd.Float(value))
}

func (self Instance) Alignment() HorizontalAlignment {
	return HorizontalAlignment(class(self).GetHorizontalAlignment())
}

func (self Instance) SetAlignment(value HorizontalAlignment) {
	class(self).SetHorizontalAlignment(value)
}

func (self Instance) Flags() gdclass.TextServerJustificationFlag {
	return gdclass.TextServerJustificationFlag(class(self).GetFlags())
}

func (self Instance) SetFlags(value gdclass.TextServerJustificationFlag) {
	class(self).SetFlags(value)
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

/*
Clears text line (removes text and inline objects).
*/
//go:nosplit
func (self class) Clear() { //gd:TextLine.clear
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetDirection(direction gdclass.TextServerDirection) { //gd:TextLine.set_direction
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDirection() gdclass.TextServerDirection { //gd:TextLine.get_direction
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOrientation(orientation gdclass.TextServerOrientation) { //gd:TextLine.set_orientation
	var frame = callframe.New()
	callframe.Arg(frame, orientation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_orientation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOrientation() gdclass.TextServerOrientation { //gd:TextLine.get_orientation
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerOrientation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_orientation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPreserveInvalid(enabled bool) { //gd:TextLine.set_preserve_invalid
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_preserve_invalid, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPreserveInvalid() bool { //gd:TextLine.get_preserve_invalid
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_preserve_invalid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPreserveControl(enabled bool) { //gd:TextLine.set_preserve_control
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_preserve_control, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPreserveControl() bool { //gd:TextLine.get_preserve_control
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_preserve_control, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Overrides BiDi for the structured text.
Override ranges should cover full source text without overlaps. BiDi algorithm will be used on each range separately.
*/
//go:nosplit
func (self class) SetBidiOverride(override Array.Any) { //gd:TextLine.set_bidi_override
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(override)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_bidi_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds text span and font to draw it.
*/
//go:nosplit
func (self class) AddString(text gd.String, font [1]gdclass.Font, font_size gd.Int, language gd.String, meta gd.Variant) bool { //gd:TextLine.add_string
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, pointers.Get(font[0])[0])
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, pointers.Get(language))
	callframe.Arg(frame, pointers.Get(meta))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_add_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds inline object to the text buffer, [param key] must be unique. In the text, object is represented as [param length] object replacement characters.
*/
//go:nosplit
func (self class) AddObject(key gd.Variant, size gd.Vector2, inline_align InlineAlignment, length gd.Int, baseline gd.Float) bool { //gd:TextLine.add_object
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(key))
	callframe.Arg(frame, size)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, length)
	callframe.Arg(frame, baseline)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_add_object, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets new size and alignment of embedded object.
*/
//go:nosplit
func (self class) ResizeObject(key gd.Variant, size gd.Vector2, inline_align InlineAlignment, baseline gd.Float) bool { //gd:TextLine.resize_object
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(key))
	callframe.Arg(frame, size)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, baseline)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_resize_object, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWidth(width gd.Float) { //gd:TextLine.set_width
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetWidth() gd.Float { //gd:TextLine.get_width
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHorizontalAlignment(alignment HorizontalAlignment) { //gd:TextLine.set_horizontal_alignment
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHorizontalAlignment() HorizontalAlignment { //gd:TextLine.get_horizontal_alignment
	var frame = callframe.New()
	var r_ret = callframe.Ret[HorizontalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Aligns text to the given tab-stops.
*/
//go:nosplit
func (self class) TabAlign(tab_stops gd.PackedFloat32Array) { //gd:TextLine.tab_align
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tab_stops))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_tab_align, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFlags(flags gdclass.TextServerJustificationFlag) { //gd:TextLine.set_flags
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFlags() gdclass.TextServerJustificationFlag { //gd:TextLine.get_flags
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerJustificationFlag](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextOverrunBehavior(overrun_behavior gdclass.TextServerOverrunBehavior) { //gd:TextLine.set_text_overrun_behavior
	var frame = callframe.New()
	callframe.Arg(frame, overrun_behavior)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextOverrunBehavior() gdclass.TextServerOverrunBehavior { //gd:TextLine.get_text_overrun_behavior
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerOverrunBehavior](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEllipsisChar(char gd.String) { //gd:TextLine.set_ellipsis_char
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(char))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_ellipsis_char, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEllipsisChar() gd.String { //gd:TextLine.get_ellipsis_char
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_ellipsis_char, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns array of inline objects.
*/
//go:nosplit
func (self class) GetObjects() Array.Any { //gd:TextLine.get_objects
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_objects, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns bounding rectangle of the inline object.
*/
//go:nosplit
func (self class) GetObjectRect(key gd.Variant) gd.Rect2 { //gd:TextLine.get_object_rect
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(key))
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_object_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns size of the bounding box of the text.
*/
//go:nosplit
func (self class) GetSize() gd.Vector2 { //gd:TextLine.get_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns TextServer buffer RID.
*/
//go:nosplit
func (self class) GetRid() gd.RID { //gd:TextLine.get_rid
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the text ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
*/
//go:nosplit
func (self class) GetLineAscent() gd.Float { //gd:TextLine.get_line_ascent
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_line_ascent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the text descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
*/
//go:nosplit
func (self class) GetLineDescent() gd.Float { //gd:TextLine.get_line_descent
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_line_descent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns width (for horizontal layout) or height (for vertical) of the text.
*/
//go:nosplit
func (self class) GetLineWidth() gd.Float { //gd:TextLine.get_line_width
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_line_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) GetLineUnderlinePosition() gd.Float { //gd:TextLine.get_line_underline_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_line_underline_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns thickness of the underline.
*/
//go:nosplit
func (self class) GetLineUnderlineThickness() gd.Float { //gd:TextLine.get_line_underline_thickness
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_line_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Draw text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) Draw(canvas gd.RID, pos gd.Vector2, color gd.Color) { //gd:TextLine.draw
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_draw, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draw text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) DrawOutline(canvas gd.RID, pos gd.Vector2, outline_size gd.Int, color gd.Color) { //gd:TextLine.draw_outline
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, outline_size)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_draw_outline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns caret character offset at the specified pixel offset at the baseline. This function always returns a valid position.
*/
//go:nosplit
func (self class) HitTest(coords gd.Float) gd.Int { //gd:TextLine.hit_test
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_hit_test, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTextLine() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTextLine() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("TextLine", func(ptr gd.Object) any { return [1]gdclass.TextLine{*(*gdclass.TextLine)(unsafe.Pointer(&ptr))} })
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
