package TextureProgressBar

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Range"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
TextureProgressBar works like [ProgressBar], but uses up to 3 textures instead of Godot's [Theme] resource. It can be used to create horizontal, vertical and radial progress bars.

*/
type Go [1]classdb.TextureProgressBar
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TextureProgressBar
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TextureProgressBar"))
	return Go{classdb.TextureProgressBar(object)}
}

func (self Go) FillMode() int {
		return int(int(class(self).GetFillMode()))
}

func (self Go) SetFillMode(value int) {
	class(self).SetFillMode(gd.Int(value))
}

func (self Go) RadialInitialAngle() float64 {
		return float64(float64(class(self).GetRadialInitialAngle()))
}

func (self Go) SetRadialInitialAngle(value float64) {
	class(self).SetRadialInitialAngle(gd.Float(value))
}

func (self Go) RadialFillDegrees() float64 {
		return float64(float64(class(self).GetFillDegrees()))
}

func (self Go) SetRadialFillDegrees(value float64) {
	class(self).SetFillDegrees(gd.Float(value))
}

func (self Go) RadialCenterOffset() gd.Vector2 {
		return gd.Vector2(class(self).GetRadialCenterOffset())
}

func (self Go) SetRadialCenterOffset(value gd.Vector2) {
	class(self).SetRadialCenterOffset(value)
}

func (self Go) NinePatchStretch() bool {
		return bool(class(self).GetNinePatchStretch())
}

func (self Go) SetNinePatchStretch(value bool) {
	class(self).SetNinePatchStretch(value)
}

func (self Go) StretchMarginLeft() int {
		return int(int(class(self).GetStretchMargin(0)))
}

func (self Go) SetStretchMarginLeft(value int) {
	class(self).SetStretchMargin(0, gd.Int(value))
}

func (self Go) StretchMarginTop() int {
		return int(int(class(self).GetStretchMargin(1)))
}

func (self Go) SetStretchMarginTop(value int) {
	class(self).SetStretchMargin(1, gd.Int(value))
}

func (self Go) StretchMarginRight() int {
		return int(int(class(self).GetStretchMargin(2)))
}

func (self Go) SetStretchMarginRight(value int) {
	class(self).SetStretchMargin(2, gd.Int(value))
}

func (self Go) StretchMarginBottom() int {
		return int(int(class(self).GetStretchMargin(3)))
}

func (self Go) SetStretchMarginBottom(value int) {
	class(self).SetStretchMargin(3, gd.Int(value))
}

func (self Go) TextureUnder() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetUnderTexture())
}

func (self Go) SetTextureUnder(value gdclass.Texture2D) {
	class(self).SetUnderTexture(value)
}

func (self Go) TextureOver() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetOverTexture())
}

func (self Go) SetTextureOver(value gdclass.Texture2D) {
	class(self).SetOverTexture(value)
}

func (self Go) TextureProgress() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetProgressTexture())
}

func (self Go) SetTextureProgress(value gdclass.Texture2D) {
	class(self).SetProgressTexture(value)
}

func (self Go) TextureProgressOffset() gd.Vector2 {
		return gd.Vector2(class(self).GetTextureProgressOffset())
}

func (self Go) SetTextureProgressOffset(value gd.Vector2) {
	class(self).SetTextureProgressOffset(value)
}

func (self Go) TintUnder() gd.Color {
		return gd.Color(class(self).GetTintUnder())
}

func (self Go) SetTintUnder(value gd.Color) {
	class(self).SetTintUnder(value)
}

func (self Go) TintOver() gd.Color {
		return gd.Color(class(self).GetTintOver())
}

func (self Go) SetTintOver(value gd.Color) {
	class(self).SetTintOver(value)
}

func (self Go) TintProgress() gd.Color {
		return gd.Color(class(self).GetTintProgress())
}

func (self Go) SetTintProgress(value gd.Color) {
	class(self).SetTintProgress(value)
}

//go:nosplit
func (self class) SetUnderTexture(tex gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(tex[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_under_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUnderTexture() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_under_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetProgressTexture(tex gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(tex[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_progress_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetProgressTexture() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_progress_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOverTexture(tex gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(tex[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_over_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOverTexture() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_over_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFillMode(mode gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_fill_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFillMode() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_fill_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTintUnder(tint gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, tint)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_tint_under, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTintUnder() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_tint_under, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTintProgress(tint gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, tint)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_tint_progress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTintProgress() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_tint_progress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTintOver(tint gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, tint)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_tint_over, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTintOver() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_tint_over, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextureProgressOffset(offset gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_texture_progress_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextureProgressOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_texture_progress_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRadialInitialAngle(mode gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_radial_initial_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRadialInitialAngle() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_radial_initial_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRadialCenterOffset(mode gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_radial_center_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRadialCenterOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_radial_center_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFillDegrees(mode gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_fill_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFillDegrees() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_fill_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the stretch margin with the specified index. See [member stretch_margin_bottom] and related properties.
*/
//go:nosplit
func (self class) SetStretchMargin(margin gd.Side, value gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_stretch_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the stretch margin with the specified index. See [member stretch_margin_bottom] and related properties.
*/
//go:nosplit
func (self class) GetStretchMargin(margin gd.Side) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_stretch_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNinePatchStretch(stretch bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, stretch)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_set_nine_patch_stretch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNinePatchStretch() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureProgressBar.Bind_get_nine_patch_stretch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTextureProgressBar() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTextureProgressBar() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRange() Range.GD { return *((*Range.GD)(unsafe.Pointer(&self))) }
func (self Go) AsRange() Range.Go { return *((*Range.Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRange(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRange(), name)
	}
}
func init() {classdb.Register("TextureProgressBar", func(ptr gd.Object) any { return classdb.TextureProgressBar(ptr) })}
type FillMode = classdb.TextureProgressBarFillMode

const (
/*The [member texture_progress] fills from left to right.*/
	FillLeftToRight FillMode = 0
/*The [member texture_progress] fills from right to left.*/
	FillRightToLeft FillMode = 1
/*The [member texture_progress] fills from top to bottom.*/
	FillTopToBottom FillMode = 2
/*The [member texture_progress] fills from bottom to top.*/
	FillBottomToTop FillMode = 3
/*Turns the node into a radial bar. The [member texture_progress] fills clockwise. See [member radial_center_offset], [member radial_initial_angle] and [member radial_fill_degrees] to control the way the bar fills up.*/
	FillClockwise FillMode = 4
/*Turns the node into a radial bar. The [member texture_progress] fills counterclockwise. See [member radial_center_offset], [member radial_initial_angle] and [member radial_fill_degrees] to control the way the bar fills up.*/
	FillCounterClockwise FillMode = 5
/*The [member texture_progress] fills from the center, expanding both towards the left and the right.*/
	FillBilinearLeftAndRight FillMode = 6
/*The [member texture_progress] fills from the center, expanding both towards the top and the bottom.*/
	FillBilinearTopAndBottom FillMode = 7
/*Turns the node into a radial bar. The [member texture_progress] fills radially from the center, expanding both clockwise and counterclockwise. See [member radial_center_offset], [member radial_initial_angle] and [member radial_fill_degrees] to control the way the bar fills up.*/
	FillClockwiseAndCounterClockwise FillMode = 8
)
