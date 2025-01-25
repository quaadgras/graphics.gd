// Package Label3D provides methods for working with Label3D object instances.
package Label3D

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
import "graphics.gd/classdb/GeometryInstance3D"
import "graphics.gd/classdb/VisualInstance3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any

/*
A node for displaying plain text in 3D space. By adjusting various properties of this node, you can configure things such as the text's appearance and whether it always faces the camera.
*/
type Instance [1]gdclass.Label3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsLabel3D() Instance
}

/*
Returns a [TriangleMesh] with the label's vertices following its current configuration (such as its [member pixel_size]).
*/
func (self Instance) GenerateTriangleMesh() [1]gdclass.TriangleMesh {
	return [1]gdclass.TriangleMesh(class(self).GenerateTriangleMesh())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Label3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Label3D"))
	casted := Instance{*(*gdclass.Label3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) PixelSize() Float.X {
	return Float.X(Float.X(class(self).GetPixelSize()))
}

func (self Instance) SetPixelSize(value Float.X) {
	class(self).SetPixelSize(gd.Float(value))
}

func (self Instance) Offset() Vector2.XY {
	return Vector2.XY(class(self).GetOffset())
}

func (self Instance) SetOffset(value Vector2.XY) {
	class(self).SetOffset(gd.Vector2(value))
}

func (self Instance) Billboard() gdclass.BaseMaterial3DBillboardMode {
	return gdclass.BaseMaterial3DBillboardMode(class(self).GetBillboardMode())
}

func (self Instance) SetBillboard(value gdclass.BaseMaterial3DBillboardMode) {
	class(self).SetBillboardMode(value)
}

func (self Instance) Shaded() bool {
	return bool(class(self).GetDrawFlag(0))
}

func (self Instance) SetShaded(value bool) {
	class(self).SetDrawFlag(0, value)
}

func (self Instance) DoubleSided() bool {
	return bool(class(self).GetDrawFlag(1))
}

func (self Instance) SetDoubleSided(value bool) {
	class(self).SetDrawFlag(1, value)
}

func (self Instance) NoDepthTest() bool {
	return bool(class(self).GetDrawFlag(2))
}

func (self Instance) SetNoDepthTest(value bool) {
	class(self).SetDrawFlag(2, value)
}

func (self Instance) FixedSize() bool {
	return bool(class(self).GetDrawFlag(3))
}

func (self Instance) SetFixedSize(value bool) {
	class(self).SetDrawFlag(3, value)
}

func (self Instance) AlphaCut() gdclass.Label3DAlphaCutMode {
	return gdclass.Label3DAlphaCutMode(class(self).GetAlphaCutMode())
}

func (self Instance) SetAlphaCut(value gdclass.Label3DAlphaCutMode) {
	class(self).SetAlphaCutMode(value)
}

func (self Instance) AlphaScissorThreshold() Float.X {
	return Float.X(Float.X(class(self).GetAlphaScissorThreshold()))
}

func (self Instance) SetAlphaScissorThreshold(value Float.X) {
	class(self).SetAlphaScissorThreshold(gd.Float(value))
}

func (self Instance) AlphaHashScale() Float.X {
	return Float.X(Float.X(class(self).GetAlphaHashScale()))
}

func (self Instance) SetAlphaHashScale(value Float.X) {
	class(self).SetAlphaHashScale(gd.Float(value))
}

func (self Instance) AlphaAntialiasingMode() gdclass.BaseMaterial3DAlphaAntiAliasing {
	return gdclass.BaseMaterial3DAlphaAntiAliasing(class(self).GetAlphaAntialiasing())
}

func (self Instance) SetAlphaAntialiasingMode(value gdclass.BaseMaterial3DAlphaAntiAliasing) {
	class(self).SetAlphaAntialiasing(value)
}

func (self Instance) AlphaAntialiasingEdge() Float.X {
	return Float.X(Float.X(class(self).GetAlphaAntialiasingEdge()))
}

func (self Instance) SetAlphaAntialiasingEdge(value Float.X) {
	class(self).SetAlphaAntialiasingEdge(gd.Float(value))
}

func (self Instance) TextureFilter() gdclass.BaseMaterial3DTextureFilter {
	return gdclass.BaseMaterial3DTextureFilter(class(self).GetTextureFilter())
}

func (self Instance) SetTextureFilter(value gdclass.BaseMaterial3DTextureFilter) {
	class(self).SetTextureFilter(value)
}

func (self Instance) RenderPriority() int {
	return int(int(class(self).GetRenderPriority()))
}

func (self Instance) SetRenderPriority(value int) {
	class(self).SetRenderPriority(gd.Int(value))
}

func (self Instance) OutlineRenderPriority() int {
	return int(int(class(self).GetOutlineRenderPriority()))
}

func (self Instance) SetOutlineRenderPriority(value int) {
	class(self).SetOutlineRenderPriority(gd.Int(value))
}

func (self Instance) Modulate() Color.RGBA {
	return Color.RGBA(class(self).GetModulate())
}

func (self Instance) SetModulate(value Color.RGBA) {
	class(self).SetModulate(gd.Color(value))
}

func (self Instance) OutlineModulate() Color.RGBA {
	return Color.RGBA(class(self).GetOutlineModulate())
}

func (self Instance) SetOutlineModulate(value Color.RGBA) {
	class(self).SetOutlineModulate(gd.Color(value))
}

func (self Instance) Text() string {
	return string(class(self).GetText().String())
}

func (self Instance) SetText(value string) {
	class(self).SetText(gd.NewString(value))
}

func (self Instance) Font() [1]gdclass.Font {
	return [1]gdclass.Font(class(self).GetFont())
}

func (self Instance) SetFont(value [1]gdclass.Font) {
	class(self).SetFont(value)
}

func (self Instance) FontSize() int {
	return int(int(class(self).GetFontSize()))
}

func (self Instance) SetFontSize(value int) {
	class(self).SetFontSize(gd.Int(value))
}

func (self Instance) OutlineSize() int {
	return int(int(class(self).GetOutlineSize()))
}

func (self Instance) SetOutlineSize(value int) {
	class(self).SetOutlineSize(gd.Int(value))
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

func (self Instance) Uppercase() bool {
	return bool(class(self).IsUppercase())
}

func (self Instance) SetUppercase(value bool) {
	class(self).SetUppercase(value)
}

func (self Instance) LineSpacing() Float.X {
	return Float.X(Float.X(class(self).GetLineSpacing()))
}

func (self Instance) SetLineSpacing(value Float.X) {
	class(self).SetLineSpacing(gd.Float(value))
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

func (self Instance) Width() Float.X {
	return Float.X(Float.X(class(self).GetWidth()))
}

func (self Instance) SetWidth(value Float.X) {
	class(self).SetWidth(gd.Float(value))
}

func (self Instance) TextDirection() gdclass.TextServerDirection {
	return gdclass.TextServerDirection(class(self).GetTextDirection())
}

func (self Instance) SetTextDirection(value gdclass.TextServerDirection) {
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
func (self class) SetHorizontalAlignment(alignment HorizontalAlignment) {
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHorizontalAlignment() HorizontalAlignment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[HorizontalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVerticalAlignment(alignment VerticalAlignment) {
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_vertical_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVerticalAlignment() VerticalAlignment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[VerticalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_vertical_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetModulate(modulate gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, modulate)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetModulate() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOutlineModulate(modulate gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, modulate)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_outline_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOutlineModulate() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_outline_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetText(text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetText() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextDirection(direction gdclass.TextServerDirection) {
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextDirection() gdclass.TextServerDirection {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLanguage(language gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(language))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLanguage() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStructuredTextBidiOverride(parser gdclass.TextServerStructuredTextParser) {
	var frame = callframe.New()
	callframe.Arg(frame, parser)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStructuredTextBidiOverride() gdclass.TextServerStructuredTextParser {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerStructuredTextParser](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStructuredTextBidiOverrideOptions(args Array.Any) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(args)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStructuredTextBidiOverrideOptions() Array.Any {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUppercase(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_uppercase, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUppercase() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_is_uppercase, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRenderPriority(priority gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_render_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRenderPriority() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_render_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOutlineRenderPriority(priority gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_outline_render_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOutlineRenderPriority() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_outline_render_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFont(font [1]gdclass.Font) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(font[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFont() [1]gdclass.Font {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Font{gd.PointerWithOwnershipTransferredToGo[gdclass.Font](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFontSize(size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFontSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOutlineSize(outline_size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, outline_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_outline_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOutlineSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_outline_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLineSpacing(line_spacing gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, line_spacing)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_line_spacing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLineSpacing() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_line_spacing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutowrapMode(autowrap_mode gdclass.TextServerAutowrapMode) {
	var frame = callframe.New()
	callframe.Arg(frame, autowrap_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutowrapMode() gdclass.TextServerAutowrapMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerAutowrapMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJustificationFlags(justification_flags gdclass.TextServerJustificationFlag) {
	var frame = callframe.New()
	callframe.Arg(frame, justification_flags)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_justification_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetJustificationFlags() gdclass.TextServerJustificationFlag {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerJustificationFlag](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_justification_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWidth(width gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetWidth() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPixelSize(pixel_size gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pixel_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_pixel_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPixelSize() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_pixel_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOffset(offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [code]true[/code], the specified flag will be enabled. See [enum Label3D.DrawFlags] for a list of flags.
*/
//go:nosplit
func (self class) SetDrawFlag(flag gdclass.Label3DDrawFlags, enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_draw_flag, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the value of the specified flag.
*/
//go:nosplit
func (self class) GetDrawFlag(flag gdclass.Label3DDrawFlags) bool {
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_draw_flag, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBillboardMode(mode gdclass.BaseMaterial3DBillboardMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_billboard_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBillboardMode() gdclass.BaseMaterial3DBillboardMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DBillboardMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_billboard_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlphaCutMode(mode gdclass.Label3DAlphaCutMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_alpha_cut_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlphaCutMode() gdclass.Label3DAlphaCutMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Label3DAlphaCutMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_alpha_cut_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlphaScissorThreshold(threshold gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, threshold)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_alpha_scissor_threshold, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlphaScissorThreshold() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_alpha_scissor_threshold, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlphaHashScale(threshold gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, threshold)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_alpha_hash_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlphaHashScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_alpha_hash_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlphaAntialiasing(alpha_aa gdclass.BaseMaterial3DAlphaAntiAliasing) {
	var frame = callframe.New()
	callframe.Arg(frame, alpha_aa)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_alpha_antialiasing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlphaAntialiasing() gdclass.BaseMaterial3DAlphaAntiAliasing {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DAlphaAntiAliasing](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_alpha_antialiasing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlphaAntialiasingEdge(edge gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, edge)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_alpha_antialiasing_edge, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlphaAntialiasingEdge() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_alpha_antialiasing_edge, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureFilter(mode gdclass.BaseMaterial3DTextureFilter) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_set_texture_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureFilter() gdclass.BaseMaterial3DTextureFilter {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BaseMaterial3DTextureFilter](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_get_texture_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [TriangleMesh] with the label's vertices following its current configuration (such as its [member pixel_size]).
*/
//go:nosplit
func (self class) GenerateTriangleMesh() [1]gdclass.TriangleMesh {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Label3D.Bind_generate_triangle_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TriangleMesh{gd.PointerWithOwnershipTransferredToGo[gdclass.TriangleMesh](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsLabel3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsLabel3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsGeometryInstance3D() GeometryInstance3D.Advanced {
	return *((*GeometryInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGeometryInstance3D() GeometryInstance3D.Instance {
	return *((*GeometryInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualInstance3D() VisualInstance3D.Advanced {
	return *((*VisualInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualInstance3D() VisualInstance3D.Instance {
	return *((*VisualInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(GeometryInstance3D.Advanced(self.AsGeometryInstance3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(GeometryInstance3D.Instance(self.AsGeometryInstance3D()), name)
	}
}
func init() {
	gdclass.Register("Label3D", func(ptr gd.Object) any { return [1]gdclass.Label3D{*(*gdclass.Label3D)(unsafe.Pointer(&ptr))} })
}

type DrawFlags = gdclass.Label3DDrawFlags

const (
	/*If set, lights in the environment affect the label.*/
	FlagShaded DrawFlags = 0
	/*If set, text can be seen from the back as well. If not, the text is invisible when looking at it from behind.*/
	FlagDoubleSided DrawFlags = 1
	/*Disables the depth test, so this object is drawn on top of all others. However, objects drawn after it in the draw order may cover it.*/
	FlagDisableDepthTest DrawFlags = 2
	/*Label is scaled by depth so that it always appears the same size on screen.*/
	FlagFixedSize DrawFlags = 3
	/*Represents the size of the [enum DrawFlags] enum.*/
	FlagMax DrawFlags = 4
)

type AlphaCutMode = gdclass.Label3DAlphaCutMode

const (
	/*This mode performs standard alpha blending. It can display translucent areas, but transparency sorting issues may be visible when multiple transparent materials are overlapping. [member GeometryInstance3D.cast_shadow] has no effect when this transparency mode is used; the [Label3D] will never cast shadows.*/
	AlphaCutDisabled AlphaCutMode = 0
	/*This mode only allows fully transparent or fully opaque pixels. Harsh edges will be visible unless some form of screen-space antialiasing is enabled (see [member ProjectSettings.rendering/anti_aliasing/quality/screen_space_aa]). This mode is also known as [i]alpha testing[/i] or [i]1-bit transparency[/i].
	  [b]Note:[/b] This mode might have issues with anti-aliased fonts and outlines, try adjusting [member alpha_scissor_threshold] or using MSDF font.
	  [b]Note:[/b] When using text with overlapping glyphs (e.g., cursive scripts), this mode might have transparency sorting issues between the main text and the outline.*/
	AlphaCutDiscard AlphaCutMode = 1
	/*This mode draws fully opaque pixels in the depth prepass. This is slower than [constant ALPHA_CUT_DISABLED] or [constant ALPHA_CUT_DISCARD], but it allows displaying translucent areas and smooth edges while using proper sorting.
	  [b]Note:[/b] When using text with overlapping glyphs (e.g., cursive scripts), this mode might have transparency sorting issues between the main text and the outline.*/
	AlphaCutOpaquePrepass AlphaCutMode = 2
	/*This mode draws cuts off all values below a spatially-deterministic threshold, the rest will remain opaque.*/
	AlphaCutHash AlphaCutMode = 3
)

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
