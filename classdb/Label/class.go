// Package Label provides methods for working with Label object instances.
package Label

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
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Rect2"

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
A control for displaying plain text. It gives you control over the horizontal and vertical alignment and can wrap the text inside the node's bounding rectangle. It doesn't support bold, italics, or other rich text formatting. For that, use [RichTextLabel] instead.
*/
type Instance [1]gdclass.Label

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsLabel() Instance
}

/*
Returns the height of the line [param line].
If [param line] is set to [code]-1[/code], returns the biggest line height.
If there are no lines, returns font size in pixels.
*/
func (self Instance) GetLineHeight() int { //gd:Label.get_line_height
	return int(int(class(self).GetLineHeight(gd.Int(-1))))
}

/*
Returns the number of lines of text the Label has.
*/
func (self Instance) GetLineCount() int { //gd:Label.get_line_count
	return int(int(class(self).GetLineCount()))
}

/*
Returns the number of lines shown. Useful if the [Label]'s height cannot currently display all lines.
*/
func (self Instance) GetVisibleLineCount() int { //gd:Label.get_visible_line_count
	return int(int(class(self).GetVisibleLineCount()))
}

/*
Returns the total number of printable characters in the text (excluding spaces and newlines).
*/
func (self Instance) GetTotalCharacterCount() int { //gd:Label.get_total_character_count
	return int(int(class(self).GetTotalCharacterCount()))
}

/*
Returns the bounding rectangle of the character at position [param pos]. If the character is a non-visual character or [param pos] is outside the valid range, an empty [Rect2] is returned. If the character is a part of a composite grapheme, the bounding rectangle of the whole grapheme is returned.
*/
func (self Instance) GetCharacterBounds(pos int) Rect2.PositionSize { //gd:Label.get_character_bounds
	return Rect2.PositionSize(class(self).GetCharacterBounds(gd.Int(pos)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Label

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Label"))
	casted := Instance{*(*gdclass.Label)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Text() string {
	return string(class(self).GetText().String())
}

func (self Instance) SetText(value string) {
	class(self).SetText(gd.NewString(value))
}

func (self Instance) LabelSettings() [1]gdclass.LabelSettings {
	return [1]gdclass.LabelSettings(class(self).GetLabelSettings())
}

func (self Instance) SetLabelSettings(value [1]gdclass.LabelSettings) {
	class(self).SetLabelSettings(value)
}

func (self Instance) HorizontalAlignment() HorizontalAlignment {
	return HorizontalAlignment(class(self).GetHorizontalAlignment())
}

func (self Instance) SetHorizontalAlignment(value HorizontalAlignment) {
	class(self).SetHorizontalAlignment(value)
}

func (self Instance) VerticalAlignment() VerticalAlignment {
	return VerticalAlignment(class(self).GetVerticalAlignment())
}

func (self Instance) SetVerticalAlignment(value VerticalAlignment) {
	class(self).SetVerticalAlignment(value)
}

func (self Instance) AutowrapMode() gdclass.TextServerAutowrapMode {
	return gdclass.TextServerAutowrapMode(class(self).GetAutowrapMode())
}

func (self Instance) SetAutowrapMode(value gdclass.TextServerAutowrapMode) {
	class(self).SetAutowrapMode(value)
}

func (self Instance) JustificationFlags() gdclass.TextServerJustificationFlag {
	return gdclass.TextServerJustificationFlag(class(self).GetJustificationFlags())
}

func (self Instance) SetJustificationFlags(value gdclass.TextServerJustificationFlag) {
	class(self).SetJustificationFlags(value)
}

func (self Instance) ClipText() bool {
	return bool(class(self).IsClippingText())
}

func (self Instance) SetClipText(value bool) {
	class(self).SetClipText(value)
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

func (self Instance) Uppercase() bool {
	return bool(class(self).IsUppercase())
}

func (self Instance) SetUppercase(value bool) {
	class(self).SetUppercase(value)
}

func (self Instance) TabStops() []float32 {
	return []float32(class(self).GetTabStops().AsSlice())
}

func (self Instance) SetTabStops(value []float32) {
	class(self).SetTabStops(gd.NewPackedFloat32Slice(value))
}

func (self Instance) LinesSkipped() int {
	return int(int(class(self).GetLinesSkipped()))
}

func (self Instance) SetLinesSkipped(value int) {
	class(self).SetLinesSkipped(gd.Int(value))
}

func (self Instance) MaxLinesVisible() int {
	return int(int(class(self).GetMaxLinesVisible()))
}

func (self Instance) SetMaxLinesVisible(value int) {
	class(self).SetMaxLinesVisible(gd.Int(value))
}

func (self Instance) VisibleCharacters() int {
	return int(int(class(self).GetVisibleCharacters()))
}

func (self Instance) SetVisibleCharacters(value int) {
	class(self).SetVisibleCharacters(gd.Int(value))
}

func (self Instance) VisibleCharactersBehavior() gdclass.TextServerVisibleCharactersBehavior {
	return gdclass.TextServerVisibleCharactersBehavior(class(self).GetVisibleCharactersBehavior())
}

func (self Instance) SetVisibleCharactersBehavior(value gdclass.TextServerVisibleCharactersBehavior) {
	class(self).SetVisibleCharactersBehavior(value)
}

func (self Instance) VisibleRatio() Float.X {
	return Float.X(Float.X(class(self).GetVisibleRatio()))
}

func (self Instance) SetVisibleRatio(value Float.X) {
	class(self).SetVisibleRatio(gd.Float(value))
}

func (self Instance) TextDirection() gdclass.ControlTextDirection {
	return gdclass.ControlTextDirection(class(self).GetTextDirection())
}

func (self Instance) SetTextDirection(value gdclass.ControlTextDirection) {
	class(self).SetTextDirection(value)
}

func (self Instance) Language() string {
	return string(class(self).GetLanguage().String())
}

func (self Instance) SetLanguage(value string) {
	class(self).SetLanguage(gd.NewString(value))
}

func (self Instance) StructuredTextBidiOverride() gdclass.TextServerStructuredTextParser {
	return gdclass.TextServerStructuredTextParser(class(self).GetStructuredTextBidiOverride())
}

func (self Instance) SetStructuredTextBidiOverride(value gdclass.TextServerStructuredTextParser) {
	class(self).SetStructuredTextBidiOverride(value)
}

func (self Instance) StructuredTextBidiOverrideOptions() []any {
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).GetStructuredTextBidiOverrideOptions())))
}

func (self Instance) SetStructuredTextBidiOverrideOptions(value []any) {
	class(self).SetStructuredTextBidiOverrideOptions(gd.EngineArrayFromSlice(value))
}

//go:nosplit
func (self class) SetHorizontalAlignment(alignment HorizontalAlignment) { //gd:Label.set_horizontal_alignment
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHorizontalAlignment() HorizontalAlignment { //gd:Label.get_horizontal_alignment
	var frame = callframe.New()
	var r_ret = callframe.Ret[HorizontalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVerticalAlignment(alignment VerticalAlignment) { //gd:Label.set_vertical_alignment
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_vertical_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVerticalAlignment() VerticalAlignment { //gd:Label.get_vertical_alignment
	var frame = callframe.New()
	var r_ret = callframe.Ret[VerticalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_vertical_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetText(text gd.String) { //gd:Label.set_text
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetText() gd.String { //gd:Label.get_text
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLabelSettings(settings [1]gdclass.LabelSettings) { //gd:Label.set_label_settings
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(settings[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_label_settings, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLabelSettings() [1]gdclass.LabelSettings { //gd:Label.get_label_settings
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_label_settings, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.LabelSettings{gd.PointerWithOwnershipTransferredToGo[gdclass.LabelSettings](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextDirection(direction gdclass.ControlTextDirection) { //gd:Label.set_text_direction
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextDirection() gdclass.ControlTextDirection { //gd:Label.get_text_direction
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.ControlTextDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLanguage(language gd.String) { //gd:Label.set_language
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(language))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLanguage() gd.String { //gd:Label.get_language
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutowrapMode(autowrap_mode gdclass.TextServerAutowrapMode) { //gd:Label.set_autowrap_mode
	var frame = callframe.New()
	callframe.Arg(frame, autowrap_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutowrapMode() gdclass.TextServerAutowrapMode { //gd:Label.get_autowrap_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerAutowrapMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJustificationFlags(justification_flags gdclass.TextServerJustificationFlag) { //gd:Label.set_justification_flags
	var frame = callframe.New()
	callframe.Arg(frame, justification_flags)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_justification_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetJustificationFlags() gdclass.TextServerJustificationFlag { //gd:Label.get_justification_flags
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerJustificationFlag](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_justification_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetClipText(enable bool) { //gd:Label.set_clip_text
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_clip_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsClippingText() bool { //gd:Label.is_clipping_text
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_is_clipping_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTabStops(tab_stops gd.PackedFloat32Array) { //gd:Label.set_tab_stops
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tab_stops))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_tab_stops, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTabStops() gd.PackedFloat32Array { //gd:Label.get_tab_stops
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_tab_stops, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedFloat32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextOverrunBehavior(overrun_behavior gdclass.TextServerOverrunBehavior) { //gd:Label.set_text_overrun_behavior
	var frame = callframe.New()
	callframe.Arg(frame, overrun_behavior)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextOverrunBehavior() gdclass.TextServerOverrunBehavior { //gd:Label.get_text_overrun_behavior
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerOverrunBehavior](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEllipsisChar(char gd.String) { //gd:Label.set_ellipsis_char
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(char))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_ellipsis_char, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEllipsisChar() gd.String { //gd:Label.get_ellipsis_char
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_ellipsis_char, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUppercase(enable bool) { //gd:Label.set_uppercase
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_uppercase, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUppercase() bool { //gd:Label.is_uppercase
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_is_uppercase, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the height of the line [param line].
If [param line] is set to [code]-1[/code], returns the biggest line height.
If there are no lines, returns font size in pixels.
*/
//go:nosplit
func (self class) GetLineHeight(line gd.Int) gd.Int { //gd:Label.get_line_height
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_line_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of lines of text the Label has.
*/
//go:nosplit
func (self class) GetLineCount() gd.Int { //gd:Label.get_line_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_line_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of lines shown. Useful if the [Label]'s height cannot currently display all lines.
*/
//go:nosplit
func (self class) GetVisibleLineCount() gd.Int { //gd:Label.get_visible_line_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_visible_line_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the total number of printable characters in the text (excluding spaces and newlines).
*/
//go:nosplit
func (self class) GetTotalCharacterCount() gd.Int { //gd:Label.get_total_character_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_total_character_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibleCharacters(amount gd.Int) { //gd:Label.set_visible_characters
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_visible_characters, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibleCharacters() gd.Int { //gd:Label.get_visible_characters
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_visible_characters, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetVisibleCharactersBehavior() gdclass.TextServerVisibleCharactersBehavior { //gd:Label.get_visible_characters_behavior
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerVisibleCharactersBehavior](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_visible_characters_behavior, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibleCharactersBehavior(behavior gdclass.TextServerVisibleCharactersBehavior) { //gd:Label.set_visible_characters_behavior
	var frame = callframe.New()
	callframe.Arg(frame, behavior)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_visible_characters_behavior, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetVisibleRatio(ratio gd.Float) { //gd:Label.set_visible_ratio
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_visible_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibleRatio() gd.Float { //gd:Label.get_visible_ratio
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_visible_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinesSkipped(lines_skipped gd.Int) { //gd:Label.set_lines_skipped
	var frame = callframe.New()
	callframe.Arg(frame, lines_skipped)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_lines_skipped, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinesSkipped() gd.Int { //gd:Label.get_lines_skipped
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_lines_skipped, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxLinesVisible(lines_visible gd.Int) { //gd:Label.set_max_lines_visible
	var frame = callframe.New()
	callframe.Arg(frame, lines_visible)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_max_lines_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxLinesVisible() gd.Int { //gd:Label.get_max_lines_visible
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_max_lines_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStructuredTextBidiOverride(parser gdclass.TextServerStructuredTextParser) { //gd:Label.set_structured_text_bidi_override
	var frame = callframe.New()
	callframe.Arg(frame, parser)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStructuredTextBidiOverride() gdclass.TextServerStructuredTextParser { //gd:Label.get_structured_text_bidi_override
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerStructuredTextParser](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStructuredTextBidiOverrideOptions(args Array.Any) { //gd:Label.set_structured_text_bidi_override_options
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(args)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStructuredTextBidiOverrideOptions() Array.Any { //gd:Label.get_structured_text_bidi_override_options
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the bounding rectangle of the character at position [param pos]. If the character is a non-visual character or [param pos] is outside the valid range, an empty [Rect2] is returned. If the character is a part of a composite grapheme, the bounding rectangle of the whole grapheme is returned.
*/
//go:nosplit
func (self class) GetCharacterBounds(pos gd.Int) gd.Rect2 { //gd:Label.get_character_bounds
	var frame = callframe.New()
	callframe.Arg(frame, pos)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_character_bounds, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsLabel() Advanced           { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsLabel() Instance        { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.Advanced { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(Control.Advanced(self.AsControl()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Control.Instance(self.AsControl()), name)
	}
}
func init() {
	gdclass.Register("Label", func(ptr gd.Object) any { return [1]gdclass.Label{*(*gdclass.Label)(unsafe.Pointer(&ptr))} })
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

type VerticalAlignment int

const (
	/*Vertical top alignment, usually for text-derived classes.*/
	VerticalAlignmentTop VerticalAlignment = 0
	/*Vertical center alignment, usually for text-derived classes.*/
	VerticalAlignmentCenter VerticalAlignment = 1
	/*Vertical bottom alignment, usually for text-derived classes.*/
	VerticalAlignmentBottom VerticalAlignment = 2
	/*Expand rows to fit height, usually for text-derived classes.*/
	VerticalAlignmentFill VerticalAlignment = 3
)
