package FontVariation

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Font"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Go [1]classdb.FontVariation
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.FontVariation
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("FontVariation"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) BaseFont() gdclass.Font {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Font(class(self).GetBaseFont(gc))
}

func (self Go) SetBaseFont(value gdclass.Font) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBaseFont(value)
}

func (self Go) VariationOpentype() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Dictionary(class(self).GetVariationOpentype(gc))
}

func (self Go) SetVariationOpentype(value gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVariationOpentype(value)
}

func (self Go) VariationFaceIndex() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetVariationFaceIndex()))
}

func (self Go) SetVariationFaceIndex(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVariationFaceIndex(gd.Int(value))
}

func (self Go) VariationEmbolden() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVariationEmbolden()))
}

func (self Go) SetVariationEmbolden(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVariationEmbolden(gd.Float(value))
}

func (self Go) VariationTransform() gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Transform2D(class(self).GetVariationTransform())
}

func (self Go) SetVariationTransform(value gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVariationTransform(value)
}

func (self Go) BaselineOffset() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetBaselineOffset()))
}

func (self Go) SetBaselineOffset(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBaselineOffset(gd.Float(value))
}

//go:nosplit
func (self class) SetBaseFont(font gdclass.Font)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontVariation.Bind_set_base_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBaseFont(ctx gd.Lifetime) gdclass.Font {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontVariation.Bind_get_base_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Font
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVariationOpentype(coords gd.Dictionary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(coords))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontVariation.Bind_set_variation_opentype, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVariationOpentype(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontVariation.Bind_get_variation_opentype, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVariationEmbolden(strength gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontVariation.Bind_set_variation_embolden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVariationEmbolden() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontVariation.Bind_get_variation_embolden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVariationFaceIndex(face_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, face_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontVariation.Bind_set_variation_face_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVariationFaceIndex() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontVariation.Bind_get_variation_face_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVariationTransform(transform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontVariation.Bind_set_variation_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVariationTransform() gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontVariation.Bind_get_variation_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOpentypeFeatures(features gd.Dictionary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(features))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontVariation.Bind_set_opentype_features, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the spacing for [param spacing] (see [enum TextServer.SpacingType]) to [param value] in pixels (not relative to the font size).
*/
//go:nosplit
func (self class) SetSpacing(spacing classdb.TextServerSpacingType, value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, spacing)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontVariation.Bind_set_spacing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetBaselineOffset(baseline_offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, baseline_offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontVariation.Bind_set_baseline_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBaselineOffset() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontVariation.Bind_get_baseline_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsFontVariation() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsFontVariation() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsFont() Font.GD { return *((*Font.GD)(unsafe.Pointer(&self))) }
func (self Go) AsFont() Font.Go { return *((*Font.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsFont(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsFont(), name)
	}
}
func init() {classdb.Register("FontVariation", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
