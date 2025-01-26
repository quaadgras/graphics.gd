// Package CharFXTransform provides methods for working with CharFXTransform object instances.
package CharFXTransform

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
import "graphics.gd/variant/Transform2D"
import "graphics.gd/variant/Vector2i"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/RID"

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
By setting various properties on this object, you can control how individual characters will be displayed in a [RichTextEffect].
*/
type Instance [1]gdclass.CharFXTransform

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCharFXTransform() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CharFXTransform

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CharFXTransform"))
	casted := Instance{*(*gdclass.CharFXTransform)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Transform() Transform2D.OriginXY {
	return Transform2D.OriginXY(class(self).GetTransform())
}

func (self Instance) SetTransform(value Transform2D.OriginXY) {
	class(self).SetTransform(gd.Transform2D(value))
}

func (self Instance) Range() Vector2i.XY {
	return Vector2i.XY(class(self).GetRange())
}

func (self Instance) SetRange(value Vector2i.XY) {
	class(self).SetRange(gd.Vector2i(value))
}

func (self Instance) ElapsedTime() Float.X {
	return Float.X(Float.X(class(self).GetElapsedTime()))
}

func (self Instance) SetElapsedTime(value Float.X) {
	class(self).SetElapsedTime(gd.Float(value))
}

func (self Instance) Visible() bool {
	return bool(class(self).IsVisible())
}

func (self Instance) SetVisible(value bool) {
	class(self).SetVisibility(value)
}

func (self Instance) Outline() bool {
	return bool(class(self).IsOutline())
}

func (self Instance) SetOutline(value bool) {
	class(self).SetOutline(value)
}

func (self Instance) Offset() Vector2.XY {
	return Vector2.XY(class(self).GetOffset())
}

func (self Instance) SetOffset(value Vector2.XY) {
	class(self).SetOffset(gd.Vector2(value))
}

func (self Instance) Color() Color.RGBA {
	return Color.RGBA(class(self).GetColor())
}

func (self Instance) SetColor(value Color.RGBA) {
	class(self).SetColor(gd.Color(value))
}

func (self Instance) Env() map[any]any {
	return map[any]any(gd.DictionaryAs[any, any](class(self).GetEnvironment()))
}

func (self Instance) SetEnv(value map[any]any) {
	class(self).SetEnvironment(gd.NewVariant(value).Interface().(gd.Dictionary))
}

func (self Instance) GlyphIndex() int {
	return int(int(class(self).GetGlyphIndex()))
}

func (self Instance) SetGlyphIndex(value int) {
	class(self).SetGlyphIndex(gd.Int(value))
}

func (self Instance) GlyphCount() int {
	return int(int(class(self).GetGlyphCount()))
}

func (self Instance) SetGlyphCount(value int) {
	class(self).SetGlyphCount(gd.Int(value))
}

func (self Instance) GlyphFlags() int {
	return int(int(class(self).GetGlyphFlags()))
}

func (self Instance) SetGlyphFlags(value int) {
	class(self).SetGlyphFlags(gd.Int(value))
}

func (self Instance) RelativeIndex() int {
	return int(int(class(self).GetRelativeIndex()))
}

func (self Instance) SetRelativeIndex(value int) {
	class(self).SetRelativeIndex(gd.Int(value))
}

func (self Instance) Font() RID.Any {
	return RID.Any(class(self).GetFont())
}

func (self Instance) SetFont(value RID.Any) {
	class(self).SetFont(gd.RID(value))
}

//go:nosplit
func (self class) GetTransform() gd.Transform2D { //gd:CharFXTransform.get_transform
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTransform(transform gd.Transform2D) { //gd:CharFXTransform.set_transform
	var frame = callframe.New()
	callframe.Arg(frame, transform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRange() gd.Vector2i { //gd:CharFXTransform.get_range
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRange(arange gd.Vector2i) { //gd:CharFXTransform.set_range
	var frame = callframe.New()
	callframe.Arg(frame, arange)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetElapsedTime() gd.Float { //gd:CharFXTransform.get_elapsed_time
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_elapsed_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetElapsedTime(time gd.Float) { //gd:CharFXTransform.set_elapsed_time
	var frame = callframe.New()
	callframe.Arg(frame, time)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_elapsed_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsVisible() bool { //gd:CharFXTransform.is_visible
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_is_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibility(visibility bool) { //gd:CharFXTransform.set_visibility
	var frame = callframe.New()
	callframe.Arg(frame, visibility)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_visibility, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsOutline() bool { //gd:CharFXTransform.is_outline
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_is_outline, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOutline(outline bool) { //gd:CharFXTransform.set_outline
	var frame = callframe.New()
	callframe.Arg(frame, outline)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_outline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOffset() gd.Vector2 { //gd:CharFXTransform.get_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOffset(offset gd.Vector2) { //gd:CharFXTransform.set_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetColor() gd.Color { //gd:CharFXTransform.get_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColor(color gd.Color) { //gd:CharFXTransform.set_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnvironment() gd.Dictionary { //gd:CharFXTransform.get_environment
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_environment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnvironment(environment gd.Dictionary) { //gd:CharFXTransform.set_environment
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(environment))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_environment, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlyphIndex() gd.Int { //gd:CharFXTransform.get_glyph_index
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_glyph_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlyphIndex(glyph_index gd.Int) { //gd:CharFXTransform.set_glyph_index
	var frame = callframe.New()
	callframe.Arg(frame, glyph_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_glyph_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRelativeIndex() gd.Int { //gd:CharFXTransform.get_relative_index
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_relative_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRelativeIndex(relative_index gd.Int) { //gd:CharFXTransform.set_relative_index
	var frame = callframe.New()
	callframe.Arg(frame, relative_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_relative_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlyphCount() gd.Int { //gd:CharFXTransform.get_glyph_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_glyph_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlyphCount(glyph_count gd.Int) { //gd:CharFXTransform.set_glyph_count
	var frame = callframe.New()
	callframe.Arg(frame, glyph_count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_glyph_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlyphFlags() gd.Int { //gd:CharFXTransform.get_glyph_flags
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_glyph_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlyphFlags(glyph_flags gd.Int) { //gd:CharFXTransform.set_glyph_flags
	var frame = callframe.New()
	callframe.Arg(frame, glyph_flags)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_glyph_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFont() gd.RID { //gd:CharFXTransform.get_font
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFont(font gd.RID) { //gd:CharFXTransform.set_font
	var frame = callframe.New()
	callframe.Arg(frame, font)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsCharFXTransform() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCharFXTransform() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("CharFXTransform", func(ptr gd.Object) any {
		return [1]gdclass.CharFXTransform{*(*gdclass.CharFXTransform)(unsafe.Pointer(&ptr))}
	})
}
