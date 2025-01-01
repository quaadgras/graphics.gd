package Label

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Control"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Rect2"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A control for displaying plain text. It gives you control over the horizontal and vertical alignment and can wrap the text inside the node's bounding rectangle. It doesn't support bold, italics, or other rich text formatting. For that, use [RichTextLabel] instead.
*/
type Instance [1]classdb.Label
type Any interface {
	gd.IsClass
	AsLabel() Instance
}

/*
Returns the height of the line [param line].
If [param line] is set to [code]-1[/code], returns the biggest line height.
If there are no lines, returns font size in pixels.
*/
func (self Instance) GetLineHeight() int {
	return int(int(class(self).GetLineHeight(gd.Int(-1))))
}

/*
Returns the number of lines of text the Label has.
*/
func (self Instance) GetLineCount() int {
	return int(int(class(self).GetLineCount()))
}

/*
Returns the number of lines shown. Useful if the [Label]'s height cannot currently display all lines.
*/
func (self Instance) GetVisibleLineCount() int {
	return int(int(class(self).GetVisibleLineCount()))
}

/*
Returns the total number of printable characters in the text (excluding spaces and newlines).
*/
func (self Instance) GetTotalCharacterCount() int {
	return int(int(class(self).GetTotalCharacterCount()))
}

/*
Returns the bounding rectangle of the character at position [param pos]. If the character is a non-visual character or [param pos] is outside the valid range, an empty [Rect2] is returned. If the character is a part of a composite grapheme, the bounding rectangle of the whole grapheme is returned.
*/
func (self Instance) GetCharacterBounds(pos int) Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetCharacterBounds(gd.Int(pos)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Label

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Label"))
	return Instance{classdb.Label(object)}
}

func (self Instance) Text() string {
	return string(class(self).GetText().String())
}

func (self Instance) SetText(value string) {
	class(self).SetText(gd.NewString(value))
}

func (self Instance) LabelSettings() objects.LabelSettings {
	return objects.LabelSettings(class(self).GetLabelSettings())
}

func (self Instance) SetLabelSettings(value objects.LabelSettings) {
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

func (self Instance) AutowrapMode() classdb.TextServerAutowrapMode {
	return classdb.TextServerAutowrapMode(class(self).GetAutowrapMode())
}

func (self Instance) SetAutowrapMode(value classdb.TextServerAutowrapMode) {
	class(self).SetAutowrapMode(value)
}

func (self Instance) JustificationFlags() classdb.TextServerJustificationFlag {
	return classdb.TextServerJustificationFlag(class(self).GetJustificationFlags())
}

func (self Instance) SetJustificationFlags(value classdb.TextServerJustificationFlag) {
	class(self).SetJustificationFlags(value)
}

func (self Instance) ClipText() bool {
	return bool(class(self).IsClippingText())
}

func (self Instance) SetClipText(value bool) {
	class(self).SetClipText(value)
}

func (self Instance) TextOverrunBehavior() classdb.TextServerOverrunBehavior {
	return classdb.TextServerOverrunBehavior(class(self).GetTextOverrunBehavior())
}

func (self Instance) SetTextOverrunBehavior(value classdb.TextServerOverrunBehavior) {
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

func (self Instance) VisibleCharactersBehavior() classdb.TextServerVisibleCharactersBehavior {
	return classdb.TextServerVisibleCharactersBehavior(class(self).GetVisibleCharactersBehavior())
}

func (self Instance) SetVisibleCharactersBehavior(value classdb.TextServerVisibleCharactersBehavior) {
	class(self).SetVisibleCharactersBehavior(value)
}

func (self Instance) VisibleRatio() Float.X {
	return Float.X(Float.X(class(self).GetVisibleRatio()))
}

func (self Instance) SetVisibleRatio(value Float.X) {
	class(self).SetVisibleRatio(gd.Float(value))
}

func (self Instance) TextDirection() classdb.ControlTextDirection {
	return classdb.ControlTextDirection(class(self).GetTextDirection())
}

func (self Instance) SetTextDirection(value classdb.ControlTextDirection) {
	class(self).SetTextDirection(value)
}

func (self Instance) Language() string {
	return string(class(self).GetLanguage().String())
}

func (self Instance) SetLanguage(value string) {
	class(self).SetLanguage(gd.NewString(value))
}

func (self Instance) StructuredTextBidiOverride() classdb.TextServerStructuredTextParser {
	return classdb.TextServerStructuredTextParser(class(self).GetStructuredTextBidiOverride())
}

func (self Instance) SetStructuredTextBidiOverride(value classdb.TextServerStructuredTextParser) {
	class(self).SetStructuredTextBidiOverride(value)
}

func (self Instance) StructuredTextBidiOverrideOptions() Array.Any {
	return Array.Any(class(self).GetStructuredTextBidiOverrideOptions())
}

func (self Instance) SetStructuredTextBidiOverrideOptions(value Array.Any) {
	class(self).SetStructuredTextBidiOverrideOptions(value)
}

//go:nosplit
func (self class) SetHorizontalAlignment(alignment HorizontalAlignment) {
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHorizontalAlignment() HorizontalAlignment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[HorizontalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVerticalAlignment(alignment VerticalAlignment) {
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_vertical_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVerticalAlignment() VerticalAlignment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[VerticalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_vertical_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetText(text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetText() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLabelSettings(settings objects.LabelSettings) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(settings[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_label_settings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLabelSettings() objects.LabelSettings {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_label_settings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.LabelSettings{classdb.LabelSettings(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextDirection(direction classdb.ControlTextDirection) {
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextDirection() classdb.ControlTextDirection {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ControlTextDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLanguage(language gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(language))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLanguage() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutowrapMode(autowrap_mode classdb.TextServerAutowrapMode) {
	var frame = callframe.New()
	callframe.Arg(frame, autowrap_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutowrapMode() classdb.TextServerAutowrapMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerAutowrapMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJustificationFlags(justification_flags classdb.TextServerJustificationFlag) {
	var frame = callframe.New()
	callframe.Arg(frame, justification_flags)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_justification_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetJustificationFlags() classdb.TextServerJustificationFlag {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerJustificationFlag](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_justification_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetClipText(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_clip_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsClippingText() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_is_clipping_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTabStops(tab_stops gd.PackedFloat32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tab_stops))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_tab_stops, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTabStops() gd.PackedFloat32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_tab_stops, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedFloat32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextOverrunBehavior(overrun_behavior classdb.TextServerOverrunBehavior) {
	var frame = callframe.New()
	callframe.Arg(frame, overrun_behavior)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextOverrunBehavior() classdb.TextServerOverrunBehavior {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerOverrunBehavior](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEllipsisChar(char gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(char))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_ellipsis_char, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEllipsisChar() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_ellipsis_char, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUppercase(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_uppercase, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUppercase() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_is_uppercase, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) GetLineHeight(line gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_line_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of lines of text the Label has.
*/
//go:nosplit
func (self class) GetLineCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_line_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of lines shown. Useful if the [Label]'s height cannot currently display all lines.
*/
//go:nosplit
func (self class) GetVisibleLineCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_visible_line_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the total number of printable characters in the text (excluding spaces and newlines).
*/
//go:nosplit
func (self class) GetTotalCharacterCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_total_character_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibleCharacters(amount gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_visible_characters, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibleCharacters() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_visible_characters, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetVisibleCharactersBehavior() classdb.TextServerVisibleCharactersBehavior {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerVisibleCharactersBehavior](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_visible_characters_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibleCharactersBehavior(behavior classdb.TextServerVisibleCharactersBehavior) {
	var frame = callframe.New()
	callframe.Arg(frame, behavior)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_visible_characters_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetVisibleRatio(ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_visible_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibleRatio() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_visible_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinesSkipped(lines_skipped gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, lines_skipped)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_lines_skipped, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinesSkipped() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_lines_skipped, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxLinesVisible(lines_visible gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, lines_visible)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_max_lines_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxLinesVisible() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_max_lines_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStructuredTextBidiOverride(parser classdb.TextServerStructuredTextParser) {
	var frame = callframe.New()
	callframe.Arg(frame, parser)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetStructuredTextBidiOverride() classdb.TextServerStructuredTextParser {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerStructuredTextParser](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStructuredTextBidiOverrideOptions(args gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(args))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_set_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetStructuredTextBidiOverrideOptions() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the bounding rectangle of the character at position [param pos]. If the character is a non-visual character or [param pos] is outside the valid range, an empty [Rect2] is returned. If the character is a part of a composite grapheme, the bounding rectangle of the whole grapheme is returned.
*/
//go:nosplit
func (self class) GetCharacterBounds(pos gd.Int) gd.Rect2 {
	var frame = callframe.New()
	callframe.Arg(frame, pos)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label.Bind_get_character_bounds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
		return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() {
	classdb.Register("Label", func(ptr gd.Object) any { return [1]classdb.Label{classdb.Label(ptr)} })
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
