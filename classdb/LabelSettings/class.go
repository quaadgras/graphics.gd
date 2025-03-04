// Package LabelSettings provides methods for working with LabelSettings object instances.
package LabelSettings

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
[LabelSettings] is a resource that provides common settings to customize the text in a [Label]. It will take priority over the properties defined in [member Control.theme]. The resource can be shared between multiple labels and changed on the fly, so it's convenient and flexible way to setup text style.
*/
type Instance [1]gdclass.LabelSettings

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsLabelSettings() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.LabelSettings

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("LabelSettings"))
	casted := Instance{*(*gdclass.LabelSettings)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) LineSpacing() Float.X {
	return Float.X(Float.X(class(self).GetLineSpacing()))
}

func (self Instance) SetLineSpacing(value Float.X) {
	class(self).SetLineSpacing(float64(value))
}

func (self Instance) ParagraphSpacing() Float.X {
	return Float.X(Float.X(class(self).GetParagraphSpacing()))
}

func (self Instance) SetParagraphSpacing(value Float.X) {
	class(self).SetParagraphSpacing(float64(value))
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
	class(self).SetFontSize(int64(value))
}

func (self Instance) FontColor() Color.RGBA {
	return Color.RGBA(class(self).GetFontColor())
}

func (self Instance) SetFontColor(value Color.RGBA) {
	class(self).SetFontColor(Color.RGBA(value))
}

func (self Instance) OutlineSize() int {
	return int(int(class(self).GetOutlineSize()))
}

func (self Instance) SetOutlineSize(value int) {
	class(self).SetOutlineSize(int64(value))
}

func (self Instance) OutlineColor() Color.RGBA {
	return Color.RGBA(class(self).GetOutlineColor())
}

func (self Instance) SetOutlineColor(value Color.RGBA) {
	class(self).SetOutlineColor(Color.RGBA(value))
}

func (self Instance) ShadowSize() int {
	return int(int(class(self).GetShadowSize()))
}

func (self Instance) SetShadowSize(value int) {
	class(self).SetShadowSize(int64(value))
}

func (self Instance) ShadowColor() Color.RGBA {
	return Color.RGBA(class(self).GetShadowColor())
}

func (self Instance) SetShadowColor(value Color.RGBA) {
	class(self).SetShadowColor(Color.RGBA(value))
}

func (self Instance) ShadowOffset() Vector2.XY {
	return Vector2.XY(class(self).GetShadowOffset())
}

func (self Instance) SetShadowOffset(value Vector2.XY) {
	class(self).SetShadowOffset(Vector2.XY(value))
}

//go:nosplit
func (self class) SetLineSpacing(spacing float64) { //gd:LabelSettings.set_line_spacing
	var frame = callframe.New()
	callframe.Arg(frame, spacing)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_set_line_spacing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLineSpacing() float64 { //gd:LabelSettings.get_line_spacing
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_get_line_spacing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetParagraphSpacing(spacing float64) { //gd:LabelSettings.set_paragraph_spacing
	var frame = callframe.New()
	callframe.Arg(frame, spacing)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_set_paragraph_spacing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetParagraphSpacing() float64 { //gd:LabelSettings.get_paragraph_spacing
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_get_paragraph_spacing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFont(font [1]gdclass.Font) { //gd:LabelSettings.set_font
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(font[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_set_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFont() [1]gdclass.Font { //gd:LabelSettings.get_font
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_get_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Font{gd.PointerWithOwnershipTransferredToGo[gdclass.Font](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFontSize(size int64) { //gd:LabelSettings.set_font_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_set_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFontSize() int64 { //gd:LabelSettings.get_font_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_get_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFontColor(color Color.RGBA) { //gd:LabelSettings.set_font_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_set_font_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFontColor() Color.RGBA { //gd:LabelSettings.get_font_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_get_font_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOutlineSize(size int64) { //gd:LabelSettings.set_outline_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_set_outline_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOutlineSize() int64 { //gd:LabelSettings.get_outline_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_get_outline_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOutlineColor(color Color.RGBA) { //gd:LabelSettings.set_outline_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_set_outline_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOutlineColor() Color.RGBA { //gd:LabelSettings.get_outline_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_get_outline_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShadowSize(size int64) { //gd:LabelSettings.set_shadow_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_set_shadow_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShadowSize() int64 { //gd:LabelSettings.get_shadow_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_get_shadow_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShadowColor(color Color.RGBA) { //gd:LabelSettings.set_shadow_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_set_shadow_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShadowColor() Color.RGBA { //gd:LabelSettings.get_shadow_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_get_shadow_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShadowOffset(offset Vector2.XY) { //gd:LabelSettings.set_shadow_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_set_shadow_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShadowOffset() Vector2.XY { //gd:LabelSettings.get_shadow_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LabelSettings.Bind_get_shadow_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsLabelSettings() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsLabelSettings() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("LabelSettings", func(ptr gd.Object) any {
		return [1]gdclass.LabelSettings{*(*gdclass.LabelSettings)(unsafe.Pointer(&ptr))}
	})
}
