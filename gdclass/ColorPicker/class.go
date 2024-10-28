package ColorPicker

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VBoxContainer"
import "grow.graphics/gd/gdclass/BoxContainer"
import "grow.graphics/gd/gdclass/Container"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
A widget that provides an interface for selecting or modifying a color. It can optionally provide functionalities like a color sampler (eyedropper), color modes, and presets.
[b]Note:[/b] This control is the color picker widget itself. You can use a [ColorPickerButton] instead if you need a button that brings up a [ColorPicker] in a popup.

*/
type Go [1]classdb.ColorPicker

/*
Adds the given color to a list of color presets. The presets are displayed in the color picker and the user will be able to select them.
[b]Note:[/b] The presets list is only for [i]this[/i] color picker.
*/
func (self Go) AddPreset(color gd.Color) {
	class(self).AddPreset(color)
}

/*
Removes the given color from the list of color presets of this color picker.
*/
func (self Go) ErasePreset(color gd.Color) {
	class(self).ErasePreset(color)
}

/*
Returns the list of colors in the presets of the color picker.
*/
func (self Go) GetPresets() []gd.Color {
	return []gd.Color(class(self).GetPresets().AsSlice())
}

/*
Adds the given color to a list of color recent presets so that it can be picked later. Recent presets are the colors that were picked recently, a new preset is automatically created and added to recent presets when you pick a new color.
[b]Note:[/b] The recent presets list is only for [i]this[/i] color picker.
*/
func (self Go) AddRecentPreset(color gd.Color) {
	class(self).AddRecentPreset(color)
}

/*
Removes the given color from the list of color recent presets of this color picker.
*/
func (self Go) EraseRecentPreset(color gd.Color) {
	class(self).EraseRecentPreset(color)
}

/*
Returns the list of colors in the recent presets of the color picker.
*/
func (self Go) GetRecentPresets() []gd.Color {
	return []gd.Color(class(self).GetRecentPresets().AsSlice())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ColorPicker
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ColorPicker"))
	return Go{classdb.ColorPicker(object)}
}

func (self Go) Color() gd.Color {
		return gd.Color(class(self).GetPickColor())
}

func (self Go) SetColor(value gd.Color) {
	class(self).SetPickColor(value)
}

func (self Go) EditAlpha() bool {
		return bool(class(self).IsEditingAlpha())
}

func (self Go) SetEditAlpha(value bool) {
	class(self).SetEditAlpha(value)
}

func (self Go) ColorMode() classdb.ColorPickerColorModeType {
		return classdb.ColorPickerColorModeType(class(self).GetColorMode())
}

func (self Go) SetColorMode(value classdb.ColorPickerColorModeType) {
	class(self).SetColorMode(value)
}

func (self Go) DeferredMode() bool {
		return bool(class(self).IsDeferredMode())
}

func (self Go) SetDeferredMode(value bool) {
	class(self).SetDeferredMode(value)
}

func (self Go) PickerShape() classdb.ColorPickerPickerShapeType {
		return classdb.ColorPickerPickerShapeType(class(self).GetPickerShape())
}

func (self Go) SetPickerShape(value classdb.ColorPickerPickerShapeType) {
	class(self).SetPickerShape(value)
}

func (self Go) CanAddSwatches() bool {
		return bool(class(self).AreSwatchesEnabled())
}

func (self Go) SetCanAddSwatches(value bool) {
	class(self).SetCanAddSwatches(value)
}

func (self Go) SamplerVisible() bool {
		return bool(class(self).IsSamplerVisible())
}

func (self Go) SetSamplerVisible(value bool) {
	class(self).SetSamplerVisible(value)
}

func (self Go) ColorModesVisible() bool {
		return bool(class(self).AreModesVisible())
}

func (self Go) SetColorModesVisible(value bool) {
	class(self).SetModesVisible(value)
}

func (self Go) SlidersVisible() bool {
		return bool(class(self).AreSlidersVisible())
}

func (self Go) SetSlidersVisible(value bool) {
	class(self).SetSlidersVisible(value)
}

func (self Go) HexVisible() bool {
		return bool(class(self).IsHexVisible())
}

func (self Go) SetHexVisible(value bool) {
	class(self).SetHexVisible(value)
}

func (self Go) PresetsVisible() bool {
		return bool(class(self).ArePresetsVisible())
}

func (self Go) SetPresetsVisible(value bool) {
	class(self).SetPresetsVisible(value)
}

//go:nosplit
func (self class) SetPickColor(color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_set_pick_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPickColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_get_pick_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDeferredMode(mode bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_set_deferred_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDeferredMode() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_is_deferred_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColorMode(color_mode classdb.ColorPickerColorModeType)  {
	var frame = callframe.New()
	callframe.Arg(frame, color_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_set_color_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColorMode() classdb.ColorPickerColorModeType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ColorPickerColorModeType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_get_color_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEditAlpha(show bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, show)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_set_edit_alpha, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEditingAlpha() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_is_editing_alpha, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCanAddSwatches(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_set_can_add_swatches, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) AreSwatchesEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_are_swatches_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPresetsVisible(visible bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_set_presets_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) ArePresetsVisible() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_are_presets_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetModesVisible(visible bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_set_modes_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) AreModesVisible() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_are_modes_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSamplerVisible(visible bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_set_sampler_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSamplerVisible() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_is_sampler_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSlidersVisible(visible bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_set_sliders_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) AreSlidersVisible() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_are_sliders_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHexVisible(visible bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_set_hex_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsHexVisible() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_is_hex_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_add_preset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the given color from the list of color presets of this color picker.
*/
//go:nosplit
func (self class) ErasePreset(color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_erase_preset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the list of colors in the presets of the color picker.
*/
//go:nosplit
func (self class) GetPresets() gd.PackedColorArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_get_presets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedColorArray](r_ret.Get())
	frame.Free()
	return ret
}
/*
Adds the given color to a list of color recent presets so that it can be picked later. Recent presets are the colors that were picked recently, a new preset is automatically created and added to recent presets when you pick a new color.
[b]Note:[/b] The recent presets list is only for [i]this[/i] color picker.
*/
//go:nosplit
func (self class) AddRecentPreset(color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_add_recent_preset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the given color from the list of color recent presets of this color picker.
*/
//go:nosplit
func (self class) EraseRecentPreset(color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_erase_recent_preset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the list of colors in the recent presets of the color picker.
*/
//go:nosplit
func (self class) GetRecentPresets() gd.PackedColorArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_get_recent_presets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedColorArray](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPickerShape(shape classdb.ColorPickerPickerShapeType)  {
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_set_picker_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPickerShape() classdb.ColorPickerPickerShapeType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ColorPickerPickerShapeType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPicker.Bind_get_picker_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnColorChanged(cb func(color gd.Color)) {
	self[0].AsObject().Connect(gd.NewStringName("color_changed"), gd.NewCallable(cb), 0)
}


func (self Go) OnPresetAdded(cb func(color gd.Color)) {
	self[0].AsObject().Connect(gd.NewStringName("preset_added"), gd.NewCallable(cb), 0)
}


func (self Go) OnPresetRemoved(cb func(color gd.Color)) {
	self[0].AsObject().Connect(gd.NewStringName("preset_removed"), gd.NewCallable(cb), 0)
}


func (self class) AsColorPicker() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsColorPicker() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVBoxContainer() VBoxContainer.GD { return *((*VBoxContainer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVBoxContainer() VBoxContainer.Go { return *((*VBoxContainer.Go)(unsafe.Pointer(&self))) }
func (self class) AsBoxContainer() BoxContainer.GD { return *((*BoxContainer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsBoxContainer() BoxContainer.Go { return *((*BoxContainer.Go)(unsafe.Pointer(&self))) }
func (self class) AsContainer() Container.GD { return *((*Container.GD)(unsafe.Pointer(&self))) }
func (self Go) AsContainer() Container.Go { return *((*Container.Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVBoxContainer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVBoxContainer(), name)
	}
}
func init() {classdb.Register("ColorPicker", func(ptr gd.Object) any { return classdb.ColorPicker(ptr) })}
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
