// Package FontVariation provides methods for working with FontVariation object instances.
package FontVariation

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Font"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Transform2D"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Provides OpenType variations, simulated bold / slant, and additional font settings like OpenType features and extra spacing.
To use simulated bold font variant:
[codeblocks]
[gdscript]
var fv = FontVariation.new()
fv.base_font = load("res://BarlowCondensed-Regular.ttf")
fv.variation_embolden = 1.2
$Label.add_theme_font_override("font", fv)
$Label.add_theme_font_size_override("font_size", 64)
[/gdscript]
[csharp]
var fv = new FontVariation();
fv.SetBaseFont(ResourceLoader.Load<FontFile>("res://BarlowCondensed-Regular.ttf"));
fv.SetVariationEmbolden(1.2);
GetNode("Label").AddThemeFontOverride("font", fv);
GetNode("Label").AddThemeFontSizeOverride("font_size", 64);
[/csharp]
[/codeblocks]
To set the coordinate of multiple variation axes:
[codeblock]
var fv = FontVariation.new();
var ts = TextServerManager.get_primary_interface()
fv.base_font = load("res://BarlowCondensed-Regular.ttf")
fv.variation_opentype = { ts.name_to_tag("wght"): 900, ts.name_to_tag("custom_hght"): 900 }
[/codeblock]
*/
type Instance [1]gdclass.FontVariation

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsFontVariation() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.FontVariation

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FontVariation"))
	casted := Instance{*(*gdclass.FontVariation)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) BaseFont() [1]gdclass.Font {
	return [1]gdclass.Font(class(self).GetBaseFont())
}

func (self Instance) SetBaseFont(value [1]gdclass.Font) {
	class(self).SetBaseFont(value)
}

func (self Instance) VariationOpentype() Dictionary.Any {
	return Dictionary.Any(class(self).GetVariationOpentype())
}

func (self Instance) SetVariationOpentype(value Dictionary.Any) {
	class(self).SetVariationOpentype(value)
}

func (self Instance) VariationFaceIndex() int {
	return int(int(class(self).GetVariationFaceIndex()))
}

func (self Instance) SetVariationFaceIndex(value int) {
	class(self).SetVariationFaceIndex(gd.Int(value))
}

func (self Instance) VariationEmbolden() Float.X {
	return Float.X(Float.X(class(self).GetVariationEmbolden()))
}

func (self Instance) SetVariationEmbolden(value Float.X) {
	class(self).SetVariationEmbolden(gd.Float(value))
}

func (self Instance) VariationTransform() Transform2D.OriginXY {
	return Transform2D.OriginXY(class(self).GetVariationTransform())
}

func (self Instance) SetVariationTransform(value Transform2D.OriginXY) {
	class(self).SetVariationTransform(gd.Transform2D(value))
}

func (self Instance) SetOpentypeFeatures(value Dictionary.Any) {
	class(self).SetOpentypeFeatures(value)
}

func (self Instance) SetSpacingGlyph(value int) {
	class(self).SetSpacing(0, gd.Int(value))
}

func (self Instance) SetSpacingSpace(value int) {
	class(self).SetSpacing(1, gd.Int(value))
}

func (self Instance) SetSpacingTop(value int) {
	class(self).SetSpacing(2, gd.Int(value))
}

func (self Instance) SetSpacingBottom(value int) {
	class(self).SetSpacing(3, gd.Int(value))
}

func (self Instance) BaselineOffset() Float.X {
	return Float.X(Float.X(class(self).GetBaselineOffset()))
}

func (self Instance) SetBaselineOffset(value Float.X) {
	class(self).SetBaselineOffset(gd.Float(value))
}

//go:nosplit
func (self class) SetBaseFont(font [1]gdclass.Font) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(font[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_set_base_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBaseFont() [1]gdclass.Font {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_get_base_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Font{gd.PointerWithOwnershipTransferredToGo[gdclass.Font](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVariationOpentype(coords gd.Dictionary) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(coords))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_set_variation_opentype, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVariationOpentype() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_get_variation_opentype, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVariationEmbolden(strength gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_set_variation_embolden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVariationEmbolden() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_get_variation_embolden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVariationFaceIndex(face_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, face_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_set_variation_face_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVariationFaceIndex() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_get_variation_face_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVariationTransform(transform gd.Transform2D) {
	var frame = callframe.New()
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_set_variation_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVariationTransform() gd.Transform2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_get_variation_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOpentypeFeatures(features gd.Dictionary) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(features))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_set_opentype_features, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the spacing for [param spacing] (see [enum TextServer.SpacingType]) to [param value] in pixels (not relative to the font size).
*/
//go:nosplit
func (self class) SetSpacing(spacing gdclass.TextServerSpacingType, value gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, spacing)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_set_spacing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetBaselineOffset(baseline_offset gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, baseline_offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_set_baseline_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBaselineOffset() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_get_baseline_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsFontVariation() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFontVariation() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsFont() Font.Advanced        { return *((*Font.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFont() Font.Instance     { return *((*Font.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Font.Advanced(self.AsFont()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Font.Instance(self.AsFont()), name)
	}
}
func init() {
	gdclass.Register("FontVariation", func(ptr gd.Object) any {
		return [1]gdclass.FontVariation{*(*gdclass.FontVariation)(unsafe.Pointer(&ptr))}
	})
}
