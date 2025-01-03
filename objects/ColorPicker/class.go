package ColorPicker

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/VBoxContainer"
import "graphics.gd/objects/BoxContainer"
import "graphics.gd/objects/Container"
import "graphics.gd/objects/Control"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Color"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A widget that provides an interface for selecting or modifying a color. It can optionally provide functionalities like a color sampler (eyedropper), color modes, and presets.
[b]Note:[/b] This control is the color picker widget itself. You can use a [ColorPickerButton] instead if you need a button that brings up a [ColorPicker] in a popup.
*/
type Instance [1]classdb.ColorPicker
type Any interface {
	gd.IsClass
	AsColorPicker() Instance
}

/*
Adds the given color to a list of color presets. The presets are displayed in the color picker and the user will be able to select them.
[b]Note:[/b] The presets list is only for [i]this[/i] color picker.
*/
func (self Instance) AddPreset(color Color.RGBA) {
	class(self).AddPreset(gd.Color(color))
}

/*
Removes the given color from the list of color presets of this color picker.
*/
func (self Instance) ErasePreset(color Color.RGBA) {
	class(self).ErasePreset(gd.Color(color))
}

/*
Returns the list of colors in the presets of the color picker.
*/
func (self Instance) GetPresets() []Color.RGBA {
	return []Color.RGBA(class(self).GetPresets().AsSlice())
}

/*
Adds the given color to a list of color recent presets so that it can be picked later. Recent presets are the colors that were picked recently, a new preset is automatically created and added to recent presets when you pick a new color.
[b]Note:[/b] The recent presets list is only for [i]this[/i] color picker.
*/
func (self Instance) AddRecentPreset(color Color.RGBA) {
	class(self).AddRecentPreset(gd.Color(color))
}

/*
Removes the given color from the list of color recent presets of this color picker.
*/
func (self Instance) EraseRecentPreset(color Color.RGBA) {
	class(self).EraseRecentPreset(gd.Color(color))
}

/*
Returns the list of colors in the recent presets of the color picker.
*/
func (self Instance) GetRecentPresets() []Color.RGBA {
	return []Color.RGBA(class(self).GetRecentPresets().AsSlice())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ColorPicker

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ColorPicker"))
	return Instance{*(*classdb.ColorPicker)(unsafe.Pointer(&object))}
}

func (self Instance) Color() Color.RGBA {
	return Color.RGBA(class(self).GetPickColor())
}

func (self Instance) SetColor(value Color.RGBA) {
	class(self).SetPickColor(gd.Color(value))
}

func (self Instance) EditAlpha() bool {
	return bool(class(self).IsEditingAlpha())
}

func (self Instance) SetEditAlpha(value bool) {
	class(self).SetEditAlpha(value)
}

func (self Instance) ColorMode() classdb.ColorPickerColorModeType {
	return classdb.ColorPickerColorModeType(class(self).GetColorMode())
}

func (self Instance) SetColorMode(value classdb.ColorPickerColorModeType) {
	class(self).SetColorMode(value)
}

func (self Instance) DeferredMode() bool {
	return bool(class(self).IsDeferredMode())
}

func (self Instance) SetDeferredMode(value bool) {
	class(self).SetDeferredMode(value)
}

func (self Instance) PickerShape() classdb.ColorPickerPickerShapeType {
	return classdb.ColorPickerPickerShapeType(class(self).GetPickerShape())
}

func (self Instance) SetPickerShape(value classdb.ColorPickerPickerShapeType) {
	class(self).SetPickerShape(value)
}

func (self Instance) CanAddSwatches() bool {
	return bool(class(self).AreSwatchesEnabled())
}

func (self Instance) SetCanAddSwatches(value bool) {
	class(self).SetCanAddSwatches(value)
}

func (self Instance) SamplerVisible() bool {
	return bool(class(self).IsSamplerVisible())
}

func (self Instance) SetSamplerVisible(value bool) {
	class(self).SetSamplerVisible(value)
}

func (self Instance) ColorModesVisible() bool {
	return bool(class(self).AreModesVisible())
}

func (self Instance) SetColorModesVisible(value bool) {
	class(self).SetModesVisible(value)
}

func (self Instance) SlidersVisible() bool {
	return bool(class(self).AreSlidersVisible())
}

func (self Instance) SetSlidersVisible(value bool) {
	class(self).SetSlidersVisible(value)
}

func (self Instance) HexVisible() bool {
	return bool(class(self).IsHexVisible())
}

func (self Instance) SetHexVisible(value bool) {
	class(self).SetHexVisible(value)
}

func (self Instance) PresetsVisible() bool {
	return bool(class(self).ArePresetsVisible())
}

func (self Instance) SetPresetsVisible(value bool) {
	class(self).SetPresetsVisible(value)
}

//go:nosplit
func (self class) SetPickColor(color gd.Color) {
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
func (self class) SetDeferredMode(mode bool) {
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
func (self class) SetColorMode(color_mode classdb.ColorPickerColorModeType) {
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
func (self class) SetEditAlpha(show bool) {
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
func (self class) SetCanAddSwatches(enabled bool) {
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
func (self class) SetPresetsVisible(visible bool) {
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
func (self class) SetModesVisible(visible bool) {
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
func (self class) SetSamplerVisible(visible bool) {
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
func (self class) SetSlidersVisible(visible bool) {
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
func (self class) SetHexVisible(visible bool) {
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
func (self class) AddPreset(color gd.Color) {
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
func (self class) ErasePreset(color gd.Color) {
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
	var ret = pointers.New[gd.PackedColorArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Adds the given color to a list of color recent presets so that it can be picked later. Recent presets are the colors that were picked recently, a new preset is automatically created and added to recent presets when you pick a new color.
[b]Note:[/b] The recent presets list is only for [i]this[/i] color picker.
*/
//go:nosplit
func (self class) AddRecentPreset(color gd.Color) {
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
func (self class) EraseRecentPreset(color gd.Color) {
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
	var ret = pointers.New[gd.PackedColorArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPickerShape(shape classdb.ColorPickerPickerShapeType) {
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
func (self Instance) OnColorChanged(cb func(color Color.RGBA)) {
	self[0].AsObject().Connect(gd.NewStringName("color_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPresetAdded(cb func(color Color.RGBA)) {
	self[0].AsObject().Connect(gd.NewStringName("preset_added"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPresetRemoved(cb func(color Color.RGBA)) {
	self[0].AsObject().Connect(gd.NewStringName("preset_removed"), gd.NewCallable(cb), 0)
}

func (self class) AsColorPicker() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsColorPicker() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsVBoxContainer() VBoxContainer.Advanced {
	return *((*VBoxContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVBoxContainer() VBoxContainer.Instance {
	return *((*VBoxContainer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsBoxContainer() BoxContainer.Advanced {
	return *((*BoxContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsBoxContainer() BoxContainer.Instance {
	return *((*BoxContainer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsContainer() Container.Advanced {
	return *((*Container.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsContainer() Container.Instance {
	return *((*Container.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(self.AsVBoxContainer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVBoxContainer(), name)
	}
}
func init() {
	classdb.Register("ColorPicker", func(ptr gd.Object) any { return [1]classdb.ColorPicker{*(*classdb.ColorPicker)(unsafe.Pointer(&ptr))} })
}

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
