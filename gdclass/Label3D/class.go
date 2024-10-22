package Label3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/GeometryInstance3D"
import "grow.graphics/gd/gdclass/VisualInstance3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A node for displaying plain text in 3D space. By adjusting various properties of this node, you can configure things such as the text's appearance and whether it always faces the camera.

*/
type Go [1]classdb.Label3D

/*
Returns a [TriangleMesh] with the label's vertices following its current configuration (such as its [member pixel_size]).
*/
func (self Go) GenerateTriangleMesh() gdclass.TriangleMesh {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.TriangleMesh(class(self).GenerateTriangleMesh(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Label3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Label3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) PixelSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetPixelSize()))
}

func (self Go) SetPixelSize(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPixelSize(gd.Float(value))
}

func (self Go) Offset() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetOffset())
}

func (self Go) SetOffset(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOffset(value)
}

func (self Go) Billboard() classdb.BaseMaterial3DBillboardMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.BaseMaterial3DBillboardMode(class(self).GetBillboardMode())
}

func (self Go) SetBillboard(value classdb.BaseMaterial3DBillboardMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBillboardMode(value)
}

func (self Go) Shaded() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetDrawFlag(0))
}

func (self Go) SetShaded(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawFlag(0, value)
}

func (self Go) DoubleSided() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetDrawFlag(1))
}

func (self Go) SetDoubleSided(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawFlag(1, value)
}

func (self Go) NoDepthTest() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetDrawFlag(2))
}

func (self Go) SetNoDepthTest(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawFlag(2, value)
}

func (self Go) FixedSize() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetDrawFlag(3))
}

func (self Go) SetFixedSize(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawFlag(3, value)
}

func (self Go) AlphaCut() classdb.Label3DAlphaCutMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.Label3DAlphaCutMode(class(self).GetAlphaCutMode())
}

func (self Go) SetAlphaCut(value classdb.Label3DAlphaCutMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAlphaCutMode(value)
}

func (self Go) AlphaScissorThreshold() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAlphaScissorThreshold()))
}

func (self Go) SetAlphaScissorThreshold(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAlphaScissorThreshold(gd.Float(value))
}

func (self Go) AlphaHashScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAlphaHashScale()))
}

func (self Go) SetAlphaHashScale(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAlphaHashScale(gd.Float(value))
}

func (self Go) AlphaAntialiasingMode() classdb.BaseMaterial3DAlphaAntiAliasing {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.BaseMaterial3DAlphaAntiAliasing(class(self).GetAlphaAntialiasing())
}

func (self Go) SetAlphaAntialiasingMode(value classdb.BaseMaterial3DAlphaAntiAliasing) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAlphaAntialiasing(value)
}

func (self Go) AlphaAntialiasingEdge() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAlphaAntialiasingEdge()))
}

func (self Go) SetAlphaAntialiasingEdge(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAlphaAntialiasingEdge(gd.Float(value))
}

func (self Go) TextureFilter() classdb.BaseMaterial3DTextureFilter {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.BaseMaterial3DTextureFilter(class(self).GetTextureFilter())
}

func (self Go) SetTextureFilter(value classdb.BaseMaterial3DTextureFilter) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextureFilter(value)
}

func (self Go) RenderPriority() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetRenderPriority()))
}

func (self Go) SetRenderPriority(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRenderPriority(gd.Int(value))
}

func (self Go) OutlineRenderPriority() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetOutlineRenderPriority()))
}

func (self Go) SetOutlineRenderPriority(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOutlineRenderPriority(gd.Int(value))
}

func (self Go) Modulate() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetModulate())
}

func (self Go) SetModulate(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetModulate(value)
}

func (self Go) OutlineModulate() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetOutlineModulate())
}

func (self Go) SetOutlineModulate(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOutlineModulate(value)
}

func (self Go) Text() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetText(gc).String())
}

func (self Go) SetText(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetText(gc.String(value))
}

func (self Go) Font() gdclass.Font {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Font(class(self).GetFont(gc))
}

func (self Go) SetFont(value gdclass.Font) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFont(value)
}

func (self Go) FontSize() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetFontSize()))
}

func (self Go) SetFontSize(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFontSize(gd.Int(value))
}

func (self Go) OutlineSize() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetOutlineSize()))
}

func (self Go) SetOutlineSize(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOutlineSize(gd.Int(value))
}

func (self Go) HorizontalAlignment() gd.HorizontalAlignment {
	gc := gd.GarbageCollector(); _ = gc
		return gd.HorizontalAlignment(class(self).GetHorizontalAlignment())
}

func (self Go) SetHorizontalAlignment(value gd.HorizontalAlignment) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHorizontalAlignment(value)
}

func (self Go) VerticalAlignment() gd.VerticalAlignment {
	gc := gd.GarbageCollector(); _ = gc
		return gd.VerticalAlignment(class(self).GetVerticalAlignment())
}

func (self Go) SetVerticalAlignment(value gd.VerticalAlignment) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVerticalAlignment(value)
}

func (self Go) Uppercase() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsUppercase())
}

func (self Go) SetUppercase(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUppercase(value)
}

func (self Go) LineSpacing() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetLineSpacing()))
}

func (self Go) SetLineSpacing(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLineSpacing(gd.Float(value))
}

func (self Go) AutowrapMode() classdb.TextServerAutowrapMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextServerAutowrapMode(class(self).GetAutowrapMode())
}

func (self Go) SetAutowrapMode(value classdb.TextServerAutowrapMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutowrapMode(value)
}

func (self Go) JustificationFlags() classdb.TextServerJustificationFlag {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextServerJustificationFlag(class(self).GetJustificationFlags())
}

func (self Go) SetJustificationFlags(value classdb.TextServerJustificationFlag) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJustificationFlags(value)
}

func (self Go) Width() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetWidth()))
}

func (self Go) SetWidth(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetWidth(gd.Float(value))
}

func (self Go) TextDirection() classdb.TextServerDirection {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextServerDirection(class(self).GetTextDirection())
}

func (self Go) SetTextDirection(value classdb.TextServerDirection) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextDirection(value)
}

func (self Go) Language() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetLanguage(gc).String())
}

func (self Go) SetLanguage(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLanguage(gc.String(value))
}

func (self Go) StructuredTextBidiOverride() classdb.TextServerStructuredTextParser {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextServerStructuredTextParser(class(self).GetStructuredTextBidiOverride())
}

func (self Go) SetStructuredTextBidiOverride(value classdb.TextServerStructuredTextParser) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStructuredTextBidiOverride(value)
}

func (self Go) StructuredTextBidiOverrideOptions() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Array(class(self).GetStructuredTextBidiOverrideOptions(gc))
}

func (self Go) SetStructuredTextBidiOverrideOptions(value gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStructuredTextBidiOverrideOptions(value)
}

//go:nosplit
func (self class) SetHorizontalAlignment(alignment gd.HorizontalAlignment)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHorizontalAlignment() gd.HorizontalAlignment {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.HorizontalAlignment](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVerticalAlignment(alignment gd.VerticalAlignment)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_vertical_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVerticalAlignment() gd.VerticalAlignment {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.VerticalAlignment](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_vertical_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetModulate(modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetModulate() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOutlineModulate(modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_outline_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOutlineModulate() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_outline_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetText(text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetText(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextDirection(direction classdb.TextServerDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextDirection() classdb.TextServerDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLanguage(language gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(language))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLanguage(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStructuredTextBidiOverride(parser classdb.TextServerStructuredTextParser)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, parser)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStructuredTextBidiOverride() classdb.TextServerStructuredTextParser {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerStructuredTextParser](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStructuredTextBidiOverrideOptions(args gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(args))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStructuredTextBidiOverrideOptions(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUppercase(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_uppercase, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUppercase() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_is_uppercase, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRenderPriority(priority gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_render_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRenderPriority() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_render_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOutlineRenderPriority(priority gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_outline_render_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOutlineRenderPriority() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_outline_render_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFont(font gdclass.Font)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFont(ctx gd.Lifetime) gdclass.Font {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Font
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFontSize(size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFontSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOutlineSize(outline_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, outline_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_outline_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOutlineSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_outline_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLineSpacing(line_spacing gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line_spacing)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_line_spacing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLineSpacing() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_line_spacing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutowrapMode(autowrap_mode classdb.TextServerAutowrapMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, autowrap_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAutowrapMode() classdb.TextServerAutowrapMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerAutowrapMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetJustificationFlags(justification_flags classdb.TextServerJustificationFlag)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, justification_flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_justification_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetJustificationFlags() classdb.TextServerJustificationFlag {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerJustificationFlag](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_justification_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWidth() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPixelSize(pixel_size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pixel_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_pixel_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPixelSize() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_pixel_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOffset(offset gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOffset() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], the specified flag will be enabled. See [enum Label3D.DrawFlags] for a list of flags.
*/
//go:nosplit
func (self class) SetDrawFlag(flag classdb.Label3DDrawFlags, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_draw_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the specified flag.
*/
//go:nosplit
func (self class) GetDrawFlag(flag classdb.Label3DDrawFlags) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_draw_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBillboardMode(mode classdb.BaseMaterial3DBillboardMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_billboard_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBillboardMode() classdb.BaseMaterial3DBillboardMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DBillboardMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_billboard_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlphaCutMode(mode classdb.Label3DAlphaCutMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_alpha_cut_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlphaCutMode() classdb.Label3DAlphaCutMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Label3DAlphaCutMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_alpha_cut_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlphaScissorThreshold(threshold gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, threshold)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_alpha_scissor_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlphaScissorThreshold() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_alpha_scissor_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlphaHashScale(threshold gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, threshold)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_alpha_hash_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlphaHashScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_alpha_hash_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlphaAntialiasing(alpha_aa classdb.BaseMaterial3DAlphaAntiAliasing)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, alpha_aa)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_alpha_antialiasing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlphaAntialiasing() classdb.BaseMaterial3DAlphaAntiAliasing {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DAlphaAntiAliasing](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_alpha_antialiasing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlphaAntialiasingEdge(edge gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, edge)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_alpha_antialiasing_edge, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlphaAntialiasingEdge() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_alpha_antialiasing_edge, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextureFilter(mode classdb.BaseMaterial3DTextureFilter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_set_texture_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextureFilter() classdb.BaseMaterial3DTextureFilter {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseMaterial3DTextureFilter](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_get_texture_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [TriangleMesh] with the label's vertices following its current configuration (such as its [member pixel_size]).
*/
//go:nosplit
func (self class) GenerateTriangleMesh(ctx gd.Lifetime) gdclass.TriangleMesh {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Label3D.Bind_generate_triangle_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.TriangleMesh
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsLabel3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsLabel3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsGeometryInstance3D() GeometryInstance3D.GD { return *((*GeometryInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsGeometryInstance3D() GeometryInstance3D.Go { return *((*GeometryInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualInstance3D() VisualInstance3D.GD { return *((*VisualInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualInstance3D() VisualInstance3D.Go { return *((*VisualInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsGeometryInstance3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsGeometryInstance3D(), name)
	}
}
func init() {classdb.Register("Label3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type DrawFlags = classdb.Label3DDrawFlags

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
type AlphaCutMode = classdb.Label3DAlphaCutMode

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
