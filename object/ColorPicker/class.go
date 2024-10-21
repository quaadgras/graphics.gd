package ColorPicker

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VBoxContainer"
import "grow.graphics/gd/object/BoxContainer"
import "grow.graphics/gd/object/Container"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A widget that provides an interface for selecting or modifying a color. It can optionally provide functionalities like a color sampler (eyedropper), color modes, and presets.
[b]Note:[/b] This control is the color picker widget itself. You can use a [ColorPickerButton] instead if you need a button that brings up a [ColorPicker] in a popup.

*/
type Simple [1]classdb.ColorPicker
func (self Simple) SetPickColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPickColor(color)
}
func (self Simple) GetPickColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetPickColor())
}
func (self Simple) SetDeferredMode(mode bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDeferredMode(mode)
}
func (self Simple) IsDeferredMode() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDeferredMode())
}
func (self Simple) SetColorMode(color_mode classdb.ColorPickerColorModeType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColorMode(color_mode)
}
func (self Simple) GetColorMode() classdb.ColorPickerColorModeType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ColorPickerColorModeType(Expert(self).GetColorMode())
}
func (self Simple) SetEditAlpha(show bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEditAlpha(show)
}
func (self Simple) IsEditingAlpha() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsEditingAlpha())
}
func (self Simple) SetCanAddSwatches(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCanAddSwatches(enabled)
}
func (self Simple) AreSwatchesEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).AreSwatchesEnabled())
}
func (self Simple) SetPresetsVisible(visible bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPresetsVisible(visible)
}
func (self Simple) ArePresetsVisible() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).ArePresetsVisible())
}
func (self Simple) SetModesVisible(visible bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetModesVisible(visible)
}
func (self Simple) AreModesVisible() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).AreModesVisible())
}
func (self Simple) SetSamplerVisible(visible bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSamplerVisible(visible)
}
func (self Simple) IsSamplerVisible() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSamplerVisible())
}
func (self Simple) SetSlidersVisible(visible bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSlidersVisible(visible)
}
func (self Simple) AreSlidersVisible() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).AreSlidersVisible())
}
func (self Simple) SetHexVisible(visible bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHexVisible(visible)
}
func (self Simple) IsHexVisible() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsHexVisible())
}
func (self Simple) AddPreset(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddPreset(color)
}
func (self Simple) ErasePreset(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ErasePreset(color)
}
func (self Simple) GetPresets() gd.PackedColorArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedColorArray(Expert(self).GetPresets(gc))
}
func (self Simple) AddRecentPreset(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddRecentPreset(color)
}
func (self Simple) EraseRecentPreset(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EraseRecentPreset(color)
}
func (self Simple) GetRecentPresets() gd.PackedColorArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedColorArray(Expert(self).GetRecentPresets(gc))
}
func (self Simple) SetPickerShape(shape classdb.ColorPickerPickerShapeType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPickerShape(shape)
}
func (self Simple) GetPickerShape() classdb.ColorPickerPickerShapeType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ColorPickerPickerShapeType(Expert(self).GetPickerShape())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ColorPicker
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetPickColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_set_pick_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPickColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_get_pick_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDeferredMode(mode bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_set_deferred_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDeferredMode() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_is_deferred_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColorMode(color_mode classdb.ColorPickerColorModeType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_set_color_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColorMode() classdb.ColorPickerColorModeType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ColorPickerColorModeType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_get_color_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEditAlpha(show bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, show)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_set_edit_alpha, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEditingAlpha() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_is_editing_alpha, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCanAddSwatches(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_set_can_add_swatches, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) AreSwatchesEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_are_swatches_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPresetsVisible(visible bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_set_presets_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) ArePresetsVisible() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_are_presets_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetModesVisible(visible bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_set_modes_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) AreModesVisible() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_are_modes_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSamplerVisible(visible bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_set_sampler_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSamplerVisible() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_is_sampler_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSlidersVisible(visible bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_set_sliders_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) AreSlidersVisible() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_are_sliders_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHexVisible(visible bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_set_hex_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsHexVisible() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_is_hex_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds the given color to a list of color presets. The presets are displayed in the color picker and the user will be able to select them.
[b]Note:[/b] The presets list is only for [i]this[/i] color picker.
*/
//go:nosplit
func (self class) AddPreset(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_add_preset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the given color from the list of color presets of this color picker.
*/
//go:nosplit
func (self class) ErasePreset(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_erase_preset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the list of colors in the presets of the color picker.
*/
//go:nosplit
func (self class) GetPresets(ctx gd.Lifetime) gd.PackedColorArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_get_presets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedColorArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Adds the given color to a list of color recent presets so that it can be picked later. Recent presets are the colors that were picked recently, a new preset is automatically created and added to recent presets when you pick a new color.
[b]Note:[/b] The recent presets list is only for [i]this[/i] color picker.
*/
//go:nosplit
func (self class) AddRecentPreset(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_add_recent_preset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the given color from the list of color recent presets of this color picker.
*/
//go:nosplit
func (self class) EraseRecentPreset(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_erase_recent_preset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the list of colors in the recent presets of the color picker.
*/
//go:nosplit
func (self class) GetRecentPresets(ctx gd.Lifetime) gd.PackedColorArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_get_recent_presets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedColorArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPickerShape(shape classdb.ColorPickerPickerShapeType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_set_picker_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPickerShape() classdb.ColorPickerPickerShapeType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ColorPickerPickerShapeType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPicker.Bind_get_picker_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsColorPicker() Expert { return self[0].AsColorPicker() }


//go:nosplit
func (self Simple) AsColorPicker() Simple { return self[0].AsColorPicker() }


//go:nosplit
func (self class) AsVBoxContainer() VBoxContainer.Expert { return self[0].AsVBoxContainer() }


//go:nosplit
func (self Simple) AsVBoxContainer() VBoxContainer.Simple { return self[0].AsVBoxContainer() }


//go:nosplit
func (self class) AsBoxContainer() BoxContainer.Expert { return self[0].AsBoxContainer() }


//go:nosplit
func (self Simple) AsBoxContainer() BoxContainer.Simple { return self[0].AsBoxContainer() }


//go:nosplit
func (self class) AsContainer() Container.Expert { return self[0].AsContainer() }


//go:nosplit
func (self Simple) AsContainer() Container.Simple { return self[0].AsContainer() }


//go:nosplit
func (self class) AsControl() Control.Expert { return self[0].AsControl() }


//go:nosplit
func (self Simple) AsControl() Control.Simple { return self[0].AsControl() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("ColorPicker", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type ColorModeType = classdb.ColorPickerColorModeType

const (
/*Allows editing the color with Red/Green/Blue sliders.*/
	ModeRgb ColorModeType = 0
/*Allows editing the color with Hue/Saturation/Value sliders.*/
	ModeHsv ColorModeType = 1
/*Allows the color R, G, B component values to go beyond 1.0, which can be used for certain special operations that require it (like tinting without darkening or rendering sprites in HDR).*/
	ModeRaw ColorModeType = 2
/*Allows editing the color with Hue/Saturation/Lightness sliders.
OKHSL is a new color space similar to HSL but that better match perception by leveraging the Oklab color space which is designed to be simple to use, while doing a good job at predicting perceived lightness, chroma and hue.
[url=https://bottosson.github.io/posts/colorpicker/]Okhsv and Okhsl color spaces[/url]*/
	ModeOkhsl ColorModeType = 3
)
type PickerShapeType = classdb.ColorPickerPickerShapeType

const (
/*HSV Color Model rectangle color space.*/
	ShapeHsvRectangle PickerShapeType = 0
/*HSV Color Model rectangle color space with a wheel.*/
	ShapeHsvWheel PickerShapeType = 1
/*HSV Color Model circle color space. Use Saturation as a radius.*/
	ShapeVhsCircle PickerShapeType = 2
/*HSL OK Color Model circle color space.*/
	ShapeOkhslCircle PickerShapeType = 3
/*The color space shape and the shape select button are hidden. Can't be selected from the shapes popup.*/
	ShapeNone PickerShapeType = 4
)
